package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_isLuaInsecureAPIEnabled(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11 + int32(_a1941)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v11 + int32(_a1942)
	v18 = m.G84
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v20 = m.G98
	v27 = m.T0[v19].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v11+int32(_a1943), v11+int32(_a1944), v9+int32(16))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v32 = m.T0[v31].(func(*base.Module, int32) int32)(m, v27)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			if v32 != int32(1) {
				v52 = m.G98
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
				v54 = m.T0[v53].(func(*base.Module, int32) int32)(m, v27)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					if v54 != int32(3) {
						v100 = m.G3
						v106 = m.G8
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
						m.T0[v107].(func(*base.Module, int32, int32, int32))(m, v100+int32(_a1945), v100+int32(_a1946), int32(358))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int32(0)
						} else {
							m.Env.Exit(m, int32(1))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v58 = m.G99
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
						v60 = m.T0[v59].(func(*base.Module, int32) int32)(m, v27)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							if v60 != int32(2) {
								v100 = m.G3
								v106 = m.G8
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
								m.T0[v107].(func(*base.Module, int32, int32, int32))(m, v100+int32(_a1945), v100+int32(_a1946), int32(358))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int32(0)
								} else {
									m.Env.Exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v64 = m.G100
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
								v66 = m.G98
								v68 = m.T0[v65].(func(*base.Module, int32, int32) int32)(m, v27, int32(1))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
									v71 = m.T0[v70].(func(*base.Module, int32) int32)(m, v68)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										if v71 != 0 {
											v112 = m.G3
											v118 = m.G8
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
											m.T0[v119].(func(*base.Module, int32, int32, int32))(m, v112+int32(_a1947), v112+int32(_a1946), int32(360))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												m.Env.Exit(m, int32(1))
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v73 = int32(0)
											v75 = m.G101
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
											v77 = m.T0[v76].(func(*base.Module, int32, int32) int32)(m, v68, v73)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
												if v79 != int32(121) {
													v89 = v73
												} else {
													v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
													if v82 != int32(101) {
														v89 = v73
													} else {
														v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+2)))
														v89 = base.B2i32(v85 == int32(115))
													}
												}
												v92 = m.G86
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
												m.T0[v93].(func(*base.Module, int32))(m, v27)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return int32(0)
												} else {
													m.G0 = v9 + int32(32)
													return v89
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
				v36 = m.G10
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
				v38 = int32(0)
				v40 = m.G101
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = m.T0[v41].(func(*base.Module, int32, int32) int32)(m, v27, v38)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v42
					v45 = m.G3
					m.T0[v37].(func(*base.Module, int32, int32, int32, int32))(m, l0, v45+int32(_a1462), v45+int32(_a1948), v9)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v89 = v38
						v92 = m.G86
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
						m.T0[v93].(func(*base.Module, int32))(m, v27)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(32)
							return v89
						}
					}
				}
			}
		}
	}
}
func F_luaCallFunction(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
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
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int64
	_ = v652
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v729 int32
	_ = v729
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v773 int32
	_ = v773
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = l9
	*(*int64)(unsafe.Add(mBase, uint32(v14)+68)) = int64(8589934595)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = l0
	v22 = m.G3
	F_lua_pushstring(m, l3, v22+int32(_a1960))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v14 + int32(56)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v34 + int32(16)
	goto L3
L3:
	;
	F_lua_settable(m, l3, int32(-10000))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l8 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v67 = int32(0)
	F_lua_createtable(m, l3, l5, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L6:
	;
	v55 = m.G396
	v57 = int32(100000)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+64)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v57
	goto L16
L7:
	;
	v41 = m.G5
	v43 = v41 + int32(1209)
	v45 = int32(100000)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+64)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v45
	goto L9
L8:
	;
	goto L5
L9:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+68)) = v43
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v52 = int32(8)
	goto L14
L13:
	;
	v52 = int32(0)
	goto L14
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v52)
	goto L8
L15:
	;
	goto L5
L16:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+68)) = v55
	if v55 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v64 = int32(12)
	goto L21
L20:
	;
	v64 = int32(0)
	goto L21
L21:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v64)
	goto L15
L22:
	;
	if l5 < int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if l2 != 0 {
		goto L34
	} else {
		goto L35
	}
L24:
	;
	v82 = v67
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l4+v82<<(uint(int32(2))%32))))
	v92 = m.G7
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v94 = m.T0[v93].(func(*base.Module, int32, int32) int32)(m, v89, v14+int32(40))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L23
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	F_lua_pushlstring(m, l3, v94, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v101 = v82 + int32(1)
	F_lua_rawseti(m, l3, int32(-2), v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v101 != l5 {
		v82 = v101
		goto L25
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	v490 = m.G3
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v490)+uint32(_consts[1009])))
	v495 = v493 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v490)+uint32(_consts[1009]))) = v495
	if v495 != int32(50) {
		goto L122
	} else {
		goto L123
	}
L32:
	;
	v477 = F_lua_pcall(m, l3, int32(2), int32(1), int32(-4))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L121
	}
L33:
	;
	if l2 != 0 {
		goto L32
	} else {
		goto L119
	}
L34:
	;
	v409 = int32(0)
	F_lua_createtable(m, l3, l7, v409)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L111
	}
L35:
	;
	v116 = int32(0)
	goto L38
L36:
	;
	v176 = m.G3
	F_lua_setfield(m, l3, int32(-10002), v176+int32(_a2000))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L52
	}
L37:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = v116
	goto L36
L38:
	;
	goto L43
L43:
	;
	switch int32(0) {
	case 0:
		goto L46
	case 1:
		goto L47
	case 2:
		goto L48
	default:
		goto L45
	}
L45:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+7)))
	v158 = m.G398
	if base.Ui32(v157) < base.Ui32(int32(0)) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v171 = l3 + int32(72)
	goto L37
L47:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+88)) = v146
	v171 = l3 + int32(88)
	goto L37
L48:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v171 = v140 + int32(96)
	goto L37
L49:
	;
	v169 = v158
	goto L51
L50:
	;
	v169 = v156 + int32(8)
	goto L51
L51:
	;
	v171 = v169
	goto L37
L52:
	;
	goto L55
L53:
	;
	F_lua_createtable(m, l3, l7, int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L69
	}
L54:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = int32(1)
	goto L53
L55:
	;
	goto L60
L60:
	;
	switch int32(0) {
	case 0:
		goto L63
	case 1:
		goto L64
	case 2:
		goto L65
	default:
		goto L62
	}
L62:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+7)))
	v222 = m.G398
	if base.Ui32(v221) < base.Ui32(int32(0)) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v235 = l3 + int32(72)
	goto L54
L64:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+88)) = v210
	v235 = l3 + int32(88)
	goto L54
L65:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v235 = v204 + int32(96)
	goto L54
L66:
	;
	v233 = v222
	goto L68
L67:
	;
	v233 = v220 + int32(8)
	goto L68
L68:
	;
	v235 = v233
	goto L54
L69:
	;
	if l7 < int32(1) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	goto L80
L71:
	;
	v253 = v116
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = int32(0)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l6+v253<<(uint(int32(2))%32))))
	v263 = m.G7
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v265 = m.T0[v264].(func(*base.Module, int32, int32) int32)(m, v260, v14+int32(40))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L70
L74:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	F_lua_pushlstring(m, l3, v265, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v272 = v253 + int32(1)
	F_lua_rawseti(m, l3, int32(-2), v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v272 != l7 {
		v253 = v272
		goto L72
	} else {
		goto L77
	}
L77:
	;
	goto L73
L78:
	;
	v346 = m.G3
	F_lua_setfield(m, l3, int32(-10002), v346+int32(_a2001))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L94
	}
L79:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	*(*int32)(unsafe.Add(mBase, uint32(v343)+8)) = int32(0)
	goto L78
L80:
	;
	goto L85
L85:
	;
	switch int32(0) {
	case 0:
		goto L88
	case 1:
		goto L89
	case 2:
		goto L90
	default:
		goto L87
	}
L87:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+7)))
	v328 = m.G398
	if base.Ui32(v327) < base.Ui32(int32(0)) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v341 = l3 + int32(72)
	goto L79
L89:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+88)) = v316
	v341 = l3 + int32(88)
	goto L79
L90:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v341 = v310 + int32(96)
	goto L79
L91:
	;
	v339 = v328
	goto L93
L92:
	;
	v339 = v326 + int32(8)
	goto L93
L93:
	;
	v341 = v339
	goto L79
L94:
	;
	goto L97
L95:
	;
	goto L33
L96:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	*(*int32)(unsafe.Add(mBase, uint32(v407)+8)) = int32(1)
	goto L95
L97:
	;
	goto L102
L102:
	;
	switch int32(0) {
	case 0:
		goto L105
	case 1:
		goto L106
	case 2:
		goto L107
	default:
		goto L104
	}
L104:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+7)))
	v392 = m.G398
	if base.Ui32(v391) < base.Ui32(int32(0)) {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	v405 = l3 + int32(72)
	goto L96
L106:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+88)) = v380
	v405 = l3 + int32(88)
	goto L96
L107:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v405 = v374 + int32(96)
	goto L96
L108:
	;
	v403 = v392
	goto L110
L109:
	;
	v403 = v390 + int32(8)
	goto L110
L110:
	;
	v405 = v403
	goto L96
L111:
	;
	if l7 < int32(1) {
		goto L32
	} else {
		goto L112
	}
L112:
	;
	v424 = v409
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = int32(0)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l6+v424<<(uint(int32(2))%32))))
	v434 = m.G7
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	v436 = m.T0[v435].(func(*base.Module, int32, int32) int32)(m, v431, v14+int32(40))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L115
	}
L114:
	;
	goto L33
L115:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	F_lua_pushlstring(m, l3, v436, v438)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v443 = v424 + int32(1)
	F_lua_rawseti(m, l3, int32(-2), v443)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	if v443 != l7 {
		v424 = v443
		goto L113
	} else {
		goto L118
	}
L118:
	;
	goto L114
L119:
	;
	v461 = F_lua_pcall(m, l3, int32(0), int32(1), int32(-2))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v488 = v461
	goto L31
L121:
	;
	v488 = v477
	goto L31
L122:
	;
	if v488 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v501 = F_lua_gc(m, l3, int32(5), int32(50))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v503 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v503)+uint32(_consts[1009]))) = int32(0)
	goto L122
L125:
	;
	v747 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+64)) = v747
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v747
	goto L192
L126:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	F_luaReplyToServerReply(m, l0, v739, l3)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L189
	}
L127:
	;
	goto L132
L128:
	;
	goto L183
L129:
	;
	v651 = v14 + int32(48)
	v652 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v651))) = v652
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v652
	F_luaExtractErrorInformation(m, l3, v14+int32(40))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L166
	}
L130:
	;
	if v567 == int32(5) {
		goto L129
	} else {
		goto L145
	}
L131:
	;
	v561 = m.G398
	if v527 != v561 {
		goto L143
	} else {
		goto L144
	}
L132:
	;
	goto L136
L136:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v527 = v524 + int32(-16)
	goto L131
L143:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v527)+8))
	v567 = v564
	goto L130
L144:
	;
	v567 = int32(-1)
	goto L130
L145:
	;
	v570 = m.G3
	goto L150
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v639
	v641 = m.G3
	v646 = m.G40
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v646)))
	v648 = m.T0[v647].(func(*base.Module, int32, int32, int32) int32)(m, l0, v641+int32(_a2002), v14+int32(32))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L165
	}
L147:
	;
	v637 = F_lua_tolstring(m, l3, int32(-1), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L164
	}
L148:
	;
	if v632 != 0 {
		goto L147
	} else {
		goto L163
	}
L149:
	;
	v622 = m.G398
	if v588 != v622 {
		goto L161
	} else {
		goto L162
	}
L150:
	;
	goto L154
L154:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v588 = v585 + int32(-16)
	goto L149
L161:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v588)+8))
	v632 = base.B2i32(base.Ui32(v625+int32(-3)) < base.Ui32(int32(2)))
	goto L148
