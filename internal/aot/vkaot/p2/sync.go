package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_sendSyncSlotsMessage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v5 == int32(0) {
		F__serverAssert(m, int32(_a447), int32(_a443), int32(2471))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+204))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v9 != int32(1) {
			v16 = v5
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+204)) = v8 | int32(131072)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
			v16 = v15
		}
		F_addReplyArrayLen(m, v16, int32(3))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
			F_addReplyBulkCString(m, v20, int32(_a448))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
				F_addReplyBulkCString(m, v24, int32(_a449))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
					F_addReplyBulkCString(m, v28, l1)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						if v8&int32(131072) != 0 {
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+204))
							*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v34 & int32(-131073)
						}
						return
					}
				}
			}
		}
	}
}
func F_syncCommand(m *base.Module, l0 int32) {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
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
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
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
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
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
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int64
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int64
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
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
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int64
	_ = v390
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int64
	_ = v431
	var v434 int32
	_ = v434
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int64
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v661 int32
	_ = v661
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v11&int32(2) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(96)
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v16 = F_valkey_calloc(m, int32(200))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v16
	goto L3
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v39 < int32(4) {
		goto L25
	} else {
		goto L26
	}
L8:
	;
	goto L7
L9:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v23 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v20 == int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v31 != int32(1) {
		goto L8
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v27 == int32(1) {
		goto L14
	} else {
		goto L16
	}
L15:
	;
	goto L12
L16:
	;
	goto L15
L17:
	;
	goto L18
L18:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v35 == int32(1) {
		goto L18
	} else {
		goto L20
	}
L19:
	;
	goto L8
L20:
	;
	goto L19
L21:
	;
	F__serverLog(m, int32(2), v652, int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L5
	} else {
		goto L200
	}
L22:
	;
	v648 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v648 {
		goto L1
	} else {
		goto L199
	}
L23:
	;
	v639 = F_startBgsaveForReplication(m, v609, v607, int32(11))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L5
	} else {
		goto L198
	}
L24:
	;
	F_addReplyError(m, l0, int32(_a1207))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L5
	} else {
		goto L197
	}
L25:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[658]))
	if v218 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L26:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v44 = F_objectGetVal(m, v43)
	mBase = m.M
	v45 = int32(_a1208)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v48 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	if v80-v82 != 0 {
		goto L25
	} else {
		goto L39
	}
L28:
	;
	v80 = F_tolower(m, v76)
	mBase = m.M
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v82 = F_tolower(m, v81)
	mBase = m.M
	goto L27
L29:
	;
	v50 = v44
	v51 = v45
	v52 = v48
	goto L32
L30:
	;
	v76 = int32(0)
	v77 = v45
	goto L28
L31:
	;
	v76 = v73 & int32(255)
	v77 = v72
	goto L28
L32:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v54 == int32(0) {
		v72 = v51
		v73 = v52
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v72 = v66
	v73 = int32(0)
	goto L31
L34:
	;
	v58 = v52 & int32(255)
	if v58 == v54 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v65 = int32(1)
	v66 = v51 + v65
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v67 != 0 {
		v50 = v50 + v65
		v51 = v66
		v52 = v67
		goto L32
	} else {
		goto L38
	}
L36:
	;
	v60 = F_tolower(m, v58)
	mBase = m.M
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v62 = F_tolower(m, v61)
	mBase = m.M
	if v60 == v62 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v72 = v51
	v73 = v64
	goto L31
L38:
	;
	goto L33
L39:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v86 = F_objectGetVal(m, v85)
	mBase = m.M
	v87 = int32(_a407)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v90 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if v122-v124 != 0 {
		goto L25
	} else {
		goto L52
	}
L41:
	;
	v122 = F_tolower(m, v118)
	mBase = m.M
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v124 = F_tolower(m, v123)
	mBase = m.M
	goto L40
L42:
	;
	v92 = v86
	v93 = v87
	v94 = v90
	goto L45
L43:
	;
	v118 = int32(0)
	v119 = v87
	goto L41
L44:
	;
	v118 = v115 & int32(255)
	v119 = v114
	goto L41
L45:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v96 == int32(0) {
		v114 = v93
		v115 = v94
		goto L44
	} else {
		goto L47
	}
L46:
	;
	v114 = v108
	v115 = int32(0)
	goto L44
L47:
	;
	v100 = v94 & int32(255)
	if v100 == v96 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v107 = int32(1)
	v108 = v93 + v107
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	if v109 != 0 {
		v92 = v92 + v107
		v93 = v108
		v94 = v109
		goto L45
	} else {
		goto L51
	}
L49:
	;
	v102 = F_tolower(m, v100)
	mBase = m.M
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	v104 = F_tolower(m, v103)
	mBase = m.M
	if v102 == v104 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v114 = v93
	v115 = v106
	goto L44
L51:
	;
	goto L46
L52:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v127 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v141 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v132 = F_objectGetVal(m, v131)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v132
	F__serverLog(m, int32(2), int32(_a1209), v9+int32(80))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	v147 = F_objectGetVal(m, v146)
	mBase = m.M
	v150 = int32(_a946)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v153 != 0 {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	F_addReplyError(m, l0, int32(_a1210))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	goto L1
L59:
	;
	if v185-v187 != 0 {
		goto L24
	} else {
		goto L71
	}
L60:
	;
	v185 = F_tolower(m, v181)
	mBase = m.M
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v187 = F_tolower(m, v186)
	mBase = m.M
	goto L59
L61:
	;
	v155 = v147
	v156 = v150
	v157 = v153
	goto L64
L62:
	;
	v181 = int32(0)
	v182 = v150
	goto L60
L63:
	;
	v181 = v178 & int32(255)
	v182 = v177
	goto L60
L64:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v159 == int32(0) {
		v177 = v156
		v178 = v157
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v177 = v171
	v178 = int32(0)
	goto L63
L66:
	;
	v163 = v157 & int32(255)
	if v163 == v159 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v170 = int32(1)
	v171 = v156 + v170
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
	if v172 != 0 {
		v155 = v155 + v170
		v156 = v171
		v157 = v172
		goto L64
	} else {
		goto L70
	}
L68:
	;
	v165 = F_tolower(m, v163)
	mBase = m.M
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	v167 = F_tolower(m, v166)
	mBase = m.M
	if v165 == v167 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	v177 = v156
	v178 = v169
	goto L63
L70:
	;
	goto L65
L71:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v190 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v197 = F_sdsempty(m)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L5
	} else {
		goto L77
	}
L73:
	;
	F_replicationUnsetPrimary(m)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L5
	} else {
		goto L76
	}
L74:
	;
	F_clusterPromoteSelfToPrimary(m)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	goto L72
L77:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _consts[209]))
	v201 = F_catClientInfoShortString(m, v197, l0, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v204 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	F_sdsfree(m, v201)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L5
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v201
	F__serverLog(m, int32(2), int32(_a1211), v9+int32(64))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	goto L25
L83:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v225 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	F_addReplyError(m, l0, int32(_a1212))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	goto L1
L86:
	;
	v235 = F_clientHasPendingReplies(m, l0)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L91
	}
