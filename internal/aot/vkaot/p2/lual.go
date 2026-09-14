package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaL_argerror(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
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
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	v6 = m.G0
	v8 = v6 - int32(144)
	m.G0 = v8
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	goto L5
L1:
	;
	m.G0 = v8 + int32(144)
	return v143
L2:
	;
	v73 = m.G3
	v78 = F_lua_getinfo(m, l0, v73+int32(_a2061), v8+int32(44))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L17
	} else {
		goto L19
	}
L3:
	;
	if v63 != 0 {
		goto L2
	} else {
		goto L16
	}
L4:
	;
	goto L3
L5:
	;
	goto L14
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(44))+96)) = v52
	v63 = int32(1)
	goto L4
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v16) <= base.Ui32(v48) {
		v63 = int32(0)
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v52 = base.I32_div_s(v16-v48, int32(24))
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
	v66 = m.G3
	v69 = F_luaL_error(m, l0, v66+int32(_a2062), v8)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	v143 = v69
	goto L1
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[1017]))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v86 == int32(0) {
		v109 = v85
		v110 = v86
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
	if v127 != 0 {
		v132 = v127
		goto L32
	} else {
		goto L33
	}
L21:
	;
	if v110-v109&int32(255) != 0 {
		v126 = l1
		goto L20
	} else {
		goto L29
	}
L22:
	;
	goto L21
L23:
	;
	if v86 != v85&int32(255) {
		v109 = v85
		v110 = v86
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v92 = v80
	v93 = v73 + int32(_a2063)
	goto L25
L25:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	if v97 == int32(0) {
		v109 = v96
		v110 = v97
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v109 = v96
	v110 = v97
	goto L22
L27:
	;
	v100 = int32(1)
	if v97 == v96&int32(255) {
		v92 = v92 + v100
		v93 = v93 + v100
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v115 = l1 + int32(-1)
	if v115 != 0 {
		v126 = v115
		goto L20
	} else {
		goto L30
	}
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l2
	v119 = m.G3
	v124 = F_luaL_error(m, l0, v119+int32(_a2064), v8+int32(16))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v143 = v124
	goto L1
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v126
	v136 = m.G3
	v141 = F_luaL_error(m, l0, v136+int32(_a2065), v8+int32(32))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L17
	} else {
		goto L34
	}
L33:
	;
	v128 = m.G3
	v130 = v128 + int32(_a357)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v130
	v132 = v130
	goto L32
L34:
	;
	v143 = v141
	goto L1
}
func F_luaL_checklstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_lua_tolstring(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			m.G0 = v8 + int32(16)
			return v10
		} else {
			v21 = m.G399
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(16))))
			if l1 < int32(1) {
				if l1 < int32(-9999) {
					switch l1 + int32(10002) {
					case 0:
						v76 = l0 + int32(72)
						v77 = m.G398
						if v76 != v77 {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
							v83 = v80
						} else {
							v83 = int32(-1)
						}
					case 1:
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v49
						v76 = l0 + int32(88)
						v77 = m.G398
						if v76 != v77 {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
							v83 = v80
						} else {
							v83 = int32(-1)
						}
					case 2:
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v76 = v72 + int32(96)
						v77 = m.G398
						if v76 != v77 {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
							v83 = v80
						} else {
							v83 = int32(-1)
						}
					default:
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
						v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+7)))
						if base.Ui32(v62) < base.Ui32(int32(-10002)-l1) {
							v83 = int32(-1)
						} else {
							v76 = v61 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
							v77 = m.G398
							if v76 != v77 {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
								v83 = v80
							} else {
								v83 = int32(-1)
							}
						}
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v76 = v40 + l1<<(uint(int32(4))%32)
					v77 = m.G398
					if v76 != v77 {
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
						v83 = v80
					} else {
						v83 = int32(-1)
					}
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v34 = v29 + l1<<(uint(int32(4))%32) + int32(-16)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v34) < base.Ui32(v35) {
					v76 = v34
					v77 = m.G398
					if v76 != v77 {
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
						v83 = v80
					} else {
						v83 = int32(-1)
					}
				} else {
					v83 = int32(-1)
				}
			}
			v85 = m.G3
			if v83 != int32(-1) {
				v90 = m.G399
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v90+v83<<(uint(int32(2))%32))))
				v95 = v94
			} else {
				v95 = v85 + int32(_a2018)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v95
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v25
			v98 = m.G3
			v101 = F_lua_pushfstring(m, l0, v98+int32(_a2066), v8)
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				v103 = F_luaL_argerror(m, l0, l1, v101)
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v10
				}
			}
		}
	}
}
func F_luaL_openlib(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v234 int32
	_ = v234
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var __phi278 int32
	_ = __phi278
	var v279 int32
	_ = v279
	var __phi279 int32
	_ = __phi279
	var v281 int64
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int64
	_ = v361
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int64
	_ = v373
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int64
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v529 int32
	_ = v529
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v385 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L2:
	;
	v15 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v16 == v15 {
		v37 = v15
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v41 = m.G3
	v45 = F_luaL_findtable(m, l0, int32(-10000), v41+int32(_a2067), int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v24 = v15
	v25 = l2
	goto L5
L5:
	;
	v28 = v24 + int32(1)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v29 != 0 {
		v24 = v28
		v25 = v25 + int32(8)
		goto L5
	} else {
		goto L7
	}
L6:
	;
	v37 = v28
	goto L3
L7:
	;
	goto L6
L8:
	;
	return
L9:
	;
	F_lua_getfield(m, l0, int32(-1), l1)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	goto L14
L11:
	;
	goto L59
L12:
	;
	if v107 == int32(5) {
		goto L11
	} else {
		goto L27
	}
L13:
	;
	v101 = m.G398
	if v67 != v101 {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	goto L18
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v67 = v64 + int32(-16)
	goto L13
L25:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	v107 = v104
	goto L12
L26:
	;
	v107 = int32(-1)
	goto L12
L27:
	;
	goto L30
L28:
	;
	v139 = F_luaL_findtable(m, l0, int32(-10002), l1, v37)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L8
	} else {
		goto L37
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v128 + int32(-16)
	goto L28
L30:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L29
L36:
	;
	goto L42
L37:
	;
	if v139 == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
	v144 = m.G3
	v147 = F_luaL_error(m, l0, v144+int32(_a2068), v11)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	F_lua_setfield(m, l0, int32(-3), l1)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L8
	} else {
		goto L56
	}
L41:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v169)))
	*(*int64)(unsafe.Add(mBase, uint32(v205))) = v206
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v205)+8)) = v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v210 + int32(16)
	goto L40
