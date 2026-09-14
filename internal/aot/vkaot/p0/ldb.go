package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_ldbCatStackValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_ldbCatStackValueRec(m, l0, l1, l2, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_ldbCatStackValueRec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 float64
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
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
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 float64
	_ = v313
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
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
	var v359 int32
	_ = v359
	var v378 int32
	_ = v378
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	if l2 < int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l3 != int32(10) {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v68 = m.G398
	if v67 != v68 {
		goto L14
	} else {
		goto L15
	}
L3:
	;
	if l2 < int32(-9999) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v25 = v20 + l2<<(uint(int32(4))%32) + int32(-16)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v25) < base.Ui32(v26) {
		v67 = v25
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v74 = int32(-1)
	goto L1
L6:
	;
	switch l2 + int32(10002) {
	case 0:
		goto L10
	case 1:
		goto L11
	case 2:
		goto L8
	default:
		goto L9
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v67 = v31 + l2<<(uint(int32(4))%32)
	goto L2
L8:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v67 = v63 + int32(96)
	goto L2
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+7)))
	if base.Ui32(v53) < base.Ui32(int32(-10002)-l2) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v67 = l1 + int32(72)
	goto L2
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v40
	v67 = l1 + int32(88)
	goto L2
L12:
	;
	v74 = int32(-1)
	goto L1
L13:
	;
	v67 = v52 + (int32(-10003)-l2)<<(uint(int32(4))%32) + int32(24)
	goto L2
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	v74 = v71
	goto L1
L15:
	;
	v74 = int32(-1)
	goto L1
L16:
	;
	m.G0 = v16 + int32(48)
	return v627
L17:
	;
	switch v74 {
	case 0:
		goto L24
	case 1:
		goto L26
	case 2, 6, 7, 8:
		goto L22
	case 3:
		goto L25
	case 4:
		goto L27
	case 5:
		goto L23
	default:
		goto L21
	}
L18:
	;
	v78 = m.G3
	v82 = m.G16
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = m.T0[v83].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), l0, v78+int32(_a_F_ldbCatStackValueRec_0), int32(44))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	v627 = l0
	goto L16
L21:
	;
	v619 = m.G3
	v623 = m.G16
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v623)))
	v625 = m.T0[v624].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), l0, v619+int32(_a_F_ldbCatStackValueRec_1), int32(20))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L19
	} else {
		goto L163
	}
L22:
	;
	v452 = m.G3
	if l2 < int32(1) {
		goto L123
	} else {
		goto L124
	}
L23:
	;
	v205 = int32(0)
	v206 = m.G3
	v208 = v206 + int32(_a_F_ldbCatStackValueRec_2)
	v210 = m.G13
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v212 = m.T0[v211].(func(*base.Module, int32, int32, int32) int32)(m, v205, v208, v205)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L19
	} else {
		goto L61
	}
L24:
	;
	v197 = m.G3
	v201 = m.G16
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v203 = m.T0[v202].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), l0, v197+int32(_a_F_ldbCatStackValueRec_3), int32(3))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L19
	} else {
		goto L60
	}
L25:
	;
	v174 = m.G7
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v176 = m.G15
	v178 = m.T0[v175].(func(*base.Module, int32, int32) int32)(m, l0, int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L19
	} else {
		goto L56
	}
L26:
	;
	v99 = m.G3
	if l2 < int32(1) {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	v88 = m.G16
	v92 = F_lua_tolstring(m, l1, l2, v16+int32(44))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v96 = m.T0[v95].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), l0, v92, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v627 = l0
	goto L16
L30:
	;
	if v165 != 0 {
		goto L49
	} else {
		goto L50
	}
L31:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v156)+8))
	switch v158 {
	case 0:
		v163 = v158
		goto L46
	case 1:
		goto L48
	default:
		goto L47
	}
L32:
	;
	if l2 < int32(-9999) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v113 = v108 + l2<<(uint(int32(4))%32) + int32(-16)
	v114 = m.G398
	if base.Ui32(v113) < base.Ui32(v107) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v116 = v113
	goto L36
L35:
	;
	v116 = v114
	goto L36
L36:
	;
	v156 = v116
	goto L31
