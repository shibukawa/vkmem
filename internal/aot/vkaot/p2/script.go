package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_scriptClusterSlotStatsInvalidateSlotIfApplicable(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_scriptClusterSlotStatsInvalidateSlotIfApplicable[0]))
	if v3 == v1 {
		F__serverAssert(m, int32(_a_F_scriptClusterSlotStatsInvalidateSlotIfApplicable_0), int32(_a_F_scriptClusterSlotStatsInvalidateSlotIfApplicable_1), int32(341))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+17)))
		if v6&int32(1) == int32(0) {
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+292)) = int32(-1)
		}
		return
	}
}
func F_scriptCommand(m *base.Module, l0 int32) {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
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
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
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
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 != int32(2) {
		v67 = v9
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(96)
	return
L2:
	;
	if v67 < int32(2) {
		v415 = v67
		goto L25
	} else {
		goto L26
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = F_objectGetVal(m, v13)
	mBase = m.M
	v15 = int32(_a_F_scriptCommand_0)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v67 = v66
	goto L2
L5:
	;
	if v50-v52 != 0 {
		goto L4
	} else {
		goto L17
	}
L6:
	;
	v50 = F_tolower(m, v46)
	mBase = m.M
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v52 = F_tolower(m, v51)
	mBase = m.M
	goto L5
L7:
	;
	v20 = v14
	v21 = v15
	v22 = v18
	goto L10
L8:
	;
	v46 = int32(0)
	v47 = v15
	goto L6
L9:
	;
	v46 = v43 & int32(255)
	v47 = v42
	goto L6
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v24 == int32(0) {
		v42 = v21
		v43 = v22
		goto L9
	} else {
		goto L12
	}
L11:
	;
	v42 = v36
	v43 = int32(0)
	goto L9
L12:
	;
	v28 = v22 & int32(255)
	if v28 == v24 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v35 = int32(1)
	v36 = v21 + v35
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v37 != 0 {
		v20 = v20 + v35
		v21 = v36
		v22 = v37
		goto L10
	} else {
		goto L16
	}
L14:
	;
	v30 = F_tolower(m, v28)
	mBase = m.M
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v32 = F_tolower(m, v31)
	mBase = m.M
	if v30 == v32 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v42 = v21
	v43 = v34
	goto L9
L16:
	;
	goto L11
L17:
	;
	goto L20
L18:
	;
	F_addReplyHelp(m, l0, v7+int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L18
L20:
	;
	v60 = F__emscripten_memcpy_bulkmem(m, v7+int32(16), int32(_a_F_scriptCommand_1), int32(72))
	mBase = m.M
	goto L19
L21:
	;
	return
L22:
	;
	goto L1
L23:
	;
	F_scriptKill(m, l0, int32(1))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L21
	} else {
		goto L247
	}
L24:
	;
	F_addReplyError(m, l0, int32(_a_F_scriptCommand_2))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L21
	} else {
		goto L246
	}
L25:
	;
	if base.Ui32(int32(1)) < base.Ui32(v415+int32(-3)) {
		goto L138
	} else {
		goto L139
	}
L26:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v72 = F_objectGetVal(m, v71)
	mBase = m.M
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v74 = int32(_a_F_scriptCommand_3)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v77 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	if v73 < int32(2) {
		v415 = v73
		goto L25
	} else {
		goto L76
	}
L28:
	;
	if v109-v111 != 0 {
		goto L27
	} else {
		goto L40
	}
L29:
	;
	v109 = F_tolower(m, v105)
	mBase = m.M
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v111 = F_tolower(m, v110)
	mBase = m.M
	goto L28
L30:
	;
	v79 = v72
	v80 = v74
	v81 = v77
	goto L33
L31:
	;
	v105 = int32(0)
	v106 = v74
	goto L29
L32:
	;
	v105 = v102 & int32(255)
	v106 = v101
	goto L29
L33:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v83 == int32(0) {
		v101 = v80
		v102 = v81
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v101 = v95
	v102 = int32(0)
	goto L32
L35:
	;
	v87 = v81 & int32(255)
	if v87 == v83 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v94 = int32(1)
	v95 = v80 + v94
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	if v96 != 0 {
		v79 = v79 + v94
		v80 = v95
		v81 = v96
		goto L33
	} else {
		goto L39
	}
L37:
	;
	v89 = F_tolower(m, v87)
	mBase = m.M
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v91 = F_tolower(m, v90)
	mBase = m.M
	if v89 == v91 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v101 = v80
	v102 = v93
	goto L32
L39:
	;
	goto L34
L40:
	;
	if v73 != int32(3) {
		v205 = v73
		goto L42
	} else {
		goto L43
	}
L41:
	;
	F_evalReset(m, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L21
	} else {
		goto L74
	}
L42:
	;
	if v205 != int32(2) {
		goto L24
	} else {
		goto L73
	}
L43:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v117 = F_objectGetVal(m, v116)
	mBase = m.M
	v118 = int32(_a_F_scriptCommand_4)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v121 != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v158 != int32(3) {
		v205 = v158
		goto L42
	} else {
		goto L58
	}
L45:
	;
	if v153-v155 != 0 {
		goto L44
	} else {
		goto L57
	}
L46:
	;
	v153 = F_tolower(m, v149)
	mBase = m.M
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v155 = F_tolower(m, v154)
	mBase = m.M
	goto L45
L47:
	;
	v123 = v117
	v124 = v118
	v125 = v121
	goto L50
L48:
	;
	v149 = int32(0)
	v150 = v118
	goto L46
L49:
	;
	v149 = v146 & int32(255)
	v150 = v145
	goto L46
L50:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v127 == int32(0) {
		v145 = v124
		v146 = v125
		goto L49
	} else {
		goto L52
	}
L51:
	;
	v145 = v139
	v146 = int32(0)
	goto L49
L52:
	;
	v131 = v125 & int32(255)
	if v131 == v127 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v138 = int32(1)
	v139 = v124 + v138
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
	if v140 != 0 {
		v123 = v123 + v138
		v124 = v139
		v125 = v140
		goto L50
	} else {
		goto L56
	}
L54:
	;
	v133 = F_tolower(m, v131)
	mBase = m.M
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	v135 = F_tolower(m, v134)
	mBase = m.M
	if v133 == v135 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v145 = v124
	v146 = v137
	goto L49
L56:
	;
	goto L51
L57:
	;
	v212 = int32(0)
	goto L41
L58:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	v163 = F_objectGetVal(m, v162)
	mBase = m.M
	v164 = int32(_a_F_scriptCommand_5)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v167 != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v205 = v204
	goto L42
L60:
	;
	if v199-v201 != 0 {
		goto L59
	} else {
		goto L72
	}
L61:
	;
	v199 = F_tolower(m, v195)
	mBase = m.M
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v201 = F_tolower(m, v200)
	mBase = m.M
	goto L60
L62:
	;
	v169 = v163
	v170 = v164
	v171 = v167
	goto L65
L63:
	;
	v195 = int32(0)
	v196 = v164
	goto L61
L64:
	;
	v195 = v192 & int32(255)
	v196 = v191
	goto L61
L65:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v173 == int32(0) {
		v191 = v170
		v192 = v171
		goto L64
	} else {
		goto L67
	}
L66:
	;
	v191 = v185
	v192 = int32(0)
	goto L64
L67:
	;
	v177 = v171 & int32(255)
	if v177 == v173 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v184 = int32(1)
	v185 = v170 + v184
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	if v186 != 0 {
		v169 = v169 + v184
		v170 = v185
		v171 = v186
		goto L65
	} else {
		goto L71
	}
L69:
	;
	v179 = F_tolower(m, v177)
	mBase = m.M
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	v181 = F_tolower(m, v180)
	mBase = m.M
	if v179 == v181 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	v191 = v170
	v192 = v183
	goto L64
L71:
	;
	goto L66
L72:
	;
	v212 = int32(1)
	goto L41
L73:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_scriptCommand[0]))
	v212 = base.B2i32(v209 != int32(0))
	goto L41
L74:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_scriptCommand[1]))
	F_addReply(m, l0, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L21
	} else {
		goto L75
	}
