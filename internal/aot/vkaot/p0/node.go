package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_addNodeReplyForClusterSlot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	v8 = int32(3)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2164))
	goto L2
L1:
	;
	F_addReplyArrayLen(m, l0, v56)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	if v9 < int32(1) {
		v56 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = v8
	v18 = int32(0)
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2168))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v18<<(uint(int32(2))%32))))
	goto L7
L5:
	;
	v56 = v47
	goto L1
L6:
	;
	v49 = v18 + int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2164))
	goto L16
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
	goto L8
L8:
	;
	if v25&int32(8) != 0 {
		v47 = v17
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
	if v29&int32(16) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v42 == int64(0) {
		v47 = v17
		goto L6
	} else {
		goto L15
	}
L11:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v24)+2248))
	v42 = v41
	goto L10
L12:
	;
	if v29&int32(2) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = *(*int64)(unsafe.Add(mBase, _consts[31]))
	v42 = v40
	goto L10
L14:
	;
	v38 = F_replicationGetReplicaOffset(m)
	mBase = m.M
	v42 = v38
	goto L10
L15:
	;
	v47 = v17 + int32(1)
	goto L6
L16:
	;
	if v49 < v50 {
		v17 = v47
		v18 = v49
		goto L4
	} else {
		goto L17
	}
L17:
	;
	goto L5
L18:
	;
	return
L19:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(l2))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(l3))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_addNodeToNodeReply(m, l0, l1)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2164))
	goto L24
L23:
	;
	if v123 == int32(3) {
		goto L42
	} else {
		goto L43
	}
L24:
	;
	if v69 < int32(1) {
		v123 = v56
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v77 = v56
	v78 = int32(0)
	goto L26
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2168))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v78<<(uint(int32(2))%32))))
	goto L29
L27:
	;
	v123 = v114
	goto L23
L28:
	;
	v116 = v78 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2164))
	goto L40
L29:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+88))
	goto L30
L30:
	;
	if v85&int32(8) != 0 {
		v114 = v77
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)+88))
	if v89&int32(16) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v102 == int64(0) {
		v114 = v77
		goto L28
	} else {
		goto L37
	}
L33:
	;
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v84)+2248))
	v102 = v101
	goto L32
L34:
	;
	if v89&int32(2) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v100 = *(*int64)(unsafe.Add(mBase, _consts[31]))
	v102 = v100
	goto L32
L36:
	;
	v98 = F_replicationGetReplicaOffset(m)
	mBase = m.M
	v102 = v98
	goto L32
L37:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2168))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105+v78<<(uint(int32(2))%32))))
	goto L38
L38:
	;
	F_addNodeToNodeReply(m, l0, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	v114 = v77 + int32(-1)
	goto L28
L40:
	;
	if v116 < v117 {
		v77 = v114
		v78 = v116
		goto L26
	} else {
		goto L41
	}
L41:
	;
	goto L27
L42:
	;
	return
L43:
	;
	F__serverAssert(m, int32(_a177), int32(_a144), int32(1507))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L18
	} else {
		goto L44
	}
L44:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_addNodeToNodeReply(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2312))
	goto L1
L1:
	;
	F_addReplyArrayLen(m, l0, int32(4))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	switch v12 {
	case 0:
		goto L5
	case 1:
		goto L8
	case 2:
		goto L7
	default:
		goto L6
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	if v37 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L5:
	;
	v32 = F_clusterNodeIp(m, l1, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L16
	}
L6:
	;
	F__serverPanic_1(m, int32(_a144), int32(1424), int32(_a175), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L15
	}
L7:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L14
	}
L8:
	;
	if v7 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_addReplyBulkCString(m, l0, int32(_a176))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L13
	}
L10:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v15 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_addReplyBulkCString(m, l0, v7)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	goto L4
L13:
	;
	goto L4
L14:
	;
	goto L4
L15:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	F_addReplyBulkCString(m, l0, v32)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	goto L4
L18:
	;
	if v50 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[89]))
	v50 = v49
	goto L18
L20:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	if v40 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v44 = F_connectionTypeTls(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v50 = base.B2i32(v43 == v44)
	goto L18
L23:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v69))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L35
	}
L24:
	;
	v69 = v67
	goto L23
L25:
	;
	if l0 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2328))
	v69 = v61
	goto L23
L27:
	;
	if v50 == int32(0) {
		goto L25
	} else {
		goto L31
	}
L28:
	;
	if l0 == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2340))
	if v55 == int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v67 = v55
	goto L24
