package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_abortCommandHandler(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	if l2 != 0 {
		v19 = m.G3
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		F_luaPushError(m, v20, v19+int32(_a1939))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = F_luaError(m, v20)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				return int32(1)
			}
		}
	} else {
		v5 = m.G3
		v11 = m.G8
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		m.T0[v12].(func(*base.Module, int32, int32, int32))(m, v5+int32(_a1940), v5+int32(_a1923), int32(632))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			m.Env.Exit(m, int32(1))
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_addCommandToBatch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	v10 = m.G0
	v12 = v10 - int32(2064)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+4)) = int64(1099511627776)
	v20 = F_getKeysFromCommand(m, l0, l1, l2, v12+int32(4))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_getKeysFreeResult(m, v12+int32(4))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L13
	}
L2:
	;
	return
L3:
	;
	if v20 < int32(1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = int32(0)
	if v24 < l4 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v27 = l4
	goto L7
L6:
	;
	v27 = v24
	goto L7
L7:
	;
	v28 = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v32 = v28
	v34 = v30
	v36 = v29
	goto L8
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	if base.Ui32(v41) <= base.Ui32(v34) {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L1
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v44 = int32(2)
	v45 = v34 << (uint(v44) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v32<<(uint(int32(3))%32))))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1+v51<<(uint(v44)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v43+v45))) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v57+v45))) = v27
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v27<<(uint(v44)%32))))
	goto L11
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+36))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v68+v69<<(uint(int32(2))%32)))) = v65
	v74 = int32(1)
	v75 = v69 + v74
	*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v75
	v78 = v32 + v74
	if v78 != v20 {
		v32 = v78
		v34 = v75
		v36 = v67
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	m.G0 = v12 + int32(2064)
	return
}
func F_commandCheckArity(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if base.B2i32(int32(0) < v11)&base.B2i32(v11 != l1) != 0 {
		v21 = int32(0)
		if l2 == v21 {
			v38 = v21
			m.G0 = v9 + int32(16)
			return v38
		} else {
			v24 = int32(0)
			v26 = F_sdsnew(m, v24)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v26
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
				v34 = F_sdscatprintf(m, v26, int32(_a1611), v9)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v34
					v38 = v24
					m.G0 = v9 + int32(16)
					return v38
				}
			}
		}
	} else {
		if int32(0)-v11 <= l1 {
			v38 = int32(1)
			m.G0 = v9 + int32(16)
			return v38
		} else {
			v21 = int32(0)
			if l2 == v21 {
				v38 = v21
				m.G0 = v9 + int32(16)
				return v38
			} else {
				v24 = int32(0)
				v26 = F_sdsnew(m, v24)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v26
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
					v34 = F_sdscatprintf(m, v26, int32(_a1611), v9)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v34
						v38 = v24
						m.G0 = v9 + int32(16)
						return v38
					}
				}
			}
		}
	}
}
func F_commandCountCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, _consts[318]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v4+v5))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_commandDocsCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	v10 = m.G0
	v12 = v10 - int32(64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v14 != int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(64)
	return
L2:
	;
	v109 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L27
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[318]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	goto L4
L4:
	;
	F_addReplyMapLen(m, l0, v19+v20)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v25 = v12 + int32(16)
	v26 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[318]))
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+14)) = uint8(v26)
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v26
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+15)) = uint8(v26)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(-1)
	if v27 == v26 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v50 = F_hashtableNext(m, v12+int32(16), v12+int32(12))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L12
	}
L8:
	;
	goto L7
L9:
	;
	goto L8
L11:
	;
	F_hashtableCleanupIterator(m, v12+int32(16))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L26
	}
L12:
	;
	if v50 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	goto L14
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+140))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+int32(-1)))))
	switch v68 & int32(7) {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		goto L17
	default:
		v85 = int32(0)
		goto L16
	}
L15:
	;
	goto L11
L16:
	;
	F_addReplyBulkCBuffer(m, l0, v65, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L22
	}
L17:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v65+int32(-17))))
	v85 = v84
	goto L16
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v65+int32(-9))))
	v85 = v81
	goto L16
L19:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65+int32(-5)))))
	v85 = v78
	goto L16
L20:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+int32(-3)))))
	v85 = v75
	goto L16
