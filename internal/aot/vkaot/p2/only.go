package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_flushAppendOnlyFile(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int64
	_ = v171
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v240 int64
	_ = v240
	var v243 int64
	_ = v243
	var v245 int64
	_ = v245
	var v249 int64
	_ = v249
	var v257 int32
	_ = v257
	var v261 int64
	_ = v261
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v277 int64
	_ = v277
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v317 int64
	_ = v317
	var v319 int32
	_ = v319
	var v322 int64
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int64
	_ = v384
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v434 int64
	_ = v434
	var v435 int64
	_ = v435
	var v440 int64
	_ = v440
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
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
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int64
	_ = v568
	var v569 int64
	_ = v569
	var v574 int64
	_ = v574
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v609 int64
	_ = v609
	var v612 int64
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v677 int64
	_ = v677
	var v683 int64
	_ = v683
	var v684 int64
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v705 int64
	_ = v705
	var v708 int64
	_ = v708
	var v710 int64
	_ = v710
	var v713 int64
	_ = v713
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v725 int64
	_ = v725
	var v729 int64
	_ = v729
	var v733 int64
	_ = v733
	var v735 int32
	_ = v735
	var v736 int64
	_ = v736
	var v738 int64
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int64
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int64
	_ = v751
	var v754 int64
	_ = v754
	var v755 int64
	_ = v755
	v11 = m.G0
	v13 = v11 - int32(64)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-1)))))
	switch v19 & int32(7) {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	case 3:
		goto L7
	case 4:
		goto L6
	default:
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(64)
	return
L2:
	;
	v665 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	if v665 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L3:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v89 != int32(2) {
		v135 = int32(0)
		goto L25
	} else {
		goto L26
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v40 != int32(2) {
		v67 = v40
		v69 = int32(0)
		goto L13
	} else {
		goto L14
	}
L5:
	;
	if v36 != 0 {
		goto L3
	} else {
		goto L11
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
	v36 = v35
	goto L5
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
	v36 = v32
	goto L5
L8:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
	v36 = v29
	goto L5
L9:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
	v36 = v26
	goto L5
L10:
	;
	v36 = int32(base.Ui32(v19) >> (uint(int32(3)) % 32))
	goto L5
L11:
	;
	goto L4
L12:
	;
	if v79 == int32(0) {
		goto L1
	} else {
		goto L23
	}
L13:
	;
	v70 = int32(1)
	if v67 == v70 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v43 = int32(1)
	v44 = int32(0)
	v45 = int32(_a44)
	v46 = *(*int64)(unsafe.Add(mBase, _consts[33]))
	v48 = *(*int64)(unsafe.Add(mBase, _consts[30]))
	if v46 == v48 {
		v79 = v43
		v80 = v44
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v50 = int32(_a44)
	v51 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	v53 = *(*int64)(unsafe.Add(mBase, _consts[36]))
	if v51-v53 < int64(1000) {
		v79 = v43
		v80 = v44
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	goto L18
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	v67 = v66
	v69 = int32(1)
	goto L13
L18:
	;
	if v63 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v658 = int32(0)
	goto L2
L20:
	;
	v73 = int32(_a44)
	v74 = *(*int64)(unsafe.Add(mBase, _consts[33]))
	v76 = *(*int64)(unsafe.Add(mBase, _consts[30]))
	if v74 != v76 {
		v658 = v69
		goto L2
	} else {
		goto L22
	}
L21:
	;
	v79 = v67
	v80 = v69
	goto L12
L22:
	;
	v79 = v70
	v80 = v69
	goto L12
L23:
	;
	if v80 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v83 = int32(_a44)
	v85 = *(*int64)(unsafe.Add(mBase, _consts[40]))
	*(*int64)(unsafe.Add(mBase, _consts[32])) = v85
	goto L1
L25:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[41]))
	if v139 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	goto L27
L27:
	;
	v99 = base.B2i32(v97 != int32(0))
	if l0 != 0 {
		v135 = v99
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v101 != int32(2) {
		v135 = v99
		goto L25
	} else {
		goto L29
	}
L29:
	;
	if v97 == int32(0) {
		v135 = v99
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v106 = int32(_a44)
	v107 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	v109 = *(*int64)(unsafe.Add(mBase, _consts[42]))
	if v109 != int64(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v107-v109 < int64(2000) {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, _consts[42])) = v107
	goto L1
L33:
	;
	v117 = int32(1)
	v118 = int32(_a44)
	v120 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v120 + v117
	v125 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v125 {
		v135 = v117
		goto L25
	} else {
		goto L34
	}
L34:
	;
	F__serverLog(m, int32(2), int32(_a114), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return
L36:
	;
	v135 = int32(1)
	goto L25
L37:
	;
	v171 = *(*int64)(unsafe.Add(mBase, _consts[44]))
	if base.B2i32(v171 == int64(0)) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L38:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143+int32(-1)))))
	switch v146 & int32(7) {
	case 0:
		goto L44
	case 1:
		goto L43
	case 2:
		goto L42
	case 3:
		goto L41
	case 4:
		goto L40
	default:
		goto L37
	}
L39:
	;
	if v163 == int32(0) {
		goto L37
	} else {
		goto L45
	}
L40:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v143+int32(-17))))
	v163 = v162
	goto L39
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v143+int32(-9))))
	v163 = v159
	goto L39
L42:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143+int32(-5)))))
	v163 = v156
	goto L39
L43:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143+int32(-3)))))
	v163 = v153
	goto L39
L44:
	;
	v163 = int32(base.Ui32(v146) >> (uint(int32(3)) % 32))
	goto L39
L45:
	;
	v166 = F_usleep(m, v139)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L35
	} else {
		goto L46
	}
