package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_handleDebugClusterCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v57 int32
	_ = v57
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
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
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
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
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v12 != int32(5) {
		v361 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return v361
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = F_objectGetVal(m, v16)
	mBase = m.M
	v18 = int32(_a339)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if v53-v55 != 0 {
		v361 = v2
		goto L1
	} else {
		goto L15
	}
L4:
	;
	v53 = F_tolower(m, v49)
	mBase = m.M
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v55 = F_tolower(m, v54)
	mBase = m.M
	goto L3
L5:
	;
	v23 = v17
	v24 = v18
	v25 = v21
	goto L8
L6:
	;
	v49 = int32(0)
	v50 = v18
	goto L4
L7:
	;
	v49 = v46 & int32(255)
	v50 = v45
	goto L4
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v27 == int32(0) {
		v45 = v24
		v46 = v25
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v45 = v39
	v46 = int32(0)
	goto L7
L10:
	;
	v31 = v25 & int32(255)
	if v31 == v27 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = int32(1)
	v39 = v24 + v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v40 != 0 {
		v23 = v23 + v38
		v24 = v39
		v25 = v40
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v33 = F_tolower(m, v31)
	mBase = m.M
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v35 = F_tolower(m, v34)
	mBase = m.M
	if v33 == v35 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v45 = v24
	v46 = v37
	goto L7
L14:
	;
	goto L9
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v59 = F_objectGetVal(m, v58)
	mBase = m.M
	v60 = int32(_a340)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v63 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v95-v97 != 0 {
		v361 = v2
		goto L1
	} else {
		goto L28
	}
L17:
	;
	v95 = F_tolower(m, v91)
	mBase = m.M
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v97 = F_tolower(m, v96)
	mBase = m.M
	goto L16
L18:
	;
	v65 = v59
	v66 = v60
	v67 = v63
	goto L21
L19:
	;
	v91 = int32(0)
	v92 = v60
	goto L17
L20:
	;
	v91 = v88 & int32(255)
	v92 = v87
	goto L17
L21:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v69 == int32(0) {
		v87 = v66
		v88 = v67
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v87 = v81
	v88 = int32(0)
	goto L20
L23:
	;
	v73 = v67 & int32(255)
	if v73 == v69 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v80 = int32(1)
	v81 = v66 + v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	if v82 != 0 {
		v65 = v65 + v80
		v66 = v81
		v67 = v82
		goto L21
	} else {
		goto L27
	}
L25:
	;
	v75 = F_tolower(m, v73)
	mBase = m.M
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v77 = F_tolower(m, v76)
	mBase = m.M
	if v75 == v77 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v87 = v66
	v88 = v79
	goto L20
L27:
	;
	goto L22
L28:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v100 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v361 = int32(1)
	goto L1
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	v108 = F_objectGetVal(m, v107)
	mBase = m.M
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	v112 = F_objectGetVal(m, v111)
	mBase = m.M
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+int32(-1)))))
	switch v115 & int32(7) {
	case 0:
		goto L39
	case 1:
		goto L38
	case 2:
		goto L37
	case 3:
		goto L36
	case 4:
		goto L35
	default:
		v132 = int32(0)
		goto L34
	}
L31:
	;
	F_addReplyError(m, l0, int32(_a341))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return int32(0)
L33:
	;
	goto L29
L34:
	;
	if v132 != int32(40) {
		v165 = int32(-1)
		goto L43
	} else {
		goto L44
	}
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v112+int32(-17))))
	v132 = v131
	goto L34
L36:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v112+int32(-9))))
	v132 = v128
	goto L34
L37:
	;
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112+int32(-5)))))
	v132 = v125
	goto L34
L38:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+int32(-3)))))
	v132 = v122
	goto L34
L39:
	;
	v132 = int32(base.Ui32(v115) >> (uint(int32(3)) % 32))
	goto L34
L40:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	if v177 != v189 {
		goto L58
	} else {
		goto L59
	}
L41:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	v182 = F_objectGetVal(m, v181)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v182
	F_addReplyErrorFormat(m, l0, int32(_a227), v9)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L32
	} else {
		goto L57
	}
L42:
	;
	if v165 != 0 {
		goto L41
	} else {
		goto L50
	}
L43:
	;
	goto L42
L44:
	;
	v140 = int32(0)
	goto L46
L45:
	;
	v165 = int32(0) - v155
	goto L43
L46:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v140))))
	v145 = int32(255)
	v155 = base.B2i32(base.Ui32((v142+int32(-123))&v145) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v142+int32(-58))&v145) < base.Ui32(int32(246)))
	if v155 != 0 {
		goto L45
	} else {
		goto L48
	}
L47:
	;
	goto L45
L48:
	;
	v157 = v140 + int32(1)
	if v157 != int32(40) {
		v140 = v157
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v166 = F_sdsnewlen(m, v108, v132)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L32
	} else {
		goto L51
	}
L51:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+32))
	v171 = F_dictFind(m, v170, v166)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L32
	} else {
		goto L52
	}
L52:
	;
	F_sdsfree(m, v166)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L32
	} else {
		goto L53
	}
L53:
	;
	if v171 == int32(0) {
		goto L41
	} else {
		goto L54
	}
L54:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v171)+8))
	goto L55
L55:
	;
	if v177 != 0 {
		goto L40
	} else {
		goto L56
	}
L56:
	;
	goto L41
L57:
	;
	goto L29
L58:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v197 = F_objectGetVal(m, v196)
	mBase = m.M
	v198 = int32(_a342)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if v201 != 0 {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	F_addReplyErrorFormat(m, l0, int32(_a343), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L32
	} else {
		goto L60
	}
L60:
	;
	goto L29
L61:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L32
	} else {
		goto L114
	}
L62:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	v244 = F_objectGetVal(m, v243)
	mBase = m.M
	v245 = int32(_a344)
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	if v248 != 0 {
		goto L81
	} else {
		goto L82
	}
L63:
	;
	if v233-v235 != 0 {
		goto L62
	} else {
		goto L75
	}
L64:
	;
	v233 = F_tolower(m, v229)
	mBase = m.M
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	v235 = F_tolower(m, v234)
	mBase = m.M
	goto L63
L65:
	;
	v203 = v197
	v204 = v198
	v205 = v201
	goto L68
L66:
	;
	v229 = int32(0)
	v230 = v198
	goto L64
L67:
	;
	v229 = v226 & int32(255)
	v230 = v225
	goto L64
L68:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v207 == int32(0) {
		v225 = v204
		v226 = v205
		goto L67
	} else {
		goto L70
	}
L69:
	;
	v225 = v219
	v226 = int32(0)
	goto L67
L70:
	;
	v211 = v205 & int32(255)
	if v211 == v207 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v218 = int32(1)
	v219 = v204 + v218
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+1)))
	if v220 != 0 {
		v203 = v203 + v218
		v204 = v219
		v205 = v220
		goto L68
	} else {
		goto L74
	}
L72:
	;
	v213 = F_tolower(m, v211)
	mBase = m.M
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	v215 = F_tolower(m, v214)
	mBase = m.M
	if v213 == v215 {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v225 = v204
	v226 = v217
	goto L67
L74:
	;
	goto L69
L75:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v177)+2348))
	if v237 == int32(0) {
		goto L61
	} else {
		goto L76
	}
L76:
	;
	F_freeClusterLink(m, v237)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L32
	} else {
		goto L77
	}
L77:
	;
	goto L61
L78:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+12))
	v291 = F_objectGetVal(m, v290)
	mBase = m.M
	v292 = int32(_a345)
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	if v295 != 0 {
		goto L97
	} else {
		goto L98
	}
L79:
	;
	if v280-v282 != 0 {
		goto L78
	} else {
		goto L91
	}
L80:
	;
	v280 = F_tolower(m, v276)
	mBase = m.M
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	v282 = F_tolower(m, v281)
	mBase = m.M
	goto L79
L81:
	;
	v250 = v244
	v251 = v245
	v252 = v248
	goto L84
L82:
	;
	v276 = int32(0)
	v277 = v245
	goto L80
L83:
	;
	v276 = v273 & int32(255)
	v277 = v272
	goto L80
L84:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v254 == int32(0) {
		v272 = v251
		v273 = v252
		goto L83
	} else {
		goto L86
	}
L85:
	;
	v272 = v266
	v273 = int32(0)
	goto L83
L86:
	;
	v258 = v252 & int32(255)
	if v258 == v254 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v265 = int32(1)
	v266 = v251 + v265
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+1)))
	if v267 != 0 {
		v250 = v250 + v265
		v251 = v266
		v252 = v267
		goto L84
	} else {
		goto L90
	}
L88:
	;
	v260 = F_tolower(m, v258)
	mBase = m.M
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	v262 = F_tolower(m, v261)
	mBase = m.M
	if v260 == v262 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v272 = v251
	v273 = v264
	goto L83
L90:
	;
	goto L85
L91:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v177)+2344))
	if v284 == int32(0) {
		goto L61
	} else {
		goto L92
	}
L92:
	;
	F_freeClusterLink(m, v284)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L32
	} else {
		goto L93
	}
L93:
	;
	goto L61
L94:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	v343 = F_objectGetVal(m, v342)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v343
	F_addReplyErrorFormat(m, l0, int32(_a346), v9+int32(16))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L32
	} else {
		goto L113
	}
L95:
	;
	if v327-v329 != 0 {
		goto L94
	} else {
		goto L107
	}
L96:
	;
	v327 = F_tolower(m, v323)
	mBase = m.M
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	v329 = F_tolower(m, v328)
	mBase = m.M
	goto L95
L97:
	;
	v297 = v291
	v298 = v292
	v299 = v295
	goto L100
L98:
	;
	v323 = int32(0)
	v324 = v292
	goto L96
L99:
	;
	v323 = v320 & int32(255)
	v324 = v319
	goto L96
L100:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if v301 == int32(0) {
		v319 = v298
		v320 = v299
		goto L99
	} else {
		goto L102
	}
L101:
	;
	v319 = v313
	v320 = int32(0)
	goto L99
L102:
	;
	v305 = v299 & int32(255)
	if v305 == v301 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v312 = int32(1)
	v313 = v298 + v312
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+1)))
	if v314 != 0 {
		v297 = v297 + v312
		v298 = v313
		v299 = v314
		goto L100
	} else {
		goto L106
	}
L104:
	;
	v307 = F_tolower(m, v305)
	mBase = m.M
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	v309 = F_tolower(m, v308)
	mBase = m.M
	if v307 == v309 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	v319 = v298
	v320 = v311
	goto L99
L106:
	;
	goto L101
L107:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v177)+2344))
	if v331 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v177)+2348))
	if v336 == int32(0) {
		goto L61
	} else {
		goto L111
	}
L109:
	;
	F_freeClusterLink(m, v331)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L32
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	F_freeClusterLink(m, v336)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L32
	} else {
		goto L112
	}
L112:
	;
	goto L61
L113:
	;
	goto L61