L21:
	;
	v85 = int32(base.Ui32(v68) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	F_addReplyCommandDocs(m, l0, v64)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v94 = F_hashtableNext(m, v12+int32(16), v12+int32(12))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	if v94 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	goto L15
L26:
	;
	goto L1
L27:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v112 < int32(3) {
		v181 = int32(0)
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_setDeferredMapLen(m, l0, v109, v181)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L5
	} else {
		goto L44
	}
L29:
	;
	v119 = int32(2)
	v124 = int32(0)
	goto L30
L30:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126+v119<<(uint(int32(2))%32))))
	v131 = F_objectGetVal(m, v130)
	mBase = m.M
	v133 = *(*int32)(unsafe.Add(mBase, _consts[318]))
	v134 = F_lookupCommandBySdsLogic(m, v133, v131)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	v181 = v168
	goto L28
L32:
	;
	v171 = v119 + int32(1)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v171 < v172 {
		v119 = v171
		v124 = v168
		goto L30
	} else {
		goto L43
	}
L33:
	;
	if v134 == int32(0) {
		v168 = v124
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+140))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139+int32(-1)))))
	switch v142 & int32(7) {
	case 0:
		goto L40
	case 1:
		goto L39
	case 2:
		goto L38
	case 3:
		goto L37
	case 4:
		goto L36
	default:
		v159 = int32(0)
		goto L35
	}
L35:
	;
	F_addReplyBulkCBuffer(m, l0, v139, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L41
	}
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v139+int32(-17))))
	v159 = v158
	goto L35
L37:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v139+int32(-9))))
	v159 = v155
	goto L35
L38:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139+int32(-5)))))
	v159 = v152
	goto L35
L39:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139+int32(-3)))))
	v159 = v149
	goto L35
L40:
	;
	v159 = int32(base.Ui32(v142) >> (uint(int32(3)) % 32))
	goto L35
L41:
	;
	F_addReplyCommandDocs(m, l0, v134)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v168 = v124 + int32(1)
	goto L32
L43:
	;
	goto L31
L44:
	;
	goto L1
}
func F_commandGetKeysAndFlagsCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_getKeysSubcommandImpl(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_commandGetKeysCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_getKeysSubcommandImpl(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_commandListWithFilter(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v14 = v11 + int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v5
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(-1)
	if l1 == v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v37 = F_hashtableNext(m, v11+int32(32), v11+int32(28))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L1
L3:
	;
	goto L2
L5:
	;
	F_hashtableCleanupIterator(m, v11+int32(32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L6
	} else {
		goto L26
	}
L6:
	;
	return
L7:
	;
	if v37 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v50 = F_shouldFilterFromCommandList(m, v49, l2)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L5
L11:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v49)+200))
	if v83 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	if v50 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+140))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(-1)))))
	switch v56 & int32(7) {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	default:
		v73 = int32(0)
		goto L14
	}
L14:
	;
	F_addReplyBulkCBuffer(m, l0, v53, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L20
	}
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-17))))
	v73 = v72
	goto L14
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-9))))
	v73 = v69
	goto L14
L17:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53+int32(-5)))))
	v73 = v66
	goto L14
L18:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(-3)))))
	v73 = v63
	goto L14
L19:
	;
	v73 = int32(base.Ui32(v56) >> (uint(int32(3)) % 32))
	goto L14
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v76 + int32(1)
	goto L11
L21:
	;
	v106 = F_hashtableNext(m, v11+int32(32), v11+int32(28))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	v86 = int32(16)
	v90 = *(*int64)(unsafe.Add(mBase, uint32(l2+v86)))
	*(*int64)(unsafe.Add(mBase, uint32(v11+v86))) = v90
	v92 = int32(8)
	v96 = *(*int64)(unsafe.Add(mBase, uint32(l2+v92)))
	*(*int64)(unsafe.Add(mBase, uint32(v11+v92))) = v96
	v98 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v98
	F_commandListWithFilter(m, l0, v83, v11, l3)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if v106 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	goto L10
L26:
	;
	m.G0 = v11 + int32(80)
	return
}
func F_commandTimeSnapshot(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, _consts[98]))
	return v2
}
func F_evalCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_replicationFeedMonitors(m, l0, v3, v5, v6, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
		if v10&int32(16) != 0 {
			v16 = F_scriptingEngineDebuggerStartSession(m, l0)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				if v16 == int32(0) {
					F_scriptingEngineDebuggerDisable(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						return
					}
				} else {
					F_evalGenericCommand(m, l0, int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						F_scriptingEngineDebuggerEndSession(m, l0)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			F_evalGenericCommand(m, l0, int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_generateCommandResponse(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v4 = m.G0
	v6 = v4 - int32(64)
	m.G0 = v6
	v8 = F_createCachedResponseClient(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[318]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	goto L3
L3:
	;
	F_addReplyArrayLen(m, v8, v14+v15)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = v6 + int32(16)
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[318]))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+14)) = uint8(v21)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v21
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+15)) = uint8(v21)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = int32(-1)
	if v22 == v21 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v45 = F_hashtableNext(m, v6+int32(16), v6+int32(12))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	goto L5
L7:
	;
	goto L6
L9:
	;
	F_hashtableCleanupIterator(m, v6+int32(16))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L17
	}
