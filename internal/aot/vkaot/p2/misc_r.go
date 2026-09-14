package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_RedisRegisterConnectionTypeUnix(m *base.Module) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_connTypeRegister(m, int32(_a_F_RedisRegisterConnectionTypeUnix_0))
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F___rem_pio2(m *base.Module, l0 float64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v111 float64
	_ = v111
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v135 float64
	_ = v135
	var v138 float64
	_ = v138
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v157 float64
	_ = v157
	var v170 float64
	_ = v170
	var v176 int32
	_ = v176
	var v177 float64
	_ = v177
	var v178 float64
	_ = v178
	var v179 float64
	_ = v179
	var v180 float64
	_ = v180
	var v183 int32
	_ = v183
	var v194 float64
	_ = v194
	var v195 float64
	_ = v195
	var v200 float64
	_ = v200
	var v201 float64
	_ = v201
	var v213 float64
	_ = v213
	var v214 float64
	_ = v214
	var v219 float64
	_ = v219
	var v220 float64
	_ = v220
	var v222 float64
	_ = v222
	var v223 float64
	_ = v223
	var v224 float64
	_ = v224
	var v231 float64
	_ = v231
	var v236 int32
	_ = v236
	var v247 float64
	_ = v247
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 float64
	_ = v268
	var v272 float64
	_ = v272
	var v282 int32
	_ = v282
	var v298 float64
	_ = v298
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 float64
	_ = v311
	var v316 float64
	_ = v316
	var v322 float64
	_ = v322
	var v328 int32
	_ = v328
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = base.I64_reinterpret_f64(l0)
	v21 = base.I32_wrap_i64(int64(base.Ui64(v18) >> (uint(int64(32)) % 64)))
	v23 = v21 & int32(2147483647)
	if base.Ui32(int32(1074752122)) < base.Ui32(v23) {
		if base.Ui32(int32(1075594811)) < base.Ui32(v23) {
			if base.Ui32(int32(1094263290)) < base.Ui32(v23) {
				if base.Ui32(v23) < base.Ui32(int32(2146435072)) {
					v236 = v16 + int32(16)
					v247 = base.F64_reinterpret_i64(v18&int64(4503599627370495) | int64(4710765210229538816))
					v251 = v236
					v257 = int32(1)
					for {
						if base.F64_lt(base.F64_abs(v247), float64(2.147483648e+09)) == int32(0) {
							v267 = int32(-2147483648)
						} else {
							v265 = base.I32_trunc_f64_s(v247)
							v267 = v265
						}
						v268 = base.F64_convert_i32_s(v267)
						*(*float64)(unsafe.Add(mBase, uint32(v251))) = v268
						v272 = base.F64_mul(base.F64_sub(v247, v268), float64(1.6777216e+07))
						if v257&int32(1) != 0 {
							v247 = v272
							v251 = v236 | int32(8)
							v257 = int32(0)
							continue
						} else {
							break
						}
						break
					}
					*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v272
					v282 = int32(2)
					for {
						v298 = *(*float64)(unsafe.Add(mBase, uint32(v16+int32(16)+v282<<(uint(int32(3))%32))))
						if base.F64_eq(v298, float64(0)) != 0 {
							v282 = v282 + int32(-1)
							continue
						} else {
							break
						}
						break
					}
					v307 = int32(1)
					v310 = F___rem_pio2_large(m, v16+int32(16), v16, int32(base.Ui32(v23)>>(uint(int32(20))%32))+int32(-1046), v282+v307, v307)
					mBase = m.M
					v311 = *(*float64)(unsafe.Add(mBase, uint32(v16)))
					if int64(-1) < v18 {
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = v311
						v322 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
						*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v322
						v328 = v310
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_neg(v311)
						v316 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
						*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_neg(v316)
						v328 = int32(0) - v310
					}
				} else {
					v231 = base.F64_sub(l0, l0)
					*(*float64)(unsafe.Add(mBase, uint32(l1))) = v231
					*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v231
					v328 = int32(0)
				}
			} else {
				v135 = base.F64_add(base.F64_add(base.F64_mul(l0, float64(0.6366197723675814)), float64(6.755399441055744e+15)), float64(-6.755399441055744e+15))
				v138 = base.F64_add(l0, base.F64_mul(v135, float64(-1.5707963267341256)))
				v140 = base.F64_mul(v135, float64(6.077100506506192e-11))
				v141 = base.F64_sub(v138, v140)
				if base.F64_lt(base.F64_abs(v135), float64(2.147483648e+09)) == int32(0) {
					v151 = int32(-2147483648)
				} else {
					v149 = base.I32_trunc_f64_s(v135)
					v151 = v149
				}
				if base.F64_lt(v141, float64(-0.7853981633974483)) == int32(0) {
					if base.F64_gt(v141, float64(0.7853981633974483)) == int32(0) {
						v176 = v151
						v177 = v135
						v178 = v138
						v179 = v140
					} else {
						v170 = base.F64_add(v135, float64(1))
						v176 = v151 + int32(1)
						v177 = v170
						v178 = base.F64_add(l0, base.F64_mul(v170, float64(-1.5707963267341256)))
						v179 = base.F64_mul(v170, float64(6.077100506506192e-11))
					}
				} else {
					v157 = base.F64_add(v135, float64(-1))
					v176 = v151 + int32(-1)
					v177 = v157
					v178 = base.F64_add(l0, base.F64_mul(v157, float64(-1.5707963267341256)))
					v179 = base.F64_mul(v157, float64(6.077100506506192e-11))
				}
				v180 = base.F64_sub(v178, v179)
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = v180
				v183 = int32(base.Ui32(v23) >> (uint(int32(20)) % 32))
				if v183-base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v180))>>(uint(int64(52))%64)))&int32(2047) < int32(17) {
					v222 = v180
					v223 = v178
					v224 = v179
				} else {
					v194 = base.F64_mul(v177, float64(6.077100506303966e-11))
					v195 = base.F64_sub(v178, v194)
					v200 = base.F64_sub(base.F64_mul(v177, float64(2.0222662487959506e-21)), base.F64_sub(base.F64_sub(v178, v195), v194))
					v201 = base.F64_sub(v195, v200)
					*(*float64)(unsafe.Add(mBase, uint32(l1))) = v201
					if int32(50) <= v183-base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v201))>>(uint(int64(52))%64)))&int32(2047) {
						v213 = base.F64_mul(v177, float64(2.0222662487111665e-21))
						v214 = base.F64_sub(v195, v213)
						v219 = base.F64_sub(base.F64_mul(v177, float64(8.4784276603689e-32)), base.F64_sub(base.F64_sub(v195, v214), v213))
						v220 = base.F64_sub(v214, v219)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = v220
						v222 = v220
						v223 = v214
						v224 = v219
					} else {
						v222 = v201
						v223 = v195
						v224 = v200
					}
				}
				*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_sub(base.F64_sub(v223, v222), v224)
				v328 = v176
			}
		} else {
			if base.Ui32(int32(1075183036)) < base.Ui32(v23) {
				if v23 == int32(1075388923) {
					v135 = base.F64_add(base.F64_add(base.F64_mul(l0, float64(0.6366197723675814)), float64(6.755399441055744e+15)), float64(-6.755399441055744e+15))
					v138 = base.F64_add(l0, base.F64_mul(v135, float64(-1.5707963267341256)))
					v140 = base.F64_mul(v135, float64(6.077100506506192e-11))
					v141 = base.F64_sub(v138, v140)
					if base.F64_lt(base.F64_abs(v135), float64(2.147483648e+09)) == int32(0) {
						v151 = int32(-2147483648)
					} else {
						v149 = base.I32_trunc_f64_s(v135)
						v151 = v149
					}
					if base.F64_lt(v141, float64(-0.7853981633974483)) == int32(0) {
						if base.F64_gt(v141, float64(0.7853981633974483)) == int32(0) {
							v176 = v151
							v177 = v135
							v178 = v138
							v179 = v140
						} else {
							v170 = base.F64_add(v135, float64(1))
							v176 = v151 + int32(1)
							v177 = v170
							v178 = base.F64_add(l0, base.F64_mul(v170, float64(-1.5707963267341256)))
							v179 = base.F64_mul(v170, float64(6.077100506506192e-11))
						}
					} else {
						v157 = base.F64_add(v135, float64(-1))
						v176 = v151 + int32(-1)
						v177 = v157
						v178 = base.F64_add(l0, base.F64_mul(v157, float64(-1.5707963267341256)))
						v179 = base.F64_mul(v157, float64(6.077100506506192e-11))
					}
					v180 = base.F64_sub(v178, v179)
					*(*float64)(unsafe.Add(mBase, uint32(l1))) = v180
					v183 = int32(base.Ui32(v23) >> (uint(int32(20)) % 32))
					if v183-base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v180))>>(uint(int64(52))%64)))&int32(2047) < int32(17) {
						v222 = v180
						v223 = v178
						v224 = v179
					} else {
						v194 = base.F64_mul(v177, float64(6.077100506303966e-11))
						v195 = base.F64_sub(v178, v194)
						v200 = base.F64_sub(base.F64_mul(v177, float64(2.0222662487959506e-21)), base.F64_sub(base.F64_sub(v178, v195), v194))
						v201 = base.F64_sub(v195, v200)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = v201
						if int32(50) <= v183-base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v201))>>(uint(int64(52))%64)))&int32(2047) {
							v213 = base.F64_mul(v177, float64(2.0222662487111665e-21))
							v214 = base.F64_sub(v195, v213)
							v219 = base.F64_sub(base.F64_mul(v177, float64(8.4784276603689e-32)), base.F64_sub(base.F64_sub(v195, v214), v213))
							v220 = base.F64_sub(v214, v219)
							*(*float64)(unsafe.Add(mBase, uint32(l1))) = v220
							v222 = v220
							v223 = v214
							v224 = v219
						} else {
							v222 = v201
							v223 = v195
							v224 = v200
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_sub(base.F64_sub(v223, v222), v224)
					v328 = v176
				} else {
					if v18 < int64(0) {
						v119 = base.F64_add(l0, float64(6.2831853069365025))
						v120 = float64(2.430840202602477e-10)
						v121 = base.F64_add(v119, v120)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = v121
						*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v119, v121), v120)
						v328 = int32(-4)
					} else {
						v109 = base.F64_add(l0, float64(-6.2831853069365025))
						v110 = float64(-2.430840202602477e-10)
						v111 = base.F64_add(v109, v110)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = v111
						*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v109, v111), v110)
						v328 = int32(4)
					}
				}
			} else {
				if v23 == int32(1074977148) {
					v135 = base.F64_add(base.F64_add(base.F64_mul(l0, float64(0.6366197723675814)), float64(6.755399441055744e+15)), float64(-6.755399441055744e+15))
					v138 = base.F64_add(l0, base.F64_mul(v135, float64(-1.5707963267341256)))
					v140 = base.F64_mul(v135, float64(6.077100506506192e-11))
					v141 = base.F64_sub(v138, v140)
					if base.F64_lt(base.F64_abs(v135), float64(2.147483648e+09)) == int32(0) {
						v151 = int32(-2147483648)
					} else {
						v149 = base.I32_trunc_f64_s(v135)
						v151 = v149
					}
					if base.F64_lt(v141, float64(-0.7853981633974483)) == int32(0) {
						if base.F64_gt(v141, float64(0.7853981633974483)) == int32(0) {
							v176 = v151
							v177 = v135
							v178 = v138
							v179 = v140
						} else {
							v170 = base.F64_add(v135, float64(1))
							v176 = v151 + int32(1)
							v177 = v170
							v178 = base.F64_add(l0, base.F64_mul(v170, float64(-1.5707963267341256)))
							v179 = base.F64_mul(v170, float64(6.077100506506192e-11))
						}
					} else {
						v157 = base.F64_add(v135, float64(-1))
						v176 = v151 + int32(-1)
						v177 = v157
						v178 = base.F64_add(l0, base.F64_mul(v157, float64(-1.5707963267341256)))
						v179 = base.F64_mul(v157, float64(6.077100506506192e-11))
					}
					v180 = base.F64_sub(v178, v179)
					*(*float64)(unsafe.Add(mBase, uint32(l1))) = v180
					v183 = int32(base.Ui32(v23) >> (uint(int32(20)) % 32))
					if v183-base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v180))>>(uint(int64(52))%64)))&int32(2047) < int32(17) {
						v222 = v180
						v223 = v178
						v224 = v179
					} else {
						v194 = base.F64_mul(v177, float64(6.077100506303966e-11))
						v195 = base.F64_sub(v178, v194)
						v200 = base.F64_sub(base.F64_mul(v177, float64(2.0222662487959506e-21)), base.F64_sub(base.F64_sub(v178, v195), v194))
						v201 = base.F64_sub(v195, v200)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = v201
						if int32(50) <= v183-base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v201))>>(uint(int64(52))%64)))&int32(2047) {
							v213 = base.F64_mul(v177, float64(2.0222662487111665e-21))
							v214 = base.F64_sub(v195, v213)
							v219 = base.F64_sub(base.F64_mul(v177, float64(8.4784276603689e-32)), base.F64_sub(base.F64_sub(v195, v214), v213))
							v220 = base.F64_sub(v214, v219)
							*(*float64)(unsafe.Add(mBase, uint32(l1))) = v220
							v222 = v220
							v223 = v214
							v224 = v219
						} else {
							v222 = v201
							v223 = v195
							v224 = v200
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_sub(base.F64_sub(v223, v222), v224)
					v328 = v176
				} else {
					if v18 < int64(0) {
						v95 = base.F64_add(l0, float64(4.712388980202377))
						v96 = float64(1.8231301519518578e-10)
						v97 = base.F64_add(v95, v96)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = v97
						*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v95, v97), v96)
						v328 = int32(-3)
					} else {
						v85 = base.F64_add(l0, float64(-4.712388980202377))
						v86 = float64(-1.8231301519518578e-10)
						v87 = base.F64_add(v85, v86)
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = v87
						*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v85, v87), v86)
						v328 = int32(3)
					}
				}
			}
		}
	} else {
		if v21&int32(1048575) == int32(598523) {
			v135 = base.F64_add(base.F64_add(base.F64_mul(l0, float64(0.6366197723675814)), float64(6.755399441055744e+15)), float64(-6.755399441055744e+15))
			v138 = base.F64_add(l0, base.F64_mul(v135, float64(-1.5707963267341256)))
			v140 = base.F64_mul(v135, float64(6.077100506506192e-11))
			v141 = base.F64_sub(v138, v140)
			if base.F64_lt(base.F64_abs(v135), float64(2.147483648e+09)) == int32(0) {
				v151 = int32(-2147483648)
			} else {
				v149 = base.I32_trunc_f64_s(v135)
				v151 = v149
			}
			if base.F64_lt(v141, float64(-0.7853981633974483)) == int32(0) {
				if base.F64_gt(v141, float64(0.7853981633974483)) == int32(0) {
					v176 = v151
					v177 = v135
					v178 = v138
					v179 = v140
				} else {
					v170 = base.F64_add(v135, float64(1))
					v176 = v151 + int32(1)
					v177 = v170
					v178 = base.F64_add(l0, base.F64_mul(v170, float64(-1.5707963267341256)))
					v179 = base.F64_mul(v170, float64(6.077100506506192e-11))
				}
			} else {
				v157 = base.F64_add(v135, float64(-1))
				v176 = v151 + int32(-1)
				v177 = v157
				v178 = base.F64_add(l0, base.F64_mul(v157, float64(-1.5707963267341256)))
				v179 = base.F64_mul(v157, float64(6.077100506506192e-11))
			}
			v180 = base.F64_sub(v178, v179)
			*(*float64)(unsafe.Add(mBase, uint32(l1))) = v180
			v183 = int32(base.Ui32(v23) >> (uint(int32(20)) % 32))
			if v183-base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v180))>>(uint(int64(52))%64)))&int32(2047) < int32(17) {
				v222 = v180
				v223 = v178
				v224 = v179
			} else {
				v194 = base.F64_mul(v177, float64(6.077100506303966e-11))
				v195 = base.F64_sub(v178, v194)
				v200 = base.F64_sub(base.F64_mul(v177, float64(2.0222662487959506e-21)), base.F64_sub(base.F64_sub(v178, v195), v194))
				v201 = base.F64_sub(v195, v200)
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = v201
				if int32(50) <= v183-base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v201))>>(uint(int64(52))%64)))&int32(2047) {
					v213 = base.F64_mul(v177, float64(2.0222662487111665e-21))
					v214 = base.F64_sub(v195, v213)
					v219 = base.F64_sub(base.F64_mul(v177, float64(8.4784276603689e-32)), base.F64_sub(base.F64_sub(v195, v214), v213))
					v220 = base.F64_sub(v214, v219)
					*(*float64)(unsafe.Add(mBase, uint32(l1))) = v220
					v222 = v220
					v223 = v214
					v224 = v219
				} else {
					v222 = v201
					v223 = v195
					v224 = v200
				}
			}
			*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_sub(base.F64_sub(v223, v222), v224)
			v328 = v176
		} else {
			if base.Ui32(int32(1073928572)) < base.Ui32(v23) {
				if v18 < int64(0) {
					v67 = base.F64_add(l0, float64(3.1415926534682512))
					v68 = float64(1.2154201013012384e-10)
					v69 = base.F64_add(v67, v68)
					*(*float64)(unsafe.Add(mBase, uint32(l1))) = v69
					*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v67, v69), v68)
					v328 = int32(-2)
				} else {
					v57 = base.F64_add(l0, float64(-3.1415926534682512))
					v58 = float64(-1.2154201013012384e-10)
					v59 = base.F64_add(v57, v58)
					*(*float64)(unsafe.Add(mBase, uint32(l1))) = v59
					*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v57, v59), v58)
					v328 = int32(2)
				}
			} else {
				if v18 < int64(0) {
					v45 = base.F64_add(l0, float64(1.5707963267341256))
					v46 = float64(6.077100506506192e-11)
					v47 = base.F64_add(v45, v46)
					*(*float64)(unsafe.Add(mBase, uint32(l1))) = v47
					*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v45, v47), v46)
					v328 = int32(-1)
				} else {
					v35 = base.F64_add(l0, float64(-1.5707963267341256))
					v36 = float64(-6.077100506506192e-11)
					v37 = base.F64_add(v35, v36)
					*(*float64)(unsafe.Add(mBase, uint32(l1))) = v37
					*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v35, v37), v36)
					v328 = int32(1)
				}
			}
		}
	}
	m.G0 = v16 + int32(48)
	return v328
}
func F___rem_pio2_large(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v82 int32
	_ = v82
	var v84 float64
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v139 int32
	_ = v139
	var v159 int32
	_ = v159
	var v169 float64
	_ = v169
	var v180 int32
	_ = v180
	var v183 float64
	_ = v183
	var v190 float64
	_ = v190
	var v192 float64
	_ = v192
	var v194 int32
	_ = v194
	var v208 float64
	_ = v208
	var v231 int32
	_ = v231
	var v245 int32
	_ = v245
	var v260 float64
	_ = v260
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v276 float64
	_ = v276
	var v288 float64
	_ = v288
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v302 float64
	_ = v302
	var v305 float64
	_ = v305
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 float64
	_ = v320
	var v321 float64
	_ = v321
	var v323 int32
	_ = v323
	var v337 float64
	_ = v337
	var v351 float64
	_ = v351
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v367 float64
	_ = v367
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 float64
	_ = v380
	var v381 int32
	_ = v381
	var v388 float64
	_ = v388
	var v394 float64
	_ = v394
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 float64
	_ = v404
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v464 int32
	_ = v464
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v522 int32
	_ = v522
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v540 float64
	_ = v540
	var v541 int32
	_ = v541
	var v542 float64
	_ = v542
	var v546 float64
	_ = v546
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v562 float64
	_ = v562
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v575 float64
	_ = v575
	var v576 int32
	_ = v576
	var v595 int32
	_ = v595
	var v597 float64
	_ = v597
	var v606 int32
	_ = v606
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v708 int32
	_ = v708
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v752 int32
	_ = v752
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v781 float64
	_ = v781
	var v786 int32
	_ = v786
	var v796 float64
	_ = v796
	var v807 int32
	_ = v807
	var v810 float64
	_ = v810
	var v817 float64
	_ = v817
	var v819 float64
	_ = v819
	var v821 int32
	_ = v821
	var v835 float64
	_ = v835
	var v852 int32
	_ = v852
	var v856 float64
	_ = v856
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v872 float64
	_ = v872
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v885 float64
	_ = v885
	var v886 int32
	_ = v886
	var v893 float64
	_ = v893
	var v901 float64
	_ = v901
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v916 float64
	_ = v916
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v972 float64
	_ = v972
	var v976 float64
	_ = v976
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v992 float64
	_ = v992
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1005 float64
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1019 int32
	_ = v1019
	var v1028 float64
	_ = v1028
	var v1047 int32
	_ = v1047
	var v1061 int32
	_ = v1061
	var v1078 float64
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1097 float64
	_ = v1097
	var v1108 int32
	_ = v1108
	var v1112 float64
	_ = v1112
	var v1117 float64
	_ = v1117
	var v1119 float64
	_ = v1119
	var v1135 float64
	_ = v1135
	var v1179 float64
	_ = v1179
	var v1187 float64
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1200 float64
	_ = v1200
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1223 float64
	_ = v1223
	var v1224 float64
	_ = v1224
	var v1238 float64
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1251 float64
	_ = v1251
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1274 float64
	_ = v1274
	var v1275 float64
	_ = v1275
	var v1294 int32
	_ = v1294
	var v1305 float64
	_ = v1305
	var v1311 float64
	_ = v1311
	var v1312 float64
	_ = v1312
	var v1339 float64
	_ = v1339
	var v1340 float64
	_ = v1340
	var v1342 float64
	_ = v1342
	var v1345 float64
	_ = v1345
	var v1359 int32
	_ = v1359
	var v1360 float64
	_ = v1360
	var v1378 float64
	_ = v1378
	var v1379 float64
	_ = v1379
	var v1392 float64
	_ = v1392
	var v1404 float64
	_ = v1404
	var v1406 float64
	_ = v1406
	var v1412 int32
	_ = v1412
	var v1421 float64
	_ = v1421
	var v1439 float64
	_ = v1439
	var v1440 float64
	_ = v1440
	var v1453 float64
	_ = v1453
	var v1465 float64
	_ = v1465
	var v1467 float64
	_ = v1467
	var v1468 float64
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1474 int32
	_ = v1474
	var v1484 float64
	_ = v1484
	var v1500 float64
	_ = v1500
	var v1501 float64
	_ = v1501
	var v1517 float64
	_ = v1517
	var v1529 float64
	_ = v1529
	var v1533 float64
	_ = v1533
	v6 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(560)
	m.G0 = v26
	v31 = base.I32_div_s(l2+int32(-3), int32(24))
	if v6 < v31 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = v31
	goto L3
L2:
	;
	v35 = v6
	goto L3
L3:
	;
	v38 = v35*int32(-24) + l2
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l4<<(uint(int32(2))%32))+uint32(_c_F___rem_pio2_large[0])))
	v45 = l3 + int32(-1)
	if v43+v45 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v120 = v38 + int32(-24)
	v121 = int32(0)
	if v121 < v43 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v54 = v35 - v45
	v58 = int32(0)
	goto L6
L6:
	;
	if int32(0) <= v54 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L4
L8:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v26+int32(320)+v58<<(uint(int32(3))%32)))) = v84
	v91 = int32(1)
	v94 = v58 + v91
	if v94 != v43+l3 {
		v54 = v54 + v91
		v58 = v94
		goto L6
	} else {
		goto L11
	}
L9:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v54<<(uint(int32(2))%32))+uint32(_c_F___rem_pio2_large[1])))
	v84 = base.F64_convert_i32_s(v82)
	goto L8
L10:
	;
	v84 = float64(0)
	goto L8
L11:
	;
	goto L7
L12:
	;
	v125 = v43
	goto L14
L13:
	;
	v125 = v121
	goto L14
L14:
	;
	v139 = v121
	goto L15
L15:
	;
	if base.B2i32(l3 < int32(1)) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v231 = int32(48) - v38
	v245 = v43
	goto L25
L17:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v26+v139<<(uint(int32(3))%32)))) = v208
	if base.B2i32(v139 == v125) == int32(0) {
		v139 = v139 + int32(1)
		goto L15
	} else {
		goto L23
	}
L18:
	;
	v159 = int32(0)
	v169 = float64(0)
	goto L20
L19:
	;
	v208 = float64(0)
	goto L17
L20:
	;
	v180 = int32(3)
	v183 = *(*float64)(unsafe.Add(mBase, uint32(l0+v159<<(uint(v180)%32))))
	v190 = *(*float64)(unsafe.Add(mBase, uint32(v26+int32(320)+(v139+v45-v159)<<(uint(v180)%32))))
	v192 = base.F64_add(base.F64_mul(v183, v190), v169)
	v194 = v159 + int32(1)
	if v194 != l3 {
		v159 = v194
		v169 = v192
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v208 = v192
	goto L17
L22:
	;
	goto L21
L23:
	;
	goto L16
L24:
	;
	v972 = float64(1)
	if v957 < int32(1024) {
		goto L149
	} else {
		goto L150
	}
L25:
	;
	v260 = *(*float64)(unsafe.Add(mBase, uint32(v26+v245<<(uint(int32(3))%32))))
	if v245 < int32(1) {
		v337 = v260
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v852 = int32(24) - v38
	if v852 < int32(1024) {
		goto L124
	} else {
		goto L125
	}
L27:
	;
	if v120 < int32(1024) {
		goto L40
	} else {
		goto L41
	}
L28:
	;
	v266 = int32(0)
	v270 = v245
	v276 = v260
	goto L29
L29:
	;
	v288 = base.F64_mul(v276, float64(5.960464477539063e-08))
	if base.F64_lt(base.F64_abs(v288), float64(2.147483648e+09)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v337 = v321
	goto L27
L31:
	;
	v302 = base.F64_convert_i32_s(v296)
	v305 = base.F64_add(base.F64_mul(v302, float64(-1.6777216e+07)), v276)
	if base.F64_lt(base.F64_abs(v305), float64(2.147483648e+09)) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v296 = int32(-2147483648)
	goto L31
L33:
	;
	v294 = base.I32_trunc_f64_s(v288)
	v296 = v294
	goto L31
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(480)+v266<<(uint(int32(2))%32)))) = v313
	v316 = v270 + int32(-1)
	v320 = *(*float64)(unsafe.Add(mBase, uint32(v26+v316<<(uint(int32(3))%32))))
	v321 = base.F64_add(v320, v302)
	v323 = v266 + int32(1)
	if v323 != v245 {
		v266 = v323
		v270 = v316
		v276 = v321
		goto L29
	} else {
		goto L37
	}
L35:
	;
	v313 = int32(-2147483648)
	goto L34
L36:
	;
	v311 = base.I32_trunc_f64_s(v305)
	v313 = v311
	goto L34
L37:
	;
	goto L30
L38:
	;
	goto L55
L39:
	;
	v388 = base.F64_mul(v380, base.F64_reinterpret_i64(base.I64_extend_i32_u(v381+int32(1023))<<(uint(int64(52))%64)))
	goto L38
L40:
	;
	if int32(-1023) < v120 {
		v380 = v337
		v381 = v120
		goto L39
	} else {
		goto L47
	}
L41:
	;
	v351 = base.F64_mul(v337, float64(8.98846567431158e+307))
	if base.Ui32(int32(2047)) <= base.Ui32(v120) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v358 = int32(3069)
	if base.Ui32(v120) < base.Ui32(v358) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v380 = v351
	v381 = v38 + int32(-1047)
	goto L39
L44:
	;
	v361 = v120
	goto L46
L45:
	;
	v361 = v358
	goto L46
L46:
	;
	v380 = base.F64_mul(v351, float64(8.98846567431158e+307))
	v381 = v361 + int32(-2046)
	goto L39
