package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_dictCStringKeyCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v6 == int32(0) {
		v29 = v5
		v30 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.B2i32(v30-v29&int32(255) == int32(0))
L2:
	;
	goto L1
L3:
	;
	if v6 != v5&int32(255) {
		v29 = v5
		v30 = v6
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v12 = l0
	v13 = l1
	goto L5
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v17 == int32(0) {
		v29 = v16
		v30 = v17
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v29 = v16
	v30 = v17
	goto L2
L7:
	;
	v20 = int32(1)
	if v17 == v16&int32(255) {
		v12 = v12 + v20
		v13 = v13 + v20
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
func F_dictCStringKeyHash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	if l0&int32(3) == int32(0) {
		v23 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v58 = F_siphash(m, l0, v56, int32(_a_F_dictCStringKeyHash_0))
	mBase = m.M
	goto L17
L2:
	;
	v56 = v48 - l0
	goto L1
L3:
	;
	v27 = v23
	goto L11
L4:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v12 = l0
	goto L7
L6:
	;
	v56 = l0 - l0
	goto L1
L7:
	;
	v16 = v12 + int32(1)
	if v16&int32(3) == int32(0) {
		v23 = v16
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v21 != 0 {
		v12 = v16
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v48 = v16
	goto L2
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v36 = int32(-2139062144)
	if (int32(16843008)-v33|v33)&v36 == v36 {
		v27 = v27 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v42 = v27
	goto L14
L13:
	;
	goto L12
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v46 != 0 {
		v42 = v42 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v48 = v42
	goto L2
L16:
	;
	goto L15
L17:
	;
	return v58
}
func F_dictClear(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	v12 = l0 + l1
	v14 = v12 + int32(26)
	v17 = l0 + l1<<(uint(int32(2))%32)
	v19 = v17 + int32(12)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+26)))
	if v20 == int32(255) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	F_valkey_free(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L9
	} else {
		goto L25
	}
L2:
	;
	v34 = int32(0)
	goto L3
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v37 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	goto L1
L5:
	;
	if l2 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(4))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v34<<(uint(int32(2))%32))))
	if v50 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	if v34&int32(65535) != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	m.T0[l2].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	goto L6
L11:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v97 == int32(255) {
		goto L1
	} else {
		goto L23
	}
L12:
	;
	v54 = v50
	goto L13
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	if v66 == int32(0) {
		v73 = v65
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L11
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	if v74 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	m.T0[v66].(func(*base.Module, int32))(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v73 = v72
	goto L15
L18:
	;
	F_valkey_free(m, v54)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L21
	}
L19:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	m.T0[v74].(func(*base.Module, int32))(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v82 + int32(-1)
	if v64 != 0 {
		v54 = v64
		goto L13
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	v101 = v34 + int32(1)
	if int32(base.Ui32(v101)>>(uint(v97)%32)) == int32(0) {
		v34 = v101
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L4
L25:
	;
	v119 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v119
	v121 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v121)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v119
	return
}
func F_dictEmpty(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(-1) {
		F_dictClear(m, l0, int32(0), l1)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			F_dictClear(m, l0, int32(1), l1)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v20)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v20)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(-1)
				return
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
		if v8 == int32(0) {
			F_dictClear(m, l0, int32(0), l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_dictClear(m, l0, int32(1), l1)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v20)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v20)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(-1)
					return
				}
			}
		} else {
			m.T0[v8].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_dictClear(m, l0, int32(0), l1)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_dictClear(m, l0, int32(1), l1)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v20 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v20)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v20)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(-1)
						return
					}
				}
			}
		}
	}
}
func F_dictEncObjKeyCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8&int32(240) != int32(16) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22&int32(-8) == int32(-16) {
		v31 = l0
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v13&int32(240) != int32(16) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = F_objectGetVal(m, l0)
	mBase = m.M
	v19 = F_objectGetVal(m, l1)
	mBase = m.M
	return base.B2i32(v18 == v19)
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v32&int32(-8) == int32(-16) {
		v39 = l1
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v27 = F_getDecodedObject(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v31 = v27
	goto L4
L8:
	;
	v40 = F_objectGetVal(m, v31)
	mBase = m.M
	v41 = F_objectGetVal(m, v39)
	mBase = m.M
	v42 = int32(0)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-1)))))
	switch v46 & int32(7) {
	case 0:
		goto L16
	case 1:
		goto L15
	case 2:
		goto L14
	case 3:
		goto L13
	case 4:
		goto L12
	default:
		v63 = v42
		goto L11
	}
L9:
	;
	v37 = F_getDecodedObject(m, l1)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v39 = v37
	goto L8
L11:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(-1)))))
	switch v66 & int32(7) {
	case 0:
		goto L22
	case 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		v83 = v42
		goto L17
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-17))))
	v63 = v62
	goto L11
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-9))))
	v63 = v59
	goto L11
L14:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+int32(-5)))))
	v63 = v56
	goto L11
L15:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-3)))))
	v63 = v53
	goto L11
L16:
	;
	v63 = int32(base.Ui32(v46) >> (uint(int32(3)) % 32))
	goto L11
L17:
	;
	if v63 != v83 {
		v152 = int32(0)
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-17))))
	v83 = v82
	goto L17
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-9))))
	v83 = v79
	goto L17
L20:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41+int32(-5)))))
	v83 = v76
	goto L17
L21:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(-3)))))
	v83 = v73
	goto L17
L22:
	;
	v83 = int32(base.Ui32(v66) >> (uint(int32(3)) % 32))
	goto L17
L23:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v153&int32(-8) == int32(-16) {
		goto L41
	} else {
		goto L42
	}
