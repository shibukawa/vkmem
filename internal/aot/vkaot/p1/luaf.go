package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaF_freeclosure(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	if v7 != 0 {
		v8 = int32(4)
	} else {
		v8 = int32(2)
	}
	if v7 != 0 {
		v12 = int32(24)
	} else {
		v12 = int32(20)
	}
	v15 = F_luaM_realloc_(m, l0, l1, v4<<(uint(v8)%32)+v12, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		return
	}
}
func F_luaF_freeupval(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v4 == l1+int32(16) {
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v9
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v11
	}
	v16 = F_luaM_realloc_(m, l0, l1, int32(32), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		return
	}
}