L162:
	;
	v632 = int32(0)
	goto L148
L163:
	;
	v639 = v570 + int32(_a2003)
	goto L146
L164:
	;
	v639 = v637
	goto L146
L165:
	;
	goto L128
L166:
	;
	v660 = m.G41
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v660)))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v663 = int32(0)
	v664 = base.B2i32(v662 == v663)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v651)))
	if v667 == v663 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	if v665 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v665
	v681 = m.G3
	v686 = m.T0[v661].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v664, v681+int32(_a79), v14+int32(16))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L172
	}
L169:
	;
	if v666 == int32(0) {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v667
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v666
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v665
	v675 = m.G3
	v678 = m.T0[v661].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v664, v675+int32(_a2004), v14)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	goto L167
L172:
	;
	goto L167
L173:
	;
	if v666 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	v690 = m.G11
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v690)))
	m.T0[v691].(func(*base.Module, int32))(m, v665)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	if v667 == int32(0) {
		goto L128
	} else {
		goto L179
	}
L177:
	;
	v696 = m.G11
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v696)))
	m.T0[v697].(func(*base.Module, int32))(m, v666)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	v702 = m.G11
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)))
	m.T0[v703].(func(*base.Module, int32))(m, v667)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	goto L128
L181:
	;
	goto L125
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v729 + int32(-16)
	goto L181
L183:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	goto L182
L189:
	;
	goto L125
L190:
	;
	v759 = m.G3
	F_lua_pushstring(m, l3, v759+int32(_a1960))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L197
	}
L192:
	;
	goto L193
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+68)) = v747
	v755 = int32(0)
	goto L195
L195:
	;
	goto L196
L196:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v755)
	goto L190
L197:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v765)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v765 + int32(16)
	goto L198
L198:
	;
	F_lua_settable(m, l3, int32(-10000))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	m.G0 = v14 + int32(80)
	return
}
func F_luaEngineFunctionMemoryOverhead(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8 = m.G332
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = m.T0[v9].(func(*base.Module, int32) int32)(m, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(0)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v16 == v14 {
			v23 = v14
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			if v24 == int32(0) {
				v31 = v14
				v34 = m.G332
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
				v36 = m.T0[v35].(func(*base.Module, int32) int32)(m, l1)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					return v23 + v10 + v31 + v36
				}
			} else {
				v27 = m.G332
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v29 = m.T0[v28].(func(*base.Module, int32) int32)(m, v24)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = v29
					v34 = m.G332
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					v36 = m.T0[v35].(func(*base.Module, int32) int32)(m, l1)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						return v23 + v10 + v31 + v36
					}
				}
			}
		} else {
			v19 = m.G332
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v21 = m.T0[v20].(func(*base.Module, int32) int32)(m, v16)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = v21
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				if v24 == int32(0) {
					v31 = v14
					v34 = m.G332
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					v36 = m.T0[v35].(func(*base.Module, int32) int32)(m, l1)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						return v23 + v10 + v31 + v36
					}
				} else {
					v27 = m.G332
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					v29 = m.T0[v28].(func(*base.Module, int32) int32)(m, v24)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = v29
						v34 = m.G332
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						v36 = m.T0[v35].(func(*base.Module, int32) int32)(m, l1)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							return v23 + v10 + v31 + v36
						}
					}
				}
			}
		}
	}
}
func F_luaEngineLoadHook(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	v6 = m.G3
	v9 = F_luaGetFromRegistry(m, l0, v6+int32(_a1955))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 == int32(0) {
			v45 = m.G3
			v51 = m.G8
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
			m.T0[v52].(func(*base.Module, int32, int32, int32))(m, v45+int32(_a1958), v45+int32(_a1952), int32(93))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				m.Env.Exit(m, int32(1))
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
			v14 = m.G385
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v16 = m.T0[v15].(func(*base.Module) int64)(m)
			mBase = m.M
			v19 = base.I64_div_u_s(v16-v13, int64(1000))
			v20 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)))
			if v20 == int64(0) {
				return
			} else {
				if base.Ui64(v19) <= base.Ui64(v20) {
					return
				} else {
					v24 = m.G5
					v26 = v24 + int32(1183)
					v28 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v28
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v28
					*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v26
					if v26 != 0 {
						v35 = int32(4)
					} else {
						v35 = int32(0)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)) = uint8(v35)
					v38 = m.G3
					F_luaPushError(m, l0, v38+int32(_a1959))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v43 = F_luaError(m, l0)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
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
func F_luaEngineResetEnv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	if base.Ui32(int32(2)) <= base.Ui32(l2) {
		v47 = m.G3
		v53 = m.G8
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
		m.T0[v54].(func(*base.Module, int32, int32, int32))(m, v47+int32(_a1313), v47+int32(_a1946), int32(373))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			m.Env.Exit(m, int32(1))
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v9 = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1+base.B2i32(l2 != v9)<<(uint(int32(2))%32))))
		if v14 == v9 {
			v59 = m.G3
			v65 = m.G8
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			m.T0[v66].(func(*base.Module, int32, int32, int32))(m, v59+int32(_a775), v59+int32(_a1946), int32(375))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				m.Env.Exit(m, int32(1))
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if l3 == int32(0) {
				v32 = int32(0)
				v35 = F_lua_gc(m, v14, int32(2), v32)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					F_lua_close(m, v14)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = v32
						v41 = F_isLuaInsecureAPIEnabled(m, l0)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v41
							F_initializeLuaState(m, l1, l2)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								return v39
							}
						}
					}
				}
			} else {
				v19 = m.G5
				v22 = m.G9
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v24 = m.T0[v23].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(8))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v19 + int32(1172)
					*(*int32)(unsafe.Add(mBase, uint32(v24))) = v14
					v39 = v24
					v41 = F_isLuaInsecureAPIEnabled(m, l0)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v41
						F_initializeLuaState(m, l1, l2)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							return v39
						}
					}
				}
			}
		}
	}
}
func F_luaFunctionLibraryCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
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
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v117 int32
	_ = v117
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v247 int32
	_ = v247
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v352 int32
	_ = v352
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
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
	var v429 int64
	_ = v429
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v466 int64
	_ = v466
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v503 int32
	_ = v503
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v622 int32
	_ = v622
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v764 int32
	_ = v764
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v892 int32
	_ = v892
	var v902 int32
	_ = v902
	var v910 int32
	_ = v910
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	v10 = m.G0
	v12 = v10 - int32(64)
	m.G0 = v12
	goto L3
L1:
	;
	v99 = int32(0)
	goto L25
L2:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	switch v70 + int32(-5) {
	case 0:
		goto L20
	default:
		goto L18
	case 2:
		goto L19
	}
L3:
	;
	goto L8
L8:
	;
	switch int32(0) {
	case 0:
		goto L11
	case 1:
		goto L12
	case 2:
		goto L13
	default:
		goto L10
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+7)))
	v55 = m.G398
	if base.Ui32(v54) < base.Ui32(int32(0)) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v67 = l0 + int32(72)
	goto L2
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v43
	v67 = l0 + int32(88)
	goto L2
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v67 = v37 + int32(96)
	goto L2
L14:
	;
	v66 = v55
	goto L16
L15:
	;
	v66 = v53 + int32(8)
	goto L16
L16:
	;
	v67 = v66
	goto L2
L17:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v86 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v85 = v79 + v70<<(uint(int32(2))%32) + int32(152)
	goto L17
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v85 = v76 + int32(8)
	goto L17
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v85 = v73 + int32(16)
	goto L17
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v86
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(16)
	goto L1
L22:
	;
	goto L1
L23:
	;
	v159 = m.G3
	F_lua_getfield(m, l0, int32(-10000), v159+int32(_a1949))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L39
	} else {
		goto L40
	}
L24:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v117+int32(-16))))
	*(*int32)(unsafe.Add(mBase, uint32(v156)+8)) = v99
	goto L23
L25:
	;
	goto L31
L31:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L24
L39:
	;
	return int32(0)
L40:
	;
	F_lua_setfield(m, l0, int32(-2), v159+int32(_a1950))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	goto L44
L42:
	;
	goto L60
L43:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+8)) = int32(1)
	goto L42
L44:
	;
	goto L49
L49:
	;
	switch int32(0) {
	case 0:
		goto L52
	case 1:
		goto L53
	case 2:
		goto L54
	default:
		goto L51
	}
L51:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+7)))
	v212 = m.G398
	if base.Ui32(v211) < base.Ui32(int32(0)) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v225 = l0 + int32(72)
	goto L43
L53:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v200
	v225 = l0 + int32(88)
	goto L43
L54:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v225 = v194 + int32(96)
	goto L43
L55:
	;
	v223 = v212
	goto L57
L56:
	;
	v223 = v210 + int32(8)
	goto L57
L57:
	;
	v225 = v223
	goto L43
L58:
	;
	if l1&int32(3) == int32(0) {
		v278 = l1
		goto L71
	} else {
		goto L72
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v247 + int32(-16)
	goto L58
L60:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L59
L66:
	;
	v923 = m.G3
	v929 = m.G8
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v929)))
	m.T0[v930].(func(*base.Module, int32, int32, int32))(m, v923+int32(_a1951), v923+int32(_a1952), int32(145))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L39
	} else {
		goto L245
	}
L67:
	;
	goto L175
L68:
	;
	goto L99
L69:
	;
	v314 = F_luaL_loadbuffer(m, l0, l1, v311, v159+int32(_a1953))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L39
	} else {
		goto L85
	}
L70:
	;
	v311 = v303 - l1
	goto L69
L71:
	;
	v282 = v278
	goto L79
L72:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v264 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v267 = l1
	goto L75
L74:
	;
	v311 = l1 - l1
	goto L69
L75:
	;
	v271 = v267 + int32(1)
	if v271&int32(3) == int32(0) {
		v278 = v271
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if v276 != 0 {
		v267 = v271
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v303 = v271
	goto L70
L79:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	v291 = int32(-2139062144)
	if (int32(16843008)-v288|v288)&v291 == v291 {
		v282 = v282 + int32(4)
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v297 = v282
	goto L82
L81:
	;
	goto L80
L82:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	if v301 != 0 {
		v297 = v297 + int32(1)
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v303 = v297
	goto L70
L84:
	;
	goto L83
L85:
	;
	if v314 == int32(0) {
		goto L68
	} else {
		goto L86
	}
L86:
	;
	v318 = m.G15
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	v322 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L39
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v322
	v326 = m.G3
	v331 = m.T0[v319].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v326+int32(_a1954), v12+int32(16))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L39
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v331
	goto L91
L89:
	;
	v659 = v99
	goto L67
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v352 + int32(-16)
	goto L89
L91:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L90
L97:
	;
	if v419 != int32(6) {
		goto L66
	} else {
		goto L112
	}
L98:
	;
	v413 = m.G398
	if v379 != v413 {
		goto L110
	} else {
		goto L111
	}
L99:
	;
	goto L103
L103:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v379 = v376 + int32(-16)
	goto L98
L110:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v379)+8))
	v419 = v416
	goto L97
L111:
	;
	v419 = int32(-1)
	goto L97
L112:
	;
	v422 = F_list_create(m)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L39
	} else {
		goto L113
	}
L113:
	;
	v424 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v422
	v427 = m.G385
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	v429 = m.T0[v428].(func(*base.Module) int64)(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v429
	v434 = m.G3
	F_luaSaveOnRegistry(m, l0, v434+int32(_a1955), v12+int32(40))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L39
	} else {
		goto L114
	}
L114:
	;
	v441 = m.G5
	v443 = v441 + int32(1183)
	v445 = int32(100000)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v445
	goto L116
L115:
	;
	v455 = int32(0)
	v458 = F_lua_pcall(m, l0, v455, v455, v455)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L39
	} else {
		goto L123
	}
L116:
	;
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v443
	if v443 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v452 = int32(8)
	goto L121
L120:
	;
	v452 = int32(0)
	goto L121
