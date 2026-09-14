package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_raxAllocSize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	return v2
}
func F_raxFreeWithCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_raxRecursiveFree(m, l0, v3, l1)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		if v6 == int64(0) {
			F_valkey_free(m, l0)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				return
			}
		} else {
			F__serverAssert(m, int32(_a_F_raxFreeWithCallback_0), int32(_a_F_raxFreeWithCallback_1), int32(1236))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
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
func F_raxInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_raxGenericInsert(m, l0, l1, l2, l3, l4, int32(1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_raxIteratorAddChars(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
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
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v19 != 0 {
		goto L58
	} else {
		goto L59
	}
L2:
	;
	return int32(1)
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = v11 + l2
	if base.Ui32(v12) <= base.Ui32(v10) {
		v36 = v11
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v41 = v40 + v36
	if v41 == l1 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = l0 + int32(24)
	if v15 == v17 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = int32(0)
	goto L8
L7:
	;
	v19 = v15
	goto L8
L8:
	;
	v21 = v12 << (uint(int32(1)) % 32)
	v22 = F_valkey_realloc(m, v19, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v22
	if v22 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v19 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v21
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v36 = v35
	goto L4
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v29 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	goto L14
L16:
	;
	v32 = F__emscripten_memcpy_bulkmem(m, v22, v17, v29)
	mBase = m.M
	goto L15
L17:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v190 + l2
	goto L2
L18:
	;
	goto L17
L19:
	;
	v45 = l2 + v41
	if base.Ui32(int32(0)-l2<<(uint(int32(1))%32)) < base.Ui32(l1-v45) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v55 = (l1 ^ v41) & int32(3)
	if base.Ui32(l1) <= base.Ui32(v41) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v52 = F___memcpy(m, v41, l1, l2)
	mBase = m.M
	goto L17
L22:
	;
	if v161 == int32(0) {
		goto L18
	} else {
		goto L54
	}
L23:
	;
	if base.Ui32(v139) <= base.Ui32(int32(3)) {
		v160 = v138
		v161 = v139
		v162 = v140
		goto L22
	} else {
		goto L50
	}
L24:
	;
	if v55 != 0 {
		v121 = l2
		goto L34
	} else {
		goto L35
	}
L25:
	;
	if v55 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if v41&int32(3) != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v160 = l1
	v161 = l2
	v162 = v41
	goto L22
L28:
	;
	v62 = l1
	v63 = l2
	v64 = v41
	goto L30
L29:
	;
	v138 = l1
	v139 = l2
	v140 = v41
	goto L23
L30:
	;
	if v63 == int32(0) {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v68)
	v70 = int32(1)
	v71 = v62 + v70
	v73 = v63 + int32(-1)
	v75 = v64 + v70
	if v75&int32(3) == int32(0) {
		v138 = v71
		v139 = v73
		v140 = v75
		goto L23
	} else {
		goto L33
	}
L33:
	;
	v62 = v71
	v63 = v73
	v64 = v75
	goto L30
L34:
	;
	if v121 == int32(0) {
		goto L18
	} else {
		goto L46
	}
L35:
	;
	if v45&int32(3) == int32(0) {
		v101 = l2
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if base.Ui32(v101) <= base.Ui32(int32(3)) {
		v121 = v101
		goto L34
	} else {
		goto L42
	}
L37:
	;
	v86 = l2
	goto L38
L38:
	;
	if v86 == int32(0) {
		goto L18
	} else {
		goto L40
	}
L39:
	;
	v101 = v92
	goto L36
L40:
	;
	v92 = v86 + int32(-1)
	v93 = v41 + v92
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v92))))
	*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v95)
	if v93&int32(3) != 0 {
		v86 = v92
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v108 = v101
	goto L43
L43:
	;
	v112 = v108 + int32(-4)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1+v112)))
	*(*int32)(unsafe.Add(mBase, uint32(v41+v112))) = v115
	if base.Ui32(int32(3)) < base.Ui32(v112) {
		v108 = v112
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v121 = v112
	goto L34
L45:
	;
	goto L44
L46:
	;
	v128 = v121
	goto L47
L47:
	;
	v132 = v128 + int32(-1)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v132))))
	*(*uint8)(unsafe.Add(mBase, uint32(v41+v132))) = uint8(v135)
	if v132 != 0 {
		v128 = v132
		goto L47
	} else {
		goto L49
	}
