package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_createObjectFromStreamID(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int64
	_ = v153
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int64
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v232 int64
	_ = v232
	var v234 int32
	_ = v234
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v259 int64
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v14 = int32(1)
	if base.Ui64(v9) < base.Ui64(int64(10)) {
		v71 = v14
		v72 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v146 = v6 + v145
	v147 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v146))) = uint8(v147)
	v149 = int32(1)
	v150 = v146 + v149
	v151 = int32(0)
	v153 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui64(v153) < base.Ui64(int64(10)) {
		v215 = v149
		v216 = v151
		goto L46
	} else {
		goto L47
	}
L2:
	;
	v75 = v71 + v72
	if base.Ui32(int32(21)) <= base.Ui32(v75) {
		goto L33
	} else {
		goto L34
	}
L3:
	;
	v22 = v2
	v23 = v9
	goto L4
L4:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v23) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v71 = v14
	v72 = v63
	goto L2
L6:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v23) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v71 = int32(2)
	v72 = v22
	goto L2
L8:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v23) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v71 = int32(3)
	v72 = v22
	goto L2
L10:
	;
	v63 = v22 + int32(12)
	v67 = base.I64_div_u_s(v23, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v23) {
		v22 = v63
		v23 = v67
		goto L4
	} else {
		goto L32
	}
L11:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v23) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v23) {
		goto L24
	} else {
		goto L25
	}
L13:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v23) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v23) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v23) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v23) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v71 = int32(4)
	v72 = v22
	goto L2
L18:
	;
	v44 = int32(6)
	goto L20
L19:
	;
	v44 = int32(5)
	goto L20
L20:
	;
	v71 = v44
	v72 = v22
	goto L2
L21:
	;
	v49 = int32(8)
	goto L23
L22:
	;
	v49 = int32(7)
	goto L23
L23:
	;
	v71 = v49
	v72 = v22
	goto L2
L24:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v23) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v23) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v56 = int32(10)
	goto L28
L27:
	;
	v56 = int32(9)
	goto L28
L28:
	;
	v71 = v56
	v72 = v22
	goto L2
L29:
	;
	v61 = int32(12)
	goto L31
L30:
	;
	v61 = int32(11)
	goto L31
L31:
	;
	v71 = v61
	v72 = v22
	goto L2
L32:
	;
	goto L5
L33:
	;
	goto L44
L34:
	;
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6+v75))) = uint8(v78)
	v81 = v75 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v9) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v117 = v6 + v114
	if base.Ui64(int64(9)) < base.Ui64(v115) {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	v88 = v9
	v90 = v81
	goto L38
L37:
	;
	v114 = v81
	v115 = v9
	goto L35
L38:
	;
	v94 = int64(100)
	v95 = base.I64_div_u_s(v88, v94)
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v88-v95*v94)<<(uint(int32(1))%32))+uint32(_consts[185]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v6+int32(-1)+v90))) = uint16(v104)
	v107 = v90 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v88) {
		v88 = v95
		v90 = v107
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v114 = v107
	v115 = v95
	goto L35
L40:
	;
	goto L39
L41:
	;
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v115)<<(uint(int32(1))%32))+uint32(_consts[185]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v117+int32(-1)))) = uint16(v131)
	v145 = v75
	goto L1
L42:
	;
	v122 = base.I32_wrap_i64(v115) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v122)
	v145 = v75
	goto L1
L43:
	;
	v145 = int32(0)
	goto L1
L44:
	;
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v135)
	goto L43
L45:
	;
	v292 = F_sdsnewlen(m, v6, v150+v289-v6)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L89
	} else {
		goto L90
	}
L46:
	;
	v219 = v215 + v216
	if base.Ui32(int32(21)) <= base.Ui32(v219) {
		goto L77
	} else {
		goto L78
	}
L47:
	;
	v166 = v151
	v167 = v153
	goto L48
L48:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v167) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v215 = v149
	v216 = v207
	goto L46
L50:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v167) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v215 = int32(2)
	v216 = v166
	goto L46
L52:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v167) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v215 = int32(3)
	v216 = v166
	goto L46
L54:
	;
	v207 = v166 + int32(12)
	v211 = base.I64_div_u_s(v167, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v167) {
		v166 = v207
		v167 = v211
		goto L48
	} else {
		goto L76
	}
L55:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v167) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v167) {
		goto L68
	} else {
		goto L69
	}
L57:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v167) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v167) {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v167) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v167) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v215 = int32(4)
	v216 = v166
	goto L46
L62:
	;
	v188 = int32(6)
	goto L64
L63:
	;
	v188 = int32(5)
	goto L64
L64:
	;
	v215 = v188
	v216 = v166
	goto L46
L65:
	;
	v193 = int32(8)
	goto L67
L66:
	;
	v193 = int32(7)
	goto L67
L67:
	;
	v215 = v193
	v216 = v166
	goto L46
L68:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v167) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v167) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v200 = int32(10)
	goto L72
L71:
	;
	v200 = int32(9)
	goto L72
L72:
	;
	v215 = v200
	v216 = v166
	goto L46
L73:
	;
	v205 = int32(12)
	goto L75
L74:
	;
	v205 = int32(11)
	goto L75
L75:
	;
	v215 = v205
	v216 = v166
	goto L46
L76:
	;
	goto L49
L77:
	;
	goto L88
L78:
	;
	v222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v150+v219))) = uint8(v222)
	v225 = v219 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v153) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v261 = v150 + v258
	if base.Ui64(int64(9)) < base.Ui64(v259) {
		goto L85
	} else {
		goto L86
	}
L80:
	;
	v232 = v153
	v234 = v225
	goto L82
L81:
	;
	v258 = v225
	v259 = v153
	goto L79
L82:
	;
	v238 = int64(100)
	v239 = base.I64_div_u_s(v232, v238)
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v232-v239*v238)<<(uint(int32(1))%32))+uint32(_consts[185]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v146+int32(0)+v234))) = uint16(v248)
	v251 = v234 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v232) {
		v232 = v239
		v234 = v251
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v258 = v251
	v259 = v239
	goto L79
L84:
	;
	goto L83
L85:
	;
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v259)<<(uint(int32(1))%32))+uint32(_consts[185]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v261+int32(-1)))) = uint16(v275)
	v289 = v219
	goto L45
L86:
	;
	v266 = base.I32_wrap_i64(v259) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v266)
	v289 = v219
	goto L45
L87:
	;
	v289 = int32(0)
	goto L45
L88:
	;
	v279 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v279)
	goto L87
L89:
	;
	return int32(0)
L90:
	;
	v296 = F_createObject(m, v151, v292)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	m.G0 = v6 + int32(48)
	return v296
}
func F_createSetObject(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_hashtableCreate(m, int32(_a846))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
		v15 = int32(12)
		v18 = F_zmalloc_usable(m, v15, v6+v15)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v9
			*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(34359738370)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23&int32(-241) | int32(32)
			m.G0 = v6 + int32(16)
			return v18
		}
	}
}
func F_objectCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v57 int32
	_ = v57
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
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
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int64
	_ = v127
	var v132 int64
	_ = v132
	var v137 int64
	_ = v137
	var v142 int64
	_ = v142
	var v145 int64
	_ = v145
	var v148 int64
	_ = v148
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v739 int64
	_ = v739
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	v7 = m.G0
	v9 = v7 - int32(64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v11 != int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(64)
	return
L2:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v158&int32(4) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v16&int32(4) == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v80 = int32(_a757)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v83 != 0 {
		goto L23
	} else {
		goto L24
	}
L5:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v76 = v75
	goto L4
L6:
	;
	if v16&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v25 = int32(16)
	goto L9
L8:
	;
	v25 = int32(8)
	goto L9
L9:
	;
	v26 = v15 + v25
	if v16&int32(2) == int32(0) {
		v57 = v26
		goto L10
	} else {
		goto L11
	}
L10:
	;
	goto L20
L11:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v32 = v26 + v31
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	switch v36 & int32(7) {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	default:
		v53 = int32(0)
		goto L12
	}
L12:
	;
	v57 = v32 + int32(1) + v53 + int32(1)
	goto L10
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(-16))))
	v53 = v52
	goto L12
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(-8))))
	v53 = v49
	goto L12
L15:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+int32(-4)))))
	v53 = v46
	goto L12
L16:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-2)))))
	v53 = v43
	goto L12
L17:
	;
	v53 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
	goto L12
L18:
	;
	v76 = v57 + v72
	goto L4
L19:
	;
	goto L18
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L19
L21:
	;
	if v115-v117 != 0 {
		goto L2
	} else {
		goto L33
	}
L22:
	;
	v115 = F_tolower(m, v111)
	mBase = m.M
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v117 = F_tolower(m, v116)
	mBase = m.M
	goto L21
L23:
	;
	v85 = v76
	v86 = v80
	v87 = v83
	goto L26
L24:
	;
	v111 = int32(0)
	v112 = v80
	goto L22
L25:
	;
	v111 = v108 & int32(255)
	v112 = v107
	goto L22
L26:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v89 == int32(0) {
		v107 = v86
		v108 = v87
		goto L25
	} else {
		goto L28
	}
L27:
	;
	v107 = v101
	v108 = int32(0)
	goto L25
L28:
	;
	v93 = v87 & int32(255)
	if v93 == v89 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v100 = int32(1)
	v101 = v86 + v100
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v102 != 0 {
		v85 = v85 + v100
		v86 = v101
		v87 = v102
		goto L26
	} else {
		goto L32
	}
L30:
	;
	v95 = F_tolower(m, v93)
	mBase = m.M
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v97 = F_tolower(m, v96)
	mBase = m.M
	if v95 == v97 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v107 = v86
	v108 = v99
	goto L25
L32:
	;
	goto L27
L33:
	;
	v121 = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, _consts[453]))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v122
	v127 = *(*int64)(unsafe.Add(mBase, _consts[454]))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(40)))) = v127
	v132 = *(*int64)(unsafe.Add(mBase, _consts[455]))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(32)))) = v132
	v137 = *(*int64)(unsafe.Add(mBase, _consts[456]))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(24)))) = v137
	v142 = *(*int64)(unsafe.Add(mBase, _consts[457]))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v142
	v145 = *(*int64)(unsafe.Add(mBase, _consts[458]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v145
	v148 = *(*int64)(unsafe.Add(mBase, _consts[459]))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v148
	F_addReplyHelp(m, l0, v9)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return
L35:
	;
	goto L1
L36:
	;
	v222 = int32(_a857)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if v225 != 0 {
		goto L57
	} else {
		goto L58
	}
L37:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v218 = v217
	goto L36
L38:
	;
	if v158&int32(1) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v167 = int32(16)
	goto L41
L40:
	;
	v167 = int32(8)
	goto L41
L41:
	;
	v168 = v157 + v167
	if v158&int32(2) == int32(0) {
		v199 = v168
		goto L42
	} else {
		goto L43
	}
L42:
	;
	goto L52
L43:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	v174 = v168 + v173
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	switch v178 & int32(7) {
	case 0:
		goto L49
	case 1:
		goto L48
	case 2:
		goto L47
	case 3:
		goto L46
	case 4:
		goto L45
	default:
		v195 = int32(0)
		goto L44
	}
L44:
	;
	v199 = v174 + int32(1) + v195 + int32(1)
	goto L42
L45:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v174+int32(-16))))
	v195 = v194
	goto L44
L46:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v174+int32(-8))))
	v195 = v191
	goto L44
L47:
	;
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v174+int32(-4)))))
	v195 = v188
	goto L44
L48:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+int32(-2)))))
	v195 = v185
	goto L44
L49:
	;
	v195 = int32(base.Ui32(v178) >> (uint(int32(3)) % 32))
	goto L44
L50:
	;
	v218 = v199 + v214
	goto L36
L51:
	;
	goto L50
L52:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L51
L53:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	if v289&int32(4) == int32(0) {
		goto L75
	} else {
		goto L76
	}
L54:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v265 != int32(3) {
		v287 = v264
		goto L53
	} else {
		goto L68
	}
L55:
	;
	if v257-v259 == int32(0) {
		goto L54
	} else {
		goto L67
	}
L56:
	;
	v257 = F_tolower(m, v253)
	mBase = m.M
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	v259 = F_tolower(m, v258)
	mBase = m.M
	goto L55
L57:
	;
	v227 = v218
	v228 = v222
	v229 = v225
	goto L60
L58:
	;
	v253 = int32(0)
	v254 = v222
	goto L56
