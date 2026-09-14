package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_evalRegisterNewScript(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
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
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int64
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int64
	_ = v372
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int64
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int64
	_ = v474
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	v11 = m.G0
	v13 = v11 - int32(64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v15 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F__serverAssert(m, int32(_a182), int32(_a550), int32(464))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L9
	} else {
		goto L129
	}
L2:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a560), int32(_a550), int32(333))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L9
	} else {
		goto L128
	}
L3:
	;
	F__serverAssert(m, int32(_a561), int32(_a550), int32(448))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L9
	} else {
		goto L127
	}
L4:
	;
	F__serverAssert(m, int32(_a562), int32(_a550), int32(436))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L9
	} else {
		goto L126
	}
L5:
	;
	F__serverAssert(m, int32(_a563), int32(_a550), int32(410))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L9
	} else {
		goto L125
	}
L6:
	;
	m.G0 = v13 + int32(64)
	return v488
L7:
	;
	v167 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v167
	v173 = F_objectGetVal(m, l1)
	mBase = m.M
	v182 = F_evalExtractShebangFlags(m, v173, v13+int32(40), v13+int32(48), v13+int32(60), v13+int32(44))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L39
	}
L8:
	;
	v17 = F_valkey_calloc(m, int32(41))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v17
	v24 = F_objectGetVal(m, l1)
	mBase = m.M
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v29 = m.G0
	v31 = v29 - int32(112)
	m.G0 = v31
	goto L13
L11:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[252]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v150 = F_dictFind(m, v148, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L33
	}
L12:
	;
	m.G0 = v31 + int32(112)
	goto L11
L13:
	;
	v76 = int32(0)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-1)))))
	switch v80 & int32(7) {
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
		v97 = v76
		goto L24
	}
L24:
	;
	v99 = v31 + int32(20)
	F_SHA1Init(m, v99)
	mBase = m.M
	F_SHA1Update(m, v99, v24, v97)
	mBase = m.M
	F_SHA1Final(m, v31, v99)
	mBase = m.M
	v107 = v76
	goto L30
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-17))))
	v97 = v96
	goto L24
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-9))))
	v97 = v93
	goto L24
L27:
	;
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(-5)))))
	v97 = v90
	goto L24
L28:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-3)))))
	v97 = v87
	goto L24
L29:
	;
	v97 = int32(base.Ui32(v80) >> (uint(int32(3)) % 32))
	goto L24
L30:
	;
	v113 = int32(1)
	v115 = v25 + v107<<(uint(v113)%32)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v107))))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119&int32(15))+uint32(_consts[259]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v115+v113))) = uint8(v124)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v119)>>(uint(int32(4))%32)))+uint32(_consts[259]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v130)
	v133 = v107 + v113
	if v133 != int32(20) {
		v107 = v133
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v136 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+40)) = uint8(v136)
	goto L12
L32:
	;
	goto L31
L33:
	;
	if v150 == int32(0) {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	goto L35
L35:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+24))
	if v155 == int32(0) {
		v488 = int32(0)
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v158 = int32(0)
	v160 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	F_listDelNode(m, v160, v155)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = int32(0)
	v488 = v158
	goto L6
L38:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	if v202 == int32(0) {
		goto L5
	} else {
		goto L49
	}
L39:
	;
	if v182 != int32(-1) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	if l0 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	if v191 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	F_addReplyErrorSds(m, l0, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v196 = int32(-1)
	if v15 != 0 {
		v488 = v196
		goto L6
	} else {
		goto L47
	}
L45:
	;
	F_valkey_free(m, v191)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_valkey_free(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v488 = v196
	goto L6
L49:
	;
	v205 = F_scriptingEngineManagerFind(m, v202)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L9
	} else {
		goto L51
	}
L50:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	F_valkey_free(m, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L9
	} else {
		goto L59
	}
L51:
	;
	if v205 != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	if l0 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	F_valkey_free(m, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L9
	} else {
		goto L56
	}
L54:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v209
	F_addReplyErrorFormat(m, l0, int32(_a564), v13)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v217 = int32(-1)
	if v15 != 0 {
		v488 = v217
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_valkey_free(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v488 = v217
	goto L6
L59:
	;
	v226 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v226
	v230 = F_objectGetVal(m, l1)
	mBase = m.M
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	v234 = F_objectGetVal(m, l1)
	mBase = m.M
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234+int32(-1)))))
	switch v237 & int32(7) {
	case 0:
		goto L65
	case 1:
		goto L64
	case 2:
		goto L63
	case 3:
		goto L62
	case 4:
		goto L61
	default:
		v254 = v226
		goto L60
	}
