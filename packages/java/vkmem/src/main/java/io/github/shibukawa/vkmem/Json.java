package io.github.shibukawa.vkmem;

import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Minimal JSON codec for the flat vkmem control-protocol messages. */
final class Json {
    private Json() {}

    static String write(Object value) {
        StringBuilder out = new StringBuilder();
        write(out, value);
        return out.toString();
    }

    @SuppressWarnings("unchecked")
    private static void write(StringBuilder out, Object value) {
        if (value == null) {
            out.append("null");
        } else if (value instanceof String) {
            writeString(out, (String) value);
        } else if (value instanceof Boolean || value instanceof Number) {
            out.append(value);
        } else if (value instanceof Map) {
            out.append('{');
            boolean first = true;
            for (Map.Entry<String, Object> entry : ((Map<String, Object>) value).entrySet()) {
                if (!first) out.append(',');
                first = false;
                writeString(out, entry.getKey());
                out.append(':');
                write(out, entry.getValue());
            }
            out.append('}');
        } else if (value instanceof Iterable) {
            out.append('[');
            boolean first = true;
            for (Object item : (Iterable<?>) value) {
                if (!first) out.append(',');
                first = false;
                write(out, item);
            }
            out.append(']');
        } else {
            throw new IllegalArgumentException("cannot encode " + value.getClass());
        }
    }

    private static void writeString(StringBuilder out, String value) {
        out.append('"');
        for (int i = 0; i < value.length(); i++) {
            char c = value.charAt(i);
            switch (c) {
                case '"': out.append("\\\""); break;
                case '\\': out.append("\\\\"); break;
                case '\n': out.append("\\n"); break;
                case '\r': out.append("\\r"); break;
                case '\t': out.append("\\t"); break;
                case '\b': out.append("\\b"); break;
                case '\f': out.append("\\f"); break;
                default:
                    if (c < 0x20) out.append(String.format("\\u%04x", (int) c));
                    else out.append(c);
            }
        }
        out.append('"');
    }

    static Object parse(String text) {
        Parser parser = new Parser(text);
        Object value = parser.value();
        parser.skipWhitespace();
        if (parser.index != text.length()) throw parser.error("trailing data");
        return value;
    }

    @SuppressWarnings("unchecked")
    static Map<String, Object> parseObject(String text) {
        Object value = parse(text);
        if (!(value instanceof Map)) throw new IllegalArgumentException("not a JSON object: " + text);
        return (Map<String, Object>) value;
    }

    private static final class Parser {
        private final String text;
        private int index;

        Parser(String text) { this.text = text; }

        IllegalArgumentException error(String message) {
            return new IllegalArgumentException("invalid JSON at " + index + ": " + message);
        }

        void skipWhitespace() {
            while (index < text.length() && Character.isWhitespace(text.charAt(index))) index++;
        }

        Object value() {
            skipWhitespace();
            if (index >= text.length()) throw error("unexpected end");
            switch (text.charAt(index)) {
                case '{': return object();
                case '[': return array();
                case '"': return string();
                case 't': expect("true"); return Boolean.TRUE;
                case 'f': expect("false"); return Boolean.FALSE;
                case 'n': expect("null"); return null;
                default: return number();
            }
        }

        void expect(String literal) {
            if (!text.startsWith(literal, index)) throw error("expected " + literal);
            index += literal.length();
        }

        Map<String, Object> object() {
            Map<String, Object> result = new LinkedHashMap<>();
            index++;
            skipWhitespace();
            if (index < text.length() && text.charAt(index) == '}') {
                index++;
                return result;
            }
            while (true) {
                skipWhitespace();
                String key = string();
                skipWhitespace();
                if (index >= text.length() || text.charAt(index) != ':') throw error("expected :");
                index++;
                result.put(key, value());
                skipWhitespace();
                if (index >= text.length()) throw error("expected , or }");
                char separator = text.charAt(index++);
                if (separator == '}') return result;
                if (separator != ',') throw error("expected , or }");
            }
        }

        List<Object> array() {
            List<Object> result = new ArrayList<>();
            index++;
            skipWhitespace();
            if (index < text.length() && text.charAt(index) == ']') {
                index++;
                return result;
            }
            while (true) {
                result.add(value());
                skipWhitespace();
                if (index >= text.length()) throw error("expected , or ]");
                char separator = text.charAt(index++);
                if (separator == ']') return result;
                if (separator != ',') throw error("expected , or ]");
            }
        }

        String string() {
            if (index >= text.length() || text.charAt(index) != '"') throw error("expected string");
            index++;
            StringBuilder result = new StringBuilder();
            while (index < text.length()) {
                char c = text.charAt(index++);
                if (c == '"') return result.toString();
                if (c != '\\') {
                    result.append(c);
                    continue;
                }
                if (index >= text.length()) throw error("bad escape");
                char escape = text.charAt(index++);
                switch (escape) {
                    case '"': result.append('"'); break;
                    case '\\': result.append('\\'); break;
                    case '/': result.append('/'); break;
                    case 'b': result.append('\b'); break;
                    case 'f': result.append('\f'); break;
                    case 'n': result.append('\n'); break;
                    case 'r': result.append('\r'); break;
                    case 't': result.append('\t'); break;
                    case 'u':
                        if (index + 4 > text.length()) throw error("short unicode escape");
                        result.append((char) Integer.parseInt(text.substring(index, index + 4), 16));
                        index += 4;
                        break;
                    default: throw error("bad escape");
                }
            }
            throw error("unterminated string");
        }

        Object number() {
            int start = index;
            while (index < text.length() && "+-0123456789.eE".indexOf(text.charAt(index)) >= 0) index++;
            String value = text.substring(start, index);
            if (value.isEmpty()) throw error("unexpected character");
            if (value.indexOf('.') < 0 && value.indexOf('e') < 0 && value.indexOf('E') < 0) {
                return Long.parseLong(value);
            }
            return Double.parseDouble(value);
        }
    }
}
