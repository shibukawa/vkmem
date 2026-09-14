package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_sdsAllocPtr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v9 = v7 & int32(7)
	if base.Ui32(int32(4)) < base.Ui32(v9) {
		v17 = int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v9<<(uint(int32(2))%32))+uint32(_c_F_sdsAllocPtr[0])))
		v17 = v16
	}
	return l0 + v17
}
func F_sdsConfigGet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v3&int32(1) == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = v14
		if v15 == int32(0) {
			v25 = F_sdsnew(m, int32(_a_F_sdsConfigGet_0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = v25
				return v27
			}
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
			if v18&int32(1) != 0 {
				v27 = v15
				return v27
			} else {
				v21 = F_sdsdup(m, v15)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return v21
				}
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		v9 = F_getModuleStringConfig(m, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v15 = v9
			if v15 == int32(0) {
				v25 = F_sdsnew(m, int32(_a_F_sdsConfigGet_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = v25
					return v27
				}
			} else {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
				if v18&int32(1) != 0 {
					v27 = v15
					return v27
				} else {
					v21 = F_sdsdup(m, v15)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						return v21
					}
				}
			}
		}
	}
}
func F_sdsHdrSize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v5 = l0 & int32(7)
	if base.Ui32(int32(4)) < base.Ui32(v5) {
		v13 = int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5<<(uint(int32(2))%32))+uint32(_c_F_sdsHdrSize[0])))
		v13 = v12
	}
	return v13
}
func F_sdsMakeRoomFor(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F__sdsMakeRoomFor(m, l0, l1, int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
