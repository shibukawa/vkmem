package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_readToQueryBuf(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int64
	_ = v208
	var v211 int64
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
	if v9&int32(4) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v19 == v18 {
		v41 = v18
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v14 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v14
	return v14
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v44 = int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v45 == int32(2) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(-1)))))
	switch v24 & int32(7) {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L6
	case 4:
		goto L5
	default:
		v41 = v18
		goto L3
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(-17))))
	v41 = v40
	goto L3
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(-9))))
	v41 = v37
	goto L3
L7:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+int32(-5)))))
	v41 = v34
	goto L3
L8:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(-3)))))
	v41 = v31
	goto L3
L9:
	;
	v41 = int32(base.Ui32(v24) >> (uint(int32(3)) % 32))
	goto L3
L10:
	;
	if v84 != 0 {
		v137 = v41
		v138 = v84
		goto L32
	} else {
		goto L33
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v49 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v84 = v19
	v85 = v44
	v86 = int32(16384)
	goto L10
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v52 < int32(32768) {
		v84 = v19
		v85 = v44
		v86 = int32(16384)
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v84 = v19
	v85 = v44
	v86 = int32(16384)
	goto L10
L15:
	;
	v55 = int32(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v60 = v52 + (v56 - v41) + int32(2)
	if v55 < v60 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v64 = v60
	goto L18
L17:
	;
	v64 = int32(16384)
	goto L18
L18:
	;
	v65 = int32(16384)
	if base.Ui32(v65) < base.Ui32(v64) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v68 = v64
	goto L21
L20:
	;
	v68 = v65
	goto L21
L21:
	;
	v70 = int32(1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v71&v70 != 0 {
		v78 = v70
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v81 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v81 = v78
	goto L22
L24:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v74 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v76 = F_isImportSlotMigrationJob(m, v74)
	mBase = m.M
	v78 = v76
	goto L23
L26:
	;
	v81 = int32(0)
	goto L22
L27:
	;
	v82 = v68
	goto L29
L28:
	;
	v82 = v64
	goto L29
L29:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v84 = v83
	v85 = v55
	v86 = v82
	goto L10
L30:
	;
	F__serverAssert(m, int32(_a775), int32(_a774), int32(4360))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L44
	} else {
		goto L100
	}
L31:
	;
	F__serverAssert(m, int32(_a776), int32(_a774), int32(4334))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L44
	} else {
		goto L99
	}
L32:
	;
	v140 = int32(0)
	v142 = *(*int32)(unsafe.Add(mBase, _consts[407]))
	v144 = v43 & int32(16384)
	if v144 != 0 {
		goto L52
	} else {
		goto L53
	}
L33:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _consts[407]))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+int32(-1)))))
	switch v92 & int32(7) {
	case 0:
		goto L40
	case 1:
		goto L39
	case 2:
		goto L38
	case 3:
		goto L37
	case 4:
		goto L36
	default:
		goto L34
	}
L34:
	;
	if v85 != 0 {
		v115 = v89
		goto L42
	} else {
		goto L43
	}
L35:
	;
	if v109 != 0 {
		goto L31
	} else {
		goto L41
	}
L36:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v89+int32(-17))))
	v109 = v108
	goto L35
L37:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v89+int32(-9))))
	v109 = v105
	goto L35
L38:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89+int32(-5)))))
	v109 = v102
	goto L35
L39:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+int32(-3)))))
	v109 = v99
	goto L35
L40:
	;
	v109 = int32(base.Ui32(v92) >> (uint(int32(3)) % 32))
	goto L35
L41:
	;
	goto L34
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v115
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115+int32(-1)))))
	switch v120 & int32(7) {
	case 0:
		goto L50
	case 1:
		goto L49
	case 2:
		goto L48
	case 3:
		goto L47
	case 4:
		goto L46
	default:
		v137 = int32(0)
		v138 = v115
		goto L32
	}
L43:
	;
	v111 = F_sdsempty(m)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return int32(0)
L45:
	;
	v115 = v111
	goto L42
L46:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v115+int32(-17))))
	v137 = v136
	v138 = v115
	goto L32
L47:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v115+int32(-9))))
	v137 = v133
	v138 = v115
	goto L32
L48:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115+int32(-5)))))
	v137 = v130
	v138 = v115
	goto L32
L49:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115+int32(-3)))))
	v137 = v127
	v138 = v115
	goto L32
L50:
	;
	v137 = int32(base.Ui32(v120) >> (uint(int32(3)) % 32))
	v138 = v115
	goto L32
L51:
	;
	if v138 != v142 {
		goto L69
	} else {
		goto L70
	}
L52:
	;
	v175 = F_sdsMakeRoomFor(m, v138, v86)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L44
	} else {
		goto L64
	}
