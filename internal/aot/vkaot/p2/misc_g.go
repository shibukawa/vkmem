package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___get_locale(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
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
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int64
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v6 != 0 {
		v164 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v168 = int32(0)
	goto L56
L2:
	;
	v7 = int32(_a2182)
	v13 = F___strchrnul(m, v7, int32(61))
	mBase = m.M
	if v13 != v7 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v61 = l0*int32(12) + int32(_a2183)
	v67 = F___strchrnul(m, v61, int32(61))
	mBase = m.M
	if v67 != v61 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	if v54 == int32(0) {
		goto L3
	} else {
		goto L18
	}
L5:
	;
	v16 = int32(0)
	v17 = v13 - v7
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[1083]))))
	if v19 != 0 {
		v46 = v16
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v54 = int32(0)
	goto L4
L7:
	;
	v54 = v46
	goto L4
L8:
	;
	v20 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, _consts[1021]))
	if v21 == v20 {
		v46 = v16
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v24 == int32(0) {
		v46 = v16
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v28 = v21
	v31 = v24
	goto L12
L11:
	;
	v46 = v34 + int32(1)
	goto L7
L12:
	;
	v32 = F_strncmp(m, v7, v31, v17)
	mBase = m.M
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v39 != 0 {
		v28 = v28 + int32(4)
		v31 = v39
		goto L12
	} else {
		goto L17
	}
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v34 = v33 + v17
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v35 == int32(61) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v46 = v16
	goto L7
L18:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v57 != 0 {
		v164 = v54
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L3
L20:
	;
	v112 = int32(_a2184)
	v118 = F___strchrnul(m, v112, int32(61))
	mBase = m.M
	if v118 != v112 {
		goto L39
	} else {
		goto L40
	}
L21:
	;
	if v108 == int32(0) {
		goto L20
	} else {
		goto L35
	}
L22:
	;
	v70 = int32(0)
	v71 = v67 - v61
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v71))))
	if v73 != 0 {
		v100 = v70
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v108 = int32(0)
	goto L21
L24:
	;
	v108 = v100
	goto L21
L25:
	;
	v74 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, _consts[1021]))
	if v75 == v74 {
		v100 = v70
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v78 == int32(0) {
		v100 = v70
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v82 = v75
	v85 = v78
	goto L29
L28:
	;
	v100 = v88 + int32(1)
	goto L24
L29:
	;
	v86 = F_strncmp(m, v61, v85, v71)
	mBase = m.M
	if v86 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v93 != 0 {
		v82 = v82 + int32(4)
		v85 = v93
		goto L29
	} else {
		goto L34
	}
L32:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v88 = v87 + v71
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v89 == int32(61) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v100 = v70
	goto L24
L35:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v111 != 0 {
		v164 = v108
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L20
L37:
	;
	v164 = int32(_a2185)
	goto L1
L38:
	;
	if v159 == int32(0) {
		goto L37
	} else {
		goto L52
	}
L39:
	;
	v121 = int32(0)
	v122 = v118 - v112
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[1084]))))
	if v124 != 0 {
		v151 = v121
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v159 = int32(0)
	goto L38
L41:
	;
	v159 = v151
	goto L38
L42:
	;
	v125 = int32(0)
	v126 = *(*int32)(unsafe.Add(mBase, _consts[1021]))
	if v126 == v125 {
		v151 = v121
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v129 == int32(0) {
		v151 = v121
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v133 = v126
	v136 = v129
	goto L46
L45:
	;
	v151 = v139 + int32(1)
	goto L41
L46:
	;
	v137 = F_strncmp(m, v112, v136, v122)
	mBase = m.M
	if v137 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v144 != 0 {
		v133 = v133 + int32(4)
		v136 = v144
		goto L46
	} else {
		goto L51
	}
L49:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v139 = v138 + v122
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v140 == int32(61) {
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v151 = v121
	goto L41
L52:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v162 != 0 {
		v164 = v159
		goto L1
	} else {
		goto L53
	}
L53:
	;
	goto L37
L54:
	;
	v184 = int32(_a2185)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v185 == int32(46) {
		v192 = v184
		goto L65
	} else {
		goto L66
	}
L55:
	;
	v183 = v168
	goto L54
L56:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+v168))))
	if v172 == int32(0) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	if v172 == int32(47) {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v177 = int32(23)
	v179 = v168 + int32(1)
	if v179 != v177 {
		v168 = v179
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v183 = v177
	goto L54
L61:
	;
	return v348
L62:
	;
	v271 = int32(0)
	v272 = *(*int32)(unsafe.Add(mBase, _consts[1085]))
	if v272 == v271 {
		goto L91
	} else {
		goto L92
	}
L63:
	;
	if l0 != 0 {
		goto L88
	} else {
		goto L89
	}
L64:
	;
	v197 = int32(_a2185)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1086])))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	if v201 == int32(0) {
		v224 = v200
		v225 = v201
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
	if v193 == int32(0) {
		v263 = v192
		goto L63
	} else {
		goto L69
	}
L66:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+v183))))
	if v189 != 0 {
		v192 = v184
		goto L65
	} else {
		goto L67
	}
L67:
	;
	if v185 != int32(67) {
		v196 = v164
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v192 = v164
	goto L65
L69:
	;
	v196 = v192
	goto L64
L70:
	;
	if v225-v224&int32(255) == int32(0) {
		v263 = v196
		goto L63
	} else {
		goto L78
	}
L71:
	;
	goto L70
L72:
	;
	if v201 != v200&int32(255) {
		v224 = v200
		v225 = v201
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v207 = v196
	v208 = v197
	goto L74
L74:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
	if v212 == int32(0) {
		v224 = v211
		v225 = v212
		goto L71
	} else {
		goto L76
	}
L75:
	;
	v224 = v211
	v225 = v212
	goto L71
L76:
	;
	v215 = int32(1)
	if v212 == v211&int32(255) {
		v207 = v207 + v215
		v208 = v208 + v215
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v231 = int32(_a2186)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1087])))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	if v235 == int32(0) {
		v258 = v234
		v259 = v235
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v259-v258&int32(255) != 0 {
		goto L62
	} else {
		goto L87
	}
L80:
	;
	goto L79
L81:
	;
	if v235 != v234&int32(255) {
		v258 = v234
		v259 = v235
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v241 = v196
	v242 = v231
	goto L83
L83:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+1)))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)))
	if v246 == int32(0) {
		v258 = v245
		v259 = v246
		goto L80
	} else {
		goto L85
	}
L84:
	;
	v258 = v245
	v259 = v246
	goto L80
L85:
	;
	v249 = int32(1)
	if v246 == v245&int32(255) {
		v241 = v241 + v249
		v242 = v242 + v249
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v263 = v196
	goto L63
L88:
	;
	return int32(0)
L89:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+1)))
	if v265 == int32(46) {
		v348 = int32(_a2187)
		goto L61
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v322 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v322 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L92:
	;
	v277 = v272
	goto L93
L93:
	;
	v281 = v277 + int32(8)
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	if v285 == int32(0) {
		v308 = v284
		v309 = v285
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L91
L95:
	;
	if v309-v308&int32(255) == int32(0) {
		v348 = v277
		goto L61
	} else {
		goto L103
	}
L96:
	;
	goto L95
L97:
	;
	if v285 != v284&int32(255) {
		v308 = v284
		v309 = v285
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v291 = v196
	v292 = v281
	goto L99
L99:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+1)))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+1)))
	if v296 == int32(0) {
		v308 = v295
		v309 = v296
		goto L96
	} else {
		goto L101
	}
L100:
	;
	v308 = v295
	v309 = v296
	goto L96
L101:
	;
	v299 = int32(1)
	if v296 == v295&int32(255) {
		v291 = v291 + v299
		v292 = v292 + v299
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v277)+32))
	if v315 != 0 {
		v277 = v315
		goto L93
	} else {
		goto L104
	}
L104:
	;
	goto L94
L105:
	;
	if l0|v322 != 0 {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	v325 = int32(0)
	v326 = *(*int64)(unsafe.Add(mBase, _consts[1088]))
	*(*int64)(unsafe.Add(mBase, uint32(v322))) = v326
	v329 = v322 + int32(8)
	if v183 == v325 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v335 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v329+v183))) = uint8(v335)
	v338 = *(*int32)(unsafe.Add(mBase, _consts[1085]))
	*(*int32)(unsafe.Add(mBase, uint32(v322)+32)) = v338
	*(*int32)(unsafe.Add(mBase, _consts[1085])) = v322
	goto L105
L108:
	;
	goto L107
L109:
	;
	v332 = F__emscripten_memcpy_bulkmem(m, v329, v196, v183)
	mBase = m.M
	goto L108
L110:
	;
	v345 = v322
	goto L112
L111:
	;
	v345 = int32(_a2187)
	goto L112
L112:
	;
	v348 = v345
	goto L61
}
func F___get_tp(m *base.Module) int32 {
	return int32(9116560)
}
func F___getf2(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32 {
	var v8 int32
	_ = v8
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	v8 = int32(-1)
	v12 = l1 & int64(9223372036854775807)
	v13 = int64(9223090561878065152)
	if v12 == v13 {
		v17 = base.B2i32(l0 != int64(0))
	} else {
		v17 = base.B2i32(base.Ui64(v13) < base.Ui64(v12))
	}
	if v17 != 0 {
		v58 = v8
		return v58
	} else {
		v21 = l3 & int64(9223372036854775807)
		v22 = int64(9223090561878065152)
		if v21 == v22 {
			v26 = base.B2i32(l2 != int64(0))
		} else {
			v26 = base.B2i32(base.Ui64(v22) < base.Ui64(v21))
		}
		if v26 != 0 {
			v58 = v8
			return v58
		} else {
			if base.B2i32(l2|l0|(v21|v12) == int64(0)) == int32(0) {
				if l3&l1 < int64(0) {
					if l1 == l3 {
						v52 = base.B2i32(base.Ui64(l2) < base.Ui64(l0))
					} else {
						v52 = base.B2i32(l3 < l1)
					}
					if v52 != 0 {
						v58 = v8
					} else {
						v58 = base.B2i32(l0^l2|(l1^l3) != int64(0))
					}
					return v58
				} else {
					if l1 == l3 {
						v42 = base.B2i32(base.Ui64(l0) < base.Ui64(l2))
					} else {
						v42 = base.B2i32(l1 < l3)
					}
					if v42 != 0 {
						v58 = v8
						return v58
					} else {
						return base.B2i32(l0^l2|(l1^l3) != int64(0))
					}
				}
			} else {
				return int32(0)
			}
		}
	}
}
func F_genInfoSectionDict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
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
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
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
	var v246 int32
	_ = v246
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = l2
	goto L3
L2:
	;
	v11 = int32(_a1629)
	goto L3
L3:
	;
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v265
L5:
	;
	v54 = F_dictCreate(m, int32(_a1446))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L20
	}
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[906]))
	if v13 != 0 {
		v265 = v13
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v16 = F_dictCreate(m, int32(_a1446))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[906])) = v16
	v22 = F_dictExpand(m, v16, int32(16))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v24 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[906]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v26 == v24 {
		v265 = v25
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v31 = v26
	v34 = v11
	goto L12
L12:
	;
	v38 = F_sdsnew(m, v31)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[906]))
	return v51
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v47 != 0 {
		v31 = v47
		v34 = v34 + int32(4)
		goto L12
	} else {
		goto L19
	}
L15:
	;
	v41 = F_dictAdd(m, v25, v38, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	if v41 != int32(1) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	F_sdsfree(m, v38)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	goto L13
L20:
	;
	v56 = int32(16)
	if l1 < v56 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v59 = l1
	goto L23
L22:
	;
	v59 = v56
	goto L23
L23:
	;
	v60 = F_dictExpand(m, v54, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	if l1 < int32(1) {
		v265 = v54
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v72 = int32(0)
	goto L26
L26:
	;
	v76 = l0 + v72<<(uint(int32(2))%32)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v78 = F_objectGetVal(m, v77)
	mBase = m.M
	v79 = int32(_a27)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v82 != 0 {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	v265 = v54
	goto L4
L28:
	;
	v257 = v72 + int32(1)
	if v257 != l1 {
		v72 = v257
		goto L26
	} else {
		goto L88
	}
L29:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v143 = F_objectGetVal(m, v142)
	mBase = m.M
	v144 = int32(_a601)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v147 != 0 {
		goto L55
	} else {
		goto L56
	}
L30:
	;
	if v114-v116 != 0 {
		goto L29
	} else {
		goto L42
	}
L31:
	;
	v114 = F_tolower(m, v110)
	mBase = m.M
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v116 = F_tolower(m, v115)
	mBase = m.M
	goto L30
L32:
	;
	v84 = v78
	v85 = v79
	v86 = v82
	goto L35
L33:
	;
	v110 = int32(0)
	v111 = v79
	goto L31
L34:
	;
	v110 = v107 & int32(255)
	v111 = v106
	goto L31
L35:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v88 == int32(0) {
		v106 = v85
		v107 = v86
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v106 = v100
	v107 = int32(0)
	goto L34
L37:
	;
	v92 = v86 & int32(255)
	if v92 == v88 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v99 = int32(1)
	v100 = v85 + v99
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if v101 != 0 {
		v84 = v84 + v99
		v85 = v100
		v86 = v101
		goto L35
	} else {
		goto L41
	}
L39:
	;
	v94 = F_tolower(m, v92)
	mBase = m.M
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v96 = F_tolower(m, v95)
	mBase = m.M
	if v94 == v96 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	v106 = v85
	v107 = v98
	goto L34
L41:
	;
	goto L36
L42:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v118 == int32(0) {
		goto L28
	} else {
		goto L43
	}
L43:
	;
	v123 = v11
	v129 = v118
	goto L44
L44:
	;
	v130 = F_sdsnew(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L47
	}
L46:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v139 != 0 {
		v123 = v123 + int32(4)
		v129 = v139
		goto L44
	} else {
		goto L51
	}
L47:
	;
	v133 = F_dictAdd(m, v54, v130, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	if v133 != int32(1) {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	F_sdsfree(m, v130)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	goto L28
L52:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v188 = F_objectGetVal(m, v187)
	mBase = m.M
	v189 = int32(_a1630)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v192 != 0 {
		goto L70
	} else {
		goto L71
	}
L53:
	;
	if v179-v181 != 0 {
		goto L52
	} else {
		goto L65
	}
L54:
	;
	v179 = F_tolower(m, v175)
	mBase = m.M
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v181 = F_tolower(m, v180)
	mBase = m.M
	goto L53
L55:
	;
	v149 = v143
	v150 = v144
	v151 = v147
	goto L58
L56:
	;
	v175 = int32(0)
	v176 = v144
	goto L54
L57:
	;
	v175 = v172 & int32(255)
	v176 = v171
	goto L54
L58:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	if v153 == int32(0) {
		v171 = v150
		v172 = v151
		goto L57
	} else {
		goto L60
	}
L59:
	;
	v171 = v165
	v172 = int32(0)
	goto L57
L60:
	;
	v157 = v151 & int32(255)
	if v157 == v153 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v164 = int32(1)
	v165 = v150 + v164
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
	if v166 != 0 {
		v149 = v149 + v164
		v150 = v165
		v151 = v166
		goto L58
	} else {
		goto L64
	}
L62:
	;
	v159 = F_tolower(m, v157)
	mBase = m.M
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v161 = F_tolower(m, v160)
	mBase = m.M
	if v159 == v161 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	v171 = v150
	v172 = v163
	goto L57
L64:
	;
	goto L59
L65:
	;
	if l3 == int32(0) {
		goto L28
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(1)
	goto L28
L67:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v237 = F_objectGetVal(m, v236)
	mBase = m.M
	v238 = F_sdsnew(m, v237)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L8
	} else {
		goto L84
	}
L68:
	;
	if v224-v226 != 0 {
		goto L67
	} else {
		goto L80
	}
L69:
	;
	v224 = F_tolower(m, v220)
	mBase = m.M
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	v226 = F_tolower(m, v225)
	mBase = m.M
	goto L68
L70:
	;
	v194 = v188
	v195 = v189
	v196 = v192
	goto L73
L71:
	;
	v220 = int32(0)
	v221 = v189
	goto L69
L72:
	;
	v220 = v217 & int32(255)
	v221 = v216
	goto L69
L73:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v198 == int32(0) {
		v216 = v195
		v217 = v196
		goto L72
	} else {
		goto L75
	}
L74:
	;
	v216 = v210
	v217 = int32(0)
	goto L72
L75:
	;
	v202 = v196 & int32(255)
	if v202 == v198 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v209 = int32(1)
	v210 = v195 + v209
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	if v211 != 0 {
		v194 = v194 + v209
		v195 = v210
		v196 = v211
		goto L73
	} else {
		goto L79
	}
L77:
	;
	v204 = F_tolower(m, v202)
	mBase = m.M
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	v206 = F_tolower(m, v205)
	mBase = m.M
	if v204 == v206 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	v216 = v195
	v217 = v208
	goto L72
L79:
	;
	goto L74
L80:
	;
	if l4 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if l3 == int32(0) {
		goto L28
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(1)
	goto L28
L84:
	;
	v241 = F_dictAdd(m, v54, v238, int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	if v241 == int32(0) {
		goto L28
	} else {
		goto L86
	}
L86:
	;
	F_sdsfree(m, v238)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	goto L28
L88:
	;
	goto L27
}
func F_genModulesInfoString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
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
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int64
	_ = v210
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v215 int64
	_ = v215
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v220 int64
	_ = v220
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[541]))
	v19 = F_dictGetIterator(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v19)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L2
	} else {
		goto L69
	}
L2:
	;
	return int32(0)
L3:
	;
	v30 = v19 + int32(20)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v126 == int32(0) {
		v285 = l0
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v37 = v30
	v38 = v34
	goto L8
L6:
	;
	v34 = int32(1)
	goto L5
L7:
	;
	v34 = int32(0)
	goto L5
L8:
	;
	switch v38 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v38 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v118
	if v118 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v42 != int32(-1) {
		v81 = v42
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v82 = int32(1)
	v83 = v81 + v82
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v83
	v85 = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+v89+int32(26)))))
	if v93 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v46 != 0 {
		v81 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v48 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	if v75 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v47)+16)))
	v56 = int64(*(*int8)(unsafe.Add(mBase, uint32(v47)+27)))
	v57 = int64(*(*int32)(unsafe.Add(mBase, uint32(v47)+8)))
	v58 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v47)+12)))
	v59 = int64(*(*int8)(unsafe.Add(mBase, uint32(v47)+26)))
	v60 = int64(*(*int32)(unsafe.Add(mBase, uint32(v47)+4)))
	v61 = F_wangHash64(m, v60)
	mBase = m.M
	v63 = F_wangHash64(m, v59+v61)
	mBase = m.M
	v65 = F_wangHash64(m, v58+v63)
	mBase = m.M
	v67 = F_wangHash64(m, v57+v65)
	mBase = m.M
	v69 = F_wangHash64(m, v56+v67)
	mBase = m.M
	v71 = F_wangHash64(m, v55+v69)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v74 = v73
	goto L17
L19:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
	v53 = v51 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v53)
	v74 = v47
	goto L17
L20:
	;
	v81 = v75 + int32(-1)
	goto L14
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v81 = v78
	goto L14