L46:
	;
	goto L37
L47:
	;
	v179 = int32(_a44)
	v180 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v183 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+int32(-1)))))
	switch v186 & int32(7) {
	case 0:
		goto L56
	case 1:
		goto L55
	case 2:
		goto L54
	case 3:
		goto L53
	case 4:
		goto L52
	default:
		v237 = int32(0)
		goto L50
	}
L48:
	;
	v177 = F_ustime(m)
	mBase = m.M
	v178 = v177
	goto L47
L49:
	;
	v178 = int64(0)
	goto L47
L50:
	;
	v240 = *(*int64)(unsafe.Add(mBase, _consts[44]))
	if v240 == int64(0) {
		v245 = v178
		goto L68
	} else {
		goto L69
	}
L51:
	;
	v204 = int32(0)
	if v203 == v204 {
		v237 = v204
		goto L50
	} else {
		goto L57
	}
L52:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v183+int32(-17))))
	v203 = v202
	goto L51
L53:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v183+int32(-9))))
	v203 = v199
	goto L51
L54:
	;
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183+int32(-5)))))
	v203 = v196
	goto L51
L55:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+int32(-3)))))
	v203 = v193
	goto L51
L56:
	;
	v203 = int32(base.Ui32(v186) >> (uint(int32(3)) % 32))
	goto L51
L57:
	;
	v207 = v183
	v209 = v203
	v215 = v204
	goto L58
L58:
	;
	v217 = F_write(m, v180, v207, v209)
	mBase = m.M
	if int32(-1) < v217 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v237 = v226
	goto L50
L60:
	;
	v226 = v217 + v215
	v228 = v209 - v217
	if v228 != 0 {
		v207 = v207 + v217
		v209 = v228
		v215 = v226
		goto L58
	} else {
		goto L67
	}
L61:
	;
	goto L62
L62:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v221 == int32(27) {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	if v215 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v225 = v215
	goto L66
L65:
	;
	v225 = int32(-1)
	goto L66
L66:
	;
	v237 = v225
	goto L50
L67:
	;
	goto L59
L68:
	;
	if v135 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v243 = F_ustime(m)
	mBase = m.M
	v245 = v243 - v178
	goto L68
L70:
	;
	v277 = *(*int64)(unsafe.Add(mBase, _consts[44]))
	if v277 == int64(0) {
		goto L83
	} else {
		goto L84
	}
L71:
	;
	F_latencyAddSample(m, v270, v245)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L35
	} else {
		goto L82
	}
L72:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	goto L76
L73:
	;
	v249 = *(*int64)(unsafe.Add(mBase, _consts[44]))
	if v249 == int64(0) {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	if v249*int64(1000) <= v245 {
		v270 = int32(_a115)
		goto L71
	} else {
		goto L75
	}
L75:
	;
	goto L70
L76:
	;
	v261 = *(*int64)(unsafe.Add(mBase, _consts[44]))
	if v261 == int64(0) {
		goto L70
	} else {
		goto L77
	}
L77:
	;
	if v245 < v261*int64(1000) {
		goto L70
	} else {
		goto L78
	}
L78:
	;
	if v257 != int32(-1) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v269 = int32(_a116)
	goto L81
L80:
	;
	v269 = int32(_a117)
	goto L81
L81:
	;
	v270 = v269
	goto L71
L82:
	;
	goto L70
L83:
	;
	v286 = int32(_a44)
	*(*int64)(unsafe.Add(mBase, _consts[42])) = int64(0)
	v291 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291+int32(-1)))))
	switch v294 & int32(7) {
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
		v311 = int32(0)
		goto L87
	}
L84:
	;
	if v245 < v277*int64(1000) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	F_latencyAddSample(m, int32(_a118), v245)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L35
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	if v237 == v311 {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v291+int32(-17))))
	v311 = v310
	goto L87
L89:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v291+int32(-9))))
	v311 = v307
	goto L87
L90:
	;
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v291+int32(-5)))))
	v311 = v304
	goto L87
L91:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291+int32(-3)))))
	v311 = v301
	goto L87
L92:
	;
	v311 = int32(base.Ui32(v294) >> (uint(int32(3)) % 32))
	goto L87
L93:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	if v547 != int32(-1) {
		v565 = v291
		goto L161
	} else {
		goto L162
	}
L94:
	;
	v314 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v316 = *(*int64)(unsafe.Add(mBase, _consts[48]))
	v317 = v314 - v316
	v319 = base.B2i32(v317 < int64(31))
	if v317 < int64(31) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v325 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v328 = v319 | base.B2i32(int32(3) < v325)
	if v237 != int32(-1) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v322 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	*(*int64)(unsafe.Add(mBase, _consts[48])) = v322
	goto L95
L97:
	;
	v410 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v406
	v413 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v413 != int32(1) {
		goto L121
	} else {
		goto L122
	}
L98:
	;
	if v328 != 0 {
		goto L105
	} else {
		goto L106
	}
L99:
	;
	v331 = int32(9116376)
	goto L100
L100:
	;
	if v328 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v406 = v341
	v409 = int32(-1)
	goto L97
L102:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v333 = F___strerror_l(m, v332, v332)
	mBase = m.M
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v333
	F__serverLog(m, int32(3), int32(_a119), v13+int32(16))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L35
	} else {
		goto L104
	}
L104:
	;
	goto L101
L105:
	;
	v380 = int32(51)
	v381 = int32(_a44)
	v382 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v384 = *(*int64)(unsafe.Add(mBase, _consts[30]))
	v385 = F_ftruncate(m, v382, v384)
	mBase = m.M
	if v385 == int32(-1) {
		goto L114
	} else {
		goto L115
	}
