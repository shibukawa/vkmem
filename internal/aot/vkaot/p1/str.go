package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_str_dump(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	v3 = m.G0
	v5 = v3 - int32(1040)
	m.G0 = v5
	F_luaL_checktype(m, l0, int32(1), int32(6))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v21 = v18 + int32(16)
		if base.Ui32(v21) <= base.Ui32(v17) {
		} else {
			v25 = v17
			for {
				*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(0)
				v29 = v25 + int32(16)
				if base.Ui32(v29) < base.Ui32(v21) {
					v25 = v29
					continue
				} else {
					break
				}
				break
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v21
		v42 = v5 + int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v42))) = v5 + int32(16)
		v49 = m.G5
		v54 = F_lua_dump(m, l0, v49+int32(1327), v5+int32(4))
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			if v54 == int32(0) {
				F_luaL_pushresult(m, v5+int32(4))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(1040)
					return int32(1)
				}
			} else {
				v58 = m.G3
				v62 = F_luaL_error(m, l0, v58+int32(_a2720), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					F_luaL_pushresult(m, v5+int32(4))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						m.G0 = v5 + int32(1040)
						return int32(1)
					}
				}
			}
		}
	}
}
func F_str_format(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
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
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
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
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
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
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 float64
	_ = v242
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
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
	var v328 int32
	_ = v328
	var v331 float64
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v420 float64
	_ = v420
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v468 int32
	_ = v468
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 float64
	_ = v603
	var v604 int32
	_ = v604
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int64
	_ = v737
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	v14 = m.G0
	v16 = v14 - int32(1680)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L1
L1:
	;
	v26 = F_luaL_checklstring(m, l0, int32(1), v16+int32(1672))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1672))
	v32 = v16 + int32(636)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v16 + int32(648)
	goto L4
L4:
	;
	if v30 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	m.G0 = v16 + int32(1680)
	return v785
L6:
	;
	F_luaL_pushresult(m, v16+int32(636))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L2
	} else {
		goto L182
	}
L7:
	;
	v44 = int32(1)
	v45 = v16 + int32(608) | v44
	v47 = v16 + int32(1672)
	v52 = v26
	v57 = v44
	goto L8
L8:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v62 == int32(37) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L6
L10:
	;
	if base.Ui32(v752) < base.Ui32(v26+v30) {
		v52 = v752
		v57 = v757
		goto L8
	} else {
		goto L181
	}
L11:
	;
	v82 = v52 + int32(1)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v83 != int32(37) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	if base.Ui32(v65) < base.Ui32(v47) {
		v73 = v62
		v74 = v65
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v75 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+636)) = v74 + v75
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v73)
	v752 = v52 + v75
	v757 = v57
	goto L10
L14:
	;
	v69 = F_luaL_prepbuffer(m, v16+int32(636))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	v73 = v71
	v74 = v72
	goto L13
L16:
	;
	v104 = v57 + int32(1)
	if v57 < (v18-v19)>>(uint(int32(4))%32) {
		v112 = v83
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	if base.Ui32(v87) < base.Ui32(v47) {
		v95 = v87
		v96 = int32(37)
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+636)) = v95 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v96)
	v752 = v52 + int32(2)
	v757 = v57
	goto L10
L19:
	;
	v91 = F_luaL_prepbuffer(m, v16+int32(636))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	v95 = v94
	v96 = v93
	goto L18
L21:
	;
	v113 = int32(0)
	if v112&int32(255) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v106 = m.G3
	v109 = F_luaL_argerror(m, l0, v104, v106+int32(_a2694))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v112 = v111
	goto L21
L24:
	;
	if base.Ui32(v150-v82) < base.Ui32(int32(6)) {
		v170 = v158
		goto L34
	} else {
		goto L35
	}
L25:
	;
	v119 = v82
	v120 = v112
	goto L27
L26:
	;
	v150 = v82
	v158 = v113
	goto L24
L27:
	;
	if base.Ui32(v120&int32(255)) <= base.Ui32(int32(63)) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v150 = v146
	v158 = v113
	goto L24
L29:
	;
	if base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v120)&int64(255))%64)&int64(325494096527361) == int64(0)) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v150 = v119
	v158 = v120
	goto L24
L31:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	v146 = v119 + int32(1)
	if v144 != 0 {
		v119 = v146
		v120 = v144
		goto L27
	} else {
		goto L33
	}
L32:
	;
	v150 = v119
	v158 = v120
	goto L24
