package io.github.shibukawa.vkmem;

import java.io.BufferedInputStream;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.net.Socket;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.StandardCopyOption;
import java.nio.file.attribute.PosixFilePermissions;
import java.time.Duration;
import java.util.ArrayList;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Locale;
import java.util.Map;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;

/**
 * A running vkmem-server process, or an isolated Valkey server started from a
 * {@link Snapshot}.
 *
 * <pre>{@code
 * try (VkmemServer template = VkmemServer.builder().start()) {
 *     template.command("SET", "seed", "yes");
 *     try (Snapshot snapshot = template.snapshot(4);
 *          Fork fork = snapshot.fork()) {
 *         // point Jedis, Lettuce or valkey-java at fork.host()/fork.port()
 *     }
 * }
 * }</pre>
 *
 * The server is located from the {@code vkmem.server.bin} system property,
 * the {@code VKMEM_SERVER_BIN} environment variable, or the classpath
 * resource shipped in the {@code vkmem-server-binaries} artifact.
 */
public class VkmemServer implements AutoCloseable {
    private final Control control;
    private final boolean ownsProcess;
    private final String id;
    private final String host;
    private final int port;
    private final String unixSocket;
    private final String version;
    private final int pid;
    private final String valkeyVersion;
    private volatile boolean closed;

    VkmemServer(Control control, Map<String, Object> endpoint, boolean ownsProcess) {
        this.control = control;
        this.ownsProcess = ownsProcess;
        this.id = String.valueOf(endpoint.getOrDefault("id", "template"));
        this.host = String.valueOf(endpoint.getOrDefault("host", "127.0.0.1"));
        this.port = number(endpoint.get("port"), "port").intValue();
        this.unixSocket = endpoint.get("unix") == null ? null : String.valueOf(endpoint.get("unix"));
        this.version = String.valueOf(endpoint.getOrDefault("version", ""));
        this.pid = number(endpoint.get("pid"), "pid").intValue();
        this.valkeyVersion = String.valueOf(endpoint.getOrDefault("valkey", ""));
    }

    private static Number number(Object value, String name) {
        if (!(value instanceof Number)) throw new VkmemException("vkmem: readiness record has no " + name);
        return (Number) value;
    }

    public static Builder builder() {
        return new Builder();
    }

    /** The template server; provided for pgmem-style setup code. */
    public VkmemServer template() {
        return this;
    }

    /** Controller id; {@code template} for the server returned by {@link Builder#start()}. */
    public String id() {
        return id;
    }

    /** Always {@code 127.0.0.1}. */
    public String host() {
        return host;
    }

    public int port() {
        return port;
    }

    /** {@code host:port}. */
    public String address() {
        return host + ":" + port;
    }

    /** {@code redis://127.0.0.1:port}, accepted by most clients. */
    public String url() {
        return "redis://" + address();
    }

    /** Alias of {@link #url()}, suitable for Redis-compatible clients. */
    public String dsn() {
        return url();
    }

    /** Unix socket path, or {@code null} when disabled. */
    public String unixSocket() {
        return unixSocket;
    }

    /** The vkmem-server version. */
    public String version() {
        return version;
    }

    /** The server process id. */
    public int pid() {
        return pid;
    }

    /** The Valkey release compiled into the server. */
    public String valkeyVersion() {
        return valkeyVersion;
    }

    /** Checkpoints the keyspace and copies it into an independent snapshot. */
    public Snapshot snapshot() {
        return snapshot(0);
    }

    /**
     * @param maxForks forks alive at once before {@link Snapshot#fork()} blocks;
     *                 0 uses the available processor count
     */
    public Snapshot snapshot(int maxForks) {
        return snapshot(maxForks, Duration.ofSeconds(30));
    }

    /**
     * Creates a data snapshot. The timeout bounds synchronous snapshot creation;
     * {@code null} waits indefinitely.
     */
    public Snapshot snapshot(int maxForks, Duration timeout) {
        Map<String, Object> fields = new LinkedHashMap<>();
        fields.put("server", id);
        if (maxForks > 0) fields.put("max_forks", maxForks);
        if (timeout != null) fields.put("timeout_ms", timeout.toMillis());
        Map<String, Object> result = control.request("snapshot", fields);
        return new Snapshot(control, String.valueOf(result.get("snapshot")), this);
    }

