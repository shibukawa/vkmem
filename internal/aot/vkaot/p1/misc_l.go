package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F___lockfile(m *base.Module, l0 int32) int32 {
	return int32(1)
}
func F___lshrti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v21 int64
	_ = v21
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	if l3&int32(64) == int32(0) {
		if l3 == int32(0) {
			v25 = l1
			v26 = l2
		} else {
			v21 = base.I64_extend_i32_u(l3)
			v25 = l2<<(uint(base.I64_extend_i32_u(int32(64)-l3))%64) | int64(base.Ui64(l1)>>(uint(v21)%64))
			v26 = int64(base.Ui64(l2) >> (uint(v21) % 64))
		}
	} else {
		v25 = int64(base.Ui64(l2) >> (uint(base.I64_extend_i32_u(l3+int32(-64))) % 64))
		v26 = int64(0)
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v25
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v26
	return
}
func F_lastsaveCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int64)(unsafe.Add(mBase, _consts[295]))
	F_addReplyLongLong(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_lazyFreeTrackingTable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	F_freeTrackingRadixTree(m, v3)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = int32(0)
		v9 = *(*int32)(unsafe.Add(mBase, _consts[373]))
		*(*int32)(unsafe.Add(mBase, _consts[373])) = v9 - v4
		v14 = *(*int32)(unsafe.Add(mBase, _consts[374]))
		*(*int32)(unsafe.Add(mBase, _consts[374])) = v4 + v14
		return
	}
}
func F_lazyfreePendingReplDataBuf(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+20))
	F_listRelease(m, v3)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = int32(0)
		v9 = *(*int32)(unsafe.Add(mBase, _consts[373]))
		*(*int32)(unsafe.Add(mBase, _consts[373])) = v9 - v4
		v14 = *(*int32)(unsafe.Add(mBase, _consts[374]))
		*(*int32)(unsafe.Add(mBase, _consts[374])) = v4 + v14
		return
	}
}
func F_lazyfreeResetStats(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	v1 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[374])) = v1
	return
}
func F_lfu_getFrequency(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v27 int32
	_ = v27
	v3 = int32(0)
	v7 = int32(*(*uint16)(unsafe.Add(mBase, _consts[376])))
	v10 = *(*int32)(unsafe.Add(mBase, _consts[377]))
	if v10 == v3 {
		v21 = v3
	} else {
		v16 = int32(65535)
		v18 = base.I32_div_s((v7-int32(base.Ui32(l0)>>(uint(int32(8))%32)))&v16, v10)
		v21 = v18 & v16
	}
	v24 = l0 & int32(255)
	v25 = v24 - v21
	if base.Ui32(v24) < base.Ui32(v25) {
		v27 = int32(0)
	} else {
		v27 = v25
	}
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v27)
	return v27 | v7<<(uint(int32(8))%32)
}
func F_lfu_import(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = int32(*(*uint16)(unsafe.Add(mBase, _consts[376])))
	return v3<<(uint(int32(8))%32) | l0
}
func F_linsertCommand(m *base.Module, l0 int32) {
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
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
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int64
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v12 = F_objectGetVal(m, v11)
	mBase = m.M
	v13 = int32(_a2382)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v8 + int32(48)
	return
L2:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L31
	} else {
		goto L59
	}
L3:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v100 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v101 = F_lookupKeyWriteOrReply(m, l0, v98, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L31
	} else {
		goto L32
	}
L4:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v55 = F_objectGetVal(m, v54)
	mBase = m.M
	v56 = int32(_a2383)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v59 != 0 {
		goto L20
	} else {
		goto L21
	}
L5:
	;
	if v48-v50 != 0 {
		goto L4
	} else {
		goto L17
	}
L6:
	;
	v48 = F_tolower(m, v44)
	mBase = m.M
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v50 = F_tolower(m, v49)
	mBase = m.M
	goto L5
L7:
	;
	v18 = v12
	v19 = v13
	v20 = v16
	goto L10
L8:
	;
	v44 = int32(0)
	v45 = v13
	goto L6
L9:
	;
	v44 = v41 & int32(255)
	v45 = v40
	goto L6
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v22 == int32(0) {
		v40 = v19
		v41 = v20
		goto L9
	} else {
		goto L12
	}
L11:
	;
	v40 = v34
	v41 = int32(0)
	goto L9
L12:
	;
	v26 = v20 & int32(255)
	if v26 == v22 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v33 = int32(1)
	v34 = v19 + v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v35 != 0 {
		v18 = v18 + v33
		v19 = v34
		v20 = v35
		goto L10
	} else {
		goto L16
	}
L14:
	;
	v28 = F_tolower(m, v26)
	mBase = m.M
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v30 = F_tolower(m, v29)
	mBase = m.M
	if v28 == v30 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v40 = v19
	v41 = v32
	goto L9
L16:
	;
	goto L11
L17:
	;
	v96 = int32(1)
	goto L3
L18:
	;
	if v91-v93 != 0 {
		goto L2
	} else {
		goto L30
	}
L19:
	;
	v91 = F_tolower(m, v87)
	mBase = m.M
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	v93 = F_tolower(m, v92)
	mBase = m.M
	goto L18
L20:
	;
	v61 = v55
	v62 = v56
	v63 = v59
	goto L23
L21:
	;
	v87 = int32(0)
	v88 = v56
	goto L19
L22:
	;
	v87 = v84 & int32(255)
	v88 = v83
	goto L19
L23:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v65 == int32(0) {
		v83 = v62
		v84 = v63
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v83 = v77
	v84 = int32(0)
	goto L22
L25:
	;
	v69 = v63 & int32(255)
	if v69 == v65 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v76 = int32(1)
	v77 = v62 + v76
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v78 != 0 {
		v61 = v61 + v76
		v62 = v77
		v63 = v78
		goto L23
	} else {
		goto L29
	}
L27:
	;
	v71 = F_tolower(m, v69)
	mBase = m.M
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v73 = F_tolower(m, v72)
	mBase = m.M
	if v71 == v73 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v83 = v62
	v84 = v75
	goto L22
L29:
	;
	goto L24
L30:
	;
	v96 = int32(0)
	goto L3
L31:
	;
	return
L32:
	;
	if v101 == int32(0) {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v106 = F_checkType(m, l0, v101, int32(1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if v106 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v110 = int32(4)
	v112 = int32(0)
	F_listTypeTryConversionRaw(m, v101, int32(1), v109, v110, v110, v112, v112)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v118 = F_listTypeInitIterator(m, v101, int32(0), int32(1))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L31
	} else {
		goto L37
	}
L37:
	;
	goto L39
L38:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
	if v177 != int32(9) {
		goto L54
	} else {
		goto L55
	}
L39:
	;
	v127 = F_listTypeNext(m, v118, v8+int32(8))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L31
	} else {
		goto L41
	}
L40:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	F_listTypeInsert(m, v8+int32(8), v142, v96)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L31
	} else {
		goto L45
	}
L41:
	;
	if v127 == int32(0) {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v135 = F_listTypeEqual(m, v8+int32(8), v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L31
	} else {
		goto L43
	}
L43:
	;
	if v135 == int32(0) {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
	if v145 != int32(9) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_valkey_free(m, v118)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L31
	} else {
		goto L49
	}
L47:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	F_quicklistReleaseIterator(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L31
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	F_signalModifiedKey(m, l0, v153, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L31
	} else {
		goto L50
	}
L50:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+28))
	F_notifyKeyspaceEvent(m, int32(16), int32(_a2384), v161, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L31
	} else {
		goto L51
	}
L51:
	;
	v166 = int32(_a20)
	v168 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	*(*int64)(unsafe.Add(mBase, _consts[180])) = v168 + int64(1)
	v172 = F_listTypeLength(m, v101)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v172))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L31
	} else {
		goto L53
	}
L53:
	;
	goto L1
L54:
	;
	F_valkey_free(m, v118)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L31
	} else {
		goto L57
	}
L55:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	F_quicklistReleaseIterator(m, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L31
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	F_addReplyLongLong(m, l0, int64(-1))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L31
	} else {
		goto L58
	}
L58:
	;
	goto L1
L59:
	;
	goto L1
}
func F_listenToPort(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v198 int32
	_ = v198
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v14 < int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v198
L2:
	;
	v198 = int32(0)
	goto L1
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v24 = int32(0)
	goto L4
L4:
	;
	v29 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, _consts[751]))
	v32 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v17+v24<<(uint(int32(2))%32))))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v40 = v36 + base.B2i32(v37 == int32(45))
	v41 = int32(58)
	v42 = F___strchrnul(m, v40, v41)
	mBase = m.M
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v44 == v41 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L2
L6:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v61 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0+v60<<(uint(v61)%32)))) = v59
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0+v65<<(uint(v61)%32))))
	if v69 != int32(-1) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v57 = F_anetTcpServer(m, int32(_a2202), v18, v40, v32, v30)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L13
	} else {
		goto L15
	}