L75:
	;
	goto L1
L76:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	v223 = F_objectGetVal(m, v222)
	mBase = m.M
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v225 = int32(_a_F_scriptCommand_6)
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	if v228 != 0 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	if v224 != int32(3) {
		v367 = v224
		goto L101
	} else {
		goto L102
	}
L78:
	;
	if v260-v262 != 0 {
		goto L77
	} else {
		goto L90
	}
L79:
	;
	v260 = F_tolower(m, v256)
	mBase = m.M
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
	v262 = F_tolower(m, v261)
	mBase = m.M
	goto L78
L80:
	;
	v230 = v223
	v231 = v225
	v232 = v228
	goto L83
L81:
	;
	v256 = int32(0)
	v257 = v225
	goto L79
L82:
	;
	v256 = v253 & int32(255)
	v257 = v252
	goto L79
L83:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v234 == int32(0) {
		v252 = v231
		v253 = v232
		goto L82
	} else {
		goto L85
	}
L84:
	;
	v252 = v246
	v253 = int32(0)
	goto L82
L85:
	;
	v238 = v232 & int32(255)
	if v238 == v234 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v245 = int32(1)
	v246 = v231 + v245
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if v247 != 0 {
		v230 = v230 + v245
		v231 = v246
		v232 = v247
		goto L83
	} else {
		goto L89
	}
