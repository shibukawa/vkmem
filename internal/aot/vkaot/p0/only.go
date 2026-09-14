package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_rewriteAppendOnlyFileRio(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v105 int32
	_ = v105
	var v121 int32
	_ = v121
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
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
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int64
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int64
	_ = v311
	var v314 int64
	_ = v314
	var v317 int64
	_ = v317
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int64
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int64
	_ = v414
	var v415 int64
	_ = v415
	var v416 int64
	_ = v416
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v419 int64
	_ = v419
	var v420 int64
	_ = v420
	var v422 int64
	_ = v422
	var v424 int64
	_ = v424
	var v426 int64
	_ = v426
	var v428 int64
	_ = v428
	var v430 int64
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int64
	_ = v508
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v521 int64
	_ = v521
	var v522 int64
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int64
	_ = v542
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int64
	_ = v560
	var v563 int64
	_ = v563
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int64
	_ = v593
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v620 int64
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v632 int64
	_ = v632
	var v639 int32
	_ = v639
	var v640 int64
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v658 int64
	_ = v658
	var v662 int64
	_ = v662
	var v664 int64
	_ = v664
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	if v19 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(32)
	return v698
L2:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	goto L36
L3:
	;
	v23 = F_genAofTimestampAnnotationIfNeeded(m, int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	F_sdsfree(m, v23)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L13
	} else {
		goto L34
	}
L5:
	;
	F_sdsfree(m, v23)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L13
	} else {
		goto L33
	}
L6:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v86&int32(6) == int32(0) {
		goto L4
	} else {
		goto L32
	}
L7:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v47&int32(6) != 0 {
		goto L5
	} else {
		goto L15
	}
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-17))))
	v46 = v45
	goto L7
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-9))))
	v46 = v42
	goto L7
L10:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(-5)))))
	v46 = v39
	goto L7
L11:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-3)))))
	v46 = v36
	goto L7
L12:
	;
	v46 = int32(base.Ui32(v29) >> (uint(int32(3)) % 32))
	goto L7
L13:
	;
	return int32(0)
L14:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-1)))))
	switch v29 & int32(7) {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	case 4:
		goto L8
	default:
		goto L6
	}
L15:
	;
	if v46 == int32(0) {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v56 = v46
	v57 = v23
	goto L17
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v65) < base.Ui32(v56) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v67 = v65
	goto L21
L20:
	;
	v67 = v56
	goto L21
L21:
	;
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v68 = v67
	goto L24
L23:
	;
	v68 = v56
	goto L24
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v69 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v75 = m.T0[v74].(func(*base.Module, int32, int32, int32) int32)(m, l0, v57, v68)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L13
	} else {
		goto L29
	}
L26:
	;
	m.T0[v69].(func(*base.Module, int32, int32, int32))(m, l0, v57, v68)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v81 + v68
	v85 = v56 - v68
	if v85 != 0 {
		v56 = v85
		v57 = v57 + v68
		goto L17
	} else {
		goto L31
	}
L29:
	;
	if v75 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v77 | int64(2)
	goto L5
L31:
	;
	goto L4
L32:
	;
	goto L5
L33:
	;
	v698 = int32(-1)
	goto L1
L34:
	;
	goto L2
L35:
	;
	F_dictReleaseIterator(m, v138)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L13
	} else {
		goto L135
	}
L36:
	;
	v138 = F_dictGetIterator(m, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	v147 = v138 + int32(20)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	if v148 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	if v243 == int32(0) {
		goto L35
	} else {
		goto L64
	}
L39:
	;
	v154 = v147
	v155 = v151
	goto L42
L40:
	;
	v151 = int32(1)
	goto L39
L41:
	;
	v151 = int32(0)
	goto L39
L42:
	;
	switch v155 {
	case 0:
		goto L47
	default:
		goto L46
	}
L44:
	;
	v155 = int32(0)
	goto L42
L45:
	;
	goto L38
L46:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+16)) = v235
	if v235 == int32(0) {
		goto L44
	} else {
		goto L63
	}
L47:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v159 != int32(-1) {
		v198 = v159
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v199 = int32(1)
	v200 = v198 + v199
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v200
	v202 = int32(0)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205+v206+int32(26)))))
	if v210 == int32(255) {
		goto L57
	} else {
		goto L58
	}