L59:
	;
	v253 = v250 & int32(255)
	v254 = v249
	goto L56
L60:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	if v231 == int32(0) {
		v249 = v228
		v250 = v229
		goto L59
	} else {
		goto L62
	}
L61:
	;
	v249 = v243
	v250 = int32(0)
	goto L59
L62:
	;
	v235 = v229 & int32(255)
	if v235 == v231 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v242 = int32(1)
	v243 = v228 + v242
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+1)))
	if v244 != 0 {
		v227 = v227 + v242
		v228 = v243
		v229 = v244
		goto L60
	} else {
		goto L66
	}
L64:
	;
	v237 = F_tolower(m, v235)
	mBase = m.M
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	v239 = F_tolower(m, v238)
	mBase = m.M
	if v237 == v239 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	v249 = v228
	v250 = v241
	goto L59
L66:
	;
	goto L61
L67:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v287 = v263
	goto L53
L68:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269<<(uint(int32(2))%32))+uint32(_consts[460])))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	v277 = F_lookupKeyReadWithFlags(m, v274, v275, int32(3))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L34
	} else {
		goto L70
	}
L69:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(int32(base.Ui32(v281)>>(uint(int32(3))%32))))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L34
	} else {
		goto L73
	}
L70:
	;
	if v277 != 0 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	F_addReplyOrErrorObject(m, l0, v273)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L34
	} else {
		goto L72
	}
L72:
	;
	goto L1
L73:
	;
	goto L1
L74:
	;
	v353 = int32(_a650)
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	if v356 != 0 {
		goto L95
	} else {
		goto L96
	}
L75:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v288)+8))
	v349 = v348
	goto L74
L76:
	;
	if v289&int32(1) != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v298 = int32(16)
	goto L79
L78:
	;
	v298 = int32(8)
	goto L79
L79:
	;
	v299 = v288 + v298
	if v289&int32(2) == int32(0) {
		v330 = v299
		goto L80
	} else {
		goto L81
	}
L80:
	;
	goto L90
L81:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	v305 = v299 + v304
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	switch v309 & int32(7) {
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
		v326 = int32(0)
		goto L82
	}
L82:
	;
	v330 = v305 + int32(1) + v326 + int32(1)
	goto L80
L83:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v305+int32(-16))))
	v326 = v325
	goto L82
L84:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v305+int32(-8))))
	v326 = v322
	goto L82
L85:
	;
	v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v305+int32(-4)))))
	v326 = v319
	goto L82
L86:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305+int32(-2)))))
	v326 = v316
	goto L82
L87:
	;
	v326 = int32(base.Ui32(v309) >> (uint(int32(3)) % 32))
	goto L82
L88:
	;
	v349 = v330 + v345
	goto L74
L89:
	;
	goto L88
L90:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L89
L91:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+4))
	if v430&int32(4) == int32(0) {
		goto L115
	} else {
		goto L116
	}
L92:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v396 != int32(3) {
		v428 = v395
		goto L91
	} else {
		goto L106
	}
L93:
	;
	if v388-v390 == int32(0) {
		goto L92
	} else {
		goto L105
	}
L94:
	;
	v388 = F_tolower(m, v384)
	mBase = m.M
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	v390 = F_tolower(m, v389)
	mBase = m.M
	goto L93
L95:
	;
	v358 = v349
	v359 = v353
	v360 = v356
	goto L98
L96:
	;
	v384 = int32(0)
	v385 = v353
	goto L94
L97:
	;
	v384 = v381 & int32(255)
	v385 = v380
	goto L94
L98:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359))))
	if v362 == int32(0) {
		v380 = v359
		v381 = v360
		goto L97
	} else {
		goto L100
	}
L99:
	;
	v380 = v374
	v381 = int32(0)
	goto L97
L100:
	;
	v366 = v360 & int32(255)
	if v366 == v362 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v373 = int32(1)
	v374 = v359 + v373
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+1)))
	if v375 != 0 {
		v358 = v358 + v373
		v359 = v374
		v360 = v375
		goto L98
	} else {
		goto L104
	}
L102:
	;
	v368 = F_tolower(m, v366)
	mBase = m.M
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359))))
	v370 = F_tolower(m, v369)
	mBase = m.M
	if v368 == v370 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	v380 = v359
	v381 = v372
	goto L97
L104:
	;
	goto L99
L105:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v428 = v394
	goto L91
L106:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v400<<(uint(int32(2))%32))+uint32(_consts[460])))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v395)+8))
	v408 = F_lookupKeyReadWithFlags(m, v405, v406, int32(3))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L34
	} else {
		goto L108
	}
L107:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	v417 = int32(base.Ui32(v413)>>(uint(int32(4))%32)) & int32(15)
	if base.Ui32(int32(11)) < base.Ui32(v417) {
		v425 = int32(_a242)
		goto L111
	} else {
		goto L112
	}
L108:
	;
	if v408 != 0 {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	F_addReplyOrErrorObject(m, l0, v404)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L34
	} else {
		goto L110
	}
L110:
	;
	goto L1
L111:
	;
	F_addReplyBulkCString(m, l0, v425)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L34
	} else {
		goto L113
	}
L112:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v417<<(uint(int32(2))%32))+uint32(_consts[461])))
	v425 = v424
	goto L111
L113:
	;
	goto L1
L114:
	;
	v494 = int32(_a858)
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490))))
	if v497 != 0 {
		goto L135
	} else {
		goto L136
	}
L115:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v429)+8))
	v490 = v489
	goto L114
L116:
	;
	if v430&int32(1) != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v439 = int32(16)
	goto L119
L118:
	;
	v439 = int32(8)
	goto L119
L119:
	;
	v440 = v429 + v439
	if v430&int32(2) == int32(0) {
		v471 = v440
		goto L120
	} else {
		goto L121
	}
L120:
	;
	goto L130
L121:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
	v446 = v440 + v445
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	switch v450 & int32(7) {
	case 0:
		goto L127
	case 1:
		goto L126
	case 2:
		goto L125
	case 3:
		goto L124
	case 4:
		goto L123
	default:
		v467 = int32(0)
		goto L122
	}
L122:
	;
	v471 = v446 + int32(1) + v467 + int32(1)
	goto L120
L123:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v446+int32(-16))))
	v467 = v466
	goto L122
L124:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v446+int32(-8))))
	v467 = v463
	goto L122
L125:
	;
	v460 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v446+int32(-4)))))
	v467 = v460
	goto L122
L126:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446+int32(-2)))))
	v467 = v457
	goto L122
L127:
	;
	v467 = int32(base.Ui32(v450) >> (uint(int32(3)) % 32))
	goto L122
L128:
	;
	v490 = v471 + v486
	goto L114
L129:
	;
	goto L128
L130:
	;
	v486 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L129
L131:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	if v575&int32(4) == int32(0) {
		goto L157
	} else {
		goto L158
	}
L132:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v537 != int32(3) {
		v573 = v536
		goto L131
	} else {
		goto L146
	}
L133:
	;
	if v529-v531 == int32(0) {
		goto L132
	} else {
		goto L145
	}
L134:
	;
	v529 = F_tolower(m, v525)
	mBase = m.M
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	v531 = F_tolower(m, v530)
	mBase = m.M
	goto L133
L135:
	;
	v499 = v490
	v500 = v494
	v501 = v497
	goto L138
L136:
	;
	v525 = int32(0)
	v526 = v494
	goto L134
L137:
	;
	v525 = v522 & int32(255)
	v526 = v521
	goto L134
L138:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	if v503 == int32(0) {
		v521 = v500
		v522 = v501
		goto L137
	} else {
		goto L140
	}
L139:
	;
	v521 = v515
	v522 = int32(0)
	goto L137
L140:
	;
	v507 = v501 & int32(255)
	if v507 == v503 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v514 = int32(1)
	v515 = v500 + v514
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+1)))
	if v516 != 0 {
		v499 = v499 + v514
		v500 = v515
		v501 = v516
		goto L138
	} else {
		goto L144
	}
L142:
	;
	v509 = F_tolower(m, v507)
	mBase = m.M
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	v511 = F_tolower(m, v510)
	mBase = m.M
	if v509 == v511 {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499))))
	v521 = v500
	v522 = v513
	goto L137
L144:
	;
	goto L139
L145:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v573 = v535
	goto L131
L146:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v541<<(uint(int32(2))%32))+uint32(_consts[460])))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v536)+8))
	v549 = F_lookupKeyReadWithFlags(m, v546, v547, int32(3))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L34
	} else {
		goto L148
	}
L147:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, _consts[167])))
	if v554&int32(2) == int32(0) {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	if v549 != 0 {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	F_addReplyOrErrorObject(m, l0, v545)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L34
	} else {
		goto L150
	}
L150:
	;
	goto L1
L151:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	v566 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	goto L154
L152:
	;
	F_addReplyError(m, l0, int32(_a859))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L34
	} else {
		goto L153
	}
L153:
	;
	goto L1
L154:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u((v566-int32(base.Ui32(v562)>>(uint(int32(8))%32)))&int32(16777215)))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L34
	} else {
		goto L155
	}
L155:
	;
	goto L1
L156:
	;
	v639 = int32(_a860)
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635))))
	if v642 != 0 {
		goto L176
	} else {
		goto L177
	}
L157:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v574)+8))
	v635 = v634
	goto L156
L158:
	;
	if v575&int32(1) != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v584 = int32(16)
	goto L161
L160:
	;
	v584 = int32(8)
	goto L161
L161:
	;
	v585 = v574 + v584
	if v575&int32(2) == int32(0) {
		v616 = v585
		goto L162
	} else {
		goto L163
	}
L162:
	;
	goto L172
L163:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585))))
	v591 = v585 + v590
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591))))
	switch v595 & int32(7) {
	case 0:
		goto L169
	case 1:
		goto L168
	case 2:
		goto L167
	case 3:
		goto L166
	case 4:
		goto L165
	default:
		v612 = int32(0)
		goto L164
	}
L164:
	;
	v616 = v591 + int32(1) + v612 + int32(1)
	goto L162
L165:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v591+int32(-16))))
	v612 = v611
	goto L164
L166:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v591+int32(-8))))
	v612 = v608
	goto L164
L167:
	;
	v605 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v591+int32(-4)))))
	v612 = v605
	goto L164
L168:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591+int32(-2)))))
	v612 = v602
	goto L164
L169:
	;
	v612 = int32(base.Ui32(v595) >> (uint(int32(3)) % 32))
	goto L164
L170:
	;
	v635 = v616 + v631
	goto L156
L171:
	;
	goto L170
L172:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L171
L173:
	;
	F_addReplySubcommandSyntaxError(m, l0)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L34
	} else {
		goto L202
	}
L174:
	;
	if v674-v676 != 0 {
		goto L173
	} else {
		goto L186
	}
L175:
	;
	v674 = F_tolower(m, v670)
	mBase = m.M
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671))))
	v676 = F_tolower(m, v675)
	mBase = m.M
	goto L174
L176:
	;
	v644 = v635
	v645 = v639
	v646 = v642
	goto L179
L177:
	;
	v670 = int32(0)
	v671 = v639
	goto L175
L178:
	;
	v670 = v667 & int32(255)
	v671 = v666
	goto L175
L179:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645))))
	if v648 == int32(0) {
		v666 = v645
		v667 = v646
		goto L178
	} else {
		goto L181
	}
L180:
	;
	v666 = v660
	v667 = int32(0)
	goto L178
L181:
	;
	v652 = v646 & int32(255)
	if v652 == v648 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v659 = int32(1)
	v660 = v645 + v659
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644)+1)))
	if v661 != 0 {
		v644 = v644 + v659
		v645 = v660
		v646 = v661
		goto L179
	} else {
		goto L185
	}
L183:
	;
	v654 = F_tolower(m, v652)
	mBase = m.M
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645))))
	v656 = F_tolower(m, v655)
	mBase = m.M
	if v654 == v656 {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644))))
	v666 = v645
	v667 = v658
	goto L178
L185:
	;
	goto L180
L186:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v678 != int32(3) {
		goto L173
	} else {
		goto L187
	}
L187:
	;
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v682<<(uint(int32(2))%32))+uint32(_consts[460])))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v688)+8))
	v691 = F_lookupKeyReadWithFlags(m, v687, v689, int32(3))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L34
	} else {
		goto L189
	}
