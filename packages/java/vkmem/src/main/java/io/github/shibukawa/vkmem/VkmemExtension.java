package io.github.shibukawa.vkmem;

import java.time.Duration;
import java.util.HashMap;
import java.util.Map;
import java.util.function.Consumer;

import org.junit.jupiter.api.extension.AfterAllCallback;
import org.junit.jupiter.api.extension.AfterEachCallback;
import org.junit.jupiter.api.extension.BeforeAllCallback;
import org.junit.jupiter.api.extension.BeforeEachCallback;
import org.junit.jupiter.api.extension.ExtensionContext;
import org.junit.jupiter.api.extension.ParameterContext;
import org.junit.jupiter.api.extension.ParameterResolver;

/**
 * JUnit 5 extension: one server per test class, {@code FLUSHALL} before each
 * test method by default. {@link #prepared(Consumer)} provides a prepared
 * template and a private {@link Fork} per test method.
 *
 * <pre>
 * &#64;RegisterExtension
 * static VkmemExtension vkmem = VkmemExtension.create();
 *
 * &#64;Test
 * void counter(VkmemServer server) {
 *     Jedis jedis = new Jedis(server.host(), server.port());
 * }
 * </pre>
 *
 * Test methods may declare a {@link VkmemServer} parameter; {@link #server()} returns the same object.
 */
public final class VkmemExtension
        implements BeforeAllCallback, AfterAllCallback, BeforeEachCallback, AfterEachCallback, ParameterResolver {
    private final Consumer<VkmemServer.Builder> configure;
    private final boolean flushEach;
    private final Consumer<VkmemServer> prepare;
    private final int maxForks;
    private final Duration forkTimeout;
    private final Map<String, Fork> forks = new HashMap<>();
    private VkmemServer server;
    private Snapshot snapshot;

    private VkmemExtension(Consumer<VkmemServer.Builder> configure, boolean flushEach) {
        this(configure, flushEach, null, 0, null);
    }

    private VkmemExtension(Consumer<VkmemServer.Builder> configure, boolean flushEach,
                           Consumer<VkmemServer> prepare, int maxForks, Duration forkTimeout) {
        this.configure = configure;
        this.flushEach = flushEach;
        this.prepare = prepare;
        this.maxForks = maxForks;
        this.forkTimeout = forkTimeout;
    }

    /** Default server, FLUSHALL before each test. */
    public static VkmemExtension create() {
        return new VkmemExtension(b -> { }, true);
    }

    /** Configured server, FLUSHALL before each test. */
    public static VkmemExtension builder(Consumer<VkmemServer.Builder> configure) {
        return new VkmemExtension(configure, true);
    }

    /** Configured server; the keyspace is left alone between tests. */
    public static VkmemExtension builder(Consumer<VkmemServer.Builder> configure, boolean flushBeforeEach) {
        return new VkmemExtension(configure, flushBeforeEach);
    }

    /** Starts a template, prepares it once, and injects a fresh Fork per test method. */
    public static VkmemExtension prepared(Consumer<VkmemServer> prepare) {
        return prepared(prepare, 0, null);
    }

    /** Prepared extension with an optional fork limit and slot timeout. */
    public static VkmemExtension prepared(Consumer<VkmemServer> prepare, int maxForks, Duration forkTimeout) {
        return new VkmemExtension(b -> { }, false, prepare, maxForks, forkTimeout);
    }

    public VkmemServer server() {
        return server;
    }

    /** The prepared data snapshot, available after {@code beforeAll}. */
    public Snapshot snapshot() {
        if (snapshot == null) throw new IllegalStateException("vkmem snapshot not started");
        return snapshot;
    }

    @Override
    public void beforeAll(ExtensionContext context) {
        VkmemServer.Builder b = VkmemServer.builder();
        configure.accept(b);
        server = b.start();
        if (prepare != null) {
            try {
                prepare.accept(server);
                snapshot = server.snapshot(maxForks);
            } catch (RuntimeException e) {
                server.close();
                server = null;
                throw e;
            }
        }
    }

    @Override
    public void afterAll(ExtensionContext context) {
        if (server != null) {
            closeForks();
            if (snapshot != null) {
                snapshot.close();
                snapshot = null;
            }
            server.close();
            server = null;
        }
    }

    @Override
    public void beforeEach(ExtensionContext context) {
        if (snapshot == null && flushEach) {
            server.flushAll();
        }
    }

    @Override
    public void afterEach(ExtensionContext context) {
        Fork fork;
        synchronized (forks) {
            fork = forks.remove(context.getUniqueId());
        }
        if (fork != null) fork.close();
    }

    @Override
    public boolean supportsParameter(ParameterContext parameterContext, ExtensionContext extensionContext) {
        Class<?> type = parameterContext.getParameter().getType();
        return type == VkmemServer.class || (snapshot != null && type == Fork.class);
    }

    @Override
    public Object resolveParameter(ParameterContext parameterContext, ExtensionContext extensionContext) {
        if (snapshot != null) {
            synchronized (forks) {
                return forks.computeIfAbsent(extensionContext.getUniqueId(), key -> snapshot.fork(forkTimeout));
            }
        }
        return server;
    }

    private void closeForks() {
        synchronized (forks) {
            for (Fork fork : forks.values()) fork.close();
            forks.clear();
        }
    }
}