L114:
	;
	goto L29
}
func F_handleParseError(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	if v9&int32(2) == int32(0) {
		if v9&int32(4) == int32(0) {
			if v9&int32(8) == int32(0) {
				if v9&int32(16) == int32(0) {
					if v9&int32(32) == int32(0) {
						if v9&int32(64) == int32(0) {
							if v9&int32(128) == int32(0) {
								if v9&int32(256) == int32(0) {
									if v9&int32(1024) == int32(0) {
										if v9&int32(4194304) == int32(0) {
											if v9&int32(512) == int32(0) {
												m.G0 = v7 + int32(16)
												return
											} else {
												v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
												if v131&int32(1) == int32(0) {
													v137 = int32(2)
													if v131&v137 == int32(0) {
														if v131&int32(262144) != 0 {
															v154 = v137
															v157 = v154
														} else {
															v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
															if v147 != 0 {
																v151 = F_isImportSlotMigrationJob(m, v147)
																mBase = m.M
																if v151 != 0 {
																	v152 = int32(4)
																} else {
																	v152 = int32(5)
																}
																v154 = v152
																v157 = v154
															} else {
																v157 = int32(0)
															}
														}
													} else {
														if v131&int32(4) != 0 {
															if v131&int32(262144) != 0 {
																v154 = v137
																v157 = v154
															} else {
																v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																if v147 != 0 {
																	v151 = F_isImportSlotMigrationJob(m, v147)
																	mBase = m.M
																	if v151 != 0 {
																		v152 = int32(4)
																	} else {
																		v152 = int32(5)
																	}
																	v154 = v152
																	v157 = v154
																} else {
																	v157 = int32(0)
																}
															}
														} else {
															v157 = int32(1)
														}
													}
												} else {
													v157 = int32(3)
												}
												v159 = *(*int32)(unsafe.Add(mBase, _consts[28]))
												if v157 != int32(4) {
													v170 = int32(_a1665)
													if int32(3) < v159 {
														v178 = v170
														F_setProtocolError(m, v178, l0)
														mBase = m.M
														v181 = m.ExcPending
														if v181 != 0 {
															return
														} else {
															m.G0 = v7 + int32(16)
															return
														}
													} else {
														F__serverLog(m, int32(3), int32(_a1666), int32(0))
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return
														} else {
															v178 = v170
															F_setProtocolError(m, v178, l0)
															mBase = m.M
															v181 = m.ExcPending
															if v181 != 0 {
																return
															} else {
																m.G0 = v7 + int32(16)
																return
															}
														}
													}
												} else {
													v162 = int32(_a1667)
													if int32(3) < v159 {
														v178 = v162
														F_setProtocolError(m, v178, l0)
														mBase = m.M
														v181 = m.ExcPending
														if v181 != 0 {
															return
														} else {
															m.G0 = v7 + int32(16)
															return
														}
													} else {
														F__serverLog(m, int32(3), int32(_a1668), int32(0))
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return
														} else {
															v178 = v162
															F_setProtocolError(m, v178, l0)
															mBase = m.M
															v181 = m.ExcPending
															if v181 != 0 {
																return
															} else {
																m.G0 = v7 + int32(16)
																return
															}
														}
													}
												}
											}
										} else {
											F_addReplyError(m, l0, int32(_a1669))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												v178 = int32(_a1670)
												F_setProtocolError(m, v178, l0)
												mBase = m.M
												v181 = m.ExcPending
												if v181 != 0 {
													return
												} else {
													m.G0 = v7 + int32(16)
													return
												}
											}
										}
									} else {
										F_addReplyError(m, l0, int32(_a1671))
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											v178 = int32(_a1672)
											F_setProtocolError(m, v178, l0)
											mBase = m.M
											v181 = m.ExcPending
											if v181 != 0 {
												return
											} else {
												m.G0 = v7 + int32(16)
												return
											}
										}
									}
								} else {
									F_addReplyError(m, l0, int32(_a1673))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return
									} else {
										v178 = int32(_a1674)
										F_setProtocolError(m, v178, l0)
										mBase = m.M
										v181 = m.ExcPending
										if v181 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								}
							} else {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v95 = int32(*(*int8)(unsafe.Add(mBase, uint32(v92+v93))))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v95
								F_addReplyErrorFormat(m, l0, int32(_a1675), v7)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									v178 = int32(_a1676)
									F_setProtocolError(m, v178, l0)
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							}
						} else {
							F_addReplyError(m, l0, int32(_a1677))
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return
							} else {
								v178 = int32(_a1678)
								F_setProtocolError(m, v178, l0)
								mBase = m.M
								v181 = m.ExcPending
								if v181 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					} else {
						F_addReplyErrorLength(m, l0, int32(_a1679), int32(43))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_afterErrorReply(m, l0, int32(_a1679), int32(43), int32(0))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								v178 = int32(_a1680)
								F_setProtocolError(m, v178, l0)
								mBase = m.M
								v181 = m.ExcPending
								if v181 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					}
				} else {
					F_addReplyErrorLength(m, l0, int32(_a1681), int32(48))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						F_afterErrorReply(m, l0, int32(_a1681), int32(48), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							v178 = int32(_a1682)
							F_setProtocolError(m, v178, l0)
							mBase = m.M
							v181 = m.ExcPending
							if v181 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				}
			} else {
				F_addReplyErrorLength(m, l0, int32(_a1683), int32(40))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					F_afterErrorReply(m, l0, int32(_a1683), int32(40), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v178 = int32(_a1684)
						F_setProtocolError(m, v178, l0)
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		} else {
			F_addReplyErrorLength(m, l0, int32(_a1685), int32(42))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				F_afterErrorReply(m, l0, int32(_a1685), int32(42), int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v178 = int32(_a1686)
					F_setProtocolError(m, v178, l0)
					mBase = m.M
					v181 = m.ExcPending
					if v181 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	} else {
		F_addReplyErrorLength(m, l0, int32(_a1687), int32(38))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_afterErrorReply(m, l0, int32(_a1687), int32(38), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v178 = int32(_a1688)
				F_setProtocolError(m, v178, l0)
				mBase = m.M
				v181 = m.ExcPending
				if v181 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	}
}
func F_handleParseResults(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	if v3&int32(4196350) == int32(0) {
		if v3&int32(2048) == int32(0) {
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			if v18&int32(1) != 0 {
			} else {
				if v18&int32(2) == int32(0) {
					if v18&int32(262144) != 0 {
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						if v31 == int32(0) {
						} else {
						}
					}
				} else {
					if v18&int32(4) == int32(0) {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						v39 = *(*int64)(unsafe.Add(mBase, _consts[109]))
						*(*int64)(unsafe.Add(mBase, uint32(v37)+80)) = v39
					} else {
						if v18&int32(262144) != 0 {
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							if v31 == int32(0) {
							} else {
							}
						}
					}
				}
			}
		}
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
		if v42&int32(2048) != 0 {
			F_resetClient(m, l0)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				v55 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v55
				return v55
			}
		} else {
			if v42&int32(4096) != 0 {
				F_resetClient(m, l0)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v55 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v55
					return v55
				}
			} else {
				if v42&int32(8192) != 0 {
					v51 = int32(0)
				} else {
					v51 = int32(-2)
				}
				return v51
			}
		}
	} else {
		F_handleParseError(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(-1)
		}
	}
}
func F_hasActiveChildProcess(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	return base.B2i32(v2 != int32(-1))
}
func F_hdr_iter_next(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v3 = m.T0[v2].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_hexistsCommand(m *base.Module, l0 int32) {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v12 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v13 = F_lookupKeyReadOrReply(m, l0, v10, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 == int32(0) {
			m.G0 = v7 + int32(16)
			return
		} else {
			v18 = F_checkType(m, l0, v13, int32(4))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				if v18 != 0 {
					m.G0 = v7 + int32(16)
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
					v22 = F_objectGetVal(m, v21)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(9223372036854775807)
					v28 = int32(12)
					v35 = F_hashTypeGetValue(m, v13, v22, v7+v28, v7+int32(8), v7, int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						if v35 != 0 {
							v37 = v28
						} else {
							v37 = int32(16)
						}
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[27])))
						F_addReply(m, l0, v39)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
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
func F_hexpireatCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_hexpireGenericCommand(m, l0, int64(0), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_hexpiretimeCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_httlGenericCommand(m, l0, int64(0), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_hexval(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v4 = l0 + int32(-48)
	if base.Ui32(v4) < base.Ui32(int32(10)) {
		v18 = v4
	} else {
		v8 = l0 | int32(32)
		if base.Ui32(v8+int32(-97)) < base.Ui32(int32(6)) {
			v16 = v8 + int32(-87)
		} else {
			v16 = int32(-1)
		}
		v18 = v16
	}
	return v18
}
func F_hlenCommand(m *base.Module, l0 int32) {
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
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v7 = F_lookupKeyReadOrReply(m, l0, v4, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		if v7 == int32(0) {
			return
		} else {
			v12 = F_checkType(m, l0, v7, int32(4))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if v12 != 0 {
					return
				} else {
					v14 = F_hashTypeLength(m, v7)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v14))
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
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
func F_hmgetCommand(m *base.Module, l0 int32) {
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = F_lookupKeyRead(m, v4, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	return
L3:
	;
	v10 = F_checkType(m, l0, v7, int32(4))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v10 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_addReplyArrayLen(m, l0, v12+int32(-2))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v17 < int32(3) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v7 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	v23 = int32(2)
	goto L9
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v23<<(uint(int32(2))%32))))
	v29 = F_objectGetVal(m, v28)
	mBase = m.M
	F_addHashFieldToReply(m, l0, v7, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L11
	}
L10:
	;
	goto L7
L11:
	;
	v33 = v23 + int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v33 < v34 {
		v23 = v33
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v41 = F_hashTypeLength(m, v7)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v41 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v46 = F_dbDelete(m, v43, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	goto L1
}
func F_hookf(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v79 int64
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v111 int32
	_ = v111
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int64
	_ = v158
	var v162 int32
	_ = v162
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	v3 = m.G3
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v3 + int32(_a2700)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11 + int32(16)
	switch int32(2) {
	case 0:
		v68 = l0 + int32(72)
	case 1:
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v44
		v68 = l0 + int32(88)
	case 2:
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v68 = v38 + int32(96)
	default:
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
		v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+7)))
		v56 = m.G398
		if base.Ui32(v55) < base.Ui32(int32(-2)) {
			v67 = v56
		} else {
			v67 = v54 + int32(-24)
		}
		v68 = v67
	}
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v73 = int32(-16)
	v75 = F_luaH_get(m, v71, v72+v73)
	mBase = m.M
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
	*(*int64)(unsafe.Add(mBase, uint32(v76+v73))) = v79
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v76+int32(-8)))) = v83
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = l0
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v90 + int32(16)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v111+int32(-32))))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v152 = int32(-16)
	v154 = F_luaH_get(m, v150, v151+v152)
	mBase = m.M
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v154)))
	*(*int64)(unsafe.Add(mBase, uint32(v155+v152))) = v158
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v155+int32(-8)))) = v162
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v181 = v178 + int32(-16)
	v215 = m.G398
	if v181 != v215 {
		v218 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
		v221 = v218
	} else {
		v221 = int32(-1)
	}
	if v221 != int32(6) {
		return
	} else {
		v224 = m.G3
		v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v231 = *(*int32)(unsafe.Add(mBase, uint32(v224+int32(_a2701)+v227<<(uint(int32(2))%32))))
		F_lua_pushstring(m, l0, v231)
		mBase = m.M
		v233 = m.ExcPending
		if v233 != 0 {
			return
		} else {
			v234 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			if v234 < int32(0) {
				v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v248)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v248 + int32(16)
			} else {
				v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v238)+8)) = int32(3)
				*(*float64)(unsafe.Add(mBase, uint32(v238))) = base.F64_convert_i32_s(v234)
				v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v243 + int32(16)
			}
			F_lua_call(m, l0, int32(2), int32(0))
			mBase = m.M
			v257 = m.ExcPending
			if v257 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_hpersistCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v70 int64
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
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
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int64
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v169 int32
	_ = v169
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v17 = F_objectGetVal(m, v16)
	mBase = m.M
	v18 = int32(_a2361)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v68 = F_getLongLongFromObjectOrReply(m, l0, v64, v11+int32(8), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L16
	} else {
		goto L18
	}
L3:
	;
	if v53-v55 == int32(0) {
		goto L2
	} else {
		goto L15
	}
L4:
	;
	v53 = F_tolower(m, v49)
	mBase = m.M
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v55 = F_tolower(m, v54)
	mBase = m.M
	goto L3
L5:
	;
	v23 = v17
	v24 = v18
	v25 = v21
	goto L8
L6:
	;
	v49 = int32(0)
	v50 = v18
	goto L4
