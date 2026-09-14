package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_SHA1Final(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	m.Env.Vkmem_hash_final(m, v3, l0, int32(20))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return
}
func F___sig_is_blocked(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	v6 = l0 + int32(-1)
	if base.Ui32(int32(63)) < base.Ui32(v6) {
		v18 = int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v6)>>(uint(int32(3))%32))&int32(536870908))+uint32(_consts[1240])))
		v18 = int32(base.Ui32(v14)>>(uint(v6)%32)) & int32(1)
	}
	return base.B2i32(v18 != int32(0))
}
func F___sigaction(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	if base.Ui32(l0) < base.Ui32(int32(65)) {
		if l2 == int32(0) {
		} else {
			v13 = int32(140)
			v20 = F__emscripten_memcpy_bulkmem(m, l2, l0*v13+int32(9117184), v13)
			mBase = m.M
		}
		if l1 == int32(0) {
		} else {
			v24 = int32(140)
			v31 = F__emscripten_memcpy_bulkmem(m, l0*v24+int32(9117184), l1, v24)
			mBase = m.M
		}
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(28)
		return int32(-1)
	}
}
func F___small_vsprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F___small_vsnprintf(m, l0, int32(2147483647), l1, l2)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F___stdio_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = l1
	v20 = v17 - v15
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v20
	v22 = v20 + l2
	v24 = v13 + int32(16)
	v25 = int32(2)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v32 = m.Wasi_snapshot_preview1.Fd_write(m, v26, v24, v25, v13+int32(12))
	mBase = m.M
	if v32 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v137
L2:
	;
	v123 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v123
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v128 | int32(32)
	if v120 == int32(2) {
		v137 = v123
		goto L1
	} else {
		goto L26
	}
L3:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v107
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v107 + v110
	v137 = l2
	goto L1
L4:
	;
	if v91 != int32(-1) {
		v118 = v90
		v120 = v92
		goto L2
	} else {
		goto L25
	}
L5:
	;
	v44 = v24
	v46 = v22
	v47 = v25
	goto L10
L6:
	;
	if v37 == int32(0) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v34 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v32
	v37 = int32(-1)
	goto L6
L8:
	;
	v37 = int32(0)
	goto L6
L9:
	;
	v90 = v24
	v91 = v22
	v92 = v25
	goto L4
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v46 == v50 {
		goto L3
	} else {
		goto L12
	}
L11:
	;
	v90 = v58
	v91 = v72
	v92 = v74
	goto L4
L12:
	;
	if int32(-1) < v50 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v55 = base.B2i32(base.Ui32(v54) < base.Ui32(v50))
	v58 = v44 + v55<<(uint(int32(3))%32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if base.Ui32(v54) < base.Ui32(v50) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v118 = v44
	v120 = v47
	goto L2
L15:
	;
	v61 = v54
	goto L17
L16:
	;
	v61 = int32(0)
	goto L17
L17:
	;
	v62 = v50 - v61
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v59 + v62
	if base.Ui32(v54) < base.Ui32(v50) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v67 = int32(12)
	goto L20
L19:
	;
	v67 = int32(4)
	goto L20
L20:
	;
	v68 = v44 + v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v69 - v62
	v72 = v46 - v50
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v74 = v47 - v55
	v77 = m.Wasi_snapshot_preview1.Fd_write(m, v73, v58, v74, v13+int32(12))
	mBase = m.M
	if v77 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v82 == int32(0) {
		v44 = v58
		v46 = v72
		v47 = v74
		goto L10
	} else {
		goto L24
	}
L22:
	;
	v79 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v77
	v82 = int32(-1)
	goto L21
L23:
	;
	v82 = int32(0)
	goto L21
L24:
	;
	goto L11
L25:
	;
	goto L3
L26:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v137 = l2 - v134
	goto L1
}
func F___stpcpy(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	if (l1^l0)&int32(3) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v77
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v60)
	if v60&int32(255) == int32(0) {
		v77 = v58
		goto L1
	} else {
		goto L15
	}
L3:
	;
	if l1&int32(3) == int32(0) {
		v29 = l0
		v30 = l1
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v58 = l0
	v59 = l1
	v60 = v10
	goto L2
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v37 = int32(-2139062144)
	if (int32(16843008)-v34|v34)&v37 != v37 {
		v58 = v29
		v59 = v30
		v60 = v34
		goto L2
	} else {
		goto L11
	}
L6:
	;
	v15 = l0
	v16 = l1
	goto L7
L7:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v19)
	if v19 == int32(0) {
		v77 = v15
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v29 = v24
	v30 = v26
	goto L5
L9:
	;
	v23 = int32(1)
	v24 = v15 + v23
	v26 = v16 + v23
	if v26&int32(3) != 0 {
		v15 = v24
		v16 = v26
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v41 = v29
	v42 = v30
	v43 = v34
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v43
	v46 = int32(4)
	v47 = v41 + v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v50 = v42 + v46
	v54 = int32(-2139062144)
	if (v48|(int32(16843008)-v48))&v54 == v54 {
		v41 = v47
		v42 = v50
		v43 = v48
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v58 = v47
	v59 = v50
	v60 = v48
	goto L2
L14:
	;
	goto L13
L15:
	;
	v67 = v58
	v68 = v59
	goto L16
L16:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)) = uint8(v71)
	v73 = int32(1)
	v74 = v67 + v73
	if v71 != 0 {
		v67 = v74
		v68 = v68 + v73
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v77 = v74
	goto L1
L18:
	;
	goto L17
}
func F___strftime_fmt_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int32
	_ = v176
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int64
	_ = v184
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v200 int64
	_ = v200
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int64
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v242 int64
	_ = v242
	var v243 int32
	_ = v243
	var v248 int64
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v253 int64
	_ = v253
	var v254 int64
	_ = v254
	var v273 int64
	_ = v273
	var v275 int64
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int64
	_ = v358
	var v360 int64
	_ = v360
	var v365 int64
	_ = v365
	var v369 int64
	_ = v369
	var v371 int64
	_ = v371
	var v374 int64
	_ = v374
	var v376 int64
	_ = v376
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v433 int64
	_ = v433
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int64
	_ = v565
	var v567 int32
	_ = v567
	var v569 int64
	_ = v569
	var v572 int32
	_ = v572
	var v574 int64
	_ = v574
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int64
	_ = v580
	var v582 int32
	_ = v582
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	switch l2 + int32(-37) {
	case 0:
		goto L14
	default:
		v674 = int32(0)
		goto L1
	case 28:
		goto L46
	case 29:
		goto L44
	case 30:
		goto L43
	case 31:
		v557 = int32(_a2362)
		goto L8
	case 33:
		goto L40
	case 34, 66:
		goto L39
	case 35:
		goto L38
	case 36:
		goto L37
	case 40:
		goto L34
	case 45:
		goto L31
	case 46:
		goto L29
	case 47:
		goto L27
	case 48:
		goto L25
	case 49:
		goto L23
	case 50:
		goto L24
	case 51:
		goto L19
	case 52:
		goto L17
	case 53:
		goto L15
	case 60:
		goto L47
	case 61, 67:
		goto L45
	case 62:
		v497 = int32(131112)
		goto L9
	case 63:
		v48 = int32(48)
		goto L41
	case 64:
		goto L42
	case 69:
		goto L36
	case 72:
		goto L35
	case 73:
		goto L33
	case 75:
		goto L32
	case 77:
		goto L10
	case 78:
		goto L30
	case 79:
		goto L28
	case 80:
		goto L26
	case 82:
		goto L22
	case 83:
		goto L20
	case 84:
		goto L18
	case 85:
		goto L16
	}
L1:
	;
	m.G0 = v16 + int32(80)
	return v674
L2:
	;
	if v613&int32(3) == int32(0) {
		v635 = v613
		goto L191
	} else {
		goto L192
	}
L3:
	;
	v613 = int32(_a216)
	goto L2
L4:
	;
	if l5 != 0 {
		goto L181
	} else {
		goto L182
	}
L5:
	;
	v578 = v572
	v579 = int32(4)
	v580 = v574
	goto L4
L6:
	;
	v578 = v567
	v579 = int32(2)
	v580 = v569
	goto L4
L7:
	;
	v567 = int32(48)
	v569 = v565
	goto L6
L8:
	;
	v560 = F___strftime_l(m, l0, int32(100), v557, l3, l4)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L115
	} else {
		goto L175
	}
L9:
	;
	if v497 != int32(14) {
		goto L150
	} else {
		goto L151
	}
L10:
	;
	v497 = int32(131115)
	goto L9
L11:
	;
	if v436 != int32(14) {
		goto L124
	} else {
		goto L125
	}
L12:
	;
	v436 = v24 | int32(131072)
	goto L11
L13:
	;
	v433 = base.I64_rem_s(v175, int64(100))
	v567 = v176
	v569 = v433
	goto L6
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v674 = int32(_a2363)
	goto L1
L15:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if int32(-1) < v417 {
		goto L120
	} else {
		goto L121
	}
L16:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if int32(-1) < v390 {
		goto L117
	} else {
		goto L118
	}
L17:
	;
	v374 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v376 = v374 + int64(1900)
	if int64(8100) <= v374 {
		goto L113
	} else {
		goto L114
	}
L18:
	;
	v365 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v369 = base.I64_rem_s(v365+int64(1900), int64(100))
	v371 = v369 >> (uint(int64(63)) % 64)
	v565 = v369 ^ v371 - v371
	goto L7
L19:
	;
	v497 = int32(131114)
	goto L9
L20:
	;
	v497 = int32(131113)
	goto L9
L21:
	;
	v578 = int32(48)
	v579 = int32(1)
	v580 = v360
	goto L4
L22:
	;
	v358 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+24)))
	v360 = v358
	goto L21
L23:
	;
	v307 = int32(53)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v312 = int32(7)
	v313 = base.I32_rem_u_s(v309+int32(6), v312)
	v318 = base.I32_div_u_s(v308-v313+v312, v312)
	v319 = v309 - v308
	v323 = base.I32_rem_u_s(v319+int32(369), v312)
	v326 = v318 + base.B2i32(base.Ui32(v323) < base.Ui32(int32(3)))
	if v326 == v307 {
		goto L104
	} else {
		goto L105
	}
L24:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v296 = int32(7)
	v297 = base.I32_rem_u_s(v293+int32(6), v296)
	v302 = base.I32_div_u_s(v292-v297+v296, v296)
	v565 = base.I64_extend_i32_u(v302)
	goto L7
L25:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v287 = int32(7)
	v290 = base.I32_div_u_s(v284-v285+v287, v287)
	v565 = base.I64_extend_i32_u(v290)
	goto L7
L26:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v280 != 0 {
		goto L99
	} else {
		goto L100
	}
L27:
	;
	v557 = int32(_a2364)
	goto L8
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v674 = int32(_a2365)
	goto L1
L29:
	;
	v275 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3))))
	v565 = v275
	goto L7
L30:
	;
	v219 = m.G0
	v221 = v219 - int32(16)
	m.G0 = v221
	v223 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui32(v224) < base.Ui32(int32(12)) {
		v242 = v223
		v243 = v224
		goto L94
	} else {
		goto L95
	}
L31:
	;
	v557 = int32(_a2366)
	goto L8
L32:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if int32(11) < v206 {
		goto L90
	} else {
		goto L91
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v674 = int32(_a397)
	goto L1
L34:
	;
	v200 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+4)))
	v565 = v200
	goto L7
L35:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v565 = base.I64_extend_i32_s(v196 + int32(1))
	goto L7
L36:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v578 = int32(48)
	v579 = int32(3)
	v580 = base.I64_extend_i32_s(v190 + int32(1))
	goto L4
L37:
	;
	v180 = int32(48)
	v181 = int32(2)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v182 != 0 {
		goto L85
	} else {
		goto L86
	}
L38:
	;
	v179 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+8)))
	v565 = v179
	goto L7
L39:
	;
	v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v53 = v51 + int64(1900)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	if int32(2) < v54 {
		goto L53
	} else {
		goto L54
	}
L40:
	;
	v557 = int32(_a2367)
	goto L8
L41:
	;
	v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+12)))
	v567 = v48
	v569 = v49
	goto L6
L42:
	;
	v48 = int32(95)
	goto L41
L43:
	;
	v42 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+20)))
	v46 = base.I64_div_s(v42+int64(1900), int64(100))
	v565 = v46
	goto L7
L44:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui32(int32(11)) < base.Ui32(v37) {
		goto L3
	} else {
		goto L51
	}
L45:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if base.Ui32(int32(11)) < base.Ui32(v32) {
		goto L3
	} else {
		goto L50
	}
L46:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if base.Ui32(int32(6)) < base.Ui32(v27) {
		goto L3
	} else {
		goto L49
	}
L47:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if base.Ui32(v24) <= base.Ui32(int32(6)) {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	goto L3
L49:
	;
	v436 = v27 + int32(131079)
	goto L11
L50:
	;
	v436 = v32 + int32(131086)
	goto L11
L51:
	;
	v436 = v37 + int32(131098)
	goto L11
L52:
	;
	v176 = int32(48)
	if l2 == int32(103) {
		goto L13
	} else {
		goto L84
	}
L53:
	;
	if base.Ui32(v54) < base.Ui32(int32(361)) {
		v175 = v53
		goto L52
	} else {
		goto L69
	}
L54:
	;
	v62 = int32(53)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v67 = int32(7)
	v68 = base.I32_rem_u_s(v64+int32(6), v67)
	v73 = base.I32_div_u_s(v63-v68+v67, v67)
	v74 = v64 - v63
	v78 = base.I32_rem_u_s(v74+int32(369), v67)
	v81 = v73 + base.B2i32(base.Ui32(v78) < base.Ui32(int32(3)))
	if v81 == v62 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	if v111 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L56:
	;
	v111 = v109
	goto L55
L57:
	;
	v103 = base.I32_rem_u_s(v74+int32(371), int32(7))
	switch v103 + int32(-3) {
	case 0:
		goto L64
	case 1:
		v109 = v62
		goto L56
	default:
		goto L63
	}
L58:
	;
	if v81 != 0 {
		v109 = v81
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v84 = int32(52)
	v88 = base.I32_rem_u_s(v74+int32(6), int32(7))
	switch v88 + int32(-4) {
	case 0:
		goto L60
	case 1:
		goto L61
	default:
		v109 = v84
		goto L56
	}
L60:
	;
	v111 = int32(53)
	goto L55
L61:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v93 = base.I32_rem_s(v91, int32(400))
	v96 = F_is_leap(m, v93+int32(-1))
	mBase = m.M
	if v96 == int32(0) {
		v109 = v84
		goto L56
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v109 = int32(1)
	goto L56
L64:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v107 = F_is_leap(m, v106)
	mBase = m.M
	if v107 != 0 {
		v109 = v62
		goto L56
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v114 = v53
	goto L68
L67:
	;
	v114 = v51 + int64(1899)
	goto L68
L68:
	;
	v175 = v114
	goto L52
L69:
	;
	v122 = int32(53)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v127 = int32(7)
	v128 = base.I32_rem_u_s(v124+int32(6), v127)
	v133 = base.I32_div_u_s(v123-v128+v127, v127)
	v134 = v124 - v123
	v138 = base.I32_rem_u_s(v134+int32(369), v127)
	v141 = v133 + base.B2i32(base.Ui32(v138) < base.Ui32(int32(3)))
	if v141 == v122 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	if v171 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L71:
	;
	v171 = v169
	goto L70
L72:
	;
	v163 = base.I32_rem_u_s(v134+int32(371), int32(7))
	switch v163 + int32(-3) {
	case 0:
		goto L79
	case 1:
		v169 = v122
		goto L71
	default:
		goto L78
	}
L73:
	;
	if v141 != 0 {
		v169 = v141
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v144 = int32(52)
	v148 = base.I32_rem_u_s(v134+int32(6), int32(7))
	switch v148 + int32(-4) {
	case 0:
		goto L75
	case 1:
		goto L76
	default:
		v169 = v144
		goto L71
	}
L75:
	;
	v171 = int32(53)
	goto L70
L76:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v153 = base.I32_rem_s(v151, int32(400))
	v156 = F_is_leap(m, v153+int32(-1))
	mBase = m.M
	if v156 == int32(0) {
		v169 = v144
		goto L71
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v169 = int32(1)
	goto L71
L79:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v167 = F_is_leap(m, v166)
	mBase = m.M
	if v167 != 0 {
		v169 = v122
		goto L71
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v174 = v51 + int64(1901)
	goto L83
L82:
	;
	v174 = v53
	goto L83
L83:
	;
	v175 = v174
	goto L52
L84:
	;
	v572 = v176
	v574 = v175
	goto L5
L85:
	;
	v184 = base.I64_extend_i32_s(v182)
	if int32(12) < v182 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v578 = v180
	v579 = v181
	v580 = int64(12)
	goto L4
L87:
	;
	v189 = v184 + int64(-12)
	goto L89
L88:
	;
	v189 = v184
	goto L89
L89:
	;
	v578 = v180
	v579 = v181
	v580 = v189
	goto L4
L90:
	;
	v209 = int32(131111)
	goto L92
L91:
	;
	v209 = int32(131110)
	goto L92
L92:
	;
	v436 = v209
	goto L11
L93:
	;
	v273 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+36)))
	v360 = v254 + (v248 + base.I64_extend_i32_s(v250) + base.I64_extend_i32_s(v251+int32(-1))*int64(86400) + v252*int64(3600) + v253*int64(60)) - v273
	goto L21
L94:
	;
	v248 = F___year_to_secs(m, v242, v221+int32(12))
	mBase = m.M
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v250 = F___month_to_secs(m, v243, v249)
	mBase = m.M
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v252 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+8)))
	v253 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3)+4)))
	v254 = int64(*(*int32)(unsafe.Add(mBase, uint32(l3))))
	m.G0 = v221 + int32(16)
	goto L93
L95:
	;
	v227 = int32(12)
	v228 = base.I32_div_s(v224, v227)
	v231 = v224 - v228*v227
	if v231 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v236 = v231 + v227
	goto L98
L97:
	;
	v236 = v231
	goto L98
L98:
	;
	v242 = base.I64_extend_i32_s(v228+v231>>(uint(int32(31))%32)) + v223
	v243 = v236
	goto L94
L99:
	;
	v282 = v280
	goto L101
L100:
	;
	v282 = int32(7)
	goto L101
L101:
	;
	v360 = base.I64_extend_i32_s(v282)
	goto L21
L102:
	;
	v565 = base.I64_extend_i32_u(v356)
	goto L7
L103:
	;
	v356 = v354
	goto L102
L104:
	;
	v348 = base.I32_rem_u_s(v319+int32(371), int32(7))
	switch v348 + int32(-3) {
	case 0:
		goto L111
	case 1:
		v354 = v307
		goto L103
	default:
		goto L110
	}
L105:
	;
	if v326 != 0 {
		v354 = v326
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v329 = int32(52)
	v333 = base.I32_rem_u_s(v319+int32(6), int32(7))
	switch v333 + int32(-4) {
	case 0:
		goto L107
	case 1:
		goto L108
	default:
		v354 = v329
		goto L103
	}
L107:
	;
	v356 = int32(53)
	goto L102
L108:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v338 = base.I32_rem_s(v336, int32(400))
	v341 = F_is_leap(m, v338+int32(-1))
	mBase = m.M
	if v341 == int32(0) {
		v354 = v329
		goto L103
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v354 = int32(1)
	goto L103
L111:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v352 = F_is_leap(m, v351)
	mBase = m.M
	if v352 != 0 {
		v354 = v307
		goto L103
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v376
	v385 = F_snprintf(m, l0, int32(100), int32(_a2368), v16+int32(48))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v572 = int32(48)
	v574 = v376
	goto L5
L115:
	;
	return int32(0)
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v385
	v674 = l0
	goto L1
L117:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	v397 = int32(3600)
	v398 = base.I32_div_s(v396, v397)
	v399 = int32(100)
	v406 = base.I32_div_s(base.I32_extend16_s(v396-v398*v397), int32(60))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v398*v399 + base.I32_extend16_s(v406)
	v414 = F_snprintf(m, l0, v399, int32(_a2369), v16+int32(64))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L115
	} else {
		goto L119
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v674 = int32(_a188)
	goto L1
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v414
	v674 = l0
	goto L1
L120:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v424 = int32(9116396)
	F___lock(m, v424)
	mBase = m.M
	F_do_tzset(m)
	mBase = m.M
	F___unlock(m, v424)
	mBase = m.M
	goto L122
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v674 = int32(_a188)
	goto L1
L122:
	;
	v613 = v423
	goto L2
L123:
	;
	v613 = v495
	goto L2
L124:
	;
	v447 = v436 >> (uint(int32(16)) % 32)
	v448 = int32(65535)
	v449 = v436 & v448
	if v449 != v448 {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v444 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v445 = int32(_a2370)
	goto L128
L127:
	;
	v445 = int32(_a2371)
	goto L128
L128:
	;
	v495 = v445
	goto L123
L129:
	;
	v462 = int32(_a188)
	switch v447 + int32(-1) {
	case 0:
		goto L139
	case 1:
		goto L138
	default:
		v489 = v462
		goto L135
	case 4:
		goto L137
	}
L130:
	;
	if int32(5) < v447 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l4+v447<<(uint(int32(2))%32))))
	if v457 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v461 = v457 + int32(8)
	goto L134
L133:
	;
	v461 = int32(_a2246)
	goto L134
L134:
	;
	v495 = v461
	goto L123
L135:
	;
	v495 = v489
	goto L123
L136:
	;
	if v449 != 0 {
		goto L143
	} else {
		goto L144
	}
L137:
	;
	if base.Ui32(int32(3)) < base.Ui32(v449) {
		v489 = v462
		goto L135
	} else {
		goto L142
	}
L138:
	;
	if base.Ui32(int32(49)) < base.Ui32(v449) {
		v489 = v462
		goto L135
	} else {
		goto L141
	}
L139:
	;
	if base.Ui32(int32(1)) < base.Ui32(v449) {
		v489 = v462
		goto L135
	} else {
		goto L140
	}
L140:
	;
	v474 = int32(_a2372)
	goto L136
L141:
	;
	v474 = int32(_a2373)
	goto L136
L142:
	;
	v474 = int32(_a2374)
	goto L136
L143:
	;
	v475 = v474
	v478 = v449
	goto L145
L144:
	;
	v495 = v474
	goto L123
L145:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
	v482 = v475 + int32(1)
	if v480 != 0 {
		v475 = v482
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v489 = v482
	goto L135
L147:
	;
	v484 = v478 + int32(-1)
	if v484 != 0 {
		v475 = v482
		v478 = v484
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v557 = v556
	goto L8
L150:
	;
	v508 = v497 >> (uint(int32(16)) % 32)
	v509 = int32(65535)
	v510 = v497 & v509
	if v510 != v509 {
		goto L155
	} else {
		goto L156
	}
L151:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v505 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v506 = int32(_a2370)
	goto L154
L153:
	;
	v506 = int32(_a2371)
	goto L154
L154:
	;
	v556 = v506
	goto L149
L155:
	;
	v523 = int32(_a188)
	switch v508 + int32(-1) {
	case 0:
		goto L165
	case 1:
		goto L164
	default:
		v550 = v523
		goto L161
	case 4:
		goto L163
	}
L156:
	;
	if int32(5) < v508 {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l4+v508<<(uint(int32(2))%32))))
	if v518 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v522 = v518 + int32(8)
	goto L160
L159:
	;
	v522 = int32(_a2246)
	goto L160
L160:
	;
	v556 = v522
	goto L149
L161:
	;
	v556 = v550
	goto L149
L162:
	;
	if v510 != 0 {
		goto L169
	} else {
		goto L170
	}
L163:
	;
	if base.Ui32(int32(3)) < base.Ui32(v510) {
		v550 = v523
		goto L161
	} else {
		goto L168
	}
L164:
	;
	if base.Ui32(int32(49)) < base.Ui32(v510) {
		v550 = v523
		goto L161
	} else {
		goto L167
	}
L165:
	;
	if base.Ui32(int32(1)) < base.Ui32(v510) {
		v550 = v523
		goto L161
	} else {
		goto L166
	}
L166:
	;
	v535 = int32(_a2372)
	goto L162
L167:
	;
	v535 = int32(_a2373)
	goto L162
L168:
	;
	v535 = int32(_a2374)
	goto L162
L169:
	;
	v536 = v535
	v539 = v510
	goto L171
L170:
	;
	v556 = v535
	goto L149
L171:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536))))
	v543 = v536 + int32(1)
	if v541 != 0 {
		v536 = v543
		goto L171
	} else {
		goto L173
	}
L172:
	;
	v550 = v543
	goto L161
L173:
	;
	v545 = v539 + int32(-1)
	if v545 != 0 {
		v536 = v543
		v539 = v545
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v560
	if v560 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v564 = l0
	goto L178
L177:
	;
	v564 = int32(0)
	goto L178
L178:
	;
	v674 = v564
	goto L1
L179:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v580
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v579
	v608 = F_snprintf(m, l0, int32(100), int32(_a2375), v16)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L115
	} else {
		goto L188
	}
L180:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v580
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v579
	v601 = F_snprintf(m, l0, int32(100), int32(_a2376), v16+int32(32))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L115
	} else {
		goto L187
	}
L181:
	;
	v582 = l5
	goto L183
L182:
	;
	v582 = v578
	goto L183
L183:
	;
	if v582 == int32(95) {
		goto L180
	} else {
		goto L184
	}
L184:
	;
	if v582 != int32(45) {
		goto L179
	} else {
		goto L185
	}
L185:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v580
	v592 = F_snprintf(m, l0, int32(100), int32(_a1006), v16+int32(16))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L115
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v592
	v674 = l0
	goto L1
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v601
	v674 = l0
	goto L1
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v608
	v674 = l0
	goto L1
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v668
	v674 = v613
	goto L1
L190:
	;
	v668 = v660 - v613
	goto L189
L191:
	;
	v639 = v635
	goto L199
L192:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613))))
	if v621 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v624 = v613
	goto L195
L194:
	;
	v668 = v613 - v613
	goto L189
L195:
	;
	v628 = v624 + int32(1)
	if v628&int32(3) == int32(0) {
		v635 = v628
		goto L191
	} else {
		goto L197
	}
L197:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628))))
	if v633 != 0 {
		v624 = v628
		goto L195
	} else {
		goto L198
	}
L198:
	;
	v660 = v628
	goto L190
L199:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v639)))
	v648 = int32(-2139062144)
	if (int32(16843008)-v645|v645)&v648 == v648 {
		v639 = v639 + int32(4)
		goto L199
	} else {
		goto L201
	}
L200:
	;
	v654 = v639
	goto L202
L201:
	;
	goto L200
L202:
	;
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654))))
	if v658 != 0 {
		v654 = v654 + int32(1)
		goto L202
	} else {
		goto L204
	}
L203:
	;
	v660 = v654
	goto L190
L204:
	;
	goto L203
}
func F___strftime_l(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	v15 = m.G0
	v17 = v15 - int32(128)
	m.G0 = v17
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(128)
	return v332
L2:
	;
	v23 = l2
	v28 = int32(0)
	goto L6
L3:
	;
	v332 = int32(0)
	goto L1
L4:
	;
	v324 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v316))) = uint8(v324)
	v332 = v315
	goto L1
L5:
	;
	if v297 == l1 {
		goto L73
	} else {
		goto L74
	}
L6:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v35 == int32(37) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v297 = v280
	goto L5
L8:
	;
	if base.Ui32(v280) < base.Ui32(l1) {
		v23 = v275 + int32(1)
		v28 = v280
		goto L6
	} else {
		goto L72
	}
L9:
	;
	v59 = v56 & int32(255)
	v62 = v23 + v55 + base.B2i32(v59 == int32(43))
	v63 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
	if base.Ui32(int32(9)) < base.Ui32(v63+int32(-48)) {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
	v54 = v40
	v55 = int32(2)
	v56 = v52
	goto L9
L11:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v28))) = uint8(v35)
	v275 = v23
	v280 = v28 + int32(1)
	goto L8
L12:
	;
	v38 = int32(0)
	v39 = int32(1)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	switch v40 + int32(-45) {
	case 0, 3:
		goto L10
	case 1, 2:
		v54 = v38
		v55 = v39
		v56 = v40
		goto L9
	default:
		goto L15
	}
L13:
	;
	if v35 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v315 = v28
	v316 = v28
	goto L4