L24:
	;
	if base.Ui32(v63) < base.Ui32(int32(4)) {
		v109 = v40
		v110 = v41
		v111 = v63
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v152 = base.B2i32(v149 == int32(0))
	goto L23
L26:
	;
	v149 = int32(0)
	goto L25
L27:
	;
	v121 = v116
	v122 = v117
	v123 = v118
	goto L37
L28:
	;
	if v111 == int32(0) {
		goto L26
	} else {
		goto L35
	}
L29:
	;
	if (v41|v40)&int32(3) != 0 {
		v116 = v40
		v117 = v41
		v118 = v63
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v93 = v40
	v94 = v41
	v95 = v63
	goto L31
L31:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v98 != v99 {
		v116 = v93
		v117 = v94
		v118 = v95
		goto L27
	} else {
		goto L33
	}
L32:
	;
	v109 = v104
	v110 = v102
	v111 = v106
	goto L28
L33:
	;
	v101 = int32(4)
	v102 = v94 + v101
	v104 = v93 + v101
	v106 = v95 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v106) {
		v93 = v104
		v94 = v102
		v95 = v106
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v116 = v109
	v117 = v110
	v118 = v111
	goto L27
L36:
	;
	v149 = v126 - v127
	goto L25
L37:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v126 != v127 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v129 = int32(1)
	v134 = v123 + int32(-1)
	if v134 == int32(0) {
		goto L26
	} else {
		goto L40
	}
L40:
	;
	v121 = v121 + v129
	v122 = v122 + v129
	v123 = v134
	goto L37
L41:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v160&int32(-8) == int32(-16) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	F_decrRefCount(m, v31)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	return v152
L45:
	;
	F_decrRefCount(m, v39)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	goto L44
}
func F_dictExpandIfNeeded(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8 != int32(-1) {
		v87 = v2
		return v87
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if v11 != int32(255) {
			v18 = int32(1)
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_dictExpandIfNeeded[0]))
			switch v20 {
			case 0:
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if int32(base.Ui32(v22)>>(uint(v11)%32)) != 0 {
					v28 = v22
					v30 = v28 + int32(1)
					v31 = int32(0)
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
					if v33 != 0 {
						if base.Ui32(int32(5)) <= base.Ui32(v30) {
							if base.Ui32(v30) <= base.Ui32(int32(2147483646)) {
								v44 = int32(32) - base.I32_clz(v28)
							} else {
								v44 = int32(31)
							}
						} else {
							v44 = int32(2)
						}
						v52 = m.T0[v33].(func(*base.Module, int32, float64) int32)(m, int32(4)<<(uint(v44)%32), base.F64_div(base.F64_convert_i32_u(v28), base.F64_convert_i32_u(int32(1)<<(uint(v11)%32))))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v64 = base.B2i32(v58 != int32(-1))
								v65 = v61 + int32(1)
								v66 = v61
								if v64 != 0 {
									v87 = v31
									return v87
								} else {
									if v66 == int32(-1) {
										v87 = v31
										return v87
									} else {
										v69 = int32(0)
										v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
										if v72 == int32(255) {
											v76 = v69
										} else {
											v76 = int32(1) << (uint(v72) % 32)
										}
										if base.Ui32(v65) <= base.Ui32(v76) {
											v87 = v69
											return v87
										} else {
											v80 = v65
											v83 = int32(0)
											v85 = F_dictResizeWithOptionalCheck(m, l0, v80, v83)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												v87 = v83
												return v87
											}
										}
									}
								}
							} else {
								return int32(0)
							}
						}
					} else {
						v64 = int32(0)
						v65 = v30
						v66 = v28
						if v64 != 0 {
							v87 = v31
							return v87
						} else {
							if v66 == int32(-1) {
								v87 = v31
								return v87
							} else {
								v69 = int32(0)
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
								if v72 == int32(255) {
									v76 = v69
								} else {
									v76 = int32(1) << (uint(v72) % 32)
								}
								if base.Ui32(v65) <= base.Ui32(v76) {
									v87 = v69
									return v87
								} else {
									v80 = v65
									v83 = int32(0)
									v85 = F_dictResizeWithOptionalCheck(m, l0, v80, v83)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										v87 = v83
										return v87
									}
								}
							}
						}
					}
				} else {
					v24 = v22
					if base.Ui32(v24) < base.Ui32(int32(4)<<(uint(v11)%32)) {
						v87 = v18
						return v87
					} else {
						v28 = v24
						v30 = v28 + int32(1)
						v31 = int32(0)
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
						if v33 != 0 {
							if base.Ui32(int32(5)) <= base.Ui32(v30) {
								if base.Ui32(v30) <= base.Ui32(int32(2147483646)) {
									v44 = int32(32) - base.I32_clz(v28)
								} else {
									v44 = int32(31)
								}
							} else {
								v44 = int32(2)
							}
							v52 = m.T0[v33].(func(*base.Module, int32, float64) int32)(m, int32(4)<<(uint(v44)%32), base.F64_div(base.F64_convert_i32_u(v28), base.F64_convert_i32_u(int32(1)<<(uint(v11)%32))))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								if v52 != 0 {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v64 = base.B2i32(v58 != int32(-1))
									v65 = v61 + int32(1)
									v66 = v61
									if v64 != 0 {
										v87 = v31
										return v87
									} else {
										if v66 == int32(-1) {
											v87 = v31
											return v87
										} else {
											v69 = int32(0)
											v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
											if v72 == int32(255) {
												v76 = v69
											} else {
												v76 = int32(1) << (uint(v72) % 32)
											}
											if base.Ui32(v65) <= base.Ui32(v76) {
												v87 = v69
												return v87
											} else {
												v80 = v65
												v83 = int32(0)
												v85 = F_dictResizeWithOptionalCheck(m, l0, v80, v83)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													v87 = v83
													return v87
												}
											}
										}
									}
								} else {
									return int32(0)
								}
							}
						} else {
							v64 = int32(0)
							v65 = v30
							v66 = v28
							if v64 != 0 {
								v87 = v31
								return v87
							} else {
								if v66 == int32(-1) {
									v87 = v31
									return v87
								} else {
									v69 = int32(0)
									v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
									if v72 == int32(255) {
										v76 = v69
									} else {
										v76 = int32(1) << (uint(v72) % 32)
									}
									if base.Ui32(v65) <= base.Ui32(v76) {
										v87 = v69
										return v87
									} else {
										v80 = v65
										v83 = int32(0)
										v85 = F_dictResizeWithOptionalCheck(m, l0, v80, v83)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											v87 = v83
											return v87
										}
									}
								}
							}
						}
					}
				}
			default:
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v24 = v21
				if base.Ui32(v24) < base.Ui32(int32(4)<<(uint(v11)%32)) {
					v87 = v18
					return v87
				} else {
					v28 = v24
					v30 = v28 + int32(1)
					v31 = int32(0)
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
					if v33 != 0 {
						if base.Ui32(int32(5)) <= base.Ui32(v30) {
							if base.Ui32(v30) <= base.Ui32(int32(2147483646)) {
								v44 = int32(32) - base.I32_clz(v28)
							} else {
								v44 = int32(31)
							}
						} else {
							v44 = int32(2)
						}
						v52 = m.T0[v33].(func(*base.Module, int32, float64) int32)(m, int32(4)<<(uint(v44)%32), base.F64_div(base.F64_convert_i32_u(v28), base.F64_convert_i32_u(int32(1)<<(uint(v11)%32))))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v64 = base.B2i32(v58 != int32(-1))
								v65 = v61 + int32(1)
								v66 = v61
								if v64 != 0 {
									v87 = v31
									return v87
								} else {
									if v66 == int32(-1) {
										v87 = v31
										return v87
									} else {
										v69 = int32(0)
										v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
										if v72 == int32(255) {
											v76 = v69
										} else {
											v76 = int32(1) << (uint(v72) % 32)
										}
										if base.Ui32(v65) <= base.Ui32(v76) {
											v87 = v69
											return v87
										} else {
											v80 = v65
											v83 = int32(0)
											v85 = F_dictResizeWithOptionalCheck(m, l0, v80, v83)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												v87 = v83
												return v87
											}
										}
									}
								}
							} else {
								return int32(0)
							}
						}
					} else {
						v64 = int32(0)
						v65 = v30
						v66 = v28
						if v64 != 0 {
							v87 = v31
							return v87
						} else {
							if v66 == int32(-1) {
								v87 = v31
								return v87
							} else {
								v69 = int32(0)
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
								if v72 == int32(255) {
									v76 = v69
								} else {
									v76 = int32(1) << (uint(v72) % 32)
								}
								if base.Ui32(v65) <= base.Ui32(v76) {
									v87 = v69
									return v87
								} else {
									v80 = v65
									v83 = int32(0)
									v85 = F_dictResizeWithOptionalCheck(m, l0, v80, v83)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										v87 = v83
										return v87
									}
								}
							}
						}
					}
				}
			case 2:
				v87 = v18
				return v87
			}
		} else {
			v14 = int32(4)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if base.Ui32(v15) <= base.Ui32(v14) {
				v80 = v14
				v83 = int32(0)
				v85 = F_dictResizeWithOptionalCheck(m, l0, v80, v83)
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					v87 = v83
					return v87
				}
			} else {
				v87 = v2
				return v87
			}
		}
	}
}
func F_dictGenCaseHashFunction(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v185 int64
	_ = v185
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v201 int64
	_ = v201
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	var v217 int64
	_ = v217
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v233 int64
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v250 int64
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v267 int64
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v284 int64
	_ = v284
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v304 int64
	_ = v304
	var v305 int64
	_ = v305
	var v306 int64
	_ = v306
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v312 int64
	_ = v312
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v323 int64
	_ = v323
	var v324 int64
	_ = v324
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v334 int64
	_ = v334
	var v335 int64
	_ = v335
	var v340 int64
	_ = v340
	var v341 int64
	_ = v341
	var v344 int64
	_ = v344
	var v350 int64
	_ = v350
	var v354 int64
	_ = v354
	var v358 int64
	_ = v358
	v3 = int32(_a_F_dictGenCaseHashFunction_0)
	v13 = *(*int64)(unsafe.Add(mBase, _c_F_dictGenCaseHashFunction[0]))
	v15 = v13 ^ int64(8317987319222330741)
	v16 = *(*int64)(unsafe.Add(mBase, _c_F_dictGenCaseHashFunction[1]))
	v18 = v16 ^ int64(7237128888997146477)
	v20 = v13 ^ int64(7816392313619706465)
	v22 = v16 ^ int64(8387220255154660723)
	v26 = l1 & int32(7)
	v27 = l0 + l1 - v26
	if l0 == v27 {
		v172 = l0
		v175 = v20
		v176 = v15
		v177 = v22
		v178 = v18
	} else {
		v29 = l0
		v32 = v20
		v33 = v15
		v34 = v22
		v35 = v18
		for {
			v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
			if base.Ui32((v41+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v50 = v41 | int32(32)
			} else {
				v50 = v41
			}
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
			v56 = v54 << (uint(int32(8)) % 32)
			if base.Ui32((v54+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v65 = v56 | int32(8192)
			} else {
				v65 = v56
			}
			v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
			if base.Ui32((v66+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v75 = v66 | int32(32)
			} else {
				v75 = v66
			}
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
			v79 = v77 << (uint(int32(16)) % 32)
			if base.Ui32((v77+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v88 = v79 | int32(2097152)
			} else {
				v88 = v79
			}
			v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)))
			v92 = v90 << (uint(int32(24)) % 32)
			if base.Ui32((v90+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v101 = v92 | int32(536870912)
			} else {
				v101 = v92
			}
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)))
			if base.Ui32((v105+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v114 = v105 | int32(32)
			} else {
				v114 = v105
			}
			v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+6)))
			if base.Ui32((v119+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v128 = v119 | int32(32)
			} else {
				v128 = v119
			}
			v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+7)))
			if base.Ui32((v133+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v142 = v133 | int32(32)
			} else {
				v142 = v133
			}
			v146 = base.I64_extend_i32_u(v50)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v65|v75|v88|v101) | base.I64_extend_i32_u(v114)<<(uint(int64(40))%64) | base.I64_extend_i32_u(v128)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v142)<<(uint(int64(56))%64)
			v147 = v146 ^ v34
			v150 = v147 + v32
			v151 = base.I64_rotl(v147, int64(16)) ^ v150
			v154 = v33 + v35
			v155 = int64(32)
			v157 = v151 + base.I64_rotl(v154, v155)
			v158 = base.I64_rotl(v151, int64(21)) ^ v157
			v159 = v157 ^ v146
			v162 = v154 ^ base.I64_rotl(v35, int64(13))
			v163 = v150 + v162
			v165 = base.I64_rotl(v163, v155)
			v168 = v163 ^ base.I64_rotl(v162, int64(17))
			v170 = v29 + int32(8)
			if v170 != v27 {
				v29 = v170
				v32 = v165
				v33 = v159
				v34 = v158
				v35 = v168
				continue
			} else {
				break
			}
			break
		}
		v172 = v27
		v175 = v165
		v176 = v159
		v177 = v158
		v178 = v168
	}
	v185 = base.I64_extend_i32_u(l1) << (uint(int64(56)) % 64)
	switch v26 {
	default:
		v300 = v185
	case 1:
		v284 = v185
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	case 2:
		v267 = v185
		v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
		v270 = v268 << (uint(int32(8)) % 32)
		if base.Ui32((v268+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v279 = v270 | int32(8192)
		} else {
			v279 = v270
		}
		v284 = v267 | base.I64_extend_i32_u(v279)
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	case 3:
		v250 = v185
		v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+2)))
		v253 = v251 << (uint(int32(16)) % 32)
		if base.Ui32((v251+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v262 = v253 | int32(2097152)
		} else {
			v262 = v253
		}
		v267 = v250 | base.I64_extend_i32_u(v262)
		v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
		v270 = v268 << (uint(int32(8)) % 32)
		if base.Ui32((v268+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v279 = v270 | int32(8192)
		} else {
			v279 = v270
		}
		v284 = v267 | base.I64_extend_i32_u(v279)
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	case 4:
		v233 = v185
		v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+3)))
		v236 = v234 << (uint(int32(24)) % 32)
		if base.Ui32((v234+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v245 = v236 | int32(536870912)
		} else {
			v245 = v236
		}
		v250 = v233 | base.I64_extend_i32_u(v245)
		v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+2)))
		v253 = v251 << (uint(int32(16)) % 32)
		if base.Ui32((v251+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v262 = v253 | int32(2097152)
		} else {
			v262 = v253
		}
		v267 = v250 | base.I64_extend_i32_u(v262)
		v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
		v270 = v268 << (uint(int32(8)) % 32)
		if base.Ui32((v268+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v279 = v270 | int32(8192)
		} else {
			v279 = v270
		}
		v284 = v267 | base.I64_extend_i32_u(v279)
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	case 5:
		v217 = v185
		v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+4)))
		if base.Ui32((v218+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v227 = v218 | int32(32)
		} else {
			v227 = v218
		}
		v233 = base.I64_extend_i32_u(v227)<<(uint(int64(32))%64) | v217
		v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+3)))
		v236 = v234 << (uint(int32(24)) % 32)
		if base.Ui32((v234+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v245 = v236 | int32(536870912)
		} else {
			v245 = v236
		}
		v250 = v233 | base.I64_extend_i32_u(v245)
		v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+2)))
		v253 = v251 << (uint(int32(16)) % 32)
		if base.Ui32((v251+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v262 = v253 | int32(2097152)
		} else {
			v262 = v253
		}
		v267 = v250 | base.I64_extend_i32_u(v262)
		v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
		v270 = v268 << (uint(int32(8)) % 32)
		if base.Ui32((v268+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v279 = v270 | int32(8192)
		} else {
			v279 = v270
		}
		v284 = v267 | base.I64_extend_i32_u(v279)
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	case 6:
		v201 = v185
		v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+5)))
		if base.Ui32((v202+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v211 = v202 | int32(32)
		} else {
			v211 = v202
		}
		v217 = base.I64_extend_i32_u(v211)<<(uint(int64(40))%64) | v201
		v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+4)))
		if base.Ui32((v218+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v227 = v218 | int32(32)
		} else {
			v227 = v218
		}
		v233 = base.I64_extend_i32_u(v227)<<(uint(int64(32))%64) | v217
		v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+3)))
		v236 = v234 << (uint(int32(24)) % 32)
		if base.Ui32((v234+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v245 = v236 | int32(536870912)
		} else {
			v245 = v236
		}
		v250 = v233 | base.I64_extend_i32_u(v245)
		v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+2)))
		v253 = v251 << (uint(int32(16)) % 32)
		if base.Ui32((v251+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v262 = v253 | int32(2097152)
		} else {
			v262 = v253
		}
		v267 = v250 | base.I64_extend_i32_u(v262)
		v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
		v270 = v268 << (uint(int32(8)) % 32)
		if base.Ui32((v268+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v279 = v270 | int32(8192)
		} else {
			v279 = v270
		}
		v284 = v267 | base.I64_extend_i32_u(v279)
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	case 7:
		v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+6)))
		if base.Ui32((v186+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v195 = v186 | int32(32)
		} else {
			v195 = v186
		}
		v201 = base.I64_extend_i32_u(v195)<<(uint(int64(48))%64) | v185
		v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+5)))
		if base.Ui32((v202+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v211 = v202 | int32(32)
		} else {
			v211 = v202
		}
		v217 = base.I64_extend_i32_u(v211)<<(uint(int64(40))%64) | v201
		v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+4)))
		if base.Ui32((v218+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v227 = v218 | int32(32)
		} else {
			v227 = v218
		}
		v233 = base.I64_extend_i32_u(v227)<<(uint(int64(32))%64) | v217
		v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+3)))
		v236 = v234 << (uint(int32(24)) % 32)
		if base.Ui32((v234+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v245 = v236 | int32(536870912)
		} else {
			v245 = v236
		}
		v250 = v233 | base.I64_extend_i32_u(v245)
		v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+2)))
		v253 = v251 << (uint(int32(16)) % 32)
		if base.Ui32((v251+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v262 = v253 | int32(2097152)
		} else {
			v262 = v253
		}
		v267 = v250 | base.I64_extend_i32_u(v262)
		v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
		v270 = v268 << (uint(int32(8)) % 32)
		if base.Ui32((v268+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v279 = v270 | int32(8192)
		} else {
			v279 = v270
		}
		v284 = v267 | base.I64_extend_i32_u(v279)
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	}
	v301 = v300 ^ v177
	v302 = int64(16)
	v304 = v301 + v175
	v305 = base.I64_rotl(v301, v302) ^ v304
	v306 = int64(21)
	v308 = v176 + v178
	v309 = int64(32)
	v311 = v305 + base.I64_rotl(v308, v309)
	v312 = base.I64_rotl(v305, v306) ^ v311
	v315 = int64(13)
	v317 = v308 ^ base.I64_rotl(v178, v315)
	v318 = v304 + v317
	v323 = base.I64_rotl(v318, v309) ^ int64(255) + v312
	v324 = base.I64_rotl(v312, v302) ^ v323
	v328 = int64(17)
	v330 = v318 ^ base.I64_rotl(v317, v328)
	v331 = v311 ^ v300 + v330
	v334 = base.I64_rotl(v331, v309) + v324
	v335 = base.I64_rotl(v324, v306) ^ v334
	v340 = v331 ^ base.I64_rotl(v330, v315)
	v341 = v340 + v323
	v344 = base.I64_rotl(v341, v309) + v335
	v350 = base.I64_rotl(v340, v328) ^ v341
	v354 = base.I64_rotl(v350, v315) ^ (v350 + v334)
	v358 = v354 + v344
	return base.I64_rotl(base.I64_rotl(v335, v302)^v344, v306) ^ base.I64_rotl(v354, v328) ^ base.I64_rotl(v358, v309) ^ v358
}
func F_dictGenHashFunction(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v151 int64
	_ = v151
	var v157 int64
	_ = v157
	var v161 int64
	_ = v161
	var v165 int64
	_ = v165
	v3 = int32(_a_F_dictGenHashFunction_0)
	v11 = *(*int64)(unsafe.Add(mBase, _c_F_dictGenHashFunction[0]))
	v13 = v11 ^ int64(8317987319222330741)
	v14 = *(*int64)(unsafe.Add(mBase, _c_F_dictGenHashFunction[1]))
	v16 = v14 ^ int64(7237128888997146477)
	v18 = v11 ^ int64(7816392313619706465)
	v20 = v14 ^ int64(8387220255154660723)
	v24 = l1 & int32(7)
	v25 = l0 + l1 - v24
	if l0 == v25 {
		v63 = l0
		v66 = v18
		v67 = v13
		v68 = v20
		v69 = v16
	} else {
		v27 = l0
		v30 = v18
		v31 = v13
		v32 = v20
		v33 = v16
		for {
			v37 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
			v38 = v37 ^ v32
			v39 = v38 + v30
			v40 = v31 + v33
			v43 = v40 ^ base.I64_rotl(v33, int64(13))
			v44 = v39 + v43
			v47 = v44 ^ base.I64_rotl(v43, int64(17))
			v50 = base.I64_rotl(v38, int64(16)) ^ v39
			v53 = int64(32)
			v55 = v50 + base.I64_rotl(v40, v53)
			v56 = base.I64_rotl(v50, int64(21)) ^ v55
			v58 = base.I64_rotl(v44, v53)
			v59 = v55 ^ v37
			v61 = v27 + int32(8)
			if v61 != v25 {
				v27 = v61
				v30 = v58
				v31 = v59
				v32 = v56
				v33 = v47
				continue
			} else {
				break
			}
			break
		}
		v63 = v25
		v66 = v58
		v67 = v59
		v68 = v56
		v69 = v47
	}
	v74 = base.I64_extend_i32_u(l1) << (uint(int64(56)) % 64)
	switch v24 {
	default:
		v107 = v74
	case 1:
		v104 = v74
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 2:
		v99 = v74
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 3:
		v94 = v74
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 4:
		v89 = v74
		v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
		v94 = v90<<(uint(int64(24))%64) | v89
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 5:
		v84 = v74
		v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
		v89 = v85<<(uint(int64(32))%64) | v84
		v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
		v94 = v90<<(uint(int64(24))%64) | v89
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 6:
		v79 = v74
		v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
		v84 = v80<<(uint(int64(40))%64) | v79
		v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
		v89 = v85<<(uint(int64(32))%64) | v84
		v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
		v94 = v90<<(uint(int64(24))%64) | v89
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 7:
		v75 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+6)))
		v79 = v75<<(uint(int64(48))%64) | v74
		v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
		v84 = v80<<(uint(int64(40))%64) | v79
		v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
		v89 = v85<<(uint(int64(32))%64) | v84
		v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
		v94 = v90<<(uint(int64(24))%64) | v89
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	}
	v108 = v107 ^ v68
	v109 = int64(16)
	v111 = v108 + v66
	v112 = base.I64_rotl(v108, v109) ^ v111
	v113 = int64(21)
	v115 = v67 + v69
	v116 = int64(32)
	v118 = v112 + base.I64_rotl(v115, v116)
	v119 = base.I64_rotl(v112, v113) ^ v118
	v122 = int64(13)
	v124 = v115 ^ base.I64_rotl(v69, v122)
	v125 = v111 + v124
	v130 = base.I64_rotl(v125, v116) ^ int64(255) + v119
	v131 = base.I64_rotl(v119, v109) ^ v130
	v135 = int64(17)
	v137 = v125 ^ base.I64_rotl(v124, v135)
	v138 = v118 ^ v107 + v137
	v141 = base.I64_rotl(v138, v116) + v131
	v142 = base.I64_rotl(v131, v113) ^ v141
	v147 = v138 ^ base.I64_rotl(v137, v122)
	v148 = v147 + v130
	v151 = base.I64_rotl(v148, v116) + v142
	v157 = base.I64_rotl(v147, v135) ^ v148
	v161 = base.I64_rotl(v157, v122) ^ (v157 + v141)
	v165 = v161 + v151
	return base.I64_rotl(base.I64_rotl(v142, v109)^v151, v113) ^ base.I64_rotl(v161, v135) ^ base.I64_rotl(v165, v116) ^ v165
}
func F_dictGenericDelete(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var __phi90 int32
	_ = __phi90
	var v96 int32
	_ = v96
	var __phi96 int32
	_ = __phi96
	var v97 int32
	_ = v97
	var __phi97 int32
	_ = __phi97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int64
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var __phi170 int32
	_ = __phi170
	var v172 int32
	_ = v172
	var __phi172 int32
	_ = __phi172
	var v174 int32
	_ = v174
	var __phi174 int32
	_ = __phi174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	v4 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v15 == v4-v17 {
		v195 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v211)+16))
	if v213 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L2:
	;
	return v195
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = m.T0[v21].(func(*base.Module, int32) int64)(m, l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v27 = int32(-1)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v28 == int32(255) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v35 = int64(0)
	goto L8
L7:
	;
	v35 = base.I64_extend_i32_u(v27<<(uint(v28)%32) ^ v27)
	goto L8
L8:
	;
	v36 = v22 & v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v37 == int32(-1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.I32_wrap_i64(v36) < v56 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v40 = base.I32_wrap_i64(v36)
	if v40 < v37 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
	if v51 != 0 {
		goto L9
	} else {
		goto L15
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v40<<(uint(int32(2))%32))))
	if v46 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	F_dictBucketRehash(m, l0, v36)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	v53 = F_dictRehash(m, l0, int32(1))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v144 = int32(-1)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v145 == int32(255) {
		goto L35
	} else {
		goto L36
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = int32(-1)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v62 == int32(255) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v126 == int32(-1) {
		v195 = v4
		goto L2
	} else {
		goto L34
	}
L20:
	;
	v69 = int64(0)
	goto L22
L21:
	;
	v69 = base.I64_extend_i32_u(v61<<(uint(v62)%32) ^ v61)
	goto L22
L22:
	;
	v71 = base.I32_wrap_i64(v69 & v22)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v59+v71<<(uint(int32(2))%32))))
	if v75 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v79 = l0 + int32(4)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if l1 != v81 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	__phi90 = v75
	__phi96 = v81
	__phi97 = int32(0)
	v90 = __phi90
	v96 = __phi96
	v97 = __phi97
	goto L26
