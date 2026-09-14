package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_functionsCreateWithLibraryCtx(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
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
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v322 int32
	_ = v322
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v359 int32
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
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int64
	_ = v572
	var v573 int64
	_ = v573
	var v574 int64
	_ = v574
	var v575 int64
	_ = v575
	var v576 int64
	_ = v576
	var v577 int64
	_ = v577
	var v578 int64
	_ = v578
	var v580 int64
	_ = v580
	var v582 int64
	_ = v582
	var v584 int64
	_ = v584
	var v586 int64
	_ = v586
	var v588 int64
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
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
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v817 int32
	_ = v817
	v20 = m.G0
	v22 = v20 - int32(64)
	m.G0 = v22
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(56)))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = int64(0)
	v32 = F_functionExtractLibMetaData(m, l0, v22+int32(48), l2)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v22 + int32(64)
	return v817
L2:
	;
	v817 = int32(0)
	goto L1
L3:
	;
	return int32(0)
L4:
	;
	if v32 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+int32(-1)))))
	v41 = v39 & int32(7)
	switch v41 {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	default:
		goto L8
	}
L6:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	if v777 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L7:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v140 = F_scriptingEngineManagerFind(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L31
	}
L8:
	;
	v136 = F_sdsnew(m, int32(_a636))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L3
	} else {
		goto L29
	}
L9:
	;
	if v56 == int32(0) {
		goto L8
	} else {
		goto L15
	}
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(-17))))
	v56 = v55
	goto L9
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(-9))))
	v56 = v52
	goto L9
L12:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36+int32(-5)))))
	v56 = v49
	goto L9
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+int32(-3)))))
	v56 = v46
	goto L9
L14:
	;
	v56 = int32(base.Ui32(v39) >> (uint(int32(3)) % 32))
	goto L9
L15:
	;
	v79 = int32(0)
	goto L16
L16:
	;
	switch v41 {
	case 0:
		goto L23
	case 1:
		goto L22
	case 2:
		goto L21
	case 3:
		goto L20
	case 4:
		goto L19
	default:
		v94 = int32(0)
		goto L18
	}
L18:
	;
	if base.Ui32(v94) <= base.Ui32(v79) {
		goto L7
	} else {
		goto L24
	}
L19:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(-17))))
	v94 = v93
	goto L18
L20:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(-9))))
	v94 = v92
	goto L18
L21:
	;
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36+int32(-5)))))
	v94 = v91
	goto L18
L22:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+int32(-3)))))
	v94 = v90
	goto L18
L23:
	;
	v94 = int32(base.Ui32(v39) >> (uint(int32(3)) % 32))
	goto L18
L24:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v79))))
	if base.Ui32((v97&int32(223)+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v79 = v79 + int32(1)
	goto L16
L26:
	;
	if v97 == int32(95) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(int32(9)) < base.Ui32((v97+int32(-48))&int32(255)) {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v136
	goto L6
L30:
	;
	v149 = int32(0)
	v150 = *(*int32)(unsafe.Add(mBase, _consts[332]))
	if v150 == v149 {
		goto L41
	} else {
		goto L42
	}
L31:
	;
	if v140 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v142 = F_sdsempty(m)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v139
	v146 = F_sdscatfmt(m, v142, int32(_a637), v22)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v146
	goto L6
L35:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	F_dictRelease(m, v743)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L3
	} else {
		goto L180
	}
L36:
	;
	F__serverAssert(m, int32(_a638), int32(_a639), int32(278))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L3
	} else {
		goto L179
	}
L37:
	;
	F__serverAssert(m, int32(_a640), int32(_a639), int32(277))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L3
	} else {
		goto L178
	}
L38:
	;
	F__serverAssert(m, int32(_a641), int32(_a639), int32(1055))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L3
	} else {
		goto L177
	}
L39:
	;
	F__serverAssert(m, int32(_a642), int32(_a639), int32(1049))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L3
	} else {
		goto L176
	}
L40:
	;
	F__serverAssert(m, int32(_a643), int32(_a639), int32(1048))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L3
	} else {
		goto L175
	}
L41:
	;
	F__serverAssert(m, int32(_a644), int32(_a639), int32(257))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L3
	} else {
		goto L174
	}
L42:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	v154 = F_dictFind(m, v153, v139)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L44
	}
L43:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v166 = F_dictFetchValue(m, v165, v36)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L3
	} else {
		goto L48
	}
L44:
	;
	if v154 != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v157 = F_valkey_calloc(m, int32(8))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _consts[332]))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	v162 = F_dictAdd(m, v161, v139, v157)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	if l1 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v166 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	if v166 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v170 = F_sdsempty(m)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v36
	v176 = F_sdscatfmt(m, v170, int32(_a645), v22+int32(32))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v176
	goto L6