L42:
	;
	goto L48
L48:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v169 = v166 + int32(-16)
	goto L41
L56:
	;
	goto L11
L57:
	;
	v297 = l3 ^ int32(-1)
	if v297 < int32(1) {
		goto L80
	} else {
		goto L81
	}
L58:
	;
	v274 = v234 + int32(-16)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v275) <= base.Ui32(v274) {
		v292 = v275
		goto L73
	} else {
		goto L74
	}
L59:
	;
	goto L65
L65:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L58
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v292 + int32(-16)
	goto L57
L74:
	;
	__phi278 = v234 + int32(-32)
	__phi279 = v274
	v278 = __phi278
	v279 = __phi279
	goto L75
L75:
	;
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v278)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v278))) = v281
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v278)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v278)+8)) = v283
	v286 = v279 + int32(16)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v286) < base.Ui32(v287) {
		__phi278 = v279
		__phi279 = v286
		v278 = __phi278
		v279 = __phi279
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v292 = v287
	goto L73
L77:
	;
	goto L76
L78:
	;
	goto L1
L79:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v353) <= base.Ui32(v352) {
		v370 = v353
		goto L94
	} else {
		goto L95
	}
L80:
	;
	if v297 < int32(-9999) {
		goto L85
	} else {
		goto L86
	}
L81:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v308 = v303 + v297<<(uint(int32(4))%32) + int32(-16)
	v309 = m.G398
	if base.Ui32(v308) < base.Ui32(v302) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v311 = v308
	goto L84
L83:
	;
	v311 = v309
	goto L84
L84:
	;
	v352 = v311
	goto L79
