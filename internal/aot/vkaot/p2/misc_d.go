package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F___dup3(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	for {
		v9 = m.Env.X__syscall_dup3(m, l0, l1, l2)
		mBase = m.M
		if v9 == int32(-10) {
			continue
		} else {
			break
		}
		break
	}
	if base.Ui32(v9) < base.Ui32(int32(-4095)) {
		v19 = v9
	} else {
		v14 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(0) - v9
		v19 = int32(-1)
	}
	return v19
}
func F_daemonize(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = m.G0
	m.G0 = v3 - int32(16)
	v7 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(52)
	m.Env.Exit(m, int32(0))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_databasesCron(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int64
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
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
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	v1 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[380]))
	if v10 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L21
L2:
	;
	v13 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v14 == v13 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	if v32 != 0 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	F_expireReplicaKeys(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v25 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v26 == v25 {
		goto L3
	} else {
		goto L10
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	goto L7
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	goto L8
L8:
	;
	if v20&int32(1) == int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L3
L10:
	;
	goto L4
L11:
	;
	return
L12:
	;
	goto L1
L13:
	;
	v36 = int32(0)
	v40 = *(*int32)(unsafe.Add(mBase, _consts[422]))
	if v40 == v36 {
		v46 = v36
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v34 = F_activeExpireCycle(m, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	if v46 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v46 = v43 + v44
	goto L17
L19:
	;
	F_flushReplicaKeysWithExpireList(m, int32(1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	goto L1
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	if v53 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v57 < int32(1) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v60 = int32(16)
	if v57 < v60 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v63 = v57
	goto L27
L26:
	;
	v63 = v60
	goto L27
L27:
	;
	v64 = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[807]))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v69 = v57
	v71 = v65
	v72 = v67
	v73 = v64
	goto L28
L28:
	;
	v77 = base.I32_rem_u_s(v71, v69)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v72+v77<<(uint(int32(2))%32))))
	v82 = int32(0)
	v84 = v71 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[807])) = v84
	if v81 == v82 {
		v106 = v69
		v107 = v84
		v108 = v72
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v112 = int32(0)
	v113 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	if v113 == v112 {
		goto L22
	} else {
		goto L36
	}
L30:
	;
	v110 = v73 + int32(1)
	if v110 != v63 {
		v69 = v106
		v71 = v107
		v72 = v108
		v73 = v110
		goto L28
	} else {
		goto L35
	}
L31:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	F_kvstoreTryResizeHashtables(m, v88, int32(16))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	F_kvstoreTryResizeHashtables(m, v92, int32(16))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	F_kvstoreTryResizeHashtables(m, v96, int32(16))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	v100 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	v103 = *(*int32)(unsafe.Add(mBase, _consts[807]))
	v105 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v106 = v101
	v107 = v103
	v108 = v105
	goto L30
L35:
	;
	goto L29
L36:
	;
	v117 = int32(0)
	v118 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v119 = base.I32_div_s(int32(1000000), v118)
	v121 = base.I32_div_s(v119, int32(100))
	v122 = base.I64_extend_i32_s(v121)
	v124 = *(*int32)(unsafe.Add(mBase, _consts[809]))
	v126 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v129 = v106
	v132 = v126
	v133 = v117
	v134 = v124
	v136 = int64(0)
	goto L37
L37:
	;
	v137 = base.I32_rem_u_s(v134, v129)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v132+v137<<(uint(int32(2))%32))))
	if v141 == int32(0) {
		v168 = v129
		v169 = v132
		v170 = v134
		v171 = v136
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L22
L39:
	;
	v173 = int32(1)
	v174 = v170 + v173
	*(*int32)(unsafe.Add(mBase, _consts[809])) = v174
	v177 = v133 + v173
	if v177 != v63 {
		v129 = v168
		v132 = v169
		v133 = v177
		v134 = v174
		v136 = v171
		goto L37
	} else {
		goto L47
	}
L40:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v146 = F_kvstoreIncrementallyRehash(m, v144, v122-v136)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	v148 = v146 + v136
	if base.Ui64(v122) <= base.Ui64(v148) {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v152 = F_kvstoreIncrementallyRehash(m, v150, v122-v148)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	v154 = v152 + v148
	if base.Ui64(v122) <= base.Ui64(v154) {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v158 = F_kvstoreIncrementallyRehash(m, v156, v122-v154)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L11
	} else {
		goto L45
	}
L45:
	;
	v160 = v158 + v154
	if base.Ui64(v122) <= base.Ui64(v160) {
		goto L22
	} else {
		goto L46
	}
L46:
	;
	v162 = int32(0)
	v163 = *(*int32)(unsafe.Add(mBase, _consts[809]))
	v165 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v167 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	v168 = v167
	v169 = v165
	v170 = v163
	v171 = v160
	goto L39
L47:
	;
	goto L38
}
func F_dbsizeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	if v4 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		if v9 != 0 {
			v11 = F_hashtableSize(m, v9)
			mBase = m.M
			v14 = base.I64_extend_i32_u(v11)
		} else {
			v14 = int64(0)
		}
	} else {
		v7 = *(*int64)(unsafe.Add(mBase, uint32(v3)+40))
		v14 = v7
	}
	F_addReplyLongLong(m, l0, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		return
	}
}
func F_debugDelay(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int64
	_ = v7
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	if int32(-1) < l0 {
		v21 = l0
	} else {
		v5 = int32(0)
		v7 = *(*int64)(unsafe.Add(mBase, _consts[408]))
		v11 = v7*int64(6364136223846793005) + int64(1)
		*(*int64)(unsafe.Add(mBase, _consts[408])) = v11
		v16 = int32(0)
		v18 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v11)>>(uint(int64(33))%64))), v16-l0)
		v21 = base.B2i32(v18 == v16)
	}
	if v21 == int32(0) {
		return
	} else {
		v24 = F_usleep(m, v21)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			return
		}
	}
}
func F_decrbyCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v13 = F_getLongLongFromObjectOrReply(m, l0, v9, v6+int32(8), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			m.G0 = v6 + int32(16)
			return
		} else {
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
			if v15 != int64(-9223372036854775807-1) {
				F_incrDecrCommand(m, l0, int64(0)-v15)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			} else {
				F_addReplyError(m, l0, int32(_a1721))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			}
		}
	}
}
func F_delCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, _consts[338]))
	F_delGenericCommand(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_delGenericCommand(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int64
	_ = v100
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int64
	_ = v118
	var v127 int32
	_ = v127
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v12 < int32(2) {
		v118 = int64(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_addReplyLongLong(m, l0, v118)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L30
	}
L2:
	;
	v20 = int32(0)
	v21 = int32(1)
	goto L3
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = v21 << (uint(int32(2)) % 32)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v30)))
	v33 = F_objectGetVal(m, v32)
	mBase = m.M
	v34 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v36 == v34 {
		v41 = v34
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v118 = base.I64_extend_i32_s(v106)
	goto L1
L5:
	;
	v42 = int32(0)
	v44 = F_expireIfNeededWithDictIndex(m, v27, v32, v42, v42, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L10
	}
L6:
	;
	v39 = F_getKeySlot(m, v33)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v41 = v39
	goto L5
L9:
	;
	v112 = v21 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v112 < v113 {
		v20 = v106
		v21 = v112
		goto L3
	} else {
		goto L29
	}
L10:
	;
	if v44 == int32(2) {
		v106 = v20
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49+v30)))
	v52 = F_objectGetVal(m, v51)
	mBase = m.M
	v54 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if l1 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v77 == int32(0) {
		v106 = v20
		goto L9
	} else {
		goto L25
	}