L106:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346+int32(-1)))))
	switch v349 & int32(7) {
	case 0:
		goto L112
	case 1:
		goto L111
	case 2:
		goto L110
	case 3:
		goto L109
	case 4:
		goto L108
	default:
		v366 = int32(0)
		goto L107
	}
L107:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = base.I64_extend_i32_u(v237)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = base.I64_extend_i32_u(v366)
	F__serverLog(m, int32(3), int32(_a120), v13+int32(48))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L35
	} else {
		goto L113
	}
L108:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v346+int32(-17))))
	v366 = v365
	goto L107
L109:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v346+int32(-9))))
	v366 = v362
	goto L107
L110:
	;
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v346+int32(-5)))))
	v366 = v359
	goto L107
L111:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346+int32(-3)))))
	v366 = v356
	goto L107
L112:
	;
	v366 = int32(base.Ui32(v349) >> (uint(int32(3)) % 32))
	goto L107
L113:
	;
	goto L105
L114:
	;
	if v317 < int64(31) {
		v406 = v380
		v409 = v237
		goto L97
	} else {
		goto L116
	}
L115:
	;
	v406 = v380
	v409 = int32(-1)
	goto L97
L116:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v392 {
		v406 = v380
		v409 = v237
		goto L97
	} else {
		goto L117
	}
L117:
	;
	goto L118
L118:
	;
	v396 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v397 = F___strerror_l(m, v396, v396)
	mBase = m.M
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v397
	F__serverLog(m, int32(3), int32(_a121), v13+int32(32))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L35
	} else {
		goto L120
	}
L120:
	;
	v406 = v380
	v409 = v237
	goto L97
L121:
	;
	*(*int32)(unsafe.Add(mBase, _consts[46])) = int32(-1)
	if v409 < int32(1) {
		goto L1
	} else {
		goto L126
	}
L122:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v417 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F__serverLog(m, int32(3), int32(_a122), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L35
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v432 = int32(_a44)
	v434 = *(*int64)(unsafe.Add(mBase, _consts[50]))
	v435 = base.I64_extend_i32_u(v409)
	*(*int64)(unsafe.Add(mBase, _consts[50])) = v434 + v435
	v440 = *(*int64)(unsafe.Add(mBase, _consts[30]))
	*(*int64)(unsafe.Add(mBase, _consts[30])) = v440 + v435
	v444 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v445 = int32(-1)
	v453 = v444 + v445
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453))))
	v456 = v454 & int32(7)
	switch v456 {
	case 0:
		goto L134
	case 1:
		goto L133
	case 2:
		goto L132
	case 3:
		goto L131
	case 4:
		goto L130
	default:
		goto L128
	}
L127:
	;
	goto L1
L128:
	;
	goto L127
L129:
	;
	if v471 == int32(0) {
		goto L128
	} else {
		goto L135
	}
L130:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v444+int32(-17))))
	v471 = v470
	goto L129
L131:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v444+int32(-9))))
	v471 = v467
	goto L129
L132:
	;
	v464 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444+int32(-5)))))
	v471 = v464
	goto L129
L133:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444+int32(-3)))))
	v471 = v461
	goto L129
L134:
	;
	v471 = int32(base.Ui32(v454) >> (uint(int32(3)) % 32))
	goto L129
L135:
	;
	v477 = int32(-1)&v471 + v445
	v481 = v409>>(uint(int32(31))%32)&v471 + v409
	v484 = v477 - v481 + int32(1)
	switch v456 {
	default:
		goto L141
	case 1:
		goto L140
	case 2:
		goto L139
	case 3:
		goto L138
	case 4:
		goto L137
	}
L136:
	;
	v500 = int32(0)
	v502 = base.B2i32(base.Ui32(v481) < base.Ui32(v499))
	if base.Ui32(v481) < base.Ui32(v499) {
		goto L143
	} else {
		goto L144
	}
L137:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v444+int32(-17))))
	v499 = v498
	goto L136
L138:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v444+int32(-9))))
	v499 = v495
	goto L136
L139:
	;
	v492 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444+int32(-5)))))
	v499 = v492
	goto L136
L140:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444+int32(-3)))))
	v499 = v489
	goto L136
L141:
	;
	v499 = int32(base.Ui32(v454) >> (uint(int32(3)) % 32))
	goto L136
L142:
	;
	v516 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v444+v510))) = uint8(v516)
	switch v456 {
	default:
		goto L160
	case 1:
		goto L159
	case 2:
		goto L158
	case 3:
		goto L157
	case 4:
		goto L156
	}
L143:
	;
	v503 = v481
	goto L145
L144:
	;
	v503 = v500
	goto L145
L145:
	;
	v504 = v499 - v503
	if base.Ui32(v484) < base.Ui32(v504) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v506 = v484
	goto L148
L147:
	;
	v506 = v504
	goto L148
L148:
	;
	if v477 < v481 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v508 = v500
	goto L151
L150:
	;
	v508 = v506
	goto L151
L151:
	;
	if base.Ui32(v481) < base.Ui32(v499) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v510 = v508
	goto L154
L153:
	;
	v510 = int32(0)
	goto L154
L154:
	;
	if v510 == int32(0) {
		goto L142
	} else {
		goto L155
	}
L155:
	;
	v514 = F_memmove(m, v444, v444+v503, v510)
	mBase = m.M
	goto L142
L156:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v444+int32(-17)))) = base.I64_extend_i32_u(v510)
	goto L128
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v444+int32(-9)))) = v510
	goto L127
L158:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v444+int32(-5)))) = uint16(v510)
	goto L127
