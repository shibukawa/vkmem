package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_VM_ACLCheckCommandPermissions(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_lookupCommand(m, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = F_ACLCheckAllUserCommandPerm(m, v15, v10, l1, l2, int32(-1), v8+int32(12))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					v29 = int32(0)
				} else {
					v24 = int32(2)
					*(*int32)(unsafe.Add(mBase, _consts[5])) = v24
					v29 = int32(1)
				}
				m.G0 = v8 + int32(16)
				return v29
			}
		} else {
			v24 = int32(44)
			*(*int32)(unsafe.Add(mBase, _consts[5])) = v24
			v29 = int32(1)
			m.G0 = v8 + int32(16)
			return v29
		}
	}
}
func F_VM_ACLCheckKeyPermissions(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v26 int64
	_ = v26
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v56 int64
	_ = v56
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v84 int64
	_ = v84
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v97 int64
	_ = v97
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v112 int64
	_ = v112
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v125 int64
	_ = v125
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v140 int64
	_ = v140
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v153 int64
	_ = v153
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v168 int64
	_ = v168
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	if l2&int32(-241) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v258
L2:
	;
	goto L44
L3:
	;
	v14 = int32(0)
	v15 = base.I64_extend_i32_u(l2)
	v26 = *(*int64)(unsafe.Add(mBase, _consts[567]))
	if base.B2i32(v26&v15 == int64(0)) == v14 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v246 = int32(28)
	goto L2
L5:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v180 = F_objectGetVal(m, l1)
	mBase = m.M
	v183 = F_objectGetVal(m, l1)
	mBase = m.M
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+int32(-1)))))
	switch v186 & int32(7) {
	case 0:
		goto L34
	case 1:
		goto L33
	case 2:
		goto L32
	case 3:
		goto L31
	case 4:
		goto L30
	default:
		v203 = int32(0)
		goto L29
	}
L6:
	;
	v41 = *(*int64)(unsafe.Add(mBase, _consts[568]))
	if v41&v15 == int64(0) {
		v51 = v38
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v37 = *(*int64)(unsafe.Add(mBase, _consts[530]))
	v38 = v37
	goto L6
L8:
	;
	v38 = int64(0)
	goto L6
L9:
	;
	v56 = *(*int64)(unsafe.Add(mBase, _consts[569]))
	if v56&v15 == int64(0) {
		v66 = v51
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v49 = *(*int64)(unsafe.Add(mBase, _consts[531]))
	v51 = v49 | v38
	goto L9
L11:
	;
	v69 = *(*int64)(unsafe.Add(mBase, _consts[570]))
	if v69&v15 == int64(0) {
		v79 = v66
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v64 = *(*int64)(unsafe.Add(mBase, _consts[532]))
	v66 = v64 | v51
	goto L11
L13:
	;
	v84 = *(*int64)(unsafe.Add(mBase, _consts[571]))
	if v84&v15 == int64(0) {
		v94 = v79
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v77 = *(*int64)(unsafe.Add(mBase, _consts[533]))
	v79 = v77 | v66
	goto L13
L15:
	;
	v97 = *(*int64)(unsafe.Add(mBase, _consts[572]))
	if v97&v15 == int64(0) {
		v107 = v94
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v92 = *(*int64)(unsafe.Add(mBase, _consts[534]))
	v94 = v92 | v79
	goto L15
L17:
	;
	v112 = *(*int64)(unsafe.Add(mBase, _consts[573]))
	if v112&v15 == int64(0) {
		v122 = v107
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v105 = *(*int64)(unsafe.Add(mBase, _consts[535]))
	v107 = v105 | v94
	goto L17
L19:
	;
	v125 = *(*int64)(unsafe.Add(mBase, _consts[574]))
	if v125&v15 == int64(0) {
		v135 = v122
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v120 = *(*int64)(unsafe.Add(mBase, _consts[536]))
	v122 = v120 | v107
	goto L19
L21:
	;
	v140 = *(*int64)(unsafe.Add(mBase, _consts[575]))
	if v140&v15 == int64(0) {
		v150 = v135
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v133 = *(*int64)(unsafe.Add(mBase, _consts[537]))
	v135 = v133 | v122
	goto L21
L23:
	;
	v153 = *(*int64)(unsafe.Add(mBase, _consts[576]))
	if v153&v15 == int64(0) {
		v163 = v150
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v148 = *(*int64)(unsafe.Add(mBase, _consts[538]))
	v150 = v148 | v135
	goto L23
L25:
	;
	v168 = *(*int64)(unsafe.Add(mBase, _consts[577]))
	if v168&v15 == int64(0) {
		v178 = v163
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v161 = *(*int64)(unsafe.Add(mBase, _consts[539]))
	v163 = v161 | v150
	goto L25
L27:
	;
	goto L5
L28:
	;
	v176 = *(*int64)(unsafe.Add(mBase, _consts[540]))
	v178 = v176 | v163
	goto L27
L29:
	;
	v204 = int32(0)
	v207 = m.G0
	v209 = v207 - int32(16)
	m.G0 = v209
	if v179 == v204 {
		v237 = v204
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v183+int32(-17))))
	v203 = v202
	goto L29
L31:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v183+int32(-9))))
	v203 = v199
	goto L29
L32:
	;
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183+int32(-5)))))
	v203 = v196
	goto L29
L33:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+int32(-3)))))
	v203 = v193
	goto L29
L34:
	;
	v203 = int32(base.Ui32(v186) >> (uint(int32(3)) % 32))
	goto L29
L35:
	;
	if v237 == int32(0) {
		v258 = v14
		goto L1
	} else {
		goto L43
	}
L36:
	;
	m.G0 = v209 + int32(16)
	goto L35
L37:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	F_listRewind(m, v214, v209+int32(8))
	mBase = m.M
	goto L38
L38:
	;
	v227 = F_listNext(m, v209+int32(8))
	mBase = m.M
	if v227 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v237 = v204
	goto L36
L40:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	v230 = F_ACLSelectorCheckKey(m, v229, v180, v203, base.I32_wrap_i64(v178), v204)
	mBase = m.M
	if v230 != 0 {
		goto L38
	} else {
		goto L42
	}
L41:
	;
	v237 = int32(3)
	goto L36
L42:
	;
	goto L39
L43:
	;
	v246 = int32(2)
	goto L2
L44:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v246
	v258 = int32(1)
	goto L1
}
func F_VM_AbortBlock(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
	v7 = F_VM_UnblockClient(m, l0, v2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_VM_AddPostNotificationJob(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v7 = int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	if v9 != 0 {
		v35 = v7
		return v35
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[130]))
		if v11 == int32(0) {
			v17 = F_valkey_malloc(m, int32(20))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v21
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+96))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v28
				v30 = int32(0)
				v32 = *(*int32)(unsafe.Add(mBase, _consts[564]))
				v33 = F_listAddNodeTail(m, v32, v17)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = v30
					return v35
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _consts[330]))
			if v15 != 0 {
				v35 = v7
				return v35
			} else {
				v17 = F_valkey_malloc(m, int32(20))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = v21
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+96))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v28
					v30 = int32(0)
					v32 = *(*int32)(unsafe.Add(mBase, _consts[564]))
					v33 = F_listAddNodeTail(m, v32, v17)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = v30
						return v35
					}
				}
			}
		}
	}
}
func F_VM_AutoMemory(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v2 | int32(1)
	return
}
func F_VM_AvoidReplicaTraffic(m *base.Module) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_isPausedActionsWithUpdate(m, int32(16))
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v2 != int32(0))
	}
}
func F_VM_BlockedClientDisconnected(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return int32(base.Ui32(v2)>>(uint(int32(5))%32)) & int32(1)
}
func F_VM_Call(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
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
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v5
	v22 = F_moduleCreateArgvFromUserFormat(m, l1, l2, v9+int32(12), v9+int32(8), l3)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[546]))
	if v27 == v26 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	F_moduleCallCommandHelper(m, l0, v59, v22, v60, v61, v9)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L8
	}
L4:
	;
	v46 = F_createClient(m, int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v30 = int32(0)
	v32 = v27 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[546])) = v32
	v35 = *(*int32)(unsafe.Add(mBase, _consts[547]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v32<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, _consts[548]))
	if base.Ui32(v41) <= base.Ui32(v32) {
		v59 = v39
		goto L3
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[548])) = v32
	v59 = v39
	goto L3
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+328)) = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+200)) = v50 | int32(1073741824)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+204)) = v54 | int32(268435456)
	v59 = v46
	goto L3
L8:
	;
	goto L14
L9:
	;
	F__serverAssert(m, int32(_a929), int32(_a917), int32(6579))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L49
	}
L10:
	;
	F__serverAssert(m, int32(_a930), int32(_a917), int32(6578))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L48
	}
L11:
	;
	if v186 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L12:
	;
	if v144 != 0 {
		goto L35
	} else {
		goto L36
	}
L13:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v138 != 0 {
		goto L32
	} else {
		goto L33
	}
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v65 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+200)))
	if v66&int32(16) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+9)))
	if v122&int32(8) == int32(0) {
		goto L10
	} else {
		goto L29
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v59)+124))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v59)+180))
	v72 = F_sdsnewlen(m, v70, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v74 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+180)) = v74
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v59)+132))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	if v82 == v74 {
		v106 = v72
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v59)+344))
	v111 = F_callReplyCreate(m, v106, v110, (v74-v69&int32(1))&l0)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L26
	}
L20:
	;
	v86 = v81
	v87 = v72
	goto L21
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v96 = F_sdscatlen(m, v87, v92+int32(13), v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v106 = v96
	goto L19
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v59)+132))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	F_listDelNode(m, v98, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v59)+132))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+20))
	if v103 != 0 {
		v86 = v102
		v87 = v96
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v113 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+344)) = v113
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+9)))
	if v115&int32(16) == v113 {
		v144 = v111
		v145 = v59
		goto L12
	} else {
		goto L27
	}
L27:
	;
	F_enableParseExactReplyTypeFlag(m, v111)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v144 = v111
	v145 = v59
	goto L12
L29:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v59)+116))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+52))
	if v128 == int32(0) {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v131 + int32(1)
	v136 = F_callReplyCreatePromise(m, v128)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v144 = v136
	v145 = int32(0)
	goto L12
L32:
	;
	v140 = F_callReplyCreateError(m, v138, l0)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	v185 = int32(0)
	v186 = v59
	goto L11
L34:
	;
	v144 = v140
	v145 = v59
	goto L12
L35:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v149&int32(1) == int32(0) {
		v185 = v144
		v186 = v145
		goto L11
	} else {
		goto L37
	}
L36:
	;
	v185 = int32(0)
	v186 = v145
	goto L11
L37:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v154 == v155 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v176 = v173 + v172<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v176)+4)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v172 + int32(1)
	v185 = v144
	v186 = v145
	goto L11
L39:
	;
	v158 = int32(8)
	if v158 < v154 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v172 = v154
	v173 = v157
	goto L38
L41:
	;
	v161 = v154
	goto L43
L42:
	;
	v161 = v158
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v161 << (uint(int32(1)) % 32)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v168 = F_valkey_realloc(m, v165, v161<<(uint(int32(4))%32))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v168
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v172 = v171
	v173 = v168
	goto L38
L45:
	;
	m.G0 = v9 + int32(16)
	return v185
L46:
	;
	F_moduleReleaseTempClient(m, v186)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_VM_CallArgv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l4 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F__serverAssert(m, int32(_a929), int32(_a917), int32(7094))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L14
	} else {
		goto L85
	}
L2:
	;
	F__serverAssert(m, int32(_a930), int32(_a917), int32(7093))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L14
	} else {
		goto L84
	}
L3:
	;
	F__serverAssert(m, int32(_a931), int32(_a917), int32(7092))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L14
	} else {
		goto L83
	}
L4:
	;
	F__serverAssert(m, int32(_a932), int32(_a917), int32(7091))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L14
	} else {
		goto L82
	}
L5:
	;
	F__serverAssert(m, int32(_a933), int32(_a917), int32(7072))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L14
	} else {
		goto L81
	}
L6:
	;
	F__serverAssert(m, int32(_a934), int32(_a917), int32(7055))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L14
	} else {
		goto L80
	}
L7:
	;
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v23
	v26 = *(*int32)(unsafe.Add(mBase, _consts[546]))
	if v26 == v23 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
	if v20 != int64(1) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+208)))
	v63 = v61 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+208)) = uint8(v63)
	F_moduleCallCommandHelper(m, l0, v60, l1, l2, l3, v16+int32(12))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L14
	} else {
		goto L16
	}
L11:
	;
	v45 = F_createClient(m, int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v29 = int32(0)
	v31 = v26 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[546])) = v31
	v34 = *(*int32)(unsafe.Add(mBase, _consts[547]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v31<<(uint(int32(2))%32))))
	v40 = *(*int32)(unsafe.Add(mBase, _consts[548]))
	if base.Ui32(v40) <= base.Ui32(v31) {
		v60 = v38
		goto L10
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[548])) = v31
	v60 = v38
	goto L10
L14:
	;
	return int32(0)
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+200)) = v51 | int32(1073741824)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v45)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+204)) = v55 | int32(268435456)
	v60 = v45
	goto L10
L16:
	;
	goto L17
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+200)))
	if v71&int32(16) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	m.G0 = v16 + int32(16)
	v309 = int32(0)
	return base.B2i32(l3&int32(256) == v309) & base.B2i32(v70 != v309)
L19:
	;
	if v70 != 0 {
		goto L4
	} else {
		goto L64
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v74 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if l4 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L22:
	;
	if v70 == int32(0) {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	if l4 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_sdsfree(m, v202)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L14
	} else {
		goto L59
	}
L25:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v79 != int32(45) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v202 = v74
	goto L24
L27:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176+int32(-1)))))
	switch v181 & int32(7) {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	case 4:
		goto L53
	default:
		v198 = int32(0)
		goto L52
	}
L28:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+int32(-1)))))
	switch v85 & int32(7) {
	case 0:
		goto L35
	case 1:
		goto L34
	case 2:
		goto L33
	case 3:
		goto L32
	case 4:
		goto L31
	default:
		v102 = int32(0)
		goto L30
	}
L29:
	;
	v176 = v74
	goto L27
L30:
	;
	if v102 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(-17))))
	v102 = v101
	goto L30
L32:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(-9))))
	v102 = v98
	goto L30
L33:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74+int32(-5)))))
	v102 = v95
	goto L30
L34:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+int32(-3)))))
	v102 = v92
	goto L30
L35:
	;
	v102 = int32(base.Ui32(v85) >> (uint(int32(3)) % 32))
	goto L30
L36:
	;
	v168 = F_sdsempty(m)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L14
	} else {
		goto L49
	}
L37:
	;
	goto L36
L38:
	;
	v118 = int32(0)
	goto L39
L39:
	;
	goto L42
L40:
	;
	goto L37
L41:
	;
	v157 = v118 + int32(1)
	if v157 != v102 {
		v118 = v157
		goto L39
	} else {
		goto L48
	}
L42:
	;
	v125 = v74 + v118
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v135 = int32(0)
	goto L43
L43:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+uint32(_consts[549]))))
	if v126&int32(255) != v139 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L41
L45:
	;
	v145 = v135 + int32(1)
	if v145 != int32(2) {
		v135 = v145
		goto L43
	} else {
		goto L47
	}
L46:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+uint32(_consts[383]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v142)
	goto L41
L47:
	;
	goto L44
L48:
	;
	goto L40
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v74
	v172 = F_sdscatfmt(m, v168, int32(_a935), v16)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	F_sdsfree(m, v74)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	v176 = v172
	goto L27
L52:
	;
	F_addReplyProto(m, v60, v176, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L14
	} else {
		goto L58
	}
L53:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v176+int32(-17))))
	v198 = v197
	goto L52
L54:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v176+int32(-9))))
	v198 = v194
	goto L52
L55:
	;
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+int32(-5)))))
	v198 = v191
	goto L52
L56:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176+int32(-3)))))
	v198 = v188
	goto L52
L57:
	;
	v198 = int32(base.Ui32(v181) >> (uint(int32(3)) % 32))
	goto L52
L58:
	;
	v202 = v176
	goto L24
L59:
	;
	goto L21
L60:
	;
	F_moduleReleaseTempClient(m, v60)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L14
	} else {
		goto L63
	}
L61:
	;
	F_invokeReplyHandlers(m, l0, v60, l4, l5)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L14
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	goto L18
L64:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v215 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	if l3&int32(2048) == int32(0) {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v60)+116))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+52))
	if v221 == int32(0) {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v227 = F_valkey_malloc(m, l2<<(uint(int32(2))%32))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L14
	} else {
		goto L68
	}
L68:
	;
	if l2 <= int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = v227
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+208)))
	v272 = v270 & int32(254)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+208)) = uint8(v272)
	if l4 == int32(0) {
		goto L18
	} else {
		goto L75
	}
L70:
	;
	v238 = int32(0)
	goto L71
L71:
	;
	v245 = v238 << (uint(int32(2)) % 32)
	v246 = l1 + v245
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	F_incrRefCount(m, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L14
	} else {
		goto L73
	}
L72:
	;
	goto L69
L73:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	*(*int32)(unsafe.Add(mBase, uint32(v227+v245))) = v251
	v254 = v238 + int32(1)
	if v254 != l2 {
		v238 = v254
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v60)+116))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+24)) = int32(1)
	goto L78
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+128)) = l5
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l4)+88))
	m.T0[v288].(func(*base.Module, int32, int32, int32))(m, l5, l0, v277)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L14
	} else {
		goto L79
	}
L77:
	;
	goto L76
L78:
	;
	v285 = F__emscripten_memcpy_bulkmem(m, v277+int32(32), l4, int32(96))
	mBase = m.M
	goto L77
L79:
	;
	goto L18
L80:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_VM_CallReplyArrayElement(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_callReplyGetArrayElement(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_VM_CallReplyAttributeElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_callReplyGetAttributeElement(m, l0, l1, l2, l3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v5 != int32(0))
	}
}
func F_VM_CallReplyMapElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_callReplyGetMapElement(m, l0, l1, l2, l3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v5 != int32(0))
	}
}
func F_VM_CallReplyPromiseSetUnblockHandler(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
	if v5 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = l1
		return
	} else {
		F__serverAssert(m, int32(_a928), int32(_a917), int32(6240))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
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
func F_VM_CallReplyStringPtr(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l1 != 0 {
		v10 = l1
	} else {
		v10 = v6 + int32(12)
	}
	v11 = F_callReplyGetString(m, l0, v10)
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v11
	}
}
func F_VM_Calloc(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_zcalloc_usable(m, l1*l0, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_VM_ChannelAtPosWithFlags(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v6&int32(1) == int32(0) {
		return
	} else {
		if l1 < int32(1) {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v13 == int32(0) {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				if v16 != v17 {
					v27 = v16
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					v31 = v28 + v27<<(uint(int32(3))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = l2<<(uint(int32(10))%32)&int32(12288) | l2<<(uint(int32(13))%32)&int32(16384) | l2<<(uint(int32(11))%32)&int32(2048)
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v27 + int32(1)
					return
				} else {
					v19 = int32(8192)
					if v16 < v19 {
						v22 = v16
					} else {
						v22 = v19
					}
					v24 = F_getKeysPrepareResult(m, v13, v22+v16)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						v27 = v26
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						v31 = v28 + v27<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = l2<<(uint(int32(10))%32)&int32(12288) | l2<<(uint(int32(13))%32)&int32(16384) | l2<<(uint(int32(11))%32)&int32(2048)
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v27 + int32(1)
						return
					}
				}
			}
		}
	}
}
func F_VM_ClusterCanonicalKeyNameInSlot(m *base.Module, l0 int32) int32 {
	var v9 int32
	_ = v9
	if base.Ui32(l0) < base.Ui32(int32(16384)) {
		v9 = int32(_a225) + l0<<(uint(int32(2))%32)
	} else {
		v9 = int32(0)
	}
	return v9
}
func F_VM_ClusterKeySlot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	v5 = F_objectGetVal(m, l0)
	mBase = m.M
	v7 = F_objectGetVal(m, l0)
	mBase = m.M
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-1)))))
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
		v299 = int32(0)
		goto L1
	}
