package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaO_chunkid(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v226 int64
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	switch v6 + int32(-61) {
	case 0:
		goto L3
	default:
		goto L1
	case 3:
		goto L2
	}
L1:
	;
	v143 = m.G3
	v149 = m.G0
	v151 = v149 - int32(32)
	m.G0 = v151
	v153 = int32(*(*int8)(unsafe.Add(mBase, uint32(v143)+uint32(_consts[1205]))))
	if v153 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L2:
	;
	v18 = l1 + int32(1)
	if v18&int32(3) == int32(0) {
		v40 = v18
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v11 = F___stpncpy(m, l0, l1+int32(1), l2)
	mBase = m.M
	goto L4
L4:
	;
	v15 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+l2+int32(-1)))) = uint8(v15)
	return
L5:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v74)
	v77 = l2 + int32(-8)
	if base.Ui32(v73) <= base.Ui32(v77) {
		v139 = v18
		goto L21
	} else {
		goto L22
	}
L6:
	;
	v73 = v65 - v18
	goto L5
L7:
	;
	v44 = v40
	goto L15
L8:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v29 = v18
	goto L11
L10:
	;
	v73 = v18 - v18
	goto L5
L11:
	;
	v33 = v29 + int32(1)
	if v33&int32(3) == int32(0) {
		v40 = v33
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v38 != 0 {
		v29 = v33
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v65 = v33
	goto L6
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v53 = int32(-2139062144)
	if (int32(16843008)-v50|v50)&v53 == v53 {
		v44 = v44 + int32(4)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v59 = v44
	goto L18
L17:
	;
	goto L16
L18:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v63 != 0 {
		v59 = v59 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v65 = v59
	goto L6
L20:
	;
	goto L19
L21:
	;
	v140 = F_strlen(m, l0)
	mBase = m.M
	v142 = F_strcpy(m, l0+v140, v139)
	mBase = m.M
	goto L39
L22:
	;
	if l0&int32(3) == int32(0) {
		v100 = l0
		goto L25
	} else {
		goto L26
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+v133))) = int32(3026478)
	v139 = v18 + (v73 - v77)
	goto L21
L24:
	;
	v133 = v125 - l0
	goto L23
L25:
	;
	v104 = v100
	goto L33
L26:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v86 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v89 = l0
	goto L29
L28:
	;
	v133 = l0 - l0
	goto L23
L29:
	;
	v93 = v89 + int32(1)
	if v93&int32(3) == int32(0) {
		v100 = v93
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v98 != 0 {
		v89 = v93
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v125 = v93
	goto L24
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v113 = int32(-2139062144)
	if (int32(16843008)-v110|v110)&v113 == v113 {
		v104 = v104 + int32(4)
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v119 = v104
	goto L36
L35:
	;
	goto L34
L36:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v123 != 0 {
		v119 = v119 + int32(1)
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v125 = v119
	goto L24
L38:
	;
	goto L37
L39:
	;
	return
L40:
	;
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+uint32(_consts[1201]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(8)))) = uint16(v224)
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v143)+uint32(_consts[1202])))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v226
	v229 = l2 + int32(-17)
	if base.Ui32(v217) < base.Ui32(v229) {
		goto L59
	} else {
		goto L60
	}
L41:
	;
	m.G0 = v151 + int32(32)
	v217 = v213 - l1
	goto L40
L42:
	;
	v158 = int32(0)
	v160 = F___memset(m, v151, v158, int32(32))
	mBase = m.M
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+uint32(_consts[1205]))))
	if v161 == v158 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v157 = F___strchrnul(m, l1, v153)
	mBase = m.M
	v213 = v157
	goto L41
L44:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+uint32(_consts[1206]))))
	if v156 != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v187 == int32(0) {
		v213 = l1
		goto L41
	} else {
		goto L51
	}
L47:
	;
	v165 = v143 + int32(_a2249)
	v167 = v161
	goto L48
L48:
	;
	v173 = v151 + int32(base.Ui32(v167)>>(uint(int32(3))%32))&int32(28)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v175 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v174 | v175<<(uint(v167)%32)
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
	if v179 != 0 {
		v165 = v165 + v175
		v167 = v179
		goto L48
	} else {
		goto L50
	}
L49:
	;
	goto L46
L50:
	;
	goto L49
L51:
	;
	v191 = l1
	v193 = v187
	goto L52
L52:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v151+int32(base.Ui32(v193)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v200)>>(uint(v193)%32))&int32(1) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v213 = v208
	goto L41
L54:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	v208 = v191 + int32(1)
	if v206 != 0 {
		v191 = v208
		v193 = v206
		goto L52
	} else {
		goto L56
	}
L55:
	;
	v213 = v191
	goto L41
L56:
	;
	goto L53
L57:
	;
	if l0&int32(3) == int32(0) {
		v347 = l0
		goto L89
	} else {
		goto L90
	}
L58:
	;
	v322 = F_strlen(m, l0)
	mBase = m.M
	v324 = F_strcpy(m, l0+v322, l1)
	mBase = m.M
	goto L86
L59:
	;
	v231 = v217
	goto L61
L60:
	;
	v231 = v229
	goto L61
L61:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v231))))
	if v233 == int32(0) {
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v238 = F_strlen(m, l0)
	mBase = m.M
	v239 = l0 + v238
	if v231 == int32(0) {
		v260 = v239
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if l0&int32(3) == int32(0) {
		v285 = l0
		goto L72
	} else {
		goto L73
	}
L64:
	;
	v262 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v260))) = uint8(v262)
	goto L63
