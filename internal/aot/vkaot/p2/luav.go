package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaV_gettable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v11 = l1
	v14 = v8
	v15 = int32(0)
	goto L2
L1:
	;
	v134 = m.G3
	F_luaG_runerror(m, l0, v134+int32(_a_F_luaV_gettable_0), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L23
	} else {
		goto L32
	}
L2:
	;
	if v14 != int32(5) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
	*(*int64)(unsafe.Add(mBase, uint32(v90))) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	*(*int64)(unsafe.Add(mBase, uint32(v95)+16)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+24)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v101 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v100)+32)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v100)+40)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(48) < v105-v106 {
		v114 = v106
		goto L28
	} else {
		goto L29
	}
L4:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v82 == int32(6) {
		goto L25
	} else {
		goto L26
	}
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	switch v47 + int32(-5) {
	case 0:
		goto L19
	default:
		goto L17
	case 2:
		goto L18
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v20 = F_luaH_get(m, v19, l2)
	mBase = m.M
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v43
	return
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v22 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
	if v25&int32(1) != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+188))
	v31 = F_luaH_getstr(m, v22, v30)
	mBase = m.M
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v32 != 0 {
		v39 = v31
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v39 != 0 {
		v81 = v39
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
	v36 = v33 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)) = uint8(v36)
	v39 = int32(0)
	goto L12
L14:
	;
	goto L7
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	if v74 != 0 {
		v81 = v73
		goto L4
	} else {
		goto L22
	}
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v64 = m.G398
	if v63 == int32(0) {
		v73 = v64
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v62 = v56 + v47<<(uint(int32(2))%32) + int32(152)
	goto L16
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v62 = v53 + int32(8)
	goto L16
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v62 = v50 + int32(16)
	goto L16
L20:
	;
	goto L15
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(0))+188))
	v72 = F_luaH_getstr(m, v63, v71)
	mBase = m.M
	v73 = v72
	goto L20
L22:
	;
	v75 = m.G3
	F_luaG_typeerror(m, l0, v11, v75+int32(_a_F_luaV_gettable_1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return
L24:
	;
	v81 = v73
	goto L4
L25:
	;
	goto L3
L26:
	;
	v86 = v15 + int32(1)
	if v86 == int32(100) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v11 = v81
	v14 = v82
	v15 = v86
	goto L2
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v114 + int32(48)
	F_luaD_call(m, l0, v114, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L23
	} else {
		goto L31
	}
L29:
	;
	F_luaD_growstack(m, l0, int32(3))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v114 = v113
	goto L28
L31:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v123 = v121 + int32(-16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v127 = v125 + (l3 - v89)
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v123)))
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = v128
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v121+int32(-8))))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+8)) = v132
	return
L32:
	;
	return
}
func F_luaV_lessthan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
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
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v7 != v8 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v120
L2:
	;
	v118 = F_luaG_ordererror(m, l0, l1, l2)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L34
	} else {
		goto L37
	}
L3:
	;
	switch v7 + int32(-3) {
	case 0:
		goto L6
	case 1:
		goto L5
	default:
		goto L4
	}
L4:
	;
	v111 = F_call_orderTM(m, l0, l1, l2, int32(13))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L34
	} else {
		goto L35
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v17 = int32(16)
	v18 = v16 + v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v21 = v19 + v17
	v22 = F___get_tp(m)
	mBase = m.M
	v23 = F___strcoll_l(m, v18, v21, v21)
	mBase = m.M
	goto L8
L6:
	;
	v12 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	return base.F64_lt(v12, v13)
L7:
	;
	return int32(base.Ui32(v103) >> (uint(int32(31)) % 32))
L8:
	;
	if v23 != 0 {
		v103 = v23
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v26 = v25
	v27 = v18
	v29 = v24
	v31 = v21
	goto L10
L10:
	;
	if v27&int32(3) == int32(0) {
		v53 = v27
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v103 = v98
	goto L7
L12:
	;
	if v86 != v29 {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	if v86 != v26 {
		goto L12
	} else {
		goto L29
	}
L14:
	;
	v86 = v78 - v27
	goto L13
L15:
	;
	v57 = v53
	goto L23
L16:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v39 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v42 = v27
	goto L19
L18:
	;
	v86 = v27 - v27
	goto L13
L19:
	;
	v46 = v42 + int32(1)
	if v46&int32(3) == int32(0) {
		v53 = v46
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v51 != 0 {
		v42 = v46
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v78 = v46
	goto L14
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v66 = int32(-2139062144)
	if (int32(16843008)-v63|v63)&v66 == v66 {
		v57 = v57 + int32(4)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v72 = v57
	goto L26
L25:
	;
	goto L24
L26:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v76 != 0 {
		v72 = v72 + int32(1)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v78 = v72
	goto L14
L28:
	;
	goto L27
L29:
	;
	v103 = int32(0)
	goto L7
L30:
	;
	v92 = v86 + int32(1)
	v95 = v27 + v92
	v96 = v31 + v92
	v97 = F___get_tp(m)
	mBase = m.M
	v98 = F___strcoll_l(m, v95, v96, v96)
	mBase = m.M
	goto L32
L31:
	;
	v103 = int32(-1)
	goto L7
L32:
	;
	if v98 == int32(0) {
		v26 = v26 - v92
		v27 = v95
		v29 = v29 - v92
		v31 = v96
		goto L10
	} else {
		goto L33
	}
L33:
	;
	goto L11
L34:
	;
	return int32(0)
L35:
	;
	if v111 != int32(-1) {
		v120 = v111
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L2
L37:
	;
	v120 = v118
	goto L1
}
func F_luaV_settable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
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
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int64
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int64
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
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
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int64
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int64
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int64
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v14 == int32(5) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return
L2:
	;
	v245 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v243))) = v245
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+8)) = v247
	v249 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v242)+6)) = uint8(v249)
	if v247 < int32(4) {
		goto L1
	} else {
		goto L62
	}
