package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___intscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v120 int32
	_ = v120
	var v130 int64
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v238 int64
	_ = v238
	var v258 int64
	_ = v258
	var v259 int64
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int64
	_ = v274
	var v276 int32
	_ = v276
	var v284 int64
	_ = v284
	var v285 int64
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v343 int32
	_ = v343
	var v345 int64
	_ = v345
	var v348 int32
	_ = v348
	var v353 int64
	_ = v353
	var v359 int32
	_ = v359
	var v361 int64
	_ = v361
	var v364 int32
	_ = v364
	var v368 int64
	_ = v368
	var v371 int64
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int64
	_ = v385
	var v388 int32
	_ = v388
	var v390 int64
	_ = v390
	var v396 int64
	_ = v396
	var v397 int64
	_ = v397
	var v399 int64
	_ = v399
	var v402 int64
	_ = v402
	var v403 int64
	_ = v403
	var v405 int64
	_ = v405
	var v406 int64
	_ = v406
	var v410 int64
	_ = v410
	var v417 int64
	_ = v417
	var v428 int64
	_ = v428
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int64
	_ = v486
	var v495 int64
	_ = v495
	var v496 int64
	_ = v496
	var v500 int32
	_ = v500
	var v505 int64
	_ = v505
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int64
	_ = v526
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v539 int64
	_ = v539
	var v548 int32
	_ = v548
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v594 int64
	_ = v594
	var v601 int64
	_ = v601
	var v604 int32
	_ = v604
	var v621 int64
	_ = v621
	var v627 int64
	_ = v627
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	if base.Ui32(int32(36)) < base.Ui32(l1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return v627
L2:
	;
	goto L7
L3:
	;
	goto L6
L4:
	;
	if l1 != int32(1) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(28)
	v627 = int64(0)
	goto L1
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v41 == v42 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	switch v52 + int32(-43) {
	case 0, 2:
		goto L17
	default:
		v77 = v52
		v78 = int32(0)
		goto L16
	}
L9:
	;
	goto L14
L10:
	;
	v48 = F___shgetc(m, l0)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v41 + int32(1)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v52 = v47
	goto L9
L12:
	;
	return int64(0)
L13:
	;
	v52 = v48
	goto L9
L14:
	;
	if base.B2i32(v52 == int32(32))|base.B2i32(base.Ui32(v52+int32(-9)) < base.Ui32(int32(5))) != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L8
L16:
	;
	if base.B2i32(l1 != int32(0))&base.B2i32(l1 != int32(16)) != 0 {
		goto L28
	} else {
		goto L29
	}
L17:
	;
	if v52 == int32(45) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v67 = int32(-1)
	goto L20
L19:
	;
	v67 = int32(0)
	goto L20
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v68 == v69 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v75 = F___shgetc(m, l0)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68 + int32(1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v77 = v74
	v78 = v67
	goto L16
L23:
	;
	v77 = v75
	v78 = v67
	goto L16
L24:
	;
	v601 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v601 < int64(0) {
		goto L146
	} else {
		goto L147
	}
L25:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537)+uint32(_consts[1281]))))
	if base.Ui32(v533) <= base.Ui32(v548) {
		v593 = v78
		v594 = v539
		goto L24
	} else {
		goto L134
	}
L26:
	;
	if v292&(v292+int32(-1)) == int32(0) {
		goto L89
	} else {
		goto L90
	}
L27:
	;
	if v187 != int32(10) {
		v292 = v187
		v293 = v188
		goto L26
	} else {
		goto L63
	}
L28:
	;
	if l1 != 0 {
		goto L52
	} else {
		goto L53
	}
L29:
	;
	if v77 != int32(48) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v86 == v87 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v95&int32(-33) != int32(88) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v93 = F___shgetc(m, l0)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v86 + int32(1)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v95 = v92
	goto L31
L34:
	;
	v95 = v93
	goto L31
L35:
	;
	if l1 != 0 {
		v187 = l1
		v188 = v95
		goto L27
	} else {
		goto L51
	}
L36:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v100 == v101 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v110 = int32(16)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+uint32(_consts[1281]))))
	if base.Ui32(v113) < base.Ui32(v110) {
		v292 = v110
		v293 = v109
		goto L26
	} else {
		goto L41
	}
L38:
	;
	v107 = F___shgetc(m, l0)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L12
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v100 + int32(1)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	v109 = v106
	goto L37
L40:
	;
	v109 = v107
	goto L37
L41:
	;
	v116 = int64(0)
	v117 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v117 < v116 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v130 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v130
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = base.I64_extend_i32_s(v135 - v136)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L48
L43:
	;
	if l2 != 0 {
		v627 = v116
		goto L1
	} else {
		goto L46
	}
L44:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v120 + int32(-1)
	if l2 == int32(0) {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v120 + int32(-2)
	v627 = v116
	goto L1
L46:
	;
	goto L42
L47:
	;
	v627 = v130
	goto L1
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v140
	goto L47
L51:
	;
	v292 = int32(8)
	v293 = v95
	goto L26
L52:
	;
	v152 = l1
	goto L54
L53:
	;
	v152 = int32(10)
	goto L54
L54:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+uint32(_consts[1281]))))
	if base.Ui32(v155) < base.Ui32(v152) {
		v187 = v152
		v188 = v77
		goto L27
	} else {
		goto L55
	}
L55:
	;
	v157 = int64(0)
	v158 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v158 < v157 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = base.I64_extend_i32_s(v169 - v170)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L59
L57:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161 + int32(-1)
	goto L56
L58:
	;
	goto L62
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v174
	goto L58
L62:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(28)
	v627 = v157
	goto L1
L63:
	;
	v193 = v188 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v193) {
		v233 = v193
		v238 = int64(0)
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if base.Ui32(int32(9)) < base.Ui32(v233) {
		v593 = v78
		v594 = v238
		goto L24
	} else {
		goto L75
	}
L65:
	;
	v199 = v193
	v202 = int32(0)
	goto L66
L66:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v211 == v212 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v233 = v225
	v238 = base.I64_extend_i32_u(v223)
	goto L64
L68:
	;
	v223 = v202*int32(10) + v199
	v225 = v220 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v225) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v218 = F___shgetc(m, l0)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L12
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v211 + int32(1)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	v220 = v217
	goto L68
L71:
	;
	v220 = v218
	goto L68
L72:
	;
	goto L67
L73:
	;
	if base.Ui32(v223) < base.Ui32(int32(429496729)) {
		v199 = v225
		v202 = v223
		goto L66
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v258 = v238 * int64(10)
	v259 = base.I64_extend_i32_u(v233)
	goto L76
L76:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v264 == v265 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v533 = int32(10)
	v537 = v273
	v539 = v274
	goto L25
L78:
	;
	v274 = v258 + v259
	v276 = v273 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v276) {
		goto L84
	} else {
		goto L85
	}
L79:
	;
	v271 = F___shgetc(m, l0)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L12
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v264 + int32(1)
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	v273 = v270
	goto L78
L81:
	;
	v273 = v271
	goto L78
L82:
	;
	goto L77
L83:
	;
	v284 = v274 * int64(10)
	v285 = base.I64_extend_i32_u(v276)
	if base.Ui64(v284) <= base.Ui64(v285^int64(-1)) {
		v258 = v284
		v259 = v285
		goto L76
	} else {
		goto L88
	}
L84:
	;
	if base.Ui32(v276) <= base.Ui32(int32(9)) {
		goto L82
	} else {
		goto L87
	}
L85:
	;
	if base.Ui64(v274) < base.Ui64(int64(1844674407370955162)) {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v593 = v78
	v594 = v274
	goto L24
L88:
	;
	goto L82
L89:
	;
	v439 = int32(*(*int8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v292*int32(23))>>(uint(int32(5))%32))&int32(7))+uint32(_consts[1282]))))
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+uint32(_consts[1281]))))
	if base.Ui32(v292) <= base.Ui32(v443) {
		v481 = v443
		v484 = v293
		v486 = int64(0)
		goto L113
	} else {
		goto L114
	}
L90:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+uint32(_consts[1281]))))
	if base.Ui32(v292) <= base.Ui32(v302) {
		v343 = v293
		v345 = int64(0)
		v348 = v302
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if base.Ui32(v292) <= base.Ui32(v348) {
		v533 = v292
		v537 = v343
		v539 = v345
		goto L25
	} else {
		goto L102
	}
L92:
	;
	v307 = int32(0)
	v315 = v302
	goto L93
L93:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v319 == v320 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v343 = v328
	v345 = base.I64_extend_i32_u(v330)
	v348 = v333
	goto L91
L95:
	;
	v330 = v315 + v307*v292
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+uint32(_consts[1281]))))
	if base.Ui32(v292) <= base.Ui32(v333) {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v326 = F___shgetc(m, l0)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L12
	} else {
		goto L98
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v319 + int32(1)
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	v328 = v325
	goto L95
L98:
	;
	v328 = v326
	goto L95
L99:
	;
	goto L94
L100:
	;
	if base.Ui32(v330) < base.Ui32(int32(119304647)) {
		v307 = v330
		v315 = v333
		goto L93
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v353 = base.I64_extend_i32_u(v292)
	v359 = v343
	v361 = v345
	v364 = v348
	goto L103
L103:
	;
	v368 = v361 * v353
	v371 = base.I64_extend_i32_u(v364) & int64(255)
	if base.Ui64(v371^int64(-1)) < base.Ui64(v368) {
		v533 = v292
		v537 = v359
		v539 = v361
		goto L25
	} else {
		goto L105
	}
L105:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v375 == v376 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v385 = v368 + v371
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+uint32(_consts[1281]))))
	if base.Ui32(v292) <= base.Ui32(v388) {
		v533 = v292
		v537 = v384
		v539 = v385
		goto L25
	} else {
		goto L110
	}
L107:
	;
	v382 = F___shgetc(m, l0)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L12
	} else {
		goto L109
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v375 + int32(1)
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375))))
	v384 = v381
	goto L106
