package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___math_invalid(m *base.Module, l0 float64) float64 {
	var v2 float64
	_ = v2
	v2 = base.F64_sub(l0, l0)
	return base.F64_div(v2, v2)
}
func F___math_uflow(m *base.Module, l0 int32) float64 {
	var v2 float64
	_ = v2
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	v2 = float64(1.2882297539194267e-231)
	if l0 != 0 {
		v4 = base.F64_neg(v2)
	} else {
		v4 = v2
	}
	v5 = F_fp_barrier_1(m, v4)
	return base.F64_mul(v2, v5)
}
func F___math_xflow(m *base.Module, l0 int32, l1 float64) float64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v6 int32
	_ = v6
	if l0 != 0 {
		v4 = base.F64_neg(l1)
	} else {
		v4 = l1
	}
	v6 = m.G0
	*(*float64)(unsafe.Add(mBase, uint32(v6-int32(16))+8)) = v4
	return base.F64_mul(l1, v4)
}
func F_math_abs(m *base.Module, l0 int32) int32 {
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
		*(*float64)(unsafe.Add(mBase, uint32(v9))) = base.F64_abs(v3)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13 + int32(16)
		return int32(1)
	}
}
func F_math_asin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v12 int64
	_ = v12
	var v17 int32
	_ = v17
	var v38 float64
	_ = v38
	var v42 float64
	_ = v42
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v52 float64
	_ = v52
	var v57 float64
	_ = v57
	var v61 float64
	_ = v61
	var v70 float64
	_ = v70
	var v77 float64
	_ = v77
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v91 float64
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	v3 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v12 = base.I64_reinterpret_f64(v3)
		v17 = base.I32_wrap_i64(int64(base.Ui64(v12)>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v17) < base.Ui32(int32(1072693248)) {
			if base.Ui32(int32(1071644671)) < base.Ui32(v17) {
				v42 = F_fabs(m, v3)
				mBase = m.M
				v45 = base.F64_mul(base.F64_sub(float64(1), v42), float64(0.5))
				v46 = F_sqrt(m, v45)
				mBase = m.M
				v47 = F_R_2(m, v45)
				mBase = m.M
				if base.Ui32(v17) < base.Ui32(int32(1072640819)) {
					v57 = float64(0.7853981633974483)
					v61 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v46) & int64(-4294967296))
					v70 = base.F64_div(base.F64_sub(v45, base.F64_mul(v61, v61)), base.F64_add(v46, v61))
					v77 = base.F64_add(base.F64_sub(base.F64_sub(v57, base.F64_add(v61, v61)), base.F64_sub(base.F64_mul(base.F64_add(v46, v46), v47), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v70, v70)))), v57)
				} else {
					v52 = base.F64_add(base.F64_mul(v46, v47), v46)
					v77 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v52, v52), float64(-6.123233995736766e-17)))
				}
				if v12 < int64(0) {
					v82 = base.F64_neg(v77)
				} else {
					v82 = v77
				}
				v83 = v82
				v91 = v83
			} else {
				if base.Ui32(v17+int32(-1048576)) < base.Ui32(int32(1044381696)) {
					v83 = v3
					v91 = v83
				} else {
					v38 = F_R_2(m, base.F64_mul(v3, v3))
					mBase = m.M
					v91 = base.F64_add(base.F64_mul(v3, v38), v3)
				}
			}
		} else {
			if v17+int32(-1072693248)|base.I32_wrap_i64(v12) != 0 {
				v91 = base.F64_div(float64(0), base.F64_sub(v3, v3))
			} else {
				v91 = base.F64_add(base.F64_mul(v3, float64(1.5707963267948966)), float64(7.52316384526264e-37))
			}
		}
		v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v93))) = v91
		v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v97 + int32(16)
		return int32(1)
	}
}
func F_math_atan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
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
		v7 = F_atan(m, v3)
		mBase = m.M
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v9))) = v7
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13 + int32(16)
		return int32(1)
	}
}
func F_math_fmod(m *base.Module, l0 int32) int32 {
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
			v10 = F_fmod(m, v3, v8)
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
func F_math_ldexp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
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
		v8 = F_luaL_checkinteger(m, l0, int32(2))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_scalbn(m, v3, v8)
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
func F_math_log(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
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
		v7 = F_log(m, v3)
		mBase = m.M
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v9))) = v7
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13 + int32(16)
		return int32(1)
	}
}
func F_math_randomseed(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_luaL_checkinteger(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, _consts[343])) = base.I64_extend_i32_u(v3 + int32(-1))
		return int32(0)
	}
}
func F_math_sin(m *base.Module, l0 int32) int32 {
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
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
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
					v36 = F___sin(m, v32, v31, int32(1))
					mBase = m.M
					v43 = v36
				case 1:
					v37 = F___cos(m, v32, v31)
					mBase = m.M
					v43 = v37
				case 2:
					v39 = F___sin(m, v32, v31, int32(1))
					mBase = m.M
					v43 = base.F64_neg(v39)
				case 3:
					v41 = F___cos(m, v32, v31)
					mBase = m.M
					v43 = base.F64_neg(v41)
				}
			} else {
				v43 = base.F64_sub(v3, v3)
			}
		} else {
			if base.Ui32(v19) < base.Ui32(int32(1045430272)) {
				v43 = v3
			} else {
				v26 = F___sin(m, v3, float64(0), int32(0))
				mBase = m.M
				v43 = v26
			}
		}
		m.G0 = v12 + int32(16)
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v50))) = v43
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v54 + int32(16)
		return int32(1)
	}
}