L60:
	;
	v255 = int32(0)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	v263 = F_scriptingEngineCallCompileCode(m, v205, v255, v230+v231, v254-v256, v255, v13+int32(32), v13+int32(36))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L9
	} else {
		goto L67
	}
L61:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v234+int32(-17))))
	v254 = v253
	goto L60
L62:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v234+int32(-9))))
	v254 = v250
	goto L60
L63:
	;
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v234+int32(-5)))))
	v254 = v247
	goto L60
L64:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234+int32(-3)))))
	v254 = v244
	goto L60
L65:
	;
	v254 = int32(base.Ui32(v237) >> (uint(int32(3)) % 32))
	goto L60
L66:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	if v287 != int32(1) {
		goto L3
	} else {
		goto L76
	}
L67:
	;
	if v263 != 0 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	if v265 == int32(0) {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	if l0 == int32(0) {
		v278 = v265
		goto L70
	} else {
		goto L71
	}
L70:
	;
	F_decrRefCount(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L9
	} else {
		goto L73
	}
L71:
	;
	v270 = F_objectGetVal(m, v265)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v270
	F_addReplyErrorFormat(m, l0, int32(_a57), v13+int32(16))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	v278 = v277
	goto L70
L73:
	;
	v281 = int32(-1)
	if v15 != 0 {
		v488 = v281
		goto L6
	} else {
		goto L74
	}
L74:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_valkey_free(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L9
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v488 = v281
	goto L6
L76:
	;
	v291 = F_valkey_calloc(m, int32(32))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = v293
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v291)+16)) = v296
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v299 = F_sdsnew(m, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	if v15 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = l1
	v425 = int32(0)
	v427 = *(*int32)(unsafe.Add(mBase, _consts[252]))
	v428 = F_dictAdd(m, v427, v299, v291)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L9
	} else {
		goto L107
	}
L80:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+20))
	if base.Ui32(v305) < base.Ui32(int32(500)) {
		v398 = v304
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v406 = F_sdsdup(m, v299)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L9
	} else {
		goto L105
	}
L82:
	;
	v310 = v304
	goto L83
L83:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _consts[252]))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+8))
	v322 = F_dictUnlink(m, v319, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L9
	} else {
		goto L85
	}
L84:
	;
	v398 = v392
	goto L81
L85:
	;
	if v322 == int32(0) {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	goto L87
L87:
	;
	v333 = v321 + int32(-1)
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	v336 = v334 & int32(7)
	switch v336 {
	case 0:
		goto L94
	case 1:
		v342 = int32(4)
		goto L89
	case 2:
		goto L93
	case 3:
		goto L92
	case 4:
		goto L91
	default:
		goto L90
	}
L88:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v326)+8))
	v368 = F_getStringObjectSdsUsedMemory(m, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L9
	} else {
		goto L101
	}
L89:
	;
	switch v336 {
	case 0:
		goto L100
	case 1:
		goto L99
	case 2:
		goto L98
	case 3:
		goto L97
	case 4:
		goto L96
	default:
		v362 = int32(0)
		goto L95
	}
L90:
	;
	v342 = int32(1)
	goto L89
L91:
	;
	v342 = int32(18)
	goto L89
L92:
	;
	v342 = int32(10)
	goto L89
L93:
	;
	v342 = int32(6)
	goto L89
L94:
	;
	v337 = F_zmalloc_usable_size(m, v333)
	mBase = m.M
	v366 = v337
	goto L88
L95:
	;
	v366 = v342 + v362
	goto L88
L96:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v321+int32(-9))))
	v362 = v361
	goto L95
