package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaD_call(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
	v8 = v6 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v8)
	v11 = v8 & int32(65535)
	if base.Ui32(v11) < base.Ui32(int32(200)) {
		v26 = F_luaD_precall(m, l0, l1, l2)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			if v26 != 0 {
				v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
				v33 = v31 + int32(-1)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v33)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+64))
				if base.Ui32(v36) < base.Ui32(v37) {
					return
				} else {
					F_luaC_step(m, l0)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				F_luaV_execute(m, l0, int32(1))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
					v33 = v31 + int32(-1)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v33)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+64))
					if base.Ui32(v36) < base.Ui32(v37) {
						return
					} else {
						F_luaC_step(m, l0)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		if v11 != int32(200) {
			if base.Ui32(int32(225)) <= base.Ui32(v8&int32(65535)) {
				F_luaD_throw(m, l0, int32(5))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v26 = F_luaD_precall(m, l0, l1, l2)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					if v26 != 0 {
						v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
						v33 = v31 + int32(-1)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v33)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+64))
						if base.Ui32(v36) < base.Ui32(v37) {
							return
						} else {
							F_luaC_step(m, l0)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						F_luaV_execute(m, l0, int32(1))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
							v33 = v31 + int32(-1)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v33)
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+64))
							if base.Ui32(v36) < base.Ui32(v37) {
								return
							} else {
								F_luaC_step(m, l0)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
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
			v16 = m.G3
			F_luaG_runerror(m, l0, v16+int32(_a2660), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v26 = F_luaD_precall(m, l0, l1, l2)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					if v26 != 0 {
						v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
						v33 = v31 + int32(-1)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v33)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+64))
						if base.Ui32(v36) < base.Ui32(v37) {
							return
						} else {
							F_luaC_step(m, l0)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						F_luaV_execute(m, l0, int32(1))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
							v33 = v31 + int32(-1)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v33)
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+64))
							if base.Ui32(v36) < base.Ui32(v37) {
								return
							} else {
								F_luaC_step(m, l0)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
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
func F_luaD_callhook(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v14 == int32(0) {
		m.G0 = v12 + int32(112)
		return
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)))
		if v17 == int32(0) {
			m.G0 = v12 + int32(112)
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
			if l1 == int32(4) {
				v33 = int32(0)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v32 = base.I32_div_s(v20-v29, int32(24))
				v33 = v32
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v33
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if int32(320) < v37-v23 {
				v53 = v20
				v54 = v23
				*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v54 + int32(320)
				v58 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v58)
				m.T0[v14].(func(*base.Module, int32, int32))(m, l0, v12+int32(12))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					v64 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v64)
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v67 + (v21 - v22)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + (v23 - v22)
					m.G0 = v12 + int32(112)
					return
				}
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v42 = int32(20)
				if v41 < v42 {
					v48 = v41 + v42
				} else {
					v48 = v41 << (uint(int32(1)) % 32)
				}
				F_luaD_reallocstack(m, l0, v48)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v53 = v51
					v54 = v52
					*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v54 + int32(320)
					v58 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v58)
					m.T0[v14].(func(*base.Module, int32, int32))(m, l0, v12+int32(12))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v64 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v64)
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v67 + (v21 - v22)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + (v23 - v22)
						m.G0 = v12 + int32(112)
						return
					}
				}
			}
		}
	}
}
func F_luaD_precall(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
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
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
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
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v355 int32
	_ = v355
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int64
	_ = v372
	var v374 int32
	_ = v374
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int64
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v468 int64
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v505 int64
	_ = v505
	var v507 int32
	_ = v507
	var v528 int32
	_ = v528
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v585 int32
	_ = v585
	var v596 int32
	_ = v596
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v810 int32
	_ = v810
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v20 != int32(6) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = v150
	v152 = v134 - v148
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+6)))
	if v153 != 0 {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	switch v26 + int32(-5) {
	case 0:
		goto L8
	default:
		goto L6
	case 2:
		goto L7
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v134 = l1
	v137 = v23
	goto L1
L4:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v54 == int32(6) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v43 = m.G398
	if v42 == int32(0) {
		v52 = v43
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v41 = v35 + v26<<(uint(int32(2))%32) + int32(152)
	goto L5
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v41 = v32 + int32(8)
	goto L5
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v41 = v29 + int32(16)
	goto L5
L9:
	;
	goto L4
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(64))+188))
	v51 = F_luaH_getstr(m, v42, v50)
	mBase = m.M
	v52 = v51
	goto L9