L53:
	;
	if v85 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v167 = F_sdsMakeRoomForNonGreedy(m, v138, v86)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L44
	} else {
		goto L61
	}
L55:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138+int32(-1)))))
	switch v149&int32(7) + int32(-2) {
	case 0:
		goto L59
	case 1:
		goto L58
	case 2:
		goto L57
	default:
		goto L54
	}
L56:
	;
	if base.Ui32(int32(16383)) < base.Ui32(v163) {
		goto L52
	} else {
		goto L60
	}
L57:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v138+int32(-9))))
	v163 = v162
	goto L56
L58:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v138+int32(-5))))
	v163 = v159
	goto L56
L59:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138+int32(-3)))))
	v163 = v156
	goto L56
L60:
	;
	goto L54
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v167
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+320))
	v171 = v137 + v86
	if base.Ui32(v170) < base.Ui32(v171) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+320)) = v171
	v214 = v167
	v215 = v86
	goto L51
L63:
	;
	v214 = v167
	v215 = v86
	goto L51
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v175
	v178 = int32(-1)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v178))))
	switch v180&int32(7) + v178 {
	case 0:
		goto L68
	case 1:
		goto L67
	case 2:
		goto L66
	case 3:
		goto L65
	default:
		v214 = v175
		v215 = v140
		goto L51
	}
L65:
	;
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v175+int32(-9))))
	v211 = *(*int64)(unsafe.Add(mBase, uint32(v175+int32(-17))))
	v214 = v175
	v215 = base.I32_wrap_i64(v208 - v211)
	goto L51
L66:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v175+int32(-5))))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v175+int32(-9))))
	v214 = v175
	v215 = v201 - v204
	goto L51
L67:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175+int32(-3)))))
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175+int32(-5)))))
	v214 = v175
	v215 = v194 - v197
	goto L51
L68:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+int32(-2)))))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+int32(-3)))))
	v214 = v175
	v215 = v187 - v190
	goto L51
L69:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+76))
	v225 = m.T0[v224].(func(*base.Module, int32, int32, int32) int32)(m, v221, v214+v137, v215)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L44
	} else {
		goto L72
	}
L70:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _consts[407]))
	if v214 != v219 {
		goto L30
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v225
	if v225 < int32(1) {
		v308 = int32(0)
		goto L73
	} else {
		goto L74
	}
L73:
	;
	return v308
L74:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_sdsIncrLen(m, v230, v225)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L44
	} else {
		goto L75
	}
L75:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v235 = v233 + int32(-1)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	switch v236 & int32(7) {
	case 0:
		goto L82
	case 1:
		goto L81
	case 2:
		goto L80
	case 3:
		goto L79
	case 4:
		goto L78
	default:
		goto L76
	}
L76:
	;
	if v144 != 0 {
		goto L84
	} else {
		goto L85
	}
L77:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+320))
	if base.Ui32(v253) <= base.Ui32(v254) {
		goto L76
	} else {
		goto L83
	}
L78:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v233+int32(-17))))
	v253 = v252
	goto L77
L79:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v233+int32(-9))))
	v253 = v249
	goto L77
L80:
	;
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v233+int32(-5)))))
	v253 = v246
	goto L77
L81:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233+int32(-3)))))
	v253 = v243
	goto L77
L82:
	;
	v253 = int32(base.Ui32(v236) >> (uint(int32(3)) % 32))
	goto L77
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+320)) = v253
	goto L76
L84:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v308 = base.B2i32(v303 == v215)
	goto L73
L85:
	;
	v258 = int32(0)
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	switch v260 & int32(7) {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	case 3:
		goto L88
	case 4:
		goto L87
	default:
		v277 = v258
		goto L86
	}
L86:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v278 == int32(0) {
		v282 = v258
		goto L92
	} else {
		goto L93
	}
L87:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v233+int32(-17))))
	v277 = v276
	goto L86
L88:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v233+int32(-9))))
	v277 = v273
	goto L86
L89:
	;
	v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v233+int32(-5)))))
	v277 = v270
	goto L86
L90:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233+int32(-3)))))
	v277 = v267
	goto L86
L91:
	;
	v277 = int32(base.Ui32(v260) >> (uint(int32(3)) % 32))
	goto L86
L92:
	;
	v283 = v282 + v277
	v285 = *(*int32)(unsafe.Add(mBase, _consts[408]))
	if base.Ui32(v283) <= base.Ui32(v285) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v278)+16))
	v282 = v281
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v295 | int32(1)
	goto L84
L95:
	;
	if base.Ui32(v283) < base.Ui32(int32(1048577)) {
		goto L84
	} else {
		goto L97
	}
L96:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v295 = v287
	goto L94
L97:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	if v290&int32(65536) == int32(0) {
		goto L84
	} else {
		goto L98
	}
