package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zslCreate(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_valkey_calloc(m, int32(272))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(1)
		return v3
	}
}
func F_zslDeleteRangeByLex(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	v15 = m.G0
	v17 = v15 - int32(128)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v19 < int32(1) {
		v173 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	if v183 != 0 {
		goto L49
	} else {
		goto L50
	}
L2:
	;
	v26 = l0
	v27 = v19
	goto L3
L3:
	;
	v37 = v27 + int32(-1)
	v39 = v37 << (uint(int32(3)) % 32)
	v47 = v26
	v51 = v26 + v39 + int32(12)
	goto L6
L4:
	;
	v173 = v47
	goto L1
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+v37<<(uint(int32(2))%32)))) = v47
	if int32(1) < v27 {
		v26 = v47
		v27 = v37
		goto L3
	} else {
		goto L47
	}
L6:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v57 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v61 = v57 + int32(16)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v65 = v61 + v62<<(uint(int32(3))%32)
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65))))
	v67 = v65 + v66
	v69 = v67 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v71 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L7
L10:
	;
	v47 = v149
	v51 = v149 + v39 + int32(12)
	goto L6
L11:
	;
	v96 = int32(0)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v96))))
	switch v103 & int32(7) {
	case 0:
		goto L33
	case 1:
		goto L32
	case 2:
		goto L31
	case 3:
		goto L30
	case 4:
		goto L29
	default:
		v120 = v96
		goto L28
	}
L12:
	;
	if v69 == v70 {
		goto L9
	} else {
		goto L21
	}
L13:
	;
	if v69 != v70 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v69 != v76 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v149 = v57
	goto L10
L16:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v70 == v79 {
		v149 = v57
		goto L10
	} else {
		goto L18
	}
L17:
	;
	v149 = v57
	goto L10
L18:
	;
	if v70 == v76 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	if v69 == v79 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v94 = int32(0)
	goto L11
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v69 != v86 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v70 == v89 {
		v149 = v57
		goto L10
	} else {
		goto L24
	}
L23:
	;
	v149 = v57
	goto L10
L24:
	;
	if v70 == v86 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	if v69 == v89 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v94 = int32(-1)
	goto L11
L27:
	;
	if v94 < v146 {
		goto L5
	} else {
		goto L46
	}
L28:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+int32(-1)))))
	switch v123 & int32(7) {
	case 0:
		goto L39
	case 1:
		goto L38
	case 2:
		goto L37
	case 3:
		goto L36
	case 4:
		goto L35
	default:
		v140 = v96
		goto L34
	}
L29:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(-16))))
	v120 = v119
	goto L28
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(-8))))
	v120 = v116
	goto L28
L31:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+int32(-4)))))
	v120 = v113
	goto L28
L32:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+int32(-2)))))
	v120 = v110
	goto L28
L33:
	;
	v120 = int32(base.Ui32(v103) >> (uint(int32(3)) % 32))
	goto L28
L34:
	;
	v141 = base.B2i32(base.Ui32(v120) < base.Ui32(v140))
	if base.Ui32(v120) < base.Ui32(v140) {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-17))))
	v140 = v139
	goto L34
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-9))))
	v140 = v136
	goto L34
L37:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70+int32(-5)))))
	v140 = v133
	goto L34
L38:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+int32(-3)))))
	v140 = v130
	goto L34
L39:
	;
	v140 = int32(base.Ui32(v123) >> (uint(int32(3)) % 32))
	goto L34
L40:
	;
	v142 = v120
	goto L42
L41:
	;
	v142 = v140
	goto L42
L42:
	;
	v143 = F_memcmp(m, v69, v70, v142)
	mBase = m.M
	if v143 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v146 = v143
	goto L45
L44:
	;
	v146 = base.B2i32(base.Ui32(v140) < base.Ui32(v120)) - v141
	goto L45
L45:
	;
	goto L27
L46:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v149 = v148
	goto L10
L47:
	;
	goto L4
L48:
	;
	m.G0 = v17 + int32(128)
	return v489
L49:
	;
	v194 = int32(0)
	v195 = v183
	goto L51
L50:
	;
	v489 = int32(0)
	goto L48
L51:
	;
	v201 = v195 + int32(16)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v205 = v201 + v202<<(uint(int32(3))%32)
	v206 = int32(*(*int8)(unsafe.Add(mBase, uint32(v205))))
	v207 = v205 + v206
	v209 = v207 + int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v211 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v489 = v480
	goto L48
L53:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v353 < int32(1) {
		v416 = v353
		goto L107
	} else {
		goto L108
	}
L54:
	;
	if v339 == int32(0) {
		v489 = v194
		goto L48
	} else {
		goto L105
	}
L55:
	;
	if v209 == v210 {
		goto L53
	} else {
		goto L81
	}
L56:
	;
	if v209 == v210 {
		v489 = v194
		goto L48
	} else {
		goto L57
	}
L57:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v209 == v216 {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v210 == v219 {
		goto L53
	} else {
		goto L59
	}
L59:
	;
	if v210 == v216 {
		v489 = v194
		goto L48
	} else {
		goto L60
	}
L60:
	;
	if v209 == v219 {
		v489 = v194
		goto L48
	} else {
		goto L61
	}
L61:
	;
	v223 = int32(0)
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v223))))
	switch v230 & int32(7) {
	case 0:
		goto L68
	case 1:
		goto L67
	case 2:
		goto L66
	case 3:
		goto L65
	case 4:
		goto L64
	default:
		v247 = v223
		goto L63
	}
L62:
	;
	v339 = int32(base.Ui32(v273) >> (uint(int32(31)) % 32))
	goto L54
L63:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210+int32(-1)))))
	switch v250 & int32(7) {
	case 0:
		goto L74
	case 1:
		goto L73
	case 2:
		goto L72
	case 3:
		goto L71
	case 4:
		goto L70
	default:
		v267 = v223
		goto L69
	}
L64:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-16))))
	v247 = v246
	goto L63
L65:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-8))))
	v247 = v243
	goto L63
L66:
	;
	v240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+int32(-4)))))
	v247 = v240
	goto L63
L67:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+int32(-2)))))
	v247 = v237
	goto L63
L68:
	;
	v247 = int32(base.Ui32(v230) >> (uint(int32(3)) % 32))
	goto L63
L69:
	;
	v268 = base.B2i32(base.Ui32(v247) < base.Ui32(v267))
	if base.Ui32(v247) < base.Ui32(v267) {
		goto L75
	} else {
		goto L76
	}
L70:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v210+int32(-17))))
	v267 = v266
	goto L69
L71:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v210+int32(-9))))
	v267 = v263
	goto L69
L72:
	;
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v210+int32(-5)))))
	v267 = v260
	goto L69
L73:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210+int32(-3)))))
	v267 = v257
	goto L69
L74:
	;
	v267 = int32(base.Ui32(v250) >> (uint(int32(3)) % 32))
	goto L69
L75:
	;
	v269 = v247
	goto L77
L76:
	;
	v269 = v267
	goto L77
L77:
	;
	v270 = F_memcmp(m, v209, v210, v269)
	mBase = m.M
	if v270 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v273 = v270
	goto L80
L79:
	;
	v273 = base.B2i32(base.Ui32(v267) < base.Ui32(v247)) - v268
	goto L80
L80:
	;
	goto L62
L81:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v209 == v278 {
		goto L53
	} else {
		goto L82
	}
L82:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v210 == v281 {
		goto L53
	} else {
		goto L83
	}
L83:
	;
	if v210 == v278 {
		v489 = v194
		goto L48
	} else {
		goto L84
	}
L84:
	;
	if v209 == v281 {
		v489 = v194
		goto L48
	} else {
		goto L85
	}
L85:
	;
	v285 = int32(0)
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v285))))
	switch v292 & int32(7) {
	case 0:
		goto L92
	case 1:
		goto L91
	case 2:
		goto L90
	case 3:
		goto L89
	case 4:
		goto L88
	default:
		v309 = v285
		goto L87
	}
L86:
	;
	v339 = base.B2i32(v335 < int32(1))
	goto L54
L87:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210+int32(-1)))))
	switch v312 & int32(7) {
	case 0:
		goto L98
	case 1:
		goto L97
	case 2:
		goto L96
	case 3:
		goto L95
	case 4:
		goto L94
	default:
		v329 = v285
		goto L93
	}
L88:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-16))))
	v309 = v308
	goto L87
L89:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-8))))
	v309 = v305
	goto L87
L90:
	;
	v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+int32(-4)))))
	v309 = v302
	goto L87
L91:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+int32(-2)))))
	v309 = v299
	goto L87
L92:
	;
	v309 = int32(base.Ui32(v292) >> (uint(int32(3)) % 32))
	goto L87
L93:
	;
	v330 = base.B2i32(base.Ui32(v309) < base.Ui32(v329))
	if base.Ui32(v309) < base.Ui32(v329) {
		goto L99
	} else {
		goto L100
	}
L94:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v210+int32(-17))))
	v329 = v328
	goto L93
L95:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v210+int32(-9))))
	v329 = v325
	goto L93
L96:
	;
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v210+int32(-5)))))
	v329 = v322
	goto L93
L97:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210+int32(-3)))))
	v329 = v319
	goto L93
L98:
	;
	v329 = int32(base.Ui32(v312) >> (uint(int32(3)) % 32))
	goto L93
L99:
	;
	v331 = v309
	goto L101
L100:
	;
	v331 = v329
	goto L101
L101:
	;
	v332 = F_memcmp(m, v209, v210, v331)
	mBase = m.M
	if v332 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v335 = v332
	goto L104
L103:
	;
	v335 = base.B2i32(base.Ui32(v329) < base.Ui32(v309)) - v330
	goto L104
L104:
	;
	goto L86
L105:
	;
	goto L53
L106:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v468 = v201 + v465<<(uint(int32(3))%32)
	v469 = int32(*(*int8)(unsafe.Add(mBase, uint32(v468))))
	v473 = F_hashtableDelete(m, l2, v468+v469+int32(1))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L128
	} else {
		goto L129
	}
L107:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	if v423 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L108:
	;
	v357 = v195 + int32(12)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+12))
	if v359 != v195 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v363 = int32(1)
	if v353 != v363 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v358)+12)) = v361
	goto L109
L111:
	;
	v374 = v363
	goto L113
L112:
	;
	v416 = int32(1)
	goto L107
L113:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v17+v374<<(uint(int32(2))%32))))
	v383 = v374 << (uint(int32(3)) % 32)
	v384 = v381 + v383
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)+12))
	if v385 != v195 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v416 = v411
	goto L107