L33:
	;
	goto L28
L34:
	;
	v171 = int32(-48)
	v173 = int32(255)
	v175 = int32(10)
	v177 = v150 + base.B2i32(base.Ui32((v170+v171)&v173) < base.Ui32(v175))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v185 = v177 + base.B2i32(base.Ui32((v178+v171)&v173) < base.Ui32(v175))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if v186 != int32(46) {
		v214 = v186
		v215 = v185
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v163 = m.G3
	v167 = F_luaL_error(m, l0, v163+int32(_a2721), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v170 = v169
	goto L34
L37:
	;
	if base.Ui32(int32(9)) < base.Ui32((v214+int32(-48))&int32(255)) {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+1)))
	v199 = base.B2i32(base.Ui32((v193+int32(-48))&int32(255)) < base.Ui32(int32(10)))
	if base.Ui32((v193+int32(-48))&int32(255)) < base.Ui32(int32(10)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v200 = v185 + int32(2)
	goto L41
L40:
	;
	v200 = v185 + int32(1)
	goto L41
L41:
	;
	if base.Ui32((v193+int32(-48))&int32(255)) < base.Ui32(int32(10)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v203 = int32(2)
	goto L44
L43:
	;
	v203 = int32(1)
	goto L44
L44:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185+v203))))
	v212 = v200 + base.B2i32(base.Ui32((v205+int32(-48))&int32(255)) < base.Ui32(int32(10)))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	v214 = v213
	v215 = v212
	goto L37
L45:
	;
	v228 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+608)) = uint8(v228)
	v232 = v215 - v82 + int32(1)
	v233 = F___stpncpy(m, v45, v82, v232)
	mBase = m.M
	goto L48
L46:
	;
	v222 = m.G3
	v226 = F_luaL_error(m, l0, v222+int32(_a2722), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v235 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45+v232))) = uint8(v235)
	v238 = v215 + int32(1)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	switch v239 + int32(-69) {
	case 0, 2, 32, 33, 34:
		goto L51
	default:
		goto L52
	case 19, 42, 48, 51:
		goto L55
	case 30:
		goto L57
	case 31, 36:
		goto L56
	case 44:
		goto L54
	case 46:
		goto L53
	}
L49:
	;
	if v104 < int32(1) {
		goto L166
	} else {
		goto L167
	}
L50:
	;
	v621 = v16 + int32(96)
	if v621&int32(3) == int32(0) {
		v645 = v621
		goto L149
	} else {
		goto L150
	}
L51:
	;
	v603 = F_luaL_checknumber(m, l0, v104)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L2
	} else {
		goto L145
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = base.I32_extend8_s(v239)
	v598 = m.G3
	v601 = F_luaL_error(m, l0, v598+int32(_a2723), v16)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L2
	} else {
		goto L144
	}
L53:
	;
	v572 = F_luaL_checklstring(m, l0, v104, v16+int32(1676))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L2
	} else {
		goto L135
	}
L54:
	;
	v443 = F_luaL_checklstring(m, l0, v104, v16+int32(1676))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L2
	} else {
		goto L105
	}
L55:
	;
	v351 = v16 + int32(608)
	if v351&int32(3) == int32(0) {
		v375 = v351
		goto L86
	} else {
		goto L87
	}
L56:
	;
	v262 = v16 + int32(608)
	if v262&int32(3) == int32(0) {
		v286 = v262
		goto L65
	} else {
		goto L66
	}
L57:
	;
	v242 = F_luaL_checknumber(m, l0, v104)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L2
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v251
	v259 = F_siprintf(m, v16+int32(96), v16+int32(608), v16+int32(16))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L2
	} else {
		goto L62
	}
L59:
	;
	v251 = int32(-2147483648)
	goto L58
L60:
	;
	if base.F64_lt(base.F64_abs(v242), float64(2.147483648e+09)) == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v249 = base.I32_trunc_f64_s(v242)
	v251 = v249
	goto L58
L62:
	;
	goto L50
L63:
	;
	v320 = v262 + v319
	v322 = v320 + int32(-1)
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	v324 = int32(108)
	*(*uint16)(unsafe.Add(mBase, uint32(v322))) = uint16(v324)
	v328 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v320+int32(1)))) = uint8(v328)
	*(*uint8)(unsafe.Add(mBase, uint32(v320))) = uint8(v323)
	v331 = F_luaL_checknumber(m, l0, v104)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L2
	} else {
		goto L81
	}
