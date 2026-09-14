package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_ACLCheckUserCredentials(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v350 int32
	_ = v350
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_objectGetVal(m, l0)
	mBase = m.M
	v13 = F_objectGetVal(m, l0)
	mBase = m.M
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
	switch v16 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v33 = int32(0)
		goto L1
	}
L1:
	;
	v34 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v34
	v37 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v39 = v9 + int32(12)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v33 == v34 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
	v33 = v32
	goto L1
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
	v33 = v29
	goto L1
L4:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
	v33 = v26
	goto L1
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
	v33 = v23
	goto L1
L6:
	;
	v33 = int32(base.Ui32(v16) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v235 != 0 {
		goto L47
	} else {
		goto L48
	}
L8:
	;
	if v192 != v33 {
		goto L35
	} else {
		goto L36
	}
L9:
	;
	v183 = int32(0)
	v189 = v48
	v190 = v49
	v192 = v183
	v196 = v183
	goto L8
L10:
	;
	if base.Ui32(v49) < base.Ui32(int32(8)) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v60 = v48
	v61 = v49
	v63 = int32(0)
	goto L13
L12:
	;
	v189 = v173
	v190 = v174
	v192 = v176
	v196 = base.B2i32(v179 != int32(0))
	goto L8
L13:
	;
	v69 = int32(base.Ui32(v61) >> (uint(int32(3)) % 32))
	v70 = int32(4)
	v71 = v60 + v70
	if v61&v70 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v173 = v164
	v174 = v165
	v176 = v149
	v179 = v154
	goto L12
L15:
	;
	v154 = int32(0)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v71+v69+(v154-v69)&int32(3)+v142<<(uint(int32(2))%32))))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if base.Ui32(v165) < base.Ui32(int32(8)) {
		v173 = v164
		v174 = v165
		v176 = v149
		v179 = v154
		goto L12
	} else {
		goto L33
	}
L16:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v63))))
	v120 = int32(0)
	goto L27
L17:
	;
	v76 = int32(0)
	if base.Ui32(v33) <= base.Ui32(v63) {
		v109 = v63
		v112 = v76
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v112 == v69 {
		v142 = v76
		v149 = v109
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v86 = v63
	v89 = v76
	goto L20
L20:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v89))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v86))))
	if v92 != v94 {
		v109 = v86
		v112 = v89
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v109 = v97
	v112 = v99
	goto L18
L22:
	;
	v96 = int32(1)
	v97 = v86 + v96
	v99 = v89 + v96
	if base.Ui32(v69) <= base.Ui32(v99) {
		v109 = v97
		v112 = v99
		goto L18
	} else {
		goto L23
	}
L23:
	;
	if base.Ui32(v97) < base.Ui32(v33) {
		v86 = v97
		v89 = v99
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v173 = v60
	v174 = v61
	v176 = v109
	v179 = v112
	goto L12
L26:
	;
	if v120 != v69 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v120))))
	if v133 == v117&int32(255) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v135 = int32(1)
	v137 = v120 + v135
	if v137 != v69 {
		v120 = v137
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v189 = v60
	v190 = v61
	v192 = v63
	v196 = v135
	goto L8
L31:
	;
	v142 = v120
	v149 = v63 + int32(1)
	goto L15
L32:
	;
	v173 = v60
	v174 = v61
	v176 = v63
	v179 = v69
	goto L12
L33:
	;
	if base.Ui32(v149) < base.Ui32(v33) {
		v60 = v164
		v61 = v165
		v63 = v149
		goto L13
	} else {
		goto L34
	}
L34:
	;
	goto L14
L35:
	;
	goto L7
L36:
	;
	if v190&int32(1) == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v204 = v190 & int32(4)
	if v196&base.B2i32(v204 != int32(0)) != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	if v39 == int32(0) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	if v190&int32(2) != 0 {
		v230 = int32(0)
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v230
	goto L35
L41:
	;
	v214 = int32(3)
	v215 = int32(base.Ui32(v190) >> (uint(v214) % 32))
	if v204 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v225 = int32(4)
	goto L44
L43:
	;
	v225 = v215 << (uint(int32(2)) % 32)
	goto L44
L44:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v189+v215+(int32(0)-v215)&v214+v225+int32(4))))
	v230 = v229
	goto L40
L45:
	;
	m.G0 = v9 + int32(16)
	return v350
L46:
	;
	v350 = int32(-1)
	goto L45
L47:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v239&int32(2) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(44)
	goto L46
L50:
	;
	if v239&int32(4) != 0 {
		v350 = int32(0)
		goto L45
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(28)
	goto L46
L53:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v235)+8))
	v252 = v9 + int32(4)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	*(*int32)(unsafe.Add(mBase, uint32(v252)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v253
	goto L54
L54:
	;
	v257 = F_objectGetVal(m, l1)
	mBase = m.M
	v259 = F_objectGetVal(m, l1)
	mBase = m.M
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259+int32(-1)))))
	switch v262 & int32(7) {
	case 0:
		goto L60
	case 1:
		goto L59
	case 2:
		goto L58
	case 3:
		goto L57
	case 4:
		goto L56
	default:
		v279 = int32(0)
		goto L55
	}
L55:
	;
	v280 = F_ACLHashPassword(m, v257, v279)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v259+int32(-17))))
	v279 = v278
	goto L55
L57:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v259+int32(-9))))
	v279 = v275
	goto L55
L58:
	;
	v272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259+int32(-5)))))
	v279 = v272
	goto L55
L59:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259+int32(-3)))))
	v279 = v269
	goto L55
L60:
	;
	v279 = int32(base.Ui32(v262) >> (uint(int32(3)) % 32))
	goto L55
L61:
	;
	return int32(0)
L62:
	;
	goto L64
L63:
	;
	F_sdsfree(m, v280)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L61
	} else {
		goto L75
	}
L64:
	;
	v291 = v9 + int32(4)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	if v293 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	F_sdsfree(m, v280)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L61
	} else {
		goto L74
	}
L66:
	;
	if v293 == int32(0) {
		goto L63
	} else {
		goto L69
	}
L67:
	;
	goto L66
L68:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v293+base.B2i32(v296 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = v302
	goto L67
L69:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v293)+8))
	v307 = int32(0)
	v309 = v307
	v313 = v307
	goto L70
L70:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306+v309))))
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v309))))
	v323 = v309 | int32(1)
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306+v323))))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v323))))
	v330 = v313 | base.I32_extend8_s(v316^v318) | base.I32_extend8_s(v325^v327)
	v332 = v309 + int32(2)
	if v332 != int32(64) {
		v309 = v332
		v313 = v330
		goto L70
	} else {
		goto L72
	}
L71:
	;
	if v330 != 0 {
		goto L64
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	goto L65
L74:
	;
	v350 = int32(0)
	goto L45
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(28)
	goto L46
}
func F_ACLCopyUser(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_listRelease(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_listRelease(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v11 = F_listDup(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v15 = F_listDup(m, v14)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v15
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v20 == int32(0) {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v25
						if v25 == int32(0) {
							return
						} else {
							F_incrRefCount(m, v25)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						F_decrRefCount(m, v20)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v25
							if v25 == int32(0) {
								return
							} else {
								F_incrRefCount(m, v25)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
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
func F_ACLCreateUser(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	v3 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if l1 == v3 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v283
L2:
	;
	if v206 != 0 {
		v283 = v3
		goto L1
	} else {
		goto L40
	}
L3:
	;
	if v165 != l1 {
		v206 = v3
		goto L30
	} else {
		goto L31
	}
L4:
	;
	v156 = int32(0)
	v163 = v22
	v165 = v156
	v169 = v156
	goto L3
L5:
	;
	if base.Ui32(v22) < base.Ui32(int32(8)) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v33 = v21
	v34 = v22
	v36 = int32(0)
	goto L8
L7:
	;
	v163 = v147
	v165 = v149
	v169 = base.B2i32(v152 != int32(0))
	goto L3
L8:
	;
	v42 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
	v43 = int32(4)
	v44 = v33 + v43
	if v34&v43 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v147 = v138
	v149 = v122
	v152 = v127
	goto L7
L10:
	;
	v127 = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v44+v42+(v127-v42)&int32(3)+v115<<(uint(int32(2))%32))))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	if base.Ui32(v138) < base.Ui32(int32(8)) {
		v147 = v138
		v149 = v122
		v152 = v127
		goto L7
	} else {
		goto L28
	}
L11:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v36))))
	v93 = int32(0)
	goto L22
L12:
	;
	v49 = int32(0)
	if base.Ui32(l1) <= base.Ui32(v36) {
		v82 = v36
		v85 = v49
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v85 == v42 {
		v115 = v49
		v122 = v82
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v59 = v36
	v62 = v49
	goto L15
L15:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v62))))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v59))))
	if v65 != v67 {
		v82 = v59
		v85 = v62
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v82 = v70
	v85 = v72
	goto L13
L17:
	;
	v69 = int32(1)
	v70 = v59 + v69
	v72 = v62 + v69
	if base.Ui32(v42) <= base.Ui32(v72) {
		v82 = v70
		v85 = v72
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if base.Ui32(v70) < base.Ui32(l1) {
		v59 = v70
		v62 = v72
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v147 = v34
	v149 = v82
	v152 = v85
	goto L7
L21:
	;
	if v93 != v42 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v93))))
	if v106 == v90&int32(255) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v108 = int32(1)
	v110 = v93 + v108
	if v110 != v42 {
		v93 = v110
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v163 = v34
	v165 = v36
	v169 = v108
	goto L3
L26:
	;
	v115 = v93
	v122 = v36 + int32(1)
	goto L10
L27:
	;
	v147 = v34
	v149 = v36
	v152 = v42
	goto L7
L28:
	;
	if base.Ui32(v122) < base.Ui32(l1) {
		v33 = v137
		v34 = v138
		v36 = v122
		goto L8
	} else {
		goto L29
	}
L29:
	;
	goto L9
L30:
	;
	goto L2
L31:
	;
	v171 = int32(0)
	if v163&int32(1) == v171 {
		v206 = v171
		goto L30
	} else {
		goto L32
	}
L32:
	;
	if v169&base.B2i32(v163&int32(4) != int32(0)) != 0 {
		v206 = v171
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v206 = int32(1)
	goto L30
L40:
	;
	v209 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	return int32(0)
L42:
	;
	v213 = F_sdsnewlen(m, l0, l1)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+4)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v213
	v218 = F_listCreate(m)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v209)+8)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v218)+12)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v218)+16)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v218)+8)) = int32(5)
	v229 = F_listCreate(m)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+12)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v229)+8)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v229)+12)) = int32(7)
	v237 = F_valkey_malloc(m, int32(160))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v240 | int32(17)
	v244 = F_listCreate(m)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+140)) = v244
	v247 = F_listCreate(m)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L41
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+144)) = v247
	v250 = F_intsetNew(m)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L41
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+136)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v237)+152)) = v250
	v255 = F_sdsempty(m)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L41
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+148)) = v255
	v258 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v244)+12)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v244)+16)) = int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v244)+8)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v247)+12)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v247)+16)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v247)+8)) = int32(5)
	v275 = F__emscripten_memset_bulkmem(m, v237+v258, base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L51
L51:
	;
	v276 = F_listAddNodeHead(m, v229, v237)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L41
	} else {
		goto L52
	}
L52:
	;
	v278 = int32(0)
	v279 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v281 = F_raxInsert(m, v279, l0, l1, v209, v278)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L41
	} else {
		goto L53
	}
