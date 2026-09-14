package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_Arith(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 float64
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v74 float64
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	switch v15 + int32(-3) {
	case 0:
		v29 = l2
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
		switch v30 + int32(-3) {
		case 0:
			v44 = l3
			v45 = *(*float64)(unsafe.Add(mBase, uint32(v44)))
			v46 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
			switch l4 + int32(-5) {
			default:
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_add(v46, v45)
			case 1:
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_sub(v46, v45)
			case 2:
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_mul(v46, v45)
			case 3:
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_div(v46, v45)
			case 4:
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_sub(v46, base.F64_mul(base.F64_floor(base.F64_div(v46, v45)), v45))
			case 5:
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
				v74 = F_pow(m, v46, v45)
				mBase = m.M
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = v74
			case 6:
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_neg(v46)
			}
			m.G0 = v13 + int32(16)
			return
		case 1:
			v34 = v13 + int32(8)
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			v40 = F_luaO_str2d(m, v35+int32(16), v34)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				if v40 == int32(0) {
					v82 = F_call_binTM(m, l0, l2, l3, l1, l4)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						if v82 != 0 {
							m.G0 = v13 + int32(16)
							return
						} else {
							F_luaG_aritherror(m, l0, l2, l3)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								m.G0 = v13 + int32(16)
								return
							}
						}
					}
				} else {
					v44 = v34
					v45 = *(*float64)(unsafe.Add(mBase, uint32(v44)))
					v46 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
					switch l4 + int32(-5) {
					default:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_add(v46, v45)
					case 1:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_sub(v46, v45)
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_mul(v46, v45)
					case 3:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_div(v46, v45)
					case 4:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_sub(v46, base.F64_mul(base.F64_floor(base.F64_div(v46, v45)), v45))
					case 5:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						v74 = F_pow(m, v46, v45)
						mBase = m.M
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = v74
					case 6:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_neg(v46)
					}
					m.G0 = v13 + int32(16)
					return
				}
			}
		default:
			v82 = F_call_binTM(m, l0, l2, l3, l1, l4)
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				if v82 != 0 {
					m.G0 = v13 + int32(16)
					return
				} else {
					F_luaG_aritherror(m, l0, l2, l3)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						m.G0 = v13 + int32(16)
						return
					}
				}
			}
		}
	case 1:
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v23 = F_luaO_str2d(m, v18+int32(16), v13+int32(8))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if v23 == int32(0) {
				v82 = F_call_binTM(m, l0, l2, l3, l1, l4)
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return
				} else {
					if v82 != 0 {
						m.G0 = v13 + int32(16)
						return
					} else {
						F_luaG_aritherror(m, l0, l2, l3)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							m.G0 = v13 + int32(16)
							return
						}
					}
				}
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
				*(*float64)(unsafe.Add(mBase, uint32(v13))) = v27
				v29 = v13
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				switch v30 + int32(-3) {
				case 0:
					v44 = l3
					v45 = *(*float64)(unsafe.Add(mBase, uint32(v44)))
					v46 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
					switch l4 + int32(-5) {
					default:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_add(v46, v45)
					case 1:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_sub(v46, v45)
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_mul(v46, v45)
					case 3:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_div(v46, v45)
					case 4:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_sub(v46, base.F64_mul(base.F64_floor(base.F64_div(v46, v45)), v45))
					case 5:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						v74 = F_pow(m, v46, v45)
						mBase = m.M
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = v74
					case 6:
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_neg(v46)
					}
					m.G0 = v13 + int32(16)
					return
				case 1:
					v34 = v13 + int32(8)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					v40 = F_luaO_str2d(m, v35+int32(16), v34)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						if v40 == int32(0) {
							v82 = F_call_binTM(m, l0, l2, l3, l1, l4)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								if v82 != 0 {
									m.G0 = v13 + int32(16)
									return
								} else {
									F_luaG_aritherror(m, l0, l2, l3)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										m.G0 = v13 + int32(16)
										return
									}
								}
							}
						} else {
							v44 = v34
							v45 = *(*float64)(unsafe.Add(mBase, uint32(v44)))
							v46 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
							switch l4 + int32(-5) {
							default:
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_add(v46, v45)
							case 1:
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_sub(v46, v45)
							case 2:
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_mul(v46, v45)
							case 3:
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_div(v46, v45)
							case 4:
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_sub(v46, base.F64_mul(base.F64_floor(base.F64_div(v46, v45)), v45))
							case 5:
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
								v74 = F_pow(m, v46, v45)
								mBase = m.M
								*(*float64)(unsafe.Add(mBase, uint32(l1))) = v74
							case 6:
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_neg(v46)
							}
							m.G0 = v13 + int32(16)
							return
						}
					}
				default:
					v82 = F_call_binTM(m, l0, l2, l3, l1, l4)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						if v82 != 0 {
							m.G0 = v13 + int32(16)
							return
						} else {
							F_luaG_aritherror(m, l0, l2, l3)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								m.G0 = v13 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	default:
		v82 = F_call_binTM(m, l0, l2, l3, l1, l4)
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return
		} else {
			if v82 != 0 {
				m.G0 = v13 + int32(16)
				return
			} else {
				F_luaG_aritherror(m, l0, l2, l3)
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					m.G0 = v13 + int32(16)
					return
				}
			}
		}
	}
}
func F___addtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v140 int64
	_ = v140
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v160 int64
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int64
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v190 int64
	_ = v190
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v208 int32
	_ = v208
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v215 int64
	_ = v215
	var v222 int64
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v250 int64
	_ = v250
	var v254 int64
	_ = v254
	var v255 int64
	_ = v255
	var v260 int32
	_ = v260
	var v277 int64
	_ = v277
	var v281 int64
	_ = v281
	var v282 int64
	_ = v282
	var v286 int64
	_ = v286
	var v287 int64
	_ = v287
	var v292 int64
	_ = v292
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v304 int64
	_ = v304
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v312 int64
	_ = v312
	var v319 int64
	_ = v319
	var v323 int64
	_ = v323
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int64
	_ = v330
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v355 int64
	_ = v355
	var v359 int64
	_ = v359
	var v360 int64
	_ = v360
	var v367 int64
	_ = v367
	var v368 int64
	_ = v368
	var v370 int64
	_ = v370
	var v373 int64
	_ = v373
	var v378 int64
	_ = v378
	var v390 int64
	_ = v390
	var v392 int64
	_ = v392
	var v393 int32
	_ = v393
	var v396 int64
	_ = v396
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v425 int64
	_ = v425
	var v429 int64
	_ = v429
	var v430 int64
	_ = v430
	var v435 int32
	_ = v435
	var v452 int64
	_ = v452
	var v456 int64
	_ = v456
	var v457 int64
	_ = v457
	var v461 int64
	_ = v461
	var v462 int64
	_ = v462
	var v467 int64
	_ = v467
	var v475 int64
	_ = v475
	var v476 int64
	_ = v476
	var v477 int64
	_ = v477
	var v478 int32
	_ = v478
	var v479 int64
	_ = v479
	var v483 int64
	_ = v483
	var v492 int64
	_ = v492
	var v495 int32
	_ = v495
	var v502 int64
	_ = v502
	var v508 int64
	_ = v508
	var v518 int64
	_ = v518
	var v528 int64
	_ = v528
	var v532 int64
	_ = v532
	var v533 int64
	_ = v533
	var v537 int64
	_ = v537
	var v538 int64
	_ = v538
	var v543 int64
	_ = v543
	var v544 int64
	_ = v544
	v7 = int64(0)
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	v18 = int64(9223372036854775807)
	v19 = l4 & v18
	v21 = base.B2i32(l1 == v7)
	v23 = l2 & v18
	v24 = int64(-9223090561878065152)
	if v23 == v7 {
		v30 = v21
	} else {
		v30 = base.B2i32(base.Ui64(v23+v24) < base.Ui64(v24))
	}
	if v30 != 0 {
		v41 = int64(9223090561878065152)
		if v23 == v41 {
			v45 = v21
		} else {
			v45 = base.B2i32(base.Ui64(v23) < base.Ui64(v41))
		}
		if v45 != 0 {
			v50 = int64(9223090561878065152)
			if v19 == v50 {
				v54 = base.B2i32(l3 == int64(0))
			} else {
				v54 = base.B2i32(base.Ui64(v19) < base.Ui64(v50))
			}
			if v54 != 0 {
				if l1|(v23^int64(9223090561878065152)) != int64(0) {
					if l3|(v19^int64(9223090561878065152)) == int64(0) {
						v543 = l3
						v544 = l4
					} else {
						if l1|v23 != int64(0) {
							if base.B2i32(l3|v19 == int64(0)) == int32(0) {
								if v19 == v23 {
									v95 = base.B2i32(base.Ui64(l1) < base.Ui64(l3))
								} else {
									v95 = base.B2i32(base.Ui64(v23) < base.Ui64(v19))
								}
								if v95 != 0 {
									v96 = l3
								} else {
									v96 = l1
								}
								if v95 != 0 {
									v97 = l4
								} else {
									v97 = l2
								}
								v99 = v97 & int64(281474976710655)
								if v95 != 0 {
									v100 = l2
								} else {
									v100 = l4
								}
								v101 = int64(48)
								v104 = int32(32767)
								v105 = base.I32_wrap_i64(int64(base.Ui64(v100)>>(uint(v101)%64))) & v104
								v110 = base.I32_wrap_i64(int64(base.Ui64(v97)>>(uint(v101)%64))) & v104
								if v110 != 0 {
									v155 = v96
									v156 = v110
									v157 = v99
								} else {
									v112 = v16 + int32(96)
									v114 = base.B2i32(v99 == int64(0))
									if v99 == int64(0) {
										v115 = v96
									} else {
										v115 = v99
									}
									v121 = base.I32_wrap_i64(base.I64_clz(v115) + base.I64_extend_i32_u(v114<<(uint(int32(6))%32)))
									v123 = v121 + int32(-15)
									if v123&int32(64) == int32(0) {
										if v123 == int32(0) {
											v144 = v96
											v145 = v99
										} else {
											v140 = base.I64_extend_i32_u(v123)
											v144 = v96 << (uint(v140) % 64)
											v145 = int64(base.Ui64(v96)>>(uint(base.I64_extend_i32_u(int32(64)-v123))%64)) | v99<<(uint(v140)%64)
										}
									} else {
										v144 = int64(0)
										v145 = v96 << (uint(base.I64_extend_i32_u(v121+int32(-79))) % 64)
									}
									*(*int64)(unsafe.Add(mBase, uint32(v112))) = v144
									*(*int64)(unsafe.Add(mBase, uint32(v112)+8)) = v145
									v153 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(104))))
									v154 = *(*int64)(unsafe.Add(mBase, uint32(v16)+96))
									v155 = v154
									v156 = int32(16) - v121
									v157 = v153
								}
								if v95 != 0 {
									v158 = l1
								} else {
									v158 = l3
								}
								v160 = v100 & int64(281474976710655)
								if v105 != 0 {
									v205 = v160
									v206 = v158
									v208 = v105
								} else {
									v162 = v16 + int32(80)
									v164 = base.B2i32(v160 == int64(0))
									if v160 == int64(0) {
										v165 = v158
									} else {
										v165 = v160
									}
									v171 = base.I32_wrap_i64(base.I64_clz(v165) + base.I64_extend_i32_u(v164<<(uint(int32(6))%32)))
									v173 = v171 + int32(-15)
									if v173&int32(64) == int32(0) {
										if v173 == int32(0) {
											v194 = v158
											v195 = v160
										} else {
											v190 = base.I64_extend_i32_u(v173)
											v194 = v158 << (uint(v190) % 64)
											v195 = int64(base.Ui64(v158)>>(uint(base.I64_extend_i32_u(int32(64)-v173))%64)) | v160<<(uint(v190)%64)
										}
									} else {
										v194 = int64(0)
										v195 = v158 << (uint(base.I64_extend_i32_u(v171+int32(-79))) % 64)
									}
									*(*int64)(unsafe.Add(mBase, uint32(v162))) = v194
									*(*int64)(unsafe.Add(mBase, uint32(v162)+8)) = v195
									v203 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(88))))
									v204 = *(*int64)(unsafe.Add(mBase, uint32(v16)+80))
									v205 = v203
									v206 = v204
									v208 = int32(16) - v171
								}
								v209 = int64(3)
								v211 = int64(61)
								v215 = v205<<(uint(v209)%64) | int64(base.Ui64(v206)>>(uint(v211)%64)) | int64(2251799813685248)
								v222 = v206 << (uint(v209) % 64)
								if v156 == v208 {
									v303 = v215
									v304 = v222
								} else {
									v225 = v156 - v208
									if base.Ui32(v225) <= base.Ui32(int32(127)) {
										v230 = int32(64)
										v231 = v16 + v230
										v233 = int32(128) - v225
										if v233&v230 == int32(0) {
											if v233 == int32(0) {
												v254 = v222
												v255 = v215
											} else {
												v250 = base.I64_extend_i32_u(v233)
												v254 = v222 << (uint(v250) % 64)
												v255 = int64(base.Ui64(v222)>>(uint(base.I64_extend_i32_u(int32(64)-v233))%64)) | v215<<(uint(v250)%64)
											}
										} else {
											v254 = int64(0)
											v255 = v222 << (uint(base.I64_extend_i32_u(v233+int32(-64))) % 64)
										}
										*(*int64)(unsafe.Add(mBase, uint32(v231))) = v254
										*(*int64)(unsafe.Add(mBase, uint32(v231)+8)) = v255
										v260 = v16 + int32(48)
										if v225&int32(64) == int32(0) {
											if v225 == int32(0) {
												v281 = v222
												v282 = v215
											} else {
												v277 = base.I64_extend_i32_u(v225)
												v281 = v215<<(uint(base.I64_extend_i32_u(int32(64)-v225))%64) | int64(base.Ui64(v222)>>(uint(v277)%64))
												v282 = int64(base.Ui64(v215) >> (uint(v277) % 64))
											}
										} else {
											v281 = int64(base.Ui64(v215) >> (uint(base.I64_extend_i32_u(v225+int32(-64))) % 64))
											v282 = int64(0)
										}
										*(*int64)(unsafe.Add(mBase, uint32(v260))) = v281
										*(*int64)(unsafe.Add(mBase, uint32(v260)+8)) = v282
										v286 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
										v287 = *(*int64)(unsafe.Add(mBase, uint32(v16)+64))
										v292 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(72))))
										v302 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(56))))
										v303 = v302
										v304 = v286 | base.I64_extend_i32_u(base.B2i32(v287|v292 != int64(0)))
									} else {
										v303 = int64(0)
										v304 = int64(1)
									}
								}
								v307 = v157<<(uint(v209)%64) | int64(base.Ui64(v155)>>(uint(v211)%64)) | int64(2251799813685248)
								v309 = v155 << (uint(int64(3)) % 64)
								if int64(-1) < l4^l2 {
									v370 = v304 + v309
									v373 = v303 + v307 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v370) < base.Ui64(v304)))
									if v373&int64(4503599627370496) == int64(0) {
										v390 = v370
										v392 = v373
										v393 = v156
									} else {
										v378 = int64(1)
										v390 = int64(base.Ui64(v370)>>(uint(v378)%64)) | v373<<(uint(int64(63))%64) | v304&v378
										v392 = int64(base.Ui64(v373) >> (uint(v378) % 64))
										v393 = v156 + int32(1)
									}
									v396 = v97 & int64(-9223372036854775807-1)
									if v393 < int32(32767) {
										v402 = int32(0)
										if v393 <= v402 {
											v406 = v16 + int32(16)
											v408 = v393 + int32(127)
											if v408&int32(64) == int32(0) {
												if v408 == int32(0) {
													v429 = v390
													v430 = v392
												} else {
													v425 = base.I64_extend_i32_u(v408)
													v429 = v390 << (uint(v425) % 64)
													v430 = int64(base.Ui64(v390)>>(uint(base.I64_extend_i32_u(int32(64)-v408))%64)) | v392<<(uint(v425)%64)
												}
											} else {
												v429 = int64(0)
												v430 = v390 << (uint(base.I64_extend_i32_u(v393+int32(63))) % 64)
											}
											*(*int64)(unsafe.Add(mBase, uint32(v406))) = v429
											*(*int64)(unsafe.Add(mBase, uint32(v406)+8)) = v430
											v435 = int32(1) - v393
											if v435&int32(64) == int32(0) {
												if v435 == int32(0) {
													v456 = v390
													v457 = v392
												} else {
													v452 = base.I64_extend_i32_u(v435)
													v456 = v392<<(uint(base.I64_extend_i32_u(int32(64)-v435))%64) | int64(base.Ui64(v390)>>(uint(v452)%64))
													v457 = int64(base.Ui64(v392) >> (uint(v452) % 64))
												}
											} else {
												v456 = int64(base.Ui64(v392) >> (uint(base.I64_extend_i32_u(v435+int32(-64))) % 64))
												v457 = int64(0)
											}
											*(*int64)(unsafe.Add(mBase, uint32(v16))) = v456
											*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v457
											v461 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
											v462 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
											v467 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(24))))
											v475 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(8))))
											v476 = v461 | base.I64_extend_i32_u(base.B2i32(v462|v467 != int64(0)))
											v477 = v475
											v478 = v402
										} else {
											v476 = v390
											v477 = v392
											v478 = v393
										}
										v479 = int64(3)
										v483 = int64(base.Ui64(v476)>>(uint(v479)%64)) | v477<<(uint(int64(61))%64)
										v492 = base.I64_extend_i32_u(v478)<<(uint(int64(48))%64) | int64(base.Ui64(v477)>>(uint(v479)%64))&int64(281474976710655) | v396
										v495 = base.I32_wrap_i64(v476) & int32(7)
										switch int32(0) {
										case 0:
											if v495 == int32(4) {
												v508 = v483 + v483&int64(1)
												v537 = v508
												v538 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v508) < base.Ui64(v483)))
												v543 = v537
												v544 = v538
											} else {
												v502 = v483 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v495)))
												v532 = v502
												v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v502) < base.Ui64(v483)))
												if v495 == int32(0) {
													v543 = v532
													v544 = v533
												} else {
													v537 = v532
													v538 = v533
													v543 = v537
													v544 = v538
												}
											}
										case 1:
											v518 = v483 + base.I64_extend_i32_u(base.B2i32(v396 != int64(0))&base.B2i32(v495 != int32(0)))
											v532 = v518
											v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v518) < base.Ui64(v483)))
											if v495 == int32(0) {
												v543 = v532
												v544 = v533
											} else {
												v537 = v532
												v538 = v533
												v543 = v537
												v544 = v538
											}
										case 2:
											v528 = v483 + base.I64_extend_i32_u(base.B2i32(v396 == int64(0))&base.B2i32(v495 != int32(0)))
											v532 = v528
											v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v528) < base.Ui64(v483)))
											if v495 == int32(0) {
												v543 = v532
												v544 = v533
											} else {
												v537 = v532
												v538 = v533
												v543 = v537
												v544 = v538
											}
										default:
											v532 = v483
											v533 = v492
											if v495 == int32(0) {
												v543 = v532
												v544 = v533
											} else {
												v537 = v532
												v538 = v533
												v543 = v537
												v544 = v538
											}
										}
									} else {
										v543 = int64(0)
										v544 = v396 | int64(9223090561878065152)
									}
								} else {
									v312 = int64(0)
									if v309^v304|(v307^v303) == v312 {
										v543 = v312
										v544 = v312
									} else {
										v319 = v309 - v304
										v323 = v307 - v303 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v309) < base.Ui64(v304)))
										if base.Ui64(int64(2251799813685247)) < base.Ui64(v323) {
											v390 = v319
											v392 = v323
											v393 = v156
										} else {
											v327 = v16 + int32(32)
											v329 = base.B2i32(v323 == int64(0))
											if v323 == int64(0) {
												v330 = v319
											} else {
												v330 = v323
											}
											v336 = base.I32_wrap_i64(base.I64_clz(v330) + base.I64_extend_i32_u(v329<<(uint(int32(6))%32)))
											v338 = v336 + int32(-12)
											if v338&int32(64) == int32(0) {
												if v338 == int32(0) {
													v359 = v319
													v360 = v323
												} else {
													v355 = base.I64_extend_i32_u(v338)
													v359 = v319 << (uint(v355) % 64)
													v360 = int64(base.Ui64(v319)>>(uint(base.I64_extend_i32_u(int32(64)-v338))%64)) | v323<<(uint(v355)%64)
												}
											} else {
												v359 = int64(0)
												v360 = v319 << (uint(base.I64_extend_i32_u(v336+int32(-76))) % 64)
											}
											*(*int64)(unsafe.Add(mBase, uint32(v327))) = v359
											*(*int64)(unsafe.Add(mBase, uint32(v327)+8)) = v360
											v367 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(40))))
											v368 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
											v390 = v368
											v392 = v367
											v393 = v156 - v338
										}
										v396 = v97 & int64(-9223372036854775807-1)
										if v393 < int32(32767) {
											v402 = int32(0)
											if v393 <= v402 {
												v406 = v16 + int32(16)
												v408 = v393 + int32(127)
												if v408&int32(64) == int32(0) {
													if v408 == int32(0) {
														v429 = v390
														v430 = v392
													} else {
														v425 = base.I64_extend_i32_u(v408)
														v429 = v390 << (uint(v425) % 64)
														v430 = int64(base.Ui64(v390)>>(uint(base.I64_extend_i32_u(int32(64)-v408))%64)) | v392<<(uint(v425)%64)
													}
												} else {
													v429 = int64(0)
													v430 = v390 << (uint(base.I64_extend_i32_u(v393+int32(63))) % 64)
												}
												*(*int64)(unsafe.Add(mBase, uint32(v406))) = v429
												*(*int64)(unsafe.Add(mBase, uint32(v406)+8)) = v430
												v435 = int32(1) - v393
												if v435&int32(64) == int32(0) {
													if v435 == int32(0) {
														v456 = v390
														v457 = v392
													} else {
														v452 = base.I64_extend_i32_u(v435)
														v456 = v392<<(uint(base.I64_extend_i32_u(int32(64)-v435))%64) | int64(base.Ui64(v390)>>(uint(v452)%64))
														v457 = int64(base.Ui64(v392) >> (uint(v452) % 64))
													}
												} else {
													v456 = int64(base.Ui64(v392) >> (uint(base.I64_extend_i32_u(v435+int32(-64))) % 64))
													v457 = int64(0)
												}
												*(*int64)(unsafe.Add(mBase, uint32(v16))) = v456
												*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v457
												v461 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
												v462 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
												v467 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(24))))
												v475 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(8))))
												v476 = v461 | base.I64_extend_i32_u(base.B2i32(v462|v467 != int64(0)))
												v477 = v475
												v478 = v402
											} else {
												v476 = v390
												v477 = v392
												v478 = v393
											}
											v479 = int64(3)
											v483 = int64(base.Ui64(v476)>>(uint(v479)%64)) | v477<<(uint(int64(61))%64)
											v492 = base.I64_extend_i32_u(v478)<<(uint(int64(48))%64) | int64(base.Ui64(v477)>>(uint(v479)%64))&int64(281474976710655) | v396
											v495 = base.I32_wrap_i64(v476) & int32(7)
											switch int32(0) {
											case 0:
												if v495 == int32(4) {
													v508 = v483 + v483&int64(1)
													v537 = v508
													v538 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v508) < base.Ui64(v483)))
													v543 = v537
													v544 = v538
												} else {
													v502 = v483 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v495)))
													v532 = v502
													v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v502) < base.Ui64(v483)))
													if v495 == int32(0) {
														v543 = v532
														v544 = v533
													} else {
														v537 = v532
														v538 = v533
														v543 = v537
														v544 = v538
													}
												}
											case 1:
												v518 = v483 + base.I64_extend_i32_u(base.B2i32(v396 != int64(0))&base.B2i32(v495 != int32(0)))
												v532 = v518
												v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v518) < base.Ui64(v483)))
												if v495 == int32(0) {
													v543 = v532
													v544 = v533
												} else {
													v537 = v532
													v538 = v533
													v543 = v537
													v544 = v538
												}
											case 2:
												v528 = v483 + base.I64_extend_i32_u(base.B2i32(v396 == int64(0))&base.B2i32(v495 != int32(0)))
												v532 = v528
												v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v528) < base.Ui64(v483)))
												if v495 == int32(0) {
													v543 = v532
													v544 = v533
												} else {
													v537 = v532
													v538 = v533
													v543 = v537
													v544 = v538
												}
											default:
												v532 = v483
												v533 = v492
												if v495 == int32(0) {
													v543 = v532
													v544 = v533
												} else {
													v537 = v532
													v538 = v533
													v543 = v537
													v544 = v538
												}
											}
										} else {
											v543 = int64(0)
											v544 = v396 | int64(9223090561878065152)
										}
									}
								}
							} else {
								v543 = l1
								v544 = l2
							}
						} else {
							if l3|v19 != int64(0) {
								v543 = l3
								v544 = l4
							} else {
								v543 = l3 & l1
								v544 = l4 & l2
							}
						}
					}
				} else {
					v69 = base.B2i32(l3^l1|(l4^l2^int64(-9223372036854775807-1)) == int64(0))
					if l3^l1|(l4^l2^int64(-9223372036854775807-1)) == int64(0) {
						v70 = int64(9223231299366420480)
					} else {
						v70 = l2
					}
					if l3^l1|(l4^l2^int64(-9223372036854775807-1)) == int64(0) {
						v72 = int64(0)
					} else {
						v72 = l1
					}
					v543 = v72
					v544 = v70
				}
			} else {
				v543 = l3
				v544 = l4 | int64(140737488355328)
			}
		} else {
			v543 = l1
			v544 = l2 | int64(140737488355328)
		}
	} else {
		v33 = int64(-9223090561878065152)
		v34 = v19 + v33
		if v34 == v33 {
			v39 = base.B2i32(l3 != int64(0))
		} else {
			v39 = base.B2i32(base.Ui64(v33) < base.Ui64(v34))
		}
		if v39 != 0 {
			if v19 == v23 {
				v95 = base.B2i32(base.Ui64(l1) < base.Ui64(l3))
			} else {
				v95 = base.B2i32(base.Ui64(v23) < base.Ui64(v19))
			}
			if v95 != 0 {
				v96 = l3
			} else {
				v96 = l1
			}
			if v95 != 0 {
				v97 = l4
			} else {
				v97 = l2
			}
			v99 = v97 & int64(281474976710655)
			if v95 != 0 {
				v100 = l2
			} else {
				v100 = l4
			}
			v101 = int64(48)
			v104 = int32(32767)
			v105 = base.I32_wrap_i64(int64(base.Ui64(v100)>>(uint(v101)%64))) & v104
			v110 = base.I32_wrap_i64(int64(base.Ui64(v97)>>(uint(v101)%64))) & v104
			if v110 != 0 {
				v155 = v96
				v156 = v110
				v157 = v99
			} else {
				v112 = v16 + int32(96)
				v114 = base.B2i32(v99 == int64(0))
				if v99 == int64(0) {
					v115 = v96
				} else {
					v115 = v99
				}
				v121 = base.I32_wrap_i64(base.I64_clz(v115) + base.I64_extend_i32_u(v114<<(uint(int32(6))%32)))
				v123 = v121 + int32(-15)
				if v123&int32(64) == int32(0) {
					if v123 == int32(0) {
						v144 = v96
						v145 = v99
					} else {
						v140 = base.I64_extend_i32_u(v123)
						v144 = v96 << (uint(v140) % 64)
						v145 = int64(base.Ui64(v96)>>(uint(base.I64_extend_i32_u(int32(64)-v123))%64)) | v99<<(uint(v140)%64)
					}
				} else {
					v144 = int64(0)
					v145 = v96 << (uint(base.I64_extend_i32_u(v121+int32(-79))) % 64)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v112))) = v144
				*(*int64)(unsafe.Add(mBase, uint32(v112)+8)) = v145
				v153 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(104))))
				v154 = *(*int64)(unsafe.Add(mBase, uint32(v16)+96))
				v155 = v154
				v156 = int32(16) - v121
				v157 = v153
			}
			if v95 != 0 {
				v158 = l1
			} else {
				v158 = l3
			}
			v160 = v100 & int64(281474976710655)
			if v105 != 0 {
				v205 = v160
				v206 = v158
				v208 = v105
			} else {
				v162 = v16 + int32(80)
				v164 = base.B2i32(v160 == int64(0))
				if v160 == int64(0) {
					v165 = v158
				} else {
					v165 = v160
				}
				v171 = base.I32_wrap_i64(base.I64_clz(v165) + base.I64_extend_i32_u(v164<<(uint(int32(6))%32)))
				v173 = v171 + int32(-15)
				if v173&int32(64) == int32(0) {
					if v173 == int32(0) {
						v194 = v158
						v195 = v160
					} else {
						v190 = base.I64_extend_i32_u(v173)
						v194 = v158 << (uint(v190) % 64)
						v195 = int64(base.Ui64(v158)>>(uint(base.I64_extend_i32_u(int32(64)-v173))%64)) | v160<<(uint(v190)%64)
					}
				} else {
					v194 = int64(0)
					v195 = v158 << (uint(base.I64_extend_i32_u(v171+int32(-79))) % 64)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v162))) = v194
				*(*int64)(unsafe.Add(mBase, uint32(v162)+8)) = v195
				v203 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(88))))
				v204 = *(*int64)(unsafe.Add(mBase, uint32(v16)+80))
				v205 = v203
				v206 = v204
				v208 = int32(16) - v171
			}
			v209 = int64(3)
			v211 = int64(61)
			v215 = v205<<(uint(v209)%64) | int64(base.Ui64(v206)>>(uint(v211)%64)) | int64(2251799813685248)
			v222 = v206 << (uint(v209) % 64)
			if v156 == v208 {
				v303 = v215
				v304 = v222
			} else {
				v225 = v156 - v208
				if base.Ui32(v225) <= base.Ui32(int32(127)) {
					v230 = int32(64)
					v231 = v16 + v230
					v233 = int32(128) - v225
					if v233&v230 == int32(0) {
						if v233 == int32(0) {
							v254 = v222
							v255 = v215
						} else {
							v250 = base.I64_extend_i32_u(v233)
							v254 = v222 << (uint(v250) % 64)
							v255 = int64(base.Ui64(v222)>>(uint(base.I64_extend_i32_u(int32(64)-v233))%64)) | v215<<(uint(v250)%64)
						}
					} else {
						v254 = int64(0)
						v255 = v222 << (uint(base.I64_extend_i32_u(v233+int32(-64))) % 64)
					}
					*(*int64)(unsafe.Add(mBase, uint32(v231))) = v254
					*(*int64)(unsafe.Add(mBase, uint32(v231)+8)) = v255
					v260 = v16 + int32(48)
					if v225&int32(64) == int32(0) {
						if v225 == int32(0) {
							v281 = v222
							v282 = v215
						} else {
							v277 = base.I64_extend_i32_u(v225)
							v281 = v215<<(uint(base.I64_extend_i32_u(int32(64)-v225))%64) | int64(base.Ui64(v222)>>(uint(v277)%64))
							v282 = int64(base.Ui64(v215) >> (uint(v277) % 64))
						}
					} else {
						v281 = int64(base.Ui64(v215) >> (uint(base.I64_extend_i32_u(v225+int32(-64))) % 64))
						v282 = int64(0)
					}
					*(*int64)(unsafe.Add(mBase, uint32(v260))) = v281
					*(*int64)(unsafe.Add(mBase, uint32(v260)+8)) = v282
					v286 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
					v287 = *(*int64)(unsafe.Add(mBase, uint32(v16)+64))
					v292 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(72))))
					v302 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(56))))
					v303 = v302
					v304 = v286 | base.I64_extend_i32_u(base.B2i32(v287|v292 != int64(0)))
				} else {
					v303 = int64(0)
					v304 = int64(1)
				}
			}
			v307 = v157<<(uint(v209)%64) | int64(base.Ui64(v155)>>(uint(v211)%64)) | int64(2251799813685248)
			v309 = v155 << (uint(int64(3)) % 64)
			if int64(-1) < l4^l2 {
				v370 = v304 + v309
				v373 = v303 + v307 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v370) < base.Ui64(v304)))
				if v373&int64(4503599627370496) == int64(0) {
					v390 = v370
					v392 = v373
					v393 = v156
				} else {
					v378 = int64(1)
					v390 = int64(base.Ui64(v370)>>(uint(v378)%64)) | v373<<(uint(int64(63))%64) | v304&v378
					v392 = int64(base.Ui64(v373) >> (uint(v378) % 64))
					v393 = v156 + int32(1)
				}
				v396 = v97 & int64(-9223372036854775807-1)
				if v393 < int32(32767) {
					v402 = int32(0)
					if v393 <= v402 {
						v406 = v16 + int32(16)
						v408 = v393 + int32(127)
						if v408&int32(64) == int32(0) {
							if v408 == int32(0) {
								v429 = v390
								v430 = v392
							} else {
								v425 = base.I64_extend_i32_u(v408)
								v429 = v390 << (uint(v425) % 64)
								v430 = int64(base.Ui64(v390)>>(uint(base.I64_extend_i32_u(int32(64)-v408))%64)) | v392<<(uint(v425)%64)
							}
						} else {
							v429 = int64(0)
							v430 = v390 << (uint(base.I64_extend_i32_u(v393+int32(63))) % 64)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v406))) = v429
						*(*int64)(unsafe.Add(mBase, uint32(v406)+8)) = v430
						v435 = int32(1) - v393
						if v435&int32(64) == int32(0) {
							if v435 == int32(0) {
								v456 = v390
								v457 = v392
							} else {
								v452 = base.I64_extend_i32_u(v435)
								v456 = v392<<(uint(base.I64_extend_i32_u(int32(64)-v435))%64) | int64(base.Ui64(v390)>>(uint(v452)%64))
								v457 = int64(base.Ui64(v392) >> (uint(v452) % 64))
							}
						} else {
							v456 = int64(base.Ui64(v392) >> (uint(base.I64_extend_i32_u(v435+int32(-64))) % 64))
							v457 = int64(0)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v16))) = v456
						*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v457
						v461 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						v462 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
						v467 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(24))))
						v475 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(8))))
						v476 = v461 | base.I64_extend_i32_u(base.B2i32(v462|v467 != int64(0)))
						v477 = v475
						v478 = v402
					} else {
						v476 = v390
						v477 = v392
						v478 = v393
					}
					v479 = int64(3)
					v483 = int64(base.Ui64(v476)>>(uint(v479)%64)) | v477<<(uint(int64(61))%64)
					v492 = base.I64_extend_i32_u(v478)<<(uint(int64(48))%64) | int64(base.Ui64(v477)>>(uint(v479)%64))&int64(281474976710655) | v396
					v495 = base.I32_wrap_i64(v476) & int32(7)
					switch int32(0) {
					case 0:
						if v495 == int32(4) {
							v508 = v483 + v483&int64(1)
							v537 = v508
							v538 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v508) < base.Ui64(v483)))
							v543 = v537
							v544 = v538
						} else {
							v502 = v483 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v495)))
							v532 = v502
							v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v502) < base.Ui64(v483)))
							if v495 == int32(0) {
								v543 = v532
								v544 = v533
							} else {
								v537 = v532
								v538 = v533
								v543 = v537
								v544 = v538
							}
						}
					case 1:
						v518 = v483 + base.I64_extend_i32_u(base.B2i32(v396 != int64(0))&base.B2i32(v495 != int32(0)))
						v532 = v518
						v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v518) < base.Ui64(v483)))
						if v495 == int32(0) {
							v543 = v532
							v544 = v533
						} else {
							v537 = v532
							v538 = v533
							v543 = v537
							v544 = v538
						}
					case 2:
						v528 = v483 + base.I64_extend_i32_u(base.B2i32(v396 == int64(0))&base.B2i32(v495 != int32(0)))
						v532 = v528
						v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v528) < base.Ui64(v483)))
						if v495 == int32(0) {
							v543 = v532
							v544 = v533
						} else {
							v537 = v532
							v538 = v533
							v543 = v537
							v544 = v538
						}
					default:
						v532 = v483
						v533 = v492
						if v495 == int32(0) {
							v543 = v532
							v544 = v533
						} else {
							v537 = v532
							v538 = v533
							v543 = v537
							v544 = v538
						}
					}
				} else {
					v543 = int64(0)
					v544 = v396 | int64(9223090561878065152)
				}
			} else {
				v312 = int64(0)
				if v309^v304|(v307^v303) == v312 {
					v543 = v312
					v544 = v312
				} else {
					v319 = v309 - v304
					v323 = v307 - v303 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v309) < base.Ui64(v304)))
					if base.Ui64(int64(2251799813685247)) < base.Ui64(v323) {
						v390 = v319
						v392 = v323
						v393 = v156
					} else {
						v327 = v16 + int32(32)
						v329 = base.B2i32(v323 == int64(0))
						if v323 == int64(0) {
							v330 = v319
						} else {
							v330 = v323
						}
						v336 = base.I32_wrap_i64(base.I64_clz(v330) + base.I64_extend_i32_u(v329<<(uint(int32(6))%32)))
						v338 = v336 + int32(-12)
						if v338&int32(64) == int32(0) {
							if v338 == int32(0) {
								v359 = v319
								v360 = v323
							} else {
								v355 = base.I64_extend_i32_u(v338)
								v359 = v319 << (uint(v355) % 64)
								v360 = int64(base.Ui64(v319)>>(uint(base.I64_extend_i32_u(int32(64)-v338))%64)) | v323<<(uint(v355)%64)
							}
						} else {
							v359 = int64(0)
							v360 = v319 << (uint(base.I64_extend_i32_u(v336+int32(-76))) % 64)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v327))) = v359
						*(*int64)(unsafe.Add(mBase, uint32(v327)+8)) = v360
						v367 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(40))))
						v368 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
						v390 = v368
						v392 = v367
						v393 = v156 - v338
					}
					v396 = v97 & int64(-9223372036854775807-1)
					if v393 < int32(32767) {
						v402 = int32(0)
						if v393 <= v402 {
							v406 = v16 + int32(16)
							v408 = v393 + int32(127)
							if v408&int32(64) == int32(0) {
								if v408 == int32(0) {
									v429 = v390
									v430 = v392
								} else {
									v425 = base.I64_extend_i32_u(v408)
									v429 = v390 << (uint(v425) % 64)
									v430 = int64(base.Ui64(v390)>>(uint(base.I64_extend_i32_u(int32(64)-v408))%64)) | v392<<(uint(v425)%64)
								}
							} else {
								v429 = int64(0)
								v430 = v390 << (uint(base.I64_extend_i32_u(v393+int32(63))) % 64)
							}
							*(*int64)(unsafe.Add(mBase, uint32(v406))) = v429
							*(*int64)(unsafe.Add(mBase, uint32(v406)+8)) = v430
							v435 = int32(1) - v393
							if v435&int32(64) == int32(0) {
								if v435 == int32(0) {
									v456 = v390
									v457 = v392
								} else {
									v452 = base.I64_extend_i32_u(v435)
									v456 = v392<<(uint(base.I64_extend_i32_u(int32(64)-v435))%64) | int64(base.Ui64(v390)>>(uint(v452)%64))
									v457 = int64(base.Ui64(v392) >> (uint(v452) % 64))
								}
							} else {
								v456 = int64(base.Ui64(v392) >> (uint(base.I64_extend_i32_u(v435+int32(-64))) % 64))
								v457 = int64(0)
							}
							*(*int64)(unsafe.Add(mBase, uint32(v16))) = v456
							*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v457
							v461 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							v462 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
							v467 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(24))))
							v475 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(8))))
							v476 = v461 | base.I64_extend_i32_u(base.B2i32(v462|v467 != int64(0)))
							v477 = v475
							v478 = v402
						} else {
							v476 = v390
							v477 = v392
							v478 = v393
						}
						v479 = int64(3)
						v483 = int64(base.Ui64(v476)>>(uint(v479)%64)) | v477<<(uint(int64(61))%64)
						v492 = base.I64_extend_i32_u(v478)<<(uint(int64(48))%64) | int64(base.Ui64(v477)>>(uint(v479)%64))&int64(281474976710655) | v396
						v495 = base.I32_wrap_i64(v476) & int32(7)
						switch int32(0) {
						case 0:
							if v495 == int32(4) {
								v508 = v483 + v483&int64(1)
								v537 = v508
								v538 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v508) < base.Ui64(v483)))
								v543 = v537
								v544 = v538
							} else {
								v502 = v483 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v495)))
								v532 = v502
								v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v502) < base.Ui64(v483)))
								if v495 == int32(0) {
									v543 = v532
									v544 = v533
								} else {
									v537 = v532
									v538 = v533
									v543 = v537
									v544 = v538
								}
							}
						case 1:
							v518 = v483 + base.I64_extend_i32_u(base.B2i32(v396 != int64(0))&base.B2i32(v495 != int32(0)))
							v532 = v518
							v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v518) < base.Ui64(v483)))
							if v495 == int32(0) {
								v543 = v532
								v544 = v533
							} else {
								v537 = v532
								v538 = v533
								v543 = v537
								v544 = v538
							}
						case 2:
							v528 = v483 + base.I64_extend_i32_u(base.B2i32(v396 == int64(0))&base.B2i32(v495 != int32(0)))
							v532 = v528
							v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v528) < base.Ui64(v483)))
							if v495 == int32(0) {
								v543 = v532
								v544 = v533
							} else {
								v537 = v532
								v538 = v533
								v543 = v537
								v544 = v538
							}
						default:
							v532 = v483
							v533 = v492
							if v495 == int32(0) {
								v543 = v532
								v544 = v533
							} else {
								v537 = v532
								v538 = v533
								v543 = v537
								v544 = v538
							}
						}
					} else {
						v543 = int64(0)
						v544 = v396 | int64(9223090561878065152)
					}
				}
			}
		} else {
			v41 = int64(9223090561878065152)
			if v23 == v41 {
				v45 = v21
			} else {
				v45 = base.B2i32(base.Ui64(v23) < base.Ui64(v41))
			}
			if v45 != 0 {
				v50 = int64(9223090561878065152)
				if v19 == v50 {
					v54 = base.B2i32(l3 == int64(0))
				} else {
					v54 = base.B2i32(base.Ui64(v19) < base.Ui64(v50))
				}
				if v54 != 0 {
					if l1|(v23^int64(9223090561878065152)) != int64(0) {
						if l3|(v19^int64(9223090561878065152)) == int64(0) {
							v543 = l3
							v544 = l4
						} else {
							if l1|v23 != int64(0) {
								if base.B2i32(l3|v19 == int64(0)) == int32(0) {
									if v19 == v23 {
										v95 = base.B2i32(base.Ui64(l1) < base.Ui64(l3))
									} else {
										v95 = base.B2i32(base.Ui64(v23) < base.Ui64(v19))
									}
									if v95 != 0 {
										v96 = l3
									} else {
										v96 = l1
									}
									if v95 != 0 {
										v97 = l4
									} else {
										v97 = l2
									}
									v99 = v97 & int64(281474976710655)
									if v95 != 0 {
										v100 = l2
									} else {
										v100 = l4
									}
									v101 = int64(48)
									v104 = int32(32767)
									v105 = base.I32_wrap_i64(int64(base.Ui64(v100)>>(uint(v101)%64))) & v104
									v110 = base.I32_wrap_i64(int64(base.Ui64(v97)>>(uint(v101)%64))) & v104
									if v110 != 0 {
										v155 = v96
										v156 = v110
										v157 = v99
									} else {
										v112 = v16 + int32(96)
										v114 = base.B2i32(v99 == int64(0))
										if v99 == int64(0) {
											v115 = v96
										} else {
											v115 = v99
										}
										v121 = base.I32_wrap_i64(base.I64_clz(v115) + base.I64_extend_i32_u(v114<<(uint(int32(6))%32)))
										v123 = v121 + int32(-15)
										if v123&int32(64) == int32(0) {
											if v123 == int32(0) {
												v144 = v96
												v145 = v99
											} else {
												v140 = base.I64_extend_i32_u(v123)
												v144 = v96 << (uint(v140) % 64)
												v145 = int64(base.Ui64(v96)>>(uint(base.I64_extend_i32_u(int32(64)-v123))%64)) | v99<<(uint(v140)%64)
											}
										} else {
											v144 = int64(0)
											v145 = v96 << (uint(base.I64_extend_i32_u(v121+int32(-79))) % 64)
										}
										*(*int64)(unsafe.Add(mBase, uint32(v112))) = v144
										*(*int64)(unsafe.Add(mBase, uint32(v112)+8)) = v145
										v153 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(104))))
										v154 = *(*int64)(unsafe.Add(mBase, uint32(v16)+96))
										v155 = v154
										v156 = int32(16) - v121
										v157 = v153
									}
									if v95 != 0 {
										v158 = l1
									} else {
										v158 = l3
									}
									v160 = v100 & int64(281474976710655)
									if v105 != 0 {
										v205 = v160
										v206 = v158
										v208 = v105
									} else {
										v162 = v16 + int32(80)
										v164 = base.B2i32(v160 == int64(0))
										if v160 == int64(0) {
											v165 = v158
										} else {
											v165 = v160
										}
										v171 = base.I32_wrap_i64(base.I64_clz(v165) + base.I64_extend_i32_u(v164<<(uint(int32(6))%32)))
										v173 = v171 + int32(-15)
										if v173&int32(64) == int32(0) {
											if v173 == int32(0) {
												v194 = v158
												v195 = v160
											} else {
												v190 = base.I64_extend_i32_u(v173)
												v194 = v158 << (uint(v190) % 64)
												v195 = int64(base.Ui64(v158)>>(uint(base.I64_extend_i32_u(int32(64)-v173))%64)) | v160<<(uint(v190)%64)
											}
										} else {
											v194 = int64(0)
											v195 = v158 << (uint(base.I64_extend_i32_u(v171+int32(-79))) % 64)
										}
										*(*int64)(unsafe.Add(mBase, uint32(v162))) = v194
										*(*int64)(unsafe.Add(mBase, uint32(v162)+8)) = v195
										v203 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(88))))
										v204 = *(*int64)(unsafe.Add(mBase, uint32(v16)+80))
										v205 = v203
										v206 = v204
										v208 = int32(16) - v171
									}
									v209 = int64(3)
									v211 = int64(61)
									v215 = v205<<(uint(v209)%64) | int64(base.Ui64(v206)>>(uint(v211)%64)) | int64(2251799813685248)
									v222 = v206 << (uint(v209) % 64)
									if v156 == v208 {
										v303 = v215
										v304 = v222
									} else {
										v225 = v156 - v208
										if base.Ui32(v225) <= base.Ui32(int32(127)) {
											v230 = int32(64)
											v231 = v16 + v230
											v233 = int32(128) - v225
											if v233&v230 == int32(0) {
												if v233 == int32(0) {
													v254 = v222
													v255 = v215
												} else {
													v250 = base.I64_extend_i32_u(v233)
													v254 = v222 << (uint(v250) % 64)
													v255 = int64(base.Ui64(v222)>>(uint(base.I64_extend_i32_u(int32(64)-v233))%64)) | v215<<(uint(v250)%64)
												}
											} else {
												v254 = int64(0)
												v255 = v222 << (uint(base.I64_extend_i32_u(v233+int32(-64))) % 64)
											}
											*(*int64)(unsafe.Add(mBase, uint32(v231))) = v254
											*(*int64)(unsafe.Add(mBase, uint32(v231)+8)) = v255
											v260 = v16 + int32(48)
											if v225&int32(64) == int32(0) {
												if v225 == int32(0) {
													v281 = v222
													v282 = v215
												} else {
													v277 = base.I64_extend_i32_u(v225)
													v281 = v215<<(uint(base.I64_extend_i32_u(int32(64)-v225))%64) | int64(base.Ui64(v222)>>(uint(v277)%64))
													v282 = int64(base.Ui64(v215) >> (uint(v277) % 64))
												}
											} else {
												v281 = int64(base.Ui64(v215) >> (uint(base.I64_extend_i32_u(v225+int32(-64))) % 64))
												v282 = int64(0)
											}
											*(*int64)(unsafe.Add(mBase, uint32(v260))) = v281
											*(*int64)(unsafe.Add(mBase, uint32(v260)+8)) = v282
											v286 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
											v287 = *(*int64)(unsafe.Add(mBase, uint32(v16)+64))
											v292 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(72))))
											v302 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(56))))
											v303 = v302
											v304 = v286 | base.I64_extend_i32_u(base.B2i32(v287|v292 != int64(0)))
										} else {
											v303 = int64(0)
											v304 = int64(1)
										}
									}
									v307 = v157<<(uint(v209)%64) | int64(base.Ui64(v155)>>(uint(v211)%64)) | int64(2251799813685248)
									v309 = v155 << (uint(int64(3)) % 64)
									if int64(-1) < l4^l2 {
										v370 = v304 + v309
										v373 = v303 + v307 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v370) < base.Ui64(v304)))
										if v373&int64(4503599627370496) == int64(0) {
											v390 = v370
											v392 = v373
											v393 = v156
										} else {
											v378 = int64(1)
											v390 = int64(base.Ui64(v370)>>(uint(v378)%64)) | v373<<(uint(int64(63))%64) | v304&v378
											v392 = int64(base.Ui64(v373) >> (uint(v378) % 64))
											v393 = v156 + int32(1)
										}
										v396 = v97 & int64(-9223372036854775807-1)
										if v393 < int32(32767) {
											v402 = int32(0)
											if v393 <= v402 {
												v406 = v16 + int32(16)
												v408 = v393 + int32(127)
												if v408&int32(64) == int32(0) {
													if v408 == int32(0) {
														v429 = v390
														v430 = v392
													} else {
														v425 = base.I64_extend_i32_u(v408)
														v429 = v390 << (uint(v425) % 64)
														v430 = int64(base.Ui64(v390)>>(uint(base.I64_extend_i32_u(int32(64)-v408))%64)) | v392<<(uint(v425)%64)
													}
												} else {
													v429 = int64(0)
													v430 = v390 << (uint(base.I64_extend_i32_u(v393+int32(63))) % 64)
												}
												*(*int64)(unsafe.Add(mBase, uint32(v406))) = v429
												*(*int64)(unsafe.Add(mBase, uint32(v406)+8)) = v430
												v435 = int32(1) - v393
												if v435&int32(64) == int32(0) {
													if v435 == int32(0) {
														v456 = v390
														v457 = v392
													} else {
														v452 = base.I64_extend_i32_u(v435)
														v456 = v392<<(uint(base.I64_extend_i32_u(int32(64)-v435))%64) | int64(base.Ui64(v390)>>(uint(v452)%64))
														v457 = int64(base.Ui64(v392) >> (uint(v452) % 64))
													}
												} else {
													v456 = int64(base.Ui64(v392) >> (uint(base.I64_extend_i32_u(v435+int32(-64))) % 64))
													v457 = int64(0)
												}
												*(*int64)(unsafe.Add(mBase, uint32(v16))) = v456
												*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v457
												v461 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
												v462 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
												v467 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(24))))
												v475 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(8))))
												v476 = v461 | base.I64_extend_i32_u(base.B2i32(v462|v467 != int64(0)))
												v477 = v475
												v478 = v402
											} else {
												v476 = v390
												v477 = v392
												v478 = v393
											}
											v479 = int64(3)
											v483 = int64(base.Ui64(v476)>>(uint(v479)%64)) | v477<<(uint(int64(61))%64)
											v492 = base.I64_extend_i32_u(v478)<<(uint(int64(48))%64) | int64(base.Ui64(v477)>>(uint(v479)%64))&int64(281474976710655) | v396
											v495 = base.I32_wrap_i64(v476) & int32(7)
											switch int32(0) {
											case 0:
												if v495 == int32(4) {
													v508 = v483 + v483&int64(1)
													v537 = v508
													v538 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v508) < base.Ui64(v483)))
													v543 = v537
													v544 = v538
												} else {
													v502 = v483 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v495)))
													v532 = v502
													v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v502) < base.Ui64(v483)))
													if v495 == int32(0) {
														v543 = v532
														v544 = v533
													} else {
														v537 = v532
														v538 = v533
														v543 = v537
														v544 = v538
													}
												}
											case 1:
												v518 = v483 + base.I64_extend_i32_u(base.B2i32(v396 != int64(0))&base.B2i32(v495 != int32(0)))
												v532 = v518
												v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v518) < base.Ui64(v483)))
												if v495 == int32(0) {
													v543 = v532
													v544 = v533
												} else {
													v537 = v532
													v538 = v533
													v543 = v537
													v544 = v538
												}
											case 2:
												v528 = v483 + base.I64_extend_i32_u(base.B2i32(v396 == int64(0))&base.B2i32(v495 != int32(0)))
												v532 = v528
												v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v528) < base.Ui64(v483)))
												if v495 == int32(0) {
													v543 = v532
													v544 = v533
												} else {
													v537 = v532
													v538 = v533
													v543 = v537
													v544 = v538
												}
											default:
												v532 = v483
												v533 = v492
												if v495 == int32(0) {
													v543 = v532
													v544 = v533
												} else {
													v537 = v532
													v538 = v533
													v543 = v537
													v544 = v538
												}
											}
										} else {
											v543 = int64(0)
											v544 = v396 | int64(9223090561878065152)
										}
									} else {
										v312 = int64(0)
										if v309^v304|(v307^v303) == v312 {
											v543 = v312
											v544 = v312
										} else {
											v319 = v309 - v304
											v323 = v307 - v303 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v309) < base.Ui64(v304)))
											if base.Ui64(int64(2251799813685247)) < base.Ui64(v323) {
												v390 = v319
												v392 = v323
												v393 = v156
											} else {
												v327 = v16 + int32(32)
												v329 = base.B2i32(v323 == int64(0))
												if v323 == int64(0) {
													v330 = v319
												} else {
													v330 = v323
												}
												v336 = base.I32_wrap_i64(base.I64_clz(v330) + base.I64_extend_i32_u(v329<<(uint(int32(6))%32)))
												v338 = v336 + int32(-12)
												if v338&int32(64) == int32(0) {
													if v338 == int32(0) {
														v359 = v319
														v360 = v323
													} else {
														v355 = base.I64_extend_i32_u(v338)
														v359 = v319 << (uint(v355) % 64)
														v360 = int64(base.Ui64(v319)>>(uint(base.I64_extend_i32_u(int32(64)-v338))%64)) | v323<<(uint(v355)%64)
													}
												} else {
													v359 = int64(0)
													v360 = v319 << (uint(base.I64_extend_i32_u(v336+int32(-76))) % 64)
												}
												*(*int64)(unsafe.Add(mBase, uint32(v327))) = v359
												*(*int64)(unsafe.Add(mBase, uint32(v327)+8)) = v360
												v367 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(40))))
												v368 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
												v390 = v368
												v392 = v367
												v393 = v156 - v338
											}
											v396 = v97 & int64(-9223372036854775807-1)
											if v393 < int32(32767) {
												v402 = int32(0)
												if v393 <= v402 {
													v406 = v16 + int32(16)
													v408 = v393 + int32(127)
													if v408&int32(64) == int32(0) {
														if v408 == int32(0) {
															v429 = v390
															v430 = v392
														} else {
															v425 = base.I64_extend_i32_u(v408)
															v429 = v390 << (uint(v425) % 64)
															v430 = int64(base.Ui64(v390)>>(uint(base.I64_extend_i32_u(int32(64)-v408))%64)) | v392<<(uint(v425)%64)
														}
													} else {
														v429 = int64(0)
														v430 = v390 << (uint(base.I64_extend_i32_u(v393+int32(63))) % 64)
													}
													*(*int64)(unsafe.Add(mBase, uint32(v406))) = v429
													*(*int64)(unsafe.Add(mBase, uint32(v406)+8)) = v430
													v435 = int32(1) - v393
													if v435&int32(64) == int32(0) {
														if v435 == int32(0) {
															v456 = v390
															v457 = v392
														} else {
															v452 = base.I64_extend_i32_u(v435)
															v456 = v392<<(uint(base.I64_extend_i32_u(int32(64)-v435))%64) | int64(base.Ui64(v390)>>(uint(v452)%64))
															v457 = int64(base.Ui64(v392) >> (uint(v452) % 64))
														}
													} else {
														v456 = int64(base.Ui64(v392) >> (uint(base.I64_extend_i32_u(v435+int32(-64))) % 64))
														v457 = int64(0)
													}
													*(*int64)(unsafe.Add(mBase, uint32(v16))) = v456
													*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v457
													v461 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
													v462 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
													v467 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(24))))
													v475 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(8))))
													v476 = v461 | base.I64_extend_i32_u(base.B2i32(v462|v467 != int64(0)))
													v477 = v475
													v478 = v402
												} else {
													v476 = v390
													v477 = v392
													v478 = v393
												}
												v479 = int64(3)
												v483 = int64(base.Ui64(v476)>>(uint(v479)%64)) | v477<<(uint(int64(61))%64)
												v492 = base.I64_extend_i32_u(v478)<<(uint(int64(48))%64) | int64(base.Ui64(v477)>>(uint(v479)%64))&int64(281474976710655) | v396
												v495 = base.I32_wrap_i64(v476) & int32(7)
												switch int32(0) {
												case 0:
													if v495 == int32(4) {
														v508 = v483 + v483&int64(1)
														v537 = v508
														v538 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v508) < base.Ui64(v483)))
														v543 = v537
														v544 = v538
													} else {
														v502 = v483 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v495)))
														v532 = v502
														v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v502) < base.Ui64(v483)))
														if v495 == int32(0) {
															v543 = v532
															v544 = v533
														} else {
															v537 = v532
															v538 = v533
															v543 = v537
															v544 = v538
														}
													}
												case 1:
													v518 = v483 + base.I64_extend_i32_u(base.B2i32(v396 != int64(0))&base.B2i32(v495 != int32(0)))
													v532 = v518
													v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v518) < base.Ui64(v483)))
													if v495 == int32(0) {
														v543 = v532
														v544 = v533
													} else {
														v537 = v532
														v538 = v533
														v543 = v537
														v544 = v538
													}
												case 2:
													v528 = v483 + base.I64_extend_i32_u(base.B2i32(v396 == int64(0))&base.B2i32(v495 != int32(0)))
													v532 = v528
													v533 = v492 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v528) < base.Ui64(v483)))
													if v495 == int32(0) {
														v543 = v532
														v544 = v533
													} else {
														v537 = v532
														v538 = v533
														v543 = v537
														v544 = v538
													}
												default:
													v532 = v483
													v533 = v492
													if v495 == int32(0) {
														v543 = v532
														v544 = v533
													} else {
														v537 = v532
														v538 = v533
														v543 = v537
														v544 = v538
													}
												}
											} else {
												v543 = int64(0)
												v544 = v396 | int64(9223090561878065152)
											}
										}
									}
								} else {
									v543 = l1
									v544 = l2
								}
							} else {
								if l3|v19 != int64(0) {
									v543 = l3
									v544 = l4
								} else {
									v543 = l3 & l1
									v544 = l4 & l2
								}
							}
						}
					} else {
						v69 = base.B2i32(l3^l1|(l4^l2^int64(-9223372036854775807-1)) == int64(0))
						if l3^l1|(l4^l2^int64(-9223372036854775807-1)) == int64(0) {
							v70 = int64(9223231299366420480)
						} else {
							v70 = l2
						}
						if l3^l1|(l4^l2^int64(-9223372036854775807-1)) == int64(0) {
							v72 = int64(0)
						} else {
							v72 = l1
						}
						v543 = v72
						v544 = v70
					}
				} else {
					v543 = l3
					v544 = l4 | int64(140737488355328)
				}
			} else {
				v543 = l1
				v544 = l2 | int64(140737488355328)
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v543
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v544
	m.G0 = v16 + int32(112)
	return
}
func F__addBulkStrRefToBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v7&int32(16777216) != 0 {
		v15 = v7
		v17 = int32(0)
		v19 = *(*int32)(unsafe.Add(mBase, _consts[484]))
		if v19 != 0 {
			v148 = v17
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
			if v21 != 0 {
				v148 = v17
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
				v24 = v22 - v23
				if v15&int32(16777216) != 0 {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
					v31 = l0 + int32(180)
					v33 = l0 + int32(184)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
					v37 = *(*int64)(unsafe.Add(mBase, _consts[480]))
					v39 = base.B2i32(v37 != int64(-1))
					v40 = int32(0)
					if base.Ui32(v24) < base.Ui32(l2) {
						v119 = v40
						v130 = v119
					} else {
						v52 = F_clusterSlotStatsEnabled(m, v35)
						mBase = m.M
						if v52 != 0 {
							v53 = v35
						} else {
							v53 = int32(-1)
						}
						if base.Ui32(v24) < base.Ui32(l2) {
							v55 = v24
						} else {
							v55 = l2
						}
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						if v56 == int32(0) {
							if base.Ui32(v24) < base.Ui32(l2+int32(12)) {
								v119 = v40
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
								v80 = v29 + v79
								*(*int32)(unsafe.Add(mBase, uint32(v33))) = v80
								v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
								v87 = v82&int32(254) | int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v87)
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
								*(*uint16)(unsafe.Add(mBase, uint32(v89)+8)) = uint16(v53)
								v92 = v24 + int32(-12)
								if base.Ui32(v92) < base.Ui32(l2) {
									v94 = v92
								} else {
									v94 = v55
								}
								*(*int32)(unsafe.Add(mBase, uint32(v89))) = v94
								v96 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v96
								v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)))
								v103 = v98&int32(253) | v39<<(uint(int32(1))%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)) = uint8(v103)
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
								*(*uint8)(unsafe.Add(mBase, uint32(v105)+11)) = uint8(v96)
								v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)))
								v110 = v108 & int32(3)
								*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)) = uint8(v110)
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
								*(*int32)(unsafe.Add(mBase, uint32(v31))) = v112 + int32(12)
								v119 = v94
							}
							v130 = v119
						} else {
							v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+10)))
							if v59&int32(1) != int32(1) {
								if base.Ui32(v24) < base.Ui32(l2+int32(12)) {
									v119 = v40
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									v80 = v29 + v79
									*(*int32)(unsafe.Add(mBase, uint32(v33))) = v80
									v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
									v87 = v82&int32(254) | int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v87)
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
									*(*uint16)(unsafe.Add(mBase, uint32(v89)+8)) = uint16(v53)
									v92 = v24 + int32(-12)
									if base.Ui32(v92) < base.Ui32(l2) {
										v94 = v92
									} else {
										v94 = v55
									}
									*(*int32)(unsafe.Add(mBase, uint32(v89))) = v94
									v96 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v96
									v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)))
									v103 = v98&int32(253) | v39<<(uint(int32(1))%32)
									*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)) = uint8(v103)
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
									*(*uint8)(unsafe.Add(mBase, uint32(v105)+11)) = uint8(v96)
									v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)))
									v110 = v108 & int32(3)
									*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)) = uint8(v110)
									v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									*(*int32)(unsafe.Add(mBase, uint32(v31))) = v112 + int32(12)
									v119 = v94
								}
								v130 = v119
							} else {
								v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+8)))
								if v53 != v63 {
									if base.Ui32(v24) < base.Ui32(l2+int32(12)) {
										v119 = v40
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
										v80 = v29 + v79
										*(*int32)(unsafe.Add(mBase, uint32(v33))) = v80
										v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
										v87 = v82&int32(254) | int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v87)
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										*(*uint16)(unsafe.Add(mBase, uint32(v89)+8)) = uint16(v53)
										v92 = v24 + int32(-12)
										if base.Ui32(v92) < base.Ui32(l2) {
											v94 = v92
										} else {
											v94 = v55
										}
										*(*int32)(unsafe.Add(mBase, uint32(v89))) = v94
										v96 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v96
										v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)))
										v103 = v98&int32(253) | v39<<(uint(int32(1))%32)
										*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)) = uint8(v103)
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										*(*uint8)(unsafe.Add(mBase, uint32(v105)+11)) = uint8(v96)
										v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)))
										v110 = v108 & int32(3)
										*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)) = uint8(v110)
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
										*(*int32)(unsafe.Add(mBase, uint32(v31))) = v112 + int32(12)
										v119 = v94
									}
									v130 = v119
								} else {
									v65 = int32(1)
									if v39 != int32(base.Ui32(v59)>>(uint(v65)%32))&v65 {
										if base.Ui32(v24) < base.Ui32(l2+int32(12)) {
											v119 = v40
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
											v80 = v29 + v79
											*(*int32)(unsafe.Add(mBase, uint32(v33))) = v80
											v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
											v87 = v82&int32(254) | int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v87)
											v89 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
											*(*uint16)(unsafe.Add(mBase, uint32(v89)+8)) = uint16(v53)
											v92 = v24 + int32(-12)
											if base.Ui32(v92) < base.Ui32(l2) {
												v94 = v92
											} else {
												v94 = v55
											}
											*(*int32)(unsafe.Add(mBase, uint32(v89))) = v94
											v96 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v96
											v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)))
											v103 = v98&int32(253) | v39<<(uint(int32(1))%32)
											*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)) = uint8(v103)
											v105 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
											*(*uint8)(unsafe.Add(mBase, uint32(v105)+11)) = uint8(v96)
											v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)))
											v110 = v108 & int32(3)
											*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)) = uint8(v110)
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
											*(*int32)(unsafe.Add(mBase, uint32(v31))) = v112 + int32(12)
											v119 = v94
										}
										v130 = v119
									} else {
										v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+11)))
										if v70 != 0 {
											if base.Ui32(v24) < base.Ui32(l2+int32(12)) {
												v119 = v40
											} else {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
												v80 = v29 + v79
												*(*int32)(unsafe.Add(mBase, uint32(v33))) = v80
												v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
												v87 = v82&int32(254) | int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v87)
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
												*(*uint16)(unsafe.Add(mBase, uint32(v89)+8)) = uint16(v53)
												v92 = v24 + int32(-12)
												if base.Ui32(v92) < base.Ui32(l2) {
													v94 = v92
												} else {
													v94 = v55
												}
												*(*int32)(unsafe.Add(mBase, uint32(v89))) = v94
												v96 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v96
												v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)))
												v103 = v98&int32(253) | v39<<(uint(int32(1))%32)
												*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)) = uint8(v103)
												v105 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
												*(*uint8)(unsafe.Add(mBase, uint32(v105)+11)) = uint8(v96)
												v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)))
												v110 = v108 & int32(3)
												*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)) = uint8(v110)
												v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
												*(*int32)(unsafe.Add(mBase, uint32(v31))) = v112 + int32(12)
												v119 = v94
											}
											v130 = v119
										} else {
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
											*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72 + v55
											v130 = v55
										}
									}
								}
							}
						}
					}
					v131 = v130
				} else {
					if base.Ui32(v24) < base.Ui32(l2) {
						v28 = v24
					} else {
						v28 = l2
					}
					v131 = v28
				}
				if v131 == int32(0) {
					v148 = v17
				} else {
					v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
					v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
					if v131 == int32(0) {
					} else {
						v139 = F__emscripten_memcpy_bulkmem(m, v134+v135, l1, v131)
						mBase = m.M
					}
					v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
					v142 = v141 + v131
					*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v142
					v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
					if base.Ui32(v142) <= base.Ui32(v144) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v142
					}
					v148 = v131
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
		if v11 != 0 {
			v148 = int32(0)
		} else {
			v13 = v7 | int32(16777216)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v13
			v15 = v13
			v17 = int32(0)
			v19 = *(*int32)(unsafe.Add(mBase, _consts[484]))
			if v19 != 0 {
				v148 = v17
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
				if v21 != 0 {
					v148 = v17
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
					v24 = v22 - v23
					if v15&int32(16777216) != 0 {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
						v31 = l0 + int32(180)
						v33 = l0 + int32(184)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
						v37 = *(*int64)(unsafe.Add(mBase, _consts[480]))
						v39 = base.B2i32(v37 != int64(-1))
						v40 = int32(0)
						if base.Ui32(v24) < base.Ui32(l2) {
							v119 = v40
							v130 = v119
						} else {
							v52 = F_clusterSlotStatsEnabled(m, v35)
							mBase = m.M
							if v52 != 0 {
								v53 = v35
							} else {
								v53 = int32(-1)
							}
							if base.Ui32(v24) < base.Ui32(l2) {
								v55 = v24
							} else {
								v55 = l2
							}
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
							if v56 == int32(0) {
								if base.Ui32(v24) < base.Ui32(l2+int32(12)) {
									v119 = v40
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									v80 = v29 + v79
									*(*int32)(unsafe.Add(mBase, uint32(v33))) = v80
									v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
									v87 = v82&int32(254) | int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v87)
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
									*(*uint16)(unsafe.Add(mBase, uint32(v89)+8)) = uint16(v53)
									v92 = v24 + int32(-12)
									if base.Ui32(v92) < base.Ui32(l2) {
										v94 = v92
									} else {
										v94 = v55
									}
									*(*int32)(unsafe.Add(mBase, uint32(v89))) = v94
									v96 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v96
									v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)))
									v103 = v98&int32(253) | v39<<(uint(int32(1))%32)
									*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)) = uint8(v103)
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
									*(*uint8)(unsafe.Add(mBase, uint32(v105)+11)) = uint8(v96)
									v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)))
									v110 = v108 & int32(3)
									*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)) = uint8(v110)
									v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									*(*int32)(unsafe.Add(mBase, uint32(v31))) = v112 + int32(12)
									v119 = v94
								}
								v130 = v119
							} else {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+10)))
								if v59&int32(1) != int32(1) {
									if base.Ui32(v24) < base.Ui32(l2+int32(12)) {
										v119 = v40
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
										v80 = v29 + v79
										*(*int32)(unsafe.Add(mBase, uint32(v33))) = v80
										v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
										v87 = v82&int32(254) | int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v87)
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										*(*uint16)(unsafe.Add(mBase, uint32(v89)+8)) = uint16(v53)
										v92 = v24 + int32(-12)
										if base.Ui32(v92) < base.Ui32(l2) {
											v94 = v92
										} else {
											v94 = v55
										}
										*(*int32)(unsafe.Add(mBase, uint32(v89))) = v94
										v96 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v96
										v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)))
										v103 = v98&int32(253) | v39<<(uint(int32(1))%32)
										*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)) = uint8(v103)
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										*(*uint8)(unsafe.Add(mBase, uint32(v105)+11)) = uint8(v96)
										v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)))
										v110 = v108 & int32(3)
										*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)) = uint8(v110)
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
										*(*int32)(unsafe.Add(mBase, uint32(v31))) = v112 + int32(12)
										v119 = v94
									}
									v130 = v119
								} else {
									v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+8)))
									if v53 != v63 {
										if base.Ui32(v24) < base.Ui32(l2+int32(12)) {
											v119 = v40
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
											v80 = v29 + v79
											*(*int32)(unsafe.Add(mBase, uint32(v33))) = v80
											v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
											v87 = v82&int32(254) | int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v87)
											v89 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
											*(*uint16)(unsafe.Add(mBase, uint32(v89)+8)) = uint16(v53)
											v92 = v24 + int32(-12)
											if base.Ui32(v92) < base.Ui32(l2) {
												v94 = v92
											} else {
												v94 = v55
											}
											*(*int32)(unsafe.Add(mBase, uint32(v89))) = v94
											v96 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v96
											v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)))
											v103 = v98&int32(253) | v39<<(uint(int32(1))%32)
											*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)) = uint8(v103)
											v105 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
											*(*uint8)(unsafe.Add(mBase, uint32(v105)+11)) = uint8(v96)
											v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)))
											v110 = v108 & int32(3)
											*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)) = uint8(v110)
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
											*(*int32)(unsafe.Add(mBase, uint32(v31))) = v112 + int32(12)
											v119 = v94
										}
										v130 = v119
									} else {
										v65 = int32(1)
										if v39 != int32(base.Ui32(v59)>>(uint(v65)%32))&v65 {
											if base.Ui32(v24) < base.Ui32(l2+int32(12)) {
												v119 = v40
											} else {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
												v80 = v29 + v79
												*(*int32)(unsafe.Add(mBase, uint32(v33))) = v80
												v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
												v87 = v82&int32(254) | int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v87)
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
												*(*uint16)(unsafe.Add(mBase, uint32(v89)+8)) = uint16(v53)
												v92 = v24 + int32(-12)
												if base.Ui32(v92) < base.Ui32(l2) {
													v94 = v92
												} else {
													v94 = v55
												}
												*(*int32)(unsafe.Add(mBase, uint32(v89))) = v94
												v96 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v96
												v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)))
												v103 = v98&int32(253) | v39<<(uint(int32(1))%32)
												*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)) = uint8(v103)
												v105 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
												*(*uint8)(unsafe.Add(mBase, uint32(v105)+11)) = uint8(v96)
												v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)))
												v110 = v108 & int32(3)
												*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)) = uint8(v110)
												v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
												*(*int32)(unsafe.Add(mBase, uint32(v31))) = v112 + int32(12)
												v119 = v94
											}
											v130 = v119
										} else {
											v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+11)))
											if v70 != 0 {
												if base.Ui32(v24) < base.Ui32(l2+int32(12)) {
													v119 = v40
												} else {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
													v80 = v29 + v79
													*(*int32)(unsafe.Add(mBase, uint32(v33))) = v80
													v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
													v87 = v82&int32(254) | int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v87)
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
													*(*uint16)(unsafe.Add(mBase, uint32(v89)+8)) = uint16(v53)
													v92 = v24 + int32(-12)
													if base.Ui32(v92) < base.Ui32(l2) {
														v94 = v92
													} else {
														v94 = v55
													}
													*(*int32)(unsafe.Add(mBase, uint32(v89))) = v94
													v96 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v96
													v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)))
													v103 = v98&int32(253) | v39<<(uint(int32(1))%32)
													*(*uint8)(unsafe.Add(mBase, uint32(v89)+10)) = uint8(v103)
													v105 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
													*(*uint8)(unsafe.Add(mBase, uint32(v105)+11)) = uint8(v96)
													v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)))
													v110 = v108 & int32(3)
													*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)) = uint8(v110)
													v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
													*(*int32)(unsafe.Add(mBase, uint32(v31))) = v112 + int32(12)
													v119 = v94
												}
												v130 = v119
											} else {
												v71 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
												*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72 + v55
												v130 = v55
											}
										}
									}
								}
							}
						}
						v131 = v130
					} else {
						if base.Ui32(v24) < base.Ui32(l2) {
							v28 = v24
						} else {
							v28 = l2
						}
						v131 = v28
					}
					if v131 == int32(0) {
						v148 = v17
					} else {
						v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
						v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
						if v131 == int32(0) {
						} else {
							v139 = F__emscripten_memcpy_bulkmem(m, v134+v135, l1, v131)
							mBase = m.M
						}
						v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
						v142 = v141 + v131
						*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v142
						v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
						if base.Ui32(v142) <= base.Ui32(v144) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v142
						}
						v148 = v131
					}
				}
			}
		}
	}
	return v148
}
func F_a_cas_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 != 0 {
		v5 = v3
	} else {
		v5 = int32(1073741823)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5
	return v3
}
func F_a_crash(m *base.Module) {
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_a_swap_3(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1042]))
	*(*int32)(unsafe.Add(mBase, _consts[1042])) = v1
	return v3
}
func F_accept(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v4 = int32(0)
	v7 = m.Env.X__syscall_accept4(m, l0, l1, l2, v4, v4, v4)
	mBase = m.M
	if base.Ui32(v7) < base.Ui32(int32(-4095)) {
		v15 = v7
	} else {
		v10 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0) - v7
		v15 = int32(-1)
	}
	return v15
}
func F_action_terminate(m *base.Module, l0 int32) {
	m.Env.X_emscripten_runtime_keepalive_clear(m)
	F__Exit(m, l0+int32(128))
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_activeDefragAlloc(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F_activeExpireCycle(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v86 int32
	_ = v86
	v2 = int64(0)
	v8 = F_isPausedActionsWithUpdate(m, int32(4))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		if v8 != 0 {
			v77 = v2
			return v77
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[325]))
			if l0 != int32(1) {
				v37 = *(*int32)(unsafe.Add(mBase, _consts[251]))
				v38 = base.I32_div_s(v13*int32(2000000)+int32(23000000), v37)
				v40 = base.I32_div_s(v38, int32(100))
				v43 = base.I64_extend_i32_s(v40)
				v45 = *(*int32)(unsafe.Add(mBase, _consts[326]))
				if v45 != 0 {
					F__serverAssert(m, int32(_a612), int32(_a613), int32(493))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int64(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[327])))
					v50 = F_activeExpireCycleJob(m, v47, l0, v43)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						v53 = F_activeExpireCycleJob(m, v47^int32(1), l0, v43-v50)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							v55 = int32(_a20)
							v57 = *(*int64)(unsafe.Add(mBase, _consts[328]))
							v58 = v53 + v50
							*(*int64)(unsafe.Add(mBase, _consts[328])) = v57 + v58
							v62 = *(*int64)(unsafe.Add(mBase, _consts[45]))
							if v62 == int64(0) {
								v71 = int32(0)
								v73 = int32(*(*uint8)(unsafe.Add(mBase, _consts[327])))
								v75 = v73 ^ int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[327])) = uint8(v75)
								v77 = v58
								return v77
							} else {
								if v58 < v62*int64(1000) {
									v71 = int32(0)
									v73 = int32(*(*uint8)(unsafe.Add(mBase, _consts[327])))
									v75 = v73 ^ int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[327])) = uint8(v75)
									v77 = v58
									return v77
								} else {
									F_latencyAddSample(m, int32(_a614), v58)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int64(0)
									} else {
										v71 = int32(0)
										v73 = int32(*(*uint8)(unsafe.Add(mBase, _consts[327])))
										v75 = v73 ^ int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[327])) = uint8(v75)
										v77 = v58
										return v77
									}
								}
							}
						}
					}
				}
			} else {
				v16 = int32(0)
				v17 = *(*int32)(unsafe.Add(mBase, _consts[34]))
				v18 = m.T0[v17].(func(*base.Module) int64)(m)
				mBase = m.M
				v23 = base.I64_extend_i32_s(v13*int32(250) + int32(750))
				v27 = *(*int64)(unsafe.Add(mBase, _consts[329]))
				if base.Ui64(v18) < base.Ui64(v23<<(uint(int64(1))%64)+v27) {
					v77 = v2
					return v77
				} else {
					*(*int64)(unsafe.Add(mBase, _consts[329])) = v18
					v43 = v23
					v45 = *(*int32)(unsafe.Add(mBase, _consts[326]))
					if v45 != 0 {
						F__serverAssert(m, int32(_a612), int32(_a613), int32(493))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int64(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[327])))
						v50 = F_activeExpireCycleJob(m, v47, l0, v43)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							v53 = F_activeExpireCycleJob(m, v47^int32(1), l0, v43-v50)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								v55 = int32(_a20)
								v57 = *(*int64)(unsafe.Add(mBase, _consts[328]))
								v58 = v53 + v50
								*(*int64)(unsafe.Add(mBase, _consts[328])) = v57 + v58
								v62 = *(*int64)(unsafe.Add(mBase, _consts[45]))
								if v62 == int64(0) {
									v71 = int32(0)
									v73 = int32(*(*uint8)(unsafe.Add(mBase, _consts[327])))
									v75 = v73 ^ int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[327])) = uint8(v75)
									v77 = v58
									return v77
								} else {
									if v58 < v62*int64(1000) {
										v71 = int32(0)
										v73 = int32(*(*uint8)(unsafe.Add(mBase, _consts[327])))
										v75 = v73 ^ int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[327])) = uint8(v75)
										v77 = v58
										return v77
									} else {
										F_latencyAddSample(m, int32(_a614), v58)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int64(0)
										} else {
											v71 = int32(0)
											v73 = int32(*(*uint8)(unsafe.Add(mBase, _consts[327])))
											v75 = v73 ^ int32(1)
											*(*uint8)(unsafe.Add(mBase, _consts[327])) = uint8(v75)
											v77 = v58
											return v77
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
func F_addExtendedReplyHelp(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v15 = F_objectGetVal(m, v14)
	mBase = m.M
	v16 = F_sdsnew(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-1)))))
	switch v25 & int32(7) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v16
	F_addReplyStatusFormat(m, l0, int32(_a1657), v11)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L16
	}
