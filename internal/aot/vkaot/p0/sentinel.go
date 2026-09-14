package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_loadSentinelConfigFromQueue(m *base.Module) {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[641]))
	if v11 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(64)
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v18 = v8 + int32(56)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v19
	goto L3
L3:
	;
	goto L7
L4:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[641]))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	F_listRelease(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L13
	} else {
		goto L39
	}
L5:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(_a535)
	v129 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	v133 = F_fiprintf(m, v129, int32(_a1139), v8+int32(48))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L13
	} else {
		goto L35
	}
L6:
	;
	v52 = v8 + int32(56)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v53
	goto L16
L7:
	;
	v29 = v8 + int32(56)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v31 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v31 == int32(0) {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31+base.B2i32(v34 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v40
	goto L10
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v47 = F_sentinelHandleConfiguration(m, v45, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v47 == int32(0) {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v120 = v44
	v123 = v47
	goto L5
L16:
	;
	goto L18
L17:
	;
	v86 = v8 + int32(56)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v87
	goto L26
L18:
	;
	v63 = v8 + int32(56)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v65 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v65 == int32(0) {
		goto L17
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65+base.B2i32(v68 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v74
	goto L21
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v81 = F_sentinelHandleConfiguration(m, v79, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	if v81 == int32(0) {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v120 = v78
	v123 = v81
	goto L5
L26:
	;
	goto L27
L27:
	;
	v97 = v8 + int32(56)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v99 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v120 = v112
	v123 = v115
	goto L5
L29:
	;
	if v99 == int32(0) {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v99+base.B2i32(v102 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v108
	goto L30
L32:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v115 = F_sentinelHandleConfiguration(m, v113, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	if v115 == int32(0) {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	goto L28
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v125
	v139 = F_fiprintf(m, v129, int32(_a1140), v8+int32(32))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L13
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v124
	v145 = F_fiprintf(m, v129, int32(_a1141), v8+int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v123
	v149 = F_fiprintf(m, v129, int32(_a1142), v8)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _consts[641]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	F_listRelease(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[641]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	F_listRelease(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L13
	} else {
		goto L41
	}
L41:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _consts[641]))
	F_valkey_free(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[641])) = int32(0)
	goto L1
}
func F_sentinelCheckSubjectivelyDown(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v121 int64
	_ = v121
	var v124 int64
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	v2 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)+56))
	if v7 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v24 == int32(0) {
		v69 = v22
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v16 = F_mstime(m)
	mBase = m.M
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v17+v15)))
	v21 = v16 - v19
	v22 = v17
	goto L1
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v11 == int32(0) {
		v21 = v2
		v22 = v6
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v15 = int32(56)
	goto L2
L5:
	;
	v15 = int32(48)
	goto L2
L6:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	if v71 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v27 = F_mstime(m)
	mBase = m.M
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
	v32 = *(*int64)(unsafe.Add(mBase, _consts[661]))
	if v27-v29 <= v32 {
		v69 = v28
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v28)+56))
	if v34 == int64(0) {
		v69 = v28
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v37 = F_mstime(m)
	mBase = m.M
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+56))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v43 = base.I64_div_s(v41, int64(2))
	if v37-v39 <= v43 {
		v69 = v38
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v45 = F_mstime(m)
	mBase = m.M
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)+72))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v51 = base.I64_div_s(v49, int64(2))
	if v45-v47 <= v51 {
		v69 = v46
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	if v53 == int32(0) {
		v69 = v46
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+8)) = int64(0)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	if v58 != v53 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+212)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = int32(1)
	F_valkeyAsyncFree(m, v53)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = int32(0)
	goto L13
L15:
	;
	return
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v69 = v68
	goto L6
L17:
	;
	v107 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	if v107 < v21 {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	v74 = F_mstime(m)
	mBase = m.M
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)+32))
	v79 = *(*int64)(unsafe.Add(mBase, _consts[661]))
	if v74-v76 <= v79 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v81 = F_mstime(m)
	mBase = m.M
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
	v86 = *(*int64)(unsafe.Add(mBase, _consts[662]))
	if v81-v83 <= v86*int64(3) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
	if v90 == int32(0) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	if v93 != v90 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v97 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v90)+212)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = int32(1)
	F_valkeyAsyncFree(m, v90)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L15
	} else {
		goto L24
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = int64(0)
	goto L22
L24:
	;
	goto L17
L25:
	;
	return
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v168
	goto L25
L27:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v154&int32(8) != 0 {
		goto L25
	} else {
		goto L38
	}
L28:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v109&int32(1) == int32(0) {
		v128 = v109
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v128&int32(8192) == int32(0) {
		v139 = v128
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v114 != int32(2) {
		v128 = v109
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v117 = F_mstime(m)
	mBase = m.M
	v118 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v121 = *(*int64)(unsafe.Add(mBase, _consts[653]))
	v124 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	if v121<<(uint(int64(1))%64)+v124 < v117-v118 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v128 = v127
	goto L29
L33:
	;
	if v139&int32(8) == int32(0) {
		goto L25
	} else {
		goto L36
	}
L34:
	;
	v133 = F_mstime(m)
	mBase = m.M
	v134 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v136 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	if v136 < v133-v134 {
		goto L27
	} else {
		goto L35
	}
L35:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v139 = v138
	goto L33
L36:
	;
	F_sentinelEvent(m, int32(3), int32(_a1232), l0, int32(_a1138), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v168 = v150 & int32(-4105)
	goto L26
L38:
	;
	F_sentinelEvent(m, int32(3), int32(_a1231), l0, int32(_a1138), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L15
	} else {
		goto L39
	}
L39:
	;
	v163 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v163
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v168 = v165 | int32(8)
	goto L26
}
func F_sentinelCommand(m *base.Module, l0 int32) {
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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v65 int32
	_ = v65
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
	var v77 int32
	_ = v77
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
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
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
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
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int64
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int64
	_ = v658
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v668 int64
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int64
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v981 int64
	_ = v981
	var v983 int32
	_ = v983
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1616 int64
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1620 int64
	_ = v1620
	var v1625 int32
	_ = v1625
	var v1630 int64
	_ = v1630
	var v1635 int64
	_ = v1635
	var v1640 int64
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1657 int32
	_ = v1657
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1718 int32
	_ = v1718
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1731 int64
	_ = v1731
	var v1732 int64
	_ = v1732
	var v1733 int64
	_ = v1733
	var v1734 int64
	_ = v1734
	var v1735 int64
	_ = v1735
	var v1736 int64
	_ = v1736
	var v1737 int64
	_ = v1737
	var v1739 int64
	_ = v1739
	var v1741 int64
	_ = v1741
	var v1743 int64
	_ = v1743
	var v1745 int64
	_ = v1745
	var v1747 int64
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1769 int32
	_ = v1769
	var v1773 int32
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1789 int32
	_ = v1789
	var v1794 int32
	_ = v1794
	var v1798 int32
	_ = v1798
	var v1802 int32
	_ = v1802
	var v1810 int32
	_ = v1810
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1833 int32
	_ = v1833
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1844 int32
	_ = v1844
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1859 int32
	_ = v1859
	var v1863 int32
	_ = v1863
	var v1865 int32
	_ = v1865
	var v1873 int32
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1887 int64
	_ = v1887
	var v1888 int64
	_ = v1888
	var v1892 int64
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1903 int32
	_ = v1903
	var v1910 int32
	_ = v1910
	var v1913 int32
	_ = v1913
	var v1916 int32
	_ = v1916
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1924 int32
	_ = v1924
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1958 int32
	_ = v1958
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1969 int32
	_ = v1969
	var v1971 int64
	_ = v1971
	var v1972 int64
	_ = v1972
	var v1973 int64
	_ = v1973
	var v1974 int64
	_ = v1974
	var v1975 int64
	_ = v1975
	var v1976 int64
	_ = v1976
	var v1977 int64
	_ = v1977
	var v1979 int64
	_ = v1979
	var v1981 int64
	_ = v1981
	var v1983 int64
	_ = v1983
	var v1985 int64
	_ = v1985
	var v1987 int64
	_ = v1987
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2001 int32
	_ = v2001
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2009 int32
	_ = v2009
	var v2013 int32
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2029 int32
	_ = v2029
	var v2034 int32
	_ = v2034
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2048 int32
	_ = v2048
	var v2051 int32
	_ = v2051
	var v2052 int64
	_ = v2052
	var v2058 int64
	_ = v2058
	var v2060 int64
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2071 int32
	_ = v2071
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2092 int32
	_ = v2092
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2108 int32
	_ = v2108
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2116 int32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2129 int64
	_ = v2129
	var v2130 int64
	_ = v2130
	var v2131 int64
	_ = v2131
	var v2132 int64
	_ = v2132
	var v2133 int64
	_ = v2133
	var v2134 int64
	_ = v2134
	var v2135 int64
	_ = v2135
	var v2137 int64
	_ = v2137
	var v2139 int64
	_ = v2139
	var v2141 int64
	_ = v2141
	var v2143 int64
	_ = v2143
	var v2145 int64
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2152 int32
	_ = v2152
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2159 int32
	_ = v2159
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2167 int32
	_ = v2167
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2187 int32
	_ = v2187
	var v2192 int32
	_ = v2192
	var v2196 int32
	_ = v2196
	var v2200 int32
	_ = v2200
	var v2214 int32
	_ = v2214
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2226 int32
	_ = v2226
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2232 int32
	_ = v2232
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2262 int32
	_ = v2262
	var v2265 int32
	_ = v2265
	var v2271 int32
	_ = v2271
	var v2278 int32
	_ = v2278
	var v2280 int32
	_ = v2280
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2293 int32
	_ = v2293
	var v2297 int32
	_ = v2297
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2330 int32
	_ = v2330
	var v2334 int32
	_ = v2334
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2341 int32
	_ = v2341
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2351 int32
	_ = v2351
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2377 int32
	_ = v2377
	var v2379 int32
	_ = v2379
	var v2384 int32
	_ = v2384
	var v2388 int32
	_ = v2388
	var v2392 int32
	_ = v2392
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2408 int32
	_ = v2408
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2420 int32
	_ = v2420
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2426 int32
	_ = v2426
	var v2430 int32
	_ = v2430
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2436 int32
	_ = v2436
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2444 int32
	_ = v2444
	var v2445 int32
	_ = v2445
	var v2448 int32
	_ = v2448
	var v2449 int32
	_ = v2449
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2458 int32
	_ = v2458
	var v2461 int32
	_ = v2461
	var v2464 int32
	_ = v2464
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2480 int32
	_ = v2480
	var v2484 int32
	_ = v2484
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2510 int32
	_ = v2510
	var v2514 int32
	_ = v2514
	var v2516 int32
	_ = v2516
	var v2518 int32
	_ = v2518
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2524 int64
	_ = v2524
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2531 int64
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2534 int32
	_ = v2534
	var v2536 int32
	_ = v2536
	var v2538 int32
	_ = v2538
	var v2540 int32
	_ = v2540
	var v2544 int32
	_ = v2544
	v10 = m.G0
	v12 = v10 - int32(256)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v14 != int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(256)
	return
L2:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v73 = F_objectGetVal(m, v72)
	mBase = m.M
	v74 = int32(_a1171)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v77 != 0 {
		goto L29
	} else {
		goto L30
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = F_objectGetVal(m, v18)
	mBase = m.M
	v20 = int32(_a757)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v55-v57 != 0 {
		goto L2
	} else {
		goto L16
	}
L5:
	;
	v55 = F_tolower(m, v51)
	mBase = m.M
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v57 = F_tolower(m, v56)
	mBase = m.M
	goto L4
L6:
	;
	v25 = v19
	v26 = v20
	v27 = v23
	goto L9
L7:
	;
	v51 = int32(0)
	v52 = v20
	goto L5
L8:
	;
	v51 = v48 & int32(255)
	v52 = v47
	goto L5
L9:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v29 == int32(0) {
		v47 = v26
		v48 = v27
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v47 = v41
	v48 = int32(0)
	goto L8
L11:
	;
	v33 = v27 & int32(255)
	if v33 == v29 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v40 = int32(1)
	v41 = v26 + v40
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v42 != 0 {
		v25 = v25 + v40
		v26 = v41
		v27 = v42
		goto L9
	} else {
		goto L15
	}
L13:
	;
	v35 = F_tolower(m, v33)
	mBase = m.M
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v37 = F_tolower(m, v36)
	mBase = m.M
	if v35 == v37 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v47 = v26
	v48 = v39
	goto L8
L15:
	;
	goto L10
L16:
	;
	goto L19
L17:
	;
	F_addReplyHelp(m, l0, v12+int32(64))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	v65 = F__emscripten_memcpy_bulkmem(m, v12+int32(64), int32(_a1172), int32(188))
	mBase = m.M
	goto L18
L20:
	;
	return
L21:
	;
	goto L1
L22:
	;
	v2532 = *(*int32)(unsafe.Add(mBase, uint32(v2529)))
	F_addReply(m, l0, v2532)
	mBase = m.M
	v2534 = m.ExcPending
	if v2534 != 0 {
		goto L20
	} else {
		goto L784
	}
L23:
	;
	v2528 = v2521
	v2529 = int32(_a619)
	v2531 = v2524
	goto L22
L24:
	;
	F_addReplyErrorArity(m, l0)
	mBase = m.M
	v2520 = m.ExcPending
	if v2520 != 0 {
		goto L20
	} else {
		goto L783
	}
L25:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v166 = F_objectGetVal(m, v165)
	mBase = m.M
	v167 = int32(_a202)
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	if v170 != 0 {
		goto L59
	} else {
		goto L60
	}
L26:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v157 != int32(2) {
		goto L24
	} else {
		goto L53
	}
L27:
	;
	if v109-v111 == int32(0) {
		goto L26
	} else {
		goto L39
	}
L28:
	;
	v109 = F_tolower(m, v105)
	mBase = m.M
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v111 = F_tolower(m, v110)
	mBase = m.M
	goto L27
L29:
	;
	v79 = v73
	v80 = v74
	v81 = v77
	goto L32
L30:
	;
	v105 = int32(0)
	v106 = v74
	goto L28
L31:
	;
	v105 = v102 & int32(255)
	v106 = v101
	goto L28
L32:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v83 == int32(0) {
		v101 = v80
		v102 = v81
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v101 = v95
	v102 = int32(0)
	goto L31
L34:
	;
	v87 = v81 & int32(255)
	if v87 == v83 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v94 = int32(1)
	v95 = v80 + v94
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	if v96 != 0 {
		v79 = v79 + v94
		v80 = v95
		v81 = v96
		goto L32
	} else {
		goto L38
	}
L36:
	;
	v89 = F_tolower(m, v87)
	mBase = m.M
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v91 = F_tolower(m, v90)
	mBase = m.M
	if v89 == v91 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v101 = v80
	v102 = v93
	goto L31
L38:
	;
	goto L33
L39:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v117 = F_objectGetVal(m, v116)
	mBase = m.M
	v118 = int32(_a1173)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v121 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if v153-v155 != 0 {
		goto L25
	} else {
		goto L52
	}
L41:
	;
	v153 = F_tolower(m, v149)
	mBase = m.M
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v155 = F_tolower(m, v154)
	mBase = m.M
	goto L40
L42:
	;
	v123 = v117
	v124 = v118
	v125 = v121
	goto L45
L43:
	;
	v149 = int32(0)
	v150 = v118
	goto L41
L44:
	;
	v149 = v146 & int32(255)
	v150 = v145
	goto L41
L45:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v127 == int32(0) {
		v145 = v124
		v146 = v125
		goto L44
	} else {
		goto L47
	}
L46:
	;
	v145 = v139
	v146 = int32(0)
	goto L44
L47:
	;
	v131 = v125 & int32(255)
	if v131 == v127 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v138 = int32(1)
	v139 = v124 + v138
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
	if v140 != 0 {
		v123 = v123 + v138
		v124 = v139
		v125 = v140
		goto L45
	} else {
		goto L51
	}
L49:
	;
	v133 = F_tolower(m, v131)
	mBase = m.M
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	v135 = F_tolower(m, v134)
	mBase = m.M
	if v133 == v135 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v145 = v124
	v146 = v137
	goto L44
L51:
	;
	goto L46
L52:
	;
	goto L26
L53:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	F_addReplyDictOfValkeyInstances(m, l0, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L20
	} else {
		goto L54
	}
L54:
	;
	goto L1
L55:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v267 = F_objectGetVal(m, v266)
	mBase = m.M
	v268 = int32(_a1174)
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	if v271 != 0 {
		goto L93
	} else {
		goto L94
	}
L56:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v250 != int32(3) {
		goto L24
	} else {
		goto L83
	}
L57:
	;
	if v202-v204 == int32(0) {
		goto L56
	} else {
		goto L69
	}
L58:
	;
	v202 = F_tolower(m, v198)
	mBase = m.M
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	v204 = F_tolower(m, v203)
	mBase = m.M
	goto L57
L59:
	;
	v172 = v166
	v173 = v167
	v174 = v170
	goto L62
L60:
	;
	v198 = int32(0)
	v199 = v167
	goto L58
L61:
	;
	v198 = v195 & int32(255)
	v199 = v194
	goto L58
L62:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	if v176 == int32(0) {
		v194 = v173
		v195 = v174
		goto L61
	} else {
		goto L64
	}
L63:
	;
	v194 = v188
	v195 = int32(0)
	goto L61
L64:
	;
	v180 = v174 & int32(255)
	if v180 == v176 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v187 = int32(1)
	v188 = v173 + v187
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
	if v189 != 0 {
		v172 = v172 + v187
		v173 = v188
		v174 = v189
		goto L62
	} else {
		goto L68
	}
L66:
	;
	v182 = F_tolower(m, v180)
	mBase = m.M
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	v184 = F_tolower(m, v183)
	mBase = m.M
	if v182 == v184 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	v194 = v173
	v195 = v186
	goto L61
L68:
	;
	goto L63
L69:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	v210 = F_objectGetVal(m, v209)
	mBase = m.M
	v211 = int32(_a200)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if v214 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	if v246-v248 != 0 {
		goto L55
	} else {
		goto L82
	}
L71:
	;
	v246 = F_tolower(m, v242)
	mBase = m.M
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v248 = F_tolower(m, v247)
	mBase = m.M
	goto L70
L72:
	;
	v216 = v210
	v217 = v211
	v218 = v214
	goto L75
L73:
	;
	v242 = int32(0)
	v243 = v211
	goto L71
L74:
	;
	v242 = v239 & int32(255)
	v243 = v238
	goto L71
L75:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v220 == int32(0) {
		v238 = v217
		v239 = v218
		goto L74
	} else {
		goto L77
	}
L76:
	;
	v238 = v232
	v239 = int32(0)
	goto L74
L77:
	;
	v224 = v218 & int32(255)
	if v224 == v220 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v231 = int32(1)
	v232 = v217 + v231
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	if v233 != 0 {
		v216 = v216 + v231
		v217 = v232
		v218 = v233
		goto L75
	} else {
		goto L81
	}
L79:
	;
	v226 = F_tolower(m, v224)
	mBase = m.M
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	v228 = F_tolower(m, v227)
	mBase = m.M
	if v226 == v228 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v238 = v217
	v239 = v230
	goto L74
L81:
	;
	goto L76
L82:
	;
	goto L56
L83:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	v257 = F_objectGetVal(m, v256)
	mBase = m.M
	v258 = F_dictFetchValue(m, v254, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L20
	} else {
		goto L85
	}
L84:
	;
	F_addReplySentinelValkeyInstance(m, l0, v258)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L20
	} else {
		goto L88
	}
L85:
	;
	if v258 != 0 {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	F_addReplyError(m, l0, int32(_a1170))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L20
	} else {
		goto L87
	}
L87:
	;
	goto L1
L88:
	;
	goto L1
L89:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	v369 = F_objectGetVal(m, v368)
	mBase = m.M
	v370 = int32(_a1175)
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369))))
	if v373 != 0 {
		goto L126
	} else {
		goto L127
	}
L90:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v351 != int32(3) {
		goto L24
	} else {
		goto L117
	}
L91:
	;
	if v303-v305 == int32(0) {
		goto L90
	} else {
		goto L103
	}
L92:
	;
	v303 = F_tolower(m, v299)
	mBase = m.M
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	v305 = F_tolower(m, v304)
	mBase = m.M
	goto L91
L93:
	;
	v273 = v267
	v274 = v268
	v275 = v271
	goto L96
L94:
	;
	v299 = int32(0)
	v300 = v268
	goto L92
L95:
	;
	v299 = v296 & int32(255)
	v300 = v295
	goto L92
L96:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	if v277 == int32(0) {
		v295 = v274
		v296 = v275
		goto L95
	} else {
		goto L98
	}
L97:
	;
	v295 = v289
	v296 = int32(0)
	goto L95
L98:
	;
	v281 = v275 & int32(255)
	if v281 == v277 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v288 = int32(1)
	v289 = v274 + v288
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+1)))
	if v290 != 0 {
		v273 = v273 + v288
		v274 = v289
		v275 = v290
		goto L96
	} else {
		goto L102
	}
L100:
	;
	v283 = F_tolower(m, v281)
	mBase = m.M
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	v285 = F_tolower(m, v284)
	mBase = m.M
	if v283 == v285 {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	v295 = v274
	v296 = v287
	goto L95
L102:
	;
	goto L97
L103:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	v311 = F_objectGetVal(m, v310)
	mBase = m.M
	v312 = int32(_a1176)
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	if v315 != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	if v347-v349 != 0 {
		goto L89
	} else {
		goto L116
	}
L105:
	;
	v347 = F_tolower(m, v343)
	mBase = m.M
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	v349 = F_tolower(m, v348)
	mBase = m.M
	goto L104
L106:
	;
	v317 = v311
	v318 = v312
	v319 = v315
	goto L109
L107:
	;
	v343 = int32(0)
	v344 = v312
	goto L105
L108:
	;
	v343 = v340 & int32(255)
	v344 = v339
	goto L105
L109:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	if v321 == int32(0) {
		v339 = v318
		v340 = v319
		goto L108
	} else {
		goto L111
	}
L110:
	;
	v339 = v333
	v340 = int32(0)
	goto L108
L111:
	;
	v325 = v319 & int32(255)
	if v325 == v321 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v332 = int32(1)
	v333 = v318 + v332
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+1)))
	if v334 != 0 {
		v317 = v317 + v332
		v318 = v333
		v319 = v334
		goto L109
	} else {
		goto L115
	}
L113:
	;
	v327 = F_tolower(m, v325)
	mBase = m.M
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	v329 = F_tolower(m, v328)
	mBase = m.M
	if v327 == v329 {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
	v339 = v318
	v340 = v331
	goto L108
L115:
	;
	goto L110
L116:
	;
	goto L90
L117:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+8))
	v358 = F_objectGetVal(m, v357)
	mBase = m.M
	v359 = F_dictFetchValue(m, v355, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L20
	} else {
		goto L119
	}
L118:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v359)+148))
	F_addReplyDictOfValkeyInstances(m, l0, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L20
	} else {
		goto L122
	}
L119:
	;
	if v359 != 0 {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	F_addReplyError(m, l0, int32(_a1170))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L20
	} else {
		goto L121
	}
L121:
	;
	goto L1
L122:
	;
	goto L1
L123:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	v423 = F_objectGetVal(m, v422)
	mBase = m.M
	v424 = int32(_a1177)
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423))))
	if v427 != 0 {
		goto L144
	} else {
		goto L145
	}
L124:
	;
	if v405-v407 != 0 {
		goto L123
	} else {
		goto L136
	}
L125:
	;
	v405 = F_tolower(m, v401)
	mBase = m.M
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402))))
	v407 = F_tolower(m, v406)
	mBase = m.M
	goto L124
L126:
	;
	v375 = v369
	v376 = v370
	v377 = v373
	goto L129
L127:
	;
	v401 = int32(0)
	v402 = v370
	goto L125
L128:
	;
	v401 = v398 & int32(255)
	v402 = v397
	goto L125
L129:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	if v379 == int32(0) {
		v397 = v376
		v398 = v377
		goto L128
	} else {
		goto L131
	}
L130:
	;
	v397 = v391
	v398 = int32(0)
	goto L128
L131:
	;
	v383 = v377 & int32(255)
	if v383 == v379 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v390 = int32(1)
	v391 = v376 + v390
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375)+1)))
	if v392 != 0 {
		v375 = v375 + v390
		v376 = v391
		v377 = v392
		goto L129
	} else {
		goto L135
	}
L133:
	;
	v385 = F_tolower(m, v383)
	mBase = m.M
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	v387 = F_tolower(m, v386)
	mBase = m.M
	if v385 == v387 {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375))))
	v397 = v376
	v398 = v389
	goto L128
L135:
	;
	goto L130
L136:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v409 != int32(3) {
		goto L24
	} else {
		goto L137
	}
L137:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)+8))
	v414 = F_sentinelGetPrimaryByNameOrReplyError(m, l0, v413)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L20
	} else {
		goto L138
	}
L138:
	;
	if v414 == int32(0) {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414)+144))
	F_addReplyDictOfValkeyInstances(m, l0, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L20
	} else {
		goto L140
	}
L140:
	;
	goto L1
L141:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	v472 = F_objectGetVal(m, v471)
	mBase = m.M
	v473 = int32(_a1178)
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472))))
	if v476 != 0 {
		goto L161
	} else {
		goto L162
	}
L142:
	;
	if v459-v461 != 0 {
		goto L141
	} else {
		goto L154
	}
L143:
	;
	v459 = F_tolower(m, v455)
	mBase = m.M
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	v461 = F_tolower(m, v460)
	mBase = m.M
	goto L142
L144:
	;
	v429 = v423
	v430 = v424
	v431 = v427
	goto L147
L145:
	;
	v455 = int32(0)
	v456 = v424
	goto L143
L146:
	;
	v455 = v452 & int32(255)
	v456 = v451
	goto L143
L147:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	if v433 == int32(0) {
		v451 = v430
		v452 = v431
		goto L146
	} else {
		goto L149
	}
