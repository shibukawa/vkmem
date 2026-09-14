# vkmem for Java

Launcher and JUnit 5 extension for `vkmem-server`, a real Valkey server
started as a child process for tests.

```java
@RegisterExtension
static VkmemExtension vkmem = VkmemExtension.create(); // FLUSHALL before each test

@Test
void counter(VkmemServer server) {
    try (Jedis jedis = new Jedis(server.host(), server.port())) { // or Lettuce, valkey-java
        jedis.incr("visits");
    }
}
```

Dependencies (Maven):

```xml
<dependency>
  <groupId>io.github.shibukawa.vkmem</groupId>
  <artifactId>vkmem</artifactId>
  <version>0.1.0</version>
  <scope>test</scope>
</dependency>
<dependency>
  <groupId>io.github.shibukawa.vkmem</groupId>
  <artifactId>vkmem-server-binaries</artifactId>
  <version>0.1.0</version>
  <classifier>linux-amd64</classifier> <!-- darwin-arm64, linux-arm64, windows-amd64, windows-arm64 -->
  <scope>test</scope>
</dependency>
```

The binary jar contains `vkmem/bin/<os>-<arch>/vkmem-server`; it is
extracted to a temp directory on first use. `-Dvkmem.server.bin=...` or
`VKMEM_SERVER_BIN` point at a locally built binary instead.

`VkmemServer.builder()` takes `port`, `unixSocket`/`noUnixSocket`, `args`
(extra `valkey-server` flags) and `binary`. `server.command(...)` is a tiny
RESP client for test plumbing (PING, FLUSHALL, CONFIG); use a real client
for application code.

Layout: `vkmem/` (launcher + extension), `binaries/` (pom-only module
that attaches the classifier jars produced by
`scripts/build-java-binaries.sh` during `mvn -Prelease deploy`). The
launcher has no dependencies beyond the JDK (17+); the JUnit extension
needs junit-jupiter-api on the classpath.