L49:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	if v163 != 0 {
		v198 = int32(-1)
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	if v165 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+20))
	if v192 != int32(-1) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v172 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v164)+16)))
	v173 = int64(*(*int8)(unsafe.Add(mBase, uint32(v164)+27)))
	v174 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+8)))
	v175 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v164)+12)))
	v176 = int64(*(*int8)(unsafe.Add(mBase, uint32(v164)+26)))
	v177 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+4)))
	v178 = F_wangHash64(m, v177)
	mBase = m.M
	v180 = F_wangHash64(m, v176+v178)
	mBase = m.M
	v182 = F_wangHash64(m, v175+v180)
	mBase = m.M
	v184 = F_wangHash64(m, v174+v182)
	mBase = m.M
	v186 = F_wangHash64(m, v173+v184)
	mBase = m.M
	v188 = F_wangHash64(m, v172+v186)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v138)+24)) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v191 = v190
	goto L51
L53:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+24)))
	v170 = v168 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v164)+24)) = uint16(v170)
	v191 = v164
	goto L51
L54:
	;
	v198 = v192 + int32(-1)
	goto L48
L55:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v198 = v195
	goto L48
L56:
	;
	v225 = int32(2)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v205+v223<<(uint(v225)%32)+int32(4))))
	v154 = v230 + v224<<(uint(v225)%32)
	v155 = int32(1)
	goto L42
L57:
	;
	v214 = v202
	goto L59
L58:
	;
	v214 = v199 << (uint(v210) % 32)
	goto L59
L59:
	;
	if v200 < v214 {
		v223 = v206
		v224 = v200
		goto L56
	} else {
		goto L60
	}
L60:
	;
	if v206 != 0 {
		v243 = v202
		goto L45
	} else {
		goto L61
	}
L61:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v205)+20))
	if v216 == int32(-1) {
		v243 = v202
		goto L45
	} else {
		goto L62
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v138)+4)) = int64(4294967296)
	v223 = int32(1)
	v224 = int32(0)
	goto L56
L63:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v239
	v243 = v235
	goto L45
L64:
	;
	v254 = v243
	goto L65
L65:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v254)+8))
	goto L67
L66:
	;
	goto L35
L67:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v265&int32(6) != 0 {
		goto L35
	} else {
		goto L68
	}
L68:
	;
	v274 = int32(4)
	v275 = int32(_a87)
	goto L69
L69:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v283) < base.Ui32(v274) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v304 = int32(24)
	v307 = int32(0)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, _consts[19])))
	*(*uint8)(unsafe.Add(mBase, uint32(v16+v304))) = uint8(v308)
	v311 = *(*int64)(unsafe.Add(mBase, _consts[20]))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(16)))) = v311
	v314 = *(*int64)(unsafe.Add(mBase, _consts[21]))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v314
	v317 = *(*int64)(unsafe.Add(mBase, _consts[22]))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v317
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v319&int32(6) != 0 {
		goto L35
	} else {
		goto L84
	}
L71:
	;
	v285 = v283
	goto L73
L72:
	;
	v285 = v274
	goto L73
L73:
	;
	if v283 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v286 = v285
	goto L76
L75:
	;
	v286 = v274
	goto L76
L76:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v287 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v293 = m.T0[v292].(func(*base.Module, int32, int32, int32) int32)(m, l0, v275, v286)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L13
	} else {
		goto L81
	}
L78:
	;
	m.T0[v287].(func(*base.Module, int32, int32, int32))(m, l0, v275, v286)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L13
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v299 + v286
	v303 = v274 - v286
	if v303 != 0 {
		v274 = v303
		v275 = v275 + v286
		goto L69
	} else {
		goto L83
	}
L81:
	;
	if v293 != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v295 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v295 | int64(2)
	goto L35
L83:
	;
	goto L70
L84:
	;
	v326 = v304
	v327 = v16
	goto L85
L85:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v335) < base.Ui32(v326) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357+int32(-1)))))
	switch v360 & int32(7) {
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
		v377 = int32(0)
		goto L100
	}
L87:
	;
	v337 = v335
	goto L89
L88:
	;
	v337 = v326
	goto L89
L89:
	;
	if v335 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v338 = v337
	goto L92