L148:
	;
	v451 = v445
	v452 = int32(0)
	goto L146
L149:
	;
	v437 = v431 & int32(255)
	if v437 == v433 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v444 = int32(1)
	v445 = v430 + v444
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+1)))
	if v446 != 0 {
		v429 = v429 + v444
		v430 = v445
		v431 = v446
		goto L147
	} else {
		goto L153
	}
L151:
	;
	v439 = F_tolower(m, v437)
	mBase = m.M
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	v441 = F_tolower(m, v440)
	mBase = m.M
	if v439 == v441 {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429))))
	v451 = v430
	v452 = v443
	goto L146
L153:
	;
	goto L148
L154:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v463 != int32(2) {
		goto L141
	} else {
		goto L155
	}
L155:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a1143), int32(40))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L20
	} else {
		goto L156
	}
L156:
	;
	goto L1
L157:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v673)+4))
	v675 = F_objectGetVal(m, v674)
	mBase = m.M
	v676 = int32(_a32)
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675))))
	if v679 != 0 {
		goto L221
	} else {
		goto L222
	}
L158:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = int64(0)
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v558 != int32(6) {
		goto L24
	} else {
		goto L185
	}
L159:
	;
	if v508-v510 == int32(0) {
		goto L158
	} else {
		goto L171
	}
L160:
	;
	v508 = F_tolower(m, v504)
	mBase = m.M
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505))))
	v510 = F_tolower(m, v509)
	mBase = m.M
	goto L159
L161:
	;
	v478 = v472
	v479 = v473
	v480 = v476
	goto L164
L162:
	;
	v504 = int32(0)
	v505 = v473
	goto L160
L163:
	;
	v504 = v501 & int32(255)
	v505 = v500
	goto L160
L164:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if v482 == int32(0) {
		v500 = v479
		v501 = v480
		goto L163
	} else {
		goto L166
	}
L165:
	;
	v500 = v494
	v501 = int32(0)
	goto L163
L166:
	;
	v486 = v480 & int32(255)
	if v486 == v482 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v493 = int32(1)
	v494 = v479 + v493
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+1)))
	if v495 != 0 {
		v478 = v478 + v493
		v479 = v494
		v480 = v495
		goto L164
	} else {
		goto L170
	}
L168:
	;
	v488 = F_tolower(m, v486)
	mBase = m.M
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	v490 = F_tolower(m, v489)
	mBase = m.M
	if v488 == v490 {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	v500 = v479
	v501 = v492
	goto L163
L170:
	;
	goto L165
L171:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	v516 = F_objectGetVal(m, v515)
	mBase = m.M
	v517 = int32(_a1179)
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
	if v520 != 0 {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	if v552-v554 != 0 {
		goto L157
	} else {
		goto L184
	}
L173:
	;
	v552 = F_tolower(m, v548)
	mBase = m.M
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549))))
	v554 = F_tolower(m, v553)
	mBase = m.M
	goto L172
L174:
	;
	v522 = v516
	v523 = v517
	v524 = v520
	goto L177
L175:
	;
	v548 = int32(0)
	v549 = v517
	goto L173
L176:
	;
	v548 = v545 & int32(255)
	v549 = v544
	goto L173
L177:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
	if v526 == int32(0) {
		v544 = v523
		v545 = v524
		goto L176
	} else {
		goto L179
	}
L178:
	;
	v544 = v538
	v545 = int32(0)
	goto L176
L179:
	;
	v530 = v524 & int32(255)
	if v530 == v526 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v537 = int32(1)
	v538 = v523 + v537
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+1)))
	if v539 != 0 {
		v522 = v522 + v537
		v523 = v538
		v524 = v539
		goto L177
	} else {
		goto L183
	}
L181:
	;
	v532 = F_tolower(m, v530)
	mBase = m.M
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
	v534 = F_tolower(m, v533)
	mBase = m.M
	if v532 == v534 {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	v544 = v523
	v545 = v536
	goto L176
L183:
	;
	goto L178
L184:
	;
	goto L158
L185:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)+12))
	v566 = F_getLongFromObjectOrReply(m, l0, v562, v12+int32(52), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L20
	} else {
		goto L186
	}
L186:
	;
	if v566 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)+16))
	v573 = F_getLongLongFromObjectOrReply(m, l0, v569, v12+int32(64), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L20
	} else {
		goto L188
	}
L188:
	;
	if v573 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v575 = int32(0)
	v576 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v577)+8))
	v579 = F_objectGetVal(m, v578)
	mBase = m.M
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v582 = F_getSentinelValkeyInstanceByAddrAndRunID(m, v576, v579, v580, v575)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L20
	} else {
		goto L190
	}
L190:
	;
	v585 = *(*int32)(unsafe.Add(mBase, _consts[645]))
	if v585 != 0 {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	if v600&int32(1) == int32(0) {
		goto L198
	} else {
		goto L199
	}
L192:
	;
	if v582 != 0 {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	if v582 == int32(0) {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	v589 = int32(9)
	v600 = v588
	v601 = base.B2i32(v588&v589 != v589)
	goto L191
L195:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	v600 = v598
	v601 = int32(1)
	goto L191
L196:
	;
	F_addReplyArrayLen(m, l0, int32(3))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L20
	} else {
		goto L197
	}
L197:
	;
	v2521 = int32(0)
	v2524 = int64(0)
	goto L23
L198:
	;
	F_addReplyArrayLen(m, l0, int32(3))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L20
	} else {
		goto L216
	}
L199:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)+20))
	v608 = F_objectGetVal(m, v607)
	mBase = m.M
	v609 = int32(_a1180)
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	if v612 != 0 {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	if v644-v646 == int32(0) {
		goto L198
	} else {
		goto L212
	}
L201:
	;
	v644 = F_tolower(m, v640)
	mBase = m.M
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641))))
	v646 = F_tolower(m, v645)
	mBase = m.M
	goto L200
L202:
	;
	v614 = v608
	v615 = v609
	v616 = v612
	goto L205
L203:
	;
	v640 = int32(0)
	v641 = v609
	goto L201
L204:
	;
	v640 = v637 & int32(255)
	v641 = v636
	goto L201
L205:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615))))
	if v618 == int32(0) {
		v636 = v615
		v637 = v616
		goto L204
	} else {
		goto L207
	}
L206:
	;
	v636 = v630
	v637 = int32(0)
	goto L204
L207:
	;
	v622 = v616 & int32(255)
	if v622 == v618 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v629 = int32(1)
	v630 = v615 + v629
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+1)))
	if v631 != 0 {
		v614 = v614 + v629
		v615 = v630
		v616 = v631
		goto L205
	} else {
		goto L211
	}
L209:
	;
	v624 = F_tolower(m, v622)
	mBase = m.M
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615))))
	v626 = F_tolower(m, v625)
	mBase = m.M
	if v624 == v626 {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	v636 = v615
	v637 = v628
	goto L204
L211:
	;
	goto L206
L212:
	;
	v650 = *(*int64)(unsafe.Add(mBase, uint32(v12)+64))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v651)+20))
	v653 = F_objectGetVal(m, v652)
	mBase = m.M
	v656 = F_sentinelVoteLeader(m, v582, v650, v653, v12+int32(56))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L20
	} else {
		goto L213
	}
L213:
	;
	v658 = *(*int64)(unsafe.Add(mBase, uint32(v12)+56))
	F_addReplyArrayLen(m, l0, int32(3))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L20
	} else {
		goto L214
	}
L214:
	;
	if v601 != 0 {
		v2521 = v656
		v2524 = v658
		goto L23
	} else {
		goto L215
	}
L215:
	;
	v2528 = v656
	v2529 = int32(_a1181)
	v2531 = v658
	goto L22
L216:
	;
	v668 = int64(0)
	v669 = int32(0)
	if v601 != 0 {
		v2521 = v669
		v2524 = v668
		goto L23
	} else {
		goto L217
	}
L217:
	;
	v2528 = v669
	v2529 = int32(_a1181)
	v2531 = v668
	goto L22
L218:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v727)+4))
	v729 = F_objectGetVal(m, v728)
	mBase = m.M
	v730 = int32(_a1182)
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v729))))
	if v733 != 0 {
		goto L239
	} else {
		goto L240
	}
L219:
	;
	if v711-v713 != 0 {
		goto L218
	} else {
		goto L231
	}
L220:
	;
	v711 = F_tolower(m, v707)
	mBase = m.M
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708))))
	v713 = F_tolower(m, v712)
	mBase = m.M
	goto L219
L221:
	;
	v681 = v675
	v682 = v676
	v683 = v679
	goto L224
L222:
	;
	v707 = int32(0)
	v708 = v676
	goto L220
L223:
	;
	v707 = v704 & int32(255)
	v708 = v703
	goto L220
L224:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682))))
	if v685 == int32(0) {
		v703 = v682
		v704 = v683
		goto L223
	} else {
		goto L226
	}
L225:
	;
	v703 = v697
	v704 = int32(0)
	goto L223
L226:
	;
	v689 = v683 & int32(255)
	if v689 == v685 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v696 = int32(1)
	v697 = v682 + v696
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681)+1)))
	if v698 != 0 {
		v681 = v681 + v696
		v682 = v697
		v683 = v698
		goto L224
	} else {
		goto L230
	}
L228:
	;
	v691 = F_tolower(m, v689)
	mBase = m.M
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682))))
	v693 = F_tolower(m, v692)
	mBase = m.M
	if v691 == v693 {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681))))
	v703 = v682
	v704 = v695
	goto L223
L230:
	;
	goto L225
L231:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v715 != int32(3) {
		goto L24
	} else {
		goto L232
	}
L232:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v718)+8))
	v720 = F_objectGetVal(m, v719)
	mBase = m.M
	v722 = F_sentinelResetPrimariesByPattern(m, v720, int32(65536))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L20
	} else {
		goto L233
	}
L233:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v722))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L20
	} else {
		goto L234
	}
L234:
	;
	goto L1
L235:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v853)+4))
	v855 = F_objectGetVal(m, v854)
	mBase = m.M
	v856 = int32(_a496)
	v859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v855))))
	if v859 != 0 {
		goto L279
	} else {
		goto L280
	}
L236:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v813 != int32(3) {
		goto L24
	} else {
		goto L263
	}
L237:
	;
	if v765-v767 == int32(0) {
		goto L236
	} else {
		goto L249
	}
L238:
	;
	v765 = F_tolower(m, v761)
	mBase = m.M
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762))))
	v767 = F_tolower(m, v766)
	mBase = m.M
	goto L237
L239:
	;
	v735 = v729
	v736 = v730
	v737 = v733
	goto L242
L240:
	;
	v761 = int32(0)
	v762 = v730
	goto L238
L241:
	;
	v761 = v758 & int32(255)
	v762 = v757
	goto L238
L242:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736))))
	if v739 == int32(0) {
		v757 = v736
		v758 = v737
		goto L241
	} else {
		goto L244
	}
L243:
	;
	v757 = v751
	v758 = int32(0)
	goto L241
L244:
	;
	v743 = v737 & int32(255)
	if v743 == v739 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v750 = int32(1)
	v751 = v736 + v750
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735)+1)))
	if v752 != 0 {
		v735 = v735 + v750
		v736 = v751
		v737 = v752
		goto L242
	} else {
		goto L248
	}
L246:
	;
	v745 = F_tolower(m, v743)
	mBase = m.M
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736))))
	v747 = F_tolower(m, v746)
	mBase = m.M
	if v745 == v747 {
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735))))
	v757 = v736
	v758 = v749
	goto L241
L248:
	;
	goto L243
L249:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v771)+4))
	v773 = F_objectGetVal(m, v772)
	mBase = m.M
	v774 = int32(_a1183)
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773))))
	if v777 != 0 {
		goto L252
	} else {
		goto L253
	}
L250:
	;
	if v809-v811 != 0 {
		goto L235
	} else {
		goto L262
	}
L251:
	;
	v809 = F_tolower(m, v805)
	mBase = m.M
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v806))))
	v811 = F_tolower(m, v810)
	mBase = m.M
	goto L250
L252:
	;
	v779 = v773
	v780 = v774
	v781 = v777
	goto L255
L253:
	;
	v805 = int32(0)
	v806 = v774
	goto L251
L254:
	;
	v805 = v802 & int32(255)
	v806 = v801
	goto L251
L255:
	;
	v783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780))))
	if v783 == int32(0) {
		v801 = v780
		v802 = v781
		goto L254
	} else {
		goto L257
	}
L256:
	;
	v801 = v795
	v802 = int32(0)
	goto L254
L257:
	;
	v787 = v781 & int32(255)
	if v787 == v783 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v794 = int32(1)
	v795 = v780 + v794
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779)+1)))
	if v796 != 0 {
		v779 = v779 + v794
		v780 = v795
		v781 = v796
		goto L255
	} else {
		goto L261
	}
L259:
	;
	v789 = F_tolower(m, v787)
	mBase = m.M
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780))))
	v791 = F_tolower(m, v790)
	mBase = m.M
	if v789 == v791 {
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779))))
	v801 = v780
	v802 = v793
	goto L254
L261:
	;
	goto L256
L262:
	;
	goto L236
L263:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)+8))
	v818 = F_objectGetVal(m, v817)
	mBase = m.M
	v819 = F_sentinelGetPrimaryByName(m, v818)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L20
	} else {
		goto L265
	}
L264:
	;
	v823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v819))))
	if v823&int32(64) == int32(0) {
		goto L269
	} else {
		goto L270
	}
L265:
	;
	if v819 != 0 {
		goto L264
	} else {
		goto L266
	}
L266:
	;
	F_addReplyNullArray(m, l0)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L20
	} else {
		goto L267
	}
L267:
	;
	goto L1
L268:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v835)+24))
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L20
	} else {
		goto L273
	}
L269:
	;
	v835 = v819
	goto L268
L270:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v819)+280))
	if v828 == int32(0) {
		goto L269
	} else {
		goto L271
	}
L271:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v819)+240))
	if int32(4) < v831 {
		v835 = v828
		goto L268
	} else {
		goto L272
	}
L272:
	;
	goto L269
L273:
	;
	v840 = int32(0)
	v841 = *(*int32)(unsafe.Add(mBase, _consts[639]))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v836+base.B2i32(v841 == v840)<<(uint(int32(2))%32))))
	F_addReplyBulkCString(m, l0, v847)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L20
	} else {
		goto L274
	}
L274:
	;
	v850 = int64(*(*int32)(unsafe.Add(mBase, uint32(v836)+8)))
	F_addReplyBulkLongLong(m, l0, v850)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L20
	} else {
		goto L275
	}
L275:
	;
	goto L1
L276:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1005)+4))
	v1007 = F_objectGetVal(m, v1006)
	mBase = m.M
	v1008 = int32(_a1184)
	v1011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007))))
	if v1011 != 0 {
		goto L331
	} else {
		goto L332
	}
L277:
	;
	if v891-v893 != 0 {
		goto L276
	} else {
		goto L289
	}
L278:
	;
	v891 = F_tolower(m, v887)
	mBase = m.M
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888))))
	v893 = F_tolower(m, v892)
	mBase = m.M
	goto L277
L279:
	;
	v861 = v855
	v862 = v856
	v863 = v859
	goto L282
L280:
	;
	v887 = int32(0)
	v888 = v856
	goto L278
L281:
	;
	v887 = v884 & int32(255)
	v888 = v883
	goto L278
L282:
	;
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862))))
	if v865 == int32(0) {
		v883 = v862
		v884 = v863
		goto L281
	} else {
		goto L284
	}
L283:
	;
	v883 = v877
	v884 = int32(0)
	goto L281
L284:
	;
	v869 = v863 & int32(255)
	if v869 == v865 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v876 = int32(1)
	v877 = v862 + v876
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+1)))
	if v878 != 0 {
		v861 = v861 + v876
		v862 = v877
		v863 = v878
		goto L282
	} else {
		goto L288
	}
L286:
	;
	v871 = F_tolower(m, v869)
	mBase = m.M
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862))))
	v873 = F_tolower(m, v872)
	mBase = m.M
	if v871 == v873 {
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861))))
	v883 = v862
	v884 = v875
	goto L281
L288:
	;
	goto L283
L289:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(v895+int32(-5)) < base.Ui32(int32(-2)) {
		goto L24
	} else {
		goto L290
	}
L290:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v900)+8))
	v902 = F_sentinelGetPrimaryByNameOrReplyError(m, l0, v901)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L20
	} else {
		goto L291
	}
L291:
	;
	if v902 == int32(0) {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v906 != int32(4) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902))))
	if v958&int32(64) == int32(0) {
		goto L312
	} else {
		goto L313
	}
L294:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v909)+12))
	v911 = F_objectGetVal(m, v910)
	mBase = m.M
	v912 = int32(_a1185)
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v911))))
	if v915 != 0 {
		goto L298
	} else {
		goto L299
	}
L295:
	;
	F_addReplyError(m, l0, int32(_a1186))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L20
	} else {
		goto L311
	}
L296:
	;
	if v947-v949 != 0 {
		goto L295
	} else {
		goto L308
	}
L297:
	;
	v947 = F_tolower(m, v943)
	mBase = m.M
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v944))))
	v949 = F_tolower(m, v948)
	mBase = m.M
	goto L296
L298:
	;
	v917 = v911
	v918 = v912
	v919 = v915
	goto L301
L299:
	;
	v943 = int32(0)
	v944 = v912
	goto L297
L300:
	;
	v943 = v940 & int32(255)
	v944 = v939
	goto L297
L301:
	;
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v918))))
	if v921 == int32(0) {
		v939 = v918
		v940 = v919
		goto L300
	} else {
		goto L303
	}
L302:
	;
	v939 = v933
	v940 = int32(0)
	goto L300
L303:
	;
	v925 = v919 & int32(255)
	if v925 == v921 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v932 = int32(1)
	v933 = v918 + v932
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v917)+1)))
	if v934 != 0 {
		v917 = v917 + v932
		v918 = v933
		v919 = v934
		goto L301
	} else {
		goto L307
	}
L305:
	;
	v927 = F_tolower(m, v925)
	mBase = m.M
	v928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v918))))
	v929 = F_tolower(m, v928)
	mBase = m.M
	if v927 == v929 {
		goto L304
	} else {
		goto L306
	}
L306:
	;
	v931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v917))))
	v939 = v918
	v940 = v931
	goto L300
L307:
	;
	goto L302
L308:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v902)+128))
	if v951 != 0 {
		goto L293
	} else {
		goto L309
	}
L309:
	;
	F_addReplyError(m, l0, int32(_a1187))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L20
	} else {
		goto L310
	}
L310:
	;
	goto L1
L311:
	;
	goto L1
L312:
	;
	v966 = F_sentinelSelectReplica(m, v902)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L20
	} else {
		goto L316
	}
L313:
	;
	F_addReplyError(m, l0, int32(_a1188))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L20
	} else {
		goto L314
	}
L314:
	;
	goto L1
L315:
	;
	v972 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v972 {
		goto L319
	} else {
		goto L320
	}
L316:
	;
	if v966 != 0 {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	F_addReplyError(m, l0, int32(_a1189))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L20
	} else {
		goto L318
	}
L318:
	;
	goto L1
L319:
	;
	v981 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v902)+56)) = v981
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v902)))
	*(*int32)(unsafe.Add(mBase, uint32(v902))) = v983 | int32(8)
	F_sentinelStartFailover(m, v902)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L20
	} else {
		goto L322
	}
L320:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v902)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v975
	F__serverLog(m, int32(2), int32(_a1190), v12)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L20
	} else {
		goto L321
	}
L321:
	;
	goto L319
L322:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v902)))
	if v906 != int32(4) {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v1002)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L20
	} else {
		goto L327
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v902))) = v989 | int32(2048)
	goto L323
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v902))) = v989 | int32(16384)
	F_sentinelAskPrimaryStateToOtherSentinels(m, v902, int32(1))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L20
	} else {
		goto L326
	}
L326:
	;
	goto L323
L327:
	;
	goto L1
L328:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+4))
	v1054 = F_objectGetVal(m, v1053)
	mBase = m.M
	v1055 = int32(_a1191)
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1054))))
	if v1058 != 0 {
		goto L347
	} else {
		goto L348
	}
L329:
	;
	if v1043-v1045 != 0 {
		goto L328
	} else {
		goto L341
	}
L330:
	;
	v1043 = F_tolower(m, v1039)
	mBase = m.M
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1040))))
	v1045 = F_tolower(m, v1044)
	mBase = m.M
	goto L329
L331:
	;
	v1013 = v1007
	v1014 = v1008
	v1015 = v1011
	goto L334
L332:
	;
	v1039 = int32(0)
	v1040 = v1008
	goto L330
L333:
	;
	v1039 = v1036 & int32(255)
	v1040 = v1035
	goto L330
L334:
	;
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014))))
	if v1017 == int32(0) {
		v1035 = v1014
		v1036 = v1015
		goto L333
	} else {
		goto L336
	}
L335:
	;
	v1035 = v1029
	v1036 = int32(0)
	goto L333
L336:
	;
	v1021 = v1015 & int32(255)
	if v1021 == v1017 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1028 = int32(1)
	v1029 = v1014 + v1028
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1013)+1)))
	if v1030 != 0 {
		v1013 = v1013 + v1028
		v1014 = v1029
		v1015 = v1030
		goto L334
	} else {
		goto L340
	}
L338:
	;
	v1023 = F_tolower(m, v1021)
	mBase = m.M
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014))))
	v1025 = F_tolower(m, v1024)
	mBase = m.M
	if v1023 == v1025 {
		goto L337
	} else {
		goto L339
	}
L339:
	;
	v1027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1013))))
	v1035 = v1014
	v1036 = v1027
	goto L333
L340:
	;
	goto L335
L341:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1047 != int32(2) {
		goto L24
	} else {
		goto L342
	}
L342:
	;
	F_sentinelPendingScriptsCommand(m, l0)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L20
	} else {
		goto L343
	}
L343:
	;
	goto L1
L344:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1181)+4))
	v1183 = F_objectGetVal(m, v1182)
	mBase = m.M
	v1184 = int32(_a1192)
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1183))))
	if v1187 != 0 {
		goto L389
	} else {
		goto L390
	}
L345:
	;
	if v1090-v1092 != 0 {
		goto L344
	} else {
		goto L357
	}
L346:
	;
	v1090 = F_tolower(m, v1086)
	mBase = m.M
	v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1087))))
	v1092 = F_tolower(m, v1091)
	mBase = m.M
	goto L345
L347:
	;
	v1060 = v1054
	v1061 = v1055
	v1062 = v1058
	goto L350
L348:
	;
	v1086 = int32(0)
	v1087 = v1055
	goto L346
L349:
	;
	v1086 = v1083 & int32(255)
	v1087 = v1082
	goto L346
L350:
	;
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1061))))
	if v1064 == int32(0) {
		v1082 = v1061
		v1083 = v1062
		goto L349
	} else {
		goto L352
	}
L351:
	;
	v1082 = v1076
	v1083 = int32(0)
	goto L349
L352:
	;
	v1068 = v1062 & int32(255)
	if v1068 == v1064 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1075 = int32(1)
	v1076 = v1061 + v1075
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1060)+1)))
	if v1077 != 0 {
		v1060 = v1060 + v1075
		v1061 = v1076
		v1062 = v1077
		goto L350
	} else {
		goto L356
	}
L354:
	;
	v1070 = F_tolower(m, v1068)
	mBase = m.M
	v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1061))))
	v1072 = F_tolower(m, v1071)
	mBase = m.M
	if v1070 == v1072 {
		goto L353
	} else {
		goto L355
	}
L355:
	;
	v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1060))))
	v1082 = v1061
	v1083 = v1074
	goto L349
L356:
	;
	goto L351
L357:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1094 != int32(6) {
		goto L24
	} else {
		goto L358
	}
L358:
	;
	v1099 = F_sentinelValidateArgs(m, l0, int32(2), int32(_a1193))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L20
	} else {
		goto L359
	}
L359:
	;
	if v1099 == int32(-1) {
		goto L1
	} else {
		goto L360
	}
L360:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1103)+20))
	v1108 = F_getLongFromObjectOrReply(m, l0, v1104, v12+int32(56), int32(_a1194))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L20
	} else {
		goto L361
	}
L361:
	;
	if v1108 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+16))
	v1115 = F_getLongFromObjectOrReply(m, l0, v1111, v12+int32(52), int32(_a1195))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L20
	} else {
		goto L363
	}
L363:
	;
	if v1115 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	if int32(0) < v1117 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1123 = int32(0)
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+12))
	v1126 = F_objectGetVal(m, v1125)
	mBase = m.M
	v1131 = *(*int32)(unsafe.Add(mBase, _consts[644]))
	v1134 = F_anetResolve(m, v1123, v1126, v12+int32(64), int32(46), base.B2i32(v1131 == v1123))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L20
	} else {
		goto L369
	}
L366:
	;
	F_addReplyError(m, l0, int32(_a1196))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L20
	} else {
		goto L367
	}
L367:
	;
	goto L1
L368:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1141)+8))
	v1143 = F_objectGetVal(m, v1142)
	mBase = m.M
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+12))
	v1147 = F_objectGetVal(m, v1146)
	mBase = m.M
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v1151 = F_createSentinelValkeyInstance(m, v1143, int32(1), v1147, v1148, v1149, int32(0))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L20
	} else {
		goto L373
	}
L369:
	;
	if v1134 != int32(-1) {
		goto L368
	} else {
		goto L370
	}
L370:
	;
	F_addReplyError(m, l0, int32(_a1197))
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L20
	} else {
		goto L371
	}
L371:
	;
	goto L1
L372:
	;
	F_sentinelFlushConfigAndReply(m, l0)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L20
	} else {
		goto L384
	}
L373:
	;
	if v1151 != 0 {
		goto L372
	} else {
		goto L374
	}
L374:
	;
	goto L379
L375:
	;
	F_addReplyError(m, l0, v1167)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L20
	} else {
		goto L383
	}
L376:
	;
	v1167 = int32(_a1198)
	goto L375
