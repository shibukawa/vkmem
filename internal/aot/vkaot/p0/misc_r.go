package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_RedisRegisterConnectionTypeTLS(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_RedisRegisterConnectionTypeTLS[0]))
	if int32(1) < v7 {
		m.G0 = v4 + int32(16)
		return int32(-1)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_RedisRegisterConnectionTypeTLS_0)
		F__serverLog(m, int32(1), int32(_a_F_RedisRegisterConnectionTypeTLS_1), v4)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			m.G0 = v4 + int32(16)
			return int32(-1)
		}
	}
}
func F_RegisterConnectionTypeRdma(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterConnectionTypeRdma[0]))
	if int32(1) < v7 {
		m.G0 = v4 + int32(16)
		return int32(-1)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_RegisterConnectionTypeRdma_0)
		F__serverLog(m, int32(1), int32(_a_F_RegisterConnectionTypeRdma_1), v4)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			m.G0 = v4 + int32(16)
			return int32(-1)
		}
	}
}
func F___release_ptc(m *base.Module) {
	return
}
func F_randomkeyCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v4 = F_dbRandomKey(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		if v4 != 0 {
			F_addReplyBulk(m, l0, v4)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				F_decrRefCount(m, v4)
				mBase = m.M
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_addReplyNull(m, l0)
			mBase = m.M
			v7 = m.ExcPending
			if v7 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_readwriteCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v2 & int32(-131073)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_readwriteCommand[0]))
	F_addReply(m, l0, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_reallymarkobject(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	v7 = l1
	v8 = v5
	goto L8
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v7
	goto L1
L3:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = v193
	goto L2
L4:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+108)) = v191
	goto L2
L5:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v189
	goto L2
L6:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v187
	goto L2
L7:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if v101 < int32(4) {
		v178 = v100
		goto L36
	} else {
		goto L37
	}
L8:
	;
	v11 = v8 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)) = uint8(v11)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
	if v13 == int32(7) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v19 = v11 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)) = uint8(v19)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v21 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	switch v13 + int32(-5) {
	case 0:
		goto L5
	case 1:
		goto L6
	default:
		goto L1
	case 3:
		goto L4
	case 4:
		goto L3
	case 5:
		goto L7
	}
L12:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+5)))
	if v97&int32(3) != 0 {
		v7 = v96
		v8 = v97
		goto L8
	} else {
		goto L35
	}
L13:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
	if v24&int32(3) == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
	v33 = v21
	v34 = v31
	goto L23
L15:
	;
	goto L12
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v33
	goto L16
L18:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+68)) = v87
	goto L17
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v85
	goto L17
L20:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v83
	goto L17
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v81
	goto L17
L22:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v61 < int32(4) {
		v72 = v60
		goto L31
	} else {
		goto L32
	}
L23:
	;
	v37 = v34 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+5)) = uint8(v37)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
	if v39 == int32(7) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v45 = v37 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+5)) = uint8(v45)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	if v47 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	switch v39 + int32(-5) {
	case 0:
		goto L20
	case 1:
		goto L21
	default:
		goto L16
	case 3:
		goto L19
	case 4:
		goto L18
	case 5:
		goto L22
	}
L27:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+5)))
	if v57&int32(3) != 0 {
		v33 = v56
		v34 = v57
		goto L23
	} else {
		goto L30
	}
L28:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+5)))
	if v50&int32(3) == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_reallymarkobject(m, l0, v47)
	mBase = m.M
	goto L27
L30:
	;
	goto L16
L31:
	;
	if v72 != v33+int32(16) {
		goto L16
	} else {
		goto L34
	}
L32:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
	if v65&int32(3) == int32(0) {
		v72 = v60
		goto L31
	} else {
		goto L33
	}
L33:
	;
	F_reallymarkobject(m, l0, v64)
	mBase = m.M
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v72 = v71
	goto L31
L34:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+5)))
	v79 = v77 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+5)) = uint8(v79)
	goto L15
L35:
	;
	goto L1
L36:
	;
	if v178 != v7+int32(16) {
		goto L1
	} else {
		goto L59
	}
L37:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+5)))
	if v105&int32(3) == int32(0) {
		v178 = v100
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+5)))
	v114 = v104
	v115 = v112
	goto L47
