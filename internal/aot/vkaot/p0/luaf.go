package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaF_findupval(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = l0 + int32(104)
	goto L3
L1:
	;
	return v61
L2:
	;
	v36 = int32(0)
	v39 = F_luaM_realloc_(m, l0, v36, v36, int32(32))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v16 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L4:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+5)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
	if v22&(v23^int32(-1))&int32(3) == int32(0) {
		v61 = v16
		goto L1
	} else {
		goto L8
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if base.Ui32(v19) < base.Ui32(l1) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if v19 != l1 {
		v12 = v16
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	v32 = v22 ^ int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+5)) = uint8(v32)
	return v16
L9:
	;
	return int32(0)
L10:
	;
	v43 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)) = uint8(v43)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = l1
	v48 = v45 & int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+5)) = uint8(v48)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v9 + int32(120)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v56)+16)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v9)+140)) = v39
	v61 = v39
	goto L1
}
func F_luaF_freeproto(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v8 = F_luaM_realloc_(m, l0, v3, v4<<(uint(int32(2))%32), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
		v15 = F_luaM_realloc_(m, l0, v10, v11<<(uint(int32(2))%32), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			v22 = F_luaM_realloc_(m, l0, v17, v18<<(uint(int32(4))%32), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				v29 = F_luaM_realloc_(m, l0, v24, v25<<(uint(int32(2))%32), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
					v36 = F_luaM_realloc_(m, l0, v31, v32*int32(12), int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
						v43 = F_luaM_realloc_(m, l0, v38, v39<<(uint(int32(2))%32), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							v47 = F_luaM_realloc_(m, l0, l1, int32(76), int32(0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_luaF_newCclosure(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v2 = l1
	v4 = int32(0)
	v11 = F_luaM_realloc_(m, l0, v4, v4, v2<<(uint(int32(4))%32)+int32(24))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(6)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v11
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)))
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v15)
		v23 = v20 & int32(3)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+5)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l2
		v26 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)) = uint8(v26)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)) = uint8(v2)
		return v11
	}
}
func F_luaF_newLclosure(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	v2 = l1
	v4 = int32(0)
	v9 = v2 << (uint(int32(2)) % 32)
	v12 = F_luaM_realloc_(m, l0, v4, v4, v9+int32(20))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(6)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v12
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+20)))
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)) = uint8(v16)
		v24 = v21 & int32(3)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)) = uint8(v24)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l2
		v27 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+6)) = uint8(v27)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)) = uint8(v2)
		if v2 == v27 {
		} else {
			v36 = F__emscripten_memset_bulkmem(m, v12+int32(20), base.I32_extend8_s(int32(0)), v9)
			mBase = m.M
		}
		return v12
	}
}
func F_luaF_newproto(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	v2 = int32(0)
	v6 = F_luaM_realloc_(m, l0, v2, v2, int32(76))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(9)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v6
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v10)
		v18 = v15 & int32(3)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)) = uint8(v18)
		v22 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6+int32(64)))) = v22
		v26 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(56)))) = v26
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(48)))) = v26
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(40)))) = v26
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(32)))) = v26
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(24)))) = v26
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(16)))) = v26
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v26
		*(*int32)(unsafe.Add(mBase, uint32(v6)+72)) = v22
		return v6
	}
}
