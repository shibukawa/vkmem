package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_latencyAllCommandsFillCDF(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(64)
	m.G0 = v10
	v13 = v10 + int32(16)
	v14 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v4
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v14)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(-1)
	if l1 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = F_hashtableNext(m, v10+int32(16), v10+int32(12))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L1
L3:
	;
	goto L4
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v13
	goto L2
L5:
	;
	F_hashtableCleanupIterator(m, v10+int32(16))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L26
	}
L6:
	;
	return
L7:
	;
	if v36 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+148))
	if v48 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L5
L11:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	if v85 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-1)))))
	switch v55 & int32(7) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	case 4:
		goto L14
	default:
		v72 = int32(0)
		goto L13
	}
L13:
	;
	F_addReplyBulkCBuffer(m, l0, v52, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L19
	}
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-17))))
	v72 = v71
	goto L13
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-9))))
	v72 = v68
	goto L13
L16:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(-5)))))
	v72 = v65
	goto L13
L17:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-3)))))
	v72 = v62
	goto L13
L18:
	;
	v72 = int32(base.Ui32(v55) >> (uint(int32(3)) % 32))
	goto L13
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v47)+148))
	F_fillCommandCDF(m, l0, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v78 + int32(1)
	goto L11
L21:
	;
	v95 = F_hashtableNext(m, v10+int32(16), v10+int32(12))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v47)+200))
	F_latencyAllCommandsFillCDF(m, l0, v88, l2)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if v95 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	goto L10
L26:
	;
	m.G0 = v10 + int32(64)
	return
}
func F_latencyCommand(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
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
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
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
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
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
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int64
	_ = v379
	var v380 int64
	_ = v380
	var v381 int64
	_ = v381
	var v382 int64
	_ = v382
	var v383 int64
	_ = v383
	var v384 int64
	_ = v384
	var v385 int64
	_ = v385
	var v387 int64
	_ = v387
	var v389 int64
	_ = v389
	var v391 int64
	_ = v391
	var v393 int64
	_ = v393
	var v395 int64
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int64
	_ = v501
	var v502 int64
	_ = v502
	var v503 int64
	_ = v503
	var v504 int64
	_ = v504
	var v505 int64
	_ = v505
	var v506 int64
	_ = v506
	var v507 int64
	_ = v507
	var v509 int64
	_ = v509
	var v511 int64
	_ = v511
	var v513 int64
	_ = v513
	var v515 int64
	_ = v515
	var v517 int64
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v582 int64
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v705 int64
	_ = v705
	var v710 int64
	_ = v710
	var v715 int64
	_ = v715
	var v720 int64
	_ = v720
	var v725 int64
	_ = v725
	var v728 int64
	_ = v728
	var v731 int64
	_ = v731
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = F_objectGetVal(m, v12)
	mBase = m.M
	v14 = int32(_a878)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v9 + int32(80)
	return
L2:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v70 = F_objectGetVal(m, v69)
	mBase = m.M
	v71 = int32(_a879)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v74 != 0 {
		goto L27
	} else {
		goto L28
	}
L3:
	;
	if v49-v51 != 0 {
		goto L2
	} else {
		goto L15
	}
L4:
	;
	v49 = F_tolower(m, v45)
	mBase = m.M
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	v51 = F_tolower(m, v50)
	mBase = m.M
	goto L3
L5:
	;
	v19 = v13
	v20 = v14
	v21 = v17
	goto L8
L6:
	;
	v45 = int32(0)
	v46 = v14
	goto L4
L7:
	;
	v45 = v42 & int32(255)
	v46 = v41
	goto L4
L8:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v23 == int32(0) {
		v41 = v20
		v42 = v21
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v41 = v35
	v42 = int32(0)
	goto L7
L10:
	;
	v27 = v21 & int32(255)
	if v27 == v23 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v34 = int32(1)
	v35 = v20 + v34
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v36 != 0 {
		v19 = v19 + v34
		v20 = v35
		v21 = v36
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v29 = F_tolower(m, v27)
	mBase = m.M
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v31 = F_tolower(m, v30)
	mBase = m.M
	if v29 == v31 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v41 = v20
	v42 = v33
	goto L7
L14:
	;
	goto L9
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v53 != int32(3) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[494]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v60 = F_objectGetVal(m, v59)
	mBase = m.M
	v61 = F_dictFetchValue(m, v57, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_latencyCommandReplyWithSamples(m, l0, v61)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L18
	} else {
		goto L22
	}
L18:
	;
	return
L19:
	;
	if v61 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	F_addReplyArrayLen(m, l0, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L1
L22:
	;
	goto L1
L23:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v739)+8))
	v741 = F_objectGetVal(m, v740)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v741
	F_addReplyErrorFormat(m, l0, int32(_a880), v9)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L18
	} else {
		goto L215
	}