L39:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v178 = v177
	goto L36
L40:
	;
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v114
	goto L40
L42:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v114)+68)) = v168
	goto L41
L43:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v114)+108)) = v166
	goto L41
L44:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v114)+32)) = v164
	goto L41
L45:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v114)+8)) = v162
	goto L41
L46:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	if v142 < int32(4) {
		v153 = v141
		goto L55
	} else {
		goto L56
	}
L47:
	;
	v118 = v115 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+5)) = uint8(v118)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
	if v120 == int32(7) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v126 = v118 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+5)) = uint8(v126)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	if v128 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	switch v120 + int32(-5) {
	case 0:
		goto L44
	case 1:
		goto L45
	default:
		goto L40
	case 3:
		goto L43
	case 4:
		goto L42
	case 5:
		goto L46
	}
L51:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+5)))
	if v138&int32(3) != 0 {
		v114 = v137
		v115 = v138
		goto L47
	} else {
		goto L54
	}
L52:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+5)))
	if v131&int32(3) == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	F_reallymarkobject(m, l0, v128)
	mBase = m.M
	goto L51
L54:
	;
	goto L40
L55:
	;
	if v153 != v114+int32(16) {
		goto L40
	} else {
		goto L58
	}
L56:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+5)))
	if v146&int32(3) == int32(0) {
		v153 = v141
		goto L55
	} else {
		goto L57
	}
L57:
	;
	F_reallymarkobject(m, l0, v145)
	mBase = m.M
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	v153 = v152
	goto L55
L58:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+5)))
	v160 = v158 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+5)) = uint8(v160)
	goto L39