L377:
	;
	F_addReplyError(m, l0, int32(_a1199))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L20
	} else {
		goto L382
	}
L378:
	;
	if v1155 != int32(10) {
		goto L376
	} else {
		goto L380
	}
L379:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	switch v1155 + int32(-28) {
	case 0:
		goto L377
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L376
	case 16:
		v1167 = int32(_a1200)
		goto L375
	default:
		goto L378
	}
L380:
	;
	F_addReplyError(m, l0, int32(_a1201))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L20
	} else {
		goto L381
	}
L381:
	;
	goto L1
L382:
	;
	goto L1
L383:
	;
	goto L1
L384:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v1172
	F_sentinelEvent(m, int32(3), int32(_a1202), v1151, int32(_a1203), v12+int32(16))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L20
	} else {
		goto L385
	}
L385:
	;
	goto L1
L386:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1228)+4))
	v1230 = F_objectGetVal(m, v1229)
	mBase = m.M
	v1231 = int32(_a1204)
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1230))))
	if v1234 != 0 {
		goto L405
	} else {
		goto L406
	}
L387:
	;
	if v1219-v1221 != 0 {
		goto L386
	} else {
		goto L399
	}
L388:
	;
	v1219 = F_tolower(m, v1215)
	mBase = m.M
	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1216))))
	v1221 = F_tolower(m, v1220)
	mBase = m.M
	goto L387
L389:
	;
	v1189 = v1183
	v1190 = v1184
	v1191 = v1187
	goto L392
L390:
	;
	v1215 = int32(0)
	v1216 = v1184
	goto L388
L391:
	;
	v1215 = v1212 & int32(255)
	v1216 = v1211
	goto L388
L392:
	;
	v1193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1190))))
	if v1193 == int32(0) {
		v1211 = v1190
		v1212 = v1191
		goto L391
	} else {
		goto L394
	}
L393:
	;
	v1211 = v1205
	v1212 = int32(0)
	goto L391
L394:
	;
	v1197 = v1191 & int32(255)
	if v1197 == v1193 {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v1204 = int32(1)
	v1205 = v1190 + v1204
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1189)+1)))
	if v1206 != 0 {
		v1189 = v1189 + v1204
		v1190 = v1205
		v1191 = v1206
		goto L392
	} else {
		goto L398
	}
L396:
	;
	v1199 = F_tolower(m, v1197)
	mBase = m.M
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1190))))
	v1201 = F_tolower(m, v1200)
	mBase = m.M
	if v1199 == v1201 {
		goto L395
	} else {
		goto L397
	}
L397:
	;
	v1203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1189))))
	v1211 = v1190
	v1212 = v1203
	goto L391
L398:
	;
	goto L393
L399:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1223 != int32(2) {
		goto L24
	} else {
		goto L400
	}
L400:
	;
	F_sentinelFlushConfigAndReply(m, l0)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L20
	} else {
		goto L401
	}
L401:
	;
	goto L1
L402:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1294)+4))
	v1296 = F_objectGetVal(m, v1295)
	mBase = m.M
	v1297 = int32(_a1205)
	v1300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1296))))
	if v1300 != 0 {
		goto L425
	} else {
		goto L426
	}
L403:
	;
	if v1266-v1268 != 0 {
		goto L402
	} else {
		goto L415
	}
L404:
	;
	v1266 = F_tolower(m, v1262)
	mBase = m.M
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1263))))
	v1268 = F_tolower(m, v1267)
	mBase = m.M
	goto L403
L405:
	;
	v1236 = v1230
	v1237 = v1231
	v1238 = v1234
	goto L408
L406:
	;
	v1262 = int32(0)
	v1263 = v1231
	goto L404
L407:
	;
	v1262 = v1259 & int32(255)
	v1263 = v1258
	goto L404
L408:
	;
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1237))))
	if v1240 == int32(0) {
		v1258 = v1237
		v1259 = v1238
		goto L407
	} else {
		goto L410
	}
L409:
	;
	v1258 = v1252
	v1259 = int32(0)
	goto L407
L410:
	;
	v1244 = v1238 & int32(255)
	if v1244 == v1240 {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v1251 = int32(1)
	v1252 = v1237 + v1251
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1236)+1)))
	if v1253 != 0 {
		v1236 = v1236 + v1251
		v1237 = v1252
		v1238 = v1253
		goto L408
	} else {
		goto L414
	}
L412:
	;
	v1246 = F_tolower(m, v1244)
	mBase = m.M
	v1247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1237))))
	v1248 = F_tolower(m, v1247)
	mBase = m.M
	if v1246 == v1248 {
		goto L411
	} else {
		goto L413
	}
L413:
	;
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1236))))
	v1258 = v1237
	v1259 = v1250
	goto L407
L414:
	;
	goto L409
L415:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1270 != int32(3) {
		goto L24
	} else {
		goto L416
	}
L416:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+8))
	v1275 = F_sentinelGetPrimaryByNameOrReplyError(m, l0, v1274)
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L20
	} else {
		goto L417
	}
L417:
	;
	if v1275 == int32(0) {
		goto L1
	} else {
		goto L418
	}
L418:
	;
	F_sentinelEvent(m, int32(3), int32(_a1206), v1275, int32(_a1138), int32(0))
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L20
	} else {
		goto L419
	}
L419:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+8))
	v1289 = F_objectGetVal(m, v1288)
	mBase = m.M
	v1290 = F_dictDelete(m, v1286, v1289)
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L20
	} else {
		goto L420
	}
L420:
	;
	F_sentinelFlushConfigAndReply(m, l0)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L20
	} else {
		goto L421
	}
L421:
	;
	goto L1
L422:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1385)+4))
	v1387 = F_objectGetVal(m, v1386)
	mBase = m.M
	v1388 = int32(_a131)
	v1391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1387))))
	if v1391 != 0 {
		goto L459
	} else {
		goto L460
	}
L423:
	;
	if v1332-v1334 != 0 {
		goto L422
	} else {
		goto L435
	}
L424:
	;
	v1332 = F_tolower(m, v1328)
	mBase = m.M
	v1333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329))))
	v1334 = F_tolower(m, v1333)
	mBase = m.M
	goto L423
L425:
	;
	v1302 = v1296
	v1303 = v1297
	v1304 = v1300
	goto L428
L426:
	;
	v1328 = int32(0)
	v1329 = v1297
	goto L424
L427:
	;
	v1328 = v1325 & int32(255)
	v1329 = v1324
	goto L424
L428:
	;
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303))))
	if v1306 == int32(0) {
		v1324 = v1303
		v1325 = v1304
		goto L427
	} else {
		goto L430
	}
L429:
	;
	v1324 = v1318
	v1325 = int32(0)
	goto L427
L430:
	;
	v1310 = v1304 & int32(255)
	if v1310 == v1306 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1317 = int32(1)
	v1318 = v1303 + v1317
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1302)+1)))
	if v1319 != 0 {
		v1302 = v1302 + v1317
		v1303 = v1318
		v1304 = v1319
		goto L428
	} else {
		goto L434
	}
L432:
	;
	v1312 = F_tolower(m, v1310)
	mBase = m.M
	v1313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303))))
	v1314 = F_tolower(m, v1313)
	mBase = m.M
	if v1312 == v1314 {
		goto L431
	} else {
		goto L433
	}
L433:
	;
	v1316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1302))))
	v1324 = v1303
	v1325 = v1316
	goto L427
L434:
	;
	goto L429
L435:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1336 != int32(3) {
		goto L24
	} else {
		goto L436
	}
L436:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+8))
	v1341 = F_sentinelGetPrimaryByNameOrReplyError(m, l0, v1340)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L20
	} else {
		goto L437
	}
L437:
	;
	if v1341 == int32(0) {
		goto L1
	} else {
		goto L438
	}
L438:
	;
	v1347 = F_sentinelIsQuorumReachable(m, v1341, v12+int32(64))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L20
	} else {
		goto L439
	}
L439:
	;
	v1349 = F_sdsempty(m)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L20
	} else {
		goto L440
	}
L440:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	if v1347 != 0 {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v1351
	v1364 = F_sdscatfmt(m, v1349, int32(_a1207), v12+int32(48))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L20
	} else {
		goto L445
	}
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v1351
	v1356 = F_sdscatfmt(m, v1349, int32(_a1208), v12+int32(32))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L20
	} else {
		goto L443
	}
L443:
	;
	F_addReplySds(m, l0, v1356)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L20
	} else {
		goto L444
	}
L444:
	;
	goto L1
L445:
	;
	if v1347&int32(1) != 0 {
		goto L448
	} else {
		goto L449
	}
L446:
	;
	F_addReplyErrorSds(m, l0, v1382)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L20
	} else {
		goto L455
	}
L447:
	;
	v1380 = F_sdscat(m, v1378, int32(_a1209))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L20
	} else {
		goto L454
	}
L448:
	;
	v1371 = F_sdscat(m, v1364, int32(_a1210))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L20
	} else {
		goto L451
	}
L449:
	;
	if v1347 != int32(1) {
		v1378 = v1364
		goto L447
	} else {
		goto L450
	}
L450:
	;
	v1382 = v1364
	goto L446
L451:
	;
	if v1347 == int32(1) {
		v1382 = v1371
		goto L446
	} else {
		goto L452
	}
L452:
	;
	v1376 = F_sdscat(m, v1371, int32(_a1211))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L20
	} else {
		goto L453
	}
L453:
	;
	v1378 = v1376
	goto L447
L454:
	;
	v1382 = v1380
	goto L446
L455:
	;
	goto L1
L456:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1429)+4))
	v1431 = F_objectGetVal(m, v1430)
	mBase = m.M
	v1432 = int32(_a1212)
	v1435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1431))))
	if v1435 != 0 {
		goto L474
	} else {
		goto L475
	}
L457:
	;
	if v1423-v1425 != 0 {
		goto L456
	} else {
		goto L469
	}
L458:
	;
	v1423 = F_tolower(m, v1419)
	mBase = m.M
	v1424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1420))))
	v1425 = F_tolower(m, v1424)
	mBase = m.M
	goto L457
L459:
	;
	v1393 = v1387
	v1394 = v1388
	v1395 = v1391
	goto L462
L460:
	;
	v1419 = int32(0)
	v1420 = v1388
	goto L458
L461:
	;
	v1419 = v1416 & int32(255)
	v1420 = v1415
	goto L458
L462:
	;
	v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1394))))
	if v1397 == int32(0) {
		v1415 = v1394
		v1416 = v1395
		goto L461
	} else {
		goto L464
	}
L463:
	;
	v1415 = v1409
	v1416 = int32(0)
	goto L461
L464:
	;
	v1401 = v1395 & int32(255)
	if v1401 == v1397 {
		goto L465
	} else {
		goto L466
	}
L465:
	;
	v1408 = int32(1)
	v1409 = v1394 + v1408
	v1410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1393)+1)))
	if v1410 != 0 {
		v1393 = v1393 + v1408
		v1394 = v1409
		v1395 = v1410
		goto L462
	} else {
		goto L468
	}
L466:
	;
	v1403 = F_tolower(m, v1401)
	mBase = m.M
	v1404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1394))))
	v1405 = F_tolower(m, v1404)
	mBase = m.M
	if v1403 == v1405 {
		goto L465
	} else {
		goto L467
	}
L467:
	;
	v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1393))))
	v1415 = v1394
	v1416 = v1407
	goto L461
L468:
	;
	goto L463
L469:
	;
	F_sentinelSetCommand(m, l0)
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L20
	} else {
		goto L470
	}
L470:
	;
	goto L1
L471:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1571)+4))
	v1573 = F_objectGetVal(m, v1572)
	mBase = m.M
	v1574 = int32(_a1213)
	v1577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1573))))
	if v1577 != 0 {
		goto L522
	} else {
		goto L523
	}
L472:
	;
	if v1467-v1469 != 0 {
		goto L471
	} else {
		goto L484
	}
L473:
	;
	v1467 = F_tolower(m, v1463)
	mBase = m.M
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1464))))
	v1469 = F_tolower(m, v1468)
	mBase = m.M
	goto L472
L474:
	;
	v1437 = v1431
	v1438 = v1432
	v1439 = v1435
	goto L477
L475:
	;
	v1463 = int32(0)
	v1464 = v1432
	goto L473
L476:
	;
	v1463 = v1460 & int32(255)
	v1464 = v1459
	goto L473
L477:
	;
	v1441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1438))))
	if v1441 == int32(0) {
		v1459 = v1438
		v1460 = v1439
		goto L476
	} else {
		goto L479
	}
L478:
	;
	v1459 = v1453
	v1460 = int32(0)
	goto L476
L479:
	;
	v1445 = v1439 & int32(255)
	if v1445 == v1441 {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	v1452 = int32(1)
	v1453 = v1438 + v1452
	v1454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1437)+1)))
	if v1454 != 0 {
		v1437 = v1437 + v1452
		v1438 = v1453
		v1439 = v1454
		goto L477
	} else {
		goto L483
	}
L481:
	;
	v1447 = F_tolower(m, v1445)
	mBase = m.M
	v1448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1438))))
	v1449 = F_tolower(m, v1448)
	mBase = m.M
	if v1447 == v1449 {
		goto L480
	} else {
		goto L482
	}
L482:
	;
	v1451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1437))))
	v1459 = v1438
	v1460 = v1451
	goto L476
L483:
	;
	goto L478
L484:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1471 < int32(4) {
		goto L24
	} else {
		goto L485
	}
L485:
	;
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+8))
	v1476 = F_objectGetVal(m, v1475)
	mBase = m.M
	v1477 = int32(_a131)
	v1480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1476))))
	if v1480 != 0 {
		goto L489
	} else {
		goto L490
	}
L486:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1521)+8))
	v1523 = F_objectGetVal(m, v1522)
	mBase = m.M
	v1524 = int32(_a515)
	v1527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1523))))
	if v1527 != 0 {
		goto L505
	} else {
		goto L506
	}
L487:
	;
	if v1512-v1514 != 0 {
		goto L486
	} else {
		goto L499
	}
L488:
	;
	v1512 = F_tolower(m, v1508)
	mBase = m.M
	v1513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1509))))
	v1514 = F_tolower(m, v1513)
	mBase = m.M
	goto L487
L489:
	;
	v1482 = v1476
	v1483 = v1477
	v1484 = v1480
	goto L492
L490:
	;
	v1508 = int32(0)
	v1509 = v1477
	goto L488
L491:
	;
	v1508 = v1505 & int32(255)
	v1509 = v1504
	goto L488
L492:
	;
	v1486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1483))))
	if v1486 == int32(0) {
		v1504 = v1483
		v1505 = v1484
		goto L491
	} else {
		goto L494
	}
L493:
	;
	v1504 = v1498
	v1505 = int32(0)
	goto L491
L494:
	;
	v1490 = v1484 & int32(255)
	if v1490 == v1486 {
		goto L495
	} else {
		goto L496
	}
L495:
	;
	v1497 = int32(1)
	v1498 = v1483 + v1497
	v1499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1482)+1)))
	if v1499 != 0 {
		v1482 = v1482 + v1497
		v1483 = v1498
		v1484 = v1499
		goto L492
	} else {
		goto L498
	}
L496:
	;
	v1492 = F_tolower(m, v1490)
	mBase = m.M
	v1493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1483))))
	v1494 = F_tolower(m, v1493)
	mBase = m.M
	if v1492 == v1494 {
		goto L495
	} else {
		goto L497
	}
L497:
	;
	v1496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1482))))
	v1504 = v1483
	v1505 = v1496
	goto L491
L498:
	;
	goto L493
L499:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1516 < int32(5) {
		goto L486
	} else {
		goto L500
	}
L500:
	;
	F_sentinelConfigSetCommand(m, l0)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L20
	} else {
		goto L501
	}
L501:
	;
	goto L1
L502:
	;
	F_addReplyError(m, l0, int32(_a1214))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L20
	} else {
		goto L518
	}
L503:
	;
	if v1559-v1561 != 0 {
		goto L502
	} else {
		goto L515
	}
L504:
	;
	v1559 = F_tolower(m, v1555)
	mBase = m.M
	v1560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1556))))
	v1561 = F_tolower(m, v1560)
	mBase = m.M
	goto L503
L505:
	;
	v1529 = v1523
	v1530 = v1524
	v1531 = v1527
	goto L508
L506:
	;
	v1555 = int32(0)
	v1556 = v1524
	goto L504
L507:
	;
	v1555 = v1552 & int32(255)
	v1556 = v1551
	goto L504
L508:
	;
	v1533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1530))))
	if v1533 == int32(0) {
		v1551 = v1530
		v1552 = v1531
		goto L507
	} else {
		goto L510
	}
L509:
	;
	v1551 = v1545
	v1552 = int32(0)
	goto L507
L510:
	;
	v1537 = v1531 & int32(255)
	if v1537 == v1533 {
		goto L511
	} else {
		goto L512
	}
L511:
	;
	v1544 = int32(1)
	v1545 = v1530 + v1544
	v1546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1529)+1)))
	if v1546 != 0 {
		v1529 = v1529 + v1544
		v1530 = v1545
		v1531 = v1546
		goto L508
	} else {
		goto L514
	}
L512:
	;
	v1539 = F_tolower(m, v1537)
	mBase = m.M
	v1540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1530))))
	v1541 = F_tolower(m, v1540)
	mBase = m.M
	if v1539 == v1541 {
		goto L511
	} else {
		goto L513
	}
L513:
	;
	v1543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1529))))
	v1551 = v1530
	v1552 = v1543
	goto L507
L514:
	;
	goto L509
L515:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1563 < int32(4) {
		goto L502
	} else {
		goto L516
	}
L516:
	;
	F_sentinelConfigGetCommand(m, l0)
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L20
	} else {
		goto L517
	}
L517:
	;
	goto L1
L518:
	;
	goto L1
L519:
	;
	v2220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2220)+4))
	v2222 = F_objectGetVal(m, v2221)
	mBase = m.M
	v2223 = int32(_a1215)
	v2226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2222))))
	if v2226 != 0 {
		goto L696
	} else {
		goto L697
	}
L520:
	;
	if v1609-v1611 != 0 {
		goto L519
	} else {
		goto L532
	}
L521:
	;
	v1609 = F_tolower(m, v1605)
	mBase = m.M
	v1610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1606))))
	v1611 = F_tolower(m, v1610)
	mBase = m.M
	goto L520
L522:
	;
	v1579 = v1573
	v1580 = v1574
	v1581 = v1577
	goto L525
L523:
	;
	v1605 = int32(0)
	v1606 = v1574
	goto L521
L524:
	;
	v1605 = v1602 & int32(255)
	v1606 = v1601
	goto L521
L525:
	;
	v1583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1580))))
	if v1583 == int32(0) {
		v1601 = v1580
		v1602 = v1581
		goto L524
	} else {
		goto L527
	}
L526:
	;
	v1601 = v1595
	v1602 = int32(0)
	goto L524
L527:
	;
	v1587 = v1581 & int32(255)
	if v1587 == v1583 {
		goto L528
	} else {
		goto L529
	}
L528:
	;
	v1594 = int32(1)
	v1595 = v1580 + v1594
	v1596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1579)+1)))
	if v1596 != 0 {
		v1579 = v1579 + v1594
		v1580 = v1595
		v1581 = v1596
		goto L525
	} else {
		goto L531
	}
L529:
	;
	v1589 = F_tolower(m, v1587)
	mBase = m.M
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1580))))
	v1591 = F_tolower(m, v1590)
	mBase = m.M
	if v1589 == v1591 {
		goto L528
	} else {
		goto L530
	}
L530:
	;
	v1593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1579))))
	v1601 = v1580
	v1602 = v1593
	goto L524
L531:
	;
	goto L526
L532:
	;
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1613 < int32(2) {
		goto L24
	} else {
		goto L533
	}
L533:
	;
	v1616 = F_mstime(m)
	mBase = m.M
	v1618 = v12 + int32(80)
	v1619 = int32(0)
	v1620 = *(*int64)(unsafe.Add(mBase, _consts[646]))
	*(*int64)(unsafe.Add(mBase, uint32(v1618))) = v1620
	v1625 = *(*int32)(unsafe.Add(mBase, _consts[647]))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(96)))) = v1625
	v1630 = *(*int64)(unsafe.Add(mBase, _consts[648]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(88)))) = v1630
	v1635 = *(*int64)(unsafe.Add(mBase, _consts[649]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(72)))) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v1618))) = v1619
	v1640 = *(*int64)(unsafe.Add(mBase, _consts[650]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v1640
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(3) <= v1642 {
		goto L535
	} else {
		goto L536
	}
L534:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1688)+16))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1688)+12))
	F_addReplyArrayLen(m, l0, (v1690+v1691)<<(uint(int32(1))%32))
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L20
	} else {
		goto L546
	}
L535:
	;
	v1649 = F_dictCreate(m, v12+int32(64))
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L20
	} else {
		goto L537
	}
L536:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	v1688 = v1646
	goto L534
L537:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1651 < int32(3) {
		v1688 = v1649
		goto L534
	} else {
		goto L538
	}
L538:
	;
	v1657 = int32(2)
	goto L539
L539:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1664+v1657<<(uint(int32(2))%32))))
	v1669 = F_objectGetVal(m, v1668)
	mBase = m.M
	v1670 = F_sentinelGetPrimaryByName(m, v1669)
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L20
	} else {
		goto L542
	}
L540:
	;
	v1688 = v1649
	goto L534
L541:
	;
	v1678 = v1657 + int32(1)
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1678 < v1679 {
		v1657 = v1678
		goto L539
	} else {
		goto L545
	}
L542:
	;
	if v1670 == int32(0) {
		goto L541
	} else {
		goto L543
	}
L543:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1670)+4))
	v1675 = F_dictAdd(m, v1649, v1674, v1670)
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L20
	} else {
		goto L544
	}
L544:
	;
	goto L541
L545:
	;
	goto L540
L546:
	;
	v1697 = F_dictGetIterator(m, v1688)
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L20
	} else {
		goto L548
	}
L547:
	;
	F_dictReleaseIterator(m, v1697)
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L20
	} else {
		goto L690
	}
L548:
	;
	v1706 = v1697 + int32(20)
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+16))
	if v1707 != 0 {
		goto L551
	} else {
		goto L552
	}
L549:
	;
	if v1802 == int32(0) {
		goto L547
	} else {
		goto L575
	}
L550:
	;
	v1713 = v1706
	v1714 = v1710
	goto L553
L551:
	;
	v1710 = int32(1)
	goto L550
L552:
	;
	v1710 = int32(0)
	goto L550
L553:
	;
	switch v1714 {
	case 0:
		goto L558
	default:
		goto L557
	}
L555:
	;
	v1714 = int32(0)
	goto L553
L556:
	;
	goto L549
L557:
	;
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(v1713)))
	*(*int32)(unsafe.Add(mBase, uint32(v1697)+16)) = v1794
	if v1794 == int32(0) {
		goto L555
	} else {
		goto L574
	}
L558:
	;
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+4))
	if v1718 != int32(-1) {
		v1757 = v1718
		goto L559
	} else {
		goto L560
	}
L559:
	;
	v1758 = int32(1)
	v1759 = v1757 + v1758
	*(*int32)(unsafe.Add(mBase, uint32(v1697)+4)) = v1759
	v1761 = int32(0)
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1697)))
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+8))
	v1769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1764+v1765+int32(26)))))
	if v1769 == int32(255) {
		goto L568
	} else {
		goto L569
	}
L560:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+8))
	if v1722 != 0 {
		v1757 = int32(-1)
		goto L559
	} else {
		goto L561
	}
L561:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1697)))
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+12))
	if v1724 == int32(0) {
		goto L563
	} else {
		goto L564
	}
L562:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1750)+20))
	if v1751 != int32(-1) {
		goto L565
	} else {
		goto L566
	}
L563:
	;
	v1731 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1723)+16)))
	v1732 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1723)+27)))
	v1733 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1723)+8)))
	v1734 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1723)+12)))
	v1735 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1723)+26)))
	v1736 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1723)+4)))
	v1737 = F_wangHash64(m, v1736)
	mBase = m.M
	v1739 = F_wangHash64(m, v1735+v1737)
	mBase = m.M
	v1741 = F_wangHash64(m, v1734+v1739)
	mBase = m.M
	v1743 = F_wangHash64(m, v1733+v1741)
	mBase = m.M
	v1745 = F_wangHash64(m, v1732+v1743)
	mBase = m.M
	v1747 = F_wangHash64(m, v1731+v1745)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v1697)+24)) = v1747
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1697)))
	v1750 = v1749
	goto L562
L564:
	;
	v1727 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1723)+24)))
	v1729 = v1727 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1723)+24)) = uint16(v1729)
	v1750 = v1723
	goto L562
L565:
	;
	v1757 = v1751 + int32(-1)
	goto L559
L566:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+4))
	v1757 = v1754
	goto L559
L567:
	;
	v1784 = int32(2)
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1764+v1782<<(uint(v1784)%32)+int32(4))))
	v1713 = v1789 + v1783<<(uint(v1784)%32)
	v1714 = int32(1)
	goto L553
L568:
	;
	v1773 = v1761
	goto L570
L569:
	;
	v1773 = v1758 << (uint(v1769) % 32)
	goto L570
L570:
	;
	if v1759 < v1773 {
		v1782 = v1765
		v1783 = v1759
		goto L567
	} else {
		goto L571
	}
L571:
	;
	if v1765 != 0 {
		v1802 = v1761
		goto L556
	} else {
		goto L572
	}
L572:
	;
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(v1764)+20))
	if v1775 == int32(-1) {
		v1802 = v1761
		goto L556
	} else {
		goto L573
	}
L573:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1697)+4)) = int64(4294967296)
	v1782 = int32(1)
	v1783 = int32(0)
	goto L567
L574:
	;
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1794)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1706))) = v1798
	v1802 = v1794
	goto L556
L575:
	;
	v1810 = v1802
	goto L576
L576:
	;
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v1810)+8))
	goto L578
L577:
	;
	goto L547