L109:
	;
	v384 = v382
	goto L106
L110:
	;
	v390 = int64(0)
	v396 = int64(32)
	v397 = int64(base.Ui64(v385) >> (uint(v396) % 64))
	v399 = int64(base.Ui64(v353) >> (uint(v396) % 64))
	v402 = int64(4294967295)
	v403 = v385 & v402
	v405 = v353 & v402
	v406 = v403 * v405
	v410 = int64(base.Ui64(v406)>>(uint(v396)%64)) + v403*v399
	v417 = v410&v402 + v397*v405
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v390*v353 + v390*v385 + v397*v399 + int64(base.Ui64(v410)>>(uint(v396)%64)) + int64(base.Ui64(v417)>>(uint(v396)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v417<<(uint(v396)%64) | v406&v402
	goto L111
L111:
	;
	v428 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	if v428 != int64(0) {
		v533 = v292
		v537 = v384
		v539 = v385
		goto L25
	} else {
		goto L112
	}
L112:
	;
	v359 = v384
	v361 = v385
	v364 = v388
	goto L103
L113:
	;
	if base.Ui32(v292) <= base.Ui32(v481) {
		v533 = v292
		v537 = v484
		v539 = v486
		goto L25
	} else {
		goto L124
	}
L114:
	;
	v448 = v443
	v456 = int32(0)
	goto L115
L115:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v460 == v461 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v481 = v474
	v484 = v469
	v486 = base.I64_extend_i32_u(v471)
	goto L113
L117:
	;
	v470 = v456 << (uint(v439) % 32)
	v471 = v448 | v470
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+uint32(_consts[1281]))))
	if base.Ui32(v292) <= base.Ui32(v474) {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	v467 = F___shgetc(m, l0)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L12
	} else {
		goto L120
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + int32(1)
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	v469 = v466
	goto L117
L120:
	;
	v469 = v467
	goto L117
L121:
	;
	goto L116
L122:
	;
	if base.Ui32(v470) < base.Ui32(int32(134217728)) {
		v448 = v474
		v456 = v471
		goto L115
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	v495 = base.I64_extend_i32_u(v439)
	v496 = int64(base.Ui64(int64(-1)) >> (uint(v495) % 64))
	if base.Ui64(v496) < base.Ui64(v486) {
		v533 = v292
		v537 = v484
		v539 = v486
		goto L25
	} else {
		goto L125
	}
L125:
	;
	v500 = v481
	v505 = v486
	goto L126
L126:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v515 == v516 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v533 = v292
	v537 = v524
	v539 = v526
	goto L25
L128:
	;
	v526 = v505<<(uint(v495)%64) | base.I64_extend_i32_u(v500)&int64(255)
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524)+uint32(_consts[1281]))))
	if base.Ui32(v292) <= base.Ui32(v529) {
		v533 = v292
		v537 = v524
		v539 = v526
		goto L25
	} else {
		goto L132
	}
L129:
	;
	v522 = F___shgetc(m, l0)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L12
	} else {
		goto L131
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v515 + int32(1)
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
	v524 = v521
	goto L128
L131:
	;
	v524 = v522
	goto L128
L132:
	;
	if base.Ui64(v526) <= base.Ui64(v496) {
		v500 = v529
		v505 = v526
		goto L126
	} else {
		goto L133
	}
L133:
	;
	goto L127
L134:
	;
	goto L135
L135:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v564 == v565 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	goto L142
L137:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+uint32(_consts[1281]))))
	if base.Ui32(v576) < base.Ui32(v533) {
		goto L135
	} else {
		goto L141
	}
L138:
	;
	v571 = F___shgetc(m, l0)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L12
	} else {
		goto L140
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v564 + int32(1)
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564))))
	v573 = v570
	goto L137
L140:
	;
	v573 = v571
	goto L137
L141:
	;
	goto L136