L8:
	;
	if v48 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L9:
	;
	v48 = v42
	goto L11
L10:
	;
	v48 = v29
	goto L11
L11:
	;
	goto L8
L12:
	;
	v52 = F_anetTcp6Server(m, int32(_a2202), v18, v40, v32, v30)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v59 = v52
	goto L6
L15:
	;
	v59 = v57
	goto L6
L16:
	;
	v181 = v24 + int32(1)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v181 < v182 {
		v24 = v181
		goto L4
	} else {
		goto L49
	}
L17:
	;
	v104 = int32(0)
	v105 = *(*int32)(unsafe.Add(mBase, _consts[753]))
	if v105 == v104 {
		v116 = v69
		goto L32
	} else {
		goto L33
	}
L18:
	;
	goto L19
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v75 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v75 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v37 != int32(45) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(_a2202)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v40
	F__serverLog(m, int32(3), int32(_a2203), v12)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	switch v73 + int32(-50) {
	case 0, 16:
		goto L16
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L26
	default:
		goto L27
	}
L24:
	;
	if v73 == int32(4) {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v96 = int32(-1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v97 == int32(0) {
		v198 = v96
		goto L1
	} else {
		goto L30
	}
L27:
	;
	switch v73 + int32(-137) {
	case 0, 2:
		goto L16
	case 1:
		goto L26
	default:
		goto L28
	}
L28:
	;
	if v73 == int32(5) {
		goto L16
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+36))
	m.T0[v101].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	v198 = v96
	goto L1
L32:
	;
	v118 = F_anetNonBlock(m, int32(0), v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L13
	} else {
		goto L35
	}
L33:
	;
	v109 = F_anetSetSockMarkId(m, int32(0), v69, v105)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0+v111<<(uint(int32(2))%32))))
	v116 = v115
	goto L32
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0+v120<<(uint(int32(2))%32))))
	v128 = m.G0
	v130 = v128 - int32(16)
	m.G0 = v130
	goto L40
L36:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v174 + int32(1)
	goto L16
L37:
	;
	m.G0 = v130 + int32(16)
	goto L36
L38:
	;
	goto L37
L39:
	;
	if v138&int32(1) != 0 {
		goto L37
	} else {
		goto L44
	}
L40:
	;
	v138 = F_fcntl(m, v124, int32(1), int32(0))
	mBase = m.M
	if v138 != int32(-1) {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v141 = F___errno_location(m)
	mBase = m.M
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v142 == int32(27) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L38
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v138 | int32(1)
	v155 = F_fcntl(m, v124, int32(2), v130)
	mBase = m.M
	if v155 != int32(-1) {
		goto L37
	} else {
		goto L47
	}
L46:
	;
	goto L38
L47:
	;
	v158 = F___errno_location(m)
	mBase = m.M
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v159 == int32(27) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	goto L5
}
func F_llroundl(m *base.Module, l0 int64, l1 int64) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v43 int64
	_ = v43
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	v4 = m.G0
	v5 = int32(16)
	v6 = v4 - v5
	m.G0 = v6
	F_roundl(m, v6, l0, l1)
	mBase = m.M
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v6+int32(8))))
	v16 = m.G0
	v18 = v16 - v5
	m.G0 = v18
	v25 = base.I32_wrap_i64(int64(base.Ui64(v12)>>(uint(int64(48))%64))) & int32(32767)
	if base.Ui32(v25) < base.Ui32(int32(16383)) {
		v49 = int64(0)
	} else {
		if base.Ui32(int32(-65)) < base.Ui32(v25+int32(-16447)) {
			F___lshrti3(m, v18, v9, v12&int64(281474976710655)|int64(281474976710656), int32(16495)-v25)
			mBase = m.M
			v43 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
			if int64(-1) < v12 {
				v48 = v43
			} else {
				v48 = int64(0) - v43
			}
			v49 = v48
		} else {
			v49 = v12>>(uint(int64(63))%64) ^ int64(9223372036854775807)
		}
	}
	m.G0 = v18 + int32(16)
	m.G0 = v6 + int32(16)
	return v49
}
func F_lmoveCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
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
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
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
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	v5 = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v9 = F_objectGetVal(m, v8)
	mBase = m.M
	v10 = int32(_a2386)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v13 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L57
	} else {
		goto L60
	}
L2:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L57
	} else {
		goto L59
	}
L3:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v95 = F_objectGetVal(m, v94)
	mBase = m.M
	v96 = int32(_a2386)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v99 != 0 {
		goto L33
	} else {
		goto L34
	}
L4:
	;
	if v45-v47 == int32(0) {
		v92 = v5
		goto L3
	} else {
		goto L16
	}
L5:
	;
	v45 = F_tolower(m, v41)
	mBase = m.M
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v47 = F_tolower(m, v46)
	mBase = m.M
	goto L4
L6:
	;
	v15 = v9
	v16 = v10
	v17 = v13
	goto L9
L7:
	;
	v41 = int32(0)
	v42 = v10
	goto L5
L8:
	;
	v41 = v38 & int32(255)
	v42 = v37
	goto L5
L9:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v19 == int32(0) {
		v37 = v16
		v38 = v17
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v37 = v31
	v38 = int32(0)
	goto L8
L11:
	;
	v23 = v17 & int32(255)
	if v23 == v19 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v30 = int32(1)
	v31 = v16 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v32 != 0 {
		v15 = v15 + v30
		v16 = v31
		v17 = v32
		goto L9
	} else {
		goto L15
	}
L13:
	;
	v25 = F_tolower(m, v23)
	mBase = m.M
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v27 = F_tolower(m, v26)
	mBase = m.M
	if v25 == v27 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v37 = v16
	v38 = v29
	goto L8
L15:
	;
	goto L10
L16:
	;
	v51 = F_objectGetVal(m, v8)
	mBase = m.M
	v52 = int32(_a2387)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if v87-v89 != 0 {
		goto L2
	} else {
		goto L29
	}
L18:
	;
	v87 = F_tolower(m, v83)
	mBase = m.M
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	v89 = F_tolower(m, v88)
	mBase = m.M
	goto L17
L19:
	;
	v57 = v51
	v58 = v52
	v59 = v55
	goto L22
L20:
	;
	v83 = int32(0)
	v84 = v52
	goto L18
L21:
	;
	v83 = v80 & int32(255)
	v84 = v79
	goto L18
L22:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v61 == int32(0) {
		v79 = v58
		v80 = v59
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v79 = v73
	v80 = int32(0)
	goto L21
L24:
	;
	v65 = v59 & int32(255)
	if v65 == v61 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v72 = int32(1)
	v73 = v58 + v72
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	if v74 != 0 {
		v57 = v57 + v72
		v58 = v73
		v59 = v74
		goto L22
	} else {
		goto L28
	}
L26:
	;
	v67 = F_tolower(m, v65)
	mBase = m.M
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v69 = F_tolower(m, v68)
	mBase = m.M
	if v67 == v69 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v79 = v58
	v80 = v71
	goto L21
L28:
	;
	goto L23
L29:
	;
	v92 = int32(0)
	goto L3
L30:
	;
	F_lmoveGenericCommand(m, l0, v92, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L57
	} else {
		goto L58
	}
L31:
	;
	if v131-v133 == int32(0) {
		v178 = v5
		goto L30
	} else {
		goto L43
	}
L32:
	;
	v131 = F_tolower(m, v127)
	mBase = m.M
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	v133 = F_tolower(m, v132)
	mBase = m.M
	goto L31
L33:
	;
	v101 = v95
	v102 = v96
	v103 = v99
	goto L36
L34:
	;
	v127 = int32(0)
	v128 = v96
	goto L32
L35:
	;
	v127 = v124 & int32(255)
	v128 = v123
	goto L32
L36:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v105 == int32(0) {
		v123 = v102
		v124 = v103
		goto L35
	} else {
		goto L38
	}
L37:
	;
	v123 = v117
	v124 = int32(0)
	goto L35
L38:
	;
	v109 = v103 & int32(255)
	if v109 == v105 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v116 = int32(1)
	v117 = v102 + v116
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v118 != 0 {
		v101 = v101 + v116
		v102 = v117
		v103 = v118
		goto L36
	} else {
		goto L42
	}
L40:
	;
	v111 = F_tolower(m, v109)
	mBase = m.M
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	v113 = F_tolower(m, v112)
	mBase = m.M
	if v111 == v113 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v123 = v102
	v124 = v115
	goto L35
L42:
	;
	goto L37
L43:
	;
	v137 = F_objectGetVal(m, v94)
	mBase = m.M
	v138 = int32(_a2387)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v141 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v173-v175 != 0 {
		goto L1
	} else {
		goto L56
	}
L45:
	;
	v173 = F_tolower(m, v169)
	mBase = m.M
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	v175 = F_tolower(m, v174)
	mBase = m.M
	goto L44
L46:
	;
	v143 = v137
	v144 = v138
	v145 = v141
	goto L49
L47:
	;
	v169 = int32(0)
	v170 = v138
	goto L45
L48:
	;
	v169 = v166 & int32(255)
	v170 = v165
	goto L45
L49:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v147 == int32(0) {
		v165 = v144
		v166 = v145
		goto L48
	} else {
		goto L51
	}
L50:
	;
	v165 = v159
	v166 = int32(0)
	goto L48
L51:
	;
	v151 = v145 & int32(255)
	if v151 == v147 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v158 = int32(1)
	v159 = v144 + v158
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+1)))
	if v160 != 0 {
		v143 = v143 + v158
		v144 = v159
		v145 = v160
		goto L49
	} else {
		goto L55
	}
