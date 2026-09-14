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
import java.util.List;
import java.util.Locale;
import java.util.concurrent.TimeUnit;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * A running vkmem-server process: a real Valkey server for tests.
 *
 * <pre>{@code
 * try (VkmemServer server = VkmemServer.builder().start()) {
 *     // point Jedis, Lettuce or valkey-java at server.host()/server.port()
 *     server.command("FLUSHALL");
 * }
 * }</pre>
 *
 * The binary is located from the {@code vkmem.server.bin} system property, the {@code VKMEM_SERVER_BIN}
 * environment variable, or the classpath resource {@code /vkmem/bin/<os>-<arch>/vkmem-server} shipped in the
 * {@code vkmem-server-binaries} artifact with the matching classifier.
 */
public final class VkmemServer implements AutoCloseable {
    private static final Pattern PORT_RE = Pattern.compile("\"port\"\\s*:\\s*(\\d+)");
    private static final Pattern UNIX_RE = Pattern.compile("\"unix\"\\s*:\\s*\"([^\"]+)\"");
    private static final Pattern VALKEY_RE = Pattern.compile("\"valkey\"\\s*:\\s*\"([^\"]+)\"");

    private final Process process;
    private final int port;
    private final String unixSocket;
    private final String valkeyVersion;

    private VkmemServer(Process process, int port, String unixSocket, String valkeyVersion) {
        this.process = process;
        this.port = port;
        this.unixSocket = unixSocket;
        this.valkeyVersion = valkeyVersion;
    }

    public static Builder builder() {
        return new Builder();
    }

    /** Always {@code 127.0.0.1}. */
    public String host() {
        return "127.0.0.1";
    }

    public int port() {
        return port;
    }

    /** {@code host:port}. */
    public String address() {
        return host() + ":" + port;
    }

    /** {@code redis://127.0.0.1:port}, accepted by most clients. */
    public String url() {
        return "redis://" + address();
    }

    /** Unix socket path, or {@code null} when disabled. */
    public String unixSocket() {
        return unixSocket;
    }

    /** The Valkey release compiled into the server. */
    public String valkeyVersion() {
        return valkeyVersion;
    }

