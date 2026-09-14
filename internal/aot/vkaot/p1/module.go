package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_addModuleEnumConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_sdsempty(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
		v18 = F_sdscatfmt(m, v13, int32(_a552), v11)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v21 = F_valkey_malloc(m, int32(88))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v21)+44)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = l5
				v27 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v27
				*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = int32(424)
				*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = int32(425)
				*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v27
				*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = int32(426)
				*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = int32(427)
				*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l2 | int32(256)
				*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v27
				*(*int32)(unsafe.Add(mBase, uint32(v21))) = v18
				*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = l3
				*(*int64)(unsafe.Add(mBase, uint32(v21+int32(52)))) = v23
				*(*int64)(unsafe.Add(mBase, uint32(v21+int32(60)))) = v23
				*(*int64)(unsafe.Add(mBase, uint32(v21+int32(68)))) = v23
				*(*int64)(unsafe.Add(mBase, uint32(v21+int32(76)))) = int64(17179869184)
				v63 = *(*int32)(unsafe.Add(mBase, _consts[247]))
				v64 = F_sdsnew(m, v18)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					v66 = F_dictAdd(m, v63, v64, v21)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						m.G0 = v11 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_createModuleObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
		v18 = int32(12)
		v21 = F_zmalloc_usable(m, v18, v7+v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v10
			*(*int64)(unsafe.Add(mBase, uint32(v21))) = int64(34359738373)
			m.G0 = v7 + int32(16)
			return v21
		}
	}
}
func F_getModuleEnumConfig(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = m.T0[v4].(func(*base.Module, int32, int32) int32)(m, v2, v3)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_moduleAllocateContext(m *base.Module) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_valkey_calloc(m, int32(72))
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_moduleBlockedClientMayTimeout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	if v3 == int32(3) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v2)+48))
		if v8 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
			return base.B2i32(v11 != int32(0))
		} else {
			return int32(0)
		}
	} else {
		return int32(1)
	}
}
func F_moduleCallCommandFilters(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[438]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v15 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return
L2:
	;
	v19 = v11 + int32(24)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v20
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v28
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	F_incrRefCount(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v36 = v11 + int32(24)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v38 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v94
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v100 != v31 {
		goto L23
	} else {
		goto L24
	}
L7:
	;
	if v38 == int32(0) {
		v92 = v24
		v93 = v26
		v94 = v28
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38+base.B2i32(v41 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v47
	goto L8
L10:
	;
	v53 = v38
	goto L11
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+8)))
	if v60&int32(1) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v92 = v88
	v93 = v87
	v94 = v86
	goto L6
L13:
	;
	v73 = v11 + int32(24)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v75 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	m.T0[v69].(func(*base.Module, int32))(m, v11+int32(8))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+40))
	if v66 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	goto L13
L18:
	;
	if v75 != 0 {
		v53 = v75
		goto L11
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v75+base.B2i32(v78 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v84
	goto L19
L21:
	;
	goto L12
L22:
	;
	F_decrRefCount(m, v31)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L28
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = int32(-1)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v107 & int32(-1966081)
	goto L26
L24:
	;
	if v94 == v34 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	F_prepareCommand(m, l0)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	goto L22
L28:
	;
	goto L1
}
func F_moduleCallCommandHelper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int64
	_ = v131
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int64
	_ = v176
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v201 int64
	_ = v201
	var v210 int64
	_ = v210
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v217 int32
	_ = v217
	var v222 int64
	_ = v222
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int64
	_ = v254
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int64
	_ = v307
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
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
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int64
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v604 int64
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v698 int64
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v893 int64
	_ = v893
	var v899 int64
	_ = v899
	var v913 int32
	_ = v913
	var v914 int64
	_ = v914
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v939 int32
	_ = v939
	var v947 int32
	_ = v947
	var v953 int32
	_ = v953
	var v959 int32
	_ = v959
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = int32(0)
	v21 = int32(9116376)
	goto L1
L1:
	;
	v22 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[18])) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = v24 & int32(2048)
	if v26 == v22 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	F__serverAssert(m, int32(_a802), int32(_a756), int32(6949))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L37
	} else {
		goto L299
	}
L3:
	;
	F__serverAssert(m, int32(_a801), int32(_a756), int32(6948))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L37
	} else {
		goto L298
	}
L4:
	;
	F__serverAssert(m, int32(_a800), int32(_a756), int32(6947))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L37
	} else {
		goto L297
	}
L5:
	;
	if v849&int32(64) == int32(0) {
		goto L272
	} else {
		goto L273
	}
L6:
	;
	if v170&int32(1024) != 0 {
		v849 = v170
		goto L5
	} else {
		goto L252
	}
L7:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l1)+292))
	F_scriptSetOriginalClientSlot(m, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L37
	} else {
		goto L251
	}
L8:
	;
	F__serverAssert(m, int32(_a804), int32(_a756), int32(6800))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L37
	} else {
		goto L250
	}
L9:
	;
	F__serverAssert(m, int32(_a805), int32(_a756), int32(6622))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L37
	} else {
		goto L249
	}
L10:
	;
	if l4&int32(2048) != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v29 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	goto L12
L12:
	;
	if base.B2i32(v30 != v29) == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+96))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v42
	v46 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+224)) = uint8(v46)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l3
	if l4&int32(8) != 0 {
		v57 = int32(3)
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v37 | int32(4096)
	goto L14
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v60 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+224)) = uint8(v57)
	goto L16
L18:
	;
	if l4&int32(16) == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+224)))
	v57 = v56
	goto L17
L20:
	;
	if l4&int32(64) != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+40)) = v63 + int32(1)
	goto L20
L22:
	;
	v85 = l4 & int32(256)
	v86 = int32(0)
	if v82&int32(32) == v86 {
		v105 = v86
		goto L29
	} else {
		goto L30
	}
L23:
	;
	if v26 == int32(0) {
		v82 = l4
		v83 = v41
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v82 = l4
	v83 = v41
	goto L22
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v71&int32(-1073742081) | int32(256)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+328))
	if v80 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v81 = l4 | int32(32)
	goto L28
L27:
	;
	v81 = l4
	goto L28
L28:
	;
	v82 = v81
	v83 = v79
	goto L22
L29:
	;
	if l2 != 0 {
		goto L39
	} else {
		goto L40
	}
L30:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v91 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+328)) = v95
	v105 = v95
	goto L29
L32:
	;
	v94 = v91
	goto L34
L33:
	;
	v94 = v83 + int32(328)
	goto L34
L34:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(138)
	if v85 == int32(0) {
		v849 = v82
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v101 = F_sdsnew(m, int32(_a806))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v101
	v849 = v82
	goto L5
L39:
	;
	F_moduleCallCommandFilters(m, l1)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L37
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(8)
	v849 = v82
	goto L5
L41:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v112 = F_lookupCommand(m, v110, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v112
	if v85 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v127 = F_commandCheckArity(m, v125, v126, v120)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L37
	} else {
		goto L50
	}
L44:
	;
	v120 = v17 + int32(76)
	goto L46
L45:
	;
	v120 = int32(0)
	goto L46
L46:
	;
	v121 = F_commandCheckExistence(m, l1, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L37
	} else {
		goto L47
	}
L47:
	;
	if v121 != 0 {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(44)
	v849 = v82
	goto L5
L49:
	;
	v131 = F_getCommandFlags(m, l1)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L37
	} else {
		goto L52
	}
L50:
	;
	if v127 != 0 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(28)
	v849 = v82
	goto L5
L52:
	;
	if v82&int32(64) == int32(0) {
		v170 = v82
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v170&int32(512) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L54:
	;
	if v131&int64(64) == int64(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v164 = F_scriptAllowsOOM(m)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L37
	} else {
		goto L65
	}
L56:
	;
	if v26 == int32(0) {
		v170 = v82
		goto L53
	} else {
		goto L64
	}
L57:
	;
	if v26 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(70)
	if v85 == int32(0) {
		v849 = v82
		goto L5
	} else {
		goto L61
	}
L59:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[424]))
	if v144 != 0 {
		goto L55
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v149 = F_sdsempty(m)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L37
	} else {
		goto L62
	}
L62:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v152
	v157 = F_sdscatfmt(m, v149, int32(_a807), v17+int32(64))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L37
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v157
	v849 = v82
	goto L5
L64:
	;
	goto L55
L65:
	;
	v166 = F_scriptIsWriteDirty(m)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L37
	} else {
		goto L66
	}
L66:
	;
	if v164|v166 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v169 = v82 & int32(-513)
	goto L69
L68:
	;
	v169 = v82
	goto L69
L69:
	;
	v170 = v169
	goto L53
L70:
	;
	v307 = v131 & int64(1)
	if v170&int32(128) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L71:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v301 | int32(32768)
	goto L70
L72:
	;
	v176 = *(*int64)(unsafe.Add(mBase, _consts[321]))
	if v176 == int64(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	if v131&int64(4) == int64(0) {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v183&int32(16) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v288 == int32(0) {
		goto L70
	} else {
		goto L107
	}
L76:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _consts[425]))
	v288 = v287
	goto L75
L77:
	;
	v196 = F_zmalloc_used_memory(m)
	mBase = m.M
	goto L79
L78:
	;
	v288 = base.B2i32(v283 == int32(-1))
	goto L75
L79:
	;
	v201 = *(*int64)(unsafe.Add(mBase, _consts[321]))
	if v201 != int64(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v283 = v276
	goto L78
L82:
	;
	v210 = base.I64_extend_i32_u(v196)
	goto L86
L83:
	;
	v276 = int32(0)
	goto L81
L85:
	;
	v214 = int32(_a20)
	v215 = *(*int64)(unsafe.Add(mBase, _consts[322]))
	v217 = *(*int32)(unsafe.Add(mBase, _consts[323]))
	if base.I64_extend_i32_u(v217) <= v215 {
		v232 = int32(0)
		goto L88
	} else {
		goto L89
	}
L86:
	;
	if base.Ui64(v201) < base.Ui64(v210) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v283 = int32(0)
	goto L78
L88:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v234 == int32(0) {
		v241 = v232
		goto L93
	} else {
		goto L94
	}
L89:
	;
	v222 = base.I64_div_s(v215, int64(16384))
	v229 = v217 - base.I32_wrap_i64(v215+v222*int64(44)) + int32(-44)
	if base.Ui32(v217) < base.Ui32(v229) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v231 = int32(0)
	goto L92
L91:
	;
	v231 = v229
	goto L92
L92:
	;
	v232 = v231
	goto L88
L93:
	;
	v242 = F_clusterIsAnySlotExporting(m)
	mBase = m.M
	if v242 == int32(0) {
		v247 = v241
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _consts[66]))
	v239 = F_sdsAllocSize(m, v238)
	mBase = m.M
	v241 = v239 + v232
	goto L93