L115:
	;
	v410 = v374 + int32(1)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v410 < v411 {
		v374 = v410
		goto L113
	} else {
		goto L118
	}
L116:
	;
	v402 = v384 + int32(16)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	*(*int32)(unsafe.Add(mBase, uint32(v402))) = v403 + int32(-1)
	goto L115
L117:
	;
	v388 = v384 + int32(16)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v195+int32(16)+v383)))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	*(*int32)(unsafe.Add(mBase, uint32(v388))) = v390 + v391 + int32(-1)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v357+v383)))
	*(*int32)(unsafe.Add(mBase, uint32(v384+int32(12)))) = v399
	goto L115
L118:
	;
	goto L114
L119:
	;
	if v429 < int32(2) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v422
	v429 = v416
	goto L119
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v423)+8)) = v422
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v429 = v427
	goto L119
L122:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v461 + int32(-1)
	goto L106
L123:
	;
	v437 = v429
	goto L124
L124:
	;
	v444 = v437 + int32(-1)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(12)+v444<<(uint(int32(3))%32))))
	if v448 != 0 {
		goto L122
	} else {
		goto L126
	}
L125:
	;
	goto L122
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v444
	if base.Ui32(int32(2)) < base.Ui32(v437) {
		v437 = v444
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	return int32(0)
L129:
	;
	F_valkey_free(m, v195)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v480 = v194 + int32(1)
	if v346 != 0 {
		v194 = v480
		v195 = v346
		goto L51
	} else {
		goto L131
	}
L131:
	;
	goto L52
}
func F_zslDeleteRangeByScore(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	v12 = m.G0
	v14 = v12 - int32(128)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v16 < int32(1) {
		v71 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	if v78 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v23 = l0
	v24 = v16
	goto L3
L3:
	;
	v31 = v24 + int32(-1)
	v42 = v23
	goto L6
L4:
	;
	v71 = v42
	goto L1
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v31<<(uint(int32(2))%32)))) = v42
	if int32(1) < v24 {
		v23 = v42
		v24 = v31
		goto L3
	} else {
		goto L13
	}
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+v31<<(uint(int32(3))%32)+int32(12))))
	if v48 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v51 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	v52 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v55 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v56 = base.F64_gt(v51, v52)
	goto L11
L10:
	;
	v56 = base.F64_ge(v51, v52)
	goto L11
L11:
	;
	if v56 != int32(1) {
		v42 = v48
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L7
L13:
	;
	goto L4
L14:
	;
	m.G0 = v14 + int32(128)
	return v245
L15:
	;
	v88 = int32(0)
	v89 = v78
	goto L17
L16:
	;
	v245 = int32(0)
	goto L14
L17:
	;
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v89)))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v96 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v245 = v237
	goto L14
L19:
	;
	v97 = base.F64_lt(v92, v93)
	goto L21
L20:
	;
	v97 = base.F64_le(v92, v93)
	goto L21
L21:
	;
	if v97 != int32(1) {
		v245 = v88
		goto L14
	} else {
		goto L22
	}
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v107 < int32(1) {
		v170 = v107
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v220 = v89 + int32(16)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v224 = v220 + v221<<(uint(int32(3))%32)
	v225 = int32(*(*int8)(unsafe.Add(mBase, uint32(v224))))
	v230 = F_hashtablePop(m, l2, v224+v225+int32(1), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L45
	} else {
		goto L46
	}
L24:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v177 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L25:
	;
	v111 = v89 + int32(12)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	if v113 != v89 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v117 = int32(1)
	if v107 != v117 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+12)) = v115
	goto L26
L28:
	;
	v128 = v117
	goto L30
L29:
	;
	v170 = int32(1)
	goto L24
L30:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v14+v128<<(uint(int32(2))%32))))
	v137 = v128 << (uint(int32(3)) % 32)
	v138 = v135 + v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	if v139 != v89 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v170 = v165
	goto L24
L32:
	;
	v164 = v128 + int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v164 < v165 {
		v128 = v164
		goto L30
	} else {
		goto L35
	}
L33:
	;
	v156 = v138 + int32(16)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v157 + int32(-1)
	goto L32
L34:
	;
	v142 = v138 + int32(16)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v89+int32(16)+v137)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v144 + v145 + int32(-1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v111+v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v138+int32(12)))) = v153
	goto L32
L35:
	;
	goto L31
L36:
	;
	if v183 < int32(2) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v176
	v183 = v170
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+8)) = v176
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v183 = v181
	goto L36
L39:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v215 + int32(-1)
	goto L23
L40:
	;
	v191 = v183
	goto L41
L41:
	;
	v198 = v191 + int32(-1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(12)+v198<<(uint(int32(3))%32))))
	if v202 != 0 {
		goto L39
	} else {
		goto L43
	}
L42:
	;
	goto L39
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v198
	if base.Ui32(int32(2)) < base.Ui32(v191) {
		v191 = v198
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	return int32(0)
L46:
	;
	F_valkey_free(m, v89)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v237 = v88 + int32(1)
	if v100 != 0 {
		v88 = v237
		v89 = v100
		goto L17
	} else {
		goto L48
	}
L48:
	;
	goto L18
}
func F_zslFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	v8 = v4
	goto L3
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	F_valkey_free(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	if v10 != 0 {
		v8 = v10
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	return
}
func F_zslInsertNode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v339 int32
	_ = v339
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	v18 = m.G0
	v20 = v18 - int32(256)
	m.G0 = v20
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v22&int64(9223372036854775807)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if v29 <= v225 {
		v339 = v225
		goto L46
	} else {
		goto L47
	}
L2:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v225 = v218
	goto L1
L3:
	;
	F__serverAssert(m, int32(_a1722), int32(_a1723), int32(259))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L44
	} else {
		goto L45
	}
L4:
	;
	v27 = int32(16)
	v28 = l1 + v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v31 = l0 + v27
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v32 <= int32(0) {
		v225 = v32
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = v32
	v42 = l0
	goto L6
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v41 == v53 {
		v59 = int32(0)
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v61 = v41 + int32(-1)
	v63 = v61 << (uint(int32(2)) % 32)
	v64 = v20 + v63
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v59
	v67 = v61 << (uint(int32(3)) % 32)
	v68 = v42 + v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	if v69 == l1 {
		v194 = v42
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v20+v41<<(uint(int32(2))%32))))
	v59 = v58
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(128)+v63))) = v194
	if v41 <= int32(1) {
		goto L2
	} else {
		goto L43
	}
L11:
	;
	if v69 == int32(0) {
		v194 = v42
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v82 = v42
	v83 = v59
	v88 = v68 + int32(12)
	v89 = v69
	goto L13
L13:
	;
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v89)))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_gt(v92, v93) != 0 {
		v194 = v82
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v194 = v178
	goto L10
L15:
	;
	if base.F64_lt(v92, v93) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v41 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L17:
	;
	v97 = v89 + int32(16)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v99 = int32(3)
	v101 = v97 + v98<<(uint(v99)%32)
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v101))))
	v103 = v101 + v102
	v104 = int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v109 = v28 + v106<<(uint(v99)%32)
	v110 = int32(*(*int8)(unsafe.Add(mBase, uint32(v109))))
	v111 = v109 + v110
	v114 = int32(0)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103+v114))))
	switch v121 & int32(7) {
	case 0:
		goto L24
	case 1:
		goto L23
	case 2:
		goto L22
	case 3:
		goto L21
	case 4:
		goto L20
	default:
		v138 = v114
		goto L19
	}
L18:
	;
	if int32(-1) < v164 {
		v194 = v82
		goto L10
	} else {
		goto L37
	}
L19:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+int32(0)))))
	switch v141 & int32(7) {
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
		v158 = v114
		goto L25
	}
L20:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v103+int32(-16))))
	v138 = v137
	goto L19
L21:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v103+int32(-8))))
	v138 = v134
	goto L19
L22:
	;
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103+int32(-4)))))
	v138 = v131
	goto L19
L23:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103+int32(-2)))))
	v138 = v128
	goto L19
L24:
	;
	v138 = int32(base.Ui32(v121) >> (uint(int32(3)) % 32))
	goto L19
L25:
	;
	v159 = base.B2i32(base.Ui32(v138) < base.Ui32(v158))
	if base.Ui32(v138) < base.Ui32(v158) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v111+int32(-16))))
	v158 = v157
	goto L25
L27:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v111+int32(-8))))
	v158 = v154
	goto L25
L28:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111+int32(-4)))))
	v158 = v151
	goto L25
L29:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+int32(-2)))))
	v158 = v148
	goto L25
L30:
	;
	v158 = int32(base.Ui32(v141) >> (uint(int32(3)) % 32))
	goto L25
L31:
	;
	v160 = v138
	goto L33
L32:
	;
	v160 = v158
	goto L33
L33:
	;
	v161 = F_memcmp(m, v103+v104, v111+v104, v160)
	mBase = m.M
	if v161 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v164 = v161
	goto L36
L35:
	;
	v164 = base.B2i32(base.Ui32(v158) < base.Ui32(v138)) - v159
	goto L36
L36:
	;
	goto L18
L37:
	;
	goto L16
L38:
	;
	v180 = v83 + v179
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v180
	v182 = v178 + v67
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	if v183 == l1 {
		v194 = v178
		goto L10
	} else {
		goto L41
	}
L39:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v178 = v175
	v179 = base.B2i32(v175 != int32(0))
	goto L38
L40:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v82+v67+int32(16))))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v178 = v174
	v179 = v173
	goto L38
L41:
	;
	if v183 != 0 {
		v82 = v178
		v83 = v180
		v88 = v182 + int32(12)
		v89 = v183
		goto L13
	} else {
		goto L42
	}
L42:
	;
	goto L14
L43:
	;
	v41 = v61
	v42 = v194
	goto L6
L44:
	;
	return int32(0)
L45:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	v350 = int32(1)
	if v29 < v350 {
		v415 = v339
		goto L62
	} else {
		goto L63
	}
L47:
	;
	v237 = int32(2)
	v238 = v225 << (uint(v237) % 32)
	v241 = v29 - v225
	v245 = F__emscripten_memset_bulkmem(m, v20+v238, base.I32_extend8_s(int32(0)), v241<<(uint(v237)%32))
	mBase = m.M
	goto L48
L48:
	;
	v246 = int32(1)
	if v241&v246 == int32(0) {
		v265 = v225
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v29 == v225+v246 {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(128)+v238))) = l0
	if v225 < int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v265 = v225 + int32(1)
	goto L49
L52:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v31+v225<<(uint(int32(3))%32)))) = v261
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v29
	v339 = v29
	goto L46
L54:
	;
	v273 = v265
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(128)+v273<<(uint(int32(2))%32)))) = l0
	if v273 < int32(1) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L53