L24:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v154 = F_objectGetVal(m, v153)
	mBase = m.M
	v155 = int32(_a881)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v158 != 0 {
		goto L55
	} else {
		goto L56
	}
L25:
	;
	if v106-v108 != 0 {
		goto L24
	} else {
		goto L37
	}
L26:
	;
	v106 = F_tolower(m, v102)
	mBase = m.M
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v108 = F_tolower(m, v107)
	mBase = m.M
	goto L25
L27:
	;
	v76 = v70
	v77 = v71
	v78 = v74
	goto L30
L28:
	;
	v102 = int32(0)
	v103 = v71
	goto L26
L29:
	;
	v102 = v99 & int32(255)
	v103 = v98
	goto L26
L30:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v80 == int32(0) {
		v98 = v77
		v99 = v78
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v98 = v92
	v99 = int32(0)
	goto L29
L32:
	;
	v84 = v78 & int32(255)
	if v84 == v80 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v91 = int32(1)
	v92 = v77 + v91
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v93 != 0 {
		v76 = v76 + v91
		v77 = v92
		v78 = v93
		goto L30
	} else {
		goto L36
	}
L34:
	;
	v86 = F_tolower(m, v84)
	mBase = m.M
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v88 = F_tolower(m, v87)
	mBase = m.M
	if v86 == v88 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v98 = v77
	v99 = v90
	goto L29
L36:
	;
	goto L31
L37:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v110 != int32(3) {
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[494]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v117 = F_objectGetVal(m, v116)
	mBase = m.M
	v118 = F_dictFind(m, v114, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	if v118 == int32(0) {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	goto L41
L41:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	goto L48
L42:
	;
	F_addReplyVerbatim(m, l0, v125, v146, int32(_a701))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L18
	} else {
		goto L50
	}
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v125+int32(-17))))
	v146 = v145
	goto L42
L44:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v125+int32(-9))))
	v146 = v142
	goto L42
L45:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125+int32(-5)))))
	v146 = v139
	goto L42
L46:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125+int32(-3)))))
	v146 = v136
	goto L42
L47:
	;
	v146 = int32(base.Ui32(v129) >> (uint(int32(3)) % 32))
	goto L42
L48:
	;
	v125 = F_latencyCommandGenSparkeline(m, v124, v122)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125+int32(-1)))))
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
L50:
	;
	F_sdsfree(m, v125)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	goto L1
L52:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	v201 = F_objectGetVal(m, v200)
	mBase = m.M
	v202 = int32(_a882)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	if v205 != 0 {
		goto L71
	} else {
		goto L72
	}
L53:
	;
	if v190-v192 != 0 {
		goto L52
	} else {
		goto L65
	}
L54:
	;
	v190 = F_tolower(m, v186)
	mBase = m.M
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v192 = F_tolower(m, v191)
	mBase = m.M
	goto L53
L55:
	;
	v160 = v154
	v161 = v155
	v162 = v158
	goto L58
L56:
	;
	v186 = int32(0)
	v187 = v155
	goto L54
L57:
	;
	v186 = v183 & int32(255)
	v187 = v182
	goto L54
L58:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v164 == int32(0) {
		v182 = v161
		v183 = v162
		goto L57
	} else {
		goto L60
	}
L59:
	;
	v182 = v176
	v183 = int32(0)
	goto L57
L60:
	;
	v168 = v162 & int32(255)
	if v168 == v164 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v175 = int32(1)
	v176 = v161 + v175
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+1)))
	if v177 != 0 {
		v160 = v160 + v175
		v161 = v176
		v162 = v177
		goto L58
	} else {
		goto L64
	}
L62:
	;
	v170 = F_tolower(m, v168)
	mBase = m.M
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	v172 = F_tolower(m, v171)
	mBase = m.M
	if v170 == v172 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v182 = v161
	v183 = v174
	goto L57
L64:
	;
	goto L59
L65:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v194 != int32(2) {
		goto L52
	} else {
		goto L66
	}
