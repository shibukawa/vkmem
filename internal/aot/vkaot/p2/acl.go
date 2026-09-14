package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_ACLAddAllowedFirstArg(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v7 != 0 {
		v12 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = int32(0)
	v15 = l1 << (uint(int32(2)) % 32)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12+v15)))
	if v17 == v13 {
		v80 = v13
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v9 = F_valkey_calloc(m, int32(4096))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v9
	v12 = v9
	goto L1
L5:
	;
	return
L6:
	;
	v83 = v80 << (uint(int32(2)) % 32)
	v86 = F_valkey_realloc(m, v17, v83+int32(8))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L3
	} else {
		goto L25
	}
L7:
	;
	v20 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v21 == v20 {
		v80 = v20
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v27 = v21
	v28 = v20
	goto L9
L9:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v80 = v71
	goto L6
L11:
	;
	if v64-v66 == int32(0) {
		goto L5
	} else {
		goto L23
	}
L12:
	;
	v64 = F_tolower(m, v60)
	mBase = m.M
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v66 = F_tolower(m, v65)
	mBase = m.M
	goto L11
L13:
	;
	v34 = v27
	v35 = l2
	v36 = v32
	goto L16
L14:
	;
	v60 = int32(0)
	v61 = l2
	goto L12
L15:
	;
	v60 = v57 & int32(255)
	v61 = v56
	goto L12
L16:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v38 == int32(0) {
		v56 = v35
		v57 = v36
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v56 = v50
	v57 = int32(0)
	goto L15
L18:
	;
	v42 = v36 & int32(255)
	if v42 == v38 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v49 = int32(1)
	v50 = v35 + v49
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v51 != 0 {
		v34 = v34 + v49
		v35 = v50
		v36 = v51
		goto L16
	} else {
		goto L22
	}
L20:
	;
	v44 = F_tolower(m, v42)
	mBase = m.M
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v46 = F_tolower(m, v45)
	mBase = m.M
	if v44 == v46 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v56 = v35
	v57 = v48
	goto L15
L22:
	;
	goto L17
L23:
	;
	v71 = v28 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v17+v71<<(uint(int32(2))%32))))
	if v75 != 0 {
		v27 = v75
		v28 = v71
		goto L9
	} else {
		goto L24
	}
L24:
	;
	goto L10
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v88+v15))) = v86
	v91 = F_sdsnew(m, l2)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v93+v15)))
	v96 = v95 + v83
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v91
	goto L5
}
func F_ACLAuthenticateUser(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
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
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v284 int64
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = F_checkModuleAuthentication(m, l0, l1, l2, l3)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v292
L2:
	;
	return int32(0)
L3:
	;
	if v14 != int32(2) {
		v292 = v14
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = F_ACLCheckUserCredentials(m, l1, l2)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	v284 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v285 = F_objectGetVal(m, l1)
	mBase = m.M
	v286 = int32(0)
	F_moduleFireAuthenticationEvent(m, v284, v285, v286, base.B2i32(v20 == v286))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L2
	} else {
		goto L58
	}
L6:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v269 = int32(2)
	v273 = int32(0)
	v274 = F_objectGetVal(m, l1)
	mBase = m.M
	F_addACLLogEntry(m, l0, int32(4), int32(base.Ui32(v268)>>(uint(v269)%32))&v269, v273, v274, v273)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L2
	} else {
		goto L57
	}
L7:
	;
	if v20 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v22 = F_objectGetVal(m, l1)
	mBase = m.M
	v23 = int32(0)
	v25 = F_objectGetVal(m, l1)
	mBase = m.M
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(-1)))))
	switch v28 & int32(7) {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	default:
		v45 = v23
		goto L9
	}
L9:
	;
	v46 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v46
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_ACLAuthenticateUser[0]))
	v51 = v12 + int32(12)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v45 == v46 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(-17))))
	v45 = v44
	goto L9
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(-9))))
	v45 = v41
	goto L9
L12:
	;
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+int32(-5)))))
	v45 = v38
	goto L9
L13:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(-3)))))
	v45 = v35
	goto L9
L14:
	;
	v45 = int32(base.Ui32(v28) >> (uint(int32(3)) % 32))
	goto L9
L15:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v247
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	goto L54
L16:
	;
	if v204 != v45 {
		goto L43
	} else {
		goto L44
	}
L17:
	;
	v195 = int32(0)
	v201 = v60
	v202 = v61
	v204 = v195
	v208 = v195
	goto L16
L18:
	;
	if base.Ui32(v61) < base.Ui32(int32(8)) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v72 = v60
	v73 = v61
	v75 = int32(0)
	goto L21
L20:
	;
	v201 = v185
	v202 = v186
	v204 = v188
	v208 = base.B2i32(v191 != int32(0))
	goto L16
L21:
	;
	v81 = int32(base.Ui32(v73) >> (uint(int32(3)) % 32))
	v82 = int32(4)
	v83 = v72 + v82
	if v73&v82 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v185 = v176
	v186 = v177
	v188 = v161
	v191 = v166
	goto L20
L23:
	;
	v166 = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v83+v81+(v166-v81)&int32(3)+v154<<(uint(int32(2))%32))))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if base.Ui32(v177) < base.Ui32(int32(8)) {
		v185 = v176
		v186 = v177
		v188 = v161
		v191 = v166
		goto L20
	} else {
		goto L41
	}
L24:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v75))))
	v132 = int32(0)
	goto L35
L25:
	;
	v88 = int32(0)
	if base.Ui32(v45) <= base.Ui32(v75) {
		v121 = v75
		v124 = v88
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if v124 == v81 {
		v154 = v88
		v161 = v121
		goto L23
	} else {
		goto L33
	}
L27:
	;
	v98 = v75
	v101 = v88
	goto L28
L28:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v101))))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v98))))
	if v104 != v106 {
		v121 = v98
		v124 = v101
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v121 = v109
	v124 = v111
	goto L26
L30:
	;
	v108 = int32(1)
	v109 = v98 + v108
	v111 = v101 + v108
	if base.Ui32(v81) <= base.Ui32(v111) {
		v121 = v109
		v124 = v111
		goto L26
	} else {
		goto L31
	}
L31:
	;
	if base.Ui32(v109) < base.Ui32(v45) {
		v98 = v109
		v101 = v111
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v185 = v72
	v186 = v73
	v188 = v121
	v191 = v124
	goto L20
L34:
	;
	if v132 != v81 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v132))))
	if v145 == v129&int32(255) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v147 = int32(1)
	v149 = v132 + v147
	if v149 != v81 {
		v132 = v149
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v201 = v72
	v202 = v73
	v204 = v75
	v208 = v147
	goto L16
L39:
	;
	v154 = v132
	v161 = v75 + int32(1)
	goto L23
L40:
	;
	v185 = v72
	v186 = v73
	v188 = v75
	v191 = v81
	goto L20
L41:
	;
	if base.Ui32(v161) < base.Ui32(v45) {
		v72 = v176
		v73 = v177
		v75 = v161
		goto L21
	} else {
		goto L42
	}
L42:
	;
	goto L22
L43:
	;
	goto L15
L44:
	;
	if v202&int32(1) == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v216 = v202 & int32(4)
	if v208&base.B2i32(v216 != int32(0)) != 0 {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	if v51 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	if v202&int32(2) != 0 {
		v242 = int32(0)
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v242
	goto L43
L49:
	;
	v226 = int32(3)
	v227 = int32(base.Ui32(v202) >> (uint(v226) % 32))
	if v216 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v237 = int32(4)
	goto L52
L51:
	;
	v237 = v227 << (uint(int32(2)) % 32)
	goto L52
L52:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v201+v227+(int32(0)-v227)&v226+v237+int32(4))))
	v242 = v241
	goto L48
L53:
	;
	F_moduleNotifyUserChanged(m, l0)
	mBase = m.M
	v279 = v23
	goto L5
L54:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v251&int32(-25165825) | int32(_a_F_ACLAuthenticateUser_0) | int32(16777216)
	goto L53
L57:
	;
	v279 = int32(1)
	goto L5
L58:
	;
	v292 = v279
	goto L1
}
func F_ACLChangeSelectorPerm(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v182 int32
	_ = v182
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v208 int32
	_ = v208
	v11 = m.G0
	v13 = v11 - int32(64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	if base.Ui32(int32(1023)) < base.Ui32(v15) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v48 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v19 = base.I64_extend_i32_u(v15)
	v20 = int64(1) << (uint(v19) % 64)
	v22 = int64(base.Ui64(v19) >> (uint(int64(6)) % 64))
	if l2 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v35 = l0 + base.I32_wrap_i64(v22)<<(uint(int32(3))%32)
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v36 & (v20 ^ int64(-1))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v41 & int32(-5)
	goto L1
L4:
	;
	v28 = l0 + base.I32_wrap_i64(v22)<<(uint(int32(3))%32)
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v29 | v20
	goto L1
L5:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	if v112 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+v15<<(uint(int32(2))%32))))
	if v54 == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v57 == int32(0) {
		v92 = v54
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_valkey_free(m, v92)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L12
	} else {
		goto L15
	}
L9:
	;
	v68 = v57
	v70 = int32(0)
	goto L10
L10:
	;
	F_sdsfree(m, v68)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v92 = v77
	goto L8
L12:
	;
	return
L13:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v74 = int32(2)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v15<<(uint(v74)%32))))
	v79 = v70 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77+v79<<(uint(v74)%32))))
	if v83 != 0 {
		v68 = v83
		v70 = v79
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v96+v15<<(uint(int32(2))%32)))) = int32(0)
	goto L5
L16:
	;
	m.G0 = v13 + int32(64)
	return
L17:
	;
	v116 = v13 + int32(16)
	v117 = int32(1)
	v118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v116)+14)) = uint8(v118)
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v116)+24)) = v118
	*(*uint8)(unsafe.Add(mBase, uint32(v116)+15)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+8)) = int32(-1)
	if v112 == v118 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v139 = F_hashtableNext(m, v13+int32(16), v13+int32(12))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L12
	} else {
		goto L23
	}
L19:
	;
	goto L18
L20:
	;
	goto L21
L21:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v112)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+24)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v112)+40)) = v116
	goto L19
L22:
	;
	F_hashtableCleanupIterator(m, v13+int32(16))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L12
	} else {
		goto L33
	}
L23:
	;
	if v139 == int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v144 = l0 + int32(8)
	goto L25
L25:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+136))
	if base.Ui32(int32(1023)) < base.Ui32(v156) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L22
L27:
	;
	v193 = F_hashtableNext(m, v13+int32(16), v13+int32(12))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L12
	} else {
		goto L31
	}
L28:
	;
	v160 = base.I64_extend_i32_u(v156)
	v161 = int64(1) << (uint(v160) % 64)
	v163 = int64(base.Ui64(v160) >> (uint(int64(6)) % 64))
	if l2 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v176 = v144 + base.I32_wrap_i64(v163)<<(uint(int32(3))%32)
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v176)))
	*(*int64)(unsafe.Add(mBase, uint32(v176))) = v177 & (v161 ^ int64(-1))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v182 & int32(-5)
	goto L27
L30:
	;
	v169 = v144 + base.I32_wrap_i64(v163)<<(uint(int32(3))%32)
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v169)))
	*(*int64)(unsafe.Add(mBase, uint32(v169))) = v170 | v161
	goto L27
L31:
	;
	if v193 != 0 {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	goto L26
L33:
	;
	goto L16
}
func F_ACLCheckAllPerm(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v12 = v10 & int32(8)
	if v12 != 0 {
		v13 = int32(112)
	} else {
		v13 = int32(96)
	}
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0+v13)))
	if v12 != 0 {
		v18 = int32(52)
	} else {
		v18 = int32(28)
	}
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15+v18)))
	v21 = F_ACLCheckAllUserCommandPerm(m, v4, v5, v6, v7, v20, l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		return v21
	}
}
func F_ACLCheckAllUserCommandPerm(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	v11 = m.G0
	v13 = v11 - int32(2080)
	m.G0 = v13
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(2080)
	return v110
L2:
	;
	v16 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+2068)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v16
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = v13 + int32(2072)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v24
	goto L4
L3:
	;
	v110 = int32(0)
	goto L1
L4:
	;
	v29 = v13 + int32(2072)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v31 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
	v110 = int32(1)
	goto L1
L6:
	;
	if v31 == int32(0) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31+base.B2i32(v34 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v40
	goto L7
L9:
	;
	v45 = v31
	v52 = int32(1)
	v53 = v16
	goto L10
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v60 = F_ACLSelectorCheckCmd(m, v55, l1, l2, l3, v13+int32(2068), v13+int32(4), l4)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v77
	if v92 == int32(0) {
		v110 = v76
		goto L1
	} else {
		goto L27
	}
L12:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+2068))
	if v52 < v60 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	return int32(0)
L14:
	;
	if v60 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v64 = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v65 == v64 {
		v110 = v64
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_getKeysFreeResult(m, v13+int32(8))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v110 = v64
	goto L1
L18:
	;
	v79 = v13 + int32(2072)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v81 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v76 = v60
	v77 = v72
	goto L18
L20:
	;
	if v60 != v52 {
		v76 = v52
		v77 = v53
		goto L18
	} else {
		goto L21
	}
L21:
	;
	if v72 <= v53 {
		v76 = v52
		v77 = v53
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	if v81 != 0 {
		v45 = v81
		v52 = v76
		v53 = v77
		goto L10
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v81+base.B2i32(v84 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v90
	goto L24
L26:
	;
	goto L11
L27:
	;
	F_getKeysFreeResult(m, v13+int32(8))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v110 = v76
	goto L1
}
func F_ACLDescribeSelectorCommandRules(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
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
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v354 int32
	_ = v354
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = F_sdsempty(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = F_valkey_malloc(m, int32(160))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_ACLDescribeSelectorCommandRules[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v22 | int32(16)
	v26 = F_listCreate(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+140)) = v26
	v29 = F_listCreate(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v29
	v33 = F_intsetNew(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+136)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+152)) = v33
	v38 = F_sdsempty(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+148)) = v38
	v41 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = int32(5)
	v58 = F__emscripten_memset_bulkmem(m, v19+v41, base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L8
L8:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	v63 = base.B2i32(int64(-1) < v61)
	if int64(-1) < v61 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v64 = int32(_a_F_ACLDescribeSelectorCommandRules_0)
	goto L11
L10:
	;
	v64 = int32(_a_F_ACLDescribeSelectorCommandRules_1)
	goto L11
L11:
	;
	v65 = F_sdscat(m, v14, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if int64(-1) < v61 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v69 = int32(_a_F_ACLDescribeSelectorCommandRules_2)
	goto L15
L14:
	;
	v69 = int32(_a_F_ACLDescribeSelectorCommandRules_3)
	goto L15
L15:
	;
	v71 = F_ACLSetSelector(m, v19, v69, int32(-1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v78 = F_sdssplitargs(m, v75, v12+int32(28))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	F__serverAssert(m, int32(_a_F_ACLDescribeSelectorCommandRules_4), int32(_a_F_ACLDescribeSelectorCommandRules_5), int32(811))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L100
	}
L18:
	;
	if v78 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v82 <= int32(0) {
		v117 = v82
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+int32(-1)))))
	switch v125 & int32(7) {
	case 0:
		goto L31
	case 1:
		goto L35
	case 2:
		goto L34
	case 3:
		goto L33
	case 4:
		goto L32
	default:
		v153 = v65
		v154 = v117
		goto L29
	}
L21:
	;
	v91 = int32(0)
	goto L22
L22:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v78+v91<<(uint(int32(2))%32))))
	v99 = F_ACLSetSelector(m, v19, v97, int32(-1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v117 = v111
	goto L20
L24:
	;
	v110 = v91 + int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v110 < v111 {
		v91 = v110
		goto L22
	} else {
		goto L28
	}
L25:
	;
	if v99 == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	F__serverAssert(m, int32(_a_F_ACLDescribeSelectorCommandRules_6), int32(_a_F_ACLDescribeSelectorCommandRules_5), int32(815))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	goto L23
L29:
	;
	F_sdsfreesplitres(m, v78, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L42
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v122
	v150 = F_sdscatfmt(m, v65, int32(_a_F_ACLDescribeSelectorCommandRules_7), v12+int32(16))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L41
	}
L31:
	;
	if int32(base.Ui32(v125)>>(uint(int32(3))%32)) == int32(0) {
		v153 = v65
		v154 = v117
		goto L29
	} else {
		goto L40
	}
L32:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v122+int32(-17))))
	if v141 != 0 {
		goto L30
	} else {
		goto L39
	}
