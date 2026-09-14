package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_R_1(m *base.Module, l0 float64) float64 {
	return base.F64_div(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), float64(1)))
}
func F_R_2(m *base.Module, l0 float64) float64 {
	return base.F64_div(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), float64(1)))
}
func F_RedisRegisterConnectionTypeSocket(m *base.Module) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_connTypeRegister(m, int32(_a2323))
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F___randname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = F___clock_gettime(m, v2, v10)
	mBase = m.M
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v18 = *(*int32)(unsafe.Add(mBase, _consts[1009]))
	v19 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, _consts[1015]))
	*(*int32)(unsafe.Add(mBase, _consts[1015])) = v21 + int32(1)
	v32 = v2
	v33 = v21 + (v15 + v16 + v18*int32(65537))
	for {
		v40 = int32(1)
		v46 = v33&int32(15) | v33<<(uint(v40)%32)&int32(32) + int32(65)
		*(*uint8)(unsafe.Add(mBase, uint32(l0+v32))) = uint8(v46)
		v51 = v32 + v40
		if v51 != int32(6) {
			v32 = v51
			v33 = int32(base.Ui32(v33) >> (uint(int32(5)) % 32))
			continue
		} else {
			break
		}
		break
	}
	m.G0 = v10 + int32(16)
	return l0
}
func F_raise(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	v9 = F_sigismember(m, int32(9116692), l0)
	mBase = m.M
	if base.B2i32(v9 != int32(0)) == int32(0) {
		v43 = l0 * int32(140)
		v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[1046]))))
		if v45&int32(4) == int32(0) {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[1047])))
			switch v60 + int32(2) {
			case 0:
				m.G0 = v6 + int32(128)
				return int32(0)
			default:
				m.Env.X__call_sighandler(m, v60, l0)
				mBase = m.M
				m.G0 = v6 + int32(128)
				return int32(0)
			case 2:
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1048])))
				if v67 == int32(0) {
					m.G0 = v6 + int32(128)
					return int32(0)
				} else {
					m.T0[v67].(func(*base.Module, int32))(m, l0)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(128)
						return int32(0)
					}
				}
			}
		} else {
			v53 = F__emscripten_memset_bulkmem(m, v6, base.I32_extend8_s(int32(0)), int32(128))
			mBase = m.M
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[1047])))
			m.T0[v55].(func(*base.Module, int32, int32, int32))(m, l0, v53, int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(128)
				return int32(0)
			}
		}
	} else {
		v17 = l0 + int32(-1)
		if base.Ui32(int32(63)) < base.Ui32(v17) {
			v24 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(28)
		} else {
			if base.Ui32(int32(2)) < base.Ui32(l0+int32(-32)) {
				v31 = int32(base.Ui32(v17)>>(uint(int32(3))%32)) & int32(536870908)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+uint32(_consts[1049])))
				*(*int32)(unsafe.Add(mBase, uint32(v31)+uint32(_consts[1049]))) = v33 | int32(1)<<(uint(v17)%32)
			} else {
				v24 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(28)
			}
		}
		m.G0 = v6 + int32(128)
		return int32(0)
	}
}
func F_rand(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v8 int64
	_ = v8
	v2 = int32(0)
	v4 = *(*int64)(unsafe.Add(mBase, _consts[343]))
	v8 = v4*int64(6364136223846793005) + int64(1)
	*(*int64)(unsafe.Add(mBase, _consts[343])) = v8
	return base.I32_wrap_i64(int64(base.Ui64(v8) >> (uint(int64(33)) % 64)))
}
func F_random(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v6 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[303]))
	v9 = *(*int32)(unsafe.Add(mBase, _consts[304]))
	if v9 != 0 {
		v18 = int32(0)
		v19 = *(*int32)(unsafe.Add(mBase, _consts[305]))
		v20 = int32(2)
		v22 = v7 + v19<<(uint(v20)%32)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		v25 = *(*int32)(unsafe.Add(mBase, _consts[306]))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v7+v25<<(uint(v20)%32))))
		v30 = v23 + v29
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v30
		v35 = v25 + int32(1)
		if v35 == v9 {
			v37 = v18
		} else {
			v37 = v35
		}
		*(*int32)(unsafe.Add(mBase, _consts[306])) = v37
		v39 = int32(0)
		v42 = v19 + int32(1)
		if v42 == v9 {
			v44 = v39
		} else {
			v44 = v42
		}
		*(*int32)(unsafe.Add(mBase, _consts[305])) = v44
		v49 = int32(base.Ui32(v30) >> (uint(int32(1)) % 32))
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		v16 = (v10*int32(1103515245) + int32(12345)) & int32(2147483647)
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v16
		v49 = v16
	}
	return v49
}
func F_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	v16 = m.Wasi_snapshot_preview1.Fd_read(m, l0, v7+int32(8), int32(1), v7+int32(4))
	mBase = m.M
	if v16 != 0 {
		v18 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = v16
	} else {
	}
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	m.G0 = v7 + int32(16)
	if v16 != 0 {
		v27 = int32(-1)
	} else {
		v27 = v22
	}
	return v27
}
func F_read_numeral(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
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
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
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
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v573 int64
	_ = v573
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v622 int32
	_ = v622
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
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = v17
	goto L1
L1:
	;
	F_save(m, l0, v21)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v56 = v46 & int32(255)
	if v56 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	return
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v33 + int32(-1)
	if v33 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v46
	if v46 == int32(46) {
		v21 = v46
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v44 = F_luaZ_fill(m, v32)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v39 + int32(1)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v46 = v43
	goto L5
L8:
	;
	v46 = v44
	goto L5
L9:
	;
	if base.Ui32(v46+int32(-48)) < base.Ui32(int32(10)) {
		v21 = v46
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L2
L11:
	;
	goto L28
L12:
	;
	F_save(m, l0, v46)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L3
	} else {
		goto L16
	}
L13:
	;
	if v56 == int32(101) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if v56 != int32(69) {
		v112 = v46
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v66 + int32(-1)
	if v66 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v79
	v83 = v79 & int32(255)
	if base.Ui32(int32(63)) < base.Ui32(v83) {
		v112 = v79
		goto L11
	} else {
		goto L21
	}
L18:
	;
	v77 = F_luaZ_fill(m, v65)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L20
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v72 + int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	v79 = v76
	goto L17
L20:
	;
	v79 = v77
	goto L17
L21:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v83))%64)&int64(43980465111041) == int64(0) {
		v112 = v79
		goto L11
	} else {
		goto L22
	}
L22:
	;
	F_save(m, l0, v79)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v96 + int32(-1)
	if v96 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
	v112 = v109
	goto L11
L25:
	;
	v107 = F_luaZ_fill(m, v95)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L27
	}
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v102 + int32(1)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	v109 = v106
	goto L24
L27:
	;
	v109 = v107
	goto L24
L28:
	;
	if v112 == int32(95) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v185 = int32(0)
	F_save(m, l0, v185)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L3
	} else {
		goto L43
	}
L30:
	;
	v132 = v112
	goto L33
L31:
	;
	if base.B2i32(base.Ui32(v112+int32(-48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v112|int32(32)+int32(-97)) < base.Ui32(int32(26))) == int32(0) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	F_save(m, l0, v132)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L3
	} else {
		goto L35
	}
L34:
	;
	goto L29
L35:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v144 + int32(-1)
	if v144 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v157
	goto L40
L37:
	;
	v155 = F_luaZ_fill(m, v143)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L3
	} else {
		goto L39
	}
L38:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = v150 + int32(1)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v157 = v154
	goto L36
L39:
	;
	v157 = v155
	goto L36
L40:
	;
	if v157 == int32(95) {
		v132 = v157
		goto L33
	} else {
		goto L41
	}