L59:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)))
	v185 = v183 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)) = uint8(v185)
	return
}
func F_rebaseReplicationBuffer(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_rebaseReplicationBuffer[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	F_raxFree(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = F_raxNew(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = int32(_a_F_rebaseReplicationBuffer_0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_rebaseReplicationBuffer[0]))
	v20 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v16
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_rebaseReplicationBuffer[1]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v25
	goto L4
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v30 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	m.G0 = v9 + int32(16)
	return
L6:
	;
	if v30 == int32(0) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v30+base.B2i32(v33 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v39
	goto L7
L9:
	;
	v45 = v30
	goto L10
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)+16))
	v51 = v50 + l0
	*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v51
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_rebaseReplicationBuffer[0]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v57 = v55 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v57
	if base.Ui32(v57) < base.Ui32(int32(64)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L5
L12:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v109 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v61 = int64(56)
	v63 = int64(65280)
	v65 = int64(40)
	v68 = int64(16711680)
	v70 = int64(24)
	v72 = int64(4278190080)
	v74 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v51<<(uint(v61)%64) | v51&v63<<(uint(v65)%64) | (v51&v68<<(uint(v70)%64) | v51&v72<<(uint(v74)%64)) | (int64(base.Ui64(v51)>>(uint(v74)%64))&v72 | int64(base.Ui64(v51)>>(uint(v70)%64))&v68 | (int64(base.Ui64(v51)>>(uint(v65)%64))&v63 | int64(base.Ui64(v51)>>(uint(v61)%64))))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v98 = int32(8)
	v102 = F_raxInsert(m, v97, v9+v98, v98, v45, int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_rebaseReplicationBuffer[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = int32(0)
	goto L12
L15:
	;
	if v109 != 0 {
		v45 = v109
		goto L10
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v109+base.B2i32(v112 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v118
	goto L16
L18:
	;
	goto L11
}
func F_receiveRDBinBioThreadSingleChannel(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_receiveRDBinBioThreadSingleChannel[0]))
	if int32(2) < v4 {
		v12 = int32(_a_F_receiveRDBinBioThreadSingleChannel_0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_receiveRDBinBioThreadSingleChannel[1]))
		v15 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_receiveRDBinBioThreadSingleChannel[1])) = v15
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
		v20 = m.T0[v19].(func(*base.Module, int32, int32) int32)(m, v13, v15)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			F_bioCreateSaveRDBToDiskJob(m, v13, int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F__serverLog(m, int32(2), int32(_a_F_receiveRDBinBioThreadSingleChannel_1), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = int32(_a_F_receiveRDBinBioThreadSingleChannel_0)
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_receiveRDBinBioThreadSingleChannel[1]))
			v15 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_receiveRDBinBioThreadSingleChannel[1])) = v15
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
			v20 = m.T0[v19].(func(*base.Module, int32, int32) int32)(m, v13, v15)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_bioCreateSaveRDBToDiskJob(m, v13, int32(0))
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
}
func F_recvfrom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v7 = m.Env.X__syscall_recvfrom(m, l0, l1, l2, l3, l4, l5)
	mBase = m.M
	if base.Ui32(v7) < base.Ui32(int32(-4095)) {
		v15 = v7
	} else {
		v10 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0) - v7
		v15 = int32(-1)
	}
	return v15
}
func F_refreshGoodReplicasCount(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_refreshGoodReplicasCount[0]))
	if v9 == int32(0) {
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_refreshGoodReplicasCount[1]))
		if v13 == int32(0) {
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_refreshGoodReplicasCount[2]))
			v19 = v6 + int32(8)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v20
			v24 = int32(0)
			v26 = v6 + int32(8)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			if v28 == v24 {
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v28+base.B2i32(v31 == int32(0))<<(uint(int32(2))%32))))
				*(*int32)(unsafe.Add(mBase, uint32(v26))) = v37
			}
			if v28 == int32(0) {
				v73 = v24
			} else {
				v42 = v24
				v43 = v28
				for {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+104))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
					if v46 != int32(9) {
						v57 = v42
					} else {
						v49 = int32(_a_F_refreshGoodReplicasCount_0)
						v50 = *(*int64)(unsafe.Add(mBase, _c_F_refreshGoodReplicasCount[3]))
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v45)+80))
						v54 = int64(*(*int32)(unsafe.Add(mBase, _c_F_refreshGoodReplicasCount[1])))
						v57 = v42 + base.B2i32(v50-v51 <= v54)
					}
					v59 = v6 + int32(8)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					if v61 == int32(0) {
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v61+base.B2i32(v64 == int32(0))<<(uint(int32(2))%32))))
						*(*int32)(unsafe.Add(mBase, uint32(v59))) = v70
					}
					if v61 != 0 {
						v42 = v57
						v43 = v61
						continue
					} else {
						break
					}
					break
				}
				v73 = v57
			}
			*(*int32)(unsafe.Add(mBase, _c_F_refreshGoodReplicasCount[4])) = v73
		}
	}
	m.G0 = v6 + int32(16)
	return
}
func F_rejectCommandSds(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	F_flagTransaction(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v7
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		if v9 == v7 {
		} else {
			v12 = *(*int64)(unsafe.Add(mBase, uint32(v9)+120))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+120)) = v12 + int64(1)
		}
		if l2 == int32(0) {
			v21 = v9
			if v21 == int32(0) {
				F_addReplyErrorSds(m, l0, l1)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					return
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
				if v24 != int32(17) {
					F_addReplyErrorSds(m, l0, l1)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						return
					}
				} else {
					F_execCommandAbort(m, l0, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_sdsfree(m, l1)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			F_moduleFireCommandRejectedEvent(m, l0, l1)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				v21 = v20
				if v21 == int32(0) {
					F_addReplyErrorSds(m, l0, l1)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						return
					}
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
					if v24 != int32(17) {
						F_addReplyErrorSds(m, l0, l1)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							return
						}
					} else {
						F_execCommandAbort(m, l0, l1)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							F_sdsfree(m, l1)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
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
func F_removeFromBucket_VECTOR(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	switch l0 + int32(1) {
	case 0:
		goto L7
	case 1:
		goto L9
	default:
		goto L8
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_removeFromBucket_VECTOR_0), int32(_a_F_removeFromBucket_VECTOR_1), int32(816))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L10
	} else {
		goto L60
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_removeFromBucket_VECTOR_0), int32(_a_F_removeFromBucket_VECTOR_1), int32(816))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L10
	} else {
		goto L59
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_removeFromBucket_VECTOR_2), int32(_a_F_removeFromBucket_VECTOR_1), int32(600))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L10
	} else {
		goto L58
	}