L53:
	;
	v153 = F_tolower(m, v151)
	mBase = m.M
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v155 = F_tolower(m, v154)
	mBase = m.M
	if v153 == v155 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v165 = v144
	v166 = v157
	goto L48
L55:
	;
	goto L50
L56:
	;
	v178 = int32(0)
	goto L30
L57:
	;
	return
L58:
	;
	return
L59:
	;
	return
L60:
	;
	return
}
func F_lmpopGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v41 int32
	_ = v41
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
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
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
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
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(-1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+l1<<(uint(int32(2))%32))))
	v28 = F_getRangeLongFromObjectOrReply(m, l0, v22, int32(1), int32(2147483647), v12+int32(12), int32(_a2388))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return
L2:
	;
	return
L3:
	;
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v31 = v30 + l1
	v33 = v31 + int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v33 < v34 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v33<<(uint(int32(2))%32))))
	v46 = F_objectGetVal(m, v45)
	mBase = m.M
	v47 = int32(_a2386)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v237)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L2
	} else {
		goto L66
	}
L9:
	;
	v134 = v31
	goto L37
L10:
	;
	if v82-v84 == int32(0) {
		v129 = int32(1)
		goto L9
	} else {
		goto L22
	}
L11:
	;
	v82 = F_tolower(m, v78)
	mBase = m.M
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v84 = F_tolower(m, v83)
	mBase = m.M
	goto L10
L12:
	;
	v52 = v46
	v53 = v47
	v54 = v50
	goto L15
L13:
	;
	v78 = int32(0)
	v79 = v47
	goto L11
L14:
	;
	v78 = v75 & int32(255)
	v79 = v74
	goto L11
L15:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v56 == int32(0) {
		v74 = v53
		v75 = v54
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v74 = v68
	v75 = int32(0)
	goto L14
L17:
	;
	v60 = v54 & int32(255)
	if v60 == v56 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v67 = int32(1)
	v68 = v53 + v67
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v69 != 0 {
		v52 = v52 + v67
		v53 = v68
		v54 = v69
		goto L15
	} else {
		goto L21
	}
L19:
	;
	v62 = F_tolower(m, v60)
	mBase = m.M
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v64 = F_tolower(m, v63)
	mBase = m.M
	if v62 == v64 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v74 = v53
	v75 = v66
	goto L14
L21:
	;
	goto L16
L22:
	;
	v88 = F_objectGetVal(m, v45)
	mBase = m.M
	v89 = int32(_a2387)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	if v124-v126 != 0 {
		goto L8
	} else {
		goto L35
	}
L24:
	;
	v124 = F_tolower(m, v120)
	mBase = m.M
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v126 = F_tolower(m, v125)
	mBase = m.M
	goto L23
L25:
	;
	v94 = v88
	v95 = v89
	v96 = v92
	goto L28
L26:
	;
	v120 = int32(0)
	v121 = v89
	goto L24
L27:
	;
	v120 = v117 & int32(255)
	v121 = v116
	goto L24
L28:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v98 == int32(0) {
		v116 = v95
		v117 = v96
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v116 = v110
	v117 = int32(0)
	goto L27
L30:
	;
	v102 = v96 & int32(255)
	if v102 == v98 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v109 = int32(1)
	v110 = v95 + v109
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	if v111 != 0 {
		v94 = v94 + v109
		v95 = v110
		v96 = v111
		goto L28
	} else {
		goto L34
	}
L32:
	;
	v104 = F_tolower(m, v102)
	mBase = m.M
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v106 = F_tolower(m, v105)
	mBase = m.M
	if v104 == v106 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v116 = v95
	v117 = v108
	goto L27
L34:
	;
	goto L29
L35:
	;
	v129 = int32(0)
	goto L9
L36:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v215 != int32(-1) {
		v221 = v215
		goto L60
	} else {
		goto L61
	}
L37:
	;
	v140 = v134 + int32(2)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v141 <= v140 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143+v140<<(uint(int32(2))%32))))
	v148 = F_objectGetVal(m, v147)
	mBase = m.M
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v149 != int32(-1) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201+v193<<(uint(int32(2))%32))))
	v211 = F_getRangeLongFromObjectOrReply(m, l0, v205, int32(1), int32(2147483647), v12+int32(8), int32(_a2389))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L2
	} else {
		goto L58
	}
L41:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L57
	}
L42:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v153 = int32(_a2390)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v156 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	if v188-v190 != 0 {
		goto L41
	} else {
		goto L55
	}
L44:
	;
	v188 = F_tolower(m, v184)
	mBase = m.M
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	v190 = F_tolower(m, v189)
	mBase = m.M
	goto L43
L45:
	;
	v158 = v148
	v159 = v153
	v160 = v156
	goto L48
L46:
	;
	v184 = int32(0)
	v185 = v153
	goto L44
L47:
	;
	v184 = v181 & int32(255)
	v185 = v180
	goto L44
L48:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v162 == int32(0) {
		v180 = v159
		v181 = v160
		goto L47
	} else {
		goto L50
	}
L49:
	;
	v180 = v174
	v181 = int32(0)
	goto L47
L50:
	;
	v166 = v160 & int32(255)
	if v166 == v162 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v173 = int32(1)
	v174 = v159 + v173
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
	if v175 != 0 {
		v158 = v158 + v173
		v159 = v174
		v160 = v175
		goto L48
	} else {
		goto L54
	}
L52:
	;
	v168 = F_tolower(m, v166)
	mBase = m.M
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v170 = F_tolower(m, v169)
	mBase = m.M
	if v168 == v170 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	v180 = v159
	v181 = v172
	goto L47
L54:
	;
	goto L49
L55:
	;
	v193 = v134 + int32(3)
	if v152 != v193 {
		goto L40
	} else {
		goto L56
	}
L56:
	;
	goto L41
L57:
	;
	goto L1
L58:
	;
	if v211 == int32(0) {
		v134 = v140
		goto L37
	} else {
		goto L59
	}
L59:
	;
	goto L1
L60:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v227 = v222 + l1<<(uint(int32(2))%32) + int32(4)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if l2 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v218 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v218
	v221 = v218
	goto L60
L62:
	;
	F_mpopGenericCommand(m, l0, v227, v228, v129, v221)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L2
	} else {
		goto L65
	}
L63:
	;
	F_blockingPopGenericCommand(m, l0, v227, v228, v129, int32(1), v221)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	goto L1
L65:
	;
	goto L1
L66:
	;
	goto L1
}
func F_loadDataFromDisk(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v47 int64
	_ = v47
	var v52 int64
	_ = v52
	var v57 int64
	_ = v57
	var v62 int64
	_ = v62
	var v67 int64
	_ = v67
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
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
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int64
	_ = v109
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int64
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
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
	var v148 int64
	_ = v148
	var v151 int64
	_ = v151
	var v156 int32
	_ = v156
	var v161 int64
	_ = v161
	var v166 int64
	_ = v166
	var v171 int64
	_ = v171
	var v176 int64
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v208 int64
	_ = v208
	var v215 int32
	_ = v215
	var v220 int64
	_ = v220
	var v225 int64
	_ = v225
	var v230 int64
	_ = v230
	var v235 int64
	_ = v235
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v243 int32
	_ = v243
	var v246 int64
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int64
	_ = v255
	var v267 int32
	_ = v267
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	v9 = F_ustime(m)
	mBase = m.M
	v11 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v11 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F__serverAssert(m, int32(_a2291), int32(_a2157), int32(7276))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L7
	} else {
		goto L49
	}
L2:
	;
	m.G0 = v7 + int32(96)
	return
L3:
	;
	v204 = int32(0)
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v7)+40))
	*(*int64)(unsafe.Add(mBase, _consts[582])) = v205
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v7)+88))
	*(*int64)(unsafe.Add(mBase, _consts[583])) = v208 + int64(1)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(80)))))
	*(*uint8)(unsafe.Add(mBase, _consts[572])) = uint8(v215)
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v7+int32(72))))
	*(*int64)(unsafe.Add(mBase, _consts[574])) = v220
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v7+int32(64))))
	*(*int64)(unsafe.Add(mBase, _consts[576])) = v225
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v7+int32(56))))
	*(*int64)(unsafe.Add(mBase, _consts[578])) = v230
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v7+int32(48))))
	*(*int64)(unsafe.Add(mBase, _consts[580])) = v235
	v239 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v240 = v208 + v239
	*(*int64)(unsafe.Add(mBase, _consts[47])) = v240
	v243 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	if v243 == v204 {
		goto L1
	} else {
		goto L47
	}