L1:
	;
	v300 = int32(0)
	if v299 < int32(1) {
		v320 = v300
		goto L95
	} else {
		goto L96
	}
L2:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-17))))
	v299 = v298
	goto L1
L3:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-9))))
	v228 = int32(0)
	if v227 < int32(1) {
		v248 = v228
		goto L74
	} else {
		goto L75
	}
L4:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(-5)))))
	v157 = int32(0)
	if v156 < int32(1) {
		v177 = v157
		goto L53
	} else {
		goto L54
	}
L5:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-3)))))
	v86 = int32(0)
	if v85 < int32(1) {
		v106 = v86
		goto L32
	} else {
		goto L33
	}
L6:
	;
	v14 = int32(base.Ui32(v10) >> (uint(int32(3)) % 32))
	v15 = int32(0)
	if v14 < int32(1) {
		v35 = v15
		goto L11
	} else {
		goto L12
	}
L7:
	;
	return v77 & int32(16383)
L8:
	;
	goto L7
L9:
	;
	v46 = v35 + int32(1)
	if v14 <= v46 {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	v44 = F_crc16(m, v5, v14)
	mBase = m.M
	v77 = v44
	goto L8
L11:
	;
	if v35 != v14 {
		goto L9
	} else {
		goto L17
	}
L12:
	;
	v23 = v15
	goto L13
L13:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v23))))
	if v27 == int32(123) {
		v35 = v23
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v31 = v23 + int32(1)
	if v31 != v14 {
		v23 = v31
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	goto L10
L18:
	;
	v74 = F_crc16(m, v5+v35+int32(1), v52+(v35^int32(-1)))
	mBase = m.M
	v77 = v74
	goto L8
L19:
	;
	v67 = F_crc16(m, v5, v14)
	mBase = m.M
	v77 = v67
	goto L8
L20:
	;
	v52 = v46
	goto L22
L21:
	;
	if v52 == v14 {
		goto L19
	} else {
		goto L26
	}
L22:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v52))))
	if v54 == int32(125) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v58 = v52 + int32(1)
	if v58 != v14 {
		v52 = v58
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	if v52 != v46 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	goto L19
L28:
	;
	return v148 & int32(16383)
L29:
	;
	goto L28
L30:
	;
	v117 = v106 + int32(1)
	if v85 <= v117 {
		goto L40
	} else {
		goto L41
	}
L31:
	;
	v115 = F_crc16(m, v5, v85)
	mBase = m.M
	v148 = v115
	goto L29
L32:
	;
	if v106 != v85 {
		goto L30
	} else {
		goto L38
	}
L33:
	;
	v94 = v86
	goto L34
L34:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v94))))
	if v98 == int32(123) {
		v106 = v94
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v102 = v94 + int32(1)
	if v102 != v85 {
		v94 = v102
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L31
L38:
	;
	goto L31
L39:
	;
	v145 = F_crc16(m, v5+v106+int32(1), v123+(v106^int32(-1)))
	mBase = m.M
	v148 = v145
	goto L29
L40:
	;
	v138 = F_crc16(m, v5, v85)
	mBase = m.M
	v148 = v138
	goto L29
L41:
	;
	v123 = v117
	goto L43
L42:
	;
	if v123 == v85 {
		goto L40
	} else {
		goto L47
	}
L43:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v123))))
	if v125 == int32(125) {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v129 = v123 + int32(1)
	if v129 != v85 {
		v123 = v129
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L40
L47:
	;
	if v123 != v117 {
		goto L39
	} else {
		goto L48
	}
L48:
	;
	goto L40
L49:
	;
	return v219 & int32(16383)
L50:
	;
	goto L49
L51:
	;
	v188 = v177 + int32(1)
	if v156 <= v188 {
		goto L61
	} else {
		goto L62
	}
L52:
	;
	v186 = F_crc16(m, v5, v156)
	mBase = m.M
	v219 = v186
	goto L50
L53:
	;
	if v177 != v156 {
		goto L51
	} else {
		goto L59
	}
L54:
	;
	v165 = v157
	goto L55
L55:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v165))))
	if v169 == int32(123) {
		v177 = v165
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v173 = v165 + int32(1)
	if v173 != v156 {
		v165 = v173
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L52
L59:
	;
	goto L52
L60:
	;
	v216 = F_crc16(m, v5+v177+int32(1), v194+(v177^int32(-1)))
	mBase = m.M
	v219 = v216
	goto L50
L61:
	;
	v209 = F_crc16(m, v5, v156)
	mBase = m.M
	v219 = v209
	goto L50
L62:
	;
	v194 = v188
	goto L64
L63:
	;
	if v194 == v156 {
		goto L61
	} else {
		goto L68
	}
L64:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v194))))
	if v196 == int32(125) {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v200 = v194 + int32(1)
	if v200 != v156 {
		v194 = v200
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L61
L68:
	;
	if v194 != v188 {
		goto L60
	} else {
		goto L69
	}
L69:
	;
	goto L61
L70:
	;
	return v290 & int32(16383)
L71:
	;
	goto L70
L72:
	;
	v259 = v248 + int32(1)
	if v227 <= v259 {
		goto L82
	} else {
		goto L83
	}
L73:
	;
	v257 = F_crc16(m, v5, v227)
	mBase = m.M
	v290 = v257
	goto L71
L74:
	;
	if v248 != v227 {
		goto L72
	} else {
		goto L80
	}
L75:
	;
	v236 = v228
	goto L76
L76:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v236))))
	if v240 == int32(123) {
		v248 = v236
		goto L74
	} else {
		goto L78
	}
L78:
	;
	v244 = v236 + int32(1)
	if v244 != v227 {
		v236 = v244
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L73
L80:
	;
	goto L73
L81:
	;
	v287 = F_crc16(m, v5+v248+int32(1), v265+(v248^int32(-1)))
	mBase = m.M
	v290 = v287
	goto L71
L82:
	;
	v280 = F_crc16(m, v5, v227)
	mBase = m.M
	v290 = v280
	goto L71
L83:
	;
	v265 = v259
	goto L85
L84:
	;
	if v265 == v227 {
		goto L82
	} else {
		goto L89
	}
L85:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v265))))
	if v267 == int32(125) {
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v271 = v265 + int32(1)
	if v271 != v227 {
		v265 = v271
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L82
L89:
	;
	if v265 != v259 {
		goto L81
	} else {
		goto L90
	}
L90:
	;
	goto L82
L91:
	;
	return v362 & int32(16383)
L92:
	;
	goto L91
L93:
	;
	v331 = v320 + int32(1)
	if v299 <= v331 {
		goto L103
	} else {
		goto L104
	}
L94:
	;
	v329 = F_crc16(m, v5, v299)
	mBase = m.M
	v362 = v329
	goto L92
L95:
	;
	if v320 != v299 {
		goto L93
	} else {
		goto L101
	}
L96:
	;
	v308 = v300
	goto L97
L97:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v308))))
	if v312 == int32(123) {
		v320 = v308
		goto L95
	} else {
		goto L99
	}
L99:
	;
	v316 = v308 + int32(1)
	if v316 != v299 {
		v308 = v316
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L94
L101:
	;
	goto L94
L102:
	;
	v359 = F_crc16(m, v5+v320+int32(1), v337+(v320^int32(-1)))
	mBase = m.M
	v362 = v359
	goto L92
L103:
	;
	v352 = F_crc16(m, v5, v299)
	mBase = m.M
	v362 = v352
	goto L92
L104:
	;
	v337 = v331
	goto L106
L105:
	;
	if v337 == v299 {
		goto L103
	} else {
		goto L110
	}
L106:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v337))))
	if v339 == int32(125) {
		goto L105
	} else {
		goto L108
	}
L108:
	;
	v343 = v337 + int32(1)
	if v343 != v299 {
		v337 = v343
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L103
L110:
	;
	if v337 != v331 {
		goto L102
	} else {
		goto L111
	}
L111:
	;
	goto L103
}
func F_VM_ClusterKeySlotC(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	v3 = int32(0)
	if l1 < int32(1) {
		v23 = v3
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v65 & int32(16383)
L2:
	;
	goto L1
L3:
	;
	v34 = v23 + int32(1)
	if l1 <= v34 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v32 = F_crc16(m, l0, l1)
	mBase = m.M
	v65 = v32
	goto L2
L5:
	;
	if v23 != l1 {
		goto L3
	} else {
		goto L11
	}
L6:
	;
	v11 = v3
	goto L7
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v11))))
	if v15 == int32(123) {
		v23 = v11
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v19 = v11 + int32(1)
	if v19 != l1 {
		v11 = v19
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	goto L4
L12:
	;
	v62 = F_crc16(m, l0+v23+int32(1), v40+(v23^int32(-1)))
	mBase = m.M
	v65 = v62
	goto L2
L13:
	;
	v55 = F_crc16(m, l0, l1)
	mBase = m.M
	v65 = v55
	goto L2
L14:
	;
	v40 = v34
	goto L16
L15:
	;
	if v40 == l1 {
		goto L13
	} else {
		goto L20
	}
L16:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v40))))
	if v42 == int32(125) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v46 = v40 + int32(1)
	if v46 != l1 {
		v40 = v46
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L13
L20:
	;
	if v40 != v34 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	goto L13
}
func F_VM_CommandFilterArgGet(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = int32(0)
	if l1 < v3 {
		v14 = v3
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v7 <= l1 {
			v14 = v3
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v9+l1<<(uint(int32(2))%32))))
			v14 = v13
		}
	}
	return v14
}
func F_VM_CreateStringFromLongDouble(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v8 = m.G0
	v9 = int32(5120)
	v10 = v8 - v9
	m.G0 = v10
	v15 = F_ld2string(m, v10, v9, l1, l2, base.B2i32(l3 != int32(0)))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = F_createStringObject_1(m, v10, v15)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if l0 == int32(0) {
				m.G0 = v10 + int32(5120)
				return v19
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
				if v23&int32(1) == int32(0) {
					m.G0 = v10 + int32(5120)
					return v19
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v28 == v29 {
						v32 = int32(8)
						if v32 < v28 {
							v35 = v28
						} else {
							v35 = v32
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v35 << (uint(int32(1)) % 32)
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v42 = F_valkey_realloc(m, v39, v35<<(uint(int32(4))%32))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v42
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v46 = v45
							v47 = v42
							v50 = v47 + v46<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v50))) = v19
							v52 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v46 + v52
							m.G0 = v10 + int32(5120)
							return v19
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v46 = v28
						v47 = v31
						v50 = v47 + v46<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v50))) = v19
						v52 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v52
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v46 + v52
						m.G0 = v10 + int32(5120)
						return v19
					}
				}
			}
		}
	}
}
func F_VM_DbSize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+96))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = F_kvstoreSize(m, v4)
	mBase = m.M
	return v5
}
func F_VM_DefragAlloc(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F_VM_DeleteKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v2&int32(2) != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v12 = F_dbDelete(m, v10, v11)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v16
				return v16
			}
		} else {
			return int32(0)
		}
	} else {
		return int32(1)
	}
}
func F_VM_DictCompare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	v9 = l0 + int32(4)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v10&int32(2) != 0 {
		v97 = int32(1)
	} else {
		v13 = F_objectGetVal(m, l2)
		mBase = m.M
		v15 = F_objectGetVal(m, l2)
		mBase = m.M
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(-1)))))
		switch v18 & int32(7) {
		case 0:
			v35 = int32(base.Ui32(v18) >> (uint(int32(3)) % 32))
		case 1:
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(-3)))))
			v35 = v25
		case 2:
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+int32(-5)))))
			v35 = v28
		case 3:
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(-9))))
			v35 = v31
		case 4:
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(-17))))
			v35 = v34
		default:
			v35 = int32(0)
		}
		v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v41 == int32(61) {
			v50 = int32(1)
		} else {
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			if v45 != int32(61) {
				v50 = int32(0)
			} else {
				v50 = int32(1)
			}
		}
		v53 = base.B2i32(v41 == int32(62))
		if v41 == int32(62) {
			v62 = int32(1)
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v65 = base.B2i32(base.Ui32(v35) < base.Ui32(v64))
			if base.Ui32(v35) < base.Ui32(v64) {
				v66 = v35
			} else {
				v66 = v64
			}
			v67 = F_memcmp(m, v63, v13, v66)
			mBase = m.M
			if v41 == int32(62) {
				if v67 != 0 {
					if v67 < int32(1) {
						v93 = v62 ^ int32(1)
					} else {
						v93 = base.B2i32(v41 == int32(62))
					}
				} else {
					if v50&base.B2i32(v35 == v64) == int32(0) {
						if v62 != 0 {
							v93 = v65 & base.B2i32(v41 == int32(62))
						} else {
							v93 = base.B2i32(base.Ui32(v64) < base.Ui32(v35))
						}
					} else {
						v93 = int32(1)
					}
				}
			} else {
				if v62 == int32(0) {
					if v67 != 0 {
						if v67 < int32(1) {
							v93 = v62 ^ int32(1)
						} else {
							v93 = base.B2i32(v41 == int32(62))
						}
					} else {
						if v50&base.B2i32(v35 == v64) == int32(0) {
							if v62 != 0 {
								v93 = v65 & base.B2i32(v41 == int32(62))
							} else {
								v93 = base.B2i32(base.Ui32(v64) < base.Ui32(v35))
							}
						} else {
							v93 = int32(1)
						}
					}
				} else {
					v93 = base.B2i32(v67 == int32(0)) & base.B2i32(v35 == v64)
				}
			}
		} else {
			if v41 == int32(60) {
				v62 = int32(0)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
				v65 = base.B2i32(base.Ui32(v35) < base.Ui32(v64))
				if base.Ui32(v35) < base.Ui32(v64) {
					v66 = v35
				} else {
					v66 = v64
				}
				v67 = F_memcmp(m, v63, v13, v66)
				mBase = m.M
				if v41 == int32(62) {
					if v67 != 0 {
						if v67 < int32(1) {
							v93 = v62 ^ int32(1)
						} else {
							v93 = base.B2i32(v41 == int32(62))
						}
					} else {
						if v50&base.B2i32(v35 == v64) == int32(0) {
							if v62 != 0 {
								v93 = v65 & base.B2i32(v41 == int32(62))
							} else {
								v93 = base.B2i32(base.Ui32(v64) < base.Ui32(v35))
							}
						} else {
							v93 = int32(1)
						}
					}
				} else {
					if v62 == int32(0) {
						if v67 != 0 {
							if v67 < int32(1) {
								v93 = v62 ^ int32(1)
							} else {
								v93 = base.B2i32(v41 == int32(62))
							}
						} else {
							if v50&base.B2i32(v35 == v64) == int32(0) {
								if v62 != 0 {
									v93 = v65 & base.B2i32(v41 == int32(62))
								} else {
									v93 = base.B2i32(base.Ui32(v64) < base.Ui32(v35))
								}
							} else {
								v93 = int32(1)
							}
						}
					} else {
						v93 = base.B2i32(v67 == int32(0)) & base.B2i32(v35 == v64)
					}
				}
			} else {
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				if v58 == int32(61) {
					v62 = int32(1)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
					v65 = base.B2i32(base.Ui32(v35) < base.Ui32(v64))
					if base.Ui32(v35) < base.Ui32(v64) {
						v66 = v35
					} else {
						v66 = v64
					}
					v67 = F_memcmp(m, v63, v13, v66)
					mBase = m.M
					if v41 == int32(62) {
						if v67 != 0 {
							if v67 < int32(1) {
								v93 = v62 ^ int32(1)
							} else {
								v93 = base.B2i32(v41 == int32(62))
							}
						} else {
							if v50&base.B2i32(v35 == v64) == int32(0) {
								if v62 != 0 {
									v93 = v65 & base.B2i32(v41 == int32(62))
								} else {
									v93 = base.B2i32(base.Ui32(v64) < base.Ui32(v35))
								}
							} else {
								v93 = int32(1)
							}
						}
					} else {
						if v62 == int32(0) {
							if v67 != 0 {
								if v67 < int32(1) {
									v93 = v62 ^ int32(1)
								} else {
									v93 = base.B2i32(v41 == int32(62))
								}
							} else {
								if v50&base.B2i32(v35 == v64) == int32(0) {
									if v62 != 0 {
										v93 = v65 & base.B2i32(v41 == int32(62))
									} else {
										v93 = base.B2i32(base.Ui32(v64) < base.Ui32(v35))
									}
								} else {
									v93 = int32(1)
								}
							}
						} else {
							v93 = base.B2i32(v67 == int32(0)) & base.B2i32(v35 == v64)
						}
					}
				} else {
					v93 = int32(0)
				}
			}
		}
		v97 = base.B2i32(v93 == int32(0))
	}
	return v97
}
func F_VM_DictGet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_objectGetVal(m, l1)
	mBase = m.M
	v14 = F_objectGetVal(m, l1)
	mBase = m.M
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-1)))))
	switch v17 & int32(7) {
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
		v34 = int32(0)
		goto L1
	}
L1:
	;
	v35 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = v10 + int32(12)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v34 == v35 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-17))))
	v34 = v33
	goto L1
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-9))))
	v34 = v30
	goto L1
L4:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(-5)))))
	v34 = v27
	goto L1
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-3)))))
	v34 = v24
	goto L1
L6:
	;
	v34 = int32(base.Ui32(v17) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	if l2 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L8:
	;
	if v192 != v34 {
		v233 = v35
		goto L35
	} else {
		goto L36
	}
L9:
	;
	v183 = int32(0)
	v189 = v48
	v190 = v49
	v192 = v183
	v196 = v183
	goto L8
L10:
	;
	if base.Ui32(v49) < base.Ui32(int32(8)) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v60 = v48
	v61 = v49
	v63 = int32(0)
	goto L13
L12:
	;
	v189 = v173
	v190 = v174
	v192 = v176
	v196 = base.B2i32(v179 != int32(0))
	goto L8
L13:
	;
	v69 = int32(base.Ui32(v61) >> (uint(int32(3)) % 32))
	v70 = int32(4)
	v71 = v60 + v70
	if v61&v70 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v173 = v164
	v174 = v165
	v176 = v149
	v179 = v154
	goto L12
L15:
	;
	v154 = int32(0)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v71+v69+(v154-v69)&int32(3)+v142<<(uint(int32(2))%32))))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if base.Ui32(v165) < base.Ui32(int32(8)) {
		v173 = v164
		v174 = v165
		v176 = v149
		v179 = v154
		goto L12
	} else {
		goto L33
	}
L16:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v63))))
	v120 = int32(0)
	goto L27