L41:
	;
	if base.B2i32(base.Ui32(v157+int32(-48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v157|int32(32)+int32(-97)) < base.Ui32(int32(26))) != 0 {
		v132 = v157
		goto L33
	} else {
		goto L42
	}
L42:
	;
	goto L34
L43:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	if v191 == int32(0) {
		v295 = v190
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v301 = F_luaO_str2d(m, v295, l1)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L3
	} else {
		goto L67
	}
L45:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	v196 = v191 & int32(3)
	if v196 == int32(0) {
		v224 = v191
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if base.Ui32(v191) < base.Ui32(int32(4)) {
		goto L53
	} else {
		goto L54
	}
L47:
	;
	v202 = v191
	v204 = v185
	goto L48
L48:
	;
	v212 = v202 + int32(-1)
	v213 = v190 + v212
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	if v214 != int32(46) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v224 = v212
	goto L46
L50:
	;
	v219 = v204 + int32(1)
	if v219 != v196 {
		v202 = v212
		v204 = v219
		goto L48
	} else {
		goto L52
	}
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v194)
	goto L50
L52:
	;
	goto L49
L53:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	v295 = v288
	goto L44
L54:
	;
	v244 = v224
	goto L55
L55:
	;
	v253 = v190 + int32(-1) + v244
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	if v254 != int32(46) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L53
L57:
	;
	v258 = v190 + int32(-2) + v244
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	if v259 != int32(46) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v253))) = uint8(v194)
	goto L57
L59:
	;
	v263 = v190 + int32(-3) + v244
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	if v264 != int32(46) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v258))) = uint8(v194)
	goto L59
L61:
	;
	v269 = v244 + int32(-4)
	v270 = v190 + v269
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if v271 != int32(46) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v263))) = uint8(v194)
	goto L61
L63:
	;
	if v269 != 0 {
		v244 = v269
		goto L55
	} else {
		goto L65
	}
L64:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v270))) = uint8(v194)
	goto L63
L65:
	;
	goto L56
L66:
	;
	m.G0 = v15 + int32(112)
	return
L67:
	;
	if v301 != 0 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	goto L71
L69:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)) = uint8(v306)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	if v312 == int32(0) {
		v414 = v311
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _consts[992]))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	goto L69
L71:
	;
	goto L70
L73:
	;
	v421 = F_luaO_str2d(m, v414, l1)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L3
	} else {
		goto L95
	}
L74:
	;
	v316 = v312 & int32(3)
	if v316 == int32(0) {
		v346 = v312
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if base.Ui32(v312) < base.Ui32(int32(4)) {
		goto L82
	} else {
		goto L83
	}
L76:
	;
	v325 = v312
	v326 = int32(0)
	goto L77
L77:
	;
	v335 = v325 + int32(-1)
	v336 = v311 + v335
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
	if v337 != v308&int32(255) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v346 = v335
	goto L75
L79:
	;
	v341 = v326 + int32(1)
	if v341 != v316 {
		v325 = v335
		v326 = v341
		goto L77
	} else {
		goto L81
	}
L80:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v336))) = uint8(v306)
	goto L79
L81:
	;
	goto L78
L82:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v414 = v408
	goto L73
L83:
	;
	v364 = v308 & int32(255)
	v368 = v346
	goto L84
L84:
	;
	v377 = v311 + int32(-1) + v368
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377))))
	if v378 != v364 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L82
L86:
	;
	v381 = v311 + int32(-2) + v368
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	if v382 != v364 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v377))) = uint8(v306)
	goto L86
L88:
	;
	v385 = v311 + int32(-3) + v368
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v386 != v364 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v381))) = uint8(v306)
	goto L88
L90:
	;
	v390 = v368 + int32(-4)
	v391 = v311 + v390
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391))))
	if v392 != v364 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v385))) = uint8(v306)
	goto L90
L92:
	;
	if v390 != 0 {
		v368 = v390
		goto L84
	} else {
		goto L94
	}
L93:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v391))) = uint8(v306)
	goto L92
L94:
	;
	goto L85
L95:
	;
	if v421 != 0 {
		goto L66
	} else {
		goto L96
	}
L96:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	if v424 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v527 = v15 + int32(32)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v530 = v528 + int32(16)
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530))))
	switch v534 + int32(-61) {
	case 0:
		goto L121
	default:
		goto L119
	case 3:
		goto L120
	}
L98:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	v430 = v424 & int32(3)
	if v430 == int32(0) {
		v461 = v424
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if base.Ui32(v424) < base.Ui32(int32(4)) {
		goto L97
	} else {
		goto L106
	}
L100:
	;
	v439 = v424
	v440 = int32(0)
	goto L101
L101:
	;
	v449 = v439 + int32(-1)
	v450 = v428 + v449
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450))))
	if v451 != v427&int32(255) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v461 = v449
	goto L99
L103:
	;
	v456 = v440 + int32(1)
	if v456 != v430 {
		v439 = v449
		v440 = v456
		goto L101
	} else {
		goto L105
	}
L104:
	;
	v453 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v450))) = uint8(v453)
	goto L103
L105:
	;
	goto L102
L106:
	;
	v479 = v427 & int32(255)
	v483 = v461
	goto L107
L107:
	;
	v492 = v428 + int32(-1) + v483
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492))))
	if v493 != v479 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L97
L109:
	;
	v497 = v428 + int32(-2) + v483
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	if v498 != v479 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v495 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v492))) = uint8(v495)
	goto L109
L111:
	;
	v502 = v428 + int32(-3) + v483
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502))))
	if v503 != v479 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v500 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v497))) = uint8(v500)
	goto L111
L113:
	;
	v508 = v483 + int32(-4)
	v509 = v428 + v508
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509))))
	if v510 != v479 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v505 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v502))) = uint8(v505)
	goto L113
L115:
	;
	if v508 != 0 {
		v483 = v508
		goto L107
	} else {
		goto L117
	}
L116:
	;
	v512 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v509))) = uint8(v512)
	goto L115
L117:
	;
	goto L108
L118:
	;
	v608 = m.G3
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v608 + int32(_a2672)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v610
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v15 + int32(32)
	v622 = F_luaO_pushfstring(m, v609, v608+int32(_a2654), v15+int32(16))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L3
	} else {
		goto L130
	}
L119:
	;
	v561 = m.G3
	v564 = F_strcspn(m, v530, v561+int32(_a2653))
	mBase = m.M
	v571 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v561)+uint32(_consts[988]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v15+int32(40)))) = uint16(v571)
	v573 = *(*int64)(unsafe.Add(mBase, uint32(v561)+uint32(_consts[989])))
	*(*int64)(unsafe.Add(mBase, uint32(v527))) = v573
	v576 = int32(63)
	if base.Ui32(v564) < base.Ui32(v576) {
		goto L126
	} else {
		goto L127
	}
L120:
	;
	v546 = v528 + int32(17)
	v547 = F_strlen(m, v546)
	mBase = m.M
	v548 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v527))) = uint8(v548)
	v551 = int32(72)
	if base.Ui32(v547) <= base.Ui32(v551) {
		v559 = v546
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v539 = F_strncpy(m, v527, v528+int32(17), int32(80))
	mBase = m.M
	v543 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v539+int32(79)))) = uint8(v543)
	goto L118
L122:
	;
	v560 = F_strcat(m, v527, v559)
	mBase = m.M
	goto L118
L123:
	;
	v553 = F_strlen(m, v527)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v527+v553))) = int32(3026478)
	v559 = v546 + (v547 - v551)
	goto L122
L124:
	;
	v590 = F_strlen(m, v527)
	mBase = m.M
	v591 = v527 + v590
	v592 = m.G3
	v595 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v592)+uint32(_consts[990]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v591))) = uint16(v595)
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592)+uint32(_consts[991]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v591+int32(2)))) = uint8(v601)
	goto L118
L125:
	;
	v588 = F_strcat(m, v527, v530)
	mBase = m.M
	goto L124
L126:
	;
	v578 = v564
	goto L128
L127:
	;
	v578 = v576
	goto L128
L128:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530+v578))))
	if v580 == int32(0) {
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v583 = F_strncat(m, v527, v530, v578)
	mBase = m.M
	v584 = F_strlen(m, v583)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v583+v584))) = int32(3026478)
	goto L124
L130:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_save(m, l0, int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v622
	v634 = F_luaO_pushfstring(m, v624, v608+int32(_a2671), v15)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L3
	} else {
		goto L132
	}
L132:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_luaD_throw(m, v636, int32(3))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L3
	} else {
		goto L133
	}