L7:
	;
	v49 = v46 & int32(255)
	v50 = v45
	goto L4
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v27 == int32(0) {
		v45 = v24
		v46 = v25
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v45 = v39
	v46 = int32(0)
	goto L7
L10:
	;
	v31 = v25 & int32(255)
	if v31 == v27 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = int32(1)
	v39 = v24 + v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v40 != 0 {
		v23 = v23 + v38
		v24 = v39
		v25 = v40
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v33 = F_tolower(m, v31)
	mBase = m.M
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v35 = F_tolower(m, v34)
	mBase = m.M
	if v33 == v35 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v45 = v24
	v46 = v37
	goto L7
L14:
	;
	goto L9
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	goto L1
L18:
	;
	if v68 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v70 == int64(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v85 = F_lookupKeyWrite(m, v82, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L16
	} else {
		goto L25
	}
L21:
	;
	F_addReplyError(m, l0, int32(_a2360))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L24
	}
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v70 == base.I64_extend_i32_s(v73+int32(-4)) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L1
L25:
	;
	v88 = F_checkType(m, l0, v85, int32(4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_initDeferredReplyBuffer(m, l0)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	F_addReplyArrayLen(m, l0, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L16
	} else {
		goto L29
	}
L29:
	;
	v96 = F_hashTypeHasVolatileFields(m, v85)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v98 <= int64(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_commitDeferredReplyBuffer(m, l0, int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L16
	} else {
		goto L47
	}
L32:
	;
	v104 = int64(0)
	v105 = int32(4)
	v108 = int32(0)
	goto L33
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v105<<(uint(int32(2))%32))))
	v115 = F_objectGetVal(m, v114)
	mBase = m.M
	v116 = F_hashTypePersist(m, v85, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L16
	} else {
		goto L36
	}
L34:
	;
	if v128 == int32(0) {
		goto L31
	} else {
		goto L40
	}
L35:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v116))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L16
	} else {
		goto L38
	}
L36:
	;
	if v116 != int32(1) {
		v128 = v108
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v120 = int32(_a20)
	v122 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	*(*int64)(unsafe.Add(mBase, _consts[180])) = v122 + int64(1)
	v128 = v108 + int32(1)
	goto L35
L38:
	;
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v136 = v104 + int64(1)
	if v136 < v134 {
		v104 = v136
		v105 = v105 + int32(1)
		v108 = v128
		goto L33
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	v140 = F_hashTypeHasVolatileFields(m, v85)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L16
	} else {
		goto L42
	}
L41:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+28))
	F_notifyKeyspaceEvent(m, int32(64), int32(_a2370), v149, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L16
	} else {
		goto L45
	}
L42:
	;
	if v96 == v140 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbUpdateObjectWithVolatileItemsTracking(m, v143, v85)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	F_signalModifiedKey(m, l0, v154, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	goto L31
L47:
	;
	goto L1
}
func F_hpexpiretimeCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_httlGenericCommand(m, l0, int64(0), int32(1))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_hrandfieldCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
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
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
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
	var v97 int32
	_ = v97
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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 < int32(3) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return
L2:
	;
	F_addReplyError(m, l0, int32(_a2375))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L46
	}
L3:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91<<(uint(int32(2))%32))+uint32(_consts[339])))
	v96 = F_lookupKeyReadOrReply(m, l0, v89, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L33
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v16 = F_getRangeLongFromObjectOrReply(m, l0, v12, int32(-2147483647), int32(2147483647), v6, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	if v16 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(4) < v18 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v76 == int32(4) {
		goto L28
	} else {
		goto L29
	}
L9:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L27
	}
L10:
	;
	if v18 == int32(4) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v29 = F_objectGetVal(m, v28)
	mBase = m.M
	v30 = int32(_a2376)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v33 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	F_hrandfieldWithCountCommand(m, l0, v23, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	goto L1
L14:
	;
	if v65-v67 == int32(0) {
		goto L8
	} else {
		goto L26
	}
L15:
	;
	v65 = F_tolower(m, v61)
	mBase = m.M
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v67 = F_tolower(m, v66)
	mBase = m.M
	goto L14
L16:
	;
	v35 = v29
	v36 = v30
	v37 = v33
	goto L19
L17:
	;
	v61 = int32(0)
	v62 = v30
	goto L15
L18:
	;
	v61 = v58 & int32(255)
	v62 = v57
	goto L15
L19:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v39 == int32(0) {
		v57 = v36
		v58 = v37
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v57 = v51
	v58 = int32(0)
	goto L18
L21:
	;
	v43 = v37 & int32(255)
	if v43 == v39 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v50 = int32(1)
	v51 = v36 + v50
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	if v52 != 0 {
		v35 = v35 + v50
		v36 = v51
		v37 = v52
		goto L19
	} else {
		goto L25
	}
L23:
	;
	v45 = F_tolower(m, v43)
	mBase = m.M
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v47 = F_tolower(m, v46)
	mBase = m.M
	if v45 == v47 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v57 = v36
	v58 = v49
	goto L18
L25:
	;
	goto L20
L26:
	;
	goto L9
L27:
	;
	goto L1
L28:
	;
	if base.Ui32(v75+int32(-1073741824)) <= base.Ui32(int32(-2147483648)) {
		goto L2
	} else {
		goto L31
	}
L29:
	;
	F_hrandfieldWithCountCommand(m, l0, v75, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	goto L1
L31:
	;
	F_hrandfieldWithCountCommand(m, l0, v75, int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	goto L1
L33:
	;
	if v96 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v101 = F_checkType(m, l0, v96, int32(4))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	if v101 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v103 = F_hashTypeLength(m, v96)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L38
	}
L37:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L45
	}
L38:
	;
	v106 = F_hashTypeRandomElement(m, v96, v103, v6, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	if v106 != 0 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v108 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	F_addReplyBulkLongLong(m, l0, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L44
	}
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	F_addReplyBulkCBuffer(m, l0, v108, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	goto L1
L44:
	;
	goto L1
L45:
	;
	goto L1
L46:
	;
	goto L1
}
func F_hrandfieldWithCountCommand(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
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
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
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
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
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
	var v361 int32
	_ = v361
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int64
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
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v703 int32
	_ = v703
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	v14 = m.G0
	v16 = v14 - int32(720)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v21 = *(*int32)(unsafe.Add(mBase, _consts[861]))
	v22 = F_lookupKeyReadOrReply(m, l0, v19, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F__serverAssert(m, int32(_a2371), int32(_a2349), int32(2294))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L5
	} else {
		goto L227
	}
L2:
	;
	F__serverAssert(m, int32(_a2372), int32(_a2349), int32(2291))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L5
	} else {
		goto L226
	}
L3:
	;
	F__serverAssert(m, int32(_a2373), int32(_a2349), int32(2264))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L5
	} else {
		goto L225
	}
L4:
	;
	m.G0 = v16 + int32(720)
	return
L5:
	;
	return
L6:
	;
	if v22 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v27 = F_checkType(m, l0, v22, int32(4))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v27 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v29 = F_hashTypeLength(m, v22)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if l1 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = F_prepareClientForFutureWrites(m, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[861]))
	F_addReply(m, l0, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	goto L4
L14:
	;
	if v35 == int32(0) {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v40 = l1 >> (uint(int32(31)) % 32)
	v42 = l1 ^ v40 - v40
	v44 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	if l1 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if l2 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L18:
	;
	if base.Ui32(l1) < base.Ui32(v29) {
		goto L58
	} else {
		goto L59
	}
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	switch int32(base.Ui32(v50)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		goto L23
	default:
		v703 = int32(0)
		goto L17
	case 9:
		goto L22
	}
L20:
	;
	if v42 != int32(1) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v109 = int32(1000)
	if base.Ui32(v42) < base.Ui32(v109) {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	v58 = l1 >> (uint(int32(31)) % 32)
	v69 = int32(0)
	v70 = v42
	goto L24
L24:
	;
	v79 = F_hashTypeRandomElement(m, v22, v29, v16+int32(16), v16+int32(704))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L26
	}
L25:
	;
	v703 = l1 ^ v58 - v58
	goto L17
L26:
	;
	if v79 != 0 {
		v703 = v69
		goto L17
	} else {
		goto L27
	}
L27:
	;
	if l2 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
	if v101&int32(4) != 0 {
		v703 = v69
		goto L17
	} else {
		goto L37
	}
L29:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	F_addWritePreparedReplyBulkCBuffer(m, v35, v97, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L36
	}
L30:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v83) < base.Ui32(int32(3)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	F_addWritePreparedReplyBulkCBuffer(m, v35, v89, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L34
	}
L32:
	;
	F_addWritePreparedReplyArrayLen(m, v35, int32(2))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v16)+704))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v16)+708))
	F_addWritePreparedReplyBulkCBuffer(m, v35, v93, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	goto L28
L36:
	;
	goto L28
L37:
	;
	v107 = v70 + int32(-1)
	if v107 != 0 {
		v69 = v69 + int32(1)
		v70 = v107
		goto L24
	} else {
		goto L38
	}
L38:
	;
	goto L25
L39:
	;
	v112 = v42
	goto L41
L40:
	;
	v112 = v109
	goto L41
L41:
	;
	v114 = v112 << (uint(int32(4)) % 32)
	v115 = F_valkey_malloc(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v117 = int32(0)
	if l2 == v117 {
		v122 = v117
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v130 = int32(0)
	v131 = v42
	goto L47
L44:
	;
	v120 = F_valkey_malloc(m, v114)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v122 = v120
	goto L43
L46:
	;
	F_valkey_free(m, v115)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L5
	} else {
		goto L56
	}
L47:
	;
	v136 = F_objectGetVal(m, v22)
	mBase = m.M
	if base.Ui32(v131) < base.Ui32(v112) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L46
L49:
	;
	v138 = v131
	goto L51
L50:
	;
	v138 = v112
	goto L51
L51:
	;
	F_lpRandomPairs(m, v136, v138, v115, v122)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	F_hrandfieldReplyWithListpack(m, v35, v138, v115, v122)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	v143 = v138 + v130
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
	if v144&int32(4) != 0 {
		goto L46
	} else {
		goto L54
	}
L54:
	;
	v147 = v131 - v138
	if v147 != 0 {
		v130 = v143
		v131 = v147
		goto L47
	} else {
		goto L55
	}
L55:
	;
	goto L48
L56:
	;
	F_valkey_free(m, v122)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v703 = v143
	goto L17
L58:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v231&int32(240) != int32(176) {
		goto L80
	} else {
		goto L81
	}
L59:
	;
	F_hashTypeInitIterator(m, v22, v16+int32(16))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	v158 = int32(0)
	v161 = F_hashTypeNext(m, v16+int32(16))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L62
	}
L61:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v219 != int32(2) {
		v703 = v213
		goto L17
	} else {
		goto L75
	}
L62:
	;
	if v161 == int32(-1) {
		v213 = v158
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v172 = v158
	goto L64
L64:
	;
	if l2 == int32(0) {
		v193 = int32(1)
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v213 = v199
	goto L61
L66:
	;
	F_addHashIteratorCursorToReply(m, v35, v16+int32(16), v193)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L5
	} else {
		goto L72
	}
L67:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v181) < base.Ui32(int32(3)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	F_addHashIteratorCursorToReply(m, v35, v16+int32(16), int32(1))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	F_addWritePreparedReplyArrayLen(m, v35, int32(2))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v193 = int32(2)
	goto L66
L72:
	;
	v199 = v172 + int32(1)
	v202 = F_hashTypeNext(m, v16+int32(16))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	if v202 != int32(-1) {
		v172 = v199
		goto L64
	} else {
		goto L74
	}
L74:
	;
	goto L65
L75:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+24)))
	if v222 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_vsetResetIterator(m, v16+int32(88))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	F_hashtableCleanupIterator(m, v16+int32(40))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v703 = v213
	goto L17
L79:
	;
	v703 = v213
	goto L17