L33:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v122+int32(-9))))
	if v136 == int32(0) {
		v153 = v65
		v154 = v117
		goto L29
	} else {
		goto L38
	}
L34:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122+int32(-5)))))
	if v133 != 0 {
		goto L30
	} else {
		goto L37
	}
L35:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+int32(-3)))))
	if v130 != 0 {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v153 = v65
	v154 = v117
	goto L29
L37:
	;
	v153 = v65
	v154 = v117
	goto L29
L38:
	;
	goto L30
L39:
	;
	v153 = v65
	v154 = v117
	goto L29
L40:
	;
	goto L30
L41:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v153 = v150
	v154 = v152
	goto L29
L42:
	;
	v166 = v153 + int32(-1)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	v169 = v167 & int32(7)
	switch v169 {
	case 0:
		goto L50
	case 1:
		goto L49
	case 2:
		goto L48
	case 3:
		goto L47
	case 4:
		goto L46
	default:
		goto L44
	}
L43:
	;
	v260 = l0 + int32(8)
	v261 = int32(128)
	goto L82
L44:
	;
	goto L43
L45:
	;
	if v184 == int32(0) {
		goto L44
	} else {
		goto L51
	}
L46:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v153+int32(-17))))
	v184 = v183
	goto L45
L47:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v153+int32(-9))))
	v184 = v180
	goto L45
L48:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153+int32(-5)))))
	v184 = v177
	goto L45
L49:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153+int32(-3)))))
	v184 = v174
	goto L45
L50:
	;
	v184 = int32(base.Ui32(v167) >> (uint(int32(3)) % 32))
	goto L45
L51:
	;
	v190 = int32(-1)&v184 + int32(-2)
	v194 = int32(0)&v184 + int32(0)
	v197 = v190 - v194 + int32(1)
	switch v169 {
	default:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	case 4:
		goto L53
	}
L52:
	;
	v213 = int32(0)
	v215 = base.B2i32(base.Ui32(v194) < base.Ui32(v212))
	if base.Ui32(v194) < base.Ui32(v212) {
		goto L59
	} else {
		goto L60
	}
L53:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v153+int32(-17))))
	v212 = v211
	goto L52
L54:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v153+int32(-9))))
	v212 = v208
	goto L52
L55:
	;
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153+int32(-5)))))
	v212 = v205
	goto L52
L56:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153+int32(-3)))))
	v212 = v202
	goto L52
L57:
	;
	v212 = int32(base.Ui32(v167) >> (uint(int32(3)) % 32))
	goto L52
L58:
	;
	v229 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v153+v223))) = uint8(v229)
	switch v169 {
	default:
		goto L76
	case 1:
		goto L75
	case 2:
		goto L74
	case 3:
		goto L73
	case 4:
		goto L72
	}
L59:
	;
	v216 = v194
	goto L61
L60:
	;
	v216 = v213
	goto L61
L61:
	;
	v217 = v212 - v216
	if base.Ui32(v197) < base.Ui32(v217) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v219 = v197
	goto L64
L63:
	;
	v219 = v217
	goto L64
L64:
	;
	if v190 < v194 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v221 = v213
	goto L67
L66:
	;
	v221 = v219
	goto L67
L67:
	;
	if base.Ui32(v194) < base.Ui32(v212) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v223 = v221
	goto L70
L69:
	;
	v223 = int32(0)
	goto L70
L70:
	;
	if v223 == int32(0) {
		goto L58
	} else {
		goto L71
	}
L71:
	;
	v227 = F_memmove(m, v153, v153+v216, v223)
	mBase = m.M
	goto L58
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v153+int32(-17)))) = base.I64_extend_i32_u(v223)
	goto L44
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153+int32(-9)))) = v223
	goto L43
L74:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v153+int32(-5)))) = uint16(v223)
	goto L43
L75:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v153+int32(-3)))) = uint8(v223)
	goto L43
L76:
	;
	v232 = v223 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v232)
	goto L43
L77:
	;
	F_ACLFreeSelector(m, v19)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L99
	}
L78:
	;
	if v325 == int32(0) {
		goto L77
	} else {
		goto L94
	}
L79:
	;
	v325 = int32(0)
	goto L78
L80:
	;
	v297 = v292
	v298 = v293
	v299 = v294
	goto L90
L81:
	;
	if v282 == int32(0) {
		goto L79
	} else {
		goto L88
	}
L82:
	;
	if (v260|v58)&int32(3) != 0 {
		v292 = v58
		v293 = v260
		v294 = v261
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v269 = v58
	v270 = v260
	v271 = v261
	goto L84
L84:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	if v274 != v275 {
		v292 = v269
		v293 = v270
		v294 = v271
		goto L80
	} else {
		goto L86
	}
L85:
	;
	goto L81
L86:
	;
	v277 = int32(4)
	v278 = v270 + v277
	v280 = v269 + v277
	v282 = v271 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v282) {
		v269 = v280
		v270 = v278
		v271 = v282
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v292 = v280
	v293 = v278
	v294 = v282
	goto L80
L89:
	;
	v325 = v302 - v303
	goto L78
L90:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if v302 != v303 {
		goto L89
	} else {
		goto L92
	}
L92:
	;
	v305 = int32(1)
	v310 = v299 + int32(-1)
	if v310 == int32(0) {
		goto L79
	} else {
		goto L93
	}
L93:
	;
	v297 = v297 + v305
	v298 = v298 + v305
	v299 = v310
	goto L90
L94:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_ACLDescribeSelectorCommandRules[1]))
	if int32(3) < v329 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	F__serverPanic_1(m, int32(_a_F_ACLDescribeSelectorCommandRules_5), int32(831), int32(_a_F_ACLDescribeSelectorCommandRules_8), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v153
	F__serverLog(m, int32(3), int32(_a_F_ACLDescribeSelectorCommandRules_9), v12)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	m.G0 = v12 + int32(32)
	return v153
L100:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ACLFreeUserAndKillClients(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ACLFreeUserAndKillClients[0]))
	v11 = v6 + int32(8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v12
	goto L1
L1:
	;
	v17 = v6 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_sdsfree(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L15
	} else {
		goto L21
	}
L3:
	;
	if v19 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19+base.B2i32(v22 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v28
	goto L4
L6:
	;
	v34 = v19
	goto L7
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+328))
	if v36 != l0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v62 = v6 + int32(8)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v64 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_ACLFreeUserAndKillClients[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+328)) = v39
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+204))
	goto L13
L11:
	;
	F_freeClientOrCloseLater(m, v35, int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+204)) = v43 & int32(-8388609)
	goto L11
L15:
	;
	return
L16:
	;
	goto L9
L17:
	;
	if v64 != 0 {
		v34 = v64
		goto L7
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v64+base.B2i32(v67 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v73
	goto L18
L20:
	;
	goto L8
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v81 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_listRelease(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L15
	} else {
		goto L25
	}
L23:
	;
	F_decrRefCount(m, v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L22
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_listRelease(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	m.G0 = v6 + int32(16)
	return
}
func F_ACLGetCommandID(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_sdsdup(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-1)))))
	switch v20 & int32(7) {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	case 3:
		goto L7
	case 4:
		goto L6
	default:
		goto L4
	}
L3:
	;
	v56 = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_ACLGetCommandID[0]))
	if v58 != 0 {
		v63 = v58
		goto L15
	} else {
		goto L16
	}
L4:
	;
	goto L3
L5:
	;
	if v37 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-17))))
	v37 = v36
	goto L5
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-9))))
	v37 = v33
	goto L5
L8:
	;
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-5)))))
	v37 = v30
	goto L5
L9:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
	v37 = v27
	goto L5
L10:
	;
	v37 = int32(base.Ui32(v20) >> (uint(int32(3)) % 32))
	goto L5
L11:
	;
	v42 = int32(0)
	goto L12
L12:
	;
	v45 = v11 + v42
	v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v45))))
	v47 = F_tolower(m, v46)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v47)
	v50 = v42 + int32(1)
	if v50 != v37 {
		v42 = v50
		goto L12
	} else {
		goto L14
	}
L13:
	;
	goto L4
L14:
	;
	goto L13
L15:
	;
	v65 = v11 + int32(-1)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	switch v66 & int32(7) {
	case 0:
		goto L23
	case 1:
		goto L22
	case 2:
		goto L21
	case 3:
		goto L20
	case 4:
		goto L19
	default:
		v83 = v56
		goto L18
	}
L16:
	;
	v60 = F_raxNew(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLGetCommandID[0])) = v60
	v63 = v60
	goto L15
L18:
	;
	v85 = v9 + int32(12)
	v86 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v83 == v86 {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-17))))
	v83 = v82
	goto L18
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-9))))
	v83 = v79
	goto L18
L21:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-5)))))
	v83 = v76
	goto L18
L22:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
	v83 = v73
	goto L18
L23:
	;
	v83 = int32(base.Ui32(v66) >> (uint(int32(3)) % 32))
	goto L18
L24:
	;
	m.G0 = v9 + int32(16)
	return v324
L25:
	;
	v286 = int32(0)
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_ACLGetCommandID[0]))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	switch v289 & int32(7) {
	case 0:
		goto L71
	case 1:
		goto L70
	case 2:
		goto L69
	case 3:
		goto L68
	case 4:
		goto L67
	default:
		v306 = v286
		goto L66
	}
L26:
	;
	if v279 == int32(0) {
		goto L25
	} else {
		goto L64
	}
L27:
	;
	if v238 != v83 {
		v279 = v86
		goto L54
	} else {
		goto L55
	}
L28:
	;
	v229 = int32(0)
	v235 = v94
	v236 = v95
	v238 = v229
	v242 = v229
	goto L27
L29:
	;
	if base.Ui32(v95) < base.Ui32(int32(8)) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v106 = v94
	v107 = v95
	v109 = int32(0)
	goto L32
L31:
	;
	v235 = v219
	v236 = v220
	v238 = v222
	v242 = base.B2i32(v225 != int32(0))
	goto L27
L32:
	;
	v115 = int32(base.Ui32(v107) >> (uint(int32(3)) % 32))
	v116 = int32(4)
	v117 = v106 + v116
	if v107&v116 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v219 = v210
	v220 = v211
	v222 = v195
	v225 = v200
	goto L31
L34:
	;
	v200 = int32(0)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v117+v115+(v200-v115)&int32(3)+v188<<(uint(int32(2))%32))))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	if base.Ui32(v211) < base.Ui32(int32(8)) {
		v219 = v210
		v220 = v211
		v222 = v195
		v225 = v200
		goto L31
	} else {
		goto L52
	}
L35:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v109))))
	v166 = int32(0)
	goto L46
L36:
	;
	v122 = int32(0)
	if base.Ui32(v83) <= base.Ui32(v109) {
		v155 = v109
		v158 = v122
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if v158 == v115 {
		v188 = v122
		v195 = v155
		goto L34
	} else {
		goto L44
	}
L38:
	;
	v132 = v109
	v135 = v122
	goto L39
L39:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v135))))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v132))))
	if v138 != v140 {
		v155 = v132
		v158 = v135
		goto L37
	} else {
		goto L41
	}
L40:
	;
	v155 = v143
	v158 = v145
	goto L37
L41:
	;
	v142 = int32(1)
	v143 = v132 + v142
	v145 = v135 + v142
	if base.Ui32(v115) <= base.Ui32(v145) {
		v155 = v143
		v158 = v145
		goto L37
	} else {
		goto L42
	}
L42:
	;
	if base.Ui32(v143) < base.Ui32(v83) {
		v132 = v143
		v135 = v145
		goto L39
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	v219 = v106
	v220 = v107
	v222 = v155
	v225 = v158
	goto L31
L45:
	;
	if v166 != v115 {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v166))))
	if v179 == v163&int32(255) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v181 = int32(1)
	v183 = v166 + v181
	if v183 != v115 {
		v166 = v183
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v235 = v106
	v236 = v107
	v238 = v109
	v242 = v181
	goto L27
L50:
	;
	v188 = v166
	v195 = v109 + int32(1)
	goto L34
L51:
	;
	v219 = v106
	v220 = v107
	v222 = v109
	v225 = v115
	goto L31
L52:
	;
	if base.Ui32(v195) < base.Ui32(v83) {
		v106 = v210
		v107 = v211
		v109 = v195
		goto L32
	} else {
		goto L53
	}
L53:
	;
	goto L33
L54:
	;
	goto L26
L55:
	;
	v244 = int32(0)
	if v236&int32(1) == v244 {
		v279 = v244
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v250 = v236 & int32(4)
	if v242&base.B2i32(v250 != int32(0)) != 0 {
		v279 = v244
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v254 = int32(1)
	if v85 == int32(0) {
		v279 = v254
		goto L54
	} else {
		goto L58
	}
L58:
	;
	if v236&int32(2) != 0 {
		v276 = int32(0)
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v276
	v279 = v254
	goto L54
L60:
	;
	v260 = int32(3)
	v261 = int32(base.Ui32(v236) >> (uint(v260) % 32))
	if v250 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v271 = int32(4)
	goto L63
L62:
	;
	v271 = v261 << (uint(int32(2)) % 32)
	goto L63
L63:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v235+v261+(int32(0)-v261)&v260+v271+int32(4))))
	v276 = v275
	goto L59
L64:
	;
	F_sdsfree(m, v11)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v324 = v285
	goto L24
L66:
	;
	v307 = int32(0)
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_ACLGetCommandID[1]))
	v310 = F_raxInsert(m, v287, v11, v306, v308, v307)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L72
	}
L67:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-17))))
	v306 = v305
	goto L66
L68:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-9))))
	v306 = v302
	goto L66
L69:
	;
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-5)))))
	v306 = v299
	goto L66
L70:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
	v306 = v296
	goto L66