L17:
	;
	v76 = int32(0)
	if base.Ui32(v34) <= base.Ui32(v63) {
		v109 = v63
		v112 = v76
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v112 == v69 {
		v142 = v76
		v149 = v109
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v86 = v63
	v89 = v76
	goto L20
L20:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v89))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v86))))
	if v92 != v94 {
		v109 = v86
		v112 = v89
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v109 = v97
	v112 = v99
	goto L18
L22:
	;
	v96 = int32(1)
	v97 = v86 + v96
	v99 = v89 + v96
	if base.Ui32(v69) <= base.Ui32(v99) {
		v109 = v97
		v112 = v99
		goto L18
	} else {
		goto L23
	}
L23:
	;
	if base.Ui32(v97) < base.Ui32(v34) {
		v86 = v97
		v89 = v99
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v173 = v60
	v174 = v61
	v176 = v109
	v179 = v112
	goto L12
L26:
	;
	if v120 != v69 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v120))))
	if v133 == v117&int32(255) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v135 = int32(1)
	v137 = v120 + v135
	if v137 != v69 {
		v120 = v137
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v189 = v60
	v190 = v61
	v192 = v63
	v196 = v135
	goto L8
L31:
	;
	v142 = v120
	v149 = v63 + int32(1)
	goto L15
L32:
	;
	v173 = v60
	v174 = v61
	v176 = v63
	v179 = v69
	goto L12
L33:
	;
	if base.Ui32(v149) < base.Ui32(v34) {
		v60 = v164
		v61 = v165
		v63 = v149
		goto L13
	} else {
		goto L34
	}
L34:
	;
	goto L14
L35:
	;
	goto L7
L36:
	;
	v198 = int32(0)
	if v190&int32(1) == v198 {
		v233 = v198
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v204 = v190 & int32(4)
	if v196&base.B2i32(v204 != int32(0)) != 0 {
		v233 = v198
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v208 = int32(1)
	if v39 == int32(0) {
		v233 = v208
		goto L35
	} else {
		goto L39
	}
L39:
	;
	if v190&int32(2) != 0 {
		v230 = int32(0)
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v230
	v233 = v208
	goto L35
L41:
	;
	v214 = int32(3)
	v215 = int32(base.Ui32(v190) >> (uint(v214) % 32))
	if v204 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v225 = int32(4)
	goto L44
L43:
	;
	v225 = v215 << (uint(int32(2)) % 32)
	goto L44
L44:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v189+v215+(int32(0)-v215)&v214+v225+int32(4))))
	v230 = v229
	goto L40
L45:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	m.G0 = v10 + int32(16)
	return v240
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = base.B2i32(v233 == int32(0))
	goto L45
}
func F_VM_DictGetC(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	v5 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v5
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = v8 + int32(12)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if l2 == v5 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L2:
	;
	if v167 != l2 {
		v208 = v5
		goto L29
	} else {
		goto L30
	}
L3:
	;
	v158 = int32(0)
	v164 = v23
	v165 = v24
	v167 = v158
	v171 = v158
	goto L2
L4:
	;
	if base.Ui32(v24) < base.Ui32(int32(8)) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v35 = v23
	v36 = v24
	v38 = int32(0)
	goto L7
L6:
	;
	v164 = v148
	v165 = v149
	v167 = v151
	v171 = base.B2i32(v154 != int32(0))
	goto L2
L7:
	;
	v44 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
	v45 = int32(4)
	v46 = v35 + v45
	if v36&v45 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v148 = v139
	v149 = v140
	v151 = v124
	v154 = v129
	goto L6
L9:
	;
	v129 = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v46+v44+(v129-v44)&int32(3)+v117<<(uint(int32(2))%32))))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if base.Ui32(v140) < base.Ui32(int32(8)) {
		v148 = v139
		v149 = v140
		v151 = v124
		v154 = v129
		goto L6
	} else {
		goto L27
	}
L10:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v38))))
	v95 = int32(0)
	goto L21