L4:
	;
	F__serverAssert(m, int32(_a_F_removeFromBucket_VECTOR_3), int32(_a_F_removeFromBucket_VECTOR_1), int32(644))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L10
	} else {
		goto L57
	}
L5:
	;
	F__serverAssert(m, int32(_a_F_removeFromBucket_VECTOR_4), int32(_a_F_removeFromBucket_VECTOR_1), int32(683))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L10
	} else {
		goto L56
	}
L6:
	;
	F__serverAssert(m, int32(_a_F_removeFromBucket_VECTOR_2), int32(_a_F_removeFromBucket_VECTOR_1), int32(600))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L10
	} else {
		goto L55
	}
L7:
	;
	F__serverAssert(m, int32(_a_F_removeFromBucket_VECTOR_5), int32(_a_F_removeFromBucket_VECTOR_1), int32(797))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L10
	} else {
		goto L54
	}
L8:
	;
	if l0&int32(7) != int32(2) {
		goto L7
	} else {
		goto L12
	}
L9:
	;
	F__serverAssert(m, int32(_a_F_removeFromBucket_VECTOR_6), int32(_a_F_removeFromBucket_VECTOR_1), int32(781))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v23 = int32(0)
	v25 = l0 & int32(-8)
	if v25 == v23 {
		v176 = l0
		v180 = v23
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v180)
	return v176
L14:
	;
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	v31 = base.I32_wrap_i64(v28) & int32(1073741823)
	if base.Ui32(int32(2)) < base.Ui32(v31) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if l3 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	if v31 == int32(0) {
		v176 = l0
		v180 = v23
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v37 = base.B2i32(v36 == l1)
	if v36 == l1 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v31 == int32(1) {
		v50 = int32(-1)
		goto L22
	} else {
		goto L23
	}
L19:
	;
	if v31 == int32(1) {
		v176 = l0
		v180 = v23
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v40 != l1 {
		v176 = l0
		v180 = v23
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	F_valkey_free(m, v25)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	if base.Ui32(v31) <= base.Ui32(v37) {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v25+v37<<(uint(int32(2))%32))+8))
	v50 = v49
	goto L22
L25:
	;
	v53 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v53)
	return v50
L26:
	;
	v176 = v167
	v180 = int32(1)
	goto L13
L27:
	;
	if v28&int64(1073741823) == int64(0) {
		v176 = l0
		v180 = v23
		goto L13
	} else {
		goto L44
	}
L28:
	;
	v64 = int32(0)
	goto L30
L29:
	;
	if base.Ui32(v64) < base.Ui32(v31) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(8)+v64<<(uint(int32(2))%32))))
	if v72 == l1 {
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v77 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v77)
	return l0
L32:
	;
	v75 = v64 + int32(1)
	if v75 != v31 {
		v64 = v75
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if v28&int64(1073741823) == int64(0) {
		goto L5
	} else {
		goto L36
	}
L35:
	;
	v81 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v81)
	return l0
L36:
	;
	v89 = v25 + int32(8)
	v90 = int32(2)
	v92 = v89 + v64<<(uint(v90)%32)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v98 = v31<<(uint(v90)%32) + v89 + int32(-4)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v93
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	if v102&int64(1073741823) == int64(0) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v109 = base.I32_wrap_i64(v102) & int32(1073741823)
	if v109 == int32(0) {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v113 = v109 + int32(-1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v89+v113<<(uint(int32(2))%32))))
	v118 = F_pvRemoveAt(m, v25, v113)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	if v118 == int32(0) {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	if v117 != l1 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F__serverAssert(m, int32(_a_F_removeFromBucket_VECTOR_7), int32(_a_F_removeFromBucket_VECTOR_1), int32(1351))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L10
	} else {
		goto L43
	}
L42:
	;
	v167 = v118 | int32(2)
	goto L26
L43:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v141 = int32(0)
	goto L46
L45:
	;
	if base.Ui32(v141) < base.Ui32(v31) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(8)+v141<<(uint(int32(2))%32))))
	if v149 == l1 {
		goto L45
	} else {
		goto L48
	}
L47:
	;
	v154 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v154)
	return l0