L54:
	;
	v184 = F_valkey_malloc(m, int32(16))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L3
	} else {
		goto L57
	}
L55:
	;
	F_libraryUnlink(m, l3, v166)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v186 = F_sdsdup(m, v36)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v189 = F_dictCreate(m, int32(_a646))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	v191 = F_sdsdup(m, l0)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184)+12)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v184)+8)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v186
	v197 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v197
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202+int32(-1)))))
	switch v205 & int32(7) {
	case 0:
		goto L66
	case 1:
		goto L65
	case 2:
		goto L64
	case 3:
		goto L63
	case 4:
		goto L62
	default:
		v222 = v197
		goto L61
	}
L61:
	;
	v228 = F_scriptingEngineCallCompileCode(m, v140, int32(1), v202, v222, l4, v22+int32(44), v22+int32(40))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L3
	} else {
		goto L68
	}
L62:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v202+int32(-17))))
	v222 = v221
	goto L61
L63:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v202+int32(-9))))
	v222 = v218
	goto L61
L64:
	;
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202+int32(-5)))))
	v222 = v215
	goto L61
L65:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202+int32(-3)))))
	v222 = v212
	goto L61
L66:
	;
	v222 = int32(base.Ui32(v205) >> (uint(int32(3)) % 32))
	goto L61
L67:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v241 != 0 {
		goto L38
	} else {
		goto L74
	}
L68:
	;
	if v228 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	if v230 != 0 {
		goto L40
	} else {
		goto L70
	}
L70:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v231 == int32(0) {
		goto L39
	} else {
		goto L71
	}
L71:
	;
	v234 = F_objectGetVal(m, v231)
	mBase = m.M
	v235 = F_sdsdup(m, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v235
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	F_decrRefCount(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	goto L35
L74:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	if v242 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_valkey_free(m, v228)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L3
	} else {
		goto L126
	}
L76:
	;
	v247 = int32(0)
	goto L77
L77:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v228+v247<<(uint(int32(2))%32))))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+8))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	if v270&int32(15) != 0 {
		goto L37
	} else {
		goto L79
	}
L78:
	;
	goto L75
L79:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v268)+16))
	if v273 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v279 = F_objectGetVal(m, v269)
	mBase = m.M
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+int32(-1)))))
	v284 = v282 & int32(7)
	switch v284 {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	case 3:
		goto L88
	case 4:
		goto L87
	default:
		goto L85
	}
L81:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	if v276&int32(15) != 0 {
		goto L36
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v485 = v247 + int32(1)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	if base.Ui32(v485) < base.Ui32(v486) {
		v247 = v485
		goto L77
	} else {
		goto L125
	}
L84:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	if base.Ui32(v432) <= base.Ui32(v247) {
		goto L118
	} else {
		goto L119
	}
L85:
	;
	v410 = F_sdsnew(m, int32(_a647))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L3
	} else {
		goto L117
	}
L86:
	;
	if v299 == int32(0) {
		goto L85
	} else {
		goto L92
	}
L87:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v279+int32(-17))))
	v299 = v298
	goto L86
L88:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v279+int32(-9))))
	v299 = v295
	goto L86
L89:
	;
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279+int32(-5)))))
	v299 = v292
	goto L86
L90:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+int32(-3)))))
	v299 = v289
	goto L86
L91:
	;
	v299 = int32(base.Ui32(v282) >> (uint(int32(3)) % 32))
	goto L86
L92:
	;
	v322 = int32(0)
	goto L93
L93:
	;
	switch v284 {
	case 0:
		goto L100
	case 1:
		goto L99
	case 2:
		goto L98
	case 3:
		goto L97
	case 4:
		goto L96
	default:
		v337 = int32(0)
		goto L95
	}
L94:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v268)+8))
	v360 = F_objectGetVal(m, v359)
	mBase = m.M
	v361 = F_sdsdup(m, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L3
	} else {
		goto L107
	}
L95:
	;
	if base.Ui32(v337) <= base.Ui32(v322) {
		goto L101
	} else {
		goto L102
	}
L96:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v279+int32(-17))))
	v337 = v336
	goto L95
L97:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v279+int32(-9))))
	v337 = v335
	goto L95
L98:
	;
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279+int32(-5)))))
	v337 = v334
	goto L95
L99:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+int32(-3)))))
	v337 = v333
	goto L95
L100:
	;
	v337 = int32(base.Ui32(v282) >> (uint(int32(3)) % 32))
	goto L95
L101:
	;
	goto L94
