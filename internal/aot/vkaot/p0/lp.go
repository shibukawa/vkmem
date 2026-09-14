package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_lpDeleteRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l2 == int32(0) {
		v66 = l0
		m.G0 = v9 + int32(16)
		return v66
	} else {
		v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
		v14 = F_lpSeek(m, l0, l1)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v14
			if v14 == int32(0) {
				v66 = l0
				m.G0 = v9 + int32(16)
				return v66
			} else {
				if v13 == int32(65535) {
					v64 = F_lpDeleteRangeWithEntry(m, l0, v9+int32(12), l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v66 = v64
						m.G0 = v9 + int32(16)
						return v66
					}
				} else {
					if v13 != int32(65535) {
						v28 = v13
					} else {
						v28 = int32(0)
					}
					v30 = l1>>(uint(int32(31))%32)&v28 + l1
					if base.Ui32(l2) < base.Ui32(v13-v30) {
						v64 = F_lpDeleteRangeWithEntry(m, l0, v9+int32(12), l2)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v66 = v64
							m.G0 = v9 + int32(16)
							return v66
						}
					} else {
						v33 = int32(255)
						*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v33)
						v35 = int32(8)
						v36 = int32(base.Ui32(v30) >> (uint(v35) % 32))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v36)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v30)
						v41 = v14 - l0 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v41)
						v44 = int32(base.Ui32(v41) >> (uint(int32(24)) % 32))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v44)
						v47 = int32(base.Ui32(v41) >> (uint(int32(16)) % 32))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v47)
						v50 = int32(base.Ui32(v41) >> (uint(v35) % 32))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v50)
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-8))))
						if base.Ui32(v54&int32(2147483647)) <= base.Ui32(v41) {
							v66 = l0
							m.G0 = v9 + int32(16)
							return v66
						} else {
							v59 = F_zrealloc_usable(m, l0, v41, int32(0))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								v66 = v59
								m.G0 = v9 + int32(16)
								return v66
							}
						}
					}
				}
			}
		}
	}
}
func F_lpFind(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v181 int64
	_ = v181
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v203 int64
	_ = v203
	var v208 int64
	_ = v208
	var v212 int32
	_ = v212
	var v214 int64
	_ = v214
	var v216 int32
	_ = v216
	var v224 int64
	_ = v224
	var v248 int64
	_ = v248
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v278 int64
	_ = v278
	var v279 int64
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(123456789)
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_lpFind_0), int32(_a_F_lpFind_1), int32(657))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L14
	} else {
		goto L100
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_lpFind_2), int32(_a_F_lpFind_1), int32(654))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L14
	} else {
		goto L99
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_lpFind_3), int32(_a_F_lpFind_1), int32(616))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L14
	} else {
		goto L98
	}
L4:
	;
	F__serverAssert(m, int32(_a_F_lpFind_4), int32(_a_F_lpFind_1), int32(610))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L14
	} else {
		goto L97
	}
L5:
	;
	v22 = l0 + int32(6)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = l0 + v23
	v30 = int32(0)
	v32 = v30
	v33 = l1
	v42 = v30
	goto L7
L6:
	;
	m.G0 = v15 + int32(32)
	return v373
L7:
	;
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	if v360+int32(1) != v24 {
		goto L1
	} else {
		goto L96
	}
L9:
	;
	if base.Ui32(v360) < base.Ui32(v22) {
		goto L2
	} else {
		goto L93
	}
L10:
	;
	v286 = int32(-1)
	v291 = int32(1)
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v293 = base.I32_extend8_s(v292)
	if v293 <= v286 {
		goto L72
	} else {
		goto L73
	}
L11:
	;
	v49 = F_lpGetWithSize(m, v33, v15+int32(24), int32(0), v15+int32(8))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v359 = l4
	v360 = v33 + v284
	v361 = v282
	goto L9
L13:
	;
	v125 = int32(255)
	v126 = v42 & v125
	if v126 == v125 {
		v282 = v42
		goto L12
	} else {
		goto L37
	}
L14:
	;
	return int32(0)
L15:
	;
	if v49 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v33) < base.Ui32(v22) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if base.Ui32(v24) <= base.Ui32(v33+v56) {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	if v59 != base.I64_extend_i32_u(l3) {
		v282 = v42
		goto L12
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(l3) < base.Ui32(int32(4)) {
		v84 = v49
		v85 = l2
		v86 = l3
		goto L23
	} else {
		goto L24
	}
L20:
	;
	if v124 != 0 {
		v282 = v42
		goto L12
	} else {
		goto L36
	}
L21:
	;
	v124 = int32(0)
	goto L20
L22:
	;
	v96 = v91
	v97 = v92
	v98 = v93
	goto L32
L23:
	;
	if v86 == int32(0) {
		goto L21
	} else {
		goto L30
	}
L24:
	;
	if (l2|v49)&int32(3) != 0 {
		v91 = v49
		v92 = l2
		v93 = l3
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v68 = v49
	v69 = l2
	v70 = l3
	goto L26
L26:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v73 != v74 {
		v91 = v68
		v92 = v69
		v93 = v70
		goto L22
	} else {
		goto L28
	}
L27:
	;
	v84 = v79
	v85 = v77
	v86 = v81
	goto L23
L28:
	;
	v76 = int32(4)
	v77 = v69 + v76
	v79 = v68 + v76
	v81 = v70 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v81) {
		v68 = v79
		v69 = v77
		v70 = v81
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v91 = v84
	v92 = v85
	v93 = v86
	goto L22
L31:
	;
	v124 = v101 - v102
	goto L20
L32:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v101 != v102 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v104 = int32(1)
	v109 = v98 + int32(-1)
	if v109 == int32(0) {
		goto L21
	} else {
		goto L35
	}
L35:
	;
	v96 = v96 + v104
	v97 = v97 + v104
	v98 = v109
	goto L32
L36:
	;
	v373 = v33
	goto L6
L37:
	;
	if v126 != 0 {
		v277 = v42
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	if v278 == v279 {
		v373 = v33
		goto L6
	} else {
		goto L69
	}
L39:
	;
	v129 = int32(255)
	if base.Ui32(l3+int32(-32)) < base.Ui32(int32(-31)) {
		v282 = v129
		goto L12
	} else {
		goto L40
	}
L40:
	;
	v132 = v15 + int32(16)
	v133 = int32(0)
	if base.Ui32(l3+int32(-21)) < base.Ui32(int32(-20)) {
		v267 = v133
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v267 == int32(0) {
		v282 = v129
		goto L12
	} else {
		goto L68
	}
L42:
	;
	goto L41
L43:
	;
	v145 = int32(1)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if l3 != v145 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v267 = int32(1)
	goto L42
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v132))) = v248
	goto L44
L46:
	;
	if v146&int32(255) == int32(45) {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v150 = v146 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v150&int32(255)) {
		v267 = v133
		goto L42
	} else {
		goto L48
	}
L48:
	;
	if v132 == int32(0) {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	v248 = base.I64_extend_i32_u(v150) & int64(255)
	goto L45
L50:
	;
	if base.Ui32(int32(8)) < base.Ui32((v169+int32(-49))&int32(255)) {
		v267 = v133
		goto L42
	} else {
		goto L53
	}
L51:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	v168 = int32(2)
	v169 = v166
	v170 = l2 + int32(1)
	goto L50
L52:
	;
	v168 = v145
	v169 = v146
	v170 = l2
	goto L50
L53:
	;
	v181 = base.I64_extend_i32_u(v169+int32(-48)) & int64(255)
	if base.Ui32(l3) <= base.Ui32(v168) {
		v224 = v181
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v146&int32(255) != int32(45) {
		goto L62
	} else {
		goto L63
	}
L55:
	;
	v187 = v168
	v189 = v181
	v191 = v170
	goto L56
L56:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	if base.Ui32((v193+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v267 = v133
		goto L42
	} else {
		goto L58
	}
L57:
	;
	v224 = v214
	goto L54
L58:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v189) {
		v267 = v133
		goto L42
	} else {
		goto L59
	}
L59:
	;
	v203 = v189 * int64(10)
	v208 = base.I64_extend_i32_u(v193+int32(-48)) & int64(255)
	if base.Ui64(v208^int64(-1)) < base.Ui64(v203) {
		v267 = v133
		goto L42
	} else {
		goto L60
	}
L60:
	;
	v212 = int32(1)
	v214 = v203 + v208
	v216 = v187 + v212
	if v216 != l3 {
		v187 = v216
		v189 = v214
		v191 = v191 + v212
		goto L56
	} else {
		goto L61
	}
L61:
	;
	goto L57
L62:
	;
	if v224 < int64(0) {
		v267 = v133
		goto L42
	} else {
		goto L66
	}
L63:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v224) {
		v267 = v133
		goto L42
	} else {
		goto L64
	}
L64:
	;
	if v132 == int32(0) {
		goto L44
	} else {
		goto L65
	}
L65:
	;
	v248 = int64(0) - v224
	goto L45
L66:
	;
	if v132 == int32(0) {
		goto L44
	} else {
		goto L67
	}
L67:
	;
	v248 = v224
	goto L45
L68:
	;
	v277 = int32(1)
	goto L38
L69:
	;
	v282 = v277
	goto L12
L70:
	;
	v359 = v32 + v286
	v360 = v33 + v356 + v355
	v361 = v42
	goto L9
L71:
	;
	goto L70
L72:
	;
	if v292&int32(192) != int32(128) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v355 = v291
	v356 = int32(1)
	goto L71
L74:
	;
	if v292&int32(224) != int32(192) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v301 = int32(1)
	v355 = v301
	v356 = v292&int32(63) + v301
	goto L71
L76:
	;
	v314 = (v293 + int32(15)) & int32(255)
	if base.Ui32(v314) < base.Ui32(int32(4)) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v355 = v291
	v356 = int32(2)
	goto L71
L78:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v314<<(uint(int32(2))%32))+uint32(_c_F_lpFind[0])))
	v355 = v291
	v356 = v354
	goto L71