L71:
	;
	v306 = int32(base.Ui32(v289) >> (uint(int32(3)) % 32))
	goto L66
L72:
	;
	F_sdsfree(m, v11)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v314 = int32(0)
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_ACLGetCommandID[1]))
	v319 = v317 + int32(1)
	if v319 == int32(1023) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v322 = int32(1024)
	goto L76
L75:
	;
	v322 = v319
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLGetCommandID[1])) = v322
	v324 = v317
	goto L24
}
func F_ACLHashPassword(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	v5 = m.G0
	v7 = v5 - int32(208)
	m.G0 = v7
	v10 = v7 + int32(96)
	F_sha256_init(m, v10)
	mBase = m.M
	F_sha256_update(m, v10, l0, l1)
	mBase = m.M
	F_sha256_final(m, v10, v7+int32(64))
	mBase = m.M
	v22 = int32(0)
	for {
		v25 = int32(1)
		v27 = v7 + v22<<(uint(v25)%32)
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(64)+v22))))
		v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31&int32(15))+uint32(_c_F_ACLHashPassword[0]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v36)
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v31)>>(uint(int32(4))%32)))+uint32(_c_F_ACLHashPassword[0]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v42)
		v45 = v22 + v25
		if v45 != int32(32) {
			v22 = v45
			continue
		} else {
			break
		}
		break
	}
	v49 = F_sdsnewlen(m, v7, int32(64))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(208)
		return v49
	}
}
func F_ACLInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int64
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	v7 = F_raxNew(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLInit[0])) = v7
	v11 = F_listCreate(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLInit[1])) = v11
	v16 = F_valkey_calloc(m, int32(1040))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLInit[2])) = v16
	v20 = *(*int64)(unsafe.Add(mBase, _c_F_ACLInit[3]))
	if v20 == int64(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F__serverAssert(m, int32(_a_F_ACLInit_0), int32(_a_F_ACLInit_1), int32(127))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L21
	}
L6:
	;
	v66 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_ACLInit[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = int32(11)
	v71 = F_listCreate(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ACLInit[4]))
	v27 = v20
	v28 = v24
	v29 = int32(_a_F_ACLInit_2)
	v30 = v23
	goto L8
L8:
	;
	if base.Ui32(int32(64)) <= base.Ui32(v28) {
		goto L5
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v35 = F_zstrdup(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v37 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_ACLInit[2]))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_ACLInit[4]))
	v41 = int32(4)
	v43 = v38 + v40<<(uint(v41)%32)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+8)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v35
	v47 = int32(1)
	v48 = v40 + v47
	*(*int32)(unsafe.Add(mBase, _c_F_ACLInit[4])) = v48
	v51 = v30 + v47
	v53 = v51 << (uint(v41) % 32)
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_ACLInit[3])))
	if v58 != int64(0) {
		v27 = v58
		v28 = v48
		v29 = v53 + int32(_a_F_ACLInit_2)
		v30 = v51
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLInit[5])) = v71
	v76 = F_ACLCreateUser(m, int32(_a_F_ACLInit_3), int32(7))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v80 = F_ACLSetUser(m, v76, int32(_a_F_ACLInit_4), int32(-1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v84 = F_ACLSetUser(m, v76, int32(_a_F_ACLInit_5), int32(-1))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v88 = F_ACLSetUser(m, v76, int32(_a_F_ACLInit_6), int32(-1))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v92 = F_ACLSetUser(m, v76, int32(_a_F_ACLInit_7), int32(-1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v96 = F_ACLSetUser(m, v76, int32(_a_F_ACLInit_8), int32(-1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v100 = F_ACLSetUser(m, v76, int32(_a_F_ACLInit_9), int32(-1))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLInit[6])) = v76
	return
L21:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ACLListFreeKeyPattern(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_sdsfree(m, v2)
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
func F_ACLListFreeSelector(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_ACLFreeSelector(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_ACLListMatchKeyPattern(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5 = int32(0)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3+int32(-1)))))
	switch v12 & int32(7) {
	case 0:
		v29 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
	case 1:
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3+int32(-3)))))
		v29 = v19
	case 2:
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3+int32(-5)))))
		v29 = v22
	case 3:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v3+int32(-9))))
		v29 = v25
	case 4:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v3+int32(-17))))
		v29 = v28
	default:
		v29 = v5
	}
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+int32(-1)))))
	switch v32 & int32(7) {
	case 0:
		v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	case 1:
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+int32(-3)))))
		v49 = v39
	case 2:
		v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4+int32(-5)))))
		v49 = v42
	case 3:
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v4+int32(-9))))
		v49 = v45
	case 4:
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v4+int32(-17))))
		v49 = v48
	default:
		v49 = v5
	}
	v50 = base.B2i32(base.Ui32(v29) < base.Ui32(v49))
	if base.Ui32(v29) < base.Ui32(v49) {
		v51 = v29
	} else {
		v51 = v49
	}
	v52 = F_memcmp(m, v3, v4, v51)
	mBase = m.M
	if v52 != 0 {
		v55 = v52
	} else {
		v55 = base.B2i32(base.Ui32(v49) < base.Ui32(v29)) - v50
	}
	return base.B2i32(v55 == int32(0))
}
func F_ACLLoadUsersAtStartup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v475 int32
	_ = v475
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v555 int32
	_ = v555
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	v1 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ACLLoadUsersAtStartup[0]))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ACLLoadUsersAtStartup[1]))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v17 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v38 = v11 + int32(52)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v39
	goto L8
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ACLLoadUsersAtStartup[2]))
	if int32(3) < v24 {
		goto L4
	} else {
		goto L5
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(_a_F_ACLLoadUsersAtStartup_0)
	F__serverLog(m, int32(3), int32(_a_F_ACLLoadUsersAtStartup_1), v11+int32(48))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	goto L4
L8:
	;
	v44 = v11 + int32(52)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v46 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_ACLLoadUsersAtStartup[2]))
	if int32(3) < v566 {
		goto L9
	} else {
		goto L136
	}
L11:
	;
	F__serverAssert(m, int32(_a_F_ACLLoadUsersAtStartup_2), int32(_a_F_ACLLoadUsersAtStartup_3), int32(2448))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L6
	} else {
		goto L135
	}
L12:
	;
	v524 = *(*int32)(unsafe.Add(mBase, _c_F_ACLLoadUsersAtStartup[1]))
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524))))
	if v525 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L13:
	;
	if v46 == int32(0) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46+base.B2i32(v49 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v55
	goto L14
L16:
	;
	v61 = v46
	goto L17
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-1)))))
	v74 = v72 & int32(7)
	switch v74 {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L23
	case 3:
		goto L22
	case 4:
		goto L21
	default:
		v148 = int32(0)
		goto L19
	}
L18:
	;
	goto L12
L19:
	;
	v155 = F_ACLCreateUser(m, v69, v148)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L42
	}
L20:
	;
	v90 = int32(0)
	if v89 == v90 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-17))))
	v89 = v88
	goto L20
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9))))
	v89 = v85
	goto L20
L23:
	;
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))))
	v89 = v82
	goto L20
L24:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))))
	v89 = v79
	goto L20
L25:
	;
	v89 = int32(base.Ui32(v72) >> (uint(int32(3)) % 32))
	goto L20
L26:
	;
	switch v74 {
	case 0:
		goto L40
	case 1:
		goto L39
	case 2:
		goto L38
	case 3:
		goto L37
	case 4:
		goto L36
	default:
		v148 = int32(0)
		goto L19
	}
L27:
	;
	v94 = v90
	goto L29
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_ACLLoadUsersAtStartup[2]))
	if int32(3) < v116 {
		goto L9
	} else {
		goto L34
	}
L29:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+v94))))
	if v102&int32(223) == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	if base.Ui32(base.I32_extend8_s(v102)+int32(-9)) <= base.Ui32(int32(4)) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v113 = v94 + int32(1)
	if v113 == v89 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v94 = v113
	goto L29
L34:
	;
	F__serverLog(m, int32(3), int32(_a_F_ACLLoadUsersAtStartup_4), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L10
L36:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-17))))
	v148 = v146
	goto L19
L37:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9))))
	v148 = v143
	goto L19
L38:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))))
	v148 = v140
	goto L19
L39:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))))
	v148 = v137
	goto L19
L40:
	;
	v148 = int32(base.Ui32(v72) >> (uint(int32(3)) % 32))
	goto L19
L41:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v398 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L42:
	;
	if v155 != 0 {
		v397 = v155
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v157 = int32(_a_F_ACLLoadUsersAtStartup_5)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ACLLoadUsersAtStartup[3])))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v161 == int32(0) {
		v184 = v160
		v185 = v161
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v185-v184&int32(255) != 0 {
		goto L11
	} else {
		goto L52
	}
L45:
	;
	goto L44
L46:
	;
	if v161 != v160&int32(255) {
		v184 = v160
		v185 = v161
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v167 = v69
	v168 = v157
	goto L48
L48:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	if v172 == int32(0) {
		v184 = v171
		v185 = v172
		goto L45
	} else {
		goto L50
	}
L49:
	;
	v184 = v171
	v185 = v172
	goto L45
L50:
	;
	v175 = int32(1)
	if v172 == v171&int32(255) {
		v167 = v167 + v175
		v168 = v168 + v175
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v189 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v189
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_ACLLoadUsersAtStartup[4]))
	v194 = int32(7)
	v196 = v11 + int32(60)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	goto L56
L53:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
	v395 = F_ACLSetUser(m, v392, int32(_a_F_ACLLoadUsersAtStartup_6), int32(-1))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L6
	} else {
		goto L91
	}
L54:
	;
	if v349 != v194 {
		goto L81
	} else {
		goto L82
	}
L55:
	;
	v340 = int32(0)
	v346 = v205
	v347 = v206
	v349 = v340
	v353 = v340
	goto L54
L56:
	;
	if base.Ui32(v206) < base.Ui32(int32(8)) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v217 = v205
	v218 = v206
	v220 = int32(0)
	goto L59
L58:
	;
	v346 = v330
	v347 = v331
	v349 = v333
	v353 = base.B2i32(v336 != int32(0))
	goto L54
L59:
	;
	v226 = int32(base.Ui32(v218) >> (uint(int32(3)) % 32))
	v227 = int32(4)
	v228 = v217 + v227
	if v218&v227 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v330 = v321
	v331 = v322
	v333 = v306
	v336 = v311
	goto L58
L61:
	;
	v311 = int32(0)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v228+v226+(v311-v226)&int32(3)+v299<<(uint(int32(2))%32))))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	if base.Ui32(v322) < base.Ui32(int32(8)) {
		v330 = v321
		v331 = v322
		v333 = v306
		v336 = v311
		goto L58
	} else {
		goto L79
	}
L62:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+uint32(_c_F_ACLLoadUsersAtStartup[3]))))
	v277 = int32(0)
	goto L73
L63:
	;
	v233 = int32(0)
	if base.Ui32(v194) <= base.Ui32(v220) {
		v266 = v220
		v269 = v233
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if v269 == v226 {
		v299 = v233
		v306 = v266
		goto L61
	} else {
		goto L71
	}
L65:
	;
	v243 = v220
	v246 = v233
	goto L66
L66:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228+v246))))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+uint32(_c_F_ACLLoadUsersAtStartup[3]))))
	if v249 != v251 {
		v266 = v243
		v269 = v246
		goto L64
	} else {
		goto L68
	}
L67:
	;
	v266 = v254
	v269 = v256
	goto L64
L68:
	;
	v253 = int32(1)
	v254 = v243 + v253
	v256 = v246 + v253
	if base.Ui32(v226) <= base.Ui32(v256) {
		v266 = v254
		v269 = v256
		goto L64
	} else {
		goto L69
	}
L69:
	;
	if base.Ui32(v254) < base.Ui32(v194) {
		v243 = v254
		v246 = v256
		goto L66
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	v330 = v217
	v331 = v218
	v333 = v266
	v336 = v269
	goto L58
L72:
	;
	if v277 != v226 {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228+v277))))
	if v290 == v274&int32(255) {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v292 = int32(1)
	v294 = v277 + v292
	if v294 != v226 {
		v277 = v294
		goto L73
	} else {
		goto L76
	}
L76:
	;
	v346 = v217
	v347 = v218
	v349 = v220
	v353 = v292
	goto L54
L77:
	;
	v299 = v277
	v306 = v220 + int32(1)
	goto L61
L78:
	;
	v330 = v217
	v331 = v218
	v333 = v220
	v336 = v226
	goto L58
L79:
	;
	if base.Ui32(v306) < base.Ui32(v194) {
		v217 = v321
		v218 = v322
		v220 = v306
		goto L59
	} else {
		goto L80
	}
L80:
	;
	goto L60
L81:
	;
	goto L53
L82:
	;
	if v347&int32(1) == int32(0) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v361 = v347 & int32(4)
	if v353&base.B2i32(v361 != int32(0)) != 0 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	if v196 == int32(0) {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	if v347&int32(2) != 0 {
		v387 = int32(0)
		goto L86
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v387
	goto L81
L87:
	;
	v371 = int32(3)
	v372 = int32(base.Ui32(v347) >> (uint(v371) % 32))
	if v361 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v382 = int32(4)
	goto L90
L89:
	;
	v382 = v372 << (uint(int32(2)) % 32)
	goto L90
L90:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v346+v372+(int32(0)-v372)&v371+v382+int32(4))))
	v387 = v386
	goto L86
L91:
	;
	v397 = v392
	goto L41
L92:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+4)))
	if v484&int32(2) == int32(0) {
		goto L119
	} else {
		goto L120
	}
L93:
	;
	v405 = v398
	v408 = int32(1)
	v409 = v68 + int32(4)
	goto L94
L94:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405+int32(-1)))))
	switch v415 & int32(7) {
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
		v432 = int32(0)
		goto L96
	}
L95:
	;
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_ACLLoadUsersAtStartup[2]))
	if int32(3) < v444 {
		goto L9
	} else {
		goto L106
	}
L96:
	;
	v433 = F_ACLSetUser(m, v397, v405, v432)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L6
	} else {
		goto L103
	}
L97:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v405+int32(-17))))
	v432 = v431
	goto L96
L98:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v405+int32(-9))))
	v432 = v428
	goto L96
L99:
	;
	v425 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v405+int32(-5)))))
	v432 = v425
	goto L96
L100:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405+int32(-3)))))
	v432 = v422
	goto L96
L101:
	;
	v432 = int32(base.Ui32(v415) >> (uint(int32(3)) % 32))
	goto L96
L102:
	;
	goto L95
L103:
	;
	if v433 != 0 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v436 = v408 + int32(1)
	v439 = v68 + v436<<(uint(int32(2))%32)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v439)))
	if v440 == int32(0) {
		goto L92
	} else {
		goto L105
	}
L105:
	;
	v405 = v440
	v408 = v436
	v409 = v439
	goto L94