L57:
	;
	v300 = v273 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(128)+v300<<(uint(int32(2))%32)))) = l0
	if v273 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v31+v273<<(uint(int32(3))%32)))) = v295
	goto L57
L59:
	;
	v313 = v273 + int32(2)
	if v313 != v29 {
		v273 = v313
		goto L55
	} else {
		goto L61
	}
L60:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v31+v300<<(uint(int32(3))%32)))) = v310
	goto L59
L61:
	;
	goto L56
L62:
	;
	if v415 <= v29 {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v20)+128))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v354)+12)) = l1
	if v29 == int32(1) {
		v415 = v339
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v378 = v350
	goto L65
L65:
	;
	v382 = v378 << (uint(int32(3)) % 32)
	v387 = v378 << (uint(int32(2)) % 32)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(128)+v387)))
	v390 = v389 + v382
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1+int32(12)+v382))) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v390)+12)) = l1
	v396 = v390 + int32(16)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v20+v387)))
	*(*int32)(unsafe.Add(mBase, uint32(v28+v382))) = v397 - v353 + v400
	*(*int32)(unsafe.Add(mBase, uint32(v396))) = v353 + int32(1) - v400
	v406 = v378 + int32(1)
	if v406 != v29 {
		v378 = v406
		goto L65
	} else {
		goto L67
	}
L66:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v415 = v408
	goto L62
L67:
	;
	goto L66
L68:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v20)+128))
	if v485 == l0 {
		goto L75
	} else {
		goto L76
	}
L69:
	;
	v431 = v29
	v433 = v415
	goto L70
L70:
	;
	if v431 < int32(1) {
		v462 = v433
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L68
L72:
	;
	v465 = v431 + int32(1)
	if v465 < v462 {
		v431 = v465
		v433 = v462
		goto L70
	} else {
		goto L74
	}
L73:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(128)+v431<<(uint(int32(2))%32))))
	v456 = v451 + v431<<(uint(int32(3))%32) + int32(16)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	*(*int32)(unsafe.Add(mBase, uint32(v456))) = v457 + int32(1)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v462 = v461
	goto L72
L74:
	;
	goto L71
L75:
	;
	v487 = int32(0)
	goto L77
L76:
	;
	v487 = v485
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v489 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v490 = v489
	goto L80
L79:
	;
	v490 = l0
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v490)+8)) = l1
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v492 + int32(1)
	m.G0 = v20 + int32(256)
	return l1
}
func F_zslNthInLexRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var __phi804 int32
	_ = __phi804
	var v805 int32
	_ = v805
	var __phi805 int32
	_ = __phi805
	var v806 int32
	_ = v806
	var __phi806 int32
	_ = __phi806
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v853 int32
	_ = v853
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v894 int32
	_ = v894
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v956 int32
	_ = v956
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1076 int32
	_ = v1076
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1138 int32
	_ = v1138
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1229 int32
	_ = v1229
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1271 int32
	_ = v1271
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1299 int32
	_ = v1299
	var __phi1299 int32
	_ = __phi1299
	var v1300 int32
	_ = v1300
	var __phi1300 int32
	_ = __phi1300
	var v1301 int32
	_ = v1301
	var __phi1301 int32
	_ = __phi1301
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1349 int32
	_ = v1349
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1402 int32
	_ = v1402
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1459 int32
	_ = v1459
	var v1469 int32
	_ = v1469
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v16 == v17 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v1469
L2:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v95 = v93 + int32(16)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v99 = v95 + v96<<(uint(int32(3))%32)
	v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v99))))
	v101 = v99 + v100
	v103 = v101 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v105 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L3:
	;
	v86 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v87 != 0 {
		v1469 = v86
		goto L1
	} else {
		goto L30
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v16 == v20 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v17 == v23 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v25 = int32(0)
	if v17 == v20 {
		v1469 = v25
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v16 == v23 {
		v1469 = v25
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v28 = int32(0)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-1)))))
	switch v36 & int32(7) {
	case 0:
		goto L15
	case 1:
		goto L14
	case 2:
		goto L13
	case 3:
		goto L12
	case 4:
		goto L11
	default:
		v53 = v28
		goto L10
	}
L9:
	;
	if int32(0) < v79 {
		v1469 = v28
		goto L1
	} else {
		goto L28
	}
L10:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-1)))))
	switch v56 & int32(7) {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		goto L17
	default:
		v73 = v28
		goto L16
	}
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
	v53 = v52
	goto L10
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
	v53 = v49
	goto L10
L13:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
	v53 = v46
	goto L10
L14:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
	v53 = v43
	goto L10
L15:
	;
	v53 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
	goto L10
L16:
	;
	v74 = base.B2i32(base.Ui32(v53) < base.Ui32(v73))
	if base.Ui32(v53) < base.Ui32(v73) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-17))))
	v73 = v72
	goto L16
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-9))))
	v73 = v69
	goto L16
L19:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(-5)))))
	v73 = v66
	goto L16
L20:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-3)))))
	v73 = v63
	goto L16
L21:
	;
	v73 = int32(base.Ui32(v56) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	v75 = v53
	goto L24
L23:
	;
	v75 = v73
	goto L24
L24:
	;
	v76 = F_memcmp(m, v16, v17, v75)
	mBase = m.M
	if v76 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v79 = v76
	goto L27
L26:
	;
	v79 = base.B2i32(base.Ui32(v73) < base.Ui32(v53)) - v74
	goto L27
L27:
	;
	goto L9
L28:
	;
	if v79 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	goto L3
L30:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v88 != 0 {
		v1469 = v86
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L2
L32:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v192 = v190 + int32(16)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v196 = v192 + v193<<(uint(int32(3))%32)
	v197 = int32(*(*int8)(unsafe.Add(mBase, uint32(v196))))
	v198 = v196 + v197
	v200 = v198 + int32(1)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v202 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L33:
	;
	v133 = int32(0)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v133))))
	switch v140 & int32(7) {
	case 0:
		goto L52
	case 1:
		goto L51
	case 2:
		goto L50
	case 3:
		goto L49
	case 4:
		goto L48
	default:
		v157 = v133
		goto L47
	}
L34:
	;
	if v103 == v104 {
		goto L32
	} else {
		goto L41
	}
L35:
	;
	v108 = int32(0)
	if v103 == v104 {
		v1469 = v108
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v103 == v111 {
		v1469 = v108
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v104 == v114 {
		v1469 = v108
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v104 == v111 {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	if v103 == v114 {
		goto L32
	} else {
		goto L40
	}
L40:
	;
	v132 = int32(0)
	goto L33
L41:
	;
	v120 = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v103 == v122 {
		v1469 = v120
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v104 == v125 {
		v1469 = v120
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v104 == v122 {
		goto L32
	} else {
		goto L44
	}
L44:
	;
	if v103 == v125 {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	v132 = int32(-1)
	goto L33
L46:
	;
	if v132 < v183 {
		goto L32
	} else {
		goto L65
	}
L47:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104+int32(-1)))))
	switch v160 & int32(7) {
	case 0:
		goto L58
	case 1:
		goto L57
	case 2:
		goto L56
	case 3:
		goto L55
	case 4:
		goto L54
	default:
		v177 = v133
		goto L53
	}
L48:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v101+int32(-16))))
	v157 = v156
	goto L47
L49:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v101+int32(-8))))
	v157 = v153
	goto L47
L50:
	;
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101+int32(-4)))))
	v157 = v150
	goto L47
L51:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+int32(-2)))))
	v157 = v147
	goto L47
L52:
	;
	v157 = int32(base.Ui32(v140) >> (uint(int32(3)) % 32))
	goto L47
L53:
	;
	v178 = base.B2i32(base.Ui32(v157) < base.Ui32(v177))
	if base.Ui32(v157) < base.Ui32(v177) {
		goto L59
	} else {
		goto L60
	}
L54:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v104+int32(-17))))
	v177 = v176
	goto L53
L55:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v104+int32(-9))))
	v177 = v173
	goto L53
L56:
	;
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104+int32(-5)))))
	v177 = v170
	goto L53
L57:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104+int32(-3)))))
	v177 = v167
	goto L53
L58:
	;
	v177 = int32(base.Ui32(v160) >> (uint(int32(3)) % 32))
	goto L53
L59:
	;
	v179 = v157
	goto L61
L60:
	;
	v179 = v177
	goto L61
L61:
	;
	v180 = F_memcmp(m, v103, v104, v179)
	mBase = m.M
	if v180 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v183 = v180
	goto L64
L63:
	;
	v183 = base.B2i32(base.Ui32(v177) < base.Ui32(v157)) - v178
	goto L64
L64:
	;
	goto L46
L65:
	;
	return int32(0)
L66:
	;
	v339 = int32(0)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v346 = (v342 + int32(-1)) << (uint(int32(3)) % 32)
	v347 = l0 + int32(12) + v346
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	if v348 == v339 {
		v482 = v339
		v483 = l0
		goto L119
	} else {
		goto L120
	}
L67:
	;
	if v333 != 0 {
		goto L66
	} else {
		goto L118
	}
L68:
	;
	if v200 == v201 {
		goto L66
	} else {
		goto L94
	}
L69:
	;
	v205 = int32(0)
	if v200 == v201 {
		v1469 = v205
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v200 == v208 {
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v201 == v211 {
		goto L66
	} else {
		goto L72
	}
L72:
	;
	if v201 == v208 {
		v1469 = v205
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v200 == v211 {
		v1469 = v205
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v215 = int32(0)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v215))))
	switch v222 & int32(7) {
	case 0:
		goto L81
	case 1:
		goto L80
	case 2:
		goto L79
	case 3:
		goto L78
	case 4:
		goto L77
	default:
		v239 = v215
		goto L76
	}
L75:
	;
	v333 = int32(base.Ui32(v265) >> (uint(int32(31)) % 32))
	goto L67
L76:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+int32(-1)))))
	switch v242 & int32(7) {
	case 0:
		goto L87
	case 1:
		goto L86
	case 2:
		goto L85
	case 3:
		goto L84
	case 4:
		goto L83
	default:
		v259 = v215
		goto L82
	}
L77:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v198+int32(-16))))
	v239 = v238
	goto L76
L78:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v198+int32(-8))))
	v239 = v235
	goto L76
L79:
	;
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198+int32(-4)))))
	v239 = v232
	goto L76
L80:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+int32(-2)))))
	v239 = v229
	goto L76
L81:
	;
	v239 = int32(base.Ui32(v222) >> (uint(int32(3)) % 32))
	goto L76