L22:
	;
	v108 = int32(2)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v88+v106<<(uint(v108)%32)+int32(4))))
	v37 = v113 + v107<<(uint(v108)%32)
	v38 = int32(1)
	goto L8
L23:
	;
	v97 = v85
	goto L25
L24:
	;
	v97 = v82 << (uint(v93) % 32)
	goto L25
L25:
	;
	if v83 < v97 {
		v106 = v89
		v107 = v83
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v89 != 0 {
		v126 = v85
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	if v99 == int32(-1) {
		v126 = v85
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+4)) = int64(4294967296)
	v106 = int32(1)
	v107 = int32(0)
	goto L22
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v122
	v126 = v118
	goto L11
L30:
	;
	v138 = l0
	v141 = v126
	goto L31
L31:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	goto L33
L32:
	;
	v285 = v170
	goto L1
L33:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	goto L34
L34:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+20))
	v153 = F_genModulesInfoStringRenderModulesList(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151)+24))
	v156 = F_genModulesInfoStringRenderModulesList(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v158 = F_genModulesInfoStringRenderModuleOptions(m, v151)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v151)+8))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v151)+28))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(24)))) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(20)))) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(16)))) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v162
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v150
	v170 = F_sdscatfmt(m, v138, int32(_a971), v15)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	F_sdsfree(m, v153)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	F_sdsfree(m, v156)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	F_sdsfree(m, v158)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v185 = v19 + int32(20)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v186 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	if v281 != 0 {
		v138 = v170
		v141 = v281
		goto L31
	} else {
		goto L68
	}
L43:
	;
	v192 = v185
	v193 = v189
	goto L46
L44:
	;
	v189 = int32(1)
	goto L43
L45:
	;
	v189 = int32(0)
	goto L43
L46:
	;
	switch v193 {
	case 0:
		goto L51
	default:
		goto L50
	}
L48:
	;
	v193 = int32(0)
	goto L46
L49:
	;
	goto L42
L50:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v273
	if v273 == int32(0) {
		goto L48
	} else {
		goto L67
	}
L51:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v197 != int32(-1) {
		v236 = v197
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v237 = int32(1)
	v238 = v236 + v237
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v238
	v240 = int32(0)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v244+int32(26)))))
	if v248 == int32(255) {
		goto L61
	} else {
		goto L62
	}
L53:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v201 != 0 {
		v236 = int32(-1)
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v203 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+20))
	if v230 != int32(-1) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v210 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v202)+16)))
	v211 = int64(*(*int8)(unsafe.Add(mBase, uint32(v202)+27)))
	v212 = int64(*(*int32)(unsafe.Add(mBase, uint32(v202)+8)))
	v213 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v202)+12)))
	v214 = int64(*(*int8)(unsafe.Add(mBase, uint32(v202)+26)))
	v215 = int64(*(*int32)(unsafe.Add(mBase, uint32(v202)+4)))
	v216 = F_wangHash64(m, v215)
	mBase = m.M
	v218 = F_wangHash64(m, v214+v216)
	mBase = m.M
	v220 = F_wangHash64(m, v213+v218)
	mBase = m.M
	v222 = F_wangHash64(m, v212+v220)
	mBase = m.M
	v224 = F_wangHash64(m, v211+v222)
	mBase = m.M
	v226 = F_wangHash64(m, v210+v224)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v226
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v229 = v228
	goto L55
L57:
	;
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202)+24)))
	v208 = v206 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v202)+24)) = uint16(v208)
	v229 = v202
	goto L55
L58:
	;
	v236 = v230 + int32(-1)
	goto L52
L59:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v236 = v233
	goto L52
L60:
	;
	v263 = int32(2)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v243+v261<<(uint(v263)%32)+int32(4))))
	v192 = v268 + v262<<(uint(v263)%32)
	v193 = int32(1)
	goto L46
L61:
	;
	v252 = v240
	goto L63
L62:
	;
	v252 = v237 << (uint(v248) % 32)
	goto L63
L63:
	;
	if v238 < v252 {
		v261 = v244
		v262 = v238
		goto L60
	} else {
		goto L64
	}
L64:
	;
	if v244 != 0 {
		v281 = v240
		goto L49
	} else {
		goto L65
	}
L65:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v243)+20))
	if v254 == int32(-1) {
		v281 = v240
		goto L49
	} else {
		goto L66
	}
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+4)) = int64(4294967296)
	v261 = int32(1)
	v262 = int32(0)
	goto L60
L67:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v273)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = v277
	v281 = v273
	goto L49
L68:
	;
	goto L32
L69:
	;
	m.G0 = v15 + int32(32)
	return v285
}
func F_genModulesInfoStringRenderModulesList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = v7 + int32(8)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v11
	goto L1
L1:
	;
	v16 = F_sdsnew(m, int32(_a968))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	v21 = v7 + int32(8)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v23 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v69 = F_sdscat(m, v66, int32(_a969))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L19
	}
L5:
	;
	if v23 == int32(0) {
		v66 = v16
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23+base.B2i32(v26 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v32
	goto L6
L8:
	;
	v38 = v16
	v39 = v23
	goto L9
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v42 = F_sdscat(m, v38, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L11
	}
L10:
	;
	v66 = v49
	goto L4
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v39 == v44 {
		v49 = v42
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v51 = v7 + int32(8)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v53 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v47 = F_sdscat(m, v42, int32(_a970))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v49 = v47
	goto L12
L15:
	;
	if v53 != 0 {
		v38 = v49
		v39 = v53
		goto L9
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v53+base.B2i32(v56 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v62
	goto L16
L18:
	;
	goto L10
L19:
	;
	m.G0 = v7 + int32(16)
	return v69
}
func F_generateSkyscraper(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v114 int64
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v16 < int32(1) {
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v21 = v19 + int32(-1)
		v22 = v21 - v16
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v33 = v29
		v34 = v21
		for {
			if v33 < int32(1) {
				v204 = v33
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v57 = v50
				v58 = v33
				v65 = v50
				v66 = v33 + v50
				for {
					if v34 != v22+int32(1) {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						if v77 == int32(0) {
							v153 = v76
						} else {
							v81 = v65 + int32(1)
							if v57 <= v81 {
								v153 = v76
							} else {
								if v66+int32(-2) <= v57 {
									v153 = v76
								} else {
									if v34 <= v22+int32(2) {
										v153 = v76
									} else {
										if v19+int32(-2) <= v34 {
											v153 = v76
										} else {
											v86 = v57 - v81
											v88 = base.I32_div_s(v86, int32(2))
											if v88&((v34-v22)&int32(1)) == int32(0) {
												v153 = v76
											} else {
												for {
													v108 = int32(0)
													v110 = *(*int64)(unsafe.Add(mBase, _consts[408]))
													v114 = v110*int64(6364136223846793005) + int64(1)
													*(*int64)(unsafe.Add(mBase, _consts[408])) = v114
													v120 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v114)>>(uint(int64(33))%64))), int32(2))
													v122 = v120 + int32(1)
													v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
													if v122 == v123 {
														continue
													} else {
														break
													}
													break
												}
												if v86&int32(1) == int32(0) {
													v153 = v122
												} else {
													v130 = v57 + int32(-1)
													v131 = int32(0)
													if v130 < v131 {
														v148 = v131
													} else {
														v136 = int32(0)
														if v34 < v136 {
															v148 = v136
														} else {
															v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															if v139 <= v130 {
																v148 = v136
															} else {
																v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																if v141 <= v34 {
																	v148 = v136
																} else {
																	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v143+v130+v139*v34))))
																	v148 = v147
																}
															}
														}
													}
													v153 = v148
												}
											}
										}
									}
								}
							}
						}
						if v57 < int32(0) {
						} else {
							if v34 < int32(0) {
							} else {
								v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if v170 <= v57 {
								} else {
									v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v172 <= v34 {
									} else {
										v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*uint8)(unsafe.Add(mBase, uint32(v174+v57+v170*v34))) = uint8(v153)
									}
								}
							}
						}
						v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v185 = v180
						v192 = v181
					} else {
						if v57 <= v65+int32(1) {
							v185 = v58
							v192 = v65
						} else {
							if v66+int32(-2) <= v57 {
								v185 = v58
								v192 = v65
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								if v77 == int32(0) {
									v153 = v76
								} else {
									v81 = v65 + int32(1)
									if v57 <= v81 {
										v153 = v76
									} else {
										if v66+int32(-2) <= v57 {
											v153 = v76
										} else {
											if v34 <= v22+int32(2) {
												v153 = v76
											} else {
												if v19+int32(-2) <= v34 {
													v153 = v76
												} else {
													v86 = v57 - v81
													v88 = base.I32_div_s(v86, int32(2))
													if v88&((v34-v22)&int32(1)) == int32(0) {
														v153 = v76
													} else {
														for {
															v108 = int32(0)
															v110 = *(*int64)(unsafe.Add(mBase, _consts[408]))
															v114 = v110*int64(6364136223846793005) + int64(1)
															*(*int64)(unsafe.Add(mBase, _consts[408])) = v114
															v120 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v114)>>(uint(int64(33))%64))), int32(2))
															v122 = v120 + int32(1)
															v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
															if v122 == v123 {
																continue
															} else {
																break
															}
															break
														}
														if v86&int32(1) == int32(0) {
															v153 = v122
														} else {
															v130 = v57 + int32(-1)
															v131 = int32(0)
															if v130 < v131 {
																v148 = v131
															} else {
																v136 = int32(0)
																if v34 < v136 {
																	v148 = v136
																} else {
																	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																	if v139 <= v130 {
																		v148 = v136
																	} else {
																		v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																		if v141 <= v34 {
																			v148 = v136
																		} else {
																			v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v143+v130+v139*v34))))
																			v148 = v147
																		}
																	}
																}
															}
															v153 = v148
														}
													}
												}
											}
										}
									}
								}
								if v57 < int32(0) {
								} else {
									if v34 < int32(0) {
									} else {
										v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										if v170 <= v57 {
										} else {
											v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											if v172 <= v34 {
											} else {
												v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*uint8)(unsafe.Add(mBase, uint32(v174+v57+v170*v34))) = uint8(v153)
											}
										}
									}
								}
								v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v185 = v180
								v192 = v181
							}
						}
					}
					v198 = v57 + int32(1)
					v199 = v185 + v192
					if v198 < v199 {
						v57 = v198
						v58 = v185
						v65 = v192
						v66 = v199
						continue
					} else {
						break
					}
					break
				}
				v204 = v185
			}
			v217 = v34 + int32(-1)
			if v22 < v217 {
				v33 = v204
				v34 = v217
				continue
			} else {
				break
			}
			break
		}
	}
	return
}
func F_geoGetPointsInRange(m *base.Module, l0 int32, l1 float64, l2 float64, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v67 int32
	_ = v67
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 float64
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v112 int64
	_ = v112
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 float64
	_ = v132
	var v133 float64
	_ = v133
	var v134 float64
	_ = v134
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v137 float64
	_ = v137
	var v139 int32
	_ = v139
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v143 float64
	_ = v143
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v147 float64
	_ = v147
	var v148 float64
	_ = v148
	var v149 int32
	_ = v149
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int64
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
	var v174 int32
	_ = v174
	var v175 float64
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 float64
	_ = v202
	var v204 float64
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v249 float64
	_ = v249
	var v250 float64
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 float64
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 float64
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var __phi289 int32
	_ = __phi289
	var v301 int32
	_ = v301
	var __phi301 int32
	_ = __phi301
	var v304 int32
	_ = v304
	var __phi304 int32
	_ = __phi304
	var v308 float64
	_ = v308
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v680 int32
	_ = v680
	var __phi680 int32
	_ = __phi680
	var v689 int32
	_ = v689
	var __phi689 int32
	_ = __phi689
	var v691 int32
	_ = v691
	var __phi691 int32
	_ = __phi691
	var v693 float64
	_ = v693
	var v696 int32
	_ = v696
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v724 int32
	_ = v724
	var v733 int32
	_ = v733
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v980 float64
	_ = v980
	var v983 int32
	_ = v983
	var v1003 int32
	_ = v1003
	var v1053 int32
	_ = v1053
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1090 float64
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1094 float64
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1101 float64
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1127 int64
	_ = v1127
	var v1136 int64
	_ = v1136
	var v1138 int64
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 float64
	_ = v1147
	var v1148 float64
	_ = v1148
	var v1149 float64
	_ = v1149
	var v1150 float64
	_ = v1150
	var v1151 float64
	_ = v1151
	var v1152 float64
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1155 float64
	_ = v1155
	var v1156 float64
	_ = v1156
	var v1158 float64
	_ = v1158
	var v1160 float64
	_ = v1160
	var v1161 float64
	_ = v1161
	var v1162 float64
	_ = v1162
	var v1163 float64
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 float64
	_ = v1165
	var v1166 float64
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1175 int32
	_ = v1175
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1192 float64
	_ = v1192
	var v1193 float64
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1223 float64
	_ = v1223
	var v1225 float64
	_ = v1225
	var v1231 int32
	_ = v1231
	var v1236 int32
	_ = v1236
	var v1240 int32
	_ = v1240
	var v1254 int32
	_ = v1254
	var v1265 int32
	_ = v1265
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	*(*int64)(unsafe.Add(mBase, uint32(v16)+72)) = int64(4294967296)
	*(*float64)(unsafe.Add(mBase, uint32(v16)+64)) = l2
	*(*float64)(unsafe.Add(mBase, uint32(v16)+56)) = l1
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v23)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L3
	default:
		goto L2
	case 4:
		goto L4
	}
L1:
	;
	m.G0 = v16 + int32(80)
	return v1265
L2:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v1265 = v1254 - v22
	goto L1
L3:
	;
	v225 = int32(0)
	v226 = F_objectGetVal(m, l0)
	mBase = m.M
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	v229 = v16 + int32(56)
	v249 = *(*float64)(unsafe.Add(mBase, uint32(v229)))
	v250 = *(*float64)(unsafe.Add(mBase, uint32(v229)+8))
	if base.F64_gt(v249, v250) != 0 {
		v1053 = v225
		goto L53
	} else {
		goto L54
	}
L4:
	;
	v30 = F_objectGetVal(m, l0)
	mBase = m.M
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = int64(0)
	v38 = F_zzlFirstInRange(m, v30, v16+int32(56))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v38
	if v38 == int32(0) {
		v1265 = v31
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v45 = F_lpNext(m, v30, v38)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v45
	v51 = l4 + int32(16)
	goto L9
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
	v68 = F_zzlGetScore(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v71 = v16 + int32(56)
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v71)+8))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	if v76 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v77 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L13:
	;
	v77 = base.F64_lt(v68, v73)
	goto L15
L14:
	;
	v77 = base.F64_le(v68, v73)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	v85 = F_lpGetValue(m, v80, v16+int32(44), v16+int32(32))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v88 = v16 + int32(16)
	v89 = int32(8)
	v90 = v16 + v89
	v91 = int32(0)
	v94 = m.G0
	v96 = v94 - int32(32)
	m.G0 = v96
	v101 = v96 + int32(24)
	v102 = int32(26)
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v96+int32(28)))) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v96)+25)) = v91
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
	*(*int64)(unsafe.Add(mBase, uint32(v96+v89))) = v112
	if base.F64_lt(v68, float64(1.8446744073709552e+19))&base.F64_ge(v68, float64(0)) == v91 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	if base.Ui32(l5+int32(-1)) < base.Ui32(v212) {
		goto L2
	} else {
		goto L49
	}
L19:
	;
	if v85 != 0 {
		goto L36
	} else {
		goto L37
	}
L20:
	;
	if v160 == int32(0) {
		goto L19
	} else {
		goto L34
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v96)+16)) = v123
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = v123
	v126 = F_geohashDecodeToLongLatWGS84(m, v96, v88)
	mBase = m.M
	if v126 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v123 = int64(0)
	goto L21
L23:
	;
	v121 = base.I64_trunc_f64_u(v68)
	v123 = v121
	goto L21
L24:
	;
	m.G0 = v96 + int32(32)
	goto L20
L25:
	;
	v128 = int32(-1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	switch v129 + v128 {
	case 0:
		goto L30
	case 1:
		goto L29
	case 2:
		goto L28
	default:
		goto L27
	}
L26:
	;
	v160 = int32(-1)
	goto L24
L27:
	;
	v160 = int32(0)
	goto L24
L28:
	;
	v150 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v151 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v154 = F_geohashGetDistanceIfInPolygon(m, v150, v151, v88, v152, v153, v90)
	mBase = m.M
	if v154 == int32(0) {
		v160 = v128
		goto L24
	} else {
		goto L33
	}
L29:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(l3)+72))
	v141 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v143 = *(*float64)(unsafe.Add(mBase, uint32(l3)+64))
	v145 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v146 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v147 = *(*float64)(unsafe.Add(mBase, uint32(v88)))
	v148 = *(*float64)(unsafe.Add(mBase, uint32(v88)+8))
	v149 = F_geohashGetDistanceIfInRectangle(m, base.F64_mul(v140, v141), base.F64_mul(v141, v143), v145, v146, v147, v148, v90)
	mBase = m.M
	if v149 != 0 {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	v132 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v133 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v134 = *(*float64)(unsafe.Add(mBase, uint32(v88)))
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v88)+8))
	v136 = *(*float64)(unsafe.Add(mBase, uint32(l3)+64))
	v137 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v139 = F_geohashGetDistanceIfInRadiusWGS84(m, v132, v133, v134, v135, base.F64_mul(v136, v137), v90)
	mBase = m.M
	if v139 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v160 = v128
	goto L24
L32:
	;
	v160 = v128
	goto L24
L33:
	;
	goto L27
L34:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v212 = v166
	goto L18
L35:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v175 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v176 != v177 {
		v196 = v174
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v171 = F_sdsnewlen(m, v85, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L5
	} else {
		goto L39
	}
L37:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
	v168 = F_sdsfromlonglong(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v173 = v168
	goto L35
L39:
	;
	v173 = v171
	goto L35
L40:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v201 = v196 + v198*int32(40)
	v202 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v201))) = v202
	v204 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v201)+32)) = v173
	*(*float64)(unsafe.Add(mBase, uint32(v201)+16)) = v175
	*(*float64)(unsafe.Add(mBase, uint32(v201)+8)) = v204
	*(*float64)(unsafe.Add(mBase, uint32(v201)+24)) = v68
	v210 = v198 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v210
	v212 = v210
	goto L18
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v176 << (uint(int32(1)) % 32)
	v183 = v176 * int32(80)
	if v174 != v51 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v193 = F_valkey_realloc(m, v174, v183)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L48
	}
L43:
	;
	v185 = F_valkey_malloc(m, v183)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v185
	goto L47
L45:
	;
	v196 = v185
	goto L40
L46:
	;
	goto L45
L47:
	;
	v191 = F__emscripten_memcpy_bulkmem(m, v185, v51, int32(320))
	mBase = m.M
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v193
	v196 = v193
	goto L40
L49:
	;
	F_zzlNext(m, v30, v16+int32(52), v16+int32(48))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	if v224 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	goto L2
L52:
	;
	if v1053 == int32(0) {
		v1265 = v225
		goto L1
	} else {
		goto L194
	}
L53:
	;
	goto L52
