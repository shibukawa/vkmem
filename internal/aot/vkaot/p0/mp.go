package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F_mp_decode_to_lua_type(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int64
	_ = v209
	var v210 int64
	_ = v210
	var v212 int64
	_ = v212
	var v214 int64
	_ = v214
	var v217 int64
	_ = v217
	var v219 int64
	_ = v219
	var v221 int64
	_ = v221
	var v223 int64
	_ = v223
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int64
	_ = v268
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v273 int64
	_ = v273
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v422 int64
	_ = v422
	var v423 int64
	_ = v423
	var v425 int64
	_ = v425
	var v427 int64
	_ = v427
	var v430 int64
	_ = v430
	var v432 int64
	_ = v432
	var v434 int64
	_ = v434
	var v436 int64
	_ = v436
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v746 int32
	_ = v746
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v8 != 0 {
		v12 = m.G3
		F_luaL_checkstack(m, l0, int32(1), v12+int32(_a_F_mp_decode_to_lua_type_0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			switch v18 + int32(-192) {
			case 0:
				v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v323)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v323 + int32(16)
				v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v329 + int32(1)
				v333 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v333 + int32(-1)
				return
			default:
				v654 = base.I32_extend8_s(v18)
				if v654 < int32(0) {
					if v654&int32(-32) == int32(-96) {
						v703 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v705 = v18 & int32(31)
						if base.Ui32(v705) < base.Ui32(v703) {
							F_lua_pushlstring(m, l0, v17+int32(1), v705)
							mBase = m.M
							v712 = m.ExcPending
							if v712 != 0 {
								return
							} else {
								v713 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v715 = v705 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v713 + v715
								v718 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v718 - v715
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
							return
						}
					} else {
						v681 = int32(224)
						if v654&v681 != v681 {
							switch v654&int32(240) + int32(-128) {
							case 0:
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17 + int32(1)
								v739 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v739 + int32(-1)
								F_mp_decode_to_lua_hash(m, l0, l1, v18&int32(15))
								mBase = m.M
								v746 = m.ExcPending
								if v746 != 0 {
									return
								} else {
									return
								}
							default:
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(2)
								return
							case 16:
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17 + int32(1)
								v728 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v728 + int32(-1)
								F_mp_decode_to_lua_array(m, l0, l1, v18&int32(15))
								mBase = m.M
								v735 = m.ExcPending
								if v735 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							v686 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v686)+8)) = int32(3)
							*(*float64)(unsafe.Add(mBase, uint32(v686))) = base.F64_convert_i32_s(v654)
							v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v691 + int32(16)
							v695 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v695 + int32(1)
							v699 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v699 + int32(-1)
							return
						}
					}
				} else {
					v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v661)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v661))) = base.F64_convert_i32_u(v654 & int32(255))
					v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v665 + int32(16)
					v669 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v669 + int32(1)
					v673 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v673 + int32(-1)
					return
				}
			case 2:
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v359)+8)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v359))) = int32(0)
				v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v365 + int32(16)
				v369 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v369 + int32(1)
				v373 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v373 + int32(-1)
				return
			case 3:
				v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v339)+8)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v339))) = int32(1)
				v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v345 + int32(16)
				v349 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v349 + int32(1)
				v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v353 + int32(-1)
				return
			case 10:
				v377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(4)) < base.Ui32(v377) {
					v382 = *(*int32)(unsafe.Add(mBase, uint32(v17)+1))
					v383 = int32(24)
					v385 = int32(65280)
					v387 = int32(8)
					v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v401)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v401))) = base.F64_promote_f32(base.F32_reinterpret_i32(v382<<(uint(v383)%32) | v382&v385<<(uint(v387)%32) | (int32(base.Ui32(v382)>>(uint(v387)%32))&v385 | int32(base.Ui32(v382)>>(uint(v383)%32)))))
					v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v405 + int32(16)
					v409 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v409 + int32(5)
					v413 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v413 + int32(-5)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 11:
				v417 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(8)) < base.Ui32(v417) {
					v422 = *(*int64)(unsafe.Add(mBase, uint32(v17)+1))
					v423 = int64(56)
					v425 = int64(65280)
					v427 = int64(40)
					v430 = int64(16711680)
					v432 = int64(24)
					v434 = int64(4278190080)
					v436 = int64(8)
					v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v460)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v460))) = base.F64_reinterpret_i64(v422<<(uint(v423)%64) | v422&v425<<(uint(v427)%64) | (v422&v430<<(uint(v432)%64) | v422&v434<<(uint(v436)%64)) | (int64(base.Ui64(v422)>>(uint(v436)%64))&v434 | int64(base.Ui64(v422)>>(uint(v432)%64))&v430 | (int64(base.Ui64(v422)>>(uint(v427)%64))&v425 | int64(base.Ui64(v422)>>(uint(v423)%64)))))
					v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v464 + int32(16)
					v468 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v468 + int32(9)
					v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v472 + int32(-9)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 12:
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(1)) < base.Ui32(v21) {
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v29))) = base.F64_convert_i32_u(v26)
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v33 + int32(16)
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37 + int32(2)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v41 + int32(-2)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 13:
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(2)) < base.Ui32(v69) {
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
					v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v81))) = base.F64_convert_i32_u(v74<<(uint(int32(8))%32) | v77)
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v85 + int32(16)
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v89 + int32(3)
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v93 + int32(-3)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 14:
				v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(4)) < base.Ui32(v126) {
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v17)+1))
					v132 = int32(24)
					v134 = int32(65280)
					v136 = int32(8)
					v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v149)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v149))) = base.F64_convert_i32_u(v131<<(uint(v132)%32) | v131&v134<<(uint(v136)%32) | (int32(base.Ui32(v131)>>(uint(v136)%32))&v134 | int32(base.Ui32(v131)>>(uint(v132)%32))))
					v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v153 + int32(16)
					v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v157 + int32(5)
					v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v161 + int32(-5)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 15:
				v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(8)) < base.Ui32(v204) {
					v209 = *(*int64)(unsafe.Add(mBase, uint32(v17)+1))
					v210 = int64(56)
					v212 = int64(65280)
					v214 = int64(40)
					v217 = int64(16711680)
					v219 = int64(24)
					v221 = int64(4278190080)
					v223 = int64(8)
					v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v247)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v247))) = base.F64_convert_i64_u(v209<<(uint(v210)%64) | v209&v212<<(uint(v214)%64) | (v209&v217<<(uint(v219)%64) | v209&v221<<(uint(v223)%64)) | (int64(base.Ui64(v209)>>(uint(v223)%64))&v221 | int64(base.Ui64(v209)>>(uint(v219)%64))&v217 | (int64(base.Ui64(v209)>>(uint(v214)%64))&v212 | int64(base.Ui64(v209)>>(uint(v210)%64)))))
					v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v251 + int32(16)
					v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v255 + int32(9)
					v259 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v259 + int32(-9)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 16:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(1)) < base.Ui32(v45) {
					v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+1)))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v52))) = base.F64_convert_i32_s(v50)
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v57 + int32(16)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v61 + int32(2)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v65 + int32(-2)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 17:
				v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(2)) < base.Ui32(v97) {
					v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
					v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v109))) = base.F64_convert_i32_s(base.I32_extend16_s(v102<<(uint(int32(8))%32)) | v106)
					v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v114 + int32(16)
					v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v118 + int32(3)
					v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v122 + int32(-3)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 18:
				v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(4)) < base.Ui32(v165) {
					v170 = *(*int32)(unsafe.Add(mBase, uint32(v17)+1))
					v171 = int32(24)
					v173 = int32(65280)
					v175 = int32(8)
					v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v187)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v187))) = base.F64_convert_i32_s(v170<<(uint(v171)%32) | v170&v173<<(uint(v175)%32) | (int32(base.Ui32(v170)>>(uint(v175)%32))&v173 | int32(base.Ui32(v170)>>(uint(v171)%32))))
					v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v192 + int32(16)
					v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v196 + int32(5)
					v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v200 + int32(-5)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 19:
				v263 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(8)) < base.Ui32(v263) {
					v268 = *(*int64)(unsafe.Add(mBase, uint32(v17)+1))
					v269 = int64(56)
					v271 = int64(65280)
					v273 = int64(40)
					v276 = int64(16711680)
					v278 = int64(24)
					v280 = int64(4278190080)
					v282 = int64(8)
					v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v306)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v306))) = base.F64_convert_i64_s(v268<<(uint(v269)%64) | v268&v271<<(uint(v273)%64) | (v268&v276<<(uint(v278)%64) | v268&v280<<(uint(v282)%64)) | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))&v276 | (int64(base.Ui64(v268)>>(uint(v273)%64))&v271 | int64(base.Ui64(v268)>>(uint(v269)%64)))))
					v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v310 + int32(16)
					v314 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v314 + int32(9)
					v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v318 + int32(-9)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 25:
				v476 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(1)) < base.Ui32(v476) {
					v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
					v483 = v481 + int32(2)
					if base.Ui32(v483) <= base.Ui32(v476) {
						F_lua_pushlstring(m, l0, v17+int32(2), v481)
						mBase = m.M
						v490 = m.ExcPending
						if v490 != 0 {
							return
						} else {
							v491 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v491 + v483
							v494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v494 - v483
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 26:
				v497 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(2)) < base.Ui32(v497) {
					v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
					v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
					v506 = v502<<(uint(int32(8))%32) | v505
					v508 = v506 + int32(3)
					if base.Ui32(v508) <= base.Ui32(v497) {
						F_lua_pushlstring(m, l0, v17+int32(3), v506)
						mBase = m.M
						v515 = m.ExcPending
						if v515 != 0 {
							return
						} else {
							v516 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v516 + v508
							v519 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v519 - v508
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 27:
				v522 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(4)) < base.Ui32(v522) {
					v527 = *(*int32)(unsafe.Add(mBase, uint32(v17)+1))
					v529 = v522 + int32(-5)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v529
					v532 = v17 + int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v532
					v534 = int32(24)
					v536 = int32(65280)
					v538 = int32(8)
					v548 = v527<<(uint(v534)%32) | v527&v536<<(uint(v538)%32) | (int32(base.Ui32(v527)>>(uint(v538)%32))&v536 | int32(base.Ui32(v527)>>(uint(v534)%32)))
					if base.Ui32(v548) <= base.Ui32(v529) {
						F_lua_pushlstring(m, l0, v532, v548)
						mBase = m.M
						v553 = m.ExcPending
						if v553 != 0 {
							return
						} else {
							v554 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v554 + v548
							v557 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v557 - v548
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 28:
				v560 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(2)) < base.Ui32(v560) {
					v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
					v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v560 + int32(-3)
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17 + int32(3)
					F_mp_decode_to_lua_array(m, l0, l1, v565|v566<<(uint(int32(8))%32))
					mBase = m.M
					v577 = m.ExcPending
					if v577 != 0 {
						return
					} else {
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 29:
				v578 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(4)) < base.Ui32(v578) {
					v583 = *(*int32)(unsafe.Add(mBase, uint32(v17)+1))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v578 + int32(-5)
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17 + int32(5)
					v590 = int32(24)
					v592 = int32(65280)
					v594 = int32(8)
					F_mp_decode_to_lua_array(m, l0, l1, v583<<(uint(v590)%32)|v583&v592<<(uint(v594)%32)|(int32(base.Ui32(v583)>>(uint(v594)%32))&v592|int32(base.Ui32(v583)>>(uint(v590)%32))))
					mBase = m.M
					v606 = m.ExcPending
					if v606 != 0 {
						return
					} else {
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 30:
				v607 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(2)) < base.Ui32(v607) {
					v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
					v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v607 + int32(-3)
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17 + int32(3)
					F_mp_decode_to_lua_hash(m, l0, l1, v612|v613<<(uint(int32(8))%32))
					mBase = m.M
					v624 = m.ExcPending
					if v624 != 0 {
						return
					} else {
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			case 31:
				v625 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(int32(4)) < base.Ui32(v625) {
					v630 = *(*int32)(unsafe.Add(mBase, uint32(v17)+1))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v625 + int32(-5)
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17 + int32(5)
					v637 = int32(24)
					v639 = int32(65280)
					v641 = int32(8)
					F_mp_decode_to_lua_hash(m, l0, l1, v630<<(uint(v637)%32)|v630&v639<<(uint(v641)%32)|(int32(base.Ui32(v630)>>(uint(v641)%32))&v639|int32(base.Ui32(v630)>>(uint(v637)%32))))
					mBase = m.M
					v653 = m.ExcPending
					if v653 != 0 {
						return
					} else {
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
					return
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
		return
	}
}
func F_mp_encode_array(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	v3 = l2
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if base.Ui64(int64(15)) < base.Ui64(v3) {
		if base.Ui64(int64(65535)) < base.Ui64(v3) {
			v30 = int32(221)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)) = uint8(v30)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)) = uint8(v3)
			v34 = int64(base.Ui64(v3) >> (uint(int64(8)) % 64))
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v34)
			v37 = int64(base.Ui64(v3) >> (uint(int64(16)) % 64))
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)) = uint8(v37)
			v40 = int64(base.Ui64(v3) >> (uint(int64(24)) % 64))
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v40)
			v43 = int32(5)
		} else {
			v23 = int32(220)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)) = uint8(v23)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)) = uint8(v3)
			v27 = int64(base.Ui64(v3) >> (uint(int64(8)) % 64))
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v27)
			v43 = int32(3)
		}
	} else {
		v18 = base.I32_wrap_i64(v3) | int32(144)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)) = uint8(v18)
		v43 = int32(1)
	}
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v44) < base.Ui32(v43) {
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v49 = v48 + v43
		if base.Ui32(v49) < base.Ui32(v48) {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		} else {
			if base.Ui32(int32(2147483647)) <= base.Ui32(v49) {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v55 = v12 + int32(12)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v55 == int32(0) {
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v55))) = v59
				}
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
				v65 = v49 << (uint(int32(1)) % 32)
				v66 = m.T0[v61].(func(*base.Module, int32, int32, int32, int32) int32)(m, v62, v53, v48+v44, v65)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v66
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v65 - v69
					v72 = v69
					v73 = v66
					if v43 == int32(0) {
					} else {
						v82 = F__emscripten_memcpy_bulkmem(m, v73+v72, v12+int32(7), v43)
						mBase = m.M
					}
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v84 + v43
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v87 - v43
					m.G0 = v12 + int32(16)
					return
				}
			}
		}
	} else {
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v72 = v46
		v73 = v47
		if v43 == int32(0) {
		} else {
			v82 = F__emscripten_memcpy_bulkmem(m, v73+v72, v12+int32(7), v43)
			mBase = m.M
		}
		v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v84 + v43
		v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v87 - v43
		m.G0 = v12 + int32(16)
		return
	}
}
func F_mp_encode_int(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int64
	_ = v30
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int64
	_ = v86
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v98 int64
	_ = v98
	var v101 int64
	_ = v101
	var v104 int32
	_ = v104
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v123 int64
	_ = v123
	var v126 int64
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
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
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	v3 = l2
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if v3 < int64(0) {
		if base.Ui64(int64(-32)) <= base.Ui64(v3) {
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+3)) = uint8(v3)
			v131 = int32(1)
		} else {
			if base.Ui64(v3) < base.Ui64(int64(-128)) {
				if base.Ui64(v3) < base.Ui64(int64(-32768)) {
					if base.Ui64(v3) < base.Ui64(int64(-2147483648)) {
						v104 = int32(211)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+3)) = uint8(v104)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)) = uint8(v3)
						v108 = int64(base.Ui64(v3) >> (uint(int64(8)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v108)
						v111 = int64(base.Ui64(v3) >> (uint(int64(16)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)) = uint8(v111)
						v114 = int64(base.Ui64(v3) >> (uint(int64(24)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v114)
						v117 = int64(base.Ui64(v3) >> (uint(int64(32)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)) = uint8(v117)
						v120 = int64(base.Ui64(v3) >> (uint(int64(40)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+6)) = uint8(v120)
						v123 = int64(base.Ui64(v3) >> (uint(int64(48)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)) = uint8(v123)
						v126 = int64(base.Ui64(v3) >> (uint(int64(56)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)) = uint8(v126)
						v131 = int32(9)
					} else {
						v91 = int32(210)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+3)) = uint8(v91)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)) = uint8(v3)
						v95 = int64(base.Ui64(v3) >> (uint(int64(8)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+6)) = uint8(v95)
						v98 = int64(base.Ui64(v3) >> (uint(int64(16)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)) = uint8(v98)
						v101 = int64(base.Ui64(v3) >> (uint(int64(24)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)) = uint8(v101)
						v131 = int32(5)
					}
				} else {
					v82 = int32(209)
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+3)) = uint8(v82)
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)) = uint8(v3)
					v86 = int64(base.Ui64(v3) >> (uint(int64(8)) % 64))
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)) = uint8(v86)
					v131 = int32(3)
				}
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)) = uint8(v3)
				v77 = int32(208)
				*(*uint8)(unsafe.Add(mBase, uint32(v12)+3)) = uint8(v77)
				v131 = int32(2)
			}
		}
	} else {
		if base.Ui64(v3) <= base.Ui64(int64(127)) {
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+3)) = uint8(v3)
			v131 = int32(1)
		} else {
			if base.Ui64(int64(255)) < base.Ui64(v3) {
				if base.Ui64(int64(65535)) < base.Ui64(v3) {
					if base.Ui64(int64(4294967295)) < base.Ui64(v3) {
						v48 = int32(207)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+3)) = uint8(v48)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)) = uint8(v3)
						v52 = int64(base.Ui64(v3) >> (uint(int64(8)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v52)
						v55 = int64(base.Ui64(v3) >> (uint(int64(16)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)) = uint8(v55)
						v58 = int64(base.Ui64(v3) >> (uint(int64(24)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v58)
						v61 = int64(base.Ui64(v3) >> (uint(int64(32)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)) = uint8(v61)
						v64 = int64(base.Ui64(v3) >> (uint(int64(40)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+6)) = uint8(v64)
						v67 = int64(base.Ui64(v3) >> (uint(int64(48)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)) = uint8(v67)
						v70 = int64(base.Ui64(v3) >> (uint(int64(56)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)) = uint8(v70)
						v131 = int32(9)
					} else {
						v35 = int32(206)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+3)) = uint8(v35)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)) = uint8(v3)
						v39 = int64(base.Ui64(v3) >> (uint(int64(8)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+6)) = uint8(v39)
						v42 = int64(base.Ui64(v3) >> (uint(int64(16)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)) = uint8(v42)
						v45 = int64(base.Ui64(v3) >> (uint(int64(24)) % 64))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)) = uint8(v45)
						v131 = int32(5)
					}
				} else {
					v26 = int32(205)
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+3)) = uint8(v26)
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)) = uint8(v3)
					v30 = int64(base.Ui64(v3) >> (uint(int64(8)) % 64))
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)) = uint8(v30)
					v131 = int32(3)
				}
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)) = uint8(v3)
				v21 = int32(204)
				*(*uint8)(unsafe.Add(mBase, uint32(v12)+3)) = uint8(v21)
				v131 = int32(2)
			}
		}
	}
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v132) < base.Ui32(v131) {
		v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v137 = v136 + v131
		if base.Ui32(v137) < base.Ui32(v136) {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		} else {
			if base.Ui32(int32(2147483647)) <= base.Ui32(v137) {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			} else {
				v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v143 = v12 + int32(12)
				v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v143 == int32(0) {
				} else {
					v147 = *(*int32)(unsafe.Add(mBase, uint32(v144)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v143))) = v147
				}
				v149 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
				v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
				v153 = v137 << (uint(int32(1)) % 32)
				v154 = m.T0[v149].(func(*base.Module, int32, int32, int32, int32) int32)(m, v150, v141, v136+v132, v153)
				mBase = m.M
				v155 = m.ExcPending
				if v155 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v154
					v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v153 - v157
					v160 = v157
					v161 = v154
					if v131 == int32(0) {
					} else {
						v170 = F__emscripten_memcpy_bulkmem(m, v161+v160, v12+int32(3), v131)
						mBase = m.M
					}
					v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v172 + v131
					v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v175 - v131
					m.G0 = v12 + int32(16)
					return
				}
			}
		}
	} else {
		v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v160 = v134
		v161 = v135
		if v131 == int32(0) {
		} else {
			v170 = F__emscripten_memcpy_bulkmem(m, v161+v160, v12+int32(3), v131)
			mBase = m.M
		}
		v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v172 + v131
		v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v175 - v131
		m.G0 = v12 + int32(16)
		return
	}
}
func F_mp_encode_lua_type(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v95 float64
	_ = v95
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v108 float64
	_ = v108
	var v109 int32
	_ = v109
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v200 int32
	_ = v200
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v33 = v30 + int32(-16)
	v67 = m.G398
	if v33 != v67 {
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
		v73 = v70
	} else {
		v73 = int32(-1)
	}
	if v73 == int32(5) {
		v76 = int32(0)
	} else {
		v76 = v73
	}
	if l2 == int32(16) {
		v79 = v76
	} else {
		v79 = v73
	}
	switch v79 + int32(-1) {
	case 0:
		F_mp_encode_lua_bool(m, l0, l1)
		mBase = m.M
		v91 = m.ExcPending
		if v91 != 0 {
			return
		} else {
			v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v200 + int32(-16)
			m.G0 = v13 + int32(16)
			return
		}
	default:
		v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v132 == int32(0) {
			v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.Ui32(int32(2147483646)) <= base.Ui32(v137) {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			} else {
				v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v142 = v13 + int32(12)
				v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v142 == int32(0) {
				} else {
					v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v142))) = v146
				}
				v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
				v149 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				v153 = v137<<(uint(int32(1))%32) + int32(2)
				v154 = m.T0[v148].(func(*base.Module, int32, int32, int32, int32) int32)(m, v149, v140, v137, v153)
				mBase = m.M
				v155 = m.ExcPending
				if v155 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v154
					v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v153 - v157
					v160 = v154
					v161 = v157
					v165 = int32(192)
					*(*uint8)(unsafe.Add(mBase, uint32(v160+v161))) = uint8(v165)
					v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v167 + int32(1)
					v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v171 + int32(-1)
					v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v200 + int32(-16)
					m.G0 = v13 + int32(16)
					return
				}
			}
		} else {
			v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v160 = v136
			v161 = v135
			v165 = int32(192)
			*(*uint8)(unsafe.Add(mBase, uint32(v160+v161))) = uint8(v165)
			v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v167 + int32(1)
			v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v171 + int32(-1)
			v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v200 + int32(-16)
			m.G0 = v13 + int32(16)
			return
		}
	case 2:
		v93 = F_lua_tonumber(m, l0, int32(-1))
		mBase = m.M
		v94 = m.ExcPending
		if v94 != 0 {
			return
		} else {
			v95 = base.F64_abs(v93)
			if base.F64_eq(v95, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_mp_encode_double(m, l0, l1, v93)
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return
				} else {
					v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v200 + int32(-16)
					m.G0 = v13 + int32(16)
					return
				}
			} else {
				if base.F64_lt(v95, float64(9.223372036854776e+18)) == int32(0) {
					v104 = int64(-9223372036854775807 - 1)
				} else {
					v102 = base.I64_trunc_f64_s(v93)
					v104 = v102
				}
				if base.F64_ne(v93, base.F64_convert_i64_s(v104)) != 0 {
					F_mp_encode_double(m, l0, l1, v93)
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return
					} else {
						v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v200 + int32(-16)
						m.G0 = v13 + int32(16)
						return
					}
				} else {
					v108 = F_lua_tonumber(m, l0, int32(-1))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return
					} else {
						if base.F64_lt(base.F64_abs(v108), float64(9.223372036854776e+18)) == int32(0) {
							F_mp_encode_int(m, l0, l1, int64(-9223372036854775807-1))
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return
							} else {
								v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v200 + int32(-16)
								m.G0 = v13 + int32(16)
								return
							}
						} else {
							v115 = base.I64_trunc_f64_s(v108)
							F_mp_encode_int(m, l0, l1, v115)
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v200 + int32(-16)
								m.G0 = v13 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	case 3:
		v85 = F_lua_tolstring(m, l0, int32(-1), v13+int32(8))
		mBase = m.M
		v86 = m.ExcPending
		if v86 != 0 {
			return
		} else {
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			F_mp_encode_bytes(m, l0, l1, v85, v87)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return
			} else {
				v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v200 + int32(-16)
				m.G0 = v13 + int32(16)
				return
			}
		}
	case 4:
		v124 = F_table_is_an_array(m, l0)
		mBase = m.M
		v125 = m.ExcPending
		if v125 != 0 {
			return
		} else {
			if v124 == int32(0) {
				F_mp_encode_lua_table_as_map(m, l0, l1, l2)
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v200 + int32(-16)
					m.G0 = v13 + int32(16)
					return
				}
			} else {
				F_mp_encode_lua_table_as_array(m, l0, l1, l2)
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return
				} else {
					v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v200 + int32(-16)
					m.G0 = v13 + int32(16)
					return
				}
			}
		}
	}
}