L49:
	;
	goto L18
L50:
	;
	v145 = v138
	v146 = v139
	v147 = v140
	goto L51
L51:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v149
	v151 = int32(4)
	v152 = v145 + v151
	v154 = v147 + v151
	v156 = v146 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v156) {
		v145 = v152
		v146 = v156
		v147 = v154
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v160 = v152
	v161 = v156
	v162 = v154
	goto L22
L53:
	;
	goto L52
L54:
	;
	v167 = v160
	v168 = v161
	v169 = v162
	goto L55
L55:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	*(*uint8)(unsafe.Add(mBase, uint32(v169))) = uint8(v171)
	v173 = int32(1)
	v178 = v168 + int32(-1)
	if v178 != 0 {
		v167 = v167 + v173
		v168 = v178
		v169 = v169 + v173
		goto L55
	} else {
		goto L57
	}
L56:
	;
	goto L18
L57:
	;
	goto L56
L58:
	;
	v199 = v19
	goto L60
L59:
	;
	v199 = v17
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v199
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxIteratorAddChars[0])) = int32(48)
	return int32(0)
}
func F_raxLowWalk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
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
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
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
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	v8 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if base.Ui32(int32(8)) <= base.Ui32(v17) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l3 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L2:
	;
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v208 = l0
	v215 = v8
	v216 = v16
	v218 = int32(0)
	goto L1
L4:
	;
	v23 = l6 + int32(12)
	v25 = l0
	v32 = int32(0)
	v33 = v16
	v34 = v17
	goto L6
L5:
	;
	v208 = l0
	v215 = v8
	v216 = v16
	v218 = int32(0)
	goto L1
L6:
	;
	v40 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
	v41 = int32(4)
	v42 = v33 + v41
	if v34&v41 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v208 = v202
	v215 = v124
	v216 = v203
	v218 = v186
	goto L1
L8:
	;
	if l6 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L9:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v32))))
	v104 = int32(0)
	goto L20
L10:
	;
	v47 = int32(0)
	if base.Ui32(l2) <= base.Ui32(v32) {
		v81 = v32
		v84 = v47
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v84 == v40 {
		v124 = v81
		v129 = v40
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v56 = v32
	v59 = v47
	goto L13
L13:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v59))))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v56))))
	if v64 != v66 {
		v81 = v56
		v84 = v59
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v81 = v69
	v84 = v71
	goto L11
L15:
	;
	v68 = int32(1)
	v69 = v56 + v68
	v71 = v59 + v68
	if base.Ui32(v40) <= base.Ui32(v71) {
		v81 = v69
		v84 = v71
		goto L11
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v69) < base.Ui32(l2) {
		v56 = v69
		v59 = v71
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v208 = v25
	v215 = v81
	v216 = v33
	v218 = v84
	goto L1
L19:
	;
	if v104 != v40 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v104))))
	if v109 == v91&int32(255) {
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v208 = v25
	v215 = v32
	v216 = v33
	v218 = v40
	goto L1
L22:
	;
	v112 = v104 + int32(1)
	if v112 != v40 {
		v104 = v112
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v124 = v32 + int32(1)
	v129 = v104
	goto L8
L25:
	;
	v208 = v25
	v215 = v32
	v216 = v33
	v218 = v40
	goto L1
L26:
	;
	v186 = int32(0)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v188 = int32(3)
	v189 = int32(base.Ui32(v187) >> (uint(v188) % 32))
	if v187&int32(4) != 0 {
		goto L45
	} else {
		goto L46
	}
L27:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v134 != v135 {
		v175 = v134
		v176 = v133
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176+v175<<(uint(int32(2))%32)))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v175 + int32(1)
	goto L26
L29:
	;
	v138 = v134 << (uint(int32(3)) % 32)
	if v133 != v23 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = v169 << (uint(int32(1)) % 32)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v175 = v174
	v176 = v170
	goto L28
L31:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v164 = v162 << (uint(int32(2)) % 32)
	if v164 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L32:
	;
	v151 = F_valkey_realloc(m, v133, v138)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L34
	} else {
		goto L39
	}