L578:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1817)+4))
	if v1818&int32(3) == int32(0) {
		v1840 = v1818
		goto L581
	} else {
		goto L582
	}
L579:
	;
	F_addReplyBulkCBuffer(m, l0, v1818, v1873)
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
		goto L20
	} else {
		goto L595
	}
L580:
	;
	v1873 = v1865 - v1818
	goto L579
L581:
	;
	v1844 = v1840
	goto L589
L582:
	;
	v1826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1818))))
	if v1826 != 0 {
		goto L583
	} else {
		goto L584
	}
L583:
	;
	v1829 = v1818
	goto L585
L584:
	;
	v1873 = v1818 - v1818
	goto L579
L585:
	;
	v1833 = v1829 + int32(1)
	if v1833&int32(3) == int32(0) {
		v1840 = v1833
		goto L581
	} else {
		goto L587
	}
L587:
	;
	v1838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1833))))
	if v1838 != 0 {
		v1829 = v1833
		goto L585
	} else {
		goto L588
	}
L588:
	;
	v1865 = v1833
	goto L580
L589:
	;
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v1844)))
	v1853 = int32(-2139062144)
	if (int32(16843008)-v1850|v1850)&v1853 == v1853 {
		v1844 = v1844 + int32(4)
		goto L589
	} else {
		goto L591
	}
L590:
	;
	v1859 = v1844
	goto L592
L591:
	;
	goto L590
L592:
	;
	v1863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1859))))
	if v1863 != 0 {
		v1859 = v1859 + int32(1)
		goto L592
	} else {
		goto L594
	}
L593:
	;
	v1865 = v1859
	goto L580
L594:
	;
	goto L593
L595:
	;
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1817)+148))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1876)+12))
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1876)+16))
	F_addReplyArrayLen(m, l0, v1877+v1878+int32(1))
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L20
	} else {
		goto L596
	}
L596:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L20
	} else {
		goto L597
	}
L597:
	;
	v1887 = int64(0)
	v1888 = *(*int64)(unsafe.Add(mBase, uint32(v1817)+96))
	if v1888 == v1887 {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	v1892 = v1887
	goto L600
L599:
	;
	v1892 = v1616 - v1888
	goto L600
L600:
	;
	F_addReplyLongLong(m, l0, v1892)
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L20
	} else {
		goto L601
	}
L601:
	;
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1817)+292))
	if v1895 == int32(0) {
		goto L603
	} else {
		goto L604
	}
L602:
	;
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1817)+148))
	v1928 = F_dictGetIterator(m, v1927)
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		goto L20
	} else {
		goto L614
	}
L603:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L20
	} else {
		goto L613
	}
L604:
	;
	v1903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1895+int32(-1)))))
	switch v1903 & int32(7) {
	case 0:
		goto L611
	case 1:
		goto L610
	case 2:
		goto L609
	case 3:
		goto L608
	case 4:
		goto L607
	default:
		v1920 = int32(0)
		goto L606
	}
L605:
	;
	F_addReplyBulkCBuffer(m, l0, v1895, v1922)
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L20
	} else {
		goto L612
	}
L606:
	;
	v1922 = v1920
	goto L605
L607:
	;
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v1895+int32(-17))))
	v1920 = v1919
	goto L606
L608:
	;
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v1895+int32(-9))))
	v1922 = v1916
	goto L605
L609:
	;
	v1913 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1895+int32(-5)))))
	v1922 = v1913
	goto L605
L610:
	;
	v1910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1895+int32(-3)))))
	v1922 = v1910
	goto L605
L611:
	;
	v1922 = int32(base.Ui32(v1903) >> (uint(int32(3)) % 32))
	goto L605
L612:
	;
	goto L602
L613:
	;
	goto L602
L614:
	;
	goto L616
L615:
	;
	F_dictReleaseIterator(m, v1928)
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L20
	} else {
		goto L662
	}
L616:
	;
	v1946 = v1928 + int32(20)
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+16))
	if v1947 != 0 {
		goto L620
	} else {
		goto L621
	}
L618:
	;
	if v2042 == int32(0) {
		goto L615
	} else {
		goto L644
	}
L619:
	;
	v1953 = v1946
	v1954 = v1950
	goto L622
L620:
	;
	v1950 = int32(1)
	goto L619
L621:
	;
	v1950 = int32(0)
	goto L619
L622:
	;
	switch v1954 {
	case 0:
		goto L627
	default:
		goto L626
	}
L624:
	;
	v1954 = int32(0)
	goto L622
L625:
	;
	goto L618
L626:
	;
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v1953)))
	*(*int32)(unsafe.Add(mBase, uint32(v1928)+16)) = v2034
	if v2034 == int32(0) {
		goto L624
	} else {
		goto L643
	}
L627:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+4))
	if v1958 != int32(-1) {
		v1997 = v1958
		goto L628
	} else {
		goto L629
	}
L628:
	;
	v1998 = int32(1)
	v1999 = v1997 + v1998
	*(*int32)(unsafe.Add(mBase, uint32(v1928)+4)) = v1999
	v2001 = int32(0)
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v1928)))
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+8))
	v2009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2004+v2005+int32(26)))))
	if v2009 == int32(255) {
		goto L637
	} else {
		goto L638
	}
L629:
	;
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+8))
	if v1962 != 0 {
		v1997 = int32(-1)
		goto L628
	} else {
		goto L630
	}
L630:
	;
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(v1928)))
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+12))
	if v1964 == int32(0) {
		goto L632
	} else {
		goto L633
	}
L631:
	;
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+20))
	if v1991 != int32(-1) {
		goto L634
	} else {
		goto L635
	}
L632:
	;
	v1971 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1963)+16)))
	v1972 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1963)+27)))
	v1973 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1963)+8)))
	v1974 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1963)+12)))
	v1975 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1963)+26)))
	v1976 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1963)+4)))
	v1977 = F_wangHash64(m, v1976)
	mBase = m.M
	v1979 = F_wangHash64(m, v1975+v1977)
	mBase = m.M
	v1981 = F_wangHash64(m, v1974+v1979)
	mBase = m.M
	v1983 = F_wangHash64(m, v1973+v1981)
	mBase = m.M
	v1985 = F_wangHash64(m, v1972+v1983)
	mBase = m.M
	v1987 = F_wangHash64(m, v1971+v1985)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v1928)+24)) = v1987
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(v1928)))
	v1990 = v1989
	goto L631
L633:
	;
	v1967 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1963)+24)))
	v1969 = v1967 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1963)+24)) = uint16(v1969)
	v1990 = v1963
	goto L631
L634:
	;
	v1997 = v1991 + int32(-1)
	goto L628
L635:
	;
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+4))
	v1997 = v1994
	goto L628
L636:
	;
	v2024 = int32(2)
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(v2004+v2022<<(uint(v2024)%32)+int32(4))))
	v1953 = v2029 + v2023<<(uint(v2024)%32)
	v1954 = int32(1)
	goto L622
L637:
	;
	v2013 = v2001
	goto L639
L638:
	;
	v2013 = v1998 << (uint(v2009) % 32)
	goto L639
L639:
	;
	if v1999 < v2013 {
		v2022 = v2005
		v2023 = v1999
		goto L636
	} else {
		goto L640
	}
L640:
	;
	if v2005 != 0 {
		v2042 = v2001
		goto L625
	} else {
		goto L641
	}
L641:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v2004)+20))
	if v2015 == int32(-1) {
		v2042 = v2001
		goto L625
	} else {
		goto L642
	}
L642:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1928)+4)) = int64(4294967296)
	v2022 = int32(1)
	v2023 = int32(0)
	goto L636
L643:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v2034)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1946))) = v2038
	v2042 = v2034
	goto L625
L644:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+8))
	goto L645
L645:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v2051 = m.ExcPending
	if v2051 != 0 {
		goto L20
	} else {
		goto L646
	}
L646:
	;
	v2052 = *(*int64)(unsafe.Add(mBase, uint32(v1817)+96))
	if base.B2i32(v2052 == int64(0)) == int32(0) {
		goto L648
	} else {
		goto L649
	}
L647:
	;
	F_addReplyLongLong(m, l0, v2060)
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L20
	} else {
		goto L650
	}
L648:
	;
	v2058 = *(*int64)(unsafe.Add(mBase, uint32(v2048)+96))
	v2060 = v1616 - v2058
	goto L647
L649:
	;
	v2060 = int64(0)
	goto L647
L650:
	;
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(v2048)+292))
	if v2063 == int32(0) {
		goto L651
	} else {
		goto L652
	}
L651:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L20
	} else {
		goto L661
	}
L652:
	;
	v2071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2063+int32(-1)))))
	switch v2071 & int32(7) {
	case 0:
		goto L659
	case 1:
		goto L658
	case 2:
		goto L657
	case 3:
		goto L656
	case 4:
		goto L655
	default:
		v2088 = int32(0)
		goto L654
	}
L653:
	;
	F_addReplyBulkCBuffer(m, l0, v2063, v2090)
	mBase = m.M
	v2092 = m.ExcPending
	if v2092 != 0 {
		goto L20
	} else {
		goto L660
	}
L654:
	;
	v2090 = v2088
	goto L653
L655:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v2063+int32(-17))))
	v2088 = v2087
	goto L654
L656:
	;
	v2084 = *(*int32)(unsafe.Add(mBase, uint32(v2063+int32(-9))))
	v2090 = v2084
	goto L653
L657:
	;
	v2081 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2063+int32(-5)))))
	v2090 = v2081
	goto L653
L658:
	;
	v2078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2063+int32(-3)))))
	v2090 = v2078
	goto L653
L659:
	;
	v2090 = int32(base.Ui32(v2071) >> (uint(int32(3)) % 32))
	goto L653
L660:
	;
	goto L616
L661:
	;
	goto L616
L662:
	;
	v2104 = v1697 + int32(20)
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+16))
	if v2105 != 0 {
		goto L665
	} else {
		goto L666
	}
L663:
	;
	if v2200 != 0 {
		v1810 = v2200
		goto L576
	} else {
		goto L689
	}
L664:
	;
	v2111 = v2104
	v2112 = v2108
	goto L667
L665:
	;
	v2108 = int32(1)
	goto L664
L666:
	;
	v2108 = int32(0)
	goto L664
L667:
	;
	switch v2112 {
	case 0:
		goto L672
	default:
		goto L671
	}
L669:
	;
	v2112 = int32(0)
	goto L667
L670:
	;
	goto L663
L671:
	;
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v2111)))
	*(*int32)(unsafe.Add(mBase, uint32(v1697)+16)) = v2192
	if v2192 == int32(0) {
		goto L669
	} else {
		goto L688
	}
L672:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+4))
	if v2116 != int32(-1) {
		v2155 = v2116
		goto L673
	} else {
		goto L674
	}
L673:
	;
	v2156 = int32(1)
	v2157 = v2155 + v2156
	*(*int32)(unsafe.Add(mBase, uint32(v1697)+4)) = v2157
	v2159 = int32(0)
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v1697)))
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+8))
	v2167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2162+v2163+int32(26)))))
	if v2167 == int32(255) {
		goto L682
	} else {
		goto L683
	}
L674:
	;
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+8))
	if v2120 != 0 {
		v2155 = int32(-1)
		goto L673
	} else {
		goto L675
	}
L675:
	;
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v1697)))
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+12))
	if v2122 == int32(0) {
		goto L677
	} else {
		goto L678
	}
L676:
	;
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v2148)+20))
	if v2149 != int32(-1) {
		goto L679
	} else {
		goto L680
	}
L677:
	;
	v2129 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2121)+16)))
	v2130 = int64(*(*int8)(unsafe.Add(mBase, uint32(v2121)+27)))
	v2131 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2121)+8)))
	v2132 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2121)+12)))
	v2133 = int64(*(*int8)(unsafe.Add(mBase, uint32(v2121)+26)))
	v2134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2121)+4)))
	v2135 = F_wangHash64(m, v2134)
	mBase = m.M
	v2137 = F_wangHash64(m, v2133+v2135)
	mBase = m.M
	v2139 = F_wangHash64(m, v2132+v2137)
	mBase = m.M
	v2141 = F_wangHash64(m, v2131+v2139)
	mBase = m.M
	v2143 = F_wangHash64(m, v2130+v2141)
	mBase = m.M
	v2145 = F_wangHash64(m, v2129+v2143)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v1697)+24)) = v2145
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v1697)))
	v2148 = v2147
	goto L676
L678:
	;
	v2125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2121)+24)))
	v2127 = v2125 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2121)+24)) = uint16(v2127)
	v2148 = v2121
	goto L676
L679:
	;
	v2155 = v2149 + int32(-1)
	goto L673
L680:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+4))
	v2155 = v2152
	goto L673
L681:
	;
	v2182 = int32(2)
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v2162+v2180<<(uint(v2182)%32)+int32(4))))
	v2111 = v2187 + v2181<<(uint(v2182)%32)
	v2112 = int32(1)
	goto L667
L682:
	;
	v2171 = v2159
	goto L684
L683:
	;
	v2171 = v2156 << (uint(v2167) % 32)
	goto L684
L684:
	;
	if v2157 < v2171 {
		v2180 = v2163
		v2181 = v2157
		goto L681
	} else {
		goto L685
	}
L685:
	;
	if v2163 != 0 {
		v2200 = v2159
		goto L670
	} else {
		goto L686
	}
L686:
	;
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v2162)+20))
	if v2173 == int32(-1) {
		v2200 = v2159
		goto L670
	} else {
		goto L687
	}
L687:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1697)+4)) = int64(4294967296)
	v2180 = int32(1)
	v2181 = int32(0)
	goto L681
L688:
	;
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(v2192)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2104))) = v2196
	v2200 = v2192
	goto L670
L689:
	;
	goto L577
L690:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	if v1688 == v2216 {
		goto L1
	} else {
		goto L691
	}
L691:
	;
	F_dictRelease(m, v1688)
	mBase = m.M
	v2219 = m.ExcPending
	if v2219 != 0 {
		goto L20
	} else {
		goto L692
	}
L692:
	;
	goto L1
L693:
	;
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v2468)+4))
	v2470 = F_objectGetVal(m, v2469)
	mBase = m.M
	v2471 = int32(_a713)
	v2474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2470))))
	if v2474 != 0 {
		goto L767
	} else {
		goto L768
	}
L694:
	;
	if v2258-v2260 != 0 {
		goto L693
	} else {
		goto L706
	}
L695:
	;
	v2258 = F_tolower(m, v2254)
	mBase = m.M
	v2259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2255))))
	v2260 = F_tolower(m, v2259)
	mBase = m.M
	goto L694
L696:
	;
	v2228 = v2222
	v2229 = v2223
	v2230 = v2226
	goto L699
L697:
	;
	v2254 = int32(0)
	v2255 = v2223
	goto L695
L698:
	;
	v2254 = v2251 & int32(255)
	v2255 = v2250
	goto L695
L699:
	;
	v2232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2229))))
	if v2232 == int32(0) {
		v2250 = v2229
		v2251 = v2230
		goto L698
	} else {
		goto L701
	}
L700:
	;
	v2250 = v2244
	v2251 = int32(0)
	goto L698
L701:
	;
	v2236 = v2230 & int32(255)
	if v2236 == v2232 {
		goto L702
	} else {
		goto L703
	}
L702:
	;
	v2243 = int32(1)
	v2244 = v2229 + v2243
	v2245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2228)+1)))
	if v2245 != 0 {
		v2228 = v2228 + v2243
		v2229 = v2244
		v2230 = v2245
		goto L699
	} else {
		goto L705
	}
L703:
	;
	v2238 = F_tolower(m, v2236)
	mBase = m.M
	v2239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2229))))
	v2240 = F_tolower(m, v2239)
	mBase = m.M
	if v2238 == v2240 {
		goto L702
	} else {
		goto L704
	}
L704:
	;
	v2242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2228))))
	v2250 = v2229
	v2251 = v2242
	goto L698
L705:
	;
	goto L700
L706:
	;
	v2262 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[651])) = v2262
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v2265 < int32(3) {
		goto L708
	} else {
		goto L709
	}
L707:
	;
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2415 = *(*int32)(unsafe.Add(mBase, uint32(v2411+v2271<<(uint(int32(2))%32))))
	v2416 = F_objectGetVal(m, v2415)
	mBase = m.M
	v2417 = int32(_a757)
	v2420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2416))))
	if v2420 != 0 {
		goto L749
	} else {
		goto L750
	}
L708:
	;
	v2408 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v2408)
	mBase = m.M
	v2410 = m.ExcPending
	if v2410 != 0 {
		goto L20
	} else {
		goto L745
	}
L709:
	;
	v2271 = int32(2)
	goto L710
L710:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2280 = v2271 << (uint(int32(2)) % 32)
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(v2278+v2280)))
	v2283 = F_objectGetVal(m, v2282)
	mBase = m.M
	v2284 = int32(_a1216)
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2283))))
	if v2287 != 0 {
		goto L717
	} else {
		goto L718
	}
L711:
	;
	goto L708
L712:
	;
	v2395 = v2271 + int32(1)
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v2395 < v2396 {
		v2271 = v2395
		goto L710
	} else {
		goto L744
	}
L713:
	;
	F__serverLog(m, int32(3), v2388, int32(0))
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L20
	} else {
		goto L743
	}
L714:
	;
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(v2334+v2280)))
	v2337 = F_objectGetVal(m, v2336)
	mBase = m.M
	v2338 = int32(_a1217)
	v2341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2337))))
	if v2341 != 0 {
		goto L731
	} else {
		goto L732
	}
L715:
	;
	if v2319-v2321 != 0 {
		goto L714
	} else {
		goto L727
	}
L716:
	;
	v2319 = F_tolower(m, v2315)
	mBase = m.M
	v2320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2316))))
	v2321 = F_tolower(m, v2320)
	mBase = m.M
	goto L715
L717:
	;
	v2289 = v2283
	v2290 = v2284
	v2291 = v2287
	goto L720
L718:
	;
	v2315 = int32(0)
	v2316 = v2284
	goto L716
L719:
	;
	v2315 = v2312 & int32(255)
	v2316 = v2311
	goto L716
L720:
	;
	v2293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2290))))
	if v2293 == int32(0) {
		v2311 = v2290
		v2312 = v2291
		goto L719
	} else {
		goto L722
	}
L721:
	;
	v2311 = v2305
	v2312 = int32(0)
	goto L719
L722:
	;
	v2297 = v2291 & int32(255)
	if v2297 == v2293 {
		goto L723
	} else {
		goto L724
	}
L723:
	;
	v2304 = int32(1)
	v2305 = v2290 + v2304
	v2306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2289)+1)))
	if v2306 != 0 {
		v2289 = v2289 + v2304
		v2290 = v2305
		v2291 = v2306
		goto L720
	} else {
		goto L726
	}
L724:
	;
	v2299 = F_tolower(m, v2297)
	mBase = m.M
	v2300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2290))))
	v2301 = F_tolower(m, v2300)
	mBase = m.M
	if v2299 == v2301 {
		goto L723
	} else {
		goto L725
	}
L725:
	;
	v2303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2289))))
	v2311 = v2290
	v2312 = v2303
	goto L719
L726:
	;
	goto L721
L727:
	;
	v2323 = int32(0)
	v2325 = *(*int32)(unsafe.Add(mBase, _consts[651]))
	*(*int32)(unsafe.Add(mBase, _consts[651])) = v2325 | int32(1)
	v2330 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v2330 {
		goto L712
	} else {
		goto L728
	}
L728:
	;
	v2388 = int32(_a1218)
	goto L713
L729:
	;
	if v2373-v2375 != 0 {
		goto L707
	} else {
		goto L741
	}
L730:
	;
	v2373 = F_tolower(m, v2369)
	mBase = m.M
	v2374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2370))))
	v2375 = F_tolower(m, v2374)
	mBase = m.M
	goto L729
L731:
	;
	v2343 = v2337
	v2344 = v2338
	v2345 = v2341
	goto L734
L732:
	;
	v2369 = int32(0)
	v2370 = v2338
	goto L730
L733:
	;
	v2369 = v2366 & int32(255)
	v2370 = v2365
	goto L730
L734:
	;
	v2347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2344))))
	if v2347 == int32(0) {
		v2365 = v2344
		v2366 = v2345
		goto L733
	} else {
		goto L736
	}
L735:
	;
	v2365 = v2359
	v2366 = int32(0)
	goto L733
L736:
	;
	v2351 = v2345 & int32(255)
	if v2351 == v2347 {
		goto L737
	} else {
		goto L738
	}
L737:
	;
	v2358 = int32(1)
	v2359 = v2344 + v2358
	v2360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2343)+1)))
	if v2360 != 0 {
		v2343 = v2343 + v2358
		v2344 = v2359
		v2345 = v2360
		goto L734
	} else {
		goto L740
	}
L738:
	;
	v2353 = F_tolower(m, v2351)
	mBase = m.M
	v2354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2344))))
	v2355 = F_tolower(m, v2354)
	mBase = m.M
	if v2353 == v2355 {
		goto L737
	} else {
		goto L739
	}
L739:
	;
	v2357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2343))))
	v2365 = v2344
	v2366 = v2357
	goto L733
L740:
	;
	goto L735
L741:
	;
	v2377 = int32(0)
	v2379 = *(*int32)(unsafe.Add(mBase, _consts[651]))
	*(*int32)(unsafe.Add(mBase, _consts[651])) = v2379 | int32(2)
	v2384 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v2384 {
		goto L712
	} else {
		goto L742
	}
L742:
	;
	v2388 = int32(_a1219)
	goto L713
L743:
	;
	goto L712
L744:
	;
	goto L711
L745:
	;
	goto L1
L746:
	;
	F_addReplyError(m, l0, int32(_a1220))
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L20
	} else {
		goto L763
	}
L747:
	;
	if v2452-v2454 != 0 {
		goto L746
	} else {
		goto L759
	}
L748:
	;
	v2452 = F_tolower(m, v2448)
	mBase = m.M
	v2453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2449))))
	v2454 = F_tolower(m, v2453)
	mBase = m.M
	goto L747
L749:
	;
	v2422 = v2416
	v2423 = v2417
	v2424 = v2420
	goto L752
L750:
	;
	v2448 = int32(0)
	v2449 = v2417
	goto L748
L751:
	;
	v2448 = v2445 & int32(255)
	v2449 = v2444
	goto L748
L752:
	;
	v2426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2423))))
	if v2426 == int32(0) {
		v2444 = v2423
		v2445 = v2424
		goto L751
	} else {
		goto L754
	}
L753:
	;
	v2444 = v2438
	v2445 = int32(0)
	goto L751
L754:
	;
	v2430 = v2424 & int32(255)
	if v2430 == v2426 {
		goto L755
	} else {
		goto L756
	}
L755:
	;
	v2437 = int32(1)
	v2438 = v2423 + v2437
	v2439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2422)+1)))
	if v2439 != 0 {
		v2422 = v2422 + v2437
		v2423 = v2438
		v2424 = v2439
		goto L752
	} else {
		goto L758
	}
L756:
	;
	v2432 = F_tolower(m, v2430)
	mBase = m.M
	v2433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2423))))
	v2434 = F_tolower(m, v2433)
	mBase = m.M
	if v2432 == v2434 {
		goto L755
	} else {
		goto L757
	}
L757:
	;
	v2436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2422))))
	v2444 = v2423
	v2445 = v2436
	goto L751
L758:
	;
	goto L753
L759:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v2458 = m.ExcPending
	if v2458 != 0 {
		goto L20
	} else {
		goto L760
	}
L760:
	;
	F_addReplyBulkCString(m, l0, int32(_a1216))
	mBase = m.M
	v2461 = m.ExcPending
	if v2461 != 0 {
		goto L20
	} else {
		goto L761
	}
L761:
	;
	F_addReplyBulkCString(m, l0, int32(_a1217))
	mBase = m.M
	v2464 = m.ExcPending
	if v2464 != 0 {
		goto L20
	} else {
		goto L762
	}
L762:
	;
	goto L1
L763:
	;
	goto L1
L764:
	;
	F_addReplySubcommandSyntaxError(m, l0)
	mBase = m.M
	v2518 = m.ExcPending
	if v2518 != 0 {
		goto L20
	} else {
		goto L782
	}
L765:
	;
	if v2506-v2508 != 0 {
		goto L764
	} else {
		goto L777
	}
L766:
	;
	v2506 = F_tolower(m, v2502)
	mBase = m.M
	v2507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2503))))
	v2508 = F_tolower(m, v2507)
	mBase = m.M
	goto L765
L767:
	;
	v2476 = v2470
	v2477 = v2471
	v2478 = v2474
	goto L770
L768:
	;
	v2502 = int32(0)
	v2503 = v2471
	goto L766
L769:
	;
	v2502 = v2499 & int32(255)
	v2503 = v2498
	goto L766
L770:
	;
	v2480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2477))))
	if v2480 == int32(0) {
		v2498 = v2477
		v2499 = v2478
		goto L769
	} else {
		goto L772
	}
L771:
	;
	v2498 = v2492
	v2499 = int32(0)
	goto L769
L772:
	;
	v2484 = v2478 & int32(255)
	if v2484 == v2480 {
		goto L773
	} else {
		goto L774
	}
L773:
	;
	v2491 = int32(1)
	v2492 = v2477 + v2491
	v2493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2476)+1)))
	if v2493 != 0 {
		v2476 = v2476 + v2491
		v2477 = v2492
		v2478 = v2493
		goto L770
	} else {
		goto L776
	}
L774:
	;
	v2486 = F_tolower(m, v2484)
	mBase = m.M
	v2487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2477))))
	v2488 = F_tolower(m, v2487)
	mBase = m.M
	if v2486 == v2488 {
		goto L773
	} else {
		goto L775
	}
L775:
	;
	v2490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2476))))
	v2498 = v2477
	v2499 = v2490
	goto L769
L776:
	;
	goto L771
L777:
	;
	v2510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v2510 != int32(2) {
		goto L778
	} else {
		goto L779
	}
L778:
	;
	F_sentinelSetDebugConfigParameters(m, l0)
	mBase = m.M
	v2516 = m.ExcPending
	if v2516 != 0 {
		goto L20
	} else {
		goto L781
	}
