package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_geohashEncode(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v23 float64
	_ = v23
	var v28 float64
	_ = v28
	var v33 float64
	_ = v33
	var v36 float64
	_ = v36
	var v49 float64
	_ = v49
	var v52 float64
	_ = v52
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 float64
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v110 int64
	_ = v110
	var v115 int64
	_ = v115
	var v120 int64
	_ = v120
	var v125 int64
	_ = v125
	var v130 int64
	_ = v130
	var v135 int64
	_ = v135
	var v147 int32
	_ = v147
	v5 = l4
	v7 = int32(0)
	if l1 == v7 {
		v147 = v7
	} else {
		if l5 == int32(0) {
			v147 = v7
		} else {
			if base.Ui32((v5+int32(-33))&int32(255)) < base.Ui32(int32(224)) {
				v147 = v7
			} else {
				v23 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				if base.F64_ne(v23, float64(0)) != 0 {
					if l0 == int32(0) {
						v147 = v7
					} else {
						v33 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
						if base.F64_ne(v33, float64(0)) != 0 {
							if base.F64_gt(base.F64_abs(l2), float64(180)) != 0 {
								v147 = v7
							} else {
								if base.F64_gt(base.F64_abs(l3), float64(85.05112878)) != 0 {
									v147 = v7
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)) = uint8(v5)
									*(*int64)(unsafe.Add(mBase, uint32(l5))) = int64(0)
									if base.F64_gt(l3, v23) != 0 {
										v147 = v7
									} else {
										v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
										if base.F64_lt(l3, v49) != 0 {
											v147 = v7
										} else {
											if base.F64_gt(l2, v33) != 0 {
												v147 = v7
											} else {
												v52 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
												if base.F64_lt(l2, v52) != 0 {
													v147 = v7
												} else {
													v60 = base.F64_convert_i64_u(int64(1) << (uint(base.I64_extend_i32_u(v5)) % 64))
													v61 = base.F64_mul(base.F64_div(base.F64_sub(l2, v52), base.F64_sub(v33, v52)), v60)
													if base.F64_lt(v61, float64(4.294967296e+09))&base.F64_ge(v61, float64(0)) == int32(0) {
														v71 = int32(0)
													} else {
														v69 = base.I32_trunc_f64_u(v61)
														v71 = v69
													}
													v75 = base.F64_mul(base.F64_div(base.F64_sub(l3, v49), base.F64_sub(v23, v49)), v60)
													if base.F64_lt(v75, float64(4.294967296e+09))&base.F64_ge(v75, float64(0)) == int32(0) {
														v85 = int32(0)
													} else {
														v83 = base.I32_trunc_f64_u(v75)
														v85 = v83
													}
													v87 = base.I64_extend_i32_u(v71)
													v88 = int64(16)
													v91 = int64(281470681808895)
													v92 = (v87<<(uint(v88)%64) | v87) & v91
													v93 = int64(8)
													v96 = int64(71777214294589695)
													v97 = (v92<<(uint(v93)%64) | v92) & v96
													v98 = int64(4)
													v101 = int64(1085102592571150095)
													v102 = (v97<<(uint(v98)%64) | v97) & v101
													v103 = int64(2)
													v106 = int64(3689348814741910323)
													v107 = (v102<<(uint(v103)%64) | v102) & v106
													v110 = int64(1)
													v115 = base.I64_extend_i32_u(v85)
													v120 = (v115<<(uint(v88)%64) | v115) & v91
													v125 = (v120<<(uint(v93)%64) | v120) & v96
													v130 = (v125<<(uint(v98)%64) | v125) & v101
													v135 = (v130<<(uint(v103)%64) | v130) & v106
													*(*int64)(unsafe.Add(mBase, uint32(l5))) = (v107<<(uint(v103)%64)|v107<<(uint(v110)%64))&int64(-6148914691236517206) | (v135<<(uint(v110)%64)|v135)&int64(6148914691236517205)
													v147 = int32(1)
												}
											}
										}
									}
								}
							}
						} else {
							v36 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
							if base.F64_eq(v36, float64(0)) != 0 {
								v147 = v7
							} else {
								if base.F64_gt(base.F64_abs(l2), float64(180)) != 0 {
									v147 = v7
								} else {
									if base.F64_gt(base.F64_abs(l3), float64(85.05112878)) != 0 {
										v147 = v7
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)) = uint8(v5)
										*(*int64)(unsafe.Add(mBase, uint32(l5))) = int64(0)
										if base.F64_gt(l3, v23) != 0 {
											v147 = v7
										} else {
											v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
											if base.F64_lt(l3, v49) != 0 {
												v147 = v7
											} else {
												if base.F64_gt(l2, v33) != 0 {
													v147 = v7
												} else {
													v52 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
													if base.F64_lt(l2, v52) != 0 {
														v147 = v7
													} else {
														v60 = base.F64_convert_i64_u(int64(1) << (uint(base.I64_extend_i32_u(v5)) % 64))
														v61 = base.F64_mul(base.F64_div(base.F64_sub(l2, v52), base.F64_sub(v33, v52)), v60)
														if base.F64_lt(v61, float64(4.294967296e+09))&base.F64_ge(v61, float64(0)) == int32(0) {
															v71 = int32(0)
														} else {
															v69 = base.I32_trunc_f64_u(v61)
															v71 = v69
														}
														v75 = base.F64_mul(base.F64_div(base.F64_sub(l3, v49), base.F64_sub(v23, v49)), v60)
														if base.F64_lt(v75, float64(4.294967296e+09))&base.F64_ge(v75, float64(0)) == int32(0) {
															v85 = int32(0)
														} else {
															v83 = base.I32_trunc_f64_u(v75)
															v85 = v83
														}
														v87 = base.I64_extend_i32_u(v71)
														v88 = int64(16)
														v91 = int64(281470681808895)
														v92 = (v87<<(uint(v88)%64) | v87) & v91
														v93 = int64(8)
														v96 = int64(71777214294589695)
														v97 = (v92<<(uint(v93)%64) | v92) & v96
														v98 = int64(4)
														v101 = int64(1085102592571150095)
														v102 = (v97<<(uint(v98)%64) | v97) & v101
														v103 = int64(2)
														v106 = int64(3689348814741910323)
														v107 = (v102<<(uint(v103)%64) | v102) & v106
														v110 = int64(1)
														v115 = base.I64_extend_i32_u(v85)
														v120 = (v115<<(uint(v88)%64) | v115) & v91
														v125 = (v120<<(uint(v93)%64) | v120) & v96
														v130 = (v125<<(uint(v98)%64) | v125) & v101
														v135 = (v130<<(uint(v103)%64) | v130) & v106
														*(*int64)(unsafe.Add(mBase, uint32(l5))) = (v107<<(uint(v103)%64)|v107<<(uint(v110)%64))&int64(-6148914691236517206) | (v135<<(uint(v110)%64)|v135)&int64(6148914691236517205)
														v147 = int32(1)
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
					if l0 == int32(0) {
						v147 = v7
					} else {
						v28 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						if base.F64_ne(v28, float64(0)) != 0 {
							v33 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
							if base.F64_ne(v33, float64(0)) != 0 {
								if base.F64_gt(base.F64_abs(l2), float64(180)) != 0 {
									v147 = v7
								} else {
									if base.F64_gt(base.F64_abs(l3), float64(85.05112878)) != 0 {
										v147 = v7
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)) = uint8(v5)
										*(*int64)(unsafe.Add(mBase, uint32(l5))) = int64(0)
										if base.F64_gt(l3, v23) != 0 {
											v147 = v7
										} else {
											v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
											if base.F64_lt(l3, v49) != 0 {
												v147 = v7
											} else {
												if base.F64_gt(l2, v33) != 0 {
													v147 = v7
												} else {
													v52 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
													if base.F64_lt(l2, v52) != 0 {
														v147 = v7
													} else {
														v60 = base.F64_convert_i64_u(int64(1) << (uint(base.I64_extend_i32_u(v5)) % 64))
														v61 = base.F64_mul(base.F64_div(base.F64_sub(l2, v52), base.F64_sub(v33, v52)), v60)
														if base.F64_lt(v61, float64(4.294967296e+09))&base.F64_ge(v61, float64(0)) == int32(0) {
															v71 = int32(0)
														} else {
															v69 = base.I32_trunc_f64_u(v61)
															v71 = v69
														}
														v75 = base.F64_mul(base.F64_div(base.F64_sub(l3, v49), base.F64_sub(v23, v49)), v60)
														if base.F64_lt(v75, float64(4.294967296e+09))&base.F64_ge(v75, float64(0)) == int32(0) {
															v85 = int32(0)
														} else {
															v83 = base.I32_trunc_f64_u(v75)
															v85 = v83
														}
														v87 = base.I64_extend_i32_u(v71)
														v88 = int64(16)
														v91 = int64(281470681808895)
														v92 = (v87<<(uint(v88)%64) | v87) & v91
														v93 = int64(8)
														v96 = int64(71777214294589695)
														v97 = (v92<<(uint(v93)%64) | v92) & v96
														v98 = int64(4)
														v101 = int64(1085102592571150095)
														v102 = (v97<<(uint(v98)%64) | v97) & v101
														v103 = int64(2)
														v106 = int64(3689348814741910323)
														v107 = (v102<<(uint(v103)%64) | v102) & v106
														v110 = int64(1)
														v115 = base.I64_extend_i32_u(v85)
														v120 = (v115<<(uint(v88)%64) | v115) & v91
														v125 = (v120<<(uint(v93)%64) | v120) & v96
														v130 = (v125<<(uint(v98)%64) | v125) & v101
														v135 = (v130<<(uint(v103)%64) | v130) & v106
														*(*int64)(unsafe.Add(mBase, uint32(l5))) = (v107<<(uint(v103)%64)|v107<<(uint(v110)%64))&int64(-6148914691236517206) | (v135<<(uint(v110)%64)|v135)&int64(6148914691236517205)
														v147 = int32(1)
													}
												}
											}
										}
									}
								}
							} else {
								v36 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
								if base.F64_eq(v36, float64(0)) != 0 {
									v147 = v7
								} else {
									if base.F64_gt(base.F64_abs(l2), float64(180)) != 0 {
										v147 = v7
									} else {
										if base.F64_gt(base.F64_abs(l3), float64(85.05112878)) != 0 {
											v147 = v7
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)) = uint8(v5)
											*(*int64)(unsafe.Add(mBase, uint32(l5))) = int64(0)
											if base.F64_gt(l3, v23) != 0 {
												v147 = v7
											} else {
												v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
												if base.F64_lt(l3, v49) != 0 {
													v147 = v7
												} else {
													if base.F64_gt(l2, v33) != 0 {
														v147 = v7
													} else {
														v52 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
														if base.F64_lt(l2, v52) != 0 {
															v147 = v7
														} else {
															v60 = base.F64_convert_i64_u(int64(1) << (uint(base.I64_extend_i32_u(v5)) % 64))
															v61 = base.F64_mul(base.F64_div(base.F64_sub(l2, v52), base.F64_sub(v33, v52)), v60)
															if base.F64_lt(v61, float64(4.294967296e+09))&base.F64_ge(v61, float64(0)) == int32(0) {
																v71 = int32(0)
															} else {
																v69 = base.I32_trunc_f64_u(v61)
																v71 = v69
															}
															v75 = base.F64_mul(base.F64_div(base.F64_sub(l3, v49), base.F64_sub(v23, v49)), v60)
															if base.F64_lt(v75, float64(4.294967296e+09))&base.F64_ge(v75, float64(0)) == int32(0) {
																v85 = int32(0)
															} else {
																v83 = base.I32_trunc_f64_u(v75)
																v85 = v83
															}
															v87 = base.I64_extend_i32_u(v71)
															v88 = int64(16)
															v91 = int64(281470681808895)
															v92 = (v87<<(uint(v88)%64) | v87) & v91
															v93 = int64(8)
															v96 = int64(71777214294589695)
															v97 = (v92<<(uint(v93)%64) | v92) & v96
															v98 = int64(4)
															v101 = int64(1085102592571150095)
															v102 = (v97<<(uint(v98)%64) | v97) & v101
															v103 = int64(2)
															v106 = int64(3689348814741910323)
															v107 = (v102<<(uint(v103)%64) | v102) & v106
															v110 = int64(1)
															v115 = base.I64_extend_i32_u(v85)
															v120 = (v115<<(uint(v88)%64) | v115) & v91
															v125 = (v120<<(uint(v93)%64) | v120) & v96
															v130 = (v125<<(uint(v98)%64) | v125) & v101
															v135 = (v130<<(uint(v103)%64) | v130) & v106
															*(*int64)(unsafe.Add(mBase, uint32(l5))) = (v107<<(uint(v103)%64)|v107<<(uint(v110)%64))&int64(-6148914691236517206) | (v135<<(uint(v110)%64)|v135)&int64(6148914691236517205)
															v147 = int32(1)
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v147 = v7
						}
					}
				}
			}
		}
	}
	return v147
}
func F_geohashGetCoordRange(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(-4582834833314545664)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(4640537203540230144)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(-4587686678794764544)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = int64(4635685358060011264)
	return
}
func F_geohashGetDistance(m *base.Module, l0 float64, l1 float64, l2 float64, l3 float64) float64 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v12 float64
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v32 float64
	_ = v32
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v45 float64
	_ = v45
	var v47 float64
	_ = v47
	var v49 float64
	_ = v49
	var v60 float64
	_ = v60
	var v70 float64
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v90 float64
	_ = v90
	var v94 int32
	_ = v94
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v99 float64
	_ = v99
	var v101 float64
	_ = v101
	var v103 float64
	_ = v103
	var v106 float64
	_ = v106
	var v109 float64
	_ = v109
	var v114 float64
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v134 float64
	_ = v134
	var v138 int32
	_ = v138
	var v139 float64
	_ = v139
	var v140 float64
	_ = v140
	var v143 float64
	_ = v143
	var v145 float64
	_ = v145
	var v147 float64
	_ = v147
	var v150 float64
	_ = v150
	var v153 float64
	_ = v153
	var v159 float64
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v179 float64
	_ = v179
	var v183 int32
	_ = v183
	var v184 float64
	_ = v184
	var v185 float64
	_ = v185
	var v189 float64
	_ = v189
	var v190 float64
	_ = v190
	var v192 float64
	_ = v192
	var v194 float64
	_ = v194
	var v196 float64
	_ = v196
	var v207 float64
	_ = v207
	var v213 int64
	_ = v213
	var v218 int32
	_ = v218
	var v239 float64
	_ = v239
	var v243 float64
	_ = v243
	var v246 float64
	_ = v246
	var v247 float64
	_ = v247
	var v248 float64
	_ = v248
	var v253 float64
	_ = v253
	var v258 float64
	_ = v258
	var v262 float64
	_ = v262
	var v271 float64
	_ = v271
	var v278 float64
	_ = v278
	var v283 float64
	_ = v283
	var v284 float64
	_ = v284
	var v292 float64
	_ = v292
	v6 = float64(0.017453292519943295)
	v12 = base.F64_mul(base.F64_sub(base.F64_mul(l2, v6), base.F64_mul(l0, v6)), float64(0.5))
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v25 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v12))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072243195)) < base.Ui32(v25) {
		if base.Ui32(v25) < base.Ui32(int32(2146435072)) {
			v36 = F___rem_pio2(m, v12, v18)
			mBase = m.M
			v37 = *(*float64)(unsafe.Add(mBase, uint32(v18)+8))
			v38 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
			switch v36 & int32(3) {
			default:
				v42 = F___sin(m, v38, v37, int32(1))
				mBase = m.M
				v49 = v42
			case 1:
				v43 = F___cos(m, v38, v37)
				mBase = m.M
				v49 = v43
			case 2:
				v45 = F___sin(m, v38, v37, int32(1))
				mBase = m.M
				v49 = base.F64_neg(v45)
			case 3:
				v47 = F___cos(m, v38, v37)
				mBase = m.M
				v49 = base.F64_neg(v47)
			}
		} else {
			v49 = base.F64_sub(v12, v12)
		}
	} else {
		if base.Ui32(v25) < base.Ui32(int32(1045430272)) {
			v49 = v12
		} else {
			v32 = F___sin(m, v12, float64(0), int32(0))
			mBase = m.M
			v49 = v32
		}
	}
	m.G0 = v18 + int32(16)
	if base.F64_le(base.F64_abs(v49), float64(1e-15)) == int32(0) {
		v70 = base.F64_mul(l1, float64(0.017453292519943295))
		v74 = m.G0
		v76 = v74 - int32(16)
		m.G0 = v76
		v83 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v70))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1072243195)) < base.Ui32(v83) {
			if base.Ui32(v83) < base.Ui32(int32(2146435072)) {
				v94 = F___rem_pio2(m, v70, v76)
				mBase = m.M
				v95 = *(*float64)(unsafe.Add(mBase, uint32(v76)+8))
				v96 = *(*float64)(unsafe.Add(mBase, uint32(v76)))
				switch v94 & int32(3) {
				default:
					v99 = F___cos(m, v96, v95)
					mBase = m.M
					v109 = v99
				case 1:
					v101 = F___sin(m, v96, v95, int32(1))
					mBase = m.M
					v109 = base.F64_neg(v101)
				case 2:
					v103 = F___cos(m, v96, v95)
					mBase = m.M
					v109 = base.F64_neg(v103)
				case 3:
					v106 = F___sin(m, v96, v95, int32(1))
					mBase = m.M
					v109 = v106
				}
			} else {
				v109 = base.F64_sub(v70, v70)
			}
		} else {
			if base.Ui32(v83) < base.Ui32(int32(1044816030)) {
				v109 = float64(1)
			} else {
				v90 = F___cos(m, v70, float64(0))
				mBase = m.M
				v109 = v90
			}
		}
		m.G0 = v76 + int32(16)
		v114 = base.F64_mul(l3, float64(0.017453292519943295))
		v118 = m.G0
		v120 = v118 - int32(16)
		m.G0 = v120
		v127 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v114))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1072243195)) < base.Ui32(v127) {
			if base.Ui32(v127) < base.Ui32(int32(2146435072)) {
				v138 = F___rem_pio2(m, v114, v120)
				mBase = m.M
				v139 = *(*float64)(unsafe.Add(mBase, uint32(v120)+8))
				v140 = *(*float64)(unsafe.Add(mBase, uint32(v120)))
				switch v138 & int32(3) {
				default:
					v143 = F___cos(m, v140, v139)
					mBase = m.M
					v153 = v143
				case 1:
					v145 = F___sin(m, v140, v139, int32(1))
					mBase = m.M
					v153 = base.F64_neg(v145)
				case 2:
					v147 = F___cos(m, v140, v139)
					mBase = m.M
					v153 = base.F64_neg(v147)
				case 3:
					v150 = F___sin(m, v140, v139, int32(1))
					mBase = m.M
					v153 = v150
				}
			} else {
				v153 = base.F64_sub(v114, v114)
			}
		} else {
			if base.Ui32(v127) < base.Ui32(int32(1044816030)) {
				v153 = float64(1)
			} else {
				v134 = F___cos(m, v114, float64(0))
				mBase = m.M
				v153 = v134
			}
		}
		m.G0 = v120 + int32(16)
		v159 = base.F64_mul(base.F64_sub(v114, v70), float64(0.5))
		v163 = m.G0
		v165 = v163 - int32(16)
		m.G0 = v165
		v172 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v159))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1072243195)) < base.Ui32(v172) {
			if base.Ui32(v172) < base.Ui32(int32(2146435072)) {
				v183 = F___rem_pio2(m, v159, v165)
				mBase = m.M
				v184 = *(*float64)(unsafe.Add(mBase, uint32(v165)+8))
				v185 = *(*float64)(unsafe.Add(mBase, uint32(v165)))
				switch v183 & int32(3) {
				default:
					v189 = F___sin(m, v185, v184, int32(1))
					mBase = m.M
					v196 = v189
				case 1:
					v190 = F___cos(m, v185, v184)
					mBase = m.M
					v196 = v190
				case 2:
					v192 = F___sin(m, v185, v184, int32(1))
					mBase = m.M
					v196 = base.F64_neg(v192)
				case 3:
					v194 = F___cos(m, v185, v184)
					mBase = m.M
					v196 = base.F64_neg(v194)
				}
			} else {
				v196 = base.F64_sub(v159, v159)
			}
		} else {
			if base.Ui32(v172) < base.Ui32(int32(1045430272)) {
				v196 = v159
			} else {
				v179 = F___sin(m, v159, float64(0), int32(0))
				mBase = m.M
				v196 = v179
			}
		}
		m.G0 = v165 + int32(16)
		v207 = base.F64_sqrt(base.F64_add(base.F64_mul(v196, v196), base.F64_mul(v49, base.F64_mul(base.F64_mul(v109, v153), v49))))
		v213 = base.I64_reinterpret_f64(v207)
		v218 = base.I32_wrap_i64(int64(base.Ui64(v213)>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v218) < base.Ui32(int32(1072693248)) {
			if base.Ui32(int32(1071644671)) < base.Ui32(v218) {
				v243 = F_fabs(m, v207)
				mBase = m.M
				v246 = base.F64_mul(base.F64_sub(float64(1), v243), float64(0.5))
				v247 = F_sqrt(m, v246)
				mBase = m.M
				v248 = F_R_2(m, v246)
				mBase = m.M
				if base.Ui32(v218) < base.Ui32(int32(1072640819)) {
					v258 = float64(0.7853981633974483)
					v262 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v247) & int64(-4294967296))
					v271 = base.F64_div(base.F64_sub(v246, base.F64_mul(v262, v262)), base.F64_add(v247, v262))
					v278 = base.F64_add(base.F64_sub(base.F64_sub(v258, base.F64_add(v262, v262)), base.F64_sub(base.F64_mul(base.F64_add(v247, v247), v248), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v271, v271)))), v258)
				} else {
					v253 = base.F64_add(base.F64_mul(v247, v248), v247)
					v278 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v253, v253), float64(-6.123233995736766e-17)))
				}
				if v213 < int64(0) {
					v283 = base.F64_neg(v278)
				} else {
					v283 = v278
				}
				v284 = v283
				v292 = v284
			} else {
				if base.Ui32(v218+int32(-1048576)) < base.Ui32(int32(1044381696)) {
					v284 = v207
					v292 = v284
				} else {
					v239 = F_R_2(m, base.F64_mul(v207, v207))
					mBase = m.M
					v292 = base.F64_add(base.F64_mul(v207, v239), v207)
				}
			}
		} else {
			if v218+int32(-1072693248)|base.I32_wrap_i64(v213) != 0 {
				v292 = base.F64_div(float64(0), base.F64_sub(v207, v207))
			} else {
				v292 = base.F64_add(base.F64_mul(v207, float64(1.5707963267948966)), float64(7.52316384526264e-37))
			}
		}
		return base.F64_mul(v292, float64(1.2745595121712e+07))
	} else {
		v60 = float64(0.017453292519943295)
		return base.F64_mul(base.F64_abs(base.F64_sub(base.F64_mul(l3, v60), base.F64_mul(l1, v60))), float64(6.372797560856e+06))
	}
}
func F_geohashGetDistanceIfInRadiusWGS84(m *base.Module, l0 float64, l1 float64, l2 float64, l3 float64, l4 float64, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v14 float64
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v34 float64
	_ = v34
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v47 float64
	_ = v47
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v62 float64
	_ = v62
	var v71 float64
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v91 float64
	_ = v91
	var v95 int32
	_ = v95
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v100 float64
	_ = v100
	var v102 float64
	_ = v102
	var v104 float64
	_ = v104
	var v107 float64
	_ = v107
	var v110 float64
	_ = v110
	var v115 float64
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v135 float64
	_ = v135
	var v139 int32
	_ = v139
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v144 float64
	_ = v144
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v151 float64
	_ = v151
	var v154 float64
	_ = v154
	var v160 float64
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v180 float64
	_ = v180
	var v184 int32
	_ = v184
	var v185 float64
	_ = v185
	var v186 float64
	_ = v186
	var v190 float64
	_ = v190
	var v191 float64
	_ = v191
	var v193 float64
	_ = v193
	var v195 float64
	_ = v195
	var v197 float64
	_ = v197
	var v208 float64
	_ = v208
	var v214 int64
	_ = v214
	var v219 int32
	_ = v219
	var v240 float64
	_ = v240
	var v244 float64
	_ = v244
	var v247 float64
	_ = v247
	var v248 float64
	_ = v248
	var v249 float64
	_ = v249
	var v254 float64
	_ = v254
	var v259 float64
	_ = v259
	var v263 float64
	_ = v263
	var v272 float64
	_ = v272
	var v279 float64
	_ = v279
	var v284 float64
	_ = v284
	var v285 float64
	_ = v285
	var v293 float64
	_ = v293
	var v296 float64
	_ = v296
	v8 = float64(0.017453292519943295)
	v14 = base.F64_mul(base.F64_sub(base.F64_mul(l2, v8), base.F64_mul(l0, v8)), float64(0.5))
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v27 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v14))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072243195)) < base.Ui32(v27) {
		if base.Ui32(v27) < base.Ui32(int32(2146435072)) {
			v38 = F___rem_pio2(m, v14, v20)
			mBase = m.M
			v39 = *(*float64)(unsafe.Add(mBase, uint32(v20)+8))
			v40 = *(*float64)(unsafe.Add(mBase, uint32(v20)))
			switch v38 & int32(3) {
			default:
				v44 = F___sin(m, v40, v39, int32(1))
				mBase = m.M
				v51 = v44
			case 1:
				v45 = F___cos(m, v40, v39)
				mBase = m.M
				v51 = v45
			case 2:
				v47 = F___sin(m, v40, v39, int32(1))
				mBase = m.M
				v51 = base.F64_neg(v47)
			case 3:
				v49 = F___cos(m, v40, v39)
				mBase = m.M
				v51 = base.F64_neg(v49)
			}
		} else {
			v51 = base.F64_sub(v14, v14)
		}
	} else {
		if base.Ui32(v27) < base.Ui32(int32(1045430272)) {
			v51 = v14
		} else {
			v34 = F___sin(m, v14, float64(0), int32(0))
			mBase = m.M
			v51 = v34
		}
	}
	m.G0 = v20 + int32(16)
	if base.F64_le(base.F64_abs(v51), float64(1e-15)) == int32(0) {
		v71 = base.F64_mul(l1, float64(0.017453292519943295))
		v75 = m.G0
		v77 = v75 - int32(16)
		m.G0 = v77
		v84 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v71))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1072243195)) < base.Ui32(v84) {
			if base.Ui32(v84) < base.Ui32(int32(2146435072)) {
				v95 = F___rem_pio2(m, v71, v77)
				mBase = m.M
				v96 = *(*float64)(unsafe.Add(mBase, uint32(v77)+8))
				v97 = *(*float64)(unsafe.Add(mBase, uint32(v77)))
				switch v95 & int32(3) {
				default:
					v100 = F___cos(m, v97, v96)
					mBase = m.M
					v110 = v100
				case 1:
					v102 = F___sin(m, v97, v96, int32(1))
					mBase = m.M
					v110 = base.F64_neg(v102)
				case 2:
					v104 = F___cos(m, v97, v96)
					mBase = m.M
					v110 = base.F64_neg(v104)
				case 3:
					v107 = F___sin(m, v97, v96, int32(1))
					mBase = m.M
					v110 = v107
				}
			} else {
				v110 = base.F64_sub(v71, v71)
			}
		} else {
			if base.Ui32(v84) < base.Ui32(int32(1044816030)) {
				v110 = float64(1)
			} else {
				v91 = F___cos(m, v71, float64(0))
				mBase = m.M
				v110 = v91
			}
		}
		m.G0 = v77 + int32(16)
		v115 = base.F64_mul(l3, float64(0.017453292519943295))
		v119 = m.G0
		v121 = v119 - int32(16)
		m.G0 = v121
		v128 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v115))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1072243195)) < base.Ui32(v128) {
			if base.Ui32(v128) < base.Ui32(int32(2146435072)) {
				v139 = F___rem_pio2(m, v115, v121)
				mBase = m.M
				v140 = *(*float64)(unsafe.Add(mBase, uint32(v121)+8))
				v141 = *(*float64)(unsafe.Add(mBase, uint32(v121)))
				switch v139 & int32(3) {
				default:
					v144 = F___cos(m, v141, v140)
					mBase = m.M
					v154 = v144
				case 1:
					v146 = F___sin(m, v141, v140, int32(1))
					mBase = m.M
					v154 = base.F64_neg(v146)
				case 2:
					v148 = F___cos(m, v141, v140)
					mBase = m.M
					v154 = base.F64_neg(v148)
				case 3:
					v151 = F___sin(m, v141, v140, int32(1))
					mBase = m.M
					v154 = v151
				}
			} else {
				v154 = base.F64_sub(v115, v115)
			}
		} else {
			if base.Ui32(v128) < base.Ui32(int32(1044816030)) {
				v154 = float64(1)
			} else {
				v135 = F___cos(m, v115, float64(0))
				mBase = m.M
				v154 = v135
			}
		}
		m.G0 = v121 + int32(16)
		v160 = base.F64_mul(base.F64_sub(v115, v71), float64(0.5))
		v164 = m.G0
		v166 = v164 - int32(16)
		m.G0 = v166
		v173 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v160))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1072243195)) < base.Ui32(v173) {
			if base.Ui32(v173) < base.Ui32(int32(2146435072)) {
				v184 = F___rem_pio2(m, v160, v166)
				mBase = m.M
				v185 = *(*float64)(unsafe.Add(mBase, uint32(v166)+8))
				v186 = *(*float64)(unsafe.Add(mBase, uint32(v166)))
				switch v184 & int32(3) {
				default:
					v190 = F___sin(m, v186, v185, int32(1))
					mBase = m.M
					v197 = v190
				case 1:
					v191 = F___cos(m, v186, v185)
					mBase = m.M
					v197 = v191
				case 2:
					v193 = F___sin(m, v186, v185, int32(1))
					mBase = m.M
					v197 = base.F64_neg(v193)
				case 3:
					v195 = F___cos(m, v186, v185)
					mBase = m.M
					v197 = base.F64_neg(v195)
				}
			} else {
				v197 = base.F64_sub(v160, v160)
			}
		} else {
			if base.Ui32(v173) < base.Ui32(int32(1045430272)) {
				v197 = v160
			} else {
				v180 = F___sin(m, v160, float64(0), int32(0))
				mBase = m.M
				v197 = v180
			}
		}
		m.G0 = v166 + int32(16)
		v208 = base.F64_sqrt(base.F64_add(base.F64_mul(v197, v197), base.F64_mul(v51, base.F64_mul(base.F64_mul(v110, v154), v51))))
		v214 = base.I64_reinterpret_f64(v208)
		v219 = base.I32_wrap_i64(int64(base.Ui64(v214)>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v219) < base.Ui32(int32(1072693248)) {
			if base.Ui32(int32(1071644671)) < base.Ui32(v219) {
				v244 = F_fabs(m, v208)
				mBase = m.M
				v247 = base.F64_mul(base.F64_sub(float64(1), v244), float64(0.5))
				v248 = F_sqrt(m, v247)
				mBase = m.M
				v249 = F_R_2(m, v247)
				mBase = m.M
				if base.Ui32(v219) < base.Ui32(int32(1072640819)) {
					v259 = float64(0.7853981633974483)
					v263 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v248) & int64(-4294967296))
					v272 = base.F64_div(base.F64_sub(v247, base.F64_mul(v263, v263)), base.F64_add(v248, v263))
					v279 = base.F64_add(base.F64_sub(base.F64_sub(v259, base.F64_add(v263, v263)), base.F64_sub(base.F64_mul(base.F64_add(v248, v248), v249), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v272, v272)))), v259)
				} else {
					v254 = base.F64_add(base.F64_mul(v248, v249), v248)
					v279 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v254, v254), float64(-6.123233995736766e-17)))
				}
				if v214 < int64(0) {
					v284 = base.F64_neg(v279)
				} else {
					v284 = v279
				}
				v285 = v284
				v293 = v285
			} else {
				if base.Ui32(v219+int32(-1048576)) < base.Ui32(int32(1044381696)) {
					v285 = v208
					v293 = v285
				} else {
					v240 = F_R_2(m, base.F64_mul(v208, v208))
					mBase = m.M
					v293 = base.F64_add(base.F64_mul(v208, v240), v208)
				}
			}
		} else {
			if v219+int32(-1072693248)|base.I32_wrap_i64(v214) != 0 {
				v293 = base.F64_div(float64(0), base.F64_sub(v208, v208))
			} else {
				v293 = base.F64_add(base.F64_mul(v208, float64(1.5707963267948966)), float64(7.52316384526264e-37))
			}
		}
		v296 = base.F64_mul(v293, float64(1.2745595121712e+07))
	} else {
		v62 = float64(0.017453292519943295)
		v296 = base.F64_mul(base.F64_abs(base.F64_sub(base.F64_mul(l3, v62), base.F64_mul(l1, v62))), float64(6.372797560856e+06))
	}
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v296
	return base.F64_gt(v296, l4) ^ int32(1)
}
func F_geohashNeighbors(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v89 int64
	_ = v89
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v112 int64
	_ = v112
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v129 int64
	_ = v129
	var v134 int64
	_ = v134
	var v139 int64
	_ = v139
	var v153 int64
	_ = v153
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v168 int64
	_ = v168
	var v178 int64
	_ = v178
	var v183 int64
	_ = v183
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v194 int64
	_ = v194
	var v206 int64
	_ = v206
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v213 int64
	_ = v213
	var v231 int64
	_ = v231
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v15
	v18 = l1 + int32(24)
	v19 = int32(8)
	v20 = l0 + v19
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v21
	v24 = l1 + int32(40)
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v25
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = v27
	v30 = l1 + v19
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v33
	v36 = l1 + int32(56)
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37
	v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = v39
	v42 = l1 + int32(88)
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v43
	v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+80)) = v45
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v47
	v50 = l1 + int32(120)
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = v51
	v53 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+64)) = v53
	v56 = l1 + int32(72)
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v56))) = v57
	v60 = l1 + int32(104)
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v60))) = v61
	v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+96)) = v63
	v65 = int64(6148914691236517205)
	v66 = int64(64)
	v67 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v68 = int64(1)
	v71 = int64(4294967294)
	v72 = (v66 - v67<<(uint(v68)%64)) & v71
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	v75 = int64(-6148914691236517206)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = (int64(base.Ui64(v65)>>(uint(v72)%64))|v74&v75+v68)&int64(base.Ui64(v75)>>(uint(v72)%64)) | v74&v65
	v89 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v94 = (v66 - v89<<(uint(v68)%64)) & v71
	v96 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = (int64(base.Ui64(v65)>>(uint(v94)%64))|v96&v75+v75>>(uint(v94)%64))&int64(base.Ui64(v75)>>(uint(v94)%64)) | v96&v65
	v112 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v119 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	v122 = int64(9223372036854775807)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(base.Ui64(v65)>>(uint((v66-v112<<(uint(v68)%64))&v71)%64))&(v119&v65+v122) | v119&v75
	v129 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v134 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v139 = (v66 - v134<<(uint(v68)%64)) & v71
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = (v129&v65|int64(base.Ui64(v75)>>(uint(v139)%64))+v68)&int64(base.Ui64(v65)>>(uint(v139)%64)) | v129&v75
	v153 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	v158 = (v66 - v153<<(uint(v68)%64)) & v71
	v159 = int64(base.Ui64(v65) >> (uint(v158) % 64))
	v168 = int64(base.Ui64(v75) >> (uint(v158) % 64))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+96)) = (v159|v63&v75+v75>>(uint(v158)%64))&v168 | (v63&v65|v168+v68)&v159
	v178 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	v183 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	v188 = (v66 - v183<<(uint(v68)%64)) & v71
	v189 = int64(base.Ui64(v75) >> (uint(v188) % 64))
	v194 = int64(base.Ui64(v65) >> (uint(v188) % 64))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+64)) = (v178&v65|v189+v68)&v194 | (v194|v178&v75+v68)&v189
	v206 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v211 = (v66 - v206<<(uint(v68)%64)) & v71
	v212 = int64(base.Ui64(v65) >> (uint(v211) % 64))
	v213 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+80)) = (v212|v213&v75+v68)&int64(base.Ui64(v75)>>(uint(v211)%64)) | v212&(v213&v65+v122)
	v231 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v236 = (v66 - v231<<(uint(v68)%64)) & v71
	v237 = int64(base.Ui64(v65) >> (uint(v236) % 64))
	v238 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = (v237|v238&v75+v75>>(uint(v236)%64))&int64(base.Ui64(v75)>>(uint(v236)%64)) | v237&(v238&v65+v122)
	return
}
