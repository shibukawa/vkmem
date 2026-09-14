package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_rioCheckType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v2 != int32(981) {
		if v2 != int32(982) {
			if v2 == int32(983) {
				v15 = int32(4)
			} else {
				v15 = int32(8)
			}
			return v15
		} else {
			return int32(2)
		}
	} else {
		return int32(1)
	}
}
func F_rioConnFlush(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F_rioConnsetFlush(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_rioConnsetWrite(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_rioConnsetTell(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	return v2
}
func F_rioFdRead(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return int32(0)
}
func F_rioFileRead(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v6 = F_fread(m, l1, l2, int32(1), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_rioFreeConn(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
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
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	if l1 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
	return
L2:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	F_sdsfree(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L48
	} else {
		goto L50
	}
L3:
	;
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(-1)))))
	switch v13 & int32(7) {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	case 3:
		goto L7
	case 4:
		goto L6
	default:
		goto L4
	}
L4:
	;
	F_sdsfree(m, v10)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L48
	} else {
		goto L49
	}
L5:
	;
	v31 = base.I32_wrap_i64(v9)
	if base.Ui32(v30) <= base.Ui32(v31) {
		goto L4
	} else {
		goto L11
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-17))))
	v30 = v29
	goto L5
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-9))))
	v30 = v26
	goto L5
L8:
	;
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(-5)))))
	v30 = v23
	goto L5
L9:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(-3)))))
	v30 = v20
	goto L5
L10:
	;
	v30 = int32(base.Ui32(v13) >> (uint(int32(3)) % 32))
	goto L5
L11:
	;
	if v9 < int64(1) {
		v137 = v10
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v137
	goto L1
L13:
	;
	v35 = int32(-1)
	v43 = v10 + v35
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v46 = v44 & int32(7)
	switch v46 {
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
		goto L15
	}
L14:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v137 = v136
	goto L12
L15:
	;
	goto L14
L16:
	;
	if v61 == int32(0) {
		goto L15
	} else {
		goto L22
	}
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-17))))
	v61 = v60
	goto L16
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-9))))
	v61 = v57
	goto L16
L19:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(-5)))))
	v61 = v54
	goto L16
L20:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(-3)))))
	v61 = v51
	goto L16
L21:
	;
	v61 = int32(base.Ui32(v44) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	v67 = int32(-1)&v61 + v35
	v71 = v31>>(uint(int32(31))%32)&v61 + v31
	v74 = v67 - v71 + int32(1)
	switch v46 {
	default:
		goto L28
	case 1:
		goto L27
	case 2:
		goto L26
	case 3:
		goto L25
	case 4:
		goto L24
	}
L23:
	;
	v90 = int32(0)
	v92 = base.B2i32(base.Ui32(v71) < base.Ui32(v89))
	if base.Ui32(v71) < base.Ui32(v89) {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-17))))
	v89 = v88
	goto L23
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-9))))
	v89 = v85
	goto L23
L26:
	;
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(-5)))))
	v89 = v82
	goto L23
L27:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(-3)))))
	v89 = v79
	goto L23
L28:
	;
	v89 = int32(base.Ui32(v44) >> (uint(int32(3)) % 32))
	goto L23
L29:
	;
	v106 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10+v100))) = uint8(v106)
	switch v46 {
	default:
		goto L47
	case 1:
		goto L46
	case 2:
		goto L45
	case 3:
		goto L44
	case 4:
		goto L43
	}
L30:
	;
	v93 = v71
	goto L32
L31:
	;
	v93 = v90
	goto L32
L32:
	;
	v94 = v89 - v93
	if base.Ui32(v74) < base.Ui32(v94) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = v74
	goto L35
L34:
	;
	v96 = v94
	goto L35
L35:
	;
	if v67 < v71 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v98 = v90
	goto L38
L37:
	;
	v98 = v96
	goto L38
L38:
	;
	if base.Ui32(v71) < base.Ui32(v89) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v100 = v98
	goto L41
