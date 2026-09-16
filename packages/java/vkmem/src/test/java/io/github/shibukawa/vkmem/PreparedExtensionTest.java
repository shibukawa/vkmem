package io.github.shibukawa.vkmem;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.RegisterExtension;

class PreparedExtensionTest {
    @RegisterExtension
    static VkmemExtension vkmem = VkmemExtension.prepared(
            template -> template.command("SET", "prepared", "yes"), 1, null);

    @Test
    void eachTestGetsPreparedFork(VkmemServer server) {
        assertEquals("yes", server.command("GET", "prepared"));
        assertEquals("OK", server.command("SET", "test-only", "yes"));
    }

    @Test
    void nextTestGetsAFreshFork(VkmemServer server) {
        assertEquals("yes", server.command("GET", "prepared"));
        assertNull(server.command("GET", "test-only"));
    }
}
