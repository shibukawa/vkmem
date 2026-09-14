package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___pthread_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	return int32(6)
}
func F__pthread_cleanup_push(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	return
}
func F_pthread_attr_getstacksize(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3
	return int32(0)
}
func F_pthread_attr_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v12 = F__emscripten_memset_bulkmem(m, v5+int32(4), base.I32_extend8_s(int32(0)), int32(44))
	mBase = m.M
	v18 = F__emscripten_memcpy_bulkmem(m, l0, v5+int32(4), int32(44))
	mBase = m.M
	v20 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, _consts[1075]))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v21
	v24 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v24
	m.G0 = v5 + int32(48)
	return int32(0)
}
func F_pthread_mutexattr_settype(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F_pthread_setcanceltype(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