L80:
	;
	if base.Ui32(l1*int32(3)) <= base.Ui32(v29) {
		goto L92
	} else {
		goto L93
	}
L81:
	;
	v237 = l1 << (uint(int32(4)) % 32)
	v238 = F_valkey_malloc(m, v237)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	if l2 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v244 = F_objectGetVal(m, v22)
	mBase = m.M
	v245 = F_lpRandomPairsUnique(m, v244, l1, v238, v243)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L5
	} else {
		goto L87
	}
L84:
	;
	v241 = F_valkey_malloc(m, v237)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L5
	} else {
		goto L86
	}
L85:
	;
	v243 = int32(0)
	goto L83
L86:
	;
	v243 = v241
	goto L83
L87:
	;
	if v245 != l1 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	F_hrandfieldReplyWithListpack(m, v35, l1, v238, v243)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	F_valkey_free(m, v238)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	F_valkey_free(m, v243)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	v703 = l1
	goto L17
L92:
	;
	v617 = F_hashtableCreate(m, int32(_a1844))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L5
	} else {
		goto L186
	}
L93:
	;
	v258 = F_hashtableCreate(m, int32(_a2374))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v260 = F_hashtableExpand(m, v258, v29)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	v262 = int32(0)
	v264 = v16 + int32(16)
	v265 = F_objectGetVal(m, v22)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+14)) = uint8(v262)
	*(*int32)(unsafe.Add(mBase, uint32(v264))) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v264)+24)) = v262
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+15)) = uint8(v262)
	*(*int32)(unsafe.Add(mBase, uint32(v264)+8)) = int32(-1)
	if v265 == v262 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v288 = F_hashtableNext(m, v16+int32(16), v16+int32(704))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L5
	} else {
		goto L101
	}
L97:
	;
	goto L96
L98:
	;
	goto L97
L100:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v258)+20))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v258)+16))
	goto L109
L101:
	;
	if v288 == int32(0) {
		v325 = v262
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v299 = v262
	goto L103
L103:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v16)+704))
	v306 = F_hashtableAdd(m, v258, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L105
	}
L104:
	;
	v325 = v311
	goto L100
L105:
	;
	if v306 == int32(0) {
		goto L2
	} else {
		goto L106
	}
L106:
	;
	v311 = v299 + int32(1)
	v316 = F_hashtableNext(m, v16+int32(16), v16+int32(704))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	if v316 != 0 {
		v299 = v311
		goto L103
	} else {
		goto L108
	}
L108:
	;
	goto L104
L109:
	;
	if v331+v332 != v325 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_hashtableCleanupIterator(m, v16+int32(16))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	if base.Ui32(v325) <= base.Ui32(v42) {
		v370 = v325
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v377 = v16 + int32(16)
	v378 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v377)+14)) = uint8(v378)
	*(*int32)(unsafe.Add(mBase, uint32(v377))) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v377)+24)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v377)+15)) = uint8(v378)
	*(*int32)(unsafe.Add(mBase, uint32(v377)+8)) = int32(-1)
	if v258 == v378 {
		goto L120
	} else {
		goto L121
	}
L113:
	;
	v347 = v325
	goto L114
L114:
	;
	v355 = F_hashtableFairRandomEntry(m, v258, v16+int32(12))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L5
	} else {
		goto L116
	}
L115:
	;
	v370 = l1
	goto L112
L116:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v358 = F_hashtableDelete(m, v258, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v361 = v347 + int32(-1)
	if base.Ui32(v42) < base.Ui32(v361) {
		v347 = v361
		goto L114
	} else {
		goto L118
	}
L118:
	;
	goto L115
L119:
	;
	v400 = F_hashtableNext(m, v16+int32(16), v16+int32(12))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L5
	} else {
		goto L124
	}
L120:
	;
	goto L119
L121:
	;
	goto L120
L123:
	;
	F_hashtableCleanupIterator(m, v16+int32(16))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L5
	} else {
		goto L184
	}
L124:
	;
	if v400 == int32(0) {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	goto L126
L126:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	goto L128
L127:
	;
	goto L123
L128:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v420 = v16 + int32(8)
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418+int32(-1)))))
	v428 = v426 & int32(7)
	if v428 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	if l2 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L130:
	;
	v530 = v522
	goto L129
L131:
	;
	v483 = F_sdsAllocPtr(m, v418)
	mBase = m.M
	v485 = v483 + int32(-4)
	if v426&int32(32) == int32(0) {
		goto L148
	} else {
		goto L149
	}
L132:
	;
	switch v428 {
	case 0:
		goto L140
	case 1:
		goto L139
	case 2:
		goto L138
	case 3:
		goto L137
	case 4:
		goto L136
	default:
		v448 = int32(0)
		goto L135
	}
L133:
	;
	if v426&int32(16) != 0 {
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	v450 = int32(1)
	v451 = F_sdsHdrSize(m, v450)
	mBase = m.M
	v452 = v418 + v448 + v451
	v454 = v452 + v450
	if v420 == int32(0) {
		v522 = v454
		goto L130
	} else {
		goto L141
	}
L136:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v418+int32(-17))))
	v448 = v447
	goto L135
L137:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v418+int32(-9))))
	v448 = v444
	goto L135
L138:
	;
	v441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v418+int32(-5)))))
	v448 = v441
	goto L135
L139:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418+int32(-3)))))
	v448 = v438
	goto L135
L140:
	;
	v448 = int32(base.Ui32(v426) >> (uint(int32(3)) % 32))
	goto L135
L141:
	;
	v457 = int32(0)
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452+v457))))
	switch v460 & int32(7) {
	case 0:
		goto L147
	case 1:
		goto L146
	case 2:
		goto L145
	case 3:
		goto L144
	case 4:
		goto L143
	default:
		v481 = v457
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v481
	v530 = v454
	goto L129
L143:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v452+int32(-16))))
	v481 = v480
	goto L142
L144:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v452+int32(-8))))
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v476
	v530 = v454
	goto L129
L145:
	;
	v472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452+int32(-4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v472
	v530 = v454
	goto L129
L146:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452+int32(-2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v468
	v530 = v454
	goto L129
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = int32(base.Ui32(v460) >> (uint(int32(3)) % 32))
	v530 = v454
	goto L129
L148:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	if v420 == int32(0) {
		v522 = v497
		goto L130
	} else {
		goto L154
	}
L149:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	if v490 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if v420 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v530 = int32(0)
	goto L129
L152:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	v530 = v496
	goto L129
L153:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v490)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v494
	goto L152
L154:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497+int32(-1)))))
	switch v503 & int32(7) {
	case 0:
		goto L160
	case 1:
		goto L159
	case 2:
		goto L158
	case 3:
		goto L157
	case 4:
		goto L156
	default:
		v520 = int32(0)
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v520
	v522 = v497
	goto L130
L156:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v497+int32(-17))))
	v520 = v519
	goto L155
L157:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v497+int32(-9))))
	v520 = v516
	goto L155
L158:
	;
	v513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v497+int32(-5)))))
	v520 = v513
	goto L155
L159:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497+int32(-3)))))
	v520 = v510
	goto L155
L160:
	;
	v520 = int32(base.Ui32(v503) >> (uint(int32(3)) % 32))
	goto L155
L161:
	;
	v595 = F_hashtableNext(m, v16+int32(16), v16+int32(12))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L5
	} else {
		goto L182
	}
L162:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+int32(-1)))))
	switch v568 & int32(7) {
	case 0:
		goto L180
	case 1:
		goto L179
	case 2:
		goto L178
	case 3:
		goto L177
	case 4:
		goto L176
	default:
		v585 = int32(0)
		goto L175
	}
L163:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v533) < base.Ui32(int32(3)) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+int32(-1)))))
	switch v542 & int32(7) {
	case 0:
		goto L172
	case 1:
		goto L171
	case 2:
		goto L170
	case 3:
		goto L169
	case 4:
		goto L168
	default:
		v559 = int32(0)
		goto L167
	}
L165:
	;
	F_addWritePreparedReplyArrayLen(m, v35, int32(2))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L5
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	F_addWritePreparedReplyBulkCBuffer(m, v35, v417, v559)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L5
	} else {
		goto L173
	}
L168:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v417+int32(-17))))
	v559 = v558
	goto L167
L169:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v417+int32(-9))))
	v559 = v555
	goto L167
L170:
	;
	v552 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v417+int32(-5)))))
	v559 = v552
	goto L167
L171:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+int32(-3)))))
	v559 = v549
	goto L167
L172:
	;
	v559 = int32(base.Ui32(v542) >> (uint(int32(3)) % 32))
	goto L167
L173:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	F_addWritePreparedReplyBulkCBuffer(m, v35, v530, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L5
	} else {
		goto L174
	}
L174:
	;
	goto L161
L175:
	;
	F_addWritePreparedReplyBulkCBuffer(m, v35, v417, v585)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L5
	} else {
		goto L181
	}
L176:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v417+int32(-17))))
	v585 = v584
	goto L175
L177:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v417+int32(-9))))
	v585 = v581
	goto L175
L178:
	;
	v578 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v417+int32(-5)))))
	v585 = v578
	goto L175
L179:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+int32(-3)))))
	v585 = v575
	goto L175
L180:
	;
	v585 = int32(base.Ui32(v568) >> (uint(int32(3)) % 32))
	goto L175
L181:
	;
	goto L161
L182:
	;
	if v595 != 0 {
		goto L126
	} else {
		goto L183
	}
L183:
	;
	goto L127
L184:
	;
	F_hashtableRelease(m, v258)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L5
	} else {
		goto L185
	}
L185:
	;
	v703 = v370
	goto L17
L186:
	;
	v619 = F_hashtableExpand(m, v617, l1)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L5
	} else {
		goto L187
	}
L187:
	;
	if base.Ui32(int32(429496729)) < base.Ui32(l1) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v626 = int32(-1)
	goto L190
L189:
	;
	v626 = l1 * int32(10)
	goto L190
L190:
	;
	v627 = int32(0)
	if l2 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v631 = v16 + int32(704)
	goto L193
L192:
	;
	v631 = v627
	goto L193
L193:
	;
	v633 = v626
	v639 = v627
	goto L195
L194:
	;
	F_hashtableRelease(m, v617)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L5
	} else {
		goto L219
	}
L195:
	;
	if v633 == int32(0) {
		v692 = v639
		goto L194
	} else {
		goto L197
	}
L196:
	;
	v692 = v687
	goto L194
L197:
	;
	v649 = F_hashTypeRandomElement(m, v22, v29, v16+int32(16), v631)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	if v649 != 0 {
		v692 = v639
		goto L194
	} else {
		goto L199
	}
L199:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v651 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v661 = F_hashtableAdd(m, v617, v660)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L5
	} else {
		goto L207
	}
L201:
	;
	v657 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
	v658 = F_sdsfromlonglong(m, v657)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L5
	} else {
		goto L204
	}
L202:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v655 = F_sdsnewlen(m, v651, v654)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L5
	} else {
		goto L203
	}
L203:
	;
	v660 = v655
	goto L200
L204:
	;
	v660 = v658
	goto L200
L205:
	;
	if base.Ui32(v687) < base.Ui32(v42) {
		v633 = v633 + int32(-1)
		v639 = v687
		goto L195
	} else {
		goto L218
	}
L206:
	;
	v666 = v639 + int32(1)
	if l2 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L207:
	;
	if v661 != 0 {
		goto L206
	} else {
		goto L208
	}
L208:
	;
	F_sdsfree(m, v660)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L5
	} else {
		goto L209
	}
L209:
	;
	v687 = v639
	goto L205
L210:
	;
	F_hashReplyFromListpackEntry(m, l0, v16+int32(16))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L5
	} else {
		goto L217
	}
L211:
	;
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v669) < base.Ui32(int32(3)) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	F_hashReplyFromListpackEntry(m, l0, v16+int32(16))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L5
	} else {
		goto L215
	}
