package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_sdslen_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v7 & int32(7) {
	case 0:
		return int32(base.Ui32(v7) >> (uint(int32(3)) % 32))
	case 1:
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		return v15
	case 2:
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		return v19
	case 3:
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		return v23
	case 4:
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v28 = v27
		return v28
	default:
		v28 = int32(0)
		return v28
	}
}
func F_sdslen_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v7 & int32(7) {
	case 0:
		return int32(base.Ui32(v7) >> (uint(int32(3)) % 32))
	case 1:
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		return v15
	case 2:
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		return v19
	case 3:
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		return v23
	case 4:
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v28 = v27
		return v28
	default:
		v28 = int32(0)
		return v28
	}
}
func F_sdslen_4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v7 & int32(7) {
	case 0:
		return int32(base.Ui32(v7) >> (uint(int32(3)) % 32))
	case 1:
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		return v15
	case 2:
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		return v19
	case 3:
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		return v23
	case 4:
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v28 = v27
		return v28
	default:
		v28 = int32(0)
		return v28
	}
}