L106:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v451 = F___errno_location(m)
	mBase = m.M
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	switch v452 + int32(-7) {
	case 0:
		goto L112
	default:
		goto L109
	case 2:
		goto L113
	case 5:
		goto L111
	case 13:
		goto L116
	case 21:
		goto L117
	case 24:
		goto L115
	case 36:
		goto L114
	case 37:
		v464 = int32(_a_F_ACLLoadUsersAtStartup_7)
		goto L108
	case 61:
		goto L110
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v447
	F__serverLog(m, int32(3), int32(_a_F_ACLLoadUsersAtStartup_8), v11+int32(32))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L6
	} else {
		goto L118
	}
L108:
	;
	v466 = v464
	goto L107
L109:
	;
	v464 = int32(_a_F_ACLLoadUsersAtStartup_9)
	goto L108
L110:
	;
	v466 = int32(_a_F_ACLLoadUsersAtStartup_10)
	goto L107
L111:
	;
	v466 = int32(_a_F_ACLLoadUsersAtStartup_11)
	goto L107
L112:
	;
	v466 = int32(_a_F_ACLLoadUsersAtStartup_12)
	goto L107
L113:
	;
	v466 = int32(_a_F_ACLLoadUsersAtStartup_13)
	goto L107
L114:
	;
	v466 = int32(_a_F_ACLLoadUsersAtStartup_14)
	goto L107
L115:
	;
	v466 = int32(_a_F_ACLLoadUsersAtStartup_15)
	goto L107
L116:
	;
	v466 = int32(_a_F_ACLLoadUsersAtStartup_16)
	goto L107
L117:
	;
	v466 = int32(_a_F_ACLLoadUsersAtStartup_17)
	goto L107
L118:
	;
	goto L10
L119:
	;
	v502 = v11 + int32(52)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	if v504 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L120:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _c_F_ACLLoadUsersAtStartup[2]))
	if int32(2) < v490 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v493
	F__serverLog(m, int32(2), int32(_a_F_ACLLoadUsersAtStartup_18), v11+int32(16))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	goto L119
L123:
	;
	if v504 != 0 {
		v61 = v504
		goto L17
	} else {
		goto L126
	}
L124:
	;
	goto L123
L125:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v502)+4))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v504+base.B2i32(v507 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v502))) = v513
	goto L124
L126:
	;
	goto L18
L127:
	;
	m.G0 = v11 + int32(64)
	return
L128:
	;
	v528 = F_ACLLoadFromFile(m, v524)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	if v528 == int32(0) {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_ACLLoadUsersAtStartup[2]))
	if int32(3) < v533 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	F_sdsfree(m, v528)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L6
	} else {
		goto L134
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_ACLLoadUsersAtStartup_0)
	F__serverLog(m, int32(3), int32(_a_F_ACLLoadUsersAtStartup_19), v11)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F__serverLog(m, int32(3), int32(_a_F_ACLLoadUsersAtStartup_20), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	goto L9
}
func F_ACLLookupCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = F_sdsnew(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_ACLLookupCommand[0]))
		v9 = F_lookupCommandBySdsLogic(m, v8, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_sdsfree(m, v3)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v9
			}
		}
	}
}
func F_ACLMergeSelectorArguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v5
	v22 = F_valkey_malloc(m, l1<<(uint(int32(2))%32))
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
	if l1 <= int32(0) {
		v199 = v22
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v15 + int32(16)
	return v199
L4:
	;
	v35 = v5
	v37 = int32(-1)
	v38 = int32(0)
	goto L5
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+v35<<(uint(int32(2))%32))))
	if v37 != int32(-1) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	if v137 == int32(-1) {
		v199 = v22
		goto L3
	} else {
		goto L31
	}
L7:
	;
	v143 = v35 + int32(1)
	if v143 != l1 {
		v35 = v143
		v37 = v137
		v38 = v138
		goto L5
	} else {
		goto L30
	}
L8:
	;
	v137 = int32(-1)
	v138 = v132
	goto L7
L9:
	;
	v121 = F_sdsdup(m, v45)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L29
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v45
	v82 = F_sdscatfmt(m, v38, int32(_a_F_ACLMergeSelectorArguments_0), v15)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L11:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v48 != int32(40) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+int32(-1)))))
	switch v54 & int32(7) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	case 4:
		goto L14
	default:
		v71 = int32(0)
		goto L13
	}
L13:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v71+int32(-1)))))
	if v75 == int32(41) {
		goto L9
	} else {
		goto L19
	}
L14:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(-17))))
	v71 = v70
	goto L13
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(-9))))
	v71 = v67
	goto L13
L16:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45+int32(-5)))))
	v71 = v64
	goto L13
L17:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+int32(-3)))))
	v71 = v61
	goto L13
L18:
	;
	v71 = int32(base.Ui32(v54) >> (uint(int32(3)) % 32))
	goto L13
L19:
	;
	v78 = F_sdsdup(m, v45)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v137 = v35
	v138 = v78
	goto L7
L21:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+int32(-1)))))
	switch v87 & int32(7) {
	case 0:
		goto L27
	case 1:
		goto L26
	case 2:
		goto L25
	case 3:
		goto L24
	case 4:
		goto L23
	default:
		v104 = int32(0)
		goto L22
	}
L22:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v104+int32(-1)))))
	if v108 != int32(41) {
		v137 = v37
		v138 = v82
		goto L7
	} else {
		goto L28
	}
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(-17))))
	v104 = v103
	goto L22
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(-9))))
	v104 = v100
	goto L22
L25:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45+int32(-5)))))
	v104 = v97
	goto L22
L26:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+int32(-3)))))
	v104 = v94
	goto L22
L27:
	;
	v104 = int32(base.Ui32(v87) >> (uint(int32(3)) % 32))
	goto L22
L28:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v22+v111<<(uint(int32(2))%32)))) = v82
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v111 + int32(1)
	v132 = v82
	goto L8
L29:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v22+v123<<(uint(int32(2))%32)))) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v123 + int32(1)
	v132 = v38
	goto L8
L30:
	;
	goto L6
L31:
	;
	v147 = int32(0)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v148 <= v147 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_valkey_free(m, v22)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L38
	}
L33:
	;
	v160 = v147
	goto L34
L34:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v22+v160<<(uint(int32(2))%32))))
	F_sdsfree(m, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L32
L36:
	;
	v170 = v160 + int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v170 < v171 {
		v160 = v170
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	F_sdsfree(m, v138)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v189 = int32(0)
	if l3 == v189 {
		v199 = v189
		goto L3
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v137
	v199 = v189
	goto L3
}
func F_ACLSetSelector(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
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
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
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
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v740 int32
	_ = v740
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v773 int32
	_ = v773
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v798 int32
	_ = v798
	var v809 int32
	_ = v809
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v988 int32
	_ = v988
	var v999 int32
	_ = v999
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1053 int32
	_ = v1053
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(_a_F_ACLSetSelector_0)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v1053
L2:
	;
	v104 = int32(_a_F_ACLSetSelector_1)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v107 != 0 {
		goto L35
	} else {
		goto L36
	}
L3:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v94 | int32(2)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	F_listEmpty(m, v98)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L30
	} else {
		goto L31
	}
L4:
	;
	if v49-v51 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L5:
	;
	v49 = F_tolower(m, v45)
	mBase = m.M
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	v51 = F_tolower(m, v50)
	mBase = m.M
	goto L4
L6:
	;
	v19 = l1
	v20 = v14
	v21 = v17
	goto L9
L7:
	;
	v45 = int32(0)
	v46 = v14
	goto L5
L8:
	;
	v45 = v42 & int32(255)
	v46 = v41
	goto L5
L9:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v23 == int32(0) {
		v41 = v20
		v42 = v21
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v41 = v35
	v42 = int32(0)
	goto L8
L11:
	;
	v27 = v21 & int32(255)
	if v27 == v23 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v34 = int32(1)
	v35 = v20 + v34
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v36 != 0 {
		v19 = v19 + v34
		v20 = v35
		v21 = v36
		goto L9
	} else {
		goto L15
	}
L13:
	;
	v29 = F_tolower(m, v27)
	mBase = m.M
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v31 = F_tolower(m, v30)
	mBase = m.M
	if v29 == v31 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v41 = v20
	v42 = v33
	goto L8
L15:
	;
	goto L10
L16:
	;
	v55 = int32(_a_F_ACLSetSelector_2)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v58 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if v90-v92 != 0 {
		goto L2
	} else {
		goto L29
	}
L18:
	;
	v90 = F_tolower(m, v86)
	mBase = m.M
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v92 = F_tolower(m, v91)
	mBase = m.M
	goto L17
L19:
	;
	v60 = l1
	v61 = v55
	v62 = v58
	goto L22
L20:
	;
	v86 = int32(0)
	v87 = v55
	goto L18
L21:
	;
	v86 = v83 & int32(255)
	v87 = v82
	goto L18
L22:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v64 == int32(0) {
		v82 = v61
		v83 = v62
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v82 = v76
	v83 = int32(0)
	goto L21
L24:
	;
	v68 = v62 & int32(255)
	if v68 == v64 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v75 = int32(1)
	v76 = v61 + v75
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	if v77 != 0 {
		v60 = v60 + v75
		v61 = v76
		v62 = v77
		goto L22
	} else {
		goto L28
	}
L26:
	;
	v70 = F_tolower(m, v68)
	mBase = m.M
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v72 = F_tolower(m, v71)
	mBase = m.M
	if v70 == v72 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	v82 = v61
	v83 = v74
	goto L21
L28:
	;
	goto L23
L29:
	;
	goto L3
L30:
	;
	return int32(0)
L31:
	;
	v1053 = int32(0)
	goto L1
L32:
	;
	v151 = int32(_a_F_ACLSetSelector_3)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v154 != 0 {
		goto L51
	} else {
		goto L52
	}
L33:
	;
	if v139-v141 != 0 {
		goto L32
	} else {
		goto L45
	}
L34:
	;
	v139 = F_tolower(m, v135)
	mBase = m.M
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v141 = F_tolower(m, v140)
	mBase = m.M
	goto L33
L35:
	;
	v109 = l1
	v110 = v104
	v111 = v107
	goto L38
L36:
	;
	v135 = int32(0)
	v136 = v104
	goto L34
L37:
	;
	v135 = v132 & int32(255)
	v136 = v131
	goto L34
L38:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v113 == int32(0) {
		v131 = v110
		v132 = v111
		goto L37
	} else {
		goto L40
	}
L39:
	;
	v131 = v125
	v132 = int32(0)
	goto L37
L40:
	;
	v117 = v111 & int32(255)
	if v117 == v113 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v124 = int32(1)
	v125 = v110 + v124
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
	if v126 != 0 {
		v109 = v109 + v124
		v110 = v125
		v111 = v126
		goto L38
	} else {
		goto L44
	}
L42:
	;
	v119 = F_tolower(m, v117)
	mBase = m.M
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v121 = F_tolower(m, v120)
	mBase = m.M
	if v119 == v121 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v131 = v110
	v132 = v123
	goto L37
L44:
	;
	goto L39
L45:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v143 & int32(-3)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	F_listEmpty(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L30
	} else {
		goto L46
	}
L46:
	;
	v1053 = int32(0)
	goto L1
L47:
	;
	v239 = int32(_a_F_ACLSetSelector_4)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v242 != 0 {
		goto L79
	} else {
		goto L80
	}
L48:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v231 | int32(8)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	F_listEmpty(m, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L30
	} else {
		goto L75
	}
L49:
	;
	if v186-v188 == int32(0) {
		goto L48
	} else {
		goto L61
	}
L50:
	;
	v186 = F_tolower(m, v182)
	mBase = m.M
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	v188 = F_tolower(m, v187)
	mBase = m.M
	goto L49
L51:
	;
	v156 = l1
	v157 = v151
	v158 = v154
	goto L54
L52:
	;
	v182 = int32(0)
	v183 = v151
	goto L50
L53:
	;
	v182 = v179 & int32(255)
	v183 = v178
	goto L50
L54:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v160 == int32(0) {
		v178 = v157
		v179 = v158
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v178 = v172
	v179 = int32(0)
	goto L53
L56:
	;
	v164 = v158 & int32(255)
	if v164 == v160 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v171 = int32(1)
	v172 = v157 + v171
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	if v173 != 0 {
		v156 = v156 + v171
		v157 = v172
		v158 = v173
		goto L54
	} else {
		goto L60
	}
L58:
	;
	v166 = F_tolower(m, v164)
	mBase = m.M
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	v168 = F_tolower(m, v167)
	mBase = m.M
	if v166 == v168 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	v178 = v157
	v179 = v170
	goto L53
L60:
	;
	goto L55
L61:
	;
	v192 = int32(_a_F_ACLSetSelector_5)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v195 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	if v227-v229 != 0 {
		goto L47
	} else {
		goto L74
	}
L63:
	;
	v227 = F_tolower(m, v223)
	mBase = m.M
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	v229 = F_tolower(m, v228)
	mBase = m.M
	goto L62
L64:
	;
	v197 = l1
	v198 = v192
	v199 = v195
	goto L67
L65:
	;
	v223 = int32(0)
	v224 = v192
	goto L63
L66:
	;
	v223 = v220 & int32(255)
	v224 = v219
	goto L63
L67:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if v201 == int32(0) {
		v219 = v198
		v220 = v199
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v219 = v213
	v220 = int32(0)
	goto L66
L69:
	;
	v205 = v199 & int32(255)
	if v205 == v201 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v212 = int32(1)
	v213 = v198 + v212
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)))
	if v214 != 0 {
		v197 = v197 + v212
		v198 = v213
		v199 = v214
		goto L67
	} else {
		goto L73
	}
L71:
	;
	v207 = F_tolower(m, v205)
	mBase = m.M
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	v209 = F_tolower(m, v208)
	mBase = m.M
	if v207 == v209 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v219 = v198
	v220 = v211
	goto L66
L73:
	;
	goto L68
L74:
	;
	goto L48
L75:
	;
	v1053 = int32(0)
	goto L1
L76:
	;
	v286 = int32(_a_F_ACLSetSelector_6)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v289 != 0 {
		goto L97
	} else {
		goto L98
	}
L77:
	;
	if v274-v276 != 0 {
		goto L76
	} else {
		goto L89
	}
L78:
	;
	v274 = F_tolower(m, v270)
	mBase = m.M
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	v276 = F_tolower(m, v275)
	mBase = m.M
	goto L77
L79:
	;
	v244 = l1
	v245 = v239
	v246 = v242
	goto L82
L80:
	;
	v270 = int32(0)
	v271 = v239
	goto L78
L81:
	;
	v270 = v267 & int32(255)
	v271 = v266
	goto L78
L82:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if v248 == int32(0) {
		v266 = v245
		v267 = v246
		goto L81
	} else {
		goto L84
	}
L83:
	;
	v266 = v260
	v267 = int32(0)
	goto L81
L84:
	;
	v252 = v246 & int32(255)
	if v252 == v248 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v259 = int32(1)
	v260 = v245 + v259
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+1)))
	if v261 != 0 {
		v244 = v244 + v259
		v245 = v260
		v246 = v261
		goto L82
	} else {
		goto L88
	}