L87:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	if v229 == int32(14) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	F_addReplyError(m, l0, int32(_a1213))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	goto L1
L90:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+162)))
	if v243&int32(3) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	if v235 == int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	F_addReplyError(m, l0, int32(_a1214))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L93
	}
L93:
	;
	goto L1
L94:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v255 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+160)))
	if v248&int32(1) != 0 {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	F_addReplyError(m, l0, int32(_a1215))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	goto L1
L98:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v269 = F_objectGetVal(m, v268)
	mBase = m.M
	v270 = int32(_a1208)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	if v273 != 0 {
		goto L106
	} else {
		goto L107
	}
L99:
	;
	v258 = F_replicationGetReplicaName(m, l0)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v258
	F__serverLog(m, int32(2), int32(_a1216), v9+int32(48))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	goto L98
L102:
	;
	v388 = int32(_a44)
	v390 = *(*int64)(unsafe.Add(mBase, _consts[659]))
	*(*int64)(unsafe.Add(mBase, _consts[659])) = v390 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = int32(6)
	v397 = *(*int32)(unsafe.Add(mBase, _consts[660]))
	if v397 == int32(0) {
		v406 = v387
		goto L136
	} else {
		goto L137
	}
L103:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v382 | int32(65536)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v387 = v386
	goto L102
L104:
	;
	if v305-v307 != 0 {
		goto L103
	} else {
		goto L116
	}
L105:
	;
	v305 = F_tolower(m, v301)
	mBase = m.M
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	v307 = F_tolower(m, v306)
	mBase = m.M
	goto L104
L106:
	;
	v275 = v269
	v276 = v270
	v277 = v273
	goto L109
L107:
	;
	v301 = int32(0)
	v302 = v270
	goto L105
L108:
	;
	v301 = v298 & int32(255)
	v302 = v297
	goto L105
L109:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	if v279 == int32(0) {
		v297 = v276
		v298 = v277
		goto L108
	} else {
		goto L111
	}
L110:
	;
	v297 = v291
	v298 = int32(0)
	goto L108
L111:
	;
	v283 = v277 & int32(255)
	if v283 == v279 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v290 = int32(1)
	v291 = v276 + v290
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+1)))
	if v292 != 0 {
		v275 = v275 + v290
		v276 = v291
		v277 = v292
		goto L109
	} else {
		goto L115
	}
L113:
	;
	v285 = F_tolower(m, v283)
	mBase = m.M
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	v287 = F_tolower(m, v286)
	mBase = m.M
	if v285 == v287 {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	v297 = v276
	v298 = v289
	goto L108
L115:
	;
	goto L110
L116:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+8))
	v314 = F_getLongLongFromObjectOrReply(m, l0, v310, v9+int32(88), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L118
	}
L117:
	;
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v9)+88))
	v332 = F_primaryTryPartialResynchronization(m, l0, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L5
	} else {
		goto L124
	}
L118:
	;
	if v314 == int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v319 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v322 = F_replicationGetReplicaName(m, l0)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v322
	F__serverLog(m, int32(3), int32(_a1217), v9+int32(32))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	goto L1
L123:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+4))
	v342 = F_objectGetVal(m, v341)
	mBase = m.M
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	if v343 == int32(63) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	if v332 != 0 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v334 = int32(_a44)
	v336 = *(*int64)(unsafe.Add(mBase, _consts[661]))
	*(*int64)(unsafe.Add(mBase, _consts[661])) = v336 + int64(1)
	goto L1
L126:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+160)))
	if v353&int32(4) == int32(0) {
		v387 = v352
		goto L102
	} else {
		goto L128
	}
L127:
	;
	v346 = int32(_a44)
	v348 = *(*int64)(unsafe.Add(mBase, _consts[662]))
	*(*int64)(unsafe.Add(mBase, _consts[662])) = v348 + int64(1)
	goto L126
L128:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v359 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)+68))
	v376 = m.T0[v375].(func(*base.Module, int32, int32, int32) int32)(m, v371, int32(_a1218), int32(18))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L5
	} else {
		goto L133
	}
