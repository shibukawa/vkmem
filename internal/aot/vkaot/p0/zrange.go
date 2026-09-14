package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zrangeResultBeginStore(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	v3 = int32(0)
	if v3 < l1 {
		v7 = l1
	} else {
		v7 = v3
	}
	v9 = *(*int32)(unsafe.Add(mBase, _consts[1103]))
	if base.Ui32(v9) < base.Ui32(v7) {
		v14 = F_createZsetObject(m)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = F_objectGetVal(m, v14)
			mBase = m.M
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v18 = F_hashtableExpand(m, v17, v7)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v14
				return
			}
		}
	} else {
		v11 = F_createZsetListpackObject(m)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v11
			return
		}
	}
}
func F_zrangeResultEmitLongLongForStore(m *base.Module, l0 int32, l1 int64, l2 float64) {
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
	v12 = F_sdsfromlonglong(m, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v20 = F_zsetAdd(m, v14, l2, v12, int32(0), v8+int32(4), v8+int32(8))
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
					F__serverAssert(m, int32(_a1627), int32(_a1609), int32(3046))
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
func F_zrangeResultFinalizeStore(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+96))
	if l1 == int32(0) {
		v34 = F_dbDelete(m, v8, v6)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			if v34 == int32(0) {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v60 = *(*int32)(unsafe.Add(mBase, _consts[233]))
				F_addReply(m, v58, v60)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					F_decrRefCount(m, v63)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+96))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				F_signalModifiedKey(m, v38, v39, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
					F_notifyKeyspaceEvent(m, int32(4), int32(_a132), v45, v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v51 = int32(_a69)
						v53 = *(*int64)(unsafe.Add(mBase, _consts[60]))
						*(*int64)(unsafe.Add(mBase, _consts[60])) = v53 + int64(1)
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v60 = *(*int32)(unsafe.Add(mBase, _consts[233]))
						F_addReply(m, v58, v60)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							F_decrRefCount(m, v63)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_setKey(m, v7, v8, v6, l0+int32(12), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
			F_notifyKeyspaceEvent(m, int32(128), int32(_a1628), v18, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = int32(_a69)
				v26 = *(*int64)(unsafe.Add(mBase, _consts[60]))
				*(*int64)(unsafe.Add(mBase, _consts[60])) = v26 + int64(1)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_addReplyLongLong(m, v30, base.I64_extend_i32_u(l1))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
