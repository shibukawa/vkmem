package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_getNodeByQuery(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int64
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
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
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v397 int32
	_ = v397
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int64
	_ = v482
	var v483 int64
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v648 int32
	_ = v648
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v667 int32
	_ = v667
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v741 int32
	_ = v741
	v30 = m.G0
	v32 = v30 - int32(2144)
	m.G0 = v32
	v35 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	goto L1
L1:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if int32(-1) < v37 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	m.G0 = v32 + int32(2144)
	return v741
L3:
	;
	if v147 != int32(-1) {
		goto L31
	} else {
		goto L32
	}
L4:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+290)))
	if v130&int32(16) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L5:
	;
	F__serverAssert(m, int32(_a231), int32(_a203), int32(1058))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L26
	} else {
		goto L27
	}
L6:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _consts[114])))
	if v46&int32(4) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+290)))
	if v40&int32(24) == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	if l1 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v741 = v36
	goto L2
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+48))
	if v56 != int32(17) {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L11
L13:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v59&int32(8) == int32(0) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if int32(1) <= v65 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v79 = int32(0)
	v80 = int32(-1)
	goto L17
L16:
	;
	v68 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v68
	v147 = v68
	goto L3
L17:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v71+v79*int32(20))+16))
	if v80 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v117
	v147 = v117
	goto L3
L19:
	;
	v119 = v79 + int32(1)
	if v119 != v65 {
		v79 = v119
		v80 = v117
		goto L17
	} else {
		goto L25
	}
L20:
	;
	if v106 == int32(-1) {
		v117 = v80
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v117 = v106
	goto L19
L22:
	;
	if v106 == v80 {
		v117 = v80
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v112 = int32(0)
	if l1 == v112 {
		v741 = v112
		goto L2
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v741 = v112
	goto L2
L25:
	;
	goto L18
L26:
	;
	return int32(0)
L27:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v147 = v140
	goto L3
L29:
	;
	v135 = int32(0)
	if l1 == v135 {
		v741 = v135
		goto L2
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v741 = v135
	goto L2
L31:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v173+v147<<(uint(int32(2))%32)+int32(52))))
	goto L34
L32:
	;
	v741 = v36
	goto L2
L33:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
	goto L39
L34:
	;
	if v179 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v180 = int32(0)
	if l1 == v180 {
		v741 = v180
		goto L2
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(6)
	v741 = v180
	goto L2
L37:
	;
	v217 = int32(1)
	v219 = F_getCommandFlags(m, l0)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L26
	} else {
		goto L49
	}
L38:
	;
	v194 = F_clusterNodeGetPrimary(m, v36)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L26
	} else {
		goto L43
	}
L39:
	;
	if v185&int32(1) != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+202)))
	if v188&int32(2) != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v192 = int32(0)
	v214 = int32(1)
	v215 = v192
	v216 = v192
	goto L37
L42:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v207 = F_getImportingSlotSource(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L26
	} else {
		goto L47
	}
L43:
	;
	if v179 != v194 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v198 = F_getMigratingSlotDest(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	if v198 == int32(0) {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v202 = int32(1)
	v214 = v202
	v215 = int32(0)
	v216 = v202
	goto L37
L47:
	;
	v209 = int32(0)
	v214 = base.B2i32(v207 == v209)
	v215 = base.B2i32(v207 != v209)
	v216 = int32(0)
	goto L37
L48:
	;
	v237 = int32(0)
	v238 = v215 | v216
	if v238 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	if v219&int64(32) != int64(0) {
		v236 = v217
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+48))
	if v227 != int32(17) {
		v236 = int32(0)
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+8)))
	v236 = int32(base.Ui32(v231&int32(32)) >> (uint(int32(5)) % 32))
	goto L48
L52:
	;
	v624 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v624)+16))
	goto L124
L53:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+48))
	if v243 != int32(17) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v599 = v237
	v600 = v217
	v602 = int32(0)
	goto L52
L55:
	;
	v271 = int32(0)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v293 = v271
	v294 = v271
	v295 = v271
	v297 = v274
	v298 = v271
	v299 = v271
	goto L61
L56:
	;
	v261 = int32(1)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v264 < v261 {
		v599 = v237
		v600 = v261
		v602 = int32(0)
		goto L52
	} else {
		goto L60
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+2092)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+2088)) = v32 + int32(2068)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+2068)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+2080)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v32)+2076)) = v256
	v270 = v32 + int32(2088)
	goto L55
L58:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v246&int32(8) != 0 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v741 = v36
	goto L2
L60:
	;
	v270 = v263
	goto L55
L61:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	v309 = v306 + v295*int32(20)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v309)+8))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v309)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = int64(1099511627776)
	v319 = F_getKeysFromCommand(m, v312, v310, v311, v32+int32(8))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L26
	} else {
		goto L63
	}
L62:
	;
	v587 = int32(0)
	v590 = base.B2i32(v571 != v587)
	v599 = v590
	v600 = base.B2i32(v572 == v587)
	v602 = base.B2i32(v567 != v587) & v590
	goto L52