L40:
	;
	v100 = int32(0)
	goto L41
L41:
	;
	if v100 == int32(0) {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	v104 = F_memmove(m, v10, v10+v93, v100)
	mBase = m.M
	goto L29
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(-17)))) = base.I64_extend_i32_u(v100)
	goto L15
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(-9)))) = v100
	goto L14
L45:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(-5)))) = uint16(v100)
	goto L14
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(-3)))) = uint8(v100)
	goto L14
L47:
	;
	v109 = v100 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v109)
	goto L14
L48:
	;
	return
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L1
L50:
	;
	goto L1
}
func F_rioFreeConnset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_valkey_free(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		F_valkey_free(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			F_sdsfree(m, v8)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_rioGenericUpdateChecksum(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v4&int32(8) != 0 {
	} else {
		v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v9 = F_crc64(m, v7, l1, base.I64_extend_i32_u(l2))
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v9
	}
	return
}
func F_rioInitWithConn(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	v8 = F__emscripten_memcpy_bulkmem(m, l0, int32(_a1070), int32(80))
	mBase = m.M
	v10 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = l1
	v18 = F_sdsnewlen(m, v10, int32(16384))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v18
		v23 = v18 + int32(-1)
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
		switch v24 & int32(7) {
		case 0:
			v27 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v27)
		case 1:
			v31 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(-3)))) = uint8(v31)
		case 2:
			v35 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v18+int32(-5)))) = uint16(v35)
		case 3:
			*(*int32)(unsafe.Add(mBase, uint32(v18+int32(-9)))) = int32(0)
		case 4:
			*(*int64)(unsafe.Add(mBase, uint32(v18+int32(-17)))) = int64(0)
		default:
		}
		v45 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v45)
		return
	}
}
func F_rioInitWithFd(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v7 = F__emscripten_memcpy_bulkmem(m, l0, int32(_a1071), int32(80))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l1
	v12 = F_sdsempty(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v12
		return
	}
}
func F_rioSetReclaimCache(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	v8 = v3&int32(254) | l1&int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v8)
	return
}
func F_rioWriteBulkStreamID(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_sdsempty(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = v15
		v21 = F_sdscatfmt(m, v11, int32(_a105), v9)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+int32(-1)))))
			switch v25 & int32(7) {
			case 0:
				v42 = int32(base.Ui32(v25) >> (uint(int32(3)) % 32))
			case 1:
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+int32(-3)))))
				v42 = v32
			case 2:
				v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+int32(-5)))))
				v42 = v35
			case 3:
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(-9))))
				v42 = v38
			case 4:
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(-17))))
				v42 = v41
			default:
				v42 = int32(0)
			}
			v43 = F_rioWriteBulkString(m, l0, v21, v42)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				F_sdsfree(m, v21)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(16)
					return v43
				}
			}
		}
	}
}
func F_rioWriteBulkString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v19 int64
	_ = v19
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v207 int32
	_ = v207
	v10 = m.G0
	v12 = v10 - int32(128)
	m.G0 = v12
	v14 = int32(36)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v14)
	v17 = v12 | int32(1)
	v19 = base.I64_extend_i32_s(l2)
	if v19 <= int64(-1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v64 = int32(2573)
	*(*uint16)(unsafe.Add(mBase, uint32(v60+v12+int32(1)))) = uint16(v64)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v66&int32(6) != 0 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v60 = int32(0)
	goto L1
L4:
	;
	v41 = F_ull2string(m, v37, v38, v39)
	mBase = m.M
	if v41 == int32(0) {
		goto L2
	} else {
		goto L8
	}
L5:
	;
	goto L7
L6:
	;
	v37 = v17
	v38 = int32(127)
	v39 = v19
	v40 = int32(0)
	goto L4
L7:
	;
	v28 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v28)
	v32 = int32(1)
	v37 = v17 + v32
	v38 = int32(126)
	v39 = int64(0) - v19
	v40 = v32
	goto L4