L91:
	;
	v338 = v326
	goto L92
L92:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v339 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v345 = m.T0[v344].(func(*base.Module, int32, int32, int32) int32)(m, l0, v327, v338)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L13
	} else {
		goto L97
	}
L94:
	;
	m.T0[v339].(func(*base.Module, int32, int32, int32))(m, l0, v327, v338)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L13
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v351 + v338
	v355 = v326 - v338
	if v355 != 0 {
		v326 = v355
		v327 = v327 + v338
		goto L85
	} else {
		goto L99
	}
L97:
	;
	if v345 != 0 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v347 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v347 | int64(2)
	goto L35
L99:
	;
	goto L86
L100:
	;
	v378 = F_rioWriteBulkString(m, l0, v357, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L13
	} else {
		goto L106
	}
L101:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v357+int32(-17))))
	v377 = v376
	goto L100
L102:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v357+int32(-9))))
	v377 = v373
	goto L100
L103:
	;
	v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357+int32(-5)))))
	v377 = v370
	goto L100
L104:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357+int32(-3)))))
	v377 = v367
	goto L100
L105:
	;
	v377 = int32(base.Ui32(v360) >> (uint(int32(3)) % 32))
	goto L100
L106:
	;
	if v378 == int32(0) {
		goto L35
	} else {
		goto L107
	}
L107:
	;
	v389 = v138 + int32(20)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	if v390 != 0 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	if v485 != 0 {
		v254 = v485
		goto L65
	} else {
		goto L134
	}
L109:
	;
	v396 = v389
	v397 = v393
	goto L112
L110:
	;
	v393 = int32(1)
	goto L109
L111:
	;
	v393 = int32(0)
	goto L109
L112:
	;
	switch v397 {
	case 0:
		goto L117
	default:
		goto L116
	}
L114:
	;
	v397 = int32(0)
	goto L112
L115:
	;
	goto L108
L116:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+16)) = v477
	if v477 == int32(0) {
		goto L114
	} else {
		goto L133
	}
L117:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v401 != int32(-1) {
		v440 = v401
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v441 = int32(1)
	v442 = v440 + v441
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v442
	v444 = int32(0)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447+v448+int32(26)))))
	if v452 == int32(255) {
		goto L127
	} else {
		goto L128
	}
L119:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	if v405 != 0 {
		v440 = int32(-1)
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	if v407 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+20))
	if v434 != int32(-1) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	v414 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v406)+16)))
	v415 = int64(*(*int8)(unsafe.Add(mBase, uint32(v406)+27)))
	v416 = int64(*(*int32)(unsafe.Add(mBase, uint32(v406)+8)))
	v417 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v406)+12)))
	v418 = int64(*(*int8)(unsafe.Add(mBase, uint32(v406)+26)))
	v419 = int64(*(*int32)(unsafe.Add(mBase, uint32(v406)+4)))
	v420 = F_wangHash64(m, v419)
	mBase = m.M
	v422 = F_wangHash64(m, v418+v420)
	mBase = m.M
	v424 = F_wangHash64(m, v417+v422)
	mBase = m.M
	v426 = F_wangHash64(m, v416+v424)
	mBase = m.M
	v428 = F_wangHash64(m, v415+v426)
	mBase = m.M
	v430 = F_wangHash64(m, v414+v428)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v138)+24)) = v430
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v433 = v432
	goto L121
L123:
	;
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v406)+24)))
	v412 = v410 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v406)+24)) = uint16(v412)
	v433 = v406
	goto L121
L124:
	;
	v440 = v434 + int32(-1)
	goto L118
L125:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v440 = v437
	goto L118
L126:
	;
	v467 = int32(2)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v447+v465<<(uint(v467)%32)+int32(4))))
	v396 = v472 + v466<<(uint(v467)%32)
	v397 = int32(1)
	goto L112
L127:
	;
	v456 = v444
	goto L129
L128:
	;
	v456 = v441 << (uint(v452) % 32)
	goto L129
L129:
	;
	if v442 < v456 {
		v465 = v448
		v466 = v442
		goto L126
	} else {
		goto L130
	}
L130:
	;
	if v448 != 0 {
		v485 = v444
		goto L115
	} else {
		goto L131
	}