L11:
	;
	v51 = int32(0)
	if base.Ui32(l2) <= base.Ui32(v38) {
		v84 = v38
		v87 = v51
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v87 == v44 {
		v117 = v51
		v124 = v84
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v61 = v38
	v64 = v51
	goto L14
L14:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v64))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v61))))
	if v67 != v69 {
		v84 = v61
		v87 = v64
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v84 = v72
	v87 = v74
	goto L12
L16:
	;
	v71 = int32(1)
	v72 = v61 + v71
	v74 = v64 + v71
	if base.Ui32(v44) <= base.Ui32(v74) {
		v84 = v72
		v87 = v74
		goto L12
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(v72) < base.Ui32(l2) {
		v61 = v72
		v64 = v74
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v148 = v35
	v149 = v36
	v151 = v84
	v154 = v87
	goto L6
L20:
	;
	if v95 != v44 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v95))))
	if v108 == v92&int32(255) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v110 = int32(1)
	v112 = v95 + v110
	if v112 != v44 {
		v95 = v112
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v164 = v35
	v165 = v36
	v167 = v38
	v171 = v110
	goto L2
L25:
	;
	v117 = v95
	v124 = v38 + int32(1)
	goto L9
L26:
	;
	v148 = v35
	v149 = v36
	v151 = v38
	v154 = v44
	goto L6
L27:
	;
	if base.Ui32(v124) < base.Ui32(l2) {
		v35 = v139
		v36 = v140
		v38 = v124
		goto L7
	} else {
		goto L28
	}
L28:
	;
	goto L8
L29:
	;
	goto L1
L30:
	;
	v173 = int32(0)
	if v165&int32(1) == v173 {
		v208 = v173
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v179 = v165 & int32(4)
	if v171&base.B2i32(v179 != int32(0)) != 0 {
		v208 = v173
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v183 = int32(1)
	if v14 == int32(0) {
		v208 = v183
		goto L29
	} else {
		goto L33
	}
L33:
	;
	if v165&int32(2) != 0 {
		v205 = int32(0)
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v205
	v208 = v183
	goto L29
L35:
	;
	v189 = int32(3)
	v190 = int32(base.Ui32(v165) >> (uint(v189) % 32))
	if v179 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v200 = int32(4)
	goto L38
L37:
	;
	v200 = v190 << (uint(int32(2)) % 32)
	goto L38
L38:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v164+v190+(int32(0)-v190)&v189+v200+int32(4))))
	v205 = v204
	goto L34
L39:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	m.G0 = v8 + int32(16)
	return v215
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = base.B2i32(v208 == int32(0))
	goto L39
}
func F_VM_DictPrevC(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v6 = F_raxPrev(m, l0+int32(4))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			if l1 == int32(0) {
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
			}
			if l2 == int32(0) {
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
			}
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			return v20
		} else {
			return int32(0)
		}
	}
}
func F_VM_DictReplaceC(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_raxInsert(m, v5, l1, l2, l3, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(1))
	}
}
func F_VM_DictSize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)+8))
	return v3
}
func F_VM_EventLoopAdd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v87 int32
	_ = v87
	if l0 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(68)
		return int32(1)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[279]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if l0 < v12 {
			if base.Ui32(int32(3)) < base.Ui32(l1) {
				*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
				return int32(1)
			} else {
				if l2 != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, _consts[279]))
					v28 = int32(0)
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					if v30 <= l0 {
						v41 = v28
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
						v35 = v32 + l0<<(uint(int32(4))%32)
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
						if v36 == int32(0) {
							v41 = v28
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
							v41 = v39
						}
					}
					if v41 != 0 {
						v47 = v41
						v48 = int32(1)
						v50 = *(*int32)(unsafe.Add(mBase, _consts[279]))
						v54 = l1 & v48
						if v54 != 0 {
							v55 = int32(567)
						} else {
							v55 = int32(568)
						}
						v56 = F_aeCreateFileEvent(m, v50, l0, l1, v55, v47)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							if v56 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = l3
								if v54 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2
								}
								if base.Ui32(l1) < base.Ui32(int32(2)) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = l2
								}
								*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(0)
								v87 = int32(0)
								return v87
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, _consts[279]))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
								if v64 <= l0 {
									v71 = int32(0)
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+l0<<(uint(int32(4))%32))))
									v71 = v70
								}
								if v71 != 0 {
									v87 = v48
									return v87
								} else {
									F_valkey_free(m, v47)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										return int32(1)
									}
								}
							}
						}
					} else {
						v43 = F_valkey_calloc(m, int32(12))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = v43
							v48 = int32(1)
							v50 = *(*int32)(unsafe.Add(mBase, _consts[279]))
							v54 = l1 & v48
							if v54 != 0 {
								v55 = int32(567)
							} else {
								v55 = int32(568)
							}
							v56 = F_aeCreateFileEvent(m, v50, l0, l1, v55, v47)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = l3
									if v54 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2
									}
									if base.Ui32(l1) < base.Ui32(int32(2)) {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = l2
									}
									*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(0)
									v87 = int32(0)
									return v87
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, _consts[279]))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
									if v64 <= l0 {
										v71 = int32(0)
									} else {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+l0<<(uint(int32(4))%32))))
										v71 = v70
									}
									if v71 != 0 {
										v87 = v48
										return v87
									} else {
										F_valkey_free(m, v47)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											return int32(1)
										}
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
					return int32(1)
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(68)
			return int32(1)
		}
	}
}
func F_VM_EventLoopAddOneShot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	if l0 != 0 {
		v7 = F_valkey_malloc(m, int32(8))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			v15 = int32(0)
			v17 = *(*int32)(unsafe.Add(mBase, _consts[566]))
			if v17 != 0 {
				v22 = v17
				v23 = F_listAddNodeTail(m, v22, v7)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, _consts[558]))
					v31 = F_write(m, v28, int32(_a941), int32(1))
					mBase = m.M
					v33 = int32(0)
					v34 = v15
					*(*int32)(unsafe.Add(mBase, _consts[5])) = v34
					return v33
				}
			} else {
				v19 = F_listCreate(m)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[566])) = v19
					v22 = v19
					v23 = F_listAddNodeTail(m, v22, v7)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, _consts[558]))
						v31 = F_write(m, v28, int32(_a941), int32(1))
						mBase = m.M
						v33 = int32(0)
						v34 = v15
						*(*int32)(unsafe.Add(mBase, _consts[5])) = v34
						return v33
					}
				}
			}
		}
	} else {
		v33 = int32(1)
		v34 = int32(28)
		*(*int32)(unsafe.Add(mBase, _consts[5])) = v34
		return v33
	}
}
func F_VM_Fork(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = F_serverFork(m, int32(4))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		switch v10 + int32(1) {
		case 0:
			v55 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(3) < v55 {
				m.G0 = v7 + int32(32)
				return v10
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				v60 = F___strerror_l(m, v59, v59)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v60
				F__serverLog(m, int32(3), int32(_a950), v7+int32(16))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(32)
					return v10
				}
			}
		case 1:
			v17 = *(*int32)(unsafe.Add(mBase, _consts[54]))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v19 = int32(_a127)
			v22 = int32(*(*int8)(unsafe.Add(mBase, _consts[55])))
			if v22 != 0 {
				v23 = int32(0)
				v24 = F_strchr(m, v18, v22)
				mBase = m.M
				if v24 == v23 {
					v44 = v23
					v47 = v44
				} else {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
					if v27 != 0 {
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
						if v28 == int32(0) {
							v44 = v23
							v47 = v44
						} else {
							v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[57])))
							if v31 != 0 {
								v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
								if v33 == int32(0) {
									v44 = v23
									v47 = v44
								} else {
									v36 = int32(*(*uint8)(unsafe.Add(mBase, _consts[58])))
									if v36 != 0 {
										v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)))
										if v38 == int32(0) {
											v44 = v23
											v47 = v44
										} else {
											v41 = int32(*(*uint8)(unsafe.Add(mBase, _consts[59])))
											if v41 != 0 {
												v43 = F_twoway_strstr(m, v24, v19)
												mBase = m.M
												v44 = v43
												v47 = v44
											} else {
												v42 = F_fourbyte_strstr(m, v24, v19)
												mBase = m.M
												v47 = v42
											}
										}
									} else {
										v37 = F_threebyte_strstr(m, v24, v19)
										mBase = m.M
										v47 = v37
									}
								}
							} else {
								v32 = F_twobyte_strstr(m, v24, v19)
								mBase = m.M
								v47 = v32
							}
						}
					} else {
						v47 = v24
					}
				}
			} else {
				v47 = v18
			}
			if v47 == int32(0) {
			} else {
			}
			m.G0 = v7 + int32(32)
			return v10
		default:
			v68 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[580])) = l1
			*(*int32)(unsafe.Add(mBase, _consts[581])) = l0
			v73 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(1) < v73 {
				m.G0 = v7 + int32(32)
				return v10
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
				F__serverLog(m, int32(1), int32(_a951), v7)
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(32)
					return v10
				}
			}
		}
	}
}
func F_VM_GetBlockedClientReadyKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	return v2
}
func F_VM_GetClientInfoById(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	v5 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	if v5 == int32(0) {
		v12 = m.G0
		v13 = int32(16)
		v14 = v12 - v13
		m.G0 = v14
		v16 = int64(56)
		v18 = int64(65280)
		v20 = int64(40)
		v23 = int64(16711680)
		v25 = int64(24)
		v27 = int64(4278190080)
		v29 = int64(8)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = l1<<(uint(v16)%64) | l1&v18<<(uint(v20)%64) | (l1&v23<<(uint(v25)%64) | l1&v27<<(uint(v29)%64)) | (int64(base.Ui64(l1)>>(uint(v29)%64))&v27 | int64(base.Ui64(l1)>>(uint(v25)%64))&v23 | (int64(base.Ui64(l1)>>(uint(v20)%64))&v18 | int64(base.Ui64(l1)>>(uint(v16)%64))))
		*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(0)
		v55 = *(*int32)(unsafe.Add(mBase, _consts[404]))
		v56 = int32(8)
		v61 = F_raxFind(m, v55, v14+v56, v56, v14+int32(4))
		mBase = m.M
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		m.G0 = v14 + v13
		if v62 != 0 {
			v68 = v62
			if l0 != 0 {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v72 = F_modulePopulateClientInfoStructure(m, l0, v68, v71)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					return v72
				}
			} else {
				return int32(0)
			}
		} else {
			return int32(1)
		}
	} else {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
		if v8 == l1 {
			v68 = v5
			if l0 != 0 {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v72 = F_modulePopulateClientInfoStructure(m, l0, v68, v71)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					return v72
				}
			} else {
				return int32(0)
			}
		} else {
			v12 = m.G0
			v13 = int32(16)
			v14 = v12 - v13
			m.G0 = v14
			v16 = int64(56)
			v18 = int64(65280)
			v20 = int64(40)
			v23 = int64(16711680)
			v25 = int64(24)
			v27 = int64(4278190080)
			v29 = int64(8)
			*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = l1<<(uint(v16)%64) | l1&v18<<(uint(v20)%64) | (l1&v23<<(uint(v25)%64) | l1&v27<<(uint(v29)%64)) | (int64(base.Ui64(l1)>>(uint(v29)%64))&v27 | int64(base.Ui64(l1)>>(uint(v25)%64))&v23 | (int64(base.Ui64(l1)>>(uint(v20)%64))&v18 | int64(base.Ui64(l1)>>(uint(v16)%64))))
			*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(0)
			v55 = *(*int32)(unsafe.Add(mBase, _consts[404]))
			v56 = int32(8)
			v61 = F_raxFind(m, v55, v14+v56, v56, v14+int32(4))
			mBase = m.M
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			m.G0 = v14 + v13
			if v62 != 0 {
				v68 = v62
				if l0 != 0 {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v72 = F_modulePopulateClientInfoStructure(m, l0, v68, v71)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						return v72
					}
				} else {
					return int32(0)
				}
			} else {
				return int32(1)
			}
		}
	}
}
func F_VM_GetClientNameById(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	v3 = int32(0)
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	v13 = int64(56)
	v15 = int64(65280)
	v17 = int64(40)
	v20 = int64(16711680)
	v22 = int64(24)
	v24 = int64(4278190080)
	v26 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = l1<<(uint(v13)%64) | l1&v15<<(uint(v17)%64) | (l1&v20<<(uint(v22)%64) | l1&v24<<(uint(v26)%64)) | (int64(base.Ui64(l1)>>(uint(v26)%64))&v24 | int64(base.Ui64(l1)>>(uint(v22)%64))&v20 | (int64(base.Ui64(l1)>>(uint(v17)%64))&v15 | int64(base.Ui64(l1)>>(uint(v13)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v3
	v52 = *(*int32)(unsafe.Add(mBase, _consts[404]))
	v53 = int32(8)
	v58 = F_raxFind(m, v52, v11+v53, v53, v11+int32(4))
	mBase = m.M
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	m.G0 = v11 + v10
	if v59 == int32(0) {
		v108 = v3
		return v108
	} else {
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)+348))
		if v65 == int32(0) {
			v108 = v3
			return v108
		} else {
			F_incrRefCount(m, v65)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int32(0)
			} else {
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
				if v72&int32(1) == int32(0) {
					v108 = v65
					return v108
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v77 == v78 {
						v81 = int32(8)
						if v81 < v77 {
							v84 = v77
						} else {
							v84 = v81
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v84 << (uint(int32(1)) % 32)
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v91 = F_valkey_realloc(m, v88, v84<<(uint(int32(4))%32))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v91
							v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v95 = v94
							v96 = v91
							v99 = v96 + v95<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v99))) = v65
							v101 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v101
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v95 + v101
							v108 = v65
							return v108
						}
					} else {
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v95 = v77
						v96 = v80
						v99 = v96 + v95<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v99))) = v65
						v101 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v101
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v95 + v101
						v108 = v65
						return v108
					}
				}
			}
		}
	}
}
func F_VM_GetClusterNodeInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_moduleGetClusterNodeInfoForClient(m, l5, int32(0), l1, l2, l3, l4, l5)
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_VM_GetClusterNodesList(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v4 != 0 {
		v7 = F_getClusterNodesList(m, l1)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v7
		}
	} else {
		return int32(0)
	}
}
func F_VM_GetCommandKeysWithFlags(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v66 int64
	_ = v66
	var v73 int64
	_ = v73
	var v80 int64
	_ = v80
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int64
	_ = v123
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int64
	_ = v214
	var v225 int64
	_ = v225
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v240 int64
	_ = v240
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v255 int64
	_ = v255
	var v263 int64
	_ = v263
	var v265 int64
	_ = v265
	var v268 int64
	_ = v268
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v283 int64
	_ = v283
	var v291 int64
	_ = v291
	var v293 int64
	_ = v293
	var v296 int64
	_ = v296
	var v304 int64
	_ = v304
	var v306 int64
	_ = v306
	var v311 int64
	_ = v311
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v324 int64
	_ = v324
	var v332 int64
	_ = v332
	var v334 int64
	_ = v334
	var v339 int64
	_ = v339
	var v347 int64
	_ = v347
	var v349 int64
	_ = v349
	var v352 int64
	_ = v352
	var v360 int64
	_ = v360
	var v362 int64
	_ = v362
	var v367 int64
	_ = v367
	var v375 int64
	_ = v375
	var v377 int64
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v391 int32
	_ = v391
	v10 = m.G0
	v12 = v10 - int32(2064)
	m.G0 = v12
	v14 = F_lookupCommand(m, l1, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v28 = int32(1)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
			if v29 != 0 {
				v131 = v28
				v144 = v131
			} else {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+58)))
				if v30&int32(32) != 0 {
					v131 = v28
					v144 = v131
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
					if int32(1) <= v33 {
						v38 = v33 & int32(3)
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
						if base.Ui32(int32(4)) <= base.Ui32(v33) {
							v47 = int32(0)
							v50 = v47
							v52 = v47
							v55 = int64(0)
							for {
								v57 = int32(48)
								v60 = *(*int64)(unsafe.Add(mBase, uint32(v39+v50*v57)+8))
								v66 = *(*int64)(unsafe.Add(mBase, uint32(v39+(v50|int32(1))*v57)+8))
								v73 = *(*int64)(unsafe.Add(mBase, uint32(v39+(v50|int32(2))*v57)+8))
								v80 = *(*int64)(unsafe.Add(mBase, uint32(v39+(v50|int32(3))*v57)+8))
								v84 = v55 | (v60&v66&v73&v80 ^ int64(-1))
								v85 = int32(4)
								v86 = v50 + v85
								v88 = v52 + v85
								if v88 != v33&int32(2147483644) {
									v50 = v86
									v52 = v88
									v55 = v84
									continue
								} else {
									break
								}
								break
							}
							v90 = v86
							v95 = v84
						} else {
							v90 = int32(0)
							v95 = int64(0)
						}
						if v38 == int32(0) {
							v123 = v95
						} else {
							v99 = v90
							v103 = int32(0)
							v104 = v95
							for {
								v109 = *(*int64)(unsafe.Add(mBase, uint32(v39+v99*int32(48))+8))
								v112 = v104 | (v109 ^ int64(-1))
								v113 = int32(1)
								v116 = v103 + v113
								if v116 != v38 {
									v99 = v99 + v113
									v103 = v116
									v104 = v112
									continue
								} else {
									break
								}
								break
							}
							v123 = v112
						}
						v131 = int32(base.Ui32(base.I32_wrap_i64(v123))>>(uint(int32(8))%32)) & int32(1)
						v144 = v131
					} else {
						v144 = int32(0)
					}
				}
			}
			if v144 != 0 {
				v149 = int32(0)
				v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
				if base.B2i32(v149 < v150)&base.B2i32(v150 != l2) != 0 {
					*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
					v391 = v149
					m.G0 = v12 + int32(2064)
					return v391
				} else {
					if int32(0)-v150 <= l2 {
						v161 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v161
						*(*int64)(unsafe.Add(mBase, uint32(v12)+4)) = int64(1099511627776)
						v168 = F_getKeysFromCommand(m, v14, l1, l2, v12+int32(4))
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return int32(0)
						} else {
							v170 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v170
							if v170 != 0 {
								v180 = v170 << (uint(int32(2)) % 32)
								v181 = F_valkey_malloc(m, v180)
								mBase = m.M
								v182 = m.ExcPending
								if v182 != 0 {
									return int32(0)
								} else {
									if l4 == int32(0) {
										v188 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
										if v188 < int32(1) {
											v391 = v181
										} else {
											v191 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
											v195 = int32(0)
											v199 = v188
											for {
												v203 = v195 << (uint(int32(2)) % 32)
												v207 = v191 + v195<<(uint(int32(3))%32)
												v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
												*(*int32)(unsafe.Add(mBase, uint32(v181+v203))) = v208
												if l4 == int32(0) {
													v380 = v199
												} else {
													v212 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
													v214 = int64(*(*int32)(unsafe.Add(mBase, uint32(v207)+4)))
													v225 = *(*int64)(unsafe.Add(mBase, _consts[567]))
													if base.B2i32(v225&v214 == int64(0)) == int32(0) {
														v236 = *(*int64)(unsafe.Add(mBase, _consts[530]))
														v237 = v236
													} else {
														v237 = int64(0)
													}
													v240 = *(*int64)(unsafe.Add(mBase, _consts[568]))
													if v240&v214 == int64(0) {
														v250 = v237
													} else {
														v248 = *(*int64)(unsafe.Add(mBase, _consts[531]))
														v250 = v248 | v237
													}
													v255 = *(*int64)(unsafe.Add(mBase, _consts[569]))
													if v255&v214 == int64(0) {
														v265 = v250
													} else {
														v263 = *(*int64)(unsafe.Add(mBase, _consts[532]))
														v265 = v263 | v250
													}
													v268 = *(*int64)(unsafe.Add(mBase, _consts[570]))
													if v268&v214 == int64(0) {
														v278 = v265
													} else {
														v276 = *(*int64)(unsafe.Add(mBase, _consts[533]))
														v278 = v276 | v265
													}
													v283 = *(*int64)(unsafe.Add(mBase, _consts[571]))
													if v283&v214 == int64(0) {
														v293 = v278
													} else {
														v291 = *(*int64)(unsafe.Add(mBase, _consts[534]))
														v293 = v291 | v278
													}
													v296 = *(*int64)(unsafe.Add(mBase, _consts[572]))
													if v296&v214 == int64(0) {
														v306 = v293
													} else {
														v304 = *(*int64)(unsafe.Add(mBase, _consts[535]))
														v306 = v304 | v293
													}
													v311 = *(*int64)(unsafe.Add(mBase, _consts[573]))
													if v311&v214 == int64(0) {
														v321 = v306
													} else {
														v319 = *(*int64)(unsafe.Add(mBase, _consts[536]))
														v321 = v319 | v306
													}
													v324 = *(*int64)(unsafe.Add(mBase, _consts[574]))
													if v324&v214 == int64(0) {
														v334 = v321
													} else {
														v332 = *(*int64)(unsafe.Add(mBase, _consts[537]))
														v334 = v332 | v321
													}
													v339 = *(*int64)(unsafe.Add(mBase, _consts[575]))
													if v339&v214 == int64(0) {
														v349 = v334
													} else {
														v347 = *(*int64)(unsafe.Add(mBase, _consts[538]))
														v349 = v347 | v334
													}
													v352 = *(*int64)(unsafe.Add(mBase, _consts[576]))
													if v352&v214 == int64(0) {
														v362 = v349
													} else {
														v360 = *(*int64)(unsafe.Add(mBase, _consts[539]))
														v362 = v360 | v349
													}
													v367 = *(*int64)(unsafe.Add(mBase, _consts[577]))
													if v367&v214 == int64(0) {
														v377 = v362
													} else {
														v375 = *(*int64)(unsafe.Add(mBase, _consts[540]))
														v377 = v375 | v362
													}
													*(*uint32)(unsafe.Add(mBase, uint32(v212+v203))) = uint32(v377)
													v379 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
													v380 = v379
												}
												v382 = v195 + int32(1)
												if v382 < v380 {
													v195 = v382
													v199 = v380
													continue
												} else {
													break
												}
												break
											}
											v391 = v181
										}
										m.G0 = v12 + int32(2064)
										return v391
									} else {
										v185 = F_valkey_malloc(m, v180)
										mBase = m.M
										v186 = m.ExcPending
										if v186 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l4))) = v185
											v188 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
											if v188 < int32(1) {
												v391 = v181
											} else {
												v191 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
												v195 = int32(0)
												v199 = v188
												for {
													v203 = v195 << (uint(int32(2)) % 32)
													v207 = v191 + v195<<(uint(int32(3))%32)
													v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
													*(*int32)(unsafe.Add(mBase, uint32(v181+v203))) = v208
													if l4 == int32(0) {
														v380 = v199
													} else {
														v212 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
														v214 = int64(*(*int32)(unsafe.Add(mBase, uint32(v207)+4)))
														v225 = *(*int64)(unsafe.Add(mBase, _consts[567]))
														if base.B2i32(v225&v214 == int64(0)) == int32(0) {
															v236 = *(*int64)(unsafe.Add(mBase, _consts[530]))
															v237 = v236
														} else {
															v237 = int64(0)
														}
														v240 = *(*int64)(unsafe.Add(mBase, _consts[568]))
														if v240&v214 == int64(0) {
															v250 = v237
														} else {
															v248 = *(*int64)(unsafe.Add(mBase, _consts[531]))
															v250 = v248 | v237
														}
														v255 = *(*int64)(unsafe.Add(mBase, _consts[569]))
														if v255&v214 == int64(0) {
															v265 = v250
														} else {
															v263 = *(*int64)(unsafe.Add(mBase, _consts[532]))
															v265 = v263 | v250
														}
														v268 = *(*int64)(unsafe.Add(mBase, _consts[570]))
														if v268&v214 == int64(0) {
															v278 = v265
														} else {
															v276 = *(*int64)(unsafe.Add(mBase, _consts[533]))
															v278 = v276 | v265
														}
														v283 = *(*int64)(unsafe.Add(mBase, _consts[571]))
														if v283&v214 == int64(0) {
															v293 = v278
														} else {
															v291 = *(*int64)(unsafe.Add(mBase, _consts[534]))
															v293 = v291 | v278
														}
														v296 = *(*int64)(unsafe.Add(mBase, _consts[572]))
														if v296&v214 == int64(0) {
															v306 = v293
														} else {
															v304 = *(*int64)(unsafe.Add(mBase, _consts[535]))
															v306 = v304 | v293
														}
														v311 = *(*int64)(unsafe.Add(mBase, _consts[573]))
														if v311&v214 == int64(0) {
															v321 = v306
														} else {
															v319 = *(*int64)(unsafe.Add(mBase, _consts[536]))
															v321 = v319 | v306
														}
														v324 = *(*int64)(unsafe.Add(mBase, _consts[574]))
														if v324&v214 == int64(0) {
															v334 = v321
														} else {
															v332 = *(*int64)(unsafe.Add(mBase, _consts[537]))
															v334 = v332 | v321
														}
														v339 = *(*int64)(unsafe.Add(mBase, _consts[575]))
														if v339&v214 == int64(0) {
															v349 = v334
														} else {
															v347 = *(*int64)(unsafe.Add(mBase, _consts[538]))
															v349 = v347 | v334
														}
														v352 = *(*int64)(unsafe.Add(mBase, _consts[576]))
														if v352&v214 == int64(0) {
															v362 = v349
														} else {
															v360 = *(*int64)(unsafe.Add(mBase, _consts[539]))
															v362 = v360 | v349
														}
														v367 = *(*int64)(unsafe.Add(mBase, _consts[577]))
														if v367&v214 == int64(0) {
															v377 = v362
														} else {
															v375 = *(*int64)(unsafe.Add(mBase, _consts[540]))
															v377 = v375 | v362
														}
														*(*uint32)(unsafe.Add(mBase, uint32(v212+v203))) = uint32(v377)
														v379 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
														v380 = v379
													}
													v382 = v195 + int32(1)
													if v382 < v380 {
														v195 = v382
														v199 = v380
														continue
													} else {
														break
													}
													break
												}
												v391 = v181
											}
											m.G0 = v12 + int32(2064)
											return v391
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(0)
								F_getKeysFreeResult(m, v12+int32(4))
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v391 = v161
									m.G0 = v12 + int32(2064)
									return v391
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
						v391 = v149
						m.G0 = v12 + int32(2064)
						return v391
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(0)
				v391 = int32(0)
				m.G0 = v12 + int32(2064)
				return v391
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(44)
			v391 = int32(0)
			m.G0 = v12 + int32(2064)
			return v391
		}
	}
}
func F_VM_GetContextFlags(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v184 int64
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int64
	_ = v193
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v205 int64
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int64
	_ = v237
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 float32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v2 {
		v77 = v2
	} else {
		v16 = l0 + int32(8)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v17 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+204))
			v23 = v19 << (uint(int32(9)) % 32) & int32(2097152)
			v27 = int32(1)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+200)))
			if v28&v27 != 0 {
				v35 = v27
				v38 = v35
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v17)+216))
				if v31 != 0 {
					v33 = F_isImportSlotMigrationJob(m, v31)
					mBase = m.M
					v35 = v33
					v38 = v35
				} else {
					v38 = int32(0)
				}
			}
			if v38 != 0 {
				v39 = v23 | int32(4096)
			} else {
				v39 = v23
			}
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+224)))
			if v43 == int32(3) {
				v46 = v39 | int32(4194304)
			} else {
				v46 = v39
			}
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+216))
			if v47 == int32(0) {
				v63 = v46
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
				if base.B2i32(v50 == int32(1)) == int32(0) {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+216))
					if v58 == int32(0) {
						v63 = v46
					} else {
						v63 = v46 | int32(67108864)
					}
				} else {
					v63 = v46 | int32(33554432)
				}
			}
		} else {
			v63 = int32(0)
		}
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v65 != 0 {
			v66 = v65
		} else {
			v66 = v16
		}
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
		if v67 == int32(0) {
			v77 = v63
		} else {
			v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+200)))
			if v72&int32(4128) != 0 {
				v75 = v63 | int32(524288)
			} else {
				v75 = v63
			}
			v77 = v75
		}
	}
	v80 = int32(0)
	v81 = *(*int32)(unsafe.Add(mBase, _consts[543]))
	v86 = v77 | base.B2i32(base.B2i32(v81 != v80) != int32(0))
	v90 = *(*int32)(unsafe.Add(mBase, _consts[544]))
	if v90 != 0 {
		v91 = v86 | int32(2)
	} else {
		v91 = v86
	}
	v95 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v95 != 0 {
		v96 = v91 | int32(32)
	} else {
		v96 = v91
	}
	v98 = *(*int32)(unsafe.Add(mBase, _consts[375]))
	if v98 == int32(0) {
		v106 = *(*int32)(unsafe.Add(mBase, _consts[131]))
		if v106 != 0 {
			v107 = v96 | int32(8192)
		} else {
			v107 = v96
		}
		v108 = v107
	} else {
		v108 = v96 | int32(_a0)
	}
	v109 = int32(_a44)
	v110 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v112 = *(*int64)(unsafe.Add(mBase, _consts[280]))
	if v112 == int64(0) {
		v127 = v108
	} else {
		if v110 == int32(0) {
			v122 = *(*int32)(unsafe.Add(mBase, _consts[97]))
			if v122 == int32(1792) {
				v125 = int32(256)
			} else {
				v125 = int32(768)
			}
			v127 = v125 | v108
		} else {
			v118 = *(*int32)(unsafe.Add(mBase, _consts[545]))
			if v118 != 0 {
				v127 = v108
			} else {
				v122 = *(*int32)(unsafe.Add(mBase, _consts[97]))
				if v122 == int32(1792) {
					v125 = int32(256)
				} else {
					v125 = int32(768)
				}
				v127 = v125 | v108
			}
		}
	}
	v131 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v131 != 0 {
		v132 = v127 | int32(64)
	} else {
		v132 = v127
	}
	v136 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	if int32(0) < v136 {
		v139 = v132 | int32(128)
	} else {
		v139 = v132
	}
	if v110 != 0 {
		v145 = *(*int32)(unsafe.Add(mBase, _consts[330]))
		if v145 != 0 {
			v146 = int32(24)
		} else {
			v146 = int32(8)
		}
		v147 = v146 | v139
		v149 = *(*int32)(unsafe.Add(mBase, _consts[193]))
		if base.Ui32(int32(1)) < base.Ui32(v149+int32(-1)) {
			switch v149 + int32(-13) {
			case 0:
				v162 = v147 | int32(65536)
			case 1:
				v162 = v147 | int32(131072)
			default:
				v162 = v147
			}
		} else {
			v162 = v147 | int32(32768)
		}
		if v149 == int32(14) {
			v167 = v162
		} else {
			v167 = v162 | int32(16384)
		}
		v169 = v167
	} else {
		v169 = v139 | int32(4)
	}
	v174 = v10 + int32(12)
	v179 = F_zmalloc_used_memory(m)
	mBase = m.M
	v184 = *(*int64)(unsafe.Add(mBase, _consts[280]))
	if v184 != int64(0) {
		v193 = base.I64_extend_i32_u(v179)
		if v174 != 0 {
			v197 = int32(_a44)
			v198 = *(*int64)(unsafe.Add(mBase, _consts[287]))
			v200 = *(*int32)(unsafe.Add(mBase, _consts[288]))
			if base.I64_extend_i32_u(v200) <= v198 {
				v215 = int32(0)
			} else {
				v205 = base.I64_div_s(v198, int64(16384))
				v212 = v200 - base.I32_wrap_i64(v198+v205*int64(44)) + int32(-44)
				if base.Ui32(v200) < base.Ui32(v212) {
					v214 = int32(0)
				} else {
					v214 = v212
				}
				v215 = v214
			}
			v217 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if v217 == int32(0) {
				v224 = v215
			} else {
				v221 = *(*int32)(unsafe.Add(mBase, _consts[34]))
				v222 = F_sdsAllocSize(m, v221)
				mBase = m.M
				v224 = v222 + v215
			}
			v225 = F_clusterIsAnySlotExporting(m)
			mBase = m.M
			if v225 == int32(0) {
				v230 = v224
			} else {
				v228 = F_clusterGetTotalSlotExportBufferMemory(m)
				mBase = m.M
				v230 = v228 + v224
			}
			v231 = int32(0)
			v233 = v179 - v230
			if base.Ui32(v179) < base.Ui32(v233) {
				v235 = v231
			} else {
				v235 = v233
			}
			v237 = *(*int64)(unsafe.Add(mBase, _consts[280]))
			if v174 == int32(0) {
			} else {
				*(*float32)(unsafe.Add(mBase, uint32(v174))) = base.F32_div(base.F32_convert_i32_u(v235), base.F32_convert_i64_u(v237))
			}
			if base.Ui64(v193) <= base.Ui64(v237) {
				v259 = v231
			} else {
				if base.Ui64(base.I64_extend_i32_u(v235)) <= base.Ui64(v237) {
					v259 = v231
				} else {
					v259 = int32(-1)
				}
			}
			v266 = v259
		} else {
			if base.Ui64(v184) < base.Ui64(v193) {
				v197 = int32(_a44)
				v198 = *(*int64)(unsafe.Add(mBase, _consts[287]))
				v200 = *(*int32)(unsafe.Add(mBase, _consts[288]))
				if base.I64_extend_i32_u(v200) <= v198 {
					v215 = int32(0)
				} else {
					v205 = base.I64_div_s(v198, int64(16384))
					v212 = v200 - base.I32_wrap_i64(v198+v205*int64(44)) + int32(-44)
					if base.Ui32(v200) < base.Ui32(v212) {
						v214 = int32(0)
					} else {
						v214 = v212
					}
					v215 = v214
				}
				v217 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				if v217 == int32(0) {
					v224 = v215
				} else {
					v221 = *(*int32)(unsafe.Add(mBase, _consts[34]))
					v222 = F_sdsAllocSize(m, v221)
					mBase = m.M
					v224 = v222 + v215
				}
				v225 = F_clusterIsAnySlotExporting(m)
				mBase = m.M
				if v225 == int32(0) {
					v230 = v224
				} else {
					v228 = F_clusterGetTotalSlotExportBufferMemory(m)
					mBase = m.M
					v230 = v228 + v224
				}
				v231 = int32(0)
				v233 = v179 - v230
				if base.Ui32(v179) < base.Ui32(v233) {
					v235 = v231
				} else {
					v235 = v233
				}
				v237 = *(*int64)(unsafe.Add(mBase, _consts[280]))
				if v174 == int32(0) {
				} else {
					*(*float32)(unsafe.Add(mBase, uint32(v174))) = base.F32_div(base.F32_convert_i32_u(v235), base.F32_convert_i64_u(v237))
				}
				if base.Ui64(v193) <= base.Ui64(v237) {
					v259 = v231
				} else {
					if base.Ui64(base.I64_extend_i32_u(v235)) <= base.Ui64(v237) {
						v259 = v231
					} else {
						v259 = int32(-1)
					}
				}
				v266 = v259
			} else {
				v266 = int32(0)
			}
		}
	} else {
		v187 = int32(0)
		if v174 == v187 {
			v259 = v187
			v266 = v259
		} else {
			v190 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v174))) = v190
			v266 = v190
		}
	}
	v267 = *(*float32)(unsafe.Add(mBase, uint32(v10)+12))
	v269 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	v272 = int32(_a44)
	v273 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	v275 = *(*int32)(unsafe.Add(mBase, _consts[525]))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+20))
	m.G0 = v10 + int32(16)
	if v266 == int32(-1) {
		v284 = v169 | int32(1024)
	} else {
		v284 = v169
	}
	if base.F32_gt(v267, float32(0.75)) != 0 {
		v289 = v284 | int32(2048)
	} else {
		v289 = v284
	}
	if v269 != int32(-1) {
		v292 = v289 | int32(262144)
	} else {
		v292 = v289
	}
	if v273 != 0 {
		v295 = v292 | int32(1048576)
	} else {
		v295 = v292
	}
	if v276 != 0 {
		v298 = v295 | int32(16777216)
	} else {
		v298 = v295
	}
	return v298
}
func F_VM_GetContextFlagsAll(m *base.Module) int32 {
	return int32(134217727)
}
func F_VM_GetCurrentCommandName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v2 = int32(0)
	if l0 == v2 {
		v14 = v2
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v6 == int32(0) {
			v14 = v2
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
			if v9 == int32(0) {
				v14 = v2
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+140))
				v14 = v12
			}
		}
	}
	return v14
}
func F_VM_GetDbIdFromDefragCtx(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	return v2
}
func F_VM_GetExpire(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v11 int32
	_ = v11
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4 != 0 {
		v7 = int64(-1)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		if v11&int32(1) == int32(0) {
			v22 = v7
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v4+(v11&int32(4)^int32(12)))))
			v22 = v21
		}
		if v22 == int64(-1) {
			v32 = v7
		} else {
			v26 = *(*int64)(unsafe.Add(mBase, _consts[98]))
			v27 = v22 - v26
			v28 = int64(0)
			if v28 < v27 {
				v31 = v27
			} else {
				v31 = v28
			}
			v32 = v31
		}
		return v32
	} else {
		return int64(-1)
	}
}
func F_VM_GetFunctionExecutionState(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	v2 = F_scriptInterrupt(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		if base.Ui32(v2+int32(-1)) < base.Ui32(int32(2)) {
			return base.B2i32(v2 != int32(2))
		} else {
			F__serverAssert(m, int32(_a954), int32(_a917), int32(14547))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_VM_GetKeyNameFromDigest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	return v2
}
func F_VM_GetKeyNameFromModuleKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	if l0 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		return v4
	} else {
		return int32(0)
	}
}
func F_VM_GetKeyNameFromOptCtx(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return v2
}
func F_VM_GetKeyspaceNotificationFlagsAll(m *base.Module) int32 {
	return int32(32767)
}
func F_VM_GetLRU(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(-1)
	v6 = int32(1)
	if l0 == int32(0) {
		v26 = v6
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v9 == int32(0) {
			v26 = v6
		} else {
			v12 = int32(0)
			v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[97])))
			if v14&int32(2) != 0 {
				v26 = v12
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v20 = F_lru_getIdleSecs(m, int32(base.Ui32(v17)>>(uint(int32(8))%32)))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = base.I64_extend_i32_u(v20 * int32(1000))
				v26 = v12
			}
		}
	}
	return v26
}
func F_VM_GetModuleOptionsAll(m *base.Module) int32 {
	return int32(63)
}
func F_VM_GetNotifyKeyspaceEvents(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _consts[289]))
	return v2
}
func F_VM_GetOpenKeyModesAll(m *base.Module) int32 {
	return int32(2031619)
}
func F_VM_GetToDbIdFromOptCtx(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	return v2
}
func F_VM_GetUsedMemoryRatio(m *base.Module) float32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int64
	_ = v21
	var v30 int64
	_ = v30
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v42 int64
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v104 float32
	_ = v104
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v11 = v5 + int32(12)
	v16 = F_zmalloc_used_memory(m)
	mBase = m.M
	v21 = *(*int64)(unsafe.Add(mBase, _consts[280]))
	if v21 != int64(0) {
		v30 = base.I64_extend_i32_u(v16)
		if v11 != 0 {
			v34 = int32(_a44)
			v35 = *(*int64)(unsafe.Add(mBase, _consts[287]))
			v37 = *(*int32)(unsafe.Add(mBase, _consts[288]))
			if base.I64_extend_i32_u(v37) <= v35 {
				v52 = int32(0)
			} else {
				v42 = base.I64_div_s(v35, int64(16384))
				v49 = v37 - base.I32_wrap_i64(v35+v42*int64(44)) + int32(-44)
				if base.Ui32(v37) < base.Ui32(v49) {
					v51 = int32(0)
				} else {
					v51 = v49
				}
				v52 = v51
			}
			v54 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if v54 == int32(0) {
				v61 = v52
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, _consts[34]))
				v59 = F_sdsAllocSize(m, v58)
				mBase = m.M
				v61 = v59 + v52
			}
			v62 = F_clusterIsAnySlotExporting(m)
			mBase = m.M
			if v62 == int32(0) {
				v67 = v61
			} else {
				v65 = F_clusterGetTotalSlotExportBufferMemory(m)
				mBase = m.M
				v67 = v65 + v61
			}
			v70 = v16 - v67
			if base.Ui32(v16) < base.Ui32(v70) {
				v72 = int32(0)
			} else {
				v72 = v70
			}
			v74 = *(*int64)(unsafe.Add(mBase, _consts[280]))
			if v11 == int32(0) {
			} else {
				*(*float32)(unsafe.Add(mBase, uint32(v11))) = base.F32_div(base.F32_convert_i32_u(v72), base.F32_convert_i64_u(v74))
			}
			if base.Ui64(v30) <= base.Ui64(v74) {
			} else {
				if base.Ui64(base.I64_extend_i32_u(v72)) <= base.Ui64(v74) {
				} else {
				}
			}
		} else {
			if base.Ui64(v21) < base.Ui64(v30) {
				v34 = int32(_a44)
				v35 = *(*int64)(unsafe.Add(mBase, _consts[287]))
				v37 = *(*int32)(unsafe.Add(mBase, _consts[288]))
				if base.I64_extend_i32_u(v37) <= v35 {
					v52 = int32(0)
				} else {
					v42 = base.I64_div_s(v35, int64(16384))
					v49 = v37 - base.I32_wrap_i64(v35+v42*int64(44)) + int32(-44)
					if base.Ui32(v37) < base.Ui32(v49) {
						v51 = int32(0)
					} else {
						v51 = v49
					}
					v52 = v51
				}
				v54 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				if v54 == int32(0) {
					v61 = v52
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, _consts[34]))
					v59 = F_sdsAllocSize(m, v58)
					mBase = m.M
					v61 = v59 + v52
				}
				v62 = F_clusterIsAnySlotExporting(m)
				mBase = m.M
				if v62 == int32(0) {
					v67 = v61
				} else {
					v65 = F_clusterGetTotalSlotExportBufferMemory(m)
					mBase = m.M
					v67 = v65 + v61
				}
				v70 = v16 - v67
				if base.Ui32(v16) < base.Ui32(v70) {
					v72 = int32(0)
				} else {
					v72 = v70
				}
				v74 = *(*int64)(unsafe.Add(mBase, _consts[280]))
				if v11 == int32(0) {
				} else {
					*(*float32)(unsafe.Add(mBase, uint32(v11))) = base.F32_div(base.F32_convert_i32_u(v72), base.F32_convert_i64_u(v74))
				}
				if base.Ui64(v30) <= base.Ui64(v74) {
				} else {
					if base.Ui64(base.I64_extend_i32_u(v72)) <= base.Ui64(v74) {
					} else {
					}
				}
			} else {
			}
		}
	} else {
		if v11 == int32(0) {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
		}
	}
	v104 = *(*float32)(unsafe.Add(mBase, uint32(v5)+12))
	m.G0 = v5 + int32(16)
	return v104
}
func F_VM_HashGet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v14 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v190
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l2
	v28 = l1 & int32(4)
	goto L6
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v18&int32(15) != int32(4) {
		v190 = int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	v190 = int32(0)
	goto L1
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v38 + int32(4)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v28 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v109 + int32(4)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if l1&int32(8) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L9:
	;
	if v42 == int32(0) {
		goto L5
	} else {
		goto L30
	}
L10:
	;
	if v42 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	if v42&int32(3) == int32(0) {
		v68 = v42
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v102 = F_createRawStringObject(m, v42, v101)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	v101 = v93 - v42
	goto L12
L14:
	;
	v72 = v68
	goto L22
L15:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v54 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v57 = v42
	goto L18
L17:
	;
	v101 = v42 - v42
	goto L12
L18:
	;
	v61 = v57 + int32(1)
	if v61&int32(3) == int32(0) {
		v68 = v61
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v66 != 0 {
		v57 = v61
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v93 = v61
	goto L13
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v81 = int32(-2139062144)
	if (int32(16843008)-v78|v78)&v81 == v81 {
		v72 = v72 + int32(4)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v87 = v72
	goto L25
L24:
	;
	goto L23
L25:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v91 != 0 {
		v87 = v87 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v93 = v87
	goto L13
L27:
	;
	goto L26
L28:
	;
	return int32(0)
L29:
	;
	v108 = v102
	goto L8
L30:
	;
	v108 = v42
	goto L8
L31:
	;
	if v28 == int32(0) {
		goto L6
	} else {
		goto L51
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(0)
	goto L31
L33:
	;
	if v113 == int32(0) {
		goto L32
	} else {
		goto L37
	}
L34:
	;
	if v113 == int32(0) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v119 = F_objectGetVal(m, v108)
	mBase = m.M
	v120 = F_hashTypeExists(m, v113, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v120
	goto L31
L37:
	;
	v125 = F_objectGetVal(m, v108)
	mBase = m.M
	v126 = F_hashTypeGetValueObject(m, v113, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L28
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v126
	if v126 == int32(0) {
		goto L31
	} else {
		goto L39
	}
L39:
	;
	v131 = F_getDecodedObject(m, v126)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L28
	} else {
		goto L40
	}
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	F_decrRefCount(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L28
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v131
	if v131 == int32(0) {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+28)))
	if v140&int32(1) == int32(0) {
		goto L31
	} else {
		goto L43
	}
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v139)+24))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
	if v145 == v146 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v167 = v164 + v163<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v131
	v169 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v163 + v169
	goto L31
L45:
	;
	v149 = int32(8)
	if v149 < v145 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
	v163 = v145
	v164 = v148
	goto L44
L47:
	;
	v152 = v145
	goto L49
L48:
	;
	v152 = v149
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+20)) = v152 << (uint(int32(1)) % 32)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
	v159 = F_valkey_realloc(m, v156, v152<<(uint(int32(4))%32))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L28
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+16)) = v159
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v139)+24))
	v163 = v162
	v164 = v159
	goto L44
