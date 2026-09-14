package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_callbackKeyDestructor(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_sdsfree(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_mapEndCallback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_lua_settable(m, v2, int32(-3))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6 + int32(-1)
		F_processCollectionElementEnd(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_mapStartCallback(m *base.Module, l0 int32, l1 int32) {
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
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
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
			v40 = int32(0)
			F_lua_createtable(m, v31, v40, v40)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				v44 = m.G3
				F_lua_pushstring(m, v31, v44+int32(_a_F_mapStartCallback_0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					F_lua_createtable(m, v31, int32(0), l1)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v54 = v52 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v54
						v56 = int32(2)
						v58 = l0 + v54<<(uint(v56)%32)
						*(*int32)(unsafe.Add(mBase, uint32(v58+int32(1040)))) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v58+int32(16)))) = int32(0)
						return
					}
				}
			}
		} else {
			F__serverPanic_2(m, int32(976))
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