L4:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L5:
	;
	v41 = int32(0)
	v42 = *(*int64)(unsafe.Add(mBase, _consts[795]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(88)))) = v42
	v47 = *(*int64)(unsafe.Add(mBase, _consts[796]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(80)))) = v47
	v52 = *(*int64)(unsafe.Add(mBase, _consts[797]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(72)))) = v52
	v57 = *(*int64)(unsafe.Add(mBase, _consts[798]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(64)))) = v57
	v62 = *(*int64)(unsafe.Add(mBase, _consts[799]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(56)))) = v62
	v67 = *(*int64)(unsafe.Add(mBase, _consts[800]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v67
	v72 = *(*int64)(unsafe.Add(mBase, _consts[801]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v72
	v75 = *(*int64)(unsafe.Add(mBase, _consts[802]))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v75
	goto L13
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v16 = F_loadAppendOnlyFiles(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if base.Ui32(v16+int32(-3)) <= base.Ui32(int32(1)) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v16 == int32(1) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v25 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v28 = F_ustime(m)
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(v7))) = base.F64_promote_f32(base.F32_div(base.F32_convert_i64_s(v28-v9), float32(1e+06)))
	F__serverLog(m, int32(2), int32(_a2294), v7)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L2
L13:
	;
	v78 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[18])) = v78
	v81 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v81 == v78 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	v103 = F_rdbLoad(m, v100, v7+int32(32), v98)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L7
	} else {
		goto L26
	}
L15:
	;
	F_createReplicationBacklog(m)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L22
	}
L16:
	;
	v91 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v93 != 0 {
		v98 = v91
		goto L14
	} else {
		goto L21
	}
L17:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	goto L18
L18:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+88))
	goto L19
L19:
	;
	if v87&int32(1) != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v98 = int32(0)
	goto L14
L21:
	;
	goto L15
L22:
	;
	v98 = int32(8)
	goto L14
L23:
	;
	v196 = int32(0)
	v197 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	if v197 == v196 {
		goto L2
	} else {
		goto L45
	}
L24:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v186 {
		goto L42
	} else {
		goto L43
	}
L25:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v106 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	switch v103 {
	case 0:
		goto L25
	case 1:
		goto L23
	default:
		goto L24
	}
L27:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
	if v122 == int32(0) {
		goto L23
	} else {
		goto L30
	}
L28:
	;
	v109 = F_ustime(m)
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(v7)+16)) = base.F64_promote_f32(base.F32_div(base.F32_convert_i64_s(v109-v9), float32(1e+06)))
	F__serverLog(m, int32(2), int32(_a2292), v7+int32(16))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v7)+88))
	if v125 == int64(-1) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	if v128 == int32(-1) {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v131 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v132 == v131 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v147 = int32(0)
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v7)+40))
	*(*int64)(unsafe.Add(mBase, _consts[581])) = v148
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v7)+88))
	*(*int64)(unsafe.Add(mBase, _consts[47])) = v151
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(80)))))
	*(*uint8)(unsafe.Add(mBase, _consts[571])) = uint8(v156)
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v7+int32(72))))
	*(*int64)(unsafe.Add(mBase, _consts[573])) = v161
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v7+int32(64))))
	*(*int64)(unsafe.Add(mBase, _consts[575])) = v166
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v7+int32(56))))
	*(*int64)(unsafe.Add(mBase, _consts[577])) = v171
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v7+int32(48))))
	*(*int64)(unsafe.Add(mBase, _consts[579])) = v176
	F_replicationCachePrimaryUsingMyself(m)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L7
	} else {
		goto L40
	}
L34:
	;
	v143 = int32(0)
	v144 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v144 == v143 {
		goto L3
	} else {
		goto L39
	}
L35:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	goto L36
L36:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+88))
	goto L37
L37:
	;
	if v138&int32(1) == int32(0) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L3
L39:
	;
	goto L33
L40:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	v183 = F_selectDb(m, v181, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	goto L2
L42:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F__serverLog(m, int32(3), int32(_a2293), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	F_freeReplicationBacklog(m)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	goto L2
L47:
	;
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v243)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v243)+24)) = v240 - v246 + int64(1)
	F_rebaseReplicationBuffer(m, v208)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v253 = int32(0)
	v255 = F___time(m, v253)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[585])) = v255
	goto L2