    /**
     * Sends one command over a fresh TCP connection and returns the decoded reply:
     * a {@link String} for simple and bulk strings, a {@link Long} for integers,
     * {@code null} for nil, or a {@link List} for arrays. Use a real client for
     * application code.
     */
    public Object command(String... args) {
        try (Socket socket = new Socket()) {
            socket.connect(new InetSocketAddress(host(), port()), 5000);
            socket.setSoTimeout(30000);
            OutputStream out = socket.getOutputStream();
            StringBuilder request = new StringBuilder("*").append(args.length).append("\r\n");
            for (String arg : args) {
                byte[] bytes = arg.getBytes(StandardCharsets.UTF_8);
                request.append('$').append(bytes.length).append("\r\n").append(arg).append("\r\n");
            }
            out.write(request.toString().getBytes(StandardCharsets.UTF_8));
            out.flush();
            return readReply(new BufferedInputStream(socket.getInputStream()));
        } catch (IOException e) {
            throw new VkmemException("vkmem: " + String.join(" ", args) + ": " + e.getMessage(), e);
        }
    }

    /** FLUSHALL: the usual reset between tests. */
    public void flushAll() {
        command("FLUSHALL");
    }

    private static Object readReply(InputStream in) throws IOException {
        int type = in.read();
        String line = readLine(in);
        switch (type) {
            case '+':
                return line;
            case '-':
                throw new VkmemException("vkmem: " + line);
            case ':':
                return Long.parseLong(line);
            case '$': {
                int length = Integer.parseInt(line);
                if (length < 0) return null;
                byte[] bytes = in.readNBytes(length);
                readLine(in);
                return new String(bytes, StandardCharsets.UTF_8);
            }
            case '*': {
                int length = Integer.parseInt(line);
                if (length < 0) return null;
                List<Object> result = new ArrayList<>(length);
                for (int i = 0; i < length; i++) result.add(readReply(in));
                return result;
            }
            default:
                throw new IOException("unexpected reply byte " + type);
        }
    }

    private static String readLine(InputStream in) throws IOException {
        StringBuilder result = new StringBuilder();
        int c;
        while ((c = in.read()) >= 0) {
            if (c == '\r') {
                in.read();
                return result.toString();
            }
            result.append((char) c);
        }
        throw new IOException("connection closed");
    }

    /** Stops this fork, or shuts down the controller for the template. */
    @Override
    public void close() {
        if (closed) return;
        closed = true;
        if (ownsProcess) {
            control.close();
        } else {
            try {
                control.request("close", Map.of("server", id));
            } catch (ServerExitedException ignored) {
                // The controller may already have exited.
            }
        }
    }

    public boolean isClosed() {
        return closed;
    }

    /** Locates the vkmem-server binary, extracting the bundled one when needed. */
    public static Path resolveBinary(Path explicit) throws IOException {
        if (explicit != null) return explicit;
        String property = System.getProperty("vkmem.server.bin");
        if (property != null && !property.isEmpty()) return Path.of(property);
        String environment = System.getenv("VKMEM_SERVER_BIN");
        if (environment != null && !environment.isEmpty()) return Path.of(environment);
        String platform = platformId();
        String name = platform.startsWith("windows") ? "vkmem-server.exe" : "vkmem-server";
        String resource = "/vkmem/bin/" + platform + "/" + name;
        try (InputStream in = VkmemServer.class.getResourceAsStream(resource)) {
            if (in == null) {
                throw new VkmemException("vkmem: no bundled binary for " + platform
                        + " (add io.github.shibukawa.vkmem:vkmem-server-binaries with classifier " + platform
                        + " or set VKMEM_SERVER_BIN)");
            }
            Path dir = Files.createTempDirectory("vkmem-server");
            Path binary = dir.resolve(name);
            Files.copy(in, binary, StandardCopyOption.REPLACE_EXISTING);
            if (!platform.startsWith("windows")) {
                Files.setPosixFilePermissions(binary, PosixFilePermissions.fromString("rwx------"));
            }
            binary.toFile().deleteOnExit();
            dir.toFile().deleteOnExit();
            return binary;
        }
    }

    static String platformId() {
        String os = System.getProperty("os.name", "").toLowerCase(Locale.ROOT);
        String arch = System.getProperty("os.arch", "").toLowerCase(Locale.ROOT);
        String operatingSystem = os.contains("mac") || os.contains("darwin") ? "darwin"
                : os.contains("win") ? "windows" : "linux";
        String architecture = arch.contains("aarch64") || arch.contains("arm64") ? "arm64" : "amd64";
        return operatingSystem + "-" + architecture;
    }

