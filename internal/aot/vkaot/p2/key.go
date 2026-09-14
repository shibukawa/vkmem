package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_keyHashSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	v3 = int32(0)
	if l1 < int32(1) {
		v23 = v3
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v317 & int32(16383)
L2:
	;
	v118 = v23 + int32(1)
	if l1 <= v118 {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	if int32(1) <= l1 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	if v23 != l1 {
		goto L2
	} else {
		goto L10
	}
L5:
	;
	v11 = v3
	goto L6
L6:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v11))))
	if v15 == int32(123) {
		v23 = v11
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v19 = v11 + int32(1)
	if v19 != l1 {
		v11 = v19
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L3
L10:
	;
	goto L3
L11:
	;
	v317 = v110 & int32(65535)
	goto L1
L12:
	;
	goto L11
L13:
	;
	v39 = int32(1)
	if l1 != v39 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v110 = int32(0)
	goto L12
L15:
	;
	if l1&v39 == int32(0) {
		v110 = v88
		goto L12
	} else {
		goto L21
	}
L16:
	;
	v46 = int32(0)
	v48 = l0
	v49 = v46
	v52 = v46
	goto L18
L17:
	;
	v87 = l0
	v88 = int32(0)
	goto L15
L18:
	;
	v54 = int32(65280)
	v56 = int32(8)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v60 = int32(1)
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32(v49&v54)>>(uint(v56)%32))^v58)<<(uint(v60)%32))+uint32(_consts[126]))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32((v64^v49<<(uint(v56)%32))&v54)>>(uint(v56)%32))^v72)<<(uint(v60)%32))+uint32(_consts[126]))))
	v81 = v78 ^ v64<<(uint(v56)%32)
	v82 = int32(2)
	v83 = v48 + v82
	v85 = v52 + v82
	if v85 != l1&int32(2147483646) {
		v48 = v83
		v49 = v81
		v52 = v85
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v87 = v83
	v88 = v81
	goto L15
L20:
	;
	goto L19
L21:
	;
	v97 = int32(8)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32(v88&int32(65280))>>(uint(v97)%32))^v99)<<(uint(int32(1))%32))+uint32(_consts[126]))))
	v110 = v105 ^ v88<<(uint(v97)%32)
	goto L12
L22:
	;
	v225 = int32(1)
	v226 = l0 + v23 + v225
	v229 = v124 + (v23 ^ int32(-1))
	if v225 <= v229 {
		goto L45
	} else {
		goto L46
	}
L23:
	;
	if int32(1) <= l1 {
		goto L34
	} else {
		goto L35
	}
L24:
	;
	v124 = v118
	goto L26
L25:
	;
	if v124 == l1 {
		goto L23
	} else {
		goto L30
	}
L26:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v124))))
	if v126 == int32(125) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v130 = v124 + int32(1)
	if v130 != l1 {
		v124 = v130
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	if v124 != v118 {
		goto L22
	} else {
		goto L31
	}
L31:
	;
	goto L23
L32:
	;
	v317 = v217 & int32(65535)
	goto L1
L33:
	;
	goto L32
L34:
	;
	v146 = int32(1)
	if l1 != v146 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v217 = int32(0)
	goto L33
L36:
	;
	if l1&v146 == int32(0) {
		v217 = v195
		goto L33
	} else {
		goto L42
	}
L37:
	;
	v153 = int32(0)
	v155 = l0
	v156 = v153
	v159 = v153
	goto L39
L38:
	;
	v194 = l0
	v195 = int32(0)
	goto L36
L39:
	;
	v161 = int32(65280)
	v163 = int32(8)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	v167 = int32(1)
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32(v156&v161)>>(uint(v163)%32))^v165)<<(uint(v167)%32))+uint32(_consts[126]))))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32((v171^v156<<(uint(v163)%32))&v161)>>(uint(v163)%32))^v179)<<(uint(v167)%32))+uint32(_consts[126]))))
	v188 = v185 ^ v171<<(uint(v163)%32)
	v189 = int32(2)
	v190 = v155 + v189
	v192 = v159 + v189
	if v192 != l1&int32(2147483646) {
		v155 = v190
		v156 = v188
		v159 = v192
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v194 = v190
	v195 = v188
	goto L36
L41:
	;
	goto L40
L42:
	;
	v204 = int32(8)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32(v195&int32(65280))>>(uint(v204)%32))^v206)<<(uint(int32(1))%32))+uint32(_consts[126]))))
	v217 = v212 ^ v195<<(uint(v204)%32)
	goto L33