L779:
	;
	F_addReplySentinelDebugInfo(m, l0)
	mBase = m.M
	v2514 = m.ExcPending
	if v2514 != 0 {
		goto L20
	} else {
		goto L780
	}
L780:
	;
	goto L1
L781:
	;
	goto L1
L782:
	;
	goto L1
L783:
	;
	goto L1
L784:
	;
	if v2528 != 0 {
		goto L785
	} else {
		goto L786
	}
L785:
	;
	v2536 = v2528
	goto L787
L786:
	;
	v2536 = int32(_a1180)
	goto L787
L787:
	;
	F_addReplyBulkCString(m, l0, v2536)
	mBase = m.M
	v2538 = m.ExcPending
	if v2538 != 0 {
		goto L20
	} else {
		goto L788
	}
L788:
	;
	F_addReplyLongLong(m, l0, v2531)
	mBase = m.M
	v2540 = m.ExcPending
	if v2540 != 0 {
		goto L20
	} else {
		goto L789
	}
L789:
	;
	if v2528 == int32(0) {
		goto L1
	} else {
		goto L790
	}
L790:
	;
	F_sdsfree(m, v2528)
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L20
	} else {
		goto L791
	}
L791:
	;
	goto L1
}
func F_sentinelDisconnectCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	if v4 == int32(0) {
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
		if v7 != l0 {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(1)
	}
	return
}
func F_sentinelEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v42 int32
	_ = v42
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
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v369 int32
	_ = v369
	v15 = m.G0
	v17 = v15 - int32(1120)
	m.G0 = v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v19 != int32(37) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v127&int32(255) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L2:
	;
	v125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+80)) = uint8(v125)
	v127 = v19
	v129 = l3
	goto L1
L3:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
	if v22 != int32(64) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v25&int32(1) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)))
	v127 = v124
	v129 = l3 + int32(2)
	goto L1
L6:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v92 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, _consts[639]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v91+base.B2i32(v93 == v92)<<(uint(int32(2))%32))))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v100
	v112 = F_snprintf(m, v17+int32(80), int32(1024), int32(_a1131), v17+int32(32))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L17
	} else {
		goto L24
	}
L7:
	;
	v32 = v25 & int32(2)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+192))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v88 = int32(_a200)
	goto L6
L9:
	;
	if v32 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v37 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, _consts[639]))
	v42 = base.B2i32(v38 == v37) << (uint(int32(2)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v36+v42)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v45+v42)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(64)))) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(68)))) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(72)))) = v50
	if v25&int32(4) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v66 = int32(_a1132)
	goto L13
L12:
	;
	v66 = int32(_a242)
	goto L13
L13:
	;
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v67 = int32(_a201)
	goto L16
L15:
	;
	v67 = v66
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v48
	v78 = F_snprintf(m, v17+int32(80), int32(1024), int32(_a1133), v17+int32(48))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	goto L5
L19:
	;
	if v25&int32(4) != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v88 = int32(_a201)
	goto L6
L21:
	;
	v87 = int32(_a1132)
	goto L23
L22:
	;
	v87 = int32(_a242)
	goto L23
L23:
	;
	v88 = v87
	goto L6
L24:
	;
	goto L5
L25:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if l0 < v207 {
		goto L44
	} else {
		goto L45
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+1116)) = l4
	v141 = v17 + int32(80)
	if v141&int32(3) == int32(0) {
		v165 = v141
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v17)+1116))
	v203 = F_vsnprintf(m, v141+v198, int32(1024)-v198, v129, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L17
	} else {
		goto L43
	}
L28:
	;
	v198 = v190 - v141
	goto L27
L29:
	;
	v169 = v165
	goto L37
L30:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v151 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v154 = v141
	goto L33
L32:
	;
	v198 = v141 - v141
	goto L27
L33:
	;
	v158 = v154 + int32(1)
	if v158&int32(3) == int32(0) {
		v165 = v158
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v163 != 0 {
		v154 = v158
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v190 = v158
	goto L28
L37:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v178 = int32(-2139062144)
	if (int32(16843008)-v175|v175)&v178 == v178 {
		v169 = v169 + int32(4)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v184 = v169
	goto L40
L39:
	;
	goto L38
L40:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v188 != 0 {
		v184 = v184 + int32(1)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v190 = v184
	goto L28
L42:
	;
	goto L41
L43:
	;
	goto L25
L44:
	;
	if l0 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	if l0&int32(255) < v207 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v17 + int32(80)
	F__serverLog(m, l0, int32(_a423), v17+int32(16))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L17
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	m.G0 = v17 + int32(1120)
	return
L49:
	;
	if l1&int32(3) == int32(0) {
		v244 = l1
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v278 = F_createStringObject_1(m, l1, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L17
	} else {
		goto L66
	}
L51:
	;
	v277 = v269 - l1
	goto L50
L52:
	;
	v248 = v244
	goto L60
L53:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v230 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v233 = l1
	goto L56
L55:
	;
	v277 = l1 - l1
	goto L50
L56:
	;
	v237 = v233 + int32(1)
	if v237&int32(3) == int32(0) {
		v244 = v237
		goto L52
	} else {
		goto L58
	}
L58:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	if v242 != 0 {
		v233 = v237
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v269 = v237
	goto L51
L60:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	v257 = int32(-2139062144)
	if (int32(16843008)-v254|v254)&v257 == v257 {
		v248 = v248 + int32(4)
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v263 = v248
	goto L63
L62:
	;
	goto L61
L63:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	if v267 != 0 {
		v263 = v263 + int32(1)
		goto L63
	} else {
		goto L65
	}
L64:
	;
	v269 = v263
	goto L51
L65:
	;
	goto L64
L66:
	;
	v281 = v17 + int32(80)
	if v281&int32(3) == int32(0) {
		v305 = v281
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v339 = F_createStringObject_1(m, v281, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L17
	} else {
		goto L83
	}
L68:
	;
	v338 = v330 - v281
	goto L67
L69:
	;
	v309 = v305
	goto L77
L70:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v291 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v294 = v281
	goto L73
L72:
	;
	v338 = v281 - v281
	goto L67
L73:
	;
	v298 = v294 + int32(1)
	if v298&int32(3) == int32(0) {
		v305 = v298
		goto L69
	} else {
		goto L75
	}
L75:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if v303 != 0 {
		v294 = v298
		goto L73
	} else {
		goto L76
	}
L76:
	;
	v330 = v298
	goto L68
L77:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v318 = int32(-2139062144)
	if (int32(16843008)-v315|v315)&v318 == v318 {
		v309 = v309 + int32(4)
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v324 = v309
	goto L80
L79:
	;
	goto L78
L80:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	if v328 != 0 {
		v324 = v324 + int32(1)
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v330 = v324
	goto L68
L82:
	;
	goto L81
L83:
	;
	v342 = F_pubsubPublishMessage(m, v278, v339, int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L17
	} else {
		goto L84
	}
L84:
	;
	F_decrRefCount(m, v278)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L17
	} else {
		goto L85
	}
L85:
	;
	F_decrRefCount(m, v339)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L17
	} else {
		goto L86
	}
L86:
	;
	if l0 != int32(3) {
		goto L48
	} else {
		goto L87
	}
L87:
	;
	if l2 == int32(0) {
		goto L48
	} else {
		goto L88
	}
L88:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v352&int32(1) != 0 {
		v358 = l2
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+284))
	if v359 == int32(0) {
		goto L48
	} else {
		goto L92
	}
L90:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l2)+192))
	if v355 == int32(0) {
		goto L48
	} else {
		goto L91
	}
L91:
	;
	v358 = v355
	goto L89
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v17 + int32(80)
	F_sentinelScheduleScriptExecution(m, v359, v17)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L17
	} else {
		goto L93
	}
L93:
	;
	goto L48
}
func F_sentinelFailoverSendReplicaOfNoOne(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v5 == int32(0) {
		v22 = F_sentinelSendReplicaOf(m, v3, int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			if v22 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
				F_sentinelEvent(m, int32(2), int32(_a1235), v26, int32(_a1138), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(4)
					v33 = F_mstime(m)
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v33
					return
				}
			}
		}
	} else {
		v8 = F_mstime(m)
		mBase = m.M
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
		if v8-v9 <= v11 {
			return
		} else {
			F_sentinelEvent(m, int32(3), int32(_a1236), l0, int32(_a1138), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_sentinelAbortFailover(m, l0)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_sentinelFailoverStateMachine(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3&int32(1) == int32(0) {
		F__serverAssert(m, int32(_a1134), int32(_a1135), int32(5335))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if v3&int32(64) == int32(0) {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
			switch v12 + int32(-1) {
			case 0:
				F_sentinelFailoverWaitStart(m, l0)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					return
				}
			case 1:
				F_sentinelFailoverSelectReplica(m, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					return
				}
			case 2:
				if v3&int32(16384) != 0 {
					F_sentinelFailoverSendFailover(m, l0)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						return
					}
				} else {
					F_sentinelFailoverSendReplicaOfNoOne(m, l0)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						return
					}
				}
			case 3:
				v25 = F_mstime(m)
				mBase = m.M
				v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
				v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
				if v25-v26 <= v28 {
					return
				} else {
					F_sentinelEvent(m, int32(3), int32(_a1236), l0, int32(_a1138), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_sentinelAbortFailover(m, l0)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							return
						}
					}
				}
			case 4:
				F_sentinelFailoverReconfNextReplica(m, l0)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					return
				}
			default:
				return
			}
		}
	}
}
func F_sentinelFailoverTo(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
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
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(112)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[639]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1+base.B2i32(v13 == v4)<<(uint(int32(2))%32))))
	v21 = v10 + int32(80)
	v23 = int64(*(*int32)(unsafe.Add(mBase, uint32(l1)+8)))
	if v23 <= int64(-1) {
		v32 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v32)
		v41 = v10 + int32(81)
		v42 = int32(31)
		v43 = int64(0) - v23
	} else {
		v41 = v21
		v42 = int32(32)
		v43 = v23
	}
	v45 = F_ull2string(m, v41, v42, v43)
	mBase = m.M
	if v45 == int32(0) {
	} else {
	}
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v68 = F_sdsnew(m, int32(_a1155))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		return int32(0)
	} else {
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
		if v72 != 0 {
			v73 = v72
		} else {
			v73 = l0
		}
		v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+104))
		v75 = F_dictFetchValue(m, v74, v68)
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return int32(0)
		} else {
			F_sdsfree(m, v68)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				if v75 != 0 {
					v80 = v75
				} else {
					v80 = int32(_a1155)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v80
				v87 = F_valkeyAsyncCommand(m, v66, int32(1004), l0, int32(_a57), v10+int32(64))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					if v87 == int32(-1) {
						v188 = int32(-1)
						m.G0 = v10 + int32(112)
						return v188
					} else {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v92 + int32(1)
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
						v98 = F_sdsnew(m, int32(_a1145))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
							if v100 != 0 {
								v101 = v100
							} else {
								v101 = l0
							}
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+104))
							v103 = F_dictFetchValue(m, v102, v98)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								F_sdsfree(m, v98)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = l2
									if v103 != 0 {
										v109 = v103
									} else {
										v109 = int32(_a1145)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v109
									v116 = F_valkeyAsyncCommand(m, v96, int32(1004), l0, int32(_a1233), v10+int32(48))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int32(0)
									} else {
										if v116 == int32(-1) {
											v188 = int32(-1)
											m.G0 = v10 + int32(112)
											return v188
										} else {
											v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = v121 + int32(1)
											v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
											v127 = F_sdsnew(m, int32(_a1004))
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return int32(0)
											} else {
												v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
												if v129 != 0 {
													v130 = v129
												} else {
													v130 = l0
												}
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+104))
												v132 = F_dictFetchValue(m, v131, v127)
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int32(0)
												} else {
													F_sdsfree(m, v127)
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return int32(0)
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v10+int32(32)))) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v19
														if v132 != 0 {
															v141 = v132
														} else {
															v141 = int32(_a1004)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v141
														*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v10 + int32(80)
														v151 = F_valkeyAsyncCommand(m, v125, int32(1004), l0, int32(_a1234), v10+int32(16))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															if v151 == int32(-1) {
																v188 = int32(-1)
																m.G0 = v10 + int32(112)
																return v188
															} else {
																v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v155)+8)) = v156 + int32(1)
																v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
																v162 = F_sdsnew(m, int32(_a1161))
																mBase = m.M
																v163 = m.ExcPending
																if v163 != 0 {
																	return int32(0)
																} else {
																	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																	if v164 != 0 {
																		v165 = v164
																	} else {
																		v165 = l0
																	}
																	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+104))
																	v167 = F_dictFetchValue(m, v166, v162)
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return int32(0)
																	} else {
																		F_sdsfree(m, v162)
																		mBase = m.M
																		v170 = m.ExcPending
																		if v170 != 0 {
																			return int32(0)
																		} else {
																			if v167 != 0 {
																				v172 = v167
																			} else {
																				v172 = int32(_a1161)
																			}
																			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v172
																			v177 = F_valkeyAsyncCommand(m, v160, int32(1004), l0, int32(_a57), v10)
																			mBase = m.M
																			v178 = m.ExcPending
																			if v178 != 0 {
																				return int32(0)
																			} else {
																				if v177 == int32(-1) {
																					v188 = int32(-1)
																				} else {
																					v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																					v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
																					*(*int32)(unsafe.Add(mBase, uint32(v181)+8)) = v182 + int32(1)
																					v188 = int32(0)
																				}
																				m.G0 = v10 + int32(112)
																				return v188
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
func F_sentinelFlushConfigAndReply(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(_a69)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[149]))
	*(*int32)(unsafe.Add(mBase, _consts[149])) = int32(10)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[178]))
	v17 = F_rewriteConfig(m, v15, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = int32(_a69)
		*(*int32)(unsafe.Add(mBase, _consts[149])) = v10
		v22 = *(*int32)(unsafe.Add(mBase, _consts[15]))
		if v17 != int32(-1) {
			if int32(2) < v22 {
				v46 = *(*int32)(unsafe.Add(mBase, _consts[77]))
				F_addReply(m, l0, v46)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				F__serverLog(m, int32(2), int32(_a1167), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, _consts[77]))
					F_addReply(m, l0, v46)
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
		} else {
			if int32(3) < v22 {
				F_addReplyError(m, l0, int32(_a1168))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				v29 = F___strerror_l(m, v28, v28)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v29
				F__serverLog(m, int32(3), int32(_a1169), v7)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_addReplyError(m, l0, int32(_a1168))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
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
func F_sentinelGetPrimaryByNameOrReplyError(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	v5 = F_objectGetVal(m, l1)
	mBase = m.M
	v6 = F_dictFetchValue(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			return v6
		} else {
			F_addReplyError(m, l0, int32(_a1170))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v6
			}
		}
	}
}
func F_sentinelInfoCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v320 int64
	_ = v320
	var v321 int64
	_ = v321
	var v322 int64
	_ = v322
	var v323 int64
	_ = v323
	var v325 int64
	_ = v325
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v331 int64
	_ = v331
	var v333 int64
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
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
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 int64
	_ = v479
	var v481 int64
	_ = v481
	var v484 int64
	_ = v484
	var v485 int64
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v494 int64
	_ = v494
	var v504 int32
	_ = v504
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int64
	_ = v550
	var v551 int64
	_ = v551
	var v552 int64
	_ = v552
	var v553 int64
	_ = v553
	var v554 int64
	_ = v554
	var v555 int64
	_ = v555
	var v556 int64
	_ = v556
	var v558 int64
	_ = v558
	var v560 int64
	_ = v560
	var v562 int64
	_ = v562
	var v564 int64
	_ = v564
	var v566 int64
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int64
	_ = v726
	var v727 int64
	_ = v727
	var v728 int64
	_ = v728
	var v729 int64
	_ = v729
	var v730 int64
	_ = v730
	var v731 int64
	_ = v731
	var v732 int64
	_ = v732
	var v734 int64
	_ = v734
	var v736 int64
	_ = v736
	var v738 int64
	_ = v738
	var v740 int64
	_ = v740
	var v742 int64
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v804 int32
	_ = v804
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(112)
	m.G0 = v20
	v25 = *(*int64)(unsafe.Add(mBase, _consts[654]))
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(96)))) = v25
	v28 = *(*int64)(unsafe.Add(mBase, _consts[655]))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+88)) = v28
	v31 = *(*int64)(unsafe.Add(mBase, _consts[656]))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+80)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v2
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v49 = F_genInfoSectionDict(m, v37+int32(4), v40+int32(-1), v20+int32(80), v20+int32(76), v20+int32(72))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v51)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L2
	} else {
		goto L82
	}
L2:
	;
	return
L3:
	;
	v51 = F_dictGetSafeIterator(m, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v60 = v51 + int32(20)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if v61 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v156 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L6:
	;
	v67 = v60
	v68 = v64
	goto L9
L7:
	;
	v64 = int32(1)
	goto L6
L8:
	;
	v64 = int32(0)
	goto L6
L9:
	;
	switch v68 {
	case 0:
		goto L14
	default:
		goto L13
	}
L11:
	;
	v68 = int32(0)
	goto L9
L12:
	;
	goto L5
L13:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v148
	if v148 == int32(0) {
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v72 != int32(-1) {
		v111 = v72
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v112 = int32(1)
	v113 = v111 + v112
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v113
	v115 = int32(0)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+v119+int32(26)))))
	if v123 == int32(255) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v76 != 0 {
		v111 = int32(-1)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	if v78 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
	if v105 != int32(-1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v85 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v77)+16)))
	v86 = int64(*(*int8)(unsafe.Add(mBase, uint32(v77)+27)))
	v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77)+8)))
	v88 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v77)+12)))
	v89 = int64(*(*int8)(unsafe.Add(mBase, uint32(v77)+26)))
	v90 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77)+4)))
	v91 = F_wangHash64(m, v90)
	mBase = m.M
	v93 = F_wangHash64(m, v89+v91)
	mBase = m.M
	v95 = F_wangHash64(m, v88+v93)
	mBase = m.M
	v97 = F_wangHash64(m, v87+v95)
	mBase = m.M
	v99 = F_wangHash64(m, v86+v97)
	mBase = m.M
	v101 = F_wangHash64(m, v85+v99)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v51)+24)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v104 = v103
	goto L18
L20:
	;
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+24)))
	v83 = v81 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v77)+24)) = uint16(v83)
	v104 = v77
	goto L18
L21:
	;
	v111 = v105 + int32(-1)
	goto L15
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v111 = v108
	goto L15
L23:
	;
	v138 = int32(2)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v118+v136<<(uint(v138)%32)+int32(4))))
	v67 = v143 + v137<<(uint(v138)%32)
	v68 = int32(1)
	goto L9
L24:
	;
	v127 = v115
	goto L26
L25:
	;
	v127 = v112 << (uint(v123) % 32)
	goto L26
L26:
	;
	if v113 < v127 {
		v136 = v119
		v137 = v113
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if v119 != 0 {
		v156 = v115
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
	if v129 == int32(-1) {
		v156 = v115
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51)+4)) = int64(4294967296)
	v136 = int32(1)
	v137 = int32(0)
	goto L23
L30:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v152
	v156 = v148
	goto L12
L31:
	;
	v166 = v156
	goto L32
L32:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	goto L34
L33:
	;
	goto L1
L34:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	if v181 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v292 = v51 + int32(20)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if v293 != 0 {
		goto L57
	} else {
		goto L58
	}
L36:
	;
	v266 = F_dictDelete(m, v49, v180)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L2
	} else {
		goto L54
	}
L37:
	;
	v188 = v181
	v189 = int32(0)
	goto L38
L38:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v203 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L36
L40:
	;
	if v235-v237 == int32(0) {
		goto L35
	} else {
		goto L52
	}
L41:
	;
	v235 = F_tolower(m, v231)
	mBase = m.M
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	v237 = F_tolower(m, v236)
	mBase = m.M
	goto L40
L42:
	;
	v205 = v188
	v206 = v180
	v207 = v203
	goto L45
L43:
	;
	v231 = int32(0)
	v232 = v180
	goto L41
L44:
	;
	v231 = v228 & int32(255)
	v232 = v227
	goto L41
L45:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v209 == int32(0) {
		v227 = v206
		v228 = v207
		goto L44
	} else {
		goto L47
	}
L46:
	;
	v227 = v221
	v228 = int32(0)
	goto L44
L47:
	;
	v213 = v207 & int32(255)
	if v213 == v209 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v220 = int32(1)
	v221 = v206 + v220
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	if v222 != 0 {
		v205 = v205 + v220
		v206 = v221
		v207 = v222
		goto L45
	} else {
		goto L51
	}
L49:
	;
	v215 = F_tolower(m, v213)
	mBase = m.M
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	v217 = F_tolower(m, v216)
	mBase = m.M
	if v215 == v217 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v227 = v206
	v228 = v219
	goto L44
L51:
	;
	goto L46
L52:
	;
	v244 = v189 + int32(1)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(80)+v244<<(uint(int32(2))%32))))
	if v248 != 0 {
		v188 = v248
		v189 = v244
		goto L38
	} else {
		goto L53
	}
L53:
	;
	goto L39
L54:
	;
	goto L35
L55:
	;
	if v388 != 0 {
		v166 = v388
		goto L32
	} else {
		goto L81
	}
L56:
	;
	v299 = v292
	v300 = v296
	goto L59
L57:
	;
	v296 = int32(1)
	goto L56
L58:
	;
	v296 = int32(0)
	goto L56
L59:
	;
	switch v300 {
	case 0:
		goto L64
	default:
		goto L63
	}
L61:
	;
	v300 = int32(0)
	goto L59
L62:
	;
	goto L55
L63:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v380
	if v380 == int32(0) {
		goto L61
	} else {
		goto L80
	}
L64:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v304 != int32(-1) {
		v343 = v304
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v344 = int32(1)
	v345 = v343 + v344
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v345
	v347 = int32(0)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350+v351+int32(26)))))
	if v355 == int32(255) {
		goto L74
	} else {
		goto L75
	}
L66:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v308 != 0 {
		v343 = int32(-1)
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	if v310 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+20))
	if v337 != int32(-1) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v317 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v309)+16)))
	v318 = int64(*(*int8)(unsafe.Add(mBase, uint32(v309)+27)))
	v319 = int64(*(*int32)(unsafe.Add(mBase, uint32(v309)+8)))
	v320 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v309)+12)))
	v321 = int64(*(*int8)(unsafe.Add(mBase, uint32(v309)+26)))
	v322 = int64(*(*int32)(unsafe.Add(mBase, uint32(v309)+4)))
	v323 = F_wangHash64(m, v322)
	mBase = m.M
	v325 = F_wangHash64(m, v321+v323)
	mBase = m.M
	v327 = F_wangHash64(m, v320+v325)
	mBase = m.M
	v329 = F_wangHash64(m, v319+v327)
	mBase = m.M
	v331 = F_wangHash64(m, v318+v329)
	mBase = m.M
	v333 = F_wangHash64(m, v317+v331)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v51)+24)) = v333
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v336 = v335
	goto L68
L70:
	;
	v313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v309)+24)))
	v315 = v313 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v309)+24)) = uint16(v315)
	v336 = v309
	goto L68
L71:
	;
	v343 = v337 + int32(-1)
	goto L65
L72:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v343 = v340
	goto L65
L73:
	;
	v370 = int32(2)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v350+v368<<(uint(v370)%32)+int32(4))))
	v299 = v375 + v369<<(uint(v370)%32)
	v300 = int32(1)
	goto L59
L74:
	;
	v359 = v347
	goto L76
L75:
	;
	v359 = v344 << (uint(v355) % 32)
	goto L76
L76:
	;
	if v345 < v359 {
		v368 = v351
		v369 = v345
		goto L73
	} else {
		goto L77
	}
L77:
	;
	if v351 != 0 {
		v388 = v347
		goto L62
	} else {
		goto L78
	}
L78:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v350)+20))
	if v361 == int32(-1) {
		v388 = v347
		goto L62
	} else {
		goto L79
	}
L79:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51)+4)) = int64(4294967296)
	v368 = int32(1)
	v369 = int32(0)
	goto L73
L80:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v380)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v292))) = v384
	v388 = v380
	goto L62
L81:
	;
	goto L33
L82:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	if v411|v412 == int32(0) {
		v431 = v49
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v433 = int32(0)
	v435 = F_genValkeyInfoString(m, v431, v433, v433)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L2
	} else {
		goto L89
	}
L84:
	;
	F_releaseInfoSectionDict(m, v49)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _consts[657]))
	if v419 != 0 {
		v431 = v419
		goto L83
	} else {
		goto L86
	}
L86:
	;
	v422 = F_dictCreate(m, int32(_a1226))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, _consts[657])) = v422
	F_addInfoSectionsToDict(m, v422, v20+int32(80))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L2
	} else {
		goto L88
	}
L88:
	;
	v430 = *(*int32)(unsafe.Add(mBase, _consts[657]))
	v431 = v430
	goto L83
L89:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	if v437 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v838 = *(*int32)(unsafe.Add(mBase, _consts[657]))
	if v431 == v838 {
		goto L175
	} else {
		goto L176
	}
L91:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435+int32(-1)))))
	switch v445 & int32(7) {
	case 0:
		goto L101
	case 1:
		goto L100
	case 2:
		goto L99
	case 3:
		goto L98
	case 4:
		goto L97
	default:
		v468 = v435
		goto L95
	}
L92:
	;
	v439 = F_dictFind(m, v431, int32(_a1132))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	if v439 == int32(0) {
		v823 = v435
		goto L90
	} else {
		goto L94
	}
L94:
	;
	goto L91
L95:
	;
	v470 = int32(0)
	v472 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+16))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v472)+12))
	v477 = *(*int32)(unsafe.Add(mBase, _consts[645]))
	if v477 != 0 {
		goto L105
	} else {
		goto L106
	}
L96:
	;
	if v462 == int32(0) {
		v468 = v435
		goto L95
	} else {
		goto L102
	}