L97:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v321+int32(-5))))
	v366 = v342 + v357
	goto L88
L98:
	;
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v321+int32(-3)))))
	v366 = v342 + v353
	goto L88
L99:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321+int32(-2)))))
	v366 = v342 + v349
	goto L88
L100:
	;
	v366 = v342 + int32(base.Ui32(v334)>>(uint(int32(3))%32))
	goto L88
L101:
	;
	v370 = int32(0)
	v372 = *(*int64)(unsafe.Add(mBase, _consts[254]))
	*(*int64)(unsafe.Add(mBase, _consts[254])) = v372 - base.I64_extend_i32_u(v368+v366)
	v378 = *(*int32)(unsafe.Add(mBase, _consts[252]))
	F_dictFreeUnlinkedEntry(m, v378, v322)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L9
	} else {
		goto L102
	}
L102:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	F_listDelNode(m, v382, v320)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L9
	} else {
		goto L103
	}
L103:
	;
	v385 = int32(_a69)
	v387 = *(*int64)(unsafe.Add(mBase, _consts[260]))
	*(*int64)(unsafe.Add(mBase, _consts[260])) = v387 + int64(1)
	v392 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+20))
	if base.Ui32(int32(499)) < base.Ui32(v393) {
		v310 = v392
		goto L83
	} else {
		goto L104
	}
L104:
	;
	goto L84
L105:
	;
	v408 = F_listAddNodeTail(m, v398, v406)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L9
	} else {
		goto L106
	}
L106:
	;
	v411 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v291)+24)) = v412
	goto L79
L107:
	;
	if v428 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v436 = v299 + int32(-1)
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
	v439 = v437 & int32(7)
	switch v439 {
	case 0:
		goto L115
	case 1:
		v445 = int32(4)
		goto L110
	case 2:
		goto L114
	case 3:
		goto L113
	case 4:
		goto L112
	default:
		goto L111
	}
L109:
	;
	v470 = F_getStringObjectSdsUsedMemory(m, l1)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L9
	} else {
		goto L122
	}
L110:
	;
	switch v439 {
	case 0:
		goto L121
	case 1:
		goto L120
	case 2:
		goto L119
	case 3:
		goto L118
	case 4:
		goto L117
	default:
		v465 = int32(0)
		goto L116
	}
L111:
	;
	v445 = int32(1)
	goto L110
L112:
	;
	v445 = int32(18)
	goto L110
L113:
	;
	v445 = int32(10)
	goto L110
L114:
	;
	v445 = int32(6)
	goto L110
L115:
	;
	v440 = F_zmalloc_usable_size(m, v436)
	mBase = m.M
	v469 = v440
	goto L109
L116:
	;
	v469 = v445 + v465
	goto L109
L117:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v299+int32(-9))))
	v465 = v464
	goto L116
L118:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v299+int32(-5))))
	v469 = v445 + v460
	goto L109
L119:
	;
	v456 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v299+int32(-3)))))
	v469 = v445 + v456
	goto L109
L120:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299+int32(-2)))))
	v469 = v445 + v452
	goto L109
L121:
	;
	v469 = v445 + int32(base.Ui32(v437)>>(uint(int32(3))%32))
	goto L109
L122:
	;
	v472 = int32(0)
	v474 = *(*int64)(unsafe.Add(mBase, _consts[254]))
	*(*int64)(unsafe.Add(mBase, _consts[254])) = v474 + base.I64_extend_i32_u(v470+v469)
	F_incrRefCount(m, l1)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L9
	} else {
		goto L123
	}
L123:
	;
	F_valkey_free(m, v263)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L9
	} else {
		goto L124
	}
L124:
	;
	v488 = v425
	goto L6