L133:
	;
	goto L66
}
func F_readdir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v4 < v5 {
		v27 = v4
		v29 = l0 + v27
		v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(40)))))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v27 + v32
		v37 = *(*int64)(unsafe.Add(mBase, uint32(v29+int32(32))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v37
		v41 = v29 + int32(24)
		return v41
	} else {
		v7 = int32(0)
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = m.Env.X__syscall_getdents64(m, v8, l0+int32(24), int32(2048))
		mBase = m.M
		if v7 < v12 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v12
			v27 = int32(0)
			v29 = l0 + v27
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(40)))))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v27 + v32
			v37 = *(*int64)(unsafe.Add(mBase, uint32(v29+int32(32))))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v37
			v41 = v29 + int32(24)
			return v41
		} else {
			if v12 == int32(-44) {
				v41 = v7
				return v41
			} else {
				if v12 == int32(0) {
					v41 = v7
					return v41
				} else {
					v20 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[18])) = v20 - v12
					return v20
				}
			}
		}
	}
}
func F_readonlyCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v2 | int32(131072)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_receiveSynchronousResponse(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v46 int32
	_ = v46
	v4 = m.G0
	v6 = v4 - int32(272)
	m.G0 = v6
	v12 = *(*int32)(unsafe.Add(mBase, _consts[599]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	v18 = m.T0[v17].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v6+int32(16), int32(256), base.I64_extend_i32_s(v12*int32(1000)))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v18 != int32(-1) {
			v38 = int32(_a20)
			v40 = *(*int64)(unsafe.Add(mBase, _consts[109]))
			*(*int64)(unsafe.Add(mBase, _consts[603])) = v40
			v44 = F_sdsnew(m, v6+int32(16))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				v46 = v44
				m.G0 = v6 + int32(272)
				return v46
			}
		} else {
			v24 = int32(0)
			v26 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if int32(3) < v26 {
				v46 = v24
				m.G0 = v6 + int32(272)
				return v46
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
					F__serverLog(m, int32(3), int32(_a1959), v6)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v46 = v24
						m.G0 = v6 + int32(272)
						return v46
					}
				}
			}
		}
	}
}
func F_recv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v5 = int32(0)
	v7 = F_recvfrom(m, l0, l1, l2, l3, v5, v5)
	return v7
}
func F_redactClientCommandArgument(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	if int32(0) < l1 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
		v12 = int32(1)
		if base.Ui32(l1) < base.Ui32(int32(32)) {
			v17 = v12 << (uint(l1) % 32)
		} else {
			v17 = v12
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v11 | v17
		return
	} else {
		F__serverAssert(m, int32(_a1754), int32(_a1630), int32(6052))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_rejectCommand(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	F_flagTransaction(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v7
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		if v9 == v7 {
			F_addReplyErrorObject(m, l0, l1)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				if l2 == int32(0) {
					return
				} else {
					v26 = F_objectGetVal(m, l1)
					mBase = m.M
					F_moduleFireCommandRejectedEvent(m, l0, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v12 = *(*int64)(unsafe.Add(mBase, uint32(v9)+120))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+120)) = v12 + int64(1)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
			if v16 != int32(17) {
				F_addReplyErrorObject(m, l0, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					if l2 == int32(0) {
						return
					} else {
						v26 = F_objectGetVal(m, l1)
						mBase = m.M
						F_moduleFireCommandRejectedEvent(m, l0, v26)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v19 = F_objectGetVal(m, l1)
				mBase = m.M
				F_execCommandAbort(m, l0, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					if l2 == int32(0) {
						return
					} else {
						v26 = F_objectGetVal(m, l1)
						mBase = m.M
						F_moduleFireCommandRejectedEvent(m, l0, v26)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_rejectCommandFormat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int64
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
	v11 = F_sdsempty(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v14 = F_sdscatvprintf(m, v11, l2, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-1)))))
	switch v26 & int32(7) {
	case 0:
		goto L11
	case 1:
		goto L10
	case 2:
		goto L9
	case 3:
		goto L8
	case 4:
		goto L7
	default:
		goto L5
	}
L4:
	;
	F_flagTransaction(m, l0)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L23
	}
L5:
	;
	goto L4
L6:
	;
	if v43 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-17))))
	v43 = v42
	goto L6
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-9))))
	v43 = v39
	goto L6
L9:
	;
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(-5)))))
	v43 = v36
	goto L6
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-3)))))
	v43 = v33
	goto L6
L11:
	;
	v43 = int32(base.Ui32(v26) >> (uint(int32(3)) % 32))
	goto L6
L12:
	;
	v53 = int32(0)
	goto L13
L13:
	;
	goto L16
L14:
	;
	goto L5
L15:
	;
	v91 = v53 + int32(1)
	if v91 != v43 {
		v53 = v91
		goto L13
	} else {
		goto L22
	}
L16:
	;
	v59 = v14 + v53
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v67 = int32(0)
	goto L17
L17:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[486]))))
	if v60&int32(255) != v73 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L15
L19:
	;
	v79 = v67 + int32(1)
	if v79 != int32(2) {
		v67 = v79
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[487]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v76)
	goto L15
L21:
	;
	goto L18
L22:
	;
	goto L14
L23:
	;
	v104 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v106 == v104 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if l1 == int32(0) {
		v118 = v106
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v106)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v106)+120)) = v109 + int64(1)
	goto L24
L26:
	;
	if v118 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	F_moduleFireCommandRejectedEvent(m, l0, v14)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v118 = v117
	goto L26
L29:
	;
	m.G0 = v8 + int32(16)
	return
L30:
	;
	F_addReplyErrorSds(m, l0, v14)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L35
	}
