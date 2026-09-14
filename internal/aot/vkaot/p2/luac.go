package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaC_callGCTM(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+48))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L3
L3:
	;
	F_GCTM(m, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
	if v10 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
}
func F_luaC_separateudata(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+112))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v11 != 0 {
		v18 = v10
		v19 = v11
		v20 = int32(0)
		for {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)))
			if v23&int32(8) == int32(0) {
				if v23&int32(3)|l1 != 0 {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
					if v31 == int32(0) {
						v50 = v23
						v53 = v50 | int32(8)
						*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)) = uint8(v53)
						v71 = v19
						v72 = v20
					} else {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+6)))
						if v34&int32(4) != 0 {
							v50 = v23
							v53 = v50 | int32(8)
							*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)) = uint8(v53)
							v71 = v19
							v72 = v20
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+196))
							v40 = F_luaH_getstr(m, v31, v39)
							mBase = m.M
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
							if v41 != 0 {
								v48 = v40
							} else {
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+6)))
								v45 = v42 | int32(4)
								*(*uint8)(unsafe.Add(mBase, uint32(v31)+6)) = uint8(v45)
								v48 = int32(0)
							}
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)))
							if v48 != 0 {
								v56 = v49 | int32(8)
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)) = uint8(v56)
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
								*(*int32)(unsafe.Add(mBase, uint32(v18))) = v59
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
								if v64 != 0 {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
									*(*int32)(unsafe.Add(mBase, uint32(v19))) = v66
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v68))) = v19
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v19))) = v19
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v19
								v71 = v18
								v72 = v20 + v58 + int32(24)
							} else {
								v50 = v49
								v53 = v50 | int32(8)
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)) = uint8(v53)
								v71 = v19
								v72 = v20
							}
						}
					}
				} else {
					v71 = v19
					v72 = v20
				}
			} else {
				v71 = v19
				v72 = v20
			}
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			if v75 != 0 {
				v18 = v71
				v19 = v75
				v20 = v72
				continue
			} else {
				break
			}
			break
		}
		return v72
	} else {
		return int32(0)
	}
}
func F_luaC_step(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+64))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+76)) = v7 - v8 + v10
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+84))
	v15 = v13 * int32(10)
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = v15
	goto L3
L2:
	;
	v17 = int32(2147483646)
	goto L3
L3:
	;
	v20 = v17
	goto L5
L4:
	;
	if v25 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v23 = F_singlestep(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	return
L8:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+21)))
	if v25 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v28 = v20 - v23
	if int32(0) < v28 {
		v20 = v28
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L6
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+72))
	v48 = base.I32_div_u_s(v46, int32(100))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v48 * v49
	return
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+76))
	if base.Ui32(int32(1023)) < base.Ui32(v34) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+76)) = v34 + int32(-1024)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v44
	return
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v37 + int32(1024)
	return
}
