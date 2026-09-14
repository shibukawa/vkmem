package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F___math_oflow(m *base.Module, l0 int32) float64 {
	var v2 float64
	_ = v2
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	v2 = float64(3.105036184601418e+231)
	if l0 != 0 {
		v4 = base.F64_neg(v2)
	} else {
		v4 = v2
	}
	v5 = F_fp_barrier_1(m, v4)
	return base.F64_mul(v2, v5)
}
func F_math_acos(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v27 float64
	_ = v27
	var v38 float64
	_ = v38
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v57 float64
	_ = v57
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v69 float64
	_ = v69
	var v75 float64
	_ = v75
	var v78 float64
	_ = v78
	var v83 float64
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	v3 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v11 = base.I64_reinterpret_f64(v3)
		v16 = base.I32_wrap_i64(int64(base.Ui64(v11)>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v16) < base.Ui32(int32(1072693248)) {
			if base.Ui32(int32(1071644671)) < base.Ui32(v16) {
				if int64(-1) < v11 {
					v62 = base.F64_mul(base.F64_sub(float64(1), v3), float64(0.5))
					v63 = F_sqrt(m, v62)
					mBase = m.M
					v64 = F_R_1(m, v62)
					mBase = m.M
					v69 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v63) & int64(-4294967296))
					v75 = base.F64_add(base.F64_add(base.F64_mul(v63, v64), base.F64_div(base.F64_sub(v62, base.F64_mul(v69, v69)), base.F64_add(v63, v69))), v69)
					v78 = base.F64_add(v75, v75)
					v83 = v78
				} else {
					v50 = base.F64_mul(base.F64_add(v3, float64(1)), float64(0.5))
					v51 = F_sqrt(m, v50)
					mBase = m.M
					v52 = F_R_1(m, v50)
					mBase = m.M
					v57 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v51, base.F64_add(base.F64_mul(v51, v52), float64(-6.123233995736766e-17))))
					v83 = base.F64_add(v57, v57)
				}
			} else {
				if base.Ui32(v16) < base.Ui32(int32(1012924417)) {
					v78 = float64(1.5707963267948966)
					v83 = v78
				} else {
					v38 = F_R_1(m, base.F64_mul(v3, v3))
					mBase = m.M
					v83 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v3, v38)), v3), float64(1.5707963267948966))
				}
			}
		} else {
			if v16+int32(-1072693248)|base.I32_wrap_i64(v11) != 0 {
				v83 = base.F64_div(float64(0), base.F64_sub(v3, v3))
			} else {
				if int64(-1) < v11 {
					v27 = float64(0)
				} else {
					v27 = float64(3.141592653589793)
				}
				v83 = v27
			}
		}
		v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v85))) = v83
		v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v89 + int32(16)
		return int32(1)
	}
}
func F_math_atan2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v17 int64
	_ = v17
	var v22 int64
	_ = v22
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 float64
	_ = v36
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v68 float64
	_ = v68
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v100 float64
	_ = v100
	var v102 float64
	_ = v102
	var v110 float64
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	v3 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_luaL_checknumber(m, l0, int32(2))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v17 = F___DOUBLE_BITS_2(m, v8)
			mBase = m.M
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v17&int64(9223372036854775807)) {
				v110 = base.F64_add(v3, v8)
			} else {
				v22 = F___DOUBLE_BITS_2(m, v3)
				mBase = m.M
				if base.Ui64(v22&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
					v28 = base.I64_reinterpret_f64(v8)
					v31 = base.I32_wrap_i64(int64(base.Ui64(v28) >> (uint(int64(32)) % 64)))
					v34 = base.I32_wrap_i64(v28)
					if v31+int32(-1072693248)|v34 != 0 {
						v40 = int32(base.Ui32(v31)>>(uint(int32(30))%32)) & int32(2)
						v41 = base.I64_reinterpret_f64(v3)
						v45 = v40 | base.I32_wrap_i64(int64(base.Ui64(v41)>>(uint(int64(63))%64)))
						v50 = base.I32_wrap_i64(int64(base.Ui64(v41)>>(uint(int64(32))%64))) & int32(2147483647)
						if v50|base.I32_wrap_i64(v41) != 0 {
							v56 = v31 & int32(2147483647)
							if v56|v34 != 0 {
								if v56 != int32(2146435072) {
									if v50 == int32(2146435072) {
										v110 = base.F64_copysign(float64(1.5707963267948966), v3)
									} else {
										if base.Ui32(v50) <= base.Ui32(v56+int32(67108864)) {
											if v40 == int32(0) {
												v84 = F_fabs(m, base.F64_div(v3, v8))
												mBase = m.M
												v85 = F_atan(m, v84)
												mBase = m.M
												v86 = v85
											} else {
												if base.Ui32(v50+int32(67108864)) < base.Ui32(v56) {
													v86 = float64(0)
												} else {
													v84 = F_fabs(m, base.F64_div(v3, v8))
													mBase = m.M
													v85 = F_atan(m, v84)
													mBase = m.M
													v86 = v85
												}
											}
											switch v45 {
											default:
												v102 = v86
												v110 = v102
											case 1:
												v110 = base.F64_neg(v86)
											case 2:
												v110 = base.F64_sub(float64(3.141592653589793), base.F64_add(v86, float64(-1.2246467991473532e-16)))
											case 3:
												v110 = base.F64_add(base.F64_add(v86, float64(-1.2246467991473532e-16)), float64(-3.141592653589793))
											}
										} else {
											v110 = base.F64_copysign(float64(1.5707963267948966), v3)
										}
									}
								} else {
									if v50 != int32(2146435072) {
										v100 = *(*float64)(unsafe.Add(mBase, uint32(v45<<(uint(int32(3))%32))+uint32(_consts[290])))
										v102 = v100
										v110 = v102
									} else {
										v68 = *(*float64)(unsafe.Add(mBase, uint32(v45<<(uint(int32(3))%32))+uint32(_consts[291])))
										v110 = v68
									}
								}
							} else {
								v110 = base.F64_copysign(float64(1.5707963267948966), v3)
							}
						} else {
							switch v45 {
							default:
								v102 = v3
								v110 = v102
							case 2:
								v110 = float64(3.141592653589793)
							case 3:
								v110 = float64(-3.141592653589793)
							}
						}
					} else {
						v36 = F_atan(m, v3)
						mBase = m.M
						v110 = v36
					}
				} else {
					v110 = base.F64_add(v3, v8)
				}
			}
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v112))) = v110
			v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v116 + int32(16)
			return int32(1)
		}
	}
}
func F_math_cos(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v26 float64
	_ = v26
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v35 float64
	_ = v35
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v42 float64
	_ = v42
	var v45 float64
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	v3 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v10 = m.G0
		v12 = v10 - int32(16)
		m.G0 = v12
		v19 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v3))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1072243195)) < base.Ui32(v19) {
			if base.Ui32(v19) < base.Ui32(int32(2146435072)) {
				v30 = F___rem_pio2(m, v3, v12)
				mBase = m.M
				v31 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
				v32 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
				switch v30 & int32(3) {
				default:
					v35 = F___cos(m, v32, v31)
					mBase = m.M
					v45 = v35
				case 1:
					v37 = F___sin(m, v32, v31, int32(1))
					mBase = m.M
					v45 = base.F64_neg(v37)
				case 2:
					v39 = F___cos(m, v32, v31)
					mBase = m.M
					v45 = base.F64_neg(v39)
				case 3:
					v42 = F___sin(m, v32, v31, int32(1))
					mBase = m.M
					v45 = v42
				}
			} else {
				v45 = base.F64_sub(v3, v3)
			}
		} else {
			if base.Ui32(v19) < base.Ui32(int32(1044816030)) {
				v45 = float64(1)
			} else {
				v26 = F___cos(m, v3, float64(0))
				mBase = m.M
				v45 = v26
			}
		}
		m.G0 = v12 + int32(16)
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v50))) = v45
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v54 + int32(16)
		return int32(1)
	}
}
func F_math_deg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v3 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v10))) = base.F64_div(v3, float64(0.017453292519943295))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v14 + int32(16)
		return int32(1)
	}
}
func F_math_exp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v57 float64
	_ = v57
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v71 float64
	_ = v71
	var v74 float64
	_ = v74
	var v77 int64
	_ = v77
	var v82 int32
	_ = v82
	var v85 float64
	_ = v85
	var v88 float64
	_ = v88
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v102 float64
	_ = v102
	var v113 float64
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	v3 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v13 = F_top12_1(m, v3)
		mBase = m.M
		v15 = v13 & int32(2047)
		v17 = F_top12_1(m, float64(5.551115123125783e-17))
		mBase = m.M
		v20 = F_top12_1(m, float64(512))
		mBase = m.M
		if base.Ui32(v20-v17) <= base.Ui32(v15-v17) {
			if base.Ui32(v17) <= base.Ui32(v15) {
				v28 = F_top12_1(m, float64(1024))
				mBase = m.M
				if base.Ui32(v15) < base.Ui32(v28) {
					v45 = int32(0)
					v46 = int32(0)
					v47 = *(*float64)(unsafe.Add(mBase, _consts[1217]))
					v50 = *(*float64)(unsafe.Add(mBase, _consts[1218]))
					v51 = base.F64_add(base.F64_mul(v3, v47), v50)
					v52 = base.F64_sub(v51, v50)
					v54 = *(*float64)(unsafe.Add(mBase, _consts[1219]))
					v57 = *(*float64)(unsafe.Add(mBase, _consts[1220]))
					v60 = base.F64_add(base.F64_mul(v52, v54), base.F64_add(base.F64_mul(v52, v57), v3))
					v61 = base.F64_mul(v60, v60)
					v64 = *(*float64)(unsafe.Add(mBase, _consts[1221]))
					v67 = *(*float64)(unsafe.Add(mBase, _consts[1222]))
					v71 = *(*float64)(unsafe.Add(mBase, _consts[1223]))
					v74 = *(*float64)(unsafe.Add(mBase, _consts[1224]))
					v77 = base.I64_reinterpret_f64(v51)
					v82 = base.I32_wrap_i64(v77) << (uint(int32(4)) % 32) & int32(2032)
					v85 = *(*float64)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1225])))
					v88 = base.F64_add(base.F64_mul(base.F64_mul(v61, v61), base.F64_add(base.F64_mul(v60, v64), v67)), base.F64_add(base.F64_mul(v61, base.F64_add(base.F64_mul(v60, v71), v74)), base.F64_add(v85, v60)))
					v91 = *(*int64)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1226])))
					v94 = v91 + v77<<(uint(int64(45))%64)
					if v45 != 0 {
						v96 = base.F64_reinterpret_i64(v94)
						v102 = base.F64_add(base.F64_mul(v96, v88), v96)
						v113 = v102
					} else {
						v95 = F_specialcase_1(m, v88, v94, v77)
						mBase = m.M
						v113 = v95
					}
				} else {
					v31 = base.I64_reinterpret_f64(v3)
					if v31 == int64(-4503599627370496) {
						v102 = float64(0)
						v113 = v102
					} else {
						v35 = F_top12_1(m, math.Float64frombits(uint64(0x7ff0000000000000)))
						mBase = m.M
						if base.Ui32(v15) < base.Ui32(v35) {
							if int64(-1) < v31 {
								v44 = F___math_oflow(m, int32(0))
								mBase = m.M
								v113 = v44
							} else {
								v42 = F___math_uflow(m, int32(0))
								mBase = m.M
								v113 = v42
							}
						} else {
							v113 = base.F64_add(v3, float64(1))
						}
					}
				}
			} else {
				v113 = base.F64_add(v3, float64(1))
			}
		} else {
			v45 = v15
			v46 = int32(0)
			v47 = *(*float64)(unsafe.Add(mBase, _consts[1217]))
			v50 = *(*float64)(unsafe.Add(mBase, _consts[1218]))
			v51 = base.F64_add(base.F64_mul(v3, v47), v50)
			v52 = base.F64_sub(v51, v50)
			v54 = *(*float64)(unsafe.Add(mBase, _consts[1219]))
			v57 = *(*float64)(unsafe.Add(mBase, _consts[1220]))
			v60 = base.F64_add(base.F64_mul(v52, v54), base.F64_add(base.F64_mul(v52, v57), v3))
			v61 = base.F64_mul(v60, v60)
			v64 = *(*float64)(unsafe.Add(mBase, _consts[1221]))
			v67 = *(*float64)(unsafe.Add(mBase, _consts[1222]))
			v71 = *(*float64)(unsafe.Add(mBase, _consts[1223]))
			v74 = *(*float64)(unsafe.Add(mBase, _consts[1224]))
			v77 = base.I64_reinterpret_f64(v51)
			v82 = base.I32_wrap_i64(v77) << (uint(int32(4)) % 32) & int32(2032)
			v85 = *(*float64)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1225])))
			v88 = base.F64_add(base.F64_mul(base.F64_mul(v61, v61), base.F64_add(base.F64_mul(v60, v64), v67)), base.F64_add(base.F64_mul(v61, base.F64_add(base.F64_mul(v60, v71), v74)), base.F64_add(v85, v60)))
			v91 = *(*int64)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1226])))
			v94 = v91 + v77<<(uint(int64(45))%64)
			if v45 != 0 {
				v96 = base.F64_reinterpret_i64(v94)
				v102 = base.F64_add(base.F64_mul(v96, v88), v96)
				v113 = v102
			} else {
				v95 = F_specialcase_1(m, v88, v94, v77)
				mBase = m.M
				v113 = v95
			}
		}
		v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v115)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v115))) = v113
		v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v119 + int32(16)
		return int32(1)
	}
}
func F_math_log10(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v17 int64
	_ = v17
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v43 int64
	_ = v43
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v72 float64
	_ = v72
	var v75 float64
	_ = v75
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v115 float64
	_ = v115
	var v127 float64
	_ = v127
	var v149 float64
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	v3 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v17 = base.I64_reinterpret_f64(v3)
		if int64(4503599627370495) < v17 {
			if base.Ui64(int64(9218868437227405311)) < base.Ui64(v17) {
				v127 = v3
				v149 = v127
			} else {
				v32 = int32(-1023)
				v34 = int64(base.Ui64(v17) >> (uint(int64(32)) % 64))
				if v34 == int64(1072693248) {
					if base.I32_wrap_i64(v17) != 0 {
						v48 = v17
						v49 = v32
						v51 = int32(1072693248)
						v53 = v51 + int32(614242)
						v57 = base.F64_convert_i32_s(v49 + int32(base.Ui32(v53)>>(uint(int32(20))%32)))
						v59 = base.F64_mul(v57, float64(0.30102999566361177))
						v72 = base.F64_add(base.F64_reinterpret_i64(base.I64_extend_i32_u(v53&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)|v48&int64(4294967295)), float64(-1))
						v75 = base.F64_mul(v72, base.F64_mul(v72, float64(0.5)))
						v80 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v72, v75)) & int64(-4294967296))
						v81 = float64(0.4342944818781689)
						v82 = base.F64_mul(v80, v81)
						v83 = base.F64_add(v59, v82)
						v88 = base.F64_div(v72, base.F64_add(v72, float64(2)))
						v89 = base.F64_mul(v88, v88)
						v90 = base.F64_mul(v89, v89)
						v115 = base.F64_add(base.F64_mul(v88, base.F64_add(v75, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v89, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v72, v80), v75))
						v127 = base.F64_add(v83, base.F64_add(base.F64_add(v82, base.F64_sub(v59, v83)), base.F64_add(base.F64_mul(v115, v81), base.F64_add(base.F64_mul(v57, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v115, v80), float64(2.5082946711645275e-11))))))
						v149 = v127
					} else {
						v149 = float64(0)
					}
				} else {
					v48 = v17
					v49 = v32
					v51 = base.I32_wrap_i64(v34)
					v53 = v51 + int32(614242)
					v57 = base.F64_convert_i32_s(v49 + int32(base.Ui32(v53)>>(uint(int32(20))%32)))
					v59 = base.F64_mul(v57, float64(0.30102999566361177))
					v72 = base.F64_add(base.F64_reinterpret_i64(base.I64_extend_i32_u(v53&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)|v48&int64(4294967295)), float64(-1))
					v75 = base.F64_mul(v72, base.F64_mul(v72, float64(0.5)))
					v80 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v72, v75)) & int64(-4294967296))
					v81 = float64(0.4342944818781689)
					v82 = base.F64_mul(v80, v81)
					v83 = base.F64_add(v59, v82)
					v88 = base.F64_div(v72, base.F64_add(v72, float64(2)))
					v89 = base.F64_mul(v88, v88)
					v90 = base.F64_mul(v89, v89)
					v115 = base.F64_add(base.F64_mul(v88, base.F64_add(v75, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v89, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v72, v80), v75))
					v127 = base.F64_add(v83, base.F64_add(base.F64_add(v82, base.F64_sub(v59, v83)), base.F64_add(base.F64_mul(v115, v81), base.F64_add(base.F64_mul(v57, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v115, v80), float64(2.5082946711645275e-11))))))
					v149 = v127
				}
			}
		} else {
			if base.F64_ne(v3, float64(0)) != 0 {
				if int64(-1) < v17 {
					v43 = base.I64_reinterpret_f64(base.F64_mul(v3, float64(1.8014398509481984e+16)))
					v48 = v43
					v49 = int32(-1077)
					v51 = base.I32_wrap_i64(int64(base.Ui64(v43) >> (uint(int64(32)) % 64)))
					v53 = v51 + int32(614242)
					v57 = base.F64_convert_i32_s(v49 + int32(base.Ui32(v53)>>(uint(int32(20))%32)))
					v59 = base.F64_mul(v57, float64(0.30102999566361177))
					v72 = base.F64_add(base.F64_reinterpret_i64(base.I64_extend_i32_u(v53&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)|v48&int64(4294967295)), float64(-1))
					v75 = base.F64_mul(v72, base.F64_mul(v72, float64(0.5)))
					v80 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v72, v75)) & int64(-4294967296))
					v81 = float64(0.4342944818781689)
					v82 = base.F64_mul(v80, v81)
					v83 = base.F64_add(v59, v82)
					v88 = base.F64_div(v72, base.F64_add(v72, float64(2)))
					v89 = base.F64_mul(v88, v88)
					v90 = base.F64_mul(v89, v89)
					v115 = base.F64_add(base.F64_mul(v88, base.F64_add(v75, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v89, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v72, v80), v75))
					v127 = base.F64_add(v83, base.F64_add(base.F64_add(v82, base.F64_sub(v59, v83)), base.F64_add(base.F64_mul(v115, v81), base.F64_add(base.F64_mul(v57, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v115, v80), float64(2.5082946711645275e-11))))))
					v149 = v127
				} else {
					v149 = base.F64_div(base.F64_sub(v3, v3), float64(0))
				}
			} else {
				v149 = base.F64_div(float64(-1), base.F64_mul(v3, v3))
			}
		}
		v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v151))) = v149
		v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v155 + int32(16)
		return int32(1)
	}
}
func F_math_max(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 float64
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 float64
	_ = v23
	var v26 float64
	_ = v26
	var v27 int32
	_ = v27
	var v29 float64
	_ = v29
	var v36 float64
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = (v8 - v9) >> (uint(int32(4)) % 32)
	goto L1
L1:
	;
	v14 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	if v12 < int32(2) {
		v36 = v14
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v40))) = v36
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v44 + int32(16)
	goto L13