L95:
	;
	v248 = int32(0)
	v250 = v196 - v247
	if base.Ui32(v196) < base.Ui32(v250) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v245 = F_clusterGetTotalSlotExportBufferMemory(m)
	mBase = m.M
	v247 = v245 + v241
	goto L95
L97:
	;
	v252 = v248
	goto L99
L98:
	;
	v252 = v250
	goto L99
L99:
	;
	v254 = *(*int64)(unsafe.Add(mBase, _consts[321]))
	goto L100
L100:
	;
	if base.Ui64(v210) <= base.Ui64(v254) {
		v276 = v248
		goto L81
	} else {
		goto L102
	}
L102:
	;
	if base.Ui64(base.I64_extend_i32_u(v252)) <= base.Ui64(v254) {
		v276 = v248
		goto L81
	} else {
		goto L103
	}
L103:
	;
	goto L104
L104:
	;
	v276 = int32(-1)
	goto L81
L107:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(51)
	if v85 == int32(0) {
		v849 = v170
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _consts[426]))
	v297 = F_objectGetVal(m, v296)
	mBase = m.M
	v298 = F_sdsdup(m, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L37
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v298
	v849 = v170
	goto L5
L110:
	;
	if v170&int32(32) == int32(0) {
		goto L116
	} else {
		goto L117
	}
L111:
	;
	if v307 == int64(0) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(51)
	if v85 == int32(0) {
		v849 = v170
		goto L5
	} else {
		goto L113
	}
L113:
	;
	v318 = F_sdsempty(m)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L37
	} else {
		goto L114
	}
L114:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v321
	v326 = F_sdscatfmt(m, v318, int32(_a808), v17+int32(48))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L37
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v326
	v849 = v170
	goto L5
L116:
	;
	v413 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v413 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L117:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+200)))
	if v333&int32(8) == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v351 = F_ACLCheckAllUserCommandPerm(m, v105, v345, v346, v347, v348, v17+int32(72))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L37
	} else {
		goto L121
	}
L119:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v344 = v341 + int32(28)
	goto L118
L120:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v344 = v338 + int32(52)
	goto L118
L121:
	;
	if v351 == int32(0) {
		goto L116
	} else {
		goto L122
	}
L122:
	;
	v357 = int32(0)
	v358 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	goto L123
L123:
	;
	if v358 != v357 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v361 = int32(4)
	goto L126
L125:
	;
	v361 = int32(3)
	goto L126
L126:
	;
	if v351 != int32(2) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v374 = F_sdsdup(m, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L37
	} else {
		goto L130
	}
L128:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v366+v367<<(uint(int32(2))%32))))
	v372 = F_objectGetVal(m, v371)
	mBase = m.M
	v373 = v372
	goto L127
L129:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+140))
	v373 = v365
	goto L127
L130:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l1)+328))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	F_addACLLogEntry(m, v376, v351, v361, int32(-1), v379, v374)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L37
	} else {
		goto L131
	}
L131:
	;
	if v85 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(2)
	v849 = v170
	goto L5
L133:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l1)+328))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v386+v387<<(uint(int32(2))%32))))
	v392 = F_objectGetVal(m, v391)
	mBase = m.M
	v394 = F_getAclErrorMessage(m, v351, v384, v385, v392, int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L37
	} else {
		goto L134
	}
L134:
	;
	v396 = F_sdsempty(m)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L37
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v394
	v402 = F_sdscatfmt(m, v396, int32(_a809), v17+int32(32))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L37
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v402
	F_sdsfree(m, v394)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L37
	} else {
		goto L137
	}
L137:
	;
	goto L132
L138:
	;
	v509 = v170 & int32(64)
	if v509 == int32(0) {
		goto L6
	} else {
		goto L168
	}
L139:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v416)))
	if v418 != int64(-1) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	if v433 != 0 {
		goto L138
	} else {
		goto L147
	}
L141:
	;
	v422 = int32(1)
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+200)))
	if v423&v422 != 0 {
		v430 = v422
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v433 = int32(1)
	goto L140
L143:
	;
	v433 = v430
	goto L140
L144:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v416)+216))
	if v426 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v428 = F_isImportSlotMigrationJob(m, v426)
	mBase = m.M
	v430 = v428
	goto L143
L146:
	;
	v433 = int32(0)
	goto L140
L147:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+200))
	v441 = v434&int32(-131073) | v438&int32(131072)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v441
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v444&int32(512) | v441&int32(-513)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v456 = F_clusterSlotByCommand(m, v451, v452, v453, l1+int32(288))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L37
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+292)) = v456
	v461 = F_getNodeByQuery(m, l1, v17+int32(72))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L37
	} else {
		goto L149
	}
L149:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v464)))
	goto L150
L150:
	;
	if v461 == v465 {
		goto L138
	} else {
		goto L151
	}
L151:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	if v467 != 0 {
		goto L8
	} else {
		goto L152
	}
L152:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	switch v468 + int32(-5) {
	case 0:
		goto L156
	default:
		goto L155
	case 2:
		goto L157
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = v505
	v849 = v170
	goto L5
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v501
	v505 = v502
	goto L153
L155:
	;
	v495 = int32(63)
	if v85 == int32(0) {
		v505 = v495
		goto L153
	} else {
		goto L166
	}
L156:
	;
	if v85 != 0 {
		goto L162
	} else {
		goto L163
	}
L157:
	;
	if v85 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v473 = F_sdsempty(m)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L37
	} else {
		goto L160
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(69)
	v849 = v170
	goto L5
L160:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v476
	v480 = F_sdscatfmt(m, v473, int32(_a810), v17)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L37
	} else {
		goto L161
	}
L161:
	;
	v501 = v480
	v502 = int32(69)
	goto L154
L162:
	;
	v484 = F_sdsempty(m)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L37
	} else {
		goto L164
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(38)
	v849 = v170
	goto L5
L164:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v486)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v487
	v493 = F_sdscatfmt(m, v484, int32(_a811), v17+int32(16))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L37
	} else {
		goto L165
	}
L165:
	;
	v501 = v493
	v502 = int32(38)
	goto L154
L166:
	;
	v499 = F_sdsnew(m, int32(_a812))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L37
	} else {
		goto L167
	}
L167:
	;
	v501 = v499
	v502 = v495
	goto L154
L168:
	;
	if v26 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	v661 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v661 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L170:
	;
	v536 = int32(0)
	v537 = int32(_a20)
	v538 = *(*int32)(unsafe.Add(mBase, _consts[427]))
	v542 = *(*int32)(unsafe.Add(mBase, _consts[428]))
	v547 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v552 = *(*int32)(unsafe.Add(mBase, _consts[429]))
	goto L183
L171:
	;
	if v307 == int64(0) {
		goto L169
	} else {
		goto L181
	}
L172:
	;
	v514 = F_scriptIsReadOnly(m)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L37
	} else {
		goto L174
	}
L173:
	;
	v528 = F_scriptIsWriteDirty(m)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L37
	} else {
		goto L178
	}
L174:
	;
	if v514 == int32(0) {
		goto L173
	} else {
		goto L175
	}
L175:
	;
	if v131&int64(65537) == int64(0) {
		goto L173
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(51)
	v525 = F_sdsnew(m, int32(_a813))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L37
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v525
	v849 = v170
	goto L5
L178:
	;
	if v528 != 0 {
		goto L169
	} else {
		goto L179
	}
L179:
	;
	if base.B2i32(v307 == int64(0)) == int32(0) {
		goto L170
	} else {
		goto L180
	}
L180:
	;
	goto L169
L181:
	;
	goto L170
L182:
	;
	v565 = int32(0)
	v567 = *(*int32)(unsafe.Add(mBase, _consts[430]))
	if v567 == v565 {
		goto L189
	} else {
		goto L190
	}
L183:
	;
	if base.B2i32(v538 == v536)|base.B2i32(v542 == v536)|base.B2i32(v547 != v536)|base.B2i32(v542 <= v552) != 0 {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(70)
	if v85 == int32(0) {
		v849 = v170
		goto L5
	} else {
		goto L185
	}
L185:
	;
	v560 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	v561 = F_objectGetVal(m, v560)
	mBase = m.M
	v562 = F_sdsdup(m, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L37
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v562
	v849 = v170
	goto L5
L187:
	;
	v601 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v601 != 0 {
		goto L198
	} else {
		goto L199
	}
L188:
	;
	goto L187
L189:
	;
	v580 = int32(0)
	v582 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v582 == v580 {
		v599 = v580
		goto L188
	} else {
		goto L193
	}
L190:
	;
	v571 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	if v571 < int32(1) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v576 = *(*int32)(unsafe.Add(mBase, _consts[431]))
	if v576 == int32(-1) {
		v599 = int32(2)
		goto L188
	} else {
		goto L192
	}
L192:
	;
	goto L189
L193:
	;
	v585 = int32(0)
	v587 = *(*int32)(unsafe.Add(mBase, _consts[51]))
	if v587 == int32(-1) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v599 = int32(1)
	goto L188
L195:
	;
	v591 = *(*int32)(unsafe.Add(mBase, _consts[50]))
	if v591 != int32(-1) {
		v599 = v585
		goto L188
	} else {
		goto L196
	}
L196:
	;
	v594 = int32(0)
	v596 = *(*int32)(unsafe.Add(mBase, _consts[432]))
	*(*int32)(unsafe.Add(mBase, _consts[433])) = v596
	goto L194
L197:
	;
	if v599 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L198:
	;
	v604 = *(*int64)(unsafe.Add(mBase, uint32(v601)))
	if v604 != int64(-1) {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v622 = int32(0)
	goto L197
L200:
	;
	v622 = base.B2i32(v619 != int32(0))
	goto L197
L201:
	;
	v608 = int32(1)
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+200)))
	if v609&v608 != 0 {
		v616 = v608
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v619 = int32(1)
	goto L200
L203:
	;
	v619 = v616
	goto L200
L204:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v601)+216))
	if v612 != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v614 = F_isImportSlotMigrationJob(m, v612)
	mBase = m.M
	v616 = v614
	goto L203
L206:
	;
	v619 = int32(0)
	goto L200
L207:
	;
	v632 = int32(_a20)
	v633 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v634 = int32(0)
	v637 = *(*int32)(unsafe.Add(mBase, _consts[434]))
	if (base.B2i32(v633 == v634)|base.B2i32(v637 == v634)|v622)&int32(1) != 0 {
		goto L212
	} else {
		goto L213
	}
