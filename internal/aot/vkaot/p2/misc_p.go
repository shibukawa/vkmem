package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_pad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v7 = m.G0
	v9 = v7 - int32(256)
	m.G0 = v9
	if l2 <= l3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(256)
	return
L2:
	;
	if l4&int32(73728) != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = l2 - l3
	v15 = int32(256)
	v17 = base.B2i32(base.Ui32(v14) < base.Ui32(v15))
	if base.Ui32(v14) < base.Ui32(v15) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v14
	goto L6
L5:
	;
	v18 = v15
	goto L6
L6:
	;
	v20 = F__emscripten_memset_bulkmem(m, v9, base.I32_extend8_s(l1), v18)
	goto L7
L7:
	;
	if base.Ui32(v14) < base.Ui32(v15) {
		v37 = v14
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_out(m, l0, v9, v37)
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L15
	}
L9:
	;
	v24 = v14
	goto L10
L10:
	;
	F_out(m, l0, v9, int32(256))
	v29 = m.ExcPending
	if v29 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v37 = v31
	goto L8
L12:
	;
	return
L13:
	;
	v31 = v24 + int32(-256)
	if base.Ui32(int32(255)) < base.Ui32(v31) {
		v24 = v31
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	goto L1
}
func F_panic(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = v9
		v14 = m.G3
		v15 = m.G397
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		v19 = F_fiprintf(m, v16, v14+int32(_a_F_panic_0), v5)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 + int32(16)
			return int32(0)
		}
	}
}
func F_parlist(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
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
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+74)) = uint8(v2)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v17 == int32(41) {
		v174 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+50)))
	v187 = v184 - v174&int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+73)) = uint8(v187)
	F_luaK_reserveregs(m, v13, v184)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L10
	} else {
		goto L31
	}
L2:
	;
	v25 = v17
	v26 = int32(0)
	goto L4
L3:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+50)))
	v76 = v75 + v73
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+50)) = uint8(v76)
	if v73 == int32(0) {
		v174 = v71
		goto L1
	} else {
		goto L21
	}
L4:
	;
	switch v25 + int32(-279) {
	case 0:
		goto L8
	default:
		goto L7
	case 6:
		goto L9
	}
L5:
	;
	v71 = int32(0)
	v73 = v62
	goto L3
L6:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+74)))
	if v63 != 0 {
		v71 = v63
		v73 = v62
		goto L3
	} else {
		goto L17
	}
L7:
	;
	v56 = m.G3
	F_luaX_syntaxerror(m, l0, v56+int32(_a_F_parlist_0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L16
	}
L8:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L13
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_luaX_next(m, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	F_new_localvar(m, l0, v34, v26)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v62 = v26 + int32(1)
	goto L6
L13:
	;
	v43 = m.G3
	v47 = F_luaX_newstring(m, l0, v43+int32(_a_F_parlist_1), int32(3))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	F_new_localvar(m, l0, v47, v26)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v51 = int32(7)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+74)) = uint8(v51)
	v71 = v51
	v73 = v26 + int32(1)
	goto L3
L16:
	;
	v62 = v26
	goto L6
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v64 != int32(44) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	goto L5
L19:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v25 = v69
	v26 = v62
	goto L4
L21:
	;
	v81 = v76 & int32(255)
	v83 = v74 + int32(172)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+24))
	v88 = v73 & int32(3)
	if v88 == int32(0) {
		v117 = v73
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if base.Ui32(v73) < base.Ui32(int32(4)) {
		v174 = v71
		goto L1
	} else {
		goto L27
	}
L23:
	;
	v92 = v73
	v102 = int32(0)
	goto L24
L24:
	;
	v104 = int32(1)
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83+(v81-v92)<<(uint(v104)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v86+v107*int32(12))+4)) = v84
	v113 = v92 + int32(-1)
	v115 = v102 + v104
	if v115 != v88 {
		v92 = v113
		v102 = v115
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v117 = v113
	goto L22
L26:
	;
	goto L25
L27:
	;
	v130 = v117
	goto L28
L28:
	;
	v144 = v83 + (v81-v130)<<(uint(int32(1))%32)
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144))))
	v146 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v86+v145*v146)+4)) = v84
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+int32(2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v86+v152*v146)+4)) = v84
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+int32(4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v86+v159*v146)+4)) = v84
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+int32(6)))))
	*(*int32)(unsafe.Add(mBase, uint32(v86+v166*v146)+4)) = v84
	v172 = v130 + int32(-4)
	if v172 != 0 {
		v130 = v172
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v174 = v71
	goto L1
L30:
	;
	goto L29
L31:
	;
	return
}
func F_parseExtendedExpireArgumentsOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
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
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
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
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
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
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l2 < int32(4) {
		v248 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v248
L2:
	;
	v18 = int32(0)
	v27 = int32(3)
	v28 = v18
	v29 = v18
	v30 = v18
	v31 = v18
	goto L5
L3:
	;
	v248 = int32(-1)
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v38
	F_addReplyErrorFormat(m, l0, int32(_a_F_parseExtendedExpireArgumentsOrReply_0), v13)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L67
	} else {
		goto L72
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v27<<(uint(int32(2))%32))))
	v38 = F_objectGetVal(m, v37)
	mBase = m.M
	v39 = int32(_a_F_parseExtendedExpireArgumentsOrReply_1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v214 = int32(0)
	if v206 == v214 {
		goto L64
	} else {
		goto L65
	}
L7:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v208 | v207
	v212 = v27 + int32(1)
	if v212 != l2 {
		v27 = v212
		v28 = v203
		v29 = v204
		v30 = v205
		v31 = v206
		goto L5
	} else {
		goto L63
	}
L8:
	;
	v80 = int32(_a_F_parseExtendedExpireArgumentsOrReply_2)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v83 != 0 {
		goto L25
	} else {
		goto L26
	}
L9:
	;
	if v74-v76 != 0 {
		goto L8
	} else {
		goto L21
	}
L10:
	;
	v74 = F_tolower(m, v70)
	mBase = m.M
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v76 = F_tolower(m, v75)
	mBase = m.M
	goto L9
L11:
	;
	v44 = v38
	v45 = v39
	v46 = v42
	goto L14
L12:
	;
	v70 = int32(0)
	v71 = v39
	goto L10
L13:
	;
	v70 = v67 & int32(255)
	v71 = v66
	goto L10
L14:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v48 == int32(0) {
		v66 = v45
		v67 = v46
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v66 = v60
	v67 = int32(0)
	goto L13
L16:
	;
	v52 = v46 & int32(255)
	if v52 == v48 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v59 = int32(1)
	v60 = v45 + v59
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v61 != 0 {
		v44 = v44 + v59
		v45 = v60
		v46 = v61
		goto L14
	} else {
		goto L20
	}
L18:
	;
	v54 = F_tolower(m, v52)
	mBase = m.M
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v56 = F_tolower(m, v55)
	mBase = m.M
	if v54 == v56 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v66 = v45
	v67 = v58
	goto L13
L20:
	;
	goto L15
L21:
	;
	v78 = int32(1)
	v203 = v28
	v204 = v29
	v205 = v30
	v206 = v78
	v207 = v78
	goto L7
L22:
	;
	v121 = int32(_a_F_parseExtendedExpireArgumentsOrReply_3)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v124 != 0 {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	if v115-v117 != 0 {
		goto L22
	} else {
		goto L35
	}
L24:
	;
	v115 = F_tolower(m, v111)
	mBase = m.M
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v117 = F_tolower(m, v116)
	mBase = m.M
	goto L23
L25:
	;
	v85 = v38
	v86 = v80
	v87 = v83
	goto L28
L26:
	;
	v111 = int32(0)
	v112 = v80
	goto L24
L27:
	;
	v111 = v108 & int32(255)
	v112 = v107
	goto L24
L28:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v89 == int32(0) {
		v107 = v86
		v108 = v87
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v107 = v101
	v108 = int32(0)
	goto L27
L30:
	;
	v93 = v87 & int32(255)
	if v93 == v89 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v100 = int32(1)
	v101 = v86 + v100
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v102 != 0 {
		v85 = v85 + v100
		v86 = v101
		v87 = v102
		goto L28
	} else {
		goto L34
	}
L32:
	;
	v95 = F_tolower(m, v93)
	mBase = m.M
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v97 = F_tolower(m, v96)
	mBase = m.M
	if v95 == v97 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v107 = v86
	v108 = v99
	goto L27
L34:
	;
	goto L29
L35:
	;
	v203 = v28
	v204 = v29
	v205 = int32(1)
	v206 = v31
	v207 = int32(2)
	goto L7
L36:
	;
	v162 = int32(_a_F_parseExtendedExpireArgumentsOrReply_4)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v165 != 0 {
		goto L52
	} else {
		goto L53
	}
L37:
	;
	if v156-v158 != 0 {
		goto L36
	} else {
		goto L49
	}
L38:
	;
	v156 = F_tolower(m, v152)
	mBase = m.M
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v158 = F_tolower(m, v157)
	mBase = m.M
	goto L37
L39:
	;
	v126 = v38
	v127 = v121
	v128 = v124
	goto L42
L40:
	;
	v152 = int32(0)
	v153 = v121
	goto L38
L41:
	;
	v152 = v149 & int32(255)
	v153 = v148
	goto L38
L42:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v130 == int32(0) {
		v148 = v127
		v149 = v128
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v148 = v142
	v149 = int32(0)
	goto L41
L44:
	;
	v134 = v128 & int32(255)
	if v134 == v130 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v141 = int32(1)
	v142 = v127 + v141
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	if v143 != 0 {
		v126 = v126 + v141
		v127 = v142
		v128 = v143
		goto L42
	} else {
		goto L48
	}
L46:
	;
	v136 = F_tolower(m, v134)
	mBase = m.M
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v138 = F_tolower(m, v137)
	mBase = m.M
	if v136 == v138 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v148 = v127
	v149 = v140
	goto L41
L48:
	;
	goto L43
L49:
	;
	v203 = v28
	v204 = int32(1)
	v205 = v30
	v206 = v31
	v207 = int32(4)
	goto L7
L50:
	;
	if v197-v199 != 0 {
		goto L4
	} else {
		goto L62
	}
L51:
	;
	v197 = F_tolower(m, v193)
	mBase = m.M
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	v199 = F_tolower(m, v198)
	mBase = m.M
	goto L50
L52:
	;
	v167 = v38
	v168 = v162
	v169 = v165
	goto L55
L53:
	;
	v193 = int32(0)
	v194 = v162
	goto L51
L54:
	;
	v193 = v190 & int32(255)
	v194 = v189
	goto L51
L55:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v171 == int32(0) {
		v189 = v168
		v190 = v169
		goto L54
	} else {
		goto L57
	}
L56:
	;
	v189 = v183
	v190 = int32(0)
	goto L54
L57:
	;
	v175 = v169 & int32(255)
	if v175 == v171 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v182 = int32(1)
	v183 = v168 + v182
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	if v184 != 0 {
		v167 = v167 + v182
		v168 = v183
		v169 = v184
		goto L55
	} else {
		goto L61
	}
L59:
	;
	v177 = F_tolower(m, v175)
	mBase = m.M
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	v179 = F_tolower(m, v178)
	mBase = m.M
	if v177 == v179 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	v189 = v168
	v190 = v181
	goto L54
L61:
	;
	goto L56
L62:
	;
	v203 = int32(1)
	v204 = v29
	v205 = v30
	v206 = v31
	v207 = int32(8)
	goto L7
L63:
	;
	goto L6
L64:
	;
	if v204 == int32(0) {
		v248 = v214
		goto L1
	} else {
		goto L69
	}
L65:
	;
	if v205|v204|v203 == int32(0) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	F_addReplyError(m, l0, int32(_a_F_parseExtendedExpireArgumentsOrReply_5))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	return int32(0)
L68:
	;
	goto L3
L69:
	;
	if v203 == int32(0) {
		v248 = v214
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_addReplyError(m, l0, int32(_a_F_parseExtendedExpireArgumentsOrReply_6))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L3
L72:
	;
	goto L3
}
func F_parseInlineBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
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
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = v14 + v15
	v17 = int32(10)
	v18 = F___strchrnul(m, v16, v17)
	mBase = m.M
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v20 == v17 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	if v24 != v16 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	if v24 != 0 {
		goto L2
	} else {
		goto L7
	}
L4:
	;
	v24 = v18
	goto L6
L5:
	;
	v24 = int32(0)
	goto L6
L6:
	;
	goto L3
L7:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-1)))))
	switch v28 & int32(7) {
	case 0:
		goto L13
	case 1:
		goto L12
	case 2:
		goto L11
	case 3:
		goto L10
	case 4:
		goto L9
	default:
		v45 = int32(0)
		goto L8
	}