L121:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)) = uint8(v452)
	goto L115
L122:
	;
	v590 = m.G9
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v592)+4))
	goto L157
L123:
	;
	if v458 == int32(0) {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v466 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(32)))) = v466
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v466
	F_luaExtractErrorInformation(m, l0, v12+int32(24))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L39
	} else {
		goto L125
	}
L125:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v474
	v477 = m.G3
	v480 = m.G15
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	v482 = m.T0[v481].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v477+int32(_a1956), v12)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L39
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v482
	goto L129
L127:
	;
	F_luaErrorInformationDiscard(m, v12+int32(24))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L39
	} else {
		goto L135
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v503 + int32(-16)
	goto L127
L129:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L128
L135:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v518 = F_list_get_iter(m, v517)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L39
	} else {
		goto L137
	}
L136:
	;
	F_list_release_iter(m, v518)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L39
	} else {
		goto L155
	}
L137:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	if v522 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if v528 == int32(0) {
		goto L136
	} else {
		goto L141
	}
L139:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v518))) = v525
	v528 = v524
	goto L138
L140:
	;
	v528 = int32(0)
	goto L138
L141:
	;
	v534 = v528
	goto L142
L142:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	v542 = m.G17
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)))
	m.T0[v543].(func(*base.Module, int32, int32))(m, int32(0), v541)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L39
	} else {
		goto L144
	}
L143:
	;
	goto L136
L144:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v534)+16))
	if v546 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v534)+12))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)+4))
	F_luaL_unref(m, l0, int32(-10000), v556)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L39
	} else {
		goto L148
	}
L146:
	;
	v550 = m.G17
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	m.T0[v551].(func(*base.Module, int32, int32))(m, int32(0), v546)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L39
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v559 = m.G11
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	m.T0[v560].(func(*base.Module, int32))(m, v555)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L39
	} else {
		goto L149
	}
L149:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	m.T0[v563].(func(*base.Module, int32))(m, v534)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L39
	} else {
		goto L150
	}
L150:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	if v568 != 0 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	if v574 != 0 {
		v534 = v574
		goto L142
	} else {
		goto L154
	}
L152:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v568)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v518))) = v571
	v574 = v570
	goto L151
L153:
	;
	v574 = int32(0)
	goto L151
L154:
	;
	goto L143
L155:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	F_list_destroy(m, v586)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L39
	} else {
		goto L156
	}
L156:
	;
	v659 = int32(0)
	goto L67
L157:
	;
	v595 = m.T0[v591].(func(*base.Module, int32, int32) int32)(m, v593, int32(4))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L39
	} else {
		goto L158
	}
L158:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v598 = F_list_get_iter(m, v597)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L39
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v598)))
	if v604 != 0 {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	F_list_release_iter(m, v598)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L39
	} else {
		goto L171
	}
L161:
	;
	if v610 == int32(0) {
		goto L160
	} else {
		goto L164
	}
L162:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v604)))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v604)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v598))) = v607
	v610 = v606
	goto L161
L163:
	;
	v610 = int32(0)
	goto L161
L164:
	;
	v614 = v610
	goto L165
L165:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v595+v622<<(uint(int32(2))%32)))) = v614
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v622 + int32(1)
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v598)))
	if v632 != 0 {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	goto L160
L167:
	;
	if v638 != 0 {
		v614 = v638
		goto L165
	} else {
		goto L170
	}
L168:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v632)))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v598))) = v635
	v638 = v634
	goto L167
L169:
	;
	v638 = int32(0)
	goto L167
L170:
	;
	goto L166
L171:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	F_list_destroy(m, v650)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L39
	} else {
		goto L172
	}
L172:
	;
	v659 = v595
	goto L67
L173:
	;
	goto L197
L174:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v715)+8))
	switch v718 + int32(-5) {
	case 0:
		goto L192
	default:
		goto L190
	case 2:
		goto L191
	}
L175:
	;
	goto L180
L180:
	;
	switch int32(0) {
	case 0:
		goto L183
	case 1:
		goto L184
	case 2:
		goto L185
	default:
		goto L182
	}
L182:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v699)+4))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v700)))
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+7)))
	v703 = m.G398
	if base.Ui32(v702) < base.Ui32(int32(0)) {
		goto L186
	} else {
		goto L187
	}
L183:
	;
	v715 = l0 + int32(72)
	goto L174
L184:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v688)+4))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v689)))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v690)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v691
	v715 = l0 + int32(88)
	goto L174
L185:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v715 = v685 + int32(96)
	goto L174
L186:
	;
	v714 = v703
	goto L188
L187:
	;
	v714 = v701 + int32(8)
	goto L188
L188:
	;
	v715 = v714
	goto L174
L189:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)))
	if v734 != 0 {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v733 = v727 + v718<<(uint(int32(2))%32) + int32(152)
	goto L189
L191:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v733 = v724 + int32(8)
	goto L189
L192:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v733 = v721 + int32(16)
	goto L189
L193:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+8)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v736))) = v734
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v740 + int32(16)
	goto L173
L194:
	;
	goto L173
L195:
	;
	v806 = m.G3
	F_lua_getfield(m, l0, int32(-10000), v806+int32(_a1957))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L39
	} else {
		goto L211
	}
L196:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v764+int32(-16))))
	*(*int32)(unsafe.Add(mBase, uint32(v803)+8)) = int32(0)
	goto L195
L197:
	;
	goto L203
L203:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L196
L211:
	;
	F_lua_setfield(m, l0, int32(-2), v806+int32(_a1950))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L39
	} else {
		goto L212
	}
L212:
	;
	goto L215
L213:
	;
	goto L231
L214:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v870)))
	*(*int32)(unsafe.Add(mBase, uint32(v872)+8)) = int32(1)
	goto L213
L215:
	;
	goto L220
L220:
	;
	switch int32(0) {
	case 0:
		goto L223
	case 1:
		goto L224
	case 2:
		goto L225
	default:
		goto L222
	}
L222:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v853)+4))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v854)))
	v856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v855)+7)))
	v857 = m.G398
	if base.Ui32(v856) < base.Ui32(int32(0)) {
		goto L226
	} else {
		goto L227
	}
L223:
	;
	v870 = l0 + int32(72)
	goto L214
L224:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v842)+4))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v843)))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v844)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v845
	v870 = l0 + int32(88)
	goto L214
L225:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v870 = v839 + int32(96)
	goto L214
L226:
	;
	v868 = v857
	goto L228
L227:
	;
	v868 = v855 + int32(8)
	goto L228
L228:
	;
	v870 = v868
	goto L214
L229:
	;
	v902 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v902
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v902
	goto L239
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v892 + int32(-16)
	goto L229
L231:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L230
L237:
	;
	F_luaSaveOnRegistry(m, l0, v806+int32(_a1955), int32(0))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L39
	} else {
		goto L244
	}
L239:
	;
	goto L240
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v902
	v910 = int32(0)
	goto L242
L242:
	;
	goto L243
L243:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)) = uint8(v910)
	goto L237
L244:
	;
	m.G0 = v12 + int32(64)
	return v659
L245:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_luaLdbLineHook(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
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
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
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
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = m.G3
	v13 = F_luaGetFromRegistry(m, l0, v10+int32(_a1960))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v202 = m.G3
	v208 = m.G8
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	m.T0[v209].(func(*base.Module, int32, int32, int32))(m, v202+int32(_a1966), v202+int32(_a1967), int32(1997))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L2
	} else {
		goto L67
	}
L2:
	;
	return
L3:
	;
	if v13 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	goto L7
L5:
	;
	v69 = m.G3
	v72 = F_lua_getinfo(m, l0, v69+int32(_a1994), l1)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L18
	}
L6:
	;
	goto L5
L7:
	;
	goto L16
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v57
	goto L6
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v21) <= base.Ui32(v53) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v57 = base.I32_div_s(v21-v53, int32(24))
	goto L15
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v75 = m.G6
	*(*int32)(unsafe.Add(mBase, uint32(v75)+280)) = v74
	goto L19
L19:
	;
	v82 = m.G6
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+260))
	if v83 < int32(1) {
		v101 = int32(0)
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v111 = l1 + int32(36)
	v113 = v69 + int32(_a1995)
	v116 = int32(*(*int8)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[1004]))))
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v105 = m.G6
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+268))
	v109 = v101 | base.B2i32(v106 != int32(0))
	goto L20
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+280))
	v89 = int32(0)
	goto L23
L23:
	;
	v92 = m.G6
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v89<<(uint(int32(2))%32))+4))
	v97 = base.B2i32(v96 == v86)
	if v96 == v86 {
		v101 = v97
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v101 = v97
	goto L21
L25:
	;
	v99 = v89 + int32(1)
	if v99 != v83 {
		v89 = v99
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	m.G0 = v8 + int32(16)
	return
L28:
	;
	if v141 == int32(0) {
		goto L27
	} else {
		goto L44
	}
L29:
	;
	v117 = int32(0)
	v118 = F_strchr(m, v111, v116)
	mBase = m.M
	if v118 == v117 {
		v138 = v117
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v141 = v111
	goto L28
L31:
	;
	v141 = v138
	goto L28
L32:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[1005]))))
	if v121 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	if v122 == int32(0) {
		v138 = v117
		goto L31
	} else {
		goto L35
	}
L34:
	;
	v141 = v118
	goto L28
L35:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[1006]))))
	if v125 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
	if v127 == int32(0) {
		v138 = v117
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v126 = F_twobyte_strstr(m, v118, v113)
	mBase = m.M
	v141 = v126
	goto L28
L38:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[1007]))))
	if v130 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
	if v132 == int32(0) {
		v138 = v117
		goto L31
	} else {
		goto L41
	}
L40:
	;
	v131 = F_threebyte_strstr(m, v118, v113)
	mBase = m.M
	v141 = v131
	goto L28
L41:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[1008]))))
	if v135 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v137 = F_twoway_strstr(m, v118, v113)
	mBase = m.M
	v138 = v137
	goto L31
L43:
	;
	v136 = F_fourbyte_strstr(m, v118, v113)
	mBase = m.M
	v141 = v136
	goto L28
L44:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v144 != int32(3) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v152 = m.G6
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+264))
	goto L49
L46:
	;
	v147 = m.G6
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+264))
	goto L47
L47:
	;
	if v148|v109 == int32(0) {
		goto L27
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	if v153|v109 == int32(0) {
		goto L27
	} else {
		goto L50
	}
L50:
	;
	v157 = m.G3
	if v109 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v169 = m.G6
	*(*int32)(unsafe.Add(mBase, uint32(v169)+264)) = int32(0)
	goto L58
L52:
	;
	v160 = m.G6
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+268))
	goto L54
L53:
	;
	v167 = v157 + int32(_a1996)
	goto L51
L54:
	;
	v162 = m.G3
	if v161 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v165 = int32(_a1997)
	goto L57
L56:
	;
	v165 = int32(_a1998)
	goto L57
L57:
	;
	v167 = v162 + v165
	goto L51
L58:
	;
	v172 = m.G6
	*(*int32)(unsafe.Add(mBase, uint32(v172)+268)) = int32(0)
	goto L59
L59:
	;
	v174 = m.G15
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v176 = m.G6
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+280))
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v177
	v181 = m.G3
	v184 = m.T0[v175].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v181+int32(_a1999), v8)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	F_ldbLog(m, v184)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v188 = m.G6
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+280))
	goto L63
L63:
	;
	F_ldbLogSourceLine(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	F_ldbSendLogs(m)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	v194 = F_ldbRepl(m, l0)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	goto L27
L67:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_luaLogCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = m.G3
	v18 = F_luaGetFromRegistry(m, l0, v15+int32(_a1960))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return v165
L2:
	;
	v158 = m.G3
	F_luaPushErrorBuff(m, l0, v158+int32(_a1965))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L40
	}
L3:
	;
	v146 = m.G3
	v152 = m.G8
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	m.T0[v153].(func(*base.Module, int32, int32, int32))(m, v146+int32(_a1966), v146+int32(_a1967), int32(1425))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L39
	}