L11:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v64) <= base.Ui32(l1) {
		v98 = v64
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v57 = m.G3
	F_luaG_typeerror(m, l0, l1, v57+int32(_a2658))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	goto L11
L15:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(16) < v106-v98 {
		v121 = v98
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v73 = v64
	goto L17
L17:
	;
	v82 = v73 + int32(-16)
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
	*(*int64)(unsafe.Add(mBase, uint32(v73))) = v83
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v73+int32(-8))))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v87
	if base.Ui32(l1) < base.Ui32(v82) {
		v73 = v82
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v98 = v90
	goto L15
L19:
	;
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v121 + int32(16)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v127 = v125 + (l1 - v53)
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = v128
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+8)) = v130
	v134 = v127
	v137 = base.I32_wrap_i64(v128)
	goto L1
L21:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v111 = int32(1)
	if v110 < v111 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v117 = v110 + v111
	goto L24
L23:
	;
	v117 = v110 << (uint(v111) % 32)
	goto L24
L24:
	;
	F_luaD_reallocstack(m, l0, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v121 = v120
	goto L20
L26:
	;
	m.G0 = v18 + int32(112)
	return v810
L27:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(320) < v686-v687 {
		v702 = v149
		goto L101
	} else {
		goto L102
	}
L28:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+73)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+75)))
	v160 = v158 + v159
	if v160<<(uint(int32(4))%32) < v154-v155 {
		v173 = v148
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v174 = v173 + v152
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+74)))
	if v175 != 0 {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v164 < v160 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v169 = v164 + v160
	goto L33
L32:
	;
	v169 = v164 << (uint(int32(1)) % 32)
	goto L33
L33:
	;
	F_luaD_reallocstack(m, l0, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v173 = v172
	goto L29
L35:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v553 != v554 {
		goto L82
	} else {
		goto L83
	}
L36:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v188 = (v185 - v174) >> (uint(int32(4)) % 32)
	v190 = v188 + int32(-1)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+73)))
	if v191 < v188 {
		v298 = v190
		v301 = v185
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v178 = v174 + int32(16)
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+73)))
	v182 = v178 + v179<<(uint(int32(4))%32)
	if base.Ui32(v176) <= base.Ui32(v182) {
		v542 = v174
		v545 = v178
		goto L35
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v182
	v542 = v174
	v545 = v178
	goto L35
L39:
	;
	v309 = int32(0)
	if v175&int32(4) == v309 {
		v412 = v301
		v417 = v309
		goto L51
	} else {
		goto L52
	}
L40:
	;
	v193 = v191 - v188
	v197 = (v193 + int32(1)) & int32(7)
	if v197 == int32(0) {
		v229 = v190
		v232 = v185
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if base.Ui32(v193) < base.Ui32(int32(7)) {
		v285 = v232
		goto L46
	} else {
		goto L47
	}
L42:
	;
	v202 = int32(0)
	v205 = v190
	v208 = v185
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = int32(0)
	v218 = int32(1)
	v219 = v205 + v218
	v221 = v208 + int32(16)
	v223 = v202 + v218
	if v223 != v197 {
		v202 = v223
		v205 = v219
		v208 = v221
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v229 = v219
	v232 = v221
	goto L41
L45:
	;
	goto L44
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v285
	v298 = v191
	v301 = v285
	goto L39
L47:
	;
	v246 = v229
	v249 = v232
	goto L48
L48:
	;
	v257 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v249)+120)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v249)+104)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v249)+88)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v249)+72)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v249)+56)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v249)+40)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v249)+24)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v249)+8)) = v257
	v274 = v249 + int32(128)
	v276 = v246 + int32(8)
	if v276 != v191 {
		v246 = v276
		v249 = v274
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v285 = v274
	goto L46
L50:
	;
	goto L49
L51:
	;
	if v191 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L52:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+68))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v314)+64))
	if base.Ui32(v315) < base.Ui32(v316) {
		v321 = v301
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v322 = v298 - v191
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+75)))
	if v325<<(uint(int32(4))%32) < v323-v321 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	F_luaC_step(m, l0)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v321 = v320
	goto L53
L56:
	;
	v339 = F_luaH_new(m, l0, v322, int32(1))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L13
	} else {
		goto L62
	}
L57:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v329 < v325 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v334 = v329 + v325
	goto L60
L59:
	;
	v334 = v329 << (uint(int32(1)) % 32)
	goto L60