L159:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v444+int32(-3)))) = uint8(v510)
	goto L127
L160:
	;
	v519 = v510 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v453))) = uint8(v519)
	goto L127
L161:
	;
	v566 = int32(_a44)
	v568 = *(*int64)(unsafe.Add(mBase, _consts[50]))
	v569 = base.I64_extend_i32_s(v237)
	*(*int64)(unsafe.Add(mBase, _consts[50])) = v568 + v569
	v574 = *(*int64)(unsafe.Add(mBase, _consts[30]))
	*(*int64)(unsafe.Add(mBase, _consts[30])) = v574 + v569
	v577 = int32(0)
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565+int32(-1)))))
	switch v581 & int32(7) {
	case 0:
		goto L171
	case 1:
		goto L170
	case 2:
		goto L169
	case 3:
		goto L168
	case 4:
		goto L167
	default:
		v616 = v577
		v617 = v577
		goto L166
	}
L162:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v551 {
		v561 = v291
		goto L163
	} else {
		goto L164
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, _consts[46])) = int32(0)
	v565 = v561
	goto L161
L164:
	;
	F__serverLog(m, int32(2), int32(_a123), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L35
	} else {
		goto L165
	}
L165:
	;
	v560 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v561 = v560
	goto L163
L166:
	;
	if base.Ui32(int32(3999)) < base.Ui32(v617+v616) {
		goto L172
	} else {
		goto L173
	}
L167:
	;
	v609 = *(*int64)(unsafe.Add(mBase, uint32(v565+int32(-9))))
	v612 = *(*int64)(unsafe.Add(mBase, uint32(v565+int32(-17))))
	v616 = base.I32_wrap_i64(v612)
	v617 = base.I32_wrap_i64(v609 - v612)
	goto L166
L168:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v565+int32(-5))))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v565+int32(-9))))
	v616 = v605
	v617 = v602 - v605
	goto L166
L169:
	;
	v595 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v565+int32(-3)))))
	v598 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v565+int32(-5)))))
	v616 = v598
	v617 = v595 - v598
	goto L166
L170:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565+int32(-2)))))
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565+int32(-3)))))
	v616 = v591
	v617 = v588 - v591
	goto L166
L171:
	;
	v616 = int32(base.Ui32(v581) >> (uint(int32(3)) % 32))
	v617 = v577
	goto L166
L172:
	;
	F_sdsfree(m, v565)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L35
	} else {
		goto L181
	}
L173:
	;
	v624 = v565 + int32(-1)
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624))))
	switch v625 & int32(7) {
	case 0:
		goto L180
	case 1:
		goto L179
	case 2:
		goto L178
	case 3:
		goto L177
	case 4:
		goto L176
	default:
		goto L175
	}
L174:
	;
	v658 = v135
	goto L2
L175:
	;
	v646 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v565))) = uint8(v646)
	goto L174
L176:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v565+int32(-17)))) = int64(0)
	goto L175
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v565+int32(-9)))) = int32(0)
	goto L175
L178:
	;
	v636 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v565+int32(-5)))) = uint16(v636)
	goto L175
L179:
	;
	v632 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v565+int32(-3)))) = uint8(v632)
	goto L175
L180:
	;
	v628 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v624))) = uint8(v628)
	goto L175
L181:
	;
	v651 = F_sdsempty(m)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L35
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, _consts[34])) = v651
	v658 = v135
	goto L2
L183:
	;
	v673 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	switch v673 + int32(-1) {
	case 0:
		goto L188
	case 1:
		goto L187
	default:
		goto L1
	}
L184:
	;
	v669 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	goto L185
L185:
	;
	if v669 != int32(-1) {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	goto L183
L187:
	;
	v735 = int32(_a44)
	v736 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	v738 = *(*int64)(unsafe.Add(mBase, _consts[36]))
	if v736-v738 < int64(1000) {
		goto L1
	} else {
		goto L204
	}
L188:
	;
	v677 = *(*int64)(unsafe.Add(mBase, _consts[44]))
	if base.B2i32(v677 == int64(0)) == int32(0) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v686 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v687 = F_fsync(m, v686)
	mBase = m.M
	if v687 != int32(-1) {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	v683 = F_ustime(m)
	mBase = m.M
	v684 = v683
	goto L189
L191:
	;
	v684 = int64(0)
	goto L189
L192:
	;
	v705 = *(*int64)(unsafe.Add(mBase, _consts[44]))
	if v705 == int64(0) {
		goto L199
	} else {
		goto L200
	}
L193:
	;
	v691 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v691 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	goto L196
L196:
	;
	v695 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v696 = F___strerror_l(m, v695, v695)
	mBase = m.M
	goto L197
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v696
	F__serverLog(m, int32(3), int32(_a124), v13)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L35
	} else {
		goto L198
	}
L198:
	;
	goto L194
L199:
	;
	v723 = int32(_a44)
	v725 = *(*int64)(unsafe.Add(mBase, _consts[30]))
	*(*int64)(unsafe.Add(mBase, _consts[33])) = v725
	v729 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	*(*int64)(unsafe.Add(mBase, _consts[36])) = v729
	v733 = *(*int64)(unsafe.Add(mBase, _consts[40]))
	*(*int64)(unsafe.Add(mBase, _consts[32])) = v733
	goto L1
L200:
	;
	v708 = F_ustime(m)
	mBase = m.M
	v710 = *(*int64)(unsafe.Add(mBase, _consts[44]))
	if v710 == int64(0) {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v713 = v708 - v684
	if v713 < v710*int64(1000) {
		goto L199
	} else {
		goto L202
	}
L202:
	;
	F_latencyAddSample(m, int32(_a125), v713)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L35
	} else {
		goto L203
	}
L203:
	;
	goto L199
L204:
	;
	if v658 != 0 {
		v755 = v736
		goto L205
	} else {
		goto L206
	}
L205:
	;
	*(*int64)(unsafe.Add(mBase, _consts[36])) = v755
	goto L1
L206:
	;
	v742 = int32(_a44)
	v743 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v745 = *(*int64)(unsafe.Add(mBase, _consts[40]))
	F_bioCreateFsyncJob(m, v743, v745, int32(1))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L35
	} else {
		goto L207
	}