L8:
	;
	v60 = v41 + v40
	goto L1
L10:
	;
	m.G0 = v12 + int32(128)
	return v207
L11:
	;
	v207 = int32(0)
	goto L10
L12:
	;
	v70 = v60 + int32(3)
	if v70 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v79 = v12
	v80 = v70
	goto L14
L14:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v82) < base.Ui32(v80) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if l2 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L16:
	;
	v84 = v82
	goto L18
L17:
	;
	v84 = v80
	goto L18
L18:
	;
	if v82 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v85 = v84
	goto L21
L20:
	;
	v85 = v80
	goto L21
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v86 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v94 = m.T0[v93].(func(*base.Module, int32, int32, int32) int32)(m, l0, v79, v85)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L24
	} else {
		goto L27
	}
L23:
	;
	m.T0[v86].(func(*base.Module, int32, int32, int32))(m, l0, v79, v85)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	goto L22
L26:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v100 + v85
	v104 = v80 - v85
	if v104 != 0 {
		v79 = v79 + v85
		v80 = v104
		goto L14
	} else {
		goto L29
	}
L27:
	;
	if v94 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v96 | int64(2)
	goto L11
L29:
	;
	goto L15
L30:
	;
	v150 = int32(0)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v151&int32(6) != 0 {
		v207 = v150
		goto L10
	} else {
		goto L48
	}
L31:
	;
	v107 = int32(0)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v108&int32(6) != 0 {
		v207 = v107
		goto L10
	} else {
		goto L32
	}
L32:
	;
	v112 = l1
	v118 = l2
	goto L33
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v120) < base.Ui32(v118) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L30
L35:
	;
	v122 = v120
	goto L37
L36:
	;
	v122 = v118
	goto L37
L37:
	;
	if v120 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v123 = v122
	goto L40
L39:
	;
	v123 = v118
	goto L40
L40:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v124 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v130 = m.T0[v129].(func(*base.Module, int32, int32, int32) int32)(m, l0, v112, v123)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L24
	} else {
		goto L45
	}
L42:
	;
	m.T0[v124].(func(*base.Module, int32, int32, int32))(m, l0, v112, v123)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L24
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v136 + v123
	v140 = v118 - v123
	if v140 != 0 {
		v112 = v112 + v123
		v118 = v140
		goto L33
	} else {
		goto L47
	}
L45:
	;
	if v130 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v132 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v132 | int64(2)
	v207 = v107
	goto L10
L47:
	;
	goto L34
L48:
	;
	v162 = int32(_a727)
	v163 = int32(2)
	goto L49
L49:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v165) < base.Ui32(v163) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v207 = l2 + v70 + int32(2)
	goto L10
L51:
	;
	v167 = v165
	goto L53
L52:
	;
	v167 = v163
	goto L53
L53:
	;
	if v165 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v168 = v167
	goto L56
L55:
	;
	v168 = v163
	goto L56
L56:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v169 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v175 = m.T0[v174].(func(*base.Module, int32, int32, int32) int32)(m, l0, v162, v168)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L24
	} else {
		goto L61
	}
L58:
	;
	m.T0[v169].(func(*base.Module, int32, int32, int32))(m, l0, v162, v168)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L24
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v181 + v168
	v185 = v163 - v168
	if v185 != 0 {
		v162 = v162 + v168
		v163 = v185
		goto L49
	} else {
		goto L63
	}
L61:
	;
	if v175 != 0 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v177 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v177 | int64(2)
	v207 = v150
	goto L10