L54:
	;
	if base.F64_ne(v249, v250) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	if v255 == int32(0) {
		v1053 = v225
		goto L53
	} else {
		goto L59
	}
L56:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v229)+16))
	if v253 != 0 {
		v1053 = v225
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v229)+20))
	if v254 != 0 {
		v1053 = v225
		goto L53
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	v258 = *(*float64)(unsafe.Add(mBase, uint32(v255)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v229)+16))
	if v261 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v262 = base.F64_gt(v258, v249)
	goto L62
L61:
	;
	v262 = base.F64_ge(v258, v249)
	goto L62
L62:
	;
	if v262 != int32(1) {
		v1053 = v225
		goto L53
	} else {
		goto L63
	}
L63:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	if v265 == int32(0) {
		v1053 = v225
		goto L53
	} else {
		goto L64
	}
L64:
	;
	v268 = *(*float64)(unsafe.Add(mBase, uint32(v265)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v229)+20))
	if v271 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v272 = base.F64_lt(v268, v250)
	goto L67
L66:
	;
	v272 = base.F64_le(v268, v250)
	goto L67
L67:
	;
	if v272 != int32(1) {
		v1053 = v225
		goto L53
	} else {
		goto L68
	}
L68:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v227)+16))
	v281 = (v277 + int32(-1)) << (uint(int32(3)) % 32)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v227+int32(12)+v281)))
	if v283 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L84
L70:
	;
	__phi289 = v283
	__phi301 = int32(0)
	__phi304 = v227
	v289 = __phi289
	v301 = __phi301
	v304 = __phi304
	goto L72
L71:
	;
	v338 = int32(0)
	v339 = v227
	goto L69
L72:
	;
	v308 = *(*float64)(unsafe.Add(mBase, uint32(v289)))
	if v261 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v338 = v320
	v339 = v289
	goto L69
L74:
	;
	if v277 < int32(2) {
		v319 = int32(1)
		goto L79
	} else {
		goto L80
	}
L75:
	;
	v311 = base.F64_gt(v308, v249)
	goto L77
L76:
	;
	v311 = base.F64_ge(v308, v249)
	goto L77
L77:
	;
	if v311 == int32(0) {
		goto L74
	} else {
		goto L78
	}
L78:
	;
	v338 = v301
	v339 = v304
	goto L69
L79:
	;
	v320 = v319 + v301
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v289+v281+int32(12))))
	if v324 != 0 {
		__phi289 = v324
		__phi301 = v320
		__phi304 = v289
		v289 = __phi289
		v301 = __phi301
		v304 = __phi304
		goto L72
	} else {
		goto L81
	}
L80:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v304+v281+int32(16))))
	v319 = v318
	goto L79
L81:
	;
	goto L73
L84:
	;
	if v277 < int32(2) {
		v733 = v339
		v748 = v338
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v752 = int32(0)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	if base.Ui32(v754) <= base.Ui32(v748+v225) {
		v1053 = v752
		goto L53
	} else {
		goto L156
	}
L139:
	;
	v648 = v339
	v654 = v277 + int32(-2)
	v663 = v338
	goto L140
L140:
	;
	v668 = v654 << (uint(int32(3)) % 32)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v648+v668)+12))
	if v670 == int32(0) {
		v709 = v648
		v724 = v663
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v733 = v709
	v748 = v724
	goto L138
L142:
	;
	if int32(0) < v654 {
		v648 = v709
		v654 = v654 + int32(-1)
		v663 = v724
		goto L140
	} else {
		goto L155
	}
L143:
	;
	__phi680 = v670
	__phi689 = v663
	__phi691 = v648
	v680 = __phi680
	v689 = __phi689
	v691 = __phi691
	goto L144
L144:
	;
	v693 = *(*float64)(unsafe.Add(mBase, uint32(v680)))
	if v261 != 0 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v709 = v680
	v724 = v705
	goto L142
L146:
	;
	if v654 != 0 {
		goto L152
	} else {
		goto L153
	}
L147:
	;
	v696 = base.F64_gt(v693, v249)
	goto L149
L148:
	;
	v696 = base.F64_ge(v693, v249)
	goto L149
L149:
	;
	if v696 == int32(0) {
		goto L146
	} else {
		goto L150
	}
L150:
	;
	v709 = v691
	v724 = v689
	goto L142
L151:
	;
	v705 = v704 + v689
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v680+v668)+12))
	if v707 != 0 {
		__phi680 = v707
		__phi689 = v705
		__phi691 = v680
		v680 = __phi680
		v689 = __phi689
		v691 = __phi691
		goto L144
	} else {
		goto L154
	}
L152:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v691+v668+int32(16))))
	v704 = v703
	goto L151
L153:
	;
	v704 = int32(1)
	goto L151
L154:
	;
	goto L145
L155:
	;
	goto L141
L156:
	;
	goto L160
L157:
	;
	v1053 = v1003
	goto L53
L158:
	;
	v980 = *(*float64)(unsafe.Add(mBase, uint32(v841)))
	if v271 != 0 {
		goto L189
	} else {
		goto L190
	}
L160:
	;
	goto L161
L161:
	;
	goto L167
L166:
	;
	if v841 != 0 {
		goto L158
	} else {
		goto L171
	}
L167:
	;
	v822 = v733
	v828 = int32(0)
	goto L168
L168:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v822)+12))
	v843 = v828 + int32(1)
	if v843 != int32(1) {
		v822 = v841
		v828 = v843
		goto L168
	} else {
		goto L170
	}
L169:
	;
	goto L166
L170:
	;
	goto L169
L171:
	;
	v1003 = int32(0)
	goto L157
L189:
	;
	v983 = base.F64_lt(v980, v250)
	goto L191
L190:
	;
	v983 = base.F64_le(v980, v250)
	goto L191
L191:
	;
	if v983 != int32(1) {
		v1053 = v752
		goto L53
	} else {
		goto L192
	}
L192:
	;
	v1003 = v841
	goto L157
L194:
	;
	v1074 = l4 + int32(16)
	v1075 = v1053
	goto L195
L195:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = int64(0)
	v1090 = *(*float64)(unsafe.Add(mBase, uint32(v1075)))
	v1092 = v16 + int32(56)
	v1094 = *(*float64)(unsafe.Add(mBase, uint32(v1092)+8))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1092)+20))
	if v1097 != 0 {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	goto L2
L197:
	;
	if v1098 == int32(0) {
		goto L2
	} else {
		goto L201
	}
L198:
	;
	v1098 = base.F64_lt(v1090, v1094)
	goto L200
L199:
	;
	v1098 = base.F64_le(v1090, v1094)
	goto L200
L200:
	;
	goto L197
L201:
	;
	v1101 = *(*float64)(unsafe.Add(mBase, uint32(v1075)))
	v1103 = v16 + int32(16)
	v1104 = int32(32)
	v1105 = v16 + v1104
	v1106 = int32(0)
	v1109 = m.G0
	v1111 = v1109 - v1104
	m.G0 = v1111
	v1116 = v1111 + int32(24)
	v1117 = int32(26)
	*(*uint8)(unsafe.Add(mBase, uint32(v1116))) = uint8(v1117)
	*(*int32)(unsafe.Add(mBase, uint32(v1111+int32(28)))) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v1111)+25)) = v1106
	v1127 = *(*int64)(unsafe.Add(mBase, uint32(v1116)))
	*(*int64)(unsafe.Add(mBase, uint32(v1111+int32(8)))) = v1127
	if base.F64_lt(v1101, float64(1.8446744073709552e+19))&base.F64_ge(v1101, float64(0)) == v1106 {
		goto L206
	} else {
		goto L207
	}
L202:
	;
	if base.Ui32(l5+int32(-1)) < base.Ui32(v1236) {
		goto L2
	} else {
		goto L230
	}
L203:
	;
	v1183 = v1075 + int32(16)
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1183)))
	v1187 = v1183 + v1184<<(uint(int32(3))%32)
	v1188 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1187))))
	goto L219
L204:
	;
	if v1175 == int32(0) {
		goto L203
	} else {
		goto L218
	}
L205:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1111)+16)) = v1138
	*(*int64)(unsafe.Add(mBase, uint32(v1111))) = v1138
	v1141 = F_geohashDecodeToLongLatWGS84(m, v1111, v1103)
	mBase = m.M
	if v1141 != 0 {
		goto L209
	} else {
		goto L210
	}
L206:
	;
	v1138 = int64(0)
	goto L205
L207:
	;
	v1136 = base.I64_trunc_f64_u(v1101)
	v1138 = v1136
	goto L205
L208:
	;
	m.G0 = v1111 + int32(32)
	goto L204
L209:
	;
	v1143 = int32(-1)
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	switch v1144 + v1143 {
	case 0:
		goto L214
	case 1:
		goto L213
	case 2:
		goto L212
	default:
		goto L211
	}
L210:
	;
	v1175 = int32(-1)
	goto L208
L211:
	;
	v1175 = int32(0)
	goto L208
L212:
	;
	v1165 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v1166 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v1169 = F_geohashGetDistanceIfInPolygon(m, v1165, v1166, v1103, v1167, v1168, v1105)
	mBase = m.M
	if v1169 == int32(0) {
		v1175 = v1143
		goto L208
	} else {
		goto L217
	}
L213:
	;
	v1155 = *(*float64)(unsafe.Add(mBase, uint32(l3)+72))
	v1156 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v1158 = *(*float64)(unsafe.Add(mBase, uint32(l3)+64))
	v1160 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v1161 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v1162 = *(*float64)(unsafe.Add(mBase, uint32(v1103)))
	v1163 = *(*float64)(unsafe.Add(mBase, uint32(v1103)+8))
	v1164 = F_geohashGetDistanceIfInRectangle(m, base.F64_mul(v1155, v1156), base.F64_mul(v1156, v1158), v1160, v1161, v1162, v1163, v1105)
	mBase = m.M
	if v1164 != 0 {
		goto L211
	} else {
		goto L216
	}
L214:
	;
	v1147 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v1148 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v1149 = *(*float64)(unsafe.Add(mBase, uint32(v1103)))
	v1150 = *(*float64)(unsafe.Add(mBase, uint32(v1103)+8))
	v1151 = *(*float64)(unsafe.Add(mBase, uint32(l3)+64))
	v1152 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v1154 = F_geohashGetDistanceIfInRadiusWGS84(m, v1147, v1148, v1149, v1150, base.F64_mul(v1151, v1152), v1105)
	mBase = m.M
	if v1154 != 0 {
		goto L211
	} else {
		goto L215
	}
L215:
	;
	v1175 = v1143
	goto L208
L216:
	;
	v1175 = v1143
	goto L208
L217:
	;
	goto L211
L218:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v1236 = v1181
	goto L202
L219:
	;
	v1192 = *(*float64)(unsafe.Add(mBase, uint32(v1075)))
	v1193 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
	v1194 = F_sdsdup(m, v1187+v1188+int32(1))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L5
	} else {
		goto L220
	}
L220:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v1197 != v1198 {
		v1217 = v1196
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v1222 = v1217 + v1219*int32(40)
	v1223 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v1222))) = v1223
	v1225 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1222)+32)) = v1194
	*(*float64)(unsafe.Add(mBase, uint32(v1222)+16)) = v1193
	*(*float64)(unsafe.Add(mBase, uint32(v1222)+8)) = v1225
	*(*float64)(unsafe.Add(mBase, uint32(v1222)+24)) = v1192
	v1231 = v1219 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v1231
	v1236 = v1231
	goto L202
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v1197 << (uint(int32(1)) % 32)
	v1204 = v1197 * int32(80)
	if v1196 != v1074 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1214 = F_valkey_realloc(m, v1196, v1204)
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L5
	} else {
		goto L229
	}
L224:
	;
	v1206 = F_valkey_malloc(m, v1204)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L5
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1206
	goto L228
L226:
	;
	v1217 = v1206
	goto L221
L227:
	;
	goto L226
L228:
	;
	v1212 = F__emscripten_memcpy_bulkmem(m, v1206, v1074, int32(320))
	mBase = m.M
	goto L227
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1214
	v1217 = v1214
	goto L221
L230:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+12))
	if v1240 != 0 {
		v1075 = v1240
		goto L195
	} else {
		goto L231
	}
L231:
	;
	goto L196
}
func F_geoWithinShape(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int64
	_ = v26
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v70 float64
	_ = v70
	var v77 float64
	_ = v77
	var v83 float64
	_ = v83
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v100 float64
	_ = v100
	var v107 float64
	_ = v107
	var v110 float64
	_ = v110
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v122 float64
	_ = v122
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v127 float64
	_ = v127
	var v128 int32
	_ = v128
	var v132 float64
	_ = v132
	var v133 float64
	_ = v133
	var v135 float64
	_ = v135
	var v143 float64
	_ = v143
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v147 float64
	_ = v147
	var v151 float64
	_ = v151
	var v160 float64
	_ = v160
	var v163 float64
	_ = v163
	var v170 float64
	_ = v170
	var v173 float64
	_ = v173
	var v178 float64
	_ = v178
	var v182 float64
	_ = v182
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v195 float64
	_ = v195
	var v202 float64
	_ = v202
	var v207 float64
	_ = v207
	var v215 int32
	_ = v215
	var v218 float64
	_ = v218
	var v219 float64
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v234 int32
	_ = v234
	var v238 float64
	_ = v238
	var v239 float64
	_ = v239
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var __phi248 int32
	_ = __phi248
	var v249 float64
	_ = v249
	var __phi249 float64
	_ = __phi249
	var v251 int32
	_ = v251
	var __phi251 int32
	_ = __phi251
	var v252 int32
	_ = v252
	var __phi252 int32
	_ = __phi252
	var v259 int32
	_ = v259
	var v260 float64
	_ = v260
	var v263 float64
	_ = v263
	var v264 float64
	_ = v264
	var v269 float64
	_ = v269
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 float64
	_ = v286
	var v287 float64
	_ = v287
	var v294 float64
	_ = v294
	var v295 float64
	_ = v295
	var v301 float64
	_ = v301
	var v310 float64
	_ = v310
	var v311 float64
	_ = v311
	var v312 float64
	_ = v312
	var v314 float64
	_ = v314
	var v315 float64
	_ = v315
	var v319 float64
	_ = v319
	var v326 float64
	_ = v326
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v15 = v10 + int32(24)
	v16 = int32(26)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v16)
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(28)))) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v10)+25)) = v5
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(8)))) = v26
	if base.F64_lt(l1, float64(1.8446744073709552e+19))&base.F64_ge(l1, float64(0)) == v5 {
		v37 = int64(0)
	} else {
		v35 = base.I64_trunc_f64_u(l1)
		v37 = v35
	}
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v37
	v41 = m.G0
	v42 = int32(16)
	v43 = v41 - v42
	m.G0 = v43
	v45 = int32(8)
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v10+v45)))
	*(*int64)(unsafe.Add(mBase, uint32(v43+v45))) = v49
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	*(*int64)(unsafe.Add(mBase, uint32(v43))) = v51
	v53 = F_geohashDecodeToLongLatType(m, v43, l2)
	mBase = m.M
	m.G0 = v43 + v42
	if v53 != 0 {
		v58 = int32(-1)
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v59 + v58 {
		case 0:
			v62 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v63 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
			v64 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
			v65 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
			v66 = *(*float64)(unsafe.Add(mBase, uint32(l0)+64))
			v67 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
			v70 = float64(0.017453292519943295)
			v77 = F_sin(m, base.F64_mul(base.F64_sub(base.F64_mul(v64, v70), base.F64_mul(v62, v70)), float64(0.5)))
			mBase = m.M
			if base.F64_le(base.F64_abs(v77), float64(1e-15)) == int32(0) {
				v91 = float64(0.017453292519943295)
				v92 = base.F64_mul(v63, v91)
				v93 = F_cos(m, v92)
				mBase = m.M
				v95 = base.F64_mul(v65, v91)
				v96 = F_cos(m, v95)
				mBase = m.M
				v100 = F_sin(m, base.F64_mul(base.F64_sub(v95, v92), float64(0.5)))
				mBase = m.M
				v107 = F_asin(m, base.F64_sqrt(base.F64_add(base.F64_mul(v100, v100), base.F64_mul(v77, base.F64_mul(base.F64_mul(v93, v96), v77)))))
				mBase = m.M
				v110 = base.F64_mul(v107, float64(1.2745595121712e+07))
			} else {
				v83 = float64(0.017453292519943295)
				v110 = base.F64_mul(base.F64_abs(base.F64_sub(base.F64_mul(v65, v83), base.F64_mul(v63, v83))), float64(6.372797560856e+06))
			}
			*(*float64)(unsafe.Add(mBase, uint32(l3))) = v110
			if base.F64_gt(v110, base.F64_mul(v66, v67))^int32(1) != 0 {
				v350 = int32(0)
			} else {
				v350 = v58
			}
		case 1:
			v119 = *(*float64)(unsafe.Add(mBase, uint32(l0)+72))
			v120 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
			v122 = *(*float64)(unsafe.Add(mBase, uint32(l0)+64))
			v124 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v125 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
			v126 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
			v127 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
			v128 = int32(0)
			v132 = float64(0.017453292519943295)
			v133 = base.F64_mul(v125, v132)
			v135 = base.F64_mul(v127, v132)
			if base.F64_gt(base.F64_mul(base.F64_abs(base.F64_sub(v133, v135)), float64(6.372797560856e+06)), base.F64_mul(base.F64_mul(v120, v122), float64(0.5))) != 0 {
				v215 = v128
			} else {
				v143 = base.F64_sub(v135, v135)
				v144 = float64(0.017453292519943295)
				v145 = base.F64_mul(v124, v144)
				v147 = base.F64_mul(v126, v144)
				v151 = F_sin(m, base.F64_mul(base.F64_sub(v145, v147), float64(0.5)))
				mBase = m.M
				if base.F64_le(base.F64_abs(v151), float64(1e-15)) == int32(0) {
					v160 = F_cos(m, v135)
					mBase = m.M
					v163 = F_sin(m, base.F64_mul(v143, float64(0.5)))
					mBase = m.M
					v170 = F_asin(m, base.F64_sqrt(base.F64_add(base.F64_mul(v163, v163), base.F64_mul(v151, base.F64_mul(base.F64_mul(v160, v160), v151)))))
					mBase = m.M
					v173 = base.F64_mul(v170, float64(1.2745595121712e+07))
				} else {
					v173 = base.F64_mul(base.F64_abs(v143), float64(6.372797560856e+06))
				}
				if base.F64_gt(v173, base.F64_mul(base.F64_mul(v119, v120), float64(0.5))) != 0 {
					v215 = v128
				} else {
					v178 = base.F64_sub(v135, v133)
					v182 = F_sin(m, base.F64_mul(base.F64_sub(v147, v145), float64(0.5)))
					mBase = m.M
					if base.F64_le(base.F64_abs(v182), float64(1e-15)) == int32(0) {
						v191 = F_cos(m, v133)
						mBase = m.M
						v192 = F_cos(m, v135)
						mBase = m.M
						v195 = F_sin(m, base.F64_mul(v178, float64(0.5)))
						mBase = m.M
						v202 = F_asin(m, base.F64_sqrt(base.F64_add(base.F64_mul(v195, v195), base.F64_mul(v182, base.F64_mul(base.F64_mul(v191, v192), v182)))))
						mBase = m.M
						v207 = base.F64_mul(v202, float64(1.2745595121712e+07))
					} else {
						v207 = base.F64_mul(base.F64_abs(v178), float64(6.372797560856e+06))
					}
					*(*float64)(unsafe.Add(mBase, uint32(l3))) = v207
					v215 = int32(1)
				}
			}
			if v215 != 0 {
				v350 = int32(0)
			} else {
				v350 = v58
			}
		case 2:
			v218 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v219 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
			v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			if int32(1) <= v221 {
				v234 = v221 + int32(-1)
				v238 = *(*float64)(unsafe.Add(mBase, uint32(v220+v234<<(uint(int32(4))%32))+8))
				v239 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				v240 = int32(0)
				__phi248 = v234
				__phi249 = v238
				__phi251 = v240
				__phi252 = v240
				v248 = __phi248
				v249 = __phi249
				v251 = __phi251
				v252 = __phi252
				for {
					v259 = v220 + v251<<(uint(int32(4))%32)
					v260 = *(*float64)(unsafe.Add(mBase, uint32(v259)+8))
					if base.F64_gt(v249, v239) == base.F64_gt(v260, v239) {
						v280 = v252
					} else {
						v263 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
						v264 = *(*float64)(unsafe.Add(mBase, uint32(v259)))
						v269 = *(*float64)(unsafe.Add(mBase, uint32(v220+v248<<(uint(int32(4))%32))))
						if base.F64_lt(v263, base.F64_add(v264, base.F64_div(base.F64_mul(base.F64_sub(v239, v260), base.F64_sub(v269, v264)), base.F64_sub(v249, v260)))) == int32(0) {
							v280 = v252
						} else {
							v280 = base.B2i32(v252 == int32(0))
						}
					}
					v283 = v251 + int32(1)
					if v283 != v221 {
						__phi248 = v251
						__phi249 = v260
						__phi251 = v283
						__phi252 = v280
						v248 = __phi248
						v249 = __phi249
						v251 = __phi251
						v252 = __phi252
						continue
					} else {
						break
					}
					break
				}
				if v280 != 0 {
					v286 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
					v287 = float64(0.017453292519943295)
					v294 = F_sin(m, base.F64_mul(base.F64_sub(base.F64_mul(v286, v287), base.F64_mul(v218, v287)), float64(0.5)))
					mBase = m.M
					v295 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
					if base.F64_le(base.F64_abs(v294), float64(1e-15)) == int32(0) {
						v310 = float64(0.017453292519943295)
						v311 = base.F64_mul(v219, v310)
						v312 = F_cos(m, v311)
						mBase = m.M
						v314 = base.F64_mul(v295, v310)
						v315 = F_cos(m, v314)
						mBase = m.M
						v319 = F_sin(m, base.F64_mul(base.F64_sub(v314, v311), float64(0.5)))
						mBase = m.M
						v326 = F_asin(m, base.F64_sqrt(base.F64_add(base.F64_mul(v319, v319), base.F64_mul(v294, base.F64_mul(base.F64_mul(v312, v315), v294)))))
						mBase = m.M
						*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_mul(v326, float64(1.2745595121712e+07))
						v344 = v280
					} else {
						v301 = float64(0.017453292519943295)
						*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_mul(base.F64_abs(base.F64_sub(base.F64_mul(v295, v301), base.F64_mul(v219, v301))), float64(6.372797560856e+06))
						v344 = v280
					}
				} else {
					v344 = int32(0)
				}
			} else {
				v344 = int32(0)
			}
			if v344 == int32(0) {
				v350 = v58
			} else {
				v350 = int32(0)
			}
		default:
			v350 = int32(0)
		}
	} else {
		v350 = int32(-1)
	}
	m.G0 = v10 + int32(32)
	return v350
}
func F_georadiusCommand(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = int32(1)
	F_georadiusGeneric(m, l0, v2, v2)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_georadiusGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
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
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v181 int32
	_ = v181
	if l2 < int32(6) {
		v123 = int32(-1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v127 != 0 {
		v132 = v127
		goto L41
	} else {
		goto L42
	}
L2:
	;
	v18 = int32(-1)
	v19 = int32(5)
	goto L3
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1+v19<<(uint(int32(2))%32))))
	v26 = F_objectGetVal(m, v25)
	mBase = m.M
	v27 = int32(_a587)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v123 = v112
	goto L1