L207:
	;
	v749 = int32(_a44)
	v751 = *(*int64)(unsafe.Add(mBase, _consts[30]))
	*(*int64)(unsafe.Add(mBase, _consts[33])) = v751
	v754 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	v755 = v754
	goto L205
}
func F_loadAppendOnlyFiles(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int64
	_ = v218
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int64
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int64
	_ = v319
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v418 int32
	_ = v418
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	if l0 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L1:
	;
	F__serverAssert(m, int32(_a105), int32(_a85), int32(1858))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L30
	} else {
		goto L109
	}
L2:
	;
	m.G0 = v13 + int32(48)
	return v402
L3:
	;
	F_stopLoading(m, base.B2i32(v385 == int32(0))|base.B2i32(v385 == int32(5)))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L30
	} else {
		goto L108
	}
L4:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+20))
	if v259 == int32(0) {
		v371 = v253
		goto L72
	} else {
		goto L73
	}
L5:
	;
	v233 = int32(1)
	if v148 == v233 {
		goto L66
	} else {
		goto L67
	}
L6:
	;
	v218 = F_ustime(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v175
	*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = base.F64_promote_f32(base.F32_div(base.F32_convert_i64_s(v218-v181), float32(1e+06)))
	F__serverLog(m, int32(2), int32(_a161), v13+int32(16))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L30
	} else {
		goto L65
	}
L7:
	;
	F__serverAssert(m, int32(_a162), int32(_a85), int32(1827))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L30
	} else {
		goto L64
	}
L8:
	;
	F__serverAssert(m, int32(_a163), int32(_a85), int32(1809))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L30
	} else {
		goto L63
	}
L9:
	;
	F__serverAssert(m, int32(_a104), int32(_a85), int32(1780))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L30
	} else {
		goto L62
	}
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	v21 = m.G0
	v22 = int32(96)
	v23 = v21 - v22
	m.G0 = v23
	v25 = F_stat(m, v18, v23)
	mBase = m.M
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	m.G0 = v23 + v22
	goto L12
L11:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v133 != 0 {
		goto L38
	} else {
		goto L39
	}
L12:
	;
	if base.B2i32(v25 == int32(0))&base.B2i32(v26&int32(61440) == int32(32768)) == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v43 = m.G0
	v44 = int32(96)
	v45 = v43 - v44
	m.G0 = v45
	v47 = F_stat(m, v40, v45)
	mBase = m.M
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	m.G0 = v45 + v44
	goto L15
L14:
	;
	F_aofUpgradePrepare(m, l0)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L30
	} else {
		goto L35
	}
L15:
	;
	if base.B2i32(v47 == int32(0))&base.B2i32(v48&int32(61440) == int32(16384)) == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v63 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if v62 != 0 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	if v62 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L11
L20:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v68 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v72 == int32(0) {
		v95 = v71
		v96 = v72
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v96-v95&int32(255) != 0 {
		goto L11
	} else {
		goto L29
	}
L22:
	;
	goto L21
L23:
	;
	if v72 != v71&int32(255) {
		v95 = v71
		v96 = v72
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v78 = v66
	v79 = v68
	goto L25
L25:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v83 == int32(0) {
		v95 = v82
		v96 = v83
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v95 = v82
	v96 = v83
	goto L22
L27:
	;
	v86 = int32(1)
	if v83 == v82&int32(255) {
		v78 = v78 + v86
		v79 = v79 + v86
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v102 = F_makePath(m, v101, v68)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	v108 = m.G0
	v109 = int32(96)
	v110 = v108 - v109
	m.G0 = v110
	v112 = F_stat(m, v102, v110)
	mBase = m.M
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	m.G0 = v110 + v109
	goto L32
L32:
	;
	F_sdsfree(m, v102)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	if base.B2i32(v112 == int32(0))&base.B2i32(v113&int32(61440) == int32(32768)) != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	goto L14
L35:
	;
	goto L11
L36:
	;
	if v148 <= int32(0) {
		goto L8
	} else {
		goto L43
	}
L37:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v132)+20))
	v148 = v146 + v144
	v149 = v145
	goto L36
L38:
	;
	v140 = l0 + int32(4)
	v141 = int32(1)
	if v132 == int32(0) {
		v148 = v141
		v149 = v140
		goto L36
	} else {
		goto L42
	}
L39:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+20))
	if v134 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v144 = int32(0)
	v145 = l0 + int32(4)
	goto L37
L41:
	;
	v402 = int32(1)
	goto L2
L42:
	;
	v144 = v141
	v145 = v140
	goto L37
L43:
	;
	v154 = F_getBaseAndIncrAppendOnlyFilesSize(m, l0, v13+int32(44))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	switch v156 {
	case 0:
		goto L45
	case 1:
		goto L46
	default:
		v402 = v156
		goto L2
	}
L45:
	;
	if base.B2i32(v154 == int64(0)) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v402 = int32(4)
	goto L2
L47:
	;
	v163 = int32(0)
	F_startLoading(m, base.I32_wrap_i64(v154), int32(1), v163)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L30
	} else {
		goto L49
	}