L87:
	;
	v240 = F_tolower(m, v238)
	mBase = m.M
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	v242 = F_tolower(m, v241)
	mBase = m.M
	if v240 == v242 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	v252 = v231
	v253 = v244
	goto L82
L89:
	;
	goto L84
L90:
	;
	F_addReplyArrayLen(m, l0, v224+int32(-2))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L21
	} else {
		goto L91
	}
L91:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v268 < int32(3) {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v274 = int32(2)
	goto L93
L93:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_scriptCommand[2]))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v281+v274<<(uint(int32(2))%32))))
	v286 = F_objectGetVal(m, v285)
	mBase = m.M
	v287 = F_dictFind(m, v280, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L21
	} else {
		goto L95
	}
L95:
	;
	if v287 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v289 = int32(16)
	goto L98
L97:
	;
	v289 = int32(12)
	goto L98
L98:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v289)+uint32(_c_F_scriptCommand[1])))
	F_addReply(m, l0, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L21
	} else {
		goto L99
	}
L99:
	;
	v295 = v274 + int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v295 < v296 {
		v274 = v295
		goto L93
	} else {
		goto L100
	}
L100:
	;
	goto L1
L101:
	;
	if v367 != int32(2) {
		v415 = v367
		goto L25
	} else {
		goto L124
	}
L102:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	v302 = F_objectGetVal(m, v301)
	mBase = m.M
	v303 = int32(_a_F_scriptCommand_7)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	if v306 != 0 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v367 = v366
	goto L101
L104:
	;
	if v338-v340 != 0 {
		goto L103
	} else {
		goto L116
	}
L105:
	;
	v338 = F_tolower(m, v334)
	mBase = m.M
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	v340 = F_tolower(m, v339)
	mBase = m.M
	goto L104
L106:
	;
	v308 = v302
	v309 = v303
	v310 = v306
	goto L109
L107:
	;
	v334 = int32(0)
	v335 = v303
	goto L105
L108:
	;
	v334 = v331 & int32(255)
	v335 = v330
	goto L105
L109:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if v312 == int32(0) {
		v330 = v309
		v331 = v310
		goto L108
	} else {
		goto L111
	}
L110:
	;
	v330 = v324
	v331 = int32(0)
	goto L108
L111:
	;
	v316 = v310 & int32(255)
	if v316 == v312 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v323 = int32(1)
	v324 = v309 + v323
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+1)))
	if v325 != 0 {
		v308 = v308 + v323
		v309 = v324
		v310 = v325
		goto L109
	} else {
		goto L115
	}
L113:
	;
	v318 = F_tolower(m, v316)
	mBase = m.M
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	v320 = F_tolower(m, v319)
	mBase = m.M
	if v318 == v320 {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	v330 = v309
	v331 = v322
	goto L108
L115:
	;
	goto L110
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(0)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+8))
	v348 = F_evalRegisterNewScript(m, l0, v345, v7+int32(16))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L21
	} else {
		goto L117
	}
