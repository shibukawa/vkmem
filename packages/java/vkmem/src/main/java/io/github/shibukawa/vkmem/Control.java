package io.github.shibukawa.vkmem;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.OutputStreamWriter;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicLong;

/** One JSON-lines control channel for a vkmem-server process. */
final class Control {
    static final int PROTOCOL = 1;

    private final Process process;
    private final BufferedReader stdout;
    private final BufferedWriter stdin;
    private final AtomicLong sequence = new AtomicLong();
    private final ConcurrentHashMap<Long, CompletableFuture<Map<String, Object>>> pending = new ConcurrentHashMap<>();
    private volatile boolean exited;
    private volatile boolean closed;

    Control(Process process, BufferedReader stdout) {
        this.process = process;
        this.stdout = stdout;
        this.stdin = new BufferedWriter(new OutputStreamWriter(process.getOutputStream(), StandardCharsets.UTF_8));
        Thread reader = new Thread(this::readLoop, "vkmem-reader");
        reader.setDaemon(true);
        reader.start();
    }

    Map<String, Object> request(String op, Map<String, Object> fields) {
        if (exited) throw new ServerExitedException("vkmem-server has exited");
        long id = sequence.incrementAndGet();
        Map<String, Object> message = new LinkedHashMap<>();
        message.put("id", id);
        message.put("op", op);
        message.putAll(fields);
        CompletableFuture<Map<String, Object>> response = new CompletableFuture<>();
        pending.put(id, response);
        synchronized (stdin) {
            try {
                stdin.write(Json.write(message));
                stdin.write('\n');
                stdin.flush();
            } catch (IOException e) {
                pending.remove(id);
                throw new ServerExitedException("cannot write to vkmem-server", e);
            }
        }
        Map<String, Object> result;
        try {
            result = response.get();
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            pending.remove(id);
            throw new VkmemException("interrupted while waiting for " + op, e);
        } catch (ExecutionException e) {
            Throwable cause = e.getCause();
            if (cause instanceof VkmemException) throw (VkmemException) cause;
            throw new VkmemException(op + " failed", cause);
        }
        if (!Boolean.TRUE.equals(result.get("ok"))) {
            @SuppressWarnings("unchecked")
            Map<String, Object> error = (Map<String, Object>) result.get("error");
            String code = error == null ? "internal" : String.valueOf(error.get("code"));
            String messageText = error == null ? "unknown error" : String.valueOf(error.get("message"));
            throw new ProtocolException(code, messageText);
        }
        return result;
    }

    void close() {
        if (closed) return;
        closed = true;
        try {
            if (!exited) request("shutdown", Map.of());
        } catch (VkmemException ignored) {
            // The process may have exited between the check and the request.
        }
        try {
            stdin.close();
        } catch (IOException ignored) {
            // already gone
        }
        try {
            if (!process.waitFor(10, TimeUnit.SECONDS)) {
                process.destroyForcibly();
                process.waitFor();
            }
        } catch (InterruptedException e) {
            process.destroyForcibly();
            Thread.currentThread().interrupt();
        }
    }

    private void readLoop() {
        try {
            String line;
            while ((line = stdout.readLine()) != null) {
                line = line.trim();
                if (line.isEmpty()) continue;
                Map<String, Object> message;
                try {
                    message = Json.parseObject(line);
                } catch (IllegalArgumentException ignored) {
                    continue;
                }
                Object id = message.get("id");
                if (id == null) continue;
                CompletableFuture<Map<String, Object>> response = pending.remove(((Number) id).longValue());
                if (response != null) response.complete(message);
            }
        } catch (IOException ignored) {
            // process gone
        } finally {
            exited = true;
            failPending();
        }
    }

    private void failPending() {
        for (Long id : new ArrayList<>(pending.keySet())) {
            CompletableFuture<Map<String, Object>> response = pending.remove(id);
            if (response != null) {
                response.completeExceptionally(new ServerExitedException("vkmem-server exited before answering"));
            }
        }
    }
}
