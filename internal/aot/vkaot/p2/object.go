package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_freeSetObject(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v5)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v12&int32(4) == int32(0) {
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_hashtableRelease(m, v73)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return
			} else {
				return
			}
		} else {
			if v12&int32(1) != 0 {
				v21 = int32(16)
			} else {
				v21 = int32(8)
			}
			v22 = l0 + v21
			if v12&int32(2) == int32(0) {
				v53 = v22
			} else {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v28 = v22 + v27
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				switch v32 & int32(7) {
				case 0:
					v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
				case 1:
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-2)))))
					v49 = v39
				case 2:
					v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-4)))))
					v49 = v42
				case 3:
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-8))))
					v49 = v45
				case 4:
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-16))))
					v49 = v48
				default:
					v49 = int32(0)
				}
				v53 = v28 + int32(1) + v49 + int32(1)
			}
			v68 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			F_hashtableRelease(m, v53+v68)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				return
			}
		}
	default:
		F__serverPanic_1(m, int32(_a1051), int32(576), int32(_a1053), int32(0))
		mBase = m.M
		v145 = m.ExcPending
		if v145 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 4, 9:
		v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v76&int32(4) == int32(0) {
			v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_valkey_free(m, v137)
			mBase = m.M
			v139 = m.ExcPending
			if v139 != 0 {
				return
			} else {
				return
			}
		} else {
			if v76&int32(1) != 0 {
				v85 = int32(16)
			} else {
				v85 = int32(8)
			}
			v86 = l0 + v85
			if v76&int32(2) == int32(0) {
				v117 = v86
			} else {
				v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
				v92 = v86 + v91
				v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
				switch v96 & int32(7) {
				case 0:
					v113 = int32(base.Ui32(v96) >> (uint(int32(3)) % 32))
				case 1:
					v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+int32(-2)))))
					v113 = v103
				case 2:
					v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92+int32(-4)))))
					v113 = v106
				case 3:
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v92+int32(-8))))
					v113 = v109
				case 4:
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v92+int32(-16))))
					v113 = v112
				default:
					v113 = int32(0)
				}
				v117 = v92 + int32(1) + v113 + int32(1)
			}
			v132 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			F_valkey_free(m, v117+v132)
			mBase = m.M
			v136 = m.ExcPending
			if v136 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_getObjectReadOnlyString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	if l0 == int32(0) {
		v85 = int32(0)
		if l1 == v85 {
			v95 = v85
		} else {
			v90 = v85
			v91 = v85
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v91
			v95 = v90
		}
		return v95
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v8&int32(15) != 0 {
			F__serverAssert(m, int32(_a181), int32(_a182), int32(671))
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if v8&int32(240) != int32(16) {
				v61 = F_objectGetVal(m, l0)
				mBase = m.M
				if l1 == int32(0) {
					v95 = v61
				} else {
					v65 = F_objectGetVal(m, l0)
					mBase = m.M
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+int32(-1)))))
					switch v68 & int32(7) {
					case 0:
						v90 = v61
						v91 = int32(base.Ui32(v68) >> (uint(int32(3)) % 32))
					case 1:
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+int32(-3)))))
						v90 = v61
						v91 = v75
					case 2:
						v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65+int32(-5)))))
						v90 = v61
						v91 = v78
					case 3:
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v65+int32(-9))))
						v90 = v61
						v91 = v81
					case 4:
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v65+int32(-17))))
						v90 = v61
						v91 = v84
					default:
						v90 = v61
						v91 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v91
					v95 = v90
				}
			} else {
				if l1 == int32(0) {
					v95 = l2
				} else {
					v18 = F_objectGetVal(m, l0)
					mBase = m.M
					v19 = base.I64_extend_i32_s(v18)
					if v19 <= int64(-1) {
						v28 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v28)
						v32 = int32(1)
						v37 = l2 + v32
						v38 = int32(20)
						v39 = int64(0) - v19
						v40 = v32
					} else {
						v37 = l2
						v38 = int32(21)
						v39 = v19
						v40 = int32(0)
					}
					v41 = F_ull2string(m, v37, v38, v39)
					mBase = m.M
					if v41 == int32(0) {
						v60 = int32(0)
					} else {
						v60 = v41 + v40
					}
					v90 = l2
					v91 = v60
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v91
					v95 = v90
				}
			}
			return v95
		}
	}
}
func F_getObjectTypeByName(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
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
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
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
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int64
	_ = v318
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[339]))
	if v4 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v47 = int32(0)
	v48 = *(*int32)(unsafe.Add(mBase, _consts[340]))
	if v48 == v47 {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if v41-v43 != 0 {
		goto L1
	} else {
		goto L15
	}
L4:
	;
	v41 = F_tolower(m, v37)
	mBase = m.M
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v43 = F_tolower(m, v42)
	mBase = m.M
	goto L3
L5:
	;
	v11 = l0
	v12 = v4
	v13 = v9
	goto L8
L6:
	;
	v37 = int32(0)
	v38 = v4
	goto L4
L7:
	;
	v37 = v34 & int32(255)
	v38 = v33
	goto L4
L8:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v15 == int32(0) {
		v33 = v12
		v34 = v13
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v33 = v27
	v34 = int32(0)
	goto L7
L10:
	;
	v19 = v13 & int32(255)
	if v19 == v15 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v26 = int32(1)
	v27 = v12 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v28 != 0 {
		v11 = v11 + v26
		v12 = v27
		v13 = v28
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v21 = F_tolower(m, v19)
	mBase = m.M
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v23 = F_tolower(m, v22)
	mBase = m.M
	if v21 == v23 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v33 = v12
	v34 = v25
	goto L7
L14:
	;
	goto L9
L15:
	;
	return int64(0)
L16:
	;
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, _consts[341]))
	if v92 == v91 {
		goto L31
	} else {
		goto L32
	}
L17:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v53 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v85-v87 != 0 {
		goto L16
	} else {
		goto L30
	}
L19:
	;
	v85 = F_tolower(m, v81)
	mBase = m.M
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v87 = F_tolower(m, v86)
	mBase = m.M
	goto L18
L20:
	;
	v55 = l0
	v56 = v48
	v57 = v53
	goto L23
L21:
	;
	v81 = int32(0)
	v82 = v48
	goto L19
L22:
	;
	v81 = v78 & int32(255)
	v82 = v77
	goto L19
L23:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v59 == int32(0) {
		v77 = v56
		v78 = v57
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v77 = v71
	v78 = int32(0)
	goto L22
L25:
	;
	v63 = v57 & int32(255)
	if v63 == v59 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v70 = int32(1)
	v71 = v56 + v70
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v72 != 0 {
		v55 = v55 + v70
		v56 = v71
		v57 = v72
		goto L23
	} else {
		goto L29
	}
L27:
	;
	v65 = F_tolower(m, v63)
	mBase = m.M
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	v67 = F_tolower(m, v66)
	mBase = m.M
	if v65 == v67 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v77 = v56
	v78 = v69
	goto L22
L29:
	;
	goto L24
L30:
	;
	return int64(1)
L31:
	;
	v135 = int32(0)
	v136 = *(*int32)(unsafe.Add(mBase, _consts[342]))
	if v136 == v135 {
		goto L46
	} else {
		goto L47
	}
L32:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v97 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if v129-v131 != 0 {
		goto L31
	} else {
		goto L45
	}
L34:
	;
	v129 = F_tolower(m, v125)
	mBase = m.M
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v131 = F_tolower(m, v130)
	mBase = m.M
	goto L33
L35:
	;
	v99 = l0
	v100 = v92
	v101 = v97
	goto L38
L36:
	;
	v125 = int32(0)
	v126 = v92
	goto L34
L37:
	;
	v125 = v122 & int32(255)
	v126 = v121
	goto L34
L38:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v103 == int32(0) {
		v121 = v100
		v122 = v101
		goto L37
	} else {
		goto L40
	}
L39:
	;
	v121 = v115
	v122 = int32(0)
	goto L37
L40:
	;
	v107 = v101 & int32(255)
	if v107 == v103 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v114 = int32(1)
	v115 = v100 + v114
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if v116 != 0 {
		v99 = v99 + v114
		v100 = v115
		v101 = v116
		goto L38
	} else {
		goto L44
	}
L42:
	;
	v109 = F_tolower(m, v107)
	mBase = m.M
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	v111 = F_tolower(m, v110)
	mBase = m.M
	if v109 == v111 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v121 = v100
	v122 = v113
	goto L37
L44:
	;
	goto L39
L45:
	;
	return int64(2)
L46:
	;
	v179 = int32(0)
	v180 = *(*int32)(unsafe.Add(mBase, _consts[343]))
	if v180 == v179 {
		goto L61
	} else {
		goto L62
	}
L47:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v141 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	if v173-v175 != 0 {
		goto L46
	} else {
		goto L60
	}
L49:
	;
	v173 = F_tolower(m, v169)
	mBase = m.M
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	v175 = F_tolower(m, v174)
	mBase = m.M
	goto L48
L50:
	;
	v143 = l0
	v144 = v136
	v145 = v141
	goto L53
L51:
	;
	v169 = int32(0)
	v170 = v136
	goto L49
L52:
	;
	v169 = v166 & int32(255)
	v170 = v165
	goto L49
L53:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v147 == int32(0) {
		v165 = v144
		v166 = v145
		goto L52
	} else {
		goto L55
	}
L54:
	;
	v165 = v159
	v166 = int32(0)
	goto L52
L55:
	;
	v151 = v145 & int32(255)
	if v151 == v147 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v158 = int32(1)
	v159 = v144 + v158
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+1)))
	if v160 != 0 {
		v143 = v143 + v158
		v144 = v159
		v145 = v160
		goto L53
	} else {
		goto L59
	}