L98:
	;
	v295 = v290
	goto L94
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
func F_trySendWriteToIOThreads(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int64
	_ = v254
	var v258 int32
	_ = v258
	var v269 int32
	_ = v269
	v9 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	if int32(2) <= v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return int32(-1)
L3:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v18 != 0 {
		v258 = int32(0)
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(-1)
L5:
	;
	F__serverAssert(m, int32(_a666), int32(_a663), int32(591))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L11
	} else {
		goto L83
	}
L6:
	;
	return v258
L7:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v19 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v24 = F_clientHasPendingReplies(m, l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return int32(-1)
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v30&int32(1) != 0 {
		v56 = v30
		goto L14
	} else {
		goto L15
	}
L11:
	;
	return int32(0)
L12:
	;
	if v24 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	return int32(-1)
L14:
	;
	if v56&int32(268435456) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	if v30&int32(2) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v51 == int32(9) {
		v56 = v30
		goto L14
	} else {
		goto L23
	}
L17:
	;
	if v30&int32(262144) != 0 {
		v56 = v30
		goto L14
	} else {
		goto L20
	}
L18:
	;
	if v30&int32(4) == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v43 == int32(0) {
		v56 = v30
		goto L14
	} else {
		goto L21
	}
L21:
	;
	goto L22
L22:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v56 = v49
	goto L14
L23:
	;
	return int32(-1)
L24:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v64 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	return int32(-1)
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v81&int32(1) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v68 = int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v69 == int32(0) {
		v76 = v68
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v76 != 0 {
		goto L26
	} else {
		goto L32
	}
L29:
	;
	goto L28
L30:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v72 != 0 {
		v76 = v68
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+156))
	v76 = base.B2i32(v73 != int32(13))
	goto L29
L32:
	;
	return int32(-1)
L33:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v133 != 0 {
		goto L52
	} else {
		goto L53
	}
L34:
	;
	v129 = v127
	v130 = int32(0)
	goto L33
L35:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v118
	if v118 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L36:
	;
	if v107 != int32(1) {
		goto L35
	} else {
		goto L49
	}
L37:
	;
	v87 = int32(2)
	if v81&v87 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v107 = int32(3)
	goto L36
L39:
	;
	if v81&int32(262144) != 0 {
		v104 = v87
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if v81&int32(4) != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v107 = int32(1)
	goto L36
L42:
	;
	v107 = v104
	goto L36
L43:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v97 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v101 = F_isImportSlotMigrationJob(m, v97)
	mBase = m.M
	if v101 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v107 = int32(0)
	goto L36
L46:
	;
	v102 = int32(4)
	goto L48
L47:
	;
	v102 = int32(5)
	goto L48
L48:
	;
	v104 = v102
	goto L42
L49:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v127 = v114 + int32(28)
	goto L34
L50:
	;
	v127 = l0 + int32(180)
	goto L34
L51:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v129 = v122 + int32(4)
	v130 = v122
	goto L33
L52:
	;
	v138 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)) = uint8(v138)
	v143 = base.B2i32(v107 == v138) << (uint(v138) % 32)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+220)) = uint16(v143)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v147 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	if v107 == int32(1) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	if v131 == int32(0) {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	v165 = int32(_a476)
	v171 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	v172 = *(*int32)(unsafe.Add(mBase, _consts[312]))
	v175 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	v179 = v171 + (v172+int32(-1))&v175<<(uint(int32(6))%32)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	if v180 != v175 {
		goto L66
	} else {
		goto L67
	}
L57:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if v150 == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)+112))
	if v153 == int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v157 = base.B2i32(v145 != int32(0))
	if v146 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v160 = v157 | int32(2)
	goto L62
L61:
	;
	v160 = v157
	goto L62
L62:
	;
	m.T0[v153].(func(*base.Module, int32, int32))(m, v147, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	goto L56
L64:
	;
	if v107 == int32(1) {
		goto L74
	} else {
		goto L75
	}
L65:
	;
	if v180 == v175 {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = l0 | int32(1)
	v183 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = v180 + v183
	v186 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	*(*int32)(unsafe.Add(mBase, _consts[315])) = v186 + v183
	goto L66
L68:
	;
	v191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)) = uint8(v191)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v193 == v191 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = int64(0)
	v208 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+220)) = uint16(v208)
	return int32(-1)
L70:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if v196 == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v196)+112))
	if v199 == int32(0) {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	m.T0[v199].(func(*base.Module, int32, int32))(m, v193, int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+202)))
	if v230&int32(64) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L75:
	;
	if v130 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
	if v223&int32(1) == int32(0) {
		goto L74
	} else {
		goto L79
	}
L77:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+12)))
	if v216&int32(1) == int32(0) {
		goto L74
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+8)) = int32(0)
	goto L74
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(0)
	goto L74
