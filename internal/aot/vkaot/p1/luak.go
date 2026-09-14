package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaK_checkstack(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = v4 + l1
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+75)))
	if v5 <= v7 {
		return
	} else {
		if base.Ui32(v5) < base.Ui32(int32(250)) {
			v18 = v6
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+75)) = uint8(v5)
			return
		} else {
			v11 = m.G3
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_luaX_syntaxerror(m, v12, v11+int32(_a2674))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v18 = v17
				*(*uint8)(unsafe.Add(mBase, uint32(v18)+75)) = uint8(v5)
				return
			}
		}
	}
}
func F_luaK_dischargevars(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v4 + int32(-6) {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
		return
	case 1:
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		v16 = F_luaK_code(m, l0, v9<<(uint(int32(23))%32)|int32(4), v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(11)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v16
			return
		}
	case 2:
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
		v28 = F_luaK_code(m, l0, v21<<(uint(int32(14))%32)|int32(5), v27)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(11)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v28
			return
		}
	case 3:
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		if v33&int32(256) != 0 {
		} else {
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
			if v33 < v36 {
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v38 + int32(-1)
			}
		}
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v42&int32(256) != 0 {
			v52 = v42
		} else {
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
			if v42 < v45 {
				v52 = v42
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v47 + int32(-1)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v52 = v51
			}
		}
		v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
		v63 = F_luaK_code(m, l0, v52<<(uint(int32(23))%32)|v55<<(uint(int32(14))%32)|int32(6), v62)
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(11)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v63
			return
		}
	default:
		return
	case 7:
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
		v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v76 = *(*int32)(unsafe.Add(mBase, uint32(v71+v72<<(uint(int32(2))%32))))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(base.Ui32(v76)>>(uint(int32(6))%32)) & int32(255)
		return
	case 8:
		v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
		v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v87 = v83 + v84<<(uint(int32(2))%32)
		v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
		*(*int32)(unsafe.Add(mBase, uint32(v87))) = v88&int32(8388607) | int32(16777216)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(11)
		return
	}
}
func F_luaK_exp2anyreg(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v32 int32
	_ = v32
	F_luaK_dischargevars(m, l0, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v9 != int32(12) {
			F_luaK_exp2nextreg(m, l0, l1)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v30 = l1 + int32(8)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				return v32
			}
		} else {
			v13 = l1 + int32(8)
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			if v14 == v15 {
				v30 = v13
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				return v32
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
				if v17 < v18 {
					F_luaK_exp2nextreg(m, l0, l1)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v30 = l1 + int32(8)
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
						return v32
					}
				} else {
					F_exp2reg(m, l0, l1, v17)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						return v22
					}
				}
			}
		}
	}
}
func F_luaK_fixline(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v4+v5<<(uint(int32(2))%32)+int32(-4)))) = l1
	return
}
func F_luaK_goiftrue(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	F_luaK_dischargevars(m, l0, l1)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v11 + int32(-2) {
	case 0, 2, 3:
		goto L3
	default:
		goto L5
	case 8:
		goto L6
	}
L3:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v122
	if v121 == int32(-1) {
		goto L24
	} else {
		goto L25
	}
L4:
	;
	if v54 == int32(-1) {
		goto L3
	} else {
		goto L12
	}
L5:
	;
	v49 = F_jumponcond(m, l0, l1, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L6:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v19 = v15 + v16<<(uint(int32(2))%32)
	if v16 < int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = base.B2i32(v35&int32(16320) == int32(0))<<(uint(int32(6))%32) | v35&int32(-16321)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v54 = v47
	goto L4
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v35 = v34
	v36 = v19
	goto L7
L9:
	;
	v23 = v19 + int32(-4)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v25 = m.G400
	v29 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25+v24&int32(63)))))
	if v29 < int32(0) {
		v35 = v24
		v36 = v23
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v54 = v49
	goto L4
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v57 == int32(-1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v54
	goto L3
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v64 = v57
	goto L16
L15:
	;
	v88 = v54 + (v64 ^ int32(-1))
	v90 = v88 >> (uint(int32(31)) % 32)
	if base.Ui32(v88^v90-v90) < base.Ui32(int32(131072)) {
		v103 = v73
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v72 = v61 + v64<<(uint(int32(2))%32)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v77 = int32(base.Ui32(v73)>>(uint(int32(14))%32)) + int32(-131071)
	if v77 == int32(-1) {
		goto L15
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	v82 = v64 + v77 + int32(1)
	if v82 != int32(-1) {
		v64 = v82
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v88<<(uint(int32(14))%32) | v103&int32(16383) + int32(2147467264)
	goto L3
L21:
	;
	v95 = m.G3
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v96, v95+int32(_a2673))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v103 = v101
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(-1)
	return
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(-1)
	return
L25:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v126 == int32(-1) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v133 = v126
	goto L28
L27:
	;
	v157 = v121 + (v133 ^ int32(-1))
	v159 = v157 >> (uint(int32(31)) % 32)
	if base.Ui32(v157^v159-v159) < base.Ui32(int32(131072)) {
		v172 = v142
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v141 = v130 + v133<<(uint(int32(2))%32)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v146 = int32(base.Ui32(v142)>>(uint(int32(14))%32)) + int32(-131071)
	if v146 == int32(-1) {
		goto L27
	} else {
		goto L30
	}
L29:
	;
	goto L27
L30:
	;
	v151 = v133 + v146 + int32(1)
	if v151 != int32(-1) {
		v133 = v151
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v157<<(uint(int32(14))%32) | v172&int32(16383) + int32(2147467264)
	goto L24
L33:
	;
	v164 = m.G3
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v165, v164+int32(_a2673))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v172 = v170
	goto L32
}
func F_luaK_indexed(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = F_luaK_exp2RK(m, l0, l2)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(9)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v4
		return
	}
}
func F_luaK_jump(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(-1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v15 = F_luaK_code(m, l0, int32(2147450902), v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v9 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v15
L4:
	;
	if v15 != int32(-1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v30 = v15
	goto L8
L6:
	;
	return v9
L7:
	;
	v52 = v9 + (v30 ^ int32(-1))
	v54 = v52 >> (uint(int32(31)) % 32)
	if base.Ui32(v52^v54-v54) < base.Ui32(int32(131072)) {
		v67 = v37
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v36 = v25 + v30<<(uint(int32(2))%32)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v41 = int32(base.Ui32(v37)>>(uint(int32(14))%32)) + int32(-131071)
	if v41 == int32(-1) {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v46 = v30 + v41 + int32(1)
	if v46 != int32(-1) {
		v30 = v46
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v52<<(uint(int32(14))%32) | v67&int32(16383) + int32(2147467264)
	goto L3
L13:
	;
	v59 = m.G3
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v60, v59+int32(_a2673))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v67 = v65
	goto L12
}
func F_luaK_self(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	F_luaK_dischargevars(m, l0, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v10 != int32(12) {
			F_luaK_exp2nextreg(m, l0, l1)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v25 != int32(12) {
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if v28&int32(256) != 0 {
					} else {
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
						if v28 < v31 {
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v33 + int32(-1)
						}
					}
				}
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v40 = v38 + int32(2)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+75)))
				if v42 < v40 {
					if base.Ui32(v40) < base.Ui32(int32(250)) {
						v56 = v41
						v57 = v40
						*(*uint8)(unsafe.Add(mBase, uint32(v56)+75)) = uint8(v40)
						v60 = v57
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v60
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v68 = F_luaK_exp2RK(m, l0, l2)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
							v77 = F_luaK_code(m, l0, v38<<(uint(int32(6))%32)|v64<<(uint(int32(23))%32)|v68<<(uint(int32(14))%32)|int32(11), v76)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								if v79 != int32(12) {
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									if v82&int32(256) != 0 {
									} else {
										v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
										if v82 < v85 {
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v87 + int32(-1)
										}
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
								return
							}
						}
					} else {
						v46 = m.G3
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						F_luaX_syntaxerror(m, v47, v46+int32(_a2674))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v56 = v55
							v57 = v52 + int32(2)
							*(*uint8)(unsafe.Add(mBase, uint32(v56)+75)) = uint8(v40)
							v60 = v57
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v60
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v68 = F_luaK_exp2RK(m, l0, l2)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
								v77 = F_luaK_code(m, l0, v38<<(uint(int32(6))%32)|v64<<(uint(int32(23))%32)|v68<<(uint(int32(14))%32)|int32(11), v76)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									if v79 != int32(12) {
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										if v82&int32(256) != 0 {
										} else {
											v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
											if v82 < v85 {
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v87 + int32(-1)
											}
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
									return
								}
							}
						}
					}
				} else {
					v60 = v40
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v60
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v68 = F_luaK_exp2RK(m, l0, l2)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
						v77 = F_luaK_code(m, l0, v38<<(uint(int32(6))%32)|v64<<(uint(int32(23))%32)|v68<<(uint(int32(14))%32)|int32(11), v76)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							if v79 != int32(12) {
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								if v82&int32(256) != 0 {
								} else {
									v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
									if v82 < v85 {
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v87 + int32(-1)
									}
								}
							}
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
							return
						}
					}
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			if v13 == v14 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v25 != int32(12) {
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if v28&int32(256) != 0 {
					} else {
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
						if v28 < v31 {
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v33 + int32(-1)
						}
					}
				}
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v40 = v38 + int32(2)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+75)))
				if v42 < v40 {
					if base.Ui32(v40) < base.Ui32(int32(250)) {
						v56 = v41
						v57 = v40
						*(*uint8)(unsafe.Add(mBase, uint32(v56)+75)) = uint8(v40)
						v60 = v57
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v60
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v68 = F_luaK_exp2RK(m, l0, l2)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
							v77 = F_luaK_code(m, l0, v38<<(uint(int32(6))%32)|v64<<(uint(int32(23))%32)|v68<<(uint(int32(14))%32)|int32(11), v76)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								if v79 != int32(12) {
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									if v82&int32(256) != 0 {
									} else {
										v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
										if v82 < v85 {
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v87 + int32(-1)
										}
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
								return
							}
						}
					} else {
						v46 = m.G3
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						F_luaX_syntaxerror(m, v47, v46+int32(_a2674))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v56 = v55
							v57 = v52 + int32(2)
							*(*uint8)(unsafe.Add(mBase, uint32(v56)+75)) = uint8(v40)
							v60 = v57
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v60
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v68 = F_luaK_exp2RK(m, l0, l2)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
								v77 = F_luaK_code(m, l0, v38<<(uint(int32(6))%32)|v64<<(uint(int32(23))%32)|v68<<(uint(int32(14))%32)|int32(11), v76)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									if v79 != int32(12) {
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										if v82&int32(256) != 0 {
										} else {
											v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
											if v82 < v85 {
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v87 + int32(-1)
											}
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
									return
								}
							}
						}
					}
				} else {
					v60 = v40
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v60
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v68 = F_luaK_exp2RK(m, l0, l2)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
						v77 = F_luaK_code(m, l0, v38<<(uint(int32(6))%32)|v64<<(uint(int32(23))%32)|v68<<(uint(int32(14))%32)|int32(11), v76)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							if v79 != int32(12) {
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								if v82&int32(256) != 0 {
								} else {
									v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
									if v82 < v85 {
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v87 + int32(-1)
									}
								}
							}
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
							return
						}
					}
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
				if v16 < v17 {
					F_luaK_exp2nextreg(m, l0, l1)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if v25 != int32(12) {
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							if v28&int32(256) != 0 {
							} else {
								v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
								if v28 < v31 {
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v33 + int32(-1)
								}
							}
						}
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v40 = v38 + int32(2)
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+75)))
						if v42 < v40 {
							if base.Ui32(v40) < base.Ui32(int32(250)) {
								v56 = v41
								v57 = v40
								*(*uint8)(unsafe.Add(mBase, uint32(v56)+75)) = uint8(v40)
								v60 = v57
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v60
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								v68 = F_luaK_exp2RK(m, l0, l2)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
									v77 = F_luaK_code(m, l0, v38<<(uint(int32(6))%32)|v64<<(uint(int32(23))%32)|v68<<(uint(int32(14))%32)|int32(11), v76)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										if v79 != int32(12) {
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											if v82&int32(256) != 0 {
											} else {
												v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
												if v82 < v85 {
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v87 + int32(-1)
												}
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
										return
									}
								}
							} else {
								v46 = m.G3
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								F_luaX_syntaxerror(m, v47, v46+int32(_a2674))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v56 = v55
									v57 = v52 + int32(2)
									*(*uint8)(unsafe.Add(mBase, uint32(v56)+75)) = uint8(v40)
									v60 = v57
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v60
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									v68 = F_luaK_exp2RK(m, l0, l2)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
										v77 = F_luaK_code(m, l0, v38<<(uint(int32(6))%32)|v64<<(uint(int32(23))%32)|v68<<(uint(int32(14))%32)|int32(11), v76)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											if v79 != int32(12) {
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
												if v82&int32(256) != 0 {
												} else {
													v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
													if v82 < v85 {
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v87 + int32(-1)
													}
												}
											}
											*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
											return
										}
									}
								}
							}
						} else {
							v60 = v40
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v60
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v68 = F_luaK_exp2RK(m, l0, l2)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
								v77 = F_luaK_code(m, l0, v38<<(uint(int32(6))%32)|v64<<(uint(int32(23))%32)|v68<<(uint(int32(14))%32)|int32(11), v76)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									if v79 != int32(12) {
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										if v82&int32(256) != 0 {
										} else {
											v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
											if v82 < v85 {
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v87 + int32(-1)
											}
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
									return
								}
							}
						}
					}
				} else {
					F_exp2reg(m, l0, l1, v16)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if v25 != int32(12) {
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							if v28&int32(256) != 0 {
							} else {
								v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
								if v28 < v31 {
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v33 + int32(-1)
								}
							}
						}
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v40 = v38 + int32(2)
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+75)))
						if v42 < v40 {
							if base.Ui32(v40) < base.Ui32(int32(250)) {
								v56 = v41
								v57 = v40
								*(*uint8)(unsafe.Add(mBase, uint32(v56)+75)) = uint8(v40)
								v60 = v57
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v60
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								v68 = F_luaK_exp2RK(m, l0, l2)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
									v77 = F_luaK_code(m, l0, v38<<(uint(int32(6))%32)|v64<<(uint(int32(23))%32)|v68<<(uint(int32(14))%32)|int32(11), v76)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										if v79 != int32(12) {
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											if v82&int32(256) != 0 {
											} else {
												v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
												if v82 < v85 {
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v87 + int32(-1)
												}
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
										return
									}
								}
							} else {
								v46 = m.G3
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								F_luaX_syntaxerror(m, v47, v46+int32(_a2674))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v56 = v55
									v57 = v52 + int32(2)
									*(*uint8)(unsafe.Add(mBase, uint32(v56)+75)) = uint8(v40)
									v60 = v57
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v60
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									v68 = F_luaK_exp2RK(m, l0, l2)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
										v77 = F_luaK_code(m, l0, v38<<(uint(int32(6))%32)|v64<<(uint(int32(23))%32)|v68<<(uint(int32(14))%32)|int32(11), v76)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											if v79 != int32(12) {
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
												if v82&int32(256) != 0 {
												} else {
													v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
													if v82 < v85 {
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v87 + int32(-1)
													}
												}
											}
											*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
											return
										}
									}
								}
							}
						} else {
							v60 = v40
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v60
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v68 = F_luaK_exp2RK(m, l0, l2)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
								v77 = F_luaK_code(m, l0, v38<<(uint(int32(6))%32)|v64<<(uint(int32(23))%32)|v68<<(uint(int32(14))%32)|int32(11), v76)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									if v79 != int32(12) {
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										if v82&int32(256) != 0 {
										} else {
											v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
											if v82 < v85 {
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v87 + int32(-1)
											}
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_luaK_setoneret(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v3 + int32(-13) {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v9+v10<<(uint(int32(2))%32))))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(base.Ui32(v14)>>(uint(int32(6))%32)) & int32(255)
		return
	case 1:
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v25 = v21 + v22<<(uint(int32(2))%32)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
		*(*int32)(unsafe.Add(mBase, uint32(v25))) = v26&int32(8388607) | int32(16777216)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(11)
		return
	default:
		return
	}
}
func F_luaK_storevar(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v6 + int32(-6) {
	case 0:
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v9 != int32(12) {
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			if v12&int32(256) != 0 {
			} else {
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
				if v12 < v15 {
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v17 + int32(-1)
				}
			}
		}
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F_exp2reg(m, l0, l2, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			return
		}
	case 1:
		F_luaK_dischargevars(m, l0, l2)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			if v27 != int32(12) {
				F_luaK_exp2nextreg(m, l0, l2)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v46 = l2 + int32(8)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v103 = v48<<(uint(int32(6))%32) | v51<<(uint(int32(23))%32) | int32(8)
					v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
					v108 = F_luaK_code(m, l0, v103, v107)
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return
					} else {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						if v113 != int32(12) {
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							if v116&int32(256) != 0 {
							} else {
								v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
								if v116 < v119 {
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v121 + int32(-1)
								}
							}
						}
						return
					}
				}
			} else {
				v31 = l2 + int32(8)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				if v32 == v33 {
					v46 = v31
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v103 = v48<<(uint(int32(6))%32) | v51<<(uint(int32(23))%32) | int32(8)
					v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
					v108 = F_luaK_code(m, l0, v103, v107)
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return
					} else {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						if v113 != int32(12) {
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							if v116&int32(256) != 0 {
							} else {
								v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
								if v116 < v119 {
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v121 + int32(-1)
								}
							}
						}
						return
					}
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
					if v35 < v36 {
						F_luaK_exp2nextreg(m, l0, l2)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v46 = l2 + int32(8)
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v103 = v48<<(uint(int32(6))%32) | v51<<(uint(int32(23))%32) | int32(8)
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
							v108 = F_luaK_code(m, l0, v103, v107)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								if v113 != int32(12) {
								} else {
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									if v116&int32(256) != 0 {
									} else {
										v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
										if v116 < v119 {
										} else {
											v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v121 + int32(-1)
										}
									}
								}
								return
							}
						}
					} else {
						F_exp2reg(m, l0, l2, v35)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v46 = v31
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v103 = v48<<(uint(int32(6))%32) | v51<<(uint(int32(23))%32) | int32(8)
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
							v108 = F_luaK_code(m, l0, v103, v107)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								if v113 != int32(12) {
								} else {
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									if v116&int32(256) != 0 {
									} else {
										v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
										if v116 < v119 {
										} else {
											v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v121 + int32(-1)
										}
									}
								}
								return
							}
						}
					}
				}
			}
		}
	case 2:
		F_luaK_dischargevars(m, l0, l2)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			if v59 != int32(12) {
				F_luaK_exp2nextreg(m, l0, l2)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					v78 = l2 + int32(8)
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v103 = v80<<(uint(int32(6))%32) | v83<<(uint(int32(14))%32) | int32(7)
					v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
					v108 = F_luaK_code(m, l0, v103, v107)
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return
					} else {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						if v113 != int32(12) {
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							if v116&int32(256) != 0 {
							} else {
								v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
								if v116 < v119 {
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v121 + int32(-1)
								}
							}
						}
						return
					}
				}
			} else {
				v63 = l2 + int32(8)
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				if v64 == v65 {
					v78 = v63
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v103 = v80<<(uint(int32(6))%32) | v83<<(uint(int32(14))%32) | int32(7)
					v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
					v108 = F_luaK_code(m, l0, v103, v107)
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return
					} else {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						if v113 != int32(12) {
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							if v116&int32(256) != 0 {
							} else {
								v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
								if v116 < v119 {
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v121 + int32(-1)
								}
							}
						}
						return
					}
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
					if v67 < v68 {
						F_luaK_exp2nextreg(m, l0, l2)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							v78 = l2 + int32(8)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v103 = v80<<(uint(int32(6))%32) | v83<<(uint(int32(14))%32) | int32(7)
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
							v108 = F_luaK_code(m, l0, v103, v107)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								if v113 != int32(12) {
								} else {
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									if v116&int32(256) != 0 {
									} else {
										v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
										if v116 < v119 {
										} else {
											v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v121 + int32(-1)
										}
									}
								}
								return
							}
						}
					} else {
						F_exp2reg(m, l0, l2, v67)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							v78 = v63
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v103 = v80<<(uint(int32(6))%32) | v83<<(uint(int32(14))%32) | int32(7)
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
							v108 = F_luaK_code(m, l0, v103, v107)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								if v113 != int32(12) {
								} else {
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									if v116&int32(256) != 0 {
									} else {
										v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
										if v116 < v119 {
										} else {
											v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v121 + int32(-1)
										}
									}
								}
								return
							}
						}
					}
				}
			}
		}
	case 3:
		v89 = F_luaK_exp2RK(m, l0, l2)
		mBase = m.M
		v90 = m.ExcPending
		if v90 != 0 {
			return
		} else {
			v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v103 = v89<<(uint(int32(14))%32) | v93<<(uint(int32(6))%32) | v97<<(uint(int32(23))%32) | int32(9)
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
			v108 = F_luaK_code(m, l0, v103, v107)
			mBase = m.M
			v109 = m.ExcPending
			if v109 != 0 {
				return
			} else {
				v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				if v113 != int32(12) {
				} else {
					v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					if v116&int32(256) != 0 {
					} else {
						v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
						if v116 < v119 {
						} else {
							v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v121 + int32(-1)
						}
					}
				}
				return
			}
		}
	default:
		v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v113 != int32(12) {
		} else {
			v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			if v116&int32(256) != 0 {
			} else {
				v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
				if v116 < v119 {
				} else {
					v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v121 + int32(-1)
				}
			}
		}
		return
	}
}
func F_luaK_stringK(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
	v11 = F_addk(m, l0, v6, v6)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v11
	}
}