L51:
	;
	F_decrRefCount(m, v108)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L28
	} else {
		goto L52
	}
L52:
	;
	goto L6
}
func F_VM_HashSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v255 int32
	_ = v255
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(28)
	if l0 == int32(0) {
		v234 = v17
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v255
L2:
	;
	goto L74
L3:
	;
	if l1&int32(-24) != 0 {
		v234 = v17
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v22 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = l2
	v67 = int32(2)
	v72 = l1 & int32(4)
	v88 = int32(0)
	goto L19
L6:
	;
	v39 = F_createHashObject(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v35&int32(2) != 0 {
		goto L5
	} else {
		goto L12
	}
L8:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v31&int32(2) != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v25&int32(15) == int32(4) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v234 = int32(138)
	goto L2
L11:
	;
	v234 = int32(8)
	goto L2
L12:
	;
	v234 = int32(8)
	goto L2
L13:
	;
	return int32(0)
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v39
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_dbAdd(m, v44, v45, v15+int32(4))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	switch v52&int32(15) + int32(-3) {
	case 0:
		goto L17
	default:
		goto L5
	case 3:
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	goto L5
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	goto L5
L18:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_dbUpdateObjectWithVolatileItemsTracking(m, v223, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L13
	} else {
		goto L71
	}
L19:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v92 = v90 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v72 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v160 + int32(4)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if l1&int32(3) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L22:
	;
	if v94 == int32(0) {
		goto L18
	} else {
		goto L42
	}
L23:
	;
	if v94 == int32(0) {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	if v94&int32(3) == int32(0) {
		v120 = v94
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v154 = F_createRawStringObject(m, v94, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L13
	} else {
		goto L41
	}
L26:
	;
	v153 = v145 - v94
	goto L25
L27:
	;
	v124 = v120
	goto L35
L28:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if v106 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v109 = v94
	goto L31
L30:
	;
	v153 = v94 - v94
	goto L25
L31:
	;
	v113 = v109 + int32(1)
	if v113&int32(3) == int32(0) {
		v120 = v113
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v118 != 0 {
		v109 = v113
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v145 = v113
	goto L26
L35:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v133 = int32(-2139062144)
	if (int32(16843008)-v130|v130)&v133 == v133 {
		v124 = v124 + int32(4)
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v139 = v124
	goto L38
L37:
	;
	goto L36
L38:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v143 != 0 {
		v139 = v139 + int32(1)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v145 = v139
	goto L26
L40:
	;
	goto L39
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v159 = v154
	v160 = v156
	goto L21
L42:
	;
	v159 = v94
	v160 = v92
	goto L21
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v159
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_hashTypeTryConversion(m, v199, v15+int32(4), int32(0), int32(1))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L13
	} else {
		goto L62
	}
L44:
	;
	F_decrRefCount(m, v159)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L13
	} else {
		goto L61
	}
L45:
	;
	if v72 == int32(0) {
		goto L19
	} else {
		goto L60
	}
L46:
	;
	if v164 != int32(1) {
		goto L43
	} else {
		goto L56
	}
L47:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v168 = F_objectGetVal(m, v159)
	mBase = m.M
	v169 = F_hashTypeExists(m, v167, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	if l1&v67 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v175 = int32(0)
	v177 = l1 & int32(1) & base.B2i32(v169 != v175)
	if v72 == v175 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v169 == int32(0) {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	if v177 != 0 {
		goto L19
	} else {
		goto L55
	}
L53:
	;
	if v177 != 0 {
		goto L44
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	goto L46
L56:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v184 = F_objectGetVal(m, v159)
	mBase = m.M
	v185 = F_hashTypeDelete(m, v183, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	v187 = v88 + v185
	if v72 == int32(0) {
		v88 = v187
		goto L19
	} else {
		goto L58
	}
L58:
	;
	F_decrRefCount(m, v159)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	v88 = v187
	goto L19
L60:
	;
	goto L44
L61:
	;
	goto L19
L62:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v207 = F_objectGetVal(m, v159)
	mBase = m.M
	v208 = F_objectGetVal(m, v164)
	mBase = m.M
	v211 = F_hashTypeSet(m, v206, v207, v208, int64(-1), int32(base.Ui32(v72)>>(uint(v67)%32)), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	if base.Ui32(l1) < base.Ui32(int32(16)) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v214 = v211
	goto L66
L65:
	;
	v214 = int32(1)
	goto L66
L66:
	;
	if v72 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v88 = v214 + v88
	goto L19
L68:
	;
	F_objectSetVal(m, v159, int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	F_decrRefCount(m, v159)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	v227 = F_moduleDelKeyIfEmpty(m, l0)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L13
	} else {
		goto L72
	}
L72:
	;
	if v88 != 0 {
		v255 = v88
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v234 = int32(44)
	goto L2
L74:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v234
	v255 = int32(0)
	goto L1
}
func F_VM_HashSetStringRef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v6 = int32(1)
	if l0 == int32(0) {
		v27 = v6
		return v27
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v9 == int32(0) {
			v27 = v6
			return v27
		} else {
			if l2 == int32(0) {
				v27 = v6
				return v27
			} else {
				if l1 == int32(0) {
					v27 = v6
					return v27
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					if v16&int32(15) != int32(4) {
						v27 = v6
						return v27
					} else {
						v21 = F_objectGetVal(m, l1)
						mBase = m.M
						v22 = F_hashTypeUpdateAsStringRef(m, v9, v21, l2, l3)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v27 = v22
							return v27
						}
					}
				}
			}
		}
	}
}
func F_VM_HoldString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v5&int32(-8) != int32(-16) {
		F_incrRefCount(m, l1)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			if l0 == int32(0) {
				v89 = l1
				return v89
			} else {
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
				if v55&int32(1) == int32(0) {
					v89 = l1
					return v89
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v60 == v61 {
						v64 = int32(8)
						if v64 < v60 {
							v67 = v60
						} else {
							v67 = v64
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v67 << (uint(int32(1)) % 32)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v74 = F_valkey_realloc(m, v71, v67<<(uint(int32(4))%32))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v74
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v78 = v77
							v79 = v74
							v82 = v79 + v78<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v82))) = l1
							v84 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v84
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v78 + v84
							v89 = l1
							return v89
						}
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v78 = v60
						v79 = v63
						v82 = v79 + v78<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v82))) = l1
						v84 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v84
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v78 + v84
						v89 = l1
						return v89
					}
				}
			}
		}
	} else {
		v10 = F_dupStringObject(m, l1)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if l0 == int32(0) {
				v89 = v10
				return v89
			} else {
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
				if v16&int32(1) == int32(0) {
					v89 = v10
					return v89
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v21 == v22 {
						v25 = int32(8)
						if v25 < v21 {
							v28 = v21
						} else {
							v28 = v25
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v28 << (uint(int32(1)) % 32)
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v35 = F_valkey_realloc(m, v32, v28<<(uint(int32(4))%32))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v35
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v39 = v38
							v40 = v35
							v43 = v40 + v39<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v43))) = v10
							v45 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v39 + v45
							return v10
						}
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v39 = v21
						v40 = v24
						v43 = v40 + v39<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v43))) = v10
						v45 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v45
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v39 + v45
						return v10
					}
				}
			}
		}
	}
}
func F_VM_InfoAddFieldLongLong(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v14 == int32(0) {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v27
			v32 = F_sdscatfmt(m, v13, int32(_a948), v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = v32
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34
				v38 = int32(0)
				m.G0 = v9 + int32(32)
				return v38
			}
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
			v22 = F_sdscatfmt(m, v13, int32(_a949), v9+int32(16))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v34 = v22
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34
				v38 = int32(0)
				m.G0 = v9 + int32(32)
				return v38
			}
		}
	} else {
		v38 = int32(1)
		m.G0 = v9 + int32(32)
		return v38
	}
}
func F_VM_IsBlockedReplyRequest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return int32(base.Ui32(v2)>>(uint(int32(2))%32)) & int32(1)
}
func F_VM_IsBlockedTimeoutRequest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return int32(base.Ui32(v2)>>(uint(int32(3))%32)) & int32(1)
}
func F_VM_IsChannelsPositionRequest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return int32(base.Ui32(v2)>>(uint(int32(8))%32)) & int32(1)
}
func F_VM_LatencyAddSample(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v12 int32
	_ = v12
	v5 = *(*int64)(unsafe.Add(mBase, _consts[44]))
	if v5 == int64(0) {
		return
	} else {
		if l1 < v5 {
			return
		} else {
			F_latencyAddSample(m, l0, l1*int64(1000))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_VM_ListPop(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	if l0 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v10 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(138)
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v13&int32(15) == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
				if v23&int32(2) != 0 {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v31 == int32(0) {
						v39 = v10
						v42 = F_listTypePop(m, v39, base.B2i32(l1 != int32(0)))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = F_getDecodedObject(m, v42)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								F_decrRefCount(m, v42)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v48 = F_moduleDelKeyIfEmpty(m, l0)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										if v48 != 0 {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+28)))
											if v56&int32(1) == int32(0) {
												return v44
											} else {
												v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
												if v61 == v62 {
													v65 = int32(8)
													if v65 < v61 {
														v68 = v61
													} else {
														v68 = v65
													}
													*(*int32)(unsafe.Add(mBase, uint32(v55)+20)) = v68 << (uint(int32(1)) % 32)
													v72 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
													v75 = F_valkey_realloc(m, v72, v68<<(uint(int32(4))%32))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = v75
														v78 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
														v79 = v78
														v80 = v75
														v83 = v80 + v79<<(uint(int32(3))%32)
														*(*int32)(unsafe.Add(mBase, uint32(v83))) = v44
														v85 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v85
														*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v79 + v85
														return v44
													}
												} else {
													v64 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
													v79 = v61
													v80 = v64
													v83 = v80 + v79<<(uint(int32(3))%32)
													*(*int32)(unsafe.Add(mBase, uint32(v83))) = v44
													v85 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v85
													*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v79 + v85
													return v44
												}
											}
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											F_listTypeTryConversion(m, v50, int32(2), int32(565), l0)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+28)))
												if v56&int32(1) == int32(0) {
													return v44
												} else {
													v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
													v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
													if v61 == v62 {
														v65 = int32(8)
														if v65 < v61 {
															v68 = v61
														} else {
															v68 = v65
														}
														*(*int32)(unsafe.Add(mBase, uint32(v55)+20)) = v68 << (uint(int32(1)) % 32)
														v72 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
														v75 = F_valkey_realloc(m, v72, v68<<(uint(int32(4))%32))
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = v75
															v78 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
															v79 = v78
															v80 = v75
															v83 = v80 + v79<<(uint(int32(3))%32)
															*(*int32)(unsafe.Add(mBase, uint32(v83))) = v44
															v85 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v85
															*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v79 + v85
															return v44
														}
													} else {
														v64 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
														v79 = v61
														v80 = v64
														v83 = v80 + v79<<(uint(int32(3))%32)
														*(*int32)(unsafe.Add(mBase, uint32(v83))) = v44
														v85 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v85
														*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v79 + v85
														return v44
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_moduleFreeKeyIterator(m, l0)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v39 = v38
							v42 = F_listTypePop(m, v39, base.B2i32(l1 != int32(0)))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = F_getDecodedObject(m, v42)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									F_decrRefCount(m, v42)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										v48 = F_moduleDelKeyIfEmpty(m, l0)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											if v48 != 0 {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+28)))
												if v56&int32(1) == int32(0) {
													return v44
												} else {
													v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
													v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
													if v61 == v62 {
														v65 = int32(8)
														if v65 < v61 {
															v68 = v61
														} else {
															v68 = v65
														}
														*(*int32)(unsafe.Add(mBase, uint32(v55)+20)) = v68 << (uint(int32(1)) % 32)
														v72 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
														v75 = F_valkey_realloc(m, v72, v68<<(uint(int32(4))%32))
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = v75
															v78 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
															v79 = v78
															v80 = v75
															v83 = v80 + v79<<(uint(int32(3))%32)
															*(*int32)(unsafe.Add(mBase, uint32(v83))) = v44
															v85 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v85
															*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v79 + v85
															return v44
														}
													} else {
														v64 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
														v79 = v61
														v80 = v64
														v83 = v80 + v79<<(uint(int32(3))%32)
														*(*int32)(unsafe.Add(mBase, uint32(v83))) = v44
														v85 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v85
														*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v79 + v85
														return v44
													}
												}
											} else {
												v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												F_listTypeTryConversion(m, v50, int32(2), int32(565), l0)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return int32(0)
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+28)))
													if v56&int32(1) == int32(0) {
														return v44
													} else {
														v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
														v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
														if v61 == v62 {
															v65 = int32(8)
															if v65 < v61 {
																v68 = v61
															} else {
																v68 = v65
															}
															*(*int32)(unsafe.Add(mBase, uint32(v55)+20)) = v68 << (uint(int32(1)) % 32)
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
															v75 = F_valkey_realloc(m, v72, v68<<(uint(int32(4))%32))
															mBase = m.M
															v76 = m.ExcPending
															if v76 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = v75
																v78 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
																v79 = v78
																v80 = v75
																v83 = v80 + v79<<(uint(int32(3))%32)
																*(*int32)(unsafe.Add(mBase, uint32(v83))) = v44
																v85 = int32(1)
																*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v85
																*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v79 + v85
																return v44
															}
														} else {
															v64 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
															v79 = v61
															v80 = v64
															v83 = v80 + v79<<(uint(int32(3))%32)
															*(*int32)(unsafe.Add(mBase, uint32(v83))) = v44
															v85 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v85
															*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v79 + v85
															return v44
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
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(8)
					return int32(0)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(138)
				return int32(0)
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
		return int32(0)
	}
}
func F_VM_ListPush(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
		v101 = int32(1)
		m.G0 = v8 + int32(16)
		return v101
	} else {
		if l2 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v17 == int32(0) {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
				if v29&int32(2) == int32(0) {
					*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(8)
					v101 = int32(1)
					m.G0 = v8 + int32(16)
					return v101
				} else {
					v43 = l0 + int32(20)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v44 == int32(0) {
						v52 = v17
						if v52 != 0 {
							v81 = v52
							v82 = int32(0)
							F_listTypeTryConversionAppend(m, v81, v8+int32(8), v82, v82, int32(565), l0)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
								F_listTypePush(m, v90, v91, base.B2i32(l1 != int32(0)))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									v101 = v82
									m.G0 = v8 + int32(16)
									return v101
								}
							}
						} else {
							v53 = int32(0)
							v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
							if v54&int32(2) == v53 {
								v81 = v53
								v82 = int32(0)
								F_listTypeTryConversionAppend(m, v81, v8+int32(8), v82, v82, int32(565), l0)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
									F_listTypePush(m, v90, v91, base.B2i32(l1 != int32(0)))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										v101 = v82
										m.G0 = v8 + int32(16)
										return v101
									}
								}
							} else {
								v59 = F_createListListpackObject(m)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v59
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									F_dbAdd(m, v62, v63, v8+int32(12))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v68
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
										switch v70&int32(15) + int32(-3) {
										case 0:
											*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
											v81 = v68
										default:
											v81 = v68
										case 3:
											*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
											v81 = v68
										}
										v82 = int32(0)
										F_listTypeTryConversionAppend(m, v81, v8+int32(8), v82, v82, int32(565), l0)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
											F_listTypePush(m, v90, v91, base.B2i32(l1 != int32(0)))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												v101 = v82
												m.G0 = v8 + int32(16)
												return v101
											}
										}
									}
								}
							}
						}
					} else {
						F_moduleFreeKeyIterator(m, l0)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v52 = v51
							if v52 != 0 {
								v81 = v52
								v82 = int32(0)
								F_listTypeTryConversionAppend(m, v81, v8+int32(8), v82, v82, int32(565), l0)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
									F_listTypePush(m, v90, v91, base.B2i32(l1 != int32(0)))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										v101 = v82
										m.G0 = v8 + int32(16)
										return v101
									}
								}
							} else {
								v53 = int32(0)
								v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
								if v54&int32(2) == v53 {
									v81 = v53
									v82 = int32(0)
									F_listTypeTryConversionAppend(m, v81, v8+int32(8), v82, v82, int32(565), l0)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
										F_listTypePush(m, v90, v91, base.B2i32(l1 != int32(0)))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											v101 = v82
											m.G0 = v8 + int32(16)
											return v101
										}
									}
								} else {
									v59 = F_createListListpackObject(m)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v59
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_dbAdd(m, v62, v63, v8+int32(12))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v68
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
											switch v70&int32(15) + int32(-3) {
											case 0:
												*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
												v81 = v68
											default:
												v81 = v68
											case 3:
												*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
												v81 = v68
											}
											v82 = int32(0)
											F_listTypeTryConversionAppend(m, v81, v8+int32(8), v82, v82, int32(565), l0)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
												F_listTypePush(m, v90, v91, base.B2i32(l1 != int32(0)))
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return int32(0)
												} else {
													v101 = v82
													m.G0 = v8 + int32(16)
													return v101
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v20 = int32(1)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				if v21&int32(15) == v20 {
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
					if v36&int32(2) == int32(0) {
						*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(8)
						v101 = int32(1)
						m.G0 = v8 + int32(16)
						return v101
					} else {
						v43 = l0 + int32(20)
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v44 == int32(0) {
							v52 = v17
							if v52 != 0 {
								v81 = v52
								v82 = int32(0)
								F_listTypeTryConversionAppend(m, v81, v8+int32(8), v82, v82, int32(565), l0)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
									F_listTypePush(m, v90, v91, base.B2i32(l1 != int32(0)))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										v101 = v82
										m.G0 = v8 + int32(16)
										return v101
									}
								}
							} else {
								v53 = int32(0)
								v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
								if v54&int32(2) == v53 {
									v81 = v53
									v82 = int32(0)
									F_listTypeTryConversionAppend(m, v81, v8+int32(8), v82, v82, int32(565), l0)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
										F_listTypePush(m, v90, v91, base.B2i32(l1 != int32(0)))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											v101 = v82
											m.G0 = v8 + int32(16)
											return v101
										}
									}
								} else {
									v59 = F_createListListpackObject(m)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v59
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_dbAdd(m, v62, v63, v8+int32(12))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v68
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
											switch v70&int32(15) + int32(-3) {
											case 0:
												*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
												v81 = v68
											default:
												v81 = v68
											case 3:
												*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
												v81 = v68
											}
											v82 = int32(0)
											F_listTypeTryConversionAppend(m, v81, v8+int32(8), v82, v82, int32(565), l0)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
												F_listTypePush(m, v90, v91, base.B2i32(l1 != int32(0)))
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return int32(0)
												} else {
													v101 = v82
													m.G0 = v8 + int32(16)
													return v101
												}
											}
										}
									}
								}
							}
						} else {
							F_moduleFreeKeyIterator(m, l0)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v52 = v51
								if v52 != 0 {
									v81 = v52
									v82 = int32(0)
									F_listTypeTryConversionAppend(m, v81, v8+int32(8), v82, v82, int32(565), l0)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
										F_listTypePush(m, v90, v91, base.B2i32(l1 != int32(0)))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											v101 = v82
											m.G0 = v8 + int32(16)
											return v101
										}
									}
								} else {
									v53 = int32(0)
									v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
									if v54&int32(2) == v53 {
										v81 = v53
										v82 = int32(0)
										F_listTypeTryConversionAppend(m, v81, v8+int32(8), v82, v82, int32(565), l0)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
											F_listTypePush(m, v90, v91, base.B2i32(l1 != int32(0)))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												v101 = v82
												m.G0 = v8 + int32(16)
												return v101
											}
										}
									} else {
										v59 = F_createListListpackObject(m)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v59
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											F_dbAdd(m, v62, v63, v8+int32(12))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v68
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
												switch v70&int32(15) + int32(-3) {
												case 0:
													*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
													v81 = v68
												default:
													v81 = v68
												case 3:
													*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
													v81 = v68
												}
												v82 = int32(0)
												F_listTypeTryConversionAppend(m, v81, v8+int32(8), v82, v82, int32(565), l0)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
													F_listTypePush(m, v90, v91, base.B2i32(l1 != int32(0)))
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return int32(0)
													} else {
														v101 = v82
														m.G0 = v8 + int32(16)
														return v101
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(138)
					v101 = v20
					m.G0 = v8 + int32(16)
					return v101
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
			v101 = int32(1)
			m.G0 = v8 + int32(16)
			return v101
		}
	}
}
func F_VM_ListSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	if l2 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v15 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(138)
			v51 = int32(1)
			m.G0 = v8 + int32(16)
			return v51
		} else {
			v18 = int32(1)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v19&int32(15) == v18 {
				v31 = int32(0)
				F_listTypeTryConversionAppend(m, v15, v8+int32(12), v31, v31, int32(565), l0)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v39 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if v39 == int32(0) {
							v51 = v18
							m.G0 = v8 + int32(16)
							return v51
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
							F_listTypeReplace(m, l0+int32(24), v45)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_moduleFreeKeyIterator(m, l0)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									v51 = int32(0)
									m.G0 = v8 + int32(16)
									return v51
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(138)
				v51 = int32(1)
				m.G0 = v8 + int32(16)
				return v51
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
		v51 = int32(1)
		m.G0 = v8 + int32(16)
		return v51
	}
}
func F_VM_LoadConfigs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	v3 = int32(1)
	if l0 == int32(0) {
		v17 = v3
		return v17
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v6 == int32(0) {
			v17 = v3
			return v17
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+72))
			if v9 == int32(0) {
				v17 = v3
				return v17
			} else {
				v12 = F_loadModuleConfigs(m, v6)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					v17 = v12
					return v17
				}
			}
		}
	}
}
func F_VM_LoadDataTypeFromStringEncver(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	v11 = F_objectGetVal(m, l0)
	mBase = m.M
	v14 = F___memcpy(m, v7+int32(32), int32(_a209), int32(80))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+20)) = int64(-4294967296)
	v20 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v7)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v7 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v31 = m.T0[v30].(func(*base.Module, int32, int32) int32)(m, v7, l2)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
		if v35 == int32(0) {
			m.G0 = v7 + int32(112)
			return v31
		} else {
			F_moduleFreeContext(m, v35)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
				F_valkey_free(m, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(112)
					return v31
				}
			}
		}
	}
}
func F_VM_LoadFloat(m *base.Module, l0 int32) float32 {
	mBase := m.M
	_ = mBase
	var v3 float32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 float32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v61 float32
	_ = v61
	v3 = float32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v12 != 0 {
		v61 = v3
		m.G0 = v9 + int32(32)
		return v61
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v15 = F_rdbLoadLen(m, v13, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return float32(0)
		} else {
			if v15 != int64(3) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+48)))
				if v31&int32(1) != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
					v61 = v3
					m.G0 = v9 + int32(32)
					return v61
				} else {
					v34 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v36 != 0 {
						v39 = *(*int32)(unsafe.Add(mBase, _consts[209]))
						if v39 == int32(0) {
							v43 = F_objectGetVal(m, v36)
							mBase = m.M
							v44 = v43
						} else {
							v44 = int32(_a938)
						}
					} else {
						v44 = int32(_a939)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v44
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 + int32(84)
					F__serverPanic_1(m, int32(_a917), int32(7585), int32(_a940), v9)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return float32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v24 = F_rdbLoadBinaryFloatValue(m, v21, v9+int32(28))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return float32(0)
				} else {
					if v24 == int32(-1) {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+48)))
						if v31&int32(1) != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
							v61 = v3
							m.G0 = v9 + int32(32)
							return v61
						} else {
							v34 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v36 != 0 {
								v39 = *(*int32)(unsafe.Add(mBase, _consts[209]))
								if v39 == int32(0) {
									v43 = F_objectGetVal(m, v36)
									mBase = m.M
									v44 = v43
								} else {
									v44 = int32(_a938)
								}
							} else {
								v44 = int32(_a939)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v44
							*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 + int32(84)
							F__serverPanic_1(m, int32(_a917), int32(7585), int32(_a940), v9)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return float32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					} else {
						v28 = *(*float32)(unsafe.Add(mBase, uint32(v9)+28))
						v61 = v28
						m.G0 = v9 + int32(32)
						return v61
					}
				}
			}
		}
	}
}
func F_VM_LoadLongDouble(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	v4 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v12 != 0 {
		v32 = v4
		v33 = v4
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v32
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v33
		m.G0 = v8 + int32(32)
		return
	} else {
		v16 = F_moduleLoadString(m, l1, int32(1), v8+int32(12))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				v23 = F_string2ld(m, v16, v20, v8+int32(16))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					F_valkey_free(m, v16)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v29 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(24))))
						v30 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
						v32 = v30
						v33 = v29
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = v32
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v33
						m.G0 = v8 + int32(32)
						return
					}
				}
			} else {
				v18 = int64(0)
				v32 = v18
				v33 = v18
				*(*int64)(unsafe.Add(mBase, uint32(l0))) = v32
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v33
				m.G0 = v8 + int32(32)
				return
			}
		}
	}
}
func F_VM_Log(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
	if l0 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v13 = v12
	} else {
		v13 = int32(0)
	}
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	F_moduleLogRaw(m, v13, l1, l2, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
func F_VM_MallocSize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-8))))
	return v4&int32(2147483647) + int32(8)
}
func F_VM_MallocSizeDict(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+24))
	return v3 + int32(4)
}
func F_VM_ModuleTypeGetType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v2 = int32(0)
	if l0 == v2 {
		v17 = v2
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v6 == int32(0) {
			v17 = v2
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			if v9&int32(15) != int32(5) {
				v17 = v2
			} else {
				v14 = F_objectGetVal(m, v6)
				mBase = m.M
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v17 = v15
			}
		}
	}
	return v17
}
func F_VM_MonotonicMicroseconds(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v3 = m.T0[v2].(func(*base.Module) int64)(m)
	mBase = m.M
	return v3
}
func F_VM_NotifyKeyspaceEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v6 = int32(1)
	if l0 == int32(0) {
		v20 = v6
		return v20
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v9 == int32(0) {
			v20 = v6
			return v20
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
			F_notifyKeyspaceEvent(m, l1, l2, l3, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v20 = int32(0)
				return v20
			}
		}
	}
}
func F_VM_OpenKey(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
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
	v7 = int32(16)
	v22 = int32(base.Ui32(l2)>>(uint(v7)%32))&int32(7) | int32(base.Ui32(l2)>>(uint(int32(15))%32))&v7 | l2<<(uint(int32(11))%32)>>(uint(int32(31))%32)&int32(23)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+96))
	if l2&int32(2) == int32(0) {
		v33 = F_lookupKeyReadWithFlags(m, v24, l1, v22)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			if v33 != 0 {
				v37 = v33
				v39 = F_valkey_malloc(m, int32(88))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v39))) = l0
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v43
					F_incrRefCount(m, l1)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = l2
						v49 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v49
						*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v37
						if v37 == v49 {
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
							switch v54&int32(15) + int32(-3) {
							case 0:
								*(*int64)(unsafe.Add(mBase, uint32(v39)+80)) = int64(4294967296)
								*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = int32(0)
							default:
							case 3:
								*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = int32(0)
							}
						}
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v65&int32(1) == int32(0) {
							return v39
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v70 == v71 {
								v74 = int32(8)
								if v74 < v70 {
									v77 = v70
								} else {
									v77 = v74
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v77 << (uint(int32(1)) % 32)
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v84 = F_valkey_realloc(m, v81, v77<<(uint(int32(4))%32))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v84
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v88 = v84
									v89 = v87
									v92 = v88 + v89<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v92))) = v39
									*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v89 + int32(1)
									return v39
								}
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v88 = v73
								v89 = v70
								v92 = v88 + v89<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v92))) = v39
								*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v89 + int32(1)
								return v39
							}
						}
					}
				}
			} else {
				return int32(0)
			}
		}
	} else {
		v29 = F_lookupKeyWriteWithFlags(m, v24, l1, v22)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v37 = v29
			v39 = F_valkey_malloc(m, int32(88))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v39))) = l0
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v43
				F_incrRefCount(m, l1)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = l2
					v49 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v49
					*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v37
					if v37 == v49 {
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
						switch v54&int32(15) + int32(-3) {
						case 0:
							*(*int64)(unsafe.Add(mBase, uint32(v39)+80)) = int64(4294967296)
							*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = int32(0)
						default:
						case 3:
							*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = int32(0)
						}
					}
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v65&int32(1) == int32(0) {
						return v39
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v70 == v71 {
							v74 = int32(8)
							if v74 < v70 {
								v77 = v70
							} else {
								v77 = v74
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v77 << (uint(int32(1)) % 32)
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v84 = F_valkey_realloc(m, v81, v77<<(uint(int32(4))%32))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v84
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v88 = v84
								v89 = v87
								v92 = v88 + v89<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v92))) = v39
								*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v89 + int32(1)
								return v39
							}
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v88 = v73
							v89 = v70
							v92 = v88 + v89<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v92))) = v39
							*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v89 + int32(1)
							return v39
						}
					}
				}
			}
		}
	}
}
func F_VM_PoolAlloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return int32(0)
L3:
	;
	if base.Ui32(v14) < base.Ui32(l1) {
		v46 = v14
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v14 = v11 - v12
	goto L3
L5:
	;
	v14 = int32(0)
	goto L3
L6:
	;
	if base.Ui32(v46) < base.Ui32(l1) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v21 = int32(4)
	goto L9
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v31 = v28 & (v21 + int32(-1))
	if v31 == int32(0) {
		v37 = v28
		goto L13
	} else {
		goto L14
	}
L9:
	;
	if base.Ui32(v21) <= base.Ui32(l1) {
		goto L8
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v25 = int32(base.Ui32(v21) >> (uint(int32(1)) % 32))
	if base.Ui32(l1) <= base.Ui32(v25) {
		v21 = v25
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v40 = v39 - v37
	if base.Ui32(v39) < base.Ui32(v40) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v35 = v28 + v21 - v31
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v35
	v37 = v35
	goto L13
L15:
	;
	v42 = int32(0)
	goto L17
L16:
	;
	v42 = v40
	goto L17
L17:
	;
	v46 = v42
	goto L6
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v69 + l1
	return v68 + v69 + int32(12)
L19:
	;
	v52 = int32(8192)
	if base.Ui32(v52) < base.Ui32(l1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v68 = v9
	v69 = v50
	goto L18
L21:
	;
	v55 = l1
	goto L23
L22:
	;
	v55 = v52
	goto L23
L23:
	;
	v58 = F_valkey_malloc(m, v55+int32(12))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v55
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v58
	v68 = v58
	v69 = int32(0)
	goto L18
}
func F_VM_PublishMessageShard(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_pubsubPublishMessageAndPropagateToCluster(m, l1, l2, int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_VM_RdbSave(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	v6 = int32(1)
	v7 = int32(28)
	if l1 == int32(0) {
		v21 = v6
		v22 = v7
		*(*int32)(unsafe.Add(mBase, _consts[5])) = v22
		v25 = v21
		return v25
	} else {
		if l2 != 0 {
			v21 = v6
			v22 = v7
			*(*int32)(unsafe.Add(mBase, _consts[5])) = v22
			v25 = v21
			return v25
		} else {
			v10 = int32(1)
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v11 != v10 {
				F__serverAssert(m, int32(_a953), int32(_a917), int32(14466))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v15 = F_rdbSaveToFile(m, v14)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v15 != 0 {
						v25 = v10
					} else {
						v19 = int32(0)
						v21 = v19
						v22 = v19
						*(*int32)(unsafe.Add(mBase, _consts[5])) = v22
						v25 = v21
					}
					return v25
				}
			}
		}
	}
}
func F_VM_RdbStreamFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v2 != int32(1) {
		F__serverAssert(m, int32(_a195), int32(_a917), int32(14371))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_valkey_free(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_valkey_free(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_VM_RegisterAuthCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v13 int32
	_ = v13
	v5 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = v7
		v11 = *(*int32)(unsafe.Add(mBase, _consts[559]))
		v12 = F_listAddNodeHead(m, v11, v5)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	}
}
func F_VM_RegisterDefragFunc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+60)) = l1
	return int32(0)
}
func F_VM_RegisterNumericConfig(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int64, l5 int64, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = F_moduleConfigValidityCheck(m, v14, l1, l3, v13)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 != 0 {
			v49 = v13
			return v49
		} else {
			v21 = F_valkey_malloc(m, int32(24))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_sdsnew(m, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l8
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v23
					*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l9
					*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l6
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
					v32 = F_listAddNodeTail(m, v31, v21)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
						F_addModuleNumericConfig(m, v34, l1, l3&int32(113), v21, l2, int32(base.Ui32(l3)>>(uint(int32(7))%32))&int32(1)|int32(base.Ui32(l3)>>(uint(int32(6))%32))&int32(8), l4, l5)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v49 = int32(0)
							return v49
						}
					}
				}
			}
		}
	}
}
func F_VM_RegisterUnsignedNumericConfig(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int64, l5 int64, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = F_moduleConfigValidityCheck(m, v14, l1, l3, v13)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 != 0 {
			v49 = v13
			return v49
		} else {
			v21 = F_valkey_malloc(m, int32(24))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_sdsnew(m, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l8
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v23
					*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l9
					*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l6
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
					v32 = F_listAddNodeTail(m, v31, v21)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
						F_addModuleUnsignedNumericConfig(m, v34, l1, l3&int32(113), v21, l2, int32(base.Ui32(l3)>>(uint(int32(7))%32))&int32(1)|int32(base.Ui32(l3)>>(uint(int32(6))%32))&int32(8), l4, l5)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v49 = int32(0)
							return v49
						}
					}
				}
			}
		}
	}
}
func F_VM_Replicate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int64
	_ = v214
	var v222 int32
	_ = v222
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v5
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 == v5 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v222
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l3
	v142 = F_moduleCreateArgvFromUserFormat(m, l1, l2, v12+int32(12), v12+int32(8), l3)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L20
	} else {
		goto L40
	}