L86:
	;
	v254 = F_tolower(m, v252)
	mBase = m.M
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	v256 = F_tolower(m, v255)
	mBase = m.M
	if v254 == v256 {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v266 = v245
	v267 = v258
	goto L81
L88:
	;
	goto L83
L89:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v278 & int32(-9)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	F_listEmpty(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L30
	} else {
		goto L90
	}
L90:
	;
	v1053 = int32(0)
	goto L1
L91:
	;
	v1053 = int32(-1)
	goto L1
L92:
	;
	v974 = l2 - v970
	v975 = l1 + v970
	if l2 == v970 {
		goto L324
	} else {
		goto L325
	}
L93:
	;
	v1053 = int32(0)
	goto L1
L94:
	;
	v338 = int32(_a_F_ACLSetSelector_7)
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v341 != 0 {
		goto L114
	} else {
		goto L115
	}
L95:
	;
	if v321-v323 != 0 {
		goto L94
	} else {
		goto L107
	}
L96:
	;
	v321 = F_tolower(m, v317)
	mBase = m.M
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	v323 = F_tolower(m, v322)
	mBase = m.M
	goto L95
L97:
	;
	v291 = l1
	v292 = v286
	v293 = v289
	goto L100
L98:
	;
	v317 = int32(0)
	v318 = v286
	goto L96
L99:
	;
	v317 = v314 & int32(255)
	v318 = v313
	goto L96
L100:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	if v295 == int32(0) {
		v313 = v292
		v314 = v293
		goto L99
	} else {
		goto L102
	}
L101:
	;
	v313 = v307
	v314 = int32(0)
	goto L99
L102:
	;
	v299 = v293 & int32(255)
	if v299 == v295 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v306 = int32(1)
	v307 = v292 + v306
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+1)))
	if v308 != 0 {
		v291 = v291 + v306
		v292 = v307
		v293 = v308
		goto L100
	} else {
		goto L106
	}
L104:
	;
	v301 = F_tolower(m, v299)
	mBase = m.M
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	v303 = F_tolower(m, v302)
	mBase = m.M
	if v301 == v303 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	v313 = v292
	v314 = v305
	goto L99
L106:
	;
	goto L101
L107:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v325 | int32(16)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v329 == int32(0) {
		goto L93
	} else {
		goto L108
	}
L108:
	;
	F_intsetFree(m, v329)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L30
	} else {
		goto L109
	}
L109:
	;
	v334 = F_intsetNew(m)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L30
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v334
	v1053 = int32(0)
	goto L1
L111:
	;
	v390 = int32(_a_F_ACLSetSelector_8)
	goto L130
L112:
	;
	if v373-v375 != 0 {
		goto L111
	} else {
		goto L124
	}
L113:
	;
	v373 = F_tolower(m, v369)
	mBase = m.M
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	v375 = F_tolower(m, v374)
	mBase = m.M
	goto L112
L114:
	;
	v343 = l1
	v344 = v338
	v345 = v341
	goto L117
L115:
	;
	v369 = int32(0)
	v370 = v338
	goto L113
L116:
	;
	v369 = v366 & int32(255)
	v370 = v365
	goto L113
L117:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	if v347 == int32(0) {
		v365 = v344
		v366 = v345
		goto L116
	} else {
		goto L119
	}
L118:
	;
	v365 = v359
	v366 = int32(0)
	goto L116
L119:
	;
	v351 = v345 & int32(255)
	if v351 == v347 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v358 = int32(1)
	v359 = v344 + v358
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+1)))
	if v360 != 0 {
		v343 = v343 + v358
		v344 = v359
		v345 = v360
		goto L117
	} else {
		goto L123
	}
L121:
	;
	v353 = F_tolower(m, v351)
	mBase = m.M
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	v355 = F_tolower(m, v354)
	mBase = m.M
	if v353 == v355 {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343))))
	v365 = v344
	v366 = v357
	goto L116
L123:
	;
	goto L118
L124:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v377 & int32(-17)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v381 == int32(0) {
		goto L93
	} else {
		goto L125
	}
L125:
	;
	F_intsetFree(m, v381)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L30
	} else {
		goto L126
	}
L126:
	;
	v386 = F_intsetNew(m)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L30
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v386
	v1053 = int32(0)
	goto L1
L128:
	;
	v451 = int32(_a_F_ACLSetSelector_9)
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v454 != 0 {
		goto L151
	} else {
		goto L152
	}
L129:
	;
	if v434-v436 != 0 {
		goto L128
	} else {
		goto L144
	}
L130:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v395 != 0 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v434 = F_tolower(m, v429)
	mBase = m.M
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	v436 = F_tolower(m, v435)
	mBase = m.M
	goto L129
L133:
	;
	v397 = l1
	v398 = v390
	v399 = int32(3)
	v400 = v395
	goto L136
L134:
	;
	v429 = int32(0)
	v430 = v390
	goto L132
L135:
	;
	v429 = v426 & int32(255)
	v430 = v424
	goto L132
L136:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398))))
	if v402 == int32(0) {
		v424 = v398
		v426 = v400
		goto L135
	} else {
		goto L138
	}
L137:
	;
	v424 = v418
	v426 = int32(0)
	goto L135
L138:
	;
	v406 = v399 + int32(-1)
	if v406 == int32(0) {
		v424 = v398
		v426 = v400
		goto L135
	} else {
		goto L139
	}
L139:
	;
	v410 = v400 & int32(255)
	if v410 == v402 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v417 = int32(1)
	v418 = v398 + v417
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+1)))
	if v419 != 0 {
		v397 = v397 + v417
		v398 = v418
		v399 = v406
		v400 = v419
		goto L136
	} else {
		goto L143
	}
L141:
	;
	v412 = F_tolower(m, v410)
	mBase = m.M
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398))))
	v414 = F_tolower(m, v413)
	mBase = m.M
	if v412 == v414 {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397))))
	v424 = v398
	v426 = v416
	goto L135
L143:
	;
	goto L137
L144:
	;
	v447 = F_ACLSetSelectorDatabasePermissions(m, l0, l1+int32(3))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L30
	} else {
		goto L145
	}
L145:
	;
	if v447 == int32(0) {
		goto L93
	} else {
		goto L146
	}
L146:
	;
	v1053 = int32(-1)
	goto L1
L147:
	;
	v571 = int32(_a_F_ACLSetSelector_10)
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v574 != 0 {
		goto L188
	} else {
		goto L189
	}
L148:
	;
	v536 = F__emscripten_memset_bulkmem(m, l0+int32(8), base.I32_extend8_s(int32(255)), int32(128))
	mBase = m.M
	goto L175
L149:
	;
	if v486-v488 == int32(0) {
		goto L148
	} else {
		goto L161
	}
L150:
	;
	v486 = F_tolower(m, v482)
	mBase = m.M
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
	v488 = F_tolower(m, v487)
	mBase = m.M
	goto L149
L151:
	;
	v456 = l1
	v457 = v451
	v458 = v454
	goto L154
L152:
	;
	v482 = int32(0)
	v483 = v451
	goto L150
L153:
	;
	v482 = v479 & int32(255)
	v483 = v478
	goto L150
L154:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457))))
	if v460 == int32(0) {
		v478 = v457
		v479 = v458
		goto L153
	} else {
		goto L156
	}
L155:
	;
	v478 = v472
	v479 = int32(0)
	goto L153
L156:
	;
	v464 = v458 & int32(255)
	if v464 == v460 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v471 = int32(1)
	v472 = v457 + v471
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456)+1)))
	if v473 != 0 {
		v456 = v456 + v471
		v457 = v472
		v458 = v473
		goto L154
	} else {
		goto L160
	}
L158:
	;
	v466 = F_tolower(m, v464)
	mBase = m.M
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457))))
	v468 = F_tolower(m, v467)
	mBase = m.M
	if v466 == v468 {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	v478 = v457
	v479 = v470
	goto L153
L160:
	;
	goto L155
L161:
	;
	v492 = int32(_a_F_ACLSetSelector_11)
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v495 != 0 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	if v527-v529 != 0 {
		goto L147
	} else {
		goto L174
	}
L163:
	;
	v527 = F_tolower(m, v523)
	mBase = m.M
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524))))
	v529 = F_tolower(m, v528)
	mBase = m.M
	goto L162
L164:
	;
	v497 = l1
	v498 = v492
	v499 = v495
	goto L167
L165:
	;
	v523 = int32(0)
	v524 = v492
	goto L163
L166:
	;
	v523 = v520 & int32(255)
	v524 = v519
	goto L163
L167:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498))))
	if v501 == int32(0) {
		v519 = v498
		v520 = v499
		goto L166
	} else {
		goto L169
	}
L168:
	;
	v519 = v513
	v520 = int32(0)
	goto L166
L169:
	;
	v505 = v499 & int32(255)
	if v505 == v501 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v512 = int32(1)
	v513 = v498 + v512
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+1)))
	if v514 != 0 {
		v497 = v497 + v512
		v498 = v513
		v499 = v514
		goto L167
	} else {
		goto L173
	}
L171:
	;
	v507 = F_tolower(m, v505)
	mBase = m.M
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498))))
	v509 = F_tolower(m, v508)
	mBase = m.M
	if v507 == v509 {
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	v519 = v498
	v520 = v511
	goto L166
L173:
	;
	goto L168
L174:
	;
	goto L148
L175:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v537 | int32(4)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v544 = v541 + int32(-1)
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544))))
	switch v545 & int32(7) {
	case 0:
		goto L182
	case 1:
		goto L181
	case 2:
		goto L180
	case 3:
		goto L179
	case 4:
		goto L178
	default:
		goto L177
	}
L176:
	;
	F_ACLResetFirstArgs(m, l0)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L30
	} else {
		goto L183
	}
L177:
	;
	v566 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v541))) = uint8(v566)
	goto L176
L178:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v541+int32(-17)))) = int64(0)
	goto L177
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v541+int32(-9)))) = int32(0)
	goto L177
L180:
	;
	v556 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v541+int32(-5)))) = uint16(v556)
	goto L177
L181:
	;
	v552 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v541+int32(-3)))) = uint8(v552)
	goto L177
L182:
	;
	v548 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v544))) = uint8(v548)
	goto L177
L183:
	;
	v1053 = int32(0)
	goto L1
L184:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	switch v691 + int32(-37) {
	case 0:
		goto L226
	case 1:
		goto L225
	case 2, 3, 4, 5, 7:
		goto L221
	case 6:
		goto L224
	case 8:
		goto L223
	default:
		goto L227
	}
L185:
	;
	v656 = F__emscripten_memset_bulkmem(m, l0+int32(8), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L212
L186:
	;
	if v606-v608 == int32(0) {
		goto L185
	} else {
		goto L198
	}
L187:
	;
	v606 = F_tolower(m, v602)
	mBase = m.M
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603))))
	v608 = F_tolower(m, v607)
	mBase = m.M
	goto L186
L188:
	;
	v576 = l1
	v577 = v571
	v578 = v574
	goto L191
L189:
	;
	v602 = int32(0)
	v603 = v571
	goto L187
L190:
	;
	v602 = v599 & int32(255)
	v603 = v598
	goto L187
L191:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577))))
	if v580 == int32(0) {
		v598 = v577
		v599 = v578
		goto L190
	} else {
		goto L193
	}
L192:
	;
	v598 = v592
	v599 = int32(0)
	goto L190
L193:
	;
	v584 = v578 & int32(255)
	if v584 == v580 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v591 = int32(1)
	v592 = v577 + v591
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+1)))
	if v593 != 0 {
		v576 = v576 + v591
		v577 = v592
		v578 = v593
		goto L191
	} else {
		goto L197
	}
L195:
	;
	v586 = F_tolower(m, v584)
	mBase = m.M
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577))))
	v588 = F_tolower(m, v587)
	mBase = m.M
	if v586 == v588 {
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576))))
	v598 = v577
	v599 = v590
	goto L190
L197:
	;
	goto L192
L198:
	;
	v612 = int32(_a_F_ACLSetSelector_12)
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v615 != 0 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	if v647-v649 != 0 {
		goto L184
	} else {
		goto L211
	}
L200:
	;
	v647 = F_tolower(m, v643)
	mBase = m.M
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644))))
	v649 = F_tolower(m, v648)
	mBase = m.M
	goto L199
L201:
	;
	v617 = l1
	v618 = v612
	v619 = v615
	goto L204
L202:
	;
	v643 = int32(0)
	v644 = v612
	goto L200
L203:
	;
	v643 = v640 & int32(255)
	v644 = v639
	goto L200
L204:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618))))
	if v621 == int32(0) {
		v639 = v618
		v640 = v619
		goto L203
	} else {
		goto L206
	}
L205:
	;
	v639 = v633
	v640 = int32(0)
	goto L203
L206:
	;
	v625 = v619 & int32(255)
	if v625 == v621 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v632 = int32(1)
	v633 = v618 + v632
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+1)))
	if v634 != 0 {
		v617 = v617 + v632
		v618 = v633
		v619 = v634
		goto L204
	} else {
		goto L210
	}
L208:
	;
	v627 = F_tolower(m, v625)
	mBase = m.M
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618))))
	v629 = F_tolower(m, v628)
	mBase = m.M
	if v627 == v629 {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617))))
	v639 = v618
	v640 = v631
	goto L203
L210:
	;
	goto L205
L211:
	;
	goto L185
L212:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v657 & int32(-5)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v664 = v661 + int32(-1)
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	switch v665 & int32(7) {
	case 0:
		goto L219
	case 1:
		goto L218
	case 2:
		goto L217
	case 3:
		goto L216
	case 4:
		goto L215
	default:
		goto L214
	}
L213:
	;
	F_ACLResetFirstArgs(m, l0)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L30
	} else {
		goto L220
	}
L214:
	;
	v686 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v661))) = uint8(v686)
	goto L213
L215:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v661+int32(-17)))) = int64(0)
	goto L214
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v661+int32(-9)))) = int32(0)
	goto L214
L217:
	;
	v676 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v661+int32(-5)))) = uint16(v676)
	goto L214
L218:
	;
	v672 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v661+int32(-3)))) = uint8(v672)
	goto L214
L219:
	;
	v668 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v664))) = uint8(v668)
	goto L214
L220:
	;
	v1053 = int32(0)
	goto L1
L221:
	;
	goto L323
L222:
	;
	v951 = F_ACLSetSelectorCategory(m, l0, l1+int32(1), base.B2i32(v691 == int32(43)))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L30
	} else {
		goto L320
	}
L223:
	;
	v927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v927 == int32(64) {
		goto L222
	} else {
		goto L313
	}
L224:
	;
	v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v838 == int32(64) {
		goto L222
	} else {
		goto L273
	}
L225:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v773&int32(8) == int32(0) {
		goto L254
	} else {
		goto L255
	}
L226:
	;
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v696&int32(2) == int32(0) {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	if v691 != int32(126) {
		goto L221
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	if v691 == int32(37) {
		goto L232
	} else {
		goto L233
	}
L230:
	;
	goto L231
L231:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[0])) = int32(20)
	goto L91
L232:
	;
	if base.Ui32(l2) < base.Ui32(int32(2)) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	v969 = int32(3)
	v970 = int32(1)
	goto L92
L234:
	;
	goto L253
L235:
	;
	v716 = int32(0)
	v718 = int32(1)
	goto L237
L236:
	;
	if v755 == int32(0) {
		goto L234
	} else {
		goto L251
	}
