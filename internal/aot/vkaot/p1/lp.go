package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_lpAppend(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v4 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = F_lpInsert(m, l0, l1, v4, l2, l0+v5+int32(-1), v4, v4)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_lpBatchDelete(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	if l2 == int32(0) {
		v298 = l0
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F__serverAssert(m, int32(_a732), int32(_a729), int32(1022))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L87
	} else {
		goto L92
	}
L2:
	;
	F__serverAssert(m, int32(_a733), int32(_a729), int32(1014))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L87
	} else {
		goto L91
	}
L3:
	;
	F__serverAssert(m, int32(_a734), int32(_a729), int32(1003))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L87
	} else {
		goto L90
	}
L4:
	;
	F__serverAssert(m, int32(_a735), int32(_a729), int32(988))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L87
	} else {
		goto L89
	}
L5:
	;
	return v298
L6:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = l0 + v12
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
	if v16 != int32(255) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = v19
	v27 = int32(0)
	goto L8
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1+v27<<(uint(int32(2))%32))))
	if v33 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v274 = v270 - v13 + v12
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v274+int32(-1)))))
	if v278 != int32(255) {
		goto L1
	} else {
		goto L82
	}
L10:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v36 == int32(255) {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v42 = int32(1)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v44 = base.I32_extend8_s(v43)
	if v44 <= int32(-1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v111 = v27 + int32(1)
	if base.Ui32(l2) <= base.Ui32(v111) {
		v118 = v13
		goto L36
	} else {
		goto L37
	}
L13:
	;
	v109 = v33 + v107 + v106
	goto L12
L14:
	;
	if v43&int32(192) != int32(128) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v106 = v42
	v107 = int32(1)
	goto L13
L16:
	;
	if v43&int32(224) != int32(192) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v52 = int32(1)
	v106 = v52
	v107 = v43&int32(63) + v52
	goto L13
L18:
	;
	v65 = (v44 + int32(15)) & int32(255)
	if base.Ui32(v65) < base.Ui32(int32(4)) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v106 = v42
	v107 = int32(2)
	goto L13
L20:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v65<<(uint(int32(2))%32))+uint32(_consts[375])))
	v106 = v42
	v107 = v105
	goto L13
L21:
	;
	if v43&int32(240) != int32(224) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if base.Ui32(v87) < base.Ui32(int32(128)) {
		v106 = v42
		v107 = v87
		goto L13
	} else {
		goto L27
	}
L23:
	;
	switch v43 + int32(-240) {
	case 0:
		goto L25
	default:
		goto L26
	case 15:
		v106 = v42
		v107 = int32(1)
		goto L13
	}
L24:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	v87 = v72 | v43<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L22
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v33)+1))
	v87 = v84 + int32(5)
	goto L22
L26:
	;
	v106 = v42
	v107 = int32(0)
	goto L13
L27:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v87) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v87) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v106 = int32(2)
	v107 = v87
	goto L13
L30:
	;
	if base.Ui32(v87) < base.Ui32(int32(268435456)) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v106 = int32(3)
	v107 = v87
	goto L13
L32:
	;
	v100 = int32(4)
	goto L34
L33:
	;
	v100 = int32(5)
	goto L34
L34:
	;
	v106 = v100
	v107 = v87
	goto L13
L35:
	;
	if v111 != l2 {
		v26 = v270
		v27 = v111
		goto L8
	} else {
		goto L81
	}
L36:
	;
	if base.Ui32(v118) <= base.Ui32(v109) {
		goto L2
	} else {
		goto L39
	}
L37:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1+v111<<(uint(int32(2))%32))))
	if v109 == v116 {
		v270 = v26
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v118 = v116
	goto L36
L39:
	;
	v120 = v118 - v109
	if v26 == v109 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v270 = v268 + v120
	goto L35
L41:
	;
	v268 = v26
	goto L40
L42:
	;
	v124 = v120 + v26
	if base.Ui32(int32(0)-v120<<(uint(int32(1))%32)) < base.Ui32(v109-v124) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v134 = (v109 ^ v26) & int32(3)
	if base.Ui32(v109) <= base.Ui32(v26) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v131 = F___memcpy(m, v26, v109, v120)
	mBase = m.M
	v268 = v131
	goto L40
L45:
	;
	if v240 == int32(0) {
		goto L41
	} else {
		goto L77
	}
L46:
	;
	if base.Ui32(v218) <= base.Ui32(int32(3)) {
		v239 = v217
		v240 = v218
		v241 = v219
		goto L45
	} else {
		goto L73
	}
L47:
	;
	if v134 != 0 {
		v200 = v120
		goto L57
	} else {
		goto L58
	}
L48:
	;
	if v134 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v26&int32(3) != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v239 = v109
	v240 = v120
	v241 = v26
	goto L45
L51:
	;
	v141 = v109
	v142 = v120
	v143 = v26
	goto L53
L52:
	;
	v217 = v109
	v218 = v120
	v219 = v26
	goto L46
L53:
	;
	if v142 == int32(0) {
		goto L41
	} else {
		goto L55
	}
L55:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v147)
	v149 = int32(1)
	v150 = v141 + v149
	v152 = v142 + int32(-1)
	v154 = v143 + v149
	if v154&int32(3) == int32(0) {
		v217 = v150
		v218 = v152
		v219 = v154
		goto L46
	} else {
		goto L56
	}
L56:
	;
	v141 = v150
	v142 = v152
	v143 = v154
	goto L53
L57:
	;
	if v200 == int32(0) {
		goto L41
	} else {
		goto L69
	}
L58:
	;
	if v124&int32(3) == int32(0) {
		v180 = v120
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if base.Ui32(v180) <= base.Ui32(int32(3)) {
		v200 = v180
		goto L57
	} else {
		goto L65
	}
L60:
	;
	v165 = v120
	goto L61
L61:
	;
	if v165 == int32(0) {
		goto L41
	} else {
		goto L63
	}
L62:
	;
	v180 = v171
	goto L59
L63:
	;
	v171 = v165 + int32(-1)
	v172 = v26 + v171
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v171))))
	*(*uint8)(unsafe.Add(mBase, uint32(v172))) = uint8(v174)
	if v172&int32(3) != 0 {
		v165 = v171
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v187 = v180
	goto L66
L66:
	;
	v191 = v187 + int32(-4)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v109+v191)))
	*(*int32)(unsafe.Add(mBase, uint32(v26+v191))) = v194
	if base.Ui32(int32(3)) < base.Ui32(v191) {
		v187 = v191
		goto L66
	} else {
		goto L68
	}
L67:
	;
	v200 = v191
	goto L57
L68:
	;
	goto L67
L69:
	;
	v207 = v200
	goto L70
L70:
	;
	v211 = v207 + int32(-1)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v211))))
	*(*uint8)(unsafe.Add(mBase, uint32(v26+v211))) = uint8(v214)
	if v211 != 0 {
		v207 = v211
		goto L70
	} else {
		goto L72
	}
L72:
	;
	goto L41
L73:
	;
	v224 = v217
	v225 = v218
	v226 = v219
	goto L74
L74:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	*(*int32)(unsafe.Add(mBase, uint32(v226))) = v228
	v230 = int32(4)
	v231 = v224 + v230
	v233 = v226 + v230
	v235 = v225 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v235) {
		v224 = v231
		v225 = v235
		v226 = v233
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v239 = v231
	v240 = v235
	v241 = v233
	goto L45
L76:
	;
	goto L75
L77:
	;
	v246 = v239
	v247 = v240
	v248 = v241
	goto L78
L78:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	*(*uint8)(unsafe.Add(mBase, uint32(v248))) = uint8(v250)
	v252 = int32(1)
	v257 = v247 + int32(-1)
	if v257 != 0 {
		v246 = v246 + v252
		v247 = v257
		v248 = v248 + v252
		goto L78
	} else {
		goto L80
	}
L79:
	;
	goto L41
L80:
	;
	goto L79
L81:
	;
	goto L9
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v274
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v282 == int32(65535) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-8))))
	goto L85
L84:
	;
	v285 = v282 - l2
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v285)
	goto L83
L85:
	;
	if base.Ui32(v289&int32(2147483647)) <= base.Ui32(v274) {
		v298 = l0
		goto L5
	} else {
		goto L86
	}
L86:
	;
	v294 = F_zrealloc_usable(m, l0, v274, int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	return int32(0)
L88:
	;
	v298 = v294
	goto L5
L89:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lpBytes(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return v2
}
func F_lpCompare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v141 int64
	_ = v141
	var v147 int32
	_ = v147
	var v149 int64
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v163 int64
	_ = v163
	var v168 int64
	_ = v168
	var v172 int32
	_ = v172
	var v174 int64
	_ = v174
	var v176 int32
	_ = v176
	var v184 int64
	_ = v184
	var v208 int64
	_ = v208
	var v227 int32
	_ = v227
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v241 int32
	_ = v241
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 == int32(255) {
		v241 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v241
L2:
	;
	v16 = int32(0)
	v18 = F_lpGetWithSize(m, l0, v8+int32(8), v16, v16)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v93 = int32(0)
	if base.Ui32(l2+int32(-21)) < base.Ui32(int32(-20)) {
		v227 = v93
		goto L25
	} else {
		goto L26
	}
L4:
	;
	return int32(0)
L5:
	;
	if v18 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
	if v24 != base.I64_extend_i32_u(l2) {
		v241 = v4
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if base.Ui32(l2) < base.Ui32(int32(4)) {
		v50 = v18
		v51 = l1
		v52 = l2
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v241 = base.B2i32(v90 == int32(0))
	goto L1
L9:
	;
	v90 = int32(0)
	goto L8
L10:
	;
	v62 = v57
	v63 = v58
	v64 = v59
	goto L20
L11:
	;
	if v52 == int32(0) {
		goto L9
	} else {
		goto L18
	}
L12:
	;
	if (l1|v18)&int32(3) != 0 {
		v57 = v18
		v58 = l1
		v59 = l2
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v34 = v18
	v35 = l1
	v36 = l2
	goto L14
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v39 != v40 {
		v57 = v34
		v58 = v35
		v59 = v36
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v50 = v45
	v51 = v43
	v52 = v47
	goto L11
L16:
	;
	v42 = int32(4)
	v43 = v35 + v42
	v45 = v34 + v42
	v47 = v36 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v47) {
		v34 = v45
		v35 = v43
		v36 = v47
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v57 = v50
	v58 = v51
	v59 = v52
	goto L10
L19:
	;
	v90 = v67 - v68
	goto L8
L20:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v67 != v68 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v70 = int32(1)
	v75 = v64 + int32(-1)
	if v75 == int32(0) {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v62 = v62 + v70
	v63 = v63 + v70
	v64 = v75
	goto L20
L24:
	;
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v241 = base.B2i32(v227 != int32(0)) & base.B2i32(v236 == v237)
	goto L1
L25:
	;
	goto L24
L26:
	;
	v105 = int32(1)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if l2 != v105 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v227 = int32(1)
	goto L25
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v208
	goto L27
L29:
	;
	if v106&int32(255) == int32(45) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v110 = v106 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v110&int32(255)) {
		v227 = v93
		goto L25
	} else {
		goto L31
	}
L31:
	;
	if v8 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v208 = base.I64_extend_i32_u(v110) & int64(255)
	goto L28
L33:
	;
	if base.Ui32(int32(8)) < base.Ui32((v129+int32(-49))&int32(255)) {
		v227 = v93
		goto L25
	} else {
		goto L36
	}
L34:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v128 = int32(2)
	v129 = v126
	v130 = l1 + int32(1)
	goto L33
L35:
	;
	v128 = v105
	v129 = v106
	v130 = l1
	goto L33
L36:
	;
	v141 = base.I64_extend_i32_u(v129+int32(-48)) & int64(255)
	if base.Ui32(l2) <= base.Ui32(v128) {
		v184 = v141
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if v106&int32(255) != int32(45) {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v147 = v128
	v149 = v141
	v151 = v130
	goto L39
L39:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	if base.Ui32((v153+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v227 = v93
		goto L25
	} else {
		goto L41
	}
L40:
	;
	v184 = v174
	goto L37
L41:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v149) {
		v227 = v93
		goto L25
	} else {
		goto L42
	}
L42:
	;
	v163 = v149 * int64(10)
	v168 = base.I64_extend_i32_u(v153+int32(-48)) & int64(255)
	if base.Ui64(v168^int64(-1)) < base.Ui64(v163) {
		v227 = v93
		goto L25
	} else {
		goto L43
	}
L43:
	;
	v172 = int32(1)
	v174 = v163 + v168
	v176 = v147 + v172
	if v176 != l2 {
		v147 = v176
		v149 = v174
		v151 = v151 + v172
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	if v184 < int64(0) {
		v227 = v93
		goto L25
	} else {
		goto L49
	}
L46:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v184) {
		v227 = v93
		goto L25
	} else {
		goto L47
	}
L47:
	;
	if v8 == int32(0) {
		goto L27
	} else {
		goto L48
	}
L48:
	;
	v208 = int64(0) - v184
	goto L28
L49:
	;
	if v8 == int32(0) {
		goto L27
	} else {
		goto L50
	}
L50:
	;
	v208 = v184
	goto L28
}
func F_lpDeleteRangeWithEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v308 int32
	_ = v308
	if l2 == int32(0) {
		v293 = l0
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a731), int32(_a729), int32(929))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L80
	} else {
		goto L85
	}
L2:
	;
	return v293
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = l0 + v13
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v25 = int32(0)
	v26 = l2
	v28 = v18
	goto L5
L4:
	;
	v115 = v14 + int32(-1) - v112 + int32(1)
	if v18 == v112 {
		goto L35
	} else {
		goto L36
	}
L5:
	;
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v103+int32(1) != v14 {
		goto L1
	} else {
		goto L33
	}
L7:
	;
	v29 = int32(1)
	v30 = v25 + v29
	v31 = int32(-1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v38 = base.I32_extend8_s(v37)
	if v38 <= v31 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v110 = l2
	v112 = v28
	goto L4
L9:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v104 != int32(255) {
		v25 = v30
		v26 = v26 + v31
		v28 = v103
		goto L5
	} else {
		goto L32
	}
L10:
	;
	v103 = v28 + v101 + v100
	goto L9
L11:
	;
	if v37&int32(192) != int32(128) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v100 = v29
	v101 = int32(1)
	goto L10
L13:
	;
	if v37&int32(224) != int32(192) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v46 = int32(1)
	v100 = v46
	v101 = v37&int32(63) + v46
	goto L10
L15:
	;
	v59 = (v38 + int32(15)) & int32(255)
	if base.Ui32(v59) < base.Ui32(int32(4)) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v100 = v29
	v101 = int32(2)
	goto L10
L17:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v59<<(uint(int32(2))%32))+uint32(_consts[375])))
	v100 = v29
	v101 = v99
	goto L10
L18:
	;
	if v37&int32(240) != int32(224) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if base.Ui32(v81) < base.Ui32(int32(128)) {
		v100 = v29
		v101 = v81
		goto L10
	} else {
		goto L24
	}
L20:
	;
	switch v37 + int32(-240) {
	case 0:
		goto L22
	default:
		goto L23
	case 15:
		v100 = v29
		v101 = int32(1)
		goto L10
	}
L21:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v81 = v66 | v37<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L19
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v28)+1))
	v81 = v78 + int32(5)
	goto L19