L82:
	;
	v260 = base.B2i32(base.Ui32(v239) < base.Ui32(v259))
	if base.Ui32(v239) < base.Ui32(v259) {
		goto L88
	} else {
		goto L89
	}
L83:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(-17))))
	v259 = v258
	goto L82
L84:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(-9))))
	v259 = v255
	goto L82
L85:
	;
	v252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201+int32(-5)))))
	v259 = v252
	goto L82
L86:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+int32(-3)))))
	v259 = v249
	goto L82
L87:
	;
	v259 = int32(base.Ui32(v242) >> (uint(int32(3)) % 32))
	goto L82
L88:
	;
	v261 = v239
	goto L90
L89:
	;
	v261 = v259
	goto L90
L90:
	;
	v262 = F_memcmp(m, v200, v201, v261)
	mBase = m.M
	if v262 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v265 = v262
	goto L93
L92:
	;
	v265 = base.B2i32(base.Ui32(v259) < base.Ui32(v239)) - v260
	goto L93
L93:
	;
	goto L75
L94:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v200 == v270 {
		goto L66
	} else {
		goto L95
	}
L95:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v201 == v273 {
		goto L66
	} else {
		goto L96
	}
L96:
	;
	v275 = int32(0)
	if v201 == v270 {
		v1469 = v275
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v200 == v273 {
		v1469 = v275
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v278 = int32(0)
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v278))))
	switch v285 & int32(7) {
	case 0:
		goto L105
	case 1:
		goto L104
	case 2:
		goto L103
	case 3:
		goto L102
	case 4:
		goto L101
	default:
		v302 = v278
		goto L100
	}
L99:
	;
	v333 = base.B2i32(v328 < int32(1))
	goto L67
L100:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+int32(-1)))))
	switch v305 & int32(7) {
	case 0:
		goto L111
	case 1:
		goto L110
	case 2:
		goto L109
	case 3:
		goto L108
	case 4:
		goto L107
	default:
		v322 = v278
		goto L106
	}
L101:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v198+int32(-16))))
	v302 = v301
	goto L100
L102:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v198+int32(-8))))
	v302 = v298
	goto L100
L103:
	;
	v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198+int32(-4)))))
	v302 = v295
	goto L100
L104:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+int32(-2)))))
	v302 = v292
	goto L100
L105:
	;
	v302 = int32(base.Ui32(v285) >> (uint(int32(3)) % 32))
	goto L100
L106:
	;
	v323 = base.B2i32(base.Ui32(v302) < base.Ui32(v322))
	if base.Ui32(v302) < base.Ui32(v322) {
		goto L112
	} else {
		goto L113
	}
L107:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(-17))))
	v322 = v321
	goto L106
L108:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(-9))))
	v322 = v318
	goto L106
L109:
	;
	v315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201+int32(-5)))))
	v322 = v315
	goto L106
L110:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+int32(-3)))))
	v322 = v312
	goto L106
L111:
	;
	v322 = int32(base.Ui32(v305) >> (uint(int32(3)) % 32))
	goto L106
L112:
	;
	v324 = v302
	goto L114
L113:
	;
	v324 = v322
	goto L114
L114:
	;
	v325 = F_memcmp(m, v200, v201, v324)
	mBase = m.M
	if v325 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v328 = v325
	goto L117
L116:
	;
	v328 = base.B2i32(base.Ui32(v322) < base.Ui32(v302)) - v323
	goto L117
L117:
	;
	goto L99
L118:
	;
	return int32(0)
L119:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if l2 < int32(0) {
		goto L162
	} else {
		goto L163
	}
L120:
	;
	v358 = v347
	v361 = v348
	v362 = int32(0)
	v363 = l0
	goto L121
L121:
	;
	v370 = v361 + int32(16)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	v374 = v370 + v371<<(uint(int32(3))%32)
	v375 = int32(*(*int8)(unsafe.Add(mBase, uint32(v374))))
	v376 = v374 + v375
	v378 = v376 + int32(1)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v380 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	v482 = v469
	v483 = v468
	goto L119
L123:
	;
	if v342 < int32(2) {
		goto L158
	} else {
		goto L159
	}
L124:
	;
	v405 = int32(0)
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376+v405))))
	switch v412 & int32(7) {
	case 0:
		goto L143
	case 1:
		goto L142
	case 2:
		goto L141
	case 3:
		goto L140
	case 4:
		goto L139
	default:
		v429 = v405
		goto L138
	}
L125:
	;
	if v378 == v379 {
		v482 = v362
		v483 = v363
		goto L119
	} else {
		goto L132
	}
L126:
	;
	if v378 == v379 {
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v378 == v385 {
		goto L123
	} else {
		goto L128
	}
L128:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v379 == v388 {
		goto L123
	} else {
		goto L129
	}
L129:
	;
	if v379 == v385 {
		v482 = v362
		v483 = v363
		goto L119
	} else {
		goto L130
	}
L130:
	;
	if v378 == v388 {
		v482 = v362
		v483 = v363
		goto L119
	} else {
		goto L131
	}
L131:
	;
	v403 = int32(0)
	goto L124
L132:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v378 == v395 {
		goto L123
	} else {
		goto L133
	}
L133:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v379 == v398 {
		goto L123
	} else {
		goto L134
	}
L134:
	;
	if v379 == v395 {
		v482 = v362
		v483 = v363
		goto L119
	} else {
		goto L135
	}
L135:
	;
	if v378 == v398 {
		v482 = v362
		v483 = v363
		goto L119
	} else {
		goto L136
	}
L136:
	;
	v403 = int32(-1)
	goto L124
L137:
	;
	if v403 < v455 {
		v482 = v362
		v483 = v363
		goto L119
	} else {
		goto L156
	}
L138:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379+int32(-1)))))
	switch v432 & int32(7) {
	case 0:
		goto L149
	case 1:
		goto L148
	case 2:
		goto L147
	case 3:
		goto L146
	case 4:
		goto L145
	default:
		v449 = v405
		goto L144
	}
L139:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v376+int32(-16))))
	v429 = v428
	goto L138
L140:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v376+int32(-8))))
	v429 = v425
	goto L138
L141:
	;
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v376+int32(-4)))))
	v429 = v422
	goto L138
L142:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376+int32(-2)))))
	v429 = v419
	goto L138
L143:
	;
	v429 = int32(base.Ui32(v412) >> (uint(int32(3)) % 32))
	goto L138
L144:
	;
	v450 = base.B2i32(base.Ui32(v429) < base.Ui32(v449))
	if base.Ui32(v429) < base.Ui32(v449) {
		goto L150
	} else {
		goto L151
	}
L145:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v379+int32(-17))))
	v449 = v448
	goto L144
L146:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v379+int32(-9))))
	v449 = v445
	goto L144
L147:
	;
	v442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v379+int32(-5)))))
	v449 = v442
	goto L144
L148:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379+int32(-3)))))
	v449 = v439
	goto L144
L149:
	;
	v449 = int32(base.Ui32(v432) >> (uint(int32(3)) % 32))
	goto L144
L150:
	;
	v451 = v429
	goto L152
L151:
	;
	v451 = v449
	goto L152
L152:
	;
	v452 = F_memcmp(m, v378, v379, v451)
	mBase = m.M
	if v452 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v455 = v452
	goto L155
L154:
	;
	v455 = base.B2i32(base.Ui32(v449) < base.Ui32(v429)) - v450
	goto L155
L155:
	;
	goto L137
L156:
	;
	goto L123
L157:
	;
	v469 = v467 + v362
	v472 = v468 + v346 + int32(12)
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	if v473 != 0 {
		v358 = v472
		v361 = v473
		v362 = v469
		v363 = v468
		goto L121
	} else {
		goto L160
	}
L158:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	v467 = base.B2i32(v464 != int32(0))
	v468 = v464
	goto L157
L159:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v363+v346+int32(16))))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	v467 = v462
	v468 = v463
	goto L157
L160:
	;
	goto L122
L161:
	;
	v1469 = v1459
	goto L1
L162:
	;
	if v489 < int32(1) {
		v1219 = v483
		v1220 = v482
		goto L300
	} else {
		goto L301
	}
L163:
	;
	if v489 < int32(2) {
		v664 = v483
		v666 = v482
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v675 = int32(0)
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v677) <= base.Ui32(v666+l2) {
		v1469 = v675
		goto L1
	} else {
		goto L211
	}
L165:
	;
	v500 = v483
	v502 = v482
	v503 = v489 + int32(-2)
	goto L166
L166:
	;
	v512 = v503 << (uint(int32(3)) % 32)
	v513 = v500 + v512
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)+12))
	if v514 == int32(0) {
		v645 = v500
		v647 = v502
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v664 = v645
	v666 = v647
	goto L164
L168:
	;
	if int32(0) < v503 {
		v500 = v645
		v502 = v647
		v503 = v503 + int32(-1)
		goto L166
	} else {
		goto L210
	}
L169:
	;
	v523 = v500
	v524 = v513 + int32(12)
	v525 = v502
	v526 = v514
	goto L170
L170:
	;
	v535 = v526 + int32(16)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)))
	v539 = v535 + v536<<(uint(int32(3))%32)
	v540 = int32(*(*int8)(unsafe.Add(mBase, uint32(v539))))
	v541 = v539 + v540
	v543 = v541 + int32(1)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v545 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L171:
	;
	v645 = v634
	v647 = v636
	goto L168
L172:
	;
	if v503 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L173:
	;
	v570 = int32(0)
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541+v570))))
	switch v577 & int32(7) {
	case 0:
		goto L192
	case 1:
		goto L191
	case 2:
		goto L190
	case 3:
		goto L189
	case 4:
		goto L188
	default:
		v594 = v570
		goto L187
	}
L174:
	;
	if v543 == v544 {
		v645 = v523
		v647 = v525
		goto L168
	} else {
		goto L181
	}
L175:
	;
	if v543 == v544 {
		goto L172
	} else {
		goto L176
	}
L176:
	;
	v550 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v543 == v550 {
		goto L172
	} else {
		goto L177
	}
L177:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v544 == v553 {
		goto L172
	} else {
		goto L178
	}
L178:
	;
	if v544 == v550 {
		v645 = v523
		v647 = v525
		goto L168
	} else {
		goto L179
	}
L179:
	;
	if v543 == v553 {
		v645 = v523
		v647 = v525
		goto L168
	} else {
		goto L180
	}
L180:
	;
	v568 = int32(0)
	goto L173
L181:
	;
	v560 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v543 == v560 {
		goto L172
	} else {
		goto L182
	}
L182:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v544 == v563 {
		goto L172
	} else {
		goto L183
	}
L183:
	;
	if v544 == v560 {
		v645 = v523
		v647 = v525
		goto L168
	} else {
		goto L184
	}