L64:
	;
	v319 = v311 - v262
	goto L63
L65:
	;
	v290 = v286
	goto L73
L66:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	if v272 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v275 = v262
	goto L69
L68:
	;
	v319 = v262 - v262
	goto L63
L69:
	;
	v279 = v275 + int32(1)
	if v279&int32(3) == int32(0) {
		v286 = v279
		goto L65
	} else {
		goto L71
	}
L71:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	if v284 != 0 {
		v275 = v279
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v311 = v279
	goto L64
L73:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v299 = int32(-2139062144)
	if (int32(16843008)-v296|v296)&v299 == v299 {
		v290 = v290 + int32(4)
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v305 = v290
	goto L76
L75:
	;
	goto L74
L76:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	if v309 != 0 {
		v305 = v305 + int32(1)
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v311 = v305
	goto L64
L78:
	;
	goto L77
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v340
	v348 = F_siprintf(m, v16+int32(96), v16+int32(608), v16+int32(32))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L2
	} else {
		goto L83
	}
L80:
	;
	v340 = int32(-2147483648)
	goto L79
L81:
	;
	if base.F64_lt(base.F64_abs(v331), float64(2.147483648e+09)) == int32(0) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v338 = base.I32_trunc_f64_s(v331)
	v340 = v338
	goto L79
L83:
	;
	goto L50
L84:
	;
	v409 = v351 + v408
	v411 = v409 + int32(-1)
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	v413 = int32(108)
	*(*uint16)(unsafe.Add(mBase, uint32(v411))) = uint16(v413)
	v417 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v409+int32(1)))) = uint8(v417)
	*(*uint8)(unsafe.Add(mBase, uint32(v409))) = uint8(v412)
	v420 = F_luaL_checknumber(m, l0, v104)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L2
	} else {
		goto L102
	}
L85:
	;
	v408 = v400 - v351
	goto L84
L86:
	;
	v379 = v375
	goto L94
L87:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
	if v361 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v364 = v351
	goto L90
L89:
	;
	v408 = v351 - v351
	goto L84
L90:
	;
	v368 = v364 + int32(1)
	if v368&int32(3) == int32(0) {
		v375 = v368
		goto L86
	} else {
		goto L92
	}
L92:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	if v373 != 0 {
		v364 = v368
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v400 = v368
	goto L85
L94:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v388 = int32(-2139062144)
	if (int32(16843008)-v385|v385)&v388 == v388 {
		v379 = v379 + int32(4)
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v394 = v379
	goto L97
L96:
	;
	goto L95
L97:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	if v398 != 0 {
		v394 = v394 + int32(1)
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v400 = v394
	goto L85
L99:
	;
	goto L98
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v431
	v439 = F_siprintf(m, v16+int32(96), v16+int32(608), v16+int32(48))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L2
	} else {
		goto L104
	}
L101:
	;
	v431 = int32(0)
	goto L100
L102:
	;
	if base.F64_lt(v420, float64(4.294967296e+09))&base.F64_ge(v420, float64(0)) == int32(0) {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v429 = base.I32_trunc_f64_u(v420)
	v431 = v429
	goto L100
L104:
	;
	goto L50
L105:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	if base.Ui32(v445) < base.Ui32(v47) {
		v452 = v445
		goto L106
	} else {
		goto L107
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+636)) = v452 + int32(1)
	v456 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v452))) = uint8(v456)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1676))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1676)) = v458 + int32(-1)
	if v458 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v449 = F_luaL_prepbuffer(m, v16+int32(636))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L2
	} else {
		goto L108
	}
L108:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	v452 = v451
	goto L106
L109:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	if base.Ui32(v557) < base.Ui32(v47) {
		v564 = v557
		goto L132
	} else {
		goto L133
	}
L110:
	;
	v468 = v443
	goto L111
L111:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	switch v477 {
	case 0:
		goto L115
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33:
		goto L114
	case 10, 34:
		goto L117
	case 13:
		goto L116
	default:
		goto L118
	}
L112:
	;
	goto L109
L113:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1676))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1676)) = v538 + int32(-1)
	if v538 != 0 {
		v468 = v468 + int32(1)
		goto L111
	} else {
		goto L131
	}
L114:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	if base.Ui32(v522) < base.Ui32(v47) {
		v530 = v477
		v531 = v522
		goto L128
	} else {
		goto L129
	}
L115:
	;
	v516 = m.G3
	F_luaL_addlstring(m, v16+int32(636), v516+int32(_a2724), int32(4))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L2
	} else {
		goto L127
	}