L33:
	;
	v140 = F_valkey_malloc(m, v138)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return int32(0)
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v140
	if v140 != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+140)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v23
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxLowWalk[0])) = int32(48)
	goto L26
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+140)) = int32(1)
	goto L41
L39:
	;
	if v151 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v151
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v169 = v156
	v170 = v151
	goto L30
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxLowWalk[0])) = int32(48)
	goto L26
L42:
	;
	v169 = v162
	v170 = v140
	goto L30
L43:
	;
	goto L42
L44:
	;
	v167 = F__emscripten_memcpy_bulkmem(m, v140, v23, v164)
	mBase = m.M
	goto L43
L45:
	;
	v199 = v186
	goto L47
L46:
	;
	v199 = v129
	goto L47
L47:
	;
	v202 = v42 + v189 + (v186-v189)&v188 + v199<<(uint(int32(2))%32)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if base.Ui32(v204) < base.Ui32(int32(8)) {
		v208 = v202
		v215 = v124
		v216 = v203
		v218 = v186
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if base.Ui32(v124) < base.Ui32(l2) {
		v25 = v202
		v32 = v124
		v33 = v203
		v34 = v204
		goto L6
	} else {
		goto L49
	}
L49:
	;
	goto L7
L50:
	;
	if l4 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v216
	goto L50
L52:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	if v228&int32(4) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v208
	goto L52
L54:
	;
	return v215
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v218
	goto L54
}
func F_raxNew(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v4 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = int64(1)
			*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = int64(0)
			v15 = F_valkey_malloc(m, int32(4))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v15 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v15
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v4+int32(-8))))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-8))))
				*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = v24&int32(2147483647) + int32(8) + (v32&int32(2147483647) + int32(8))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
				if v39 == int32(0) {
					F_valkey_free(m, v4)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					return v4
				}
			}
		}
	}
}
func F_raxSeekGreatest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
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
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
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
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v352 int32
	_ = v352
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if base.Ui32(v11) < base.Ui32(int32(8)) {
		v352 = int32(1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v352
L2:
	;
	v15 = l0 + int32(168)
	v17 = l0 + int32(24)
	v20 = v10
	v21 = v11
	goto L6
L3:
	;
	goto L101
L4:
	;
	if v220 != 0 {
		goto L98
	} else {
		goto L99
	}
L5:
	;
	if v41 != 0 {
		goto L95
	} else {
		goto L96
	}
L6:
	;
	v26 = int32(4)
	v27 = v20 + v26
	v29 = int32(base.Ui32(v21) >> (uint(int32(3)) % 32))
	if v21&v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v251 + v248
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v257 != v258 {
		v290 = v256
		v291 = v257
		goto L77
	} else {
		goto L78
	}
L9:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v216 = v214 + int32(1)
	if base.Ui32(v216) <= base.Ui32(v213) {
		v238 = v212
		v239 = v214
		goto L65
	} else {
		goto L66
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v37 = v36 + v29
	if base.Ui32(v37) <= base.Ui32(v35) {
		v61 = v34
		v62 = v36
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v63 = v61 + v62
	if v63 == v27 {
		goto L25
	} else {
		goto L26
	}
L12:
	;
	if v34 == v17 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v41 = int32(0)
	goto L15
L14:
	;
	v41 = v34
	goto L15
L15:
	;
	v43 = v37 << (uint(int32(1)) % 32)
	v44 = F_valkey_realloc(m, v41, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v44
	if v44 == int32(0) {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	if v41 != 0 {
		v57 = v44
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v43
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v61 = v57
	v62 = v59
	goto L11
L20:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v51 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v57 = v56
	goto L19
L22:
	;
	goto L21
L23:
	;
	v54 = F__emscripten_memcpy_bulkmem(m, v44, v17, v51)
	mBase = m.M
	goto L22
L24:
	;
	v248 = v29
	goto L8
L25:
	;
	goto L24
L26:
	;
	v67 = v29 + v63
	if base.Ui32(int32(0)-v29<<(uint(int32(1))%32)) < base.Ui32(v27-v67) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v77 = (v27 ^ v63) & int32(3)
	if base.Ui32(v27) <= base.Ui32(v63) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v74 = F___memcpy(m, v63, v27, v29)
	mBase = m.M
	goto L24
L29:
	;
	if v183 == int32(0) {
		goto L25
	} else {
		goto L61
	}
L30:
	;
	if base.Ui32(v161) <= base.Ui32(int32(3)) {
		v182 = v160
		v183 = v161
		v184 = v162
		goto L29
	} else {
		goto L57
	}
L31:
	;
	if v77 != 0 {
		v143 = v29
		goto L41
	} else {
		goto L42
	}
L32:
	;
	if v77 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if v63&int32(3) != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v182 = v27
	v183 = v29
	v184 = v63
	goto L29
L35:
	;
	v84 = v27
	v85 = v29
	v86 = v63
	goto L37
L36:
	;
	v160 = v27
	v161 = v29
	v162 = v63
	goto L30
L37:
	;
	if v85 == int32(0) {
		goto L25
	} else {
		goto L39
	}
L39:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v90)
	v92 = int32(1)
	v93 = v84 + v92
	v95 = v85 + int32(-1)
	v97 = v86 + v92
	if v97&int32(3) == int32(0) {
		v160 = v93
		v161 = v95
		v162 = v97
		goto L30
	} else {
		goto L40
	}
L40:
	;
	v84 = v93
	v85 = v95
	v86 = v97
	goto L37
L41:
	;
	if v143 == int32(0) {
		goto L25
	} else {
		goto L53
	}
L42:
	;
	if v67&int32(3) == int32(0) {
		v123 = v29
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if base.Ui32(v123) <= base.Ui32(int32(3)) {
		v143 = v123
		goto L41
	} else {
		goto L49
	}
L44:
	;
	v108 = v29
	goto L45
L45:
	;
	if v108 == int32(0) {
		goto L25
	} else {
		goto L47
	}
L46:
	;
	v123 = v114
	goto L43
L47:
	;
	v114 = v108 + int32(-1)
	v115 = v63 + v114
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v114))))
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v117)
	if v115&int32(3) != 0 {
		v108 = v114
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v130 = v123
	goto L50
L50:
	;
	v134 = v130 + int32(-4)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v27+v134)))
	*(*int32)(unsafe.Add(mBase, uint32(v63+v134))) = v137
	if base.Ui32(int32(3)) < base.Ui32(v134) {
		v130 = v134
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v143 = v134
	goto L41
L52:
	;
	goto L51
L53:
	;
	v150 = v143
	goto L54
L54:
	;
	v154 = v150 + int32(-1)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v154))))
	*(*uint8)(unsafe.Add(mBase, uint32(v63+v154))) = uint8(v157)
	if v154 != 0 {
		v150 = v154
		goto L54
	} else {
		goto L56
	}