L25:
	;
	v211 = v75
	v213 = int32(0)
	v214 = v71
	v215 = v79
	v216 = int32(0)
	goto L1
L26:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	if v99 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v211 = v107
	v213 = v90
	v214 = v71
	v215 = v79
	v216 = int32(0)
	goto L1
L28:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	if v107 == int32(0) {
		goto L19
	} else {
		goto L32
	}
L29:
	;
	v102 = m.T0[v99].(func(*base.Module, int32, int32) int32)(m, l1, v96)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	if v102 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v211 = v90
	v213 = v97
	v214 = v71
	v215 = v79
	v216 = int32(0)
	goto L1
L32:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if l1 != v110 {
		__phi90 = v107
		__phi96 = v110
		__phi97 = v90
		v90 = __phi90
		v96 = __phi96
		v97 = __phi97
		goto L26
	} else {
		goto L33
	}
L33:
	;
	goto L27
L34:
	;
	goto L17
L35:
	;
	v152 = int64(0)
	goto L37
L36:
	;
	v152 = base.I64_extend_i32_u(v144<<(uint(v145)%32) ^ v144)
	goto L37
L37:
	;
	v154 = base.I32_wrap_i64(v152 & v22)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v142+v154<<(uint(int32(2))%32))))
	if v158 == int32(0) {
		v195 = v4
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v162 = l0 + int32(8)
	v163 = int32(0)
	v164 = int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if l1 == v165 {
		v211 = v158
		v213 = v163
		v214 = v154
		v215 = v162
		v216 = v164
		goto L1
	} else {
		goto L39
	}
L39:
	;
	__phi170 = v165
	__phi172 = v158
	__phi174 = v163
	v170 = __phi170
	v172 = __phi172
	v174 = __phi174
	goto L40
L40:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	if v181 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v211 = v188
	v213 = v172
	v214 = v154
	v215 = v162
	v216 = v164
	goto L1
L42:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v172)+16))
	if v188 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v184 = m.T0[v181].(func(*base.Module, int32, int32) int32)(m, l1, v170)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	if v184 == int32(0) {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v211 = v172
	v213 = v174
	v214 = v154
	v215 = v162
	v216 = v164
	goto L1
L46:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	if l1 != v190 {
		__phi170 = v190
		__phi172 = v188
		__phi174 = v172
		v170 = __phi170
		v172 = __phi172
		v174 = __phi174
		goto L40
	} else {
		goto L48
	}
L47:
	;
	v195 = int32(0)
	goto L2
L48:
	;
	goto L41
L49:
	;
	if l2 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	*(*int32)(unsafe.Add(mBase, uint32(v223+v214<<(uint(int32(2))%32)))) = v219
	goto L49
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+16)) = v219
	goto L49