L5:
	;
	goto L4
L6:
	;
	if v42 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
	v42 = v41
	goto L6
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
	v42 = v38
	goto L6
L9:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
	v42 = v35
	goto L6
L10:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
	v42 = v32
	goto L6
L11:
	;
	v42 = int32(base.Ui32(v25) >> (uint(int32(3)) % 32))
	goto L6
L12:
	;
	v47 = int32(0)
	goto L13
L13:
	;
	v50 = v16 + v47
	v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
	v52 = F_toupper(m, v51)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v50))) = uint8(v52)
	v55 = v47 + int32(1)
	if v55 != v42 {
		v47 = v55
		goto L13
	} else {
		goto L15
	}
L14:
	;
	goto L5
L15:
	;
	goto L14
L16:
	;
	F_sdsfree(m, v16)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v67 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if l2 != 0 {
		goto L42
	} else {
		goto L43
	}
L19:
	;
	v74 = v67
	v76 = int32(0)
	goto L21
L20:
	;
	v144 = int32(3)
	goto L18
L21:
	;
	if v74&int32(3) == int32(0) {
		v99 = v74
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v144 = v76 + int32(4)
	goto L18
L23:
	;
	F_addReplyStatusLength(m, l0, v74, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L39
	}
L24:
	;
	v132 = v124 - v74
	goto L23
L25:
	;
	v103 = v99
	goto L33
L26:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v85 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v88 = v74
	goto L29
L28:
	;
	v132 = v74 - v74
	goto L23
L29:
	;
	v92 = v88 + int32(1)
	if v92&int32(3) == int32(0) {
		v99 = v92
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v97 != 0 {
		v88 = v92
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v124 = v92
	goto L24
L33:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v112 = int32(-2139062144)
	if (int32(16843008)-v109|v109)&v112 == v112 {
		v103 = v103 + int32(4)
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v118 = v103
	goto L36
L35:
	;
	goto L34
L36:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v122 != 0 {
		v118 = v118 + int32(1)
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v124 = v118
	goto L24
L38:
	;
	goto L37
L39:
	;
	v136 = v76 + int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1+v136<<(uint(int32(2))%32))))
	if v140 != 0 {
		v74 = v140
		v76 = v136
		goto L21
	} else {
		goto L40
	}
L40:
	;
	goto L22
L41:
	;
	F_addReplyStatusLength(m, l0, int32(_a1658), int32(4))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L65
	}
L42:
	;
	v152 = int32(0)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v153 == v152 {
		v233 = v152
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v233 = int32(0)
	goto L41
L44:
	;
	v160 = v153
	v162 = v152
	goto L45
L45:
	;
	if v160&int32(3) == int32(0) {
		v185 = v160
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v233 = v222
	goto L41
L47:
	;
	F_addReplyStatusLength(m, l0, v160, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L63
	}
L48:
	;
	v218 = v210 - v160
	goto L47
L49:
	;
	v189 = v185
	goto L57
L50:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v171 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v174 = v160
	goto L53
L52:
	;
	v218 = v160 - v160
	goto L47
L53:
	;
	v178 = v174 + int32(1)
	if v178&int32(3) == int32(0) {
		v185 = v178
		goto L49
	} else {
		goto L55
	}
L55:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v183 != 0 {
		v174 = v178
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v210 = v178
	goto L48
L57:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v198 = int32(-2139062144)
	if (int32(16843008)-v195|v195)&v198 == v198 {
		v189 = v189 + int32(4)
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v204 = v189
	goto L60
L59:
	;
	goto L58
L60:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v208 != 0 {
		v204 = v204 + int32(1)
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v210 = v204
	goto L48
L62:
	;
	goto L61
L63:
	;
	v222 = v162 + int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l2+v222<<(uint(int32(2))%32))))
	if v226 != 0 {
		v160 = v226
		v162 = v222
		goto L45
	} else {
		goto L64
	}
L64:
	;
	goto L46
L65:
	;
	F_addReplyStatusLength(m, l0, int32(_a1659), int32(20))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_setDeferredAggregateLen(m, l0, v18, v144+v233, int32(42))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	m.G0 = v11 + int32(16)
	return
}
func F_addWritePreparedReplyBulkCBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	v6 = m.G0
	v8 = v6 - int32(128)
	m.G0 = v8
	if base.Ui32(int32(31)) < base.Ui32(l2) {
		v25 = int32(36)
		*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v25)
		v28 = v8 | int32(1)
		v30 = base.I64_extend_i32_u(l2)
		if v30 <= int64(-1) {
			v39 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v39)
			v43 = int32(1)
			v48 = v28 + v43
			v49 = int32(126)
			v50 = int64(0) - v30
			v51 = v43
		} else {
			v48 = v28
			v49 = int32(127)
			v50 = v30
			v51 = int32(0)
		}
		v52 = F_ull2string(m, v48, v49, v50)
		mBase = m.M
		if v52 == int32(0) {
			v71 = int32(0)
		} else {
			v71 = v52 + v51
		}
		v75 = int32(2573)
		*(*uint16)(unsafe.Add(mBase, uint32(v71+v8+int32(1)))) = uint16(v75)
		F__addReplyToBufferOrList(m, l0, v8, v71+int32(3))
		mBase = m.M
		v80 = m.ExcPending
		if v80 != 0 {
			return
		} else {
			F__addReplyToBufferOrList(m, l0, l1, l2)
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				F__addReplyToBufferOrList(m, l0, int32(_a823), int32(2))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return
				} else {
					m.G0 = v8 + int32(128)
					return
				}
			}
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_consts[488])))
		v17 = F_objectGetVal(m, v16)
		mBase = m.M
		if base.Ui32(l2) < base.Ui32(int32(10)) {
			v22 = int32(4)
		} else {
			v22 = int32(5)
		}
		F__addReplyToBufferOrList(m, l0, v17, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			F__addReplyToBufferOrList(m, l0, l1, l2)
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				F__addReplyToBufferOrList(m, l0, int32(_a823), int32(2))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return
				} else {
					m.G0 = v8 + int32(128)
					return
				}
			}
		}
	}
}
func F_addZpopInitialReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	if l2|l1 != 0 {
		if l1 != 0 {
			if l1 == int32(0) {
				if l1 == int32(0) {
					return
				} else {
					if l2 == int32(0) {
						return
					} else {
						F_addReplyArrayLen(m, l0, int32(2))
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							F_addReplyBulk(m, l0, l4)
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								F_addReplyArrayLen(m, l0, l3)
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			} else {
				if l2 != 0 {
					if l1 == int32(0) {
						return
					} else {
						if l2 == int32(0) {
							return
						} else {
							F_addReplyArrayLen(m, l0, int32(2))
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								F_addReplyBulk(m, l0, l4)
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									F_addReplyArrayLen(m, l0, l3)
									v31 = m.ExcPending
									if v31 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				} else {
					v13 = int32(1)
					F_addReplyArrayLen(m, l0, l3<<(uint(v13)%32)|v13)
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						F_addReplyBulk(m, l0, l4)
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			if l2 != 0 {
				F_addReplyArrayLen(m, l0, l3)
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					return
				}
			} else {
				if l1 == int32(0) {
					if l1 == int32(0) {
						return
					} else {
						if l2 == int32(0) {
							return
						} else {
							F_addReplyArrayLen(m, l0, int32(2))
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								F_addReplyBulk(m, l0, l4)
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									F_addReplyArrayLen(m, l0, l3)
									v31 = m.ExcPending
									if v31 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				} else {
					if l2 != 0 {
						if l1 == int32(0) {
							return
						} else {
							if l2 == int32(0) {
								return
							} else {
								F_addReplyArrayLen(m, l0, int32(2))
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									F_addReplyBulk(m, l0, l4)
									v29 = m.ExcPending
									if v29 != 0 {
										return
									} else {
										F_addReplyArrayLen(m, l0, l3)
										v31 = m.ExcPending
										if v31 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					} else {
						v13 = int32(1)
						F_addReplyArrayLen(m, l0, l3<<(uint(v13)%32)|v13)
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							F_addReplyBulk(m, l0, l4)
							v20 = m.ExcPending
							if v20 != 0 {
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
		F_addReplyArrayLen(m, l0, l3<<(uint(int32(1))%32))
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	}
}
func F_addk(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 float64
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v143 int32
	_ = v143
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_luaH_set(m, v12, v13, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		if v18 != int32(3) {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v14))) = base.F64_convert_i32_s(v33)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v40 = v31 + int32(40)
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
			if v38 < v41 {
				v54 = v41
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
				if v54 <= v32 {
				} else {
					v59 = (v54 - v32) & int32(7)
					if v59 == int32(0) {
						v91 = v32
					} else {
						v64 = int32(0)
						v70 = v32
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v55+v70<<(uint(int32(4))%32))+8)) = int32(0)
							v79 = int32(1)
							v80 = v70 + v79
							v82 = v64 + v79
							if v82 != v59 {
								v64 = v82
								v70 = v80
								continue
							} else {
								break
							}
							break
						}
						v91 = v80
					}
					if base.Ui32(int32(-8)) < base.Ui32(v32-v54) {
					} else {
						v105 = v91
						for {
							v111 = v55 + v105<<(uint(int32(4))%32)
							v112 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = v112
							*(*int32)(unsafe.Add(mBase, uint32(v111+int32(24)))) = v112
							*(*int32)(unsafe.Add(mBase, uint32(v111+int32(40)))) = v112
							*(*int32)(unsafe.Add(mBase, uint32(v111+int32(56)))) = v112
							*(*int32)(unsafe.Add(mBase, uint32(v111+int32(72)))) = v112
							*(*int32)(unsafe.Add(mBase, uint32(v111+int32(88)))) = v112
							*(*int32)(unsafe.Add(mBase, uint32(v111+int32(104)))) = v112
							*(*int32)(unsafe.Add(mBase, uint32(v111+int32(120)))) = v112
							v143 = v105 + int32(8)
							if v143 != v54 {
								v105 = v143
								continue
							} else {
								break
							}
							break
						}
					}
				}
				v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v157 = int32(4)
				v159 = v55 + v156<<(uint(v157)%32)
				v160 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				*(*int64)(unsafe.Add(mBase, uint32(v159))) = v160
				v162 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v159)+8)) = v162
				if v162 < v157 {
				} else {
					v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+5)))
					if v167&int32(3) == int32(0) {
					} else {
						v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+5)))
						if v172&int32(4) == int32(0) {
						} else {
							v177 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
							v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+21)))
							if v178 != int32(1) {
								v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+20)))
								v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+5)))
								v188 = v182&int32(3) | v185&int32(248)
								*(*uint8)(unsafe.Add(mBase, uint32(v31)+5)) = uint8(v188)
							} else {
								F_reallymarkobject(m, v177, v166)
								mBase = m.M
							}
						}
					}
				}
				v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v191 + int32(1)
				return v191
			} else {
				v43 = m.G3
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
				v49 = F_luaM_growaux_(m, v12, v44, v40, int32(16), int32(262143), v43+int32(_a2675))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v49
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
					v54 = v52
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					if v54 <= v32 {
					} else {
						v59 = (v54 - v32) & int32(7)
						if v59 == int32(0) {
							v91 = v32
						} else {
							v64 = int32(0)
							v70 = v32
							for {
								*(*int32)(unsafe.Add(mBase, uint32(v55+v70<<(uint(int32(4))%32))+8)) = int32(0)
								v79 = int32(1)
								v80 = v70 + v79
								v82 = v64 + v79
								if v82 != v59 {
									v64 = v82
									v70 = v80
									continue
								} else {
									break
								}
								break
							}
							v91 = v80
						}
						if base.Ui32(int32(-8)) < base.Ui32(v32-v54) {
						} else {
							v105 = v91
							for {
								v111 = v55 + v105<<(uint(int32(4))%32)
								v112 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = v112
								*(*int32)(unsafe.Add(mBase, uint32(v111+int32(24)))) = v112
								*(*int32)(unsafe.Add(mBase, uint32(v111+int32(40)))) = v112
								*(*int32)(unsafe.Add(mBase, uint32(v111+int32(56)))) = v112
								*(*int32)(unsafe.Add(mBase, uint32(v111+int32(72)))) = v112
								*(*int32)(unsafe.Add(mBase, uint32(v111+int32(88)))) = v112
								*(*int32)(unsafe.Add(mBase, uint32(v111+int32(104)))) = v112
								*(*int32)(unsafe.Add(mBase, uint32(v111+int32(120)))) = v112
								v143 = v105 + int32(8)
								if v143 != v54 {
									v105 = v143
									continue
								} else {
									break
								}
								break
							}
						}
					}
					v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v157 = int32(4)
					v159 = v55 + v156<<(uint(v157)%32)
					v160 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					*(*int64)(unsafe.Add(mBase, uint32(v159))) = v160
					v162 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v159)+8)) = v162
					if v162 < v157 {
					} else {
						v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+5)))
						if v167&int32(3) == int32(0) {
						} else {
							v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+5)))
							if v172&int32(4) == int32(0) {
							} else {
								v177 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
								v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+21)))
								if v178 != int32(1) {
									v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+20)))
									v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+5)))
									v188 = v182&int32(3) | v185&int32(248)
									*(*uint8)(unsafe.Add(mBase, uint32(v31)+5)) = uint8(v188)
								} else {
									F_reallymarkobject(m, v177, v166)
									mBase = m.M
								}
							}
						}
					}
					v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v191 + int32(1)
					return v191
				}
			}
		} else {
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
			if base.F64_lt(base.F64_abs(v21), float64(2.147483648e+09)) == int32(0) {
				return int32(-2147483648)
			} else {
				v27 = base.I32_trunc_f64_s(v21)
				return v27
			}
		}
	}
}
func F_afterCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	F_postExecutionUnitOperations(m)
	mBase = m.M
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		F_trackingHandlePendingKeyInvalidations(m)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			F_clusterSlotStatsAddNetworkBytesOutForUserClient(m, l0)
			mBase = m.M
			v7 = m.ExcPending
			if v7 != 0 {
				return
			} else {
				v9 = *(*int32)(unsafe.Add(mBase, _consts[177]))
				if v9 != 0 {
				} else {
					v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					v11 = int32(0)
					v12 = *(*int32)(unsafe.Add(mBase, _consts[485]))
					v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					if v16 == v11 {
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v20
						if v20 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v19
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v19
						}
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v26
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v28 + v16
						*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12))) = int64(0)
					}
				}
				return
			}
		}
	}
}
func F_allocatorDefragGetFragSmallbins(m *base.Module) int32 {
	return int32(0)
}
func F_applyTlsCfg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v28 int32
	_ = v28
	v3 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v3 != 0 {
		v14 = F_connectionTypeTls(m)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
			v19 = m.T0[v18].(func(*base.Module, int32, int32) int32)(m, int32(_a526), int32(1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 != int32(-1) {
					F_tlsResetCertInfo(m)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						return int32(1)
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a527)
					return int32(0)
				}
			}
		}
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, _consts[170]))
		if v5 != 0 {
			v14 = F_connectionTypeTls(m)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
				v19 = m.T0[v18].(func(*base.Module, int32, int32) int32)(m, int32(_a526), int32(1))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v19 != int32(-1) {
						F_tlsResetCertInfo(m)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							return int32(1)
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a527)
						return int32(0)
					}
				}
			}
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, _consts[107]))
			if v7 == int32(0) {
				F_tlsResetCertInfo(m)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					return int32(1)
				}
			} else {
				v14 = F_connectionTypeTls(m)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
					v19 = m.T0[v18].(func(*base.Module, int32, int32) int32)(m, int32(_a526), int32(1))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						if v19 != int32(-1) {
							F_tlsResetCertInfo(m)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								return int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a527)
							return int32(0)
						}
					}
				}
			}
		}
	}
}
func F_arrayEndCallback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2 + int32(-1)
	F_processCollectionElementEnd(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_arrayStartCallback(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 < int32(0) {
		v31 = v6
	} else {
		v12 = l0 + v7<<(uint(int32(2))%32)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
		v14 = int32(1)
		v15 = v13 + v14
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v15
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1040))
		if v17 != v14 {
			v31 = v6
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v22))) = base.F64_convert_i32_u(v15)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v26 + int32(16)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v31 = v30
		}
	}
	v35 = F_lua_checkstack(m, v31, int32(3))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		return
	} else {
		if v35 != 0 {
			F_lua_createtable(m, v31, l1, int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v44 = int32(1)
				v45 = v43 + v44
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v45
				v49 = l0 + v45<<(uint(int32(2))%32)
				*(*int32)(unsafe.Add(mBase, uint32(v49+int32(1040)))) = v44
				*(*int32)(unsafe.Add(mBase, uint32(v49+int32(16)))) = int32(0)
				return
			}
		} else {
			F__serverPanic_2(m, int32(1026))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_asin(m *base.Module, l0 float64) float64 {
	var v7 int64
	_ = v7
	var v12 int32
	_ = v12
	var v34 float64
	_ = v34
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v105 float64
	_ = v105
	var v110 float64
	_ = v110
	var v115 float64
	_ = v115
	var v119 float64
	_ = v119
	var v128 float64
	_ = v128
	var v135 float64
	_ = v135
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	v7 = base.I64_reinterpret_f64(l0)
	v12 = base.I32_wrap_i64(int64(base.Ui64(v7)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v12) < base.Ui32(int32(1072693248)) {
		if base.Ui32(int32(1071644671)) < base.Ui32(v12) {
			v73 = base.F64_mul(base.F64_sub(float64(1), base.F64_abs(l0)), float64(0.5))
			v74 = base.F64_sqrt(v73)
			v105 = base.F64_div(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), float64(1)))
			if base.Ui32(v12) < base.Ui32(int32(1072640819)) {
				v115 = float64(0.7853981633974483)
				v119 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v74) & int64(-4294967296))
				v128 = base.F64_div(base.F64_sub(v73, base.F64_mul(v119, v119)), base.F64_add(v74, v119))
				v135 = base.F64_add(base.F64_sub(base.F64_sub(v115, base.F64_add(v119, v119)), base.F64_sub(base.F64_mul(base.F64_add(v74, v74), v105), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v128, v128)))), v115)
			} else {
				v110 = base.F64_add(base.F64_mul(v74, v105), v74)
				v135 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v110, v110), float64(-6.123233995736766e-17)))
			}
			if v7 < int64(0) {
				v140 = base.F64_neg(v135)
			} else {
				v140 = v135
			}
			v141 = v140
			return v141
		} else {
			if base.Ui32(v12+int32(-1048576)) < base.Ui32(int32(1044381696)) {
				v141 = l0
				return v141
			} else {
				v34 = base.F64_mul(l0, l0)
				return base.F64_add(base.F64_mul(l0, base.F64_div(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), float64(1)))), l0)
			}
		}
	} else {
		if v12+int32(-1072693248)|base.I32_wrap_i64(v7) != 0 {
			return base.F64_div(float64(0), base.F64_sub(l0, l0))
		} else {
			return base.F64_add(base.F64_mul(l0, float64(1.5707963267948966)), float64(7.52316384526264e-37))
		}
	}
}
func F_askingCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v3 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v7 | int32(512)
		v12 = *(*int32)(unsafe.Add(mBase, _consts[27]))
		F_addReply(m, l0, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			return
		}
	} else {
		F_addReplyError(m, l0, int32(_a209))
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_assignment(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v14+int32(-6)) < base.Ui32(int32(4)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v26 + int32(-44) {
	case 0:
		goto L11
	default:
		goto L10
	case 17:
		goto L9
	}
L2:
	;
	v19 = m.G3
	F_luaX_syntaxerror(m, l0, v19+int32(_a2681))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	goto L1
L5:
	;
	F_luaK_storevar(m, v266, l1+int32(8), v12+int32(48))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L3
	} else {
		goto L58
	}
L6:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(12)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v258 + int32(-1)
	v266 = v257
	goto L5
L7:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202)+52)))
	v204 = int32(200) - v203
	if l2 <= v204 {
		goto L49
	} else {
		goto L50
	}
