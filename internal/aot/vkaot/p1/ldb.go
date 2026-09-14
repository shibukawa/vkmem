package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_ldbDisable(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	v1 = int32(0)
	v2 = m.G6
	*(*int32)(unsafe.Add(mBase, uint32(v2))) = v1
	*(*int32)(unsafe.Add(mBase, uint32(v2)+264)) = v1
	return
}
func F_ldbEnable(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = m.G6
	*(*int32)(unsafe.Add(mBase, uint32(v2))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+268)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2)+260)) = int64(4294967296)
	return
}
func F_ldbLog(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = m.G12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	m.T0[v4].(func(*base.Module, int32, int32))(m, l0, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_ldbLogCString(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	v2 = int32(0)
	v4 = m.G13
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = m.G12
	if l0&int32(3) == v2 {
		v29 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v63 = m.T0[v5].(func(*base.Module, int32, int32, int32) int32)(m, v2, l0, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v62 = v54 - l0
	goto L1
L3:
	;
	v33 = v29
	goto L11
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v18 = l0
	goto L7
L6:
	;
	v62 = l0 - l0
	goto L1
L7:
	;
	v22 = v18 + int32(1)
	if v22&int32(3) == int32(0) {
		v29 = v22
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v27 != 0 {
		v18 = v22
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v54 = v22
	goto L2
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v42 = int32(-2139062144)
	if (int32(16843008)-v39|v39)&v42 == v42 {
		v33 = v33 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v48 = v33
	goto L14
L13:
	;
	goto L12
L14:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v52 != 0 {
		v48 = v48 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v54 = v48
	goto L2
L16:
	;
	goto L15
L17:
	;
	return
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	m.T0[v66].(func(*base.Module, int32, int32))(m, v63, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	return
}
