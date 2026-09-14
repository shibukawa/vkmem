package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_hashtableChannelsGetKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(44))))
	return v4
}
func F_hashtableClientHash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	return v2
}
func F_hashtableCommandGetCurrentName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	return v2
}
func F_hashtableDelete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_hashtablePop(m, l0, l1, v6+int32(12))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			m.G0 = v6 + int32(16)
			return v10
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
			if v17 == int32(0) {
				m.G0 = v6 + int32(16)
				return v10
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				m.T0[v17].(func(*base.Module, int32))(m, v20)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v10
				}
			}
		}
	}
}
func F_hashtableFairRandomEntry(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
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
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
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
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v17 = v15 + v16
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v20 == int32(255) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = v3
	goto L3
L2:
	;
	v24 = int32(1) << (uint(v20) % 32)
	goto L3
L3:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v27 == int32(255) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v31 = int32(0)
	goto L6
L5:
	;
	v31 = int32(1) << (uint(v27) % 32)
	goto L6
L6:
	;
	if base.Ui32(v17) < base.Ui32(v24+v31) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v34 = int32(12)
	goto L9
L8:
	;
	v34 = int32(120)
	goto L9
L9:
	;
	v37 = v10 - v34<<(uint(int32(2))%32)
	m.G0 = v37
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
	if base.Ui32(v34) < base.Ui32(v17) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = v34
	goto L12
L11:
	;
	v43 = v17
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v43
	if v17 == int32(0) {
		v117 = v3
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v121 == int32(-1) {
		v130 = v117
		goto L30
	} else {
		goto L31
	}
L14:
	;
	goto L15
L15:
	;
	v54 = int32(0)
	F___lock(m, int32(9116960))
	mBase = m.M
	v61 = *(*int32)(unsafe.Add(mBase, _consts[303]))
	v63 = *(*int32)(unsafe.Add(mBase, _consts[304]))
	if v63 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v117 = v112
	goto L13
L17:
	;
	v106 = int32(0)
	v108 = F_hashtableScanDefrag(m, l0, v98, int32(539), v10+int32(4), v106, v106)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	F___unlock(m, int32(9116960))
	mBase = m.M
	goto L17
L19:
	;
	v67 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, _consts[305]))
	v69 = int32(2)
	v71 = v61 + v68<<(uint(v69)%32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v74 = *(*int32)(unsafe.Add(mBase, _consts[306]))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v61+v74<<(uint(v69)%32))))
	v79 = v72 + v78
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v79
	v84 = v74 + int32(1)
	if v84 == v63 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v65 = F_lcg31(m, v64)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v65
	v98 = v65
	goto L18
L21:
	;
	v86 = v67
	goto L23
L22:
	;
	v86 = v84
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[306])) = v86
	v88 = int32(0)
	v91 = v68 + int32(1)
	if v91 == v63 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v93 = v88
	goto L26
L25:
	;
	v93 = v91
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[305])) = v93
	v98 = int32(base.Ui32(v79) >> (uint(int32(1)) % 32))
	goto L18
L27:
	;
	return int32(0)
L28:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if base.Ui32(v112) < base.Ui32(v43) {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	goto L16
L30:
	;
	if base.Ui32(v130) < base.Ui32(v43) {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
	if v124 != 0 {
		v130 = v117
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _consts[340]))
	if v126 != 0 {
		v130 = v117
		goto L30
	} else {
		goto L33
	}
L33:
	;
	F_rehashStep(m, l0)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v130 = v129
	goto L30
L35:
	;
	m.G0 = v10 + int32(16)
	return base.B2i32(v132 != int32(0))
L36:
	;
	v132 = v130
	goto L38
L37:
	;
	v132 = v43
	goto L38
L38:
	;
	if v132 == int32(0) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v135 = int32(0)
	F___lock(m, int32(9116960))
	mBase = m.M
	v142 = *(*int32)(unsafe.Add(mBase, _consts[303]))
	v144 = *(*int32)(unsafe.Add(mBase, _consts[304]))
	if v144 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v184 = base.I32_rem_u_s(v179, v132)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v37+v184<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v188
	goto L35
L41:
	;
	F___unlock(m, int32(9116960))
	mBase = m.M
	goto L40
L42:
	;
	v148 = int32(0)
	v149 = *(*int32)(unsafe.Add(mBase, _consts[305]))
	v150 = int32(2)
	v152 = v142 + v149<<(uint(v150)%32)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v155 = *(*int32)(unsafe.Add(mBase, _consts[306]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v142+v155<<(uint(v150)%32))))
	v160 = v153 + v159
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v160
	v165 = v155 + int32(1)
	if v165 == v144 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v146 = F_lcg31(m, v145)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v146
	v179 = v146
	goto L41
