package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_generic_reader(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	v5 = m.G3
	F_luaL_checkstack(m, l0, int32(2), v5+int32(_a2302))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v23 = v18 + int32(0)
		v24 = m.G398
		if base.Ui32(v23) < base.Ui32(v17) {
			v26 = v23
		} else {
			v26 = v24
		}
		v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v69 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
		*(*int64)(unsafe.Add(mBase, uint32(v68))) = v69
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v71
		v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73 + int32(16)
		F_lua_call(m, l0, int32(0), int32(1))
		mBase = m.M
		v80 = m.ExcPending
		if v80 != 0 {
			return int32(0)
		} else {
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v98 = v95 + int32(-16)
			v132 = m.G398
			if v98 != v132 {
				v135 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
				v138 = v135
			} else {
				v138 = int32(-1)
			}
			if v138 != 0 {
				v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v160 = v157 + int32(-16)
				v194 = m.G398
				if v160 != v194 {
					v197 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
					v204 = base.B2i32(base.Ui32(v197+int32(-3)) < base.Ui32(int32(2)))
				} else {
					v204 = int32(0)
				}
				if v204 == int32(0) {
					v214 = m.G3
					v218 = F_luaL_error(m, l0, v214+int32(_a2303), int32(0))
					mBase = m.M
					v219 = m.ExcPending
					if v219 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					F_lua_replace(m, l0, int32(3))
					mBase = m.M
					v209 = m.ExcPending
					if v209 != 0 {
						return int32(0)
					} else {
						v211 = F_lua_tolstring(m, l0, int32(3), l2)
						mBase = m.M
						v212 = m.ExcPending
						if v212 != 0 {
							return int32(0)
						} else {
							return v211
						}
					}
				}
			} else {
				v139 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v139
				return v139
			}
		}
	}
}
func F_pushGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = F_lookupKeyWrite(m, v10, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v13
	v17 = F_checkType(m, l0, v13, int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v8 + int32(16)
	return
L4:
	;
	if v17 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v13 != 0 {
		v36 = v13
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v37 = int32(2)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v44 = int32(0)
	F_listTypeTryConversionRaw(m, v36, int32(1), v39, v37, v41+int32(-1), v44, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	if l2 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v25 = F_createListListpackObject(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[233]))
	F_addReply(m, l0, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L3
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v25
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	F_dbAdd(m, v28, v30, v8+int32(12))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v36 = v35
	goto L6
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v48 < int32(3) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	F_signalModifiedKey(m, l0, v79, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L20
	}
L15:
	;
	v53 = v37
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v53<<(uint(int32(2))%32))))
	F_listTypePush(m, v56, v61, l1)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L14
L18:
	;
	v64 = int32(_a69)
	v66 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v66 + int64(1)
	v71 = v53 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v71 < v72 {
		v53 = v71
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	if l1 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v87 = int32(_a1499)
	goto L23
L22:
	;
	v87 = int32(_a1500)
	goto L23
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+28))
	F_notifyKeyspaceEvent(m, int32(16), v87, v89, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v95 = F_listTypeLength(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v95))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L3
}