L47:
	;
	v367 = base.F64_mul(v337, float64(2.004168360008973e-292))
	if base.Ui32(v120) <= base.Ui32(int32(-1992)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v374 = int32(-2960)
	if base.Ui32(v374) < base.Ui32(v120) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v380 = v367
	v381 = v38 + int32(945)
	goto L39
L50:
	;
	v377 = v120
	goto L52
L51:
	;
	v377 = v374
	goto L52
L52:
	;
	v380 = base.F64_mul(v367, float64(2.004168360008973e-292))
	v381 = v377 + int32(1938)
	goto L39
L53:
	;
	v404 = base.F64_sub(v394, base.F64_convert_i32_s(v402))
	v406 = base.B2i32(v120 < int32(1))
	if v120 < int32(1) {
		goto L61
	} else {
		goto L62
	}
L54:
	;
	v402 = int32(-2147483648)
	goto L53
L55:
	;
	v394 = base.F64_add(v388, base.F64_mul(base.F64_floor(base.F64_mul(v388, float64(0.125))), float64(-8)))
	if base.F64_lt(base.F64_abs(v394), float64(2.147483648e+09)) == int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v400 = base.I32_trunc_f64_s(v394)
	v402 = v400
	goto L53
L57:
	;
	if base.F64_ne(v597, float64(0)) != 0 {
		goto L98
	} else {
		goto L99
	}
L58:
	;
	v445 = int32(0)
	v447 = int32(1)
	if v245 < v447 {
		v503 = v447
		goto L66
	} else {
		goto L67
	}
L59:
	;
	if base.F64_ge(v404, float64(0.5)) != 0 {
		v443 = v402
		v444 = int32(2)
		goto L58
	} else {
		goto L65
	}
L60:
	;
	if v434 < int32(1) {
		v595 = v433
		v597 = v404
		v606 = v434
		goto L57
	} else {
		goto L64
	}
L61:
	;
	if v120 != 0 {
		goto L59
	} else {
		goto L63
	}
L62:
	;
	v413 = v245<<(uint(int32(2))%32) + (v26 + int32(480)) + int32(-4)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	v415 = v414 >> (uint(v231) % 32)
	v417 = v414 - v415<<(uint(v231)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v413))) = v417
	v433 = v415 + v402
	v434 = v417 >> (uint(int32(47)-v38) % 32)
	goto L60
L63:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v245<<(uint(int32(2))%32)+(v26+int32(480))+int32(-4))))
	v433 = v402
	v434 = v428 >> (uint(int32(23)) % 32)
	goto L60
L64:
	;
	v443 = v433
	v444 = v434
	goto L58
L65:
	;
	v595 = v402
	v597 = v404
	v606 = int32(0)
	goto L57
L66:
	;
	if v120 < int32(1) {
		goto L77
	} else {
		goto L78
	}
L67:
	;
	v452 = v445
	v464 = v445
	goto L68
L68:
	;
	v477 = v26 + int32(480) + v452<<(uint(int32(2))%32)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	if v464 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	v503 = v492
	goto L66
L70:
	;
	v495 = v452 + int32(1)
	if v495 != v245 {
		v452 = v495
		v464 = v493
		goto L68
	} else {
		goto L76
	}
L71:
	;
	v492 = int32(1)
	v493 = int32(0)
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v477))) = v485 - v478
	v492 = int32(0)
	v493 = int32(1)
	goto L70
L73:
	;
	if v478 == int32(0) {
		goto L71
	} else {
		goto L75
	}
L74:
	;
	v485 = int32(16777215)
	goto L72
L75:
	;
	v485 = int32(16777216)
	goto L72
L76:
	;
	goto L69
L77:
	;
	v536 = v443 + int32(1)
	if v444 != int32(2) {
		v595 = v536
		v597 = v404
		v606 = v444
		goto L57
	} else {
		goto L81
	}
L78:
	;
	switch v38 + int32(-25) {
	case 0:
		v522 = int32(8388607)
		goto L79
	case 1:
		goto L80
	default:
		goto L77
	}
L79:
	;
	v529 = v245<<(uint(int32(2))%32) + (v26 + int32(480)) + int32(-4)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	*(*int32)(unsafe.Add(mBase, uint32(v529))) = v530 & v522
	goto L77
L80:
	;
	v522 = int32(4194303)
	goto L79
L81:
	;
	v540 = base.F64_sub(float64(1), v404)
	v541 = int32(2)
	if v503 != 0 {
		v595 = v536
		v597 = v540
		v606 = v541
		goto L57
	} else {
		goto L82
	}
L82:
	;
	v542 = float64(1)
	if v120 < int32(1024) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v595 = v536
	v597 = base.F64_sub(v540, base.F64_mul(v575, base.F64_reinterpret_i64(base.I64_extend_i32_u(v576+int32(1023))<<(uint(int64(52))%64))))
	v606 = v541
	goto L57
L84:
	;
	goto L83
L85:
	;
	if int32(-1023) < v120 {
		v575 = v542
		v576 = v120
		goto L84
	} else {
		goto L92
	}
L86:
	;
	v546 = base.F64_mul(v542, float64(8.98846567431158e+307))
	if base.Ui32(int32(2047)) <= base.Ui32(v120) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v553 = int32(3069)
	if base.Ui32(v120) < base.Ui32(v553) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v575 = v546
	v576 = v38 + int32(-1047)
	goto L84
L89:
	;
	v556 = v120
	goto L91
L90:
	;
	v556 = v553
	goto L91
L91:
	;
	v575 = base.F64_mul(v546, float64(8.98846567431158e+307))
	v576 = v556 + int32(-2046)
	goto L84
L92:
	;
	v562 = base.F64_mul(v542, float64(2.004168360008973e-292))
	if base.Ui32(v120) <= base.Ui32(int32(-1992)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v569 = int32(-2960)
	if base.Ui32(v569) < base.Ui32(v120) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v575 = v562
	v576 = v38 + int32(945)
	goto L84
L95:
	;
	v572 = v120
	goto L97
L96:
	;
	v572 = v569
	goto L97
L97:
	;
	v575 = base.F64_mul(v562, float64(2.004168360008973e-292))
	v576 = v572 + int32(1938)
	goto L84
L98:
	;
	goto L26
L99:
	;
	if v245 <= v43 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v708 = int32(1)
	goto L109
L101:
	;
	v614 = v245
	v618 = int32(0)
	goto L102
L102:
	;
	v638 = v614 + int32(-1)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(480)+v638<<(uint(int32(2))%32))))
	v643 = v642 | v618
	if v43 < v638 {
		v614 = v638
		v618 = v643
		goto L102
	} else {
		goto L104
	}
L103:
	;
	if v643 == int32(0) {
		goto L100
	} else {
		goto L105
	}
L104:
	;
	goto L103
L105:
	;
	v655 = v120
	v658 = v245
	goto L106
L106:
	;
	v671 = v655 + int32(-24)
	v675 = v658 + int32(-1)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(480)+v675<<(uint(int32(2))%32))))
	if v679 == int32(0) {
		v655 = v671
		v658 = v675
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v957 = v671
	v960 = v675
	goto L24
L109:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(480)+(v43-v708)<<(uint(int32(2))%32))))
	if v737 == int32(0) {
		v708 = v708 + int32(1)
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v740 = v708 + v245
	v752 = v245
	goto L112
L111:
	;
	goto L110
L112:
	;
	v766 = v752 + l3
	v770 = int32(1)
	v771 = v752 + v770
	v777 = *(*int32)(unsafe.Add(mBase, uint32((v771+v35)<<(uint(int32(2))%32))+uint32(_c_F___rem_pio2_large[1])))
	*(*float64)(unsafe.Add(mBase, uint32(v26+int32(320)+v766<<(uint(int32(3))%32)))) = base.F64_convert_i32_s(v777)
	v781 = float64(0)
	if l3 < v770 {
		v835 = v781
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v245 = v740
	goto L25
L114:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v26+v771<<(uint(int32(3))%32)))) = v835
	if v771 < v740 {
		v752 = v771
		goto L112
	} else {
		goto L119
	}
L115:
	;
	v786 = int32(0)
	v796 = v781
	goto L116
L116:
	;
	v807 = int32(3)
	v810 = *(*float64)(unsafe.Add(mBase, uint32(l0+v786<<(uint(v807)%32))))
	v817 = *(*float64)(unsafe.Add(mBase, uint32(v26+int32(320)+(v766-v786)<<(uint(v807)%32))))
	v819 = base.F64_add(base.F64_mul(v810, v817), v796)
	v821 = v786 + int32(1)
	if v821 != l3 {
		v786 = v821
		v796 = v819
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v835 = v819
	goto L114
L118:
	;
	goto L117
L119:
	;
	goto L113
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(480)+v940<<(uint(int32(2))%32)))) = v936
	v957 = v939
	v960 = v940
	goto L24
L121:
	;
	if base.F64_lt(base.F64_abs(v893), float64(2.147483648e+09)) == int32(0) {
		goto L145
	} else {
		goto L146
	}
L122:
	;
	if base.F64_ge(v893, float64(1.6777216e+07)) == int32(0) {
		goto L121
	} else {
		goto L137
	}
L123:
	;
	v893 = base.F64_mul(v885, base.F64_reinterpret_i64(base.I64_extend_i32_u(v886+int32(1023))<<(uint(int64(52))%64)))
	goto L122
L124:
	;
	if int32(-1023) < v852 {
		v885 = v597
		v886 = v852
		goto L123
	} else {
		goto L131
	}
L125:
	;
	v856 = base.F64_mul(v597, float64(8.98846567431158e+307))
	if base.Ui32(int32(2047)) <= base.Ui32(v852) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v863 = int32(3069)
	if base.Ui32(v852) < base.Ui32(v863) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v885 = v856
	v886 = v852 + int32(-1023)
	goto L123
L128:
	;
	v866 = v852
	goto L130
L129:
	;
	v866 = v863
	goto L130
L130:
	;
	v885 = base.F64_mul(v856, float64(8.98846567431158e+307))
	v886 = v866 + int32(-2046)
	goto L123
L131:
	;
	v872 = base.F64_mul(v597, float64(2.004168360008973e-292))
	if base.Ui32(v852) <= base.Ui32(int32(-1992)) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v879 = int32(-2960)
	if base.Ui32(v879) < base.Ui32(v852) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v885 = v872
	v886 = v852 + int32(969)
	goto L123
L134:
	;
	v882 = v852
	goto L136
L135:
	;
	v882 = v879
	goto L136
L136:
	;
	v885 = base.F64_mul(v872, float64(2.004168360008973e-292))
	v886 = v882 + int32(1938)
	goto L123
L137:
	;
	v901 = base.F64_mul(v893, float64(5.960464477539063e-08))
	if base.F64_lt(base.F64_abs(v901), float64(2.147483648e+09)) == int32(0) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v916 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v909), float64(-1.6777216e+07)), v893)
	if base.F64_lt(base.F64_abs(v916), float64(2.147483648e+09)) == int32(0) {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v909 = int32(-2147483648)
	goto L138
L140:
	;
	v907 = base.I32_trunc_f64_s(v901)
	v909 = v907
	goto L138
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(480)+v245<<(uint(int32(2))%32)))) = v924
	v936 = v909
	v939 = v38
	v940 = v245 + int32(1)
	goto L120
L142:
	;
	v924 = int32(-2147483648)
	goto L141
L143:
	;
	v922 = base.I32_trunc_f64_s(v916)
	v924 = v922
	goto L141
L144:
	;
	v936 = v935
	v939 = v120
	v940 = v245
	goto L120
L145:
	;
	v935 = int32(-2147483648)
	goto L144
L146:
	;
	v933 = base.I32_trunc_f64_s(v893)
	v935 = v933
	goto L144
L147:
	;
	if v960 < int32(0) {
		goto L162
	} else {
		goto L163
	}
L148:
	;
	goto L147
L149:
	;
	if int32(-1023) < v957 {
		v1005 = v972
		v1006 = v957
		goto L148
	} else {
		goto L156
	}
L150:
	;
	v976 = base.F64_mul(v972, float64(8.98846567431158e+307))
	if base.Ui32(int32(2047)) <= base.Ui32(v957) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v983 = int32(3069)
	if base.Ui32(v957) < base.Ui32(v983) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v1005 = v976
	v1006 = v957 + int32(-1023)
	goto L148
L153:
	;
	v986 = v957
	goto L155
L154:
	;
	v986 = v983
	goto L155
L155:
	;
	v1005 = base.F64_mul(v976, float64(8.98846567431158e+307))
	v1006 = v986 + int32(-2046)
	goto L148
L156:
	;
	v992 = base.F64_mul(v972, float64(2.004168360008973e-292))
	if base.Ui32(v957) <= base.Ui32(int32(-1992)) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v999 = int32(-2960)
	if base.Ui32(v999) < base.Ui32(v957) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v1005 = v992
	v1006 = v957 + int32(969)
	goto L148
L159:
	;
	v1002 = v957
	goto L161
L160:
	;
	v1002 = v999
	goto L161
L161:
	;
	v1005 = base.F64_mul(v992, float64(2.004168360008973e-292))
	v1006 = v1002 + int32(1938)
	goto L148
L162:
	;
	switch l4 {
	case 0:
		goto L181
	case 1, 2:
		goto L180
	case 3:
		goto L182
	default:
		goto L178
	}
L163:
	;
	v1019 = v960
	v1028 = base.F64_mul(v1005, base.F64_reinterpret_i64(base.I64_extend_i32_u(v1006+int32(1023))<<(uint(int64(52))%64)))
	goto L164
L164:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(480)+v1019<<(uint(int32(2))%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v26+v1019<<(uint(int32(3))%32)))) = base.F64_mul(v1028, base.F64_convert_i32_s(v1047))
	if v1019 != 0 {
		v1019 = v1019 + int32(-1)
		v1028 = base.F64_mul(v1028, float64(5.960464477539063e-08))
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v1061 = v960
	goto L167
L166:
	;
	goto L165
L167:
	;
	v1078 = float64(0)
	v1080 = v960 - v1061
	if v43 < v1080 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L162
L169:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v26+int32(160)+v1080<<(uint(int32(3))%32)))) = v1135
	if int32(0) < v1061 {
		v1061 = v1061 + int32(-1)
		goto L167
	} else {
		goto L177
	}
L170:
	;
	v1082 = v43
	goto L172
L171:
	;
	v1082 = v1080
	goto L172
L172:
	;
	if v1082 < int32(0) {
		v1135 = v1078
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v1087 = int32(0)
	v1097 = v1078
	goto L174
L174:
	;
	v1108 = int32(3)
	v1112 = *(*float64)(unsafe.Add(mBase, uint32(v1087<<(uint(v1108)%32))+uint32(_c_F___rem_pio2_large[2])))
	v1117 = *(*float64)(unsafe.Add(mBase, uint32(v26+(v1087+v1061)<<(uint(v1108)%32))))
	v1119 = base.F64_add(base.F64_mul(v1112, v1117), v1097)
	if v1087 != v1082 {
		v1087 = v1087 + int32(1)
		v1097 = v1119
		goto L174
	} else {
		goto L176
	}
L175:
	;
	v1135 = v1119
	goto L169
L176:
	;
	goto L175
L177:
	;
	goto L168
L178:
	;
	m.G0 = v26 + int32(560)
	return v595 & int32(7)
L179:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_neg(v1340)
	v1533 = *(*float64)(unsafe.Add(mBase, uint32(v26)+168))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_neg(v1339)
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_neg(v1533)
	goto L178
L180:
	;
	v1406 = float64(0)
	if v960 < int32(0) {
		v1453 = v1406
		goto L204
	} else {
		goto L205
	}
L181:
	;
	v1345 = float64(0)
	if v960 < int32(0) {
		v1392 = v1345
		goto L196
	} else {
		goto L197
	}
L182:
	;
	v1179 = float64(0)
	if v960 < int32(1) {
		v1339 = v1179
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v1340 = *(*float64)(unsafe.Add(mBase, uint32(v26)+160))
	if v606 != 0 {
		goto L179
	} else {
		goto L195
	}
L184:
	;
	v1187 = *(*float64)(unsafe.Add(mBase, uint32(v26+int32(160)+v960<<(uint(int32(3))%32))))
	v1190 = v960
	v1200 = v1187
	goto L185
L185:
	;
	v1212 = v26 + int32(160)
	v1213 = int32(3)
	v1219 = v1190 + int32(-1)
	v1222 = v1212 + v1219<<(uint(v1213)%32)
	v1223 = *(*float64)(unsafe.Add(mBase, uint32(v1222)))
	v1224 = base.F64_add(v1223, v1200)
	*(*float64)(unsafe.Add(mBase, uint32(v1212+v1190<<(uint(v1213)%32)))) = base.F64_add(v1200, base.F64_sub(v1223, v1224))
	*(*float64)(unsafe.Add(mBase, uint32(v1222))) = v1224
	if base.Ui32(int32(1)) < base.Ui32(v1190) {
		v1190 = v1219
		v1200 = v1224
		goto L185
	} else {
		goto L187
	}
L186:
	;
	if v960 == int32(1) {
		v1339 = v1179
		goto L183
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	v1238 = *(*float64)(unsafe.Add(mBase, uint32(v26+int32(160)+v960<<(uint(int32(3))%32))))
	v1241 = v960
	v1251 = v1238
	goto L189
L189:
	;
	v1263 = v26 + int32(160)
	v1264 = int32(3)
	v1270 = v1241 + int32(-1)
	v1273 = v1263 + v1270<<(uint(v1264)%32)
	v1274 = *(*float64)(unsafe.Add(mBase, uint32(v1273)))
	v1275 = base.F64_add(v1274, v1251)
	*(*float64)(unsafe.Add(mBase, uint32(v1263+v1241<<(uint(v1264)%32)))) = base.F64_add(v1251, base.F64_sub(v1274, v1275))
	*(*float64)(unsafe.Add(mBase, uint32(v1273))) = v1275
	if base.Ui32(int32(2)) < base.Ui32(v1241) {
		v1241 = v1270
		v1251 = v1275
		goto L189
	} else {
		goto L191
	}
L190:
	;
	v1294 = v960
	v1305 = float64(0)
	goto L192
L191:
	;
	goto L190
L192:
	;
	v1311 = *(*float64)(unsafe.Add(mBase, uint32(v26+int32(160)+v1294<<(uint(int32(3))%32))))
	v1312 = base.F64_add(v1305, v1311)
	if int32(2) < v1294 {
		v1294 = v1294 + int32(-1)
		v1305 = v1312
		goto L192
	} else {
		goto L194
	}
L193:
	;
	v1339 = v1312
	goto L183
L194:
	;
	goto L193
L195:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v1340
	v1342 = *(*float64)(unsafe.Add(mBase, uint32(v26)+168))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = v1339
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v1342
	goto L178
L196:
	;
	if v606 != 0 {
		goto L201
	} else {
		goto L202
	}
L197:
	;
	v1359 = v960
	v1360 = v1345
	goto L198
L198:
	;
	v1378 = *(*float64)(unsafe.Add(mBase, uint32(v26+int32(160)+v1359<<(uint(int32(3))%32))))
	v1379 = base.F64_add(v1360, v1378)
	if v1359 != 0 {
		v1359 = v1359 + int32(-1)
		v1360 = v1379
		goto L198
	} else {
		goto L200
	}
L199:
	;
	v1392 = v1379
	goto L196
L200:
	;
	goto L199
L201:
	;
	v1404 = base.F64_neg(v1392)
	goto L203
L202:
	;
	v1404 = v1392
	goto L203
L203:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v1404
	goto L178
L204:
	;
	if v606 != 0 {
		goto L209
	} else {
		goto L210
	}
L205:
	;
	v1412 = v960
	v1421 = v1406
	goto L206
L206:
	;
	v1439 = *(*float64)(unsafe.Add(mBase, uint32(v26+int32(160)+v1412<<(uint(int32(3))%32))))
	v1440 = base.F64_add(v1421, v1439)
	if v1412 != 0 {
		v1412 = v1412 + int32(-1)
		v1421 = v1440
		goto L206
	} else {
		goto L208
	}
L207:
	;
	v1453 = v1440
	goto L204
L208:
	;
	goto L207
L209:
	;
	v1465 = base.F64_neg(v1453)
	goto L211
L210:
	;
	v1465 = v1453
	goto L211
L211:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v1465
	v1467 = *(*float64)(unsafe.Add(mBase, uint32(v26)+160))
	v1468 = base.F64_sub(v1467, v1453)
	v1469 = int32(1)
	if v960 < v1469 {
		v1517 = v1468
		goto L212
	} else {
		goto L213
	}
L212:
	;
	if v606 != 0 {
		goto L217
	} else {
		goto L218
	}
L213:
	;
	v1474 = v1469
	v1484 = v1468
	goto L214
L214:
	;
	v1500 = *(*float64)(unsafe.Add(mBase, uint32(v26+int32(160)+v1474<<(uint(int32(3))%32))))
	v1501 = base.F64_add(v1484, v1500)
	if v1474 != v960 {
		v1474 = v1474 + int32(1)
		v1484 = v1501
		goto L214
	} else {
		goto L216
	}
L215:
	;
	v1517 = v1501
	goto L212
L216:
	;
	goto L215
L217:
	;
	v1529 = base.F64_neg(v1517)
	goto L219
L218:
	;
	v1529 = v1517
	goto L219
L219:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v1529
	goto L178
}
func F_readQueryFromClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v17 int32
	_ = v17
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
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_readQueryFromClient[0]))
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+223)))
	if v13 != 0 {
		goto L1
	} else {
		goto L7
	}
L3:
	;
	v9 = F_trySendReadToIOThreads(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v9 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L2
L7:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+222)))
	if v14 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v17 = int32(0)
	goto L9
L9:
	;
	v21 = F_readToQueryBuf(m, v6)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L1
L11:
	;
	v23 = F_handleReadResult(m, v6)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+200))
	F_beforeNextClient(m, v6)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L18
	}
L13:
	;
	if v23 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v25 = F_processInputBuffer(m, v6)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if v25 == int32(-1) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_trimCommandQueue(m, v6)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	goto L12
L18:
	;
	v39 = base.B2i32(v31&int32(1025) == int32(1))
	if v39&base.B2i32(base.Ui32(v17) < base.Ui32(int32(24)))&v21 != 0 {
		v17 = v17 + v39
		goto L9
	} else {
		goto L19
	}
L19:
	;
	goto L10
}
func F_receiveRDBinBioThreadDualChannel(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_receiveRDBinBioThreadDualChannel[0]))
	if int32(2) < v4 {
		v12 = int32(_a_F_receiveRDBinBioThreadDualChannel_0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_receiveRDBinBioThreadDualChannel[1]))
		v15 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_receiveRDBinBioThreadDualChannel[1])) = v15
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
		v20 = m.T0[v19].(func(*base.Module, int32, int32) int32)(m, v13, v15)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			F_bioCreateSaveRDBToDiskJob(m, v13, int32(1))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F__serverLog(m, int32(2), int32(_a_F_receiveRDBinBioThreadDualChannel_1), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = int32(_a_F_receiveRDBinBioThreadDualChannel_0)
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_receiveRDBinBioThreadDualChannel[1]))
			v15 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_receiveRDBinBioThreadDualChannel[1])) = v15
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
			v20 = m.T0[v19].(func(*base.Module, int32, int32) int32)(m, v13, v15)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_bioCreateSaveRDBToDiskJob(m, v13, int32(1))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_recfield(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v14 != int32(285) {
		F_luaX_next(m, l0)
		mBase = m.M
		v85 = m.ExcPending
		if v85 != 0 {
			return
		} else {
			v89 = F_subexpr(m, l0, v10+int32(104), int32(0))
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				F_luaK_exp2val(m, v91, v10+int32(104))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v96 == int32(93) {
						F_luaX_next(m, l0)
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return
						} else {
							v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v118 + int32(1)
							v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							if v122 == int32(61) {
								F_luaX_next(m, l0)
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return
								} else {
									v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return
									} else {
										v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
										mBase = m.M
										v148 = m.ExcPending
										if v148 != 0 {
											return
										} else {
											v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
											v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return
											} else {
												v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
													m.G0 = v10 + int32(128)
													return
												}
											}
										}
									}
								}
							} else {
								v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v127 = F_luaX_token2str(m, l0, int32(61))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v127
									v130 = m.G3
									v133 = F_luaO_pushfstring(m, v125, v130+int32(_a_F_recfield_0), v10)
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return
									} else {
										F_luaX_syntaxerror(m, l0, v133)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return
										} else {
											F_luaX_next(m, l0)
											mBase = m.M
											v139 = m.ExcPending
											if v139 != 0 {
												return
											} else {
												v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
												mBase = m.M
												v143 = m.ExcPending
												if v143 != 0 {
													return
												} else {
													v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return
													} else {
														v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
														v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
														v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return
														} else {
															v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
																m.G0 = v10 + int32(128)
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
					} else {
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v101 = F_luaX_token2str(m, l0, int32(93))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v101
							v104 = m.G3
							v109 = F_luaO_pushfstring(m, v99, v104+int32(_a_F_recfield_0), v10+int32(64))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return
							} else {
								F_luaX_syntaxerror(m, l0, v109)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return
								} else {
									F_luaX_next(m, l0)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return
									} else {
										v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v118 + int32(1)
										v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v122 == int32(61) {
											F_luaX_next(m, l0)
											mBase = m.M
											v139 = m.ExcPending
											if v139 != 0 {
												return
											} else {
												v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
												mBase = m.M
												v143 = m.ExcPending
												if v143 != 0 {
													return
												} else {
													v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return
													} else {
														v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
														v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
														v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return
														} else {
															v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
																m.G0 = v10 + int32(128)
																return
															}
														}
													}
												}
											}
										} else {
											v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											v127 = F_luaX_token2str(m, l0, int32(61))
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = v127
												v130 = m.G3
												v133 = F_luaO_pushfstring(m, v125, v130+int32(_a_F_recfield_0), v10)
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return
												} else {
													F_luaX_syntaxerror(m, l0, v133)
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return
													} else {
														F_luaX_next(m, l0)
														mBase = m.M
														v139 = m.ExcPending
														if v139 != 0 {
															return
														} else {
															v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return
															} else {
																v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
																mBase = m.M
																v148 = m.ExcPending
																if v148 != 0 {
																	return
																} else {
																	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
																	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
																	v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
																	mBase = m.M
																	v155 = m.ExcPending
																	if v155 != 0 {
																		return
																	} else {
																		v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
																		mBase = m.M
																		v157 = m.ExcPending
																		if v157 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
																			m.G0 = v10 + int32(128)
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
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		if v17 < int32(2147483646) {
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			F_luaX_next(m, l0)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return
			} else {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v77 = F_luaK_stringK(m, v76, v73)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v77
					*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = int32(4)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = int64(-1)
					v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v118 + int32(1)
					v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v122 == int32(61) {
						F_luaX_next(m, l0)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return
							} else {
								v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
								mBase = m.M
								v148 = m.ExcPending
								if v148 != 0 {
									return
								} else {
									v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
									v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
									mBase = m.M
									v155 = m.ExcPending
									if v155 != 0 {
										return
									} else {
										v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
											m.G0 = v10 + int32(128)
											return
										}
									}
								}
							}
						}
					} else {
						v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v127 = F_luaX_token2str(m, l0, int32(61))
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v127
							v130 = m.G3
							v133 = F_luaO_pushfstring(m, v125, v130+int32(_a_F_recfield_0), v10)
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return
							} else {
								F_luaX_syntaxerror(m, l0, v133)
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
									return
								} else {
									F_luaX_next(m, l0)
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return
									} else {
										v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
										mBase = m.M
										v143 = m.ExcPending
										if v143 != 0 {
											return
										} else {
											v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
											mBase = m.M
											v148 = m.ExcPending
											if v148 != 0 {
												return
											} else {
												v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
												v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
												v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return
												} else {
													v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
														m.G0 = v10 + int32(128)
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
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
			if v22 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = int32(2147483645)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v22
				v38 = m.G3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v38 + int32(_a_F_recfield_1)
				v46 = F_luaO_pushfstring(m, v20, v38+int32(_a_F_recfield_2), v10+int32(48))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					v48 = v46
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					F_luaX_lexerror(m, v50, v48, int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v54 == int32(285) {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							F_luaX_next(m, l0)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v77 = F_luaK_stringK(m, v76, v73)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v77
									*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = int32(4)
									*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = int64(-1)
									v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v118 + int32(1)
									v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v122 == int32(61) {
										F_luaX_next(m, l0)
										mBase = m.M
										v139 = m.ExcPending
										if v139 != 0 {
											return
										} else {
											v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
											mBase = m.M
											v143 = m.ExcPending
											if v143 != 0 {
												return
											} else {
												v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return
												} else {
													v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
													v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
													v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return
													} else {
														v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
															m.G0 = v10 + int32(128)
															return
														}
													}
												}
											}
										}
									} else {
										v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v127 = F_luaX_token2str(m, l0, int32(61))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v127
											v130 = m.G3
											v133 = F_luaO_pushfstring(m, v125, v130+int32(_a_F_recfield_0), v10)
											mBase = m.M
											v134 = m.ExcPending
											if v134 != 0 {
												return
											} else {
												F_luaX_syntaxerror(m, l0, v133)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return
												} else {
													F_luaX_next(m, l0)
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
														return
													} else {
														v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return
														} else {
															v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
															mBase = m.M
															v148 = m.ExcPending
															if v148 != 0 {
																return
															} else {
																v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
																v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
																v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return
																} else {
																	v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
																	mBase = m.M
																	v157 = m.ExcPending
																	if v157 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
																		m.G0 = v10 + int32(128)
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
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v59 = F_luaX_token2str(m, l0, int32(285))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v59
								v62 = m.G3
								v67 = F_luaO_pushfstring(m, v57, v62+int32(_a_F_recfield_0), v10+int32(16))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									F_luaX_syntaxerror(m, l0, v67)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										F_luaX_next(m, l0)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v77 = F_luaK_stringK(m, v76, v73)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v77
												*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = int32(4)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = int64(-1)
												v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
												*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v118 + int32(1)
												v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v122 == int32(61) {
													F_luaX_next(m, l0)
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
														return
													} else {
														v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return
														} else {
															v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
															mBase = m.M
															v148 = m.ExcPending
															if v148 != 0 {
																return
															} else {
																v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
																v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
																v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return
																} else {
																	v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
																	mBase = m.M
																	v157 = m.ExcPending
																	if v157 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
																		m.G0 = v10 + int32(128)
																		return
																	}
																}
															}
														}
													}
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
													v127 = F_luaX_token2str(m, l0, int32(61))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10))) = v127
														v130 = m.G3
														v133 = F_luaO_pushfstring(m, v125, v130+int32(_a_F_recfield_0), v10)
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return
														} else {
															F_luaX_syntaxerror(m, l0, v133)
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return
															} else {
																F_luaX_next(m, l0)
																mBase = m.M
																v139 = m.ExcPending
																if v139 != 0 {
																	return
																} else {
																	v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
																	mBase = m.M
																	v143 = m.ExcPending
																	if v143 != 0 {
																		return
																	} else {
																		v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
																		mBase = m.M
																		v148 = m.ExcPending
																		if v148 != 0 {
																			return
																		} else {
																			v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
																			v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
																			v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
																			mBase = m.M
																			v155 = m.ExcPending
																			if v155 != 0 {
																				return
																			} else {
																				v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
																				mBase = m.M
																				v157 = m.ExcPending
																				if v157 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
																					m.G0 = v10 + int32(128)
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
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(2147483645)
				v25 = m.G3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v25 + int32(_a_F_recfield_1)
				v33 = F_luaO_pushfstring(m, v20, v25+int32(_a_F_recfield_3), v10+int32(32))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v48 = v33
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					F_luaX_lexerror(m, v50, v48, int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v54 == int32(285) {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							F_luaX_next(m, l0)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v77 = F_luaK_stringK(m, v76, v73)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v77
									*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = int32(4)
									*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = int64(-1)
									v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v118 + int32(1)
									v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v122 == int32(61) {
										F_luaX_next(m, l0)
										mBase = m.M
										v139 = m.ExcPending
										if v139 != 0 {
											return
										} else {
											v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
											mBase = m.M
											v143 = m.ExcPending
											if v143 != 0 {
												return
											} else {
												v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return
												} else {
													v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
													v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
													v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return
													} else {
														v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
															m.G0 = v10 + int32(128)
															return
														}
													}
												}
											}
										}
									} else {
										v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v127 = F_luaX_token2str(m, l0, int32(61))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v127
											v130 = m.G3
											v133 = F_luaO_pushfstring(m, v125, v130+int32(_a_F_recfield_0), v10)
											mBase = m.M
											v134 = m.ExcPending
											if v134 != 0 {
												return
											} else {
												F_luaX_syntaxerror(m, l0, v133)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return
												} else {
													F_luaX_next(m, l0)
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
														return
													} else {
														v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return
														} else {
															v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
															mBase = m.M
															v148 = m.ExcPending
															if v148 != 0 {
																return
															} else {
																v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
																v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
																v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return
																} else {
																	v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
																	mBase = m.M
																	v157 = m.ExcPending
																	if v157 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
																		m.G0 = v10 + int32(128)
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
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v59 = F_luaX_token2str(m, l0, int32(285))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v59
								v62 = m.G3
								v67 = F_luaO_pushfstring(m, v57, v62+int32(_a_F_recfield_0), v10+int32(16))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									F_luaX_syntaxerror(m, l0, v67)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										F_luaX_next(m, l0)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v77 = F_luaK_stringK(m, v76, v73)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v77
												*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = int32(4)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = int64(-1)
												v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
												*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v118 + int32(1)
												v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v122 == int32(61) {
													F_luaX_next(m, l0)
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
														return
													} else {
														v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return
														} else {
															v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
															mBase = m.M
															v148 = m.ExcPending
															if v148 != 0 {
																return
															} else {
																v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
																v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
																v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return
																} else {
																	v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
																	mBase = m.M
																	v157 = m.ExcPending
																	if v157 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
																		m.G0 = v10 + int32(128)
																		return
																	}
																}
															}
														}
													}
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
													v127 = F_luaX_token2str(m, l0, int32(61))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10))) = v127
														v130 = m.G3
														v133 = F_luaO_pushfstring(m, v125, v130+int32(_a_F_recfield_0), v10)
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return
														} else {
															F_luaX_syntaxerror(m, l0, v133)
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return
															} else {
																F_luaX_next(m, l0)
																mBase = m.M
																v139 = m.ExcPending
																if v139 != 0 {
																	return
																} else {
																	v142 = F_luaK_exp2RK(m, v12, v10+int32(104))
																	mBase = m.M
																	v143 = m.ExcPending
																	if v143 != 0 {
																		return
																	} else {
																		v147 = F_subexpr(m, l0, v10+int32(80), int32(0))
																		mBase = m.M
																		v148 = m.ExcPending
																		if v148 != 0 {
																			return
																		} else {
																			v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
																			v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
																			v154 = F_luaK_exp2RK(m, v12, v10+int32(80))
																			mBase = m.M
																			v155 = m.ExcPending
																			if v155 != 0 {
																				return
																			} else {
																				v156 = F_luaK_codeABC(m, v12, int32(9), v151, v142, v154)
																				mBase = m.M
																				v157 = m.ExcPending
																				if v157 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v13
																					m.G0 = v10 + int32(128)
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
func F_reclaimFilePageCache(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return int32(0)
}
func F_redis_check_aof_main(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int64
	_ = v148
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
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	v5 = m.G0
	v7 = v5 - int32(4144)
	m.G0 = v7
	if l0 < int32(2) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.Env.Exit(m, int32(0))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	F_checkMultiPartAof(m, v298, v157, v158)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L25
	} else {
		goto L104
	}
L3:
	;
	v310 = F_puts(m, int32(_a_F_redis_check_aof_main_0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L25
	} else {
		goto L103
	}
L4:
	;
	if v157&int32(3) == int32(0) {
		v202 = v157
		goto L60
	} else {
		goto L61
	}
L5:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v172
	v175 = F_iprintf(m, int32(_a_F_redis_check_aof_main_1), v7)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L25
	} else {
		goto L57
	}
L6:
	;
	switch l0 + int32(-2) {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	default:
		goto L5
	}
L7:
	;
	v159 = int32(4097)
	v162 = F_memchr(m, v157, int32(0), v159)
	mBase = m.M
	if v162 != 0 {
		goto L52
	} else {
		goto L53
	}
L8:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v107 = int32(_a_F_redis_check_aof_main_2)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_redis_check_aof_main[0])))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v111 == int32(0) {
		v134 = v110
		v135 = v111
		goto L39
	} else {
		goto L40
	}
L9:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v72 = int32(_a_F_redis_check_aof_main_3)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_redis_check_aof_main[1])))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v76 == int32(0) {
		v99 = v75
		v100 = v76
		goto L30
	} else {
		goto L31
	}