L53:
	;
	v283 = v209
	goto L1
}
func F_ACLFreeLogEntry(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_sdsfree(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		F_sdsfree(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			F_sdsfree(m, v8)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_valkey_free(m, l0)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_ACLFreeSelector(m *base.Module, l0 int32) {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	F_listRelease(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	F_listRelease(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	F_intsetFree(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	F_sdsfree(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L21
	}
L7:
	;
	v24 = int32(0)
	goto L8
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v31 = v24 << (uint(int32(2)) % 32)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v31)))
	if v33 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	F_valkey_free(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L20
	}
L10:
	;
	v72 = v24 + int32(1)
	if v72 != int32(1024) {
		v24 = v72
		goto L8
	} else {
		goto L19
	}
L11:
	;
	v36 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v37 == v36 {
		v60 = v33
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_valkey_free(m, v60)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	v44 = v36
	v45 = v37
	goto L14
L14:
	;
	F_sdsfree(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v60 = v50
	goto L12
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48+v31)))
	v52 = v44 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50+v52<<(uint(int32(2))%32))))
	if v56 != 0 {
		v44 = v52
		v45 = v56
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L10
L19:
	;
	goto L9
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = int32(0)
	goto L6
L21:
	;
	return
}
func F_ACLFreeUser(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_sdsfree(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v6 == int32(0) {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_listRelease(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_listRelease(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			F_decrRefCount(m, v6)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				F_listRelease(m, v13)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					F_listRelease(m, v16)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						F_valkey_free(m, l0)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
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
func F_ACLFreeUserVoid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_sdsfree(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v6 == int32(0) {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_listRelease(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_listRelease(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			F_decrRefCount(m, v6)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				F_listRelease(m, v13)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					F_listRelease(m, v16)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						F_valkey_free(m, l0)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
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
func F_ACLGetCommandCategoryFlagByName(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
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
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int64
	_ = v116
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
	if base.B2i32(v9 == int64(0)) == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	return int64(0)
L3:
	;
	return v116
L4:
	;
	if v51-v53 == int32(0) {
		v116 = v9
		goto L3
	} else {
		goto L16
	}
L5:
	;
	v51 = F_tolower(m, v47)
	mBase = m.M
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v53 = F_tolower(m, v52)
	mBase = m.M
	goto L4
L6:
	;
	v21 = l0
	v22 = v16
	v23 = v19
	goto L9
L7:
	;
	v47 = int32(0)
	v48 = v16
	goto L5
L8:
	;
	v47 = v44 & int32(255)
	v48 = v43
	goto L5
L9:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v25 == int32(0) {
		v43 = v22
		v44 = v23
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v43 = v37
	v44 = int32(0)
	goto L8
L11:
	;
	v29 = v23 & int32(255)
	if v29 == v25 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v36 = int32(1)
	v37 = v22 + v36
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v38 != 0 {
		v21 = v21 + v36
		v22 = v37
		v23 = v38
		goto L9
	} else {
		goto L15
	}
L13:
	;
	v31 = F_tolower(m, v29)
	mBase = m.M
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v33 = F_tolower(m, v32)
	mBase = m.M
	if v31 == v33 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v43 = v22
	v44 = v35
	goto L8
L15:
	;
	goto L10
L16:
	;
	v58 = v2
	goto L17
L17:
	;
	v63 = v58 + int32(1)
	v66 = v8 + v63<<(uint(int32(4))%32)
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
	if base.B2i32(v67 == int64(0)) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v116 = v67
	goto L3
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	return int64(0)
L21:
	;
	if v109-v111 != 0 {
		v58 = v63
		goto L17
	} else {
		goto L33
	}
L22:
	;
	v109 = F_tolower(m, v105)
	mBase = m.M
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v111 = F_tolower(m, v110)
	mBase = m.M
	goto L21
L23:
	;
	v79 = l0
	v80 = v74
	v81 = v77
	goto L26
L24:
	;
	v105 = int32(0)
	v106 = v74
	goto L22
L25:
	;
	v105 = v102 & int32(255)
	v106 = v101
	goto L22
L26:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v83 == int32(0) {
		v101 = v80
		v102 = v81
		goto L25
	} else {
		goto L28
	}
L27:
	;
	v101 = v95
	v102 = int32(0)
	goto L25
L28:
	;
	v87 = v81 & int32(255)
	if v87 == v83 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v94 = int32(1)
	v95 = v80 + v94
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	if v96 != 0 {
		v79 = v79 + v94
		v80 = v95
		v81 = v96
		goto L26
	} else {
		goto L32
	}
L30:
	;
	v89 = F_tolower(m, v87)
	mBase = m.M
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v91 = F_tolower(m, v90)
	mBase = m.M
	if v89 == v91 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v101 = v80
	v102 = v93
	goto L25
L32:
	;
	goto L27
L33:
	;
	goto L18
}
func F_ACLGetUserByName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v3
	v11 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v13 = v6 + int32(12)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if l1 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	m.G0 = v6 + int32(16)
	return v209
L2:
	;
	if v166 != l1 {
		goto L29
	} else {
		goto L30
	}
L3:
	;
	v157 = int32(0)
	v163 = v22
	v164 = v23
	v166 = v157
	v170 = v157
	goto L2
L4:
	;
	if base.Ui32(v23) < base.Ui32(int32(8)) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v34 = v22
	v35 = v23
	v37 = int32(0)
	goto L7
L6:
	;
	v163 = v147
	v164 = v148
	v166 = v150
	v170 = base.B2i32(v153 != int32(0))
	goto L2
L7:
	;
	v43 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
	v44 = int32(4)
	v45 = v34 + v44
	if v35&v44 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v147 = v138
	v148 = v139
	v150 = v123
	v153 = v128
	goto L6
L9:
	;
	v128 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v45+v43+(v128-v43)&int32(3)+v116<<(uint(int32(2))%32))))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if base.Ui32(v139) < base.Ui32(int32(8)) {
		v147 = v138
		v148 = v139
		v150 = v123
		v153 = v128
		goto L6
	} else {
		goto L27
	}
L10:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v37))))
	v94 = int32(0)
	goto L21
L11:
	;
	v50 = int32(0)
	if base.Ui32(l1) <= base.Ui32(v37) {
		v83 = v37
		v86 = v50
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v86 == v43 {
		v116 = v50
		v123 = v83
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v60 = v37
	v63 = v50
	goto L14
L14:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v63))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v60))))
	if v66 != v68 {
		v83 = v60
		v86 = v63
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v83 = v71
	v86 = v73
	goto L12
L16:
	;
	v70 = int32(1)
	v71 = v60 + v70
	v73 = v63 + v70
	if base.Ui32(v43) <= base.Ui32(v73) {
		v83 = v71
		v86 = v73
		goto L12
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(v71) < base.Ui32(l1) {
		v60 = v71
		v63 = v73
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v147 = v34
	v148 = v35
	v150 = v83
	v153 = v86
	goto L6
L20:
	;
	if v94 != v43 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v94))))
	if v107 == v91&int32(255) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v109 = int32(1)
	v111 = v94 + v109
	if v111 != v43 {
		v94 = v111
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v163 = v34
	v164 = v35
	v166 = v37
	v170 = v109
	goto L2
L25:
	;
	v116 = v94
	v123 = v37 + int32(1)
	goto L9
L26:
	;
	v147 = v34
	v148 = v35
	v150 = v37
	v153 = v43
	goto L6
L27:
	;
	if base.Ui32(v123) < base.Ui32(l1) {
		v34 = v138
		v35 = v139
		v37 = v123
		goto L7
	} else {
		goto L28
	}
L28:
	;
	goto L8
L29:
	;
	goto L1
L30:
	;
	if v164&int32(1) == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v178 = v164 & int32(4)
	if v170&base.B2i32(v178 != int32(0)) != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v13 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	if v164&int32(2) != 0 {
		v204 = int32(0)
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v204
	goto L29
L35:
	;
	v188 = int32(3)
	v189 = int32(base.Ui32(v164) >> (uint(v188) % 32))
	if v178 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v199 = int32(4)
	goto L38
L37:
	;
	v199 = v189 << (uint(int32(2)) % 32)
	goto L38
L38:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v163+v189+(int32(0)-v189)&v188+v199+int32(4))))
	v204 = v203
	goto L34
}
func F_ACLListDupKeyPattern(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_sdsdup(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = F_valkey_malloc(m, int32(8))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v9
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v5
			return v11
		}
	}
}
func F_ACLListDupSds(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_sdsdup(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_ACLListMatchLoadedUser(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = int32(0)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3+int32(-1)))))
	switch v11 & int32(7) {
	case 0:
		v28 = int32(base.Ui32(v11) >> (uint(int32(3)) % 32))
	case 1:
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3+int32(-3)))))
		v28 = v18
	case 2:
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3+int32(-5)))))
		v28 = v21
	case 3:
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v3+int32(-9))))
		v28 = v24
	case 4:
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v3+int32(-17))))
		v28 = v27
	default:
		v28 = v4
	}
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v31 & int32(7) {
	case 0:
		v48 = int32(base.Ui32(v31) >> (uint(int32(3)) % 32))
	case 1:
		v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
		v48 = v38
	case 2:
		v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
		v48 = v41
	case 3:
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
		v48 = v44
	case 4:
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
		v48 = v47
	default:
		v48 = v4
	}
	v49 = base.B2i32(base.Ui32(v28) < base.Ui32(v48))
	if base.Ui32(v28) < base.Ui32(v48) {
		v50 = v28
	} else {
		v50 = v48
	}
	v51 = F_memcmp(m, v3, l1, v50)
	mBase = m.M
	if v51 != 0 {
		v54 = v51
	} else {
		v54 = base.B2i32(base.Ui32(v48) < base.Ui32(v28)) - v49
	}
	return base.B2i32(v54 == int32(0))
}
func F_ACLListMatchSds(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	v3 = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v10 & int32(7) {
	case 0:
		v27 = int32(base.Ui32(v10) >> (uint(int32(3)) % 32))
	case 1:
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v27 = v17
	case 2:
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v27 = v20
	case 3:
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v27 = v23
	case 4:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v27 = v26
	default:
		v27 = v3
	}
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v30 & int32(7) {
	case 0:
		v47 = int32(base.Ui32(v30) >> (uint(int32(3)) % 32))
	case 1:
		v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
		v47 = v37
	case 2:
		v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
		v47 = v40
	case 3:
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
		v47 = v43
	case 4:
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
		v47 = v46
	default:
		v47 = v3
	}
	v48 = base.B2i32(base.Ui32(v27) < base.Ui32(v47))
	if base.Ui32(v27) < base.Ui32(v47) {
		v49 = v27
	} else {
		v49 = v47
	}
	v50 = F_memcmp(m, l0, l1, v49)
	mBase = m.M
	if v50 != 0 {
		v53 = v50
	} else {
		v53 = base.B2i32(base.Ui32(v47) < base.Ui32(v27)) - v48
	}
	return base.B2i32(v53 == int32(0))
}
func F_ACLModuleHasCommandRules(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int64
	_ = v25
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(320)
	m.G0 = v13
	v17 = v13 + int32(16)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(128)
	v25 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+12)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v17)+296)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v17)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v13 + int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = v13 + int32(184)
	goto L1
L1:
	;
	v40 = int32(0)
	v42 = F_raxSeek(m, v13+int32(16), int32(_a4), v40, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	v48 = F_raxNext(m, v13+int32(16))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v13 + int32(320)
	return v285
L5:
	;
	F_raxStop(m, v13+int32(16))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L2
	} else {
		goto L67
	}
L6:
	;
	if v48 == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	goto L8
L8:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v65 = v13 + int32(8)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v66
	goto L10
L9:
	;
	goto L5
L10:
	;
	v71 = v13 + int32(8)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v73 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v266 = F_raxNext(m, v13+int32(16))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L2
	} else {
		goto L65
	}
L12:
	;
	if v73 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v73+base.B2i32(v76 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v82
	goto L13
L15:
	;
	v90 = v73
	goto L16
L16:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+148))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+int32(-1)))))
	switch v100 & int32(7) {
	case 0:
		goto L24
	case 1:
		goto L23
	case 2:
		goto L22
	case 3:
		goto L21
	case 4:
		goto L20
	default:
		goto L18
	}
L17:
	;
	goto L11
L18:
	;
	v241 = v13 + int32(8)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	if v243 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L19:
	;
	if v117 == int32(0) {
		goto L18
	} else {
		goto L25
	}
L20:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v97+int32(-17))))
	v117 = v116
	goto L19
L21:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v97+int32(-9))))
	v117 = v113
	goto L19
L22:
	;
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97+int32(-5)))))
	v117 = v110
	goto L19
L23:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+int32(-3)))))
	v117 = v107
	goto L19
L24:
	;
	v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
	goto L19
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(0)
	v124 = F_sdssplitargs(m, v97, v13+int32(4))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	if v124 == int32(0) {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v129 < int32(1) {
		v223 = v129
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_sdsfreesplitres(m, v124, v223)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L2
	} else {
		goto L60
	}
L29:
	;
	v136 = int32(0)
	goto L30
L30:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v124+v136<<(uint(int32(2))%32))))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	switch v146 + int32(-43) {
	case 0, 2:
		goto L33
	default:
		goto L32
	}