L208:
	;
	if v622 != 0 {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(70)
	if v85 == int32(0) {
		v849 = v170
		goto L5
	} else {
		goto L210
	}
L210:
	;
	v629 = F_writeCommandsGetDiskErrorMessage(m, v599)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L37
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v629
	v849 = v170
	goto L5
L212:
	;
	if v26 == int32(0) {
		goto L169
	} else {
		goto L216
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(70)
	if v85 == int32(0) {
		v849 = v170
		goto L5
	} else {
		goto L214
	}
L214:
	;
	v649 = *(*int32)(unsafe.Add(mBase, _consts[435]))
	v650 = F_objectGetVal(m, v649)
	mBase = m.M
	v651 = F_sdsdup(m, v650)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L37
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v651
	v849 = v170
	goto L5
L216:
	;
	F_scriptSetWriteDirtyFlag(m)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L37
	} else {
		goto L217
	}
L217:
	;
	goto L169
L218:
	;
	if v26 == int32(0) {
		goto L6
	} else {
		goto L228
	}
L219:
	;
	v665 = *(*int32)(unsafe.Add(mBase, _consts[423]))
	if v665 == int32(14) {
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v669 = *(*int32)(unsafe.Add(mBase, _consts[436]))
	if v669 != 0 {
		goto L218
	} else {
		goto L221
	}
L221:
	;
	if v131&int64(1024) != int64(0) {
		goto L218
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(70)
	if v85 == int32(0) {
		v849 = v170
		goto L5
	} else {
		goto L223
	}
L223:
	;
	if v26 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v685 = *(*int32)(unsafe.Add(mBase, _consts[437]))
	v686 = F_objectGetVal(m, v685)
	mBase = m.M
	v687 = F_sdsdup(m, v686)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L37
	} else {
		goto L227
	}
L225:
	;
	v681 = F_sdsnew(m, int32(_a814))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L37
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v681
	v849 = v170
	goto L5
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v687
	v849 = v170
	goto L5
L228:
	;
	v693 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v693 == int32(0) {
		goto L6
	} else {
		goto L229
	}
L229:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v698 = *(*int64)(unsafe.Add(mBase, uint32(v696)))
	if v698 != int64(-1) {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	if v713 != 0 {
		goto L6
	} else {
		goto L237
	}
L231:
	;
	v702 = int32(1)
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696)+200)))
	if v703&v702 != 0 {
		v710 = v702
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v713 = int32(1)
	goto L230
L233:
	;
	v713 = v710
	goto L230
L234:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v696)+216))
	if v706 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v708 = F_isImportSlotMigrationJob(m, v706)
	mBase = m.M
	v710 = v708
	goto L233
L236:
	;
	v713 = int32(0)
	goto L230
L237:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l1)+292))
	if v714 == int32(-1) {
		goto L7
	} else {
		goto L238
	}
L238:
	;
	v717 = F_scriptAllowsCrossSlot(m)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L37
	} else {
		goto L239
	}
L239:
	;
	if v717 != 0 {
		goto L7
	} else {
		goto L240
	}
L240:
	;
	v719 = F_scriptGetSlot(m)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L37
	} else {
		goto L242
	}
L241:
	;
	v726 = F_scriptGetSlot(m)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L37
	} else {
		goto L245
	}
L242:
	;
	if v719 != int32(-1) {
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l1)+292))
	F_scriptSetSlot(m, v723)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L37
	} else {
		goto L244
	}
L244:
	;
	goto L7
L245:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l1)+292))
	if v726 == v728 {
		goto L7
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(70)
	if v85 == int32(0) {
		v849 = v170
		goto L5
	} else {
		goto L247
	}
L247:
	;
	v735 = F_sdsnew(m, int32(_a803))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L37
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v735
	v849 = v170
	goto L5
L249:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L250:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	goto L6
L252:
	;
	v757 = int32(_a20)
	v759 = l4 & int32(1)
	v761 = *(*int32)(unsafe.Add(mBase, _consts[283]))
	v762 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[283])) = v759 & base.B2i32(v761 != v762)
	if v759 == v762 {
		v781 = int32(4)
		goto L253
	} else {
		goto L254
	}
L253:
	;
	F_call(m, l1, v781)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L37
	} else {
		goto L258
	}
L254:
	;
	if v170&int32(2) != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v773 = int32(4)
	goto L257
L256:
	;
	v773 = int32(5)
	goto L257
L257:
	;
	v776 = int32(2)
	v781 = v773 | int32(base.Ui32(v170)>>(uint(int32(1))%32))&v776 ^ v776
	goto L253
L258:
	;
	if v26 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v795 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[18])) = v795
	*(*int32)(unsafe.Add(mBase, _consts[283])) = v761
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+200)))
	if v799&int32(16) == v795 {
		v849 = v170
		goto L5
	} else {
		goto L263
	}
L260:
	;
	if v509 == int32(0) {
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v789)+96))
	if v788 == v790 {
		goto L259
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v789)+96)) = v788
	goto L259
L263:
	;
	if v26 != 0 {
		goto L4
	} else {
		goto L264
	}
L264:
	;
	if v170&int32(2048) == int32(0) {
		goto L3
	} else {
		goto L265
	}
L265:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v808 == int32(0) {
		goto L2
	} else {
		goto L266
	}
L266:
	;
	v812 = F_valkey_malloc(m, int32(136))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L37
	} else {
		goto L267
	}
L267:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v812)+16)) = l1
	v817 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v812)+12)) = v817
	*(*int32)(unsafe.Add(mBase, uint32(v812)+8)) = v815
	*(*int64)(unsafe.Add(mBase, uint32(v812))) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v812)+20)) = (v817 - v814&int32(1)) & l0
	v833 = F__emscripten_memset_bulkmem(m, v812+int32(24), base.I32_extend8_s(v817), int32(112))
	mBase = m.M
	goto L268
L268:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v834)+52)) = v812
	if v781&int32(1) != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	if v781&int32(2) != 0 {
		v849 = v170
		goto L5
	} else {
		goto L271
	}
L270:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v838 | int32(524288)
	goto L269
L271:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v844 | int32(1048576)
	v849 = v170
	goto L5
L272:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	if v925 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L273:
	;
	v860 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	if v860 == int32(0) {
		goto L272
	} else {
		goto L274
	}
L274:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864+int32(-1)))))
	switch v867 & int32(7) {
	case 0:
		goto L280
	case 1:
		goto L279
	case 2:
		goto L278
	case 3:
		goto L277
	case 4:
		goto L276
	default:
		v884 = int32(0)
		goto L275
	}
L275:
	;
	F_afterErrorReply(m, l1, v864, v884, int32(0))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L37
	} else {
		goto L281
	}
L276:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v864+int32(-17))))
	v884 = v883
	goto L275
L277:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v864+int32(-9))))
	v884 = v880
	goto L275
L278:
	;
	v877 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v864+int32(-5)))))
	v884 = v877
	goto L275
L279:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864+int32(-3)))))
	v884 = v874
	goto L275
L280:
	;
	v884 = int32(base.Ui32(v867) >> (uint(int32(3)) % 32))
	goto L275
L281:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v891 = int32(0)
	v893 = *(*int64)(unsafe.Add(mBase, _consts[421]))
	if v888 == v891 {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	goto L272
L283:
	;
	*(*int64)(unsafe.Add(mBase, _consts[422])) = v893
	goto L282
L284:
	;
	v899 = *(*int64)(unsafe.Add(mBase, _consts[422]))
	if v893 <= v899 {
		goto L283
	} else {
		goto L285
	}
L285:
	;
	goto L288
L286:
	;
	v913 = v888 + int32(120)
	v914 = *(*int64)(unsafe.Add(mBase, uint32(v913)))
	*(*int64)(unsafe.Add(mBase, uint32(v913))) = v914 + int64(1)
	goto L283
L288:
	;
	goto L286
L290:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v929 == int32(0) {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v925
	goto L290
L292:
	;
	if v26 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v929)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v929)+40)) = v932 + int32(-1)
	goto L292
L294:
	;
	m.G0 = v17 + int32(80)
	return
L295:
	;
	F_scriptClusterSlotStatsInvalidateSlotIfApplicable(m)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L37
	} else {
		goto L296
	}
L296:
	;
	goto L294
L297:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L298:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L299:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_moduleCallCommandUnblockedHandler(m *base.Module, l0 int32) {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int64
	_ = v107
	var v111 int64
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v133 int64
	_ = v133
	var v137 int64
	_ = v137
	var v140 int32
	_ = v140
	var v148 int64
	_ = v148
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
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a758), int32(_a756), int32(932))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L11
	} else {
		goto L39
	}
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_moduleReleaseTempClient(m, l0)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L11
	} else {
		goto L38
	}
L4:
	;
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(32)))) = v23
	v25 = int32(64)
	*(*int64)(unsafe.Add(mBase, uint32(v10+v25))) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(56)))) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(48)))) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(72)))) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(561)
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	if v59 == v58 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v18 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v90
	v94 = *(*int32)(unsafe.Add(mBase, _consts[384]))
	v95 = int32(0)
	v96 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v97 = m.T0[v96].(func(*base.Module) int64)(m)
	mBase = m.M
	if v94 == v95 {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v78 = F_createClient(m, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v62 = int32(0)
	v64 = v59 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[383])) = v64
	v67 = *(*int32)(unsafe.Add(mBase, _consts[385]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v64<<(uint(int32(2))%32))))
	v73 = *(*int32)(unsafe.Add(mBase, _consts[386]))
	if base.Ui32(v73) <= base.Ui32(v64) {
		v90 = v71
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[386])) = v64
	v90 = v71
	goto L7
L11:
	;
	return
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+328)) = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+200)) = v82 | int32(1073741824)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v78)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+204)) = v86 | int32(268435456)
	v90 = v78
	goto L7
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v111
	v115 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	*(*int32)(unsafe.Add(mBase, _consts[177])) = v119 + int32(1)
	goto L18
L14:
	;
	v107 = *(*int64)(unsafe.Add(mBase, _consts[387]))
	v111 = v107*int64(1000) + v97
	goto L13
L15:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v103 = base.I32_div_s(int32(1000000), v102)
	v111 = v97 + base.I64_extend_i32_s(v103)
	goto L13
L16:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+28))
	v154 = F_selectDb(m, v90, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L11
	} else {
		goto L22
	}
L17:
	;
	goto L16
L18:
	;
	if v119 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	goto L21