L15:
	;
	if v40 == int32(95) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if v40 != 0 {
		v54 = v38
		v55 = v39
		v56 = v40
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L11
L18:
	;
	v79 = int32(0)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v82 = v80 + int32(-67)
	if base.Ui32(int32(22)) < base.Ui32(v82) {
		v92 = v79
		goto L22
	} else {
		goto L23
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v62
	v77 = int32(0)
	v78 = v62
	goto L18
L20:
	;
	v72 = F_strtox_2(m, v62, v17+int32(12), int32(10), int64(4294967295))
	mBase = m.M
	goto L21
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v77 = base.I32_wrap_i64(v72)
	v78 = v74
	goto L18
L22:
	;
	if v80 == int32(79) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	if int32(1)<<(uint(v82)%32)&int32(4194329) == int32(0) {
		v92 = v79
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v77 != 0 {
		v92 = v77
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v92 = base.B2i32(v78 != v62)
	goto L22
L26:
	;
	v107 = F___strftime_fmt_1(m, v17+int32(16), v17+int32(124), base.I32_extend8_s(v101), l3, l4, v54)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	v100 = v78 + int32(1)
	v101 = v99
	goto L26
L28:
	;
	if v80 == int32(69) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v100 = v78
	v101 = v80
	goto L26
L30:
	;
	return int32(0)
L31:
	;
	if v107 == int32(0) {
		v297 = v28
		goto L5
	} else {
		goto L32
	}
L32:
	;
	if v92 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v262 = l1 - v255
	if base.Ui32(v257) < base.Ui32(v262) {
		goto L66
	} else {
		goto L67
	}
L34:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	switch v114 + int32(-43) {
	case 0, 2:
		goto L37
	default:
		goto L38
	}
L35:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v17)+124))
	v255 = v28
	v256 = v107
	v257 = v113
	goto L33
L36:
	;
	if v124&int32(255) != int32(48) {
		v164 = v125
		v165 = v126
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v17)+124))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	v124 = v121
	v125 = v107 + int32(1)
	v126 = v118 + int32(-1)
	goto L36
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v17)+124))
	v124 = v114
	v125 = v107
	v126 = v117
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+124)) = v165
	v178 = int32(0)
	goto L45
L40:
	;
	v139 = v125
	v140 = v126
	goto L41
L41:
	;
	v145 = int32(*(*int8)(unsafe.Add(mBase, uint32(v139)+1)))
	if base.Ui32(int32(9)) < base.Ui32(v145+int32(-48)) {
		v164 = v139
		v165 = v140
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v164 = v151
	v165 = v153
	goto L39
L43:
	;
	v151 = v139 + int32(1)
	v153 = v140 + int32(-1)
	if v145 == int32(48) {
		v139 = v151
		v140 = v153
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v189 = int32(*(*int8)(unsafe.Add(mBase, uint32(v164+v178))))
	if base.Ui32(v189+int32(-48)) < base.Ui32(int32(10)) {
		v178 = v178 + int32(1)
		goto L45
	} else {
		goto L47
	}
L46:
	;
	if base.Ui32(v165) < base.Ui32(v92) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	v195 = v92
	goto L50
L49:
	;
	v195 = v165
	goto L50
L50:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if int32(-1900) <= v196 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	if base.Ui32(v220) <= base.Ui32(v165) {
		v255 = v221
		v256 = v164
		v257 = v165
		goto L33
	} else {
		goto L60
	}
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v28))) = uint8(v213)
	v220 = v195 + int32(-1)
	v221 = v28 + int32(1)
	goto L51
L53:
	;
	if v59 != int32(43) {
		v220 = v195
		v221 = v28
		goto L51
	} else {
		goto L55
	}
L54:
	;
	v213 = int32(45)
	goto L52
L55:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v207 == int32(67) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v210 = int32(3)
	goto L58
L57:
	;
	v210 = int32(5)
	goto L58
L58:
	;
	if base.Ui32(v195-v165+v178) < base.Ui32(v210) {
		v220 = v195
		v221 = v28
		goto L51
	} else {
		goto L59
	}
L59:
	;
	v213 = int32(43)
	goto L52
L60:
	;
	if base.Ui32(l1) <= base.Ui32(v221) {
		v255 = v221
		v256 = v164
		v257 = v165
		goto L33
	} else {
		goto L61
	}
L61:
	;
	v231 = v220
	v232 = v221
	goto L62
L62:
	;
	v240 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v232))) = uint8(v240)
	v243 = v232 + int32(1)
	v245 = v231 + int32(-1)
	if base.Ui32(v245) <= base.Ui32(v165) {
		v255 = v243
		v256 = v164
		v257 = v165
		goto L33
	} else {
		goto L64
	}
L63:
	;
	v255 = v243
	v256 = v164
	v257 = v165
	goto L33
L64:
	;
	if base.Ui32(v243) < base.Ui32(l1) {
		v231 = v245
		v232 = v243
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v264 = v257
	goto L68
L67:
	;
	v264 = v262
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+124)) = v264
	if v264 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v17)+124))
	v275 = v100
	v280 = v271 + v255
	goto L8
L70:
	;
	goto L69
L71:
	;
	v269 = F__emscripten_memcpy_bulkmem(m, l0+v255, v256, v264)
	mBase = m.M
	goto L70
L72:
	;
	goto L7
L73:
	;
	v307 = l1 + int32(-1)
	goto L75
L74:
	;
	v307 = v297
	goto L75
L75:
	;
	v315 = int32(0)
	v316 = v307
	goto L4
}
func F___synccall(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	m.T0[l0].(func(*base.Module, int32))(m, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_saddCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v183 int64
	_ = v183
	var v188 int32
	_ = v188
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = F_lookupKeyWrite(m, v14, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v17
	v21 = F_checkType(m, l0, v17, int32(2))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v12 + int32(16)
	return
L4:
	;
	if v21 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(3) <= v104 {
		goto L38
	} else {
		goto L39
	}
L7:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v83 = v81 + int32(-2)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v86 = v84 & int32(240)
	if v86 != int32(176) {
		goto L31
	} else {
		goto L32
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v25 = F_objectGetVal(m, v24)
	mBase = m.M
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v28 = v26 + int32(-2)
	v29 = int32(0)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(-1)))))
	switch v35 & int32(7) {
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
		v52 = v29
		goto L12
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	F_dbAdd(m, v74, v76, v12+int32(12))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L29
	}
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[1079]))
	if base.Ui32(v63) < base.Ui32(v28) {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	if v56 != 0 {
		goto L10
	} else {
		goto L21
	}
L12:
	;
	v55 = F_string2ll(m, v25, v52, v29)
	mBase = m.M
	if v55 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(-17))))
	v52 = v51
	goto L12
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(-9))))
	v52 = v48
	goto L12
L15:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+int32(-5)))))
	v52 = v45
	goto L12
L16:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(-3)))))
	v52 = v42
	goto L12
L17:
	;
	v52 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
	goto L12
L18:
	;
	v56 = int32(0)
	goto L20
L19:
	;
	v56 = int32(-1)
	goto L20
L20:
	;
	goto L11
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[1078]))
	if base.Ui32(v58) < base.Ui32(v28) {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v60 = F_createIntsetObject(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v72 = v60
	goto L9
L24:
	;
	v67 = F_createSetObject(m)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	v65 = F_createSetListpackObject(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v72 = v65
	goto L9
L27:
	;
	v69 = F_objectGetVal(m, v67)
	mBase = m.M
	v70 = F_hashtableExpand(m, v69, v28)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v72 = v67
	goto L9
L29:
	;
	goto L6
L30:
	;
	v99 = F_setTypeConvertAndExpand(m, v17, int32(2), v83, int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L36
	}
L31:
	;
	if v86 != int32(96) {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[1079]))
	if base.Ui32(v90) < base.Ui32(v83) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[1078]))
	if base.Ui32(v83) <= base.Ui32(v95) {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	goto L6
L37:
	;
	F_addReplyLongLong(m, l0, v183)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L54
	}
L38:
	;
	v112 = int32(2)
	v116 = int32(0)
	goto L40
L39:
	;
	v183 = int64(0)
	goto L37
L40:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v112<<(uint(int32(2))%32))))
	v126 = F_objectGetVal(m, v125)
	mBase = m.M
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+int32(-1)))))
	switch v129 & int32(7) {
	case 0:
		goto L47
	case 1:
		goto L46
	case 2:
		goto L45
	case 3:
		goto L44
	case 4:
		goto L43
	default:
		v146 = int32(0)
		goto L42
	}
L41:
	;
	if v153 != 0 {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	v149 = F_setTypeAddAux(m, v119, v126, v146, int64(0), int32(1))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L48
	}
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v126+int32(-17))))
	v146 = v145
	goto L42
L44:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v126+int32(-9))))
	v146 = v142
	goto L42
L45:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126+int32(-5)))))
	v146 = v139
	goto L42
L46:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+int32(-3)))))
	v146 = v136
	goto L42
L47:
	;
	v146 = int32(base.Ui32(v129) >> (uint(int32(3)) % 32))
	goto L42
L48:
	;
	v153 = v116 + base.B2i32(v149 != int32(0))
	v155 = v112 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v155 < v156 {
		v112 = v155
		v116 = v153
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L41
L50:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	F_signalModifiedKey(m, l0, v159, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L52
	}
L51:
	;
	v183 = int64(0)
	goto L37
L52:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+28))
	F_notifyKeyspaceEvent(m, int32(32), int32(_a1516), v167, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v172 = int32(_a69)
	v174 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	v175 = base.I64_extend_i32_u(v153)
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v174 + v175
	v183 = v175
	goto L37
L54:
	;
	goto L3
}
func F_sampleEntriesScanFn(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v5) <= base.Ui32(v4) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4 + int32(1)
		v18 = int32(0)
		F___lock(m, int32(9116960))
		mBase = m.M
		v25 = *(*int32)(unsafe.Add(mBase, _consts[245]))
		v27 = *(*int32)(unsafe.Add(mBase, _consts[246]))
		if v27 != 0 {
			v31 = int32(0)
			v32 = *(*int32)(unsafe.Add(mBase, _consts[247]))
			v33 = int32(2)
			v35 = v25 + v32<<(uint(v33)%32)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			v38 = *(*int32)(unsafe.Add(mBase, _consts[248]))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v25+v38<<(uint(v33)%32))))
			v43 = v36 + v42
			*(*int32)(unsafe.Add(mBase, uint32(v35))) = v43
			v48 = v38 + int32(1)
			if v48 == v27 {
				v50 = v31
			} else {
				v50 = v48
			}
			*(*int32)(unsafe.Add(mBase, _consts[248])) = v50
			v52 = int32(0)
			v55 = v32 + int32(1)
			if v55 == v27 {
				v57 = v52
			} else {
				v57 = v55
			}
			*(*int32)(unsafe.Add(mBase, _consts[247])) = v57
			v62 = int32(base.Ui32(v43) >> (uint(int32(1)) % 32))
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
			v29 = F_lcg31(m, v28)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v25))) = v29
			v62 = v29
		}
		F___unlock(m, int32(9116960))
		mBase = m.M
		v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v68 = base.I32_rem_u_s(v62, v67)
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v69) <= base.Ui32(v68) {
		} else {
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v71+v68<<(uint(int32(2))%32)))) = l1
		}
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4 + int32(1)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10+v4<<(uint(int32(2))%32)))) = l1
		return
	}
}
func F_save(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	v2 = l1
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v15 = v13 + int32(1)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(v16) < base.Ui32(v15) {
		if base.Ui32(v16) < base.Ui32(int32(2147483646)) {
			v124 = v16
			v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v127 = v124 << (uint(int32(1)) % 32)
			if v127 == int32(-2) {
				v133 = F_luaM_toobig(m, v125)
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return
				} else {
					v135 = v133
					*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v127
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v135
					v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v141 = v135
					v142 = v138
					v143 = v138 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v143
					*(*uint8)(unsafe.Add(mBase, uint32(v141+v142))) = uint8(v2)
					m.G0 = v10 + int32(96)
					return
				}
			} else {
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v131 = F_luaM_realloc_(m, v125, v130, v124, v127)
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return
				} else {
					v135 = v131
					*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v127
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v135
					v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v141 = v135
					v142 = v138
					v143 = v138 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v143
					*(*uint8)(unsafe.Add(mBase, uint32(v141+v142))) = uint8(v2)
					m.G0 = v10 + int32(96)
					return
				}
			}
		} else {
			v21 = int32(16)
			v22 = v10 + v21
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			v25 = v23 + v21
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
			switch v29 + int32(-61) {
			case 0:
				v34 = F_strncpy(m, v22, v23+int32(17), int32(80))
				mBase = m.M
				v38 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v34+int32(79)))) = uint8(v38)
			default:
				v56 = m.G3
				v59 = F_strcspn(m, v25, v56+int32(_a2249))
				mBase = m.M
				v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+uint32(_consts[1201]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(24)))) = uint16(v66)
				v68 = *(*int64)(unsafe.Add(mBase, uint32(v56)+uint32(_consts[1202])))
				*(*int64)(unsafe.Add(mBase, uint32(v22))) = v68
				v71 = int32(63)
				if base.Ui32(v59) < base.Ui32(v71) {
					v73 = v59
				} else {
					v73 = v71
				}
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v73))))
				if v75 == int32(0) {
					v83 = F_strcat(m, v22, v25)
					mBase = m.M
				} else {
					v78 = F_strncat(m, v22, v25, v73)
					mBase = m.M
					v79 = F_strlen(m, v78)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v78+v79))) = int32(3026478)
				}
				v85 = F_strlen(m, v22)
				mBase = m.M
				v86 = v22 + v85
				v87 = m.G3
				v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[1203]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v86))) = uint16(v90)
				v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[1204]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v86+int32(2)))) = uint8(v96)
			case 3:
				v41 = v23 + int32(17)
				v42 = F_strlen(m, v41)
				mBase = m.M
				v43 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v43)
				v46 = int32(72)
				if base.Ui32(v42) <= base.Ui32(v46) {
					v54 = v41
				} else {
					v48 = F_strlen(m, v22)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v22+v48))) = int32(3026478)
					v54 = v41 + (v42 - v46)
				}
				v55 = F_strcat(m, v22, v54)
				mBase = m.M
			}
			v103 = m.G3
			v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v103 + int32(_a2259)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v105
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(16)
			v115 = F_luaO_pushfstring(m, v104, v103+int32(_a2260), v10)
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return
			} else {
				v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				F_luaD_throw(m, v117, int32(3))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return
				} else {
					v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					v124 = v121
					v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v127 = v124 << (uint(int32(1)) % 32)
					if v127 == int32(-2) {
						v133 = F_luaM_toobig(m, v125)
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return
						} else {
							v135 = v133
							*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v127
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v135
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v141 = v135
							v142 = v138
							v143 = v138 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v143
							*(*uint8)(unsafe.Add(mBase, uint32(v141+v142))) = uint8(v2)
							m.G0 = v10 + int32(96)
							return
						}
					} else {
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v131 = F_luaM_realloc_(m, v125, v130, v124, v127)
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
							return
						} else {
							v135 = v131
							*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v127
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v135
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v141 = v135
							v142 = v138
							v143 = v138 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v143
							*(*uint8)(unsafe.Add(mBase, uint32(v141+v142))) = uint8(v2)
							m.G0 = v10 + int32(96)
							return
						}
					}
				}
			}
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v141 = v18
		v142 = v13
		v143 = v15
		*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v143
		*(*uint8)(unsafe.Add(mBase, uint32(v141+v142))) = uint8(v2)
		m.G0 = v10 + int32(96)
		return
	}
}
func F_scanDatabaseForDeletedKeys(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
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
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v321 int32
	_ = v321
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = F_dictGetSafeIterator(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v17)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L2
	} else {
		goto L80
	}
L2:
	;
	return
L3:
	;
	v26 = v17 + int32(20)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v122 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v33 = v26
	v34 = v30
	goto L8
L6:
	;
	v30 = int32(1)
	goto L5
L7:
	;
	v30 = int32(0)
	goto L5
L8:
	;
	switch v34 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v34 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v114
	if v114 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v38 != int32(-1) {
		v77 = v38
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v78 = int32(1)
	v79 = v77 + v78
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v79
	v81 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v85+int32(26)))))
	if v89 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v42 != 0 {
		v77 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v44 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	if v71 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v51 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+16)))
	v52 = int64(*(*int8)(unsafe.Add(mBase, uint32(v43)+27)))
	v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(v43)+8)))
	v54 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+12)))
	v55 = int64(*(*int8)(unsafe.Add(mBase, uint32(v43)+26)))
	v56 = int64(*(*int32)(unsafe.Add(mBase, uint32(v43)+4)))
	v57 = F_wangHash64(m, v56)
	mBase = m.M
	v59 = F_wangHash64(m, v55+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v54+v59)
	mBase = m.M
	v63 = F_wangHash64(m, v53+v61)
	mBase = m.M
	v65 = F_wangHash64(m, v52+v63)
	mBase = m.M
	v67 = F_wangHash64(m, v51+v65)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v70 = v69
	goto L17
L19:
	;
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+24)))
	v49 = v47 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+24)) = uint16(v49)
	v70 = v43
	goto L17
L20:
	;
	v77 = v71 + int32(-1)
	goto L14
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v77 = v74
	goto L14
L22:
	;
	v104 = int32(2)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v84+v102<<(uint(v104)%32)+int32(4))))
	v33 = v109 + v103<<(uint(v104)%32)
	v34 = int32(1)
	goto L8
L23:
	;
	v93 = v81
	goto L25
L24:
	;
	v93 = v78 << (uint(v89) % 32)
	goto L25
L25:
	;
	if v79 < v93 {
		v102 = v85
		v103 = v79
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v85 != 0 {
		v122 = v81
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	if v95 == int32(-1) {
		v122 = v81
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+4)) = int64(4294967296)
	v102 = int32(1)
	v103 = int32(0)
	goto L22
L29:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v118
	v122 = v114
	goto L11
L30:
	;
	v132 = v122
	goto L31
L31:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	goto L33
L32:
	;
	goto L1
L33:
	;
	v140 = F_objectGetVal(m, v139)
	mBase = m.M
	v141 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v143 == v141 {
		v148 = v141
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(0)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v154 = F_kvstoreHashtableFind(m, v151, v148, v140, v14+int32(8))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L37
	}
L35:
	;
	v146 = F_getKeySlot(m, v140)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v148 = v146
	goto L34
L37:
	;
	v156 = int32(-1)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v158 == int32(0) {
		v164 = v156
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if l1 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v164 = v161 & int32(15)
	goto L38
L40:
	;
	if base.B2i32(v158 != int32(0))&(v191^int32(1)) != 0 {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	v166 = F_objectGetVal(m, v139)
	mBase = m.M
	v167 = int32(0)
	v170 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v170 == v167 {
		v175 = v167
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v190 = v156
	v191 = int32(0)
	goto L40
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = int32(0)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v181 = F_kvstoreHashtableFind(m, v178, v175, v166, v14+int32(12))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L46
	}
L44:
	;
	v173 = F_getKeySlot(m, v166)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v175 = v173
	goto L43
L46:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v183 == int32(0) {
		v190 = v156
		v191 = v167
		goto L40
	} else {
		goto L47
	}
L47:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v190 = v186 & int32(15)
	v191 = int32(1)
	goto L40
L48:
	;
	v209 = v17 + int32(20)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v210 != 0 {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	F_signalDeletedKeyAsReady(m, l0, v139, v164)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L52
	}
L50:
	;
	if v164 == v190 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	goto L48
L53:
	;
	if v305 != 0 {
		v132 = v305
		goto L31
	} else {
		goto L79
	}
L54:
	;
	v216 = v209
	v217 = v213
	goto L57
L55:
	;
	v213 = int32(1)
	goto L54
L56:
	;
	v213 = int32(0)
	goto L54
L57:
	;
	switch v217 {
	case 0:
		goto L62
	default:
		goto L61
	}
L59:
	;
	v217 = int32(0)
	goto L57
L60:
	;
	goto L53
L61:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v297
	if v297 == int32(0) {
		goto L59
	} else {
		goto L78
	}
L62:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v221 != int32(-1) {
		v260 = v221
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v261 = int32(1)
	v262 = v260 + v261
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v262
	v264 = int32(0)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267+v268+int32(26)))))
	if v272 == int32(255) {
		goto L72
	} else {
		goto L73
	}
L64:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v225 != 0 {
		v260 = int32(-1)
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v227 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+20))
	if v254 != int32(-1) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v234 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v226)+16)))
	v235 = int64(*(*int8)(unsafe.Add(mBase, uint32(v226)+27)))
	v236 = int64(*(*int32)(unsafe.Add(mBase, uint32(v226)+8)))
	v237 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v226)+12)))
	v238 = int64(*(*int8)(unsafe.Add(mBase, uint32(v226)+26)))
	v239 = int64(*(*int32)(unsafe.Add(mBase, uint32(v226)+4)))
	v240 = F_wangHash64(m, v239)
	mBase = m.M
	v242 = F_wangHash64(m, v238+v240)
	mBase = m.M
	v244 = F_wangHash64(m, v237+v242)
	mBase = m.M
	v246 = F_wangHash64(m, v236+v244)
	mBase = m.M
	v248 = F_wangHash64(m, v235+v246)
	mBase = m.M
	v250 = F_wangHash64(m, v234+v248)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v250
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v253 = v252
	goto L66
L68:
	;
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226)+24)))
	v232 = v230 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v226)+24)) = uint16(v232)
	v253 = v226
	goto L66
L69:
	;
	v260 = v254 + int32(-1)
	goto L63
L70:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v260 = v257
	goto L63
L71:
	;
	v287 = int32(2)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v267+v285<<(uint(v287)%32)+int32(4))))
	v216 = v292 + v286<<(uint(v287)%32)
	v217 = int32(1)
	goto L57
L72:
	;
	v276 = v264
	goto L74
L73:
	;
	v276 = v261 << (uint(v272) % 32)
	goto L74
L74:
	;
	if v262 < v276 {
		v285 = v268
		v286 = v262
		goto L71
	} else {
		goto L75
	}
L75:
	;
	if v268 != 0 {
		v305 = v264
		goto L60
	} else {
		goto L76
	}
L76:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v267)+20))
	if v278 == int32(-1) {
		v305 = v264
		goto L60
	} else {
		goto L77
	}
L77:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+4)) = int64(4294967296)
	v285 = int32(1)
	v286 = int32(0)
	goto L71
L78:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v301
	v305 = v297
	goto L60
L79:
	;
	goto L32
L80:
	;
	m.G0 = v14 + int32(16)
	return
}
func F_sdscatrepr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
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
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v566 int32
	_ = v566
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v17 = F__sdsMakeRoomFor(m, l0, l2+int32(2), int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	v41 = int32(1)
	v43 = F__sdsMakeRoomFor(m, v17, v41, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L10
	}
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-17))))
	v40 = v39
	goto L1
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-9))))
	v40 = v36
	goto L1
L4:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(-5)))))
	v40 = v33
	goto L1
L5:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-3)))))
	v40 = v30
	goto L1
L6:
	;
	v40 = int32(base.Ui32(v23) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	return int32(0)
L8:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-1)))))
	switch v23 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v40 = int32(0)
		goto L1
	}
L9:
	;
	if l2 == int32(0) {
		v504 = v43
		goto L18
	} else {
		goto L19
	}
L10:
	;
	if v43 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v48 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v43+v40))) = uint8(v48)
	v51 = v40 + int32(1)
	v53 = v43 + int32(-1)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	switch v54 & int32(7) {
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
		goto L12
	}
L12:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43+v51))) = uint8(v74)
	goto L9
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v43+int32(-17)))) = base.I64_extend_i32_u(v51)
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+int32(-9)))) = v51
	goto L12
L15:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v43+int32(-5)))) = uint16(v51)
	goto L12
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v43+int32(-3)))) = uint8(v51)
	goto L12
L17:
	;
	v58 = v51 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v58)
	goto L12
L18:
	;
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504+int32(-1)))))
	switch v515 & int32(7) {
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
		v532 = int32(0)
		goto L138
	}
L19:
	;
	v80 = v43
	v81 = l1
	v82 = l2
	goto L20
L20:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v89 = base.I32_extend8_s(v88)
	if base.Ui32(v89+int32(-127)) < base.Ui32(int32(-95)) {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	v504 = v496
	goto L18
L22:
	;
	if v498 != 0 {
		v80 = v496
		v81 = v497
		v82 = v498
		goto L20
	} else {
		goto L137
	}
L23:
	;
	v496 = v488
	v497 = v81 + int32(1)
	v498 = v82 + int32(-1)
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v187
	v486 = F_sdscatprintf(m, v80, int32(_a1129), v11)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L7
	} else {
		goto L136
	}
L25:
	;
	v427 = int32(0)
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-1)))))
	switch v431 & int32(7) {
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
		v448 = v427
		goto L122
	}
L26:
	;
	v370 = int32(0)
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-1)))))
	switch v374 & int32(7) {
	case 0:
		goto L113
	case 1:
		goto L112
	case 2:
		goto L111
	case 3:
		goto L110
	case 4:
		goto L109
	default:
		v391 = v370
		goto L108
	}
L27:
	;
	v313 = int32(0)
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-1)))))
	switch v317 & int32(7) {
	case 0:
		goto L99
	case 1:
		goto L98
	case 2:
		goto L97
	case 3:
		goto L96
	case 4:
		goto L95
	default:
		v334 = v313
		goto L94
	}
L28:
	;
	v256 = int32(0)
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-1)))))
	switch v260 & int32(7) {
	case 0:
		goto L85
	case 1:
		goto L84
	case 2:
		goto L83
	case 3:
		goto L82
	case 4:
		goto L81
	default:
		v277 = v256
		goto L80
	}
L29:
	;
	v199 = int32(0)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-1)))))
	switch v203 & int32(7) {
	case 0:
		goto L71
	case 1:
		goto L70
	case 2:
		goto L69
	case 3:
		goto L68
	case 4:
		goto L67
	default:
		v220 = v199
		goto L66
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v89
	v197 = F_sdscatprintf(m, v80, int32(_a1130), v11+int32(16))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L7
	} else {
		goto L65
	}
L31:
	;
	v187 = v88 & int32(255)
	switch v187 + int32(-7) {
	case 0:
		goto L26
	case 1:
		goto L25
	case 2:
		goto L27
	case 3:
		goto L29
	case 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26:
		goto L24
	case 6:
		goto L28
	case 27:
		goto L30
	default:
		goto L63
	}
L32:
	;
	v95 = v88 & int32(255)
	if v95 == int32(34) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	if v95 == int32(92) {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v101 = v81
	v103 = v82
	goto L36
L35:
	;
	v130 = v127 - v81
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-1)))))
	switch v134 & int32(7) {
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
		v151 = int32(0)
		goto L45
	}
L36:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if base.Ui32(int32(-95)) <= base.Ui32(base.I32_extend8_s(v109)+int32(-127)) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v127 = v81 + v82
	v128 = int32(0)
	goto L35
L38:
	;
	v116 = v109 & int32(255)
	if v116 != int32(34) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v127 = v101
	v128 = v103
	goto L35
L40:
	;
	if v116 != int32(92) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v127 = v101
	v128 = v103
	goto L35
L42:
	;
	v124 = v103 + int32(-1)
	if v124 != 0 {
		v101 = v101 + int32(1)
		v103 = v124
		goto L36
	} else {
		goto L44
	}
L43:
	;
	v127 = v101
	v128 = v103
	goto L35
L44:
	;
	goto L37
L45:
	;
	v153 = F__sdsMakeRoomFor(m, v80, v130, int32(1))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L52
	}
L46:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-17))))
	v151 = v150
	goto L45
L47:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-9))))
	v151 = v147
	goto L45
L48:
	;
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+int32(-5)))))
	v151 = v144
	goto L45
L49:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-3)))))
	v151 = v141
	goto L45
L50:
	;
	v151 = int32(base.Ui32(v134) >> (uint(int32(3)) % 32))
	goto L45
L51:
	;
	if v130 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	if v153 != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v496 = int32(0)
	v497 = v127
	v498 = v128
	goto L22
L54:
	;
	v161 = v151 + v130
	v163 = v153 + int32(-1)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	switch v164 & int32(7) {
	case 0:
		goto L62
	case 1:
		goto L61
	case 2:
		goto L60
	case 3:
		goto L59
	case 4:
		goto L58
	default:
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v159 = F__emscripten_memcpy_bulkmem(m, v153+v151, v81, v130)
	mBase = m.M
	goto L55
L57:
	;
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v153+v161))) = uint8(v184)
	v496 = v153
	v497 = v127
	v498 = v128
	goto L22
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v153+int32(-17)))) = base.I64_extend_i32_u(v161)
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153+int32(-9)))) = v161
	goto L57
L60:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v153+int32(-5)))) = uint16(v161)
	goto L57
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v153+int32(-3)))) = uint8(v161)
	goto L57
L62:
	;
	v168 = v161 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v163))) = uint8(v168)
	goto L57
L63:
	;
	if v187 != int32(92) {
		goto L24
	} else {
		goto L64
	}
L64:
	;
	goto L30
L65:
	;
	v488 = v197
	goto L23
L66:
	;
	v223 = F__sdsMakeRoomFor(m, v80, int32(2), int32(1))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L7
	} else {
		goto L72
	}
L67:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-17))))
	v220 = v219
	goto L66
L68:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-9))))
	v220 = v216
	goto L66