L60:
	;
	F_luaD_reallocstack(m, l0, v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	goto L56
L62:
	;
	if v322 < int32(1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v392 = m.G3
	v396 = F_luaS_newlstr(m, l0, v392+int32(_a2659), int32(1))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L13
	} else {
		goto L69
	}
L64:
	;
	v343 = int32(0)
	v355 = v343
	goto L65
L65:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v367 = v355 + int32(1)
	v368 = F_luaH_setnum(m, l0, v339, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L13
	} else {
		goto L67
	}
L66:
	;
	goto L63
L67:
	;
	v371 = v355<<(uint(int32(4))%32) + (v365 + (v343-v322)<<(uint(int32(4))%32))
	v372 = *(*int64)(unsafe.Add(mBase, uint32(v371)))
	*(*int64)(unsafe.Add(mBase, uint32(v368))) = v372
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v371)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v368)+8)) = v374
	if v367 != v322 {
		v355 = v367
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v398 = F_luaH_setstr(m, l0, v339, v396)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v398)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v398))) = base.F64_convert_i32_s(v322)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v412 = v404
	v417 = v339
	goto L51
L71:
	;
	if v417 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	v424 = v412 - v298<<(uint(int32(4))%32)
	v425 = int32(1)
	if v191 == v425 {
		v482 = int32(0)
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if v191&v425 == int32(0) {
		goto L71
	} else {
		goto L78
	}
L74:
	;
	v432 = int32(0)
	v435 = v432
	v440 = v432
	goto L75
L75:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v450 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v449 + v450
	v455 = v424 + v435<<(uint(int32(4))%32)
	v456 = *(*int64)(unsafe.Add(mBase, uint32(v455)))
	*(*int64)(unsafe.Add(mBase, uint32(v449))) = v456
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v455)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v449)+8)) = v458
	v460 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v455)+8)) = v460
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v462 + v450
	v468 = *(*int64)(unsafe.Add(mBase, uint32(v455+v450)))
	*(*int64)(unsafe.Add(mBase, uint32(v462))) = v468
	v471 = v455 + int32(24)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	*(*int32)(unsafe.Add(mBase, uint32(v462)+8)) = v472
	*(*int32)(unsafe.Add(mBase, uint32(v471))) = v460
	v476 = int32(2)
	v477 = v435 + v476
	v479 = v440 + v476
	if v479 != v191&int32(254) {
		v435 = v477
		v440 = v479
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v482 = v477
	goto L73
L77:
	;
	goto L76
L78:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v498 + int32(16)
	v504 = v424 + v482<<(uint(int32(4))%32)
	v505 = *(*int64)(unsafe.Add(mBase, uint32(v504)))
	*(*int64)(unsafe.Add(mBase, uint32(v498))) = v505
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v504)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v498)+8)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v504)+8)) = int32(0)
	goto L71
L79:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v542 = v536 + v152
	v545 = v412
	goto L35
L80:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v528 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v528)+8)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v528))) = v417
	goto L79
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v561))) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v561)+4)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v545
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+75)))
	v568 = v545 + v565<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v561)+8)) = v568
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v561)+16)) = l2
	v573 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v561)+20)) = v573
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v568) <= base.Ui32(v576) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v559 = v553 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v559
	v561 = v559
	goto L81
L83:
	;
	v556 = F_growCI(m, l0)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	v561 = v556
	goto L81
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v568
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v614&int32(1) == int32(0) {
		v810 = v573
		goto L26
	} else {
		goto L90
	}
L86:
	;
	v585 = v576
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v585)+8)) = int32(0)
	v596 = v585 + int32(16)
	if base.Ui32(v596) < base.Ui32(v568) {
		v585 = v596
		goto L87
	} else {
		goto L89
	}
L88:
	;
	goto L85
L89:
	;
	goto L88
L90:
	;
	v620 = v570 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v620
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v622 == int32(0) {
		v680 = v620
		goto L91
	} else {
		goto L92
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v680 + int32(-4)
	v810 = v573
	goto L26
L92:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)))
	if v625 == int32(0) {
		v680 = v620
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)+8))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(0)
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v638 = base.I32_div_s(v628-v635, int32(24))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v638
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(320) < v642-v568 {
		v658 = v568
		v659 = v628
		goto L94
	} else {
		goto L95
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v659)+8)) = v658 + int32(320)
	v663 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v663)
	m.T0[v622].(func(*base.Module, int32, int32))(m, l0, v18+int32(12))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L13
	} else {
		goto L100
	}
L95:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v647 = int32(20)
	if v646 < v647 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v653 = v646 + v647
	goto L98
L97:
	;
	v653 = v646 << (uint(int32(1)) % 32)
	goto L98