L48:
	;
	v152 = v141 + int32(1)
	if v152 != v31 {
		v141 = v152
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v161 = F_pvRemoveAt(m, v25, v141)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L10
	} else {
		goto L52
	}
L51:
	;
	v158 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v158)
	return l0
L52:
	;
	if v161 == int32(0) {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v167 = v161 | int32(2)
	goto L26
L54:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_rename(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	v3 = int32(-100)
	v5 = m.Env.X__syscall_renameat(m, v3, l0, v3, l1)
	mBase = m.M
	if base.Ui32(v5) < base.Ui32(int32(-4095)) {
		v13 = v5
	} else {
		v8 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0) - v5
		v13 = int32(-1)
	}
	return v13
}
func F_renameCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_renameGenericCommand(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_renamenxCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_renameGenericCommand(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_replicaofCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int64
	_ = v229
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_replicaofCommand[0]))
	if v9 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v6 + int32(32)
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_replicaofCommand[1]))
	if v16 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_addReplyError(m, l0, int32(_a_F_replicaofCommand_0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	goto L1
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = F_objectGetVal(m, v23)
	mBase = m.M
	v25 = int32(_a_F_replicaofCommand_1)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v28 != 0 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	F_addReplyError(m, l0, int32(_a_F_replicaofCommand_2))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_replicaofCommand[2]))
	F_addReply(m, l0, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L78
	}
L10:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v129&int32(2) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L11:
	;
	if v60-v62 != 0 {
		goto L10
	} else {
		goto L23
	}
L12:
	;
	v60 = F_tolower(m, v56)
	mBase = m.M
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v62 = F_tolower(m, v61)
	mBase = m.M
	goto L11
L13:
	;
	v30 = v24
	v31 = v25
	v32 = v28
	goto L16
L14:
	;
	v56 = int32(0)
	v57 = v25
	goto L12
L15:
	;
	v56 = v53 & int32(255)
	v57 = v52
	goto L12
L16:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v34 == int32(0) {
		v52 = v31
		v53 = v32
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v52 = v46
	v53 = int32(0)
	goto L15
L18:
	;
	v38 = v32 & int32(255)
	if v38 == v34 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v45 = int32(1)
	v46 = v31 + v45
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v47 != 0 {
		v30 = v30 + v45
		v31 = v46
		v32 = v47
		goto L16
	} else {
		goto L22
	}
L20:
	;
	v40 = F_tolower(m, v38)
	mBase = m.M
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v42 = F_tolower(m, v41)
	mBase = m.M
	if v40 == v42 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v52 = v31
	v53 = v44
	goto L15
L22:
	;
	goto L17
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v66 = F_objectGetVal(m, v65)
	mBase = m.M
	v67 = int32(_a_F_replicaofCommand_3)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v70 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	if v102-v104 != 0 {
		goto L10
	} else {
		goto L36
	}
L25:
	;
	v102 = F_tolower(m, v98)
	mBase = m.M
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v104 = F_tolower(m, v103)
	mBase = m.M
	goto L24
L26:
	;
	v72 = v66
	v73 = v67
	v74 = v70
	goto L29
L27:
	;
	v98 = int32(0)
	v99 = v67
	goto L25
L28:
	;
	v98 = v95 & int32(255)
	v99 = v94
	goto L25
L29:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v76 == int32(0) {
		v94 = v73
		v95 = v74
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v94 = v88
	v95 = int32(0)
	goto L28
L31:
	;
	v80 = v74 & int32(255)
	if v80 == v76 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v87 = int32(1)
	v88 = v73 + v87
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v89 != 0 {
		v72 = v72 + v87
		v73 = v88
		v74 = v89
		goto L29
	} else {
		goto L35
	}
L33:
	;
	v82 = F_tolower(m, v80)
	mBase = m.M
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v84 = F_tolower(m, v83)
	mBase = m.M
	if v82 == v84 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	v94 = v73
	v95 = v86
	goto L28
L35:
	;
	goto L30
L36:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_replicaofCommand[3]))
	if v107 == int32(0) {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	F_replicationUnsetPrimary(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v112 = F_sdsempty(m)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_replicaofCommand[4]))
	v116 = F_catClientInfoShortString(m, v112, l0, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_replicaofCommand[5]))
	if int32(2) < v119 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_sdsfree(m, v116)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v116
	F__serverLog(m, int32(2), int32(_a_F_replicaofCommand_4), v6)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L9