L10:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v14 != int32(45) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v57 = F_getVersion(m)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L25
	} else {
		goto L26
	}
L12:
	;
	v23 = int32(0)
	v24 = int32(_a_F_redis_check_aof_main_4)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_redis_check_aof_main[2])))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v28 == v23 {
		v51 = v27
		v52 = v28
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v17 != int32(118) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	if v20 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	if v52-v51&int32(255) != 0 {
		v157 = v13
		v158 = v23
		goto L7
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	if v28 != v27&int32(255) {
		v51 = v27
		v52 = v28
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v34 = v13
	v35 = v24
	goto L20
L20:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v39 == int32(0) {
		v51 = v38
		v52 = v39
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v51 = v38
	v52 = v39
	goto L17
L22:
	;
	v42 = int32(1)
	if v39 == v38&int32(255) {
		v34 = v34 + v42
		v35 = v35 + v42
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L11
L25:
	;
	return int32(0)
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v57
	v65 = F_iprintf(m, int32(_a_F_redis_check_aof_main_5), v7+int32(16))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	F_sdsfree(m, v57)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	m.Env.Exit(m, int32(0))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	if v100-v99&int32(255) != 0 {
		goto L5
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	if v76 != v75&int32(255) {
		v99 = v75
		v100 = v76
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v82 = v71
	v83 = v72
	goto L33
L33:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v87 == int32(0) {
		v99 = v86
		v100 = v87
		goto L30
	} else {
		goto L35
	}
L34:
	;
	v99 = v86
	v100 = v87
	goto L30
L35:
	;
	v90 = int32(1)
	if v87 == v86&int32(255) {
		v82 = v82 + v90
		v83 = v83 + v90
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v157 = v104
	v158 = int32(1)
	goto L7
L38:
	;
	if v135-v134&int32(255) != 0 {
		goto L5
	} else {
		goto L46
	}
L39:
	;
	goto L38
L40:
	;
	if v111 != v110&int32(255) {
		v134 = v110
		v135 = v111
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v117 = v106
	v118 = v107
	goto L42
L42:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if v122 == int32(0) {
		v134 = v121
		v135 = v122
		goto L39
	} else {
		goto L44
	}
L43:
	;
	v134 = v121
	v135 = v122
	goto L39
L44:
	;
	v125 = int32(1)
	if v122 == v121&int32(255) {
		v117 = v117 + v125
		v118 = v118 + v125
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v139 = int32(9116376)
	goto L47
L47:
	;
	v140 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_aof_main[3])) = v140
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v148 = F_strtox_2(m, v143, v7+int32(32), int32(10), int64(2147483648))
	mBase = m.M
	goto L48
L48:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_redis_check_aof_main[4])) = base.I64_extend_i32_s(base.I32_wrap_i64(v148))
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_aof_main[3]))
	if v152 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v154 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v157 = v155
	v158 = int32(0)
	goto L7
L51:
	;
	if base.Ui32(v164) < base.Ui32(int32(4097)) {
		goto L4
	} else {
		goto L55
	}
L52:
	;
	v164 = v162 - v157
	goto L54
L53:
	;
	v164 = v159
	goto L54
L54:
	;
	goto L51
L55:
	;
	v168 = F_puts(m, int32(_a_F_redis_check_aof_main_6))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L25
	} else {
		goto L56
	}
L56:
	;
	goto L5
L57:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	v237 = v235 + int32(1)
	if v237 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L59:
	;
	v235 = v227 - v157
	goto L58
L60:
	;
	v206 = v202
	goto L68
L61:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v188 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v191 = v157
	goto L64
L63:
	;
	v235 = v157 - v157
	goto L58
L64:
	;
	v195 = v191 + int32(1)
	if v195&int32(3) == int32(0) {
		v202 = v195
		goto L60
	} else {
		goto L66
	}
L66:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v200 != 0 {
		v191 = v195
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v227 = v195
	goto L59
L68:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v215 = int32(-2139062144)
	if (int32(16843008)-v212|v212)&v215 == v215 {
		v206 = v206 + int32(4)
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v221 = v206
	goto L71
L70:
	;
	goto L69
L71:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	if v225 != 0 {
		v221 = v221 + int32(1)
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v227 = v221
	goto L59
L73:
	;
	goto L72
L74:
	;
	v243 = v7 + int32(32)
	if v243 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	goto L74
L76:
	;
	v240 = F__emscripten_memcpy_bulkmem(m, v7+int32(32), v157, v237)
	mBase = m.M
	goto L75
L77:
	;
	v299 = F_fileIsManifest(m, v157)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L25
	} else {
		goto L96
	}
L78:
	;
	v298 = int32(_a_F_redis_check_aof_main_7)
	goto L77
L79:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	if v248 == int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v251 = F_strlen(m, v243)
	mBase = m.M
	v253 = v251
	goto L83
L81:
	;
	v298 = int32(_a_F_redis_check_aof_main_8)
	goto L77
L82:
	;
	v273 = v267
	goto L92
L83:
	;
	v256 = v253 + int32(-1)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v256))))
	if v258 == int32(47) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if v256 != 0 {
		v253 = v256
		goto L83
	} else {
		goto L91
	}
L86:
	;
	v262 = v256
	goto L87
L87:
	;
	if v262 == int32(0) {
		goto L78
	} else {
		goto L89
	}
L88:
	;
	goto L82
L89:
	;
	v267 = v262 + int32(-1)
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v267))))
	if v269 != int32(47) {
		v262 = v267
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	goto L81
L92:
	;
	if v273 == int32(0) {
		goto L81
	} else {
		goto L94
	}
L93:
	;
	v285 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v279+int32(1)))) = uint8(v285)
	v298 = v243
	goto L77
L94:
	;
	v278 = v273 + int32(-1)
	v279 = v243 + v278
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	if v280 == int32(47) {
		v273 = v278
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	if v299 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	v301 = F_fileIsRDB(m, v157)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L25
	} else {
		goto L99
	}
L98:
	;
	F_checkOldStyleAof(m, v157, v158, int32(1))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L25
	} else {
		goto L102
	}
L99:
	;
	if v301 != 0 {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	F_checkOldStyleAof(m, v157, v158, int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L25
	} else {
		goto L101
	}
L101:
	;
	goto L1
L102:
	;
	goto L1
L103:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	goto L1
}
func F_redis_check_rdb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v354 int32
	_ = v354
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v389 int64
	_ = v389
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v408 int64
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v422 int64
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int64
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v485 int64
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v494 int64
	_ = v494
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v515 int64
	_ = v515
	var v516 int32
	_ = v516
	var v521 int64
	_ = v521
	var v522 int32
	_ = v522
	var v527 int64
	_ = v527
	var v528 int32
	_ = v528
	var v533 int64
	_ = v533
	var v534 int32
	_ = v534
	var v539 int64
	_ = v539
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v583 int64
	_ = v583
	var v584 int32
	_ = v584
	var v585 int64
	_ = v585
	var v586 int64
	_ = v586
	var v602 int64
	_ = v602
	var v607 int64
	_ = v607
	var v608 int32
	_ = v608
	var v613 int64
	_ = v613
	var v614 int32
	_ = v614
	var v618 int64
	_ = v618
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v697 int64
	_ = v697
	var v698 int32
	_ = v698
	var v703 int64
	_ = v703
	var v704 int32
	_ = v704
	var v709 int64
	_ = v709
	var v710 int32
	_ = v710
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v747 int32
	_ = v747
	var v755 int32
	_ = v755
	var v763 int32
	_ = v763
	var v771 int32
	_ = v771
	var v779 int32
	_ = v779
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v873 int32
	_ = v873
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v892 int64
	_ = v892
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int64
	_ = v900
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v924 int32
	_ = v924
	var v936 int32
	_ = v936
	var v951 int32
	_ = v951
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1017 int64
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	v16 = m.G0
	v18 = v16 - int32(1296)
	m.G0 = v18
	v21 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_redis_check_rdb[0])) = v21
	if l1 != 0 {
		v26 = l1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(1296)
	return v1120
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+76))
	if int32(-1) < v29 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v24 = F_fopen(m, l0, int32(_a_F_redis_check_rdb_0))
	mBase = m.M
	if v24 != 0 {
		v26 = v24
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v1120 = int32(1)
	goto L1
L5:
	;
	F_startLoadingFile(m, v63, l0, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+200)) = int64(0)
	v63 = int32(0)
	goto L5
L7:
	;
	if int32(-1) < v46 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	if int32(-1) < v38 {
		v46 = v38
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v33 = F___lockfile(m, v26)
	mBase = m.M
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+60))
	if v33 == int32(0) {
		v38 = v34
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+60))
	v38 = v32
	goto L8
L11:
	;
	F___unlockfile(m, v26)
	mBase = m.M
	v38 = v34
	goto L8
L12:
	;
	goto L7
L13:
	;
	v42 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(8)
	v46 = int32(-1)
	goto L12
L14:
	;
	if v56 == int32(-1) {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	v55 = F___fstatat(m, v46, int32(_a_F_redis_check_rdb_1), v18+int32(176), int32(4096))
	mBase = m.M
	v56 = v55
	goto L14
L16:
	;
	v52 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v56 = v52
	goto L14
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v18)+200))
	v63 = v59
	goto L5
L18:
	;
	return int32(0)
L19:
	;
	v72 = F___memcpy(m, int32(_a_F_redis_check_rdb_2), int32(_a_F_redis_check_rdb_3), int32(80))
	mBase = m.M
	v73 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v72)+56)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v72)+48)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v72+int32(64)))) = v73
	v82 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v72+int32(72)))) = uint8(v82)
	goto L20
L20:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[1])) = int32(967)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[2])) = int32(_a_F_redis_check_rdb_2)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_redis_check_rdb[3])))
	if v91&int32(5) != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	F_stopLoading(m, v1107)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L18
	} else {
		goto L274
	}
L22:
	;
	v1101 = F_fclose(m, v26)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L18
	} else {
		goto L273
	}
L23:
	;
	v1084 = int32(1)
	v1085 = int32(0)
	if l1 != 0 {
		v1103 = v1084
		v1107 = v1085
		goto L21
	} else {
		goto L272
	}
L24:
	;
	v1057 = int32(0)
	v1058 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[4]))
	if v1058 == v1057 {
		goto L268
	} else {
		goto L269
	}
L25:
	;
	v97 = int32(9)
	v102 = v18 + int32(272)
	goto L26
L26:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[5]))
	if base.Ui32(v114) < base.Ui32(v97) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v142 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+281)) = uint8(v142)
	v145 = v18 + int32(272)
	v146 = int32(_a_F_redis_check_rdb_4)
	v147 = int32(6)
	goto L46
L28:
	;
	v128 = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[1]))
	if v129 == v128 {
		goto L37
	} else {
		goto L38
	}
L29:
	;
	v116 = v114
	goto L31
L30:
	;
	v116 = v97
	goto L31
L31:
	;
	if v114 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v117 = v116
	goto L34
L33:
	;
	v117 = v97
	goto L34
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[6]))
	v120 = m.T0[v119].(func(*base.Module, int32, int32, int32) int32)(m, int32(_a_F_redis_check_rdb_2), v102, v117)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	if v120 != 0 {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v122 = int32(0)
	v124 = *(*int64)(unsafe.Add(mBase, _c_F_redis_check_rdb[3]))
	*(*int64)(unsafe.Add(mBase, _c_F_redis_check_rdb[3])) = v124 | int64(1)
	goto L24
L37:
	;
	v135 = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[7])) = v137 + v117
	v141 = v97 - v117
	if v141 != 0 {
		v97 = v141
		v102 = v102 + v117
		goto L26
	} else {
		goto L40
	}
L38:
	;
	m.T0[v129].(func(*base.Module, int32, int32, int32))(m, int32(_a_F_redis_check_rdb_2), v102, v117)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L27
L41:
	;
	v289 = base.B2i32(v211 == int32(0))
	v297 = v18 + int32(272) | int32(6)
	goto L80
L42:
	;
	if v211 == int32(0) {
		goto L41
	} else {
		goto L58
	}
L43:
	;
	v211 = int32(0)
	goto L42
L44:
	;
	v183 = v178
	v184 = v179
	v185 = v180
	goto L54
L45:
	;
	if v168 == int32(0) {
		goto L43
	} else {
		goto L52
	}
L46:
	;
	if (v146|v145)&int32(3) != 0 {
		v178 = v145
		v179 = v146
		v180 = v147
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v155 = v145
	v156 = v146
	v157 = v147
	goto L48
L48:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v160 != v161 {
		v178 = v155
		v179 = v156
		v180 = v157
		goto L44
	} else {
		goto L50
	}
L49:
	;
	goto L45
L50:
	;
	v163 = int32(4)
	v164 = v156 + v163
	v166 = v155 + v163
	v168 = v157 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v168) {
		v155 = v166
		v156 = v164
		v157 = v168
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v178 = v166
	v179 = v164
	v180 = v168
	goto L44
L53:
	;
	v211 = v188 - v189
	goto L42
L54:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v188 != v189 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v191 = int32(1)
	v196 = v185 + int32(-1)
	if v196 == int32(0) {
		goto L43
	} else {
		goto L57
	}
L57:
	;
	v183 = v183 + v191
	v184 = v184 + v191
	v185 = v196
	goto L54
L58:
	;
	v215 = v18 + int32(272)
	v216 = int32(_a_F_redis_check_rdb_5)
	v217 = int32(6)
	goto L63
L59:
	;
	if v281 == int32(0) {
		goto L41
	} else {
		goto L75
	}
L60:
	;
	v281 = int32(0)
	goto L59
L61:
	;
	v253 = v248
	v254 = v249
	v255 = v250
	goto L71
L62:
	;
	if v238 == int32(0) {
		goto L60
	} else {
		goto L69
	}
L63:
	;
	if (v216|v215)&int32(3) != 0 {
		v248 = v215
		v249 = v216
		v250 = v217
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v225 = v215
	v226 = v216
	v227 = v217
	goto L65
L65:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	if v230 != v231 {
		v248 = v225
		v249 = v226
		v250 = v227
		goto L61
	} else {
		goto L67
	}
L66:
	;
	goto L62
L67:
	;
	v233 = int32(4)
	v234 = v226 + v233
	v236 = v225 + v233
	v238 = v227 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v238) {
		v225 = v236
		v226 = v234
		v227 = v238
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v248 = v236
	v249 = v234
	v250 = v238
	goto L61
L70:
	;
	v281 = v258 - v259
	goto L59
L71:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if v258 != v259 {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v261 = int32(1)
	v266 = v255 + int32(-1)
	if v266 == int32(0) {
		goto L60
	} else {
		goto L74
	}
L74:
	;
	v253 = v253 + v261
	v254 = v254 + v261
	v255 = v266
	goto L71
L75:
	;
	F_rdbCheckError(m, int32(_a_F_redis_check_rdb_6), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L18
	} else {
		goto L76
	}
L76:
	;
	goto L23
L77:
	;
	if base.Ui32(v342) <= base.Ui32(int32(80)) {
		goto L100
	} else {
		goto L101
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v342
	F_rdbCheckError(m, int32(_a_F_redis_check_rdb_7), v18)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L18
	} else {
		goto L97
	}
L79:
	;
	if v289&base.B2i32(base.Ui32(int32(79)) < base.Ui32(v342)) != 0 {
		goto L78
	} else {
		goto L94
	}
L80:
	;
	v302 = v297 + int32(1)
	v303 = int32(*(*int8)(unsafe.Add(mBase, uint32(v297))))
	v304 = F___isspace_1(m, v303)
	mBase = m.M
	if v304 != 0 {
		v297 = v302
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v305 = int32(1)
	switch v303&int32(255) + int32(-43) {
	case 0:
		v311 = v305
		goto L84
	default:
		v313 = v297
		v314 = v303
		v315 = v305
		goto L83
	case 2:
		goto L85
	}
L82:
	;
	goto L81
L83:
	;
	v318 = v314 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v318) {
		v336 = int32(0)
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v312 = int32(*(*int8)(unsafe.Add(mBase, uint32(v302))))
	v313 = v302
	v314 = v312
	v315 = v311
	goto L83
L85:
	;
	v311 = int32(0)
	goto L84
L86:
	;
	if v315 != 0 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v322 = int32(0)
	v323 = v313
	v324 = v318
	goto L88
L88:
	;
	v326 = int32(10)
	v328 = v322*v326 - v324
	v329 = int32(*(*int8)(unsafe.Add(mBase, uint32(v323)+1)))
	v333 = v329 + int32(-48)
	if base.Ui32(v333) < base.Ui32(v326) {
		v322 = v328
		v323 = v323 + int32(1)
		v324 = v333
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v336 = v328
	goto L86
L90:
	;
	goto L89
L91:
	;
	v342 = int32(0) - v336
	goto L93
L92:
	;
	v342 = v336
	goto L93
L93:
	;
	goto L79
L94:
	;
	if v342 < int32(1) {
		goto L78
	} else {
		goto L95
	}
L95:
	;
	if v289|base.B2i32(base.Ui32(int32(11)) < base.Ui32(v342)) != 0 {
		goto L77
	} else {
		goto L96
	}
L96:
	;
	goto L78
L97:
	;
	goto L23
L98:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[8])) = v342
	v389 = int64(-1)
	v390 = int32(-1)
	goto L104
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+144)) = v342
	F_rdbCheckInfo(m, v365, v18+int32(144))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L18
	} else {
		goto L103
	}
L100:
	;
	if base.Ui32(int32(67)) < base.Ui32(v342+int32(-12)) {
		v373 = int32(0)
		goto L98
	} else {
		goto L102
	}
L101:
	;
	v365 = int32(_a_F_redis_check_rdb_8)
	goto L99
L102:
	;
	v365 = int32(_a_F_redis_check_rdb_9)
	goto L99
L103:
	;
	v373 = base.B2i32(base.Ui32(v342) < base.Ui32(int32(81)))
	goto L98
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[9])) = int32(1)
	v400 = F_rdbLoadType(m, int32(_a_F_redis_check_rdb_2))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L18
	} else {
		goto L124
	}
L106:
	;
	if v373^int32(1)|base.B2i32(base.Ui32(int32(221)) < base.Ui32(v400+int32(-22))) != 0 {
		goto L240
	} else {
		goto L241
	}
L107:
	;
	F_rdbCheckError(m, int32(_a_F_redis_check_rdb_10), int32(0))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L18
	} else {
		goto L239
	}
L108:
	;
	if v400 != int32(-1) {
		goto L106
	} else {
		goto L238
	}
L109:
	;
	if v342 < int32(5) {
		goto L227
	} else {
		goto L228
	}
L110:
	;
	if v873 == int32(4) {
		goto L104
	} else {
		goto L226
	}
L111:
	;
	v873 = int32(4)
	goto L110
L112:
	;
	v816 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[9])) = int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+160)) = v816
	v826 = F_rdbFunctionLoad(m, int32(_a_F_redis_check_rdb_2), v342, v816, v816, v18+int32(160))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L18
	} else {
		goto L222
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v342
	F_rdbCheckError(m, int32(_a_F_redis_check_rdb_11), v18+int32(112))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L18
	} else {
		goto L220
	}
L114:
	;
	v692 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[9])) = int32(8)
	v697 = F_rdbLoadLen(m, int32(_a_F_redis_check_rdb_2), v692)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L18
	} else {
		goto L207
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[9])) = int32(7)
	v624 = F_rdbLoadStringObject(m, int32(_a_F_redis_check_rdb_2))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L18
	} else {
		goto L184
	}
L116:
	;
	v544 = F_rdbLoadStringObject(m, int32(_a_F_redis_check_rdb_2))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L18
	} else {
		goto L161
	}
L117:
	;
	v527 = F_rdbLoadLen(m, int32(_a_F_redis_check_rdb_2), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L18
	} else {
		goto L155
	}
L118:
	;
	v510 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[9])) = int32(6)
	v515 = F_rdbLoadLen(m, int32(_a_F_redis_check_rdb_2), v510)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L18
	} else {
		goto L151
	}
L119:
	;
	v489 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[9])) = int32(6)
	v494 = F_rdbLoadLen(m, int32(_a_F_redis_check_rdb_2), v489)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L18
	} else {
		goto L147
	}
L120:
	;
	v485 = F_rdbLoadLen(m, int32(_a_F_redis_check_rdb_2), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L18
	} else {
		goto L145
	}
L121:
	;
	v430 = int32(2)
	v432 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_redis_check_rdb[3])))
	if v432&int32(5) != 0 {
		v873 = v430
		goto L110
	} else {
		goto L129
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[9])) = int32(2)
	v422 = F_rdbLoadMillisecondTime(m, int32(_a_F_redis_check_rdb_2), v342)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L18
	} else {
		goto L127
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[9])) = int32(2)
	v408 = F_rdbLoadTime(m, int32(_a_F_redis_check_rdb_2))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L18
	} else {
		goto L125
	}
L124:
	;
	switch v400 + int32(-243) {
	case 0:
		goto L116
	case 1:
		goto L117
	case 2:
		goto L112
	case 3:
		goto L113
	case 4:
		goto L114
	case 5:
		goto L120
	case 6:
		goto L121
	case 7:
		goto L115
	case 8:
		goto L118
	case 9:
		goto L122
	case 10:
		goto L123
	case 11:
		goto L119
	case 12:
		goto L109
	default:
		goto L108
	}
L125:
	;
	v412 = int32(0)
	v413 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_redis_check_rdb[3])))
	if v413&int32(1) == v412 {
		v389 = v408 * int64(1000)
		goto L104
	} else {
		goto L126
	}
L126:
	;
	goto L24
L127:
	;
	v424 = int32(0)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_redis_check_rdb[3])))
	if v425&int32(1) == v424 {
		v389 = v422
		goto L104
	} else {
		goto L128
	}
L128:
	;
	goto L24
L129:
	;
	v438 = int32(1)
	v443 = v18 + int32(160)
	goto L130
