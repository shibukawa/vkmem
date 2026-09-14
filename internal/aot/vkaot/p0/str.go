package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_str_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v34 int32
	_ = v34
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
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v73 int32
	_ = v73
	v8 = m.G0
	v10 = v8 - int32(1040)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = (v12 - v13) >> (uint(int32(4)) % 32)
	goto L1
L1:
	;
	v18 = v10 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v10 + int32(16)
	goto L2
L2:
	;
	if v16 < int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_luaL_pushresult(m, v10+int32(4))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L8
	} else {
		goto L16
	}
L4:
	;
	v34 = int32(1)
	goto L5
L5:
	;
	v37 = F_luaL_checkinteger(m, l0, v34)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L3
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if base.Ui32(v48) < base.Ui32(v10+int32(1040)) {
		v55 = v48
		goto L12
	} else {
		goto L13
	}
L8:
	;
	return int32(0)
L9:
	;
	if base.Ui32(v37) < base.Ui32(int32(256)) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v43 = m.G3
	v46 = F_luaL_argerror(m, l0, v34, v43+int32(_a2322))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	v56 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v55 + v56
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v37)
	if v34 != v16 {
		v34 = v34 + v56
		goto L5
	} else {
		goto L15
	}
L13:
	;
	v52 = F_luaL_prepbuffer(m, v10+int32(4))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v55 = v54
	goto L12
L15:
	;
	goto L6
L16:
	;
	m.G0 = v10 + int32(1040)
	return int32(1)
}
func F_str_find_aux(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v502 int32
	_ = v502
	var v515 int32
	_ = v515
	v11 = m.G0
	v13 = v11 - int32(288)
	m.G0 = v13
	v18 = F_luaL_checklstring(m, l0, int32(1), v13+int32(284))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = F_luaL_checklstring(m, l0, int32(2), v13+int32(280))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = F_luaL_optinteger(m, l0, int32(3), int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)+284))
	v37 = v29 + v29>>(uint(int32(31))%32)&(v33+int32(1))
	v38 = int32(0)
	v40 = base.B2i32(v38 < v37)
	if v38 < v37 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v41 = v37
	goto L7
L6:
	;
	v41 = v38
	goto L7
L7:
	;
	v43 = v41 + int32(-1)
	if base.Ui32(v43) < base.Ui32(v33) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v45 = v43
	goto L10
L9:
	;
	v45 = v33
	goto L10
L10:
	;
	if v38 < v37 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v47 = v45
	goto L13
L12:
	;
	v47 = int32(0)
	goto L13
L13:
	;
	if l1 == int32(0) {
		v124 = v33
		goto L17
	} else {
		goto L18
	}
L14:
	;
	m.G0 = v13 + int32(288)
	return v515
L15:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v502)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v502 + int32(16)
	goto L128
L16:
	;
	v253 = v18 + v47
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v13)+280))
	if v254 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L17:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v18 + v124
	v132 = base.B2i32(v125 == int32(94))
	v139 = v18 + v47
	goto L44
L18:
	;
	goto L22
L19:
	;
	if v112 != 0 {
		goto L16
	} else {
		goto L38
	}
L20:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	switch v105 {
	case 0:
		v110 = v105
		goto L35
	case 1:
		goto L37
	default:
		goto L36
	}
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v60 = v55 + int32(48)
	v61 = m.G398
	if base.Ui32(v60) < base.Ui32(v54) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v63 = v60
	goto L25
L24:
	;
	v63 = v61
	goto L25
L25:
	;
	goto L20
L35:
	;
	v112 = v110
	goto L19
L36:
	;
	v110 = int32(1)
	goto L35
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v112 = base.B2i32(v106 != int32(0))
	goto L19
L38:
	;
	v113 = m.G3
	v116 = F_strcspn(m, v25, v113+int32(_a2323))
	mBase = m.M
	v117 = v25 + v116
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v119 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v120 == int32(0) {
		goto L16
	} else {
		goto L43
	}
L40:
	;
	v120 = v117
	goto L42
L41:
	;
	v120 = int32(0)
	goto L42
L42:
	;
	goto L39
L43:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v13)+284))
	v124 = v123
	goto L17
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(0)
	v148 = F_match(m, v13+int32(8), v139, v25+v132)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	if v125 == int32(94) {
		goto L15
	} else {
		goto L72
	}
L47:
	;
	if v148 == int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	if l1 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v220 = m.G3
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v219 != 0 {
		goto L60
	} else {
		goto L61
	}
L50:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v158))) = base.F64_convert_i32_s(v139 - v18 + int32(1))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v163 + int32(16)
	goto L51