L45:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	v144 = F_getRangeLongFromObjectOrReply(m, l0, v138, int32(0), int32(65535), v6+int32(28), int32(_a_F_replicaofCommand_5))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	F_addReplyError(m, l0, int32(_a_F_replicaofCommand_6))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	goto L1
L48:
	;
	if v144 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_replicaofCommand[3]))
	if v147 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	v211 = F_objectGetVal(m, v210)
	mBase = m.M
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
	F_replicationSetPrimary(m, v211, v212, int32(0), int32(1))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L71
	}
L51:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v152 = F_objectGetVal(m, v151)
	mBase = m.M
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v155 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	if v187-v189 != 0 {
		goto L50
	} else {
		goto L64
	}
L53:
	;
	v187 = F_tolower(m, v183)
	mBase = m.M
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	v189 = F_tolower(m, v188)
	mBase = m.M
	goto L52
L54:
	;
	v157 = v147
	v158 = v152
	v159 = v155
	goto L57
L55:
	;
	v183 = int32(0)
	v184 = v152
	goto L53
L56:
	;
	v183 = v180 & int32(255)
	v184 = v179
	goto L53
L57:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v161 == int32(0) {
		v179 = v158
		v180 = v159
		goto L56
	} else {
		goto L59
	}
L58:
	;
	v179 = v173
	v180 = int32(0)
	goto L56
L59:
	;
	v165 = v159 & int32(255)
	if v165 == v161 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v172 = int32(1)
	v173 = v158 + v172
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
	if v174 != 0 {
		v157 = v157 + v172
		v158 = v173
		v159 = v174
		goto L57
	} else {
		goto L63
	}
L61:
	;
	v167 = F_tolower(m, v165)
	mBase = m.M
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	v169 = F_tolower(m, v168)
	mBase = m.M
	if v167 == v169 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	v179 = v158
	v180 = v171
	goto L56
L63:
	;
	goto L58
L64:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_replicaofCommand[6]))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
	if v192 != v193 {
		goto L50
	} else {
		goto L65
	}