L213:
	;
	F_addWritePreparedReplyArrayLen(m, v35, int32(2))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L5
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	F_hashReplyFromListpackEntry(m, l0, v16+int32(704))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L5
	} else {
		goto L216
	}
L216:
	;
	v687 = v666
	goto L205
L217:
	;
	v687 = v666
	goto L205
L218:
	;
	goto L196
L219:
	;
	v703 = v692
	goto L17
L220:
	;
	F_setDeferredArrayLen(m, l0, v44, v703)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L5
	} else {
		goto L224
	}
L221:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v711 != int32(2) {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	F_setDeferredArrayLen(m, l0, v44, v703<<(uint(int32(1))%32))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L5
	} else {
		goto L223
	}
L223:
	;
	goto L4
L224:
	;
	goto L4
L225:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L227:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hscanCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v10 = F_objectGetVal(m, v9)
	mBase = m.M
	v13 = F_parseScanCursorOrReply(m, l0, v10, v6+int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 == int32(-1) {
			m.G0 = v6 + int32(16)
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v20 = *(*int32)(unsafe.Add(mBase, _consts[860]))
			v21 = F_lookupKeyReadOrReply(m, l0, v18, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if v21 == int32(0) {
					m.G0 = v6 + int32(16)
					return
				} else {
					v26 = F_checkType(m, l0, v21, int32(4))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						if v26 != 0 {
							m.G0 = v6 + int32(16)
							return
						} else {
							v28 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
							F_scanGenericCommand(m, l0, v21, v28)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_hsetexCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
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
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
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
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v173 int64
	_ = v173
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int64
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
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
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
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
	var v440 int32
	_ = v440
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
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
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
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
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
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v626 int64
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v740 int64
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int64
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v877 int64
	_ = v877
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v965 int64
	_ = v965
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1080 int32
	_ = v1080
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(64)
	m.G0 = v19
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(-1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v34 < int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v19 + int32(64)
	return
L2:
	;
	F_addReplyError(m, l0, int32(_a2360))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L23
	} else {
		goto L241
	}
L3:
	;
	v39 = v2
	goto L7
L4:
	;
	v386 = F_hashTypeHasVolatileFields(m, v374)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L23
	} else {
		goto L84
	}
L5:
	;
	v359 = F_createHashObject(m)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L23
	} else {
		goto L82
	}
L6:
	;
	if v215&int32(4096) == int32(0) {
		goto L5
	} else {
		goto L80
	}
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v39<<(uint(int32(2))%32))))
	v58 = F_objectGetVal(m, v57)
	mBase = m.M
	v59 = int32(_a2361)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v62 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v330 = v39 + int32(1)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v330 < v331+int32(-1) {
		v39 = v330
		goto L7
	} else {
		goto L79
	}
L10:
	;
	if v94-v96 != 0 {
		goto L9
	} else {
		goto L22
	}
L11:
	;
	v94 = F_tolower(m, v90)
	mBase = m.M
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v96 = F_tolower(m, v95)
	mBase = m.M
	goto L10
L12:
	;
	v64 = v58
	v65 = v59
	v66 = v62
	goto L15
L13:
	;
	v90 = int32(0)
	v91 = v59
	goto L11
L14:
	;
	v90 = v87 & int32(255)
	v91 = v86
	goto L11
L15:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v68 == int32(0) {
		v86 = v65
		v87 = v66
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v86 = v80
	v87 = int32(0)
	goto L14
L17:
	;
	v72 = v66 & int32(255)
	if v72 == v68 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v79 = int32(1)
	v80 = v65 + v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v81 != 0 {
		v64 = v64 + v79
		v65 = v80
		v66 = v81
		goto L15
	} else {
		goto L21
	}
L19:
	;
	v74 = F_tolower(m, v72)
	mBase = m.M
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v76 = F_tolower(m, v75)
	mBase = m.M
	if v74 == v76 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v86 = v65
	v87 = v78
	goto L14
L21:
	;
	goto L16
L22:
	;
	v109 = F_parseExtendedCommandArgumentsOrReply(m, l0, int32(3), int32(2), v39, v19+int32(28), v19+int32(32), int32(0), v19+int32(40), v19+int32(36))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return
L24:
	;
	if v109 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+v39<<(uint(int32(2))%32))+4))
	v119 = F_getLongLongFromObjectOrReply(m, l0, v115, v19+int32(16), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	if v119 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	if v121 == int64(0) {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	if int64(4611686018427387903) < v121 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v130 = v39 + int32(2)
	if v121<<(uint(int64(1))%64) != base.I64_extend_i32_s(v128-v130) {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v137 = F_lookupKeyWrite(m, v134, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v137
	v141 = F_checkType(m, l0, v137, int32(4))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	if v141 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if v144&int32(1) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v160 = base.B2i32(v144&int32(6223) != int32(0))
	v162 = int32(1)
	if v144&int32(16) != 0 {
		v214 = v143
		v215 = v144
		v216 = v160
		v217 = int32(4)
		v218 = v162
		goto L42
	} else {
		goto L43
	}
L35:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	F_addReply(m, l0, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L23
	} else {
		goto L41
	}
L36:
	;
	if v144&int32(2) == int32(0) {
		goto L34
	} else {
		goto L39
	}
L37:
	;
	if v143 != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	if v143 != 0 {
		goto L34
	} else {
		goto L40
	}
L40:
	;
	goto L35
L41:
	;
	goto L1
L42:
	;
	if v215&int32(6144) == int32(0) {
		v317 = v214
		goto L61
	} else {
		goto L62
	}
L43:
	;
	v165 = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	if v166 == v165 {
		v214 = v143
		v215 = v144
		v216 = v160
		v217 = v165
		v218 = v162
		goto L42
	} else {
		goto L44
	}
L44:
	;
	if v144&int32(192) != 0 {
		v175 = int64(0)
		v176 = v166
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v180 = F_convertExpireArgumentToUnixTime(m, l0, v176, v175, v177, v19+int32(8))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L23
	} else {
		goto L48
	}
L46:
	;
	v173 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	goto L47
L47:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v175 = v173
	v176 = v174
	goto L45
L48:
	;
	if v180 == int32(-1) {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	v189 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v189 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v207 != 0 {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	goto L50
L52:
	;
	v195 = int32(0)
	v196 = F_commandTimeSnapshot(m)
	mBase = m.M
	if v196 < v185 {
		v207 = v195
		goto L51
	} else {
		goto L55
	}
L53:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189)+216))
	if v193 != 0 {
		v207 = int32(0)
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v198 = int32(_a20)
	v199 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v201 = *(*int32)(unsafe.Add(mBase, _consts[384]))
	if v201 != 0 {
		v207 = v195
		goto L51
	} else {
		goto L56
	}
L56:
	;
	if v199 != 0 {
		v207 = v195
		goto L51
	} else {
		goto L57
	}
L57:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _consts[505]))
	v207 = base.B2i32(v203 == int32(0))
	goto L51
L58:
	;
	v208 = int32(1)
	goto L60
L59:
	;
	v208 = v160
	goto L60
L60:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v214 = v211
	v215 = v212
	v216 = v208
	v217 = v165
	v218 = base.B2i32(v207 == int32(0))
	goto L42
L61:
	;
	if v317 != 0 {
		v374 = v317
		goto L4
	} else {
		goto L78
	}
L62:
	;
	if v214 == int32(0) {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v226 <= v130 {
		v374 = v214
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v232 = v130
	goto L65
L65:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if v244&int32(2048) == int32(0) {
		v272 = v244
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v317 = v312
	goto L61
L67:
	;
	v309 = v232 + int32(2)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v309 < v310 {
		v232 = v309
		goto L65
	} else {
		goto L77
	}
L68:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	F_addReply(m, l0, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L23
	} else {
		goto L76
	}
L69:
	;
	if v272&int32(4096) == int32(0) {
		goto L67
	} else {
		goto L73
	}
L70:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v232<<(uint(int32(2))%32))))
	v255 = F_objectGetVal(m, v254)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = int64(9223372036854775807)
	v267 = F_hashTypeGetValue(m, v249, v255, v19+int32(60), v19+int32(56), v19+int32(48), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L23
	} else {
		goto L71
	}
L71:
	;
	if v267 == int32(0) {
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v272 = v271
	goto L69
L73:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279+v232<<(uint(int32(2))%32))))
	v284 = F_objectGetVal(m, v283)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = int64(9223372036854775807)
	v296 = F_hashTypeGetValue(m, v278, v284, v19+int32(60), v19+int32(56), v19+int32(48), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L23
	} else {
		goto L74
	}
L74:
	;
	if v296 == int32(0) {
		goto L67
	} else {
		goto L75
	}
L75:
	;
	goto L68
L76:
	;
	goto L1
L77:
	;
	goto L66
L78:
	;
	goto L5
L79:
	;
	goto L2
L80:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	F_addReply(m, l0, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L23
	} else {
		goto L81
	}
L81:
	;
	goto L1
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v359
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
	F_dbAdd(m, v362, v364, v19+int32(44))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L23
	} else {
		goto L83
	}
L83:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v374 = v369
	goto L4
L84:
	;
	if v218 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v678 <= v130 {
		v983 = int32(0)
		goto L158
	} else {
		goto L159
	}
L86:
	;
	if v216 != 0 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v388 = int32(2)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v394 = F_valkey_malloc(m, v389<<(uint(v388)%32)+int32(8))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L23
	} else {
		goto L88
	}
L88:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _consts[551]))
	*(*int32)(unsafe.Add(mBase, uint32(v394))) = v397
	F_incrRefCount(m, v397)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L23
	} else {
		goto L89
	}
L89:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v394)+4)) = v402
	F_incrRefCount(m, v402)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L23
	} else {
		goto L90
	}
L90:
	;
	v665 = v388
	v672 = v394
	goto L85
L91:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v412 = F_valkey_malloc(m, v409<<(uint(int32(2))%32))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L23
	} else {
		goto L93
	}
L92:
	;
	v406 = int32(0)
	v665 = v406
	v672 = v406
	goto L85
L93:
	;
	v419 = int32(0)
	v420 = int32(0)
	goto L94
L94:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v433 = v420 << (uint(int32(2)) % 32)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v431+v433)))
	v436 = F_objectGetVal(m, v435)
	mBase = m.M
	v437 = int32(_a2362)
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
	if v440 != 0 {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	v665 = v652
	v672 = v412
	goto L85
L96:
	;
	v659 = v653 + int32(1)
	if v659 < v130 {
		v419 = v652
		v420 = v659
		goto L94
	} else {
		goto L156
	}
L97:
	;
	if v472-v474 == int32(0) {
		v652 = v419
		v653 = v420
		goto L96
	} else {
		goto L109
	}
L98:
	;
	v472 = F_tolower(m, v468)
	mBase = m.M
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469))))
	v474 = F_tolower(m, v473)
	mBase = m.M
	goto L97
L99:
	;
	v442 = v436
	v443 = v437
	v444 = v440
	goto L102
L100:
	;
	v468 = int32(0)
	v469 = v437
	goto L98
L101:
	;
	v468 = v465 & int32(255)
	v469 = v464
	goto L98
L102:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v446 == int32(0) {
		v464 = v443
		v465 = v444
		goto L101
	} else {
		goto L104
	}
L103:
	;
	v464 = v458
	v465 = int32(0)
	goto L101
L104:
	;
	v450 = v444 & int32(255)
	if v450 == v446 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v457 = int32(1)
	v458 = v443 + v457
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442)+1)))
	if v459 != 0 {
		v442 = v442 + v457
		v443 = v458
		v444 = v459
		goto L102
	} else {
		goto L108
	}