L3:
	;
	if v70 != 0 {
		goto L36
	} else {
		goto L37
	}
L4:
	;
	v80 = int32(0)
	v83 = m.G0
	v85 = v83 - int32(16)
	m.G0 = v85
	v89 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v89 == v80 {
		v121 = v80
		goto L24
	} else {
		goto L25
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v23 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	v32 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v32 == v23 {
		v64 = v23
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v70 = v22 & int32(16)
	if v70 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	m.G0 = v28 + int32(16)
	goto L6
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	if v36 == int32(0) {
		v64 = v23
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[212])))
	v41 = v28 + int32(8)
	F_listRewind(m, v39, v41)
	mBase = m.M
	v43 = int32(0)
	v46 = F_listNext(m, v41)
	mBase = m.M
	if v46 == v43 {
		v64 = v43
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v51 = v46
	goto L11
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v53 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v64 = v43
	goto L7
L13:
	;
	v62 = F_listNext(m, v28+int32(8))
	mBase = m.M
	if v62 != 0 {
		v51 = v62
		goto L11
	} else {
		goto L16
	}
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+156))
	if base.Ui32(v54+int32(-18)) <= base.Ui32(int32(2)) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v64 = int32(1)
	goto L7
L16:
	;
	goto L12
L17:
	;
	v74 = F_lookupCommandByCString(m, l1)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v64 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v133 = v5
	v134 = int32(0)
	goto L2