L66:
	;
	F_latencyCommandReplyWithLatestEvents(m, l0)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L18
	} else {
		goto L67
	}
L67:
	;
	goto L1
L68:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	v274 = F_objectGetVal(m, v273)
	mBase = m.M
	v275 = int32(_a50)
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	if v278 != 0 {
		goto L95
	} else {
		goto L96
	}
L69:
	;
	if v237-v239 != 0 {
		goto L68
	} else {
		goto L81
	}
L70:
	;
	v237 = F_tolower(m, v233)
	mBase = m.M
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	v239 = F_tolower(m, v238)
	mBase = m.M
	goto L69
L71:
	;
	v207 = v201
	v208 = v202
	v209 = v205
	goto L74
L72:
	;
	v233 = int32(0)
	v234 = v202
	goto L70
L73:
	;
	v233 = v230 & int32(255)
	v234 = v229
	goto L70
L74:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	if v211 == int32(0) {
		v229 = v208
		v230 = v209
		goto L73
	} else {
		goto L76
	}
L75:
	;
	v229 = v223
	v230 = int32(0)
	goto L73
L76:
	;
	v215 = v209 & int32(255)
	if v215 == v211 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v222 = int32(1)
	v223 = v208 + v222
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
	if v224 != 0 {
		v207 = v207 + v222
		v208 = v223
		v209 = v224
		goto L74
	} else {
		goto L80
	}
L78:
	;
	v217 = F_tolower(m, v215)
	mBase = m.M
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	v219 = F_tolower(m, v218)
	mBase = m.M
	if v217 == v219 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	v229 = v208
	v230 = v221
	goto L73
L80:
	;
	goto L75
L81:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v241 != int32(2) {
		goto L68
	} else {
		goto L82
	}
L82:
	;
	v245 = F_createLatencyReport(m)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L18
	} else {
		goto L89
	}
L83:
	;
	F_addReplyVerbatim(m, l0, v245, v266, int32(_a701))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L18
	} else {
		goto L90
	}
L84:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v245+int32(-17))))
	v266 = v265
	goto L83
L85:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v245+int32(-9))))
	v266 = v262
	goto L83
L86:
	;
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v245+int32(-5)))))
	v266 = v259
	goto L83
L87:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+int32(-3)))))
	v266 = v256
	goto L83
L88:
	;
	v266 = int32(base.Ui32(v249) >> (uint(int32(3)) % 32))
	goto L83
L89:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+int32(-1)))))
	switch v249 & int32(7) {
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
		v266 = int32(0)
		goto L83
	}
L90:
	;
	F_sdsfree(m, v245)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L18
	} else {
		goto L91
	}
L91:
	;
	goto L1
L92:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v589)+4))
	v591 = F_objectGetVal(m, v590)
	mBase = m.M
	v592 = int32(_a883)
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591))))
	if v595 != 0 {
		goto L180
	} else {
		goto L181
	}
L93:
	;
	if v310-v312 != 0 {
		goto L92
	} else {
		goto L105
	}
L94:
	;
	v310 = F_tolower(m, v306)
	mBase = m.M
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	v312 = F_tolower(m, v311)
	mBase = m.M
	goto L93
L95:
	;
	v280 = v274
	v281 = v275
	v282 = v278
	goto L98
L96:
	;
	v306 = int32(0)
	v307 = v275
	goto L94
L97:
	;
	v306 = v303 & int32(255)
	v307 = v302
	goto L94
L98:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v284 == int32(0) {
		v302 = v281
		v303 = v282
		goto L97
	} else {
		goto L100
	}
L99:
	;
	v302 = v296
	v303 = int32(0)
	goto L97
L100:
	;
	v288 = v282 & int32(255)
	if v288 == v284 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v295 = int32(1)
	v296 = v281 + v295
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)))
	if v297 != 0 {
		v280 = v280 + v295
		v281 = v296
		v282 = v297
		goto L98
	} else {
		goto L104
	}
L102:
	;
	v290 = F_tolower(m, v288)
	mBase = m.M
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	v292 = F_tolower(m, v291)
	mBase = m.M
	if v290 == v292 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	v302 = v281
	v303 = v294
	goto L97
L104:
	;
	goto L99
L105:
	;
	v314 = int32(2)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v315 < v314 {
		goto L92
	} else {
		goto L106
	}