L63:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v312)+48))
	if v322 != int32(54) {
		v335 = v297
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if v319 < int32(1) {
		v566 = v293
		v567 = v294
		v571 = v298
		v572 = v299
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	v326 = F_getLongLongFromObject(m, v325, v32)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L26
	} else {
		goto L66
	}
L66:
	;
	if v326 != 0 {
		v335 = v297
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v329 = F_selectDb(m, l0, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L26
	} else {
		goto L68
	}
L68:
	;
	if v329 != 0 {
		v335 = v297
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v274)+28))
	v333 = F_selectDb(m, l0, v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L26
	} else {
		goto L70
	}
L70:
	;
	v335 = v331
	goto L64
L71:
	;
	F_getKeysFreeResult(m, v32+int32(8))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L26
	} else {
		goto L121
	}
L72:
	;
	v357 = v293
	v358 = v294
	v362 = v298
	v363 = v299
	v367 = int32(0)
	goto L73
L73:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v321+v367<<(uint(int32(3))%32))))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v310+v373<<(uint(int32(2))%32))))
	if v357 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v566 = v385
	v567 = v386
	v571 = v545
	v572 = v546
	goto L71
L75:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v312)+48))
	if v387 != int32(55) {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	if v214|base.B2i32(v358 != int32(0)) != 0 {
		v385 = v357
		v386 = v358
		goto L75
	} else {
		goto L78
	}
L77:
	;
	v385 = v377
	v386 = v358
	goto L75
L78:
	;
	v381 = F_equalStringObjects(m, v357, v377)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L26
	} else {
		goto L79
	}
L79:
	;
	v385 = v357
	v386 = base.B2i32(v381 == int32(0))
	goto L75
L80:
	;
	if v387 != int32(56) {
		goto L88
	} else {
		goto L89
	}
L81:
	;
	if l1 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_getKeysFreeResult(m, v32+int32(8))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L26
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	goto L82
L84:
	;
	v741 = int32(0)
	goto L2
L85:
	;
	v548 = v367 + int32(1)
	if v548 != v319 {
		v357 = v385
		v358 = v386
		v362 = v545
		v363 = v546
		v367 = v548
		goto L73
	} else {
		goto L120
	}
L86:
	;
	v545 = v362
	v546 = v363 + int32(1)
	goto L85
L87:
	;
	if l1 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L88:
	;
	if v236 != 0 {
		v545 = v362
		v546 = v363
		goto L85
	} else {
		goto L111
	}
L89:
	;
	if v311 < int32(5) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v408 = int32(3)
	v409 = int32(4)
	goto L91
L91:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v310+v408<<(uint(int32(2))%32))))
	v436 = F_objectGetVal(m, v435)
	mBase = m.M
	v437 = int32(_a232)
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
	if v440 != 0 {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	goto L88
L93:
	;
	v487 = v485 + int32(2)
	if v487 < v311 {
		v408 = v485 + int32(1)
		v409 = v487
		goto L91
	} else {
		goto L110
	}
L94:
	;
	if v472-v474 != 0 {
		v485 = v408
		goto L93
	} else {
		goto L106
	}
L95:
	;
	v472 = F_tolower(m, v468)
	mBase = m.M
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469))))
	v474 = F_tolower(m, v473)
	mBase = m.M
	goto L94
L96:
	;
	v442 = v436
	v443 = v437
	v444 = v440
	goto L99
L97:
	;
	v468 = int32(0)
	v469 = v437
	goto L95
L98:
	;
	v468 = v465 & int32(255)
	v469 = v464
	goto L95
L99:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v446 == int32(0) {
		v464 = v443
		v465 = v444
		goto L98
	} else {
		goto L101
	}
L100:
	;
	v464 = v458
	v465 = int32(0)
	goto L98
L101:
	;
	v450 = v444 & int32(255)
	if v450 == v446 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v457 = int32(1)
	v458 = v443 + v457
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442)+1)))
	if v459 != 0 {
		v442 = v442 + v457
		v443 = v458
		v444 = v459
		goto L99
	} else {
		goto L105
	}
L103:
	;
	v452 = F_tolower(m, v450)
	mBase = m.M
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	v454 = F_tolower(m, v453)
	mBase = m.M
	if v452 == v454 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
	v464 = v443
	v465 = v456
	goto L98
L105:
	;
	goto L100
L106:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v310+v409<<(uint(int32(2))%32))))
	v480 = F_getLongLongFromObject(m, v479, v32)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L26
	} else {
		goto L107
	}
L107:
	;
	if v480 != 0 {
		goto L87
	} else {
		goto L108
	}
L108:
	;
	v482 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
	v483 = int64(*(*int32)(unsafe.Add(mBase, uint32(v335)+28)))
	if v482 != v483 {
		goto L87
	} else {
		goto L109
	}
L109:
	;
	v485 = v409
	goto L93
L110:
	;
	goto L92
L111:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v520&int32(8) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v530 = F_lookupKeyReadWithFlags(m, v335, v377, int32(23))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L26
	} else {
		goto L115
	}