L97:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v435+int32(-17))))
	v462 = v461
	goto L96
L98:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v435+int32(-9))))
	v462 = v458
	goto L96
L99:
	;
	v455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v435+int32(-5)))))
	v462 = v455
	goto L96
L100:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435+int32(-3)))))
	v462 = v452
	goto L96
L101:
	;
	v462 = int32(base.Ui32(v445) >> (uint(int32(3)) % 32))
	goto L96
L102:
	;
	v466 = F_sdscat(m, v435, int32(_a727))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L2
	} else {
		goto L103
	}
L103:
	;
	v468 = v466
	goto L95
L104:
	;
	v486 = int32(0)
	v487 = *(*int32)(unsafe.Add(mBase, _consts[658]))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)+20))
	v494 = *(*int64)(unsafe.Add(mBase, _consts[659]))
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(48)))) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(56)))) = v488
	v504 = *(*int32)(unsafe.Add(mBase, _consts[651]))
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(60)))) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v473 + v474
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v477
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v485
	v512 = F_sdscatprintf(m, v468, int32(_a1227), v20+int32(32))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L2
	} else {
		goto L107
	}
L105:
	;
	v479 = F_mstime(m)
	mBase = m.M
	v481 = *(*int64)(unsafe.Add(mBase, _consts[660]))
	v484 = base.I64_div_s(v479-v481, int64(1000))
	v485 = v484
	goto L104
L106:
	;
	v485 = int64(-1)
	goto L104
L107:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	v516 = F_dictGetIterator(m, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L2
	} else {
		goto L109
	}
L108:
	;
	F_dictReleaseIterator(m, v516)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L2
	} else {
		goto L174
	}
L109:
	;
	v525 = v516 + int32(20)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v516)+16))
	if v526 != 0 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	if v621 == int32(0) {
		v804 = v512
		goto L108
	} else {
		goto L136
	}
L111:
	;
	v532 = v525
	v533 = v529
	goto L114
L112:
	;
	v529 = int32(1)
	goto L111
L113:
	;
	v529 = int32(0)
	goto L111
L114:
	;
	switch v533 {
	case 0:
		goto L119
	default:
		goto L118
	}
L116:
	;
	v533 = int32(0)
	goto L114
L117:
	;
	goto L110
L118:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	*(*int32)(unsafe.Add(mBase, uint32(v516)+16)) = v613
	if v613 == int32(0) {
		goto L116
	} else {
		goto L135
	}
L119:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v537 != int32(-1) {
		v576 = v537
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v577 = int32(1)
	v578 = v576 + v577
	*(*int32)(unsafe.Add(mBase, uint32(v516)+4)) = v578
	v580 = int32(0)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v516)+8))
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583+v584+int32(26)))))
	if v588 == int32(255) {
		goto L129
	} else {
		goto L130
	}
L121:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v516)+8))
	if v541 != 0 {
		v576 = int32(-1)
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	if v543 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)+20))
	if v570 != int32(-1) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v550 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v542)+16)))
	v551 = int64(*(*int8)(unsafe.Add(mBase, uint32(v542)+27)))
	v552 = int64(*(*int32)(unsafe.Add(mBase, uint32(v542)+8)))
	v553 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v542)+12)))
	v554 = int64(*(*int8)(unsafe.Add(mBase, uint32(v542)+26)))
	v555 = int64(*(*int32)(unsafe.Add(mBase, uint32(v542)+4)))
	v556 = F_wangHash64(m, v555)
	mBase = m.M
	v558 = F_wangHash64(m, v554+v556)
	mBase = m.M
	v560 = F_wangHash64(m, v553+v558)
	mBase = m.M
	v562 = F_wangHash64(m, v552+v560)
	mBase = m.M
	v564 = F_wangHash64(m, v551+v562)
	mBase = m.M
	v566 = F_wangHash64(m, v550+v564)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v516)+24)) = v566
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	v569 = v568
	goto L123
L125:
	;
	v546 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v542)+24)))
	v548 = v546 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v542)+24)) = uint16(v548)
	v569 = v542
	goto L123
L126:
	;
	v576 = v570 + int32(-1)
	goto L120
L127:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v576 = v573
	goto L120
L128:
	;
	v603 = int32(2)
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v583+v601<<(uint(v603)%32)+int32(4))))
	v532 = v608 + v602<<(uint(v603)%32)
	v533 = int32(1)
	goto L114
L129:
	;
	v592 = v580
	goto L131
L130:
	;
	v592 = v577 << (uint(v588) % 32)
	goto L131
L131:
	;
	if v578 < v592 {
		v601 = v584
		v602 = v578
		goto L128
	} else {
		goto L132
	}
L132:
	;
	if v584 != 0 {
		v621 = v580
		goto L117
	} else {
		goto L133
	}
L133:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v583)+20))
	if v594 == int32(-1) {
		v621 = v580
		goto L117
	} else {
		goto L134
	}
L134:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v516)+4)) = int64(4294967296)
	v601 = int32(1)
	v602 = int32(0)
	goto L128
L135:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v613)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v525))) = v617
	v621 = v613
	goto L117
L136:
	;
	v634 = v512
	v635 = v621
	v636 = v470
	goto L137
L137:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v635)+8))
	goto L139
L138:
	;
	v804 = v692
	goto L108
L139:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v648)+24))
	v650 = int32(0)
	v651 = *(*int32)(unsafe.Add(mBase, _consts[639]))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v649+base.B2i32(v651 == v650)<<(uint(int32(2))%32))))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v648)))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v648)+4))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v648)+144))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v660)+16))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v660)+12))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v648)+148))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v663)+16))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v663)+12))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v649)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(16)))) = v668
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(20)))) = v664 + v665
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(24)))) = v662 + v661 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v659
	if v658&int32(8) != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v682 = int32(_a1228)
	goto L142
L141:
	;
	v682 = int32(_a294)
	goto L142
L142:
	;
	if v658&int32(16) != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v685 = int32(_a1229)
	goto L145
L144:
	;
	v685 = v682
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v685
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v657
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v636
	v692 = F_sdscatprintf(m, v634, int32(_a1230), v20)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	v701 = v516 + int32(20)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v516)+16))
	if v702 != 0 {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	if v797 != 0 {
		v634 = v692
		v635 = v797
		v636 = v636 + int32(1)
		goto L137
	} else {
		goto L173
	}
L148:
	;
	v708 = v701
	v709 = v705
	goto L151
L149:
	;
	v705 = int32(1)
	goto L148
L150:
	;
	v705 = int32(0)
	goto L148
L151:
	;
	switch v709 {
	case 0:
		goto L156
	default:
		goto L155
	}
L153:
	;
	v709 = int32(0)
	goto L151
L154:
	;
	goto L147
L155:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v708)))
	*(*int32)(unsafe.Add(mBase, uint32(v516)+16)) = v789
	if v789 == int32(0) {
		goto L153
	} else {
		goto L172
	}
L156:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v713 != int32(-1) {
		v752 = v713
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v753 = int32(1)
	v754 = v752 + v753
	*(*int32)(unsafe.Add(mBase, uint32(v516)+4)) = v754
	v756 = int32(0)
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v516)+8))
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759+v760+int32(26)))))
	if v764 == int32(255) {
		goto L166
	} else {
		goto L167
	}
L158:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v516)+8))
	if v717 != 0 {
		v752 = int32(-1)
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	if v719 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v745)+20))
	if v746 != int32(-1) {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	v726 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v718)+16)))
	v727 = int64(*(*int8)(unsafe.Add(mBase, uint32(v718)+27)))
	v728 = int64(*(*int32)(unsafe.Add(mBase, uint32(v718)+8)))
	v729 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v718)+12)))
	v730 = int64(*(*int8)(unsafe.Add(mBase, uint32(v718)+26)))
	v731 = int64(*(*int32)(unsafe.Add(mBase, uint32(v718)+4)))
	v732 = F_wangHash64(m, v731)
	mBase = m.M
	v734 = F_wangHash64(m, v730+v732)
	mBase = m.M
	v736 = F_wangHash64(m, v729+v734)
	mBase = m.M
	v738 = F_wangHash64(m, v728+v736)
	mBase = m.M
	v740 = F_wangHash64(m, v727+v738)
	mBase = m.M
	v742 = F_wangHash64(m, v726+v740)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v516)+24)) = v742
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	v745 = v744
	goto L160
L162:
	;
	v722 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v718)+24)))
	v724 = v722 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v718)+24)) = uint16(v724)
	v745 = v718
	goto L160
L163:
	;
	v752 = v746 + int32(-1)
	goto L157
L164:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v752 = v749
	goto L157
L165:
	;
	v779 = int32(2)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v759+v777<<(uint(v779)%32)+int32(4))))
	v708 = v784 + v778<<(uint(v779)%32)
	v709 = int32(1)
	goto L151
L166:
	;
	v768 = v756
	goto L168
L167:
	;
	v768 = v753 << (uint(v764) % 32)
	goto L168
L168:
	;
	if v754 < v768 {
		v777 = v760
		v778 = v754
		goto L165
	} else {
		goto L169
	}
L169:
	;
	if v760 != 0 {
		v797 = v756
		goto L154
	} else {
		goto L170
	}
L170:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v759)+20))
	if v770 == int32(-1) {
		v797 = v756
		goto L154
	} else {
		goto L171
	}
L171:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v516)+4)) = int64(4294967296)
	v777 = int32(1)
	v778 = int32(0)
	goto L165
L172:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v789)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v701))) = v793
	v797 = v789
	goto L154
L173:
	;
	goto L138
L174:
	;
	v823 = v804
	goto L90
L175:
	;
	F_addReplyBulkSds(m, l0, v823)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L2
	} else {
		goto L178
	}
L176:
	;
	F_releaseInfoSectionDict(m, v431)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L2
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	m.G0 = v20 + int32(112)
	return
}
func F_sentinelInfoReplyCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	if l1 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
		if v6 == int32(0) {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			v10 = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v9 + v10
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			switch v13 + v10 {
			case 0, 13:
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				F_sentinelRefreshInstanceInfo(m, l2, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					return
				}
			default:
				return
			}
		}
	}
}
func F_sentinelIsQuorumReachable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
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
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v183 int64
	_ = v183
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v11 = int32(1)
	v12 = F_dictGetIterator(m, v8)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v12)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L2
	} else {
		goto L61
	}
L2:
	;
	return int32(0)
L3:
	;
	v23 = v12 + int32(20)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v119 == int32(0) {
		v251 = v11
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v30 = v23
	v31 = v27
	goto L8
L6:
	;
	v27 = int32(1)
	goto L5
L7:
	;
	v27 = int32(0)
	goto L5
L8:
	;
	switch v31 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v31 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v111
	if v111 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v35 != int32(-1) {
		v74 = v35
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v75 = int32(1)
	v76 = v74 + v75
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v76
	v78 = int32(0)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v82+int32(26)))))
	if v86 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v39 != 0 {
		v74 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v41 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	if v68 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v48 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v40)+16)))
	v49 = int64(*(*int8)(unsafe.Add(mBase, uint32(v40)+27)))
	v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v40)+8)))
	v51 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v40)+12)))
	v52 = int64(*(*int8)(unsafe.Add(mBase, uint32(v40)+26)))
	v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(v40)+4)))
	v54 = F_wangHash64(m, v53)
	mBase = m.M
	v56 = F_wangHash64(m, v52+v54)
	mBase = m.M
	v58 = F_wangHash64(m, v51+v56)
	mBase = m.M
	v60 = F_wangHash64(m, v50+v58)
	mBase = m.M
	v62 = F_wangHash64(m, v49+v60)
	mBase = m.M
	v64 = F_wangHash64(m, v48+v62)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v67 = v66
	goto L17
L19:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+24)))
	v46 = v44 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v40)+24)) = uint16(v46)
	v67 = v40
	goto L17
L20:
	;
	v74 = v68 + int32(-1)
	goto L14
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v74 = v71
	goto L14
L22:
	;
	v101 = int32(2)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v81+v99<<(uint(v101)%32)+int32(4))))
	v30 = v106 + v100<<(uint(v101)%32)
	v31 = int32(1)
	goto L8
L23:
	;
	v90 = v78
	goto L25
L24:
	;
	v90 = v75 << (uint(v86) % 32)
	goto L25
L25:
	;
	if v76 < v90 {
		v99 = v82
		v100 = v76
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v82 != 0 {
		v119 = v78
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	if v92 == int32(-1) {
		v119 = v78
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+4)) = int64(4294967296)
	v99 = int32(1)
	v100 = int32(0)
	goto L22
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v115
	v119 = v111
	goto L11
L30:
	;
	v127 = v119
	v130 = v11
	goto L31
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	goto L33
L32:
	;
	v251 = v138
	goto L1
L33:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v138 = v130 + base.B2i32(v133&int32(24) == int32(0))
	v146 = v12 + int32(20)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v147 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	if v242 != 0 {
		v127 = v242
		v130 = v138
		goto L31
	} else {
		goto L60
	}
L35:
	;
	v153 = v146
	v154 = v150
	goto L38
L36:
	;
	v150 = int32(1)
	goto L35
L37:
	;
	v150 = int32(0)
	goto L35
L38:
	;
	switch v154 {
	case 0:
		goto L43
	default:
		goto L42
	}
L40:
	;
	v154 = int32(0)
	goto L38
L41:
	;
	goto L34
L42:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v234
	if v234 == int32(0) {
		goto L40
	} else {
		goto L59
	}
L43:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v158 != int32(-1) {
		v197 = v158
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v198 = int32(1)
	v199 = v197 + v198
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v199
	v201 = int32(0)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+v205+int32(26)))))
	if v209 == int32(255) {
		goto L53
	} else {
		goto L54
	}
L45:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v162 != 0 {
		v197 = int32(-1)
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v164 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+20))
	if v191 != int32(-1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v171 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v163)+16)))
	v172 = int64(*(*int8)(unsafe.Add(mBase, uint32(v163)+27)))
	v173 = int64(*(*int32)(unsafe.Add(mBase, uint32(v163)+8)))
	v174 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v163)+12)))
	v175 = int64(*(*int8)(unsafe.Add(mBase, uint32(v163)+26)))
	v176 = int64(*(*int32)(unsafe.Add(mBase, uint32(v163)+4)))
	v177 = F_wangHash64(m, v176)
	mBase = m.M
	v179 = F_wangHash64(m, v175+v177)
	mBase = m.M
	v181 = F_wangHash64(m, v174+v179)
	mBase = m.M
	v183 = F_wangHash64(m, v173+v181)
	mBase = m.M
	v185 = F_wangHash64(m, v172+v183)
	mBase = m.M
	v187 = F_wangHash64(m, v171+v185)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v190 = v189
	goto L47
L49:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+24)))
	v169 = v167 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v163)+24)) = uint16(v169)
	v190 = v163
	goto L47
L50:
	;
	v197 = v191 + int32(-1)
	goto L44
L51:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v197 = v194
	goto L44
L52:
	;
	v224 = int32(2)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v204+v222<<(uint(v224)%32)+int32(4))))
	v153 = v229 + v223<<(uint(v224)%32)
	v154 = int32(1)
	goto L38
L53:
	;
	v213 = v201
	goto L55
L54:
	;
	v213 = v198 << (uint(v209) % 32)
	goto L55
L55:
	;
	if v199 < v213 {
		v222 = v205
		v223 = v199
		goto L52
	} else {
		goto L56
	}
L56:
	;
	if v205 != 0 {
		v242 = v201
		goto L41
	} else {
		goto L57
	}
L57:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v204)+20))
	if v215 == int32(-1) {
		v242 = v201
		goto L41
	} else {
		goto L58
	}
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+4)) = int64(4294967296)
	v222 = int32(1)
	v223 = int32(0)
	goto L52
L59:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v238
	v242 = v234
	goto L41
L60:
	;
	goto L32
L61:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if l1 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v259 = base.B2i32(v251 < v255)
	v260 = int32(2)
	v266 = base.I32_div_s(v10+v9+int32(1), v260)
	if v266 < v251 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v251
	goto L62
L64:
	;
	v268 = v259
	goto L66
L65:
	;
	v268 = v259 | v260
	goto L66
L66:
	;
	return v268
}
func F_sentinelKillClients(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v91 int32
	_ = v91
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
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v11 != 0 {
		v14 = F_sdsnew(m, int32(_a1155))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
			if v18 != 0 {
				v19 = v18
			} else {
				v19 = l0
			}
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
			v21 = F_dictFetchValue(m, v20, v14)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_sdsfree(m, v14)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v21 != 0 {
						v26 = v21
					} else {
						v26 = int32(_a1155)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v26
					v33 = F_valkeyAsyncCommand(m, v11, int32(1004), l0, int32(_a57), v8+int32(80))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						if v33 == int32(-1) {
							v187 = int32(-1)
							m.G0 = v8 + int32(96)
							return v187
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v38 + int32(1)
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
							v44 = F_sdsnew(m, int32(_a1156))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
								if v46 != 0 {
									v47 = v46
								} else {
									v47 = l0
								}
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+104))
								v49 = F_dictFetchValue(m, v48, v44)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									F_sdsfree(m, v44)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										if v49 != 0 {
											v54 = v49
										} else {
											v54 = int32(_a1156)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v54
										v61 = F_valkeyAsyncCommand(m, v42, int32(1004), l0, int32(_a1157), v8+int32(64))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											if v61 == int32(-1) {
												v187 = int32(-1)
												m.G0 = v8 + int32(96)
												return v187
											} else {
												v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v66 + int32(1)
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
												v72 = F_sdsnew(m, int32(_a1145))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
													if v74 != 0 {
														v75 = v74
													} else {
														v75 = l0
													}
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+104))
													v77 = F_dictFetchValue(m, v76, v72)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return int32(0)
													} else {
														F_sdsfree(m, v72)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = int32(_a1158)
															if v77 != 0 {
																v84 = v77
															} else {
																v84 = int32(_a1145)
															}
															*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v84
															v91 = F_valkeyAsyncCommand(m, v70, int32(1004), l0, int32(_a1159), v8+int32(48))
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int32(0)
															} else {
																if v91 == int32(-1) {
																	v187 = int32(-1)
																	m.G0 = v8 + int32(96)
																	return v187
																} else {
																	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v96 + int32(1)
																	v100 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
																	v102 = F_sdsnew(m, int32(_a1145))
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return int32(0)
																	} else {
																		v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																		if v104 != 0 {
																			v105 = v104
																		} else {
																			v105 = l0
																		}
																		v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+104))
																		v107 = F_dictFetchValue(m, v106, v102)
																		mBase = m.M
																		v108 = m.ExcPending
																		if v108 != 0 {
																			return int32(0)
																		} else {
																			F_sdsfree(m, v102)
																			mBase = m.M
																			v110 = m.ExcPending
																			if v110 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(_a1152)
																				if v107 != 0 {
																					v114 = v107
																				} else {
																					v114 = int32(_a1145)
																				}
																				*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v114
																				v121 = F_valkeyAsyncCommand(m, v100, int32(1004), l0, int32(_a1159), v8+int32(32))
																				mBase = m.M
																				v122 = m.ExcPending
																				if v122 != 0 {
																					return int32(0)
																				} else {
																					if v121 == int32(-1) {
																						v187 = int32(-1)
																						m.G0 = v8 + int32(96)
																						return v187
																					} else {
																						v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																						v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v125)+8)) = v126 + int32(1)
																						v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
																						v132 = F_sdsnew(m, int32(_a1145))
																						mBase = m.M
																						v133 = m.ExcPending
																						if v133 != 0 {
																							return int32(0)
																						} else {
																							v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																							if v134 != 0 {
																								v135 = v134
																							} else {
																								v135 = l0
																							}
																							v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+104))
																							v137 = F_dictFetchValue(m, v136, v132)
																							mBase = m.M
																							v138 = m.ExcPending
																							if v138 != 0 {
																								return int32(0)
																							} else {
																								F_sdsfree(m, v132)
																								mBase = m.M
																								v140 = m.ExcPending
																								if v140 != 0 {
																									return int32(0)
																								} else {
																									if v137 != 0 {
																										v142 = v137
																									} else {
																										v142 = int32(_a1145)
																									}
																									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v142
																									v149 = F_valkeyAsyncCommand(m, v130, int32(1004), l0, int32(_a1160), v8+int32(16))
																									mBase = m.M
																									v150 = m.ExcPending
																									if v150 != 0 {
																										return int32(0)
																									} else {
																										if v149 == int32(-1) {
																											v187 = int32(-1)
																											m.G0 = v8 + int32(96)
																											return v187
																										} else {
																											v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																											v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
																											*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = v154 + int32(1)
																											v158 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
																											v160 = F_sdsnew(m, int32(_a1161))
																											mBase = m.M
																											v161 = m.ExcPending
																											if v161 != 0 {
																												return int32(0)
																											} else {
																												v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																												if v162 != 0 {
																													v163 = v162
																												} else {
																													v163 = l0
																												}
																												v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+104))
																												v165 = F_dictFetchValue(m, v164, v160)
																												mBase = m.M
																												v166 = m.ExcPending
																												if v166 != 0 {
																													return int32(0)
																												} else {
																													F_sdsfree(m, v160)
																													mBase = m.M
																													v168 = m.ExcPending
																													if v168 != 0 {
																														return int32(0)
																													} else {
																														if v165 != 0 {
																															v170 = v165
																														} else {
																															v170 = int32(_a1161)
																														}
																														*(*int32)(unsafe.Add(mBase, uint32(v8))) = v170
																														v175 = F_valkeyAsyncCommand(m, v158, int32(1004), l0, int32(_a57), v8)
																														mBase = m.M
																														v176 = m.ExcPending
																														if v176 != 0 {
																															return int32(0)
																														} else {
																															if v175 == int32(-1) {
																																v187 = int32(-1)
																															} else {
																																v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																																v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
																																*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v180 + int32(1)
																																v187 = int32(0)
																															}
																															m.G0 = v8 + int32(96)
																															return v187
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
	} else {
		v187 = int32(-1)
		m.G0 = v8 + int32(96)
		return v187
	}
}
func F_sentinelReconnectInstance(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int64
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(48)
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v16 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = F_mstime(m)
	mBase = m.M
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)+80))
	v24 = *(*int64)(unsafe.Add(mBase, _consts[643]))
	if v19-v21 < v24 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+80)) = v19
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v179&int32(3) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v29 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, _consts[644]))
	if v30 == v29 {
		v48 = v28
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v53 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	v54 = F_valkeyAsyncConnectBind(m, v50, v51, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L16
	}
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v36 = F_createSentinelAddr(m, v33, v34, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	F_sdsfree(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v48 = v38
	goto L7
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	F_sdsfree(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	F_valkey_free(m, v38)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v36
	v48 = v36
	goto L7
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v54
	if v54 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+204))
	if v119 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L18:
	;
	v112 = int32(0)
	F_sentinelEvent(m, v112, int32(_a1147), l0, int32(_a1148), v112)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L35
	}
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+204))
	if v59 != 0 {
		v118 = v54
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+136))
	v64 = m.G0
	v66 = v64 - int32(16)
	m.G0 = v66
	goto L25
L21:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v110 != 0 {
		v118 = v110
		goto L17
	} else {
		goto L34
	}
L22:
	;
	m.G0 = v66 + int32(16)
	goto L21
L23:
	;
	goto L22
L24:
	;
	if v74&int32(1) != 0 {
		goto L22
	} else {
		goto L29
	}
L25:
	;
	v74 = F_fcntl(m, v60, int32(1), int32(0))
	mBase = m.M
	if v74 != int32(-1) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v77 = F___errno_location(m)
	mBase = m.M
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v78 == int32(27) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L23
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v74 | int32(1)
	v91 = F_fcntl(m, v60, int32(2), v66)
	mBase = m.M
	if v91 != int32(-1) {
		goto L22
	} else {
		goto L32
	}
L31:
	;
	goto L23
L32:
	;
	v94 = F___errno_location(m)
	mBase = m.M
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95 == int32(27) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L18
L35:
	;
	goto L5
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(0)
	v148 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v150)+212)) = v11
	v153 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	F_valkeyAeAttach(m, v153, v150)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L43
	}
L37:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v122
	F_sentinelEvent(m, int32(0), int32(_a1147), l0, int32(_a1149), v9+int32(32))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v131 == int32(0) {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v136 != v131 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+212)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(1)
	F_valkeyAsyncFree(m, v131)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L9
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(0)
	goto L40
L42:
	;
	goto L5
L43:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v158 = F_valkeyAsyncSetConnectCallback(m, v156, int32(1006))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160)+248))
	if v164 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_sentinelSendAuthIfNeeded(m, l0, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L9
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+248)) = int32(1007)
	goto L46
L48:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_sentinelSetClientName(m, l0, v171, int32(_a1150))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	v175 = F_sentinelSendPing(m, l0)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	goto L5
L51:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v315 == int32(0) {
		goto L1
	} else {
		goto L89
	}
L52:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v184 != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
	v189 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	v190 = F_valkeyAsyncConnectBind(m, v186, v187, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v190
	if v190 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+204))
	if v255 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L56:
	;
	v248 = int32(0)
	F_sentinelEvent(m, v248, int32(_a1151), l0, int32(_a1148), v248)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L9
	} else {
		goto L73
	}
L57:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v190)+204))
	if v195 != 0 {
		v254 = v190
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v190)+136))
	v200 = m.G0
	v202 = v200 - int32(16)
	m.G0 = v202
	goto L63
L59:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v246 != 0 {
		v254 = v246
		goto L55
	} else {
		goto L72
	}
L60:
	;
	m.G0 = v202 + int32(16)
	goto L59
L61:
	;
	goto L60
L62:
	;
	if v210&int32(1) != 0 {
		goto L60
	} else {
		goto L67
	}