L23:
	;
	v100 = v29
	v101 = int32(0)
	goto L10
L24:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v81) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v81) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v100 = int32(2)
	v101 = v81
	goto L10
L27:
	;
	if base.Ui32(v81) < base.Ui32(int32(268435456)) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v100 = int32(3)
	v101 = v81
	goto L10
L29:
	;
	v94 = int32(4)
	goto L31
L30:
	;
	v94 = int32(5)
	goto L31
L31:
	;
	v100 = v94
	v101 = v81
	goto L10
L32:
	;
	goto L6
L33:
	;
	v110 = v30
	v112 = v103
	goto L4
L34:
	;
	v265 = v263 - v112 + v13
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v265
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v267 == int32(65535) {
		goto L75
	} else {
		goto L76
	}
L35:
	;
	v263 = v18
	goto L34
L36:
	;
	v119 = v115 + v18
	if base.Ui32(int32(0)-v115<<(uint(int32(1))%32)) < base.Ui32(v112-v119) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v129 = (v112 ^ v18) & int32(3)
	if base.Ui32(v112) <= base.Ui32(v18) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v126 = F___memcpy(m, v18, v112, v115)
	mBase = m.M
	v263 = v126
	goto L34
L39:
	;
	if v235 == int32(0) {
		goto L35
	} else {
		goto L71
	}
L40:
	;
	if base.Ui32(v213) <= base.Ui32(int32(3)) {
		v234 = v212
		v235 = v213
		v236 = v214
		goto L39
	} else {
		goto L67
	}
L41:
	;
	if v129 != 0 {
		v195 = v115
		goto L51
	} else {
		goto L52
	}
L42:
	;
	if v129 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v18&int32(3) != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v234 = v112
	v235 = v115
	v236 = v18
	goto L39
L45:
	;
	v136 = v112
	v137 = v115
	v138 = v18
	goto L47
L46:
	;
	v212 = v112
	v213 = v115
	v214 = v18
	goto L40
L47:
	;
	if v137 == int32(0) {
		goto L35
	} else {
		goto L49
	}
L49:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v142)
	v144 = int32(1)
	v145 = v136 + v144
	v147 = v137 + int32(-1)
	v149 = v138 + v144
	if v149&int32(3) == int32(0) {
		v212 = v145
		v213 = v147
		v214 = v149
		goto L40
	} else {
		goto L50
	}
L50:
	;
	v136 = v145
	v137 = v147
	v138 = v149
	goto L47
L51:
	;
	if v195 == int32(0) {
		goto L35
	} else {
		goto L63
	}
L52:
	;
	if v119&int32(3) == int32(0) {
		v175 = v115
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if base.Ui32(v175) <= base.Ui32(int32(3)) {
		v195 = v175
		goto L51
	} else {
		goto L59
	}
L54:
	;
	v160 = v115
	goto L55
L55:
	;
	if v160 == int32(0) {
		goto L35
	} else {
		goto L57
	}
L56:
	;
	v175 = v166
	goto L53
L57:
	;
	v166 = v160 + int32(-1)
	v167 = v18 + v166
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+v166))))
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v169)
	if v167&int32(3) != 0 {
		v160 = v166
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v182 = v175
	goto L60
L60:
	;
	v186 = v182 + int32(-4)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v112+v186)))
	*(*int32)(unsafe.Add(mBase, uint32(v18+v186))) = v189
	if base.Ui32(int32(3)) < base.Ui32(v186) {
		v182 = v186
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v195 = v186
	goto L51
L62:
	;
	goto L61
L63:
	;
	v202 = v195
	goto L64
L64:
	;
	v206 = v202 + int32(-1)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+v206))))
	*(*uint8)(unsafe.Add(mBase, uint32(v18+v206))) = uint8(v209)
	if v206 != 0 {
		v202 = v206
		goto L64
	} else {
		goto L66
	}
L66:
	;
	goto L35
L67:
	;
	v219 = v212
	v220 = v213
	v221 = v214
	goto L68
L68:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = v223
	v225 = int32(4)
	v226 = v219 + v225
	v228 = v221 + v225
	v230 = v220 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v230) {
		v219 = v226
		v220 = v230
		v221 = v228
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v234 = v226
	v235 = v230
	v236 = v228
	goto L39
L70:
	;
	goto L69
L71:
	;
	v241 = v234
	v242 = v235
	v243 = v236
	goto L72
L72:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	*(*uint8)(unsafe.Add(mBase, uint32(v243))) = uint8(v245)
	v247 = int32(1)
	v252 = v242 + int32(-1)
	if v252 != 0 {
		v241 = v241 + v247
		v242 = v252
		v243 = v243 + v247
		goto L72
	} else {
		goto L74
	}
L73:
	;
	goto L35
L74:
	;
	goto L73
L75:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-8))))
	goto L78
L76:
	;
	v270 = v267 - v110
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v270)
	goto L75
L77:
	;
	v285 = v284 + (v263 - l0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v285
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	if v288 == int32(255) {
		goto L82
	} else {
		goto L83
	}
L78:
	;
	if base.Ui32(v275&int32(2147483647)) <= base.Ui32(v265) {
		v284 = l0
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v280 = F_zrealloc_usable(m, l0, v265, int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	return int32(0)
L81:
	;
	v284 = v280
	goto L77
L82:
	;
	v291 = int32(0)
	goto L84
L83:
	;
	v291 = v285
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v291
	v293 = v284
	goto L2
L85:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lpFreeVoid(m *base.Module, l0 int32) {
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
func F_lpGetEdgeStreamID(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v81 int64
	_ = v81
	var v87 int32
	_ = v87
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v103 int64
	_ = v103
	var v108 int64
	_ = v108
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v124 int64
	_ = v124
	var v148 int64
	_ = v148
	var v167 int32
	_ = v167
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v190 int64
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int64
	_ = v196
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v228 int64
	_ = v228
	var v233 int32
	_ = v233
	var v239 int64
	_ = v239
	var v244 int32
	_ = v244
	var v250 int64
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v262 int64
	_ = v262
	var v267 int32
	_ = v267
	var v279 int64
	_ = v279
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int64
	_ = v298
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int64
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v362 int64
	_ = v362
	var v368 int32
	_ = v368
	var v370 int64
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v384 int64
	_ = v384
	var v389 int64
	_ = v389
	var v393 int32
	_ = v393
	var v395 int64
	_ = v395
	var v397 int32
	_ = v397
	var v405 int64
	_ = v405
	var v429 int64
	_ = v429
	var v448 int32
	_ = v448
	var v457 int64
	_ = v457
	var v458 int64
	_ = v458
	var v465 int32
	_ = v465
	var v471 int64
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int64
	_ = v477
	var v481 int32
	_ = v481
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int64
	_ = v492
	var v493 int64
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int64
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v549 int64
	_ = v549
	var v555 int32
	_ = v555
	var v557 int64
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v571 int64
	_ = v571
	var v576 int64
	_ = v576
	var v580 int32
	_ = v580
	var v582 int64
	_ = v582
	var v584 int32
	_ = v584
	var v592 int64
	_ = v592
	var v616 int64
	_ = v616
	var v635 int32
	_ = v635
	var v644 int64
	_ = v644
	var v645 int64
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int64
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v703 int64
	_ = v703
	var v709 int32
	_ = v709
	var v711 int64
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v725 int64
	_ = v725
	var v730 int64
	_ = v730
	var v734 int32
	_ = v734
	var v736 int64
	_ = v736
	var v738 int32
	_ = v738
	var v746 int64
	_ = v746
	var v770 int64
	_ = v770
	var v789 int32
	_ = v789
	var v798 int64
	_ = v798
	var v799 int64
	_ = v799
	var v805 int32
	_ = v805
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v841 int32
	_ = v841
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F__serverAssert(m, int32(_a2401), int32(_a2402), int32(282))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L13
	} else {
		goto L183
	}
L2:
	;
	F__serverAssert(m, int32(_a2401), int32(_a2402), int32(282))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L13
	} else {
		goto L182
	}
L3:
	;
	F__serverAssert(m, int32(_a2401), int32(_a2402), int32(282))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L13
	} else {
		goto L181
	}
L4:
	;
	F__serverAssert(m, int32(_a2401), int32(_a2402), int32(282))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L13
	} else {
		goto L180
	}
L5:
	;
	m.G0 = v13 + int32(16)
	return v805
L6:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v805 = int32(0)
	goto L5
L8:
	;
	v490 = F_lpNext(m, l0, v481)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L13
	} else {
		goto L114
	}
L9:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v216 == int32(7) {
		v303 = int32(0)
		goto L59
	} else {
		goto L60
	}