L44:
	;
	v167 = v148
	goto L46
L45:
	;
	v167 = v165
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, _consts[306])) = v167
	v169 = int32(0)
	v172 = v149 + int32(1)
	if v172 == v144 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v174 = v169
	goto L49
L48:
	;
	v174 = v172
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, _consts[305])) = v174
	v179 = int32(base.Ui32(v160) >> (uint(int32(1)) % 32))
	goto L41
}
func F_hashtableFreeStats(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_valkey_free(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_valkey_free(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_hashtableGetStatsMsg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+72)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v13
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = int32(_a660)
	goto L3
L2:
	;
	v19 = int32(_a661)
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v19
	v24 = F_snprintf(m, l0, l1, int32(_a662), v11+int32(64))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v28 != 0 {
		v39 = v24
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if l3 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v29
	v36 = F_snprintf(m, l0+v24, l1-v24, int32(_a663), v11+int32(48))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v39 = v36 + v24
	goto L6
L9:
	;
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+l1+int32(-1)))) = uint8(v113)
	if l0&int32(3) == v113 {
		v136 = l0
		goto L21
	} else {
		goto L22
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v11+int32(32)))) = base.F64_promote_f32(base.F32_div(base.F32_convert_i32_u(v45), base.F32_convert_i32_u(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v47
	v60 = F_snprintf(m, l0+v39, l1-v39, int32(_a664), v11+int32(16))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v67 = v60 + v39
	v69 = int32(0)
	goto L12
L12:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v69<<(uint(int32(2))%32))))
	if v76 == int32(0) {
		v96 = v67
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	v99 = v69 + int32(1)
	if v99 != int32(49) {
		v67 = v96
		v69 = v99
		goto L12
	} else {
		goto L18
	}
L15:
	;
	if base.Ui32(l1) <= base.Ui32(v67) {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v76
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = base.F64_promote_f32(base.F32_mul(base.F32_div(base.F32_convert_i32_u(v76), base.F32_convert_i32_u(v80)), float32(100)))
	v93 = F_snprintf(m, l0+v67, l1-v67, int32(_a665), v11)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v96 = v93 + v67
	goto L14
L18:
	;
	goto L13
L19:
	;
	m.G0 = v11 + int32(80)
	return v169
L20:
	;
	v169 = v161 - l0
	goto L19
L21:
	;
	v140 = v136
	goto L29
L22:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v122 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v125 = l0
	goto L25
L24:
	;
	v169 = l0 - l0
	goto L19
L25:
	;
	v129 = v125 + int32(1)
	if v129&int32(3) == int32(0) {
		v136 = v129
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	if v134 != 0 {
		v125 = v129
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v161 = v129
	goto L20
L29:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v149 = int32(-2139062144)
	if (int32(16843008)-v146|v146)&v149 == v149 {
		v140 = v140 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v155 = v140
	goto L32
L31:
	;
	goto L30
L32:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if v159 != 0 {
		v155 = v155 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v161 = v155
	goto L20
L34:
	;
	goto L33
}
func F_hashtableIsRehashing(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	return base.B2i32(v2 != int32(-1))
}
func F_hashtableMemUsage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	if v11 != 0 {
		v13 = m.T0[v11].(func(*base.Module) int32)(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v19 = v13 + int32(44)
			if v9 == int32(255) {
				v26 = int32(0)
			} else {
				v26 = int32(1) << (uint(v9) % 32)
			}
			if v8 == int32(255) {
				v33 = int32(0)
			} else {
				v33 = int32(1) << (uint(v8) % 32)
			}
			return v19 + (v6+v7+v26+v33)<<(uint(int32(6))%32)
		}
	} else {
		v19 = int32(44)
		if v9 == int32(255) {
			v26 = int32(0)
		} else {
			v26 = int32(1) << (uint(v9) % 32)
		}
		if v8 == int32(255) {
			v33 = int32(0)
		} else {
			v33 = int32(1) << (uint(v8) % 32)
		}
		return v19 + (v6+v7+v26+v33)<<(uint(int32(6))%32)
	}
}
func F_hashtableObjectPrefetchValue(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v2)>>(uint(int32(4))%32))&int32(15) + int32(-1) {
	case 0, 7:
	default:
		v9 = F_objectGetVal(m, l0)
		mBase = m.M
	}
	return
}
func F_hashtablePauseAutoShrink(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	v4 = v2 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v4)
	return
}
func F_hashtableReplaceReallocatedEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v150 int64
	_ = v150
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v188 int64
	_ = v188
	var v194 int64
	_ = v194
	var v198 int64
	_ = v198
	var v202 int64
	_ = v202
	var v207 int64
	_ = v207
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
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v21 == int32(0) {
		v29 = l2
		v30 = v20
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v32 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v24 = m.T0[v21].(func(*base.Module, int32) int32)(m, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = v24
	v30 = v28
	goto L1
L5:
	;
	v210 = base.I32_wrap_i64(int64(base.Ui64(v207) >> (uint(int64(56)) % 64)))
	v211 = base.I32_wrap_i64(v207)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v212 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L6:
	;
	v38 = v18 + int32(12)
	v39 = int32(4)
	v40 = int32(_a657)
	v48 = *(*int64)(unsafe.Add(mBase, _consts[341]))
	v50 = v48 ^ int64(8317987319222330741)
	v51 = *(*int64)(unsafe.Add(mBase, _consts[342]))
	v53 = v51 ^ int64(7237128888997146477)
	v55 = v48 ^ int64(7816392313619706465)
	v57 = v51 ^ int64(8387220255154660723)
	v62 = v18 + int32(16) - v39
	if v38 == v62 {
		v100 = v38
		v103 = v55
		v104 = v50
		v105 = v57
		v106 = v53
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v35 = m.T0[v32].(func(*base.Module, int32) int64)(m, v29)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v207 = v35
	goto L5
L9:
	;
	v207 = base.I64_rotl(base.I64_rotl(v179, v146)^v188, v150) ^ base.I64_rotl(v198, v172) ^ base.I64_rotl(v202, v153) ^ v202
	goto L5
L10:
	;
	v111 = base.I64_extend_i32_u(v39) << (uint(int64(56)) % 64)
	switch v39 {
	default:
		v144 = v111
		goto L15
	case 1:
		v141 = v111
		goto L16
	case 2:
		v136 = v111
		goto L17
	case 3:
		v131 = v111
		goto L18
	case 4:
		v126 = v111
		goto L19
	case 5:
		v121 = v111
		goto L20
	case 6:
		v116 = v111
		goto L21
	case 7:
		goto L22
	}
L11:
	;
	v64 = v38
	v67 = v55
	v68 = v50
	v69 = v57
	v70 = v53
	goto L12
L12:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v75 = v74 ^ v69
	v76 = v75 + v67
	v77 = v68 + v70
	v80 = v77 ^ base.I64_rotl(v70, int64(13))
	v81 = v76 + v80
	v84 = v81 ^ base.I64_rotl(v80, int64(17))
	v87 = base.I64_rotl(v75, int64(16)) ^ v76
	v90 = int64(32)
	v92 = v87 + base.I64_rotl(v77, v90)
	v93 = base.I64_rotl(v87, int64(21)) ^ v92
	v95 = base.I64_rotl(v81, v90)
	v96 = v92 ^ v74
	v98 = v64 + int32(8)
	if v98 != v62 {
		v64 = v98
		v67 = v95
		v68 = v96
		v69 = v93
		v70 = v84
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v100 = v62
	v103 = v95
	v104 = v96
	v105 = v93
	v106 = v84
	goto L10
L14:
	;
	goto L13
L15:
	;
	v145 = v144 ^ v105
	v146 = int64(16)
	v148 = v145 + v103
	v149 = base.I64_rotl(v145, v146) ^ v148
	v150 = int64(21)
	v152 = v104 + v106
	v153 = int64(32)
	v155 = v149 + base.I64_rotl(v152, v153)
	v156 = base.I64_rotl(v149, v150) ^ v155
	v159 = int64(13)
	v161 = v152 ^ base.I64_rotl(v106, v159)
	v162 = v148 + v161
	v167 = base.I64_rotl(v162, v153) ^ int64(255) + v156
	v168 = base.I64_rotl(v156, v146) ^ v167
	v172 = int64(17)
	v174 = v162 ^ base.I64_rotl(v161, v172)
	v175 = v155 ^ v144 + v174
	v178 = base.I64_rotl(v175, v153) + v168
	v179 = base.I64_rotl(v168, v150) ^ v178
	v184 = v175 ^ base.I64_rotl(v174, v159)
	v185 = v184 + v167
	v188 = base.I64_rotl(v185, v153) + v179
	v194 = base.I64_rotl(v184, v172) ^ v185
	v198 = base.I64_rotl(v194, v159) ^ (v194 + v178)
	v202 = v198 + v188
	goto L9
L16:
	;
	v142 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	v144 = v141 | v142
	goto L15
L17:
	;
	v137 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	v141 = v137<<(uint(int64(8))%64) | v136
	goto L16
L18:
	;
	v132 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
	v136 = v132<<(uint(int64(16))%64) | v131
	goto L17
L19:
	;
	v127 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
	v131 = v127<<(uint(int64(24))%64) | v126
	goto L18
L20:
	;
	v122 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)))
	v126 = v122<<(uint(int64(32))%64) | v121
	goto L19
L21:
	;
	v117 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v100)+5)))
	v121 = v117<<(uint(int64(40))%64) | v116
	goto L20