L69:
	;
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+int32(-5)))))
	v220 = v213
	goto L66
L70:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-3)))))
	v220 = v210
	goto L66
L71:
	;
	v220 = int32(base.Ui32(v203) >> (uint(int32(3)) % 32))
	goto L66
L72:
	;
	if v223 == int32(0) {
		v488 = v199
		goto L23
	} else {
		goto L73
	}
L73:
	;
	v228 = int32(28252)
	*(*uint16)(unsafe.Add(mBase, uint32(v223+v220))) = uint16(v228)
	v231 = v220 + int32(2)
	v233 = v223 + int32(-1)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	switch v234 & int32(7) {
	case 0:
		goto L79
	case 1:
		goto L78
	case 2:
		goto L77
	case 3:
		goto L76
	case 4:
		goto L75
	default:
		goto L74
	}
L74:
	;
	v254 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v223+v231))) = uint8(v254)
	v488 = v223
	goto L23
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v223+int32(-17)))) = base.I64_extend_i32_u(v231)
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223+int32(-9)))) = v231
	goto L74
L77:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v223+int32(-5)))) = uint16(v231)
	goto L74
L78:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v223+int32(-3)))) = uint8(v231)
	goto L74
L79:
	;
	v238 = v231 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v238)
	goto L74
L80:
	;
	v280 = F__sdsMakeRoomFor(m, v80, int32(2), int32(1))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L7
	} else {
		goto L86
	}
L81:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-17))))
	v277 = v276
	goto L80
L82:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-9))))
	v277 = v273
	goto L80
L83:
	;
	v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+int32(-5)))))
	v277 = v270
	goto L80
L84:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-3)))))
	v277 = v267
	goto L80
L85:
	;
	v277 = int32(base.Ui32(v260) >> (uint(int32(3)) % 32))
	goto L80
L86:
	;
	if v280 == int32(0) {
		v488 = v256
		goto L23
	} else {
		goto L87
	}
L87:
	;
	v285 = int32(29276)
	*(*uint16)(unsafe.Add(mBase, uint32(v280+v277))) = uint16(v285)
	v288 = v277 + int32(2)
	v290 = v280 + int32(-1)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	switch v291 & int32(7) {
	case 0:
		goto L93
	case 1:
		goto L92
	case 2:
		goto L91
	case 3:
		goto L90
	case 4:
		goto L89
	default:
		goto L88
	}
L88:
	;
	v311 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v280+v288))) = uint8(v311)
	v488 = v280
	goto L23
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v280+int32(-17)))) = base.I64_extend_i32_u(v288)
	goto L88
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280+int32(-9)))) = v288
	goto L88
L91:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v280+int32(-5)))) = uint16(v288)
	goto L88
L92:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v280+int32(-3)))) = uint8(v288)
	goto L88
L93:
	;
	v295 = v288 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v295)
	goto L88
L94:
	;
	v337 = F__sdsMakeRoomFor(m, v80, int32(2), int32(1))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L7
	} else {
		goto L100
	}
L95:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-17))))
	v334 = v333
	goto L94
L96:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-9))))
	v334 = v330
	goto L94
L97:
	;
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+int32(-5)))))
	v334 = v327
	goto L94
L98:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-3)))))
	v334 = v324
	goto L94
L99:
	;
	v334 = int32(base.Ui32(v317) >> (uint(int32(3)) % 32))
	goto L94
L100:
	;
	if v337 == int32(0) {
		v488 = v313
		goto L23
	} else {
		goto L101
	}
L101:
	;
	v342 = int32(29788)
	*(*uint16)(unsafe.Add(mBase, uint32(v337+v334))) = uint16(v342)
	v345 = v334 + int32(2)
	v347 = v337 + int32(-1)
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347))))
	switch v348 & int32(7) {
	case 0:
		goto L107
	case 1:
		goto L106
	case 2:
		goto L105
	case 3:
		goto L104
	case 4:
		goto L103
	default:
		goto L102
	}
L102:
	;
	v368 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v337+v345))) = uint8(v368)
	v488 = v337
	goto L23
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v337+int32(-17)))) = base.I64_extend_i32_u(v345)
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v337+int32(-9)))) = v345
	goto L102
L105:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v337+int32(-5)))) = uint16(v345)
	goto L102
L106:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-3)))) = uint8(v345)
	goto L102
L107:
	;
	v352 = v345 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v347))) = uint8(v352)
	goto L102
L108:
	;
	v394 = F__sdsMakeRoomFor(m, v80, int32(2), int32(1))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L7
	} else {
		goto L114
	}
L109:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-17))))
	v391 = v390
	goto L108
L110:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-9))))
	v391 = v387
	goto L108
L111:
	;
	v384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+int32(-5)))))
	v391 = v384
	goto L108
L112:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-3)))))
	v391 = v381
	goto L108
L113:
	;
	v391 = int32(base.Ui32(v374) >> (uint(int32(3)) % 32))
	goto L108
L114:
	;
	if v394 == int32(0) {
		v488 = v370
		goto L23
	} else {
		goto L115
	}
L115:
	;
	v399 = int32(24924)
	*(*uint16)(unsafe.Add(mBase, uint32(v394+v391))) = uint16(v399)
	v402 = v391 + int32(2)
	v404 = v394 + int32(-1)
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	switch v405 & int32(7) {
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
		goto L116
	}
L116:
	;
	v425 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v394+v402))) = uint8(v425)
	v488 = v394
	goto L23
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v394+int32(-17)))) = base.I64_extend_i32_u(v402)
	goto L116
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v394+int32(-9)))) = v402
	goto L116
L119:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v394+int32(-5)))) = uint16(v402)
	goto L116
L120:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v394+int32(-3)))) = uint8(v402)
	goto L116
L121:
	;
	v409 = v402 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v404))) = uint8(v409)
	goto L116
L122:
	;
	v451 = F__sdsMakeRoomFor(m, v80, int32(2), int32(1))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L7
	} else {
		goto L128
	}
L123:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-17))))
	v448 = v447
	goto L122
L124:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-9))))
	v448 = v444
	goto L122
L125:
	;
	v441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+int32(-5)))))
	v448 = v441
	goto L122
L126:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-3)))))
	v448 = v438
	goto L122
L127:
	;
	v448 = int32(base.Ui32(v431) >> (uint(int32(3)) % 32))
	goto L122
L128:
	;
	if v451 == int32(0) {
		v488 = v427
		goto L23
	} else {
		goto L129
	}
L129:
	;
	v456 = int32(25180)
	*(*uint16)(unsafe.Add(mBase, uint32(v451+v448))) = uint16(v456)
	v459 = v448 + int32(2)
	v461 = v451 + int32(-1)
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	switch v462 & int32(7) {
	case 0:
		goto L135
	case 1:
		goto L134
	case 2:
		goto L133
	case 3:
		goto L132
	case 4:
		goto L131
	default:
		goto L130
	}
L130:
	;
	v482 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v451+v459))) = uint8(v482)
	v488 = v451
	goto L23
L131:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v451+int32(-17)))) = base.I64_extend_i32_u(v459)
	goto L130
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451+int32(-9)))) = v459
	goto L130
L133:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v451+int32(-5)))) = uint16(v459)
	goto L130
L134:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v451+int32(-3)))) = uint8(v459)
	goto L130
L135:
	;
	v466 = v459 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v461))) = uint8(v466)
	goto L130
L136:
	;
	v488 = v486
	goto L23
L137:
	;
	goto L21
L138:
	;
	v533 = int32(1)
	v535 = F__sdsMakeRoomFor(m, v504, v533, v533)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L7
	} else {
		goto L145
	}
L139:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v504+int32(-17))))
	v532 = v531
	goto L138
L140:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v504+int32(-9))))
	v532 = v528
	goto L138
L141:
	;
	v525 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v504+int32(-5)))))
	v532 = v525
	goto L138
L142:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504+int32(-3)))))
	v532 = v522
	goto L138
L143:
	;
	v532 = int32(base.Ui32(v515) >> (uint(int32(3)) % 32))
	goto L138
L144:
	;
	m.G0 = v11 + int32(32)
	return v535
L145:
	;
	if v535 == int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v540 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v535+v532))) = uint8(v540)
	v543 = v532 + int32(1)
	v545 = v535 + int32(-1)
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545))))
	switch v546 & int32(7) {
	case 0:
		goto L152
	case 1:
		goto L151
	case 2:
		goto L150
	case 3:
		goto L149
	case 4:
		goto L148
	default:
		goto L147
	}
L147:
	;
	v566 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v535+v543))) = uint8(v566)
	goto L144
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v535+int32(-17)))) = base.I64_extend_i32_u(v543)
	goto L147
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v535+int32(-9)))) = v543
	goto L147
L150:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v535+int32(-5)))) = uint16(v543)
	goto L147
L151:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v535+int32(-3)))) = uint8(v543)
	goto L147
L152:
	;
	v550 = v543 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v545))) = uint8(v550)
	goto L147
}
func F_sdscpy(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	if l1&int32(3) == int32(0) {
		v24 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v58 = F_sdscpylen(m, l0, l1, v57)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v57 = v49 - l1
	goto L1
L3:
	;
	v28 = v24
	goto L11
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v13 = l1
	goto L7
L6:
	;
	v57 = l1 - l1
	goto L1
L7:
	;
	v17 = v13 + int32(1)
	if v17&int32(3) == int32(0) {
		v24 = v17
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v22 != 0 {
		v13 = v17
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v49 = v17
	goto L2
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v37 = int32(-2139062144)
	if (int32(16843008)-v34|v34)&v37 == v37 {
		v28 = v28 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v43 = v28
	goto L14
L13:
	;
	goto L12
L14:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v47 != 0 {
		v43 = v43 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v49 = v43
	goto L2
L16:
	;
	goto L15
L17:
	;
	return int32(0)
L18:
	;
	return v58
}
func F_sdsfree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	if l0 == int32(0) {
		return
	} else {
		v11 = l0 + int32(-1)
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
		v14 = v12 & int32(7)
		switch v14 {
		case 0:
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v51 = v17 & int32(2147483647)
			v53 = v11
		case 1:
			v32 = int32(4)
			v33 = l0 + int32(-3)
			switch v14 {
			case 0:
				v49 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
			case 1:
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
				v49 = v39
			case 2:
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v49 = v42
			case 3:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
				v49 = v45
			case 4:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v49 = v48
			default:
				v49 = int32(0)
			}
			v51 = v49 + v32
			v53 = v33
		case 2:
			v32 = int32(6)
			v33 = l0 + int32(-5)
			switch v14 {
			case 0:
				v49 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
			case 1:
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
				v49 = v39
			case 2:
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v49 = v42
			case 3:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
				v49 = v45
			case 4:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v49 = v48
			default:
				v49 = int32(0)
			}
			v51 = v49 + v32
			v53 = v33
		case 3:
			v32 = int32(10)
			v33 = l0 + int32(-9)
			switch v14 {
			case 0:
				v49 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
			case 1:
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
				v49 = v39
			case 2:
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v49 = v42
			case 3:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
				v49 = v45
			case 4:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v49 = v48
			default:
				v49 = int32(0)
			}
			v51 = v49 + v32
			v53 = v33
		case 4:
			v32 = int32(18)
			v33 = l0 + int32(-17)
			switch v14 {
			case 0:
				v49 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
			case 1:
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
				v49 = v39
			case 2:
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v49 = v42
			case 3:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
				v49 = v45
			case 4:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v49 = v48
			default:
				v49 = int32(0)
			}
			v51 = v49 + v32
			v53 = v33
		default:
			v32 = int32(1)
			v33 = l0
			switch v14 {
			case 0:
				v49 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
			case 1:
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
				v49 = v39
			case 2:
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v49 = v42
			case 3:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
				v49 = v45
			case 4:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v49 = v48
			default:
				v49 = int32(0)
			}
			v51 = v49 + v32
			v53 = v33
		}
		F_zfree_with_size(m, v53, v51)
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			return
		}
	}
}
func F_sdsfreeVoid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	if l0 == int32(0) {
		return
	} else {
		v11 = l0 + int32(-1)
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
		v14 = v12 & int32(7)
		switch v14 {
		case 0:
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v51 = v17 & int32(2147483647)
			v53 = v11
		case 1:
			v32 = int32(4)
			v33 = l0 + int32(-3)
			switch v14 {
			case 0:
				v49 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
			case 1:
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
				v49 = v39
			case 2:
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v49 = v42
			case 3:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
				v49 = v45
			case 4:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v49 = v48
			default:
				v49 = int32(0)
			}
			v51 = v49 + v32
			v53 = v33
		case 2:
			v32 = int32(6)
			v33 = l0 + int32(-5)
			switch v14 {
			case 0:
				v49 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
			case 1:
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
				v49 = v39
			case 2:
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v49 = v42
			case 3:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
				v49 = v45
			case 4:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v49 = v48
			default:
				v49 = int32(0)
			}
			v51 = v49 + v32
			v53 = v33
		case 3:
			v32 = int32(10)
			v33 = l0 + int32(-9)
			switch v14 {
			case 0:
				v49 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
			case 1:
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
				v49 = v39
			case 2:
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v49 = v42
			case 3:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
				v49 = v45
			case 4:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v49 = v48
			default:
				v49 = int32(0)
			}
			v51 = v49 + v32
			v53 = v33
		case 4:
			v32 = int32(18)
			v33 = l0 + int32(-17)
			switch v14 {
			case 0:
				v49 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
			case 1:
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
				v49 = v39
			case 2:
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v49 = v42
			case 3:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
				v49 = v45
			case 4:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v49 = v48
			default:
				v49 = int32(0)
			}
			v51 = v49 + v32
			v53 = v33
		default:
			v32 = int32(1)
			v33 = l0
			switch v14 {
			case 0:
				v49 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
			case 1:
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
				v49 = v39
			case 2:
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v49 = v42
			case 3:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
				v49 = v45
			case 4:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v49 = v48
			default:
				v49 = int32(0)
			}
			v51 = v49 + v32
			v53 = v33
		}
		F_zfree_with_size(m, v53, v51)
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			return
		}
	}
}
func F_sdsfreesplitres(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v92 int32
	_ = v92
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L23
	} else {
		goto L26
	}
L4:
	;
	v14 = l1
	goto L5
L5:
	;
	v22 = v14 + int32(-1)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+v22<<(uint(int32(2))%32))))
	if v26 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	if v22 != 0 {
		v14 = v22
		goto L5
	} else {
		goto L25
	}
L8:
	;
	v31 = v26 + int32(-1)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v34 = v32 & int32(7)
	switch v34 {
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
		v52 = int32(1)
		v53 = v26
		goto L10
	}
L9:
	;
	F_zfree_with_size(m, v73, v71)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L23
	} else {
		goto L24
	}
L10:
	;
	switch v34 {
	case 0:
		goto L22
	case 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		v69 = int32(0)
		goto L17
	}
L11:
	;
	v52 = int32(18)
	v53 = v26 + int32(-17)
	goto L10
L12:
	;
	v52 = int32(10)
	v53 = v26 + int32(-9)
	goto L10
L13:
	;
	v52 = int32(6)
	v53 = v26 + int32(-5)
	goto L10
L14:
	;
	v52 = int32(4)
	v53 = v26 + int32(-3)
	goto L10
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(-9))))
	goto L16
L16:
	;
	v71 = v37 & int32(2147483647)
	v73 = v31
	goto L9
L17:
	;
	v71 = v69 + v52
	v73 = v53
	goto L9
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(-9))))
	v69 = v68
	goto L17
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(-5))))
	v69 = v65
	goto L17
L20:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26+int32(-3)))))
	v69 = v62
	goto L17
L21:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+int32(-2)))))
	v69 = v59
	goto L17
L22:
	;
	v69 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	goto L17
L23:
	;
	return
L24:
	;
	goto L7
L25:
	;
	goto L6
L26:
	;
	goto L1
}
func F_sdsfromlonglong(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	if l0 <= int64(-1) {
		v17 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v17)
		v21 = int32(1)
		v26 = v6 + v21
		v27 = int32(20)
		v28 = int64(0) - l0
		v29 = v21
	} else {
		v26 = v6
		v27 = int32(21)
		v28 = l0
		v29 = int32(0)
	}
	v30 = F_ull2string(m, v26, v27, v28)
	mBase = m.M
	if v30 == int32(0) {
		v49 = int32(0)
	} else {
		v49 = v30 + v29
	}
	v51 = F__sdsnewlen(m, v6, v49, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(32)
		return v51
	}
}
func F_sdsgrowzero(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v66 int32
	_ = v66
	v2 = l1
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v8 & int32(7) {
	case 0:
		v25 = int32(base.Ui32(v8) >> (uint(int32(3)) % 32))
	case 1:
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v25 = v15
	case 2:
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v25 = v18
	case 3:
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v25 = v21
	case 4:
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v25 = v24
	default:
		v25 = int32(0)
	}
	if base.Ui32(v2) <= base.Ui32(v25) {
		v66 = l0
		return v66
	} else {
		v27 = v2 - v25
		v29 = F__sdsMakeRoomFor(m, l0, v27, int32(1))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			if v29 != 0 {
				v40 = F__emscripten_memset_bulkmem(m, v29+v25, base.I32_extend8_s(int32(0)), v27+int32(1))
				mBase = m.M
				v42 = v29 + int32(-1)
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
				switch v43 & int32(7) {
				case 0:
					v47 = v2 << (uint(int32(3)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v47)
					return v29
				case 1:
					*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))) = uint8(v2)
					return v29
				case 2:
					*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))) = uint16(v2)
					return v29
				case 3:
					*(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9)))) = v2
					return v29
				case 4:
					*(*int64)(unsafe.Add(mBase, uint32(v29+int32(-17)))) = base.I64_extend_i32_u(v2)
					v66 = v29
					return v66
				default:
					v66 = v29
					return v66
				}
			} else {
				return int32(0)
			}
		}
	}
}
func F_sdsjoin(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_zmalloc_usable(m, int32(4), v10+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l1 < int32(1) {
		v67 = v36
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v24 = v20 + int32(-4)
	if base.Ui32(v24) < base.Ui32(int32(65531)) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	return int32(0)
L4:
	;
	if v15 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v36 = int32(0)
	goto L1
L6:
	;
	v27 = int32(2)
	goto L8
L7:
	;
	v27 = int32(3)
	goto L8
L8:
	;
	if base.Ui32(int32(255)) < base.Ui32(v24) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = v27
	goto L11
L10:
	;
	v31 = int32(1)
	goto L11
L11:
	;
	v34 = F_sdswrite(m, v15, v20, v31, int32(_a188), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v36 = v34
	goto L1
L13:
	;
	m.G0 = v10 + int32(16)
	return v67
L14:
	;
	v46 = int32(0)
	v47 = v36
	goto L15
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0+v46<<(uint(int32(2))%32))))
	v53 = F_sdscat(m, v47, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L17
	}
L16:
	;
	v67 = v58
	goto L13
L17:
	;
	if v46 == l1+int32(-1) {
		v58 = v53
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = v46 + int32(1)
	if v60 != l1 {
		v46 = v60
		v47 = v58
		goto L15
	} else {
		goto L21
	}
L19:
	;
	v56 = F_sdscat(m, v53, l2)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v58 = v56
	goto L18
L21:
	;
	goto L16
}
func F_sdssplitlen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v208 int32
	_ = v208
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	v6 = int32(0)
	if l1 < int32(1) {
		v303 = v6
		v304 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v304
	return v303
L2:
	;
	if l3 < int32(1) {
		v303 = v6
		v304 = v6
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v28 = int32(1)
	v29 = l1 - l3 + v28
	if v28 <= v29 {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	return int32(0)
L6:
	;
	if v21 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	return int32(0)
L8:
	;
	if v195 < int32(1) {
		goto L45
	} else {
		goto L46
	}
L9:
	;
	v180 = F__sdsnewlen(m, l0+v170, l1-v170, int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L5
	} else {
		goto L43
	}
L10:
	;
	v37 = int32(0)
	v45 = v37
	v46 = v37
	v47 = v21
	v49 = v37
	v51 = int32(5)
	goto L12
L11:
	;
	v32 = int32(0)
	v167 = v32
	v168 = v21
	v170 = v32
	goto L9
L12:
	;
	if v46+int32(2) <= v51 {
		v64 = v47
		v65 = v51
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v167 = v155
	v168 = v64
	v170 = v156
	goto L9
L14:
	;
	if l3 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v58 = F_valkey_realloc(m, v47, v51<<(uint(int32(3))%32))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	if v58 == int32(0) {
		v195 = v46
		v196 = v47
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v64 = v58
	v65 = v51 << (uint(int32(1)) % 32)
	goto L14
L18:
	;
	v159 = v157 + int32(1)
	if v159 < v29 {
		v45 = v159
		v46 = v155
		v47 = v64
		v49 = v156
		v51 = v65
		goto L12
	} else {
		goto L42
	}
L19:
	;
	v146 = F__sdsnewlen(m, l0+v49, v45-v49, int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L40
	}
L20:
	;
	v73 = l0 + v45
	if base.Ui32(l3) < base.Ui32(int32(4)) {
		v97 = v73
		v98 = l2
		v99 = l3
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v45))))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v70 == v71 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v137 == int32(0) {
		goto L19
	} else {
		goto L39
	}
L24:
	;
	v137 = int32(0)
	goto L23
L25:
	;
	v109 = v104
	v110 = v105
	v111 = v106
	goto L35
L26:
	;
	if v99 == int32(0) {
		goto L24
	} else {
		goto L33
	}
L27:
	;
	if (l2|v73)&int32(3) != 0 {
		v104 = v73
		v105 = l2
		v106 = l3
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v81 = v73
	v82 = l2
	v83 = l3
	goto L29
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v86 != v87 {
		v104 = v81
		v105 = v82
		v106 = v83
		goto L25
	} else {
		goto L31
	}
L30:
	;
	v97 = v92
	v98 = v90
	v99 = v94
	goto L26
L31:
	;
	v89 = int32(4)
	v90 = v82 + v89
	v92 = v81 + v89
	v94 = v83 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v94) {
		v81 = v92
		v82 = v90
		v83 = v94
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v104 = v97
	v105 = v98
	v106 = v99
	goto L25
L34:
	;
	v137 = v114 - v115
	goto L23
L35:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v114 != v115 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v117 = int32(1)
	v122 = v111 + int32(-1)
	if v122 == int32(0) {
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v109 = v109 + v117
	v110 = v110 + v117
	v111 = v122
	goto L35
L39:
	;
	v155 = v46
	v156 = v49
	v157 = v45
	goto L18
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64+v46<<(uint(int32(2))%32)))) = v146
	if v146 == int32(0) {
		v195 = v46
		v196 = v64
		goto L8
	} else {
		goto L41
	}
L41:
	;
	v155 = v46 + int32(1)
	v156 = v45 + l3
	v157 = v45 + (l3 + int32(-1))
	goto L18
L42:
	;
	goto L13
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168+v167<<(uint(int32(2))%32)))) = v180
	if v180 == int32(0) {
		v195 = v167
		v196 = v168
		goto L8
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v167 + int32(1)
	return v168
L45:
	;
	F_valkey_free(m, v196)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L5
	} else {
		goto L67
	}
L46:
	;
	v208 = int32(0)
	goto L47
L47:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v196+v208<<(uint(int32(2))%32))))
	if v221 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L45
L49:
	;
	v279 = v208 + int32(1)
	if v279 != v195 {
		v208 = v279
		goto L47
	} else {
		goto L66
	}
L50:
	;
	v226 = v221 + int32(-1)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	v229 = v227 & int32(7)
	switch v229 {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	case 4:
		goto L53
	default:
		v247 = int32(1)
		v248 = v221
		goto L52
	}
L51:
	;
	F_zfree_with_size(m, v268, v267)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L65
	}
L52:
	;
	switch v229 {
	case 0:
		goto L64
	case 1:
		goto L63
	case 2:
		goto L62
	case 3:
		goto L61
	case 4:
		goto L60
	default:
		v264 = int32(0)
		goto L59
	}
L53:
	;
	v247 = int32(18)
	v248 = v221 + int32(-17)
	goto L52
L54:
	;
	v247 = int32(10)
	v248 = v221 + int32(-9)
	goto L52
L55:
	;
	v247 = int32(6)
	v248 = v221 + int32(-5)
	goto L52
L56:
	;
	v247 = int32(4)
	v248 = v221 + int32(-3)
	goto L52
L57:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v221+int32(-9))))
	goto L58
L58:
	;
	v267 = v232 & int32(2147483647)
	v268 = v226
	goto L51
L59:
	;
	v267 = v264 + v247
	v268 = v248
	goto L51
L60:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v221+int32(-9))))
	v264 = v263
	goto L59
L61:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v221+int32(-5))))
	v264 = v260
	goto L59
L62:
	;
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221+int32(-3)))))
	v264 = v257
	goto L59
L63:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+int32(-2)))))
	v264 = v254
	goto L59
L64:
	;
	v264 = int32(base.Ui32(v227) >> (uint(int32(3)) % 32))
	goto L59
L65:
	;
	goto L49
L66:
	;
	goto L48
L67:
	;
	v296 = int32(0)
	v303 = v296
	v304 = v296
	goto L1
}
func F_sdstrim(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	v8 = l0 + int32(-1)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	switch v9 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v26 = int32(0)
		goto L1
	}
L1:
	;
	v29 = l0 + v26 + int32(-1)
	if base.Ui32(v29) < base.Ui32(l0) {
		v53 = l0
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v26 = v25
	goto L1
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v26 = v22
	goto L1
L4:
	;
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v26 = v19
	goto L1
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v26 = v16
	goto L1
L6:
	;
	v26 = int32(base.Ui32(v9) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	if base.Ui32(v29) <= base.Ui32(v53) {
		v75 = v29
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v35 = l0
	goto L9
L9:
	;
	v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35))))
	v37 = F___strchrnul(m, l1, v36)
	mBase = m.M
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v39 == v36&int32(255) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v53 = v47
	goto L7
L11:
	;
	if v43 == int32(0) {
		v53 = v35
		goto L7
	} else {
		goto L15
	}
L12:
	;
	v43 = v37
	goto L14
L13:
	;
	v43 = int32(0)
	goto L14
L14:
	;
	goto L11
L15:
	;
	v47 = v35 + int32(1)
	if base.Ui32(v47) <= base.Ui32(v29) {
		v35 = v47
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	v80 = v75 - v53 + int32(1)
	if l0 == v53 {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	v57 = v29
	goto L19
L19:
	;
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57))))
	v61 = F___strchrnul(m, l1, v60)
	mBase = m.M
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v63 == v60&int32(255) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v75 = v53
	goto L17
L21:
	;
	if v67 == int32(0) {
		v75 = v57
		goto L17
	} else {
		goto L25
	}
L22:
	;
	v67 = v61
	goto L24
L23:
	;
	v67 = int32(0)
	goto L24
L24:
	;
	goto L21
L25:
	;
	v71 = v57 + int32(-1)
	if base.Ui32(v53) < base.Ui32(v71) {
		v57 = v71
		goto L19
	} else {
		goto L26
	}
L26:
	;
	goto L20
L27:
	;
	v231 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v80))) = uint8(v231)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	switch v233 & int32(7) {
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
		goto L70
	}
L28:
	;
	if l0 == v53 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L27
L30:
	;
	goto L29
L31:
	;
	v85 = v80 + l0
	if base.Ui32(int32(0)-v80<<(uint(int32(1))%32)) < base.Ui32(v53-v85) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v95 = (v53 ^ l0) & int32(3)
	if base.Ui32(v53) <= base.Ui32(l0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v92 = F___memcpy(m, l0, v53, v80)
	mBase = m.M
	goto L29
L34:
	;
	if v201 == int32(0) {
		goto L30
	} else {
		goto L66
	}
L35:
	;
	if base.Ui32(v179) <= base.Ui32(int32(3)) {
		v200 = v178
		v201 = v179
		v202 = v180
		goto L34
	} else {
		goto L62
	}
L36:
	;
	if v95 != 0 {
		v161 = v80
		goto L46
	} else {
		goto L47
	}
L37:
	;
	if v95 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if l0&int32(3) != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v200 = v53
	v201 = v80
	v202 = l0
	goto L34
L40:
	;
	v102 = v53
	v103 = v80
	v104 = l0
	goto L42
L41:
	;
	v178 = v53
	v179 = v80
	v180 = l0
	goto L35
L42:
	;
	if v103 == int32(0) {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	*(*uint8)(unsafe.Add(mBase, uint32(v104))) = uint8(v108)
	v110 = int32(1)
	v111 = v102 + v110
	v113 = v103 + int32(-1)
	v115 = v104 + v110
	if v115&int32(3) == int32(0) {
		v178 = v111
		v179 = v113
		v180 = v115
		goto L35
	} else {
		goto L45
	}
L45:
	;
	v102 = v111
	v103 = v113
	v104 = v115
	goto L42
L46:
	;
	if v161 == int32(0) {
		goto L30
	} else {
		goto L58
	}
L47:
	;
	if v85&int32(3) == int32(0) {
		v141 = v80
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if base.Ui32(v141) <= base.Ui32(int32(3)) {
		v161 = v141
		goto L46
	} else {
		goto L54
	}
L49:
	;
	v126 = v80
	goto L50
L50:
	;
	if v126 == int32(0) {
		goto L30
	} else {
		goto L52
	}
L51:
	;
	v141 = v132
	goto L48
L52:
	;
	v132 = v126 + int32(-1)
	v133 = l0 + v132
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v132))))
	*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v135)
	if v133&int32(3) != 0 {
		v126 = v132
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v148 = v141
	goto L55
L55:
	;
	v152 = v148 + int32(-4)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v53+v152)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v152))) = v155
	if base.Ui32(int32(3)) < base.Ui32(v152) {
		v148 = v152
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v161 = v152
	goto L46
L57:
	;
	goto L56
L58:
	;
	v168 = v161
	goto L59
L59:
	;
	v172 = v168 + int32(-1)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v172))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v172))) = uint8(v175)
	if v172 != 0 {
		v168 = v172
		goto L59
	} else {
		goto L61
	}
