package io.github.shibukawa.vkmem;

import java.util.Map;

/** A server started from a {@link Snapshot}; closing it releases its fork slot. */
public final class Fork extends VkmemServer {
    Fork(Control control, Map<String, Object> endpoint) {
        super(control, endpoint, false);
    }
}