L184:
	;
	if v543 == v563 {
		v645 = v523
		v647 = v525
		goto L168
	} else {
		goto L185
	}
L185:
	;
	v568 = int32(-1)
	goto L173
L186:
	;
	if v568 < v620 {
		v645 = v523
		v647 = v525
		goto L168
	} else {
		goto L205
	}
L187:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544+int32(-1)))))
	switch v597 & int32(7) {
	case 0:
		goto L198
	case 1:
		goto L197
	case 2:
		goto L196
	case 3:
		goto L195
	case 4:
		goto L194
	default:
		v614 = v570
		goto L193
	}
L188:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v541+int32(-16))))
	v594 = v593
	goto L187
L189:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v541+int32(-8))))
	v594 = v590
	goto L187
L190:
	;
	v587 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v541+int32(-4)))))
	v594 = v587
	goto L187
L191:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541+int32(-2)))))
	v594 = v584
	goto L187
L192:
	;
	v594 = int32(base.Ui32(v577) >> (uint(int32(3)) % 32))
	goto L187
L193:
	;
	v615 = base.B2i32(base.Ui32(v594) < base.Ui32(v614))
	if base.Ui32(v594) < base.Ui32(v614) {
		goto L199
	} else {
		goto L200
	}
L194:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v544+int32(-17))))
	v614 = v613
	goto L193
L195:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v544+int32(-9))))
	v614 = v610
	goto L193
L196:
	;
	v607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v544+int32(-5)))))
	v614 = v607
	goto L193
L197:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544+int32(-3)))))
	v614 = v604
	goto L193
L198:
	;
	v614 = int32(base.Ui32(v597) >> (uint(int32(3)) % 32))
	goto L193
L199:
	;
	v616 = v594
	goto L201
L200:
	;
	v616 = v614
	goto L201
L201:
	;
	v617 = F_memcmp(m, v543, v544, v616)
	mBase = m.M
	if v617 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v620 = v617
	goto L204
L203:
	;
	v620 = base.B2i32(base.Ui32(v614) < base.Ui32(v594)) - v615
	goto L204
L204:
	;
	goto L186
L205:
	;
	goto L172
L206:
	;
	v636 = v635 + v525
	v637 = v634 + v512
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v637)+12))
	if v640 != 0 {
		v523 = v634
		v524 = v637 + int32(12)
		v525 = v636
		v526 = v640
		goto L170
	} else {
		goto L209
	}
L207:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	v634 = v631
	v635 = base.B2i32(v631 != int32(0))
	goto L206
L208:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v523+v512+int32(16))))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	v634 = v630
	v635 = v629
	goto L206
L209:
	;
	goto L171
L210:
	;
	goto L167
L211:
	;
	if int32(9) < l2 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v865 = v853 + int32(16)
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v865)))
	v869 = v865 + v866<<(uint(int32(3))%32)
	v870 = int32(*(*int8)(unsafe.Add(mBase, uint32(v869))))
	v871 = v869 + v870
	v873 = v871 + int32(1)
	v874 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v875 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L213:
	;
	v769 = int32(0)
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v770 <= v769 {
		v1459 = v769
		goto L161
	} else {
		goto L226
	}
L214:
	;
	v682 = l2 + int32(1)
	v683 = int32(7)
	v684 = v682 & v683
	if base.Ui32(l2) < base.Ui32(v683) {
		v720 = v664
		goto L215
	} else {
		goto L216
	}
L215:
	;
	if v684 == int32(0) {
		v757 = v720
		goto L220
	} else {
		goto L221
	}
L216:
	;
	v693 = int32(0)
	v694 = v664
	goto L217
L217:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v694)+12))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+12))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v706)+12))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)+12))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v708)+12))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v709)+12))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v710)+12))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v711)+12))
	v714 = v693 + int32(8)
	if v714 != v682&int32(-8) {
		v693 = v714
		v694 = v712
		goto L217
	} else {
		goto L219
	}
L218:
	;
	v720 = v712
	goto L215
L219:
	;
	goto L218
L220:
	;
	if v757 != 0 {
		v853 = v757
		goto L212
	} else {
		goto L225
	}
L221:
	;
	v737 = int32(0)
	v738 = v720
	goto L222
L222:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v738)+12))
	v751 = v737 + int32(1)
	if v751 != v684 {
		v737 = v751
		v738 = v749
		goto L222
	} else {
		goto L224
	}
L223:
	;
	v757 = v749
	goto L220
L224:
	;
	goto L223
L225:
	;
	v1459 = int32(0)
	goto L161
L226:
	;
	v776 = l2 - v482 + v666 + int32(1)
	v782 = v483
	v783 = int32(0)
	v792 = v770
	goto L228
L227:
	;
	if v806 != v776 {
		v1459 = v769
		goto L161
	} else {
		goto L242
	}
L228:
	;
	v794 = v792 + int32(-1)
	v796 = v794 << (uint(int32(3)) % 32)
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v782+v796)+12))
	if v798 == int32(0) {
		v834 = v782
		v835 = v783
		goto L230
	} else {
		goto L231
	}
L230:
	;
	if v835 == v776 {
		v853 = v834
		goto L212
	} else {
		goto L240
	}
L231:
	;
	__phi804 = v798
	__phi805 = v782
	__phi806 = v783
	v804 = __phi804
	v805 = __phi805
	v806 = __phi806
	goto L232
L232:
	;
	if v794 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	v834 = v804
	v835 = v827
	goto L230
L234:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v804+v796)+12))
	if v829 != 0 {
		__phi804 = v829
		__phi805 = v804
		__phi806 = v827
		v804 = __phi804
		v805 = __phi805
		v806 = __phi806
		goto L232
	} else {
		goto L239
	}
L235:
	;
	v825 = v806 + int32(1)
	if base.Ui32(v776) < base.Ui32(v825) {
		goto L227
	} else {
		goto L238
	}
L236:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v805+v796+int32(16))))
	v822 = v821 + v806
	if base.Ui32(v822) <= base.Ui32(v776) {
		v827 = v822
		goto L234
	} else {
		goto L237
	}
L237:
	;
	v834 = v805
	v835 = v806
	goto L230
L238:
	;
	v827 = v825
	goto L234
L239:
	;
	goto L233
L240:
	;
	if v792 < int32(2) {
		v1459 = v769
		goto L161
	} else {
		goto L241
	}
L241:
	;
	v782 = v834
	v783 = v835
	v792 = v794
	goto L228
L242:
	;
	v853 = v805
	goto L212
L243:
	;
	if v1002 != 0 {
		v1459 = v853
		goto L161
	} else {
		goto L299
	}
L244:
	;
	if v873 != v874 {
		goto L272
	} else {
		goto L273
	}
L245:
	;
	if v873 == v874 {
		v1469 = v675
		goto L1
	} else {
		goto L246
	}
L246:
	;
	v880 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v873 != v880 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v883 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v874 != v883 {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v1459 = v853
	goto L161
L249:
	;
	if v874 == v880 {
		v1469 = v675
		goto L1
	} else {
		goto L251
	}
L250:
	;
	v1459 = v853
	goto L161
L251:
	;
	if v873 == v883 {
		v1469 = v675
		goto L1
	} else {
		goto L252
	}
L252:
	;
	v887 = int32(0)
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871+v887))))
	switch v894 & int32(7) {
	case 0:
		goto L259
	case 1:
		goto L258
	case 2:
		goto L257
	case 3:
		goto L256
	case 4:
		goto L255
	default:
		v911 = v887
		goto L254
	}
L253:
	;
	v1002 = int32(base.Ui32(v937) >> (uint(int32(31)) % 32))
	goto L243
L254:
	;
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874+int32(-1)))))
	switch v914 & int32(7) {
	case 0:
		goto L265
	case 1:
		goto L264
	case 2:
		goto L263
	case 3:
		goto L262
	case 4:
		goto L261
	default:
		v931 = v887
		goto L260
	}
L255:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v871+int32(-16))))
	v911 = v910
	goto L254
L256:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v871+int32(-8))))
	v911 = v907
	goto L254
L257:
	;
	v904 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v871+int32(-4)))))
	v911 = v904
	goto L254
L258:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871+int32(-2)))))
	v911 = v901
	goto L254
L259:
	;
	v911 = int32(base.Ui32(v894) >> (uint(int32(3)) % 32))
	goto L254
L260:
	;
	v932 = base.B2i32(base.Ui32(v911) < base.Ui32(v931))
	if base.Ui32(v911) < base.Ui32(v931) {
		goto L266
	} else {
		goto L267
	}
L261:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v874+int32(-17))))
	v931 = v930
	goto L260
L262:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v874+int32(-9))))
	v931 = v927
	goto L260
L263:
	;
	v924 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v874+int32(-5)))))
	v931 = v924
	goto L260
L264:
	;
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874+int32(-3)))))
	v931 = v921
	goto L260
L265:
	;
	v931 = int32(base.Ui32(v914) >> (uint(int32(3)) % 32))
	goto L260
L266:
	;
	v933 = v911
	goto L268
L267:
	;
	v933 = v931
	goto L268
L268:
	;
	v934 = F_memcmp(m, v873, v874, v933)
	mBase = m.M
	if v934 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v937 = v934
	goto L271
L270:
	;
	v937 = base.B2i32(base.Ui32(v931) < base.Ui32(v911)) - v932
	goto L271
L271:
	;
	goto L253
L272:
	;
	v942 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v873 != v942 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	v1459 = v853
	goto L161
L274:
	;
	v945 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v874 != v945 {
		goto L276
	} else {
		goto L277
	}
L275:
	;
	v1459 = v853
	goto L161
L276:
	;
	if v874 == v942 {
		v1469 = v675
		goto L1
	} else {
		goto L278
	}
L277:
	;
	v1459 = v853
	goto L161
L278:
	;
	if v873 == v945 {
		v1469 = v675
		goto L1
	} else {
		goto L279
	}
L279:
	;
	v949 = int32(0)
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871+v949))))
	switch v956 & int32(7) {
	case 0:
		goto L286
	case 1:
		goto L285
	case 2:
		goto L284
	case 3:
		goto L283
	case 4:
		goto L282
	default:
		v973 = v949
		goto L281
	}
L280:
	;
	v1002 = base.B2i32(v999 < int32(1))
	goto L243
L281:
	;
	v976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874+int32(-1)))))
	switch v976 & int32(7) {
	case 0:
		goto L292
	case 1:
		goto L291
	case 2:
		goto L290
	case 3:
		goto L289
	case 4:
		goto L288
	default:
		v993 = v949
		goto L287
	}
L282:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v871+int32(-16))))
	v973 = v972
	goto L281
