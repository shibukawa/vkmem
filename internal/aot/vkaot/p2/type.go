package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_setTypeConvert(m *base.Module, l0 int32, l1 int32) {
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = F_setTypeSize(m, l0)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v6 = F_setTypeConvertAndExpand(m, l0, l1, v3, int32(1))
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	}
}
func F_setTypeDup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11&int32(15) != int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a1685), int32(_a1683), int32(562))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L9
	} else {
		goto L44
	}
L2:
	;
	switch int32(base.Ui32(v11)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		goto L6
	default:
		goto L5
	case 4:
		goto L4
	case 9:
		goto L7
	}
L3:
	;
	m.G0 = v9 + int32(16)
	return v145
L4:
	;
	v122 = F_objectGetVal(m, l0)
	mBase = m.M
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v127 = v123*v124 + int32(8)
	goto L38
L5:
	;
	F__serverPanic_1(m, int32(_a1683), int32(592), int32(_a1684), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L9
	} else {
		goto L37
	}
L6:
	;
	v41 = F_createSetObject(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L15
	}
L7:
	;
	v23 = F_objectGetVal(m, l0)
	mBase = m.M
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	goto L8
L8:
	;
	v25 = F_valkey_malloc(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if v24 == int32(0) {
		v32 = v25
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v33 = F_createObject(m, int32(2), v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v31 = F__emscripten_memcpy_bulkmem(m, v25, v23, v24)
	mBase = m.M
	v32 = v31
	goto L12
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v35&int32(-241) | int32(176)
	v145 = v33
	goto L3
L15:
	;
	v43 = F_objectGetVal(m, l0)
	mBase = m.M
	v44 = F_objectGetVal(m, v41)
	mBase = m.M
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	goto L16
L16:
	;
	v48 = F_hashtableExpand(m, v44, v45+v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v50 = F_setTypeInitIterator(m, l0)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v106 != int32(2) {
		goto L33
	} else {
		goto L34
	}
L19:
	;
	v56 = F_setTypeNext(m, v50, v9+int32(12), v9+int32(8), v9)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if v56 == int32(-1) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L22
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+int32(-1)))))
	switch v70 & int32(7) {
	case 0:
		goto L29
	case 1:
		goto L28
	case 2:
		goto L27
	case 3:
		goto L26
	case 4:
		goto L25
	default:
		v87 = int32(0)
		goto L24
	}
L23:
	;
	goto L18
L24:
	;
	v90 = F_setTypeAddAux(m, v41, v67, v87, int64(0), int32(1))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L30
	}
L25:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(-17))))
	v87 = v86
	goto L24
L26:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(-9))))
	v87 = v83
	goto L24
L27:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+int32(-5)))))
	v87 = v80
	goto L24
L28:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+int32(-3)))))
	v87 = v77
	goto L24
L29:
	;
	v87 = int32(base.Ui32(v70) >> (uint(int32(3)) % 32))
	goto L24
L30:
	;
	v96 = F_setTypeNext(m, v50, v9+int32(12), v9+int32(8), v9)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	if v96 != int32(-1) {
		goto L22
	} else {
		goto L32
	}
L32:
	;
	goto L23