L37:
	;
	switch l2 + int32(10002) {
	case 0:
		goto L40
	case 1:
		goto L41
	case 2:
		goto L42
	default:
		goto L39
	}
L38:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v156 = v119 + l2<<(uint(int32(4))%32)
	goto L31
L39:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+7)))
	v143 = m.G398
	if base.Ui32(v142) < base.Ui32(int32(-10002)-l2) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v156 = l1 + int32(72)
	goto L31
L41:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v131
	v156 = l1 + int32(88)
	goto L31
L42:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v156 = v125 + int32(96)
	goto L31
L43:
	;
	v154 = v143
	goto L45
L44:
	;
	v154 = v141 + (int32(-10003)-l2)<<(uint(int32(4))%32) + int32(24)
	goto L45
L45:
	;
	v156 = v154
	goto L31
L46:
	;
	v165 = v163
	goto L30
L47:
	;
	v163 = int32(1)
	goto L46
L48:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v165 = base.B2i32(v159 != int32(0))
	goto L30
L49:
	;
	v166 = v99 + int32(_a_F_ldbCatStackValueRec_4)
	goto L51
L50:
	;
	v166 = v99 + int32(_a_F_ldbCatStackValueRec_5)
	goto L51
L51:
	;
	if v165 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v169 = int32(4)
	goto L54
L53:
	;
	v169 = int32(5)
	goto L54
L54:
	;
	v170 = m.G16
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v172 = m.T0[v171].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), l0, v166, v169)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L19
	} else {
		goto L55
	}
L55:
	;
	v627 = l0
	goto L16
L56:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v181 = F_lua_tonumber(m, l1, l2)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L19
	} else {
		goto L57
	}
L57:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v178
	v185 = m.G3
	v186 = m.G17
	v190 = m.T0[v180].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v185+int32(_a_F_ldbCatStackValueRec_6), v16)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L19
	} else {
		goto L58
	}
L58:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	m.T0[v193].(func(*base.Module, int32, int32))(m, int32(0), l0)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L19
	} else {
		goto L59
	}
L59:
	;
	v627 = v190
	goto L16
L60:
	;
	v627 = l0
	goto L16
L61:
	;
	v214 = int32(0)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v217 = m.T0[v216].(func(*base.Module, int32, int32, int32) int32)(m, v214, v208, v214)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L19
	} else {
		goto L62
	}
L62:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v220 + int32(16)
	goto L63
L63:
	;
	v227 = l2 + int32(-1)
	v228 = F_lua_next(m, l1, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L19
	} else {
		goto L66
	}
L64:
	;
	v406 = m.G3
	v410 = m.G16
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v412 = m.T0[v411].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), l0, v406+int32(_a_F_ldbCatStackValueRec_7), int32(1))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L19
	} else {
		goto L108
	}
L65:
	;
	v231 = int32(1)
	v232 = l3 + v231
	v240 = v217
	v242 = v212
	v245 = v231
	v246 = v231
	goto L68
L66:
	;
	if v228 != 0 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v397 = v217
	v399 = v212
	v401 = int32(0)
	goto L64
L68:
	;
	v249 = int32(0)
	if v246 == v249 {
		v321 = int32(1)
		v322 = v249
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v397 = v354
	v399 = v327
	v401 = v321
	goto L64
L70:
	;
	v323 = m.G3
	v324 = m.G16
	v327 = F_ldbCatStackValueRec(m, v242, l1, int32(-1), v232)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L19
	} else {
		goto L91
	}
L71:
	;
	goto L75
L72:
	;
	v321 = int32(1)
	v322 = int32(0)
	goto L70
L73:
	;
	if v309 != int32(3) {
		goto L72
	} else {
		goto L88
	}
L74:
	;
	v303 = m.G398
	if v269 != v303 {
		goto L86
	} else {
		goto L87
	}
L75:
	;
	goto L79
L79:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v269 = v266 + int32(-32)
	goto L74
L86:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v269)+8))
	v309 = v306
	goto L73
L87:
	;
	v309 = int32(-1)
	goto L73
L88:
	;
	v313 = F_lua_tonumber(m, l1, int32(-2))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L19
	} else {
		goto L89
	}