L49:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_loadingIncrProgress(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
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
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	v2 = int32(_a20)
	v4 = *(*int64)(unsafe.Add(mBase, _consts[553]))
	*(*int64)(unsafe.Add(mBase, _consts[553])) = v4 + l0
	v8 = *(*int32)(unsafe.Add(mBase, _consts[554]))
	v9 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	if v18 < int32(261) {
		if v18 < int32(1) {
			v95 = v9
		} else {
			v26 = v9
			v27 = v18
			v29 = v27 & int32(3)
			if base.Ui32(int32(4)) <= base.Ui32(v27) {
				v36 = int32(0)
				v38 = v26
				v39 = v36
				v43 = v36
				for {
					v46 = v39 << (uint(int32(2)) % 32)
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[317])))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[318])))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[319])))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[320])))
					v62 = v49 + (v52 + (v55 + (v58 + v38)))
					v63 = int32(4)
					v64 = v39 + v63
					v66 = v43 + v63
					if v66 != v27&int32(2147483644) {
						v38 = v62
						v39 = v64
						v43 = v66
						continue
					} else {
						break
					}
					break
				}
				v68 = v62
				v69 = v64
			} else {
				v68 = v26
				v69 = int32(0)
			}
			if v29 == int32(0) {
				v95 = v68
			} else {
				v77 = v68
				v78 = v69
				v80 = int32(0)
				for {
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(int32(2))%32))+uint32(_consts[320])))
					v89 = v88 + v77
					v90 = int32(1)
					v93 = v80 + v90
					if v93 != v29 {
						v77 = v89
						v78 = v78 + v90
						v80 = v93
						continue
					} else {
						break
					}
					break
				}
				v95 = v89
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _consts[316]))
		v26 = v22
		v27 = int32(260)
		v29 = v27 & int32(3)
		if base.Ui32(int32(4)) <= base.Ui32(v27) {
			v36 = int32(0)
			v38 = v26
			v39 = v36
			v43 = v36
			for {
				v46 = v39 << (uint(int32(2)) % 32)
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[317])))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[318])))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[319])))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[320])))
				v62 = v49 + (v52 + (v55 + (v58 + v38)))
				v63 = int32(4)
				v64 = v39 + v63
				v66 = v43 + v63
				if v66 != v27&int32(2147483644) {
					v38 = v62
					v39 = v64
					v43 = v66
					continue
				} else {
					break
				}
				break
			}
			v68 = v62
			v69 = v64
		} else {
			v68 = v26
			v69 = int32(0)
		}
		if v29 == int32(0) {
			v95 = v68
		} else {
			v77 = v68
			v78 = v69
			v80 = int32(0)
			for {
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(int32(2))%32))+uint32(_consts[320])))
				v89 = v88 + v77
				v90 = int32(1)
				v93 = v80 + v90
				if v93 != v29 {
					v77 = v89
					v78 = v78 + v90
					v80 = v93
					continue
				} else {
					break
				}
				break
			}
			v95 = v89
		}
	}
	if base.Ui32(v95) <= base.Ui32(v8) {
	} else {
		v104 = int32(0)
		v113 = *(*int32)(unsafe.Add(mBase, _consts[315]))
		if v113 < int32(261) {
			if v113 < int32(1) {
				v190 = v104
			} else {
				v121 = v104
				v122 = v113
				v124 = v122 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(v122) {
					v131 = int32(0)
					v133 = v121
					v134 = v131
					v138 = v131
					for {
						v141 = v134 << (uint(int32(2)) % 32)
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[317])))
						v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[318])))
						v150 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[319])))
						v153 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[320])))
						v157 = v144 + (v147 + (v150 + (v153 + v133)))
						v158 = int32(4)
						v159 = v134 + v158
						v161 = v138 + v158
						if v161 != v122&int32(2147483644) {
							v133 = v157
							v134 = v159
							v138 = v161
							continue
						} else {
							break
						}
						break
					}
					v163 = v157
					v164 = v159
				} else {
					v163 = v121
					v164 = int32(0)
				}
				if v124 == int32(0) {
					v190 = v163
				} else {
					v172 = v163
					v173 = v164
					v175 = int32(0)
					for {
						v183 = *(*int32)(unsafe.Add(mBase, uint32(v173<<(uint(int32(2))%32))+uint32(_consts[320])))
						v184 = v183 + v172
						v185 = int32(1)
						v188 = v175 + v185
						if v188 != v124 {
							v172 = v184
							v173 = v173 + v185
							v175 = v188
							continue
						} else {
							break
						}
						break
					}
					v190 = v184
				}
			}
		} else {
			v117 = *(*int32)(unsafe.Add(mBase, _consts[316]))
			v121 = v117
			v122 = int32(260)
			v124 = v122 & int32(3)
			if base.Ui32(int32(4)) <= base.Ui32(v122) {
				v131 = int32(0)
				v133 = v121
				v134 = v131
				v138 = v131
				for {
					v141 = v134 << (uint(int32(2)) % 32)
					v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[317])))
					v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[318])))
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[319])))
					v153 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[320])))
					v157 = v144 + (v147 + (v150 + (v153 + v133)))
					v158 = int32(4)
					v159 = v134 + v158
					v161 = v138 + v158
					if v161 != v122&int32(2147483644) {
						v133 = v157
						v134 = v159
						v138 = v161
						continue
					} else {
						break
					}
					break
				}
				v163 = v157
				v164 = v159
			} else {
				v163 = v121
				v164 = int32(0)
			}
			if v124 == int32(0) {
				v190 = v163
			} else {
				v172 = v163
				v173 = v164
				v175 = int32(0)
				for {
					v183 = *(*int32)(unsafe.Add(mBase, uint32(v173<<(uint(int32(2))%32))+uint32(_consts[320])))
					v184 = v183 + v172
					v185 = int32(1)
					v188 = v175 + v185
					if v188 != v124 {
						v172 = v184
						v173 = v173 + v185
						v175 = v188
						continue
					} else {
						break
					}
					break
				}
				v190 = v184
			}
		}
		*(*int32)(unsafe.Add(mBase, _consts[554])) = v190
	}
	return
}
func F_locking_putc_1(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	v1 = l0
	v7 = l1 + int32(76)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v9 != 0 {
		v11 = v9
	} else {
		v11 = int32(1073741823)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v11
	if v9 == int32(0) {
	} else {
	}
	v17 = v1 & int32(255)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v17 == v18 {
		v28 = F___overflow(m, l1, v17)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = v28
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
			if v35&int32(1073741824) == int32(0) {
			} else {
				v43 = F_emscripten_futex_wake(m, v7, int32(1))
				mBase = m.M
			}
			return v32
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		if v20 == v21 {
			v28 = F___overflow(m, l1, v17)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = v28
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
				if v35&int32(1073741824) == int32(0) {
				} else {
					v43 = F_emscripten_futex_wake(m, v7, int32(1))
					mBase = m.M
				}
				return v32
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v20 + int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v1)
			v32 = v17
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
			if v35&int32(1073741824) == int32(0) {
			} else {
				v43 = F_emscripten_futex_wake(m, v7, int32(1))
				mBase = m.M
			}
			return v32
		}
	}
}
func F_logInvalidUseAndFreeClientAsync(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v15 int32
	_ = v15
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v10 = F_sdsempty(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		v13 = F_sdscatvprintf(m, v10, l1, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = F_sdsempty(m)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _consts[481]))
				v19 = F_catClientInfoString(m, v15, l0, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, _consts[28]))
					if int32(3) < v22 {
						F_sdsfree(m, v13)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							F_sdsfree(m, v19)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								if v35&int32(1280) != 0 {
									m.G0 = v7 + int32(16)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v35 | int32(1024)
									v42 = *(*int32)(unsafe.Add(mBase, _consts[113]))
									if v42 == int32(0) {
										v50 = *(*int32)(unsafe.Add(mBase, _consts[483]))
										v51 = F_listAddNodeTail(m, v50, l0)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, _consts[483]))
										v47 = F_listSearchKey(m, v46, l0)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											if v47 != 0 {
												F__serverAssertWithInfo(m, l0, int32(0), int32(_a1641), int32(_a1630), int32(2277))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v50 = *(*int32)(unsafe.Add(mBase, _consts[483]))
												v51 = F_listAddNodeTail(m, v50, l0)
												mBase = m.M
												v52 = m.ExcPending
												if v52 != 0 {
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
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v19
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
						F__serverLog(m, int32(3), int32(_a1643), v7)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							F_sdsfree(m, v13)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_sdsfree(m, v19)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
									if v35&int32(1280) != 0 {
										m.G0 = v7 + int32(16)
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v35 | int32(1024)
										v42 = *(*int32)(unsafe.Add(mBase, _consts[113]))
										if v42 == int32(0) {
											v50 = *(*int32)(unsafe.Add(mBase, _consts[483]))
											v51 = F_listAddNodeTail(m, v50, l0)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return
											} else {
												m.G0 = v7 + int32(16)
												return
											}
										} else {
											v46 = *(*int32)(unsafe.Add(mBase, _consts[483]))
											v47 = F_listSearchKey(m, v46, l0)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return
											} else {
												if v47 != 0 {
													F__serverAssertWithInfo(m, l0, int32(0), int32(_a1641), int32(_a1630), int32(2277))
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v50 = *(*int32)(unsafe.Add(mBase, _consts[483]))
													v51 = F_listAddNodeTail(m, v50, l0)
													mBase = m.M
													v52 = m.ExcPending
													if v52 != 0 {
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
							}
						}
					}
				}
			}
		}
	}
}
func F_log_inline(m *base.Module, l0 int64, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int64
	_ = v12
	var v16 float64
	_ = v16
	var v18 float64
	_ = v18
	var v26 int32
	_ = v26
	var v29 float64
	_ = v29
	var v33 int64
	_ = v33
	var v38 float64
	_ = v38
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v50 float64
	_ = v50
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v72 float64
	_ = v72
	var v75 float64
	_ = v75
	var v78 float64
	_ = v78
	var v82 float64
	_ = v82
	var v85 float64
	_ = v85
	var v90 float64
	_ = v90
	var v93 float64
	_ = v93
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	v5 = int32(0)
	v12 = l0 + int64(-4604531861337669632)
	v16 = base.F64_convert_i32_s(base.I32_wrap_i64(v12 >> (uint(int64(52)) % 64)))
	v18 = *(*float64)(unsafe.Add(mBase, _consts[1019]))
	v26 = base.I32_wrap_i64(int64(base.Ui64(v12)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
	v29 = *(*float64)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1020])))
	v33 = l0 - v12&int64(-4503599627370496)
	v38 = base.F64_reinterpret_i64((v33 + int64(2147483648)) & int64(-4294967296))
	v41 = *(*float64)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1021])))
	v44 = base.F64_add(base.F64_mul(v38, v41), float64(-1))
	v47 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v33), v38), v41)
	v48 = base.F64_add(v44, v47)
	v50 = *(*float64)(unsafe.Add(mBase, _consts[1022]))
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1023])))
	v55 = base.F64_add(base.F64_mul(v16, v50), v54)
	v56 = base.F64_add(v48, v55)
	v61 = *(*float64)(unsafe.Add(mBase, _consts[1024]))
	v62 = base.F64_mul(v48, v61)
	v63 = base.F64_mul(v44, v61)
	v67 = base.F64_mul(v44, v63)
	v68 = base.F64_add(v56, v67)
	v72 = base.F64_mul(v48, v62)
	v75 = *(*float64)(unsafe.Add(mBase, _consts[1025]))
	v78 = *(*float64)(unsafe.Add(mBase, _consts[1026]))
	v82 = *(*float64)(unsafe.Add(mBase, _consts[1027]))
	v85 = *(*float64)(unsafe.Add(mBase, _consts[1028]))
	v90 = *(*float64)(unsafe.Add(mBase, _consts[1029]))
	v93 = *(*float64)(unsafe.Add(mBase, _consts[1030]))
	v97 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v16, v18), v29), base.F64_add(v48, base.F64_sub(v55, v56))), base.F64_mul(v47, base.F64_add(v62, v63))), base.F64_add(v67, base.F64_sub(v56, v68))), base.F64_mul(base.F64_mul(v48, v72), base.F64_add(base.F64_mul(v72, base.F64_add(base.F64_mul(v72, base.F64_add(base.F64_mul(v48, v75), v78)), base.F64_add(base.F64_mul(v48, v82), v85))), base.F64_add(base.F64_mul(v48, v90), v93))))
	v98 = base.F64_add(v68, v97)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_add(v97, base.F64_sub(v68, v98))
	return v98
}
func F_lookupSubcommand(m *base.Module, l0 int32, l1 int32) int32 {
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
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v13 = F_hashtableFind(m, v10, l1, v6+int32(12))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		m.G0 = v6 + int32(16)
		return v17
	}
}
func F_lpopCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_popGenericCommand(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_lrulfu_updateClockAndPolicy(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int64
	_ = v7
	var v11 int64
	_ = v11
	v2 = l1
	v3 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[378])) = uint8(v2)
	v7 = base.I64_div_s(l0, int64(60000))
	*(*uint16)(unsafe.Add(mBase, _consts[376])) = uint16(v7)
	v11 = base.I64_div_s(l0, int64(1000))
	*(*int32)(unsafe.Add(mBase, _consts[379])) = base.I32_wrap_i64(v11) & int32(16777215)
	return
}
func F_luaA_pushobject(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v4))) = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9 + int32(16)
	return
}
func F_luaE_freethread(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	F_luaF_close(m, l1, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
		v11 = F_luaM_realloc_(m, l0, v6, v7*int32(24), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
			v18 = F_luaM_realloc_(m, l0, v13, v14<<(uint(int32(4))%32), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v22 = F_luaM_realloc_(m, l0, l1, int32(120), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_luaG_aritherror(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	v8 = m.G0
	v10 = v8 - int32(64)
	m.G0 = v10
	v14 = F_luaV_tonumber(m, l1, v10+int32(40))
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = int32(0)
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v18 = l2
	goto L5
L4:
	;
	v18 = l1
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v20 = m.G399
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v19<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if base.Ui32(v27) <= base.Ui32(v26) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	m.G0 = v10 + int32(64)
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v24
	v71 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v71 + int32(_a2655)
	F_luaG_runerror(m, l0, v71+int32(_a2656), v10)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L16
	}
L8:
	;
	v30 = v26
	goto L10
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v46 = F_getobjname(m, l0, v25, (v18-v40)>>(uint(int32(4))%32), v10+int32(60))
	mBase = m.M
	if v46 == int32(0) {
		goto L7
	} else {
		goto L14
	}
L10:
	;
	if v18 == v30 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v38 = v30 + int32(16)
	if base.Ui32(v27) <= base.Ui32(v38) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v30 = v38
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v46
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v24
	v53 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v53 + int32(_a2655)
	F_luaG_runerror(m, l0, v53+int32(_a2657), v10+int32(16))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L6
L16:
	;
	goto L6
}
func F_luaG_errormsg(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v4 == int32(0) {
		F_luaD_throw(m, l0, int32(2))
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			return
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v8 = v7 + v4
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		if v9 == int32(6) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v16 = int32(-16)
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v15+v16)))
			*(*int64)(unsafe.Add(mBase, uint32(v15))) = v18
			v20 = int32(-8)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v15+v20)))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v22
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
			*(*int64)(unsafe.Add(mBase, uint32(v24+v16))) = v27
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v24+v20))) = v31
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if int32(16) < v33-v34 {
				v42 = v34
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v42 + int32(16)
				F_luaD_call(m, l0, v42+int32(-16), int32(1))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					F_luaD_throw(m, l0, int32(2))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				F_luaD_growstack(m, l0, int32(1))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v42 = v41
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v42 + int32(16)
					F_luaD_call(m, l0, v42+int32(-16), int32(1))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						F_luaD_throw(m, l0, int32(2))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			F_luaD_throw(m, l0, int32(5))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v16 = int32(-16)
				v18 = *(*int64)(unsafe.Add(mBase, uint32(v15+v16)))
				*(*int64)(unsafe.Add(mBase, uint32(v15))) = v18
				v20 = int32(-8)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v15+v20)))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v22
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
				*(*int64)(unsafe.Add(mBase, uint32(v24+v16))) = v27
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v24+v20))) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if int32(16) < v33-v34 {
					v42 = v34
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v42 + int32(16)
					F_luaD_call(m, l0, v42+int32(-16), int32(1))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						F_luaD_throw(m, l0, int32(2))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					F_luaD_growstack(m, l0, int32(1))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v42 = v41
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v42 + int32(16)
						F_luaD_call(m, l0, v42+int32(-16), int32(1))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							F_luaD_throw(m, l0, int32(2))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
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
func F_luaG_runerror(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
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
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
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
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
	v12 = F_luaO_pushvfstring(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
		if v16 != int32(6) {
			F_luaG_errormsg(m, l0)
			mBase = m.M
			v134 = m.ExcPending
			if v134 != 0 {
				return
			} else {
				m.G0 = v9 + int32(80)
				return
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+6)))
			if v20 != 0 {
				F_luaG_errormsg(m, l0)
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return
				} else {
					m.G0 = v9 + int32(80)
					return
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v21
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
				v26 = v21 - v25
				if int32(4) <= v26 {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
					if v30 != 0 {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v30+v26+int32(-4))))
						v36 = v35
					} else {
						v36 = int32(0)
					}
				} else {
					v36 = int32(-1)
				}
				v38 = int32(16)
				v39 = v9 + v38
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
				v42 = v40 + v38
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
				switch v46 + int32(-61) {
				case 0:
					v51 = F_strncpy(m, v39, v40+int32(17), int32(60))
					mBase = m.M
					v55 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v51+int32(59)))) = uint8(v55)
				default:
					v73 = m.G3
					v76 = F_strcspn(m, v42, v73+int32(_a2653))
					mBase = m.M
					v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[988]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v9+int32(24)))) = uint16(v83)
					v85 = *(*int64)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[989])))
					*(*int64)(unsafe.Add(mBase, uint32(v39))) = v85
					v88 = int32(43)
					if base.Ui32(v76) < base.Ui32(v88) {
						v90 = v76
					} else {
						v90 = v88
					}
					v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v90))))
					if v92 == int32(0) {
						v100 = F_strcat(m, v39, v42)
						mBase = m.M
					} else {
						v95 = F_strncat(m, v39, v42, v90)
						mBase = m.M
						v96 = F_strlen(m, v95)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v95+v96))) = int32(3026478)
					}
					v102 = F_strlen(m, v39)
					mBase = m.M
					v103 = v39 + v102
					v104 = m.G3
					v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+uint32(_consts[990]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v103))) = uint16(v107)
					v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+uint32(_consts[991]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v103+int32(2)))) = uint8(v113)
				case 3:
					v58 = v40 + int32(17)
					v59 = F_strlen(m, v58)
					mBase = m.M
					v60 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v60)
					v63 = int32(52)
					if base.Ui32(v59) <= base.Ui32(v63) {
						v71 = v58
					} else {
						v65 = F_strlen(m, v39)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v39+v65))) = int32(3026478)
						v71 = v58 + (v59 - v63)
					}
					v72 = F_strcat(m, v39, v71)
					mBase = m.M
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v36
				v122 = m.G3
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(16)
				v128 = F_luaO_pushfstring(m, l0, v122+int32(_a2654), v9)
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return
				} else {
					F_luaG_errormsg(m, l0)
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						m.G0 = v9 + int32(80)
						return
					}
				}
			}
		}
	}
}
func F_luaM_toobig(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	v2 = m.G3
	F_luaG_runerror(m, l0, v2+int32(_a2664), int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_luaS_resize(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
	if v11 == int32(2) {
		return
	} else {
		if base.Ui32(int32(1073741823)) < base.Ui32(l1+int32(1)) {
			v24 = F_luaM_toobig(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = v24
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if l1 < int32(1) {
				} else {
					v34 = F__emscripten_memset_bulkmem(m, v26, base.I32_extend8_s(int32(0)), l1<<(uint(int32(2))%32))
					mBase = m.M
				}
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				if v35 < int32(1) {
					v92 = v35
				} else {
					v45 = v35
					v47 = int32(0)
					for {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v47<<(uint(int32(2))%32))))
						if v54 == int32(0) {
							v80 = v45
						} else {
							v64 = v54
							for {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
								v71 = v26 + v67&(l1+int32(-1))<<(uint(int32(2))%32)
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
								*(*int32)(unsafe.Add(mBase, uint32(v64))) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v71))) = v64
								if v66 != 0 {
									v64 = v66
									continue
								} else {
									break
								}
								break
							}
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
							v80 = v75
						}
						v86 = v47 + int32(1)
						if v86 < v80 {
							v45 = v80
							v47 = v86
							continue
						} else {
							break
						}
						break
					}
					v92 = v80
				}
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v101 = F_luaM_realloc_(m, l0, v97, v92<<(uint(int32(2))%32), int32(0))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = l1
					return
				}
			}
		} else {
			v18 = int32(0)
			v22 = F_luaM_realloc_(m, l0, v18, v18, l1<<(uint(int32(2))%32))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v26 = v22
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if l1 < int32(1) {
				} else {
					v34 = F__emscripten_memset_bulkmem(m, v26, base.I32_extend8_s(int32(0)), l1<<(uint(int32(2))%32))
					mBase = m.M
				}
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				if v35 < int32(1) {
					v92 = v35
				} else {
					v45 = v35
					v47 = int32(0)
					for {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v47<<(uint(int32(2))%32))))
						if v54 == int32(0) {
							v80 = v45
						} else {
							v64 = v54
							for {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
								v71 = v26 + v67&(l1+int32(-1))<<(uint(int32(2))%32)
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
								*(*int32)(unsafe.Add(mBase, uint32(v64))) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v71))) = v64
								if v66 != 0 {
									v64 = v66
									continue
								} else {
									break
								}
								break
							}
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
							v80 = v75
						}
						v86 = v47 + int32(1)
						if v86 < v80 {
							v45 = v80
							v47 = v86
							continue
						} else {
							break
						}
						break
					}
					v92 = v80
				}
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v101 = F_luaM_realloc_(m, l0, v97, v92<<(uint(int32(2))%32), int32(0))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = l1
					return
				}
			}
		}
	}
}
func F_luaT_gettm(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v7 = int32(-1)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v16 = v5 + v6&(v7<<(uint(v8)%32)^v7)<<(uint(int32(5))%32)
	goto L3
L1:
	;
	return v36
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v29 != 0 {
		v36 = v28
		goto L1
	} else {
		goto L9
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v19 != int32(4) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v28 = v25
	goto L2
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v25 = m.G398
	if v24 != 0 {
		v16 = v24
		goto L3
	} else {
		goto L8
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v22 != l2 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v28 = v16
	goto L2
L8:
	;
	goto L4
L9:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	v33 = v30 | int32(1)<<(uint(l1)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v33)
	v36 = int32(0)
	goto L1
}
func F_luaY_parser(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	v6 = m.G0
	v8 = v6 - int32(656)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+644)) = l2
	if l3&int32(3) == int32(0) {
		v32 = l3
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v66 = F_luaS_newlstr(m, l0, l3, v65)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v65 = v57 - l3
	goto L1
L3:
	;
	v36 = v32
	goto L11
L4:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = l3
	goto L7
L6:
	;
	v65 = l3 - l3
	goto L1
L7:
	;
	v25 = v21 + int32(1)
	if v25&int32(3) == int32(0) {
		v32 = v25
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v30 != 0 {
		v21 = v25
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v57 = v25
	goto L2
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v45 = int32(-2139062144)
	if (int32(16843008)-v42|v42)&v45 == v45 {
		v36 = v36 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v51 = v36
	goto L14
L13:
	;
	goto L12
L14:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != 0 {
		v51 = v51 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v57 = v51
	goto L2
L16:
	;
	goto L15
L17:
	;
	return int32(0)
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v66
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(16) < v74-v75 {
		v83 = v75
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v83 + int32(16)
	F_luaX_setinput(m, l0, v8+int32(584), l1, v66)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	F_luaD_growstack(m, l0, int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v83 = v82
	goto L19
L22:
	;
	F_open_func(m, v8+int32(584), v8+int32(12))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v98 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+74)) = uint8(v98)
	F_luaX_next(m, v8+int32(584))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	F_chunk(m, v8+int32(584))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v8)+600))
	if v108 == int32(287) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_close_func(m, v8+int32(584))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L17
	} else {
		goto L31
	}
L27:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v8)+636))
	v115 = F_luaX_token2str(m, v8+int32(584), int32(287))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v115
	v120 = m.G3
	v123 = F_luaO_pushfstring(m, v111, v120+int32(_a2677), v8)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	F_luaX_syntaxerror(m, v8+int32(584), v123)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v132 + int32(-16)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	m.G0 = v8 + int32(656)
	return v136
}
func F_luaZ_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
	return
}
func F_luaZ_lookahead(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v18 = m.T0[v17].(func(*base.Module, int32, int32, int32) int32)(m, v13, v14, v7+int32(12))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v18 == int32(0) {
				v32 = int32(-1)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				if v24 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
					v29 = v18
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
					v32 = v31
				} else {
					v32 = int32(-1)
				}
			}
			m.G0 = v7 + int32(16)
			return v32
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v29 = v12
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
		v32 = v31
		m.G0 = v7 + int32(16)
		return v32
	}
}
func F_luaZ_openspace(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v6) < base.Ui32(l2) {
		v10 = int32(32)
		if base.Ui32(v10) < base.Ui32(l2) {
			v13 = l2
		} else {
			v13 = v10
		}
		if base.Ui32(int32(-3)) < base.Ui32(v13+int32(1)) {
			v23 = F_luaM_toobig(m, l0)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = v23
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v13
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25
				return v25
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v19 = F_luaM_realloc_(m, l0, v18, v6, v13)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v25 = v19
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v13
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25
				return v25
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		return v8
	}
}
func F_lwGetPixel(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v4 = int32(0)
	if l1 < v4 {
		v21 = v4
	} else {
		v9 = int32(0)
		if l2 < v9 {
			v21 = v9
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v12 <= l1 {
				v21 = v9
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v14 <= l2 {
					v21 = v9
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v20 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16+l1+v12*l2))))
					v21 = v20
				}
			}
		}
	}
	return v21
}
func F_lzf_decompress(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
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
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v517 int32
	_ = v517
	if l1 == int32(0) {
		v517 = l2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v517 - l2
L2:
	;
	v13 = l2 + l3
	v14 = l0 + l1
	v15 = l0
	v19 = l2
	goto L3
L3:
	;
	v26 = v15 + int32(1)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if base.Ui32(int32(31)) < base.Ui32(v27) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v517 = v506
	goto L1
L5:
	;
	if base.Ui32(v502) < base.Ui32(v14) {
		v15 = v502
		v19 = v506
		goto L3
	} else {
		goto L84
	}
L6:
	;
	if base.Ui32(v26) < base.Ui32(v14) {
		goto L46
	} else {
		goto L47
	}
L7:
	;
	v31 = v27 + int32(1)
	if base.Ui32(v19+v31) <= base.Ui32(v13) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if base.Ui32(v26+v31) <= base.Ui32(v14) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(1)
	return int32(0)
L11:
	;
	switch v27 {
	default:
		v292 = v26
		v293 = v19
		goto L14
	case 1:
		v284 = v26
		v285 = v19
		goto L15
	case 2:
		v276 = v26
		v277 = v19
		goto L16
	case 3:
		v268 = v26
		v269 = v19
		goto L17
	case 4:
		v260 = v26
		v261 = v19
		goto L18
	case 5:
		v252 = v26
		v253 = v19
		goto L19
	case 6:
		v244 = v26
		v245 = v19
		goto L20
	case 7:
		v236 = v26
		v237 = v19
		goto L21
	case 8:
		v228 = v26
		v229 = v19
		goto L22
	case 9:
		v220 = v26
		v221 = v19
		goto L23
	case 10:
		v212 = v26
		v213 = v19
		goto L24
	case 11:
		v204 = v26
		v205 = v19
		goto L25
	case 12:
		v196 = v26
		v197 = v19
		goto L26
	case 13:
		v188 = v26
		v189 = v19
		goto L27
	case 14:
		v180 = v26
		v181 = v19
		goto L28
	case 15:
		v172 = v26
		v173 = v19
		goto L29
	case 16:
		v164 = v26
		v165 = v19
		goto L30
	case 17:
		v156 = v26
		v157 = v19
		goto L31
	case 18:
		v148 = v26
		v149 = v19
		goto L32
	case 19:
		v140 = v26
		v141 = v19
		goto L33
	case 20:
		v132 = v26
		v133 = v19
		goto L34
	case 21:
		v124 = v26
		v125 = v19
		goto L35
	case 22:
		v116 = v26
		v117 = v19
		goto L36
	case 23:
		v108 = v26
		v109 = v19
		goto L37
	case 24:
		v100 = v26
		v101 = v19
		goto L38
	case 25:
		v92 = v26
		v93 = v19
		goto L39
	case 26:
		v84 = v26
		v85 = v19
		goto L40
	case 27:
		v76 = v26
		v77 = v19
		goto L41
	case 28:
		v68 = v26
		v69 = v19
		goto L42
	case 29:
		v60 = v26
		v61 = v19
		goto L43
	case 30:
		v52 = v26
		v53 = v19
		goto L44
	case 31:
		goto L45
	}
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(28)
	return int32(0)
L14:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	*(*uint8)(unsafe.Add(mBase, uint32(v293))) = uint8(v294)
	v296 = int32(1)
	v502 = v292 + v296
	v506 = v293 + v296
	goto L5
L15:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	*(*uint8)(unsafe.Add(mBase, uint32(v285))) = uint8(v286)
	v288 = int32(1)
	v292 = v284 + v288
	v293 = v285 + v288
	goto L14
L16:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	*(*uint8)(unsafe.Add(mBase, uint32(v277))) = uint8(v278)
	v280 = int32(1)
	v284 = v276 + v280
	v285 = v277 + v280
	goto L15
L17:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v270)
	v272 = int32(1)
	v276 = v268 + v272
	v277 = v269 + v272
	goto L16
L18:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v262)
	v264 = int32(1)
	v268 = v260 + v264
	v269 = v261 + v264
	goto L17