L79:
	;
	if v292&int32(240) != int32(224) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if base.Ui32(v336) < base.Ui32(int32(128)) {
		v355 = v291
		v356 = v336
		goto L71
	} else {
		goto L85
	}
L81:
	;
	switch v292 + int32(-240) {
	case 0:
		goto L83
	default:
		goto L84
	case 15:
		v355 = v291
		v356 = int32(1)
		goto L71
	}
L82:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	v336 = v321 | v292<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L80
L83:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v33)+1))
	v336 = v333 + int32(5)
	goto L80
L84:
	;
	v355 = v291
	v356 = int32(0)
	goto L71
L85:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v336) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v336) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v355 = int32(2)
	v356 = v336
	goto L71
L88:
	;
	if base.Ui32(v336) < base.Ui32(int32(268435456)) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v355 = int32(3)
	v356 = v336
	goto L71
L90:
	;
	v349 = int32(4)
	goto L92
L91:
	;
	v349 = int32(5)
	goto L92
L92:
	;
	v355 = v349
	v356 = v336
	goto L71
L93:
	;
	if base.Ui32(v24) <= base.Ui32(v360) {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	if v365 != int32(255) {
		v32 = v359
		v33 = v360
		v42 = v361
		goto L7
	} else {
		goto L95
	}
L95:
	;
	goto L8
L96:
	;
	v373 = int32(0)
	goto L6
L97:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lpFree(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_valkey_free(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_lpGetIntegerIfValid(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v70 int64
	_ = v70
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v92 int64
	_ = v92
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v103 int64
	_ = v103
	var v105 int32
	_ = v105
	var v113 int64
	_ = v113
	var v137 int64
	_ = v137
	var v156 int32
	_ = v156
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v179 int32
	_ = v179
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v12 = F_lpGet(m, l0, v7+int32(8), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_lpGetIntegerIfValid_0), int32(_a_F_lpGetIntegerIfValid_1), int32(282))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L40
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return v170
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v22 = int32(0)
	if base.Ui32(v21+int32(-21)) < base.Ui32(int32(-20)) {
		v156 = v22
		goto L10
	} else {
		goto L11
	}
L4:
	;
	return int64(0)
L5:
	;
	if v12 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if l1 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
	v170 = v20
	goto L2
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	goto L7
L9:
	;
	if l1 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L10:
	;
	goto L9
L11:
	;
	v34 = int32(1)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v21 != v34 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v156 = int32(1)
	goto L10
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v137
	goto L12
L14:
	;
	if v35&int32(255) == int32(45) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v39 = v35 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v39&int32(255)) {
		v156 = v22
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if v7 == int32(0) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v137 = base.I64_extend_i32_u(v39) & int64(255)
	goto L13
L18:
	;
	if base.Ui32(int32(8)) < base.Ui32((v58+int32(-49))&int32(255)) {
		v156 = v22
		goto L10
	} else {
		goto L21
	}
L19:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	v57 = int32(2)
	v58 = v55
	v59 = v12 + int32(1)
	goto L18
L20:
	;
	v57 = v34
	v58 = v35
	v59 = v12
	goto L18
L21:
	;
	v70 = base.I64_extend_i32_u(v58+int32(-48)) & int64(255)
	if base.Ui32(v21) <= base.Ui32(v57) {
		v113 = v70
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v35&int32(255) != int32(45) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v76 = v57
	v78 = v70
	v80 = v59
	goto L24
L24:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	if base.Ui32((v82+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v156 = v22
		goto L10
	} else {
		goto L26
	}
L25:
	;
	v113 = v103
	goto L22
L26:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v78) {
		v156 = v22
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v92 = v78 * int64(10)
	v97 = base.I64_extend_i32_u(v82+int32(-48)) & int64(255)
	if base.Ui64(v97^int64(-1)) < base.Ui64(v92) {
		v156 = v22
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v101 = int32(1)
	v103 = v92 + v97
	v105 = v76 + v101
	if v105 != v21 {
		v76 = v105
		v78 = v103
		v80 = v80 + v101
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	if v113 < int64(0) {
		v156 = v22
		goto L10
	} else {
		goto L34
	}
L31:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v113) {
		v156 = v22
		goto L10
	} else {
		goto L32
	}
L32:
	;
	if v7 == int32(0) {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v137 = int64(0) - v113
	goto L13
L34:
	;
	if v7 == int32(0) {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v137 = v113
	goto L13
L36:
	;
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v170 = v168
	goto L2
L37:
	;
	if v156 == int32(0) {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v156
	goto L36
L39:
	;
	goto L36
L40:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lpGetObject(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a_F_lpGetObject_0), int32(_a_F_lpGetObject_1), int32(894))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v11 = F_lpGetValue(m, l0, v5+int32(12), v5)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v11 == int32(0) {
				v20 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
				v21 = F_sdsfromlonglong(m, v20)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = v21
					m.G0 = v5 + int32(16)
					return v23
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				v18 = F_sdsnewlen(m, v11, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v23 = v18
					m.G0 = v5 + int32(16)
					return v23
				}
			}
		}
	}
}
func F_lpInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v167 int64
	_ = v167
	var v173 int32
	_ = v173
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v189 int64
	_ = v189
	var v194 int64
	_ = v194
	var v198 int32
	_ = v198
	var v200 int64
	_ = v200
	var v202 int32
	_ = v202
	var v210 int64
	_ = v210
	var v234 int64
	_ = v234
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int64
	_ = v264
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v298 int64
	_ = v298
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v315 int64
	_ = v315
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v335 int32
	_ = v335
	var v357 int32
	_ = v357
	var v358 int64
	_ = v358
	var v368 int32
	_ = v368
	var v371 int64
	_ = v371
	var v374 int64
	_ = v374
	var v377 int64
	_ = v377
	var v380 int64
	_ = v380
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int64
	_ = v388
	var v393 int32
	_ = v393
	var v396 int64
	_ = v396
	var v399 int32
	_ = v399
	var v405 int64
	_ = v405
	var v409 int32
	_ = v409
	var v415 int64
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v430 int64
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v448 int64
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int64
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int64
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v502 int32
	_ = v502
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v574 int32
	_ = v574
	var v591 int64
	_ = v591
	var v593 int64
	_ = v593
	var v594 int64
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v837 int32
	_ = v837
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v943 int32
	_ = v943
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v960 int32
	_ = v960
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v998 int32
	_ = v998
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1016 int32
	_ = v1016
	var v1023 int32
	_ = v1023
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	v4 = l3
	v24 = m.G0
	v26 = v24 - int32(32)
	m.G0 = v26
	v29 = l1 | l2
	if v29 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	m.G0 = v26 + int32(32)
	return v1046
L2:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v936))) = uint32(v594)
	v1046 = v936
	goto L1
L3:
	;
	v1036 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v936)+4)))
	if v1036 == int32(65535) {
		goto L2
	} else {
		goto L249
	}
