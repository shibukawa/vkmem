//go:build !unix && !windows

package aot

// reserveMemory falls back to the Go heap where no lazy mapping is
// available.
func reserveMemory(n int) (mem []byte, release func(), err error) {
	return make([]byte, n), func() {}, nil
}

func commitMemory(mem []byte, from, to uint64) error { return nil }