L130:
	;
	v362 = F_replicationGetReplicaName(m, l0)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v362
	F__serverLog(m, int32(2), int32(_a1219), v9+int32(16))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	goto L129
L133:
	;
	if v376 == int32(18) {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_freeClientAsync(m, l0)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L5
	} else {
		goto L135
	}
L135:
	;
	goto L1
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v406)+8)) = int32(-1)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v409 | int32(2)
	v414 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v415 = F_listAddNodeTail(m, v414, l0)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L5
	} else {
		goto L139
	}
L137:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+12))
	v403 = F_anetDisableTcpNoDelay(m, int32(0), v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L5
	} else {
		goto L138
	}
L138:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v406 = v405
	goto L136
L139:
	;
	v418 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+20))
	if v419 != int32(1) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v478 = int32(_a44)
	v479 = *(*int32)(unsafe.Add(mBase, _consts[651]))
	v481 = *(*int32)(unsafe.Add(mBase, _consts[66]))
	if v481 != int32(1) {
		goto L147
	} else {
		goto L148
	}
L141:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	if v423 != 0 {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v426 = int32(_a946)
	F_getRandomHexChars(m, v426, int32(40))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L5
	} else {
		goto L143
	}
L143:
	;
	v430 = int32(_a44)
	v431 = int64(3472328296227680304)
	*(*int64)(unsafe.Add(mBase, _consts[384])) = v431
	v434 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[663])) = uint8(v434)
	*(*int64)(unsafe.Add(mBase, _consts[390])) = v431
	*(*int64)(unsafe.Add(mBase, _consts[389])) = v431
	*(*int64)(unsafe.Add(mBase, _consts[388])) = v431
	*(*int64)(unsafe.Add(mBase, _consts[387])) = v431
	*(*int64)(unsafe.Add(mBase, _consts[385])) = int64(-1)
	*(*uint8)(unsafe.Add(mBase, _consts[386])) = uint8(v434)
	F_createReplicationBacklog(m)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L5
	} else {
		goto L144
	}
L144:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v465 {
		goto L140
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(_a945)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v426
	F__serverLog(m, int32(2), int32(_a1220), v9)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L5
	} else {
		goto L146
	}
L146:
	;
	goto L140
L147:
	;
	if v481 != int32(1) {
		goto L173
	} else {
		goto L174
	}
L148:
	;
	if v479 != int32(1) {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v487 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v489 = v9 + int32(88)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	*(*int32)(unsafe.Add(mBase, uint32(v489)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v489))) = v490
	goto L150
L150:
	;
	v495 = v9 + int32(88)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	if v497 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	if v497 == int32(0) {
		goto L22
	} else {
		goto L154
	}
L152:
	;
	goto L151
L153:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v497+base.B2i32(v500 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v506
	goto L152
L154:
	;
	v512 = v497
	goto L156
L155:
	;
	v543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v517)+160)))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v545 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v544)+160)))
	if v543&v545 != v543 {
		goto L22
	} else {
		goto L166
	}
L156:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v512)+8))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)+104))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v517)))
	if v518 != int32(7) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v530 = v9 + int32(88)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	if v532 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L159:
	;
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+205)))
	if v521&int32(32) == int32(0) {
		goto L155
	} else {
		goto L160
	}
L160:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
	if v526&int32(32) != 0 {
		goto L155
	} else {
		goto L161
	}
L161:
	;
	goto L158
L162:
	;
	if v532 != 0 {
		v512 = v532
		goto L156
	} else {
		goto L165
	}
L163:
	;
	goto L162
L164:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v530)+4))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v532+base.B2i32(v535 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v530))) = v541
	goto L163
L165:
	;
	goto L22
L166:
	;
	v548 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v544)+162)))
	v549 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v517)+162)))
	if v548 != v549 {
		goto L22
	} else {
		goto L167
	}
L167:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
	if v551&int32(32) != 0 {
		v557 = v517
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v558 = *(*int64)(unsafe.Add(mBase, uint32(v557)+96))
	v559 = F_replicationSetupReplicaForFullResync(m, l0, v558)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L5
	} else {
		goto L171
	}
L169:
	;
	F_copyReplicaOutputBuffer(m, l0, v516)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L5
	} else {
		goto L170
	}
L170:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v516)+104))
	v557 = v556
	goto L168
L171:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v562 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v652 = int32(_a1221)
	goto L21
L173:
	;
	v580 = *(*int32)(unsafe.Add(mBase, _consts[655]))
	if v580 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L174:
	;
	if v479 != int32(2) {
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v571 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v571 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F__serverLog(m, int32(2), int32(_a1222), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L5
	} else {
		goto L177
	}
L177:
	;
	goto L1
L178:
	;
	v603 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	goto L185
L179:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+160)))
	if v584&int32(1) == int32(0) {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v590 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	if v590 == int32(0) {
		goto L178
	} else {
		goto L181
	}
L181:
	;
	v594 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v594 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	F__serverLog(m, int32(2), int32(_a1223), int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L5
	} else {
		goto L183
	}
L183:
	;
	goto L1
L184:
	;
	v627 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v627 {
		goto L1
	} else {
		goto L195
	}
L185:
	;
	if v603 != int32(-1) {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v607 = int32(*(*int16)(unsafe.Add(mBase, uint32(v606)+162)))
	v608 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v606)+160)))
	v609 = base.I32_extend16_s(v608)
	if v608&int32(1) != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v606)+156))
	if v615 <= int32(589823) {
		goto L191
	} else {
		goto L192
	}
L188:
	;
	v613 = F_startBgsaveForReplication(m, v609, v607, int32(80))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L5
	} else {
		goto L189
	}
