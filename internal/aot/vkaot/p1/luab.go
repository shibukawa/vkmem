package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaB_collectgarbage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	v4 = m.G3
	v12 = F_luaL_checkoption(m, l0, int32(1), v4+int32(_a2696), v4+int32(_a2697))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(2)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v4+int32(_a2698)+v12<<(uint(v16)%32))))
		v22 = F_luaL_optinteger(m, l0, v16, int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_lua_gc(m, l0, v19, v22)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				switch v19 + int32(-3) {
				case 0:
					v30 = F_lua_gc(m, l0, int32(4), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(v38))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v30), float64(0.0009765625)), base.F64_convert_i32_s(v24))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v42 + int32(16)
						return int32(1)
					}
				default:
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v63))) = base.F64_convert_i32_s(v24)
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + int32(16)
					return int32(1)
				case 2:
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v49))) = base.B2i32(v24 != int32(0))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v55 + int32(16)
					return int32(1)
				}
			}
		}
	}
}
func F_luaB_dofile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v2 = int32(0)
	v7 = F_luaL_optlstring(m, l0, int32(1), v2, v2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v16 = F_luaL_loadfile(m, l0, v7)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v16 == int32(0) {
				F_lua_call(m, l0, int32(0), int32(-1))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					return (v26-v27)>>(uint(int32(4))%32) - (v11-v12)>>(uint(int32(4))%32)
				}
			} else {
				v20 = F_lua_error(m, l0)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_lua_call(m, l0, int32(0), int32(-1))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						return (v26-v27)>>(uint(int32(4))%32) - (v11-v12)>>(uint(int32(4))%32)
					}
				}
			}
		}
	}
}
