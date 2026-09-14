package vkaot

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	_ "unsafe"
)
//go:linkname F___wasm_call_ctors github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___wasm_call_ctors
func F___wasm_call_ctors(m *base.Module)
//go:linkname F_vkmem_main github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_vkmem_main
func F_vkmem_main(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_vkmem_dlopen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_vkmem_dlopen
func F_vkmem_dlopen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_vkmem_dlsym github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_vkmem_dlsym
func F_vkmem_dlsym(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_vkmem_dlclose github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_vkmem_dlclose
func F_vkmem_dlclose(m *base.Module, l0 int32) int32
//go:linkname F_vkmem_dlerror github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_vkmem_dlerror
func F_vkmem_dlerror(m *base.Module) int32
//go:linkname F_vkmem_call_sighandler github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_vkmem_call_sighandler
func F_vkmem_call_sighandler(m *base.Module, l0 int32, l1 int32)
//go:linkname F_main github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_main
func F_main(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___errno_location github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___errno_location
func F___errno_location(m *base.Module) int32
//go:linkname F_fflush github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fflush
func F_fflush(m *base.Module, l0 int32) int32
//go:linkname F_htonl github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_htonl
func F_htonl(m *base.Module, l0 int32) int32
//go:linkname F_htons github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_htons
func F_htons(m *base.Module, l0 int32) int32
//go:linkname F_ntohs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ntohs
func F_ntohs(m *base.Module, l0 int32) int32
//go:linkname F__emscripten_timeout github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__emscripten_timeout
func F__emscripten_timeout(m *base.Module, l0 int32, l1 float64)
//go:linkname F_strerror github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strerror
func F_strerror(m *base.Module, l0 int32) int32
//go:linkname F_emscripten_builtin_malloc github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_emscripten_builtin_malloc
func F_emscripten_builtin_malloc(m *base.Module, l0 int32) int32
//go:linkname F_emscripten_builtin_free github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_emscripten_builtin_free
func F_emscripten_builtin_free(m *base.Module, l0 int32)
//go:linkname F_emscripten_builtin_memalign github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_emscripten_builtin_memalign
func F_emscripten_builtin_memalign(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_emscripten_stack_init github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_emscripten_stack_init
func F_emscripten_stack_init(m *base.Module)
//go:linkname F_emscripten_stack_get_free github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_emscripten_stack_get_free
func F_emscripten_stack_get_free(m *base.Module) int32
//go:linkname F_emscripten_stack_get_base github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_emscripten_stack_get_base
func F_emscripten_stack_get_base(m *base.Module) int32
//go:linkname F_emscripten_stack_get_end github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_emscripten_stack_get_end
func F_emscripten_stack_get_end(m *base.Module) int32
//go:linkname F___funcs_on_exit github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___funcs_on_exit
func F___funcs_on_exit(m *base.Module)
//go:linkname F__emscripten_stack_restore github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__emscripten_stack_restore
func F__emscripten_stack_restore(m *base.Module, l0 int32)
//go:linkname F__emscripten_stack_alloc github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__emscripten_stack_alloc
func F__emscripten_stack_alloc(m *base.Module, l0 int32) int32
//go:linkname F_emscripten_stack_get_current github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_emscripten_stack_get_current
func F_emscripten_stack_get_current(m *base.Module) int32
//go:linkname InitElemSeg_0_0 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.InitElemSeg_0_0
func InitElemSeg_0_0(m *base.Module)
//go:linkname InitElemSeg_0_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.InitElemSeg_0_1
func InitElemSeg_0_1(m *base.Module)
//go:linkname InitElemSeg_1_0 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.InitElemSeg_1_0
func InitElemSeg_1_0(m *base.Module)
//go:linkname InitElemSeg_1_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.InitElemSeg_1_1
func InitElemSeg_1_1(m *base.Module)
//go:linkname InitElemSeg_2_0 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.InitElemSeg_2_0
func InitElemSeg_2_0(m *base.Module)
//go:linkname InitElemSeg_2_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.InitElemSeg_2_1
func InitElemSeg_2_1(m *base.Module)