L52:
	;
	v251 = l0 + int32(12) + v216<<(uint(int32(2))%32)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v252 + int32(-1)
	v256 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	if v256 <= int32(0) {
		goto L61
	} else {
		goto L62
	}
L53:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	if v231 == int32(0) {
		v238 = v230
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	if v239 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	m.T0[v231].(func(*base.Module, int32))(m, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v238 = v237
	goto L54
L57:
	;
	F_valkey_free(m, v211)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	m.T0[v239].(func(*base.Module, int32))(m, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	goto L52
L61:
	;
	v260 = F_dictShrinkIfNeeded(m, l0)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L63
	}
L62:
	;
	return v211
L63:
	;
	return v211
}
func F_dictGetIterator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+12)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+4)) = int64(4294967295)
		return v4
	}
}
func F_dictGetKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return v2
}
func F_dictGetRandomKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v7 == v2-v9 {
		v294 = v2
		return v294
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v12 == int32(-1) {
			v121 = int32(-1)
			v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
			if v122 == int32(255) {
				v128 = int32(0)
			} else {
				v128 = v121<<(uint(v122)%32) ^ v121
			}
			for {
				v134 = int32(0)
				F___lock(m, int32(9116960))
				mBase = m.M
				v141 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[0]))
				v143 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[1]))
				if v143 != 0 {
					v147 = int32(0)
					v148 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2]))
					v149 = int32(2)
					v151 = v141 + v148<<(uint(v149)%32)
					v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
					v154 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3]))
					v158 = *(*int32)(unsafe.Add(mBase, uint32(v141+v154<<(uint(v149)%32))))
					v159 = v152 + v158
					*(*int32)(unsafe.Add(mBase, uint32(v151))) = v159
					v164 = v154 + int32(1)
					if v164 == v143 {
						v166 = v147
					} else {
						v166 = v164
					}
					*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3])) = v166
					v168 = int32(0)
					v171 = v148 + int32(1)
					if v171 == v143 {
						v173 = v168
					} else {
						v173 = v171
					}
					*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2])) = v173
					v178 = int32(base.Ui32(v159) >> (uint(int32(1)) % 32))
				} else {
					v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
					v145 = F_lcg31(m, v144)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v141))) = v145
					v178 = v145
				}
				F___unlock(m, int32(9116960))
				mBase = m.M
				v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v188 = *(*int32)(unsafe.Add(mBase, uint32(v183+v178&v128<<(uint(int32(2))%32))))
				if v188 == int32(0) {
					continue
				} else {
					break
				}
				break
			}
			v192 = v188
			v197 = int32(0)
			v199 = v192
			for {
				v203 = v197 + int32(1)
				v204 = *(*int32)(unsafe.Add(mBase, uint32(v199)+16))
				if v204 != 0 {
					v197 = v203
					v199 = v204
					continue
				} else {
					break
				}
				break
			}
			v205 = int32(0)
			F___lock(m, int32(9116960))
			mBase = m.M
			v212 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[0]))
			v214 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[1]))
			if v214 != 0 {
				v218 = int32(0)
				v219 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2]))
				v220 = int32(2)
				v222 = v212 + v219<<(uint(v220)%32)
				v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
				v225 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3]))
				v229 = *(*int32)(unsafe.Add(mBase, uint32(v212+v225<<(uint(v220)%32))))
				v230 = v223 + v229
				*(*int32)(unsafe.Add(mBase, uint32(v222))) = v230
				v235 = v225 + int32(1)
				if v235 == v214 {
					v237 = v218
				} else {
					v237 = v235
				}
				*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3])) = v237
				v239 = int32(0)
				v242 = v219 + int32(1)
				if v242 == v214 {
					v244 = v239
				} else {
					v244 = v242
				}
				*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2])) = v244
				v249 = int32(base.Ui32(v230) >> (uint(int32(1)) % 32))
			} else {
				v215 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
				v216 = F_lcg31(m, v215)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v212))) = v216
				v249 = v216
			}
			F___unlock(m, int32(9116960))
			mBase = m.M
			v254 = base.I32_rem_s(v249, v203)
			if v254 == int32(0) {
				v294 = v192
			} else {
				v258 = v254 & int32(7)
				if v258 != 0 {
					v260 = v254
					v261 = v192
					v262 = int32(0)
					for {
						v266 = v260 + int32(-1)
						v267 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
						v269 = v262 + int32(1)
						if v269 != v258 {
							v260 = v266
							v261 = v267
							v262 = v269
							continue
						} else {
							break
						}
						break
					}
					v271 = v266
					v272 = v267
				} else {
					v271 = v254
					v272 = v192
				}
				if base.Ui32(v254) < base.Ui32(int32(8)) {
					v294 = v272
				} else {
					v278 = v271
					v279 = v272
					for {
						v283 = *(*int32)(unsafe.Add(mBase, uint32(v279)+16))
						v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+16))
						v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
						v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+16))
						v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+16))
						v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+16))
						v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+16))
						v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+16))
						v292 = v278 + int32(-8)
						if v292 != 0 {
							v278 = v292
							v279 = v290
							continue
						} else {
							break
						}
						break
					}
					v294 = v290
				}
			}
			return v294
		} else {
			v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
			if v15 != 0 {
				v24 = v12
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
				if v27 == int32(255) {
					v31 = int32(0)
				} else {
					v31 = int32(1) << (uint(v27) % 32)
				}
				v34 = v24
				for {
					v37 = int32(0)
					F___lock(m, int32(9116960))
					mBase = m.M
					v44 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[0]))
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[1]))
					if v46 != 0 {
						v50 = int32(0)
						v51 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2]))
						v52 = int32(2)
						v54 = v44 + v51<<(uint(v52)%32)
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
						v57 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3]))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v44+v57<<(uint(v52)%32))))
						v62 = v55 + v61
						*(*int32)(unsafe.Add(mBase, uint32(v54))) = v62
						v67 = v57 + int32(1)
						if v67 == v46 {
							v69 = v50
						} else {
							v69 = v67
						}
						*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3])) = v69
						v71 = int32(0)
						v74 = v51 + int32(1)
						if v74 == v46 {
							v76 = v71
						} else {
							v76 = v74
						}
						*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2])) = v76
						v81 = int32(base.Ui32(v62) >> (uint(int32(1)) % 32))
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
						v48 = F_lcg31(m, v47)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v44))) = v48
						v81 = v48
					}
					F___unlock(m, int32(9116960))
					mBase = m.M
					v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
					if v88 == int32(255) {
						v92 = int32(0)
					} else {
						v92 = int32(1) << (uint(v88) % 32)
					}
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
					if v97 == int32(255) {
						v101 = int32(0)
					} else {
						v101 = int32(1) << (uint(v97) % 32)
					}
					v103 = base.I32_rem_u_s(v81, v92-v93+v101)
					v104 = v103 + v34
					if base.Ui32(v104) < base.Ui32(v31) {
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v115 = v111 + v104<<(uint(int32(2))%32)
					} else {
						v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v115 = v106 + (v104-v31)<<(uint(int32(2))%32)
					}
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
					if v116 == int32(0) {
						v34 = v93
						continue
					} else {
						break
					}
					break
				}
				v192 = v116
				v197 = int32(0)
				v199 = v192
				for {
					v203 = v197 + int32(1)
					v204 = *(*int32)(unsafe.Add(mBase, uint32(v199)+16))
					if v204 != 0 {
						v197 = v203
						v199 = v204
						continue
					} else {
						break
					}
					break
				}
				v205 = int32(0)
				F___lock(m, int32(9116960))
				mBase = m.M
				v212 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[0]))
				v214 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[1]))
				if v214 != 0 {
					v218 = int32(0)
					v219 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2]))
					v220 = int32(2)
					v222 = v212 + v219<<(uint(v220)%32)
					v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
					v225 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3]))
					v229 = *(*int32)(unsafe.Add(mBase, uint32(v212+v225<<(uint(v220)%32))))
					v230 = v223 + v229
					*(*int32)(unsafe.Add(mBase, uint32(v222))) = v230
					v235 = v225 + int32(1)
					if v235 == v214 {
						v237 = v218
					} else {
						v237 = v235
					}
					*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3])) = v237
					v239 = int32(0)
					v242 = v219 + int32(1)
					if v242 == v214 {
						v244 = v239
					} else {
						v244 = v242
					}
					*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2])) = v244
					v249 = int32(base.Ui32(v230) >> (uint(int32(1)) % 32))
				} else {
					v215 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
					v216 = F_lcg31(m, v215)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v212))) = v216
					v249 = v216
				}
				F___unlock(m, int32(9116960))
				mBase = m.M
				v254 = base.I32_rem_s(v249, v203)
				if v254 == int32(0) {
					v294 = v192
				} else {
					v258 = v254 & int32(7)
					if v258 != 0 {
						v260 = v254
						v261 = v192
						v262 = int32(0)
						for {
							v266 = v260 + int32(-1)
							v267 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
							v269 = v262 + int32(1)
							if v269 != v258 {
								v260 = v266
								v261 = v267
								v262 = v269
								continue
							} else {
								break
							}
							break
						}
						v271 = v266
						v272 = v267
					} else {
						v271 = v254
						v272 = v192
					}
					if base.Ui32(v254) < base.Ui32(int32(8)) {
						v294 = v272
					} else {
						v278 = v271
						v279 = v272
						for {
							v283 = *(*int32)(unsafe.Add(mBase, uint32(v279)+16))
							v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+16))
							v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
							v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+16))
							v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+16))
							v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+16))
							v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+16))
							v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+16))
							v292 = v278 + int32(-8)
							if v292 != 0 {
								v278 = v292
								v279 = v290
								continue
							} else {
								break
							}
							break
						}
						v294 = v290
					}
				}
				return v294
			} else {
				v17 = F_dictRehash(m, l0, int32(1))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v21 == int32(-1) {
						v121 = int32(-1)
						v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
						if v122 == int32(255) {
							v128 = int32(0)
						} else {
							v128 = v121<<(uint(v122)%32) ^ v121
						}
						for {
							v134 = int32(0)
							F___lock(m, int32(9116960))
							mBase = m.M
							v141 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[0]))
							v143 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[1]))
							if v143 != 0 {
								v147 = int32(0)
								v148 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2]))
								v149 = int32(2)
								v151 = v141 + v148<<(uint(v149)%32)
								v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
								v154 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3]))
								v158 = *(*int32)(unsafe.Add(mBase, uint32(v141+v154<<(uint(v149)%32))))
								v159 = v152 + v158
								*(*int32)(unsafe.Add(mBase, uint32(v151))) = v159
								v164 = v154 + int32(1)
								if v164 == v143 {
									v166 = v147
								} else {
									v166 = v164
								}
								*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3])) = v166
								v168 = int32(0)
								v171 = v148 + int32(1)
								if v171 == v143 {
									v173 = v168
								} else {
									v173 = v171
								}
								*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2])) = v173
								v178 = int32(base.Ui32(v159) >> (uint(int32(1)) % 32))
							} else {
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
								v145 = F_lcg31(m, v144)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v141))) = v145
								v178 = v145
							}
							F___unlock(m, int32(9116960))
							mBase = m.M
							v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v188 = *(*int32)(unsafe.Add(mBase, uint32(v183+v178&v128<<(uint(int32(2))%32))))
							if v188 == int32(0) {
								continue
							} else {
								break
							}
							break
						}
						v192 = v188
					} else {
						v24 = v21
						v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
						if v27 == int32(255) {
							v31 = int32(0)
						} else {
							v31 = int32(1) << (uint(v27) % 32)
						}
						v34 = v24
						for {
							v37 = int32(0)
							F___lock(m, int32(9116960))
							mBase = m.M
							v44 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[0]))
							v46 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[1]))
							if v46 != 0 {
								v50 = int32(0)
								v51 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2]))
								v52 = int32(2)
								v54 = v44 + v51<<(uint(v52)%32)
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
								v57 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3]))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v44+v57<<(uint(v52)%32))))
								v62 = v55 + v61
								*(*int32)(unsafe.Add(mBase, uint32(v54))) = v62
								v67 = v57 + int32(1)
								if v67 == v46 {
									v69 = v50
								} else {
									v69 = v67
								}
								*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3])) = v69
								v71 = int32(0)
								v74 = v51 + int32(1)
								if v74 == v46 {
									v76 = v71
								} else {
									v76 = v74
								}
								*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2])) = v76
								v81 = int32(base.Ui32(v62) >> (uint(int32(1)) % 32))
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
								v48 = F_lcg31(m, v47)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v44))) = v48
								v81 = v48
							}
							F___unlock(m, int32(9116960))
							mBase = m.M
							v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
							if v88 == int32(255) {
								v92 = int32(0)
							} else {
								v92 = int32(1) << (uint(v88) % 32)
							}
							v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
							if v97 == int32(255) {
								v101 = int32(0)
							} else {
								v101 = int32(1) << (uint(v97) % 32)
							}
							v103 = base.I32_rem_u_s(v81, v92-v93+v101)
							v104 = v103 + v34
							if base.Ui32(v104) < base.Ui32(v31) {
								v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v115 = v111 + v104<<(uint(int32(2))%32)
							} else {
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v115 = v106 + (v104-v31)<<(uint(int32(2))%32)
							}
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
							if v116 == int32(0) {
								v34 = v93
								continue
							} else {
								break
							}
							break
						}
						v192 = v116
					}
					v197 = int32(0)
					v199 = v192
					for {
						v203 = v197 + int32(1)
						v204 = *(*int32)(unsafe.Add(mBase, uint32(v199)+16))
						if v204 != 0 {
							v197 = v203
							v199 = v204
							continue
						} else {
							break
						}
						break
					}
					v205 = int32(0)
					F___lock(m, int32(9116960))
					mBase = m.M
					v212 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[0]))
					v214 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[1]))
					if v214 != 0 {
						v218 = int32(0)
						v219 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2]))
						v220 = int32(2)
						v222 = v212 + v219<<(uint(v220)%32)
						v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
						v225 = *(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3]))
						v229 = *(*int32)(unsafe.Add(mBase, uint32(v212+v225<<(uint(v220)%32))))
						v230 = v223 + v229
						*(*int32)(unsafe.Add(mBase, uint32(v222))) = v230
						v235 = v225 + int32(1)
						if v235 == v214 {
							v237 = v218
						} else {
							v237 = v235
						}
						*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[3])) = v237
						v239 = int32(0)
						v242 = v219 + int32(1)
						if v242 == v214 {
							v244 = v239
						} else {
							v244 = v242
						}
						*(*int32)(unsafe.Add(mBase, _c_F_dictGetRandomKey[2])) = v244
						v249 = int32(base.Ui32(v230) >> (uint(int32(1)) % 32))
					} else {
						v215 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
						v216 = F_lcg31(m, v215)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v212))) = v216
						v249 = v216
					}
					F___unlock(m, int32(9116960))
					mBase = m.M
					v254 = base.I32_rem_s(v249, v203)
					if v254 == int32(0) {
						v294 = v192
					} else {
						v258 = v254 & int32(7)
						if v258 != 0 {
							v260 = v254
							v261 = v192
							v262 = int32(0)
							for {
								v266 = v260 + int32(-1)
								v267 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
								v269 = v262 + int32(1)
								if v269 != v258 {
									v260 = v266
									v261 = v267
									v262 = v269
									continue
								} else {
									break
								}
								break
							}
							v271 = v266
							v272 = v267
						} else {
							v271 = v254
							v272 = v192
						}
						if base.Ui32(v254) < base.Ui32(int32(8)) {
							v294 = v272
						} else {
							v278 = v271
							v279 = v272
							for {
								v283 = *(*int32)(unsafe.Add(mBase, uint32(v279)+16))
								v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+16))
								v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
								v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+16))
								v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+16))
								v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+16))
								v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+16))
								v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+16))
								v292 = v278 + int32(-8)
								if v292 != 0 {
									v278 = v292
									v279 = v290
									continue
								} else {
									break
								}
								break
							}
							v294 = v290
						}
					}
					return v294
				}
			}
		}
	}
}
func F_dictInitIterator(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(4294967295)
	return
}
func F_dictInitSafeIterator(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(1)
	return
}
func F_dictObjectDestructor(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	if l0 == int32(0) {
		return
	} else {
		F_decrRefCount(m, l0)
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			return
		}
	}
}
func F_dictPtrCompare(m *base.Module, l0 int32, l1 int32) int32 {
	return base.B2i32(l0 == l1)
}
func F_dictRehash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int64
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v208 int32
	_ = v208
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_dictRehash[0]))
	if v13 == int32(2) {
		v195 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_dictRehash_0), int32(_a_F_dictRehash_1), int32(343))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L36
	} else {
		goto L48
	}