L33:
	;
	F_valkey_free(m, v50)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L36
	}
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	F_hashtableReleaseIterator(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v145 = v41
	goto L3
L37:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v128 = F_valkey_malloc(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	if v127 == int32(0) {
		v133 = v128
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v134 = F_createObject(m, int32(2), v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v132 = F__emscripten_memcpy_bulkmem(m, v128, v122, v127)
	mBase = m.M
	v133 = v132
	goto L41
L43:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v136&int32(-241) | int32(96)
	v145 = v134
	goto L3
L44:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_setTypeNextObject(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v11 = F_setTypeNext(m, l0, v5+int32(4), v5, v5+int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != int32(-1) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			if v18 == int32(0) {
				v24 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
				v25 = F_sdsfromlonglong(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = v25
					m.G0 = v5 + int32(16)
					return v27
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v22 = F_sdsnewlen(m, v18, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v27 = v22
					m.G0 = v5 + int32(16)
					return v27
				}
			}
		} else {
			v27 = int32(0)
			m.G0 = v5 + int32(16)
			return v27
		}
	}
}
func F_setTypeRandomElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v12)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		v19 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v19
		v22 = F_objectGetVal(m, l0)
		mBase = m.M
		v25 = F_hashtableFairRandomEntry(m, v22, v10+int32(12))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-1)))))
			switch v33 & int32(7) {
			case 0:
				v50 = int32(base.Ui32(v33) >> (uint(int32(3)) % 32))
			case 1:
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
				v50 = v40
			case 2:
				v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))))
				v50 = v43
			case 3:
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
				v50 = v46
			case 4:
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-17))))
				v50 = v49
			default:
				v50 = v19
			}
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v50
			*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(-123456789)
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			m.G0 = v10 + int32(16)
			return int32(base.Ui32(v95)>>(uint(int32(4))%32)) & int32(15)
		}
	default:
		F__serverPanic_1(m, int32(_a1683), int32(439), int32(_a1684), int32(0))
		mBase = m.M
		v84 = m.ExcPending
		if v84 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 4:
		v86 = F_objectGetVal(m, l0)
		mBase = m.M
		v87 = F_intsetRandom(m, v86)
		mBase = m.M
		v88 = m.ExcPending
		if v88 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l3))) = v87
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			m.G0 = v10 + int32(16)
			return int32(base.Ui32(v95)>>(uint(int32(4))%32)) & int32(15)
		}
	case 9:
		v54 = F_objectGetVal(m, l0)
		mBase = m.M
		v56 = int32(0)
		v58 = *(*int64)(unsafe.Add(mBase, _consts[408]))
		v62 = v58*int64(6364136223846793005) + int64(1)
		*(*int64)(unsafe.Add(mBase, _consts[408])) = v62
		v67 = F_lpLength(m, v54)
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return int32(0)
		} else {
			v69 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v62)>>(uint(int64(33))%64))), v67)
			v70 = F_lpSeek(m, v54, v69)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int32(0)
			} else {
				v74 = F_lpGetValue(m, v70, v10+int32(8), l3)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v74
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					m.G0 = v10 + int32(16)
					return int32(base.Ui32(v95)>>(uint(int32(4))%32)) & int32(15)
				}
			}
		}
	}
}
func F_setTypeReleaseIterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2 != int32(2) {
		F_valkey_free(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_hashtableReleaseIterator(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_valkey_free(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_setTypeRemoveAux(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v159 int64
	_ = v159
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v181 int64
	_ = v181
	var v186 int64
	_ = v186
	var v190 int32
	_ = v190
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v202 int64
	_ = v202
	var v226 int64
	_ = v226
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if l1 != 0 {
		v75 = l1
		v76 = l2
		v77 = l4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v278
L2:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v78)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		goto L22
	default:
		goto L19
	case 4:
		goto L20
	case 9:
		goto L21
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13&int32(240) != int32(96) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = v11 + int32(16)
	v30 = int32(0)
	if l3 <= int64(-1) {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v18 = F_objectGetVal(m, l0)
	mBase = m.M
	v21 = F_intsetRemove(m, v18, l3, v11+int32(16))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	F_objectSetVal(m, l0, v21)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v278 = v27
	goto L1
L9:
	;
	v75 = v29
	v76 = v74
	v77 = v30
	goto L2
L10:
	;
	v74 = v30
	goto L9
L12:
	;
	v55 = F_ull2string(m, v51, v52, v53)
	mBase = m.M
	if v55 == int32(0) {
		goto L10
	} else {
		goto L16
	}
L13:
	;
	goto L15
L14:
	;
	v51 = v29
	v52 = int32(21)
	v53 = l3
	v54 = int32(0)
	goto L12
L15:
	;
	v42 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v42)
	v51 = v11 + int32(17)
	v52 = int32(20)
	v53 = int64(0) - l3
	v54 = int32(1)
	goto L12
L16:
	;
	v74 = v55 + v54
	goto L9
L18:
	;
	v278 = int32(1)
	goto L1
L19:
	;
	F__serverPanic_1(m, int32(_a1683), int32(275), int32(_a1684), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L6
	} else {
		goto L68
	}
L20:
	;
	v110 = v11 + int32(8)
	v111 = int32(0)
	if base.Ui32(v76+int32(-21)) < base.Ui32(int32(-20)) {
		v245 = v111
		goto L38
	} else {
		goto L39
	}
L21:
	;
	v94 = F_objectGetVal(m, l0)
	mBase = m.M
	v95 = F_lpFirst(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L30
	}
L22:
	;
	if v77 != 0 {
		v87 = v75
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v88 = F_objectGetVal(m, l0)
	mBase = m.M
	v89 = F_hashtableDelete(m, v88, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	v85 = F_sdsnewlen(m, v75, v76)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v87 = v85
	goto L23
L26:
	;
	if v87 == v75 {
		v278 = v89
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_sdsfree(m, v87)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v278 = v89
	goto L1
L29:
	;
	v98 = int32(0)
	v100 = F_lpFind(m, v94, v95, v75, v76, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L6
	} else {
		goto L32
	}
L30:
	;
	if v95 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v278 = int32(0)
	goto L1
L32:
	;
	if v100 == int32(0) {
		v278 = v98
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v105 = F_lpDelete(m, v94, v100, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_objectSetVal(m, l0, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L18
L36:
	;
	v278 = int32(0)
	goto L1
L37:
	;
	if v245 == int32(0) {
		goto L36
	} else {
		goto L64
	}
L38:
	;
	goto L37
L39:
	;
	v123 = int32(1)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v76 != v123 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v245 = int32(1)
	goto L38
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v110))) = v226
	goto L40
L42:
	;
	if v124&int32(255) == int32(45) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v128 = v124 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v128&int32(255)) {
		v245 = v111
		goto L38
	} else {
		goto L44
	}
L44:
	;
	if v110 == int32(0) {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	v226 = base.I64_extend_i32_u(v128) & int64(255)
	goto L41
L46:
	;
	if base.Ui32(int32(8)) < base.Ui32((v147+int32(-49))&int32(255)) {
		v245 = v111
		goto L38
	} else {
		goto L49
	}
L47:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	v146 = int32(2)
	v147 = v144
	v148 = v75 + int32(1)
	goto L46
L48:
	;
	v146 = v123
	v147 = v124
	v148 = v75
	goto L46
L49:
	;
	v159 = base.I64_extend_i32_u(v147+int32(-48)) & int64(255)
	if base.Ui32(v76) <= base.Ui32(v146) {
		v202 = v159
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v124&int32(255) != int32(45) {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	v165 = v146
	v167 = v159
	v169 = v148
	goto L52
L52:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	if base.Ui32((v171+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v245 = v111
		goto L38
	} else {
		goto L54
	}
L53:
	;
	v202 = v192
	goto L50
L54:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v167) {
		v245 = v111
		goto L38
	} else {
		goto L55
	}
L55:
	;
	v181 = v167 * int64(10)
	v186 = base.I64_extend_i32_u(v171+int32(-48)) & int64(255)
	if base.Ui64(v186^int64(-1)) < base.Ui64(v181) {
		v245 = v111
		goto L38
	} else {
		goto L56
	}
L56:
	;
	v190 = int32(1)
	v192 = v181 + v186
	v194 = v165 + v190
	if v194 != v76 {
		v165 = v194
		v167 = v192
		v169 = v169 + v190
		goto L52
	} else {
		goto L57
	}
L57:
	;
	goto L53
L58:
	;
	if v202 < int64(0) {
		v245 = v111
		goto L38
	} else {
		goto L62
	}
L59:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v202) {
		v245 = v111
		goto L38
	} else {
		goto L60
	}
L60:
	;
	if v110 == int32(0) {
		goto L40
	} else {
		goto L61
	}
L61:
	;
	v226 = int64(0) - v202
	goto L41
L62:
	;
	if v110 == int32(0) {
		goto L40
	} else {
		goto L63
	}
L63:
	;
	v226 = v202
	goto L41
L64:
	;
	v254 = F_objectGetVal(m, l0)
	mBase = m.M
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v258 = F_intsetRemove(m, v254, v255, v11+int32(4))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	F_objectSetVal(m, l0, v258)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v262 != 0 {
		goto L18
	} else {
		goto L67
	}
L67:
	;
	goto L36
L68:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
