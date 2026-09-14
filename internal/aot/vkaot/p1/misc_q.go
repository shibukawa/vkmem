package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_qsort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v7 int32
	_ = v7
	F___qsort_r(m, l0, l1, l2, int32(1381), l3)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_qsortCompareSetsByCardinality(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_setTypeSize(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v11 = F_setTypeSize(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if base.Ui32(v11) < base.Ui32(v6) {
				v24 = int32(1)
				return v24
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v17 = F_setTypeSize(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v20 = F_setTypeSize(m, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						if base.Ui32(v17) < base.Ui32(v20) {
							v23 = int32(-1)
						} else {
							v23 = int32(0)
						}
						v24 = v23
						return v24
					}
				}
			}
		}
	}
}
func F_qsortCompareSetsByRevCardinality(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 == v3 {
		v15 = v3
		if v5 == int32(0) {
			v20 = v3
			return base.B2i32(base.Ui32(v15) < base.Ui32(v20)) - base.B2i32(base.Ui32(v20) < base.Ui32(v15))
		} else {
			v18 = F_setTypeSize(m, v5)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = v18
				return base.B2i32(base.Ui32(v15) < base.Ui32(v20)) - base.B2i32(base.Ui32(v20) < base.Ui32(v15))
			}
		}
	} else {
		v11 = F_setTypeSize(m, v8)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			if v5 == int32(0) {
				v20 = v3
				return base.B2i32(base.Ui32(v15) < base.Ui32(v20)) - base.B2i32(base.Ui32(v20) < base.Ui32(v15))
			} else {
				v18 = F_setTypeSize(m, v5)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = v18
					return base.B2i32(base.Ui32(v15) < base.Ui32(v20)) - base.B2i32(base.Ui32(v20) < base.Ui32(v15))
				}
			}
		}
	}
}