L283:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v871+int32(-8))))
	v973 = v969
	goto L281
L284:
	;
	v966 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v871+int32(-4)))))
	v973 = v966
	goto L281
L285:
	;
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871+int32(-2)))))
	v973 = v963
	goto L281
L286:
	;
	v973 = int32(base.Ui32(v956) >> (uint(int32(3)) % 32))
	goto L281
L287:
	;
	v994 = base.B2i32(base.Ui32(v973) < base.Ui32(v993))
	if base.Ui32(v973) < base.Ui32(v993) {
		goto L293
	} else {
		goto L294
	}
L288:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v874+int32(-17))))
	v993 = v992
	goto L287
L289:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v874+int32(-9))))
	v993 = v989
	goto L287
L290:
	;
	v986 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v874+int32(-5)))))
	v993 = v986
	goto L287
L291:
	;
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874+int32(-3)))))
	v993 = v983
	goto L287
L292:
	;
	v993 = int32(base.Ui32(v976) >> (uint(int32(3)) % 32))
	goto L287
L293:
	;
	v995 = v973
	goto L295
L294:
	;
	v995 = v993
	goto L295
L295:
	;
	v996 = F_memcmp(m, v873, v874, v995)
	mBase = m.M
	if v996 != 0 {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v999 = v996
	goto L298
L297:
	;
	v999 = base.B2i32(base.Ui32(v993) < base.Ui32(v973)) - v994
	goto L298
L298:
	;
	goto L280
L299:
	;
	v1469 = v675
	goto L1
L300:
	;
	v1229 = int32(0)
	if v1220 < v1229-l2 {
		v1469 = v1229
		goto L1
	} else {
		goto L365
	}
L301:
	;
	v1011 = v483
	v1012 = v482
	v1018 = v489
	goto L302
L302:
	;
	v1022 = v1018 + int32(-1)
	v1024 = v1022 << (uint(int32(3)) % 32)
	v1032 = v1011 + v1024 + int32(12)
	v1033 = v1011
	v1034 = v1012
	goto L305
L303:
	;
	v1219 = v1033
	v1220 = v1034
	goto L300
L304:
	;
	if int32(1) < v1018 {
		v1011 = v1033
		v1012 = v1034
		v1018 = v1022
		goto L302
	} else {
		goto L364
	}
L305:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1032)))
	if v1043 == int32(0) {
		goto L304
	} else {
		goto L307
	}
L307:
	;
	v1047 = v1043 + int32(16)
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1047)))
	v1051 = v1047 + v1048<<(uint(int32(3))%32)
	v1052 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1051))))
	v1053 = v1051 + v1052
	v1055 = v1053 + int32(1)
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1057 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L308:
	;
	if v1018 == int32(1) {
		goto L362
	} else {
		goto L363
	}
L309:
	;
	if v1184 == int32(0) {
		goto L304
	} else {
		goto L360
	}
L310:
	;
	if v1055 == v1056 {
		goto L308
	} else {
		goto L336
	}
L311:
	;
	if v1055 == v1056 {
		goto L304
	} else {
		goto L312
	}
L312:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v1055 == v1062 {
		goto L308
	} else {
		goto L313
	}
L313:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v1056 == v1065 {
		goto L308
	} else {
		goto L314
	}
L314:
	;
	if v1056 == v1062 {
		goto L304
	} else {
		goto L315
	}
L315:
	;
	if v1055 == v1065 {
		goto L304
	} else {
		goto L316
	}
L316:
	;
	v1069 = int32(0)
	v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053+v1069))))
	switch v1076 & int32(7) {
	case 0:
		goto L323
	case 1:
		goto L322
	case 2:
		goto L321
	case 3:
		goto L320
	case 4:
		goto L319
	default:
		v1093 = v1069
		goto L318
	}
L317:
	;
	v1184 = int32(base.Ui32(v1119) >> (uint(int32(31)) % 32))
	goto L309
L318:
	;
	v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056+int32(-1)))))
	switch v1096 & int32(7) {
	case 0:
		goto L329
	case 1:
		goto L328
	case 2:
		goto L327
	case 3:
		goto L326
	case 4:
		goto L325
	default:
		v1113 = v1069
		goto L324
	}
L319:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1053+int32(-16))))
	v1093 = v1092
	goto L318
L320:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1053+int32(-8))))
	v1093 = v1089
	goto L318
L321:
	;
	v1086 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1053+int32(-4)))))
	v1093 = v1086
	goto L318
L322:
	;
	v1083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053+int32(-2)))))
	v1093 = v1083
	goto L318
L323:
	;
	v1093 = int32(base.Ui32(v1076) >> (uint(int32(3)) % 32))
	goto L318
L324:
	;
	v1114 = base.B2i32(base.Ui32(v1093) < base.Ui32(v1113))
	if base.Ui32(v1093) < base.Ui32(v1113) {
		goto L330
	} else {
		goto L331
	}
L325:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1056+int32(-17))))
	v1113 = v1112
	goto L324
L326:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1056+int32(-9))))
	v1113 = v1109
	goto L324
L327:
	;
	v1106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1056+int32(-5)))))
	v1113 = v1106
	goto L324
L328:
	;
	v1103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056+int32(-3)))))
	v1113 = v1103
	goto L324
L329:
	;
	v1113 = int32(base.Ui32(v1096) >> (uint(int32(3)) % 32))
	goto L324
L330:
	;
	v1115 = v1093
	goto L332
L331:
	;
	v1115 = v1113
	goto L332
L332:
	;
	v1116 = F_memcmp(m, v1055, v1056, v1115)
	mBase = m.M
	if v1116 != 0 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v1119 = v1116
	goto L335
L334:
	;
	v1119 = base.B2i32(base.Ui32(v1113) < base.Ui32(v1093)) - v1114
	goto L335
L335:
	;
	goto L317
L336:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v1055 == v1124 {
		goto L308
	} else {
		goto L337
	}
L337:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v1056 == v1127 {
		goto L308
	} else {
		goto L338
	}
L338:
	;
	if v1056 == v1124 {
		goto L304
	} else {
		goto L339
	}
L339:
	;
	if v1055 == v1127 {
		goto L304
	} else {
		goto L340
	}
L340:
	;
	v1131 = int32(0)
	v1138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053+v1131))))
	switch v1138 & int32(7) {
	case 0:
		goto L347
	case 1:
		goto L346
	case 2:
		goto L345
	case 3:
		goto L344
	case 4:
		goto L343
	default:
		v1155 = v1131
		goto L342
	}
L341:
	;
	v1184 = base.B2i32(v1181 < int32(1))
	goto L309
L342:
	;
	v1158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056+int32(-1)))))
	switch v1158 & int32(7) {
	case 0:
		goto L353
	case 1:
		goto L352
	case 2:
		goto L351
	case 3:
		goto L350
	case 4:
		goto L349
	default:
		v1175 = v1131
		goto L348
	}
L343:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1053+int32(-16))))
	v1155 = v1154
	goto L342
L344:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1053+int32(-8))))
	v1155 = v1151
	goto L342
L345:
	;
	v1148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1053+int32(-4)))))
	v1155 = v1148
	goto L342
L346:
	;
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053+int32(-2)))))
	v1155 = v1145
	goto L342
L347:
	;
	v1155 = int32(base.Ui32(v1138) >> (uint(int32(3)) % 32))
	goto L342
L348:
	;
	v1176 = base.B2i32(base.Ui32(v1155) < base.Ui32(v1175))
	if base.Ui32(v1155) < base.Ui32(v1175) {
		goto L354
	} else {
		goto L355
	}
L349:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1056+int32(-17))))
	v1175 = v1174
	goto L348
L350:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1056+int32(-9))))
	v1175 = v1171
	goto L348
L351:
	;
	v1168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1056+int32(-5)))))
	v1175 = v1168
	goto L348
L352:
	;
	v1165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056+int32(-3)))))
	v1175 = v1165
	goto L348
L353:
	;
	v1175 = int32(base.Ui32(v1158) >> (uint(int32(3)) % 32))
	goto L348
L354:
	;
	v1177 = v1155
	goto L356
L355:
	;
	v1177 = v1175
	goto L356
L356:
	;
	v1178 = F_memcmp(m, v1055, v1056, v1177)
	mBase = m.M
	if v1178 != 0 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1181 = v1178
	goto L359
L358:
	;
	v1181 = base.B2i32(base.Ui32(v1175) < base.Ui32(v1155)) - v1176
	goto L359
L359:
	;
	goto L341
L360:
	;
	goto L308
L361:
	;
	v1032 = v1202 + v1024 + int32(12)
	v1033 = v1202
	v1034 = v1203 + v1034
	goto L305
L362:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1032)))
	v1202 = v1199
	v1203 = base.B2i32(v1199 != int32(0))
	goto L361
L363:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1033+v1024+int32(16))))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1032)))
	v1202 = v1198
	v1203 = v1197
	goto L361
L364:
	;
	goto L303
L365:
	;
	if l2 < int32(-10) {
		goto L367
	} else {
		goto L368
	}
L366:
	;
	v1360 = v1349 + int32(16)
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v1360)))
	v1364 = v1360 + v1361<<(uint(int32(3))%32)
	v1365 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1364))))
	v1366 = v1364 + v1365
	v1368 = v1366 + int32(1)
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1370 == int32(0) {
		goto L395
	} else {
		goto L396
	}
L367:
	;
	v1264 = int32(0)
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1265 <= v1264 {
		v1459 = v1264
		goto L161
	} else {
		goto L377
	}
L368:
	;
	if l2 == int32(-1) {
		v1349 = v1219
		goto L366
	} else {
		goto L369
	}
L369:
	;
	v1237 = int32(-2)
	if l2 < v1237 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v1240 = l2
	goto L372
L371:
	;
	v1240 = v1237
	goto L372
L372:
	;
	v1245 = int32(0)
	v1247 = v1219
	goto L373
L373:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+8))
	v1259 = v1245 + int32(1)
	if v1240^v1259 != int32(-1) {
		v1245 = v1259
		v1247 = v1257
		goto L373
	} else {
		goto L375
	}
L374:
	;
	if v1257 != 0 {
		v1349 = v1257
		goto L366
	} else {
		goto L376
	}
L375:
	;
	goto L374
L376:
	;
	v1459 = int32(0)
	goto L161
L377:
	;
	v1271 = l2 - v482 + v1220 + int32(1)
	v1277 = int32(0)
	v1278 = v483
	v1287 = v1265
	goto L379
L378:
	;
	if v1300 != v1271 {
		v1459 = v1264
		goto L161
	} else {
		goto L393
	}
