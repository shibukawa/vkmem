package io.github.shibukawa.vkmem;

import java.util.function.Consumer;

import org.junit.jupiter.api.extension.AfterAllCallback;
import org.junit.jupiter.api.extension.BeforeAllCallback;
import org.junit.jupiter.api.extension.BeforeEachCallback;
import org.junit.jupiter.api.extension.ExtensionContext;
import org.junit.jupiter.api.extension.ParameterContext;
import org.junit.jupiter.api.extension.ParameterResolver;

/**
 * JUnit 5 extension: one server per test class, {@code FLUSHALL} before each test method.
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
        implements BeforeAllCallback, AfterAllCallback, BeforeEachCallback, ParameterResolver {
    private final Consumer<VkmemServer.Builder> configure;
    private final boolean flushEach;
    private VkmemServer server;

    private VkmemExtension(Consumer<VkmemServer.Builder> configure, boolean flushEach) {
        this.configure = configure;
        this.flushEach = flushEach;
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

    public VkmemServer server() {
        return server;
    }

    @Override
    public void beforeAll(ExtensionContext context) {
        VkmemServer.Builder b = VkmemServer.builder();
        configure.accept(b);
        server = b.start();
    }

    @Override
    public void afterAll(ExtensionContext context) {
        if (server != null) {
            server.close();
            server = null;
        }
    }

    @Override
    public void beforeEach(ExtensionContext context) {
        if (flushEach) {
            server.flushAll();
        }
    }

    @Override
    public boolean supportsParameter(ParameterContext parameterContext, ExtensionContext extensionContext) {
        return parameterContext.getParameter().getType() == VkmemServer.class;
    }

    @Override
    public Object resolveParameter(ParameterContext parameterContext, ExtensionContext extensionContext) {
        return server;
    }
}