L130:
	;
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[5]))
	if base.Ui32(v455) < base.Ui32(v438) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v469 = int32(0)
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[1]))
	if v470 == v469 {
		goto L141
	} else {
		goto L142
	}
L133:
	;
	v457 = v455
	goto L135
L134:
	;
	v457 = v438
	goto L135
L135:
	;
	if v455 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v458 = v457
	goto L138
L137:
	;
	v458 = v438
	goto L138
L138:
	;
	v460 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[6]))
	v461 = m.T0[v460].(func(*base.Module, int32, int32, int32) int32)(m, int32(_a_F_redis_check_rdb_2), v443, v458)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L18
	} else {
		goto L139
	}
L139:
	;
	if v461 != 0 {
		goto L132
	} else {
		goto L140
	}
L140:
	;
	v463 = int32(0)
	v465 = *(*int64)(unsafe.Add(mBase, _c_F_redis_check_rdb[3]))
	*(*int64)(unsafe.Add(mBase, _c_F_redis_check_rdb[3])) = v465 | int64(1)
	v873 = v430
	goto L110
L141:
	;
	v476 = int32(0)
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[7])) = v478 + v458
	v482 = v438 - v458
	if v482 != 0 {
		v438 = v482
		v443 = v443 + v458
		goto L130
	} else {
		goto L144
	}
L142:
	;
	m.T0[v470].(func(*base.Module, int32, int32, int32))(m, int32(_a_F_redis_check_rdb_2), v443, v458)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L18
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	goto L111
L145:
	;
	if v485 != int64(-1) {
		goto L104
	} else {
		goto L146
	}
L146:
	;
	goto L24
L147:
	;
	if v494 == int64(-1) {
		goto L24
	} else {
		goto L148
	}
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v494
	F_rdbCheckInfo(m, int32(_a_F_redis_check_rdb_12), v18+int32(64))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[10]))
	v506 = base.I32_wrap_i64(v494)
	if v506 <= v505 {
		v390 = v506
		goto L104
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[10])) = v506
	v390 = v506
	goto L104
L151:
	;
	if v515 == int64(-1) {
		goto L24
	} else {
		goto L152
	}
L152:
	;
	v521 = F_rdbLoadLen(m, int32(_a_F_redis_check_rdb_2), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L18
	} else {
		goto L153
	}
L153:
	;
	if v521 != int64(-1) {
		goto L104
	} else {
		goto L154
	}
L154:
	;
	goto L24
L155:
	;
	if v527 == int64(-1) {
		goto L24
	} else {
		goto L156
	}
L156:
	;
	v533 = F_rdbLoadLen(m, int32(_a_F_redis_check_rdb_2), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L18
	} else {
		goto L157
	}
L157:
	;
	if v533 == int64(-1) {
		goto L24
	} else {
		goto L158
	}
L158:
	;
	v539 = F_rdbLoadLen(m, int32(_a_F_redis_check_rdb_2), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L18
	} else {
		goto L159
	}
L159:
	;
	if v539 != int64(-1) {
		goto L104
	} else {
		goto L160
	}
L160:
	;
	goto L24
L161:
	;
	if v544 == int32(0) {
		goto L24
	} else {
		goto L162
	}
L162:
	;
	v548 = F_objectGetVal(m, v544)
	mBase = m.M
	v549 = int32(-1)
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548+v549))))
	switch v551&int32(7) + v549 {
	case 0:
		goto L169
	case 1:
		goto L168
	case 2:
		goto L167
	case 3:
		goto L166
	default:
		goto L164
	}
L163:
	;
	F_decrRefCount(m, v544)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L18
	} else {
		goto L173
	}
L164:
	;
	F_rdbCheckError(m, int32(_a_F_redis_check_rdb_13), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L18
	} else {
		goto L171
	}
L165:
	;
	if v568 == int32(40) {
		goto L163
	} else {
		goto L170
	}
L166:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v548+int32(-17))))
	v568 = v567
	goto L165
L167:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v548+int32(-9))))
	v568 = v564
	goto L165
L168:
	;
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v548+int32(-5)))))
	v568 = v561
	goto L165
L169:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548+int32(-3)))))
	v568 = v558
	goto L165
L170:
	;
	goto L164
L171:
	;
	F_decrRefCount(m, v544)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L18
	} else {
		goto L172
	}
L172:
	;
	goto L23
L173:
	;
	v583 = F_rdbLoadLen(m, int32(_a_F_redis_check_rdb_2), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L18
	} else {
		goto L175
	}
L174:
	;
	v602 = int64(0)
	goto L177
L175:
	;
	v585 = int64(1)
	v586 = v583 + v585
	if base.Ui64(v585) < base.Ui64(v586) {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	switch base.I32_wrap_i64(v586) {
	default:
		goto L24
	case 1:
		goto L104
	}
L177:
	;
	v607 = F_rdbLoadLen(m, int32(_a_F_redis_check_rdb_2), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L18
	} else {
		goto L179
	}
L179:
	;
	if v607 == int64(-1) {
		goto L24
	} else {
		goto L180
	}
L180:
	;
	v613 = F_rdbLoadLen(m, int32(_a_F_redis_check_rdb_2), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L18
	} else {
		goto L181
	}
L181:
	;
	if v613 == int64(-1) {
		goto L24
	} else {
		goto L182
	}
L182:
	;
	v618 = v602 + int64(1)
	if v618 == v583 {
		goto L104
	} else {
		goto L183
	}
L183:
	;
	v602 = v618
	goto L177
L184:
	;
	if v624 == int32(0) {
		goto L24
	} else {
		goto L185
	}
L185:
	;
	v629 = F_rdbLoadStringObject(m, int32(_a_F_redis_check_rdb_2))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L18
	} else {
		goto L187
	}
L186:
	;
	v633 = F_objectGetVal(m, v624)
	mBase = m.M
	v634 = int32(_a_F_redis_check_rdb_14)
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633))))
	if v637 != 0 {
		goto L193
	} else {
		goto L194
	}
L187:
	;
	if v629 != 0 {
		goto L186
	} else {
		goto L188
	}
L188:
	;
	F_decrRefCount(m, v624)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L18
	} else {
		goto L189
	}
L189:
	;
	goto L24
L190:
	;
	v679 = F_objectGetVal(m, v624)
	mBase = m.M
	v680 = F_objectGetVal(m, v629)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v680
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v679
	F_rdbCheckInfo(m, int32(_a_F_redis_check_rdb_15), v18+int32(80))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L18
	} else {
		goto L204
	}
L191:
	;
	if v669-v671 != 0 {
		goto L190
	} else {
		goto L203
	}
L192:
	;
	v669 = F_tolower(m, v665)
	mBase = m.M
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	v671 = F_tolower(m, v670)
	mBase = m.M
	goto L191
L193:
	;
	v639 = v633
	v640 = v634
	v641 = v637
	goto L196
L194:
	;
	v665 = int32(0)
	v666 = v634
	goto L192
L195:
	;
	v665 = v662 & int32(255)
	v666 = v661
	goto L192
L196:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640))))
	if v643 == int32(0) {
		v661 = v640
		v662 = v641
		goto L195
	} else {
		goto L198
	}
L197:
	;
	v661 = v655
	v662 = int32(0)
	goto L195
L198:
	;
	v647 = v641 & int32(255)
	if v647 == v643 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v654 = int32(1)
	v655 = v640 + v654
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639)+1)))
	if v656 != 0 {
		v639 = v639 + v654
		v640 = v655
		v641 = v656
		goto L196
	} else {
		goto L202
	}
L200:
	;
	v649 = F_tolower(m, v647)
	mBase = m.M
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640))))
	v651 = F_tolower(m, v650)
	mBase = m.M
	if v649 == v651 {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639))))
	v661 = v640
	v662 = v653
	goto L195
L202:
	;
	goto L197
L203:
	;
	v673 = int32(0)
	v675 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[11])) = v675 + int32(1)
	goto L190
L204:
	;
	F_decrRefCount(m, v624)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L18
	} else {
		goto L205
	}
L205:
	;
	F_decrRefCount(m, v629)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L18
	} else {
		goto L206
	}
L206:
	;
	goto L104
L207:
	;
	if v697 == int64(-1) {
		goto L24
	} else {
		goto L208
	}
L208:
	;
	v703 = F_rdbLoadLen(m, int32(_a_F_redis_check_rdb_2), int32(0))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L18
	} else {
		goto L209
	}
L209:
	;
	if v703 == int64(-1) {
		goto L24
	} else {
		goto L210
	}
L210:
	;
	v709 = F_rdbLoadLen(m, int32(_a_F_redis_check_rdb_2), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L18
	} else {
		goto L211
	}
L211:
	;
	if v709 == int64(-1) {
		goto L24
	} else {
		goto L212
	}
L212:
	;
	if v703 == int64(2) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v720 = v18 + int32(160)
	v721 = int32(0)
	v724 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[12]))
	*(*uint8)(unsafe.Add(mBase, uint32(v720)+9)) = uint8(v721)
	v727 = base.I32_wrap_i64(v697)
	v730 = int32(63)
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724+int32(base.Ui32(v727)>>(uint(int32(10))%32))&v730))))
	*(*uint8)(unsafe.Add(mBase, uint32(v720)+8)) = uint8(v733)
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724+int32(base.Ui32(v727)>>(uint(int32(16))%32))&v730))))
	*(*uint8)(unsafe.Add(mBase, uint32(v720)+7)) = uint8(v740)
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724+int32(base.Ui32(v727)>>(uint(int32(22))%32))&v730))))
	*(*uint8)(unsafe.Add(mBase, uint32(v720)+6)) = uint8(v747)
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724+base.I32_wrap_i64(int64(base.Ui64(v697)>>(uint(int64(28))%64)))&v730))))
	*(*uint8)(unsafe.Add(mBase, uint32(v720)+5)) = uint8(v755)
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724+base.I32_wrap_i64(int64(base.Ui64(v697)>>(uint(int64(34))%64)))&v730))))
	*(*uint8)(unsafe.Add(mBase, uint32(v720)+4)) = uint8(v763)
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724+base.I32_wrap_i64(int64(base.Ui64(v697)>>(uint(int64(40))%64)))&v730))))
	*(*uint8)(unsafe.Add(mBase, uint32(v720)+3)) = uint8(v771)
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724+base.I32_wrap_i64(int64(base.Ui64(v697)>>(uint(int64(46))%64)))&v730))))
	*(*uint8)(unsafe.Add(mBase, uint32(v720)+2)) = uint8(v779)
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724+base.I32_wrap_i64(int64(base.Ui64(v697)>>(uint(int64(52))%64)))&v730))))
	*(*uint8)(unsafe.Add(mBase, uint32(v720)+1)) = uint8(v787)
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724+base.I32_wrap_i64(int64(base.Ui64(v697)>>(uint(int64(58))%64)))))))
	*(*uint8)(unsafe.Add(mBase, uint32(v720))) = uint8(v793)
	goto L216
L214:
	;
	F_rdbCheckError(m, int32(_a_F_redis_check_rdb_16), int32(0))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L18
	} else {
		goto L215
	}
L215:
	;
	goto L23
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v18 + int32(160)
	F_rdbCheckInfo(m, int32(_a_F_redis_check_rdb_17), v18+int32(96))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L18
	} else {
		goto L217
	}
L217:
	;
	v806 = F_rdbLoadCheckModuleValue(m, int32(_a_F_redis_check_rdb_2), v18+int32(160))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L18
	} else {
		goto L218
	}
L218:
	;
	F_decrRefCount(m, v806)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L18
	} else {
		goto L219
	}
L219:
	;
	goto L104
L220:
	;
	goto L23
L221:
	;
	v841 = int32(0)
	v843 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[13])) = v843 + int32(1)
	goto L111
L222:
	;
	if v826 == int32(0) {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v18)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v830
	F_rdbCheckError(m, int32(_a_F_redis_check_rdb_18), v18+int32(128))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L18
	} else {
		goto L224
	}
L224:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v18)+160))
	F_sdsfree(m, v837)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L18
	} else {
		goto L225
	}
L225:
	;
	v873 = int32(3)
	goto L110
L226:
	;
	switch v873 + int32(-2) {
	default:
		goto L24
	case 1:
		goto L23
	}
L227:
	;
	v915 = int32(0)
	v916 = int32(1)
	if l1 == v915 {
		v1086 = v915
		v1090 = v916
		goto L22
	} else {
		goto L237
	}
L228:
	;
	v885 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[14]))
	if v885 == int32(0) {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v888 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[9])) = int32(5)
	v892 = *(*int64)(unsafe.Add(mBase, _c_F_redis_check_rdb[15]))
	v896 = F_rioRead_2(m, v18+int32(160), int32(8))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L18
	} else {
		goto L230
	}
L230:
	;
	if v896 == int32(0) {
		goto L24
	} else {
		goto L231
	}
L231:
	;
	v900 = *(*int64)(unsafe.Add(mBase, uint32(v18)+160))
	if base.B2i32(v900 == int64(0)) == int32(0) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	F_rdbCheckInfo(m, v908, int32(0))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L18
	} else {
		goto L236
	}
L233:
	;
	if v900 != v892 {
		goto L107
	} else {
		goto L235
	}
L234:
	;
	v908 = int32(_a_F_redis_check_rdb_19)
	goto L232
L235:
	;
	v908 = int32(_a_F_redis_check_rdb_20)
	goto L232
L236:
	;
	goto L227
L237:
	;
	v1103 = v915
	v1107 = v916
	goto L21
L238:
	;
	goto L24
L239:
	;
	goto L23
L240:
	;
	if base.Ui32(v400) < base.Ui32(int32(8)) {
		goto L243
	} else {
		goto L244
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v400
	F_rdbCheckError(m, int32(_a_F_redis_check_rdb_21), v18+int32(16))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L18
	} else {
		goto L242
	}
L242:
	;
	goto L23
L243:
	;
	v958 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[9])) = int32(3)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[16])) = v400
	v964 = F_rdbLoadStringObject(m, int32(_a_F_redis_check_rdb_2))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L18
	} else {
		goto L250
	}
L244:
	;
	if base.Ui32(v400+int32(-9)) < base.Ui32(int32(14)) {
		goto L243
	} else {
		goto L245
	}
L245:
	;
	if base.Ui32(v342) < base.Ui32(int32(81)) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v400
	F_rdbCheckError(m, int32(_a_F_redis_check_rdb_22), v18+int32(32))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L18
	} else {
		goto L249
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v400
	F_rdbCheckError(m, int32(_a_F_redis_check_rdb_23), v18+int32(48))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L18
	} else {
		goto L248
	}
L248:
	;
	goto L23
L249:
	;
	goto L23
L250:
	;
	if v964 == int32(0) {
		goto L24
	} else {
		goto L251
	}
L251:
	;
	v968 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[17])) = v964
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[9])) = int32(4)
	v975 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[18]))
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[18])) = v975 + int32(1)
	v980 = F_objectGetVal(m, v964)
	mBase = m.M
	v984 = F_rdbLoadObject(m, v400, int32(_a_F_redis_check_rdb_2), v980, v390, v968, v968, int64(0))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L18
	} else {
		goto L252
	}
L252:
	;
	if v984 == int32(0) {
		goto L24
	} else {
		goto L253
	}
L253:
	;
	v988 = int32(0)
	v989 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[19]))
	if v989 == v988 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1015 = base.B2i32(v389 == int64(-1))
	if v389 == int64(-1) {
		goto L262
	} else {
		goto L263
	}
L255:
	;
	v992 = int32(0)
	v993 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[10]))
	v994 = int32(7)
	v997 = v993*v994 + v994
	v999 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[20]))
	if v997 <= v999 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	F_computeDatasetProfile(m, v390, v964, v984, v389)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L18
	} else {
		goto L259
	}
L257:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[21]))
	v1003 = F_tryExpandRdbStats(m, v1002, v999, v997)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L18
	} else {
		goto L258
	}
L258:
	;
	v1005 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[20])) = v997
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[21])) = v1003
	goto L256
L259:
	;
	goto L254
L260:
	;
	v1031 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[17])) = v1031
	F_decrRefCount(m, v964)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L18
	} else {
		goto L266
	}
L261:
	;
	v1025 = int32(0)
	v1027 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[22])) = v1027 + int32(1)
	goto L260
L262:
	;
	if v389 == int64(-1) {
		goto L260
	} else {
		goto L265
	}
L263:
	;
	v1017 = *(*int64)(unsafe.Add(mBase, _c_F_redis_check_rdb[0]))
	if v1017 <= v389 {
		goto L262
	} else {
		goto L264
	}
L264:
	;
	v1019 = int32(0)
	v1021 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[23])) = v1021 + int32(1)
	goto L261
L265:
	;
	goto L261
L266:
	;
	F_decrRefCount(m, v984)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L18
	} else {
		goto L267
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb[16])) = int32(-1)
	v389 = int64(-1)
	goto L104
L268:
	;
	F_rdbCheckError(m, int32(_a_F_redis_check_rdb_24), int32(0))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L18
	} else {
		goto L271
	}
L269:
	;
	F_rdbCheckError(m, int32(_a_F_redis_check_rdb_25), int32(0))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L18
	} else {
		goto L270
	}
L270:
	;
	goto L23
L271:
	;
	goto L23
L272:
	;
	v1086 = v1084
	v1090 = v1085
	goto L22
L273:
	;
	v1103 = v1086
	v1107 = v1090
	goto L21
L274:
	;
	v1120 = v1103
	goto L1
}
func F_redis_check_rdb_main(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v58 int64
	_ = v58
	var v61 int64
	_ = v61
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v88 int64
	_ = v88
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v296 int32
	_ = v296
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	v6 = m.G0
	v8 = v6 - int32(192)
	m.G0 = v8
	F_parseCheckRdbOptions(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = F___gettimeofday(m, v8+int32(32), int32(0))
	mBase = m.M
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
	v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+40)))
	v23 = F___syscall_getpid(m)
	mBase = m.M
	goto L3
L3:
	;
	v25 = v18*int64(1000000) + v21 ^ base.I64_extend_i32_s(v23)
	*(*int64)(unsafe.Add(mBase, _c_F_redis_check_rdb_main[0])) = v25
	v31 = v25
	v32 = int64(1)
	goto L6
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb_main[2])) = int32(-1)
	v99 = F_valkey_malloc(m, int32(28))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L9
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb_main[1])) = int32(312)
	goto L4
L6:
	;
	v35 = int32(3)
	v39 = int64(62)
	v42 = int64(6364136223846793005)
	v44 = (int64(base.Ui64(v31)>>(uint(v39)%64))^v31)*v42 + v32
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v32)<<(uint(v35)%32))+uint32(_c_F_redis_check_rdb_main[0]))) = v44
	v47 = v32 + int64(1)
	v58 = (int64(base.Ui64(v44)>>(uint(v39)%64))^v44)*v42 + v47
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v47)<<(uint(v35)%32))+uint32(_c_F_redis_check_rdb_main[0]))) = v58
	v61 = v32 + int64(2)
	v72 = (int64(base.Ui64(v58)>>(uint(v39)%64))^v58)*v42 + v61
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v61)<<(uint(v35)%32))+uint32(_c_F_redis_check_rdb_main[0]))) = v72
	v75 = v32 + int64(3)
	if v75 == int64(312) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v88 = (int64(base.Ui64(v72)>>(uint(int64(62))%64))^v72)*int64(6364136223846793005) + v75
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v75)<<(uint(int32(3))%32))+uint32(_c_F_redis_check_rdb_main[0]))) = v88
	v31 = v88
	v32 = v32 + int64(4)
	goto L6
L9:
	;
	v102 = F_valkey_calloc(m, int32(48))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v102
	v124 = F_valkey_calloc(m, int32(48))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L16
	}
L11:
	;
	if v102 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = int32(0)
	v113 = F_hdr_init(m, int64(1), int64(204800), int32(3), v102+int32(40))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v120 = F_hdr_init(m, int64(1), int64(1048576), int32(3), v102+int32(44))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v124
	v146 = F_valkey_calloc(m, int32(48))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L21
	}
L16:
	;
	if v124 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = int32(1)
	v135 = F_hdr_init(m, int64(1), int64(204800), int32(3), v124+int32(40))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v142 = F_hdr_init(m, int64(1), int64(1048576), int32(3), v124+int32(44))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v146
	v168 = F_valkey_calloc(m, int32(48))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L26
	}
L21:
	;
	if v146 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = int32(2)
	v157 = F_hdr_init(m, int64(1), int64(204800), int32(3), v146+int32(40))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v164 = F_hdr_init(m, int64(1), int64(1048576), int32(3), v146+int32(44))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v168
	v190 = F_valkey_calloc(m, int32(48))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L31
	}
L26:
	;
	if v168 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v172 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v172
	v179 = F_hdr_init(m, int64(1), int64(204800), v172, v168+int32(40))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v186 = F_hdr_init(m, int64(1), int64(1048576), int32(3), v168+int32(44))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+16)) = v190
	v212 = F_valkey_calloc(m, int32(48))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L36
	}
L31:
	;
	if v190 == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = int32(4)
	v201 = F_hdr_init(m, int64(1), int64(204800), int32(3), v190+int32(40))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v208 = F_hdr_init(m, int64(1), int64(1048576), int32(3), v190+int32(44))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+20)) = v212
	v234 = F_valkey_calloc(m, int32(48))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L41
	}
L36:
	;
	if v212 == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = int32(5)
	v223 = F_hdr_init(m, int64(1), int64(204800), int32(3), v212+int32(40))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v230 = F_hdr_init(m, int64(1), int64(1048576), int32(3), v212+int32(44))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+24)) = v234
	v255 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb_main[3])) = int32(7)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb_main[4])) = v99
	*(*int64)(unsafe.Add(mBase, _c_F_redis_check_rdb_main[5])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb_main[6])) = v255
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb_main[7]))
	if v267 != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	if v234 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(6)
	v245 = F_hdr_init(m, int64(1), int64(204800), int32(3), v234+int32(40))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v252 = F_hdr_init(m, int64(1), int64(1048576), int32(3), v234+int32(44))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb_main[8])) = int32(1)
	*(*int64)(unsafe.Add(mBase, _c_F_redis_check_rdb_main[9])) = int64(0)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v276
	F_rdbCheckInfo(m, int32(_a_F_redis_check_rdb_main_0), v8+int32(16))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	F_createSharedObjects(m)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(56)))) = int64(0)
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = int32(1127)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = int32(-1073741820)
	v296 = v8 + int32(52)
	goto L51
L50:
	;
	v324 = v8 + int32(52)
	goto L58
L51:
	;
	goto L53
L53:
	;
	if v296 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	goto L50
L56:
	;
	v319 = F___memcpy(m, int32(9118724), v296, int32(140))
	mBase = m.M
	goto L55
L57:
	;
	v352 = v8 + int32(52)
	goto L65
L58:
	;
	goto L60
L60:
	;
	if v324 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	goto L57
L63:
	;
	v347 = F___memcpy(m, int32(9118164), v324, int32(140))
	mBase = m.M
	goto L62
L64:
	;
	v380 = v8 + int32(52)
	goto L72
L65:
	;
	goto L67
L67:
	;
	if v352 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	goto L64
L70:
	;
	v375 = F___memcpy(m, int32(9118304), v352, int32(140))
	mBase = m.M
	goto L69
L71:
	;
	v408 = v8 + int32(52)
	goto L79
L72:
	;
	goto L74
L74:
	;
	if v380 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	goto L71
L77:
	;
	v403 = F___memcpy(m, int32(9117744), v380, int32(140))
	mBase = m.M
	goto L76
L78:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v435 = F_redis_check_rdb(m, v434, l2)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L85
	}
L79:
	;
	goto L81
L81:
	;
	if v408 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	goto L78
L84:
	;
	v431 = F___memcpy(m, int32(9118024), v408, int32(140))
	mBase = m.M
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v435
	F_rdbCheckInfo(m, int32(_a_F_redis_check_rdb_main_1), v8)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if v435 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	if l2 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L88:
	;
	v441 = int32(_a_F_redis_check_rdb_main_2)
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb_main[10]))
	if base.Ui32(v445+int32(-12)) < base.Ui32(int32(68)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v450 = v441
	goto L91
L90:
	;
	v450 = int32(_a_F_redis_check_rdb_main_3)
	goto L91
L91:
	;
	if int32(80) < v445 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v453 = v441
	goto L94
L93:
	;
	v453 = v450
	goto L94
L94:
	;
	F_rdbCheckInfo(m, v453, int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_rdbShowGenericInfo(m)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L87
L97:
	;
	v469 = int32(0)
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb_main[4]))
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_redis_check_rdb_main[3]))
	F_freeRdbProfile(m, v470, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L102
	}
L98:
	;
	m.G0 = v8 + int32(192)
	if v435 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v467 = int32(-1)
	goto L101
L100:
	;
	v467 = int32(0)
	goto L101
L101:
	;
	return v467
L102:
	;
	m.Env.Exit(m, v435)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_rehashStep(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
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
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int64
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v354 int64
	_ = v354
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v359 int64
	_ = v359
	var v361 int64
	_ = v361
	var v363 int64
	_ = v363
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int64
	_ = v373
	var v374 int64
	_ = v374
	var v375 int64
	_ = v375
	var v376 int64
	_ = v376
	var v380 int64
	_ = v380
	var v381 int64
	_ = v381
	var v382 int64
	_ = v382
	var v383 int64
	_ = v383
	var v386 int64
	_ = v386
	var v387 int64
	_ = v387
	var v390 int64
	_ = v390
	var v393 int64
	_ = v393
	var v396 int64
	_ = v396
	var v398 int64
	_ = v398
	var v399 int64
	_ = v399
	var v401 int64
	_ = v401
	var v402 int64
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int64
	_ = v409
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v412 int64
	_ = v412
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v422 int64
	_ = v422
	var v423 int64
	_ = v423
	var v427 int64
	_ = v427
	var v428 int64
	_ = v428
	var v432 int64
	_ = v432
	var v433 int64
	_ = v433
	var v437 int64
	_ = v437
	var v438 int64
	_ = v438
	var v442 int64
	_ = v442
	var v443 int64
	_ = v443
	var v447 int64
	_ = v447
	var v448 int64
	_ = v448
	var v450 int64
	_ = v450
	var v451 int64
	_ = v451
	var v452 int64
	_ = v452
	var v454 int64
	_ = v454
	var v455 int64
	_ = v455
	var v456 int64
	_ = v456
	var v458 int64
	_ = v458
	var v459 int64
	_ = v459
	var v461 int64
	_ = v461
	var v462 int64
	_ = v462
	var v465 int64
	_ = v465
	var v467 int64
	_ = v467
	var v468 int64
	_ = v468
	var v473 int64
	_ = v473
	var v474 int64
	_ = v474
	var v478 int64
	_ = v478
	var v480 int64
	_ = v480
	var v481 int64
	_ = v481
	var v484 int64
	_ = v484
	var v485 int64
	_ = v485
	var v490 int64
	_ = v490
	var v491 int64
	_ = v491
	var v494 int64
	_ = v494
	var v500 int64
	_ = v500
	var v504 int64
	_ = v504
	var v508 int64
	_ = v508
	var v513 int64
	_ = v513
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v534 int64
	_ = v534
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v568 int32
	_ = v568
	var v577 int32
	_ = v577
	var v592 int32
	_ = v592
	v13 = m.G0
	v15 = v13 - int32(400)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(400)
	return
L2:
	;
	F_rehashStepFinalize(m, l0)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L11
	} else {
		goto L80
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_rehashStep_0), int32(_a_F_rehashStep_1), int32(716))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L11
	} else {
		goto L79
	}
L4:
	;
	v20 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+25)))
	v21 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v21 <= v20 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v148 == int32(0) {
		goto L2
	} else {
		goto L26
	}
L6:
	;
	v26 = v17
	v27 = int32(10)
	goto L8
L7:
	;
	v54 = v39
	goto L15
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = v36 + v26<<(uint(int32(6))%32)
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39))))
	if v40&int32(8191) != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_rehashStepFinalize(m, l0)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v45 == int32(-1) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v49 = v27 + int32(-1)
	if v49 != 0 {
		v26 = v45
		v27 = v49
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L1
L15:
	;
	v62 = int32(0)
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54))))
	if v63&int32(1) == v62 {
		v69 = v62
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v77 = int32(0)
	v78 = v63
	goto L19
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v54)+60))
	v69 = v68
	goto L17