L10:
	;
	if v45 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	goto L12
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	F_addReplyCommandInfo(m, v8, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	v59 = F_hashtableNext(m, v6+int32(16), v6+int32(12))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v59 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v68 = F_aggregateClientOutputBuffer(m, v8)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_deleteCachedResponseClient(m, v8)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	m.G0 = v6 + int32(64)
	return v68
}
func F_lookupCommandOrOriginal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[318]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v3
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = F_objectGetVal(m, v14)
	mBase = m.M
	v18 = F_hashtableFind(m, v11, v15, v8+int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		if v18 == int32(0) {
			v40 = v22
			if v40 != 0 {
				v71 = v40
				m.G0 = v8 + int32(16)
				return v71
			} else {
				v42 = int32(0)
				v43 = *(*int32)(unsafe.Add(mBase, _consts[7]))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v42
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v47 = F_objectGetVal(m, v46)
				mBase = m.M
				v50 = F_hashtableFind(m, v43, v47, v8+int32(8))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
					if v50 == int32(0) {
						v71 = v52
						m.G0 = v8 + int32(16)
						return v71
					} else {
						if l1 == int32(1) {
							v71 = v52
							m.G0 = v8 + int32(16)
							return v71
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+200))
							if v57 == int32(0) {
								v71 = v52
								m.G0 = v8 + int32(16)
								return v71
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v61 = F_objectGetVal(m, v60)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v52)+200))
								v67 = F_hashtableFind(m, v64, v61, v8+int32(12))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
									v71 = v69
									m.G0 = v8 + int32(16)
									return v71
								}
							}
						}
					}
				}
			}
		} else {
			if l1 == int32(1) {
				v71 = v22
				m.G0 = v8 + int32(16)
				return v71
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+200))
				if v27 == int32(0) {
					v71 = v22
					m.G0 = v8 + int32(16)
					return v71
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v31 = F_objectGetVal(m, v30)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+200))
					v37 = F_hashtableFind(m, v34, v31, v8+int32(12))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						v40 = v39
						if v40 != 0 {
							v71 = v40
							m.G0 = v8 + int32(16)
							return v71
						} else {
							v42 = int32(0)
							v43 = *(*int32)(unsafe.Add(mBase, _consts[7]))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v42
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v47 = F_objectGetVal(m, v46)
							mBase = m.M
							v50 = F_hashtableFind(m, v43, v47, v8+int32(8))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
								if v50 == int32(0) {
									v71 = v52
									m.G0 = v8 + int32(16)
									return v71
								} else {
									if l1 == int32(1) {
										v71 = v52
										m.G0 = v8 + int32(16)
										return v71
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+200))
										if v57 == int32(0) {
											v71 = v52
											m.G0 = v8 + int32(16)
											return v71
										} else {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v61 = F_objectGetVal(m, v60)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v52)+200))
											v67 = F_hashtableFind(m, v64, v61, v8+int32(12))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
												v71 = v69
												m.G0 = v8 + int32(16)
												return v71
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_restoreCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int64
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int64
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int64
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int64
	_ = v387
	var v391 int64
	_ = v391
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int64
	_ = v454
	var v461 int64
	_ = v461
	var v462 int64
	_ = v462
	var v463 int64
	_ = v463
	var v465 int64
	_ = v465
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int64
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int64
	_ = v518
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int64
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int64
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int64
	_ = v555
	var v556 int64
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int64
	_ = v574
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	v12 = m.G0
	v14 = v12 - int32(144)
	m.G0 = v14
	v16 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v14)+120)) = v16
	v21 = l0 + int32(20)
	v22 = int32(4)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v23 <= v22 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v14 + int32(144)
	return
L2:
	;
	v579 = *(*int32)(unsafe.Add(mBase, _consts[129]))
	F_addReplyErrorObject(m, l0, v579)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L54
	} else {
		goto L164
	}
