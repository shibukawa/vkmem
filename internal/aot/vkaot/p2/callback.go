package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_callbackValDestructor(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = m.G4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	m.T0[v3].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_errorCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
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
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
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
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v334 int64
	_ = v334
	var v340 int64
	_ = v340
	var v342 int64
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = F_lua_checkstack(m, v13, int32(3))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverPanic_2(m, int32(944))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L2
	} else {
		goto L109
	}
L2:
	;
	return
L3:
	;
	if v15 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = int32(9116376)
	goto L7
L5:
	;
	v383 = m.G3
	F_lua_pushstring(m, v13, v383+int32(_a2007))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L2
	} else {
		goto L106
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
	v368 = m.G3
	v371 = F_lm_asprintf(m, v368+int32(_a2008), v11)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L103
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v20 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v20
	v26 = m.G3
	v33 = m.G10
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	m.T0[v34].(func(*base.Module, int32, int32, int32, int32))(m, v23, v26+int32(_a774), v26+int32(_a2009), v11+int32(32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	switch v37 + int32(-28) {
	case 0:
		goto L13
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L10
	case 16:
		goto L12
	default:
		goto L14
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	v354 = m.G3
	v359 = F_lm_asprintf(m, v354+int32(_a2008), v11+int32(16))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L2
	} else {
		goto L100
	}
L11:
	;
	v206 = m.G3
	v208 = v206 + int32(_a2010)
	goto L63
L12:
	;
	v152 = m.G3
	v154 = v152 + int32(_a2011)
	goto L48
L13:
	;
	v98 = m.G3
	v100 = v98 + int32(_a2012)
	goto L33
L14:
	;
	if v37 == int32(2) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	if v37 != int32(70) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v44 = m.G3
	v46 = v44 + int32(_a2013)
	goto L18
L17:
	;
	if v80-v85 != 0 {
		goto L10
	} else {
		goto L30
	}
L18:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v51 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	goto L17
L21:
	;
	v53 = l1
	v54 = v46
	v55 = int32(12)
	v56 = v51
	goto L24
L22:
	;
	v80 = int32(0)
	v81 = v46
	goto L20
L23:
	;
	v80 = v77 & int32(255)
	v81 = v75
	goto L20
L24:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v56&int32(255) != v60 {
		v75 = v54
		v77 = v56
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v75 = v69
	v77 = int32(0)
	goto L23
L26:
	;
	if v60 == int32(0) {
		v75 = v54
		v77 = v56
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v65 = v55 + int32(-1)
	if v65 == int32(0) {
		v75 = v54
		v77 = v56
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v68 = int32(1)
	v69 = v54 + v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	if v70 != 0 {
		v53 = v53 + v68
		v54 = v69
		v55 = v65
		v56 = v70
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v93 = m.G3
	F_luaPushErrorBuff(m, v13, v93+int32(_a2014))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	goto L5
L32:
	;
	if v134-v139 != 0 {
		goto L10
	} else {
		goto L45
	}
L33:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v105 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	goto L32
L36:
	;
	v107 = l1
	v108 = v100
	v109 = int32(34)
	v110 = v105
	goto L39
L37:
	;
	v134 = int32(0)
	v135 = v100
	goto L35
L38:
	;
	v134 = v131 & int32(255)
	v135 = v129
	goto L35
L39:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v110&int32(255) != v114 {
		v129 = v108
		v131 = v110
		goto L38
	} else {
		goto L41
	}
L40:
	;
	v129 = v123
	v131 = int32(0)
	goto L38
L41:
	;
	if v114 == int32(0) {
		v129 = v108
		v131 = v110
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v119 = v109 + int32(-1)
	if v119 == int32(0) {
		v129 = v108
		v131 = v110
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v122 = int32(1)
	v123 = v108 + v122
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	if v124 != 0 {
		v107 = v107 + v122
		v108 = v123
		v109 = v119
		v110 = v124
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	v147 = m.G3
	F_luaPushErrorBuff(m, v13, v147+int32(_a2015))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	goto L5
L47:
	;
	if v188-v193 != 0 {
		goto L10
	} else {
		goto L60
	}
L48:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v159 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	goto L47
L51:
	;
	v161 = l1
	v162 = v154
	v163 = int32(21)
	v164 = v159
	goto L54
L52:
	;
	v188 = int32(0)
	v189 = v154
	goto L50
L53:
	;
	v188 = v185 & int32(255)
	v189 = v183
	goto L50
L54:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v164&int32(255) != v168 {
		v183 = v162
		v185 = v164
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v183 = v177
	v185 = int32(0)
	goto L53
L56:
	;
	if v168 == int32(0) {
		v183 = v162
		v185 = v164
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v173 = v163 + int32(-1)
	if v173 == int32(0) {
		v183 = v162
		v185 = v164
		goto L53
	} else {
		goto L58
	}
L58:
	;
	v176 = int32(1)
	v177 = v162 + v176
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	if v178 != 0 {
		v161 = v161 + v176
		v162 = v177
		v163 = v173
		v164 = v178
		goto L54
	} else {
		goto L59
	}
L59:
	;
	goto L55
L60:
	;
	v201 = m.G3
	F_luaPushErrorBuff(m, v13, v201+int32(_a2016))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	goto L5
L62:
	;
	if v242-v247 != 0 {
		goto L10
	} else {
		goto L75
	}
L63:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v213 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	goto L62
L66:
	;
	v215 = l1
	v216 = v208
	v217 = int32(7)
	v218 = v213
	goto L69
L67:
	;
	v242 = int32(0)
	v243 = v208
	goto L65
L68:
	;
	v242 = v239 & int32(255)
	v243 = v237
	goto L65
L69:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	if v218&int32(255) != v222 {
		v237 = v216
		v239 = v218
		goto L68
	} else {
		goto L71
	}
L70:
	;
	v237 = v231
	v239 = int32(0)
	goto L68
L71:
	;
	if v222 == int32(0) {
		v237 = v216
		v239 = v218
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v227 = v217 + int32(-1)
	if v227 == int32(0) {
		v237 = v216
		v239 = v218
		goto L68
	} else {
		goto L73
	}
L73:
	;
	v230 = int32(1)
	v231 = v216 + v230
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	if v232 != 0 {
		v215 = v215 + v230
		v216 = v231
		v217 = v227
		v218 = v232
		goto L69
	} else {
		goto L74
	}
L74:
	;
	goto L70
L75:
	;
	v256 = l1 + int32(7)
	if v256&int32(3) == int32(0) {
		v278 = v256
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v313 = v311 + int32(24)
	v314 = m.G22
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v316 = m.T0[v315].(func(*base.Module, int32) int32)(m, v313)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L2
	} else {
		goto L92
	}
L77:
	;
	v311 = v303 - v256
	goto L76
L78:
	;
	v282 = v278
	goto L86
L79:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
	if v264 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v267 = v256
	goto L82
L81:
	;
	v311 = v256 - v256
	goto L76
L82:
	;
	v271 = v267 + int32(1)
	if v271&int32(3) == int32(0) {
		v278 = v271
		goto L78
	} else {
		goto L84
	}
L84:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if v276 != 0 {
		v267 = v271
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v303 = v271
	goto L77
L86:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	v291 = int32(-2139062144)
	if (int32(16843008)-v288|v288)&v291 == v291 {
		v282 = v282 + int32(4)
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v297 = v282
	goto L89
L88:
	;
	goto L87
L89:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	if v301 != 0 {
		v297 = v297 + int32(1)
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v303 = v297
	goto L77
L91:
	;
	goto L90
L92:
	;
	v320 = int32(0)
	if base.Ui32(v313) < base.Ui32(int32(25)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v324 = v320
	goto L95
L94:
	;
	v324 = v311
	goto L95
L95:
	;
	v326 = F__emscripten_memset_bulkmem(m, v316+int32(24), base.I32_extend8_s(v320), v324)
	mBase = m.M
	goto L96
L96:
	;
	v329 = m.G3
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v329)+uint32(_consts[1010])))
	*(*int64)(unsafe.Add(mBase, uint32(v316+int32(16)))) = v334
	v340 = *(*int64)(unsafe.Add(mBase, uint32(v329)+uint32(_consts[1011])))
	*(*int64)(unsafe.Add(mBase, uint32(v316+int32(8)))) = v340
	v342 = *(*int64)(unsafe.Add(mBase, uint32(v329)+uint32(_consts[1012])))
	*(*int64)(unsafe.Add(mBase, uint32(v316))) = v342
	v344 = F_strlen(m, v316)
	mBase = m.M
	v346 = F_strcpy(m, v316+v344, v256)
	mBase = m.M
	goto L97
L97:
	;
	F_luaPushErrorBuff(m, v13, v316)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	v349 = m.G11
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	m.T0[v350].(func(*base.Module, int32))(m, v316)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	goto L5
L100:
	;
	F_luaPushErrorBuff(m, v13, v359)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L2
	} else {
		goto L101
	}
L101:
	;
	v363 = m.G11
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	m.T0[v364].(func(*base.Module, int32))(m, v359)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	goto L5
L103:
	;
	F_luaPushErrorBuff(m, v13, v371)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v375 = m.G11
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	m.T0[v376].(func(*base.Module, int32))(m, v371)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	goto L5
L106:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v390))) = int32(1)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v396 + int32(16)
	goto L107
L107:
	;
	F_lua_settable(m, v13, int32(-3))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L2
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
	m.G0 = v11 + int32(48)
	return
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
