package io.github.shibukawa.vkmem;

/** Thrown when vkmem-server cannot be started or a command fails. */
public class VkmemException extends RuntimeException {
    private static final long serialVersionUID = 1L;

    public VkmemException(String message) {
        super(message);
    }

    public VkmemException(String message, Throwable cause) {
        super(message, cause);
    }
}