L31:
	;
	v223 = v216
	goto L28
L32:
	;
	v215 = v136 + int32(1)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v215 < v216 {
		v136 = v215
		goto L30
	} else {
		goto L59
	}
L33:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+1)))
	if v149 == int32(64) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v153 = v145 + int32(1)
	v154 = F_sdsnew(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v158 = F_lookupCommandBySdsLogic(m, v157, v154)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_sdsfree(m, v154)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	if v158 != 0 {
		v189 = v158
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+56)))
	if v190&int32(8) == int32(0) {
		goto L32
	} else {
		goto L51
	}
L39:
	;
	v162 = int32(124)
	v163 = F___strchrnul(m, v153, v162)
	mBase = m.M
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v165 == v162 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v169 == int32(0) {
		goto L32
	} else {
		goto L44
	}
L41:
	;
	v169 = v163
	goto L43
L42:
	;
	v169 = int32(0)
	goto L43
L43:
	;
	goto L40
L44:
	;
	v173 = F_sdsnewlen(m, v153, v169-v153)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v175 = F_sdsnew(m, v173)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v179 = F_lookupCommandBySdsLogic(m, v178, v175)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	F_sdsfree(m, v175)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	F_sdsfree(m, v173)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	if v179 == int32(0) {
		goto L32
	} else {
		goto L50
	}
L50:
	;
	v189 = v179
	goto L38
L51:
	;
	v195 = F_moduleFromCommand(m, v189)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	if v195 != l0 {
		goto L32
	} else {
		goto L53
	}
L53:
	;
	if l1 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	F_sdsfreesplitres(m, v124, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L2
	} else {
		goto L57
	}
L55:
	;
	v200 = F_sdsdup(m, v145)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v200
	goto L54
L57:
	;
	F_raxStop(m, v13+int32(16))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	v285 = int32(1)
	goto L4
L59:
	;
	goto L31
L60:
	;
	goto L18
L61:
	;
	if v243 != 0 {
		v90 = v243
		goto L16
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v243+base.B2i32(v246 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = v252
	goto L62
L64:
	;
	goto L17
L65:
	;
	if v266 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	goto L9
L67:
	;
	v285 = v3
	goto L4
}
func F_ACLRecomputeCommandBitsFromCommandRulesAllUsers(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int64
	_ = v21
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	v8 = m.G0
	v10 = v8 - int32(320)
	m.G0 = v10
	v13 = v10 + int32(16)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(128)
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+12)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v13)+296)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v13)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v10 + int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+156)) = v10 + int32(184)
	goto L1
L1:
	;
	v36 = int32(0)
	v38 = F_raxSeek(m, v10+int32(16), int32(_a4), v36, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	v42 = F_raxNext(m, v10+int32(16))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L7
	}
L4:
	;
	F__serverAssert(m, int32(_a5), int32(_a6), int32(697))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L53
	}
L5:
	;
	F__serverAssert(m, int32(_a7), int32(_a6), int32(704))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L2
	} else {
		goto L52
	}
L6:
	;
	F_raxStop(m, v10+int32(16))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L2
	} else {
		goto L51
	}
L7:
	;
	if v42 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v56 = v10 + int32(8)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v57
	goto L11
L10:
	;
	goto L6
L11:
	;
	v62 = v10 + int32(8)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v64 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v195 = F_raxNext(m, v10+int32(16))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L49
	}
L13:
	;
	if v64 == int32(0) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v64+base.B2i32(v67 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v73
	goto L14
L16:
	;
	v78 = v64
	goto L17
L17:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+148))
	v90 = F_sdssplitargs(m, v87, v10+int32(4))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L19
	}
L18:
	;
	goto L12
L19:
	;
	if v90 == int32(0) {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v84)+128))
	if int64(-1) < v94 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v113 = int32(0)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v114 <= v113 {
		v167 = v114
		goto L29
	} else {
		goto L30
	}
L22:
	;
	v111 = F_ACLSetSelector(m, v84, int32(_a8), int32(-1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L27
	}
L23:
	;
	v99 = F_ACLSetSelector(m, v84, int32(_a9), int32(-1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	if v99 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	F__serverAssert(m, int32(_a7), int32(_a6), int32(701))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	if v111 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	goto L21
L29:
	;
	F_sdsfreesplitres(m, v90, v167)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L44
	}
L30:
	;
	v118 = v113
	goto L31
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v90+v118<<(uint(int32(2))%32))))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))))
	switch v131 & int32(7) {
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
		v148 = int32(0)
		goto L33
	}
L32:
	;
	v167 = v161
	goto L29
L33:
	;
	v149 = F_ACLSetSelector(m, v84, v128, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L40
	}
L34:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v128+int32(-17))))
	v148 = v147
	goto L33
L35:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v128+int32(-9))))
	v148 = v144
	goto L33
L36:
	;
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128+int32(-5)))))
	v148 = v141
	goto L33
L37:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))))
	v148 = v138
	goto L33
L38:
	;
	v148 = int32(base.Ui32(v131) >> (uint(int32(3)) % 32))
	goto L33
L39:
	;
	v160 = v118 + int32(1)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v160 < v161 {
		v118 = v160
		goto L31
	} else {
		goto L43
	}
L40:
	;
	if v149 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	F__serverAssert(m, int32(_a7), int32(_a6), int32(710))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	goto L32
L44:
	;
	v173 = v10 + int32(8)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	if v175 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v175 != 0 {
		v78 = v175
		goto L17
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v175+base.B2i32(v178 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v184
	goto L46
L48:
	;
	goto L18
L49:
	;
	if v195 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	goto L10
L51:
	;
	m.G0 = v10 + int32(320)
	return
L52:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ACLSelectorCheckCmd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v124 int32
	_ = v124
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int64
	_ = v173
	var v175 int32
	_ = v175
	var v178 int64
	_ = v178
	var v184 int64
	_ = v184
	var v191 int64
	_ = v191
	var v198 int64
	_ = v198
	var v202 int64
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int64
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v227 int64
	_ = v227
	var v230 int64
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v241 int64
	_ = v241
	var v249 int32
	_ = v249
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int64
	_ = v276
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	var v327 int64
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int64
	_ = v364
	var v367 int32
	_ = v367
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int64
	_ = v403
	var v405 int32
	_ = v405
	var v408 int64
	_ = v408
	var v414 int64
	_ = v414
	var v421 int64
	_ = v421
	var v428 int64
	_ = v428
	var v432 int64
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v443 int64
	_ = v443
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v452 int64
	_ = v452
	var v457 int64
	_ = v457
	var v460 int64
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v471 int64
	_ = v471
	var v479 int32
	_ = v479
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v506 int64
	_ = v506
	var v509 int32
	_ = v509
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v554 int32
	_ = v554
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v584 int64
	_ = v584
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v692 int32
	_ = v692
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v733 int64
	_ = v733
	var v735 int32
	_ = v735
	var v738 int64
	_ = v738
	var v744 int64
	_ = v744
	var v751 int64
	_ = v751
	var v758 int64
	_ = v758
	var v762 int64
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v773 int64
	_ = v773
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int64
	_ = v782
	var v787 int64
	_ = v787
	var v790 int64
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v801 int64
	_ = v801
	var v809 int32
	_ = v809
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v851 int32
	_ = v851
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v989 int32
	_ = v989
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1040 int64
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1082 int32
	_ = v1082
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1209 int32
	_ = v1209
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1242 int32
	_ = v1242
	v16 = m.G0
	v18 = v16 - int32(2064)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v18 + int32(2064)
	return v1242
L2:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v571&int32(4) != 0 {
		v692 = v571
		goto L126
	} else {
		goto L127
	}
L3:
	;
	F__serverAssert(m, int32(_a49), int32(_a6), int32(1863))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L8
	} else {
		goto L125
	}
L4:
	;
	v1242 = int32(1)
	goto L1
L5:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+59)))
	if v302&int32(32) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+2060)) = int32(0)
	v28 = m.T0[v21].(func(*base.Module, int32, int32, int32) int32)(m, l2, l3, v18+int32(2060))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v146 = int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v147 != 0 {
		v249 = v146
		goto L37
	} else {
		goto L38
	}
L8:
	;
	return int32(0)
L9:
	;
	if v28 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)+2060))
	if v34 < int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_valkey_free(m, v28)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L35
	}
L12:
	;
	v47 = int32(0)
	goto L13
L13:
	;
	v53 = int32(2)
	v55 = v28 + v47<<(uint(v53)%32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l2+v56<<(uint(v53)%32))))
	v61 = F_getLongLongFromObject(m, v60, v18)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L8
	} else {
		goto L15
	}
L14:
	;
	goto L11
L15:
	;
	if v61 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v63&int32(16) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v105 = v47 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v18)+2060))
	if v105 < v106 {
		v47 = v105
		goto L13
	} else {
		goto L34
	}
L18:
	;
	v66 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18))))
	if v66 < int64(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v98
	F_valkey_free(m, v28)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L33
	}
L20:
	;
	v70 = int64(*(*int32)(unsafe.Add(mBase, _consts[10])))
	if v70 <= v66 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v72 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if base.Ui64(v66+int64(-32768)) < base.Ui64(int64(-65536)) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	if v96 != 0 {
		goto L17
	} else {
		goto L32
	}
L24:
	;
	goto L23
L25:
	;
	v85 = int32(4)
	goto L27
L26:
	;
	v85 = int32(2)
	goto L27
L27:
	;
	if base.Ui64(v66+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v90 = int32(8)
	goto L30
L29:
	;
	v90 = v85
	goto L30
L30:
	;
	if base.Ui32(v77) < base.Ui32(v90) {
		v96 = int32(0)
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v92 = int32(0)
	v93 = F_intsetSearch(m, v72, v66, v92)
	mBase = m.M
	v96 = base.B2i32(v93 != v92)
	goto L24
L32:
	;
	goto L19
L33:
	;
	goto L4
L34:
	;
	goto L14
L35:
	;
	goto L7
L36:
	;
	if v262 == int32(0) {
		goto L2
	} else {
		goto L53
	}
L37:
	;
	v262 = v249
	goto L36
L38:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+58)))
	if v148&int32(32) != 0 {
		v249 = v146
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if int32(1) <= v151 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v156 = v151 & int32(3)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if base.Ui32(int32(4)) <= base.Ui32(v151) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v262 = int32(0)
	goto L36
L42:
	;
	if v156 == int32(0) {
		v241 = v213
		goto L48
	} else {
		goto L49
	}
L43:
	;
	v165 = int32(0)
	v168 = v165
	v170 = v165
	v173 = int64(0)
	goto L45
L44:
	;
	v208 = int32(0)
	v213 = int64(0)
	goto L42
L45:
	;
	v175 = int32(48)
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v157+v168*v175)+8))
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v157+(v168|int32(1))*v175)+8))
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v157+(v168|int32(2))*v175)+8))
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v157+(v168|int32(3))*v175)+8))
	v202 = v173 | (v178&v184&v191&v198 ^ int64(-1))
	v203 = int32(4)
	v204 = v168 + v203
	v206 = v170 + v203
	if v206 != v151&int32(2147483644) {
		v168 = v204
		v170 = v206
		v173 = v202
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v208 = v204
	v213 = v202
	goto L42
L47:
	;
	goto L46
L48:
	;
	v249 = int32(base.Ui32(base.I32_wrap_i64(v241))>>(uint(int32(8))%32)) & int32(1)
	goto L37
L49:
	;
	v217 = v208
	v221 = int32(0)
	v222 = v213
	goto L50
L50:
	;
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v157+v217*int32(48))+8))
	v230 = v222 | (v227 ^ int64(-1))
	v231 = int32(1)
	v234 = v221 + v231
	if v234 != v156 {
		v217 = v217 + v231
		v221 = v234
		v222 = v230
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v241 = v230
	goto L48
L52:
	;
	goto L51
L53:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v265&int32(16) != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	if l6 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L4
L56:
	;
	v271 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v271 <= l6 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v273 == int32(0) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v276 = base.I64_extend_i32_s(l6)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	if base.Ui64(v276+int64(-32768)) < base.Ui64(int64(-65536)) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	if v298 != 0 {
		goto L2
	} else {
		goto L68
	}
L60:
	;
	goto L59
L61:
	;
	v287 = int32(4)
	goto L63
L62:
	;
	v287 = int32(2)
	goto L63
L63:
	;
	if base.Ui64(v276+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v292 = int32(8)
	goto L66
L65:
	;
	v292 = v287
	goto L66
L66:
	;
	if base.Ui32(v279) < base.Ui32(v292) {
		v298 = int32(0)
		goto L60
	} else {
		goto L67
	}
L67:
	;
	v294 = int32(0)
	v295 = F_intsetSearch(m, v273, v276, v294)
	mBase = m.M
	v298 = base.B2i32(v295 != v294)
	goto L60
L68:
	;
	goto L55
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L4
L70:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)))
	if v367&int32(7) != 0 {
		goto L90
	} else {
		goto L91
	}