L131:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v447)+20))
	if v458 == int32(-1) {
		v485 = v444
		goto L115
	} else {
		goto L132
	}
L132:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v138)+4)) = int64(4294967296)
	v465 = int32(1)
	v466 = int32(0)
	goto L126
L133:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v477)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v389))) = v481
	v485 = v477
	goto L115
L134:
	;
	goto L66
L135:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v505 < int32(1) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v693 = int32(-1)
	if v688 == int32(0) {
		v698 = v693
		goto L1
	} else {
		goto L180
	}
L137:
	;
	v698 = int32(0)
	goto L1
L138:
	;
	v508 = int64(0)
	v509 = int32(0)
	v514 = v509
	v520 = v509
	v521 = v508
	v522 = v508
	goto L139
L139:
	;
	v525 = base.I32_wrap_i64(v521)
	v527 = int32(1)
	if v525 < int32(0) {
		v546 = v527
		goto L143
	} else {
		goto L144
	}
L140:
	;
	goto L137
L141:
	;
	v662 = v521 + int64(1)
	v664 = int64(*(*int32)(unsafe.Add(mBase, _consts[10])))
	if v662 < v664 {
		v514 = v650
		v520 = v656
		v521 = v662
		v522 = v658
		goto L139
	} else {
		goto L179
	}
L142:
	;
	if v546 != 0 {
		v650 = v514
		v656 = v520
		v658 = v522
		goto L141
	} else {
		goto L147
	}
L143:
	;
	goto L142
L144:
	;
	v531 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v531 <= v525 {
		v546 = v527
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v533 = int32(0)
	v534 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v534+v525<<(uint(int32(2))%32))))
	if v538 == v533 {
		v546 = v527
		goto L143
	} else {
		goto L146
	}
L146:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	v542 = F_kvstoreSize(m, v541)
	mBase = m.M
	v546 = base.B2i32(v542 == int64(0))
	goto L143
L147:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v548+v525<<(uint(int32(2))%32))))
	v553 = int32(16)
	v556 = int32(0)
	v557 = int32(*(*uint8)(unsafe.Add(mBase, _consts[24])))
	*(*uint8)(unsafe.Add(mBase, uint32(v16+v553))) = uint8(v557)
	v560 = *(*int64)(unsafe.Add(mBase, _consts[25]))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v560
	v563 = *(*int64)(unsafe.Add(mBase, _consts[26]))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v563
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v565&int32(6) != 0 {
		v688 = v520
		goto L136
	} else {
		goto L148
	}
L148:
	;
	v572 = v553
	v573 = v16
	goto L149
L149:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v581) < base.Ui32(v572) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v602 = F_rioWriteBulkLongLong(m, l0, v521)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L13
	} else {
		goto L164
	}
L151:
	;
	v583 = v581
	goto L153
L152:
	;
	v583 = v572
	goto L153
L153:
	;
	if v581 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v584 = v583
	goto L156
L155:
	;
	v584 = v572
	goto L156
L156:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v585 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v591 = m.T0[v590].(func(*base.Module, int32, int32, int32) int32)(m, l0, v573, v584)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L13
	} else {
		goto L161
	}
L158:
	;
	m.T0[v585].(func(*base.Module, int32, int32, int32))(m, l0, v573, v584)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L13
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v597 + v584
	v601 = v572 - v584
	if v601 != 0 {
		v572 = v601
		v573 = v573 + v584
		goto L149
	} else {
		goto L163
	}
L161:
	;
	if v591 != 0 {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v593 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v593 | int64(2)
	v688 = v520
	goto L136
L163:
	;
	goto L150
L164:
	;
	if v602 == int32(0) {
		v688 = v520
		goto L136
	} else {
		goto L165
	}
L165:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	v608 = F_kvstoreIteratorInit(m, v606, int32(3))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L13
	} else {
		goto L166
	}
L166:
	;
	v612 = v514
	v620 = v522
	goto L168
L167:
	;
	F_kvstoreIteratorRelease(m, v608)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L13
	} else {
		goto L178
	}
L168:
	;
	v623 = F_kvstoreIteratorNext(m, v608, v16)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L13
	} else {
		goto L170
	}
L170:
	;
	if v623 == int32(0) {
		goto L167
	} else {
		goto L171
	}