L13:
	;
	if v54 != 0 {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	if v54 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v62 = int32(1)
	v64 = F_getKeySlot(m, v52)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L18
	}
L16:
	;
	v57 = int32(1)
	v60 = F_dbGenericDeleteWithDictIndex(m, v48, v51, v57, v57, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v77 = v60
	goto L12
L18:
	;
	v66 = F_dbGenericDeleteWithDictIndex(m, v48, v51, v62, v62, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v77 = v66
	goto L12
L20:
	;
	v74 = F_dbGenericDeleteWithDictIndex(m, v48, v51, int32(0), int32(1), v71)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L24
	}
L21:
	;
	v69 = F_getKeySlot(m, v52)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L23
	}
L22:
	;
	v71 = int32(0)
	goto L20
L23:
	;
	v71 = v69
	goto L20
L24:
	;
	v77 = v74
	goto L12
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81+v30)))
	F_touchWatchedKey(m, v80, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	F_trackingInvalidateKey(m, l0, v83, int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v91+v30)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a213), v93, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v98 = int32(_a44)
	v100 = *(*int64)(unsafe.Add(mBase, _consts[83]))
	*(*int64)(unsafe.Add(mBase, _consts[83])) = v100 + int64(1)
	v106 = v20 + int32(1)
	goto L9
L29:
	;
	goto L4
L30:
	;
	return
}
func F_deleteExpiredKeyFromOverwriteAndPropagate(m *base.Module, l0 int32, l1 int32) {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(_a44)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[350]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v14 = F_objectGetVal(m, l1)
	mBase = m.M
	v16 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v16 != 0 {
		v18 = F_getKeySlot(m, v14)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = v18
			v22 = F_dbGenericDeleteWithDictIndex(m, v13, l1, v12, int32(2), v20)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				if v22 != 0 {
					v30 = int32(_a44)
					v32 = *(*int64)(unsafe.Add(mBase, _consts[83]))
					*(*int64)(unsafe.Add(mBase, _consts[83])) = v32 + int64(1)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
					v41 = *(*int32)(unsafe.Add(mBase, _consts[350]))
					if v41 != 0 {
						v42 = int32(244)
					} else {
						v42 = int32(240)
					}
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[84])))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v44
					F_rewriteClientCommandVector(m, l0, int32(2), v9)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						F_touchWatchedKey(m, v49, l1)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							F_trackingInvalidateKey(m, l0, l1, int32(1))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
								F_notifyKeyspaceEvent(m, int32(256), int32(_a583), l1, v58)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									v61 = int32(_a44)
									v63 = *(*int64)(unsafe.Add(mBase, _consts[351]))
									*(*int64)(unsafe.Add(mBase, _consts[351])) = v63 + int64(1)
									m.G0 = v9 + int32(16)
									return
								}
							}
						}
					}
				} else {
					F__serverAssertWithInfo(m, l0, l1, int32(_a555), int32(_a550), int32(1974))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
		}
	} else {
		v20 = int32(0)
		v22 = F_dbGenericDeleteWithDictIndex(m, v13, l1, v12, int32(2), v20)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			if v22 != 0 {
				v30 = int32(_a44)
				v32 = *(*int64)(unsafe.Add(mBase, _consts[83]))
				*(*int64)(unsafe.Add(mBase, _consts[83])) = v32 + int64(1)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
				v41 = *(*int32)(unsafe.Add(mBase, _consts[350]))
				if v41 != 0 {
					v42 = int32(244)
				} else {
					v42 = int32(240)
				}
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[84])))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v44
				F_rewriteClientCommandVector(m, l0, int32(2), v9)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					F_touchWatchedKey(m, v49, l1)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_trackingInvalidateKey(m, l0, l1, int32(1))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
							F_notifyKeyspaceEvent(m, int32(256), int32(_a583), l1, v58)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v61 = int32(_a44)
								v63 = *(*int64)(unsafe.Add(mBase, _consts[351]))
								*(*int64)(unsafe.Add(mBase, _consts[351])) = v63 + int64(1)
								m.G0 = v9 + int32(16)
								return
							}
						}
					}
				}
			} else {
				F__serverAssertWithInfo(m, l0, l1, int32(_a555), int32(_a550), int32(1974))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
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
	}
}
func F_dirname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(_a955)
L2:
	;
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l0&int32(3) == int32(0) {
		v30 = l0
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v65 = v63
	goto L22
L5:
	;
	v63 = v55 - l0
	goto L4
L6:
	;
	v34 = v30
	goto L14
L7:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v19 = l0
	goto L10
L9:
	;
	v63 = l0 - l0
	goto L4
L10:
	;
	v23 = v19 + int32(1)
	if v23&int32(3) == int32(0) {
		v30 = v23
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v28 != 0 {
		v19 = v23
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v55 = v23
	goto L5
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v43 = int32(-2139062144)
	if (int32(16843008)-v40|v40)&v43 == v43 {
		v34 = v34 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v49 = v34
	goto L17
L16:
	;
	goto L15
L17:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v53 != 0 {
		v49 = v49 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v55 = v49
	goto L5
L19:
	;
	goto L18
L20:
	;
	return int32(_a1760)
L21:
	;
	v85 = v79
	goto L31
L22:
	;
	v68 = v65 + int32(-1)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v68))))
	if v70 == int32(47) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v68 != 0 {
		v65 = v68
		goto L22
	} else {
		goto L30
	}
L25:
	;
	v74 = v68
	goto L26
L26:
	;
	if v74 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L21
L28:
	;
	v79 = v74 + int32(-1)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v79))))
	if v81 != int32(47) {
		v74 = v79
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L20
L31:
	;
	if v85 == int32(0) {
		goto L20
	} else {
		goto L33
	}
L32:
	;
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v91+int32(1)))) = uint8(v97)
	return l0