L142:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(68)
	if l3&int64(1) == int64(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v586 = v78
	goto L145
L144:
	;
	v586 = int32(0)
	goto L145
L145:
	;
	v593 = v586
	v594 = l3
	goto L24
L146:
	;
	if base.Ui64(v594) < base.Ui64(l3) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604 + int32(-1)
	goto L146
L148:
	;
	v621 = base.I64_extend_i32_s(v593)
	v627 = v594 ^ v621 - v621
	goto L1
L149:
	;
	if base.I32_wrap_i64(l3)&int32(1) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if base.Ui64(v594) <= base.Ui64(l3) {
		goto L148
	} else {
		goto L154
	}
L151:
	;
	if v593 != 0 {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(68)
	v627 = l3 + int64(-1)
	goto L1
L154:
	;
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(68)
	v627 = l3
	goto L1
}
func F___isspace_3(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 == int32(32)) | base.B2i32(base.Ui32(l0+int32(-9)) < base.Ui32(int32(5)))
}
func F___isspace_5(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 == int32(32)) | base.B2i32(base.Ui32(l0+int32(-9)) < base.Ui32(int32(5)))
}
func F_incrCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_incrDecrCommand(m, l0, int64(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_incrCommandStatsOnError(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	v4 = int32(0)
	v6 = *(*int64)(unsafe.Add(mBase, _consts[844]))
	if l0 == v4 {
		v32 = v4
	} else {
		v10 = int32(0)
		v12 = *(*int64)(unsafe.Add(mBase, _consts[845]))
		if v6 <= v12 {
			v32 = v10
		} else {
			if l1&int32(1) == int32(0) {
				v19 = int32(0)
				if l1&int32(2) == v19 {
					v32 = v19
				} else {
					v25 = int32(128)
					v26 = l0 + v25
					v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
					*(*int64)(unsafe.Add(mBase, uint32(v26))) = v27 + int64(1)
					v32 = int32(1)
				}
			} else {
				v25 = int32(120)
				v26 = l0 + v25
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
				*(*int64)(unsafe.Add(mBase, uint32(v26))) = v27 + int64(1)
				v32 = int32(1)
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, _consts[845])) = v6
	return v32
}
func F_incrRefCount(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(int32(-17)) < base.Ui32(v3) {
		if base.Ui32(int32(-9)) < base.Ui32(v3) {
			return
		} else {
			F__serverPanic_1(m, int32(_a838), int32(622), int32(_a847), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3 + int32(8)
		return
	}
}
func F_incrbyfloatCommand(m *base.Module, l0 int32) {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v88 int64
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	v8 = m.G0
	v10 = v8 - int32(64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = F_lookupKeyWrite(m, v12, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v18 = F_checkType(m, l0, v15, int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			if v18 != 0 {
				m.G0 = v10 + int32(64)
				return
			} else {
				v23 = F_getLongDoubleFromObjectOrReply(m, l0, v15, v10+int32(32), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					if v23 != 0 {
						m.G0 = v10 + int32(64)
						return
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
						v30 = F_getLongDoubleFromObjectOrReply(m, l0, v26, v10+int32(48), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							if v30 != 0 {
								m.G0 = v10 + int32(64)
								return
							} else {
								v34 = *(*int64)(unsafe.Add(mBase, uint32(v10)+48))
								v39 = *(*int64)(unsafe.Add(mBase, uint32(v10+int32(56))))
								v40 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
								v44 = v10 + int32(40)
								v45 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
								F___addtf3(m, v10+int32(8), v34, v39, v40, v45)
								mBase = m.M
								v51 = *(*int64)(unsafe.Add(mBase, uint32(v10+int32(16))))
								*(*int64)(unsafe.Add(mBase, uint32(v44))) = v51
								v53 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v53
								v59 = v51 & int64(281474976710655)
								v63 = int32(32767)
								v64 = base.I32_wrap_i64(int64(base.Ui64(v51)>>(uint(int64(48))%64))) & v63
								if v64 == v63 {
									v77 = base.B2i32(v59|v53 == int64(0))
									v79 = v77
								} else {
									if v64 != 0 {
										v77 = int32(4)
										v79 = v77
									} else {
										if v59|v53 == int64(0) {
											v73 = int32(2)
										} else {
											v73 = int32(3)
										}
										v79 = v73
									}
								}
								if v79 == int32(0) {
									F_addReplyError(m, l0, int32(_a1595))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return
									} else {
										m.G0 = v10 + int32(64)
										return
									}
								} else {
									v82 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
									v83 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
									v88 = v83 & int64(281474976710655)
									v92 = int32(32767)
									v93 = base.I32_wrap_i64(int64(base.Ui64(v83)>>(uint(int64(48))%64))) & v92
									if v93 == v92 {
										v106 = base.B2i32(v88|v82 == int64(0))
										v108 = v106
									} else {
										if v93 != 0 {
											v106 = int32(4)
											v108 = v106
										} else {
											if v88|v82 == int64(0) {
												v102 = int32(2)
											} else {
												v102 = int32(3)
											}
											v108 = v102
										}
									}
									if v108 != int32(1) {
										v114 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
										v117 = *(*int64)(unsafe.Add(mBase, uint32(v10+int32(40))))
										v119 = F_createStringObjectFromLongDouble(m, v114, v117, int32(1))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v119
											v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
											if v15 == int32(0) {
												F_dbAdd(m, v122, v124, v10+int32(28))
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return
												} else {
													v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
													v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
													F_signalModifiedKey(m, l0, v135, v137)
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
														return
													} else {
														v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
														v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
														v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+28))
														F_notifyKeyspaceEvent(m, int32(8), int32(_a1596), v143, v145)
														mBase = m.M
														v147 = m.ExcPending
														if v147 != 0 {
															return
														} else {
															v148 = int32(_a69)
															v150 = *(*int64)(unsafe.Add(mBase, _consts[60]))
															*(*int64)(unsafe.Add(mBase, _consts[60])) = v150 + int64(1)
															v154 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
															F_addReplyBulk(m, l0, v154)
															mBase = m.M
															v156 = m.ExcPending
															if v156 != 0 {
																return
															} else {
																v159 = *(*int32)(unsafe.Add(mBase, _consts[751]))
																F_rewriteClientCommandArgument(m, l0, int32(0), v159)
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return
																} else {
																	v163 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
																	F_rewriteClientCommandArgument(m, l0, int32(2), v163)
																	mBase = m.M
																	v165 = m.ExcPending
																	if v165 != 0 {
																		return
																	} else {
																		v168 = *(*int32)(unsafe.Add(mBase, _consts[766]))
																		F_rewriteClientCommandArgument(m, l0, int32(3), v168)
																		mBase = m.M
																		v170 = m.ExcPending
																		if v170 != 0 {
																			return
																		} else {
																			m.G0 = v10 + int32(64)
																			return
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												F_dbReplaceValue(m, v122, v124, v10+int32(28))
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return
												} else {
													v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
													v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
													F_signalModifiedKey(m, l0, v135, v137)
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
														return
													} else {
														v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
														v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
														v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+28))
														F_notifyKeyspaceEvent(m, int32(8), int32(_a1596), v143, v145)
														mBase = m.M
														v147 = m.ExcPending
														if v147 != 0 {
															return
														} else {
															v148 = int32(_a69)
															v150 = *(*int64)(unsafe.Add(mBase, _consts[60]))
															*(*int64)(unsafe.Add(mBase, _consts[60])) = v150 + int64(1)
															v154 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
															F_addReplyBulk(m, l0, v154)
															mBase = m.M
															v156 = m.ExcPending
															if v156 != 0 {
																return
															} else {
																v159 = *(*int32)(unsafe.Add(mBase, _consts[751]))
																F_rewriteClientCommandArgument(m, l0, int32(0), v159)
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return
																} else {
																	v163 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
																	F_rewriteClientCommandArgument(m, l0, int32(2), v163)
																	mBase = m.M
																	v165 = m.ExcPending
																	if v165 != 0 {
																		return
																	} else {
																		v168 = *(*int32)(unsafe.Add(mBase, _consts[766]))
																		F_rewriteClientCommandArgument(m, l0, int32(3), v168)
																		mBase = m.M
																		v170 = m.ExcPending
																		if v170 != 0 {
																			return
																		} else {
																			m.G0 = v10 + int32(64)
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
									} else {
										F_addReplyError(m, l0, int32(_a1595))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return
										} else {
											m.G0 = v10 + int32(64)
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
func F_incrementErrorCount(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
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
	var v219 int64
	_ = v219
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[847]))
	v12 = v7 + int32(12)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if l1 == v3 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return
L2:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v218)))
	*(*int64)(unsafe.Add(mBase, uint32(v218))) = v219 + int64(1)
	goto L1
L3:
	;
	if v206 != 0 {
		goto L2
	} else {
		goto L41
	}
L4:
	;
	if v165 != l1 {
		v206 = v3
		goto L31
	} else {
		goto L32
	}
L5:
	;
	v156 = int32(0)
	v162 = v21
	v163 = v22
	v165 = v156
	v169 = v156
	goto L4
L6:
	;
	if base.Ui32(v22) < base.Ui32(int32(8)) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v33 = v21
	v34 = v22
	v36 = int32(0)
	goto L9
L8:
	;
	v162 = v146
	v163 = v147
	v165 = v149
	v169 = base.B2i32(v152 != int32(0))
	goto L4
L9:
	;
	v42 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
	v43 = int32(4)
	v44 = v33 + v43
	if v34&v43 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v146 = v137
	v147 = v138
	v149 = v122
	v152 = v127
	goto L8
L11:
	;
	v127 = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v44+v42+(v127-v42)&int32(3)+v115<<(uint(int32(2))%32))))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	if base.Ui32(v138) < base.Ui32(int32(8)) {
		v146 = v137
		v147 = v138
		v149 = v122
		v152 = v127
		goto L8
	} else {
		goto L29
	}
L12:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v36))))
	v93 = int32(0)
	goto L23
L13:
	;
	v49 = int32(0)
	if base.Ui32(l1) <= base.Ui32(v36) {
		v82 = v36
		v85 = v49
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v85 == v42 {
		v115 = v49
		v122 = v82
		goto L11
	} else {
		goto L21
	}
L15:
	;
	v59 = v36
	v62 = v49
	goto L16
L16:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v62))))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v59))))
	if v65 != v67 {
		v82 = v59
		v85 = v62
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v82 = v70
	v85 = v72
	goto L14
L18:
	;
	v69 = int32(1)
	v70 = v59 + v69
	v72 = v62 + v69
	if base.Ui32(v42) <= base.Ui32(v72) {
		v82 = v70
		v85 = v72
		goto L14
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(v70) < base.Ui32(l1) {
		v59 = v70
		v62 = v72
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v146 = v33
	v147 = v34
	v149 = v82
	v152 = v85
	goto L8
L22:
	;
	if v93 != v42 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v93))))
	if v106 == v90&int32(255) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v108 = int32(1)
	v110 = v93 + v108
	if v110 != v42 {
		v93 = v110
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v162 = v33
	v163 = v34
	v165 = v36
	v169 = v108
	goto L4
L27:
	;
	v115 = v93
	v122 = v36 + int32(1)
	goto L11
L28:
	;
	v146 = v33
	v147 = v34
	v149 = v36
	v152 = v42
	goto L8
L29:
	;
	if base.Ui32(v122) < base.Ui32(l1) {
		v33 = v137
		v34 = v138
		v36 = v122
		goto L9
	} else {
		goto L30
	}
L30:
	;
	goto L10
L31:
	;
	goto L3
L32:
	;
	v171 = int32(0)
	if v163&int32(1) == v171 {
		v206 = v171
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v177 = v163 & int32(4)
	if v169&base.B2i32(v177 != int32(0)) != 0 {
		v206 = v171
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v181 = int32(1)
	if v12 == int32(0) {
		v206 = v181
		goto L31
	} else {
		goto L35
	}
L35:
	;
	if v163&int32(2) != 0 {
		v203 = int32(0)
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v203
	v206 = v181
	goto L31
L37:
	;
	v187 = int32(3)
	v188 = int32(base.Ui32(v163) >> (uint(v187) % 32))
	if v177 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v198 = int32(4)
	goto L40
L39:
	;
	v198 = v188 << (uint(int32(2)) % 32)
	goto L40
L40:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v162+v188+(int32(0)-v188)&v187+v198+int32(4))))
	v203 = v202
	goto L36
L41:
	;
	v209 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	return
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v209))) = int64(1)
	v213 = int32(0)
	v214 = *(*int32)(unsafe.Add(mBase, _consts[847]))
	v216 = F_raxInsert(m, v214, l0, l1, v209, v213)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	goto L1
}
func F_inet_pton(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
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
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v493 int32
	_ = v493
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l0 == int32(10) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v493
L2:
	;
	v493 = int32(0)
	goto L1
L3:
	;
	goto L118
L4:
	;
	v493 = int32(1)
	goto L1
L5:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v91 != int32(58) {
		v99 = l1
		goto L26
	} else {
		goto L27
	}
L6:
	;
	if l0 != int32(2) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v20 = l1
	v23 = int32(0)
	goto L8
L8:
	;
	v28 = int32(0)
	v30 = v28
	v35 = v28
	goto L13
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2+v23))) = uint8(v68)
	v72 = v20 + v69
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v23 != int32(3) {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	v61 = int32(0)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v62 == int32(48) {
		v493 = v61
		goto L1
	} else {
		goto L19
	}
L12:
	;
	switch v30 {
	case 0:
		v493 = v30
		goto L1
	case 1:
		goto L17
	default:
		v59 = v35
		v60 = v30
		goto L11
	}
L13:
	;
	v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20+v30))))
	if base.Ui32(int32(9)) < base.Ui32(v40+int32(-48)) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v49 = v35*int32(10) + v40 + int32(-48)
	v50 = int32(3)
	v52 = v30 + int32(1)
	if v52 != v50 {
		v30 = v52
		v35 = v49
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v59 = v49
	v60 = v50
	goto L11
L17:
	;
	if v35 <= int32(255) {
		v68 = v35
		v69 = int32(1)
		goto L10
	} else {
		goto L18
	}
L18:
	;
	goto L2
L19:
	;
	if int32(255) < v59 {
		v493 = v61
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v68 = v59
	v69 = v60
	goto L10
L21:
	;
	v80 = int32(0)
	if v73&int32(255) != int32(46) {
		v493 = v80
		goto L1
	} else {
		goto L24
	}
L22:
	;
	if v73&int32(255) == int32(0) {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v85 = int32(1)
	if base.Ui32(v23) < base.Ui32(int32(3)) {
		v20 = v72 + v85
		v23 = v23 + v85
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v493 = v80
	goto L1
L26:
	;
	v103 = v99
	v108 = v91
	v109 = int32(0)
	v110 = int32(-1)
	goto L29
L27:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v94 != int32(58) {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v99 = l1 + int32(1)
	goto L26
L29:
	;
	v111 = int32(0)
	if v108&int32(255) != int32(58) {
		v134 = v111
		v135 = v111
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v103 = v458
	v108 = v464
	v109 = v109 + int32(1)
	v110 = v466
	goto L29
L32:
	;
	v454 = v103 + v179
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454)+1)))
	v458 = v454 + int32(1)
	v464 = v457
	v466 = v110
	goto L31
L33:
	;
	v223 = int32(0)
	if v222 < v223 {
		goto L63
	} else {
		goto L64
	}
L34:
	;
	v136 = v134
	v141 = v135
	goto L42
L35:
	;
	v117 = int32(0)
	if int32(-1) < v110 {
		v134 = v117
		v135 = v117
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v121 = int32(1)
	v127 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12|v109&int32(7)<<(uint(v121)%32)))) = uint16(v127)
	v130 = v103 + v121
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	if v131 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if v109 == int32(7) {
		goto L2
	} else {
		goto L39
	}