L10:
	;
	v18 = F_lpFirst(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v179 = F_lpNext(m, l0, v24)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L13
	} else {
		goto L47
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v33 = int32(0)
	if base.Ui32(v32+int32(-21)) < base.Ui32(int32(-20)) {
		v167 = v33
		goto L20
	} else {
		goto L21
	}
L13:
	;
	return int32(0)
L14:
	;
	v22 = F_lpNext(m, l0, v18)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v24 = F_lpNext(m, l0, v22)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v29 = F_lpGet(m, v24, v13+int32(8), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	if v29 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v177 = v31
	goto L11
L19:
	;
	if v167 == int32(0) {
		goto L4
	} else {
		goto L46
	}
L20:
	;
	goto L19
L21:
	;
	v45 = int32(1)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v32 != v45 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v167 = int32(1)
	goto L20
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v148
	goto L22
L24:
	;
	if v46&int32(255) == int32(45) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v50 = v46 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v50&int32(255)) {
		v167 = v33
		goto L20
	} else {
		goto L26
	}
L26:
	;
	if v13 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v148 = base.I64_extend_i32_u(v50) & int64(255)
	goto L23
L28:
	;
	if base.Ui32(int32(8)) < base.Ui32((v69+int32(-49))&int32(255)) {
		v167 = v33
		goto L20
	} else {
		goto L31
	}
L29:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	v68 = int32(2)
	v69 = v66
	v70 = v29 + int32(1)
	goto L28
L30:
	;
	v68 = v45
	v69 = v46
	v70 = v29
	goto L28
L31:
	;
	v81 = base.I64_extend_i32_u(v69+int32(-48)) & int64(255)
	if base.Ui32(v32) <= base.Ui32(v68) {
		v124 = v81
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v46&int32(255) != int32(45) {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	v87 = v68
	v89 = v81
	v91 = v70
	goto L34
L34:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if base.Ui32((v93+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v167 = v33
		goto L20
	} else {
		goto L36
	}
L35:
	;
	v124 = v114
	goto L32
L36:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v89) {
		v167 = v33
		goto L20
	} else {
		goto L37
	}
L37:
	;
	v103 = v89 * int64(10)
	v108 = base.I64_extend_i32_u(v93+int32(-48)) & int64(255)
	if base.Ui64(v108^int64(-1)) < base.Ui64(v103) {
		v167 = v33
		goto L20
	} else {
		goto L38
	}
L38:
	;
	v112 = int32(1)
	v114 = v103 + v108
	v116 = v87 + v112
	if v116 != v32 {
		v87 = v116
		v89 = v114
		v91 = v91 + v112
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	if v124 < int64(0) {
		v167 = v33
		goto L20
	} else {
		goto L44
	}
L41:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v124) {
		v167 = v33
		goto L20
	} else {
		goto L42
	}
L42:
	;
	if v13 == int32(0) {
		goto L22
	} else {
		goto L43
	}
L43:
	;
	v148 = int64(0) - v124
	goto L23
L44:
	;
	if v13 == int32(0) {
		goto L22
	} else {
		goto L45
	}
L45:
	;
	v148 = v124
	goto L23
L46:
	;
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v177 = v176
	goto L11
L47:
	;
	if v177 <= int64(0) {
		v199 = v179
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v208 = F_lpNext(m, l0, v199)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L13
	} else {
		goto L54
	}
L49:
	;
	v184 = v179
	v190 = int64(0)
	goto L50
L50:
	;
	v193 = F_lpNext(m, l0, v184)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L13
	} else {
		goto L52
	}
L51:
	;
	v199 = v193
	goto L48
L52:
	;
	v196 = v190 + int64(1)
	if v196 != v177 {
		v184 = v193
		v190 = v196
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	if v208 != 0 {
		v481 = v208
		goto L8
	} else {
		goto L55
	}
L55:
	;
	v805 = int32(0)
	goto L5
L56:
	;
	if base.B2i32(v458 == int64(0)) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L57:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v314 = int32(0)
	if base.Ui32(v313+int32(-21)) < base.Ui32(int32(-20)) {
		v448 = v314
		goto L81
	} else {
		goto L82
	}
L58:
	;
	v310 = F_lpGet(m, v303, v13+int32(8), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L13
	} else {
		goto L78
	}
L59:
	;
	goto L58
L60:
	;
	v219 = int32(-1)
	v220 = l0 + v216
	v225 = int32(*(*int8)(unsafe.Add(mBase, uint32(v220+int32(-2)))))
	v228 = base.I64_extend_i32_u(v225 & int32(127))
	if v219 < v225 {
		v296 = v219
		v298 = v228
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v303 = v220 + v219 + (v296 - base.I32_wrap_i64(v298))
	goto L59
L62:
	;
	v233 = int32(*(*int8)(unsafe.Add(mBase, uint32(v220+int32(-3)))))
	v239 = base.I64_extend_i32_u(v233&int32(127))<<(uint(int64(7))%64) | v228
	if int32(-1) < v233 {
		v279 = v239
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v296 = v253
	v298 = int64(-1)
	goto L61
L64:
	;
	if base.Ui64(int64(128)) <= base.Ui64(v279) {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v244 = int32(*(*int8)(unsafe.Add(mBase, uint32(v220+int32(-4)))))
	v250 = base.I64_extend_i32_u(v244&int32(127))<<(uint(int64(14))%64) | v239
	if int32(-1) < v244 {
		v279 = v250
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v253 = int32(-5)
	v256 = int32(*(*int8)(unsafe.Add(mBase, uint32(v220+v253))))
	v262 = base.I64_extend_i32_u(v256&int32(127))<<(uint(int64(21))%64) | v250
	if int32(-1) < v256 {
		v279 = v262
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v267 = int32(*(*int8)(unsafe.Add(mBase, uint32(v220+int32(-6)))))
	if v267 < int32(0) {
		goto L63
	} else {
		goto L68
	}
L68:
	;
	v279 = base.I64_extend_i32_u(v267&int32(127))<<(uint(int64(28))%64) | v262
	goto L64
L69:
	;
	if base.Ui64(int64(16384)) <= base.Ui64(v279) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v296 = int32(-1)
	v298 = v279
	goto L61
L71:
	;
	if base.Ui64(int64(2097152)) <= base.Ui64(v279) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v296 = int32(-2)
	v298 = v279
	goto L61
L73:
	;
	if base.Ui64(v279) < base.Ui64(int64(268435456)) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v296 = int32(-3)
	v298 = v279
	goto L61
L75:
	;
	v293 = int32(-4)
	goto L77
L76:
	;
	v293 = int32(-5)
	goto L77
L77:
	;
	v296 = v293
	v298 = v279
	goto L61
L78:
	;
	if v310 != 0 {
		goto L57
	} else {
		goto L79
	}
L79:
	;
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v458 = v312
	goto L56
L80:
	;
	if v448 == int32(0) {
		goto L3
	} else {
		goto L107
	}
L81:
	;
	goto L80
L82:
	;
	v326 = int32(1)
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	if v313 != v326 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v448 = int32(1)
	goto L81
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v429
	goto L83
L85:
	;
	if v327&int32(255) == int32(45) {
		goto L90
	} else {
		goto L91
	}
L86:
	;
	v331 = v327 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v331&int32(255)) {
		v448 = v314
		goto L81
	} else {
		goto L87
	}
L87:
	;
	if v13 == int32(0) {
		goto L83
	} else {
		goto L88
	}
L88:
	;
	v429 = base.I64_extend_i32_u(v331) & int64(255)
	goto L84
L89:
	;
	if base.Ui32(int32(8)) < base.Ui32((v350+int32(-49))&int32(255)) {
		v448 = v314
		goto L81
	} else {
		goto L92
	}
L90:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
	v349 = int32(2)
	v350 = v347
	v351 = v310 + int32(1)
	goto L89
L91:
	;
	v349 = v326
	v350 = v327
	v351 = v310
	goto L89
L92:
	;
	v362 = base.I64_extend_i32_u(v350+int32(-48)) & int64(255)
	if base.Ui32(v313) <= base.Ui32(v349) {
		v405 = v362
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v327&int32(255) != int32(45) {
		goto L101
	} else {
		goto L102
	}
L94:
	;
	v368 = v349
	v370 = v362
	v372 = v351
	goto L95
L95:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+1)))
	if base.Ui32((v374+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v448 = v314
		goto L81
	} else {
		goto L97
	}
L96:
	;
	v405 = v395
	goto L93
L97:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v370) {
		v448 = v314
		goto L81
	} else {
		goto L98
	}
L98:
	;
	v384 = v370 * int64(10)
	v389 = base.I64_extend_i32_u(v374+int32(-48)) & int64(255)
	if base.Ui64(v389^int64(-1)) < base.Ui64(v384) {
		v448 = v314
		goto L81
	} else {
		goto L99
	}
L99:
	;
	v393 = int32(1)
	v395 = v384 + v389
	v397 = v368 + v393
	if v397 != v313 {
		v368 = v397
		v370 = v395
		v372 = v372 + v393
		goto L95
	} else {
		goto L100
	}
L100:
	;
	goto L96
L101:
	;
	if v405 < int64(0) {
		v448 = v314
		goto L81
	} else {
		goto L105
	}
L102:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v405) {
		v448 = v314
		goto L81
	} else {
		goto L103
	}
L103:
	;
	if v13 == int32(0) {
		goto L83
	} else {
		goto L104
	}
L104:
	;
	v429 = int64(0) - v405
	goto L84
L105:
	;
	if v13 == int32(0) {
		goto L83
	} else {
		goto L106
	}
L106:
	;
	v429 = v405
	goto L84
L107:
	;
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v458 = v457
	goto L56
L108:
	;
	v465 = v303
	v471 = v458
	goto L110
L109:
	;
	v805 = int32(0)
	goto L5
L110:
	;
	v474 = F_lpPrev(m, l0, v465)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L13
	} else {
		goto L112
	}
L111:
	;
	v481 = v474
	goto L8
L112:
	;
	v477 = v471 + int64(-1)
	if v477 != int64(0) {
		v465 = v474
		v471 = v477
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v492 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v493 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v497 = F_lpGet(m, v490, v13+int32(8), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L13
	} else {
		goto L117
	}
L115:
	;
	v646 = F_lpNext(m, l0, v490)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L13
	} else {
		goto L149
	}
L116:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v501 = int32(0)
	if base.Ui32(v500+int32(-21)) < base.Ui32(int32(-20)) {
		v635 = v501
		goto L120
	} else {
		goto L121
	}
L117:
	;
	if v497 != 0 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v499 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v645 = v499
	goto L115
L119:
	;
	if v635 == int32(0) {
		goto L2
	} else {
		goto L146
	}
L120:
	;
	goto L119
L121:
	;
	v513 = int32(1)
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	if v500 != v513 {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	v635 = int32(1)
	goto L120
L123:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v616
	goto L122
L124:
	;
	if v514&int32(255) == int32(45) {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	v518 = v514 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v518&int32(255)) {
		v635 = v501
		goto L120
	} else {
		goto L126
	}
L126:
	;
	if v13 == int32(0) {
		goto L122
	} else {
		goto L127
	}
L127:
	;
	v616 = base.I64_extend_i32_u(v518) & int64(255)
	goto L123
L128:
	;
	if base.Ui32(int32(8)) < base.Ui32((v537+int32(-49))&int32(255)) {
		v635 = v501
		goto L120
	} else {
		goto L131
	}
L129:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+1)))
	v536 = int32(2)
	v537 = v534
	v538 = v497 + int32(1)
	goto L128
L130:
	;
	v536 = v513
	v537 = v514
	v538 = v497
	goto L128
L131:
	;
	v549 = base.I64_extend_i32_u(v537+int32(-48)) & int64(255)
	if base.Ui32(v500) <= base.Ui32(v536) {
		v592 = v549
		goto L132
	} else {
		goto L133
	}
L132:
	;
	if v514&int32(255) != int32(45) {
		goto L140
	} else {
		goto L141
	}
L133:
	;
	v555 = v536
	v557 = v549
	v559 = v538
	goto L134
L134:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+1)))
	if base.Ui32((v561+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v635 = v501
		goto L120
	} else {
		goto L136
	}
L135:
	;
	v592 = v582
	goto L132
L136:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v557) {
		v635 = v501
		goto L120
	} else {
		goto L137
	}
L137:
	;
	v571 = v557 * int64(10)
	v576 = base.I64_extend_i32_u(v561+int32(-48)) & int64(255)
	if base.Ui64(v576^int64(-1)) < base.Ui64(v571) {
		v635 = v501
		goto L120
	} else {
		goto L138
	}
L138:
	;
	v580 = int32(1)
	v582 = v571 + v576
	v584 = v555 + v580
	if v584 != v500 {
		v555 = v584
		v557 = v582
		v559 = v559 + v580
		goto L134
	} else {
		goto L139
	}
L139:
	;
	goto L135
L140:
	;
	if v592 < int64(0) {
		v635 = v501
		goto L120
	} else {
		goto L144
	}
L141:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v592) {
		v635 = v501
		goto L120
	} else {
		goto L142
	}
L142:
	;
	if v13 == int32(0) {
		goto L122
	} else {
		goto L143
	}
L143:
	;
	v616 = int64(0) - v592
	goto L123
L144:
	;
	if v13 == int32(0) {
		goto L122
	} else {
		goto L145
	}
L145:
	;
	v616 = v592
	goto L123
L146:
	;
	v644 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v645 = v644
	goto L115
L147:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v645 + v493
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v799 + v492
	v805 = int32(1)
	goto L5
L148:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v655 = int32(0)
	if base.Ui32(v654+int32(-21)) < base.Ui32(int32(-20)) {
		v789 = v655
		goto L153
	} else {
		goto L154
	}
L149:
	;
	v651 = F_lpGet(m, v646, v13+int32(8), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L13
	} else {
		goto L150
	}
L150:
	;
	if v651 != 0 {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	v653 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v799 = v653
	goto L147
L152:
	;
	if v789 == int32(0) {
		goto L1
	} else {
		goto L179
	}
L153:
	;
	goto L152
L154:
	;
	v667 = int32(1)
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651))))
	if v654 != v667 {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	v789 = int32(1)
	goto L153
L156:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v770
	goto L155
L157:
	;
	if v668&int32(255) == int32(45) {
		goto L162
	} else {
		goto L163
	}
L158:
	;
	v672 = v668 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v672&int32(255)) {
		v789 = v655
		goto L153
	} else {
		goto L159
	}
L159:
	;
	if v13 == int32(0) {
		goto L155
	} else {
		goto L160
	}
L160:
	;
	v770 = base.I64_extend_i32_u(v672) & int64(255)
	goto L156
L161:
	;
	if base.Ui32(int32(8)) < base.Ui32((v691+int32(-49))&int32(255)) {
		v789 = v655
		goto L153
	} else {
		goto L164
	}
L162:
	;
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651)+1)))
	v690 = int32(2)
	v691 = v688
	v692 = v651 + int32(1)
	goto L161
L163:
	;
	v690 = v667
	v691 = v668
	v692 = v651
	goto L161
L164:
	;
	v703 = base.I64_extend_i32_u(v691+int32(-48)) & int64(255)
	if base.Ui32(v654) <= base.Ui32(v690) {
		v746 = v703
		goto L165
	} else {
		goto L166
	}
L165:
	;
	if v668&int32(255) != int32(45) {
		goto L173
	} else {
		goto L174
	}
L166:
	;
	v709 = v690
	v711 = v703
	v713 = v692
	goto L167
L167:
	;
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713)+1)))
	if base.Ui32((v715+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v789 = v655
		goto L153
	} else {
		goto L169
	}
L168:
	;
	v746 = v736
	goto L165
L169:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v711) {
		v789 = v655
		goto L153
	} else {
		goto L170
	}
L170:
	;
	v725 = v711 * int64(10)
	v730 = base.I64_extend_i32_u(v715+int32(-48)) & int64(255)
	if base.Ui64(v730^int64(-1)) < base.Ui64(v725) {
		v789 = v655
		goto L153
	} else {
		goto L171
	}
L171:
	;
	v734 = int32(1)
	v736 = v725 + v730
	v738 = v709 + v734
	if v738 != v654 {
		v709 = v738
		v711 = v736
		v713 = v713 + v734
		goto L167
	} else {
		goto L172
	}
L172:
	;
	goto L168
L173:
	;
	if v746 < int64(0) {
		v789 = v655
		goto L153
	} else {
		goto L177
	}
L174:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v746) {
		v789 = v655
		goto L153
	} else {
		goto L175
	}
L175:
	;
	if v13 == int32(0) {
		goto L155
	} else {
		goto L176
	}