L85:
	;
	switch v297 + int32(10002) {
	case 0:
		goto L88
	case 1:
		goto L89
	case 2:
		goto L90
	default:
		goto L87
	}
L86:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v352 = v314 + v297<<(uint(int32(4))%32)
	goto L79
L87:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+7)))
	v338 = m.G398
	if base.Ui32(v337) < base.Ui32(int32(-10002)-v297) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v352 = l0 + int32(72)
	goto L79
L89:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v326
	v352 = l0 + int32(88)
	goto L79
L90:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v352 = v320 + int32(96)
	goto L79
L91:
	;
	v349 = v338
	goto L93
L92:
	;
	v349 = v336 + (int32(-10003)-v297)<<(uint(int32(4))%32) + int32(24)
	goto L93
L93:
	;
	v352 = v349
	goto L79
L94:
	;
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v370)))
	*(*int64)(unsafe.Add(mBase, uint32(v352))) = v373
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v370)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v352)+8)) = v375
	goto L78
L95:
	;
	v356 = v353
	goto L96
L96:
	;
	v360 = v356 + int32(-16)
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v360)))
	*(*int64)(unsafe.Add(mBase, uint32(v356))) = v361
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v356+int32(-8))))
	*(*int32)(unsafe.Add(mBase, uint32(v356)+8)) = v365
	if base.Ui32(v352) < base.Ui32(v360) {
		v356 = v360
		goto L96
	} else {
		goto L98
	}
L97:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v370 = v368
	goto L94
L98:
	;
	goto L97
L99:
	;
	v504 = l3 ^ int32(-1)
	if v504 < int32(0) {
		goto L129
	} else {
		goto L130
	}
L100:
	;
	v391 = int32(0) - l3
	v396 = l2
	goto L101
L101:
	;
	if l3 < int32(1) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L99
L103:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	F_lua_pushcclosure(m, l0, v486, l3)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L8
	} else {
		goto L124
	}
L104:
	;
	v409 = int32(0)
	goto L105
L105:
	;
	if v391 < int32(1) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	goto L103
L107:
	;
	v476 = v409 + int32(1)
	if v476 != l3 {
		v409 = v476
		goto L105
	} else {
		goto L123
	}
L108:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v467 = *(*int64)(unsafe.Add(mBase, uint32(v463)))
	*(*int64)(unsafe.Add(mBase, uint32(v466))) = v467
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v463)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v466)+8)) = v469
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v471 + int32(16)
	goto L107
L109:
	;
	if v391 < int32(-9999) {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v421 = v416 + v391<<(uint(int32(4))%32) + int32(-16)
	v422 = m.G398
	if base.Ui32(v421) < base.Ui32(v415) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v424 = v421
	goto L113
L112:
	;
	v424 = v422
	goto L113
L113:
	;
	v463 = v424
	goto L108
L114:
	;
	switch v391 + int32(10002) {
	case 0:
		goto L117
	case 1:
		goto L118
	case 2:
		goto L119
	default:
		goto L116
	}
L115:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v463 = v427 + v391<<(uint(int32(4))%32)
	goto L108
L116:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)+4))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449)+7)))
	v451 = m.G398
	if base.Ui32(v450) < base.Ui32(int32(-10002)-v391) {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v463 = l0 + int32(72)
	goto L108
L118:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+4))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v439
	v463 = l0 + int32(88)
	goto L108
L119:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v463 = v433 + int32(96)
	goto L108
L120:
	;
	v462 = v451
	goto L122
L121:
	;
	v462 = v449 + (int32(-10003)-v391)<<(uint(int32(4))%32) + int32(24)
	goto L122
L122:
	;
	v463 = v462
	goto L108
L123:
	;
	goto L106
L124:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	F_lua_setfield(m, l0, int32(-2)-l3, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L8
	} else {
		goto L125
	}
L125:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v396)+8))
	if v492 != 0 {
		v396 = v396 + int32(8)
		goto L101
	} else {
		goto L126
	}
L126:
	;
	goto L102
L127:
	;
	m.G0 = v11 + int32(16)
	return
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v529
	goto L127
L129:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v529 = v522 + v504<<(uint(int32(4))%32) + int32(16)
	goto L128