L80:
	;
	v245 = int32(0)
	v248 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	*(*int32)(unsafe.Add(mBase, _consts[220])) = v248 + int32(1)
	v252 = int32(_a69)
	v254 = *(*int64)(unsafe.Add(mBase, _consts[316]))
	*(*int64)(unsafe.Add(mBase, _consts[316])) = v254 + int64(1)
	v258 = v245
	goto L6
L81:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	F_listUnlinkNode(m, v236, l0+int32(168))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v241 & int32(-4194305)
	goto L80
L83:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_writeToReplica(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	v16 = m.G0
	v18 = v16 - int32(8192)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a804), int32(_a774), int32(2495))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L14
	} else {
		goto L55
	}
L2:
	;
	F__serverAssert(m, int32(_a805), int32(_a774), int32(2487))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L14
	} else {
		goto L54
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	if v22 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	goto L8
L5:
	;
	m.G0 = v18 + int32(8192)
	return
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+184))
	if v46 != v42 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v37 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L8:
	;
	if base.B2i32(v24 == v23) == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v31 == int32(0) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v42 = v31
	v43 = v34 + int32(28)
	goto L6
L11:
	;
	v42 = v37
	v43 = l0 + int32(140)
	goto L6
L12:
	;
	if v46 == int32(0) {
		goto L5
	} else {
		goto L17
	}
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+188))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+68))
	v57 = m.T0[v56].(func(*base.Module, int32, int32, int32) int32)(m, v48, v49+v50+int32(32), v44-v50)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v57
	if int32(0) < v57 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+220)))
	v64 = v62 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+220)) = uint16(v64)
	goto L5
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+20)))
	if v69 == int32(0) {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v72 = int32(1024)
	if base.Ui32(v69) < base.Ui32(v72) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v75 = v69
	goto L21
L20:
	;
	v75 = v72
	goto L21
L21:
	;
	v76 = int32(0)
	v81 = v46
	v87 = v76
	v88 = v76
	goto L22
L22:
	;
	if v81 != v46 {
		v96 = int32(0)
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v124 == int32(0) {
		goto L5
	} else {
		goto L34
	}
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v98 = base.B2i32(v81 == v42)
	if v81 == v42 {
		v100 = v44
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v45)+188))
	v96 = v95
	goto L24
L26:
	;
	if v100 == v96 {
		v114 = v87
		v115 = v88
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	v100 = v99
	goto L26
L28:
	;
	goto L23
L29:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v118 == int32(0) {
		v123 = v114
		v124 = v115
		goto L28
	} else {
		goto L32
	}
L30:
	;
	v104 = v18 + v87<<(uint(int32(3))%32)
	v105 = v100 - v96
	*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v97 + v96 + int32(32)
	v112 = v87 + int32(1)
	v113 = v105 + v88
	if v81 == v42 {
		v123 = v112
		v124 = v113
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v114 = v112
	v115 = v113
	goto L29
L32:
	;
	if v114 < v75 {
		v81 = v118
		v87 = v114
		v88 = v115
		goto L22
	} else {
		goto L33
	}
L33:
	;
	v123 = v114
	v124 = v115
	goto L28
L34:
	;
	v129 = int32(0)
	if v123 < int32(1) {
		v206 = v129
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v206
	goto L5
L36:
	;
	v139 = v18
	v141 = v123
	v143 = v129
	goto L37
L37:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+72))
	v150 = m.T0[v149].(func(*base.Module, int32, int32, int32) int32)(m, v147, v139, v141)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L14
	} else {
		goto L40
	}
L38:
	;
	v206 = v161
	goto L35
L39:
	;
	v161 = v150 + v143
	if v161 != v124 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	if int32(0) < v150 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+220)))
	v156 = v154 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+220)) = uint16(v156)
	if int32(0) < v143 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v160 = v143
	goto L44
L43:
	;
	v160 = v150
	goto L44
L44:
	;
	v206 = v160
	goto L35
L45:
	;
	v166 = v150
	v170 = v139
	v172 = v141
	goto L48
L46:
	;
	v206 = v124
	goto L35
L47:
	;
	if int32(0) < v192 {
		v139 = v191
		v141 = v192
		v143 = v161
		goto L37
	} else {
		goto L53
	}
L48:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if base.Ui32(v178) <= base.Ui32(v166) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v191 = v188
	v192 = v186
	goto L47
L50:
	;
	v186 = v172 + int32(-1)
	v188 = v170 + int32(8)
	v189 = v166 - v178
	if v189 != 0 {
		v166 = v189
		v170 = v188
		v172 = v186
		goto L48
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = v178 - v166
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v182 + v166
	v191 = v170
	v192 = v172
	goto L47
L52:
	;
	goto L49
L53:
	;
	goto L38
L54:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