L33:
	;
	v90 = v85 + int32(-1)
	v91 = l0 + v90
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v92 == int32(47) {
		v85 = v90
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
}
func F_discardCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v2&int32(8) != 0 {
		F_resetClientMultiState(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v10 & int32(-4137)
			F_unwatchAllKeys(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _consts[84]))
				F_addReply(m, l0, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		F_addReplyError(m, l0, int32(_a974))
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	}
}
func F_discardCommandQueue(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	v6 = l0 + int32(56)
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	if base.Ui32(v8) <= base.Ui32(v7) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	F_valkey_free(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L9
	} else {
		goto L14
	}
L2:
	;
	v12 = v7
	goto L3
L3:
	;
	v14 = int32(1)
	v15 = v12 + v14
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)) = uint16(v15)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v22 = v17 + v12&int32(65535)*int32(40)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v23 < v14 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	F_valkey_free(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L12
	}
L6:
	;
	v29 = int32(0)
	goto L7
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v29<<(uint(int32(2))%32))))
	F_decrRefCount(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	return
L10:
	;
	v39 = v29 + int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v39 < v40 {
		v29 = v39
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	if base.Ui32(v49) < base.Ui32(v50) {
		v12 = v49
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L4
L14:
	;
	v61 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(64)))) = uint16(v61)
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(0)
	return
}
func F_discardTempDb(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
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
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
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
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	v8 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v8 < int32(1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F__serverAssert(m, int32(_a556), int32(_a550), int32(714))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L22
	} else {
		goto L47
	}
L2:
	;
	F__serverAssert(m, int32(_a557), int32(_a550), int32(713))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L22
	} else {
		goto L46
	}
L3:
	;
	F__serverAssert(m, int32(_a558), int32(_a550), int32(712))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L22
	} else {
		goto L45
	}
L4:
	;
	F__serverAssert(m, int32(_a559), int32(_a550), int32(711))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L22
	} else {
		goto L44
	}
L5:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L22
	} else {
		goto L43
	}
L6:
	;
	v17 = int32(0)
	goto L7
L7:
	;
	v22 = l0 + v17<<(uint(int32(2))%32)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v23 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v75 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v77 <= v75 {
		goto L5
	} else {
		goto L25
	}
L9:
	;
	if v17 != v8+int32(-1) {
		v17 = v17 + int32(1)
		goto L7
	} else {
		goto L24
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v27 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v37 == int64(0) {
		goto L9
	} else {
		goto L16
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v26)+40))
	v37 = v30
	goto L11
L14:
	;
	v34 = F_hashtableSize(m, v32)
	mBase = m.M
	v37 = base.I64_extend_i32_u(v34)
	goto L11
L15:
	;
	v37 = int64(0)
	goto L11
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	if v42 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	F_emptyDbAsync(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v47 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	v49 = F_hashtableSize(m, v47)
	mBase = m.M
	goto L17
L21:
	;
	goto L17
L22:
	;
	return
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v57 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v56)+32)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v56+int32(56)))) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v56+int32(48)))) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v56+int32(40)))) = v57
	goto L9
L24:
	;
	goto L8
L25:
	;
	v83 = v75
	v84 = v77
	goto L26
L26:
	;
	v88 = l0 + v83<<(uint(int32(2))%32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v89 == int32(0) {
		v150 = v84
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L5
L28:
	;
	v153 = v83 + int32(1)
	if v153 < v150 {
		v83 = v153
		v84 = v150
		goto L26
	} else {
		goto L42
	}
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	F_kvstoreRelease(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	F_kvstoreRelease(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	F_kvstoreRelease(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L22
	} else {
		goto L32
	}
L32:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	if v105 != int32(0)-v107 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	if v111 != int32(0)-v113 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)+20))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	if v117 != int32(0)-v119 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v103)+24))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	if v123 != int32(0)-v125 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_dictRelease(m, v104)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L22
	} else {
		goto L37
	}
L37:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	F_dictRelease(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L22
	} else {
		goto L38
	}
L38:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	F_dictRelease(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L22
	} else {
		goto L39
	}
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+24))
	F_dictRelease(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L22
	} else {
		goto L40
	}