L19:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	*(*uint8)(unsafe.Add(mBase, uint32(v253))) = uint8(v254)
	v256 = int32(1)
	v260 = v252 + v256
	v261 = v253 + v256
	goto L18
L20:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	*(*uint8)(unsafe.Add(mBase, uint32(v245))) = uint8(v246)
	v248 = int32(1)
	v252 = v244 + v248
	v253 = v245 + v248
	goto L19
L21:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	*(*uint8)(unsafe.Add(mBase, uint32(v237))) = uint8(v238)
	v240 = int32(1)
	v244 = v236 + v240
	v245 = v237 + v240
	goto L20
L22:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	*(*uint8)(unsafe.Add(mBase, uint32(v229))) = uint8(v230)
	v232 = int32(1)
	v236 = v228 + v232
	v237 = v229 + v232
	goto L21
L23:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	*(*uint8)(unsafe.Add(mBase, uint32(v221))) = uint8(v222)
	v224 = int32(1)
	v228 = v220 + v224
	v229 = v221 + v224
	goto L22
L24:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v214)
	v216 = int32(1)
	v220 = v212 + v216
	v221 = v213 + v216
	goto L23
L25:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	*(*uint8)(unsafe.Add(mBase, uint32(v205))) = uint8(v206)
	v208 = int32(1)
	v212 = v204 + v208
	v213 = v205 + v208
	goto L24