L89:
	;
	if base.F64_ne(v313, base.F64_convert_i32_u(v245)) != 0 {
		goto L72
	} else {
		goto L90
	}
L90:
	;
	v321 = int32(0)
	v322 = int32(1)
	goto L70
L91:
	;
	v330 = v323 + int32(_a_F_ldbCatStackValueRec_8)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v333 = m.T0[v332].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), v327, v330, int32(2))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L19
	} else {
		goto L92
	}
L92:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v340 = m.T0[v339].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), v240, v323+int32(_a_F_ldbCatStackValueRec_9), int32(1))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L19
	} else {
		goto L93
	}
L93:
	;
	v344 = F_ldbCatStackValueRec(m, v240, l1, int32(-2), v232)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L19
	} else {
		goto L94
	}
L94:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v350 = m.T0[v349].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), v344, v323+int32(_a_F_ldbCatStackValueRec_10), int32(2))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L19
	} else {
		goto L95
	}
L95:
	;
	v354 = F_ldbCatStackValueRec(m, v344, l1, int32(-1), v232)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L19
	} else {
		goto L96
	}
L96:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v358 = m.T0[v357].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), v354, v330, int32(2))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L19
	} else {
		goto L97
	}
L97:
	;
	goto L100
L98:
	;
	v390 = F_lua_next(m, l1, v227)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L19
	} else {
		goto L106
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v378 + int32(-16)
	goto L98
L100:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	goto L99
L106:
	;
	if v390 != 0 {
		v240 = v354
		v242 = v327
		v245 = v245 + int32(1)
		v246 = v322
		goto L68
	} else {
		goto L107
	}
L107:
	;
	goto L69
L108:
	;
	v416 = m.G7
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	v418 = m.T0[v417].(func(*base.Module, int32, int32) int32)(m, v399, v16+int32(44))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L19
	} else {
		goto L109
	}
L109:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	v424 = m.T0[v423].(func(*base.Module, int32, int32) int32)(m, v397, v16+int32(40))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L19
	} else {
		goto L110
	}
L110:
	;
	if v401 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v426 = v424
	goto L113
L112:
	;
	v426 = v418
	goto L113
L113:
	;
	if v401 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v431 = v16 + int32(40)
	goto L116
L115:
	;
	v431 = v16 + int32(44)
	goto L116
L116:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v434 = m.T0[v433].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), l0, v426, v432)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L19
	} else {
		goto L117
	}
L117:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v441 = m.T0[v440].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), l0, v406+int32(_a_F_ldbCatStackValueRec_11), int32(1))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L19
	} else {
		goto L118
	}
L118:
	;
	v444 = m.G17
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	m.T0[v445].(func(*base.Module, int32, int32))(m, int32(0), v399)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L19
	} else {
		goto L119
	}
L119:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	m.T0[v449].(func(*base.Module, int32, int32))(m, int32(0), v397)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L19
	} else {
		goto L120
	}
L120:
	;
	v627 = l0
	goto L16
L121:
	;
	v582 = v74 + int32(-2)
	if base.Ui32(v582) <= base.Ui32(int32(6)) {
		goto L158
	} else {
		goto L159
	}
L122:
	;
	v507 = int32(0)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v505)+8))
	switch v508 + int32(-2) {
	case 0, 5:
		goto L138
	default:
		v575 = v507
		goto L137
	case 3, 4, 6:
		goto L139
	}
L123:
	;
	if l2 < int32(-9999) {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v463 = v458 + l2<<(uint(int32(4))%32) + int32(-16)
	v464 = m.G398
	if base.Ui32(v463) < base.Ui32(v457) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v466 = v463
	goto L127
L126:
	;
	v466 = v464
	goto L127
L127:
	;
	v505 = v466
	goto L122
L128:
	;
	switch l2 + int32(10002) {
	case 0:
		goto L131
	case 1:
		goto L132
	case 2:
		goto L133
	default:
		goto L130
	}
L129:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v505 = v469 + l2<<(uint(int32(4))%32)
	goto L122
L130:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+7)))
	v493 = m.G398
	if base.Ui32(v492) < base.Ui32(int32(-10002)-l2) {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	v505 = l1 + int32(72)
	goto L122
L132:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v481
	v505 = l1 + int32(88)
	goto L122
L133:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v505 = v475 + int32(96)
	goto L122
L134:
	;
	v504 = v493
	goto L136
L135:
	;
	v504 = v491 + (int32(-10003)-l2)<<(uint(int32(4))%32) + int32(24)
	goto L136
L136:
	;
	v505 = v504
	goto L122
L137:
	;
	v580 = v575
	goto L121
L138:
	;
	if l2 < int32(1) {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v505)))
	v580 = v511
	goto L121