L38:
	;
	v215 = v130
	v218 = v121
	v221 = v109
	v222 = v109
	goto L33
L39:
	;
	v458 = v130
	v464 = v131
	v466 = v109
	goto L31
L40:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v12|v109&int32(7)<<(uint(int32(1))%32)))) = uint16(v177)
	v187 = v178 & int32(255)
	if v187 != 0 {
		goto L53
	} else {
		goto L54
	}
L41:
	;
	if v141 == int32(0) {
		goto L2
	} else {
		goto L52
	}
L42:
	;
	v146 = int32(*(*int8)(unsafe.Add(mBase, uint32(v103+v141))))
	v149 = v146 + int32(-48)
	if base.Ui32(v149) < base.Ui32(int32(10)) {
		v163 = v149
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
	v177 = v168
	v178 = v173
	v179 = int32(4)
	goto L40
L44:
	;
	if v163 < int32(0) {
		goto L41
	} else {
		goto L50
	}
L45:
	;
	goto L44
L46:
	;
	v153 = v146 | int32(32)
	if base.Ui32(v153+int32(-97)) < base.Ui32(int32(6)) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v161 = v153 + int32(-87)
	goto L49
L48:
	;
	v161 = int32(-1)
	goto L49
L49:
	;
	v163 = v161
	goto L45
L50:
	;
	v166 = int32(4)
	v168 = v163 + v136<<(uint(v166)%32)
	v170 = v141 + int32(1)
	if v170 != v166 {
		v136 = v168
		v141 = v170
		goto L42
	} else {
		goto L51
	}
L51:
	;
	goto L43
L52:
	;
	v177 = v136
	v178 = v146
	v179 = v141
	goto L40
L53:
	;
	v193 = int32(0)
	if v109 == int32(7) {
		v493 = v193
		goto L1
	} else {
		goto L57
	}
L54:
	;
	v188 = int32(1)
	if int32(-1) < v110 {
		v215 = v103
		v218 = v188
		v221 = v109
		v222 = v110
		goto L33
	} else {
		goto L55
	}
L55:
	;
	if v109 != int32(7) {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v215 = v103
	v218 = v188
	v221 = v109
	v222 = v110
	goto L33
L57:
	;
	if v187 == int32(58) {
		goto L32
	} else {
		goto L58
	}
L58:
	;
	if v187 != int32(46) {
		v493 = v193
		goto L1
	} else {
		goto L59
	}
L59:
	;
	if base.Ui32(int32(5)) < base.Ui32(v109) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v204 = int32(0)
	v205 = int32(1)
	v206 = v109 + v205
	*(*uint16)(unsafe.Add(mBase, uint32(v12|v206&int32(7)<<(uint(v205)%32)))) = uint16(v204)
	v215 = v103
	v218 = v204
	v221 = v206
	v222 = v110
	goto L33
L61:
	;
	if v110 < int32(0) {
		v493 = v193
		goto L1
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v423 = v223
	v425 = l2
	goto L113
L64:
	;
	v226 = int32(1)
	v228 = v12 + v222<<(uint(v226)%32)
	v233 = v228 - v221<<(uint(v226)%32) + int32(14)
	v238 = (v221-v222)<<(uint(v226)%32) + int32(2)
	if v233 == v228 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if int32(6) < v221 {
		goto L63
	} else {
		goto L106
	}
L66:
	;
	goto L65
L67:
	;
	v242 = v238 + v233
	if base.Ui32(int32(0)-v238<<(uint(int32(1))%32)) < base.Ui32(v228-v242) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v252 = (v228 ^ v233) & int32(3)
	if base.Ui32(v228) <= base.Ui32(v233) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v249 = F___memcpy(m, v233, v228, v238)
	mBase = m.M
	goto L65
L70:
	;
	if v358 == int32(0) {
		goto L66
	} else {
		goto L102
	}
L71:
	;
	if base.Ui32(v336) <= base.Ui32(int32(3)) {
		v357 = v335
		v358 = v336
		v359 = v337
		goto L70
	} else {
		goto L98
	}
L72:
	;
	if v252 != 0 {
		v318 = v238
		goto L82
	} else {
		goto L83
	}
L73:
	;
	if v252 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	if v233&int32(3) != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v357 = v228
	v358 = v238
	v359 = v233
	goto L70
L76:
	;
	v259 = v228
	v260 = v238
	v261 = v233
	goto L78
L77:
	;
	v335 = v228
	v336 = v238
	v337 = v233
	goto L71
L78:
	;
	if v260 == int32(0) {
		goto L66
	} else {
		goto L80
	}
L80:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v265)
	v267 = int32(1)
	v268 = v259 + v267
	v270 = v260 + int32(-1)
	v272 = v261 + v267
	if v272&int32(3) == int32(0) {
		v335 = v268
		v336 = v270
		v337 = v272
		goto L71
	} else {
		goto L81
	}
L81:
	;
	v259 = v268
	v260 = v270
	v261 = v272
	goto L78
L82:
	;
	if v318 == int32(0) {
		goto L66
	} else {
		goto L94
	}
L83:
	;
	if v242&int32(3) == int32(0) {
		v298 = v238
		goto L84
	} else {
		goto L85
	}
L84:
	;
	if base.Ui32(v298) <= base.Ui32(int32(3)) {
		v318 = v298
		goto L82
	} else {
		goto L90
	}
L85:
	;
	v283 = v238
	goto L86
L86:
	;
	if v283 == int32(0) {
		goto L66
	} else {
		goto L88
	}
L87:
	;
	v298 = v289
	goto L84
L88:
	;
	v289 = v283 + int32(-1)
	v290 = v233 + v289
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228+v289))))
	*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v292)
	if v290&int32(3) != 0 {
		v283 = v289
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v305 = v298
	goto L91
L91:
	;
	v309 = v305 + int32(-4)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v228+v309)))
	*(*int32)(unsafe.Add(mBase, uint32(v233+v309))) = v312
	if base.Ui32(int32(3)) < base.Ui32(v309) {
		v305 = v309
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v318 = v309
	goto L82
L93:
	;
	goto L92
L94:
	;
	v325 = v318
	goto L95
L95:
	;
	v329 = v325 + int32(-1)
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228+v329))))
	*(*uint8)(unsafe.Add(mBase, uint32(v233+v329))) = uint8(v332)
	if v329 != 0 {
		v325 = v329
		goto L95
	} else {
		goto L97
	}
L97:
	;
	goto L66
L98:
	;
	v342 = v335
	v343 = v336
	v344 = v337
	goto L99
L99:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = v346
	v348 = int32(4)
	v349 = v342 + v348
	v351 = v344 + v348
	v353 = v343 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v353) {
		v342 = v349
		v343 = v353
		v344 = v351
		goto L99
	} else {
		goto L101
	}
L100:
	;
	v357 = v349
	v358 = v353
	v359 = v351
	goto L70
L101:
	;
	goto L100
L102:
	;
	v364 = v357
	v365 = v358
	v366 = v359
	goto L103
L103:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
	*(*uint8)(unsafe.Add(mBase, uint32(v366))) = uint8(v368)
	v370 = int32(1)
	v375 = v365 + int32(-1)
	if v375 != 0 {
		v364 = v364 + v370
		v365 = v375
		v366 = v366 + v370
		goto L103
	} else {
		goto L105
	}
L104:
	;
	goto L66
L105:
	;
	goto L104
L106:
	;
	v390 = int32(7) - v221
	v391 = int32(1)
	if v391 < v390 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v394 = v390
	goto L109
L108:
	;
	v394 = v391
	goto L109
L109:
	;
	v401 = int32(0)
	goto L110
L110:
	;
	v406 = int32(1)
	v409 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12+(v401+v222)<<(uint(v406)%32)))) = uint16(v409)
	v412 = v401 + v406
	if v412 != v394 {
		v401 = v412
		goto L110
	} else {
		goto L112
	}
L111:
	;
	goto L63
L112:
	;
	goto L111
L113:
	;
	v432 = int32(1)
	v435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+v423<<(uint(v432)%32)))))
	v436 = int32(8)
	v440 = v435<<(uint(v436)%32) | int32(base.Ui32(v435)>>(uint(v436)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v425))) = uint16(v440)
	v445 = v423 + v432
	if v445 != v436 {
		v423 = v445
		v425 = v425 + int32(2)
		goto L113
	} else {
		goto L115
	}
L114:
	;
	if v218 != 0 {
		goto L4
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	v451 = F_inet_pton(m, int32(2), v215, v425+int32(-2))
	mBase = m.M
	if v451 < int32(1) {
		goto L2
	} else {
		goto L117
	}
L117:
	;
	goto L4
L118:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(5)
	v493 = int32(-1)
	goto L1
}
func F_infoCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	if v13 == v2 {
		v18 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v18
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v33 = F_genInfoSectionDict(m, v22+int32(4), v25+int32(-1), v18, v9+int32(12), v9+int32(8))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v37 = F_genValkeyInfoString(m, v33, v35, v36)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-1)))))
				switch v41 & int32(7) {
				case 0:
					v58 = int32(base.Ui32(v41) >> (uint(int32(3)) % 32))
				case 1:
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-3)))))
					v58 = v48
				case 2:
					v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(-5)))))
					v58 = v51
				case 3:
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-9))))
					v58 = v54
				case 4:
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-17))))
					v58 = v57
				default:
					v58 = v2
				}
				F_addReplyVerbatim(m, l0, v37, v58, int32(_a683))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					F_sdsfree(m, v37)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, _consts[1052]))
						if v33 == v65 {
							m.G0 = v9 + int32(16)
							return
						} else {
							F_dictRelease(m, v33)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_sentinelInfoCommand(m, l0)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			m.G0 = v9 + int32(16)
			return
		}
	}
}
func F_initListeners(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int64
	_ = v137
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	v1 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	if v13 == v1 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	v180 = int32(0)
	v184 = v180
	v185 = v180
	goto L42
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(_a902)
	F__serverPanic_1(m, int32(_a1240), int32(3150), int32(_a1344), v10+int32(32))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L8
	} else {
		goto L41
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(_a1345)
	F__serverPanic_1(m, int32(_a1240), int32(3140), int32(_a1344), v10+int32(48))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L8
	} else {
		goto L40
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = int32(_a1346)
	F__serverPanic_1(m, int32(_a1240), int32(3131), int32(_a1344), v10+int32(64))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L8
	} else {
		goto L39
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = int32(_a1347)
	F__serverPanic_1(m, int32(_a1240), int32(3109), int32(_a1344), v10+int32(80))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L8
	} else {
		goto L38
	}
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[105]))
	if v36 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v17 = F_connectionByType(m, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v17 == int32(0) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v21 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[824])) = int32(_a452)
	*(*int32)(unsafe.Add(mBase, _consts[825])) = v17
	v28 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	*(*int32)(unsafe.Add(mBase, _consts[826])) = v28
	v32 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	*(*int32)(unsafe.Add(mBase, _consts[827])) = v32
	goto L6