L2:
	;
	return v195
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v16 == int32(-1) {
		v195 = v3
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v13 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l1 != 0 {
		goto L20
	} else {
		goto L21
	}
L6:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v23 == int32(255) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if base.Ui32(v34) <= base.Ui32(v27) {
		goto L5
	} else {
		goto L16
	}
L8:
	;
	v27 = int32(0)
	goto L10
L9:
	;
	v27 = int32(1) << (uint(v23) % 32)
	goto L10
L10:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v30 == int32(255) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v34 = int32(0)
	goto L13
L12:
	;
	v34 = int32(1) << (uint(v30) % 32)
	goto L13
L13:
	;
	if base.Ui32(v27) <= base.Ui32(v34) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(v27) < base.Ui32(v34<<(uint(int32(2))%32)) {
		v195 = v3
		goto L2
	} else {
		goto L15
	}
L15:
	;
	goto L7
L16:
	;
	if base.Ui32(v34) < base.Ui32(v27<<(uint(int32(5))%32)) {
		v195 = v3
		goto L2
	} else {
		goto L17
	}
L17:
	;
	goto L5
L18:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+28))
	if v172 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L19:
	;
	if v153 != 0 {
		v195 = int32(1)
		goto L2
	} else {
		goto L43
	}
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v50 = l1
	v52 = v48
	v53 = v16
	v55 = l1 * int32(10)
	goto L22