L71:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v307&int32(16) != 0 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v311 < int32(1) {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	v324 = v311
	v327 = int64(0)
	goto L74
L74:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v330&int32(16) != 0 {
		v362 = v324
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v364 = v327 + int64(1)
	if v364 < base.I64_extend_i32_s(v362) {
		v324 = v362
		v327 = v364
		goto L74
	} else {
		goto L89
	}
L77:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v333 == int32(0) {
		goto L69
	} else {
		goto L78
	}
L78:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	if base.Ui64(v327+int64(-32768)) < base.Ui64(int64(-65536)) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	if v357 == int32(0) {
		goto L69
	} else {
		goto L88
	}
L80:
	;
	goto L79
L81:
	;
	v346 = int32(4)
	goto L83
L82:
	;
	v346 = int32(2)
	goto L83
L83:
	;
	if base.Ui64(v327+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v351 = int32(8)
	goto L86
L85:
	;
	v351 = v346
	goto L86
L86:
	;
	if base.Ui32(v338) < base.Ui32(v351) {
		v357 = int32(0)
		goto L80
	} else {
		goto L87
	}
L87:
	;
	v353 = int32(0)
	v354 = F_intsetSearch(m, v333, v327, v353)
	mBase = m.M
	v357 = base.B2i32(v354 != v353)
	goto L80
L88:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v362 = v361
	goto L76
L89:
	;
	goto L2
L90:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v495&int32(16) != 0 {
		goto L2
	} else {
		goto L110
	}
L91:
	;
	v376 = int32(1)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v377 != 0 {
		v479 = v376
		goto L93
	} else {
		goto L94
	}
L92:
	;
	if v492 == int32(0) {
		goto L2
	} else {
		goto L109
	}
L93:
	;
	v492 = v479
	goto L92
L94:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+58)))
	if v378&int32(32) != 0 {
		v479 = v376
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if int32(1) <= v381 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v386 = v381 & int32(3)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if base.Ui32(int32(4)) <= base.Ui32(v381) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v492 = int32(0)
	goto L92
L98:
	;
	if v386 == int32(0) {
		v471 = v443
		goto L104
	} else {
		goto L105
	}
L99:
	;
	v395 = int32(0)
	v398 = v395
	v400 = v395
	v403 = int64(0)
	goto L101
L100:
	;
	v438 = int32(0)
	v443 = int64(0)
	goto L98
L101:
	;
	v405 = int32(48)
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v387+v398*v405)+8))
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v387+(v398|int32(1))*v405)+8))
	v421 = *(*int64)(unsafe.Add(mBase, uint32(v387+(v398|int32(2))*v405)+8))
	v428 = *(*int64)(unsafe.Add(mBase, uint32(v387+(v398|int32(3))*v405)+8))
	v432 = v403 | (v408&v414&v421&v428 ^ int64(-1))
	v433 = int32(4)
	v434 = v398 + v433
	v436 = v400 + v433
	if v436 != v381&int32(2147483644) {
		v398 = v434
		v400 = v436
		v403 = v432
		goto L101
	} else {
		goto L103
	}
L102:
	;
	v438 = v434
	v443 = v432
	goto L98
L103:
	;
	goto L102
L104:
	;
	v479 = int32(base.Ui32(base.I32_wrap_i64(v471))>>(uint(int32(8))%32)) & int32(1)
	goto L93
L105:
	;
	v447 = v438
	v451 = int32(0)
	v452 = v443
	goto L106
L106:
	;
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v387+v447*int32(48))+8))
	v460 = v452 | (v457 ^ int64(-1))
	v461 = int32(1)
	v464 = v451 + v461
	if v464 != v386 {
		v447 = v447 + v461
		v451 = v464
		v452 = v460
		goto L106
	} else {
		goto L108
	}
L107:
	;
	v471 = v460
	goto L104
L108:
	;
	goto L107
L109:
	;
	goto L90
L110:
	;
	if l6 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L4
L112:
	;
	v501 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v501 <= l6 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v503 == int32(0) {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	v506 = base.I64_extend_i32_s(l6)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	if base.Ui64(v506+int64(-32768)) < base.Ui64(int64(-65536)) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	if v528 != 0 {
		goto L2
	} else {
		goto L124
	}
L116:
	;
	goto L115
L117:
	;
	v517 = int32(4)
	goto L119
L118:
	;
	v517 = int32(2)
	goto L119
L119:
	;
	if base.Ui64(v506+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v522 = int32(8)
	goto L122
L121:
	;
	v522 = v517
	goto L122
L122:
	;
	if base.Ui32(v509) < base.Ui32(v522) {
		v528 = int32(0)
		goto L116
	} else {
		goto L123
	}
L123:
	;
	v524 = int32(0)
	v525 = F_intsetSearch(m, v503, v506, v524)
	mBase = m.M
	v528 = base.B2i32(v525 != v524)
	goto L116
L124:
	;
	goto L111
L125:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	if v692&int32(2) != 0 {
		goto L154
	} else {
		goto L155
	}
L127:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+57)))
	if v574&int32(128) != 0 {
		v692 = v571
		goto L126
	} else {
		goto L128
	}
L128:
	;
	if base.Ui32(int32(1023)) < base.Ui32(v20) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v592 = int32(2)
	if l3 < v592 {
		v1242 = v592
		goto L1
	} else {
		goto L132
	}
L130:
	;
	v584 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(base.Ui32(v20)>>(uint(int32(3))%32))&int32(536870904))+8))
	if base.I32_wrap_i64(int64(base.Ui64(v584)>>(uint(base.I64_extend_i32_u(v20&int32(63)))%64)))&int32(1) != 0 {
		v692 = v571
		goto L126
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v595 == int32(0) {
		v1242 = v592
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v598 = int32(2)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v595+v20<<(uint(v598)%32))))
	if v602 == int32(0) {
		v1242 = v598
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v616 = int32(0)
	v618 = v602
	goto L135
L135:
	;
	v623 = v616 << (uint(int32(2)) % 32)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v618+v623)))
	if v625 == int32(0) {
		v1242 = int32(2)
		goto L1
	} else {
		goto L137
	}
L136:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v692 = v682
	goto L126
L137:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	if v632 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v633 = int32(8)
	goto L140
L139:
	;
	v633 = int32(4)
	goto L140
L140:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l2+v633)))
	v636 = F_objectGetVal(m, v635)
	mBase = m.M
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v637+v20<<(uint(int32(2))%32))))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v641+v623)))
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636))))
	if v646 != 0 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	if v678-v680 != 0 {
		v616 = v616 + int32(1)
		v618 = v641
		goto L135
	} else {
		goto L153
	}
L142:
	;
	v678 = F_tolower(m, v674)
	mBase = m.M
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675))))
	v680 = F_tolower(m, v679)
	mBase = m.M
	goto L141
L143:
	;
	v648 = v636
	v649 = v643
	v650 = v646
	goto L146
L144:
	;
	v674 = int32(0)
	v675 = v643
	goto L142
L145:
	;
	v674 = v671 & int32(255)
	v675 = v670
	goto L142
L146:
	;
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649))))
	if v652 == int32(0) {
		v670 = v649
		v671 = v650
		goto L145
	} else {
		goto L148
	}
L147:
	;
	v670 = v664
	v671 = int32(0)
	goto L145
L148:
	;
	v656 = v650 & int32(255)
	if v656 == v652 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v663 = int32(1)
	v664 = v649 + v663
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v648)+1)))
	if v665 != 0 {
		v648 = v648 + v663
		v649 = v664
		v650 = v665
		goto L146
	} else {
		goto L152
	}
L150:
	;
	v658 = F_tolower(m, v656)
	mBase = m.M
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649))))
	v660 = F_tolower(m, v659)
	mBase = m.M
	if v658 == v660 {
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v648))))
	v670 = v649
	v671 = v662
	goto L145
L152:
	;
	goto L147
L153:
	;
	goto L136
L154:
	;
	v1015 = int32(0)
	v1016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v1016&int32(8) != 0 {
		v1242 = v1015
		goto L1
	} else {
		goto L212
	}
L155:
	;
	v706 = int32(1)
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v707 != 0 {
		v809 = v706
		goto L157
	} else {
		goto L158
	}
L156:
	;
	if v822 == int32(0) {
		goto L154
	} else {
		goto L173
	}
L157:
	;
	v822 = v809
	goto L156
L158:
	;
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+58)))
	if v708&int32(32) != 0 {
		v809 = v706
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if int32(1) <= v711 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v716 = v711 & int32(3)
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if base.Ui32(int32(4)) <= base.Ui32(v711) {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	v822 = int32(0)
	goto L156
L162:
	;
	if v716 == int32(0) {
		v801 = v773
		goto L168
	} else {
		goto L169
	}
L163:
	;
	v725 = int32(0)
	v728 = v725
	v730 = v725
	v733 = int64(0)
	goto L165
L164:
	;
	v768 = int32(0)
	v773 = int64(0)
	goto L162
L165:
	;
	v735 = int32(48)
	v738 = *(*int64)(unsafe.Add(mBase, uint32(v717+v728*v735)+8))
	v744 = *(*int64)(unsafe.Add(mBase, uint32(v717+(v728|int32(1))*v735)+8))
	v751 = *(*int64)(unsafe.Add(mBase, uint32(v717+(v728|int32(2))*v735)+8))
	v758 = *(*int64)(unsafe.Add(mBase, uint32(v717+(v728|int32(3))*v735)+8))
	v762 = v733 | (v738&v744&v751&v758 ^ int64(-1))
	v763 = int32(4)
	v764 = v728 + v763
	v766 = v730 + v763
	if v766 != v711&int32(2147483644) {
		v728 = v764
		v730 = v766
		v733 = v762
		goto L165
	} else {
		goto L167
	}
L166:
	;
	v768 = v764
	v773 = v762
	goto L162
L167:
	;
	goto L166
L168:
	;
	v809 = int32(base.Ui32(base.I32_wrap_i64(v801))>>(uint(int32(8))%32)) & int32(1)
	goto L157
L169:
	;
	v777 = v768
	v781 = int32(0)
	v782 = v773
	goto L170
L170:
	;
	v787 = *(*int64)(unsafe.Add(mBase, uint32(v717+v777*int32(48))+8))
	v790 = v782 | (v787 ^ int64(-1))
	v791 = int32(1)
	v794 = v781 + v791
	if v794 != v716 {
		v777 = v777 + v791
		v781 = v794
		v782 = v790
		goto L170
	} else {
		goto L172
	}
L171:
	;
	v801 = v790
	goto L168
L172:
	;
	goto L171
L173:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v825 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v837 = int32(0)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v838 <= v837 {
		goto L154
	} else {
		goto L177
	}
L175:
	;
	v826 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v826
	*(*int64)(unsafe.Add(mBase, uint32(l5)+4)) = int64(1099511627776)
	v833 = F_getKeysFromCommandWithSpecs(m, l1, l2, l3, v826, l5+int32(4))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L8
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(1)
	goto L174
L177:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v851 = v837
	goto L178
L178:
	;
	v859 = v841 + v851<<(uint(int32(3))%32)
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v859)))
	v863 = l2 + v860<<(uint(int32(2))%32)
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v863)))
	v865 = F_objectGetVal(m, v864)
	mBase = m.M
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v863)))
	v868 = F_objectGetVal(m, v867)
	mBase = m.M
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v868+int32(-1)))))
	switch v871 & int32(7) {
	case 0:
		goto L185
	case 1:
		goto L184
	case 2:
		goto L183
	case 3:
		goto L182
	case 4:
		goto L181
	default:
		v888 = int32(0)
		goto L180
	}
L179:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v859)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v998
	v1242 = v989
	goto L1
