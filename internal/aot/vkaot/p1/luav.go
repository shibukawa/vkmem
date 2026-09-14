package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaV_concat(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 float64
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
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
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 float64
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v16 = l1
	v17 = l2
	goto L1
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = v25 + v17<<(uint(int32(4))%32)
	v30 = v28 + int32(-16)
	v32 = v28 + int32(-8)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if base.Ui32(int32(1)) < base.Ui32(v33+int32(-3)) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	m.G0 = v13 + int32(80)
	return
L3:
	;
	v383 = v16 - v379
	if int32(1) < v383 {
		v16 = v383
		v17 = v17 - v379
		goto L1
	} else {
		goto L87
	}
L4:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	if v124 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L5:
	;
	v47 = *(*float64)(unsafe.Add(mBase, uint32(v28)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+32)) = v47
	v51 = m.G3
	v56 = F___small_sprintf(m, v13+int32(48), v51+int32(_a2688), v13+int32(32))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L12
	}
L6:
	;
	v41 = int32(1)
	v43 = F_call_binTM(m, l0, v30, v28, v30, int32(15))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	switch v38 + int32(-3) {
	case 0:
		goto L5
	case 1:
		goto L4
	default:
		goto L6
	}
L8:
	;
	return
L9:
	;
	if v43 != 0 {
		v379 = v41
		goto L3
	} else {
		goto L10
	}
L10:
	;
	F_luaG_concaterror(m, l0, v30, v28)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v379 = v41
	goto L3