L48:
	;
	v402 = int32(2)
	goto L2
L49:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v169 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v169)+16))
	if v172 != int32(98) {
		goto L7
	} else {
		goto L52
	}
L51:
	;
	v253 = int32(0)
	v254 = v163
	v255 = int64(0)
	goto L4
L52:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v175
	goto L53
L53:
	;
	v179 = F_getAppendOnlyFileSize(m, v175, int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L30
	} else {
		goto L54
	}
L54:
	;
	v181 = F_ustime(m)
	mBase = m.M
	v182 = F_loadSingleAppendOnlyFile(m, v175)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L30
	} else {
		goto L56
	}
L55:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v195 <= int32(2) {
		goto L6
	} else {
		goto L61
	}
L56:
	;
	if v182 == int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	if v148 != int32(1) {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	if v182 != int32(5) {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) <= v191 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	goto L6
L61:
	;
	v253 = int32(0)
	v254 = int32(1)
	v255 = v179
	goto L4
L62:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
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
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	goto L5
L66:
	;
	if base.Ui32(v182+int32(-3)) < base.Ui32(int32(2)) {
		v385 = v182
		goto L3
	} else {
		goto L71
	}
L67:
	;
	if v182 != int32(5) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v238 = int32(4)
	v240 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v240 {
		v385 = v238
		goto L3
	} else {
		goto L69
	}
L69:
	;
	F__serverLog(m, int32(3), int32(_a164), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L30
	} else {
		goto L70
	}
L70:
	;
	v385 = v238
	goto L3
L71:
	;
	v253 = v182
	v254 = v233
	v255 = v179
	goto L4
L72:
	;
	v379 = int32(_a44)
	*(*int64)(unsafe.Add(mBase, _consts[75])) = v255
	*(*int64)(unsafe.Add(mBase, _consts[50])) = v154
	v385 = v371
	goto L3
L73:
	;
	v263 = v13 + int32(36)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v264
	goto L74
L74:
	;
	v269 = v13 + int32(36)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	if v271 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v271 == int32(0) {
		v371 = v253
		goto L72
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v271+base.B2i32(v274 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = v280
	goto L76
L78:
	;
	v284 = v271
	v290 = v254
	goto L79
L79:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+16))
	if v295 != int32(105) {
		goto L1
	} else {
		goto L81
	}
L80:
	;
	v385 = v304
	goto L3
L81:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v298
	goto L82
L82:
	;
	v302 = v290 + int32(1)
	v303 = F_ustime(m)
	mBase = m.M
	v304 = F_loadSingleAppendOnlyFile(m, v298)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L30
	} else {
		goto L86
	}
L83:
	;
	if v304 == int32(2) {
		goto L94
	} else {
		goto L95
	}
L84:
	;
	v319 = F_ustime(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v298
	*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = base.F64_promote_f32(base.F32_div(base.F32_convert_i64_s(v319-v303), float32(1e+06)))
	F__serverLog(m, int32(2), int32(_a165), v13)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L30
	} else {
		goto L92
	}
L85:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v316 {
		goto L83
	} else {
		goto L91
	}
L86:
	;
	if v304 == int32(0) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	if v304 != int32(5) {
		goto L83
	} else {
		goto L88
	}
L88:
	;
	if v302 != v148 {
		goto L83
	} else {
		goto L89
	}
L89:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v312 < int32(3) {
		goto L84
	} else {
		goto L90
	}
L90:
	;
	goto L83
L91:
	;
	goto L84
L92:
	;
	goto L83
L93:
	;
	if base.Ui32(int32(-2)) <= base.Ui32(v335+int32(-5)) {
		goto L102
	} else {
		goto L103
	}
L94:
	;
	v335 = int32(0)
	goto L96
L95:
	;
	v335 = v304
	goto L96
L96:
	;
	if v335 != int32(5) {
		goto L93
	} else {
		goto L97
	}
L97:
	;
	if v302 == v148 {
		goto L93
	} else {
		goto L98
	}
L98:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v340 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v385 = int32(4)
	goto L3
L100:
	;
	F__serverLog(m, int32(3), int32(_a164), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L30
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	goto L80
L103:
	;
	v354 = v13 + int32(36)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	if v356 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v356 == int32(0) {
		v371 = v335
		goto L72
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v356+base.B2i32(v359 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v354))) = v365
	goto L105
L107:
	;
	v284 = v356
	v290 = v302
	goto L79
L108:
	;
	v402 = v385
	goto L2