L4:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L5:
	;
	F__serverAssert(m, int32(_a_F_lpInsert_0), int32(_a_F_lpInsert_1), int32(754))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L128
	} else {
		goto L248
	}
L6:
	;
	F__serverAssert(m, int32(_a_F_lpInsert_2), int32(_a_F_lpInsert_1), int32(718))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L128
	} else {
		goto L247
	}
L7:
	;
	v113 = int32(0)
	v114 = base.B2i32(v29 == v113)
	if l1 == v113 {
		goto L43
	} else {
		goto L44
	}
L8:
	;
	v30 = l5
	goto L10
L9:
	;
	v30 = int32(2)
	goto L10
L10:
	;
	if v30 != int32(1) {
		v111 = l4
		v112 = v30
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v36 = int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v38 = base.I32_extend8_s(v37)
	if v38 <= int32(-1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if base.Ui32(v103) < base.Ui32(l0+int32(6)) {
		goto L6
	} else {
		goto L35
	}
L13:
	;
	v103 = l4 + v101 + v100
	goto L12
L14:
	;
	if v37&int32(192) != int32(128) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v100 = v36
	v101 = int32(1)
	goto L13
L16:
	;
	if v37&int32(224) != int32(192) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v46 = int32(1)
	v100 = v46
	v101 = v37&int32(63) + v46
	goto L13
L18:
	;
	v59 = (v38 + int32(15)) & int32(255)
	if base.Ui32(v59) < base.Ui32(int32(4)) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v100 = v36
	v101 = int32(2)
	goto L13
L20:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v59<<(uint(int32(2))%32))+uint32(_c_F_lpInsert[0])))
	v100 = v36
	v101 = v99
	goto L13
L21:
	;
	if v37&int32(240) != int32(224) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if base.Ui32(v81) < base.Ui32(int32(128)) {
		v100 = v36
		v101 = v81
		goto L13
	} else {
		goto L27
	}
L23:
	;
	switch v37 + int32(-240) {
	case 0:
		goto L25
	default:
		goto L26
	case 15:
		v100 = v36
		v101 = int32(1)
		goto L13
	}
L24:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	v81 = v66 | v37<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L22
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l4)+1))
	v81 = v78 + int32(5)
	goto L22
L26:
	;
	v100 = v36
	v101 = int32(0)
	goto L13
L27:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v81) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v81) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v100 = int32(2)
	v101 = v81
	goto L13
L30:
	;
	if base.Ui32(v81) < base.Ui32(int32(268435456)) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v100 = int32(3)
	v101 = v81
	goto L13
L32:
	;
	v94 = int32(4)
	goto L34
L33:
	;
	v94 = int32(5)
	goto L34
L34:
	;
	v100 = v94
	v101 = v81
	goto L13
L35:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(l0+v108) <= base.Ui32(v103) {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v111 = v103
	v112 = int32(0)
	goto L7
L37:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	v488 = int32(0)
	if v112 != int32(2) {
		v574 = v488
		goto L97
	} else {
		goto L98
	}
L38:
	;
	v479 = v472
	v480 = v473
	v481 = v474
	v483 = v470
	goto L37
L39:
	;
	if base.Ui64(int64(16383)) < base.Ui64(v388) {
		goto L91
	} else {
		goto L92
	}
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)) = uint8(v396)
	v470 = v393
	v472 = v399
	v473 = v396
	v474 = int32(1)
	goto L38
L41:
	;
	if base.Ui64(int64(127)) < base.Ui64(v388) {
		goto L39
	} else {
		goto L90
	}
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+19)) = uint8(v322)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+18)) = uint8(v325)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+17)) = uint8(v305)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+16)) = uint8(v326)
	v368 = int32(244)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)) = uint8(v368)
	v371 = int64(base.Ui64(v264) >> (uint(int64(56)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+23)) = uint8(v371)
	v374 = int64(base.Ui64(v264) >> (uint(int64(48)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)) = uint8(v374)
	v377 = int64(base.Ui64(v264) >> (uint(int64(40)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+21)) = uint8(v377)
	v380 = int64(base.Ui64(v264) >> (uint(int64(32)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+20)) = uint8(v380)
	v385 = v26 + int32(15)
	v387 = v261
	v388 = int64(9)
	goto L41
L43:
	;
	if l2 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L44:
	;
	v118 = v26 + int32(24)
	v119 = int32(0)
	if base.Ui32(v4+int32(-21)) < base.Ui32(int32(-20)) {
		v253 = v119
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v260 = int32(0)
	v261 = base.B2i32(v253 != v260)
	if v253 == v260 {
		goto L72
	} else {
		goto L73
	}
L46:
	;
	goto L45
L47:
	;
	v131 = int32(1)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v4 != v131 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v253 = int32(1)
	goto L46
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v118))) = v234
	goto L48
L50:
	;
	if v132&int32(255) == int32(45) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v136 = v132 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v136&int32(255)) {
		v253 = v119
		goto L46
	} else {
		goto L52
	}
L52:
	;
	if v118 == int32(0) {
		goto L48
	} else {
		goto L53
	}
L53:
	;
	v234 = base.I64_extend_i32_u(v136) & int64(255)
	goto L49
L54:
	;
	if base.Ui32(int32(8)) < base.Ui32((v155+int32(-49))&int32(255)) {
		v253 = v119
		goto L46
	} else {
		goto L57
	}
L55:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v154 = int32(2)
	v155 = v152
	v156 = l1 + int32(1)
	goto L54
L56:
	;
	v154 = v131
	v155 = v132
	v156 = l1
	goto L54
L57:
	;
	v167 = base.I64_extend_i32_u(v155+int32(-48)) & int64(255)
	if base.Ui32(v4) <= base.Ui32(v154) {
		v210 = v167
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v132&int32(255) != int32(45) {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	v173 = v154
	v175 = v167
	v177 = v156
	goto L60
L60:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)))
	if base.Ui32((v179+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v253 = v119
		goto L46
	} else {
		goto L62
	}
L61:
	;
	v210 = v200
	goto L58
L62:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v175) {
		v253 = v119
		goto L46
	} else {
		goto L63
	}
L63:
	;
	v189 = v175 * int64(10)
	v194 = base.I64_extend_i32_u(v179+int32(-48)) & int64(255)
	if base.Ui64(v194^int64(-1)) < base.Ui64(v189) {
		v253 = v119
		goto L46
	} else {
		goto L64
	}
L64:
	;
	v198 = int32(1)
	v200 = v189 + v194
	v202 = v173 + v198
	if v202 != v4 {
		v173 = v202
		v175 = v200
		v177 = v177 + v198
		goto L60
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	if v210 < int64(0) {
		v253 = v119
		goto L46
	} else {
		goto L70
	}
L67:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v210) {
		v253 = v119
		goto L46
	} else {
		goto L68
	}
L68:
	;
	if v118 == int32(0) {
		goto L48
	} else {
		goto L69
	}
L69:
	;
	v234 = int64(0) - v210
	goto L49
L70:
	;
	if v118 == int32(0) {
		goto L48
	} else {
		goto L71
	}
L71:
	;
	v234 = v210
	goto L49
L72:
	;
	if base.Ui32(int32(63)) < base.Ui32(v4) {
		goto L83
	} else {
		goto L84
	}
L73:
	;
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
	if base.Ui64(int64(127)) < base.Ui64(v264) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	if base.Ui64(int64(8191)) < base.Ui64(v264+int64(4096)) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)) = uint8(v264)
	v385 = v26 + int32(15)
	v387 = v261
	v388 = int64(1)
	goto L41
L76:
	;
	if base.Ui64(int64(65535)) < base.Ui64(v264+int64(32768)) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+16)) = uint8(v264)
	v285 = int32(base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v264)>>(uint(int64(50))%64))&int64(8192)+v264))>>(uint(int32(8))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)) = uint8(v285)
	v385 = v26 + int32(15)
	v387 = v261
	v388 = int64(2)
	goto L41