L8:
	;
	if base.Ui32(v45-v15) < base.Ui32(int32(65537)) {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-17))))
	v45 = v44
	goto L8
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-9))))
	v45 = v41
	goto L8
L11:
	;
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(-5)))))
	v45 = v38
	goto L8
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-3)))))
	v45 = v35
	goto L8
L13:
	;
	v45 = int32(base.Ui32(v28) >> (uint(int32(3)) % 32))
	goto L8
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v13 | int32(2)
	goto L1
L15:
	;
	v65 = v64 - v16
	v68 = F_sdsnsplitargs(m, v16, v65, v11+int32(12))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L25
	} else {
		goto L26
	}
L16:
	;
	v55 = v24 + int32(-1)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v58 = base.B2i32(v56 == int32(13))
	if v56 == int32(13) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v63 = int32(1)
	v64 = v16
	goto L15
L18:
	;
	v59 = v55
	goto L20
L19:
	;
	v59 = v24
	goto L20
L20:
	;
	if v56 == int32(13) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v62 = int32(2)
	goto L23
L22:
	;
	v62 = int32(1)
	goto L23
L23:
	;
	v63 = v62
	v64 = v59
	goto L15
L24:
	;
	if v64 != v16 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	return
L26:
	;
	if v68 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v70 | int32(1024)
	goto L1
L28:
	;
	if v64 == v16 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v75 | int32(2048)
	goto L28
L30:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v65 + v63 + v92
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v95 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	if v13&int32(16384) == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_sdsfreesplitres(m, v68, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v87 | int32(512)
	goto L1
L34:
	;
	F_valkey_free(m, v68)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L25
	} else {
		goto L53
	}
L35:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v98 == int32(0) {
		v104 = v95
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v104
	v109 = F_valkey_malloc(m, v104<<(uint(int32(2))%32))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L25
	} else {
		goto L40
	}
L38:
	;
	F_valkey_free(m, v98)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L25
	} else {
		goto L39
	}
L39:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v104 = v103
	goto L37
L40:
	;
	v111 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v109
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v111
	if v114 < int32(1) {
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v125 = int32(0)
	goto L42
L42:
	;
	v129 = v68 + v125<<(uint(int32(2))%32)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v132 = F_sdsRemoveFreeSpace(m, v130, int32(1))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L25
	} else {
		goto L44
	}
L43:
	;
	goto L34
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v132
	v135 = int32(0)
	v137 = F_createObject(m, v135, v132)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v139+v140<<(uint(int32(2))%32)))) = v137
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v140 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+int32(-1)))))
	switch v151 & int32(7) {
	case 0:
		goto L51
	case 1:
		goto L50
	case 2:
		goto L49
	case 3:
		goto L48
	case 4:
		goto L47
	default:
		v168 = v135
		goto L46
	}
L46:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v169 + v168
	v173 = v125 + int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v173 < v174 {
		v125 = v173
		goto L42
	} else {
		goto L52
	}
L47:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-17))))
	v168 = v167
	goto L46
L48:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-9))))
	v168 = v164
	goto L46
L49:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148+int32(-5)))))
	v168 = v161
	goto L46
L50:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+int32(-3)))))
	v168 = v158
	goto L46
L51:
	;
	v168 = int32(base.Ui32(v151) >> (uint(int32(3)) % 32))
	goto L46
L52:
	;
	goto L43
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v188 | int32(8192)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+240)) = base.I64_extend_i32_u(v192 + v193 + int32(1))
	goto L1
}
func F_parseScanOptionsOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int64
	_ = v330
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v530 int32
	_ = v530
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int64)(unsafe.Add(mBase, uint32(l4)+32)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(l4)+16)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = int64(9223372036854775807)
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(10)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if l2 < v26 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v530
L2:
	;
	v31 = l2
	v35 = v26
	goto L6
L3:
	;
	v530 = int32(0)
	goto L1
L4:
	;
	v530 = int32(-1)
	goto L1
L5:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_parseScanOptionsOrReply[0]))
	F_addReplyErrorObject(m, l0, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L25
	} else {
		goto L159
	}
L6:
	;
	v40 = v35 - v31
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v43 = v31 << (uint(int32(2)) % 32)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v43)))
	v46 = F_objectGetVal(m, v45)
	mBase = m.M
	v47 = int32(_a_F_parseScanOptionsOrReply_0)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v530 = int32(0)
	goto L1
L8:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v508 < v514 {
		v31 = v508
		v35 = v514
		goto L6
	} else {
		goto L158
	}
L9:
	;
	v508 = v31 + int32(2)
	goto L8
L10:
	;
	v106 = int32(_a_F_parseScanOptionsOrReply_1)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v109 != 0 {
		goto L33
	} else {
		goto L34
	}
L11:
	;
	if v82-v84 != 0 {
		goto L10
	} else {
		goto L23
	}
L12:
	;
	v82 = F_tolower(m, v78)
	mBase = m.M
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v84 = F_tolower(m, v83)
	mBase = m.M
	goto L11
L13:
	;
	v52 = v46
	v53 = v47
	v54 = v50
	goto L16
L14:
	;
	v78 = int32(0)
	v79 = v47
	goto L12
L15:
	;
	v78 = v75 & int32(255)
	v79 = v74
	goto L12
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v56 == int32(0) {
		v74 = v53
		v75 = v54
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v74 = v68
	v75 = int32(0)
	goto L15
L18:
	;
	v60 = v54 & int32(255)
	if v60 == v56 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v67 = int32(1)
	v68 = v53 + v67
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v69 != 0 {
		v52 = v52 + v67
		v53 = v68
		v54 = v69
		goto L16
	} else {
		goto L22
	}
L20:
	;
	v62 = F_tolower(m, v60)
	mBase = m.M
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v64 = F_tolower(m, v63)
	mBase = m.M
	if v62 == v64 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v74 = v53
	v75 = v66
	goto L15
L22:
	;
	goto L17
L23:
	;
	if v40 < int32(2) {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v88 = int32(-1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v43+int32(4))))
	v95 = F_getLongFromObjectOrReply(m, l0, v93, l4, int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	if v95 != 0 {
		v530 = v88
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if int32(0) < v99 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_parseScanOptionsOrReply[0]))
	F_addReplyErrorObject(m, l0, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v530 = v88
	goto L1
L30:
	;
	v283 = int32(_a_F_parseScanOptionsOrReply_2)
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v286 != 0 {
		goto L84
	} else {
		goto L85
	}
L31:
	;
	if v141-v143 != 0 {
		goto L30
	} else {
		goto L43
	}
L32:
	;
	v141 = F_tolower(m, v137)
	mBase = m.M
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v143 = F_tolower(m, v142)
	mBase = m.M
	goto L31
L33:
	;
	v111 = v46
	v112 = v106
	v113 = v109
	goto L36
L34:
	;
	v137 = int32(0)
	v138 = v106
	goto L32
L35:
	;
	v137 = v134 & int32(255)
	v138 = v133
	goto L32
L36:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v115 == int32(0) {
		v133 = v112
		v134 = v113
		goto L35
	} else {
		goto L38
	}
L37:
	;
	v133 = v127
	v134 = int32(0)
	goto L35
L38:
	;
	v119 = v113 & int32(255)
	if v119 == v115 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v126 = int32(1)
	v127 = v112 + v126
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
	if v128 != 0 {
		v111 = v111 + v126
		v112 = v127
		v113 = v128
		goto L36
	} else {
		goto L42
	}
L40:
	;
	v121 = F_tolower(m, v119)
	mBase = m.M
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v123 = F_tolower(m, v122)
	mBase = m.M
	if v121 == v123 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v133 = v112
	v134 = v125
	goto L35
L42:
	;
	goto L37
L43:
	;
	if v40 < int32(2) {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147+v43+int32(4))))
	v152 = F_objectGetVal(m, v151)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v152
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+int32(-1)))))
	switch v156 & int32(7) {
	case 0:
		goto L51
	case 1:
		goto L50
	case 2:
		goto L49
	case 3:
		goto L48
	case 4:
		goto L47
	default:
		goto L52
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v187
	v189 = int32(-1)
	if v187 == int32(0) {
		v281 = v189
		goto L54
	} else {
		goto L55
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v177
	v179 = int32(1)
	if v177 != v179 {
		v186 = v177
		v187 = v179
		goto L45
	} else {
		goto L53
	}
L47:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v152+int32(-17))))
	v177 = v176
	goto L46
L48:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v152+int32(-9))))
	v177 = v173
	goto L46
L49:
	;
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152+int32(-5)))))
	v177 = v170
	goto L46
