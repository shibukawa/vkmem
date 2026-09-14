package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaO_log2(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	v5 = int32(-1)
	if base.Ui32(int32(256)) <= base.Ui32(l0) {
		v8 = l0
		v9 = v5
		for {
			v12 = int32(8)
			v13 = v9 + v12
			v17 = int32(base.Ui32(v8) >> (uint(v12) % 32))
			if base.Ui32(int32(65535)) < base.Ui32(v8) {
				v8 = v17
				v9 = v13
				continue
			} else {
				break
			}
			break
		}
		v19 = v13
		v20 = v17
	} else {
		v19 = v5
		v20 = l0
	}
	v22 = m.G3
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(_a2665)+v20))))
	return v19 + v26
}
func F_luaO_pushfstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v10 = F_luaO_pushvfstring(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v10
	}
}
func F_luaO_pushvfstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v264 float64
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
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
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = m.G3
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = F_luaS_newlstr(m, l0, v14+int32(_a320), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(16) < v26-v27 {
		v35 = v27
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v37 = v35 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v37
	v39 = int32(37)
	v40 = F___strchrnul(m, l1, v39)
	mBase = m.M
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v42 == v39 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_luaD_growstack(m, l0, int32(1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = v34
	goto L3
L6:
	;
	if v493&int32(3) == int32(0) {
		v522 = v493
		goto L130
	} else {
		goto L131
	}
L7:
	;
	v50 = l1
	v51 = l2
	v53 = v46
	v54 = v37
	v55 = int32(1)
	goto L13
L8:
	;
	if v46 != 0 {
		goto L7
	} else {
		goto L12
	}
L9:
	;
	v46 = v40
	goto L11
L10:
	;
	v46 = int32(0)
	goto L11
L11:
	;
	goto L8
L12:
	;
	v493 = l1
	v497 = v37
	v498 = int32(1)
	goto L6
L13:
	;
	v59 = F_luaS_newlstr(m, l0, v50, v53-v50)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v493 = v483
	v497 = v476
	v498 = v481
	goto L6
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v59
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(16) < v64-v65 {
		v73 = v65
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73 + int32(16)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	switch v77 + int32(-99) {
	case 0:
		goto L26
	case 1:
		goto L25
	case 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15:
		goto L21
	case 3:
		goto L24
	case 13:
		goto L23
	case 16:
		goto L27
	default:
		goto L22
	}
L17:
	;
	F_luaD_growstack(m, l0, int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v73 = v72
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v476
	v480 = int32(2)
	v481 = v55 + v480
	v483 = v53 + v480
	v484 = int32(37)
	v485 = F___strchrnul(m, v483, v484)
	mBase = m.M
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	if v487 == v484 {
		goto L124
	} else {
		goto L125
	}
L20:
	;
	v453 = m.G3
	v457 = F_luaS_newlstr(m, l0, v453+int32(_a2666), int32(1))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L119
	}
L21:
	;
	v372 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+18)) = uint8(v372)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+17)) = uint8(v77)
	v375 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v375)
	v378 = v12 + int32(16)
	if v378&int32(3) == v372 {
		v402 = v378
		goto L101
	} else {
		goto L102
	}
L22:
	;
	if v77 == int32(37) {
		goto L20
	} else {
		goto L98
	}
L23:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v282
	v286 = m.G3
	v289 = F_siprintf(m, v12+int32(16), v286+int32(_a2667), v12)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L77
	}
L24:
	;
	v263 = (v51 + int32(7)) & int32(-8)
	v264 = *(*float64)(unsafe.Add(mBase, uint32(v263)))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v73)+16)) = v264
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(16) < v268-v269 {
		v277 = v269
		goto L74
	} else {
		goto L75
	}
L25:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v73)+16)) = base.F64_convert_i32_s(v241)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(16) < v246-v247 {
		v255 = v247
		goto L71
	} else {
		goto L72
	}
L26:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+17)) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v159)
	v164 = v12 + int32(16)
	if v164&int32(3) == v160 {
		v188 = v164
		goto L53
	} else {
		goto L54
	}