L65:
	;
	v243 = l1
	v244 = v231
	v245 = v239
	goto L66
L66:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	if v247 == int32(0) {
		v260 = v245
		goto L64
	} else {
		goto L68
	}
L67:
	;
	v260 = v252
	goto L64
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v245))) = uint8(v247)
	v251 = int32(1)
	v252 = v245 + v251
	v256 = v244 + int32(-1)
	if v256 != 0 {
		v243 = v243 + v251
		v244 = v256
		v245 = v252
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+v318))) = int32(3026478)
	goto L57
L71:
	;
	v318 = v310 - l0
	goto L70
L72:
	;
	v289 = v285
	goto L80
L73:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v271 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v274 = l0
	goto L76
L75:
	;
	v318 = l0 - l0
	goto L70
L76:
	;
	v278 = v274 + int32(1)
	if v278&int32(3) == int32(0) {
		v285 = v278
		goto L72
	} else {
		goto L78
	}
L78:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	if v283 != 0 {
		v274 = v278
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v310 = v278
	goto L71
L80:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v298 = int32(-2139062144)
	if (int32(16843008)-v295|v295)&v298 == v298 {
		v289 = v289 + int32(4)
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v304 = v289
	goto L83
L82:
	;
	goto L81
L83:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	if v308 != 0 {
		v304 = v304 + int32(1)
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v310 = v304
	goto L71
L85:
	;
	goto L84
L86:
	;
	goto L57
L87:
	;
	v381 = l0 + v380
	v382 = m.G3
	v385 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v382)+uint32(_consts[1203]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v381))) = uint16(v385)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+uint32(_consts[1204]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v381+int32(2)))) = uint8(v391)
	return
L88:
	;
	v380 = v372 - l0
	goto L87
L89:
	;
	v351 = v347
	goto L97
L90:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v333 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v336 = l0
	goto L93
L92:
	;
	v380 = l0 - l0
	goto L87
L93:
	;
	v340 = v336 + int32(1)
	if v340&int32(3) == int32(0) {
		v347 = v340
		goto L89
	} else {
		goto L95
	}
L95:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	if v345 != 0 {
		v336 = v340
		goto L93
	} else {
		goto L96
	}
L96:
	;
	v372 = v340
	goto L88
L97:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	v360 = int32(-2139062144)
	if (int32(16843008)-v357|v357)&v360 == v360 {
		v351 = v351 + int32(4)
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v366 = v351
	goto L100
L99:
	;
	goto L98
L100:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	if v370 != 0 {
		v366 = v366 + int32(1)
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v372 = v366
	goto L88
L102:
	;
	goto L101
}
func F_luaO_fb2int(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v10 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(31)
	if v10 != 0 {
		v14 = (l0&int32(7) | int32(8)) << (uint(v10+int32(-1)) % 32)
	} else {
		v14 = l0
	}
	return v14
}
func F_luaO_int2fb(m *base.Module, l0 int32) int32 {
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	if base.Ui32(int32(16)) <= base.Ui32(l0) {
		v8 = l0
		v9 = int32(0)
		for {
			v11 = int32(1)
			v12 = v9 + v11
			v14 = v8 + v11
			v16 = int32(base.Ui32(v14) >> (uint(v11) % 32))
			if base.Ui32(int32(31)) < base.Ui32(v14) {
				v8 = v16
				v9 = v12
				continue
			} else {
				break
			}
			break
		}
		v23 = v16
		v24 = v12<<(uint(int32(3))%32) + int32(8)
	} else {
		v23 = l0
		v24 = int32(8)
	}
	if base.Ui32(v23) < base.Ui32(int32(8)) {
		v31 = v23
	} else {
		v31 = v24 | (v23 + int32(-8))
	}
	return v31
}
func F_luaO_rawequalObj(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v5 == v6 {
		switch v5 {
		case 0:
			v26 = int32(1)
			return v26
		case 1:
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			return base.B2i32(v15 == v16)
		case 2:
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			return base.B2i32(v19 == v20)
		case 3:
			v11 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
			v12 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
			return base.F64_eq(v11, v12)
		default:
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v26 = base.B2i32(v23 == v24)
			return v26
		}
	} else {
		return int32(0)
	}
}
func F_luaO_str2d(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 float64
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int64
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_strtod(m, l0, v8+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(l1))) = v12
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		if v17 != l0 {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			if v20|int32(32) != int32(120) {
				v35 = v17
				v36 = v20
			} else {
				v29 = F_strtox_2(m, l0, v8+int32(12), int32(16), int64(4294967295))
				mBase = m.M
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_convert_i32_u(base.I32_wrap_i64(v29))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
				v35 = v33
				v36 = v34
			}
			if v36&int32(255) != 0 {
				v43 = v35
				v44 = v36
				for {
					v46 = v44 & int32(255)
					if base.Ui32(v46+int32(-9)) < base.Ui32(int32(5)) {
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
						v43 = v43 + int32(1)
						v44 = v53
						continue
					} else {
					}
					if v46 != int32(32) {
						break
					} else {
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
						v43 = v43 + int32(1)
						v44 = v53
						continue
					}
					break
				}
				v63 = base.B2i32(v44&int32(255) == int32(0))
			} else {
				v63 = int32(1)
			}
		} else {
			v63 = int32(0)
		}
		m.G0 = v8 + int32(16)
		return v63
	}
}