L57:
	;
	v153 = F_tolower(m, v151)
	mBase = m.M
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v155 = F_tolower(m, v154)
	mBase = m.M
	if v153 == v155 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v165 = v144
	v166 = v157
	goto L52
L59:
	;
	goto L54
L60:
	;
	return int64(3)
L61:
	;
	v223 = int32(0)
	v224 = *(*int32)(unsafe.Add(mBase, _consts[344]))
	if v224 == v223 {
		goto L76
	} else {
		goto L77
	}
L62:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v185 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	if v217-v219 != 0 {
		goto L61
	} else {
		goto L75
	}
L64:
	;
	v217 = F_tolower(m, v213)
	mBase = m.M
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	v219 = F_tolower(m, v218)
	mBase = m.M
	goto L63
L65:
	;
	v187 = l0
	v188 = v180
	v189 = v185
	goto L68
L66:
	;
	v213 = int32(0)
	v214 = v180
	goto L64
L67:
	;
	v213 = v210 & int32(255)
	v214 = v209
	goto L64
L68:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v191 == int32(0) {
		v209 = v188
		v210 = v189
		goto L67
	} else {
		goto L70
	}
L69:
	;
	v209 = v203
	v210 = int32(0)
	goto L67
L70:
	;
	v195 = v189 & int32(255)
	if v195 == v191 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v202 = int32(1)
	v203 = v188 + v202
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	if v204 != 0 {
		v187 = v187 + v202
		v188 = v203
		v189 = v204
		goto L68
	} else {
		goto L74
	}
