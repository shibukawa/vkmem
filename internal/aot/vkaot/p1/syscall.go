package p1

import base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"

func F___syscall_getuid32(m *base.Module) int32 {
	return int32(0)
}
func F___syscall_shutdown(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	m.Env.Emscripten_err(m, int32(_a2764))
	return int32(-52)
}
func F___syscall_wait4(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	m.Env.Emscripten_err(m, int32(_a2765))
	return int32(-52)
}
