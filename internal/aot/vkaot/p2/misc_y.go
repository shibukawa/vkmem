package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_yesnotoi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	v2 = int32(_a_F_yesnotoi_0)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v45 = int32(_a_F_yesnotoi_1)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v48 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	if v37-v39 != 0 {
		goto L1
	} else {
		goto L14
	}
L3:
	;
	v37 = F_tolower(m, v33)
	mBase = m.M
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v39 = F_tolower(m, v38)
	mBase = m.M
	goto L2
L4:
	;
	v7 = l0
	v8 = v2
	v9 = v5
	goto L7
L5:
	;
	v33 = int32(0)
	v34 = v2
	goto L3
L6:
	;
	v33 = v30 & int32(255)
	v34 = v29
	goto L3
L7:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v11 == int32(0) {
		v29 = v8
		v30 = v9
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v29 = v23
	v30 = int32(0)
	goto L6
L9:
	;
	v15 = v9 & int32(255)
	if v15 == v11 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v22 = int32(1)
	v23 = v8 + v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v24 != 0 {
		v7 = v7 + v22
		v8 = v23
		v9 = v24
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v17 = F_tolower(m, v15)
	mBase = m.M
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v19 = F_tolower(m, v18)
	mBase = m.M
	if v17 == v19 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v29 = v8
	v30 = v21
	goto L6
L13:
	;
	goto L8
L14:
	;
	return int32(1)
L15:
	;
	if v80-v82 != 0 {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v80 = F_tolower(m, v76)
	mBase = m.M
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v82 = F_tolower(m, v81)
	mBase = m.M
	goto L15
L17:
	;
	v50 = l0
	v51 = v45
	v52 = v48
	goto L20
L18:
	;
	v76 = int32(0)
	v77 = v45
	goto L16
L19:
	;
	v76 = v73 & int32(255)
	v77 = v72
	goto L16
L20:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v54 == int32(0) {
		v72 = v51
		v73 = v52
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v72 = v66
	v73 = int32(0)
	goto L19
L22:
	;
	v58 = v52 & int32(255)
	if v58 == v54 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v65 = int32(1)
	v66 = v51 + v65
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v67 != 0 {
		v50 = v50 + v65
		v51 = v66
		v52 = v67
		goto L20
	} else {
		goto L26
	}
L24:
	;
	v60 = F_tolower(m, v58)
	mBase = m.M
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v62 = F_tolower(m, v61)
	mBase = m.M
	if v60 == v62 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v72 = v51
	v73 = v64
	goto L19
L26:
	;
	goto L21
L27:
	;
	v84 = int32(-1)
	goto L29
L28:
	;
	v84 = int32(0)
	goto L29
L29:
	;
	return v84
}