L78:
	;
	v305 = base.I32_wrap_i64(int64(base.Ui64(v264) >> (uint(int64(8)) % 64)))
	if base.Ui64(int64(16777215)) < base.Ui64(v264+int64(8388608)) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v294 = int32(241)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)) = uint8(v294)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+16)) = uint8(v264)
	v298 = int64(base.Ui64(v264) >> (uint(int64(8)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+17)) = uint8(v298)
	v385 = v26 + int32(15)
	v387 = v261
	v388 = int64(3)
	goto L41
L80:
	;
	v322 = base.I32_wrap_i64(int64(base.Ui64(v264) >> (uint(int64(24)) % 64)))
	v325 = base.I32_wrap_i64(int64(base.Ui64(v264) >> (uint(int64(16)) % 64)))
	v326 = base.I32_wrap_i64(v264)
	if base.Ui64(int64(4294967295)) < base.Ui64(v264+int64(2147483648)) {
		goto L42
	} else {
		goto L82
	}
L81:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+17)) = uint8(v305)
	v311 = int32(242)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)) = uint8(v311)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+16)) = uint8(v264)
	v315 = int64(base.Ui64(v264) >> (uint(int64(16)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+18)) = uint8(v315)
	v385 = v26 + int32(15)
	v387 = v261
	v388 = int64(4)
	goto L41
L82:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+19)) = uint8(v322)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+18)) = uint8(v325)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+17)) = uint8(v305)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+16)) = uint8(v326)
	v335 = int32(243)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)) = uint8(v335)
	v385 = v26 + int32(15)
	v387 = v261
	v388 = int64(5)
	goto L41
L83:
	;
	if base.Ui32(int32(4095)) < base.Ui32(v4) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v385 = l2
	v387 = v261
	v388 = base.I64_extend_i32_u(v4 + int32(1))
	goto L41
L85:
	;
	v385 = l2
	v387 = v261
	v388 = base.I64_extend_i32_u(v4) + int64(5)
	goto L41
L86:
	;
	v385 = l2
	v387 = v261
	v388 = base.I64_extend_i32_u(v4 + int32(2))
	goto L41
L87:
	;
	v357 = int32(0)
	v358 = int64(0)
	if v114 == v357 {
		v393 = v357
		v396 = v358
		v399 = v357
		goto L40
	} else {
		goto L89
	}
L88:
	;
	v385 = l2
	v387 = int32(1)
	v388 = base.I64_extend_i32_u(v4)
	goto L41
L89:
	;
	v479 = v357
	v480 = v358
	v481 = v357
	v483 = v357
	goto L37
L90:
	;
	v393 = v385
	v396 = v388
	v399 = v387
	goto L40
L91:
	;
	if base.Ui64(int64(2097151)) < base.Ui64(v388) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v405 = int64(base.Ui64(v388) >> (uint(int64(7)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)) = uint8(v405)
	v409 = base.I32_wrap_i64(v388) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+25)) = uint8(v409)
	v470 = v385
	v472 = v387
	v473 = v388
	v474 = int32(2)
	goto L38
L93:
	;
	if base.Ui64(int64(268435455)) < base.Ui64(v388) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v415 = int64(base.Ui64(v388) >> (uint(int64(14)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)) = uint8(v415)
	v417 = base.I32_wrap_i64(v388)
	v418 = int32(128)
	v419 = v417 | v418
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+26)) = uint8(v419)
	v424 = int32(base.Ui32(v417)>>(uint(int32(7))%32)) | v418
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+25)) = uint8(v424)
	v470 = v385
	v472 = v387
	v473 = v388
	v474 = int32(3)
	goto L38
L95:
	;
	v448 = int64(base.Ui64(v388) >> (uint(int64(28)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)) = uint8(v448)
	v450 = base.I32_wrap_i64(v388)
	v451 = int32(128)
	v452 = v450 | v451
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+28)) = uint8(v452)
	v457 = int32(base.Ui32(v450)>>(uint(int32(7))%32)) | v451
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+27)) = uint8(v457)
	v462 = int32(base.Ui32(v450)>>(uint(int32(14))%32)) | v451
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+26)) = uint8(v462)
	v467 = int32(base.Ui32(v450)>>(uint(int32(21))%32)) | v451
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+25)) = uint8(v467)
	v470 = v385
	v472 = v387
	v473 = v388
	v474 = int32(5)
	goto L38
L96:
	;
	v430 = int64(base.Ui64(v388) >> (uint(int64(21)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)) = uint8(v430)
	v432 = base.I32_wrap_i64(v388)
	v433 = int32(128)
	v434 = v432 | v433
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+27)) = uint8(v434)
	v439 = int32(base.Ui32(v432)>>(uint(int32(7))%32)) | v433
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+26)) = uint8(v439)
	v444 = int32(base.Ui32(v432)>>(uint(int32(14))%32)) | v433
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+25)) = uint8(v444)
	v470 = v385
	v472 = v387
	v473 = v388
	v474 = int32(4)
	goto L38
L97:
	;
	v591 = base.I64_extend_i32_u(v485)<<(uint(int64(8))%64) | base.I64_extend_i32_u(v486) | base.I64_extend_i32_u(v484)<<(uint(int64(16))%64) | base.I64_extend_i32_u(v487)<<(uint(int64(24))%64)
	v593 = base.I64_extend_i32_u(v574)
	v594 = v480 + base.I64_extend_i32_u(v481) + v591 - v593
	if base.Ui64(int64(4294967295)) < base.Ui64(v594) {
		v1046 = v488
		goto L1
	} else {
		goto L123
	}
L98:
	;
	v492 = int32(1)
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v494 = base.I32_extend8_s(v493)
	if v494 <= int32(-1) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if base.Ui32(v111) < base.Ui32(l0+int32(6)) {
		goto L5
	} else {
		goto L121
	}
L100:
	;
	if v493&int32(192) != int32(128) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v556 = v492
	v557 = int32(1)
	goto L99
L102:
	;
	if v493&int32(224) != int32(192) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v502 = int32(1)
	v556 = v502
	v557 = v493&int32(63) + v502
	goto L99
L104:
	;
	v515 = (v494 + int32(15)) & int32(255)
	if base.Ui32(v515) < base.Ui32(int32(4)) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v556 = v492
	v557 = int32(2)
	goto L99
L106:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v515<<(uint(int32(2))%32))+uint32(_c_F_lpInsert[0])))
	v556 = v492
	v557 = v555
	goto L99
L107:
	;
	if v493&int32(240) != int32(224) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	if base.Ui32(v537) < base.Ui32(int32(128)) {
		v556 = v492
		v557 = v537
		goto L99
	} else {
		goto L113
	}
L109:
	;
	switch v493 + int32(-240) {
	case 0:
		goto L111
	default:
		goto L112
	case 15:
		v556 = v492
		v557 = int32(1)
		goto L99
	}
L110:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
	v537 = v522 | v493<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L108
L111:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v111)+1))
	v537 = v534 + int32(5)
	goto L108
L112:
	;
	v556 = v492
	v557 = int32(0)
	goto L99
L113:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v537) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v537) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v556 = int32(2)
	v557 = v537
	goto L99
L116:
	;
	if base.Ui32(v537) < base.Ui32(int32(268435456)) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v556 = int32(3)
	v557 = v537
	goto L99
L118:
	;
	v550 = int32(4)
	goto L120
L119:
	;
	v550 = int32(5)
	goto L120
L120:
	;
	v556 = v550
	v557 = v537
	goto L99
L121:
	;
	v561 = v556 + v557
	if base.Ui32(l0+(v485<<(uint(int32(8))%32)|v486|v484<<(uint(int32(16))%32)|v487<<(uint(int32(24))%32))) <= base.Ui32(v111+v561) {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v574 = v561
	goto L97
L123:
	;
	v597 = v111 - l0
	v598 = l0 + v597
	if base.Ui64(v594) <= base.Ui64(v591) {
		v617 = l0
		v619 = v598
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v620 = base.I32_wrap_i64(v480)
	v622 = v619 + v620 + v481
	if v112 != 0 {
		goto L132
	} else {
		goto L133
	}
L125:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-8))))
	goto L126