L26:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v198)
	v200 = int32(1)
	v204 = v196 + v200
	v205 = v197 + v200
	goto L25
L27:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	*(*uint8)(unsafe.Add(mBase, uint32(v189))) = uint8(v190)
	v192 = int32(1)
	v196 = v188 + v192
	v197 = v189 + v192
	goto L26
L28:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	*(*uint8)(unsafe.Add(mBase, uint32(v181))) = uint8(v182)
	v184 = int32(1)
	v188 = v180 + v184
	v189 = v181 + v184
	goto L27
L29:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v174)
	v176 = int32(1)
	v180 = v172 + v176
	v181 = v173 + v176
	goto L28
L30:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	*(*uint8)(unsafe.Add(mBase, uint32(v165))) = uint8(v166)
	v168 = int32(1)
	v172 = v164 + v168
	v173 = v165 + v168
	goto L29
L31:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v158)
	v160 = int32(1)
	v164 = v156 + v160
	v165 = v157 + v160
	goto L30
L32:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v150)
	v152 = int32(1)
	v156 = v148 + v152
	v157 = v149 + v152
	goto L31
L33:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v142)
	v144 = int32(1)
	v148 = v140 + v144
	v149 = v141 + v144
	goto L32
L34:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v134)
	v136 = int32(1)
	v140 = v132 + v136
	v141 = v133 + v136
	goto L33