L117:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v348 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_addReplyBulkCBuffer(m, l0, v350, int32(40))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L21
	} else {
		goto L122
	}
L119:
	;
	if v350 == int32(0) {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F__serverAssert(m, int32(_a_F_scriptCommand_8), int32(_a_F_scriptCommand_9), int32(629))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L21
	} else {
		goto L121
	}
L121:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	F_valkey_free(m, v350)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L21
	} else {
		goto L123
	}
L123:
	;
	goto L1
L124:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	v372 = F_objectGetVal(m, v371)
	mBase = m.M
	v373 = int32(_a_F_scriptCommand_10)
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	if v376 != 0 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	if v408-v410 == int32(0) {
		goto L23
	} else {
		goto L137
	}
L126:
	;
	v408 = F_tolower(m, v404)
	mBase = m.M
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405))))
	v410 = F_tolower(m, v409)
	mBase = m.M
	goto L125
L127:
	;
	v378 = v372
	v379 = v373
	v380 = v376
	goto L130
L128:
	;
	v404 = int32(0)
	v405 = v373
	goto L126
L129:
	;
	v404 = v401 & int32(255)
	v405 = v400
	goto L126
L130:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	if v382 == int32(0) {
		v400 = v379
		v401 = v380
		goto L129
	} else {
		goto L132
	}
L131:
	;
	v400 = v394
	v401 = int32(0)
	goto L129
L132:
	;
	v386 = v380 & int32(255)
	if v386 == v382 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v393 = int32(1)
	v394 = v379 + v393
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+1)))
	if v395 != 0 {
		v378 = v378 + v393
		v379 = v394
		v380 = v395
		goto L130
	} else {
		goto L136
	}
L134:
	;
	v388 = F_tolower(m, v386)
	mBase = m.M
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	v390 = F_tolower(m, v389)
	mBase = m.M
	if v388 == v390 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	v400 = v379
	v401 = v392
	goto L129
L136:
	;
	goto L131
L137:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v415 = v414
	goto L25
L138:
	;
	F_addReplySubcommandSyntaxError(m, l0)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L21
	} else {
		goto L245
	}
L139:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	v423 = F_objectGetVal(m, v422)
	mBase = m.M
	v424 = int32(_a_F_scriptCommand_11)
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423))))
	if v427 != 0 {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v649 != int32(3) {
		goto L138
	} else {
		goto L219
	}
L141:
	;
	if v459-v461 != 0 {
		goto L140
	} else {
		goto L153
	}
L142:
	;
	v459 = F_tolower(m, v455)
	mBase = m.M
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	v461 = F_tolower(m, v460)
	mBase = m.M
	goto L141
L143:
	;
	v429 = v423
	v430 = v424
	v431 = v427
	goto L146
L144:
	;
	v455 = int32(0)
	v456 = v424
	goto L142
L145:
	;
	v455 = v452 & int32(255)
	v456 = v451
	goto L142
L146:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	if v433 == int32(0) {
		v451 = v430
		v452 = v431
		goto L145
	} else {
		goto L148
	}
L147:
	;
	v451 = v445
	v452 = int32(0)
	goto L145
L148:
	;
	v437 = v431 & int32(255)
	if v437 == v433 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v444 = int32(1)
	v445 = v430 + v444
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+1)))
	if v446 != 0 {
		v429 = v429 + v444
		v430 = v445
		v431 = v446
		goto L146
	} else {
		goto L152
	}
L150:
	;
	v439 = F_tolower(m, v437)
	mBase = m.M
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	v441 = F_tolower(m, v440)
	mBase = m.M
	if v439 == v441 {
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429))))
	v451 = v430
	v452 = v443
	goto L145
L152:
	;
	goto L147
L153:
	;
	v463 = F_clientHasPendingReplies(m, l0)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L21
	} else {
		goto L155
	}
L154:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v471 != int32(4) {
		v477 = int32(_a_F_scriptCommand_12)
		goto L158
	} else {
		goto L159
	}