L4:
	;
	return int32(0)
L5:
	;
	if v18 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = (v24 - v25) >> (uint(int32(4)) % 32)
	goto L8
L7:
	;
	v39 = int32(0) - v28
	v40 = F_lua_isnumber(m, l0, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L13
	}
L8:
	;
	if int32(1) < v28 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v31 = m.G3
	F_luaPushErrorBuff(m, l0, v31+int32(_a1968))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v36 = F_lua_error(m, l0)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v165 = v36
	goto L1
L12:
	;
	v49 = F_lua_tonumber(m, l0, v39)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L19
	}
L13:
	;
	if v40 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v42 = m.G3
	F_luaPushErrorBuff(m, l0, v42+int32(_a1969))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v47 = F_lua_error(m, l0)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v165 = v47
	goto L1
L17:
	;
	if base.Ui32(int32(3)) < base.Ui32(v58) {
		goto L2
	} else {
		goto L21
	}
L18:
	;
	v58 = int32(-2147483648)
	goto L17
L19:
	;
	if base.F64_lt(base.F64_abs(v49), float64(2.147483648e+09)) == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v56 = base.I32_trunc_f64_s(v49)
	v58 = v56
	goto L17
L21:
	;
	v65 = F_lua_tolstring(m, l0, int32(1)-v28, v13+int32(44))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	v77 = int32(2)
	if v28 == v77 {
		v123 = v76
		goto L27
	} else {
		goto L28
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v65
	v69 = m.G3
	v74 = F_lm_asprintf(m, v69+int32(_a79), v13+int32(32))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	if v65 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v76 = int32(0)
	goto L22
L26:
	;
	v76 = v74
	goto L22
L27:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v123
	v128 = m.G3
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v128+int32(_a1970)+v58<<(uint(int32(2))%32))))
	v137 = m.G10
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	m.T0[v138].(func(*base.Module, int32, int32, int32, int32))(m, v126, v134, v128+int32(_a79), v13)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L37
	}
L28:
	;
	v84 = v77
	v87 = v76
	goto L29
L29:
	;
	v93 = F_lua_tolstring(m, l0, v84-v28, v13+int32(44))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	v123 = v110
	goto L27
L31:
	;
	v114 = v84 + int32(1)
	if v114 != v28 {
		v84 = v114
		v87 = v110
		goto L29
	} else {
		goto L36
	}