L106:
	;
	if v315 == int32(2) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _consts[494]))
	v345 = F_dictGetSafeIterator(m, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L18
	} else {
		goto L116
	}
L108:
	;
	v323 = v314
	v325 = int32(0)
	goto L109
L109:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v327+v323<<(uint(int32(2))%32))))
	v332 = F_objectGetVal(m, v331)
	mBase = m.M
	v333 = F_latencyResetEvent(m, v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L18
	} else {
		goto L111
	}
L110:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v335))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L18
	} else {
		goto L113
	}
L111:
	;
	v335 = v333 + v325
	v337 = v323 + int32(1)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v337 < v338 {
		v323 = v337
		v325 = v335
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	goto L1
L114:
	;
	F_dictReleaseIterator(m, v345)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L18
	} else {
		goto L175
	}
L115:
	;
	v458 = int32(0)
	v460 = v450
	goto L144
L116:
	;
	v354 = v345 + int32(20)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v345)+16))
	if v355 != 0 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	if v450 != 0 {
		goto L115
	} else {
		goto L143
	}
L118:
	;
	v361 = v354
	v362 = v358
	goto L121
L119:
	;
	v358 = int32(1)
	goto L118
L120:
	;
	v358 = int32(0)
	goto L118
L121:
	;
	switch v362 {
	case 0:
		goto L126
	default:
		goto L125
	}
L123:
	;
	v362 = int32(0)
	goto L121
L124:
	;
	goto L117
L125:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v345)+16)) = v442
	if v442 == int32(0) {
		goto L123
	} else {
		goto L142
	}
L126:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if v366 != int32(-1) {
		v405 = v366
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v406 = int32(1)
	v407 = v405 + v406
	*(*int32)(unsafe.Add(mBase, uint32(v345)+4)) = v407
	v409 = int32(0)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v345)+8))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412+v413+int32(26)))))
	if v417 == int32(255) {
		goto L136
	} else {
		goto L137
	}
L128:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v345)+8))
	if v370 != 0 {
		v405 = int32(-1)
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	if v372 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)+20))
	if v399 != int32(-1) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v379 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v371)+16)))
	v380 = int64(*(*int8)(unsafe.Add(mBase, uint32(v371)+27)))
	v381 = int64(*(*int32)(unsafe.Add(mBase, uint32(v371)+8)))
	v382 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v371)+12)))
	v383 = int64(*(*int8)(unsafe.Add(mBase, uint32(v371)+26)))
	v384 = int64(*(*int32)(unsafe.Add(mBase, uint32(v371)+4)))
	v385 = F_wangHash64(m, v384)
	mBase = m.M
	v387 = F_wangHash64(m, v383+v385)
	mBase = m.M
	v389 = F_wangHash64(m, v382+v387)
	mBase = m.M
	v391 = F_wangHash64(m, v381+v389)
	mBase = m.M
	v393 = F_wangHash64(m, v380+v391)
	mBase = m.M
	v395 = F_wangHash64(m, v379+v393)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v345)+24)) = v395
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v398 = v397
	goto L130
L132:
	;
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v371)+24)))
	v377 = v375 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+24)) = uint16(v377)
	v398 = v371
	goto L130
L133:
	;
	v405 = v399 + int32(-1)
	goto L127
L134:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	v405 = v402
	goto L127
L135:
	;
	v432 = int32(2)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v412+v430<<(uint(v432)%32)+int32(4))))
	v361 = v437 + v431<<(uint(v432)%32)
	v362 = int32(1)
	goto L121
L136:
	;
	v421 = v409
	goto L138
L137:
	;
	v421 = v406 << (uint(v417) % 32)
	goto L138
L138:
	;
	if v407 < v421 {
		v430 = v413
		v431 = v407
		goto L135
	} else {
		goto L139
	}
L139:
	;
	if v413 != 0 {
		v450 = v409
		goto L124
	} else {
		goto L140
	}
L140:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v412)+20))
	if v423 == int32(-1) {
		v450 = v409
		goto L124
	} else {
		goto L141
	}
L141:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v345)+4)) = int64(4294967296)
	v430 = int32(1)
	v431 = int32(0)
	goto L135
L142:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v442)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v354))) = v446
	v450 = v442
	goto L124
L143:
	;
	v582 = int64(0)
	goto L114
L144:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	goto L146
L145:
	;
	v582 = base.I64_extend_i32_u(v468)
	goto L114