L61:
	;
	goto L30
L62:
	;
	v185 = v178
	v186 = v179
	v187 = v180
	goto L63
L63:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v189
	v191 = int32(4)
	v192 = v185 + v191
	v194 = v187 + v191
	v196 = v186 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v196) {
		v185 = v192
		v186 = v196
		v187 = v194
		goto L63
	} else {
		goto L65
	}
L64:
	;
	v200 = v192
	v201 = v196
	v202 = v194
	goto L34
L65:
	;
	goto L64
L66:
	;
	v207 = v200
	v208 = v201
	v209 = v202
	goto L67
L67:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	*(*uint8)(unsafe.Add(mBase, uint32(v209))) = uint8(v211)
	v213 = int32(1)
	v218 = v208 + int32(-1)
	if v218 != 0 {
		v207 = v207 + v213
		v208 = v218
		v209 = v209 + v213
		goto L67
	} else {
		goto L69
	}
L68:
	;
	goto L30
L69:
	;
	goto L68
L70:
	;
	return l0
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(-17)))) = base.I64_extend_i32_u(v80)
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9)))) = v80
	return l0
L73:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))) = uint16(v80)
	return l0
L74:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))) = uint8(v80)
	return l0
L75:
	;
	v237 = v80 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v237)
	return l0
}
func F_sdstrynewlen(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F__sdsnewlen(m, l0, l1, int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_sendto(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v7 = m.Env.X__syscall_sendto(m, l0, l1, l2, l3, l4, l5)
	mBase = m.M
	if base.Ui32(v7) < base.Ui32(int32(-4095)) {
		v15 = v7
	} else {
		v10 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0) - v7
		v15 = int32(-1)
	}
	return v15
}
func F_setMigratingSlotDest(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v6 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
	v8 = F_dictFind(m, v7, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		if l1 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[90]))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
			if v8 == int32(0) {
				v23 = F_dictAdd(m, v19, l0, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
				return
			}
		} else {
			if v8 == int32(0) {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, _consts[90]))
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
				v15 = F_dictDelete(m, v14, l0)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_setNumericType(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	v2 = l1
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	switch v4 {
	case 0:
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*uint32)(unsafe.Add(mBase, uint32(v5))) = uint32(v2)
		return int32(1)
	case 1:
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v2)
		return int32(1)
	case 2:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*uint32)(unsafe.Add(mBase, uint32(v13))) = uint32(v2)
		return int32(1)
	case 3:
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*uint32)(unsafe.Add(mBase, uint32(v17))) = uint32(v2)
		return int32(1)
	case 4:
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
		if v21&int32(1) == int32(0) {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v32))) = v2
			return int32(1)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			v27 = F_setModuleNumericConfig(m, v26, v2, l2)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				return v27
			}
		}
	case 5:
		v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
		if v36&int32(1) == int32(0) {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v45))) = v2
			return int32(1)
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			v42 = F_setModuleUnsignedNumericConfig(m, v41, v2, l2)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				return v42
			}
		}
	case 6:
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*uint32)(unsafe.Add(mBase, uint32(v49))) = uint32(v2)
		return int32(1)
	case 7:
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*uint32)(unsafe.Add(mBase, uint32(v53))) = uint32(v2)
		return int32(1)
	case 8:
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int64)(unsafe.Add(mBase, uint32(v57))) = v2
		return int32(1)
	case 9:
		v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int64)(unsafe.Add(mBase, uint32(v61))) = v2
		return int32(1)
	default:
		return int32(1)
	}
}
func F_setOOMScoreAdj(m *base.Module, l0 int32) int32 {
	var v12 int32
	_ = v12
	if base.Ui32(l0+int32(1)) < base.Ui32(int32(4)) {
		return int32(-1)
	} else {
		F__serverAssert(m, int32(_a1343), int32(_a1240), int32(2477))
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_setexCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v6&int32(1) != 0 {
		v13 = v4
		v14 = v5
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		v18 = int32(0)
		F_setGenericCommand(m, l0, int32(1028), v16, v14, v17, v18, v18, v18, v18)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			return
		}
	} else {
		v9 = F_tryObjectEncoding(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v9
			v13 = v11
			v14 = v9
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			v18 = int32(0)
			F_setGenericCommand(m, l0, int32(1028), v16, v14, v17, v18, v18, v18, v18)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_setlocale(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int64
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v297 int64
	_ = v297
	var v302 int64
	_ = v302
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int64
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if base.Ui32(int32(6)) < base.Ui32(l0) {
		v559 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v559
L2:
	;
	goto L3
L3:
	;
	if l0 != int32(6) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v458 = int32(0)
	v460 = v458
	v461 = int32(9117040)
	v465 = v458
	goto L142
L5:
	;
	if l1 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L6:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v23 = int32(0)
	v24 = *(*int64)(unsafe.Add(mBase, _consts[1248]))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(16)))) = v24
	v27 = *(*int64)(unsafe.Add(mBase, _consts[1249]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v27
	v30 = *(*int64)(unsafe.Add(mBase, _consts[1250]))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v30
	v33 = v23
	v34 = l1
	goto L9
L8:
	;
	goto L87
L9:
	;
	goto L17
L10:
	;
	v291 = int32(0)
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
	*(*int64)(unsafe.Add(mBase, _consts[1251])) = v292
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(40))))
	*(*int64)(unsafe.Add(mBase, _consts[1252])) = v297
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(32))))
	*(*int64)(unsafe.Add(mBase, _consts[1253])) = v302
	goto L4
L11:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v160 != 0 {
		v180 = v11
		goto L42
	} else {
		goto L43
	}
L12:
	;
	v142 = v131 - v34
	if int32(23) < v142 {
		v156 = v34
		goto L11
	} else {
		goto L34
	}
L13:
	;
	goto L12
L14:
	;
	v122 = v117
	goto L30
L15:
	;
	v117 = v108
	goto L14
L17:
	;
	if v34&int32(3) == int32(0) {
		v68 = v34
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v77 = int32(-2139062144)
	if (int32(16843008)-v74|v74)&v77 != v77 {
		v108 = v68
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v55 = v34
	goto L20
L20:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v60 == int32(0) {
		v131 = v55
		goto L13
	} else {
		goto L22
	}
L21:
	;
	v68 = v65
	goto L18
L22:
	;
	if v60 == int32(59) {
		v131 = v55
		goto L13
	} else {
		goto L23
	}
L23:
	;
	v65 = v55 + int32(1)
	if v65&int32(3) != 0 {
		v55 = v65
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v83 = v68
	v86 = v74
	goto L26
L26:
	;
	v89 = v86 ^ int32(993737531)
	v92 = int32(-2139062144)
	if (int32(16843008)-v89|v89)&v92 != v92 {
		v108 = v83
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v98 = v83 + int32(4)
	v102 = int32(-2139062144)
	if (v96|(int32(16843008)-v96))&v102 == v102 {
		v83 = v98
		v86 = v96
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v117 = v98
	goto L14
L30:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v123 == int32(0) {
		v131 = v122
		goto L13
	} else {
		goto L32
	}
L31:
	;
	v131 = v122
	goto L13
L32:
	;
	if v123 != int32(59) {
		v122 = v122 + int32(1)
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if v142 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11+v142))) = uint8(v150)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v154 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L35
L37:
	;
	v147 = F__emscripten_memcpy_bulkmem(m, v11, v34, v142)
	mBase = m.M
	goto L36
L38:
	;
	v155 = v131 + int32(1)
	goto L40
L39:
	;
	v155 = v34
	goto L40
L40:
	;
	v156 = v155
	goto L11
L41:
	;
	if v278 == int32(-1) {
		goto L8
	} else {
		goto L85
	}
L42:
	;
	v184 = int32(0)
	goto L55
L43:
	;
	v162 = F_getenv(m, int32(_a2355))
	mBase = m.M
	if v162 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v170 = F_getenv(m, v33*int32(12)+int32(_a2356))
	mBase = m.M
	if v170 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v165 != 0 {
		v180 = v162
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v175 = F_getenv(m, int32(_a2357))
	mBase = m.M
	if v175 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v173 != 0 {
		v180 = v170
		goto L42
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v180 = int32(_a2358)
	goto L42
L51:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v178 != 0 {
		v180 = v175
		goto L42
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v200 = int32(_a2358)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v201 == int32(46) {
		v208 = v200
		goto L64
	} else {
		goto L65
	}
L54:
	;
	v199 = v184
	goto L53
L55:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+v184))))
	if v188 == int32(0) {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	if v188 == int32(47) {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v193 = int32(23)
	v195 = v184 + int32(1)
	if v195 != v193 {
		v184 = v195
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v199 = v193
	goto L53
L60:
	;
	v278 = v270
	goto L41
L61:
	;
	v226 = int32(0)
	v227 = *(*int32)(unsafe.Add(mBase, _consts[1254]))
	if v227 == v226 {
		goto L74
	} else {
		goto L75
	}
L62:
	;
	if v33 != 0 {
		goto L71
	} else {
		goto L72
	}
L63:
	;
	v214 = F_strcmp(m, v212, int32(_a2358))
	mBase = m.M
	if v214 == int32(0) {
		v219 = v212
		goto L62
	} else {
		goto L69
	}
L64:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	if v209 == int32(0) {
		v219 = v208
		goto L62
	} else {
		goto L68
	}
L65:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+v199))))
	if v205 != 0 {
		v208 = v200
		goto L64
	} else {
		goto L66
	}
L66:
	;
	if v201 != int32(67) {
		v212 = v180
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v208 = v180
	goto L64
L68:
	;
	v212 = v208
	goto L63
L69:
	;
	v218 = F_strcmp(m, v212, int32(_a2359))
	mBase = m.M
	if v218 != 0 {
		goto L61
	} else {
		goto L70
	}
L70:
	;
	v219 = v212
	goto L62
L71:
	;
	v278 = int32(0)
	goto L41
L72:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
	if v221 == int32(46) {
		v270 = int32(_a2360)
		goto L60
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v247 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v247 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L75:
	;
	v232 = v227
	goto L76
L76:
	;
	v237 = F_strcmp(m, v212, v232+int32(8))
	mBase = m.M
	if v237 == int32(0) {
		v270 = v232
		goto L60
	} else {
		goto L78
	}
L77:
	;
	goto L74
L78:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v232)+32))
	if v240 != 0 {
		v232 = v240
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	if v33|v247 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v250 = int32(0)
	v251 = *(*int64)(unsafe.Add(mBase, _consts[1255]))
	*(*int64)(unsafe.Add(mBase, uint32(v247))) = v251
	v254 = v247 + int32(8)
	v255 = F___memcpy(m, v254, v212, v199)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v254+v199))) = uint8(v250)
	v260 = *(*int32)(unsafe.Add(mBase, _consts[1254]))
	*(*int32)(unsafe.Add(mBase, uint32(v247)+32)) = v260
	*(*int32)(unsafe.Add(mBase, _consts[1254])) = v247
	goto L80
L82:
	;
	v267 = v247
	goto L84
L83:
	;
	v267 = int32(_a2360)
	goto L84
L84:
	;
	v270 = v267
	goto L60
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(24)+v33<<(uint(int32(2))%32)))) = v278
	v288 = v33 + int32(1)
	if v288 != int32(6) {
		v33 = v288
		v34 = v156
		goto L9
	} else {
		goto L86
	}
L86:
	;
	goto L10
L87:
	;
	v559 = int32(0)
	goto L1
L88:
	;
	goto L138
L89:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1251])))
	v443 = v442
	goto L88
L90:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v311 != 0 {
		v331 = l1
		goto L93
	} else {
		goto L94
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1251]))) = v429
	v443 = v429
	goto L88
L92:
	;
	if v429 != int32(-1) {
		goto L91
	} else {
		goto L136
	}
L93:
	;
	v335 = int32(0)
	goto L106
L94:
	;
	v313 = F_getenv(m, int32(_a2355))
	mBase = m.M
	if v313 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v321 = F_getenv(m, l0*int32(12)+int32(_a2356))
	mBase = m.M
	if v321 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
	if v316 != 0 {
		v331 = v313
		goto L93
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v326 = F_getenv(m, int32(_a2357))
	mBase = m.M
	if v326 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if v324 != 0 {
		v331 = v321
		goto L93
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v331 = int32(_a2358)
	goto L93
L102:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	if v329 != 0 {
		v331 = v326
		goto L93
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v351 = int32(_a2358)
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v352 == int32(46) {
		v359 = v351
		goto L115
	} else {
		goto L116
	}
L105:
	;
	v350 = v335
	goto L104
L106:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331+v335))))
	if v339 == int32(0) {
		goto L105
	} else {
		goto L108
	}
L108:
	;
	if v339 == int32(47) {
		goto L105
	} else {
		goto L109
	}
L109:
	;
	v344 = int32(23)
	v346 = v335 + int32(1)
	if v346 != v344 {
		v335 = v346
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v350 = v344
	goto L104
L111:
	;
	v429 = v421
	goto L92
L112:
	;
	v377 = int32(0)
	v378 = *(*int32)(unsafe.Add(mBase, _consts[1254]))
	if v378 == v377 {
		goto L125
	} else {
		goto L126
	}
L113:
	;
	if l0 != 0 {
		goto L122
	} else {
		goto L123
	}
L114:
	;
	v365 = F_strcmp(m, v363, int32(_a2358))
	mBase = m.M
	if v365 == int32(0) {
		v370 = v363
		goto L113
	} else {
		goto L120
	}
L115:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+1)))
	if v360 == int32(0) {
		v370 = v359
		goto L113
	} else {
		goto L119
	}
L116:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331+v350))))
	if v356 != 0 {
		v359 = v351
		goto L115
	} else {
		goto L117
	}
L117:
	;
	if v352 != int32(67) {
		v363 = v331
		goto L114
	} else {
		goto L118
	}
L118:
	;
	v359 = v331
	goto L115
L119:
	;
	v363 = v359
	goto L114
L120:
	;
	v369 = F_strcmp(m, v363, int32(_a2359))
	mBase = m.M
	if v369 != 0 {
		goto L112
	} else {
		goto L121
	}
L121:
	;
	v370 = v363
	goto L113
L122:
	;
	v429 = int32(0)
	goto L92
L123:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+1)))
	if v372 == int32(46) {
		v421 = int32(_a2360)
		goto L111
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	v398 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v398 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L126:
	;
	v383 = v378
	goto L127
L127:
	;
	v388 = F_strcmp(m, v363, v383+int32(8))
	mBase = m.M
	if v388 == int32(0) {
		v421 = v383
		goto L111
	} else {
		goto L129
	}
L128:
	;
	goto L125
L129:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v383)+32))
	if v391 != 0 {
		v383 = v391
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	if l0|v398 != 0 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v401 = int32(0)
	v402 = *(*int64)(unsafe.Add(mBase, _consts[1255]))
	*(*int64)(unsafe.Add(mBase, uint32(v398))) = v402
	v405 = v398 + int32(8)
	v406 = F___memcpy(m, v405, v363, v350)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v405+v350))) = uint8(v401)
	v411 = *(*int32)(unsafe.Add(mBase, _consts[1254]))
	*(*int32)(unsafe.Add(mBase, uint32(v398)+32)) = v411
	*(*int32)(unsafe.Add(mBase, _consts[1254])) = v398
	goto L131
L133:
	;
	v418 = v398
	goto L135
L134:
	;
	v418 = int32(_a2360)
	goto L135
L135:
	;
	v421 = v418
	goto L111
L136:
	;
	goto L137
L137:
	;
	v559 = v3
	goto L1
L138:
	;
	if v443 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v448 = v443 + int32(8)
	goto L141
L140:
	;
	v448 = int32(_a2246)
	goto L141
L141:
	;
	v559 = v448
	goto L1
L142:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v460<<(uint(int32(2))%32))+uint32(_consts[1251])))
	if v472 != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v549 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v538))) = uint8(v549)
	goto L167
L144:
	;
	v476 = v472 + int32(8)
	goto L146
L145:
	;
	v476 = int32(_a2246)
	goto L146
L146:
	;
	v477 = int32(0)
	v478 = *(*int32)(unsafe.Add(mBase, _consts[1251]))
	if v476&int32(3) == v477 {
		v500 = v476
		goto L149
	} else {
		goto L150
	}
L147:
	;
	if v533 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L148:
	;
	v533 = v525 - v476
	goto L147
L149:
	;
	v504 = v500
	goto L157
L150:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
	if v486 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v489 = v476
	goto L153
L152:
	;
	v533 = v476 - v476
	goto L147
L153:
	;
	v493 = v489 + int32(1)
	if v493&int32(3) == int32(0) {
		v500 = v493
		goto L149
	} else {
		goto L155
	}
L155:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
	if v498 != 0 {
		v489 = v493
		goto L153
	} else {
		goto L156
	}
L156:
	;
	v525 = v493
	goto L148
L157:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	v513 = int32(-2139062144)
	if (int32(16843008)-v510|v510)&v513 == v513 {
		v504 = v504 + int32(4)
		goto L157
	} else {
		goto L159
	}
L158:
	;
	v519 = v504
	goto L160
L159:
	;
	goto L158
L160:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
	if v523 != 0 {
		v519 = v519 + int32(1)
		goto L160
	} else {
		goto L162
	}
L161:
	;
	v525 = v519
	goto L148
L162:
	;
	goto L161
L163:
	;
	v538 = v461 + v533
	v539 = int32(59)
	*(*uint8)(unsafe.Add(mBase, uint32(v538))) = uint8(v539)
	v541 = int32(1)
	v544 = v465 + base.B2i32(v472 == v478)
	v546 = v460 + v541
	if v546 != int32(6) {
		v460 = v546
		v461 = v538 + v541
		v465 = v544
		goto L142
	} else {
		goto L166
	}
L164:
	;
	goto L163
L165:
	;
	v536 = F__emscripten_memcpy_bulkmem(m, v461, v476, v533)
	mBase = m.M
	goto L164
L166:
	;
	goto L143
L167:
	;
	if v544 == int32(6) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v555 = v476
	goto L170
L169:
	;
	v555 = int32(9117040)
	goto L170
L170:
	;
	v559 = v555
	goto L1
}
func F_setnxCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v6&int32(1) != 0 {
		v13 = v4
		v14 = v5
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v17 = int32(0)
		v19 = int32(_a918)
		v20 = *(*int32)(unsafe.Add(mBase, _consts[234]))
		v22 = *(*int32)(unsafe.Add(mBase, _consts[233]))
		F_setGenericCommand(m, l0, int32(1), v16, v14, v17, v17, v20, v22, v17)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			return
		}
	} else {
		v9 = F_tryObjectEncoding(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v9
			v13 = v11
			v14 = v9
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v17 = int32(0)
			v19 = int32(_a918)
			v20 = *(*int32)(unsafe.Add(mBase, _consts[234]))
			v22 = *(*int32)(unsafe.Add(mBase, _consts[233]))
			F_setGenericCommand(m, l0, int32(1), v16, v14, v17, v17, v20, v22, v17)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_setrangeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int64
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int64
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v209 int64
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
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
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int64
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v17 = F_objectGetVal(m, v16)
	mBase = m.M
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v23 = F_getLongFromObjectOrReply(m, l0, v19, v13+int32(8), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return
L2:
	;
	return
L3:
	;
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if int32(-1) < v25 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v34 = F_lookupKeyWrite(m, v31, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L10
	}
L6:
	;
	F_addReplyError(m, l0, int32(_a1592))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	v225 = F_objectGetVal(m, v220)
	mBase = m.M
	v226 = int32(0)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v230 = v17 + int32(-1)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	switch v231 & int32(7) {
	case 0:
		goto L88
	case 1:
		goto L87
	case 2:
		goto L86
	case 3:
		goto L85
	case 4:
		goto L84
	default:
		v248 = v226
		goto L83
	}
L9:
	;
	v142 = F_checkType(m, l0, v34, int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L50
	}
L10:
	;
	if v34 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v37 = v17 + int32(-1)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v40 = v38 & int32(7)
	switch v40 {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	default:
		goto L13
	}
L12:
	;
	switch v40 {
	default:
		goto L27
	case 1:
		goto L26
	case 2:
		goto L25
	case 3:
		goto L24
	case 4:
		goto L23
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[233]))
	F_addReply(m, l0, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L2
	} else {
		goto L21
	}
L14:
	;
	if v55 != 0 {
		goto L12
	} else {
		goto L20
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-17))))
	v55 = v54
	goto L14
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-9))))
	v55 = v51
	goto L14
L17:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(-5)))))
	v55 = v48
	goto L14
L18:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-3)))))
	v55 = v45
	goto L14
L19:
	;
	v55 = int32(base.Ui32(v38) >> (uint(int32(3)) % 32))
	goto L14
L20:
	;
	goto L13
L21:
	;
	goto L1
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v78 != int64(-1) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-17))))
	v75 = v74
	goto L22
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-9))))
	v75 = v71
	goto L22
L25:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(-5)))))
	v75 = v68
	goto L22
L26:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-3)))))
	v75 = v65
	goto L22
L27:
	;
	v75 = int32(base.Ui32(v38) >> (uint(int32(3)) % 32))
	goto L22
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	switch v107 & int32(7) {
	case 0:
		goto L46
	case 1:
		goto L45
	case 2:
		goto L44
	case 3:
		goto L43
	case 4:
		goto L42
	default:
		v124 = int32(0)
		goto L41
	}
L29:
	;
	if v93 != 0 {
		goto L28
	} else {
		goto L36
	}
L30:
	;
	v82 = int32(1)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v83&v82 != 0 {
		v90 = v82
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v93 = int32(1)
	goto L29
L32:
	;
	v93 = v90
	goto L29
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v86 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v88 = F_isImportSlotMigrationJob(m, v86)
	mBase = m.M
	v90 = v88
	goto L32
L35:
	;
	v93 = int32(0)
	goto L29
L36:
	;
	if v76 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_addReplyError(m, l0, int32(_a1591))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L40
	}
L38:
	;
	v100 = *(*int64)(unsafe.Add(mBase, _consts[1093]))
	if base.I64_extend_i32_u(v75)+base.I64_extend_i32_s(v76) <= v100 {
		goto L28
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L1
L41:
	;
	v125 = int32(0)
	v128 = F_sdsnewlen(m, v125, v124+v105)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L47
	}
L42:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-17))))
	v124 = v123
	goto L41
L43:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-9))))
	v124 = v120
	goto L41
L44:
	;
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(-5)))))
	v124 = v117
	goto L41
L45:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-3)))))
	v124 = v114
	goto L41
L46:
	;
	v124 = int32(base.Ui32(v107) >> (uint(int32(3)) % 32))
	goto L41
L47:
	;
	v130 = F_createObject(m, v125, v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v130
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	F_dbAdd(m, v133, v135, v13+int32(12))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v220 = v140
	goto L8
L50:
	;
	if v142 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v144 = F_stringObjectLen(m, v34)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-1)))))
	v150 = v148 & int32(7)
	switch v150 {
	case 0:
		goto L60
	case 1:
		goto L59
	case 2:
		goto L58
	case 3:
		goto L57
	case 4:
		goto L56
	default:
		goto L54
	}
L53:
	;
	switch v150 {
	default:
		goto L68
	case 1:
		goto L67
	case 2:
		goto L66
	case 3:
		goto L65
	case 4:
		goto L64
	}
L54:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v144))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L62
	}
L55:
	;
	if v165 != 0 {
		goto L53
	} else {
		goto L61
	}
L56:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-17))))
	v165 = v164
	goto L55
L57:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-9))))
	v165 = v161
	goto L55
L58:
	;
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(-5)))))
	v165 = v158
	goto L55
L59:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-3)))))
	v165 = v155
	goto L55
L60:
	;
	v165 = int32(base.Ui32(v148) >> (uint(int32(3)) % 32))
	goto L55
L61:
	;
	goto L54
L62:
	;
	goto L1
L63:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v187 != int64(-1) {
		goto L71
	} else {
		goto L72
	}
L64:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-17))))
	v184 = v183
	goto L63
L65:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-9))))
	v184 = v180
	goto L63
L66:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(-5)))))
	v184 = v177
	goto L63
L67:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-3)))))
	v184 = v174
	goto L63
L68:
	;
	v184 = int32(base.Ui32(v148) >> (uint(int32(3)) % 32))
	goto L63
L69:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v217 = F_dbUnshareStringValue(m, v214, v216, v34)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L2
	} else {
		goto L82
	}
L70:
	;
	if v202 != 0 {
		goto L69
	} else {
		goto L77
	}
L71:
	;
	v191 = int32(1)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v192&v191 != 0 {
		v199 = v191
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v202 = int32(1)
	goto L70
L73:
	;
	v202 = v199
	goto L70
L74:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v195 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v197 = F_isImportSlotMigrationJob(m, v195)
	mBase = m.M
	v199 = v197
	goto L73
L76:
	;
	v202 = int32(0)
	goto L70
L77:
	;
	if v185 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	F_addReplyError(m, l0, int32(_a1591))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L81
	}
L79:
	;
	v209 = *(*int64)(unsafe.Add(mBase, _consts[1093]))
	if base.I64_extend_i32_u(v184)+base.I64_extend_i32_s(v185) <= v209 {
		goto L69
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	goto L1
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v217
	v220 = v217
	goto L8
L83:
	;
	v250 = F_sdsgrowzero(m, v225, v248+v227)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L2
	} else {
		goto L89
	}
L84:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-17))))
	v248 = v247
	goto L83
L85:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-9))))
	v248 = v244
	goto L83
L86:
	;
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(-5)))))
	v248 = v241
	goto L83
L87:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-3)))))
	v248 = v238
	goto L83
L88:
	;
	v248 = int32(base.Ui32(v231) >> (uint(int32(3)) % 32))
	goto L83
L89:
	;
	F_objectSetVal(m, v220, v250)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L90
	}
L90:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v255 = F_objectGetVal(m, v254)
	mBase = m.M
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	switch v258 & int32(7) {
	case 0:
		goto L96
	case 1:
		goto L95
	case 2:
		goto L94
	case 3:
		goto L93
	case 4:
		goto L92
	default:
		v275 = v226
		goto L91
	}
L91:
	;
	if v275 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L92:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-17))))
	v275 = v274
	goto L91
L93:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-9))))
	v275 = v271
	goto L91
L94:
	;
	v268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(-5)))))
	v275 = v268
	goto L91
L95:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-3)))))
	v275 = v265
	goto L91
L96:
	;
	v275 = int32(base.Ui32(v258) >> (uint(int32(3)) % 32))
	goto L91
L97:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	F_signalModifiedKey(m, l0, v280, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L2
	} else {
		goto L100
	}
L98:
	;
	goto L97
L99:
	;
	v278 = F__emscripten_memcpy_bulkmem(m, v255+v256, v17, v275)
	mBase = m.M
	goto L98
L100:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+28))
	F_notifyKeyspaceEvent(m, int32(8), int32(_a1590), v288, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L2
	} else {
		goto L101
	}
L101:
	;
	v293 = int32(_a69)
	v295 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v295 + int64(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v301 = F_objectGetVal(m, v300)
	mBase = m.M
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301+int32(-1)))))
	switch v304 & int32(7) {
	case 0:
		goto L107
	case 1:
		goto L106
	case 2:
		goto L105
	case 3:
		goto L104
	case 4:
		goto L103
	default:
		v321 = int32(0)
		goto L102
	}
L102:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v321))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L2
	} else {
		goto L108
	}
L103:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v301+int32(-17))))
	v321 = v320
	goto L102
L104:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v301+int32(-9))))
	v321 = v317
	goto L102
L105:
	;
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v301+int32(-5)))))
	v321 = v314
	goto L102
L106:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301+int32(-3)))))
	v321 = v311
	goto L102
L107:
	;
	v321 = int32(base.Ui32(v304) >> (uint(int32(3)) % 32))
	goto L102