L12:
	;
	v59 = v13 + int32(48)
	if v59&int32(3) == int32(0) {
		v83 = v59
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v117 = F_luaS_newlstr(m, l0, v59, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L8
	} else {
		goto L29
	}
L14:
	;
	v116 = v108 - v59
	goto L13
L15:
	;
	v87 = v83
	goto L23
L16:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v69 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v72 = v59
	goto L19
L18:
	;
	v116 = v59 - v59
	goto L13
L19:
	;
	v76 = v72 + int32(1)
	if v76&int32(3) == int32(0) {
		v83 = v76
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v81 != 0 {
		v72 = v76
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v108 = v76
	goto L14
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v96 = int32(-2139062144)
	if (int32(16843008)-v93|v93)&v96 == v96 {
		v87 = v87 + int32(4)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v102 = v87
	goto L26
L25:
	;
	goto L24
L26:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v106 != 0 {
		v102 = v102 + int32(1)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v108 = v102
	goto L14
L28:
	;
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v117
	goto L4
L30:
	;
	v295 = int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v296 != int32(3) {
		v379 = v295
		goto L3
	} else {
		goto L68
	}
L31:
	;
	v128 = v28 + int32(16)
	v129 = int32(1)
	if v16 <= v129 {
		v249 = v124
		v250 = v129
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v256 = F_luaZ_openspace(m, l0, v253+int32(52), v249)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L8
	} else {
		goto L60
	}
L33:
	;
	v139 = v124
	v140 = v129
	goto L34
L34:
	;
	v144 = v128 - v140<<(uint(int32(4))%32)
	v146 = v144 + int32(-16)
	v148 = v144 + int32(-8)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	switch v149 + int32(-3) {
	case 0:
		goto L37
	case 1:
		goto L36
	default:
		v249 = v139
		v250 = v140
		goto L32
	}
L35:
	;
	v249 = v238
	v250 = v16
	goto L32
L36:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	if base.Ui32(v228) < base.Ui32(int32(-3)-v139) {
		goto L56
	} else {
		goto L57
	}
L37:
	;
	v152 = *(*float64)(unsafe.Add(mBase, uint32(v146)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = v152
	v156 = m.G3
	v161 = F___small_sprintf(m, v13+int32(48), v156+int32(_a2688), v13+int32(16))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v164 = v13 + int32(48)
	if v164&int32(3) == int32(0) {
		v188 = v164
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v222 = F_luaS_newlstr(m, l0, v164, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L8
	} else {
		goto L55
	}
L40:
	;
	v221 = v213 - v164
	goto L39
L41:
	;
	v192 = v188
	goto L49
L42:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v174 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v177 = v164
	goto L45
L44:
	;
	v221 = v164 - v164
	goto L39
L45:
	;
	v181 = v177 + int32(1)
	if v181&int32(3) == int32(0) {
		v188 = v181
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v186 != 0 {
		v177 = v181
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v213 = v181
	goto L40
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v201 = int32(-2139062144)
	if (int32(16843008)-v198|v198)&v201 == v201 {
		v192 = v192 + int32(4)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v207 = v192
	goto L52
L51:
	;
	goto L50
L52:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v211 != 0 {
		v207 = v207 + int32(1)
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v213 = v207
	goto L40
L54:
	;
	goto L53
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = int32(4)
	goto L36
L56:
	;
	v238 = v228 + v139
	v240 = v140 + int32(1)
	if v240 != v16 {
		v139 = v238
		v140 = v240
		goto L34
	} else {
		goto L59
	}
L57:
	;
	v232 = m.G3
	F_luaG_runerror(m, l0, v232+int32(_a2689), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	goto L35
L60:
	;
	v262 = int32(0)
	v265 = v250
	goto L61
L61:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v128-v265<<(uint(int32(4))%32))))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	if v275 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v285 = F_luaS_newlstr(m, l0, v256, v280)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L8
	} else {
		goto L67
	}
L63:
	;
	v280 = v275 + v262
	if int32(1) < v265 {
		v262 = v280
		v265 = v265 + int32(-1)
		goto L61
	} else {
		goto L66
	}
L64:
	;
	goto L63
L65:
	;
	v278 = F__emscripten_memcpy_bulkmem(m, v256+v262, v272+int32(16), v275)
	mBase = m.M
	goto L64
L66:
	;
	goto L62
L67:
	;
	v287 = int32(4)
	v289 = v128 - v250<<(uint(v287)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v289)+8)) = v287
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v285
	v379 = v250 + int32(-1)
	goto L3
L68:
	;
	v299 = *(*float64)(unsafe.Add(mBase, uint32(v30)))
	*(*float64)(unsafe.Add(mBase, uint32(v13))) = v299
	v303 = m.G3
	v306 = F___small_sprintf(m, v13+int32(48), v303+int32(_a2688), v13)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	v309 = v13 + int32(48)
	if v309&int32(3) == int32(0) {
		v333 = v309
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v367 = F_luaS_newlstr(m, l0, v309, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L8
	} else {
		goto L86
	}
L71:
	;
	v366 = v358 - v309
	goto L70
L72:
	;
	v337 = v333
	goto L80
L73:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if v319 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v322 = v309
	goto L76
L75:
	;
	v366 = v309 - v309
	goto L70
L76:
	;
	v326 = v322 + int32(1)
	if v326&int32(3) == int32(0) {
		v333 = v326
		goto L72
	} else {
		goto L78
	}
L78:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	if v331 != 0 {
		v322 = v326
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v358 = v326
	goto L71
L80:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	v346 = int32(-2139062144)
	if (int32(16843008)-v343|v343)&v346 == v346 {
		v337 = v337 + int32(4)
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v352 = v337
	goto L83
L82:
	;
	goto L81
L83:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	if v356 != 0 {
		v352 = v352 + int32(1)
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v358 = v352
	goto L71
L85:
	;
	goto L84
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(4)
	v379 = v295
	goto L3
L87:
	;
	goto L2
}
func F_luaV_equalval(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
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
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
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
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int64
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	v8 = int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	switch v9 {
	case 0:
		v219 = v8
		return v219
	case 1:
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		return base.B2i32(v14 == v15)
	case 2:
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		return base.B2i32(v18 == v19)
	case 3:
		v10 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
		v11 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
		return base.F64_eq(v10, v11)
	default:
		v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v225 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		return base.B2i32(v224 == v225)
	case 5:
		v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v89 == v90 {
			v219 = v8
			return v219
		} else {
			v92 = int32(0)
			v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
			if v93 == v92 {
				v219 = v92
				return v219
			} else {
				v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+6)))
				if v96&int32(16) != 0 {
					v219 = v92
					return v219
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
					v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+204))
					v103 = F_luaH_getstr(m, v93, v102)
					mBase = m.M
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
					if v104 != 0 {
						v111 = v103
					} else {
						v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+6)))
						v108 = v105 | int32(16)
						*(*uint8)(unsafe.Add(mBase, uint32(v93)+6)) = uint8(v108)
						v111 = int32(0)
					}
					if v111 == int32(0) {
						v219 = v92
						return v219
					} else {
						if v93 == v99 {
							v160 = v111
							v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v164 = *(*int64)(unsafe.Add(mBase, uint32(v160)))
							*(*int64)(unsafe.Add(mBase, uint32(v163))) = v164
							v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v166
							v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v169 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
							*(*int64)(unsafe.Add(mBase, uint32(v168)+16)) = v169
							v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v168)+24)) = v171
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v174 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							*(*int64)(unsafe.Add(mBase, uint32(v173)+32)) = v174
							v176 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v173)+40)) = v176
							v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if int32(48) < v178-v179 {
								v189 = v179
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v189 + int32(48)
								F_luaD_call(m, l0, v189, int32(1))
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return int32(0)
								} else {
									v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v198 = v196 + int32(-16)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v198
									v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v202 = v200 + (v163 - v162)
									v203 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
									*(*int64)(unsafe.Add(mBase, uint32(v202))) = v203
									v207 = *(*int32)(unsafe.Add(mBase, uint32(v196+int32(-8))))
									*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = v207
									v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
									switch v211 {
									case 0:
										v219 = int32(0)
										return v219
									case 1:
										v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
										return base.B2i32(v212 != int32(0))
									default:
										v219 = int32(1)
										return v219
									}
								}
							} else {
								F_luaD_growstack(m, l0, int32(3))
								mBase = m.M
								v187 = m.ExcPending
								if v187 != 0 {
									return int32(0)
								} else {
									v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v189 = v188
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v189 + int32(48)
									F_luaD_call(m, l0, v189, int32(1))
									mBase = m.M
									v195 = m.ExcPending
									if v195 != 0 {
										return int32(0)
									} else {
										v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v198 = v196 + int32(-16)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v198
										v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v202 = v200 + (v163 - v162)
										v203 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
										*(*int64)(unsafe.Add(mBase, uint32(v202))) = v203
										v207 = *(*int32)(unsafe.Add(mBase, uint32(v196+int32(-8))))
										*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = v207
										v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
										switch v211 {
										case 0:
											v219 = int32(0)
											return v219
										case 1:
											v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
											return base.B2i32(v212 != int32(0))
										default:
											v219 = int32(1)
											return v219
										}
									}
								}
							}
						} else {
							if v99 == int32(0) {
								v219 = v92
								return v219
							} else {
								v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+6)))
								if v117&int32(16) != 0 {
									v219 = v92
									return v219
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+204))
									v123 = F_luaH_getstr(m, v99, v122)
									mBase = m.M
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
									if v124 != 0 {
										v131 = v123
									} else {
										v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+6)))
										v128 = v125 | int32(16)
										*(*uint8)(unsafe.Add(mBase, uint32(v99)+6)) = uint8(v128)
										v131 = int32(0)
									}
									if v131 == int32(0) {
										v219 = v92
										return v219
									} else {
										v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
										v137 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
										if v136 == v137 {
											switch v136 {
											case 0:
												v153 = int32(1)
												v155 = v153
											case 1:
												v144 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
												v145 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
												v155 = base.B2i32(v144 == v145)
											case 2:
												v147 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
												v155 = base.B2i32(v147 == v148)
											case 3:
												v141 = *(*float64)(unsafe.Add(mBase, uint32(v111)))
												v142 = *(*float64)(unsafe.Add(mBase, uint32(v131)))
												v155 = base.F64_eq(v141, v142)
											default:
												v150 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
												v151 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
												v153 = base.B2i32(v150 == v151)
												v155 = v153
											}
										} else {
											v155 = int32(0)
										}
										if v155 == int32(0) {
											v219 = v92
											return v219
										} else {
											v160 = v111
											v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v164 = *(*int64)(unsafe.Add(mBase, uint32(v160)))
											*(*int64)(unsafe.Add(mBase, uint32(v163))) = v164
											v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v166
											v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v169 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
											*(*int64)(unsafe.Add(mBase, uint32(v168)+16)) = v169
											v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v168)+24)) = v171
											v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v174 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
											*(*int64)(unsafe.Add(mBase, uint32(v173)+32)) = v174
											v176 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v173)+40)) = v176
											v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if int32(48) < v178-v179 {
												v189 = v179
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v189 + int32(48)
												F_luaD_call(m, l0, v189, int32(1))
												mBase = m.M
												v195 = m.ExcPending
												if v195 != 0 {
													return int32(0)
												} else {
													v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v198 = v196 + int32(-16)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v198
													v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													v202 = v200 + (v163 - v162)
													v203 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
													*(*int64)(unsafe.Add(mBase, uint32(v202))) = v203
													v207 = *(*int32)(unsafe.Add(mBase, uint32(v196+int32(-8))))
													*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = v207
													v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
													switch v211 {
													case 0:
														v219 = int32(0)
														return v219
													case 1:
														v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
														return base.B2i32(v212 != int32(0))
													default:
														v219 = int32(1)
														return v219
													}
												}
											} else {
												F_luaD_growstack(m, l0, int32(3))
												mBase = m.M
												v187 = m.ExcPending
												if v187 != 0 {
													return int32(0)
												} else {
													v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v189 = v188
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v189 + int32(48)
													F_luaD_call(m, l0, v189, int32(1))
													mBase = m.M
													v195 = m.ExcPending
													if v195 != 0 {
														return int32(0)
													} else {
														v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v198 = v196 + int32(-16)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v198
														v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														v202 = v200 + (v163 - v162)
														v203 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
														*(*int64)(unsafe.Add(mBase, uint32(v202))) = v203
														v207 = *(*int32)(unsafe.Add(mBase, uint32(v196+int32(-8))))
														*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = v207
														v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
														switch v211 {
														case 0:
															v219 = int32(0)
															return v219
														case 1:
															v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
															return base.B2i32(v212 != int32(0))
														default:
															v219 = int32(1)
															return v219
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
	case 7:
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v22 == v23 {
			v219 = v8
			return v219
		} else {
			v25 = int32(0)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
			if v26 == v25 {
				v219 = v25
				return v219
			} else {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+6)))
				if v29&int32(16) != 0 {
					v219 = v25
					return v219
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+204))
					v36 = F_luaH_getstr(m, v26, v35)
					mBase = m.M
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
					if v37 != 0 {
						v44 = v36
					} else {
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+6)))
						v41 = v38 | int32(16)
						*(*uint8)(unsafe.Add(mBase, uint32(v26)+6)) = uint8(v41)
						v44 = int32(0)
					}
					if v44 == int32(0) {
						v219 = v25
						return v219
					} else {
						if v26 == v32 {
							v160 = v44
							v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v164 = *(*int64)(unsafe.Add(mBase, uint32(v160)))
							*(*int64)(unsafe.Add(mBase, uint32(v163))) = v164
							v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v166
							v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v169 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
							*(*int64)(unsafe.Add(mBase, uint32(v168)+16)) = v169
							v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v168)+24)) = v171
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v174 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							*(*int64)(unsafe.Add(mBase, uint32(v173)+32)) = v174
							v176 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v173)+40)) = v176
							v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if int32(48) < v178-v179 {
								v189 = v179
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v189 + int32(48)
								F_luaD_call(m, l0, v189, int32(1))
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return int32(0)
								} else {
									v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v198 = v196 + int32(-16)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v198
									v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v202 = v200 + (v163 - v162)
									v203 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
									*(*int64)(unsafe.Add(mBase, uint32(v202))) = v203
									v207 = *(*int32)(unsafe.Add(mBase, uint32(v196+int32(-8))))
									*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = v207
									v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
									switch v211 {
									case 0:
										v219 = int32(0)
										return v219
									case 1:
										v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
										return base.B2i32(v212 != int32(0))
									default:
										v219 = int32(1)
										return v219
									}
								}
							} else {
								F_luaD_growstack(m, l0, int32(3))
								mBase = m.M
								v187 = m.ExcPending
								if v187 != 0 {
									return int32(0)
								} else {
									v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v189 = v188
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v189 + int32(48)
									F_luaD_call(m, l0, v189, int32(1))
									mBase = m.M
									v195 = m.ExcPending
									if v195 != 0 {
										return int32(0)
									} else {
										v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v198 = v196 + int32(-16)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v198
										v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v202 = v200 + (v163 - v162)
										v203 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
										*(*int64)(unsafe.Add(mBase, uint32(v202))) = v203
										v207 = *(*int32)(unsafe.Add(mBase, uint32(v196+int32(-8))))
										*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = v207
										v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
										switch v211 {
										case 0:
											v219 = int32(0)
											return v219
										case 1:
											v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
											return base.B2i32(v212 != int32(0))
										default:
											v219 = int32(1)
											return v219
										}
									}
								}
							}
						} else {
							if v32 == int32(0) {
								v219 = v25
								return v219
							} else {
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)))
								if v50&int32(16) != 0 {
									v219 = v25
									return v219
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+204))
									v56 = F_luaH_getstr(m, v32, v55)
									mBase = m.M
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
									if v57 != 0 {
										v64 = v56
									} else {
										v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)))
										v61 = v58 | int32(16)
										*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v61)
										v64 = int32(0)
									}
									if v64 == int32(0) {
										v219 = v25
										return v219
									} else {
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
										if v69 == v70 {
											switch v69 {
											case 0:
												v86 = int32(1)
												v88 = v86
											case 1:
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
												v88 = base.B2i32(v77 == v78)
											case 2:
												v80 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
												v88 = base.B2i32(v80 == v81)
											case 3:
												v74 = *(*float64)(unsafe.Add(mBase, uint32(v44)))
												v75 = *(*float64)(unsafe.Add(mBase, uint32(v64)))
												v88 = base.F64_eq(v74, v75)
											default:
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
												v86 = base.B2i32(v83 == v84)
												v88 = v86
											}
										} else {
											v88 = int32(0)
										}
										if v88 != 0 {
											v160 = v44
											v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v164 = *(*int64)(unsafe.Add(mBase, uint32(v160)))
											*(*int64)(unsafe.Add(mBase, uint32(v163))) = v164
											v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v166
											v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v169 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
											*(*int64)(unsafe.Add(mBase, uint32(v168)+16)) = v169
											v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v168)+24)) = v171
											v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v174 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
											*(*int64)(unsafe.Add(mBase, uint32(v173)+32)) = v174
											v176 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v173)+40)) = v176
											v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if int32(48) < v178-v179 {
												v189 = v179
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v189 + int32(48)
												F_luaD_call(m, l0, v189, int32(1))
												mBase = m.M
												v195 = m.ExcPending
												if v195 != 0 {
													return int32(0)
												} else {
													v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v198 = v196 + int32(-16)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v198
													v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													v202 = v200 + (v163 - v162)
													v203 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
													*(*int64)(unsafe.Add(mBase, uint32(v202))) = v203
													v207 = *(*int32)(unsafe.Add(mBase, uint32(v196+int32(-8))))
													*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = v207
													v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
													switch v211 {
													case 0:
														v219 = int32(0)
														return v219
													case 1:
														v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
														return base.B2i32(v212 != int32(0))
													default:
														v219 = int32(1)
														return v219
													}
												}
											} else {
												F_luaD_growstack(m, l0, int32(3))
												mBase = m.M
												v187 = m.ExcPending
												if v187 != 0 {
													return int32(0)
												} else {
													v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v189 = v188
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v189 + int32(48)
													F_luaD_call(m, l0, v189, int32(1))
													mBase = m.M
													v195 = m.ExcPending
													if v195 != 0 {
														return int32(0)
													} else {
														v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v198 = v196 + int32(-16)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v198
														v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														v202 = v200 + (v163 - v162)
														v203 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
														*(*int64)(unsafe.Add(mBase, uint32(v202))) = v203
														v207 = *(*int32)(unsafe.Add(mBase, uint32(v196+int32(-8))))
														*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = v207
														v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
														switch v211 {
														case 0:
															v219 = int32(0)
															return v219
														case 1:
															v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
															return base.B2i32(v212 != int32(0))
														default:
															v219 = int32(1)
															return v219
														}
													}
												}
											}
										} else {
											v219 = v25
											return v219
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
func F_luaV_tonumber(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 float64
	_ = v23
	var v28 int32
	_ = v28
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch v9 + int32(-3) {
	case 0:
		v28 = l0
		m.G0 = v7 + int32(16)
		return v28
	case 1:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = F_luaO_str2d(m, v12+int32(16), v7+int32(8))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v17 == int32(0) {
				v28 = int32(0)
			} else {
				v23 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(3)
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = v23
				v28 = l1
			}
			m.G0 = v7 + int32(16)
			return v28
		}
	default:
		v28 = int32(0)
		m.G0 = v7 + int32(16)
		return v28
	}
}
func F_luaV_tostring(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 float64
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v10 != int32(3) {
		v89 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(48)
	return v89
L2:
	;
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	*(*float64)(unsafe.Add(mBase, uint32(v7))) = v13
	v17 = m.G3
	v20 = F___small_sprintf(m, v7+int32(16), v17+int32(_a2688), v7)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v25 = v7 + int32(16)
	if v25&int32(3) == int32(0) {
		v49 = v25
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v83 = F_luaS_newlstr(m, l0, v25, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L3
	} else {
		goto L21
	}
L6:
	;
	v82 = v74 - v25
	goto L5
L7:
	;
	v53 = v49
	goto L15
L8:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v38 = v25
	goto L11
L10:
	;
	v82 = v25 - v25
	goto L5
L11:
	;
	v42 = v38 + int32(1)
	if v42&int32(3) == int32(0) {
		v49 = v42
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v47 != 0 {
		v38 = v42
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v74 = v42
	goto L6
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v62 = int32(-2139062144)
	if (int32(16843008)-v59|v59)&v62 == v62 {
		v53 = v53 + int32(4)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v68 = v53
	goto L18
L17:
	;
	goto L16
L18:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v72 != 0 {
		v68 = v68 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v74 = v68
	goto L6
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v83
	v89 = int32(1)
	goto L1
}