L11:
	;
	v97 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	if v98 == v97 {
		goto L31
	} else {
		goto L32
	}
L12:
	;
	v43 = F_connectionTypeTls(m)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L17
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v38 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v39 = int32(0)
	v40 = *(*int32)(unsafe.Add(mBase, _consts[89]))
	if v40 == v39 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v59 = m.T0[v58].(func(*base.Module, int32, int32) int32)(m, int32(_a450), int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L23
	}
L17:
	;
	if v43 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v46 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	F__serverLog(m, int32(3), int32(_a1348), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v74 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, _consts[105]))
	if v75 == v74 {
		goto L11
	} else {
		goto L28
	}
L23:
	;
	if v59 != int32(-1) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v64 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	F__serverLog(m, int32(3), int32(_a1349), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v79 = F_connectionByType(m, int32(2))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	if v79 == int32(0) {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[828])) = int32(_a452)
	*(*int32)(unsafe.Add(mBase, _consts[829])) = v79
	v90 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	*(*int32)(unsafe.Add(mBase, _consts[830])) = v90
	v94 = *(*int32)(unsafe.Add(mBase, _consts[105]))
	*(*int32)(unsafe.Add(mBase, _consts[831])) = v94
	goto L11
L31:
	;
	v118 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	if v119 == v118 {
		goto L1
	} else {
		goto L35
	}
L32:
	;
	v102 = F_connectionByType(m, int32(1))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	if v102 == int32(0) {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v106 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[832])) = int32(_a1350)
	*(*int32)(unsafe.Add(mBase, _consts[833])) = v102
	*(*int32)(unsafe.Add(mBase, _consts[834])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[835])) = int32(_a1351)
	goto L31
L35:
	;
	v123 = F_connectionByType(m, int32(3))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	if v123 == int32(0) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v127 = int32(0)
	v128 = int32(_a462)
	*(*int32)(unsafe.Add(mBase, _consts[836])) = v128
	*(*int32)(unsafe.Add(mBase, _consts[837])) = v128
	*(*int32)(unsafe.Add(mBase, _consts[838])) = v123
	v137 = *(*int64)(unsafe.Add(mBase, _consts[205]))
	*(*int64)(unsafe.Add(mBase, _consts[839])) = v137
	goto L1
L38:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v190 = v184 * int32(88)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[825])))
	if v193 == int32(0) {
		v327 = v185
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v327 != 0 {
		goto L77
	} else {
		goto L78
	}
L44:
	;
	v332 = v184 + int32(1)
	if v332 != int32(4) {
		v184 = v332
		v185 = v327
		goto L42
	} else {
		goto L76
	}
L45:
	;
	v197 = v190 + int32(_a1337)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v193)+32))
	v199 = m.T0[v198].(func(*base.Module, int32) int32)(m, v197)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L47
	}
L46:
	;
	v232 = int32(0)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[825])))
	if v234 == v232 {
		v238 = v232
		goto L55
	} else {
		goto L56
	}
L47:
	;
	if v199 != int32(-1) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v204 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[827])))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[825])))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v211 = m.T0[v210].(func(*base.Module) int32)(m)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L8
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v207
	F__serverLog(m, int32(3), int32(_a1352), v10)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L8
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(int32(3)) < base.Ui32(v211) {
		v220 = int32(_a1353)
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v211<<(uint(int32(2))%32))+uint32(_consts[840])))
	v220 = v219
	goto L51
L54:
	;
	goto L49
L55:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[841])))
	if v239 < int32(1) {
		v320 = v239
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v234)+20))
	v238 = v237
	goto L55
L57:
	;
	v327 = v320 + v185
	goto L44
L58:
	;
	v243 = v232
	goto L59
L59:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v197+v243<<(uint(int32(2))%32))))
	v256 = F_aeCreateFileEvent(m, v250, v254, int32(1), v238, v197)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L8
	} else {
		goto L62
	}
L60:
	;
	v320 = v314
	goto L57
L61:
	;
	v313 = v243 + int32(1)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[841])))
	if v313 < v314 {
		v243 = v313
		goto L59
	} else {
		goto L75
	}
L62:
	;
	if v256 != int32(-1) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	if v243 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[825])))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v291 = m.T0[v290].(func(*base.Module) int32)(m)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L8
	} else {
		goto L70
	}
L65:
	;
	v263 = v243
	goto L66
L66:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v272 = v263 + int32(-1)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v197+v272<<(uint(int32(2))%32))))
	F_aeDeleteFileEvent(m, v270, v276, int32(1))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L8
	} else {
		goto L68
	}
L67:
	;
	goto L64
L68:
	;
	if int32(1) < v263 {
		v263 = v272
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	if base.Ui32(int32(3)) < base.Ui32(v291) {
		v302 = int32(_a1353)
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v302
	F__serverPanic_1(m, int32(_a1240), int32(3172), int32(_a1354), v10+int32(16))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L8
	} else {
		goto L74
	}
L72:
	;
	goto L71
L73:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v291<<(uint(int32(2))%32))+uint32(_consts[840])))
	v302 = v301
	goto L72
L74:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	goto L60
L76:
	;
	goto L43
L77:
	;
	m.G0 = v10 + int32(96)
	return
L78:
	;
	v336 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v336 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F__serverLog(m, int32(3), int32(_a1355), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	goto L79
}
func F_initializeRandomSeed(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1112])))
	if v11 != 0 {
		F__serverAssert(m, int32(_a1650), int32(_a1651), int32(1097))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v14 = F_fopen(m, int32(_a1652), int32(_a178))
		mBase = m.M
		if v14 == int32(0) {
			v29 = int32(0)
			for {
				v32 = F___gettimeofday(m, v8, int32(0))
				mBase = m.M
				v33 = F___syscall_getpid(m)
				mBase = m.M
				v36 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
				v40 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)))
				v41 = v36 ^ base.I64_extend_i32_u(v14) ^ base.I64_extend_i32_u(v33) ^ v40
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1113]))) = uint8(v41)
				v44 = v29 + int32(1)
				if v44 != int32(64) {
					v29 = v44
					continue
				} else {
					break
				}
				break
			}
			if v14 == int32(0) {
				m.G0 = v8 + int32(16)
				return
			} else {
				v63 = F_fclose(m, v14)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		} else {
			v20 = F_fread(m, int32(_a1653), int32(64), int32(1), v14)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				if v20 == int32(1) {
					v56 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _consts[1112])) = uint8(v56)
					v63 = F_fclose(m, v14)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				} else {
					v29 = int32(0)
					for {
						v32 = F___gettimeofday(m, v8, int32(0))
						mBase = m.M
						v33 = F___syscall_getpid(m)
						mBase = m.M
						v36 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
						v40 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)))
						v41 = v36 ^ base.I64_extend_i32_u(v14) ^ base.I64_extend_i32_u(v33) ^ v40
						*(*uint8)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1113]))) = uint8(v41)
						v44 = v29 + int32(1)
						if v44 != int32(64) {
							v29 = v44
							continue
						} else {
							break
						}
						break
					}
					if v14 == int32(0) {
						m.G0 = v8 + int32(16)
						return
					} else {
						v63 = F_fclose(m, v14)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_internal_memalign(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var __phi167 int32
	_ = __phi167
	var v168 int32
	_ = v168
	var __phi168 int32
	_ = __phi168
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var __phi329 int32
	_ = __phi329
	var v330 int32
	_ = v330
	var __phi330 int32
	_ = __phi330
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v412 int32
	_ = v412
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v522 int32
	_ = v522
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var __phi635 int32
	_ = __phi635
	var v636 int32
	_ = v636
	var __phi636 int32
	_ = __phi636
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v668 int32
	_ = v668
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	var __phi797 int32
	_ = __phi797
	var v798 int32
	_ = v798
	var __phi798 int32
	_ = __phi798
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v822 int32
	_ = v822
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v880 int32
	_ = v880
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v990 int32
	_ = v990
	v8 = int32(16)
	if base.Ui32(v8) < base.Ui32(l0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if base.Ui32(l1) < base.Ui32(int32(-64)-v26) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v18 = v8
	goto L7
L3:
	;
	v12 = l0
	goto L5
L4:
	;
	v12 = v8
	goto L5
L5:
	;
	if v12&(v12+int32(-1)) != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v26 = v12
	goto L1
L7:
	;
	if base.Ui32(v18) < base.Ui32(v12) {
		v18 = v18 << (uint(int32(1)) % 32)
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v26 = v18
	goto L1
L9:
	;
	goto L8
L10:
	;
	v42 = int32(11)
	if base.Ui32(l1) < base.Ui32(v42) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(48)
	return int32(0)
L13:
	;
	v56 = v52 + int32(-8)
	if (v26+int32(-1))&v52 != 0 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v48 = int32(16)
	goto L16
L15:
	;
	v48 = (l1 + v42) & int32(-8)
	goto L16
L16:
	;
	v52 = F_emscripten_builtin_malloc(m, v48+v26+int32(12))
	mBase = m.M
	if v52 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	return int32(0)
L18:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	if v553&int32(3) == int32(0) {
		goto L118
	} else {
		goto L119
	}
L19:
	;
	v61 = v52 + int32(-4)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = int32(-8)
	v68 = int32(0)
	v72 = (v52+v26+int32(-1))&(v68-v26) + v63
	if base.Ui32(int32(15)) < base.Ui32(v72-v56) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v547 = v56
	goto L18
L21:
	;
	v77 = v68
	goto L23
L22:
	;
	v77 = v26
	goto L23
L23:
	;
	v78 = v72 + v77
	v79 = v78 - v56
	v80 = v62&v63 - v79
	if v62&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v88 = int32(1)
	v91 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v80 | v87&v88 | v91
	v94 = v78 + v80
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v95 | v88
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v79 | v99&v88 | v91
	v106 = v56 + v79
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v107 | v88
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v119&v88 != 0 {
		v240 = v56
		v241 = v79
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v83 + v79
	v547 = v78
	goto L18
L26:
	;
	v547 = v78
	goto L18
L27:
	;
	goto L26
L28:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v249&int32(2) != 0 {
		goto L66
	} else {
		goto L67
	}
L29:
	;
	if v119&int32(2) == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v127 = v126 + v79
	v128 = v56 - v126
	v130 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v128 == v130 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	if v146 == int32(0) {
		v240 = v128
		v241 = v127
		goto L28
	} else {
		goto L50
	}
L32:
	;
	v200 = int32(0)
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v135
	v240 = v128
	v241 = v127
	goto L28
L34:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v181 = int32(3)
	if v180&v181 != v181 {
		v240 = v128
		v241 = v127
		goto L28
	} else {
		goto L49
	}
L35:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	if base.Ui32(int32(255)) < base.Ui32(v126) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v128)+24))
	if v132 == v128 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	if v132 != v135 {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v137 = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v139 & base.I32_rotl(int32(-2), int32(base.Ui32(v126)>>(uint(int32(3))%32)))
	v240 = v128
	v241 = v127
	goto L28
L39:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v128)+20))
	if v151 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v148
	v200 = v132
	goto L31