L155:
	;
	if v463 == int32(0) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	F_addReplyError(m, l0, int32(_a_F_scriptCommand_13))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L21
	} else {
		goto L157
	}
L157:
	;
	goto L1
L158:
	;
	v478 = F_scriptingEngineManagerFind(m, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L21
	} else {
		goto L161
	}
L159:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v474)+12))
	v476 = F_objectGetVal(m, v475)
	mBase = m.M
	v477 = v476
	goto L158
L160:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+8))
	v486 = F_objectGetVal(m, v485)
	mBase = m.M
	v487 = int32(_a_F_scriptCommand_14)
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486))))
	if v490 != 0 {
		goto L167
	} else {
		goto L168
	}
L161:
	;
	if v478 != 0 {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v477
	F_addReplyErrorFormat(m, l0, int32(_a_F_scriptCommand_15), v7)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L21
	} else {
		goto L163
	}
L163:
	;
	goto L1
L164:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)+8))
	v534 = F_objectGetVal(m, v533)
	mBase = m.M
	v535 = int32(_a_F_scriptCommand_16)
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	if v538 != 0 {
		goto L183
	} else {
		goto L184
	}
L165:
	;
	if v522-v524 != 0 {
		goto L164
	} else {
		goto L177
	}
L166:
	;
	v522 = F_tolower(m, v518)
	mBase = m.M
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
	v524 = F_tolower(m, v523)
	mBase = m.M
	goto L165
L167:
	;
	v492 = v486
	v493 = v487
	v494 = v490
	goto L170
L168:
	;
	v518 = int32(0)
	v519 = v487
	goto L166
L169:
	;
	v518 = v515 & int32(255)
	v519 = v514
	goto L166
L170:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
	if v496 == int32(0) {
		v514 = v493
		v515 = v494
		goto L169
	} else {
		goto L172
	}
L171:
	;
	v514 = v508
	v515 = int32(0)
	goto L169
L172:
	;
	v500 = v494 & int32(255)
	if v500 == v496 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v507 = int32(1)
	v508 = v493 + v507
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+1)))
	if v509 != 0 {
		v492 = v492 + v507
		v493 = v508
		v494 = v509
		goto L170
	} else {
		goto L176
	}
L174:
	;
	v502 = F_tolower(m, v500)
	mBase = m.M
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
	v504 = F_tolower(m, v503)
	mBase = m.M
	if v502 == v504 {
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492))))
	v514 = v493
	v515 = v506
	goto L169
L176:
	;
	goto L171
L177:
	;
	F_scriptingEngineDebuggerDisable(m, l0)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L21
	} else {
		goto L178
	}
L178:
	;
	v529 = *(*int32)(unsafe.Add(mBase, _c_F_scriptCommand[1]))
	F_addReply(m, l0, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L21
	} else {
		goto L179
	}
L179:
	;
	goto L1
L180:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)+8))
	v589 = F_objectGetVal(m, v588)
	mBase = m.M
	v590 = int32(_a_F_scriptCommand_4)
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	if v593 != 0 {
		goto L202
	} else {
		goto L203
	}
L181:
	;
	if v570-v572 != 0 {
		goto L180
	} else {
		goto L193
	}
L182:
	;
	v570 = F_tolower(m, v566)
	mBase = m.M
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567))))
	v572 = F_tolower(m, v571)
	mBase = m.M
	goto L181
L183:
	;
	v540 = v534
	v541 = v535
	v542 = v538
	goto L186
L184:
	;
	v566 = int32(0)
	v567 = v535
	goto L182
L185:
	;
	v566 = v563 & int32(255)
	v567 = v562
	goto L182
L186:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	if v544 == int32(0) {
		v562 = v541
		v563 = v542
		goto L185
	} else {
		goto L188
	}
L187:
	;
	v562 = v556
	v563 = int32(0)
	goto L185
L188:
	;
	v548 = v542 & int32(255)
	if v548 == v544 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v555 = int32(1)
	v556 = v541 + v555
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+1)))
	if v557 != 0 {
		v540 = v540 + v555
		v541 = v556
		v542 = v557
		goto L186
	} else {
		goto L192
	}