L5:
	;
	v117 = v113 + int32(1)
	if v117 < l2 {
		v18 = v112
		v19 = v117
		goto L3
	} else {
		goto L39
	}
L6:
	;
	v108 = v19 + int32(1)
	v109 = base.B2i32(v108 < l2)
	if v108 < l2 {
		goto L33
	} else {
		goto L34
	}
L7:
	;
	if v62-v64 == int32(0) {
		goto L6
	} else {
		goto L19
	}
L8:
	;
	v62 = F_tolower(m, v58)
	mBase = m.M
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v64 = F_tolower(m, v63)
	mBase = m.M
	goto L7
L9:
	;
	v32 = v26
	v33 = v27
	v34 = v30
	goto L12
L10:
	;
	v58 = int32(0)
	v59 = v27
	goto L8
L11:
	;
	v58 = v55 & int32(255)
	v59 = v54
	goto L8
L12:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v36 == int32(0) {
		v54 = v33
		v55 = v34
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v54 = v48
	v55 = int32(0)
	goto L11
L14:
	;
	v40 = v34 & int32(255)
	if v40 == v36 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v47 = int32(1)
	v48 = v33 + v47
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if v49 != 0 {
		v32 = v32 + v47
		v33 = v48
		v34 = v49
		goto L12
	} else {
		goto L18
	}
L16:
	;
	v42 = F_tolower(m, v40)
	mBase = m.M
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v44 = F_tolower(m, v43)
	mBase = m.M
	if v42 == v44 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v54 = v33
	v55 = v46
	goto L11
L18:
	;
	goto L13
L19:
	;
	v68 = int32(_a599)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v71 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v103-v105 != 0 {
		v112 = v18
		v113 = v19
		goto L5
	} else {
		goto L32
	}
L21:
	;
	v103 = F_tolower(m, v99)
	mBase = m.M
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	v105 = F_tolower(m, v104)
	mBase = m.M
	goto L20
L22:
	;
	v73 = v26
	v74 = v68
	v75 = v71
	goto L25
L23:
	;
	v99 = int32(0)
	v100 = v68
	goto L21
L24:
	;
	v99 = v96 & int32(255)
	v100 = v95
	goto L21
L25:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v77 == int32(0) {
		v95 = v74
		v96 = v75
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v95 = v89
	v96 = int32(0)
	goto L24
L27:
	;
	v81 = v75 & int32(255)
	if v81 == v77 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v88 = int32(1)
	v89 = v74 + v88
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v90 != 0 {
		v73 = v73 + v88
		v74 = v89
		v75 = v90
		goto L25
	} else {
		goto L31
	}
L29:
	;
	v83 = F_tolower(m, v81)
	mBase = m.M
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	v85 = F_tolower(m, v84)
	mBase = m.M
	if v83 == v85 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v95 = v74
	v96 = v87
	goto L24
L31:
	;
	goto L26
L32:
	;
	goto L6
L33:
	;
	v110 = v108
	goto L35
L34:
	;
	v110 = v19
	goto L35
L35:
	;
	if v108 < l2 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v111 = v108
	goto L38
L37:
	;
	v111 = v18
	goto L38
L38:
	;
	v112 = v111
	v113 = v110
	goto L5
L39:
	;
	goto L4
L40:
	;
	F__serverAssert(m, int32(_a585), int32(_a550), int32(2292))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L52
	} else {
		goto L61
	}
L41:
	;
	if v123 == int32(-1) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v128 != 0 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v130 = l3 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v130
	v132 = v130
	goto L41
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v166))) = int64(73014444033)
	if v123 == int32(-1) {
		goto L59
	} else {
		goto L60
	}
L45:
	;
	v137 = int32(1)
	goto L47
L46:
	;
	v137 = int32(2)
	goto L47
L47:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v137 <= v138 {
		v166 = v132
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v141 = v137 << (uint(int32(3)) % 32)
	v143 = l3 + int32(12)
	if v132 == v143 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v137
	v166 = v163
	goto L44
L50:
	;
	v150 = F_valkey_malloc(m, v141)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L52
	} else {
		goto L54
	}
L51:
	;
	v145 = F_valkey_realloc(m, v132, v141)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	return int32(0)
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v145
	v163 = v145
	goto L49
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v150
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v153 == int32(0) {
		v163 = v150
		goto L49
	} else {
		goto L55
	}
L55:
	;
	v157 = v153 << (uint(int32(3)) % 32)
	if v157 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v163 = v150
	goto L49
L57:
	;
	goto L56
L58:
	;
	v160 = F__emscripten_memcpy_bulkmem(m, v150, v143, v157)
	mBase = m.M
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v137
	return v137
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166)+12)) = int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v166)+8)) = v123
	goto L59
L61:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_georadiusbymemberroCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_georadiusGeneric(m, l0, int32(1), int32(6))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_getAbsolutePath(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
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
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	v18 = m.G0
	v20 = v18 - int32(1024)
	m.G0 = v20
	v22 = F_sdsnew(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v20 + int32(1024)
	return v526
L2:
	;
	v121 = F_getcwd(m, v20, int32(1024))
	mBase = m.M
	if v121 != 0 {
		goto L34
	} else {
		goto L35
	}
L3:
	;
	return int32(0)
L4:
	;
	v26 = int32(_a1759)
	v32 = v22 + int32(-1)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	switch v33 & int32(7) {
	case 0:
		goto L11
	case 1:
		goto L10
	case 2:
		goto L9
	case 3:
		goto L8
	case 4:
		goto L7
	default:
		v50 = int32(0)
		goto L6
	}
L5:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v117 != int32(47) {
		goto L2
	} else {
		goto L32
	}
L6:
	;
	v53 = v22 + v50 + int32(-1)
	if base.Ui32(v53) < base.Ui32(v22) {
		v71 = v22
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(-17))))
	v50 = v49
	goto L6
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(-9))))
	v50 = v46
	goto L6
L9:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(-5)))))
	v50 = v43
	goto L6
L10:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-3)))))
	v50 = v40
	goto L6
L11:
	;
	v50 = int32(base.Ui32(v33) >> (uint(int32(3)) % 32))
	goto L6
L12:
	;
	if base.Ui32(v53) <= base.Ui32(v71) {
		v87 = v53
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v59 = v22
	goto L14
L14:
	;
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v59))))
	v61 = F_strchr(m, v26, v60)
	mBase = m.M
	if v61 == int32(0) {
		v71 = v59
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v71 = v65
	goto L12
L16:
	;
	v65 = v59 + int32(1)
	if base.Ui32(v65) <= base.Ui32(v53) {
		v59 = v65
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v92 = v87 - v71 + int32(1)
	if v22 == v71 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v75 = v53
	goto L20
L20:
	;
	v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v75))))
	v79 = F_strchr(m, v26, v78)
	mBase = m.M
	if v79 == int32(0) {
		v87 = v75
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v87 = v71
	goto L18
L22:
	;
	v83 = v75 + int32(-1)
	if base.Ui32(v71) < base.Ui32(v83) {
		v75 = v83
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v96 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22+v92))) = uint8(v96)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	switch v98 & int32(7) {
	case 0:
		goto L31
	case 1:
		goto L30
	case 2:
		goto L29
	case 3:
		goto L28
	case 4:
		goto L27
	default:
		goto L26
	}
L25:
	;
	v94 = F_memmove(m, v22, v71, v92)
	mBase = m.M
	goto L24
L26:
	;
	goto L5
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(-17)))) = base.I64_extend_i32_u(v92)
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(-9)))) = v92
	goto L5
L29:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(-5)))) = uint16(v92)
	goto L5
L30:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-3)))) = uint8(v92)
	goto L5
L31:
	;
	v102 = v92 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v102)
	goto L5
L32:
	;
	v526 = v22
	goto L1
L33:
	;
	F_sdsfree(m, v22)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L3
	} else {
		goto L153
	}
L34:
	;
	v123 = F_sdsnew(m, v20)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L43
	}
L35:
	;
	v507 = int32(0)
	goto L33
L36:
	;
	v173 = int32(-3)
	v174 = v172 + v173
	v175 = int32(-5)
	v176 = v172 + v175
	v177 = int32(-9)
	v178 = v172 + v177
	v179 = int32(-17)
	v180 = v172 + v179
	v181 = int32(-1)
	goto L54
L37:
	;
	if v144 == int32(0) {
		v172 = v123
		goto L36
	} else {
		goto L44
	}
L38:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v123+int32(-17))))
	v144 = v143
	goto L37
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v123+int32(-9))))
	v144 = v140
	goto L37
L40:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123+int32(-5)))))
	v144 = v137
	goto L37
L41:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+int32(-3)))))
	v144 = v134
	goto L37
L42:
	;
	v144 = int32(base.Ui32(v127) >> (uint(int32(3)) % 32))
	goto L37
L43:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+int32(-1)))))
	v129 = v127 & int32(7)
	switch v129 {
	case 0:
		goto L42
	case 1:
		goto L41
	case 2:
		goto L40
	case 3:
		goto L39
	case 4:
		goto L38
	default:
		v172 = v123
		goto L36
	}
L44:
	;
	switch v129 {
	default:
		goto L50
	case 1:
		goto L49
	case 2:
		goto L48
	case 3:
		goto L47
	case 4:
		goto L46
	}
L45:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+v161+int32(-1)))))
	if v165 == int32(47) {
		v172 = v123
		goto L36
	} else {
		goto L51
	}
L46:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v123+int32(-17))))
	v161 = v160
	goto L45
L47:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v123+int32(-9))))
	v161 = v157
	goto L45
L48:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123+int32(-5)))))
	v161 = v154
	goto L45
L49:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+int32(-3)))))
	v161 = v151
	goto L45
L50:
	;
	v161 = int32(base.Ui32(v127) >> (uint(int32(3)) % 32))
	goto L45
L51:
	;
	v169 = F_sdscat(m, v123, int32(_a1760))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v172 = v169
	goto L36
L53:
	;
	v505 = F_sdscatsds(m, v172, v22)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L3
	} else {
		goto L152
	}
L54:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v181))))
	switch v212 & int32(7) {
	case 0:
		goto L61
	case 1:
		goto L60
	case 2:
		goto L59
	case 3:
		goto L58
	case 4:
		goto L57
	default:
		goto L53
	}
L56:
	;
	if base.Ui32(v221) < base.Ui32(int32(3)) {
		goto L53
	} else {
		goto L62
	}
L57:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v22+v179)))
	v221 = v220
	goto L56
L58:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v22+v177)))
	v221 = v219
	goto L56
L59:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+v175))))
	v221 = v218
	goto L56
L60:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v173))))
	v221 = v217
	goto L56
L61:
	;
	v221 = int32(base.Ui32(v212) >> (uint(int32(3)) % 32))
	goto L56
L62:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v224 != int32(46) {
		goto L53
	} else {
		goto L63
	}
L63:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v227 != int32(46) {
		goto L53
	} else {
		goto L64
	}
L64:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
	if v230 != int32(47) {
		goto L53
	} else {
		goto L65
	}
L65:
	;
	v234 = int32(-1)
	v242 = v22 + v234
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	v245 = v243 & int32(7)
	switch v245 {
	case 0:
		goto L73
	case 1:
		goto L72
	case 2:
		goto L71
	case 3:
		goto L70
	case 4:
		goto L69
	default:
		goto L67
	}
L66:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172+v181))))
	v337 = v335 & int32(7)
	switch v337 {
	case 0:
		goto L105
	case 1:
		goto L104
	case 2:
		goto L103
	case 3:
		goto L102
	case 4:
		goto L101
	default:
		goto L54
	}
L67:
	;
	goto L66
L68:
	;
	if v260 == int32(0) {
		goto L67
	} else {
		goto L74
	}
L69:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(-17))))
	v260 = v259
	goto L68
L70:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(-9))))
	v260 = v256
	goto L68
L71:
	;
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(-5)))))
	v260 = v253
	goto L68
L72:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-3)))))
	v260 = v250
	goto L68
L73:
	;
	v260 = int32(base.Ui32(v243) >> (uint(int32(3)) % 32))
	goto L68
L74:
	;
	v266 = int32(-1)&v260 + v234
	v270 = int32(0)&v260 + int32(3)
	v273 = v266 - v270 + int32(1)
	switch v245 {
	default:
		goto L80
	case 1:
		goto L79
	case 2:
		goto L78
	case 3:
		goto L77
	case 4:
		goto L76
	}
L75:
	;
	v289 = int32(0)
	v291 = base.B2i32(base.Ui32(v270) < base.Ui32(v288))
	if base.Ui32(v270) < base.Ui32(v288) {
		goto L82
	} else {
		goto L83
	}
L76:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(-17))))
	v288 = v287
	goto L75
L77:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(-9))))
	v288 = v284
	goto L75
L78:
	;
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(-5)))))
	v288 = v281
	goto L75
L79:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-3)))))
	v288 = v278
	goto L75
L80:
	;
	v288 = int32(base.Ui32(v243) >> (uint(int32(3)) % 32))
	goto L75
L81:
	;
	v305 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22+v299))) = uint8(v305)
	switch v245 {
	default:
		goto L99
	case 1:
		goto L98
	case 2:
		goto L97
	case 3:
		goto L96
	case 4:
		goto L95
	}
L82:
	;
	v292 = v270
	goto L84
L83:
	;
	v292 = v289
	goto L84
L84:
	;
	v293 = v288 - v292
	if base.Ui32(v273) < base.Ui32(v293) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v295 = v273
	goto L87
L86:
	;
	v295 = v293
	goto L87
L87:
	;
	if v266 < v270 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v297 = v289
	goto L90
L89:
	;
	v297 = v295
	goto L90
L90:
	;
	if base.Ui32(v270) < base.Ui32(v288) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v299 = v297
	goto L93
L92:
	;
	v299 = int32(0)
	goto L93
L93:
	;
	if v299 == int32(0) {
		goto L81
	} else {
		goto L94
	}
L94:
	;
	v303 = F_memmove(m, v22, v22+v292, v299)
	mBase = m.M
	goto L81
L95:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(-17)))) = base.I64_extend_i32_u(v299)
	goto L67
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(-9)))) = v299
	goto L66
L97:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(-5)))) = uint16(v299)
	goto L66
L98:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-3)))) = uint8(v299)
	goto L66
L99:
	;
	v308 = v299 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v242))) = uint8(v308)
	goto L66
L100:
	;
	if base.Ui32(v344) < base.Ui32(int32(2)) {
		goto L54
	} else {
		goto L106
	}
L101:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v344 = v343
	goto L100
L102:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v344 = v342
	goto L100
L103:
	;
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176))))
	v344 = v341
	goto L100
L104:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v344 = v340
	goto L100
L105:
	;
	v344 = int32(base.Ui32(v335) >> (uint(int32(3)) % 32))
	goto L100
L106:
	;
	switch v337 {
	default:
		goto L112
	case 1:
		goto L111
	case 2:
		goto L110
	case 3:
		goto L109
	case 4:
		goto L108
	}