L102:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v322))))
	if base.Ui32((v340&int32(223)+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v322 = v322 + int32(1)
	goto L93
L104:
	;
	if v340 == int32(95) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	if base.Ui32(int32(9)) < base.Ui32((v340+int32(-48))&int32(255)) {
		goto L85
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v364 = F_dictFetchValue(m, v363, v361)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L3
	} else {
		goto L109
	}
L108:
	;
	v375 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L3
	} else {
		goto L113
	}
L109:
	;
	if v364 == int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v369 = F_sdsnew(m, int32(_a648))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v369
	F_sdsfree(m, v361)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	goto L84
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v375)+4)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v375))) = v268
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v380 = F_dictAdd(m, v379, v361, v375)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L3
	} else {
		goto L114
	}
L114:
	;
	if v380 == int32(0) {
		goto L83
	} else {
		goto L115
	}
L115:
	;
	F__serverAssert(m, int32(_a649), int32(_a639), int32(300))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v410
	goto L84
L118:
	;
	F_valkey_free(m, v228)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L3
	} else {
		goto L124
	}
L119:
	;
	v435 = v247
	goto L120
L120:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v228+v435<<(uint(int32(2))%32))))
	F_scriptingEngineCallFreeFunction(m, v140, int32(1), v457)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L3
	} else {
		goto L122
	}
L121:
	;
	goto L118
L122:
	;
	v461 = v435 + int32(1)
	if v461 != v432 {
		v435 = v461
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	goto L35
L125:
	;
	goto L78
L126:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+12))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v509)+16))
	if v510 != int32(0)-v512 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v519 = F_dictGetIterator(m, v509)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L3
	} else {
		goto L130
	}
L128:
	;
	v516 = F_sdsnew(m, int32(_a650))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v516
	goto L35
L130:
	;
	goto L132
L131:
	;
	F_dictReleaseIterator(m, v519)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L3
	} else {
		goto L168
	}
L132:
	;
	v547 = v519 + int32(20)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v519)+16))
	if v548 != 0 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v658 = F_sdsempty(m)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L3
	} else {
		goto L164
	}
L134:
	;
	if v643 == int32(0) {
		goto L131
	} else {
		goto L160
	}
L135:
	;
	v554 = v547
	v555 = v551
	goto L138
L136:
	;
	v551 = int32(1)
	goto L135
L137:
	;
	v551 = int32(0)
	goto L135
L138:
	;
	switch v555 {
	case 0:
		goto L143
	default:
		goto L142
	}
L140:
	;
	v555 = int32(0)
	goto L138
L141:
	;
	goto L134
L142:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v554)))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+16)) = v635
	if v635 == int32(0) {
		goto L140
	} else {
		goto L159
	}
L143:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v519)+4))
	if v559 != int32(-1) {
		v598 = v559
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v599 = int32(1)
	v600 = v598 + v599
	*(*int32)(unsafe.Add(mBase, uint32(v519)+4)) = v600
	v602 = int32(0)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v519)+8))
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605+v606+int32(26)))))
	if v610 == int32(255) {
		goto L153
	} else {
		goto L154
	}
L145:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v519)+8))
	if v563 != 0 {
		v598 = int32(-1)
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v519)+12))
	if v565 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v591)+20))
	if v592 != int32(-1) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v572 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v564)+16)))
	v573 = int64(*(*int8)(unsafe.Add(mBase, uint32(v564)+27)))
	v574 = int64(*(*int32)(unsafe.Add(mBase, uint32(v564)+8)))
	v575 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v564)+12)))
	v576 = int64(*(*int8)(unsafe.Add(mBase, uint32(v564)+26)))
	v577 = int64(*(*int32)(unsafe.Add(mBase, uint32(v564)+4)))
	v578 = F_wangHash64(m, v577)
	mBase = m.M
	v580 = F_wangHash64(m, v576+v578)
	mBase = m.M
	v582 = F_wangHash64(m, v575+v580)
	mBase = m.M
	v584 = F_wangHash64(m, v574+v582)
	mBase = m.M
	v586 = F_wangHash64(m, v573+v584)
	mBase = m.M
	v588 = F_wangHash64(m, v572+v586)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v519)+24)) = v588
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	v591 = v590
	goto L147
L149:
	;
	v568 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v564)+24)))
	v570 = v568 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v564)+24)) = uint16(v570)
	v591 = v564
	goto L147
L150:
	;
	v598 = v592 + int32(-1)
	goto L144
L151:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v519)+4))
	v598 = v595
	goto L144
L152:
	;
	v625 = int32(2)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v605+v623<<(uint(v625)%32)+int32(4))))
	v554 = v630 + v624<<(uint(v625)%32)
	v555 = int32(1)
	goto L138