L171:
	;
	v628 = v612 + int32(1)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v612&int32(1023) != 0 {
		v640 = v620
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v642 = F_rewriteObjectRio(m, l0, v629, v525)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L13
	} else {
		goto L176
	}
L173:
	;
	v632 = F_mstime(m)
	mBase = m.M
	if v632-v620 < int64(1000) {
		v640 = v620
		goto L172
	} else {
		goto L174
	}
L174:
	;
	F_sendChildInfo(m, int32(0), v628, int32(_a88))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L13
	} else {
		goto L175
	}
L175:
	;
	v640 = v632
	goto L172
L176:
	;
	if v642 == int32(-1) {
		v688 = v608
		goto L136
	} else {
		goto L177
	}
L177:
	;
	v612 = v628
	v620 = v640
	goto L168
L178:
	;
	v650 = v612
	v656 = v608
	v658 = v620
	goto L141
L179:
	;
	goto L140
L180:
	;
	F_kvstoreIteratorRelease(m, v688)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L13
	} else {
		goto L181
	}
L181:
	;
	v698 = v693
	goto L1
}
func F_startAppendOnly(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	v4 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v4 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v121
L2:
	;
	v82 = int32(_a69)
	v84 = *(*int64)(unsafe.Add(mBase, _consts[32]))
	*(*int64)(unsafe.Add(mBase, _consts[33])) = v84
	v87 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v87 != int32(-1) {
		goto L27
	} else {
		goto L28
	}
L3:
	;
	F__serverAssert(m, int32(_a96), int32(_a68), int32(964))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L14
	} else {
		goto L26
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[27])) = int32(2)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	goto L5
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if base.B2i32(v9 != int32(-1)) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v13 != int32(2) {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	F__serverLog(m, int32(2), v38, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	if v27 == int32(0) {
		goto L6
	} else {
		goto L12
	}
L9:
	;
	if v13 == int32(2) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v18 = int32(_a69)
	*(*int32)(unsafe.Add(mBase, _consts[45])) = int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v22 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v38 = int32(_a102)
	goto L7
L12:
	;
	v30 = int32(_a69)
	*(*int32)(unsafe.Add(mBase, _consts[45])) = int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v34 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v38 = int32(_a101)
	goto L7
L14:
	;
	return int32(0)
L15:
	;
	goto L2
L16:
	;
	v58 = int32(-1)
	v59 = F_rewriteAppendOnlyFileBackground(m)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L14
	} else {
		goto L22
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v48 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_killAppendOnlyChild(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L14
	} else {
		goto L21
	}
L19:
	;
	F__serverLog(m, int32(2), int32(_a100), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L16
L22:
	;
	if v59 != int32(-1) {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v63 = int32(_a69)
	*(*int32)(unsafe.Add(mBase, _consts[27])) = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v67 {
		v121 = v58
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F__serverLog(m, int32(3), int32(_a99), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v121 = v58
	goto L1
L26:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v104 != int32(-1) {
		v121 = int32(0)
		goto L1
	} else {
		goto L32
	}
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v91 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	goto L27
L30:
	;
	F__serverLog(m, int32(3), int32(_a98), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v108 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v117 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[44])) = v117
	return v117
L34:
	;
	F__serverLog(m, int32(3), int32(_a97), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	goto L33
}
func F_updateAppendOnly(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	v5 = int32(_a69)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	v8 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	if v8 != 0 {
		v17 = int32(1)
		if v8 == int32(0) {
			v27 = v17
			return v27
		} else {
			if v6 != 0 {
				v27 = v17
				return v27
			} else {
				v20 = F_startAppendOnly(m)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					if v20 != int32(-1) {
						v27 = v17
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a424)
						v27 = int32(0)
					}
					return v27
				}
			}
		}
	} else {
		if v6 == int32(0) {
			v17 = int32(1)
			if v8 == int32(0) {
				v27 = v17
				return v27
			} else {
				if v6 != 0 {
					v27 = v17
					return v27
				} else {
					v20 = F_startAppendOnly(m)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						if v20 != int32(-1) {
							v27 = v17
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a424)
							v27 = int32(0)
						}
						return v27
					}
				}
			}
		} else {
			F_stopAppendOnly(m)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return int32(1)
			}
		}
	}
}