L188:
	;
	v696 = int32(*(*uint8)(unsafe.Add(mBase, _consts[167])))
	if v696&int32(2) != 0 {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	if v691 != 0 {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	F_addReplyOrErrorObject(m, l0, v686)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L34
	} else {
		goto L191
	}
L191:
	;
	goto L1
L192:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v691)))
	v704 = int32(base.Ui32(v702) >> (uint(int32(8)) % 32))
	v705 = int32(0)
	v709 = int32(*(*uint16)(unsafe.Add(mBase, _consts[335])))
	v712 = *(*int32)(unsafe.Add(mBase, _consts[336]))
	if v712 == v705 {
		v723 = v705
		goto L196
	} else {
		goto L197
	}
L193:
	;
	F_addReplyError(m, l0, int32(_a861))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L34
	} else {
		goto L194
	}
L194:
	;
	goto L1
L195:
	;
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691))))
	*(*int32)(unsafe.Add(mBase, uint32(v691))) = v734 | (v729|v709<<(uint(int32(8))%32))<<(uint(int32(8))%32)
	v739 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	F_addReplyLongLong(m, l0, v739)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L34
	} else {
		goto L201
	}
L196:
	;
	v726 = v704 & int32(255)
	v727 = v726 - v723
	if base.Ui32(v726) < base.Ui32(v727) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v718 = int32(65535)
	v720 = base.I32_div_s((v709-int32(base.Ui32(v704)>>(uint(int32(8))%32)))&v718, v712)
	v723 = v720 & v718
	goto L196
L198:
	;
	v729 = int32(0)
	goto L200
L199:
	;
	v729 = v727
	goto L200
L200:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v729)
	goto L195
L201:
	;
	goto L1
L202:
	;
	goto L1
}
func F_objectComputeSize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
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
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v262 float64
	_ = v262
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 float64
	_ = v532
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v690 int32
	_ = v690
	var v698 int32
	_ = v698
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v952 float64
	_ = v952
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v974 int32
	_ = v974
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1317 float64
	_ = v1317
	var v1325 int32
	_ = v1325
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1354 int32
	_ = v1354
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1376 int32
	_ = v1376
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1409 int64
	_ = v1409
	var v1421 int32
	_ = v1421
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int64
	_ = v1471
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int64
	_ = v1479
	var v1488 int32
	_ = v1488
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1506 int32
	_ = v1506
	var v1511 int32
	_ = v1511
	var v1515 int32
	_ = v1515
	var v1522 int32
	_ = v1522
	var v1526 int64
	_ = v1526
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1561 int32
	_ = v1561
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1585 int64
	_ = v1585
	var v1597 int32
	_ = v1597
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1608 int32
	_ = v1608
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int64
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1642 int64
	_ = v1642
	var v1654 int32
	_ = v1654
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1666 int32
	_ = v1666
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1692 int32
	_ = v1692
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1734 int32
	_ = v1734
	var v1735 int64
	_ = v1735
	var v1739 float64
	_ = v1739
	var v1747 int32
	_ = v1747
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1763 int32
	_ = v1763
	var v1767 int32
	_ = v1767
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1788 int32
	_ = v1788
	var v1789 int64
	_ = v1789
	var v1793 float64
	_ = v1793
	var v1801 int32
	_ = v1801
	var v1823 int32
	_ = v1823
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1857 int32
	_ = v1857
	v16 = m.G0
	v18 = v16 - int32(608)
	m.G0 = v18
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-8))))
	v24 = v22 & int32(2147483647)
	v26 = v24 + int32(8)
	goto L1
L1:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v27 & int32(15) {
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
	case 5:
		goto L10
	case 6:
		goto L11
	default:
		goto L9
	}
L2:
	;
	m.G0 = v18 + int32(608)
	return v1857
L3:
	;
	if v1112 == int32(0) {
		goto L445
	} else {
		goto L446
	}
L4:
	;
	v1857 = int32(0)
	goto L2
L5:
	;
	F_raxStop(m, v18+int32(304))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L49
	} else {
		goto L399
	}
L6:
	;
	v1534 = int32(0)
	v1536 = F_raxSeek(m, v18+int32(304), int32(_a828), v1534, v1534)
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L49
	} else {
		goto L396
	}
L7:
	;
	F_hashtableCleanupIterator(m, v18+int32(304))
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L49
	} else {
		goto L395
	}
L8:
	;
	F_hashtableCleanupIterator(m, v18+int32(304))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L49
	} else {
		goto L394
	}
L9:
	;
	F__serverPanic_1(m, int32(_a838), int32(1354), int32(_a123), int32(0))
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L49
	} else {
		goto L393
	}
L10:
	;
	v1498 = F_moduleGetMemUsage(m, l0, l1, l2, l3)
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L49
	} else {
		goto L392
	}
L11:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1334&int32(4) == int32(0) {
		goto L360
	} else {
		goto L361
	}
L12:
	;
	switch int32(base.Ui32(v27)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		goto L255
	default:
		goto L254
	case 9:
		goto L256
	}
L13:
	;
	switch int32(base.Ui32(v27)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L186
	default:
		goto L185
	case 4:
		goto L187
	}
L14:
	;
	switch int32(base.Ui32(v27)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		goto L100
	default:
		goto L97
	case 4:
		goto L99
	case 9:
		goto L98
	}
L15:
	;
	switch int32(base.Ui32(v27)>>(uint(int32(4))%32))&int32(15) + int32(-9) {
	case 0:
		goto L53
	default:
		goto L51
	case 2:
		goto L52
	}
L16:
	;
	switch int32(base.Ui32(v27)>>(uint(int32(4))%32)) & int32(15) {
	case 0:
		goto L18
	case 1, 8:
		v1857 = v26
		goto L2
	default:
		goto L17
	}
L17:
	;
	F__serverPanic_1(m, int32(_a838), int32(1202), int32(_a848), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L49
	} else {
		goto L50
	}
L18:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v34&int32(4) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v104 = v95 + int32(-1)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v107 = v105 & int32(7)
	switch v107 {
	case 0:
		goto L42
	case 1:
		v113 = int32(4)
		goto L37
	case 2:
		goto L41
	case 3:
		goto L40
	case 4:
		goto L39
	default:
		goto L38
	}
L20:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v95 = v93
	goto L19
L21:
	;
	if v34&int32(1) != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v43 = int32(16)
	goto L24
L23:
	;
	v43 = int32(8)
	goto L24
L24:
	;
	v44 = l1 + v43
	if v34&int32(2) == int32(0) {
		v76 = v44
		goto L25
	} else {
		goto L26
	}
L25:
	;
	goto L35
L26:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v50 = v44 + v49
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	switch v54 & int32(7) {
	case 0:
		goto L32
	case 1:
		goto L31
	case 2:
		goto L30
	case 3:
		goto L29
	case 4:
		goto L28
	default:
		v71 = int32(0)
		goto L27
	}
L27:
	;
	v76 = v50 + int32(1) + v71 + int32(1)
	goto L25
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v50+int32(-16))))
	v71 = v70
	goto L27
L29:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v50+int32(-8))))
	v71 = v67
	goto L27
L30:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50+int32(-4)))))
	v71 = v64
	goto L27
L31:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+int32(-2)))))
	v71 = v61
	goto L27
L32:
	;
	v71 = int32(base.Ui32(v54) >> (uint(int32(3)) % 32))
	goto L27
L33:
	;
	v95 = v76 + v90
	goto L19
L34:
	;
	goto L33
L35:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L34
L36:
	;
	v1857 = v137 + v26
	goto L2
L37:
	;
	switch v107 {
	case 0:
		goto L48
	case 1:
		goto L47
	case 2:
		goto L46
	case 3:
		goto L45
	case 4:
		goto L44
	default:
		v133 = int32(0)
		goto L43
	}
L38:
	;
	v113 = int32(1)
	goto L37
L39:
	;
	v113 = int32(18)
	goto L37
L40:
	;
	v113 = int32(10)
	goto L37
L41:
	;
	v113 = int32(6)
	goto L37
L42:
	;
	v108 = F_zmalloc_usable_size(m, v104)
	mBase = m.M
	v137 = v108
	goto L36
L43:
	;
	v137 = v113 + v133
	goto L36
L44:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v95+int32(-9))))
	v133 = v132
	goto L43
L45:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v95+int32(-5))))
	v137 = v113 + v128
	goto L36
L46:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95+int32(-3)))))
	v137 = v113 + v124
	goto L36
L47:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+int32(-2)))))
	v137 = v113 + v120
	goto L36
L48:
	;
	v137 = v113 + int32(base.Ui32(v105)>>(uint(int32(3))%32))
	goto L36
L49:
	;
	return int32(0)
L50:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F__serverPanic_1(m, int32(_a838), int32(1217), int32(_a852), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L49
	} else {
		goto L96
	}
L52:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v271&int32(4) == int32(0) {
		goto L79
	} else {
		goto L80
	}
L53:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v154&int32(4) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v219 = int32(0)
	v222 = v219
	v224 = v219
	v227 = v218
	goto L72
L55:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v214 = v213
	goto L54
L56:
	;
	if v154&int32(1) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v163 = int32(16)
	goto L59
L58:
	;
	v163 = int32(8)
	goto L59
L59:
	;
	v164 = l1 + v163
	if v154&int32(2) == int32(0) {
		v196 = v164
		goto L60
	} else {
		goto L61
	}
L60:
	;
	goto L70
L61:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	v170 = v164 + v169
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	switch v174 & int32(7) {
	case 0:
		goto L67
	case 1:
		goto L66
	case 2:
		goto L65
	case 3:
		goto L64
	case 4:
		goto L63
	default:
		v191 = int32(0)
		goto L62
	}
L62:
	;
	v196 = v170 + int32(1) + v191 + int32(1)
	goto L60
L63:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v170+int32(-16))))
	v191 = v190
	goto L62
L64:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v170+int32(-8))))
	v191 = v187
	goto L62
L65:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170+int32(-4)))))
	v191 = v184
	goto L62
L66:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+int32(-2)))))
	v191 = v181
	goto L62
L67:
	;
	v191 = int32(base.Ui32(v174) >> (uint(int32(3)) % 32))
	goto L62
L68:
	;
	v214 = v196 + v210
	goto L54
L69:
	;
	goto L68
L70:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L69
L71:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	v262 = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v248), base.F64_convert_i32_u(v237)), base.F64_convert_i32_u(v256)), base.F64_convert_i32_u(v24+int32(28)))
	if base.F64_lt(v262, float64(4.294967296e+09))&base.F64_ge(v262, float64(0)) == int32(0) {
		goto L4
	} else {
		goto L77
	}
L72:
	;
	v237 = v222 + int32(1)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v238+int32(-8))))
	goto L74
L73:
	;
	goto L71
L74:
	;
	v248 = v224 + (v241&int32(2147483647) + int32(8)) + int32(20)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v249 == int32(0) {
		goto L71
	} else {
		goto L75
	}
L75:
	;
	if base.Ui32(v237) < base.Ui32(l2) {
		v222 = v237
		v224 = v248
		v227 = v249
		goto L72
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	v270 = base.I32_trunc_f64_u(v262)
	v1857 = v270
	goto L2
L78:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v332+int32(-8))))
	goto L95
L79:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v332 = v330
	goto L78
L80:
	;
	if v271&int32(1) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v280 = int32(16)
	goto L83
L82:
	;
	v280 = int32(8)
	goto L83
L83:
	;
	v281 = l1 + v280
	if v271&int32(2) == int32(0) {
		v313 = v281
		goto L84
	} else {
		goto L85
	}
L84:
	;
	goto L94
L85:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	v287 = v281 + v286
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	switch v291 & int32(7) {
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
		v308 = int32(0)
		goto L86
	}
L86:
	;
	v313 = v287 + int32(1) + v308 + int32(1)
	goto L84
L87:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v287+int32(-16))))
	v308 = v307
	goto L86
L88:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v287+int32(-8))))
	v308 = v304
	goto L86
L89:
	;
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287+int32(-4)))))
	v308 = v301
	goto L86
L90:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287+int32(-2)))))
	v308 = v298
	goto L86
L91:
	;
	v308 = int32(base.Ui32(v291) >> (uint(int32(3)) % 32))
	goto L86
L92:
	;
	v332 = v313 + v327
	goto L78
L93:
	;
	goto L92
L94:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L93
L95:
	;
	v1857 = v337&int32(2147483647) + int32(8) + v26
	goto L2
L96:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F__serverPanic_1(m, int32(_a838), int32(1239), int32(_a853), int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L49
	} else {
		goto L184
	}
L98:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v613&int32(4) == int32(0) {
		goto L167
	} else {
		goto L168
	}