L190:
	;
	v550 = F_tolower(m, v548)
	mBase = m.M
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	v552 = F_tolower(m, v551)
	mBase = m.M
	if v550 == v552 {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540))))
	v562 = v541
	v563 = v554
	goto L185
L192:
	;
	goto L187
L193:
	;
	v576 = F_scriptingEngineDebuggerEnable(m, l0, v478, v7+int32(16))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L21
	} else {
		goto L195
	}
L194:
	;
	v584 = *(*int32)(unsafe.Add(mBase, _c_F_scriptCommand[1]))
	F_addReply(m, l0, v584)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L21
	} else {
		goto L198
	}
L195:
	;
	if v576 == int32(0) {
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	F_addReplyErrorSds(m, l0, v580)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L21
	} else {
		goto L197
	}
L197:
	;
	goto L1
L198:
	;
	goto L1
L199:
	;
	F_addReplyError(m, l0, int32(_a_F_scriptCommand_17))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L21
	} else {
		goto L218
	}
L200:
	;
	if v625-v627 != 0 {
		goto L199
	} else {
		goto L212
	}
L201:
	;
	v625 = F_tolower(m, v621)
	mBase = m.M
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622))))
	v627 = F_tolower(m, v626)
	mBase = m.M
	goto L200
L202:
	;
	v595 = v589
	v596 = v590
	v597 = v593
	goto L205
L203:
	;
	v621 = int32(0)
	v622 = v590
	goto L201
L204:
	;
	v621 = v618 & int32(255)
	v622 = v617
	goto L201
L205:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596))))
	if v599 == int32(0) {
		v617 = v596
		v618 = v597
		goto L204
	} else {
		goto L207
	}
L206:
	;
	v617 = v611
	v618 = int32(0)
	goto L204
L207:
	;
	v603 = v597 & int32(255)
	if v603 == v599 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v610 = int32(1)
	v611 = v596 + v610
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595)+1)))
	if v612 != 0 {
		v595 = v595 + v610
		v596 = v611
		v597 = v612
		goto L205
	} else {
		goto L211
	}
L209:
	;
	v605 = F_tolower(m, v603)
	mBase = m.M
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596))))
	v607 = F_tolower(m, v606)
	mBase = m.M
	if v605 == v607 {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595))))
	v617 = v596
	v618 = v609
	goto L204
L211:
	;
	goto L206
L212:
	;
	v631 = F_scriptingEngineDebuggerEnable(m, l0, v478, v7+int32(16))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L21
	} else {
		goto L214
	}
L213:
	;
	v639 = *(*int32)(unsafe.Add(mBase, _c_F_scriptCommand[1]))
	F_addReply(m, l0, v639)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L21
	} else {
		goto L217
	}
L214:
	;
	if v631 == int32(0) {
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	F_addReplyErrorSds(m, l0, v635)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L21
	} else {
		goto L216
	}
L216:
	;
	goto L1
L217:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v642 | int32(536870912)
	goto L1
L218:
	;
	goto L1
L219:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v652)+4))
	v654 = F_objectGetVal(m, v653)
	mBase = m.M
	v655 = int32(_a_F_scriptCommand_18)
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654))))
	if v658 != 0 {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	if v690-v692 != 0 {
		goto L138
	} else {
		goto L232
	}
L221:
	;
	v690 = F_tolower(m, v686)
	mBase = m.M
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687))))
	v692 = F_tolower(m, v691)
	mBase = m.M
	goto L220
L222:
	;
	v660 = v654
	v661 = v655
	v662 = v658
	goto L225
L223:
	;
	v686 = int32(0)
	v687 = v655
	goto L221
L224:
	;
	v686 = v683 & int32(255)
	v687 = v682
	goto L221
L225:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661))))
	if v664 == int32(0) {
		v682 = v661
		v683 = v662
		goto L224
	} else {
		goto L227
	}
L226:
	;
	v682 = v676
	v683 = int32(0)
	goto L224
