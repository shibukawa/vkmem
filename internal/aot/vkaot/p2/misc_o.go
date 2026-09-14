package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___ofl_unlock(m *base.Module) {
	return
}
func F_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1&int32(64) != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2 + int32(4)
		v21 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2))))
		v22 = v21
	} else {
		v13 = int32(4259840)
		if l1&v13 != v13 {
			v22 = int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2 + int32(4)
			v21 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2))))
			v22 = v21
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v22
	v27 = m.Env.X__syscall_openat(m, int32(-100), l0, l1|int32(32768), v8)
	mBase = m.M
	if base.Ui32(v27) < base.Ui32(int32(-4095)) {
		v35 = v27
	} else {
		v30 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(0) - v27
		v35 = int32(-1)
	}
	m.G0 = v8 + int32(16)
	return v35
}
func F_open_func(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v7 = F_luaF_newproto(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v7
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v6
		*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(l1)+28)) = int64(-1)
		v17 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l1)+20)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(l1)+36)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(l1+int32(43)))) = v17
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
		v26 = int32(2)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+75)) = uint8(v26)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v25
		v29 = int32(0)
		v31 = F_luaH_new(m, v6, v29, v29)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v31
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = int32(5)
			*(*int32)(unsafe.Add(mBase, uint32(v34))) = v31
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			if int32(16) < v38-v39 {
				v47 = v39
				v48 = int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v47 + v48
				*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = int32(9)
				*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v7
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				if v48 < v54-v55 {
					v63 = v55
					*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v63 + int32(16)
					return
				} else {
					F_luaD_growstack(m, v6, int32(1))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						v63 = v62
						*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v63 + int32(16)
						return
					}
				}
			} else {
				F_luaD_growstack(m, v6, int32(1))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
					v47 = v46
					v48 = int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v47 + v48
					*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = int32(9)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v7
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
					if v48 < v54-v55 {
						v63 = v55
						*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v63 + int32(16)
						return
					} else {
						F_luaD_growstack(m, v6, int32(1))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
							v63 = v62
							*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v63 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_os_clock(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(-1)
	v12 = F___clock_gettime(m, int32(2), v8)
	mBase = m.M
	if v12 != 0 {
		v27 = v10
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
		if int64(2147) < v13 {
			v27 = v10
		} else {
			v18 = v13 * int64(1000000)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			v22 = base.I32_div_s(v20, int32(1000))
			if int64(2147483647)-v18 < base.I64_extend_i32_s(v22) {
				v27 = v10
			} else {
				v27 = v22 + base.I32_wrap_i64(v18)
			}
		}
	}
	m.G0 = v8 + int32(16)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v37))) = base.F64_div(base.F64_convert_i32_s(v27), float64(1e+06))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41 + int32(16)
	return int32(1)
}