L116:
	;
	v508 = m.G3
	F_luaL_addlstring(m, v16+int32(636), v508+int32(_a2725), int32(2))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L2
	} else {
		goto L126
	}
L117:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	if base.Ui32(v480) < base.Ui32(v47) {
		v487 = v480
		goto L120
	} else {
		goto L121
	}
L118:
	;
	if v477 != int32(92) {
		goto L114
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+636)) = v487 + int32(1)
	v491 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v487))) = uint8(v491)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	if base.Ui32(v493) < base.Ui32(v47) {
		v500 = v493
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v484 = F_luaL_prepbuffer(m, v16+int32(636))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	v487 = v486
	goto L120
L123:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+636)) = v500 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v500))) = uint8(v501)
	goto L113
L124:
	;
	v497 = F_luaL_prepbuffer(m, v16+int32(636))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	v500 = v499
	goto L123
L126:
	;
	goto L113
L127:
	;
	goto L113
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+636)) = v531 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v531))) = uint8(v530)
	goto L113
L129:
	;
	v526 = F_luaL_prepbuffer(m, v16+int32(636))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L2
	} else {
		goto L130
	}
L130:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	v530 = v528
	v531 = v529
	goto L128
L131:
	;
	goto L112
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+636)) = v564 + int32(1)
	v568 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v564))) = uint8(v568)
	v752 = v238
	v757 = v104
	goto L10
L133:
	;
	v561 = F_luaL_prepbuffer(m, v16+int32(636))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L2
	} else {
		goto L134
	}
L134:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v16)+636))
	v564 = v563
	goto L132
L135:
	;
	v576 = int32(46)
	v577 = F___strchrnul(m, v16+int32(608), v576)
	mBase = m.M
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577))))
	if v579 == v576 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v572
	v594 = F_siprintf(m, v16+int32(96), v16+int32(608), v16+int32(80))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L2
	} else {
		goto L143
	}
L137:
	;
	if v583 != 0 {
		goto L136
	} else {
		goto L141
	}
L138:
	;
	v583 = v577
	goto L140
L139:
	;
	v583 = int32(0)
	goto L140
L140:
	;
	goto L137
L141:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1676))
	if base.Ui32(int32(99)) < base.Ui32(v584) {
		goto L49
	} else {
		goto L142
	}
L142:
	;
	goto L136
L143:
	;
	goto L50
L144:
	;
	v785 = v601
	goto L5
L145:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16)+64)) = v603
	v612 = F___small_sprintf(m, v16+int32(96), v16+int32(608), v16+int32(64))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	goto L50
L147:
	;
	F_luaL_addlstring(m, v16+int32(636), v621, v678)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L2
	} else {
		goto L163
	}
L148:
	;
	v678 = v670 - v621
	goto L147
L149:
	;
	v649 = v645
	goto L157
L150:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621))))
	if v631 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v634 = v621
	goto L153
L152:
	;
	v678 = v621 - v621
	goto L147
L153:
	;
	v638 = v634 + int32(1)
	if v638&int32(3) == int32(0) {
		v645 = v638
		goto L149
	} else {
		goto L155
	}
L155:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638))))
	if v643 != 0 {
		v634 = v638
		goto L153
	} else {
		goto L156
	}
L156:
	;
	v670 = v638
	goto L148
L157:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	v658 = int32(-2139062144)
	if (int32(16843008)-v655|v655)&v658 == v658 {
		v649 = v649 + int32(4)
		goto L157
	} else {
		goto L159
	}
L158:
	;
	v664 = v649
	goto L160
L159:
	;
	goto L158
L160:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	if v668 != 0 {
		v664 = v664 + int32(1)
		goto L160
	} else {
		goto L162
	}
L161:
	;
	v670 = v664
	goto L148
L162:
	;
	goto L161
L163:
	;
	v752 = v238
	v757 = v104
	goto L10
L164:
	;
	F_luaL_addvalue(m, v16+int32(636))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L2
	} else {
		goto L180
	}
L165:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v737 = *(*int64)(unsafe.Add(mBase, uint32(v733)))
	*(*int64)(unsafe.Add(mBase, uint32(v736))) = v737
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v733)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+8)) = v739
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v741 + int32(16)
	goto L164
L166:
	;
	if v104 < int32(-9999) {
		goto L171
	} else {
		goto L172
	}