L8:
	;
	v184 = int32(0)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v187 = F_luaK_codeABC(m, v39, v184, v185, v183, v184)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L3
	} else {
		goto L47
	}
L9:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L3
	} else {
		goto L29
	}
L10:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v75 = F_luaX_token2str(m, l0, int32(61))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L3
	} else {
		goto L26
	}
L11:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l1
	F_primaryexp(m, l0, v12+int32(56))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	if v36 != int32(6) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v43 = l1
	v49 = int32(0)
	goto L15
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if v51 != int32(9) {
		v64 = v49
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v40
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v72 != 0 {
		v43 = v72
		v49 = int32(1)
		goto L15
	} else {
		goto L25
	}
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v66 != 0 {
		v43 = v66
		v49 = v64
		goto L15
	} else {
		goto L23
	}
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	if v54 != v55 {
		v60 = v49
		v61 = v55
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v62 == v61 {
		goto L17
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v40
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v60 = int32(1)
	v61 = v59
	goto L20
L22:
	;
	v64 = v60
	goto L18
L23:
	;
	if v64 == int32(0) {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v183 = v69
	goto L8
L25:
	;
	v183 = v61
	goto L8
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v75
	v78 = m.G3
	v81 = F_luaO_pushfstring(m, v73, v78+int32(_a2677), v12)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	F_luaX_syntaxerror(m, l0, v81)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	goto L9
L29:
	;
	v91 = F_subexpr(m, l0, v12+int32(48), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v94 != int32(44) {
		v125 = int32(1)
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if l2 == v125 {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v99 = int32(1)
	goto L33
L33:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L35
	}
L34:
	;
	v125 = v120
	goto L31
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_exp2nextreg(m, v109, v12+int32(48))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v117 = F_subexpr(m, l0, v12+int32(48), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v120 = v99 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v121 == int32(44) {
		v99 = v120
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v146 = v12 + int32(48)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	switch v147 + int32(-13) {
	case 0:
		goto L46
	case 1:
		goto L45
	default:
		goto L44
	}
L40:
	;
	F_adjust_assign(m, l0, l2, v125, v12+int32(48))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	if v125 <= l2 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+36)) = l2 - v125 + v141
	goto L6
L43:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v266 = v180
	goto L5
L44:
	;
	goto L43
L45:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	v169 = v165 + v166<<(uint(int32(2))%32)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v170&int32(8388607) | int32(16777216)
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = int32(11)
	goto L44
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = int32(12)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153+v154<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v146)+8)) = int32(base.Ui32(v158)>>(uint(int32(6))%32)) & int32(255)
	goto L43
L47:
	;
	F_luaK_reserveregs(m, v39, int32(1))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	goto L7
L49:
	;
	F_assignment(m, l0, v12+int32(48), l2+int32(1))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L3
	} else {
		goto L57
	}
L50:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+16))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+60))
	if v209 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v206)+12))
	F_luaX_lexerror(m, v234, v233, int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L3
	} else {
		goto L56
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v209
	v223 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v223 + int32(_a2682)
	v231 = F_luaO_pushfstring(m, v207, v223+int32(_a2683), v12+int32(32))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L3
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v204
	v211 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v211 + int32(_a2682)
	v219 = F_luaO_pushfstring(m, v207, v211+int32(_a2684), v12+int32(16))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v233 = v219
	goto L51
L55:
	;
	v233 = v231
	goto L51
L56:
	;
	goto L49
L57:
	;
	goto L6
L58:
	;
	m.G0 = v12 + int32(80)
	return
}
func F_atan(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v23 float64
	_ = v23
	var v30 float64
	_ = v30
	var v61 float64
	_ = v61
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v78 float64
	_ = v78
	var v95 float64
	_ = v95
	var v103 int32
	_ = v103
	var v106 float64
	_ = v106
	var v111 float64
	_ = v111
	var v114 float64
	_ = v114
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	v8 = base.I64_reinterpret_f64(l0)
	v13 = base.I32_wrap_i64(int64(base.Ui64(v8)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v13) < base.Ui32(int32(1141899264)) {
		if base.Ui32(int32(1071382527)) < base.Ui32(v13) {
			v30 = base.F64_abs(l0)
			if base.Ui32(int32(1072889855)) < base.Ui32(v13) {
				if base.Ui32(int32(1073971199)) < base.Ui32(v13) {
					v61 = base.F64_div(float64(-1), v30)
					v62 = int32(3)
				} else {
					v61 = base.F64_div(base.F64_add(v30, float64(-1.5)), base.F64_add(base.F64_mul(v30, float64(1.5)), float64(1)))
					v62 = int32(2)
				}
			} else {
				if base.Ui32(int32(1072037887)) < base.Ui32(v13) {
					v61 = base.F64_div(base.F64_add(v30, float64(-1)), base.F64_add(v30, float64(1)))
					v62 = int32(1)
				} else {
					v61 = base.F64_div(base.F64_add(base.F64_add(v30, v30), float64(-1)), base.F64_add(v30, float64(2)))
					v62 = int32(0)
				}
			}
			v63 = base.F64_mul(v61, v61)
			v64 = base.F64_mul(v63, v63)
			v78 = base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
			v95 = base.F64_mul(v63, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
			if base.Ui32(int32(1071382527)) < base.Ui32(v13) {
				v103 = v62 << (uint(int32(3)) % 32)
				v106 = *(*float64)(unsafe.Add(mBase, uint32(v103)+uint32(_consts[1004])))
				v111 = *(*float64)(unsafe.Add(mBase, uint32(v103)+uint32(_consts[1005])))
				v114 = base.F64_sub(v106, base.F64_sub(base.F64_sub(base.F64_mul(v61, base.F64_add(v78, v95)), v111), v61))
				if v8 < int64(0) {
					v118 = base.F64_neg(v114)
				} else {
					v118 = v114
				}
				v119 = v118
				return v119
			} else {
				return base.F64_sub(v61, base.F64_mul(v61, base.F64_add(v78, v95)))
			}
		} else {
			if base.Ui32(int32(1044381696)) <= base.Ui32(v13) {
				v61 = l0
				v62 = int32(-1)
				v63 = base.F64_mul(v61, v61)
				v64 = base.F64_mul(v63, v63)
				v78 = base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
				v95 = base.F64_mul(v63, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
				if base.Ui32(int32(1071382527)) < base.Ui32(v13) {
					v103 = v62 << (uint(int32(3)) % 32)
					v106 = *(*float64)(unsafe.Add(mBase, uint32(v103)+uint32(_consts[1004])))
					v111 = *(*float64)(unsafe.Add(mBase, uint32(v103)+uint32(_consts[1005])))
					v114 = base.F64_sub(v106, base.F64_sub(base.F64_sub(base.F64_mul(v61, base.F64_add(v78, v95)), v111), v61))
					if v8 < int64(0) {
						v118 = base.F64_neg(v114)
					} else {
						v118 = v114
					}
					v119 = v118
					return v119
				} else {
					return base.F64_sub(v61, base.F64_mul(v61, base.F64_add(v78, v95)))
				}
			} else {
				v119 = l0
				return v119
			}
		}
	} else {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) {
			v23 = l0
		} else {
			v23 = base.F64_copysign(float64(1.5707963267948966), l0)
		}
		return v23
	}
}
func F_atan2(m *base.Module, l0 float64, l1 float64) float64 {
	mBase := m.M
	_ = mBase
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 float64
	_ = v30
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v66 float64
	_ = v66
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v103 float64
	_ = v103
	var v105 float64
	_ = v105
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l1)&int64(9223372036854775807)) {
		return base.F64_add(l0, l1)
	} else {
		if base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			v22 = base.I64_reinterpret_f64(l1)
			v25 = base.I32_wrap_i64(int64(base.Ui64(v22) >> (uint(int64(32)) % 64)))
			v28 = base.I32_wrap_i64(v22)
			if v25+int32(-1072693248)|v28 != 0 {
				v35 = int32(base.Ui32(v25)>>(uint(int32(30))%32)) & int32(2)
				v36 = base.I64_reinterpret_f64(l0)
				v40 = v35 | base.I32_wrap_i64(int64(base.Ui64(v36)>>(uint(int64(63))%64)))
				v45 = base.I32_wrap_i64(int64(base.Ui64(v36)>>(uint(int64(32))%64))) & int32(2147483647)
				if v45|base.I32_wrap_i64(v36) != 0 {
					v53 = v25 & int32(2147483647)
					if v53|v28 != 0 {
						if v53 != int32(2146435072) {
							if v45 == int32(2146435072) {
								return base.F64_copysign(float64(1.5707963267948966), l0)
							} else {
								if base.Ui32(v45) <= base.Ui32(v53+int32(67108864)) {
									if v35 == int32(0) {
										v85 = F_atan(m, base.F64_abs(base.F64_div(l0, l1)))
										mBase = m.M
										v86 = v85
									} else {
										if base.Ui32(v45+int32(67108864)) < base.Ui32(v53) {
											v86 = float64(0)
										} else {
											v85 = F_atan(m, base.F64_abs(base.F64_div(l0, l1)))
											mBase = m.M
											v86 = v85
										}
									}
									switch v40 {
									default:
										v105 = v86
										return v105
									case 1:
										return base.F64_neg(v86)
									case 2:
										return base.F64_sub(float64(3.141592653589793), base.F64_add(v86, float64(-1.2246467991473532e-16)))
									case 3:
										return base.F64_add(base.F64_add(v86, float64(-1.2246467991473532e-16)), float64(-3.141592653589793))
									}
								} else {
									return base.F64_copysign(float64(1.5707963267948966), l0)
								}
							}
						} else {
							if v45 != int32(2146435072) {
								v103 = *(*float64)(unsafe.Add(mBase, uint32(v40<<(uint(int32(3))%32))+uint32(_consts[1006])))
								v105 = v103
								return v105
							} else {
								v66 = *(*float64)(unsafe.Add(mBase, uint32(v40<<(uint(int32(3))%32))+uint32(_consts[1007])))
								return v66
							}
						}
					} else {
						return base.F64_copysign(float64(1.5707963267948966), l0)
					}
				} else {
					switch v40 {
					default:
						v105 = l0
						return v105
					case 2:
						return float64(3.141592653589793)
					case 3:
						return float64(-3.141592653589793)
					}
				}
			} else {
				v30 = F_atan(m, l0)
				mBase = m.M
				return v30
			}
		} else {
			return base.F64_add(l0, l1)
		}
	}
}
func F_atoi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	v5 = l0
	for {
		v10 = v5 + int32(1)
		v11 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5))))
		if base.B2i32(v11 == int32(32))|base.B2i32(base.Ui32(v11+int32(-9)) < base.Ui32(int32(5))) != 0 {
			v5 = v10
			continue
		} else {
			break
		}
		break
	}
	v19 = int32(1)
	switch v11&int32(255) + int32(-43) {
	case 0:
		v25 = v19
		v26 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
		v27 = v10
		v28 = v26
		v29 = v25
	default:
		v27 = v5
		v28 = v11
		v29 = v19
	case 2:
		v25 = int32(0)
		v26 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
		v27 = v10
		v28 = v26
		v29 = v25
	}
	v32 = v28 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v32) {
		v50 = int32(0)
	} else {
		v36 = int32(0)
		v37 = v27
		v38 = v32
		for {
			v40 = int32(10)
			v42 = v36*v40 - v38
			v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(v37)+1)))
			v47 = v43 + int32(-48)
			if base.Ui32(v47) < base.Ui32(v40) {
				v36 = v42
				v37 = v37 + int32(1)
				v38 = v47
				continue
			} else {
				break
			}
			break
		}
		v50 = v42
	}
	if v29 != 0 {
		v56 = int32(0) - v50
	} else {
		v56 = v50
	}
	return v56
}
func F_authCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 < int32(4) {
		F_redactClientCommandArgument(m, l0, int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v19 != int32(2) {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				F_redactClientCommandArgument(m, l0, int32(2))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					v42 = v38
					v43 = v37
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
					v48 = F_ACLAuthenticateUser(m, l0, v42, v43, v7+int32(12))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						switch v48 {
						case 0:
							v51 = *(*int32)(unsafe.Add(mBase, _consts[27]))
							F_addReply(m, l0, v51)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
								v65 = v64
								if v65 == int32(0) {
									m.G0 = v7 + int32(16)
									return
								} else {
									F_decrRefCount(m, v65)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							}
						case 1:
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+207)))
							if v55&int32(64) != 0 {
								v65 = v54
								if v65 == int32(0) {
									m.G0 = v7 + int32(16)
									return
								} else {
									F_decrRefCount(m, v65)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							} else {
								if v54 != 0 {
									v59 = F_objectGetVal(m, v54)
									mBase = m.M
									v60 = v59
								} else {
									v60 = int32(_a101)
								}
								F_addReplyError(m, l0, v60)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
									v65 = v64
									if v65 == int32(0) {
										m.G0 = v7 + int32(16)
										return
									} else {
										F_decrRefCount(m, v65)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								}
							}
						default:
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							v65 = v64
							if v65 == int32(0) {
								m.G0 = v7 + int32(16)
								return
							} else {
								F_decrRefCount(m, v65)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					}
				}
			} else {
				v22 = int32(0)
				v23 = *(*int32)(unsafe.Add(mBase, _consts[22]))
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
				if v24&int32(4) == v22 {
					v33 = *(*int32)(unsafe.Add(mBase, _consts[32]))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
					v42 = v33
					v43 = v35
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
					v48 = F_ACLAuthenticateUser(m, l0, v42, v43, v7+int32(12))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						switch v48 {
						case 0:
							v51 = *(*int32)(unsafe.Add(mBase, _consts[27]))
							F_addReply(m, l0, v51)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
								v65 = v64
								if v65 == int32(0) {
									m.G0 = v7 + int32(16)
									return
								} else {
									F_decrRefCount(m, v65)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							}
						case 1:
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+207)))
							if v55&int32(64) != 0 {
								v65 = v54
								if v65 == int32(0) {
									m.G0 = v7 + int32(16)
									return
								} else {
									F_decrRefCount(m, v65)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							} else {
								if v54 != 0 {
									v59 = F_objectGetVal(m, v54)
									mBase = m.M
									v60 = v59
								} else {
									v60 = int32(_a101)
								}
								F_addReplyError(m, l0, v60)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
									v65 = v64
									if v65 == int32(0) {
										m.G0 = v7 + int32(16)
										return
									} else {
										F_decrRefCount(m, v65)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								}
							}
						default:
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							v65 = v64
							if v65 == int32(0) {
								m.G0 = v7 + int32(16)
								return
							} else {
								F_decrRefCount(m, v65)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					}
				} else {
					F_addReplyError(m, l0, int32(_a102))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[33]))
		F_addReplyErrorObject(m, l0, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_authenticateClientWithUser(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v92 int32
	_ = v92
	v9 = int32(1)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v10&int32(2) != 0 {
		v92 = v9
		return v92
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v13 == int32(0) {
			v92 = v9
			return v92
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+203)))
			if v16&int32(64) != 0 {
				v92 = v9
				return v92
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
				if v19 == int32(0) {
					v35 = v13
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
					if v22 == int32(0) {
						v35 = v13
					} else {
						v25 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						m.T0[v22].(func(*base.Module, int64, int32))(m, v25, v26)
						mBase = m.M
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = int64(0)
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v35 = v33
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v35)+328)) = l1
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+204))
				*(*int32)(unsafe.Add(mBase, uint32(v35)+204)) = v40&int32(-25165825) | int32(_a821) | int32(16777216)
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+108))
				if v56 == int32(0) {
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
					if v59 == int32(0) {
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+204))
						*(*int32)(unsafe.Add(mBase, uint32(v55)+204)) = v62 | int32(262144)
					}
				}
				if l2 == int32(0) {
					v85 = int32(0)
					if l4 == v85 {
						v92 = v85
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v89 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
						*(*int64)(unsafe.Add(mBase, uint32(l4))) = v89
						v92 = v85
					}
					return v92
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+108))
					if v69 != 0 {
						v78 = v69
						*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = l2
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v81
						v85 = int32(0)
						if l4 == v85 {
							v92 = v85
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v89 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
							*(*int64)(unsafe.Add(mBase, uint32(l4))) = v89
							v92 = v85
						}
						return v92
					} else {
						v71 = F_valkey_calloc(m, int32(20))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v68)+108)) = v71
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+108))
							v78 = v77
							*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = l2
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v81
							v85 = int32(0)
							if l4 == v85 {
								v92 = v85
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v89 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
								*(*int64)(unsafe.Add(mBase, uint32(l4))) = v89
								v92 = v85
							}
							return v92
						}
					}
				}
			}
		}
	}
}