L237:
	;
	v722 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+v718))))
	if base.Ui32(v722+int32(-97)) < base.Ui32(int32(26)) {
		goto L242
	} else {
		goto L243
	}
L238:
	;
	v755 = v750
	v756 = l2
	v758 = v751
	goto L236
L239:
	;
	v750 = v716 | v749
	v751 = int32(1)
	v753 = v718 + v751
	if v753 != l2 {
		v716 = v750
		v718 = v753
		goto L237
	} else {
		goto L250
	}
L240:
	;
	if v729 != int32(87) {
		goto L247
	} else {
		goto L248
	}
L241:
	;
	if v729 != int32(82) {
		goto L240
	} else {
		goto L245
	}
L242:
	;
	v729 = v722 & int32(95)
	goto L244
L243:
	;
	v729 = v722
	goto L244
L244:
	;
	goto L241
L245:
	;
	v732 = int32(1)
	if v716&v732 == int32(0) {
		v749 = v732
		goto L239
	} else {
		goto L246
	}
L246:
	;
	goto L240
L247:
	;
	v747 = base.B2i32(v722 == int32(126))
	v755 = v716
	v756 = v718 + v747
	v758 = v747
	goto L236
L248:
	;
	v740 = int32(2)
	if v716&v740 == int32(0) {
		v749 = v740
		goto L239
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	goto L238
L251:
	;
	if v758 != 0 {
		v969 = v755
		v970 = v756
		goto L92
	} else {
		goto L252
	}
L252:
	;
	goto L234
L253:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[0])) = int32(28)
	goto L91
L254:
	;
	v782 = l1 + int32(1)
	v784 = l2 + int32(-1)
	if v784 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L255:
	;
	goto L256
L256:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[0])) = int32(31)
	goto L91
L257:
	;
	v823 = F_sdsnewlen(m, v782, v784)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L30
	} else {
		goto L266
	}
L258:
	;
	v789 = int32(0)
	goto L260
L259:
	;
	goto L265
L260:
	;
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v782+v789))))
	if v798&int32(223) == int32(0) {
		goto L259
	} else {
		goto L262
	}
L262:
	;
	if base.Ui32(base.I32_extend8_s(v798)+int32(-9)) <= base.Ui32(int32(4)) {
		goto L259
	} else {
		goto L263
	}
L263:
	;
	v809 = v789 + int32(1)
	if v809 == v784 {
		goto L257
	} else {
		goto L264
	}
L264:
	;
	v789 = v809
	goto L260
L265:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[0])) = int32(28)
	goto L91
L266:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v826 = F_listSearchKey(m, v825, v823)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L30
	} else {
		goto L269
	}
L267:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v833 & int32(-9)
	v1053 = int32(0)
	goto L1
L268:
	;
	F_sdsfree(m, v823)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L30
	} else {
		goto L272
	}
L269:
	;
	if v826 != 0 {
		goto L268
	} else {
		goto L270
	}
L270:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v829 = F_listAddNodeTail(m, v828, v823)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L30
	} else {
		goto L271
	}
L271:
	;
	goto L267
L272:
	;
	goto L267
L273:
	;
	v841 = int32(1)
	v842 = l1 + v841
	v844 = F_strlen(m, l1)
	mBase = m.M
	v847 = F___memrchr(m, l1, int32(124), v844+v841)
	mBase = m.M
	goto L275
L274:
	;
	v863 = F_zstrdup(m, v842)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L30
	} else {
		goto L283
	}
L275:
	;
	if v847 != 0 {
		goto L274
	} else {
		goto L276
	}
L276:
	;
	v848 = F_ACLLookupCommand(m, v842)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L30
	} else {
		goto L278
	}
L277:
	;
	goto L282
L278:
	;
	if v848 == int32(0) {
		goto L277
	} else {
		goto L279
	}
L279:
	;
	F_ACLChangeSelectorPerm(m, l0, v848, int32(1))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L30
	} else {
		goto L280
	}
L280:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v848)+140))
	F_ACLUpdateCommandRules(m, l0, v855, int32(1))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L30
	} else {
		goto L281
	}
L281:
	;
	v1053 = int32(0)
	goto L1
L282:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[0])) = int32(44)
	goto L91
L283:
	;
	v866 = F_strlen(m, v863)
	mBase = m.M
	v869 = F___memrchr(m, v863, int32(124), v866+int32(1))
	mBase = m.M
	goto L284
L284:
	;
	v870 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v869))) = uint8(v870)
	v872 = F_ACLLookupCommand(m, v863)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L30
	} else {
		goto L286
	}
L285:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v872)+204))
	if v879 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L286:
	;
	if v872 != 0 {
		goto L285
	} else {
		goto L287
	}
L287:
	;
	F_valkey_free(m, v863)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L30
	} else {
		goto L288
	}
L288:
	;
	goto L289
L289:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[0])) = int32(44)
	goto L91
L290:
	;
	v888 = v869 + int32(1)
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888))))
	if v889 != 0 {
		goto L294
	} else {
		goto L295
	}
L291:
	;
	F_valkey_free(m, v863)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L30
	} else {
		goto L292
	}
L292:
	;
	goto L293
L293:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[0])) = int32(12)
	goto L91
L294:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v872)+200))
	if v895 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L295:
	;
	F_valkey_free(m, v863)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L30
	} else {
		goto L296
	}
L296:
	;
	goto L297
L297:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[0])) = int32(28)
	goto L91
L298:
	;
	F_ACLUpdateCommandRules(m, l0, v842, int32(1))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L30
	} else {
		goto L311
	}
L299:
	;
	v909 = *(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[1]))
	if int32(3) < v909 {
		goto L307
	} else {
		goto L308
	}
L300:
	;
	v898 = F_ACLLookupCommand(m, v842)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L30
	} else {
		goto L302
	}
L301:
	;
	F_ACLChangeSelectorPerm(m, l0, v898, int32(1))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L30
	} else {
		goto L306
	}
L302:
	;
	if v898 != 0 {
		goto L301
	} else {
		goto L303
	}
L303:
	;
	F_valkey_free(m, v863)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L30
	} else {
		goto L304
	}
L304:
	;
	goto L305
L305:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[0])) = int32(44)
	goto L91
L306:
	;
	goto L298
L307:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v872)+136))
	F_ACLAddAllowedFirstArg(m, l0, v917, v888)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L30
	} else {
		goto L310
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v842
	F__serverLog(m, int32(3), int32(_a_F_ACLSetSelector_13), v12)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L30
	} else {
		goto L309
	}
L309:
	;
	goto L307
L310:
	;
	goto L298
L311:
	;
	F_valkey_free(m, v863)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L30
	} else {
		goto L312
	}
L312:
	;
	v1053 = int32(0)
	goto L1
L313:
	;
	v932 = F_ACLLookupCommand(m, l1+int32(1))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L30
	} else {
		goto L315
	}
L314:
	;
	goto L319
L315:
	;
	if v932 == int32(0) {
		goto L314
	} else {
		goto L316
	}
L316:
	;
	F_ACLChangeSelectorPerm(m, l0, v932, int32(0))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L30
	} else {
		goto L317
	}
L317:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v932)+140))
	F_ACLUpdateCommandRules(m, l0, v939, int32(0))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L30
	} else {
		goto L318
	}
L318:
	;
	goto L93
L319:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[0])) = int32(44)
	goto L91
L320:
	;
	if v951 != int32(-1) {
		goto L93
	} else {
		goto L321
	}
L321:
	;
	goto L322
L322:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[0])) = int32(44)
	v1053 = int32(-1)
	goto L1
L323:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[0])) = int32(28)
	goto L91
L324:
	;
	v1013 = F_sdsnewlen(m, v975, v974)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L30
	} else {
		goto L333
	}
L325:
	;
	v979 = int32(0)
	goto L327
L326:
	;
	goto L332
L327:
	;
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v975+v979))))
	if v988&int32(223) == int32(0) {
		goto L326
	} else {
		goto L329
	}
L329:
	;
	if base.Ui32(base.I32_extend8_s(v988)+int32(-9)) <= base.Ui32(int32(4)) {
		goto L326
	} else {
		goto L330
	}
L330:
	;
	v999 = v979 + int32(1)
	if v999 == v974 {
		goto L324
	} else {
		goto L331
	}
L331:
	;
	v979 = v999
	goto L327
L332:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelector[0])) = int32(28)
	goto L91
L333:
	;
	v1016 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L30
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1016))) = v969
	*(*int32)(unsafe.Add(mBase, uint32(v1016)+4)) = v1013
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v1021 = F_listSearchKey(m, v1020, v1016)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L30
	} else {
		goto L337
	}
L335:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1036 & int32(-3)
	v1053 = int32(0)
	goto L1
L336:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+8))
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1026)))
	*(*int32)(unsafe.Add(mBase, uint32(v1026))) = v1027 | v969
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+4))
	F_sdsfree(m, v1030)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L30
	} else {
		goto L340
	}
L337:
	;
	if v1021 != 0 {
		goto L336
	} else {
		goto L338
	}
L338:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v1024 = F_listAddNodeTail(m, v1023, v1016)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L30
	} else {
		goto L339
	}
L339:
	;
	goto L335
L340:
	;
	F_valkey_free(m, v1016)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L30
	} else {
		goto L341
	}
L341:
	;
	goto L335
}
func F_ACLSetSelectorDatabasePermissions(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = F_intsetNew(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19&int32(16) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = F_sdsnew(m, l1)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L14
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v19 & int32(-17)
	goto L3
L5:
	;
	m.G0 = v13 + int32(16)
	return v196
L6:
	;
	v75 = F_sdssplitlen(m, v27, v48, int32(_a_F_ACLSetSelectorDatabasePermissions_0), int32(1), v13+int32(12))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L7:
	;
	if v15 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	if v48 == int32(0) {
		goto L7
	} else {
		goto L15
	}
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-17))))
	v48 = v47
	goto L8
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-9))))
	v48 = v44
	goto L8
L11:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(-5)))))
	v48 = v41
	goto L8
L12:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-3)))))
	v48 = v38
	goto L8
L13:
	;
	v48 = int32(base.Ui32(v31) >> (uint(int32(3)) % 32))
	goto L8
L14:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-1)))))
	switch v31 & int32(7) {
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
		goto L7
	}
L15:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v51 == int32(44) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v48+int32(-1)))))
	if v57 != int32(44) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L7
L18:
	;
	F_sdsfree(m, v27)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	F_intsetFree(m, v15)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelectorDatabasePermissions[0])) = int32(28)
	v196 = int32(-1)
	goto L5
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v77 < int32(1) {
		v175 = v15
		v178 = v77
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_sdsfreesplitres(m, v75, v178)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L61
	}
L25:
	;
	v84 = v15
	v85 = int32(0)
	v87 = v77
	goto L26
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v75+v85<<(uint(int32(2))%32))))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+int32(-1)))))
	switch v97 & int32(7) {
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
		goto L29
	}
L27:
	;
	v175 = v166
	v178 = v170
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(0)
	v130 = int32(9116376)
	goto L43
L29:
	;
	F_sdsfreesplitres(m, v75, v87)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L37
	}
L30:
	;
	if v114 != 0 {
		goto L28
	} else {
		goto L36
	}
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(-17))))
	v114 = v113
	goto L30
L32:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(-9))))
	v114 = v110
	goto L30
L33:
	;
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94+int32(-5)))))
	v114 = v107
	goto L30
L34:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+int32(-3)))))
	v114 = v104
	goto L30
L35:
	;
	v114 = int32(base.Ui32(v97) >> (uint(int32(3)) % 32))
	goto L30
L36:
	;
	goto L29
L37:
	;
	if v84 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_sdsfree(m, v27)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	F_intsetFree(m, v84)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelectorDatabasePermissions[0])) = int32(28)
	v196 = int32(-1)
	goto L5
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelectorDatabasePermissions[0])) = int32(0)
	v137 = F_strtox_2(m, v94, v13+int32(8), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L44
L44:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v139 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v166 = F_intsetAdd(m, v84, v137, int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L59
	}
L46:
	;
	F_sdsfree(m, v27)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L58
	}
L47:
	;
	F_intsetFree(m, v84)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L57
	}
L48:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelectorDatabasePermissions[0]))
	if v146 == int32(68) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_sdsfreesplitres(m, v75, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v145 = int32(28)
	if v84 != 0 {
		v157 = v145
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v160 = v145
	goto L46
L52:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_sdsfreesplitres(m, v75, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	if base.Ui64(v137) < base.Ui64(int64(2147483648)) {
		goto L45
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v154 = int32(68)
	if v84 == int32(0) {
		v160 = v154
		goto L46
	} else {
		goto L56
	}
L56:
	;
	v157 = v154
	goto L47
L57:
	;
	v160 = v157
	goto L46
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ACLSetSelectorDatabasePermissions[0])) = v160
	v196 = int32(-1)
	goto L5
L59:
	;
	v169 = v85 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v169 < v170 {
		v84 = v166
		v85 = v169
		v87 = v170
		goto L26
	} else {
		goto L60
	}
L60:
	;
	goto L27
L61:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v184 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v175
	F_sdsfree(m, v27)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	F_intsetFree(m, v184)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v196 = int32(0)
	goto L5
}
func F_ACLShouldKillPubsubClient(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v458 int32
	_ = v458
	var v467 int32
	_ = v467
	var v479 int32
	_ = v479
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v13&int32(1) != 0 {
		v479 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(128)
	return v479
L2:
	;
	if v13&int32(2) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v13&int32(262144) != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v13&int32(4) == int32(0) {
		v479 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	v33 = v10 + int32(80)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+14)) = uint8(v36)
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v36
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+15)) = uint8(v36)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = int32(-1)
	if v35 == v36 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v26 == int32(0) {
		v479 = v3
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v479 = v3
	goto L1
L10:
	;
	goto L16
L11:
	;
	goto L10
L12:
	;
	goto L11
L14:
	;
	v479 = int32(1)
	goto L1
L15:
	;
	F_hashtableCleanupIterator(m, v10+int32(80))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L18
	} else {
		goto L47
	}
L16:
	;
	v65 = F_hashtableNext(m, v10+int32(80), v10+int32(76))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_hashtableCleanupIterator(m, v10+int32(80))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L18
	} else {
		goto L46
	}
L18:
	;
	return int32(0)
L19:
	;
	if v65 == int32(0) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
	v72 = F_objectGetVal(m, v71)
	mBase = m.M
	v74 = F_objectGetVal(m, v71)
	mBase = m.M
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+int32(-1)))))
	switch v77 & int32(7) {
	case 0:
		goto L26
	case 1:
		goto L25
	case 2:
		goto L24
	case 3:
		goto L23
	case 4:
		goto L22
	default:
		goto L21
	}
L21:
	;
	v99 = m.G0
	v101 = v99 - int32(16)
	m.G0 = v101
	v104 = v101 + int32(8)
	F_listRewind(m, l1, v104)
	mBase = m.M
	v108 = F_listNext(m, v104)
	mBase = m.M
	if v108 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	goto L21
L23:
	;
	goto L21
L24:
	;
	goto L21
L25:
	;
	goto L21
L26:
	;
	goto L21