L189:
	;
	goto L1
L190:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	v624 = F_startBgsaveForReplication(m, v609, v607, v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L5
	} else {
		goto L194
	}
L191:
	;
	if v615 < int32(459264) {
		goto L23
	} else {
		goto L193
	}
L192:
	;
	v622 = int32(_a1204)
	goto L190
L193:
	;
	v622 = int32(_a1205)
	goto L190
L194:
	;
	goto L1
L195:
	;
	F__serverLog(m, int32(2), int32(_a1224), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L5
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
	goto L1
L199:
	;
	v652 = int32(_a1225)
	goto L21
L200:
	;
	goto L1
}
func F_syncRead(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v102 int32
	_ = v102
	var v117 int32
	_ = v117
	v5 = int32(0)
	v12 = F_mstime(m)
	mBase = m.M
	if l2 == v5 {
		v117 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v117
L2:
	;
	v16 = l1
	v17 = l2
	v19 = v5
	v21 = l3
	goto L5
L3:
	;
	v117 = int32(-1)
	goto L1
L4:
	;
	goto L27
L5:
	;
	v26 = F_read(m, l0, v16, v17)
	mBase = m.M
	switch v26 + int32(1) {
	case 0:
		goto L9
	case 1:
		v102 = int32(15)
		goto L4
	default:
		goto L8
	}
L6:
	;
	v102 = int32(73)
	goto L4
L7:
	;
	v42 = int64(10)
	if v42 < v21 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v33 = v26 + v19
	v34 = v17 - v26
	if v34 == int32(0) {
		v117 = v33
		goto L1
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v30 == int32(6) {
		v38 = v16
		v39 = v17
		v40 = v19
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L3
L12:
	;
	v38 = v16 + v26
	v39 = v34
	v40 = v33
	goto L7
L13:
	;
	v45 = v21
	goto L15
L14:
	;
	v45 = v42
	goto L15
L15:
	;
	v47 = m.G0
	v49 = v47 - int32(16)
	m.G0 = v49
	*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = l0
	goto L18
L16:
	;
	v93 = F_mstime(m)
	mBase = m.M
	v94 = v93 - v12
	if v94 < l3 {
		v16 = v38
		v17 = v39
		v19 = v40
		v21 = l3 - v94
		goto L5
	} else {
		goto L26
	}
L17:
	;
	goto L19
L18:
	;
	v59 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+12)) = uint16(v59)
	goto L17
L19:
	;
	v70 = int32(1)
	v72 = F_poll(m, v49+int32(8), v70, base.I32_wrap_i64(v45))
	mBase = m.M
	if v72 != v70 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	m.G0 = v49 + int32(16)
	goto L16
L22:
	;
	goto L23
L23:
	;
	goto L25
L25:
	;
	goto L21
L26:
	;
	goto L6
L27:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v102
	goto L3
}
func F_syncWithPrimary(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
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
	var v223 int32
	_ = v223
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
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int64
	_ = v332
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int64
	_ = v361
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int64
	_ = v389
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int64
	_ = v417
	var v418 int32
	_ = v418
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int64
	_ = v445
	var v446 int32
	_ = v446
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
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
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int64
	_ = v680
	var v683 int64
	_ = v683
	var v701 int64
	_ = v701
	var v705 int32
	_ = v705
	v6 = m.G0
	v8 = v6 - int32(656)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+652)) = l0
	v12 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(656)
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 == int32(3) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	m.T0[v14].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
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
	switch v12 + int32(-2) {
	case 0:
		goto L31
	case 1:
		goto L30
	case 2:
		goto L29
	case 3:
		goto L28
	case 4:
		goto L27
	case 5:
		goto L26
	case 6:
		goto L25
	case 7:
		goto L24
	case 8:
		goto L23
	case 9:
		goto L21
	case 10:
		goto L19
	default:
		goto L20
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v21 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	m.T0[v36].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L12
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
	v26 = m.T0[v25].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+336)) = v26
	F__serverLog(m, int32(3), int32(_a1275), v8+int32(336))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v39 = int32(_a44)
	v40 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[680])) = v40
	v43 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	if v43 == v40 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_cleanupTransferResources(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L18
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	m.T0[v47].(func(*base.Module, int32))(m, v43)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v50 = int32(_a44)
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[699])) = v51
	v54 = *(*int32)(unsafe.Add(mBase, _consts[680]))
	if v54 == v51 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+52))
	m.T0[v58].(func(*base.Module, int32))(m, v54)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[680])) = int32(0)
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(1)
	goto L1
L19:
	;
	v250 = F_replicaProcessPsyncReply(m, l0)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L90
	}
L20:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v207 {
		goto L80
	} else {
		goto L81
	}
L21:
	;
	v179 = F_replicaSendPsyncCommand(m, l0)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L71
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(11)
	goto L21
L23:
	;
	v168 = F_syncWithPrimaryHandleReceiveNodeIDReplyState(m, l0)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L67
	}
L24:
	;
	v153 = F_syncWithPrimaryHandleReceiveVersionReplyState(m, l0)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L63
	}
L25:
	;
	v142 = F_syncWithPrimaryHandleReceiveCapaReplyState(m, l0)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L59
	}
L26:
	;
	v129 = F_syncWithPrimaryHandleReceiveIPReplyState(m, l0)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L54
	}
L27:
	;
	v118 = F_syncWithPrimaryHandleReceivePortReplyState(m, l0)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L50
	}
L28:
	;
	v105 = F_syncWithPrimaryHandleReceiveAuthReplyState(m, l0)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L45
	}
