package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaF_close(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v7 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = v7
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if base.Ui32(v17) < base.Ui32(l1) {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	goto L1
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v19
	v22 = v13 + int32(16)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
	if v23&(v24^int32(-1))&int32(3) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v99 != 0 {
		v13 = v99
		goto L3
	} else {
		goto L22
	}
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v22
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v56)+28)) = v13
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)))
	if v60&int32(7) != 0 {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	if v17 == v22 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v41 = F_luaM_realloc_(m, l0, v13, int32(32), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v36
	goto L9
L11:
	;
	return
L12:
	;
	goto L6
L13:
	;
	goto L6
L14:
	;
	goto L13
L15:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+21)))
	if v63 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+20)))
	v92 = v89&int32(3) | v60
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)) = uint8(v92)
	goto L14
L17:
	;
	v66 = int32(4)
	v67 = v60 | v66
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)) = uint8(v67)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	if v70 < v66 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+5)))
	if v74&int32(3) == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+21)))
	if v80 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+20)))
	v87 = v84&int32(3) | v60
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)) = uint8(v87)
	goto L13
L21:
	;
	F_reallymarkobject(m, v79, v73)
	mBase = m.M
	goto L13
L22:
	;
	goto L4
}
func F_luaF_getlocalname(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if int32(1) <= v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = int32(0)
	v15 = l1
	goto L4
L2:
	;
	return int32(0)
L3:
	;
	return int32(0)
L4:
	;
	v22 = v12 + v14*int32(12)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if l2 < v23 {
		goto L3
	} else {
		goto L6
	}
L5:
	;
	goto L3
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v25 <= l2 {
		v33 = v15
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v35 = v14 + int32(1)
	if v35 != v7 {
		v14 = v35
		v15 = v33
		goto L4
	} else {
		goto L10
	}
L8:
	;
	v28 = v15 + int32(-1)
	if v28 != 0 {
		v33 = v28
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	return v29 + int32(16)
L10:
	;
	goto L5
}
func F_luaF_newupval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v2 = int32(0)
	v6 = F_luaM_realloc_(m, l0, v2, v2, int32(32))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(10)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v6
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v10)
		v18 = v15 & int32(3)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v6 + int32(16)
		return v6
	}
}