L19:
	;
	v87 = int32(1)
	if int32(base.Ui32(int32(base.Ui32(v78)>>(uint(v87)%32))&int32(4095))>>(uint(v77)%32))&v87 == int32(0) {
		v137 = v78
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v69 != 0 {
		v54 = v69
		goto L15
	} else {
		goto L25
	}
L21:
	;
	v141 = int32(1)
	v142 = v77 + v141
	if base.Ui32(v142) < base.Ui32(int32(12)-v137&v141) {
		v77 = v142
		v78 = v137
		goto L19
	} else {
		goto L24
	}
L22:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+int32(2)+v77))))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v54+int32(16)+v77<<(uint(int32(2))%32))))
	v102 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+4)))
	v106 = F_findBucketForInsert(m, l0, v102, v15+int32(192), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+192))
	v109 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v106+v108<<(uint(v109)%32)+int32(16)))) = v101
	*(*uint8)(unsafe.Add(mBase, uint32(v106+v108+v109))) = uint8(v97)
	v119 = int32(1)
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106))))
	v126 = v119<<(uint(v108)%32)<<(uint(v119)%32)&int32(8190) | v125
	*(*uint16)(unsafe.Add(mBase, uint32(v106))) = uint16(v126)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v128 + int32(-1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v132 + v119
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54))))
	v137 = v136
	goto L21
L24:
	;
	goto L20
L25:
	;
	goto L2
L26:
	;
	v154 = int32(0)
	v162 = v154
	v164 = v148 + v17<<(uint(int32(6))%32)
	v165 = v154
	v167 = v154
	goto L27
L27:
	;
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164))))
	v175 = int32(0)
	v176 = v171
	v181 = v165
	v183 = v167
	goto L29
L29:
	;
	v185 = int32(1)
	if int32(base.Ui32(int32(base.Ui32(v176)>>(uint(v185)%32))&int32(4095))>>(uint(v175)%32))&v185 == int32(0) {
		v207 = v176
		v208 = v181
		v209 = v183
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v217 = int32(0)
	if v214 == v217 {
		v226 = v217
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v210 = int32(1)
	v211 = v175 + v210
	v214 = v207 & v210
	if base.Ui32(v211) < base.Ui32(int32(12)-v214) {
		v175 = v211
		v176 = v207
		v181 = v208
		v183 = v209
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v196 = int32(2)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v164+int32(16)+v175<<(uint(v196)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(192)+v183<<(uint(v196)%32)))) = v202
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164))))
	v206 = v183 + int32(1)
	v207 = v204
	v208 = v206
	v209 = v206
	goto L31
L33:
	;
	goto L30
L34:
	;
	if v208 < int32(1) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v164)+60))
	if base.Ui32(int32(2)) < base.Ui32(v162) {
		v226 = v220
		goto L34
	} else {
		goto L36
	}
L36:
	;
	if v220 != 0 {
		v162 = v162 + int32(1)
		v164 = v220
		v165 = v208
		v167 = v209
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v226 = v220
	goto L34
L38:
	;
	v568 = int32(0)
	if v226 == v568 {
		goto L2
	} else {
		goto L78
	}
L39:
	;
	if v208 == int32(1) {
		v287 = int32(0)
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v208&int32(1) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L41:
	;
	v234 = int32(0)
	v238 = v234
	v246 = v234
	goto L42
L42:
	;
	v251 = v238 << (uint(int32(2)) % 32)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(192)+v251)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	if v255 == int32(0) {
		v260 = v253
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v287 = v281
	goto L40
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v251))) = v260
	v268 = (v238 | int32(1)) << (uint(int32(2)) % 32)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(192)+v268)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if v272 == int32(0) {
		v277 = v270
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v258 = m.T0[v255].(func(*base.Module, int32) int32)(m, v253)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	v260 = v258
	goto L44
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v268))) = v277
	v280 = int32(2)
	v281 = v238 + v280
	v283 = v246 + v280
	if v283 != v208&int32(-2) {
		v238 = v281
		v246 = v283
		goto L42
	} else {
		goto L50
	}
L48:
	;
	v275 = m.T0[v272].(func(*base.Module, int32) int32)(m, v270)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	v277 = v275
	goto L47
L50:
	;
	goto L43
L51:
	;
	v323 = int32(0)
	goto L56
L52:
	;
	v304 = v287 << (uint(int32(2)) % 32)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(192)+v304)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	if v308 == int32(0) {
		v313 = v306
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v304))) = v313
	goto L51
L54:
	;
	v311 = m.T0[v308].(func(*base.Module, int32) int32)(m, v306)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	v313 = v311
	goto L53
L56:
	;
	v333 = v323 << (uint(int32(2)) % 32)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v15+v333)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+396)) = v335
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v338 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L38
L58:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(192)+v333)))
	v521 = F_findBucketForInsert(m, l0, v513, v15+int32(396), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L11
	} else {
		goto L76
	}
L59:
	;
	v344 = v15 + int32(396)
	v345 = int32(4)
	v346 = int32(_a_F_rehashStep_2)
	v354 = *(*int64)(unsafe.Add(mBase, _c_F_rehashStep[0]))
	v356 = v354 ^ int64(8317987319222330741)
	v357 = *(*int64)(unsafe.Add(mBase, _c_F_rehashStep[1]))
	v359 = v357 ^ int64(7237128888997146477)
	v361 = v354 ^ int64(7816392313619706465)
	v363 = v357 ^ int64(8387220255154660723)
	v368 = v15 + int32(400) - v345
	if v344 == v368 {
		v406 = v344
		v409 = v361
		v410 = v356
		v411 = v363
		v412 = v359
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v341 = m.T0[v338].(func(*base.Module, int32) int64)(m, v335)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	v513 = v341
	goto L58
L62:
	;
	v513 = base.I64_rotl(base.I64_rotl(v485, v452)^v494, v456) ^ base.I64_rotl(v504, v478) ^ base.I64_rotl(v508, v459) ^ v508
	goto L58
L63:
	;
	v417 = base.I64_extend_i32_u(v345) << (uint(int64(56)) % 64)
	switch v345 {
	default:
		v450 = v417
		goto L68
	case 1:
		v447 = v417
		goto L69
	case 2:
		v442 = v417
		goto L70
	case 3:
		v437 = v417
		goto L71
	case 4:
		v432 = v417
		goto L72
	case 5:
		v427 = v417
		goto L73
	case 6:
		v422 = v417
		goto L74
	case 7:
		goto L75
	}
L64:
	;
	v370 = v344
	v373 = v361
	v374 = v356
	v375 = v363
	v376 = v359
	goto L65
L65:
	;
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v370)))
	v381 = v380 ^ v375
	v382 = v381 + v373
	v383 = v374 + v376
	v386 = v383 ^ base.I64_rotl(v376, int64(13))
	v387 = v382 + v386
	v390 = v387 ^ base.I64_rotl(v386, int64(17))
	v393 = base.I64_rotl(v381, int64(16)) ^ v382
	v396 = int64(32)
	v398 = v393 + base.I64_rotl(v383, v396)
	v399 = base.I64_rotl(v393, int64(21)) ^ v398
	v401 = base.I64_rotl(v387, v396)
	v402 = v398 ^ v380
	v404 = v370 + int32(8)
	if v404 != v368 {
		v370 = v404
		v373 = v401
		v374 = v402
		v375 = v399
		v376 = v390
		goto L65
	} else {
		goto L67
	}
L66:
	;
	v406 = v368
	v409 = v401
	v410 = v402
	v411 = v399
	v412 = v390
	goto L63
L67:
	;
	goto L66
L68:
	;
	v451 = v450 ^ v411
	v452 = int64(16)
	v454 = v451 + v409
	v455 = base.I64_rotl(v451, v452) ^ v454
	v456 = int64(21)
	v458 = v410 + v412
	v459 = int64(32)
	v461 = v455 + base.I64_rotl(v458, v459)
	v462 = base.I64_rotl(v455, v456) ^ v461
	v465 = int64(13)
	v467 = v458 ^ base.I64_rotl(v412, v465)
	v468 = v454 + v467
	v473 = base.I64_rotl(v468, v459) ^ int64(255) + v462
	v474 = base.I64_rotl(v462, v452) ^ v473
	v478 = int64(17)
	v480 = v468 ^ base.I64_rotl(v467, v478)
	v481 = v461 ^ v450 + v480
	v484 = base.I64_rotl(v481, v459) + v474
	v485 = base.I64_rotl(v474, v456) ^ v484
	v490 = v481 ^ base.I64_rotl(v480, v465)
	v491 = v490 + v473
	v494 = base.I64_rotl(v491, v459) + v485
	v500 = base.I64_rotl(v490, v478) ^ v491
	v504 = base.I64_rotl(v500, v465) ^ (v500 + v484)
	v508 = v504 + v494
	goto L62
L69:
	;
	v448 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	v450 = v447 | v448
	goto L68
L70:
	;
	v443 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v406)+1)))
	v447 = v443<<(uint(int64(8))%64) | v442
	goto L69
L71:
	;
	v438 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v406)+2)))
	v442 = v438<<(uint(int64(16))%64) | v437
	goto L70
L72:
	;
	v433 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v406)+3)))
	v437 = v433<<(uint(int64(24))%64) | v432
	goto L71
L73:
	;
	v428 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v406)+4)))
	v432 = v428<<(uint(int64(32))%64) | v427
	goto L72
L74:
	;
	v423 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v406)+5)))
	v427 = v423<<(uint(int64(40))%64) | v422
	goto L73
L75:
	;
	v418 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v406)+6)))
	v422 = v418<<(uint(int64(48))%64) | v417
	goto L74
L76:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v15)+396))
	v524 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v521+v523<<(uint(v524)%32)+int32(16)))) = v517
	v534 = int64(base.Ui64(v513) >> (uint(int64(56)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v521+v523+v524))) = uint8(v534)
	v536 = int32(1)
	v542 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v521))))
	v543 = v536<<(uint(v523)%32)<<(uint(v536)%32)&int32(8190) | v542
	*(*uint16)(unsafe.Add(mBase, uint32(v521))) = uint16(v543)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v545 + int32(-1)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v549 + v536
	v554 = v323 + v536
	if v554 != v208 {
		v323 = v554
		goto L56
	} else {
		goto L77
	}
L77:
	;
	goto L57
L78:
	;
	v162 = v568
	v164 = v226
	v165 = v568
	v167 = v568
	goto L27
L79:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	goto L1
}
func F_rehashStepFinalize(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = v8 << (uint(int32(6)) % 32)
	v11 = v7 + v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
	if v12&int32(1) == int32(0) {
		v53 = v7
		v57 = v12
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v60 = v57 & int32(57344)
	*(*uint16)(unsafe.Add(mBase, uint32(v53+v10))) = uint16(v60)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v64 = v62 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_rehashStepFinalize[0]))
	if v67 != 0 {
		v73 = v64
		v74 = v67
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
	if v17 == int32(0) {
		v53 = v7
		v57 = v12
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = v17
	goto L4
L4:
	;
	v26 = int32(0)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v27&int32(1) == v26 {
		v33 = v26
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+v8<<(uint(int32(6))%32)))))
	v53 = v47
	v57 = v51
	goto L1
L6:
	;
	F_valkey_free(m, v24)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	v33 = v32
	goto L6
L8:
	;
	return
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	if v37 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v43 + int32(-1)
	if v33 != 0 {
		v24 = v33
		goto L4
	} else {
		goto L13
	}
L11:
	;
	m.T0[v37].(func(*base.Module, int32, int32))(m, l0, int32(-64))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	goto L5
L14:
	;
	v78 = int32(base.Ui32(v74)>>(uint(int32(2))%32)) & int32(1073741808)
	v79 = base.I32_rem_u_s(v73, v78)
	if v79 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v70 = F_sysconf(m, int32(30))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_rehashStepFinalize[0])) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v73 = v72
	v74 = v70
	goto L14
L16:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v95 != 0 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = int32(6)
	if base.Ui32((v80+v73<<(uint(v81)%32)-v78<<(uint(v81)%32))&(int32(0)-v74)) < base.Ui32(v80) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	return
L21:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v99 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v96 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_rehashingCompleted(m, l0)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	F__serverAssert(m, int32(_a_F_rehashStepFinalize_0), int32(_a_F_rehashStepFinalize_1), int32(623))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L28
	}
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(base.Ui32(v102)>>(uint(v99)%32)) == int32(0) {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_rehashingCompleted(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+32))
	if v5 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v10 == int32(0) {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v30 != int32(-1) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(-1)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v44 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
				v48 = int32(255)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v48)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v47)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v44
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v51
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v44
				return
			} else {
				F__serverAssert(m, int32(_a_F_rehashingCompleted_0), int32(_a_F_rehashingCompleted_1), int32(484))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			F_valkey_free(m, v10)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
				if v16 == int32(0) {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v30 != int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(-1)
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v44 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
						v48 = int32(255)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v48)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v47)
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v44
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v51
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v44
						return
					} else {
						F__serverAssert(m, int32(_a_F_rehashingCompleted_0), int32(_a_F_rehashingCompleted_1), int32(484))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				} else {
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
					if v21 == int32(255) {
						v25 = int32(0)
					} else {
						v25 = int32(-64) << (uint(v21) % 32)
					}
					m.T0[v16].(func(*base.Module, int32, int32))(m, l0, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v30 != int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(-1)
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v44 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
							v48 = int32(255)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v48)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v47)
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v44
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v51
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v44
							return
						} else {
							F__serverAssert(m, int32(_a_F_rehashingCompleted_0), int32(_a_F_rehashingCompleted_1), int32(484))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
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
				}
			}
		}
	} else {
		m.T0[v5].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v10 == int32(0) {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v30 != int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(-1)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v44 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
					v48 = int32(255)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v48)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v47)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v44
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v51
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v44
					return
				} else {
					F__serverAssert(m, int32(_a_F_rehashingCompleted_0), int32(_a_F_rehashingCompleted_1), int32(484))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				F_valkey_free(m, v10)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
					if v16 == int32(0) {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v30 != int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(-1)
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v44 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
							v48 = int32(255)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v48)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v47)
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v44
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v51
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v44
							return
						} else {
							F__serverAssert(m, int32(_a_F_rehashingCompleted_0), int32(_a_F_rehashingCompleted_1), int32(484))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					} else {
						v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v21 == int32(255) {
							v25 = int32(0)
						} else {
							v25 = int32(-64) << (uint(v21) % 32)
						}
						m.T0[v16].(func(*base.Module, int32, int32))(m, l0, v25)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v30 != int32(-1) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(-1)
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v44 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43
								v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
								v48 = int32(255)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v48)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v47)
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v44
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v51
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v44
								return
							} else {
								F__serverAssert(m, int32(_a_F_rehashingCompleted_0), int32(_a_F_rehashingCompleted_1), int32(484))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
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
					}
				}
			}
		}
	}
}
func F_releaseBufReferences(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
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
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	v7 = l0 + l1
	if l1 == int32(0) {
		v67 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v67 == v7 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v10 = l0
	goto L3
L3:
	;
	v17 = v10 + int32(12)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+10)))
	if v18&int32(1) == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v67 = v65
	goto L1
L5:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v65 = v17 + v64
	if base.Ui32(v65) < base.Ui32(v7) {
		v10 = v65
		goto L3
	} else {
		goto L19
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
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+10)))
	if v35&int32(2) != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
	v26 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)) = uint8(v26)
	if v25 == v26 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+272))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+272)) = v30 - v31
	goto L7
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v42 == int32(0) {
		goto L5
	} else {
		goto L14
	}
L11:
	;
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+8)))
	v39 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)))
	F_clusterSlotStatsAddNetworkBytesOutForSlot(m, v38, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	goto L10
L14:
	;
	v46 = v17
	v50 = v42
	goto L15
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	F_decrRefCount(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L17
	}
L16:
	;
	goto L5
L17:
	;
	v57 = v50 + int32(-8)
	if v57 != 0 {
		v46 = v46 + int32(8)
		v50 = v57
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	goto L4
L20:
	;
	return
L21:
	;
	F__serverAssert(m, int32(_a_F_releaseBufReferences_0), int32(_a_F_releaseBufReferences_1), int32(2989))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_releaseInstanceLink(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 <= int32(0) {
		F__serverAssert(m, int32(_a_F_releaseInstanceLink_0), int32(_a_F_releaseInstanceLink_1), int32(1055))
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = v4 + int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v8
		if v8 == int32(0) {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v32 == int32(0) {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v49 == int32(0) {
					F_valkey_free(m, l0)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v67 = int32(0)
						return v67
					}
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v52 != v49 {
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
					}
					v56 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v56
					*(*int32)(unsafe.Add(mBase, uint32(v49)+212)) = v56
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
					F_valkeyAsyncFree(m, v49)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						F_valkey_free(m, l0)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v67 = int32(0)
							return v67
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v37 != v32 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v32)+212)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
				F_valkeyAsyncFree(m, v32)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v49 == int32(0) {
						F_valkey_free(m, l0)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v67 = int32(0)
							return v67
						}
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v52 != v49 {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
						}
						v56 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v49)+212)) = v56
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
						F_valkeyAsyncFree(m, v49)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_valkey_free(m, l0)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								v67 = int32(0)
								return v67
							}
						}
					}
				}
			}
		} else {
			if l1 == int32(0) {
				v67 = l0
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
				if v15 == int32(0) {
					v67 = l0
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+256))
					if v19 == int32(0) {
						v67 = l0
					} else {
						v24 = v19
						for {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
							if v25 != l1 {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(1004)
							}
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
							if v31 != 0 {
								v24 = v31
								continue
							} else {
								break
							}
							break
						}
						v67 = l0
					}
				}
			}
			return v67
		}
	}
}
func F_removeMatchingSentinelFromPrimary(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v208 int64
	_ = v208
	var v209 int64
	_ = v209
	var v210 int64
	_ = v210
	var v212 int64
	_ = v212
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v220 int64
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v11 = F_dictGetSafeIterator(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return int32(0)
L3:
	;
	F_dictReleaseIterator(m, v11)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L75
	}
L4:
	;
	return int32(0)
L5:
	;
	v22 = v11 + int32(20)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v118 == int32(0) {
		v281 = v9
		goto L3
	} else {
		goto L32
	}
L7:
	;
	v29 = v22
	v30 = v26
	goto L10
L8:
	;
	v26 = int32(1)
	goto L7
L9:
	;
	v26 = int32(0)
	goto L7
L10:
	;
	switch v30 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v30 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v110
	if v110 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v34 != int32(-1) {
		v73 = v34
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v74 = int32(1)
	v75 = v73 + v74
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v75
	v77 = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v81+int32(26)))))
	if v85 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v38 != 0 {
		v73 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v40 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	if v67 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v47 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v39)+16)))
	v48 = int64(*(*int8)(unsafe.Add(mBase, uint32(v39)+27)))
	v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(v39)+8)))
	v50 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v39)+12)))
	v51 = int64(*(*int8)(unsafe.Add(mBase, uint32(v39)+26)))
	v52 = int64(*(*int32)(unsafe.Add(mBase, uint32(v39)+4)))
	v53 = F_wangHash64(m, v52)
	mBase = m.M
	v55 = F_wangHash64(m, v51+v53)
	mBase = m.M
	v57 = F_wangHash64(m, v50+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v49+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v48+v59)
	mBase = m.M
	v63 = F_wangHash64(m, v47+v61)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v66 = v65
	goto L19
L21:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+24)))
	v45 = v43 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+24)) = uint16(v45)
	v66 = v39
	goto L19
L22:
	;
	v73 = v67 + int32(-1)
	goto L16
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v73 = v70
	goto L16
L24:
	;
	v100 = int32(2)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v80+v98<<(uint(v100)%32)+int32(4))))
	v29 = v105 + v99<<(uint(v100)%32)
	v30 = int32(1)
	goto L10
L25:
	;
	v89 = v77
	goto L27
L26:
	;
	v89 = v74 << (uint(v85) % 32)
	goto L27
L27:
	;
	if v75 < v89 {
		v98 = v81
		v99 = v75
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v81 != 0 {
		v118 = v77
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	if v91 == int32(-1) {
		v118 = v77
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(4294967296)
	v98 = int32(1)
	v99 = int32(0)
	goto L24
L31:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v114
	v118 = v110
	goto L13
L32:
	;
	v126 = v9
	v128 = v118
	goto L33
L33:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	goto L36
L34:
	;
	v281 = v171
	goto L3
L35:
	;
	v179 = v11 + int32(20)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v180 != 0 {
		goto L50
	} else {
		goto L51
	}
L36:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	if v131 == int32(0) {
		v171 = v126
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v137 == int32(0) {
		v160 = v136
		v161 = v137
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v161-v160&int32(255) != 0 {
		v171 = v126
		goto L35
	} else {
		goto L46
	}
L39:
	;
	goto L38
L40:
	;
	if v137 != v136&int32(255) {
		v160 = v136
		v161 = v137
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v143 = v131
	v144 = l1
	goto L42
L42:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+1)))
	if v148 == int32(0) {
		v160 = v147
		v161 = v148
		goto L39
	} else {
		goto L44
	}
L43:
	;
	v160 = v147
	v161 = v148
	goto L39
L44:
	;
	v151 = int32(1)
	if v148 == v147&int32(255) {
		v143 = v143 + v151
		v144 = v144 + v151
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v167 = F_dictDelete(m, v165, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v171 = v126 + int32(1)
	goto L35
L48:
	;
	if v275 != 0 {
		v126 = v171
		v128 = v275
		goto L33
	} else {
		goto L74
	}
L49:
	;
	v186 = v179
	v187 = v183
	goto L52
L50:
	;
	v183 = int32(1)
	goto L49
L51:
	;
	v183 = int32(0)
	goto L49
L52:
	;
	switch v187 {
	case 0:
		goto L57
	default:
		goto L56
	}
L54:
	;
	v187 = int32(0)
	goto L52
L55:
	;
	goto L48
L56:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v267
	if v267 == int32(0) {
		goto L54
	} else {
		goto L73
	}
L57:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v191 != int32(-1) {
		v230 = v191
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v231 = int32(1)
	v232 = v230 + v231
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v232
	v234 = int32(0)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237+v238+int32(26)))))
	if v242 == int32(255) {
		goto L67
	} else {
		goto L68
	}
L59:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v195 != 0 {
		v230 = int32(-1)
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v197 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+20))
	if v224 != int32(-1) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v204 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v196)+16)))
	v205 = int64(*(*int8)(unsafe.Add(mBase, uint32(v196)+27)))
	v206 = int64(*(*int32)(unsafe.Add(mBase, uint32(v196)+8)))
	v207 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v196)+12)))
	v208 = int64(*(*int8)(unsafe.Add(mBase, uint32(v196)+26)))
	v209 = int64(*(*int32)(unsafe.Add(mBase, uint32(v196)+4)))
	v210 = F_wangHash64(m, v209)
	mBase = m.M
	v212 = F_wangHash64(m, v208+v210)
	mBase = m.M
	v214 = F_wangHash64(m, v207+v212)
	mBase = m.M
	v216 = F_wangHash64(m, v206+v214)
	mBase = m.M
	v218 = F_wangHash64(m, v205+v216)
	mBase = m.M
	v220 = F_wangHash64(m, v204+v218)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v223 = v222
	goto L61
L63:
	;
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196)+24)))
	v202 = v200 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v196)+24)) = uint16(v202)
	v223 = v196
	goto L61
L64:
	;
	v230 = v224 + int32(-1)
	goto L58
L65:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v230 = v227
	goto L58
L66:
	;
	v257 = int32(2)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v237+v255<<(uint(v257)%32)+int32(4))))
	v186 = v262 + v256<<(uint(v257)%32)
	v187 = int32(1)
	goto L52
L67:
	;
	v246 = v234
	goto L69
L68:
	;
	v246 = v231 << (uint(v242) % 32)
	goto L69
L69:
	;
	if v232 < v246 {
		v255 = v238
		v256 = v232
		goto L66
	} else {
		goto L70
	}
L70:
	;
	if v238 != 0 {
		v275 = v234
		goto L55
	} else {
		goto L71
	}
L71:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v237)+20))
	if v248 == int32(-1) {
		v275 = v234
		goto L55
	} else {
		goto L72
	}
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(4294967296)
	v255 = int32(1)
	v256 = int32(0)
	goto L66
L73:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = v271
	v275 = v267
	goto L55
L74:
	;
	goto L34