L176:
	;
	v770 = int64(0) - v746
	goto L156
L177:
	;
	if v13 == int32(0) {
		goto L155
	} else {
		goto L178
	}
L178:
	;
	v770 = v746
	goto L156
L179:
	;
	v798 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v799 = v798
	goto L147
L180:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L182:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lpGetValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v14 = F_lpGetWithSize(m, l0, v8+int32(8), v4, v4)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
		if v14 == int32(0) {
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v18
		} else {
			*(*uint32)(unsafe.Add(mBase, uint32(l1))) = uint32(v18)
		}
		m.G0 = v8 + int32(16)
		return v14
	}
}
func F_lpLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v4 != int32(65535) {
		v122 = v4
		return v122
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
		if v7 == int32(255) {
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v100 == int32(7) {
				v118 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v118)
				v122 = v118
				return v122
			} else {
				F__serverAssert(m, int32(_a730), int32(_a729), int32(426))
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			v14 = int32(0)
			v15 = l0 + int32(6)
			for {
				v16 = int32(1)
				v17 = v14 + v16
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				v23 = base.I32_extend8_s(v22)
				if v23 <= int32(-1) {
					if v22&int32(192) != int32(128) {
						if v22&int32(224) != int32(192) {
							v44 = (v23 + int32(15)) & int32(255)
							if base.Ui32(v44) < base.Ui32(int32(4)) {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v44<<(uint(int32(2))%32))+uint32(_consts[375])))
								v85 = v16
								v86 = v84
							} else {
								if v22&int32(240) != int32(224) {
									switch v22 + int32(-240) {
									case 0:
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1))
										v66 = v63 + int32(5)
										if base.Ui32(v66) < base.Ui32(int32(128)) {
											v85 = v16
											v86 = v66
										} else {
											if base.Ui32(int32(16384)) <= base.Ui32(v66) {
												if base.Ui32(int32(2097152)) <= base.Ui32(v66) {
													if base.Ui32(v66) < base.Ui32(int32(268435456)) {
														v79 = int32(4)
													} else {
														v79 = int32(5)
													}
													v85 = v79
													v86 = v66
												} else {
													v85 = int32(3)
													v86 = v66
												}
											} else {
												v85 = int32(2)
												v86 = v66
											}
										}
									default:
										v85 = v16
										v86 = int32(0)
									case 15:
										v85 = v16
										v86 = int32(1)
									}
								} else {
									v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
									v66 = v51 | v22<<(uint(int32(8))%32)&int32(3840) + int32(2)
									if base.Ui32(v66) < base.Ui32(int32(128)) {
										v85 = v16
										v86 = v66
									} else {
										if base.Ui32(int32(16384)) <= base.Ui32(v66) {
											if base.Ui32(int32(2097152)) <= base.Ui32(v66) {
												if base.Ui32(v66) < base.Ui32(int32(268435456)) {
													v79 = int32(4)
												} else {
													v79 = int32(5)
												}
												v85 = v79
												v86 = v66
											} else {
												v85 = int32(3)
												v86 = v66
											}
										} else {
											v85 = int32(2)
											v86 = v66
										}
									}
								}
							}
						} else {
							v85 = v16
							v86 = int32(2)
						}
					} else {
						v31 = int32(1)
						v85 = v31
						v86 = v22&int32(63) + v31
					}
				} else {
					v85 = v16
					v86 = int32(1)
				}
				v88 = v15 + v86 + v85
				v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
				if v89 != int32(255) {
					v14 = v17
					v15 = v88
					continue
				} else {
					break
				}
				break
			}
			v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v88+int32(1) != l0+v94 {
				F__serverAssert(m, int32(_a730), int32(_a729), int32(400))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.Ui32(int32(65534)) < base.Ui32(v17) {
					v122 = v17
				} else {
					v118 = v17
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v118)
					v122 = v118
				}
				return v122
			}
		}
	}
}
func F_lpMerge(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int64
	_ = v225
	var v228 int64
	_ = v228
	var v231 int64
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v253 int32
	_ = v253
	v3 = int32(0)
	if l0 == v3 {
		v240 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a736), int32(_a729), int32(1081))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L8
	} else {
		goto L73
	}
L2:
	;
	return v240
L3:
	;
	if l1 == int32(0) {
		v240 = v3
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 == int32(0) {
		v240 = v3
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v19 == int32(0) {
		v240 = v3
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if v16 == v19 {
		v240 = v3
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v24 = F_lpLength(m, v16)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v30 = F_lpLength(m, v28)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v36 = base.I64_extend_i32_u(v23) + base.I64_extend_i32_u(v29) + int64(-7)
	if base.Ui64(int64(4294967295)) <= base.Ui64(v36) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v39 = base.B2i32(base.Ui32(v23) < base.Ui32(v29))
	if base.Ui32(v23) < base.Ui32(v29) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v40 = l0
	goto L14
L13:
	;
	v40 = l1
	goto L14
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = v30 + v24
	v43 = int32(65535)
	if base.Ui32(v42) < base.Ui32(v43) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v46 = v42
	goto L17
L16:
	;
	v46 = v43
	goto L17
L17:
	;
	v48 = v29 + int32(-6)
	if base.Ui32(v23) < base.Ui32(v29) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v49 = l1
	goto L20
L19:
	;
	v49 = l0
	goto L20
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v53 = F_zrealloc_usable(m, v50, base.I32_wrap_i64(v36), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v57 = v53 + v23 + int32(-1)
	if base.Ui32(v23) < base.Ui32(v29) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v36)
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+4)) = uint16(v46)
	v225 = int64(base.Ui64(v36) >> (uint(int64(24)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+3)) = uint8(v225)
	v228 = int64(base.Ui64(v36) >> (uint(int64(16)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+2)) = uint8(v228)
	v231 = int64(base.Ui64(v36) >> (uint(int64(8)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)) = uint8(v231)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	F_valkey_free(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L8
	} else {
		goto L72
	}
L23:
	;
	v65 = v53 + int32(6)
	if v57 == v65 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	if v48 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v220 = l1
	v221 = l0
	goto L22
L26:
	;
	goto L25
L27:
	;
	v62 = F__emscripten_memcpy_bulkmem(m, v57, v41+int32(6), v48)
	mBase = m.M
	goto L26
L28:
	;
	v215 = v23 + int32(-1)
	if v215 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L29:
	;
	goto L28
L30:
	;
	v69 = v48 + v57
	if base.Ui32(int32(0)-v48<<(uint(int32(1))%32)) < base.Ui32(v65-v69) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v79 = (v65 ^ v57) & int32(3)
	if base.Ui32(v65) <= base.Ui32(v57) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v76 = F___memcpy(m, v57, v65, v48)
	mBase = m.M
	goto L28
L33:
	;
	if v185 == int32(0) {
		goto L29
	} else {
		goto L65
	}
L34:
	;
	if base.Ui32(v163) <= base.Ui32(int32(3)) {
		v184 = v162
		v185 = v163
		v186 = v164
		goto L33
	} else {
		goto L61
	}
L35:
	;
	if v79 != 0 {
		v145 = v48
		goto L45
	} else {
		goto L46
	}
L36:
	;
	if v79 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if v57&int32(3) != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v184 = v65
	v185 = v48
	v186 = v57
	goto L33
L39:
	;
	v86 = v65
	v87 = v48
	v88 = v57
	goto L41
L40:
	;
	v162 = v65
	v163 = v48
	v164 = v57
	goto L34
L41:
	;
	if v87 == int32(0) {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v92)
	v94 = int32(1)
	v95 = v86 + v94
	v97 = v87 + int32(-1)
	v99 = v88 + v94
	if v99&int32(3) == int32(0) {
		v162 = v95
		v163 = v97
		v164 = v99
		goto L34
	} else {
		goto L44
	}
L44:
	;
	v86 = v95
	v87 = v97
	v88 = v99
	goto L41
L45:
	;
	if v145 == int32(0) {
		goto L29
	} else {
		goto L57
	}
L46:
	;
	if v69&int32(3) == int32(0) {
		v125 = v48
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if base.Ui32(v125) <= base.Ui32(int32(3)) {
		v145 = v125
		goto L45
	} else {
		goto L53
	}
L48:
	;
	v110 = v48
	goto L49
L49:
	;
	if v110 == int32(0) {
		goto L29
	} else {
		goto L51
	}
L50:
	;
	v125 = v116
	goto L47
L51:
	;
	v116 = v110 + int32(-1)
	v117 = v57 + v116
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+v116))))
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v119)
	if v117&int32(3) != 0 {
		v110 = v116
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v132 = v125
	goto L54
L54:
	;
	v136 = v132 + int32(-4)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v65+v136)))
	*(*int32)(unsafe.Add(mBase, uint32(v57+v136))) = v139
	if base.Ui32(int32(3)) < base.Ui32(v136) {
		v132 = v136
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v145 = v136
	goto L45
L56:
	;
	goto L55
L57:
	;
	v152 = v145
	goto L58
L58:
	;
	v156 = v152 + int32(-1)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+v156))))
	*(*uint8)(unsafe.Add(mBase, uint32(v57+v156))) = uint8(v159)
	if v156 != 0 {
		v152 = v156
		goto L58
	} else {
		goto L60
	}
L60:
	;
	goto L29
L61:
	;
	v169 = v162
	v170 = v163
	v171 = v164
	goto L62