L32:
	;
	if v93 == int32(0) {
		v110 = v87
		goto L31
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v87
	v99 = m.G3
	v100 = m.G11
	v105 = F_lm_asprintf(m, v99+int32(_a1412), v13+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	m.T0[v107].(func(*base.Module, int32))(m, v87)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v110 = v105
	goto L31
L36:
	;
	goto L30
L37:
	;
	v141 = m.G11
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	m.T0[v142].(func(*base.Module, int32))(m, v123)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v165 = int32(0)
	goto L1
L39:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v163 = F_lua_error(m, l0)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v165 = v163
	goto L1
}
func F_luaMaskCountHook(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = m.G3
	v13 = F_luaGetFromRegistry(m, l0, v10+int32(_a1960))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 == int32(0) {
			v63 = m.G3
			v69 = m.G8
			v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
			m.T0[v70].(func(*base.Module, int32, int32, int32))(m, v63+int32(_a1966), v63+int32(_a1967), int32(1924))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				m.Env.Exit(m, int32(1))
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v18 = m.G381
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			v20 = m.T0[v19].(func(*base.Module, int32) int32)(m, v17)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				if v20 != int32(1) {
					m.G0 = v8 + int32(16)
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					v25 = m.G3
					if v24 != 0 {
						v28 = int32(_a2005)
					} else {
						v28 = int32(_a2006)
					}
					v29 = v25 + v28
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v29
					v36 = m.G10
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					m.T0[v37].(func(*base.Module, int32, int32, int32, int32))(m, int32(0), v25+int32(_a1461), v25+int32(_a79), v8)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v40 = m.G5
						v42 = v40 + int32(1209)
						v44 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v44
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v44
						*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v42
						if v42 != 0 {
							v51 = int32(4)
						} else {
							v51 = int32(0)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)) = uint8(v51)
						F_luaPushErrorBuff(m, l0, v29)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v56 = F_lua_error(m, l0)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
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
}
func F_luaMemory(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_lua_gc(m, l0, int32(3), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4 << (uint(int32(10)) % 32)
	}
}
func F_luaProtectedTableError(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = m.G3
	v12 = F_luaGetFromRegistry(m, l0, v9+int32(_a1960))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if (v16-v17)>>(uint(int32(4))%32) == int32(2) {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v57 = v54 + int32(-16)
			v91 = m.G398
			if v57 != v91 {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
				v101 = base.B2i32(base.Ui32(v94+int32(-3)) < base.Ui32(int32(2)))
			} else {
				v101 = int32(0)
			}
			if v101 != 0 {
				v113 = F_lua_tolstring(m, l0, int32(-1), int32(0))
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v113
					v116 = m.G3
					v119 = F_luaL_error(m, l0, v116+int32(_a1961), v7)
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return int32(0)
					}
				}
			} else {
				v103 = F_lua_isnumber(m, l0, int32(-1))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return int32(0)
				} else {
					if v103 != 0 {
						v113 = F_lua_tolstring(m, l0, int32(-1), int32(0))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v113
							v116 = m.G3
							v119 = F_luaL_error(m, l0, v116+int32(_a1961), v7)
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(16)
								return int32(0)
							}
						}
					} else {
						v105 = m.G3
						v109 = F_luaL_error(m, l0, v105+int32(_a1962), int32(0))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int32(0)
						} else {
							v113 = F_lua_tolstring(m, l0, int32(-1), int32(0))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v113
								v116 = m.G3
								v119 = F_luaL_error(m, l0, v116+int32(_a1961), v7)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(16)
									return int32(0)
								}
							}
						}
					}
				}
			}
		} else {
			v23 = m.G3
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v30 = m.G10
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
			m.T0[v31].(func(*base.Module, int32, int32, int32, int32))(m, v24, v23+int32(_a1462), v23+int32(_a1963), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v37 = F_luaL_error(m, l0, v23+int32(_a1964), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v57 = v54 + int32(-16)
					v91 = m.G398
					if v57 != v91 {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
						v101 = base.B2i32(base.Ui32(v94+int32(-3)) < base.Ui32(int32(2)))
					} else {
						v101 = int32(0)
					}
					if v101 != 0 {
						v113 = F_lua_tolstring(m, l0, int32(-1), int32(0))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v113
							v116 = m.G3
							v119 = F_luaL_error(m, l0, v116+int32(_a1961), v7)
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(16)
								return int32(0)
							}
						}
					} else {
						v103 = F_lua_isnumber(m, l0, int32(-1))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							if v103 != 0 {
								v113 = F_lua_tolstring(m, l0, int32(-1), int32(0))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v113
									v116 = m.G3
									v119 = F_luaL_error(m, l0, v116+int32(_a1961), v7)
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										m.G0 = v7 + int32(16)
										return int32(0)
									}
								}
							} else {
								v105 = m.G3
								v109 = F_luaL_error(m, l0, v105+int32(_a1962), int32(0))
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return int32(0)
								} else {
									v113 = F_lua_tolstring(m, l0, int32(-1), int32(0))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v113
										v116 = m.G3
										v119 = F_luaL_error(m, l0, v116+int32(_a1961), v7)
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											m.G0 = v7 + int32(16)
											return int32(0)
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
func F_luaPushError(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	F_luaPushErrorBuff(m, l0, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_luaRegisterServerAPI(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v114 int32
	_ = v114
	var v130 int64
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
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
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	switch int32(0) {
	case 0:
		v65 = l1 + int32(72)
	case 1:
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(5)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v41
		v65 = l1 + int32(88)
	case 2:
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v65 = v35 + int32(96)
	default:
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
		v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+7)))
		v53 = m.G398
		if base.Ui32(v52) < base.Ui32(int32(0)) {
			v64 = v53
		} else {
			v64 = v51 + int32(8)
		}
		v65 = v64
	}
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v73 + int32(16)
	v77 = int32(0)
	F_lua_createtable(m, l1, v77, v77)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		return
	} else {
		v81 = m.G5
		F_lua_pushcclosure(m, l1, v81+int32(1187), int32(0))
		mBase = m.M
		v86 = m.ExcPending
		if v86 != 0 {
			return
		} else {
			v88 = m.G3
			F_lua_setfield(m, l1, int32(-2), v88+int32(_a1971))
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return
			} else {
				v94 = F_lua_setmetatable(m, l1, int32(-2))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return
				} else {
					v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v114 + int32(-16)
					*(*int32)(unsafe.Add(mBase, uint32(v10+int32(24)))) = int32(0)
					v130 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v10+int32(16)))) = v130
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v130
					v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v134
					v137 = v88 + int32(_a1960)
					F_lua_pushstring(m, l1, v137)
					mBase = m.M
					v139 = m.ExcPending
					if v139 != 0 {
						return
					} else {
						v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v143)+8)) = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v143))) = v10 + int32(8)
						v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v147 + int32(16)
						F_lua_settable(m, l1, int32(-10000))
						mBase = m.M
						v153 = m.ExcPending
						if v153 != 0 {
							return
						} else {
							v154 = m.G386
							F_lua_pushcclosure(m, l1, v154, int32(0))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return
							} else {
								F_lua_pushstring(m, l1, v88+int32(_a139))
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return
								} else {
									F_lua_call(m, l1, int32(1), int32(0))
									mBase = m.M
									v165 = m.ExcPending
									if v165 != 0 {
										return
									} else {
										v166 = m.G387
										F_lua_pushcclosure(m, l1, v166, int32(0))
										mBase = m.M
										v169 = m.ExcPending
										if v169 != 0 {
											return
										} else {
											F_lua_pushstring(m, l1, v88+int32(_a1972))
											mBase = m.M
											v173 = m.ExcPending
											if v173 != 0 {
												return
											} else {
												F_lua_call(m, l1, int32(1), int32(0))
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
													return
												} else {
													v178 = m.G388
													F_lua_pushcclosure(m, l1, v178, int32(0))
													mBase = m.M
													v181 = m.ExcPending
													if v181 != 0 {
														return
													} else {
														F_lua_pushstring(m, l1, v88+int32(_a679))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return
														} else {
															F_lua_call(m, l1, int32(1), int32(0))
															mBase = m.M
															v189 = m.ExcPending
															if v189 != 0 {
																return
															} else {
																v190 = m.G389
																F_lua_pushcclosure(m, l1, v190, int32(0))
																mBase = m.M
																v193 = m.ExcPending
																if v193 != 0 {
																	return
																} else {
																	v195 = v88 + int32(_a1973)
																	F_lua_pushstring(m, l1, v195)
																	mBase = m.M
																	v197 = m.ExcPending
																	if v197 != 0 {
																		return
																	} else {
																		F_lua_call(m, l1, int32(1), int32(0))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return
																		} else {
																			v202 = m.G390
																			F_lua_pushcclosure(m, l1, v202, int32(0))
																			mBase = m.M
																			v205 = m.ExcPending
																			if v205 != 0 {
																				return
																			} else {
																				F_lua_pushstring(m, l1, v88+int32(_a774))
																				mBase = m.M
																				v209 = m.ExcPending
																				if v209 != 0 {
																					return
																				} else {
																					F_lua_call(m, l1, int32(1), int32(0))
																					mBase = m.M
																					v213 = m.ExcPending
																					if v213 != 0 {
																						return
																					} else {
																						v214 = m.G391
																						F_lua_pushcclosure(m, l1, v214, int32(0))
																						mBase = m.M
																						v217 = m.ExcPending
																						if v217 != 0 {
																							return
																						} else {
																							F_lua_pushstring(m, l1, v88+int32(_a1974))
																							mBase = m.M
																							v221 = m.ExcPending
																							if v221 != 0 {
																								return
																							} else {
																								F_lua_call(m, l1, int32(1), int32(0))
																								mBase = m.M
																								v225 = m.ExcPending
																								if v225 != 0 {
																									return
																								} else {
																									v226 = m.G392
																									F_lua_pushcclosure(m, l1, v226, int32(0))
																									mBase = m.M
																									v229 = m.ExcPending
																									if v229 != 0 {
																										return
																									} else {
																										F_lua_pushstring(m, l1, v88+int32(_a1975))
																										mBase = m.M
																										v233 = m.ExcPending
																										if v233 != 0 {
																											return
																										} else {
																											F_lua_call(m, l1, int32(1), int32(0))
																											mBase = m.M
																											v237 = m.ExcPending
																											if v237 != 0 {
																												return
																											} else {
																												v238 = m.G393
																												F_lua_pushcclosure(m, l1, v238, int32(0))
																												mBase = m.M
																												v241 = m.ExcPending
																												if v241 != 0 {
																													return
																												} else {
																													F_lua_pushstring(m, l1, v88+int32(_a1976))
																													mBase = m.M
																													v245 = m.ExcPending
																													if v245 != 0 {
																														return
																													} else {
																														F_lua_call(m, l1, int32(1), int32(0))
																														mBase = m.M
																														v249 = m.ExcPending
																														if v249 != 0 {
																															return
																														} else {
																															v250 = m.G394
																															F_lua_pushcclosure(m, l1, v250, int32(0))
																															mBase = m.M
																															v253 = m.ExcPending
																															if v253 != 0 {
																																return
																															} else {
																																F_lua_pushstring(m, l1, v88+int32(_a1977))
																																mBase = m.M
																																v257 = m.ExcPending
																																if v257 != 0 {
																																	return
																																} else {
																																	F_lua_call(m, l1, int32(1), int32(0))
																																	mBase = m.M
																																	v261 = m.ExcPending
																																	if v261 != 0 {
																																		return
																																	} else {
																																		v262 = m.G395
																																		F_lua_pushcclosure(m, l1, v262, int32(0))
																																		mBase = m.M
																																		v265 = m.ExcPending
																																		if v265 != 0 {
																																			return
																																		} else {
																																			F_lua_pushstring(m, l1, v88+int32(_a1978))
																																			mBase = m.M
																																			v269 = m.ExcPending
																																			if v269 != 0 {
																																				return
																																			} else {
																																				F_lua_call(m, l1, int32(1), int32(0))
																																				mBase = m.M
																																				v273 = m.ExcPending
																																				if v273 != 0 {
																																					return
																																				} else {
																																					F_lua_pushstring(m, l1, v137)
																																					mBase = m.M
																																					v275 = m.ExcPending
																																					if v275 != 0 {
																																						return
																																					} else {
																																						v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																																						*(*int32)(unsafe.Add(mBase, uint32(v277)+8)) = int32(0)
																																						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v277 + int32(16)
																																						F_lua_settable(m, l1, int32(-10000))
																																						mBase = m.M
																																						v285 = m.ExcPending
																																						if v285 != 0 {
																																							return
																																						} else {
																																							F_lua_pushcclosure(m, l1, v81+int32(1198), int32(0))
																																							mBase = m.M
																																							v290 = m.ExcPending
																																							if v290 != 0 {
																																								return
																																							} else {
																																								v293 = v88 + int32(_a1979)
																																								F_lua_setfield(m, l1, int32(-10002), v293)
																																								mBase = m.M
																																								v295 = m.ExcPending
																																								if v295 != 0 {
																																									return
																																								} else {
																																									v296 = int32(0)
																																									F_lua_createtable(m, l1, v296, v296)
																																									mBase = m.M
																																									v299 = m.ExcPending
																																									if v299 != 0 {
																																										return
																																									} else {
																																										F_lua_pushstring(m, l1, v88+int32(_a1980))
																																										mBase = m.M
																																										v303 = m.ExcPending
																																										if v303 != 0 {
																																											return
																																										} else {
																																											F_lua_pushcclosure(m, l1, v81+int32(1199), int32(0))
																																											mBase = m.M
																																											v308 = m.ExcPending
																																											if v308 != 0 {
																																												return
																																											} else {
																																												F_lua_settable(m, l1, int32(-3))
																																												mBase = m.M
																																												v311 = m.ExcPending
																																												if v311 != 0 {
																																													return
																																												} else {
																																													F_lua_pushstring(m, l1, v293)
																																													mBase = m.M
																																													v313 = m.ExcPending
																																													if v313 != 0 {
																																														return
																																													} else {
																																														F_lua_pushcclosure(m, l1, v81+int32(1200), int32(0))
																																														mBase = m.M
																																														v318 = m.ExcPending
																																														if v318 != 0 {
																																															return
																																														} else {
																																															F_lua_settable(m, l1, int32(-3))
																																															mBase = m.M
																																															v321 = m.ExcPending
																																															if v321 != 0 {
																																																return
																																															} else {
																																																F_luaRegisterLogFunction(m, l1)
																																																mBase = m.M
																																																v323 = m.ExcPending
																																																if v323 != 0 {
																																																	return
																																																} else {
																																																	F_luaRegisterVersion(m, l0, l1)
																																																	mBase = m.M
																																																	v325 = m.ExcPending
																																																	if v325 != 0 {
																																																		return
																																																	} else {
																																																		F_lua_pushstring(m, l1, v88+int32(_a1981))
																																																		mBase = m.M
																																																		v329 = m.ExcPending
																																																		if v329 != 0 {
																																																			return
																																																		} else {
																																																			F_lua_pushcclosure(m, l1, v81+int32(1201), int32(0))
																																																			mBase = m.M
																																																			v334 = m.ExcPending
																																																			if v334 != 0 {
																																																				return
																																																			} else {
																																																				F_lua_settable(m, l1, int32(-3))
																																																				mBase = m.M
																																																				v337 = m.ExcPending
																																																				if v337 != 0 {
																																																					return
																																																				} else {
																																																					F_lua_pushstring(m, l1, v88+int32(_a1982))
																																																					mBase = m.M
																																																					v341 = m.ExcPending
																																																					if v341 != 0 {
																																																						return
																																																					} else {
																																																						F_lua_pushcclosure(m, l1, v81+int32(1202), int32(0))
																																																						mBase = m.M
																																																						v346 = m.ExcPending
																																																						if v346 != 0 {
																																																							return
																																																						} else {
																																																							F_lua_settable(m, l1, int32(-3))
																																																							mBase = m.M
																																																							v349 = m.ExcPending
																																																							if v349 != 0 {
																																																								return
																																																							} else {
																																																								F_lua_pushstring(m, l1, v88+int32(_a1983))
																																																								mBase = m.M
																																																								v353 = m.ExcPending
																																																								if v353 != 0 {
																																																									return
																																																								} else {
																																																									F_lua_pushcclosure(m, l1, v81+int32(1203), int32(0))
																																																									mBase = m.M
																																																									v358 = m.ExcPending
																																																									if v358 != 0 {
																																																										return
																																																									} else {
																																																										F_lua_settable(m, l1, int32(-3))
																																																										mBase = m.M
																																																										v361 = m.ExcPending
																																																										if v361 != 0 {
																																																											return
																																																										} else {
																																																											F_lua_pushstring(m, l1, v88+int32(_a1984))
																																																											mBase = m.M
																																																											v365 = m.ExcPending
																																																											if v365 != 0 {
																																																												return
																																																											} else {
																																																												F_lua_pushcclosure(m, l1, v81+int32(1204), int32(0))
																																																												mBase = m.M
																																																												v370 = m.ExcPending
																																																												if v370 != 0 {
																																																													return
																																																												} else {
																																																													F_lua_settable(m, l1, int32(-3))
																																																													mBase = m.M
																																																													v373 = m.ExcPending
																																																													if v373 != 0 {
																																																														return
																																																													} else {
																																																														F_lua_pushstring(m, l1, v88+int32(_a1985))
																																																														mBase = m.M
																																																														v377 = m.ExcPending
																																																														if v377 != 0 {
																																																															return
																																																														} else {
																																																															F_lua_pushcclosure(m, l1, v81+int32(1205), int32(0))
																																																															mBase = m.M
																																																															v382 = m.ExcPending
																																																															if v382 != 0 {
																																																																return
																																																															} else {
																																																																F_lua_settable(m, l1, int32(-3))
																																																																mBase = m.M
																																																																v385 = m.ExcPending
																																																																if v385 != 0 {
																																																																	return
																																																																} else {
																																																																	F_lua_pushstring(m, l1, v88+int32(_a1986))
																																																																	mBase = m.M
																																																																	v389 = m.ExcPending
																																																																	if v389 != 0 {
																																																																		return
																																																																	} else {
																																																																		v392 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																																																																		*(*int32)(unsafe.Add(mBase, uint32(v392)+8)) = int32(3)
																																																																		*(*float64)(unsafe.Add(mBase, uint32(v392))) = float64(0)
																																																																		v396 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																																																																		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v396 + int32(16)
																																																																		F_lua_settable(m, l1, int32(-3))
																																																																		mBase = m.M
																																																																		v402 = m.ExcPending
																																																																		if v402 != 0 {
																																																																			return
																																																																		} else {
																																																																			F_lua_pushstring(m, l1, v88+int32(_a1987))
																																																																			mBase = m.M
																																																																			v406 = m.ExcPending
																																																																			if v406 != 0 {
																																																																				return
																																																																			} else {
																																																																				v409 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																																																																				*(*int32)(unsafe.Add(mBase, uint32(v409)+8)) = int32(3)
																																																																				*(*float64)(unsafe.Add(mBase, uint32(v409))) = float64(1)
																																																																				v413 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																																																																				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v413 + int32(16)
																																																																				F_lua_settable(m, l1, int32(-3))
																																																																				mBase = m.M
																																																																				v419 = m.ExcPending
																																																																				if v419 != 0 {
																																																																					return
																																																																				} else {
																																																																					F_lua_pushstring(m, l1, v88+int32(_a1988))
																																																																					mBase = m.M
																																																																					v423 = m.ExcPending
																																																																					if v423 != 0 {
																																																																						return
																																																																					} else {
																																																																						v426 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																																																																						*(*int32)(unsafe.Add(mBase, uint32(v426)+8)) = int32(3)
																																																																						*(*float64)(unsafe.Add(mBase, uint32(v426))) = float64(2)
																																																																						v430 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																																																																						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v430 + int32(16)
																																																																						F_lua_settable(m, l1, int32(-3))
																																																																						mBase = m.M
																																																																						v436 = m.ExcPending
																																																																						if v436 != 0 {
																																																																							return
																																																																						} else {
																																																																							F_lua_pushstring(m, l1, v88+int32(_a1989))
																																																																							mBase = m.M
																																																																							v440 = m.ExcPending
																																																																							if v440 != 0 {
																																																																								return
																																																																							} else {
																																																																								v443 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																																																																								*(*int32)(unsafe.Add(mBase, uint32(v443)+8)) = int32(3)
																																																																								*(*float64)(unsafe.Add(mBase, uint32(v443))) = float64(2)
																																																																								v447 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																																																																								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v447 + int32(16)
																																																																								F_lua_settable(m, l1, int32(-3))
																																																																								mBase = m.M
																																																																								v453 = m.ExcPending
																																																																								if v453 != 0 {
																																																																									return
																																																																								} else {
																																																																									F_lua_pushstring(m, l1, v88+int32(_a1990))
																																																																									mBase = m.M
																																																																									v457 = m.ExcPending
																																																																									if v457 != 0 {
																																																																										return
																																																																									} else {
																																																																										v460 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																																																																										*(*int32)(unsafe.Add(mBase, uint32(v460)+8)) = int32(3)
																																																																										*(*float64)(unsafe.Add(mBase, uint32(v460))) = float64(3)
																																																																										v464 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																																																																										*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v464 + int32(16)
																																																																										F_lua_settable(m, l1, int32(-3))
																																																																										mBase = m.M
																																																																										v470 = m.ExcPending
																																																																										if v470 != 0 {
																																																																											return
																																																																										} else {
																																																																											F_lua_pushstring(m, l1, v88+int32(_a1991))
																																																																											mBase = m.M
																																																																											v474 = m.ExcPending
																																																																											if v474 != 0 {
																																																																												return
																																																																											} else {
																																																																												F_lua_pushcclosure(m, l1, v81+int32(1206), int32(0))
																																																																												mBase = m.M
																																																																												v479 = m.ExcPending
																																																																												if v479 != 0 {
																																																																													return
																																																																												} else {
																																																																													F_lua_settable(m, l1, int32(-3))
																																																																													mBase = m.M
																																																																													v482 = m.ExcPending
																																																																													if v482 != 0 {
																																																																														return
																																																																													} else {
																																																																														v485 = v88 + int32(_a980)
																																																																														F_lua_setfield(m, l1, int32(-10002), v485)
																																																																														mBase = m.M
																																																																														v487 = m.ExcPending
																																																																														if v487 != 0 {
																																																																															return
																																																																														} else {
																																																																															F_lua_getfield(m, l1, int32(-10002), v485)
																																																																															mBase = m.M
																																																																															v490 = m.ExcPending
																																																																															if v490 != 0 {
																																																																																return
																																																																															} else {
																																																																																F_lua_setfield(m, l1, int32(-10002), v88+int32(_a1039))
																																																																																mBase = m.M
																																																																																v495 = m.ExcPending
																																																																																if v495 != 0 {
																																																																																	return
																																																																																} else {
																																																																																	F_lua_getfield(m, l1, int32(-10002), v195)
																																																																																	mBase = m.M
																																																																																	v498 = m.ExcPending
																																																																																	if v498 != 0 {
																																																																																		return
																																																																																	} else {
																																																																																		F_lua_pushstring(m, l1, v88+int32(_a1992))
																																																																																		mBase = m.M
																																																																																		v502 = m.ExcPending
																																																																																		if v502 != 0 {
																																																																																			return
																																																																																		} else {
																																																																																			F_lua_pushcclosure(m, l1, v81+int32(1207), int32(0))
																																																																																			mBase = m.M
																																																																																			v507 = m.ExcPending
																																																																																			if v507 != 0 {
																																																																																				return
																																																																																			} else {
																																																																																				F_lua_settable(m, l1, int32(-3))
																																																																																				mBase = m.M
																																																																																				v510 = m.ExcPending
																																																																																				if v510 != 0 {
																																																																																					return
																																																																																				} else {
																																																																																					F_lua_pushstring(m, l1, v88+int32(_a1993))
																																																																																					mBase = m.M
																																																																																					v514 = m.ExcPending
																																																																																					if v514 != 0 {
																																																																																						return
																																																																																					} else {
																																																																																						F_lua_pushcclosure(m, l1, v81+int32(1208), int32(0))
																																																																																						mBase = m.M
																																																																																						v519 = m.ExcPending
																																																																																						if v519 != 0 {
																																																																																							return
																																																																																						} else {
																																																																																							F_lua_settable(m, l1, int32(-3))
																																																																																							mBase = m.M
																																																																																							v522 = m.ExcPending
																																																																																							if v522 != 0 {
																																																																																								return
																																																																																							} else {
																																																																																								F_lua_setfield(m, l1, int32(-10002), v195)
																																																																																								mBase = m.M
																																																																																								v525 = m.ExcPending
																																																																																								if v525 != 0 {
																																																																																									return
																																																																																								} else {
																																																																																									m.G0 = v10 + int32(32)
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
}
func F_lua_atpanic(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+88)) = l1
	return v5
}
func F_lua_call(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_luaD_call(m, l0, v4+(l1^int32(-1))<<(uint(int32(4))%32), l2)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if l2 != int32(-1) {
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
			if base.Ui32(v14) < base.Ui32(v16) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v14
			}
		}
		return
	}
}
func F_lua_cjson_new(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v106 int64
	_ = v106
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v282 int32
	_ = v282
	var v283 int64
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v357 int32
	_ = v357
	var v358 int64
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v432 int32
	_ = v432
	var v433 int64
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v507 int32
	_ = v507
	var v508 int64
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v582 int32
	_ = v582
	var v583 int64
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v657 int32
	_ = v657
	var v658 int64
	_ = v658
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v732 int32
	_ = v732
	var v733 int64
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v807 int32
	_ = v807
	var v808 int64
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v882 int32
	_ = v882
	var v883 int64
	_ = v883
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v919 int32
	_ = v919
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	F_fpconv_init(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(0)
		F_lua_createtable(m, l0, v9, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = F_lua_newuserdata(m, l0, int32(1340))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = int32(0)
				F_lua_createtable(m, l0, v17, v17)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = m.G5
					F_lua_pushcclosure(m, l0, v21+int32(1344), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v28 = m.G3
						F_lua_setfield(m, l0, int32(-2), v28+int32(_a2108))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v34 = F_lua_setmetatable(m, l0, int32(-2))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v15)+1312)) = int64(4294967296010)
								*(*int64)(unsafe.Add(mBase, uint32(v15)+1304)) = int64(8589934592)
								*(*int64)(unsafe.Add(mBase, uint32(v15)+1332)) = int64(4294967296001)
								v42 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v15)+1320)) = v42
								*(*int64)(unsafe.Add(mBase, uint32(v15)+1324)) = int64(4294967310)
								F_strbuf_init(m, v15+int32(1280), v42)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v52 = v9
									for {
										v57 = v15 + v52<<(uint(int32(2))%32)
										v58 = int64(51539607564)
										*(*int64)(unsafe.Add(mBase, uint32(v57))) = v58
										v60 = int32(8)
										*(*int64)(unsafe.Add(mBase, uint32(v57+v60))) = v58
										*(*int64)(unsafe.Add(mBase, uint32(v57+int32(16)))) = v58
										*(*int64)(unsafe.Add(mBase, uint32(v57+int32(24)))) = v58
										v73 = v52 + v60
										if v73 != int32(256) {
											v52 = v73
											continue
										} else {
											break
										}
										break
									}
									*(*int32)(unsafe.Add(mBase, uint32(v15)+500)) = int32(1)
									v78 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+492)) = v78
									*(*int32)(unsafe.Add(mBase, uint32(v15)+372)) = int32(3)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+364)) = int32(2)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+232)) = int32(8)
									v86 = int32(11)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v86
									*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(10)
									v90 = int32(13)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+420)) = v90
									*(*int32)(unsafe.Add(mBase, uint32(v15)+408)) = v90
									*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v86
									*(*int64)(unsafe.Add(mBase, uint32(v15)+36)) = int64(47244640267)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+440)) = v90
									*(*int32)(unsafe.Add(mBase, uint32(v15)+292)) = v90
									*(*int32)(unsafe.Add(mBase, uint32(v15)+464)) = v90
									*(*int32)(unsafe.Add(mBase, uint32(v15)+312)) = v90
									v106 = int64(55834574861)
									*(*int64)(unsafe.Add(mBase, uint32(v15)+224)) = v106
									*(*int64)(unsafe.Add(mBase, uint32(v15)+216)) = v106
									*(*int64)(unsafe.Add(mBase, uint32(v15)+208)) = v106
									*(*int64)(unsafe.Add(mBase, uint32(v15)+200)) = v106
									*(*int64)(unsafe.Add(mBase, uint32(v15)+192)) = v106
									*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v90
									*(*int64)(unsafe.Add(mBase, uint32(v15)+172)) = int64(38654705677)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v90
									v127 = F__emscripten_memset_bulkmem(m, v15+int32(1024), base.I32_extend8_s(v78), int32(256))
									mBase = m.M
									v128 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+1116)) = uint8(v128)
									v130 = int32(34)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+1058)) = uint8(v130)
									v132 = int32(8)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+1122)) = uint8(v132)
									v134 = int32(47)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+1071)) = uint8(v134)
									v136 = int32(10)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+1134)) = uint8(v136)
									v138 = int32(29961)
									*(*uint16)(unsafe.Add(mBase, uint32(v15)+1140)) = uint16(v138)
									v140 = int32(13)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+1138)) = uint8(v140)
									v142 = int32(12)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+1126)) = uint8(v142)
									v145 = m.G3
									F_luaL_checkstack(m, l0, int32(1), v145+int32(_a2109))
									mBase = m.M
									v149 = m.ExcPending
									if v149 != 0 {
										return int32(0)
									} else {
										v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v170 = v167 + int32(-16)
										v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v207 = *(*int64)(unsafe.Add(mBase, uint32(v170)))
										*(*int64)(unsafe.Add(mBase, uint32(v206))) = v207
										v209 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v206)+8)) = v209
										v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v211 + int32(16)
										v215 = m.G5
										F_lua_pushcclosure(m, l0, v215+int32(1345), int32(1))
										mBase = m.M
										v220 = m.ExcPending
										if v220 != 0 {
											return int32(0)
										} else {
											F_lua_setfield(m, l0, int32(-3), v145+int32(_a2110))
											mBase = m.M
											v225 = m.ExcPending
											if v225 != 0 {
												return int32(0)
											} else {
												v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v246 = v243 + int32(-16)
												v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v283 = *(*int64)(unsafe.Add(mBase, uint32(v246)))
												*(*int64)(unsafe.Add(mBase, uint32(v282))) = v283
												v285 = *(*int32)(unsafe.Add(mBase, uint32(v246)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v282)+8)) = v285
												v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v287 + int32(16)
												F_lua_pushcclosure(m, l0, v215+int32(1346), int32(1))
												mBase = m.M
												v295 = m.ExcPending
												if v295 != 0 {
													return int32(0)
												} else {
													F_lua_setfield(m, l0, int32(-3), v145+int32(_a2111))
													mBase = m.M
													v300 = m.ExcPending
													if v300 != 0 {
														return int32(0)
													} else {
														v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v321 = v318 + int32(-16)
														v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v358 = *(*int64)(unsafe.Add(mBase, uint32(v321)))
														*(*int64)(unsafe.Add(mBase, uint32(v357))) = v358
														v360 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v357)+8)) = v360
														v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v362 + int32(16)
														F_lua_pushcclosure(m, l0, v215+int32(1347), int32(1))
														mBase = m.M
														v370 = m.ExcPending
														if v370 != 0 {
															return int32(0)
														} else {
															F_lua_setfield(m, l0, int32(-3), v145+int32(_a2112))
															mBase = m.M
															v375 = m.ExcPending
															if v375 != 0 {
																return int32(0)
															} else {
																v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																v396 = v393 + int32(-16)
																v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																v433 = *(*int64)(unsafe.Add(mBase, uint32(v396)))
																*(*int64)(unsafe.Add(mBase, uint32(v432))) = v433
																v435 = *(*int32)(unsafe.Add(mBase, uint32(v396)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v432)+8)) = v435
																v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v437 + int32(16)
																F_lua_pushcclosure(m, l0, v215+int32(1348), int32(1))
																mBase = m.M
																v445 = m.ExcPending
																if v445 != 0 {
																	return int32(0)
																} else {
																	F_lua_setfield(m, l0, int32(-3), v145+int32(_a2113))
																	mBase = m.M
																	v450 = m.ExcPending
																	if v450 != 0 {
																		return int32(0)
																	} else {
																		v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		v471 = v468 + int32(-16)
																		v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		v508 = *(*int64)(unsafe.Add(mBase, uint32(v471)))
																		*(*int64)(unsafe.Add(mBase, uint32(v507))) = v508
																		v510 = *(*int32)(unsafe.Add(mBase, uint32(v471)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(v507)+8)) = v510
																		v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v512 + int32(16)
																		F_lua_pushcclosure(m, l0, v215+int32(1349), int32(1))
																		mBase = m.M
																		v520 = m.ExcPending
																		if v520 != 0 {
																			return int32(0)
																		} else {
																			F_lua_setfield(m, l0, int32(-3), v145+int32(_a2114))
																			mBase = m.M
																			v525 = m.ExcPending
																			if v525 != 0 {
																				return int32(0)
																			} else {
																				v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				v546 = v543 + int32(-16)
																				v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				v583 = *(*int64)(unsafe.Add(mBase, uint32(v546)))
																				*(*int64)(unsafe.Add(mBase, uint32(v582))) = v583
																				v585 = *(*int32)(unsafe.Add(mBase, uint32(v546)+8))
																				*(*int32)(unsafe.Add(mBase, uint32(v582)+8)) = v585
																				v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v587 + int32(16)
																				F_lua_pushcclosure(m, l0, v215+int32(1350), int32(1))
																				mBase = m.M
																				v595 = m.ExcPending
																				if v595 != 0 {
																					return int32(0)
																				} else {
																					F_lua_setfield(m, l0, int32(-3), v145+int32(_a2115))
																					mBase = m.M
																					v600 = m.ExcPending
																					if v600 != 0 {
																						return int32(0)
																					} else {
																						v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						v621 = v618 + int32(-16)
																						v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						v658 = *(*int64)(unsafe.Add(mBase, uint32(v621)))
																						*(*int64)(unsafe.Add(mBase, uint32(v657))) = v658
																						v660 = *(*int32)(unsafe.Add(mBase, uint32(v621)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v657)+8)) = v660
																						v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v662 + int32(16)
																						F_lua_pushcclosure(m, l0, v215+int32(1351), int32(1))
																						mBase = m.M
																						v670 = m.ExcPending
																						if v670 != 0 {
																							return int32(0)
																						} else {
																							F_lua_setfield(m, l0, int32(-3), v145+int32(_a2116))
																							mBase = m.M
																							v675 = m.ExcPending
																							if v675 != 0 {
																								return int32(0)
																							} else {
																								v693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								v696 = v693 + int32(-16)
																								v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								v733 = *(*int64)(unsafe.Add(mBase, uint32(v696)))
																								*(*int64)(unsafe.Add(mBase, uint32(v732))) = v733
																								v735 = *(*int32)(unsafe.Add(mBase, uint32(v696)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v732)+8)) = v735
																								v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v737 + int32(16)
																								F_lua_pushcclosure(m, l0, v215+int32(1352), int32(1))
																								mBase = m.M
																								v745 = m.ExcPending
																								if v745 != 0 {
																									return int32(0)
																								} else {
																									F_lua_setfield(m, l0, int32(-3), v145+int32(_a2117))
																									mBase = m.M
																									v750 = m.ExcPending
																									if v750 != 0 {
																										return int32(0)
																									} else {
																										v768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										v771 = v768 + int32(-16)
																										v807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										v808 = *(*int64)(unsafe.Add(mBase, uint32(v771)))
																										*(*int64)(unsafe.Add(mBase, uint32(v807))) = v808
																										v810 = *(*int32)(unsafe.Add(mBase, uint32(v771)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v807)+8)) = v810
																										v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v812 + int32(16)
																										F_lua_pushcclosure(m, l0, v215+int32(1353), int32(1))
																										mBase = m.M
																										v820 = m.ExcPending
																										if v820 != 0 {
																											return int32(0)
																										} else {
																											F_lua_setfield(m, l0, int32(-3), v145+int32(_a2118))
																											mBase = m.M
																											v825 = m.ExcPending
																											if v825 != 0 {
																												return int32(0)
																											} else {
																												v843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																												v846 = v843 + int32(-16)
																												v882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																												v883 = *(*int64)(unsafe.Add(mBase, uint32(v846)))
																												*(*int64)(unsafe.Add(mBase, uint32(v882))) = v883
																												v885 = *(*int32)(unsafe.Add(mBase, uint32(v846)+8))
																												*(*int32)(unsafe.Add(mBase, uint32(v882)+8)) = v885
																												v887 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v887 + int32(16)
																												F_lua_pushcclosure(m, l0, v215+int32(1354), int32(1))
																												mBase = m.M
																												v895 = m.ExcPending
																												if v895 != 0 {
																													return int32(0)
																												} else {
																													F_lua_setfield(m, l0, int32(-3), v145+int32(_a2119))
																													mBase = m.M
																													v900 = m.ExcPending
																													if v900 != 0 {
																														return int32(0)
																													} else {
																														v919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v919 + int32(-16)
																														v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																														*(*int32)(unsafe.Add(mBase, uint32(v931)+8)) = int32(2)
																														*(*int32)(unsafe.Add(mBase, uint32(v931))) = int32(0)
																														v935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v935 + int32(16)
																														F_lua_setfield(m, l0, int32(-2), v145+int32(_a684))
																														mBase = m.M
																														v943 = m.ExcPending
																														if v943 != 0 {
																															return int32(0)
																														} else {
																															F_lua_pushlstring(m, l0, v145+int32(_a1975), int32(5))
																															mBase = m.M
																															v948 = m.ExcPending
																															if v948 != 0 {
																																return int32(0)
																															} else {
																																F_lua_setfield(m, l0, int32(-2), v145+int32(_a2120))
																																mBase = m.M
																																v953 = m.ExcPending
																																if v953 != 0 {
																																	return int32(0)
																																} else {
																																	F_lua_pushlstring(m, l0, v145+int32(_a2121), int32(5))
																																	mBase = m.M
																																	v958 = m.ExcPending
																																	if v958 != 0 {
																																		return int32(0)
																																	} else {
																																		F_lua_setfield(m, l0, int32(-2), v145+int32(_a2122))
																																		mBase = m.M
																																		v963 = m.ExcPending
																																		if v963 != 0 {
																																			return int32(0)
																																		} else {
																																			return int32(1)
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
	}
}
func F_lua_enablereadonlytable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v56 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v56 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v56 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v56 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v56 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v56 = v16
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = l2
	return
}
func F_lua_getfield(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 < int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l2&int32(3) == int32(0) {
		v85 = l2
		goto L18
	} else {
		goto L19
	}
L2:
	;
	if l1 < int32(-9999) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = v14 + l1<<(uint(int32(4))%32) + int32(-16)
	v20 = m.G398
	if base.Ui32(v19) < base.Ui32(v13) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = v19
	goto L6
L5:
	;
	v22 = v20
	goto L6
L6:
	;
	v61 = v22
	goto L1
L7:
	;
	switch l1 + int32(10002) {
	case 0:
		goto L10
	case 1:
		goto L11
	case 2:
		goto L12
	default:
		goto L9
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v61 = v25 + l1<<(uint(int32(4))%32)
	goto L1
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+7)))
	v49 = m.G398
	if base.Ui32(v48) < base.Ui32(int32(-10002)-l1) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v61 = l0 + int32(72)
	goto L1
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v37
	v61 = l0 + int32(88)
	goto L1
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v61 = v31 + int32(96)
	goto L1
L13:
	;
	v60 = v49
	goto L15
L14:
	;
	v60 = v47 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
	goto L15
L15:
	;
	v61 = v60
	goto L1
L16:
	;
	v119 = F_luaS_newlstr(m, l0, l2, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L32
	} else {
		goto L33
	}
L17:
	;
	v118 = v110 - l2
	goto L16
L18:
	;
	v89 = v85
	goto L26
L19:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v71 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v74 = l2
	goto L22
L21:
	;
	v118 = l2 - l2
	goto L16
L22:
	;
	v78 = v74 + int32(1)
	if v78&int32(3) == int32(0) {
		v85 = v78
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v83 != 0 {
		v74 = v78
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v110 = v78
	goto L17
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v98 = int32(-2139062144)
	if (int32(16843008)-v95|v95)&v98 == v98 {
		v89 = v89 + int32(4)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v104 = v89
	goto L29
L28:
	;
	goto L27
L29:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v108 != 0 {
		v104 = v104 + int32(1)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v110 = v104
	goto L17
L31:
	;
	goto L30
L32:
	;
	return
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v119
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_luaV_gettable(m, l0, v61, v9, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v128 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v127 + v128
	m.G0 = v9 + v128
	return
}
func F_lua_gettable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v55 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v55 = v16
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v60 = v58 + int32(-16)
	F_luaV_gettable(m, l0, v55, v60, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		return
	} else {
		return
	}
}
func F_lua_getupvalue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v56 = l0 + int32(72)
			case 1:
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v32
				v56 = l0 + int32(88)
			case 2:
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v56 = v26 + int32(96)
			default:
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+7)))
				v44 = m.G398
				if base.Ui32(v43) < base.Ui32(int32(-10002)-l1) {
					v55 = v44
				} else {
					v55 = v42 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v56 = v55
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v56 = v20 + l1<<(uint(int32(4))%32)
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = v9 + l1<<(uint(int32(4))%32) + int32(-16)
		v15 = m.G398
		if base.Ui32(v14) < base.Ui32(v8) {
			v17 = v14
		} else {
			v17 = v15
		}
		v56 = v17
	}
	v59 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v60 != int32(6) {
		v111 = v59
	} else {
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
		v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+6)))
		if v64 == int32(0) {
			if l2 < int32(1) {
				v111 = v59
			} else {
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+36))
				if v82 < l2 {
					v111 = v59
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+28))
					v88 = l2<<(uint(int32(2))%32) + int32(-4)
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v84+v88)))
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v63+v88)+20))
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
					v96 = v95
					v98 = v90 + int32(16)
					v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v101 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
					*(*int64)(unsafe.Add(mBase, uint32(v100))) = v101
					v103 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v103
					v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v105 + int32(16)
					v111 = v98
				}
			}
		} else {
			if l2 < int32(1) {
				v111 = v59
			} else {
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+7)))
				if base.Ui32(v69) < base.Ui32(l2) {
					v111 = v59
				} else {
					v71 = m.G3
					v96 = l2<<(uint(int32(4))%32) + v63 + int32(8)
					v98 = v71 + int32(_a139)
					v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v101 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
					*(*int64)(unsafe.Add(mBase, uint32(v100))) = v101
					v103 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v103
					v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v105 + int32(16)
					v111 = v98
				}
			}
		}
	}
	return v111
}
func F_lua_iscfunction(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v30
				v55 = l0 + int32(88)
			case 2:
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v24 + int32(96)
			default:
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+7)))
				v42 = m.G398
				if base.Ui32(v41) < base.Ui32(int32(-10002)-l1) {
					v53 = v42
				} else {
					v53 = v40 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v53
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v18 + l1<<(uint(int32(4))%32)
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v12 = v7 + l1<<(uint(int32(4))%32) + int32(-16)
		v13 = m.G398
		if base.Ui32(v12) < base.Ui32(v6) {
			v15 = v12
		} else {
			v15 = v13
		}
		v55 = v15
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v58 != int32(6) {
		v65 = int32(0)
	} else {
		v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+6)))
		v65 = base.B2i32(v62 != int32(0))
	}
	return v65
}
func F_lua_newuserdata(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+68))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+64))
	if base.Ui32(v5) < base.Ui32(v6) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v12 != v13 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v21 = v18 + int32(12)
		} else {
			v21 = l0 + int32(72)
		}
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		v23 = F_luaS_newudata(m, l0, l1, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(7)
			*(*int32)(unsafe.Add(mBase, uint32(v25))) = v23
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v29 + int32(16)
			return v23 + int32(24)
		}
	} else {
		F_luaC_step(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v12 != v13 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v21 = v18 + int32(12)
			} else {
				v21 = l0 + int32(72)
			}
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v23 = F_luaS_newudata(m, l0, l1, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(7)
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = v23
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v29 + int32(16)
				return v23 + int32(24)
			}
		}
	}
}
func F_lua_next(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v55 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v55 = v16
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v62 = F_luaH_next(m, l0, v58, v59+int32(-16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		return int32(0)
	} else {
		v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v62 != 0 {
			v69 = int32(16)
		} else {
			v69 = int32(-16)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v66 + v69
		return v62
	}
}
func F_lua_pushnumber(m *base.Module, l0 int32, l1 float64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v4))) = l1
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v8 + int32(16)
	return
}
func F_lua_pushvalue(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v55 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v55 = v16
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v63 + int32(16)
	return
}
func F_lua_rawgeti(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 float64
	_ = v73
	var v74 int32
	_ = v74
	var v75 float64
	_ = v75
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 float64
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 float64
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	if l1 < int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if l2 < int32(1) {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	if l1 < int32(-9999) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = v9 + l1<<(uint(int32(4))%32) + int32(-16)
	v15 = m.G398
	if base.Ui32(v14) < base.Ui32(v8) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = v14
	goto L6
L5:
	;
	v17 = v15
	goto L6
L6:
	;
	v56 = v17
	goto L1
L7:
	;
	switch l1 + int32(10002) {
	case 0:
		goto L10
	case 1:
		goto L11
	case 2:
		goto L12
	default:
		goto L9
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v56 = v20 + l1<<(uint(int32(4))%32)
	goto L1
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+7)))
	v44 = m.G398
	if base.Ui32(v43) < base.Ui32(int32(-10002)-l1) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v56 = l0 + int32(72)
	goto L1
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v32
	v56 = l0 + int32(88)
	goto L1
L12:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v56 = v26 + int32(96)
	goto L1
L13:
	;
	v55 = v44
	goto L15
L14:
	;
	v55 = v42 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
	goto L15
L15:
	;
	v56 = v55
	goto L1
L16:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v111)))
	*(*int64)(unsafe.Add(mBase, uint32(v112))) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v117 + int32(16)
	return
L17:
	;
	v97 = v93
	goto L24
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	v77 = base.I64_reinterpret_f64(v75)
	v82 = int32(-1)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+12)))
	v89 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v77)>>(uint(int64(32))%64))+v77), v82<<(uint(v83)%32)^v82|int32(1))
	v93 = v76 + v89<<(uint(int32(5))%32)
	v94 = v75
	goto L17
L19:
	;
	v73 = base.F64_convert_i32_s(l2)
	if l2 != 0 {
		v75 = v73
		goto L18
	} else {
		goto L23
	}
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+36))
	if l2 <= v64 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	v111 = v67 + l2<<(uint(int32(4))%32) + int32(-16)
	goto L16
L22:
	;
	v75 = base.F64_convert_i32_u(l2)
	goto L18
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	v93 = v74
	v94 = v73
	goto L17
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
	if v100 != int32(3) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v111 = v106
	goto L16
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	v106 = m.G398
	if v105 != 0 {
		v97 = v105
		goto L24
	} else {
		goto L29
	}
L27:
	;
	v103 = *(*float64)(unsafe.Add(mBase, uint32(v97)+16))
	if base.F64_ne(v103, v94) != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v111 = v97
	goto L16
L29:
	;
	goto L25
}
func F_lua_rawset(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v56 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v56 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v56 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v56 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v56 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v56 = v16
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	if v59 == int32(0) {
		v69 = v58
		v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v73 = F_luaH_set(m, l0, v69, v70+int32(-32))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return
		} else {
			v77 = *(*int64)(unsafe.Add(mBase, uint32(v70+int32(-16))))
			*(*int64)(unsafe.Add(mBase, uint32(v73))) = v77
			v79 = int32(-8)
			v81 = *(*int32)(unsafe.Add(mBase, uint32(v70+v79)))
			*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v81
			v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)))
			if v86 < int32(4) {
				v112 = v83
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(-16))))
				v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+5)))
				if v92&int32(3) == int32(0) {
					v112 = v83
				} else {
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
					v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)))
					if v98&int32(4) == int32(0) {
						v112 = v83
					} else {
						v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)))
						v106 = v104 & int32(251)
						*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)) = uint8(v106)
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v103)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v97)+32)) = v108
						*(*int32)(unsafe.Add(mBase, uint32(v103)+40)) = v97
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v112 = v111
					}
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v112 + int32(-32)
			return
		}
	} else {
		v62 = m.G3
		F_luaG_runerror(m, l0, v62+int32(_a2019), int32(0))
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return
		} else {
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
			v69 = v68
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v73 = F_luaH_set(m, l0, v69, v70+int32(-32))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return
			} else {
				v77 = *(*int64)(unsafe.Add(mBase, uint32(v70+int32(-16))))
				*(*int64)(unsafe.Add(mBase, uint32(v73))) = v77
				v79 = int32(-8)
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v70+v79)))
				*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v81
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v83+v79)))
				if v86 < int32(4) {
					v112 = v83
				} else {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(-16))))
					v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+5)))
					if v92&int32(3) == int32(0) {
						v112 = v83
					} else {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
						v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)))
						if v98&int32(4) == int32(0) {
							v112 = v83
						} else {
							v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)))
							v106 = v104 & int32(251)
							*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)) = uint8(v106)
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v103)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v97)+32)) = v108
							*(*int32)(unsafe.Add(mBase, uint32(v103)+40)) = v97
							v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v112 = v111
						}
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v112 + int32(-32)
				return
			}
		}
	}
}
func F_lua_sethook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = l3
	if l2 != 0 {
		v8 = l1
	} else {
		v8 = int32(0)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v8
	if l1 != 0 {
		v11 = l2
	} else {
		v11 = int32(0)
	}
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)) = uint8(v11)
	return int32(1)
}
func F_lua_setlevel(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+52)) = uint16(v3)
	return
}
func F_lua_settable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v55 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v55 = v16
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_luaV_settable(m, l0, v55, v58+int32(-32), v58+int32(-16))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		return
	} else {
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v65 + int32(-32)
		return
	}
}
func F_lua_status(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	return v2
}
func F_lua_tolstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v56 = l0 + int32(72)
			case 1:
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v32
				v56 = l0 + int32(88)
			case 2:
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v56 = v26 + int32(96)
			default:
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+7)))
				v44 = m.G398
				if base.Ui32(v43) < base.Ui32(int32(-10002)-l1) {
					v55 = v44
				} else {
					v55 = v42 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v56 = v55
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v56 = v20 + l1<<(uint(int32(4))%32)
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = v9 + l1<<(uint(int32(4))%32) + int32(-16)
		v15 = m.G398
		if base.Ui32(v14) < base.Ui32(v8) {
			v17 = v14
		} else {
			v17 = v15
		}
		v56 = v17
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v58 == int32(4) {
		v130 = v56
		v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
		if l2 == int32(0) {
			v137 = v131
		} else {
			v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v134
			v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
			v137 = v136
		}
		v140 = v137 + int32(16)
		return v140
	} else {
		v61 = F_luaV_tostring(m, l0, v56)
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int32(0)
		} else {
			if v61 != 0 {
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+68))
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+64))
				if base.Ui32(v73) < base.Ui32(v74) {
					if l1 < int32(1) {
						if l1 < int32(-9999) {
							switch l1 + int32(10002) {
							case 0:
								v130 = l0 + int32(72)
							case 1:
								v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v104
								v130 = l0 + int32(88)
							case 2:
								v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v130 = v98 + int32(96)
							default:
								v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
								v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)))
								v116 = m.G398
								if base.Ui32(v115) < base.Ui32(int32(-10002)-l1) {
									v127 = v116
								} else {
									v127 = v114 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
								}
								v130 = v127
							}
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v130 = v92 + l1<<(uint(int32(4))%32)
						}
					} else {
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v86 = v81 + l1<<(uint(int32(4))%32) + int32(-16)
						v87 = m.G398
						if base.Ui32(v86) < base.Ui32(v80) {
							v89 = v86
						} else {
							v89 = v87
						}
						v130 = v89
					}
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
					if l2 == int32(0) {
						v137 = v131
					} else {
						v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v134
						v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
						v137 = v136
					}
					v140 = v137 + int32(16)
					return v140
				} else {
					F_luaC_step(m, l0)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						if l1 < int32(1) {
							if l1 < int32(-9999) {
								switch l1 + int32(10002) {
								case 0:
									v130 = l0 + int32(72)
								case 1:
									v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
									v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v104
									v130 = l0 + int32(88)
								case 2:
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v130 = v98 + int32(96)
								default:
									v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
									v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)))
									v116 = m.G398
									if base.Ui32(v115) < base.Ui32(int32(-10002)-l1) {
										v127 = v116
									} else {
										v127 = v114 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
									}
									v130 = v127
								}
							} else {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v130 = v92 + l1<<(uint(int32(4))%32)
							}
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v86 = v81 + l1<<(uint(int32(4))%32) + int32(-16)
							v87 = m.G398
							if base.Ui32(v86) < base.Ui32(v80) {
								v89 = v86
							} else {
								v89 = v87
							}
							v130 = v89
						}
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
						if l2 == int32(0) {
							v137 = v131
						} else {
							v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v134
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
							v137 = v136
						}
						v140 = v137 + int32(16)
						return v140
					}
				}
			} else {
				v65 = int32(0)
				if l2 == v65 {
					v140 = v65
					return v140
				} else {
					v68 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v68
					return v68
				}
			}
		}
	}
}
func F_lua_tothread(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v30
				v55 = l0 + int32(88)
			case 2:
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v24 + int32(96)
			default:
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+7)))
				v42 = m.G398
				if base.Ui32(v41) < base.Ui32(int32(-10002)-l1) {
					v53 = v42
				} else {
					v53 = v40 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v53
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v18 + l1<<(uint(int32(4))%32)
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v12 = v7 + l1<<(uint(int32(4))%32) + int32(-16)
		v13 = m.G398
		if base.Ui32(v12) < base.Ui32(v6) {
			v15 = v12
		} else {
			v15 = v13
		}
		v55 = v15
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v58 != int32(8) {
		v62 = int32(0)
	} else {
		v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		v62 = v61
	}
	return v62
}
func F_lua_type(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v54 = l0 + int32(72)
				v55 = m.G398
				if v54 != v55 {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
					return v59
				} else {
					return int32(-1)
				}
			case 1:
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v26
				v54 = l0 + int32(88)
				v55 = m.G398
				if v54 != v55 {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
					return v59
				} else {
					return int32(-1)
				}
			case 2:
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v54 = v50 + int32(96)
				v55 = m.G398
				if v54 != v55 {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
					return v59
				} else {
					return int32(-1)
				}
			default:
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+7)))
				if base.Ui32(v39) < base.Ui32(int32(-10002)-l1) {
					return int32(-1)
				} else {
					v54 = v38 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
					v55 = m.G398
					if v54 != v55 {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
						return v59
					} else {
						return int32(-1)
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v54 = v17 + l1<<(uint(int32(4))%32)
			v55 = m.G398
			if v54 != v55 {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
				return v59
			} else {
				return int32(-1)
			}
		}
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v10 = v5 + l1<<(uint(int32(4))%32) + int32(-16)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v10) < base.Ui32(v11) {
			v54 = v10
			v55 = m.G398
			if v54 != v55 {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
				return v59
			} else {
				return int32(-1)
			}
		} else {
			return int32(-1)
		}
	}
}
func F_lua_typename(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v4 = m.G3
	if l1 != int32(-1) {
		v10 = m.G399
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10+l1<<(uint(int32(2))%32))))
		return v14
	} else {
		return v4 + int32(_a2018)
	}
}
func F_lua_yield(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
	if base.Ui32(v3) <= base.Ui32(v4) {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v14)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v16 - l1<<(uint(int32(4))%32)
		return int32(-1)
	} else {
		v6 = m.G3
		F_luaG_runerror(m, l0, v6+int32(_a2022), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v14)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v16 - l1<<(uint(int32(4))%32)
			return int32(-1)
		}
	}
}
func F_updateLuaEnableInsecureApi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v3 = int32(_a44)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	if v4 == v6 {
		v18 = v4
		*(*int32)(unsafe.Add(mBase, _consts[272])) = v18
		return int32(1)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[138]))
		F_evalReset(m, base.B2i32(v9 != int32(0)))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[273]))
			v18 = v17
			*(*int32)(unsafe.Add(mBase, _consts[272])) = v18
			return int32(1)
		}
	}
}
