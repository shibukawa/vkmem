package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zmalloc_default_oom(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_zmalloc_default_oom[0]))
	v11 = F_fiprintf(m, v9, int32(_a_F_zmalloc_default_oom_0), v5)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = F_fflush(m, v9)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
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
func F_zmalloc_get_allocator_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v6
	if l3 == v6 {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	}
	if l4 == int32(0) {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	}
	return int32(1)
}
func F_zmalloc_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-8))))
	return v4&int32(2147483647) + int32(8)
}
func F_zmalloc_usable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	if base.Ui32(int32(2147483646)) < base.Ui32(l0) {
		v49 = int32(0)
		v51 = *(*int32)(unsafe.Add(mBase, _c_F_zmalloc_usable[0]))
		m.T0[v51].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			v58 = int32(0)
			v60 = v49
			if l1 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v58
			}
			return v60
		}
	} else {
		if l0 != 0 {
			v9 = l0
		} else {
			v9 = int32(4)
		}
		v11 = v9 + int32(8)
		v12 = F_emscripten_builtin_malloc(m, v11)
		mBase = m.M
		if v12 == int32(0) {
			v49 = int32(0)
			v51 = *(*int32)(unsafe.Add(mBase, _c_F_zmalloc_usable[0]))
			m.T0[v51].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				v58 = int32(0)
				v60 = v49
				if l1 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v58
				}
				return v60
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v9
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_zmalloc_usable[1]))
			if v17 != int32(-1) {
				v28 = v17
			} else {
				v20 = int32(0)
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_zmalloc_usable[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_zmalloc_usable[1])) = v22
				*(*int32)(unsafe.Add(mBase, _c_F_zmalloc_usable[2])) = v22 + int32(1)
				v28 = v22
			}
			if v28 < int32(260) {
				v37 = v28 << (uint(int32(2)) % 32)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_zmalloc_usable[3])))
				*(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_zmalloc_usable[3]))) = v40 + v11
			} else {
				v31 = int32(0)
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_zmalloc_usable[4]))
				*(*int32)(unsafe.Add(mBase, _c_F_zmalloc_usable[4])) = v33 + v11
			}
			v58 = v9
			v60 = v12 + int32(8)
			if l1 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v58
			}
			return v60
		}
	}
}