L50:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+int32(-3)))))
	v177 = v167
	goto L46
L51:
	;
	v177 = int32(base.Ui32(v156) >> (uint(int32(3)) % 32))
	goto L46
L52:
	;
	v159 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v159
	v186 = v159
	v187 = int32(1)
	goto L45
L53:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	v186 = int32(1)
	v187 = base.B2i32(v182 != int32(42))
	goto L45
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v281
	goto L9
L55:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_parseScanOptionsOrReply[1]))
	if v193 == int32(0) {
		v281 = v189
		goto L54
	} else {
		goto L56
	}
L56:
	;
	if v186 < int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v281 = v280
	goto L54
L58:
	;
	v269 = F_crc16(m, v152, v186)
	mBase = m.M
	v280 = v269 & int32(16383)
	goto L57
L59:
	;
	v204 = int32(-1)
	v209 = v204
	v210 = int32(0)
	goto L60
L60:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+v210))))
	v218 = v216 + int32(-63)
	if base.Ui32(int32(29)) < base.Ui32(v218) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	goto L58
L62:
	;
	v259 = v210 + int32(1)
	if v259 != v186 {
		v209 = v256
		v210 = v259
		goto L60
	} else {
		goto L81
	}
L63:
	;
	v280 = v254
	goto L57
L64:
	;
	if v216 == int32(42) {
		v254 = v204
		goto L63
	} else {
		goto L67
	}
L65:
	;
	if int32(1)<<(uint(v218)%32)&int32(805306369) != 0 {
		v254 = v204
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	if v209 != int32(-1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if int32(0) <= v209 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	if v216 == int32(123) {
		v256 = v210
		goto L62
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v237 = base.B2i32(v210 == v209+int32(1))
	if v210 == v209+int32(1) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v256 = v209
	goto L62
L73:
	;
	v238 = int32(-2)
	goto L75
L74:
	;
	v238 = v209
	goto L75
L75:
	;
	if v216 == int32(125) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v241 = v238
	goto L78
L77:
	;
	v241 = v209
	goto L78
L78:
	;
	if v216 != int32(125) {
		v256 = v241
		goto L62
	} else {
		goto L79
	}
L79:
	;
	if v210 == v209+int32(1) {
		v256 = v241
		goto L62
	} else {
		goto L80
	}
L80:
	;
	v250 = F_crc16(m, v152+v209+int32(1), v210+(v209^int32(-1)))
	mBase = m.M
	v254 = v250 & int32(16383)
	goto L63
L81:
	;
	goto L61
L82:
	;
	if l1 != 0 {
		goto L94
	} else {
		goto L95
	}
L83:
	;
	v318 = F_tolower(m, v314)
	mBase = m.M
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	v320 = F_tolower(m, v319)
	mBase = m.M
	goto L82
L84:
	;
	v288 = v46
	v289 = v283
	v290 = v286
	goto L87
L85:
	;
	v314 = int32(0)
	v315 = v283
	goto L83
L86:
	;
	v314 = v311 & int32(255)
	v315 = v310
	goto L83
L87:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	if v292 == int32(0) {
		v310 = v289
		v311 = v290
		goto L86
	} else {
		goto L89
	}
L88:
	;
	v310 = v304
	v311 = int32(0)
	goto L86
L89:
	;
	v296 = v290 & int32(255)
	if v296 == v292 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v303 = int32(1)
	v304 = v289 + v303
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+1)))
	if v305 != 0 {
		v288 = v288 + v303
		v289 = v304
		v290 = v305
		goto L87
	} else {
		goto L93
	}
L91:
	;
	v298 = F_tolower(m, v296)
	mBase = m.M
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	v300 = F_tolower(m, v299)
	mBase = m.M
	if v298 == v300 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	v310 = v289
	v311 = v302
	goto L86
L93:
	;
	goto L88
L94:
	;
	v339 = int32(_a_F_parseScanOptionsOrReply_3)
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v342 != 0 {
		goto L106
	} else {
		goto L107
	}
L95:
	;
	if v318-v320 != 0 {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	if v40 < int32(2) {
		goto L94
	} else {
		goto L97
	}
L97:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v324+v43+int32(4))))
	v329 = F_objectGetVal(m, v328)
	mBase = m.M
	v330 = F_getObjectTypeByName(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L25
	} else {
		goto L98
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = v330
	if v330 != int64(9223372036854775807) {
		goto L9
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v329
	F_addReplyErrorFormat(m, l0, int32(_a_F_parseScanOptionsOrReply_4), v14)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L25
	} else {
		goto L100
	}
L100:
	;
	goto L4
L101:
	;
	if l3 == int32(0) {
		goto L5
	} else {
		goto L138
	}
L102:
	;
	v437 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+24)) = v437
	v508 = v31 + v437
	goto L8
L103:
	;
	v388 = int32(_a_F_parseScanOptionsOrReply_5)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v391 != 0 {
		goto L123
	} else {
		goto L124
	}
L104:
	;
	if v374-v376 != 0 {
		goto L103
	} else {
		goto L116
	}
L105:
	;
	v374 = F_tolower(m, v370)
	mBase = m.M
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371))))
	v376 = F_tolower(m, v375)
	mBase = m.M
	goto L104
L106:
	;
	v344 = v46
	v345 = v339
	v346 = v342
	goto L109
L107:
	;
	v370 = int32(0)
	v371 = v339
	goto L105
L108:
	;
	v370 = v367 & int32(255)
	v371 = v366
	goto L105
L109:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	if v348 == int32(0) {
		v366 = v345
		v367 = v346
		goto L108
	} else {
		goto L111
	}
L110:
	;
	v366 = v360
	v367 = int32(0)
	goto L108
L111:
	;
	v352 = v346 & int32(255)
	if v352 == v348 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v359 = int32(1)
	v360 = v345 + v359
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+1)))
	if v361 != 0 {
		v344 = v344 + v359
		v345 = v360
		v346 = v361
		goto L109
	} else {
		goto L115
	}
L113:
	;
	v354 = F_tolower(m, v352)
	mBase = m.M
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	v356 = F_tolower(m, v355)
	mBase = m.M
	if v354 == v356 {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	v366 = v345
	v367 = v358
	goto L108
L115:
	;
	goto L110
L116:
	;
	if l1 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_addReplyError(m, l0, int32(_a_F_parseScanOptionsOrReply_6))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L25
	} else {
		goto L120
	}
L118:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v380&int32(15) == int32(4) {
		goto L102
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	goto L4
L121:
	;
	if v423-v425 != 0 {
		goto L101
	} else {
		goto L133
	}
L122:
	;
	v423 = F_tolower(m, v419)
	mBase = m.M
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v425 = F_tolower(m, v424)
	mBase = m.M
	goto L121
L123:
	;
	v393 = v46
	v394 = v388
	v395 = v391
	goto L126
L124:
	;
	v419 = int32(0)
	v420 = v388
	goto L122
L125:
	;
	v419 = v416 & int32(255)
	v420 = v415
	goto L122
L126:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	if v397 == int32(0) {
		v415 = v394
		v416 = v395
		goto L125
	} else {
		goto L128
	}
L127:
	;
	v415 = v409
	v416 = int32(0)
	goto L125
L128:
	;
	v401 = v395 & int32(255)
	if v401 == v397 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v408 = int32(1)
	v409 = v394 + v408
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+1)))
	if v410 != 0 {
		v393 = v393 + v408
		v394 = v409
		v395 = v410
		goto L126
	} else {
		goto L132
	}
L130:
	;
	v403 = F_tolower(m, v401)
	mBase = m.M
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	v405 = F_tolower(m, v404)
	mBase = m.M
	if v403 == v405 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	v415 = v394
	v416 = v407
	goto L125
L132:
	;
	goto L127
L133:
	;
	if l1 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	F_addReplyError(m, l0, int32(_a_F_parseScanOptionsOrReply_7))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L25
	} else {
		goto L137
	}
L135:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v429&int32(15) == int32(3) {
		goto L102
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	goto L4
L138:
	;
	v443 = int32(_a_F_parseScanOptionsOrReply_8)
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v446 != 0 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	if v478-v480 != 0 {
		goto L5
	} else {
		goto L151
	}
L140:
	;
	v478 = F_tolower(m, v474)
	mBase = m.M
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
	v480 = F_tolower(m, v479)
	mBase = m.M
	goto L139
L141:
	;
	v448 = v46
	v449 = v443
	v450 = v446
	goto L144
L142:
	;
	v474 = int32(0)
	v475 = v443
	goto L140
L143:
	;
	v474 = v471 & int32(255)
	v475 = v470
	goto L140
L144:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
	if v452 == int32(0) {
		v470 = v449
		v471 = v450
		goto L143
	} else {
		goto L146
	}
L145:
	;
	v470 = v464
	v471 = int32(0)
	goto L143
L146:
	;
	v456 = v450 & int32(255)
	if v456 == v452 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v463 = int32(1)
	v464 = v449 + v463
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+1)))
	if v465 != 0 {
		v448 = v448 + v463
		v449 = v464
		v450 = v465
		goto L144
	} else {
		goto L150
	}
L148:
	;
	v458 = F_tolower(m, v456)
	mBase = m.M
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
	v460 = F_tolower(m, v459)
	mBase = m.M
	if v458 == v460 {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	v470 = v449
	v471 = v462
	goto L143
L150:
	;
	goto L145
L151:
	;
	if v40 < int32(2) {
		goto L5
	} else {
		goto L152
	}
L152:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	if v484 == int32(-1) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v490+v43+int32(4))))
	v495 = F_getSlotOrReply(m, l0, v494)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L25
	} else {
		goto L156
	}
L154:
	;
	F_addReplyError(m, l0, int32(_a_F_parseScanOptionsOrReply_9))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L25
	} else {
		goto L155
	}
L155:
	;
	goto L4
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+28)) = v495
	v498 = int32(-1)
	if v495 == v498 {
		v530 = v498
		goto L1
	} else {
		goto L157
	}
L157:
	;
	goto L9
L158:
	;
	goto L7