L108:
	;
	goto L1
}
func F_shutdownCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
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
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
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
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(2) <= v11 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v451 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L114
	} else {
		goto L149
	}
L3:
	;
	if v399 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L4:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
	if v389&int32(16) == int32(0) {
		v399 = v385
		v400 = v386
		goto L3
	} else {
		goto L128
	}
L5:
	;
	v16 = int32(0)
	v21 = int32(1)
	v22 = v16
	v23 = v16
	goto L7
L6:
	;
	v385 = int32(1)
	v386 = int32(0)
	goto L4
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = v21 << (uint(int32(2)) % 32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v27)))
	v30 = F_objectGetVal(m, v29)
	mBase = m.M
	v31 = int32(_a490)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	if v342 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L9:
	;
	v344 = v21 + int32(1)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v344 < v345 {
		v21 = v344
		v22 = v341
		v23 = v342
		goto L7
	} else {
		goto L107
	}
L10:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72+v27)))
	v75 = F_objectGetVal(m, v74)
	mBase = m.M
	v76 = int32(_a491)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v79 != 0 {
		goto L27
	} else {
		goto L28
	}
L11:
	;
	if v66-v68 != 0 {
		goto L10
	} else {
		goto L23
	}
L12:
	;
	v66 = F_tolower(m, v62)
	mBase = m.M
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v68 = F_tolower(m, v67)
	mBase = m.M
	goto L11
L13:
	;
	v36 = v30
	v37 = v31
	v38 = v34
	goto L16
L14:
	;
	v62 = int32(0)
	v63 = v31
	goto L12
L15:
	;
	v62 = v59 & int32(255)
	v63 = v58
	goto L12
L16:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v40 == int32(0) {
		v58 = v37
		v59 = v38
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v58 = v52
	v59 = int32(0)
	goto L15
L18:
	;
	v44 = v38 & int32(255)
	if v44 == v40 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v51 = int32(1)
	v52 = v37 + v51
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v53 != 0 {
		v36 = v36 + v51
		v37 = v52
		v38 = v53
		goto L16
	} else {
		goto L22
	}
L20:
	;
	v46 = F_tolower(m, v44)
	mBase = m.M
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v48 = F_tolower(m, v47)
	mBase = m.M
	if v46 == v48 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v58 = v37
	v59 = v50
	goto L15
L22:
	;
	goto L17
L23:
	;
	v341 = v22 | int32(2)
	v342 = v23
	goto L9
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v117+v27)))
	v120 = F_objectGetVal(m, v119)
	mBase = m.M
	v121 = int32(_a492)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v124 != 0 {
		goto L41
	} else {
		goto L42
	}
L25:
	;
	if v111-v113 != 0 {
		goto L24
	} else {
		goto L37
	}
L26:
	;
	v111 = F_tolower(m, v107)
	mBase = m.M
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v113 = F_tolower(m, v112)
	mBase = m.M
	goto L25
L27:
	;
	v81 = v75
	v82 = v76
	v83 = v79
	goto L30
L28:
	;
	v107 = int32(0)
	v108 = v76
	goto L26
L29:
	;
	v107 = v104 & int32(255)
	v108 = v103
	goto L26
L30:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v85 == int32(0) {
		v103 = v82
		v104 = v83
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v103 = v97
	v104 = int32(0)
	goto L29
L32:
	;
	v89 = v83 & int32(255)
	if v89 == v85 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = int32(1)
	v97 = v82 + v96
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	if v98 != 0 {
		v81 = v81 + v96
		v82 = v97
		v83 = v98
		goto L30
	} else {
		goto L36
	}
L34:
	;
	v91 = F_tolower(m, v89)
	mBase = m.M
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v93 = F_tolower(m, v92)
	mBase = m.M
	if v91 == v93 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v103 = v82
	v104 = v95
	goto L29
L36:
	;
	goto L31
L37:
	;
	v341 = v22 | int32(1)
	v342 = v23
	goto L9
L38:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162+v27)))
	v165 = F_objectGetVal(m, v164)
	mBase = m.M
	v166 = int32(_a493)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v169 != 0 {
		goto L55
	} else {
		goto L56
	}
L39:
	;
	if v156-v158 != 0 {
		goto L38
	} else {
		goto L51
	}
L40:
	;
	v156 = F_tolower(m, v152)
	mBase = m.M
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v158 = F_tolower(m, v157)
	mBase = m.M
	goto L39
L41:
	;
	v126 = v120
	v127 = v121
	v128 = v124
	goto L44
L42:
	;
	v152 = int32(0)
	v153 = v121
	goto L40
L43:
	;
	v152 = v149 & int32(255)
	v153 = v148
	goto L40
L44:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v130 == int32(0) {
		v148 = v127
		v149 = v128
		goto L43
	} else {
		goto L46
	}
L45:
	;
	v148 = v142
	v149 = int32(0)
	goto L43
L46:
	;
	v134 = v128 & int32(255)
	if v134 == v130 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v141 = int32(1)
	v142 = v127 + v141
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	if v143 != 0 {
		v126 = v126 + v141
		v127 = v142
		v128 = v143
		goto L44
	} else {
		goto L50
	}
L48:
	;
	v136 = F_tolower(m, v134)
	mBase = m.M
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v138 = F_tolower(m, v137)
	mBase = m.M
	if v136 == v138 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v148 = v127
	v149 = v140
	goto L43
L50:
	;
	goto L45
L51:
	;
	v341 = v22 | int32(4)
	v342 = v23
	goto L9
L52:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v207+v27)))
	v210 = F_objectGetVal(m, v209)
	mBase = m.M
	v211 = int32(_a494)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if v214 != 0 {
		goto L69
	} else {
		goto L70
	}
L53:
	;
	if v201-v203 != 0 {
		goto L52
	} else {
		goto L65
	}
L54:
	;
	v201 = F_tolower(m, v197)
	mBase = m.M
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	v203 = F_tolower(m, v202)
	mBase = m.M
	goto L53
L55:
	;
	v171 = v165
	v172 = v166
	v173 = v169
	goto L58
L56:
	;
	v197 = int32(0)
	v198 = v166
	goto L54
L57:
	;
	v197 = v194 & int32(255)
	v198 = v193
	goto L54
L58:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v175 == int32(0) {
		v193 = v172
		v194 = v173
		goto L57
	} else {
		goto L60
	}
L59:
	;
	v193 = v187
	v194 = int32(0)
	goto L57
L60:
	;
	v179 = v173 & int32(255)
	if v179 == v175 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v186 = int32(1)
	v187 = v172 + v186
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+1)))
	if v188 != 0 {
		v171 = v171 + v186
		v172 = v187
		v173 = v188
		goto L58
	} else {
		goto L64
	}
L62:
	;
	v181 = F_tolower(m, v179)
	mBase = m.M
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	v183 = F_tolower(m, v182)
	mBase = m.M
	if v181 == v183 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	v193 = v172
	v194 = v185
	goto L57
L64:
	;
	goto L59
L65:
	;
	v341 = v22 | int32(8)
	v342 = v23
	goto L9
L66:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v251+v27)))
	v254 = F_objectGetVal(m, v253)
	mBase = m.M
	v255 = int32(_a495)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if v258 != 0 {
		goto L83
	} else {
		goto L84
	}
L67:
	;
	if v246-v248 != 0 {
		goto L66
	} else {
		goto L79
	}
L68:
	;
	v246 = F_tolower(m, v242)
	mBase = m.M
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v248 = F_tolower(m, v247)
	mBase = m.M
	goto L67
L69:
	;
	v216 = v210
	v217 = v211
	v218 = v214
	goto L72
L70:
	;
	v242 = int32(0)
	v243 = v211
	goto L68
L71:
	;
	v242 = v239 & int32(255)
	v243 = v238
	goto L68
L72:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v220 == int32(0) {
		v238 = v217
		v239 = v218
		goto L71
	} else {
		goto L74
	}
L73:
	;
	v238 = v232
	v239 = int32(0)
	goto L71
L74:
	;
	v224 = v218 & int32(255)
	if v224 == v220 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v231 = int32(1)
	v232 = v217 + v231
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	if v233 != 0 {
		v216 = v216 + v231
		v217 = v232
		v218 = v233
		goto L72
	} else {
		goto L78
	}
L76:
	;
	v226 = F_tolower(m, v224)
	mBase = m.M
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	v228 = F_tolower(m, v227)
	mBase = m.M
	if v226 == v228 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v238 = v217
	v239 = v230
	goto L71
L78:
	;
	goto L73
L79:
	;
	v341 = v22
	v342 = int32(1)
	goto L9
L80:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v296+v27)))
	v299 = F_objectGetVal(m, v298)
	mBase = m.M
	v300 = int32(_a496)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	if v303 != 0 {
		goto L96
	} else {
		goto L97
	}
L81:
	;
	if v290-v292 != 0 {
		goto L80
	} else {
		goto L93
	}
L82:
	;
	v290 = F_tolower(m, v286)
	mBase = m.M
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	v292 = F_tolower(m, v291)
	mBase = m.M
	goto L81
L83:
	;
	v260 = v254
	v261 = v255
	v262 = v258
	goto L86
L84:
	;
	v286 = int32(0)
	v287 = v255
	goto L82
L85:
	;
	v286 = v283 & int32(255)
	v287 = v282
	goto L82
L86:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	if v264 == int32(0) {
		v282 = v261
		v283 = v262
		goto L85
	} else {
		goto L88
	}
L87:
	;
	v282 = v276
	v283 = int32(0)
	goto L85
L88:
	;
	v268 = v262 & int32(255)
	if v268 == v264 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v275 = int32(1)
	v276 = v261 + v275
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	if v277 != 0 {
		v260 = v260 + v275
		v261 = v276
		v262 = v277
		goto L86
	} else {
		goto L92
	}
L90:
	;
	v270 = F_tolower(m, v268)
	mBase = m.M
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	v272 = F_tolower(m, v271)
	mBase = m.M
	if v270 == v272 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	v282 = v261
	v283 = v274
	goto L85
L92:
	;
	goto L87
L93:
	;
	v341 = v22 | int32(16)
	v342 = v23
	goto L9
L94:
	;
	if v335-v337 != 0 {
		goto L2
	} else {
		goto L106
	}
L95:
	;
	v335 = F_tolower(m, v331)
	mBase = m.M
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	v337 = F_tolower(m, v336)
	mBase = m.M
	goto L94
L96:
	;
	v305 = v299
	v306 = v300
	v307 = v303
	goto L99
L97:
	;
	v331 = int32(0)
	v332 = v300
	goto L95
L98:
	;
	v331 = v328 & int32(255)
	v332 = v327
	goto L95
L99:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	if v309 == int32(0) {
		v327 = v306
		v328 = v307
		goto L98
	} else {
		goto L101
	}
L100:
	;
	v327 = v321
	v328 = int32(0)
	goto L98
L101:
	;
	v313 = v307 & int32(255)
	if v313 == v309 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v320 = int32(1)
	v321 = v306 + v320
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+1)))
	if v322 != 0 {
		v305 = v305 + v320
		v306 = v321
		v307 = v322
		goto L99
	} else {
		goto L105
	}
L103:
	;
	v315 = F_tolower(m, v313)
	mBase = m.M
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	v317 = F_tolower(m, v316)
	mBase = m.M
	if v315 == v317 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	v327 = v306
	v328 = v319
	goto L98
L105:
	;
	goto L100
L106:
	;
	v341 = v22 | int32(32)
	v342 = v23
	goto L9
L107:
	;
	goto L8
L108:
	;
	if v341&int32(32) == int32(0) {
		goto L116
	} else {
		goto L117
	}
L109:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v354)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	v349 = int32(3)
	if v341&v349 != v349 {
		goto L108
	} else {
		goto L113
	}
L111:
	;
	if v341 != 0 {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	goto L109
L114:
	;
	return
L115:
	;
	goto L1
L116:
	;
	if v342 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v362 != 0 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	F_addReplyError(m, l0, int32(_a497))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L114
	} else {
		goto L119
	}
L119:
	;
	goto L1
L120:
	;
	v380 = base.B2i32(v341&int32(2) == int32(0))
	if v341&int32(4) != 0 {
		v399 = v380
		v400 = v341
		goto L3
	} else {
		goto L127
	}
L121:
	;
	v368 = F_abortShutdown(m)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L114
	} else {
		goto L123
	}
L122:
	;
	F_addReplyError(m, l0, int32(_a498))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L114
	} else {
		goto L126
	}
L123:
	;
	if v368 != 0 {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L114
	} else {
		goto L125
	}
L125:
	;
	goto L1
L126:
	;
	goto L1
L127:
	;
	v385 = v380
	v386 = v341
	goto L4
L128:
	;
	F_addReplyError(m, l0, int32(_a499))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L114
	} else {
		goto L129
	}
L129:
	;
	goto L1
L130:
	;
	F_blockClientShutdown(m, l0)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L114
	} else {
		goto L146
	}
L131:
	;
	v405 = F_scriptIsTimedout(m)
	mBase = m.M
	v406 = int32(0)
	v407 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	goto L132
L132:
	;
	if base.B2i32(v405|v407 != v406) == int32(0) {
		goto L130
	} else {
		goto L133
	}
L133:
	;
	v414 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	if v414 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	if v414 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v418 = *(*int32)(unsafe.Add(mBase, _consts[227]))
	if v418 == int32(0) {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v418
	F_addReplyErrorFormat(m, l0, int32(_a500), v9)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L114
	} else {
		goto L137
	}
L137:
	;
	goto L1
L138:
	;
	v432 = F_scriptIsEval(m)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L114
	} else {
		goto L142
	}
L139:
	;
	v429 = *(*int32)(unsafe.Add(mBase, _consts[228]))
	F_addReplyErrorObject(m, l0, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L114
	} else {
		goto L140
	}
L140:
	;
	goto L1
L141:
	;
	v441 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	F_addReplyErrorObject(m, l0, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L114
	} else {
		goto L145
	}
L142:
	;
	if v432 == int32(0) {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v437 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	F_addReplyErrorObject(m, l0, v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L114
	} else {
		goto L144
	}
L144:
	;
	goto L1
L145:
	;
	goto L1
L146:
	;
	v446 = F_prepareForShutdown(m, l0, v400)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L114
	} else {
		goto L147
	}
L147:
	;
	if v446 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	m.Env.Exit(m, int32(0))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	goto L1
}
func F_sift(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	v14 = m.G0
	v16 = v14 - int32(240)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	if l4 < int32(2) {
		v77 = int32(1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v88 = m.G0
	v90 = v88 - int32(256)
	m.G0 = v90
	if v77 < int32(2) {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v29 = l4
	v32 = int32(1)
	v34 = l0
	goto L3
L3:
	;
	v38 = v34 + (int32(0) - l1)
	v40 = v29 + int32(-2)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l5+v40<<(uint(int32(2))%32))))
	v45 = v38 - v44
	v46 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, l0, v45, l3)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v77 = v64
	goto L1
L5:
	;
	v57 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v45, v38, l3)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L11
	}
L6:
	;
	return
L7:
	;
	if v46 < int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v50 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, l0, v38, l3)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if int32(-1) < v50 {
		v77 = v32
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L5
L11:
	;
	v60 = base.B2i32(int32(-1) < v57)
	if int32(-1) < v57 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v61 = v45
	goto L14
L13:
	;
	v61 = v38
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+v32<<(uint(int32(2))%32)))) = v61
	v64 = v32 + int32(1)
	if int32(-1) < v57 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v67 = v29 + int32(-1)
	goto L17
L16:
	;
	v67 = v40
	goto L17
L17:
	;
	if int32(1) < v67 {
		v29 = v67
		v32 = v64
		v34 = v61
		goto L3
	} else {
		goto L18
	}
L18:
	;
	goto L4
L19:
	;
	m.G0 = v16 + int32(240)
	return
L20:
	;
	m.G0 = v90 + int32(256)
	goto L19
L21:
	;
	v96 = v16 + v77<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v90
	if l1 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v100 = l1
	goto L23
L23:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v110 = int32(256)
	if base.Ui32(v100) < base.Ui32(v110) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L20
L25:
	;
	v113 = v100
	goto L27
L26:
	;
	v113 = v110
	goto L27
L27:
	;
	v114 = F___memcpy(m, v108, v109, v113)
	mBase = m.M
	v122 = int32(0)
	goto L28
L28:
	;
	v124 = int32(2)
	v126 = v16 + v122<<(uint(v124)%32)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v129 = v122 + int32(1)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v16+v129<<(uint(v124)%32))))
	v134 = F___memcpy(m, v127, v133, v113)
	mBase = m.M
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v135 + v113
	if v129 != v77 {
		v122 = v129
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v139 = v100 - v113
	if v139 != 0 {
		v100 = v139
		goto L23
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	goto L24
}
func F_sigalrmSignalHandler(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int64
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int64
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v7 != 0 {
		v8 = int32(_a528)
	} else {
		v8 = int32(_a529)
	}
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v19 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v19 {
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _consts[239]))
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
		if v23&int32(255) != 0 {
			if v23&int32(255) != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(420)
				v34 = F_open(m, v22, int32(1089), v14)
				mBase = m.M
				if v34 == int32(-1) {
				} else {
					v37 = v34
					v45 = v14 + int32(16)
					v47 = F_getpid(m)
					mBase = m.M
					v49 = F_ll2string(m, v45, int32(64), base.I64_extend_i32_s(v47))
					mBase = m.M
					v54 = F_strlen(m, v45)
					mBase = m.M
					v55 = F_write(m, v37, v45, v54)
					mBase = m.M
					if v55 == int32(-1) {
					} else {
						v60 = F_write(m, v37, int32(_a530), int32(17))
						mBase = m.M
						if v60 == int32(-1) {
						} else {
							v64 = v14 + int32(16)
							v67 = F___time(m, int32(0))
							mBase = m.M
							v68 = F_ll2string(m, v64, int32(64), v67)
							mBase = m.M
							v73 = F_strlen(m, v64)
							mBase = m.M
							v74 = F_write(m, v37, v64, v73)
							mBase = m.M
							if v74 == int32(-1) {
							} else {
								v79 = F_write(m, v37, int32(_a531), int32(2))
								mBase = m.M
								if v79 == int32(-1) {
								} else {
									v82 = F_strlen(m, v8)
									mBase = m.M
									v83 = F_write(m, v37, v8, v82)
									mBase = m.M
									if v83 == int32(-1) {
									} else {
										v88 = F_write(m, v37, int32(_a397), int32(1))
										mBase = m.M
									}
								}
							}
						}
					}
					if v23&int32(255) == int32(0) {
					} else {
						v93 = F_close(m, v37)
						mBase = m.M
					}
				}
			} else {
				v37 = int32(1)
				v45 = v14 + int32(16)
				v47 = F_getpid(m)
				mBase = m.M
				v49 = F_ll2string(m, v45, int32(64), base.I64_extend_i32_s(v47))
				mBase = m.M
				v54 = F_strlen(m, v45)
				mBase = m.M
				v55 = F_write(m, v37, v45, v54)
				mBase = m.M
				if v55 == int32(-1) {
				} else {
					v60 = F_write(m, v37, int32(_a530), int32(17))
					mBase = m.M
					if v60 == int32(-1) {
					} else {
						v64 = v14 + int32(16)
						v67 = F___time(m, int32(0))
						mBase = m.M
						v68 = F_ll2string(m, v64, int32(64), v67)
						mBase = m.M
						v73 = F_strlen(m, v64)
						mBase = m.M
						v74 = F_write(m, v37, v64, v73)
						mBase = m.M
						if v74 == int32(-1) {
						} else {
							v79 = F_write(m, v37, int32(_a531), int32(2))
							mBase = m.M
							if v79 == int32(-1) {
							} else {
								v82 = F_strlen(m, v8)
								mBase = m.M
								v83 = F_write(m, v37, v8, v82)
								mBase = m.M
								if v83 == int32(-1) {
								} else {
									v88 = F_write(m, v37, int32(_a397), int32(1))
									mBase = m.M
								}
							}
						}
					}
				}
				if v23&int32(255) == int32(0) {
				} else {
					v93 = F_close(m, v37)
					mBase = m.M
				}
			}
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, _consts[240]))
			if v27 != 0 {
			} else {
				if v23&int32(255) != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(420)
					v34 = F_open(m, v22, int32(1089), v14)
					mBase = m.M
					if v34 == int32(-1) {
					} else {
						v37 = v34
						v45 = v14 + int32(16)
						v47 = F_getpid(m)
						mBase = m.M
						v49 = F_ll2string(m, v45, int32(64), base.I64_extend_i32_s(v47))
						mBase = m.M
						v54 = F_strlen(m, v45)
						mBase = m.M
						v55 = F_write(m, v37, v45, v54)
						mBase = m.M
						if v55 == int32(-1) {
						} else {
							v60 = F_write(m, v37, int32(_a530), int32(17))
							mBase = m.M
							if v60 == int32(-1) {
							} else {
								v64 = v14 + int32(16)
								v67 = F___time(m, int32(0))
								mBase = m.M
								v68 = F_ll2string(m, v64, int32(64), v67)
								mBase = m.M
								v73 = F_strlen(m, v64)
								mBase = m.M
								v74 = F_write(m, v37, v64, v73)
								mBase = m.M
								if v74 == int32(-1) {
								} else {
									v79 = F_write(m, v37, int32(_a531), int32(2))
									mBase = m.M
									if v79 == int32(-1) {
									} else {
										v82 = F_strlen(m, v8)
										mBase = m.M
										v83 = F_write(m, v37, v8, v82)
										mBase = m.M
										if v83 == int32(-1) {
										} else {
											v88 = F_write(m, v37, int32(_a397), int32(1))
											mBase = m.M
										}
									}
								}
							}
						}
						if v23&int32(255) == int32(0) {
						} else {
							v93 = F_close(m, v37)
							mBase = m.M
						}
					}
				} else {
					v37 = int32(1)
					v45 = v14 + int32(16)
					v47 = F_getpid(m)
					mBase = m.M
					v49 = F_ll2string(m, v45, int32(64), base.I64_extend_i32_s(v47))
					mBase = m.M
					v54 = F_strlen(m, v45)
					mBase = m.M
					v55 = F_write(m, v37, v45, v54)
					mBase = m.M
					if v55 == int32(-1) {
					} else {
						v60 = F_write(m, v37, int32(_a530), int32(17))
						mBase = m.M
						if v60 == int32(-1) {
						} else {
							v64 = v14 + int32(16)
							v67 = F___time(m, int32(0))
							mBase = m.M
							v68 = F_ll2string(m, v64, int32(64), v67)
							mBase = m.M
							v73 = F_strlen(m, v64)
							mBase = m.M
							v74 = F_write(m, v37, v64, v73)
							mBase = m.M
							if v74 == int32(-1) {
							} else {
								v79 = F_write(m, v37, int32(_a531), int32(2))
								mBase = m.M
								if v79 == int32(-1) {
								} else {
									v82 = F_strlen(m, v8)
									mBase = m.M
									v83 = F_write(m, v37, v8, v82)
									mBase = m.M
									if v83 == int32(-1) {
									} else {
										v88 = F_write(m, v37, int32(_a397), int32(1))
										mBase = m.M
									}
								}
							}
						}
					}
					if v23&int32(255) == int32(0) {
					} else {
						v93 = F_close(m, v37)
						mBase = m.M
					}
				}
			}
		}
	}
	m.G0 = v14 + int32(80)
	v100 = int32(_a532)
	v104 = m.G0
	v106 = v104 - int32(80)
	m.G0 = v106
	v111 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v111 {
	} else {
		v114 = *(*int32)(unsafe.Add(mBase, _consts[239]))
		v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
		if v115&int32(255) != 0 {
			if v115&int32(255) != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(420)
				v126 = F_open(m, v114, int32(1089), v106)
				mBase = m.M
				if v126 == int32(-1) {
				} else {
					v129 = v126
					v137 = v106 + int32(16)
					v139 = F_getpid(m)
					mBase = m.M
					v141 = F_ll2string(m, v137, int32(64), base.I64_extend_i32_s(v139))
					mBase = m.M
					v146 = F_strlen(m, v137)
					mBase = m.M
					v147 = F_write(m, v129, v137, v146)
					mBase = m.M
					if v147 == int32(-1) {
					} else {
						v152 = F_write(m, v129, int32(_a530), int32(17))
						mBase = m.M
						if v152 == int32(-1) {
						} else {
							v156 = v106 + int32(16)
							v159 = F___time(m, int32(0))
							mBase = m.M
							v160 = F_ll2string(m, v156, int32(64), v159)
							mBase = m.M
							v165 = F_strlen(m, v156)
							mBase = m.M
							v166 = F_write(m, v129, v156, v165)
							mBase = m.M
							if v166 == int32(-1) {
							} else {
								v171 = F_write(m, v129, int32(_a531), int32(2))
								mBase = m.M
								if v171 == int32(-1) {
								} else {
									v174 = F_strlen(m, v100)
									mBase = m.M
									v175 = F_write(m, v129, v100, v174)
									mBase = m.M
									if v175 == int32(-1) {
									} else {
										v180 = F_write(m, v129, int32(_a397), int32(1))
										mBase = m.M
									}
								}
							}
						}
					}
					if v115&int32(255) == int32(0) {
					} else {
						v185 = F_close(m, v129)
						mBase = m.M
					}
				}
			} else {
				v129 = int32(1)
				v137 = v106 + int32(16)
				v139 = F_getpid(m)
				mBase = m.M
				v141 = F_ll2string(m, v137, int32(64), base.I64_extend_i32_s(v139))
				mBase = m.M
				v146 = F_strlen(m, v137)
				mBase = m.M
				v147 = F_write(m, v129, v137, v146)
				mBase = m.M
				if v147 == int32(-1) {
				} else {
					v152 = F_write(m, v129, int32(_a530), int32(17))
					mBase = m.M
					if v152 == int32(-1) {
					} else {
						v156 = v106 + int32(16)
						v159 = F___time(m, int32(0))
						mBase = m.M
						v160 = F_ll2string(m, v156, int32(64), v159)
						mBase = m.M
						v165 = F_strlen(m, v156)
						mBase = m.M
						v166 = F_write(m, v129, v156, v165)
						mBase = m.M
						if v166 == int32(-1) {
						} else {
							v171 = F_write(m, v129, int32(_a531), int32(2))
							mBase = m.M
							if v171 == int32(-1) {
							} else {
								v174 = F_strlen(m, v100)
								mBase = m.M
								v175 = F_write(m, v129, v100, v174)
								mBase = m.M
								if v175 == int32(-1) {
								} else {
									v180 = F_write(m, v129, int32(_a397), int32(1))
									mBase = m.M
								}
							}
						}
					}
				}
				if v115&int32(255) == int32(0) {
				} else {
					v185 = F_close(m, v129)
					mBase = m.M
				}
			}
		} else {
			v119 = *(*int32)(unsafe.Add(mBase, _consts[240]))
			if v119 != 0 {
			} else {
				if v115&int32(255) != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(420)
					v126 = F_open(m, v114, int32(1089), v106)
					mBase = m.M
					if v126 == int32(-1) {
					} else {
						v129 = v126
						v137 = v106 + int32(16)
						v139 = F_getpid(m)
						mBase = m.M
						v141 = F_ll2string(m, v137, int32(64), base.I64_extend_i32_s(v139))
						mBase = m.M
						v146 = F_strlen(m, v137)
						mBase = m.M
						v147 = F_write(m, v129, v137, v146)
						mBase = m.M
						if v147 == int32(-1) {
						} else {
							v152 = F_write(m, v129, int32(_a530), int32(17))
							mBase = m.M
							if v152 == int32(-1) {
							} else {
								v156 = v106 + int32(16)
								v159 = F___time(m, int32(0))
								mBase = m.M
								v160 = F_ll2string(m, v156, int32(64), v159)
								mBase = m.M
								v165 = F_strlen(m, v156)
								mBase = m.M
								v166 = F_write(m, v129, v156, v165)
								mBase = m.M
								if v166 == int32(-1) {
								} else {
									v171 = F_write(m, v129, int32(_a531), int32(2))
									mBase = m.M
									if v171 == int32(-1) {
									} else {
										v174 = F_strlen(m, v100)
										mBase = m.M
										v175 = F_write(m, v129, v100, v174)
										mBase = m.M
										if v175 == int32(-1) {
										} else {
											v180 = F_write(m, v129, int32(_a397), int32(1))
											mBase = m.M
										}
									}
								}
							}
						}
						if v115&int32(255) == int32(0) {
						} else {
							v185 = F_close(m, v129)
							mBase = m.M
						}
					}
				} else {
					v129 = int32(1)
					v137 = v106 + int32(16)
					v139 = F_getpid(m)
					mBase = m.M
					v141 = F_ll2string(m, v137, int32(64), base.I64_extend_i32_s(v139))
					mBase = m.M
					v146 = F_strlen(m, v137)
					mBase = m.M
					v147 = F_write(m, v129, v137, v146)
					mBase = m.M
					if v147 == int32(-1) {
					} else {
						v152 = F_write(m, v129, int32(_a530), int32(17))
						mBase = m.M
						if v152 == int32(-1) {
						} else {
							v156 = v106 + int32(16)
							v159 = F___time(m, int32(0))
							mBase = m.M
							v160 = F_ll2string(m, v156, int32(64), v159)
							mBase = m.M
							v165 = F_strlen(m, v156)
							mBase = m.M
							v166 = F_write(m, v129, v156, v165)
							mBase = m.M
							if v166 == int32(-1) {
							} else {
								v171 = F_write(m, v129, int32(_a531), int32(2))
								mBase = m.M
								if v171 == int32(-1) {
								} else {
									v174 = F_strlen(m, v100)
									mBase = m.M
									v175 = F_write(m, v129, v100, v174)
									mBase = m.M
									if v175 == int32(-1) {
									} else {
										v180 = F_write(m, v129, int32(_a397), int32(1))
										mBase = m.M
									}
								}
							}
						}
					}
					if v115&int32(255) == int32(0) {
					} else {
						v185 = F_close(m, v129)
						mBase = m.M
					}
				}
			}
		}
	}
	m.G0 = v106 + int32(80)
	v192 = int32(_a533)
	v196 = m.G0
	v198 = v196 - int32(80)
	m.G0 = v198
	v203 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v203 {
	} else {
		v206 = *(*int32)(unsafe.Add(mBase, _consts[239]))
		v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
		if v207&int32(255) != 0 {
			if v207&int32(255) != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v198))) = int32(420)
				v218 = F_open(m, v206, int32(1089), v198)
				mBase = m.M
				if v218 == int32(-1) {
				} else {
					v221 = v218
					v229 = v198 + int32(16)
					v231 = F_getpid(m)
					mBase = m.M
					v233 = F_ll2string(m, v229, int32(64), base.I64_extend_i32_s(v231))
					mBase = m.M
					v238 = F_strlen(m, v229)
					mBase = m.M
					v239 = F_write(m, v221, v229, v238)
					mBase = m.M
					if v239 == int32(-1) {
					} else {
						v244 = F_write(m, v221, int32(_a530), int32(17))
						mBase = m.M
						if v244 == int32(-1) {
						} else {
							v248 = v198 + int32(16)
							v251 = F___time(m, int32(0))
							mBase = m.M
							v252 = F_ll2string(m, v248, int32(64), v251)
							mBase = m.M
							v257 = F_strlen(m, v248)
							mBase = m.M
							v258 = F_write(m, v221, v248, v257)
							mBase = m.M
							if v258 == int32(-1) {
							} else {
								v263 = F_write(m, v221, int32(_a531), int32(2))
								mBase = m.M
								if v263 == int32(-1) {
								} else {
									v266 = F_strlen(m, v192)
									mBase = m.M
									v267 = F_write(m, v221, v192, v266)
									mBase = m.M
									if v267 == int32(-1) {
									} else {
										v272 = F_write(m, v221, int32(_a397), int32(1))
										mBase = m.M
									}
								}
							}
						}
					}
					if v207&int32(255) == int32(0) {
					} else {
						v277 = F_close(m, v221)
						mBase = m.M
					}
				}
			} else {
				v221 = int32(1)
				v229 = v198 + int32(16)
				v231 = F_getpid(m)
				mBase = m.M
				v233 = F_ll2string(m, v229, int32(64), base.I64_extend_i32_s(v231))
				mBase = m.M
				v238 = F_strlen(m, v229)
				mBase = m.M
				v239 = F_write(m, v221, v229, v238)
				mBase = m.M
				if v239 == int32(-1) {
				} else {
					v244 = F_write(m, v221, int32(_a530), int32(17))
					mBase = m.M
					if v244 == int32(-1) {
					} else {
						v248 = v198 + int32(16)
						v251 = F___time(m, int32(0))
						mBase = m.M
						v252 = F_ll2string(m, v248, int32(64), v251)
						mBase = m.M
						v257 = F_strlen(m, v248)
						mBase = m.M
						v258 = F_write(m, v221, v248, v257)
						mBase = m.M
						if v258 == int32(-1) {
						} else {
							v263 = F_write(m, v221, int32(_a531), int32(2))
							mBase = m.M
							if v263 == int32(-1) {
							} else {
								v266 = F_strlen(m, v192)
								mBase = m.M
								v267 = F_write(m, v221, v192, v266)
								mBase = m.M
								if v267 == int32(-1) {
								} else {
									v272 = F_write(m, v221, int32(_a397), int32(1))
									mBase = m.M
								}
							}
						}
					}
				}
				if v207&int32(255) == int32(0) {
				} else {
					v277 = F_close(m, v221)
					mBase = m.M
				}
			}
		} else {
			v211 = *(*int32)(unsafe.Add(mBase, _consts[240]))
			if v211 != 0 {
			} else {
				if v207&int32(255) != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v198))) = int32(420)
					v218 = F_open(m, v206, int32(1089), v198)
					mBase = m.M
					if v218 == int32(-1) {
					} else {
						v221 = v218
						v229 = v198 + int32(16)
						v231 = F_getpid(m)
						mBase = m.M
						v233 = F_ll2string(m, v229, int32(64), base.I64_extend_i32_s(v231))
						mBase = m.M
						v238 = F_strlen(m, v229)
						mBase = m.M
						v239 = F_write(m, v221, v229, v238)
						mBase = m.M
						if v239 == int32(-1) {
						} else {
							v244 = F_write(m, v221, int32(_a530), int32(17))
							mBase = m.M
							if v244 == int32(-1) {
							} else {
								v248 = v198 + int32(16)
								v251 = F___time(m, int32(0))
								mBase = m.M
								v252 = F_ll2string(m, v248, int32(64), v251)
								mBase = m.M
								v257 = F_strlen(m, v248)
								mBase = m.M
								v258 = F_write(m, v221, v248, v257)
								mBase = m.M
								if v258 == int32(-1) {
								} else {
									v263 = F_write(m, v221, int32(_a531), int32(2))
									mBase = m.M
									if v263 == int32(-1) {
									} else {
										v266 = F_strlen(m, v192)
										mBase = m.M
										v267 = F_write(m, v221, v192, v266)
										mBase = m.M
										if v267 == int32(-1) {
										} else {
											v272 = F_write(m, v221, int32(_a397), int32(1))
											mBase = m.M
										}
									}
								}
							}
						}
						if v207&int32(255) == int32(0) {
						} else {
							v277 = F_close(m, v221)
							mBase = m.M
						}
					}
				} else {
					v221 = int32(1)
					v229 = v198 + int32(16)
					v231 = F_getpid(m)
					mBase = m.M
					v233 = F_ll2string(m, v229, int32(64), base.I64_extend_i32_s(v231))
					mBase = m.M
					v238 = F_strlen(m, v229)
					mBase = m.M
					v239 = F_write(m, v221, v229, v238)
					mBase = m.M
					if v239 == int32(-1) {
					} else {
						v244 = F_write(m, v221, int32(_a530), int32(17))
						mBase = m.M
						if v244 == int32(-1) {
						} else {
							v248 = v198 + int32(16)
							v251 = F___time(m, int32(0))
							mBase = m.M
							v252 = F_ll2string(m, v248, int32(64), v251)
							mBase = m.M
							v257 = F_strlen(m, v248)
							mBase = m.M
							v258 = F_write(m, v221, v248, v257)
							mBase = m.M
							if v258 == int32(-1) {
							} else {
								v263 = F_write(m, v221, int32(_a531), int32(2))
								mBase = m.M
								if v263 == int32(-1) {
								} else {
									v266 = F_strlen(m, v192)
									mBase = m.M
									v267 = F_write(m, v221, v192, v266)
									mBase = m.M
									if v267 == int32(-1) {
									} else {
										v272 = F_write(m, v221, int32(_a397), int32(1))
										mBase = m.M
									}
								}
							}
						}
					}
					if v207&int32(255) == int32(0) {
					} else {
						v277 = F_close(m, v221)
						mBase = m.M
					}
				}
			}
		}
	}
	m.G0 = v198 + int32(80)
	return
}
func F_signalModifiedKey(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	F_touchWatchedKey(m, l1, l2)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_trackingInvalidateKey(m, l0, l2, int32(1))
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_sigsegvHandler(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	v7 = m.G0
	v9 = v7 - int32(64)
	m.G0 = v9
	goto L2
L1:
	;
	goto L23
L2:
	;
	goto L1
L23:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, _consts[238])))
	if v110 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	goto L33