L65:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_replicaofCommand[5]))
	if int32(2) < v196 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v205 = F_sdsnew(m, int32(_a_F_replicaofCommand_7))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	F__serverLog(m, int32(2), int32(_a_F_replicaofCommand_8), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	F_addReplySds(m, l0, v205)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	goto L1
L71:
	;
	v217 = F_sdsempty(m)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_replicaofCommand[4]))
	v221 = F_catClientInfoShortString(m, v217, l0, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_replicaofCommand[5]))
	if int32(2) < v224 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_sdsfree(m, v221)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L77
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v221
	v229 = *(*int64)(unsafe.Add(mBase, _c_F_replicaofCommand[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v229
	F__serverLog(m, int32(2), int32(_a_F_replicaofCommand_9), v6+int32(16))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	goto L9
L78:
	;
	goto L1
}
func F_representSlotRangeList(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
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
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = F_sdsempty(m)
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
	v13 = v6 + int32(24)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v14
	goto L3
L3:
	;
	v19 = v6 + int32(24)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v21 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v6 + int32(32)
	return v83
L5:
	;
	if v21 == int32(0) {
		v83 = v8
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+base.B2i32(v24 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v30
	goto L6
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v35
	v40 = F_sdscatfmt(m, v8, int32(_a_F_representSlotRangeList_0), v6+int32(16))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v43 = v6 + int32(24)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v45 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v45 == int32(0) {
		v83 = v40
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v45+base.B2i32(v48 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v54
	goto L11
L13:
	;
	v58 = v45
	v60 = v40
	goto L14
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v62
	v65 = F_sdscatfmt(m, v60, int32(_a_F_representSlotRangeList_1), v6)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v83 = v65
	goto L4
L16:
	;
	v68 = v6 + int32(24)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v70 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v70 != 0 {
		v58 = v70
		v60 = v65
		goto L14
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70+base.B2i32(v73 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v79
	goto L18
L20:
	;
	goto L15
}
func F_reqresAppendResponse(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F_reqresSaveClientReplyOffset(m *base.Module, l0 int32) {
	return
}
func F_resetManualFailover(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_resetManualFailover[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_resetManualFailover[1])))
	if v4 == int32(0) {
		v12 = v3
		v13 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_resetManualFailover[2]))) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_resetManualFailover[3]))) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_resetManualFailover[4]))) = int64(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_resetManualFailover[1]))) = v13
		return
	} else {
		F_unpauseActions(m, int32(2))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_resetManualFailover[0]))
			v12 = v11
			v13 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_resetManualFailover[2]))) = v13
			*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_resetManualFailover[3]))) = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_resetManualFailover[4]))) = int64(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_resetManualFailover[1]))) = v13
			return
		}
	}
}
func F_resetSharedQueryBuf(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_resetSharedQueryBuf[0]))
	if v7 != v9 {
		F__serverAssert(m, int32(_a_F_resetSharedQueryBuf_0), int32(_a_F_resetSharedQueryBuf_1), int32(2317))
		mBase = m.M
		v99 = m.ExcPending
		if v99 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-1)))))
		switch v13 & int32(7) {
		case 0:
			v30 = int32(base.Ui32(v13) >> (uint(int32(3)) % 32))
		case 1:
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-3)))))
			v30 = v20
		case 2:
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(-5)))))
			v30 = v23
		case 3:
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-9))))
			v30 = v26
		case 4:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-17))))
			v30 = v29
		default:
			v30 = v2
		}
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v30 == v31 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
			v69 = v9 + int32(-1)
			v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
			switch v70 & int32(7) {
			case 0:
				v73 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v73)
			case 1:
				v77 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-3)))) = uint8(v77)
			case 2:
				v81 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v9+int32(-5)))) = uint16(v81)
			case 3:
				*(*int32)(unsafe.Add(mBase, uint32(v9+int32(-9)))) = int32(0)
			case 4:
				*(*int64)(unsafe.Add(mBase, uint32(v9+int32(-17)))) = int64(0)
			default:
			}
			v91 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v91)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
			return
		} else {
			v33 = int32(0)
			v36 = F_sdsnewlen(m, v33, int32(16384))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_resetSharedQueryBuf[0])) = v36
				v41 = v36 + int32(-1)
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
				switch v42 & int32(7) {
				case 0:
					v45 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v45)
				case 1:
					v49 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v36+int32(-3)))) = uint8(v49)
				case 2:
					v53 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v36+int32(-5)))) = uint16(v53)
				case 3:
					*(*int32)(unsafe.Add(mBase, uint32(v36+int32(-9)))) = int32(0)
				case 4:
					*(*int64)(unsafe.Add(mBase, uint32(v36+int32(-17)))) = int64(0)
				default:
				}
				v63 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v63)
				return
			}
		}
	}
}
func F_resize_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 != int32(-1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L1
L3:
	;
	F__serverAssert(m, int32(_a_F_resize_1_0), int32(_a_F_resize_1_1), int32(750))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L25
	} else {
		goto L63
	}