L126:
	;
	if base.Ui64(v594) <= base.Ui64(base.I64_extend_i32_u(v602&int32(2147483647))) {
		v617 = l0
		v619 = v598
		goto L124
	} else {
		goto L127
	}
L127:
	;
	v607 = int32(0)
	v610 = F_zrealloc_usable(m, l0, base.I32_wrap_i64(v594), v607)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	return int32(0)
L129:
	;
	if v610 == int32(0) {
		v1046 = v607
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v617 = v610
	v619 = v610 + v597
	goto L124
L131:
	;
	if base.Ui64(v594) < base.Ui64(v591) {
		goto L217
	} else {
		goto L218
	}
L132:
	;
	v773 = v619 + v574
	v777 = base.I32_wrap_i64(v591 - (base.I64_extend_i32_u(v597) + v593))
	if v622 == v773 {
		goto L176
	} else {
		goto L177
	}
L133:
	;
	v624 = base.I32_wrap_i64(v591) - v597
	if v622 == v619 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	goto L131
L135:
	;
	goto L134
L136:
	;
	v628 = v624 + v622
	if base.Ui32(int32(0)-v624<<(uint(int32(1))%32)) < base.Ui32(v619-v628) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v638 = (v619 ^ v622) & int32(3)
	if base.Ui32(v619) <= base.Ui32(v622) {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v635 = F___memcpy(m, v622, v619, v624)
	mBase = m.M
	goto L134
L139:
	;
	if v744 == int32(0) {
		goto L135
	} else {
		goto L171
	}
L140:
	;
	if base.Ui32(v722) <= base.Ui32(int32(3)) {
		v743 = v721
		v744 = v722
		v745 = v723
		goto L139
	} else {
		goto L167
	}
L141:
	;
	if v638 != 0 {
		v704 = v624
		goto L151
	} else {
		goto L152
	}
L142:
	;
	if v638 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	if v622&int32(3) != 0 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v743 = v619
	v744 = v624
	v745 = v622
	goto L139
L145:
	;
	v645 = v619
	v646 = v624
	v647 = v622
	goto L147
L146:
	;
	v721 = v619
	v722 = v624
	v723 = v622
	goto L140
L147:
	;
	if v646 == int32(0) {
		goto L135
	} else {
		goto L149
	}
L149:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645))))
	*(*uint8)(unsafe.Add(mBase, uint32(v647))) = uint8(v651)
	v653 = int32(1)
	v654 = v645 + v653
	v656 = v646 + int32(-1)
	v658 = v647 + v653
	if v658&int32(3) == int32(0) {
		v721 = v654
		v722 = v656
		v723 = v658
		goto L140
	} else {
		goto L150
	}
L150:
	;
	v645 = v654
	v646 = v656
	v647 = v658
	goto L147
L151:
	;
	if v704 == int32(0) {
		goto L135
	} else {
		goto L163
	}
L152:
	;
	if v628&int32(3) == int32(0) {
		v684 = v624
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if base.Ui32(v684) <= base.Ui32(int32(3)) {
		v704 = v684
		goto L151
	} else {
		goto L159
	}
L154:
	;
	v669 = v624
	goto L155
L155:
	;
	if v669 == int32(0) {
		goto L135
	} else {
		goto L157
	}
L156:
	;
	v684 = v675
	goto L153
L157:
	;
	v675 = v669 + int32(-1)
	v676 = v622 + v675
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619+v675))))
	*(*uint8)(unsafe.Add(mBase, uint32(v676))) = uint8(v678)
	if v676&int32(3) != 0 {
		v669 = v675
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v691 = v684
	goto L160
L160:
	;
	v695 = v691 + int32(-4)
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v619+v695)))
	*(*int32)(unsafe.Add(mBase, uint32(v622+v695))) = v698
	if base.Ui32(int32(3)) < base.Ui32(v695) {
		v691 = v695
		goto L160
	} else {
		goto L162
	}
L161:
	;
	v704 = v695
	goto L151
L162:
	;
	goto L161
L163:
	;
	v711 = v704
	goto L164
L164:
	;
	v715 = v711 + int32(-1)
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619+v715))))
	*(*uint8)(unsafe.Add(mBase, uint32(v622+v715))) = uint8(v718)
	if v715 != 0 {
		v711 = v715
		goto L164
	} else {
		goto L166
	}
L166:
	;
	goto L135
L167:
	;
	v728 = v721
	v729 = v722
	v730 = v723
	goto L168
L168:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v728)))
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = v732
	v734 = int32(4)
	v735 = v728 + v734
	v737 = v730 + v734
	v739 = v729 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v739) {
		v728 = v735
		v729 = v739
		v730 = v737
		goto L168
	} else {
		goto L170
	}
L169:
	;
	v743 = v735
	v744 = v739
	v745 = v737
	goto L139
L170:
	;
	goto L169
L171:
	;
	v750 = v743
	v751 = v744
	v752 = v745
	goto L172
L172:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750))))
	*(*uint8)(unsafe.Add(mBase, uint32(v752))) = uint8(v754)
	v756 = int32(1)
	v761 = v751 + int32(-1)
	if v761 != 0 {
		v750 = v750 + v756
		v751 = v761
		v752 = v752 + v756
		goto L172
	} else {
		goto L174
	}
L173:
	;
	goto L135
L174:
	;
	goto L173
L175:
	;
	goto L131
L176:
	;
	goto L175
L177:
	;
	v781 = v777 + v622
	if base.Ui32(int32(0)-v777<<(uint(int32(1))%32)) < base.Ui32(v773-v781) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v791 = (v773 ^ v622) & int32(3)
	if base.Ui32(v773) <= base.Ui32(v622) {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	v788 = F___memcpy(m, v622, v773, v777)
	mBase = m.M
	goto L175
L180:
	;
	if v897 == int32(0) {
		goto L176
	} else {
		goto L212
	}
L181:
	;
	if base.Ui32(v875) <= base.Ui32(int32(3)) {
		v896 = v874
		v897 = v875
		v898 = v876
		goto L180
	} else {
		goto L208
	}
L182:
	;
	if v791 != 0 {
		v857 = v777
		goto L192
	} else {
		goto L193
	}
L183:
	;
	if v791 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	if v622&int32(3) != 0 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v896 = v773
	v897 = v777
	v898 = v622
	goto L180
L186:
	;
	v798 = v773
	v799 = v777
	v800 = v622
	goto L188
L187:
	;
	v874 = v773
	v875 = v777
	v876 = v622
	goto L181
L188:
	;
	if v799 == int32(0) {
		goto L176
	} else {
		goto L190
	}
L190:
	;
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
	*(*uint8)(unsafe.Add(mBase, uint32(v800))) = uint8(v804)
	v806 = int32(1)
	v807 = v798 + v806
	v809 = v799 + int32(-1)
	v811 = v800 + v806
	if v811&int32(3) == int32(0) {
		v874 = v807
		v875 = v809
		v876 = v811
		goto L181
	} else {
		goto L191
	}
L191:
	;
	v798 = v807
	v799 = v809
	v800 = v811
	goto L188
L192:
	;
	if v857 == int32(0) {
		goto L176
	} else {
		goto L204
	}
L193:
	;
	if v781&int32(3) == int32(0) {
		v837 = v777
		goto L194
	} else {
		goto L195
	}
L194:
	;
	if base.Ui32(v837) <= base.Ui32(int32(3)) {
		v857 = v837
		goto L192
	} else {
		goto L200
	}
L195:
	;
	v822 = v777
	goto L196
L196:
	;
	if v822 == int32(0) {
		goto L176
	} else {
		goto L198
	}
L197:
	;
	v837 = v828
	goto L194
L198:
	;
	v828 = v822 + int32(-1)
	v829 = v622 + v828
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773+v828))))
	*(*uint8)(unsafe.Add(mBase, uint32(v829))) = uint8(v831)
	if v829&int32(3) != 0 {
		v822 = v828
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	v844 = v837
	goto L201
L201:
	;
	v848 = v844 + int32(-4)
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v773+v848)))
	*(*int32)(unsafe.Add(mBase, uint32(v622+v848))) = v851
	if base.Ui32(int32(3)) < base.Ui32(v848) {
		v844 = v848
		goto L201
	} else {
		goto L203
	}
L202:
	;
	v857 = v848
	goto L192
L203:
	;
	goto L202