L107:
	;
	v356 = v172 + int32(-2) + v353
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356))))
	if v357 == int32(47) {
		v390 = int32(-2)
		goto L113
	} else {
		goto L114
	}
L108:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v353 = v352
	goto L107
L109:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v353 = v351
	goto L107
L110:
	;
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176))))
	v353 = v350
	goto L107
L111:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v353 = v349
	goto L107
L112:
	;
	v353 = int32(base.Ui32(v335) >> (uint(int32(3)) % 32))
	goto L107
L113:
	;
	v411 = v172 + int32(-1)
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	v414 = v412 & int32(7)
	switch v414 {
	case 0:
		goto L125
	case 1:
		goto L124
	case 2:
		goto L123
	case 3:
		goto L122
	case 4:
		goto L121
	default:
		goto L119
	}
L114:
	;
	v360 = int32(1)
	v365 = v356
	goto L115
L115:
	;
	v380 = v365 + int32(-1)
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	if v381 != int32(47) {
		v360 = v360 + int32(1)
		v365 = v380
		goto L115
	} else {
		goto L117
	}
L116:
	;
	v390 = int32(-2) - v360
	goto L113
L117:
	;
	goto L116
L118:
	;
	goto L54
L119:
	;
	goto L118
L120:
	;
	if v429 == int32(0) {
		goto L119
	} else {
		goto L126
	}
L121:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v172+int32(-17))))
	v429 = v428
	goto L120
L122:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v172+int32(-9))))
	v429 = v425
	goto L120
L123:
	;
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172+int32(-5)))))
	v429 = v422
	goto L120
L124:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172+int32(-3)))))
	v429 = v419
	goto L120
L125:
	;
	v429 = int32(base.Ui32(v412) >> (uint(int32(3)) % 32))
	goto L120
L126:
	;
	v435 = v390>>(uint(int32(31))%32)&v429 + v390
	v439 = int32(0)&v429 + int32(0)
	v442 = v435 - v439 + int32(1)
	switch v414 {
	default:
		goto L132
	case 1:
		goto L131
	case 2:
		goto L130
	case 3:
		goto L129
	case 4:
		goto L128
	}
L127:
	;
	v458 = int32(0)
	v460 = base.B2i32(base.Ui32(v439) < base.Ui32(v457))
	if base.Ui32(v439) < base.Ui32(v457) {
		goto L134
	} else {
		goto L135
	}
L128:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v172+int32(-17))))
	v457 = v456
	goto L127
L129:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v172+int32(-9))))
	v457 = v453
	goto L127
L130:
	;
	v450 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172+int32(-5)))))
	v457 = v450
	goto L127
L131:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172+int32(-3)))))
	v457 = v447
	goto L127
L132:
	;
	v457 = int32(base.Ui32(v412) >> (uint(int32(3)) % 32))
	goto L127
L133:
	;
	v474 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v172+v468))) = uint8(v474)
	switch v414 {
	default:
		goto L151
	case 1:
		goto L150
	case 2:
		goto L149
	case 3:
		goto L148
	case 4:
		goto L147
	}
L134:
	;
	v461 = v439
	goto L136
L135:
	;
	v461 = v458
	goto L136
L136:
	;
	v462 = v457 - v461
	if base.Ui32(v442) < base.Ui32(v462) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v464 = v442
	goto L139
L138:
	;
	v464 = v462
	goto L139
L139:
	;
	if v435 < v439 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v466 = v458
	goto L142
L141:
	;
	v466 = v464
	goto L142
L142:
	;
	if base.Ui32(v439) < base.Ui32(v457) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v468 = v466
	goto L145
L144:
	;
	v468 = int32(0)
	goto L145
L145:
	;
	if v468 == int32(0) {
		goto L133
	} else {
		goto L146
	}
L146:
	;
	v472 = F_memmove(m, v172, v172+v461, v468)
	mBase = m.M
	goto L133
L147:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v172+int32(-17)))) = base.I64_extend_i32_u(v468)
	goto L119
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172+int32(-9)))) = v468
	goto L118
L149:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v172+int32(-5)))) = uint16(v468)
	goto L118
L150:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v172+int32(-3)))) = uint8(v468)
	goto L118
L151:
	;
	v477 = v468 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v411))) = uint8(v477)
	goto L118
L152:
	;
	v507 = v505
	goto L33
L153:
	;
	v526 = v507
	goto L1
}
func F_getArgvReprString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v7 = F_getDecodedObject(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_sdsempty(m)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_objectGetVal(m, v7)
			mBase = m.M
			v14 = F_objectGetVal(m, v7)
			mBase = m.M
			v15 = int32(-1)
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v15))))
			switch v17&int32(7) + v15 {
			case 0:
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-3)))))
				v34 = v24
				if base.Ui32(v34) <= base.Ui32(int32(127)) {
					v40 = F_objectGetVal(m, v7)
					mBase = m.M
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-1)))))
					switch v43 & int32(7) {
					case 0:
						v60 = int32(base.Ui32(v43) >> (uint(int32(3)) % 32))
					case 1:
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-3)))))
						v60 = v50
					case 2:
						v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+int32(-5)))))
						v60 = v53
					case 3:
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-9))))
						v60 = v56
					case 4:
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-17))))
						v60 = v59
					default:
						v60 = int32(0)
					}
				} else {
					v60 = int32(128)
				}
			case 1:
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(-5)))))
				v34 = v27
				if base.Ui32(v34) <= base.Ui32(int32(127)) {
					v40 = F_objectGetVal(m, v7)
					mBase = m.M
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-1)))))
					switch v43 & int32(7) {
					case 0:
						v60 = int32(base.Ui32(v43) >> (uint(int32(3)) % 32))
					case 1:
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-3)))))
						v60 = v50
					case 2:
						v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+int32(-5)))))
						v60 = v53
					case 3:
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-9))))
						v60 = v56
					case 4:
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-17))))
						v60 = v59
					default:
						v60 = int32(0)
					}
				} else {
					v60 = int32(128)
				}
			case 2:
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-9))))
				v34 = v30
				if base.Ui32(v34) <= base.Ui32(int32(127)) {
					v40 = F_objectGetVal(m, v7)
					mBase = m.M
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-1)))))
					switch v43 & int32(7) {
					case 0:
						v60 = int32(base.Ui32(v43) >> (uint(int32(3)) % 32))
					case 1:
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-3)))))
						v60 = v50
					case 2:
						v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+int32(-5)))))
						v60 = v53
					case 3:
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-9))))
						v60 = v56
					case 4:
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-17))))
						v60 = v59
					default:
						v60 = int32(0)
					}
				} else {
					v60 = int32(128)
				}
			case 3:
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-17))))
				v34 = v33
				if base.Ui32(v34) <= base.Ui32(int32(127)) {
					v40 = F_objectGetVal(m, v7)
					mBase = m.M
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-1)))))
					switch v43 & int32(7) {
					case 0:
						v60 = int32(base.Ui32(v43) >> (uint(int32(3)) % 32))
					case 1:
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-3)))))
						v60 = v50
					case 2:
						v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+int32(-5)))))
						v60 = v53
					case 3:
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-9))))
						v60 = v56
					case 4:
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-17))))
						v60 = v59
					default:
						v60 = int32(0)
					}
				} else {
					v60 = int32(128)
				}
			default:
				v40 = F_objectGetVal(m, v7)
				mBase = m.M
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-1)))))
				switch v43 & int32(7) {
				case 0:
					v60 = int32(base.Ui32(v43) >> (uint(int32(3)) % 32))
				case 1:
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-3)))))
					v60 = v50
				case 2:
					v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+int32(-5)))))
					v60 = v53
				case 3:
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-9))))
					v60 = v56
				case 4:
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-17))))
					v60 = v59
				default:
					v60 = int32(0)
				}
			}
			v63 = F_sdscatrepr(m, v11, v13, v60)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				F_decrRefCount(m, v7)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					return v63
				}
			}
		}
	}
}
func F_getDecodedObject(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int64
	_ = v101
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int32(base.Ui32(v10)>>(uint(int32(4))%32)) & int32(15)
	switch v14 {
	case 0, 8:
		goto L3
	default:
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(48)
	return v241
L2:
	;
	if v10&int32(15) != 0 {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(int32(-17)) < base.Ui32(v15) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if base.Ui32(int32(-9)) < base.Ui32(v15) {
		v241 = l0
		goto L1
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15 + int32(8)
	v241 = l0
	goto L1
L6:
	;
	F__serverPanic_1(m, int32(_a1051), int32(622), int32(_a1058), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L9:
	;
	v229 = F_sdsnewlen(m, v8, v197)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L62
	}
L10:
	;
	F__serverPanic_1(m, int32(_a1051), int32(942), int32(_a1059), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L7
	} else {
		goto L61
	}
L11:
	;
	if v14 != int32(1) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v36&int32(4) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v101 = base.I64_extend_i32_s(v96)
	if v101 <= int64(-1) {
		goto L34
	} else {
		goto L35
	}
L14:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v96 = v95
	goto L13
L15:
	;
	if v36&int32(1) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v45 = int32(16)
	goto L18
L17:
	;
	v45 = int32(8)
	goto L18
L18:
	;
	v46 = l0 + v45
	if v36&int32(2) == int32(0) {
		v77 = v46
		goto L19
	} else {
		goto L20
	}
L19:
	;
	goto L29
L20:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	v52 = v46 + v51
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	switch v56 & int32(7) {
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
		v73 = int32(0)
		goto L21
	}
L21:
	;
	v77 = v52 + int32(1) + v73 + int32(1)
	goto L19
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-16))))
	v73 = v72
	goto L21
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-8))))
	v73 = v69
	goto L21
L24:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(-4)))))
	v73 = v66
	goto L21
L25:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-2)))))
	v73 = v63
	goto L21
L26:
	;
	v73 = int32(base.Ui32(v56) >> (uint(int32(3)) % 32))
	goto L21
L27:
	;
	v96 = v77 + v92
	goto L13
L28:
	;
	goto L27
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	goto L28
L30:
	;
	if v8&int32(3) == int32(0) {
		v164 = v8
		goto L41
	} else {
		goto L42
	}
L31:
	;
	goto L30
L33:
	;
	v123 = F_ull2string(m, v119, v120, v121)
	mBase = m.M
	if v123 == int32(0) {
		goto L31
	} else {
		goto L37
	}
L34:
	;
	goto L36
L35:
	;
	v119 = v8
	v120 = int32(32)
	v121 = v101
	goto L33
L36:
	;
	v110 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v110)
	v119 = v8 + int32(1)
	v120 = int32(31)
	v121 = int64(0) - v101
	goto L33
L37:
	;
	goto L30
L39:
	;
	if base.Ui32(int32(255)) < base.Ui32(v197) {
		goto L9
	} else {
		goto L55
	}
L40:
	;
	v197 = v189 - v8
	goto L39
L41:
	;
	v168 = v164
	goto L49
L42:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v150 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v153 = v8
	goto L45
L44:
	;
	v197 = v8 - v8
	goto L39
L45:
	;
	v157 = v153 + int32(1)
	if v157&int32(3) == int32(0) {
		v164 = v157
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v162 != 0 {
		v153 = v157
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v189 = v157
	goto L40
L49:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v177 = int32(-2139062144)
	if (int32(16843008)-v174|v174)&v177 == v177 {
		v168 = v168 + int32(4)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v183 = v168
	goto L52
L51:
	;
	goto L50
L52:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v187 != 0 {
		v183 = v183 + int32(1)
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v189 = v183
	goto L40
L54:
	;
	goto L53
L55:
	;
	goto L58
L56:
	;
	if base.Ui32(int32(128)) < base.Ui32(v197+v211+int32(9)) {
		goto L9
	} else {
		goto L59
	}
L57:
	;
	goto L56
L58:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	goto L57
L59:
	;
	v220 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v8, v197, int32(0), int64(-1))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v241 = v220
	goto L1
L61:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(0)
	v236 = F_zmalloc_usable(m, int32(12), v8+int32(44))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236)+8)) = v229
	*(*int64)(unsafe.Add(mBase, uint32(v236))) = int64(34359738368)
	v241 = v236
	goto L1
}
func F_getF(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v5 == int32(0) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
		if int32(-1) < v20 {
			v24 = F___lockfile(m, v17)
			mBase = m.M
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			if v24 == int32(0) {
				v29 = v25
			} else {
				F___unlockfile(m, v17)
				mBase = m.M
				v29 = v25
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v29 = v23
		}
		if int32(base.Ui32(v29)>>(uint(int32(4))%32))&int32(1) != 0 {
			v48 = int32(0)
			return v48
		} else {
			v36 = l1 + int32(8)
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v40 = F_fread(m, v36, int32(1), int32(1024), v39)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v40
				if v40 != 0 {
					v46 = v36
				} else {
					v46 = int32(0)
				}
				v48 = v46
				return v48
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
		v12 = m.G3
		return v12 + int32(_a247)
	}
}
func F_getFailoverStateString(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v5 = *(*int32)(unsafe.Add(mBase, _consts[658]))
	if base.Ui32(int32(2)) < base.Ui32(v5) {
		v13 = int32(_a288)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5<<(uint(int32(2))%32))+uint32(_consts[718])))
		v13 = v12
	}
	return v13
}
func F_getIntFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v16 = F_getRangeLongFromObjectOrReply(m, l0, l1, int32(-2147483648), int32(2147483647), v9+int32(12), l3)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 != 0 {
			v23 = int32(-1)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v20
			v23 = int32(0)
		}
		m.G0 = v9 + int32(16)
		return v23
	}
}
func F_getListensInfoString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = l0
	v14 = int32(0)
	goto L1
L1:
	;
	v20 = v14 * int32(88)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[325])))
	if v22 == int32(0) {
		v91 = v12
		goto L3
	} else {
		goto L4
	}
L2:
	;
	m.G0 = v9 + int32(48)
	return v91
L3:
	;
	v98 = v14 + int32(1)
	if v98 != int32(4) {
		v12 = v91
		v14 = v98
		goto L1
	} else {
		goto L20
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v29 = m.T0[v28].(func(*base.Module) int32)(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v14
	v46 = F_sdscatfmt(m, v12, int32(_a544), v9+int32(32))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L9
	}
L6:
	;
	return int32(0)
L7:
	;
	if base.Ui32(int32(3)) < base.Ui32(v29) {
		v40 = int32(_a545)
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29<<(uint(int32(2))%32))+uint32(_consts[320])))
	v40 = v39
	goto L5
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[326])))
	if v48 < int32(1) {
		v77 = v46
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[327])))
	if v79 == int32(0) {
		v86 = v77
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v55 = int32(0)
	v56 = v46
	goto L12
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[328])))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v55<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v62
	v67 = F_sdscatfmt(m, v56, int32(_a546), v9+int32(16))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L14
	}
L13:
	;
	v77 = v67
	goto L10