L63:
	;
	v210 = F_fcntl(m, v196, int32(1), int32(0))
	mBase = m.M
	if v210 != int32(-1) {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v213 = F___errno_location(m)
	mBase = m.M
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	if v214 == int32(27) {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L61
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v210 | int32(1)
	v227 = F_fcntl(m, v196, int32(2), v202)
	mBase = m.M
	if v227 != int32(-1) {
		goto L60
	} else {
		goto L70
	}
L69:
	;
	goto L61
L70:
	;
	v230 = F___errno_location(m)
	mBase = m.M
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v231 == int32(27) {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	goto L56
L73:
	;
	goto L51
L74:
	;
	v270 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v270
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v272)+212)) = v11
	v275 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	F_valkeyAeAttach(m, v275, v272)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L9
	} else {
		goto L78
	}
L75:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v254)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v258
	F_sentinelEvent(m, int32(0), int32(_a1151), l0, int32(_a1149), v9+int32(16))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	F_instanceLinkCloseConnection(m, v11, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	goto L51
L78:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v280 = F_valkeyAsyncSetConnectCallback(m, v278, int32(1006))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v282)+248))
	if v286 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	F_sentinelSendAuthIfNeeded(m, l0, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L9
	} else {
		goto L83
	}
L81:
	;
	goto L80
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+248)) = int32(1007)
	goto L81
L83:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	F_sentinelSetClientName(m, l0, v293, int32(_a1152))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L9
	} else {
		goto L84
	}
L84:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v299 = F_sentinelInstanceMapCommand(m, l0, int32(_a1153))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(_a1154)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v299
	v306 = F_valkeyAsyncCommand(m, v297, int32(1008), l0, int32(_a423), v9)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L9
	} else {
		goto L86
	}
L86:
	;
	if v306 == int32(0) {
		goto L51
	} else {
		goto L87
	}
L87:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	F_instanceLinkCloseConnection(m, v11, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	goto L1
L89:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v318&int32(4) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
	goto L1
L91:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v321 == int32(0) {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	goto L90
}
func F_sentinelResetPrimariesByPattern(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
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
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v193 int64
	_ = v193
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v199 int64
	_ = v199
	var v201 int64
	_ = v201
	var v203 int64
	_ = v203
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	v10 = F_dictGetIterator(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v10)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L2
	} else {
		goto L66
	}
L2:
	;
	return int32(0)
L3:
	;
	v21 = v10 + int32(20)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v117 == int32(0) {
		v268 = v3
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v28 = v21
	v29 = v25
	goto L8
L6:
	;
	v25 = int32(1)
	goto L5
L7:
	;
	v25 = int32(0)
	goto L5
L8:
	;
	switch v29 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v29 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v109
	if v109 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v33 != int32(-1) {
		v72 = v33
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v73 = int32(1)
	v74 = v72 + v73
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v74
	v76 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+v80+int32(26)))))
	if v84 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v37 != 0 {
		v72 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v39 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	if v66 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v46 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v38)+16)))
	v47 = int64(*(*int8)(unsafe.Add(mBase, uint32(v38)+27)))
	v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v38)+8)))
	v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v38)+12)))
	v50 = int64(*(*int8)(unsafe.Add(mBase, uint32(v38)+26)))
	v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v38)+4)))
	v52 = F_wangHash64(m, v51)
	mBase = m.M
	v54 = F_wangHash64(m, v50+v52)
	mBase = m.M
	v56 = F_wangHash64(m, v49+v54)
	mBase = m.M
	v58 = F_wangHash64(m, v48+v56)
	mBase = m.M
	v60 = F_wangHash64(m, v47+v58)
	mBase = m.M
	v62 = F_wangHash64(m, v46+v60)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v65 = v64
	goto L17
L19:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+24)))
	v44 = v42 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+24)) = uint16(v44)
	v65 = v38
	goto L17
L20:
	;
	v72 = v66 + int32(-1)
	goto L14
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v72 = v69
	goto L14
L22:
	;
	v99 = int32(2)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v79+v97<<(uint(v99)%32)+int32(4))))
	v28 = v104 + v98<<(uint(v99)%32)
	v29 = int32(1)
	goto L8
L23:
	;
	v88 = v76
	goto L25
L24:
	;
	v88 = v73 << (uint(v84) % 32)
	goto L25
L25:
	;
	if v74 < v88 {
		v97 = v80
		v98 = v74
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v80 != 0 {
		v117 = v76
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	if v90 == int32(-1) {
		v117 = v76
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(4294967296)
	v97 = int32(1)
	v98 = int32(0)
	goto L22
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v113
	v117 = v109
	goto L11
L30:
	;
	v125 = v3
	v127 = v117
	goto L31
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	goto L34
L32:
	;
	v268 = v158
	goto L1
L33:
	;
	v166 = v10 + int32(20)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v167 != 0 {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v130 == int32(0) {
		v158 = v125
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v133 = int32(0)
	v137 = m.G0
	v138 = int32(16)
	v139 = v137 - v138
	m.G0 = v139
	v141 = F_strlen(m, l0)
	mBase = m.M
	v142 = F_strlen(m, v130)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v139)+12)) = v133
	v148 = F_stringmatchlen_impl(m, l0, v141, v130, v142, v133, v139+int32(12), v133)
	mBase = m.M
	m.G0 = v139 + v138
	goto L36
L36:
	;
	if v148 == int32(0) {
		v158 = v125
		goto L33
	} else {
		goto L37
	}
L37:
	;
	F_sentinelResetPrimary(m, v129, l1)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v158 = v125 + int32(1)
	goto L33
L39:
	;
	if v262 != 0 {
		v125 = v158
		v127 = v262
		goto L31
	} else {
		goto L65
	}
L40:
	;
	v173 = v166
	v174 = v170
	goto L43
L41:
	;
	v170 = int32(1)
	goto L40
L42:
	;
	v170 = int32(0)
	goto L40
L43:
	;
	switch v174 {
	case 0:
		goto L48
	default:
		goto L47
	}
L45:
	;
	v174 = int32(0)
	goto L43
L46:
	;
	goto L39
L47:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v254
	if v254 == int32(0) {
		goto L45
	} else {
		goto L64
	}
L48:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v178 != int32(-1) {
		v217 = v178
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v218 = int32(1)
	v219 = v217 + v218
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v219
	v221 = int32(0)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224+v225+int32(26)))))
	if v229 == int32(255) {
		goto L58
	} else {
		goto L59
	}
L50:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v182 != 0 {
		v217 = int32(-1)
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v184 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+20))
	if v211 != int32(-1) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v191 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v183)+16)))
	v192 = int64(*(*int8)(unsafe.Add(mBase, uint32(v183)+27)))
	v193 = int64(*(*int32)(unsafe.Add(mBase, uint32(v183)+8)))
	v194 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v183)+12)))
	v195 = int64(*(*int8)(unsafe.Add(mBase, uint32(v183)+26)))
	v196 = int64(*(*int32)(unsafe.Add(mBase, uint32(v183)+4)))
	v197 = F_wangHash64(m, v196)
	mBase = m.M
	v199 = F_wangHash64(m, v195+v197)
	mBase = m.M
	v201 = F_wangHash64(m, v194+v199)
	mBase = m.M
	v203 = F_wangHash64(m, v193+v201)
	mBase = m.M
	v205 = F_wangHash64(m, v192+v203)
	mBase = m.M
	v207 = F_wangHash64(m, v191+v205)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v210 = v209
	goto L52
L54:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+24)))
	v189 = v187 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v183)+24)) = uint16(v189)
	v210 = v183
	goto L52
L55:
	;
	v217 = v211 + int32(-1)
	goto L49
L56:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v217 = v214
	goto L49
L57:
	;
	v244 = int32(2)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v224+v242<<(uint(v244)%32)+int32(4))))
	v173 = v249 + v243<<(uint(v244)%32)
	v174 = int32(1)
	goto L43
L58:
	;
	v233 = v221
	goto L60
L59:
	;
	v233 = v218 << (uint(v229) % 32)
	goto L60
L60:
	;
	if v219 < v233 {
		v242 = v225
		v243 = v219
		goto L57
	} else {
		goto L61
	}
L61:
	;
	if v225 != 0 {
		v262 = v221
		goto L46
	} else {
		goto L62
	}
L62:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v224)+20))
	if v235 == int32(-1) {
		v262 = v221
		goto L46
	} else {
		goto L63
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(4294967296)
	v242 = int32(1)
	v243 = int32(0)
	goto L57
L64:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v254)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v258
	v262 = v254
	goto L46
L65:
	;
	goto L32
L66:
	;
	return v268
}
func F_sentinelResetPrimary(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int64
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v6&int32(1) == int32(0) {
		F__serverAssert(m, int32(_a1134), int32(_a1135), int32(1529))
		mBase = m.M
		v120 = m.ExcPending
		if v120 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
		F_dictRelease(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = F_dictCreate(m, int32(_a1136))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v15
				if l1&int32(1) != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
					if v28 == int32(0) {
						v44 = v27
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
						if v45 == int32(0) {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v60 & int32(1)
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							if v64 == int32(0) {
								v71 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
								*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
								v75 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								F_sdsfree(m, v81)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
									F_sdsfree(m, v84)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										v87 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
										v91 = F_mstime(m)
										mBase = m.M
										v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
										v96 = F_mstime(m)
										mBase = m.M
										v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
										v99 = F_mstime(m)
										mBase = m.M
										v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
										v102 = F_mstime(m)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
										if l1&int32(65536) == v87 {
											return
										} else {
											F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								F_sdsfree(m, v64)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = int32(0)
									v71 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
									*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
									v75 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
									v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									F_sdsfree(m, v81)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
										F_sdsfree(m, v84)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											v87 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
											v91 = F_mstime(m)
											mBase = m.M
											v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
											v96 = F_mstime(m)
											mBase = m.M
											v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
											v99 = F_mstime(m)
											mBase = m.M
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
											v102 = F_mstime(m)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
											if l1&int32(65536) == v87 {
												return
											} else {
												F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								}
							}
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
							if v48 != v45 {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v44)+8)) = int64(0)
							}
							v52 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v45)+212)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = int32(1)
							F_valkeyAsyncFree(m, v45)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v60 & int32(1)
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								if v64 == int32(0) {
									v71 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
									*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
									v75 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
									v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									F_sdsfree(m, v81)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
										F_sdsfree(m, v84)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											v87 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
											v91 = F_mstime(m)
											mBase = m.M
											v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
											v96 = F_mstime(m)
											mBase = m.M
											v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
											v99 = F_mstime(m)
											mBase = m.M
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
											v102 = F_mstime(m)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
											if l1&int32(65536) == v87 {
												return
											} else {
												F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								} else {
									F_sdsfree(m, v64)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = int32(0)
										v71 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
										*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
										v75 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
										*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_sdsfree(m, v81)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
											F_sdsfree(m, v84)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												v87 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
												v91 = F_mstime(m)
												mBase = m.M
												v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
												v96 = F_mstime(m)
												mBase = m.M
												v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
												v99 = F_mstime(m)
												mBase = m.M
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
												v102 = F_mstime(m)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
												if l1&int32(65536) == v87 {
													return
												} else {
													F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
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
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = int64(0)
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
						if v33 != v28 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = int32(0)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v28)+212)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(1)
						F_valkeyAsyncFree(m, v28)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v44 = v43
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
							if v45 == int32(0) {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v60 & int32(1)
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								if v64 == int32(0) {
									v71 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
									*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
									v75 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
									v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									F_sdsfree(m, v81)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
										F_sdsfree(m, v84)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											v87 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
											v91 = F_mstime(m)
											mBase = m.M
											v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
											v96 = F_mstime(m)
											mBase = m.M
											v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
											v99 = F_mstime(m)
											mBase = m.M
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
											v102 = F_mstime(m)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
											if l1&int32(65536) == v87 {
												return
											} else {
												F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								} else {
									F_sdsfree(m, v64)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = int32(0)
										v71 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
										*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
										v75 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
										*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_sdsfree(m, v81)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
											F_sdsfree(m, v84)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												v87 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
												v91 = F_mstime(m)
												mBase = m.M
												v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
												v96 = F_mstime(m)
												mBase = m.M
												v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
												v99 = F_mstime(m)
												mBase = m.M
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
												v102 = F_mstime(m)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
												if l1&int32(65536) == v87 {
													return
												} else {
													F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														return
													}
												}
											}
										}
									}
								}
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
								if v48 != v45 {
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v44)+8)) = int64(0)
								}
								v52 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v45)+212)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = int32(1)
								F_valkeyAsyncFree(m, v45)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v60 & int32(1)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
									if v64 == int32(0) {
										v71 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
										*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
										v75 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
										*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_sdsfree(m, v81)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
											F_sdsfree(m, v84)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												v87 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
												v91 = F_mstime(m)
												mBase = m.M
												v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
												v96 = F_mstime(m)
												mBase = m.M
												v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
												v99 = F_mstime(m)
												mBase = m.M
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
												v102 = F_mstime(m)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
												if l1&int32(65536) == v87 {
													return
												} else {
													F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														return
													}
												}
											}
										}
									} else {
										F_sdsfree(m, v64)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = int32(0)
											v71 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
											*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
											v75 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											F_sdsfree(m, v81)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
												F_sdsfree(m, v84)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													v87 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
													v91 = F_mstime(m)
													mBase = m.M
													v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
													v96 = F_mstime(m)
													mBase = m.M
													v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
													v99 = F_mstime(m)
													mBase = m.M
													v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
													v102 = F_mstime(m)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
													if l1&int32(65536) == v87 {
														return
													} else {
														F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
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
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
					F_dictRelease(m, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v24 = F_dictCreate(m, int32(_a1136))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v24
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
							if v28 == int32(0) {
								v44 = v27
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
								if v45 == int32(0) {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v60 & int32(1)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
									if v64 == int32(0) {
										v71 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
										*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
										v75 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
										*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_sdsfree(m, v81)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
											F_sdsfree(m, v84)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												v87 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
												v91 = F_mstime(m)
												mBase = m.M
												v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
												v96 = F_mstime(m)
												mBase = m.M
												v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
												v99 = F_mstime(m)
												mBase = m.M
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
												v102 = F_mstime(m)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
												if l1&int32(65536) == v87 {
													return
												} else {
													F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														return
													}
												}
											}
										}
									} else {
										F_sdsfree(m, v64)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = int32(0)
											v71 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
											*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
											v75 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											F_sdsfree(m, v81)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
												F_sdsfree(m, v84)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													v87 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
													v91 = F_mstime(m)
													mBase = m.M
													v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
													v96 = F_mstime(m)
													mBase = m.M
													v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
													v99 = F_mstime(m)
													mBase = m.M
													v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
													v102 = F_mstime(m)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
													if l1&int32(65536) == v87 {
														return
													} else {
														F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return
														} else {
															return
														}
													}
												}
											}
										}
									}
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
									if v48 != v45 {
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v44)+8)) = int64(0)
									}
									v52 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v45)+212)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = int32(1)
									F_valkeyAsyncFree(m, v45)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v60 & int32(1)
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
										if v64 == int32(0) {
											v71 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
											*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
											v75 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											F_sdsfree(m, v81)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
												F_sdsfree(m, v84)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													v87 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
													v91 = F_mstime(m)
													mBase = m.M
													v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
													v96 = F_mstime(m)
													mBase = m.M
													v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
													v99 = F_mstime(m)
													mBase = m.M
													v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
													v102 = F_mstime(m)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
													if l1&int32(65536) == v87 {
														return
													} else {
														F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return
														} else {
															return
														}
													}
												}
											}
										} else {
											F_sdsfree(m, v64)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = int32(0)
												v71 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
												*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
												v75 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
												v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												F_sdsfree(m, v81)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
													F_sdsfree(m, v84)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														v87 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
														v91 = F_mstime(m)
														mBase = m.M
														v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
														v96 = F_mstime(m)
														mBase = m.M
														v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
														v99 = F_mstime(m)
														mBase = m.M
														v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
														v102 = F_mstime(m)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
														*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
														if l1&int32(65536) == v87 {
															return
														} else {
															F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
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
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = int64(0)
								v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
								if v33 != v28 {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = int32(0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v28)+212)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(1)
								F_valkeyAsyncFree(m, v28)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									v44 = v43
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
									if v45 == int32(0) {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v60 & int32(1)
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
										if v64 == int32(0) {
											v71 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
											*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
											v75 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											F_sdsfree(m, v81)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
												F_sdsfree(m, v84)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													v87 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
													v91 = F_mstime(m)
													mBase = m.M
													v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
													v96 = F_mstime(m)
													mBase = m.M
													v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
													v99 = F_mstime(m)
													mBase = m.M
													v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
													v102 = F_mstime(m)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
													if l1&int32(65536) == v87 {
														return
													} else {
														F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return
														} else {
															return
														}
													}
												}
											}
										} else {
											F_sdsfree(m, v64)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = int32(0)
												v71 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
												*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
												v75 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
												v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												F_sdsfree(m, v81)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
													F_sdsfree(m, v84)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														v87 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
														v91 = F_mstime(m)
														mBase = m.M
														v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
														v96 = F_mstime(m)
														mBase = m.M
														v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
														v99 = F_mstime(m)
														mBase = m.M
														v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
														v102 = F_mstime(m)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
														*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
														if l1&int32(65536) == v87 {
															return
														} else {
															F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return
															} else {
																return
															}
														}
													}
												}
											}
										}
									} else {
										v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
										if v48 != v45 {
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v44)+8)) = int64(0)
										}
										v52 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v52
										*(*int32)(unsafe.Add(mBase, uint32(v45)+212)) = v52
										*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = int32(1)
										F_valkeyAsyncFree(m, v45)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = v60 & int32(1)
											v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
											if v64 == int32(0) {
												v71 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
												*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
												v75 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
												v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												F_sdsfree(m, v81)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
													F_sdsfree(m, v84)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														v87 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
														v91 = F_mstime(m)
														mBase = m.M
														v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
														v96 = F_mstime(m)
														mBase = m.M
														v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
														v99 = F_mstime(m)
														mBase = m.M
														v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
														v102 = F_mstime(m)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
														*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
														if l1&int32(65536) == v87 {
															return
														} else {
															F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return
															} else {
																return
															}
														}
													}
												}
											} else {
												F_sdsfree(m, v64)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = int32(0)
													v71 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v71
													*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v71
													v75 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v75
													*(*int64)(unsafe.Add(mBase, uint32(l0+int32(256)))) = v75
													v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													F_sdsfree(m, v81)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return
													} else {
														v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
														F_sdsfree(m, v84)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return
														} else {
															v87 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v87
															*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
															v91 = F_mstime(m)
															mBase = m.M
															v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
															*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v92)+56)) = v91
															v96 = F_mstime(m)
															mBase = m.M
															v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
															*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = v96
															v99 = F_mstime(m)
															mBase = m.M
															v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
															*(*int64)(unsafe.Add(mBase, uint32(v100)+72)) = v99
															v102 = F_mstime(m)
															mBase = m.M
															*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
															*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v102
															if l1&int32(65536) == v87 {
																return
															} else {
																F_sentinelEvent(m, int32(3), int32(_a1137), l0, int32(_a1138), int32(0))
																mBase = m.M
																v115 = m.ExcPending
																if v115 != 0 {
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
					}
				}
			}
		}
	}
}
func F_sentinelSelectReplica(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v16 = F_valkey_malloc(m, (v11+v12)<<(uint(int32(2))%32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v21&int32(8) == int32(0) {
		v29 = int64(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v32 = F_dictGetIterator(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v26 = F_mstime(m)
	mBase = m.M
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v29 = v26 - v27
	goto L3
L5:
	;
	F_dictReleaseIterator(m, v32)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L80
	}
L6:
	;
	v41 = v32 + int32(20)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v42 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v137 == int32(0) {
		goto L5
	} else {
		goto L33
	}
L8:
	;
	v48 = v41
	v49 = v45
	goto L11
L9:
	;
	v45 = int32(1)
	goto L8
L10:
	;
	v45 = int32(0)
	goto L8
L11:
	;
	switch v49 {
	case 0:
		goto L16
	default:
		goto L15
	}
L13:
	;
	v49 = int32(0)
	goto L11
L14:
	;
	goto L7
L15:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v129
	if v129 == int32(0) {
		goto L13
	} else {
		goto L32
	}
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v53 != int32(-1) {
		v92 = v53
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v93 = int32(1)
	v94 = v92 + v93
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v94
	v96 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v100+int32(26)))))
	if v104 == int32(255) {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v57 != 0 {
		v92 = int32(-1)
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v59 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+20))
	if v86 != int32(-1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v66 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+16)))
	v67 = int64(*(*int8)(unsafe.Add(mBase, uint32(v58)+27)))
	v68 = int64(*(*int32)(unsafe.Add(mBase, uint32(v58)+8)))
	v69 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+12)))
	v70 = int64(*(*int8)(unsafe.Add(mBase, uint32(v58)+26)))
	v71 = int64(*(*int32)(unsafe.Add(mBase, uint32(v58)+4)))
	v72 = F_wangHash64(m, v71)
	mBase = m.M
	v74 = F_wangHash64(m, v70+v72)
	mBase = m.M
	v76 = F_wangHash64(m, v69+v74)
	mBase = m.M
	v78 = F_wangHash64(m, v68+v76)
	mBase = m.M
	v80 = F_wangHash64(m, v67+v78)
	mBase = m.M
	v82 = F_wangHash64(m, v66+v80)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v32)+24)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v85 = v84
	goto L20
L22:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+24)))
	v64 = v62 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v58)+24)) = uint16(v64)
	v85 = v58
	goto L20
L23:
	;
	v92 = v86 + int32(-1)
	goto L17
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v92 = v89
	goto L17
L25:
	;
	v119 = int32(2)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v99+v117<<(uint(v119)%32)+int32(4))))
	v48 = v124 + v118<<(uint(v119)%32)
	v49 = int32(1)
	goto L11
L26:
	;
	v108 = v96
	goto L28
L27:
	;
	v108 = v93 << (uint(v104) % 32)
	goto L28
L28:
	;
	if v94 < v108 {
		v117 = v100
		v118 = v94
		goto L25
	} else {
		goto L29
	}
L29:
	;
	if v100 != 0 {
		v137 = v96
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	if v110 == int32(-1) {
		v137 = v96
		goto L14
	} else {
		goto L31
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+4)) = int64(4294967296)
	v117 = int32(1)
	v118 = int32(0)
	goto L25
L32:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v133
	v137 = v129
	goto L14
L33:
	;
	v148 = v137
	v154 = int32(0)
	goto L34
L34:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	goto L37
L35:
	;
	F_dictReleaseIterator(m, v32)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L74
	}
L36:
	;
	v205 = v32 + int32(20)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v206 != 0 {
		goto L49
	} else {
		goto L50
	}
L37:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v157&int32(24) != 0 {
		v196 = v154
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156)+28))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v161 != 0 {
		v196 = v154
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v162 = F_mstime(m)
	mBase = m.M
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v156)+28))
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v163)+48))
	v167 = *(*int64)(unsafe.Add(mBase, _consts[643]))
	v169 = v167 * int64(5)
	if v169 < v162-v164 {
		v196 = v154
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v156)+176))
	if v171 == int32(0) {
		v196 = v154
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v176 = *(*int64)(unsafe.Add(mBase, _consts[653]))
	v177 = F_mstime(m)
	mBase = m.M
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v156)+96))
	if v174&int32(8) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v184 = v169
	goto L44
L43:
	;
	v184 = v176 * int64(3)
	goto L44
L44:
	;
	if v184 < v177-v178 {
		v196 = v154
		goto L36
	} else {
		goto L45
	}
L45:
	;
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v156)+168))
	if v30*int64(10)+v29 < v186 {
		v196 = v154
		goto L36
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+v154<<(uint(int32(2))%32)))) = v156
	v196 = v154 + int32(1)
	goto L36
L47:
	;
	if v301 != 0 {
		v148 = v301
		v154 = v196
		goto L34
	} else {
		goto L73
	}
L48:
	;
	v212 = v205
	v213 = v209
	goto L51
L49:
	;
	v209 = int32(1)
	goto L48
L50:
	;
	v209 = int32(0)
	goto L48
L51:
	;
	switch v213 {
	case 0:
		goto L56
	default:
		goto L55
	}
L53:
	;
	v213 = int32(0)
	goto L51
L54:
	;
	goto L47
L55:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v293
	if v293 == int32(0) {
		goto L53
	} else {
		goto L72
	}
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v217 != int32(-1) {
		v256 = v217
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v257 = int32(1)
	v258 = v256 + v257
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v258
	v260 = int32(0)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263+v264+int32(26)))))
	if v268 == int32(255) {
		goto L66
	} else {
		goto L67
	}
L58:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v221 != 0 {
		v256 = int32(-1)
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v223 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+20))
	if v250 != int32(-1) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v230 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v222)+16)))
	v231 = int64(*(*int8)(unsafe.Add(mBase, uint32(v222)+27)))
	v232 = int64(*(*int32)(unsafe.Add(mBase, uint32(v222)+8)))
	v233 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v222)+12)))
	v234 = int64(*(*int8)(unsafe.Add(mBase, uint32(v222)+26)))
	v235 = int64(*(*int32)(unsafe.Add(mBase, uint32(v222)+4)))
	v236 = F_wangHash64(m, v235)
	mBase = m.M
	v238 = F_wangHash64(m, v234+v236)
	mBase = m.M
	v240 = F_wangHash64(m, v233+v238)
	mBase = m.M
	v242 = F_wangHash64(m, v232+v240)
	mBase = m.M
	v244 = F_wangHash64(m, v231+v242)
	mBase = m.M
	v246 = F_wangHash64(m, v230+v244)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v32)+24)) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v249 = v248
	goto L60
L62:
	;
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+24)))
	v228 = v226 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v222)+24)) = uint16(v228)
	v249 = v222
	goto L60
L63:
	;
	v256 = v250 + int32(-1)
	goto L57
L64:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v256 = v253
	goto L57
L65:
	;
	v283 = int32(2)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v263+v281<<(uint(v283)%32)+int32(4))))
	v212 = v288 + v282<<(uint(v283)%32)
	v213 = int32(1)
	goto L51
L66:
	;
	v272 = v260
	goto L68
L67:
	;
	v272 = v257 << (uint(v268) % 32)
	goto L68