L204:
	;
	v864 = v857
	goto L205
L205:
	;
	v868 = v864 + int32(-1)
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773+v868))))
	*(*uint8)(unsafe.Add(mBase, uint32(v622+v868))) = uint8(v871)
	if v868 != 0 {
		v864 = v868
		goto L205
	} else {
		goto L207
	}
L207:
	;
	goto L176
L208:
	;
	v881 = v874
	v882 = v875
	v883 = v876
	goto L209
L209:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v881)))
	*(*int32)(unsafe.Add(mBase, uint32(v883))) = v885
	v887 = int32(4)
	v888 = v881 + v887
	v890 = v883 + v887
	v892 = v882 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v892) {
		v881 = v888
		v882 = v892
		v883 = v890
		goto L209
	} else {
		goto L211
	}
L210:
	;
	v896 = v888
	v897 = v892
	v898 = v890
	goto L180
L211:
	;
	goto L210
L212:
	;
	v903 = v896
	v904 = v897
	v905 = v898
	goto L213
L213:
	;
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v903))))
	*(*uint8)(unsafe.Add(mBase, uint32(v905))) = uint8(v907)
	v909 = int32(1)
	v914 = v904 + int32(-1)
	if v914 != 0 {
		v903 = v903 + v909
		v904 = v914
		v905 = v905 + v909
		goto L213
	} else {
		goto L215
	}
L214:
	;
	goto L176
L215:
	;
	goto L214
L216:
	;
	if l6 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L217:
	;
	v927 = int32(0)
	v930 = F_zrealloc_usable(m, v617, base.I32_wrap_i64(v594), v927)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L128
	} else {
		goto L219
	}
L218:
	;
	v936 = v617
	v937 = v619
	goto L216
L219:
	;
	if v930 == int32(0) {
		v1046 = v927
		goto L1
	} else {
		goto L220
	}
L220:
	;
	v936 = v930
	v937 = v930 + v597
	goto L216
L221:
	;
	if v479 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L222:
	;
	if v29 == v113 {
		goto L3
	} else {
		goto L226
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v937
	if v114 == int32(0) {
		goto L221
	} else {
		goto L224
	}
L224:
	;
	v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937))))
	if v943 != int32(255) {
		goto L3
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0)
	goto L3
L226:
	;
	goto L221
L227:
	;
	if v481 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L228:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L233
	}
L229:
	;
	v950 = base.I32_wrap_i64(v480)
	if v950 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	v992 = v950
	goto L227
L231:
	;
	goto L230
L232:
	;
	v953 = F__emscripten_memcpy_bulkmem(m, v937, v483, v950)
	mBase = m.M
	goto L231
L233:
	;
	if base.Ui32(int32(63)) < base.Ui32(v4) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	if v4 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L235:
	;
	if base.Ui32(int32(4095)) < base.Ui32(v4) {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v960 = v4 | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v937))) = uint8(v960)
	v985 = int32(1)
	goto L234
L237:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v937)+1)) = uint8(v4)
	v973 = int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v937))) = uint8(v973)
	v976 = int32(base.Ui32(v4) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v937)+4)) = uint8(v976)
	v979 = int32(base.Ui32(v4) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v937)+3)) = uint8(v979)
	v982 = int32(base.Ui32(v4) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v937)+2)) = uint8(v982)
	v985 = int32(5)
	goto L234
L238:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v937)+1)) = uint8(v4)
	v969 = int32(base.Ui32(v4)>>(uint(int32(8))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v937))) = uint8(v969)
	v985 = int32(2)
	goto L234
L239:
	;
	v992 = v620
	goto L227
L240:
	;
	goto L239
L241:
	;
	v989 = F__emscripten_memcpy_bulkmem(m, v937+v985, l1, v4)
	mBase = m.M
	goto L240
L242:
	;
	if v114|base.B2i32(v112 != int32(2)) == int32(0) {
		goto L2
	} else {
		goto L245
	}
L243:
	;
	goto L242
L244:
	;
	v998 = F__emscripten_memcpy_bulkmem(m, v937+v992, v26+int32(24), v481)
	mBase = m.M
	goto L243
L245:
	;
	v1005 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v936)+4)))
	if v1005 == int32(65535) {
		goto L2
	} else {
		goto L246
	}
L246:
	;
	v1009 = v1005 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v936)+4)) = uint16(v1009)
	goto L2
L247:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	goto L4
L249:
	;
	v1040 = v1036 + int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v936)+4)) = uint16(v1040)
	goto L2
}
func F_lpInsertInteger(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int64
	_ = v44
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int64
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v2 = l1
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if base.Ui64(int64(127)) < base.Ui64(v2) {
		if base.Ui64(int64(8191)) < base.Ui64(v2+int64(4096)) {
			if base.Ui64(int64(65535)) < base.Ui64(v2+int64(32768)) {
				v49 = base.I32_wrap_i64(int64(base.Ui64(v2) >> (uint(int64(8)) % 64)))
				if base.Ui64(int64(16777215)) < base.Ui64(v2+int64(8388608)) {
					v64 = base.I32_wrap_i64(int64(base.Ui64(v2) >> (uint(int64(24)) % 64)))
					v67 = base.I32_wrap_i64(int64(base.Ui64(v2) >> (uint(int64(16)) % 64)))
					v68 = base.I32_wrap_i64(v2)
					if base.Ui64(int64(4294967295)) < base.Ui64(v2+int64(2147483648)) {
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)) = uint8(v64)
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v67)
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+9)) = uint8(v49)
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)) = uint8(v68)
						v84 = int32(244)
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+7)) = uint8(v84)
						v87 = int64(base.Ui64(v2) >> (uint(int64(56)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v87)
						v90 = int64(base.Ui64(v2) >> (uint(int64(48)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v90)
						v93 = int64(base.Ui64(v2) >> (uint(int64(40)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v93)
						v96 = int64(base.Ui64(v2) >> (uint(int64(32)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)) = uint8(v96)
						v99 = int32(9)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)) = uint8(v64)
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v67)
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+9)) = uint8(v49)
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)) = uint8(v68)
						v77 = int32(243)
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+7)) = uint8(v77)
						v99 = int32(5)
					}
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v13)+9)) = uint8(v49)
					v55 = int32(242)
					*(*uint8)(unsafe.Add(mBase, uint32(v13)+7)) = uint8(v55)
					*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)) = uint8(v2)
					v59 = int64(base.Ui64(v2) >> (uint(int64(16)) % 64))
					*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v59)
					v99 = int32(4)
				}
			} else {
				v40 = int32(241)
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+7)) = uint8(v40)
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)) = uint8(v2)
				v44 = int64(base.Ui64(v2) >> (uint(int64(8)) % 64))
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+9)) = uint8(v44)
				v99 = int32(3)
			}
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)) = uint8(v2)
			v33 = int32(base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v2)>>(uint(int64(50))%64))&int64(8192)+v2))>>(uint(int32(8))%32)) | int32(192)
			*(*uint8)(unsafe.Add(mBase, uint32(v13)+7)) = uint8(v33)
			v99 = int32(2)
		}
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v13)+7)) = uint8(v2)
		v99 = int32(1)
	}
	v106 = F_lpInsert(m, l0, int32(0), v13+int32(7), v99, l2, l3, l4)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		return int32(0)
	} else {
		m.G0 = v13 + int32(16)
		return v106
	}
}
func F_lpNextRandom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v129 int64
	_ = v129
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	v6 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v12 = F_lpLength(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v243
L2:
	;
	return int32(0)
L3:
	;
	if base.Ui32(v12) <= base.Ui32(v11) {
		v243 = v6
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l1 == int32(0) {
		v243 = v6
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v23 = l1
	v28 = v11
	goto L7
L6:
	;
	v243 = int32(0)
	goto L1
L7:
	;
	if l4 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v233 = v28 + int32(1)
	if base.Ui32(v233) < base.Ui32(v12) {
		v23 = v231
		v28 = v233
		goto L7
	} else {
		goto L68
	}
L10:
	;
	v123 = int32(0)
	v125 = *(*int64)(unsafe.Add(mBase, _c_F_lpNextRandom[0]))
	v129 = v125*int64(6364136223846793005) + int64(1)
	*(*int64)(unsafe.Add(mBase, _c_F_lpNextRandom[0])) = v129
	goto L40
L11:
	;
	if v28&int32(1) == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v40 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v42 = base.I32_extend8_s(v41)
	if v42 <= int32(-1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v108 != int32(255) {
		v231 = v107
		goto L9
	} else {
		goto L36
	}
L14:
	;
	v107 = v23 + v105 + v104
	goto L13
L15:
	;
	if v41&int32(192) != int32(128) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v104 = v40
	v105 = int32(1)
	goto L14
L17:
	;
	if v41&int32(224) != int32(192) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v50 = int32(1)
	v104 = v50
	v105 = v41&int32(63) + v50
	goto L14
L19:
	;
	v63 = (v42 + int32(15)) & int32(255)
	if base.Ui32(v63) < base.Ui32(int32(4)) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v104 = v40
	v105 = int32(2)
	goto L14
L21:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v63<<(uint(int32(2))%32))+uint32(_c_F_lpNextRandom[1])))
	v104 = v40
	v105 = v103
	goto L14
L22:
	;
	if v41&int32(240) != int32(224) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if base.Ui32(v85) < base.Ui32(int32(128)) {
		v104 = v40
		v105 = v85
		goto L14
	} else {
		goto L28
	}
L24:
	;
	switch v41 + int32(-240) {
	case 0:
		goto L26
	default:
		goto L27
	case 15:
		v104 = v40
		v105 = int32(1)
		goto L14
	}
L25:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v85 = v70 | v41<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L23
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v23)+1))
	v85 = v82 + int32(5)
	goto L23