L51:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v169))) = base.F64_convert_i32_s(v148 - v18)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v174 + int32(16)
	goto L52
L52:
	;
	v178 = m.G3
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_luaL_checkstack(m, v179, v180, v178+int32(_a2324))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v180 < int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v515 = v180 + int32(2)
	goto L14
L55:
	;
	v193 = int32(0)
	goto L56
L56:
	;
	v200 = int32(0)
	F_push_onecapture(m, v13+int32(8), v193, v200, v200)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	goto L54
L58:
	;
	v205 = v193 + int32(1)
	if v205 != v180 {
		v193 = v205
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v223 = v219
	goto L62
L61:
	;
	v223 = int32(1)
	goto L62
L62:
	;
	if v139 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v224 = v223
	goto L65
L64:
	;
	v224 = v219
	goto L65
L65:
	;
	F_luaL_checkstack(m, v221, v224, v220+int32(_a2324))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	if v224 < int32(1) {
		v515 = v219
		goto L14
	} else {
		goto L67
	}
L67:
	;
	v238 = int32(0)
	goto L68
L68:
	;
	F_push_onecapture(m, v13+int32(8), v238, v139, v148)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	v515 = v224
	goto L14
L70:
	;
	v247 = v238 + int32(1)
	if v247 != v224 {
		v238 = v247
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if base.Ui32(v139) < base.Ui32(v249) {
		v139 = v139 + int32(1)
		goto L44
	} else {
		goto L73
	}
L73:
	;
	goto L15
L74:
	;
	v465 = v462 - v18
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v469)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v469))) = base.F64_convert_i32_s(v465 + int32(1))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v474 + int32(16)
	goto L126
L75:
	;
	if v18 == int32(0) {
		goto L15
	} else {
		goto L125
	}
L76:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v13)+284))
	v258 = v257 - v47
	if base.Ui32(v258) < base.Ui32(v254) {
		goto L15
	} else {
		goto L77
	}
L77:
	;
	v261 = v254 + int32(-1)
	v262 = v258 - v261
	if v262 == int32(0) {
		goto L15
	} else {
		goto L78
	}
L78:
	;
	v266 = v25 + int32(1)
	v267 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25))))
	v272 = v253
	v273 = v262
	goto L79
L79:
	;
	v278 = int32(0)
	v281 = base.B2i32(v273 != v278)
	if v272&int32(3) == v278 {
		v307 = v272
		v309 = v273
		v310 = v281
		goto L84
	} else {
		goto L85
	}
L81:
	;
	if v380 == int32(0) {
		goto L15
	} else {
		goto L106
	}
L82:
	;
	v380 = int32(0)
	goto L81
L83:
	;
	v358 = v351
	v360 = v353
	goto L101
L84:
	;
	if v310 == int32(0) {
		goto L82
	} else {
		goto L92
	}
