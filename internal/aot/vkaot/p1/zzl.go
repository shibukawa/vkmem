package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zzlDeleteRangeByLex(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
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
	var v48 int32
	_ = v48
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	v14 = F_zzlFirstInLexRange(m, l0, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v14
	if v14 == int32(0) {
		v48 = l0
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v9 + int32(16)
	return v48
L4:
	;
	v21 = l0
	v25 = v4
	v26 = v14
	goto L6
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v45
	v48 = v44
	goto L3
L6:
	;
	v27 = F_lpNext(m, v21, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v44 = v41
	v45 = v37
	goto L5
L8:
	;
	if v27 == int32(0) {
		v44 = v21
		v45 = v25
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v32 = F_zzlLexValueLteMax(m, v31, l1)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v32 == int32(0) {
		v44 = v21
		v45 = v25
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v37 = v25 + int32(1)
	v41 = F_lpDeleteRangeWithEntry(m, v21, v9+int32(12), int32(2))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v43 != 0 {
		v21 = v41
		v25 = v37
		v26 = v43
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L7
}
func F_zzlInsert(m *base.Module, l0 int32, l1 int32, l2 float64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int64
	_ = v78
	var v83 int64
	_ = v83
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v104 float64
	_ = v104
	var v108 int64
	_ = v108
	var v110 float64
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v22 = F_lpSeek(m, l0, v4)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_zzlInsert_0), int32(_a_F_zzlInsert_1), int32(1190))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L3
	} else {
		goto L70
	}
L2:
	;
	v270 = F_zzlInsertAt(m, l0, v259, l1, l2)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L3
	} else {
		goto L69
	}
L3:
	;
	return int32(0)
L4:
	;
	if v22 == int32(0) {
		v259 = v4
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v42 = v22
	goto L6
L6:
	;
	v53 = F_lpNext(m, l0, v42)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v259 = int32(0)
	goto L2
L8:
	;
	if v53 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v59 = F_lpGetValue(m, v53, v18+int32(32), v18)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	if base.F64_gt(v110, l2) != 0 {
		v259 = v42
		goto L2
	} else {
		goto L22
	}
L11:
	;
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v110 = base.F64_convert_i64_s(v108)
	goto L10
L12:
	;
	if v59 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v64 = int32(0)
	v68 = m.G0
	v70 = v68 - int32(32)
	m.G0 = v70
	v72 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v64
	v78 = *(*int64)(unsafe.Add(mBase, _c_F_zzlInsert[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v70+int32(8)))) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v70)+24)) = int64(0)
	v83 = *(*int64)(unsafe.Add(mBase, _c_F_zzlInsert[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v70))) = v83
	F_ffc_from_chars_double_options(m, v70+int32(16), v59, v59+v63, v70+int32(24), v70)
	mBase = m.M
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	if v91 == v64 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v110 = v104
	goto L10
L15:
	;
	goto L20
L16:
	;
	if v91 == int32(2) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v98 = int32(68)
	goto L19
L18:
	;
	v98 = int32(28)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v98
	goto L15
L20:
	;
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v70)+24))
	m.G0 = v70 + int32(32)
	goto L14
L22:
	;
	if base.F64_ne(v110, l2) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v252 = F_lpNext(m, l0, v53)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L3
	} else {
		goto L67
	}
L24:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v114 & int32(7) {
	case 0:
		goto L30
	case 1:
		goto L29
	case 2:
		goto L28
	case 3:
		goto L27
	case 4:
		goto L26
	default:
		v123 = int32(0)
		goto L25
	}
L25:
	;
	v128 = F_lpGetValue(m, v42, v18+int32(44), v18+int32(32))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L33
	}
L26:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v123 = v122
	goto L25
L27:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v123 = v121
	goto L25
L28:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v123 = v120
	goto L25
L29:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v123 = v119
	goto L25
L30:
	;
	v123 = int32(base.Ui32(v114) >> (uint(int32(3)) % 32))
	goto L25
L31:
	;
	if base.Ui32(v177) < base.Ui32(v123) {
		goto L44
	} else {
		goto L45
	}
L32:
	;
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	if v134 <= int64(-1) {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	if v128 == int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v177 = v132
	v178 = v128
	goto L31
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v175
	v177 = v175
	v178 = v18
	goto L31
L36:
	;
	v175 = int32(0)
	goto L35
L38:
	;
	v156 = F_ull2string(m, v152, v153, v154)
	mBase = m.M
	if v156 == int32(0) {
		goto L36
	} else {
		goto L42
	}
L39:
	;
	goto L41
L40:
	;
	v152 = v18
	v153 = int32(32)
	v154 = v134
	v155 = int32(0)
	goto L38
L41:
	;
	v143 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v143)
	v147 = int32(1)
	v152 = v18 + v147
	v153 = int32(31)
	v154 = int64(0) - v134
	v155 = v147
	goto L38