L41:
	;
	__phi167 = v161
	__phi168 = v162
	v167 = __phi167
	v168 = __phi168
	goto L45
L42:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	if v156 == int32(0) {
		goto L32
	} else {
		goto L44
	}
L43:
	;
	v161 = v151
	v162 = v128 + int32(20)
	goto L41
L44:
	;
	v161 = v156
	v162 = v128 + int32(16)
	goto L41
L45:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	if v174 != 0 {
		__phi167 = v174
		__phi168 = v167 + int32(20)
		v167 = __phi167
		v168 = __phi168
		goto L45
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = int32(0)
	v200 = v167
	goto L31
L47:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	if v177 != 0 {
		__phi167 = v177
		__phi168 = v167 + int32(16)
		v167 = __phi167
		v168 = __phi168
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v180 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v127 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v127
	goto L26
L50:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v128)+28))
	v210 = v208 << (uint(int32(2)) % 32)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_consts[1211])))
	if v128 != v213 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+24)) = v146
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	if v230 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L52:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v146)+16))
	if v223 != v128 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_consts[1211]))) = v200
	if v200 != 0 {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v216 = int32(0)
	v218 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v218 & base.I32_rotl(int32(-2), v208)
	v240 = v128
	v241 = v127
	goto L28
L55:
	;
	if v200 == int32(0) {
		v240 = v128
		v241 = v127
		goto L28
	} else {
		goto L58
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+20)) = v200
	goto L55
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+16)) = v200
	goto L55
L58:
	;
	goto L51
L59:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v128)+20))
	if v235 == int32(0) {
		v240 = v128
		v241 = v127
		goto L28
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v230)+24)) = v200
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+20)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v235)+24)) = v200
	v240 = v128
	v241 = v127
	goto L28
L62:
	;
	if base.Ui32(int32(255)) < base.Ui32(v412) {
		goto L100
	} else {
		goto L101
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+4)) = v291 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v240+v291))) = v291
	if v240 != v275 {
		v412 = v291
		goto L62
	} else {
		goto L99
	}
L64:
	;
	if v308 == int32(0) {
		goto L63
	} else {
		goto L87
	}
L65:
	;
	v354 = int32(0)
	goto L64
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v249 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v240)+4)) = v241 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v240+v241))) = v241
	v412 = v241
	goto L62
L67:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _consts[1213]))
	if v106 != v253 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v106 != v275 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v255 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1213])) = v240
	v259 = *(*int32)(unsafe.Add(mBase, _consts[1214]))
	v260 = v259 + v241
	*(*int32)(unsafe.Add(mBase, _consts[1214])) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v240)+4)) = v260 | int32(1)
	v266 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v240 != v266 {
		goto L27
	} else {
		goto L70
	}
L70:
	;
	v268 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v268
	*(*int32)(unsafe.Add(mBase, _consts[1208])) = v268
	goto L26
L71:
	;
	v291 = v249&int32(-8) + v241
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	if base.Ui32(int32(255)) < base.Ui32(v249) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v277 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1208])) = v240
	v281 = *(*int32)(unsafe.Add(mBase, _consts[1210]))
	v282 = v281 + v241
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v240)+4)) = v282 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v240+v282))) = v282
	goto L26
L73:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v106)+24))
	if v292 == v106 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	if v292 != v295 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v292)+8)) = v295
	goto L63
L76:
	;
	v297 = int32(0)
	v299 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v299 & base.I32_rotl(int32(-2), int32(base.Ui32(v249)>>(uint(int32(3))%32)))
	goto L63
L77:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	if v313 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v310)+12)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v292)+8)) = v310
	v354 = v292
	goto L64
L79:
	;
	__phi329 = v323
	__phi330 = v324
	v329 = __phi329
	v330 = __phi330
	goto L83
L80:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	if v318 == int32(0) {
		goto L65
	} else {
		goto L82
	}
L81:
	;
	v323 = v313
	v324 = v106 + int32(20)
	goto L79
L82:
	;
	v323 = v318
	v324 = v106 + int32(16)
	goto L79
L83:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v329)+20))
	if v336 != 0 {
		__phi329 = v336
		__phi330 = v329 + int32(20)
		v329 = __phi329
		v330 = __phi330
		goto L83
	} else {
		goto L85
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330))) = int32(0)
	v354 = v329
	goto L64
L85:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v329)+16))
	if v339 != 0 {
		__phi329 = v339
		__phi330 = v329 + int32(16)
		v329 = __phi329
		v330 = __phi330
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v106)+28))
	v364 = v362 << (uint(int32(2)) % 32)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v364)+uint32(_consts[1211])))
	if v106 != v367 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+24)) = v308
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	if v384 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L89:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v308)+16))
	if v377 != v106 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364)+uint32(_consts[1211]))) = v354
	if v354 != 0 {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v370 = int32(0)
	v372 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v372 & base.I32_rotl(int32(-2), v362)
	goto L63
L92:
	;
	if v354 == int32(0) {
		goto L63
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v308)+20)) = v354
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v308)+16)) = v354
	goto L92
L95:
	;
	goto L88
L96:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	if v389 == int32(0) {
		goto L63
	} else {
		goto L98
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+16)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v384)+24)) = v354
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+20)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v389)+24)) = v354
	goto L63
L99:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v291
	goto L26
L100:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v412) {
		v458 = int32(31)
		goto L105
	} else {
		goto L106
	}
L101:
	;
	v423 = v412 & int32(-8)
	v425 = v423 + int32(9128464)
	v427 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	v431 = int32(1) << (uint(int32(base.Ui32(v412)>>(uint(int32(3))%32))) % 32)
	if v427&v431 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v423)+uint32(_consts[1215]))) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v437)+12)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v240)+12)) = v425
	*(*int32)(unsafe.Add(mBase, uint32(v240)+8)) = v437
	goto L26
L103:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v423)+uint32(_consts[1215])))
	v437 = v436
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v427 | v431
	v437 = v425
	goto L102
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+28)) = v458
	*(*int64)(unsafe.Add(mBase, uint32(v240)+16)) = int64(0)
	v463 = v458 << (uint(int32(2)) % 32)
	v467 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	v469 = int32(1) << (uint(v458) % 32)
	if v467&v469 != 0 {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v448 = base.I32_clz(int32(base.Ui32(v412) >> (uint(int32(8)) % 32)))
	v451 = int32(1)
	v458 = int32(base.Ui32(v412)>>(uint(int32(38)-v448)%32))&v451 - v448<<(uint(v451)%32) + int32(62)
	goto L105
L107:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v491)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v522)+12)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v491)+8)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v240)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v240)+12)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v240)+8)) = v522
	goto L27
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+12)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v240)+8)) = v240
	goto L26
L109:
	;
	if v458 == int32(31) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v467 | v469
	*(*int32)(unsafe.Add(mBase, uint32(v463)+uint32(_consts[1211]))) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v240)+24)) = v463 + int32(9128728)
	goto L108
L111:
	;
	v483 = int32(0)
	goto L113
L112:
	;
	v483 = int32(25) - int32(base.Ui32(v458)>>(uint(int32(1))%32))
	goto L113
L113:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v463)+uint32(_consts[1211])))
	v489 = v412 << (uint(v483) % 32)
	v491 = v485
	goto L114
L114:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v491)+4))
	if v495&int32(-8) == v412 {
		goto L107
	} else {
		goto L116
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v505+int32(16)))) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v240)+24)) = v491
	goto L108