L99:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v541&int32(4) == int32(0) {
		goto L149
	} else {
		goto L150
	}
L100:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v356&int32(4) == int32(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v420 = F_hashtableMemUsage(m, v418)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L49
	} else {
		goto L118
	}
L102:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v418 = v415
	goto L101
L103:
	;
	if v356&int32(1) != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v365 = int32(16)
	goto L106
L105:
	;
	v365 = int32(8)
	goto L106
L106:
	;
	v366 = l1 + v365
	if v356&int32(2) == int32(0) {
		v398 = v366
		goto L107
	} else {
		goto L108
	}
L107:
	;
	goto L117
L108:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	v372 = v366 + v371
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	switch v376 & int32(7) {
	case 0:
		goto L114
	case 1:
		goto L113
	case 2:
		goto L112
	case 3:
		goto L111
	case 4:
		goto L110
	default:
		v393 = int32(0)
		goto L109
	}
L109:
	;
	v398 = v372 + int32(1) + v393 + int32(1)
	goto L107
L110:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v372+int32(-16))))
	v393 = v392
	goto L109
L111:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v372+int32(-8))))
	v393 = v389
	goto L109
L112:
	;
	v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v372+int32(-4)))))
	v393 = v386
	goto L109
L113:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372+int32(-2)))))
	v393 = v383
	goto L109
L114:
	;
	v393 = int32(base.Ui32(v376) >> (uint(int32(3)) % 32))
	goto L109
L115:
	;
	v418 = v398 + v412
	goto L101
L116:
	;
	goto L115
L117:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L116
L118:
	;
	v422 = int32(0)
	v424 = v18 + int32(304)
	*(*uint8)(unsafe.Add(mBase, uint32(v424)+14)) = uint8(v422)
	*(*int32)(unsafe.Add(mBase, uint32(v424))) = v418
	*(*int32)(unsafe.Add(mBase, uint32(v424)+24)) = v422
	*(*uint8)(unsafe.Add(mBase, uint32(v424)+15)) = uint8(v422)
	*(*int32)(unsafe.Add(mBase, uint32(v424)+8)) = int32(-1)
	if v418 == v422 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v443 = v420 + v26
	v446 = F_hashtableNext(m, v18+int32(304), v18)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L49
	} else {
		goto L123
	}
L120:
	;
	goto L119
L121:
	;
	goto L120
L123:
	;
	if v446 == int32(0) {
		goto L8
	} else {
		goto L124
	}
L124:
	;
	if l2 == int32(0) {
		goto L8
	} else {
		goto L125
	}
L125:
	;
	v454 = int32(0)
	v459 = v422
	goto L127
L126:
	;
	F_hashtableCleanupIterator(m, v18+int32(304))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L49
	} else {
		goto L145
	}
L127:
	;
	v469 = v454 + int32(1)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v477 = v470 + int32(-1)
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	v480 = v478 & int32(7)
	switch v480 {
	case 0:
		goto L135
	case 1:
		v486 = int32(4)
		goto L130
	case 2:
		goto L134
	case 3:
		goto L133
	case 4:
		goto L132
	default:
		goto L131
	}
L128:
	;
	goto L126
L129:
	;
	v511 = v510 + v459
	v514 = F_hashtableNext(m, v18+int32(304), v18)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L49
	} else {
		goto L142
	}
L130:
	;
	switch v480 {
	case 0:
		goto L141
	case 1:
		goto L140
	case 2:
		goto L139
	case 3:
		goto L138
	case 4:
		goto L137
	default:
		v506 = int32(0)
		goto L136
	}
L131:
	;
	v486 = int32(1)
	goto L130
L132:
	;
	v486 = int32(18)
	goto L130
L133:
	;
	v486 = int32(10)
	goto L130
L134:
	;
	v486 = int32(6)
	goto L130
L135:
	;
	v481 = F_zmalloc_usable_size(m, v477)
	mBase = m.M
	v510 = v481
	goto L129
L136:
	;
	v510 = v486 + v506
	goto L129
L137:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v470+int32(-9))))
	v506 = v505
	goto L136
L138:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v470+int32(-5))))
	v510 = v486 + v501
	goto L129
L139:
	;
	v497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v470+int32(-3)))))
	v510 = v486 + v497
	goto L129
L140:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470+int32(-2)))))
	v510 = v486 + v493
	goto L129
L141:
	;
	v510 = v486 + int32(base.Ui32(v478)>>(uint(int32(3))%32))
	goto L129
L142:
	;
	if v514 == int32(0) {
		goto L126
	} else {
		goto L143
	}
L143:
	;
	if base.Ui32(v469) < base.Ui32(l2) {
		v454 = v469
		v459 = v511
		goto L127
	} else {
		goto L144
	}
L144:
	;
	goto L128
L145:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v418)+20))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v418)+16))
	goto L146
L146:
	;
	v532 = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v511), base.F64_convert_i32_u(v469)), base.F64_convert_i32_u(v526+v527)), base.F64_convert_i32_u(v443))
	if base.F64_lt(v532, float64(4.294967296e+09))&base.F64_ge(v532, float64(0)) == int32(0) {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	v540 = base.I32_trunc_f64_u(v532)
	v1857 = v540
	goto L2
L148:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v602+int32(-8))))
	goto L165
L149:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v602 = v600
	goto L148
L150:
	;
	if v541&int32(1) != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v550 = int32(16)
	goto L153
L152:
	;
	v550 = int32(8)
	goto L153
L153:
	;
	v551 = l1 + v550
	if v541&int32(2) == int32(0) {
		v583 = v551
		goto L154
	} else {
		goto L155
	}
L154:
	;
	goto L164
L155:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	v557 = v551 + v556
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	switch v561 & int32(7) {
	case 0:
		goto L161
	case 1:
		goto L160
	case 2:
		goto L159
	case 3:
		goto L158
	case 4:
		goto L157
	default:
		v578 = int32(0)
		goto L156
	}
L156:
	;
	v583 = v557 + int32(1) + v578 + int32(1)
	goto L154
L157:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v557+int32(-16))))
	v578 = v577
	goto L156
L158:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v557+int32(-8))))
	v578 = v574
	goto L156
L159:
	;
	v571 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v557+int32(-4)))))
	v578 = v571
	goto L156
L160:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557+int32(-2)))))
	v578 = v568
	goto L156
L161:
	;
	v578 = int32(base.Ui32(v561) >> (uint(int32(3)) % 32))
	goto L156
L162:
	;
	v602 = v583 + v597
	goto L148
L163:
	;
	goto L162
L164:
	;
	v597 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L163
L165:
	;
	v1857 = v607&int32(2147483647) + int32(8) + v26
	goto L2
L166:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v674+int32(-8))))
	goto L183
L167:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v674 = v672
	goto L166
L168:
	;
	if v613&int32(1) != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v622 = int32(16)
	goto L171
L170:
	;
	v622 = int32(8)
	goto L171
L171:
	;
	v623 = l1 + v622
	if v613&int32(2) == int32(0) {
		v655 = v623
		goto L172
	} else {
		goto L173
	}
L172:
	;
	goto L182
L173:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623))))
	v629 = v623 + v628
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629))))
	switch v633 & int32(7) {
	case 0:
		goto L179
	case 1:
		goto L178
	case 2:
		goto L177
	case 3:
		goto L176
	case 4:
		goto L175
	default:
		v650 = int32(0)
		goto L174
	}
L174:
	;
	v655 = v629 + int32(1) + v650 + int32(1)
	goto L172
L175:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v629+int32(-16))))
	v650 = v649
	goto L174
L176:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v629+int32(-8))))
	v650 = v646
	goto L174
L177:
	;
	v643 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v629+int32(-4)))))
	v650 = v643
	goto L174
L178:
	;
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629+int32(-2)))))
	v650 = v640
	goto L174
L179:
	;
	v650 = int32(base.Ui32(v633) >> (uint(int32(3)) % 32))
	goto L174
L180:
	;
	v674 = v655 + v669
	goto L166
L181:
	;
	goto L180
L182:
	;
	v669 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L181
L183:
	;
	v1857 = v679&int32(2147483647) + int32(8) + v26
	goto L2
L184:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L185:
	;
	F__serverPanic_1(m, int32(_a838), int32(1257), int32(_a854), int32(0))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L49
	} else {
		goto L253
	}
L186:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v770&int32(4) == int32(0) {
		goto L207
	} else {
		goto L208
	}
L187:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v698&int32(4) == int32(0) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v759+int32(-8))))
	goto L205
L189:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v759 = v757
	goto L188
L190:
	;
	if v698&int32(1) != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v707 = int32(16)
	goto L193
L192:
	;
	v707 = int32(8)
	goto L193
L193:
	;
	v708 = l1 + v707
	if v698&int32(2) == int32(0) {
		v740 = v708
		goto L194
	} else {
		goto L195
	}
L194:
	;
	goto L204
L195:
	;
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708))))
	v714 = v708 + v713
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714))))
	switch v718 & int32(7) {
	case 0:
		goto L201
	case 1:
		goto L200
	case 2:
		goto L199
	case 3:
		goto L198
	case 4:
		goto L197
	default:
		v735 = int32(0)
		goto L196
	}
L196:
	;
	v740 = v714 + int32(1) + v735 + int32(1)
	goto L194
L197:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v714+int32(-16))))
	v735 = v734
	goto L196
L198:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v714+int32(-8))))
	v735 = v731
	goto L196
L199:
	;
	v728 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v714+int32(-4)))))
	v735 = v728
	goto L196
L200:
	;
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714+int32(-2)))))
	v735 = v725
	goto L196
L201:
	;
	v735 = int32(base.Ui32(v718) >> (uint(int32(3)) % 32))
	goto L196
L202:
	;
	v759 = v740 + v754
	goto L188
L203:
	;
	goto L202
L204:
	;
	v754 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L203
L205:
	;
	v1857 = v764&int32(2147483647) + int32(8) + v26
	goto L2
L206:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v832)))
	if v833&int32(4) == int32(0) {
		goto L224
	} else {
		goto L225
	}
L207:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v832 = v830
	v833 = v770
	goto L206
L208:
	;
	if v770&int32(1) != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v779 = int32(16)
	goto L211
L210:
	;
	v779 = int32(8)
	goto L211
L211:
	;
	v780 = l1 + v779
	if v770&int32(2) == int32(0) {
		v812 = v780
		goto L212
	} else {
		goto L213
	}
L212:
	;
	goto L222
L213:
	;
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780))))
	v786 = v780 + v785
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786))))
	switch v790 & int32(7) {
	case 0:
		goto L219
	case 1:
		goto L218
	case 2:
		goto L217
	case 3:
		goto L216
	case 4:
		goto L215
	default:
		v807 = int32(0)
		goto L214
	}
L214:
	;
	v812 = v786 + int32(1) + v807 + int32(1)
	goto L212
L215:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v786+int32(-16))))
	v807 = v806
	goto L214
L216:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v786+int32(-8))))
	v807 = v803
	goto L214
L217:
	;
	v800 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v786+int32(-4)))))
	v807 = v800
	goto L214
L218:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786+int32(-2)))))
	v807 = v797
	goto L214
L219:
	;
	v807 = int32(base.Ui32(v790) >> (uint(int32(3)) % 32))
	goto L214
L220:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v832 = v812 + v826
	v833 = v829
	goto L206
L221:
	;
	goto L220
L222:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L221
L223:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v895)+4))
	goto L240
L224:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v895 = v894
	goto L223
L225:
	;
	if v833&int32(1) != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v844 = int32(16)
	goto L228
L227:
	;
	v844 = int32(8)
	goto L228
L228:
	;
	v845 = l1 + v844
	if v833&int32(2) == int32(0) {
		v876 = v845
		goto L229
	} else {
		goto L230
	}
L229:
	;
	goto L239
L230:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845))))
	v851 = v845 + v850
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851))))
	switch v855 & int32(7) {
	case 0:
		goto L236
	case 1:
		goto L235
	case 2:
		goto L234
	case 3:
		goto L233
	case 4:
		goto L232
	default:
		v872 = int32(0)
		goto L231
	}
L231:
	;
	v876 = v851 + int32(1) + v872 + int32(1)
	goto L229
L232:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v851+int32(-16))))
	v872 = v871
	goto L231
L233:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v851+int32(-8))))
	v872 = v868
	goto L231
L234:
	;
	v865 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v851+int32(-4)))))
	v872 = v865
	goto L231
L235:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851+int32(-2)))))
	v872 = v862
	goto L231