L31:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+48))
	if v121 != int32(17) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	F_execCommandAbort(m, l0, v14)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_sdsfree(m, v14)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	goto L29
}
func F_releaseBlockedEntry(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v10 = F_dictFetchValue(m, v9, v6)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if v10 == int32(0) {
			F__serverAssertWithInfo(m, l0, v6, int32(_a187), int32(_a184), int32(589))
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
			F_listUnlinkNode(m, v10, v7)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
				if v16 != 0 {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
					if v22 == int32(0) {
						if l2 == int32(0) {
							return
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
							v47 = F_dictDelete(m, v46, v6)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
						v27 = F_dictFind(m, v26, v6)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							if v27 == int32(0) {
								F__serverAssertWithInfo(m, l0, v6, int32(_a188), int32(_a184), int32(604))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v32 = *(*int64)(unsafe.Add(mBase, uint32(v27)+8))
								v33 = v32 + int64(-1)
								*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = v33
								if v33 != int64(0) {
									if l2 == int32(0) {
										return
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
										v47 = F_dictDelete(m, v46, v6)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											return
										}
									}
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
									v40 = F_dictDelete(m, v39, v6)
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return
									} else {
										if l2 == int32(0) {
											return
										} else {
											v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
											v47 = F_dictDelete(m, v46, v6)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						}
					}
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
					v19 = F_dictDelete(m, v18, v6)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
						v40 = F_dictDelete(m, v39, v6)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							if l2 == int32(0) {
								return
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
								v47 = F_dictDelete(m, v46, v6)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
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
func F_releaseInfoSectionDict(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, _consts[791]))
	if l0 == v3 {
		return
	} else {
		F_dictRelease(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_rememberReplicaKeyWithExpire(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	v5 = *(*int32)(unsafe.Add(mBase, _consts[330]))
	if v5 != 0 {
		v11 = v5
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if int32(63) < v12 {
			return
		} else {
			v15 = F_objectGetVal(m, l1)
			mBase = m.M
			v16 = F_dictAddOrFind(m, v11, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v19 = F_objectGetVal(m, l1)
				mBase = m.M
				if v18 != v19 {
					v30 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
					v32 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v30 | int64(1)<<(uint(v32)%64)
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, _consts[330]))
					v23 = F_objectGetVal(m, l1)
					mBase = m.M
					v24 = F_sdsdup(m, v23)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						F_dictSetKey(m, v22, v16, v24)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
							v30 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
							v32 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v30 | int64(1)<<(uint(v32)%64)
							return
						}
					}
				}
			}
		}
	} else {
		v8 = F_dictCreate(m, int32(_a615))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[330])) = v8
			v11 = v8
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if int32(63) < v12 {
				return
			} else {
				v15 = F_objectGetVal(m, l1)
				mBase = m.M
				v16 = F_dictAddOrFind(m, v11, v15)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v19 = F_objectGetVal(m, l1)
					mBase = m.M
					if v18 != v19 {
						v30 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
						v32 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v30 | int64(1)<<(uint(v32)%64)
						return
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, _consts[330]))
						v23 = F_objectGetVal(m, l1)
						mBase = m.M
						v24 = F_sdsdup(m, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_dictSetKey(m, v22, v16, v24)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
								v30 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
								v32 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
								*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v30 | int64(1)<<(uint(v32)%64)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_removeFromBucket_HASHTABLE(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	switch l0 + int32(1) {
	case 0:
		F__serverAssert(m, int32(_a2523), int32(_a2500), int32(802))
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 1:
		F__serverAssert(m, int32(_a2513), int32(_a2500), int32(781))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	default:
		if l0&int32(7) != int32(4) {
			F__serverAssert(m, int32(_a2523), int32(_a2500), int32(802))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v25 = l0 & int32(-8)
			v26 = F_hashtableDelete(m, v25, l1)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v69 = l0
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v26)
					m.G0 = v8 + int32(64)
					return v69
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					if v30+v31 == int32(0) {
						F__serverAssert(m, int32(_a2524), int32(_a2500), int32(1372))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
						if v35+v36 != int32(1) {
							v69 = l0
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v26)
							m.G0 = v8 + int32(64)
							return v69
						} else {
							v41 = v8 + int32(16)
							v42 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v41)+14)) = uint8(v42)
							*(*int32)(unsafe.Add(mBase, uint32(v41))) = v25
							*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v42
							*(*uint8)(unsafe.Add(mBase, uint32(v41)+15)) = uint8(v42)
							*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(-1)
							if v25 == v42 {
							} else {
							}
							v64 = F_hashtableNext(m, v8+int32(16), v8+int32(12))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								F_hashtableRelease(m, v25)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
									v69 = v68
									*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v26)
									m.G0 = v8 + int32(64)
									return v69
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_removeSigSegvHandlers(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v19 int32
	_ = v19
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v154 int32
	_ = v154
	v2 = m.G0
	v4 = v2 - int32(144)
	m.G0 = v4
	*(*int64)(unsafe.Add(mBase, uint32(v4+int32(8)))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+136)) = int32(-1073741824)
	v19 = v4 + int32(4)
	if v19 == int32(0) {
	} else {
		v42 = F___memcpy(m, int32(9118724), v19, int32(140))
		mBase = m.M
	}
	v47 = v4 + int32(4)
	if v47 == int32(0) {
	} else {
		v70 = F___memcpy(m, int32(9118164), v47, int32(140))
		mBase = m.M
	}
	v75 = v4 + int32(4)
	if v75 == int32(0) {
	} else {
		v98 = F___memcpy(m, int32(9118304), v75, int32(140))
		mBase = m.M
	}
	v103 = v4 + int32(4)
	if v103 == int32(0) {
	} else {
		v126 = F___memcpy(m, int32(9117744), v103, int32(140))
		mBase = m.M
	}
	v131 = v4 + int32(4)
	if v131 == int32(0) {
	} else {
		v154 = F___memcpy(m, int32(9118024), v131, int32(140))
		mBase = m.M
	}
	m.G0 = v4 + int32(144)
	return
}
func F_resetErrorTableStats(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, _consts[770]))
	F_freeErrorsRadixTreeAsync(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v6 = F_raxNew(m)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[770])) = v6
			return
		}
	}
}
func F_resize_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v139 int32
	_ = v139
	var v173 int32
	_ = v173
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v205 float64
	_ = v205
	var v206 int32
	_ = v206
	var v212 float64
	_ = v212
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 float64
	_ = v231
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 float64
	_ = v251
	var v253 int32
	_ = v253
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v304 int32
	_ = v304
	var v307 int64
	_ = v307
	var v309 int32
	_ = v309
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v359 int32
	_ = v359
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 float64
	_ = v396
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int64
	_ = v414
	var v416 int32
	_ = v416
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if l2 <= v22 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_setnodevector(m, l0, l1, l3)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L20
	}
L2:
	;
	if base.Ui32(int32(268435455)) < base.Ui32(l2+int32(1)) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if l2 <= v39 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v35 = F_luaM_toobig(m, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L8
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v29 = int32(4)
	v33 = F_luaM_realloc_(m, l0, v28, v22<<(uint(v29)%32), l2<<(uint(v29)%32))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v37 = v33
	goto L3
L8:
	;
	v37 = v35
	goto L3
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = l2
	goto L1
L10:
	;
	v43 = (l2 - v39) & int32(7)
	if v43 == int32(0) {
		v82 = v39
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if base.Ui32(int32(-8)) < base.Ui32(v39-l2) {
		goto L9
	} else {
		goto L16
	}
L12:
	;
	v57 = v39
	v59 = int32(0)
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37+v57<<(uint(int32(4))%32))+8)) = int32(0)
	v67 = int32(1)
	v68 = v57 + v67
	v70 = v59 + v67
	if v70 != v43 {
		v57 = v68
		v59 = v70
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v82 = v68
	goto L11
L15:
	;
	goto L14
L16:
	;
	v100 = v82
	goto L17
L17:
	;
	v107 = v37 + v100<<(uint(int32(4))%32)
	v108 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v107+int32(24)))) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v107+int32(40)))) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v107+int32(56)))) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v107+int32(72)))) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v107+int32(88)))) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v107+int32(104)))) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v107+int32(120)))) = v108
	v139 = v100 + int32(8)
	if v139 != l2 {
		v100 = v139
		goto L17
	} else {
		goto L19
	}
L18:
	;
	goto L9
L19:
	;
	goto L18
L20:
	;
	if v22 <= l2 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v21 == int32(31) {
		goto L51
	} else {
		goto L52
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = l2
	v188 = l2
	goto L23
L23:
	;
	v192 = v188 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v196 = v193 + v188<<(uint(int32(4))%32)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	if v197 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if base.Ui32(int32(268435455)) < base.Ui32(l2+int32(1)) {
		goto L47
	} else {
		goto L48
	}
L25:
	;
	if v192 != v22 {
		v188 = v192
		goto L23
	} else {
		goto L45
	}
L26:
	;
	if v188 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
	*(*int64)(unsafe.Add(mBase, uint32(v304))) = v307
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v304)+8)) = v309
	goto L25
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v18))) = base.F64_convert_i32_s(v192)
	v290 = F_newkey(m, l0, l1, v18)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L6
	} else {
		goto L44
	}
L29:
	;
	v269 = m.G398
	if v266 != v269 {
		v304 = v266
		goto L27
	} else {
		goto L43
	}
L30:
	;
	v245 = v230
	goto L37
L31:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v214 = base.I64_reinterpret_f64(v212)
	v219 = int32(-1)
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v226 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v214)>>(uint(int64(32))%64))+v214), v219<<(uint(v220)%32)^v219|int32(1))
	v230 = v213 + v226<<(uint(int32(5))%32)
	v231 = v212
	goto L30
L32:
	;
	v266 = v193 + v192<<(uint(int32(4))%32) + int32(-16)
	goto L29
L33:
	;
	v205 = base.F64_convert_i32_s(v192)
	if v192 != 0 {
		v212 = v205
		goto L31
	} else {
		goto L36
	}
L34:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v188 < v202 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v212 = base.F64_convert_i32_u(v192)
	goto L31
L36:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v230 = v206
	v231 = v205
	goto L30
L37:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v245)+24))
	if v248 != int32(3) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v245)+28))
	if v253 != 0 {
		v245 = v253
		goto L37
	} else {
		goto L42
	}
L40:
	;
	v251 = *(*float64)(unsafe.Add(mBase, uint32(v245)+16))
	if base.F64_eq(v251, v231) != 0 {
		v266 = v245
		goto L29
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L28
L43:
	;
	goto L28
L44:
	;
	v304 = v290
	goto L27
L45:
	;
	goto L24
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v340
	goto L21
L47:
	;
	v338 = F_luaM_toobig(m, l0)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L6
	} else {
		goto L50
	}
