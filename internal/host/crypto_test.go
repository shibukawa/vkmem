package host

import "testing"

func TestCRC64Jones(t *testing.T) {
	// The check value from Valkey's crc64.c self-test.
	if got := crc64Raw(0, []byte("123456789")); got != 0xe9c6d914c4b8d9ca {
		t.Fatalf("crc64 = %#x", got)
	}
	// Incremental updates compose like the C version.
	a := crc64Raw(0, []byte("12345"))
	if got := crc64Raw(a, []byte("6789")); got != 0xe9c6d914c4b8d9ca {
		t.Fatalf("incremental crc64 = %#x", got)
	}
}