L40:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	F_valkey_free(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L22
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = int32(0)
	v148 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	v150 = v148
	goto L28
L42:
	;
	goto L27
L43:
	;
	return
L44:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_discharge2reg(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v55 int32
	_ = v55
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
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 float64
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	F_luaK_dischargevars(m, l0, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		switch v15 + int32(-1) {
		case 0:
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v18 <= v19 {
				v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
				v133 = F_luaK_code(m, l0, l2<<(uint(int32(23))%32)|l2<<(uint(int32(6))%32)|int32(3), v132)
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
					m.G0 = v11 + int32(16)
					return
				}
			} else {
				if v18 != 0 {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
					v29 = v24 + v18<<(uint(int32(2))%32) + int32(-4)
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					if v30&int32(63) != int32(3) {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
						v133 = F_luaK_code(m, l0, l2<<(uint(int32(23))%32)|l2<<(uint(int32(6))%32)|int32(3), v132)
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
							m.G0 = v11 + int32(16)
							return
						}
					} else {
						if l2 < int32(base.Ui32(v30)>>(uint(int32(6))%32))&int32(255) {
							v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
							v133 = F_luaK_code(m, l0, l2<<(uint(int32(23))%32)|l2<<(uint(int32(6))%32)|int32(3), v132)
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
								m.G0 = v11 + int32(16)
								return
							}
						} else {
							v41 = int32(base.Ui32(v30) >> (uint(int32(23)) % 32))
							if v41+int32(1) < l2 {
								v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
								v133 = F_luaK_code(m, l0, l2<<(uint(int32(23))%32)|l2<<(uint(int32(6))%32)|int32(3), v132)
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
									m.G0 = v11 + int32(16)
									return
								}
							} else {
								if base.Ui32(l2) <= base.Ui32(v41) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v29))) = v30&int32(8388547) | l2<<(uint(int32(23))%32)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
								m.G0 = v11 + int32(16)
								return
							}
						}
					}
				} else {
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
					if v21 <= l2 {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
						m.G0 = v11 + int32(16)
						return
					} else {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
						v133 = F_luaK_code(m, l0, l2<<(uint(int32(23))%32)|l2<<(uint(int32(6))%32)|int32(3), v132)
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
							m.G0 = v11 + int32(16)
							return
						}
					}
				}
			}
		case 1, 2:
			v55 = int32(2)
			if v15 == v55 {
				v58 = int32(_a0)
			} else {
				v58 = v55
			}
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
			v64 = F_luaK_code(m, l0, l2<<(uint(int32(6))%32)|v58|int32(2), v63)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
				m.G0 = v11 + int32(16)
				return
			}
		case 3:
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
			v76 = F_luaK_code(m, l0, l2<<(uint(int32(6))%32)|v68<<(uint(int32(14))%32)|int32(1), v75)
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
				m.G0 = v11 + int32(16)
				return
			}
		case 4:
			v78 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v11))) = v78
			v84 = F_addk(m, l0, v11, v11)
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
				v93 = F_luaK_code(m, l0, l2<<(uint(int32(6))%32)|v84<<(uint(int32(14))%32)|int32(1), v92)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
					m.G0 = v11 + int32(16)
					return
				}
			}
		default:
			m.G0 = v11 + int32(16)
			return
		case 10:
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v100 = v96 + v97<<(uint(int32(2))%32)
			v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
			*(*int32)(unsafe.Add(mBase, uint32(v100))) = v101&int32(-16321) | l2<<(uint(int32(6))%32)&int32(16320)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
			m.G0 = v11 + int32(16)
			return
		case 11:
			v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if l2 == v110 {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
				m.G0 = v11 + int32(16)
				return
			} else {
				v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
				v119 = F_luaK_code(m, l0, v110<<(uint(int32(23))%32)|l2<<(uint(int32(6))%32), v118)
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
					m.G0 = v11 + int32(16)
					return
				}
			}
		}
	}
}
func F_do_putc_1(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v1 = l0
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v5 < int32(0) {
		v16 = v1 & int32(255)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
		if v16 == v17 {
			v28 = F___overflow(m, l1, v16)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				return v28
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			if v19 == v20 {
				v28 = F___overflow(m, l1, v16)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					return v28
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v19 + int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v1)
				return v16
			}
		}
	} else {
		if v5 == int32(0) {
			v33 = F_locking_putc_1(m, v1, l1)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				return v33
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[1]))
			if v5&int32(1073741823) != v13 {
				v33 = F_locking_putc_1(m, v1, l1)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					return v33
				}
			} else {
				v16 = v1 & int32(255)
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
				if v16 == v17 {
					v28 = F___overflow(m, l1, v16)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						return v28
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					if v19 == v20 {
						v28 = F___overflow(m, l1, v16)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							return v28
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v19 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v1)
						return v16
					}
				}
			}
		}
	}
}
func F_double2ll(m *base.Module, l0 float64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 float64
	_ = v7
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	v3 = int32(0)
	v7 = base.F64_abs(l0)
	if base.F64_gt(v7, float64(4.611686018427388e+18)) != 0 {
		v21 = v3
	} else {
		if base.F64_lt(v7, float64(9.223372036854776e+18)) == int32(0) {
			v16 = int64(-9223372036854775807 - 1)
		} else {
			v14 = base.I64_trunc_f64_s(l0)
			v16 = v14
		}
		if base.F64_ne(l0, base.F64_convert_i64_s(v16)) != 0 {
			v21 = v3
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v16
			v21 = int32(1)
		}
	}
	return v21
}
func F_drainIOThreadsQueue(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v20 int32
	_ = v20
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
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	v3 = *(*int32)(unsafe.Add(mBase, _consts[430]))
	if v3 != 0 {
		F__serverAssert(m, int32(_a845), int32(_a846), int32(85))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, _consts[431]))
		if v5 < int32(2) {
		} else {
			v9 = int32(1)
			for {
				v11 = v9 * int32(192)
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[432])))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[433])))
				if v15 == v16 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[433]))) = v15
				}
				v20 = v9 + int32(1)
				v22 = *(*int32)(unsafe.Add(mBase, _consts[431]))
				if v20 < v22 {
					v9 = v20
					continue
				} else {
					break
				}
				break
			}
		}
		v25 = int32(0)
		v26 = *(*int32)(unsafe.Add(mBase, _consts[434]))
		v28 = *(*int32)(unsafe.Add(mBase, _consts[435]))
		if v26 == v28 {
		} else {
			for {
				v31 = int32(0)
				v32 = *(*int32)(unsafe.Add(mBase, _consts[434]))
				v34 = *(*int32)(unsafe.Add(mBase, _consts[435]))
				if v32 != v34 {
					continue
				} else {
					break
				}
				break
			}
		}
		return
	}
}
func F_dropInstanceConnections(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
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
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int64
	_ = v225
	var v226 int64
	_ = v226
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5&int32(1) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a1355), int32(_a1333), int32(1129))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L7
	} else {
		goto L84
	}
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v11 == int32(0) {
		v27 = v10
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v28 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v16 != v11 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+212)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(1)
	F_valkeyAsyncFree(m, v11)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(0)
	goto L5
L7:
	;
	return
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v27 = v26
	goto L3
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v44 = F_dictGetIterator(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L15
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v31 != v28 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v28)+212)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(1)
	F_valkeyAsyncFree(m, v28)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L13
	}
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = int64(0)
	goto L11
L13:
	;
	goto L9
L14:
	;
	F_dictReleaseIterator(m, v44)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L7
	} else {
		goto L83
	}