L180:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v859)+4))
	v894 = m.G0
	v896 = v894 - int32(16)
	m.G0 = v896
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v899&int32(2) != 0 {
		v989 = int32(0)
		goto L188
	} else {
		goto L189
	}
L181:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v868+int32(-17))))
	v888 = v887
	goto L180
L182:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v868+int32(-9))))
	v888 = v884
	goto L180
L183:
	;
	v881 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v868+int32(-5)))))
	v888 = v881
	goto L180
L184:
	;
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v868+int32(-3)))))
	v888 = v878
	goto L180
L185:
	;
	v888 = int32(base.Ui32(v871) >> (uint(int32(3)) % 32))
	goto L180
L186:
	;
	goto L179
L187:
	;
	if v989 != 0 {
		goto L186
	} else {
		goto L210
	}
L188:
	;
	m.G0 = v896 + int32(16)
	goto L187
L189:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v904 = v896 + int32(8)
	F_listRewind(m, v902, v904)
	mBase = m.M
	v908 = F_listNext(m, v904)
	mBase = m.M
	if v908 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v989 = int32(3)
	goto L188
L191:
	;
	v911 = int32(2)
	if v889&int32(160) != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v918 = v911
	goto L194
L193:
	;
	v918 = int32(base.Ui32(v889)>>(uint(int32(5))%32)) & v911
	goto L194
L194:
	;
	v923 = v918 | int32(base.Ui32(v889)>>(uint(int32(4))%32))&int32(1)
	v924 = v908
	goto L195
L195:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v924)+8))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v932)))
	if v933&v923 != v923 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	goto L190
L197:
	;
	v973 = F_listNext(m, v896+int32(8))
	mBase = m.M
	if v973 != 0 {
		v924 = v973
		goto L195
	} else {
		goto L209
	}
L198:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v932)+4))
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937+int32(-1)))))
	switch v940 & int32(7) {
	case 0:
		goto L204
	case 1:
		goto L203
	case 2:
		goto L202
	case 3:
		goto L201
	case 4:
		goto L200
	default:
		v957 = int32(0)
		goto L199
	}
L199:
	;
	goto L205
L200:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v937+int32(-17))))
	v957 = v956
	goto L199
L201:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v937+int32(-9))))
	v957 = v953
	goto L199
L202:
	;
	v950 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v937+int32(-5)))))
	v957 = v950
	goto L199
L203:
	;
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937+int32(-3)))))
	v957 = v947
	goto L199
L204:
	;
	v957 = int32(base.Ui32(v940) >> (uint(int32(3)) % 32))
	goto L199
L205:
	;
	v965 = int32(0)
	v967 = F_stringmatchlen(m, v937, v957, v865, v888, v965)
	mBase = m.M
	if v967 != 0 {
		v989 = v965
		goto L188
	} else {
		goto L208
	}
L208:
	;
	goto L197
L209:
	;
	goto L196
L210:
	;
	v995 = v851 + int32(1)
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v996 <= v995 {
		goto L154
	} else {
		goto L211
	}
L211:
	;
	v851 = v995
	goto L178
L212:
	;
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+59)))
	if v1024&int32(8) != 0 {
		v1050 = int32(1)
		goto L214
	} else {
		goto L215
	}
L213:
	;
	if v1058 == int32(0) {
		v1242 = v1015
		goto L1
	} else {
		goto L222
	}
L214:
	;
	v1058 = v1050
	goto L213
L215:
	;
	v1027 = int32(0)
	v1029 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	if v1029 == v1027 {
		v1050 = v1027
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1034 = int32(_a50)
	v1037 = v1029
	goto L217
L217:
	;
	if v1032 != v1037 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v1050 = v1027
	goto L214
L219:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+24))
	if v1045 != 0 {
		v1034 = v1034 + int32(24)
		v1037 = v1045
		goto L217
	} else {
		goto L221
	}
L220:
	;
	v1040 = *(*int64)(unsafe.Add(mBase, uint32(v1034)+8))
	v1058 = base.B2i32(v1040&base.I64_extend_i32_s(int32(20480)) != int64(0))
	goto L213
L221:
	;
	goto L218
L222:
	;
	v1061 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v1061
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(1099511627776)
	v1066 = F_getChannelsFromCommand(m, l1, l2, l3, v18)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L8
	} else {
		goto L223
	}
L223:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v1068 < int32(1) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1090)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1228
	F_getKeysFreeResult(m, v18)
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L8
	} else {
		goto L258
	}
L225:
	;
	F_getKeysFreeResult(m, v18)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L8
	} else {
		goto L257
	}
L226:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v1074 = v1068
	v1082 = int32(0)
	goto L227
L227:
	;
	v1090 = v1071 + v1082<<(uint(int32(3))%32)
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+4))
	if v1091&int32(20480) == int32(0) {
		v1202 = v1074
		goto L229
	} else {
		goto L230
	}
L228:
	;
	goto L225
L229:
	;
	v1209 = v1082 + int32(1)
	if v1209 < v1202 {
		v1074 = v1202
		v1082 = v1209
		goto L227
	} else {
		goto L256
	}
L230:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1090)))
	v1100 = l2 + v1097<<(uint(int32(2))%32)
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1100)))
	v1102 = F_objectGetVal(m, v1101)
	mBase = m.M
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1100)))
	v1107 = F_objectGetVal(m, v1106)
	mBase = m.M
	v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1107+int32(-1)))))
	switch v1110 & int32(7) {
	case 0:
		goto L236
	case 1:
		goto L235
	case 2:
		goto L234
	case 3:
		goto L233
	case 4:
		goto L232
	default:
		v1127 = int32(0)
		goto L231
	}
L231:
	;
	v1131 = m.G0
	v1133 = v1131 - int32(16)
	m.G0 = v1133
	v1136 = v1133 + int32(8)
	F_listRewind(m, v1096, v1136)
	mBase = m.M
	v1140 = F_listNext(m, v1136)
	mBase = m.M
	if v1140 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L232:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1107+int32(-17))))
	v1127 = v1126
	goto L231
L233:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1107+int32(-9))))
	v1127 = v1123
	goto L231
L234:
	;
	v1120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1107+int32(-5)))))
	v1127 = v1120
	goto L231
L235:
	;
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1107+int32(-3)))))
	v1127 = v1117
	goto L231
L236:
	;
	v1127 = int32(base.Ui32(v1110) >> (uint(int32(3)) % 32))
	goto L231
L237:
	;
	if v1197 != 0 {
		goto L224
	} else {
		goto L255
	}
L238:
	;
	m.G0 = v1133 + int32(16)
	goto L237
L239:
	;
	v1197 = int32(5)
	goto L238
L240:
	;
	v1143 = v1140
	goto L241
L241:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+8))
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151+int32(-1)))))
	switch v1154 & int32(7) {
	case 0:
		goto L248
	case 1:
		goto L247
	case 2:
		goto L246
	case 3:
		goto L245
	case 4:
		goto L244
	default:
		v1171 = int32(0)
		goto L243
	}
L242:
	;
	goto L239
L243:
	;
	if v1091&int32(2048) == int32(0) {
		goto L250
	} else {
		goto L251
	}
L244:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1151+int32(-17))))
	v1171 = v1170
	goto L243
L245:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1151+int32(-9))))
	v1171 = v1167
	goto L243
L246:
	;
	v1164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1151+int32(-5)))))
	v1171 = v1164
	goto L243
L247:
	;
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151+int32(-3)))))
	v1171 = v1161
	goto L243
L248:
	;
	v1171 = int32(base.Ui32(v1154) >> (uint(int32(3)) % 32))
	goto L243
L249:
	;
	v1182 = F_listNext(m, v1133+int32(8))
	mBase = m.M
	if v1182 != 0 {
		v1143 = v1182
		goto L241
	} else {
		goto L254
	}
L250:
	;
	v1176 = int32(0)
	v1178 = F_stringmatchlen(m, v1151, v1171, v1102, v1127, v1176)
	mBase = m.M
	if v1178 != 0 {
		v1197 = v1176
		goto L238
	} else {
		goto L253
	}
L251:
	;
	v1174 = F_strcmp(m, v1151, v1102)
	mBase = m.M
	if v1174 != 0 {
		goto L249
	} else {
		goto L252
	}
L252:
	;
	v1197 = int32(0)
	goto L238
L253:
	;
	goto L249
L254:
	;
	goto L242
L255:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v1202 = v1201
	goto L229
L256:
	;
	goto L228
L257:
	;
	v1242 = v1061
	goto L1
L258:
	;
	v1242 = v1197
	goto L1
}
func F_ACLSelectorCheckKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v173 int32
	_ = v173
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v14&int32(2) != 0 {
		v173 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v173
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v19 = v11 + int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v20
	goto L3
L3:
	;
	v25 = v11 + int32(8)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v27 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v173 = int32(3)
	goto L1
L5:
	;
	if v27 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27+base.B2i32(v30 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v36
	goto L6
L8:
	;
	v40 = int32(2)
	if l3&int32(160) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v47 = v40
	goto L11
L10:
	;
	v47 = int32(base.Ui32(l3)>>(uint(int32(5))%32)) & v40
	goto L11
L11:
	;
	v52 = v47 | int32(base.Ui32(l3)>>(uint(int32(4))%32))&int32(1)
	v53 = v27
	goto L12
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v62&v52 != v52 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L4
L14:
	;
	v145 = v11 + int32(8)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	if v147 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+int32(-1)))))
	switch v69 & int32(7) {
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
		v86 = int32(0)
		goto L16
	}
L16:
	;
	if l4 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v66+int32(-17))))
	v86 = v85
	goto L16
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v66+int32(-9))))
	v86 = v82
	goto L16
L19:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66+int32(-5)))))
	v86 = v79
	goto L16
L20:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+int32(-3)))))
	v86 = v76
	goto L16
L21:
	;
	v86 = int32(base.Ui32(v69) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	v125 = int32(0)
	v128 = m.G0
	v129 = int32(16)
	v130 = v128 - v129
	m.G0 = v130
	*(*int32)(unsafe.Add(mBase, uint32(v130)+12)) = v125
	v137 = F_stringmatchlen_impl(m, v66, v86, l1, l2, v125, v130+int32(12), v125)
	mBase = m.M
	m.G0 = v130 + v129
	goto L34
L23:
	;
	v89 = int32(0)
	v93 = m.G0
	v95 = v93 - int32(16)
	m.G0 = v95
	v97 = int32(1)
	if v86 != v97 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if v119 == int32(0) {
		goto L14
	} else {
		goto L33
	}
L25:
	;
	m.G0 = v95 + int32(16)
	goto L24
L26:
	;
	v113 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = v113
	v118 = F_stringmatchlen_impl(m, v66, v86, l1, l2, v89, v95+int32(12), v113)
	mBase = m.M
	v119 = v118
	goto L25
L27:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v86+int32(-1)))))
	if v109 != int32(42) {
		v119 = int32(0)
		goto L25
	} else {
		goto L32
	}
L28:
	;
	if v86 < int32(1) {
		goto L26
	} else {
		goto L31
	}
L29:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v100 != int32(42) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v119 = v97
	goto L25
L31:
	;
	goto L27
L32:
	;
	goto L26
L33:
	;
	v173 = v89
	goto L1