L42:
	;
	v175 = v156 + v155
	goto L35
L44:
	;
	v180 = v177
	goto L46
L45:
	;
	v180 = v123
	goto L46
L46:
	;
	if base.Ui32(v180) < base.Ui32(int32(4)) {
		v204 = v178
		v205 = l1
		v206 = v180
		goto L50
	} else {
		goto L51
	}
L47:
	;
	if v244 != 0 {
		goto L63
	} else {
		goto L64
	}
L48:
	;
	v244 = int32(0)
	goto L47
L49:
	;
	v216 = v211
	v217 = v212
	v218 = v213
	goto L59
L50:
	;
	if v206 == int32(0) {
		goto L48
	} else {
		goto L57
	}
L51:
	;
	if (l1|v178)&int32(3) != 0 {
		v211 = v178
		v212 = l1
		v213 = v180
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v188 = v178
	v189 = l1
	v190 = v180
	goto L53
L53:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v193 != v194 {
		v211 = v188
		v212 = v189
		v213 = v190
		goto L49
	} else {
		goto L55
	}
L54:
	;
	v204 = v199
	v205 = v197
	v206 = v201
	goto L50
L55:
	;
	v196 = int32(4)
	v197 = v189 + v196
	v199 = v188 + v196
	v201 = v190 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v201) {
		v188 = v199
		v189 = v197
		v190 = v201
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v211 = v204
	v212 = v205
	v213 = v206
	goto L49
L58:
	;
	v244 = v221 - v222
	goto L47
L59:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v221 != v222 {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v224 = int32(1)
	v229 = v218 + int32(-1)
	if v229 == int32(0) {
		goto L48
	} else {
		goto L62
	}
L62:
	;
	v216 = v216 + v224
	v217 = v217 + v224
	v218 = v229
	goto L59
L63:
	;
	v246 = v244
	goto L65
L64:
	;
	v246 = v177 - v123
	goto L65
L65:
	;
	if int32(0) < v246 {
		v259 = v42
		goto L2
	} else {
		goto L66
	}
L66:
	;
	goto L23
L67:
	;
	if v252 != 0 {
		v42 = v252
		goto L6
	} else {
		goto L68
	}
L68:
	;
	goto L7
L69:
	;
	m.G0 = v18 + int32(48)
	return v270
L70:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_zzlInsertAt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v22 float64
	_ = v22
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(160)
	m.G0 = v13
	v22 = base.F64_abs(l3)
	if base.F64_gt(v22, float64(4.611686018427388e+18)) != 0 {
		v36 = v5
	} else {
		if base.F64_lt(v22, float64(9.223372036854776e+18)) == int32(0) {
			v31 = int64(-9223372036854775807 - 1)
		} else {
			v29 = base.I64_trunc_f64_s(l3)
			v31 = v29
		}
		if base.F64_ne(l3, base.F64_convert_i64_s(v31)) != 0 {
			v36 = v5
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v13+int32(8)))) = v31
			v36 = int32(1)
		}
	}
	if v36 != 0 {
		v45 = v5
		v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
		v50 = v48 & int32(7)
		if l1 != 0 {
			switch v50 {
			case 0:
				v93 = int32(base.Ui32(v48) >> (uint(int32(3)) % 32))
			case 1:
				v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
				v93 = v83
			case 2:
				v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
				v93 = v86
			case 3:
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
				v93 = v89
			case 4:
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
				v93 = v92
			default:
				v93 = int32(0)
			}
			v97 = F_lpInsertString(m, l0, l2, v93, l1, int32(0), v13+int32(156))
			mBase = m.M
			v98 = m.ExcPending
			if v98 != 0 {
				return int32(0)
			} else {
				if v36 == int32(0) {
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v13)+156))
					v112 = F_lpInsertString(m, v97, v13+int32(16), v45, v109, int32(1), int32(0))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int32(0)
					} else {
						v115 = v112
						m.G0 = v13 + int32(160)
						return v115
					}
				} else {
					v101 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v13)+156))
					v105 = F_lpInsertInteger(m, v97, v101, v102, int32(1), int32(0))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						v115 = v105
						m.G0 = v13 + int32(160)
						return v115
					}
				}
			}
		} else {
			switch v50 {
			case 0:
				v66 = int32(base.Ui32(v48) >> (uint(int32(3)) % 32))
			case 1:
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
				v66 = v56
			case 2:
				v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
				v66 = v59
			case 3:
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
				v66 = v62
			case 4:
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
				v66 = v65
			default:
				v66 = int32(0)
			}
			v67 = F_lpAppend(m, l0, l2, v66)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				if v36 == int32(0) {
					v76 = F_lpAppend(m, v67, v13+int32(16), v45)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						v115 = v76
						m.G0 = v13 + int32(160)
						return v115
					}
				} else {
					v71 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
					v72 = F_lpAppendInteger(m, v67, v71)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v115 = v72
						m.G0 = v13 + int32(160)
						return v115
					}
				}
			}
		}
	} else {
		v41 = F_d2string(m, v13+int32(16), int32(128), l3)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			v45 = v41
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
			v50 = v48 & int32(7)
			if l1 != 0 {
				switch v50 {
				case 0:
					v93 = int32(base.Ui32(v48) >> (uint(int32(3)) % 32))
				case 1:
					v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
					v93 = v83
				case 2:
					v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
					v93 = v86
				case 3:
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
					v93 = v89
				case 4:
					v92 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
					v93 = v92
				default:
					v93 = int32(0)
				}
				v97 = F_lpInsertString(m, l0, l2, v93, l1, int32(0), v13+int32(156))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					if v36 == int32(0) {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v13)+156))
						v112 = F_lpInsertString(m, v97, v13+int32(16), v45, v109, int32(1), int32(0))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							v115 = v112
							m.G0 = v13 + int32(160)
							return v115
						}
					} else {
						v101 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
						v102 = *(*int32)(unsafe.Add(mBase, uint32(v13)+156))
						v105 = F_lpInsertInteger(m, v97, v101, v102, int32(1), int32(0))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							v115 = v105
							m.G0 = v13 + int32(160)
							return v115
						}
					}
				}
			} else {
				switch v50 {
				case 0:
					v66 = int32(base.Ui32(v48) >> (uint(int32(3)) % 32))
				case 1:
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
					v66 = v56
				case 2:
					v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
					v66 = v59
				case 3:
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
					v66 = v62
				case 4:
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
					v66 = v65
				default:
					v66 = int32(0)
				}
				v67 = F_lpAppend(m, l0, l2, v66)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					if v36 == int32(0) {
						v76 = F_lpAppend(m, v67, v13+int32(16), v45)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							v115 = v76
							m.G0 = v13 + int32(160)
							return v115
						}
					} else {
						v71 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
						v72 = F_lpAppendInteger(m, v67, v71)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							v115 = v72
							m.G0 = v13 + int32(160)
							return v115
						}
					}
				}
			}
		}
	}
}
func F_zzlIsInLexRange(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v113 int32
	_ = v113
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v8 == v9 {
		v78 = int32(0)
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v79 != 0 {
			v107 = v78
			return v107
		} else {
			v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if v80 != 0 {
				v107 = v78
				return v107
			} else {
				v85 = int32(0)
				v87 = F_lpSeek(m, l0, int32(-2))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					if v87 == int32(0) {
						v107 = v85
						return v107
					} else {
						v93 = F_zzlLexValueGteMin(m, v87, l1)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							if v93 == int32(0) {
								v107 = v85
								return v107
							} else {
								v98 = F_lpSeek(m, l0, int32(0))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									if v98 == int32(0) {
										F__serverAssert(m, int32(_a_F_zzlIsInLexRange_0), int32(_a_F_zzlIsInLexRange_1), int32(1076))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v102 = F_zzlLexValueLteMax(m, v98, l1)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int32(0)
										} else {
											v107 = v102
											return v107
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_zzlIsInLexRange[0]))
		if v8 == v12 {
			v85 = int32(0)
			v87 = F_lpSeek(m, l0, int32(-2))
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int32(0)
			} else {
				if v87 == int32(0) {
					v107 = v85
					return v107
				} else {
					v93 = F_zzlLexValueGteMin(m, v87, l1)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						if v93 == int32(0) {
							v107 = v85
							return v107
						} else {
							v98 = F_lpSeek(m, l0, int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								if v98 == int32(0) {
									F__serverAssert(m, int32(_a_F_zzlIsInLexRange_0), int32(_a_F_zzlIsInLexRange_1), int32(1076))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v102 = F_zzlLexValueLteMax(m, v98, l1)
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int32(0)
									} else {
										v107 = v102
										return v107
									}
								}
							}
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_zzlIsInLexRange[1]))
			if v9 == v15 {
				v85 = int32(0)
				v87 = F_lpSeek(m, l0, int32(-2))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					if v87 == int32(0) {
						v107 = v85
						return v107
					} else {
						v93 = F_zzlLexValueGteMin(m, v87, l1)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							if v93 == int32(0) {
								v107 = v85
								return v107
							} else {
								v98 = F_lpSeek(m, l0, int32(0))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									if v98 == int32(0) {
										F__serverAssert(m, int32(_a_F_zzlIsInLexRange_0), int32(_a_F_zzlIsInLexRange_1), int32(1076))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v102 = F_zzlLexValueLteMax(m, v98, l1)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int32(0)
										} else {
											v107 = v102
											return v107
										}
									}
								}
							}
						}
					}
				}
			} else {
				v17 = int32(0)
				if v9 == v12 {
					v107 = v17
					return v107
				} else {
					if v8 == v15 {
						v107 = v17
						return v107
					} else {
						v20 = int32(0)
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-1)))))
						switch v28 & int32(7) {
						case 0:
							v45 = int32(base.Ui32(v28) >> (uint(int32(3)) % 32))
						case 1:
							v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
							v45 = v35
						case 2:
							v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
							v45 = v38
						case 3:
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
							v45 = v41
						case 4:
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
							v45 = v44
						default:
							v45 = v20
						}
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-1)))))
						switch v48 & int32(7) {
						case 0:
							v65 = int32(base.Ui32(v48) >> (uint(int32(3)) % 32))
						case 1:
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-3)))))
							v65 = v55
						case 2:
							v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+int32(-5)))))
							v65 = v58
						case 3:
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-9))))
							v65 = v61
						case 4:
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-17))))
							v65 = v64
						default:
							v65 = v20
						}
						v66 = base.B2i32(base.Ui32(v45) < base.Ui32(v65))
						if base.Ui32(v45) < base.Ui32(v65) {
							v67 = v45
						} else {
							v67 = v65
						}
						v68 = F_memcmp(m, v8, v9, v67)
						mBase = m.M
						if v68 != 0 {
							v71 = v68
						} else {
							v71 = base.B2i32(base.Ui32(v65) < base.Ui32(v45)) - v66
						}
						if int32(0) < v71 {
							v107 = v20
							return v107
						} else {
							if v71 != 0 {
								v85 = int32(0)
								v87 = F_lpSeek(m, l0, int32(-2))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int32(0)
								} else {
									if v87 == int32(0) {
										v107 = v85
										return v107
									} else {
										v93 = F_zzlLexValueGteMin(m, v87, l1)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int32(0)
										} else {
											if v93 == int32(0) {
												v107 = v85
												return v107
											} else {
												v98 = F_lpSeek(m, l0, int32(0))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int32(0)
												} else {
													if v98 == int32(0) {
														F__serverAssert(m, int32(_a_F_zzlIsInLexRange_0), int32(_a_F_zzlIsInLexRange_1), int32(1076))
														mBase = m.M
														v113 = m.ExcPending
														if v113 != 0 {
															return int32(0)
														} else {
															F_abort(m)
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v102 = F_zzlLexValueLteMax(m, v98, l1)
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return int32(0)
														} else {
															v107 = v102
															return v107
														}
													}
												}
											}
										}
									}
								}
							} else {
								v78 = int32(0)
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								if v79 != 0 {
									v107 = v78
									return v107
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
									if v80 != 0 {
										v107 = v78
										return v107
									} else {
										v85 = int32(0)
										v87 = F_lpSeek(m, l0, int32(-2))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int32(0)
										} else {
											if v87 == int32(0) {
												v107 = v85
												return v107
											} else {
												v93 = F_zzlLexValueGteMin(m, v87, l1)
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													if v93 == int32(0) {
														v107 = v85
														return v107
													} else {
														v98 = F_lpSeek(m, l0, int32(0))
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															if v98 == int32(0) {
																F__serverAssert(m, int32(_a_F_zzlIsInLexRange_0), int32(_a_F_zzlIsInLexRange_1), int32(1076))
																mBase = m.M
																v113 = m.ExcPending
																if v113 != 0 {
																	return int32(0)
																} else {
																	F_abort(m)
																	mBase = m.M
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																v102 = F_zzlLexValueLteMax(m, v98, l1)
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return int32(0)
																} else {
																	v107 = v102
																	return v107
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_zzlIsInRange(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int64
	_ = v47
	var v52 int64
	_ = v52
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v73 float64
	_ = v73
	var v77 int64
	_ = v77
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int64
	_ = v113
	var v118 int64
	_ = v118
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v139 float64
	_ = v139
	var v143 int64
	_ = v143
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v14 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_gt(v13, v14) != 0 {
		v151 = v3
		m.G0 = v10 + int32(16)
		return v151
	} else {
		if base.F64_ne(v13, v14) != 0 {
			v20 = F_lpSeek(m, l0, int32(-1))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v20 == int32(0) {
					v151 = v3
					m.G0 = v10 + int32(16)
					return v151
				} else {
					v28 = F_lpGetValue(m, v20, v10+int32(12), v10)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						if v28 == int32(0) {
							v77 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
							v79 = base.F64_convert_i64_s(v77)
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v33 = int32(0)
							v37 = m.G0
							v39 = v37 - int32(32)
							m.G0 = v39
							v41 = F___errno_location(m)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v41))) = v33
							v47 = *(*int64)(unsafe.Add(mBase, _c_F_zzlIsInRange[0]))
							*(*int64)(unsafe.Add(mBase, uint32(v39+int32(8)))) = v47
							*(*int64)(unsafe.Add(mBase, uint32(v39)+24)) = int64(0)
							v52 = *(*int64)(unsafe.Add(mBase, _c_F_zzlIsInRange[1]))
							*(*int64)(unsafe.Add(mBase, uint32(v39))) = v52
							F_ffc_from_chars_double_options(m, v39+int32(16), v28, v28+v32, v39+int32(24), v39)
							mBase = m.M
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
							if v60 == v33 {
							} else {
								if v60 == int32(2) {
									v67 = int32(68)
								} else {
									v67 = int32(28)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v41))) = v67
							}
							v73 = *(*float64)(unsafe.Add(mBase, uint32(v39)+24))
							m.G0 = v39 + int32(32)
							v79 = v73
						}
						v80 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						if v83 != 0 {
							v84 = base.F64_gt(v79, v80)
						} else {
							v84 = base.F64_ge(v79, v80)
						}
						if v84 != int32(1) {
							v151 = v3
							m.G0 = v10 + int32(16)
							return v151
						} else {
							v88 = F_lpSeek(m, l0, int32(1))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								if v88 == int32(0) {
									F__serverAssert(m, int32(_a_F_zzlIsInRange_0), int32(_a_F_zzlIsInRange_1), int32(982))
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v94 = F_lpGetValue(m, v88, v10+int32(12), v10)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										if v94 == int32(0) {
											v143 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
											v145 = base.F64_convert_i64_s(v143)
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
											v99 = int32(0)
											v103 = m.G0
											v105 = v103 - int32(32)
											m.G0 = v105
											v107 = F___errno_location(m)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v107))) = v99
											v113 = *(*int64)(unsafe.Add(mBase, _c_F_zzlIsInRange[0]))
											*(*int64)(unsafe.Add(mBase, uint32(v105+int32(8)))) = v113
											*(*int64)(unsafe.Add(mBase, uint32(v105)+24)) = int64(0)
											v118 = *(*int64)(unsafe.Add(mBase, _c_F_zzlIsInRange[1]))
											*(*int64)(unsafe.Add(mBase, uint32(v105))) = v118
											F_ffc_from_chars_double_options(m, v105+int32(16), v94, v94+v98, v105+int32(24), v105)
											mBase = m.M
											v126 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
											if v126 == v99 {
											} else {
												if v126 == int32(2) {
													v133 = int32(68)
												} else {
													v133 = int32(28)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v107))) = v133
											}
											v139 = *(*float64)(unsafe.Add(mBase, uint32(v105)+24))
											m.G0 = v105 + int32(32)
											v145 = v139
										}
										v146 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
										v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
										if v149 != 0 {
											v150 = base.F64_lt(v145, v146)
										} else {
											v150 = base.F64_le(v145, v146)
										}
										v151 = v150
										m.G0 = v10 + int32(16)
										return v151
									}
								}
							}
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			if v17 != 0 {
				v151 = v3
				m.G0 = v10 + int32(16)
				return v151
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				if v18 != 0 {
					v151 = v3
					m.G0 = v10 + int32(16)
					return v151
				} else {
					v20 = F_lpSeek(m, l0, int32(-1))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						if v20 == int32(0) {
							v151 = v3
							m.G0 = v10 + int32(16)
							return v151
						} else {
							v28 = F_lpGetValue(m, v20, v10+int32(12), v10)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								if v28 == int32(0) {
									v77 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
									v79 = base.F64_convert_i64_s(v77)
								} else {
									v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
									v33 = int32(0)
									v37 = m.G0
									v39 = v37 - int32(32)
									m.G0 = v39
									v41 = F___errno_location(m)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v41))) = v33
									v47 = *(*int64)(unsafe.Add(mBase, _c_F_zzlIsInRange[0]))
									*(*int64)(unsafe.Add(mBase, uint32(v39+int32(8)))) = v47
									*(*int64)(unsafe.Add(mBase, uint32(v39)+24)) = int64(0)
									v52 = *(*int64)(unsafe.Add(mBase, _c_F_zzlIsInRange[1]))
									*(*int64)(unsafe.Add(mBase, uint32(v39))) = v52
									F_ffc_from_chars_double_options(m, v39+int32(16), v28, v28+v32, v39+int32(24), v39)
									mBase = m.M
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
									if v60 == v33 {
									} else {
										if v60 == int32(2) {
											v67 = int32(68)
										} else {
											v67 = int32(28)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v41))) = v67
									}
									v73 = *(*float64)(unsafe.Add(mBase, uint32(v39)+24))
									m.G0 = v39 + int32(32)
									v79 = v73
								}
								v80 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								if v83 != 0 {
									v84 = base.F64_gt(v79, v80)
								} else {
									v84 = base.F64_ge(v79, v80)
								}
								if v84 != int32(1) {
									v151 = v3
									m.G0 = v10 + int32(16)
									return v151
								} else {
									v88 = F_lpSeek(m, l0, int32(1))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										if v88 == int32(0) {
											F__serverAssert(m, int32(_a_F_zzlIsInRange_0), int32(_a_F_zzlIsInRange_1), int32(982))
											mBase = m.M
											v163 = m.ExcPending
											if v163 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v94 = F_lpGetValue(m, v88, v10+int32(12), v10)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												if v94 == int32(0) {
													v143 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
													v145 = base.F64_convert_i64_s(v143)
												} else {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
													v99 = int32(0)
													v103 = m.G0
													v105 = v103 - int32(32)
													m.G0 = v105
													v107 = F___errno_location(m)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v107))) = v99
													v113 = *(*int64)(unsafe.Add(mBase, _c_F_zzlIsInRange[0]))
													*(*int64)(unsafe.Add(mBase, uint32(v105+int32(8)))) = v113
													*(*int64)(unsafe.Add(mBase, uint32(v105)+24)) = int64(0)
													v118 = *(*int64)(unsafe.Add(mBase, _c_F_zzlIsInRange[1]))
													*(*int64)(unsafe.Add(mBase, uint32(v105))) = v118
													F_ffc_from_chars_double_options(m, v105+int32(16), v94, v94+v98, v105+int32(24), v105)
													mBase = m.M
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
													if v126 == v99 {
													} else {
														if v126 == int32(2) {
															v133 = int32(68)
														} else {
															v133 = int32(28)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v107))) = v133
													}
													v139 = *(*float64)(unsafe.Add(mBase, uint32(v105)+24))
													m.G0 = v105 + int32(32)
													v145 = v139
												}
												v146 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
												v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
												if v149 != 0 {
													v150 = base.F64_lt(v145, v146)
												} else {
													v150 = base.F64_le(v145, v146)
												}
												v151 = v150
												m.G0 = v10 + int32(16)
												return v151
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_zzlLastInLexRange(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v3 = int32(0)
	v7 = F_lpSeek(m, l0, int32(-2))
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = F_zzlIsInLexRange(m, l0, l1)
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	return v43
L4:
	;
	if v11 == int32(0) {
		v43 = v3
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v7 == int32(0) {
		v43 = v3
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v20 = v7
	goto L7
L7:
	;
	v21 = F_zzlLexValueLteMax(m, v20, l1)
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	F__serverAssert(m, int32(_a_F_zzlLastInLexRange_0), int32(_a_F_zzlLastInLexRange_1), int32(1125))
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L20
	}
L9:
	;
	v29 = F_lpPrev(m, l0, v20)
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L16
	}
L10:
	;
	if v21 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v26 = F_zzlLexValueGteMin(m, v20, l1)
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v26 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v28 = v20
	goto L15
L14:
	;
	v28 = int32(0)
	goto L15
L15:
	;
	v43 = v28
	goto L3
L16:
	;
	if v29 == int32(0) {
		v43 = v3
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v33 = F_lpPrev(m, l0, v29)
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v33 != 0 {
		v20 = v33
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L8
L20:
	;
	F_abort(m)
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_zzlLexValueLteMax(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a_F_zzlLexValueLteMax_0), int32(_a_F_zzlLexValueLteMax_1), int32(894))
		mBase = m.M
		v175 = m.ExcPending
		if v175 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v15 = F_lpGetValue(m, l0, v9+int32(12), v9)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				v24 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
				v25 = F_sdsfromlonglong(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = v25
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if v29 == int32(0) {
						v98 = int32(1)
						if v27 == v28 {
							v162 = v98
						} else {
							v101 = *(*int32)(unsafe.Add(mBase, _c_F_zzlLexValueLteMax[0]))
							if v27 == v101 {
								v162 = v98
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, _c_F_zzlLexValueLteMax[1]))
								if v28 == v104 {
									v162 = v98
								} else {
									v106 = int32(0)
									if v28 == v101 {
										v162 = v106
									} else {
										if v27 == v104 {
											v162 = v106
										} else {
											v109 = int32(0)
											v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-1)))))
											switch v116 & int32(7) {
											case 0:
												v133 = int32(base.Ui32(v116) >> (uint(int32(3)) % 32))
											case 1:
												v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-3)))))
												v133 = v123
											case 2:
												v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(-5)))))
												v133 = v126
											case 3:
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-9))))
												v133 = v129
											case 4:
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-17))))
												v133 = v132
											default:
												v133 = v109
											}
											v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-1)))))
											switch v136 & int32(7) {
											case 0:
												v153 = int32(base.Ui32(v136) >> (uint(int32(3)) % 32))
											case 1:
												v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))))
												v153 = v143
											case 2:
												v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))))
												v153 = v146
											case 3:
												v149 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9))))
												v153 = v149
											case 4:
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-17))))
												v153 = v152
											default:
												v153 = v109
											}
											v154 = base.B2i32(base.Ui32(v133) < base.Ui32(v153))
											if base.Ui32(v133) < base.Ui32(v153) {
												v155 = v133
											} else {
												v155 = v153
											}
											v156 = F_memcmp(m, v27, v28, v155)
											mBase = m.M
											if v156 != 0 {
												v159 = v156
											} else {
												v159 = base.B2i32(base.Ui32(v153) < base.Ui32(v133)) - v154
											}
											v162 = base.B2i32(v159 < int32(1))
										}
									}
								}
							}
						}
					} else {
						if v27 != v28 {
							v34 = int32(1)
							v36 = *(*int32)(unsafe.Add(mBase, _c_F_zzlLexValueLteMax[0]))
							if v27 == v36 {
								v162 = v34
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, _c_F_zzlLexValueLteMax[1]))
								if v28 == v39 {
									v162 = v34
								} else {
									if v28 != v36 {
										if v27 == v39 {
											v162 = int32(0)
										} else {
											v45 = int32(0)
											v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-1)))))
											switch v52 & int32(7) {
											case 0:
												v69 = int32(base.Ui32(v52) >> (uint(int32(3)) % 32))
											case 1:
												v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-3)))))
												v69 = v59
											case 2:
												v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(-5)))))
												v69 = v62
											case 3:
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-9))))
												v69 = v65
											case 4:
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-17))))
												v69 = v68
											default:
												v69 = v45
											}
											v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-1)))))
											switch v72 & int32(7) {
											case 0:
												v89 = int32(base.Ui32(v72) >> (uint(int32(3)) % 32))
											case 1:
												v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))))
												v89 = v79
											case 2:
												v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))))
												v89 = v82
											case 3:
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9))))
												v89 = v85
											case 4:
												v88 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-17))))
												v89 = v88
											default:
												v89 = v45
											}
											v90 = base.B2i32(base.Ui32(v69) < base.Ui32(v89))
											if base.Ui32(v69) < base.Ui32(v89) {
												v91 = v69
											} else {
												v91 = v89
											}
											v92 = F_memcmp(m, v27, v28, v91)
											mBase = m.M
											if v92 != 0 {
												v95 = v92
											} else {
												v95 = base.B2i32(base.Ui32(v89) < base.Ui32(v69)) - v90
											}
											v162 = int32(base.Ui32(v95) >> (uint(int32(31)) % 32))
										}
									} else {
										v162 = int32(0)
									}
								}
							}
						} else {
							v162 = int32(0)
						}
					}
					F_sdsfree(m, v27)
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v162
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				v22 = F_sdsnewlen(m, v15, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v27 = v22
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if v29 == int32(0) {
						v98 = int32(1)
						if v27 == v28 {
							v162 = v98
						} else {
							v101 = *(*int32)(unsafe.Add(mBase, _c_F_zzlLexValueLteMax[0]))
							if v27 == v101 {
								v162 = v98
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, _c_F_zzlLexValueLteMax[1]))
								if v28 == v104 {
									v162 = v98
								} else {
									v106 = int32(0)
									if v28 == v101 {
										v162 = v106
									} else {
										if v27 == v104 {
											v162 = v106
										} else {
											v109 = int32(0)
											v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-1)))))
											switch v116 & int32(7) {
											case 0:
												v133 = int32(base.Ui32(v116) >> (uint(int32(3)) % 32))
											case 1:
												v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-3)))))
												v133 = v123
											case 2:
												v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(-5)))))
												v133 = v126
											case 3:
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-9))))
												v133 = v129
											case 4:
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-17))))
												v133 = v132
											default:
												v133 = v109
											}
											v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-1)))))
											switch v136 & int32(7) {
											case 0:
												v153 = int32(base.Ui32(v136) >> (uint(int32(3)) % 32))
											case 1:
												v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))))
												v153 = v143
											case 2:
												v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))))
												v153 = v146
											case 3:
												v149 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9))))
												v153 = v149
											case 4:
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-17))))
												v153 = v152
											default:
												v153 = v109
											}
											v154 = base.B2i32(base.Ui32(v133) < base.Ui32(v153))
											if base.Ui32(v133) < base.Ui32(v153) {
												v155 = v133
											} else {
												v155 = v153
											}
											v156 = F_memcmp(m, v27, v28, v155)
											mBase = m.M
											if v156 != 0 {
												v159 = v156
											} else {
												v159 = base.B2i32(base.Ui32(v153) < base.Ui32(v133)) - v154
											}
											v162 = base.B2i32(v159 < int32(1))
										}
									}
								}
							}
						}
					} else {
						if v27 != v28 {
							v34 = int32(1)
							v36 = *(*int32)(unsafe.Add(mBase, _c_F_zzlLexValueLteMax[0]))
							if v27 == v36 {
								v162 = v34
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, _c_F_zzlLexValueLteMax[1]))
								if v28 == v39 {
									v162 = v34
								} else {
									if v28 != v36 {
										if v27 == v39 {
											v162 = int32(0)
										} else {
											v45 = int32(0)
											v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-1)))))
											switch v52 & int32(7) {
											case 0:
												v69 = int32(base.Ui32(v52) >> (uint(int32(3)) % 32))
											case 1:
												v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-3)))))
												v69 = v59
											case 2:
												v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(-5)))))
												v69 = v62
											case 3:
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-9))))
												v69 = v65
											case 4:
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-17))))
												v69 = v68
											default:
												v69 = v45
											}
											v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-1)))))
											switch v72 & int32(7) {
											case 0:
												v89 = int32(base.Ui32(v72) >> (uint(int32(3)) % 32))
											case 1:
												v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))))
												v89 = v79
											case 2:
												v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))))
												v89 = v82
											case 3:
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9))))
												v89 = v85
											case 4:
												v88 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-17))))
												v89 = v88
											default:
												v89 = v45
											}
											v90 = base.B2i32(base.Ui32(v69) < base.Ui32(v89))
											if base.Ui32(v69) < base.Ui32(v89) {
												v91 = v69
											} else {
												v91 = v89
											}
											v92 = F_memcmp(m, v27, v28, v91)
											mBase = m.M
											if v92 != 0 {
												v95 = v92
											} else {
												v95 = base.B2i32(base.Ui32(v89) < base.Ui32(v69)) - v90
											}
											v162 = int32(base.Ui32(v95) >> (uint(int32(31)) % 32))
										}
									} else {
										v162 = int32(0)
									}
								}
							}
						} else {
							v162 = int32(0)
						}
					}
					F_sdsfree(m, v27)
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v162
					}
				}
			}
		}
	}
}
func F_zzlNext(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v5 == int32(0) {
		F__serverAssert(m, int32(_a_F_zzlNext_0), int32(_a_F_zzlNext_1), int32(933))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v8 == int32(0) {
			F__serverAssert(m, int32(_a_F_zzlNext_0), int32(_a_F_zzlNext_1), int32(933))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v11 = F_lpNext(m, l0, v8)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				if v11 != 0 {
					v14 = F_lpNext(m, l0, v11)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						if v14 == int32(0) {
							F__serverAssert(m, int32(_a_F_zzlNext_2), int32(_a_F_zzlNext_1), int32(938))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v18 = v14
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v11
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
							return
						}
					}
				} else {
					v18 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v11
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
					return
				}
			}
		}
	}
}
func F_zzlPrev(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v5 == int32(0) {
		F__serverAssert(m, int32(_a_F_zzlPrev_0), int32(_a_F_zzlPrev_1), int32(952))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v8 == int32(0) {
			F__serverAssert(m, int32(_a_F_zzlPrev_0), int32(_a_F_zzlPrev_1), int32(952))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v11 = F_lpPrev(m, l0, v5)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				if v11 != 0 {
					v14 = F_lpPrev(m, l0, v11)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						if v14 == int32(0) {
							F__serverAssert(m, int32(_a_F_zzlPrev_2), int32(_a_F_zzlPrev_1), int32(957))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v18 = v14
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v18
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v11
							return
						}
					}
				} else {
					v18 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v18
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v11
					return
				}
			}
		}
	}
}