L15:
	;
	v53 = v44 + int32(20)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	if v54 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v149 == int32(0) {
		goto L14
	} else {
		goto L42
	}
L17:
	;
	v60 = v53
	v61 = v57
	goto L20
L18:
	;
	v57 = int32(1)
	goto L17
L19:
	;
	v57 = int32(0)
	goto L17
L20:
	;
	switch v61 {
	case 0:
		goto L25
	default:
		goto L24
	}
L22:
	;
	v61 = int32(0)
	goto L20
L23:
	;
	goto L16
L24:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v141
	if v141 == int32(0) {
		goto L22
	} else {
		goto L41
	}
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v65 != int32(-1) {
		v104 = v65
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v105 = int32(1)
	v106 = v104 + v105
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v106
	v108 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+v112+int32(26)))))
	if v116 == int32(255) {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v69 != 0 {
		v104 = int32(-1)
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if v71 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	if v98 != int32(-1) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v78 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v70)+16)))
	v79 = int64(*(*int8)(unsafe.Add(mBase, uint32(v70)+27)))
	v80 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+8)))
	v81 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v70)+12)))
	v82 = int64(*(*int8)(unsafe.Add(mBase, uint32(v70)+26)))
	v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+4)))
	v84 = F_wangHash64(m, v83)
	mBase = m.M
	v86 = F_wangHash64(m, v82+v84)
	mBase = m.M
	v88 = F_wangHash64(m, v81+v86)
	mBase = m.M
	v90 = F_wangHash64(m, v80+v88)
	mBase = m.M
	v92 = F_wangHash64(m, v79+v90)
	mBase = m.M
	v94 = F_wangHash64(m, v78+v92)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v44)+24)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v97 = v96
	goto L29
L31:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+24)))
	v76 = v74 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+24)) = uint16(v76)
	v97 = v70
	goto L29
L32:
	;
	v104 = v98 + int32(-1)
	goto L26
L33:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v104 = v101
	goto L26
L34:
	;
	v131 = int32(2)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v111+v129<<(uint(v131)%32)+int32(4))))
	v60 = v136 + v130<<(uint(v131)%32)
	v61 = int32(1)
	goto L20
L35:
	;
	v120 = v108
	goto L37
L36:
	;
	v120 = v105 << (uint(v116) % 32)
	goto L37
L37:
	;
	if v106 < v120 {
		v129 = v112
		v130 = v106
		goto L34
	} else {
		goto L38
	}
L38:
	;
	if v112 != 0 {
		v149 = v108
		goto L23
	} else {
		goto L39
	}
L39:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	if v122 == int32(-1) {
		v149 = v108
		goto L23
	} else {
		goto L40
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v44)+4)) = int64(4294967296)
	v129 = int32(1)
	v130 = int32(0)
	goto L34
L41:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v145
	v149 = v141
	goto L23
L42:
	;
	v156 = v149
	goto L43
L43:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)+8))
	goto L46
L44:
	;
	goto L14
L45:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	if v178 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+28))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	if v161 == int32(0) {
		v177 = v160
		goto L45
	} else {
		goto L47
	}
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v160)+8)) = int64(0)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+16))
	if v166 != v161 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+212)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = int32(1)
	F_valkeyAsyncFree(m, v161)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L7
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+16)) = int32(0)
	goto L48
L50:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v159)+28))
	v177 = v176
	goto L45
L51:
	;
	v200 = v44 + int32(20)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	if v201 != 0 {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	if v181 != v178 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v185 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v177)+16)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v178)+212)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v177)+4)) = int32(1)
	F_valkeyAsyncFree(m, v178)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L55
	}
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v177)+8)) = int64(0)
	goto L53
L55:
	;
	goto L51
L56:
	;
	if v296 != 0 {
		v156 = v296
		goto L43
	} else {
		goto L82
	}
L57:
	;
	v207 = v200
	v208 = v204
	goto L60
L58:
	;
	v204 = int32(1)
	goto L57
L59:
	;
	v204 = int32(0)
	goto L57
L60:
	;
	switch v208 {
	case 0:
		goto L65
	default:
		goto L64
	}
L62:
	;
	v208 = int32(0)
	goto L60
L63:
	;
	goto L56
L64:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v288
	if v288 == int32(0) {
		goto L62
	} else {
		goto L81
	}
L65:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v212 != int32(-1) {
		v251 = v212
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v252 = int32(1)
	v253 = v251 + v252
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v253
	v255 = int32(0)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258+v259+int32(26)))))
	if v263 == int32(255) {
		goto L75
	} else {
		goto L76
	}
L67:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v216 != 0 {
		v251 = int32(-1)
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if v218 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+20))
	if v245 != int32(-1) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v225 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v217)+16)))
	v226 = int64(*(*int8)(unsafe.Add(mBase, uint32(v217)+27)))
	v227 = int64(*(*int32)(unsafe.Add(mBase, uint32(v217)+8)))
	v228 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v217)+12)))
	v229 = int64(*(*int8)(unsafe.Add(mBase, uint32(v217)+26)))
	v230 = int64(*(*int32)(unsafe.Add(mBase, uint32(v217)+4)))
	v231 = F_wangHash64(m, v230)
	mBase = m.M
	v233 = F_wangHash64(m, v229+v231)
	mBase = m.M
	v235 = F_wangHash64(m, v228+v233)
	mBase = m.M
	v237 = F_wangHash64(m, v227+v235)
	mBase = m.M
	v239 = F_wangHash64(m, v226+v237)
	mBase = m.M
	v241 = F_wangHash64(m, v225+v239)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v44)+24)) = v241
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v244 = v243
	goto L69
L71:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+24)))
	v223 = v221 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v217)+24)) = uint16(v223)
	v244 = v217
	goto L69
L72:
	;
	v251 = v245 + int32(-1)
	goto L66
L73:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v251 = v248
	goto L66
L74:
	;
	v278 = int32(2)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v258+v276<<(uint(v278)%32)+int32(4))))
	v207 = v283 + v277<<(uint(v278)%32)
	v208 = int32(1)
	goto L60
L75:
	;
	v267 = v255
	goto L77
L76:
	;
	v267 = v252 << (uint(v263) % 32)
	goto L77
L77:
	;
	if v253 < v267 {
		v276 = v259
		v277 = v253
		goto L74
	} else {
		goto L78
	}