L236:
	;
	v872 = int32(base.Ui32(v855) >> (uint(int32(3)) % 32))
	goto L231
L237:
	;
	v895 = v876 + v891
	goto L223
L238:
	;
	goto L237
L239:
	;
	v891 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L238
L240:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v899)+12))
	goto L241
L241:
	;
	v903 = F_hashtableMemUsage(m, v835)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L49
	} else {
		goto L242
	}
L242:
	;
	v907 = v24 + int32(280) + v903 + int32(8)
	if v900 == int32(0) {
		v1857 = v907
		goto L2
	} else {
		goto L243
	}
L243:
	;
	if l2 == int32(0) {
		v1857 = v907
		goto L2
	} else {
		goto L244
	}
L244:
	;
	v912 = int32(0)
	v915 = v900
	v917 = v912
	v920 = v912
	goto L246
L245:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v835)+20))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v835)+16))
	goto L251
L246:
	;
	v930 = v920 + int32(1)
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v915+int32(-8))))
	goto L248
L247:
	;
	goto L245
L248:
	;
	v938 = v933&int32(2147483647) + int32(8) + v917
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v915)+12))
	if v939 == int32(0) {
		goto L245
	} else {
		goto L249
	}
L249:
	;
	if base.Ui32(v930) < base.Ui32(l2) {
		v915 = v939
		v917 = v938
		v920 = v930
		goto L246
	} else {
		goto L250
	}
L250:
	;
	goto L247
L251:
	;
	v952 = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v938), base.F64_convert_i32_u(v930)), base.F64_convert_i32_u(v946+v947)), base.F64_convert_i32_u(v907))
	if base.F64_lt(v952, float64(4.294967296e+09))&base.F64_ge(v952, float64(0)) == int32(0) {
		goto L4
	} else {
		goto L252
	}
L252:
	;
	v960 = base.I32_trunc_f64_u(v952)
	v1857 = v960
	goto L2
L253:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L254:
	;
	F__serverPanic_1(m, int32(_a838), int32(1278), int32(_a855), int32(0))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L49
	} else {
		goto L358
	}
L255:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1046&int32(4) == int32(0) {
		goto L276
	} else {
		goto L277
	}
L256:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v974&int32(4) == int32(0) {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1035+int32(-8))))
	goto L274
L258:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1035 = v1033
	goto L257
L259:
	;
	if v974&int32(1) != 0 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v983 = int32(16)
	goto L262
L261:
	;
	v983 = int32(8)
	goto L262
L262:
	;
	v984 = l1 + v983
	if v974&int32(2) == int32(0) {
		v1016 = v984
		goto L263
	} else {
		goto L264
	}
L263:
	;
	goto L273
L264:
	;
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v984))))
	v990 = v984 + v989
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
	switch v994 & int32(7) {
	case 0:
		goto L270
	case 1:
		goto L269
	case 2:
		goto L268
	case 3:
		goto L267
	case 4:
		goto L266
	default:
		v1011 = int32(0)
		goto L265
	}
L265:
	;
	v1016 = v990 + int32(1) + v1011 + int32(1)
	goto L263
L266:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v990+int32(-16))))
	v1011 = v1010
	goto L265
L267:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v990+int32(-8))))
	v1011 = v1007
	goto L265
L268:
	;
	v1004 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v990+int32(-4)))))
	v1011 = v1004
	goto L265
L269:
	;
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990+int32(-2)))))
	v1011 = v1001
	goto L265
L270:
	;
	v1011 = int32(base.Ui32(v994) >> (uint(int32(3)) % 32))
	goto L265
L271:
	;
	v1035 = v1016 + v1030
	goto L257
L272:
	;
	goto L271
L273:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L272
L274:
	;
	v1857 = v1040&int32(2147483647) + int32(8) + v26
	goto L2
L275:
	;
	v1112 = v1108 + int32(44)
	goto L292
L276:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1108 = v1105
	goto L275
L277:
	;
	if v1046&int32(1) != 0 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1055 = int32(16)
	goto L280
L279:
	;
	v1055 = int32(8)
	goto L280
L280:
	;
	v1056 = l1 + v1055
	if v1046&int32(2) == int32(0) {
		v1088 = v1056
		goto L281
	} else {
		goto L282
	}
L281:
	;
	goto L291
L282:
	;
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056))))
	v1062 = v1056 + v1061
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062))))
	switch v1066 & int32(7) {
	case 0:
		goto L288
	case 1:
		goto L287
	case 2:
		goto L286
	case 3:
		goto L285
	case 4:
		goto L284
	default:
		v1083 = int32(0)
		goto L283
	}
L283:
	;
	v1088 = v1062 + int32(1) + v1083 + int32(1)
	goto L281
L284:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1062+int32(-16))))
	v1083 = v1082
	goto L283
L285:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1062+int32(-8))))
	v1083 = v1079
	goto L283
L286:
	;
	v1076 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1062+int32(-4)))))
	v1083 = v1076
	goto L283
L287:
	;
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062+int32(-2)))))
	v1083 = v1073
	goto L283
L288:
	;
	v1083 = int32(base.Ui32(v1066) >> (uint(int32(3)) % 32))
	goto L283
L289:
	;
	v1108 = v1088 + v1102
	goto L275
L290:
	;
	goto L289
L291:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L290
L292:
	;
	v1114 = v18 + int32(304)
	v1115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1114)+14)) = uint8(v1115)
	*(*int32)(unsafe.Add(mBase, uint32(v1114))) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v1114)+24)) = v1115
	*(*uint8)(unsafe.Add(mBase, uint32(v1114)+15)) = uint8(v1115)
	*(*int32)(unsafe.Add(mBase, uint32(v1114)+8)) = int32(-1)
	if v1108 == v1115 {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	v1133 = F_hashtableMemUsage(m, v1108)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L49
	} else {
		goto L297
	}
L294:
	;
	goto L293
L295:
	;
	goto L294
L297:
	;
	v1135 = v1133 + v26
	v1138 = F_hashtableNext(m, v18+int32(304), v18)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L49
	} else {
		goto L298
	}
L298:
	;
	if v1138 == int32(0) {
		goto L7
	} else {
		goto L299
	}
L299:
	;
	if l2 == int32(0) {
		goto L7
	} else {
		goto L300
	}
L300:
	;
	v1146 = int32(0)
	v1151 = int32(0)
	goto L302
L301:
	;
	F_hashtableCleanupIterator(m, v18+int32(304))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L49
	} else {
		goto L354
	}
L302:
	;
	v1161 = v1146 + int32(1)
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v1168 = v1162 + int32(-1)
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168))))
	v1171 = v1169 & int32(7)
	if v1171 == int32(0) {
		goto L307
	} else {
		goto L308
	}
L303:
	;
	goto L301
L304:
	;
	v1296 = v1295 + v1151
	v1299 = F_hashtableNext(m, v18+int32(304), v18)
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L49
	} else {
		goto L351
	}
L305:
	;
	v1238 = v1233 & int32(7)
	if v1238 == int32(0) {
		goto L334
	} else {
		goto L335
	}
L306:
	;
	v1207 = F_sdsAllocPtr(m, v1162)
	mBase = m.M
	v1210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168))))
	if int32(base.Ui32(v1210&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L322
	} else {
		goto L323
	}
L307:
	;
	switch v1171 {
	case 0:
		goto L315
	case 1:
		goto L314
	case 2:
		goto L313
	case 3:
		goto L312
	case 4:
		goto L311
	default:
		v1191 = int32(0)
		goto L310
	}
L308:
	;
	if v1169&int32(16) != 0 {
		goto L306
	} else {
		goto L309
	}
L309:
	;
	goto L307
L310:
	;
	v1192 = F_sdsHdrSize(m, v1171)
	mBase = m.M
	v1193 = v1191 + v1192
	v1197 = v1193 + int32(1)
	v1198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168))))
	if int32(base.Ui32(v1198&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L316
	} else {
		goto L317
	}
L311:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1162+int32(-17))))
	v1191 = v1190
	goto L310
L312:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1162+int32(-9))))
	v1191 = v1187
	goto L310
L313:
	;
	v1184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1162+int32(-5)))))
	v1191 = v1184
	goto L310
L314:
	;
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1162+int32(-3)))))
	v1191 = v1181
	goto L310
L315:
	;
	v1191 = int32(base.Ui32(v1169) >> (uint(int32(3)) % 32))
	goto L310
L316:
	;
	v1203 = v1193 + int32(9)
	goto L318
L317:
	;
	v1203 = v1197
	goto L318
L318:
	;
	if v1198&int32(7) != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1206 = v1203
	goto L321
L320:
	;
	v1206 = v1197
	goto L321
L321:
	;
	v1233 = v1198
	v1235 = v1206
	goto L305
L322:
	;
	v1215 = int32(-4)
	goto L324
L323:
	;
	v1215 = int32(0)
	goto L324
L324:
	;
	v1218 = v1210 & int32(7)
	if v1218 != 0 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1219 = v1215
	goto L327
L326:
	;
	v1219 = int32(0)
	goto L327
L327:
	;
	if int32(base.Ui32(v1210&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v1227 = int32(-8)
	goto L330
L329:
	;
	v1227 = int32(0)
	goto L330
L330:
	;
	if v1218 != 0 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1229 = v1227
	goto L333
L332:
	;
	v1229 = int32(0)
	goto L333
L333:
	;
	v1231 = F_zmalloc_usable_size(m, v1207+v1219+v1229)
	mBase = m.M
	v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168))))
	v1233 = v1232
	v1235 = v1231
	goto L305
L334:
	;
	if v1238 == int32(0) {
		goto L339
	} else {
		goto L340
	}
L335:
	;
	v1241 = int32(48)
	if v1233&v1241 != v1241 {
		goto L334
	} else {
		goto L336
	}
L336:
	;
	v1245 = F_sdsAllocPtr(m, v1162)
	mBase = m.M
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1245+int32(-4))))
	v1249 = F_zmalloc_usable_size(m, v1248)
	mBase = m.M
	v1295 = v1249 + v1235
	goto L304
L337:
	;
	v1291 = F_sdsAllocSize(m, v1289)
	mBase = m.M
	v1295 = v1291 + v1235
	goto L304
L338:
	;
	v1279 = F_sdsAllocPtr(m, v1162)
	mBase = m.M
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1279+int32(-4))))
	if v1233&int32(32) == int32(0) {
		v1289 = v1282
		goto L337
	} else {
		goto L348
	}
L339:
	;
	switch v1238 {
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
		v1272 = int32(0)
		goto L342
	}
L340:
	;
	if v1233&int32(16) != 0 {
		goto L338
	} else {
		goto L341
	}
L341:
	;
	goto L339
L342:
	;
	v1274 = int32(1)
	v1275 = F_sdsHdrSize(m, v1274)
	mBase = m.M
	v1289 = v1162 + v1272 + v1275 + v1274
	goto L337
L343:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1162+int32(-17))))
	v1272 = v1271
	goto L342
L344:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1162+int32(-9))))
	v1272 = v1268
	goto L342
L345:
	;
	v1265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1162+int32(-5)))))
	v1272 = v1265
	goto L342
L346:
	;
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1162+int32(-3)))))
	v1272 = v1262
	goto L342
L347:
	;
	v1272 = int32(base.Ui32(v1233&int32(248)) >> (uint(int32(3)) % 32))
	goto L342
L348:
	;
	if v1282 != 0 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1282)))
	v1289 = v1288
	goto L337
L350:
	;
	v1289 = int32(0)
	goto L337
L351:
	;
	if v1299 == int32(0) {
		goto L301
	} else {
		goto L352
	}
L352:
	;
	if base.Ui32(v1161) < base.Ui32(l2) {
		v1146 = v1161
		v1151 = v1296
		goto L302
	} else {
		goto L353
	}
L353:
	;
	goto L303
L354:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+20))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+16))
	goto L356
L355:
	;
	v1823 = int32(0)
	goto L3
L356:
	;
	v1317 = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v1296), base.F64_convert_i32_u(v1161)), base.F64_convert_i32_u(v1311+v1312)), base.F64_convert_i32_u(v1135))
	if base.F64_lt(v1317, float64(4.294967296e+09))&base.F64_ge(v1317, float64(0)) == int32(0) {
		goto L355
	} else {
		goto L357
	}
L357:
	;
	v1325 = base.I32_trunc_f64_u(v1317)
	v1823 = v1325
	goto L3
L358:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L359:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+4))
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1399)+24))
	goto L376