L72:
	;
	v197 = F_tolower(m, v195)
	mBase = m.M
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	v199 = F_tolower(m, v198)
	mBase = m.M
	if v197 == v199 {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v209 = v188
	v210 = v201
	goto L67
L74:
	;
	goto L69
L75:
	;
	return int64(4)
L76:
	;
	v267 = int32(0)
	v268 = *(*int32)(unsafe.Add(mBase, _consts[345]))
	if v268 == v267 {
		goto L91
	} else {
		goto L92
	}
L77:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v229 != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	if v261-v263 != 0 {
		goto L76
	} else {
		goto L90
	}
L79:
	;
	v261 = F_tolower(m, v257)
	mBase = m.M
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	v263 = F_tolower(m, v262)
	mBase = m.M
	goto L78
L80:
	;
	v231 = l0
	v232 = v224
	v233 = v229
	goto L83
L81:
	;
	v257 = int32(0)
	v258 = v224
	goto L79
L82:
	;
	v257 = v254 & int32(255)
	v258 = v253
	goto L79
L83:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	if v235 == int32(0) {
		v253 = v232
		v254 = v233
		goto L82
	} else {
		goto L85
	}
L84:
	;
	v253 = v247
	v254 = int32(0)
	goto L82
L85:
	;
	v239 = v233 & int32(255)
	if v239 == v235 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v246 = int32(1)
	v247 = v232 + v246
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
	if v248 != 0 {
		v231 = v231 + v246
		v232 = v247
		v233 = v248
		goto L83
	} else {
		goto L89
	}
L87:
	;
	v241 = F_tolower(m, v239)
	mBase = m.M
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	v243 = F_tolower(m, v242)
	mBase = m.M
	if v241 == v243 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	v253 = v232
	v254 = v245
	goto L82
L89:
	;
	goto L84
L90:
	;
	return int64(5)
L91:
	;
	v311 = F_moduleTypeLookupModuleByNameIgnoreCase(m, l0)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L107
	} else {
		goto L108
	}
L92:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v273 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	if v305-v307 != 0 {
		goto L91
	} else {
		goto L105
	}
L94:
	;
	v305 = F_tolower(m, v301)
	mBase = m.M
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	v307 = F_tolower(m, v306)
	mBase = m.M
	goto L93
L95:
	;
	v275 = l0
	v276 = v268
	v277 = v273
	goto L98
L96:
	;
	v301 = int32(0)
	v302 = v268
	goto L94
L97:
	;
	v301 = v298 & int32(255)
	v302 = v297
	goto L94
L98:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	if v279 == int32(0) {
		v297 = v276
		v298 = v277
		goto L97
	} else {
		goto L100
	}
L99:
	;
	v297 = v291
	v298 = int32(0)
	goto L97
L100:
	;
	v283 = v277 & int32(255)
	if v283 == v279 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v290 = int32(1)
	v291 = v276 + v290
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+1)))
	if v292 != 0 {
		v275 = v275 + v290
		v276 = v291
		v277 = v292
		goto L98
	} else {
		goto L104
	}
L102:
	;
	v285 = F_tolower(m, v283)
	mBase = m.M
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	v287 = F_tolower(m, v286)
	mBase = m.M
	if v285 == v287 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	v297 = v276
	v298 = v289
	goto L97
L104:
	;
	goto L99
L105:
	;
	return int64(6)
L106:
	;
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v311)))
	return int64(0) - int64(base.Ui64(v318)>>(uint(int64(10))%64))
L107:
	;
	return int64(0)