L78:
	;
	if v259 != 0 {
		v296 = v255
		goto L63
	} else {
		goto L79
	}
L79:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v258)+20))
	if v269 == int32(-1) {
		v296 = v255
		goto L63
	} else {
		goto L80
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v44)+4)) = int64(4294967296)
	v276 = int32(1)
	v277 = int32(0)
	goto L74
L81:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v292
	v296 = v288
	goto L63
L82:
	;
	goto L44
L83:
	;
	return
L84:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dualChannelSyncSuccess(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
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
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v54 int64
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v63 int64
	_ = v63
	var v68 int64
	_ = v68
	var v75 int64
	_ = v75
	var v82 int64
	_ = v82
	var v89 int64
	_ = v89
	var v96 int64
	_ = v96
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int64
	_ = v141
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	v4 = int32(_a44)
	v6 = *(*int64)(unsafe.Add(mBase, _consts[679]))
	*(*int64)(unsafe.Add(mBase, _consts[672])) = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[680]))
	v11 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	F_replicationCreatePrimaryClientWithHandler(m, v9, v11, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = int32(_a44)
		v16 = *(*int32)(unsafe.Add(mBase, _consts[188]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
		v19 = *(*int64)(unsafe.Add(mBase, _consts[682]))
		*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = v19
		v21 = int32(144)
		v26 = int32(*(*uint8)(unsafe.Add(mBase, _consts[683])))
		*(*uint8)(unsafe.Add(mBase, uint32(v17+v21))) = uint8(v26)
		v28 = int32(136)
		v33 = *(*int64)(unsafe.Add(mBase, _consts[684]))
		*(*int64)(unsafe.Add(mBase, uint32(v17+v28))) = v33
		v35 = int32(128)
		v40 = *(*int64)(unsafe.Add(mBase, _consts[685]))
		*(*int64)(unsafe.Add(mBase, uint32(v17+v35))) = v40
		v42 = int32(120)
		v47 = *(*int64)(unsafe.Add(mBase, _consts[686]))
		*(*int64)(unsafe.Add(mBase, uint32(v17+v42))) = v47
		v49 = int32(112)
		v54 = *(*int64)(unsafe.Add(mBase, _consts[687]))
		*(*int64)(unsafe.Add(mBase, uint32(v17+v49))) = v54
		v57 = *(*int32)(unsafe.Add(mBase, _consts[188]))
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
		v60 = *(*int64)(unsafe.Add(mBase, _consts[679]))
		*(*int64)(unsafe.Add(mBase, uint32(v58)+48)) = v60
		v63 = *(*int64)(unsafe.Add(mBase, _consts[688]))
		*(*int64)(unsafe.Add(mBase, uint32(v58)+40)) = v63
		*(*int64)(unsafe.Add(mBase, _consts[40])) = v60
		v68 = *(*int64)(unsafe.Add(mBase, uint32(v58)+104))
		*(*int64)(unsafe.Add(mBase, _consts[689])) = v68
		v75 = *(*int64)(unsafe.Add(mBase, uint32(v58+v49)))
		*(*int64)(unsafe.Add(mBase, _consts[690])) = v75
		v82 = *(*int64)(unsafe.Add(mBase, uint32(v58+v42)))
		*(*int64)(unsafe.Add(mBase, _consts[691])) = v82
		v89 = *(*int64)(unsafe.Add(mBase, uint32(v58+v35)))
		*(*int64)(unsafe.Add(mBase, _consts[692])) = v89
		v96 = *(*int64)(unsafe.Add(mBase, uint32(v58+v28)))
		*(*int64)(unsafe.Add(mBase, _consts[693])) = v96
		v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v21))))
		*(*uint8)(unsafe.Add(mBase, _consts[663])) = uint8(v103)
		v105 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v57
		v107 = *(*int32)(unsafe.Add(mBase, uint32(v57)+200))
		*(*int32)(unsafe.Add(mBase, uint32(v57)+200)) = v107 & int32(-65)
		v112 = *(*int32)(unsafe.Add(mBase, _consts[188]))
		v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+200))
		*(*int32)(unsafe.Add(mBase, uint32(v112)+200)) = v113 & int32(-1025)
		v118 = *(*int32)(unsafe.Add(mBase, _consts[188]))
		*(*int32)(unsafe.Add(mBase, uint32(v118)+328)) = int32(0)
		v123 = *(*int32)(unsafe.Add(mBase, uint32(v118)+204))
		*(*int32)(unsafe.Add(mBase, uint32(v118)+204)) = v123&int32(-25165825) | int32(_a0) | int32(16777216)
		v138 = int32(_a44)
		v139 = *(*int32)(unsafe.Add(mBase, _consts[188]))
		v141 = *(*int64)(unsafe.Add(mBase, _consts[47]))
		*(*int64)(unsafe.Add(mBase, uint32(v139)+88)) = v141
		*(*int64)(unsafe.Add(mBase, _consts[694])) = int64(0)
		*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(14)
		v150 = int32(0)
		F_moduleFireServerEvent(m, int64(7), v150, v150)
		mBase = m.M
		v153 = m.ExcPending
		if v153 != 0 {
			return
		} else {
			v155 = *(*int32)(unsafe.Add(mBase, _consts[695]))
			if v155 == int32(0) {
				v188 = int32(_a44)
				*(*int32)(unsafe.Add(mBase, _consts[696])) = int32(0)
				*(*int64)(unsafe.Add(mBase, _consts[695])) = int64(0)
				v195 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(2) < v195 {
					F_replicationSteadyStateInit(m)
					mBase = m.M
					v204 = m.ExcPending
					if v204 != 0 {
						return
					} else {
						F_replicationSendAck(m)
						mBase = m.M
						v206 = m.ExcPending
						if v206 != 0 {
							return
						} else {
							v207 = int32(_a44)
							*(*int32)(unsafe.Add(mBase, _consts[697])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[698])) = int64(-1)
							return
						}
					}
				} else {
					F__serverLog(m, int32(2), int32(_a1254), int32(0))
					mBase = m.M
					v202 = m.ExcPending
					if v202 != 0 {
						return
					} else {
						F_replicationSteadyStateInit(m)
						mBase = m.M
						v204 = m.ExcPending
						if v204 != 0 {
							return
						} else {
							F_replicationSendAck(m)
							mBase = m.M
							v206 = m.ExcPending
							if v206 != 0 {
								return
							} else {
								v207 = int32(_a44)
								*(*int32)(unsafe.Add(mBase, _consts[697])) = int32(0)
								*(*int64)(unsafe.Add(mBase, _consts[698])) = int64(-1)
								return
							}
						}
					}
				}
			} else {
				v159 = *(*int32)(unsafe.Add(mBase, _consts[188]))
				v160 = F_streamReplDataBufToDb(m, v159)
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return
				} else {
					if v160 != int32(-1) {
						v182 = *(*int32)(unsafe.Add(mBase, _consts[695]))
						if v182 == int32(0) {
							v188 = int32(_a44)
							*(*int32)(unsafe.Add(mBase, _consts[696])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[695])) = int64(0)
							v195 = *(*int32)(unsafe.Add(mBase, _consts[6]))
							if int32(2) < v195 {
								F_replicationSteadyStateInit(m)
								mBase = m.M
								v204 = m.ExcPending
								if v204 != 0 {
									return
								} else {
									F_replicationSendAck(m)
									mBase = m.M
									v206 = m.ExcPending
									if v206 != 0 {
										return
									} else {
										v207 = int32(_a44)
										*(*int32)(unsafe.Add(mBase, _consts[697])) = int32(0)
										*(*int64)(unsafe.Add(mBase, _consts[698])) = int64(-1)
										return
									}
								}
							} else {
								F__serverLog(m, int32(2), int32(_a1254), int32(0))
								mBase = m.M
								v202 = m.ExcPending
								if v202 != 0 {
									return
								} else {
									F_replicationSteadyStateInit(m)
									mBase = m.M
									v204 = m.ExcPending
									if v204 != 0 {
										return
									} else {
										F_replicationSendAck(m)
										mBase = m.M
										v206 = m.ExcPending
										if v206 != 0 {
											return
										} else {
											v207 = int32(_a44)
											*(*int32)(unsafe.Add(mBase, _consts[697])) = int32(0)
											*(*int64)(unsafe.Add(mBase, _consts[698])) = int64(-1)
											return
										}
									}
								}
							}
						} else {
							F_freePendingReplDataBufAsync(m, v182)
							mBase = m.M
							v186 = m.ExcPending
							if v186 != 0 {
								return
							} else {
								v188 = int32(_a44)
								*(*int32)(unsafe.Add(mBase, _consts[696])) = int32(0)
								*(*int64)(unsafe.Add(mBase, _consts[695])) = int64(0)
								v195 = *(*int32)(unsafe.Add(mBase, _consts[6]))
								if int32(2) < v195 {
									F_replicationSteadyStateInit(m)
									mBase = m.M
									v204 = m.ExcPending
									if v204 != 0 {
										return
									} else {
										F_replicationSendAck(m)
										mBase = m.M
										v206 = m.ExcPending
										if v206 != 0 {
											return
										} else {
											v207 = int32(_a44)
											*(*int32)(unsafe.Add(mBase, _consts[697])) = int32(0)
											*(*int64)(unsafe.Add(mBase, _consts[698])) = int64(-1)
											return
										}
									}
								} else {
									F__serverLog(m, int32(2), int32(_a1254), int32(0))
									mBase = m.M
									v202 = m.ExcPending
									if v202 != 0 {
										return
									} else {
										F_replicationSteadyStateInit(m)
										mBase = m.M
										v204 = m.ExcPending
										if v204 != 0 {
											return
										} else {
											F_replicationSendAck(m)
											mBase = m.M
											v206 = m.ExcPending
											if v206 != 0 {
												return
											} else {
												v207 = int32(_a44)
												*(*int32)(unsafe.Add(mBase, _consts[697])) = int32(0)
												*(*int64)(unsafe.Add(mBase, _consts[698])) = int64(-1)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v165 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if int32(3) < v165 {
							v174 = *(*int32)(unsafe.Add(mBase, _consts[697]))
							if v174 == int32(0) {
								return
							} else {
								F_replicationAbortDualChannelSyncTransfer(m)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return
								} else {
									F_replicationUnsetPrimary(m)
									mBase = m.M
									v180 = m.ExcPending
									if v180 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							F__serverLog(m, int32(3), int32(_a1255), int32(0))
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return
							} else {
								v174 = *(*int32)(unsafe.Add(mBase, _consts[697]))
								if v174 == int32(0) {
									return
								} else {
									F_replicationAbortDualChannelSyncTransfer(m)
									mBase = m.M
									v178 = m.ExcPending
									if v178 != 0 {
										return
									} else {
										F_replicationUnsetPrimary(m)
										mBase = m.M
										v180 = m.ExcPending
										if v180 != 0 {
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
		}
	}
}
func F_dummy_2(m *base.Module, l0 int32) {
	return
}
func F_dummy_5(m *base.Module, l0 int32, l1 int32) {
	return
}
func F_dup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v2 = m.Env.X__syscall_dup(m, l0)
	mBase = m.M
	if base.Ui32(v2) < base.Ui32(int32(-4095)) {
		v10 = v2
	} else {
		v5 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(0) - v2
		v10 = int32(-1)
	}
	return v10
}
func F_dupStringObject(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v373 int32
	_ = v373
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11&int32(15) != 0 {
		F__serverAssert(m, int32(_a1055), int32(_a1051), int32(467))
		mBase = m.M
		v373 = m.ExcPending
		if v373 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		switch int32(base.Ui32(v11)>>(uint(int32(4))%32)) & int32(15) {
		case 0:
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v18&int32(4) == int32(0) {
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v79 = v18
				v80 = v78
			} else {
				if v18&int32(1) != 0 {
					v27 = int32(16)
				} else {
					v27 = int32(8)
				}
				v28 = l0 + v27
				if v18&int32(2) == int32(0) {
					v60 = v28
				} else {
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
					v34 = v28 + v33
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
					switch v38 & int32(7) {
					case 0:
						v55 = int32(base.Ui32(v38) >> (uint(int32(3)) % 32))
					case 1:
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+int32(-2)))))
						v55 = v45
					case 2:
						v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(-4)))))
						v55 = v48
					case 3:
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(-8))))
						v55 = v51
					case 4:
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(-16))))
						v55 = v54
					default:
						v55 = int32(0)
					}
					v60 = v34 + int32(1) + v55 + int32(1)
				}
				v74 = *(*int32)(unsafe.Add(mBase, _consts[601]))
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v79 = v77
				v80 = v60 + v74
			}
			if v79&int32(4) == int32(0) {
				v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v143 = v141
			} else {
				if v79&int32(1) != 0 {
					v91 = int32(16)
				} else {
					v91 = int32(8)
				}
				v92 = l0 + v91
				if v79&int32(2) == int32(0) {
					v123 = v92
				} else {
					v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
					v98 = v92 + v97
					v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
					switch v102 & int32(7) {
					case 0:
						v119 = int32(base.Ui32(v102) >> (uint(int32(3)) % 32))
					case 1:
						v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+int32(-2)))))
						v119 = v109
					case 2:
						v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98+int32(-4)))))
						v119 = v112
					case 3:
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v98+int32(-8))))
						v119 = v115
					case 4:
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v98+int32(-16))))
						v119 = v118
					default:
						v119 = int32(0)
					}
					v123 = v98 + int32(1) + v119 + int32(1)
				}
				v138 = *(*int32)(unsafe.Add(mBase, _consts[601]))
				v143 = v123 + v138
			}
			v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143+int32(-1)))))
			switch v149 & int32(7) {
			case 0:
				v166 = int32(base.Ui32(v149) >> (uint(int32(3)) % 32))
			case 1:
				v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143+int32(-3)))))
				v166 = v156
			case 2:
				v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143+int32(-5)))))
				v166 = v159
			case 3:
				v162 = *(*int32)(unsafe.Add(mBase, uint32(v143+int32(-9))))
				v166 = v162
			case 4:
				v165 = *(*int32)(unsafe.Add(mBase, uint32(v143+int32(-17))))
				v166 = v165
			default:
				v166 = int32(0)
			}
			v167 = F_sdsnewlen(m, v80, v166)
			mBase = m.M
			v170 = m.ExcPending
			if v170 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(0)
				v176 = F_zmalloc_usable(m, int32(12), v9+int32(8))
				mBase = m.M
				v177 = m.ExcPending
				if v177 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v176)+8)) = v167
					*(*int64)(unsafe.Add(mBase, uint32(v176))) = int64(34359738368)
					v361 = v176
					m.G0 = v9 + int32(16)
					return v361
				}
			}
		case 1:
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
			v343 = int32(12)
			v346 = F_zmalloc_usable(m, v343, v9+v343)
			mBase = m.M
			v347 = m.ExcPending
			if v347 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v346)+8)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v346))) = int64(34359738368)
				v352 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
				*(*int32)(unsafe.Add(mBase, uint32(v346))) = v352&int32(-241) | int32(16)
				v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v346)+8)) = v358
				v361 = v346
				m.G0 = v9 + int32(16)
				return v361
			}
		default:
			F__serverPanic_1(m, int32(_a1051), int32(477), int32(_a1056), int32(0))
			mBase = m.M
			v339 = m.ExcPending
			if v339 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 8:
			v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v181&int32(4) == int32(0) {
				v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v242 = v181
				v243 = v241
			} else {
				if v181&int32(1) != 0 {
					v190 = int32(16)
				} else {
					v190 = int32(8)
				}
				v191 = l0 + v190
				if v181&int32(2) == int32(0) {
					v223 = v191
				} else {
					v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
					v197 = v191 + v196
					v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
					switch v201 & int32(7) {
					case 0:
						v218 = int32(base.Ui32(v201) >> (uint(int32(3)) % 32))
					case 1:
						v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197+int32(-2)))))
						v218 = v208
					case 2:
						v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197+int32(-4)))))
						v218 = v211
					case 3:
						v214 = *(*int32)(unsafe.Add(mBase, uint32(v197+int32(-8))))
						v218 = v214
					case 4:
						v217 = *(*int32)(unsafe.Add(mBase, uint32(v197+int32(-16))))
						v218 = v217
					default:
						v218 = int32(0)
					}
					v223 = v197 + int32(1) + v218 + int32(1)
				}
				v237 = *(*int32)(unsafe.Add(mBase, _consts[601]))
				v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v242 = v240
				v243 = v223 + v237
			}
			if v242&int32(4) == int32(0) {
				v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v306 = v304
			} else {
				if v242&int32(1) != 0 {
					v254 = int32(16)
				} else {
					v254 = int32(8)
				}
				v255 = l0 + v254
				if v242&int32(2) == int32(0) {
					v286 = v255
				} else {
					v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
					v261 = v255 + v260
					v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
					switch v265 & int32(7) {
					case 0:
						v282 = int32(base.Ui32(v265) >> (uint(int32(3)) % 32))
					case 1:
						v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261+int32(-2)))))
						v282 = v272
					case 2:
						v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261+int32(-4)))))
						v282 = v275
					case 3:
						v278 = *(*int32)(unsafe.Add(mBase, uint32(v261+int32(-8))))
						v282 = v278
					case 4:
						v281 = *(*int32)(unsafe.Add(mBase, uint32(v261+int32(-16))))
						v282 = v281
					default:
						v282 = int32(0)
					}
					v286 = v261 + int32(1) + v282 + int32(1)
				}
				v301 = *(*int32)(unsafe.Add(mBase, _consts[601]))
				v306 = v286 + v301
			}
			v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306+int32(-1)))))
			switch v312 & int32(7) {
			case 0:
				v329 = int32(base.Ui32(v312) >> (uint(int32(3)) % 32))
			case 1:
				v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306+int32(-3)))))
				v329 = v319
			case 2:
				v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v306+int32(-5)))))
				v329 = v322
			case 3:
				v325 = *(*int32)(unsafe.Add(mBase, uint32(v306+int32(-9))))
				v329 = v325
			case 4:
				v328 = *(*int32)(unsafe.Add(mBase, uint32(v306+int32(-17))))
				v329 = v328
			default:
				v329 = int32(0)
			}
			v332 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v243, v329, int32(0), int64(-1))
			mBase = m.M
			v333 = m.ExcPending
			if v333 != 0 {
				return int32(0)
			} else {
				v361 = v332
				m.G0 = v9 + int32(16)
				return v361
			}
		}
	}
}