L29:
	;
	v94 = F_syncWithPrimaryHandleSendHandshakeState(m, l0)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L41
	}
L30:
	;
	v83 = F_syncWithPrimaryHandleReceivePingReplyState(m, l0)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L37
	}
L31:
	;
	v72 = F_syncWithPrimaryHandleConnectingState(m, l0)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(3)
	goto L1
L33:
	;
	if v72 != int32(-1) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	goto L1
L36:
	;
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(4)
	goto L29
L37:
	;
	if v83 != int32(-1) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	goto L1
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(5)
	goto L1
L41:
	;
	if v94 != int32(-1) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	goto L1
L44:
	;
	v113 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(6)
	v117 = *(*int32)(unsafe.Add(mBase, _consts[215]))
	if v117 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	if v105 != int32(-1) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	goto L1
L48:
	;
	goto L27
L49:
	;
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(7)
	goto L1
L50:
	;
	if v118 != int32(-1) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	goto L1
L53:
	;
	v137 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(8)
	v141 = *(*int32)(unsafe.Add(mBase, _consts[702]))
	if v141 != 0 {
		goto L1
	} else {
		goto L57
	}
L54:
	;
	if v129 != int32(-1) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	goto L1
L57:
	;
	goto L25
L58:
	;
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(9)
	goto L1
L59:
	;
	if v142 != int32(-1) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	goto L1
L62:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v162 == int32(0) {
		goto L22
	} else {
		goto L66
	}
L63:
	;
	if v153 != int32(-1) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	goto L1
L66:
	;
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(10)
	goto L1
L67:
	;
	if v168 != int32(-1) {
		goto L22
	} else {
		goto L68
	}
L68:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	goto L1
L70:
	;
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(12)
	goto L1
L71:
	;
	if v179 != 0 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v182 = F_sdsnew(m, int32(_a1262))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	F_abortFailover(m, v182)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v187 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_sdsfree(m, v182)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L78
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v182
	F__serverLog(m, int32(3), int32(_a1263), v8+int32(16))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	goto L1
L80:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+52))
	m.T0[v216].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
	F__serverLog(m, int32(3), int32(_a1276), v8)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v219 = int32(_a44)
	v220 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[680])) = v220
	v223 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	if v223 == v220 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	F_cleanupTransferResources(m)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L89
	}
L85:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+52))
	m.T0[v227].(func(*base.Module, int32))(m, v223)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v230 = int32(_a44)
	v231 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[699])) = v231
	v234 = *(*int32)(unsafe.Add(mBase, _consts[680]))
	if v234 == v231 {
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+52))
	m.T0[v238].(func(*base.Module, int32))(m, v234)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, _consts[680])) = int32(0)
	goto L84
L89:
	;
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(1)
	goto L1
L90:
	;
	if v250 == int32(1) {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _consts[658]))
	if v255 != int32(2) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	F_abortFailover(m, int32(_a1277))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L4
	} else {
		goto L195
	}
L93:
	;
	switch v250 + int32(-2) {
	case 0:
		goto L101
	default:
		goto L99
	case 2:
		goto L100
	case 3:
		goto L102
	}
L94:
	;
	if v250 == int32(6) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	F_clearFailoverState(m, int32(1))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L98
	}
L96:
	;
	if v250&int32(6) != int32(2) {
		goto L92
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	goto L93
L99:
	;
	v329 = F_useDisklessLoad(m)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L120
	}
L100:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v289 {
		goto L109
	} else {
		goto L110
	}
L101:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v274 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	goto L1
L104:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	if v283 != int32(2) {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	F__serverLog(m, int32(2), int32(_a1264), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	goto L108
L108:
	;
	goto L1
L109:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+92))
	v306 = m.T0[v305].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, int32(_a1272), int32(6), base.I64_extend_i32_s(v300*int32(1000)))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L112
	}
L110:
	;
	F__serverLog(m, int32(2), int32(_a1274), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	if v306 != int32(-1) {
		goto L99
	} else {
		goto L113
	}
L113:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v311 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L118
	}
L115:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+88))
	v316 = m.T0[v315].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+320)) = v316
	F__serverLog(m, int32(3), int32(_a1273), v8+int32(320))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	goto L114
L118:
	;
	goto L1
L119:
	;
	if v250 != int32(6) {
		goto L150
	} else {
		goto L151
	}
L120:
	;
	if v329 != 0 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v332 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v333 = F___syscall_getpid(m)
	mBase = m.M
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+308)) = v333
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+304)) = uint32(v332)
	v342 = F_snprintf(m, v8+int32(384), int32(256), int32(_a1270), v8+int32(304))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+288)) = int32(438)
	v351 = F_open(m, v8+int32(384), int32(193), v8+int32(288))
	mBase = m.M
	if v351 != int32(-1) {
		v488 = v351
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v493 = F_zstrdup(m, v8+int32(384))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L149
	}
L125:
	;
	v354 = int32(9116376)
	goto L126
L126:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v357 = F_sleep(m, int32(1))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v355
	v361 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v362 = F___syscall_getpid(m)
	mBase = m.M
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+276)) = v362
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+272)) = uint32(v361)
	v371 = F_snprintf(m, v8+int32(384), int32(256), int32(_a1270), v8+int32(272))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+256)) = int32(438)
	v380 = F_open(m, v8+int32(384), int32(193), v8+int32(256))
	mBase = m.M
	if v380 != int32(-1) {
		v488 = v380
		goto L124
	} else {
		goto L130
	}