L167:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v691 = v686 + v104<<(uint(int32(4))%32) + int32(-16)
	v692 = m.G398
	if base.Ui32(v691) < base.Ui32(v685) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v694 = v691
	goto L170
L169:
	;
	v694 = v692
	goto L170
L170:
	;
	v733 = v694
	goto L165
L171:
	;
	switch v57 + int32(10003) {
	case 0:
		goto L174
	case 1:
		goto L175
	case 2:
		goto L176
	default:
		goto L173
	}
L172:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v733 = v697 + v104<<(uint(int32(4))%32)
	goto L165
L173:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+4))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v718)))
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v719)+7)))
	v721 = m.G398
	if base.Ui32(v720) < base.Ui32(int32(-10002)-v104) {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	v733 = l0 + int32(72)
	goto L165
L175:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v708)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v709
	v733 = l0 + int32(88)
	goto L165
L176:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v733 = v703 + int32(96)
	goto L165
L177:
	;
	v732 = v721
	goto L179
L178:
	;
	v732 = v719 + (int32(-10003)-v104)<<(uint(int32(4))%32) + int32(24)
	goto L179
L179:
	;
	v733 = v732
	goto L165
L180:
	;
	v752 = v238
	v757 = v104
	goto L10
L181:
	;
	goto L9
L182:
	;
	v785 = int32(1)
	goto L5
}
func F_str_lower(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
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
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(1040)
	m.G0 = v9
	v14 = F_luaL_checklstring(m, l0, int32(1), v9+int32(1036))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(12)
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1036))
	if v24 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_luaL_pushresult(m, v9)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L16
	}
L5:
	;
	v34 = int32(0)
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if base.Ui32(v36) < base.Ui32(v9+int32(1036)) {
		v41 = v36
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L4
L8:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v34))))
	if base.Ui32(v43+int32(-65)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v38 = F_luaL_prepbuffer(m, v9)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v41 = v40
	goto L8
L11:
	;
	v51 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v41 + v51
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v50)
	v56 = v34 + v51
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1036))
	if base.Ui32(v56) < base.Ui32(v57) {
		v34 = v56
		goto L6
	} else {
		goto L15
	}
L12:
	;
	v50 = v43 | int32(32)
	goto L14
L13:
	;
	v50 = v43
	goto L14
L14:
	;
	goto L11
L15:
	;
	goto L7
L16:
	;
	m.G0 = v9 + int32(1040)
	return int32(1)
}
func F_str_reverse(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	v6 = m.G0
	v8 = v6 - int32(1040)
	m.G0 = v8
	v13 = F_luaL_checklstring(m, l0, int32(1), v8+int32(1036))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(12)
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+1036))
	v25 = v23 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+1036)) = v25
	if v23 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_luaL_pushresult(m, v8)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L12
	}
L5:
	;
	v34 = v25
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if base.Ui32(v36) < base.Ui32(v8+int32(1036)) {
		v42 = v36
		v43 = v34
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L4
L8:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v43))))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v42 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v45)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)+1036))
	v52 = v50 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+1036)) = v52
	if v50 != 0 {
		v34 = v52
		goto L6
	} else {
		goto L11
	}
L9:
	;
	v38 = F_luaL_prepbuffer(m, v8)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+1036))
	v42 = v40
	v43 = v41
	goto L8
L11:
	;
	goto L7
L12:
	;
	m.G0 = v8 + int32(1040)
	return int32(1)
}
func F_str_upper(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
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
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(1040)
	m.G0 = v9
	v14 = F_luaL_checklstring(m, l0, int32(1), v9+int32(1036))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(12)
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1036))
	if v24 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_luaL_pushresult(m, v9)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L16
	}
L5:
	;
	v34 = int32(0)
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if base.Ui32(v36) < base.Ui32(v9+int32(1036)) {
		v41 = v36
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L4
L8:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v34))))
	if base.Ui32(v43+int32(-97)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v38 = F_luaL_prepbuffer(m, v9)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v41 = v40
	goto L8
L11:
	;
	v51 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v41 + v51
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v50)
	v56 = v34 + v51
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1036))
	if base.Ui32(v56) < base.Ui32(v57) {
		v34 = v56
		goto L6
	} else {
		goto L15
	}
L12:
	;
	v50 = v43 & int32(95)
	goto L14
L13:
	;
	v50 = v43
	goto L14
L14:
	;
	goto L11
L15:
	;
	goto L7
L16:
	;
	m.G0 = v9 + int32(1040)
	return int32(1)
}
