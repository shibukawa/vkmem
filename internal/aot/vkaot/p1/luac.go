package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaC_freeall(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = int32(67)
	*(*uint8)(unsafe.Add(mBase, uint32(v4)+20)) = uint8(v5)
	v10 = F_sweeplist(m, l0, v4+int32(28), int32(-3))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v12 < int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v18 = int32(0)
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v24 = F_sweeplist(m, l0, v19+v18<<(uint(int32(2))%32), int32(-3))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L3
L7:
	;
	v27 = v18 + int32(1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v27 < v28 {
		v18 = v27
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
func F_luaC_fullgc(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+21)))
	if base.Ui32(int32(1)) < base.Ui32(v5) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v32 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v34)+36)) = int64(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+112))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+5)))
	if v40&int32(3) == v32 {
		v47 = v39
		goto L12
	} else {
		goto L13
	}
L2:
	;
	goto L6
L3:
	;
	if v5 == int32(4) {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+44)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v4)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = v8
	v14 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v4)+21)) = uint8(v14)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = v4 + int32(28)
	goto L2
L5:
	;
	goto L2
L6:
	;
	v24 = F_singlestep(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	return
L9:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+21)))
	if v26 != int32(4) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+21)))
	if v74 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+80))
	if v48 < int32(4) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	F_reallymarkobject(m, v34, v39)
	mBase = m.M
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+112))
	v47 = v46
	goto L12
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	if v60 < int32(4) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+72))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+5)))
	if v52&int32(3) == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	F_reallymarkobject(m, v34, v51)
	mBase = m.M
	goto L14
L17:
	;
	F_markmt(m, v34)
	mBase = m.M
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+21)) = uint8(v72)
	goto L11
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+96))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
	if v64&int32(3) == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_reallymarkobject(m, v34, v63)
	mBase = m.M
	goto L17
L20:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v4)+72))
	v88 = base.I32_div_u_s(v86, int32(100))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v4)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+64)) = v88 * v89
	return
L21:
	;
	goto L22
L22:
	;
	v80 = F_singlestep(m, l0)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L24
	}
L23:
	;
	goto L20
L24:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+21)))
	if v82 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
}