L146:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _consts[494]))
	v465 = F_dictDelete(m, v464, v462)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L18
	} else {
		goto L147
	}
L147:
	;
	v468 = v458 + int32(1)
	v476 = v345 + int32(20)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v345)+16))
	if v477 != 0 {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	if v572 != 0 {
		v458 = v468
		v460 = v572
		goto L144
	} else {
		goto L174
	}
L149:
	;
	v483 = v476
	v484 = v480
	goto L152
L150:
	;
	v480 = int32(1)
	goto L149
L151:
	;
	v480 = int32(0)
	goto L149
L152:
	;
	switch v484 {
	case 0:
		goto L157
	default:
		goto L156
	}
L154:
	;
	v484 = int32(0)
	goto L152
L155:
	;
	goto L148
L156:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	*(*int32)(unsafe.Add(mBase, uint32(v345)+16)) = v564
	if v564 == int32(0) {
		goto L154
	} else {
		goto L173
	}
L157:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if v488 != int32(-1) {
		v527 = v488
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v528 = int32(1)
	v529 = v527 + v528
	*(*int32)(unsafe.Add(mBase, uint32(v345)+4)) = v529
	v531 = int32(0)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v345)+8))
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534+v535+int32(26)))))
	if v539 == int32(255) {
		goto L167
	} else {
		goto L168
	}
L159:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v345)+8))
	if v492 != 0 {
		v527 = int32(-1)
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	if v494 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v520)+20))
	if v521 != int32(-1) {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v501 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v493)+16)))
	v502 = int64(*(*int8)(unsafe.Add(mBase, uint32(v493)+27)))
	v503 = int64(*(*int32)(unsafe.Add(mBase, uint32(v493)+8)))
	v504 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v493)+12)))
	v505 = int64(*(*int8)(unsafe.Add(mBase, uint32(v493)+26)))
	v506 = int64(*(*int32)(unsafe.Add(mBase, uint32(v493)+4)))
	v507 = F_wangHash64(m, v506)
	mBase = m.M
	v509 = F_wangHash64(m, v505+v507)
	mBase = m.M
	v511 = F_wangHash64(m, v504+v509)
	mBase = m.M
	v513 = F_wangHash64(m, v503+v511)
	mBase = m.M
	v515 = F_wangHash64(m, v502+v513)
	mBase = m.M
	v517 = F_wangHash64(m, v501+v515)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v345)+24)) = v517
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v520 = v519
	goto L161
L163:
	;
	v497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v493)+24)))
	v499 = v497 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v493)+24)) = uint16(v499)
	v520 = v493
	goto L161
L164:
	;
	v527 = v521 + int32(-1)
	goto L158
L165:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	v527 = v524
	goto L158
L166:
	;
	v554 = int32(2)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v534+v552<<(uint(v554)%32)+int32(4))))
	v483 = v559 + v553<<(uint(v554)%32)
	v484 = int32(1)
	goto L152
L167:
	;
	v543 = v531
	goto L169
L168:
	;
	v543 = v528 << (uint(v539) % 32)
	goto L169
L169:
	;
	if v529 < v543 {
		v552 = v535
		v553 = v529
		goto L166
	} else {
		goto L170
	}
L170:
	;
	if v535 != 0 {
		v572 = v531
		goto L155
	} else {
		goto L171
	}
L171:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v534)+20))
	if v545 == int32(-1) {
		v572 = v531
		goto L155
	} else {
		goto L172
	}
L172:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v345)+4)) = int64(4294967296)
	v552 = int32(1)
	v553 = int32(0)
	goto L166
L173:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v564)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v476))) = v568
	v572 = v564
	goto L155
L174:
	;
	goto L145
L175:
	;
	F_addReplyLongLong(m, l0, v582)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L18
	} else {
		goto L176
	}
L176:
	;
	goto L1
L177:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v652)+4))
	v654 = F_objectGetVal(m, v653)
	mBase = m.M
	v655 = int32(_a621)
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654))))
	if v658 != 0 {
		goto L201
	} else {
		goto L202
	}
L178:
	;
	if v627-v629 != 0 {
		goto L177
	} else {
		goto L190
	}
L179:
	;
	v627 = F_tolower(m, v623)
	mBase = m.M
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624))))
	v629 = F_tolower(m, v628)
	mBase = m.M
	goto L178
L180:
	;
	v597 = v591
	v598 = v592
	v599 = v595
	goto L183
