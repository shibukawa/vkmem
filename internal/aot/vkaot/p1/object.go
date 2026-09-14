package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_createObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	v11 = int32(12)
	v14 = F_zmalloc_usable(m, v11, v7+v11)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0 & int32(15)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(8)
		m.G0 = v7 + int32(16)
		return v14
	}
}
func F_initObjectLRUOrLFU(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(int32(-9)) < base.Ui32(v3) {
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[378])))
		if v7 != int32(1) {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[379]))
			v20 = v17 & int32(16777215)
		} else {
			v11 = int32(*(*uint16)(unsafe.Add(mBase, _consts[376])))
			v20 = v11<<(uint(int32(8))%32) | int32(5)
		}
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v21 | v20<<(uint(int32(8))%32)
	}
	return
}
func F_makeObjectShared(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3&int32(-8) == int32(8) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3 | int32(-8)
		return l0
	} else {
		F__serverAssert(m, int32(_a1757), int32(_a1758), int32(132))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_objectGetLFUFrequency(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = int32(base.Ui32(v8) >> (uint(int32(8)) % 32))
	v17 = int32(*(*uint16)(unsafe.Add(mBase, _consts[376])))
	v20 = *(*int32)(unsafe.Add(mBase, _consts[377]))
	if v20 == v2 {
		v31 = v2
	} else {
		v26 = int32(65535)
		v28 = base.I32_div_s((v17-int32(base.Ui32(v10)>>(uint(int32(8))%32)))&v26, v20)
		v31 = v28 & v26
	}
	v34 = v10 & int32(255)
	v35 = v34 - v31
	if base.Ui32(v34) < base.Ui32(v35) {
		v37 = int32(0)
	} else {
		v37 = v35
	}
	*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(15)))) = uint8(v37)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v42 | (v37|v17<<(uint(int32(8))%32))<<(uint(int32(8))%32)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
	m.G0 = v6 + int32(16)
	return v47
}
func F_tryObjectEncoding(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_tryObjectEncodingEx(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
