package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___DOUBLE_BITS_3(m *base.Module, l0 float64) int64 {
	return base.I64_reinterpret_f64(l0)
}
func F___DOUBLE_BITS_5(m *base.Module, l0 float64) int64 {
	return base.I64_reinterpret_f64(l0)
}
func F_createDoubleObject(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	if l3 != int32(-1) {
		v13 = m.G4
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v15 = m.T0[v14].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(48))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(7)
				v26 = m.G4
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				v28 = m.T0[v27].(func(*base.Module, int32) int32)(m, l3+int32(1))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v28
					if v28 != 0 {
						if l3 == int32(0) {
						} else {
							v37 = F__emscripten_memcpy_bulkmem(m, v28, l2, l3)
							mBase = m.M
						}
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
						v41 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v39+l3))) = uint8(v41)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l3
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v44 == v41 {
							return v15
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
							if base.Ui32(v48+int32(-9)) < base.Ui32(int32(4)) {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+40))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v55+v56<<(uint(int32(2))%32)))) = v15
								return v15
							} else {
								if v48 != int32(2) {
									v64 = m.G3
									m.Env.X__assert_fail(m, v64+int32(_a1917), v64+int32(_a1910), int32(250), v64+int32(_a1918))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+40))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v55+v56<<(uint(int32(2))%32)))) = v15
									return v15
								}
							}
						}
					} else {
						F_freeReplyObject(m, v15)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				}
			} else {
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func F_getDoubleFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 float64
	_ = v27
	var v30 int32
	_ = v30
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_getDoubleFromObject(m, l1, v8+int32(8))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v27 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			*(*float64)(unsafe.Add(mBase, uint32(l2))) = v27
			v30 = int32(0)
			m.G0 = v8 + int32(16)
			return v30
		} else {
			if l3 == int32(0) {
				F_addReplyError(m, l0, int32(_a1060))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v30 = int32(-1)
					m.G0 = v8 + int32(16)
					return v30
				}
			} else {
				F_addReplyError(m, l0, l3)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v30 = int32(-1)
					m.G0 = v8 + int32(16)
					return v30
				}
			}
		}
	}
}
func F_trimDoubleString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	v5 = int32(46)
	v6 = F___strchrnul(m, l0, v5)
	mBase = m.M
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v8 == v5 {
		v12 = v6
	} else {
		v12 = int32(0)
	}
	if v12 == int32(0) {
		v32 = l1
	} else {
		v17 = l1
		v18 = l0 + l1
		for {
			v21 = v18 + int32(-1)
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
			if v22 == int32(48) {
				v17 = v17 + int32(-1)
				v18 = v21
				continue
			} else {
				break
			}
			break
		}
		if v22 != int32(46) {
			v32 = v17
		} else {
			v32 = v17 + int32(-1)
		}
	}
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v32))) = uint8(v36)
	return v32
}
