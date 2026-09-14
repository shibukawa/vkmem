package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_getLongDoubleFromObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v65 int32
	_ = v65
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int64
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v271 int64
	_ = v271
	var v286 int64
	_ = v286
	var v287 int64
	_ = v287
	var v288 int64
	_ = v288
	var v297 int64
	_ = v297
	var v298 int64
	_ = v298
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v309 int32
	_ = v309
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	if l0 != 0 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v16&int32(15) != 0 {
			F__serverAssertWithInfo(m, int32(0), l0, int32(_a479), int32(_a838), int32(1065))
			mBase = m.M
			v325 = m.ExcPending
			if v325 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			switch int32(base.Ui32(v16)>>(uint(int32(4))%32)) & int32(15) {
			case 0, 8:
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v23&int32(4) == int32(0) {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v84 = v23
					v86 = v83
				} else {
					if v23&int32(1) != 0 {
						v32 = int32(16)
					} else {
						v32 = int32(8)
					}
					v33 = l0 + v32
					if v23&int32(2) == int32(0) {
						v65 = v33
					} else {
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
						v39 = v33 + v38
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
						switch v43 & int32(7) {
						case 0:
							v60 = int32(base.Ui32(v43) >> (uint(int32(3)) % 32))
						case 1:
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+int32(-2)))))
							v60 = v50
						case 2:
							v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39+int32(-4)))))
							v60 = v53
						case 3:
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(-8))))
							v60 = v56
						case 4:
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(-16))))
							v60 = v59
						default:
							v60 = int32(0)
						}
						v65 = v39 + int32(1) + v60 + int32(1)
					}
					v79 = *(*int32)(unsafe.Add(mBase, _consts[249]))
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v84 = v82
					v86 = v65 + v79
				}
				if v84&int32(4) == int32(0) {
					v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v149 = v146
				} else {
					if v84&int32(1) != 0 {
						v96 = int32(16)
					} else {
						v96 = int32(8)
					}
					v97 = l0 + v96
					if v84&int32(2) == int32(0) {
						v128 = v97
					} else {
						v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
						v103 = v97 + v102
						v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
						switch v107 & int32(7) {
						case 0:
							v124 = int32(base.Ui32(v107) >> (uint(int32(3)) % 32))
						case 1:
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103+int32(-2)))))
							v124 = v114
						case 2:
							v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103+int32(-4)))))
							v124 = v117
						case 3:
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v103+int32(-8))))
							v124 = v120
						case 4:
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v103+int32(-16))))
							v124 = v123
						default:
							v124 = int32(0)
						}
						v128 = v103 + int32(1) + v124 + int32(1)
					}
					v143 = *(*int32)(unsafe.Add(mBase, _consts[249]))
					v149 = v128 + v143
				}
				v151 = int32(-1)
				v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+v151))))
				switch v155 & int32(7) {
				case 0:
					v172 = int32(base.Ui32(v155) >> (uint(int32(3)) % 32))
				case 1:
					v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+int32(-3)))))
					v172 = v162
				case 2:
					v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149+int32(-5)))))
					v172 = v165
				case 3:
					v168 = *(*int32)(unsafe.Add(mBase, uint32(v149+int32(-9))))
					v172 = v168
				case 4:
					v171 = *(*int32)(unsafe.Add(mBase, uint32(v149+int32(-17))))
					v172 = v171
				default:
					v172 = int32(0)
				}
				v175 = F_string2ld(m, v86, v172, v12+int32(16))
				mBase = m.M
				v178 = m.ExcPending
				if v178 != 0 {
					return int32(0)
				} else {
					if v175 == int32(0) {
						v309 = v151
					} else {
						v183 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(24))))
						v184 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
						v300 = v184
						v301 = v183
						*(*int64)(unsafe.Add(mBase, uint32(l1))) = v300
						*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v301
						v309 = int32(0)
					}
					m.G0 = v12 + int32(32)
					return v309
				}
			case 1:
				v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v185&int32(4) == int32(0) {
					v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v245 = v244
				} else {
					if v185&int32(1) != 0 {
						v194 = int32(16)
					} else {
						v194 = int32(8)
					}
					v195 = l0 + v194
					if v185&int32(2) == int32(0) {
						v226 = v195
					} else {
						v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
						v201 = v195 + v200
						v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
						switch v205 & int32(7) {
						case 0:
							v222 = int32(base.Ui32(v205) >> (uint(int32(3)) % 32))
						case 1:
							v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+int32(-2)))))
							v222 = v212
						case 2:
							v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201+int32(-4)))))
							v222 = v215
						case 3:
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(-8))))
							v222 = v218
						case 4:
							v221 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(-16))))
							v222 = v221
						default:
							v222 = int32(0)
						}
						v226 = v201 + int32(1) + v222 + int32(1)
					}
					v241 = *(*int32)(unsafe.Add(mBase, _consts[249]))
					v245 = v226 + v241
				}
				v253 = m.G0
				v255 = v253 - int32(16)
				m.G0 = v255
				if v245 != 0 {
					v260 = v245 >> (uint(int32(31)) % 32)
					v262 = v245 ^ v260 - v260
					v265 = base.I32_clz(v262)
					F___ashlti3(m, v255, base.I64_extend_i32_u(v262), int64(0), v265+int32(81))
					mBase = m.M
					v271 = *(*int64)(unsafe.Add(mBase, uint32(v255+int32(8))))
					v286 = *(*int64)(unsafe.Add(mBase, uint32(v255)))
					v287 = v286
					v288 = v271 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v265)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v245&int32(-2147483648))<<(uint(int64(32))%64)
				} else {
					v257 = int64(0)
					v287 = v257
					v288 = v257
				}
				*(*int64)(unsafe.Add(mBase, uint32(v12))) = v287
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v288
				m.G0 = v255 + int32(16)
				v297 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(8))))
				v298 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
				v300 = v298
				v301 = v297
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = v300
				*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v301
				v309 = int32(0)
				m.G0 = v12 + int32(32)
				return v309
			default:
				F__serverPanic_1(m, int32(_a838), int32(1071), int32(_a848), int32(0))
				mBase = m.M
				v332 = m.ExcPending
				if v332 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v14 = int64(0)
		v300 = v14
		v301 = v14
		*(*int64)(unsafe.Add(mBase, uint32(l1))) = v300
		*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v301
		v309 = int32(0)
		m.G0 = v12 + int32(32)
		return v309
	}
}
func F_getLongLongFromObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v221 int64
	_ = v221
	var v227 int32
	_ = v227
	var v229 int64
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v243 int64
	_ = v243
	var v248 int64
	_ = v248
	var v252 int32
	_ = v252
	var v254 int64
	_ = v254
	var v256 int32
	_ = v256
	var v264 int64
	_ = v264
	var v288 int64
	_ = v288
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v370 int32
	_ = v370
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int64
	_ = v386
	var v397 int32
	_ = v397
	var v400 int64
	_ = v400
	var v402 int32
	_ = v402
	var v417 int32
	_ = v417
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F__serverAssertWithInfo(m, int32(0), l0, int32(_a479), int32(_a838), int32(1098))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L95
	} else {
		goto L98
	}