L125:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_scriptAllowsCrossSlot(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v3 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
		return v12 & int32(256)
	} else {
		F__serverAssert(m, int32(_a1090), int32(_a1073), int32(341))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_scriptAllowsOOM(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v3 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
		return v12 & int32(64)
	} else {
		F__serverAssert(m, int32(_a1090), int32(_a1073), int32(321))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_scriptInterrupt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v72 int64
	_ = v72
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v87 int64
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v9&int32(8) != 0 {
		F_processEventsWhileBlocked(m)
		mBase = m.M
		v97 = m.ExcPending
		if v97 != 0 {
			return int32(0)
		} else {
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
			if v100&int32(16) != 0 {
				v103 = int32(1)
			} else {
				v103 = int32(2)
			}
			v104 = v103
			m.G0 = v7 + int32(16)
			return v104
		}
	} else {
		v12 = int32(2)
		v14 = *(*int64)(unsafe.Add(mBase, _consts[348]))
		if v14 == int64(0) {
			v104 = v12
			m.G0 = v7 + int32(16)
			return v104
		} else {
			v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			v19 = *(*int32)(unsafe.Add(mBase, _consts[271]))
			v20 = m.T0[v19].(func(*base.Module) int64)(m)
			mBase = m.M
			v23 = base.I64_div_u_s(v20-v17, int64(1000))
			v25 = *(*int64)(unsafe.Add(mBase, _consts[348]))
			if v23 < v25 {
				v104 = v12
				m.G0 = v7 + int32(16)
				return v104
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _consts[15]))
				if int32(3) < v28 {
					v49 = *(*int32)(unsafe.Add(mBase, _consts[163]))
					if l0 != v49 {
						F__serverAssert(m, int32(_a1072), int32(_a1073), int32(56))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+16)))
						if v51&int32(8) != 0 {
							F__serverAssert(m, int32(_a1074), int32(_a1073), int32(57))
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v54 | int32(8)
							v58 = int32(0)
							v63 = *(*int32)(unsafe.Add(mBase, _consts[487]))
							*(*int32)(unsafe.Add(mBase, _consts[487])) = v63 + int32(1)
							if v63 != 0 {
							} else {
								v67 = int32(0)
								v68 = F_ustime(m)
								mBase = m.M
								*(*int64)(unsafe.Add(mBase, _consts[277])) = v68
								v72 = base.I64_div_s(v68, int64(1000))
								*(*int64)(unsafe.Add(mBase, _consts[32])) = v72
								v76 = base.I64_div_s(v68, int64(1000000))
								*(*int64)(unsafe.Add(mBase, _consts[37])) = v76
								v79 = *(*int32)(unsafe.Add(mBase, _consts[167]))
								F_lrulfu_updateClockAndPolicy(m, v72, int32(base.Ui32(v79&int32(2))>>(uint(int32(1))%32)))
								mBase = m.M
								v87 = *(*int64)(unsafe.Add(mBase, _consts[32]))
								*(*int64)(unsafe.Add(mBase, _consts[488])) = v87
							}
							v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							F_protectClient(m, v91)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								F_processEventsWhileBlocked(m)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
									if v100&int32(16) != 0 {
										v103 = int32(1)
									} else {
										v103 = int32(2)
									}
									v104 = v103
									m.G0 = v7 + int32(16)
									return v104
								}
							}
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v32
					*(*int64)(unsafe.Add(mBase, uint32(v7))) = v23
					if v31&int32(128) != 0 {
						v39 = int32(_a1075)
					} else {
						v39 = int32(_a1076)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v39
					F__serverLog(m, int32(3), int32(_a1077), v7)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, _consts[163]))
						if l0 != v49 {
							F__serverAssert(m, int32(_a1072), int32(_a1073), int32(56))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+16)))
							if v51&int32(8) != 0 {
								F__serverAssert(m, int32(_a1074), int32(_a1073), int32(57))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v54 | int32(8)
								v58 = int32(0)
								v63 = *(*int32)(unsafe.Add(mBase, _consts[487]))
								*(*int32)(unsafe.Add(mBase, _consts[487])) = v63 + int32(1)
								if v63 != 0 {
								} else {
									v67 = int32(0)
									v68 = F_ustime(m)
									mBase = m.M
									*(*int64)(unsafe.Add(mBase, _consts[277])) = v68
									v72 = base.I64_div_s(v68, int64(1000))
									*(*int64)(unsafe.Add(mBase, _consts[32])) = v72
									v76 = base.I64_div_s(v68, int64(1000000))
									*(*int64)(unsafe.Add(mBase, _consts[37])) = v76
									v79 = *(*int32)(unsafe.Add(mBase, _consts[167]))
									F_lrulfu_updateClockAndPolicy(m, v72, int32(base.Ui32(v79&int32(2))>>(uint(int32(1))%32)))
									mBase = m.M
									v87 = *(*int64)(unsafe.Add(mBase, _consts[32]))
									*(*int64)(unsafe.Add(mBase, _consts[488])) = v87
								}
								v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								F_protectClient(m, v91)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									F_processEventsWhileBlocked(m)
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
										if v100&int32(16) != 0 {
											v103 = int32(1)
										} else {
											v103 = int32(2)
										}
										v104 = v103
										m.G0 = v7 + int32(16)
										return v104
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
func F_scriptIsReadOnly(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v3 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
		return v12 & int32(32)
	} else {
		F__serverAssert(m, int32(_a1090), int32(_a1073), int32(326))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_scriptKill(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v7 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v7 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		if v13 != int64(-1) {
			v17 = int32(1)
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+200)))
			if v18&v17 != 0 {
				v25 = v17
				v28 = v25
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
				if v21 != 0 {
					v23 = F_isImportSlotMigrationJob(m, v21)
					mBase = m.M
					v25 = v23
					v28 = v25
				} else {
					v28 = int32(0)
				}
			}
		} else {
			v28 = int32(1)
		}
		if v28 == int32(0) {
			v34 = int32(0)
			v35 = *(*int32)(unsafe.Add(mBase, _consts[163]))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
			if v36&int32(1) == v34 {
				v45 = v36 & int32(128)
				if l1 == int32(0) {
					if l1 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v36 | int32(16)
						v62 = *(*int32)(unsafe.Add(mBase, _consts[77]))
						F_addReply(m, l0, v62)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							return
						}
					} else {
						if v45 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v36 | int32(16)
							v62 = *(*int32)(unsafe.Add(mBase, _consts[77]))
							F_addReply(m, l0, v62)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								return
							}
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, _consts[230]))
							F_addReplyErrorObject(m, l0, v55)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					if v45 != 0 {
						if l1 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v36 | int32(16)
							v62 = *(*int32)(unsafe.Add(mBase, _consts[77]))
							F_addReply(m, l0, v62)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								return
							}
						} else {
							if v45 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v36 | int32(16)
								v62 = *(*int32)(unsafe.Add(mBase, _consts[77]))
								F_addReply(m, l0, v62)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									return
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, _consts[230]))
								F_addReplyErrorObject(m, l0, v55)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, _consts[229]))
						F_addReplyErrorObject(m, l0, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				F_addReplyError(m, l0, int32(_a1087))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_addReplyError(m, l0, int32(_a1088))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_addReplyError(m, l0, int32(_a1089))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	}
}
func F_scriptPrepareForRun(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v224 int64
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int64
	_ = v250
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v301 int32
	_ = v301
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a1078), int32(_a1073), int32(133))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L29
	} else {
		goto L85
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+204))
	v21 = v19 & int32(32768)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	if v23 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	m.G0 = v15 + int32(32)
	return v288
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l3
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l2)+292))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v243
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v245
	v247 = int32(0)
	v249 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	v250 = m.T0[v249].(func(*base.Module) int64)(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(12884901888)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v250
	if l5 != 0 {
		goto L76
	} else {
		goto L77
	}
L5:
	;
	if v31 != 0 {
		v239 = int32(0)
		goto L4
	} else {
		goto L73
	}
L6:
	;
	if l4&int64(8) == int64(0) {
		goto L26
	} else {
		goto L27
	}
L7:
	;
	v57 = int32(0)
	v59 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	if v59 != int64(-1) {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	if v27 == int32(14) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[610]))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	if v33 != int64(-1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if base.B2i32(l4&int64(16) == int64(0)) == int32(0) {
		goto L5
	} else {
		goto L17
	}
L11:
	;
	v37 = int32(1)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+200)))
	if v38&v37 != 0 {
		v45 = v37
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v48 = int32(1)
	goto L10
L13:
	;
	v48 = v45
	goto L10
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+216))
	if v41 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = F_isImportSlotMigrationJob(m, v41)
	mBase = m.M
	v45 = v43
	goto L13
