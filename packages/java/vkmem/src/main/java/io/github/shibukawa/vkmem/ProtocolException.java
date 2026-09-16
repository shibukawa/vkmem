package io.github.shibukawa.vkmem;

/** The server rejected a JSON-lines control request. */
public final class ProtocolException extends VkmemException {
    private static final long serialVersionUID = 1L;
    private final String code;

    public ProtocolException(String code, String message) {
        super(code + ": " + message);
        this.code = code;
    }

    /** Error code such as {@code pool_timeout}, {@code snapshot_closed} or {@code internal}. */
    public String code() { return code; }
}