L3:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v309)+8))
	v317 = F_getLongLongFromObjectOrReply(m, l0, v313, v14+int32(136), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L54
	} else {
		goto L84
	}
L4:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v298 = F_lookupKeyWrite(m, v297, v289)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L54
	} else {
		goto L82
	}
L5:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	v288 = v21
	v289 = v284
	v290 = int32(0)
	goto L4
L6:
	;
	v26 = int32(0)
	v31 = v22
	v32 = v23
	v33 = v26
	v34 = v26
	goto L7
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v41 = v31 << (uint(int32(2)) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v41)))
	v44 = F_objectGetVal(m, v43)
	mBase = m.M
	v45 = int32(_a202)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v275 = l0 + int32(20)
	v276 = int32(0)
	v278 = base.B2i32(v266 != v276)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	if v265 == v276 {
		v288 = v275
		v289 = v280
		v290 = v278
		goto L4
	} else {
		goto L81
	}
L9:
	;
	v271 = v263 + int32(1)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v271 < v272 {
		v31 = v271
		v32 = v272
		v33 = v265
		v34 = v266
		goto L7
	} else {
		goto L80
	}
L10:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v41)))
	v88 = F_objectGetVal(m, v87)
	mBase = m.M
	v89 = int32(_a203)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 != 0 {
		goto L27
	} else {
		goto L28
	}
L11:
	;
	if v80-v82 != 0 {
		goto L10
	} else {
		goto L23
	}
L12:
	;
	v80 = F_tolower(m, v76)
	mBase = m.M
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v82 = F_tolower(m, v81)
	mBase = m.M
	goto L11
L13:
	;
	v50 = v44
	v51 = v45
	v52 = v48
	goto L16
L14:
	;
	v76 = int32(0)
	v77 = v45
	goto L12
L15:
	;
	v76 = v73 & int32(255)
	v77 = v72
	goto L12
L16:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v54 == int32(0) {
		v72 = v51
		v73 = v52
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v72 = v66
	v73 = int32(0)
	goto L15
L18:
	;
	v58 = v52 & int32(255)
	if v58 == v54 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v65 = int32(1)
	v66 = v51 + v65
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v67 != 0 {
		v50 = v50 + v65
		v51 = v66
		v52 = v67
		goto L16
	} else {
		goto L22
	}
L20:
	;
	v60 = F_tolower(m, v58)
	mBase = m.M
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v62 = F_tolower(m, v61)
	mBase = m.M
	if v60 == v62 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v72 = v51
	v73 = v64
	goto L15
L22:
	;
	goto L17
L23:
	;
	v263 = v31
	v265 = int32(1)
	v266 = v34
	goto L9
L24:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129+v41)))
	v132 = F_objectGetVal(m, v131)
	mBase = m.M
	v133 = int32(_a204)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v136 != 0 {
		goto L40
	} else {
		goto L41
	}
L25:
	;
	if v124-v126 != 0 {
		goto L24
	} else {
		goto L37
	}
L26:
	;
	v124 = F_tolower(m, v120)
	mBase = m.M
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v126 = F_tolower(m, v125)
	mBase = m.M
	goto L25
L27:
	;
	v94 = v88
	v95 = v89
	v96 = v92
	goto L30
L28:
	;
	v120 = int32(0)
	v121 = v89
	goto L26
L29:
	;
	v120 = v117 & int32(255)
	v121 = v116
	goto L26