L106:
	;
	v452 = F_tolower(m, v450)
	mBase = m.M
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	v454 = F_tolower(m, v453)
	mBase = m.M
	if v452 == v454 {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
	v464 = v443
	v465 = v456
	goto L101
L108:
	;
	goto L103
L109:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v478+v433)))
	v481 = F_objectGetVal(m, v480)
	mBase = m.M
	v482 = int32(_a2363)
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481))))
	if v485 != 0 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	if v517-v519 == int32(0) {
		v652 = v419
		v653 = v420
		goto L96
	} else {
		goto L122
	}
L111:
	;
	v517 = F_tolower(m, v513)
	mBase = m.M
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
	v519 = F_tolower(m, v518)
	mBase = m.M
	goto L110
L112:
	;
	v487 = v481
	v488 = v482
	v489 = v485
	goto L115
L113:
	;
	v513 = int32(0)
	v514 = v482
	goto L111
L114:
	;
	v513 = v510 & int32(255)
	v514 = v509
	goto L111
L115:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	if v491 == int32(0) {
		v509 = v488
		v510 = v489
		goto L114
	} else {
		goto L117
	}
L116:
	;
	v509 = v503
	v510 = int32(0)
	goto L114
L117:
	;
	v495 = v489 & int32(255)
	if v495 == v491 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v502 = int32(1)
	v503 = v488 + v502
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+1)))
	if v504 != 0 {
		v487 = v487 + v502
		v488 = v503
		v489 = v504
		goto L115
	} else {
		goto L121
	}
L119:
	;
	v497 = F_tolower(m, v495)
	mBase = m.M
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	v499 = F_tolower(m, v498)
	mBase = m.M
	if v497 == v499 {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487))))
	v509 = v488
	v510 = v501
	goto L114
L121:
	;
	goto L116
L122:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v523+v433)))
	v526 = F_objectGetVal(m, v525)
	mBase = m.M
	v527 = int32(_a2364)
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	if v530 != 0 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	if v562-v564 == int32(0) {
		v652 = v419
		v653 = v420
		goto L96
	} else {
		goto L135
	}
L124:
	;
	v562 = F_tolower(m, v558)
	mBase = m.M
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559))))
	v564 = F_tolower(m, v563)
	mBase = m.M
	goto L123
L125:
	;
	v532 = v526
	v533 = v527
	v534 = v530
	goto L128
L126:
	;
	v558 = int32(0)
	v559 = v527
	goto L124
L127:
	;
	v558 = v555 & int32(255)
	v559 = v554
	goto L124
L128:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533))))
	if v536 == int32(0) {
		v554 = v533
		v555 = v534
		goto L127
	} else {
		goto L130
	}
L129:
	;
	v554 = v548
	v555 = int32(0)
	goto L127
L130:
	;
	v540 = v534 & int32(255)
	if v540 == v536 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v547 = int32(1)
	v548 = v533 + v547
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532)+1)))
	if v549 != 0 {
		v532 = v532 + v547
		v533 = v548
		v534 = v549
		goto L128
	} else {
		goto L134
	}
L132:
	;
	v542 = F_tolower(m, v540)
	mBase = m.M
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533))))
	v544 = F_tolower(m, v543)
	mBase = m.M
	if v542 == v544 {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532))))
	v554 = v533
	v555 = v546
	goto L127
L134:
	;
	goto L129
L135:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v568+v433)))
	v571 = F_objectGetVal(m, v570)
	mBase = m.M
	v572 = int32(_a2365)
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571))))
	if v575 != 0 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	if v607-v609 == int32(0) {
		v652 = v419
		v653 = v420
		goto L96
	} else {
		goto L148
	}
L137:
	;
	v607 = F_tolower(m, v603)
	mBase = m.M
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604))))
	v609 = F_tolower(m, v608)
	mBase = m.M
	goto L136
L138:
	;
	v577 = v571
	v578 = v572
	v579 = v575
	goto L141
L139:
	;
	v603 = int32(0)
	v604 = v572
	goto L137
L140:
	;
	v603 = v600 & int32(255)
	v604 = v599
	goto L137
L141:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578))))
	if v581 == int32(0) {
		v599 = v578
		v600 = v579
		goto L140
	} else {
		goto L143
	}
L142:
	;
	v599 = v593
	v600 = int32(0)
	goto L140
L143:
	;
	v585 = v579 & int32(255)
	if v585 == v581 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v592 = int32(1)
	v593 = v578 + v592
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577)+1)))
	if v594 != 0 {
		v577 = v577 + v592
		v578 = v593
		v579 = v594
		goto L141
	} else {
		goto L147
	}
L145:
	;
	v587 = F_tolower(m, v585)
	mBase = m.M
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578))))
	v589 = F_tolower(m, v588)
	mBase = m.M
	if v587 == v589 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577))))
	v599 = v578
	v600 = v591
	goto L140
L147:
	;
	goto L142
L148:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	if v613 != 0 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v640+v433)))
	*(*int32)(unsafe.Add(mBase, uint32(v412+v419<<(uint(int32(2))%32)))) = v646
	F_incrRefCount(m, v646)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L23
	} else {
		goto L155
	}
L150:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+28)))
	if v616&int32(128) != 0 {
		v640 = v615
		goto L149
	} else {
		goto L152
	}
L151:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v640 = v614
	goto L149
L152:
	;
	v620 = v420 + int32(1)
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v615+v620<<(uint(int32(2))%32))))
	if v624 != v613 {
		v640 = v615
		goto L149
	} else {
		goto L153
	}
L153:
	;
	v626 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	v627 = F_createStringObjectFromLongLong(m, v626)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L23
	} else {
		goto L154
	}
L154:
	;
	v629 = int32(2)
	v631 = v412 + v419<<(uint(v629)%32)
	v633 = *(*int32)(unsafe.Add(mBase, _consts[859]))
	*(*int32)(unsafe.Add(mBase, uint32(v631))) = v633
	*(*int32)(unsafe.Add(mBase, uint32(v631+int32(4)))) = v627
	v652 = v419 + v629
	v653 = v620
	goto L96
L155:
	;
	v652 = v419 + int32(1)
	v653 = v420
	goto L96
L156:
	;
	goto L95
L157:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1011)+4))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1013)+28))
	F_notifyKeyspaceEvent(m, int32(64), int32(_a2358), v1012, v1014)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L23
	} else {
		goto L225
	}
L158:
	;
	if v218 != 0 {
		goto L220
	} else {
		goto L221
	}
L159:
	;
	v680 = int32(0)
	v685 = v130
	v687 = v665
	v688 = v680
	v696 = v680
	v697 = v680
	goto L160
L160:
	;
	if v218 != 0 {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	if v853 == int32(0) {
		v983 = v857
		goto L158
	} else {
		goto L201
	}
L162:
	;
	v860 = v685 + int32(2)
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v860 < v861 {
		v685 = v860
		v687 = v852
		v688 = v853
		v696 = v856
		v697 = v857
		goto L160
	} else {
		goto L200
	}
L163:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v778 = int32(2)
	v779 = v685 << (uint(v778) % 32)
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v777+v779)))
	v782 = F_objectGetVal(m, v781)
	mBase = m.M
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v787 = (v685 + int32(1)) << (uint(v778) % 32)
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v783+v787)))
	v790 = F_objectGetVal(m, v789)
	mBase = m.M
	v791 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	v794 = F_hashTypeSet(m, v776, v782, v790, v791, v217, v19+int32(48))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L23
	} else {
		goto L187
	}
L164:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v699)))
	if v700&int32(240) != int32(32) {
		v711 = v699
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v714 = v685 << (uint(int32(2)) % 32)
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v712+v714)))
	v717 = F_objectGetVal(m, v716)
	mBase = m.M
	v718 = F_hashTypeDelete(m, v711, v717)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L23
	} else {
		goto L169
	}
L166:
	;
	v705 = F_objectGetVal(m, v699)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v705))) = int32(_a2352)
	goto L167
L167:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v711 = v710
	goto L165
L168:
	;
	v738 = int32(_a20)
	v740 = *(*int64)(unsafe.Add(mBase, _consts[858]))
	*(*int64)(unsafe.Add(mBase, _consts[858])) = v740 + int64(1)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v744)))
	if v745&int32(240) != int32(32) {
		v852 = v735
		v853 = v736
		v856 = v696
		v857 = v697
		goto L162
	} else {
		goto L172
	}
L169:
	;
	if v718 == int32(0) {
		v735 = v687
		v736 = v688
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v725+v714)))
	*(*int32)(unsafe.Add(mBase, uint32(v672+v687<<(uint(int32(2))%32)))) = v727
	F_incrRefCount(m, v727)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L23
	} else {
		goto L171
	}
L171:
	;
	v731 = int32(1)
	v735 = v687 + v731
	v736 = v688 + v731
	goto L168
L172:
	;
	v750 = F_objectGetVal(m, v744)
	mBase = m.M
	v752 = v750 + int32(44)
	goto L173
L173:
	;
	if v752 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	v767 = F_objectGetVal(m, v744)
	mBase = m.M
	if v766 != 0 {
		goto L180
	} else {
		goto L181
	}
L175:
	;
	goto L174
L176:
	;
	v766 = int32(0)
	goto L175
L177:
	;
	v756 = int32(1)
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v752)))
	switch v757 + v756 {
	case 0:
		v766 = v756
		goto L175
	case 1:
		goto L176
	default:
		goto L178
	}
L178:
	;
	if v757&int32(7) != 0 {
		v766 = v756
		goto L175
	} else {
		goto L179
	}
L179:
	;
	goto L176
L180:
	;
	v770 = int32(_a2353)
	goto L182
L181:
	;
	v770 = int32(_a2352)
	goto L182
L182:
	;
	if v752 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v772 = v770
	goto L185
L184:
	;
	v772 = int32(_a2352)
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v767))) = v772
	goto L186
L186:
	;
	v852 = v735
	v853 = v736
	v856 = v696
	v857 = v697
	goto L162
L187:
	;
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	if v796 != int32(1) {
		v823 = v696
		v824 = v697
		goto L188
	} else {
		goto L189
	}
L188:
	;
	if v216 == int32(0) {
		v847 = v687
		goto L196
	} else {
		goto L197
	}
L189:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+28)))
	if v799&int32(16) == int32(0) {
		v819 = v697
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v823 = v696 + int32(1)
	v824 = v819
	goto L188
L191:
	;
	if v697 != 0 {
		v809 = v697
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v813+v779)))
	*(*int32)(unsafe.Add(mBase, uint32(v809+v696<<(uint(int32(2))%32)))) = v815
	F_incrRefCount(m, v815)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L23
	} else {
		goto L195
	}
L193:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v807 = F_valkey_malloc(m, v804<<(uint(int32(2))%32))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L23
	} else {
		goto L194
	}
L194:
	;
	v809 = v807
	goto L192
L195:
	;
	v819 = v809
	goto L190
L196:
	;
	v852 = v847
	v853 = v688 + int32(1)
	v856 = v823
	v857 = v824
	goto L162
L197:
	;
	v830 = v672 + v687<<(uint(int32(2))%32)
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v831+v779)))
	*(*int32)(unsafe.Add(mBase, uint32(v830))) = v833
	F_incrRefCount(m, v833)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L23
	} else {
		goto L198
	}
L198:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v839+v787)))
	*(*int32)(unsafe.Add(mBase, uint32(v830+int32(4)))) = v841
	F_incrRefCount(m, v841)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L23
	} else {
		goto L199
	}
L199:
	;
	v847 = v687 + int32(2)
	goto L196
L200:
	;
	goto L161
L201:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v866 = F_hashTypeHasVolatileFields(m, v865)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L23
	} else {
		goto L203
	}
L202:
	;
	if v856 < int32(1) {
		v952 = v857
		goto L206
	} else {
		goto L207
	}
L203:
	;
	if v386 == v866 {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	F_dbUpdateObjectWithVolatileItemsTracking(m, v869, v870)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L23
	} else {
		goto L205
	}
L205:
	;
	goto L202
