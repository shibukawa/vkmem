package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaC_barrierback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	v6 = v4 & int32(251)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v6)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v3)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = l1
	return
}
func F_luaC_barrierf(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+21)))
	if v5 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+20)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	v81 = v75&int32(3) | v78&int32(248)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v81)
	return
L2:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	v12 = l2
	v13 = v10
	goto L11
L3:
	;
	return
L4:
	;
	goto L3
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v12
	goto L4
L6:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v66
	goto L5
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v64
	goto L5
L8:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v62
	goto L5
L9:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v60
	goto L5
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v40 < int32(4) {
		v51 = v39
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v16 = v13 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)) = uint8(v16)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)))
	if v18 == int32(7) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v24 = v16 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)) = uint8(v24)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v26 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	switch v18 + int32(-5) {
	case 0:
		goto L8
	case 1:
		goto L9
	default:
		goto L4
	case 3:
		goto L7
	case 4:
		goto L6
	case 5:
		goto L10
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+5)))
	if v36&int32(3) != 0 {
		v12 = v35
		v13 = v36
		goto L11
	} else {
		goto L18
	}
L16:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+5)))
	if v29&int32(3) == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_reallymarkobject(m, v4, v26)
	mBase = m.M
	goto L15
L18:
	;
	goto L4
L19:
	;
	if v51 != v12+int32(16) {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+5)))
	if v44&int32(3) == int32(0) {
		v51 = v39
		goto L19
	} else {
		goto L21
	}
L21:
	;
	F_reallymarkobject(m, v4, v43)
	mBase = m.M
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v51 = v50
	goto L19
L22:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)))
	v58 = v56 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)) = uint8(v58)
	goto L3
}
func F_luaC_link(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v3 = l2
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = l1
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v3)
	v11 = v8 & int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v11)
	return
}
func F_luaC_linkupval(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = l1
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v9&int32(7) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+21)))
	if v12 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+20)))
	v107 = v104&int32(3) | v9
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v107)
	goto L1
L4:
	;
	v15 = int32(4)
	v16 = v9 | v15
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v16)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v19 < v15 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
	if v23&int32(3) == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+21)))
	if v29 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+20)))
	v102 = v99&int32(3) | v9
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v102)
	return
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
	v36 = v22
	v37 = v34
	goto L17
L9:
	;
	return
L10:
	;
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v36
	goto L10
L12:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v90
	goto L11
L13:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v88
	goto L11
L14:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v86
	goto L11
L15:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v84
	goto L11
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v64 < int32(4) {
		v75 = v63
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v40 = v37 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+5)) = uint8(v40)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
	if v42 == int32(7) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v48 = v40 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+5)) = uint8(v48)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v50 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	switch v42 + int32(-5) {
	case 0:
		goto L14
	case 1:
		goto L15
	default:
		goto L10
	case 3:
		goto L13
	case 4:
		goto L12
	case 5:
		goto L16
	}
L21:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+5)))
	if v60&int32(3) != 0 {
		v36 = v59
		v37 = v60
		goto L17
	} else {
		goto L24
	}
L22:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+5)))
	if v53&int32(3) == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_reallymarkobject(m, v28, v50)
	mBase = m.M
	goto L21
L24:
	;
	goto L10
L25:
	;
	if v75 != v36+int32(16) {
		goto L10
	} else {
		goto L28
	}
L26:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+5)))
	if v68&int32(3) == int32(0) {
		v75 = v63
		goto L25
	} else {
		goto L27
	}
L27:
	;
	F_reallymarkobject(m, v28, v67)
	mBase = m.M
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v75 = v74
	goto L25
L28:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+5)))
	v82 = v80 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+5)) = uint8(v82)
	goto L9
}