L85:
	;
	if v273 == int32(0) {
		v307 = v272
		v309 = v273
		v310 = v281
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v290 = v272
	v292 = v273
	goto L87
L87:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	if v295 == v267&int32(255) {
		v351 = v290
		v353 = v292
		goto L83
	} else {
		goto L89
	}
L88:
	;
	v307 = v302
	v309 = v298
	v310 = v300
	goto L84
L89:
	;
	v298 = v292 + int32(-1)
	v299 = int32(0)
	v300 = base.B2i32(v298 != v299)
	v302 = v290 + int32(1)
	if v302&int32(3) == v299 {
		v307 = v302
		v309 = v298
		v310 = v300
		goto L84
	} else {
		goto L90
	}
L90:
	;
	if v298 != 0 {
		v290 = v302
		v292 = v298
		goto L87
	} else {
		goto L91
	}
L91:
	;
	goto L88
L92:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	if v314 == v267&int32(255) {
		v344 = v307
		v346 = v309
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v346 == int32(0) {
		goto L82
	} else {
		goto L100
	}
L94:
	;
	if base.Ui32(v309) < base.Ui32(int32(4)) {
		v344 = v307
		v346 = v309
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v324 = v307
	v326 = v309
	goto L96
L96:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v331 = v330 ^ v267&int32(255)*int32(16843009)
	v334 = int32(-2139062144)
	if (int32(16843008)-v331|v331)&v334 != v334 {
		v351 = v324
		v353 = v326
		goto L83
	} else {
		goto L98
	}
L97:
	;
	v344 = v339
	v346 = v341
	goto L93
L98:
	;
	v339 = v324 + int32(4)
	v341 = v326 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v341) {
		v324 = v339
		v326 = v341
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v351 = v344
	v353 = v346
	goto L83
L101:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	if v363 != v267&int32(255) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L82
L103:
	;
	v368 = v360 + int32(-1)
	if v368 != 0 {
		v358 = v358 + int32(1)
		v360 = v368
		goto L101
	} else {
		goto L105
	}
L104:
	;
	v380 = v358
	goto L81
L105:
	;
	goto L102
L106:
	;
	v384 = v380 + int32(1)
	if base.Ui32(v261) < base.Ui32(int32(4)) {
		v408 = v384
		v409 = v266
		v410 = v261
		goto L110
	} else {
		goto L111
	}
L107:
	;
	if v448 == int32(0) {
		v462 = v380
		goto L74
	} else {
		goto L123
	}
L108:
	;
	v448 = int32(0)
	goto L107
L109:
	;
	v420 = v415
	v421 = v416
	v422 = v417
	goto L119
L110:
	;
	if v410 == int32(0) {
		goto L108
	} else {
		goto L117
	}
L111:
	;
	if (v266|v384)&int32(3) != 0 {
		v415 = v384
		v416 = v266
		v417 = v261
		goto L109
	} else {
		goto L112
	}
L112:
	;
	v392 = v384
	v393 = v266
	v394 = v261
	goto L113
L113:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	if v397 != v398 {
		v415 = v392
		v416 = v393
		v417 = v394
		goto L109
	} else {
		goto L115
	}
L114:
	;
	v408 = v403
	v409 = v401
	v410 = v405
	goto L110
L115:
	;
	v400 = int32(4)
	v401 = v393 + v400
	v403 = v392 + v400
	v405 = v394 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v405) {
		v392 = v403
		v393 = v401
		v394 = v405
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v415 = v408
	v416 = v409
	v417 = v410
	goto L109
L118:
	;
	v448 = v425 - v426
	goto L107
L119:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421))))
	if v425 != v426 {
		goto L118
	} else {
		goto L121
	}
L121:
	;
	v428 = int32(1)
	v433 = v422 + int32(-1)
	if v433 == int32(0) {
		goto L108
	} else {
		goto L122
	}
L122:
	;
	v420 = v420 + v428
	v421 = v421 + v428
	v422 = v433
	goto L119
L123:
	;
	v452 = v273 + v272 - v384
	if v452 != 0 {
		v272 = v384
		v273 = v452
		goto L79
	} else {
		goto L124
	}
L124:
	;
	goto L15
L125:
	;
	v462 = v253
	goto L74
L126:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v13)+280))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v481)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v481))) = base.F64_convert_i32_s(v465 + v478)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v486 + int32(16)
	goto L127
L127:
	;
	v515 = int32(2)
	goto L14