L30:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v98 == int32(0) {
		v116 = v95
		v117 = v96
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v116 = v110
	v117 = int32(0)
	goto L29
L32:
	;
	v102 = v96 & int32(255)
	if v102 == v98 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v109 = int32(1)
	v110 = v95 + v109
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	if v111 != 0 {
		v94 = v94 + v109
		v95 = v110
		v96 = v111
		goto L30
	} else {
		goto L36
	}
L34:
	;
	v104 = F_tolower(m, v102)
	mBase = m.M
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v106 = F_tolower(m, v105)
	mBase = m.M
	if v104 == v106 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v116 = v95
	v117 = v108
	goto L29
L36:
	;
	goto L31
L37:
	;
	v263 = v31
	v265 = v33
	v266 = int32(1)
	goto L9
L38:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v175 = base.B2i32(int32(-2) < v31-v32)
	if int32(-2) < v31-v32 {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v168 = F_tolower(m, v164)
	mBase = m.M
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	v170 = F_tolower(m, v169)
	mBase = m.M
	goto L38
L40:
	;
	v138 = v132
	v139 = v133
	v140 = v136
	goto L43
L41:
	;
	v164 = int32(0)
	v165 = v133
	goto L39
L42:
	;
	v164 = v161 & int32(255)
	v165 = v160
	goto L39
L43:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v142 == int32(0) {
		v160 = v139
		v161 = v140
		goto L42
	} else {
		goto L45
	}
L44:
	;
	v160 = v154
	v161 = int32(0)
	goto L42
L45:
	;
	v146 = v140 & int32(255)
	if v146 == v142 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v153 = int32(1)
	v154 = v139 + v153
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	if v155 != 0 {
		v138 = v138 + v153
		v139 = v154
		v140 = v155
		goto L43
	} else {
		goto L49
	}
L47:
	;
	v148 = F_tolower(m, v146)
	mBase = m.M
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	v150 = F_tolower(m, v149)
	mBase = m.M
	if v148 == v150 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v160 = v139
	v161 = v152
	goto L42
L49:
	;
	goto L44
L50:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v172+v41)))
	v198 = F_objectGetVal(m, v197)
	mBase = m.M
	v199 = int32(_a205)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if v202 != 0 {
		goto L61
	} else {
		goto L62
	}
L51:
	;
	if v168-v170 != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v14)+128))
	if v176 != int64(-1) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v180 = v31 + int32(1)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v172+v180<<(uint(int32(2))%32))))
	v188 = F_getLongLongFromObjectOrReply(m, l0, v184, v14+int32(120), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	return
L55:
	;
	if v188 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v14)+120))
	if int64(-1) < v190 {
		v263 = v180
		v265 = v33
		v266 = v34
		goto L9
	} else {
		goto L57
	}
L57:
	;
	F_addReplyError(m, l0, int32(_a206))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	goto L1
L59:
	;
	if int32(-2) < v31-v32 {
		goto L71
	} else {
		goto L72
	}
L60:
	;
	v234 = F_tolower(m, v230)
	mBase = m.M
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	v236 = F_tolower(m, v235)
	mBase = m.M
	goto L59
L61:
	;
	v204 = v198
	v205 = v199
	v206 = v202
	goto L64
L62:
	;
	v230 = int32(0)
	v231 = v199
	goto L60
L63:
	;
	v230 = v227 & int32(255)
	v231 = v226
	goto L60
L64:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v208 == int32(0) {
		v226 = v205
		v227 = v206
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v226 = v220
	v227 = int32(0)
	goto L63
L66:
	;
	v212 = v206 & int32(255)
	if v212 == v208 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v219 = int32(1)
	v220 = v205 + v219
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+1)))
	if v221 != 0 {
		v204 = v204 + v219
		v205 = v220
		v206 = v221
		goto L64
	} else {
		goto L70
	}
L68:
	;
	v214 = F_tolower(m, v212)
	mBase = m.M
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v216 = F_tolower(m, v215)
	mBase = m.M
	if v214 == v216 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	v226 = v205
	v227 = v218
	goto L63
L70:
	;
	goto L65
L71:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_addReplyErrorObject(m, l0, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L54
	} else {
		goto L79
	}
L72:
	;
	if v234-v236 != 0 {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v14)+120))
	if v238 != int64(-1) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v243 = v31 + int32(1)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v241+v243<<(uint(int32(2))%32))))
	v251 = F_getLongLongFromObjectOrReply(m, l0, v247, v14+int32(128), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L54
	} else {
		goto L75
	}
L75:
	;
	if v251 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v14)+128))
	if base.Ui64(v253) < base.Ui64(int64(256)) {
		v263 = v243
		v265 = v33
		v266 = v34
		goto L9
	} else {
		goto L77
	}
L77:
	;
	F_addReplyError(m, l0, int32(_a207))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L54
	} else {
		goto L78
	}
L78:
	;
	goto L1
L79:
	;
	goto L1
L80:
	;
	goto L8
L81:
	;
	v304 = v275
	v305 = v280
	v306 = v278
	v309 = v279
	v311 = v276
	goto L3
L82:
	;
	if v298 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v304 = v288
	v305 = v289
	v306 = v290
	v309 = v300
	v311 = int32(1)
	goto L3
L84:
	;
	if v317 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v14)+136))
	if int64(-1) < v319 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+12))
	v327 = F_objectGetVal(m, v326)
	mBase = m.M
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	v330 = F_objectGetVal(m, v329)
	mBase = m.M
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330+int32(-1)))))
	switch v333 & int32(7) {
	case 0:
		goto L96
	case 1:
		goto L95
	case 2:
		goto L94
	case 3:
		goto L93
	case 4:
		goto L92
	default:
		goto L90
	}