L20:
	;
	v129 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[178])) = v127
	v133 = base.I64_div_s(v127, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[54])) = v133
	v137 = base.I64_div_s(v127, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[109])) = v137
	v140 = *(*int32)(unsafe.Add(mBase, _consts[179]))
	F_lrulfu_updateClockAndPolicy(m, v133, int32(base.Ui32(v140&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v148 = *(*int64)(unsafe.Add(mBase, _consts[54]))
	*(*int64)(unsafe.Add(mBase, _consts[12])) = v148
	goto L17
L21:
	;
	v127 = F_ustime(m)
	mBase = m.M
	goto L20
L22:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v156 + int32(1)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	if v160 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v224 + int32(-1)
	F_moduleFreeContext(m, v10+int32(8))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L11
	} else {
		goto L37
	}
L24:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v13)+128))
	F_invokeReplyHandlers(m, v10+int32(8), l0, v13+int32(32), v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L11
	} else {
		goto L36
	}
L25:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v163 = F_sdsnewlen(m, v161, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v165 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	if v168 == v165 {
		v195 = v163
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v200 = F_callReplyCreate(m, v195, v198, int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L11
	} else {
		goto L34
	}
L28:
	;
	v175 = v163
	v176 = v167
	goto L29
L29:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v183 = F_sdscatlen(m, v175, v179+int32(13), v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L11
	} else {
		goto L31
	}
L30:
	;
	v195 = v183
	goto L27
L31:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	F_listDelNode(m, v185, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+20))
	if v190 != 0 {
		v175 = v183
		v176 = v189
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = int32(0)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	m.T0[v207].(func(*base.Module, int32, int32, int32))(m, v10+int32(8), v200, v206)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	goto L23
L36:
	;
	goto L23
L37:
	;
	goto L3
L38:
	;
	m.G0 = v10 + int32(80)
	return
L39:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_moduleClientIsBlockedOnKeys(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+48))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+44))
	return v4
}
func F_moduleCloseKey(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v5 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v13&int32(2) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+48)))
	v12 = base.B2i32(v7&int32(2) == int32(0))
	goto L1
L3:
	;
	v12 = int32(1)
	goto L1
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v25 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	if v12 == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_signalModifiedKey(m, v20, v21, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	goto L4
L9:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v60 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v28 == int32(0) {
		v34 = v25
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	switch v35&int32(15) + int32(-3) {
	case 0:
		goto L15
	default:
		goto L9
	case 3:
		goto L14
	}
L12:
	;
	F_moduleFreeKeyIterator(m, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = v33
	goto L11
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v51 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v40 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	goto L9
L17:
	;
	F_zsetFreeLexRange(m, l0+int32(56))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_signalKeyAsReady(m, v54, v55, int32(6))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_decrRefCount(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	F__serverAssert(m, int32(_a797), int32(_a756), int32(4351))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	return
}
func F_moduleConfigValidityCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v124
L2:
	;
	goto L32
L3:
	;
	v15 = base.I64_extend_i32_u(l2)
	if v15&int64(4294966284) == int64(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v114 = int32(10)
	goto L2
L5:
	;
	v42 = int32(28)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v43&int32(255) == int32(0) {
		v114 = v42
		goto L2
	} else {
		goto L16
	}
L6:
	;
	F_serverLogRaw(m, int32(3), v35)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	if l3 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v35 = int32(_a1618)
	goto L6
L9:
	;
	if l3 == int32(4) {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	if v15&int64(640) == int64(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v35 = int32(_a1619)
	goto L6
L12:
	;
	if v15&int64(256) == int64(0) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v35 = int32(_a1620)
	goto L6
L14:
	;
	return int32(0)
L15:
	;
	v114 = int32(28)
	goto L2
L16:
	;
	v51 = v43
	v55 = int32(0)
	goto L17
L17:
	;
	if base.Ui32((v51+int32(-48))&int32(255)) < base.Ui32(int32(10)) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v98 = F_listSearchKey(m, v97, l1)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L14
	} else {
		goto L28
	}
L19:
	;
	v92 = v55 + int32(1)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v92))))
	if v94&int32(255) != 0 {
		v51 = v94
		v55 = v92
		goto L17
	} else {
		goto L26
	}
L20:
	;
	if base.Ui32((v51&int32(-33)+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v72 = v51 & int32(255)
	if v72 == int32(45) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	if v72 == int32(95) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v78 {
		v114 = v42
		goto L2
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = base.I32_extend8_s(v51)
	F__serverLog(m, int32(3), int32(_a1621), v11+int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v114 = v42
	goto L2
L26:
	;
	goto L18
L27:
	;
	v101 = int32(7)
	v103 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v103 {
		v114 = v101
		goto L2
	} else {
		goto L30
	}
L28:
	;
	if v98 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v124 = int32(0)
	goto L1
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
	F__serverLog(m, int32(3), int32(_a1622), v11)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	v114 = v101
	goto L2
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = v114
	v124 = int32(1)
	goto L1
}
func F_moduleDelKeyIfEmpty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v50 int32
	_ = v50
	v2 = int32(0)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v5&int32(2) == v2 {
		v50 = v2
		return v50
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v10 == int32(0) {
			v50 = v2
			return v50
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			switch v13&int32(15) + int32(-1) {
			case 0:
				v36 = F_listTypeLength(m, v10)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					if v36 != 0 {
						v50 = v2
						return v50
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v38 == int32(0) {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v45 = F_dbDelete(m, v43, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
								v50 = int32(1)
								return v50
							}
						} else {
							F_moduleFreeKeyIterator(m, l0)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v45 = F_dbDelete(m, v43, v44)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
									v50 = int32(1)
									return v50
								}
							}
						}
					}
				}
			case 1:
				v18 = F_setTypeSize(m, v10)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					if v18 == int32(0) {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v38 == int32(0) {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v45 = F_dbDelete(m, v43, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
								v50 = int32(1)
								return v50
							}
						} else {
							F_moduleFreeKeyIterator(m, l0)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v45 = F_dbDelete(m, v43, v44)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
									v50 = int32(1)
									return v50
								}
							}
						}
					} else {
						v50 = v2
						return v50
					}
				}
			case 2:
				v24 = F_zsetLength(m, v10)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v24 == int32(0) {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v38 == int32(0) {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v45 = F_dbDelete(m, v43, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
								v50 = int32(1)
								return v50
							}
						} else {
							F_moduleFreeKeyIterator(m, l0)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v45 = F_dbDelete(m, v43, v44)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
									v50 = int32(1)
									return v50
								}
							}
						}
					} else {
						v50 = v2
						return v50
					}
				}
			case 3:
				v28 = F_hashTypeLength(m, v10)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					if v28 == int32(0) {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v38 == int32(0) {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v45 = F_dbDelete(m, v43, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
								v50 = int32(1)
								return v50
							}
						} else {
							F_moduleFreeKeyIterator(m, l0)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v45 = F_dbDelete(m, v43, v44)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
									v50 = int32(1)
									return v50
								}
							}
						}
					} else {
						v50 = v2
						return v50
					}
				}
			default:
				v50 = v2
				return v50
			case 5:
				v32 = F_objectGetVal(m, v10)
				mBase = m.M
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
				if v33 == int32(0) {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v38 == int32(0) {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v45 = F_dbDelete(m, v43, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
							v50 = int32(1)
							return v50
						}
					} else {
						F_moduleFreeKeyIterator(m, l0)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v45 = F_dbDelete(m, v43, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
								v50 = int32(1)
								return v50
							}
						}
					}
				} else {
					v50 = v2
					return v50
				}
			}
		}
	}
}
func F_moduleFreeContext(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	if v9&int32(528) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_autoMemoryCollect(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L6
	}
L2:
	;
	v12 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	*(*int32)(unsafe.Add(mBase, _consts[177])) = v14 + int32(-1)
	goto L3
L3:
	;
	F_postExecutionUnitOperations(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	goto L1
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v22 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v36 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v38 == v36 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v27 = v22
	goto L9
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	F_valkey_free(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L7
L11:
	;
	if v29 != 0 {
		v27 = v29
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v56&int32(64) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	F_valkey_free(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v46 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v50
	F__serverLog(m, int32(3), int32(_a754), v7)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	m.G0 = v7 + int32(16)
	return
L19:
	;
	if v56&int32(128) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_moduleReleaseTempClient(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	if v56&int32(2048) == int32(0) {
		goto L18
	} else {
		goto L25
	}
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v69 = F_freeClient(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	goto L18
}
func F_moduleFreeModuleStructure(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_listRelease(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_listRelease(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_listRelease(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_listRelease(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_listRelease(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_sdsfree(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v22 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	F_sdsfree(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v28 < int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	F_valkey_free(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L17
	}
L12:
	;
	v34 = int32(0)
	goto L13
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v34<<(uint(int32(2))%32))))
	F_decrRefCount(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L11
L15:
	;
	v43 = v34 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v43 < v44 {
		v34 = v43
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_valkey_free(m, v22)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L8
L19:
	;
	return
}
func F_moduleFromCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v2 == int32(562) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		return v14
	} else {
		F__serverAssert(m, int32(_a815), int32(_a756), int32(7307))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
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
func F_moduleGetClusterNodeInfoForClient(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int64
	_ = v136
	var v138 int32
	_ = v138
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v154 int64
	_ = v154
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	v9 = int32(1)
	if l2&int32(3) == int32(0) {
		v31 = l2
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v263
L2:
	;
	v65 = F_clusterLookupNode(m, l2, v64)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L19
	}
L3:
	;
	v64 = v56 - l2
	goto L2
L4:
	;
	v35 = v31
	goto L12
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v20 = l2
	goto L8
L7:
	;
	v64 = l2 - l2
	goto L2
L8:
	;
	v24 = v20 + int32(1)
	if v24&int32(3) == int32(0) {
		v31 = v24
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v29 != 0 {
		v20 = v24
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v56 = v24
	goto L3
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v44 = int32(-2139062144)
	if (int32(16843008)-v41|v41)&v44 == v44 {
		v35 = v35 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v50 = v35
	goto L15
L14:
	;
	goto L13
L15:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v54 != 0 {
		v50 = v50 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v56 = v50
	goto L3
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	if v65 == int32(0) {
		v263 = v9
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+88))
	goto L21
L21:
	;
	if v71&int32(96) != 0 {
		v263 = v9
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if l3 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if l4 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L24:
	;
	v76 = F_clusterNodeIp(m, v65, l1)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	goto L29
L26:
	;
	goto L23
L27:
	;
	goto L26
L28:
	;
	v108 = v87
	goto L35
L29:
	;
	v83 = l3
	v85 = int32(46)
	v87 = v76
	goto L31
L30:
	;
	v98 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v98)
	goto L28
L31:
	;
	v89 = v85 + int32(-1)
	if v89 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v92)
	v94 = int32(1)
	if v92 != 0 {
		v83 = v83 + v94
		v85 = v89
		v87 = v87 + v94
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L27
L35:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v110 != 0 {
		v108 = v108 + int32(1)
		goto L35
	} else {
		goto L37
	}
L36:
	;
	goto L27
L37:
	;
	goto L36
L38:
	;
	if l5 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L39:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v65)+88))
	goto L41
L40:
	;
	v162 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = v162
	*(*int64)(unsafe.Add(mBase, uint32(l4+int32(32)))) = v162
	*(*int64)(unsafe.Add(mBase, uint32(l4+int32(24)))) = v162
	*(*int64)(unsafe.Add(mBase, uint32(l4+int32(16)))) = v162
	*(*int64)(unsafe.Add(mBase, uint32(l4+int32(8)))) = v162
	goto L38
L41:
	;
	if v123&int32(2) == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v128 = F_clusterNodeGetPrimary(m, v65)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	if v128 == int32(0) {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v132 = F_clusterNodeGetPrimary(m, v65)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	goto L46
L46:
	;
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v132+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = v136
	v138 = int32(32)
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v132+int32(40))))
	*(*int64)(unsafe.Add(mBase, uint32(l4+v138))) = v142
	v144 = int32(24)
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v132+v138)))
	*(*int64)(unsafe.Add(mBase, uint32(l4+v144))) = v148
	v150 = int32(16)
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v132+v144)))
	*(*int64)(unsafe.Add(mBase, uint32(l4+v150))) = v154
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v132+v150)))
	*(*int64)(unsafe.Add(mBase, uint32(l4+int32(8)))) = v160
	goto L38
L47:
	;
	if l6 != 0 {
		goto L61
	} else {
		goto L62
	}
L48:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	if v184 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v203
	goto L47
L50:
	;
	v203 = v201
	goto L49
L51:
	;
	if l1 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v65)+2328))
	v203 = v195
	goto L49
