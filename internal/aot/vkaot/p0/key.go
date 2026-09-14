package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_getKeySlot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int64
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_getKeySlot[0]))
	if v5 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_getKeySlot_0), int32(_a_F_getKeySlot_1), int32(241))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L46
	} else {
		goto L84
	}
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_getKeySlot[1]))
	if v9 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v258
L4:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v147 & int32(7) {
	case 0:
		goto L53
	case 1:
		goto L52
	case 2:
		goto L51
	case 3:
		goto L50
	case 4:
		goto L49
	default:
		v164 = int32(0)
		goto L48
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+292))
	if v12 < int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+204)))
	if v15&int32(1) == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	if v21 != int64(-1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v36 != 0 {
		goto L4
	} else {
		goto L15
	}
L9:
	;
	v25 = int32(1)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+200)))
	if v26&v25 != 0 {
		v33 = v25
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v36 = int32(1)
	goto L8
L11:
	;
	v36 = v33
	goto L8
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+216))
	if v29 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v31 = F_isImportSlotMigrationJob(m, v29)
	mBase = m.M
	v33 = v31
	goto L11
L14:
	;
	v36 = int32(0)
	goto L8
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_getKeySlot[2]))
	if v38 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v46 & int32(7) {
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
		v63 = int32(0)
		goto L18
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_getKeySlot[1]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+292))
	return v41
L18:
	;
	v64 = int32(0)
	if v63 < int32(1) {
		v84 = v64
		goto L28
	} else {
		goto L29
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v63 = v62
	goto L18
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v63 = v59
	goto L18
L21:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v63 = v56
	goto L18
L22:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v63 = v53
	goto L18
L23:
	;
	v63 = int32(base.Ui32(v46) >> (uint(int32(3)) % 32))
	goto L18
L24:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_getKeySlot[1]))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+292))
	if v130 == v133 {
		v258 = v130
		goto L3
	} else {
		goto L45
	}
L25:
	;
	v130 = v126 & int32(16383)
	goto L24
L26:
	;
	v95 = v84 + int32(1)
	if v63 <= v95 {
		goto L36
	} else {
		goto L37
	}
L27:
	;
	v93 = F_crc16(m, l0, v63)
	mBase = m.M
	v126 = v93
	goto L25
L28:
	;
	if v84 != v63 {
		goto L26
	} else {
		goto L34
	}
L29:
	;
	v72 = v64
	goto L30
L30:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v72))))
	if v76 == int32(123) {
		v84 = v72
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v80 = v72 + int32(1)
	if v80 != v63 {
		v72 = v80
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L27
L34:
	;
	goto L27
L35:
	;
	v123 = F_crc16(m, l0+v84+int32(1), v101+(v84^int32(-1)))
	mBase = m.M
	v126 = v123
	goto L25
L36:
	;
	v116 = F_crc16(m, l0, v63)
	mBase = m.M
	v126 = v116
	goto L25
L37:
	;
	v101 = v95
	goto L39
L38:
	;
	if v101 == v63 {
		goto L36
	} else {
		goto L43
	}
L39:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v101))))
	if v103 == int32(125) {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v107 = v101 + int32(1)
	if v107 != v63 {
		v101 = v107
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L36
L43:
	;
	if v101 != v95 {
		goto L35
	} else {
		goto L44
	}
L44:
	;
	goto L36
L45:
	;
	F__serverAssertWithInfo(m, v132, int32(0), int32(_a_F_getKeySlot_2), int32(_a_F_getKeySlot_1), int32(255))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	return int32(0)
L47:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v165 = int32(0)
	if v164 < int32(1) {
		v185 = v165
		goto L58
	} else {
		goto L59
	}
L49:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v164 = v163
	goto L48
L50:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v164 = v160
	goto L48
L51:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v164 = v157
	goto L48
L52:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v164 = v154
	goto L48
L53:
	;
	v164 = int32(base.Ui32(v147) >> (uint(int32(3)) % 32))
	goto L48
L54:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_getKeySlot[1]))
	if v233 == int32(0) {
		v258 = v231
		goto L3
	} else {
		goto L75
	}
L55:
	;
	v231 = v227 & int32(16383)
	goto L54
L56:
	;
	v196 = v185 + int32(1)
	if v164 <= v196 {
		goto L66
	} else {
		goto L67
	}
L57:
	;
	v194 = F_crc16(m, l0, v164)
	mBase = m.M
	v227 = v194
	goto L55
L58:
	;
	if v185 != v164 {
		goto L56
	} else {
		goto L64
	}
L59:
	;
	v173 = v165
	goto L60
L60:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v173))))
	if v177 == int32(123) {
		v185 = v173
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v181 = v173 + int32(1)
	if v181 != v164 {
		v173 = v181
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L57
L64:
	;
	goto L57
L65:
	;
	v224 = F_crc16(m, l0+v185+int32(1), v202+(v185^int32(-1)))
	mBase = m.M
	v227 = v224
	goto L55
L66:
	;
	v217 = F_crc16(m, l0, v164)
	mBase = m.M
	v227 = v217
	goto L55
L67:
	;
	v202 = v196
	goto L69
L68:
	;
	if v202 == v164 {
		goto L66
	} else {
		goto L73
	}
L69:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v202))))
	if v204 == int32(125) {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v208 = v202 + int32(1)
	if v208 != v164 {
		v202 = v208
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L66
L73:
	;
	if v202 != v196 {
		goto L65
	} else {
		goto L74
	}
L74:
	;
	goto L66
L75:
	;
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v233)))
	if v237 != int64(-1) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if v252 == int32(0) {
		v258 = v231
		goto L3
	} else {
		goto L83
	}