L159:
	;
	goto L4
}
func F_patchlistaux(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	if l1 == int32(-1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = l1
	goto L3
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v37 = v34 + v22<<(uint(int32(2))%32)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v22 < int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L1
L5:
	;
	v58 = int32(base.Ui32(v38)>>(uint(int32(14))%32)) + int32(-131071)
	if v56&int32(63) != int32(27) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v55 = v37
	v56 = v38
	goto L5
L7:
	;
	v44 = v37 + int32(-4)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = m.G400
	v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(v46+v45&int32(63)))))
	if v50 < int32(0) {
		v55 = v44
		v56 = v45
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v111<<(uint(int32(14))%32) | v109&int32(16383) + int32(2147467264)
	if v58 == int32(-1) {
		goto L1
	} else {
		goto L22
	}
L10:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v109 = v108
	v111 = v107
	goto L9
L11:
	;
	v99 = m.G3
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v100, v99+int32(_a_F_patchlistaux_0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	v89 = l4 + (v22 ^ int32(-1))
	v91 = v89 >> (uint(int32(31)) % 32)
	if base.Ui32(v89^v91-v91) < base.Ui32(int32(131072)) {
		v109 = v38
		v111 = v89
		goto L9
	} else {
		goto L19
	}
L13:
	;
	if l3 == int32(255) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v76
	v80 = l2 + (v22 ^ int32(-1))
	v82 = v80 >> (uint(int32(31)) % 32)
	if base.Ui32(int32(131071)) < base.Ui32(v80^v82-v82) {
		v98 = v80
		goto L11
	} else {
		goto L18
	}
L15:
	;
	v76 = int32(base.Ui32(v56)>>(uint(int32(17))%32))&int32(32704) | v56&int32(8372250)
	goto L14
L16:
	;
	if l3 == int32(base.Ui32(v56)>>(uint(int32(23))%32)) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v76 = v56&int32(-16357) | l3<<(uint(int32(6))%32)&int32(16320)
	goto L14
L18:
	;
	v107 = v80
	goto L10
L19:
	;
	v98 = v89
	goto L11
L20:
	;
	return
L21:
	;
	v107 = v98
	goto L10
L22:
	;
	v124 = v22 + v58 + int32(1)
	if v124 != int32(-1) {
		v22 = v124
		goto L3
	} else {
		goto L23
	}
L23:
	;
	goto L4
}
func F_patternHashSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
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
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
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
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	if l1 < int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if int32(1) <= l1 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	v11 = int32(-1)
	v16 = v11
	v17 = int32(0)
	goto L3
L3:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v17))))
	v25 = v23 + int32(-63)
	if base.Ui32(int32(29)) < base.Ui32(v25) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L1
L5:
	;
	v151 = v17 + int32(1)
	if v151 != l1 {
		v16 = v148
		v17 = v151
		goto L3
	} else {
		goto L35
	}
L6:
	;
	return v145
L7:
	;
	if v23 == int32(42) {
		v145 = v11
		goto L6
	} else {
		goto L10
	}
L8:
	;
	if int32(1)<<(uint(v25)%32)&int32(805306369) != 0 {
		v145 = v11
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	if v16 != int32(-1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if int32(0) <= v16 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v23 == int32(123) {
		v148 = v17
		goto L5
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v44 = base.B2i32(v17 == v16+int32(1))
	if v17 == v16+int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v148 = v16
	goto L5
L16:
	;
	v45 = int32(-2)
	goto L18
L17:
	;
	v45 = v16
	goto L18
L18:
	;
	if v23 == int32(125) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v48 = v45
	goto L21
L20:
	;
	v48 = v16
	goto L21
L21:
	;
	if v23 != int32(125) {
		v148 = v48
		goto L5
	} else {
		goto L22
	}
L22:
	;
	if v17 == v16+int32(1) {
		v148 = v48
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v52 = int32(1)
	v53 = l0 + v16 + v52
	v56 = v17 + (v16 ^ int32(-1))
	if v52 <= v56 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v145 = v135 & int32(65535) & int32(16383)
	goto L6
L25:
	;
	goto L24
L26:
	;
	v64 = int32(1)
	if v56 != v64 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v135 = int32(0)
	goto L25
L28:
	;
	if v56&v64 == int32(0) {
		v135 = v113
		goto L25
	} else {
		goto L34
	}
L29:
	;
	v71 = int32(0)
	v73 = v53
	v74 = v71
	v77 = v71
	goto L31
L30:
	;
	v112 = v53
	v113 = int32(0)
	goto L28
L31:
	;
	v79 = int32(65280)
	v81 = int32(8)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v85 = int32(1)
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32(v74&v79)>>(uint(v81)%32))^v83)<<(uint(v85)%32))+uint32(_c_F_patternHashSlot[0]))))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32((v89^v74<<(uint(v81)%32))&v79)>>(uint(v81)%32))^v97)<<(uint(v85)%32))+uint32(_c_F_patternHashSlot[0]))))
	v106 = v103 ^ v89<<(uint(v81)%32)
	v107 = int32(2)
	v108 = v73 + v107
	v110 = v77 + v107
	if v110 != v56&int32(2147483646) {
		v73 = v108
		v74 = v106
		v77 = v110
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v112 = v108
	v113 = v106
	goto L28
L33:
	;
	goto L32
L34:
	;
	v122 = int32(8)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32(v113&int32(65280))>>(uint(v122)%32))^v124)<<(uint(int32(1))%32))+uint32(_c_F_patternHashSlot[0]))))
	v135 = v130 ^ v113<<(uint(v122)%32)
	goto L25
L35:
	;
	goto L4
L36:
	;
	return v239 & int32(65535) & int32(16383)
L37:
	;
	goto L36
L38:
	;
	v168 = int32(1)
	if l1 != v168 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v239 = int32(0)
	goto L37
L40:
	;
	if l1&v168 == int32(0) {
		v239 = v217
		goto L37
	} else {
		goto L46
	}
L41:
	;
	v175 = int32(0)
	v177 = l0
	v178 = v175
	v181 = v175
	goto L43
L42:
	;
	v216 = l0
	v217 = int32(0)
	goto L40
L43:
	;
	v183 = int32(65280)
	v185 = int32(8)
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v189 = int32(1)
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32(v178&v183)>>(uint(v185)%32))^v187)<<(uint(v189)%32))+uint32(_c_F_patternHashSlot[0]))))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)))
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32((v193^v178<<(uint(v185)%32))&v183)>>(uint(v185)%32))^v201)<<(uint(v189)%32))+uint32(_c_F_patternHashSlot[0]))))
	v210 = v207 ^ v193<<(uint(v185)%32)
	v211 = int32(2)
	v212 = v177 + v211
	v214 = v181 + v211
	if v214 != l1&int32(2147483646) {
		v177 = v212
		v178 = v210
		v181 = v214
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v216 = v212
	v217 = v210
	goto L40
L45:
	;
	goto L44
L46:
	;
	v226 = int32(8)
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32(v217&int32(65280))>>(uint(v226)%32))^v228)<<(uint(int32(1))%32))+uint32(_c_F_patternHashSlot[0]))))
	v239 = v234 ^ v217<<(uint(v226)%32)
	goto L37
}
func F_pauseActions(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v6 = l0 << (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_pauseActions[0]))) = l2
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_pauseActions[1])))
	if l1 <= v9 {
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_pauseActions[1]))) = l1
	}
	F_updatePausedActions(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_pauseActions[2]))
		if v17 == int32(0) {
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_pauseActions[3])) = int32(1)
		}
		return
	}
}
func F_performModuleConfigSetDefaultFromName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v27 int32
	_ = v27
	var v28 int64
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
	var v54 int32
	_ = v54
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_performModuleConfigSetDefaultFromName[0]))
	v5 = F_dictFind(m, v4, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			F__serverAssert(m, int32(_a_F_performModuleConfigSetDefaultFromName_0), int32(_a_F_performModuleConfigSetDefaultFromName_1), int32(747))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			if v11 == int32(0) {
				F__serverAssert(m, int32(_a_F_performModuleConfigSetDefaultFromName_0), int32(_a_F_performModuleConfigSetDefaultFromName_1), int32(747))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+9)))
				if v14&int32(1) != 0 {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
					switch v21 {
					case 0:
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
						v46 = F_setModuleBoolConfig(m, v44, v45, l1)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							return v46
						}
					case 1:
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
						v28 = *(*int64)(unsafe.Add(mBase, uint32(v11)+64))
						v29 = F_setModuleNumericConfig(m, v27, v28, l1)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							return v29
						}
					default:
						F__serverPanic_1(m, int32(_a_F_performModuleConfigSetDefaultFromName_1), int32(757), int32(_a_F_performModuleConfigSetDefaultFromName_2), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					case 3:
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
						v24 = F_setModuleStringConfig(m, v22, v23, l1)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return v24
						}
					case 4:
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
						v34 = F_setModuleEnumConfig(m, v32, v33, l1)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							return v34
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a_F_performModuleConfigSetDefaultFromName_3)
					return int32(0)
				}
			}
		}
	}
}
func F_performSlotImportJobFailover(m *base.Module, l0 int32) {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_performSlotImportJobFailover_0), int32(_a_F_performSlotImportJobFailover_1), int32(858))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L30
	}
L2:
	;
	v12 = F_clusterBumpConfigEpochWithoutConsensus(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v16 = v7 + int32(8)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v17
	goto L5
L5:
	;
	v22 = v7 + int32(8)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v24 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L3
	} else {
		goto L26
	}
L7:
	;
	if v24 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24+base.B2i32(v27 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v33
	goto L8
L10:
	;
	v37 = v24
	goto L11
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v43 < v42 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L6
L13:
	;
	v96 = v7 + int32(8)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	if v98 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L14:
	;
	v45 = v42
	goto L15
L15:
	;
	v49 = F_clusterDelSlot(m, v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L17
	}
L16:
	;
	goto L13
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_performSlotImportJobFailover[0]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v60 = v45 << (uint(int32(2)) % 32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v52+v60+int32(52))))
	if v64 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v45 < v87 {
		v45 = v45 + int32(1)
		goto L15
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	F_clusterNodeSetSlotBit(m, v53, v45)
	mBase = m.M
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_performSlotImportJobFailover[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v67+v60+int32(52)))) = v53
	v73 = base.I32_div_s(v45, int32(8))
	v74 = v67 + v73
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_performSlotImportJobFailover[1]))))
	v82 = v77 & base.I32_rotl(int32(-2), v45&int32(7))
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_performSlotImportJobFailover[1]))) = uint8(v82)
	F_clusterSlotStatReset(m, v45)
	mBase = m.M
	goto L19
L21:
	;
	goto L16