L153:
	;
	v614 = v602
	goto L155
L154:
	;
	v614 = v599 << (uint(v610) % 32)
	goto L155
L155:
	;
	if v600 < v614 {
		v623 = v606
		v624 = v600
		goto L152
	} else {
		goto L156
	}
L156:
	;
	if v606 != 0 {
		v643 = v602
		goto L141
	} else {
		goto L157
	}
L157:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v605)+20))
	if v616 == int32(-1) {
		v643 = v602
		goto L141
	} else {
		goto L158
	}
L158:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+4)) = int64(4294967296)
	v623 = int32(1)
	v624 = int32(0)
	goto L152
L159:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v635)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v547))) = v639
	v643 = v635
	goto L141
L160:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v643)+8))
	goto L161
L161:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v651)+8))
	v653 = F_objectGetVal(m, v652)
	mBase = m.M
	v654 = F_dictFetchValue(m, v650, v653)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	if v654 == int32(0) {
		goto L132
	} else {
		goto L163
	}
L163:
	;
	goto L133
L164:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v660)+8))
	v662 = F_objectGetVal(m, v661)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v662
	v667 = F_sdscatfmt(m, v658, int32(_a651), v22+int32(16))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L3
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v667
	if v519 == int32(0) {
		goto L35
	} else {
		goto L166
	}
L166:
	;
	F_dictReleaseIterator(m, v519)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L3
	} else {
		goto L167
	}
L167:
	;
	goto L35
L168:
	;
	F_libraryLink(m, l3, v184)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L3
	} else {
		goto L169
	}
L169:
	;
	if v166 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = int32(0)
	F_functionFreeLibMetaData(m, v22+int32(48))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L3
	} else {
		goto L173
	}
L171:
	;
	F_engineLibraryFree(m, v166)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L3
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v817 = v36
	goto L1
L174:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	F_sdsfree(m, v746)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L3
	} else {
		goto L181
	}
L181:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	F_sdsfree(m, v749)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L3
	} else {
		goto L182
	}
L182:
	;
	F_valkey_free(m, v184)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L3
	} else {
		goto L183
	}
L183:
	;
	if v166 == int32(0) {
		goto L6
	} else {
		goto L184
	}
L184:
	;
	F_libraryLink(m, l3, v166)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L3
	} else {
		goto L185
	}
L185:
	;
	goto L6
L186:
	;
	if v36 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	F_sdsfree(m, v777)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L3
	} else {
		goto L188
	}
L188:
	;
	goto L186
L189:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v786 == int32(0) {
		goto L2
	} else {
		goto L192
	}
L190:
	;
	F_sdsfree(m, v36)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L3
	} else {
		goto L191
	}
L191:
	;
	goto L189
L192:
	;
	F_sdsfree(m, v786)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L3
	} else {
		goto L193
	}
L193:
	;
	goto L2
}
func F_functionsInit(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v3 = F_valkey_malloc(m, int32(16))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_dictCreate(m, int32(_a617))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3))) = v8
			v12 = F_dictCreate(m, int32(_a618))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v3)+4)) = v12
				v16 = F_dictCreate(m, int32(_a619))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v16
					F_scriptingEngineManagerForEachEngine(m, int32(529), v3)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = v22
						*(*int32)(unsafe.Add(mBase, _consts[332])) = v3
						return v22
					}
				}
			}
		}
	}
}
func F_functionsLibCtxFree(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	F_functionsLibCtxClear(m, l0, l1)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_dictRelease(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_dictRelease(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_dictRelease(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if l2 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v20 = F_listGetIterator(m, l2, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L11
L10:
	;
	F_listReleaseIterator(m, v20)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L20
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v26 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v26 == int32(0) {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26+base.B2i32(v29 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v35
	goto L14
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v39 == int32(0) {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	m.T0[v43].(func(*base.Module, int32))(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_valkey_free(m, v39)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L11
L20:
	;
	F_listRelease(m, l2)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L7
}
func F_functionsLibCtxReleaseCurrent(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	if l0 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[332]))
		F_functionsLibCtxClear(m, v15, l1)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			F_dictRelease(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				F_dictRelease(m, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
					F_dictRelease(m, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_valkey_free(m, v15)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							F_scriptingEngineManagerForEachEngine(m, int32(528), int32(0))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	} else {
		v6 = F_listCreate(m)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_scriptingEngineManagerForEachEngine(m, int32(528), v6)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, _consts[332]))
				F_freeFunctionsAsync(m, v11, v6)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_functionsLibGet(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[332]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	return v3
}