L16:
	;
	v48 = int32(0)
	goto L10
L17:
	;
	v80 = v48
	v81 = base.B2i32(v31 == int32(0))
	goto L6
L18:
	;
	if l4&int64(16) != int64(0) {
		v239 = v57
		goto L4
	} else {
		goto L25
	}
L19:
	;
	v63 = int32(1)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+200)))
	if v64&v63 != 0 {
		v71 = v63
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v74 = int32(1)
	goto L18
L21:
	;
	v74 = v71
	goto L18
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+216))
	if v67 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v69 = F_isImportSlotMigrationJob(m, v67)
	mBase = m.M
	v71 = v69
	goto L21
L24:
	;
	v74 = int32(0)
	goto L18
L25:
	;
	v80 = v74
	v81 = v57
	goto L6
L26:
	;
	if base.I32_wrap_i64(int64(base.Ui64(l4&int64(4))>>(uint(int64(2))%64)))|(v81^int32(1)) != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v87 == int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	F_addReplyError(m, l2, int32(_a1085))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	v288 = int32(-1)
	goto L3
L31:
	;
	if base.I32_wrap_i64(l4)&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	F_addReplyError(m, l2, int32(_a1084))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v288 = int32(-1)
	goto L3
L34:
	;
	v214 = int32(1)
	if l4&int64(3) != int64(0) {
		v239 = v214
		goto L4
	} else {
		goto L68
	}