L43:
	;
	v317 = v308 & int32(65535)
	goto L1
L44:
	;
	goto L43
L45:
	;
	v237 = int32(1)
	if v229 != v237 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v308 = int32(0)
	goto L44
L47:
	;
	if v229&v237 == int32(0) {
		v308 = v286
		goto L44
	} else {
		goto L53
	}
L48:
	;
	v244 = int32(0)
	v246 = v226
	v247 = v244
	v250 = v244
	goto L50
L49:
	;
	v285 = v226
	v286 = int32(0)
	goto L47
L50:
	;
	v252 = int32(65280)
	v254 = int32(8)
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	v258 = int32(1)
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32(v247&v252)>>(uint(v254)%32))^v256)<<(uint(v258)%32))+uint32(_consts[126]))))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)))
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32((v262^v247<<(uint(v254)%32))&v252)>>(uint(v254)%32))^v270)<<(uint(v258)%32))+uint32(_consts[126]))))
	v279 = v276 ^ v262<<(uint(v254)%32)
	v280 = int32(2)
	v281 = v246 + v280
	v283 = v250 + v280
	if v283 != v229&int32(2147483646) {
		v246 = v281
		v247 = v279
		v250 = v283
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v285 = v281
	v286 = v279
	goto L47
L52:
	;
	goto L51
L53:
	;
	v295 = int32(8)
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32(v286&int32(65280))>>(uint(v295)%32))^v297)<<(uint(int32(1))%32))+uint32(_consts[126]))))
	v308 = v303 ^ v286<<(uint(v295)%32)
	goto L44
}
func F_keyIsExpired(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_objectGetVal(m, l1)
	mBase = m.M
	v16 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v16 == v3 {
		v23 = v3
		v25 = *(*int32)(unsafe.Add(mBase, _consts[131]))
		if v25 != 0 {
			v77 = v3
			m.G0 = v10 + int32(16)
			return v77
		} else {
			v26 = F_objectGetVal(m, l1)
			mBase = m.M
			v27 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v27
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v33 = F_kvstoreHashtableFind(m, v30, v23, v26, v10+int32(12))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				if v35 != 0 {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
					if v40&int32(1) == int32(0) {
						v51 = int64(-1)
					} else {
						v50 = *(*int64)(unsafe.Add(mBase, uint32(v35+(v40&int32(4)^int32(12)))))
						v51 = v50
					}
					v52 = v51
				} else {
					v52 = int64(-1)
				}
				if int64(0) <= v52 {
					v56 = F_commandTimeSnapshot(m)
					mBase = m.M
					v58 = base.B2i32(v52 < v56)
				} else {
					v58 = int32(0)
				}
				if v58 == int32(0) {
					v77 = v27
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, _consts[130]))
					if v62 != 0 {
						v77 = int32(1)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, _consts[132]))
						if v64 == int32(0) {
							v77 = int32(1)
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, _consts[67]))
							if v68 == int32(0) {
								v77 = int32(1)
							} else {
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+207)))
								if v71&int32(32) != 0 {
									v77 = v27
								} else {
									v77 = int32(1)
								}
							}
						}
					}
				}
				m.G0 = v10 + int32(16)
				return v77
			}
		}
	} else {
		v19 = F_getKeySlot(m, v12)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = v19
			v25 = *(*int32)(unsafe.Add(mBase, _consts[131]))
			if v25 != 0 {
				v77 = v3
				m.G0 = v10 + int32(16)
				return v77
			} else {
				v26 = F_objectGetVal(m, l1)
				mBase = m.M
				v27 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v27
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v33 = F_kvstoreHashtableFind(m, v30, v23, v26, v10+int32(12))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v35 != 0 {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
						if v40&int32(1) == int32(0) {
							v51 = int64(-1)
						} else {
							v50 = *(*int64)(unsafe.Add(mBase, uint32(v35+(v40&int32(4)^int32(12)))))
							v51 = v50
						}
						v52 = v51
					} else {
						v52 = int64(-1)
					}
					if int64(0) <= v52 {
						v56 = F_commandTimeSnapshot(m)
						mBase = m.M
						v58 = base.B2i32(v52 < v56)
					} else {
						v58 = int32(0)
					}
					if v58 == int32(0) {
						v77 = v27
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, _consts[130]))
						if v62 != 0 {
							v77 = int32(1)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, _consts[132]))
							if v64 == int32(0) {
								v77 = int32(1)
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, _consts[67]))
								if v68 == int32(0) {
									v77 = int32(1)
								} else {
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+207)))
									if v71&int32(32) != 0 {
										v77 = v27
									} else {
										v77 = int32(1)
									}
								}
							}
						}
					}
					m.G0 = v10 + int32(16)
					return v77
				}
			}
		}
	}
}
func F_lookupKey(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v127 float64
	_ = v127
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v160 int64
	_ = v160
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_objectGetVal(m, l1)
	mBase = m.M
	v14 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v14 == v4 {
		v21 = v4
		v22 = F_objectGetVal(m, l1)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v28 = F_kvstoreHashtableFind(m, v25, v21, v22, v9+int32(12))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			if v30 == int32(0) {
				if l2&int32(10) != 0 {
					v171 = int32(0)
					if l2&int32(12) != 0 {
						v182 = v171
					} else {
						v174 = int32(_a44)
						v176 = *(*int64)(unsafe.Add(mBase, _consts[329]))
						*(*int64)(unsafe.Add(mBase, _consts[329])) = v176 + int64(1)
						v182 = v171
					}
					m.G0 = v9 + int32(16)
					return v182
				} else {
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					F_notifyKeyspaceEvent(m, int32(2048), int32(_a548), l1, v168)
					mBase = m.M
					v170 = m.ExcPending
					if v170 != 0 {
						return int32(0)
					} else {
						v171 = int32(0)
						if l2&int32(12) != 0 {
							v182 = v171
						} else {
							v174 = int32(_a44)
							v176 = *(*int64)(unsafe.Add(mBase, _consts[329]))
							*(*int64)(unsafe.Add(mBase, _consts[329])) = v176 + int64(1)
							v182 = v171
						}
						m.G0 = v9 + int32(16)
						return v182
					}
				}
			} else {
				v35 = int32(_a44)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[130]))
				v37 = int32(0)
				v40 = *(*int32)(unsafe.Add(mBase, _consts[330]))
				v47 = F_expireIfNeededWithDictIndex(m, l0, l1, v30, int32(base.Ui32(l2)>>(uint(int32(3))%32))&(base.B2i32(v36 == v37)|base.B2i32(v40 == v37)|int32(2)), v21)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					if v47 != 0 {
						if l2&int32(10) != 0 {
							v171 = int32(0)
							if l2&int32(12) != 0 {
								v182 = v171
							} else {
								v174 = int32(_a44)
								v176 = *(*int64)(unsafe.Add(mBase, _consts[329]))
								*(*int64)(unsafe.Add(mBase, _consts[329])) = v176 + int64(1)
								v182 = v171
							}
							m.G0 = v9 + int32(16)
							return v182
						} else {
							v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							F_notifyKeyspaceEvent(m, int32(2048), int32(_a548), l1, v168)
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return int32(0)
							} else {
								v171 = int32(0)
								if l2&int32(12) != 0 {
									v182 = v171
								} else {
									v174 = int32(_a44)
									v176 = *(*int64)(unsafe.Add(mBase, _consts[329]))
									*(*int64)(unsafe.Add(mBase, _consts[329])) = v176 + int64(1)
									v182 = v171
								}
								m.G0 = v9 + int32(16)
								return v182
							}
						}
					} else {
						if l2&int32(1) != 0 {
							v70 = l2
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, _consts[67]))
							if v52 == int32(0) {
								v70 = l2
							} else {
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+206)))
								if v55&int32(1) == int32(0) {
									v70 = l2
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, _consts[68]))
									if v61 == int32(0) {
										v70 = l2
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+68))
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
										v70 = l2 | base.B2i32(v65 != int32(211))
									}
								}
							}
						}
						v72 = *(*int32)(unsafe.Add(mBase, _consts[45]))
						if base.B2i32(v72 != int32(-1))|v70&int32(1) != 0 {
							if v70&int32(12) != 0 {
								v182 = v30
							} else {
								v158 = int32(_a44)
								v160 = *(*int64)(unsafe.Add(mBase, _consts[331]))
								*(*int64)(unsafe.Add(mBase, _consts[331])) = v160 + int64(1)
								v182 = v30
							}
							m.G0 = v9 + int32(16)
							return v182
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
							if base.Ui32(int32(-8)) <= base.Ui32(v78) {
								F__serverAssert(m, int32(_a549), int32(_a550), int32(113))
								mBase = m.M
								v191 = m.ExcPending
								if v191 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
								v83 = int32(base.Ui32(v81) >> (uint(int32(8)) % 32))
								v84 = int32(0)
								v89 = int32(*(*uint8)(unsafe.Add(mBase, _consts[332])))
								if v89 != int32(1) {
									v142 = *(*int32)(unsafe.Add(mBase, _consts[333]))
									v149 = v142 & int32(16777215)
								} else {
									v92 = int32(0)
									v93 = int32(*(*uint16)(unsafe.Add(mBase, _consts[334])))
									v95 = *(*int32)(unsafe.Add(mBase, _consts[335]))
									if v95 == v92 {
										v106 = v84
									} else {
										v101 = int32(65535)
										v103 = base.I32_div_s((v93-int32(base.Ui32(v83)>>(uint(int32(8))%32)))&v101, v95)
										v106 = v103 & v101
									}
									v107 = int32(255)
									v110 = v83 & v107
									v111 = v110 - v106
									if base.Ui32(v110) < base.Ui32(v111) {
										v113 = int32(0)
									} else {
										v113 = v111
									}
									if v113 == int32(255) {
										v137 = v107
									} else {
										v116 = F_rand(m)
										mBase = m.M
										if base.Ui32(v113) < base.Ui32(int32(5)) {
											v127 = float64(0)
										} else {
											v127 = base.F64_convert_i32_s(v113 + int32(-5))
										}
										v129 = *(*int32)(unsafe.Add(mBase, _consts[336]))
										v137 = v113 + base.F64_lt(base.F64_div(base.F64_convert_i32_s(v116), float64(2.147483647e+09)), base.F64_div(float64(1), base.F64_add(base.F64_mul(v127, base.F64_convert_i32_s(v129)), float64(1))))
									}
									v149 = v137 + v93<<(uint(int32(8))%32)
								}
								v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
								*(*int32)(unsafe.Add(mBase, uint32(v30))) = v150 | v149<<(uint(int32(8))%32)
								if v70&int32(12) != 0 {
									v182 = v30
								} else {
									v158 = int32(_a44)
									v160 = *(*int64)(unsafe.Add(mBase, _consts[331]))
									*(*int64)(unsafe.Add(mBase, _consts[331])) = v160 + int64(1)
									v182 = v30
								}
								m.G0 = v9 + int32(16)
								return v182
							}
						}
					}
				}
			}
		}
	} else {
		v17 = F_getKeySlot(m, v11)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = v17
			v22 = F_objectGetVal(m, l1)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v28 = F_kvstoreHashtableFind(m, v25, v21, v22, v9+int32(12))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				if v30 == int32(0) {
					if l2&int32(10) != 0 {
						v171 = int32(0)
						if l2&int32(12) != 0 {
							v182 = v171
						} else {
							v174 = int32(_a44)
							v176 = *(*int64)(unsafe.Add(mBase, _consts[329]))
							*(*int64)(unsafe.Add(mBase, _consts[329])) = v176 + int64(1)
							v182 = v171
						}
						m.G0 = v9 + int32(16)
						return v182
					} else {
						v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						F_notifyKeyspaceEvent(m, int32(2048), int32(_a548), l1, v168)
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return int32(0)
						} else {
							v171 = int32(0)
							if l2&int32(12) != 0 {
								v182 = v171
							} else {
								v174 = int32(_a44)
								v176 = *(*int64)(unsafe.Add(mBase, _consts[329]))
								*(*int64)(unsafe.Add(mBase, _consts[329])) = v176 + int64(1)
								v182 = v171
							}
							m.G0 = v9 + int32(16)
							return v182
						}
					}
				} else {
					v35 = int32(_a44)
					v36 = *(*int32)(unsafe.Add(mBase, _consts[130]))
					v37 = int32(0)
					v40 = *(*int32)(unsafe.Add(mBase, _consts[330]))
					v47 = F_expireIfNeededWithDictIndex(m, l0, l1, v30, int32(base.Ui32(l2)>>(uint(int32(3))%32))&(base.B2i32(v36 == v37)|base.B2i32(v40 == v37)|int32(2)), v21)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						if v47 != 0 {
							if l2&int32(10) != 0 {
								v171 = int32(0)
								if l2&int32(12) != 0 {
									v182 = v171
								} else {
									v174 = int32(_a44)
									v176 = *(*int64)(unsafe.Add(mBase, _consts[329]))
									*(*int64)(unsafe.Add(mBase, _consts[329])) = v176 + int64(1)
									v182 = v171
								}
								m.G0 = v9 + int32(16)
								return v182
							} else {
								v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								F_notifyKeyspaceEvent(m, int32(2048), int32(_a548), l1, v168)
								mBase = m.M
								v170 = m.ExcPending
								if v170 != 0 {
									return int32(0)
								} else {
									v171 = int32(0)
									if l2&int32(12) != 0 {
										v182 = v171
									} else {
										v174 = int32(_a44)
										v176 = *(*int64)(unsafe.Add(mBase, _consts[329]))
										*(*int64)(unsafe.Add(mBase, _consts[329])) = v176 + int64(1)
										v182 = v171
									}
									m.G0 = v9 + int32(16)
									return v182
								}
							}
						} else {
							if l2&int32(1) != 0 {
								v70 = l2
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, _consts[67]))
								if v52 == int32(0) {
									v70 = l2
								} else {
									v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+206)))
									if v55&int32(1) == int32(0) {
										v70 = l2
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, _consts[68]))
										if v61 == int32(0) {
											v70 = l2
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+68))
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
											v70 = l2 | base.B2i32(v65 != int32(211))
										}
									}
								}
							}
							v72 = *(*int32)(unsafe.Add(mBase, _consts[45]))
							if base.B2i32(v72 != int32(-1))|v70&int32(1) != 0 {
								if v70&int32(12) != 0 {
									v182 = v30
								} else {
									v158 = int32(_a44)
									v160 = *(*int64)(unsafe.Add(mBase, _consts[331]))
									*(*int64)(unsafe.Add(mBase, _consts[331])) = v160 + int64(1)
									v182 = v30
								}
								m.G0 = v9 + int32(16)
								return v182
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
								if base.Ui32(int32(-8)) <= base.Ui32(v78) {
									F__serverAssert(m, int32(_a549), int32(_a550), int32(113))
									mBase = m.M
									v191 = m.ExcPending
									if v191 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
									v83 = int32(base.Ui32(v81) >> (uint(int32(8)) % 32))
									v84 = int32(0)
									v89 = int32(*(*uint8)(unsafe.Add(mBase, _consts[332])))
									if v89 != int32(1) {
										v142 = *(*int32)(unsafe.Add(mBase, _consts[333]))
										v149 = v142 & int32(16777215)
									} else {
										v92 = int32(0)
										v93 = int32(*(*uint16)(unsafe.Add(mBase, _consts[334])))
										v95 = *(*int32)(unsafe.Add(mBase, _consts[335]))
										if v95 == v92 {
											v106 = v84
										} else {
											v101 = int32(65535)
											v103 = base.I32_div_s((v93-int32(base.Ui32(v83)>>(uint(int32(8))%32)))&v101, v95)
											v106 = v103 & v101
										}
										v107 = int32(255)
										v110 = v83 & v107
										v111 = v110 - v106
										if base.Ui32(v110) < base.Ui32(v111) {
											v113 = int32(0)
										} else {
											v113 = v111
										}
										if v113 == int32(255) {
											v137 = v107
										} else {
											v116 = F_rand(m)
											mBase = m.M
											if base.Ui32(v113) < base.Ui32(int32(5)) {
												v127 = float64(0)
											} else {
												v127 = base.F64_convert_i32_s(v113 + int32(-5))
											}
											v129 = *(*int32)(unsafe.Add(mBase, _consts[336]))
											v137 = v113 + base.F64_lt(base.F64_div(base.F64_convert_i32_s(v116), float64(2.147483647e+09)), base.F64_div(float64(1), base.F64_add(base.F64_mul(v127, base.F64_convert_i32_s(v129)), float64(1))))
										}
										v149 = v137 + v93<<(uint(int32(8))%32)
									}
									v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
									*(*int32)(unsafe.Add(mBase, uint32(v30))) = v150 | v149<<(uint(int32(8))%32)
									if v70&int32(12) != 0 {
										v182 = v30
									} else {
										v158 = int32(_a44)
										v160 = *(*int64)(unsafe.Add(mBase, _consts[331]))
										*(*int64)(unsafe.Add(mBase, _consts[331])) = v160 + int64(1)
										v182 = v30
									}
									m.G0 = v9 + int32(16)
									return v182
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_lookupKeyReadWithFlags(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	if l2&int32(8) == int32(0) {
		v16 = F_lookupKey(m, l0, l1, l2)
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			return v16
		}
	} else {
		F__serverAssert(m, int32(_a551), int32(_a550), int32(138))
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_lookupKeyWriteOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v6 = F_lookupKey(m, v4, l1, int32(8))
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
