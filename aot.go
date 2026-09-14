package valkeymem

import (
	"github.com/shibukawa/valkeymem/internal/aot"
	"github.com/shibukawa/valkeymem/internal/engine"
)

// The backend is Go code generated ahead of time by wasm2go from the
// Valkey wasm module; no wasm runtime is involved at run time.
func init() {
	engine.Factory = aot.Factory{}
}