L130:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v385 = F_sleep(m, int32(1))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v383
	v389 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v390 = F___syscall_getpid(m)
	mBase = m.M
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+244)) = v390
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+240)) = uint32(v389)
	v399 = F_snprintf(m, v8+int32(384), int32(256), int32(_a1270), v8+int32(240))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+224)) = int32(438)
	v408 = F_open(m, v8+int32(384), int32(193), v8+int32(224))
	mBase = m.M
	if v408 != int32(-1) {
		v488 = v408
		goto L124
	} else {
		goto L134
	}
L134:
	;
	v411 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v413 = F_sleep(m, int32(1))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v411
	v417 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v418 = F___syscall_getpid(m)
	mBase = m.M
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+212)) = v418
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+208)) = uint32(v417)
	v427 = F_snprintf(m, v8+int32(384), int32(256), int32(_a1270), v8+int32(208))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+192)) = int32(438)
	v436 = F_open(m, v8+int32(384), int32(193), v8+int32(192))
	mBase = m.M
	if v436 != int32(-1) {
		v488 = v436
		goto L124
	} else {
		goto L138
	}
L138:
	;
	v439 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v441 = F_sleep(m, int32(1))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v439
	v445 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v446 = F___syscall_getpid(m)
	mBase = m.M
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+180)) = v446
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+176)) = uint32(v445)
	v455 = F_snprintf(m, v8+int32(384), int32(256), int32(_a1270), v8+int32(176))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+160)) = int32(438)
	v464 = F_open(m, v8+int32(384), int32(193), v8+int32(160))
	mBase = m.M
	if v464 != int32(-1) {
		v488 = v464
		goto L124
	} else {
		goto L142
	}
L142:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v469 = F_sleep(m, int32(1))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v467
	v473 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v473 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L4
	} else {
		goto L148
	}
L145:
	;
	v476 = F___strerror_l(m, v467, v467)
	mBase = m.M
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v476
	F__serverLog(m, int32(3), int32(_a1271), v8+int32(32))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	goto L144
L148:
	;
	goto L1
L149:
	;
	v495 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[711])) = v488
	*(*int32)(unsafe.Add(mBase, _consts[712])) = v493
	goto L119
L150:
	;
	v600 = F_useDisklessLoad(m)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L4
	} else {
		goto L173
	}
L151:
	;
	v505 = F_connTypeOfReplication(m)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v505)+40))
	v508 = m.T0[v507].(func(*base.Module) int32)(m)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, _consts[699])) = v508
	v511 = int32(_a44)
	v512 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v514 = *(*int32)(unsafe.Add(mBase, _consts[290]))
	v516 = *(*int32)(unsafe.Add(mBase, _consts[202]))
	v518 = *(*int32)(unsafe.Add(mBase, _consts[710]))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v520)+56))
	v522 = m.T0[v521].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v508, v512, v514, v516, v518, int32(975))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L155
	}
L154:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v8)+652))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v557)))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v559)+84))
	v561 = m.T0[v560].(func(*base.Module, int32, int32) int32)(m, v557, int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L4
	} else {
		goto L164
	}
L155:
	;
	if v522 != int32(-1) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v527 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v527 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v545 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)+52))
	m.T0[v547].(func(*base.Module, int32))(m, v545)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L4
	} else {
		goto L161
	}
L158:
	;
	v531 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)+88))
	v534 = m.T0[v533].(func(*base.Module, int32) int32)(m, v531)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v534
	F__serverLog(m, int32(3), int32(_a1269), v8+int32(48))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	goto L157
L161:
	;
	*(*int32)(unsafe.Add(mBase, _consts[699])) = int32(0)
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	goto L1
L163:
	;
	*(*int32)(unsafe.Add(mBase, _consts[697])) = int32(1)
	goto L1
L164:
	;
	if v561 != int32(-1) {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v566 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v566 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L4
	} else {
		goto L172
	}
L167:
	;
	goto L168
L168:
	;
	v570 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v571 = F___strerror_l(m, v570, v570)
	mBase = m.M
	goto L169
L169:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v557)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v572
	v580 = F_snprintf(m, v8+int32(352), int32(31), int32(_a1265), v8+int32(80))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v571
	*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v8 + int32(352)
	F__serverLog(m, int32(3), int32(_a1268), v8+int32(64))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	goto L166
L172:
	;
	goto L1
L173:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v8)+652))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v602)))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)+84))
	if v600 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v679 = int32(_a44)
	v680 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[704])) = v680
	v683 = int64(-1)
	*(*int64)(unsafe.Add(mBase, _consts[705])) = v683
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(13)
	*(*int64)(unsafe.Add(mBase, _consts[706])) = v680
	*(*int64)(unsafe.Add(mBase, _consts[707])) = v683
	*(*int64)(unsafe.Add(mBase, _consts[708])) = v680
	v701 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	*(*int64)(unsafe.Add(mBase, _consts[709])) = v701
	goto L1
L175:
	;
	v643 = m.T0[v604].(func(*base.Module, int32, int32) int32)(m, v602, int32(977))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L4
	} else {
		goto L186
	}
L176:
	;
	v606 = m.T0[v604].(func(*base.Module, int32, int32) int32)(m, v602, int32(976))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	if v606 != int32(-1) {
		goto L174
	} else {
		goto L178
	}
L178:
	;
	v611 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v611 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L4
	} else {
		goto L185
	}
L180:
	;
	goto L181
L181:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v616 = F___strerror_l(m, v615, v615)
	mBase = m.M
	goto L182
L182:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v602)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = v617
	v625 = F_snprintf(m, v8+int32(352), int32(31), int32(_a1265), v8+int32(112))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v616
	*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = v8 + int32(352)
	F__serverLog(m, int32(3), int32(_a1267), v8+int32(96))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	goto L179