L206:
	;
	if v216 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L207:
	;
	v875 = int32(_a20)
	v877 = *(*int64)(unsafe.Add(mBase, _consts[858]))
	*(*int64)(unsafe.Add(mBase, _consts[858])) = v877 + base.I64_extend_i32_u(v856)
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v883)+4))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v885)+28))
	F_notifyKeyspaceEvent(m, int32(64), int32(_a2359), v884, v886)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L23
	} else {
		goto L208
	}
L208:
	;
	if v857 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v952 = int32(0)
	goto L206
L210:
	;
	v894 = int32(0)
	goto L211
L211:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v915 = F_propagateFieldsDeletion(m, v908, v909, v856-v894, v857+v894<<(uint(int32(2))%32), v914)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L23
	} else {
		goto L213
	}
L212:
	;
	F_valkey_free(m, v857)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L23
	} else {
		goto L215
	}
L213:
	;
	v917 = v915 + v894
	if v917 < v856 {
		v894 = v917
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	goto L209
L216:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v959)+4))
	F_signalModifiedKey(m, l0, v958, v960)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L23
	} else {
		goto L219
	}
L217:
	;
	F_replaceClientCommandVector(m, l0, v852, v672)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L23
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v963 = int32(_a20)
	v965 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	*(*int64)(unsafe.Add(mBase, _consts[180])) = v965 + base.I64_extend_i32_s(v853)
	v1007 = v952
	goto L157
L220:
	;
	if v672 == int32(0) {
		v1007 = v983
		goto L157
	} else {
		goto L223
	}
L221:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v985)+4))
	F_decrRefCount(m, v986)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L23
	} else {
		goto L222
	}
L222:
	;
	goto L220
L223:
	;
	F_valkey_free(m, v672)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L23
	} else {
		goto L224
	}
L224:
	;
	v1007 = v983
	goto L157
L225:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	if v1017 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	if v218 != 0 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+4))
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+28))
	F_notifyKeyspaceEvent(m, int32(64), int32(_a2366), v1023, v1025)
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L23
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v1037 = F_hashTypeLength(m, v1036)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L23
	} else {
		goto L233
	}
L230:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+4))
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+28))
	F_notifyKeyspaceEvent(m, int32(64), int32(_a2359), v1031, v1033)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L23
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	if v1007 != 0 {
		goto L237
	} else {
		goto L238
	}
L233:
	;
	if v1037 != 0 {
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+4))
	v1042 = F_dbDelete(m, v1039, v1041)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L23
	} else {
		goto L235
	}
L235:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1046)+4))
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a308), v1047, v1049)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L23
	} else {
		goto L236
	}
L236:
	;
	goto L232
L237:
	;
	F__serverAssert(m, int32(_a2367), int32(_a2349), int32(1589))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L23
	} else {
		goto L240
	}
L238:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, _consts[331]))
	F_addReply(m, l0, v1053)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L23
	} else {
		goto L239
	}
L239:
	;
	goto L1
