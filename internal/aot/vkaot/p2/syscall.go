package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___syscall_getpid(m *base.Module) int32 {
	return int32(42)
}
func F___syscall_setrlimit(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F___syscall_ugetrlimit(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	m.Env.Emscripten_err(m, int32(_a_F___syscall_ugetrlimit_0))
	mBase = m.M
	v7 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l1+int32(8)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v7
	return int32(0)
}
func F___syscall_uname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	if l0 != 0 {
		v6 = int32(0)
		v7 = *(*int32)(unsafe.Add(mBase, _c_F___syscall_uname[0]))
		*(*int32)(unsafe.Add(mBase, uint32(l0+int32(7)))) = v7
		v10 = *(*int64)(unsafe.Add(mBase, _c_F___syscall_uname[1]))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v10
		v13 = *(*int64)(unsafe.Add(mBase, _c_F___syscall_uname[2]))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+65)) = v13
		v18 = *(*int32)(unsafe.Add(mBase, _c_F___syscall_uname[3]))
		*(*int32)(unsafe.Add(mBase, uint32(l0+int32(72)))) = v18
		v21 = *(*int32)(unsafe.Add(mBase, _c_F___syscall_uname[4]))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+130)) = v21
		v26 = *(*int32)(unsafe.Add(mBase, _c_F___syscall_uname[5]))
		*(*int32)(unsafe.Add(mBase, uint32(l0+int32(133)))) = v26
		v29 = int32(*(*uint16)(unsafe.Add(mBase, _c_F___syscall_uname[6])))
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+195)) = uint16(v29)
		v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F___syscall_uname[7])))
		*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(197)))) = uint8(v34)
		v39 = *(*int32)(unsafe.Add(mBase, _c_F___syscall_uname[8]))
		*(*int32)(unsafe.Add(mBase, uint32(l0+int32(263)))) = v39
		v42 = *(*int32)(unsafe.Add(mBase, _c_F___syscall_uname[9]))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v42
		return v6
	} else {
		return int32(-21)
	}
}
