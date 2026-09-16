package io.github.shibukawa.vkmem;

import java.time.Duration;
import java.util.LinkedHashMap;
import java.util.Map;

/** A frozen copy of a server's serialized Valkey keyspace. */
public final class Snapshot implements AutoCloseable {
    private final Control control;
    private final String id;
    private final VkmemServer origin;
    private volatile boolean closed;

    Snapshot(Control control, String id, VkmemServer origin) {
        this.control = control;
        this.id = id;
        this.origin = origin;
    }

    public String id() {
        return id;
    }

    /** The server this snapshot was taken from. */
    public VkmemServer origin() {
        return origin;
    }

    /** Starts a fresh server on a private copy of this snapshot. */
    public Fork fork() {
        return fork(null);
    }

    /**
     * Like {@link #fork()}, but bounds waiting for a free {@code maxForks} slot.
     * A {@code null} timeout waits indefinitely.
     */
    public Fork fork(Duration timeout) {
        Map<String, Object> fields = new LinkedHashMap<>();
        fields.put("snapshot", id);
        if (timeout != null) fields.put("timeout_ms", timeout.toMillis());
        Map<String, Object> result = control.request("fork", fields);
        @SuppressWarnings("unchecked")
        Map<String, Object> endpoint = (Map<String, Object>) result.get("server");
        return new Fork(control, endpoint);
    }

    /** Rejects new forks; live forks keep running. Idempotent. */
    @Override
    public void close() {
        if (closed) return;
        closed = true;
        try {
            control.request("close", Map.of("snapshot", id));
        } catch (ServerExitedException ignored) {
            // The controller may already have exited.
        }
    }

    @Override
    public String toString() {
        return "Snapshot(" + id + " of " + origin.id() + ")";
    }
}