L22:
	;
	v112 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v100)+6)))
	v116 = v112<<(uint(int64(48))%64) | v111
	goto L21
L23:
	;
	m.G0 = v18 + int32(16)
	return v414
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v407))) = l2
	v414 = int32(1)
	goto L23
L25:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v312 != 0 {
		goto L44
	} else {
		goto L45
	}
L26:
	;
	v216 = int32(-1)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v217 == int32(255) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v223 = int32(0)
	goto L29
L28:
	;
	v223 = v216<<(uint(v217)%32) ^ v216
	goto L29
L29:
	;
	v224 = v223 & v211
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v225 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v242 = v229 + v224<<(uint(int32(6))%32)
	goto L33
L31:
	;
	if base.Ui32(v224) < base.Ui32(v225) {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v242))))
	v253 = int32(1)
	v259 = v252 & v253
	v266 = int32(0)
	goto L35
L34:
	;
	goto L25
L35:
	;
	if int32(base.Ui32(int32(base.Ui32(v252)>>(uint(v253)%32))&int32(4095))>>(uint(v266)%32))&int32(1) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v259 == int32(0) {
		goto L25
	} else {
		goto L42
	}
L37:
	;
	v292 = v266 + int32(1)
	if v292 != int32(12)-v259 {
		v266 = v292
		goto L35
	} else {
		goto L41
	}