L87:
	;
	F_addReplyError(m, l0, int32(_a208))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L54
	} else {
		goto L88
	}
L88:
	;
	goto L1
L89:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+12))
	v405 = F_objectGetVal(m, v404)
	mBase = m.M
	v408 = F___memcpy(m, v14+int32(40), int32(_a209), int32(80))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v408)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v408)+48)) = v405
	goto L111
L90:
	;
	F_addReplyError(m, l0, int32(_a210))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L54
	} else {
		goto L110
	}
L91:
	;
	if base.Ui32(v350) < base.Ui32(int32(10)) {
		goto L90
	} else {
		goto L97
	}
L92:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v330+int32(-17))))
	v350 = v349
	goto L91
L93:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v330+int32(-9))))
	v350 = v346
	goto L91
L94:
	;
	v343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v330+int32(-5)))))
	v350 = v343
	goto L91
L95:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330+int32(-3)))))
	v350 = v340
	goto L91
L96:
	;
	v350 = int32(base.Ui32(v333) >> (uint(int32(3)) % 32))
	goto L91
L97:
	;
	v353 = v327 + v350
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v353+int32(-10)))))
	v357 = int32(0)
	if v356 < int32(1) {
		v378 = v357
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v378 == int32(0) {
		goto L90
	} else {
		goto L107
	}
L99:
	;
	goto L98
L100:
	;
	if v357&base.B2i32(v356 < int32(80)) != 0 {
		v378 = v357
		goto L99
	} else {
		goto L101
	}
L101:
	;
	if base.B2i32(int32(79) < v356)&v357 != 0 {
		v378 = v357
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	if v370 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v378 = int32(1)
	goto L99
L104:
	;
	if base.Ui32(int32(80)) < base.Ui32(v356) {
		v378 = v357
		goto L99
	} else {
		goto L105
	}
L105:
	;
	if base.Ui32(v356+int32(-12)) < base.Ui32(int32(68)) {
		v378 = v357
		goto L99
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	if v382 != 0 {
		goto L89
	} else {
		goto L108
	}
L108:
	;
	v384 = int32(-8)
	v387 = F_crc64(m, int64(0), v327, base.I64_extend_i32_u(v350+v384))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v387
	v391 = *(*int64)(unsafe.Add(mBase, uint32(v353+v384)))
	if v387 == v391 {
		goto L89
	} else {
		goto L109
	}
L109:
	;
	goto L90
L110:
	;
	goto L1
L111:
	;
	v414 = F_rdbLoadObjectType(m, v14+int32(40))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L54
	} else {
		goto L113
	}
L112:
	;
	if base.Ui32(int32(67)) < base.Ui32((v356+int32(-12))&int32(65535)) {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	if v414 != int32(-1) {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	F_addReplyError(m, l0, int32(_a211))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L54
	} else {
		goto L115
	}
L115:
	;
	goto L1
L116:
	;
	v435 = F_objectGetVal(m, v305)
	mBase = m.M
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+28))
	v438 = int32(0)
	v441 = F_rdbLoadObject(m, v414, v14+int32(40), v435, v437, v438, v438, int64(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L54
	} else {
		goto L120
	}
L117:
	;
	if v414 < int32(22) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v414
	F_addReplyErrorFormat(m, l0, int32(_a212), v14)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L54
	} else {
		goto L119
	}
L119:
	;
	goto L1
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v441
	if v441 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	if v311 != 0 {
		v453 = int32(1)
		goto L124
	} else {
		goto L125
	}
L122:
	;
	F_addReplyError(m, l0, int32(_a211))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L54
	} else {
		goto L123
	}
L123:
	;
	goto L1
L124:
	;
	v454 = *(*int64)(unsafe.Add(mBase, uint32(v14)+136))
	if (base.B2i32(v454 == int64(0))|v306)&int32(1) != 0 {
		v465 = v454
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v449 = F_dbDelete(m, v448, v305)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L54
	} else {
		goto L126
	}
L126:
	;
	v453 = base.B2i32(v449 == int32(0))
	goto L124
L127:
	;
	if v465 == int64(0) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v461 = *(*int64)(unsafe.Add(mBase, _consts[98]))
	goto L129
L129:
	;
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v14)+136))
	v463 = v461 + v462
	*(*int64)(unsafe.Add(mBase, uint32(v14)+136)) = v463
	v465 = v463
	goto L127