L77:
	;
	v241 = int32(1)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+200)))
	if v242&v241 != 0 {
		v249 = v241
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v252 = int32(1)
	goto L76
L79:
	;
	v252 = v249
	goto L76
L80:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v233)+216))
	if v245 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v247 = F_isImportSlotMigrationJob(m, v245)
	mBase = m.M
	v249 = v247
	goto L79
L82:
	;
	v252 = int32(0)
	goto L76
L83:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_getKeySlot[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+292)) = v231
	v258 = v231
	goto L3
L84:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lookupKeyByPattern(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
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
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	v12 = F_objectGetVal(m, l1)
	mBase = m.M
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 != int32(35) {
		v22 = F_getDecodedObject(m, l2)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_objectGetVal(m, v22)
			mBase = m.M
			v25 = int32(42)
			v26 = F___strchrnul(m, v12, v25)
			mBase = m.M
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
			if v28 == v25 {
				v32 = v26
			} else {
				v32 = int32(0)
			}
			if v32 != 0 {
				v39 = v32 + int32(1)
				v40 = int32(_a_F_lookupKeyByPattern_0)
				v43 = int32(*(*int8)(unsafe.Add(mBase, _c_F_lookupKeyByPattern[0])))
				if v43 != 0 {
					v44 = int32(0)
					v45 = F_strchr(m, v39, v43)
					mBase = m.M
					if v45 == v44 {
						v65 = v44
						v68 = v65
					} else {
						v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lookupKeyByPattern[1])))
						if v48 != 0 {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
							if v49 == int32(0) {
								v65 = v44
								v68 = v65
							} else {
								v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lookupKeyByPattern[2])))
								if v52 != 0 {
									v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+2)))
									if v54 == int32(0) {
										v65 = v44
										v68 = v65
									} else {
										v57 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lookupKeyByPattern[3])))
										if v57 != 0 {
											v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+3)))
											if v59 == int32(0) {
												v65 = v44
												v68 = v65
											} else {
												v62 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lookupKeyByPattern[4])))
												if v62 != 0 {
													v64 = F_twoway_strstr(m, v45, v40)
													mBase = m.M
													v65 = v64
													v68 = v65
												} else {
													v63 = F_fourbyte_strstr(m, v45, v40)
													mBase = m.M
													v68 = v63
												}
											}
										} else {
											v58 = F_threebyte_strstr(m, v45, v40)
											mBase = m.M
											v68 = v58
										}
									}
								} else {
									v53 = F_twobyte_strstr(m, v45, v40)
									mBase = m.M
									v68 = v53
								}
							}
						} else {
							v68 = v45
						}
					}
				} else {
					v68 = v39
				}
				if v68 == int32(0) {
					v103 = int32(0)
					v105 = v103
					v106 = v103
					v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-1)))))
					switch v111 & int32(7) {
					case 0:
						v128 = int32(base.Ui32(v111) >> (uint(int32(3)) % 32))
					case 1:
						v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-3)))))
						v128 = v118
					case 2:
						v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(-5)))))
						v128 = v121
					case 3:
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-9))))
						v128 = v124
					case 4:
						v127 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-17))))
						v128 = v127
					default:
						v128 = int32(0)
					}
					v129 = v32 - v12
					v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-1)))))
					switch v133 & int32(7) {
					case 0:
						v150 = int32(base.Ui32(v133) >> (uint(int32(3)) % 32))
					case 1:
						v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
						v150 = v140
					case 2:
						v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-5)))))
						v150 = v143
					case 3:
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
						v150 = v146
					case 4:
						v149 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-17))))
						v150 = v149
					default:
						v150 = int32(0)
					}
					if v105 != 0 {
						v157 = int32(-3) - v105
					} else {
						v157 = int32(-1)
					}
					v158 = v150 - v129 + v157
					v160 = F_createStringObject_1(m, int32(0), v128+v129+v158)
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return int32(0)
					} else {
						v162 = F_objectGetVal(m, v160)
						mBase = m.M
						if v129 == int32(0) {
							v166 = v162
						} else {
							v165 = F__emscripten_memcpy_bulkmem(m, v162, v12, v129)
							mBase = m.M
							v166 = v165
						}
						v167 = v166 + v129
						if v128 == int32(0) {
							v171 = v167
						} else {
							v170 = F__emscripten_memcpy_bulkmem(m, v167, v24, v128)
							mBase = m.M
							v171 = v170
						}
						if v158 == int32(0) {
						} else {
							v175 = F__emscripten_memcpy_bulkmem(m, v171+v128, v39, v158)
							mBase = m.M
						}
						F_decrRefCount(m, v22)
						mBase = m.M
						v178 = m.ExcPending
						if v178 != 0 {
							return int32(0)
						} else {
							v179 = F_lookupKeyRead(m, l0, v160)
							mBase = m.M
							v180 = m.ExcPending
							if v180 != 0 {
								return int32(0)
							} else {
								if v179 == int32(0) {
									F_decrRefCount(m, v160)
									mBase = m.M
									v202 = m.ExcPending
									if v202 != 0 {
										return int32(0)
									} else {
										v203 = int32(0)
										if v105 == v203 {
											v210 = v203
											return v210
										} else {
											v206 = v203
											F_decrRefCount(m, v106)
											mBase = m.M
											v209 = m.ExcPending
											if v209 != 0 {
												return int32(0)
											} else {
												v210 = v206
												return v210
											}
										}
									}
								} else {
									v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
									v185 = v183 & int32(15)
									if v106 == int32(0) {
										if v185 != 0 {
											F_decrRefCount(m, v160)
											mBase = m.M
											v202 = m.ExcPending
											if v202 != 0 {
												return int32(0)
											} else {
												v203 = int32(0)
												if v105 == v203 {
													v210 = v203
													return v210
												} else {
													v206 = v203
													F_decrRefCount(m, v106)
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return int32(0)
													} else {
														v210 = v206
														return v210
													}
												}
											}
										} else {
											F_incrRefCount(m, v179)
											mBase = m.M
											v196 = m.ExcPending
											if v196 != 0 {
												return int32(0)
											} else {
												F_decrRefCount(m, v160)
												mBase = m.M
												v198 = m.ExcPending
												if v198 != 0 {
													return int32(0)
												} else {
													return v179
												}
											}
										}
									} else {
										if v185 != int32(4) {
											F_decrRefCount(m, v160)
											mBase = m.M
											v202 = m.ExcPending
											if v202 != 0 {
												return int32(0)
											} else {
												v203 = int32(0)
												if v105 == v203 {
													v210 = v203
													return v210
												} else {
													v206 = v203
													F_decrRefCount(m, v106)
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return int32(0)
													} else {
														v210 = v206
														return v210
													}
												}
											}
										} else {
											v190 = F_objectGetVal(m, v106)
											mBase = m.M
											v191 = F_hashTypeGetValueObject(m, v179, v190)
											mBase = m.M
											v192 = m.ExcPending
											if v192 != 0 {
												return int32(0)
											} else {
												F_decrRefCount(m, v160)
												mBase = m.M
												v194 = m.ExcPending
												if v194 != 0 {
													return int32(0)
												} else {
													v206 = v191
													F_decrRefCount(m, v106)
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return int32(0)
													} else {
														v210 = v206
														return v210
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
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+2)))
					if v71 == int32(0) {
						v103 = int32(0)
						v105 = v103
						v106 = v103
						v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-1)))))
						switch v111 & int32(7) {
						case 0:
							v128 = int32(base.Ui32(v111) >> (uint(int32(3)) % 32))
						case 1:
							v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-3)))))
							v128 = v118
						case 2:
							v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(-5)))))
							v128 = v121
						case 3:
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-9))))
							v128 = v124
						case 4:
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-17))))
							v128 = v127
						default:
							v128 = int32(0)
						}
						v129 = v32 - v12
						v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-1)))))
						switch v133 & int32(7) {
						case 0:
							v150 = int32(base.Ui32(v133) >> (uint(int32(3)) % 32))
						case 1:
							v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
							v150 = v140
						case 2:
							v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-5)))))
							v150 = v143
						case 3:
							v146 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
							v150 = v146
						case 4:
							v149 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-17))))
							v150 = v149
						default:
							v150 = int32(0)
						}
						if v105 != 0 {
							v157 = int32(-3) - v105
						} else {
							v157 = int32(-1)
						}
						v158 = v150 - v129 + v157
						v160 = F_createStringObject_1(m, int32(0), v128+v129+v158)
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return int32(0)
						} else {
							v162 = F_objectGetVal(m, v160)
							mBase = m.M
							if v129 == int32(0) {
								v166 = v162
							} else {
								v165 = F__emscripten_memcpy_bulkmem(m, v162, v12, v129)
								mBase = m.M
								v166 = v165
							}
							v167 = v166 + v129
							if v128 == int32(0) {
								v171 = v167
							} else {
								v170 = F__emscripten_memcpy_bulkmem(m, v167, v24, v128)
								mBase = m.M
								v171 = v170
							}
							if v158 == int32(0) {
							} else {
								v175 = F__emscripten_memcpy_bulkmem(m, v171+v128, v39, v158)
								mBase = m.M
							}
							F_decrRefCount(m, v22)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = F_lookupKeyRead(m, l0, v160)
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return int32(0)
								} else {
									if v179 == int32(0) {
										F_decrRefCount(m, v160)
										mBase = m.M
										v202 = m.ExcPending
										if v202 != 0 {
											return int32(0)
										} else {
											v203 = int32(0)
											if v105 == v203 {
												v210 = v203
												return v210
											} else {
												v206 = v203
												F_decrRefCount(m, v106)
												mBase = m.M
												v209 = m.ExcPending
												if v209 != 0 {
													return int32(0)
												} else {
													v210 = v206
													return v210
												}
											}
										}
									} else {
										v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
										v185 = v183 & int32(15)
										if v106 == int32(0) {
											if v185 != 0 {
												F_decrRefCount(m, v160)
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return int32(0)
												} else {
													v203 = int32(0)
													if v105 == v203 {
														v210 = v203
														return v210
													} else {
														v206 = v203
														F_decrRefCount(m, v106)
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return int32(0)
														} else {
															v210 = v206
															return v210
														}
													}
												}
											} else {
												F_incrRefCount(m, v179)
												mBase = m.M
												v196 = m.ExcPending
												if v196 != 0 {
													return int32(0)
												} else {
													F_decrRefCount(m, v160)
													mBase = m.M
													v198 = m.ExcPending
													if v198 != 0 {
														return int32(0)
													} else {
														return v179
													}
												}
											}
										} else {
											if v185 != int32(4) {
												F_decrRefCount(m, v160)
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return int32(0)
												} else {
													v203 = int32(0)
													if v105 == v203 {
														v210 = v203
														return v210
													} else {
														v206 = v203
														F_decrRefCount(m, v106)
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return int32(0)
														} else {
															v210 = v206
															return v210
														}
													}
												}
											} else {
												v190 = F_objectGetVal(m, v106)
												mBase = m.M
												v191 = F_hashTypeGetValueObject(m, v179, v190)
												mBase = m.M
												v192 = m.ExcPending
												if v192 != 0 {
													return int32(0)
												} else {
													F_decrRefCount(m, v160)
													mBase = m.M
													v194 = m.ExcPending
													if v194 != 0 {
														return int32(0)
													} else {
														v206 = v191
														F_decrRefCount(m, v106)
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return int32(0)
														} else {
															v210 = v206
															return v210
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
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-1)))))
						switch v79 & int32(7) {
						case 0:
							v96 = int32(base.Ui32(v79) >> (uint(int32(3)) % 32))
						case 1:
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
							v96 = v86
						case 2:
							v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-5)))))
							v96 = v89
						case 3:
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
							v96 = v92
						case 4:
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-17))))
							v96 = v95
						default:
							v96 = int32(0)
						}
						v100 = v12 - v68 + v96 + int32(-2)
						v101 = F_createStringObject_1(m, v68+int32(2), v100)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							v105 = v100
							v106 = v101
							v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-1)))))
							switch v111 & int32(7) {
							case 0:
								v128 = int32(base.Ui32(v111) >> (uint(int32(3)) % 32))
							case 1:
								v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-3)))))
								v128 = v118
							case 2:
								v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(-5)))))
								v128 = v121
							case 3:
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-9))))
								v128 = v124
							case 4:
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-17))))
								v128 = v127
							default:
								v128 = int32(0)
							}
							v129 = v32 - v12
							v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-1)))))
							switch v133 & int32(7) {
							case 0:
								v150 = int32(base.Ui32(v133) >> (uint(int32(3)) % 32))
							case 1:
								v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
								v150 = v140
							case 2:
								v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-5)))))
								v150 = v143
							case 3:
								v146 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
								v150 = v146
							case 4:
								v149 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-17))))
								v150 = v149
							default:
								v150 = int32(0)
							}
							if v105 != 0 {
								v157 = int32(-3) - v105
							} else {
								v157 = int32(-1)
							}
							v158 = v150 - v129 + v157
							v160 = F_createStringObject_1(m, int32(0), v128+v129+v158)
							mBase = m.M
							v161 = m.ExcPending
							if v161 != 0 {
								return int32(0)
							} else {
								v162 = F_objectGetVal(m, v160)
								mBase = m.M
								if v129 == int32(0) {
									v166 = v162
								} else {
									v165 = F__emscripten_memcpy_bulkmem(m, v162, v12, v129)
									mBase = m.M
									v166 = v165
								}
								v167 = v166 + v129
								if v128 == int32(0) {
									v171 = v167
								} else {
									v170 = F__emscripten_memcpy_bulkmem(m, v167, v24, v128)
									mBase = m.M
									v171 = v170
								}
								if v158 == int32(0) {
								} else {
									v175 = F__emscripten_memcpy_bulkmem(m, v171+v128, v39, v158)
									mBase = m.M
								}
								F_decrRefCount(m, v22)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = F_lookupKeyRead(m, l0, v160)
									mBase = m.M
									v180 = m.ExcPending
									if v180 != 0 {
										return int32(0)
									} else {
										if v179 == int32(0) {
											F_decrRefCount(m, v160)
											mBase = m.M
											v202 = m.ExcPending
											if v202 != 0 {
												return int32(0)
											} else {
												v203 = int32(0)
												if v105 == v203 {
													v210 = v203
													return v210
												} else {
													v206 = v203
													F_decrRefCount(m, v106)
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return int32(0)
													} else {
														v210 = v206
														return v210
													}
												}
											}
										} else {
											v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
											v185 = v183 & int32(15)
											if v106 == int32(0) {
												if v185 != 0 {
													F_decrRefCount(m, v160)
													mBase = m.M
													v202 = m.ExcPending
													if v202 != 0 {
														return int32(0)
													} else {
														v203 = int32(0)
														if v105 == v203 {
															v210 = v203
															return v210
														} else {
															v206 = v203
															F_decrRefCount(m, v106)
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return int32(0)
															} else {
																v210 = v206
																return v210
															}
														}
													}
												} else {
													F_incrRefCount(m, v179)
													mBase = m.M
													v196 = m.ExcPending
													if v196 != 0 {
														return int32(0)
													} else {
														F_decrRefCount(m, v160)
														mBase = m.M
														v198 = m.ExcPending
														if v198 != 0 {
															return int32(0)
														} else {
															return v179
														}
													}
												}
											} else {
												if v185 != int32(4) {
													F_decrRefCount(m, v160)
													mBase = m.M
													v202 = m.ExcPending
													if v202 != 0 {
														return int32(0)
													} else {
														v203 = int32(0)
														if v105 == v203 {
															v210 = v203
															return v210
														} else {
															v206 = v203
															F_decrRefCount(m, v106)
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return int32(0)
															} else {
																v210 = v206
																return v210
															}
														}
													}
												} else {
													v190 = F_objectGetVal(m, v106)
													mBase = m.M
													v191 = F_hashTypeGetValueObject(m, v179, v190)
													mBase = m.M
													v192 = m.ExcPending
													if v192 != 0 {
														return int32(0)
													} else {
														F_decrRefCount(m, v160)
														mBase = m.M
														v194 = m.ExcPending
														if v194 != 0 {
															return int32(0)
														} else {
															v206 = v191
															F_decrRefCount(m, v106)
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return int32(0)
															} else {
																v210 = v206
																return v210
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
				F_decrRefCount(m, v22)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
		if v16 != 0 {
			v22 = F_getDecodedObject(m, l2)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = F_objectGetVal(m, v22)
				mBase = m.M
				v25 = int32(42)
				v26 = F___strchrnul(m, v12, v25)
				mBase = m.M
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
				if v28 == v25 {
					v32 = v26
				} else {
					v32 = int32(0)
				}
				if v32 != 0 {
					v39 = v32 + int32(1)
					v40 = int32(_a_F_lookupKeyByPattern_0)
					v43 = int32(*(*int8)(unsafe.Add(mBase, _c_F_lookupKeyByPattern[0])))
					if v43 != 0 {
						v44 = int32(0)
						v45 = F_strchr(m, v39, v43)
						mBase = m.M
						if v45 == v44 {
							v65 = v44
							v68 = v65
						} else {
							v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lookupKeyByPattern[1])))
							if v48 != 0 {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
								if v49 == int32(0) {
									v65 = v44
									v68 = v65
								} else {
									v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lookupKeyByPattern[2])))
									if v52 != 0 {
										v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+2)))
										if v54 == int32(0) {
											v65 = v44
											v68 = v65
										} else {
											v57 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lookupKeyByPattern[3])))
											if v57 != 0 {
												v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+3)))
												if v59 == int32(0) {
													v65 = v44
													v68 = v65
												} else {
													v62 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lookupKeyByPattern[4])))
													if v62 != 0 {
														v64 = F_twoway_strstr(m, v45, v40)
														mBase = m.M
														v65 = v64
														v68 = v65
													} else {
														v63 = F_fourbyte_strstr(m, v45, v40)
														mBase = m.M
														v68 = v63
													}
												}
											} else {
												v58 = F_threebyte_strstr(m, v45, v40)
												mBase = m.M
												v68 = v58
											}
										}
									} else {
										v53 = F_twobyte_strstr(m, v45, v40)
										mBase = m.M
										v68 = v53
									}
								}
							} else {
								v68 = v45
							}
						}
					} else {
						v68 = v39
					}
					if v68 == int32(0) {
						v103 = int32(0)
						v105 = v103
						v106 = v103
						v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-1)))))
						switch v111 & int32(7) {
						case 0:
							v128 = int32(base.Ui32(v111) >> (uint(int32(3)) % 32))
						case 1:
							v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-3)))))
							v128 = v118
						case 2:
							v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(-5)))))
							v128 = v121
						case 3:
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-9))))
							v128 = v124
						case 4:
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-17))))
							v128 = v127
						default:
							v128 = int32(0)
						}
						v129 = v32 - v12
						v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-1)))))
						switch v133 & int32(7) {
						case 0:
							v150 = int32(base.Ui32(v133) >> (uint(int32(3)) % 32))
						case 1:
							v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
							v150 = v140
						case 2:
							v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-5)))))
							v150 = v143
						case 3:
							v146 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
							v150 = v146
						case 4:
							v149 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-17))))
							v150 = v149
						default:
							v150 = int32(0)
						}
						if v105 != 0 {
							v157 = int32(-3) - v105
						} else {
							v157 = int32(-1)
						}
						v158 = v150 - v129 + v157
						v160 = F_createStringObject_1(m, int32(0), v128+v129+v158)
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return int32(0)
						} else {
							v162 = F_objectGetVal(m, v160)
							mBase = m.M
							if v129 == int32(0) {
								v166 = v162
							} else {
								v165 = F__emscripten_memcpy_bulkmem(m, v162, v12, v129)
								mBase = m.M
								v166 = v165
							}
							v167 = v166 + v129
							if v128 == int32(0) {
								v171 = v167
							} else {
								v170 = F__emscripten_memcpy_bulkmem(m, v167, v24, v128)
								mBase = m.M
								v171 = v170
							}
							if v158 == int32(0) {
							} else {
								v175 = F__emscripten_memcpy_bulkmem(m, v171+v128, v39, v158)
								mBase = m.M
							}
							F_decrRefCount(m, v22)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = F_lookupKeyRead(m, l0, v160)
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return int32(0)
								} else {
									if v179 == int32(0) {
										F_decrRefCount(m, v160)
										mBase = m.M
										v202 = m.ExcPending
										if v202 != 0 {
											return int32(0)
										} else {
											v203 = int32(0)
											if v105 == v203 {
												v210 = v203
												return v210
											} else {
												v206 = v203
												F_decrRefCount(m, v106)
												mBase = m.M
												v209 = m.ExcPending
												if v209 != 0 {
													return int32(0)
												} else {
													v210 = v206
													return v210
												}
											}
										}
									} else {
										v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
										v185 = v183 & int32(15)
										if v106 == int32(0) {
											if v185 != 0 {
												F_decrRefCount(m, v160)
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return int32(0)
												} else {
													v203 = int32(0)
													if v105 == v203 {
														v210 = v203
														return v210
													} else {
														v206 = v203
														F_decrRefCount(m, v106)
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return int32(0)
														} else {
															v210 = v206
															return v210
														}
													}
												}
											} else {
												F_incrRefCount(m, v179)
												mBase = m.M
												v196 = m.ExcPending
												if v196 != 0 {
													return int32(0)
												} else {
													F_decrRefCount(m, v160)
													mBase = m.M
													v198 = m.ExcPending
													if v198 != 0 {
														return int32(0)
													} else {
														return v179
													}
												}
											}
										} else {
											if v185 != int32(4) {
												F_decrRefCount(m, v160)
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return int32(0)
												} else {
													v203 = int32(0)
													if v105 == v203 {
														v210 = v203
														return v210
													} else {
														v206 = v203
														F_decrRefCount(m, v106)
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return int32(0)
														} else {
															v210 = v206
															return v210
														}
													}
												}
											} else {
												v190 = F_objectGetVal(m, v106)
												mBase = m.M
												v191 = F_hashTypeGetValueObject(m, v179, v190)
												mBase = m.M
												v192 = m.ExcPending
												if v192 != 0 {
													return int32(0)
												} else {
													F_decrRefCount(m, v160)
													mBase = m.M
													v194 = m.ExcPending
													if v194 != 0 {
														return int32(0)
													} else {
														v206 = v191
														F_decrRefCount(m, v106)
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return int32(0)
														} else {
															v210 = v206
															return v210
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
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+2)))
						if v71 == int32(0) {
							v103 = int32(0)
							v105 = v103
							v106 = v103
							v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-1)))))
							switch v111 & int32(7) {
							case 0:
								v128 = int32(base.Ui32(v111) >> (uint(int32(3)) % 32))
							case 1:
								v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-3)))))
								v128 = v118
							case 2:
								v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(-5)))))
								v128 = v121
							case 3:
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-9))))
								v128 = v124
							case 4:
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-17))))
								v128 = v127
							default:
								v128 = int32(0)
							}
							v129 = v32 - v12
							v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-1)))))
							switch v133 & int32(7) {
							case 0:
								v150 = int32(base.Ui32(v133) >> (uint(int32(3)) % 32))
							case 1:
								v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
								v150 = v140
							case 2:
								v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-5)))))
								v150 = v143
							case 3:
								v146 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
								v150 = v146
							case 4:
								v149 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-17))))
								v150 = v149
							default:
								v150 = int32(0)
							}
							if v105 != 0 {
								v157 = int32(-3) - v105
							} else {
								v157 = int32(-1)
							}
							v158 = v150 - v129 + v157
							v160 = F_createStringObject_1(m, int32(0), v128+v129+v158)
							mBase = m.M
							v161 = m.ExcPending
							if v161 != 0 {
								return int32(0)
							} else {
								v162 = F_objectGetVal(m, v160)
								mBase = m.M
								if v129 == int32(0) {
									v166 = v162
								} else {
									v165 = F__emscripten_memcpy_bulkmem(m, v162, v12, v129)
									mBase = m.M
									v166 = v165
								}
								v167 = v166 + v129
								if v128 == int32(0) {
									v171 = v167
								} else {
									v170 = F__emscripten_memcpy_bulkmem(m, v167, v24, v128)
									mBase = m.M
									v171 = v170
								}
								if v158 == int32(0) {
								} else {
									v175 = F__emscripten_memcpy_bulkmem(m, v171+v128, v39, v158)
									mBase = m.M
								}
								F_decrRefCount(m, v22)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = F_lookupKeyRead(m, l0, v160)
									mBase = m.M
									v180 = m.ExcPending
									if v180 != 0 {
										return int32(0)
									} else {
										if v179 == int32(0) {
											F_decrRefCount(m, v160)
											mBase = m.M
											v202 = m.ExcPending
											if v202 != 0 {
												return int32(0)
											} else {
												v203 = int32(0)
												if v105 == v203 {
													v210 = v203
													return v210
												} else {
													v206 = v203
													F_decrRefCount(m, v106)
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return int32(0)
													} else {
														v210 = v206
														return v210
													}
												}
											}
										} else {
											v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
											v185 = v183 & int32(15)
											if v106 == int32(0) {
												if v185 != 0 {
													F_decrRefCount(m, v160)
													mBase = m.M
													v202 = m.ExcPending
													if v202 != 0 {
														return int32(0)
													} else {
														v203 = int32(0)
														if v105 == v203 {
															v210 = v203
															return v210
														} else {
															v206 = v203
															F_decrRefCount(m, v106)
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return int32(0)
															} else {
																v210 = v206
																return v210
															}
														}
													}
												} else {
													F_incrRefCount(m, v179)
													mBase = m.M
													v196 = m.ExcPending
													if v196 != 0 {
														return int32(0)
													} else {
														F_decrRefCount(m, v160)
														mBase = m.M
														v198 = m.ExcPending
														if v198 != 0 {
															return int32(0)
														} else {
															return v179
														}
													}
												}
											} else {
												if v185 != int32(4) {
													F_decrRefCount(m, v160)
													mBase = m.M
													v202 = m.ExcPending
													if v202 != 0 {
														return int32(0)
													} else {
														v203 = int32(0)
														if v105 == v203 {
															v210 = v203
															return v210
														} else {
															v206 = v203
															F_decrRefCount(m, v106)
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return int32(0)
															} else {
																v210 = v206
																return v210
															}
														}
													}
												} else {
													v190 = F_objectGetVal(m, v106)
													mBase = m.M
													v191 = F_hashTypeGetValueObject(m, v179, v190)
													mBase = m.M
													v192 = m.ExcPending
													if v192 != 0 {
														return int32(0)
													} else {
														F_decrRefCount(m, v160)
														mBase = m.M
														v194 = m.ExcPending
														if v194 != 0 {
															return int32(0)
														} else {
															v206 = v191
															F_decrRefCount(m, v106)
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return int32(0)
															} else {
																v210 = v206
																return v210
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
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-1)))))
							switch v79 & int32(7) {
							case 0:
								v96 = int32(base.Ui32(v79) >> (uint(int32(3)) % 32))
							case 1:
								v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
								v96 = v86
							case 2:
								v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-5)))))
								v96 = v89
							case 3:
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
								v96 = v92
							case 4:
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-17))))
								v96 = v95
							default:
								v96 = int32(0)
							}
							v100 = v12 - v68 + v96 + int32(-2)
							v101 = F_createStringObject_1(m, v68+int32(2), v100)
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int32(0)
							} else {
								v105 = v100
								v106 = v101
								v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-1)))))
								switch v111 & int32(7) {
								case 0:
									v128 = int32(base.Ui32(v111) >> (uint(int32(3)) % 32))
								case 1:
									v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-3)))))
									v128 = v118
								case 2:
									v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(-5)))))
									v128 = v121
								case 3:
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-9))))
									v128 = v124
								case 4:
									v127 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-17))))
									v128 = v127
								default:
									v128 = int32(0)
								}
								v129 = v32 - v12
								v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-1)))))
								switch v133 & int32(7) {
								case 0:
									v150 = int32(base.Ui32(v133) >> (uint(int32(3)) % 32))
								case 1:
									v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
									v150 = v140
								case 2:
									v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-5)))))
									v150 = v143
								case 3:
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
									v150 = v146
								case 4:
									v149 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-17))))
									v150 = v149
								default:
									v150 = int32(0)
								}
								if v105 != 0 {
									v157 = int32(-3) - v105
								} else {
									v157 = int32(-1)
								}
								v158 = v150 - v129 + v157
								v160 = F_createStringObject_1(m, int32(0), v128+v129+v158)
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return int32(0)
								} else {
									v162 = F_objectGetVal(m, v160)
									mBase = m.M
									if v129 == int32(0) {
										v166 = v162
									} else {
										v165 = F__emscripten_memcpy_bulkmem(m, v162, v12, v129)
										mBase = m.M
										v166 = v165
									}
									v167 = v166 + v129
									if v128 == int32(0) {
										v171 = v167
									} else {
										v170 = F__emscripten_memcpy_bulkmem(m, v167, v24, v128)
										mBase = m.M
										v171 = v170
									}
									if v158 == int32(0) {
									} else {
										v175 = F__emscripten_memcpy_bulkmem(m, v171+v128, v39, v158)
										mBase = m.M
									}
									F_decrRefCount(m, v22)
									mBase = m.M
									v178 = m.ExcPending
									if v178 != 0 {
										return int32(0)
									} else {
										v179 = F_lookupKeyRead(m, l0, v160)
										mBase = m.M
										v180 = m.ExcPending
										if v180 != 0 {
											return int32(0)
										} else {
											if v179 == int32(0) {
												F_decrRefCount(m, v160)
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return int32(0)
												} else {
													v203 = int32(0)
													if v105 == v203 {
														v210 = v203
														return v210
													} else {
														v206 = v203
														F_decrRefCount(m, v106)
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return int32(0)
														} else {
															v210 = v206
															return v210
														}
													}
												}
											} else {
												v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
												v185 = v183 & int32(15)
												if v106 == int32(0) {
													if v185 != 0 {
														F_decrRefCount(m, v160)
														mBase = m.M
														v202 = m.ExcPending
														if v202 != 0 {
															return int32(0)
														} else {
															v203 = int32(0)
															if v105 == v203 {
																v210 = v203
																return v210
															} else {
																v206 = v203
																F_decrRefCount(m, v106)
																mBase = m.M
																v209 = m.ExcPending
																if v209 != 0 {
																	return int32(0)
																} else {
																	v210 = v206
																	return v210
																}
															}
														}
													} else {
														F_incrRefCount(m, v179)
														mBase = m.M
														v196 = m.ExcPending
														if v196 != 0 {
															return int32(0)
														} else {
															F_decrRefCount(m, v160)
															mBase = m.M
															v198 = m.ExcPending
															if v198 != 0 {
																return int32(0)
															} else {
																return v179
															}
														}
													}
												} else {
													if v185 != int32(4) {
														F_decrRefCount(m, v160)
														mBase = m.M
														v202 = m.ExcPending
														if v202 != 0 {
															return int32(0)
														} else {
															v203 = int32(0)
															if v105 == v203 {
																v210 = v203
																return v210
															} else {
																v206 = v203
																F_decrRefCount(m, v106)
																mBase = m.M
																v209 = m.ExcPending
																if v209 != 0 {
																	return int32(0)
																} else {
																	v210 = v206
																	return v210
																}
															}
														}
													} else {
														v190 = F_objectGetVal(m, v106)
														mBase = m.M
														v191 = F_hashTypeGetValueObject(m, v179, v190)
														mBase = m.M
														v192 = m.ExcPending
														if v192 != 0 {
															return int32(0)
														} else {
															F_decrRefCount(m, v160)
															mBase = m.M
															v194 = m.ExcPending
															if v194 != 0 {
																return int32(0)
															} else {
																v206 = v191
																F_decrRefCount(m, v106)
																mBase = m.M
																v209 = m.ExcPending
																if v209 != 0 {
																	return int32(0)
																} else {
																	v210 = v206
																	return v210
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
					F_decrRefCount(m, v22)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				}
			}
		} else {
			F_incrRefCount(m, l2)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return l2
			}
		}
	}
}
func F_lookupKeyRead(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_lookupKey(m, l0, l1, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_signalKeyAsReady(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v6 int32
	_ = v6
	F_signalKeyAsReadyLogic(m, l0, l1, l2, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_signalKeyAsReadyLogic(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = l2 + int32(-1)
	if base.Ui32(int32(5)) < base.Ui32(v11) {
		m.G0 = v8 + int32(16)
		return
	} else {
		if int32(base.Ui32(int32(53))>>(uint(v11&int32(255))%32))&int32(1) == int32(0) {
			m.G0 = v8 + int32(16)
			return
		} else {
			v23 = int32(2)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v11<<(uint(v23)%32))+uint32(_c_F_signalKeyAsReadyLogic[0])))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v27<<(uint(v23)%32))+uint32(_c_F_signalKeyAsReadyLogic[1])))
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_signalKeyAsReadyLogic[2]))
			if v31|v33 == int32(0) {
				m.G0 = v8 + int32(16)
				return
			} else {
				if l3 == int32(0) {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v43 = F_dictFind(m, v42, l1)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						if v43 == int32(0) {
							m.G0 = v8 + int32(16)
							return
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v50 = F_dictAddRaw(m, v47, l1, v8+int32(12))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								if v50 == int32(0) {
									m.G0 = v8 + int32(16)
									return
								} else {
									F_incrRefCount(m, l1)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										v57 = F_valkey_malloc(m, int32(8))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v57))) = l0
											*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = l1
											F_incrRefCount(m, l1)
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return
											} else {
												v64 = *(*int32)(unsafe.Add(mBase, _c_F_signalKeyAsReadyLogic[3]))
												v65 = F_listAddNodeTail(m, v64, v57)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return
												} else {
													m.G0 = v8 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v40 = F_dictFind(m, v39, l1)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						if v40 != 0 {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v50 = F_dictAddRaw(m, v47, l1, v8+int32(12))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								if v50 == int32(0) {
									m.G0 = v8 + int32(16)
									return
								} else {
									F_incrRefCount(m, l1)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										v57 = F_valkey_malloc(m, int32(8))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v57))) = l0
											*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = l1
											F_incrRefCount(m, l1)
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return
											} else {
												v64 = *(*int32)(unsafe.Add(mBase, _c_F_signalKeyAsReadyLogic[3]))
												v65 = F_listAddNodeTail(m, v64, v57)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return
												} else {
													m.G0 = v8 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