L109:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_rewriteAppendOnlyFileBackground(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
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
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int64
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	v3 = m.G0
	v5 = v3 - int32(336)
	m.G0 = v5
	v7 = int32(-1)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	if v9 != v7 {
		v222 = v7
		m.G0 = v5 + int32(336)
		return v222
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[26]))
		v14 = int32(0)
		v17 = m.G0
		v19 = v17 - int32(96)
		m.G0 = v19
		v23 = F_mkdir(m, v13, int32(493))
		mBase = m.M
		if v23 == v14 {
			v39 = v14
		} else {
			v26 = F___errno_location(m)
			mBase = m.M
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			if v27 != int32(20) {
				v39 = int32(-1)
			} else {
				v30 = F_stat(m, v13, v19)
				mBase = m.M
				if v30 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(54)
					v39 = int32(-1)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
					if v31&int32(61440) == int32(16384) {
						v39 = v14
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(54)
						v39 = int32(-1)
					}
				}
			}
		}
		m.G0 = v19 + int32(96)
		if v39 != int32(-1) {
			*(*int32)(unsafe.Add(mBase, _consts[51])) = int32(-1)
			F_flushAppendOnlyFile(m, int32(1))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				v74 = F_openNewIncrAofForAppend(m)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					if v74 == int32(0) {
						v83 = *(*int32)(unsafe.Add(mBase, _consts[28]))
						if v83 != int32(2) {
							v96 = int32(_a44)
							v98 = *(*int64)(unsafe.Add(mBase, _consts[52]))
							*(*int64)(unsafe.Add(mBase, _consts[52])) = v98 + int64(1)
							v103 = F_serverFork(m, int32(2))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								switch v103 + int32(1) {
								case 0:
									v181 = int32(-1)
									v182 = int32(_a44)
									*(*int32)(unsafe.Add(mBase, _consts[53])) = v181
									v186 = *(*int32)(unsafe.Add(mBase, _consts[6]))
									if int32(3) < v186 {
										v222 = v181
										m.G0 = v5 + int32(336)
										return v222
									} else {
										v190 = *(*int32)(unsafe.Add(mBase, _consts[5]))
										v191 = F___strerror_l(m, v190, v190)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v5)+64)) = v191
										F__serverLog(m, int32(3), int32(_a126), v5+int32(64))
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return int32(0)
										} else {
											v222 = v181
											m.G0 = v5 + int32(336)
											return v222
										}
									}
								case 1:
									v110 = *(*int32)(unsafe.Add(mBase, _consts[54]))
									v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
									v112 = int32(_a127)
									v115 = int32(*(*int8)(unsafe.Add(mBase, _consts[55])))
									if v115 != 0 {
										v117 = F_strchr(m, v111, v115)
										mBase = m.M
										if v117 == int32(0) {
										} else {
											v120 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
											if v120 != 0 {
												v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
												if v121 == int32(0) {
												} else {
													v124 = int32(*(*uint8)(unsafe.Add(mBase, _consts[57])))
													if v124 != 0 {
														v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+2)))
														if v126 == int32(0) {
														} else {
															v129 = int32(*(*uint8)(unsafe.Add(mBase, _consts[58])))
															if v129 != 0 {
																v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+3)))
																if v131 == int32(0) {
																} else {
																	v134 = int32(*(*uint8)(unsafe.Add(mBase, _consts[59])))
																	if v134 != 0 {
																		v136 = F_twoway_strstr(m, v117, v112)
																		mBase = m.M
																	} else {
																		v135 = F_fourbyte_strstr(m, v117, v112)
																		mBase = m.M
																	}
																}
															} else {
																v130 = F_threebyte_strstr(m, v117, v112)
																mBase = m.M
															}
														}
													} else {
														v125 = F_twobyte_strstr(m, v117, v112)
														mBase = m.M
													}
												}
											} else {
											}
										}
									} else {
									}
									v145 = F___syscall_getpid(m)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v5)+48)) = v145
									v153 = F_snprintf(m, v5+int32(80), int32(256), int32(_a128), v5+int32(48))
									mBase = m.M
									v154 = m.ExcPending
									if v154 != 0 {
										return int32(0)
									} else {
										v157 = F_rewriteAppendOnlyFile(m, v5+int32(80))
										mBase = m.M
										v158 = m.ExcPending
										if v158 != 0 {
											return int32(0)
										} else {
											if v157 != 0 {
												F__exit(m, int32(1))
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											} else {
												v160 = *(*int32)(unsafe.Add(mBase, _consts[6]))
												if int32(2) < v160 {
													F_sendChildCowInfo(m, int32(1), int32(_a129))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return int32(0)
													} else {
														F__exit(m, int32(0))
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = v5 + int32(80)
													F__serverLog(m, int32(2), int32(_a130), v5+int32(32))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														F_sendChildCowInfo(m, int32(1), int32(_a129))
														mBase = m.M
														v175 = m.ExcPending
														if v175 != 0 {
															return int32(0)
														} else {
															F__exit(m, int32(0))
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
								default:
									v200 = *(*int32)(unsafe.Add(mBase, _consts[6]))
									if int32(2) < v200 {
										v210 = int32(0)
										v211 = int32(_a44)
										*(*int32)(unsafe.Add(mBase, _consts[29])) = v210
										v216 = F___time(m, v210)
										mBase = m.M
										*(*int64)(unsafe.Add(mBase, _consts[60])) = v216
										v220 = *(*int32)(unsafe.Add(mBase, _consts[61]))
										*(*int32)(unsafe.Add(mBase, _consts[62])) = v220
										v222 = v210
										m.G0 = v5 + int32(336)
										return v222
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v103
										F__serverLog(m, int32(2), int32(_a131), v5+int32(16))
										mBase = m.M
										v209 = m.ExcPending
										if v209 != 0 {
											return int32(0)
										} else {
											v210 = int32(0)
											v211 = int32(_a44)
											*(*int32)(unsafe.Add(mBase, _consts[29])) = v210
											v216 = F___time(m, v210)
											mBase = m.M
											*(*int64)(unsafe.Add(mBase, _consts[60])) = v216
											v220 = *(*int32)(unsafe.Add(mBase, _consts[61]))
											*(*int32)(unsafe.Add(mBase, _consts[62])) = v220
											v222 = v210
											m.G0 = v5 + int32(336)
											return v222
										}
									}
								}
							}
						} else {
							F_bioDrainWorker(m, int32(1))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								v89 = int32(_a44)
								*(*int64)(unsafe.Add(mBase, _consts[31])) = int64(0)
								v94 = *(*int64)(unsafe.Add(mBase, _consts[40]))
								*(*int64)(unsafe.Add(mBase, _consts[32])) = v94
								v96 = int32(_a44)
								v98 = *(*int64)(unsafe.Add(mBase, _consts[52]))
								*(*int64)(unsafe.Add(mBase, _consts[52])) = v98 + int64(1)
								v103 = F_serverFork(m, int32(2))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									switch v103 + int32(1) {
									case 0:
										v181 = int32(-1)
										v182 = int32(_a44)
										*(*int32)(unsafe.Add(mBase, _consts[53])) = v181
										v186 = *(*int32)(unsafe.Add(mBase, _consts[6]))
										if int32(3) < v186 {
											v222 = v181
											m.G0 = v5 + int32(336)
											return v222
										} else {
											v190 = *(*int32)(unsafe.Add(mBase, _consts[5]))
											v191 = F___strerror_l(m, v190, v190)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v5)+64)) = v191
											F__serverLog(m, int32(3), int32(_a126), v5+int32(64))
											mBase = m.M
											v198 = m.ExcPending
											if v198 != 0 {
												return int32(0)
											} else {
												v222 = v181
												m.G0 = v5 + int32(336)
												return v222
											}
										}
									case 1:
										v110 = *(*int32)(unsafe.Add(mBase, _consts[54]))
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
										v112 = int32(_a127)
										v115 = int32(*(*int8)(unsafe.Add(mBase, _consts[55])))
										if v115 != 0 {
											v117 = F_strchr(m, v111, v115)
											mBase = m.M
											if v117 == int32(0) {
											} else {
												v120 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
												if v120 != 0 {
													v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
													if v121 == int32(0) {
													} else {
														v124 = int32(*(*uint8)(unsafe.Add(mBase, _consts[57])))
														if v124 != 0 {
															v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+2)))
															if v126 == int32(0) {
															} else {
																v129 = int32(*(*uint8)(unsafe.Add(mBase, _consts[58])))
																if v129 != 0 {
																	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+3)))
																	if v131 == int32(0) {
																	} else {
																		v134 = int32(*(*uint8)(unsafe.Add(mBase, _consts[59])))
																		if v134 != 0 {
																			v136 = F_twoway_strstr(m, v117, v112)
																			mBase = m.M
																		} else {
																			v135 = F_fourbyte_strstr(m, v117, v112)
																			mBase = m.M
																		}
																	}
																} else {
																	v130 = F_threebyte_strstr(m, v117, v112)
																	mBase = m.M
																}
															}
														} else {
															v125 = F_twobyte_strstr(m, v117, v112)
															mBase = m.M
														}
													}
												} else {
												}
											}
										} else {
										}
										v145 = F___syscall_getpid(m)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v5)+48)) = v145
										v153 = F_snprintf(m, v5+int32(80), int32(256), int32(_a128), v5+int32(48))
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return int32(0)
										} else {
											v157 = F_rewriteAppendOnlyFile(m, v5+int32(80))
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return int32(0)
											} else {
												if v157 != 0 {
													F__exit(m, int32(1))
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												} else {
													v160 = *(*int32)(unsafe.Add(mBase, _consts[6]))
													if int32(2) < v160 {
														F_sendChildCowInfo(m, int32(1), int32(_a129))
														mBase = m.M
														v175 = m.ExcPending
														if v175 != 0 {
															return int32(0)
														} else {
															F__exit(m, int32(0))
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = v5 + int32(80)
														F__serverLog(m, int32(2), int32(_a130), v5+int32(32))
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return int32(0)
														} else {
															F_sendChildCowInfo(m, int32(1), int32(_a129))
															mBase = m.M
															v175 = m.ExcPending
															if v175 != 0 {
																return int32(0)
															} else {
																F__exit(m, int32(0))
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
									default:
										v200 = *(*int32)(unsafe.Add(mBase, _consts[6]))
										if int32(2) < v200 {
											v210 = int32(0)
											v211 = int32(_a44)
											*(*int32)(unsafe.Add(mBase, _consts[29])) = v210
											v216 = F___time(m, v210)
											mBase = m.M
											*(*int64)(unsafe.Add(mBase, _consts[60])) = v216
											v220 = *(*int32)(unsafe.Add(mBase, _consts[61]))
											*(*int32)(unsafe.Add(mBase, _consts[62])) = v220
											v222 = v210
											m.G0 = v5 + int32(336)
											return v222
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v103
											F__serverLog(m, int32(2), int32(_a131), v5+int32(16))
											mBase = m.M
											v209 = m.ExcPending
											if v209 != 0 {
												return int32(0)
											} else {
												v210 = int32(0)
												v211 = int32(_a44)
												*(*int32)(unsafe.Add(mBase, _consts[29])) = v210
												v216 = F___time(m, v210)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, _consts[60])) = v216
												v220 = *(*int32)(unsafe.Add(mBase, _consts[61]))
												*(*int32)(unsafe.Add(mBase, _consts[62])) = v220
												v222 = v210
												m.G0 = v5 + int32(336)
												return v222
											}
										}
									}
								}
							}
						}
					} else {
						v78 = int32(-1)
						*(*int32)(unsafe.Add(mBase, _consts[53])) = v78
						v222 = v78
						m.G0 = v5 + int32(336)
						return v222
					}
				}
			}
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(3) < v47 {
				v64 = int32(-1)
				*(*int32)(unsafe.Add(mBase, _consts[53])) = v64
				v222 = v64
				m.G0 = v5 + int32(336)
				return v222
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, _consts[26]))
				v53 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				v54 = F___strerror_l(m, v53, v53)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v54
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v51
				F__serverLog(m, int32(3), int32(_a110), v5)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v64 = int32(-1)
					*(*int32)(unsafe.Add(mBase, _consts[53])) = v64
					v222 = v64
					m.G0 = v5 + int32(336)
					return v222
				}
			}
		}
	}
}