L48:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v332 = int32(4)
	v336 = F_luaM_realloc_(m, l0, v331, v22<<(uint(v332)%32), l2<<(uint(v332)%32))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	v340 = v336
	goto L46
L50:
	;
	v340 = v338
	goto L46
L51:
	;
	v441 = m.G3
	if v20 == v441+int32(_a2685) {
		goto L66
	} else {
		goto L67
	}
L52:
	;
	v359 = int32(-1)
	v375 = v359<<(uint(v21)%32) ^ v359
	goto L53
L53:
	;
	v380 = v20 + v375<<(uint(int32(5))%32)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+8))
	if v381 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L51
L55:
	;
	if int32(0) < v375 {
		v375 = v375 + int32(-1)
		goto L53
	} else {
		goto L65
	}
L56:
	;
	v385 = v380 + int32(16)
	v386 = F_luaH_get(m, l1, v385)
	mBase = m.M
	v387 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)) = uint8(v387)
	v389 = m.G398
	if v386 != v389 {
		v412 = v386
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v380)))
	*(*int64)(unsafe.Add(mBase, uint32(v412))) = v414
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v380)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v412)+8)) = v416
	goto L55
L58:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v380)+24))
	v392 = m.G3
	switch v391 {
	case 0:
		v400 = v392 + int32(_a2686)
		goto L60
	default:
		goto L59
	case 3:
		goto L61
	}
L59:
	;
	v409 = F_newkey(m, l0, l1, v385)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L64
	}
L60:
	;
	F_luaG_runerror(m, l0, v400, int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L6
	} else {
		goto L63
	}
L61:
	;
	v395 = m.G3
	v396 = *(*float64)(unsafe.Add(mBase, uint32(v385)))
	if base.F64_eq(v396, v396) != 0 {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v400 = v395 + int32(_a2687)
	goto L60
L63:
	;
	goto L59
L64:
	;
	v412 = v409
	goto L57
L65:
	;
	goto L54
L66:
	;
	m.G0 = v18 + int32(16)
	return
L67:
	;
	v448 = F_luaM_realloc_(m, l0, v20, int32(32)<<(uint(v21)%32), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	goto L66
}
func F_restartAOFAfterSYNC(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	v1 = F_startAppendOnly(m)
	mBase = m.M
	v2 = m.ExcPending
	if v2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	return
L3:
	;
	if v1 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v6 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v6 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v15 = F_sleep(m, int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	F__serverLog(m, int32(3), int32(_a1942), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v17 = F_startAppendOnly(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v17 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v22 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = F_sleep(m, int32(1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	F__serverLog(m, int32(3), int32(_a1942), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v33 = F_startAppendOnly(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v33 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v38 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v47 = F_sleep(m, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	F__serverLog(m, int32(3), int32(_a1942), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v49 = F_startAppendOnly(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	if v49 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v54 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v63 = F_sleep(m, int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L26
	}
L24:
	;
	F__serverLog(m, int32(3), int32(_a1942), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v65 = F_startAppendOnly(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	if v65 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v70 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v79 = F_sleep(m, int32(1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L32
	}
L30:
	;
	F__serverLog(m, int32(3), int32(_a1942), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v81 = F_startAppendOnly(m)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	if v81 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v86 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v95 = F_sleep(m, int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	F__serverLog(m, int32(3), int32(_a1942), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v97 = F_startAppendOnly(m)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	if v97 == int32(0) {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v102 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v111 = F_sleep(m, int32(1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L44
	}
L42:
	;
	F__serverLog(m, int32(3), int32(_a1942), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v113 = F_startAppendOnly(m)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	if v113 == int32(0) {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v118 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v127 = F_sleep(m, int32(1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L50
	}
L48:
	;
	F__serverLog(m, int32(3), int32(_a1942), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v129 = F_startAppendOnly(m)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	if v129 == int32(0) {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v134 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v143 = F_sleep(m, int32(1))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L56
	}
L54:
	;
	F__serverLog(m, int32(3), int32(_a1942), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v145 = F_startAppendOnly(m)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	if v145 == int32(0) {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v150 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v159 = F_sleep(m, int32(1))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L2
	} else {
		goto L62
	}
L60:
	;
	F__serverLog(m, int32(3), int32(_a1942), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v162 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F__serverLog(m, int32(3), int32(_a1941), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	goto L63
}
func F_restartAOFWithSyncRdb(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v243 int32
	_ = v243
	var v252 int64
	_ = v252
	var v260 int64
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
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
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	v10 = m.G0
	v12 = v10 - int32(128)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v15 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v12 + int32(128)
	return v411
L2:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _consts[58]))
	if v342 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L3:
	;
	v331 = int32(0)
	v333 = int32(-1)
	v334 = int32(1)
	v335 = int32(0)
	v336 = v331
	v337 = v331
	goto L2
L4:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	v305 = F_rename(m, v84, v304)
	mBase = m.M
	if v305 != int32(-1) {
		v327 = int32(1)
		goto L74
	} else {
		goto L75
	}
L5:
	;
	F__serverAssert(m, int32(_a153), int32(_a123), int32(1032))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L22
	} else {
		goto L73
	}
L6:
	;
	F__serverAssert(m, int32(_a137), int32(_a123), int32(1028))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L22
	} else {
		goto L72
	}
L7:
	;
	F__serverAssert(m, int32(_a152), int32(_a123), int32(1013))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L22
	} else {
		goto L71
	}
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v19 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(96)
	m.G0 = v24
	v28 = F_mkdir(m, v18, int32(493))
	mBase = m.M
	if v28 == v19 {
		v44 = v19
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	if v71 == int32(0) {
		goto L6
	} else {
		goto L24
	}
L10:
	;
	if v44 != int32(-1) {
		goto L9
	} else {
		goto L18
	}
L11:
	;
	m.G0 = v24 + int32(96)
	goto L10
L12:
	;
	v31 = F___errno_location(m)
	mBase = m.M
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 != int32(20) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v44 = int32(-1)
	goto L11
L14:
	;
	v35 = F_stat(m, v18, v24)
	mBase = m.M
	if v35 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(54)
	goto L13
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v36&int32(61440) == int32(16384) {
		v44 = v19
		goto L11
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v54 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	goto L20
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v61 = F___strerror_l(m, v60, v60)
	mBase = m.M
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v58
	F__serverLog(m, int32(3), int32(_a143), v12)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	goto L3
L24:
	;
	v74 = F_aofManifestDup(m, v71)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	v78 = F_getNewBaseFileNameAndMarkPreAsHistory(m, v74, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v78 == int32(0) {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v84 = F_makePath(m, v83, v78)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v86 = int32(-1)
	v88 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	v89 = F_rename(m, v88, v84)
	mBase = m.M
	if v89 != v86 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_markRewrittenIncrAofAsHistory(m, v74)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L22
	} else {
		goto L35
	}
L30:
	;
	v92 = int32(1)
	v93 = int32(0)
	v95 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v95 {
		v333 = v86
		v334 = v92
		v335 = v74
		v336 = v84
		v337 = v93
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	goto L32
L32:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v102 = F___strerror_l(m, v101, v101)
	mBase = m.M
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v99
	F__serverLog(m, int32(3), int32(_a157), v12+int32(16))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	v333 = v86
	v334 = v92
	v335 = v74
	v336 = v84
	v337 = v93
	goto L2
L35:
	;
	v115 = F_valkey_calloc(m, int32(24))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L22
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+16)) = int32(105)
	v119 = F_sdsempty(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L22
	} else {
		goto L37
	}
L37:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v74)+24))
	v125 = v123 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v74)+24)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(116)))) = int32(_a129)
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(112)))) = int32(_a134)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+104)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v122
	v140 = F_sdscatprintf(m, v119, int32(_a131), v12+int32(96))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L22
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v140
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v74)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v115)+8)) = v143
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v146 = F_listAddNodeTail(m, v145, v115)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L22
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+32)) = int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v153 = F_makePath(m, v151, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L22
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = int32(438)
	v160 = F_open(m, v153, int32(577), v12+int32(80))
	mBase = m.M
	if v160 != int32(-1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v74)+32))
	if v178 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v164 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	goto L44
L44:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v169 = F___strerror_l(m, v168, v168)
	mBase = m.M
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v152
	F__serverLog(m, int32(3), int32(_a141), v12+int32(48))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L22
	} else {
		goto L46
	}
L46:
	;
	goto L4
L47:
	;
	F_aofManifestFreeAndUpdate(m, v74)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L22
	} else {
		goto L53
	}
L48:
	;
	v181 = F_getAofManifestAsString(m, v74)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L22
	} else {
		goto L49
	}
L49:
	;
	v183 = F_writeAofManifestFile(m, v181)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L22
	} else {
		goto L50
	}
L50:
	;
	F_sdsfree(m, v181)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L22
	} else {
		goto L51
	}