L2:
	;
	m.G0 = v11 + int32(16)
	return v402
L3:
	;
	v397 = int32(0)
	if l1 == v397 {
		v402 = v397
		goto L2
	} else {
		goto L97
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v386
	goto L3
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14&int32(15) != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v386 = int64(0)
	goto L4
L7:
	;
	switch int32(base.Ui32(v14)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		goto L11
	case 1:
		goto L10
	default:
		goto L9
	}
L8:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v386 = base.I64_extend_i32_s(v383)
	goto L4
L9:
	;
	F__serverPanic_1(m, int32(_a838), int32(1104), int32(_a848), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L95
	} else {
		goto L96
	}
L10:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v314&int32(4) == int32(0) {
		goto L8
	} else {
		goto L80
	}
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21&int32(4) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v82&int32(4) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v82 = v21
	v84 = v81
	goto L12
L14:
	;
	if v21&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v30 = int32(16)
	goto L17
L16:
	;
	v30 = int32(8)
	goto L17
L17:
	;
	v31 = l0 + v30
	if v21&int32(2) == int32(0) {
		v63 = v31
		goto L18
	} else {
		goto L19
	}
L18:
	;
	goto L28
L19:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v37 = v31 + v36
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	switch v41 & int32(7) {
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
		v58 = int32(0)
		goto L20
	}
L20:
	;
	v63 = v37 + int32(1) + v58 + int32(1)
	goto L18
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-16))))
	v58 = v57
	goto L20
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-8))))
	v58 = v54
	goto L20
L23:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(-4)))))
	v58 = v51
	goto L20
L24:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-2)))))
	v58 = v48
	goto L20
L25:
	;
	v58 = int32(base.Ui32(v41) >> (uint(int32(3)) % 32))
	goto L20
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v82 = v80
	v84 = v63 + v77
	goto L12
L27:
	;
	goto L26
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L27
L29:
	;
	v149 = int32(-1)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v149))))
	switch v153 & int32(7) {
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
		v170 = int32(0)
		goto L46
	}
L30:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v147 = v144
	goto L29
L31:
	;
	if v82&int32(1) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v94 = int32(16)
	goto L34
L33:
	;
	v94 = int32(8)
	goto L34
L34:
	;
	v95 = l0 + v94
	if v82&int32(2) == int32(0) {
		v126 = v95
		goto L35
	} else {
		goto L36
	}
L35:
	;
	goto L45
L36:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v101 = v95 + v100
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	switch v105 & int32(7) {
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
		v122 = int32(0)
		goto L37
	}
L37:
	;
	v126 = v101 + int32(1) + v122 + int32(1)
	goto L35
L38:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v101+int32(-16))))
	v122 = v121
	goto L37
L39:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v101+int32(-8))))
	v122 = v118
	goto L37
L40:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101+int32(-4)))))
	v122 = v115
	goto L37