L25:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v112 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v128 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[238])) = uint8(v128)
	goto L24
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _consts[237]))
	if v118 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v119 = int32(_a518)
	goto L30
L29:
	;
	v119 = int32(_a519)
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v119
	F__serverLog(m, int32(1027), int32(_a520), v9+int32(48))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return
L32:
	;
	goto L26
L33:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v133 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v164 != 0 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(_a535)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l0
	F__serverLog(m, int32(3), int32(_a536), v9+int32(32))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	if base.B2i32(l0 != int32(11))&base.B2i32(l0 != int32(7)) != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v153 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v156
	F__serverLog(m, int32(3), int32(_a537), v9+int32(16))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L31
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	goto L46
L41:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v165 == int32(-1) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v169 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v165
	F__serverLog(m, int32(3), int32(_a534), v9)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L31
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	F_bugReportEnd(m, int32(1), l0)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L31
	} else {
		goto L48
	}
L46:
	;
	F_printCrashReport(m)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L31
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	m.G0 = v9 + int32(64)
	return
}
func F_sinterCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = int32(0)
	F_sinterGenericCommand(m, l0, v2+int32(4), v5+int32(-1), v8, v8, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		return
	}
}
func F_sortROGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v8 != 0 {
		v13 = v8
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
		if int32(0) < v14 {
			v43 = v13
			*(*int64)(unsafe.Add(mBase, uint32(v43))) = int64(73014444033)
			v48 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
			return v48
		} else {
			v18 = l3 + int32(12)
			if v13 == v18 {
				v27 = F_valkey_malloc(m, int32(8))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v27
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					if v30 == int32(0) {
						v39 = v27
					} else {
						v34 = v30 << (uint(int32(3)) % 32)
						if v34 == int32(0) {
						} else {
							v37 = F__emscripten_memcpy_bulkmem(m, v27, v18, v34)
							mBase = m.M
						}
						v39 = v27
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
					v43 = v39
					*(*int64)(unsafe.Add(mBase, uint32(v43))) = int64(73014444033)
					v48 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
					return v48
				}
			} else {
				v21 = F_valkey_realloc(m, v13, int32(8))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v21
					v39 = v21
					*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
					v43 = v39
					*(*int64)(unsafe.Add(mBase, uint32(v43))) = int64(73014444033)
					v48 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
					return v48
				}
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		if v9 != 0 {
			F__serverAssert(m, int32(_a512), int32(_a474), int32(2292))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v11 = l3 + int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v11
			v13 = v11
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			if int32(0) < v14 {
				v43 = v13
				*(*int64)(unsafe.Add(mBase, uint32(v43))) = int64(73014444033)
				v48 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
				return v48
			} else {
				v18 = l3 + int32(12)
				if v13 == v18 {
					v27 = F_valkey_malloc(m, int32(8))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v27
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						if v30 == int32(0) {
							v39 = v27
						} else {
							v34 = v30 << (uint(int32(3)) % 32)
							if v34 == int32(0) {
							} else {
								v37 = F__emscripten_memcpy_bulkmem(m, v27, v18, v34)
								mBase = m.M
							}
							v39 = v27
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
						v43 = v39
						*(*int64)(unsafe.Add(mBase, uint32(v43))) = int64(73014444033)
						v48 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
						return v48
					}
				} else {
					v21 = F_valkey_realloc(m, v13, int32(8))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v21
						v39 = v21
						*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
						v43 = v39
						*(*int64)(unsafe.Add(mBase, uint32(v43))) = int64(73014444033)
						v48 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
						return v48
					}
				}
			}
		}
	}
}
func F_sort_comp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v348 int32
	_ = v348
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = v7 + int32(16)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v12) < base.Ui32(v13) {
		v55 = m.G398
		if v12 != v55 {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			v61 = v58
		} else {
			v61 = int32(-1)
		}
	} else {
		v61 = int32(-1)
	}
	if v61 == int32(0) {
		v359 = F_lua_lessthan(m, l0, l1, l2)
		mBase = m.M
		v360 = m.ExcPending
		if v360 != 0 {
			return int32(0)
		} else {
			return v359
		}
	} else {
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v75 = v70 + int32(16)
		v76 = m.G398
		if base.Ui32(v75) < base.Ui32(v69) {
			v78 = v75
		} else {
			v78 = v76
		}
		v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v121 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
		*(*int64)(unsafe.Add(mBase, uint32(v120))) = v121
		v123 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = v123
		v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v125 + int32(16)
		v130 = l1 + int32(-1)
		if v130 < int32(1) {
			if v130 < int32(-9999) {
				switch l1 + int32(10001) {
				case 0:
					v183 = l0 + int32(72)
				case 1:
					v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
					v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
					v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v159
					v183 = l0 + int32(88)
				case 2:
					v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v183 = v153 + int32(96)
				default:
					v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
					v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
					v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+7)))
					v171 = m.G398
					if base.Ui32(v170) < base.Ui32(int32(-10002)-v130) {
						v182 = v171
					} else {
						v182 = v169 + (int32(-10003)-v130)<<(uint(int32(4))%32) + int32(24)
					}
					v183 = v182
				}
			} else {
				v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v183 = v147 + v130<<(uint(int32(4))%32)
			}
		} else {
			v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v141 = v136 + v130<<(uint(int32(4))%32) + int32(-16)
			v142 = m.G398
			if base.Ui32(v141) < base.Ui32(v135) {
				v144 = v141
			} else {
				v144 = v142
			}
			v183 = v144
		}
		v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v187 = *(*int64)(unsafe.Add(mBase, uint32(v183)))
		*(*int64)(unsafe.Add(mBase, uint32(v186))) = v187
		v189 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v186)+8)) = v189
		v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v191 + int32(16)
		v196 = l2 + int32(-2)
		if v196 < int32(1) {
			if v196 < int32(-9999) {
				switch l2 + int32(10000) {
				case 0:
					v249 = l0 + int32(72)
				case 1:
					v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
					v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
					v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v225
					v249 = l0 + int32(88)
				case 2:
					v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v249 = v219 + int32(96)
				default:
					v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
					v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
					v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+7)))
					v237 = m.G398
					if base.Ui32(v236) < base.Ui32(int32(-10002)-v196) {
						v248 = v237
					} else {
						v248 = v235 + (int32(-10003)-v196)<<(uint(int32(4))%32) + int32(24)
					}
					v249 = v248
				}
			} else {
				v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v249 = v213 + v196<<(uint(int32(4))%32)
			}
		} else {
			v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v207 = v202 + v196<<(uint(int32(4))%32) + int32(-16)
			v208 = m.G398
			if base.Ui32(v207) < base.Ui32(v201) {
				v210 = v207
			} else {
				v210 = v208
			}
			v249 = v210
		}
		v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v253 = *(*int64)(unsafe.Add(mBase, uint32(v249)))
		*(*int64)(unsafe.Add(mBase, uint32(v252))) = v253
		v255 = *(*int32)(unsafe.Add(mBase, uint32(v249)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v252)+8)) = v255
		v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v257 + int32(16)
		F_lua_call(m, l0, int32(2), int32(1))
		mBase = m.M
		v266 = m.ExcPending
		if v266 != 0 {
			return int32(0)
		} else {
			v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v286 = v283 + int32(-16)
			v322 = *(*int32)(unsafe.Add(mBase, uint32(v286)+8))
			switch v322 {
			case 0:
				v327 = v322
				v329 = v327
			case 1:
				v323 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
				v329 = base.B2i32(v323 != int32(0))
			default:
				v327 = int32(1)
				v329 = v327
			}
			v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v348 + int32(-16)
			return v329
		}
	}
}
func F_sort_gp_asc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_ne(v8, v9) != 0 {
		v11 = int32(-1)
	} else {
		v11 = int32(0)
	}
	if base.F64_gt(v8, v9) != 0 {
		v13 = int32(1)
	} else {
		v13 = v11
	}
	return v13
}
func F_sparklineRender(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v9 < int32(1) {
		v38 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v38
L2:
	;
	v13 = l0
	v18 = v9
	v19 = int32(0)
	goto L3
L3:
	;
	v21 = v18 - v19
	if v21 < l2 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v38 = v33
	goto L1
L5:
	;
	v23 = v21
	goto L7
L6:
	;
	v23 = l2
	goto L7
L7:
	;
	if v19 == int32(0) {
		v32 = v13
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v33 = F_sparklineRenderRange(m, v32, l1, l3, v19, v23, l4)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L12
	}
L9:
	;
	v28 = F_sdscatlen(m, v13, int32(_a397), int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v32 = v28
	goto L8
L12:
	;
	v35 = v19 + l2
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v35 < v36 {
		v13 = v33
		v18 = v36
		v19 = v35
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L4
}
func F_specialcase_1(m *base.Module, l0 float64, l1 int64, l2 int64) float64 {
	mBase := m.M
	_ = mBase
	var v13 float64
	_ = v13
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 float64
	_ = v34
	var v37 int32
	_ = v37
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v52 float64
	_ = v52
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	if l2&int64(2147483648) != int64(0) {
		v21 = base.F64_reinterpret_i64(l1 + int64(4602678819172646912))
		v22 = base.F64_mul(v21, l0)
		v23 = base.F64_add(v22, v21)
		if base.F64_lt(v23, float64(1)) == int32(0) {
			v56 = v23
		} else {
			v29 = m.G0
			v31 = v29 - int32(16)
			*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(4503599627370496)
			v34 = *(*float64)(unsafe.Add(mBase, uint32(v31)+8))
			v37 = m.G0
			*(*float64)(unsafe.Add(mBase, uint32(v37-int32(16))+8)) = base.F64_mul(v34, float64(2.2250738585072014e-308))
			v41 = float64(0)
			v42 = float64(1)
			v43 = base.F64_add(v23, v42)
			v52 = base.F64_add(base.F64_add(v43, base.F64_add(base.F64_add(v22, base.F64_sub(v21, v23)), base.F64_add(v23, base.F64_sub(v42, v43)))), float64(-1))
			if base.F64_eq(v52, v41) != 0 {
				v55 = v41
			} else {
				v55 = v52
			}
			v56 = v55
		}
		return base.F64_mul(v56, float64(2.2250738585072014e-308))
	} else {
		v13 = base.F64_reinterpret_i64(l1 + int64(-4544132024016830464))
		return base.F64_mul(base.F64_add(base.F64_mul(v13, l0), v13), float64(5.486124068793689e+303))
	}
}
func F_specialcase_2(m *base.Module, l0 float64, l1 int64, l2 int64) float64 {
	mBase := m.M
	_ = mBase
	var v14 float64
	_ = v14
	var v21 int64
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v30 float64
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v58 float64
	_ = v58
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	if l2&int64(2147483648) != int64(0) {
		v21 = l1 + int64(4602678819172646912)
		v22 = base.F64_reinterpret_i64(v21)
		v23 = base.F64_mul(v22, l0)
		v24 = base.F64_add(v23, v22)
		if base.F64_lt(base.F64_abs(v24), float64(1)) == int32(0) {
			v62 = v24
		} else {
			v30 = float64(2.2250738585072014e-308)
			v32 = m.G0
			*(*float64)(unsafe.Add(mBase, uint32(v32-int32(16))+8)) = v30
			v39 = m.G0
			*(*float64)(unsafe.Add(mBase, uint32(v39-int32(16))+8)) = base.F64_mul(v30, float64(2.2250738585072014e-308))
			if base.F64_lt(v24, float64(0)) != 0 {
				v50 = float64(-1)
			} else {
				v50 = float64(1)
			}
			v51 = base.F64_add(v24, v50)
			v58 = base.F64_sub(base.F64_add(v51, base.F64_add(base.F64_add(v23, base.F64_sub(v22, v24)), base.F64_add(v24, base.F64_sub(v50, v51)))), v50)
			if base.F64_eq(v58, float64(0)) != 0 {
				v61 = base.F64_reinterpret_i64(v21 & int64(-9223372036854775807-1))
			} else {
				v61 = v58
			}
			v62 = v61
		}
		return base.F64_mul(v62, float64(2.2250738585072014e-308))
	} else {
		v14 = base.F64_reinterpret_i64(l1 + int64(-4544132024016830464))
		return base.F64_mul(base.F64_add(base.F64_mul(v14, l0), v14), float64(5.486124068793689e+303))
	}
}
func F_spmcDequeue(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = v6
	goto L1
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v19 = v12 + (v13+int32(-1))&v8<<(uint(int32(6))%32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = v8 + int32(1)
	if v20 != v22 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v22 <= v20 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v24 == v8 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = v20
	goto L7
L6:
	;
	v26 = v24
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v26
	if v24 != v8 {
		v8 = v24
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v29 + v8
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	return v32
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = v37
	goto L1
L10:
	;
	return int32(0)
}
func F_spmcIsEmpty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v4 == v5 {
		v11 = int32(1)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v7
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
		v11 = base.B2i32(v7 == v9)
	}
	return v11
}
func F_spmcSize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = v4 - v5
	if base.Ui32(v4) < base.Ui32(v6) {
		v8 = int32(0)
	} else {
		v8 = v6
	}
	return v8
}
func F_spopCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v10 != int32(3) {
		if v10 < int32(4) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_consts[460])))
			v29 = F_lookupKeyWriteOrReply(m, l0, v22, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				if v29 == int32(0) {
					m.G0 = v8 + int32(16)
					return
				} else {
					v34 = F_checkType(m, l0, v29, int32(2))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						if v34 != 0 {
							m.G0 = v8 + int32(16)
							return
						} else {
							v36 = F_setTypePopRandom(m, v29)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
								F_notifyKeyspaceEvent(m, int32(32), int32(_a1517), v41, v43)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
									v49 = *(*int32)(unsafe.Add(mBase, _consts[744]))
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v47
									*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v36
									F_rewriteClientCommandVector(m, l0, int32(3), v8)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										F_addReplyBulk(m, l0, v36)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return
										} else {
											F_decrRefCount(m, v36)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												v60 = F_setTypeSize(m, v29)
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													if v60 != 0 {
														v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
														v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
														F_signalModifiedKey(m, l0, v75, v77)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return
														} else {
															v80 = int32(_a69)
															v82 = *(*int64)(unsafe.Add(mBase, _consts[60]))
															*(*int64)(unsafe.Add(mBase, _consts[60])) = v82 + int64(1)
															m.G0 = v8 + int32(16)
															return
														}
													} else {
														v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
														v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
														v65 = F_dbDelete(m, v62, v64)
														mBase = m.M
														v66 = m.ExcPending
														if v66 != 0 {
															return
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
															v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+28))
															F_notifyKeyspaceEvent(m, int32(4), int32(_a132), v70, v72)
															mBase = m.M
															v74 = m.ExcPending
															if v74 != 0 {
																return
															} else {
																v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
																F_signalModifiedKey(m, l0, v75, v77)
																mBase = m.M
																v79 = m.ExcPending
																if v79 != 0 {
																	return
																} else {
																	v80 = int32(_a69)
																	v82 = *(*int64)(unsafe.Add(mBase, _consts[60]))
																	*(*int64)(unsafe.Add(mBase, _consts[60])) = v82 + int64(1)
																	m.G0 = v8 + int32(16)
																	return
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
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[59]))
			F_addReplyErrorObject(m, l0, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	} else {
		F_spopWithCountCommand(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	}
}
func F_spscInit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int64
	_ = v21
	if base.I32_popcnt(l1) == int32(1) {
		v15 = F_valkey_malloc(m, l1<<(uint(int32(2))%32))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v15
			*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = int32(0)
			v21 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v21
			return
		}
	} else {
		F__serverAssert(m, int32(_a1686), int32(_a1685), int32(210))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
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
func F_spscIsEmpty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v4 == v5 {
		v11 = int32(1)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v7
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v11 = base.B2i32(v7 == v9)
	}
	return v11
}
func F_spscIsFull(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if base.Ui32(v6-v7) < base.Ui32(v9) {
		v22 = v2
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v11
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		if base.Ui32(v6-v11) < base.Ui32(v14) {
			v22 = v2
		} else {
			v16 = int32(1)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			if v17 == v18 {
				v22 = v16
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v20
				v22 = v16
			}
		}
	}
	return v22
}
func F_srand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, _consts[298])) = base.I64_extend_i32_u(l0 + int32(-1))
	return
}
func F_srandmemberWithCountCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v268 int32
	_ = v268
	var v269 int64
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
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
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int64
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v479 int32
	_ = v479
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int64
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v22 = F_getRangeLongFromObjectOrReply(m, l0, v16, int32(-2147483647), int32(2147483647), v13+int32(92), int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(96)
	return
L2:
	;
	return
L3:
	;
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v28 = *(*int32)(unsafe.Add(mBase, _consts[289]))
	v29 = F_lookupKeyReadOrReply(m, l0, v26, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v29 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v34 = F_checkType(m, l0, v29, int32(2))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v34 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v36 = F_setTypeSize(m, v29)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = v24 >> (uint(int32(31)) % 32)
	v45 = v24 ^ v43 - v43
	if v24 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[289]))
	F_addReply(m, l0, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	goto L1
L13:
	;
	v479 = v45
	goto L150
L14:
	;
	if base.Ui32(v24) < base.Ui32(v36) {
		goto L45
	} else {
		goto L46
	}
L15:
	;
	F_addReplyArrayLen(m, l0, v45)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L18
	}
L16:
	;
	if v45 != int32(1) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v52&int32(240) != int32(176) {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	if v45 == int32(1) {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v59 = int32(1000)
	if base.Ui32(v45) < base.Ui32(v59) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v62 = v45
	goto L23
L22:
	;
	v62 = v59
	goto L23
L23:
	;
	v65 = F_valkey_malloc(m, v62<<(uint(int32(4))%32))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v73 = v45
	goto L25
L25:
	;
	v77 = F_objectGetVal(m, v29)
	mBase = m.M
	if base.Ui32(v73) < base.Ui32(v62) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_valkey_free(m, v65)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L44
	}
L27:
	;
	v79 = v73
	goto L29
L28:
	;
	v79 = v62
	goto L29
L29:
	;
	F_lpRandomEntries(m, v77, v79, v65)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v82 = int32(0)
	if v73 == v82 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
	if v120&int32(4) != 0 {
		goto L41
	} else {
		goto L42
	}
L32:
	;
	v89 = v82
	goto L33
L33:
	;
	v97 = v65 + v89<<(uint(int32(4))%32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v98 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L31
L35:
	;
	v108 = v89 + int32(1)
	if v108 != v79 {
		v89 = v108
		goto L33
	} else {
		goto L40
	}
L36:
	;
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v97)+8))
	F_addReplyBulkLongLong(m, l0, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L39
	}
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	F_addReplyBulkCBuffer(m, l0, v98, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	goto L35
L40:
	;
	goto L34
L41:
	;
	goto L26
L42:
	;
	v123 = v73 - v79
	if v123 != 0 {
		v73 = v123
		goto L25
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L1
L45:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v191&int32(240) != int32(176) {
		goto L64
	} else {
		goto L65
	}
L46:
	;
	F_addReplyArrayLen(m, l0, v36)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v130 = F_setTypeInitIterator(m, v29)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L49
	}
L48:
	;
	F_setTypeReleaseIterator(m, v130)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L61
	}
L49:
	;
	v138 = F_setTypeNext(m, v130, v13+int32(88), v13+int32(84), v13+int32(72))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	if v138 == int32(-1) {
		v175 = v36
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v146 = v36
	goto L52
L52:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
	if v152 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v175 = v160
	goto L48
L54:
	;
	v160 = v146 + int32(-1)
	v167 = F_setTypeNext(m, v130, v13+int32(88), v13+int32(84), v13+int32(72))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L59
	}
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v13)+84))
	F_addReplyBulkCBuffer(m, l0, v152, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L58
	}