L116:
	;
	v505 = v491 + int32(base.Ui32(v489)>>(uint(int32(29))%32))&int32(4)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v505)+16))
	if v506 != 0 {
		v489 = v489 << (uint(int32(1)) % 32)
		v491 = v506
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	return v547 + int32(8)
L119:
	;
	v559 = v553 & int32(-8)
	if base.Ui32(v559) <= base.Ui32(v48+int32(16)) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v563 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v547)+4)) = v48 | v553&v563 | int32(2)
	v569 = v547 + v48
	v570 = v559 - v48
	*(*int32)(unsafe.Add(mBase, uint32(v569)+4)) = v570 | int32(3)
	v574 = v547 + v559
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v574)+4)) = v575 | v563
	v586 = v569 + v570
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v569)+4))
	if v587&v563 != 0 {
		v708 = v569
		v709 = v570
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L118
L122:
	;
	goto L121
L123:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v586)+4))
	if v717&int32(2) != 0 {
		goto L161
	} else {
		goto L162
	}
L124:
	;
	if v587&int32(2) == int32(0) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	v595 = v594 + v570
	v596 = v569 - v594
	v598 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v596 == v598 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	if v614 == int32(0) {
		v708 = v596
		v709 = v595
		goto L123
	} else {
		goto L145
	}
L127:
	;
	v668 = int32(0)
	goto L126
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+12)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v600)+8)) = v603
	v708 = v596
	v709 = v595
	goto L123
L129:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v586)+4))
	v649 = int32(3)
	if v648&v649 != v649 {
		v708 = v596
		v709 = v595
		goto L123
	} else {
		goto L144
	}
L130:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v596)+12))
	if base.Ui32(int32(255)) < base.Ui32(v594) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v596)+24))
	if v600 == v596 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v596)+8))
	if v600 != v603 {
		goto L128
	} else {
		goto L133
	}
L133:
	;
	v605 = int32(0)
	v607 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v607 & base.I32_rotl(int32(-2), int32(base.Ui32(v594)>>(uint(int32(3))%32)))
	v708 = v596
	v709 = v595
	goto L123
L134:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v596)+20))
	if v619 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v596)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v616)+12)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v600)+8)) = v616
	v668 = v600
	goto L126
L136:
	;
	__phi635 = v629
	__phi636 = v630
	v635 = __phi635
	v636 = __phi636
	goto L140
L137:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v596)+16))
	if v624 == int32(0) {
		goto L127
	} else {
		goto L139
	}
L138:
	;
	v629 = v619
	v630 = v596 + int32(20)
	goto L136
L139:
	;
	v629 = v624
	v630 = v596 + int32(16)
	goto L136
L140:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v635)+20))
	if v642 != 0 {
		__phi635 = v642
		__phi636 = v635 + int32(20)
		v635 = __phi635
		v636 = __phi636
		goto L140
	} else {
		goto L142
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v636))) = int32(0)
	v668 = v635
	goto L126
L142:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v635)+16))
	if v645 != 0 {
		__phi635 = v645
		__phi636 = v635 + int32(16)
		v635 = __phi635
		v636 = __phi636
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v595
	*(*int32)(unsafe.Add(mBase, uint32(v586)+4)) = v648 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v596)+4)) = v595 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v586))) = v595
	goto L121
L145:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v596)+28))
	v678 = v676 << (uint(int32(2)) % 32)
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v678)+uint32(_consts[1211])))
	if v596 != v681 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v668)+24)) = v614
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v596)+16))
	if v698 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L147:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v614)+16))
	if v691 != v596 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v678)+uint32(_consts[1211]))) = v668
	if v668 != 0 {
		goto L146
	} else {
		goto L149
	}
L149:
	;
	v684 = int32(0)
	v686 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v686 & base.I32_rotl(int32(-2), v676)
	v708 = v596
	v709 = v595
	goto L123
L150:
	;
	if v668 == int32(0) {
		v708 = v596
		v709 = v595
		goto L123
	} else {
		goto L153
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v614)+20)) = v668
	goto L150
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v614)+16)) = v668
	goto L150
L153:
	;
	goto L146
L154:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v596)+20))
	if v703 == int32(0) {
		v708 = v596
		v709 = v595
		goto L123
	} else {
		goto L156
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v668)+16)) = v698
	*(*int32)(unsafe.Add(mBase, uint32(v698)+24)) = v668
	goto L154
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v668)+20)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v703)+24)) = v668
	v708 = v596
	v709 = v595
	goto L123
L157:
	;
	if base.Ui32(int32(255)) < base.Ui32(v880) {
		goto L195
	} else {
		goto L196
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+4)) = v759 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v708+v759))) = v759
	if v708 != v743 {
		v880 = v759
		goto L157
	} else {
		goto L194
	}
L159:
	;
	if v776 == int32(0) {
		goto L158
	} else {
		goto L182
	}
L160:
	;
	v822 = int32(0)
	goto L159
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v586)+4)) = v717 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+4)) = v709 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v708+v709))) = v709
	v880 = v709
	goto L157
L162:
	;
	v721 = *(*int32)(unsafe.Add(mBase, _consts[1213]))
	if v586 != v721 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v743 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v586 != v743 {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v723 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1213])) = v708
	v727 = *(*int32)(unsafe.Add(mBase, _consts[1214]))
	v728 = v727 + v709
	*(*int32)(unsafe.Add(mBase, _consts[1214])) = v728
	*(*int32)(unsafe.Add(mBase, uint32(v708)+4)) = v728 | int32(1)
	v734 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v708 != v734 {
		goto L122
	} else {
		goto L165
	}
L165:
	;
	v736 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v736
	*(*int32)(unsafe.Add(mBase, _consts[1208])) = v736
	goto L121
L166:
	;
	v759 = v717&int32(-8) + v709
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v586)+12))
	if base.Ui32(int32(255)) < base.Ui32(v717) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v745 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1208])) = v708
	v749 = *(*int32)(unsafe.Add(mBase, _consts[1210]))
	v750 = v749 + v709
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v750
	*(*int32)(unsafe.Add(mBase, uint32(v708)+4)) = v750 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v708+v750))) = v750
	goto L121
L168:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v586)+24))
	if v760 == v586 {
		goto L172
	} else {
		goto L173
	}
L169:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v586)+8))
	if v760 != v763 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v763)+12)) = v760
	*(*int32)(unsafe.Add(mBase, uint32(v760)+8)) = v763
	goto L158
L171:
	;
	v765 = int32(0)
	v767 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v767 & base.I32_rotl(int32(-2), int32(base.Ui32(v717)>>(uint(int32(3))%32)))
	goto L158
L172:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v586)+20))
	if v781 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v586)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v778)+12)) = v760
	*(*int32)(unsafe.Add(mBase, uint32(v760)+8)) = v778
	v822 = v760
	goto L159
L174:
	;
	__phi797 = v791
	__phi798 = v792
	v797 = __phi797
	v798 = __phi798
	goto L178
L175:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v586)+16))
	if v786 == int32(0) {
		goto L160
	} else {
		goto L177
	}
L176:
	;
	v791 = v781
	v792 = v586 + int32(20)
	goto L174
L177:
	;
	v791 = v786
	v792 = v586 + int32(16)
	goto L174
L178:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v797)+20))
	if v804 != 0 {
		__phi797 = v804
		__phi798 = v797 + int32(20)
		v797 = __phi797
		v798 = __phi798
		goto L178
	} else {
		goto L180
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v798))) = int32(0)
	v822 = v797
	goto L159
L180:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v797)+16))
	if v807 != 0 {
		__phi797 = v807
		__phi798 = v797 + int32(16)
		v797 = __phi797
		v798 = __phi798
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v586)+28))
	v832 = v830 << (uint(int32(2)) % 32)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v832)+uint32(_consts[1211])))
	if v586 != v835 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v822)+24)) = v776
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v586)+16))
	if v852 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L184:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v776)+16))
	if v845 != v586 {
		goto L188
	} else {
		goto L189
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832)+uint32(_consts[1211]))) = v822
	if v822 != 0 {
		goto L183
	} else {
		goto L186
	}
L186:
	;
	v838 = int32(0)
	v840 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v840 & base.I32_rotl(int32(-2), v830)
	goto L158
L187:
	;
	if v822 == int32(0) {
		goto L158
	} else {
		goto L190
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v776)+20)) = v822
	goto L187
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v776)+16)) = v822
	goto L187
L190:
	;
	goto L183
L191:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v586)+20))
	if v857 == int32(0) {
		goto L158
	} else {
		goto L193
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v822)+16)) = v852
	*(*int32)(unsafe.Add(mBase, uint32(v852)+24)) = v822
	goto L191
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v822)+20)) = v857
	*(*int32)(unsafe.Add(mBase, uint32(v857)+24)) = v822
	goto L158
L194:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v759
	goto L121
L195:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v880) {
		v926 = int32(31)
		goto L200
	} else {
		goto L201
	}
L196:
	;
	v891 = v880 & int32(-8)
	v893 = v891 + int32(9128464)
	v895 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	v899 = int32(1) << (uint(int32(base.Ui32(v880)>>(uint(int32(3))%32))) % 32)
	if v895&v899 != 0 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891)+uint32(_consts[1215]))) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v905)+12)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v708)+12)) = v893
	*(*int32)(unsafe.Add(mBase, uint32(v708)+8)) = v905
	goto L121
L198:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v891)+uint32(_consts[1215])))
	v905 = v904
	goto L197
L199:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v895 | v899
	v905 = v893
	goto L197
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+28)) = v926
	*(*int64)(unsafe.Add(mBase, uint32(v708)+16)) = int64(0)
	v931 = v926 << (uint(int32(2)) % 32)
	v935 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	v937 = int32(1) << (uint(v926) % 32)
	if v935&v937 != 0 {
		goto L204
	} else {
		goto L205
	}
L201:
	;
	v916 = base.I32_clz(int32(base.Ui32(v880) >> (uint(int32(8)) % 32)))
	v919 = int32(1)
	v926 = int32(base.Ui32(v880)>>(uint(int32(38)-v916)%32))&v919 - v916<<(uint(v919)%32) + int32(62)
	goto L200
L202:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v959)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v990)+12)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v959)+8)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v708)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+12)) = v959
	*(*int32)(unsafe.Add(mBase, uint32(v708)+8)) = v990
	goto L122
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+12)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v708)+8)) = v708
	goto L121