L130:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbAdd(m, v528, v305, v14+int32(36))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L54
	} else {
		goto L151
	}
L131:
	;
	v471 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	if v471 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	if v489 == int32(0) {
		goto L130
	} else {
		goto L140
	}
L133:
	;
	goto L132
L134:
	;
	v477 = int32(0)
	v478 = F_commandTimeSnapshot(m)
	mBase = m.M
	if v478 < v465 {
		v489 = v477
		goto L133
	} else {
		goto L137
	}
L135:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v471)+216))
	if v475 != 0 {
		v489 = int32(0)
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v480 = int32(_a44)
	v481 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v483 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	if v483 != 0 {
		v489 = v477
		goto L133
	} else {
		goto L138
	}
L138:
	;
	if v481 != 0 {
		v489 = v477
		goto L133
	} else {
		goto L139
	}
L139:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	v489 = base.B2i32(v485 == int32(0))
	goto L133
L140:
	;
	if v453 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	F_decrRefCount(m, v441)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L54
	} else {
		goto L149
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v305
	v497 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	if v497 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v498 = int32(244)
	goto L145
L144:
	;
	v498 = int32(240)
	goto L145
L145:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v498)+uint32(_consts[84])))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v500
	F_rewriteClientCommandVector(m, l0, int32(2), v14+int32(16))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L54
	} else {
		goto L146
	}
L146:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v507, v305)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L54
	} else {
		goto L147
	}
L147:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a213), v305, v513)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L54
	} else {
		goto L148
	}
L148:
	;
	v516 = int32(_a44)
	v518 = *(*int64)(unsafe.Add(mBase, _consts[83]))
	*(*int64)(unsafe.Add(mBase, _consts[83])) = v518 + int64(1)
	goto L141
L149:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	F_addReply(m, l0, v525)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L54
	} else {
		goto L150
	}
L150:
	;
	goto L1
L151:
	;
	v533 = *(*int64)(unsafe.Add(mBase, uint32(v14)+136))
	if v533 == int64(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v555 = *(*int64)(unsafe.Add(mBase, uint32(v14)+128))
	v556 = *(*int64)(unsafe.Add(mBase, uint32(v14)+120))
	v557 = F_objectSetLRUOrLFU(m, v554, v555, v556)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L54
	} else {
		goto L160
	}
L153:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v537 = F_setExpire(m, l0, v536, v305, v533)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L54
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v537
	if v306 != 0 {
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v541 = *(*int64)(unsafe.Add(mBase, uint32(v14)+136))
	v542 = F_createStringObjectFromLongLong(m, v541)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L54
	} else {
		goto L156
	}
L156:
	;
	F_rewriteClientCommandArgument(m, l0, int32(2), v542)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L54
	} else {
		goto L157
	}
L157:
	;
	F_decrRefCount(m, v542)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L54
	} else {
		goto L158
	}
L158:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v550 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	F_rewriteClientCommandArgument(m, l0, v548, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L54
	} else {
		goto L159
	}
L159:
	;
	goto L152
L160:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v559, v305)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L54
	} else {
		goto L161
	}
L161:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a214), v305, v565)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L54
	} else {
		goto L162
	}
L162:
	;
	v569 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	F_addReply(m, l0, v569)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L54
	} else {
		goto L163
	}
L163:
	;
	v572 = int32(_a44)
	v574 = *(*int64)(unsafe.Add(mBase, _consts[83]))
	*(*int64)(unsafe.Add(mBase, _consts[83])) = v574 + int64(1)
	goto L1
L164:
	;
	goto L1
}
func F_sendCommandArgv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = F_sdsempty(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l1
	v22 = F_sdscatfmt(m, v14, int32(_a1258), v12+int32(32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l1 < int32(1) {
		v121 = v22
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v124 = int32(0)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121+int32(-1)))))
	switch v128 & int32(7) {
	case 0:
		goto L36
	case 1:
		goto L35
	case 2:
		goto L34
	case 3:
		goto L33
	case 4:
		goto L32
	default:
		v145 = v124
		goto L31
	}
L5:
	;
	v32 = int32(0)
	v33 = v22
	goto L6
L6:
	;
	v37 = v32 << (uint(int32(2)) % 32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2+v37)))
	if l3 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v121 = v110
	goto L4
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v99
	v104 = F_sdscatfmt(m, v33, int32(_a1259), v12+int32(16))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L27
	}