L27:
	;
	v104 = v40
	v105 = int32(0)
	goto L14
L28:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v85) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v85) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v104 = int32(2)
	v105 = v85
	goto L14
L31:
	;
	if base.Ui32(v85) < base.Ui32(int32(268435456)) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v104 = int32(3)
	v105 = v85
	goto L14
L33:
	;
	v98 = int32(4)
	goto L35
L34:
	;
	v98 = int32(5)
	goto L35
L35:
	;
	v104 = v98
	v105 = v85
	goto L14
L36:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v107+int32(1) == l0+v113 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	F__serverAssert(m, int32(_a_F_lpNextRandom_0), int32(_a_F_lpNextRandom_1), int32(400))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v149 = int32(1)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v151 = base.I32_extend8_s(v150)
	if v151 <= int32(-1) {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	if base.F64_le(base.F64_div(base.F64_convert_i32_s(base.I32_wrap_i64(int64(base.Ui64(v129)>>(uint(int64(33))%64)))), float64(2.147483647e+09)), base.F64_div(base.F64_convert_i32_u(l3), base.F64_convert_i32_u(int32(base.Ui32(v12-v28)>>(uint(base.B2i32(l4 != int32(0)))%32))))) == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v28
	return v23
L42:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	if v217 != int32(255) {
		v231 = v216
		goto L9
	} else {
		goto L65
	}
L43:
	;
	v216 = v23 + v214 + v213
	goto L42
L44:
	;
	if v150&int32(192) != int32(128) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v213 = v149
	v214 = int32(1)
	goto L43
L46:
	;
	if v150&int32(224) != int32(192) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v159 = int32(1)
	v213 = v159
	v214 = v150&int32(63) + v159
	goto L43
L48:
	;
	v172 = (v151 + int32(15)) & int32(255)
	if base.Ui32(v172) < base.Ui32(int32(4)) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v213 = v149
	v214 = int32(2)
	goto L43
L50:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v172<<(uint(int32(2))%32))+uint32(_c_F_lpNextRandom[1])))
	v213 = v149
	v214 = v212
	goto L43
L51:
	;
	if v150&int32(240) != int32(224) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v194) < base.Ui32(int32(128)) {
		v213 = v149
		v214 = v194
		goto L43
	} else {
		goto L57
	}
L53:
	;
	switch v150 + int32(-240) {
	case 0:
		goto L55
	default:
		goto L56
	case 15:
		v213 = v149
		v214 = int32(1)
		goto L43
	}
L54:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v194 = v179 | v150<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L52
L55:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v23)+1))
	v194 = v191 + int32(5)
	goto L52
L56:
	;
	v213 = v149
	v214 = int32(0)
	goto L43
L57:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v194) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v194) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v213 = int32(2)
	v214 = v194
	goto L43
L60:
	;
	if base.Ui32(v194) < base.Ui32(int32(268435456)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v213 = int32(3)
	v214 = v194
	goto L43
L62:
	;
	v207 = int32(4)
	goto L64
L63:
	;
	v207 = int32(5)
	goto L64
L64:
	;
	v213 = v207
	v214 = v194
	goto L43
L65:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v216+int32(1) == l0+v222 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	F__serverAssert(m, int32(_a_F_lpNextRandom_0), int32(_a_F_lpNextRandom_1), int32(400))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	goto L8
}
func F_lpPrependInteger(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v4 == int32(255) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v9 != int32(7) {
			F__serverAssert(m, int32(_a_F_lpPrependInteger_0), int32(_a_F_lpPrependInteger_1), int32(426))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v14 = l0 + int32(6)
			v15 = int32(0)
			v17 = F_lpInsertInteger(m, l0, l1, v14, v15, v15)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v17
			}
		}
	} else {
		v14 = l0 + int32(6)
		v15 = int32(0)
		v17 = F_lpInsertInteger(m, l0, l1, v14, v15, v15)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			return v17
		}
	}
}
func F_lpPrev(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v28 int64
	_ = v28
	var v33 int32
	_ = v33
	var v39 int64
	_ = v39
	var v44 int32
	_ = v44
	var v50 int64
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v67 int64
	_ = v67
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v98 int32
	_ = v98
	if l1 == int32(0) {
		F__serverAssert(m, int32(_a_F_lpPrev_0), int32(_a_F_lpPrev_1), int32(410))
		mBase = m.M
		v88 = m.ExcPending
		if v88 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if l1-l0 == int32(6) {
			v98 = int32(0)
		} else {
			v11 = int32(-1)
			v14 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+v11))))
			v17 = base.I64_extend_i32_u(v14 & int32(127))
			if v11 < v14 {
				v92 = v11
				v93 = v17
			} else {
				v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(-2)))))
				v28 = base.I64_extend_i32_u(v22&int32(127))<<(uint(int64(7))%64) | v17
				if int32(-1) < v22 {
					v67 = v28
					if base.Ui64(int64(128)) <= base.Ui64(v67) {
						if base.Ui64(int64(16384)) <= base.Ui64(v67) {
							if base.Ui64(int64(2097152)) <= base.Ui64(v67) {
								if base.Ui64(v67) < base.Ui64(int64(268435456)) {
									v81 = int32(-4)
								} else {
									v81 = int32(-5)
								}
								v92 = v81
								v93 = v67
							} else {
								v92 = int32(-3)
								v93 = v67
							}
						} else {
							v92 = int32(-2)
							v93 = v67
						}
					} else {
						v92 = int32(-1)
						v93 = v67
					}
				} else {
					v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
					v39 = base.I64_extend_i32_u(v33&int32(127))<<(uint(int64(14))%64) | v28
					if int32(-1) < v33 {
						v67 = v39
						if base.Ui64(int64(128)) <= base.Ui64(v67) {
							if base.Ui64(int64(16384)) <= base.Ui64(v67) {
								if base.Ui64(int64(2097152)) <= base.Ui64(v67) {
									if base.Ui64(v67) < base.Ui64(int64(268435456)) {
										v81 = int32(-4)
									} else {
										v81 = int32(-5)
									}
									v92 = v81
									v93 = v67
								} else {
									v92 = int32(-3)
									v93 = v67
								}
							} else {
								v92 = int32(-2)
								v93 = v67
							}
						} else {
							v92 = int32(-1)
							v93 = v67
						}
					} else {
						v44 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(-4)))))
						v50 = base.I64_extend_i32_u(v44&int32(127))<<(uint(int64(21))%64) | v39
						if int32(-1) < v44 {
							v67 = v50
							if base.Ui64(int64(128)) <= base.Ui64(v67) {
								if base.Ui64(int64(16384)) <= base.Ui64(v67) {
									if base.Ui64(int64(2097152)) <= base.Ui64(v67) {
										if base.Ui64(v67) < base.Ui64(int64(268435456)) {
											v81 = int32(-4)
										} else {
											v81 = int32(-5)
										}
										v92 = v81
										v93 = v67
									} else {
										v92 = int32(-3)
										v93 = v67
									}
								} else {
									v92 = int32(-2)
									v93 = v67
								}
							} else {
								v92 = int32(-1)
								v93 = v67
							}
						} else {
							v53 = int32(-5)
							v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+v53))))
							if v56 < int32(0) {
								v92 = v53
								v93 = int64(-1)
							} else {
								v67 = base.I64_extend_i32_u(v56&int32(127))<<(uint(int64(28))%64) | v50
								if base.Ui64(int64(128)) <= base.Ui64(v67) {
									if base.Ui64(int64(16384)) <= base.Ui64(v67) {
										if base.Ui64(int64(2097152)) <= base.Ui64(v67) {
											if base.Ui64(v67) < base.Ui64(int64(268435456)) {
												v81 = int32(-4)
											} else {
												v81 = int32(-5)
											}
											v92 = v81
											v93 = v67
										} else {
											v92 = int32(-3)
											v93 = v67
										}
									} else {
										v92 = int32(-2)
										v93 = v67
									}
								} else {
									v92 = int32(-1)
									v93 = v67
								}
							}
						}
					}
				}
			}
			v98 = l1 + (v92 - base.I32_wrap_i64(v93))
		}
		return v98
	}
}
func F_lpRandomEntries(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v21 = F_zmalloc_usable(m, l1<<(uint(int32(3))%32), v4)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v23 = F_lpLength(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_lpRandomEntries_0), int32(_a_F_lpRandomEntries_1), int32(1343))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L68
	}