L22:
	;
	if v98 != 0 {
		v37 = v98
		goto L11
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v98+base.B2i32(v101 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107
	goto L23
L25:
	;
	goto L12
L26:
	;
	F_clusterUpdateState(m)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	F_clusterDoBeforeSleep(m, int32(12))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	F_clusterDoBeforeSleep(m, int32(32))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	m.G0 = v7 + int32(16)
	return
L30:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pexpireCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int64)(unsafe.Add(mBase, _c_F_pexpireCommand[0]))
	F_expireGenericCommand(m, l0, v3, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_pipe(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v2 = m.Env.X__syscall_pipe(m, l0)
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
func F_pntz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_a_ctz_32(m, v3+int32(-1))
	mBase = m.M
	if v6 != 0 {
		v14 = v6
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = F_a_ctz_32(m, v7)
		mBase = m.M
		if v8 != 0 {
			v12 = v8 | int32(32)
		} else {
			v12 = int32(0)
		}
		v14 = v12
	}
	return v14
}
func F_pop_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 float64
	_ = v127
	var v130 int32
	_ = v130
	switch l1 + int32(-9) {
	case 0:
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v7 + int32(4)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11
		return
	case 1:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v13 + int32(4)
		v17 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v17
		return
	case 2:
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19 + int32(4)
		v23 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v19))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v23
		return
	case 3:
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v41 = (v37 + int32(7)) & int32(-8)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v41 + int32(8)
		v45 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v45
		return
	case 4:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v25 + int32(4)
		v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v29
		return
	case 5:
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31 + int32(4)
		v35 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v35
		return
	case 6:
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v47 + int32(4)
		v51 = int64(*(*int16)(unsafe.Add(mBase, uint32(v47))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v51
		return
	case 7:
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53 + int32(4)
		v57 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v53))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v57
		return
	case 8:
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v59 + int32(4)
		v63 = int64(*(*int8)(unsafe.Add(mBase, uint32(v59))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v63
		return
	case 9:
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + int32(4)
		v69 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v69
		return
	case 10:
		v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v75 = (v71 + int32(7)) & int32(-8)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v75 + int32(8)
		v79 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v79
		return
	case 11:
		v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v81 + int32(4)
		v85 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v81))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v85
		return
	case 12:
		v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v91 = (v87 + int32(7)) & int32(-8)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v91 + int32(8)
		v95 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v95
		return
	case 13:
		v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v101 = (v97 + int32(7)) & int32(-8)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v101 + int32(8)
		v105 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v105
		return
	case 14:
		v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v107 + int32(4)
		v111 = int64(*(*int32)(unsafe.Add(mBase, uint32(v107))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v111
		return
	case 15:
		v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v113 + int32(4)
		v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v113))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v117
		return
	case 16:
		v119 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v123 = (v119 + int32(7)) & int32(-8)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v123 + int32(8)
		v127 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
		*(*float64)(unsafe.Add(mBase, uint32(l0))) = v127
		return
	case 17:
		m.T0[l3].(func(*base.Module, int32, int32))(m, l0, l2)
		mBase = m.M
		v130 = m.ExcPending
		if v130 != 0 {
			return
		} else {
			return
		}
	default:
		return
	}
}
func F_popcountScalar(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int64
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int64
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int64
	_ = v261
	var v263 int32
	_ = v263
	var v266 int64
	_ = v266
	var v268 int32
	_ = v268
	var v271 int64
	_ = v271
	var v273 int32
	_ = v273
	var v276 int64
	_ = v276
	var v277 int64
	_ = v277
	var v281 int32
	_ = v281
	var v284 int64
	_ = v284
	if l0&int32(3) == int32(0) {
		v54 = l1
		v55 = int64(0)
		v56 = l0
	} else {
		if l1 != 0 {
			v14 = l1 + int32(-1)
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			v18 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_popcountScalar[0]))))
			v20 = l0 + int32(1)
			if v20&int32(3) != 0 {
				if v14 != 0 {
					v24 = l1 + int32(-2)
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
					v28 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_popcountScalar[0]))))
					v29 = v18 + v28
					v31 = l0 + int32(2)
					if v31&int32(3) != 0 {
						if v24 != 0 {
							v35 = l1 + int32(-3)
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
							v39 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v36)+uint32(_c_F_popcountScalar[0]))))
							v40 = v29 + v39
							v41 = int32(3)
							v42 = l0 + v41
							if v42&v41 != 0 {
								if v35 != 0 {
									v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
									v52 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+uint32(_c_F_popcountScalar[0]))))
									v54 = l1 + int32(-4)
									v55 = v40 + v52
									v56 = l0 + int32(4)
								} else {
									v54 = v35
									v55 = v40
									v56 = v42
								}
							} else {
								v54 = v35
								v55 = v40
								v56 = v42
							}
						} else {
							v54 = v24
							v55 = v29
							v56 = v31
						}
					} else {
						v54 = v24
						v55 = v29
						v56 = v31
					}
				} else {
					v54 = v14
					v55 = v18
					v56 = v20
				}
			} else {
				v54 = v14
				v55 = v18
				v56 = v20
			}
		} else {
			v54 = l1
			v55 = int64(0)
			v56 = l0
		}
	}
	if int32(27) < v54 {
		v61 = v54
		v62 = v55
		v63 = v56
		for {
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v68 = int32(1)
			v70 = int32(1431655765)
			v72 = v67 - int32(base.Ui32(v67)>>(uint(v68)%32))&v70
			v73 = int32(2)
			v75 = int32(858993459)
			v79 = int32(base.Ui32(v72)>>(uint(v73)%32))&v75 + v72&v75
			v80 = int32(4)
			v83 = int32(252645135)
			v85 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v90 = v85 - int32(base.Ui32(v85)>>(uint(v68)%32))&v70
			v97 = int32(base.Ui32(v90)>>(uint(v73)%32))&v75 + v90&v75
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
			v109 = v104 - int32(base.Ui32(v104)>>(uint(v68)%32))&v70
			v116 = int32(base.Ui32(v109)>>(uint(v73)%32))&v75 + v109&v75
			v123 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
			v128 = v123 - int32(base.Ui32(v123)>>(uint(v68)%32))&v70
			v135 = int32(base.Ui32(v128)>>(uint(v73)%32))&v75 + v128&v75
			v142 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
			v147 = v142 - int32(base.Ui32(v142)>>(uint(v68)%32))&v70
			v154 = int32(base.Ui32(v147)>>(uint(v73)%32))&v75 + v147&v75
			v161 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
			v166 = v161 - int32(base.Ui32(v161)>>(uint(v68)%32))&v70
			v173 = int32(base.Ui32(v166)>>(uint(v73)%32))&v75 + v166&v75
			v180 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
			v185 = v180 - int32(base.Ui32(v180)>>(uint(v68)%32))&v70
			v192 = int32(base.Ui32(v185)>>(uint(v73)%32))&v75 + v185&v75
			v204 = v62 + base.I64_extend_i32_u(int32(base.Ui32(((int32(base.Ui32(v79)>>(uint(v80)%32))+v79)&v83+(int32(base.Ui32(v97)>>(uint(v80)%32))+v97)&v83+(int32(base.Ui32(v116)>>(uint(v80)%32))+v116)&v83+(int32(base.Ui32(v135)>>(uint(v80)%32))+v135)&v83+(int32(base.Ui32(v154)>>(uint(v80)%32))+v154)&v83+(int32(base.Ui32(v173)>>(uint(v80)%32))+v173)&v83+(int32(base.Ui32(v192)>>(uint(v80)%32))+v192)&v83)*int32(16843009))>>(uint(int32(24))%32)))
			v206 = v63 + int32(28)
			v210 = v61 + int32(-28)
			if base.Ui32(int32(55)) < base.Ui32(v61) {
				v61 = v210
				v62 = v204
				v63 = v206
				continue
			} else {
				break
			}
			break
		}
		v213 = v204
		v214 = v206
		v216 = v210
	} else {
		v213 = v55
		v214 = v56
		v216 = v54
	}
	if v216 == int32(0) {
		v284 = v213
	} else {
		v221 = v216 & int32(3)
		if v221 != 0 {
			v223 = v214
			v224 = v216
			v225 = v213
			v227 = int32(0)
			for {
				v231 = v224 + int32(-1)
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
				v235 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v232)+uint32(_c_F_popcountScalar[0]))))
				v236 = v225 + v235
				v237 = int32(1)
				v238 = v223 + v237
				v240 = v227 + v237
				if v240 != v221 {
					v223 = v238
					v224 = v231
					v225 = v236
					v227 = v240
					continue
				} else {
					break
				}
				break
			}
			v243 = v231
			v244 = v236
			v245 = v238
		} else {
			v243 = v216
			v244 = v213
			v245 = v214
		}
		if base.Ui32(v216) < base.Ui32(int32(4)) {
			v284 = v244
		} else {
			v252 = v243
			v253 = v244
			v254 = v245
			for {
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
				v261 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v258)+uint32(_c_F_popcountScalar[0]))))
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+1)))
				v266 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_popcountScalar[0]))))
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+2)))
				v271 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v268)+uint32(_c_F_popcountScalar[0]))))
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+3)))
				v276 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v273)+uint32(_c_F_popcountScalar[0]))))
				v277 = v253 + v261 + v266 + v271 + v276
				v281 = v252 + int32(-4)
				if v281 != 0 {
					v252 = v281
					v253 = v277
					v254 = v254 + int32(4)
					continue
				} else {
					break
				}
				break
			}
			v284 = v277
		}
	}
	return v284
}
func F_populateArgsStructure(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 == v6 {
		v27 = v6
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return int32(0)
L3:
	;
	F__serverAssert(m, int32(_a_F_populateArgsStructure_0), int32(_a_F_populateArgsStructure_1), int32(1596))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L12
	}
L4:
	;
	return v27
L5:
	;
	v10 = l0
	v11 = v6
	goto L6
L6:
	;
	if v11 == int32(2147483647) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v27 = v22
	goto L4
L8:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v16 = F_populateArgsStructure(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v16
	v22 = v11 + int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	if v23 != 0 {
		v10 = v10 + int32(44)
		v11 = v22
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_populateCommandStructure(m *base.Module, l0 int32) int32 {
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
	var v16 int64
	_ = v16
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v70 int32
	_ = v70
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_populateCommandStructure[0]))
	v15 = int32(-1)
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if v16&int64(131072) != int64(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_populateCommandStructure_0), int32(_a_F_populateCommandStructure_1), int32(3309))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L28
	}
L2:
	;
	m.G0 = v11 + int32(16)
	return v94