L4:
	;
	v16 = int32(0)
	v20 = int32(1)
	if base.Ui32(v20) < base.Ui32(l1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v196
L6:
	;
	v23 = l1
	goto L8
L7:
	;
	v23 = v20
	goto L8
L8:
	;
	v25 = v23 * int32(3)
	if base.Ui32(v25) < base.Ui32(int32(33)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v34 = v16
	goto L11
L10:
	;
	v34 = int32(32) - base.I32_clz(int32(base.Ui32(v25+int32(-1))>>(uint(int32(5))%32)))
	goto L11
L11:
	;
	v36 = v34 & int32(255)
	if base.Ui32(int32(12)<<(uint(v36)%32)) < base.Ui32(l1) {
		v196 = v16
		goto L5
	} else {
		goto L12
	}
L12:
	;
	if base.Ui32(int32(25)) < base.Ui32(v34) {
		v196 = v16
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v34&int32(255) == v43 {
		v196 = v16
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v45 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_resize_1[0]))
	if v47 != int32(2) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v52 = int32(64) << (uint(v36) % 32)
	v53 = base.I32_extend8_s(v43)
	if v34 <= v53 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v50 != 0 {
		v196 = v45
		goto L5
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	if l2 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	if v56 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if v53 == int32(-1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v68 = float64(0)
	goto L23
L22:
	;
	v68 = base.F64_mul(base.F64_convert_i32_u(int32(1)<<(uint(v53)%32)), float64(12))
	goto L23
L23:
	;
	v69 = base.F64_div(base.F64_convert_i32_u(v23), v68)
	if base.F64_lt(base.F64_mul(v69, float64(100)), float64(500)) == int32(0) {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v76 = m.T0[v56].(func(*base.Module, int32, float64) int32)(m, v52, v69)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	if v76 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	return int32(0)
L28:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+36))
	if v143 == int32(0) {
		v149 = v142
		goto L45
	} else {
		goto L46
	}
L29:
	;
	v139 = F_valkey_calloc(m, v52)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L25
	} else {
		goto L44
	}
L30:
	;
	v86 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(v52) {
		v132 = v86
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v132 != 0 {
		v141 = v132
		goto L28
	} else {
		goto L43
	}
L32:
	;
	goto L31
L33:
	;
	if v52 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v94 = v52
	goto L36
L35:
	;
	v94 = int32(4)
	goto L36
L36:
	;
	v96 = v94 + int32(8)
	v97 = F_emscripten_builtin_calloc(m, int32(1), v96)
	mBase = m.M
	if v97 == int32(0) {
		v132 = v86
		goto L32
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v94
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_resize_1[1]))
	if v102 != int32(-1) {
		v113 = v102
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if v113 < int32(260) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v105 = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_resize_1[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_resize_1[1])) = v107
	*(*int32)(unsafe.Add(mBase, _c_F_resize_1[2])) = v107 + int32(1)
	v113 = v107
	goto L38
L40:
	;
	v132 = v97 + int32(8)
	goto L32
L41:
	;
	v122 = v113 << (uint(int32(2)) % 32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)+uint32(_c_F_resize_1[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+uint32(_c_F_resize_1[3]))) = v125 + v96
	goto L40
L42:
	;
	v116 = int32(0)
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_resize_1[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_resize_1[4])) = v118 + v96
	goto L40
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
	return int32(0)
L44:
	;
	v141 = v139
	goto L28
L45:
	;
	v150 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v141
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v34)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v150
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v149)+28))
	if v156 == v150 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	m.T0[v143].(func(*base.Module, int32, int32))(m, l0, v52)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L25
	} else {
		goto L47
	}
L47:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v149 = v148
	goto L45
L48:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v161 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	m.T0[v156].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v170 = int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+44)))
	if v172&v170 == int32(0) {
		v196 = v170
		goto L5
	} else {
		goto L57
	}
L52:
	;
	F_rehashingCompleted(m, l0)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L25
	} else {
		goto L56
	}
L53:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v164 != 0 {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v165 != 0 {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	return int32(1)
L57:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v177 == int32(-1) {
		v196 = v170
		goto L5
	} else {
		goto L58
	}
L58:
	;
	goto L59
L59:
	;
	F_rehashStep(m, l0)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L25
	} else {
		goto L61
	}
L60:
	;
	v196 = v170
	goto L5
L61:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v190 != int32(-1) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_resume(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v4 != 0 {
		v11 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v11)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+6)))
		if v16 == v11 {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v26
			v29 = v13
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v33 = base.I32_div_s(v29-v30, int32(24))
			F_luaV_execute(m, l0, v33)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				return
			}
		} else {
			v19 = F_luaD_poscall(m, l0, l1)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v19 == int32(0) {
					v29 = v21
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24
					v29 = v21
				}
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v33 = base.I32_div_s(v29-v30, int32(24))
				F_luaV_execute(m, l0, v33)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v8 = F_luaD_precall(m, l0, l1+int32(-16), int32(-1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			if v8 != 0 {
				return
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = v10
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v33 = base.I32_div_s(v29-v30, int32(24))
				F_luaV_execute(m, l0, v33)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_rpopCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_popGenericCommand(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_rpushCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_pushGenericCommand(m, l0, int32(1), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