L4:
	;
	if v23 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_qsort(m, v21, l1, int32(8), int32(560))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	v33 = v4
	goto L8
L8:
	;
	v42 = int32(0)
	v44 = *(*int64)(unsafe.Add(mBase, _c_F_lpRandomEntries[0]))
	v48 = v44*int64(6364136223846793005) + int64(1)
	*(*int64)(unsafe.Add(mBase, _c_F_lpRandomEntries[0])) = v48
	goto L10
L9:
	;
	goto L6
L10:
	;
	v55 = v21 + v33<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v33
	v57 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v48)>>(uint(int64(33))%64))), v23)
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v57
	v60 = v33 + int32(1)
	if v60 != l1 {
		v33 = v60
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v78 == int32(255) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	F__serverAssert(m, int32(_a_F_lpRandomEntries_2), int32(_a_F_lpRandomEntries_1), int32(400))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L67
	}
L14:
	;
	F__serverAssert(m, int32(_a_F_lpRandomEntries_3), int32(_a_F_lpRandomEntries_1), int32(395))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L66
	}
L15:
	;
	F__serverAssert(m, int32(_a_F_lpRandomEntries_2), int32(_a_F_lpRandomEntries_1), int32(426))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L65
	}
L16:
	;
	if l1 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v84 != int32(7) {
		goto L15
	} else {
		goto L19
	}
L18:
	;
	v87 = l0 + int32(6)
	goto L16
L19:
	;
	v87 = int32(0)
	goto L16
L20:
	;
	F_valkey_free(m, v21)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L64
	}
L21:
	;
	v90 = int32(0)
	v96 = v87
	v98 = v90
	v99 = v90
	goto L22
L22:
	;
	v106 = v21 + v98<<(uint(int32(3))%32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	if base.Ui32(v107) <= base.Ui32(v99) {
		v212 = v96
		v215 = v99
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L20
L24:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v223 = int32(0)
	v225 = F_lpGetWithSize(m, v212, v15+int32(8), v223, v223)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L56
	}
L25:
	;
	v113 = v96
	v116 = v99
	goto L26
L26:
	;
	if v113 == int32(0) {
		goto L14
	} else {
		goto L28
	}
L27:
	;
	v212 = v203
	v215 = v107
	goto L24
L28:
	;
	v126 = int32(1)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	v128 = base.I32_extend8_s(v127)
	if v128 <= int32(-1) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v206 = v116 + int32(1)
	if v206 != v107 {
		v113 = v203
		v116 = v206
		goto L26
	} else {
		goto L55
	}
L30:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	if v194 != int32(255) {
		v203 = v193
		goto L29
	} else {
		goto L53
	}
L31:
	;
	v193 = v113 + v191 + v190
	goto L30
L32:
	;
	if v127&int32(192) != int32(128) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v190 = v126
	v191 = int32(1)
	goto L31
L34:
	;
	if v127&int32(224) != int32(192) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v136 = int32(1)
	v190 = v136
	v191 = v127&int32(63) + v136
	goto L31
L36:
	;
	v149 = (v128 + int32(15)) & int32(255)
	if base.Ui32(v149) < base.Ui32(int32(4)) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v190 = v126
	v191 = int32(2)
	goto L31
L38:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v149<<(uint(int32(2))%32))+uint32(_c_F_lpRandomEntries[1])))
	v190 = v126
	v191 = v189
	goto L31
L39:
	;
	if v127&int32(240) != int32(224) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if base.Ui32(v171) < base.Ui32(int32(128)) {
		v190 = v126
		v191 = v171
		goto L31
	} else {
		goto L45
	}
L41:
	;
	switch v127 + int32(-240) {
	case 0:
		goto L43
	default:
		goto L44
	case 15:
		v190 = v126
		v191 = int32(1)
		goto L31
	}
L42:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	v171 = v156 | v127<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L40
L43:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v113)+1))
	v171 = v168 + int32(5)
	goto L40
L44:
	;
	v190 = v126
	v191 = int32(0)
	goto L31
L45:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v171) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v171) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v190 = int32(2)
	v191 = v171
	goto L31
L48:
	;
	if base.Ui32(v171) < base.Ui32(int32(268435456)) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v190 = int32(3)
	v191 = v171
	goto L31
L50:
	;
	v184 = int32(4)
	goto L52
L51:
	;
	v184 = int32(5)
	goto L52
L52:
	;
	v190 = v184
	v191 = v171
	goto L31
L53:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v193+int32(1) != l0+v200 {
		goto L13
	} else {
		goto L54
	}
L54:
	;
	v203 = int32(0)
	goto L29
L55:
	;
	goto L27
L56:
	;
	v229 = l2 + v220<<(uint(int32(4))%32)
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	if v225 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v232 = int64(0)
	goto L59
L58:
	;
	v232 = v231
	goto L59
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v229)+8)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = v225
	if v225 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v237 = base.I32_wrap_i64(v231)
	goto L62
L61:
	;
	v237 = int32(0)
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+4)) = v237
	v240 = v98 + int32(1)
	if v240 != l1 {
		v96 = v212
		v98 = v240
		v99 = v215
		goto L22
	} else {
		goto L63
	}
L63:
	;
	goto L23
L64:
	;
	m.G0 = v15 + int32(16)
	return
L65:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lpSafeToAdd(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	if l0 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v5 = v4
	} else {
		v5 = int32(0)
	}
	return base.B2i32(base.Ui32(v5+l1) < base.Ui32(int32(1073741825)))
}
func F_lpShrinkToFit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-8))))
	if base.Ui32(v6&int32(2147483647)) <= base.Ui32(v3) {
		v15 = l0
		return v15
	} else {
		v11 = F_zrealloc_usable(m, l0, v3, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			return v15
		}
	}
}
func F_lpValidateFirst(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v5 == int32(255) {
		v8 = int32(0)
	} else {
		v8 = l0 + int32(6)
	}
	return v8
}
func F_lpValidateIntegrityAndDups(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l2
	v16 = F_lpValidateIntegrity(m, l0, l1, int32(964), v8+int32(4))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		if v21 == int32(0) {
			m.G0 = v8 + int32(16)
			if v20&int32(1) != 0 {
				v32 = int32(0)
			} else {
				v32 = v16
			}
			if l2 != 0 {
				v33 = v32
			} else {
				v33 = v16
			}
			return v33
		} else {
			F_hashtableRelease(m, v21)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				if v20&int32(1) != 0 {
					v32 = int32(0)
				} else {
					v32 = v16
				}
				if l2 != 0 {
					v33 = v32
				} else {
					v33 = v16
				}
				return v33
			}
		}
	}
}