L379:
	;
	v1289 = v1287 + int32(-1)
	v1291 = v1289 << (uint(int32(3)) % 32)
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1278+v1291)+12))
	if v1293 == int32(0) {
		v1329 = v1277
		v1330 = v1278
		goto L381
	} else {
		goto L382
	}
L381:
	;
	if v1329 == v1271 {
		v1349 = v1330
		goto L366
	} else {
		goto L391
	}
L382:
	;
	__phi1299 = v1293
	__phi1300 = v1277
	__phi1301 = v1278
	v1299 = __phi1299
	v1300 = __phi1300
	v1301 = __phi1301
	goto L383
L383:
	;
	if v1289 == int32(0) {
		goto L386
	} else {
		goto L387
	}
L384:
	;
	v1329 = v1322
	v1330 = v1299
	goto L381
L385:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1299+v1291)+12))
	if v1324 != 0 {
		__phi1299 = v1324
		__phi1300 = v1322
		__phi1301 = v1299
		v1299 = __phi1299
		v1300 = __phi1300
		v1301 = __phi1301
		goto L383
	} else {
		goto L390
	}
L386:
	;
	v1320 = v1300 + int32(1)
	if base.Ui32(v1271) < base.Ui32(v1320) {
		goto L378
	} else {
		goto L389
	}
L387:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1301+v1291+int32(16))))
	v1317 = v1316 + v1300
	if base.Ui32(v1317) <= base.Ui32(v1271) {
		v1322 = v1317
		goto L385
	} else {
		goto L388
	}
L388:
	;
	v1329 = v1300
	v1330 = v1301
	goto L381
L389:
	;
	v1322 = v1320
	goto L385
L390:
	;
	goto L384
L391:
	;
	if v1287 < int32(2) {
		v1459 = v1264
		goto L161
	} else {
		goto L392
	}
L392:
	;
	v1277 = v1329
	v1278 = v1330
	v1287 = v1289
	goto L379
L393:
	;
	v1349 = v1301
	goto L366
L394:
	;
	v1395 = int32(0)
	v1402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1366+v1395))))
	switch v1402 & int32(7) {
	case 0:
		goto L418
	case 1:
		goto L417
	case 2:
		goto L416
	case 3:
		goto L415
	case 4:
		goto L414
	default:
		v1419 = v1395
		goto L413
	}
L395:
	;
	if v1368 != v1369 {
		goto L404
	} else {
		goto L405
	}
L396:
	;
	if v1368 == v1369 {
		v1469 = v1229
		goto L1
	} else {
		goto L397
	}
L397:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v1368 == v1375 {
		v1469 = v1229
		goto L1
	} else {
		goto L398
	}
L398:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v1369 == v1378 {
		v1469 = v1229
		goto L1
	} else {
		goto L399
	}
L399:
	;
	if v1369 != v1375 {
		goto L400
	} else {
		goto L401
	}
L400:
	;
	if v1368 != v1378 {
		goto L402
	} else {
		goto L403
	}
L401:
	;
	v1459 = v1349
	goto L161
L402:
	;
	v1393 = int32(0)
	goto L394
L403:
	;
	v1459 = v1349
	goto L161
L404:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v1368 == v1385 {
		v1469 = v1229
		goto L1
	} else {
		goto L406
	}
L405:
	;
	v1459 = v1349
	goto L161
L406:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	if v1369 == v1388 {
		v1469 = v1229
		goto L1
	} else {
		goto L407
	}
L407:
	;
	if v1369 != v1385 {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	if v1368 != v1388 {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	v1459 = v1349
	goto L161
L410:
	;
	v1393 = int32(-1)
	goto L394
L411:
	;
	v1459 = v1349
	goto L161
L412:
	;
	if v1445 <= v1393 {
		v1469 = v1229
		goto L1
	} else {
		goto L431
	}
L413:
	;
	v1422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1369+int32(-1)))))
	switch v1422 & int32(7) {
	case 0:
		goto L424
	case 1:
		goto L423
	case 2:
		goto L422
	case 3:
		goto L421
	case 4:
		goto L420
	default:
		v1439 = v1395
		goto L419
	}
L414:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v1366+int32(-16))))
	v1419 = v1418
	goto L413
L415:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1366+int32(-8))))
	v1419 = v1415
	goto L413
L416:
	;
	v1412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1366+int32(-4)))))
	v1419 = v1412
	goto L413
L417:
	;
	v1409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1366+int32(-2)))))
	v1419 = v1409
	goto L413
L418:
	;
	v1419 = int32(base.Ui32(v1402) >> (uint(int32(3)) % 32))
	goto L413
L419:
	;
	v1440 = base.B2i32(base.Ui32(v1419) < base.Ui32(v1439))
	if base.Ui32(v1419) < base.Ui32(v1439) {
		goto L425
	} else {
		goto L426
	}
L420:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1369+int32(-17))))
	v1439 = v1438
	goto L419
L421:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v1369+int32(-9))))
	v1439 = v1435
	goto L419
L422:
	;
	v1432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1369+int32(-5)))))
	v1439 = v1432
	goto L419
L423:
	;
	v1429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1369+int32(-3)))))
	v1439 = v1429
	goto L419
L424:
	;
	v1439 = int32(base.Ui32(v1422) >> (uint(int32(3)) % 32))
	goto L419
L425:
	;
	v1441 = v1419
	goto L427
L426:
	;
	v1441 = v1439
	goto L427
L427:
	;
	v1442 = F_memcmp(m, v1368, v1369, v1441)
	mBase = m.M
	if v1442 != 0 {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1445 = v1442
	goto L430
L429:
	;
	v1445 = base.B2i32(base.Ui32(v1439) < base.Ui32(v1419)) - v1440
	goto L430
L430:
	;
	goto L412
L431:
	;
	v1459 = v1349
	goto L161
}
func F_zslNthInRange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 float64
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 float64
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var __phi62 int32
	_ = __phi62
	var v74 int32
	_ = v74
	var __phi74 int32
	_ = __phi74
	var v77 int32
	_ = v77
	var __phi77 int32
	_ = __phi77
	var v81 float64
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var __phi160 int32
	_ = __phi160
	var v169 int32
	_ = v169
	var __phi169 int32
	_ = __phi169
	var v171 int32
	_ = v171
	var __phi171 int32
	_ = __phi171
	var v173 float64
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v230 int32
	_ = v230
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var __phi308 int32
	_ = __phi308
	var v314 int32
	_ = v314
	var __phi314 int32
	_ = __phi314
	var v322 int32
	_ = v322
	var __phi322 int32
	_ = __phi322
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v356 int32
	_ = v356
	var v368 int32
	_ = v368
	var v387 float64
	_ = v387
	var v390 int32
	_ = v390
	var v410 int32
	_ = v410
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v453 int32
	_ = v453
	var __phi453 int32
	_ = __phi453
	var v462 int32
	_ = v462
	var __phi462 int32
	_ = __phi462
	var v464 int32
	_ = v464
	var __phi464 int32
	_ = __phi464
	var v466 float64
	_ = v466
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v497 int32
	_ = v497
	var v506 int32
	_ = v506
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var __phi674 int32
	_ = __phi674
	var v680 int32
	_ = v680
	var __phi680 int32
	_ = __phi680
	var v689 int32
	_ = v689
	var __phi689 int32
	_ = __phi689
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v723 int32
	_ = v723
	var v734 int32
	_ = v734
	var v753 float64
	_ = v753
	var v756 int32
	_ = v756
	var v776 int32
	_ = v776
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v819 int32
	_ = v819
	var v826 int32
	_ = v826
	v5 = int32(0)
	v22 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v23 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_gt(v22, v23) != 0 {
		v826 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v826
L2:
	;
	if base.F64_ne(v22, v23) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v28 == int32(0) {
		v826 = v5
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v26 != 0 {
		v826 = v5
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v27 != 0 {
		v826 = v5
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	v31 = *(*float64)(unsafe.Add(mBase, uint32(v28)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v35 = base.F64_gt(v31, v22)
	goto L10
L9:
	;
	v35 = base.F64_ge(v31, v22)
	goto L10
L10:
	;
	if v35 != int32(1) {
		v826 = v5
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v38 == int32(0) {
		v826 = v5
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v41 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v44 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v45 = base.F64_lt(v41, v23)
	goto L15
L14:
	;
	v45 = base.F64_le(v41, v23)
	goto L15
L15:
	;
	if v45 != int32(1) {
		v826 = v5
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v52 = v50 + int32(-1)
	v54 = v52 << (uint(int32(3)) % 32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(12)+v54)))
	if v56 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if int32(-1) < l2 {
		goto L32
	} else {
		goto L33
	}
L18:
	;
	__phi62 = v56
	__phi74 = int32(0)
	__phi77 = l0
	v62 = __phi62
	v74 = __phi74
	v77 = __phi77
	goto L20
L19:
	;
	v111 = int32(0)
	v112 = l0
	goto L17
L20:
	;
	v81 = *(*float64)(unsafe.Add(mBase, uint32(v62)))
	if v34 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v111 = v93
	v112 = v62
	goto L17
L22:
	;
	if v50 < int32(2) {
		v92 = int32(1)
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v84 = base.F64_gt(v81, v22)
	goto L25
L24:
	;
	v84 = base.F64_ge(v81, v22)
	goto L25
L25:
	;
	if v84 == int32(0) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v111 = v74
	v112 = v77
	goto L17
L27:
	;
	v93 = v92 + v74
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v62+v54+int32(12))))
	if v97 != 0 {
		__phi62 = v97
		__phi74 = v93
		__phi77 = v62
		v62 = __phi62
		v74 = __phi74
		v77 = __phi77
		goto L20
	} else {
		goto L29
	}
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v77+v54+int32(16))))
	v92 = v91
	goto L27
L29:
	;
	goto L21
L30:
	;
	v826 = v819
	goto L1
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v800
	v819 = v798
	goto L30
L32:
	;
	if v50 < int32(2) {
		v506 = v112
		v521 = v111
		goto L86
	} else {
		goto L87
	}
L33:
	;
	v120 = int32(0)
	if v50 <= v120 {
		v215 = v112
		v230 = v111
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v230 < int32(0)-l2 {
		v826 = v120
		goto L1
	} else {
		goto L51
	}
L35:
	;
	v124 = v112
	v139 = v111
	v140 = v50
	goto L36
L36:
	;
	v144 = v140 + int32(-1)
	v146 = v144 << (uint(int32(3)) % 32)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v124+v146+int32(12))))
	if v150 == int32(0) {
		v193 = v124
		v208 = v139
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v215 = v193
	v230 = v208
	goto L34
L38:
	;
	if int32(1) < v140 {
		v124 = v193
		v139 = v208
		v140 = v144
		goto L36
	} else {
		goto L50
	}
L39:
	;
	__phi160 = v150
	__phi169 = v139
	__phi171 = v124
	v160 = __phi160
	v169 = __phi169
	v171 = __phi171
	goto L40
L40:
	;
	v173 = *(*float64)(unsafe.Add(mBase, uint32(v160)))
	if v44 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v193 = v160
	v208 = v187
	goto L38
L42:
	;
	v179 = int32(1)
	if v140 == v179 {
		v186 = v179
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v176 = base.F64_lt(v173, v23)
	goto L45
L44:
	;
	v176 = base.F64_le(v173, v23)
	goto L45
L45:
	;
	if v176 == int32(1) {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v193 = v171
	v208 = v169
	goto L38
L47:
	;
	v187 = v186 + v169
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v160+v146+int32(12))))
	if v191 != 0 {
		__phi160 = v191
		__phi169 = v187
		__phi171 = v160
		v160 = __phi160
		v169 = __phi169
		v171 = __phi171
		goto L40
	} else {
		goto L49
	}