L41:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+int32(-2)))))
	v122 = v112
	goto L37
L42:
	;
	v122 = int32(base.Ui32(v105) >> (uint(int32(3)) % 32))
	goto L37
L43:
	;
	v147 = v126 + v141
	goto L29
L44:
	;
	goto L43
L45:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L44
L46:
	;
	v172 = v11 + int32(8)
	v173 = int32(0)
	if base.Ui32(v170+int32(-21)) < base.Ui32(int32(-20)) {
		v307 = v173
		goto L53
	} else {
		goto L54
	}
L47:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v147+int32(-17))))
	v170 = v169
	goto L46
L48:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v147+int32(-9))))
	v170 = v166
	goto L46
L49:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147+int32(-5)))))
	v170 = v163
	goto L46
L50:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+int32(-3)))))
	v170 = v160
	goto L46
L51:
	;
	v170 = int32(base.Ui32(v153) >> (uint(int32(3)) % 32))
	goto L46
L52:
	;
	if v307 != 0 {
		goto L3
	} else {
		goto L79
	}
L53:
	;
	goto L52
L54:
	;
	v185 = int32(1)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v170 != v185 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v307 = int32(1)
	goto L53
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v172))) = v288
	goto L55
L57:
	;
	if v186&int32(255) == int32(45) {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	v190 = v186 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v190&int32(255)) {
		v307 = v173
		goto L53
	} else {
		goto L59
	}
L59:
	;
	if v172 == int32(0) {
		goto L55
	} else {
		goto L60
	}
L60:
	;
	v288 = base.I64_extend_i32_u(v190) & int64(255)
	goto L56
L61:
	;
	if base.Ui32(int32(8)) < base.Ui32((v209+int32(-49))&int32(255)) {
		v307 = v173
		goto L53
	} else {
		goto L64
	}
L62:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v208 = int32(2)
	v209 = v206
	v210 = v84 + int32(1)
	goto L61
L63:
	;
	v208 = v185
	v209 = v186
	v210 = v84
	goto L61
L64:
	;
	v221 = base.I64_extend_i32_u(v209+int32(-48)) & int64(255)
	if base.Ui32(v170) <= base.Ui32(v208) {
		v264 = v221
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if v186&int32(255) != int32(45) {
		goto L73
	} else {
		goto L74
	}
L66:
	;
	v227 = v208
	v229 = v221
	v231 = v210
	goto L67
L67:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
	if base.Ui32((v233+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v307 = v173
		goto L53
	} else {
		goto L69
	}
L68:
	;
	v264 = v254
	goto L65
L69:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v229) {
		v307 = v173
		goto L53
	} else {
		goto L70
	}
L70:
	;
	v243 = v229 * int64(10)
	v248 = base.I64_extend_i32_u(v233+int32(-48)) & int64(255)
	if base.Ui64(v248^int64(-1)) < base.Ui64(v243) {
		v307 = v173
		goto L53
	} else {
		goto L71
	}
L71:
	;
	v252 = int32(1)
	v254 = v243 + v248
	v256 = v227 + v252
	if v256 != v170 {
		v227 = v256
		v229 = v254
		v231 = v231 + v252
		goto L67
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	if v264 < int64(0) {
		v307 = v173
		goto L53
	} else {
		goto L77
	}
L74:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v264) {
		v307 = v173
		goto L53
	} else {
		goto L75
	}
L75:
	;
	if v172 == int32(0) {
		goto L55
	} else {
		goto L76
	}
L76:
	;
	v288 = int64(0) - v264
	goto L56
L77:
	;
	if v172 == int32(0) {
		goto L55
	} else {
		goto L78
	}
L78:
	;
	v288 = v264
	goto L56
L79:
	;
	v402 = v149
	goto L2
L80:
	;
	if v314&int32(1) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v323 = int32(16)
	goto L83
L82:
	;
	v323 = int32(8)
	goto L83
L83:
	;
	v324 = l0 + v323
	if v314&int32(2) == int32(0) {
		v355 = v324
		goto L84
	} else {
		goto L85
	}
L84:
	;
	goto L94
L85:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	v330 = v324 + v329
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	switch v334 & int32(7) {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	case 3:
		goto L88
	case 4:
		goto L87
	default:
		v351 = int32(0)
		goto L86
	}
L86:
	;
	v355 = v330 + int32(1) + v351 + int32(1)
	goto L84
L87:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v330+int32(-16))))
	v351 = v350
	goto L86
L88:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v330+int32(-8))))
	v351 = v347
	goto L86
L89:
	;
	v344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v330+int32(-4)))))
	v351 = v344
	goto L86
L90:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330+int32(-2)))))
	v351 = v341
	goto L86
L91:
	;
	v351 = int32(base.Ui32(v334) >> (uint(int32(3)) % 32))
	goto L86
L92:
	;
	v386 = base.I64_extend_i32_s(v355 + v370)
	goto L4
L93:
	;
	goto L92
L94:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L93
L95:
	;
	return int32(0)
L96:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v400 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v400
	v402 = v397
	goto L2
L98:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