L3:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v88 == int32(6) {
		v192 = l1
		v196 = v85
		goto L27
	} else {
		goto L28
	}
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v53 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	switch v19 + int32(-5) {
	case 0:
		goto L10
	default:
		goto L8
	case 2:
		goto L9
	}
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v46 != 0 {
		v85 = v45
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v36 = m.G398
	if v35 == int32(0) {
		v45 = v36
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v34 = v28 + v19<<(uint(int32(2))%32) + int32(152)
	goto L7
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v34 = v25 + int32(8)
	goto L7
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v34 = v22 + int32(16)
	goto L7
L11:
	;
	goto L6
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(4))+188))
	v44 = F_luaH_getstr(m, v35, v43)
	mBase = m.M
	v45 = v44
	goto L11
L13:
	;
	v47 = m.G3
	F_luaG_typeerror(m, l0, l1, v47+int32(_a_F_luaV_settable_0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	v85 = v45
	goto L3
L16:
	;
	v62 = F_luaH_set(m, l0, v52, l2)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	v56 = m.G3
	F_luaG_runerror(m, l0, v56+int32(_a_F_luaV_settable_1), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v64 != 0 {
		v242 = v52
		v243 = v62
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v65 == int32(0) {
		v242 = v52
		v243 = v62
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+6)))
	if v68&int32(2) != 0 {
		v242 = v52
		v243 = v62
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+192))
	v74 = F_luaH_getstr(m, v65, v73)
	mBase = m.M
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	if v75 != 0 {
		v82 = v74
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v82 == int32(0) {
		v242 = v52
		v243 = v62
		goto L2
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+6)))
	v79 = v76 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v65)+6)) = uint8(v79)
	v82 = int32(0)
	goto L24
L26:
	;
	v85 = v82
	goto L3
L27:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
	*(*int64)(unsafe.Add(mBase, uint32(v200))) = v201
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+8)) = v203
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v192)))
	*(*int64)(unsafe.Add(mBase, uint32(v205)+16)) = v206
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v205)+24)) = v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v211 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v210)+32)) = v211
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v210)+40)) = v213
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v216 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v215)+48)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+56)) = v218
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(64) < v220-v221 {
		v229 = v221
		goto L58
	} else {
		goto L59
	}
L28:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v91
	v96 = int32(1)
	v101 = v88
	v103 = v91
	goto L29
L29:
	;
	if v101 != int32(5) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v185 = m.G3
	F_luaG_runerror(m, l0, v185+int32(_a_F_luaV_settable_2), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L14
	} else {
		goto L57
	}
L31:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v172)+8))
	if v175 != int32(6) {
		goto L54
	} else {
		goto L55
	}
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	switch v139 + int32(-5) {
	case 0:
		goto L49
	default:
		goto L47
	case 2:
		goto L48
	}
L33:
	;
	v106 = base.I32_wrap_i64(v103)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	if v107 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v116 = F_luaH_set(m, l0, v106, l2)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L14
	} else {
		goto L37
	}
L35:
	;
	v110 = m.G3
	F_luaG_runerror(m, l0, v110+int32(_a_F_luaV_settable_1), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L14
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	if v118 != 0 {
		v242 = v106
		v243 = v116
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	if v119 == int32(0) {
		v242 = v106
		v243 = v116
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)))
	if v122&int32(2) != 0 {
		v242 = v106
		v243 = v116
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+192))
	v128 = F_luaH_getstr(m, v119, v127)
	mBase = m.M
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	if v129 != 0 {
		v136 = v128
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v136 != 0 {
		v172 = v136
		goto L31
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)))
	v133 = v130 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)) = uint8(v133)
	v136 = int32(0)
	goto L42
L44:
	;
	v242 = v106
	v243 = v116
	goto L2
L45:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	if v166 != 0 {
		v172 = v165
		goto L31
	} else {
		goto L52
	}
L46:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v156 = m.G398
	if v155 == int32(0) {
		v165 = v156
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v154 = v148 + v139<<(uint(int32(2))%32) + int32(152)
	goto L46
L48:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v154 = v145 + int32(8)
	goto L46
L49:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v154 = v142 + int32(16)
	goto L46
L50:
	;
	goto L45
L51:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+int32(4))+188))
	v164 = F_luaH_getstr(m, v155, v163)
	mBase = m.M
	v165 = v164
	goto L50
L52:
	;
	v167 = m.G3
	F_luaG_typeerror(m, l0, v12, v167+int32(_a_F_luaV_settable_0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L14
	} else {
		goto L53
	}
L53:
	;
	v172 = v165
	goto L31
L54:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v172)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v175
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v178
	v182 = v96 + int32(1)
	if v182 != int32(100) {
		v96 = v182
		v101 = v175
		v103 = v178
		goto L29
	} else {
		goto L56
	}
L55:
	;
	v192 = v12
	v196 = v172
	goto L27
L56:
	;
	goto L30
L57:
	;
	goto L1
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v229 + int32(64)
	F_luaD_call(m, l0, v229, int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L14
	} else {
		goto L61
	}
L59:
	;
	F_luaD_growstack(m, l0, int32(4))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v229 = v228
	goto L58
L61:
	;
	goto L1
L62:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+5)))
	if v254&int32(3) == int32(0) {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+5)))
	if v259&int32(4) == int32(0) {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+5)))
	v267 = v265 & int32(251)
	*(*uint8)(unsafe.Add(mBase, uint32(v242)+5)) = uint8(v267)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v264)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v242)+32)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v264)+40)) = v242
	goto L65
L65:
	;
	goto L1
}