L56:
	;
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
	F_addReplyBulkLongLong(m, l0, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	goto L54
L59:
	;
	if v167 != int32(-1) {
		v146 = v160
		goto L52
	} else {
		goto L60
	}
L60:
	;
	goto L53
L61:
	;
	if v175 == int32(0) {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F__serverAssert(m, int32(_a1518), int32(_a1511), int32(1090))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v239 = F_hashtableCreate(m, int32(_a1519))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L2
	} else {
		goto L79
	}
L65:
	;
	v196 = F_objectGetVal(m, v29)
	mBase = m.M
	v197 = F_lpFirst(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(0)
	F_addReplyArrayLen(m, l0, v24)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v208 = v197
	v209 = v45
	goto L68
L68:
	;
	v216 = F_lpNextRandom(m, v196, v208, v13+int32(16), v209, int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L2
	} else {
		goto L72
	}
L70:
	;
	v230 = F_lpNext(m, v196, v216)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L2
	} else {
		goto L77
	}
L71:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
	F_addReplyBulkCBuffer(m, l0, v222, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L2
	} else {
		goto L76
	}
L72:
	;
	v222 = F_lpGetValue(m, v216, v13+int32(88), v13+int32(72))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	if v222 != 0 {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
	F_addReplyBulkLongLong(m, l0, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	goto L70
L76:
	;
	goto L70
L77:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v232 + int32(1)
	v237 = v209 + int32(-1)
	if v237 != 0 {
		v208 = v230
		v209 = v237
		goto L68
	} else {
		goto L78
	}
L78:
	;
	goto L1
L79:
	;
	if base.Ui32(v24*int32(3)) <= base.Ui32(v36) {
		goto L84
	} else {
		goto L85
	}
L80:
	;
	F__serverAssert(m, int32(_a1520), int32(_a1511), int32(1192))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L2
	} else {
		goto L149
	}
L81:
	;
	F__serverAssert(m, int32(_a1521), int32(_a1511), int32(1148))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L2
	} else {
		goto L148
	}
L82:
	;
	F__serverAssert(m, int32(_a1522), int32(_a1511), int32(1144))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L2
	} else {
		goto L147
	}
L83:
	;
	v387 = v13 + int32(16)
	v388 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v387)+14)) = uint8(v388)
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v387)+24)) = v388
	*(*uint8)(unsafe.Add(mBase, uint32(v387)+15)) = uint8(v388)
	*(*int32)(unsafe.Add(mBase, uint32(v387)+8)) = int32(-1)
	if v239 == v388 {
		goto L131
	} else {
		goto L132
	}
L84:
	;
	v337 = F_hashtableExpand(m, v239, v24)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L2
	} else {
		goto L115
	}
L85:
	;
	v244 = F_setTypeInitIterator(m, v29)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	v246 = F_hashtableExpand(m, v239, v36)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	v254 = F_setTypeNext(m, v244, v13+int32(88), v13+int32(84), v13+int32(72))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L2
	} else {
		goto L89
	}
L88:
	;
	F_setTypeReleaseIterator(m, v244)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L2
	} else {
		goto L105
	}
L89:
	;
	if v254 == int32(-1) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	goto L91
L91:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
	if v268 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L88
L93:
	;
	v293 = F_setTypeNext(m, v244, v13+int32(88), v13+int32(84), v13+int32(72))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L2
	} else {
		goto L103
	}
L94:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v13)+84))
	v281 = F_sdsnewlen(m, v268, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L100
	}
L95:
	;
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
	v270 = F_sdsfromlonglong(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	v272 = F_hashtableAdd(m, v239, v270)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	if v272 != 0 {
		goto L93
	} else {
		goto L98
	}
L98:
	;
	F__serverAssert(m, int32(_a1523), int32(_a1511), int32(1142))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L2
	} else {
		goto L99
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
	v283 = F_hashtableAdd(m, v239, v281)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L2
	} else {
		goto L101
	}
L101:
	;
	if v283 == int32(0) {
		goto L82
	} else {
		goto L102
	}
L102:
	;
	goto L93
L103:
	;
	if v293 != int32(-1) {
		goto L91
	} else {
		goto L104
	}
L104:
	;
	goto L92
L105:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v239)+20))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v239)+16))
	goto L106
L106:
	;
	if v309+v310 != v36 {
		goto L81
	} else {
		goto L107
	}
L107:
	;
	if base.Ui32(v36) <= base.Ui32(v45) {
		goto L83
	} else {
		goto L108
	}
L108:
	;
	v318 = v36
	goto L109
L109:
	;
	v326 = F_hashtableFairRandomEntry(m, v239, v13+int32(16))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v329 = F_hashtableDelete(m, v239, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	F_sdsfree(m, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	v335 = v318 + int32(-1)
	if base.Ui32(v45) < base.Ui32(v335) {
		v318 = v335
		goto L109
	} else {
		goto L114
	}
L114:
	;
	goto L83
L115:
	;
	v342 = int32(0)
	goto L116
L116:
	;
	v356 = F_setTypeRandomElement(m, v29, v13+int32(88), v13+int32(84), v13+int32(72))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L2
	} else {
		goto L118
	}
L117:
	;
	goto L83
L118:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
	if v358 != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v366 = F_hashtableAdd(m, v239, v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L2
	} else {
		goto L126
	}
L120:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v13)+84))
	v363 = F_sdsnewlen(m, v358, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L2
	} else {
		goto L123
	}
L121:
	;
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
	v360 = F_sdsfromlonglong(m, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v365 = v360
	goto L119
L123:
	;
	v365 = v363
	goto L119
L124:
	;
	if base.Ui32(v374) < base.Ui32(v45) {
		v342 = v374
		goto L116
	} else {
		goto L129
	}
L125:
	;
	F_sdsfree(m, v365)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L2
	} else {
		goto L128
	}
L126:
	;
	if v366 == int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v374 = v342 + int32(1)
	goto L124
L128:
	;
	v374 = v342
	goto L124
L129:
	;
	goto L117
L130:
	;
	F_addReplyArrayLen(m, l0, v45)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L2
	} else {
		goto L134
	}
L131:
	;
	goto L130
L132:
	;
	goto L131
L134:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v239)+20))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v239)+16))
	goto L135
L135:
	;
	if v45 != v408+v409 {
		goto L80
	} else {
		goto L136
	}
L136:
	;
	v416 = F_hashtableNext(m, v13+int32(16), v13+int32(12))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L2
	} else {
		goto L138
	}
L137:
	;
	F_hashtableCleanupIterator(m, v13+int32(16))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L2
	} else {
		goto L145
	}
L138:
	;
	if v416 == int32(0) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	goto L140
L140:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_addReplyBulkSds(m, l0, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L2
	} else {
		goto L142
	}
L141:
	;
	goto L137
L142:
	;
	v437 = F_hashtableNext(m, v13+int32(16), v13+int32(12))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L2
	} else {
		goto L143
	}
L143:
	;
	if v437 != 0 {
		goto L140
	} else {
		goto L144
	}
L144:
	;
	goto L141
L145:
	;
	F_hashtableRelease(m, v239)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	goto L1
L147:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	if v479 == int32(0) {
		goto L1
	} else {
		goto L152
	}
L151:
	;
	goto L1
L152:
	;
	v491 = F_setTypeRandomElement(m, v29, v13+int32(88), v13+int32(84), v13+int32(72))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L2
	} else {
		goto L153
	}
L153:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
	if v493 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
	if v502&int32(4) == int32(0) {
		v479 = v479 + int32(-1)
		goto L150
	} else {
		goto L159
	}
L155:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v13)+84))
	F_addReplyBulkCBuffer(m, l0, v493, v497)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L2
	} else {
		goto L158
	}
L156:
	;
	v494 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
	F_addReplyBulkLongLong(m, l0, v494)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L2
	} else {
		goto L157
	}
L157:
	;
	goto L154
L158:
	;
	goto L154
L159:
	;
	goto L151
}
func F_ssubscribeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
	if v12&int32(16) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return
L2:
	;
	F_addReplyError(m, l0, int32(_a864))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L8
	} else {
		goto L12
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v15 < int32(2) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v62&int32(262144) != 0 {
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v30 = int32(1)
	goto L6
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v30<<(uint(int32(2))%32))))
	v37 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, _consts[468]))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v38
	v41 = *(*int64)(unsafe.Add(mBase, _consts[469]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(16)))) = v41
	v44 = *(*int64)(unsafe.Add(mBase, _consts[470]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(8)))) = v44
	v47 = *(*int64)(unsafe.Add(mBase, _consts[471]))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v47
	v49 = F_pubsubSubscribeChannel(m, l0, v36, v10)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L4
L8:
	;
	return
L9:
	;
	v52 = v30 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v52 < v53 {
		v30 = v52
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v62 | int32(262144)
	v68 = int32(_a69)
	v70 = *(*int32)(unsafe.Add(mBase, _consts[467]))
	*(*int32)(unsafe.Add(mBase, _consts[467])) = v70 + int32(1)
	goto L1
L12:
	;
	goto L1
}
func F_startEvictionTimeProc(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _consts[262])))
	if v2 != 0 {
		return
	} else {
		v3 = int32(0)
		v4 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _consts[262])) = uint8(v4)
		v7 = *(*int32)(unsafe.Add(mBase, _consts[156]))
		v12 = F_aeCreateTimeEvent(m, v7, int64(0), int32(524), v3, v3)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	}
}
func F_startLoading(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v63 int64
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	v6 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[116])) = v6
	if l2 != v6 {
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[482])) = int32(1)
	}
	v13 = int32(0)
	v14 = F___time(m, v13)
	mBase = m.M
	v15 = int32(_a69)
	v16 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[397])) = v16
	*(*int64)(unsafe.Add(mBase, _consts[483])) = v14
	*(*int64)(unsafe.Add(mBase, _consts[484])) = v16
	*(*int64)(unsafe.Add(mBase, _consts[485])) = v16
	*(*int64)(unsafe.Add(mBase, _consts[396])) = base.I64_extend_i32_u(l0)
	*(*int64)(unsafe.Add(mBase, _consts[486])) = v16
	v39 = *(*int32)(unsafe.Add(mBase, _consts[487]))
	*(*int32)(unsafe.Add(mBase, _consts[487])) = v39 + int32(1)
	if v39 != 0 {
	} else {
		v43 = int32(0)
		v44 = F_ustime(m)
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, _consts[277])) = v44
		v48 = base.I64_div_s(v44, int64(1000))
		*(*int64)(unsafe.Add(mBase, _consts[32])) = v48
		v52 = base.I64_div_s(v44, int64(1000000))
		*(*int64)(unsafe.Add(mBase, _consts[37])) = v52
		v55 = *(*int32)(unsafe.Add(mBase, _consts[167]))
		F_lrulfu_updateClockAndPolicy(m, v48, int32(base.Ui32(v55&int32(2))>>(uint(int32(1))%32)))
		mBase = m.M
		v63 = *(*int64)(unsafe.Add(mBase, _consts[32]))
		*(*int64)(unsafe.Add(mBase, _consts[488])) = v63
	}
	F_clusterCleanSlotImportsBeforeLoad(m)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		return
	} else {
		v70 = int32(1)
		if l1&v70 != 0 {
			v75 = v70
		} else {
			v75 = l1 & int32(2)
		}
		F_moduleFireServerEvent(m, int64(3), v75, int32(0))
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return
		} else {
			return
		}
	}
}
func F_stat(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	v5 = F___fstatat(m, int32(-100), l0, l1, int32(0))
	return v5
}
func F_stopSaving(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	if l0 != 0 {
		v5 = int32(3)
	} else {
		v5 = int32(4)
	}
	F_moduleFireServerEvent(m, int64(1), v5, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_store_int(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	v3 = l2
	if l0 == int32(0) {
		return
	} else {
		switch l1 + int32(2) {
		case 0:
			*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v3)
			return
		case 1:
			*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v3)
			return
		case 2, 3:
			*(*uint32)(unsafe.Add(mBase, uint32(l0))) = uint32(v3)
			return
		default:
			return
		case 5:
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v3
			return
		}
	}
}
func F_strbuf_init(m *base.Module, l0 int32, l1 int32) {
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
	var v18 int64
	_ = v18
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 != 0 {
		v12 = l1 + int32(1)
	} else {
		v12 = int32(1023)
	}
	if base.Ui32(v12) < base.Ui32(l1) {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
		v31 = m.G3
		F_die(m, v31+int32(_a2327), v7)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
		v15 = F_emscripten_builtin_malloc(m, v12)
		mBase = m.M
		v18 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l0+int32(16)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v15
		if v15 == int32(0) {
			v36 = m.G3
			F_die(m, v36+int32(_a1692), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v25 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v25)
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_strbuf_new(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v10 == int32(0) {
		v37 = m.G3
		F_die(m, v37+int32(_a1692), int32(0))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if l0 != 0 {
			v16 = l0 + int32(1)
		} else {
			v16 = int32(1023)
		}
		if base.Ui32(v16) < base.Ui32(l0) {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			v46 = m.G3
			F_die(m, v46+int32(_a2327), v7)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v16
			v19 = F_emscripten_builtin_malloc(m, v16)
			mBase = m.M
			v22 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v10+int32(16)))) = v22
			*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v19
			if v19 == int32(0) {
				v51 = m.G3
				F_die(m, v51+int32(_a1692), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v29 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v29)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(1)
				m.G0 = v7 + int32(16)
				return v10
			}
		}
	}
}
func F_strcat(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	if l0&int32(3) == int32(0) {
		v24 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v59 = F___stpcpy(m, l0+v57, l1)
	mBase = m.M
	goto L17
L2:
	;
	v57 = v49 - l0
	goto L1
L3:
	;
	v28 = v24
	goto L11
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v13 = l0
	goto L7
L6:
	;
	v57 = l0 - l0
	goto L1
L7:
	;
	v17 = v13 + int32(1)
	if v17&int32(3) == int32(0) {
		v24 = v17
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v22 != 0 {
		v13 = v17
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v49 = v17
	goto L2
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v37 = int32(-2139062144)
	if (int32(16843008)-v34|v34)&v37 == v37 {
		v28 = v28 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v43 = v28
	goto L14
L13:
	;
	goto L12
L14:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v47 != 0 {
		v43 = v43 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v49 = v43
	goto L2
L16:
	;
	goto L15
L17:
	;
	return l0
}
func F_strcmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v6 == int32(0) {
		v29 = v5
		v30 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v30 - v29&int32(255)
L2:
	;
	if v6 != v5&int32(255) {
		v29 = v5
		v30 = v6
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = l0
	v13 = l1
	goto L4
L4:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v17 == int32(0) {
		v29 = v16
		v30 = v17
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v29 = v16
	v30 = v17
	goto L1
L6:
	;
	v20 = int32(1)
	if v17 == v16&int32(255) {
		v12 = v12 + v20
		v13 = v13 + v20
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
}
func F_strcpy(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	if (l1^l0)&int32(3) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return l0
L2:
	;
	goto L1
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v60)
	if v60&int32(255) == int32(0) {
		goto L2
	} else {
		goto L16
	}
L4:
	;
	if l1&int32(3) == int32(0) {
		v29 = l0
		v30 = l1
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v58 = l0
	v59 = l1
	v60 = v10
	goto L3
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v37 = int32(-2139062144)
	if (int32(16843008)-v34|v34)&v37 != v37 {
		v58 = v29
		v59 = v30
		v60 = v34
		goto L3
	} else {
		goto L12
	}
L7:
	;
	v15 = l0
	v16 = l1
	goto L8
L8:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v19)
	if v19 == int32(0) {
		goto L2
	} else {
		goto L10
	}
L9:
	;
	v29 = v24
	v30 = v26
	goto L6
L10:
	;
	v23 = int32(1)
	v24 = v15 + v23
	v26 = v16 + v23
	if v26&int32(3) != 0 {
		v15 = v24
		v16 = v26
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v41 = v29
	v42 = v30
	v43 = v34
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v43
	v46 = int32(4)
	v47 = v41 + v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v50 = v42 + v46
	v54 = int32(-2139062144)
	if (v48|(int32(16843008)-v48))&v54 == v54 {
		v41 = v47
		v42 = v50
		v43 = v48
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v58 = v47
	v59 = v50
	v60 = v48
	goto L3
L15:
	;
	goto L14
L16:
	;
	v67 = v58
	v68 = v59
	goto L17
L17:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)) = uint8(v71)
	v73 = int32(1)
	if v71 != 0 {
		v67 = v67 + v73
		v68 = v68 + v73
		goto L17
	} else {
		goto L19
	}
L18:
	;
	goto L2
L19:
	;
	goto L18
}
func F_strerror(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	if base.Ui32(int32(153)) < base.Ui32(l0) {
		v5 = int32(0)
	} else {
		v5 = l0
	}
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5<<(uint(int32(1))%32))+uint32(_consts[1256]))))
	return v10 + int32(_a2361)
}
func F_strerror_r(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	v5 = F___strerror_l(m, l0, l0)
	mBase = m.M
	goto L3
L1:
	;
	return v83
L2:
	;
	v77 = v60 + int32(1)
	if v77 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L3:
	;
	if v5&int32(3) == int32(0) {
		v27 = v5
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if base.Ui32(v60) < base.Ui32(l2) {
		goto L2
	} else {
		goto L20
	}
L5:
	;
	v60 = v52 - v5
	goto L4
L6:
	;
	v31 = v27
	goto L14
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v13 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v16 = v5
	goto L10
L9:
	;
	v60 = v5 - v5
	goto L4
L10:
	;
	v20 = v16 + int32(1)
	if v20&int32(3) == int32(0) {
		v27 = v20
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 != 0 {
		v16 = v20
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v52 = v20
	goto L5
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v40 = int32(-2139062144)
	if (int32(16843008)-v37|v37)&v40 == v40 {
		v31 = v31 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v46 = v31
	goto L17
L16:
	;
	goto L15
L17:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		v46 = v46 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v52 = v46
	goto L5
L19:
	;
	goto L18
L20:
	;
	if l2 == int32(0) {
		v83 = int32(68)
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v66 = l2 + int32(-1)
	if v66 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v66))) = uint8(v72)
	return int32(68)
L23:
	;
	goto L22
L24:
	;
	v69 = F__emscripten_memcpy_bulkmem(m, l1, v5, v66)
	mBase = m.M
	goto L23
L25:
	;
	v83 = int32(0)
	goto L1
L26:
	;
	goto L25
L27:
	;
	v80 = F__emscripten_memcpy_bulkmem(m, l1, v5, v77)
	mBase = m.M
	goto L26
}
func F_stringmatchlen_fuzz_test(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v26 int64
	_ = v26
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v38 int64
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v66 int64
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v100 int64
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v14 = int32(9999999)
	v15 = int32(0)
	for {
		v20 = int32(0)
		v22 = *(*int64)(unsafe.Add(mBase, _consts[298]))
		v26 = v22*int64(6364136223846793005) + int64(1)
		*(*int64)(unsafe.Add(mBase, _consts[298])) = v26
		v32 = int32(0)
		v34 = *(*int64)(unsafe.Add(mBase, _consts[298]))
		v38 = v34*int64(6364136223846793005) + int64(1)
		*(*int64)(unsafe.Add(mBase, _consts[298])) = v38
		v43 = int32(31)
		v44 = base.I32_wrap_i64(int64(base.Ui64(v38)>>(uint(int64(33))%64))) & v43
		v45 = int32(0)
		v47 = base.I32_wrap_i64(int64(base.Ui64(v26)>>(uint(int64(33))%64))) & v43
		if v47 == v45 {
		} else {
			v51 = v45
			for {
				v60 = int32(0)
				v62 = *(*int64)(unsafe.Add(mBase, _consts[298]))
				v66 = v62*int64(6364136223846793005) + int64(1)
				*(*int64)(unsafe.Add(mBase, _consts[298])) = v66
				v72 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v66)>>(uint(int64(33))%64))), int32(128))
				*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(32)+v51))) = uint8(v72)
				v75 = v51 + int32(1)
				if v75 != v47 {
					v51 = v75
					continue
				} else {
					break
				}
				break
			}
		}
		v83 = int32(0)
		if v44 == v83 {
		} else {
			v87 = v83
			for {
				v94 = int32(0)
				v96 = *(*int64)(unsafe.Add(mBase, _consts[298]))
				v100 = v96*int64(6364136223846793005) + int64(1)
				*(*int64)(unsafe.Add(mBase, _consts[298])) = v100
				v106 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v100)>>(uint(int64(33))%64))), int32(128))
				*(*uint8)(unsafe.Add(mBase, uint32(v9+v87))) = uint8(v106)
				v109 = v87 + int32(1)
				if v109 != v44 {
					v87 = v109
					continue
				} else {
					break
				}
				break
			}
		}
		v117 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+76)) = v117
		v125 = F_stringmatchlen_impl(m, v9, v44, v9+int32(32), v47, v117, v9+int32(76), v117)
		mBase = m.M
		v126 = v125 + v15
		if v14 != 0 {
			v14 = v14 + int32(-1)
			v15 = v126
			continue
		} else {
			break
		}
		break
	}
	m.G0 = v9 + int32(80)
	return v126
}
func F_stringmatchlen_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v355 int32
	_ = v355
	v8 = int32(0)
	if l6 <= int32(1000) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = l0
	v20 = l1
	v21 = l2
	v22 = l3
	goto L5
L2:
	;
	return int32(0)
L3:
	;
	return v355
L4:
	;
	v355 = base.B2i32(v333|v335 == int32(0))
	goto L3
L5:
	;
	if v22 == int32(0) {
		v333 = v20
		v335 = v22
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v294 = int32(0)
	if v287 == v294 {
		goto L111
	} else {
		goto L112
	}
L7:
	;
	if v20 == int32(0) {
		v333 = v20
		v335 = v22
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	switch v36 + int32(-63) {
	case 0:
		v273 = v19
		v274 = v20
		goto L9
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27:
		v123 = v19
		v124 = v20
		goto L12
	case 28:
		goto L14
	case 29:
		goto L13
	default:
		goto L15
	}
L9:
	;
	v286 = int32(-1)
	v287 = v274 + v286
	v288 = int32(1)
	v289 = v273 + v288
	v293 = v22 + v286
	if v293 != 0 {
		v19 = v289
		v20 = v287
		v21 = v21 + v288
		v22 = v293
		goto L5
	} else {
		goto L110
	}
L10:
	;
	v159 = v148
	v160 = v149
	v161 = int32(0)
	goto L47
L11:
	;
	v148 = v146
	v149 = v19 + int32(1)
	v150 = int32(0)
	goto L10
L12:
	;
	v125 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123))))
	if l4 != 0 {
		goto L34
	} else {
		goto L35
	}
L13:
	;
	if v20 < int32(2) {
		v123 = v19
		v124 = v20
		goto L12
	} else {
		goto L33
	}
L14:
	;
	v107 = v20 + int32(-1)
	if v107 != 0 {
		goto L30
	} else {
		goto L31
	}
L15:
	;
	if v36 != int32(42) {
		v123 = v19
		v124 = v20
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v42 = v19
	v43 = v20
	goto L19
L17:
	;
	v69 = int32(1)
	v78 = v21
	v79 = v22
	goto L24
L18:
	;
	v63 = int32(1)
	if v43 == v63 {
		v355 = v63
		goto L3
	} else {
		goto L23
	}
L19:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v55 != int32(42) {
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v66 = v19 + v20
	v67 = int32(0)
	goto L17
L21:
	;
	v61 = v43 + int32(-1)
	if v61 != 0 {
		v42 = v42 + int32(1)
		v43 = v61
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v66 = v42
	v67 = v43
	goto L17
L24:
	;
	v89 = F_stringmatchlen_impl(m, v66+v69, v67+int32(-1), v78, v79, l4, l5, l6+v69)
	mBase = m.M
	if v89 != 0 {
		v355 = v69
		goto L3
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(1)
	return int32(0)
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v90 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v98 = v79 + int32(-1)
	if v98 != 0 {
		v78 = v78 + int32(1)
		v79 = v98
		goto L24
	} else {
		goto L29
	}
L28:
	;
	return int32(0)
L29:
	;
	goto L25
L30:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v109 != int32(94) {
		v146 = v107
		goto L11
	} else {
		goto L32
	}
L31:
	;
	v146 = int32(0)
	goto L11
L32:
	;
	v148 = v20 + int32(-2)
	v149 = v19 + int32(2)
	v150 = int32(1)
	goto L10
L33:
	;
	v123 = v19 + int32(1)
	v124 = v20 + int32(-1)
	goto L12
L34:
	;
	if base.Ui32(v125+int32(-65)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v125&int32(255) == v128 {
		v273 = v123
		v274 = v124
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v355 = v8
	goto L3
L37:
	;
	v137 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21))))
	if base.Ui32(v137+int32(-65)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v136 = v125 | int32(32)
	goto L40
L39:
	;
	v136 = v125
	goto L40
L40:
	;
	goto L37
L41:
	;
	if v136 != v144 {
		v355 = v8
		goto L3
	} else {
		goto L45
	}
L42:
	;
	v144 = v137 | int32(32)
	goto L44
L43:
	;
	v144 = v137
	goto L44
L44:
	;
	goto L41
L45:
	;
	v273 = v123
	v274 = v124
	goto L9
L46:
	;
	if v150 != 0 {
		goto L106
	} else {
		goto L107
	}
L47:
	;
	if v159 < int32(2) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v159 = v258 + int32(-1)
	v160 = v259 + int32(1)
	v161 = v260
	goto L47
L50:
	;
	v184 = v182 & int32(255)
	if v184 != int32(93) {
		goto L59
	} else {
		goto L60
	}
L51:
	;
	if v159 != 0 {
		goto L57
	} else {
		goto L58
	}
L52:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v166 != int32(92) {
		v182 = v166
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+1)))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v170 == v171 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v173 = int32(1)
	goto L56
L55:
	;
	v173 = v161
	goto L56
L56:
	;
	v258 = v159 + int32(-1)
	v259 = v160 + int32(1)
	v260 = v173
	goto L49
L57:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v182 = v181
	goto L50
L58:
	;
	v266 = v160 + int32(-1)
	v267 = int32(1)
	goto L46
L59:
	;
	v187 = base.I32_extend8_s(v182)
	if v159 < int32(3) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v266 = v160
	v267 = v159
	goto L46
L61:
	;
	if l4 != 0 {
		goto L90
	} else {
		goto L91
	}
L62:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+1)))
	if v190 != int32(45) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v193 = int32(*(*int8)(unsafe.Add(mBase, uint32(v160)+2)))
	if v193 < v187 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v195 = v187
	goto L66
L65:
	;
	v195 = v193
	goto L66
L66:
	;
	if v187 < v193 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v197 = v187
	goto L69
L68:
	;
	v197 = v193
	goto L69
L69:
	;
	v198 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21))))
	if l4 == int32(0) {
		v222 = v198
		v223 = v197
		v224 = v195
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if v224 < v222 {
		goto L84
	} else {
		goto L85
	}
L71:
	;
	if base.Ui32(v197+int32(-65)) < base.Ui32(int32(26)) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if base.Ui32(v195+int32(-65)) < base.Ui32(int32(26)) {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	v207 = v197 | int32(32)
	goto L75
L74:
	;
	v207 = v197
	goto L75
L75:
	;
	goto L72
L76:
	;
	if base.Ui32(v198+int32(-65)) < base.Ui32(int32(26)) {
		goto L81
	} else {
		goto L82
	}
L77:
	;
	v214 = v195 | int32(32)
	goto L79
L78:
	;
	v214 = v195
	goto L79
L79:
	;
	goto L76
L80:
	;
	v222 = v221
	v223 = v207
	v224 = v214
	goto L70
L81:
	;
	v221 = v198 | int32(32)
	goto L83
L82:
	;
	v221 = v198
	goto L83
L83:
	;
	goto L80
L84:
	;
	v229 = v161
	goto L86
L85:
	;
	v229 = int32(1)
	goto L86
L86:
	;
	if v222 < v223 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v231 = v161
	goto L89
L88:
	;
	v231 = v229
	goto L89
L89:
	;
	v258 = v159 + int32(-2)
	v259 = v160 + int32(2)
	v260 = v231
	goto L49
L90:
	;
	if base.Ui32(v187+int32(-65)) < base.Ui32(int32(26)) {
		goto L96
	} else {
		goto L97
	}
L91:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v184 == v235 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v237 = int32(1)
	goto L94
L93:
	;
	v237 = v161
	goto L94
L94:
	;
	v258 = v159
	v259 = v160
	v260 = v237
	goto L49
L95:
	;
	v246 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21))))
	if base.Ui32(v246+int32(-65)) < base.Ui32(int32(26)) {
		goto L100
	} else {
		goto L101
	}
L96:
	;
	v245 = v187 | int32(32)
	goto L98
L97:
	;
	v245 = v187
	goto L98
L98:
	;
	goto L95
L99:
	;
	if v245 == v253 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v253 = v246 | int32(32)
	goto L102
L101:
	;
	v253 = v246
	goto L102
L102:
	;
	goto L99
L103:
	;
	v255 = int32(1)
	goto L105
L104:
	;
	v255 = v161
	goto L105
L105:
	;
	v258 = v159
	v259 = v160
	v260 = v255
	goto L49
L106:
	;
	v270 = base.B2i32(v161 == int32(0))
	goto L108
L107:
	;
	v270 = v161
	goto L108
L108:
	;
	if v270 != 0 {
		v273 = v266
		v274 = v267
		goto L9
	} else {
		goto L109
	}
L109:
	;
	return int32(0)
L110:
	;
	goto L6
L111:
	;
	v333 = int32(0)
	v335 = v294
	goto L4
L112:
	;
	v297 = v289
	v298 = v287
	goto L113
L113:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	if v310 == int32(42) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	goto L111
L115:
	;
	v317 = v298 + int32(-1)
	if v317 != 0 {
		v297 = v297 + int32(1)
		v298 = v317
		goto L113
	} else {
		goto L117
	}