L27:
	;
	v80 = m.G3
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v81 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v84 = v81
	goto L30
L29:
	;
	v84 = v80 + int32(_a2668)
	goto L30
L30:
	;
	if v84&int32(3) == int32(0) {
		v106 = v84
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v140 = F_luaS_newlstr(m, l0, v84, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L47
	}
L32:
	;
	v139 = v131 - v84
	goto L31
L33:
	;
	v110 = v106
	goto L41
L34:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v92 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v95 = v84
	goto L37
L36:
	;
	v139 = v84 - v84
	goto L31
L37:
	;
	v99 = v95 + int32(1)
	if v99&int32(3) == int32(0) {
		v106 = v99
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v104 != 0 {
		v95 = v99
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v131 = v99
	goto L32
L41:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v119 = int32(-2139062144)
	if (int32(16843008)-v116|v116)&v119 == v119 {
		v110 = v110 + int32(4)
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v125 = v110
	goto L44
L43:
	;
	goto L42
L44:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if v129 != 0 {
		v125 = v125 + int32(1)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v131 = v125
	goto L32
L46:
	;
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = v140
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(16) < v145-v146 {
		v154 = v146
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v475 = v51 + int32(4)
	v476 = v154 + int32(16)
	goto L19
L49:
	;
	F_luaD_growstack(m, l0, int32(1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v154 = v153
	goto L48
L51:
	;
	v222 = F_luaS_newlstr(m, l0, v164, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L67
	}
L52:
	;
	v221 = v213 - v164
	goto L51
L53:
	;
	v192 = v188
	goto L61
L54:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v174 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v177 = v164
	goto L57
L56:
	;
	v221 = v164 - v164
	goto L51
L57:
	;
	v181 = v177 + int32(1)
	if v181&int32(3) == int32(0) {
		v188 = v181
		goto L53
	} else {
		goto L59
	}
L59:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v186 != 0 {
		v177 = v181
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v213 = v181
	goto L52
L61:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v201 = int32(-2139062144)
	if (int32(16843008)-v198|v198)&v201 == v201 {
		v192 = v192 + int32(4)
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v207 = v192
	goto L64
L63:
	;
	goto L62
L64:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v211 != 0 {
		v207 = v207 + int32(1)
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v213 = v207
	goto L52
L66:
	;
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = v222
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(16) < v227-v228 {
		v236 = v228
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v475 = v51 + int32(4)
	v476 = v236 + int32(16)
	goto L19
L69:
	;
	F_luaD_growstack(m, l0, int32(1))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v236 = v235
	goto L68
L71:
	;
	v475 = v51 + int32(4)
	v476 = v255 + int32(16)
	goto L19
L72:
	;
	F_luaD_growstack(m, l0, int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v255 = v254
	goto L71
L74:
	;
	v475 = v263 + int32(8)
	v476 = v277 + int32(16)
	goto L19
L75:
	;
	F_luaD_growstack(m, l0, int32(1))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v277 = v276
	goto L74
L77:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v293 = v12 + int32(16)
	if v293&int32(3) == int32(0) {
		v317 = v293
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v351 = F_luaS_newlstr(m, l0, v293, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L94
	}
L79:
	;
	v350 = v342 - v293
	goto L78
L80:
	;
	v321 = v317
	goto L88
L81:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	if v303 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v306 = v293
	goto L84
L83:
	;
	v350 = v293 - v293
	goto L78
L84:
	;
	v310 = v306 + int32(1)
	if v310&int32(3) == int32(0) {
		v317 = v310
		goto L80
	} else {
		goto L86
	}
L86:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	if v315 != 0 {
		v306 = v310
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v342 = v310
	goto L79
L88:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v330 = int32(-2139062144)
	if (int32(16843008)-v327|v327)&v330 == v330 {
		v321 = v321 + int32(4)
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v336 = v321
	goto L91
L90:
	;
	goto L89
L91:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
	if v340 != 0 {
		v336 = v336 + int32(1)
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v342 = v336
	goto L79
L93:
	;
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = v351
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(16) < v356-v357 {
		v365 = v357
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v475 = v51 + int32(4)
	v476 = v365 + int32(16)
	goto L19
L96:
	;
	F_luaD_growstack(m, l0, int32(1))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v365 = v364
	goto L95
L98:
	;
	goto L21
L99:
	;
	v436 = F_luaS_newlstr(m, l0, v378, v435)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L115
	}
L100:
	;
	v435 = v427 - v378
	goto L99
L101:
	;
	v406 = v402
	goto L109
L102:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	if v388 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v391 = v378
	goto L105
L104:
	;
	v435 = v378 - v378
	goto L99
L105:
	;
	v395 = v391 + int32(1)
	if v395&int32(3) == int32(0) {
		v402 = v395
		goto L101
	} else {
		goto L107
	}
L107:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	if v400 != 0 {
		v391 = v395
		goto L105
	} else {
		goto L108
	}
L108:
	;
	v427 = v395
	goto L100
L109:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	v415 = int32(-2139062144)
	if (int32(16843008)-v412|v412)&v415 == v415 {
		v406 = v406 + int32(4)
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v421 = v406
	goto L112
L111:
	;
	goto L110
L112:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421))))
	if v425 != 0 {
		v421 = v421 + int32(1)
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v427 = v421
	goto L100
L114:
	;
	goto L113
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = v436
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(16) < v441-v442 {
		v450 = v442
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v475 = v51
	v476 = v450 + int32(16)
	goto L19
L117:
	;
	F_luaD_growstack(m, l0, int32(1))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v450 = v449
	goto L116
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = v457
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(16) < v462-v463 {
		v471 = v463
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v475 = v51
	v476 = v471 + int32(16)
	goto L19
L121:
	;
	F_luaD_growstack(m, l0, int32(1))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v471 = v470
	goto L120
L123:
	;
	if v491 != 0 {
		v50 = v483
		v51 = v475
		v53 = v491
		v54 = v476
		v55 = v481
		goto L13
	} else {
		goto L127
	}
L124:
	;
	v491 = v485
	goto L126
L125:
	;
	v491 = int32(0)
	goto L126
L126:
	;
	goto L123
L127:
	;
	goto L14
L128:
	;
	v556 = F_luaS_newlstr(m, l0, v493, v555)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L144
	}
L129:
	;
	v555 = v547 - v493
	goto L128
L130:
	;
	v526 = v522
	goto L138
L131:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
	if v508 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v511 = v493
	goto L134
L133:
	;
	v555 = v493 - v493
	goto L128
L134:
	;
	v515 = v511 + int32(1)
	if v515&int32(3) == int32(0) {
		v522 = v515
		goto L130
	} else {
		goto L136
	}
L136:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
	if v520 != 0 {
		v511 = v515
		goto L134
	} else {
		goto L137
	}
L137:
	;
	v547 = v515
	goto L129
L138:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v526)))
	v535 = int32(-2139062144)
	if (int32(16843008)-v532|v532)&v535 == v535 {
		v526 = v526 + int32(4)
		goto L138
	} else {
		goto L140
	}
L139:
	;
	v541 = v526
	goto L141
L140:
	;
	goto L139
L141:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	if v545 != 0 {
		v541 = v541 + int32(1)
		goto L141
	} else {
		goto L143
	}
L142:
	;
	v547 = v541
	goto L129
L143:
	;
	goto L142
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v497))) = v556
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(16) < v561-v562 {
		v570 = v562
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v572 = v570 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v572
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaV_concat(m, l0, v498+int32(1), (v572-v576)>>(uint(int32(4))%32)+int32(-1))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	F_luaD_growstack(m, l0, int32(1))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v570 = v569
	goto L145
L148:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v587 = v584 - v498<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v587
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v587+int32(-16))))
	m.G0 = v12 + int32(48)
	return v591 + int32(16)
}