L140:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v563)+8))
	switch v565 + int32(-2) {
	case 0:
		goto L155
	default:
		v575 = v507
		goto L137
	case 5:
		goto L156
	}
L141:
	;
	if l2 < int32(-9999) {
		goto L146
	} else {
		goto L147
	}
L142:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v520 = v515 + l2<<(uint(int32(4))%32) + int32(-16)
	v521 = m.G398
	if base.Ui32(v520) < base.Ui32(v514) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v523 = v520
	goto L145
L144:
	;
	v523 = v521
	goto L145
L145:
	;
	v563 = v523
	goto L140
L146:
	;
	switch l2 + int32(10002) {
	case 0:
		goto L149
	case 1:
		goto L150
	case 2:
		goto L151
	default:
		goto L148
	}
L147:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v563 = v526 + l2<<(uint(int32(4))%32)
	goto L140
L148:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+7)))
	v550 = m.G398
	if base.Ui32(v549) < base.Ui32(int32(-10002)-l2) {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v563 = l1 + int32(72)
	goto L140
L150:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v537)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v538
	v563 = l1 + int32(88)
	goto L140
L151:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v563 = v532 + int32(96)
	goto L140
L152:
	;
	v561 = v550
	goto L154
L153:
	;
	v561 = v548 + (int32(-10003)-l2)<<(uint(int32(4))%32) + int32(24)
	goto L154
L154:
	;
	v563 = v561
	goto L140
L155:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	v575 = v571
	goto L137
L156:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	v580 = v568 + int32(24)
	goto L121
L157:
	;
	v596 = m.G7
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v596)))
	v598 = m.T0[v597].(func(*base.Module, int32, int32) int32)(m, l0, int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L19
	} else {
		goto L160
	}
L158:
	;
	v587 = m.G3
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v587+int32(_a_F_ldbCatStackValueRec_12)+v582<<(uint(int32(2))%32))))
	v594 = v593
	goto L157
L159:
	;
	v594 = v452 + int32(_a_F_ldbCatStackValueRec_13)
	goto L157
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v594
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v580
	v603 = m.G3
	v604 = m.G15
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v604)))
	v606 = m.G17
	v612 = m.T0[v605].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v603+int32(_a_F_ldbCatStackValueRec_14), v16+int32(16))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L19
	} else {
		goto L161
	}
L161:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v606)))
	m.T0[v615].(func(*base.Module, int32, int32))(m, int32(0), l0)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L19
	} else {
		goto L162
	}
L162:
	;
	v627 = v612
	goto L16
