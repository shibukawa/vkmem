package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_evalCalcScriptHash(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	if l0 == int32(0) {
		v54 = int32(0)
		v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
		switch v58 & int32(7) {
		case 0:
			v75 = int32(base.Ui32(v58) >> (uint(int32(3)) % 32))
		case 1:
			v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
			v75 = v65
		case 2:
			v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
			v75 = v68
		case 3:
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
			v75 = v71
		case 4:
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
			v75 = v74
		default:
			v75 = v54
		}
		v77 = v9 + int32(20)
		F_SHA1Init(m, v77)
		mBase = m.M
		F_SHA1Update(m, v77, l1, v75)
		mBase = m.M
		F_SHA1Final(m, v9, v77)
		mBase = m.M
		v85 = v54
		for {
			v91 = int32(1)
			v93 = l2 + v85<<(uint(v91)%32)
			v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v85))))
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97&int32(15))+uint32(_consts[310]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v93+v91))) = uint8(v102)
			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v97)>>(uint(int32(4))%32)))+uint32(_consts[310]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v108)
			v111 = v85 + v91
			if v111 != int32(20) {
				v85 = v111
				continue
			} else {
				break
			}
			break
		}
		v114 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2)+40)) = uint8(v114)
	} else {
		v14 = int32(0)
		for {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v14))))
			if base.Ui32((v22+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v31 = v22 + int32(32)
			} else {
				v31 = v22
			}
			*(*uint8)(unsafe.Add(mBase, uint32(l2+v14))) = uint8(v31)
			v34 = v14 | int32(1)
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v34))))
			if base.Ui32((v37+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v46 = v37 + int32(32)
			} else {
				v46 = v37
			}
			*(*uint8)(unsafe.Add(mBase, uint32(l2+v34))) = uint8(v46)
			v49 = v14 + int32(2)
			if v49 != int32(40) {
				v14 = v49
				continue
			} else {
				break
			}
			break
		}
		v52 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2)+40)) = uint8(v52)
	}
	m.G0 = v9 + int32(112)
	return
}
func F_scriptFlagsToCmdFlags(m *base.Module, l0 int64, l1 int64) int64 {
	var v11 int64
	_ = v11
	var v18 int64
	_ = v18
	v11 = l0 & int64(-66566)
	if l1&int64(3) == int64(0) {
		v18 = v11 | int64(4)
	} else {
		v18 = v11
	}
	return l1<<(uint(int64(8))%64)&int64(1024) | l1&int64(1) | v18 ^ int64(1)
}
func F_scriptGetSlot(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	if v3 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
		return v12
	} else {
		F__serverAssert(m, int32(_a2020), int32(_a2021), int32(346))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
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
func F_scriptIsRunning(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	v1 = int32(0)
	v2 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	return base.B2i32(v2 != v1)
}
func F_scriptIsTimedout(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	v1 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	if v5 == v1 {
		v13 = v1
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
		v13 = int32(base.Ui32(v8)>>(uint(int32(3))%32)) & int32(1)
	}
	return v13
}
func F_scriptSetSlot(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	if v4 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = l0
		return
	} else {
		F__serverAssert(m, int32(_a2020), int32(_a2021), int32(351))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
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
