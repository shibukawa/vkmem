package io.github.shibukawa.vkmem;

/** The vkmem-server process ended while a control request was pending. */
public final class ServerExitedException extends VkmemException {
    private static final long serialVersionUID = 1L;

    public ServerExitedException(String message) { super(message); }
    public ServerExitedException(String message, Throwable cause) { super(message, cause); }
}
