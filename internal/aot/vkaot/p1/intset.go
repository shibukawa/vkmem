package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_intsetLen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	return v2
}
func F_intsetMin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v2 + int32(-4) {
	case 0:
		v7 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+8)))
		return v7
	default:
		v9 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
		return v9
	case 4:
		v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		return v5
	}
}
func F_intsetNew(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(2)
		return v3
	}
}
