package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___pthread_join(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(28)
}
func F___pthread_mutex_unlock(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F___pthread_self_internal(m *base.Module) int32 {
	return int32(9116560)
}
func F_init_pthread_self(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	v1 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_init_pthread_self[0])) = int32(9116516)
	v5 = F___syscall_getpid(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_init_pthread_self[1])) = v5
	return
}
func F_pthread_attr_destroy(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F_pthread_attr_setstacksize(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	if base.Ui32(l1+int32(-1073743872)) < base.Ui32(int32(-1073741824)) {
		v13 = int32(28)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v10
		v13 = v10
	}
	return v13
}
func F_pthread_cancel(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F_pthread_cond_init(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F_pthread_mutex_init(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F_pthread_setcancelstate(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F_pthread_testcancel(m *base.Module) {
	return
}
