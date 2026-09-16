package io.github.shibukawa.vkmem;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

import java.time.Duration;
import java.util.List;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.RegisterExtension;

class VkmemServerTest {
    @RegisterExtension
    static VkmemExtension vkmem = VkmemExtension.builder(b -> b.args("--maxmemory", "32mb"));

    @Test
    void commandsAndConfig(VkmemServer server) {
        assertTrue(server.port() > 0);
        assertEquals("redis://127.0.0.1:" + server.port(), server.url());
        assertTrue(server.valkeyVersion().startsWith("9."));
        assertEquals("PONG", server.command("PING"));
        assertEquals("OK", server.command("SET", "k", "v"));
        assertEquals("v", server.command("GET", "k"));
        assertEquals(1L, server.command("INCR", "n"));
        assertEquals(List.of("maxmemory", "33554432"), server.command("CONFIG", "GET", "maxmemory"));
        assertThrows(VkmemException.class, () -> server.command("NOSUCHCOMMAND"));
    }

    @Test
    void eachTestStartsEmpty(VkmemServer server) {
        assertNull(server.command("GET", "k"));
        assertEquals(0L, server.command("DBSIZE"));
    }

    @Test
    void snapshotForkStartsPreparedAndStaysIsolated(VkmemServer template) {
        assertEquals("OK", template.command("SET", "prepared", "yes"));
        try (Snapshot snapshot = template.snapshot(2)) {
            assertEquals("OK", template.command("SET", "template-only", "yes"));
            try (Fork first = snapshot.fork(); Fork second = snapshot.fork()) {
                assertEquals("yes", first.command("GET", "prepared"));
                assertEquals("OK", first.command("SET", "fork-only", "yes"));
                assertNull(second.command("GET", "fork-only"));
                assertNull(second.command("GET", "template-only"));
                assertNull(template.command("GET", "fork-only"));
            }
        }
    }

    @Test
    void snapshotForkLimitTimesOutAndReleases() {
        VkmemServer template = VkmemServer.builder().noUnixSocket().start();
        try (Snapshot snapshot = template.snapshot(1)) {
            Fork first = snapshot.fork();
            try {
                ProtocolException timeout = assertThrows(ProtocolException.class,
                        () -> snapshot.fork(Duration.ofMillis(100)));
                assertEquals("pool_timeout", timeout.code());
            } finally {
                first.close();
            }
            try (Fork second = snapshot.fork(Duration.ofSeconds(1))) {
                assertEquals("PONG", second.command("PING"));
            }
        } finally {
            template.close();
        }
    }

    @Test
    void closeStopsTheProcess() {
        VkmemServer s = VkmemServer.builder().noUnixSocket().start();
        int port = s.port();
        assertEquals("PONG", s.command("PING"));
        s.close();
        assertThrows(java.io.IOException.class, () -> {
            try (java.net.Socket sock = new java.net.Socket()) {
                sock.connect(new java.net.InetSocketAddress("127.0.0.1", port), 1000);
            }
        });
    }
}
