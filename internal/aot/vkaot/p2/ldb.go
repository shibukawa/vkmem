package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_ldbEnd(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v1 = int32(0)
	v4 = m.G6
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+276))
	if v5 <= v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = m.G6
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+272))
	v28 = m.G11
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	m.T0[v29].(func(*base.Module, int32))(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	v8 = v1
	goto L3
L3:
	;
	v10 = m.G6
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+272))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11+v8<<(uint(int32(2))%32))))
	v16 = m.G11
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	m.T0[v17].(func(*base.Module, int32))(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v21 = v8 + int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+276))
	if v21 < v22 {
		v8 = v21
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v32
	return
}
func F_ldbIsEnabled(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v1 = int32(0)
	v2 = m.G6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+264))
	return base.B2i32(v3 != v1) & base.B2i32(v6 != v1)
}
func F_ldbSendLogs(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v1 = m.G14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(v1)))
	m.T0[v2].(func(*base.Module))(m)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_ldbSetCurrentLine(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = m.G6
	*(*int32)(unsafe.Add(mBase, uint32(v2)+280)) = l0
	return
}
func F_ldbStart(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = m.G6
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
	v20 = m.G7
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = m.T0[v21].(func(*base.Module, int32, int32) int32)(m, l0, v13+int32(12))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v24 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if v45 == int32(0) {
		v158 = v126
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v81 = int32(0)
	v88 = int32(1)
	v89 = v81
	v93 = v81
	goto L14
L5:
	;
	v67 = m.G3
	v73 = m.G8
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	m.T0[v74].(func(*base.Module, int32, int32, int32))(m, v67+int32(_a1922), v67+int32(_a1923), int32(67))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v30 = v24
	v35 = int32(0)
	goto L7
L7:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-1)+v30))))
	switch v41 + int32(-10) {
	case 0, 3:
		goto L9
	default:
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	v53 = v30 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v53
	if v53 != 0 {
		v30 = v53
		v35 = v35 + int32(1)
		goto L7
	} else {
		goto L12
	}
L10:
	;
	v45 = v30 & int32(3)
	if base.Ui32(v35-v24) <= base.Ui32(int32(-4)) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v126 = int32(1)
	v127 = int32(0)
	goto L3
L12:
	;
	goto L8
L13:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v94 = v22 + v89
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v96 = int32(10)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+int32(1)))))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+int32(2)))))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+int32(3)))))
	v116 = v88 + base.B2i32(v95 == v96) + base.B2i32(v101 == v96) + base.B2i32(v107 == v96) + base.B2i32(v113 == v96)
	v117 = int32(4)
	v118 = v89 + v117
	v120 = v93 + v117
	if v120 != v30&int32(-4) {
		v88 = v116
		v89 = v118
		v93 = v120
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v126 = v116
	v127 = v118
	goto L3
L16:
	;
	goto L15
L17:
	;
	v165 = m.G9
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v167 = m.T0[v166].(func(*base.Module, int32, int32) int32)(m, v158, int32(4))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L24
	}
L18:
	;
	v138 = v126
	v139 = v127
	v141 = int32(0)
	goto L19
L19:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v139))))
	v148 = v138 + base.B2i32(v145 == int32(10))
	v149 = int32(1)
	v152 = v141 + v149
	if v152 != v45 {
		v138 = v148
		v139 = v139 + v149
		v141 = v152
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v158 = v148
	goto L17
L21:
	;
	goto L20
L22:
	;
	v247 = m.G6
	*(*int32)(unsafe.Add(mBase, uint32(v247)+272)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v247)+276)) = v244
	m.G0 = v13 + int32(16)
	return
L23:
	;
	v225 = int32(0)
	v227 = m.G3
	v233 = m.G10
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	m.T0[v234].(func(*base.Module, int32, int32, int32, int32))(m, v225, v227+int32(_a707), v227+int32(_a1924), v225)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L39
	}
L24:
	;
	if v167 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v171 = int32(0)
	v178 = v171
	v179 = v171
	v181 = v171
	goto L26
L26:
	;
	if v179 == v30 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if base.Ui32(v221) <= base.Ui32(v30) {
		v178 = v220
		v179 = v221
		v181 = v222
		goto L26
	} else {
		goto L38
	}
L29:
	;
	v191 = v179 - v178
	v192 = int32(1)
	v195 = m.G9
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v197 = m.T0[v196].(func(*base.Module, int32, int32) int32)(m, v191+v192, v192)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L33
	}
L30:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v179))))
	if v186 == int32(10) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v220 = v178
	v221 = v179 + int32(1)
	v222 = v181
	goto L28
L32:
	;
	v218 = v179 + int32(1)
	v220 = v218
	v221 = v218
	v222 = v216
	goto L28
L33:
	;
	if v197 == int32(0) {
		v216 = v181
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if v191 == int32(0) {
		v205 = v197
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v207 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v205+v191))) = uint8(v207)
	*(*int32)(unsafe.Add(mBase, uint32(v167+v181<<(uint(int32(2))%32)))) = v205
	v216 = v181 + int32(1)
	goto L32
L36:
	;
	goto L35
L37:
	;
	v204 = F__emscripten_memcpy_bulkmem(m, v197, v22+v178, v191)
	mBase = m.M
	v205 = v204
	goto L36
L38:
	;
	v244 = v222
	goto L22
L39:
	;
	v244 = v225
	goto L22
}