L20:
	;
	return int32(0)
L21:
	;
	if v74 == int32(0) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v133 = v74
	v134 = v64
	goto L2
L23:
	;
	v126 = F_lookupCommandByCString(m, l1)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L20
	} else {
		goto L34
	}
L24:
	;
	m.G0 = v85 + int32(16)
	goto L23
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	if v93 == int32(0) {
		v121 = v80
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_consts[212])))
	v98 = v85 + int32(8)
	F_listRewind(m, v96, v98)
	mBase = m.M
	v100 = int32(0)
	v103 = F_listNext(m, v98)
	mBase = m.M
	if v103 == v100 {
		v121 = v100
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v108 = v103
	goto L28
L28:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v110 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v121 = v100
	goto L24
L30:
	;
	v119 = F_listNext(m, v85+int32(8))
	mBase = m.M
	if v119 != 0 {
		v108 = v119
		goto L28
	} else {
		goto L33
	}
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+156))
	if base.Ui32(v111+int32(-18)) <= base.Ui32(int32(2)) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v121 = int32(1)
	goto L24
L33:
	;
	goto L29
L34:
	;
	if v126 != 0 {
		v133 = v126
		v134 = v121
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v222 = int32(1)
	goto L1
L36:
	;
	F_clusterFailAllSlotExportsWithMessage(m, int32(_a922))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L20
	} else {
		goto L38
	}
L37:
	;
	v222 = int32(1)
	goto L1
L38:
	;
	v133 = v5
	v134 = v64
	goto L2
L39:
	;
	v145 = int32(0)
	if base.B2i32(v133 != v145)&v134 != int32(1) {
		v165 = int32(-1)
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if v142 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v222 = int32(1)
	goto L1
L42:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+96))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+28))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	F_alsoPropagate(m, v168, v142, v169, int32(base.Ui32(v170^int32(-1))>>(uint(int32(1))%32))&int32(3), v165)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L20
	} else {
		goto L48
	}
L43:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v153 = F_clusterSlotByCommand(m, v133, v142, v152, v12)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L20
	} else {
		goto L44
	}
L44:
	;
	if v153 != int32(-1) {
		v165 = v153
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+2)))
	if v157&int32(16) == int32(0) {
		v165 = v153
		goto L42
	} else {
		goto L46
	}
L46:
	;
	F_clusterFailAllSlotExportsWithMessage(m, int32(_a923))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L20
	} else {
		goto L47
	}
L47:
	;
	v165 = v153
	goto L42
L48:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v179 < int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	F_valkey_free(m, v142)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L20
	} else {
		goto L55
	}
L50:
	;
	v185 = v145
	goto L51
L51:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v142+v185<<(uint(int32(2))%32))))
	F_decrRefCount(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L20
	} else {
		goto L53
	}
L52:
	;
	goto L49