L98:
	;
	F_luaD_reallocstack(m, l0, v653)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L13
	} else {
		goto L99
	}
L99:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v658 = v657
	v659 = v656
	goto L94
L100:
	;
	v669 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v669)
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v671)+8)) = v672 + (v629 - v630)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v672 + (v568 - v630)
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v680 = v677
	goto L91
L101:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v702 != v704 {
		goto L108
	} else {
		goto L109
	}
L102:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v692 = int32(20)
	if v691 < v692 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v698 = v691 + v692
	goto L105
L104:
	;
	v698 = v691 << (uint(int32(1)) % 32)
	goto L105
L105:
	;
	F_luaD_reallocstack(m, l0, v698)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L13
	} else {
		goto L106
	}
L106:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v702 = v701
	goto L101
L107:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v713 = v712 + v152
	*(*int32)(unsafe.Add(mBase, uint32(v711)+4)) = v713
	v716 = v713 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v711))) = v716
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v716
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v711)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v711)+8)) = v719 + int32(320)
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v724&int32(1) == int32(0) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v709 = v702 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v709
	v711 = v709
	goto L107
L109:
	;
	v706 = F_growCI(m, l0)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L13
	} else {
		goto L110
	}
L110:
	;
	v711 = v706
	goto L107
L111:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v789)+4))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v790)))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v791)+16))
	v793 = m.T0[v792].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L13
	} else {
		goto L123
	}
L112:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v729 == int32(0) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)))
	if v732 == int32(0) {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v735)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(0)
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v744 = base.I32_div_s(v735-v741, int32(24))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v744
	v748 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(320) < v748-v719 {
		v764 = v719
		v765 = v735
		goto L115
	} else {
		goto L116
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v765)+8)) = v764 + int32(320)
	v769 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v769)
	m.T0[v729].(func(*base.Module, int32, int32))(m, l0, v18+int32(12))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L13
	} else {
		goto L121
	}
L116:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v753 = int32(20)
	if v752 < v753 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v759 = v752 + v753
	goto L119
L118:
	;
	v759 = v752 << (uint(int32(1)) % 32)
	goto L119
L119:
	;
	F_luaD_reallocstack(m, l0, v759)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L13
	} else {
		goto L120
	}
L120:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v764 = v763
	v765 = v762
	goto L115
L121:
	;
	v775 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v775)
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v777)+8)) = v778 + (v736 - v712)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v778 + (v719 - v712)
	goto L111
L122:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v802 = F_luaD_poscall(m, l0, v798-v793<<(uint(int32(4))%32))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L13
	} else {
		goto L125
	}
L123:
	;
	if int32(0) <= v793 {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v810 = int32(2)
	goto L26
L125:
	;
	v810 = int32(1)
	goto L26
}
func F_luaD_protectedparser(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v4
	v15 = m.G5
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v24 = F_luaD_pcall(m, l0, v15+int32(1231), v7+int32(12), v20-v21, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
		v31 = F_luaM_realloc_(m, l0, v28, v29, int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(32)
			return v24
		}
	}
}
func F_luaD_rawrunprotected(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v22 = v14
	v24 = int32(-1)
	v26 = v4
	v27 = v4
	v28 = v4
	goto L1
L1:
	;
	if v24 == int32(1) {
		v50 = v22
		v51 = v26
		v52 = v27
		v53 = v28
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(112)))) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	m.G0 = v14 + int32(16)
	return v82
L3:
	;
	if v53 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v33 = v22 + int32(-176)
	m.G0 = v33
	v35 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v38
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v33
	v42 = v22 + int32(-172)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v14 + int32(12)
	goto L5
L5:
	;
	v50 = v33
	v51 = v22 + int32(-16)
	v52 = v33
	v53 = v35
	goto L3
L6:
	;
	goto L2
L7:
	;
	goto L8
L8:
	;
	m.T0[l1].(func(*base.Module, int32, int32))(m, l0, l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L6
L10:
	;
	v56 = int32(m.ExcTag)
	v57 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v56 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v61 = int32(v57)
	m.G0 = v50
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(4))))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v71 != v14+int32(12) {
		v74 = int32(0)
		goto L16
	} else {
		goto L17
	}
L13:
	;
	m.ExcPending = 1
	goto L19
L14:
	;
	v22 = v50
	v24 = v74
	v26 = v51
	v27 = v52
	v28 = v65
	goto L1
L15:
	;
	if v74 != 0 {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v74 = v73
	goto L16
L18:
	;
	F___wasm_longjmp(m, v66, v65)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