L3:
	;
	if v16&int64(262144) == int64(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v14 != 0 {
		v94 = v15
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+212)) = int64(0)
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v29
	F_populateCommandLegacyRangeSpec(m, l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v14 == int32(0) {
		v94 = v15
		goto L2
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v37 = F_ACLGetCommandID(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v37
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v40 == int32(0) {
		v94 = v29
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v43 == int32(0) {
		v94 = v29
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v49 = v40
	v50 = int32(0)
	v52 = v43
	goto L14
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = F_sdsempty(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v94 = int32(0)
	goto L2
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v55
	v61 = F_sdscatfmt(m, v56, int32(_a_F_populateCommandStructure_2), v11)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+140)) = v61
	v64 = F_populateCommandStructure(m, v49)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v85 = v50 + int32(1)
	v88 = v83 + v85*int32(224)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v89 != 0 {
		v49 = v88
		v50 = v85
		v52 = v89
		goto L14
	} else {
		goto L27
	}
L19:
	;
	if v64 == int32(-1) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v68 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+204)) = l0
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v49)+140))
	v75 = F_ACLGetCommandID(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	v70 = F_hashtableCreate(m, int32(_a_F_populateCommandStructure_3))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v70
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+136)) = v75
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v79 = F_hashtableAdd(m, v78, v49)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	if v79 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	goto L15
L28:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_postExecutionUnitOperations(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
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
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_postExecutionUnitOperations[0]))
	if v6 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_postExecutionUnitOperations_0), int32(_a_F_postExecutionUnitOperations_1), int32(3728))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L39
	}
L2:
	;
	return
L3:
	;
	F_firePostExecutionUnitJobs(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v9 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_postExecutionUnitOperations[1]))
	if v10 == v9 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_modulePostExecutionUnitOperations(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L38
	}
L7:
	;
	v13 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_postExecutionUnitOperations[2]))
	if v15 == v13 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	if v44 < int32(1) {
		v73 = v44
		goto L17
	} else {
		goto L18
	}
L9:
	;
	v34 = int32(1)
	v35 = int32(-1)
	F_propagateNow(m, v35, int32(_a_F_postExecutionUnitOperations_2), v34, int32(3), v35)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L16
	}
L10:
	;
	if v10 < int32(2) {
		v44 = v10
		v45 = v13
		goto L8
	} else {
		goto L15
	}
L11:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	if v18 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v21 = int32(0)
	if v10 < int32(2) {
		v44 = v10
		v45 = v21
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	if v24&int64(268435456) == int64(0) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v44 = v10
	v45 = v21
	goto L8
L15:
	;
	goto L9
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_postExecutionUnitOperations[1]))
	v44 = v43
	v45 = v34
	goto L8
L17:
	;
	if v45 == int32(0) {
		v88 = v73
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v52 = int32(0)
	goto L19
L19:
	;
	v54 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_postExecutionUnitOperations[3]))
	v58 = v55 + v52*int32(20)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	if v59 == v54 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v73 = v71
	goto L17
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	F_propagateNow(m, v62, v63, v64, v59, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v69 = v52 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_postExecutionUnitOperations[1]))
	if v69 < v71 {
		v52 = v69
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	if v88 == int32(0) {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	v79 = int32(-1)
	F_propagateNow(m, v79, int32(_a_F_postExecutionUnitOperations_3), int32(1), int32(3), v79)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_postExecutionUnitOperations[1]))
	v88 = v87
	goto L24
L27:
	;
	v91 = v88
	goto L28
L28:
	;
	v95 = int32(0)
	v97 = v91 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_postExecutionUnitOperations[1])) = v97
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_postExecutionUnitOperations[3]))
	v104 = v101 + v97*int32(20)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v105 < int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L6
L30:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	F_valkey_free(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L36
	}
L31:
	;
	v108 = v95
	goto L32
L32:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v108<<(uint(int32(2))%32))))
	F_decrRefCount(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L34
	}
L33:
	;
	goto L30
L34:
	;
	v120 = v108 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v120 < v121 {
		v108 = v120
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_postExecutionUnitOperations[1]))
	if v131 != 0 {
		v91 = v131
		goto L28
	} else {
		goto L37
	}
L37:
	;
	goto L29
L38:
	;
	goto L2
L39:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_postWriteToClient(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int64
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
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
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
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
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int64
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = int64(0)
	v14 = int32(_a_F_postWriteToClient_0)
	v16 = *(*int64)(unsafe.Add(mBase, _c_F_postWriteToClient[0]))
	*(*int64)(unsafe.Add(mBase, _c_F_postWriteToClient[0])) = v16 + int64(1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v20&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L1:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_postWriteToClient_1), int32(_a_F_postWriteToClient_2), int32(2277))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L27
	} else {
		goto L102
	}
L2:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_postWriteToClient_1), int32(_a_F_postWriteToClient_2), int32(2277))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L27
	} else {
		goto L101
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_postWriteToClient_3), int32(_a_F_postWriteToClient_2), int32(3056))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L27
	} else {
		goto L100
	}
L4:
	;
	F__serverAssert(m, int32(_a_F_postWriteToClient_4), int32(_a_F_postWriteToClient_2), int32(3023))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L27
	} else {
		goto L99
	}
L5:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+220)))
	if v194&int32(1) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L6:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v83 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L7:
	;
	v75 = int32(_a_F_postWriteToClient_0)
	v77 = *(*int64)(unsafe.Add(mBase, _c_F_postWriteToClient[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_postWriteToClient[1])) = v77 + base.I64_extend_i32_s(v74)
	goto L6
L8:
	;
	F_postWriteToReplica(m, l0)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L27
	} else {
		goto L28
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v40 < int32(1) {
		goto L5
	} else {
		goto L17
	}
L10:
	;
	if v20&int32(2) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v20&int32(262144) != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	if v20&int32(4) == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L16
L16:
	;
	goto L9
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v43&int32(1) != 0 {
		v74 = v40
		goto L7
	} else {
		goto L18
	}
L18:
	;
	if v43&int32(2) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v56 == int32(0) {
		v74 = v40
		goto L7
	} else {
		goto L24
	}
L20:
	;
	if v43&int32(262144) != 0 {
		v74 = v40
		goto L7
	} else {
		goto L23
	}
L21:
	;
	if v43&int32(262148) == int32(4) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v74 = v40
	goto L7
L23:
	;
	goto L19
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	goto L25
L25:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v59 == int32(1) {
		v74 = v62
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v63 = int32(_a_F_postWriteToClient_0)
	v65 = *(*int64)(unsafe.Add(mBase, _c_F_postWriteToClient[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_postWriteToClient[2])) = v65 + base.I64_extend_i32_s(v62)
	goto L6
L27:
	;
	return int32(0)
L28:
	;
	goto L5
L29:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v122 = v10 + int32(8)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v123
	goto L43
L30:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v88 = base.B2i32(v86 != v87)
	if v86 != v87 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v86 != v87 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_postWriteToClient[3]))
	if v90 == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if base.Ui32(v83) < base.Ui32(v93) {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(0)
	goto L5
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v97&int32(16777216) == int32(0) {
		v105 = v97
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v83 != v95 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+180)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v105 & int32(-16777217)
	if v86 != v87 {
		goto L29
	} else {
		goto L42
	}
L40:
	;
	F_releaseBufReferences(m, v86, v83, l0)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L27
	} else {
		goto L41
	}
L41:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v105 = v104
	goto L39
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = int64(0)
	goto L5
L43:
	;
	goto L44
L44:
	;
	v135 = v10 + int32(8)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v137 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = int64(0)
	goto L5
L46:
	;
	if v137 == int32(0) {
		goto L5
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v137+base.B2i32(v140 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v146
	goto L47
L49:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	v152 = v150 + int32(13)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v154 = base.B2i32(v152 != v153)
	if v152 != v153 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v152 != v153 {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_postWriteToClient[3]))
	if v156 == int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v159 == int32(0) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	if base.Ui32(v162) < base.Ui32(v159) {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v168 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v169 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v150))))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v168 - v169
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+12)))
	if v172&int32(1) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v165 != v166 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_listDelNode(m, v180, v137)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L27
	} else {
		goto L61
	}
L59:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	F_releaseBufReferences(m, v152, v177, l0)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L27
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	if v152 != v153 {
		goto L44
	} else {
		goto L62
	}
L62:
	;
	goto L45
L63:
	;
	m.G0 = v10 + int32(16)
	return v303
L64:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v236 < int32(1) {
		goto L77
	} else {
		goto L78
	}
L65:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v200 == int32(3) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_postWriteToClient[4]))
	if int32(1) < v204 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v216 = int32(-1)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v217&int32(1280) != 0 {
		v303 = v216
		goto L63
	} else {
		goto L71
	}
L68:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+88))
	v209 = m.T0[v208].(func(*base.Module, int32) int32)(m, v199)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L27
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v209
	F__serverLog(m, int32(1), int32(_a_F_postWriteToClient_5), v10)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L27
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(1024)
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_postWriteToClient[3]))
	if v224 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_postWriteToClient[5]))
	v233 = F_listAddNodeTail(m, v232, l0)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L27
	} else {
		goto L76
	}
L73:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_postWriteToClient[5]))
	v229 = F_listSearchKey(m, v228, l0)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L27
	} else {
		goto L74
	}
L74:
	;
	if v229 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	v303 = v216
	goto L63
L77:
	;
	v259 = F_clientHasPendingReplies(m, l0)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L27
	} else {
		goto L86
	}
L78:
	;
	v239 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v239 + base.I64_extend_i32_u(v236)
	v244 = int32(1)
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v245&v244 != 0 {
		v252 = v244
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v255 != 0 {
		goto L77
	} else {
		goto L84
	}
L80:
	;
	v255 = v252
	goto L79
L81:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v248 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v250 = F_isImportSlotMigrationJob(m, v248)
	mBase = m.M
	v252 = v250
	goto L80
L83:
	;
	v255 = int32(0)
	goto L79
L84:
	;
	v257 = *(*int64)(unsafe.Add(mBase, _c_F_postWriteToClient[6]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v257
	goto L77
L85:
	;
	v300 = F_updateClientMemUsageAndBucket(m, l0)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L27
	} else {
		goto L98
	}
L86:
	;
	if v259 != 0 {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v261 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v261
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = int64(0)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+32))
	if v266 == v261 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v275&int32(64) == int32(0) {
		goto L85
	} else {
		goto L91
	}
L89:
	;
	v269 = int32(0)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+80))
	v273 = m.T0[v272].(func(*base.Module, int32, int32, int32) int32)(m, v265, v269, v269)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L27
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v280 = int32(-1)
	if v275&int32(1280) != 0 {
		v303 = v280
		goto L63
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v275 | int32(1024)
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_postWriteToClient[3]))
	if v287 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_postWriteToClient[5]))
	v296 = F_listAddNodeTail(m, v295, l0)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L27
	} else {
		goto L97
	}
L94:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_postWriteToClient[5]))
	v292 = F_listSearchKey(m, v291, l0)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L27
	} else {
		goto L95
	}
L95:
	;
	if v292 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L93
L97:
	;
	v303 = v280
	goto L63
L98:
	;
	v303 = int32(0)
	goto L63