L185:
	;
	goto L1
L186:
	;
	if v643 != int32(-1) {
		goto L174
	} else {
		goto L187
	}
L187:
	;
	v648 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v648 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	F_syncWithPrimaryHandleError(m, v8+int32(652))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L4
	} else {
		goto L194
	}
L189:
	;
	goto L190
L190:
	;
	v652 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v653 = F___strerror_l(m, v652, v652)
	mBase = m.M
	goto L191
L191:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v602)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = v654
	v662 = F_snprintf(m, v8+int32(352), int32(31), int32(_a1265), v8+int32(144))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+128)) = v653
	*(*int32)(unsafe.Add(mBase, uint32(v8)+132)) = v8 + int32(352)
	F__serverLog(m, int32(3), int32(_a1266), v8+int32(128))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	goto L188
L194:
	;
	goto L1
L195:
	;
	goto L1
}
func F_syncWithPrimaryHandleError(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+52))
	m.T0[v5].(func(*base.Module, int32))(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v8
		v10 = int32(_a44)
		*(*int32)(unsafe.Add(mBase, _consts[680])) = v8
		v14 = *(*int32)(unsafe.Add(mBase, _consts[699]))
		if v14 == v8 {
			F_cleanupTransferResources(m)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(1)
				return
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
			m.T0[v18].(func(*base.Module, int32))(m, v14)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = int32(_a44)
				v22 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[699])) = v22
				v25 = *(*int32)(unsafe.Add(mBase, _consts[680]))
				if v25 == v22 {
					F_cleanupTransferResources(m)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(1)
						return
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
					m.T0[v29].(func(*base.Module, int32))(m, v25)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[680])) = int32(0)
						F_cleanupTransferResources(m)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(1)
							return
						}
					}
				}
			}
		}
	}
}
func F_syncWithPrimaryHandleReceiveCapaReplyState(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	v4 = m.G0
	v6 = v4 - int32(288)
	m.G0 = v6
	v8 = int32(-1)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v19 = m.T0[v18].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v6+int32(32), int32(256), base.I64_extend_i32_s(v13*int32(1000)))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 != int32(-1) {
			v38 = int32(_a44)
			v40 = *(*int64)(unsafe.Add(mBase, _consts[47]))
			*(*int64)(unsafe.Add(mBase, _consts[709])) = v40
			v44 = F_sdsnew(m, v6+int32(32))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				if v44 == int32(0) {
					v66 = v8
					m.G0 = v6 + int32(288)
					return v66
				} else {
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
					if v48 != int32(45) {
						F_sdsfree(m, v44)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v66 = int32(0)
							m.G0 = v6 + int32(288)
							return v66
						}
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if int32(2) < v52 {
							F_sdsfree(m, v44)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v66 = int32(0)
								m.G0 = v6 + int32(288)
								return v66
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v44
							F__serverLog(m, int32(2), int32(_a1285), v6+int32(16))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_sdsfree(m, v44)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v66 = int32(0)
									m.G0 = v6 + int32(288)
									return v66
								}
							}
						}
					}
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(3) < v26 {
				v66 = v8
				m.G0 = v6 + int32(288)
				return v66
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
				v31 = m.T0[v30].(func(*base.Module, int32) int32)(m, l0)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v31
					F__serverLog(m, int32(3), int32(_a1278), v6)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v66 = v8
						m.G0 = v6 + int32(288)
						return v66
					}
				}
			}
		}
	}
}
func F_syncWithPrimaryHandleReceivePingReplyState(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	v4 = m.G0
	v6 = v4 - int32(288)
	m.G0 = v6
	v8 = int32(-1)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v19 = m.T0[v18].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v6+int32(32), int32(256), base.I64_extend_i32_s(v13*int32(1000)))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v6 + int32(288)
	return v224
L2:
	;
	v38 = int32(_a44)
	v40 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	*(*int64)(unsafe.Add(mBase, _consts[709])) = v40
	v44 = F_sdsnew(m, v6+int32(32))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L9
	}
L3:
	;
	return int32(0)