L62:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v173
	v175 = int32(4)
	v176 = v169 + v175
	v178 = v171 + v175
	v180 = v170 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v180) {
		v169 = v176
		v170 = v180
		v171 = v178
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v184 = v176
	v185 = v180
	v186 = v178
	goto L33
L64:
	;
	goto L63
L65:
	;
	v191 = v184
	v192 = v185
	v193 = v186
	goto L66
L66:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	*(*uint8)(unsafe.Add(mBase, uint32(v193))) = uint8(v195)
	v197 = int32(1)
	v202 = v192 + int32(-1)
	if v202 != 0 {
		v191 = v191 + v197
		v192 = v202
		v193 = v193 + v197
		goto L66
	} else {
		goto L68
	}
L67:
	;
	goto L29
L68:
	;
	goto L67
L69:
	;
	v220 = l0
	v221 = l1
	goto L22
L70:
	;
	goto L69
L71:
	;
	v218 = F__emscripten_memcpy_bulkmem(m, v53, v41, v215)
	mBase = m.M
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = v53
	v240 = v53
	goto L2
L73:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lpNext(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	if l1 == int32(0) {
		F__serverAssert(m, int32(_a728), int32(_a729), int32(395))
		mBase = m.M
		v95 = m.ExcPending
		if v95 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v9 = int32(1)
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		v11 = base.I32_extend8_s(v10)
		if v11 <= int32(-1) {
			if v10&int32(192) != int32(128) {
				if v10&int32(224) != int32(192) {
					v32 = (v11 + int32(15)) & int32(255)
					if base.Ui32(v32) < base.Ui32(int32(4)) {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_consts[375])))
						v73 = v9
						v74 = v72
					} else {
						if v10&int32(240) != int32(224) {
							switch v10 + int32(-240) {
							case 0:
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1))
								v54 = v51 + int32(5)
								if base.Ui32(v54) < base.Ui32(int32(128)) {
									v73 = v9
									v74 = v54
								} else {
									if base.Ui32(int32(16384)) <= base.Ui32(v54) {
										if base.Ui32(int32(2097152)) <= base.Ui32(v54) {
											if base.Ui32(v54) < base.Ui32(int32(268435456)) {
												v67 = int32(4)
											} else {
												v67 = int32(5)
											}
											v73 = v67
											v74 = v54
										} else {
											v73 = int32(3)
											v74 = v54
										}
									} else {
										v73 = int32(2)
										v74 = v54
									}
								}
							default:
								v73 = v9
								v74 = int32(0)
							case 15:
								v73 = v9
								v74 = int32(1)
							}
						} else {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
							v54 = v39 | v10<<(uint(int32(8))%32)&int32(3840) + int32(2)
							if base.Ui32(v54) < base.Ui32(int32(128)) {
								v73 = v9
								v74 = v54
							} else {
								if base.Ui32(int32(16384)) <= base.Ui32(v54) {
									if base.Ui32(int32(2097152)) <= base.Ui32(v54) {
										if base.Ui32(v54) < base.Ui32(int32(268435456)) {
											v67 = int32(4)
										} else {
											v67 = int32(5)
										}
										v73 = v67
										v74 = v54
									} else {
										v73 = int32(3)
										v74 = v54
									}
								} else {
									v73 = int32(2)
									v74 = v54
								}
							}
						}
					}
				} else {
					v73 = v9
					v74 = int32(2)
				}
			} else {
				v19 = int32(1)
				v73 = v19
				v74 = v10&int32(63) + v19
			}
		} else {
			v73 = v9
			v74 = int32(1)
		}
		v76 = l1 + v74 + v73
		v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
		if v77 != int32(255) {
			v86 = v76
			return v86
		} else {
			v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v76+int32(1) != l0+v83 {
				F__serverAssert(m, int32(_a730), int32(_a729), int32(400))
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v86 = int32(0)
				return v86
			}
		}
	}
}
func F_lpRandomPair(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1 == int32(0) {
		F__serverAssert(m, int32(_a737), int32(_a729), int32(1322))
		mBase = m.M
		v153 = m.ExcPending
		if v153 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v15 = int32(0)
		v17 = *(*int64)(unsafe.Add(mBase, _consts[343]))
		v21 = v17*int64(6364136223846793005) + int64(1)
		*(*int64)(unsafe.Add(mBase, _consts[343])) = v21
		v26 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v21)>>(uint(int64(33))%64))), l1)
		v29 = F_lpSeek(m, l0, v26<<(uint(int32(1))%32))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			if v29 == int32(0) {
				F__serverAssert(m, int32(_a738), int32(_a729), int32(1326))
				mBase = m.M
				v159 = m.ExcPending
				if v159 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v35 = int32(0)
				v37 = F_lpGetWithSize(m, v29, v10+int32(8), v35, v35)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
					if v37 == int32(0) {
						*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v39
					} else {
						*(*uint32)(unsafe.Add(mBase, uint32(l2)+4)) = uint32(v39)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37
					if l3 == int32(0) {
						m.G0 = v10 + int32(16)
						return
					} else {
						v50 = int32(1)
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
						v52 = base.I32_extend8_s(v51)
						if v52 <= int32(-1) {
							if v51&int32(192) != int32(128) {
								if v51&int32(224) != int32(192) {
									v73 = (v52 + int32(15)) & int32(255)
									if base.Ui32(v73) < base.Ui32(int32(4)) {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v73<<(uint(int32(2))%32))+uint32(_consts[375])))
										v114 = v50
										v115 = v113
									} else {
										if v51&int32(240) != int32(224) {
											switch v51 + int32(-240) {
											case 0:
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v29)+1))
												v95 = v92 + int32(5)
												if base.Ui32(v95) < base.Ui32(int32(128)) {
													v114 = v50
													v115 = v95
												} else {
													if base.Ui32(int32(16384)) <= base.Ui32(v95) {
														if base.Ui32(int32(2097152)) <= base.Ui32(v95) {
															if base.Ui32(v95) < base.Ui32(int32(268435456)) {
																v108 = int32(4)
															} else {
																v108 = int32(5)
															}
															v114 = v108
															v115 = v95
														} else {
															v114 = int32(3)
															v115 = v95
														}
													} else {
														v114 = int32(2)
														v115 = v95
													}
												}
											default:
												v114 = v50
												v115 = int32(0)
											case 15:
												v114 = v50
												v115 = int32(1)
											}
										} else {
											v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
											v95 = v80 | v51<<(uint(int32(8))%32)&int32(3840) + int32(2)
											if base.Ui32(v95) < base.Ui32(int32(128)) {
												v114 = v50
												v115 = v95
											} else {
												if base.Ui32(int32(16384)) <= base.Ui32(v95) {
													if base.Ui32(int32(2097152)) <= base.Ui32(v95) {
														if base.Ui32(v95) < base.Ui32(int32(268435456)) {
															v108 = int32(4)
														} else {
															v108 = int32(5)
														}
														v114 = v108
														v115 = v95
													} else {
														v114 = int32(3)
														v115 = v95
													}
												} else {
													v114 = int32(2)
													v115 = v95
												}
											}
										}
									}
								} else {
									v114 = v50
									v115 = int32(2)
								}
							} else {
								v60 = int32(1)
								v114 = v60
								v115 = v51&int32(63) + v60
							}
						} else {
							v114 = v50
							v115 = int32(1)
						}
						v117 = v29 + v115 + v114
						v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
						if v118 != int32(255) {
							v134 = int32(0)
							v136 = F_lpGetWithSize(m, v117, v10+int32(8), v134, v134)
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return
							} else {
								v138 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
								if v136 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v138
								} else {
									*(*uint32)(unsafe.Add(mBase, uint32(l3)+4)) = uint32(v138)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v136
								m.G0 = v10 + int32(16)
								return
							}
						} else {
							v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							if v117+int32(1) != l0+v123 {
								F__serverAssert(m, int32(_a730), int32(_a729), int32(400))
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								F__serverAssert(m, int32(_a739), int32(_a729), int32(1330))
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
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
func F_lpRandomPairs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v222 int64
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	v5 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v27 = F_zmalloc_usable(m, l1<<(uint(int32(3))%32), v5)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = F_lpLength(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F__serverAssert(m, int32(_a740), int32(_a729), int32(1389))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L99
	}
L4:
	;
	if base.Ui32(v29) < base.Ui32(int32(2)) {
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
	F_qsort(m, v27, l1, int32(8), int32(560))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	v42 = v5
	goto L8
L8:
	;
	v56 = int32(0)
	v58 = *(*int64)(unsafe.Add(mBase, _consts[343]))
	v62 = v58*int64(6364136223846793005) + int64(1)
	*(*int64)(unsafe.Add(mBase, _consts[343])) = v62
	goto L10
L9:
	;
	goto L6
L10:
	;
	v69 = v27 + v42<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v42
	v71 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v62)>>(uint(int64(33))%64))), int32(base.Ui32(v29)>>(uint(int32(1))%32)))
	v72 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v71 << (uint(v72) % 32)
	v76 = v42 + v72
	if v76 != l1 {
		v42 = v76
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v101 = F_lpSeek(m, l0, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if l1 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F__serverAssert(m, int32(_a730), int32(_a729), int32(400))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L98
	}
L15:
	;
	F_valkey_free(m, v27)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L97
	}
L16:
	;
	if v101 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v107 = int64(0)
	v108 = int32(0)
	v117 = v108
	v119 = v101
	v122 = v100
	v123 = v107
	v124 = v107
	v125 = v108
	v126 = v108
	goto L18
L18:
	;
	v132 = int32(0)
	v134 = F_lpGetWithSize(m, v119, v21+int32(8), v132, v132)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L15
L20:
	;
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	v140 = int32(1)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v142 = base.I32_extend8_s(v141)
	if v142 <= int32(-1) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	if v134 != 0 {
		goto L48
	} else {
		goto L49
	}
L22:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v208 != int32(255) {
		goto L21
	} else {
		goto L45
	}
L23:
	;
	v207 = v119 + v205 + v204
	goto L22
L24:
	;
	if v141&int32(192) != int32(128) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v204 = v140
	v205 = int32(1)
	goto L23
L26:
	;
	if v141&int32(224) != int32(192) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v150 = int32(1)
	v204 = v150
	v205 = v141&int32(63) + v150
	goto L23
L28:
	;
	v163 = (v142 + int32(15)) & int32(255)
	if base.Ui32(v163) < base.Ui32(int32(4)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v204 = v140
	v205 = int32(2)
	goto L23
L30:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v163<<(uint(int32(2))%32))+uint32(_consts[375])))
	v204 = v140
	v205 = v203
	goto L23
L31:
	;
	if v141&int32(240) != int32(224) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if base.Ui32(v185) < base.Ui32(int32(128)) {
		v204 = v140
		v205 = v185
		goto L23
	} else {
		goto L37
	}
L33:
	;
	switch v141 + int32(-240) {
	case 0:
		goto L35
	default:
		goto L36
	case 15:
		v204 = v140
		v205 = int32(1)
		goto L23
	}
L34:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	v185 = v170 | v141<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L32
L35:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v119)+1))
	v185 = v182 + int32(5)
	goto L32
L36:
	;
	v204 = v140
	v205 = int32(0)
	goto L23
L37:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v185) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v185) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v204 = int32(2)
	v205 = v185
	goto L23
L40:
	;
	if base.Ui32(v185) < base.Ui32(int32(268435456)) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v204 = int32(3)
	v205 = v185
	goto L23
L42:
	;
	v198 = int32(4)
	goto L44
L43:
	;
	v198 = int32(5)
	goto L44
L44:
	;
	v204 = v198
	v205 = v185
	goto L23
L45:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v207+int32(1) != l0+v213 {
		goto L14
	} else {
		goto L46
	}
L46:
	;
	F__serverAssert(m, int32(_a739), int32(_a729), int32(1406))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v222 = v124
	goto L50
L49:
	;
	v222 = v136
	goto L50
L50:
	;
	if v134 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v224 = base.I32_wrap_i64(v136)
	goto L53
L52:
	;
	v224 = v126
	goto L53
L53:
	;
	v227 = int32(0)
	v229 = F_lpGetWithSize(m, v207, v21+int32(8), v227, v227)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	if v229 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v232 = v123
	goto L57
L56:
	;
	v232 = v231
	goto L57
L57:
	;
	if v229 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v234 = base.I32_wrap_i64(v231)
	goto L60
L59:
	;
	v234 = v125
	goto L60
L60:
	;
	if base.Ui32(l1) <= base.Ui32(v117) {
		v281 = v117
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v297 = int32(1)
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	v299 = base.I32_extend8_s(v298)
	if v299 <= int32(-1) {
		goto L72
	} else {
		goto L73
	}
L62:
	;
	v241 = v117
	goto L63
L63:
	;
	v256 = v27 + v241<<(uint(int32(3))%32)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if v122 != v257 {
		v281 = v241
		goto L61
	} else {
		goto L65
	}
L64:
	;
	v281 = l1
	goto L61
L65:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v261 = v259 << (uint(int32(4)) % 32)
	v262 = l2 + v261
	*(*int64)(unsafe.Add(mBase, uint32(v262)+8)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v262)+4)) = v224
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v134
	if l3 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v274 = v241 + int32(1)
	if v274 != l1 {
		v241 = v274
		goto L63
	} else {
		goto L68
	}
L67:
	;
	v268 = l3 + v261
	*(*int64)(unsafe.Add(mBase, uint32(v268)+8)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v229
	goto L66
L68:
	;
	goto L64
L69:
	;
	if base.Ui32(v281) < base.Ui32(l1) {
		v117 = v281
		v119 = v364
		v122 = v122 + int32(2)
		v123 = v232
		v124 = v222
		v125 = v234
		v126 = v224
		goto L18
	} else {
		goto L96
	}
L70:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
	if v365 != int32(255) {
		goto L69
	} else {
		goto L93
	}
L71:
	;
	v364 = v207 + v362 + v361
	goto L70
L72:
	;
	if v298&int32(192) != int32(128) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v361 = v297
	v362 = int32(1)
	goto L71
L74:
	;
	if v298&int32(224) != int32(192) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v307 = int32(1)
	v361 = v307
	v362 = v298&int32(63) + v307
	goto L71
L76:
	;
	v320 = (v299 + int32(15)) & int32(255)
	if base.Ui32(v320) < base.Ui32(int32(4)) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v361 = v297
	v362 = int32(2)
	goto L71
L78:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v320<<(uint(int32(2))%32))+uint32(_consts[375])))
	v361 = v297
	v362 = v360
	goto L71
L79:
	;
	if v298&int32(240) != int32(224) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if base.Ui32(v342) < base.Ui32(int32(128)) {
		v361 = v297
		v362 = v342
		goto L71
	} else {
		goto L85
	}
L81:
	;
	switch v298 + int32(-240) {
	case 0:
		goto L83
	default:
		goto L84
	case 15:
		v361 = v297
		v362 = int32(1)
		goto L71
	}
L82:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
	v342 = v327 | v298<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L80
L83:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v207)+1))
	v342 = v339 + int32(5)
	goto L80
L84:
	;
	v361 = v297
	v362 = int32(0)
	goto L71
L85:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v342) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v342) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v361 = int32(2)
	v362 = v342
	goto L71
L88:
	;
	if base.Ui32(v342) < base.Ui32(int32(268435456)) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v361 = int32(3)
	v362 = v342
	goto L71
L90:
	;
	v355 = int32(4)
	goto L92
L91:
	;
	v355 = int32(5)
	goto L92
L92:
	;
	v361 = v355
	v362 = v342
	goto L71
L93:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v364+int32(1) == l0+v370 {
		goto L15
	} else {
		goto L94
	}
L94:
	;
	F__serverAssert(m, int32(_a730), int32(_a729), int32(400))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	goto L19
L97:
	;
	m.G0 = v21 + int32(16)
	return
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
}
func F_lpReplace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = F_lpInsert(m, l0, l2, int32(0), l3, v6, int32(2), l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_lpSeek(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int64
	_ = v158
	var v163 int32
	_ = v163
	var v169 int64
	_ = v169
	var v174 int32
	_ = v174
	var v180 int64
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v192 int64
	_ = v192
	var v197 int32
	_ = v197
	var v209 int64
	_ = v209
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int64
	_ = v228
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int64
	_ = v253
	var v259 int32
	_ = v259
	var v265 int64
	_ = v265
	var v270 int32
	_ = v270
	var v276 int64
	_ = v276
	var v281 int32
	_ = v281
	var v287 int64
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v304 int64
	_ = v304
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int64
	_ = v322
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v8 == int32(65535) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v332
L2:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v144 != int32(7) {
		goto L48
	} else {
		goto L49
	}
L3:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v25 != int32(255) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	if l1 < int32(0) {
		v143 = l1
		goto L2
	} else {
		goto L9
	}
L5:
	;
	v14 = l1>>(uint(int32(31))%32)&v8 + l1
	if base.Ui32(v14) < base.Ui32(v8) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if base.Ui32(v14) <= base.Ui32(int32(base.Ui32(v8)>>(uint(int32(1))%32))) {
		v24 = v14
		goto L3
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v143 = v14 - v8
	goto L2
L9:
	;
	v24 = l1
	goto L3
L10:
	;
	v42 = l0 + int32(6)
	if v24 == int32(0) {
		v332 = v42
		goto L1
	} else {
		goto L16
	}
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v28 != int32(7) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F__serverAssert(m, int32(_a730), int32(_a729), int32(426))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	return int32(0)
L14:
	;
	return int32(0)
L15:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	v46 = v24
	v47 = v42
	goto L17
L17:
	;
	v55 = int32(1)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v57 = base.I32_extend8_s(v56)
	if v57 <= int32(-1) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	if base.Ui32(int32(1)) < base.Ui32(v46) {
		v46 = v46 + int32(-1)
		v47 = v122
		goto L17
	} else {
		goto L47
	}
L20:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v123 != int32(255) {
		goto L19
	} else {
		goto L43
	}
L21:
	;
	v122 = v47 + v120 + v119
	goto L20
L22:
	;
	if v56&int32(192) != int32(128) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v119 = v55
	v120 = int32(1)
	goto L21
L24:
	;
	if v56&int32(224) != int32(192) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v65 = int32(1)
	v119 = v65
	v120 = v56&int32(63) + v65
	goto L21
L26:
	;
	v78 = (v57 + int32(15)) & int32(255)
	if base.Ui32(v78) < base.Ui32(int32(4)) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v119 = v55
	v120 = int32(2)
	goto L21
L28:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(int32(2))%32))+uint32(_consts[375])))
	v119 = v55
	v120 = v118
	goto L21