L27:
	;
	if v165 != int32(5) {
		goto L16
	} else {
		goto L45
	}
L28:
	;
	m.G0 = v101 + int32(16)
	goto L27
L29:
	;
	v165 = int32(5)
	goto L28
L30:
	;
	v111 = v108
	goto L31
L31:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+int32(-1)))))
	switch v122 & int32(7) {
	case 0:
		goto L38
	case 1:
		goto L37
	case 2:
		goto L36
	case 3:
		goto L35
	case 4:
		goto L34
	default:
		goto L33
	}
L32:
	;
	goto L29
L33:
	;
	goto L41
L34:
	;
	goto L33
L35:
	;
	goto L33
L36:
	;
	goto L33
L37:
	;
	goto L33
L38:
	;
	goto L33
L39:
	;
	v150 = F_listNext(m, v101+int32(8))
	mBase = m.M
	if v150 != 0 {
		v111 = v150
		goto L31
	} else {
		goto L44
	}
L41:
	;
	v142 = F_strcmp(m, v119, v72)
	mBase = m.M
	if v142 != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v165 = int32(0)
	goto L28
L44:
	;
	goto L32
L45:
	;
	goto L17
L46:
	;
	goto L14
L47:
	;
	v180 = v10 + int32(16)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v183 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v180)+14)) = uint8(v183)
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v180)+24)) = v183
	*(*uint8)(unsafe.Add(mBase, uint32(v180)+15)) = uint8(v183)
	*(*int32)(unsafe.Add(mBase, uint32(v180)+8)) = int32(-1)
	if v182 == v183 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L53
L49:
	;
	goto L48
L50:
	;
	goto L49
L52:
	;
	F_hashtableCleanupIterator(m, v10+int32(16))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L18
	} else {
		goto L83
	}
L53:
	;
	v212 = F_hashtableNext(m, v10+int32(16), v10+int32(12))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L18
	} else {
		goto L55
	}
L54:
	;
	F_hashtableCleanupIterator(m, v10+int32(16))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L18
	} else {
		goto L82
	}
L55:
	;
	if v212 == int32(0) {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v217 = F_objectGetVal(m, v216)
	mBase = m.M
	v219 = F_objectGetVal(m, v216)
	mBase = m.M
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219+int32(-1)))))
	switch v222 & int32(7) {
	case 0:
		goto L62
	case 1:
		goto L61
	case 2:
		goto L60
	case 3:
		goto L59
	case 4:
		goto L58
	default:
		v239 = int32(0)
		goto L57
	}
L57:
	;
	v244 = m.G0
	v246 = v244 - int32(16)
	m.G0 = v246
	v249 = v246 + int32(8)
	F_listRewind(m, l1, v249)
	mBase = m.M
	v253 = F_listNext(m, v249)
	mBase = m.M
	if v253 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v219+int32(-17))))
	v239 = v238
	goto L57
L59:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v219+int32(-9))))
	v239 = v235
	goto L57
L60:
	;
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219+int32(-5)))))
	v239 = v232
	goto L57
L61:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219+int32(-3)))))
	v239 = v229
	goto L57
L62:
	;
	v239 = int32(base.Ui32(v222) >> (uint(int32(3)) % 32))
	goto L57
L63:
	;
	if v310 != int32(5) {
		goto L53
	} else {
		goto L81
	}
L64:
	;
	m.G0 = v246 + int32(16)
	goto L63
L65:
	;
	v310 = int32(5)
	goto L64
L66:
	;
	v256 = v253
	goto L67
L67:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264+int32(-1)))))
	switch v267 & int32(7) {
	case 0:
		goto L74
	case 1:
		goto L73
	case 2:
		goto L72
	case 3:
		goto L71
	case 4:
		goto L70
	default:
		v284 = int32(0)
		goto L69
	}
L68:
	;
	goto L65
L69:
	;
	goto L76
L70:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v264+int32(-17))))
	v284 = v283
	goto L69
L71:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v264+int32(-9))))
	v284 = v280
	goto L69
L72:
	;
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v264+int32(-5)))))
	v284 = v277
	goto L69
L73:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264+int32(-3)))))
	v284 = v274
	goto L69
L74:
	;
	v284 = int32(base.Ui32(v267) >> (uint(int32(3)) % 32))
	goto L69
L75:
	;
	v295 = F_listNext(m, v246+int32(8))
	mBase = m.M
	if v295 != 0 {
		v256 = v295
		goto L67
	} else {
		goto L80
	}
L76:
	;
	v289 = int32(0)
	v291 = F_stringmatchlen(m, v264, v284, v217, v239, v289)
	mBase = m.M
	if v291 != 0 {
		v310 = v289
		goto L64
	} else {
		goto L79
	}
L79:
	;
	goto L75
L80:
	;
	goto L68
L81:
	;
	goto L54
L82:
	;
	goto L14
L83:
	;
	v325 = v10 + int32(16)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+8))
	v328 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v325)+14)) = uint8(v328)
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = v327
	*(*int32)(unsafe.Add(mBase, uint32(v325)+24)) = v328
	*(*uint8)(unsafe.Add(mBase, uint32(v325)+15)) = uint8(v328)
	*(*int32)(unsafe.Add(mBase, uint32(v325)+8)) = int32(-1)
	if v327 == v328 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L88
L85:
	;
	goto L84
L86:
	;
	goto L85
L88:
	;
	v357 = F_hashtableNext(m, v10+int32(16), v10+int32(12))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L18
	} else {
		goto L91
	}
L89:
	;
	F_hashtableCleanupIterator(m, v10+int32(16))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L18
	} else {
		goto L119
	}
L90:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v365 = F_objectGetVal(m, v364)
	mBase = m.M
	v367 = F_objectGetVal(m, v364)
	mBase = m.M
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367+int32(-1)))))
	switch v370 & int32(7) {
	case 0:
		goto L99
	case 1:
		goto L98
	case 2:
		goto L97
	case 3:
		goto L96
	case 4:
		goto L95
	default:
		v387 = int32(0)
		goto L94
	}
L91:
	;
	if v357 != 0 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	F_hashtableCleanupIterator(m, v10+int32(16))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L18
	} else {
		goto L93
	}
L93:
	;
	v479 = int32(0)
	goto L1
L94:
	;
	v392 = m.G0
	v394 = v392 - int32(16)
	m.G0 = v394
	v397 = v394 + int32(8)
	F_listRewind(m, l1, v397)
	mBase = m.M
	v401 = F_listNext(m, v397)
	mBase = m.M
	if v401 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L95:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v367+int32(-17))))
	v387 = v386
	goto L94
L96:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v367+int32(-9))))
	v387 = v383
	goto L94
L97:
	;
	v380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v367+int32(-5)))))
	v387 = v380
	goto L94
L98:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367+int32(-3)))))
	v387 = v377
	goto L94
L99:
	;
	v387 = int32(base.Ui32(v370) >> (uint(int32(3)) % 32))
	goto L94
L100:
	;
	if v458 != int32(5) {
		goto L88
	} else {
		goto L118
	}
L101:
	;
	m.G0 = v394 + int32(16)
	goto L100
L102:
	;
	v458 = int32(5)
	goto L101
L103:
	;
	v404 = v401
	goto L104
L104:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v404)+8))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412+int32(-1)))))
	switch v415 & int32(7) {
	case 0:
		goto L111
	case 1:
		goto L110
	case 2:
		goto L109
	case 3:
		goto L108
	case 4:
		goto L107
	default:
		v432 = int32(0)
		goto L106
	}
L105:
	;
	goto L102
L106:
	;
	goto L113
L107:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v412+int32(-17))))
	v432 = v431
	goto L106
L108:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v412+int32(-9))))
	v432 = v428
	goto L106
L109:
	;
	v425 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v412+int32(-5)))))
	v432 = v425
	goto L106
L110:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412+int32(-3)))))
	v432 = v422
	goto L106
L111:
	;
	v432 = int32(base.Ui32(v415) >> (uint(int32(3)) % 32))
	goto L106
L112:
	;
	v443 = F_listNext(m, v394+int32(8))
	mBase = m.M
	if v443 != 0 {
		v404 = v443
		goto L104
	} else {
		goto L117
	}
L113:
	;
	v437 = int32(0)
	v439 = F_stringmatchlen(m, v412, v432, v365, v387, v437)
	mBase = m.M
	if v439 != 0 {
		v458 = v437
		goto L101
	} else {
		goto L116
	}
L116:
	;
	goto L112
L117:
	;
	goto L105
L118:
	;
	goto L89
L119:
	;
	goto L14
}
func F_ACLStringSetUser(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
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
	var v37 int32
	_ = v37
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
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
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v265 int32
	_ = v265
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	if l0|l1 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_ACLStringSetUser_0), int32(_a_F_ACLStringSetUser_1), int32(2349))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L6
	} else {
		goto L93
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_ACLStringSetUser_2), int32(_a_F_ACLStringSetUser_1), int32(2308))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L6
	} else {
		goto L92
	}
L3:
	;
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v19
	v27 = F_ACLMergeSelectorArguments(m, l2, l3, v14+int32(36), v14+int32(32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v14 + int32(48)
	return v349
L5:
	;
	v42 = F_ACLCreateUnlinkedUser(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L11
	}
L6:
	;
	return int32(0)
L7:
	;
	if v27 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v31 = F_sdsempty(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2+v33<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v37
	v40 = F_sdscatfmt(m, v31, int32(_a_F_ACLStringSetUser_3), v14)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v349 = v40
	goto L4
L11:
	;
	if l0 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	if v49 <= v48 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	F_ACLCopyUser(m, v42, l0)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	F_sdsfree(m, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L6
	} else {
		goto L78
	}
L16:
	;
	if l0 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L17:
	;
	v54 = v48
	goto L18
L18:
	;
	v66 = v27 + v54<<(uint(int32(2))%32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+int32(-1)))))
	switch v70 & int32(7) {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L23
	case 3:
		goto L22
	case 4:
		goto L21
	default:
		v87 = int32(0)
		goto L20
	}
L19:
	;
	v95 = F___errno_location(m)
	mBase = m.M
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	switch v96 + int32(-7) {
	case 0:
		goto L35
	default:
		goto L32
	case 2:
		goto L36
	case 5:
		goto L34
	case 13:
		goto L39
	case 21:
		goto L40
	case 24:
		goto L38
	case 36:
		goto L37
	case 37:
		v108 = int32(_a_F_ACLStringSetUser_4)
		goto L31
	case 61:
		goto L33
	}
L20:
	;
	v88 = F_ACLSetUser(m, v42, v67, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L27
	}
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(-17))))
	v87 = v86
	goto L20
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(-9))))
	v87 = v83
	goto L20
L23:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+int32(-5)))))
	v87 = v80
	goto L20
L24:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+int32(-3)))))
	v87 = v77
	goto L20
L25:
	;
	v87 = int32(base.Ui32(v70) >> (uint(int32(3)) % 32))
	goto L20
L26:
	;
	goto L19
L27:
	;
	if v88 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v91 = v54 + int32(1)
	if v91 == v49 {
		goto L16
	} else {
		goto L29
	}
L29:
	;
	v54 = v91
	goto L18
L30:
	;
	v111 = F_sdsempty(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L6
	} else {
		goto L41
	}
L31:
	;
	v110 = v108
	goto L30
L32:
	;
	v108 = int32(_a_F_ACLStringSetUser_5)
	goto L31
L33:
	;
	v110 = int32(_a_F_ACLStringSetUser_6)
	goto L30
L34:
	;
	v110 = int32(_a_F_ACLStringSetUser_7)
	goto L30
L35:
	;
	v110 = int32(_a_F_ACLStringSetUser_8)
	goto L30
L36:
	;
	v110 = int32(_a_F_ACLStringSetUser_9)
	goto L30
L37:
	;
	v110 = int32(_a_F_ACLStringSetUser_10)
	goto L30
L38:
	;
	v110 = int32(_a_F_ACLStringSetUser_11)
	goto L30
L39:
	;
	v110 = int32(_a_F_ACLStringSetUser_12)
	goto L30
L40:
	;
	v110 = int32(_a_F_ACLStringSetUser_13)
	goto L30
L41:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v113
	v119 = F_sdscatfmt(m, v111, int32(_a_F_ACLStringSetUser_14), v14+int32(16))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v284 = v119
	goto L15
L43:
	;
	F_ACLCopyUser(m, v265, v42)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L6
	} else {
		goto L77
	}
L44:
	;
	if v252 == int32(0) {
		goto L1
	} else {
		goto L76
	}
L45:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v232 & int32(7) {
	case 0:
		goto L74
	case 1:
		goto L73
	case 2:
		goto L72
	case 3:
		goto L71
	case 4:
		goto L70
	default:
		v249 = int32(0)
		goto L69
	}
L46:
	;
	v136 = int32(_a_F_ACLStringSetUser_15)
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_ACLStringSetUser[0]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_ACLStringSetUser[1]))
	v142 = F_kvstoreSize(m, v141)
	mBase = m.M
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_ACLStringSetUser[2]))
	v148 = F_kvstoreSize(m, v147)
	mBase = m.M
	goto L47
L47:
	;
	if base.I32_wrap_i64(v142+base.I64_extend_i32_u(v139+v138)+v148) == int32(0) {
		v265 = l0
		goto L43
	} else {
		goto L48
	}
L48:
	;
	v153 = F_getUpcomingChannelList(m, v42, l0)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	if v153 == int32(0) {
		v265 = l0
		goto L43
	} else {
		goto L50
	}
L50:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_ACLStringSetUser[3]))
	v160 = v14 + int32(40)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v161
	goto L51
L51:
	;
	v166 = v14 + int32(40)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	if v168 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	F_listRelease(m, v153)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L6
	} else {
		goto L68
	}
L53:
	;
	if v168 == int32(0) {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L53
L55:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v168+base.B2i32(v171 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v177
	goto L54
L56:
	;
	v183 = v168
	goto L57
L57:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+328))
	if v193 != l0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L52
L59:
	;
	v203 = v14 + int32(40)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if v205 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v195 = F_ACLShouldKillPubsubClient(m, v192, v153)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	if v195 == int32(0) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	F_freeClientOrCloseLater(m, v192, int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	if v205 != 0 {
		v183 = v205
		goto L57
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v205+base.B2i32(v208 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v214
	goto L65
L67:
	;
	goto L58
L68:
	;
	v252 = l0
	goto L44
L69:
	;
	v250 = F_ACLCreateUser(m, l1, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L6
	} else {
		goto L75
	}
L70:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v249 = v248
	goto L69
L71:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v249 = v245
	goto L69
L72:
	;
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v249 = v242
	goto L69
L73:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v249 = v239
	goto L69
L74:
	;
	v249 = int32(base.Ui32(v232) >> (uint(int32(3)) % 32))
	goto L69
L75:
	;
	v252 = v250
	goto L44
L76:
	;
	v265 = v252
	goto L43
L77:
	;
	v284 = int32(0)
	goto L15
L78:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v293 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	F_listRelease(m, v300)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L6
	} else {
		goto L82
	}
L80:
	;
	F_decrRefCount(m, v293)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = int32(0)
	goto L79
L82:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	F_listRelease(m, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	F_valkey_free(m, v42)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	v308 = int32(0)
	if v49 <= v308 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	F_valkey_free(m, v27)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L6
	} else {
		goto L91
	}
L86:
	;
	v313 = v308
	goto L87
L87:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v27+v313<<(uint(int32(2))%32))))
	F_sdsfree(m, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L6
	} else {
		goto L89
	}