L51:
	;
	if v183 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+32)) = int32(0)
	goto L47
L53:
	;
	v193 = F_aofDelHistoryFiles(m)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L22
	} else {
		goto L54
	}
L54:
	;
	F_sdsfree(m, v84)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L22
	} else {
		goto L55
	}
L55:
	;
	F_sdsfree(m, v153)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L22
	} else {
		goto L56
	}
L56:
	;
	F_bioDrainWorker(m, int32(1))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L22
	} else {
		goto L57
	}
L57:
	;
	v202 = int32(_a20)
	v204 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	*(*int64)(unsafe.Add(mBase, _consts[48])) = v204
	*(*int64)(unsafe.Add(mBase, _consts[49])) = v204
	v209 = *(*int32)(unsafe.Add(mBase, _consts[50]))
	if v209 != int32(-1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _consts[51]))
	if v225 != int32(-1) {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v213 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, _consts[50])) = int32(0)
	goto L58
L61:
	;
	F__serverLog(m, int32(3), int32(_a156), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L22
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v240 = int32(_a20)
	v241 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[52])) = v241
	v243 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v243
	*(*int64)(unsafe.Add(mBase, _consts[42])) = v241
	v252 = *(*int64)(unsafe.Add(mBase, _consts[54]))
	*(*int64)(unsafe.Add(mBase, _consts[55])) = v252
	*(*int64)(unsafe.Add(mBase, _consts[41])) = v241
	v260 = F_getAppendOnlyFileSize(m, v78, v243)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L22
	} else {
		goto L68
	}
L64:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v229 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, _consts[51])) = int32(0)
	goto L63
L66:
	;
	F__serverLog(m, int32(3), int32(_a155), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L22
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v262 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v160
	*(*int64)(unsafe.Add(mBase, _consts[56])) = v260
	*(*int64)(unsafe.Add(mBase, _consts[57])) = v260
	*(*int32)(unsafe.Add(mBase, _consts[38])) = int32(1)
	v272 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v272 {
		v411 = v243
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v78
	F__serverLog(m, int32(2), int32(_a154), v12+int32(64))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L22
	} else {
		goto L70
	}
L70:
	;
	v411 = v243
	goto L1
L71:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v333 = v160
	v334 = v327
	v335 = v74
	v336 = v84
	v337 = v153
	goto L2
L75:
	;
	v308 = int32(0)
	v310 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v310 {
		v327 = v308
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	goto L77
L77:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v317 = F___strerror_l(m, v316, v316)
	mBase = m.M
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v314
	F__serverLog(m, int32(3), int32(_a158), v12+int32(32))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L22
	} else {
		goto L79
	}
L79:
	;
	v327 = v308
	goto L74
L80:
	;
	if v333 == int32(-1) {
		goto L91
	} else {
		goto L92
	}
L81:
	;
	v345 = int32(0)
	v346 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	v348 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	goto L82
L82:
	;
	if base.B2i32(v346|v348 == v345) == int32(0) {
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v355 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	if v334 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	F__serverLog(m, int32(2), int32(_a159), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L22
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v365 = v364
	goto L89
L88:
	;
	v365 = v336
	goto L89
L89:
	;
	v366 = F_bg_unlink(m, v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L22
	} else {
		goto L90
	}
L90:
	;
	goto L80
L91:
	;
	if v337 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v370 = F_close(m, v333)
	mBase = m.M
	goto L91
L93:
	;
	if v335 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v373 = F_bg_unlink(m, v337)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L22
	} else {
		goto L95
	}
L95:
	;
	F_sdsfree(m, v337)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L22
	} else {
		goto L96
	}
L96:
	;
	goto L93
L97:
	;
	v404 = int32(-1)
	if v336 == int32(0) {
		v411 = v404
		goto L1
	} else {
		goto L112
	}
L98:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	if v379 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	if v390 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L100:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	if v382 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	F_valkey_free(m, v379)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L22
	} else {
		goto L104
	}
L102:
	;
	F_sdsfree(m, v382)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L22
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	goto L99
L105:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v335)+8))
	if v395 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	F_listRelease(m, v390)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L22
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	F_valkey_free(m, v335)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L22
	} else {
		goto L111
	}
L109:
	;
	F_listRelease(m, v395)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L22
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	goto L97
L112:
	;
	F_sdsfree(m, v336)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L22
	} else {
		goto L113
	}