L14:
	;
	v70 = v55 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[326])))
	if v70 < v71 {
		v55 = v70
		v56 = v67
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v89 = F_sdscatfmt(m, v86, int32(_a132), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v79
	v84 = F_sdscatfmt(m, v77, int32(_a547), v9)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v86 = v84
	goto L16
L19:
	;
	v91 = v89
	goto L3
L20:
	;
	goto L2
}
func F_getMigratingSlotDest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+44))
	v5 = F_dictFind(m, v4, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			return v11
		} else {
			return int32(0)
		}
	}
}
func F_getMyShardSlotCount(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+88)))
	if v5&int32(2) == int32(0) {
		v13 = v4
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+2160))
		return v14
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)+2172))
		if v10 != 0 {
			v13 = v10
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+2160))
			return v14
		} else {
			return int32(0)
		}
	}
}
func F_getPositiveLongFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v19 int64
	_ = v19
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l3 == int32(0) {
		v30 = F_getLongLongFromObject(m, l1, v9+int32(8))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			if v30 == int32(0) {
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
				if base.Ui64(int64(-4294967297)) < base.Ui64(v37+int64(-2147483648)) {
					*(*uint32)(unsafe.Add(mBase, uint32(l2))) = uint32(v37)
					if v37 <= int64(-1) {
						F_addReplyError(m, l0, int32(_a1063))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v57 = int32(-1)
							m.G0 = v9 + int32(16)
							return v57
						}
					} else {
						v57 = int32(0)
						m.G0 = v9 + int32(16)
						return v57
					}
				} else {
					F_addReplyError(m, l0, int32(_a1063))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v57 = int32(-1)
						m.G0 = v9 + int32(16)
						return v57
					}
				}
			} else {
				F_addReplyError(m, l0, int32(_a1063))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v57 = int32(-1)
					m.G0 = v9 + int32(16)
					return v57
				}
			}
		}
	} else {
		v15 = F_getLongLongFromObject(m, l1, v9+int32(8))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 != 0 {
				F_addReplyError(m, l0, l3)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v57 = int32(-1)
					m.G0 = v9 + int32(16)
					return v57
				}
			} else {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
				if base.Ui64(v19+int64(-2147483648)) <= base.Ui64(int64(-4294967297)) {
					F_addReplyError(m, l0, l3)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v57 = int32(-1)
						m.G0 = v9 + int32(16)
						return v57
					}
				} else {
					*(*uint32)(unsafe.Add(mBase, uint32(l2))) = uint32(v19)
					if v19 <= int64(-1) {
						F_addReplyError(m, l0, l3)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v57 = int32(-1)
							m.G0 = v9 + int32(16)
							return v57
						}
					} else {
						v57 = int32(0)
						m.G0 = v9 + int32(16)
						return v57
					}
				}
			}
		}
	}
}
func F_getPsyncInitialOffset(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, _consts[40]))
	return v2
}
func F_getRandomHexChars(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	F_getRandomBytes(m, l0, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		if l1 == int32(0) {
		} else {
			v11 = int32(1)
			if l1 == v11 {
				v51 = int32(0)
			} else {
				v18 = int32(0)
				v23 = v18
				v25 = v18
				for {
					v26 = l0 + v23
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
					v28 = int32(15)
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27&v28)+uint32(_consts[797]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v32)
					v35 = v26 + int32(1)
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36&v28)+uint32(_consts[797]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v35))) = uint8(v41)
					v43 = int32(2)
					v44 = v23 + v43
					v46 = v25 + v43
					if v46 != l1&int32(-2) {
						v23 = v44
						v25 = v46
						continue
					} else {
						break
					}
					break
				}
				v51 = v44
			}
			if l1&v11 == int32(0) {
			} else {
				v56 = l0 + v51
				v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57&int32(15))+uint32(_consts[797]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v62)
			}
		}
		return
	}
}
func F_getTimeoutFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v134 int64
	_ = v134
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v143 int64
	_ = v143
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v159 int64
	_ = v159
	var v164 int64
	_ = v164
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v179 int64
	_ = v179
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v189 int64
	_ = v189
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v198 int32
	_ = v198
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v212 int64
	_ = v212
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v215 int64
	_ = v215
	var v225 int64
	_ = v225
	var v230 int64
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v261 int64
	_ = v261
	var v266 int64
	_ = v266
	var v267 int64
	_ = v267
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int64
	_ = v278
	var v281 int64
	_ = v281
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v302 int64
	_ = v302
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v14 = *(*int64)(unsafe.Add(mBase, _consts[98]))
	if l3 != 0 {
		v276 = F_getLongLongFromObjectOrReply(m, l0, l1, v11+int32(56), int32(_a1745))
		mBase = m.M
		v277 = m.ExcPending
		if v277 != 0 {
			return int32(0)
		} else {
			if v276 != 0 {
				v306 = int32(-1)
				m.G0 = v11 + int32(64)
				return v306
			} else {
				v278 = *(*int64)(unsafe.Add(mBase, uint32(v11)+56))
				v281 = v278
				if int64(-1) < v281 {
					if base.B2i32(v281 == int64(0)) == int32(0) {
						if base.Ui64(v281) <= base.Ui64(v14^int64(9223372036854775807)) {
							v302 = v281 + v14
							*(*int64)(unsafe.Add(mBase, uint32(l2))) = v302
							v306 = int32(0)
							m.G0 = v11 + int32(64)
							return v306
						} else {
							F_addReplyError(m, l0, int32(_a1746))
							mBase = m.M
							v299 = m.ExcPending
							if v299 != 0 {
								return int32(0)
							} else {
								v306 = int32(-1)
								m.G0 = v11 + int32(64)
								return v306
							}
						}
					} else {
						v302 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l2))) = v302
						v306 = int32(0)
						m.G0 = v11 + int32(64)
						return v306
					}
				} else {
					F_addReplyError(m, l0, int32(_a1747))
					mBase = m.M
					v287 = m.ExcPending
					if v287 != 0 {
						return int32(0)
					} else {
						v306 = int32(-1)
						m.G0 = v11 + int32(64)
						return v306
					}
				}
			}
		}
	} else {
		v15 = int32(-1)
		v19 = F_getLongDoubleFromObjectOrReply(m, l0, l1, v11+int32(40), int32(_a1748))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v19 != 0 {
				v306 = v15
				m.G0 = v11 + int32(64)
				return v306
			} else {
				v25 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
				v29 = v11 + int32(48)
				v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
				v31 = int64(0)
				F___multf3(m, v11+int32(24), v25, v30, v31, int64(4614206099078250496))
				mBase = m.M
				v38 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(32))))
				*(*int64)(unsafe.Add(mBase, uint32(v29))) = v38
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v40
				v42 = int64(-1125899906842624)
				v43 = int64(4629137466983448575)
				v47 = int32(-1)
				v51 = v38 & int64(9223372036854775807)
				v52 = int64(9223090561878065152)
				if v51 == v52 {
					v56 = base.B2i32(v40 != v31)
				} else {
					v56 = base.B2i32(base.Ui64(v52) < base.Ui64(v51))
				}
				if v56 != 0 {
					v95 = v47
					v99 = v95
				} else {
					if base.B2i32(v42|v40|(int64(4629137466983448575)|v51) == int64(0)) == int32(0) {
						if v43&v38 < int64(0) {
							if v38 == v43 {
								v89 = base.B2i32(base.Ui64(v42) < base.Ui64(v40))
							} else {
								v89 = base.B2i32(v43 < v38)
							}
							if v89 != 0 {
								v95 = v47
							} else {
								v95 = base.B2i32(v40^v42|(v38^v43) != int64(0))
							}
							v99 = v95
						} else {
							if v38 == v43 {
								v80 = base.B2i32(base.Ui64(v40) < base.Ui64(v42))
							} else {
								v80 = base.B2i32(v38 < v43)
							}
							if v80 != 0 {
								v95 = v47
								v99 = v95
							} else {
								v99 = base.B2i32(v40^v42|(v38^v43) != int64(0))
							}
						}
					} else {
						v99 = int32(0)
					}
				}
				if v99 < int32(1) {
					v106 = v11 + int32(8)
					v110 = int64(0)
					v112 = m.G0
					v114 = v112 - int32(112)
					m.G0 = v114
					v118 = F___letf2(m, v40, v38, v110, v110)
					mBase = m.M
					if v118 == int32(0) {
						v214 = v40
						v215 = v38
					} else {
						v123 = base.I32_wrap_i64(int64(base.Ui64(v38) >> (uint(int64(48)) % 64)))
						v125 = v123 & int32(32767)
						if base.Ui32(int32(16494)) < base.Ui32(v125) {
							v214 = v40
							v215 = v38
						} else {
							if base.Ui32(int32(16382)) < base.Ui32(v125) {
								v138 = int64(0)
								v139 = int64(4642929740842270720)
								F___addtf3(m, v114+int32(96), v40, v38, v138, v139)
								mBase = m.M
								v143 = *(*int64)(unsafe.Add(mBase, uint32(v114)+96))
								v148 = *(*int64)(unsafe.Add(mBase, uint32(v114+int32(104))))
								v150 = int64(-4580442296012505088)
								F___addtf3(m, v114+int32(80), v143, v148, v138, v150)
								mBase = m.M
								F___addtf3(m, v114+int32(64), v40, v38, v138, v150)
								mBase = m.M
								v159 = *(*int64)(unsafe.Add(mBase, uint32(v114)+64))
								v164 = *(*int64)(unsafe.Add(mBase, uint32(v114+int32(72))))
								F___addtf3(m, v114+int32(48), v159, v164, v138, v139)
								mBase = m.M
								v170 = *(*int64)(unsafe.Add(mBase, uint32(v114)+80))
								v171 = *(*int64)(unsafe.Add(mBase, uint32(v114)+48))
								v173 = base.B2i32(base.Ui32(v123) < base.Ui32(int32(32768)))
								if base.Ui32(v123) < base.Ui32(int32(32768)) {
									v174 = v170
								} else {
									v174 = v171
								}
								v179 = *(*int64)(unsafe.Add(mBase, uint32(v114+int32(88))))
								v184 = *(*int64)(unsafe.Add(mBase, uint32(v114+int32(56))))
								if base.Ui32(v123) < base.Ui32(int32(32768)) {
									v185 = v179
								} else {
									v185 = v184
								}
								F___subtf3(m, v114+int32(32), v174, v185, v40, v38)
								mBase = m.M
								v189 = *(*int64)(unsafe.Add(mBase, uint32(v114)+32))
								v194 = *(*int64)(unsafe.Add(mBase, uint32(v114+int32(40))))
								F___addtf3(m, v114+int32(16), v40, v38, v189, v194)
								mBase = m.M
								v196 = int64(0)
								v198 = F___letf2(m, v189, v194, v196, v196)
								mBase = m.M
								v203 = *(*int64)(unsafe.Add(mBase, uint32(v114+int32(24))))
								v204 = *(*int64)(unsafe.Add(mBase, uint32(v114)+16))
								if int32(-1) < v198 {
									v214 = v204
									v215 = v203
								} else {
									F___addtf3(m, v114, v204, v203, int64(0), int64(4611404543450677248))
									mBase = m.M
									v212 = *(*int64)(unsafe.Add(mBase, uint32(v114+int32(8))))
									v213 = *(*int64)(unsafe.Add(mBase, uint32(v114)))
									v214 = v213
									v215 = v212
								}
							} else {
								if base.Ui32(v123) < base.Ui32(int32(32768)) {
									v134 = int64(4611404543450677248)
								} else {
									v134 = int64(-9223372036854775807 - 1)
								}
								v214 = int64(0)
								v215 = v134
							}
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v106))) = v214
					*(*int64)(unsafe.Add(mBase, uint32(v106)+8)) = v215
					m.G0 = v114 + int32(112)
					v225 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
					v316 = int32(16)
					v230 = *(*int64)(unsafe.Add(mBase, uint32(v11+v316)))
					v234 = m.G0
					v236 = v234 - v316
					m.G0 = v236
					v243 = base.I32_wrap_i64(int64(base.Ui64(v230)>>(uint(int64(48))%64))) & int32(32767)
					if base.Ui32(v243) < base.Ui32(int32(16383)) {
						v267 = int64(0)
					} else {
						if base.Ui32(int32(-65)) < base.Ui32(v243+int32(-16447)) {
							F___lshrti3(m, v236, v225, v230&int64(281474976710655)|int64(281474976710656), int32(16495)-v243)
							mBase = m.M
							v261 = *(*int64)(unsafe.Add(mBase, uint32(v236)))
							if int64(-1) < v230 {
								v266 = v261
							} else {
								v266 = int64(0) - v261
							}
							v267 = v266
						} else {
							v267 = v230>>(uint(int64(63))%64) ^ int64(9223372036854775807)
						}
					}
					m.G0 = v236 + int32(16)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v267
					v281 = v267
					if int64(-1) < v281 {
						if base.B2i32(v281 == int64(0)) == int32(0) {
							if base.Ui64(v281) <= base.Ui64(v14^int64(9223372036854775807)) {
								v302 = v281 + v14
								*(*int64)(unsafe.Add(mBase, uint32(l2))) = v302
								v306 = int32(0)
								m.G0 = v11 + int32(64)
								return v306
							} else {
								F_addReplyError(m, l0, int32(_a1746))
								mBase = m.M
								v299 = m.ExcPending
								if v299 != 0 {
									return int32(0)
								} else {
									v306 = int32(-1)
									m.G0 = v11 + int32(64)
									return v306
								}
							}
						} else {
							v302 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l2))) = v302
							v306 = int32(0)
							m.G0 = v11 + int32(64)
							return v306
						}
					} else {
						F_addReplyError(m, l0, int32(_a1747))
						mBase = m.M
						v287 = m.ExcPending
						if v287 != 0 {
							return int32(0)
						} else {
							v306 = int32(-1)
							m.G0 = v11 + int32(64)
							return v306
						}
					}
				} else {
					F_addReplyError(m, l0, int32(_a1746))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						v306 = v15
						m.G0 = v11 + int32(64)
						return v306
					}
				}
			}
		}
	}
}
func F_getbitCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v18 = F_getBitOffsetFromArgument(m, l0, v13, v10+int32(8), v2, v2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		if v18 != 0 {
			m.G0 = v10 + int32(48)
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
			v23 = *(*int32)(unsafe.Add(mBase, _consts[85]))
			v24 = F_lookupKeyReadOrReply(m, l0, v21, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				if v24 == int32(0) {
					m.G0 = v10 + int32(48)
					return
				} else {
					v29 = F_checkType(m, l0, v24, int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						if v29 != 0 {
							m.G0 = v10 + int32(48)
							return
						} else {
							v31 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
							v34 = base.I32_wrap_i64(int64(base.Ui64(v31) >> (uint(int64(3)) % 64)))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
							switch int32(base.Ui32(v35)>>(uint(int32(4))%32)) & int32(15) {
							case 0, 8:
								v40 = F_objectGetVal(m, v24)
								mBase = m.M
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-1)))))
								switch v43 & int32(7) {
								case 0:
									v60 = int32(base.Ui32(v43) >> (uint(int32(3)) % 32))
									if base.Ui32(v60) <= base.Ui32(v34) {
										v135 = *(*int32)(unsafe.Add(mBase, _consts[85]))
										v136 = v135
									} else {
										v62 = F_objectGetVal(m, v24)
										mBase = m.M
										v114 = v62 + v34
										v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
										if int32(1)<<(uint((base.I32_wrap_i64(v31)^int32(-1))&int32(7))%32)&v127 != 0 {
											v129 = int32(16)
										} else {
											v129 = int32(12)
										}
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+uint32(_consts[84])))
										v136 = v131
									}
								case 1:
									v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-3)))))
									v60 = v50
									if base.Ui32(v60) <= base.Ui32(v34) {
										v135 = *(*int32)(unsafe.Add(mBase, _consts[85]))
										v136 = v135
									} else {
										v62 = F_objectGetVal(m, v24)
										mBase = m.M
										v114 = v62 + v34
										v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
										if int32(1)<<(uint((base.I32_wrap_i64(v31)^int32(-1))&int32(7))%32)&v127 != 0 {
											v129 = int32(16)
										} else {
											v129 = int32(12)
										}
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+uint32(_consts[84])))
										v136 = v131
									}
								case 2:
									v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+int32(-5)))))
									v60 = v53
									if base.Ui32(v60) <= base.Ui32(v34) {
										v135 = *(*int32)(unsafe.Add(mBase, _consts[85]))
										v136 = v135
									} else {
										v62 = F_objectGetVal(m, v24)
										mBase = m.M
										v114 = v62 + v34
										v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
										if int32(1)<<(uint((base.I32_wrap_i64(v31)^int32(-1))&int32(7))%32)&v127 != 0 {
											v129 = int32(16)
										} else {
											v129 = int32(12)
										}
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+uint32(_consts[84])))
										v136 = v131
									}
								case 3:
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-9))))
									v60 = v56
									if base.Ui32(v60) <= base.Ui32(v34) {
										v135 = *(*int32)(unsafe.Add(mBase, _consts[85]))
										v136 = v135
									} else {
										v62 = F_objectGetVal(m, v24)
										mBase = m.M
										v114 = v62 + v34
										v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
										if int32(1)<<(uint((base.I32_wrap_i64(v31)^int32(-1))&int32(7))%32)&v127 != 0 {
											v129 = int32(16)
										} else {
											v129 = int32(12)
										}
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+uint32(_consts[84])))
										v136 = v131
									}
								case 4:
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-17))))
									v60 = v59
									if base.Ui32(v60) <= base.Ui32(v34) {
										v135 = *(*int32)(unsafe.Add(mBase, _consts[85]))
										v136 = v135
									} else {
										v62 = F_objectGetVal(m, v24)
										mBase = m.M
										v114 = v62 + v34
										v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
										if int32(1)<<(uint((base.I32_wrap_i64(v31)^int32(-1))&int32(7))%32)&v127 != 0 {
											v129 = int32(16)
										} else {
											v129 = int32(12)
										}
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+uint32(_consts[84])))
										v136 = v131
									}
								default:
									v135 = *(*int32)(unsafe.Add(mBase, _consts[85]))
									v136 = v135
								}
							default:
								v65 = v10 + int32(16)
								v67 = F_objectGetVal(m, v24)
								mBase = m.M
								v68 = base.I64_extend_i32_s(v67)
								if v68 <= int64(-1) {
									v77 = int32(45)
									*(*uint8)(unsafe.Add(mBase, uint32(v65))) = uint8(v77)
									v86 = v10 + int32(17)
									v87 = int32(31)
									v88 = int64(0) - v68
									v89 = int32(1)
								} else {
									v86 = v65
									v87 = int32(32)
									v88 = v68
									v89 = int32(0)
								}
								v90 = F_ull2string(m, v86, v87, v88)
								mBase = m.M
								if v90 == int32(0) {
									v109 = int32(0)
								} else {
									v109 = v90 + v89
								}
								if base.Ui32(v109) <= base.Ui32(v34) {
									v135 = *(*int32)(unsafe.Add(mBase, _consts[85]))
									v136 = v135
								} else {
									v114 = v10 + int32(16) + v34
									v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
									if int32(1)<<(uint((base.I32_wrap_i64(v31)^int32(-1))&int32(7))%32)&v127 != 0 {
										v129 = int32(16)
									} else {
										v129 = int32(12)
									}
									v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+uint32(_consts[84])))
									v136 = v131
								}
							}
							F_addReply(m, l0, v136)
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return
							} else {
								m.G0 = v10 + int32(48)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_getc(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_do_getc(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_getcwd(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v7 = m.G0
	v8 = int32(4096)
	if l0 != 0 {
		v11 = int32(16)
	} else {
		v11 = v8
	}
	v12 = v7 - v11
	m.G0 = v12
	if l0 == int32(0) {
		v20 = v12
		v21 = v8
		v22 = int32(0)
		v23 = m.Env.X__syscall_getcwd(m, v20, v21)
		mBase = m.M
		if base.Ui32(v23) < base.Ui32(int32(-4095)) {
			v31 = v23
		} else {
			v26 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(0) - v23
			v31 = int32(-1)
		}
		if v31 < int32(0) {
			v52 = v22
		} else {
			if v31 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(44)
				v52 = v22
			} else {
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				if v36 == int32(47) {
					if v20 == v12 {
						v45 = F_strlen(m, v20)
						mBase = m.M
						v47 = v45 + int32(1)
						v48 = F_emscripten_builtin_malloc(m, v47)
						mBase = m.M
						if v48 != 0 {
							v50 = F___memcpy(m, v48, v20, v47)
							mBase = m.M
							v51 = v50
						} else {
							v51 = int32(0)
						}
						v52 = v51
					} else {
						v52 = v20
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(44)
					v52 = v22
				}
			}
		}
	} else {
		if l1 != 0 {
			v20 = l0
			v21 = l1
			v22 = int32(0)
			v23 = m.Env.X__syscall_getcwd(m, v20, v21)
			mBase = m.M
			if base.Ui32(v23) < base.Ui32(int32(-4095)) {
				v31 = v23
			} else {
				v26 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(0) - v23
				v31 = int32(-1)
			}
			if v31 < int32(0) {
				v52 = v22
			} else {
				if v31 == int32(0) {
					*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(44)
					v52 = v22
				} else {
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
					if v36 == int32(47) {
						if v20 == v12 {
							v45 = F_strlen(m, v20)
							mBase = m.M
							v47 = v45 + int32(1)
							v48 = F_emscripten_builtin_malloc(m, v47)
							mBase = m.M
							if v48 != 0 {
								v50 = F___memcpy(m, v48, v20, v47)
								mBase = m.M
								v51 = v50
							} else {
								v51 = int32(0)
							}
							v52 = v51
						} else {
							v52 = v20
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(44)
						v52 = v22
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
			v52 = int32(0)
		}
	}
	m.G0 = v7
	return v52
}
func F_getdelCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v65 int32
	_ = v65
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_initDeferredReplyBuffer(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13<<(uint(int32(2))%32))+uint32(_consts[927])))
		v18 = F_lookupKeyReadOrReply(m, l0, v11, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			if v18 == int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = F_dbSyncDelete(m, v27, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					if v30 == int32(0) {
						F_commitDeferredReplyBuffer(m, l0, int32(1))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							m.G0 = v6 + int32(16)
							return
						}
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
						v37 = *(*int32)(unsafe.Add(mBase, _consts[937]))
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v37
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v35
						F_rewriteClientCommandVector(m, l0, int32(2), v6)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
							F_signalModifiedKey(m, l0, v43, v45)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
								F_notifyKeyspaceEvent(m, int32(4), int32(_a213), v51, v53)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									v56 = int32(_a44)
									v58 = *(*int64)(unsafe.Add(mBase, _consts[83]))
									*(*int64)(unsafe.Add(mBase, _consts[83])) = v58 + int64(1)
									F_commitDeferredReplyBuffer(m, l0, int32(1))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
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
			} else {
				v23 = F_checkType(m, l0, v18, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					if v23 != 0 {
						m.G0 = v6 + int32(16)
						return
					} else {
						F_addReplyBulk(m, l0, v18)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
							v30 = F_dbSyncDelete(m, v27, v29)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								if v30 == int32(0) {
									F_commitDeferredReplyBuffer(m, l0, int32(1))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
									v37 = *(*int32)(unsafe.Add(mBase, _consts[937]))
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v37
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v35
									F_rewriteClientCommandVector(m, l0, int32(2), v6)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
										v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
										F_signalModifiedKey(m, l0, v43, v45)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
											v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
											F_notifyKeyspaceEvent(m, int32(4), int32(_a213), v51, v53)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return
											} else {
												v56 = int32(_a44)
												v58 = *(*int64)(unsafe.Add(mBase, _consts[83]))
												*(*int64)(unsafe.Add(mBase, _consts[83])) = v58 + int64(1)
												F_commitDeferredReplyBuffer(m, l0, int32(1))
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
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
				}
			}
		}
	}
}
func F_getenv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	goto L7
L1:
	;
	v110 = int32(0)
	v111 = v96 - l0
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v111))))
	if v113 != 0 {
		v184 = v110
		goto L25
	} else {
		goto L26
	}
L2:
	;
	if v96 != l0 {
		goto L1
	} else {
		goto L24
	}
L3:
	;
	goto L2
L4:
	;
	v87 = v82
	goto L20
L5:
	;
	v82 = v73
	goto L4
L7:
	;
	if l0&int32(3) == int32(0) {
		v33 = l0
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v42 = int32(-2139062144)
	if (int32(16843008)-v39|v39)&v42 != v42 {
		v73 = v33
		goto L5
	} else {
		goto L15
	}
L9:
	;
	v20 = l0
	goto L10
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 == int32(0) {
		v96 = v20
		goto L3
	} else {
		goto L12
	}
L11:
	;
	v33 = v30
	goto L8
L12:
	;
	if v25 == int32(61) {
		v96 = v20
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v30 = v20 + int32(1)
	if v30&int32(3) != 0 {
		v20 = v30
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v48 = v33
	v51 = v39
	goto L16
L16:
	;
	v54 = v51 ^ int32(1027423549)
	v57 = int32(-2139062144)
	if (int32(16843008)-v54|v54)&v57 != v57 {
		v73 = v48
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v63 = v48 + int32(4)
	v67 = int32(-2139062144)
	if (v61|(int32(16843008)-v61))&v67 == v67 {
		v48 = v63
		v51 = v61
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v82 = v63
	goto L4
L20:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v88 == int32(0) {
		v96 = v87
		goto L3
	} else {
		goto L22
	}
L21:
	;
	v96 = v87
	goto L3
L22:
	;
	if v88 != int32(61) {
		v87 = v87 + int32(1)
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	return int32(0)
L25:
	;
	return v184
L26:
	;
	v114 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, _consts[1021]))
	if v115 == v114 {
		v184 = v110
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v118 == int32(0) {
		v184 = v110
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v122 = v115
	v125 = v118
	goto L30
L29:
	;
	v184 = v172 + int32(1)
	goto L25
L30:
	;
	if v111 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v177 != 0 {
		v122 = v122 + int32(4)
		v125 = v177
		goto L30
	} else {
		goto L48
	}
L33:
	;
	if v170 != 0 {
		goto L32
	} else {
		goto L46
	}
L34:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v129 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v170 = int32(0)
	goto L33
L36:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v170 = v158 - v163
	goto L33
L37:
	;
	v131 = l0
	v132 = v125
	v133 = v111
	v134 = v129
	goto L40
L38:
	;
	v158 = int32(0)
	v159 = v125
	goto L36
L39:
	;
	v158 = v155 & int32(255)
	v159 = v153
	goto L36
L40:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v134&int32(255) != v138 {
		v153 = v132
		v155 = v134
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v153 = v147
	v155 = int32(0)
	goto L39
L42:
	;
	if v138 == int32(0) {
		v153 = v132
		v155 = v134
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v143 = v133 + int32(-1)
	if v143 == int32(0) {
		v153 = v132
		v155 = v134
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v146 = int32(1)
	v147 = v132 + v146
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)))
	if v148 != 0 {
		v131 = v131 + v146
		v132 = v147
		v133 = v143
		v134 = v148
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v172 = v171 + v111
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v173 == int32(61) {
		goto L29
	} else {
		goto L47
	}
L47:
	;
	goto L32
L48:
	;
	v184 = v110
	goto L25
}
func F_getint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8))))
	v11 = v9 + int32(-48)
	if base.Ui32(v11) <= base.Ui32(int32(9)) {
		v17 = int32(0)
		v18 = v8
		v19 = v11
		for {
			if base.Ui32(int32(214748364)) < base.Ui32(v17) {
				v34 = int32(-1)
			} else {
				v27 = v17 * int32(10)
				if base.Ui32(v27^int32(2147483647)) < base.Ui32(v19) {
					v32 = int32(-1)
				} else {
					v32 = v19 + v27
				}
				v34 = v32
			}
			v36 = v18 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v36
			v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18)+1)))
			v40 = v38 + int32(-48)
			if base.Ui32(v40) < base.Ui32(int32(10)) {
				v17 = v34
				v18 = v36
				v19 = v40
				continue
			} else {
				break
			}
			break
		}
		return v34
	} else {
		return int32(0)
	}
}
func F_getpwuid_r(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	return int32(44)
}
func F_getrangeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int64
	_ = v43
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
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
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v137 int64
	_ = v137
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v146 int64
	_ = v146
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v17 = F_getLongLongFromObjectOrReply(m, l0, v13, v10+int32(40), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		if v17 != 0 {
			m.G0 = v10 + int32(48)
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
			v24 = F_getLongLongFromObjectOrReply(m, l0, v20, v10+int32(32), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				if v24 != 0 {
					m.G0 = v10 + int32(48)
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
					v29 = *(*int32)(unsafe.Add(mBase, _consts[939]))
					v30 = F_lookupKeyReadOrReply(m, l0, v27, v29)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						if v30 == int32(0) {
							m.G0 = v10 + int32(48)
							return
						} else {
							v35 = F_checkType(m, l0, v30, int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								if v35 != 0 {
									m.G0 = v10 + int32(48)
									return
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
									v38 = F_objectGetVal(m, v30)
									mBase = m.M
									if v37&int32(240) != int32(16) {
										v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(-1)))))
										switch v91 & int32(7) {
										case 0:
											v110 = int32(base.Ui32(v91) >> (uint(int32(3)) % 32))
										case 1:
											v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(-3)))))
											v110 = v98
										case 2:
											v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38+int32(-5)))))
											v110 = v101
										case 3:
											v104 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(-9))))
											v110 = v104
										case 4:
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(-17))))
											v108 = v107
											v110 = v108
										default:
											v108 = int32(0)
											v110 = v108
										}
										v111 = v38
										v112 = v110
									} else {
										v43 = base.I64_extend_i32_s(v38)
										if v43 <= int64(-1) {
											v53 = int32(45)
											*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v53)
											v57 = int32(1)
											v62 = v10 + v57
											v63 = int32(31)
											v64 = int64(0) - v43
											v65 = v57
										} else {
											v62 = v10
											v63 = int32(32)
											v64 = v43
											v65 = int32(0)
										}
										v66 = F_ull2string(m, v62, v63, v64)
										mBase = m.M
										if v66 == int32(0) {
											v85 = int32(0)
										} else {
											v85 = v66 + v65
										}
										v111 = v10
										v112 = v85
									}
									v114 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
									v115 = *(*int64)(unsafe.Add(mBase, uint32(v10)+40))
									v117 = base.B2i32(int64(-1) < v115)
									if int64(-1) < v115 {
										if int64(-1) < v115 {
											v128 = v115
										} else {
											v126 = v115 + base.I64_extend_i32_u(v112)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v126
											v128 = v126
										}
										if int64(-1) < v114 {
											v134 = v114
										} else {
											v132 = v114 + base.I64_extend_i32_u(v112)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v132
											v134 = v132
										}
										if int64(-1) < v128 {
											v140 = v128
										} else {
											v137 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v137
											v140 = v137
										}
										if int64(-1) < v134 {
											v146 = v134
										} else {
											v143 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v143
											v146 = v143
										}
										if base.Ui64(v146) < base.Ui64(base.I64_extend_i32_u(v112)) {
											v153 = v146
										} else {
											v151 = base.I64_extend_i32_u(v112 + int32(-1))
											*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v151
											v153 = v151
										}
										if base.Ui64(v153) < base.Ui64(v140) {
											v156 = *(*int32)(unsafe.Add(mBase, _consts[939]))
											F_addReply(m, l0, v156)
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return
											} else {
												m.G0 = v10 + int32(48)
												return
											}
										} else {
											if v112 != 0 {
												F_addReplyBulkCBuffer(m, l0, v111+base.I32_wrap_i64(v140), base.I32_wrap_i64(v153-v140)+int32(1))
												mBase = m.M
												v166 = m.ExcPending
												if v166 != 0 {
													return
												} else {
													m.G0 = v10 + int32(48)
													return
												}
											} else {
												v156 = *(*int32)(unsafe.Add(mBase, _consts[939]))
												F_addReply(m, l0, v156)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return
												} else {
													m.G0 = v10 + int32(48)
													return
												}
											}
										}
									} else {
										if int64(-1) < v114 {
											if int64(-1) < v115 {
												v128 = v115
											} else {
												v126 = v115 + base.I64_extend_i32_u(v112)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v126
												v128 = v126
											}
											if int64(-1) < v114 {
												v134 = v114
											} else {
												v132 = v114 + base.I64_extend_i32_u(v112)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v132
												v134 = v132
											}
											if int64(-1) < v128 {
												v140 = v128
											} else {
												v137 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v137
												v140 = v137
											}
											if int64(-1) < v134 {
												v146 = v134
											} else {
												v143 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v143
												v146 = v143
											}
											if base.Ui64(v146) < base.Ui64(base.I64_extend_i32_u(v112)) {
												v153 = v146
											} else {
												v151 = base.I64_extend_i32_u(v112 + int32(-1))
												*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v151
												v153 = v151
											}
											if base.Ui64(v153) < base.Ui64(v140) {
												v156 = *(*int32)(unsafe.Add(mBase, _consts[939]))
												F_addReply(m, l0, v156)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return
												} else {
													m.G0 = v10 + int32(48)
													return
												}
											} else {
												if v112 != 0 {
													F_addReplyBulkCBuffer(m, l0, v111+base.I32_wrap_i64(v140), base.I32_wrap_i64(v153-v140)+int32(1))
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return
													} else {
														m.G0 = v10 + int32(48)
														return
													}
												} else {
													v156 = *(*int32)(unsafe.Add(mBase, _consts[939]))
													F_addReply(m, l0, v156)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return
													} else {
														m.G0 = v10 + int32(48)
														return
													}
												}
											}
										} else {
											if v115 <= v114 {
												if int64(-1) < v115 {
													v128 = v115
												} else {
													v126 = v115 + base.I64_extend_i32_u(v112)
													*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v126
													v128 = v126
												}
												if int64(-1) < v114 {
													v134 = v114
												} else {
													v132 = v114 + base.I64_extend_i32_u(v112)
													*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v132
													v134 = v132
												}
												if int64(-1) < v128 {
													v140 = v128
												} else {
													v137 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v137
													v140 = v137
												}
												if int64(-1) < v134 {
													v146 = v134
												} else {
													v143 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v143
													v146 = v143
												}
												if base.Ui64(v146) < base.Ui64(base.I64_extend_i32_u(v112)) {
													v153 = v146
												} else {
													v151 = base.I64_extend_i32_u(v112 + int32(-1))
													*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v151
													v153 = v151
												}
												if base.Ui64(v153) < base.Ui64(v140) {
													v156 = *(*int32)(unsafe.Add(mBase, _consts[939]))
													F_addReply(m, l0, v156)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return
													} else {
														m.G0 = v10 + int32(48)
														return
													}
												} else {
													if v112 != 0 {
														F_addReplyBulkCBuffer(m, l0, v111+base.I32_wrap_i64(v140), base.I32_wrap_i64(v153-v140)+int32(1))
														mBase = m.M
														v166 = m.ExcPending
														if v166 != 0 {
															return
														} else {
															m.G0 = v10 + int32(48)
															return
														}
													} else {
														v156 = *(*int32)(unsafe.Add(mBase, _consts[939]))
														F_addReply(m, l0, v156)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return
														} else {
															m.G0 = v10 + int32(48)
															return
														}
													}
												}
											} else {
												v122 = *(*int32)(unsafe.Add(mBase, _consts[939]))
												F_addReply(m, l0, v122)
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return
												} else {
													m.G0 = v10 + int32(48)
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
func F_getrlimit(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int64
	_ = v16
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 == v3 {
	} else {
		v16 = int64(-1)
		*(*int64)(unsafe.Add(mBase, uint32(l1))) = v16
		*(*int64)(unsafe.Add(mBase, uint32(l1+int32(8)))) = v16
	}
	m.G0 = v9 + int32(16)
	return v3
}
func F_getsetCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
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
	var v77 int64
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_initDeferredReplyBuffer(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13<<(uint(int32(2))%32))+uint32(_consts[927])))
		v18 = F_lookupKeyReadOrReply(m, l0, v11, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			if v18 == int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v28
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
				if v30&int32(1) == int32(0) {
					v50 = F_tryObjectEncoding(m, v28)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v50
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
						F_setKey(m, l0, v53, v55, v6+int32(12), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
							F_incrRefCount(m, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v65
								v67 = v64
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+28))
								F_notifyKeyspaceEvent(m, int32(8), int32(_a188), v70, v72)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									v75 = int32(_a44)
									v77 = *(*int64)(unsafe.Add(mBase, _consts[83]))
									*(*int64)(unsafe.Add(mBase, _consts[83])) = v77 + int64(1)
									F_commitDeferredReplyBuffer(m, l0, int32(1))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										v86 = *(*int32)(unsafe.Add(mBase, _consts[938]))
										F_rewriteClientCommandArgument(m, l0, int32(0), v86)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
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
				} else {
					F_incrRefCount(m, v28)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						F_setKey(m, l0, v37, v39, v6+int32(12), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
							F_rewriteClientCommandArgument(m, l0, int32(2), v46)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v67 = v49
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+28))
								F_notifyKeyspaceEvent(m, int32(8), int32(_a188), v70, v72)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									v75 = int32(_a44)
									v77 = *(*int64)(unsafe.Add(mBase, _consts[83]))
									*(*int64)(unsafe.Add(mBase, _consts[83])) = v77 + int64(1)
									F_commitDeferredReplyBuffer(m, l0, int32(1))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										v86 = *(*int32)(unsafe.Add(mBase, _consts[938]))
										F_rewriteClientCommandArgument(m, l0, int32(0), v86)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
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
			} else {
				v23 = F_checkType(m, l0, v18, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					if v23 != 0 {
						m.G0 = v6 + int32(16)
						return
					} else {
						F_addReplyBulk(m, l0, v18)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v28
							v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
							if v30&int32(1) == int32(0) {
								v50 = F_tryObjectEncoding(m, v28)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v50
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
									F_setKey(m, l0, v53, v55, v6+int32(12), int32(0))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
										F_incrRefCount(m, v61)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v65
											v67 = v64
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+28))
											F_notifyKeyspaceEvent(m, int32(8), int32(_a188), v70, v72)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return
											} else {
												v75 = int32(_a44)
												v77 = *(*int64)(unsafe.Add(mBase, _consts[83]))
												*(*int64)(unsafe.Add(mBase, _consts[83])) = v77 + int64(1)
												F_commitDeferredReplyBuffer(m, l0, int32(1))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													v86 = *(*int32)(unsafe.Add(mBase, _consts[938]))
													F_rewriteClientCommandArgument(m, l0, int32(0), v86)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
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
							} else {
								F_incrRefCount(m, v28)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									F_setKey(m, l0, v37, v39, v6+int32(12), int32(0))
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
										F_rewriteClientCommandArgument(m, l0, int32(2), v46)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v67 = v49
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+28))
											F_notifyKeyspaceEvent(m, int32(8), int32(_a188), v70, v72)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return
											} else {
												v75 = int32(_a44)
												v77 = *(*int64)(unsafe.Add(mBase, _consts[83]))
												*(*int64)(unsafe.Add(mBase, _consts[83])) = v77 + int64(1)
												F_commitDeferredReplyBuffer(m, l0, int32(1))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													v86 = *(*int32)(unsafe.Add(mBase, _consts[938]))
													F_rewriteClientCommandArgument(m, l0, int32(0), v86)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
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
					}
				}
			}
		}
	}
}
func F_getsockname(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v4 = int32(0)
	v7 = m.Env.X__syscall_getsockname(m, l0, l1, l2, v4, v4, v4)
	mBase = m.M
	if base.Ui32(v7) < base.Ui32(int32(-4095)) {
		v15 = v7
	} else {
		v10 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0) - v7
		v15 = int32(-1)
	}
	return v15
}
func F_getuid(m *base.Module) int32 {
	return int32(0)
}
func F_glob(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var __phi244 int32
	_ = __phi244
	var v245 int32
	_ = v245
	var __phi245 int32
	_ = __phi245
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var __phi411 int32
	_ = __phi411
	var v412 int32
	_ = v412
	var __phi412 int32
	_ = __phi412
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v435 int32
	_ = v435
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v493 int32
	_ = v493
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v689 int32
	_ = v689
	var v699 int32
	_ = v699
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v756 int32
	_ = v756
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v794 int32
	_ = v794
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v960 int32
	_ = v960
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v1000 int32
	_ = v1000
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(4128)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1047]))) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1048]))) = v13 + int32(4124)
	if l1&int32(8) == v5 {
		v26 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = l1 & int32(32)
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v26 = v25
	goto L1
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v26
	goto L3
L5:
	;
	m.G0 = v13 + int32(4128)
	return v1000
L6:
	;
	v1000 = int32(1)
	goto L5
L7:
	;
	if v28 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L8:
	;
	if l1&int32(16) == int32(0) {
		goto L163
	} else {
		goto L164
	}
L9:
	;
	v36 = F_strlen(m, l0)
	mBase = m.M
	v38 = v36 + int32(1)
	v39 = F_emscripten_builtin_malloc(m, v38)
	mBase = m.M
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v689 = int32(0)
	goto L8
L11:
	;
	if v42 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	v41 = F___memcpy(m, v39, l0, v38)
	mBase = m.M
	v42 = v41
	goto L11
L13:
	;
	v42 = int32(0)
	goto L11
L14:
	;
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+16)) = uint8(v45)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v42
	if l1&int32(20480) == v45 {
		v161 = v45
		v162 = v42
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v42 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L16:
	;
	if l2 != 0 {
		goto L45
	} else {
		goto L46
	}
L17:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v56 != int32(126) {
		v161 = int32(0)
		v162 = v42
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v60 = v13 + int32(8)
	v62 = v13 + int32(16)
	v69 = m.G0
	v71 = v69 - int32(32)
	m.G0 = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v75 = v73 + int32(1)
	v77 = F___strchrnul(m, v75, int32(47))
	mBase = m.M
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v78 == int32(0) {
		v85 = v77
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v154 != 0 {
		v174 = v154
		goto L15
	} else {
		goto L44
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v85
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v87 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v81 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v81)
	v85 = v77 + int32(1)
	goto L20
L22:
	;
	m.G0 = v71 + int32(32)
	goto L19
L23:
	;
	v114 = v111
	v119 = int32(0)
	goto L37
L24:
	;
	if v104 != int32(48) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v98 = F_getuid(m)
	mBase = m.M
	v102 = F_getpwuid_r(m, v98, v71+int32(4), v62, int32(4096), v71)
	mBase = m.M
	v104 = v102
	goto L24
L26:
	;
	v97 = F_getpwnam_r(m, v75, v71+int32(4), v62, int32(4096), v71)
	mBase = m.M
	v104 = v97
	goto L24
L27:
	;
	v89 = F_getenv(m, int32(_a2172))
	mBase = m.M
	if v89 != 0 {
		v111 = v89
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v90 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	if v104 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v154 = int32(1)
	goto L22
L32:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v71)+24))
	v111 = v110
	goto L23
L33:
	;
	v154 = int32(3)
	goto L22
L34:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v108 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v62+v136))) = uint8(v78)
	if v78 == int32(0) {
		v146 = v136
		goto L42
	} else {
		goto L43
	}