L108:
	;
	if v311 != 0 {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	return int64(9223372036854775807)
}
func F_objectCommandLookupOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v6 = F_lookupKeyReadWithFlags(m, v4, l1, int32(3))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			return v6
		} else {
			F_addReplyOrErrorObject(m, l0, l2)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v6
			}
		}
	}
}
func F_objectGetExpire(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5&int32(1) == int32(0) {
		v16 = int64(-1)
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0+(v5&int32(4)^int32(12)))))
		v16 = v15
	}
	return v16
}
func F_objectGetKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5&int32(2) == v2 {
		v25 = v2
	} else {
		v19 = l0 + (v5&int32(4) ^ int32(12)) + v5<<(uint(int32(3))%32)&int32(8)
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
		v25 = v19 + v20 + int32(1)
	}
	return v25
}
func F_objectGetLRUIdleSecs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	return (v6 - int32(base.Ui32(v2)>>(uint(int32(8))%32))) & int32(16777215)
}
func F_objectSetExpire(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5&int32(1) == int32(0) {
		if l1 == int64(-1) {
			v39 = l0
			return v39
		} else {
			v19 = int32(0)
			if v5&int32(2) == v19 {
				v34 = v19
			} else {
				v28 = l0 + (v5&int32(4) ^ int32(12))
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v34 = v28 + v29 + int32(1)
			}
			v35 = F_objectSetKeyAndExpire(m, l0, v34, l1)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v39 = v35
				return v39
			}
		}
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(l0+(v5&int32(4)^int32(12))))) = l1
		return l0
	}
}
func F_objectSetKeyAndExpire(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = v10 & int32(255)
	if v12 != int32(128) {
		v255 = v10 & int32(15)
		if v9&int32(-8) != int32(8) {
			switch v12 {
			case 0:
				v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v264 = F_sdsdup(m, v263)
				mBase = m.M
				v265 = m.ExcPending
				if v265 != 0 {
					return int32(0)
				} else {
					v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v279 = v264
					v280 = v266 & int32(15)
					v281 = F_createUnembeddedObjectWithKeyAndExpire(m, v280, v279, l1, l2)
					mBase = m.M
					v282 = m.ExcPending
					if v282 != 0 {
						return int32(0)
					} else {
						v283 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
						v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v289 = v283&int32(-241) | v286&int32(240)
						*(*int32)(unsafe.Add(mBase, uint32(v281))) = v289
						v291 = v289
						v292 = v281
						v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
						F_decrRefCount(m, l0)
						mBase = m.M
						v305 = m.ExcPending
						if v305 != 0 {
							return int32(0)
						} else {
							return v292
						}
					}
				}
			default:
				if v255 == int32(0) {
					F__serverAssert(m, int32(_a1050), int32(_a1051), int32(376))
					mBase = m.M
					v311 = m.ExcPending
					if v311 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					F__serverPanic_1(m, int32(_a1051), int32(379), int32(_a1052), int32(0))
					mBase = m.M
					v276 = m.ExcPending
					if v276 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			case 16:
				v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v279 = v278
				v280 = v255
				v281 = F_createUnembeddedObjectWithKeyAndExpire(m, v280, v279, l1, l2)
				mBase = m.M
				v282 = m.ExcPending
				if v282 != 0 {
					return int32(0)
				} else {
					v283 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
					v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v289 = v283&int32(-241) | v286&int32(240)
					*(*int32)(unsafe.Add(mBase, uint32(v281))) = v289
					v291 = v289
					v292 = v281
					v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
					F_decrRefCount(m, l0)
					mBase = m.M
					v305 = m.ExcPending
					if v305 != 0 {
						return int32(0)
					} else {
						return v292
					}
				}
			}
		} else {
			v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
			v279 = v260
			v280 = v255
			v281 = F_createUnembeddedObjectWithKeyAndExpire(m, v280, v279, l1, l2)
			mBase = m.M
			v282 = m.ExcPending
			if v282 != 0 {
				return int32(0)
			} else {
				v283 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
				v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v289 = v283&int32(-241) | v286&int32(240)
				*(*int32)(unsafe.Add(mBase, uint32(v281))) = v289
				v291 = v289
				v292 = v281
				v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
				F_decrRefCount(m, l0)
				mBase = m.M
				v305 = m.ExcPending
				if v305 != 0 {
					return int32(0)
				} else {
					return v292
				}
			}
		}
	} else {
		if v9&int32(4) == int32(0) {
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v75 = v9
			v77 = v74
		} else {
			if v9&int32(1) != 0 {
				v23 = int32(16)
			} else {
				v23 = int32(8)
			}
			v24 = l0 + v23
			if v9&int32(2) == int32(0) {
				v56 = v24
			} else {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v30 = v24 + v29
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
				switch v34 & int32(7) {
				case 0:
					v51 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
				case 1:
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+int32(-2)))))
					v51 = v41
				case 2:
					v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(-4)))))
					v51 = v44
				case 3:
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(-8))))
					v51 = v47
				case 4:
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(-16))))
					v51 = v50
				default:
					v51 = int32(0)
				}
				v56 = v30 + int32(1) + v51 + int32(1)
			}
			v70 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v75 = v73
			v77 = v56 + v70
		}
		if v75&int32(4) == int32(0) {
			v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v139 = v137
		} else {
			if v75&int32(1) != 0 {
				v87 = int32(16)
			} else {
				v87 = int32(8)
			}
			v88 = l0 + v87
			if v75&int32(2) == int32(0) {
				v120 = v88
			} else {
				v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
				v94 = v88 + v93
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
				switch v98 & int32(7) {
				case 0:
					v115 = int32(base.Ui32(v98) >> (uint(int32(3)) % 32))
				case 1:
					v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+int32(-2)))))
					v115 = v105
				case 2:
					v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94+int32(-4)))))
					v115 = v108
				case 3:
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(-8))))
					v115 = v111
				case 4:
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(-16))))
					v115 = v114
				default:
					v115 = int32(0)
				}
				v120 = v94 + int32(1) + v115 + int32(1)
			}
			v134 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			v139 = v120 + v134
		}
		v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139+int32(-1)))))
		switch v145 & int32(7) {
		case 0:
			v165 = int32(base.Ui32(v145) >> (uint(int32(3)) % 32))
			if l1 != 0 {
				v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
				switch v170 & int32(7) {
				case 0:
					v187 = int32(base.Ui32(v170) >> (uint(int32(3)) % 32))
				case 1:
					v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
					v187 = v177
				case 2:
					v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
					v187 = v180
				case 3:
					v183 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
					v187 = v183
				case 4:
					v186 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
					v187 = v186
				default:
					v187 = int32(0)
				}
				if base.Ui32(int32(32)) <= base.Ui32(v187) {
					if base.Ui32(int32(253)) <= base.Ui32(v187) {
						if base.Ui32(v187) < base.Ui32(int32(65531)) {
							v198 = int32(2)
						} else {
							v198 = int32(3)
						}
						v199 = v198
					} else {
						v199 = int32(1)
					}
				} else {
					v199 = int32(0)
				}
				v203 = v199 & int32(7)
				if base.Ui32(int32(4)) < base.Ui32(v203) {
					v211 = int32(0)
				} else {
					v210 = *(*int32)(unsafe.Add(mBase, uint32(v203<<(uint(int32(2))%32))+uint32(_consts[602])))
					v211 = v210
				}
				v215 = v187 + v211 + int32(10)
			} else {
				v215 = int32(8)
			}
			if l2 == int64(-1) {
				v221 = int32(1)
			} else {
				v221 = int32(9)
			}
			v235 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			if base.Ui32(int32(128)) < base.Ui32(v221+v165+v215+v235) {
				v245 = v165
				v249 = F_sdsnewlen(m, v77, v245)
				mBase = m.M
				v250 = m.ExcPending
				if v250 != 0 {
					return int32(0)
				} else {
					v251 = F_createUnembeddedObjectWithKeyAndExpire(m, int32(0), v249, l1, l2)
					mBase = m.M
					v252 = m.ExcPending
					if v252 != 0 {
						return int32(0)
					} else {
						v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
						v291 = v253
						v292 = v251
						v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
						F_decrRefCount(m, l0)
						mBase = m.M
						v305 = m.ExcPending
						if v305 != 0 {
							return int32(0)
						} else {
							return v292
						}
					}
				}
			} else {
				v240 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v77, v165, l1, l2)
				mBase = m.M
				v243 = m.ExcPending
				if v243 != 0 {
					return int32(0)
				} else {
					v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
					v291 = v244
					v292 = v240
					v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
					F_decrRefCount(m, l0)
					mBase = m.M
					v305 = m.ExcPending
					if v305 != 0 {
						return int32(0)
					} else {
						return v292
					}
				}
			}
		case 1:
			v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139+int32(-3)))))
			v165 = v152
			if l1 != 0 {
				v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
				switch v170 & int32(7) {
				case 0:
					v187 = int32(base.Ui32(v170) >> (uint(int32(3)) % 32))
				case 1:
					v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
					v187 = v177
				case 2:
					v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
					v187 = v180
				case 3:
					v183 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
					v187 = v183
				case 4:
					v186 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
					v187 = v186
				default:
					v187 = int32(0)
				}
				if base.Ui32(int32(32)) <= base.Ui32(v187) {
					if base.Ui32(int32(253)) <= base.Ui32(v187) {
						if base.Ui32(v187) < base.Ui32(int32(65531)) {
							v198 = int32(2)
						} else {
							v198 = int32(3)
						}
						v199 = v198
					} else {
						v199 = int32(1)
					}
				} else {
					v199 = int32(0)
				}
				v203 = v199 & int32(7)
				if base.Ui32(int32(4)) < base.Ui32(v203) {
					v211 = int32(0)
				} else {
					v210 = *(*int32)(unsafe.Add(mBase, uint32(v203<<(uint(int32(2))%32))+uint32(_consts[602])))
					v211 = v210
				}
				v215 = v187 + v211 + int32(10)
			} else {
				v215 = int32(8)
			}
			if l2 == int64(-1) {
				v221 = int32(1)
			} else {
				v221 = int32(9)
			}
			v235 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			if base.Ui32(int32(128)) < base.Ui32(v221+v165+v215+v235) {
				v245 = v165
				v249 = F_sdsnewlen(m, v77, v245)
				mBase = m.M
				v250 = m.ExcPending
				if v250 != 0 {
					return int32(0)
				} else {
					v251 = F_createUnembeddedObjectWithKeyAndExpire(m, int32(0), v249, l1, l2)
					mBase = m.M
					v252 = m.ExcPending
					if v252 != 0 {
						return int32(0)
					} else {
						v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
						v291 = v253
						v292 = v251
						v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
						F_decrRefCount(m, l0)
						mBase = m.M
						v305 = m.ExcPending
						if v305 != 0 {
							return int32(0)
						} else {
							return v292
						}
					}
				}
			} else {
				v240 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v77, v165, l1, l2)
				mBase = m.M
				v243 = m.ExcPending
				if v243 != 0 {
					return int32(0)
				} else {
					v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
					v291 = v244
					v292 = v240
					v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
					F_decrRefCount(m, l0)
					mBase = m.M
					v305 = m.ExcPending
					if v305 != 0 {
						return int32(0)
					} else {
						return v292
					}
				}
			}
		case 2:
			v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139+int32(-5)))))
			v162 = v155
			if base.Ui32(int32(255)) < base.Ui32(v162) {
				v245 = v162
				v249 = F_sdsnewlen(m, v77, v245)
				mBase = m.M
				v250 = m.ExcPending
				if v250 != 0 {
					return int32(0)
				} else {
					v251 = F_createUnembeddedObjectWithKeyAndExpire(m, int32(0), v249, l1, l2)
					mBase = m.M
					v252 = m.ExcPending
					if v252 != 0 {
						return int32(0)
					} else {
						v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
						v291 = v253
						v292 = v251
						v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
						F_decrRefCount(m, l0)
						mBase = m.M
						v305 = m.ExcPending
						if v305 != 0 {
							return int32(0)
						} else {
							return v292
						}
					}
				}
			} else {
				v165 = v162
				if l1 != 0 {
					v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
					switch v170 & int32(7) {
					case 0:
						v187 = int32(base.Ui32(v170) >> (uint(int32(3)) % 32))
					case 1:
						v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
						v187 = v177
					case 2:
						v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
						v187 = v180
					case 3:
						v183 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
						v187 = v183
					case 4:
						v186 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
						v187 = v186
					default:
						v187 = int32(0)
					}
					if base.Ui32(int32(32)) <= base.Ui32(v187) {
						if base.Ui32(int32(253)) <= base.Ui32(v187) {
							if base.Ui32(v187) < base.Ui32(int32(65531)) {
								v198 = int32(2)
							} else {
								v198 = int32(3)
							}
							v199 = v198
						} else {
							v199 = int32(1)
						}
					} else {
						v199 = int32(0)
					}
					v203 = v199 & int32(7)
					if base.Ui32(int32(4)) < base.Ui32(v203) {
						v211 = int32(0)
					} else {
						v210 = *(*int32)(unsafe.Add(mBase, uint32(v203<<(uint(int32(2))%32))+uint32(_consts[602])))
						v211 = v210
					}
					v215 = v187 + v211 + int32(10)
				} else {
					v215 = int32(8)
				}
				if l2 == int64(-1) {
					v221 = int32(1)
				} else {
					v221 = int32(9)
				}
				v235 = *(*int32)(unsafe.Add(mBase, _consts[601]))
				if base.Ui32(int32(128)) < base.Ui32(v221+v165+v215+v235) {
					v245 = v165
					v249 = F_sdsnewlen(m, v77, v245)
					mBase = m.M
					v250 = m.ExcPending
					if v250 != 0 {
						return int32(0)
					} else {
						v251 = F_createUnembeddedObjectWithKeyAndExpire(m, int32(0), v249, l1, l2)
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return int32(0)
						} else {
							v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
							v291 = v253
							v292 = v251
							v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
							F_decrRefCount(m, l0)
							mBase = m.M
							v305 = m.ExcPending
							if v305 != 0 {
								return int32(0)
							} else {
								return v292
							}
						}
					}
				} else {
					v240 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v77, v165, l1, l2)
					mBase = m.M
					v243 = m.ExcPending
					if v243 != 0 {
						return int32(0)
					} else {
						v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
						v291 = v244
						v292 = v240
						v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
						F_decrRefCount(m, l0)
						mBase = m.M
						v305 = m.ExcPending
						if v305 != 0 {
							return int32(0)
						} else {
							return v292
						}
					}
				}
			}
		case 3:
			v158 = *(*int32)(unsafe.Add(mBase, uint32(v139+int32(-9))))
			v162 = v158
			if base.Ui32(int32(255)) < base.Ui32(v162) {
				v245 = v162
				v249 = F_sdsnewlen(m, v77, v245)
				mBase = m.M
				v250 = m.ExcPending
				if v250 != 0 {
					return int32(0)
				} else {
					v251 = F_createUnembeddedObjectWithKeyAndExpire(m, int32(0), v249, l1, l2)
					mBase = m.M
					v252 = m.ExcPending
					if v252 != 0 {
						return int32(0)
					} else {
						v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
						v291 = v253
						v292 = v251
						v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
						F_decrRefCount(m, l0)
						mBase = m.M
						v305 = m.ExcPending
						if v305 != 0 {
							return int32(0)
						} else {
							return v292
						}
					}
				}
			} else {
				v165 = v162
				if l1 != 0 {
					v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
					switch v170 & int32(7) {
					case 0:
						v187 = int32(base.Ui32(v170) >> (uint(int32(3)) % 32))
					case 1:
						v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
						v187 = v177
					case 2:
						v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
						v187 = v180
					case 3:
						v183 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
						v187 = v183
					case 4:
						v186 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
						v187 = v186
					default:
						v187 = int32(0)
					}
					if base.Ui32(int32(32)) <= base.Ui32(v187) {
						if base.Ui32(int32(253)) <= base.Ui32(v187) {
							if base.Ui32(v187) < base.Ui32(int32(65531)) {
								v198 = int32(2)
							} else {
								v198 = int32(3)
							}
							v199 = v198
						} else {
							v199 = int32(1)
						}
					} else {
						v199 = int32(0)
					}
					v203 = v199 & int32(7)
					if base.Ui32(int32(4)) < base.Ui32(v203) {
						v211 = int32(0)
					} else {
						v210 = *(*int32)(unsafe.Add(mBase, uint32(v203<<(uint(int32(2))%32))+uint32(_consts[602])))
						v211 = v210
					}
					v215 = v187 + v211 + int32(10)
				} else {
					v215 = int32(8)
				}
				if l2 == int64(-1) {
					v221 = int32(1)
				} else {
					v221 = int32(9)
				}
				v235 = *(*int32)(unsafe.Add(mBase, _consts[601]))
				if base.Ui32(int32(128)) < base.Ui32(v221+v165+v215+v235) {
					v245 = v165
					v249 = F_sdsnewlen(m, v77, v245)
					mBase = m.M
					v250 = m.ExcPending
					if v250 != 0 {
						return int32(0)
					} else {
						v251 = F_createUnembeddedObjectWithKeyAndExpire(m, int32(0), v249, l1, l2)
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return int32(0)
						} else {
							v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
							v291 = v253
							v292 = v251
							v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
							F_decrRefCount(m, l0)
							mBase = m.M
							v305 = m.ExcPending
							if v305 != 0 {
								return int32(0)
							} else {
								return v292
							}
						}
					}
				} else {
					v240 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v77, v165, l1, l2)
					mBase = m.M
					v243 = m.ExcPending
					if v243 != 0 {
						return int32(0)
					} else {
						v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
						v291 = v244
						v292 = v240
						v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
						F_decrRefCount(m, l0)
						mBase = m.M
						v305 = m.ExcPending
						if v305 != 0 {
							return int32(0)
						} else {
							return v292
						}
					}
				}
			}
		case 4:
			v161 = *(*int32)(unsafe.Add(mBase, uint32(v139+int32(-17))))
			v162 = v161
			if base.Ui32(int32(255)) < base.Ui32(v162) {
				v245 = v162
				v249 = F_sdsnewlen(m, v77, v245)
				mBase = m.M
				v250 = m.ExcPending
				if v250 != 0 {
					return int32(0)
				} else {
					v251 = F_createUnembeddedObjectWithKeyAndExpire(m, int32(0), v249, l1, l2)
					mBase = m.M
					v252 = m.ExcPending
					if v252 != 0 {
						return int32(0)
					} else {
						v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
						v291 = v253
						v292 = v251
						v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
						F_decrRefCount(m, l0)
						mBase = m.M
						v305 = m.ExcPending
						if v305 != 0 {
							return int32(0)
						} else {
							return v292
						}
					}
				}
			} else {
				v165 = v162
				if l1 != 0 {
					v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
					switch v170 & int32(7) {
					case 0:
						v187 = int32(base.Ui32(v170) >> (uint(int32(3)) % 32))
					case 1:
						v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
						v187 = v177
					case 2:
						v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
						v187 = v180
					case 3:
						v183 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
						v187 = v183
					case 4:
						v186 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
						v187 = v186
					default:
						v187 = int32(0)
					}
					if base.Ui32(int32(32)) <= base.Ui32(v187) {
						if base.Ui32(int32(253)) <= base.Ui32(v187) {
							if base.Ui32(v187) < base.Ui32(int32(65531)) {
								v198 = int32(2)
							} else {
								v198 = int32(3)
							}
							v199 = v198
						} else {
							v199 = int32(1)
						}
					} else {
						v199 = int32(0)
					}
					v203 = v199 & int32(7)
					if base.Ui32(int32(4)) < base.Ui32(v203) {
						v211 = int32(0)
					} else {
						v210 = *(*int32)(unsafe.Add(mBase, uint32(v203<<(uint(int32(2))%32))+uint32(_consts[602])))
						v211 = v210
					}
					v215 = v187 + v211 + int32(10)
				} else {
					v215 = int32(8)
				}
				if l2 == int64(-1) {
					v221 = int32(1)
				} else {
					v221 = int32(9)
				}
				v235 = *(*int32)(unsafe.Add(mBase, _consts[601]))
				if base.Ui32(int32(128)) < base.Ui32(v221+v165+v215+v235) {
					v245 = v165
					v249 = F_sdsnewlen(m, v77, v245)
					mBase = m.M
					v250 = m.ExcPending
					if v250 != 0 {
						return int32(0)
					} else {
						v251 = F_createUnembeddedObjectWithKeyAndExpire(m, int32(0), v249, l1, l2)
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return int32(0)
						} else {
							v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
							v291 = v253
							v292 = v251
							v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
							F_decrRefCount(m, l0)
							mBase = m.M
							v305 = m.ExcPending
							if v305 != 0 {
								return int32(0)
							} else {
								return v292
							}
						}
					}
				} else {
					v240 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v77, v165, l1, l2)
					mBase = m.M
					v243 = m.ExcPending
					if v243 != 0 {
						return int32(0)
					} else {
						v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
						v291 = v244
						v292 = v240
						v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
						F_decrRefCount(m, l0)
						mBase = m.M
						v305 = m.ExcPending
						if v305 != 0 {
							return int32(0)
						} else {
							return v292
						}
					}
				}
			}
		default:
			v165 = int32(0)
			if l1 != 0 {
				v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
				switch v170 & int32(7) {
				case 0:
					v187 = int32(base.Ui32(v170) >> (uint(int32(3)) % 32))
				case 1:
					v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
					v187 = v177
				case 2:
					v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
					v187 = v180
				case 3:
					v183 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
					v187 = v183
				case 4:
					v186 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
					v187 = v186
				default:
					v187 = int32(0)
				}
				if base.Ui32(int32(32)) <= base.Ui32(v187) {
					if base.Ui32(int32(253)) <= base.Ui32(v187) {
						if base.Ui32(v187) < base.Ui32(int32(65531)) {
							v198 = int32(2)
						} else {
							v198 = int32(3)
						}
						v199 = v198
					} else {
						v199 = int32(1)
					}
				} else {
					v199 = int32(0)
				}
				v203 = v199 & int32(7)
				if base.Ui32(int32(4)) < base.Ui32(v203) {
					v211 = int32(0)
				} else {
					v210 = *(*int32)(unsafe.Add(mBase, uint32(v203<<(uint(int32(2))%32))+uint32(_consts[602])))
					v211 = v210
				}
				v215 = v187 + v211 + int32(10)
			} else {
				v215 = int32(8)
			}
			if l2 == int64(-1) {
				v221 = int32(1)
			} else {
				v221 = int32(9)
			}
			v235 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			if base.Ui32(int32(128)) < base.Ui32(v221+v165+v215+v235) {
				v245 = v165
				v249 = F_sdsnewlen(m, v77, v245)
				mBase = m.M
				v250 = m.ExcPending
				if v250 != 0 {
					return int32(0)
				} else {
					v251 = F_createUnembeddedObjectWithKeyAndExpire(m, int32(0), v249, l1, l2)
					mBase = m.M
					v252 = m.ExcPending
					if v252 != 0 {
						return int32(0)
					} else {
						v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
						v291 = v253
						v292 = v251
						v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
						F_decrRefCount(m, l0)
						mBase = m.M
						v305 = m.ExcPending
						if v305 != 0 {
							return int32(0)
						} else {
							return v292
						}
					}
				}
			} else {
				v240 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v77, v165, l1, l2)
				mBase = m.M
				v243 = m.ExcPending
				if v243 != 0 {
					return int32(0)
				} else {
					v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
					v291 = v244
					v292 = v240
					v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v292))) = v291&int32(255) | v299&int32(-256)
					F_decrRefCount(m, l0)
					mBase = m.M
					v305 = m.ExcPending
					if v305 != 0 {
						return int32(0)
					} else {
						return v292
					}
				}
			}
		}
	}
}
func F_rewriteSetObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
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
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_setTypeSize(m, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = F_setTypeInitIterator(m, l2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_setTypeReleaseIterator(m, v17)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L33
	}
L4:
	;
	v101 = int32(1)
	goto L3
L5:
	;
	v23 = F_setTypeNext(m, v17, v11+int32(12), v11+int32(8), v11)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v23 == int32(-1) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v35 = base.I64_extend_i32_u(v13)
	v36 = int64(0)
	goto L8
L8:
	;
	if v36 != int64(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L4
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v63 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	v39 = int32(0)
	v41 = int64(64)
	if v35 < v41 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = v35
	goto L14
L13:
	;
	v44 = v41
	goto L14
L14:
	;
	v48 = F_rioWriteBulkCount(m, l0, int32(42), base.I32_wrap_i64(v44)+int32(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v48 == int32(0) {
		v101 = v39
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v54 = F_rioWriteBulkString(m, l0, int32(_a167), int32(4))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v54 == int32(0) {
		v101 = v39
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v58 = F_rioWriteBulkObject(m, l0, l1)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v58 == int32(0) {
		v101 = v39
		goto L3
	} else {
		goto L20
	}
L20:
	;
	goto L10
L21:
	;
	if v72 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v70 = F_rioWriteBulkLongLong(m, l0, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v67 = F_rioWriteBulkString(m, l0, v63, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v72 = v67
	goto L21
L25:
	;
	v72 = v70
	goto L21
L26:
	;
	v76 = v36 + int64(1)
	if v76 == int64(64) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v101 = int32(0)
	goto L3
L28:
	;
	v79 = int64(0)
	goto L30
L29:
	;
	v79 = v76
	goto L30
L30:
	;
	v86 = F_setTypeNext(m, v17, v11+int32(12), v11+int32(8), v11)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v86 != int32(-1) {
		v35 = v35 + int64(-1)
		v36 = v79
		goto L8
	} else {
		goto L32
	}
L32:
	;
	goto L9
L33:
	;
	m.G0 = v11 + int32(16)
	return v101
}