L21:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v153 = v45
	goto L19
L22:
	;
	if v52 == int32(0) {
		goto L18
	} else {
		goto L24
	}
L23:
	;
	v153 = v136
	goto L19
L24:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v61 == int32(255) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if int32(base.Ui32(v53)>>(uint(v61)%32)) != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v66 = v50 + int32(-1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v72 = v53
	v74 = v55
	goto L28
L27:
	;
	v91 = v81
	v92 = v52
	goto L32
L28:
	;
	v79 = v72 << (uint(int32(2)) % 32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v67+v79)))
	if v81 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v82 = int32(1)
	v84 = v72 + v82
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v84
	v87 = v74 + int32(-1)
	if v87 != 0 {
		v72 = v84
		v74 = v87
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v195 = v82
	goto L2
L32:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+27)))
	v101 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v100 <= v101 {
		v112 = v92
		v113 = v100
		v114 = base.I64_extend_i32_u(v72)
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v142+v79))) = int32(0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v148 = v146 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v148
	if v66 != 0 {
		v50 = v66
		v52 = v136
		v53 = v148
		v55 = v74
		goto L22
	} else {
		goto L42
	}
L34:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v117 = int32(-1)
	v118 = int32(255)
	v119 = v113 & v118
	if v119 == v118 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v106 = m.T0[v105].(func(*base.Module, int32) int64)(m, v103)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	return int32(0)
L37:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	v112 = v110
	v113 = v111
	v114 = v106
	goto L34
L38:
	;
	v126 = int64(0)
	goto L40
L39:
	;
	v126 = base.I64_extend_i32_u(v117<<(uint(v119)%32) ^ v117)
	goto L40
L40:
	;
	v131 = v115 + base.I32_wrap_i64(v126&v114)<<(uint(int32(2))%32)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+16)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v91
	v136 = v112 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v138 + int32(1)
	if v99 != 0 {
		v91 = v99
		v92 = v136
		goto L32
	} else {
		goto L41
	}