L37:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v121 == int32(0) {
		v136 = v119
		goto L36
	} else {
		goto L39
	}
L38:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v133 != 0 {
		v154 = int32(3)
		goto L22
	} else {
		goto L41
	}
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v62+v119))) = uint8(v121)
	v126 = int32(1)
	v127 = v114 + v126
	v129 = v119 + v126
	if v129 != int32(4094) {
		v114 = v127
		v119 = v129
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v136 = int32(4094)
	goto L36
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(12)))) = v146
	v154 = int32(0)
	goto L22
L43:
	;
	v142 = v136 + int32(1)
	v144 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62+v142))) = uint8(v144)
	v146 = v142
	goto L42
L44:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v161 = v160
	v162 = v159
	goto L16
L45:
	;
	v167 = l2
	goto L47
L46:
	;
	v167 = int32(1379)
	goto L47
L47:
	;
	v170 = F_do_glob(m, v13+int32(16), v161, int32(0), v162, l1, v167, v13+int32(4120))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	return int32(0)
L49:
	;
	v174 = v170
	goto L15
L50:
	;
	v650 = int32(1)
	if v174 != v650 {
		goto L150
	} else {
		goto L151
	}
L51:
	;
	goto L50
L52:
	;
	v186 = int32(-8)
	v187 = v42 + v186
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(-4))))
	v192 = v190 & v186
	v193 = v187 + v192
	if v190&int32(1) != 0 {
		v317 = v192
		v318 = v187
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if base.Ui32(v193) <= base.Ui32(v318) {
		goto L51
	} else {
		goto L88
	}
L54:
	;
	if v190&int32(2) == int32(0) {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v201 = v187 - v200
	v203 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	if base.Ui32(v201) < base.Ui32(v203) {
		goto L51
	} else {
		goto L56
	}
L56:
	;
	v205 = v200 + v192
	v207 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v201 == v207 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	if v223 == int32(0) {
		v317 = v205
		v318 = v201
		goto L53
	} else {
		goto L76
	}
L58:
	;
	v276 = int32(0)
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+12)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v209)+8)) = v212
	v317 = v205
	v318 = v201
	goto L53
