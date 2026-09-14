package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F___math_divzero(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v6 int32
	_ = v6
	if l0 != 0 {
		v4 = float64(-1)
	} else {
		v4 = float64(1)
	}
	v6 = m.G0
	*(*float64)(unsafe.Add(mBase, uint32(v6-int32(16))+8)) = v4
	return base.F64_div(v4, float64(0))
}
func F_math_ceil(m *base.Module, l0 int32) int32 {
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
		*(*float64)(unsafe.Add(mBase, uint32(v9))) = base.F64_ceil(v3)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13 + int32(16)
		return int32(1)
	}
}
func F_math_cosh(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v9 float64
	_ = v9
	var v10 int64
	_ = v10
	var v16 float64
	_ = v16
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v26 float64
	_ = v26
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v37 float64
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v3 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v9 = base.F64_abs(v3)
		v10 = base.I64_reinterpret_f64(v9)
		if base.Ui64(int64(4604418530035630079)) < base.Ui64(v10) {
			if base.Ui64(int64(4649454526309335039)) < base.Ui64(v10) {
				v33 = F___expo2(m, v9, float64(1))
				mBase = m.M
				v34 = v33
				v37 = v34
			} else {
				v26 = F_exp(m, v9)
				mBase = m.M
				v37 = base.F64_mul(base.F64_add(v26, base.F64_div(float64(1), v26)), float64(0.5))
			}
		} else {
			if base.Ui64(v10) < base.Ui64(int64(4490088828488384512)) {
				v34 = float64(1)
				v37 = v34
			} else {
				v16 = F_expm1(m, v9)
				mBase = m.M
				v18 = float64(1)
				v19 = base.F64_add(v16, v18)
				v37 = base.F64_add(base.F64_div(base.F64_mul(v16, v16), base.F64_add(v19, v19)), v18)
			}
		}
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v39))) = v37
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v43 + int32(16)
		return int32(1)
	}
}
func F_math_floor(m *base.Module, l0 int32) int32 {
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
		*(*float64)(unsafe.Add(mBase, uint32(v9))) = base.F64_floor(v3)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13 + int32(16)
		return int32(1)
	}
}
func F_math_frexp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 float64
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 float64
	_ = v29
	var v30 int32
	_ = v30
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v44 float64
	_ = v44
	var v47 float64
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = F_luaL_checknumber(m, l0, int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = v5 + int32(12)
		v16 = base.I64_reinterpret_f64(v8)
		v20 = int32(2047)
		v21 = base.I32_wrap_i64(int64(base.Ui64(v16)>>(uint(int64(52))%64))) & v20
		if v21 == v20 {
			v44 = v8
			v47 = v44
		} else {
			if v21 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v21 + int32(-1022)
				v44 = base.F64_reinterpret_i64(v16&int64(-9218868437227405313) | int64(4602678819172646912))
				v47 = v44
			} else {
				if base.F64_ne(v8, float64(0)) != 0 {
					v29 = F_frexp(m, base.F64_mul(v8, float64(1.8446744073709552e+19)), v13)
					mBase = m.M
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v33 = v29
					v34 = v30 + int32(-64)
				} else {
					v33 = v8
					v34 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v34
				v47 = v33
			}
		}
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v49))) = v47
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v53 + int32(16)
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v59))) = base.F64_convert_i32_s(v57)
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v64 + int32(16)
		m.G0 = v5 + int32(16)
		return int32(2)
	}
}
func F_math_rad(m *base.Module, l0 int32) int32 {
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
		*(*float64)(unsafe.Add(mBase, uint32(v10))) = base.F64_mul(v3, float64(0.017453292519943295))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v14 + int32(16)
		return int32(1)
	}
}
func F_math_sinh(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v13 int64
	_ = v13
	var v16 float64
	_ = v16
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v39 float64
	_ = v39
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
		v11 = base.F64_copysign(float64(0.5), v3)
		v12 = base.F64_abs(v3)
		v13 = base.I64_reinterpret_f64(v12)
		if base.Ui64(int64(4649454526309335039)) < base.Ui64(v13) {
			v34 = F___expo2(m, v12, base.F64_add(v11, v11))
			mBase = m.M
			v35 = v34
			v39 = v35
		} else {
			v16 = F_expm1(m, v12)
			mBase = m.M
			if base.Ui64(int64(4607182418800017407)) < base.Ui64(v13) {
				v39 = base.F64_mul(v11, base.F64_add(v16, base.F64_div(v16, base.F64_add(v16, float64(1)))))
			} else {
				if base.Ui64(v13) < base.Ui64(int64(4490088828488384512)) {
					v35 = v3
					v39 = v35
				} else {
					v39 = base.F64_mul(v11, base.F64_sub(base.F64_add(v16, v16), base.F64_div(base.F64_mul(v16, v16), base.F64_add(v16, float64(1)))))
				}
			}
		}
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v41))) = v39
		v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
		return int32(1)
	}
}
func F_math_tanh(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v9 float64
	_ = v9
	var v10 int64
	_ = v10
	var v20 float64
	_ = v20
	var v22 float64
	_ = v22
	var v30 float64
	_ = v30
	var v38 float64
	_ = v38
	var v43 float64
	_ = v43
	var v48 float64
	_ = v48
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
		v9 = base.F64_abs(v3)
		v10 = base.I64_reinterpret_f64(v9)
		if base.Ui64(v10) < base.Ui64(int64(4603122931675955200)) {
			if base.Ui64(v10) < base.Ui64(int64(4598272728187797504)) {
				if base.Ui64(v10) < base.Ui64(int64(4503599627370496)) {
					v43 = v9
				} else {
					v38 = F_expm1(m, base.F64_mul(v9, float64(-2)))
					mBase = m.M
					v43 = base.F64_div(base.F64_neg(v38), base.F64_add(v38, float64(2)))
				}
			} else {
				v30 = F_expm1(m, base.F64_add(v9, v9))
				mBase = m.M
				v43 = base.F64_div(v30, base.F64_add(v30, float64(2)))
			}
		} else {
			if base.Ui64(v10) < base.Ui64(int64(4626322721511309312)) {
				v20 = float64(2)
				v22 = F_expm1(m, base.F64_add(v9, v9))
				mBase = m.M
				v43 = base.F64_sub(float64(1), base.F64_div(v20, base.F64_add(v22, v20)))
			} else {
				v43 = base.F64_add(base.F64_div(math.Float64frombits(uint64(0x8000000000000000)), v9), float64(1))
			}
		}
		if base.I64_reinterpret_f64(v3) < int64(0) {
			v48 = base.F64_neg(v43)
		} else {
			v48 = v43
		}
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v50))) = v48
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v54 + int32(16)
		return int32(1)
	}
}