    /**
     * Sends one command over a fresh TCP connection and returns the decoded reply: a {@link String} for
     * simple and bulk strings, a {@link Long} for integers, {@code null} for nil, a {@link List} for arrays.
     * Error replies throw {@link VkmemException}. Meant for test plumbing (PING, FLUSHALL, CONFIG); use a
     * real client for application code.
     */
    public Object command(String... args) {
        try (Socket sock = new Socket()) {
            sock.connect(new InetSocketAddress(host(), port), 5000);
            sock.setSoTimeout(30000);
            OutputStream out = sock.getOutputStream();
            StringBuilder sb = new StringBuilder("*").append(args.length).append("\r\n");
            for (String a : args) {
                byte[] b = a.getBytes(StandardCharsets.UTF_8);
                sb.append('$').append(b.length).append("\r\n").append(a).append("\r\n");
            }
            out.write(sb.toString().getBytes(StandardCharsets.UTF_8));
            out.flush();
            return readReply(new BufferedInputStream(sock.getInputStream()));
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
                int n = Integer.parseInt(line);
                if (n < 0) {
                    return null;
                }
                byte[] b = in.readNBytes(n);
                readLine(in);
                return new String(b, StandardCharsets.UTF_8);
            }
            case '*': {
                int n = Integer.parseInt(line);
                if (n < 0) {
                    return null;
                }
                List<Object> out = new ArrayList<>(n);
                for (int i = 0; i < n; i++) {
                    out.add(readReply(in));
                }
                return out;
            }
            default:
                throw new IOException("unexpected reply byte " + type);
        }
    }

    private static String readLine(InputStream in) throws IOException {
        StringBuilder sb = new StringBuilder();
        int c;
        while ((c = in.read()) >= 0) {
            if (c == '\r') {
                in.read(); // '\n'
                return sb.toString();
            }
            sb.append((char) c);
        }
        throw new IOException("connection closed");
    }

    /** Stops the process: closes stdin, then destroys it after a grace period. */
    @Override
    public void close() {
        if (!process.isAlive()) {
            return;
        }
        try {
            process.getOutputStream().close();
        } catch (IOException ignored) {
            // already closed
        }
        try {
            if (!process.waitFor(5, TimeUnit.SECONDS)) {
                process.destroyForcibly();
            }
        } catch (InterruptedException e) {
            process.destroyForcibly();
            Thread.currentThread().interrupt();
        }
    }

    /** Locates the vkmem-server binary, extracting the bundled one when needed. */
    public static Path resolveBinary(Path explicit) throws IOException {
        if (explicit != null) {
            return explicit;
        }
        String prop = System.getProperty("vkmem.server.bin");
        if (prop != null && !prop.isEmpty()) {
            return Path.of(prop);
        }
        String env = System.getenv("VKMEM_SERVER_BIN");
        if (env != null && !env.isEmpty()) {
            return Path.of(env);
        }
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
            Path bin = dir.resolve(name);
            Files.copy(in, bin, StandardCopyOption.REPLACE_EXISTING);
            if (!platform.startsWith("windows")) {
                Files.setPosixFilePermissions(bin, PosixFilePermissions.fromString("rwx------"));
            }
            bin.toFile().deleteOnExit();
            dir.toFile().deleteOnExit();
            return bin;
        }
    }

    static String platformId() {
        String os = System.getProperty("os.name", "").toLowerCase(Locale.ROOT);
        String arch = System.getProperty("os.arch", "").toLowerCase(Locale.ROOT);
        String o = os.contains("mac") || os.contains("darwin") ? "darwin" : os.contains("win") ? "windows" : "linux";
        String a = arch.contains("aarch64") || arch.contains("arm64") ? "arm64" : "amd64";
        return o + "-" + a;
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
            List<String> cmd = new ArrayList<>();
            try {
                cmd.add(resolveBinary(binary).toString());
            } catch (IOException e) {
                throw new VkmemException("vkmem: cannot prepare binary: " + e.getMessage(), e);
            }
            cmd.add("--parent-pid");
            cmd.add(Long.toString(ProcessHandle.current().pid()));
            if (port != 0) {
                cmd.add("--port");
                cmd.add(Integer.toString(port));
            }
            if (noUnixSocket) {
                cmd.add("--no-unixsocket");
            } else if (unixSocket != null) {
                cmd.add("--unixsocket");
                cmd.add(unixSocket);
            }
            if (quiet) {
                cmd.add("--quiet");
            }
            if (!args.isEmpty()) {
                cmd.add("--");
                cmd.addAll(args);
            }
            ProcessBuilder pb = new ProcessBuilder(cmd).redirectError(ProcessBuilder.Redirect.INHERIT);
            Process p;
            try {
                p = pb.start();
            } catch (IOException e) {
                throw new VkmemException("vkmem: failed to start " + cmd.get(0) + ": " + e.getMessage(), e);
            }
            String ready = readReadyLine(p, startupTimeout);
            Matcher m = PORT_RE.matcher(ready);
            if (!m.find()) {
                p.destroyForcibly();
                throw new VkmemException("vkmem: malformed ready line: " + ready);
            }
            Matcher u = UNIX_RE.matcher(ready);
            Matcher v = VALKEY_RE.matcher(ready);
            return new VkmemServer(p, Integer.parseInt(m.group(1)), u.find() ? u.group(1) : null,
                    v.find() ? v.group(1) : "");
        }

        private static String readReadyLine(Process p, Duration timeout) {
            String[] result = new String[1];
            Thread t = new Thread(() -> {
                try (BufferedReader r = new BufferedReader(new InputStreamReader(p.getInputStream(), StandardCharsets.UTF_8))) {
                    String line;
                    while ((line = r.readLine()) != null) {
                        if (line.contains("\"port\"")) {
                            result[0] = line;
                            return;
                        }
                    }
                } catch (IOException ignored) {
                    // process died
                }
            }, "vkmem-ready-reader");
            t.setDaemon(true);
            t.start();
            try {
                t.join(timeout.toMillis());
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
            if (result[0] == null) {
                p.destroyForcibly();
                throw new VkmemException("vkmem: server did not start within " + timeout);
            }
            return result[0];
        }
    }
}