L56:
	;
	goto L25
L57:
	;
	v167 = v160
	v168 = v161
	v169 = v162
	goto L58
L58:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v171
	v173 = int32(4)
	v174 = v167 + v173
	v176 = v169 + v173
	v178 = v168 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v178) {
		v167 = v174
		v168 = v178
		v169 = v176
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v182 = v174
	v183 = v178
	v184 = v176
	goto L29
L60:
	;
	goto L59
L61:
	;
	v189 = v182
	v190 = v183
	v191 = v184
	goto L62
L62:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v193)
	v195 = int32(1)
	v200 = v190 + int32(-1)
	if v200 != 0 {
		v189 = v189 + v195
		v190 = v200
		v191 = v191 + v195
		goto L62
	} else {
		goto L64
	}
L63:
	;
	goto L25
L64:
	;
	goto L63
L65:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v29+int32(-1)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v238+v239))) = uint8(v244)
	v248 = int32(1)
	goto L8
L66:
	;
	if v212 == v17 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v220 = int32(0)
	goto L69
L68:
	;
	v220 = v212
	goto L69
L69:
	;
	v222 = v216 << (uint(int32(1)) % 32)
	v223 = F_valkey_realloc(m, v220, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L16
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v223
	if v223 == int32(0) {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	if v220 != 0 {
		v234 = v223
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v222
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v238 = v234
	v239 = v236
	goto L65
L73:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v228 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v234 = v233
	goto L72
L75:
	;
	goto L74
L76:
	;
	v231 = F__emscripten_memcpy_bulkmem(m, v223, v17, v228)
	mBase = m.M
	goto L75
L77:
	;
	v292 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v290+v291<<(uint(v292)%32)))) = v254
	v296 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v291 + v296
	v300 = int32(3)
	v301 = int32(base.Ui32(v255) >> (uint(v300) % 32))
	v308 = int32(4)
	if v255&v308 != 0 {
		goto L91
	} else {
		goto L92
	}
