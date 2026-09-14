package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaX_next(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v3
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v5 == int32(287) {
		v24 = F_llex(m, l0, l0+int32(24))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v24
			return
		}
	} else {
		v11 = l0 + int32(32)
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		*(*int64)(unsafe.Add(mBase, uint32(l0+int32(16)))) = v12
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(40))))
		*(*int64)(unsafe.Add(mBase, uint32(l0+int32(24)))) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(287)
		return
	}
}