L181:
	;
	v623 = int32(0)
	v624 = v592
	goto L179
L182:
	;
	v623 = v620 & int32(255)
	v624 = v619
	goto L179
L183:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598))))
	if v601 == int32(0) {
		v619 = v598
		v620 = v599
		goto L182
	} else {
		goto L185
	}
L184:
	;
	v619 = v613
	v620 = int32(0)
	goto L182
L185:
	;
	v605 = v599 & int32(255)
	if v605 == v601 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v612 = int32(1)
	v613 = v598 + v612
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597)+1)))
	if v614 != 0 {
		v597 = v597 + v612
		v598 = v613
		v599 = v614
		goto L183
	} else {
		goto L189
	}
L187:
	;
	v607 = F_tolower(m, v605)
	mBase = m.M
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598))))
	v609 = F_tolower(m, v608)
	mBase = m.M
	if v607 == v609 {
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597))))
	v619 = v598
	v620 = v611
	goto L182
L189:
	;
	goto L184
L190:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v631 < int32(2) {
		goto L177
	} else {
		goto L191
	}
L191:
	;
	if v631 != int32(2) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	F_latencySpecificCommandsFillCDF(m, l0)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L18
	} else {
		goto L197
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(0)
	v638 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L18
	} else {
		goto L194
	}
L194:
	;
	v641 = *(*int32)(unsafe.Add(mBase, _consts[318]))
	F_latencyAllCommandsFillCDF(m, l0, v641, v9+int32(16))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L18
	} else {
		goto L195
	}
L195:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	F_setDeferredMapLen(m, l0, v638, v646)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L18
	} else {
		goto L196
	}
L196:
	;
	goto L1
L197:
	;
	goto L1
L198:
	;
	F_addReplySubcommandSyntaxError(m, l0)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L18
	} else {
		goto L214
	}
L199:
	;
	if v690-v692 != 0 {
		goto L198
	} else {
		goto L211
	}
L200:
	;
	v690 = F_tolower(m, v686)
	mBase = m.M
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687))))
	v692 = F_tolower(m, v691)
	mBase = m.M
	goto L199
L201:
	;
	v660 = v654
	v661 = v655
	v662 = v658
	goto L204
L202:
	;
	v686 = int32(0)
	v687 = v655
	goto L200
L203:
	;
	v686 = v683 & int32(255)
	v687 = v682
	goto L200
L204:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661))))
	if v664 == int32(0) {
		v682 = v661
		v683 = v662
		goto L203
	} else {
		goto L206
	}
L205:
	;
	v682 = v676
	v683 = int32(0)
	goto L203
L206:
	;
	v668 = v662 & int32(255)
	if v668 == v664 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v675 = int32(1)
	v676 = v661 + v675
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+1)))
	if v677 != 0 {
		v660 = v660 + v675
		v661 = v676
		v662 = v677
		goto L204
	} else {
		goto L210
	}
L208:
	;
	v670 = F_tolower(m, v668)
	mBase = m.M
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661))))
	v672 = F_tolower(m, v671)
	mBase = m.M
	if v670 == v672 {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660))))
	v682 = v661
	v683 = v674
	goto L203
L210:
	;
	goto L205
L211:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v694 != int32(2) {
		goto L198
	} else {
		goto L212
	}
L212:
	;
	v699 = int32(0)
	v700 = *(*int32)(unsafe.Add(mBase, _consts[495]))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(72)))) = v700
	v705 = *(*int64)(unsafe.Add(mBase, _consts[496]))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(64)))) = v705
	v710 = *(*int64)(unsafe.Add(mBase, _consts[497]))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(56)))) = v710
	v715 = *(*int64)(unsafe.Add(mBase, _consts[498]))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v715
	v720 = *(*int64)(unsafe.Add(mBase, _consts[499]))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(40)))) = v720
	v725 = *(*int64)(unsafe.Add(mBase, _consts[500]))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(32)))) = v725
	v728 = *(*int64)(unsafe.Add(mBase, _consts[501]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v728
	v731 = *(*int64)(unsafe.Add(mBase, _consts[502]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v731
	F_addReplyHelp(m, l0, v9+int32(16))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L18
	} else {
		goto L213
	}
L213:
	;
	goto L1
L214:
	;
	goto L1
L215:
	;
	goto L1
}
func F_latencyMonitorInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = F_dictCreate(m, int32(_a877))
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[494])) = v3
		return
	}
}