L75:
	;
	return v281
}
func F_renameGenericCommand(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
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
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
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
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
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
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = F_objectGetVal(m, v15)
	mBase = m.M
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v19 = F_objectGetVal(m, v18)
	mBase = m.M
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-1)))))
	switch v27 & int32(7) {
	case 0:
		v44 = int32(base.Ui32(v27) >> (uint(int32(3)) % 32))
	case 1:
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
		v44 = v34
	case 2:
		v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
		v44 = v37
	case 3:
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
		v44 = v40
	case 4:
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
		v44 = v43
	default:
		v44 = v3
	}
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(-1)))))
	switch v47 & int32(7) {
	case 0:
		v64 = int32(base.Ui32(v47) >> (uint(int32(3)) % 32))
	case 1:
		v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(-3)))))
		v64 = v54
	case 2:
		v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+int32(-5)))))
		v64 = v57
	case 3:
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(-9))))
		v64 = v60
	case 4:
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(-17))))
		v64 = v63
	default:
		v64 = v3
	}
	v65 = base.B2i32(base.Ui32(v44) < base.Ui32(v64))
	if base.Ui32(v44) < base.Ui32(v64) {
		v66 = v44
	} else {
		v66 = v64
	}
	v67 = F_memcmp(m, v16, v19, v66)
	mBase = m.M
	if v67 != 0 {
		v70 = v67
	} else {
		v70 = base.B2i32(base.Ui32(v64) < base.Ui32(v44)) - v65
	}
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_renameGenericCommand[0]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v77 = F_lookupKey(m, v73, v75, int32(8))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		return
	} else {
		if v77 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v77
			if v70 != 0 {
				F_incrRefCount(m, v77)
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
					if v95&int32(1) == int32(0) {
						v106 = int64(-1)
					} else {
						v105 = *(*int64)(unsafe.Add(mBase, uint32(v77+(v95&int32(4)^int32(12)))))
						v106 = v105
					}
					v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
					v111 = F_lookupKey(m, v107, v109, int32(8))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						if v111 == int32(0) {
							v142 = int32(_a_F_renameGenericCommand_0)
							v143 = *(*int32)(unsafe.Add(mBase, _c_F_renameGenericCommand[1]))
							v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
							v147 = F_objectGetVal(m, v146)
							mBase = m.M
							v148 = int32(0)
							v150 = *(*int32)(unsafe.Add(mBase, _c_F_renameGenericCommand[2]))
							if v150 == v148 {
								v155 = v148
								v157 = F_dbGenericDeleteWithDictIndex(m, v144, v146, v143, int32(1), v155)
								mBase = m.M
								v158 = m.ExcPending
								if v158 != 0 {
									return
								} else {
									v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
									v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
									F_dbAddInternal(m, v159, v161, v12+int32(12), int32(0))
									mBase = m.M
									v166 = m.ExcPending
									if v166 != 0 {
										return
									} else {
										if v106 == int64(-1) {
											v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
											F_touchWatchedKey(m, v174, v176)
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return
											} else {
												F_trackingInvalidateKey(m, l0, v176, int32(1))
												mBase = m.M
												v181 = m.ExcPending
												if v181 != 0 {
													return
												} else {
													v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
													v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
													F_touchWatchedKey(m, v182, v184)
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return
													} else {
														F_trackingInvalidateKey(m, l0, v184, int32(1))
														mBase = m.M
														v189 = m.ExcPending
														if v189 != 0 {
															return
														} else {
															v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
															v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
															F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_1), v193, v195)
															mBase = m.M
															v197 = m.ExcPending
															if v197 != 0 {
																return
															} else {
																v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
																v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
																F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_2), v201, v203)
																mBase = m.M
																v205 = m.ExcPending
																if v205 != 0 {
																	return
																} else {
																	v206 = int32(_a_F_renameGenericCommand_0)
																	v208 = *(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3]))
																	*(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3])) = v208 + int64(1)
																	if l1 != 0 {
																		v216 = int32(_a_F_renameGenericCommand_3)
																	} else {
																		v216 = int32(_a_F_renameGenericCommand_4)
																	}
																	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
																	F_addReply(m, l0, v217)
																	mBase = m.M
																	v219 = m.ExcPending
																	if v219 != 0 {
																		return
																	} else {
																		m.G0 = v12 + int32(16)
																		return
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
											v172 = F_setExpire(m, l0, v169, v171, v106)
											mBase = m.M
											v173 = m.ExcPending
											if v173 != 0 {
												return
											} else {
												v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
												v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
												F_touchWatchedKey(m, v174, v176)
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													F_trackingInvalidateKey(m, l0, v176, int32(1))
													mBase = m.M
													v181 = m.ExcPending
													if v181 != 0 {
														return
													} else {
														v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
														v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
														F_touchWatchedKey(m, v182, v184)
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return
														} else {
															F_trackingInvalidateKey(m, l0, v184, int32(1))
															mBase = m.M
															v189 = m.ExcPending
															if v189 != 0 {
																return
															} else {
																v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
																v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
																F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_1), v193, v195)
																mBase = m.M
																v197 = m.ExcPending
																if v197 != 0 {
																	return
																} else {
																	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
																	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
																	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_2), v201, v203)
																	mBase = m.M
																	v205 = m.ExcPending
																	if v205 != 0 {
																		return
																	} else {
																		v206 = int32(_a_F_renameGenericCommand_0)
																		v208 = *(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3]))
																		*(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3])) = v208 + int64(1)
																		if l1 != 0 {
																			v216 = int32(_a_F_renameGenericCommand_3)
																		} else {
																			v216 = int32(_a_F_renameGenericCommand_4)
																		}
																		v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
																		F_addReply(m, l0, v217)
																		mBase = m.M
																		v219 = m.ExcPending
																		if v219 != 0 {
																			return
																		} else {
																			m.G0 = v12 + int32(16)
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
							} else {
								v153 = F_getKeySlot(m, v147)
								mBase = m.M
								v154 = m.ExcPending
								if v154 != 0 {
									return
								} else {
									v155 = v153
									v157 = F_dbGenericDeleteWithDictIndex(m, v144, v146, v143, int32(1), v155)
									mBase = m.M
									v158 = m.ExcPending
									if v158 != 0 {
										return
									} else {
										v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
										v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
										F_dbAddInternal(m, v159, v161, v12+int32(12), int32(0))
										mBase = m.M
										v166 = m.ExcPending
										if v166 != 0 {
											return
										} else {
											if v106 == int64(-1) {
												v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
												v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
												F_touchWatchedKey(m, v174, v176)
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													F_trackingInvalidateKey(m, l0, v176, int32(1))
													mBase = m.M
													v181 = m.ExcPending
													if v181 != 0 {
														return
													} else {
														v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
														v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
														F_touchWatchedKey(m, v182, v184)
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return
														} else {
															F_trackingInvalidateKey(m, l0, v184, int32(1))
															mBase = m.M
															v189 = m.ExcPending
															if v189 != 0 {
																return
															} else {
																v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
																v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
																F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_1), v193, v195)
																mBase = m.M
																v197 = m.ExcPending
																if v197 != 0 {
																	return
																} else {
																	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
																	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
																	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_2), v201, v203)
																	mBase = m.M
																	v205 = m.ExcPending
																	if v205 != 0 {
																		return
																	} else {
																		v206 = int32(_a_F_renameGenericCommand_0)
																		v208 = *(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3]))
																		*(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3])) = v208 + int64(1)
																		if l1 != 0 {
																			v216 = int32(_a_F_renameGenericCommand_3)
																		} else {
																			v216 = int32(_a_F_renameGenericCommand_4)
																		}
																		v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
																		F_addReply(m, l0, v217)
																		mBase = m.M
																		v219 = m.ExcPending
																		if v219 != 0 {
																			return
																		} else {
																			m.G0 = v12 + int32(16)
																			return
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
												v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
												v172 = F_setExpire(m, l0, v169, v171, v106)
												mBase = m.M
												v173 = m.ExcPending
												if v173 != 0 {
													return
												} else {
													v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
													v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
													F_touchWatchedKey(m, v174, v176)
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return
													} else {
														F_trackingInvalidateKey(m, l0, v176, int32(1))
														mBase = m.M
														v181 = m.ExcPending
														if v181 != 0 {
															return
														} else {
															v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
															F_touchWatchedKey(m, v182, v184)
															mBase = m.M
															v186 = m.ExcPending
															if v186 != 0 {
																return
															} else {
																F_trackingInvalidateKey(m, l0, v184, int32(1))
																mBase = m.M
																v189 = m.ExcPending
																if v189 != 0 {
																	return
																} else {
																	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
																	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
																	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_1), v193, v195)
																	mBase = m.M
																	v197 = m.ExcPending
																	if v197 != 0 {
																		return
																	} else {
																		v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
																		v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																		v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
																		F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_2), v201, v203)
																		mBase = m.M
																		v205 = m.ExcPending
																		if v205 != 0 {
																			return
																		} else {
																			v206 = int32(_a_F_renameGenericCommand_0)
																			v208 = *(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3]))
																			*(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3])) = v208 + int64(1)
																			if l1 != 0 {
																				v216 = int32(_a_F_renameGenericCommand_3)
																			} else {
																				v216 = int32(_a_F_renameGenericCommand_4)
																			}
																			v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
																			F_addReply(m, l0, v217)
																			mBase = m.M
																			v219 = m.ExcPending
																			if v219 != 0 {
																				return
																			} else {
																				m.G0 = v12 + int32(16)
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
						} else {
							if l1 == int32(0) {
								v123 = int32(_a_F_renameGenericCommand_0)
								v124 = *(*int32)(unsafe.Add(mBase, _c_F_renameGenericCommand[1]))
								v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
								v128 = F_objectGetVal(m, v127)
								mBase = m.M
								v130 = *(*int32)(unsafe.Add(mBase, _c_F_renameGenericCommand[2]))
								if v130 != 0 {
									v132 = F_getKeySlot(m, v128)
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										v134 = v132
										v136 = F_dbGenericDeleteWithDictIndex(m, v125, v127, v124, int32(1), v134)
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return
										} else {
											v142 = int32(_a_F_renameGenericCommand_0)
											v143 = *(*int32)(unsafe.Add(mBase, _c_F_renameGenericCommand[1]))
											v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
											v147 = F_objectGetVal(m, v146)
											mBase = m.M
											v148 = int32(0)
											v150 = *(*int32)(unsafe.Add(mBase, _c_F_renameGenericCommand[2]))
											if v150 == v148 {
												v155 = v148
												v157 = F_dbGenericDeleteWithDictIndex(m, v144, v146, v143, int32(1), v155)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return
												} else {
													v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
													v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
													F_dbAddInternal(m, v159, v161, v12+int32(12), int32(0))
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return
													} else {
														if v106 == int64(-1) {
															v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
															F_touchWatchedKey(m, v174, v176)
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return
															} else {
																F_trackingInvalidateKey(m, l0, v176, int32(1))
																mBase = m.M
																v181 = m.ExcPending
																if v181 != 0 {
																	return
																} else {
																	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
																	F_touchWatchedKey(m, v182, v184)
																	mBase = m.M
																	v186 = m.ExcPending
																	if v186 != 0 {
																		return
																	} else {
																		F_trackingInvalidateKey(m, l0, v184, int32(1))
																		mBase = m.M
																		v189 = m.ExcPending
																		if v189 != 0 {
																			return
																		} else {
																			v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
																			v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																			v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
																			F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_1), v193, v195)
																			mBase = m.M
																			v197 = m.ExcPending
																			if v197 != 0 {
																				return
																			} else {
																				v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
																				v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																				v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
																				F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_2), v201, v203)
																				mBase = m.M
																				v205 = m.ExcPending
																				if v205 != 0 {
																					return
																				} else {
																					v206 = int32(_a_F_renameGenericCommand_0)
																					v208 = *(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3]))
																					*(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3])) = v208 + int64(1)
																					if l1 != 0 {
																						v216 = int32(_a_F_renameGenericCommand_3)
																					} else {
																						v216 = int32(_a_F_renameGenericCommand_4)
																					}
																					v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
																					F_addReply(m, l0, v217)
																					mBase = m.M
																					v219 = m.ExcPending
																					if v219 != 0 {
																						return
																					} else {
																						m.G0 = v12 + int32(16)
																						return
																					}
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
															v172 = F_setExpire(m, l0, v169, v171, v106)
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return
															} else {
																v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
																F_touchWatchedKey(m, v174, v176)
																mBase = m.M
																v178 = m.ExcPending
																if v178 != 0 {
																	return
																} else {
																	F_trackingInvalidateKey(m, l0, v176, int32(1))
																	mBase = m.M
																	v181 = m.ExcPending
																	if v181 != 0 {
																		return
																	} else {
																		v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																		v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
																		F_touchWatchedKey(m, v182, v184)
																		mBase = m.M
																		v186 = m.ExcPending
																		if v186 != 0 {
																			return
																		} else {
																			F_trackingInvalidateKey(m, l0, v184, int32(1))
																			mBase = m.M
																			v189 = m.ExcPending
																			if v189 != 0 {
																				return
																			} else {
																				v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
																				v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																				v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
																				F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_1), v193, v195)
																				mBase = m.M
																				v197 = m.ExcPending
																				if v197 != 0 {
																					return
																				} else {
																					v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
																					v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																					v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
																					F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_2), v201, v203)
																					mBase = m.M
																					v205 = m.ExcPending
																					if v205 != 0 {
																						return
																					} else {
																						v206 = int32(_a_F_renameGenericCommand_0)
																						v208 = *(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3]))
																						*(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3])) = v208 + int64(1)
																						if l1 != 0 {
																							v216 = int32(_a_F_renameGenericCommand_3)
																						} else {
																							v216 = int32(_a_F_renameGenericCommand_4)
																						}
																						v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
																						F_addReply(m, l0, v217)
																						mBase = m.M
																						v219 = m.ExcPending
																						if v219 != 0 {
																							return
																						} else {
																							m.G0 = v12 + int32(16)
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
											} else {
												v153 = F_getKeySlot(m, v147)
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return
												} else {
													v155 = v153
													v157 = F_dbGenericDeleteWithDictIndex(m, v144, v146, v143, int32(1), v155)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return
													} else {
														v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
														v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
														F_dbAddInternal(m, v159, v161, v12+int32(12), int32(0))
														mBase = m.M
														v166 = m.ExcPending
														if v166 != 0 {
															return
														} else {
															if v106 == int64(-1) {
																v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
																F_touchWatchedKey(m, v174, v176)
																mBase = m.M
																v178 = m.ExcPending
																if v178 != 0 {
																	return
																} else {
																	F_trackingInvalidateKey(m, l0, v176, int32(1))
																	mBase = m.M
																	v181 = m.ExcPending
																	if v181 != 0 {
																		return
																	} else {
																		v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																		v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
																		F_touchWatchedKey(m, v182, v184)
																		mBase = m.M
																		v186 = m.ExcPending
																		if v186 != 0 {
																			return
																		} else {
																			F_trackingInvalidateKey(m, l0, v184, int32(1))
																			mBase = m.M
																			v189 = m.ExcPending
																			if v189 != 0 {
																				return
																			} else {
																				v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
																				v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																				v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
																				F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_1), v193, v195)
																				mBase = m.M
																				v197 = m.ExcPending
																				if v197 != 0 {
																					return
																				} else {
																					v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
																					v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																					v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
																					F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_2), v201, v203)
																					mBase = m.M
																					v205 = m.ExcPending
																					if v205 != 0 {
																						return
																					} else {
																						v206 = int32(_a_F_renameGenericCommand_0)
																						v208 = *(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3]))
																						*(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3])) = v208 + int64(1)
																						if l1 != 0 {
																							v216 = int32(_a_F_renameGenericCommand_3)
																						} else {
																							v216 = int32(_a_F_renameGenericCommand_4)
																						}
																						v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
																						F_addReply(m, l0, v217)
																						mBase = m.M
																						v219 = m.ExcPending
																						if v219 != 0 {
																							return
																						} else {
																							m.G0 = v12 + int32(16)
																							return
																						}
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
																v172 = F_setExpire(m, l0, v169, v171, v106)
																mBase = m.M
																v173 = m.ExcPending
																if v173 != 0 {
																	return
																} else {
																	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
																	F_touchWatchedKey(m, v174, v176)
																	mBase = m.M
																	v178 = m.ExcPending
																	if v178 != 0 {
																		return
																	} else {
																		F_trackingInvalidateKey(m, l0, v176, int32(1))
																		mBase = m.M
																		v181 = m.ExcPending
																		if v181 != 0 {
																			return
																		} else {
																			v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																			v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
																			F_touchWatchedKey(m, v182, v184)
																			mBase = m.M
																			v186 = m.ExcPending
																			if v186 != 0 {
																				return
																			} else {
																				F_trackingInvalidateKey(m, l0, v184, int32(1))
																				mBase = m.M
																				v189 = m.ExcPending
																				if v189 != 0 {
																					return
																				} else {
																					v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
																					v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																					v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
																					F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_1), v193, v195)
																					mBase = m.M
																					v197 = m.ExcPending
																					if v197 != 0 {
																						return
																					} else {
																						v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																						v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
																						v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																						v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
																						F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_2), v201, v203)
																						mBase = m.M
																						v205 = m.ExcPending
																						if v205 != 0 {
																							return
																						} else {
																							v206 = int32(_a_F_renameGenericCommand_0)
																							v208 = *(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3]))
																							*(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3])) = v208 + int64(1)
																							if l1 != 0 {
																								v216 = int32(_a_F_renameGenericCommand_3)
																							} else {
																								v216 = int32(_a_F_renameGenericCommand_4)
																							}
																							v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
																							F_addReply(m, l0, v217)
																							mBase = m.M
																							v219 = m.ExcPending
																							if v219 != 0 {
																								return
																							} else {
																								m.G0 = v12 + int32(16)
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
								} else {
									v134 = int32(0)
									v136 = F_dbGenericDeleteWithDictIndex(m, v125, v127, v124, int32(1), v134)
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return
									} else {
										v142 = int32(_a_F_renameGenericCommand_0)
										v143 = *(*int32)(unsafe.Add(mBase, _c_F_renameGenericCommand[1]))
										v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
										v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
										v147 = F_objectGetVal(m, v146)
										mBase = m.M
										v148 = int32(0)
										v150 = *(*int32)(unsafe.Add(mBase, _c_F_renameGenericCommand[2]))
										if v150 == v148 {
											v155 = v148
											v157 = F_dbGenericDeleteWithDictIndex(m, v144, v146, v143, int32(1), v155)
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return
											} else {
												v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
												v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
												F_dbAddInternal(m, v159, v161, v12+int32(12), int32(0))
												mBase = m.M
												v166 = m.ExcPending
												if v166 != 0 {
													return
												} else {
													if v106 == int64(-1) {
														v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
														v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
														F_touchWatchedKey(m, v174, v176)
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return
														} else {
															F_trackingInvalidateKey(m, l0, v176, int32(1))
															mBase = m.M
															v181 = m.ExcPending
															if v181 != 0 {
																return
															} else {
																v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
																F_touchWatchedKey(m, v182, v184)
																mBase = m.M
																v186 = m.ExcPending
																if v186 != 0 {
																	return
																} else {
																	F_trackingInvalidateKey(m, l0, v184, int32(1))
																	mBase = m.M
																	v189 = m.ExcPending
																	if v189 != 0 {
																		return
																	} else {
																		v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
																		v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																		v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
																		F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_1), v193, v195)
																		mBase = m.M
																		v197 = m.ExcPending
																		if v197 != 0 {
																			return
																		} else {
																			v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
																			v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																			v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
																			F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_2), v201, v203)
																			mBase = m.M
																			v205 = m.ExcPending
																			if v205 != 0 {
																				return
																			} else {
																				v206 = int32(_a_F_renameGenericCommand_0)
																				v208 = *(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3]))
																				*(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3])) = v208 + int64(1)
																				if l1 != 0 {
																					v216 = int32(_a_F_renameGenericCommand_3)
																				} else {
																					v216 = int32(_a_F_renameGenericCommand_4)
																				}
																				v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
																				F_addReply(m, l0, v217)
																				mBase = m.M
																				v219 = m.ExcPending
																				if v219 != 0 {
																					return
																				} else {
																					m.G0 = v12 + int32(16)
																					return
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
														v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
														v172 = F_setExpire(m, l0, v169, v171, v106)
														mBase = m.M
														v173 = m.ExcPending
														if v173 != 0 {
															return
														} else {
															v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
															F_touchWatchedKey(m, v174, v176)
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return
															} else {
																F_trackingInvalidateKey(m, l0, v176, int32(1))
																mBase = m.M
																v181 = m.ExcPending
																if v181 != 0 {
																	return
																} else {
																	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
																	F_touchWatchedKey(m, v182, v184)
																	mBase = m.M
																	v186 = m.ExcPending
																	if v186 != 0 {
																		return
																	} else {
																		F_trackingInvalidateKey(m, l0, v184, int32(1))
																		mBase = m.M
																		v189 = m.ExcPending
																		if v189 != 0 {
																			return
																		} else {
																			v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
																			v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																			v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
																			F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_1), v193, v195)
																			mBase = m.M
																			v197 = m.ExcPending
																			if v197 != 0 {
																				return
																			} else {
																				v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
																				v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																				v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
																				F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_2), v201, v203)
																				mBase = m.M
																				v205 = m.ExcPending
																				if v205 != 0 {
																					return
																				} else {
																					v206 = int32(_a_F_renameGenericCommand_0)
																					v208 = *(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3]))
																					*(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3])) = v208 + int64(1)
																					if l1 != 0 {
																						v216 = int32(_a_F_renameGenericCommand_3)
																					} else {
																						v216 = int32(_a_F_renameGenericCommand_4)
																					}
																					v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
																					F_addReply(m, l0, v217)
																					mBase = m.M
																					v219 = m.ExcPending
																					if v219 != 0 {
																						return
																					} else {
																						m.G0 = v12 + int32(16)
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
										} else {
											v153 = F_getKeySlot(m, v147)
											mBase = m.M
											v154 = m.ExcPending
											if v154 != 0 {
												return
											} else {
												v155 = v153
												v157 = F_dbGenericDeleteWithDictIndex(m, v144, v146, v143, int32(1), v155)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return
												} else {
													v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
													v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
													F_dbAddInternal(m, v159, v161, v12+int32(12), int32(0))
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return
													} else {
														if v106 == int64(-1) {
															v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
															F_touchWatchedKey(m, v174, v176)
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return
															} else {
																F_trackingInvalidateKey(m, l0, v176, int32(1))
																mBase = m.M
																v181 = m.ExcPending
																if v181 != 0 {
																	return
																} else {
																	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
																	F_touchWatchedKey(m, v182, v184)
																	mBase = m.M
																	v186 = m.ExcPending
																	if v186 != 0 {
																		return
																	} else {
																		F_trackingInvalidateKey(m, l0, v184, int32(1))
																		mBase = m.M
																		v189 = m.ExcPending
																		if v189 != 0 {
																			return
																		} else {
																			v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
																			v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																			v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
																			F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_1), v193, v195)
																			mBase = m.M
																			v197 = m.ExcPending
																			if v197 != 0 {
																				return
																			} else {
																				v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
																				v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																				v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
																				F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_2), v201, v203)
																				mBase = m.M
																				v205 = m.ExcPending
																				if v205 != 0 {
																					return
																				} else {
																					v206 = int32(_a_F_renameGenericCommand_0)
																					v208 = *(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3]))
																					*(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3])) = v208 + int64(1)
																					if l1 != 0 {
																						v216 = int32(_a_F_renameGenericCommand_3)
																					} else {
																						v216 = int32(_a_F_renameGenericCommand_4)
																					}
																					v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
																					F_addReply(m, l0, v217)
																					mBase = m.M
																					v219 = m.ExcPending
																					if v219 != 0 {
																						return
																					} else {
																						m.G0 = v12 + int32(16)
																						return
																					}
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
															v172 = F_setExpire(m, l0, v169, v171, v106)
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return
															} else {
																v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
																F_touchWatchedKey(m, v174, v176)
																mBase = m.M
																v178 = m.ExcPending
																if v178 != 0 {
																	return
																} else {
																	F_trackingInvalidateKey(m, l0, v176, int32(1))
																	mBase = m.M
																	v181 = m.ExcPending
																	if v181 != 0 {
																		return
																	} else {
																		v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																		v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
																		F_touchWatchedKey(m, v182, v184)
																		mBase = m.M
																		v186 = m.ExcPending
																		if v186 != 0 {
																			return
																		} else {
																			F_trackingInvalidateKey(m, l0, v184, int32(1))
																			mBase = m.M
																			v189 = m.ExcPending
																			if v189 != 0 {
																				return
																			} else {
																				v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
																				v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																				v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
																				F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_1), v193, v195)
																				mBase = m.M
																				v197 = m.ExcPending
																				if v197 != 0 {
																					return
																				} else {
																					v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
																					v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																					v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
																					F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_renameGenericCommand_2), v201, v203)
																					mBase = m.M
																					v205 = m.ExcPending
																					if v205 != 0 {
																						return
																					} else {
																						v206 = int32(_a_F_renameGenericCommand_0)
																						v208 = *(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3]))
																						*(*int64)(unsafe.Add(mBase, _c_F_renameGenericCommand[3])) = v208 + int64(1)
																						if l1 != 0 {
																							v216 = int32(_a_F_renameGenericCommand_3)
																						} else {
																							v216 = int32(_a_F_renameGenericCommand_4)
																						}
																						v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
																						F_addReply(m, l0, v217)
																						mBase = m.M
																						v219 = m.ExcPending
																						if v219 != 0 {
																							return
																						} else {
																							m.G0 = v12 + int32(16)
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
							} else {
								F_decrRefCount(m, v77)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return
								} else {
									v120 = *(*int32)(unsafe.Add(mBase, _c_F_renameGenericCommand[4]))
									F_addReply(m, l0, v120)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										m.G0 = v12 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			} else {
				if l1 != 0 {
					v86 = int32(_a_F_renameGenericCommand_5)
				} else {
					v86 = int32(_a_F_renameGenericCommand_4)
				}
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
				F_addReply(m, l0, v87)
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return
				} else {
					m.G0 = v12 + int32(16)
					return
				}
			}
		} else {
			F_addReplyOrErrorObject(m, l0, v72)
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return
			} else {
				m.G0 = v12 + int32(16)
				return
			}
		}
	}
}
func F_replconfCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
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
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
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
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int64
	_ = v501
	var v502 int32
	_ = v502
	var v503 int64
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int64
	_ = v564
	var v565 int32
	_ = v565
	var v566 int64
	_ = v566
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int64
	_ = v573
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
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
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v963 int32
	_ = v963
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int64
	_ = v1106
	var v1107 int64
	_ = v1107
	var v1109 int64
	_ = v1109
	var v1111 int64
	_ = v1111
	var v1114 int64
	_ = v1114
	var v1116 int64
	_ = v1116
	var v1118 int64
	_ = v1118
	var v1120 int64
	_ = v1120
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1188 int32
	_ = v1188
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1254 int32
	_ = v1254
	var v1261 int32
	_ = v1261
	var v1266 int32
	_ = v1266
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1316 int32
	_ = v1316
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1347 int64
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v15&int32(1) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(112)
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v22 != 0 {
		v28 = v15
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_replconfCommand[0]))
	F_addReplyErrorObject(m, l0, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	goto L1
L6:
	;
	if v28 < int32(2) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v24 = F_valkey_calloc(m, int32(200))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v28 = v27
	goto L6
L9:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, _c_F_replconfCommand[1]))
	F_addReply(m, l0, v1520)
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L4
	} else {
		goto L393
	}
L10:
	;
	v35 = int32(1)
	goto L11
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v44 = v35 << (uint(int32(2)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44)))
	v47 = F_objectGetVal(m, v46)
	mBase = m.M
	v48 = int32(_a_F_replconfCommand_0)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v51 != 0 {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	goto L9
L13:
	;
	v1506 = v35 + int32(2)
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1506 < v1507 {
		v35 = v1506
		goto L11
	} else {
		goto L392
	}
L14:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v98+v44)))
	v101 = F_objectGetVal(m, v100)
	mBase = m.M
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v103 = v102 + v44
	v104 = int32(_a_F_replconfCommand_1)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v107 != 0 {
		goto L33
	} else {
		goto L34
	}
L15:
	;
	if v83-v85 != 0 {
		goto L14
	} else {
		goto L27
	}
L16:
	;
	v83 = F_tolower(m, v79)
	mBase = m.M
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v85 = F_tolower(m, v84)
	mBase = m.M
	goto L15
L17:
	;
	v53 = v47
	v54 = v48
	v55 = v51
	goto L20
L18:
	;
	v79 = int32(0)
	v80 = v48
	goto L16
L19:
	;
	v79 = v76 & int32(255)
	v80 = v75
	goto L16
L20:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v57 == int32(0) {
		v75 = v54
		v76 = v55
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v75 = v69
	v76 = int32(0)
	goto L19
L22:
	;
	v61 = v55 & int32(255)
	if v61 == v57 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v68 = int32(1)
	v69 = v54 + v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	if v70 != 0 {
		v53 = v53 + v68
		v54 = v69
		v55 = v70
		goto L20
	} else {
		goto L26
	}
L24:
	;
	v63 = F_tolower(m, v61)
	mBase = m.M
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v65 = F_tolower(m, v64)
	mBase = m.M
	if v63 == v65 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v75 = v54
	v76 = v67
	goto L19
L26:
	;
	goto L21
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v44)+4))
	v93 = F_getLongFromObjectOrReply(m, l0, v89, v13+int32(104), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	if v93 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+148)) = v96
	goto L13
L30:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v192 = F_objectGetVal(m, v191)
	mBase = m.M
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v194 = int32(_a_F_replconfCommand_2)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v197 != 0 {
		goto L63
	} else {
		goto L64
	}
L31:
	;
	if v139-v141 != 0 {
		goto L30
	} else {
		goto L43
	}
L32:
	;
	v139 = F_tolower(m, v135)
	mBase = m.M
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v141 = F_tolower(m, v140)
	mBase = m.M
	goto L31
L33:
	;
	v109 = v101
	v110 = v104
	v111 = v107
	goto L36
L34:
	;
	v135 = int32(0)
	v136 = v104
	goto L32
L35:
	;
	v135 = v132 & int32(255)
	v136 = v131
	goto L32
L36:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v113 == int32(0) {
		v131 = v110
		v132 = v111
		goto L35
	} else {
		goto L38
	}
L37:
	;
	v131 = v125
	v132 = int32(0)
	goto L35
L38:
	;
	v117 = v111 & int32(255)
	if v117 == v113 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v124 = int32(1)
	v125 = v110 + v124
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
	if v126 != 0 {
		v109 = v109 + v124
		v110 = v125
		v111 = v126
		goto L36
	} else {
		goto L42
	}
L40:
	;
	v119 = F_tolower(m, v117)
	mBase = m.M
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v121 = F_tolower(m, v120)
	mBase = m.M
	if v119 == v121 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v131 = v110
	v132 = v123
	goto L35
L42:
	;
	goto L37
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v144 = F_objectGetVal(m, v143)
	mBase = m.M
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+int32(-1)))))
	v149 = v147 & int32(7)
	switch v149 + int32(-2) {
	case 0:
		goto L49
	case 1:
		goto L48
	case 2:
		goto L47
	default:
		goto L45
	}
L44:
	;
	switch v149 + int32(-2) {
	default:
		goto L58
	case 1:
		goto L57
	case 2:
		goto L56
	}
L45:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+152))
	if v166 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	if base.Ui32(int32(255)) < base.Ui32(v161) {
		goto L44
	} else {
		goto L50
	}
L47:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-17))))
	v161 = v160
	goto L46
L48:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-9))))
	v161 = v157
	goto L46
L49:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+int32(-5)))))
	v161 = v154
	goto L46
L50:
	;
	goto L45
L51:
	;
	v171 = F_sdsdup(m, v144)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	F_sdsfree(m, v166)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+152)) = v171
	goto L13
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v186
	F_addReplyErrorFormat(m, l0, int32(_a_F_replconfCommand_3), v13)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L59
	}
L56:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-17))))
	v186 = v185
	goto L55
L57:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-9))))
	v186 = v182
	goto L55
L58:
	;
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+int32(-5)))))
	v186 = v179
	goto L55
L59:
	;
	goto L1
L60:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v193+v44)))
	v447 = F_objectGetVal(m, v446)
	mBase = m.M
	v448 = int32(_a_F_replconfCommand_4)
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447))))
	if v451 != 0 {
		goto L137
	} else {
		goto L138
	}
L61:
	;
	if v229-v231 != 0 {
		goto L60
	} else {
		goto L73
	}
L62:
	;
	v229 = F_tolower(m, v225)
	mBase = m.M
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	v231 = F_tolower(m, v230)
	mBase = m.M
	goto L61
L63:
	;
	v199 = v192
	v200 = v194
	v201 = v197
	goto L66
L64:
	;
	v225 = int32(0)
	v226 = v194
	goto L62
L65:
	;
	v225 = v222 & int32(255)
	v226 = v221
	goto L62
L66:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v203 == int32(0) {
		v221 = v200
		v222 = v201
		goto L65
	} else {
		goto L68
	}
L67:
	;
	v221 = v215
	v222 = int32(0)
	goto L65
L68:
	;
	v207 = v201 & int32(255)
	if v207 == v203 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v214 = int32(1)
	v215 = v200 + v214
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+1)))
	if v216 != 0 {
		v199 = v199 + v214
		v200 = v215
		v201 = v216
		goto L66
	} else {
		goto L72
	}
L70:
	;
	v209 = F_tolower(m, v207)
	mBase = m.M
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	v211 = F_tolower(m, v210)
	mBase = m.M
	if v209 == v211 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	v221 = v200
	v222 = v213
	goto L65
L72:
	;
	goto L67