    /** Builder returned by {@link VkmemServer#builder()}; {@link #start()} launches the server. */
    public static final class Builder {
        private int port;
        private String unixSocket;
        private boolean noUnixSocket;
        private final List<String> args = new ArrayList<>();
        private Path binary;
        private boolean quiet = true;
        private Duration startupTimeout = Duration.ofSeconds(30);

        Builder() {}

        /** TCP port; 0 (the default) picks a free one. */
        public Builder port(int port) {
            this.port = port;
            return this;
        }

        /** Unix socket path (default: a generated temp path). */
        public Builder unixSocket(String path) {
            this.unixSocket = path;
            return this;
        }

        /** Serve TCP only. */
        public Builder noUnixSocket() {
            this.noUnixSocket = true;
            return this;
        }

        /** Extra valkey-server arguments, e.g. {@code args("--maxmemory", "64mb")}. */
        public Builder args(String... valkeyArgs) {
            Collections.addAll(args, valkeyArgs);
            return this;
        }

        public Builder binary(Path binary) {
            this.binary = binary;
            return this;
        }

        /** Forward the Valkey log to stderr (dropped by default). */
        public Builder quiet(boolean quiet) {
            this.quiet = quiet;
            return this;
        }

        public Builder startupTimeout(Duration timeout) {
            this.startupTimeout = timeout;
            return this;
        }

        /** Starts the server and waits for its ready line. */
        public VkmemServer start() {
            List<String> command = new ArrayList<>();
            try {
                command.add(resolveBinary(binary).toString());
            } catch (IOException e) {
                throw new VkmemException("vkmem: cannot prepare binary: " + e.getMessage(), e);
            }
            command.add("--parent-pid");
            command.add(Long.toString(ProcessHandle.current().pid()));
            if (port != 0) {
                command.add("--port");
                command.add(Integer.toString(port));
            }
            if (noUnixSocket) {
                command.add("--no-unixsocket");
            } else if (unixSocket != null) {
                command.add("--unixsocket");
                command.add(unixSocket);
            }
            if (quiet) command.add("--quiet");
            if (!args.isEmpty()) {
                command.add("--");
                command.addAll(args);
            }
            Process process;
            try {
                process = new ProcessBuilder(command).redirectError(ProcessBuilder.Redirect.INHERIT).start();
            } catch (IOException e) {
                throw new VkmemException("vkmem: failed to start " + command.get(0) + ": " + e.getMessage(), e);
            }
            BufferedReader stdout = new BufferedReader(new InputStreamReader(process.getInputStream(), StandardCharsets.UTF_8));
            Map<String, Object> ready = readReady(process, stdout, startupTimeout);
            return new VkmemServer(new Control(process, stdout), ready, true);
        }

        private static Map<String, Object> readReady(Process process, BufferedReader stdout, Duration timeout) {
            CompletableFuture<String> first = CompletableFuture.supplyAsync(() -> {
                try {
                    return stdout.readLine();
                } catch (IOException e) {
                    return null;
                }
            });
            String line;
            try {
                line = first.get(timeout.toMillis(), TimeUnit.MILLISECONDS);
            } catch (TimeoutException e) {
                process.destroyForcibly();
                throw new VkmemException("vkmem: server did not start within " + timeout, e);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                process.destroyForcibly();
                throw new VkmemException("vkmem: interrupted while starting", e);
            } catch (ExecutionException e) {
                process.destroyForcibly();
                throw new VkmemException("vkmem: cannot read readiness", e.getCause());
            }
            if (line == null) {
                int code;
                try {
                    code = process.waitFor();
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                    code = -1;
                }
                throw new VkmemException("vkmem: server exited with status " + code + " before becoming ready");
            }
            Map<String, Object> ready;
            try {
                ready = Json.parseObject(line);
            } catch (IllegalArgumentException e) {
                process.destroyForcibly();
                throw new VkmemException("vkmem: malformed ready line: " + line, e);
            }
            if (!"ready".equals(ready.get("event")) || !(ready.get("port") instanceof Number)) {
                process.destroyForcibly();
                throw new VkmemException("vkmem: unexpected ready line: " + line);
            }
            Object protocol = ready.get("protocol");
            if (!(protocol instanceof Number) || ((Number) protocol).intValue() != Control.PROTOCOL) {
                process.destroyForcibly();
                throw new VkmemException("vkmem: server speaks protocol " + protocol
                        + ", this client needs " + Control.PROTOCOL);
            }
            return ready;
        }
    }
}
