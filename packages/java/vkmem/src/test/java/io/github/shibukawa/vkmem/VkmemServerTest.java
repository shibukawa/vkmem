package io.github.shibukawa.vkmem;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

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