L31:
	;
	goto L26
L32:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
	v67 = v66
	goto L24
L33:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2336))
	if v64 != 0 {
		v67 = v64
		goto L24
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	goto L36
L36:
	;
	F_addReplyBulkCBuffer(m, l0, l1+int32(8), int32(40))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v81 = base.B2i32(v79 != int32(0))
	if v79 == int32(1) {
		v91 = v81
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2320))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+int32(-1)))))
	switch v95 & int32(7) {
	case 0:
		goto L53
	case 1:
		goto L52
	case 2:
		goto L51
	case 3:
		goto L50
	case 4:
		goto L49
	default:
		v117 = v91
		goto L47
	}
L39:
	;
	if v7 == int32(0) {
		v91 = v81
		goto L38
	} else {
		goto L40
	}
L40:
	;
	if v79 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v88 = int32(2)
	goto L43
L42:
	;
	v88 = int32(1)
	goto L43
L43:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v89 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v90 = v88
	goto L46
L45:
	;
	v90 = v81
	goto L46
L46:
	;
	v91 = v90
	goto L38
L47:
	;
	F_addReplyMapLen(m, l0, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L55
	}
L48:
	;
	if v112 == int32(0) {
		v117 = v91
		goto L47
	} else {
		goto L54
	}
L49:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v92+int32(-17))))
	v112 = v111
	goto L48
L50:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v92+int32(-9))))
	v112 = v108
	goto L48
L51:
	;
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92+int32(-5)))))
	v112 = v105
	goto L48
L52:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+int32(-3)))))
	v112 = v102
	goto L48
L53:
	;
	v112 = int32(base.Ui32(v95) >> (uint(int32(3)) % 32))
	goto L48
L54:
	;
	v117 = v91 + int32(1)
	goto L47
L55:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v122 == int32(0) {
		v138 = v117
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2320))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+int32(-1)))))
	switch v155 & int32(7) {
	case 0:
		goto L73
	case 1:
		goto L72
	case 2:
		goto L71
	case 3:
		goto L70
	case 4:
		goto L69
	default:
		v184 = v151
		goto L67
	}
L57:
	;
	if v7 == int32(0) {
		v151 = v138
		goto L56
	} else {
		goto L63
	}
L58:
	;
	F_addReplyBulkCString(m, l0, int32(_a174))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v128 = F_clusterNodeIp(m, l1, l0)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	F_addReplyBulkCString(m, l0, v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v133 = v117 + int32(-1)
	v135 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v135 == int32(1) {
		v151 = v133
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v138 = v133
	goto L57
L63:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v141 == int32(0) {
		v151 = v138
		goto L56
	} else {
		goto L64
	}
L64:
	;
	F_addReplyBulkCString(m, l0, int32(_a173))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	F_addReplyBulkCString(m, l0, v7)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v151 = v138 + int32(-1)
	goto L56
L67:
	;
	if v184 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L68:
	;
	if v172 == int32(0) {
		v184 = v151
		goto L67
	} else {
		goto L74
	}
L69:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v152+int32(-17))))
	v172 = v171
	goto L68
L70:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v152+int32(-9))))
	v172 = v168
	goto L68
L71:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152+int32(-5)))))
	v172 = v165
	goto L68
L72:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+int32(-3)))))
	v172 = v162
	goto L68
L73:
	;
	v172 = int32(base.Ui32(v155) >> (uint(int32(3)) % 32))
	goto L68
L74:
	;
	F_addReplyBulkCString(m, l0, int32(_a172))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2320))
	F_addReplyBulkCString(m, l0, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v184 = v151 + int32(-1)
	goto L67
L77:
	;
	return
L78:
	;
	F__serverAssert(m, int32(_a171), int32(_a144), int32(1468))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_getNodeReplicationOffset(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v3&int32(16) == int32(0) {
		v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+2248))
		return v40
	} else {
		if v3&int32(2) == int32(0) {
			v38 = *(*int64)(unsafe.Add(mBase, _consts[31]))
			return v38
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[64]))
			if v16 == int32(0) {
				v30 = int64(0)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, _consts[133]))
				if v20 != 0 {
					v27 = v20
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+104))
					v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
					v30 = v29
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, _consts[134]))
					if v23 == int32(0) {
						v30 = int64(0)
					} else {
						v27 = v23
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+104))
						v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
						v30 = v29
					}
				}
			}
			v32 = int64(0)
			if v32 < v30 {
				v35 = v30
			} else {
				v35 = v32
			}
			return v35
		}
	}
}