L5:
	;
	v21 = int32(2)
	v23 = v14
	goto L6
L6:
	;
	v26 = F_luaL_checknumber(m, l0, v21)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v36 = v29
	goto L4
L8:
	;
	if base.F64_gt(v26, v23) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v29 = v26
	goto L11
L10:
	;
	v29 = v23
	goto L11
L11:
	;
	if v21 != v12 {
		v21 = v21 + int32(1)
		v23 = v29
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L7
L13:
	;
	return int32(1)
}
func F_math_min(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 float64
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 float64
	_ = v23
	var v26 float64
	_ = v26
	var v27 int32
	_ = v27
	var v29 float64
	_ = v29
	var v36 float64
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = (v8 - v9) >> (uint(int32(4)) % 32)
	goto L1
L1:
	;
	v14 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	if v12 < int32(2) {
		v36 = v14
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v40))) = v36
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v44 + int32(16)
	goto L13
L5:
	;
	v21 = int32(2)
	v23 = v14
	goto L6
L6:
	;
	v26 = F_luaL_checknumber(m, l0, v21)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v36 = v29
	goto L4
L8:
	;
	if base.F64_lt(v26, v23) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v29 = v26
	goto L11
L10:
	;
	v29 = v23
	goto L11
L11:
	;
	if v21 != v12 {
		v21 = v21 + int32(1)
		v23 = v29
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L7
L13:
	;
	return int32(1)
}
func F_math_modf(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 float64
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v44 int64
	_ = v44
	var v56 int64
	_ = v56
	var v60 float64
	_ = v60
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v6 + int32(8)
		v19 = base.I64_reinterpret_f64(v9)
		v24 = base.I32_wrap_i64(int64(base.Ui64(v19)>>(uint(int64(52))%64))) & int32(2047)
		v26 = v24 + int32(-1023)
		if base.Ui32(v24) < base.Ui32(int32(1075)) {
			if base.Ui32(int32(1022)) < base.Ui32(v24) {
				v44 = base.I64_extend_i32_u(v26)
				if v19<<(uint(v44)%64)&int64(4503599627370495) != int64(0) {
					v56 = int64(-4503599627370496) >> (uint(v44) % 64) & v19
					*(*int64)(unsafe.Add(mBase, uint32(v14))) = v56
					v60 = base.F64_sub(v9, base.F64_reinterpret_i64(v56))
					v66 = v60
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v14))) = v9
					v66 = base.F64_reinterpret_i64(v19 & int64(-9223372036854775807-1))
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v14))) = v19 & int64(-9223372036854775807-1)
				v66 = v9
			}
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v14))) = v9
			if v19&int64(4503599627370495) == int64(0) {
				v66 = base.F64_reinterpret_i64(v19 & int64(-9223372036854775807-1))
			} else {
				if v26 == int32(1024) {
					v60 = v9
					v66 = v60
				} else {
					v66 = base.F64_reinterpret_i64(v19 & int64(-9223372036854775807-1))
				}
			}
		}
		v67 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v69))) = v67
		v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73 + int32(16)
		v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v78))) = v66
		v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82 + int32(16)
		m.G0 = v6 + int32(16)
		return int32(2)
	}
}
func F_math_pow(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	v3 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_luaL_checknumber(m, l0, int32(2))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_pow(m, v3, v8)
			mBase = m.M
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v12))) = v10
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v16 + int32(16)
			return int32(1)
		}
	}
}
func F_math_random(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int64
	_ = v8
	var v12 int64
	_ = v12
	var v18 int32
	_ = v18
	var v21 float64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	v3 = int32(0)
	v8 = *(*int64)(unsafe.Add(mBase, _consts[298]))
	v12 = v8*int64(6364136223846793005) + int64(1)
	*(*int64)(unsafe.Add(mBase, _consts[298])) = v12
	v18 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v12)>>(uint(int64(33))%64))), int32(2147483647))
	v21 = base.F64_div(base.F64_convert_i32_s(v18), float64(2.147483647e+09))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	switch (v22 - v23) >> (uint(int32(4)) % 32) {
	case 0:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v28))) = v21
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32 + int32(16)
		return int32(1)
	case 1:
		v39 = F_luaL_checkinteger(m, l0, int32(1))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			if int32(0) < v39 {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(3)
				*(*float64)(unsafe.Add(mBase, uint32(v57))) = base.F64_add(base.F64_floor(base.F64_mul(v21, base.F64_convert_i32_s(v39))), float64(1))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v61 + int32(16)
				return int32(1)
			} else {
				v46 = m.G3
				v49 = F_luaL_argerror(m, l0, int32(1), v46+int32(_a2320))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v57))) = base.F64_add(base.F64_floor(base.F64_mul(v21, base.F64_convert_i32_s(v39))), float64(1))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v61 + int32(16)
					return int32(1)
				}
			}
		}
	case 2:
		v68 = F_luaL_checkinteger(m, l0, int32(1))
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return int32(0)
		} else {
			v71 = F_luaL_checkinteger(m, l0, int32(2))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				if v68 <= v71 {
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v89))) = base.F64_add(base.F64_floor(base.F64_mul(v21, base.F64_convert_i32_s(v71-v68+int32(1)))), base.F64_convert_i32_s(v68))
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v93 + int32(16)
					return int32(1)
				} else {
					v75 = m.G3
					v78 = F_luaL_argerror(m, l0, int32(2), v75+int32(_a2320))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(v89))) = base.F64_add(base.F64_floor(base.F64_mul(v21, base.F64_convert_i32_s(v71-v68+int32(1)))), base.F64_convert_i32_s(v68))
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v93 + int32(16)
						return int32(1)
					}
				}
			}
		}
	default:
		v99 = m.G3
		v103 = F_luaL_error(m, l0, v99+int32(_a2214), int32(0))
		mBase = m.M
		v104 = m.ExcPending
		if v104 != 0 {
			return int32(0)
		} else {
			return v103
		}
	}
}
func F_math_sqrt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v3 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v9))) = base.F64_sqrt(v3)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13 + int32(16)
		return int32(1)
	}
}
func F_math_tan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v25 float64
	_ = v25
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	v3 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v9 = m.G0
		v11 = v9 - int32(16)
		m.G0 = v11
		v18 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v3))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1072243195)) < base.Ui32(v18) {
			if base.Ui32(v18) < base.Ui32(int32(2146435072)) {
				v29 = F___rem_pio2(m, v3, v11)
				mBase = m.M
				v30 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
				v31 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v34 = F___tan(m, v30, v31, v29&int32(1))
				mBase = m.M
				v35 = v34
			} else {
				v35 = base.F64_sub(v3, v3)
			}
		} else {
			if base.Ui32(v18) < base.Ui32(int32(1044381696)) {
				v35 = v3
			} else {
				v25 = F___tan(m, v3, float64(0), int32(0))
				mBase = m.M
				v35 = v25
			}
		}
		m.G0 = v11 + int32(16)
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v41))) = v35
		v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
		return int32(1)
	}
}
