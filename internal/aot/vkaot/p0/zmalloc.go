package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zmalloc_get_memory_size(m *base.Module) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = F_sysconf(m, int32(85))
	v4 = F_sysconf(m, int32(30))
	return v2 * v4
}
func F_zmalloc_get_smap_bytes_by_field(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F_zmalloc_set_oom_handler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_zmalloc_set_oom_handler[0])) = l0
	return
}