L9:
	;
	if v39&int32(3) == int32(0) {
		v65 = v39
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3+v37)))
	v99 = v43
	goto L8
L11:
	;
	v99 = v98
	goto L8
L12:
	;
	v98 = v90 - v39
	goto L11
L13:
	;
	v69 = v65
	goto L21
L14:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v51 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = v39
	goto L17
L16:
	;
	v98 = v39 - v39
	goto L11
L17:
	;
	v58 = v54 + int32(1)
	if v58&int32(3) == int32(0) {
		v65 = v58
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v63 != 0 {
		v54 = v58
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v90 = v58
	goto L12
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v78 = int32(-2139062144)
	if (int32(16843008)-v75|v75)&v78 == v78 {
		v69 = v69 + int32(4)
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v84 = v69
	goto L24
L23:
	;
	goto L22
L24:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v88 != 0 {
		v84 = v84 + int32(1)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v90 = v84
	goto L12
L26:
	;
	goto L25
L27:
	;
	v106 = F_sdscatlen(m, v104, v39, v99)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v110 = F_sdscatlen(m, v106, int32(_a132), int32(2))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v113 = v32 + int32(1)
	if v113 != l1 {
		v32 = v113
		v33 = v110
		goto L6
	} else {
		goto L30
	}
L30:
	;
	goto L7
L31:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+92))
	v153 = m.T0[v152].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v121, v145, base.I64_extend_i32_s(v147*int32(1000)))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L38
	}
L32:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v121+int32(-17))))
	v145 = v144
	goto L31
L33:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v121+int32(-9))))
	v145 = v141
	goto L31
L34:
	;
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121+int32(-5)))))
	v145 = v138
	goto L31
L35:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121+int32(-3)))))
	v145 = v135
	goto L31
L36:
	;
	v145 = int32(base.Ui32(v128) >> (uint(int32(3)) % 32))
	goto L31
L37:
	;
	F_sdsfree(m, v121)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L43
	}
L38:
	;
	if v153 != int32(-1) {
		v168 = v124
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v157 = F_sdsempty(m)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+88))
	v161 = m.T0[v160].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v161
	v165 = F_sdscatprintf(m, v157, int32(_a1260), v12)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v168 = v165
	goto L37
L43:
	;
	m.G0 = v12 + int32(48)
	return v168
}
func F_setCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v2
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v27 = F_parseExtendedCommandArgumentsOrReply(m, l0, int32(1), int32(3), v19, v7, v7+int32(4), v2, v7+int32(12), v7+int32(8))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return
	} else {
		if v27 != 0 {
			m.G0 = v7 + int32(16)
			return
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
			if v31&int32(1) != 0 {
				v38 = v29
				v39 = v30
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				v44 = int32(0)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				F_setGenericCommand(m, l0, v40, v41, v39, v42, v43, v44, v44, v46)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				v34 = F_tryObjectEncoding(m, v30)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v34
					v38 = v36
					v39 = v34
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v44 = int32(0)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					F_setGenericCommand(m, l0, v40, v41, v39, v42, v43, v44, v44, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_trimCommandQueue(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
	if v5&int32(4) != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if v8 == int32(0) {
			return
		} else {
			v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
			switch v11 {
			case 0:
				F_valkey_free(m, v8)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					v14 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v14)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v14
					return
				}
			case 1:
				v27 = v11
				if base.Ui32(v27) < base.Ui32(v11) {
					F__serverAssert(m, int32(_a976), int32(_a977), int32(4064))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v29 = int32(16)
					if base.Ui32(v29) < base.Ui32(v27) {
						v32 = v27
					} else {
						v32 = v29
					}
					v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
					if base.Ui32(v33) <= base.Ui32(v32) {
						return
					} else {
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v32)
						v38 = F_valkey_realloc(m, v8, v32*int32(40))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v38
							return
						}
					}
				}
			default:
				v27 = int32(1) << (uint(int32(32)-base.I32_clz(v11+int32(-1))) % 32) & int32(65535)
				if base.Ui32(v27) < base.Ui32(v11) {
					F__serverAssert(m, int32(_a976), int32(_a977), int32(4064))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v29 = int32(16)
					if base.Ui32(v29) < base.Ui32(v27) {
						v32 = v27
					} else {
						v32 = v29
					}
					v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
					if base.Ui32(v33) <= base.Ui32(v32) {
						return
					} else {
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v32)
						v38 = F_valkey_realloc(m, v8, v32*int32(40))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v38
							return
						}
					}
				}
			}
		}
	}
}