L78:
	;
	v261 = v257 << (uint(int32(3)) % 32)
	if v256 != v15 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v285 << (uint(int32(1)) % 32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v290 = v284
	v291 = v289
	goto L77
L80:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v279 = v277 << (uint(int32(2)) % 32)
	if v279 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L81:
	;
	v269 = F_valkey_realloc(m, v256, v261)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L16
	} else {
		goto L86
	}
L82:
	;
	v263 = F_valkey_malloc(m, v261)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L16
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v263
	if v263 != 0 {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v15
	goto L3
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = int32(1)
	goto L3
L86:
	;
	if v269 == int32(0) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v269
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v284 = v269
	v285 = v274
	goto L79
L88:
	;
	v284 = v263
	v285 = v277
	goto L79
L89:
	;
	goto L88
L90:
	;
	v282 = F__emscripten_memcpy_bulkmem(m, v263, v15, v279)
	mBase = m.M
	goto L89
L91:
	;
	v313 = v308
	goto L93
L92:
	;
	v313 = v301 << (uint(v292) % 32)
	goto L93
L93:
	;
	v315 = int32(1)
	v316 = v255 << (uint(v315) % 32)
	v322 = int32(0) - v255&v315
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v254+v301+(int32(0)-v301)&v300+v313+(v316^int32(-1))&v322&int32(4)+v322&(v316|int32(-5)+v315))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v333
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	if base.Ui32(int32(8)) <= base.Ui32(v335) {
		v20 = v333
		v21 = v335
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v352 = v296
	goto L1
L95:
	;
	v338 = v41
	goto L97
L96:
	;
	v338 = v17
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v338
	goto L3
L98:
	;
	v340 = v220
	goto L100
L99:
	;
	v340 = v17
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v340
	goto L3
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxSeekGreatest[0])) = int32(48)
	v352 = int32(0)
	goto L1
}
func F_raxSetData(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5 | int32(3)
		return
	} else {
		v8 = int32(3)
		v9 = int32(base.Ui32(v5) >> (uint(v8) % 32))
		v16 = int32(4)
		if v5&v16 != 0 {
			v21 = v16
		} else {
			v21 = v9 << (uint(int32(2)) % 32)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0+v9+(int32(0)-v9)&v8+v21+int32(4)))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5&int32(-4) | int32(1)
		return
	}
}