L128:
	;
	v515 = int32(1)
	goto L14
}
func F_str_len(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v10 = F_luaL_checklstring(m, l0, int32(1), v5+int32(12))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v16))) = base.F64_convert_i32_s(v14)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v21 + int32(16)
		m.G0 = v5 + int32(16)
		return int32(1)
	}
}
func F_str_match(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_str_find_aux(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_str_next(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 != 0 {
		v13 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
		if int32(-1) < v13 {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
			v137 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
			v138 = v137
		} else {
			v17 = v8 + int32(12)
			if l0 != 0 {
				if l1 == int32(0) {
					v114 = F___errno_location(m)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(25)
					v118 = int32(-1)
					v126 = v118
				} else {
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					v24 = base.I32_extend8_s(v23)
					if v24 < int32(0) {
						v32 = F___get_tp(m)
						mBase = m.M
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						if v34 != 0 {
							v43 = v23 + int32(-194)
							if base.Ui32(int32(50)) < base.Ui32(v43) {
								v114 = F___errno_location(m)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(25)
								v118 = int32(-1)
								v126 = v118
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v43<<(uint(int32(2))%32))+uint32(_consts[1236])))
								if base.Ui32(int32(3)) < base.Ui32(l1) {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
									v62 = int32(base.Ui32(v60) >> (uint(int32(3)) % 32))
									if base.Ui32(int32(7)) < base.Ui32(v62+int32(-16)|(v62+v50>>(uint(int32(26))%32))) {
										v114 = F___errno_location(m)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(25)
										v118 = int32(-1)
										v126 = v118
									} else {
										v75 = v60 + int32(-128) | v50<<(uint(int32(6))%32)
										if v75 < int32(0) {
											v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
											v85 = v83 + int32(-128)
											if base.Ui32(int32(63)) < base.Ui32(v85) {
												v114 = F___errno_location(m)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(25)
												v118 = int32(-1)
												v126 = v118
											} else {
												v89 = v75 << (uint(int32(6)) % 32)
												v90 = v85 | v89
												if v89 < int32(0) {
													v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
													v100 = v98 + int32(-128)
													if base.Ui32(int32(63)) < base.Ui32(v100) {
														v114 = F___errno_location(m)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(25)
														v118 = int32(-1)
														v126 = v118
													} else {
														if v17 == int32(0) {
															v118 = int32(4)
															v126 = v118
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v17))) = v100 | v90<<(uint(int32(6))%32)
															v126 = int32(4)
														}
													}
												} else {
													if v17 == int32(0) {
														v118 = int32(3)
														v126 = v118
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v17))) = v90
														v126 = int32(3)
													}
												}
											}
										} else {
											if v17 == int32(0) {
												v118 = int32(2)
												v126 = v118
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v17))) = v75
												v126 = int32(2)
											}
										}
									}
								} else {
									if v50<<(uint(l1*int32(6)+int32(-6))%32) < int32(0) {
										v114 = F___errno_location(m)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(25)
										v118 = int32(-1)
										v126 = v118
									} else {
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
										v62 = int32(base.Ui32(v60) >> (uint(int32(3)) % 32))
										if base.Ui32(int32(7)) < base.Ui32(v62+int32(-16)|(v62+v50>>(uint(int32(26))%32))) {
											v114 = F___errno_location(m)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(25)
											v118 = int32(-1)
											v126 = v118
										} else {
											v75 = v60 + int32(-128) | v50<<(uint(int32(6))%32)
											if v75 < int32(0) {
												v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
												v85 = v83 + int32(-128)
												if base.Ui32(int32(63)) < base.Ui32(v85) {
													v114 = F___errno_location(m)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(25)
													v118 = int32(-1)
													v126 = v118
												} else {
													v89 = v75 << (uint(int32(6)) % 32)
													v90 = v85 | v89
													if v89 < int32(0) {
														v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
														v100 = v98 + int32(-128)
														if base.Ui32(int32(63)) < base.Ui32(v100) {
															v114 = F___errno_location(m)
															mBase = m.M
															*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(25)
															v118 = int32(-1)
															v126 = v118
														} else {
															if v17 == int32(0) {
																v118 = int32(4)
																v126 = v118
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v17))) = v100 | v90<<(uint(int32(6))%32)
																v126 = int32(4)
															}
														}
													} else {
														if v17 == int32(0) {
															v118 = int32(3)
															v126 = v118
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v17))) = v90
															v126 = int32(3)
														}
													}
												}
											} else {
												if v17 == int32(0) {
													v118 = int32(2)
													v126 = v118
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v17))) = v75
													v126 = int32(2)
												}
											}
										}
									}
								}
							}
						} else {
							if v17 == int32(0) {
								v118 = int32(1)
								v126 = v118
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v17))) = v24 & int32(57343)
								v126 = int32(1)
							}
						}
					} else {
						if v17 == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v17))) = v23
						}
						v126 = base.B2i32(v24 != int32(0))
					}
				}
			} else {
				v126 = int32(0)
			}
			v127 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			v130 = base.B2i32(v126 < int32(0))
			if v126 < int32(0) {
				v131 = int32(1)
			} else {
				v131 = v126
			}
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v131
			if v126 < int32(0) {
				v134 = int32(-1)
			} else {
				v134 = v127
			}
			v138 = v134
		}
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v10
		v138 = v10
	}
	m.G0 = v8 + int32(16)
	return v138
}
func F_str_sub(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_luaL_checklstring(m, l0, int32(1), v10+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = F_luaL_checkinteger(m, l0, int32(2))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			v25 = F_luaL_optinteger(m, l0, int32(3), int32(-1))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				v33 = v25 + v25>>(uint(int32(31))%32)&(v29+int32(1))
				v34 = int32(0)
				if v34 < v33 {
					v37 = v33
				} else {
					v37 = v34
				}
				if v37 < v29 {
					v39 = v37
				} else {
					v39 = v29
				}
				v42 = int32(1)
				v45 = v20 + v20>>(uint(int32(31))%32)&(v22+v42)
				if v42 < v45 {
					v49 = v45
				} else {
					v49 = v42
				}
				if v39 < v49 {
					v59 = m.G3
					F_lua_pushlstring(m, l0, v59+int32(_a188), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(16)
						return int32(1)
					}
				} else {
					F_lua_pushlstring(m, l0, v15+v49+int32(-1), v39-v49+int32(1))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(16)
						return int32(1)
					}
				}
			}
		}
	}
}