L227:
	;
	v668 = v662 & int32(255)
	if v668 == v664 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v675 = int32(1)
	v676 = v661 + v675
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+1)))
	if v677 != 0 {
		v660 = v660 + v675
		v661 = v676
		v662 = v677
		goto L225
	} else {
		goto L231
	}
L229:
	;
	v670 = F_tolower(m, v668)
	mBase = m.M
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661))))
	v672 = F_tolower(m, v671)
	mBase = m.M
	if v670 == v672 {
		goto L228
	} else {
		goto L230
	}
L230:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660))))
	v682 = v661
	v683 = v674
	goto L224
L231:
	;
	goto L226
L232:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)+8))
	v696 = F_objectGetVal(m, v695)
	mBase = m.M
	v697 = int32(-1)
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696+v697))))
	switch v699&int32(7) + v697 {
	case 0:
		goto L238
	case 1:
		goto L237
	case 2:
		goto L236
	case 3:
		goto L235
	default:
		goto L233
	}
L233:
	;
	v734 = *(*int32)(unsafe.Add(mBase, _c_F_scriptCommand[3]))
	F_addReplyErrorObject(m, l0, v734)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L21
	} else {
		goto L244
	}
L234:
	;
	if v716 != int32(40) {
		goto L233
	} else {
		goto L239
	}
L235:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v696+int32(-17))))
	v716 = v715
	goto L234
L236:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v696+int32(-9))))
	v716 = v712
	goto L234
L237:
	;
	v709 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v696+int32(-5)))))
	v716 = v709
	goto L234
L238:
	;
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696+int32(-3)))))
	v716 = v706
	goto L234
L239:
	;
	v720 = *(*int32)(unsafe.Add(mBase, _c_F_scriptCommand[2]))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v721)+8))
	v723 = F_objectGetVal(m, v722)
	mBase = m.M
	v724 = F_dictFind(m, v720, v723)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L21
	} else {
		goto L240
	}
L240:
	;
	if v724 == int32(0) {
		goto L233
	} else {
		goto L241
	}
L241:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v724)+8))
	goto L242
L242:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v728)+8))
	F_addReplyBulk(m, l0, v729)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L21
	} else {
		goto L243
	}
L243:
	;
	goto L1
L244:
	;
	goto L1
L245:
	;
	goto L1
L246:
	;
	goto L1
L247:
	;
	goto L1
}
func F_scriptCurrFunction(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_scriptCurrFunction[0]))
	if v3 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		return v12
	} else {
		F__serverAssert(m, int32(_a_F_scriptCurrFunction_0), int32(_a_F_scriptCurrFunction_1), int32(274))
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
func F_scriptGetCaller(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_scriptGetCaller[0]))
	if v3 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
		return v12
	} else {
		F__serverAssert(m, int32(_a_F_scriptGetCaller_0), int32(_a_F_scriptGetCaller_1), int32(68))
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
func F_scriptGetRunningEngineName(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_scriptGetRunningEngineName[0]))
	if v3 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		return v13
	} else {
		F__serverAssert(m, int32(_a_F_scriptGetRunningEngineName_0), int32(_a_F_scriptGetRunningEngineName_1), int32(361))
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
func F_scriptIsEval(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_scriptIsEval[0]))
	if v3 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
		return v12 & int32(128)
	} else {
		F__serverAssert(m, int32(_a_F_scriptIsEval_0), int32(_a_F_scriptIsEval_1), int32(279))
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
func F_scriptIsWriteDirty(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_scriptIsWriteDirty[0]))
	if v3 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
		return v12 & int32(1)
	} else {
		F__serverAssert(m, int32(_a_F_scriptIsWriteDirty_0), int32(_a_F_scriptIsWriteDirty_1), int32(331))
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
func F_scriptRunDuration(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_scriptRunDuration[0]))
	if v4 != 0 {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v4)+24))
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_scriptRunDuration[1]))
		v16 = m.T0[v15].(func(*base.Module) int64)(m)
		mBase = m.M
		v19 = base.I64_div_u_s(v16-v13, int64(1000))
		return v19
	} else {
		F__serverAssert(m, int32(_a_F_scriptRunDuration_0), int32(_a_F_scriptRunDuration_1), int32(316))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