L99:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_prepareForShutdown(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int64
	_ = v118
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(-1)
	v11 = *(*int64)(unsafe.Add(mBase, _c_F_prepareForShutdown[0]))
	if v11 != int64(0) {
		v149 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v149
L2:
	;
	v14 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_prepareForShutdown[1]))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_prepareForShutdown[2]))
	if v20|v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = l1&int32(-4) | int32(2)
	goto L5
L4:
	;
	v24 = l1
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_prepareForShutdown[3])) = v24
	if l0 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v24&int32(4) != 0 {
		goto L18
	} else {
		goto L19
	}
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_prepareForShutdown[4]))
	if int32(2) < v48 {
		goto L6
	} else {
		goto L16
	}
L8:
	;
	v28 = F_sdsempty(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_prepareForShutdown[5]))
	v34 = F_catClientInfoShortString(m, v28, l0, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_prepareForShutdown[4]))
	if int32(2) < v37 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_sdsfree(m, v34)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v34
	F__serverLog(m, int32(2), int32(_a_F_prepareForShutdown_0), v7)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L6
L16:
	;
	F__serverLog(m, int32(2), int32(_a_F_prepareForShutdown_1), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L6
L18:
	;
	v144 = F_finishShutdown(m)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L9
	} else {
		goto L37
	}
L19:
	;
	v59 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_prepareForShutdown[6]))
	if v60 == v59 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v63 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_prepareForShutdown[7]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v65 == v63 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v69 = v7 + int32(4)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v70
	goto L22
L22:
	;
	goto L23
L23:
	;
	v79 = v7 + int32(4)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v81 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v100 = int32(0)
	v102 = *(*int64)(unsafe.Add(mBase, _c_F_prepareForShutdown[8]))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_prepareForShutdown[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_prepareForShutdown[0])) = v102 + base.I64_extend_i32_s(v104*int32(1000))
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_prepareForShutdown[9]))
	goto L31
L25:
	;
	if v81 == int32(0) {
		goto L18
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v81+base.B2i32(v84 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v90
	goto L26
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+104))
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)+64))
	v98 = *(*int64)(unsafe.Add(mBase, _c_F_prepareForShutdown[10]))
	if v96 == v98 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	F_pauseActions(m, int32(1), int64(9223372036854775807), int32(29))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L9
	} else {
		goto L34
	}
L31:
	;
	if v112&int32(16) != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v114 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_prepareForShutdown[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v115
	v118 = *(*int64)(unsafe.Add(mBase, _c_F_prepareForShutdown[12]))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v118
	F_replicationFeedReplicas(m, int32(-1), v7+int32(4), int32(3))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_prepareForShutdown[4]))
	if int32(2) < v132 {
		v149 = v9
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F__serverLog(m, int32(2), int32(_a_F_prepareForShutdown_2), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v149 = v9
	goto L1
L37:
	;
	v149 = v144
	goto L1
}
func F_primaryexp(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
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
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 == int32(285) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	goto L14
L2:
	;
	v31 = m.G3
	F_luaX_syntaxerror(m, l0, v31+int32(_a_F_primaryexp_0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L12
	}
L3:
	;
	F_singlevar(m, l0, l1)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L11
	}
L4:
	;
	if v11 != int32(40) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_luaX_next(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v20 = F_subexpr(m, l0, l1, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_check_match(m, l0, int32(41), int32(40), v16)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_dischargevars(m, v26, l1)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L1
L11:
	;
	goto L1
L12:
	;
	goto L1
L13:
	;
	m.G0 = v8 + int32(48)
	return
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(90) < v42 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L39
	}
L17:
	;
	v59 = F_luaK_exp2anyreg(m, v10, l1)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L28
	}
L18:
	;
	F_field(m, l0, l1)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L27
	}
L19:
	;
	F_luaK_exp2nextreg(m, v10, l1)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L25
	}
L20:
	;
	if v42 == int32(91) {
		goto L17
	} else {
		goto L22
	}
L21:
	;
	switch v42 + int32(-40) {
	case 0:
		goto L19
	default:
		goto L13
	case 6:
		goto L18
	case 18:
		goto L16
	}
L22:
	;
	if v42 == int32(123) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	if v42 != int32(286) {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	goto L19
L25:
	;
	F_funcargs(m, l0, l1)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L14
L27:
	;
	goto L14
L28:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v66 = F_subexpr(m, l0, v8+int32(24), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_exp2val(m, v68, v8+int32(24))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v73 == int32(93) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L37
	}
L33:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v78 = F_luaX_token2str(m, l0, int32(93))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v78
	v81 = m.G3
	v84 = F_luaO_pushfstring(m, v76, v81+int32(_a_F_primaryexp_1), v8)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	F_luaX_syntaxerror(m, l0, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	F_luaK_indexed(m, v10, l1, v8+int32(24))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	goto L14
L39:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v97 == int32(285) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_luaX_next(m, l0)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L45
	}
L41:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v102 = F_luaX_token2str(m, l0, int32(285))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v102
	v105 = m.G3
	v110 = F_luaO_pushfstring(m, v100, v105+int32(_a_F_primaryexp_1), v8+int32(16))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	F_luaX_syntaxerror(m, l0, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v119 = F_luaK_stringK(m, v118, v115)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(-1)
	F_luaK_self(m, v10, l1, v8+int32(24))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	F_funcargs(m, l0, l1)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	goto L14
}
func F_printCrashReport(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	*(*int32)(unsafe.Add(mBase, _c_F_printCrashReport[0])) = int32(1)
	F_logServerInfo(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_printCrashReport[1]))
		F_logCurrentClient(m, v8, int32(_a_F_printCrashReport_0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_printCrashReport[2]))
			F_logCurrentClient(m, v13, int32(_a_F_printCrashReport_1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_serverLogRaw(m, int32(1027), int32(_a_F_printCrashReport_2))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v22 = F_sdsempty(m)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = int32(0)
						v27 = F_modulesCollectInfo(m, v22, v24, int32(1), v24)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							F_serverLogRaw(m, int32(1027), v27)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								F_sdsfree(m, v27)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									v33 = F_getConfigDebugInfo(m)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										F_serverLogRaw(m, int32(1027), int32(_a_F_printCrashReport_3))
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											F_serverLogRaw(m, int32(1027), v33)
											mBase = m.M
											v41 = m.ExcPending
											if v41 != 0 {
												return
											} else {
												F_sdsfree(m, v33)
												mBase = m.M
												v43 = m.ExcPending
												if v43 != 0 {
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
func F_processPendingCommandAndInputBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v2 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v6&int32(16) != 0 {
		v72 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v72
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v9&int32(2) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v40 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v9 & int32(-3)
	v17 = int32(_a_F_processPendingCommandAndInputBuffer_0)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_processPendingCommandAndInputBuffer[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_processPendingCommandAndInputBuffer[0])) = l0
	v21 = F_processCommand(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v32 = int32(_a_F_processPendingCommandAndInputBuffer_0)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_processPendingCommandAndInputBuffer[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_processPendingCommandAndInputBuffer[0])) = v18
	if v33 != 0 {
		goto L3
	} else {
		goto L12
	}
L6:
	;
	return int32(0)
L7:
	;
	if v21 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	F_commandProcessed(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v27 == int32(0) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v30 = F_updateClientMemUsageAndBucket(m, l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L5
L12:
	;
	return int32(-1)
L13:
	;
	v70 = F_processInputBuffer(m, l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L24
	}
L14:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	if base.Ui32(v66) <= base.Ui32(v65) {
		v72 = v2
		goto L1
	} else {
		goto L23
	}
L15:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-1)))))
	switch v45 & int32(7) {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		goto L17
	default:
		goto L14
	}
L16:
	;
	if v62 != 0 {
		goto L13
	} else {
		goto L22
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-17))))
	v62 = v61
	goto L16
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-9))))
	v62 = v58
	goto L16
L19:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+int32(-5)))))
	v62 = v55
	goto L16
L20:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-3)))))
	v62 = v52
	goto L16
L21:
	;
	v62 = int32(base.Ui32(v45) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	goto L14
L23:
	;
	goto L13
L24:
	;
	v72 = v70
	goto L1
}
func F_processRESP(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int64
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v206 int32
	_ = v206
	v8 = m.G0
	v10 = v8 - int32(1104)
	m.G0 = v10
	v15 = F_readLong(m, l0, int32(42), v10+int32(1100))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(1104)
	return v206
L2:
	;
	v206 = int32(0)
	goto L1
L3:
	;
	return int32(0)
L4:
	;
	if v15 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v21 = int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1100))
	if v22 < v21 {
		v206 = v21
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v27 = F_readString(m, l0, v10+int32(1096))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v27 == int32(0) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1096))
	v32 = int32(_a_F_processRESP_0)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v35 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	F_valkey_free(m, v31)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L3
	} else {
		goto L54
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
	v178 = F_snprintf(m, v10+int32(64), int32(1024), int32(_a_F_processRESP_1), v10+int32(16))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L3
	} else {
		goto L52
	}
L11:
	;
	F_valkey_free(m, v31)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L43
	}
L12:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v140 + int32(1)
	if v140 != 0 {
		goto L10
	} else {
		goto L42
	}
L13:
	;
	if v67-v69 == int32(0) {
		goto L12
	} else {
		goto L25
	}
L14:
	;
	v67 = F_tolower(m, v63)
	mBase = m.M
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v69 = F_tolower(m, v68)
	mBase = m.M
	goto L13
L15:
	;
	v37 = v31
	v38 = v32
	v39 = v35
	goto L18
L16:
	;
	v63 = int32(0)
	v64 = v32
	goto L14
L17:
	;
	v63 = v60 & int32(255)
	v64 = v59
	goto L14
L18:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v41 == int32(0) {
		v59 = v38
		v60 = v39
		goto L17
	} else {
		goto L20
	}
L19:
	;
	v59 = v53
	v60 = int32(0)
	goto L17
L20:
	;
	v45 = v39 & int32(255)
	if v45 == v41 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v52 = int32(1)
	v53 = v38 + v52
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v54 != 0 {
		v37 = v37 + v52
		v38 = v53
		v39 = v54
		goto L18
	} else {
		goto L24
	}
L22:
	;
	v47 = F_tolower(m, v45)
	mBase = m.M
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v49 = F_tolower(m, v48)
	mBase = m.M
	if v47 == v49 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v59 = v38
	v60 = v51
	goto L17
L24:
	;
	goto L19
L25:
	;
	v73 = int32(_a_F_processRESP_2)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v76 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	if v108-v110 != 0 {
		goto L11
	} else {
		goto L38
	}
L27:
	;
	v108 = F_tolower(m, v104)
	mBase = m.M
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v110 = F_tolower(m, v109)
	mBase = m.M
	goto L26