L34:
	;
	if v137 != 0 {
		v173 = v125
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L14
L36:
	;
	if v147 != 0 {
		v53 = v147
		goto L12
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v147+base.B2i32(v150 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v156
	goto L37
L39:
	;
	goto L13
}
func F_ACLSetSelectorCategory(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
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
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int64
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	v4 = int32(0)
	v11 = int32(-1)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	if v14 == int64(0) {
		v138 = v11
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v138
L2:
	;
	v18 = l1 + int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_ACLUpdateCommandRules(m, l0, l1, l2)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L33
	} else {
		goto L34
	}
L4:
	;
	if v54-v56 == int32(0) {
		v122 = v14
		goto L3
	} else {
		goto L16
	}
L5:
	;
	v54 = F_tolower(m, v50)
	mBase = m.M
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v56 = F_tolower(m, v55)
	mBase = m.M
	goto L4
L6:
	;
	v24 = v18
	v25 = v19
	v26 = v22
	goto L9
L7:
	;
	v50 = int32(0)
	v51 = v19
	goto L5
L8:
	;
	v50 = v47 & int32(255)
	v51 = v46
	goto L5
L9:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v28 == int32(0) {
		v46 = v25
		v47 = v26
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v46 = v40
	v47 = int32(0)
	goto L8
L11:
	;
	v32 = v26 & int32(255)
	if v32 == v28 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v39 = int32(1)
	v40 = v25 + v39
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v41 != 0 {
		v24 = v24 + v39
		v25 = v40
		v26 = v41
		goto L9
	} else {
		goto L15
	}
L13:
	;
	v34 = F_tolower(m, v32)
	mBase = m.M
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v36 = F_tolower(m, v35)
	mBase = m.M
	if v34 == v36 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v46 = v25
	v47 = v38
	goto L8
L15:
	;
	goto L10
L16:
	;
	v63 = v4
	goto L17
L17:
	;
	v70 = v63 + int32(1)
	v73 = v13 + v70<<(uint(int32(4))%32)
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v73)+8))
	if v74 == int64(0) {
		v138 = v11
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v122 = v74
	goto L3
L19:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v80 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v112-v114 != 0 {
		v63 = v70
		goto L17
	} else {
		goto L32
	}
L21:
	;
	v112 = F_tolower(m, v108)
	mBase = m.M
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v114 = F_tolower(m, v113)
	mBase = m.M
	goto L20
L22:
	;
	v82 = v18
	v83 = v77
	v84 = v80
	goto L25
L23:
	;
	v108 = int32(0)
	v109 = v77
	goto L21
L24:
	;
	v108 = v105 & int32(255)
	v109 = v104
	goto L21
L25:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v86 == int32(0) {
		v104 = v83
		v105 = v84
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v104 = v98
	v105 = int32(0)
	goto L24
L27:
	;
	v90 = v84 & int32(255)
	if v90 == v86 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v97 = int32(1)
	v98 = v83 + v97
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v99 != 0 {
		v82 = v82 + v97
		v83 = v98
		v84 = v99
		goto L25
	} else {
		goto L31
	}
L29:
	;
	v92 = F_tolower(m, v90)
	mBase = m.M
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v94 = F_tolower(m, v93)
	mBase = m.M
	if v92 == v94 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v104 = v83
	v105 = v96
	goto L24
L31:
	;
	goto L26
L32:
	;
	goto L18
L33:
	;
	return int32(0)
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_ACLSetSelectorCommandBitsForCategory(m, v130, l0, v122, l2)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v138 = int32(0)
	goto L1
}
func F_ACLSetUser(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
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
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
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
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
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
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v696 int32
	_ = v696
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v723 int32
	_ = v723
	var v729 int32
	_ = v729
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l2 != int32(-1) {
		v76 = l2
		goto L5
	} else {
		goto L6
	}
L2:
	;
	F_decrRefCount(m, v10)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L1
L5:
	;
	v77 = int32(0)
	if v76 == v77 {
		v696 = v77
		goto L31
	} else {
		goto L32
	}
L6:
	;
	if l1&int32(3) == int32(0) {
		v42 = l1
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v76 = v75
	goto L5
L8:
	;
	v75 = v67 - l1
	goto L7
L9:
	;
	v46 = v42
	goto L17
L10:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = l1
	goto L13
L12:
	;
	v75 = l1 - l1
	goto L7
L13:
	;
	v35 = v31 + int32(1)
	if v35&int32(3) == int32(0) {
		v42 = v35
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v40 != 0 {
		v31 = v35
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v67 = v35
	goto L8
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v55 = int32(-2139062144)
	if (int32(16843008)-v52|v52)&v55 == v55 {
		v46 = v46 + int32(4)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v61 = v46
	goto L20
L19:
	;
	goto L18
L20:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v65 != 0 {
		v61 = v61 + int32(1)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v67 = v61
	goto L8
L22:
	;
	goto L21
L23:
	;
	F__serverAssert(m, int32(_a17), int32(_a6), int32(1484))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L3
	} else {
		goto L238
	}
L24:
	;
	F__serverAssert(m, int32(_a18), int32(_a6), int32(1483))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L3
	} else {
		goto L237
	}
L25:
	;
	F__serverAssert(m, int32(_a19), int32(_a6), int32(1482))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L3
	} else {
		goto L236
	}
L26:
	;
	F__serverAssert(m, int32(_a20), int32(_a6), int32(1480))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L3
	} else {
		goto L235
	}
L27:
	;
	F__serverAssert(m, int32(_a21), int32(_a6), int32(1478))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L3
	} else {
		goto L234
	}
L28:
	;
	F__serverAssert(m, int32(_a22), int32(_a6), int32(1477))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L3
	} else {
		goto L233
	}
L29:
	;
	F__serverAssert(m, int32(_a23), int32(_a6), int32(1476))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L3
	} else {
		goto L232
	}
L30:
	;
	F__serverAssert(m, int32(_a24), int32(_a6), int32(1470))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L3
	} else {
		goto L231
	}
L31:
	;
	m.G0 = v8 + int32(16)
	return v696
L32:
	;
	v80 = int32(_a25)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v83 != 0 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v696 = int32(0)
	goto L31
L34:
	;
	v125 = int32(_a26)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v128 != 0 {
		goto L51
	} else {
		goto L52
	}
L35:
	;
	if v115-v117 != 0 {
		goto L34
	} else {
		goto L47
	}
L36:
	;
	v115 = F_tolower(m, v111)
	mBase = m.M
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v117 = F_tolower(m, v116)
	mBase = m.M
	goto L35
L37:
	;
	v85 = l1
	v86 = v80
	v87 = v83
	goto L40
L38:
	;
	v111 = int32(0)
	v112 = v80
	goto L36
L39:
	;
	v111 = v108 & int32(255)
	v112 = v107
	goto L36
L40:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v89 == int32(0) {
		v107 = v86
		v108 = v87
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v107 = v101
	v108 = int32(0)
	goto L39
L42:
	;
	v93 = v87 & int32(255)
	if v93 == v89 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v100 = int32(1)
	v101 = v86 + v100
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v102 != 0 {
		v85 = v85 + v100
		v86 = v101
		v87 = v102
		goto L40
	} else {
		goto L46
	}
L44:
	;
	v95 = F_tolower(m, v93)
	mBase = m.M
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v97 = F_tolower(m, v96)
	mBase = m.M
	if v95 == v97 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v107 = v86
	v108 = v99
	goto L39
L46:
	;
	goto L41
L47:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v119&int32(-4) | int32(1)
	goto L33
L48:
	;
	v170 = int32(_a27)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v173 != 0 {
		goto L64
	} else {
		goto L65
	}
L49:
	;
	if v160-v162 != 0 {
		goto L48
	} else {
		goto L61
	}
L50:
	;
	v160 = F_tolower(m, v156)
	mBase = m.M
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	v162 = F_tolower(m, v161)
	mBase = m.M
	goto L49
L51:
	;
	v130 = l1
	v131 = v125
	v132 = v128
	goto L54
L52:
	;
	v156 = int32(0)
	v157 = v125
	goto L50
L53:
	;
	v156 = v153 & int32(255)
	v157 = v152
	goto L50
L54:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v134 == int32(0) {
		v152 = v131
		v153 = v132
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v152 = v146
	v153 = int32(0)
	goto L53
L56:
	;
	v138 = v132 & int32(255)
	if v138 == v134 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v145 = int32(1)
	v146 = v131 + v145
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	if v147 != 0 {
		v130 = v130 + v145
		v131 = v146
		v132 = v147
		goto L54
	} else {
		goto L60
	}
L58:
	;
	v140 = F_tolower(m, v138)
	mBase = m.M
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	v142 = F_tolower(m, v141)
	mBase = m.M
	if v140 == v142 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	v152 = v131
	v153 = v144
	goto L53
L60:
	;
	goto L55
L61:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v164&int32(-4) | int32(2)
	goto L33
L62:
	;
	if v205-v207 == int32(0) {
		goto L33
	} else {
		goto L74
	}
L63:
	;
	v205 = F_tolower(m, v201)
	mBase = m.M
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	v207 = F_tolower(m, v206)
	mBase = m.M
	goto L62
L64:
	;
	v175 = l1
	v176 = v170
	v177 = v173
	goto L67
L65:
	;
	v201 = int32(0)
	v202 = v170
	goto L63
L66:
	;
	v201 = v198 & int32(255)
	v202 = v197
	goto L63
L67:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	if v179 == int32(0) {
		v197 = v176
		v198 = v177
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v197 = v191
	v198 = int32(0)
	goto L66
L69:
	;
	v183 = v177 & int32(255)
	if v183 == v179 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v190 = int32(1)
	v191 = v176 + v190
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	if v192 != 0 {
		v175 = v175 + v190
		v176 = v191
		v177 = v192
		goto L67
	} else {
		goto L73
	}
L71:
	;
	v185 = F_tolower(m, v183)
	mBase = m.M
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v187 = F_tolower(m, v186)
	mBase = m.M
	if v185 == v187 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v197 = v176
	v198 = v189
	goto L66
L73:
	;
	goto L68
L74:
	;
	v211 = int32(_a28)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v214 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	if v246-v248 == int32(0) {
		goto L33
	} else {
		goto L87
	}
L76:
	;
	v246 = F_tolower(m, v242)
	mBase = m.M
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v248 = F_tolower(m, v247)
	mBase = m.M
	goto L75
L77:
	;
	v216 = l1
	v217 = v211
	v218 = v214
	goto L80
L78:
	;
	v242 = int32(0)
	v243 = v211
	goto L76
L79:
	;
	v242 = v239 & int32(255)
	v243 = v238
	goto L76
L80:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v220 == int32(0) {
		v238 = v217
		v239 = v218
		goto L79
	} else {
		goto L82
	}
L81:
	;
	v238 = v232
	v239 = int32(0)
	goto L79
L82:
	;
	v224 = v218 & int32(255)
	if v224 == v220 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v231 = int32(1)
	v232 = v217 + v231
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	if v233 != 0 {
		v216 = v216 + v231
		v217 = v232
		v218 = v233
		goto L80
	} else {
		goto L86
	}
L84:
	;
	v226 = F_tolower(m, v224)
	mBase = m.M
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	v228 = F_tolower(m, v227)
	mBase = m.M
	if v226 == v228 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v238 = v217
	v239 = v230
	goto L79
L86:
	;
	goto L81
L87:
	;
	v252 = int32(_a29)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v255 != 0 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v298 = int32(_a30)
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v301 != 0 {
		goto L106
	} else {
		goto L107
	}
L89:
	;
	if v287-v289 != 0 {
		goto L88
	} else {
		goto L101
	}
L90:
	;
	v287 = F_tolower(m, v283)
	mBase = m.M
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	v289 = F_tolower(m, v288)
	mBase = m.M
	goto L89
L91:
	;
	v257 = l1
	v258 = v252
	v259 = v255
	goto L94
L92:
	;
	v283 = int32(0)
	v284 = v252
	goto L90
L93:
	;
	v283 = v280 & int32(255)
	v284 = v279
	goto L90
L94:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	if v261 == int32(0) {
		v279 = v258
		v280 = v259
		goto L93
	} else {
		goto L96
	}
L95:
	;
	v279 = v273
	v280 = int32(0)
	goto L93
L96:
	;
	v265 = v259 & int32(255)
	if v265 == v261 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v272 = int32(1)
	v273 = v258 + v272
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+1)))
	if v274 != 0 {
		v257 = v257 + v272
		v258 = v273
		v259 = v274
		goto L94
	} else {
		goto L100
	}
L98:
	;
	v267 = F_tolower(m, v265)
	mBase = m.M
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	v269 = F_tolower(m, v268)
	mBase = m.M
	if v267 == v269 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
	v279 = v258
	v280 = v271
	goto L93
L100:
	;
	goto L95
L101:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v291 | int32(4)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_listEmpty(m, v295)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L3
	} else {
		goto L102
	}
L102:
	;
	goto L33
L103:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	switch v344 + int32(-33) {
	case 0:
		goto L121
	default:
		goto L118
	case 2:
		goto L125
	case 7:
		goto L119
	case 27:
		goto L122
	case 29:
		goto L126
	}
L104:
	;
	if v333-v335 != 0 {
		goto L103
	} else {
		goto L116
	}
L105:
	;
	v333 = F_tolower(m, v329)
	mBase = m.M
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	v335 = F_tolower(m, v334)
	mBase = m.M
	goto L104
L106:
	;
	v303 = l1
	v304 = v298
	v305 = v301
	goto L109