L360:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1398 = v1393
	goto L359
L361:
	;
	if v1334&int32(1) != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1343 = int32(16)
	goto L364
L363:
	;
	v1343 = int32(8)
	goto L364
L364:
	;
	v1344 = l1 + v1343
	if v1334&int32(2) == int32(0) {
		v1376 = v1344
		goto L365
	} else {
		goto L366
	}
L365:
	;
	goto L375
L366:
	;
	v1349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1344))))
	v1350 = v1344 + v1349
	v1354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1350))))
	switch v1354 & int32(7) {
	case 0:
		goto L372
	case 1:
		goto L371
	case 2:
		goto L370
	case 3:
		goto L369
	case 4:
		goto L368
	default:
		v1371 = int32(0)
		goto L367
	}
L367:
	;
	v1376 = v1350 + int32(1) + v1371 + int32(1)
	goto L365
L368:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1350+int32(-16))))
	v1371 = v1370
	goto L367
L369:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1350+int32(-8))))
	v1371 = v1367
	goto L367
L370:
	;
	v1364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1350+int32(-4)))))
	v1371 = v1364
	goto L367
L371:
	;
	v1361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1350+int32(-2)))))
	v1371 = v1361
	goto L367
L372:
	;
	v1371 = int32(base.Ui32(v1354) >> (uint(int32(3)) % 32))
	goto L367
L373:
	;
	v1398 = v1376 + v1390
	goto L359
L374:
	;
	goto L373
L375:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L374
L376:
	;
	v1402 = v18 + int32(304)
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+4)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v1402))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+20)) = int32(128)
	v1409 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1402)+12)) = v1409
	*(*int64)(unsafe.Add(mBase, uint32(v1402)+296)) = v1409
	*(*int64)(unsafe.Add(mBase, uint32(v1402)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+8)) = v18 + int32(328)
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+156)) = v18 + int32(472)
	goto L377
L377:
	;
	v1421 = int32(0)
	v1427 = F_raxSeek(m, v18+int32(304), int32(_a4), v1421, v1421)
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L49
	} else {
		goto L378
	}
L378:
	;
	v1431 = v26 + v1400 + int32(72)
	if l2 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L379:
	;
	v1561 = v1488 + v1431
	goto L5
L380:
	;
	v1477 = int32(0)
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+4))
	v1479 = *(*int64)(unsafe.Add(mBase, uint32(v1478)+8))
	if v1479 != int64(0) {
		v1522 = v1477
		v1526 = v1479
		goto L6
	} else {
		goto L391
	}
L381:
	;
	v1436 = v1421
	v1441 = int32(0)
	goto L383
L382:
	;
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+4))
	v1471 = *(*int64)(unsafe.Add(mBase, uint32(v1470)+8))
	if base.Ui64(v1471) <= base.Ui64(base.I64_extend_i32_u(v1468)) {
		v1488 = v1469
		goto L379
	} else {
		goto L389
	}
L383:
	;
	v1452 = F_raxNext(m, v18+int32(304))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L49
	} else {
		goto L385
	}
L384:
	;
	v1468 = l2
	v1469 = v1464
	goto L382
L385:
	;
	if v1452 == int32(0) {
		v1468 = v1436
		v1469 = v1441
		goto L382
	} else {
		goto L386
	}
L386:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v18)+316))
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1456+int32(-8))))
	goto L387
L387:
	;
	v1464 = v1459&int32(2147483647) + int32(8) + v1441
	v1466 = v1436 + int32(1)
	if v1466 != l2 {
		v1436 = v1466
		v1441 = v1464
		goto L383
	} else {
		goto L388
	}
L388:
	;
	goto L384
L389:
	;
	if v1468 == int32(0) {
		v1522 = v1469
		v1526 = v1471
		goto L6
	} else {
		goto L390
	}
L390:
	;
	v1476 = base.I32_div_u_s(v1469, v1468)
	v1522 = v1476
	v1526 = v1471
	goto L6
L391:
	;
	v1488 = v1477
	goto L379
L392:
	;
	v1857 = v1498 + v26
	goto L2
L393:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L394:
	;
	v1857 = v443
	goto L2
L395:
	;
	v1823 = v1135
	goto L3
L396:
	;
	v1540 = F_raxNext(m, v18+int32(304))
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L49
	} else {
		goto L397
	}
L397:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v18)+316))
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1547+int32(-8))))
	goto L398
L398:
	;
	v1561 = v1522*(base.I32_wrap_i64(v1526)+int32(-1)) + v1431 + (v1550&int32(2147483647) + int32(8))
	goto L5
L399:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1398)))
	if v1575 == int32(0) {
		v1857 = v1561
		goto L2
	} else {
		goto L400
	}
L400:
	;
	v1579 = v18 + int32(304)
	*(*int32)(unsafe.Add(mBase, uint32(v1579)+4)) = v1575
	*(*int32)(unsafe.Add(mBase, uint32(v1579))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1579)+20)) = int32(128)
	v1585 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1579)+12)) = v1585
	*(*int64)(unsafe.Add(mBase, uint32(v1579)+296)) = v1585
	*(*int64)(unsafe.Add(mBase, uint32(v1579)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v1579)+8)) = v18 + int32(328)
	*(*int32)(unsafe.Add(mBase, uint32(v1579)+156)) = v18 + int32(472)
	goto L401
L401:
	;
	v1597 = int32(0)
	v1603 = F_raxSeek(m, v18+int32(304), int32(_a4), v1597, v1597)
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L49
	} else {
		goto L402
	}
L402:
	;
	if l2 != 0 {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v1621 = v1597
	v1622 = int32(0)
	goto L408
L404:
	;
	F_raxStop(m, v18+int32(304))
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L49
	} else {
		goto L405
	}
L405:
	;
	v1857 = v1561
	goto L2
L406:
	;
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v1398)))
	v1789 = *(*int64)(unsafe.Add(mBase, uint32(v1788)+8))
	goto L441
L407:
	;
	F_raxStop(m, v18+int32(304))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L49
	} else {
		goto L439
	}
L408:
	;
	v1627 = F_raxNext(m, v18+int32(304))
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L49
	} else {
		goto L410
	}
L409:
	;
	F_raxStop(m, v18+int32(304))
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L49
	} else {
		goto L438
	}
L410:
	;
	if v1627 == int32(0) {
		goto L407
	} else {
		goto L411
	}
L411:
	;
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v18)+316))
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+24))
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1632)+24))
	goto L412
L412:
	;
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+24))
	v1635 = *(*int64)(unsafe.Add(mBase, uint32(v1634)+8))
	goto L413
L413:
	;
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(128)
	v1642 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+12)) = v1642
	*(*int64)(unsafe.Add(mBase, uint32(v18)+296)) = v1642
	*(*int64)(unsafe.Add(mBase, uint32(v18)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v18 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+156)) = v18 + int32(168)
	goto L414
L414:
	;
	v1654 = int32(0)
	v1658 = F_raxSeek(m, v18, int32(_a4), v1654, v1654)
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L49
	} else {
		goto L415
	}
L415:
	;
	v1666 = v1621 + v1633 + base.I32_wrap_i64(v1635)*int32(24) + int32(32)
	v1671 = v1654
	v1674 = int32(0)
	goto L419
L416:
	;
	v1758 = v1622 + int32(1)
	if v1758 != l2 {
		v1621 = v1755
		v1622 = v1758
		goto L408
	} else {
		goto L437
	}
L417:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+28))
	v1735 = *(*int64)(unsafe.Add(mBase, uint32(v1734)+8))
	goto L435
L418:
	;
	F_raxStop(m, v18)
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L49
	} else {
		goto L432
	}
L419:
	;
	v1683 = F_raxNext(m, v18)
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L49
	} else {
		goto L421
	}
L420:
	;
	F_raxStop(m, v18)
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L49
	} else {
		goto L431
	}
L421:
	;
	if v1683 == int32(0) {
		goto L418
	} else {
		goto L422
	}
L422:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1688)+16))
	v1692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1689+int32(-1)))))
	switch v1692 & int32(7) {
	case 0:
		goto L428
	case 1:
		goto L427
	case 2:
		goto L426
	case 3:
		goto L425
	case 4:
		goto L424
	default:
		v1709 = int32(0)
		goto L423
	}
L423:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1688)+20))
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+24))
	goto L429
L424:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1689+int32(-17))))
	v1709 = v1708
	goto L423
L425:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v1689+int32(-9))))
	v1709 = v1705
	goto L423
L426:
	;
	v1702 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1689+int32(-5)))))
	v1709 = v1702
	goto L423
L427:
	;
	v1699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1689+int32(-3)))))
	v1709 = v1699
	goto L423
L428:
	;
	v1709 = int32(base.Ui32(v1692) >> (uint(int32(3)) % 32))
	goto L423
L429:
	;
	v1715 = v1671 + int32(24) + v1709 + v1714
	v1717 = v1674 + int32(1)
	if v1717 != l2 {
		v1671 = v1715
		v1674 = v1717
		goto L419
	} else {
		goto L430
	}
L430:
	;
	goto L420
L431:
	;
	v1727 = v1715
	v1728 = l2
	goto L417
L432:
	;
	if v1674 == int32(0) {
		v1755 = v1666
		goto L416
	} else {
		goto L433
	}
L433:
	;
	v1727 = v1671
	v1728 = v1674
	goto L417
L434:
	;
	v1755 = int32(0)
	goto L416
L435:
	;
	v1739 = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v1727), base.F64_convert_i32_u(v1728)), base.F64_convert_i64_u(v1735)), base.F64_convert_i32_u(v1666))
	if base.F64_lt(v1739, float64(4.294967296e+09))&base.F64_ge(v1739, float64(0)) == int32(0) {
		goto L434
	} else {
		goto L436
	}
L436:
	;
	v1747 = base.I32_trunc_f64_u(v1739)
	v1755 = v1747
	goto L416
L437:
	;
	goto L409
L438:
	;
	v1781 = v1755
	v1782 = l2
	goto L406
L439:
	;
	if v1622 == int32(0) {
		v1857 = v1561
		goto L2
	} else {
		goto L440
	}
L440:
	;
	v1781 = v1621
	v1782 = v1622
	goto L406
L441:
	;
	v1793 = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v1781), base.F64_convert_i32_u(v1782)), base.F64_convert_i64_u(v1789)), base.F64_convert_i32_u(v1561))
	if base.F64_lt(v1793, float64(4.294967296e+09))&base.F64_ge(v1793, float64(0)) == int32(0) {
		goto L4
	} else {
		goto L442
	}
L442:
	;
	v1801 = base.I32_trunc_f64_u(v1793)
	v1857 = v1801
	goto L2
L443:
	;
	if v1846 == int32(0) {
		v1857 = v1823
		goto L2
	} else {
		goto L449
	}
L444:
	;
	goto L443
L445:
	;
	v1846 = int32(0)
	goto L444
L446:
	;
	v1836 = int32(1)
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1112)))
	switch v1837 + v1836 {
	case 0:
		v1846 = v1836
		goto L444
	case 1:
		goto L445
	default:
		goto L447
	}
L447:
	;
	if v1837&int32(7) != 0 {
		v1846 = v1836
		goto L444
	} else {
		goto L448
	}
L448:
	;
	goto L445
L449:
	;
	v1849 = F_vsetMemUsage(m, v1112)
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L49
	} else {
		goto L450
	}