L29:
	;
	if v56&int32(240) != int32(224) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if base.Ui32(v100) < base.Ui32(int32(128)) {
		v119 = v55
		v120 = v100
		goto L21
	} else {
		goto L35
	}
L31:
	;
	switch v56 + int32(-240) {
	case 0:
		goto L33
	default:
		goto L34
	case 15:
		v119 = v55
		v120 = int32(1)
		goto L21
	}
L32:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	v100 = v85 | v56<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L30
L33:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v47)+1))
	v100 = v97 + int32(5)
	goto L30
L34:
	;
	v119 = v55
	v120 = int32(0)
	goto L21
L35:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v100) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v100) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v119 = int32(2)
	v120 = v100
	goto L21
L38:
	;
	if base.Ui32(v100) < base.Ui32(int32(268435456)) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v119 = int32(3)
	v120 = v100
	goto L21
L40:
	;
	v113 = int32(4)
	goto L42
L41:
	;
	v113 = int32(5)
	goto L42
L42:
	;
	v119 = v113
	v120 = v100
	goto L21
L43:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v122+int32(1) != l0+v128 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F__serverAssert(m, int32(_a730), int32(_a729), int32(400))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L14
	} else {
		goto L46
	}
L45:
	;
	return int32(0)
L46:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v332 = v122
	goto L1
L48:
	;
	v149 = int32(-1)
	v150 = l0 + v144
	v155 = int32(*(*int8)(unsafe.Add(mBase, uint32(v150+int32(-2)))))
	v158 = base.I64_extend_i32_u(v155 & int32(127))
	if v149 < v155 {
		v226 = v149
		v228 = v158
		goto L50
	} else {
		goto L51
	}
L49:
	;
	return int32(0)
L50:
	;
	v231 = v150 + v149 + (v226 - base.I32_wrap_i64(v228))
	if int32(-2) < v143 {
		v332 = v231
		goto L1
	} else {
		goto L67
	}