L35:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	if v112 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v123 = int32(0)
	v125 = *(*int32)(unsafe.Add(mBase, _consts[605]))
	if v125 == v123 {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _consts[231]))
	if v116 == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	if v80 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	F_addReplyError(m, l2, int32(_a1083))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L29
	} else {
		goto L40
	}
L40:
	;
	v288 = int32(-1)
	goto L3
L41:
	;
	if l5 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L42:
	;
	if v157 == int32(0) {
		goto L41
	} else {
		goto L52
	}
L43:
	;
	goto L42
L44:
	;
	v138 = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v140 == v138 {
		v157 = v138
		goto L43
	} else {
		goto L48
	}
L45:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	if v129 < int32(1) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v134 == int32(-1) {
		v157 = int32(2)
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v143 = int32(0)
	v145 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v145 == int32(-1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v157 = int32(1)
	goto L43
L50:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v149 != int32(-1) {
		v157 = v143
		goto L43
	} else {
		goto L51
	}
L51:
	;
	v152 = int32(0)
	v154 = *(*int32)(unsafe.Add(mBase, _consts[606]))
	*(*int32)(unsafe.Add(mBase, _consts[607])) = v154
	goto L49
L52:
	;
	if v80 != 0 {
		goto L41
	} else {
		goto L53
	}
L53:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _consts[237]))
	if v163 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v164 = int32(_a513)
	goto L56
L55:
	;
	v164 = int32(_a184)
	goto L56
L56:
	;
	if v157 != int32(2) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _consts[607]))
	v174 = F___strerror_l(m, v173, v173)
	mBase = m.M
	goto L60
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v164
	F_addReplyErrorFormat(m, l2, int32(_a1082), v15)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L29
	} else {
		goto L59
	}
L59:
	;
	v288 = int32(-1)
	goto L3
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v164
	F_addReplyErrorFormat(m, l2, int32(_a1081), v15+int32(16))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L29
	} else {
		goto L61
	}
L61:
	;
	v288 = int32(-1)
	goto L3
L62:
	;
	v189 = int32(0)
	v190 = int32(_a69)
	v191 = *(*int32)(unsafe.Add(mBase, _consts[187]))
	v195 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	v200 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	v205 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	goto L65
L63:
	;
	F_addReplyError(m, l2, int32(_a1080))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L29
	} else {
		goto L64
	}
L64:
	;
	v288 = int32(-1)
	goto L3
