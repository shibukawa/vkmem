// Package guest defines the engine's view of an instantiated module (Go
// code generated ahead of time by wasm2go).
package guest

import (
	"context"

	"github.com/shibukawa/valkeymem/internal/host"
)

// Instance is one running copy of a module bound to one host.Host.
type Instance interface {
	// Call invokes an export by its wasm name. Guest exits surface as
	// *host.ExitError, aborts as *host.AbortError, traps as other errors.
	// The instance stays usable after an unwinding call.
	Call(name string, args ...uint64) ([]uint64, error)
	// CallI32 is Call for exports returning a single i32.
	CallI32(name string, args ...uint64) (int32, error)
	// Memory is the instance's linear memory.
	Memory() host.Memory
	// Malloc/Free use the guest allocator.
	Malloc(n uint32) (uint32, error)
	Free(p uint32)
	// WriteCString copies s NUL-terminated into freshly allocated memory.
	WriteCString(s string) (uint32, error)
	// Close releases the instance.
	Close(ctx context.Context) error
}

// Factory creates instances of one module.
type Factory interface {
	Instantiate(ctx context.Context, h *host.Host) (Instance, error)
}