L53:
	;
	v198 = v185 + int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v198 < v199 {
		v185 = v198
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v212 = int32(_a44)
	v214 = *(*int64)(unsafe.Add(mBase, _consts[83]))
	*(*int64)(unsafe.Add(mBase, _consts[83])) = v214 + int64(1)
	v222 = int32(0)
	goto L1
}
func F_VM_ReplicateVerbatim(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+96))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v2)+292))
	F_alsoPropagate(m, v4, v5, v6, int32(3), v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(_a44)
		v15 = *(*int64)(unsafe.Add(mBase, _consts[83]))
		*(*int64)(unsafe.Add(mBase, _consts[83])) = v15 + int64(1)
		return int32(0)
	}
}
func F_VM_ReplySetArrayLength(m *base.Module, l0 int32, l1 int32) {
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
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v6 = m.G0
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v10&v7 == int32(0) {
		v22 = l0 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		if v23 == int32(0) {
			m.G0 = v8 + int32(16)
			return
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v26 != 0 {
				v39 = v26 + int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v39
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v39<<(uint(int32(2))%32))))
				F_setDeferredArrayLen(m, v23, v45, l1)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v48 != 0 {
						m.G0 = v8 + int32(16)
						return
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						F_valkey_free(m, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(3) < v28 {
					m.G0 = v8 + int32(16)
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v32
					F__serverLog(m, int32(3), int32(_a921), v8)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v15 == int32(0) {
			m.G0 = v8 + int32(16)
			return
		} else {
			v22 = v15 + int32(36)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			if v23 == int32(0) {
				m.G0 = v8 + int32(16)
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v26 != 0 {
					v39 = v26 + int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v39<<(uint(int32(2))%32))))
					F_setDeferredArrayLen(m, v23, v45, l1)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v48 != 0 {
							m.G0 = v8 + int32(16)
							return
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							F_valkey_free(m, v49)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					if int32(3) < v28 {
						m.G0 = v8 + int32(16)
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v32
						F__serverLog(m, int32(3), int32(_a921), v8)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_VM_ReplySetAttributeLength(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+224)))
	if v11 == int32(2) {
		m.G0 = v8 + int32(16)
		return
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
		if v14&int32(16) == int32(0) {
			v25 = v10
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v26 != 0 {
				v39 = v26 + int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v39
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v39<<(uint(int32(2))%32))))
				F_setDeferredAttributeLen(m, v25, v45, l1)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v48 != 0 {
						m.G0 = v8 + int32(16)
						return
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						F_valkey_free(m, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(3) < v28 {
					m.G0 = v8 + int32(16)
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v32
					F__serverLog(m, int32(3), int32(_a921), v8)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v19 == int32(0) {
				m.G0 = v8 + int32(16)
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
				if v22 == int32(0) {
					m.G0 = v8 + int32(16)
					return
				} else {
					v25 = v22
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v26 != 0 {
						v39 = v26 + int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v39
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v39<<(uint(int32(2))%32))))
						F_setDeferredAttributeLen(m, v25, v45, l1)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v48 != 0 {
								m.G0 = v8 + int32(16)
								return
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								F_valkey_free(m, v49)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if int32(3) < v28 {
							m.G0 = v8 + int32(16)
							return
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v32
							F__serverLog(m, int32(3), int32(_a921), v8)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_VM_ReplySetMapLength(m *base.Module, l0 int32, l1 int32) {
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
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v6 = m.G0
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v10&v7 == int32(0) {
		v22 = l0 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		if v23 == int32(0) {
			m.G0 = v8 + int32(16)
			return
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v26 != 0 {
				v39 = v26 + int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v39
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v39<<(uint(int32(2))%32))))
				F_setDeferredMapLen(m, v23, v45, l1)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v48 != 0 {
						m.G0 = v8 + int32(16)
						return
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						F_valkey_free(m, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(3) < v28 {
					m.G0 = v8 + int32(16)
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v32
					F__serverLog(m, int32(3), int32(_a921), v8)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v15 == int32(0) {
			m.G0 = v8 + int32(16)
			return
		} else {
			v22 = v15 + int32(36)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			if v23 == int32(0) {
				m.G0 = v8 + int32(16)
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v26 != 0 {
					v39 = v26 + int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v39<<(uint(int32(2))%32))))
					F_setDeferredMapLen(m, v23, v45, l1)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v48 != 0 {
							m.G0 = v8 + int32(16)
							return
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							F_valkey_free(m, v49)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					if int32(3) < v28 {
						m.G0 = v8 + int32(16)
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v32
						F__serverLog(m, int32(3), int32(_a921), v8)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_VM_ReplySetSetLength(m *base.Module, l0 int32, l1 int32) {
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
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v6 = m.G0
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v10&v7 == int32(0) {
		v22 = l0 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		if v23 == int32(0) {
			m.G0 = v8 + int32(16)
			return
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v26 != 0 {
				v39 = v26 + int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v39
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v39<<(uint(int32(2))%32))))
				F_setDeferredSetLen(m, v23, v45, l1)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v48 != 0 {
						m.G0 = v8 + int32(16)
						return
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						F_valkey_free(m, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(3) < v28 {
					m.G0 = v8 + int32(16)
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v32
					F__serverLog(m, int32(3), int32(_a921), v8)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v15 == int32(0) {
			m.G0 = v8 + int32(16)
			return
		} else {
			v22 = v15 + int32(36)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			if v23 == int32(0) {
				m.G0 = v8 + int32(16)
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v26 != 0 {
					v39 = v26 + int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v39<<(uint(int32(2))%32))))
					F_setDeferredSetLen(m, v23, v45, l1)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v48 != 0 {
							m.G0 = v8 + int32(16)
							return
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							F_valkey_free(m, v49)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					if int32(3) < v28 {
						m.G0 = v8 + int32(16)
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v32
						F__serverLog(m, int32(3), int32(_a921), v8)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_VM_ReplyWithArray(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_moduleReplyWithCollection(m, l0, l1, int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_VM_ReplyWithBool(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v3&int32(16) == int32(0) {
		v15 = l0 + int32(8)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		if v16 == int32(0) {
			return int32(0)
		} else {
			F_addReplyBool(m, v16, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 == int32(0) {
			return int32(0)
		} else {
			v15 = v8 + int32(36)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v16 == int32(0) {
				return int32(0)
			} else {
				F_addReplyBool(m, v16, l1)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ReplyWithCString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v3&int32(16) == int32(0) {
		v15 = l0 + int32(8)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		if v16 == int32(0) {
			return int32(0)
		} else {
			F_addReplyBulkCString(m, v16, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 == int32(0) {
			return int32(0)
		} else {
			v15 = v8 + int32(36)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v16 == int32(0) {
				return int32(0)
			} else {
				F_addReplyBulkCString(m, v16, l1)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ReplyWithEmptyArray(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v2&int32(16) == int32(0) {
		v14 = l0 + int32(8)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		if v15 == int32(0) {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[542]))
			F_addReply(m, v15, v19)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7 == int32(0) {
			return int32(0)
		} else {
			v14 = v7 + int32(36)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if v15 == int32(0) {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _consts[542]))
				F_addReply(m, v15, v19)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ReplyWithLongDouble(m *base.Module, l0 int32, l1 int64, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v4&int32(16) == int32(0) {
		v16 = l0 + int32(8)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v17 == int32(0) {
			return int32(0)
		} else {
			F_addReplyHumanLongDouble(m, v17, l1, l2)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v9 == int32(0) {
			return int32(0)
		} else {
			v16 = v9 + int32(36)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v17 == int32(0) {
				return int32(0)
			} else {
				F_addReplyHumanLongDouble(m, v17, l1, l2)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ReplyWithSet(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_moduleReplyWithCollection(m, l0, l1, int32(3))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_VM_ReplyWithSimpleString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v3&int32(16) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v16 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L3:
	;
	v15 = l0 + int32(8)
	goto L2
L4:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v15 = v8 + int32(36)
	goto L2
L6:
	;
	F_addReplyProto(m, v16, int32(_a920), int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	if l1&int32(3) == int32(0) {
		v46 = l1
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_addReplyProto(m, v16, l1, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L25
	}
L10:
	;
	v79 = v71 - l1
	goto L9
L11:
	;
	v50 = v46
	goto L19
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v32 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v35 = l1
	goto L15
L14:
	;
	v79 = l1 - l1
	goto L9
L15:
	;
	v39 = v35 + int32(1)
	if v39&int32(3) == int32(0) {
		v46 = v39
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v44 != 0 {
		v35 = v39
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v71 = v39
	goto L10
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v59 = int32(-2139062144)
	if (int32(16843008)-v56|v56)&v59 == v59 {
		v50 = v50 + int32(4)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v65 = v50
	goto L22
L21:
	;
	goto L20
L22:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v69 != 0 {
		v65 = v65 + int32(1)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v71 = v65
	goto L10
L24:
	;
	goto L23
L25:
	;
	F_addReplyProto(m, v16, int32(_a132), int32(2))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	goto L1
}
func F_VM_ReplyWithString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v3&int32(16) == int32(0) {
		v15 = l0 + int32(8)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		if v16 == int32(0) {
			return int32(0)
		} else {
			F_addReplyBulk(m, v16, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 == int32(0) {
			return int32(0)
		} else {
			v15 = v8 + int32(36)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v16 == int32(0) {
				return int32(0)
			} else {
				F_addReplyBulk(m, v16, l1)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_RetainString(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = int32(3)
	v70 = v16 + int32(-1)
	if v65 == v70 {
		v80 = v65
		goto L17
	} else {
		goto L18
	}
L2:
	;
	F_incrRefCount(m, l1)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v11&int32(1) == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v16 < int32(1) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v19 = int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v30 = int32(0)
	goto L6
L6:
	;
	v35 = v16 + (v30 ^ int32(-1))
	v38 = v23 + v35<<(uint(int32(3))%32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v39 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L2
L8:
	;
	v46 = v23 + v30<<(uint(int32(3))%32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v47 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v42 == l1 {
		v65 = v35
		v66 = v38
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v53 = v30 + int32(1)
	if v53 != int32(base.Ui32(v16+v19)>>(uint(v19)%32)) {
		v30 = v53
		goto L6
	} else {
		goto L14
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v50 != l1 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v65 = v30
	v66 = v46
	goto L1
L14:
	;
	goto L7
L15:
	;
	return
L16:
	;
	return
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v80
	return
L18:
	;
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v23+v70<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v66))) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v80 = v77 + int32(-1)
	goto L17
}
func F_VM_SaveSigned(m *base.Module, l0 int32, l1 int64) {
	var v4 int32
	_ = v4
	F_VM_SaveUnsigned(m, l0, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_VM_SaveUnsigned(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v8 == int32(0) {
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v51 = F_rdbSaveLen(m, v49, int64(2))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				if v51 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
					return
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v55 + v51
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v59 = F_rdbSaveLen(m, v58, l1)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						if v59 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
							return
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v63 + v59
							return
						}
					}
				}
			}
		} else {
			v11 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v11
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-1)))))
			switch v17 & int32(7) {
			case 0:
				v34 = int32(base.Ui32(v17) >> (uint(int32(3)) % 32))
			case 1:
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
				v34 = v24
			case 2:
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
				v34 = v27
			case 3:
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
				v34 = v30
			case 4:
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
				v34 = v33
			default:
				v34 = v11
			}
			v35 = F_rdbWriteRaw(m, v13, v8, v34)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_sdsfree(m, v8)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					if v35 < int32(0) {
						if v35 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
							return
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v51 = F_rdbSaveLen(m, v49, int64(2))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								if v51 == int32(-1) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v55 + v51
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v59 = F_rdbSaveLen(m, v58, l1)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										if v59 == int32(-1) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
											return
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = v63 + v59
											return
										}
									}
								}
							}
						}
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v41 + v35
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v51 = F_rdbSaveLen(m, v49, int64(2))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							if v51 == int32(-1) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
								return
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v55 + v51
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v59 = F_rdbSaveLen(m, v58, l1)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									if v59 == int32(-1) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
										return
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v63 + v59
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
}
func F_VM_ScanCursorCreate(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_valkey_malloc(m, int32(16))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(0)
		return v3
	}
}
func F_VM_ScanKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v16 = int32(28)
	if l0 == v5 {
		v211 = v5
		v212 = v16
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L60
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v19 == int32(0) {
		v211 = v5
		v212 = v16
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	switch v22&int32(15) + int32(-2) {
	case 0:
		goto L7
	case 1:
		goto L5
	case 2:
		goto L6
	default:
		v211 = v5
		v212 = v16
		goto L1
	}
L4:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v47 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	if v22&int32(240) == int32(112) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	if v22&int32(240) == int32(32) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	if v22&int32(240) == int32(32) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v32 = F_objectGetVal(m, v19)
	mBase = m.M
	v46 = v32
	goto L4
L9:
	;
	v46 = int32(0)
	goto L4
L10:
	;
	v38 = F_objectGetVal(m, v19)
	mBase = m.M
	v46 = v38
	goto L4
L11:
	;
	v46 = int32(0)
	goto L4
L12:
	;
	v44 = F_objectGetVal(m, v19)
	mBase = m.M
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = v45
	goto L4
L13:
	;
	v46 = int32(0)
	goto L4
L14:
	;
	if v46 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v211 = v5
	v212 = int32(44)
	goto L1
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	switch v74&int32(15) + int32(-2) {
	case 0:
		goto L24
	case 1, 2:
		goto L23
	default:
		v211 = int32(1)
		v212 = int32(0)
		goto L1
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l0
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v60 = F_hashtableScan(m, v46, v56, int32(570), v13+int32(16))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = base.I64_extend_i32_u(v60)
	if v60 != 0 {
		v70 = int32(1)
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v211 = v70
	v212 = int32(0)
	goto L1
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
	v70 = int32(0)
	goto L20
L22:
	;
	v211 = int32(0)
	v212 = v201
	goto L1
L23:
	;
	v122 = int32(0)
	v123 = F_objectGetVal(m, v19)
	mBase = m.M
	v125 = F_lpSeek(m, v123, v122)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L18
	} else {
		goto L38
	}
L24:
	;
	v79 = F_setTypeInitIterator(m, v19)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L18
	} else {
		goto L26
	}
L25:
	;
	F_setTypeReleaseIterator(m, v79)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L18
	} else {
		goto L36
	}
L26:
	;
	v81 = F_setTypeNextObject(m, v79)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	if v81 == int32(0) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v91 = v81
	goto L29
L29:
	;
	v96 = F_createObject(m, int32(0), v91)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L18
	} else {
		goto L31
	}
L30:
	;
	goto L25
L31:
	;
	m.T0[l2].(func(*base.Module, int32, int32, int32, int32))(m, l0, v96, int32(0), l3)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	F_decrRefCount(m, v96)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v103 = F_setTypeNextObject(m, v79)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	if v103 != 0 {
		v91 = v103
		goto L29
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(1)
	v201 = int32(0)
	goto L22
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(1)
	v201 = v122
	goto L22
L38:
	;
	if v125 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v134 = v125
	goto L40
L40:
	;
	v143 = F_lpGetValue(m, v134, v13+int32(12), v13+int32(16))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L18
	} else {
		goto L44
	}
L41:
	;
	goto L37
L42:
	;
	v154 = F_objectGetVal(m, v19)
	mBase = m.M
	v155 = F_lpNext(m, v154, v134)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L18
	} else {
		goto L50
	}
L43:
	;
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	v151 = F_createStringObjectFromLongLongWithSds(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L18
	} else {
		goto L47
	}
L44:
	;
	if v143 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v148 = F_createStringObject_1(m, v143, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L18
	} else {
		goto L46
	}
L46:
	;
	v153 = v148
	goto L42
L47:
	;
	v153 = v151
	goto L42
L48:
	;
	m.T0[l2].(func(*base.Module, int32, int32, int32, int32))(m, l0, v153, v171, l3)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L18
	} else {
		goto L55
	}
L49:
	;
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	v169 = F_createStringObjectFromLongLongWithSds(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L18
	} else {
		goto L54
	}
L50:
	;
	v161 = F_lpGetValue(m, v155, v13+int32(12), v13+int32(16))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	if v161 == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v166 = F_createStringObject_1(m, v161, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L18
	} else {
		goto L53
	}
L53:
	;
	v171 = v166
	goto L48
L54:
	;
	v171 = v169
	goto L48
L55:
	;
	v174 = F_objectGetVal(m, v19)
	mBase = m.M
	v175 = F_lpNext(m, v174, v155)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L18
	} else {
		goto L56
	}
L56:
	;
	F_decrRefCount(m, v153)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L18
	} else {
		goto L57
	}
L57:
	;
	F_decrRefCount(m, v171)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L18
	} else {
		goto L58
	}
L58:
	;
	if v175 != 0 {
		v134 = v175
		goto L40
	} else {
		goto L59
	}
L59:
	;
	goto L41
L60:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v212
	m.G0 = v13 + int32(32)
	return v211
}
func F_VM_ScriptingEngineDebuggerFlushLogs(m *base.Module) {
	var v2 int32
	_ = v2
	F_scriptingEngineDebuggerFlushLogs(m)
	v2 = m.ExcPending
	if v2 != 0 {
		return
	} else {
		return
	}
}
func F_VM_SelectDb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = F_selectDb(m, v3, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v4 != int32(0))
	}
}
func F_VM_SetExpire(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v5 = int32(1)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v6&int32(2) == int32(0) {
		v36 = v5
		return v36
	} else {
		if l1 < int64(-1) {
			v36 = v5
			return v36
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v13 == int32(0) {
				v36 = v5
				return v36
			} else {
				if l1 == int64(-1) {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v32 = F_removeExpire(m, v30, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v36 = int32(0)
						return v36
					}
				} else {
					v19 = *(*int64)(unsafe.Add(mBase, _consts[98]))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v25 = F_setExpire(m, v21, v22, v23, v19+l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v25
						v36 = int32(0)
						return v36
					}
				}
			}
		}
	}
}
func F_VM_SetModuleOptions(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+48)) = l1
	return
}
func F_VM_SetModuleUserACL(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = F_ACLSetUser(m, v3, l1, int32(-1))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_VM_StreamIteratorStop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	if l0 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(138)
			return int32(1)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if v11&int32(15) == int32(6) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v21 != 0 {
					F_streamIteratorStop(m, v21)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						F_valkey_free(m, v31)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v34
							return v34
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(8)
					return int32(1)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(138)
				return int32(1)
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
		return int32(1)
	}
}
func F_VM_StreamTrimByLength(m *base.Module, l0 int32, l1 int32, l2 int64) int64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
		return int64(-1)
	} else {
		if base.Ui32(int32(1)) < base.Ui32(l1) {
			*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
			return int64(-1)
		} else {
			if int64(-1) < l2 {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v16 == int32(0) {
					*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(138)
					return int64(-1)
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					if v19&int32(15) == int32(6) {
						v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v29&int32(2) != 0 {
							v37 = F_objectGetVal(m, v16)
							mBase = m.M
							v38 = F_streamTrimByLength(m, v37, l2, l1)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int64(0)
							} else {
								return v38
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(8)
							return int64(-1)
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(138)
						return int64(-1)
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
				return int64(-1)
			}
		}
	}
}
func F_VM_StringAppendBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v5&int32(-8) == int32(8) {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		switch int32(base.Ui32(v24)>>(uint(int32(4))%32))&int32(15) + int32(-1) {
		case 0:
			v33 = F_objectGetVal(m, l1)
			mBase = m.M
			v35 = F_sdsfromlonglong(m, base.I64_extend_i32_s(v33))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				F_objectSetVal(m, l1, v35)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v39 & int32(-241)
					v43 = F_objectGetVal(m, l1)
					mBase = m.M
					v44 = F_sdscatlen(m, v43, l2, l3)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						F_objectSetVal(m, l1, v44)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v49 = int32(0)
							return v49
						}
					}
				}
			}
		default:
			v43 = F_objectGetVal(m, l1)
			mBase = m.M
			v44 = F_sdscatlen(m, v43, l2, l3)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				F_objectSetVal(m, l1, v44)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v49 = int32(0)
					return v49
				}
			}
		case 7:
			F_objectUnembedVal(m, l1)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v43 = F_objectGetVal(m, l1)
				mBase = m.M
				v44 = F_sdscatlen(m, v43, l2, l3)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_objectSetVal(m, l1, v44)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v49 = int32(0)
						return v49
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[6]))
		if int32(3) < v12 {
			v49 = int32(1)
			return v49
		} else {
			F__serverLog(m, int32(3), int32(_a919), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return int32(1)
			}
		}
	}
}
func F_VM_StringCompare(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_compareStringObjects(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_VM_StringTruncate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int64
	_ = v208
	var v211 int64
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v12&int32(2) == int32(0) {
		v227 = v11
		m.G0 = v9 + int32(16)
		return v227
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v17 == int32(0) {
			if base.Ui32(int32(536870912)) < base.Ui32(l1) {
				v227 = v11
				m.G0 = v9 + int32(16)
				return v227
			} else {
				v41 = int32(0)
				if l1 == v41 {
					v227 = v41
					m.G0 = v9 + int32(16)
					return v227
				} else {
					v44 = int32(0)
					v46 = F_sdsnewlen(m, v44, l1)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v48 = F_createObject(m, v44, v46)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v48
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							F_setKey(m, v52, v53, v54, v9+int32(12), int32(10))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v60
								v227 = v41
								m.G0 = v9 + int32(16)
								return v227
							}
						}
					}
				}
			}
		} else {
			if base.Ui32(int32(536870912)) < base.Ui32(l1) {
				v227 = v11
				m.G0 = v9 + int32(16)
				return v227
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				if v22&int32(15) != 0 {
					v227 = v11
					m.G0 = v9 + int32(16)
					return v227
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v27 = F_dbUnshareStringValue(m, v25, v26, v17)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v27
						v33 = F_objectGetVal(m, v27)
						mBase = m.M
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+int32(-1)))))
						switch v36 & int32(7) {
						case 0:
							v76 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
						case 1:
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+int32(-3)))))
							v76 = v66
						case 2:
							v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33+int32(-5)))))
							v76 = v69
						case 3:
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(-9))))
							v76 = v72
						case 4:
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(-17))))
							v76 = v75
						default:
							v76 = int32(0)
						}
						if base.Ui32(l1) <= base.Ui32(v76) {
							if base.Ui32(l1) < base.Ui32(v76) {
								v87 = int32(0)
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v89 = F_objectGetVal(m, v88)
								mBase = m.M
								v97 = v89 + int32(-1)
								v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
								v100 = v98 & int32(7)
								switch v100 {
								case 0:
									v115 = int32(base.Ui32(v98) >> (uint(int32(3)) % 32))
								case 1:
									v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+int32(-3)))))
									v115 = v105
								case 2:
									v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89+int32(-5)))))
									v115 = v108
								case 3:
									v111 = *(*int32)(unsafe.Add(mBase, uint32(v89+int32(-9))))
									v115 = v111
								case 4:
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v89+int32(-17))))
									v115 = v114
								default:
									v115 = v87
								}
								v117 = base.B2i32(base.Ui32(v87) < base.Ui32(v115))
								if base.Ui32(v87) < base.Ui32(v115) {
									v118 = v87
								} else {
									v118 = int32(0)
								}
								v119 = v115 - v118
								if base.Ui32(l1) < base.Ui32(v119) {
									v121 = l1
								} else {
									v121 = v119
								}
								if base.Ui32(v87) < base.Ui32(v115) {
									v123 = v121
								} else {
									v123 = int32(0)
								}
								if v123 == int32(0) {
								} else {
									v127 = F_memmove(m, v89, v89+v118, v123)
									mBase = m.M
								}
								v129 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v89+v123))) = uint8(v129)
								switch v100 {
								case 0:
									v132 = v123 << (uint(int32(3)) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v132)
								case 1:
									*(*uint8)(unsafe.Add(mBase, uint32(v89+int32(-3)))) = uint8(v123)
								case 2:
									*(*uint16)(unsafe.Add(mBase, uint32(v89+int32(-5)))) = uint16(v123)
								case 3:
									*(*int32)(unsafe.Add(mBase, uint32(v89+int32(-9)))) = v123
								case 4:
									*(*int64)(unsafe.Add(mBase, uint32(v89+int32(-17)))) = base.I64_extend_i32_u(v123)
								default:
								}
								v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v148 = F_objectGetVal(m, v147)
								mBase = m.M
								v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+int32(-1)))))
								switch v154 & int32(7) {
								case 0:
									v173 = int32(base.Ui32(v154) >> (uint(int32(3)) % 32))
								case 1:
									v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+int32(-3)))))
									v173 = v161
								case 2:
									v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148+int32(-5)))))
									v173 = v164
								case 3:
									v167 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-9))))
									v173 = v167
								case 4:
									v170 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-17))))
									v171 = v170
									v173 = v171
								default:
									v171 = int32(0)
									v173 = v171
								}
								v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v175 = F_objectGetVal(m, v174)
								mBase = m.M
								v178 = int32(-1)
								v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v178))))
								switch v180&int32(7) + v178 {
								case 0:
									v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+int32(-2)))))
									v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+int32(-3)))))
									v216 = v187 - v190
								case 1:
									v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175+int32(-3)))))
									v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175+int32(-5)))))
									v216 = v194 - v197
								case 2:
									v201 = *(*int32)(unsafe.Add(mBase, uint32(v175+int32(-5))))
									v204 = *(*int32)(unsafe.Add(mBase, uint32(v175+int32(-9))))
									v216 = v201 - v204
								case 3:
									v208 = *(*int64)(unsafe.Add(mBase, uint32(v175+int32(-9))))
									v211 = *(*int64)(unsafe.Add(mBase, uint32(v175+int32(-17))))
									v214 = base.I32_wrap_i64(v208 - v211)
									v216 = v214
								default:
									v214 = int32(0)
									v216 = v214
								}
								if base.Ui32(v216) <= base.Ui32(v173) {
									v227 = v87
									m.G0 = v9 + int32(16)
									return v227
								} else {
									v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v219 = int32(0)
									v220 = F_objectGetVal(m, v218)
									mBase = m.M
									v222 = F_sdsRemoveFreeSpace(m, v220, v219)
									mBase = m.M
									v223 = m.ExcPending
									if v223 != 0 {
										return int32(0)
									} else {
										F_objectSetVal(m, v218, v222)
										mBase = m.M
										v225 = m.ExcPending
										if v225 != 0 {
											return int32(0)
										} else {
											v227 = v219
											m.G0 = v9 + int32(16)
											return v227
										}
									}
								}
							} else {
								v227 = int32(0)
								m.G0 = v9 + int32(16)
								return v227
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v79 = F_objectGetVal(m, v78)
							mBase = m.M
							v80 = F_sdsgrowzero(m, v79, l1)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_objectSetVal(m, v78, v80)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									v227 = int32(0)
									m.G0 = v9 + int32(16)
									return v227
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_VM_ThreadSafeContextTryLock(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v33 int64
	_ = v33
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v48 int64
	_ = v48
	var v60 int32
	_ = v60
	v12 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	if v12 != 0 {
		F__serverAssert(m, int32(_a947), int32(_a917), int32(9239))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v15 = int32(0)
		v19 = *(*int32)(unsafe.Add(mBase, _consts[95]))
		*(*int32)(unsafe.Add(mBase, _consts[95])) = v19 + int32(1)
		if v19 != 0 {
		} else {
			v27 = F_ustime(m)
			mBase = m.M
			v29 = int32(0)
			*(*int64)(unsafe.Add(mBase, _consts[96])) = v27
			v33 = base.I64_div_s(v27, int64(1000))
			*(*int64)(unsafe.Add(mBase, _consts[35])) = v33
			v37 = base.I64_div_s(v27, int64(1000000))
			*(*int64)(unsafe.Add(mBase, _consts[47])) = v37
			v40 = *(*int32)(unsafe.Add(mBase, _consts[97]))
			F_lrulfu_updateClockAndPolicy(m, v33, int32(base.Ui32(v40&int32(2))>>(uint(int32(1))%32)))
			mBase = m.M
			v48 = *(*int64)(unsafe.Add(mBase, _consts[35]))
			*(*int64)(unsafe.Add(mBase, _consts[98])) = v48
		}
		return int32(0)
	}
}
func F_VM_TryAlloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	if base.Ui32(int32(2147483646)) < base.Ui32(l0) {
		v50 = int32(0)
	} else {
		if l0 != 0 {
			v11 = l0
		} else {
			v11 = int32(4)
		}
		v13 = v11 + int32(8)
		v14 = F_emscripten_builtin_malloc(m, v13)
		mBase = m.M
		if v14 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v11
			v19 = *(*int32)(unsafe.Add(mBase, _consts[411]))
			if v19 != int32(-1) {
				v30 = v19
			} else {
				v22 = int32(0)
				v24 = *(*int32)(unsafe.Add(mBase, _consts[281]))
				*(*int32)(unsafe.Add(mBase, _consts[411])) = v24
				*(*int32)(unsafe.Add(mBase, _consts[281])) = v24 + int32(1)
				v30 = v24
			}
			if v30 < int32(260) {
				v39 = v30 << (uint(int32(2)) % 32)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[285])))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[285]))) = v42 + v13
			} else {
				v33 = int32(0)
				v35 = *(*int32)(unsafe.Add(mBase, _consts[286]))
				*(*int32)(unsafe.Add(mBase, _consts[286])) = v35 + v13
			}
			v50 = v14 + int32(8)
		} else {
			v50 = int32(0)
		}
	}
	return v50
}
func F_VM_ValueLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v2 = int32(0)
	if l0 == v2 {
		v32 = v2
		return v32
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v6 == int32(0) {
			v32 = v2
			return v32
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			switch v9 & int32(15) {
			case 0:
				v12 = F_stringObjectLen(m, v6)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					return v12
				}
			case 1:
				v17 = F_listTypeLength(m, v6)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			case 2:
				v20 = F_setTypeSize(m, v6)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v20
				}
			case 3:
				v23 = F_zsetLength(m, v6)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return v23
				}
			case 4:
				v26 = F_hashTypeLength(m, v6)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					return v26
				}
			default:
				v32 = v2
				return v32
			case 6:
				v29 = F_objectGetVal(m, v6)
				mBase = m.M
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
				v32 = v30
				return v32
			}
		}
	}
}
func F_VM_ZsetAdd(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v5
	v15 = int32(1)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v16&int32(2) == v5 {
		v120 = v15
		m.G0 = v11 + int32(16)
		return v120
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v21 == int32(0) {
			v29 = F_createZsetListpackObject(m)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v29
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				F_dbAdd(m, v34, v35, v11+int32(12))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v40
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					switch v42&int32(15) + int32(-3) {
					case 0:
						*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
						v53 = v40
					default:
						v53 = v40
					case 3:
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
						v53 = v40
					}
					v54 = int32(0)
					if l3 == v54 {
						v94 = v54
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						v58 = int32(2)
						v60 = int32(24)
						v65 = base.I32_rotl(v57&int32(3), v60)
						v70 = int32(4)
						v72 = int32(251662080)
						v78 = int32(base.Ui32(v65)>>(uint(v70)%32))&v72 | v65&v72<<(uint(v70)%32)
						v94 = int32(base.Ui32(v57)>>(uint(v58)%32))&v60 | int32(base.Ui32(v65<<(uint(int32(5))%32)&int32(1073741824)|(int32(base.Ui32(v78)>>(uint(v58)%32))|v78&int32(268435456)<<(uint(v58)%32))&int32(1342177280)<<(uint(int32(1))%32))>>(uint(int32(29))%32))
					}
					v95 = F_objectGetVal(m, l2)
					mBase = m.M
					v99 = F_zsetAdd(m, v53, l1, v95, v94, v11+int32(8), int32(0))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						if v99 != 0 {
							v107 = int32(0)
							if l3 == v107 {
								v120 = v107
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v110<<(uint(int32(4))%32)&int32(16) | v110&int32(12)
								v120 = v107
							}
							m.G0 = v11 + int32(16)
							return v120
						} else {
							if l3 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
							}
							v105 = F_moduleDelKeyIfEmpty(m, l0)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								v120 = v15
								m.G0 = v11 + int32(16)
								return v120
							}
						}
					}
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			if v24&int32(15) == int32(3) {
				v53 = v21
				v54 = int32(0)
				if l3 == v54 {
					v94 = v54
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					v58 = int32(2)
					v60 = int32(24)
					v65 = base.I32_rotl(v57&int32(3), v60)
					v70 = int32(4)
					v72 = int32(251662080)
					v78 = int32(base.Ui32(v65)>>(uint(v70)%32))&v72 | v65&v72<<(uint(v70)%32)
					v94 = int32(base.Ui32(v57)>>(uint(v58)%32))&v60 | int32(base.Ui32(v65<<(uint(int32(5))%32)&int32(1073741824)|(int32(base.Ui32(v78)>>(uint(v58)%32))|v78&int32(268435456)<<(uint(v58)%32))&int32(1342177280)<<(uint(int32(1))%32))>>(uint(int32(29))%32))
				}
				v95 = F_objectGetVal(m, l2)
				mBase = m.M
				v99 = F_zsetAdd(m, v53, l1, v95, v94, v11+int32(8), int32(0))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					if v99 != 0 {
						v107 = int32(0)
						if l3 == v107 {
							v120 = v107
						} else {
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v110<<(uint(int32(4))%32)&int32(16) | v110&int32(12)
							v120 = v107
						}
						m.G0 = v11 + int32(16)
						return v120
					} else {
						if l3 == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
						}
						v105 = F_moduleDelKeyIfEmpty(m, l0)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							v120 = v15
							m.G0 = v11 + int32(16)
							return v120
						}
					}
				}
			} else {
				v120 = v15
				m.G0 = v11 + int32(16)
				return v120
			}
		}
	}
}
func F_VM_ZsetLastInLexRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_zsetInitLexRange(m, l0, l1, l2, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_VM_ZsetRangeStop(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		if v6&int32(15) != int32(3) {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v11 != int32(1) {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
				return
			} else {
				F_zsetFreeLexRange(m, l0+int32(56))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
					return
				}
			}
		}
	}
}