L35:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v126)
	v128 = int32(1)
	v132 = v124 + v128
	v133 = v125 + v128
	goto L34
L36:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v118)
	v120 = int32(1)
	v124 = v116 + v120
	v125 = v117 + v120
	goto L35
L37:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v110)
	v112 = int32(1)
	v116 = v108 + v112
	v117 = v109 + v112
	goto L36
L38:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v102)
	v104 = int32(1)
	v108 = v100 + v104
	v109 = v101 + v104
	goto L37
L39:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v94)
	v96 = int32(1)
	v100 = v92 + v96
	v101 = v93 + v96
	goto L38
L40:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v86)
	v88 = int32(1)
	v92 = v84 + v88
	v93 = v85 + v88
	goto L39
L41:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v78)
	v80 = int32(1)
	v84 = v76 + v80
	v85 = v77 + v80
	goto L40
L42:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v70)
	v72 = int32(1)
	v76 = v68 + v72
	v77 = v69 + v72
	goto L41
L43:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v62)
	v64 = int32(1)
	v68 = v60 + v64
	v69 = v61 + v64
	goto L42
L44:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v54)
	v56 = int32(1)
	v60 = v52 + v56
	v61 = v53 + v56
	goto L43
L45:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v46)
	v52 = v15 + int32(2)
	v53 = v19 + int32(1)
	goto L44
L46:
	;
	v307 = int32(base.Ui32(v27) >> (uint(int32(5)) % 32))
	if v307 != int32(7) {
		v316 = v26
		v317 = v307
		goto L51
	} else {
		goto L52
	}
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(28)
	return int32(0)
L49:
	;
	if v347 == int32(0) {
		v500 = v19
		goto L82
	} else {
		goto L83
	}
L50:
	;
	goto L80
L51:
	;
	if base.Ui32(v19+v317+int32(2)) <= base.Ui32(v13) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v311 = v15 + int32(2)
	if base.Ui32(v14) <= base.Ui32(v311) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	v316 = v311
	v317 = v313 + int32(7)
	goto L51
L54:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316))))
	v335 = (v27^int32(-1))<<(uint(int32(8))%32) | int32(-7937) + v19 - v334
	if base.Ui32(l2) <= base.Ui32(v335) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(1)
	return int32(0)
L57:
	;
	v343 = v316 + int32(1)
	switch v317 + int32(-1) {
	case 0:
		v482 = v335
		v483 = v19
		goto L60
	case 1:
		v474 = v335
		v475 = v19
		goto L61
	case 2:
		v466 = v335
		v467 = v19
		goto L62
	case 3:
		v458 = v335
		v459 = v19
		goto L63
	case 4:
		v450 = v335
		v451 = v19
		goto L64
	case 5:
		v442 = v335
		v443 = v19
		goto L65
	case 6:
		v434 = v335
		v435 = v19
		goto L66
	case 7:
		v426 = v335
		v427 = v19
		goto L67
	case 8:
		goto L68
	default:
		goto L69
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(28)
	return int32(0)
L60:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	*(*uint8)(unsafe.Add(mBase, uint32(v483))) = uint8(v484)
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)) = uint8(v486)
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v483)+2)) = uint8(v488)
	v502 = v343
	v506 = v483 + int32(3)
	goto L5
L61:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	*(*uint8)(unsafe.Add(mBase, uint32(v475))) = uint8(v476)
	v478 = int32(1)
	v482 = v474 + v478
	v483 = v475 + v478
	goto L60
L62:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466))))
	*(*uint8)(unsafe.Add(mBase, uint32(v467))) = uint8(v468)
	v470 = int32(1)
	v474 = v466 + v470
	v475 = v467 + v470
	goto L61
L63:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	*(*uint8)(unsafe.Add(mBase, uint32(v459))) = uint8(v460)
	v462 = int32(1)
	v466 = v458 + v462
	v467 = v459 + v462
	goto L62
L64:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450))))
	*(*uint8)(unsafe.Add(mBase, uint32(v451))) = uint8(v452)
	v454 = int32(1)
	v458 = v450 + v454
	v459 = v451 + v454
	goto L63
L65:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
	*(*uint8)(unsafe.Add(mBase, uint32(v443))) = uint8(v444)
	v446 = int32(1)
	v450 = v442 + v446
	v451 = v443 + v446
	goto L64
L66:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	*(*uint8)(unsafe.Add(mBase, uint32(v435))) = uint8(v436)
	v438 = int32(1)
	v442 = v434 + v438
	v443 = v435 + v438
	goto L65
L67:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	*(*uint8)(unsafe.Add(mBase, uint32(v427))) = uint8(v428)
	v430 = int32(1)
	v434 = v426 + v430
	v435 = v427 + v430
	goto L66
L68:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v420)
	v422 = int32(1)
	v426 = v335 + v422
	v427 = v19 + v422
	goto L67
L69:
	;
	v347 = v317 + int32(2)
	if base.Ui32(v335+v347) <= base.Ui32(v19) {
		goto L49
	} else {
		goto L70
	}
L70:
	;
	v350 = int32(0)
	v352 = v347 & int32(7)
	if v352 == v350 {
		v377 = v347
		v379 = v335
		v380 = v19
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if base.Ui32(v317) < base.Ui32(int32(6)) {
		v502 = v343
		v506 = v380
		goto L5
	} else {
		goto L76
	}
L72:
	;
	v356 = v347
	v358 = v335
	v359 = v19
	v363 = v350
	goto L73
L73:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	*(*uint8)(unsafe.Add(mBase, uint32(v359))) = uint8(v365)
	v368 = v356 + int32(-1)
	v369 = int32(1)
	v370 = v359 + v369
	v372 = v358 + v369
	v374 = v363 + v369
	if v374 != v352 {
		v356 = v368
		v358 = v372
		v359 = v370
		v363 = v374
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v377 = v368
	v379 = v372
	v380 = v370
	goto L71
L75:
	;
	goto L74
L76:
	;
	v389 = v377
	v391 = v379
	v392 = v380
	goto L77
L77:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391))))
	*(*uint8)(unsafe.Add(mBase, uint32(v392))) = uint8(v398)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v392)+1)) = uint8(v400)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v392)+2)) = uint8(v402)
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v392)+3)) = uint8(v404)
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v392)+4)) = uint8(v406)
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v392)+5)) = uint8(v408)
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v392)+6)) = uint8(v410)
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v392)+7)) = uint8(v412)
	v414 = int32(8)
	v415 = v392 + v414
	v419 = v389 + int32(-8)
	if v419 != 0 {
		v389 = v419
		v391 = v391 + v414
		v392 = v415
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v502 = v343
	v506 = v415
	goto L5
L80:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(28)
	return int32(0)
L81:
	;
	v502 = v343
	v506 = v500 + v347
	goto L5
L82:
	;
	goto L81
L83:
	;
	v499 = F__emscripten_memcpy_bulkmem(m, v19, v335, v347)
	mBase = m.M
	v500 = v499
	goto L82
L84:
	;
	goto L4
}
