package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zrangeResultEmitCBufferForStore(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(0)
	v12 = F_sdsnewlen(m, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v20 = F_zsetAdd(m, v14, l3, v12, int32(0), v8+int32(4), v8+int32(8))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			F_sdsfree(m, v12)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				if v20 != 0 {
					m.G0 = v8 + int32(16)
					return
				} else {
					F__serverAssert(m, int32(_a1107), int32(_a1723), int32(3037))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
func F_zrangeResultEmitLongLongToClient(m *base.Module, l0 int32, l1 int64, l2 float64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v4 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_addReplyBulkLongLong(m, v11, l1)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v14 == int32(0) {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_addReplyDouble(m, v17, l2)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_addReplyArrayLen(m, v7, int32(2))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			F_addReplyBulkLongLong(m, v11, l1)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v14 == int32(0) {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_addReplyDouble(m, v17, l2)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_zrangeResultFinalizeClient(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v5 == int32(0) {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v9 == int32(0) {
			v16 = l1
		} else {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+224)))
			v16 = l1 << (uint(base.B2i32(v12 == int32(2))) % 32)
		}
		F_setDeferredArrayLen(m, v8, v5, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			return
		}
	}
}