L73:
	;
	v236 = (v35 + int32(1)) << (uint(int32(2)) % 32)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v193+v236)))
	v239 = F_objectGetVal(m, v238)
	mBase = m.M
	v240 = int32(_a_F_replconfCommand_5)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	if v243 != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v284+v236)))
	v287 = F_objectGetVal(m, v286)
	mBase = m.M
	v288 = int32(_a_F_replconfCommand_6)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	if v291 != 0 {
		goto L91
	} else {
		goto L92
	}
L75:
	;
	if v275-v277 != 0 {
		goto L74
	} else {
		goto L87
	}
L76:
	;
	v275 = F_tolower(m, v271)
	mBase = m.M
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	v277 = F_tolower(m, v276)
	mBase = m.M
	goto L75
L77:
	;
	v245 = v239
	v246 = v240
	v247 = v243
	goto L80
L78:
	;
	v271 = int32(0)
	v272 = v240
	goto L76
L79:
	;
	v271 = v268 & int32(255)
	v272 = v267
	goto L76
L80:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	if v249 == int32(0) {
		v267 = v246
		v268 = v247
		goto L79
	} else {
		goto L82
	}
L81:
	;
	v267 = v261
	v268 = int32(0)
	goto L79
L82:
	;
	v253 = v247 & int32(255)
	if v253 == v249 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v260 = int32(1)
	v261 = v246 + v260
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+1)))
	if v262 != 0 {
		v245 = v245 + v260
		v246 = v261
		v247 = v262
		goto L80
	} else {
		goto L86
	}
L84:
	;
	v255 = F_tolower(m, v253)
	mBase = m.M
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	v257 = F_tolower(m, v256)
	mBase = m.M
	if v255 == v257 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	v267 = v246
	v268 = v259
	goto L79
L86:
	;
	goto L81
L87:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279)+160)))
	v282 = v280 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v279)+160)) = uint16(v282)
	goto L13
L88:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v332+v236)))
	v335 = F_objectGetVal(m, v334)
	mBase = m.M
	v336 = int32(_a_F_replconfCommand_7)
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	if v339 != 0 {
		goto L105
	} else {
		goto L106
	}
L89:
	;
	if v323-v325 != 0 {
		goto L88
	} else {
		goto L101
	}
L90:
	;
	v323 = F_tolower(m, v319)
	mBase = m.M
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320))))
	v325 = F_tolower(m, v324)
	mBase = m.M
	goto L89
L91:
	;
	v293 = v287
	v294 = v288
	v295 = v291
	goto L94
L92:
	;
	v319 = int32(0)
	v320 = v288
	goto L90
L93:
	;
	v319 = v316 & int32(255)
	v320 = v315
	goto L90
L94:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v297 == int32(0) {
		v315 = v294
		v316 = v295
		goto L93
	} else {
		goto L96
	}
L95:
	;
	v315 = v309
	v316 = int32(0)
	goto L93
L96:
	;
	v301 = v295 & int32(255)
	if v301 == v297 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v308 = int32(1)
	v309 = v294 + v308
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+1)))
	if v310 != 0 {
		v293 = v293 + v308
		v294 = v309
		v295 = v310
		goto L94
	} else {
		goto L100
	}
L98:
	;
	v303 = F_tolower(m, v301)
	mBase = m.M
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	v305 = F_tolower(m, v304)
	mBase = m.M
	if v303 == v305 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	v315 = v294
	v316 = v307
	goto L93
L100:
	;
	goto L95
L101:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v327)+160)))
	v330 = v328 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v327)+160)) = uint16(v330)
	goto L13
L102:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v397+v236)))
	v400 = F_objectGetVal(m, v399)
	mBase = m.M
	v401 = int32(_a_F_replconfCommand_8)
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	if v404 != 0 {
		goto L123
	} else {
		goto L124
	}
L103:
	;
	if v371-v373 != 0 {
		goto L102
	} else {
		goto L115
	}
L104:
	;
	v371 = F_tolower(m, v367)
	mBase = m.M
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	v373 = F_tolower(m, v372)
	mBase = m.M
	goto L103
L105:
	;
	v341 = v335
	v342 = v336
	v343 = v339
	goto L108
L106:
	;
	v367 = int32(0)
	v368 = v336
	goto L104
L107:
	;
	v367 = v364 & int32(255)
	v368 = v363
	goto L104
L108:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	if v345 == int32(0) {
		v363 = v342
		v364 = v343
		goto L107
	} else {
		goto L110
	}
L109:
	;
	v363 = v357
	v364 = int32(0)
	goto L107
L110:
	;
	v349 = v343 & int32(255)
	if v349 == v345 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v356 = int32(1)
	v357 = v342 + v356
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+1)))
	if v358 != 0 {
		v341 = v341 + v356
		v342 = v357
		v343 = v358
		goto L108
	} else {
		goto L114
	}
L112:
	;
	v351 = F_tolower(m, v349)
	mBase = m.M
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	v353 = F_tolower(m, v352)
	mBase = m.M
	if v351 == v353 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	v363 = v342
	v364 = v355
	goto L107
L114:
	;
	goto L109
L115:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_replconfCommand[2]))
	if v376 == int32(0) {
		goto L102
	} else {
		goto L116
	}
L116:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_replconfCommand[3]))
	if v380 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_replconfCommand[4]))
	if int32(2) < v389 {
		goto L13
	} else {
		goto L119
	}
L118:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v383)+160)))
	v386 = v384 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v383)+160)) = uint16(v386)
	goto L13
L119:
	;
	F__serverLog(m, int32(2), int32(_a_F_replconfCommand_9), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	goto L13
L121:
	;
	if v436-v438 != 0 {
		goto L13
	} else {
		goto L133
	}
L122:
	;
	v436 = F_tolower(m, v432)
	mBase = m.M
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433))))
	v438 = F_tolower(m, v437)
	mBase = m.M
	goto L121
L123:
	;
	v406 = v400
	v407 = v401
	v408 = v404
	goto L126
L124:
	;
	v432 = int32(0)
	v433 = v401
	goto L122
L125:
	;
	v432 = v429 & int32(255)
	v433 = v428
	goto L122
L126:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407))))
	if v410 == int32(0) {
		v428 = v407
		v429 = v408
		goto L125
	} else {
		goto L128
	}
L127:
	;
	v428 = v422
	v429 = int32(0)
	goto L125
L128:
	;
	v414 = v408 & int32(255)
	if v414 == v410 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v421 = int32(1)
	v422 = v407 + v421
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+1)))
	if v423 != 0 {
		v406 = v406 + v421
		v407 = v422
		v408 = v423
		goto L126
	} else {
		goto L132
	}
L130:
	;
	v416 = F_tolower(m, v414)
	mBase = m.M
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407))))
	v418 = F_tolower(m, v417)
	mBase = m.M
	if v416 == v418 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	v428 = v407
	v429 = v420
	goto L125
L132:
	;
	goto L127
L133:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v440)+160)))
	v443 = v441 | int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v440)+160)) = uint16(v443)
	goto L13
L134:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v602+v44)))
	v605 = F_objectGetVal(m, v604)
	mBase = m.M
	v606 = int32(_a_F_replconfCommand_10)
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605))))
	if v609 != 0 {
		goto L185
	} else {
		goto L186
	}
L135:
	;
	if v483-v485 != 0 {
		goto L134
	} else {
		goto L147
	}
L136:
	;
	v483 = F_tolower(m, v479)
	mBase = m.M
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	v485 = F_tolower(m, v484)
	mBase = m.M
	goto L135
L137:
	;
	v453 = v447
	v454 = v448
	v455 = v451
	goto L140
L138:
	;
	v479 = int32(0)
	v480 = v448
	goto L136
L139:
	;
	v479 = v476 & int32(255)
	v480 = v475
	goto L136
L140:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454))))
	if v457 == int32(0) {
		v475 = v454
		v476 = v455
		goto L139
	} else {
		goto L142
	}
L141:
	;
	v475 = v469
	v476 = int32(0)
	goto L139
L142:
	;
	v461 = v455 & int32(255)
	if v461 == v457 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v468 = int32(1)
	v469 = v454 + v468
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+1)))
	if v470 != 0 {
		v453 = v453 + v468
		v454 = v469
		v455 = v470
		goto L140
	} else {
		goto L146
	}
L144:
	;
	v463 = F_tolower(m, v461)
	mBase = m.M
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454))))
	v465 = F_tolower(m, v464)
	mBase = m.M
	if v463 == v465 {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453))))
	v475 = v454
	v476 = v467
	goto L139
L146:
	;
	goto L141
L147:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v487&int32(2) == int32(0) {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v492+v35<<(uint(int32(2))%32))+4))
	v499 = F_getLongLongFromObject(m, v496, v13+int32(104))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	if v499 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v501 = *(*int64)(unsafe.Add(mBase, uint32(v13)+104))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v503 = *(*int64)(unsafe.Add(mBase, uint32(v502)+64))
	if v501 <= v503 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v508 = v35 + int32(3)
	if v506 <= v508 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v502)+64)) = v501
	goto L151
L153:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v572 = int32(_a_F_replconfCommand_11)
	v573 = *(*int64)(unsafe.Add(mBase, _c_F_replconfCommand[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v571)+80)) = v573
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_replconfCommand[6]))
	if v576 != int32(1) {
		v585 = v571
		goto L171
	} else {
		goto L172
	}
L154:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v510+v35<<(uint(int32(2))%32))+8))
	v515 = F_objectGetVal(m, v514)
	mBase = m.M
	v516 = int32(_a_F_replconfCommand_12)
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
	if v519 != 0 {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	if v551-v553 != 0 {
		goto L153
	} else {
		goto L167
	}
L156:
	;
	v551 = F_tolower(m, v547)
	mBase = m.M
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548))))
	v553 = F_tolower(m, v552)
	mBase = m.M
	goto L155
L157:
	;
	v521 = v515
	v522 = v516
	v523 = v519
	goto L160
L158:
	;
	v547 = int32(0)
	v548 = v516
	goto L156
L159:
	;
	v547 = v544 & int32(255)
	v548 = v543
	goto L156
L160:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	if v525 == int32(0) {
		v543 = v522
		v544 = v523
		goto L159
	} else {
		goto L162
	}
L161:
	;
	v543 = v537
	v544 = int32(0)
	goto L159
L162:
	;
	v529 = v523 & int32(255)
	if v529 == v525 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v536 = int32(1)
	v537 = v522 + v536
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521)+1)))
	if v538 != 0 {
		v521 = v521 + v536
		v522 = v537
		v523 = v538
		goto L160
	} else {
		goto L166
	}
L164:
	;
	v531 = F_tolower(m, v529)
	mBase = m.M
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	v533 = F_tolower(m, v532)
	mBase = m.M
	if v531 == v533 {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521))))
	v543 = v522
	v544 = v535
	goto L159
L166:
	;
	goto L161
L167:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v555+v508<<(uint(int32(2))%32))))
	v562 = F_getLongLongFromObject(m, v559, v13+int32(104))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	if v562 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v564 = *(*int64)(unsafe.Add(mBase, uint32(v13)+104))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v565)+72))
	if v564 <= v566 {
		goto L153
	} else {
		goto L170
	}
L170:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v565)+72)) = v564
	goto L153
L171:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v585)+4))
	if v586 == int32(0) {
		v595 = v585
		goto L176
	} else {
		goto L177
	}
L172:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v571)))
	if v579 != int32(7) {
		v585 = v571
		goto L171
	} else {
		goto L173
	}
L173:
	;
	F_checkChildrenDone(m)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v585 = v584
	goto L171
L175:
	;
	if v597 != int32(11) {
		goto L1
	} else {
		goto L180
	}
L176:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v595)))
	v597 = v596
	goto L175
L177:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v585)))
	if v589 != int32(9) {
		v597 = v589
		goto L175
	} else {
		goto L178
	}
L178:
	;
	F_replicaStartCommandStream(m, l0)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v595 = v594
	goto L176
L180:
	;
	v600 = F_replicaPutOnline(m, l0)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	goto L1
L182:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v655+v44)))
	v658 = F_objectGetVal(m, v657)
	mBase = m.M
	v659 = int32(_a_F_replconfCommand_13)
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v658))))
	if v662 != 0 {
		goto L202
	} else {
		goto L203
	}
L183:
	;
	if v641-v643 != 0 {
		goto L182
	} else {
		goto L195
	}
L184:
	;
	v641 = F_tolower(m, v637)
	mBase = m.M
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638))))
	v643 = F_tolower(m, v642)
	mBase = m.M
	goto L183
L185:
	;
	v611 = v605
	v612 = v606
	v613 = v609
	goto L188
L186:
	;
	v637 = int32(0)
	v638 = v606
	goto L184
L187:
	;
	v637 = v634 & int32(255)
	v638 = v633
	goto L184
L188:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612))))
	if v615 == int32(0) {
		v633 = v612
		v634 = v613
		goto L187
	} else {
		goto L190
	}
L189:
	;
	v633 = v627
	v634 = int32(0)
	goto L187
L190:
	;
	v619 = v613 & int32(255)
	if v619 == v615 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v626 = int32(1)
	v627 = v612 + v626
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+1)))
	if v628 != 0 {
		v611 = v611 + v626
		v612 = v627
		v613 = v628
		goto L188
	} else {
		goto L194
	}
L192:
	;
	v621 = F_tolower(m, v619)
	mBase = m.M
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612))))
	v623 = F_tolower(m, v622)
	mBase = m.M
	if v621 == v623 {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
	v633 = v612
	v634 = v625
	goto L187
L194:
	;
	goto L189
L195:
	;
	v646 = *(*int32)(unsafe.Add(mBase, _c_F_replconfCommand[7]))
	if v646 == int32(0) {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v650 = *(*int32)(unsafe.Add(mBase, _c_F_replconfCommand[8]))
	if v650 == int32(0) {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_replicationSendAck(m)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L4
	} else {
		goto L198
	}
L198:
	;
	goto L1
L199:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v720+v44)))
	v723 = F_objectGetVal(m, v722)
	mBase = m.M
	v724 = int32(_a_F_replconfCommand_14)
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723))))
	if v727 != 0 {
		goto L220
	} else {
		goto L221
	}
L200:
	;
	if v694-v696 != 0 {
		goto L199
	} else {
		goto L212
	}
L201:
	;
	v694 = F_tolower(m, v690)
	mBase = m.M
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691))))
	v696 = F_tolower(m, v695)
	mBase = m.M
	goto L200
L202:
	;
	v664 = v658
	v665 = v659
	v666 = v662
	goto L205
L203:
	;
	v690 = int32(0)
	v691 = v659
	goto L201
L204:
	;
	v690 = v687 & int32(255)
	v691 = v686
	goto L201
L205:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665))))
	if v668 == int32(0) {
		v686 = v665
		v687 = v666
		goto L204
	} else {
		goto L207
	}
L206:
	;
	v686 = v680
	v687 = int32(0)
	goto L204
L207:
	;
	v672 = v666 & int32(255)
	if v672 == v668 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v679 = int32(1)
	v680 = v665 + v679
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664)+1)))
	if v681 != 0 {
		v664 = v664 + v679
		v665 = v680
		v666 = v681
		goto L205
	} else {
		goto L211
	}
L209:
	;
	v674 = F_tolower(m, v672)
	mBase = m.M
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665))))
	v676 = F_tolower(m, v675)
	mBase = m.M
	if v674 == v676 {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	v686 = v665
	v687 = v678
	goto L204
L211:
	;
	goto L206
L212:
	;
	v698 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v698
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v700+v44)+4))
	v708 = F_getRangeLongFromObjectOrReply(m, l0, v702, v698, int32(1), v13+int32(104), v698)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	if v708 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v710&int32(-8193) | base.B2i32(v713 == int32(1))<<(uint(int32(13))%32)
	goto L13
L215:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v1493)+168)) = v1347
	goto L13
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v798
	F_addReplyErrorFormat(m, l0, int32(_a_F_replconfCommand_15), v13+int32(16))
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L4
	} else {
		goto L390
	}
L217:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v854+v44)))
	v857 = F_objectGetVal(m, v856)
	mBase = m.M
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v859 = int32(_a_F_replconfCommand_16)
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857))))
	if v862 != 0 {
		goto L257
	} else {
		goto L258
	}
L218:
	;
	if v759-v761 != 0 {
		goto L217
	} else {
		goto L230
	}
L219:
	;
	v759 = F_tolower(m, v755)
	mBase = m.M
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756))))
	v761 = F_tolower(m, v760)
	mBase = m.M
	goto L218
L220:
	;
	v729 = v723
	v730 = v724
	v731 = v727
	goto L223
L221:
	;
	v755 = int32(0)
	v756 = v724
	goto L219
L222:
	;
	v755 = v752 & int32(255)
	v756 = v751
	goto L219
L223:
	;
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730))))
	if v733 == int32(0) {
		v751 = v730
		v752 = v731
		goto L222
	} else {
		goto L225
	}
L224:
	;
	v751 = v745
	v752 = int32(0)
	goto L222
L225:
	;
	v737 = v731 & int32(255)
	if v737 == v733 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v744 = int32(1)
	v745 = v730 + v744
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v729)+1)))
	if v746 != 0 {
		v729 = v729 + v744
		v730 = v745
		v731 = v746
		goto L223
	} else {
		goto L229
	}
L227:
	;
	v739 = F_tolower(m, v737)
	mBase = m.M
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730))))
	v741 = F_tolower(m, v740)
	mBase = m.M
	if v739 == v741 {
		goto L226
	} else {
		goto L228
	}
L228:
	;
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v729))))
	v751 = v730
	v752 = v743
	goto L222
L229:
	;
	goto L224
L230:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v763+v44)+4))
	v766 = F_objectGetVal(m, v765)
	mBase = m.M
	v769 = F_sdssplitargs(m, v766, v13+int32(104))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L4
	} else {
		goto L232
	}
L231:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v775 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v774)+162)))
	v777 = v775 | int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v774)+162)) = uint16(v777)
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	if v779 < int32(1) {
		goto L235
	} else {
		goto L236
	}
L232:
	;
	if v769 != 0 {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	F_addReplyError(m, l0, int32(_a_F_replconfCommand_17))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L4
	} else {
		goto L234
	}
L234:
	;
	goto L1
L235:
	;
	F_sdsfreesplitres(m, v769, v779)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L4
	} else {
		goto L253
	}
L236:
	;
	v783 = v777 & int32(65533)
	v787 = int32(0)
	goto L237
L237:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v769+v787<<(uint(int32(2))%32))))
	v799 = int32(_a_F_replconfCommand_18)
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
	if v802 != 0 {
		goto L241
	} else {
		goto L242
	}
L238:
	;
	goto L235
L239:
	;
	if v834-v836 != 0 {
		goto L216
	} else {
		goto L251
	}
L240:
	;
	v834 = F_tolower(m, v830)
	mBase = m.M
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v831))))
	v836 = F_tolower(m, v835)
	mBase = m.M
	goto L239
L241:
	;
	v804 = v798
	v805 = v799
	v806 = v802
	goto L244
L242:
	;
	v830 = int32(0)
	v831 = v799
	goto L240
L243:
	;
	v830 = v827 & int32(255)
	v831 = v826
	goto L240
L244:
	;
	v808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v805))))
	if v808 == int32(0) {
		v826 = v805
		v827 = v806
		goto L243
	} else {
		goto L246
	}
L245:
	;
	v826 = v820
	v827 = int32(0)
	goto L243
L246:
	;
	v812 = v806 & int32(255)
	if v812 == v808 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v819 = int32(1)
	v820 = v805 + v819
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804)+1)))
	if v821 != 0 {
		v804 = v804 + v819
		v805 = v820
		v806 = v821
		goto L244
	} else {
		goto L250
	}
L248:
	;
	v814 = F_tolower(m, v812)
	mBase = m.M
	v815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v805))))
	v816 = F_tolower(m, v815)
	mBase = m.M
	if v814 == v816 {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804))))
	v826 = v805
	v827 = v818
	goto L243
L250:
	;
	goto L245
L251:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v774)+162)) = uint16(v783)
	v840 = v787 + int32(1)
	if v840 != v779 {
		v787 = v840
		goto L237
	} else {
		goto L252
	}
L252:
	;
	goto L238
L253:
	;
	goto L13
L254:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v858+v44)))
	v989 = F_objectGetVal(m, v988)
	mBase = m.M
	v990 = int32(_a_F_replconfCommand_19)
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v989))))
	if v993 != 0 {
		goto L288
	} else {
		goto L289
	}
L255:
	;
	if v894-v896 != 0 {
		goto L254
	} else {
		goto L267
	}
L256:
	;
	v894 = F_tolower(m, v890)
	mBase = m.M
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891))))
	v896 = F_tolower(m, v895)
	mBase = m.M
	goto L255
L257:
	;
	v864 = v857
	v865 = v859
	v866 = v862
	goto L260
L258:
	;
	v890 = int32(0)
	v891 = v859
	goto L256
L259:
	;
	v890 = v887 & int32(255)
	v891 = v886
	goto L256
L260:
	;
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v865))))
	if v868 == int32(0) {
		v886 = v865
		v887 = v866
		goto L259
	} else {
		goto L262
	}
L261:
	;
	v886 = v880
	v887 = int32(0)
	goto L259
L262:
	;
	v872 = v866 & int32(255)
	if v872 == v868 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v879 = int32(1)
	v880 = v865 + v879
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864)+1)))
	if v881 != 0 {
		v864 = v864 + v879
		v865 = v880
		v866 = v881
		goto L260
	} else {
		goto L266
	}
L264:
	;
	v874 = F_tolower(m, v872)
	mBase = m.M
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v865))))
	v876 = F_tolower(m, v875)
	mBase = m.M
	if v874 == v876 {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864))))
	v886 = v865
	v887 = v878
	goto L259
L266:
	;
	goto L261
L267:
	;
	v899 = v35 + int32(1)
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v858+v899<<(uint(int32(2))%32))))
	v904 = F_objectGetVal(m, v903)
	mBase = m.M
	v905 = int32(0)
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904))))
	v914 = v904
	v915 = v910
	v916 = v905
	v917 = v905
	v918 = v905
	goto L271
L268:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v975+v899<<(uint(int32(2))%32))))
	v980 = F_objectGetVal(m, v979)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v980
	F_addReplyErrorFormat(m, l0, int32(_a_F_replconfCommand_20), v13+int32(32))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L4
	} else {
		goto L283
	}
L269:
	;
	if v970 < int32(0) {
		goto L268
	} else {
		goto L282
	}
L270:
	;
	v970 = v963
	goto L269
L271:
	;
	v923 = (v915 + int32(-48)) & int32(255)
	if base.Ui32(int32(9)) < base.Ui32(v923) {
		goto L274
	} else {
		goto L275
	}
L272:
	;
	if v947 == int32(2) {
		goto L280
	} else {
		goto L281
	}
L273:
	;
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914)+1)))
	if v949 != 0 {
		v914 = v914 + int32(1)
		v915 = v949
		v916 = v945
		v917 = v946
		v918 = v947
		goto L271
	} else {
		goto L279
	}
L274:
	;
	v932 = int32(-1)
	if v915&int32(255) != int32(46) {
		v963 = v932
		goto L270
	} else {
		goto L277
	}
L275:
	;
	v928 = v917*int32(10) + v923
	if v928 <= int32(255) {
		v945 = v916
		v946 = v928
		v947 = v918
		goto L273
	} else {
		goto L276
	}
L276:
	;
	v970 = int32(-1)
	goto L269
L277:
	;
	if int32(1) < v918 {
		v963 = v932
		goto L270
	} else {
		goto L278
	}
L278:
	;
	v945 = v916<<(uint(int32(8))%32) | v917
	v946 = int32(0)
	v947 = v918 + int32(1)
	goto L273
L279:
	;
	goto L272
L280:
	;
	v963 = v945<<(uint(int32(8))%32) | v946
	goto L270
L281:
	;
	v970 = int32(-1)
	goto L269
L282:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v973)+156)) = v970
	goto L13
L283:
	;
	goto L1
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v1041 & int32(-67108865)
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v1480 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1479)+162)))
	v1482 = v1480 & int32(65531)
	*(*uint16)(unsafe.Add(mBase, uint32(v1479)+162)) = uint16(v1482)
	goto L13
L285:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1053+v44)))
	v1056 = F_objectGetVal(m, v1055)
	mBase = m.M
	v1057 = int32(_a_F_replconfCommand_21)
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056))))
	if v1060 != 0 {
		goto L305
	} else {
		goto L306
	}
L286:
	;
	if v1025-v1027 != 0 {
		goto L285
	} else {
		goto L298
	}
L287:
	;
	v1025 = F_tolower(m, v1021)
	mBase = m.M
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1022))))
	v1027 = F_tolower(m, v1026)
	mBase = m.M
	goto L286
L288:
	;
	v995 = v989
	v996 = v990
	v997 = v993
	goto L291
L289:
	;
	v1021 = int32(0)
	v1022 = v990
	goto L287
L290:
	;
	v1021 = v1018 & int32(255)
	v1022 = v1017
	goto L287
L291:
	;
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996))))
	if v999 == int32(0) {
		v1017 = v996
		v1018 = v997
		goto L290
	} else {
		goto L293
	}
L292:
	;
	v1017 = v1011
	v1018 = int32(0)
	goto L290
L293:
	;
	v1003 = v997 & int32(255)
	if v1003 == v999 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1010 = int32(1)
	v1011 = v996 + v1010
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+1)))
	if v1012 != 0 {
		v995 = v995 + v1010
		v996 = v1011
		v997 = v1012
		goto L291
	} else {
		goto L297
	}
L295:
	;
	v1005 = F_tolower(m, v1003)
	mBase = m.M
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996))))
	v1007 = F_tolower(m, v1006)
	mBase = m.M
	if v1005 == v1007 {
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995))))
	v1017 = v996
	v1018 = v1009
	goto L290
L297:
	;
	goto L292
L298:
	;
	v1029 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v1029
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1031+v44)+4))
	v1039 = F_getRangeLongFromObjectOrReply(m, l0, v1033, v1029, int32(1), v13+int32(104), v1029)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L4
	} else {
		goto L299
	}
L299:
	;
	if v1039 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	if v1042 != int32(1) {
		goto L284
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v1041 | int32(67108864)
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v1049 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1048)+162)))
	v1051 = v1049 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1048)+162)) = uint16(v1051)
	goto L13
L302:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1355+v44)))
	v1358 = F_objectGetVal(m, v1357)
	mBase = m.M
	v1359 = int32(_a_F_replconfCommand_22)
	v1362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1358))))
	if v1362 != 0 {
		goto L361
	} else {
		goto L362
	}
L303:
	;
	if v1092-v1094 != 0 {
		goto L302
	} else {
		goto L315
	}
L304:
	;
	v1092 = F_tolower(m, v1088)
	mBase = m.M
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1089))))
	v1094 = F_tolower(m, v1093)
	mBase = m.M
	goto L303
L305:
	;
	v1062 = v1056
	v1063 = v1057
	v1064 = v1060
	goto L308
L306:
	;
	v1088 = int32(0)
	v1089 = v1057
	goto L304
L307:
	;
	v1088 = v1085 & int32(255)
	v1089 = v1084
	goto L304
L308:
	;
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1063))))
	if v1066 == int32(0) {
		v1084 = v1063
		v1085 = v1064
		goto L307
	} else {
		goto L310
	}
L309:
	;
	v1084 = v1078
	v1085 = int32(0)
	goto L307
L310:
	;
	v1070 = v1064 & int32(255)
	if v1070 == v1066 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1077 = int32(1)
	v1078 = v1063 + v1077
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062)+1)))
	if v1079 != 0 {
		v1062 = v1062 + v1077
		v1063 = v1078
		v1064 = v1079
		goto L308
	} else {
		goto L314
	}
L312:
	;
	v1072 = F_tolower(m, v1070)
	mBase = m.M
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1063))))
	v1074 = F_tolower(m, v1073)
	mBase = m.M
	if v1072 == v1074 {
		goto L311
	} else {
		goto L313
	}
L313:
	;
	v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062))))
	v1084 = v1063
	v1085 = v1076
	goto L307
L314:
	;
	goto L309
L315:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = int64(0)
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1098+v44)+4))
	v1104 = F_getLongLongFromObjectOrReply(m, l0, v1100, v13+int32(88), int32(0))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L4
	} else {
		goto L316
	}
L316:
	;
	if v1104 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	v1106 = *(*int64)(unsafe.Add(mBase, uint32(v13)+88))
	v1107 = int64(56)
	v1109 = int64(65280)
	v1111 = int64(40)
	v1114 = int64(16711680)
	v1116 = int64(24)
	v1118 = int64(4278190080)
	v1120 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+104)) = v1106<<(uint(v1107)%64) | v1106&v1109<<(uint(v1111)%64) | (v1106&v1114<<(uint(v1116)%64) | v1106&v1118<<(uint(v1120)%64)) | (int64(base.Ui64(v1106)>>(uint(v1120)%64))&v1118 | int64(base.Ui64(v1106)>>(uint(v1116)%64))&v1114 | (int64(base.Ui64(v1106)>>(uint(v1111)%64))&v1109 | int64(base.Ui64(v1106)>>(uint(v1107)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = int32(0)
	v1146 = *(*int32)(unsafe.Add(mBase, _c_F_replconfCommand[9]))
	v1148 = v13 + int32(104)
	v1149 = int32(8)
	v1151 = v13 + int32(100)
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1146)))
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1160)))
	goto L321