L53:
	;
	if v184 == int32(0) {
		goto L51
	} else {
		goto L57
	}
L54:
	;
	if l1 == int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v65)+2340))
	if v189 == int32(0) {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v201 = v189
	goto L50
L57:
	;
	goto L52
L58:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v65)+2324))
	v201 = v200
	goto L50
L59:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v65)+2336))
	if v198 != 0 {
		v201 = v198
		goto L50
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0)
	v209 = F_getMyClusterNode(m)
	mBase = m.M
	goto L64
L62:
	;
	return int32(0)
L63:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v65)+88))
	goto L67
L64:
	;
	if base.B2i32(v65 == v209) == int32(0) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v213 | int32(1)
	goto L63
L66:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v65)+88))
	goto L70
L67:
	;
	if v217&int32(1) == int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v222 | int32(2)
	goto L66
L69:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v65)+88))
	goto L73
L70:
	;
	if v226&int32(2) == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v231 | int32(4)
	goto L69
L72:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v65)+88))
	goto L76
L73:
	;
	if v235&int32(4) == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v240 | int32(8)
	goto L72
L75:
	;
	v253 = int32(0)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v65)+88))
	goto L78
L76:
	;
	if v244&int32(8) == int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v249 | int32(16)
	goto L75
L78:
	;
	if v254&int32(512) == int32(0) {
		v263 = v253
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v259 | int32(32)
	v263 = v253
	goto L1
}
func F_moduleGetFreeEffort(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_objectGetVal(m, l1)
	mBase = m.M
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	if v13 == int32(0) {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
		if v27 != 0 {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v30 = m.T0[v27].(func(*base.Module, int32, int32) int32)(m, l0, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = v30
				m.G0 = v9 + int32(16)
				return v32
			}
		} else {
			v32 = int32(1)
			m.G0 = v9 + int32(16)
			return v32
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v23 = m.T0[v13].(func(*base.Module, int32, int32) int32)(m, v9, v22)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v32 = v23
			m.G0 = v9 + int32(16)
			return v32
		}
	}
}
func F_moduleInitModulesSystem(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = F_listCreate(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[440])) = v8
		v12 = F_listCreate(m)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[454])) = v12
			v16 = F_dictCreate(m, int32(_a831))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = int32(_a20)
				v19 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[414])) = v19
				*(*int32)(unsafe.Add(mBase, _consts[248])) = v16
				v25 = F_dictCreate(m, int32(_a832))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[81])) = v25
					v29 = F_listCreate(m)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[455])) = v29
						v33 = F_listCreate(m)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[441])) = v33
							v37 = F_listCreate(m)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[442])) = v37
								v41 = F_listCreate(m)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[438])) = v41
									F_moduleRegisterCoreAPI(m)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										v48 = int32(_a833)
										v53 = m.G0
										v55 = v53 - int32(64)
										m.G0 = v55
										v58 = F_pipe(m, v48)
										mBase = m.M
										if v58 != 0 {
											v115 = int32(-1)
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, _consts[456]))
											*(*int32)(unsafe.Add(mBase, uint32(v55)+48)) = int32(1)
											v69 = F_fcntl(m, v63, int32(2), v55+int32(48))
											mBase = m.M
											if v69 != 0 {
												v108 = *(*int32)(unsafe.Add(mBase, _consts[456]))
												v109 = F_close(m, v108)
												mBase = m.M
												v110 = *(*int32)(unsafe.Add(mBase, _consts[415]))
												v111 = F_close(m, v110)
												mBase = m.M
												v115 = int32(-1)
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, _consts[415]))
												*(*int32)(unsafe.Add(mBase, uint32(v55)+32)) = int32(1)
												v81 = F_fcntl(m, v75, int32(2), v55+int32(32))
												mBase = m.M
												if v81 != 0 {
													v108 = *(*int32)(unsafe.Add(mBase, _consts[456]))
													v109 = F_close(m, v108)
													mBase = m.M
													v110 = *(*int32)(unsafe.Add(mBase, _consts[415]))
													v111 = F_close(m, v110)
													mBase = m.M
													v115 = int32(-1)
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, _consts[456]))
													*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = int32(2048)
													v92 = F_fcntl(m, v87, int32(4), v55+int32(16))
													mBase = m.M
													if v92 != 0 {
														v108 = *(*int32)(unsafe.Add(mBase, _consts[456]))
														v109 = F_close(m, v108)
														mBase = m.M
														v110 = *(*int32)(unsafe.Add(mBase, _consts[415]))
														v111 = F_close(m, v110)
														mBase = m.M
														v115 = int32(-1)
													} else {
														v99 = *(*int32)(unsafe.Add(mBase, _consts[415]))
														*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(2048)
														v102 = F_fcntl(m, v99, int32(4), v55)
														mBase = m.M
														if v102 == int32(0) {
															v115 = int32(0)
														} else {
															v108 = *(*int32)(unsafe.Add(mBase, _consts[456]))
															v109 = F_close(m, v108)
															mBase = m.M
															v110 = *(*int32)(unsafe.Add(mBase, _consts[415]))
															v111 = F_close(m, v110)
															mBase = m.M
															v115 = int32(-1)
														}
													}
												}
											}
										}
										m.G0 = v55 + int32(64)
										if v115 != int32(-1) {
											v136 = F_raxNew(m)
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[457])) = v136
												v140 = F_listCreate(m)
												mBase = m.M
												v141 = m.ExcPending
												if v141 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[449])) = v140
													m.G0 = v5 + int32(16)
													return
												}
											}
										} else {
											v122 = *(*int32)(unsafe.Add(mBase, _consts[28]))
											if int32(3) < v122 {
												m.Env.Exit(m, int32(1))
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, _consts[18]))
												v127 = F___strerror_l(m, v126, v126)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v5))) = v127
												F__serverLog(m, int32(3), int32(_a834), v5)
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return
												} else {
													m.Env.Exit(m, int32(1))
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
			}
		}
	}
}
func F_moduleInitPostOnLoadResolved(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int64
	_ = v105
	var v109 int64
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v131 int64
	_ = v131
	var v135 int64
	_ = v135
	var v138 int32
	_ = v138
	var v146 int64
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v383 int32
	_ = v383
	v8 = int32(0)
	v11 = int64(0)
	v12 = m.G0
	v14 = v12 - int32(128)
	m.G0 = v14
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(80)))) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(112)))) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(104)))) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(96)))) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(88)))) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(72)))) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = int32(64)
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(120)))) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = int32(561)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v8
	v55 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	if v55 == v8 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v88
	v92 = *(*int32)(unsafe.Add(mBase, _consts[384]))
	v93 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v95 = m.T0[v94].(func(*base.Module) int64)(m)
	mBase = m.M
	if v92 == v93 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v74 = F_createClient(m, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v58 = int32(0)
	v60 = v55 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[383])) = v60
	v63 = *(*int32)(unsafe.Add(mBase, _consts[385]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v60<<(uint(int32(2))%32))))
	v69 = *(*int32)(unsafe.Add(mBase, _consts[386]))
	if base.Ui32(v69) <= base.Ui32(v60) {
		v88 = v67
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[386])) = v60
	v88 = v67
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+328)) = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v74)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+200)) = v80 | int32(1073741824)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v74)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+204)) = v84 | int32(268435456)
	v88 = v74
	goto L1
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+112)) = v109
	v113 = int32(0)
	v117 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	*(*int32)(unsafe.Add(mBase, _consts[177])) = v117 + int32(1)
	goto L12
L8:
	;
	v105 = *(*int64)(unsafe.Add(mBase, _consts[387]))
	v109 = v105*int64(1000) + v95
	goto L7
L9:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v101 = base.I32_div_s(int32(1000000), v100)
	v109 = v95 + base.I64_extend_i32_s(v101)
	goto L7
L10:
	;
	v152 = m.T0[l0].(func(*base.Module, int32, int32, int32) int32)(m, v14+int32(56), l3, l4)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	if v117 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	goto L15