L113:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+48))
	if v526 != int32(17) {
		v545 = v362
		v546 = v363
		goto L85
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	if v530 != 0 {
		goto L86
	} else {
		goto L116
	}
L116:
	;
	v545 = v362 + int32(1)
	v546 = v363
	goto L85
L117:
	;
	F_getKeysFreeResult(m, v32+int32(8))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L26
	} else {
		goto L119
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	goto L117
L119:
	;
	v741 = int32(0)
	goto L2
L120:
	;
	goto L74
L121:
	;
	v584 = v295 + int32(1)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v584 < v585 {
		v293 = v566
		v294 = v567
		v295 = v584
		v297 = v335
		v298 = v571
		v299 = v572
		goto L61
	} else {
		goto L122
	}
L122:
	;
	goto L62
L123:
	;
	if v238 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L124:
	;
	if v625 == int32(0) {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	if v236 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v638 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	if v638 != 0 {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	if v631 != 0 {
		goto L123
	} else {
		goto L128
	}
L128:
	;
	v632 = int32(0)
	if l1 == v632 {
		v741 = v632
		goto L2
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(5)
	v741 = v632
	goto L2
L130:
	;
	if v219&int64(1) == int64(0) {
		goto L123
	} else {
		goto L133
	}
L131:
	;
	v639 = int32(0)
	if l1 == v639 {
		v741 = v639
		goto L2
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(5)
	v741 = v639
	goto L2
L133:
	;
	v648 = int32(0)
	if l1 == v648 {
		v741 = v648
		goto L2
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(7)
	v741 = v648
	goto L2
L135:
	;
	if v216&v599 != int32(1) {
		goto L140
	} else {
		goto L141
	}
L136:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v655)+48))
	if v656 != int32(57) {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
	goto L138
L138:
	;
	if v659&int32(1) == int32(0) {
		goto L135
	} else {
		goto L139
	}
L139:
	;
	v741 = v36
	goto L2
L140:
	;
	if v215 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L141:
	;
	if v600 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	if l1 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v667 = int32(0)
	if l1 == v667 {
		v741 = v667
		goto L2
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	v741 = v667
	goto L2
L145:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v677 = F_getMigratingSlotDest(m, v676)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L26
	} else {
		goto L147
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	goto L145
L147:
	;
	v741 = v677
	goto L2
L148:
	;
	if base.I32_wrap_i64(v219)&int32(1) != 0 {
		goto L158
	} else {
		goto L159
	}
L149:
	;
	if v219&int64(8192) != int64(0) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v690 = int32(0)
	if base.B2i32(l1 == v690)|(v602^int32(1)) == v690 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v685&int32(512) == int32(0) {
		goto L148
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	v741 = v690
	goto L2
L154:
	;
	if v602 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v699 = int32(0)
	goto L157
L156:
	;
	v699 = v36
	goto L157
L157:
	;
	v741 = v699
	goto L2
L158:
	;
	if l1 == int32(0) {
		v741 = v179
		goto L2
	} else {
		goto L168
	}
L159:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v706)+48))
	if v707 != int32(17) {
		v714 = int32(0)
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+202)))
	v716 = int32(1)
	if (int32(base.Ui32(v715)>>(uint(v716)%32))|v236)&v716 == int32(0) {
		goto L158
	} else {
		goto L162
	}
L161:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710)+8)))
	v714 = v711 & int32(1)
	goto L160
L162:
	;
	if v714 != 0 {
		goto L158
	} else {
		goto L163
	}
L163:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
	goto L164
L164:
	;
	if v723&int32(2) == int32(0) {
		goto L158
	} else {
		goto L165
	}
L165:
	;
	v728 = F_clusterNodeGetPrimary(m, v36)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L26
	} else {
		goto L166
	}
L166:
	;
	if v728 != v179 {
		goto L158
	} else {
		goto L167
	}
L167:
	;
	v741 = v36
	goto L2
L168:
	;
	if v179 == v36 {
		v741 = v179
		goto L2
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(4)
	v741 = v179
	goto L2
}
func F_isNodeAvailable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v4&int32(8) != 0 {
		v24 = int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		if v8&int32(16) == int32(0) {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+2248))
			v21 = v20
		} else {
			if v8&int32(2) == int32(0) {
				v19 = *(*int64)(unsafe.Add(mBase, _consts[47]))
				v21 = v19
			} else {
				v17 = F_replicationGetReplicaOffset(m)
				mBase = m.M
				v21 = v17
			}
		}
		v24 = base.B2i32(v21 != int64(0))
	}
	return v24
}
func F_markNodeAsFailing(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v4 | int32(8)
	v8 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+2216)) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v10 & int32(-5)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	if v16&int32(2) == v2 {
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+2172))
		if v21 != l0 {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = v16 | int32(8192)
			v27 = *(*int32)(unsafe.Add(mBase, _consts[111]))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[122])))
			*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[122]))) = v28 | int32(1)
		}
	}
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return
	} else {
		v36 = *(*int32)(unsafe.Add(mBase, _consts[111]))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[122])))
		*(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[122]))) = v37 | int32(6)
		return
	}
}