L68:
	;
	if v258 < v272 {
		v281 = v264
		v282 = v258
		goto L65
	} else {
		goto L69
	}
L69:
	;
	if v264 != 0 {
		v301 = v260
		goto L54
	} else {
		goto L70
	}
L70:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v263)+20))
	if v274 == int32(-1) {
		v301 = v260
		goto L54
	} else {
		goto L71
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+4)) = int64(4294967296)
	v281 = int32(1)
	v282 = int32(0)
	goto L65
L72:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v293)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v297
	v301 = v293
	goto L54
L73:
	;
	goto L35
L74:
	;
	if v196 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_qsort(m, v16, v196, int32(4), int32(1019))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	F_valkey_free(m, v16)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	return int32(0)
L78:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	F_valkey_free(m, v16)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	return v315
L80:
	;
	F_valkey_free(m, v16)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	return int32(0)
}
func F_sentinelSendReplicaOf(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
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
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	v7 = m.G0
	v9 = v7 - int32(144)
	m.G0 = v9
	if l1 != 0 {
		v14 = int32(0)
		v15 = *(*int32)(unsafe.Add(mBase, _consts[639]))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1+base.B2i32(v15 == v14)<<(uint(int32(2))%32))))
		v23 = v9 + int32(112)
		v25 = int64(*(*int32)(unsafe.Add(mBase, uint32(l1)+8)))
		if v25 <= int64(-1) {
			v34 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v34)
			v43 = v9 + int32(113)
			v44 = int32(31)
			v45 = int64(0) - v25
		} else {
			v43 = v23
			v44 = int32(32)
			v45 = v25
		}
		v47 = F_ull2string(m, v43, v44, v45)
		mBase = m.M
		if v47 == int32(0) {
		} else {
		}
		v67 = v21
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = int32(4542031)
		v67 = int32(_a1162)
	}
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v71 = F_sdsnew(m, int32(_a1155))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		return int32(0)
	} else {
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
		if v75 != 0 {
			v76 = v75
		} else {
			v76 = l0
		}
		v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+104))
		v78 = F_dictFetchValue(m, v77, v71)
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return int32(0)
		} else {
			F_sdsfree(m, v71)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				if v78 != 0 {
					v83 = v78
				} else {
					v83 = int32(_a1155)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v83
				v90 = F_valkeyAsyncCommand(m, v69, int32(1004), l0, int32(_a57), v9+int32(96))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					if v90 == int32(-1) {
						v281 = int32(-1)
						m.G0 = v9 + int32(144)
						return v281
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v95 + int32(1)
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						if v99 != int32(2) {
							v130 = v94
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
							v135 = F_sdsnew(m, int32(_a1163))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return int32(0)
							} else {
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
								if v137 != 0 {
									v138 = v137
								} else {
									v138 = l0
								}
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+104))
								v140 = F_dictFetchValue(m, v139, v135)
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return int32(0)
								} else {
									F_sdsfree(m, v135)
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v67
										if v140 != 0 {
											v146 = v140
										} else {
											v146 = int32(_a1163)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v146
										*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v9 + int32(112)
										v156 = F_valkeyAsyncCommand(m, v133, int32(1004), l0, int32(_a1164), v9+int32(64))
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return int32(0)
										} else {
											if v156 == int32(-1) {
												v281 = int32(-1)
												m.G0 = v9 + int32(144)
												return v281
											} else {
												v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v161 + int32(1)
												v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
												v167 = F_sdsnew(m, int32(_a1156))
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return int32(0)
												} else {
													v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
													if v169 != 0 {
														v170 = v169
													} else {
														v170 = l0
													}
													v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+104))
													v172 = F_dictFetchValue(m, v171, v167)
													mBase = m.M
													v173 = m.ExcPending
													if v173 != 0 {
														return int32(0)
													} else {
														F_sdsfree(m, v167)
														mBase = m.M
														v175 = m.ExcPending
														if v175 != 0 {
															return int32(0)
														} else {
															if v172 != 0 {
																v177 = v172
															} else {
																v177 = int32(_a1156)
															}
															*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v177
															v184 = F_valkeyAsyncCommand(m, v165, int32(1004), l0, int32(_a1157), v9+int32(48))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return int32(0)
															} else {
																if v184 == int32(-1) {
																	v281 = int32(-1)
																	m.G0 = v9 + int32(144)
																	return v281
																} else {
																	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v188)+8)) = v189 + int32(1)
																	v193 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
																	v195 = F_sdsnew(m, int32(_a1145))
																	mBase = m.M
																	v196 = m.ExcPending
																	if v196 != 0 {
																		return int32(0)
																	} else {
																		v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																		if v197 != 0 {
																			v198 = v197
																		} else {
																			v198 = l0
																		}
																		v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+104))
																		v200 = F_dictFetchValue(m, v199, v195)
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return int32(0)
																		} else {
																			F_sdsfree(m, v195)
																			mBase = m.M
																			v203 = m.ExcPending
																			if v203 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(_a1158)
																				if v200 != 0 {
																					v207 = v200
																				} else {
																					v207 = int32(_a1145)
																				}
																				*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v207
																				v214 = F_valkeyAsyncCommand(m, v193, int32(1004), l0, int32(_a1159), v9+int32(32))
																				mBase = m.M
																				v215 = m.ExcPending
																				if v215 != 0 {
																					return int32(0)
																				} else {
																					if v214 == int32(-1) {
																						v281 = int32(-1)
																						m.G0 = v9 + int32(144)
																						return v281
																					} else {
																						v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																						v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v218)+8)) = v219 + int32(1)
																						v223 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
																						v225 = F_sdsnew(m, int32(_a1145))
																						mBase = m.M
																						v226 = m.ExcPending
																						if v226 != 0 {
																							return int32(0)
																						} else {
																							v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																							if v227 != 0 {
																								v228 = v227
																							} else {
																								v228 = l0
																							}
																							v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+104))
																							v230 = F_dictFetchValue(m, v229, v225)
																							mBase = m.M
																							v231 = m.ExcPending
																							if v231 != 0 {
																								return int32(0)
																							} else {
																								F_sdsfree(m, v225)
																								mBase = m.M
																								v233 = m.ExcPending
																								if v233 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a1152)
																									if v230 != 0 {
																										v237 = v230
																									} else {
																										v237 = int32(_a1145)
																									}
																									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v237
																									v244 = F_valkeyAsyncCommand(m, v223, int32(1004), l0, int32(_a1159), v9+int32(16))
																									mBase = m.M
																									v245 = m.ExcPending
																									if v245 != 0 {
																										return int32(0)
																									} else {
																										if v244 == int32(-1) {
																											v281 = int32(-1)
																											m.G0 = v9 + int32(144)
																											return v281
																										} else {
																											v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																											v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
																											*(*int32)(unsafe.Add(mBase, uint32(v248)+8)) = v249 + int32(1)
																											v253 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
																											v255 = F_sdsnew(m, int32(_a1161))
																											mBase = m.M
																											v256 = m.ExcPending
																											if v256 != 0 {
																												return int32(0)
																											} else {
																												v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																												if v257 != 0 {
																													v258 = v257
																												} else {
																													v258 = l0
																												}
																												v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+104))
																												v260 = F_dictFetchValue(m, v259, v255)
																												mBase = m.M
																												v261 = m.ExcPending
																												if v261 != 0 {
																													return int32(0)
																												} else {
																													F_sdsfree(m, v255)
																													mBase = m.M
																													v263 = m.ExcPending
																													if v263 != 0 {
																														return int32(0)
																													} else {
																														if v260 != 0 {
																															v265 = v260
																														} else {
																															v265 = int32(_a1161)
																														}
																														*(*int32)(unsafe.Add(mBase, uint32(v9))) = v265
																														v270 = F_valkeyAsyncCommand(m, v253, int32(1004), l0, int32(_a57), v9)
																														mBase = m.M
																														v271 = m.ExcPending
																														if v271 != 0 {
																															return int32(0)
																														} else {
																															if v270 == int32(-1) {
																																v281 = int32(-1)
																															} else {
																																v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																																v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+8))
																																*(*int32)(unsafe.Add(mBase, uint32(v274)+8)) = v275 + int32(1)
																																v281 = int32(0)
																															}
																															m.G0 = v9 + int32(144)
																															return v281
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
						} else {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
							v104 = F_sdsnew(m, int32(_a1004))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return int32(0)
							} else {
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
								if v106 != 0 {
									v107 = v106
								} else {
									v107 = l0
								}
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+104))
								v109 = F_dictFetchValue(m, v108, v104)
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return int32(0)
								} else {
									F_sdsfree(m, v104)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int32(0)
									} else {
										if v109 != 0 {
											v114 = v109
										} else {
											v114 = int32(_a1004)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v114
										v121 = F_valkeyAsyncCommand(m, v102, int32(1004), l0, int32(_a1165), v9+int32(80))
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int32(0)
										} else {
											if v121 == int32(-1) {
												v281 = int32(-1)
												m.G0 = v9 + int32(144)
												return v281
											} else {
												v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v125)+8)) = v126 + int32(1)
												v130 = v125
												v133 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
												v135 = F_sdsnew(m, int32(_a1163))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
													if v137 != 0 {
														v138 = v137
													} else {
														v138 = l0
													}
													v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+104))
													v140 = F_dictFetchValue(m, v139, v135)
													mBase = m.M
													v141 = m.ExcPending
													if v141 != 0 {
														return int32(0)
													} else {
														F_sdsfree(m, v135)
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v67
															if v140 != 0 {
																v146 = v140
															} else {
																v146 = int32(_a1163)
															}
															*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v146
															*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v9 + int32(112)
															v156 = F_valkeyAsyncCommand(m, v133, int32(1004), l0, int32(_a1164), v9+int32(64))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																if v156 == int32(-1) {
																	v281 = int32(-1)
																	m.G0 = v9 + int32(144)
																	return v281
																} else {
																	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v161 + int32(1)
																	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
																	v167 = F_sdsnew(m, int32(_a1156))
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return int32(0)
																	} else {
																		v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																		if v169 != 0 {
																			v170 = v169
																		} else {
																			v170 = l0
																		}
																		v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+104))
																		v172 = F_dictFetchValue(m, v171, v167)
																		mBase = m.M
																		v173 = m.ExcPending
																		if v173 != 0 {
																			return int32(0)
																		} else {
																			F_sdsfree(m, v167)
																			mBase = m.M
																			v175 = m.ExcPending
																			if v175 != 0 {
																				return int32(0)
																			} else {
																				if v172 != 0 {
																					v177 = v172
																				} else {
																					v177 = int32(_a1156)
																				}
																				*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v177
																				v184 = F_valkeyAsyncCommand(m, v165, int32(1004), l0, int32(_a1157), v9+int32(48))
																				mBase = m.M
																				v185 = m.ExcPending
																				if v185 != 0 {
																					return int32(0)
																				} else {
																					if v184 == int32(-1) {
																						v281 = int32(-1)
																						m.G0 = v9 + int32(144)
																						return v281
																					} else {
																						v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																						v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v188)+8)) = v189 + int32(1)
																						v193 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
																						v195 = F_sdsnew(m, int32(_a1145))
																						mBase = m.M
																						v196 = m.ExcPending
																						if v196 != 0 {
																							return int32(0)
																						} else {
																							v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																							if v197 != 0 {
																								v198 = v197
																							} else {
																								v198 = l0
																							}
																							v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+104))
																							v200 = F_dictFetchValue(m, v199, v195)
																							mBase = m.M
																							v201 = m.ExcPending
																							if v201 != 0 {
																								return int32(0)
																							} else {
																								F_sdsfree(m, v195)
																								mBase = m.M
																								v203 = m.ExcPending
																								if v203 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(_a1158)
																									if v200 != 0 {
																										v207 = v200
																									} else {
																										v207 = int32(_a1145)
																									}
																									*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v207
																									v214 = F_valkeyAsyncCommand(m, v193, int32(1004), l0, int32(_a1159), v9+int32(32))
																									mBase = m.M
																									v215 = m.ExcPending
																									if v215 != 0 {
																										return int32(0)
																									} else {
																										if v214 == int32(-1) {
																											v281 = int32(-1)
																											m.G0 = v9 + int32(144)
																											return v281
																										} else {
																											v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																											v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+8))
																											*(*int32)(unsafe.Add(mBase, uint32(v218)+8)) = v219 + int32(1)
																											v223 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
																											v225 = F_sdsnew(m, int32(_a1145))
																											mBase = m.M
																											v226 = m.ExcPending
																											if v226 != 0 {
																												return int32(0)
																											} else {
																												v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																												if v227 != 0 {
																													v228 = v227
																												} else {
																													v228 = l0
																												}
																												v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+104))
																												v230 = F_dictFetchValue(m, v229, v225)
																												mBase = m.M
																												v231 = m.ExcPending
																												if v231 != 0 {
																													return int32(0)
																												} else {
																													F_sdsfree(m, v225)
																													mBase = m.M
																													v233 = m.ExcPending
																													if v233 != 0 {
																														return int32(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a1152)
																														if v230 != 0 {
																															v237 = v230
																														} else {
																															v237 = int32(_a1145)
																														}
																														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v237
																														v244 = F_valkeyAsyncCommand(m, v223, int32(1004), l0, int32(_a1159), v9+int32(16))
																														mBase = m.M
																														v245 = m.ExcPending
																														if v245 != 0 {
																															return int32(0)
																														} else {
																															if v244 == int32(-1) {
																																v281 = int32(-1)
																																m.G0 = v9 + int32(144)
																																return v281
																															} else {
																																v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																																v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
																																*(*int32)(unsafe.Add(mBase, uint32(v248)+8)) = v249 + int32(1)
																																v253 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
																																v255 = F_sdsnew(m, int32(_a1161))
																																mBase = m.M
																																v256 = m.ExcPending
																																if v256 != 0 {
																																	return int32(0)
																																} else {
																																	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																																	if v257 != 0 {
																																		v258 = v257
																																	} else {
																																		v258 = l0
																																	}
																																	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+104))
																																	v260 = F_dictFetchValue(m, v259, v255)
																																	mBase = m.M
																																	v261 = m.ExcPending
																																	if v261 != 0 {
																																		return int32(0)
																																	} else {
																																		F_sdsfree(m, v255)
																																		mBase = m.M
																																		v263 = m.ExcPending
																																		if v263 != 0 {
																																			return int32(0)
																																		} else {
																																			if v260 != 0 {
																																				v265 = v260
																																			} else {
																																				v265 = int32(_a1161)
																																			}
																																			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v265
																																			v270 = F_valkeyAsyncCommand(m, v253, int32(1004), l0, int32(_a57), v9)
																																			mBase = m.M
																																			v271 = m.ExcPending
																																			if v271 != 0 {
																																				return int32(0)
																																			} else {
																																				if v270 == int32(-1) {
																																					v281 = int32(-1)
																																				} else {
																																					v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																																					v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+8))
																																					*(*int32)(unsafe.Add(mBase, uint32(v274)+8)) = v275 + int32(1)
																																					v281 = int32(0)
																																				}
																																				m.G0 = v9 + int32(144)
																																				return v281
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
						}
					}
				}
			}
		}
	}
}
func F_sentinelSetClientName(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v25 int32
	_ = v25
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(_a1143)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l2
	v19 = F_snprintf(m, v8+int32(32), int32(64), int32(_a1144), v8+int32(16))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v22 = F_sdsnew(m, int32(_a1145))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
			if v24 != 0 {
				v25 = v24
			} else {
				v25 = l0
			}
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+104))
			v27 = F_dictFetchValue(m, v26, v22)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_sdsfree(m, v22)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					if v27 != 0 {
						v32 = v27
					} else {
						v32 = int32(_a1145)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v32
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v8 + int32(32)
					v39 = F_valkeyAsyncCommand(m, l1, int32(1004), l0, int32(_a1146), v8)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						if v39 != 0 {
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v42 + int32(1)
						}
						m.G0 = v8 + int32(96)
						return
					}
				}
			}
		}
	}
}
func F_sentinelSimFailureCrash(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v2 {
		m.Env.Exit(m, int32(99))
		mBase = m.M
		base.Wasm_trap_unreachable()
		for {
		}
	} else {
		F__serverLog(m, int32(3), int32(_a1166), int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			m.Env.Exit(m, int32(99))
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_sentinelStartFailover(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v56 int32
	_ = v56
	var v60 int64
	_ = v60
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9&int32(1) != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v9 | int32(64)
		v23 = int32(0)
		v25 = *(*int64)(unsafe.Add(mBase, _consts[652]))
		v27 = v25 + int64(1)
		*(*int64)(unsafe.Add(mBase, _consts[652])) = v27
		*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v27
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v27
		F_sentinelEvent(m, int32(3), int32(_a1221), l0, int32(_a408), v7)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			F_sentinelEvent(m, int32(3), int32(_a1224), l0, int32(_a1138), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				v42 = F_mstime(m)
				mBase = m.M
				v44 = int32(0)
				v46 = *(*int64)(unsafe.Add(mBase, _consts[298]))
				v50 = v46*int64(6364136223846793005) + int64(1)
				*(*int64)(unsafe.Add(mBase, _consts[298])) = v50
				v56 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v50)>>(uint(int64(33))%64))), int32(1000))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+256)) = v42 + base.I64_extend_i32_s(v56)
				v60 = F_mstime(m)
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v60
				m.G0 = v7 + int32(16)
				return
			}
		}
	} else {
		F__serverAssert(m, int32(_a1225), int32(_a1135), int32(4939))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
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
func F_sentinelTimer(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	v2 = F_mstime(m)
	mBase = m.M
	v4 = *(*int64)(unsafe.Add(mBase, _consts[663]))
	v5 = v2 - v4
	if v5 < int64(0) {
		v11 = int32(0)
		v12 = int32(1)
		*(*int32)(unsafe.Add(mBase, _consts[645])) = v12
		v15 = F_mstime(m)
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, _consts[660])) = v15
		v19 = *(*int32)(unsafe.Add(mBase, _consts[659]))
		*(*int32)(unsafe.Add(mBase, _consts[659])) = v19 + v12
		F_sentinelEvent(m, int32(3), int32(_a1237), v11, int32(_a1238), v11)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v30 = int32(0)
			v31 = F_mstime(m)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, _consts[663])) = v31
			v34 = *(*int32)(unsafe.Add(mBase, _consts[640]))
			F_sentinelHandleDictOfValkeyInstances(m, v34)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_sentinelRunPendingScripts(m)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_sentinelCollectTerminatedScripts(m)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_sentinelKillTimedoutScripts(m)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v45 = int32(0)
							v47 = *(*int64)(unsafe.Add(mBase, _consts[298]))
							v51 = v47*int64(6364136223846793005) + int64(1)
							*(*int64)(unsafe.Add(mBase, _consts[298])) = v51
							v56 = int32(10)
							v57 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v51)>>(uint(int64(33))%64))), v56)
							*(*int32)(unsafe.Add(mBase, _consts[149])) = v57 + v56
							return
						}
					}
				}
			}
		}
	} else {
		v9 = *(*int64)(unsafe.Add(mBase, _consts[664]))
		if v5 <= v9 {
			v30 = int32(0)
			v31 = F_mstime(m)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, _consts[663])) = v31
			v34 = *(*int32)(unsafe.Add(mBase, _consts[640]))
			F_sentinelHandleDictOfValkeyInstances(m, v34)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_sentinelRunPendingScripts(m)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_sentinelCollectTerminatedScripts(m)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_sentinelKillTimedoutScripts(m)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v45 = int32(0)
							v47 = *(*int64)(unsafe.Add(mBase, _consts[298]))
							v51 = v47*int64(6364136223846793005) + int64(1)
							*(*int64)(unsafe.Add(mBase, _consts[298])) = v51
							v56 = int32(10)
							v57 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v51)>>(uint(int64(33))%64))), v56)
							*(*int32)(unsafe.Add(mBase, _consts[149])) = v57 + v56
							return
						}
					}
				}
			}
		} else {
			v11 = int32(0)
			v12 = int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[645])) = v12
			v15 = F_mstime(m)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, _consts[660])) = v15
			v19 = *(*int32)(unsafe.Add(mBase, _consts[659]))
			*(*int32)(unsafe.Add(mBase, _consts[659])) = v19 + v12
			F_sentinelEvent(m, int32(3), int32(_a1237), v11, int32(_a1238), v11)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v30 = int32(0)
				v31 = F_mstime(m)
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, _consts[663])) = v31
				v34 = *(*int32)(unsafe.Add(mBase, _consts[640]))
				F_sentinelHandleDictOfValkeyInstances(m, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					F_sentinelRunPendingScripts(m)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						F_sentinelCollectTerminatedScripts(m)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_sentinelKillTimedoutScripts(m)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								v45 = int32(0)
								v47 = *(*int64)(unsafe.Add(mBase, _consts[298]))
								v51 = v47*int64(6364136223846793005) + int64(1)
								*(*int64)(unsafe.Add(mBase, _consts[298])) = v51
								v56 = int32(10)
								v57 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v51)>>(uint(int64(33))%64))), v56)
								*(*int32)(unsafe.Add(mBase, _consts[149])) = v57 + v56
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_sentinelVoteLeader(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int64
	_ = v168
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v176 int64
	_ = v176
	var v182 int32
	_ = v182
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v14 = *(*int64)(unsafe.Add(mBase, _consts[652]))
	if base.Ui64(l1) <= base.Ui64(v14) {
		v67 = v14
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	if base.Ui64(l1) <= base.Ui64(v70) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v16 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[652])) = l1
	v18 = int32(_a69)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[149]))
	*(*int32)(unsafe.Add(mBase, _consts[149])) = int32(10)
	v24 = *(*int32)(unsafe.Add(mBase, _consts[178]))
	v26 = F_rewriteConfig(m, v24, v16)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v30 = int32(_a69)
	*(*int32)(unsafe.Add(mBase, _consts[149])) = v19
	v33 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v26 != int32(-1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v56 = *(*int64)(unsafe.Add(mBase, _consts[652]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v56
	F_sentinelEvent(m, int32(3), int32(_a1221), l0, int32(_a408), v11+int32(32))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L3
	} else {
		goto L14
	}
L6:
	;
	if int32(2) < v33 {
		goto L5
	} else {
		goto L12
	}
L7:
	;
	if int32(3) < v33 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v40 = F___strerror_l(m, v39, v39)
	mBase = m.M
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v40
	F__serverLog(m, int32(3), int32(_a1169), v11+int32(48))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	goto L5
L12:
	;
	F__serverLog(m, int32(2), int32(_a1167), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L5
L14:
	;
	v66 = *(*int64)(unsafe.Add(mBase, _consts[652]))
	v67 = v66
	goto L1
L15:
	;
	v188 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v190 != 0 {
		goto L46
	} else {
		goto L47
	}
L16:
	;
	if base.Ui64(l1) < base.Ui64(v67) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	F_sdsfree(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v76 = F_sdsnew(m, l2)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v76
	v79 = int32(0)
	v80 = *(*int64)(unsafe.Add(mBase, _consts[652]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v80
	v82 = int32(_a69)
	v83 = *(*int32)(unsafe.Add(mBase, _consts[149]))
	*(*int32)(unsafe.Add(mBase, _consts[149])) = int32(10)
	v88 = *(*int32)(unsafe.Add(mBase, _consts[178]))
	v90 = F_rewriteConfig(m, v88, v79)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v92 = int32(_a69)
	*(*int32)(unsafe.Add(mBase, _consts[149])) = v83
	v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v90 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v118 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v117
	F_sentinelEvent(m, int32(3), int32(_a1222), l0, int32(_a1223), v11)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L30
	}
L22:
	;
	if int32(2) < v95 {
		goto L21
	} else {
		goto L28
	}
L23:
	;
	if int32(3) < v95 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L25
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v102 = F___strerror_l(m, v101, v101)
	mBase = m.M
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v102
	F__serverLog(m, int32(3), int32(_a1169), v11+int32(16))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	F__serverLog(m, int32(2), int32(_a1167), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	goto L21
L30:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v127 = int32(_a1143)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v162-v164 == int32(0) {
		goto L15
	} else {
		goto L43
	}
L32:
	;
	v162 = F_tolower(m, v158)
	mBase = m.M
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v164 = F_tolower(m, v163)
	mBase = m.M
	goto L31
L33:
	;
	v132 = v126
	v133 = v127
	v134 = v130
	goto L36
L34:
	;
	v158 = int32(0)
	v159 = v127
	goto L32
L35:
	;
	v158 = v155 & int32(255)
	v159 = v154
	goto L32
L36:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v136 == int32(0) {
		v154 = v133
		v155 = v134
		goto L35
	} else {
		goto L38
	}
L37:
	;
	v154 = v148
	v155 = int32(0)
	goto L35
L38:
	;
	v140 = v134 & int32(255)
	if v140 == v136 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v147 = int32(1)
	v148 = v133 + v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v149 != 0 {
		v132 = v132 + v147
		v133 = v148
		v134 = v149
		goto L36
	} else {
		goto L42
	}
L40:
	;
	v142 = F_tolower(m, v140)
	mBase = m.M
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v144 = F_tolower(m, v143)
	mBase = m.M
	if v142 == v144 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v154 = v133
	v155 = v146
	goto L35
L42:
	;
	goto L37
L43:
	;
	v168 = F_mstime(m)
	mBase = m.M
	v170 = int32(0)
	v172 = *(*int64)(unsafe.Add(mBase, _consts[298]))
	v176 = v172*int64(6364136223846793005) + int64(1)
	*(*int64)(unsafe.Add(mBase, _consts[298])) = v176
	goto L44
L44:
	;
	v182 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v176)>>(uint(int64(33))%64))), int32(1000))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+256)) = v168 + base.I64_extend_i32_s(v182)
	goto L15
L45:
	;
	m.G0 = v11 + int32(64)
	return v194
L46:
	;
	v192 = F_sdsnew(m, v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L3
	} else {
		goto L48
	}
L47:
	;
	v194 = int32(0)
	goto L45
L48:
	;
	v194 = v192
	goto L45
}