L48:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v171+v146+int32(16))))
	v186 = v185
	goto L47
L49:
	;
	goto L41
L50:
	;
	goto L37
L51:
	;
	if l2 < int32(-10) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	if l3 == int32(0) {
		v819 = v410
		goto L30
	} else {
		goto L85
	}
L53:
	;
	v387 = *(*float64)(unsafe.Add(mBase, uint32(v368)))
	if v34 != 0 {
		goto L81
	} else {
		goto L82
	}
L54:
	;
	v273 = int32(0)
	if v50 <= v273 {
		v410 = v273
		goto L52
	} else {
		goto L64
	}
L55:
	;
	if l2 == int32(-1) {
		v368 = v215
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v241 = int32(-2)
	if l2 < v241 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v244 = l2
	goto L59
L58:
	;
	v244 = v241
	goto L59
L59:
	;
	v247 = v215
	v253 = int32(0)
	goto L60
L60:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v247)+8))
	v268 = v253 + int32(1)
	if v244^v268 != int32(-1) {
		v247 = v266
		v253 = v268
		goto L60
	} else {
		goto L62
	}
L61:
	;
	if v266 != 0 {
		v368 = v266
		goto L53
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v410 = int32(0)
	goto L52
L64:
	;
	v279 = l2 - v111 + v230 + int32(1)
	v282 = v112
	v293 = v52
	v296 = int32(0)
	goto L66
L65:
	;
	if v322 != v279 {
		v410 = v273
		goto L52
	} else {
		goto L80
	}
L66:
	;
	v302 = v293 << (uint(int32(3)) % 32)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v282+v302)+12))
	if v304 == int32(0) {
		v342 = v282
		v356 = v296
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if v356 == v279 {
		v368 = v342
		goto L53
	} else {
		goto L78
	}
L69:
	;
	__phi308 = v282
	__phi314 = v304
	__phi322 = v296
	v308 = __phi308
	v314 = __phi314
	v322 = __phi322
	goto L70
L70:
	;
	if v293 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v342 = v314
	v356 = v338
	goto L68
L72:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v314+v302)+12))
	if v340 != 0 {
		__phi308 = v314
		__phi314 = v340
		__phi322 = v338
		v308 = __phi308
		v314 = __phi314
		v322 = __phi322
		goto L70
	} else {
		goto L77
	}
L73:
	;
	v336 = v322 + int32(1)
	if base.Ui32(v279) < base.Ui32(v336) {
		goto L65
	} else {
		goto L76
	}
L74:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v308+v302+int32(16))))
	v333 = v332 + v322
	if base.Ui32(v333) <= base.Ui32(v279) {
		v338 = v333
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v342 = v308
	v356 = v322
	goto L68
L76:
	;
	v338 = v336
	goto L72
L77:
	;
	goto L71
L78:
	;
	if v293 < int32(1) {
		v410 = v273
		goto L52
	} else {
		goto L79
	}
L79:
	;
	v282 = v342
	v293 = v293 + int32(-1)
	v296 = v356
	goto L66
L80:
	;
	v368 = v308
	goto L53
L81:
	;
	v390 = base.F64_gt(v387, v22)
	goto L83
L82:
	;
	v390 = base.F64_ge(v387, v22)
	goto L83
L83:
	;
	if v390 != int32(1) {
		v826 = v120
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v410 = v368
	goto L52
L85:
	;
	v798 = v410
	v800 = v230 + l2
	goto L31
L86:
	;
	v525 = int32(0)
	v526 = v521 + l2
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v527) <= base.Ui32(v526) {
		v826 = v525
		goto L1
	} else {
		goto L104
	}
L87:
	;
	v421 = v112
	v427 = v50 + int32(-2)
	v436 = v111
	goto L88
L88:
	;
	v441 = v427 << (uint(int32(3)) % 32)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v421+v441)+12))
	if v443 == int32(0) {
		v482 = v421
		v497 = v436
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v506 = v482
	v521 = v497
	goto L86
L90:
	;
	if int32(0) < v427 {
		v421 = v482
		v427 = v427 + int32(-1)
		v436 = v497
		goto L88
	} else {
		goto L103
	}
L91:
	;
	__phi453 = v443
	__phi462 = v436
	__phi464 = v421
	v453 = __phi453
	v462 = __phi462
	v464 = __phi464
	goto L92
L92:
	;
	v466 = *(*float64)(unsafe.Add(mBase, uint32(v453)))
	if v34 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v482 = v453
	v497 = v478
	goto L90
L94:
	;
	if v427 != 0 {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v469 = base.F64_gt(v466, v22)
	goto L97
L96:
	;
	v469 = base.F64_ge(v466, v22)
	goto L97
L97:
	;
	if v469 == int32(0) {
		goto L94
	} else {
		goto L98
	}
L98:
	;
	v482 = v464
	v497 = v462
	goto L90
L99:
	;
	v478 = v477 + v462
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v453+v441)+12))
	if v480 != 0 {
		__phi453 = v480
		__phi462 = v478
		__phi464 = v453
		v453 = __phi453
		v462 = __phi462
		v464 = __phi464
		goto L92
	} else {
		goto L102
	}
L100:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v464+v441+int32(16))))
	v477 = v476
	goto L99
L101:
	;
	v477 = int32(1)
	goto L99
L102:
	;
	goto L93
L103:
	;
	goto L89
L104:
	;
	if int32(9) < l2 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	if l3 == int32(0) {
		v826 = v776
		goto L1
	} else {
		goto L141
	}
L106:
	;
	v753 = *(*float64)(unsafe.Add(mBase, uint32(v734)))
	if v44 != 0 {
		goto L137
	} else {
		goto L138
	}
L107:
	;
	v639 = int32(0)
	if v50 <= v639 {
		v776 = v639
		goto L105
	} else {
		goto L120
	}
L108:
	;
	v532 = l2 + int32(1)
	v533 = int32(7)
	v534 = v532 & v533
	if base.Ui32(l2) < base.Ui32(v533) {
		v572 = v506
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if v534 == int32(0) {
		v619 = v572
		goto L114
	} else {
		goto L115
	}
L110:
	;
	v541 = v506
	v547 = int32(0)
	goto L111
L111:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v541)+12))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)+12))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)+12))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)+12))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)+12))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)+12))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)+12))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	v569 = v547 + int32(8)
	if v569 != v532&int32(-8) {
		v541 = v567
		v547 = v569
		goto L111
	} else {
		goto L113
	}
L112:
	;
	v572 = v567
	goto L109
L113:
	;
	goto L112
L114:
	;
	if v619 != 0 {
		v734 = v619
		goto L106
	} else {
		goto L119
	}
L115:
	;
	v595 = v572
	v601 = int32(0)
	goto L116
L116:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v595)+12))
	v616 = v601 + int32(1)
	if v616 != v534 {
		v595 = v614
		v601 = v616
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v619 = v614
	goto L114
L118:
	;
	goto L117
L119:
	;
	v776 = int32(0)
	goto L105
L120:
	;
	v645 = l2 - v111 + v521 + int32(1)
	v648 = v112
	v659 = v52
	v663 = int32(0)
	goto L122
L121:
	;
	if v689 != v645 {
		v776 = v639
		goto L105
	} else {
		goto L136
	}
L122:
	;
	v668 = v659 << (uint(int32(3)) % 32)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v648+v668)+12))
	if v670 == int32(0) {
		v708 = v648
		v723 = v663
		goto L124
	} else {
		goto L125
	}
L124:
	;
	if v723 == v645 {
		v734 = v708
		goto L106
	} else {
		goto L134
	}
L125:
	;
	__phi674 = v648
	__phi680 = v670
	__phi689 = v663
	v674 = __phi674
	v680 = __phi680
	v689 = __phi689
	goto L126
L126:
	;
	if v659 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v708 = v680
	v723 = v704
	goto L124
L128:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v680+v668)+12))
	if v706 != 0 {
		__phi674 = v680
		__phi680 = v706
		__phi689 = v704
		v674 = __phi674
		v680 = __phi680
		v689 = __phi689
		goto L126
	} else {
		goto L133
	}
L129:
	;
	v702 = v689 + int32(1)
	if base.Ui32(v645) < base.Ui32(v702) {
		goto L121
	} else {
		goto L132
	}
L130:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v674+v668+int32(16))))
	v699 = v698 + v689
	if base.Ui32(v699) <= base.Ui32(v645) {
		v704 = v699
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v708 = v674
	v723 = v689
	goto L124
L132:
	;
	v704 = v702
	goto L128
L133:
	;
	goto L127
L134:
	;
	if v659 < int32(1) {
		v776 = v639
		goto L105
	} else {
		goto L135
	}
L135:
	;
	v648 = v708
	v659 = v659 + int32(-1)
	v663 = v723
	goto L122
L136:
	;
	v734 = v674
	goto L106
L137:
	;
	v756 = base.F64_lt(v753, v23)
	goto L139
L138:
	;
	v756 = base.F64_le(v753, v23)
	goto L139
L139:
	;
	if v756 != int32(1) {
		v826 = v525
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v776 = v734
	goto L105
L141:
	;
	v798 = v776
	v800 = v526
	goto L31
}
func F_zslValueGteMin(m *base.Module, l0 float64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v7 != 0 {
		v8 = base.F64_gt(l0, v4)
	} else {
		v8 = base.F64_ge(l0, v4)
	}
	return v8
}
func F_zslValueLteMax(m *base.Module, l0 float64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v7 != 0 {
		v8 = base.F64_lt(l0, v4)
	} else {
		v8 = base.F64_le(l0, v4)
	}
	return v8
}