L163:
	;
	v627 = l0
	goto L16
}
func F_ldbGenerateDebuggerCommandsArray(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v5 = m.G3
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_ldbGenerateDebuggerCommandsArray[0])))
	if v8 != 0 {
		v30 = v8
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v30
		v35 = m.G3
		v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_ldbGenerateDebuggerCommandsArray[1]))))
		if v38 != 0 {
			v39 = int32(13)
		} else {
			v39 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v39
		return
	} else {
		v9 = m.G3
		v12 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_ldbGenerateDebuggerCommandsArray[1]))) = uint8(v12)
		v18 = m.G9
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		v20 = m.T0[v19].(func(*base.Module, int32, int32) int32)(m, int32(13), int32(40))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_ldbGenerateDebuggerCommandsArray[0]))) = v20
			v28 = F__emscripten_memcpy_bulkmem(m, v20, v9+int32(_a_F_ldbGenerateDebuggerCommandsArray_0), int32(520))
			mBase = m.M
			v30 = v20
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v30
			v35 = m.G3
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_ldbGenerateDebuggerCommandsArray[1]))))
			if v38 != 0 {
				v39 = int32(13)
			} else {
				v39 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v39
			return
		}
	}
}
func F_ldbGetCurrentLine(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	v1 = m.G6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(v1)+280))
	return v2
}
func F_ldbIsActive(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	v1 = m.G6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(v1)))
	return v2
}
func F_ldbIsBreakpointOnNextLineEnabled(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	v1 = m.G6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(v1)+268))
	return v2
}
func F_ldbIsStepEnabled(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	v1 = m.G6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(v1)+264))
	return v2
}
func F_ldbLogSourceLine(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = m.G3
	v13 = v11 + int32(_a_F_ldbLogSourceLine_0)
	v14 = m.G6
	if l0 < int32(1) {
		v27 = v13
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = m.G3
	v29 = int32(0)
	v30 = m.G6
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+260))
	if v29 < v31 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+276))
	if v17 < l0 {
		v27 = v13
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = m.G6
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+272))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20+l0<<(uint(int32(2))%32)+int32(-4))))
	v27 = v26
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
	v71 = m.G6
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+280))
	if v72 == l0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v41 = v29
	goto L8
L6:
	;
	v66 = v28 + int32(_a_F_ldbLogSourceLine_1)
	v68 = v28 + int32(_a_F_ldbLogSourceLine_2)
	goto L4
L7:
	;
	v66 = v44 + int32(_a_F_ldbLogSourceLine_3)
	v68 = v44 + int32(_a_F_ldbLogSourceLine_4)
	goto L4
L8:
	;
	v44 = m.G3
	v45 = m.G6
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v41<<(uint(int32(2))%32))+4))
	if v49 == l0 {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v66 = v51 + int32(_a_F_ldbLogSourceLine_1)
	v68 = v51 + int32(_a_F_ldbLogSourceLine_2)
	goto L4
L10:
	;
	v51 = m.G3
	v53 = v41 + int32(1)
	if v53 != v31 {
		v41 = v53
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v74 = v66
	goto L14
L13:
	;
	v74 = v68
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v74
	v76 = m.G3
	v77 = m.G15
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v79 = m.G12
	v83 = m.T0[v78].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v76+int32(_a_F_ldbLogSourceLine_5), v9)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	m.T0[v86].(func(*base.Module, int32, int32))(m, v83, int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	m.G0 = v9 + int32(16)
	return
}
func F_ldbRepl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v2
	v18 = m.G20
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	m.T0[v19].(func(*base.Module, int32, int32))(m, v7+int32(12), v7+int32(8))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		if v24 == int32(0) {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			if v41 == int32(0) {
				v48 = v2
			} else {
				v44 = m.G6
				*(*int64)(unsafe.Add(mBase, uint32(v44)+260)) = int64(0)
				v48 = int32(-1)
			}
			m.G0 = v7 + int32(16)
			return v48
		} else {
			v28 = m.G7
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			v30 = m.T0[v29].(func(*base.Module, int32, int32) int32)(m, v24, int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_luaPushError(m, l0, v30)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = m.G11
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					m.T0[v36].(func(*base.Module, int32))(m, v35)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = F_luaError(m, l0)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v48 = v2
							m.G0 = v7 + int32(16)
							return v48
						}
					}
				}
			}
		}
	}
}
func F_ldbSetBreakpointOnNextLine(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = m.G6
	*(*int32)(unsafe.Add(mBase, uint32(v2)+268)) = l0
	return
}
func F_ldbSetStepMode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = m.G6
	*(*int32)(unsafe.Add(mBase, uint32(v2)+264)) = l0
	return
}
func F_ldbShouldBreak(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v6 = m.G6
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+260))
	if v7 < int32(1) {
		v25 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = m.G6
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+268))
	return v25 | base.B2i32(v30 != int32(0))
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+280))
	v13 = int32(0)
	goto L3
L3:
	;
	v16 = m.G6
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v13<<(uint(int32(2))%32))+4))
	v21 = base.B2i32(v20 == v10)
	if v20 == v10 {
		v25 = v21
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v25 = v21
	goto L1
L5:
	;
	v23 = v13 + int32(1)
	if v23 != v7 {
		v13 = v23
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
}