L14:
	;
	v127 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[178])) = v125
	v131 = base.I64_div_s(v125, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[54])) = v131
	v135 = base.I64_div_s(v125, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[109])) = v135
	v138 = *(*int32)(unsafe.Add(mBase, _consts[179]))
	F_lrulfu_updateClockAndPolicy(m, v131, int32(base.Ui32(v138&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v146 = *(*int64)(unsafe.Add(mBase, _consts[54]))
	*(*int64)(unsafe.Add(mBase, _consts[12])) = v146
	goto L11
L15:
	;
	v125 = F_ustime(m)
	mBase = m.M
	goto L14
L16:
	;
	m.G0 = v14 + int32(128)
	return v383
L17:
	;
	if l1 == int32(0) {
		v215 = l1
		goto L42
	} else {
		goto L43
	}
L18:
	;
	if v152 != int32(1) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if v158 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_moduleFreeContext(m, v14+int32(56))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L5
	} else {
		goto L39
	}
L21:
	;
	if int32(3) < v157 {
		goto L20
	} else {
		goto L34
	}
L22:
	;
	if int32(3) < v157 {
		v175 = v158
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_moduleUnregisterCleanup(m, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L5
	} else {
		goto L29
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l2
	if l6 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v166 = int32(_a1589)
	goto L27
L26:
	;
	v166 = int32(_a320)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v166
	F__serverLog(m, int32(3), int32(_a1590), v14+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v175 = v174
	goto L23
L29:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+76))
	if v179 == int32(0) {
		v185 = v178
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_moduleFreeModuleStructure(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	F_ACLCleanupCategoriesOnFailure(m, v179)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v185 = v184
	goto L30
L33:
	;
	goto L20
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l2
	if l6 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v193 = int32(_a1589)
	goto L37
L36:
	;
	v193 = int32(_a320)
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v193
	F__serverLog(m, int32(3), int32(_a1591), v14)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	goto L20
L39:
	;
	v205 = int32(-1)
	if l1 == int32(0) {
		v383 = v205
		goto L16
	} else {
		goto L40
	}
L40:
	;
	goto L41
L41:
	;
	v383 = v205
	goto L16
L42:
	;
	v216 = int32(0)
	v218 = *(*int32)(unsafe.Add(mBase, _consts[81]))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v221 = F_dictAdd(m, v218, v220, v219)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L5
	} else {
		goto L46
	}
L43:
	;
	if l6 == int32(0) {
		v215 = l1
		goto L42
	} else {
		goto L44
	}
L44:
	;
	goto L45
L45:
	;
	v215 = int32(0)
	goto L42
L46:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+80)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v223)+52)) = int32(0)
	v229 = F_valkey_malloc(m, int32(12))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+64)) = v229
	v233 = F_sdsnew(m, l2)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v233
	if l4 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v278)+68))
	if v288 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L50:
	;
	v243 = F_valkey_malloc(m, l4<<(uint(int32(2))%32))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L5
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236)+4)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v236)+8)) = int32(0)
	v278 = v235
	goto L49
L52:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v246)+8)) = v243
	if l4 <= int32(0) {
		v278 = v245
		goto L49
	} else {
		goto L53
	}
L53:
	;
	v259 = v216
	goto L54
L54:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+64))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	v266 = v259 << (uint(int32(2)) % 32)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l3+v266)))
	*(*int32)(unsafe.Add(mBase, uint32(v264+v266))) = v269
	F_incrRefCount(m, v269)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L5
	} else {
		goto L56
	}
L55:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v278 = v276
	goto L49
L56:
	;
	v274 = v259 + int32(1)
	if v274 != l4 {
		v259 = v274
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if l6 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	F_ACLRecomputeCommandBitsFromCommandRulesAllUsers(m)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v321 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v320)+72)) = v321
	v323 = int32(1)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v320)+32))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+20))
	if v325 == v321 {
		v334 = v323
		goto L68
	} else {
		goto L69
	}
L62:
	;
	if int32(2) < v294 {
		goto L61
	} else {
		goto L66
	}
L63:
	;
	if int32(2) < v294 {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v300
	F__serverLog(m, int32(2), int32(_a1592), v14+int32(32))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = l2
	F__serverLog(m, int32(2), int32(_a1593), v14+int32(48))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	goto L61
L68:
	;
	if l5 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L69:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v320)+36))
	if v328 != 0 {
		v334 = v323
		goto L68
	} else {
		goto L70
	}
L70:
	;
	F_serverLogRaw(m, int32(3), int32(_a1594))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	v334 = int32(0)
	goto L68
L72:
	;
	v366 = int32(0)
	F_moduleFireServerEvent(m, int64(9), v366, v350)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L5
	} else {
		goto L82
	}
L73:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _consts[81]))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	v356 = F_dictFetchValue(m, v354, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L5
	} else {
		goto L79
	}
L74:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if v334 != 0 {
		goto L72
	} else {
		goto L78
	}
L75:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _consts[248]))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v338)+16))
	if v339 == int32(0)-v341 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	F_serverLogRaw(m, int32(3), int32(_a1595))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v351 = v348
	goto L73
L78:
	;
	v351 = v350
	goto L73
L79:
	;
	v359 = F_moduleUnloadInternal(m, v356, int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	F_moduleFreeContext(m, v14+int32(56))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	v383 = int32(-1)
	goto L16
L82:
	;
	F_moduleFreeContext(m, v14+int32(56))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	v383 = v366
	goto L16
}
func F_moduleLoadFromQueue(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v164 int64
	_ = v164
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v279 int64
	_ = v279
	var v280 int64
	_ = v280
	var v281 int64
	_ = v281
	var v282 int64
	_ = v282
	var v283 int64
	_ = v283
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v293 int64
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	v12 = v7 + int32(24)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13
	goto L1
L1:
	;
	v18 = v7 + int32(24)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v20 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[248]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	if v115 == int32(0)-v117 {
		goto L30
	} else {
		goto L31
	}
L3:
	;
	if v20 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20+base.B2i32(v23 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v29
	goto L4
L6:
	;
	v34 = v20
	goto L7
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v42 = F_moduleLoad(m, v38, v39, v40, int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L2
L9:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	F_sdsfree(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L16
	}
L10:
	;
	return
L11:
	;
	if v42 != int32(-1) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v47 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v50
	F__serverLog(m, int32(3), int32(_a1586), v7+int32(16))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v64 < int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	F_valkey_free(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L10
	} else {
		goto L23
	}
L18:
	;
	v70 = int32(0)
	goto L19
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v70<<(uint(int32(2))%32))))
	F_decrRefCount(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L10
	} else {
		goto L21
	}
L20:
	;
	goto L17
L21:
	;
	v79 = v70 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v79 < v80 {
		v70 = v79
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	F_valkey_free(m, v37)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	F_listDelNode(m, v92, v34)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v96 = v7 + int32(24)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	if v98 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v98 != 0 {
		v34 = v98
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v98+base.B2i32(v101 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107
	goto L27
L29:
	;
	goto L8
L30:
	;
	m.G0 = v7 + int32(32)
	return
L31:
	;
	v120 = F_dictGetSafeIterator(m, v114)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L33
	}
L32:
	;
	F_dictReleaseIterator(m, v120)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L10
	} else {
		goto L94
	}
L33:
	;
	v129 = v120 + int32(20)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	if v130 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	if v225 == int32(0) {
		goto L32
	} else {
		goto L60
	}
L35:
	;
	v136 = v129
	v137 = v133
	goto L38
L36:
	;
	v133 = int32(1)
	goto L35
L37:
	;
	v133 = int32(0)
	goto L35
L38:
	;
	switch v137 {
	case 0:
		goto L43
	default:
		goto L42
	}
L40:
	;
	v137 = int32(0)
	goto L38
L41:
	;
	goto L34
L42:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+16)) = v217
	if v217 == int32(0) {
		goto L40
	} else {
		goto L59
	}
L43:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v141 != int32(-1) {
		v180 = v141
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v181 = int32(1)
	v182 = v180 + v181
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v182
	v184 = int32(0)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+v188+int32(26)))))
	if v192 == int32(255) {
		goto L53
	} else {
		goto L54
	}
L45:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	if v145 != 0 {
		v180 = int32(-1)
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	if v147 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+20))
	if v174 != int32(-1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v154 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v146)+16)))
	v155 = int64(*(*int8)(unsafe.Add(mBase, uint32(v146)+27)))
	v156 = int64(*(*int32)(unsafe.Add(mBase, uint32(v146)+8)))
	v157 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v146)+12)))
	v158 = int64(*(*int8)(unsafe.Add(mBase, uint32(v146)+26)))
	v159 = int64(*(*int32)(unsafe.Add(mBase, uint32(v146)+4)))
	v160 = F_wangHash64(m, v159)
	mBase = m.M
	v162 = F_wangHash64(m, v158+v160)
	mBase = m.M
	v164 = F_wangHash64(m, v157+v162)
	mBase = m.M
	v166 = F_wangHash64(m, v156+v164)
	mBase = m.M
	v168 = F_wangHash64(m, v155+v166)
	mBase = m.M
	v170 = F_wangHash64(m, v154+v168)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v120)+24)) = v170
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v173 = v172
	goto L47
L49:
	;
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+24)))
	v152 = v150 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v146)+24)) = uint16(v152)
	v173 = v146
	goto L47
L50:
	;
	v180 = v174 + int32(-1)
	goto L44
L51:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v180 = v177
	goto L44
L52:
	;
	v207 = int32(2)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v187+v205<<(uint(v207)%32)+int32(4))))
	v136 = v212 + v206<<(uint(v207)%32)
	v137 = int32(1)
	goto L38
L53:
	;
	v196 = v184
	goto L55
L54:
	;
	v196 = v181 << (uint(v192) % 32)
	goto L55
L55:
	;
	if v182 < v196 {
		v205 = v188
		v206 = v182
		goto L52
	} else {
		goto L56
	}
L56:
	;
	if v188 != 0 {
		v225 = v184
		goto L41
	} else {
		goto L57
	}
L57:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
	if v198 == int32(-1) {
		v225 = v184
		goto L41
	} else {
		goto L58
	}
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v120)+4)) = int64(4294967296)
	v205 = int32(1)
	v206 = int32(0)
	goto L52
L59:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v221
	v225 = v217
	goto L41
L60:
	;
	v234 = v225
	goto L61
L61:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	goto L63
L62:
	;
	goto L32
L63:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v237 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v252 = v120 + int32(20)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	if v253 != 0 {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v235
	F__serverLog(m, int32(3), int32(_a1587), v7)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	if v348 != 0 {
		v234 = v348
		goto L61
	} else {
		goto L93
	}
L68:
	;
	v259 = v252
	v260 = v256
	goto L71
L69:
	;
	v256 = int32(1)
	goto L68
L70:
	;
	v256 = int32(0)
	goto L68
L71:
	;
	switch v260 {
	case 0:
		goto L76
	default:
		goto L75
	}
L73:
	;
	v260 = int32(0)
	goto L71
L74:
	;
	goto L67
L75:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+16)) = v340
	if v340 == int32(0) {
		goto L73
	} else {
		goto L92
	}
L76:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v264 != int32(-1) {
		v303 = v264
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v304 = int32(1)
	v305 = v303 + v304
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v305
	v307 = int32(0)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310+v311+int32(26)))))
	if v315 == int32(255) {
		goto L86
	} else {
		goto L87
	}