L240:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	goto L1
}
func F_hsetnxCommand(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
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
	var v94 int32
	_ = v94
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
	var v102 int64
	_ = v102
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = F_lookupKeyWrite(m, v9, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = F_checkType(m, l0, v12, int32(4))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v15 != 0 {
				m.G0 = v7 + int32(16)
				return
			} else {
				if v12 != 0 {
					v26 = v12
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
					v29 = F_objectGetVal(m, v28)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(9223372036854775807)
					v39 = F_hashTypeGetValue(m, v26, v29, v7+int32(12), v7+int32(8), v7, int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						if v39 != 0 {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							F_hashTypeTryConversion(m, v26, v45, int32(2), int32(3))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v50 = F_hashTypeHasVolatileFields(m, v26)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									v52 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v52)
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
									v56 = F_objectGetVal(m, v55)
									mBase = m.M
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
									v59 = F_objectGetVal(m, v58)
									mBase = m.M
									v62 = F_hashTypeSet(m, v26, v56, v59, int64(-1), v52, v7)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										v64 = F_hashTypeHasVolatileFields(m, v26)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											if v50 == v64 {
												v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
												F_signalModifiedKey(m, l0, v70, v72)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return
												} else {
													v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
													if v75 != int32(1) {
														v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
														v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
														F_notifyKeyspaceEvent(m, int32(64), int32(_a2358), v95, v97)
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return
														} else {
															v100 = int32(_a20)
															v102 = *(*int64)(unsafe.Add(mBase, _consts[180]))
															*(*int64)(unsafe.Add(mBase, _consts[180])) = v102 + int64(1)
															if v50 == int32(0) {
																v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																F_addReply(m, l0, v114)
																mBase = m.M
																v116 = m.ExcPending
																if v116 != 0 {
																	return
																} else {
																	m.G0 = v7 + int32(16)
																	return
																}
															} else {
																v110 = *(*int32)(unsafe.Add(mBase, _consts[857]))
																F_rewriteClientCommandArgument(m, l0, int32(0), v110)
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return
																} else {
																	v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																	F_addReply(m, l0, v114)
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		m.G0 = v7 + int32(16)
																		return
																	}
																}
															}
														}
													} else {
														v78 = int32(_a20)
														v80 = *(*int64)(unsafe.Add(mBase, _consts[858]))
														*(*int64)(unsafe.Add(mBase, _consts[858])) = v80 + int64(1)
														v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
														v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+28))
														F_notifyKeyspaceEvent(m, int32(64), int32(_a2359), v87, v89)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return
														} else {
															v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
															v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
															F_notifyKeyspaceEvent(m, int32(64), int32(_a2358), v95, v97)
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return
															} else {
																v100 = int32(_a20)
																v102 = *(*int64)(unsafe.Add(mBase, _consts[180]))
																*(*int64)(unsafe.Add(mBase, _consts[180])) = v102 + int64(1)
																if v50 == int32(0) {
																	v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																	F_addReply(m, l0, v114)
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		m.G0 = v7 + int32(16)
																		return
																	}
																} else {
																	v110 = *(*int32)(unsafe.Add(mBase, _consts[857]))
																	F_rewriteClientCommandArgument(m, l0, int32(0), v110)
																	mBase = m.M
																	v112 = m.ExcPending
																	if v112 != 0 {
																		return
																	} else {
																		v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																		F_addReply(m, l0, v114)
																		mBase = m.M
																		v116 = m.ExcPending
																		if v116 != 0 {
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
												v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
												F_dbUpdateObjectWithVolatileItemsTracking(m, v67, v26)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return
												} else {
													v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
													v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
													F_signalModifiedKey(m, l0, v70, v72)
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return
													} else {
														v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
														if v75 != int32(1) {
															v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
															v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
															F_notifyKeyspaceEvent(m, int32(64), int32(_a2358), v95, v97)
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return
															} else {
																v100 = int32(_a20)
																v102 = *(*int64)(unsafe.Add(mBase, _consts[180]))
																*(*int64)(unsafe.Add(mBase, _consts[180])) = v102 + int64(1)
																if v50 == int32(0) {
																	v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																	F_addReply(m, l0, v114)
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		m.G0 = v7 + int32(16)
																		return
																	}
																} else {
																	v110 = *(*int32)(unsafe.Add(mBase, _consts[857]))
																	F_rewriteClientCommandArgument(m, l0, int32(0), v110)
																	mBase = m.M
																	v112 = m.ExcPending
																	if v112 != 0 {
																		return
																	} else {
																		v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																		F_addReply(m, l0, v114)
																		mBase = m.M
																		v116 = m.ExcPending
																		if v116 != 0 {
																			return
																		} else {
																			m.G0 = v7 + int32(16)
																			return
																		}
																	}
																}
															}
														} else {
															v78 = int32(_a20)
															v80 = *(*int64)(unsafe.Add(mBase, _consts[858]))
															*(*int64)(unsafe.Add(mBase, _consts[858])) = v80 + int64(1)
															v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
															v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+28))
															F_notifyKeyspaceEvent(m, int32(64), int32(_a2359), v87, v89)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return
															} else {
																v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
																v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
																F_notifyKeyspaceEvent(m, int32(64), int32(_a2358), v95, v97)
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return
																} else {
																	v100 = int32(_a20)
																	v102 = *(*int64)(unsafe.Add(mBase, _consts[180]))
																	*(*int64)(unsafe.Add(mBase, _consts[180])) = v102 + int64(1)
																	if v50 == int32(0) {
																		v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																		F_addReply(m, l0, v114)
																		mBase = m.M
																		v116 = m.ExcPending
																		if v116 != 0 {
																			return
																		} else {
																			m.G0 = v7 + int32(16)
																			return
																		}
																	} else {
																		v110 = *(*int32)(unsafe.Add(mBase, _consts[857]))
																		F_rewriteClientCommandArgument(m, l0, int32(0), v110)
																		mBase = m.M
																		v112 = m.ExcPending
																		if v112 != 0 {
																			return
																		} else {
																			v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																			F_addReply(m, l0, v114)
																			mBase = m.M
																			v116 = m.ExcPending
																			if v116 != 0 {
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
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, _consts[71]))
							F_addReply(m, l0, v42)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				} else {
					v17 = F_createHashObject(m)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v17
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						F_dbAdd(m, v20, v11, v7)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
							if v23 == int32(0) {
								m.G0 = v7 + int32(16)
								return
							} else {
								v26 = v23
								v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
								v29 = F_objectGetVal(m, v28)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(-1)
								*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(9223372036854775807)
								v39 = F_hashTypeGetValue(m, v26, v29, v7+int32(12), v7+int32(8), v7, int32(0))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									if v39 != 0 {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										F_hashTypeTryConversion(m, v26, v45, int32(2), int32(3))
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											v50 = F_hashTypeHasVolatileFields(m, v26)
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return
											} else {
												v52 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v52)
												v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
												v56 = F_objectGetVal(m, v55)
												mBase = m.M
												v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
												v59 = F_objectGetVal(m, v58)
												mBase = m.M
												v62 = F_hashTypeSet(m, v26, v56, v59, int64(-1), v52, v7)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													v64 = F_hashTypeHasVolatileFields(m, v26)
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return
													} else {
														if v50 == v64 {
															v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
															F_signalModifiedKey(m, l0, v70, v72)
															mBase = m.M
															v74 = m.ExcPending
															if v74 != 0 {
																return
															} else {
																v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
																if v75 != int32(1) {
																	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
																	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
																	F_notifyKeyspaceEvent(m, int32(64), int32(_a2358), v95, v97)
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return
																	} else {
																		v100 = int32(_a20)
																		v102 = *(*int64)(unsafe.Add(mBase, _consts[180]))
																		*(*int64)(unsafe.Add(mBase, _consts[180])) = v102 + int64(1)
																		if v50 == int32(0) {
																			v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																			F_addReply(m, l0, v114)
																			mBase = m.M
																			v116 = m.ExcPending
																			if v116 != 0 {
																				return
																			} else {
																				m.G0 = v7 + int32(16)
																				return
																			}
																		} else {
																			v110 = *(*int32)(unsafe.Add(mBase, _consts[857]))
																			F_rewriteClientCommandArgument(m, l0, int32(0), v110)
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return
																			} else {
																				v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																				F_addReply(m, l0, v114)
																				mBase = m.M
																				v116 = m.ExcPending
																				if v116 != 0 {
																					return
																				} else {
																					m.G0 = v7 + int32(16)
																					return
																				}
																			}
																		}
																	}
																} else {
																	v78 = int32(_a20)
																	v80 = *(*int64)(unsafe.Add(mBase, _consts[858]))
																	*(*int64)(unsafe.Add(mBase, _consts[858])) = v80 + int64(1)
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
																	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+28))
																	F_notifyKeyspaceEvent(m, int32(64), int32(_a2359), v87, v89)
																	mBase = m.M
																	v91 = m.ExcPending
																	if v91 != 0 {
																		return
																	} else {
																		v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
																		v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																		v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
																		F_notifyKeyspaceEvent(m, int32(64), int32(_a2358), v95, v97)
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return
																		} else {
																			v100 = int32(_a20)
																			v102 = *(*int64)(unsafe.Add(mBase, _consts[180]))
																			*(*int64)(unsafe.Add(mBase, _consts[180])) = v102 + int64(1)
																			if v50 == int32(0) {
																				v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																				F_addReply(m, l0, v114)
																				mBase = m.M
																				v116 = m.ExcPending
																				if v116 != 0 {
																					return
																				} else {
																					m.G0 = v7 + int32(16)
																					return
																				}
																			} else {
																				v110 = *(*int32)(unsafe.Add(mBase, _consts[857]))
																				F_rewriteClientCommandArgument(m, l0, int32(0), v110)
																				mBase = m.M
																				v112 = m.ExcPending
																				if v112 != 0 {
																					return
																				} else {
																					v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																					F_addReply(m, l0, v114)
																					mBase = m.M
																					v116 = m.ExcPending
																					if v116 != 0 {
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
															v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															F_dbUpdateObjectWithVolatileItemsTracking(m, v67, v26)
															mBase = m.M
															v69 = m.ExcPending
															if v69 != 0 {
																return
															} else {
																v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
																F_signalModifiedKey(m, l0, v70, v72)
																mBase = m.M
																v74 = m.ExcPending
																if v74 != 0 {
																	return
																} else {
																	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
																	if v75 != int32(1) {
																		v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
																		v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																		v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
																		F_notifyKeyspaceEvent(m, int32(64), int32(_a2358), v95, v97)
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return
																		} else {
																			v100 = int32(_a20)
																			v102 = *(*int64)(unsafe.Add(mBase, _consts[180]))
																			*(*int64)(unsafe.Add(mBase, _consts[180])) = v102 + int64(1)
																			if v50 == int32(0) {
																				v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																				F_addReply(m, l0, v114)
																				mBase = m.M
																				v116 = m.ExcPending
																				if v116 != 0 {
																					return
																				} else {
																					m.G0 = v7 + int32(16)
																					return
																				}
																			} else {
																				v110 = *(*int32)(unsafe.Add(mBase, _consts[857]))
																				F_rewriteClientCommandArgument(m, l0, int32(0), v110)
																				mBase = m.M
																				v112 = m.ExcPending
																				if v112 != 0 {
																					return
																				} else {
																					v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																					F_addReply(m, l0, v114)
																					mBase = m.M
																					v116 = m.ExcPending
																					if v116 != 0 {
																						return
																					} else {
																						m.G0 = v7 + int32(16)
																						return
																					}
																				}
																			}
																		}
																	} else {
																		v78 = int32(_a20)
																		v80 = *(*int64)(unsafe.Add(mBase, _consts[858]))
																		*(*int64)(unsafe.Add(mBase, _consts[858])) = v80 + int64(1)
																		v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
																		v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																		v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+28))
																		F_notifyKeyspaceEvent(m, int32(64), int32(_a2359), v87, v89)
																		mBase = m.M
																		v91 = m.ExcPending
																		if v91 != 0 {
																			return
																		} else {
																			v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
																			v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																			v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
																			F_notifyKeyspaceEvent(m, int32(64), int32(_a2358), v95, v97)
																			mBase = m.M
																			v99 = m.ExcPending
																			if v99 != 0 {
																				return
																			} else {
																				v100 = int32(_a20)
																				v102 = *(*int64)(unsafe.Add(mBase, _consts[180]))
																				*(*int64)(unsafe.Add(mBase, _consts[180])) = v102 + int64(1)
																				if v50 == int32(0) {
																					v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																					F_addReply(m, l0, v114)
																					mBase = m.M
																					v116 = m.ExcPending
																					if v116 != 0 {
																						return
																					} else {
																						m.G0 = v7 + int32(16)
																						return
																					}
																				} else {
																					v110 = *(*int32)(unsafe.Add(mBase, _consts[857]))
																					F_rewriteClientCommandArgument(m, l0, int32(0), v110)
																					mBase = m.M
																					v112 = m.ExcPending
																					if v112 != 0 {
																						return
																					} else {
																						v114 = *(*int32)(unsafe.Add(mBase, _consts[331]))
																						F_addReply(m, l0, v114)
																						mBase = m.M
																						v116 = m.ExcPending
																						if v116 != 0 {
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
									} else {
										v42 = *(*int32)(unsafe.Add(mBase, _consts[71]))
										F_addReply(m, l0, v42)
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
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
func F_httlCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	F_httlGenericCommand(m, l0, v3, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_humanNodename(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2316))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(-1)))))
	switch v13 & int32(7) {
	case 0:
		v30 = int32(base.Ui32(v13) >> (uint(int32(3)) % 32))
		if v30 != 0 {
			v71 = v10
			m.G0 = v8 + int32(16)
			return v71
		} else {
			v32 = int32(0)
			v34 = *(*int32)(unsafe.Add(mBase, _consts[118]))
			v38 = base.I32_rem_s(v34+int32(1), int32(8))
			*(*int32)(unsafe.Add(mBase, _consts[118])) = v38
			v41 = l0 + int32(2256)
			v42 = int32(58)
			v43 = F___strchrnul(m, v41, v42)
			mBase = m.M
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
			if v45 == v42 {
				v49 = v43
			} else {
				v49 = v32
			}
			v53 = *(*int32)(unsafe.Add(mBase, _consts[107]))
			if v53 != 0 {
				v54 = int32(2328)
			} else {
				v54 = int32(2324)
			}
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l0+v54)))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v56
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v41
			v62 = v34<<(uint(int32(7))%32) + int32(_a242)
			if v49 != 0 {
				v66 = int32(_a243)
			} else {
				v66 = int32(_a244)
			}
			v67 = F_snprintf(m, v62, int32(128), v66, v8)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v71 = v62
				m.G0 = v8 + int32(16)
				return v71
			}
		}
	case 1:
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(-3)))))
		v30 = v20
		if v30 != 0 {
			v71 = v10
			m.G0 = v8 + int32(16)
			return v71
		} else {
			v32 = int32(0)
			v34 = *(*int32)(unsafe.Add(mBase, _consts[118]))
			v38 = base.I32_rem_s(v34+int32(1), int32(8))
			*(*int32)(unsafe.Add(mBase, _consts[118])) = v38
			v41 = l0 + int32(2256)
			v42 = int32(58)
			v43 = F___strchrnul(m, v41, v42)
			mBase = m.M
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
			if v45 == v42 {
				v49 = v43
			} else {
				v49 = v32
			}
			v53 = *(*int32)(unsafe.Add(mBase, _consts[107]))
			if v53 != 0 {
				v54 = int32(2328)
			} else {
				v54 = int32(2324)
			}
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l0+v54)))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v56
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v41
			v62 = v34<<(uint(int32(7))%32) + int32(_a242)
			if v49 != 0 {
				v66 = int32(_a243)
			} else {
				v66 = int32(_a244)
			}
			v67 = F_snprintf(m, v62, int32(128), v66, v8)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v71 = v62
				m.G0 = v8 + int32(16)
				return v71
			}
		}
	case 2:
		v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(-5)))))
		v30 = v23
		if v30 != 0 {
			v71 = v10
			m.G0 = v8 + int32(16)
			return v71
		} else {
			v32 = int32(0)
			v34 = *(*int32)(unsafe.Add(mBase, _consts[118]))
			v38 = base.I32_rem_s(v34+int32(1), int32(8))
			*(*int32)(unsafe.Add(mBase, _consts[118])) = v38
			v41 = l0 + int32(2256)
			v42 = int32(58)
			v43 = F___strchrnul(m, v41, v42)
			mBase = m.M
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
			if v45 == v42 {
				v49 = v43
			} else {
				v49 = v32
			}
			v53 = *(*int32)(unsafe.Add(mBase, _consts[107]))
			if v53 != 0 {
				v54 = int32(2328)
			} else {
				v54 = int32(2324)
			}
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l0+v54)))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v56
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v41
			v62 = v34<<(uint(int32(7))%32) + int32(_a242)
			if v49 != 0 {
				v66 = int32(_a243)
			} else {
				v66 = int32(_a244)
			}
			v67 = F_snprintf(m, v62, int32(128), v66, v8)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v71 = v62
				m.G0 = v8 + int32(16)
				return v71
			}
		}
	case 3:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-9))))
		v30 = v26
		if v30 != 0 {
			v71 = v10
			m.G0 = v8 + int32(16)
			return v71
		} else {
			v32 = int32(0)
			v34 = *(*int32)(unsafe.Add(mBase, _consts[118]))
			v38 = base.I32_rem_s(v34+int32(1), int32(8))
			*(*int32)(unsafe.Add(mBase, _consts[118])) = v38
			v41 = l0 + int32(2256)
			v42 = int32(58)
			v43 = F___strchrnul(m, v41, v42)
			mBase = m.M
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
			if v45 == v42 {
				v49 = v43
			} else {
				v49 = v32
			}
			v53 = *(*int32)(unsafe.Add(mBase, _consts[107]))
			if v53 != 0 {
				v54 = int32(2328)
			} else {
				v54 = int32(2324)
			}
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l0+v54)))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v56
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v41
			v62 = v34<<(uint(int32(7))%32) + int32(_a242)
			if v49 != 0 {
				v66 = int32(_a243)
			} else {
				v66 = int32(_a244)
			}
			v67 = F_snprintf(m, v62, int32(128), v66, v8)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v71 = v62
				m.G0 = v8 + int32(16)
				return v71
			}
		}
	case 4:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-17))))
		v30 = v29
		if v30 != 0 {
			v71 = v10
			m.G0 = v8 + int32(16)
			return v71
		} else {
			v32 = int32(0)
			v34 = *(*int32)(unsafe.Add(mBase, _consts[118]))
			v38 = base.I32_rem_s(v34+int32(1), int32(8))
			*(*int32)(unsafe.Add(mBase, _consts[118])) = v38
			v41 = l0 + int32(2256)
			v42 = int32(58)
			v43 = F___strchrnul(m, v41, v42)
			mBase = m.M
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
			if v45 == v42 {
				v49 = v43
			} else {
				v49 = v32
			}
			v53 = *(*int32)(unsafe.Add(mBase, _consts[107]))
			if v53 != 0 {
				v54 = int32(2328)
			} else {
				v54 = int32(2324)
			}
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l0+v54)))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v56
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v41
			v62 = v34<<(uint(int32(7))%32) + int32(_a242)
			if v49 != 0 {
				v66 = int32(_a243)
			} else {
				v66 = int32(_a244)
			}
			v67 = F_snprintf(m, v62, int32(128), v66, v8)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v71 = v62
				m.G0 = v8 + int32(16)
				return v71
			}
		}
	default:
		v32 = int32(0)
		v34 = *(*int32)(unsafe.Add(mBase, _consts[118]))
		v38 = base.I32_rem_s(v34+int32(1), int32(8))
		*(*int32)(unsafe.Add(mBase, _consts[118])) = v38
		v41 = l0 + int32(2256)
		v42 = int32(58)
		v43 = F___strchrnul(m, v41, v42)
		mBase = m.M
		v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
		if v45 == v42 {
			v49 = v43
		} else {
			v49 = v32
		}
		v53 = *(*int32)(unsafe.Add(mBase, _consts[107]))
		if v53 != 0 {
			v54 = int32(2328)
		} else {
			v54 = int32(2324)
		}
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l0+v54)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v56
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v41
		v62 = v34<<(uint(int32(7))%32) + int32(_a242)
		if v49 != 0 {
			v66 = int32(_a243)
		} else {
			v66 = int32(_a244)
		}
		v67 = F_snprintf(m, v62, int32(128), v66, v8)
		mBase = m.M
		v70 = m.ExcPending
		if v70 != 0 {
			return int32(0)
		} else {
			v71 = v62
			m.G0 = v8 + int32(16)
			return v71
		}
	}
}
func F_hvalsCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_genericHgetallCommand(m, l0, int32(2))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