L107:
	;
	v329 = int32(0)
	v330 = v298
	goto L105
L108:
	;
	v329 = v326 & int32(255)
	v330 = v325
	goto L105
L109:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	if v307 == int32(0) {
		v325 = v304
		v326 = v305
		goto L108
	} else {
		goto L111
	}
L110:
	;
	v325 = v319
	v326 = int32(0)
	goto L108
L111:
	;
	v311 = v305 & int32(255)
	if v311 == v307 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v318 = int32(1)
	v319 = v304 + v318
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+1)))
	if v320 != 0 {
		v303 = v303 + v318
		v304 = v319
		v305 = v320
		goto L109
	} else {
		goto L115
	}
L113:
	;
	v313 = F_tolower(m, v311)
	mBase = m.M
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	v315 = F_tolower(m, v314)
	mBase = m.M
	if v313 == v315 {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	v325 = v304
	v326 = v317
	goto L108
L115:
	;
	goto L110
L116:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v337 & int32(-5)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_listEmpty(m, v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L3
	} else {
		goto L117
	}
L117:
	;
	goto L33
L118:
	;
	v494 = int32(_a31)
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v497 != 0 {
		goto L168
	} else {
		goto L169
	}
L119:
	;
	v479 = int32(-1)
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v76+v479))))
	if v483 != int32(41) {
		goto L118
	} else {
		goto L161
	}
L120:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v468 = F_listSearchKey(m, v467, v464)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L3
	} else {
		goto L155
	}
L121:
	;
	if v76 != int32(65) {
		goto L145
	} else {
		goto L146
	}
L122:
	;
	v416 = F_ACLHashPassword(m, l1+int32(1), v76+int32(-1))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L3
	} else {
		goto L143
	}
L123:
	;
	goto L142
L124:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v392 = F_listSearchKey(m, v391, v388)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L3
	} else {
		goto L138
	}
L125:
	;
	if v76 != int32(65) {
		goto L123
	} else {
		goto L128
	}
L126:
	;
	v351 = F_ACLHashPassword(m, l1+int32(1), v76+int32(-1))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	v388 = v351
	goto L124
L128:
	;
	v356 = l1 + int32(1)
	v360 = int32(0)
	goto L130
L129:
	;
	if v377 != 0 {
		goto L123
	} else {
		goto L134
	}
L130:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356+v360))))
	v367 = int32(255)
	v377 = base.B2i32(base.Ui32((v364+int32(-103))&v367) < base.Ui32(int32(250))) & base.B2i32(base.Ui32((v364+int32(-58))&v367) < base.Ui32(int32(246)))
	if v377 != 0 {
		goto L129
	} else {
		goto L132
	}
L131:
	;
	goto L129
L132:
	;
	v379 = v360 + int32(1)
	if v379 != int32(64) {
		v360 = v379
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	v384 = F_sdsnewlen(m, v356, int32(64))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L3
	} else {
		goto L135
	}
L135:
	;
	v388 = v384
	goto L124
L136:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399 & int32(-5)
	goto L33
L137:
	;
	F_sdsfree(m, v388)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L3
	} else {
		goto L141
	}
L138:
	;
	if v392 != 0 {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v395 = F_listAddNodeTail(m, v394, v388)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L3
	} else {
		goto L140
	}
L140:
	;
	goto L136
L141:
	;
	goto L136
L142:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(9)
	v696 = int32(-1)
	goto L31
L143:
	;
	v464 = v416
	goto L120
L144:
	;
	v460 = F_sdsnewlen(m, v421, int32(64))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L3
	} else {
		goto L154
	}
L145:
	;
	goto L153
L146:
	;
	v421 = l1 + int32(1)
	v425 = int32(0)
	goto L148
L147:
	;
	if v442 == int32(0) {
		goto L144
	} else {
		goto L152
	}
L148:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421+v425))))
	v432 = int32(255)
	v442 = base.B2i32(base.Ui32((v429+int32(-103))&v432) < base.Ui32(int32(250))) & base.B2i32(base.Ui32((v429+int32(-58))&v432) < base.Ui32(int32(246)))
	if v442 != 0 {
		goto L147
	} else {
		goto L150
	}
L149:
	;
	goto L147
L150:
	;
	v444 = v425 + int32(1)
	if v444 != int32(64) {
		v425 = v444
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	goto L145
L153:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(9)
	v696 = int32(-1)
	goto L31
L154:
	;
	v464 = v460
	goto L120
L155:
	;
	F_sdsfree(m, v464)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L3
	} else {
		goto L156
	}
L156:
	;
	if v468 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_listDelNode(m, v476, v468)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L3
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(43)
	v696 = int32(-1)
	goto L31
L160:
	;
	goto L33
L161:
	;
	v486 = F_aclCreateSelectorFromOpSet(m, l1, v76)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	if v486 == int32(0) {
		v696 = v479
		goto L31
	} else {
		goto L163
	}
L163:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v491 = F_listAddNodeTail(m, v490, v486)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L3
	} else {
		goto L164
	}
L164:
	;
	goto L33
L165:
	;
	v594 = int32(_a32)
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v597 != 0 {
		goto L198
	} else {
		goto L199
	}
L166:
	;
	if v529-v531 != 0 {
		goto L165
	} else {
		goto L178
	}
L167:
	;
	v529 = F_tolower(m, v525)
	mBase = m.M
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	v531 = F_tolower(m, v530)
	mBase = m.M
	goto L166
L168:
	;
	v499 = l1
	v500 = v494
	v501 = v497
	goto L171
L169:
	;
	v525 = int32(0)
	v526 = v494
	goto L167
L170:
	;
	v525 = v522 & int32(255)
	v526 = v521
	goto L167
L171:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	if v503 == int32(0) {
		v521 = v500
		v522 = v501
		goto L170
	} else {
		goto L173
	}
L172:
	;
	v521 = v515
	v522 = int32(0)
	goto L170
L173:
	;
	v507 = v501 & int32(255)
	if v507 == v503 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v514 = int32(1)
	v515 = v500 + v514
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+1)))
	if v516 != 0 {
		v499 = v499 + v514
		v500 = v515
		v501 = v516
		goto L171
	} else {
		goto L177
	}
L175:
	;
	v509 = F_tolower(m, v507)
	mBase = m.M
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	v511 = F_tolower(m, v510)
	mBase = m.M
	if v509 == v511 {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499))))
	v521 = v500
	v522 = v513
	goto L170
L177:
	;
	goto L172
L178:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v535 = v8 + int32(8)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	*(*int32)(unsafe.Add(mBase, uint32(v535)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v535))) = v536
	goto L179
L179:
	;
	v541 = v8 + int32(8)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v541)))
	if v543 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	if v543 == int32(0) {
		goto L30
	} else {
		goto L183
	}
L181:
	;
	goto L180
L182:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v543+base.B2i32(v546 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v541))) = v552
	goto L181
L183:
	;
	v557 = v8 + int32(8)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v557)))
	if v559 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	if v559 == int32(0) {
		goto L33
	} else {
		goto L187
	}
L185:
	;
	goto L184
L186:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v557)+4))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v559+base.B2i32(v562 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v557))) = v568
	goto L185
L187:
	;
	v574 = v559
	goto L188
L188:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_listDelNode(m, v577, v574)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L3
	} else {
		goto L190
	}
L190:
	;
	v581 = v8 + int32(8)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	if v583 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	if v583 != 0 {
		v574 = v583
		goto L188
	} else {
		goto L194
	}
L192:
	;
	goto L191
L193:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v583+base.B2i32(v586 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v581))) = v592
	goto L192
L194:
	;
	goto L33
L195:
	;
	v680 = F_ACLUserGetRootSelector(m, l0)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L3
	} else {
		goto L228
	}
L196:
	;
	if v629-v631 != 0 {
		goto L195
	} else {
		goto L208
	}
L197:
	;
	v629 = F_tolower(m, v625)
	mBase = m.M
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	v631 = F_tolower(m, v630)
	mBase = m.M
	goto L196
L198:
	;
	v599 = l1
	v600 = v594
	v601 = v597
	goto L201
L199:
	;
	v625 = int32(0)
	v626 = v594
	goto L197
L200:
	;
	v625 = v622 & int32(255)
	v626 = v621
	goto L197
L201:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600))))
	if v603 == int32(0) {
		v621 = v600
		v622 = v601
		goto L200
	} else {
		goto L203
	}
L202:
	;
	v621 = v615
	v622 = int32(0)
	goto L200
L203:
	;
	v607 = v601 & int32(255)
	if v607 == v603 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v614 = int32(1)
	v615 = v600 + v614
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599)+1)))
	if v616 != 0 {
		v599 = v599 + v614
		v600 = v615
		v601 = v616
		goto L201
	} else {
		goto L207
	}
L205:
	;
	v609 = F_tolower(m, v607)
	mBase = m.M
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600))))
	v611 = F_tolower(m, v610)
	mBase = m.M
	if v609 == v611 {
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
	v621 = v600
	v622 = v613
	goto L200
L207:
	;
	goto L202
L208:
	;
	v635 = F_ACLSetUser(m, l0, int32(_a30), int32(-1))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L3
	} else {
		goto L209
	}
L209:
	;
	if v635 != 0 {
		goto L29
	} else {
		goto L210
	}
L210:
	;
	v639 = F_ACLSetUser(m, l0, int32(_a33), int32(-1))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L3
	} else {
		goto L211
	}
L211:
	;
	if v639 != 0 {
		goto L28
	} else {
		goto L212
	}
L212:
	;
	v643 = F_ACLSetUser(m, l0, int32(_a34), int32(-1))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L3
	} else {
		goto L213
	}
L213:
	;
	if v643 != 0 {
		goto L27
	} else {
		goto L214
	}
L214:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, _consts[7])))
	if v646&int32(8) == int32(0) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v657 = F_ACLSetUser(m, l0, int32(_a35), int32(-1))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L3
	} else {
		goto L219
	}
L216:
	;
	v653 = F_ACLSetUser(m, l0, int32(_a36), int32(-1))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L3
	} else {
		goto L217
	}
L217:
	;
	if v653 != 0 {
		goto L26
	} else {
		goto L218
	}
L218:
	;
	goto L215
L219:
	;
	if v657 != 0 {
		goto L25
	} else {
		goto L220
	}
L220:
	;
	v661 = F_ACLSetUser(m, l0, int32(_a26), int32(-1))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L3
	} else {
		goto L221
	}
L221:
	;
	if v661 != 0 {
		goto L24
	} else {
		goto L222
	}
L222:
	;
	v665 = F_ACLSetUser(m, l0, int32(_a31), int32(-1))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L3
	} else {
		goto L223
	}
L223:
	;
	if v665 != 0 {
		goto L23
	} else {
		goto L224
	}
L224:
	;
	v669 = F_ACLSetUser(m, l0, int32(_a8), int32(-1))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L3
	} else {
		goto L225
	}
L225:
	;
	if v669 == int32(0) {
		goto L33
	} else {
		goto L226
	}
L226:
	;
	F__serverAssert(m, int32(_a37), int32(_a6), int32(1485))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L3
	} else {
		goto L227
	}
L227:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	v682 = F_ACLSetSelector(m, v680, l1, v76)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L3
	} else {
		goto L229
	}
L229:
	;
	if v682 == int32(-1) {
		v696 = int32(-1)
		goto L31
	} else {
		goto L230
	}
L230:
	;
	goto L33
L231:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L233:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L238:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ACLSetUserStringError(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	switch v4 + int32(-7) {
	case 0:
		return int32(_a38)
	default:
		v24 = int32(_a39)
		return v24
	case 2:
		return int32(_a40)
	case 5:
		return int32(_a41)
	case 13:
		return int32(_a42)
	case 21:
		return int32(_a43)
	case 24:
		return int32(_a44)
	case 36:
		return int32(_a45)
	case 37:
		v24 = int32(_a46)
		return v24
	case 61:
		return int32(_a47)
	}
}
func F_ACLUpdateCommandRules(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v308 int32
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
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = F_sdsnew(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-1)))))
	switch v21 & int32(7) {
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
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-1)))))
	switch v60 & int32(7) {
	case 0:
		goto L20
	case 1:
		goto L19
	case 2:
		goto L18
	case 3:
		goto L17
	case 4:
		goto L16
	default:
		v77 = int32(0)
		goto L15
	}
L4:
	;
	goto L3