L130:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v512 = v509 + v504<<(uint(int32(4))%32)
	if base.Ui32(v512) <= base.Ui32(v508) {
		v529 = v512
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v516 = v508
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v516)+8)) = int32(0)
	v520 = v516 + int32(16)
	if base.Ui32(v520) < base.Ui32(v512) {
		v516 = v520
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v529 = v512
	goto L128
}
func F_luaL_ref(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v93 int32
	_ = v93
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v195 int32
	_ = v195
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int64
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	if base.Ui32(l1+int32(-1)) < base.Ui32(int32(-10000)) {
		v16 = l1
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v16 = l1 + (v8-v9)>>(uint(int32(4))%32) + int32(1)
	}
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = v31 + int32(-16)
	v68 = m.G398
	if v34 != v68 {
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
		v74 = v71
	} else {
		v74 = int32(-1)
	}
	if v74 != 0 {
		if v16 < int32(1) {
			if v16 < int32(-9999) {
				switch v16 + int32(10002) {
				case 0:
					v158 = l0 + int32(72)
				case 1:
					v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
					v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v134
					v158 = l0 + int32(88)
				case 2:
					v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v158 = v128 + int32(96)
				default:
					v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
					v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
					v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+7)))
					v146 = m.G398
					if base.Ui32(v145) < base.Ui32(int32(-10002)-v16) {
						v157 = v146
					} else {
						v157 = v144 + (int32(-10003)-v16)<<(uint(int32(4))%32) + int32(24)
					}
					v158 = v157
				}
			} else {
				v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v158 = v122 + v16<<(uint(int32(4))%32)
			}
		} else {
			v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v116 = v111 + v16<<(uint(int32(4))%32) + int32(-16)
			v117 = m.G398
			if base.Ui32(v116) < base.Ui32(v110) {
				v119 = v116
			} else {
				v119 = v117
			}
			v158 = v119
		}
		v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
		v162 = F_luaH_getnum(m, v161, int32(0))
		mBase = m.M
		v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v164 = *(*int64)(unsafe.Add(mBase, uint32(v162)))
		*(*int64)(unsafe.Add(mBase, uint32(v163))) = v164
		v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v166
		v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v168 + int32(16)
		v173 = F_lua_tointeger(m, l0, int32(-1))
		mBase = m.M
		v176 = m.ExcPending
		if v176 != 0 {
			return int32(0)
		} else {
			v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v195 + int32(-16)
			if v173 == int32(0) {
				v279 = F_lua_objlen(m, l0, v16)
				mBase = m.M
				v280 = m.ExcPending
				if v280 != 0 {
					return int32(0)
				} else {
					v282 = v279 + int32(1)
					F_lua_rawseti(m, l0, v16, v282)
					mBase = m.M
					v284 = m.ExcPending
					if v284 != 0 {
						return int32(0)
					} else {
						return v282
					}
				}
			} else {
				if v16 < int32(1) {
					if v16 < int32(-9999) {
						switch v16 + int32(10002) {
						case 0:
							v259 = l0 + int32(72)
						case 1:
							v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
							v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
							v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v235
							v259 = l0 + int32(88)
						case 2:
							v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v259 = v229 + int32(96)
						default:
							v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
							v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
							v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+7)))
							v247 = m.G398
							if base.Ui32(v246) < base.Ui32(int32(-10002)-v16) {
								v258 = v247
							} else {
								v258 = v245 + (int32(-10003)-v16)<<(uint(int32(4))%32) + int32(24)
							}
							v259 = v258
						}
					} else {
						v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v259 = v223 + v16<<(uint(int32(4))%32)
					}
				} else {
					v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v217 = v212 + v16<<(uint(int32(4))%32) + int32(-16)
					v218 = m.G398
					if base.Ui32(v217) < base.Ui32(v211) {
						v220 = v217
					} else {
						v220 = v218
					}
					v259 = v220
				}
				v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
				v263 = F_luaH_getnum(m, v262, v173)
				mBase = m.M
				v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v265 = *(*int64)(unsafe.Add(mBase, uint32(v263)))
				*(*int64)(unsafe.Add(mBase, uint32(v264))) = v265
				v267 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v264)+8)) = v267
				v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v269 + int32(16)
				F_lua_rawseti(m, l0, v16, int32(0))
				mBase = m.M
				v275 = m.ExcPending
				if v275 != 0 {
					return int32(0)
				} else {
					F_lua_rawseti(m, l0, v16, v173)
					mBase = m.M
					v277 = m.ExcPending
					if v277 != 0 {
						return int32(0)
					} else {
						return v173
					}
				}
			}
		}
	} else {
		v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v93 + int32(-16)
		return int32(-1)
	}
}
func F_luaL_typerror(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v58 = l0 + int32(72)
				v59 = m.G398
				if v58 != v59 {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
					v65 = v62
				} else {
					v65 = int32(-1)
				}
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v58 = l0 + int32(88)
				v59 = m.G398
				if v58 != v59 {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
					v65 = v62
				} else {
					v65 = int32(-1)
				}
			case 2:
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v58 = v54 + int32(96)
				v59 = m.G398
				if v58 != v59 {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
					v65 = v62
				} else {
					v65 = int32(-1)
				}
			default:
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+7)))
				if base.Ui32(v44) < base.Ui32(int32(-10002)-l1) {
					v65 = int32(-1)
				} else {
					v58 = v43 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
					v59 = m.G398
					if v58 != v59 {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
						v65 = v62
					} else {
						v65 = int32(-1)
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v58 = v22 + l1<<(uint(int32(4))%32)
			v59 = m.G398
			if v58 != v59 {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
				v65 = v62
			} else {
				v65 = int32(-1)
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v16 = v11 + l1<<(uint(int32(4))%32) + int32(-16)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v16) < base.Ui32(v17) {
			v58 = v16
			v59 = m.G398
			if v58 != v59 {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
				v65 = v62
			} else {
				v65 = int32(-1)
			}
		} else {
			v65 = int32(-1)
		}
	}
	v67 = m.G3
	if v65 != int32(-1) {
		v72 = m.G399
		v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v65<<(uint(int32(2))%32))))
		v77 = v76
	} else {
		v77 = v67 + int32(_a2018)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
	v80 = m.G3
	v83 = F_lua_pushfstring(m, l0, v80+int32(_a2066), v7)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		return int32(0)
	} else {
		v87 = F_luaL_argerror(m, l0, l1, v83)
		mBase = m.M
		v88 = m.ExcPending
		if v88 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v87
		}
	}
}
func F_luaL_unref(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	if l2 < int32(0) {
		return
	} else {
		if base.Ui32(l1+int32(-1)) < base.Ui32(int32(-10000)) {
			v18 = l1
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = l1 + (v10-v11)>>(uint(int32(4))%32) + int32(1)
		}
		if v18 < int32(1) {
			if v18 < int32(-9999) {
				switch v18 + int32(10002) {
				case 0:
					v72 = l0 + int32(72)
				case 1:
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v48
					v72 = l0 + int32(88)
				case 2:
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v72 = v42 + int32(96)
				default:
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+7)))
					v60 = m.G398
					if base.Ui32(v59) < base.Ui32(int32(-10002)-v18) {
						v71 = v60
					} else {
						v71 = v58 + (int32(-10003)-v18)<<(uint(int32(4))%32) + int32(24)
					}
					v72 = v71
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v72 = v36 + v18<<(uint(int32(4))%32)
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v30 = v25 + v18<<(uint(int32(4))%32) + int32(-16)
			v31 = m.G398
			if base.Ui32(v30) < base.Ui32(v24) {
				v33 = v30
			} else {
				v33 = v31
			}
			v72 = v33
		}
		v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
		v76 = F_luaH_getnum(m, v75, int32(0))
		mBase = m.M
		v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v78 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
		*(*int64)(unsafe.Add(mBase, uint32(v77))) = v78
		v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = v80
		v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82 + int32(16)
		F_lua_rawseti(m, l0, v18, l2)
		mBase = m.M
		v87 = m.ExcPending
		if v87 != 0 {
			return
		} else {
			v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v89))) = base.F64_convert_i32_s(l2)
			v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v94 + int32(16)
			F_lua_rawseti(m, l0, v18, int32(0))
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return
			} else {
				return
			}
		}
	}
}