L63:
	;
	goto L50
}
func F_rioWriteStreamEmptyConsumer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v6 = int32(0)
	v10 = F_rioWriteBulkCount(m, l0, int32(42), int32(5))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			v64 = v6
			return v64
		} else {
			v18 = F_rioWriteBulkString(m, l0, int32(_a112), int32(6))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					v64 = v6
					return v64
				} else {
					v24 = F_rioWriteBulkString(m, l0, int32(_a113), int32(14))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						if v24 == int32(0) {
							v64 = v6
							return v64
						} else {
							v28 = F_rioWriteBulkObject(m, l0, l1)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								if v28 == int32(0) {
									v64 = v6
									return v64
								} else {
									v32 = F_rioWriteBulkString(m, l0, l2, l3)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return int32(0)
									} else {
										if v32 == int32(0) {
											v64 = v6
											return v64
										} else {
											v37 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
											v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-1)))))
											switch v40 & int32(7) {
											case 0:
												v57 = int32(base.Ui32(v40) >> (uint(int32(3)) % 32))
											case 1:
												v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-3)))))
												v57 = v47
											case 2:
												v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(-5)))))
												v57 = v50
											case 3:
												v53 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-9))))
												v57 = v53
											case 4:
												v56 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-17))))
												v57 = v56
											default:
												v57 = int32(0)
											}
											v58 = F_rioWriteBulkString(m, l0, v37, v57)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int32(0)
											} else {
												v64 = base.B2i32(v58 != int32(0))
												return v64
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
func F_rioWriteStreamPendingEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int64
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int64
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = v11 + int32(16)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l5)+8))
	v19 = int64(56)
	v21 = int64(65280)
	v23 = int64(40)
	v26 = int64(16711680)
	v28 = int64(24)
	v30 = int64(4278190080)
	v32 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v18<<(uint(v19)%64) | v18&v21<<(uint(v23)%64) | (v18&v26<<(uint(v28)%64) | v18&v30<<(uint(v32)%64)) | (int64(base.Ui64(v18)>>(uint(v32)%64))&v30 | int64(base.Ui64(v18)>>(uint(v28)%64))&v26 | (int64(base.Ui64(v18)>>(uint(v23)%64))&v21 | int64(base.Ui64(v18)>>(uint(v19)%64))))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v17<<(uint(v19)%64) | v17&v21<<(uint(v23)%64) | (v17&v26<<(uint(v28)%64) | v17&v30<<(uint(v32)%64)) | (int64(base.Ui64(v17)>>(uint(v32)%64))&v30 | int64(base.Ui64(v17)>>(uint(v28)%64))&v26 | (int64(base.Ui64(v17)>>(uint(v23)%64))&v21 | int64(base.Ui64(v17)>>(uint(v19)%64))))
	v91 = int32(0)
	v94 = F_rioWriteBulkCount(m, l0, int32(42), int32(12))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		return int32(0)
	} else {
		if v94 == int32(0) {
			v221 = v91
			m.G0 = v11 + int32(32)
			return v221
		} else {
			v102 = F_rioWriteBulkString(m, l0, int32(_a106), int32(6))
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return int32(0)
			} else {
				if v102 == int32(0) {
					v221 = v91
					m.G0 = v11 + int32(32)
					return v221
				} else {
					v106 = F_rioWriteBulkObject(m, l0, l1)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						if v106 == int32(0) {
							v221 = v91
							m.G0 = v11 + int32(32)
							return v221
						} else {
							v110 = F_rioWriteBulkString(m, l0, l2, l3)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								if v110 == int32(0) {
									v221 = v91
									m.G0 = v11 + int32(32)
									return v221
								} else {
									v114 = int32(0)
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
									v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+int32(-1)))))
									switch v119 & int32(7) {
									case 0:
										v136 = int32(base.Ui32(v119) >> (uint(int32(3)) % 32))
									case 1:
										v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+int32(-3)))))
										v136 = v126
									case 2:
										v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116+int32(-5)))))
										v136 = v129
									case 3:
										v132 = *(*int32)(unsafe.Add(mBase, uint32(v116+int32(-9))))
										v136 = v132
									case 4:
										v135 = *(*int32)(unsafe.Add(mBase, uint32(v116+int32(-17))))
										v136 = v135
									default:
										v136 = v114
									}
									v137 = F_rioWriteBulkString(m, l0, v116, v136)
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return int32(0)
									} else {
										if v137 == int32(0) {
											v221 = v114
											m.G0 = v11 + int32(32)
											return v221
										} else {
											v143 = F_rioWriteBulkString(m, l0, int32(_a107), int32(1))
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return int32(0)
											} else {
												if v143 == int32(0) {
													v221 = v114
													m.G0 = v11 + int32(32)
													return v221
												} else {
													v147 = F_sdsempty(m)
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														v149 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
														*(*int64)(unsafe.Add(mBase, uint32(v11))) = v149
														v151 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
														*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v151
														v153 = int32(0)
														v156 = F_sdscatfmt(m, v147, int32(_a105), v11)
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+int32(-1)))))
															switch v160 & int32(7) {
															case 0:
																v177 = int32(base.Ui32(v160) >> (uint(int32(3)) % 32))
															case 1:
																v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+int32(-3)))))
																v177 = v167
															case 2:
																v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156+int32(-5)))))
																v177 = v170
															case 3:
																v173 = *(*int32)(unsafe.Add(mBase, uint32(v156+int32(-9))))
																v177 = v173
															case 4:
																v176 = *(*int32)(unsafe.Add(mBase, uint32(v156+int32(-17))))
																v177 = v176
															default:
																v177 = v153
															}
															v178 = F_rioWriteBulkString(m, l0, v156, v177)
															mBase = m.M
															v179 = m.ExcPending
															if v179 != 0 {
																return int32(0)
															} else {
																F_sdsfree(m, v156)
																mBase = m.M
																v181 = m.ExcPending
																if v181 != 0 {
																	return int32(0)
																} else {
																	if v178 == int32(0) {
																		v221 = v153
																		m.G0 = v11 + int32(32)
																		return v221
																	} else {
																		v186 = F_rioWriteBulkString(m, l0, int32(_a108), int32(4))
																		mBase = m.M
																		v187 = m.ExcPending
																		if v187 != 0 {
																			return int32(0)
																		} else {
																			if v186 == int32(0) {
																				v221 = v153
																				m.G0 = v11 + int32(32)
																				return v221
																			} else {
																				v190 = *(*int64)(unsafe.Add(mBase, uint32(l6)))
																				v191 = F_rioWriteBulkLongLong(m, l0, v190)
																				mBase = m.M
																				v192 = m.ExcPending
																				if v192 != 0 {
																					return int32(0)
																				} else {
																					if v191 == int32(0) {
																						v221 = v153
																						m.G0 = v11 + int32(32)
																						return v221
																					} else {
																						v197 = F_rioWriteBulkString(m, l0, int32(_a109), int32(10))
																						mBase = m.M
																						v198 = m.ExcPending
																						if v198 != 0 {
																							return int32(0)
																						} else {
																							if v197 == int32(0) {
																								v221 = v153
																								m.G0 = v11 + int32(32)
																								return v221
																							} else {
																								v201 = *(*int64)(unsafe.Add(mBase, uint32(l6)+8))
																								v202 = F_rioWriteBulkLongLong(m, l0, v201)
																								mBase = m.M
																								v203 = m.ExcPending
																								if v203 != 0 {
																									return int32(0)
																								} else {
																									if v202 == int32(0) {
																										v221 = v153
																										m.G0 = v11 + int32(32)
																										return v221
																									} else {
																										v208 = F_rioWriteBulkString(m, l0, int32(_a110), int32(6))
																										mBase = m.M
																										v209 = m.ExcPending
																										if v209 != 0 {
																											return int32(0)
																										} else {
																											if v208 == int32(0) {
																												v221 = v153
																												m.G0 = v11 + int32(32)
																												return v221
																											} else {
																												v214 = F_rioWriteBulkString(m, l0, int32(_a111), int32(5))
																												mBase = m.M
																												v215 = m.ExcPending
																												if v215 != 0 {
																													return int32(0)
																												} else {
																													v221 = base.B2i32(v214 != int32(0))
																													m.G0 = v11 + int32(32)
																													return v221
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
				}
			}
		}
	}
}
