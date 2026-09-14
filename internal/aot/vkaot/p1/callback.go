package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_callbackHash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v7 & int32(7) {
	case 0:
		v13 = F_siphash(m, l0, int32(base.Ui32(v7)>>(uint(int32(3))%32)), int32(_a245))
		mBase = m.M
		return v13
	case 1:
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v19 = F_siphash(m, l0, v17, int32(_a245))
		mBase = m.M
		return v19
	case 2:
		v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v25 = F_siphash(m, l0, v23, int32(_a245))
		mBase = m.M
		return v25
	case 3:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v31 = F_siphash(m, l0, v29, int32(_a245))
		mBase = m.M
		return v31
	case 4:
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v36 = v35
		v38 = F_siphash(m, l0, v36, int32(_a245))
		mBase = m.M
		return v38
	default:
		v36 = int32(0)
		v38 = F_siphash(m, l0, v36, int32(_a245))
		mBase = m.M
		return v38
	}
}
func F_callbackKeyCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	v3 = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v10 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v27 = v3
		goto L1
	}
L1:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v30 & int32(7) {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	case 4:
		goto L8
	default:
		v47 = v3
		goto L7
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v27 = v26
	goto L1
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v27 = v23
	goto L1
L4:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v27 = v20
	goto L1
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v27 = v17
	goto L1
L6:
	;
	v27 = int32(base.Ui32(v10) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	if v27 != v47 {
		v116 = int32(0)
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v47 = v46
	goto L7
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v47 = v43
	goto L7
L10:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v47 = v40
	goto L7
L11:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v47 = v37
	goto L7
L12:
	;
	v47 = int32(base.Ui32(v30) >> (uint(int32(3)) % 32))
	goto L7
L13:
	;
	return v116
L14:
	;
	if base.Ui32(v27) < base.Ui32(int32(4)) {
		v73 = l0
		v74 = l1
		v75 = v27
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v116 = base.B2i32(v113 == int32(0))
	goto L13
L16:
	;
	v113 = int32(0)
	goto L15
L17:
	;
	v85 = v80
	v86 = v81
	v87 = v82
	goto L27
L18:
	;
	if v75 == int32(0) {
		goto L16
	} else {
		goto L25
	}
L19:
	;
	if (l1|l0)&int32(3) != 0 {
		v80 = l0
		v81 = l1
		v82 = v27
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v57 = l0
	v58 = l1
	v59 = v27
	goto L21
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v62 != v63 {
		v80 = v57
		v81 = v58
		v82 = v59
		goto L17
	} else {
		goto L23
	}
L22:
	;
	v73 = v68
	v74 = v66
	v75 = v70
	goto L18
L23:
	;
	v65 = int32(4)
	v66 = v58 + v65
	v68 = v57 + v65
	v70 = v59 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v70) {
		v57 = v68
		v58 = v66
		v59 = v70
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v80 = v73
	v81 = v74
	v82 = v75
	goto L17
L26:
	;
	v113 = v90 - v91
	goto L15
L27:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v90 != v91 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v93 = int32(1)
	v98 = v87 + int32(-1)
	if v98 == int32(0) {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v85 = v85 + v93
	v86 = v86 + v93
	v87 = v98
	goto L27
}
func F_setEndCallback(m *base.Module, l0 int32) {
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
func F_setStartCallback(m *base.Module, l0 int32, l1 int32) {
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
				F_lua_pushstring(m, v31, v44+int32(_a2432))
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
						v58 = l0 + v54<<(uint(int32(2))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v58+int32(1040)))) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v58+int32(16)))) = int32(0)
						return
					}
				}
			}
		} else {
			F__serverPanic_2(m, int32(1001))
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