L51:
	;
	v163 = int32(*(*int8)(unsafe.Add(mBase, uint32(v150+int32(-3)))))
	v169 = base.I64_extend_i32_u(v163&int32(127))<<(uint(int64(7))%64) | v158
	if int32(-1) < v163 {
		v209 = v169
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v226 = v183
	v228 = int64(-1)
	goto L50
L53:
	;
	if base.Ui64(int64(128)) <= base.Ui64(v209) {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	v174 = int32(*(*int8)(unsafe.Add(mBase, uint32(v150+int32(-4)))))
	v180 = base.I64_extend_i32_u(v174&int32(127))<<(uint(int64(14))%64) | v169
	if int32(-1) < v174 {
		v209 = v180
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v183 = int32(-5)
	v186 = int32(*(*int8)(unsafe.Add(mBase, uint32(v150+v183))))
	v192 = base.I64_extend_i32_u(v186&int32(127))<<(uint(int64(21))%64) | v180
	if int32(-1) < v186 {
		v209 = v192
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v197 = int32(*(*int8)(unsafe.Add(mBase, uint32(v150+int32(-6)))))
	if v197 < int32(0) {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	v209 = base.I64_extend_i32_u(v197&int32(127))<<(uint(int64(28))%64) | v192
	goto L53
L58:
	;
	if base.Ui64(int64(16384)) <= base.Ui64(v209) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v226 = int32(-1)
	v228 = v209
	goto L50
L60:
	;
	if base.Ui64(int64(2097152)) <= base.Ui64(v209) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v226 = int32(-2)
	v228 = v209
	goto L50
L62:
	;
	if base.Ui64(v209) < base.Ui64(int64(268435456)) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v226 = int32(-3)
	v228 = v209
	goto L50
L64:
	;
	v223 = int32(-4)
	goto L66
L65:
	;
	v223 = int32(-5)
	goto L66
L66:
	;
	v226 = v223
	v228 = v209
	goto L50
L67:
	;
	if v231 == int32(0) {
		v332 = v231
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v237 = v143
	v238 = v231
	goto L69
L69:
	;
	if v238-l0 != int32(6) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v332 = v325
	goto L1
L71:
	;
	v248 = int32(-1)
	v250 = int32(*(*int8)(unsafe.Add(mBase, uint32(v238+v248))))
	v253 = base.I64_extend_i32_u(v250 & int32(127))
	if v250 <= v248 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	return int32(0)
L73:
	;
	v325 = v238 + (v320 - base.I32_wrap_i64(v322))
	if base.Ui32(int32(-3)) < base.Ui32(v237) {
		v332 = v325
		goto L1
	} else {
		goto L91
	}
L74:
	;
	v259 = int32(*(*int8)(unsafe.Add(mBase, uint32(v238+int32(-2)))))
	v265 = base.I64_extend_i32_u(v259&int32(127))<<(uint(int64(7))%64) | v253
	if int32(-1) < v259 {
		v304 = v265
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v320 = int32(-1)
	v322 = v253
	goto L73
L76:
	;
	v320 = v290
	v322 = int64(-1)
	goto L73
L77:
	;
	if base.Ui64(int64(128)) <= base.Ui64(v304) {
		goto L82
	} else {
		goto L83
	}
L78:
	;
	v270 = int32(*(*int8)(unsafe.Add(mBase, uint32(v238+int32(-3)))))
	v276 = base.I64_extend_i32_u(v270&int32(127))<<(uint(int64(14))%64) | v265
	if int32(-1) < v270 {
		v304 = v276
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v281 = int32(*(*int8)(unsafe.Add(mBase, uint32(v238+int32(-4)))))
	v287 = base.I64_extend_i32_u(v281&int32(127))<<(uint(int64(21))%64) | v276
	if int32(-1) < v281 {
		v304 = v287
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v290 = int32(-5)
	v293 = int32(*(*int8)(unsafe.Add(mBase, uint32(v238+v290))))
	if v293 < int32(0) {
		goto L76
	} else {
		goto L81
	}
L81:
	;
	v304 = base.I64_extend_i32_u(v293&int32(127))<<(uint(int64(28))%64) | v287
	goto L77
L82:
	;
	if base.Ui64(int64(16384)) <= base.Ui64(v304) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v320 = int32(-1)
	v322 = v304
	goto L73
L84:
	;
	if base.Ui64(int64(2097152)) <= base.Ui64(v304) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v320 = int32(-2)
	v322 = v304
	goto L73
L86:
	;
	if base.Ui64(v304) < base.Ui64(int64(268435456)) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v320 = int32(-3)
	v322 = v304
	goto L73
L88:
	;
	v318 = int32(-4)
	goto L90
L89:
	;
	v318 = int32(-5)
	goto L90
L90:
	;
	v320 = v318
	v322 = v304
	goto L73
L91:
	;
	if v325 != 0 {
		v237 = v237 + int32(1)
		v238 = v325
		goto L69
	} else {
		goto L92
	}
L92:
	;
	goto L70
}
func F_lpValidateIntegrity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int64
	_ = v173
	var v178 int32
	_ = v178
	var v184 int64
	_ = v184
	var v189 int32
	_ = v189
	var v195 int64
	_ = v195
	var v200 int32
	_ = v200
	var v206 int64
	_ = v206
	var v211 int32
	_ = v211
	var v221 int64
	_ = v221
	var v223 int64
	_ = v223
	var v228 int32
	_ = v228
	var v239 int32
	_ = v239
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if base.Ui32(l1) < base.Ui32(int32(7)) {
		v286 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v286
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18 != l1 {
		v286 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = l0 + l1 + int32(-1)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v23 != int32(255) {
		v286 = v5
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v28 = l0 + int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v28
	v39 = v28
	v40 = int32(0)
	goto L6
L5:
	;
	if v273 != v22 {
		v286 = int32(0)
		goto L1
	} else {
		goto L68
	}
L6:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v41 == int32(255) {
		v273 = v39
		v274 = v40
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v273 = int32(0)
	v274 = v269
	goto L5
L8:
	;
	v44 = int32(0)
	v46 = v13 + int32(12)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v56 == v44 {
		v239 = v44
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v257 == int32(0) {
		v286 = v44
		goto L1
	} else {
		goto L61
	}
L10:
	;
	v257 = v239
	goto L9
L11:
	;
	v60 = l0 + int32(6)
	if base.Ui32(v56) < base.Ui32(v60) {
		v239 = v44
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v62 = l0 + l1
	v64 = v62 + int32(-1)
	if base.Ui32(v64) < base.Ui32(v56) {
		v239 = v44
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v66 != int32(255) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v228
	v239 = int32(1)
	goto L10
L15:
	;
	v74 = base.I32_extend8_s(v66)
	v75 = int32(192)
	v76 = v66 & v75
	v77 = int32(1)
	v79 = v66 & int32(224)
	if v79 == v75 {
		v100 = v77
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v56+int32(1) == v62 {
		v228 = int32(0)
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v257 = int32(0)
	goto L9
L18:
	;
	v101 = v56 + v100
	if base.Ui32(v101) < base.Ui32(v60) {
		v239 = v44
		goto L10
	} else {
		goto L26
	}
L19:
	;
	if int32(-1) < v74 {
		v100 = v77
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32((v74+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v100 = v77
		goto L18
	} else {
		goto L21
	}
L21:
	;
	if v76 == int32(128) {
		v100 = v77
		goto L18
	} else {
		goto L22
	}
L22:
	;
	if v66&int32(240) != int32(224) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v74 != int32(-16) {
		v239 = v44
		goto L10
	} else {
		goto L25
	}
L24:
	;
	v100 = int32(2)
	goto L18
L25:
	;
	v100 = int32(5)
	goto L18
L26:
	;
	if base.Ui32(v64) < base.Ui32(v101) {
		v239 = v44
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v104 = int32(1)
	if v74 <= int32(-1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v164 = v163 + v161
	v165 = v56 + v164
	if base.Ui32(v165) < base.Ui32(v60) {
		v239 = v44
		goto L10
	} else {
		goto L50
	}
L29:
	;
	if v76 != int32(128) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v161 = int32(1)
	v163 = v104
	goto L28
L31:
	;
	if v79 != int32(192) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v110 = int32(1)
	v161 = v66&int32(63) + v110
	v163 = v110
	goto L28
L33:
	;
	v121 = (v74 + int32(15)) & int32(255)
	if base.Ui32(v121) < base.Ui32(int32(4)) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v161 = int32(2)
	v163 = v104
	goto L28
L35:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v121<<(uint(int32(2))%32))+uint32(_consts[375])))
	v161 = v160
	v163 = v104
	goto L28
L36:
	;
	if v66&int32(240) != int32(224) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if base.Ui32(v142) < base.Ui32(int32(128)) {
		v161 = v142
		v163 = v104
		goto L28
	} else {
		goto L42
	}
L38:
	;
	if v74 == int32(-16) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v142 = v128 | v66<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L37
L40:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v56)+1))
	v142 = v139 + int32(5)
	goto L37
L41:
	;
	v161 = int32(0)
	v163 = v104
	goto L28
L42:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v142) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v142) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v161 = v142
	v163 = int32(2)
	goto L28
L45:
	;
	if base.Ui32(v142) < base.Ui32(int32(268435456)) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v161 = v142
	v163 = int32(3)
	goto L28
L47:
	;
	v155 = int32(4)
	goto L49
L48:
	;
	v155 = int32(5)
	goto L49
L49:
	;
	v161 = v142
	v163 = v155
	goto L28
L50:
	;
	if base.Ui32(v64) < base.Ui32(v165) {
		v239 = v44
		goto L10
	} else {
		goto L51
	}
L51:
	;
	v168 = int32(-1)
	v170 = int32(*(*int8)(unsafe.Add(mBase, uint32(v165+v168))))
	v173 = base.I64_extend_i32_u(v170 & int32(127))
	if v168 < v170 {
		v223 = v173
		goto L52
	} else {
		goto L53
	}
L52:
	;
	if v223+base.I64_extend_i32_u(v163) != base.I64_extend_i32_u(v164) {
		v239 = v44
		goto L10
	} else {
		goto L60
	}
L53:
	;
	v178 = int32(*(*int8)(unsafe.Add(mBase, uint32(v165+int32(-2)))))
	v184 = base.I64_extend_i32_u(v178&int32(127))<<(uint(int64(7))%64) | v173
	if int32(-1) < v178 {
		v223 = v184
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v189 = int32(*(*int8)(unsafe.Add(mBase, uint32(v165+int32(-3)))))
	v195 = base.I64_extend_i32_u(v189&int32(127))<<(uint(int64(14))%64) | v184
	if int32(-1) < v189 {
		v223 = v195
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v200 = int32(*(*int8)(unsafe.Add(mBase, uint32(v165+int32(-4)))))
	v206 = base.I64_extend_i32_u(v200&int32(127))<<(uint(int64(21))%64) | v195
	if int32(-1) < v200 {
		v223 = v206
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v211 = int32(*(*int8)(unsafe.Add(mBase, uint32(v165+int32(-5)))))
	if int32(-1) < v211 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v221 = base.I64_extend_i32_u(v211&int32(127))<<(uint(int64(28))%64) | v206
	goto L59
L58:
	;
	v221 = int64(-1)
	goto L59
L59:
	;
	v223 = v221
	goto L52
L60:
	;
	v228 = v165
	goto L14
L61:
	;
	if l2 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v269 = v40 + int32(1)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v270 != 0 {
		v39 = v270
		v40 = v269
		goto L6
	} else {
		goto L67
	}
L63:
	;
	v262 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v39, v26, l3)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	return int32(0)
L65:
	;
	if v262 == int32(0) {
		v286 = v44
		goto L1
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	goto L7
L68:
	;
	v286 = base.B2i32(v26 == int32(65535)) | base.B2i32(v26 == v274)
	goto L1
}
func F_lpValidateNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int64
	_ = v131
	var v136 int32
	_ = v136
	var v142 int64
	_ = v142
	var v147 int32
	_ = v147
	var v153 int64
	_ = v153
	var v158 int32
	_ = v158
	var v164 int64
	_ = v164
	var v169 int32
	_ = v169
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	v4 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v13 == v4 {
		v197 = v4
		return v197
	} else {
		v17 = l0 + int32(6)
		if base.Ui32(v13) < base.Ui32(v17) {
			v197 = v4
			return v197
		} else {
			v19 = l0 + l2
			v21 = v19 + int32(-1)
			if base.Ui32(v21) < base.Ui32(v13) {
				v197 = v4
				return v197
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				if v23 != int32(255) {
					v32 = base.I32_extend8_s(v23)
					v33 = int32(192)
					v34 = v23 & v33
					v35 = int32(1)
					v37 = v23 & int32(224)
					if v37 == v33 {
						v58 = v35
						v59 = v13 + v58
						if base.Ui32(v59) < base.Ui32(v17) {
							v197 = v4
						} else {
							if base.Ui32(v21) < base.Ui32(v59) {
								v197 = v4
							} else {
								v62 = int32(1)
								if v32 <= int32(-1) {
									if v34 != int32(128) {
										if v37 != int32(192) {
											v79 = (v32 + int32(15)) & int32(255)
											if base.Ui32(v79) < base.Ui32(int32(4)) {
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v79<<(uint(int32(2))%32))+uint32(_consts[375])))
												v119 = v118
												v121 = v62
											} else {
												if v23&int32(240) != int32(224) {
													if v32 == int32(-16) {
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)+1))
														v100 = v97 + int32(5)
														if base.Ui32(v100) < base.Ui32(int32(128)) {
															v119 = v100
															v121 = v62
														} else {
															if base.Ui32(int32(16384)) <= base.Ui32(v100) {
																if base.Ui32(int32(2097152)) <= base.Ui32(v100) {
																	if base.Ui32(v100) < base.Ui32(int32(268435456)) {
																		v113 = int32(4)
																	} else {
																		v113 = int32(5)
																	}
																	v119 = v100
																	v121 = v113
																} else {
																	v119 = v100
																	v121 = int32(3)
																}
															} else {
																v119 = v100
																v121 = int32(2)
															}
														}
													} else {
														v119 = int32(0)
														v121 = v62
													}
												} else {
													v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
													v100 = v86 | v23<<(uint(int32(8))%32)&int32(3840) + int32(2)
													if base.Ui32(v100) < base.Ui32(int32(128)) {
														v119 = v100
														v121 = v62
													} else {
														if base.Ui32(int32(16384)) <= base.Ui32(v100) {
															if base.Ui32(int32(2097152)) <= base.Ui32(v100) {
																if base.Ui32(v100) < base.Ui32(int32(268435456)) {
																	v113 = int32(4)
																} else {
																	v113 = int32(5)
																}
																v119 = v100
																v121 = v113
															} else {
																v119 = v100
																v121 = int32(3)
															}
														} else {
															v119 = v100
															v121 = int32(2)
														}
													}
												}
											}
										} else {
											v119 = int32(2)
											v121 = v62
										}
									} else {
										v68 = int32(1)
										v119 = v23&int32(63) + v68
										v121 = v68
									}
								} else {
									v119 = int32(1)
									v121 = v62
								}
								v122 = v121 + v119
								v123 = v13 + v122
								if base.Ui32(v123) < base.Ui32(v17) {
									v197 = v4
								} else {
									if base.Ui32(v21) < base.Ui32(v123) {
										v197 = v4
									} else {
										v126 = int32(-1)
										v128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+v126))))
										v131 = base.I64_extend_i32_u(v128 & int32(127))
										if v126 < v128 {
											v181 = v131
										} else {
											v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-2)))))
											v142 = base.I64_extend_i32_u(v136&int32(127))<<(uint(int64(7))%64) | v131
											if int32(-1) < v136 {
												v181 = v142
											} else {
												v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-3)))))
												v153 = base.I64_extend_i32_u(v147&int32(127))<<(uint(int64(14))%64) | v142
												if int32(-1) < v147 {
													v181 = v153
												} else {
													v158 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-4)))))
													v164 = base.I64_extend_i32_u(v158&int32(127))<<(uint(int64(21))%64) | v153
													if int32(-1) < v158 {
														v181 = v164
													} else {
														v169 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-5)))))
														if int32(-1) < v169 {
															v179 = base.I64_extend_i32_u(v169&int32(127))<<(uint(int64(28))%64) | v164
														} else {
															v179 = int64(-1)
														}
														v181 = v179
													}
												}
											}
										}
										if v181+base.I64_extend_i32_u(v121) != base.I64_extend_i32_u(v122) {
											v197 = v4
										} else {
											v186 = v123
											*(*int32)(unsafe.Add(mBase, uint32(l1))) = v186
											v197 = int32(1)
										}
									}
								}
							}
						}
					} else {
						if int32(-1) < v32 {
							v58 = v35
							v59 = v13 + v58
							if base.Ui32(v59) < base.Ui32(v17) {
								v197 = v4
							} else {
								if base.Ui32(v21) < base.Ui32(v59) {
									v197 = v4
								} else {
									v62 = int32(1)
									if v32 <= int32(-1) {
										if v34 != int32(128) {
											if v37 != int32(192) {
												v79 = (v32 + int32(15)) & int32(255)
												if base.Ui32(v79) < base.Ui32(int32(4)) {
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v79<<(uint(int32(2))%32))+uint32(_consts[375])))
													v119 = v118
													v121 = v62
												} else {
													if v23&int32(240) != int32(224) {
														if v32 == int32(-16) {
															v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)+1))
															v100 = v97 + int32(5)
															if base.Ui32(v100) < base.Ui32(int32(128)) {
																v119 = v100
																v121 = v62
															} else {
																if base.Ui32(int32(16384)) <= base.Ui32(v100) {
																	if base.Ui32(int32(2097152)) <= base.Ui32(v100) {
																		if base.Ui32(v100) < base.Ui32(int32(268435456)) {
																			v113 = int32(4)
																		} else {
																			v113 = int32(5)
																		}
																		v119 = v100
																		v121 = v113
																	} else {
																		v119 = v100
																		v121 = int32(3)
																	}
																} else {
																	v119 = v100
																	v121 = int32(2)
																}
															}
														} else {
															v119 = int32(0)
															v121 = v62
														}
													} else {
														v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
														v100 = v86 | v23<<(uint(int32(8))%32)&int32(3840) + int32(2)
														if base.Ui32(v100) < base.Ui32(int32(128)) {
															v119 = v100
															v121 = v62
														} else {
															if base.Ui32(int32(16384)) <= base.Ui32(v100) {
																if base.Ui32(int32(2097152)) <= base.Ui32(v100) {
																	if base.Ui32(v100) < base.Ui32(int32(268435456)) {
																		v113 = int32(4)
																	} else {
																		v113 = int32(5)
																	}
																	v119 = v100
																	v121 = v113
																} else {
																	v119 = v100
																	v121 = int32(3)
																}
															} else {
																v119 = v100
																v121 = int32(2)
															}
														}
													}
												}
											} else {
												v119 = int32(2)
												v121 = v62
											}
										} else {
											v68 = int32(1)
											v119 = v23&int32(63) + v68
											v121 = v68
										}
									} else {
										v119 = int32(1)
										v121 = v62
									}
									v122 = v121 + v119
									v123 = v13 + v122
									if base.Ui32(v123) < base.Ui32(v17) {
										v197 = v4
									} else {
										if base.Ui32(v21) < base.Ui32(v123) {
											v197 = v4
										} else {
											v126 = int32(-1)
											v128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+v126))))
											v131 = base.I64_extend_i32_u(v128 & int32(127))
											if v126 < v128 {
												v181 = v131
											} else {
												v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-2)))))
												v142 = base.I64_extend_i32_u(v136&int32(127))<<(uint(int64(7))%64) | v131
												if int32(-1) < v136 {
													v181 = v142
												} else {
													v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-3)))))
													v153 = base.I64_extend_i32_u(v147&int32(127))<<(uint(int64(14))%64) | v142
													if int32(-1) < v147 {
														v181 = v153
													} else {
														v158 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-4)))))
														v164 = base.I64_extend_i32_u(v158&int32(127))<<(uint(int64(21))%64) | v153
														if int32(-1) < v158 {
															v181 = v164
														} else {
															v169 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-5)))))
															if int32(-1) < v169 {
																v179 = base.I64_extend_i32_u(v169&int32(127))<<(uint(int64(28))%64) | v164
															} else {
																v179 = int64(-1)
															}
															v181 = v179
														}
													}
												}
											}
											if v181+base.I64_extend_i32_u(v121) != base.I64_extend_i32_u(v122) {
												v197 = v4
											} else {
												v186 = v123
												*(*int32)(unsafe.Add(mBase, uint32(l1))) = v186
												v197 = int32(1)
											}
										}
									}
								}
							}
						} else {
							if base.Ui32((v32+int32(15))&int32(255)) < base.Ui32(int32(4)) {
								v58 = v35
								v59 = v13 + v58
								if base.Ui32(v59) < base.Ui32(v17) {
									v197 = v4
								} else {
									if base.Ui32(v21) < base.Ui32(v59) {
										v197 = v4
									} else {
										v62 = int32(1)
										if v32 <= int32(-1) {
											if v34 != int32(128) {
												if v37 != int32(192) {
													v79 = (v32 + int32(15)) & int32(255)
													if base.Ui32(v79) < base.Ui32(int32(4)) {
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v79<<(uint(int32(2))%32))+uint32(_consts[375])))
														v119 = v118
														v121 = v62
													} else {
														if v23&int32(240) != int32(224) {
															if v32 == int32(-16) {
																v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)+1))
																v100 = v97 + int32(5)
																if base.Ui32(v100) < base.Ui32(int32(128)) {
																	v119 = v100
																	v121 = v62
																} else {
																	if base.Ui32(int32(16384)) <= base.Ui32(v100) {
																		if base.Ui32(int32(2097152)) <= base.Ui32(v100) {
																			if base.Ui32(v100) < base.Ui32(int32(268435456)) {
																				v113 = int32(4)
																			} else {
																				v113 = int32(5)
																			}
																			v119 = v100
																			v121 = v113
																		} else {
																			v119 = v100
																			v121 = int32(3)
																		}
																	} else {
																		v119 = v100
																		v121 = int32(2)
																	}
																}
															} else {
																v119 = int32(0)
																v121 = v62
															}
														} else {
															v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
															v100 = v86 | v23<<(uint(int32(8))%32)&int32(3840) + int32(2)
															if base.Ui32(v100) < base.Ui32(int32(128)) {
																v119 = v100
																v121 = v62
															} else {
																if base.Ui32(int32(16384)) <= base.Ui32(v100) {
																	if base.Ui32(int32(2097152)) <= base.Ui32(v100) {
																		if base.Ui32(v100) < base.Ui32(int32(268435456)) {
																			v113 = int32(4)
																		} else {
																			v113 = int32(5)
																		}
																		v119 = v100
																		v121 = v113
																	} else {
																		v119 = v100
																		v121 = int32(3)
																	}
																} else {
																	v119 = v100
																	v121 = int32(2)
																}
															}
														}
													}
												} else {
													v119 = int32(2)
													v121 = v62
												}
											} else {
												v68 = int32(1)
												v119 = v23&int32(63) + v68
												v121 = v68
											}
										} else {
											v119 = int32(1)
											v121 = v62
										}
										v122 = v121 + v119
										v123 = v13 + v122
										if base.Ui32(v123) < base.Ui32(v17) {
											v197 = v4
										} else {
											if base.Ui32(v21) < base.Ui32(v123) {
												v197 = v4
											} else {
												v126 = int32(-1)
												v128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+v126))))
												v131 = base.I64_extend_i32_u(v128 & int32(127))
												if v126 < v128 {
													v181 = v131
												} else {
													v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-2)))))
													v142 = base.I64_extend_i32_u(v136&int32(127))<<(uint(int64(7))%64) | v131
													if int32(-1) < v136 {
														v181 = v142
													} else {
														v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-3)))))
														v153 = base.I64_extend_i32_u(v147&int32(127))<<(uint(int64(14))%64) | v142
														if int32(-1) < v147 {
															v181 = v153
														} else {
															v158 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-4)))))
															v164 = base.I64_extend_i32_u(v158&int32(127))<<(uint(int64(21))%64) | v153
															if int32(-1) < v158 {
																v181 = v164
															} else {
																v169 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-5)))))
																if int32(-1) < v169 {
																	v179 = base.I64_extend_i32_u(v169&int32(127))<<(uint(int64(28))%64) | v164
																} else {
																	v179 = int64(-1)
																}
																v181 = v179
															}
														}
													}
												}
												if v181+base.I64_extend_i32_u(v121) != base.I64_extend_i32_u(v122) {
													v197 = v4
												} else {
													v186 = v123
													*(*int32)(unsafe.Add(mBase, uint32(l1))) = v186
													v197 = int32(1)
												}
											}
										}
									}
								}
							} else {
								if v34 == int32(128) {
									v58 = v35
									v59 = v13 + v58
									if base.Ui32(v59) < base.Ui32(v17) {
										v197 = v4
									} else {
										if base.Ui32(v21) < base.Ui32(v59) {
											v197 = v4
										} else {
											v62 = int32(1)
											if v32 <= int32(-1) {
												if v34 != int32(128) {
													if v37 != int32(192) {
														v79 = (v32 + int32(15)) & int32(255)
														if base.Ui32(v79) < base.Ui32(int32(4)) {
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v79<<(uint(int32(2))%32))+uint32(_consts[375])))
															v119 = v118
															v121 = v62
														} else {
															if v23&int32(240) != int32(224) {
																if v32 == int32(-16) {
																	v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)+1))
																	v100 = v97 + int32(5)
																	if base.Ui32(v100) < base.Ui32(int32(128)) {
																		v119 = v100
																		v121 = v62
																	} else {
																		if base.Ui32(int32(16384)) <= base.Ui32(v100) {
																			if base.Ui32(int32(2097152)) <= base.Ui32(v100) {
																				if base.Ui32(v100) < base.Ui32(int32(268435456)) {
																					v113 = int32(4)
																				} else {
																					v113 = int32(5)
																				}
																				v119 = v100
																				v121 = v113
																			} else {
																				v119 = v100
																				v121 = int32(3)
																			}
																		} else {
																			v119 = v100
																			v121 = int32(2)
																		}
																	}
																} else {
																	v119 = int32(0)
																	v121 = v62
																}
															} else {
																v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
																v100 = v86 | v23<<(uint(int32(8))%32)&int32(3840) + int32(2)
																if base.Ui32(v100) < base.Ui32(int32(128)) {
																	v119 = v100
																	v121 = v62
																} else {
																	if base.Ui32(int32(16384)) <= base.Ui32(v100) {
																		if base.Ui32(int32(2097152)) <= base.Ui32(v100) {
																			if base.Ui32(v100) < base.Ui32(int32(268435456)) {
																				v113 = int32(4)
																			} else {
																				v113 = int32(5)
																			}
																			v119 = v100
																			v121 = v113
																		} else {
																			v119 = v100
																			v121 = int32(3)
																		}
																	} else {
																		v119 = v100
																		v121 = int32(2)
																	}
																}
															}
														}
													} else {
														v119 = int32(2)
														v121 = v62
													}
												} else {
													v68 = int32(1)
													v119 = v23&int32(63) + v68
													v121 = v68
												}
											} else {
												v119 = int32(1)
												v121 = v62
											}
											v122 = v121 + v119
											v123 = v13 + v122
											if base.Ui32(v123) < base.Ui32(v17) {
												v197 = v4
											} else {
												if base.Ui32(v21) < base.Ui32(v123) {
													v197 = v4
												} else {
													v126 = int32(-1)
													v128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+v126))))
													v131 = base.I64_extend_i32_u(v128 & int32(127))
													if v126 < v128 {
														v181 = v131
													} else {
														v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-2)))))
														v142 = base.I64_extend_i32_u(v136&int32(127))<<(uint(int64(7))%64) | v131
														if int32(-1) < v136 {
															v181 = v142
														} else {
															v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-3)))))
															v153 = base.I64_extend_i32_u(v147&int32(127))<<(uint(int64(14))%64) | v142
															if int32(-1) < v147 {
																v181 = v153
															} else {
																v158 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-4)))))
																v164 = base.I64_extend_i32_u(v158&int32(127))<<(uint(int64(21))%64) | v153
																if int32(-1) < v158 {
																	v181 = v164
																} else {
																	v169 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-5)))))
																	if int32(-1) < v169 {
																		v179 = base.I64_extend_i32_u(v169&int32(127))<<(uint(int64(28))%64) | v164
																	} else {
																		v179 = int64(-1)
																	}
																	v181 = v179
																}
															}
														}
													}
													if v181+base.I64_extend_i32_u(v121) != base.I64_extend_i32_u(v122) {
														v197 = v4
													} else {
														v186 = v123
														*(*int32)(unsafe.Add(mBase, uint32(l1))) = v186
														v197 = int32(1)
													}
												}
											}
										}
									}
								} else {
									if v23&int32(240) != int32(224) {
										if v32 != int32(-16) {
											v197 = v4
										} else {
											v58 = int32(5)
											v59 = v13 + v58
											if base.Ui32(v59) < base.Ui32(v17) {
												v197 = v4
											} else {
												if base.Ui32(v21) < base.Ui32(v59) {
													v197 = v4
												} else {
													v62 = int32(1)
													if v32 <= int32(-1) {
														if v34 != int32(128) {
															if v37 != int32(192) {
																v79 = (v32 + int32(15)) & int32(255)
																if base.Ui32(v79) < base.Ui32(int32(4)) {
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v79<<(uint(int32(2))%32))+uint32(_consts[375])))
																	v119 = v118
																	v121 = v62
																} else {
																	if v23&int32(240) != int32(224) {
																		if v32 == int32(-16) {
																			v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)+1))
																			v100 = v97 + int32(5)
																			if base.Ui32(v100) < base.Ui32(int32(128)) {
																				v119 = v100
																				v121 = v62
																			} else {
																				if base.Ui32(int32(16384)) <= base.Ui32(v100) {
																					if base.Ui32(int32(2097152)) <= base.Ui32(v100) {
																						if base.Ui32(v100) < base.Ui32(int32(268435456)) {
																							v113 = int32(4)
																						} else {
																							v113 = int32(5)
																						}
																						v119 = v100
																						v121 = v113
																					} else {
																						v119 = v100
																						v121 = int32(3)
																					}
																				} else {
																					v119 = v100
																					v121 = int32(2)
																				}
																			}
																		} else {
																			v119 = int32(0)
																			v121 = v62
																		}
																	} else {
																		v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
																		v100 = v86 | v23<<(uint(int32(8))%32)&int32(3840) + int32(2)
																		if base.Ui32(v100) < base.Ui32(int32(128)) {
																			v119 = v100
																			v121 = v62
																		} else {
																			if base.Ui32(int32(16384)) <= base.Ui32(v100) {
																				if base.Ui32(int32(2097152)) <= base.Ui32(v100) {
																					if base.Ui32(v100) < base.Ui32(int32(268435456)) {
																						v113 = int32(4)
																					} else {
																						v113 = int32(5)
																					}
																					v119 = v100
																					v121 = v113
																				} else {
																					v119 = v100
																					v121 = int32(3)
																				}
																			} else {
																				v119 = v100
																				v121 = int32(2)
																			}
																		}
																	}
																}
															} else {
																v119 = int32(2)
																v121 = v62
															}
														} else {
															v68 = int32(1)
															v119 = v23&int32(63) + v68
															v121 = v68
														}
													} else {
														v119 = int32(1)
														v121 = v62
													}
													v122 = v121 + v119
													v123 = v13 + v122
													if base.Ui32(v123) < base.Ui32(v17) {
														v197 = v4
													} else {
														if base.Ui32(v21) < base.Ui32(v123) {
															v197 = v4
														} else {
															v126 = int32(-1)
															v128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+v126))))
															v131 = base.I64_extend_i32_u(v128 & int32(127))
															if v126 < v128 {
																v181 = v131
															} else {
																v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-2)))))
																v142 = base.I64_extend_i32_u(v136&int32(127))<<(uint(int64(7))%64) | v131
																if int32(-1) < v136 {
																	v181 = v142
																} else {
																	v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-3)))))
																	v153 = base.I64_extend_i32_u(v147&int32(127))<<(uint(int64(14))%64) | v142
																	if int32(-1) < v147 {
																		v181 = v153
																	} else {
																		v158 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-4)))))
																		v164 = base.I64_extend_i32_u(v158&int32(127))<<(uint(int64(21))%64) | v153
																		if int32(-1) < v158 {
																			v181 = v164
																		} else {
																			v169 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-5)))))
																			if int32(-1) < v169 {
																				v179 = base.I64_extend_i32_u(v169&int32(127))<<(uint(int64(28))%64) | v164
																			} else {
																				v179 = int64(-1)
																			}
																			v181 = v179
																		}
																	}
																}
															}
															if v181+base.I64_extend_i32_u(v121) != base.I64_extend_i32_u(v122) {
																v197 = v4
															} else {
																v186 = v123
																*(*int32)(unsafe.Add(mBase, uint32(l1))) = v186
																v197 = int32(1)
															}
														}
													}
												}
											}
										}
									} else {
										v58 = int32(2)
										v59 = v13 + v58
										if base.Ui32(v59) < base.Ui32(v17) {
											v197 = v4
										} else {
											if base.Ui32(v21) < base.Ui32(v59) {
												v197 = v4
											} else {
												v62 = int32(1)
												if v32 <= int32(-1) {
													if v34 != int32(128) {
														if v37 != int32(192) {
															v79 = (v32 + int32(15)) & int32(255)
															if base.Ui32(v79) < base.Ui32(int32(4)) {
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v79<<(uint(int32(2))%32))+uint32(_consts[375])))
																v119 = v118
																v121 = v62
															} else {
																if v23&int32(240) != int32(224) {
																	if v32 == int32(-16) {
																		v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)+1))
																		v100 = v97 + int32(5)
																		if base.Ui32(v100) < base.Ui32(int32(128)) {
																			v119 = v100
																			v121 = v62
																		} else {
																			if base.Ui32(int32(16384)) <= base.Ui32(v100) {
																				if base.Ui32(int32(2097152)) <= base.Ui32(v100) {
																					if base.Ui32(v100) < base.Ui32(int32(268435456)) {
																						v113 = int32(4)
																					} else {
																						v113 = int32(5)
																					}
																					v119 = v100
																					v121 = v113
																				} else {
																					v119 = v100
																					v121 = int32(3)
																				}
																			} else {
																				v119 = v100
																				v121 = int32(2)
																			}
																		}
																	} else {
																		v119 = int32(0)
																		v121 = v62
																	}
																} else {
																	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
																	v100 = v86 | v23<<(uint(int32(8))%32)&int32(3840) + int32(2)
																	if base.Ui32(v100) < base.Ui32(int32(128)) {
																		v119 = v100
																		v121 = v62
																	} else {
																		if base.Ui32(int32(16384)) <= base.Ui32(v100) {
																			if base.Ui32(int32(2097152)) <= base.Ui32(v100) {
																				if base.Ui32(v100) < base.Ui32(int32(268435456)) {
																					v113 = int32(4)
																				} else {
																					v113 = int32(5)
																				}
																				v119 = v100
																				v121 = v113
																			} else {
																				v119 = v100
																				v121 = int32(3)
																			}
																		} else {
																			v119 = v100
																			v121 = int32(2)
																		}
																	}
																}
															}
														} else {
															v119 = int32(2)
															v121 = v62
														}
													} else {
														v68 = int32(1)
														v119 = v23&int32(63) + v68
														v121 = v68
													}
												} else {
													v119 = int32(1)
													v121 = v62
												}
												v122 = v121 + v119
												v123 = v13 + v122
												if base.Ui32(v123) < base.Ui32(v17) {
													v197 = v4
												} else {
													if base.Ui32(v21) < base.Ui32(v123) {
														v197 = v4
													} else {
														v126 = int32(-1)
														v128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+v126))))
														v131 = base.I64_extend_i32_u(v128 & int32(127))
														if v126 < v128 {
															v181 = v131
														} else {
															v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-2)))))
															v142 = base.I64_extend_i32_u(v136&int32(127))<<(uint(int64(7))%64) | v131
															if int32(-1) < v136 {
																v181 = v142
															} else {
																v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-3)))))
																v153 = base.I64_extend_i32_u(v147&int32(127))<<(uint(int64(14))%64) | v142
																if int32(-1) < v147 {
																	v181 = v153
																} else {
																	v158 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-4)))))
																	v164 = base.I64_extend_i32_u(v158&int32(127))<<(uint(int64(21))%64) | v153
																	if int32(-1) < v158 {
																		v181 = v164
																	} else {
																		v169 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123+int32(-5)))))
																		if int32(-1) < v169 {
																			v179 = base.I64_extend_i32_u(v169&int32(127))<<(uint(int64(28))%64) | v164
																		} else {
																			v179 = int64(-1)
																		}
																		v181 = v179
																	}
																}
															}
														}
														if v181+base.I64_extend_i32_u(v121) != base.I64_extend_i32_u(v122) {
															v197 = v4
														} else {
															v186 = v123
															*(*int32)(unsafe.Add(mBase, uint32(l1))) = v186
															v197 = int32(1)
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
					return v197
				} else {
					if v13+int32(1) == v19 {
						v186 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v186
						v197 = int32(1)
						return v197
					} else {
						return int32(0)
					}
				}
			}
		}
	}
}