L88:
	;
	goto L85
L89:
	;
	v329 = v313 + int32(1)
	if v329 != v49 {
		v313 = v329
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v349 = v284
	goto L4
L92:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ACLUserCheckKeyPerm(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	v6 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v6 {
		v154 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v154
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = v10 + int32(8)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v18
	goto L3
L3:
	;
	goto L4
L4:
	;
	v30 = v10 + int32(8)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v32 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v154 = v6
	goto L1
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v48 = m.G0
	v50 = v48 - int32(16)
	m.G0 = v50
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v53&int32(2) != 0 {
		v143 = int32(0)
		goto L12
	} else {
		goto L13
	}
L7:
	;
	if v32 != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v32+base.B2i32(v35 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v41
	goto L8
L10:
	;
	v154 = int32(3)
	goto L1
L11:
	;
	if v143 != 0 {
		goto L4
	} else {
		goto L34
	}
L12:
	;
	m.G0 = v50 + int32(16)
	goto L11
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v44)+140))
	v58 = v50 + int32(8)
	F_listRewind(m, v56, v58)
	mBase = m.M
	v62 = F_listNext(m, v58)
	mBase = m.M
	if v62 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v143 = int32(3)
	goto L12
L15:
	;
	v65 = int32(2)
	if l3&int32(160) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v72 = v65
	goto L18
L17:
	;
	v72 = int32(base.Ui32(l3)>>(uint(int32(5))%32)) & v65
	goto L18
L18:
	;
	v77 = v72 | int32(base.Ui32(l3)>>(uint(int32(4))%32))&int32(1)
	v78 = v62
	goto L19
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v87&v77 != v77 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L14
L21:
	;
	v127 = F_listNext(m, v50+int32(8))
	mBase = m.M
	if v127 != 0 {
		v78 = v127
		goto L19
	} else {
		goto L33
	}
L22:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+int32(-1)))))
	switch v94 & int32(7) {
	case 0:
		goto L28
	case 1:
		goto L27
	case 2:
		goto L26
	case 3:
		goto L25
	case 4:
		goto L24
	default:
		v111 = int32(0)
		goto L23
	}
L23:
	;
	if l4 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v91+int32(-17))))
	v111 = v110
	goto L23
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v91+int32(-9))))
	v111 = v107
	goto L23
L26:
	;
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91+int32(-5)))))
	v111 = v104
	goto L23
L27:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+int32(-3)))))
	v111 = v101
	goto L23
L28:
	;
	v111 = int32(base.Ui32(v94) >> (uint(int32(3)) % 32))
	goto L23
L29:
	;
	v119 = int32(0)
	v121 = F_stringmatchlen(m, v91, v111, l1, l2, v119)
	mBase = m.M
	if v121 != 0 {
		v143 = v119
		goto L12
	} else {
		goto L32
	}
L30:
	;
	v114 = int32(0)
	v116 = F_prefixmatchlen(m, v91, v111, l1, l2, v114)
	mBase = m.M
	if v116 == v114 {
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v143 = v114
	goto L12
L32:
	;
	goto L21
L33:
	;
	goto L20
L34:
	;
	goto L5
}
func F_aclAddReplySelectorDescription(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
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
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
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
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v223 int64
	_ = v223
	var v227 int64
	_ = v227
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v235 int32
	_ = v235
	var v239 int64
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int64
	_ = v272
	var v276 int64
	_ = v276
	var v280 int64
	_ = v280
	var v281 int64
	_ = v281
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	v7 = m.G0
	v9 = v7 - int32(64)
	m.G0 = v9
	F_addReplyBulkCString(m, l0, int32(_a_F_aclAddReplySelectorDescription_0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = F_ACLDescribeSelectorCommandRules(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_addReplyBulkSds(m, l0, v14)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_aclAddReplySelectorDescription_1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v21&int32(2) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_aclAddReplySelectorDescription_2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L28
	}
L7:
	;
	v30 = F_sdsempty(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_aclAddReplySelectorDescription_3), int32(2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v34 = v9 + int32(56)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v35
	goto L11
L11:
	;
	v40 = v9 + int32(56)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v42 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	F_addReplyBulkSds(m, l0, v88)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L27
	}
L13:
	;
	if v42 == int32(0) {
		v88 = v30
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42+base.B2i32(v45 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v51
	goto L14
L16:
	;
	v58 = v30
	v59 = v42
	goto L17
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v59 == v63 {
		v68 = v58
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v88 = v69
	goto L12
L19:
	;
	v69 = F_sdsCatPatternString(m, v68, v61)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v66 = F_sdscat(m, v58, int32(_a_F_aclAddReplySelectorDescription_4))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v68 = v66
	goto L19
L22:
	;
	v72 = v9 + int32(56)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v74 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v74 != 0 {
		v58 = v69
		v59 = v74
		goto L17
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v74+base.B2i32(v77 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v83
	goto L24
L26:
	;
	goto L18
L27:
	;
	goto L6
L28:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v102&int32(8) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_aclAddReplySelectorDescription_5))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L51
	}
L30:
	;
	v111 = F_sdsempty(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_aclAddReplySelectorDescription_6), int32(2))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
	v115 = v9 + int32(56)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v116
	goto L34
L34:
	;
	v121 = v9 + int32(56)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v123 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	F_addReplyBulkSds(m, l0, v173)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L50
	}
L36:
	;
	if v123 == int32(0) {
		v173 = v111
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v123+base.B2i32(v126 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v132
	goto L37
L39:
	;
	v139 = v111
	v140 = v123
	goto L40
L40:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v140 == v144 {
		v149 = v139
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v173 = v154
	goto L35
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v142
	v154 = F_sdscatfmt(m, v149, int32(_a_F_aclAddReplySelectorDescription_7), v9+int32(32))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	v147 = F_sdscat(m, v139, int32(_a_F_aclAddReplySelectorDescription_4))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v149 = v147
	goto L42
L45:
	;
	v157 = v9 + int32(56)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v159 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v159 != 0 {
		v139 = v154
		v140 = v159
		goto L40
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v159+base.B2i32(v162 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v168
	goto L47
L49:
	;
	goto L41
L50:
	;
	goto L29
L51:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v187&int32(16) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	m.G0 = v9 + int32(64)
	return
L53:
	;
	v196 = F_sdsempty(m)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_aclAddReplySelectorDescription_8), int32(6))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	goto L58
L57:
	;
	F_addReplyBulkSds(m, l0, v305)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L86
	}
L58:
	;
	if v199 == int32(0) {
		v305 = v196
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v204 = F_sdscatlen(m, v196, int32(_a_F_aclAddReplySelectorDescription_9), int32(3))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	v207 = int32(0)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if base.Ui32(v213) <= base.Ui32(v207) {
		v235 = v207
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v247 = int32(1)
	if v199 == v247 {
		v305 = v246
		goto L57
	} else {
		goto L71
	}
L62:
	;
	if v235 == int32(0) {
		v246 = v204
		goto L61
	} else {
		goto L69
	}
L63:
	;
	goto L62
L64:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	switch v217 + int32(-4) {
	case 0:
		goto L67
	default:
		goto L66
	case 4:
		goto L68
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v232
	v235 = int32(1)
	goto L63
L66:
	;
	v231 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+int32(8)))))
	v232 = v231
	goto L65
L67:
	;
	v227 = int64(*(*int32)(unsafe.Add(mBase, uint32(v206+int32(8)))))
	v232 = v227
	goto L65
L68:
	;
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v206+int32(8))))
	v232 = v223
	goto L65
L69:
	;
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v9)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v239
	v244 = F_sdscatfmt(m, v204, int32(_a_F_aclAddReplySelectorDescription_10), v9+int32(16))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v246 = v244
	goto L61
L71:
	;
	v253 = v247
	v254 = v246
	goto L72
L72:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if base.Ui32(v262) <= base.Ui32(v253) {
		v284 = int32(0)
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v305 = v297
	goto L57
L74:
	;
	v299 = v253 + int32(1)
	if v299 != v199 {
		v253 = v299
		v254 = v297
		goto L72
	} else {
		goto L85
	}
L75:
	;
	if v284 == int32(0) {
		v297 = v254
		goto L74
	} else {
		goto L82
	}
L76:
	;
	goto L75
L77:
	;
	v265 = v256 + int32(8)
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
	switch v266 + int32(-4) {
	case 0:
		goto L80
	default:
		goto L79
	case 4:
		goto L81
	}
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v281
	v284 = int32(1)
	goto L76
L79:
	;
	v280 = int64(*(*int16)(unsafe.Add(mBase, uint32(v265+v253<<(uint(int32(1))%32)))))
	v281 = v280
	goto L78
L80:
	;
	v276 = int64(*(*int32)(unsafe.Add(mBase, uint32(v265+v253<<(uint(int32(2))%32)))))
	v281 = v276
	goto L78
L81:
	;
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v265+v253<<(uint(int32(3))%32))))
	v281 = v272
	goto L78
L82:
	;
	v290 = F_sdscatlen(m, v254, int32(_a_F_aclAddReplySelectorDescription_11), int32(1))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v9)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v292
	v295 = F_sdscatfmt(m, v290, int32(_a_F_aclAddReplySelectorDescription_10), v9)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v297 = v295
	goto L74
L85:
	;
	goto L73
L86:
	;
	goto L52
}
func F_aclCatWithFlags(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
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
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v14 = v11 + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v5
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(-1)
	if l1 == v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v37 = F_hashtableNext(m, v11+int32(16), v11+int32(12))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L1
L3:
	;
	goto L2
L5:
	;
	F_hashtableCleanupIterator(m, v11+int32(16))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L25
	}
L6:
	;
	return
L7:
	;
	if v37 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)+64))
	if v50&l2 == int64(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L5
L11:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v49)+200))
	if v85 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+140))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+int32(-1)))))
	switch v58 & int32(7) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	case 4:
		goto L14
	default:
		v75 = int32(0)
		goto L13
	}
L13:
	;
	F_addReplyBulkCBuffer(m, l0, v55, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L19
	}
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v55+int32(-17))))
	v75 = v74
	goto L13
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v55+int32(-9))))
	v75 = v71
	goto L13
L16:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55+int32(-5)))))
	v75 = v68
	goto L13
L17:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+int32(-3)))))
	v75 = v65
	goto L13
L18:
	;
	v75 = int32(base.Ui32(v58) >> (uint(int32(3)) % 32))
	goto L13
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v78 + int32(1)
	goto L11
L20:
	;
	v94 = F_hashtableNext(m, v11+int32(16), v11+int32(12))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L23
	}
L21:
	;
	F_aclCatWithFlags(m, l0, v85, l2, l3)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v94 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	goto L10
L25:
	;
	m.G0 = v11 + int32(64)
	return
}
func F_aclCreateSelectorFromOpSet(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v13 != int32(40) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_aclCreateSelectorFromOpSet_0), int32(_a_F_aclCreateSelectorFromOpSet_1), int32(1083))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L28
	}
L2:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+l1+int32(-1)))))
	if v19 != int32(41) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = F_valkey_malloc(m, int32(160))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_aclCreateSelectorFromOpSet[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v28 | int32(16)
	v32 = F_listCreate(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+140)) = v32
	v35 = F_listCreate(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = v35
	v39 = F_intsetNew(m)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+136)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+152)) = v39
	v44 = F_sdsempty(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+148)) = v44
	v47 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = int32(5)
	v64 = F__emscripten_memset_bulkmem(m, v23+v47, base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
	v73 = F_sdsnsplitargs(m, l0+int32(1), l1+int32(-2), v11+int32(12))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v75 <= int32(0) {
		v123 = v75
		v126 = v23
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_sdsfreesplitres(m, v73, v123)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L27
	}
L13:
	;
	v84 = int32(0)
	goto L14
L14:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v73+v84<<(uint(int32(2))%32))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+int32(-1)))))
	switch v93 & int32(7) {
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
		v110 = int32(0)
		goto L16
	}
L15:
	;
	F_ACLFreeSelector(m, v23)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L26
	}
L16:
	;
	v111 = F_ACLSetSelector(m, v23, v90, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L23
	}
L17:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v90+int32(-17))))
	v110 = v109
	goto L16
L18:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v90+int32(-9))))
	v110 = v106
	goto L16
L19:
	;
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90+int32(-5)))))
	v110 = v103
	goto L16
L20:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+int32(-3)))))
	v110 = v100
	goto L16
L21:
	;
	v110 = int32(base.Ui32(v93) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	goto L15
L23:
	;
	if v111 == int32(-1) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v116 = v84 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v117 <= v116 {
		v123 = v117
		v126 = v23
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v84 = v116
	goto L14
L26:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v123 = v122
	v126 = int32(0)
	goto L12
L27:
	;
	m.G0 = v11 + int32(16)
	return v126
L28:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_getAclErrorMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v7 = m.G0
	v9 = v7 - int32(64)
	m.G0 = v9
	switch l0 + int32(-1) {
	case 0:
		if l4 == int32(0) {
			v58 = F_sdsnew(m, int32(_a_F_getAclErrorMessage_0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v76 = v58
				m.G0 = v9 + int32(64)
				return v76
			}
		} else {
			v47 = F_sdsempty(m)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v49
				v55 = F_sdscatfmt(m, v47, int32(_a_F_getAclErrorMessage_1), v9+int32(48))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v76 = v55
					m.G0 = v9 + int32(64)
					return v76
				}
			}
		}
	case 1:
		v67 = F_sdsempty(m)
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return int32(0)
		} else {
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+140))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v70
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v69
			v74 = F_sdscatfmt(m, v67, int32(_a_F_getAclErrorMessage_2), v9)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				v76 = v74
				m.G0 = v9 + int32(64)
				return v76
			}
		}
	case 2:
		if l4 == int32(0) {
			v28 = F_sdsnew(m, int32(_a_F_getAclErrorMessage_3))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v76 = v28
				m.G0 = v9 + int32(64)
				return v76
			}
		} else {
			v15 = F_sdsempty(m)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v19
				v25 = F_sdscatfmt(m, v15, int32(_a_F_getAclErrorMessage_4), v9+int32(16))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v76 = v25
					m.G0 = v9 + int32(64)
					return v76
				}
			}
		}
	default:
		F__serverPanic_1(m, int32(_a_F_getAclErrorMessage_5), int32(3002), int32(_a_F_getAclErrorMessage_6), int32(0))
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 4:
		if l4 == int32(0) {
			v43 = F_sdsnew(m, int32(_a_F_getAclErrorMessage_7))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v76 = v43
				m.G0 = v9 + int32(64)
				return v76
			}
		} else {
			v32 = F_sdsempty(m)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v34
				v40 = F_sdscatfmt(m, v32, int32(_a_F_getAclErrorMessage_8), v9+int32(32))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v76 = v40
					m.G0 = v9 + int32(64)
					return v76
				}
			}
		}
	}
}