L4:
	;
	if v19 != int32(-1) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v26 {
		v224 = v8
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
	v31 = m.T0[v30].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v31
	F__serverLog(m, int32(3), int32(_a1278), v6)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v224 = v8
	goto L1
L9:
	;
	if v44 == int32(0) {
		v224 = v8
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v48 == int32(43) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v212 {
		goto L59
	} else {
		goto L60
	}
L12:
	;
	v51 = int32(_a1279)
	goto L14
L13:
	;
	if v85-v90 == int32(0) {
		goto L11
	} else {
		goto L26
	}
L14:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v56 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	goto L13
L17:
	;
	v58 = v44
	v59 = v51
	v60 = int32(7)
	v61 = v56
	goto L20
L18:
	;
	v85 = int32(0)
	v86 = v51
	goto L16
L19:
	;
	v85 = v82 & int32(255)
	v86 = v80
	goto L16
L20:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v61&int32(255) != v65 {
		v80 = v59
		v82 = v61
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v80 = v74
	v82 = int32(0)
	goto L19
L22:
	;
	if v65 == int32(0) {
		v80 = v59
		v82 = v61
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v70 = v60 + int32(-1)
	if v70 == int32(0) {
		v80 = v59
		v82 = v61
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v73 = int32(1)
	v74 = v59 + v73
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v75 != 0 {
		v58 = v58 + v73
		v59 = v74
		v60 = v70
		v61 = v75
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v100 = int32(_a1280)
	goto L28
L27:
	;
	if v134-v139 == int32(0) {
		goto L11
	} else {
		goto L40
	}
L28:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v105 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	goto L27
L31:
	;
	v107 = v44
	v108 = v100
	v109 = int32(7)
	v110 = v105
	goto L34
L32:
	;
	v134 = int32(0)
	v135 = v100
	goto L30
L33:
	;
	v134 = v131 & int32(255)
	v135 = v129
	goto L30
L34:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v110&int32(255) != v114 {
		v129 = v108
		v131 = v110
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v129 = v123
	v131 = int32(0)
	goto L33
L36:
	;
	if v114 == int32(0) {
		v129 = v108
		v131 = v110
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v119 = v109 + int32(-1)
	if v119 == int32(0) {
		v129 = v108
		v131 = v110
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v122 = int32(1)
	v123 = v108 + v122
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	if v124 != 0 {
		v107 = v107 + v122
		v108 = v123
		v109 = v119
		v110 = v124
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v149 = int32(_a1281)
	goto L42
L41:
	;
	if v183-v188 == int32(0) {
		goto L11
	} else {
		goto L54
	}
L42:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v154 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	goto L41
L45:
	;
	v156 = v44
	v157 = v149
	v158 = int32(28)
	v159 = v154
	goto L48
L46:
	;
	v183 = int32(0)
	v184 = v149
	goto L44
L47:
	;
	v183 = v180 & int32(255)
	v184 = v178
	goto L44
L48:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v159&int32(255) != v163 {
		v178 = v157
		v180 = v159
		goto L47
	} else {
		goto L50
	}
L49:
	;
	v178 = v172
	v180 = int32(0)
	goto L47
L50:
	;
	if v163 == int32(0) {
		v178 = v157
		v180 = v159
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v168 = v158 + int32(-1)
	if v168 == int32(0) {
		v178 = v157
		v180 = v159
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v171 = int32(1)
	v172 = v157 + v171
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	if v173 != 0 {
		v156 = v156 + v171
		v157 = v172
		v158 = v168
		v159 = v173
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L49
L54:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v199 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	F_sdsfree(m, v44)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L3
	} else {
		goto L58
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v44
	F__serverLog(m, int32(3), int32(_a1282), v6+int32(16))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v224 = v8
	goto L1
L59:
	;
	F_sdsfree(m, v44)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L3
	} else {
		goto L62
	}
L60:
	;
	F__serverLog(m, int32(2), int32(_a1283), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v224 = int32(0)
	goto L1
}
func F_syncWithPrimaryHandleReceivePortReplyState(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	v4 = m.G0
	v6 = v4 - int32(288)
	m.G0 = v6
	v8 = int32(-1)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v19 = m.T0[v18].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v6+int32(32), int32(256), base.I64_extend_i32_s(v13*int32(1000)))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 != int32(-1) {
			v38 = int32(_a44)
			v40 = *(*int64)(unsafe.Add(mBase, _consts[47]))
			*(*int64)(unsafe.Add(mBase, _consts[709])) = v40
			v44 = F_sdsnew(m, v6+int32(32))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				if v44 == int32(0) {
					v66 = v8
					m.G0 = v6 + int32(288)
					return v66
				} else {
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
					if v48 != int32(45) {
						F_sdsfree(m, v44)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v66 = int32(0)
							m.G0 = v6 + int32(288)
							return v66
						}
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if int32(2) < v52 {
							F_sdsfree(m, v44)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v66 = int32(0)
								m.G0 = v6 + int32(288)
								return v66
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v44
							F__serverLog(m, int32(2), int32(_a1284), v6+int32(16))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_sdsfree(m, v44)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v66 = int32(0)
									m.G0 = v6 + int32(288)
									return v66
								}
							}
						}
					}
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(3) < v26 {
				v66 = v8
				m.G0 = v6 + int32(288)
				return v66
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
				v31 = m.T0[v30].(func(*base.Module, int32) int32)(m, l0)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v31
					F__serverLog(m, int32(3), int32(_a1278), v6)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v66 = v8
						m.G0 = v6 + int32(288)
						return v66
					}
				}
			}
		}
	}
}
func F_syncWithPrimaryHandleReceiveVersionReplyState(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	v4 = m.G0
	v6 = v4 - int32(288)
	m.G0 = v6
	v8 = int32(-1)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v19 = m.T0[v18].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v6+int32(32), int32(256), base.I64_extend_i32_s(v13*int32(1000)))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 != int32(-1) {
			v38 = int32(_a44)
			v40 = *(*int64)(unsafe.Add(mBase, _consts[47]))
			*(*int64)(unsafe.Add(mBase, _consts[709])) = v40
			v44 = F_sdsnew(m, v6+int32(32))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				if v44 == int32(0) {
					v66 = v8
					m.G0 = v6 + int32(288)
					return v66
				} else {
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
					if v48 != int32(45) {
						F_sdsfree(m, v44)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v66 = int32(0)
							m.G0 = v6 + int32(288)
							return v66
						}
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if int32(2) < v52 {
							F_sdsfree(m, v44)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v66 = int32(0)
								m.G0 = v6 + int32(288)
								return v66
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v44
							F__serverLog(m, int32(2), int32(_a1286), v6+int32(16))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_sdsfree(m, v44)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v66 = int32(0)
									m.G0 = v6 + int32(288)
									return v66
								}
							}
						}
					}
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(3) < v26 {
				v66 = v8
				m.G0 = v6 + int32(288)
				return v66
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
				v31 = m.T0[v30].(func(*base.Module, int32) int32)(m, l0)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v31
					F__serverLog(m, int32(3), int32(_a1278), v6)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v66 = v8
						m.G0 = v6 + int32(288)
						return v66
					}
				}
			}
		}
	}
}