L60:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v258 = int32(3)
	if v257&v258 != v258 {
		v317 = v205
		v318 = v201
		goto L53
	} else {
		goto L75
	}
L61:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	if base.Ui32(int32(255)) < base.Ui32(v200) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v201)+24))
	if v209 == v201 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	if v209 != v212 {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v214 = int32(0)
	v216 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v216 & base.I32_rotl(int32(-2), int32(base.Ui32(v200)>>(uint(int32(3))%32)))
	v317 = v205
	v318 = v201
	goto L53
L65:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v201)+20))
	if v228 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v225)+12)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v209)+8)) = v225
	v276 = v209
	goto L57
L67:
	;
	__phi244 = v238
	__phi245 = v239
	v244 = __phi244
	v245 = __phi245
	goto L71
L68:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v201)+16))
	if v233 == int32(0) {
		goto L58
	} else {
		goto L70
	}
L69:
	;
	v238 = v228
	v239 = v201 + int32(20)
	goto L67
L70:
	;
	v238 = v233
	v239 = v201 + int32(16)
	goto L67
L71:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v244)+20))
	if v251 != 0 {
		__phi244 = v251
		__phi245 = v244 + int32(20)
		v244 = __phi244
		v245 = __phi245
		goto L71
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = int32(0)
	v276 = v244
	goto L57
L73:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
	if v254 != 0 {
		__phi244 = v254
		__phi245 = v244 + int32(16)
		v244 = __phi244
		v245 = __phi245
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v193)+4)) = v257 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v205 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v205
	goto L50
L76:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v201)+28))
	v287 = v285 << (uint(int32(2)) % 32)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v287)+uint32(_consts[519])))
	if v201 != v290 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+24)) = v223
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v201)+16))
	if v307 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L78:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	if v300 != v201 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+uint32(_consts[519]))) = v276
	if v276 != 0 {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v293 = int32(0)
	v295 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v295 & base.I32_rotl(int32(-2), v285)
	v317 = v205
	v318 = v201
	goto L53
L81:
	;
	if v276 == int32(0) {
		v317 = v205
		v318 = v201
		goto L53
	} else {
		goto L84
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+20)) = v276
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+16)) = v276
	goto L81
L84:
	;
	goto L77
L85:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v201)+20))
	if v312 == int32(0) {
		v317 = v205
		v318 = v201
		goto L53
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+16)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v307)+24)) = v276
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+20)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v312)+24)) = v276
	v317 = v205
	v318 = v201
	goto L53
L88:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if v327&int32(1) == int32(0) {
		goto L51
	} else {
		goto L89
	}
L89:
	;
	if v327&int32(2) != 0 {
		goto L94
	} else {
		goto L95
	}
L90:
	;
	if base.Ui32(int32(255)) < base.Ui32(v493) {
		goto L128
	} else {
		goto L129
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+4)) = v373 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v318+v373))) = v373
	if v318 != v357 {
		v493 = v373
		goto L90
	} else {
		goto L127
	}
L92:
	;
	if v390 == int32(0) {
		goto L91
	} else {
		goto L115
	}
L93:
	;
	v435 = int32(0)
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+4)) = v327 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v318)+4)) = v317 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v318+v317))) = v317
	v493 = v317
	goto L90
L95:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	if v193 != v335 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v193 != v357 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v337 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v318
	v341 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v342 = v341 + v317
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v318)+4)) = v342 | int32(1)
	v348 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v318 != v348 {
		goto L51
	} else {
		goto L98
	}
L98:
	;
	v350 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v350
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v350
	goto L50
L99:
	;
	v373 = v327&int32(-8) + v317
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	if base.Ui32(int32(255)) < base.Ui32(v327) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v359 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v318
	v363 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	v364 = v363 + v317
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v318)+4)) = v364 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v318+v364))) = v364
	goto L50
L101:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v193)+24))
	if v374 == v193 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	if v374 != v377 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v377)+12)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v374)+8)) = v377
	goto L91
L104:
	;
	v379 = int32(0)
	v381 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v381 & base.I32_rotl(int32(-2), int32(base.Ui32(v327)>>(uint(int32(3))%32)))
	goto L91
L105:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v193)+20))
	if v395 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v392)+12)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v374)+8)) = v392
	v435 = v374
	goto L92
L107:
	;
	__phi411 = v405
	__phi412 = v406
	v411 = __phi411
	v412 = __phi412
	goto L111
L108:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	if v400 == int32(0) {
		goto L93
	} else {
		goto L110
	}
L109:
	;
	v405 = v395
	v406 = v193 + int32(20)
	goto L107
L110:
	;
	v405 = v400
	v406 = v193 + int32(16)
	goto L107
L111:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v411)+20))
	if v418 != 0 {
		__phi411 = v418
		__phi412 = v411 + int32(20)
		v411 = __phi411
		v412 = __phi412
		goto L111
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v412))) = int32(0)
	v435 = v411
	goto L92
L113:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v411)+16))
	if v421 != 0 {
		__phi411 = v421
		__phi412 = v411 + int32(16)
		v411 = __phi411
		v412 = __phi412
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v193)+28))
	v446 = v444 << (uint(int32(2)) % 32)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v446)+uint32(_consts[519])))
	if v193 != v449 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435)+24)) = v390
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	if v466 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L117:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v390)+16))
	if v459 != v193 {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v446)+uint32(_consts[519]))) = v435
	if v435 != 0 {
		goto L116
	} else {
		goto L119
	}
L119:
	;
	v452 = int32(0)
	v454 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v454 & base.I32_rotl(int32(-2), v444)
	goto L91
L120:
	;
	if v435 == int32(0) {
		goto L91
	} else {
		goto L123
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+20)) = v435
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+16)) = v435
	goto L120
L123:
	;
	goto L116
L124:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v193)+20))
	if v471 == int32(0) {
		goto L91
	} else {
		goto L126
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435)+16)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v466)+24)) = v435
	goto L124
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435)+20)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v471)+24)) = v435
	goto L91
L127:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v373
	goto L50
L128:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v493) {
		v540 = int32(31)
		goto L133
	} else {
		goto L134
	}
L129:
	;
	v505 = v493 & int32(-8)
	v507 = v505 + int32(9128464)
	v509 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v513 = int32(1) << (uint(int32(base.Ui32(v493)>>(uint(int32(3))%32))) % 32)
	if v509&v513 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v505)+uint32(_consts[523]))) = v318
	*(*int32)(unsafe.Add(mBase, uint32(v519)+12)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(v318)+12)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v318)+8)) = v519
	goto L50
L131:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v505)+uint32(_consts[523])))
	v519 = v518
	goto L130
L132:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v509 | v513
	v519 = v507
	goto L130
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+28)) = v540
	*(*int64)(unsafe.Add(mBase, uint32(v318)+16)) = int64(0)
	v545 = v540 << (uint(int32(2)) % 32)
	v549 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	v551 = int32(1) << (uint(v540) % 32)
	if v549&v551 != 0 {
		goto L138
	} else {
		goto L139
	}
L134:
	;
	v530 = base.I32_clz(int32(base.Ui32(v493) >> (uint(int32(8)) % 32)))
	v533 = int32(1)
	v540 = int32(base.Ui32(v493)>>(uint(int32(38)-v530)%32))&v533 - v530<<(uint(v533)%32) + int32(62)
	goto L133
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318+v612))) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v318)+12)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(v318+v610))) = v613
	v624 = int32(0)
	v626 = *(*int32)(unsafe.Add(mBase, _consts[524]))
	v627 = int32(-1)
	v628 = v626 + v627
	if v628 != 0 {
		goto L147
	} else {
		goto L148
	}
L136:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v574)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v604)+12)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(v574)+8)) = v318
	v610 = int32(24)
	v612 = int32(8)
	v613 = int32(0)
	v614 = v574
	v615 = v604
	goto L135
L137:
	;
	v610 = v595
	v612 = v597
	v613 = v318
	v614 = v318
	v615 = v600
	goto L135
L138:
	;
	if v540 == int32(31) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v549 | v551
	*(*int32)(unsafe.Add(mBase, uint32(v545)+uint32(_consts[519]))) = v318
	v595 = int32(8)
	v597 = int32(24)
	v600 = v545 + int32(9128728)
	goto L137
L140:
	;
	v566 = int32(0)
	goto L142
L141:
	;
	v566 = int32(25) - int32(base.Ui32(v540)>>(uint(int32(1))%32))
	goto L142
L142:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v545)+uint32(_consts[519])))
	v571 = v493 << (uint(v566) % 32)
	v574 = v568
	goto L143
L143:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	if v578&int32(-8) == v493 {
		goto L136
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v588+int32(16)))) = v318
	v595 = int32(8)
	v597 = int32(24)
	v600 = v574
	goto L137
L145:
	;
	v588 = v574 + int32(base.Ui32(v571)>>(uint(int32(29))%32))&int32(4)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v588)+16))
	if v589 != 0 {
		v571 = v571 << (uint(int32(1)) % 32)
		v574 = v589
		goto L143
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	v630 = v628
	goto L149
L148:
	;
	v630 = v627
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, _consts[524])) = v630
	goto L51
L150:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1047])))
	if v665 == int32(0) {
		v689 = v174
		goto L8
	} else {
		goto L158
	}
L151:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1047])))
	if v656 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v1000 = v650
	goto L5
L153:
	;
	goto L152
L154:
	;
	v659 = v656
	goto L155
L155:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v659)))
	F_emscripten_builtin_free(m, v659)
	mBase = m.M
	if v661 != 0 {
		v659 = v661
		goto L155
	} else {
		goto L157
	}
L156:
	;
	goto L153
L157:
	;
	goto L156
L158:
	;
	v671 = v665
	v677 = int32(0)
	goto L159
L159:
	;
	v680 = v677 + int32(1)
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v671)))
	if v681 != 0 {
		v671 = v681
		v677 = v680
		goto L159
	} else {
		goto L161
	}
L160:
	;
	if v680 != 0 {
		v804 = v174
		v805 = v680
		goto L7
	} else {
		goto L162
	}
L161:
	;
	goto L160
L162:
	;
	v689 = v174
	goto L8
L163:
	;
	if v689 != 0 {
		v804 = v689
		v805 = int32(0)
		goto L7
	} else {
		goto L189
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1048]))) = v13 + int32(4124)
	v699 = int32(1)
	if l0&int32(3) == int32(0) {
		v723 = l0
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v763 = F_emscripten_builtin_malloc(m, v756+int32(6))
	mBase = m.M
	if v763 != 0 {
		goto L182
	} else {
		goto L183
	}
L166:
	;
	v756 = v748 - l0
	goto L165
L167:
	;
	v727 = v723
	goto L175
L168:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v709 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v712 = l0
	goto L171
L170:
	;
	v756 = l0 - l0
	goto L165
L171:
	;
	v716 = v712 + int32(1)
	if v716&int32(3) == int32(0) {
		v723 = v716
		goto L167
	} else {
		goto L173
	}
L173:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v716))))
	if v721 != 0 {
		v712 = v716
		goto L171
	} else {
		goto L174
	}
L174:
	;
	v748 = v716
	goto L166
L175:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v727)))
	v736 = int32(-2139062144)
	if (int32(16843008)-v733|v733)&v736 == v736 {
		v727 = v727 + int32(4)
		goto L175
	} else {
		goto L177
	}
L176:
	;
	v742 = v727
	goto L178
L177:
	;
	goto L176
L178:
	;
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742))))
	if v746 != 0 {
		v742 = v742 + int32(1)
		goto L178
	} else {
		goto L180
	}
L179:
	;
	v748 = v742
	goto L166
L180:
	;
	goto L179
L181:
	;
	if v794 != 0 {
		v1000 = v699
		goto L5
	} else {
		goto L188
	}
L182:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1048])))
	*(*int32)(unsafe.Add(mBase, uint32(v765))) = v763
	v767 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v763))) = v767
	v773 = F___memcpy(m, v763+int32(4), l0, v756+int32(1))
	mBase = m.M
	if v756 == v767 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v794 = int32(-1)
	goto L181
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1048]))) = v763
	v794 = int32(0)
	goto L181
L185:
	;
	goto L184
L188:
	;
	v804 = v689
	v805 = v699
	goto L7
L189:
	;
	v1000 = int32(3)
	goto L5
L190:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1047])))
	if v972 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L191:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1047])))
	if v805 == int32(0) {
		v943 = v914
		goto L225
	} else {
		goto L226
	}
L192:
	;
	v879 = F_emscripten_builtin_malloc(m, (v805+v26)<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v879
	if v879 == int32(0) {
		goto L190
	} else {
		goto L220
	}
L193:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v811 = v810 + v26
	v814 = (v811 + v805) << (uint(int32(2)) % 32)
	v816 = v814 + int32(4)
	if v809 != 0 {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v860
	v904 = v811
	v913 = v860
	goto L191
L195:
	;
	if v860 != 0 {
		goto L194
	} else {
		goto L213
	}
L196:
	;
	if base.Ui32(v816) < base.Ui32(int32(-64)) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v819 = F_emscripten_builtin_malloc(m, v816)
	mBase = m.M
	v860 = v819
	goto L195
L198:
	;
	v826 = int32(-8)
	if base.Ui32(v816) < base.Ui32(int32(11)) {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v822 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v822))) = int32(48)
	v860 = int32(0)
	goto L195
L200:
	;
	v841 = F_emscripten_builtin_malloc(m, v816)
	mBase = m.M
	if v841 != 0 {
		goto L205
	} else {
		goto L206
	}
L201:
	;
	v835 = int32(16)
	goto L203
L202:
	;
	v835 = (v814 + int32(15)) & v826
	goto L203
L203:
	;
	v836 = F_try_realloc_chunk(m, v809+v826, v835)
	mBase = m.M
	if v836 == int32(0) {
		goto L200
	} else {
		goto L204
	}
L204:
	;
	v860 = v836 + int32(8)
	goto L195
L205:
	;
	v843 = int32(-4)
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v809+v843)))
	if v847&int32(3) != 0 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v860 = int32(0)
	goto L195
L207:
	;
	v850 = v843
	goto L209
L208:
	;
	v850 = int32(-8)
	goto L209
L209:
	;
	v853 = v850 + v847&int32(-8)
	if base.Ui32(v853) < base.Ui32(v816) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v855 = v853
	goto L212
L211:
	;
	v855 = v816
	goto L212
L212:
	;
	v856 = F___memcpy(m, v841, v809, v855)
	mBase = m.M
	F_emscripten_builtin_free(m, v809)
	mBase = m.M
	v860 = v841
	goto L195
L213:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1047])))
	if v864 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	goto L6
L215:
	;
	goto L214
L216:
	;
	v867 = v864
	goto L217
L217:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v867)))
	F_emscripten_builtin_free(m, v867)
	mBase = m.M
	if v869 != 0 {
		v867 = v869
		goto L217
	} else {
		goto L219
	}
L218:
	;
	goto L215
L219:
	;
	goto L218
L220:
	;
	v883 = int32(0)
	if v26 == v883 {
		v904 = v883
		v913 = v879
		goto L191
	} else {
		goto L221
	}
L221:
	;
	v886 = v883
	goto L222
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v879+v886<<(uint(int32(2))%32)))) = int32(0)
	v902 = v886 + int32(1)
	if v902 != v26 {
		v886 = v902
		goto L222
	} else {
		goto L224
	}
L223:
	;
	v904 = v26
	v913 = v879
	goto L191
L224:
	;
	goto L223
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1048]))) = v943
	v952 = int32(2)
	v954 = v913 + v904<<(uint(v952)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v954+v805<<(uint(v952)%32)))) = int32(0)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v960 + v805
	if l1&int32(4) != 0 {
		goto L230
	} else {
		goto L231
	}
L226:
	;
	v923 = v914
	v926 = int32(0)
	goto L227
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v913+v904<<(uint(int32(2))%32)+v926<<(uint(int32(2))%32)))) = v923 + int32(4)
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v923)))
	v939 = v926 + int32(1)
	if v939 != v805 {
		v923 = v937
		v926 = v939
		goto L227
	} else {
		goto L229
	}
L228:
	;
	v943 = v937
	goto L225
L229:
	;
	goto L228
L230:
	;
	v1000 = v804
	goto L5
L231:
	;
	F_qsort(m, v954, v805, int32(4), int32(1380))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L48
	} else {
		goto L232
	}
L232:
	;
	goto L230
L233:
	;
	goto L6
L234:
	;
	goto L233
L235:
	;
	v975 = v972
	goto L236
L236:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v975)))
	F_emscripten_builtin_free(m, v975)
	mBase = m.M
	if v977 != 0 {
		v975 = v977
		goto L236
	} else {
		goto L238
	}
L237:
	;
	goto L234
L238:
	;
	goto L237
}