L5:
	;
	if v38 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-17))))
	v38 = v37
	goto L5
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-9))))
	v38 = v34
	goto L5
L8:
	;
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(-5)))))
	v38 = v31
	goto L5
L9:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-3)))))
	v38 = v28
	goto L5
L10:
	;
	v38 = int32(base.Ui32(v21) >> (uint(int32(3)) % 32))
	goto L5
L11:
	;
	v43 = int32(0)
	goto L12
L12:
	;
	v46 = v14 + v43
	v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v46))))
	v48 = F_tolower(m, v47)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v48)
	v51 = v43 + int32(1)
	if v51 != v38 {
		v43 = v51
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
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v79 == int32(0) {
		v455 = v78
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-17))))
	v77 = v76
	goto L15
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-9))))
	v77 = v73
	goto L15
L18:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(-5)))))
	v77 = v70
	goto L15
L19:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-3)))))
	v77 = v67
	goto L15
L20:
	;
	v77 = int32(base.Ui32(v60) >> (uint(int32(3)) % 32))
	goto L15
L21:
	;
	v460 = F_strlen(m, v455)
	mBase = m.M
	v462 = v455 + int32(-1)
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462))))
	switch v463 & int32(7) {
	case 0:
		goto L138
	case 1:
		goto L137
	case 2:
		goto L136
	case 3:
		goto L135
	case 4:
		goto L134
	default:
		goto L133
	}
L22:
	;
	v88 = v78
	goto L23
L23:
	;
	v92 = v88 + int32(1)
	v93 = int32(32)
	v94 = F___strchrnul(m, v92, v93)
	mBase = m.M
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if v96 == v93 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v455 = v448
	goto L21
L25:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v165 = v160 + base.B2i32(v162 == int32(32))
	v166 = v160 - v92
	if base.Ui32(v166) < base.Ui32(v77) {
		goto L49
	} else {
		goto L50
	}
L26:
	;
	if v100 != 0 {
		v160 = v100
		v161 = v88
		goto L25
	} else {
		goto L30
	}
L27:
	;
	v100 = v94
	goto L29
L28:
	;
	v100 = int32(0)
	goto L29
L29:
	;
	goto L26
L30:
	;
	if v92&int32(3) == int32(0) {
		v122 = v92
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v160 = v92 + v155
	v161 = v88 - base.B2i32(v88 != v157)
	goto L25
L32:
	;
	v155 = v147 - v92
	goto L31
L33:
	;
	v126 = v122
	goto L41
L34:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v108 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v111 = v92
	goto L37
L36:
	;
	v155 = v92 - v92
	goto L31
L37:
	;
	v115 = v111 + int32(1)
	if v115&int32(3) == int32(0) {
		v122 = v115
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v120 != 0 {
		v111 = v115
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v147 = v115
	goto L32
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v135 = int32(-2139062144)
	if (int32(16843008)-v132|v132)&v135 == v135 {
		v126 = v126 + int32(4)
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v141 = v126
	goto L44
L43:
	;
	goto L42
L44:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v145 != 0 {
		v141 = v141 + int32(1)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v147 = v141
	goto L32
L46:
	;
	goto L45
L47:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	if v447 != 0 {
		v88 = v446
		goto L23
	} else {
		goto L131
	}
L48:
	;
	if v166 == v77 {
		goto L69
	} else {
		goto L70
	}
L49:
	;
	v168 = v166
	goto L51
L50:
	;
	v168 = v77
	goto L51
L51:
	;
	if base.Ui32(v168) < base.Ui32(int32(4)) {
		v192 = v92
		v193 = v14
		v194 = v168
		goto L55
	} else {
		goto L56
	}
L52:
	;
	if v232 == int32(0) {
		goto L48
	} else {
		goto L68
	}
L53:
	;
	v232 = int32(0)
	goto L52
L54:
	;
	v204 = v199
	v205 = v200
	v206 = v201
	goto L64
L55:
	;
	if v194 == int32(0) {
		goto L53
	} else {
		goto L62
	}
L56:
	;
	if (v14|v92)&int32(3) != 0 {
		v199 = v92
		v200 = v14
		v201 = v168
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v176 = v92
	v177 = v14
	v178 = v168
	goto L58
L58:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	if v181 != v182 {
		v199 = v176
		v200 = v177
		v201 = v178
		goto L54
	} else {
		goto L60
	}
L59:
	;
	v192 = v187
	v193 = v185
	v194 = v189
	goto L55
L60:
	;
	v184 = int32(4)
	v185 = v177 + v184
	v187 = v176 + v184
	v189 = v178 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v189) {
		v176 = v187
		v177 = v185
		v178 = v189
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v199 = v192
	v200 = v193
	v201 = v194
	goto L54
L63:
	;
	v232 = v209 - v210
	goto L52
L64:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v209 != v210 {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v212 = int32(1)
	v217 = v206 + int32(-1)
	if v217 == int32(0) {
		goto L53
	} else {
		goto L67
	}
L67:
	;
	v204 = v204 + v212
	v205 = v205 + v212
	v206 = v217
	goto L64
L68:
	;
	v446 = v165
	goto L47
L69:
	;
	if v165&int32(3) == int32(0) {
		v262 = v165
		goto L76
	} else {
		goto L77
	}
L70:
	;
	if base.Ui32(v77) < base.Ui32(v166) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v77))))
	if v238 == int32(124) {
		goto L69
	} else {
		goto L73
	}
L72:
	;
	v446 = v165
	goto L47
L73:
	;
	v446 = v165
	goto L47
L74:
	;
	v297 = v295 + int32(1)
	if v161 == v165 {
		goto L91
	} else {
		goto L92
	}
L75:
	;
	v295 = v287 - v165
	goto L74
L76:
	;
	v266 = v262
	goto L84
L77:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v248 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v251 = v165
	goto L80
L79:
	;
	v295 = v165 - v165
	goto L74
L80:
	;
	v255 = v251 + int32(1)
	if v255&int32(3) == int32(0) {
		v262 = v255
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if v260 != 0 {
		v251 = v255
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v287 = v255
	goto L75
L84:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v275 = int32(-2139062144)
	if (int32(16843008)-v272|v272)&v275 == v275 {
		v266 = v266 + int32(4)
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v281 = v266
	goto L87
L86:
	;
	goto L85
L87:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v285 != 0 {
		v281 = v281 + int32(1)
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v287 = v281
	goto L75
L89:
	;
	goto L88
L90:
	;
	v446 = v161
	goto L47
L91:
	;
	goto L90
L92:
	;
	v301 = v297 + v161
	if base.Ui32(int32(0)-v297<<(uint(int32(1))%32)) < base.Ui32(v165-v301) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v311 = (v165 ^ v161) & int32(3)
	if base.Ui32(v165) <= base.Ui32(v161) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v308 = F___memcpy(m, v161, v165, v297)
	mBase = m.M
	goto L90
L95:
	;
	if v417 == int32(0) {
		goto L91
	} else {
		goto L127
	}
L96:
	;
	if base.Ui32(v395) <= base.Ui32(int32(3)) {
		v416 = v394
		v417 = v395
		v418 = v396
		goto L95
	} else {
		goto L123
	}
L97:
	;
	if v311 != 0 {
		v377 = v297
		goto L107
	} else {
		goto L108
	}
L98:
	;
	if v311 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if v161&int32(3) != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v416 = v165
	v417 = v297
	v418 = v161
	goto L95
L101:
	;
	v318 = v165
	v319 = v297
	v320 = v161
	goto L103
L102:
	;
	v394 = v165
	v395 = v297
	v396 = v161
	goto L96
L103:
	;
	if v319 == int32(0) {
		goto L91
	} else {
		goto L105
	}
L105:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	*(*uint8)(unsafe.Add(mBase, uint32(v320))) = uint8(v324)
	v326 = int32(1)
	v327 = v318 + v326
	v329 = v319 + int32(-1)
	v331 = v320 + v326
	if v331&int32(3) == int32(0) {
		v394 = v327
		v395 = v329
		v396 = v331
		goto L96
	} else {
		goto L106
	}
L106:
	;
	v318 = v327
	v319 = v329
	v320 = v331
	goto L103
L107:
	;
	if v377 == int32(0) {
		goto L91
	} else {
		goto L119
	}
L108:
	;
	if v301&int32(3) == int32(0) {
		v357 = v297
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if base.Ui32(v357) <= base.Ui32(int32(3)) {
		v377 = v357
		goto L107
	} else {
		goto L115
	}
L110:
	;
	v342 = v297
	goto L111
L111:
	;
	if v342 == int32(0) {
		goto L91
	} else {
		goto L113
	}
L112:
	;
	v357 = v348
	goto L109
L113:
	;
	v348 = v342 + int32(-1)
	v349 = v161 + v348
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v348))))
	*(*uint8)(unsafe.Add(mBase, uint32(v349))) = uint8(v351)
	if v349&int32(3) != 0 {
		v342 = v348
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v364 = v357
	goto L116
L116:
	;
	v368 = v364 + int32(-4)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v165+v368)))
	*(*int32)(unsafe.Add(mBase, uint32(v161+v368))) = v371
	if base.Ui32(int32(3)) < base.Ui32(v368) {
		v364 = v368
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v377 = v368
	goto L107
L118:
	;
	goto L117
L119:
	;
	v384 = v377
	goto L120
L120:
	;
	v388 = v384 + int32(-1)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v388))))
	*(*uint8)(unsafe.Add(mBase, uint32(v161+v388))) = uint8(v391)
	if v388 != 0 {
		v384 = v388
		goto L120
	} else {
		goto L122
	}
L122:
	;
	goto L91
L123:
	;
	v401 = v394
	v402 = v395
	v403 = v396
	goto L124
L124:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	*(*int32)(unsafe.Add(mBase, uint32(v403))) = v405
	v407 = int32(4)
	v408 = v401 + v407
	v410 = v403 + v407
	v412 = v402 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v412) {
		v401 = v408
		v402 = v412
		v403 = v410
		goto L124
	} else {
		goto L126
	}
L125:
	;
	v416 = v408
	v417 = v412
	v418 = v410
	goto L95
L126:
	;
	goto L125
L127:
	;
	v423 = v416
	v424 = v417
	v425 = v418
	goto L128
L128:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423))))
	*(*uint8)(unsafe.Add(mBase, uint32(v425))) = uint8(v427)
	v429 = int32(1)
	v434 = v424 + int32(-1)
	if v434 != 0 {
		v423 = v423 + v429
		v424 = v434
		v425 = v425 + v429
		goto L128
	} else {
		goto L130
	}
L129:
	;
	goto L91
L130:
	;
	goto L129
L131:
	;
	goto L24
L132:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482+int32(-1)))))
	switch v485 & int32(7) {
	case 0:
		goto L145
	case 1:
		goto L144
	case 2:
		goto L143
	case 3:
		goto L142
	case 4:
		goto L141
	default:
		v509 = v482
		goto L139
	}
L133:
	;
	goto L132
L134:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v455+int32(-17)))) = base.I64_extend_i32_u(v460)
	goto L133
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v455+int32(-9)))) = v460
	goto L132
L136:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v455+int32(-5)))) = uint16(v460)
	goto L132
L137:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v455+int32(-3)))) = uint8(v460)
	goto L132
L138:
	;
	v467 = v460 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v462))) = uint8(v467)
	goto L132
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v14
	if l2 != 0 {
		goto L148
	} else {
		goto L149
	}
L140:
	;
	if v502 == int32(0) {
		v509 = v482
		goto L139
	} else {
		goto L146
	}
L141:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v482+int32(-17))))
	v502 = v501
	goto L140
L142:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v482+int32(-9))))
	v502 = v498
	goto L140
L143:
	;
	v495 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482+int32(-5)))))
	v502 = v495
	goto L140
L144:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482+int32(-3)))))
	v502 = v492
	goto L140
L145:
	;
	v502 = int32(base.Ui32(v485) >> (uint(int32(3)) % 32))
	goto L140
L146:
	;
	v506 = F_sdscat(m, v482, int32(_a10))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v506
	v509 = v506
	goto L139
L148:
	;
	v514 = int32(_a11)
	goto L150
L149:
	;
	v514 = int32(_a12)
	goto L150
L150:
	;
	v515 = F_sdscatfmt(m, v509, v514, v12)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v515
	F_sdsfree(m, v14)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	m.G0 = v12 + int32(16)
	return
}
