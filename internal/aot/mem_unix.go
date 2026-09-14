//go:build unix

package aot

import "syscall"

// reserveMemory maps n bytes of anonymous memory outside the Go heap. Pages
// are materialized by the OS on first touch, so a large reservation costs
// nothing until the guest uses it, and release returns everything at once.
func reserveMemory(n int) (mem []byte, release func(), err error) {
	mem, err = syscall.Mmap(-1, 0, n, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		return nil, nil, err
	}
	return mem, func() { syscall.Munmap(mem) }, nil
}

// commitMemory is a no-op on unix: anonymous mappings commit lazily.
func commitMemory(mem []byte, from, to uint64) error { return nil }
