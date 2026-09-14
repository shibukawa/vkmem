package valkeymem

import (
	"context"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"testing"
)

// TestHashVectors pins the outputs that depend on sha1/sha256/crc64/lzf so
// a host-side replacement of those routines is checked byte for byte.
func TestHashVectors(t *testing.T) {
	s := startTestServer(t)
	c := newClient(t, s)
	ctx := context.Background()
	// SCRIPT LOAD returns the sha1 of the body.
	body := "return {KEYS[1], ARGV[1]}"
	got, err := c.Do(ctx, c.B().ScriptLoad().Script(body).Build()).ToString()
	if err != nil {
		t.Fatal(err)
	}
	sum := sha1.Sum([]byte(body))
	if want := hex.EncodeToString(sum[:]); got != want {
		t.Fatalf("SCRIPT LOAD sha1 = %s, want %s", got, want)
	}
	// ACL stores sha256 of the password.
	if err := c.Do(ctx, c.B().Arbitrary("ACL", "SETUSER", "vec", "on", ">hunter2").Build()).Error(); err != nil {
		t.Fatal(err)
	}
	acl, err := c.Do(ctx, c.B().Arbitrary("ACL", "LIST").Build()).AsStrSlice()
	if err != nil {
		t.Fatal(err)
	}
	pw := sha256.Sum256([]byte("hunter2"))
	if want := "#" + hex.EncodeToString(pw[:]); !strings.Contains(strings.Join(acl, "\n"), want) {
		t.Fatalf("ACL LIST %v lacks %s", acl, want)
	}
	// DUMP payload: RDB type+value, then version and crc64; a 60-byte string
	// is lzf-compressed. These bytes came from the unmodified C build.
	c.Do(ctx, c.B().Set().Key("vec").Value(strings.Repeat("abcabcabc", 7)).Build())
	dump, err := c.Do(ctx, c.B().Dump().Key("vec").Build()).ToString()
	if err != nil {
		t.Fatal(err)
	}
	const wantDump = "00c30b3f0361626361e030020162635000cd08be753b142dce"
	if got := hex.EncodeToString([]byte(dump)); got != wantDump {
		t.Fatalf("DUMP = %s, want %s", got, wantDump)
	}
	if err := c.Do(ctx, c.B().Restore().Key("vec2").Ttl(0).SerializedValue(dump).Build()).Error(); err != nil {
		t.Fatal(err)
	}
	if v, _ := c.Do(ctx, c.B().Get().Key("vec2").Build()).ToString(); v != strings.Repeat("abcabcabc", 7) {
		t.Fatalf("RESTORE round trip: %q", v)
	}
}