L450:
	;
	v1857 = v1849 + v1823
	goto L2
}
func F_objectGetIdleness(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = int32(base.Ui32(v8) >> (uint(int32(8)) % 32))
	v12 = v6 + int32(12)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[334])))
	if v18 != int32(1) {
		v49 = *(*int32)(unsafe.Add(mBase, _consts[333]))
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = (v49 - v10) & int32(16777215)
		v58 = v10
	} else {
		v21 = int32(0)
		v22 = int32(*(*uint16)(unsafe.Add(mBase, _consts[335])))
		v24 = *(*int32)(unsafe.Add(mBase, _consts[336]))
		if v24 == v21 {
			v35 = v2
		} else {
			v30 = int32(65535)
			v32 = base.I32_div_s((v22-int32(base.Ui32(v10)>>(uint(int32(8))%32)))&v30, v24)
			v35 = v32 & v30
		}
		v38 = v10 & int32(255)
		v39 = v38 - v35
		if base.Ui32(v38) < base.Ui32(v39) {
			v41 = int32(0)
		} else {
			v41 = v39
		}
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v41 ^ int32(255)
		v58 = v41 | v22<<(uint(int32(8))%32)
	}
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v59 | v58<<(uint(int32(8))%32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	m.G0 = v6 + int32(16)
	return v64
}
func F_objectGetVal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5&int32(4) == int32(0) {
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		return v65
	} else {
		if v5&int32(1) != 0 {
			v14 = int32(16)
		} else {
			v14 = int32(8)
		}
		v15 = l0 + v14
		if v5&int32(2) == int32(0) {
			v46 = v15
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v21 = v15 + v20
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
			switch v25 & int32(7) {
			case 0:
				v42 = int32(base.Ui32(v25) >> (uint(int32(3)) % 32))
			case 1:
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+int32(-2)))))
				v42 = v32
			case 2:
				v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+int32(-4)))))
				v42 = v35
			case 3:
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(-8))))
				v42 = v38
			case 4:
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(-16))))
				v42 = v41
			default:
				v42 = int32(0)
			}
			v46 = v21 + int32(1) + v42 + int32(1)
		}
		v61 = *(*int32)(unsafe.Add(mBase, _consts[249]))
		return v46 + v61
	}
}
func F_objectSetLRUOrLFU(m *base.Module, l0 int32, l1 int64, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _consts[334])))
	if v6 == int32(0) {
		if int64(0) <= l2 {
			v27 = *(*int32)(unsafe.Add(mBase, _consts[333]))
			v31 = (v27 - base.I32_wrap_i64(l2)) & int32(16777215)
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v32 | v31<<(uint(int32(8))%32)
			return int32(1)
		} else {
			return int32(0)
		}
	} else {
		if int64(0) <= l1 {
			if base.Ui64(int64(256)) <= base.Ui64(l1) {
				F__serverAssert(m, int32(_a856), int32(_a838), int32(1671))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v17 = int32(*(*uint16)(unsafe.Add(mBase, _consts[335])))
				v31 = v17<<(uint(int32(8))%32) | base.I32_wrap_i64(l1)
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v32 | v31<<(uint(int32(8))%32)
				return int32(1)
			}
		} else {
			return int32(0)
		}
	}
}
func F_objectSetVal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v3&int32(4) == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
		return
	} else {
		F__serverAssert(m, int32(_a840), int32(_a838), int32(326))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
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
func F_objectUnembedVal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
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
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
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
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7&int32(4) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a841), int32(_a838), int32(338))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L40
	} else {
		goto L88
	}
L2:
	;
	F__serverAssert(m, int32(_a842), int32(_a838), int32(335))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L40
	} else {
		goto L87
	}
L3:
	;
	F__serverAssert(m, int32(_a843), int32(_a838), int32(334))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L40
	} else {
		goto L86
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12&int32(240) != int32(128) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v7&int32(1) != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v21 = int32(16)
	goto L8
L7:
	;
	v21 = int32(8)
	goto L8
L8:
	;
	v22 = l0 + v21
	if v7&int32(2) == int32(0) {
		v54 = v22
		goto L9
	} else {
		goto L10
	}
L9:
	;
	goto L19
L10:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v28 = v22 + v27
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	switch v32 & int32(7) {
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
		v49 = int32(0)
		goto L11
	}
L11:
	;
	v54 = v28 + int32(1) + v49 + int32(1)
	goto L9
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-16))))
	v49 = v48
	goto L11
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-8))))
	v49 = v45
	goto L11
L14:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-4)))))
	v49 = v42
	goto L11
L15:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-2)))))
	v49 = v39
	goto L11
L16:
	;
	v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	goto L11
L17:
	;
	v70 = v54 + v68
	v77 = v70 + int32(-1)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v80 = v78 & int32(7)
	switch v80 {
	case 0:
		goto L26
	case 1:
		v86 = int32(4)
		goto L21
	case 2:
		goto L25
	case 3:
		goto L24
	case 4:
		goto L23
	default:
		goto L22
	}
L18:
	;
	goto L17
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L18
L20:
	;
	if base.Ui32(v110) <= base.Ui32(int32(3)) {
		goto L1
	} else {
		goto L33
	}
L21:
	;
	switch v80 {
	case 0:
		goto L32
	case 1:
		goto L31
	case 2:
		goto L30
	case 3:
		goto L29
	case 4:
		goto L28
	default:
		v106 = int32(0)
		goto L27
	}
L22:
	;
	v86 = int32(1)
	goto L21
L23:
	;
	v86 = int32(18)
	goto L21
L24:
	;
	v86 = int32(10)
	goto L21
L25:
	;
	v86 = int32(6)
	goto L21
L26:
	;
	v81 = F_zmalloc_usable_size(m, v77)
	mBase = m.M
	v110 = v81
	goto L20
L27:
	;
	v110 = v86 + v106
	goto L20
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-9))))
	v106 = v105
	goto L27
L29:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-5))))
	v110 = v86 + v101
	goto L20
L30:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70+int32(-3)))))
	v110 = v86 + v97
	goto L20
L31:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+int32(-2)))))
	v110 = v86 + v93
	goto L20
L32:
	;
	v110 = v86 + int32(base.Ui32(v78)>>(uint(int32(3))%32))
	goto L20
L33:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+int32(-1)))))
	switch v116 & int32(7) {
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
		v133 = int32(0)
		goto L34
	}
L34:
	;
	v134 = F_sdsnewlen(m, v70, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-17))))
	v133 = v132
	goto L34
L36:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-9))))
	v133 = v129
	goto L34
L37:
	;
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70+int32(-5)))))
	v133 = v126
	goto L34
L38:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+int32(-3)))))
	v133 = v123
	goto L34
L39:
	;
	v133 = int32(base.Ui32(v116) >> (uint(int32(3)) % 32))
	goto L34
L40:
	;
	return
L41:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L44
L42:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v151 = int32(4)
	v153 = int32(12)
	v155 = l0 + (v150&v151 ^ v153)
	v157 = v155 + v151
	v164 = v70 - (v148 + (l0 + (v136&v151 ^ v153)))
	if v157 == v155 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	goto L42
L44:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v134
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v314 & int32(-5)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v318 & int32(-241)
	return
L46:
	;
	goto L45
L47:
	;
	v168 = v164 + v157
	if base.Ui32(int32(0)-v164<<(uint(int32(1))%32)) < base.Ui32(v155-v168) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v178 = (v155 ^ v157) & int32(3)
	if base.Ui32(v155) <= base.Ui32(v157) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v175 = F___memcpy(m, v157, v155, v164)
	mBase = m.M
	goto L45
L50:
	;
	if v284 == int32(0) {
		goto L46
	} else {
		goto L82
	}
L51:
	;
	if base.Ui32(v262) <= base.Ui32(int32(3)) {
		v283 = v261
		v284 = v262
		v285 = v263
		goto L50
	} else {
		goto L78
	}
L52:
	;
	if v178 != 0 {
		v244 = v164
		goto L62
	} else {
		goto L63
	}
L53:
	;
	if v178 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v157&int32(3) != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v283 = v155
	v284 = v164
	v285 = v157
	goto L50
L56:
	;
	v185 = v155
	v186 = v164
	v187 = v157
	goto L58
L57:
	;
	v261 = v155
	v262 = v164
	v263 = v157
	goto L51
L58:
	;
	if v186 == int32(0) {
		goto L46
	} else {
		goto L60
	}
L60:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	*(*uint8)(unsafe.Add(mBase, uint32(v187))) = uint8(v191)
	v193 = int32(1)
	v194 = v185 + v193
	v196 = v186 + int32(-1)
	v198 = v187 + v193
	if v198&int32(3) == int32(0) {
		v261 = v194
		v262 = v196
		v263 = v198
		goto L51
	} else {
		goto L61
	}
L61:
	;
	v185 = v194
	v186 = v196
	v187 = v198
	goto L58
L62:
	;
	if v244 == int32(0) {
		goto L46
	} else {
		goto L74
	}
L63:
	;
	if v168&int32(3) == int32(0) {
		v224 = v164
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if base.Ui32(v224) <= base.Ui32(int32(3)) {
		v244 = v224
		goto L62
	} else {
		goto L70
	}
L65:
	;
	v209 = v164
	goto L66
L66:
	;
	if v209 == int32(0) {
		goto L46
	} else {
		goto L68
	}
L67:
	;
	v224 = v215
	goto L64
L68:
	;
	v215 = v209 + int32(-1)
	v216 = v157 + v215
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+v215))))
	*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v218)
	if v216&int32(3) != 0 {
		v209 = v215
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v231 = v224
	goto L71
L71:
	;
	v235 = v231 + int32(-4)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v155+v235)))
	*(*int32)(unsafe.Add(mBase, uint32(v157+v235))) = v238
	if base.Ui32(int32(3)) < base.Ui32(v235) {
		v231 = v235
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v244 = v235
	goto L62
L73:
	;
	goto L72
L74:
	;
	v251 = v244
	goto L75
L75:
	;
	v255 = v251 + int32(-1)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+v255))))
	*(*uint8)(unsafe.Add(mBase, uint32(v157+v255))) = uint8(v258)
	if v255 != 0 {
		v251 = v255
		goto L75
	} else {
		goto L77
	}
L77:
	;
	goto L46
L78:
	;
	v268 = v261
	v269 = v262
	v270 = v263
	goto L79
L79:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = v272
	v274 = int32(4)
	v275 = v268 + v274
	v277 = v270 + v274
	v279 = v269 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v279) {
		v268 = v275
		v269 = v279
		v270 = v277
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v283 = v275
	v284 = v279
	v285 = v277
	goto L50
L81:
	;
	goto L80
L82:
	;
	v290 = v283
	v291 = v284
	v292 = v285
	goto L83
L83:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	*(*uint8)(unsafe.Add(mBase, uint32(v292))) = uint8(v294)
	v296 = int32(1)
	v301 = v291 + int32(-1)
	if v301 != 0 {
		v290 = v290 + v296
		v291 = v301
		v292 = v292 + v296
		goto L83
	} else {
		goto L85
	}
L84:
	;
	goto L46
L85:
	;
	goto L84
L86:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_rewriteObjectRio(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int64
	_ = v198
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int64
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v283 int32
	_ = v283
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18&int32(2) == v4 {
		v38 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v38
	*(*int64)(unsafe.Add(mBase, uint32(v12)+4)) = int64(-68719476736)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v45&int32(1) == int32(0) {
		v56 = int64(-1)
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L1
L3:
	;
	v32 = l1 + (v18&int32(4) ^ int32(12)) + v18<<(uint(int32(3))%32)&int32(8)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v38 = v32 + v33 + int32(1)
	goto L2
L4:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v57 & int32(15) {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	case 5:
		goto L11
	case 6:
		goto L12
	default:
		goto L9
	}
L5:
	;
	goto L4
L6:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l1+(v45&int32(4)^int32(12)))))
	v56 = v55
	goto L5
L7:
	;
	m.G0 = v12 + int32(48)
	return v283
L8:
	;
	v283 = int32(-1)
	goto L7
L9:
	;
	F__serverPanic_1(m, int32(_a68), int32(2393), int32(_a123), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L29
	} else {
		goto L82
	}
L10:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	if v185 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L11:
	;
	v144 = F_objectGetVal(m, l1)
	mBase = m.M
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = l2
	v149 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v149
	*(*int64)(unsafe.Add(mBase, uint32(v12)+28)) = int64(0)
	v156 = v12 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v156
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v145)+20))
	m.T0[v163].(func(*base.Module, int32, int32, int32))(m, v12+int32(16), v156, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L29
	} else {
		goto L49
	}
L12:
	;
	v140 = F_rewriteStreamObject(m, l0, v12+int32(4), l1)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L29
	} else {
		goto L47
	}
L13:
	;
	v134 = F_rewriteHashObject(m, l0, v12+int32(4), l1)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L29
	} else {
		goto L45
	}
L14:
	;
	v128 = F_rewriteSortedSetObject(m, l0, v12+int32(4), l1)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L29
	} else {
		goto L43
	}
L15:
	;
	v122 = F_rewriteSetObject(m, l0, v12+int32(4), l1)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L41
	}
L16:
	;
	v116 = F_rewriteListObject(m, l0, v12+int32(4), l1)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L29
	} else {
		goto L39
	}
L17:
	;
	v60 = int32(0)
	v61 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+22)) = v61
	v64 = *(*int64)(unsafe.Add(mBase, _consts[48]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v64
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v66&int32(6) != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v78 = v12 + int32(16)
	v79 = int32(13)
	goto L19
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v81) < base.Ui32(v79) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v106 = F_rioWriteBulkObject(m, l0, v12+int32(4))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L29
	} else {
		goto L35
	}
