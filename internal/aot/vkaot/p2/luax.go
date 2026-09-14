package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaX_lookahead(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = F_llex(m, l0, l0+int32(40))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v4
		return
	}
}
func F_luaX_setinput(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v5 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+68)) = uint8(v5)
	v7 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(287)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = int64(4294967297)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v22 = F_luaM_realloc_(m, l0, v19, v20, int32(32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
		*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = int32(32)
		*(*int32)(unsafe.Add(mBase, uint32(v24))) = v22
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
		*(*int32)(unsafe.Add(mBase, uint32(v28))) = v29 + int32(-1)
		if v29 == int32(0) {
			v41 = F_luaZ_fill(m, v28)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v41
				return
			}
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v35 + int32(1)
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v39
			return
		}
	}
}
func F_luaX_syntaxerror(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_luaX_lexerror(m, l0, l1, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