L78:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	if v268 != 0 {
		v303 = int32(-1)
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	if v270 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+20))
	if v297 != int32(-1) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v277 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v269)+16)))
	v278 = int64(*(*int8)(unsafe.Add(mBase, uint32(v269)+27)))
	v279 = int64(*(*int32)(unsafe.Add(mBase, uint32(v269)+8)))
	v280 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v269)+12)))
	v281 = int64(*(*int8)(unsafe.Add(mBase, uint32(v269)+26)))
	v282 = int64(*(*int32)(unsafe.Add(mBase, uint32(v269)+4)))
	v283 = F_wangHash64(m, v282)
	mBase = m.M
	v285 = F_wangHash64(m, v281+v283)
	mBase = m.M
	v287 = F_wangHash64(m, v280+v285)
	mBase = m.M
	v289 = F_wangHash64(m, v279+v287)
	mBase = m.M
	v291 = F_wangHash64(m, v278+v289)
	mBase = m.M
	v293 = F_wangHash64(m, v277+v291)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v120)+24)) = v293
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v296 = v295
	goto L80
L82:
	;
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269)+24)))
	v275 = v273 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v269)+24)) = uint16(v275)
	v296 = v269
	goto L80
L83:
	;
	v303 = v297 + int32(-1)
	goto L77
L84:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v303 = v300
	goto L77
L85:
	;
	v330 = int32(2)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v310+v328<<(uint(v330)%32)+int32(4))))
	v259 = v335 + v329<<(uint(v330)%32)
	v260 = int32(1)
	goto L71
L86:
	;
	v319 = v307
	goto L88
L87:
	;
	v319 = v304 << (uint(v315) % 32)
	goto L88
L88:
	;
	if v305 < v319 {
		v328 = v311
		v329 = v305
		goto L85
	} else {
		goto L89
	}
L89:
	;
	if v311 != 0 {
		v348 = v307
		goto L74
	} else {
		goto L90
	}
L90:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v310)+20))
	if v321 == int32(-1) {
		v348 = v307
		goto L74
	} else {
		goto L91
	}
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v120)+4)) = int64(4294967296)
	v328 = int32(1)
	v329 = int32(0)
	goto L85
L92:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v340)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v344
	v348 = v340
	goto L74
L93:
	;
	goto L62
L94:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v359 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F__serverLog(m, int32(3), int32(_a1588), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L10
	} else {
		goto L97
	}
L97:
	;
	goto L95
}
func F_moduleLoadQueueEntryToLoadmoduleOptionStr(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	v5 = F_sdsnew(m, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = F_sdscatlen(m, v5, int32(_a6), int32(1))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v15 = F_sdscatsds(m, v11, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v17 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v19 <= v17 {
		v46 = v15
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return v46
L6:
	;
	v23 = v17
	v24 = v15
	goto L7
L7:
	;
	v28 = F_sdscatlen(m, v24, int32(_a6), int32(1))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v46 = v37
	goto L5
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v23<<(uint(int32(2))%32))))
	v36 = F_objectGetVal(m, v35)
	mBase = m.M
	v37 = F_sdscatsds(m, v28, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v40 = v23 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v40 < v42 {
		v23 = v40
		v24 = v37
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
}
func F_moduleNameFromCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v2 == int32(562) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		return v15
	} else {
		F__serverAssert(m, int32(_a815), int32(_a756), int32(7300))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
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
func F_moduleReleaseGIL(m *base.Module) {
	return
}
func F_moduleScriptingEngineInitContext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	if l2 == int32(0) {
		if l2 != 0 {
			v11 = int32(2048)
		} else {
			v11 = int32(80)
		}
		F_moduleCreateContext(m, l0, l1, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if l2 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l3
			}
			return
		}
	} else {
		if l3 == int32(0) {
			F__serverAssert(m, int32(_a759), int32(_a756), int32(1022))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if l2 != 0 {
				v11 = int32(2048)
			} else {
				v11 = int32(80)
			}
			F_moduleCreateContext(m, l0, l1, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if l2 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l3
				}
				return
			}
		}
	}
}
func F_moduleTypeEncodeId(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int64
	_ = v167
	var v171 int64
	_ = v171
	var v212 int64
	_ = v212
	v3 = int32(0)
	v4 = int64(0)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[439]))
	if l0&int32(3) == v3 {
		v37 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if base.Ui32(int32(1023)) < base.Ui32(l1) {
		v212 = v4
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v70 = v62 - l0
	goto L1
L3:
	;
	v41 = v37
	goto L11
L4:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = l0
	goto L7
L6:
	;
	v70 = l0 - l0
	goto L1
L7:
	;
	v30 = v26 + int32(1)
	if v30&int32(3) == int32(0) {
		v37 = v30
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v35 != 0 {
		v26 = v30
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v62 = v30
	goto L2
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v50 = int32(-2139062144)
	if (int32(16843008)-v47|v47)&v50 == v50 {
		v41 = v41 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v56 = v41
	goto L14
L13:
	;
	goto L12
L14:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v60 != 0 {
		v56 = v56 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v62 = v56
	goto L2
L16:
	;
	goto L15
L17:
	;
	return v212
L18:
	;
	if v70 != int32(9) {
		v212 = v4
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v75 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	v76 = F___strchrnul(m, v14, v75)
	mBase = m.M
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v78 == v75&int32(255) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v82 == int32(0) {
		v212 = v4
		goto L17
	} else {
		goto L24
	}
L21:
	;
	v82 = v76
	goto L23
L22:
	;
	v82 = int32(0)
	goto L23
L23:
	;
	goto L20
L24:
	;
	v85 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
	v86 = F___strchrnul(m, v14, v85)
	mBase = m.M
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v88 == v85&int32(255) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v92 == int32(0) {
		v212 = v4
		goto L17
	} else {
		goto L29
	}
L26:
	;
	v92 = v86
	goto L28
L27:
	;
	v92 = int32(0)
	goto L28
L28:
	;
	goto L25
L29:
	;
	v95 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
	v96 = F___strchrnul(m, v14, v95)
	mBase = m.M
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v98 == v95&int32(255) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v102 == int32(0) {
		v212 = v4
		goto L17
	} else {
		goto L34
	}
L31:
	;
	v102 = v96
	goto L33
L32:
	;
	v102 = int32(0)
	goto L33
L33:
	;
	goto L30
L34:
	;
	v105 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+3)))
	v106 = F___strchrnul(m, v14, v105)
	mBase = m.M
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v108 == v105&int32(255) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v112 == int32(0) {
		v212 = v4
		goto L17
	} else {
		goto L39
	}
L36:
	;
	v112 = v106
	goto L38
L37:
	;
	v112 = int32(0)
	goto L38
L38:
	;
	goto L35
L39:
	;
	v115 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+4)))
	v116 = F___strchrnul(m, v14, v115)
	mBase = m.M
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v118 == v115&int32(255) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v122 == int32(0) {
		v212 = v4
		goto L17
	} else {
		goto L44
	}
L41:
	;
	v122 = v116
	goto L43
L42:
	;
	v122 = int32(0)
	goto L43
L43:
	;
	goto L40
L44:
	;
	v125 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+5)))
	v126 = F___strchrnul(m, v14, v125)
	mBase = m.M
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v128 == v125&int32(255) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v132 == int32(0) {
		v212 = v4
		goto L17
	} else {
		goto L49
	}
L46:
	;
	v132 = v126
	goto L48
L47:
	;
	v132 = int32(0)
	goto L48
L48:
	;
	goto L45
L49:
	;
	v135 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+6)))
	v136 = F___strchrnul(m, v14, v135)
	mBase = m.M
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v138 == v135&int32(255) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v142 == int32(0) {
		v212 = v4
		goto L17
	} else {
		goto L54
	}
L51:
	;
	v142 = v136
	goto L53
L52:
	;
	v142 = int32(0)
	goto L53
L53:
	;
	goto L50
L54:
	;
	v145 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+7)))
	v146 = F___strchrnul(m, v14, v145)
	mBase = m.M
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v148 == v145&int32(255) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v152 == int32(0) {
		v212 = v4
		goto L17
	} else {
		goto L59
	}
L56:
	;
	v152 = v146
	goto L58
L57:
	;
	v152 = int32(0)
	goto L58
L58:
	;
	goto L55
L59:
	;
	v155 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+8)))
	v156 = F___strchrnul(m, v14, v155)
	mBase = m.M
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v158 == v155&int32(255) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v162 == int32(0) {
		v212 = v4
		goto L17
	} else {
		goto L64
	}
L61:
	;
	v162 = v156
	goto L63
L62:
	;
	v162 = int32(0)
	goto L63
L63:
	;
	goto L60
L64:
	;
	v167 = int64(12)
	v171 = int64(6)
	v212 = ((((base.I64_extend_i32_u(v82-v14)<<(uint(v167)%64)|base.I64_extend_i32_u(v92-v14)<<(uint(v171)%64)|base.I64_extend_i32_u(v102-v14))<<(uint(v167)%64)|base.I64_extend_i32_u(v112-v14)<<(uint(v171)%64)|base.I64_extend_i32_u(v122-v14))<<(uint(v167)%64)|base.I64_extend_i32_u(v132-v14)<<(uint(v171)%64)|base.I64_extend_i32_u(v142-v14))<<(uint(v167)%64)|base.I64_extend_i32_u(v152-v14)<<(uint(v171)%64)|base.I64_extend_i32_u(v162-v14))<<(uint(int64(10))%64) | base.I64_extend_i32_u(l1)
	goto L17
}
func F_moduleTypeLookupModuleByNameIgnoreCase(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_moduleTypeLookupModuleByNameInternal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_moduleTypeModuleName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v2 = int32(0)
	if l0 == v2 {
		v11 = v2
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v6 == int32(0) {
			v11 = v2
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v11 = v9
		}
	}
	return v11
}
func F_moduleUnloadAllModules(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v193 int64
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[81]))
	v10 = F_dictGetSafeIterator(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v10)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L2
	} else {
		goto L66
	}
L2:
	;
	return
L3:
	;
	v19 = v10 + int32(20)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v115 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v26 = v19
	v27 = v23
	goto L8
L6:
	;
	v23 = int32(1)
	goto L5
L7:
	;
	v23 = int32(0)
	goto L5
L8:
	;
	switch v27 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v27 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v107
	if v107 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v31 != int32(-1) {
		v70 = v31
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = int32(1)
	v72 = v70 + v71
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v72
	v74 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78+int32(26)))))
	if v82 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v35 != 0 {
		v70 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v37 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	if v64 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v44 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v36)+16)))
	v45 = int64(*(*int8)(unsafe.Add(mBase, uint32(v36)+27)))
	v46 = int64(*(*int32)(unsafe.Add(mBase, uint32(v36)+8)))
	v47 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v36)+12)))
	v48 = int64(*(*int8)(unsafe.Add(mBase, uint32(v36)+26)))
	v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(v36)+4)))
	v50 = F_wangHash64(m, v49)
	mBase = m.M
	v52 = F_wangHash64(m, v48+v50)
	mBase = m.M
	v54 = F_wangHash64(m, v47+v52)
	mBase = m.M
	v56 = F_wangHash64(m, v46+v54)
	mBase = m.M
	v58 = F_wangHash64(m, v45+v56)
	mBase = m.M
	v60 = F_wangHash64(m, v44+v58)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v63 = v62
	goto L17