L113:
	;
	v411 = v404
	goto L1
}
func F_restartServer(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v13 = F_access(m, v11, int32(1))
	mBase = m.M
	if v13 != int32(-1) {
		if l1&int32(2) == int32(0) {
			if l1&int32(1) == int32(0) {
				v82 = *(*int32)(unsafe.Add(mBase, _consts[750]))
				if v82 < int32(-1020) {
				} else {
					v89 = int32(3)
					for {
						v93 = F_fcntl(m, v89, int32(1), int32(0))
						mBase = m.M
						if v93 == int32(-1) {
						} else {
							v96 = F_close(m, v89)
							mBase = m.M
						}
						v98 = *(*int32)(unsafe.Add(mBase, _consts[750]))
						if v89 < v98+int32(1023) {
							v89 = v89 + int32(1)
							continue
						} else {
							break
						}
						break
					}
				}
				if l2 == int64(0) {
					v117 = *(*int32)(unsafe.Add(mBase, _consts[531]))
					v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
					F_valkey_free(m, v118)
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						v122 = *(*int32)(unsafe.Add(mBase, _consts[749]))
						v123 = F_zstrdup(m, v122)
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, _consts[531]))
							*(*int32)(unsafe.Add(mBase, uint32(v126))) = v123
							v132 = F___errno_location(m)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(45)
							F__Exit(m, int32(1))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				} else {
					v114 = F_usleep(m, base.I32_wrap_i64(l2)*int32(1000))
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return int32(0)
					} else {
						v117 = *(*int32)(unsafe.Add(mBase, _consts[531]))
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
						F_valkey_free(m, v118)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							v122 = *(*int32)(unsafe.Add(mBase, _consts[749]))
							v123 = F_zstrdup(m, v122)
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int32(0)
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, _consts[531]))
								*(*int32)(unsafe.Add(mBase, uint32(v126))) = v123
								v132 = F___errno_location(m)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(45)
								F__Exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v62 = F_prepareForShutdown(m, l0, int32(4))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					if v62 == int32(0) {
						v82 = *(*int32)(unsafe.Add(mBase, _consts[750]))
						if v82 < int32(-1020) {
						} else {
							v89 = int32(3)
							for {
								v93 = F_fcntl(m, v89, int32(1), int32(0))
								mBase = m.M
								if v93 == int32(-1) {
								} else {
									v96 = F_close(m, v89)
									mBase = m.M
								}
								v98 = *(*int32)(unsafe.Add(mBase, _consts[750]))
								if v89 < v98+int32(1023) {
									v89 = v89 + int32(1)
									continue
								} else {
									break
								}
								break
							}
						}
						if l2 == int64(0) {
							v117 = *(*int32)(unsafe.Add(mBase, _consts[531]))
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
							F_valkey_free(m, v118)
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								v122 = *(*int32)(unsafe.Add(mBase, _consts[749]))
								v123 = F_zstrdup(m, v122)
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int32(0)
								} else {
									v126 = *(*int32)(unsafe.Add(mBase, _consts[531]))
									*(*int32)(unsafe.Add(mBase, uint32(v126))) = v123
									v132 = F___errno_location(m)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(45)
									F__Exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						} else {
							v114 = F_usleep(m, base.I32_wrap_i64(l2)*int32(1000))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int32(0)
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, _consts[531]))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
								F_valkey_free(m, v118)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v122 = *(*int32)(unsafe.Add(mBase, _consts[749]))
									v123 = F_zstrdup(m, v122)
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return int32(0)
									} else {
										v126 = *(*int32)(unsafe.Add(mBase, _consts[531]))
										*(*int32)(unsafe.Add(mBase, uint32(v126))) = v123
										v132 = F___errno_location(m)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(45)
										F__Exit(m, int32(1))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, _consts[28]))
						if int32(3) < v67 {
							m.G0 = v8 + int32(32)
							return int32(-1)
						} else {
							F__serverLog(m, int32(3), int32(_a2199), int32(0))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(32)
								return int32(-1)
							}
						}
					}
				}
			}
		} else {
			v33 = int32(0)
			v34 = *(*int32)(unsafe.Add(mBase, _consts[690]))
			if v34 == v33 {
				if l1&int32(1) == int32(0) {
					v82 = *(*int32)(unsafe.Add(mBase, _consts[750]))
					if v82 < int32(-1020) {
					} else {
						v89 = int32(3)
						for {
							v93 = F_fcntl(m, v89, int32(1), int32(0))
							mBase = m.M
							if v93 == int32(-1) {
							} else {
								v96 = F_close(m, v89)
								mBase = m.M
							}
							v98 = *(*int32)(unsafe.Add(mBase, _consts[750]))
							if v89 < v98+int32(1023) {
								v89 = v89 + int32(1)
								continue
							} else {
								break
							}
							break
						}
					}
					if l2 == int64(0) {
						v117 = *(*int32)(unsafe.Add(mBase, _consts[531]))
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
						F_valkey_free(m, v118)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							v122 = *(*int32)(unsafe.Add(mBase, _consts[749]))
							v123 = F_zstrdup(m, v122)
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int32(0)
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, _consts[531]))
								*(*int32)(unsafe.Add(mBase, uint32(v126))) = v123
								v132 = F___errno_location(m)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(45)
								F__Exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					} else {
						v114 = F_usleep(m, base.I32_wrap_i64(l2)*int32(1000))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, _consts[531]))
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
							F_valkey_free(m, v118)
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								v122 = *(*int32)(unsafe.Add(mBase, _consts[749]))
								v123 = F_zstrdup(m, v122)
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int32(0)
								} else {
									v126 = *(*int32)(unsafe.Add(mBase, _consts[531]))
									*(*int32)(unsafe.Add(mBase, uint32(v126))) = v123
									v132 = F___errno_location(m)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(45)
									F__Exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v62 = F_prepareForShutdown(m, l0, int32(4))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						if v62 == int32(0) {
							v82 = *(*int32)(unsafe.Add(mBase, _consts[750]))
							if v82 < int32(-1020) {
							} else {
								v89 = int32(3)
								for {
									v93 = F_fcntl(m, v89, int32(1), int32(0))
									mBase = m.M
									if v93 == int32(-1) {
									} else {
										v96 = F_close(m, v89)
										mBase = m.M
									}
									v98 = *(*int32)(unsafe.Add(mBase, _consts[750]))
									if v89 < v98+int32(1023) {
										v89 = v89 + int32(1)
										continue
									} else {
										break
									}
									break
								}
							}
							if l2 == int64(0) {
								v117 = *(*int32)(unsafe.Add(mBase, _consts[531]))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
								F_valkey_free(m, v118)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v122 = *(*int32)(unsafe.Add(mBase, _consts[749]))
									v123 = F_zstrdup(m, v122)
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return int32(0)
									} else {
										v126 = *(*int32)(unsafe.Add(mBase, _consts[531]))
										*(*int32)(unsafe.Add(mBase, uint32(v126))) = v123
										v132 = F___errno_location(m)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(45)
										F__Exit(m, int32(1))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							} else {
								v114 = F_usleep(m, base.I32_wrap_i64(l2)*int32(1000))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									v117 = *(*int32)(unsafe.Add(mBase, _consts[531]))
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
									F_valkey_free(m, v118)
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										v122 = *(*int32)(unsafe.Add(mBase, _consts[749]))
										v123 = F_zstrdup(m, v122)
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return int32(0)
										} else {
											v126 = *(*int32)(unsafe.Add(mBase, _consts[531]))
											*(*int32)(unsafe.Add(mBase, uint32(v126))) = v123
											v132 = F___errno_location(m)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(45)
											F__Exit(m, int32(1))
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, _consts[28]))
							if int32(3) < v67 {
								m.G0 = v8 + int32(32)
								return int32(-1)
							} else {
								F__serverLog(m, int32(3), int32(_a2199), int32(0))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(32)
									return int32(-1)
								}
							}
						}
					}
				}
			} else {
				v38 = F_rewriteConfig(m, v34, int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					if v38 != int32(-1) {
						if l1&int32(1) == int32(0) {
							v82 = *(*int32)(unsafe.Add(mBase, _consts[750]))
							if v82 < int32(-1020) {
							} else {
								v89 = int32(3)
								for {
									v93 = F_fcntl(m, v89, int32(1), int32(0))
									mBase = m.M
									if v93 == int32(-1) {
									} else {
										v96 = F_close(m, v89)
										mBase = m.M
									}
									v98 = *(*int32)(unsafe.Add(mBase, _consts[750]))
									if v89 < v98+int32(1023) {
										v89 = v89 + int32(1)
										continue
									} else {
										break
									}
									break
								}
							}
							if l2 == int64(0) {
								v117 = *(*int32)(unsafe.Add(mBase, _consts[531]))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
								F_valkey_free(m, v118)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v122 = *(*int32)(unsafe.Add(mBase, _consts[749]))
									v123 = F_zstrdup(m, v122)
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return int32(0)
									} else {
										v126 = *(*int32)(unsafe.Add(mBase, _consts[531]))
										*(*int32)(unsafe.Add(mBase, uint32(v126))) = v123
										v132 = F___errno_location(m)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(45)
										F__Exit(m, int32(1))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							} else {
								v114 = F_usleep(m, base.I32_wrap_i64(l2)*int32(1000))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									v117 = *(*int32)(unsafe.Add(mBase, _consts[531]))
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
									F_valkey_free(m, v118)
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										v122 = *(*int32)(unsafe.Add(mBase, _consts[749]))
										v123 = F_zstrdup(m, v122)
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return int32(0)
										} else {
											v126 = *(*int32)(unsafe.Add(mBase, _consts[531]))
											*(*int32)(unsafe.Add(mBase, uint32(v126))) = v123
											v132 = F___errno_location(m)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(45)
											F__Exit(m, int32(1))
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v62 = F_prepareForShutdown(m, l0, int32(4))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								if v62 == int32(0) {
									v82 = *(*int32)(unsafe.Add(mBase, _consts[750]))
									if v82 < int32(-1020) {
									} else {
										v89 = int32(3)
										for {
											v93 = F_fcntl(m, v89, int32(1), int32(0))
											mBase = m.M
											if v93 == int32(-1) {
											} else {
												v96 = F_close(m, v89)
												mBase = m.M
											}
											v98 = *(*int32)(unsafe.Add(mBase, _consts[750]))
											if v89 < v98+int32(1023) {
												v89 = v89 + int32(1)
												continue
											} else {
												break
											}
											break
										}
									}
									if l2 == int64(0) {
										v117 = *(*int32)(unsafe.Add(mBase, _consts[531]))
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
										F_valkey_free(m, v118)
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											v122 = *(*int32)(unsafe.Add(mBase, _consts[749]))
											v123 = F_zstrdup(m, v122)
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return int32(0)
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, _consts[531]))
												*(*int32)(unsafe.Add(mBase, uint32(v126))) = v123
												v132 = F___errno_location(m)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(45)
												F__Exit(m, int32(1))
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										v114 = F_usleep(m, base.I32_wrap_i64(l2)*int32(1000))
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return int32(0)
										} else {
											v117 = *(*int32)(unsafe.Add(mBase, _consts[531]))
											v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
											F_valkey_free(m, v118)
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int32(0)
											} else {
												v122 = *(*int32)(unsafe.Add(mBase, _consts[749]))
												v123 = F_zstrdup(m, v122)
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int32(0)
												} else {
													v126 = *(*int32)(unsafe.Add(mBase, _consts[531]))
													*(*int32)(unsafe.Add(mBase, uint32(v126))) = v123
													v132 = F___errno_location(m)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(45)
													F__Exit(m, int32(1))
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, _consts[28]))
									if int32(3) < v67 {
										m.G0 = v8 + int32(32)
										return int32(-1)
									} else {
										F__serverLog(m, int32(3), int32(_a2199), int32(0))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											m.G0 = v8 + int32(32)
											return int32(-1)
										}
									}
								}
							}
						}
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, _consts[28]))
						if int32(3) < v43 {
							m.G0 = v8 + int32(32)
							return int32(-1)
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, _consts[18]))
							v48 = F___strerror_l(m, v47, v47)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v48
							F__serverLog(m, int32(3), int32(_a2200), v8+int32(16))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(32)
								return int32(-1)
							}
						}
					}
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		if int32(3) < v17 {
			m.G0 = v8 + int32(32)
			return int32(-1)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[749]))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v21
			F__serverLog(m, int32(3), int32(_a2201), v8)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(32)
				return int32(-1)
			}
		}
	}
}
func F_roundl(m *base.Module, l0 int32, l1 int64, l2 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int64
	_ = v22
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v41 int64
	_ = v41
	var v46 int64
	_ = v46
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v75 int64
	_ = v75
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v95 int32
	_ = v95
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v146 int64
	_ = v146
	var v151 int64
	_ = v151
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v178 int32
	_ = v178
	var v202 int32
	_ = v202
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v243 int64
	_ = v243
	var v244 int64
	_ = v244
	var v245 int64
	_ = v245
	var v246 int64
	_ = v246
	var v252 int64
	_ = v252
	var v253 int64
	_ = v253
	var v254 int64
	_ = v254
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	v15 = base.I32_wrap_i64(int64(base.Ui64(l2) >> (uint(int64(48)) % 64)))
	v17 = v15 & int32(32767)
	if base.Ui32(int32(16494)) < base.Ui32(v17) {
		v253 = l1
		v254 = l2
	} else {
		if base.Ui32(int32(16381)) < base.Ui32(v17) {
			if base.Ui32(v15) < base.Ui32(int32(32768)) {
				v35 = l2
			} else {
				v35 = l2 ^ int64(-9223372036854775807-1)
			}
			v36 = int64(0)
			F___addtf3(m, v11+int32(112), l1, v35, v36, int64(4642929740842270720))
			mBase = m.M
			v41 = *(*int64)(unsafe.Add(mBase, uint32(v11)+112))
			v46 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(120))))
			F___addtf3(m, v11+int32(96), v41, v46, v36, int64(-4580442296012505088))
			mBase = m.M
			v51 = v11 + int32(80)
			v52 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
			v57 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(104))))
			v59 = m.G0
			v60 = int32(16)
			v61 = v59 - v60
			m.G0 = v61
			F___addtf3(m, v61, v52, v57, l1, v35^int64(-9223372036854775807-1))
			mBase = m.M
			v66 = *(*int64)(unsafe.Add(mBase, uint32(v61)))
			v69 = *(*int64)(unsafe.Add(mBase, uint32(v61+int32(8))))
			*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v69
			*(*int64)(unsafe.Add(mBase, uint32(v51))) = v66
			m.G0 = v61 + v60
			v75 = *(*int64)(unsafe.Add(mBase, uint32(v11)+80))
			v80 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(88))))
			v81 = int64(0)
			v82 = int64(4611123068473966592)
			v86 = int32(-1)
			v90 = v80 & int64(9223372036854775807)
			v91 = int64(9223090561878065152)
			if v90 == v91 {
				v95 = base.B2i32(v75 != v81)
			} else {
				v95 = base.B2i32(base.Ui64(v91) < base.Ui64(v90))
			}
			if v95 != 0 {
				v134 = v86
				v138 = v134
			} else {
				if base.B2i32(v81|v75|(int64(4611123068473966592)|v90) == int64(0)) == int32(0) {
					if v82&v80 < int64(0) {
						if v80 == v82 {
							v128 = base.B2i32(base.Ui64(v81) < base.Ui64(v75))
						} else {
							v128 = base.B2i32(v82 < v80)
						}
						if v128 != 0 {
							v134 = v86
						} else {
							v134 = base.B2i32(v75^v81|(v80^v82) != int64(0))
						}
						v138 = v134
					} else {
						if v80 == v82 {
							v119 = base.B2i32(base.Ui64(v75) < base.Ui64(v81))
						} else {
							v119 = base.B2i32(v80 < v82)
						}
						if v119 != 0 {
							v134 = v86
							v138 = v134
						} else {
							v138 = base.B2i32(v75^v81|(v80^v82) != int64(0))
						}
					}
				} else {
					v138 = int32(0)
				}
			}
			if v138 < int32(1) {
				F___addtf3(m, v11+int32(64), l1, v35, v75, v80)
				mBase = m.M
				v164 = int64(0)
				v165 = int64(-4612248968380809216)
				v173 = v80 & int64(9223372036854775807)
				v174 = int64(9223090561878065152)
				if v173 == v174 {
					v178 = base.B2i32(v75 != v164)
				} else {
					v178 = base.B2i32(base.Ui64(v174) < base.Ui64(v173))
				}
				if v178 != 0 {
					v223 = int32(1)
					v227 = v223
				} else {
					if base.B2i32(v164|v75|(int64(4611123068473966592)|v173) == int64(0)) == int32(0) {
						if v165&v80 < int64(0) {
							if v80 == v165 {
								v214 = base.B2i32(base.Ui64(v164) < base.Ui64(v75))
							} else {
								v214 = base.B2i32(v165 < v80)
							}
							if v214 == int32(0) {
								v223 = base.B2i32(v75^v164|(v80^v165) != int64(0))
								v227 = v223
							} else {
								v227 = int32(-1)
							}
						} else {
							if v80 == v165 {
								v202 = base.B2i32(base.Ui64(v75) < base.Ui64(v164))
							} else {
								v202 = base.B2i32(v80 < v165)
							}
							if v202 == int32(0) {
								v227 = base.B2i32(v75^v164|(v80^v165) != int64(0))
							} else {
								v227 = int32(-1)
							}
						}
					} else {
						v227 = int32(0)
					}
				}
				v232 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(72))))
				v233 = *(*int64)(unsafe.Add(mBase, uint32(v11)+64))
				if int32(0) < v227 {
					v245 = v233
					v246 = v232
				} else {
					F___addtf3(m, v11+int32(48), v233, v232, int64(0), int64(4611404543450677248))
					mBase = m.M
					v243 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(56))))
					v244 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
					v245 = v244
					v246 = v243
				}
			} else {
				F___addtf3(m, v11+int32(32), l1, v35, v75, v80)
				mBase = m.M
				v146 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
				v151 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(40))))
				F___addtf3(m, v11+int32(16), v146, v151, int64(0), int64(-4611967493404098560))
				mBase = m.M
				v159 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(24))))
				v160 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
				v245 = v160
				v246 = v159
			}
			if base.Ui32(v15) < base.Ui32(int32(32768)) {
				v252 = v246
			} else {
				v252 = v246 ^ int64(-9223372036854775807-1)
			}
			v253 = v245
			v254 = v252
		} else {
			v22 = int64(0)
			F___multf3(m, v11, l1, l2, v22, v22)
			mBase = m.M
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(8))))
			v28 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
			v253 = v28
			v254 = v27
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v253
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v254
	m.G0 = v11 + int32(128)
	return
}
func F_rpushxCommand(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = int32(1)
	F_pushGenericCommand(m, l0, v2, v2)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