L318:
	;
	v1347 = *(*int64)(unsafe.Add(mBase, uint32(v13)+88))
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
	if v1348 != 0 {
		goto L215
	} else {
		goto L356
	}
L319:
	;
	if v1304 != v1149 {
		goto L346
	} else {
		goto L347
	}
L320:
	;
	v1295 = int32(0)
	v1301 = v1160
	v1302 = v1161
	v1304 = v1295
	v1308 = v1295
	goto L319
L321:
	;
	if base.Ui32(v1161) < base.Ui32(int32(8)) {
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1172 = v1160
	v1173 = v1161
	v1175 = int32(0)
	goto L324
L323:
	;
	v1301 = v1285
	v1302 = v1286
	v1304 = v1288
	v1308 = base.B2i32(v1291 != int32(0))
	goto L319
L324:
	;
	v1181 = int32(base.Ui32(v1173) >> (uint(int32(3)) % 32))
	v1182 = int32(4)
	v1183 = v1172 + v1182
	if v1173&v1182 == int32(0) {
		goto L327
	} else {
		goto L328
	}
L325:
	;
	v1285 = v1276
	v1286 = v1277
	v1288 = v1261
	v1291 = v1266
	goto L323
L326:
	;
	v1266 = int32(0)
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1183+v1181+(v1266-v1181)&int32(3)+v1254<<(uint(int32(2))%32))))
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1276)))
	if base.Ui32(v1277) < base.Ui32(int32(8)) {
		v1285 = v1276
		v1286 = v1277
		v1288 = v1261
		v1291 = v1266
		goto L323
	} else {
		goto L344
	}
L327:
	;
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1148+v1175))))
	v1232 = int32(0)
	goto L338
L328:
	;
	v1188 = int32(0)
	if base.Ui32(v1149) <= base.Ui32(v1175) {
		v1221 = v1175
		v1224 = v1188
		goto L329
	} else {
		goto L330
	}
L329:
	;
	if v1224 == v1181 {
		v1254 = v1188
		v1261 = v1221
		goto L326
	} else {
		goto L336
	}
L330:
	;
	v1198 = v1175
	v1201 = v1188
	goto L331
L331:
	;
	v1204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1183+v1201))))
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1148+v1198))))
	if v1204 != v1206 {
		v1221 = v1198
		v1224 = v1201
		goto L329
	} else {
		goto L333
	}
L332:
	;
	v1221 = v1209
	v1224 = v1211
	goto L329
L333:
	;
	v1208 = int32(1)
	v1209 = v1198 + v1208
	v1211 = v1201 + v1208
	if base.Ui32(v1181) <= base.Ui32(v1211) {
		v1221 = v1209
		v1224 = v1211
		goto L329
	} else {
		goto L334
	}
L334:
	;
	if base.Ui32(v1209) < base.Ui32(v1149) {
		v1198 = v1209
		v1201 = v1211
		goto L331
	} else {
		goto L335
	}
L335:
	;
	goto L332
L336:
	;
	v1285 = v1172
	v1286 = v1173
	v1288 = v1221
	v1291 = v1224
	goto L323
L337:
	;
	if v1232 != v1181 {
		goto L342
	} else {
		goto L343
	}
L338:
	;
	v1245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1183+v1232))))
	if v1245 == v1229&int32(255) {
		goto L337
	} else {
		goto L340
	}
L340:
	;
	v1247 = int32(1)
	v1249 = v1232 + v1247
	if v1249 != v1181 {
		v1232 = v1249
		goto L338
	} else {
		goto L341
	}
L341:
	;
	v1301 = v1172
	v1302 = v1173
	v1304 = v1175
	v1308 = v1247
	goto L319
L342:
	;
	v1254 = v1232
	v1261 = v1175 + int32(1)
	goto L326
L343:
	;
	v1285 = v1172
	v1286 = v1173
	v1288 = v1175
	v1291 = v1181
	goto L323
L344:
	;
	if base.Ui32(v1261) < base.Ui32(v1149) {
		v1172 = v1276
		v1173 = v1277
		v1175 = v1261
		goto L324
	} else {
		goto L345
	}
L345:
	;
	goto L325
L346:
	;
	goto L318
L347:
	;
	if v1302&int32(1) == int32(0) {
		goto L346
	} else {
		goto L348
	}
L348:
	;
	v1316 = v1302 & int32(4)
	if v1308&base.B2i32(v1316 != int32(0)) != 0 {
		goto L346
	} else {
		goto L349
	}
L349:
	;
	if v1151 == int32(0) {
		goto L346
	} else {
		goto L350
	}
L350:
	;
	if v1302&int32(2) != 0 {
		v1342 = int32(0)
		goto L351
	} else {
		goto L352
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1151))) = v1342
	goto L346
L352:
	;
	v1326 = int32(3)
	v1327 = int32(base.Ui32(v1302) >> (uint(v1326) % 32))
	if v1316 != 0 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1337 = int32(4)
	goto L355
L354:
	;
	v1337 = v1327 << (uint(int32(2)) % 32)
	goto L355
L355:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1301+v1327+(int32(0)-v1327)&v1326+v1337+int32(4))))
	v1342 = v1341
	goto L351
L356:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v1347
	F_addReplyErrorFormat(m, l0, int32(_a_F_replconfCommand_23), v13+int32(48))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L4
	} else {
		goto L357
	}
L357:
	;
	goto L1
L358:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1464+v35<<(uint(int32(2))%32))))
	v1469 = F_objectGetVal(m, v1468)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v1469
	F_addReplyErrorFormat(m, l0, int32(_a_F_replconfCommand_24), v13+int32(80))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L4
	} else {
		goto L389
	}
L359:
	;
	if v1394-v1396 != 0 {
		goto L358
	} else {
		goto L371
	}
L360:
	;
	v1394 = F_tolower(m, v1390)
	mBase = m.M
	v1395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391))))
	v1396 = F_tolower(m, v1395)
	mBase = m.M
	goto L359
L361:
	;
	v1364 = v1358
	v1365 = v1359
	v1366 = v1362
	goto L364
L362:
	;
	v1390 = int32(0)
	v1391 = v1359
	goto L360
L363:
	;
	v1390 = v1387 & int32(255)
	v1391 = v1386
	goto L360
L364:
	;
	v1368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1365))))
	if v1368 == int32(0) {
		v1386 = v1365
		v1387 = v1366
		goto L363
	} else {
		goto L366
	}
L365:
	;
	v1386 = v1380
	v1387 = int32(0)
	goto L363
L366:
	;
	v1372 = v1366 & int32(255)
	if v1372 == v1368 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1379 = int32(1)
	v1380 = v1365 + v1379
	v1381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1364)+1)))
	if v1381 != 0 {
		v1364 = v1364 + v1379
		v1365 = v1380
		v1366 = v1381
		goto L364
	} else {
		goto L370
	}
L368:
	;
	v1374 = F_tolower(m, v1372)
	mBase = m.M
	v1375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1365))))
	v1376 = F_tolower(m, v1375)
	mBase = m.M
	if v1374 == v1376 {
		goto L367
	} else {
		goto L369
	}
L369:
	;
	v1378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1364))))
	v1386 = v1365
	v1387 = v1378
	goto L363
L370:
	;
	goto L365
L371:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, _c_F_replconfCommand[10]))
	if v1399 != 0 {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1405 = v35 + int32(1)
	v1407 = v1405 << (uint(int32(2)) % 32)
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1403+v1407)))
	v1410 = F_objectGetVal(m, v1409)
	mBase = m.M
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1412+v1407)))
	v1415 = F_objectGetVal(m, v1414)
	mBase = m.M
	v1418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1415+int32(-1)))))
	switch v1418 & int32(7) {
	case 0:
		goto L380
	case 1:
		goto L379
	case 2:
		goto L378
	case 3:
		goto L377
	case 4:
		goto L376
	default:
		v1435 = int32(0)
		goto L375
	}
L373:
	;
	F_addReplyError(m, l0, int32(_a_F_replconfCommand_25))
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L4
	} else {
		goto L374
	}
L374:
	;
	goto L1
L375:
	;
	v1436 = F_clusterLookupNode(m, v1410, v1435)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L4
	} else {
		goto L382
	}
L376:
	;
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1415+int32(-17))))
	v1435 = v1434
	goto L375
L377:
	;
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v1415+int32(-9))))
	v1435 = v1431
	goto L375
L378:
	;
	v1428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1415+int32(-5)))))
	v1435 = v1428
	goto L375
L379:
	;
	v1425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1415+int32(-3)))))
	v1435 = v1425
	goto L375
L380:
	;
	v1435 = int32(base.Ui32(v1418) >> (uint(int32(3)) % 32))
	goto L375
L381:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+192))
	if v1451 == int32(0) {
		goto L385
	} else {
		goto L386
	}
L382:
	;
	if v1436 != 0 {
		goto L381
	} else {
		goto L383
	}
L383:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1438+v1405<<(uint(int32(2))%32))))
	v1443 = F_objectGetVal(m, v1442)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v1443
	F_addReplyErrorFormat(m, l0, int32(_a_F_replconfCommand_26), v13+int32(64))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L4
	} else {
		goto L384
	}
L384:
	;
	goto L1
L385:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1456+v1407)))
	v1459 = F_objectGetVal(m, v1458)
	mBase = m.M
	v1460 = F_sdsdup(m, v1459)
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L4
	} else {
		goto L388
	}
L386:
	;
	F_sdsfree(m, v1451)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L4
	} else {
		goto L387
	}
L387:
	;
	goto L385
L388:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v1462)+192)) = v1460
	goto L13
L389:
	;
	goto L1
L390:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	F_sdsfreesplitres(m, v769, v1490)
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L4
	} else {
		goto L391
	}
L391:
	;
	goto L1
L392:
	;
	goto L12
L393:
	;
	goto L1
}
func F_representClusterNodeFlags(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v8 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v25 = int32(0)
		goto L1
	}
L1:
	;
	if l1&int32(16) == int32(0) {
		v35 = l0
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v25 = v24
	goto L1
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v25 = v21
	goto L1
L4:
	;
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v25 = v18
	goto L1
L5:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v25 = v15
	goto L1
L6:
	;
	v25 = int32(base.Ui32(v8) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	if l1&int32(1) == int32(0) {
		v43 = v35
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v31 = F_sdscat(m, l0, int32(_a_F_representClusterNodeFlags_6))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v35 = v31
	goto L7
L11:
	;
	if l1&int32(2) == int32(0) {
		v51 = v43
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v41 = F_sdscat(m, v35, int32(_a_F_representClusterNodeFlags_7))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v43 = v41
	goto L11
L14:
	;
	if l1&int32(4) == int32(0) {
		v59 = v51
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v49 = F_sdscat(m, v43, int32(_a_F_representClusterNodeFlags_8))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v51 = v49
	goto L14
L17:
	;
	if l1&int32(8) == int32(0) {
		v67 = v59
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v57 = F_sdscat(m, v51, int32(_a_F_representClusterNodeFlags_5))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v59 = v57
	goto L17
L20:
	;
	if l1&int32(32) == int32(0) {
		v75 = v67
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v65 = F_sdscat(m, v59, int32(_a_F_representClusterNodeFlags_4))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v67 = v65
	goto L20
L23:
	;
	if l1&int32(64) == int32(0) {
		v83 = v75
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v73 = F_sdscat(m, v67, int32(_a_F_representClusterNodeFlags_3))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	v75 = v73
	goto L23
L26:
	;
	if l1&int32(512) == int32(0) {
		v91 = v83
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v81 = F_sdscat(m, v75, int32(_a_F_representClusterNodeFlags_2))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v83 = v81
	goto L26
L29:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+int32(-1)))))
	switch v95 & int32(7) {
	case 0:
		goto L37
	case 1:
		goto L36
	case 2:
		goto L35
	case 3:
		goto L34
	case 4:
		goto L33
	default:
		v112 = int32(0)
		goto L32
	}
L30:
	;
	v89 = F_sdscat(m, v83, int32(_a_F_representClusterNodeFlags_1))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	v91 = v89
	goto L29
L32:
	;
	if v112 != v25 {
		v117 = v91
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v91+int32(-17))))
	v112 = v111
	goto L32
L34:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v91+int32(-9))))
	v112 = v108
	goto L32
L35:
	;
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91+int32(-5)))))
	v112 = v105
	goto L32
L36:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+int32(-3)))))
	v112 = v102
	goto L32
L37:
	;
	v112 = int32(base.Ui32(v95) >> (uint(int32(3)) % 32))
	goto L32
L38:
	;
	F_sdsIncrLen(m, v117, int32(-1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L9
	} else {
		goto L41
	}
L39:
	;
	v115 = F_sdscat(m, v91, int32(_a_F_representClusterNodeFlags_0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	v117 = v115
	goto L38
L41:
	;
	return v117
}
func F_reqresAppendRequest(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F_resetLastWrittenBuf(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = int64(0)
	return
}
func F_resizeReplicationBacklog(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v2 = *(*int64)(unsafe.Add(mBase, _c_F_resizeReplicationBacklog[0]))
	if int64(16383) < v2 {
	} else {
		*(*int64)(unsafe.Add(mBase, _c_F_resizeReplicationBacklog[0])) = int64(16384)
	}
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_resizeReplicationBacklog[1]))
	if v9 == int32(0) {
		return
	} else {
		F_incrementalTrimReplicationBacklog(m, int32(64))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			return
		}
	}
}
func F_restoreBackupConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(_a_F_restoreBackupConfig_0)
	if l2 < int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l4 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v25 = int32(0)
	goto L3
L3:
	;
	v29 = v25 << (uint(int32(2)) % 32)
	v30 = l0 + v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v32 = l1 + v29
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v36 = F_performInterfaceSet(m, v31, v33, v12+int32(44))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if l3 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L5:
	;
	v57 = v25 + int32(1)
	if v57 != l2 {
		v25 = v57
		goto L3
	} else {
		goto L11
	}
L6:
	;
	return
L7:
	;
	if v36 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_restoreBackupConfig[0]))
	if int32(3) < v39 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v42
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v47
	F__serverLog(m, int32(3), int32(_a_F_restoreBackupConfig_1), v12+int32(32))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L5
L11:
	;
	goto L4
L12:
	;
	v68 = int32(0)
	goto L13
L13:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l3+v68<<(uint(int32(2))%32))))
	if v74 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L1
L15:
	;
	v79 = m.T0[v74].(func(*base.Module, int32) int32)(m, v12+int32(44))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L17
	}
L16:
	;
	v94 = v68 + int32(1)
	if v94 != l2 {
		v68 = v94
		goto L13
	} else {
		goto L21
	}
L17:
	;
	if v79 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_restoreBackupConfig[0]))
	if int32(3) < v82 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v85
	F__serverLog(m, int32(3), int32(_a_F_restoreBackupConfig_2), v12+int32(16))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	goto L14
L22:
	;
	m.G0 = v12 + int32(48)
	return
L23:
	;
	v110 = F_moduleConfigApplyConfig(m, l4, v12+int32(44), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	if v110 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_restoreBackupConfig[0]))
	if int32(3) < v113 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v116
	F__serverLog(m, int32(3), int32(_a_F_restoreBackupConfig_2), v12)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	goto L22
}
func F_rewriteSortedSetObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v58 int64
	_ = v58
	var v61 int64
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v106 int64
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int64
	_ = v141
	var v144 int64
	_ = v144
	var v146 int32
	_ = v146
	var v150 int64
	_ = v150
	var v153 int64
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 float64
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
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
	var v215 int32
	_ = v215
	var v219 int64
	_ = v219
	var v222 int64
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v239 int32
	_ = v239
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	v10 = m.G0
	v12 = v10 - int32(64)
	m.G0 = v12
	v14 = F_zsetLength(m, l2)
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
	v18 = base.I64_extend_i32_u(v14)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch int32(base.Ui32(v19)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L9
	default:
		goto L5
	case 4:
		goto L10
	}
L3:
	;
	m.G0 = v12 + int32(64)
	return v270
L4:
	;
	v270 = int32(0)
	goto L3
L5:
	;
	F__serverPanic_1(m, int32(_a_F_rewriteSortedSetObject_0), int32(2054), int32(_a_F_rewriteSortedSetObject_1), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L86
	}
L6:
	;
	F__serverAssert(m, int32(_a_F_rewriteSortedSetObject_2), int32(_a_F_rewriteSortedSetObject_0), int32(2004))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L85
	}
L7:
	;
	F__serverAssert(m, int32(_a_F_rewriteSortedSetObject_3), int32(_a_F_rewriteSortedSetObject_0), int32(2002))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L84
	}
L8:
	;
	v270 = int32(1)
	goto L3
L9:
	;
	v110 = F_objectGetVal(m, l2)
	mBase = m.M
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v112
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(-1)
	if v111 == v112 {
		goto L45
	} else {
		goto L46
	}
L10:
	;
	v26 = F_objectGetVal(m, l2)
	mBase = m.M
	v28 = F_lpSeek(m, v26, int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v28
	if v28 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v33 = F_lpNext(m, v26, v28)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v33
	if v33 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v41 = v28
	v43 = v18
	v46 = int64(0)
	goto L15
L15:
	;
	v50 = F_lpGetValue(m, v41, v12+int32(52), v12)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v53 = F_zzlGetScore(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v46 != int64(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v81 = F_rioWriteBulkDouble(m, l0, v53)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L30
	}
L20:
	;
	v58 = int64(64)
	if v43 < v58 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v61 = v43
	goto L23
L22:
	;
	v61 = v58
	goto L23
L23:
	;
	v67 = F_rioWriteBulkCount(m, l0, int32(42), base.I32_wrap_i64(v61)<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v67 == int32(0) {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v73 = F_rioWriteBulkString(m, l0, int32(_a_F_rewriteSortedSetObject_4), int32(4))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v73 == int32(0) {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v77 = F_rioWriteBulkObject(m, l0, l1)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v77 == int32(0) {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	goto L19
L30:
	;
	if v81 == int32(0) {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v50 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	F_zzlNext(m, v26, v12+int32(60), v12+int32(56))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L39
	}
L33:
	;
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v91 = F_rioWriteBulkLongLong(m, l0, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L37
	}
L34:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v88 = F_rioWriteBulkString(m, l0, v50, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v88 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L4
L37:
	;
	if v91 == int32(0) {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	goto L32
L39:
	;
	v103 = v46 + int64(1)
	if v103 == int64(64) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v106 = int64(0)
	goto L42
L41:
	;
	v106 = v103
	goto L42
L42:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	if v109 != 0 {
		v41 = v109
		v43 = v43 + int64(-1)
		v46 = v106
		goto L15
	} else {
		goto L43
	}
L43:
	;
	goto L8
L44:
	;
	v132 = F_hashtableNext(m, v12, v12+int32(60))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L49
	}
L45:
	;
	goto L44
L46:
	;
	goto L45
L48:
	;
	F_hashtableCleanupIterator(m, v12)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L83
	}
L49:
	;
	if v132 == int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v141 = v18
	v144 = int64(0)
	goto L51
L51:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	if v144 != int64(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L48
L53:
	;
	v219 = v144 + int64(1)
	if v219 == int64(64) {
		goto L78
	} else {
		goto L79
	}
L54:
	;
	F_hashtableCleanupIterator(m, v12)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L77
	}
L55:
	;
	v174 = v146 + int32(16)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v178 = v174 + v175<<(uint(int32(3))%32)
	v179 = int32(*(*int8)(unsafe.Add(mBase, uint32(v178))))
	v180 = v178 + v179
	goto L66
L56:
	;
	v150 = int64(64)
	if v141 < v150 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v153 = v141
	goto L59
L58:
	;
	v153 = v150
	goto L59
L59:
	;
	v159 = F_rioWriteBulkCount(m, l0, int32(42), base.I32_wrap_i64(v153)<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v159 == int32(0) {
		goto L54
	} else {
		goto L61
	}
L61:
	;
	v165 = F_rioWriteBulkString(m, l0, int32(_a_F_rewriteSortedSetObject_4), int32(4))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	if v165 == int32(0) {
		goto L54
	} else {
		goto L63
	}
L63:
	;
	v169 = F_rioWriteBulkObject(m, l0, l1)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v169 == int32(0) {
		goto L54
	} else {
		goto L65
	}
L65:
	;
	goto L55
L66:
	;
	v183 = *(*float64)(unsafe.Add(mBase, uint32(v146)))
	v184 = F_rioWriteBulkDouble(m, l0, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v184 == int32(0) {
		goto L54
	} else {
		goto L68
	}
L68:
	;
	v188 = int32(0)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+v188))))
	switch v191 & int32(7) {
	case 0:
		goto L74
	case 1:
		goto L73
	case 2:
		goto L72
	case 3:
		goto L71
	case 4:
		goto L70
	default:
		v208 = v188
		goto L69
	}
L69:
	;
	v209 = F_rioWriteBulkString(m, l0, v180+int32(1), v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L75
	}
L70:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v180+int32(-16))))
	v208 = v207
	goto L69
L71:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v180+int32(-8))))
	v208 = v204
	goto L69
L72:
	;
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180+int32(-4)))))
	v208 = v201
	goto L69
L73:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+int32(-2)))))
	v208 = v198
	goto L69
L74:
	;
	v208 = int32(base.Ui32(v191) >> (uint(int32(3)) % 32))
	goto L69
L75:
	;
	if v209 != 0 {
		goto L53
	} else {
		goto L76
	}
L76:
	;
	goto L54
L77:
	;
	v270 = int32(0)
	goto L3
L78:
	;
	v222 = int64(0)
	goto L80
L79:
	;
	v222 = v219
	goto L80
L80:
	;
	v227 = F_hashtableNext(m, v12, v12+int32(60))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	if v227 != 0 {
		v141 = v141 + int64(-1)
		v144 = v222
		goto L51
	} else {
		goto L82
	}
L82:
	;
	goto L52
L83:
	;
	goto L8
L84:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_rint(m *base.Module, l0 float64) float64 {
	return base.F64_nearest(l0)
}
func F_roleCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int64
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v170 int32
	_ = v170
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_roleCommand[0]))
	if v14 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(64)
	return
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_roleCommand[1]))
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_sentinelRoleCommand(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	goto L1
L6:
	;
	F_addReplyArrayLen(m, l0, int32(5))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L37
	}
L7:
	;
	F_addReplyArrayLen(m, l0, int32(3))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_roleCommand_0), int32(6))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int64)(unsafe.Add(mBase, _c_F_roleCommand[2]))
	F_addReplyLongLong(m, l0, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v32 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_roleCommand[3]))
	v37 = v11 + int32(56)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v38
	goto L12
L12:
	;
	v42 = int32(0)
	v44 = v11 + int32(56)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v46 == v42 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	F_setDeferredArrayLen(m, l0, v32, v125)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L36
	}
L14:
	;
	if v46 == int32(0) {
		v125 = v42
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46+base.B2i32(v49 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v55
	goto L15
L17:
	;
	v62 = v42
	v63 = v46
	goto L18
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+104))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+152))
	if v69 != 0 {
		v85 = v68
		v86 = v69
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v125 = v105
	goto L13
L20:
	;
	v109 = v11 + int32(56)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v111 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v87 != int32(9) {
		v105 = v62
		goto L20
	} else {
		goto L27
	}
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v70 == int32(0) {
		v105 = v62
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+24))
	if v74 == int32(0) {
		v105 = v62
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v80 = m.T0[v74].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v70, v11, int32(46), int32(0), int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	if v80 == int32(-1) {
		v105 = v62
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v67)+104))
	v85 = v84
	v86 = v11
	goto L21
L27:
	;
	F_addReplyArrayLen(m, l0, int32(3))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	F_addReplyBulkCString(m, l0, v86)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v67)+104))
	v96 = int64(*(*int32)(unsafe.Add(mBase, uint32(v95)+148)))
	F_addReplyBulkLongLong(m, l0, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v67)+104))
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v99)+64))
	F_addReplyBulkLongLong(m, l0, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v105 = v62 + int32(1)
	goto L20
L32:
	;
	if v111 != 0 {
		v62 = v105
		v63 = v111
		goto L18
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v111+base.B2i32(v114 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v120
	goto L33
L35:
	;
	goto L19
L36:
	;
	goto L1
L37:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_roleCommand_1), int32(5))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_roleCommand[1]))
	F_addReplyBulkCString(m, l0, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v144 = int64(*(*int32)(unsafe.Add(mBase, _c_F_roleCommand[4])))
	F_addReplyLongLong(m, l0, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_roleCommand[5]))
	if base.Ui32(v148+int32(-13)) <= base.Ui32(int32(-11)) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	F_addReplyBulkCString(m, l0, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L49
	}
L42:
	;
	switch v148 {
	case 0:
		v160 = int32(_a_F_roleCommand_2)
		goto L41
	case 1:
		goto L48
	case 2:
		goto L47
	default:
		goto L44
	case 13:
		goto L46
	case 14:
		goto L45
	}
L43:
	;
	v160 = int32(_a_F_roleCommand_3)
	goto L41
L44:
	;
	v160 = int32(_a_F_roleCommand_4)
	goto L41
L45:
	;
	v160 = int32(_a_F_roleCommand_5)
	goto L41
L46:
	;
	v160 = int32(_a_F_roleCommand_6)
	goto L41
L47:
	;
	v160 = int32(_a_F_roleCommand_7)
	goto L41
L48:
	;
	v160 = int32(_a_F_roleCommand_8)
	goto L41
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_roleCommand[6]))
	if v164 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	F_addReplyLongLong(m, l0, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v164)+104))
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v166)+48))
	v168 = v167
	goto L50
L52:
	;
	v168 = int64(-1)
	goto L50
L53:
	;
	goto L1
}
func F_round(m *base.Module, l0 float64) float64 {
	var v5 int64
	_ = v5
	var v10 int32
	_ = v10
	var v18 float64
	_ = v18
	var v23 float64
	_ = v23
	var v31 float64
	_ = v31
	var v38 float64
	_ = v38
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	v5 = base.I64_reinterpret_f64(l0)
	v10 = base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(52))%64))) & int32(2047)
	if base.Ui32(int32(1074)) < base.Ui32(v10) {
		v43 = l0
		return v43
	} else {
		if base.Ui32(int32(1021)) < base.Ui32(v10) {
			v18 = base.F64_abs(l0)
			v23 = base.F64_sub(base.F64_add(base.F64_add(v18, float64(4.503599627370496e+15)), float64(-4.503599627370496e+15)), v18)
			if base.F64_gt(v23, float64(0.5)) == int32(0) {
				v31 = base.F64_add(v18, v23)
				if base.F64_le(v23, float64(-0.5)) == int32(0) {
					v38 = v31
				} else {
					v38 = base.F64_add(v31, float64(1))
				}
			} else {
				v38 = base.F64_add(base.F64_add(v18, v23), float64(-1))
			}
			if v5 < int64(0) {
				v42 = base.F64_neg(v38)
			} else {
				v42 = v38
			}
			v43 = v42
			return v43
		} else {
			return base.F64_mul(l0, float64(0))
		}
	}
}
func F_roundf(m *base.Module, l0 float32) float32 {
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v17 float32
	_ = v17
	var v22 float32
	_ = v22
	var v30 float32
	_ = v30
	var v37 float32
	_ = v37
	var v41 float32
	_ = v41
	var v42 float32
	_ = v42
	v5 = base.I32_reinterpret_f32(l0)
	v9 = int32(base.Ui32(v5)>>(uint(int32(23))%32)) & int32(255)
	if base.Ui32(int32(149)) < base.Ui32(v9) {
		v42 = l0
		return v42
	} else {
		if base.Ui32(int32(125)) < base.Ui32(v9) {
			v17 = base.F32_abs(l0)
			v22 = base.F32_sub(base.F32_add(base.F32_add(v17, float32(8.388608e+06)), float32(-8.388608e+06)), v17)
			if base.F32_gt(v22, float32(0.5)) == int32(0) {
				v30 = base.F32_add(v17, v22)
				if base.F32_le(v22, float32(-0.5)) == int32(0) {
					v37 = v30
				} else {
					v37 = base.F32_add(v30, float32(1))
				}
			} else {
				v37 = base.F32_add(base.F32_add(v17, v22), float32(-1))
			}
			if v5 < int32(0) {
				v41 = base.F32_neg(v37)
			} else {
				v41 = v37
			}
			v42 = v41
			return v42
		} else {
			return base.F32_mul(l0, float32(0))
		}
	}
}
func F_rpoplpushCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_lmoveGenericCommand(m, l0, int32(1), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