L19:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+24)))
	v42 = v40 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+24)) = uint16(v42)
	v63 = v36
	goto L17
L20:
	;
	v70 = v64 + int32(-1)
	goto L14
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v70 = v67
	goto L14
L22:
	;
	v97 = int32(2)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v77+v95<<(uint(v97)%32)+int32(4))))
	v26 = v102 + v96<<(uint(v97)%32)
	v27 = int32(1)
	goto L8
L23:
	;
	v86 = v74
	goto L25
L24:
	;
	v86 = v71 << (uint(v82) % 32)
	goto L25
L25:
	;
	if v72 < v86 {
		v95 = v78
		v96 = v72
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v78 != 0 {
		v115 = v74
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	if v88 == int32(-1) {
		v115 = v74
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(4294967296)
	v95 = int32(1)
	v96 = int32(0)
	goto L22
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v111
	v115 = v107
	goto L11
L30:
	;
	v123 = v115
	goto L31
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	goto L33
L32:
	;
	goto L1
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
	v129 = F_moduleUnloadInternal(m, v124, v6+int32(12))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L2
	} else {
		goto L35
	}
L34:
	;
	v152 = v10 + int32(20)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v153 != 0 {
		goto L41
	} else {
		goto L42
	}
L35:
	;
	if v129 != int32(-1) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v134 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v139
	F__serverLog(m, int32(3), int32(_a1614), v6)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	if v248 != 0 {
		v123 = v248
		goto L31
	} else {
		goto L65
	}
L40:
	;
	v159 = v152
	v160 = v156
	goto L43
L41:
	;
	v156 = int32(1)
	goto L40
L42:
	;
	v156 = int32(0)
	goto L40
L43:
	;
	switch v160 {
	case 0:
		goto L48
	default:
		goto L47
	}
L45:
	;
	v160 = int32(0)
	goto L43
L46:
	;
	goto L39
L47:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v240
	if v240 == int32(0) {
		goto L45
	} else {
		goto L64
	}
L48:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v164 != int32(-1) {
		v203 = v164
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v204 = int32(1)
	v205 = v203 + v204
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v205
	v207 = int32(0)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210+v211+int32(26)))))
	if v215 == int32(255) {
		goto L58
	} else {
		goto L59
	}
L50:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v168 != 0 {
		v203 = int32(-1)
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v170 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+20))
	if v197 != int32(-1) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v177 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v169)+16)))
	v178 = int64(*(*int8)(unsafe.Add(mBase, uint32(v169)+27)))
	v179 = int64(*(*int32)(unsafe.Add(mBase, uint32(v169)+8)))
	v180 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v169)+12)))
	v181 = int64(*(*int8)(unsafe.Add(mBase, uint32(v169)+26)))
	v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v169)+4)))
	v183 = F_wangHash64(m, v182)
	mBase = m.M
	v185 = F_wangHash64(m, v181+v183)
	mBase = m.M
	v187 = F_wangHash64(m, v180+v185)
	mBase = m.M
	v189 = F_wangHash64(m, v179+v187)
	mBase = m.M
	v191 = F_wangHash64(m, v178+v189)
	mBase = m.M
	v193 = F_wangHash64(m, v177+v191)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v193
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v196 = v195
	goto L52
L54:
	;
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+24)))
	v175 = v173 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v169)+24)) = uint16(v175)
	v196 = v169
	goto L52
L55:
	;
	v203 = v197 + int32(-1)
	goto L49
L56:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v203 = v200
	goto L49
L57:
	;
	v230 = int32(2)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v210+v228<<(uint(v230)%32)+int32(4))))
	v159 = v235 + v229<<(uint(v230)%32)
	v160 = int32(1)
	goto L43
L58:
	;
	v219 = v207
	goto L60
L59:
	;
	v219 = v204 << (uint(v215) % 32)
	goto L60
L60:
	;
	if v205 < v219 {
		v228 = v211
		v229 = v205
		goto L57
	} else {
		goto L61
	}
L61:
	;
	if v211 != 0 {
		v248 = v207
		goto L46
	} else {
		goto L62
	}
L62:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v210)+20))
	if v221 == int32(-1) {
		v248 = v207
		goto L46
	} else {
		goto L63
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(4294967296)
	v228 = int32(1)
	v229 = int32(0)
	goto L57
L64:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v244
	v248 = v240
	goto L46
L65:
	;
	goto L32
L66:
	;
	m.G0 = v6 + int32(16)
	return
}
func F_moduleUnloadInternal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v29 int32
	_ = v29
	var v35 int64
	_ = v35
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	v5 = m.G0
	v7 = v5 - int32(400)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v10 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v7 + int32(400)
	return v283
L2:
	;
	F_moduleUnregisterCleanup(m, l0)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L11
	} else {
		goto L69
	}
L3:
	;
	v283 = int32(-1)
	goto L1
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v16 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a1600)
	goto L3
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v21 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a1601)
	goto L3
L8:
	;
	v27 = v7 + int32(96)
	v29 = *(*int32)(unsafe.Add(mBase, _consts[457]))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = int32(128)
	v35 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+12)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v27)+296)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v27)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v7 + int32(120)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+156)) = v7 + int32(264)
	goto L10
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a1602)
	goto L3
L10:
	;
	v50 = int32(0)
	v52 = F_raxSeek(m, v7+int32(96), int32(_a67), v50, v50)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	goto L14
L13:
	;
	F_raxStop(m, v7+int32(96))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L11
	} else {
		goto L20
	}
L14:
	;
	v62 = F_raxNext(m, v7+int32(96))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L11
	} else {
		goto L16
	}
L15:
	;
	F_raxStop(m, v7+int32(96))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L19
	}
L16:
	;
	if v62 == int32(0) {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)+108))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v67 != l0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a1603)
	goto L3
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+92)) = int32(0)
	v83 = F_ACLModuleHasCommandRules(m, l0, v7+int32(92))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L11
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+88)) = int32(0)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v113 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L22:
	;
	if v83 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v88 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v7)+92))
	if v104 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+80)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v7)+92))
	if v93 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v95 = v93
	goto L28
L27:
	;
	v95 = int32(_a86)
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+84)) = v95
	F__serverLog(m, int32(3), int32(_a1604), v7+int32(80))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a1605)
	goto L3
L31:
	;
	F_sdsfree(m, v104)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	F_moduleCreateContext(m, v7+int32(96), l0, int32(64))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L11
	} else {
		goto L61
	}
L34:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v7)+88))
	if v185 == int32(0) {
		goto L2
	} else {
		goto L60
	}
L35:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v174 {
		v188 = v149
		goto L33
	} else {
		goto L58
	}
L36:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v154 = F_moduleLoadStaticSymbol(m, v7+int32(88), l0, int32(_a1606), v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L11
	} else {
		goto L52
	}
L37:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v117 = int32(_a1606)
	if v116 != int32(_a1) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v132 != 0 {
		v188 = v132
		goto L33
	} else {
		goto L44
	}
L39:
	;
	v128 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a1607)
	v132 = v128
	goto L38
L40:
	;
	v121 = F_strcmp(m, int32(_a1608), v117)
	mBase = m.M
	if v121 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v125 = F_strcmp(m, int32(_a1609), v117)
	mBase = m.M
	if v125 != 0 {
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v132 = v123
	goto L38
L43:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	v132 = v127
	goto L38
L44:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v134 = int32(_a1610)
	if v133 != int32(_a1) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v149 != 0 {
		goto L35
	} else {
		goto L51
	}
L46:
	;
	v145 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a1607)
	v149 = v145
	goto L45
L47:
	;
	v138 = F_strcmp(m, int32(_a1608), v134)
	mBase = m.M
	if v138 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v142 = F_strcmp(m, int32(_a1609), v134)
	mBase = m.M
	if v142 != 0 {
		goto L46
	} else {
		goto L50
	}
L49:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v149 = v140
	goto L45
L50:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	v149 = v144
	goto L45
L51:
	;
	goto L2
L52:
	;
	if v154 == int32(0) {
		goto L34
	} else {
		goto L53
	}
L53:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v159 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	goto L57
L55:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v162
	F__serverLog(m, int32(3), int32(_a1611), v7+int32(48))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(11)
	goto L3
L58:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v177
	F__serverLog(m, int32(2), int32(_a1612), v7+int32(64))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	v188 = v149
	goto L33
L60:
	;
	v188 = v185
	goto L33
L61:
	;
	v196 = m.T0[v188].(func(*base.Module, int32) int32)(m, v7+int32(96))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	F_moduleFreeContext(m, v7+int32(96))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	if v196 != int32(1) {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v205 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	goto L68
L66:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v208
	F__serverLog(m, int32(3), int32(_a1611), v7+int32(32))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(11)
	goto L3
L69:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v227 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	F_moduleFireServerEvent(m, int64(9), int32(1), l0)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L11
	} else {
		goto L80
	}
L71:
	;
	goto L72
L72:
	;
	goto L70
L80:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v261 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v270 = int32(0)
	v272 = *(*int32)(unsafe.Add(mBase, _consts[81]))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v274 = F_dictDelete(m, v272, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L11
	} else {
		goto L84
	}
L82:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v264
	F__serverLog(m, int32(2), int32(_a1613), v7)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	F_moduleFreeModuleStructure(m, l0)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L11
	} else {
		goto L85
	}
L85:
	;
	F_ACLRecomputeCommandBitsFromCommandRulesAllUsers(m)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L11
	} else {
		goto L86
	}
L86:
	;
	v283 = v270
	goto L1
}
func F_moduleUnregisterFilters(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = v7 + int32(8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v12
	goto L1
L1:
	;
	v16 = int32(0)
	v18 = v7 + int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v20 == v16 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return v71
L3:
	;
	if v20 == int32(0) {
		v71 = v16
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20+base.B2i32(v23 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v29
	goto L4
L6:
	;
	v33 = v20
	v35 = v16
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[438]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v40 = F_listSearchKey(m, v38, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v71 = v52
	goto L2
L9:
	;
	F_valkey_free(m, v39)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L14
	}
L10:
	;
	return int32(0)
L11:
	;
	if v40 == int32(0) {
		v52 = v35
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[438]))
	F_listDelNode(m, v47, v40)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v52 = v35 + int32(1)
	goto L9
L14:
	;
	v56 = v7 + int32(8)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v58 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v58 != 0 {
		v33 = v58
		v35 = v52
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v58+base.B2i32(v61 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v67
	goto L16
L18:
	;
	goto L8
}