L204:
	;
	if v926 == int32(31) {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v935 | v937
	*(*int32)(unsafe.Add(mBase, uint32(v931)+uint32(_consts[1211]))) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v708)+24)) = v931 + int32(9128728)
	goto L203
L206:
	;
	v951 = int32(0)
	goto L208
L207:
	;
	v951 = int32(25) - int32(base.Ui32(v926)>>(uint(int32(1))%32))
	goto L208
L208:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v931)+uint32(_consts[1211])))
	v957 = v880 << (uint(v951) % 32)
	v959 = v953
	goto L209
L209:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v959)+4))
	if v963&int32(-8) == v880 {
		goto L202
	} else {
		goto L211
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v973+int32(16)))) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v708)+24)) = v959
	goto L203
L211:
	;
	v973 = v959 + int32(base.Ui32(v957)>>(uint(int32(29))%32))&int32(4)
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v973)+16))
	if v974 != 0 {
		v957 = v957 << (uint(int32(1)) % 32)
		v959 = v974
		goto L209
	} else {
		goto L212
	}
L212:
	;
	goto L210
}
func F_invalidateClusterSlotsResp(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	F_clearCachedClusterSlotsResponse(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_ipairsaux(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	v4 = F_luaL_checkinteger(m, l0, int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_luaL_checktype(m, l0, int32(1), int32(5))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v13 = v4 + int32(1)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v15))) = base.F64_convert_i32_s(v13)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v20 + int32(16)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v35 = v30 + int32(0)
			v36 = m.G398
			if base.Ui32(v35) < base.Ui32(v29) {
				v38 = v35
			} else {
				v38 = v36
			}
			v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
			v81 = F_luaH_getnum(m, v80, v13)
			mBase = m.M
			v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v83 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
			*(*int64)(unsafe.Add(mBase, uint32(v82))) = v83
			v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v85
			v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87 + int32(16)
			v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v108 = v105 + int32(-16)
			v142 = m.G398
			if v108 != v142 {
				v145 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
				v148 = v145
			} else {
				v148 = int32(-1)
			}
			return base.B2i32(v148 != int32(0)) << (uint(int32(1)) % 32)
		}
	}
}
func F_iprintf(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1
	v10 = F_vfiprintf(m, int32(_a2354), l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v10
	}
}
func F_isAnySlotInManualMigratingState(m *base.Module) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+44))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	return base.B2i32(v5 != int32(0)-v7)
}
func F_isCopyAvoidPreferred(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v6&int32(268435456) != 0 {
		v160 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v160
L2:
	;
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
	if v9 != int64(-1) {
		v160 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v6&int32(131072) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v40&int32(1) != 0 {
		v160 = v3
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	if l0 != v17 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v20 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	if v23 == int32(0) {
		v160 = v3
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	if v26 == int32(288) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v26 == int32(286) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v26 == int32(284) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if v26 == int32(282) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v26 == int32(287) {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v26 != int32(289) {
		v160 = v3
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L4
L15:
	;
	if v40&int32(2) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v40&int32(262144) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if v40&int32(4) == int32(0) {
		v160 = v3
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	if l1 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v53 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	goto L22
L22:
	;
	return int32(0)
L23:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v71&int32(240) != 0 {
		v160 = v3
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v62 = int32(_a69)
	v63 = *(*int32)(unsafe.Add(mBase, _consts[419]))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[418]))
	return base.B2i32(v63 != int32(0)) & base.B2i32(v63 <= v67)
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v74&int32(-8) == int32(-16) {
		v160 = v3
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v79 = int32(_a69)
	v80 = *(*int32)(unsafe.Add(mBase, _consts[418]))
	v82 = *(*int32)(unsafe.Add(mBase, _consts[419]))
	if v82 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v80 != int32(1) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v82 <= v80 {
		v160 = int32(1)
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v125 = int32(0)
	v127 = *(*int32)(unsafe.Add(mBase, _consts[420]))
	if v127 == v125 {
		v160 = v125
		goto L1
	} else {
		goto L40
	}
L31:
	;
	v90 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, _consts[421]))
	if v92 == v90 {
		v160 = v90
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v95 = F_objectGetVal(m, l1)
	mBase = m.M
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+int32(-1)))))
	switch v101 & int32(7) {
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
		v118 = int32(0)
		goto L34
	}
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[421]))
	return base.B2i32(base.Ui32(v122) <= base.Ui32(v120))
L34:
	;
	v120 = v118
	goto L33
L35:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v95+int32(-17))))
	v118 = v117
	goto L34
L36:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v95+int32(-9))))
	v120 = v114
	goto L33
L37:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95+int32(-5)))))
	v120 = v111
	goto L33
L38:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+int32(-3)))))
	v120 = v108
	goto L33
L39:
	;
	v120 = int32(base.Ui32(v101) >> (uint(int32(3)) % 32))
	goto L33
L40:
	;
	v130 = F_objectGetVal(m, l1)
	mBase = m.M
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+int32(-1)))))
	switch v136 & int32(7) {
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
		v153 = int32(0)
		goto L42
	}
L41:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[420]))
	v160 = base.B2i32(base.Ui32(v157) <= base.Ui32(v155))
	goto L1
L42:
	;
	v155 = v153
	goto L41
L43:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v130+int32(-17))))
	v153 = v152
	goto L42
L44:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v130+int32(-9))))
	v155 = v149
	goto L41
L45:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130+int32(-5)))))
	v155 = v146
	goto L41
L46:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+int32(-3)))))
	v155 = v143
	goto L41
L47:
	;
	v155 = int32(base.Ui32(v136) >> (uint(int32(3)) % 32))
	goto L41
}
func F_isHLLObjectOrReply(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v7 = F_checkType(m, l0, l1, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v47 = int32(-1)
			return v47
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			switch int32(base.Ui32(v11)>>(uint(int32(4))%32)) & int32(15) {
			case 0, 8:
				v16 = F_stringObjectLen(m, l1)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					if base.Ui32(v16) < base.Ui32(int32(16)) {
						F_addReplyError(m, l0, int32(_a640))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v47 = int32(-1)
							return v47
						}
					} else {
						v20 = F_objectGetVal(m, l1)
						mBase = m.M
						v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
						if v21 != int32(72) {
							F_addReplyError(m, l0, int32(_a640))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v47 = int32(-1)
								return v47
							}
						} else {
							v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
							if v24 != int32(89) {
								F_addReplyError(m, l0, int32(_a640))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v47 = int32(-1)
									return v47
								}
							} else {
								v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
								if v27 != int32(76) {
									F_addReplyError(m, l0, int32(_a640))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										v47 = int32(-1)
										return v47
									}
								} else {
									v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
									if v30 != int32(76) {
										F_addReplyError(m, l0, int32(_a640))
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int32(0)
										} else {
											v47 = int32(-1)
											return v47
										}
									} else {
										v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
										if base.Ui32(int32(1)) < base.Ui32(v33) {
											F_addReplyError(m, l0, int32(_a640))
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int32(0)
											} else {
												v47 = int32(-1)
												return v47
											}
										} else {
											v36 = int32(0)
											if v33 != 0 {
												v47 = v36
												return v47
											} else {
												v37 = F_stringObjectLen(m, l1)
												mBase = m.M
												v38 = m.ExcPending
												if v38 != 0 {
													return int32(0)
												} else {
													if v37 == int32(12304) {
														v47 = v36
														return v47
													} else {
														F_addReplyError(m, l0, int32(_a640))
														mBase = m.M
														v45 = m.ExcPending
														if v45 != 0 {
															return int32(0)
														} else {
															v47 = int32(-1)
															return v47
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
			default:
				F_addReplyError(m, l0, int32(_a640))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v47 = int32(-1)
					return v47
				}
			}
		}
	}
}
func F_isImportSlotMigrationJob(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return base.B2i32(v2 == int32(1))
}
func F_isblank(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 == int32(32)) | base.B2i32(l0 == int32(9))
}
func F_iscntrl(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0) < base.Ui32(int32(32))) | base.B2i32(l0 == int32(127))
}
func F_iswlower(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	v3 = F_casemap(m, l0, int32(1))
	return base.B2i32(v3 != l0)
}
func F_iswprint(m *base.Module, l0 int32) int32 {
	var v12 int32
	_ = v12
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	if base.Ui32(int32(254)) < base.Ui32(l0) {
		v12 = int32(1)
		if base.Ui32(l0+int32(-57344)) < base.Ui32(int32(8185)) {
			v32 = v12
		} else {
			if base.Ui32(l0) < base.Ui32(int32(8232)) {
				v32 = v12
			} else {
				if base.Ui32(l0+int32(-8234)) < base.Ui32(int32(47062)) {
					v32 = v12
				} else {
					v27 = int32(65534)
					v32 = base.B2i32(base.Ui32(l0+int32(-65532)) < base.Ui32(int32(1048580))) & base.B2i32(l0&v27 != v27)
				}
			}
		}
		return v32
	} else {
		return base.B2i32(base.Ui32(int32(32)) < base.Ui32((l0+int32(1))&int32(127)))
	}
}
func F_iswpunct(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	if base.Ui32(int32(131071)) < base.Ui32(l0) {
		v26 = int32(0)
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_consts[1238]))))
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10<<(uint(int32(5))%32)|int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31))+uint32(_consts[1238]))))
		v26 = int32(base.Ui32(v20)>>(uint(l0&int32(7))%32)) & int32(1)
	}
	return v26
}
func F_iswupper(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	v3 = F_casemap(m, l0, int32(0))
	return base.B2i32(v3 != l0)
}
func F_iswxdigit(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0+int32(-48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(l0|int32(32)+int32(-97)) < base.Ui32(int32(6)))
}
func F_isxdigit(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0+int32(-48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(l0|int32(32)+int32(-97)) < base.Ui32(int32(6)))
}