L38:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242+int32(2)+v266))))
	if v283 != v210 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v287 = v242 + int32(16) + v266<<(uint(int32(2))%32)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	if v288 == l1 {
		v407 = v287
		goto L24
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	goto L36
L42:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v242)+60))
	if v296 != 0 {
		v242 = v296
		goto L33
	} else {
		goto L43
	}
L43:
	;
	goto L34
L44:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v316 = int32(-1)
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v317 == int32(255) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v414 = int32(0)
	goto L23
L46:
	;
	v323 = int32(0)
	goto L48
L47:
	;
	v323 = v316<<(uint(v317)%32) ^ v316
	goto L48
L48:
	;
	v328 = v314 + v323&v211<<(uint(int32(6))%32)
	goto L49
L49:
	;
	v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v328))))
	v348 = int32(1)
	v354 = v347 & v348
	v361 = int32(0)
	goto L51
L51:
	;
	if int32(base.Ui32(int32(base.Ui32(v347)>>(uint(v348)%32))&int32(4095))>>(uint(v361)%32))&int32(1) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v389 = int32(0)
	if v354 == v389 {
		v414 = v389
		goto L23
	} else {
		goto L58
	}
L53:
	;
	v387 = v361 + int32(1)
	if v387 != int32(12)-v354 {
		v361 = v387
		goto L51
	} else {
		goto L57
	}