L28:
	;
	v78 = v31
	v79 = v73
	v80 = v76
	goto L31
L29:
	;
	v104 = int32(0)
	v105 = v73
	goto L27
L30:
	;
	v104 = v101 & int32(255)
	v105 = v100
	goto L27
L31:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v82 == int32(0) {
		v100 = v79
		v101 = v80
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v100 = v94
	v101 = int32(0)
	goto L30
L33:
	;
	v86 = v80 & int32(255)
	if v86 == v82 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v93 = int32(1)
	v94 = v79 + v93
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v95 != 0 {
		v78 = v78 + v93
		v79 = v94
		v80 = v95
		goto L31
	} else {
		goto L37
	}
L35:
	;
	v88 = F_tolower(m, v86)
	mBase = m.M
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v90 = F_tolower(m, v89)
	mBase = m.M
	if v88 == v90 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v100 = v79
	v101 = v92
	goto L30
L37:
	;
	goto L32
L38:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v114 = v112 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v114
	if v114 == int32(0) {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l1
	v125 = F_snprintf(m, v10+int32(64), int32(1024), int32(_a_F_processRESP_3), v10+int32(48))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	v128 = *(*int64)(unsafe.Add(mBase, _c_F_processRESP[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v10 + int32(64)
	v138 = F_snprintf(m, int32(_a_F_processRESP_4), int32(1044), int32(_a_F_processRESP_5), v10+int32(32))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	goto L9
L42:
	;
	goto L11
L43:
	;
	v147 = int32(1)
	if v22 != v147 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v153 = v147
	goto L46
L45:
	;
	v206 = int32(1)
	goto L1
L46:
	;
	v160 = F_readString(m, l0, v10+int32(1096))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	if v160 == int32(0) {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1096))
	F_valkey_free(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v167 = int32(1)
	v169 = v153 + v167
	if v169 != v22 {
		v153 = v169
		goto L46
	} else {
		goto L51
	}
L51:
	;
	v206 = v167
	goto L1
L52:
	;
	v181 = *(*int64)(unsafe.Add(mBase, _c_F_processRESP[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v10 + int32(64)
	v189 = F_snprintf(m, int32(_a_F_processRESP_4), int32(1044), int32(_a_F_processRESP_5), v10)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	goto L9
L54:
	;
	goto L2
}
func F_propagateDeletion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(_a_F_propagateDeletion_0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_propagateDeletion[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_propagateDeletion[0])) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	if l2 != 0 {
		v20 = int32(244)
	} else {
		v20 = int32(240)
	}
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_propagateDeletion[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_alsoPropagate(m, v24, v9+int32(8), int32(2), int32(3), l3)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_propagateDeletion[0])) = v12
		m.G0 = v9 + int32(16)
		return
	}
}
func F_propagateNow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	v6 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[0]))
	if v11 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_propagateNow_0), int32(_a_F_propagateNow_1), int32(3606))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L51
	} else {
		goto L58
	}
L2:
	;
	return
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[1]))
	if v15 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v17 = l3 & int32(1)
	if v17 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[2]))
	goto L26
L6:
	;
	if l3&int32(2) == int32(0) {
		goto L2
	} else {
		goto L9
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[3]))
	if v21 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v26 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(16)
	m.G0 = v31
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[4]))
	if v35 == v26 {
		v67 = v26
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v67 != 0 {
		goto L5
	} else {
		goto L21
	}
L11:
	;
	m.G0 = v31 + int32(16)
	goto L10
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[5]))
	if v39 == int32(0) {
		v67 = v26
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_propagateNow[6])))
	v44 = v31 + int32(8)
	F_listRewind(m, v42, v44)
	mBase = m.M
	v46 = int32(0)
	v49 = F_listNext(m, v44)
	mBase = m.M
	if v49 == v46 {
		v67 = v46
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v54 = v49
	goto L15
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v56 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v67 = v46
	goto L11
L17:
	;
	v65 = F_listNext(m, v31+int32(8))
	mBase = m.M
	if v65 != 0 {
		v54 = v65
		goto L15
	} else {
		goto L20
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+156))
	if base.Ui32(v57+int32(-18)) <= base.Ui32(int32(2)) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v67 = int32(1)
	goto L11
L20:
	;
	goto L16
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[7]))
	if v73 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[8]))
	if v75 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v76 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[9]))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	if v78 == v76 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	goto L5
L25:
	;
	v93 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[3]))
	v97 = v17 & base.B2i32(v94 != v93)
	v99 = l3 & int32(2)
	if v99 == v93 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	if v83&int32(16) == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[10]))
	if v88 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v89 = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[11]))
	if v90 == v89 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	if v97 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L31:
	;
	v121 = int32(0)
	v124 = m.G0
	v126 = v124 - int32(16)
	m.G0 = v126
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[4]))
	if v130 == v121 {
		v162 = v121
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v113 = int32(0)
	if v99 == v113 {
		v167 = int32(1)
		v169 = v113
		goto L30
	} else {
		goto L37
	}
L33:
	;
	if v97 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[7]))
	if v104 != 0 {
		v119 = int32(1)
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v105 = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[8]))
	if v107 != 0 {
		v119 = v105
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v108 = int32(0)
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[9]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+20))
	v119 = base.B2i32(v110 == v108)
	goto L31
L37:
	;
	v119 = v113
	goto L31
L38:
	;
	v167 = v119
	v169 = v162
	goto L30
L39:
	;
	m.G0 = v126 + int32(16)
	goto L38
L40:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_propagateNow[5]))
	if v134 == int32(0) {
		v162 = v121
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v134)+uint32(_c_F_propagateNow[6])))
	v139 = v126 + int32(8)
	F_listRewind(m, v137, v139)
	mBase = m.M
	v141 = int32(0)
	v144 = F_listNext(m, v139)
	mBase = m.M
	if v144 == v141 {
		v162 = v141
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v149 = v144
	goto L43
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if v151 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v162 = v141
	goto L39
L45:
	;
	v160 = F_listNext(m, v126+int32(8))
	mBase = m.M
	if v160 != 0 {
		v149 = v160
		goto L43
	} else {
		goto L48
	}
L46:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v150)+156))
	if base.Ui32(v152+int32(-18)) <= base.Ui32(int32(2)) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v162 = int32(1)
	goto L39
L48:
	;
	goto L44
L49:
	;
	if v167 != 0 {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	F_feedAppendOnlyFile(m, l0, l1, l2)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	return
L52:
	;
	goto L49
L53:
	;
	if v169 == int32(0) {
		goto L2
	} else {
		goto L56
	}
L54:
	;
	F_replicationFeedReplicas(m, l0, l1, l2)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	F_clusterFeedSlotExportJobs(m, l0, l1, l2, l4)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L51
	} else {
		goto L57
	}
L57:
	;
	goto L2
L58:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_publishCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int64
	_ = v27
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_publishCommand[0]))
	if v11 == int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		v21 = int32(0)
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_publishCommand[1]))
		*(*int32)(unsafe.Add(mBase, uint32(v8+int32(24)))) = v22
		v27 = *(*int64)(unsafe.Add(mBase, _c_F_publishCommand[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(16)))) = v27
		v32 = *(*int64)(unsafe.Add(mBase, _c_F_publishCommand[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(8)))) = v32
		v35 = *(*int64)(unsafe.Add(mBase, _c_F_publishCommand[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v35
		v37 = F_pubsubPublishMessageInternal(m, v18, v17, v8)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, _c_F_publishCommand[5]))
			if v40 == int32(0) {
				F_forceCommandPropagation(m, l0, int32(2))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v37))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						m.G0 = v8 + int32(32)
						return
					}
				}
			} else {
				F_clusterPropagatePublish(m, v18, v17, int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_publishCommand[5]))
					if v47 != 0 {
						F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v37))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					} else {
						F_forceCommandPropagation(m, l0, int32(2))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v37))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								m.G0 = v8 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_sentinelPublishCommand(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	}
}
func F_punsubscribeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v3 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	goto L12
L2:
	;
	v24 = F_pubsubUnsubscribeAllPatterns(m, l0, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L10
	}
L3:
	;
	if v3 < int32(2) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v10 = int32(1)
	goto L5
L5:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11+v10<<(uint(int32(2))%32))))
	v17 = F_pubsubUnsubscribePattern(m, l0, v15, int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v20 = v10 + int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v20 < v21 {
		v10 = v20
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L1
L10:
	;
	goto L1
L11:
	;
	return
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	goto L13
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	goto L14
L14:
	;
	if v30+v31+(v35+v36) != int32(0)-(v42+v43) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v47&int32(262144) == int32(0) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v47 & int32(-262145)
	v55 = int32(_a_F_punsubscribeCommand_0)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_punsubscribeCommand[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_punsubscribeCommand[0])) = v57 + int32(-1)
	goto L11
}
func F_pushclosure(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	if v10 < v12 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v43 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v41+v43<<(uint(int32(2))%32)))) = v42
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+5)))
	if v51&int32(3) == int32(0) {
		v78 = v43
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v31 = int32(2)
	v39 = F__emscripten_memset_bulkmem(m, v30+v12<<(uint(v31)%32), base.I32_extend8_s(int32(0)), (v26-v12)<<(uint(v31)%32))
	mBase = m.M
	goto L8
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v41 = v29
	goto L1
L4:
	;
	v14 = m.G3
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v23 = F_luaM_growaux_(m, v15, v16, v11+int32(52), int32(4), int32(262143), v14+int32(_a_F_pushclosure_0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	if v12 < v26 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L3
L8:
	;
	v41 = v30
	goto L1
L9:
	;
	v81 = F_luaK_codeABx(m, v9, int32(36), int32(0), v78)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L15
	}
L10:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+5)))
	if v56&int32(4) == int32(0) {
		v78 = v43
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+21)))
	if v63 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	v78 = v75 + int32(-1)
	goto L9
L13:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+20)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+5)))
	v73 = v67&int32(3) | v70&int32(248)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+5)) = uint8(v73)
	goto L12
L14:
	;
	F_reallymarkobject(m, v62, v42)
	mBase = m.M
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = int64(-1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+72)))
	if v89 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	v99 = int32(0)
	goto L18
L18:
	;
	v105 = l1 + int32(51) + v99<<(uint(int32(1))%32)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v111 = int32(0)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	v114 = F_luaK_codeABC(m, v9, base.B2i32(v106 != int32(6))<<(uint(int32(2))%32), v111, v112, v111)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L20
	}
L19:
	;
	goto L16
L20:
	;
	v117 = v99 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+72)))
	if base.Ui32(v117) < base.Ui32(v119) {
		v99 = v117
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
}
func F_putchar(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_do_putc_2(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