L65:
	;
	if base.B2i32(v191 == v189)|base.B2i32(v195 == v189)|base.B2i32(v200 != v189)|base.B2i32(v195 <= v205) != 0 {
		goto L34
	} else {
		goto L66
	}
L66:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _consts[609]))
	F_addReplyErrorObject(m, l2, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L29
	} else {
		goto L67
	}
L67:
	;
	v288 = int32(-1)
	goto L3
L68:
	;
	if v21 != 0 {
		v239 = v214
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _consts[604]))
	if v220 == int32(0) {
		v239 = v214
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v224 = *(*int64)(unsafe.Add(mBase, _consts[265]))
	if v224 == int64(0) {
		v239 = v214
		goto L4
	} else {
		goto L71
	}
L71:
	;
	F_addReplyError(m, l2, int32(_a1079))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L29
	} else {
		goto L72
	}
L72:
	;
	v288 = int32(-1)
	goto L3
L73:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _consts[611]))
	F_addReplyErrorObject(m, l2, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L29
	} else {
		goto L74
	}
L74:
	;
	v288 = int32(-1)
	goto L3
L75:
	;
	if v21 != 0 {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v261 = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v261
	v264 = v261
	goto L75
L77:
	;
	if base.B2i32(l4&int64(1) == int64(0))|(v239^int32(1)) != 0 {
		v264 = v247
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	if base.B2i32(l4&int64(32) == int64(0))&v239 != 0 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v273 = v264 | int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v273
	v275 = v273
	goto L79
L81:
	;
	if base.B2i32(l4&int64(2) == int64(0))|(v239^int32(1)) != 0 {
		v275 = v264
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v284 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[163])) = l0
	v288 = v284
	goto L3
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v275 | int32(256)
	goto L83
L85:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_scriptResetRun(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v4 == v2 {
		F__serverAssert(m, int32(_a1086), int32(_a1073), int32(249))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+16)))
		if v7&int32(8) == int32(0) {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v42)+96)) = v43
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(-1)
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+200))
			*(*int32)(unsafe.Add(mBase, uint32(v42)+200)) = v47 | int32(2097152)
			v51 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[163])) = v51
			return
		} else {
			if l0 != v4 {
				F__serverAssert(m, int32(_a1072), int32(_a1073), int32(47))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v13 & int32(-9)
				v17 = int32(0)
				v20 = *(*int32)(unsafe.Add(mBase, _consts[487]))
				v22 = v20 + int32(-1)
				*(*int32)(unsafe.Add(mBase, _consts[487])) = v22
				if v22 != 0 {
				} else {
					*(*int64)(unsafe.Add(mBase, _consts[488])) = int64(0)
				}
				v28 = *(*int32)(unsafe.Add(mBase, _consts[64]))
				if v28 == int32(0) {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_unprotectClient(m, v38)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v42)+96)) = v43
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(-1)
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+200))
						*(*int32)(unsafe.Add(mBase, uint32(v42)+200)) = v47 | int32(2097152)
						v51 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[163])) = v51
						return
					}
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _consts[133]))
					if v32 == int32(0) {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						F_unprotectClient(m, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v42)+96)) = v43
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(-1)
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+200))
							*(*int32)(unsafe.Add(mBase, uint32(v42)+200)) = v47 | int32(2097152)
							v51 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[163])) = v51
							return
						}
					} else {
						F_queueClientForReprocessing(m, v32)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							F_unprotectClient(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v42)+96)) = v43
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(-1)
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+200))
								*(*int32)(unsafe.Add(mBase, uint32(v42)+200)) = v47 | int32(2097152)
								v51 = int32(0)
								*(*int32)(unsafe.Add(mBase, _consts[163])) = v51
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_scriptSetOriginalClientSlot(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v4 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+292)) = l0
		return
	} else {
		F__serverAssert(m, int32(_a1090), int32(_a1073), int32(356))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
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
func F_scriptSetWriteDirtyFlag(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v3 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = v10 | int32(1)
		return
	} else {
		F__serverAssert(m, int32(_a1090), int32(_a1073), int32(336))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
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