L54:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328+int32(2)+v361))))
	if v378 != v210 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v382 = v328 + int32(16) + v361<<(uint(int32(2))%32)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	if v383 == l1 {
		v407 = v382
		goto L24
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	goto L52
L58:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v328)+60))
	if v392 != 0 {
		v328 = v392
		goto L49
	} else {
		goto L59
	}
L59:
	;
	v414 = v389
	goto L23
}
func F_hashtableResumeAutoShrink(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	v6 = v4 + int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v6)
	if v6&int32(65535) != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v10 != int32(-1) {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v18 == int32(255) {
				v22 = int32(0)
			} else {
				v22 = int32(12) << (uint(v18) % 32)
			}
			v26 = *(*int32)(unsafe.Add(mBase, _consts[340]))
			if v26 != 0 {
				v27 = int32(3)
			} else {
				v27 = int32(13)
			}
			if base.Ui32(v22*v27) < base.Ui32(v13*int32(100)) {
				return
			} else {
				v31 = F_resize_1(m, l0, v13, int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_hashtableSdsKeyCompare(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	v3 = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v10 & int32(7) {
	case 0:
		v27 = int32(base.Ui32(v10) >> (uint(int32(3)) % 32))
	case 1:
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v27 = v17
	case 2:
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v27 = v20
	case 3:
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v27 = v23
	case 4:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v27 = v26
	default:
		v27 = v3
	}
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v30 & int32(7) {
	case 0:
		v47 = int32(base.Ui32(v30) >> (uint(int32(3)) % 32))
	case 1:
		v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
		v47 = v37
	case 2:
		v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
		v47 = v40
	case 3:
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
		v47 = v43
	case 4:
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
		v47 = v46
	default:
		v47 = v3
	}
	if v27 != v47 {
		v103 = int32(1)
	} else {
		v50 = int32(0)
		v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
		switch v57 & int32(7) {
		case 0:
			v74 = int32(base.Ui32(v57) >> (uint(int32(3)) % 32))
		case 1:
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
			v74 = v64
		case 2:
			v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
			v74 = v67
		case 3:
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v74 = v70
		case 4:
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
			v74 = v73
		default:
			v74 = v50
		}
		v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
		switch v77 & int32(7) {
		case 0:
			v94 = int32(base.Ui32(v77) >> (uint(int32(3)) % 32))
		case 1:
			v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
			v94 = v84
		case 2:
			v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
			v94 = v87
		case 3:
			v90 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
			v94 = v90
		case 4:
			v93 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
			v94 = v93
		default:
			v94 = v50
		}
		v95 = base.B2i32(base.Ui32(v74) < base.Ui32(v94))
		if base.Ui32(v74) < base.Ui32(v94) {
			v96 = v74
		} else {
			v96 = v94
		}
		v97 = F_memcmp(m, l0, l1, v96)
		mBase = m.M
		if v97 != 0 {
			v100 = v97
		} else {
			v100 = base.B2i32(base.Ui32(v94) < base.Ui32(v74)) - v95
		}
		v103 = base.B2i32(v100 != int32(0))
	}
	return v103
}
func F_hashtableSetResizePolicy(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[340])) = l0
	return
}
func F_hashtableTwoPhasePopFindRef(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v73 int64
	_ = v73
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v177 int64
	_ = v177
	var v183 int64
	_ = v183
	var v187 int64
	_ = v187
	var v191 int64
	_ = v191
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v12 == v4-v14 {
		v232 = v4
		m.G0 = v9 + int32(16)
		return v232
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
		if v19 == int32(0) {
			v27 = v9 + int32(12)
			v28 = int32(4)
			v29 = int32(_a657)
			v37 = *(*int64)(unsafe.Add(mBase, _consts[341]))
			v39 = v37 ^ int64(8317987319222330741)
			v40 = *(*int64)(unsafe.Add(mBase, _consts[342]))
			v42 = v40 ^ int64(7237128888997146477)
			v44 = v37 ^ int64(7816392313619706465)
			v46 = v40 ^ int64(8387220255154660723)
			v51 = v9 + int32(16) - v28
			if v27 == v51 {
				v89 = v27
				v92 = v44
				v93 = v39
				v94 = v46
				v95 = v42
			} else {
				v53 = v27
				v56 = v44
				v57 = v39
				v58 = v46
				v59 = v42
				for {
					v63 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
					v64 = v63 ^ v58
					v65 = v64 + v56
					v66 = v57 + v59
					v69 = v66 ^ base.I64_rotl(v59, int64(13))
					v70 = v65 + v69
					v73 = v70 ^ base.I64_rotl(v69, int64(17))
					v76 = base.I64_rotl(v64, int64(16)) ^ v65
					v79 = int64(32)
					v81 = v76 + base.I64_rotl(v66, v79)
					v82 = base.I64_rotl(v76, int64(21)) ^ v81
					v84 = base.I64_rotl(v70, v79)
					v85 = v81 ^ v63
					v87 = v53 + int32(8)
					if v87 != v51 {
						v53 = v87
						v56 = v84
						v57 = v85
						v58 = v82
						v59 = v73
						continue
					} else {
						break
					}
					break
				}
				v89 = v51
				v92 = v84
				v93 = v85
				v94 = v82
				v95 = v73
			}
			v100 = base.I64_extend_i32_u(v28) << (uint(int64(56)) % 64)
			switch v28 {
			default:
				v133 = v100
			case 1:
				v130 = v100
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			case 2:
				v125 = v100
				v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
				v130 = v126<<(uint(int64(8))%64) | v125
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			case 3:
				v120 = v100
				v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+2)))
				v125 = v121<<(uint(int64(16))%64) | v120
				v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
				v130 = v126<<(uint(int64(8))%64) | v125
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			case 4:
				v115 = v100
				v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+3)))
				v120 = v116<<(uint(int64(24))%64) | v115
				v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+2)))
				v125 = v121<<(uint(int64(16))%64) | v120
				v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
				v130 = v126<<(uint(int64(8))%64) | v125
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			case 5:
				v110 = v100
				v111 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
				v115 = v111<<(uint(int64(32))%64) | v110
				v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+3)))
				v120 = v116<<(uint(int64(24))%64) | v115
				v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+2)))
				v125 = v121<<(uint(int64(16))%64) | v120
				v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
				v130 = v126<<(uint(int64(8))%64) | v125
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			case 6:
				v105 = v100
				v106 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+5)))
				v110 = v106<<(uint(int64(40))%64) | v105
				v111 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
				v115 = v111<<(uint(int64(32))%64) | v110
				v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+3)))
				v120 = v116<<(uint(int64(24))%64) | v115
				v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+2)))
				v125 = v121<<(uint(int64(16))%64) | v120
				v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
				v130 = v126<<(uint(int64(8))%64) | v125
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			case 7:
				v101 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+6)))
				v105 = v101<<(uint(int64(48))%64) | v100
				v106 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+5)))
				v110 = v106<<(uint(int64(40))%64) | v105
				v111 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
				v115 = v111<<(uint(int64(32))%64) | v110
				v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+3)))
				v120 = v116<<(uint(int64(24))%64) | v115
				v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+2)))
				v125 = v121<<(uint(int64(16))%64) | v120
				v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
				v130 = v126<<(uint(int64(8))%64) | v125
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			}
			v134 = v133 ^ v94
			v135 = int64(16)
			v137 = v134 + v92
			v138 = base.I64_rotl(v134, v135) ^ v137
			v139 = int64(21)
			v141 = v93 + v95
			v142 = int64(32)
			v144 = v138 + base.I64_rotl(v141, v142)
			v145 = base.I64_rotl(v138, v139) ^ v144
			v148 = int64(13)
			v150 = v141 ^ base.I64_rotl(v95, v148)
			v151 = v137 + v150
			v156 = base.I64_rotl(v151, v142) ^ int64(255) + v145
			v157 = base.I64_rotl(v145, v135) ^ v156
			v161 = int64(17)
			v163 = v151 ^ base.I64_rotl(v150, v161)
			v164 = v144 ^ v133 + v163
			v167 = base.I64_rotl(v164, v142) + v157
			v168 = base.I64_rotl(v157, v139) ^ v167
			v173 = v164 ^ base.I64_rotl(v163, v148)
			v174 = v173 + v156
			v177 = base.I64_rotl(v174, v142) + v168
			v183 = base.I64_rotl(v173, v161) ^ v174
			v187 = base.I64_rotl(v183, v148) ^ (v183 + v167)
			v191 = v187 + v177
			v196 = base.I64_rotl(base.I64_rotl(v168, v135)^v177, v139) ^ base.I64_rotl(v187, v161) ^ base.I64_rotl(v191, v142) ^ v191
			v197 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v197
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v197
			v206 = F_findBucket_1(m, l0, v196, l1, v9+int32(8), v9+int32(4))
			mBase = m.M
			v207 = m.ExcPending
			if v207 != 0 {
				return int32(0)
			} else {
				if v206 == int32(0) {
					v232 = v197
					m.G0 = v9 + int32(16)
					return v232
				} else {
					v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
					v211 = int32(1)
					v212 = v210 + v211
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)) = uint16(v212)
					v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
					v216 = v214 + v211
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v216)
					if l2 == int32(0) {
						F__serverAssert(m, int32(_a658), int32(_a659), int32(1839))
						mBase = m.M
						v242 = m.ExcPending
						if v242 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v206
						v221 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v221)
						v223 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v223)
						v232 = v206 + v221<<(uint(int32(2))%32) + int32(16)
						m.G0 = v9 + int32(16)
						return v232
					}
				}
			}
		} else {
			v22 = m.T0[v19].(func(*base.Module, int32) int64)(m, l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v196 = v22
				v197 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v197
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v197
				v206 = F_findBucket_1(m, l0, v196, l1, v9+int32(8), v9+int32(4))
				mBase = m.M
				v207 = m.ExcPending
				if v207 != 0 {
					return int32(0)
				} else {
					if v206 == int32(0) {
						v232 = v197
						m.G0 = v9 + int32(16)
						return v232
					} else {
						v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
						v211 = int32(1)
						v212 = v210 + v211
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)) = uint16(v212)
						v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
						v216 = v214 + v211
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v216)
						if l2 == int32(0) {
							F__serverAssert(m, int32(_a658), int32(_a659), int32(1839))
							mBase = m.M
							v242 = m.ExcPending
							if v242 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v206
							v221 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
							*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v221)
							v223 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v223)
							v232 = v206 + v221<<(uint(int32(2))%32) + int32(16)
							m.G0 = v9 + int32(16)
							return v232
						}
					}
				}
			}
		}
	}
}