L21:
	;
	v83 = v81
	goto L23
L22:
	;
	v83 = v79
	goto L23
L23:
	;
	if v81 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v84 = v83
	goto L26
L25:
	;
	v84 = v79
	goto L26
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v85 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v93 = m.T0[v92].(func(*base.Module, int32, int32, int32) int32)(m, l0, v78, v84)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L29
	} else {
		goto L32
	}
L28:
	;
	m.T0[v85].(func(*base.Module, int32, int32, int32))(m, l0, v78, v84)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	goto L27
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v99 + v84
	v103 = v79 - v84
	if v103 != 0 {
		v78 = v78 + v84
		v79 = v103
		goto L19
	} else {
		goto L34
	}
L32:
	;
	if v93 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v95 | int64(2)
	goto L8
L34:
	;
	goto L20
L35:
	;
	if v106 == int32(0) {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v110 = F_rioWriteBulkObject(m, l0, l1)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	if v110 == int32(0) {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	goto L10
L39:
	;
	if v116 == int32(0) {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	goto L10
L41:
	;
	if v122 == int32(0) {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	goto L10
L43:
	;
	if v128 == int32(0) {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	goto L10
L45:
	;
	if v134 == int32(0) {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	goto L10
L47:
	;
	if v140 == int32(0) {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	goto L10
L49:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v166 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v174 != 0 {
		goto L8
	} else {
		goto L54
	}
L51:
	;
	F_moduleFreeContext(m, v166)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L29
	} else {
		goto L52
	}
L52:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	F_valkey_free(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L29
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	goto L10
L55:
	;
	if v56 == int64(-1) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	goto L55
L58:
	;
	v259 = int32(0)
	v261 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	if v261 == v259 {
		v283 = v259
		goto L7
	} else {
		goto L80
	}
L59:
	;
	v194 = int32(0)
	v195 = *(*int32)(unsafe.Add(mBase, _consts[50]))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(32)))) = v195
	v198 = *(*int64)(unsafe.Add(mBase, _consts[51]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v198
	v201 = *(*int64)(unsafe.Add(mBase, _consts[52]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v201
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v203&int32(6) != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	v211 = int32(19)
	v216 = v12 + int32(16)
	goto L61
L61:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v218) < base.Ui32(v211) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v242 = F_rioWriteBulkObject(m, l0, v12+int32(4))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L29
	} else {
		goto L76
	}
L63:
	;
	v220 = v218
	goto L65
L64:
	;
	v220 = v211
	goto L65
L65:
	;
	if v218 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v221 = v220
	goto L68
L67:
	;
	v221 = v211
	goto L68
L68:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v222 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v228 = m.T0[v227].(func(*base.Module, int32, int32, int32) int32)(m, l0, v216, v221)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L29
	} else {
		goto L73
	}
L70:
	;
	m.T0[v222].(func(*base.Module, int32, int32, int32))(m, l0, v216, v221)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L29
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v235 + v221
	v239 = v211 - v221
	if v239 != 0 {
		v211 = v239
		v216 = v216 + v221
		goto L61
	} else {
		goto L75
	}
L73:
	;
	if v228 != 0 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v230 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v230 | int64(2)
	v283 = int32(-1)
	goto L7
L75:
	;
	goto L62
L76:
	;
	if v242 == int32(0) {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	v246 = F_rioWriteBulkLongLong(m, l0, v56)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L29
	} else {
		goto L78
	}
L78:
	;
	if v246 == int32(0) {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	goto L58
L80:
	;
	F_debugDelay(m, v261)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L29
	} else {
		goto L81
	}
L81:
	;
	v283 = v259
	goto L7
L82:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tryObjectEncodingEx(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
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
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v158 int64
	_ = v158
	var v164 int32
	_ = v164
	var v166 int64
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v180 int64
	_ = v180
	var v185 int64
	_ = v185
	var v189 int32
	_ = v189
	var v191 int64
	_ = v191
	var v193 int32
	_ = v193
	var v201 int64
	_ = v201
	var v221 int64
	_ = v221
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int64
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v406 int32
	_ = v406
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12&int32(4) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v76&int32(15) != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v74 = v71
	goto L1
L3:
	;
	if v12&int32(1) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = int32(16)
	goto L6
L5:
	;
	v21 = int32(8)
	goto L6
L6:
	;
	v22 = l0 + v21
	if v12&int32(2) == int32(0) {
		v54 = v22
		goto L7
	} else {
		goto L8
	}
L7:
	;
	goto L17
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v28 = v22 + v27
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	switch v32 & int32(7) {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	default:
		v49 = int32(0)
		goto L9
	}
L9:
	;
	v54 = v28 + int32(1) + v49 + int32(1)
	goto L7
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-16))))
	v49 = v48
	goto L9
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-8))))
	v49 = v45
	goto L9
L12:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-4)))))
	v49 = v42
	goto L9
L13:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-2)))))
	v49 = v39
	goto L9
L14:
	;
	v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	goto L9
L15:
	;
	v74 = v54 + v68
	goto L1
L16:
	;
	goto L15
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L16
L18:
	;
	F__serverAssertWithInfo(m, int32(0), l0, int32(_a479), int32(_a838), int32(874))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L79
	} else {
		goto L99
	}
L19:
	;
	switch int32(base.Ui32(v76)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		goto L21
	default:
		v393 = l0
		goto L20
	}
L20:
	;
	m.G0 = v10 + int32(16)
	return v393
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(int32(15)) < base.Ui32(v83) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v393 = l0
	goto L20
L23:
	;
	v88 = v74 + int32(-1)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	switch v89 & int32(7) {
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
		v109 = int32(0)
		goto L27
	}
L24:
	;
	if l1 == int32(0) {
		goto L22
	} else {
		goto L97
	}
L25:
	;
	goto L86
L26:
	;
	if base.Ui32(int32(255)) < base.Ui32(v106) {
		goto L24
	} else {
		goto L83
	}
L27:
	;
	v112 = int32(0)
	if base.Ui32(v109+int32(-21)) < base.Ui32(int32(-20)) {
		v234 = v112
		goto L36
	} else {
		goto L37
	}
L28:
	;
	if base.Ui32(int32(20)) < base.Ui32(v106) {
		goto L26
	} else {
		goto L34
	}
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(-17))))
	v106 = v105
	goto L28
L30:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(-9))))
	v106 = v102
	goto L28
L31:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74+int32(-5)))))
	v106 = v99
	goto L28
L32:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+int32(-3)))))
	v106 = v96
	goto L28
L33:
	;
	v106 = int32(base.Ui32(v89) >> (uint(int32(3)) % 32))
	goto L28
L34:
	;
	v109 = v106
	goto L27
L35:
	;
	if v234 == int32(0) {
		v332 = v109
		goto L25
	} else {
		goto L59
	}
L36:
	;
	goto L35
L37:
	;
	v124 = int32(1)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v109 != v124 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if base.Ui64(v221+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		v234 = v112
		goto L36
	} else {
		goto L58
	}
L39:
	;
	if v125&int32(255) == int32(45) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v129 = v125 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v129&int32(255)) {
		v234 = v112
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v221 = base.I64_extend_i32_u(v129) & int64(255)
	goto L38
L42:
	;
	if base.Ui32(int32(8)) < base.Ui32((v146+int32(-49))&int32(255)) {
		v234 = v112
		goto L36
	} else {
		goto L45
	}
L43:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	v145 = int32(2)
	v146 = v143
	v147 = v74 + int32(1)
	goto L42
L44:
	;
	v145 = v124
	v146 = v125
	v147 = v74
	goto L42
L45:
	;
	v158 = base.I64_extend_i32_u(v146+int32(-48)) & int64(255)
	if base.Ui32(v109) <= base.Ui32(v145) {
		v201 = v158
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if v125&int32(255) != int32(45) {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	v164 = v145
	v166 = v158
	v168 = v147
	goto L48
L48:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	if base.Ui32((v170+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v234 = v112
		goto L36
	} else {
		goto L50
	}
L49:
	;
	v201 = v191
	goto L46
L50:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v166) {
		v234 = v112
		goto L36
	} else {
		goto L51
	}
L51:
	;
	v180 = v166 * int64(10)
	v185 = base.I64_extend_i32_u(v170+int32(-48)) & int64(255)
	if base.Ui64(v185^int64(-1)) < base.Ui64(v180) {
		v234 = v112
		goto L36
	} else {
		goto L52
	}
L52:
	;
	v189 = int32(1)
	v191 = v180 + v185
	v193 = v164 + v189
	if v193 != v109 {
		v164 = v193
		v166 = v191
		v168 = v168 + v189
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L49
L54:
	;
	if v201 < int64(0) {
		v234 = v112
		goto L36
	} else {
		goto L57
	}
L55:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v201) {
		v234 = v112
		goto L36
	} else {
		goto L56
	}
L56:
	;
	v221 = int64(0) - v201
	goto L38
L57:
	;
	v221 = v201
	goto L38
L58:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v10+int32(12)))) = uint32(v221)
	v234 = int32(1)
	goto L36
L59:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v243)>>(uint(int32(4))%32)) & int32(15) {
	case 0:
		goto L61
	default:
		v332 = v109
		goto L25
	case 8:
		goto L60
	}
L60:
	;
	F_decrRefCount(m, l0)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L79
	} else {
		goto L81
	}
L61:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v248&int32(4) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	F_sdsfree(m, v309)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L79
	} else {
		goto L80
	}
L63:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v309 = v307
	goto L62
L64:
	;
	if v248&int32(1) != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v257 = int32(16)
	goto L67
L66:
	;
	v257 = int32(8)
	goto L67
L67:
	;
	v258 = l0 + v257
	if v248&int32(2) == int32(0) {
		v291 = v258
		goto L68
	} else {
		goto L69
	}
L68:
	;
	goto L78
L69:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	v264 = v258 + v263
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	switch v268 & int32(7) {
	case 0:
		goto L75
	case 1:
		goto L74
	case 2:
		goto L73
	case 3:
		goto L72
	case 4:
		goto L71
	default:
		v285 = int32(0)
		goto L70
	}
L70:
	;
	v291 = v264 + int32(1) + v285 + int32(1)
	goto L68
L71:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v264+int32(-16))))
	v285 = v284
	goto L70
L72:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v264+int32(-8))))
	v285 = v281
	goto L70
L73:
	;
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v264+int32(-4)))))
	v285 = v278
	goto L70
L74:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264+int32(-2)))))
	v285 = v275
	goto L70
L75:
	;
	v285 = int32(base.Ui32(v268) >> (uint(int32(3)) % 32))
	goto L70
L76:
	;
	v309 = v291 + v304
	goto L62
L77:
	;
	goto L76
L78:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L77
L79:
	;
	return int32(0)
L80:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v316&int32(-241) | int32(16)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v322
	goto L22
L81:
	;
	v326 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10)+12)))
	v328 = F_createStringObjectFromLongLongWithOptions(m, v326, int32(1))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	v393 = v328
	goto L20
L83:
	;
	v332 = v106
	goto L25
L84:
	;
	if base.Ui32(int32(128)) < base.Ui32(v332+v344+int32(9)) {
		goto L24
	} else {
		goto L87
	}
L85:
	;
	goto L84
L86:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L85
L87:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v351&int32(240) == int32(128) {
		goto L22
	} else {
		goto L88
	}
L88:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	switch v357 & int32(7) {
	case 0:
		goto L94
	case 1:
		goto L93
	case 2:
		goto L92
	case 3:
		goto L91
	case 4:
		goto L90
	default:
		v374 = int32(0)
		goto L89
	}
L89:
	;
	v377 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v74, v374, int32(0), int64(-1))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L79
	} else {
		goto L95
	}
L90:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(-17))))
	v374 = v373
	goto L89
L91:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(-9))))
	v374 = v370
	goto L89
L92:
	;
	v367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74+int32(-5)))))
	v374 = v367
	goto L89
L93:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+int32(-3)))))
	v374 = v364
	goto L89
L94:
	;
	v374 = int32(base.Ui32(v357) >> (uint(int32(3)) % 32))
	goto L89
L95:
	;
	F_decrRefCount(m, l0)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L79
	} else {
		goto L96
	}
L96:
	;
	v393 = v377
	goto L20
L97:
	;
	F_trimStringObjectIfNeeded(m, l0, int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L79
	} else {
		goto L98
	}
L98:
	;
	goto L22
L99:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