L41:
	;
	goto L33
L42:
	;
	goto L23
L43:
	;
	goto L18
L44:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_valkey_free(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L36
	} else {
		goto L47
	}
L45:
	;
	m.T0[v172].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L36
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v180
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v182
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	v185 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v185)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v184)
	v188 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v188
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
	v195 = v188
	goto L2
L48:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dictScriptDestructor(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	if l0 == int32(0) {
		return
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_scriptingEngineCallFreeFunction(m, v4, int32(0), v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_decrRefCount(m, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_valkey_free(m, l0)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_dictSdsCaseHash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v7 & int32(7) {
	case 0:
		v13 = F_siphash_nocase(m, l0, int32(base.Ui32(v7)>>(uint(int32(3))%32)), int32(_a_F_dictSdsCaseHash_0))
		mBase = m.M
		return v13
	case 1:
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v19 = F_siphash_nocase(m, l0, v17, int32(_a_F_dictSdsCaseHash_0))
		mBase = m.M
		return v19
	case 2:
		v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v25 = F_siphash_nocase(m, l0, v23, int32(_a_F_dictSdsCaseHash_0))
		mBase = m.M
		return v25
	case 3:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v31 = F_siphash_nocase(m, l0, v29, int32(_a_F_dictSdsCaseHash_0))
		mBase = m.M
		return v31
	case 4:
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v36 = v35
		v38 = F_siphash_nocase(m, l0, v36, int32(_a_F_dictSdsCaseHash_0))
		mBase = m.M
		return v38
	default:
		v36 = int32(0)
		v38 = F_siphash_nocase(m, l0, v36, int32(_a_F_dictSdsCaseHash_0))
		mBase = m.M
		return v38
	}
}
func F_dictSdsDestructor(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_sdsfree(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_dictSdsKeyCaseCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return base.B2i32(v37-v39 == int32(0))
L2:
	;
	v37 = F_tolower(m, v33)
	mBase = m.M
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v39 = F_tolower(m, v38)
	mBase = m.M
	goto L1
L3:
	;
	v7 = l0
	v8 = l1
	v9 = v5
	goto L6
L4:
	;
	v33 = int32(0)
	v34 = l1
	goto L2
L5:
	;
	v33 = v30 & int32(255)
	v34 = v29
	goto L2
L6:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v11 == int32(0) {
		v29 = v8
		v30 = v9
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v29 = v23
	v30 = int32(0)
	goto L5
L8:
	;
	v15 = v9 & int32(255)
	if v15 == v11 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v22 = int32(1)
	v23 = v8 + v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v24 != 0 {
		v7 = v7 + v22
		v8 = v23
		v9 = v24
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v17 = F_tolower(m, v15)
	mBase = m.M
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v19 = F_tolower(m, v18)
	mBase = m.M
	if v17 == v19 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v29 = v8
	v30 = v21
	goto L5
L12:
	;
	goto L7
}
func F_dictSdsKeyCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	v3 = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v10 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v27 = v3
		goto L1
	}
L1:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v30 & int32(7) {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	case 4:
		goto L8
	default:
		v47 = v3
		goto L7
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v27 = v26
	goto L1
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v27 = v23
	goto L1
L4:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v27 = v20
	goto L1
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v27 = v17
	goto L1
L6:
	;
	v27 = int32(base.Ui32(v10) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	if v27 != v47 {
		v116 = int32(0)
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v47 = v46
	goto L7
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v47 = v43
	goto L7
L10:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v47 = v40
	goto L7
L11:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v47 = v37
	goto L7
L12:
	;
	v47 = int32(base.Ui32(v30) >> (uint(int32(3)) % 32))
	goto L7
L13:
	;
	return v116
L14:
	;
	if base.Ui32(v27) < base.Ui32(int32(4)) {
		v73 = l0
		v74 = l1
		v75 = v27
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v116 = base.B2i32(v113 == int32(0))
	goto L13
L16:
	;
	v113 = int32(0)
	goto L15
L17:
	;
	v85 = v80
	v86 = v81
	v87 = v82
	goto L27
L18:
	;
	if v75 == int32(0) {
		goto L16
	} else {
		goto L25
	}
L19:
	;
	if (l1|l0)&int32(3) != 0 {
		v80 = l0
		v81 = l1
		v82 = v27
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v57 = l0
	v58 = l1
	v59 = v27
	goto L21
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v62 != v63 {
		v80 = v57
		v81 = v58
		v82 = v59
		goto L17
	} else {
		goto L23
	}
L22:
	;
	v73 = v68
	v74 = v66
	v75 = v70
	goto L18
L23:
	;
	v65 = int32(4)
	v66 = v58 + v65
	v68 = v57 + v65
	v70 = v59 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v70) {
		v57 = v68
		v58 = v66
		v59 = v70
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v80 = v73
	v81 = v74
	v82 = v75
	goto L17
L26:
	;
	v113 = v90 - v91
	goto L15
L27:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v90 != v91 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v93 = int32(1)
	v98 = v87 + int32(-1)
	if v98 == int32(0) {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v85 = v85 + v93
	v86 = v86 + v93
	v87 = v98
	goto L27
}
func F_dictSetResizeEnabled(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_dictSetResizeEnabled[0])) = l0
	return
}
func F_dictStringHash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	if l0&int32(3) == int32(0) {
		v23 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v58 = F_siphash(m, l0, v56, int32(_a_F_dictStringHash_0))
	mBase = m.M
	goto L17
L2:
	;
	v56 = v48 - l0
	goto L1
L3:
	;
	v27 = v23
	goto L11
L4:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v12 = l0
	goto L7
L6:
	;
	v56 = l0 - l0
	goto L1
L7:
	;
	v16 = v12 + int32(1)
	if v16&int32(3) == int32(0) {
		v23 = v16
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v21 != 0 {
		v12 = v16
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v48 = v16
	goto L2
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v36 = int32(-2139062144)
	if (int32(16843008)-v33|v33)&v36 == v36 {
		v27 = v27 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v42 = v27
	goto L14
L13:
	;
	goto L12
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v46 != 0 {
		v42 = v42 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v48 = v42
	goto L2
L16:
	;
	goto L15
L17:
	;
	return v58
}
func F_dictUnlink(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_dictGenericDelete(m, l0, l1, int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
