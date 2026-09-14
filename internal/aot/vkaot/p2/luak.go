package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaK_concat(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	if l2 == int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l2
	return
L2:
	;
	return
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v10 == int32(-1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v18 = v10
	goto L6
L5:
	;
	v40 = l2 + (v18 ^ int32(-1))
	v42 = v40 >> (uint(int32(31)) % 32)
	if base.Ui32(v40^v42-v42) < base.Ui32(int32(131072)) {
		v55 = v25
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v24 = v14 + v18<<(uint(int32(2))%32)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v29 = int32(base.Ui32(v25)>>(uint(int32(14))%32)) + int32(-131071)
	if v29 == int32(-1) {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v34 = v18 + v29 + int32(1)
	if v34 != int32(-1) {
		v18 = v34
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v40<<(uint(int32(14))%32) | v55&int32(16383) + int32(2147467264)
	goto L2
L11:
	;
	v47 = m.G3
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v48, v47+int32(_a2034))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v55 = v53
	goto L10
}
func F_luaK_exp2RK(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_luaK_dischargevars(m, l0, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v12 == v11 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v33 + int32(-1) {
	case 0, 1, 2, 4:
		goto L14
	case 3:
		goto L13
	default:
		goto L12
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v18 != int32(12) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_luaK_exp2nextreg(m, l0, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v21 == v22 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v24 < v25 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	F_exp2reg(m, l0, l1, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L3
L10:
	;
	goto L3
L11:
	;
	m.G0 = v9 + int32(32)
	return v111
L12:
	;
	F_luaK_dischargevars(m, l0, l1)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L24
	}
L13:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(255) < v80 {
		goto L12
	} else {
		goto L23
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if int32(255) < v36 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	switch v33 + int32(-1) {
	case 0:
		goto L19
	default:
		goto L17
	case 4:
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v72
	v111 = v72 | int32(256)
	goto L11
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = base.B2i32(v33 == int32(2))
	v67 = v9 + int32(16)
	v70 = F_addk(m, l0, v67, v67)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L22
	}
L18:
	;
	v51 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v51
	v56 = v9 + int32(16)
	v59 = F_addk(m, l0, v56, v56)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v43
	v49 = F_addk(m, l0, v9+int32(16), v9)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v72 = v49
	goto L16
L21:
	;
	v72 = v59
	goto L16
L22:
	;
	v72 = v70
	goto L16
L23:
	;
	v111 = v80 | int32(256)
	goto L11
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v88 != int32(12) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v111 = v109
	goto L11
L26:
	;
	F_luaK_exp2nextreg(m, l0, l1)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L31
	}
L27:
	;
	v92 = l1 + int32(8)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v93 == v94 {
		v107 = v92
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v96 < v97 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	F_exp2reg(m, l0, l1, v96)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v107 = v92
	goto L25
L31:
	;
	v107 = l1 + int32(8)
	goto L25
}
func F_luaK_exp2val(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_luaK_dischargevars(m, l0, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		if v6 == v5 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v10 != int32(12) {
				F_luaK_exp2nextreg(m, l0, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				if v13 == v14 {
					return
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
					if v16 < v17 {
						F_luaK_exp2nextreg(m, l0, l1)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							return
						}
					} else {
						F_exp2reg(m, l0, l1, v16)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_luaK_infix(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	switch l1 {
	case 0, 1, 2, 3, 4, 5:
		goto L5
	case 6:
		goto L6
	default:
		goto L4
	case 13:
		goto L8
	case 14:
		goto L7
	}
L1:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v103
	if v102 == int32(-1) {
		goto L33
	} else {
		goto L34
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v20
	goto L1
L3:
	;
	return
L4:
	;
	v91 = F_luaK_exp2RK(m, l0, l2)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L32
	}
L5:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v80 != int32(5) {
		goto L27
	} else {
		goto L28
	}
L6:
	;
	F_luaK_exp2nextreg(m, l0, l2)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L26
	}
L7:
	;
	F_luaK_dischargevars(m, l0, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L9
	} else {
		goto L11
	}
L8:
	;
	F_luaK_goiftrue(m, l0, l2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	return
L11:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v13 + int32(-1) {
	case 0, 2:
		goto L1
	default:
		goto L13
	case 9:
		goto L14
	}
L12:
	;
	if v20 == int32(-1) {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v18 = F_jumponcond(m, l0, l2, int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L15
	}
L14:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v20 = v16
	goto L12
L15:
	;
	v20 = v18
	goto L12
L16:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v23 == int32(-1) {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v29 = v23
	goto L19
L18:
	;
	v54 = v20 + (v29 ^ int32(-1))
	v56 = v54 >> (uint(int32(31)) % 32)
	if base.Ui32(v54^v56-v56) < base.Ui32(int32(131072)) {
		v69 = v39
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v38 = v27 + v29<<(uint(int32(2))%32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v43 = int32(base.Ui32(v39)>>(uint(int32(14))%32)) + int32(-131071)
	if v43 == int32(-1) {
		goto L18
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	v48 = v29 + v43 + int32(1)
	if v48 != int32(-1) {
		v29 = v48
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v54<<(uint(int32(14))%32) | v69&int32(16383) + int32(2147467264)
	goto L1
L24:
	;
	v61 = m.G3
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v62, v61+int32(_a2034))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v69 = v67
	goto L23
L26:
	;
	return
L27:
	;
	v89 = F_luaK_exp2RK(m, l0, l2)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L31
	}
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v83 != int32(-1) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v86 == int32(-1) {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	return
L32:
	;
	goto L3
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = int32(-1)
	return
L34:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v107 == int32(-1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v102
	goto L33
L36:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v113 = v107
	goto L38
L37:
	;
	v138 = v102 + (v113 ^ int32(-1))
	v140 = v138 >> (uint(int32(31)) % 32)
	if base.Ui32(v138^v140-v140) < base.Ui32(int32(131072)) {
		v153 = v123
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v122 = v111 + v113<<(uint(int32(2))%32)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v127 = int32(base.Ui32(v123)>>(uint(int32(14))%32)) + int32(-131071)
	if v127 == int32(-1) {
		goto L37
	} else {
		goto L40
	}
L39:
	;
	goto L37
L40:
	;
	v132 = v113 + v127 + int32(1)
	if v132 != int32(-1) {
		v113 = v132
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v138<<(uint(int32(14))%32) | v153&int32(16383) + int32(2147467264)
	goto L33
L43:
	;
	v145 = m.G3
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v146, v145+int32(_a2034))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v153 = v151
	goto L42
}
func F_luaK_numberK(m *base.Module, l0 int32, l1 float64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v6))) = l1
	v11 = F_addk(m, l0, v6, v6)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v11
	}
}
func F_luaK_reserveregs(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = v6 + l1
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+75)))
	if v9 < v7 {
		if base.Ui32(v7) < base.Ui32(int32(250)) {
			v23 = v8
			v24 = v7
			*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v7)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v24
			return
		} else {
			v14 = m.G3
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_luaX_syntaxerror(m, v15, v14+int32(_a2035))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v23 = v22
				v24 = v20 + l1
				*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v7)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v24
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v7
		return
	}
}