L116:
	;
	v333 = int32(1)
	v335 = v294
	goto L4
L117:
	;
	goto L114
}
func F_strmapchars(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	if l0&int32(3) == int32(0) {
		v31 = l0
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return l0
L2:
	;
	if v64 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L3:
	;
	v64 = v56 - l0
	goto L2
L4:
	;
	v35 = v31
	goto L12
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v20 = l0
	goto L8
L7:
	;
	v64 = l0 - l0
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
	v73 = int32(0)
	goto L19
L19:
	;
	if l3 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L1
L21:
	;
	v112 = v73 + int32(1)
	if v112 != v64 {
		v73 = v112
		goto L19
	} else {
		goto L28
	}
L22:
	;
	v80 = l0 + v73
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v90 = int32(0)
	goto L23
L23:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v90))))
	if v81&int32(255) != v94 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L21
L25:
	;
	v100 = v90 + int32(1)
	if v100 != l3 {
		v90 = v100
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v90))))
	*(*uint8)(unsafe.Add(mBase, uint32(v80))) = uint8(v97)
	goto L21
L27:
	;
	goto L24
L28:
	;
	goto L20
}
func F_strncat(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	if l0&int32(3) == int32(0) {
		v27 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v61 = l0 + v60
	if l2 == int32(0) {
		v82 = v61
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v60 = v52 - l0
	goto L1
L3:
	;
	v31 = v27
	goto L11
L4:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v13 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v16 = l0
	goto L7
L6:
	;
	v60 = l0 - l0
	goto L1
L7:
	;
	v20 = v16 + int32(1)
	if v20&int32(3) == int32(0) {
		v27 = v20
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 != 0 {
		v16 = v20
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v52 = v20
	goto L2
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v40 = int32(-2139062144)
	if (int32(16843008)-v37|v37)&v40 == v40 {
		v31 = v31 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v46 = v31
	goto L14
L13:
	;
	goto L12
L14:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		v46 = v46 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v52 = v46
	goto L2
L16:
	;
	goto L15
L17:
	;
	v84 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v82))) = uint8(v84)
	return l0
L18:
	;
	v65 = l1
	v66 = l2
	v67 = v61
	goto L19
L19:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v69 == int32(0) {
		v82 = v67
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v82 = v74
	goto L17
L21:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v69)
	v73 = int32(1)
	v74 = v67 + v73
	v78 = v66 + int32(-1)
	if v78 != 0 {
		v65 = v65 + v73
		v66 = v78
		v67 = v74
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
}
func F_strncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return int32(0)
L3:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	return v37 - v42
L4:
	;
	v10 = l0
	v11 = l1
	v12 = l2
	v13 = v8
	goto L7
L5:
	;
	v37 = int32(0)
	v38 = l1
	goto L3
L6:
	;
	v37 = v34 & int32(255)
	v38 = v32
	goto L3
L7:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v13&int32(255) != v17 {
		v32 = v11
		v34 = v13
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v32 = v26
	v34 = int32(0)
	goto L6
L9:
	;
	if v17 == int32(0) {
		v32 = v11
		v34 = v13
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v22 = v12 + int32(-1)
	if v22 == int32(0) {
		v32 = v11
		v34 = v13
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v25 = int32(1)
	v26 = v11 + v25
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v27 != 0 {
		v10 = v10 + v25
		v11 = v26
		v12 = v22
		v13 = v27
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
}
func F_strncpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	if (l1^l0)&int32(3) != 0 {
		v68 = l0
		v69 = l1
		v70 = l2
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return l0
L2:
	;
	v102 = F___memset(m, v97, int32(0), v99)
	mBase = m.M
	goto L1
L3:
	;
	v97 = v92
	v99 = int32(0)
	goto L2
L4:
	;
	v78 = v74
	v79 = v75
	v80 = v76
	goto L23
L5:
	;
	if v70 == int32(0) {
		v92 = v68
		goto L3
	} else {
		goto L22
	}
L6:
	;
	v8 = int32(0)
	v9 = base.B2i32(l2 != v8)
	if l1&int32(3) == v8 {
		v36 = l0
		v37 = l1
		v38 = l2
		v39 = v9
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v39 == int32(0) {
		v92 = v36
		goto L3
	} else {
		goto L15
	}
L8:
	;
	if l2 == int32(0) {
		v36 = l0
		v37 = l1
		v38 = l2
		v39 = v9
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v16 = l0
	v17 = l1
	v18 = l2
	goto L10
L10:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v20)
	if v20 == int32(0) {
		v97 = v16
		v99 = v18
		goto L2
	} else {
		goto L12
	}
L11:
	;
	v36 = v25
	v37 = v31
	v38 = v27
	v39 = v29
	goto L7
L12:
	;
	v24 = int32(1)
	v25 = v16 + v24
	v27 = v18 + int32(-1)
	v28 = int32(0)
	v29 = base.B2i32(v27 != v28)
	v31 = v17 + v24
	if v31&int32(3) == v28 {
		v36 = v25
		v37 = v31
		v38 = v27
		v39 = v29
		goto L7
	} else {
		goto L13
	}
L13:
	;
	if v27 != 0 {
		v16 = v25
		v17 = v31
		v18 = v27
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v42 == int32(0) {
		v97 = v36
		v99 = v38
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v38) < base.Ui32(int32(4)) {
		v68 = v36
		v69 = v37
		v70 = v38
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v47 = v36
	v48 = v37
	v49 = v38
	goto L18
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v55 = int32(-2139062144)
	if (int32(16843008)-v52|v52)&v55 != v55 {
		v74 = v47
		v75 = v48
		v76 = v49
		goto L4
	} else {
		goto L20
	}
L19:
	;
	v68 = v61
	v69 = v63
	v70 = v65
	goto L5
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v52
	v60 = int32(4)
	v61 = v47 + v60
	v63 = v48 + v60
	v65 = v49 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v65) {
		v47 = v61
		v48 = v63
		v49 = v65
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v74 = v68
	v75 = v69
	v76 = v70
	goto L4
L23:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v82)
	if v82 == int32(0) {
		v97 = v78
		v99 = v80
		goto L2
	} else {
		goto L25
	}
L24:
	;
	v92 = v87
	goto L3
L25:
	;
	v86 = int32(1)
	v87 = v78 + v86
	v91 = v80 + int32(-1)
	if v91 != 0 {
		v78 = v87
		v79 = v79 + v86
		v80 = v91
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
}
func F_strrchr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	if l0&int32(3) == int32(0) {
		v24 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v65 = v57 + int32(1)
	goto L18
L2:
	;
	v57 = v49 - l0
	goto L1
L3:
	;
	v28 = v24
	goto L11
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v13 = l0
	goto L7
L6:
	;
	v57 = l0 - l0
	goto L1
L7:
	;
	v17 = v13 + int32(1)
	if v17&int32(3) == int32(0) {
		v24 = v17
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v22 != 0 {
		v13 = v17
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v49 = v17
	goto L2
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v37 = int32(-2139062144)
	if (int32(16843008)-v34|v34)&v37 == v37 {
		v28 = v28 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v43 = v28
	goto L14
L13:
	;
	goto L12
L14:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v47 != 0 {
		v43 = v43 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v49 = v43
	goto L2
L16:
	;
	goto L15
L17:
	;
	return v75
L18:
	;
	if v65 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v75 = v70
	goto L17
L20:
	;
	v69 = v65 + int32(-1)
	v70 = l0 + v69
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v71 != l1&int32(255) {
		v65 = v69
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v75 = int32(0)
	goto L17
L22:
	;
	goto L19
}
func F_strspn(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v6 = m.G0
	v8 = v6 - int32(32)
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(24)))) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(16)))) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v11
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return int32(0)
L3:
	;
	v37 = l1
	v39 = v21
	goto L8
L4:
	;
	v26 = l0
	goto L5
L5:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v32 == v21 {
		v26 = v26 + int32(1)
		goto L5
	} else {
		goto L7
	}
L6:
	;
	return v26 - l0
L7:
	;
	goto L6
L8:
	;
	v45 = v8 + int32(base.Ui32(v39)>>(uint(int32(3))%32))&int32(28)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v47 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v46 | v47<<(uint(v39)%32)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v51 != 0 {
		v37 = v37 + v47
		v39 = v51
		goto L8
	} else {
		goto L10
	}
L9:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v54 == int32(0) {
		v78 = l0
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	return v78 - l0
L12:
	;
	v58 = l0
	v60 = v54
	goto L13
L13:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(base.Ui32(v60)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v67)>>(uint(v60)%32))&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v78 = v73
	goto L11
L15:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	v73 = v58 + int32(1)
	if v71 != 0 {
		v58 = v73
		v60 = v71
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v78 = v58
	goto L11
L17:
	;
	goto L14
}
func F_strtoul(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, l2, int64(4294967295))
	return base.I32_wrap_i64(v5)
}
func F_subscribeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
	if v12&int32(16) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v23 < int32(2) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v17&int32(8) != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	F_addReplyError(m, l0, int32(_a863))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	goto L1
L7:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v70&int32(262144) != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	v38 = int32(1)
	goto L9
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v38<<(uint(int32(2))%32))))
	v45 = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[463]))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v46
	v49 = *(*int64)(unsafe.Add(mBase, _consts[464]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(16)))) = v49
	v52 = *(*int64)(unsafe.Add(mBase, _consts[465]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(8)))) = v52
	v55 = *(*int64)(unsafe.Add(mBase, _consts[466]))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v55
	v57 = F_pubsubSubscribeChannel(m, l0, v44, v10)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L7
L11:
	;
	v60 = v38 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v60 < v61 {
		v38 = v60
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v70 | int32(262144)
	v76 = int32(_a69)
	v78 = *(*int32)(unsafe.Add(mBase, _consts[467]))
	*(*int32)(unsafe.Add(mBase, _consts[467])) = v78 + int32(1)
	goto L1
}
func F_sunionCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = int32(0)
	F_sunionDiffGenericCommand(m, l0, v2+int32(4), v5+int32(-1), v8, v8)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		return
	}
}
func F_swapMainDbWithTempDb(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v76 int64
	_ = v76
	var v84 int64
	_ = v84
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v104 int64
	_ = v104
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v148 int64
	_ = v148
	var v152 int64
	_ = v152
	var v156 int64
	_ = v156
	var v160 int64
	_ = v160
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
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	v13 = m.G0
	v15 = v13 - int32(64)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v18 < int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_trackingInvalidateKeysOnFlush(m, int32(1))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L9
	} else {
		goto L18
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v32 = v18
	v33 = v22
	v35 = int32(0)
	goto L3
L3:
	;
	v43 = v35 << (uint(int32(2)) % 32)
	v44 = l0 + v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v45 != 0 {
		v55 = v33
		v56 = v45
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L1
L5:
	;
	v176 = v35 + int32(1)
	if v176 < v168 {
		v32 = v168
		v33 = v169
		v35 = v176
		goto L3
	} else {
		goto L17
	}
L6:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55+v43)))
	if v58 != 0 {
		v66 = v58
		v67 = v56
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v33+v43)))
	if v47 == int32(0) {
		v168 = v32
		v169 = v33
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v50 = F_createDatabase(m, v35)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v50
	v54 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v55 = v54
	v56 = v50
	goto L6
L11:
	;
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v66+int32(20))))
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(16)))) = v76
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v66+int32(28))))
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(24)))) = v84
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v66+int32(36))))
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(32)))) = v92
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v66+int32(44))))
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(40)))) = v96
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v66+int32(52))))
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(48)))) = v104
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v66+int32(60))))
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(56)))) = v112
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v66)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v114
	F_touchAllWatchedKeysInDb(m, v66, v67)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	v59 = F_createDatabase(m, v35)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v62+v43))) = v59
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v66 = v59
	v67 = v65
	goto L11
L14:
	;
	F_scanDatabaseForDeletedKeys(m, v66, v67)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v124
	v126 = int32(40)
	v129 = v67 + v126
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v129)))
	*(*int64)(unsafe.Add(mBase, uint32(v66+v126))) = v130
	v132 = int32(48)
	v135 = v67 + v132
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v135)))
	*(*int64)(unsafe.Add(mBase, uint32(v66+v132))) = v136
	v138 = int32(56)
	v141 = v67 + v138
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
	*(*int64)(unsafe.Add(mBase, uint32(v66+v138))) = v142
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v67)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v66)+32)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = v68
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(28))))
	*(*int64)(unsafe.Add(mBase, uint32(v67)+32)) = v148
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(36))))
	*(*int64)(unsafe.Add(mBase, uint32(v129))) = v152
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(44))))
	*(*int64)(unsafe.Add(mBase, uint32(v135))) = v156
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(52))))
	*(*int64)(unsafe.Add(mBase, uint32(v141))) = v160
	F_scanDatabaseForReadyKeys(m, v66)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v164 = int32(_a69)
	v165 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v167 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v168 = v167
	v169 = v165
	goto L5
L17:
	;
	goto L4
L18:
	;
	F_flushReplicaKeysWithExpireList(m, int32(1))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	m.G0 = v15 + int32(64)
	return
}
func F_swapdbCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v8 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		v19 = F_getIntFromObjectOrReply(m, l0, v15, v5+int32(28), int32(_a509))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			if v19 != 0 {
				m.G0 = v5 + int32(32)
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
				v26 = F_getIntFromObjectOrReply(m, l0, v22, v5+int32(24), int32(_a510))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					if v26 != 0 {
						m.G0 = v5 + int32(32)
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
						v30 = F_dbSwapDatabases(m, v28, v29)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							if v30 != int32(-1) {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v37
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v39
								*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(1)
								F_moduleFireServerEvent(m, int64(11), int32(0), v5+int32(8))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									v49 = int32(_a69)
									v51 = *(*int64)(unsafe.Add(mBase, _consts[60]))
									*(*int64)(unsafe.Add(mBase, _consts[60])) = v51 + int64(1)
									v56 = *(*int32)(unsafe.Add(mBase, _consts[77]))
									F_addReply(m, l0, v56)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										m.G0 = v5 + int32(32)
										return
									}
								}
							} else {
								F_addReplyError(m, l0, int32(_a483))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									m.G0 = v5 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_addReplyError(m, l0, int32(_a511))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			m.G0 = v5 + int32(32)
			return
		}
	}
}
func F_symbexec(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v489 int32
	_ = v489
	var v499 int32
	_ = v499
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	v4 = int32(0)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+75)))
	if base.Ui32(int32(250)) < base.Ui32(v22) {
		v518 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v518
L2:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+74)))
	if v25&int32(5) == int32(4) {
		v518 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+73)))
	if base.Ui32(v22) < base.Ui32(v25&int32(1)+v32) {
		v518 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v36 < v35 {
		v518 = v4
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v38 < int32(1) {
		v518 = v4
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v42 = int32(0)
	if base.B2i32(v41 == v42)|base.B2i32(v41 == v38) == v42 {
		v518 = v4
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+v38<<(uint(int32(2))%32)+int32(-4))))
	if v54&int32(63) != int32(30) {
		v518 = v4
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v60 = v38 + int32(-1)
	if int32(1) <= l1 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v48+v499<<(uint(int32(2))%32))))
	v518 = v514
	goto L1
L10:
	;
	v73 = int32(0)
	v76 = v60
	goto L12
L11:
	;
	v499 = v60
	goto L9
L12:
	;
	v90 = v48 + v73<<(uint(int32(2))%32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v93 = v91 & int32(63)
	if base.Ui32(v93) <= base.Ui32(int32(37)) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v499 = v476
	goto L9
L14:
	;
	v101 = int32(base.Ui32(v91)>>(uint(int32(6))%32)) & int32(255)
	if base.Ui32(v101) < base.Ui32(v22) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	return int32(0)
L16:
	;
	v105 = m.G400
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+v93))))
	v108 = base.I32_extend8_s(v107)
	v109 = int32(0)
	switch v107 & int32(3) {
	default:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		v244 = v109
		v245 = v109
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	if int32(-1) < v108 {
		goto L54
	} else {
		goto L55
	}
L19:
	;
	v176 = int32(base.Ui32(v91)>>(uint(int32(14))%32)) + int32(-131071)
	if v107&int32(48) != int32(32) {
		v244 = v109
		v245 = v176
		goto L18
	} else {
		goto L43
	}
L20:
	;
	v164 = int32(base.Ui32(v91) >> (uint(int32(14)) % 32))
	v165 = int32(48)
	if v107&v165 != v165 {
		v244 = v109
		v245 = v164
		goto L18
	} else {
		goto L41
	}
L21:
	;
	v114 = int32(base.Ui32(v91) >> (uint(int32(23)) % 32))
	switch int32(base.Ui32(v107)>>(uint(int32(4))%32)) & int32(3) {
	default:
		goto L25
	case 1:
		goto L22
	case 2:
		goto L24
	case 3:
		goto L23
	}
L22:
	;
	v138 = int32(base.Ui32(v91) >> (uint(int32(14)) % 32))
	v140 = v138 & int32(511)
	switch int32(base.Ui32(v108)>>(uint(int32(2))%32)) & int32(3) {
	default:
		goto L34
	case 1:
		v244 = v140
		v245 = v114
		goto L18
	case 2:
		goto L33
	case 3:
		goto L32
	}
L23:
	;
	if int32(-1) < v91 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if base.Ui32(v114) < base.Ui32(v22) {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	if base.Ui32(v91) < base.Ui32(int32(_a14)) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	return int32(0)
L27:
	;
	return int32(0)
L28:
	;
	if base.Ui32(v114) < base.Ui32(v22) {
		goto L22
	} else {
		goto L31
	}
L29:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v114&int32(255) < v130 {
		goto L22
	} else {
		goto L30
	}
L30:
	;
	return int32(0)
L31:
	;
	return int32(0)
L32:
	;
	if base.Ui32(v140) < base.Ui32(int32(256)) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	if base.Ui32(v140) < base.Ui32(v22) {
		v244 = v140
		v245 = v114
		goto L18
	} else {
		goto L36
	}
L34:
	;
	v145 = int32(0)
	if v140 == v145 {
		v244 = v145
		v245 = v114
		goto L18
	} else {
		goto L35
	}
L35:
	;
	v518 = v145
	goto L1
L36:
	;
	return int32(0)
L37:
	;
	if base.Ui32(v140) < base.Ui32(v22) {
		v244 = v140
		v245 = v114
		goto L18
	} else {
		goto L40
	}
L38:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v138&int32(255) < v156 {
		v244 = v140
		v245 = v114
		goto L18
	} else {
		goto L39
	}
L39:
	;
	return int32(0)
L40:
	;
	return int32(0)
L41:
	;
	v169 = int32(0)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v164 < v171 {
		v244 = v169
		v245 = v164
		goto L18
	} else {
		goto L42
	}
L42:
	;
	v518 = v169
	goto L1
L43:
	;
	v181 = int32(0)
	v182 = v73 + v176
	v184 = v182 + int32(1)
	if v184 < v181 {
		v518 = v181
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v38 <= v184 {
		v518 = v181
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v184 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v198 = int32(0)
	goto L49
L47:
	;
	v244 = int32(0)
	v245 = v176
	goto L18
L48:
	;
	v223 = int32(0)
	if v222&int32(1) != 0 {
		v518 = v223
		goto L1
	} else {
		goto L53
	}
L49:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v48+(v182-v198)<<(uint(int32(2))%32))))
	if v214&int32(8372287) != int32(34) {
		v222 = v198
		goto L48
	} else {
		goto L51
	}
L50:
	;
	v222 = v184
	goto L48
L51:
	;
	v220 = v198 + int32(1)
	if v220 != v184 {
		v198 = v220
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v244 = v223
	v245 = v176
	goto L18
L54:
	;
	if int32(base.Ui32(v108&int32(64))>>(uint(int32(6))%32)) != 0 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v249 = int32(0)
	if v38 <= v73+int32(2) {
		v518 = v249
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v90+int32(4))))
	if v255&int32(63) != int32(22) {
		v518 = v249
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	v265 = v73
	goto L60
L59:
	;
	v265 = v76
	goto L60
L60:
	;
	if v101 == l2 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v267 = v265
	goto L63
L62:
	;
	v267 = v76
	goto L63
L63:
	;
	switch v93 + int32(-2) {
	case 0:
		goto L79
	case 1:
		goto L78
	case 2, 6:
		goto L77
	case 3, 5:
		goto L76
	default:
		v473 = v73
		v476 = v267
		goto L64
	case 9:
		goto L75
	case 19:
		goto L74
	case 20:
		goto L71
	case 26, 27:
		goto L70
	case 28:
		goto L69
	case 29, 30:
		goto L72
	case 31:
		goto L73
	case 32:
		goto L68
	case 34:
		goto L67
	case 35:
		goto L66
	}
L64:
	;
	v489 = v473 + int32(1)
	if v489 < l1 {
		v73 = v489
		v76 = v476
		goto L12
	} else {
		goto L150
	}
L65:
	;
	if l2 == int32(255) {
		goto L147
	} else {
		goto L148
	}
L66:
	;
	v425 = int32(0)
	if v25&int32(6) != int32(2) {
		v518 = v425
		goto L1
	} else {
		goto L138
	}
L67:
	;
	v382 = int32(0)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v383 <= v245 {
		v518 = v382
		goto L1
	} else {
		goto L131
	}
L68:
	;
	v374 = int32(0)
	if v245 < int32(1) {
		goto L126
	} else {
		goto L127
	}
L69:
	;
	if v245 < int32(2) {
		v473 = v73
		v476 = v267
		goto L64
	} else {
		goto L124
	}
L70:
	;
	v338 = int32(0)
	if v245 == v338 {
		goto L109
	} else {
		goto L110
	}
L71:
	;
	v323 = int32(0)
	v326 = int32(1)
	v327 = v73 + v245 + v326
	if (base.B2i32(v327 <= v73)|base.B2i32(l1 < v327))&v326 != 0 {
		goto L103
	} else {
		goto L104
	}
L72:
	;
	if base.Ui32(v101+int32(3)) < base.Ui32(v22) {
		goto L71
	} else {
		goto L102
	}
L73:
	;
	v309 = int32(0)
	if v244 == v309 {
		v518 = v309
		goto L1
	} else {
		goto L97
	}
L74:
	;
	if v245 < v244 {
		v473 = v73
		v476 = v267
		goto L64
	} else {
		goto L96
	}
L75:
	;
	v300 = v101 + int32(1)
	if base.Ui32(v300) < base.Ui32(v22) {
		goto L91
	} else {
		goto L92
	}
L76:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v291 = int32(4)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v290+v245<<(uint(v291)%32))+8))
	if v294 == v291 {
		v473 = v73
		v476 = v267
		goto L64
	} else {
		goto L90
	}
L77:
	;
	if v245 < v36 {
		v473 = v73
		v476 = v267
		goto L64
	} else {
		goto L89
	}
L78:
	;
	if v245 < l2 {
		goto L83
	} else {
		goto L84
	}
L79:
	;
	if v244 != int32(1) {
		v473 = v73
		v476 = v267
		goto L64
	} else {
		goto L80
	}
L80:
	;
	v272 = int32(0)
	if v38 <= v73+int32(2) {
		v518 = v272
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v90+int32(4))))
	if v278&int32(8372287) != int32(34) {
		v473 = v73
		v476 = v267
		goto L64
	} else {
		goto L82
	}
L82:
	;
	v518 = v272
	goto L1
L83:
	;
	v284 = v267
	goto L85
L84:
	;
	v284 = v73
	goto L85
L85:
	;
	if l2 < v101 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v286 = v267
	goto L88
L87:
	;
	v286 = v284
	goto L88
L88:
	;
	v473 = v73
	v476 = v286
	goto L64
L89:
	;
	return int32(0)
L90:
	;
	return int32(0)
L91:
	;
	if l2 == v300 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	return int32(0)
L93:
	;
	v305 = v73
	goto L95
L94:
	;
	v305 = v267
	goto L95
L95:
	;
	v473 = v73
	v476 = v305
	goto L64
L96:
	;
	return int32(0)
L97:
	;
	v313 = v101 + int32(2)
	if base.Ui32(v22) <= base.Ui32(v244+v313) {
		v518 = v309
		goto L1
	} else {
		goto L98
	}
L98:
	;
	if l2 < v313 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v317 = v267
	goto L101
L100:
	;
	v317 = v73
	goto L101
L101:
	;
	v473 = v73
	v476 = v317
	goto L64
L102:
	;
	return int32(0)
L103:
	;
	v333 = v323
	goto L105
L104:
	;
	v333 = v245
	goto L105
L105:
	;
	if l2 == int32(255) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v336 = v323
	goto L108
L107:
	;
	v336 = v333
	goto L108
L108:
	;
	v473 = v336 + v73
	v476 = v267
	goto L64
L109:
	;
	if v244 != 0 {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	if v22 < v245+v101 {
		v518 = v338
		goto L1
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	if l2 < v101 {
		goto L121
	} else {
		goto L122
	}
L113:
	;
	v357 = v244 + int32(-1)
	if v357 == int32(0) {
		goto L112
	} else {
		goto L119
	}
L114:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v90+int32(4))))
	v347 = v345 & int32(63)
	if base.Ui32(v347+int32(-28)) < base.Ui32(int32(3)) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	if base.Ui32(v345) <= base.Ui32(int32(8388607)) {
		goto L112
	} else {
		goto L118
	}
L116:
	;
	if v347 != int32(34) {
		v518 = v338
		goto L1
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v518 = v338
	goto L1
L119:
	;
	if v22 < v357+v101 {
		v518 = v338
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L112
L121:
	;
	v365 = v76
	goto L123
L122:
	;
	v365 = v73
	goto L123
L123:
	;
	v473 = v73
	v476 = v365
	goto L64
L124:
	;
	if v101+v245+int32(-1) <= v22 {
		v473 = v73
		v476 = v267
		goto L64
	} else {
		goto L125
	}
L125:
	;
	return int32(0)
L126:
	;
	if v244 != 0 {
		v473 = v73
		v476 = v267
		goto L64
	} else {
		goto L129
	}
L127:
	;
	if v22 <= v245+v101 {
		v518 = v374
		goto L1
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v380 = v73 + int32(1)
	if v380 < v60 {
		v473 = v380
		v476 = v267
		goto L64
	} else {
		goto L130
	}
L130:
	;
	v518 = v374
	goto L1
L131:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v385+v245<<(uint(int32(2))%32))))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+72)))
	v391 = v73 + v390
	if v38 <= v391 {
		v518 = v382
		goto L1
	} else {
		goto L132
	}
L132:
	;
	if v390 == int32(0) {
		goto L65
	} else {
		goto L133
	}
L133:
	;
	v414 = int32(1)
	goto L134
L134:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v414<<(uint(int32(2))%32)))))
	if v419&int32(59) != 0 {
		v518 = v382
		goto L1
	} else {
		goto L136
	}
L136:
	;
	if v414 == v390 {
		goto L65
	} else {
		goto L137
	}
L137:
	;
	v414 = v414 + int32(1)
	goto L134
L138:
	;
	v427 = v245 + int32(-1)
	if v245 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	if v22 < v427+v101 {
		v518 = v425
		goto L1
	} else {
		goto L146
	}
L140:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v90+int32(4))))
	v432 = v430 & int32(63)
	if base.Ui32(v432+int32(-28)) < base.Ui32(int32(3)) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	if base.Ui32(int32(8388607)) < base.Ui32(v430) {
		v518 = v425
		goto L1
	} else {
		goto L144
	}
L142:
	;
	if v432 != int32(34) {
		v518 = v425
		goto L1
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	if v427+v101 <= v22 {
		v473 = v73
		v476 = v267
		goto L64
	} else {
		goto L145
	}
L145:
	;
	v518 = v425
	goto L1
L146:
	;
	v473 = v73
	v476 = v267
	goto L64
L147:
	;
	v467 = v73
	goto L149
L148:
	;
	v467 = v391
	goto L149
L149:
	;
	v473 = v467
	v476 = v267
	goto L64
L150:
	;
	goto L13
}
func F_sysconf(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	if base.Ui32(int32(250)) < base.Ui32(l0) {
		*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(28)
		return int32(-1)
	} else {
		v9 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_consts[1262]))))
		if v9 != 0 {
			if int32(-2) < v9 {
				v40 = v9
				return v40
			} else {
				switch v9&int32(255) + int32(-1) {
				case 0:
					v40 = int32(200809)
					return v40
				case 1:
					return int32(131072)
				case 2:
					return int32(32768)
				case 3:
					return int32(65536)
				case 4, 10:
					return int32(2147483647)
				case 5, 6:
					return int32(1)
				case 7, 8:
					v33 = m.Env.Emscripten_get_heap_max(m)
					mBase = m.M
					return int32(base.Ui32(v33) >> (uint(int32(16)) % 32))
				case 9:
					return int32(0)
				default:
					v40 = v9
					return v40
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(28)
			return int32(-1)
		}
	}
}
func F_syslog(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	F___vsyslog(m, l0, l1, l2)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
