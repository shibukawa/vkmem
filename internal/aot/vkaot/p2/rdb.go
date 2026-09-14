package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_rdbCheckError(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	v5 = m.G0
	v7 = v5 - int32(1104)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+76)) = l1
	v13 = F_vsnprintf(m, v7+int32(80), int32(1024), l0, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v16 = F_puts(m, int32(_a1836))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[977]))
			if v19 != 0 {
				v21 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v19)+40)))
				v22 = v21
			} else {
				v22 = int64(0)
			}
			*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v7 + int32(80)
			v30 = F_iprintf(m, int32(_a1837), v7+int32(48))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _consts[978]))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v33<<(uint(int32(2))%32))+uint32(_consts[979])))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v38
				v43 = F_iprintf(m, int32(_a1838), v7+int32(32))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v45 = int32(0)
					v46 = *(*int32)(unsafe.Add(mBase, _consts[980]))
					if v46 == v45 {
						v57 = *(*int32)(unsafe.Add(mBase, _consts[981]))
						if v57 == int32(-1) {
							F_rdbShowGenericInfo(m)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								m.G0 = v7 + int32(1104)
								return
							}
						} else {
							if base.Ui32(int32(22)) < base.Ui32(v57) {
								v68 = int32(_a288)
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(int32(2))%32))+uint32(_consts[982])))
								v68 = v67
							}
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v57
							v72 = F_iprintf(m, int32(_a1839), v7)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								F_rdbShowGenericInfo(m)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									m.G0 = v7 + int32(1104)
									return
								}
							}
						}
					} else {
						v49 = F_objectGetVal(m, v46)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v49
						v54 = F_iprintf(m, int32(_a1840), v7+int32(16))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, _consts[981]))
							if v57 == int32(-1) {
								F_rdbShowGenericInfo(m)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									m.G0 = v7 + int32(1104)
									return
								}
							} else {
								if base.Ui32(int32(22)) < base.Ui32(v57) {
									v68 = int32(_a288)
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(int32(2))%32))+uint32(_consts[982])))
									v68 = v67
								}
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v68
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v57
								v72 = F_iprintf(m, int32(_a1839), v7)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									F_rdbShowGenericInfo(m)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										m.G0 = v7 + int32(1104)
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
func F_rdbCheckInfo(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(1056)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = l1
	v12 = v8 + int32(32)
	v16 = F_vsnprintf(m, v12, int32(1024), l0, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 < int32(1024) {
			v54 = v12
			v56 = *(*int32)(unsafe.Add(mBase, _consts[977]))
			if v56 != 0 {
				v58 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v56)+40)))
				v59 = v58
			} else {
				v59 = int64(0)
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = v59
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v54
			v63 = F_iprintf(m, int32(_a1837), v8)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				if v54 == v8+int32(32) {
					m.G0 = v8 + int32(1056)
					return
				} else {
					F_sdsfree(m, v54)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						m.G0 = v8 + int32(1056)
						return
					}
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[72]))
			v23 = F_sdsnewlen(m, v21, int32(2048))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v27 = v23 + int32(-1)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
				switch v28 & int32(7) {
				case 0:
					v31 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v31)
				case 1:
					v35 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-3)))) = uint8(v35)
				case 2:
					v39 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(-5)))) = uint16(v39)
				case 3:
					*(*int32)(unsafe.Add(mBase, uint32(v23+int32(-9)))) = int32(0)
				case 4:
					*(*int64)(unsafe.Add(mBase, uint32(v23+int32(-17)))) = int64(0)
				default:
				}
				v49 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v49)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = l1
				v52 = F_sdscatvprintf(m, v23, l0, l1)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					v54 = v52
					v56 = *(*int32)(unsafe.Add(mBase, _consts[977]))
					if v56 != 0 {
						v58 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v56)+40)))
						v59 = v58
					} else {
						v59 = int64(0)
					}
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = v59
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v54
					v63 = F_iprintf(m, int32(_a1837), v8)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						if v54 == v8+int32(32) {
							m.G0 = v8 + int32(1056)
							return
						} else {
							F_sdsfree(m, v54)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								m.G0 = v8 + int32(1056)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_rdbGenericLoadStringObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
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
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int64
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int64
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v252 int32
	_ = v252
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(64)
	m.G0 = v13
	v20 = F_rdbLoadLenByRef(m, l0, v13+int32(52), v13+int32(56))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(64)
	return v284
L2:
	;
	return int32(0)
L3:
	;
	if v20 == int32(-1) {
		v284 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v13)+56))
	if v26 == int64(-1) {
		v284 = v4
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	if v29 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if l1&int32(6) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	if base.Ui64(int64(3)) < base.Ui64(v26) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v26
	F_rdbReportError(m, int32(1), int32(586), int32(_a1124), v13+int32(32))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L14
	}
L9:
	;
	v34 = base.I32_wrap_i64(v26)
	switch v34 {
	default:
		goto L11
	case 3:
		goto L10
	}
L10:
	;
	v37 = F_rdbLoadLzfStringObject(m, l0, l1, l2)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L13
	}
L11:
	;
	v35 = F_rdbLoadIntegerObject(m, l0, v34, l1, l2)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v284 = v35
	goto L1
L13:
	;
	v284 = v37
	goto L1
L14:
	;
	v284 = v4
	goto L1
L15:
	;
	if v52 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L16:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v177 = base.I32_wrap_i64(v26)
	v178 = F_tryCreateStringObject(m, v176, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L2
	} else {
		goto L69
	}
L17:
	;
	v52 = l1 & int32(2)
	if v52 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v109 != 0 {
		goto L34
	} else {
		goto L35
	}
L19:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v107 = F_sdstrynewlen(m, v105, base.I32_wrap_i64(v26))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L33
	}
L20:
	;
	v55 = base.I32_wrap_i64(v26)
	v56 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(v55) {
		v101 = v56
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v109 = v101
	goto L18
L22:
	;
	goto L21
L23:
	;
	if v55 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v63 = v55
	goto L26
L25:
	;
	v63 = int32(4)
	goto L26
L26:
	;
	v65 = v63 + int32(8)
	v66 = F_emscripten_builtin_malloc(m, v65)
	mBase = m.M
	if v66 == int32(0) {
		v101 = v56
		goto L22
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v63
	v71 = *(*int32)(unsafe.Add(mBase, _consts[411]))
	if v71 != int32(-1) {
		v82 = v71
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v82 < int32(260) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v74 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v76
	*(*int32)(unsafe.Add(mBase, _consts[281])) = v76 + int32(1)
	v82 = v76
	goto L28
L30:
	;
	v101 = v66 + int32(8)
	goto L22
L31:
	;
	v91 = v82 << (uint(int32(2)) % 32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[285])))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[285]))) = v94 + v65
	goto L30
L32:
	;
	v85 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	*(*int32)(unsafe.Add(mBase, _consts[286])) = v87 + v65
	goto L30
L33:
	;
	v109 = v107
	goto L18
L34:
	;
	if l2 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L35:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	if v111 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v26
	F__serverLog(m, v126, int32(_a1125), v13+int32(16))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L2
	} else {
		goto L44
	}
L37:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v111)))
	if v119 == int64(-1) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v112 = int32(3)
	v114 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v114 <= v112 {
		v126 = v112
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v284 = v4
	goto L1
L40:
	;
	v122 = int32(3)
	goto L42
L41:
	;
	v122 = int32(1)
	goto L42
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v122 < v124 {
		v284 = v4
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v126 = v122
	goto L36
L44:
	;
	v284 = v4
	goto L1
L45:
	;
	if base.B2i32(v26 == int64(0)) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l2))) = uint32(v26)
	goto L45
L47:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v140&int32(5) != 0 {
		goto L15
	} else {
		goto L49
	}
L48:
	;
	v284 = v109
	goto L1
L49:
	;
	v143 = base.I32_wrap_i64(v26)
	if v143 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v145 = v143
	v152 = v109
	goto L52
L51:
	;
	v284 = v109
	goto L1
L52:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v154) < base.Ui32(v145) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v284 = v109
	goto L1
L54:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v165 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L55:
	;
	v156 = v154
	goto L57
L56:
	;
	v156 = v145
	goto L57
L57:
	;
	if v154 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v157 = v156
	goto L60
L59:
	;
	v157 = v145
	goto L60
L60:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v159 = m.T0[v158].(func(*base.Module, int32, int32, int32) int32)(m, l0, v152, v157)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	if v159 != 0 {
		goto L54
	} else {
		goto L62
	}
L62:
	;
	v161 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v161 | int64(1)
	goto L15
L63:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v170 + v157
	v174 = v145 - v157
	if v174 != 0 {
		v145 = v174
		v152 = v152 + v157
		goto L52
	} else {
		goto L66
	}
L64:
	;
	m.T0[v165].(func(*base.Module, int32, int32, int32))(m, l0, v152, v157)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	goto L53
L67:
	;
	v284 = int32(0)
	goto L1
L68:
	;
	if v26 == int64(0) {
		v284 = v178
		goto L1
	} else {
		goto L80
	}
L69:
	;
	if v178 != 0 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	if v181 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v26
	F__serverLog(m, v196, int32(_a1125), v13)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L79
	}
L72:
	;
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v181)))
	if v189 == int64(-1) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v182 = int32(3)
	v184 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v184 <= v182 {
		v196 = v182
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L67
L75:
	;
	v192 = int32(3)
	goto L77
L76:
	;
	v192 = int32(1)
	goto L77
L77:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v192 < v194 {
		goto L67
	} else {
		goto L78
	}
L78:
	;
	v196 = v192
	goto L71
L79:
	;
	v284 = int32(0)
	goto L1
L80:
	;
	v204 = F_objectGetVal(m, v178)
	mBase = m.M
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v205&int32(5) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	F_decrRefCount(m, v178)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L2
	} else {
		goto L99
	}
L82:
	;
	if v177 == int32(0) {
		v284 = v178
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v211 = v177
	v218 = v204
	goto L84
L84:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v220) < base.Ui32(v211) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v231 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L87:
	;
	v222 = v220
	goto L89
L88:
	;
	v222 = v211
	goto L89
L89:
	;
	if v220 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v223 = v222
	goto L92
L91:
	;
	v223 = v211
	goto L92
L92:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v225 = m.T0[v224].(func(*base.Module, int32, int32, int32) int32)(m, l0, v218, v223)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	if v225 != 0 {
		goto L86
	} else {
		goto L94
	}
L94:
	;
	v227 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v227 | int64(1)
	goto L81
L95:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v236 + v223
	v240 = v211 - v223
	if v240 != 0 {
		v211 = v240
		v218 = v218 + v223
		goto L84
	} else {
		goto L98
	}
L96:
	;
	m.T0[v231].(func(*base.Module, int32, int32, int32))(m, l0, v218, v223)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v284 = v178
	goto L1
L99:
	;
	goto L67
L100:
	;
	F_sdsfree(m, v109)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L2
	} else {
		goto L103
	}
L101:
	;
	F_valkey_free(m, v109)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	v284 = v4
	goto L1
L103:
	;
	v284 = v4
	goto L1
}
func F_rdbLoadProgressCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int64
	_ = v261
	v7 = *(*int32)(unsafe.Add(mBase, _consts[634]))
	if v7 == int32(0) {
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v10&int32(8) != 0 {
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			v15 = F_crc64(m, v13, l1, base.I64_extend_i32_u(l2))
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v15
		}
	}
	v18 = *(*int64)(unsafe.Add(mBase, _consts[641]))
	if v18 == int64(0) {
		v241 = *(*int32)(unsafe.Add(mBase, _consts[193]))
		if v241 != int32(13) {
		} else {
			v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v244 != int32(981) {
				if v244 != int32(982) {
					if v244 == int32(983) {
						v255 = int32(4)
					} else {
						v255 = int32(8)
					}
					v256 = v255
				} else {
					v256 = int32(2)
				}
			} else {
				v256 = int32(1)
			}
			if v256 != int32(4) {
			} else {
				v259 = int32(_a44)
				v261 = *(*int64)(unsafe.Add(mBase, _consts[303]))
				*(*int64)(unsafe.Add(mBase, _consts[303])) = v261 + base.I64_extend_i32_u(l2)
			}
		}
		return
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v24 = base.I64_div_s(base.I64_extend_i32_u(v21+l2), v18)
		v25 = base.I64_extend_i32_u(v21)
		v26 = base.I64_div_s(v25, v18)
		if v24 <= v26 {
			v241 = *(*int32)(unsafe.Add(mBase, _consts[193]))
			if v241 != int32(13) {
			} else {
				v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v244 != int32(981) {
					if v244 != int32(982) {
						if v244 == int32(983) {
							v255 = int32(4)
						} else {
							v255 = int32(8)
						}
						v256 = v255
					} else {
						v256 = int32(2)
					}
				} else {
					v256 = int32(1)
				}
				if v256 != int32(4) {
				} else {
					v259 = int32(_a44)
					v261 = *(*int64)(unsafe.Add(mBase, _consts[303]))
					*(*int64)(unsafe.Add(mBase, _consts[303])) = v261 + base.I64_extend_i32_u(l2)
				}
			}
			return
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, _consts[130]))
			if v29 == int32(0) {
				v39 = v25
				v40 = int32(_a44)
				*(*int64)(unsafe.Add(mBase, _consts[70])) = v39
				v43 = *(*int32)(unsafe.Add(mBase, _consts[71]))
				v44 = int32(0)
				v53 = *(*int32)(unsafe.Add(mBase, _consts[281]))
				if v53 < int32(261) {
					if v53 < int32(1) {
						v130 = v44
					} else {
						v61 = v44
						v62 = v53
						v64 = v62 & int32(3)
						if base.Ui32(int32(4)) <= base.Ui32(v62) {
							v71 = int32(0)
							v73 = v61
							v74 = v71
							v78 = v71
							for {
								v81 = v74 << (uint(int32(2)) % 32)
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[282])))
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[283])))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[284])))
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[285])))
								v97 = v84 + (v87 + (v90 + (v93 + v73)))
								v98 = int32(4)
								v99 = v74 + v98
								v101 = v78 + v98
								if v101 != v62&int32(2147483644) {
									v73 = v97
									v74 = v99
									v78 = v101
									continue
								} else {
									break
								}
								break
							}
							v103 = v97
							v104 = v99
						} else {
							v103 = v61
							v104 = int32(0)
						}
						if v64 == int32(0) {
							v130 = v103
						} else {
							v112 = v103
							v113 = v104
							v115 = int32(0)
							for {
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v113<<(uint(int32(2))%32))+uint32(_consts[285])))
								v124 = v123 + v112
								v125 = int32(1)
								v128 = v115 + v125
								if v128 != v64 {
									v112 = v124
									v113 = v113 + v125
									v115 = v128
									continue
								} else {
									break
								}
								break
							}
							v130 = v124
						}
					}
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, _consts[286]))
					v61 = v57
					v62 = int32(260)
					v64 = v62 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(v62) {
						v71 = int32(0)
						v73 = v61
						v74 = v71
						v78 = v71
						for {
							v81 = v74 << (uint(int32(2)) % 32)
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[282])))
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[283])))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[284])))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[285])))
							v97 = v84 + (v87 + (v90 + (v93 + v73)))
							v98 = int32(4)
							v99 = v74 + v98
							v101 = v78 + v98
							if v101 != v62&int32(2147483644) {
								v73 = v97
								v74 = v99
								v78 = v101
								continue
							} else {
								break
							}
							break
						}
						v103 = v97
						v104 = v99
					} else {
						v103 = v61
						v104 = int32(0)
					}
					if v64 == int32(0) {
						v130 = v103
					} else {
						v112 = v103
						v113 = v104
						v115 = int32(0)
						for {
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v113<<(uint(int32(2))%32))+uint32(_consts[285])))
							v124 = v123 + v112
							v125 = int32(1)
							v128 = v115 + v125
							if v128 != v64 {
								v112 = v124
								v113 = v113 + v125
								v115 = v128
								continue
							} else {
								break
							}
							break
						}
						v130 = v124
					}
				}
				if base.Ui32(v130) <= base.Ui32(v43) {
				} else {
					v139 = int32(0)
					v148 = *(*int32)(unsafe.Add(mBase, _consts[281]))
					if v148 < int32(261) {
						if v148 < int32(1) {
							v225 = v139
						} else {
							v156 = v139
							v157 = v148
							v159 = v157 & int32(3)
							if base.Ui32(int32(4)) <= base.Ui32(v157) {
								v166 = int32(0)
								v168 = v156
								v169 = v166
								v173 = v166
								for {
									v176 = v169 << (uint(int32(2)) % 32)
									v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[282])))
									v182 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[283])))
									v185 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[284])))
									v188 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[285])))
									v192 = v179 + (v182 + (v185 + (v188 + v168)))
									v193 = int32(4)
									v194 = v169 + v193
									v196 = v173 + v193
									if v196 != v157&int32(2147483644) {
										v168 = v192
										v169 = v194
										v173 = v196
										continue
									} else {
										break
									}
									break
								}
								v198 = v192
								v199 = v194
							} else {
								v198 = v156
								v199 = int32(0)
							}
							if v159 == int32(0) {
								v225 = v198
							} else {
								v207 = v198
								v208 = v199
								v210 = int32(0)
								for {
									v218 = *(*int32)(unsafe.Add(mBase, uint32(v208<<(uint(int32(2))%32))+uint32(_consts[285])))
									v219 = v218 + v207
									v220 = int32(1)
									v223 = v210 + v220
									if v223 != v159 {
										v207 = v219
										v208 = v208 + v220
										v210 = v223
										continue
									} else {
										break
									}
									break
								}
								v225 = v219
							}
						}
					} else {
						v152 = *(*int32)(unsafe.Add(mBase, _consts[286]))
						v156 = v152
						v157 = int32(260)
						v159 = v157 & int32(3)
						if base.Ui32(int32(4)) <= base.Ui32(v157) {
							v166 = int32(0)
							v168 = v156
							v169 = v166
							v173 = v166
							for {
								v176 = v169 << (uint(int32(2)) % 32)
								v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[282])))
								v182 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[283])))
								v185 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[284])))
								v188 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[285])))
								v192 = v179 + (v182 + (v185 + (v188 + v168)))
								v193 = int32(4)
								v194 = v169 + v193
								v196 = v173 + v193
								if v196 != v157&int32(2147483644) {
									v168 = v192
									v169 = v194
									v173 = v196
									continue
								} else {
									break
								}
								break
							}
							v198 = v192
							v199 = v194
						} else {
							v198 = v156
							v199 = int32(0)
						}
						if v159 == int32(0) {
							v225 = v198
						} else {
							v207 = v198
							v208 = v199
							v210 = int32(0)
							for {
								v218 = *(*int32)(unsafe.Add(mBase, uint32(v208<<(uint(int32(2))%32))+uint32(_consts[285])))
								v219 = v218 + v207
								v220 = int32(1)
								v223 = v210 + v220
								if v223 != v159 {
									v207 = v219
									v208 = v208 + v220
									v210 = v223
									continue
								} else {
									break
								}
								break
							}
							v225 = v219
						}
					}
					*(*int32)(unsafe.Add(mBase, _consts[71])) = v225
				}
				F_processEventsWhileBlocked(m)
				mBase = m.M
				v234 = m.ExcPending
				if v234 != 0 {
					return
				} else {
					F_processModuleLoadingProgressEvent(m, int32(0))
					mBase = m.M
					v237 = m.ExcPending
					if v237 != 0 {
						return
					} else {
						v241 = *(*int32)(unsafe.Add(mBase, _consts[193]))
						if v241 != int32(13) {
						} else {
							v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							if v244 != int32(981) {
								if v244 != int32(982) {
									if v244 == int32(983) {
										v255 = int32(4)
									} else {
										v255 = int32(8)
									}
									v256 = v255
								} else {
									v256 = int32(2)
								}
							} else {
								v256 = int32(1)
							}
							if v256 != int32(4) {
							} else {
								v259 = int32(_a44)
								v261 = *(*int64)(unsafe.Add(mBase, _consts[303]))
								*(*int64)(unsafe.Add(mBase, _consts[303])) = v261 + base.I64_extend_i32_u(l2)
							}
						}
						return
					}
				}
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _consts[193]))
				if v33 != int32(13) {
					v39 = v25
					v40 = int32(_a44)
					*(*int64)(unsafe.Add(mBase, _consts[70])) = v39
					v43 = *(*int32)(unsafe.Add(mBase, _consts[71]))
					v44 = int32(0)
					v53 = *(*int32)(unsafe.Add(mBase, _consts[281]))
					if v53 < int32(261) {
						if v53 < int32(1) {
							v130 = v44
						} else {
							v61 = v44
							v62 = v53
							v64 = v62 & int32(3)
							if base.Ui32(int32(4)) <= base.Ui32(v62) {
								v71 = int32(0)
								v73 = v61
								v74 = v71
								v78 = v71
								for {
									v81 = v74 << (uint(int32(2)) % 32)
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[282])))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[283])))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[284])))
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[285])))
									v97 = v84 + (v87 + (v90 + (v93 + v73)))
									v98 = int32(4)
									v99 = v74 + v98
									v101 = v78 + v98
									if v101 != v62&int32(2147483644) {
										v73 = v97
										v74 = v99
										v78 = v101
										continue
									} else {
										break
									}
									break
								}
								v103 = v97
								v104 = v99
							} else {
								v103 = v61
								v104 = int32(0)
							}
							if v64 == int32(0) {
								v130 = v103
							} else {
								v112 = v103
								v113 = v104
								v115 = int32(0)
								for {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(v113<<(uint(int32(2))%32))+uint32(_consts[285])))
									v124 = v123 + v112
									v125 = int32(1)
									v128 = v115 + v125
									if v128 != v64 {
										v112 = v124
										v113 = v113 + v125
										v115 = v128
										continue
									} else {
										break
									}
									break
								}
								v130 = v124
							}
						}
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, _consts[286]))
						v61 = v57
						v62 = int32(260)
						v64 = v62 & int32(3)
						if base.Ui32(int32(4)) <= base.Ui32(v62) {
							v71 = int32(0)
							v73 = v61
							v74 = v71
							v78 = v71
							for {
								v81 = v74 << (uint(int32(2)) % 32)
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[282])))
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[283])))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[284])))
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[285])))
								v97 = v84 + (v87 + (v90 + (v93 + v73)))
								v98 = int32(4)
								v99 = v74 + v98
								v101 = v78 + v98
								if v101 != v62&int32(2147483644) {
									v73 = v97
									v74 = v99
									v78 = v101
									continue
								} else {
									break
								}
								break
							}
							v103 = v97
							v104 = v99
						} else {
							v103 = v61
							v104 = int32(0)
						}
						if v64 == int32(0) {
							v130 = v103
						} else {
							v112 = v103
							v113 = v104
							v115 = int32(0)
							for {
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v113<<(uint(int32(2))%32))+uint32(_consts[285])))
								v124 = v123 + v112
								v125 = int32(1)
								v128 = v115 + v125
								if v128 != v64 {
									v112 = v124
									v113 = v113 + v125
									v115 = v128
									continue
								} else {
									break
								}
								break
							}
							v130 = v124
						}
					}
					if base.Ui32(v130) <= base.Ui32(v43) {
					} else {
						v139 = int32(0)
						v148 = *(*int32)(unsafe.Add(mBase, _consts[281]))
						if v148 < int32(261) {
							if v148 < int32(1) {
								v225 = v139
							} else {
								v156 = v139
								v157 = v148
								v159 = v157 & int32(3)
								if base.Ui32(int32(4)) <= base.Ui32(v157) {
									v166 = int32(0)
									v168 = v156
									v169 = v166
									v173 = v166
									for {
										v176 = v169 << (uint(int32(2)) % 32)
										v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[282])))
										v182 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[283])))
										v185 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[284])))
										v188 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[285])))
										v192 = v179 + (v182 + (v185 + (v188 + v168)))
										v193 = int32(4)
										v194 = v169 + v193
										v196 = v173 + v193
										if v196 != v157&int32(2147483644) {
											v168 = v192
											v169 = v194
											v173 = v196
											continue
										} else {
											break
										}
										break
									}
									v198 = v192
									v199 = v194
								} else {
									v198 = v156
									v199 = int32(0)
								}
								if v159 == int32(0) {
									v225 = v198
								} else {
									v207 = v198
									v208 = v199
									v210 = int32(0)
									for {
										v218 = *(*int32)(unsafe.Add(mBase, uint32(v208<<(uint(int32(2))%32))+uint32(_consts[285])))
										v219 = v218 + v207
										v220 = int32(1)
										v223 = v210 + v220
										if v223 != v159 {
											v207 = v219
											v208 = v208 + v220
											v210 = v223
											continue
										} else {
											break
										}
										break
									}
									v225 = v219
								}
							}
						} else {
							v152 = *(*int32)(unsafe.Add(mBase, _consts[286]))
							v156 = v152
							v157 = int32(260)
							v159 = v157 & int32(3)
							if base.Ui32(int32(4)) <= base.Ui32(v157) {
								v166 = int32(0)
								v168 = v156
								v169 = v166
								v173 = v166
								for {
									v176 = v169 << (uint(int32(2)) % 32)
									v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[282])))
									v182 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[283])))
									v185 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[284])))
									v188 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[285])))
									v192 = v179 + (v182 + (v185 + (v188 + v168)))
									v193 = int32(4)
									v194 = v169 + v193
									v196 = v173 + v193
									if v196 != v157&int32(2147483644) {
										v168 = v192
										v169 = v194
										v173 = v196
										continue
									} else {
										break
									}
									break
								}
								v198 = v192
								v199 = v194
							} else {
								v198 = v156
								v199 = int32(0)
							}
							if v159 == int32(0) {
								v225 = v198
							} else {
								v207 = v198
								v208 = v199
								v210 = int32(0)
								for {
									v218 = *(*int32)(unsafe.Add(mBase, uint32(v208<<(uint(int32(2))%32))+uint32(_consts[285])))
									v219 = v218 + v207
									v220 = int32(1)
									v223 = v210 + v220
									if v223 != v159 {
										v207 = v219
										v208 = v208 + v220
										v210 = v223
										continue
									} else {
										break
									}
									break
								}
								v225 = v219
							}
						}
						*(*int32)(unsafe.Add(mBase, _consts[71])) = v225
					}
					F_processEventsWhileBlocked(m)
					mBase = m.M
					v234 = m.ExcPending
					if v234 != 0 {
						return
					} else {
						F_processModuleLoadingProgressEvent(m, int32(0))
						mBase = m.M
						v237 = m.ExcPending
						if v237 != 0 {
							return
						} else {
							v241 = *(*int32)(unsafe.Add(mBase, _consts[193]))
							if v241 != int32(13) {
							} else {
								v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if v244 != int32(981) {
									if v244 != int32(982) {
										if v244 == int32(983) {
											v255 = int32(4)
										} else {
											v255 = int32(8)
										}
										v256 = v255
									} else {
										v256 = int32(2)
									}
								} else {
									v256 = int32(1)
								}
								if v256 != int32(4) {
								} else {
									v259 = int32(_a44)
									v261 = *(*int64)(unsafe.Add(mBase, _consts[303]))
									*(*int64)(unsafe.Add(mBase, _consts[303])) = v261 + base.I64_extend_i32_u(l2)
								}
							}
							return
						}
					}
				} else {
					F_replicationSendNewlineToPrimary(m)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+40)))
						v39 = v38
						v40 = int32(_a44)
						*(*int64)(unsafe.Add(mBase, _consts[70])) = v39
						v43 = *(*int32)(unsafe.Add(mBase, _consts[71]))
						v44 = int32(0)
						v53 = *(*int32)(unsafe.Add(mBase, _consts[281]))
						if v53 < int32(261) {
							if v53 < int32(1) {
								v130 = v44
							} else {
								v61 = v44
								v62 = v53
								v64 = v62 & int32(3)
								if base.Ui32(int32(4)) <= base.Ui32(v62) {
									v71 = int32(0)
									v73 = v61
									v74 = v71
									v78 = v71
									for {
										v81 = v74 << (uint(int32(2)) % 32)
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[282])))
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[283])))
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[284])))
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[285])))
										v97 = v84 + (v87 + (v90 + (v93 + v73)))
										v98 = int32(4)
										v99 = v74 + v98
										v101 = v78 + v98
										if v101 != v62&int32(2147483644) {
											v73 = v97
											v74 = v99
											v78 = v101
											continue
										} else {
											break
										}
										break
									}
									v103 = v97
									v104 = v99
								} else {
									v103 = v61
									v104 = int32(0)
								}
								if v64 == int32(0) {
									v130 = v103
								} else {
									v112 = v103
									v113 = v104
									v115 = int32(0)
									for {
										v123 = *(*int32)(unsafe.Add(mBase, uint32(v113<<(uint(int32(2))%32))+uint32(_consts[285])))
										v124 = v123 + v112
										v125 = int32(1)
										v128 = v115 + v125
										if v128 != v64 {
											v112 = v124
											v113 = v113 + v125
											v115 = v128
											continue
										} else {
											break
										}
										break
									}
									v130 = v124
								}
							}
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, _consts[286]))
							v61 = v57
							v62 = int32(260)
							v64 = v62 & int32(3)
							if base.Ui32(int32(4)) <= base.Ui32(v62) {
								v71 = int32(0)
								v73 = v61
								v74 = v71
								v78 = v71
								for {
									v81 = v74 << (uint(int32(2)) % 32)
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[282])))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[283])))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[284])))
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[285])))
									v97 = v84 + (v87 + (v90 + (v93 + v73)))
									v98 = int32(4)
									v99 = v74 + v98
									v101 = v78 + v98
									if v101 != v62&int32(2147483644) {
										v73 = v97
										v74 = v99
										v78 = v101
										continue
									} else {
										break
									}
									break
								}
								v103 = v97
								v104 = v99
							} else {
								v103 = v61
								v104 = int32(0)
							}
							if v64 == int32(0) {
								v130 = v103
							} else {
								v112 = v103
								v113 = v104
								v115 = int32(0)
								for {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(v113<<(uint(int32(2))%32))+uint32(_consts[285])))
									v124 = v123 + v112
									v125 = int32(1)
									v128 = v115 + v125
									if v128 != v64 {
										v112 = v124
										v113 = v113 + v125
										v115 = v128
										continue
									} else {
										break
									}
									break
								}
								v130 = v124
							}
						}
						if base.Ui32(v130) <= base.Ui32(v43) {
						} else {
							v139 = int32(0)
							v148 = *(*int32)(unsafe.Add(mBase, _consts[281]))
							if v148 < int32(261) {
								if v148 < int32(1) {
									v225 = v139
								} else {
									v156 = v139
									v157 = v148
									v159 = v157 & int32(3)
									if base.Ui32(int32(4)) <= base.Ui32(v157) {
										v166 = int32(0)
										v168 = v156
										v169 = v166
										v173 = v166
										for {
											v176 = v169 << (uint(int32(2)) % 32)
											v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[282])))
											v182 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[283])))
											v185 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[284])))
											v188 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[285])))
											v192 = v179 + (v182 + (v185 + (v188 + v168)))
											v193 = int32(4)
											v194 = v169 + v193
											v196 = v173 + v193
											if v196 != v157&int32(2147483644) {
												v168 = v192
												v169 = v194
												v173 = v196
												continue
											} else {
												break
											}
											break
										}
										v198 = v192
										v199 = v194
									} else {
										v198 = v156
										v199 = int32(0)
									}
									if v159 == int32(0) {
										v225 = v198
									} else {
										v207 = v198
										v208 = v199
										v210 = int32(0)
										for {
											v218 = *(*int32)(unsafe.Add(mBase, uint32(v208<<(uint(int32(2))%32))+uint32(_consts[285])))
											v219 = v218 + v207
											v220 = int32(1)
											v223 = v210 + v220
											if v223 != v159 {
												v207 = v219
												v208 = v208 + v220
												v210 = v223
												continue
											} else {
												break
											}
											break
										}
										v225 = v219
									}
								}
							} else {
								v152 = *(*int32)(unsafe.Add(mBase, _consts[286]))
								v156 = v152
								v157 = int32(260)
								v159 = v157 & int32(3)
								if base.Ui32(int32(4)) <= base.Ui32(v157) {
									v166 = int32(0)
									v168 = v156
									v169 = v166
									v173 = v166
									for {
										v176 = v169 << (uint(int32(2)) % 32)
										v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[282])))
										v182 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[283])))
										v185 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[284])))
										v188 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[285])))
										v192 = v179 + (v182 + (v185 + (v188 + v168)))
										v193 = int32(4)
										v194 = v169 + v193
										v196 = v173 + v193
										if v196 != v157&int32(2147483644) {
											v168 = v192
											v169 = v194
											v173 = v196
											continue
										} else {
											break
										}
										break
									}
									v198 = v192
									v199 = v194
								} else {
									v198 = v156
									v199 = int32(0)
								}
								if v159 == int32(0) {
									v225 = v198
								} else {
									v207 = v198
									v208 = v199
									v210 = int32(0)
									for {
										v218 = *(*int32)(unsafe.Add(mBase, uint32(v208<<(uint(int32(2))%32))+uint32(_consts[285])))
										v219 = v218 + v207
										v220 = int32(1)
										v223 = v210 + v220
										if v223 != v159 {
											v207 = v219
											v208 = v208 + v220
											v210 = v223
											continue
										} else {
											break
										}
										break
									}
									v225 = v219
								}
							}
							*(*int32)(unsafe.Add(mBase, _consts[71])) = v225
						}
						F_processEventsWhileBlocked(m)
						mBase = m.M
						v234 = m.ExcPending
						if v234 != 0 {
							return
						} else {
							F_processModuleLoadingProgressEvent(m, int32(0))
							mBase = m.M
							v237 = m.ExcPending
							if v237 != 0 {
								return
							} else {
								v241 = *(*int32)(unsafe.Add(mBase, _consts[193]))
								if v241 != int32(13) {
								} else {
									v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									if v244 != int32(981) {
										if v244 != int32(982) {
											if v244 == int32(983) {
												v255 = int32(4)
											} else {
												v255 = int32(8)
											}
											v256 = v255
										} else {
											v256 = int32(2)
										}
									} else {
										v256 = int32(1)
									}
									if v256 != int32(4) {
									} else {
										v259 = int32(_a44)
										v261 = *(*int64)(unsafe.Add(mBase, _consts[303]))
										*(*int64)(unsafe.Add(mBase, _consts[303])) = v261 + base.I64_extend_i32_u(l2)
									}
								}
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_rdbLoadRioWithLoadingCtx(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
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
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v357 int64
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int64
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v384 int64
	_ = v384
	var v390 int64
	_ = v390
	var v404 int32
	_ = v404
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v412 int64
	_ = v412
	var v413 int32
	_ = v413
	var v414 int64
	_ = v414
	var v415 int64
	_ = v415
	var v416 int64
	_ = v416
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int64
	_ = v463
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v485 int32
	_ = v485
	var v495 int32
	_ = v495
	var v498 int64
	_ = v498
	var v499 int64
	_ = v499
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int64
	_ = v544
	var v546 int64
	_ = v546
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int64
	_ = v558
	var v561 int64
	_ = v561
	var v587 int64
	_ = v587
	var v589 int64
	_ = v589
	var v596 int64
	_ = v596
	var v597 int64
	_ = v597
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int64
	_ = v642
	var v644 int64
	_ = v644
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int64
	_ = v656
	var v657 int64
	_ = v657
	var v683 int64
	_ = v683
	var v685 int64
	_ = v685
	var v692 int32
	_ = v692
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int64
	_ = v735
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int64
	_ = v749
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int64
	_ = v757
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int64
	_ = v767
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int64
	_ = v807
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int64
	_ = v818
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int64
	_ = v828
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int64
	_ = v838
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int64
	_ = v848
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1043 int32
	_ = v1043
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1066 int64
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1078 int64
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1084 int64
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1090 int64
	_ = v1090
	var v1096 int64
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1146 int64
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1335 int64
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1340 int64
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1347 int64
	_ = v1347
	var v1348 int64
	_ = v1348
	var v1351 int64
	_ = v1351
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1403 int64
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1417 int32
	_ = v1417
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1464 int64
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1520 int64
	_ = v1520
	var v1524 int32
	_ = v1524
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1620 int32
	_ = v1620
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int64
	_ = v1715
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1721 int64
	_ = v1721
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1734 int32
	_ = v1734
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1744 int64
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1761 int32
	_ = v1761
	var v1768 int32
	_ = v1768
	var v1775 int32
	_ = v1775
	var v1783 int32
	_ = v1783
	var v1791 int32
	_ = v1791
	var v1799 int32
	_ = v1799
	var v1807 int32
	_ = v1807
	var v1815 int32
	_ = v1815
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1837 int32
	_ = v1837
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1855 int32
	_ = v1855
	var v1862 int32
	_ = v1862
	var v1873 int32
	_ = v1873
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1896 int32
	_ = v1896
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1912 int32
	_ = v1912
	var v1919 int32
	_ = v1919
	var v1927 int32
	_ = v1927
	var v1935 int32
	_ = v1935
	var v1943 int32
	_ = v1943
	var v1951 int32
	_ = v1951
	var v1959 int32
	_ = v1959
	var v1965 int32
	_ = v1965
	var v1968 int32
	_ = v1968
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1979 int64
	_ = v1979
	var v1983 int32
	_ = v1983
	var v1987 int32
	_ = v1987
	var v1995 int32
	_ = v1995
	var v2003 int32
	_ = v2003
	var v2008 int32
	_ = v2008
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2016 int32
	_ = v2016
	var v2019 int32
	_ = v2019
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2029 int32
	_ = v2029
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2045 int32
	_ = v2045
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2051 int32
	_ = v2051
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2059 int32
	_ = v2059
	var v2068 int32
	_ = v2068
	var v2079 int32
	_ = v2079
	var v2084 int32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2110 int32
	_ = v2110
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2135 int32
	_ = v2135
	var v2145 int32
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2153 int64
	_ = v2153
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2166 int64
	_ = v2166
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2189 int32
	_ = v2189
	var v2194 int32
	_ = v2194
	var v2201 int32
	_ = v2201
	var v2204 int32
	_ = v2204
	var v2212 int32
	_ = v2212
	var v2218 int32
	_ = v2218
	var v2222 int32
	_ = v2222
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2240 int64
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2244 int64
	_ = v2244
	var v2245 int64
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2263 int32
	_ = v2263
	var v2271 int32
	_ = v2271
	var v2287 int32
	_ = v2287
	var v2289 int32
	_ = v2289
	var v2292 int64
	_ = v2292
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2326 int32
	_ = v2326
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2333 int64
	_ = v2333
	var v2337 int32
	_ = v2337
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2346 int32
	_ = v2346
	var v2348 int32
	_ = v2348
	var v2352 int32
	_ = v2352
	var v2353 int32
	_ = v2353
	var v2359 int32
	_ = v2359
	var v2363 int64
	_ = v2363
	var v2367 int32
	_ = v2367
	var v2371 int32
	_ = v2371
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2380 int32
	_ = v2380
	var v2390 int32
	_ = v2390
	var v2396 int32
	_ = v2396
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2440 int32
	_ = v2440
	var v2441 int64
	_ = v2441
	var v2444 int64
	_ = v2444
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2458 int32
	_ = v2458
	var v2459 int64
	_ = v2459
	var v2462 int64
	_ = v2462
	var v2469 int32
	_ = v2469
	var v2470 int64
	_ = v2470
	var v2483 int32
	_ = v2483
	var v2489 int64
	_ = v2489
	var v2490 int64
	_ = v2490
	var v2491 int64
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2493 int64
	_ = v2493
	var v2494 int64
	_ = v2494
	var v2495 int64
	_ = v2495
	var v2496 int64
	_ = v2496
	var v2497 int64
	_ = v2497
	var v2502 int32
	_ = v2502
	var v2537 int32
	_ = v2537
	var v2539 int32
	_ = v2539
	var v2546 int32
	_ = v2546
	var v2547 int32
	_ = v2547
	var v2552 int32
	_ = v2552
	var v2558 int32
	_ = v2558
	v31 = m.G0
	v33 = v31 - int32(1456)
	m.G0 = v33
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(967)
	v38 = *(*int64)(unsafe.Add(mBase, _consts[641]))
	*(*uint32)(unsafe.Add(mBase, uint32(l0)+44)) = uint32(v38)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v40&int32(5) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v33 + int32(1456)
	return v2558
L2:
	;
	v2537 = int32(3)
	v2539 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v2537 < v2539 {
		goto L648
	} else {
		goto L649
	}
L3:
	;
	v51 = v33 + int32(416)
	v52 = int32(9)
	goto L4
L4:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v76) < base.Ui32(v52) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v99 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+425)) = uint8(v99)
	v102 = v33 + int32(416)
	v103 = int32(_a140)
	v104 = int32(6)
	goto L25
L6:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v89 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	v78 = v76
	goto L9
L8:
	;
	v78 = v52
	goto L9
L9:
	;
	if v76 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v79 = v78
	goto L12
L11:
	;
	v79 = v52
	goto L12
L12:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = m.T0[v80].(func(*base.Module, int32, int32, int32) int32)(m, l0, v51, v79)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v81 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v85 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v85 | int64(1)
	goto L2
L16:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v94 + v79
	v98 = v52 - v79
	if v98 != 0 {
		v51 = v51 + v79
		v52 = v98
		goto L4
	} else {
		goto L19
	}
L17:
	;
	m.T0[v89].(func(*base.Module, int32, int32, int32))(m, l0, v51, v79)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	goto L5
L20:
	;
	v264 = v33 + int32(416) | int32(6)
	goto L60
L21:
	;
	if v168 == int32(0) {
		goto L20
	} else {
		goto L37
	}
L22:
	;
	v168 = int32(0)
	goto L21
L23:
	;
	v140 = v135
	v141 = v136
	v142 = v137
	goto L33
L24:
	;
	if v125 == int32(0) {
		goto L22
	} else {
		goto L31
	}
L25:
	;
	if (v103|v102)&int32(3) != 0 {
		v135 = v102
		v136 = v103
		v137 = v104
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v112 = v102
	v113 = v103
	v114 = v104
	goto L27
L27:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if v117 != v118 {
		v135 = v112
		v136 = v113
		v137 = v114
		goto L23
	} else {
		goto L29
	}
L28:
	;
	goto L24
L29:
	;
	v120 = int32(4)
	v121 = v113 + v120
	v123 = v112 + v120
	v125 = v114 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v125) {
		v112 = v123
		v113 = v121
		v114 = v125
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v135 = v123
	v136 = v121
	v137 = v125
	goto L23
L32:
	;
	v168 = v145 - v146
	goto L21
L33:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v145 != v146 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v148 = int32(1)
	v153 = v142 + int32(-1)
	if v153 == int32(0) {
		goto L22
	} else {
		goto L36
	}
L36:
	;
	v140 = v140 + v148
	v141 = v141 + v148
	v142 = v153
	goto L33
L37:
	;
	v172 = v33 + int32(416)
	v173 = int32(_a141)
	v174 = int32(6)
	goto L42
L38:
	;
	if v238 == int32(0) {
		goto L20
	} else {
		goto L54
	}
L39:
	;
	v238 = int32(0)
	goto L38
L40:
	;
	v210 = v205
	v211 = v206
	v212 = v207
	goto L50
L41:
	;
	if v195 == int32(0) {
		goto L39
	} else {
		goto L48
	}
L42:
	;
	if (v173|v172)&int32(3) != 0 {
		v205 = v172
		v206 = v173
		v207 = v174
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v182 = v172
	v183 = v173
	v184 = v174
	goto L44
L44:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v187 != v188 {
		v205 = v182
		v206 = v183
		v207 = v184
		goto L40
	} else {
		goto L46
	}
L45:
	;
	goto L41
L46:
	;
	v190 = int32(4)
	v191 = v183 + v190
	v193 = v182 + v190
	v195 = v184 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v195) {
		v182 = v193
		v183 = v191
		v184 = v195
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v205 = v193
	v206 = v191
	v207 = v195
	goto L40
L49:
	;
	v238 = v215 - v216
	goto L38
L50:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v215 != v216 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v218 = int32(1)
	v223 = v212 + int32(-1)
	if v223 == int32(0) {
		goto L39
	} else {
		goto L53
	}
L53:
	;
	v210 = v210 + v218
	v211 = v211 + v218
	v212 = v223
	goto L50
L54:
	;
	v241 = int32(2)
	v243 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v243 {
		v2558 = v241
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+352)) = v33 + int32(416)
	F__serverLog(m, int32(3), int32(_a1131), v33+int32(352))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	v2558 = v241
	goto L1
L57:
	;
	if l1&int32(32) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L58:
	;
	v328 = int32(2)
	v330 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v330 {
		v2558 = v328
		goto L1
	} else {
		goto L80
	}
L59:
	;
	if base.B2i32(v168 == int32(0))&base.B2i32(int32(79) < v309) != 0 {
		goto L58
	} else {
		goto L74
	}
L60:
	;
	v269 = v264 + int32(1)
	v270 = int32(*(*int8)(unsafe.Add(mBase, uint32(v264))))
	v271 = F___isspace_1(m, v270)
	mBase = m.M
	if v271 != 0 {
		v264 = v269
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v272 = int32(1)
	switch v270&int32(255) + int32(-43) {
	case 0:
		v278 = v272
		goto L64
	default:
		v280 = v264
		v281 = v270
		v282 = v272
		goto L63
	case 2:
		goto L65
	}
L62:
	;
	goto L61
L63:
	;
	v285 = v281 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v285) {
		v303 = int32(0)
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v279 = int32(*(*int8)(unsafe.Add(mBase, uint32(v269))))
	v280 = v269
	v281 = v279
	v282 = v278
	goto L63
L65:
	;
	v278 = int32(0)
	goto L64
L66:
	;
	if v282 != 0 {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	v289 = int32(0)
	v290 = v280
	v291 = v285
	goto L68
L68:
	;
	v293 = int32(10)
	v295 = v289*v293 - v291
	v296 = int32(*(*int8)(unsafe.Add(mBase, uint32(v290)+1)))
	v300 = v296 + int32(-48)
	if base.Ui32(v300) < base.Ui32(v293) {
		v289 = v295
		v290 = v290 + int32(1)
		v291 = v300
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v303 = v295
	goto L66
L70:
	;
	goto L69
L71:
	;
	v309 = int32(0) - v303
	goto L73
L72:
	;
	v309 = v303
	goto L73
L73:
	;
	goto L59
L74:
	;
	if v309 < int32(1) {
		goto L58
	} else {
		goto L75
	}
L75:
	;
	if base.B2i32(v309 < int32(80))&base.B2i32(v168 != int32(0)) != 0 {
		goto L58
	} else {
		goto L76
	}
L76:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	if v321 != 0 {
		goto L57
	} else {
		goto L77
	}
L77:
	;
	if base.Ui32(int32(80)) < base.Ui32(v309) {
		goto L58
	} else {
		goto L78
	}
L78:
	;
	if base.Ui32(int32(67)) < base.Ui32(v309+int32(-12)) {
		goto L57
	} else {
		goto L79
	}
L79:
	;
	goto L58
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v309
	F__serverLog(m, int32(3), int32(_a1132), v33)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	v2558 = v328
	goto L1
L82:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	if v364 != 0 {
		v370 = v364
		goto L89
	} else {
		goto L90
	}
L83:
	;
	v342 = int32(_a44)
	v343 = *(*int32)(unsafe.Add(mBase, _consts[210]))
	v347 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v347 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v357 = F_emptyData(m, int32(-1), base.B2i32(v343 != int32(0)), int32(968))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L13
	} else {
		goto L87
	}
L85:
	;
	F__serverLog(m, int32(2), int32(_a1133), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L13
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v360 = *(*int32)(unsafe.Add(mBase, _consts[423]))
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v360
	goto L82
L89:
	;
	v371 = F_mstime(m)
	mBase = m.M
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v372&int32(5) != 0 {
		goto L2
	} else {
		goto L92
	}
L90:
	;
	v366 = F_createDatabase(m, int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L13
	} else {
		goto L91
	}
L91:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v368))) = v366
	v370 = v366
	goto L89
L92:
	;
	v377 = int32(8)
	v383 = int32(3)
	v384 = int64(0)
	v390 = int64(-1)
	v404 = v370
	v410 = v384
	v411 = v384
	v412 = v384
	v413 = int32(0)
	v414 = v384
	v415 = v384
	v416 = v390
	v417 = v390
	v418 = v390
	goto L93
L93:
	;
	v432 = int32(1)
	v434 = v33 + int32(360)
	goto L95
L94:
	;
	goto L2
L95:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v456) < base.Ui32(v432) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+360)))
	if v168 != 0 {
		goto L110
	} else {
		goto L111
	}
L97:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v467 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L98:
	;
	v458 = v456
	goto L100
L99:
	;
	v458 = v432
	goto L100
L100:
	;
	if v456 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v459 = v458
	goto L103
L102:
	;
	v459 = v432
	goto L103
L103:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v461 = m.T0[v460].(func(*base.Module, int32, int32, int32) int32)(m, l0, v434, v459)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L13
	} else {
		goto L104
	}
L104:
	;
	if v461 != 0 {
		goto L97
	} else {
		goto L105
	}
L105:
	;
	v463 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v463 | int64(1)
	goto L2
L106:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v472 + v459
	v476 = v432 - v459
	if v476 != 0 {
		v432 = v476
		v434 = v434 + v459
		goto L95
	} else {
		goto L109
	}
L107:
	;
	m.T0[v467].(func(*base.Module, int32, int32, int32))(m, l0, v434, v459)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L13
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	goto L96
L110:
	;
	switch v477 + int32(-243) {
	case 0:
		goto L124
	case 1:
		goto L130
	case 2:
		goto L126
	case 3:
		goto L127
	case 4:
		goto L128
	case 5:
		goto L133
	case 6:
		goto L134
	case 7:
		goto L129
	case 8:
		goto L131
	case 9:
		goto L135
	case 10:
		goto L136
	case 11:
		goto L132
	case 12:
		goto L120
	default:
		goto L125
	}
L111:
	;
	if base.Ui32(int32(221)) < base.Ui32((v477+int32(-22))&int32(255)) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v485 {
		v2558 = v383
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v477
	F__serverLog(m, int32(3), int32(_a1134), v33+int32(16))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L13
	} else {
		goto L114
	}
L114:
	;
	v2558 = v383
	goto L1
L115:
	;
	v2502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v2502&int32(5) == int32(0) {
		v404 = v2483
		v410 = v2489
		v411 = v2490
		v412 = v2491
		v413 = v2492
		v414 = v2493
		v415 = v2494
		v416 = v2495
		v417 = v2496
		v418 = v2497
		goto L93
	} else {
		goto L647
	}
L116:
	;
	v2470 = int64(-1)
	v2483 = v404
	v2489 = v410
	v2490 = v411
	v2491 = v412
	v2492 = v2246
	v2493 = v2244
	v2494 = v2245
	v2495 = v2470
	v2496 = v2470
	v2497 = int64(-1)
	goto L115
L117:
	;
	v2428 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v414 == int64(0) {
		goto L641
	} else {
		goto L642
	}
L118:
	;
	v2289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v2289&int32(5) != 0 {
		goto L2
	} else {
		goto L610
	}
L119:
	;
	F__serverAssert(m, int32(_a1135), int32(_a1122), int32(3534))
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L13
	} else {
		goto L609
	}
L120:
	;
	if int32(4) < v309 {
		goto L118
	} else {
		goto L608
	}
L121:
	;
	if v2271 == int32(9) {
		v2483 = v404
		v2489 = v410
		v2490 = v411
		v2491 = v412
		v2492 = v413
		v2493 = v414
		v2494 = v415
		v2495 = v416
		v2496 = v417
		v2497 = v418
		goto L115
	} else {
		goto L606
	}
L122:
	;
	v2271 = int32(9)
	goto L121
L123:
	;
	v2260 = F_rdbLoadCheckModuleValue(m, l0, v33+int32(392))
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L13
	} else {
		goto L604
	}
L124:
	;
	v2254 = F_clusterRDBLoadSlotImport(m, l0)
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L13
	} else {
		goto L602
	}
L125:
	;
	if v413 == int32(0) {
		goto L530
	} else {
		goto L531
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+360)) = int32(0)
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v2011 = F_rdbFunctionLoad(m, l0, l0, v2008, l1, v33+int32(360))
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L13
	} else {
		goto L524
	}
L127:
	;
	F_rdbReportError(m, int32(1), int32(3471), int32(_a1136), int32(0))
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L13
	} else {
		goto L523
	}
L128:
	;
	v1713 = F_rdbLoadLenByRef(m, l0, int32(0), v33+int32(360))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L13
	} else {
		goto L476
	}
L129:
	;
	v872 = int32(0)
	v874 = F_rdbGenericLoadStringObject(m, l0, v872, v872)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L13
	} else {
		goto L224
	}
L130:
	;
	v824 = F_rdbLoadLenByRef(m, l0, int32(0), v33+int32(360))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L13
	} else {
		goto L208
	}
L131:
	;
	v803 = F_rdbLoadLenByRef(m, l0, int32(0), v33+int32(360))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L13
	} else {
		goto L202
	}
L132:
	;
	v763 = F_rdbLoadLenByRef(m, l0, int32(0), v33+int32(360))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L13
	} else {
		goto L192
	}
L133:
	;
	v753 = F_rdbLoadLenByRef(m, l0, int32(0), v33+int32(360))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L13
	} else {
		goto L189
	}
L134:
	;
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v692&int32(5) != 0 {
		goto L2
	} else {
		goto L173
	}
L135:
	;
	v596 = int64(9223372036854775807)
	v597 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	if v597&int64(5) != int64(0) {
		v683 = v596
		v685 = v597
		goto L155
	} else {
		goto L156
	}
L136:
	;
	v498 = int64(-1000)
	v499 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	if v499&int64(5) != int64(0) {
		v587 = v498
		v589 = v499
		goto L137
	} else {
		goto L138
	}
L137:
	;
	if v589&int64(1) == int64(0) {
		v2483 = v404
		v2489 = v410
		v2490 = v411
		v2491 = v412
		v2492 = v413
		v2493 = v414
		v2494 = v415
		v2495 = v416
		v2496 = v417
		v2497 = v587
		goto L115
	} else {
		goto L154
	}
L138:
	;
	v513 = int32(4)
	v515 = v33 + int32(360)
	goto L139
L139:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v537) < base.Ui32(v513) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v558 = int64(*(*int32)(unsafe.Add(mBase, uint32(v33)+360)))
	v561 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v587 = v558 * int64(1000)
	v589 = v561
	goto L137
L141:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v548 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L142:
	;
	v539 = v537
	goto L144
L143:
	;
	v539 = v513
	goto L144
L144:
	;
	if v537 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v540 = v539
	goto L147
L146:
	;
	v540 = v513
	goto L147
L147:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v542 = m.T0[v541].(func(*base.Module, int32, int32, int32) int32)(m, l0, v515, v540)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L13
	} else {
		goto L148
	}
L148:
	;
	if v542 != 0 {
		goto L141
	} else {
		goto L149
	}
L149:
	;
	v544 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v546 = v544 | int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v546
	v587 = v498
	v589 = v546
	goto L137
L150:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v553 + v540
	v557 = v513 - v540
	if v557 != 0 {
		v513 = v557
		v515 = v515 + v540
		goto L139
	} else {
		goto L153
	}
L151:
	;
	m.T0[v548].(func(*base.Module, int32, int32, int32))(m, l0, v515, v540)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L13
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	goto L140
L154:
	;
	goto L2
L155:
	;
	if v685&int64(1) == int64(0) {
		v2483 = v404
		v2489 = v410
		v2490 = v411
		v2491 = v412
		v2492 = v413
		v2493 = v414
		v2494 = v415
		v2495 = v416
		v2496 = v417
		v2497 = v683
		goto L115
	} else {
		goto L172
	}
L156:
	;
	v611 = int32(8)
	v613 = v33 + int32(360)
	goto L157
L157:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v635) < base.Ui32(v611) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	v656 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v657 = *(*int64)(unsafe.Add(mBase, uint32(v33)+360))
	v683 = v657
	v685 = v656
	goto L155
L159:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v646 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L160:
	;
	v637 = v635
	goto L162
L161:
	;
	v637 = v611
	goto L162
L162:
	;
	if v635 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v638 = v637
	goto L165
L164:
	;
	v638 = v611
	goto L165
L165:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v640 = m.T0[v639].(func(*base.Module, int32, int32, int32) int32)(m, l0, v613, v638)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L13
	} else {
		goto L166
	}
L166:
	;
	if v640 != 0 {
		goto L159
	} else {
		goto L167
	}
L167:
	;
	v642 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v644 = v642 | int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v644
	v683 = v596
	v685 = v644
	goto L155
L168:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v651 + v638
	v655 = v611 - v638
	if v655 != 0 {
		v611 = v655
		v613 = v613 + v638
		goto L157
	} else {
		goto L171
	}
L169:
	;
	m.T0[v646].(func(*base.Module, int32, int32, int32))(m, l0, v613, v638)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L13
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	goto L158
L172:
	;
	goto L2
L173:
	;
	v704 = int32(1)
	v706 = v33 + int32(360)
	goto L174
L174:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v728) < base.Ui32(v704) {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v749 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v33)+360)))
	v2483 = v404
	v2489 = v410
	v2490 = v411
	v2491 = v412
	v2492 = v413
	v2493 = v414
	v2494 = v415
	v2495 = v416
	v2496 = v749
	v2497 = v418
	goto L115
L176:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v739 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L177:
	;
	v730 = v728
	goto L179
L178:
	;
	v730 = v704
	goto L179
L179:
	;
	if v728 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v731 = v730
	goto L182
L181:
	;
	v731 = v704
	goto L182
L182:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v733 = m.T0[v732].(func(*base.Module, int32, int32, int32) int32)(m, l0, v706, v731)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L13
	} else {
		goto L183
	}
L183:
	;
	if v733 != 0 {
		goto L176
	} else {
		goto L184
	}
L184:
	;
	v735 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v735 | int64(1)
	goto L2
L185:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v744 + v731
	v748 = v704 - v731
	if v748 != 0 {
		v704 = v748
		v706 = v706 + v731
		goto L174
	} else {
		goto L188
	}
L186:
	;
	m.T0[v739].(func(*base.Module, int32, int32, int32))(m, l0, v706, v731)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L13
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	goto L175
L189:
	;
	if v753 == int32(-1) {
		goto L2
	} else {
		goto L190
	}
L190:
	;
	v757 = *(*int64)(unsafe.Add(mBase, uint32(v33)+360))
	if v757 != int64(-1) {
		v2483 = v404
		v2489 = v410
		v2490 = v411
		v2491 = v412
		v2492 = v413
		v2493 = v414
		v2494 = v415
		v2495 = v757
		v2496 = v417
		v2497 = v418
		goto L115
	} else {
		goto L191
	}
L191:
	;
	goto L2
L192:
	;
	if v763 == int32(-1) {
		goto L2
	} else {
		goto L193
	}
L193:
	;
	v767 = *(*int64)(unsafe.Add(mBase, uint32(v33)+360))
	if v767 == int64(-1) {
		goto L2
	} else {
		goto L194
	}
L194:
	;
	v771 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if base.Ui64(v767) < base.Ui64(base.I64_extend_i32_u(v771)) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v790 = base.I32_wrap_i64(v767)
	v792 = v790 << (uint(int32(2)) % 32)
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v789+v792)))
	if v794 != 0 {
		v2483 = v794
		v2489 = v767
		v2490 = v411
		v2491 = v412
		v2492 = v413
		v2493 = v414
		v2494 = v415
		v2495 = v416
		v2496 = v417
		v2497 = v418
		goto L115
	} else {
		goto L200
	}
L196:
	;
	v775 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v775 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = int32(_a46)
	F__serverLog(m, int32(3), int32(_a1137), v33+int32(160))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L13
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	v795 = F_createDatabase(m, v790)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L13
	} else {
		goto L201
	}
L201:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v797+v792))) = v795
	v2483 = v795
	v2489 = v767
	v2490 = v411
	v2491 = v412
	v2492 = v413
	v2493 = v414
	v2494 = v415
	v2495 = v416
	v2496 = v417
	v2497 = v418
	goto L115
L202:
	;
	if v803 == int32(-1) {
		goto L2
	} else {
		goto L203
	}
L203:
	;
	v807 = *(*int64)(unsafe.Add(mBase, uint32(v33)+360))
	if v807 == int64(-1) {
		goto L2
	} else {
		goto L204
	}
L204:
	;
	v813 = F_rdbLoadLenByRef(m, l0, int32(0), v33+int32(360))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L13
	} else {
		goto L205
	}
L205:
	;
	if v813 == int32(-1) {
		goto L2
	} else {
		goto L206
	}
L206:
	;
	v818 = *(*int64)(unsafe.Add(mBase, uint32(v33)+360))
	if v818 != int64(-1) {
		v2483 = v404
		v2489 = v410
		v2490 = v807
		v2491 = v818
		v2492 = int32(1)
		v2493 = v414
		v2494 = v415
		v2495 = v416
		v2496 = v417
		v2497 = v418
		goto L115
	} else {
		goto L207
	}
L207:
	;
	goto L2
L208:
	;
	if v824 == int32(-1) {
		goto L2
	} else {
		goto L209
	}
L209:
	;
	v828 = *(*int64)(unsafe.Add(mBase, uint32(v33)+360))
	if v828 == int64(-1) {
		goto L2
	} else {
		goto L210
	}
L210:
	;
	v834 = F_rdbLoadLenByRef(m, l0, int32(0), v33+int32(360))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L13
	} else {
		goto L211
	}
L211:
	;
	if v834 == int32(-1) {
		goto L2
	} else {
		goto L212
	}
L212:
	;
	v838 = *(*int64)(unsafe.Add(mBase, uint32(v33)+360))
	if v838 == int64(-1) {
		goto L2
	} else {
		goto L213
	}
L213:
	;
	v844 = F_rdbLoadLenByRef(m, l0, int32(0), v33+int32(360))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L13
	} else {
		goto L214
	}
L214:
	;
	if v844 == int32(-1) {
		goto L2
	} else {
		goto L215
	}
L215:
	;
	v848 = *(*int64)(unsafe.Add(mBase, uint32(v33)+360))
	if v848 == int64(-1) {
		goto L2
	} else {
		goto L216
	}
L216:
	;
	if base.Ui64(int64(16383)) < base.Ui64(v828) {
		v2483 = v404
		v2489 = v410
		v2490 = v411
		v2491 = v412
		v2492 = v413
		v2493 = v414
		v2494 = v415
		v2495 = v416
		v2496 = v417
		v2497 = v418
		goto L115
	} else {
		goto L217
	}
L217:
	;
	v854 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v854 == int32(0) {
		v2483 = v404
		v2489 = v410
		v2490 = v411
		v2491 = v412
		v2492 = v413
		v2493 = v414
		v2494 = v415
		v2495 = v416
		v2496 = v417
		v2497 = v418
		goto L115
	} else {
		goto L218
	}
L218:
	;
	if v838 == int64(0) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v864 = int32(0)
	if v848 == int64(0) {
		v2483 = v404
		v2489 = v410
		v2490 = v411
		v2491 = v412
		v2492 = v864
		v2493 = v414
		v2494 = v415
		v2495 = v416
		v2496 = v417
		v2497 = v418
		goto L115
	} else {
		goto L222
	}
L220:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v862 = F_kvstoreHashtableExpand(m, v859, base.I32_wrap_i64(v828), base.I32_wrap_i64(v838))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L13
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	v870 = F_kvstoreHashtableExpand(m, v867, base.I32_wrap_i64(v828), base.I32_wrap_i64(v848))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L13
	} else {
		goto L223
	}
L223:
	;
	v2483 = v404
	v2489 = v410
	v2490 = v411
	v2491 = v412
	v2492 = v864
	v2493 = v414
	v2494 = v415
	v2495 = v416
	v2496 = v417
	v2497 = v418
	goto L115
L224:
	;
	if v874 == int32(0) {
		goto L2
	} else {
		goto L225
	}
L225:
	;
	v878 = int32(0)
	v880 = F_rdbGenericLoadStringObject(m, l0, v878, v878)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L13
	} else {
		goto L227
	}
L226:
	;
	v884 = F_objectGetVal(m, v874)
	mBase = m.M
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v884))))
	if v885 != int32(37) {
		goto L231
	} else {
		goto L232
	}
L227:
	;
	if v880 != 0 {
		goto L226
	} else {
		goto L228
	}
L228:
	;
	F_decrRefCount(m, v874)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L13
	} else {
		goto L229
	}
L229:
	;
	goto L2
L230:
	;
	F_decrRefCount(m, v874)
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L13
	} else {
		goto L474
	}
L231:
	;
	v902 = F_objectGetVal(m, v874)
	mBase = m.M
	v903 = int32(_a1138)
	v906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902))))
	if v906 != 0 {
		goto L238
	} else {
		goto L239
	}
L232:
	;
	v889 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v889 {
		v1703 = v413
		goto L230
	} else {
		goto L233
	}
L233:
	;
	v892 = F_objectGetVal(m, v874)
	mBase = m.M
	v893 = F_objectGetVal(m, v880)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v893
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v892
	F__serverLog(m, int32(2), int32(_a1139), v33+int32(176))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L13
	} else {
		goto L234
	}
L234:
	;
	v1703 = v413
	goto L230
L235:
	;
	v995 = F_objectGetVal(m, v874)
	mBase = m.M
	v996 = int32(_a1140)
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995))))
	if v999 != 0 {
		goto L268
	} else {
		goto L269
	}
L236:
	;
	if v938-v940 != 0 {
		goto L235
	} else {
		goto L248
	}
L237:
	;
	v938 = F_tolower(m, v934)
	mBase = m.M
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935))))
	v940 = F_tolower(m, v939)
	mBase = m.M
	goto L236
L238:
	;
	v908 = v902
	v909 = v903
	v910 = v906
	goto L241
L239:
	;
	v934 = int32(0)
	v935 = v903
	goto L237
L240:
	;
	v934 = v931 & int32(255)
	v935 = v930
	goto L237
L241:
	;
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909))))
	if v912 == int32(0) {
		v930 = v909
		v931 = v910
		goto L240
	} else {
		goto L243
	}
L242:
	;
	v930 = v924
	v931 = int32(0)
	goto L240
L243:
	;
	v916 = v910 & int32(255)
	if v916 == v912 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v923 = int32(1)
	v924 = v909 + v923
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908)+1)))
	if v925 != 0 {
		v908 = v908 + v923
		v909 = v924
		v910 = v925
		goto L241
	} else {
		goto L247
	}
L245:
	;
	v918 = F_tolower(m, v916)
	mBase = m.M
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909))))
	v920 = F_tolower(m, v919)
	mBase = m.M
	if v918 == v920 {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908))))
	v930 = v909
	v931 = v922
	goto L240
L247:
	;
	goto L242
L248:
	;
	if l2 == int32(0) {
		v1703 = v413
		goto L230
	} else {
		goto L249
	}
L249:
	;
	v944 = F_objectGetVal(m, v880)
	mBase = m.M
	v948 = v944
	goto L251
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v993
	v1703 = v413
	goto L230
L251:
	;
	v953 = v948 + int32(1)
	v954 = int32(*(*int8)(unsafe.Add(mBase, uint32(v948))))
	v955 = F___isspace_1(m, v954)
	mBase = m.M
	if v955 != 0 {
		v948 = v953
		goto L251
	} else {
		goto L253
	}
L252:
	;
	v956 = int32(1)
	switch v954&int32(255) + int32(-43) {
	case 0:
		v962 = v956
		goto L255
	default:
		v964 = v948
		v965 = v954
		v966 = v956
		goto L254
	case 2:
		goto L256
	}
L253:
	;
	goto L252
L254:
	;
	v969 = v965 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v969) {
		v987 = int32(0)
		goto L257
	} else {
		goto L258
	}
L255:
	;
	v963 = int32(*(*int8)(unsafe.Add(mBase, uint32(v953))))
	v964 = v953
	v965 = v963
	v966 = v962
	goto L254
L256:
	;
	v962 = int32(0)
	goto L255
L257:
	;
	if v966 != 0 {
		goto L262
	} else {
		goto L263
	}
L258:
	;
	v973 = int32(0)
	v974 = v964
	v975 = v969
	goto L259
L259:
	;
	v977 = int32(10)
	v979 = v973*v977 - v975
	v980 = int32(*(*int8)(unsafe.Add(mBase, uint32(v974)+1)))
	v984 = v980 + int32(-48)
	if base.Ui32(v984) < base.Ui32(v977) {
		v973 = v979
		v974 = v974 + int32(1)
		v975 = v984
		goto L259
	} else {
		goto L261
	}
L260:
	;
	v987 = v979
	goto L257
L261:
	;
	goto L260
L262:
	;
	v993 = int32(0) - v987
	goto L264
L263:
	;
	v993 = v987
	goto L264
L264:
	;
	goto L250
L265:
	;
	v1100 = F_objectGetVal(m, v874)
	mBase = m.M
	v1101 = int32(_a1141)
	v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1100))))
	if v1104 != 0 {
		goto L291
	} else {
		goto L292
	}
L266:
	;
	if v1031-v1033 != 0 {
		goto L265
	} else {
		goto L278
	}
L267:
	;
	v1031 = F_tolower(m, v1027)
	mBase = m.M
	v1032 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1028))))
	v1033 = F_tolower(m, v1032)
	mBase = m.M
	goto L266
L268:
	;
	v1001 = v995
	v1002 = v996
	v1003 = v999
	goto L271
L269:
	;
	v1027 = int32(0)
	v1028 = v996
	goto L267
L270:
	;
	v1027 = v1024 & int32(255)
	v1028 = v1023
	goto L267
L271:
	;
	v1005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1002))))
	if v1005 == int32(0) {
		v1023 = v1002
		v1024 = v1003
		goto L270
	} else {
		goto L273
	}
L272:
	;
	v1023 = v1017
	v1024 = int32(0)
	goto L270
L273:
	;
	v1009 = v1003 & int32(255)
	if v1009 == v1005 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v1016 = int32(1)
	v1017 = v1002 + v1016
	v1018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001)+1)))
	if v1018 != 0 {
		v1001 = v1001 + v1016
		v1002 = v1017
		v1003 = v1018
		goto L271
	} else {
		goto L277
	}
L275:
	;
	v1011 = F_tolower(m, v1009)
	mBase = m.M
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1002))))
	v1013 = F_tolower(m, v1012)
	mBase = m.M
	if v1011 == v1013 {
		goto L274
	} else {
		goto L276
	}
L276:
	;
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001))))
	v1023 = v1002
	v1024 = v1015
	goto L270
L277:
	;
	goto L272
L278:
	;
	if l2 == int32(0) {
		v1703 = v413
		goto L230
	} else {
		goto L279
	}
L279:
	;
	v1037 = F_objectGetVal(m, v880)
	mBase = m.M
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1037+int32(-1)))))
	switch v1043 & int32(7) {
	case 0:
		goto L286
	case 1:
		goto L285
	case 2:
		goto L284
	case 3:
		goto L283
	case 4:
		goto L282
	default:
		v1060 = int32(0)
		goto L281
	}
L280:
	;
	if v1062 != int32(40) {
		v1703 = v413
		goto L230
	} else {
		goto L287
	}
L281:
	;
	v1062 = v1060
	goto L280
L282:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1037+int32(-17))))
	v1060 = v1059
	goto L281
L283:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1037+int32(-9))))
	v1062 = v1056
	goto L280
L284:
	;
	v1053 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1037+int32(-5)))))
	v1062 = v1053
	goto L280
L285:
	;
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1037+int32(-3)))))
	v1062 = v1050
	goto L280
L286:
	;
	v1062 = int32(base.Ui32(v1043) >> (uint(int32(3)) % 32))
	goto L280
L287:
	;
	v1065 = F_objectGetVal(m, v880)
	mBase = m.M
	v1066 = *(*int64)(unsafe.Add(mBase, uint32(v1065)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v377))) = v1066
	v1068 = int32(40)
	v1072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1065+v1068))))
	*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(48)))) = uint8(v1072)
	v1074 = int32(32)
	v1078 = *(*int64)(unsafe.Add(mBase, uint32(v1065+v1074)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v1068))) = v1078
	v1080 = int32(24)
	v1084 = *(*int64)(unsafe.Add(mBase, uint32(v1065+v1080)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v1074))) = v1084
	v1086 = int32(16)
	v1090 = *(*int64)(unsafe.Add(mBase, uint32(v1065+v1086)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v1080))) = v1090
	v1096 = *(*int64)(unsafe.Add(mBase, uint32(v1065+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v1086))) = v1096
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = int32(1)
	v1703 = v413
	goto L230
L288:
	;
	v1148 = F_objectGetVal(m, v874)
	mBase = m.M
	v1149 = int32(_a775)
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1148))))
	if v1152 != 0 {
		goto L306
	} else {
		goto L307
	}
L289:
	;
	if v1136-v1138 != 0 {
		goto L288
	} else {
		goto L301
	}
L290:
	;
	v1136 = F_tolower(m, v1132)
	mBase = m.M
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1133))))
	v1138 = F_tolower(m, v1137)
	mBase = m.M
	goto L289
L291:
	;
	v1106 = v1100
	v1107 = v1101
	v1108 = v1104
	goto L294
L292:
	;
	v1132 = int32(0)
	v1133 = v1101
	goto L290
L293:
	;
	v1132 = v1129 & int32(255)
	v1133 = v1128
	goto L290
L294:
	;
	v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1107))))
	if v1110 == int32(0) {
		v1128 = v1107
		v1129 = v1108
		goto L293
	} else {
		goto L296
	}
L295:
	;
	v1128 = v1122
	v1129 = int32(0)
	goto L293
L296:
	;
	v1114 = v1108 & int32(255)
	if v1114 == v1110 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1121 = int32(1)
	v1122 = v1107 + v1121
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1106)+1)))
	if v1123 != 0 {
		v1106 = v1106 + v1121
		v1107 = v1122
		v1108 = v1123
		goto L294
	} else {
		goto L300
	}
L298:
	;
	v1116 = F_tolower(m, v1114)
	mBase = m.M
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1107))))
	v1118 = F_tolower(m, v1117)
	mBase = m.M
	if v1116 == v1118 {
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1106))))
	v1128 = v1107
	v1129 = v1120
	goto L293
L300:
	;
	goto L295
L301:
	;
	if l2 == int32(0) {
		v1703 = v413
		goto L230
	} else {
		goto L302
	}
L302:
	;
	v1142 = F_objectGetVal(m, v880)
	mBase = m.M
	v1146 = F_strtox_2(m, v1142, int32(0), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L303
L303:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+56)) = v1146
	v1703 = v413
	goto L230
L304:
	;
	if v1184-v1186 == int32(0) {
		v1703 = v413
		goto L230
	} else {
		goto L316
	}
L305:
	;
	v1184 = F_tolower(m, v1180)
	mBase = m.M
	v1185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1181))))
	v1186 = F_tolower(m, v1185)
	mBase = m.M
	goto L304
L306:
	;
	v1154 = v1148
	v1155 = v1149
	v1156 = v1152
	goto L309
L307:
	;
	v1180 = int32(0)
	v1181 = v1149
	goto L305
L308:
	;
	v1180 = v1177 & int32(255)
	v1181 = v1176
	goto L305
L309:
	;
	v1158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1155))))
	if v1158 == int32(0) {
		v1176 = v1155
		v1177 = v1156
		goto L308
	} else {
		goto L311
	}
L310:
	;
	v1176 = v1170
	v1177 = int32(0)
	goto L308
L311:
	;
	v1162 = v1156 & int32(255)
	if v1162 == v1158 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1169 = int32(1)
	v1170 = v1155 + v1169
	v1171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1154)+1)))
	if v1171 != 0 {
		v1154 = v1154 + v1169
		v1155 = v1170
		v1156 = v1171
		goto L309
	} else {
		goto L315
	}
L313:
	;
	v1164 = F_tolower(m, v1162)
	mBase = m.M
	v1165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1155))))
	v1166 = F_tolower(m, v1165)
	mBase = m.M
	if v1164 == v1166 {
		goto L312
	} else {
		goto L314
	}
L314:
	;
	v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1154))))
	v1176 = v1155
	v1177 = v1168
	goto L308
L315:
	;
	goto L310
L316:
	;
	v1190 = F_objectGetVal(m, v874)
	mBase = m.M
	v1191 = int32(_a1142)
	v1194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1190))))
	if v1194 != 0 {
		goto L320
	} else {
		goto L321
	}
L317:
	;
	v1242 = F_objectGetVal(m, v874)
	mBase = m.M
	v1243 = int32(_a1143)
	v1246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242))))
	if v1246 != 0 {
		goto L336
	} else {
		goto L337
	}
L318:
	;
	if v1226-v1228 != 0 {
		goto L317
	} else {
		goto L330
	}
L319:
	;
	v1226 = F_tolower(m, v1222)
	mBase = m.M
	v1227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1223))))
	v1228 = F_tolower(m, v1227)
	mBase = m.M
	goto L318
L320:
	;
	v1196 = v1190
	v1197 = v1191
	v1198 = v1194
	goto L323
L321:
	;
	v1222 = int32(0)
	v1223 = v1191
	goto L319
L322:
	;
	v1222 = v1219 & int32(255)
	v1223 = v1218
	goto L319
L323:
	;
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1197))))
	if v1200 == int32(0) {
		v1218 = v1197
		v1219 = v1198
		goto L322
	} else {
		goto L325
	}
L324:
	;
	v1218 = v1212
	v1219 = int32(0)
	goto L322
L325:
	;
	v1204 = v1198 & int32(255)
	if v1204 == v1200 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1211 = int32(1)
	v1212 = v1197 + v1211
	v1213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1196)+1)))
	if v1213 != 0 {
		v1196 = v1196 + v1211
		v1197 = v1212
		v1198 = v1213
		goto L323
	} else {
		goto L329
	}
L327:
	;
	v1206 = F_tolower(m, v1204)
	mBase = m.M
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1197))))
	v1208 = F_tolower(m, v1207)
	mBase = m.M
	if v1206 == v1208 {
		goto L326
	} else {
		goto L328
	}
L328:
	;
	v1210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1196))))
	v1218 = v1197
	v1219 = v1210
	goto L322
L329:
	;
	goto L324
L330:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v1231 {
		v1703 = v413
		goto L230
	} else {
		goto L331
	}
L331:
	;
	v1234 = F_objectGetVal(m, v880)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1234
	F__serverLog(m, int32(2), int32(_a1144), v33+int32(192))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L13
	} else {
		goto L332
	}
L332:
	;
	v1703 = v413
	goto L230
L333:
	;
	v1294 = F_objectGetVal(m, v874)
	mBase = m.M
	v1295 = int32(_a1145)
	v1298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294))))
	if v1298 != 0 {
		goto L352
	} else {
		goto L353
	}
L334:
	;
	if v1278-v1280 != 0 {
		goto L333
	} else {
		goto L346
	}
L335:
	;
	v1278 = F_tolower(m, v1274)
	mBase = m.M
	v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1275))))
	v1280 = F_tolower(m, v1279)
	mBase = m.M
	goto L334
L336:
	;
	v1248 = v1242
	v1249 = v1243
	v1250 = v1246
	goto L339
L337:
	;
	v1274 = int32(0)
	v1275 = v1243
	goto L335
L338:
	;
	v1274 = v1271 & int32(255)
	v1275 = v1270
	goto L335
L339:
	;
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1249))))
	if v1252 == int32(0) {
		v1270 = v1249
		v1271 = v1250
		goto L338
	} else {
		goto L341
	}
L340:
	;
	v1270 = v1264
	v1271 = int32(0)
	goto L338
L341:
	;
	v1256 = v1250 & int32(255)
	if v1256 == v1252 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1263 = int32(1)
	v1264 = v1249 + v1263
	v1265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248)+1)))
	if v1265 != 0 {
		v1248 = v1248 + v1263
		v1249 = v1264
		v1250 = v1265
		goto L339
	} else {
		goto L345
	}
L343:
	;
	v1258 = F_tolower(m, v1256)
	mBase = m.M
	v1259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1249))))
	v1260 = F_tolower(m, v1259)
	mBase = m.M
	if v1258 == v1260 {
		goto L342
	} else {
		goto L344
	}
L344:
	;
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248))))
	v1270 = v1249
	v1271 = v1262
	goto L338
L345:
	;
	goto L340
L346:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v1283 {
		v1703 = v413
		goto L230
	} else {
		goto L347
	}
L347:
	;
	v1286 = F_objectGetVal(m, v880)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1286
	F__serverLog(m, int32(2), int32(_a1146), v33+int32(208))
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L13
	} else {
		goto L348
	}
L348:
	;
	v1703 = v413
	goto L230
L349:
	;
	v1359 = F_objectGetVal(m, v874)
	mBase = m.M
	v1360 = int32(_a1147)
	v1363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359))))
	if v1363 != 0 {
		goto L372
	} else {
		goto L373
	}
L350:
	;
	if v1330-v1332 != 0 {
		goto L349
	} else {
		goto L362
	}
L351:
	;
	v1330 = F_tolower(m, v1326)
	mBase = m.M
	v1331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1327))))
	v1332 = F_tolower(m, v1331)
	mBase = m.M
	goto L350
L352:
	;
	v1300 = v1294
	v1301 = v1295
	v1302 = v1298
	goto L355
L353:
	;
	v1326 = int32(0)
	v1327 = v1295
	goto L351
L354:
	;
	v1326 = v1323 & int32(255)
	v1327 = v1322
	goto L351
L355:
	;
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1301))))
	if v1304 == int32(0) {
		v1322 = v1301
		v1323 = v1302
		goto L354
	} else {
		goto L357
	}
L356:
	;
	v1322 = v1316
	v1323 = int32(0)
	goto L354
L357:
	;
	v1308 = v1302 & int32(255)
	if v1308 == v1304 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1315 = int32(1)
	v1316 = v1301 + v1315
	v1317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1300)+1)))
	if v1317 != 0 {
		v1300 = v1300 + v1315
		v1301 = v1316
		v1302 = v1317
		goto L355
	} else {
		goto L361
	}
L359:
	;
	v1310 = F_tolower(m, v1308)
	mBase = m.M
	v1311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1301))))
	v1312 = F_tolower(m, v1311)
	mBase = m.M
	if v1310 == v1312 {
		goto L358
	} else {
		goto L360
	}
L360:
	;
	v1314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1300))))
	v1322 = v1301
	v1323 = v1314
	goto L354
L361:
	;
	goto L356
L362:
	;
	v1334 = int32(0)
	v1335 = F___time(m, v1334)
	mBase = m.M
	v1336 = F_objectGetVal(m, v880)
	mBase = m.M
	v1340 = F_strtox_2(m, v1336, v1334, int32(10), int64(2147483648))
	mBase = m.M
	goto L363
L363:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v1343 {
		v1703 = v413
		goto L230
	} else {
		goto L364
	}
L364:
	;
	v1347 = v1335 - base.I64_extend_i32_s(base.I32_wrap_i64(v1340))
	v1348 = int64(0)
	if v1348 < v1347 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1351 = v1347
	goto L367
L366:
	;
	v1351 = v1348
	goto L367
L367:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+224)) = uint32(v1351)
	F__serverLog(m, int32(2), int32(_a1148), v33+int32(224))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L13
	} else {
		goto L368
	}
L368:
	;
	v1703 = v413
	goto L230
L369:
	;
	v1420 = F_objectGetVal(m, v874)
	mBase = m.M
	v1421 = int32(_a1149)
	v1424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1420))))
	if v1424 != 0 {
		goto L390
	} else {
		goto L391
	}
L370:
	;
	if v1395-v1397 != 0 {
		goto L369
	} else {
		goto L382
	}
L371:
	;
	v1395 = F_tolower(m, v1391)
	mBase = m.M
	v1396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1392))))
	v1397 = F_tolower(m, v1396)
	mBase = m.M
	goto L370
L372:
	;
	v1365 = v1359
	v1366 = v1360
	v1367 = v1363
	goto L375
L373:
	;
	v1391 = int32(0)
	v1392 = v1360
	goto L371
L374:
	;
	v1391 = v1388 & int32(255)
	v1392 = v1387
	goto L371
L375:
	;
	v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1366))))
	if v1369 == int32(0) {
		v1387 = v1366
		v1388 = v1367
		goto L374
	} else {
		goto L377
	}
L376:
	;
	v1387 = v1381
	v1388 = int32(0)
	goto L374
L377:
	;
	v1373 = v1367 & int32(255)
	if v1373 == v1369 {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v1380 = int32(1)
	v1381 = v1366 + v1380
	v1382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1365)+1)))
	if v1382 != 0 {
		v1365 = v1365 + v1380
		v1366 = v1381
		v1367 = v1382
		goto L375
	} else {
		goto L381
	}
L379:
	;
	v1375 = F_tolower(m, v1373)
	mBase = m.M
	v1376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1366))))
	v1377 = F_tolower(m, v1376)
	mBase = m.M
	if v1375 == v1377 {
		goto L378
	} else {
		goto L380
	}
L380:
	;
	v1379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1365))))
	v1387 = v1366
	v1388 = v1379
	goto L374
L381:
	;
	goto L376
L382:
	;
	v1399 = F_objectGetVal(m, v880)
	mBase = m.M
	v1403 = F_strtox_2(m, v1399, int32(0), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L383
L383:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v1405 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	*(*int64)(unsafe.Add(mBase, _consts[638])) = v1403
	v1703 = v413
	goto L230
L385:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v33)+240)) = base.F64_mul(base.F64_convert_i64_s(v1403), float64(9.5367431640625e-07))
	F__serverLog(m, int32(2), int32(_a1150), v33+int32(240))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L13
	} else {
		goto L386
	}
L386:
	;
	goto L384
L387:
	;
	v1476 = F_objectGetVal(m, v874)
	mBase = m.M
	v1477 = int32(_a1151)
	v1480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1476))))
	if v1480 != 0 {
		goto L408
	} else {
		goto L409
	}
L388:
	;
	if v1456-v1458 != 0 {
		goto L387
	} else {
		goto L400
	}
L389:
	;
	v1456 = F_tolower(m, v1452)
	mBase = m.M
	v1457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1453))))
	v1458 = F_tolower(m, v1457)
	mBase = m.M
	goto L388
L390:
	;
	v1426 = v1420
	v1427 = v1421
	v1428 = v1424
	goto L393
L391:
	;
	v1452 = int32(0)
	v1453 = v1421
	goto L389
L392:
	;
	v1452 = v1449 & int32(255)
	v1453 = v1448
	goto L389
L393:
	;
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1427))))
	if v1430 == int32(0) {
		v1448 = v1427
		v1449 = v1428
		goto L392
	} else {
		goto L395
	}
L394:
	;
	v1448 = v1442
	v1449 = int32(0)
	goto L392
L395:
	;
	v1434 = v1428 & int32(255)
	if v1434 == v1430 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v1441 = int32(1)
	v1442 = v1427 + v1441
	v1443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1426)+1)))
	if v1443 != 0 {
		v1426 = v1426 + v1441
		v1427 = v1442
		v1428 = v1443
		goto L393
	} else {
		goto L399
	}
L397:
	;
	v1436 = F_tolower(m, v1434)
	mBase = m.M
	v1437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1427))))
	v1438 = F_tolower(m, v1437)
	mBase = m.M
	if v1436 == v1438 {
		goto L396
	} else {
		goto L398
	}
L398:
	;
	v1440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1426))))
	v1448 = v1427
	v1449 = v1440
	goto L392
L399:
	;
	goto L394
L400:
	;
	v1460 = F_objectGetVal(m, v880)
	mBase = m.M
	v1464 = F_strtox_2(m, v1460, int32(0), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L401
L401:
	;
	if v1464 == int64(0) {
		v1703 = v413
		goto L230
	} else {
		goto L402
	}
L402:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v1468 {
		v1703 = v413
		goto L230
	} else {
		goto L403
	}
L403:
	;
	F__serverLog(m, int32(2), int32(_a1152), int32(0))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L13
	} else {
		goto L404
	}
L404:
	;
	v1703 = v413
	goto L230
L405:
	;
	v1532 = F_objectGetVal(m, v874)
	mBase = m.M
	v1533 = int32(_a1153)
	v1536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1532))))
	if v1536 != 0 {
		goto L425
	} else {
		goto L426
	}
L406:
	;
	if v1512-v1514 != 0 {
		goto L405
	} else {
		goto L418
	}
L407:
	;
	v1512 = F_tolower(m, v1508)
	mBase = m.M
	v1513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1509))))
	v1514 = F_tolower(m, v1513)
	mBase = m.M
	goto L406
L408:
	;
	v1482 = v1476
	v1483 = v1477
	v1484 = v1480
	goto L411
L409:
	;
	v1508 = int32(0)
	v1509 = v1477
	goto L407
L410:
	;
	v1508 = v1505 & int32(255)
	v1509 = v1504
	goto L407
L411:
	;
	v1486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1483))))
	if v1486 == int32(0) {
		v1504 = v1483
		v1505 = v1484
		goto L410
	} else {
		goto L413
	}
L412:
	;
	v1504 = v1498
	v1505 = int32(0)
	goto L410
L413:
	;
	v1490 = v1484 & int32(255)
	if v1490 == v1486 {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v1497 = int32(1)
	v1498 = v1483 + v1497
	v1499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1482)+1)))
	if v1499 != 0 {
		v1482 = v1482 + v1497
		v1483 = v1498
		v1484 = v1499
		goto L411
	} else {
		goto L417
	}
L415:
	;
	v1492 = F_tolower(m, v1490)
	mBase = m.M
	v1493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1483))))
	v1494 = F_tolower(m, v1493)
	mBase = m.M
	if v1492 == v1494 {
		goto L414
	} else {
		goto L416
	}
L416:
	;
	v1496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1482))))
	v1504 = v1483
	v1505 = v1496
	goto L410
L417:
	;
	goto L412
L418:
	;
	v1516 = F_objectGetVal(m, v880)
	mBase = m.M
	v1520 = F_strtox_2(m, v1516, int32(0), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L419
L419:
	;
	if v1520 == int64(0) {
		v1703 = v413
		goto L230
	} else {
		goto L420
	}
L420:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v1524 {
		v1703 = v413
		goto L230
	} else {
		goto L421
	}
L421:
	;
	F__serverLog(m, int32(2), int32(_a1154), int32(0))
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L13
	} else {
		goto L422
	}
L422:
	;
	v1703 = v413
	goto L230
L423:
	;
	if v1568-v1570 == int32(0) {
		v1703 = v413
		goto L230
	} else {
		goto L435
	}
L424:
	;
	v1568 = F_tolower(m, v1564)
	mBase = m.M
	v1569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1565))))
	v1570 = F_tolower(m, v1569)
	mBase = m.M
	goto L423
L425:
	;
	v1538 = v1532
	v1539 = v1533
	v1540 = v1536
	goto L428
L426:
	;
	v1564 = int32(0)
	v1565 = v1533
	goto L424
L427:
	;
	v1564 = v1561 & int32(255)
	v1565 = v1560
	goto L424
L428:
	;
	v1542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1539))))
	if v1542 == int32(0) {
		v1560 = v1539
		v1561 = v1540
		goto L427
	} else {
		goto L430
	}
L429:
	;
	v1560 = v1554
	v1561 = int32(0)
	goto L427
L430:
	;
	v1546 = v1540 & int32(255)
	if v1546 == v1542 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1553 = int32(1)
	v1554 = v1539 + v1553
	v1555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1538)+1)))
	if v1555 != 0 {
		v1538 = v1538 + v1553
		v1539 = v1554
		v1540 = v1555
		goto L428
	} else {
		goto L434
	}
L432:
	;
	v1548 = F_tolower(m, v1546)
	mBase = m.M
	v1549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1539))))
	v1550 = F_tolower(m, v1549)
	mBase = m.M
	if v1548 == v1550 {
		goto L431
	} else {
		goto L433
	}
L433:
	;
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1538))))
	v1560 = v1539
	v1561 = v1552
	goto L427
L434:
	;
	goto L429
L435:
	;
	v1574 = F_objectGetVal(m, v874)
	mBase = m.M
	v1575 = int32(_a1155)
	v1578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1574))))
	if v1578 != 0 {
		goto L441
	} else {
		goto L442
	}
L436:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(0) < v1691 {
		v1703 = v413
		goto L230
	} else {
		goto L472
	}
L437:
	;
	F_decrRefCount(m, v874)
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L13
	} else {
		goto L470
	}
L438:
	;
	v1668 = int32(0)
	v1669 = *(*int32)(unsafe.Add(mBase, _consts[632]))
	if v1669 == v1668 {
		goto L436
	} else {
		goto L464
	}
L439:
	;
	if v1610-v1612 != 0 {
		goto L438
	} else {
		goto L451
	}
L440:
	;
	v1610 = F_tolower(m, v1606)
	mBase = m.M
	v1611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1607))))
	v1612 = F_tolower(m, v1611)
	mBase = m.M
	goto L439
L441:
	;
	v1580 = v1574
	v1581 = v1575
	v1582 = v1578
	goto L444
L442:
	;
	v1606 = int32(0)
	v1607 = v1575
	goto L440
L443:
	;
	v1606 = v1603 & int32(255)
	v1607 = v1602
	goto L440
L444:
	;
	v1584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1581))))
	if v1584 == int32(0) {
		v1602 = v1581
		v1603 = v1582
		goto L443
	} else {
		goto L446
	}
L445:
	;
	v1602 = v1596
	v1603 = int32(0)
	goto L443
L446:
	;
	v1588 = v1582 & int32(255)
	if v1588 == v1584 {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v1595 = int32(1)
	v1596 = v1581 + v1595
	v1597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1580)+1)))
	if v1597 != 0 {
		v1580 = v1580 + v1595
		v1581 = v1596
		v1582 = v1597
		goto L444
	} else {
		goto L450
	}
L448:
	;
	v1590 = F_tolower(m, v1588)
	mBase = m.M
	v1591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1581))))
	v1592 = F_tolower(m, v1591)
	mBase = m.M
	if v1590 == v1592 {
		goto L447
	} else {
		goto L449
	}
L449:
	;
	v1594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1580))))
	v1602 = v1581
	v1603 = v1594
	goto L443
L450:
	;
	goto L445
L451:
	;
	v1614 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+392)) = v1614
	*(*int32)(unsafe.Add(mBase, uint32(v33)+1448)) = v1614
	*(*int32)(unsafe.Add(mBase, uint32(v33)+404)) = v1614
	v1620 = F_objectGetVal(m, v880)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v33)+268)) = v33 + int32(404)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+264)) = v33 + int32(1448)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+260)) = v33 + int32(392)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+256)) = v33 + int32(360)
	v1636 = F_sscanf(m, v1620, int32(_a1156), v33+int32(256))
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L13
	} else {
		goto L452
	}
L452:
	;
	if v1636 < int32(3) {
		goto L437
	} else {
		goto L453
	}
L453:
	;
	v1641 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v1641 == int32(0) {
		v1703 = v413
		goto L230
	} else {
		goto L454
	}
L454:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v33)+360))
	if base.Ui32(int32(16383)) < base.Ui32(v1644) {
		v1703 = v413
		goto L230
	} else {
		goto L455
	}
L455:
	;
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v33)+392))
	if v1647 == int32(0) {
		goto L456
	} else {
		goto L457
	}
L456:
	;
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v33)+1448))
	if v1653 == int32(0) {
		goto L459
	} else {
		goto L460
	}
L457:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v1651 = F_kvstoreHashtableExpand(m, v1650, v1644, v1647)
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L13
	} else {
		goto L458
	}
L458:
	;
	goto L456
L459:
	;
	v1660 = int32(0)
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v33)+404))
	if v1661 == v1660 {
		v1703 = v1660
		goto L230
	} else {
		goto L462
	}
L460:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v33)+360))
	v1658 = F_kvstoreHashtableExpand(m, v1656, v1657, v1653)
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L13
	} else {
		goto L461
	}
L461:
	;
	goto L459
L462:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v404)+8))
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v33)+360))
	v1666 = F_kvstoreHashtableExpand(m, v1664, v1665, v1661)
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L13
	} else {
		goto L463
	}
L463:
	;
	v1703 = v1660
	goto L230
L464:
	;
	v1672 = F_objectGetVal(m, v874)
	mBase = m.M
	v1673 = F_dictFind(m, v1669, v1672)
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L13
	} else {
		goto L465
	}
L465:
	;
	if v1673 == int32(0) {
		goto L436
	} else {
		goto L466
	}
L466:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+8))
	goto L467
L467:
	;
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+4))
	v1679 = F_objectGetVal(m, v880)
	mBase = m.M
	v1680 = m.T0[v1678].(func(*base.Module, int32, int32) int32)(m, l1, v1679)
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L13
	} else {
		goto L468
	}
L468:
	;
	if v1680 != int32(-1) {
		v1703 = v413
		goto L230
	} else {
		goto L469
	}
L469:
	;
	goto L437
L470:
	;
	F_decrRefCount(m, v880)
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L13
	} else {
		goto L471
	}
L471:
	;
	goto L2
L472:
	;
	v1694 = F_objectGetVal(m, v874)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v33)+272)) = v1694
	F__serverLog(m, int32(0), int32(_a1157), v33+int32(272))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L13
	} else {
		goto L473
	}
L473:
	;
	v1703 = v413
	goto L230
L474:
	;
	F_decrRefCount(m, v880)
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L13
	} else {
		goto L475
	}
L475:
	;
	v2483 = v404
	v2489 = v410
	v2490 = v411
	v2491 = v412
	v2492 = v1703
	v2493 = v414
	v2494 = v415
	v2495 = v416
	v2496 = v417
	v2497 = v418
	goto L115
L476:
	;
	v1715 = *(*int64)(unsafe.Add(mBase, uint32(v33)+360))
	v1719 = F_rdbLoadLenByRef(m, l0, int32(0), v33+int32(360))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L13
	} else {
		goto L477
	}
L477:
	;
	v1721 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v33)+360)))
	v1725 = F_rdbLoadLenByRef(m, l0, int32(0), v33+int32(360))
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L13
	} else {
		goto L478
	}
L478:
	;
	v1727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v1727&int32(1) != 0 {
		goto L2
	} else {
		goto L479
	}
L479:
	;
	if v1719 == int32(-1) {
		goto L481
	} else {
		goto L482
	}
L480:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v33)+360))
	if v1713 == int32(-1) {
		goto L485
	} else {
		goto L486
	}
L481:
	;
	v1734 = int32(0)
	F_rdbReportError(m, v1734, int32(3417), int32(_a1158), v1734)
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L13
	} else {
		goto L484
	}
L482:
	;
	if v1721 == int64(2) {
		goto L480
	} else {
		goto L483
	}
L483:
	;
	goto L481
L484:
	;
	goto L2
L485:
	;
	v1744 = int64(-1)
	goto L487
L486:
	;
	v1744 = v1715
	goto L487
L487:
	;
	v1745 = F_moduleTypeLookupModuleByID(m, v1744)
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L13
	} else {
		goto L488
	}
L488:
	;
	v1748 = v33 + int32(392)
	v1749 = int32(0)
	v1752 = *(*int32)(unsafe.Add(mBase, _consts[556]))
	*(*uint8)(unsafe.Add(mBase, uint32(v1748)+9)) = uint8(v1749)
	v1755 = base.I32_wrap_i64(v1744)
	v1758 = int32(63)
	v1761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1752+int32(base.Ui32(v1755)>>(uint(int32(10))%32))&v1758))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1748)+8)) = uint8(v1761)
	v1768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1752+int32(base.Ui32(v1755)>>(uint(int32(16))%32))&v1758))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1748)+7)) = uint8(v1768)
	v1775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1752+int32(base.Ui32(v1755)>>(uint(int32(22))%32))&v1758))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1748)+6)) = uint8(v1775)
	v1783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1752+base.I32_wrap_i64(int64(base.Ui64(v1744)>>(uint(int64(28))%64)))&v1758))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1748)+5)) = uint8(v1783)
	v1791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1752+base.I32_wrap_i64(int64(base.Ui64(v1744)>>(uint(int64(34))%64)))&v1758))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1748)+4)) = uint8(v1791)
	v1799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1752+base.I32_wrap_i64(int64(base.Ui64(v1744)>>(uint(int64(40))%64)))&v1758))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1748)+3)) = uint8(v1799)
	v1807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1752+base.I32_wrap_i64(int64(base.Ui64(v1744)>>(uint(int64(46))%64)))&v1758))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1748)+2)) = uint8(v1807)
	v1815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1752+base.I32_wrap_i64(int64(base.Ui64(v1744)>>(uint(int64(52))%64)))&v1758))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1748)+1)) = uint8(v1815)
	v1821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1752+base.I32_wrap_i64(int64(base.Ui64(v1744)>>(uint(int64(58))%64)))))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1748))) = uint8(v1821)
	goto L489
L489:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	if v1824 != 0 {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	if v1824 != 0 {
		goto L123
	} else {
		goto L496
	}
L491:
	;
	if v1745 != 0 {
		goto L490
	} else {
		goto L492
	}
L492:
	;
	v1826 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v1826 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+288)) = v33 + int32(392)
	F__serverLog(m, int32(3), int32(_a1159), v33+int32(288))
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L13
	} else {
		goto L495
	}
L495:
	;
	goto L493
L496:
	;
	if v1745 == int32(0) {
		goto L123
	} else {
		goto L497
	}
L497:
	;
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1745)+52))
	if v1842 != 0 {
		goto L498
	} else {
		goto L499
	}
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+368)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+364)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v33)+380)) = int64(-4294967296)
	v1862 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+360)) = v1862
	*(*int32)(unsafe.Add(mBase, uint32(v33)+388)) = v1862
	*(*int64)(unsafe.Add(mBase, uint32(v33)+372)) = int64(0)
	v1873 = int32(-1)
	if v1725 == v1873 {
		goto L503
	} else {
		goto L504
	}
L499:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v1844 {
		goto L500
	} else {
		goto L501
	}
L500:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+304)) = v33 + int32(392)
	F__serverLog(m, int32(3), int32(_a1160), v33+int32(304))
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L13
	} else {
		goto L502
	}
L502:
	;
	goto L500
L503:
	;
	v1876 = v1873
	goto L505
L504:
	;
	v1876 = v1740
	goto L505
L505:
	;
	v1877 = m.T0[v1842].(func(*base.Module, int32, int32, int32) int32)(m, v33+int32(360), base.I32_wrap_i64(v1744)&int32(1023), v1876)
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L13
	} else {
		goto L506
	}
L506:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v33)+376))
	if v1879 == int32(0) {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v33)+372))
	if v1877|v1887 == int32(0) {
		goto L513
	} else {
		goto L514
	}
L508:
	;
	F_moduleFreeContext(m, v1879)
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L13
	} else {
		goto L509
	}
L509:
	;
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v33)+376))
	F_valkey_free(m, v1884)
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L13
	} else {
		goto L510
	}
L510:
	;
	goto L507
L511:
	;
	v2271 = int32(2)
	goto L121
L512:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+320)) = v33 + int32(392)
	F__serverLog(m, int32(3), v1987, v33+int32(320))
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L13
	} else {
		goto L522
	}
L513:
	;
	v1975 = F_rdbLoadLenByRef(m, l0, int32(0), v33+int32(1448))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L13
	} else {
		goto L518
	}
L514:
	;
	v1892 = v33 + int32(392)
	v1893 = int32(0)
	v1896 = *(*int32)(unsafe.Add(mBase, _consts[556]))
	*(*uint8)(unsafe.Add(mBase, uint32(v1892)+9)) = uint8(v1893)
	v1899 = base.I32_wrap_i64(v1744)
	v1902 = int32(63)
	v1905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1896+int32(base.Ui32(v1899)>>(uint(int32(10))%32))&v1902))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1892)+8)) = uint8(v1905)
	v1912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1896+int32(base.Ui32(v1899)>>(uint(int32(16))%32))&v1902))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1892)+7)) = uint8(v1912)
	v1919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1896+int32(base.Ui32(v1899)>>(uint(int32(22))%32))&v1902))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1892)+6)) = uint8(v1919)
	v1927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1896+base.I32_wrap_i64(int64(base.Ui64(v1744)>>(uint(int64(28))%64)))&v1902))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1892)+5)) = uint8(v1927)
	v1935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1896+base.I32_wrap_i64(int64(base.Ui64(v1744)>>(uint(int64(34))%64)))&v1902))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1892)+4)) = uint8(v1935)
	v1943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1896+base.I32_wrap_i64(int64(base.Ui64(v1744)>>(uint(int64(40))%64)))&v1902))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1892)+3)) = uint8(v1943)
	v1951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1896+base.I32_wrap_i64(int64(base.Ui64(v1744)>>(uint(int64(46))%64)))&v1902))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1892)+2)) = uint8(v1951)
	v1959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1896+base.I32_wrap_i64(int64(base.Ui64(v1744)>>(uint(int64(52))%64)))&v1902))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1892)+1)) = uint8(v1959)
	v1965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1896+base.I32_wrap_i64(int64(base.Ui64(v1744)>>(uint(int64(58))%64)))))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1892))) = uint8(v1965)
	goto L515
L515:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v1968 {
		goto L511
	} else {
		goto L516
	}
L516:
	;
	v1987 = int32(_a1161)
	goto L512
L517:
	;
	v1983 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v1983 {
		goto L511
	} else {
		goto L521
	}
L518:
	;
	if v1975 == int32(-1) {
		goto L517
	} else {
		goto L519
	}
L519:
	;
	v1979 = *(*int64)(unsafe.Add(mBase, uint32(v33)+1448))
	if v1979 == int64(0) {
		goto L122
	} else {
		goto L520
	}
L520:
	;
	goto L517
L521:
	;
	v1987 = int32(_a1162)
	goto L512
L522:
	;
	goto L511
L523:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L524:
	;
	if v2011 == int32(0) {
		goto L122
	} else {
		goto L525
	}
L525:
	;
	v2016 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v2016 {
		goto L526
	} else {
		goto L527
	}
L526:
	;
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v33)+360))
	F_sdsfree(m, v2027)
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L13
	} else {
		goto L529
	}
L527:
	;
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v33)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+336)) = v2019
	F__serverLog(m, int32(3), int32(_a1163), v33+int32(336))
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L13
	} else {
		goto L528
	}
L528:
	;
	goto L526
L529:
	;
	v2271 = int32(2)
	goto L121
L530:
	;
	v2041 = F_rdbGenericLoadStringObject(m, l0, int32(4), int32(0))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L13
	} else {
		goto L534
	}
L531:
	;
	v2034 = F_dbExpand(m, v404, v411, int32(0))
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L13
	} else {
		goto L532
	}
L532:
	;
	v2037 = F_dbExpandExpires(m, v404, v412, int32(0))
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L13
	} else {
		goto L533
	}
L533:
	;
	goto L530
L534:
	;
	if v2041 == int32(0) {
		goto L2
	} else {
		goto L535
	}
L535:
	;
	v2045 = *(*int32)(unsafe.Add(mBase, uint32(v404)+28))
	v2048 = F_rdbLoadObject(m, v477, l0, v2041, v2045, v33+int32(412), l1, v371)
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L13
	} else {
		goto L536
	}
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+408)) = v2048
	if v2048 != 0 {
		goto L539
	} else {
		goto L540
	}
L537:
	;
	v2246 = int32(0)
	v2249 = *(*int32)(unsafe.Add(mBase, _consts[73]))
	if v2249 == v2246 {
		goto L116
	} else {
		goto L600
	}
L538:
	;
	F_sdsfree(m, v2041)
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L13
	} else {
		goto L599
	}
L539:
	;
	v2101 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v2101 != 0 {
		goto L560
	} else {
		goto L561
	}
L540:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v33)+412))
	switch v2051 + int32(-1) {
	case 0:
		goto L544
	case 1:
		goto L543
	default:
		goto L541
	case 3:
		goto L542
	}
L541:
	;
	F_sdsfree(m, v2041)
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L13
	} else {
		goto L557
	}
L542:
	;
	F_sdsfree(m, v2041)
	mBase = m.M
	v2095 = m.ExcPending
	if v2095 != 0 {
		goto L13
	} else {
		goto L556
	}
L543:
	;
	F_sdsfree(m, v2041)
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L13
	} else {
		goto L553
	}
L544:
	;
	if int64(9) < v414 {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v2240 = v414 + int64(1)
	goto L538
L546:
	;
	v2056 = int32(_a44)
	v2057 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v2059 = *(*int32)(unsafe.Add(mBase, _consts[209]))
	if v2059 == int32(0) {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	if int32(2) < v2057 {
		goto L545
	} else {
		goto L551
	}
L548:
	;
	if int32(2) < v2057 {
		goto L545
	} else {
		goto L549
	}
L549:
	;
	F__serverLog(m, int32(2), int32(_a1164), int32(0))
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		goto L13
	} else {
		goto L550
	}
L550:
	;
	v2240 = v414 + int64(1)
	goto L538
L551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v2041
	F__serverLog(m, int32(2), int32(_a1165), v33+int32(32))
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L13
	} else {
		goto L552
	}
L552:
	;
	goto L545
L553:
	;
	v2086 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v2086 {
		v2558 = v383
		goto L1
	} else {
		goto L554
	}
L554:
	;
	F__serverLog(m, int32(3), int32(_a1166), int32(0))
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L13
	} else {
		goto L555
	}
L555:
	;
	v2558 = v383
	goto L1
L556:
	;
	v2244 = v414
	v2245 = v415 + int64(1)
	goto L537
L557:
	;
	goto L2
L558:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+368)) = v2041
	*(*int64)(unsafe.Add(mBase, uint32(v33)+360)) = int64(-68719476736)
	v2162 = F_dbAddRDBLoad(m, v404, v2041, v33+int32(408))
	mBase = m.M
	v2163 = m.ExcPending
	if v2163 != 0 {
		goto L13
	} else {
		goto L576
	}
L559:
	;
	if v2110 == int32(0) {
		goto L558
	} else {
		goto L562
	}
L560:
	;
	v2106 = F_getMyClusterNode(m)
	mBase = m.M
	v2107 = F_clusterNodeIsPrimary(m, v2106)
	mBase = m.M
	v2110 = base.B2i32(v2107 != int32(0))
	goto L559
L561:
	;
	v2102 = int32(0)
	v2103 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v2110 = base.B2i32(v2103 == v2102)
	goto L559
L562:
	;
	if l1&int32(1) != 0 {
		goto L558
	} else {
		goto L563
	}
L563:
	;
	if v418 == int64(-1) {
		goto L558
	} else {
		goto L564
	}
L564:
	;
	if v371 <= v418 {
		goto L558
	} else {
		goto L565
	}
L565:
	;
	if l1&v377 == int32(0) {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	F_sdsfree(m, v2041)
	mBase = m.M
	v2147 = m.ExcPending
	if v2147 != 0 {
		goto L13
	} else {
		goto L574
	}
L567:
	;
	v2119 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	if v2119 == int32(0) {
		goto L119
	} else {
		goto L568
	}
L568:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v2123)+20))
	if v2124 != 0 {
		goto L119
	} else {
		goto L569
	}
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+368)) = v2041
	*(*int64)(unsafe.Add(mBase, uint32(v33)+360)) = int64(-68719476736)
	v2132 = *(*int32)(unsafe.Add(mBase, _consts[350]))
	if v2132 != 0 {
		goto L570
	} else {
		goto L571
	}
L570:
	;
	v2133 = int32(244)
	goto L572
L571:
	;
	v2133 = int32(240)
	goto L572
L572:
	;
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v2133)+uint32(_consts[84])))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+392)) = v2135
	*(*int32)(unsafe.Add(mBase, uint32(v33)+396)) = v33 + int32(360)
	F_replicationFeedReplicas(m, base.I32_wrap_i64(v410), v33+int32(392), int32(2))
	mBase = m.M
	v2145 = m.ExcPending
	if v2145 != 0 {
		goto L13
	} else {
		goto L573
	}
L573:
	;
	goto L566
L574:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v33)+408))
	F_decrRefCount(m, v2148)
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		goto L13
	} else {
		goto L575
	}
L575:
	;
	v2151 = int32(_a44)
	v2153 = *(*int64)(unsafe.Add(mBase, _consts[637]))
	*(*int64)(unsafe.Add(mBase, _consts[637])) = v2153 + int64(1)
	v2244 = v414
	v2245 = v415
	goto L537
L576:
	;
	v2164 = int32(_a44)
	v2166 = *(*int64)(unsafe.Add(mBase, _consts[640]))
	*(*int64)(unsafe.Add(mBase, _consts[640])) = v2166 + int64(1)
	if v2162 != 0 {
		goto L577
	} else {
		goto L578
	}
L577:
	;
	if v418 != int64(-1) {
		goto L594
	} else {
		goto L595
	}
L578:
	;
	if l1&int32(4) == int32(0) {
		goto L579
	} else {
		goto L580
	}
L579:
	;
	v2186 = int32(_a44)
	v2187 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v2189 = *(*int32)(unsafe.Add(mBase, _consts[209]))
	if v2189 == int32(0) {
		goto L586
	} else {
		goto L587
	}
L580:
	;
	v2174 = F_dbSyncDelete(m, v404, v33+int32(360))
	mBase = m.M
	v2175 = m.ExcPending
	if v2175 != 0 {
		goto L13
	} else {
		goto L581
	}
L581:
	;
	v2178 = F_dbAddRDBLoad(m, v404, v2041, v33+int32(408))
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L13
	} else {
		goto L582
	}
L582:
	;
	if v2178 != 0 {
		goto L577
	} else {
		goto L583
	}
L583:
	;
	F__serverAssert(m, int32(_a1167), int32(_a1122), int32(3559))
	mBase = m.M
	v2184 = m.ExcPending
	if v2184 != 0 {
		goto L13
	} else {
		goto L584
	}
L584:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L585:
	;
	F__serverPanic_1(m, int32(_a1122), int32(3566), int32(_a1168), int32(0))
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L13
	} else {
		goto L592
	}
L586:
	;
	if int32(3) < v2187 {
		goto L585
	} else {
		goto L590
	}
L587:
	;
	if int32(3) < v2187 {
		goto L585
	} else {
		goto L588
	}
L588:
	;
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v404)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = v2194
	F__serverLog(m, int32(3), int32(_a1169), v33+int32(64))
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L13
	} else {
		goto L589
	}
L589:
	;
	goto L585
L590:
	;
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(v404)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v2041
	F__serverLog(m, int32(3), int32(_a1170), v33+int32(48))
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L13
	} else {
		goto L591
	}
L591:
	;
	goto L585
L592:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L593:
	;
	v2230 = F_objectSetLRUOrLFU(m, v2229, v417, v416)
	mBase = m.M
	v2231 = m.ExcPending
	if v2231 != 0 {
		goto L13
	} else {
		goto L597
	}
L594:
	;
	v2226 = F_setExpire(m, int32(0), v404, v33+int32(360), v418)
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L13
	} else {
		goto L596
	}
L595:
	;
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(v33)+408))
	v2229 = v2222
	goto L593
L596:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+408)) = v2226
	v2229 = v2226
	goto L593
L597:
	;
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(v404)+28))
	F_moduleNotifyKeyspaceEvent(m, int32(4096), int32(_a1171), v33+int32(360), v2236)
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L13
	} else {
		goto L598
	}
L598:
	;
	v2240 = v414
	goto L538
L599:
	;
	v2244 = v2240
	v2245 = v415
	goto L537
L600:
	;
	F_debugDelay(m, v2249)
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L13
	} else {
		goto L601
	}
L601:
	;
	goto L116
L602:
	;
	if v2254 != int32(-1) {
		v2483 = v404
		v2489 = v410
		v2490 = v411
		v2491 = v412
		v2492 = v413
		v2493 = v414
		v2494 = v415
		v2495 = v416
		v2496 = v417
		v2497 = v418
		goto L115
	} else {
		goto L603
	}
L603:
	;
	goto L2
L604:
	;
	F_decrRefCount(m, v2260)
	mBase = m.M
	v2263 = m.ExcPending
	if v2263 != 0 {
		goto L13
	} else {
		goto L605
	}
L605:
	;
	goto L122
L606:
	;
	if v2271 != int32(2) {
		v2558 = v383
		goto L1
	} else {
		goto L607
	}
L607:
	;
	goto L2
L608:
	;
	goto L117
L609:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L610:
	;
	v2292 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v2301 = v33 + int32(360)
	v2302 = int32(8)
	goto L611
L611:
	;
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v2326) < base.Ui32(v2302) {
		goto L614
	} else {
		goto L615
	}
L612:
	;
	v2348 = *(*int32)(unsafe.Add(mBase, _consts[634]))
	if v2348 == int32(0) {
		goto L117
	} else {
		goto L626
	}
L613:
	;
	v2337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2337 == int32(0) {
		goto L622
	} else {
		goto L623
	}
L614:
	;
	v2328 = v2326
	goto L616
L615:
	;
	v2328 = v2302
	goto L616
L616:
	;
	if v2326 != 0 {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	v2329 = v2328
	goto L619
L618:
	;
	v2329 = v2302
	goto L619
L619:
	;
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2331 = m.T0[v2330].(func(*base.Module, int32, int32, int32) int32)(m, l0, v2301, v2329)
	mBase = m.M
	v2332 = m.ExcPending
	if v2332 != 0 {
		goto L13
	} else {
		goto L620
	}
L620:
	;
	if v2331 != 0 {
		goto L613
	} else {
		goto L621
	}
L621:
	;
	v2333 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v2333 | int64(1)
	goto L2
L622:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v2342 + v2329
	v2346 = v2302 - v2329
	if v2346 != 0 {
		v2301 = v2301 + v2329
		v2302 = v2346
		goto L611
	} else {
		goto L625
	}
L623:
	;
	m.T0[v2337].(func(*base.Module, int32, int32, int32))(m, l0, v2301, v2329)
	mBase = m.M
	v2341 = m.ExcPending
	if v2341 != 0 {
		goto L13
	} else {
		goto L624
	}
L624:
	;
	goto L622
L625:
	;
	goto L612
L626:
	;
	v2352 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	if v2352 != 0 {
		goto L117
	} else {
		goto L627
	}
L627:
	;
	v2353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v2353&int32(8) == int32(0) {
		goto L630
	} else {
		goto L631
	}
L628:
	;
	if v2363 == v2292 {
		goto L117
	} else {
		goto L636
	}
L629:
	;
	F__serverLog(m, int32(2), v2371, int32(0))
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L13
	} else {
		goto L635
	}
L630:
	;
	v2363 = *(*int64)(unsafe.Add(mBase, uint32(v33)+360))
	if v2363 != int64(0) {
		goto L628
	} else {
		goto L633
	}
L631:
	;
	v2359 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v2359 {
		goto L117
	} else {
		goto L632
	}
L632:
	;
	v2371 = int32(_a1172)
	goto L629
L633:
	;
	v2367 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v2367 {
		goto L117
	} else {
		goto L634
	}
L634:
	;
	v2371 = int32(_a1173)
	goto L629
L635:
	;
	goto L117
L636:
	;
	v2378 = int32(3)
	v2380 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v2378 < v2380 {
		goto L637
	} else {
		goto L638
	}
L637:
	;
	F_rdbReportError(m, int32(1), int32(3611), int32(_a1174), int32(0))
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L13
	} else {
		goto L640
	}
L638:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v33)+152)) = v2363
	*(*int64)(unsafe.Add(mBase, uint32(v33)+144)) = v2292
	F__serverLog(m, int32(3), int32(_a1175), v33+int32(144))
	mBase = m.M
	v2390 = m.ExcPending
	if v2390 != 0 {
		goto L13
	} else {
		goto L639
	}
L639:
	;
	goto L637
L640:
	;
	v2558 = v2378
	goto L1
L641:
	;
	v2452 = int32(0)
	if int32(2) < v2428 {
		v2558 = v2452
		goto L1
	} else {
		goto L645
	}
L642:
	;
	v2431 = int32(0)
	if int32(2) < v2428 {
		v2558 = v2431
		goto L1
	} else {
		goto L643
	}
L643:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(128)))) = v414
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(136)))) = v415
	v2440 = int32(_a44)
	v2441 = *(*int64)(unsafe.Add(mBase, _consts[640]))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+112)) = v2441
	v2444 = *(*int64)(unsafe.Add(mBase, _consts[637]))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+120)) = v2444
	F__serverLog(m, int32(2), int32(_a1176), v33+int32(112))
	mBase = m.M
	v2451 = m.ExcPending
	if v2451 != 0 {
		goto L13
	} else {
		goto L644
	}
L644:
	;
	v2558 = v2431
	goto L1
L645:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(96)))) = v415
	v2458 = int32(_a44)
	v2459 = *(*int64)(unsafe.Add(mBase, _consts[640]))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+80)) = v2459
	v2462 = *(*int64)(unsafe.Add(mBase, _consts[637]))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+88)) = v2462
	F__serverLog(m, int32(2), int32(_a1177), v33+int32(80))
	mBase = m.M
	v2469 = m.ExcPending
	if v2469 != 0 {
		goto L13
	} else {
		goto L646
	}
L646:
	;
	v2558 = v2452
	goto L1
L647:
	;
	goto L94
L648:
	;
	v2547 = int32(0)
	F_rdbReportError(m, v2547, int32(3632), int32(_a1178), v2547)
	mBase = m.M
	v2552 = m.ExcPending
	if v2552 != 0 {
		goto L13
	} else {
		goto L651
	}
L649:
	;
	F__serverLog(m, int32(3), int32(_a1179), int32(0))
	mBase = m.M
	v2546 = m.ExcPending
	if v2546 != 0 {
		goto L13
	} else {
		goto L650
	}
L650:
	;
	goto L648
L651:
	;
	v2558 = v2537
	goto L1
}
func F_rdbLoadStringObject(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_rdbGenericLoadStringObject(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_rdbPipeReadHandler(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
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
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v384 int32
	_ = v384
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _consts[670]))
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[649]))
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v22 = F_valkey_malloc(m, int32(16384))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[670])) = v22
	goto L1
L5:
	;
	m.G0 = v16 + int32(48)
	return
L6:
	;
	v251 = int32(_a44)
	v252 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v254 = *(*int32)(unsafe.Add(mBase, _consts[643]))
	F_aeDeleteFileEvent(m, v252, v254, int32(1))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L3
	} else {
		goto L52
	}
L7:
	;
	F__serverAssert(m, int32(_a1248), int32(_a1190), int32(1800))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L3
	} else {
		goto L51
	}
L8:
	;
	goto L9
L9:
	;
	v40 = int32(_a44)
	v42 = *(*int32)(unsafe.Add(mBase, _consts[670]))
	v44 = F_read(m, l1, v42, int32(16384))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[671])) = v44
	if int32(-1) < v44 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v238 = int32(_a44)
	v239 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v241 = *(*int32)(unsafe.Add(mBase, _consts[643]))
	F_aeDeleteFileEvent(m, v239, v241, int32(1))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L3
	} else {
		goto L50
	}
L11:
	;
	if v44 == int32(0) {
		goto L6
	} else {
		goto L28
	}
L12:
	;
	goto L13
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v49 == int32(6) {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v53 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[646]))
	if v63 < int32(1) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v56 = F___strerror_l(m, v49, v49)
	mBase = m.M
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v56
	F__serverLog(m, int32(3), int32(_a1249), v16)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	F_killRDBChild(m)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L27
	}
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[645]))
	v74 = int32(0)
	v75 = v63
	v76 = v67
	goto L21
L21:
	;
	v83 = v74 << (uint(int32(2)) % 32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v76+v83)))
	if v85 == int32(0) {
		v98 = v75
		v99 = v76
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L19
L23:
	;
	v101 = v74 + int32(1)
	if v101 < v98 {
		v74 = v101
		v75 = v98
		v76 = v99
		goto L21
	} else {
		goto L26
	}
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+24))
	v89 = F_freeClient(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v91 = int32(_a44)
	v92 = *(*int32)(unsafe.Add(mBase, _consts[645]))
	*(*int32)(unsafe.Add(mBase, uint32(v92+v83))) = int32(0)
	v97 = *(*int32)(unsafe.Add(mBase, _consts[646]))
	v98 = v97
	v99 = v92
	goto L23
L26:
	;
	goto L22
L27:
	;
	goto L5
L28:
	;
	v120 = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, _consts[646]))
	if v122 <= v120 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _consts[649]))
	if v235 == int32(0) {
		goto L9
	} else {
		goto L49
	}
L30:
	;
	v131 = v120
	goto L31
L31:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[645]))
	v141 = v131 << (uint(int32(2)) % 32)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139+v141)))
	if v143 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L29
L33:
	;
	v217 = v131 + int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, _consts[646]))
	if v217 < v219 {
		v131 = v217
		goto L31
	} else {
		goto L48
	}
L34:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+24))
	v147 = int32(_a44)
	v148 = *(*int32)(unsafe.Add(mBase, _consts[670]))
	v150 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+68))
	v153 = m.T0[v152].(func(*base.Module, int32, int32, int32) int32)(m, v143, v148, v150)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L37
	}
L35:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	if v153 == v195 {
		goto L33
	} else {
		goto L46
	}
L36:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v146)+104))
	v186 = base.I64_extend_i32_s(v153)
	*(*int64)(unsafe.Add(mBase, uint32(v185)+16)) = v186
	v188 = int32(_a44)
	v190 = *(*int64)(unsafe.Add(mBase, _consts[124]))
	*(*int64)(unsafe.Add(mBase, _consts[124])) = v190 + v186
	goto L35
L37:
	;
	if v153 != int32(-1) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if v157 == int32(3) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v146)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+16)) = int64(0)
	goto L35
L40:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v161 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v175 = F_freeClient(m, v146)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L45
	}
L42:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+88))
	v166 = m.T0[v165].(func(*base.Module, int32) int32)(m, v143)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v166
	F__serverLog(m, int32(3), int32(_a1250), v16+int32(32))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _consts[645]))
	*(*int32)(unsafe.Add(mBase, uint32(v178+v141))) = int32(0)
	goto L33
L46:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v146)+104))
	v198 = int32(_a44)
	v199 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	*(*int64)(unsafe.Add(mBase, uint32(v197)+88)) = v199
	v203 = *(*int32)(unsafe.Add(mBase, _consts[649]))
	*(*int32)(unsafe.Add(mBase, _consts[649])) = v203 + int32(1)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+80))
	v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v143, int32(971), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	goto L33
L48:
	;
	goto L32
L49:
	;
	goto L10
L50:
	;
	goto L5
L51:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _consts[646]))
	if v259 < int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v398 = int32(_a44)
	v399 = *(*int32)(unsafe.Add(mBase, _consts[644]))
	v400 = F_close(m, v399)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[644])) = int32(-1)
	goto L5
L54:
	;
	v263 = v259 & int32(3)
	v265 = *(*int32)(unsafe.Add(mBase, _consts[645]))
	v266 = int32(0)
	if base.Ui32(v259) < base.Ui32(int32(4)) {
		v324 = v266
		v325 = v266
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if v263 == int32(0) {
		v365 = v325
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v273 = int32(0)
	v277 = v273
	v281 = v273
	v282 = v273
	goto L57
L57:
	;
	v291 = v265 + v281<<(uint(int32(2))%32)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v293 = int32(0)
	v296 = int32(4)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v291+v296)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v291+int32(8))))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v291+int32(12))))
	v313 = v282 + base.B2i32(v292 != v293) + base.B2i32(v298 != v293) + base.B2i32(v304 != v293) + base.B2i32(v310 != v293)
	v315 = v281 + v296
	v317 = v277 + v296
	if v317 != v259&int32(2147483644) {
		v277 = v317
		v281 = v315
		v282 = v313
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v324 = v315
	v325 = v313
	goto L55
L59:
	;
	goto L58
L60:
	;
	if v365 == int32(0) {
		goto L53
	} else {
		goto L65
	}
L61:
	;
	v339 = v324
	v340 = v325
	v342 = v266
	goto L62
L62:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v265+v339<<(uint(int32(2))%32))))
	v353 = v340 + base.B2i32(v350 != int32(0))
	v354 = int32(1)
	v357 = v342 + v354
	if v357 != v263 {
		v339 = v339 + v354
		v340 = v353
		v342 = v357
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v365 = v353
	goto L60
L64:
	;
	goto L63
L65:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v375 {
		goto L53
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v365
	F__serverLog(m, int32(2), int32(_a1251), v16+int32(16))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	goto L53
}
func F_rdbPipeWriteHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	if v13 <= int32(0) {
		F__serverAssert(m, int32(_a1252), int32(_a1190), int32(1773))
		mBase = m.M
		v105 = m.ExcPending
		if v105 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[670]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
		v25 = m.T0[v24].(func(*base.Module, int32, int32, int32) int32)(m, l0, v17+v20, v13-v20)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			if v25 != int32(-1) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
				v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)+16))
				v49 = base.I64_extend_i32_s(v25)
				v50 = v48 + v49
				*(*int64)(unsafe.Add(mBase, uint32(v47)+16)) = v50
				v52 = int32(_a44)
				v54 = *(*int64)(unsafe.Add(mBase, _consts[124]))
				*(*int64)(unsafe.Add(mBase, _consts[124])) = v54 + v49
				v58 = int64(*(*int32)(unsafe.Add(mBase, _consts[671])))
				if v58 <= v50 {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					if v64 == int32(0) {
						m.G0 = v10 + int32(16)
						return
					} else {
						v67 = int32(0)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+80))
						v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int32)(m, l0, v67, v67)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+104))
							*(*int64)(unsafe.Add(mBase, uint32(v74)+88)) = int64(0)
							v77 = int32(_a44)
							v79 = *(*int32)(unsafe.Add(mBase, _consts[649]))
							v81 = v79 + int32(-1)
							*(*int32)(unsafe.Add(mBase, _consts[649])) = v81
							if v81 != 0 {
								m.G0 = v10 + int32(16)
								return
							} else {
								v83 = int32(_a44)
								v84 = *(*int32)(unsafe.Add(mBase, _consts[279]))
								v86 = *(*int32)(unsafe.Add(mBase, _consts[643]))
								v90 = F_aeCreateFileEvent(m, v84, v86, int32(1), int32(969), int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									if v90 == int32(-1) {
										F__serverPanic_1(m, int32(_a1190), int32(1765), int32(_a1180), int32(0))
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										m.G0 = v10 + int32(16)
										return
									}
								}
							}
						}
					}
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
					v62 = *(*int64)(unsafe.Add(mBase, _consts[47]))
					*(*int64)(unsafe.Add(mBase, uint32(v60)+88)) = v62
					m.G0 = v10 + int32(16)
					return
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v29 == int32(3) {
					m.G0 = v10 + int32(16)
					return
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					if int32(3) < v33 {
						v45 = F_freeClient(m, v18)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v10 + int32(16)
							return
						}
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
						v38 = m.T0[v37].(func(*base.Module, int32) int32)(m, l0)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v38
							F__serverLog(m, int32(3), int32(_a1253), v10)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v45 = F_freeClient(m, v18)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
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
func F_rdbRegisterAuxField(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	v6 = *(*int32)(unsafe.Add(mBase, _consts[632]))
	if v6 != 0 {
		v15 = F_valkey_malloc(m, int32(8))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
			v20 = int32(0)
			v22 = *(*int32)(unsafe.Add(mBase, _consts[632]))
			v23 = F_sdsnew(m, l0)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = F_dictAdd(m, v22, v23, v15)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v25 != 0 {
						v27 = int32(-1)
					} else {
						v27 = v20
					}
					return v27
				}
			}
		}
	} else {
		v9 = F_dictCreate(m, int32(_a1120))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[632])) = v9
			v15 = F_valkey_malloc(m, int32(8))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
				v20 = int32(0)
				v22 = *(*int32)(unsafe.Add(mBase, _consts[632]))
				v23 = F_sdsnew(m, l0)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = F_dictAdd(m, v22, v23, v15)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 != 0 {
							v27 = int32(-1)
						} else {
							v27 = v20
						}
						return v27
					}
				}
			}
		}
	}
}
func F_rdbSaveLongLongAsStringObject(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v46 int64
	_ = v46
	var v49 int64
	_ = v49
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v199 int32
	_ = v199
	var v212 int32
	_ = v212
	v2 = l1
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if base.Ui64(int64(255)) < base.Ui64(v2+int64(128)) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F__serverAssert(m, int32(_a1121), int32(_a1122), int32(537))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L23
	} else {
		goto L62
	}
L2:
	;
	m.G0 = v13 + int32(32)
	return v199
L3:
	;
	if v2 <= int64(-1) {
		goto L33
	} else {
		goto L34
	}
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v54)
	if l0 != 0 {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	if base.Ui64(int64(65535)) < base.Ui64(v2+int64(32768)) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v22 = int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v22)
	v54 = v2
	v55 = int32(2)
	v56 = v13 | int32(1)
	goto L4
L7:
	;
	if base.Ui64(int64(4294967295)) < base.Ui64(v2+int64(2147483648)) {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v30 = int32(193)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v30)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)) = uint8(v2)
	v54 = int64(base.Ui64(v2) >> (uint(int64(8)) % 64))
	v55 = int32(3)
	v56 = v13 | int32(2)
	goto L4
L9:
	;
	v42 = int32(194)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v42)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)) = uint8(v2)
	v46 = int64(base.Ui64(v2) >> (uint(int64(16)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)) = uint8(v46)
	v49 = int64(base.Ui64(v2) >> (uint(int64(8)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)) = uint8(v49)
	v54 = int64(base.Ui64(v2) >> (uint(int64(24)) % 64))
	v55 = int32(5)
	v56 = v13 | int32(4)
	goto L4
L10:
	;
	v58 = int32(-1)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v59&int32(6) != 0 {
		v199 = v58
		goto L2
	} else {
		goto L12
	}
L11:
	;
	v199 = v55
	goto L2
L12:
	;
	v68 = v13
	v69 = v55
	goto L13
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v72) < base.Ui32(v69) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v199 = v55
	goto L2
L15:
	;
	v74 = v72
	goto L17
L16:
	;
	v74 = v69
	goto L17
L17:
	;
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v75 = v74
	goto L20
L19:
	;
	v75 = v69
	goto L20
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v76 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v84 = m.T0[v83].(func(*base.Module, int32, int32, int32) int32)(m, l0, v68, v75)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L23
	} else {
		goto L26
	}
L22:
	;
	m.T0[v76].(func(*base.Module, int32, int32, int32))(m, l0, v68, v75)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	goto L21
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v90 + v75
	v94 = v69 - v75
	if v94 != 0 {
		v68 = v68 + v75
		v69 = v94
		goto L13
	} else {
		goto L28
	}
L26:
	;
	if v84 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v86 | int64(2)
	v199 = v58
	goto L2
L28:
	;
	goto L14
L29:
	;
	if int32(32) <= v136 {
		goto L1
	} else {
		goto L38
	}
L30:
	;
	v136 = int32(0)
	goto L29
L32:
	;
	v117 = F_ull2string(m, v113, v114, v115)
	mBase = m.M
	if v117 == int32(0) {
		goto L30
	} else {
		goto L36
	}
L33:
	;
	goto L35
L34:
	;
	v113 = v13
	v114 = int32(32)
	v115 = v2
	v116 = int32(0)
	goto L32
L35:
	;
	v104 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v104)
	v108 = int32(1)
	v113 = v13 + v108
	v114 = int32(31)
	v115 = int64(0) - v2
	v116 = v108
	goto L32
L36:
	;
	v136 = v117 + v116
	goto L29
L38:
	;
	v139 = int32(-1)
	v141 = F_rdbSaveLen(m, l0, base.I64_extend_i32_s(v136))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L23
	} else {
		goto L39
	}
L39:
	;
	if v141 == int32(-1) {
		v199 = v139
		goto L2
	} else {
		goto L40
	}
L40:
	;
	if l0 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if v136 == int32(-1) {
		v199 = v139
		goto L2
	} else {
		goto L61
	}
L42:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v147&int32(6) != 0 {
		v199 = v139
		goto L2
	} else {
		goto L43
	}
L43:
	;
	if v136 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v156 = v13
	v157 = v136
	goto L46
L45:
	;
	v199 = v141
	goto L2
L46:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v160) < base.Ui32(v157) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L41
L48:
	;
	v162 = v160
	goto L50
L49:
	;
	v162 = v157
	goto L50
L50:
	;
	if v160 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v163 = v162
	goto L53
L52:
	;
	v163 = v157
	goto L53
L53:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v164 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v170 = m.T0[v169].(func(*base.Module, int32, int32, int32) int32)(m, l0, v156, v163)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L23
	} else {
		goto L58
	}
L55:
	;
	m.T0[v164].(func(*base.Module, int32, int32, int32))(m, l0, v156, v163)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L23
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v176 + v163
	v180 = v157 - v163
	if v180 != 0 {
		v156 = v156 + v163
		v157 = v180
		goto L46
	} else {
		goto L60
	}
L58:
	;
	if v170 != 0 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v172 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v172 | int64(2)
	v199 = v139
	goto L2
L60:
	;
	goto L47
L61:
	;
	v199 = v141 + v136
	goto L2
L62:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_rdbSaveLzfBlob(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
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
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v151 int32
	_ = v151
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(195)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v15)
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v151
L2:
	;
	v71 = F_rdbSaveLen(m, l0, base.I64_extend_i32_u(l2))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L15
	} else {
		goto L21
	}
L3:
	;
	v19 = int32(-1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v20&int32(6) != 0 {
		v151 = v19
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = v13 + int32(15)
	v33 = int32(1)
	goto L5
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v36) < base.Ui32(v33) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	v38 = v36
	goto L9
L8:
	;
	v38 = v33
	goto L9
L9:
	;
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v39 = v38
	goto L12
L11:
	;
	v39 = v33
	goto L12
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v40 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v48 = m.T0[v47].(func(*base.Module, int32, int32, int32) int32)(m, l0, v32, v39)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L15
	} else {
		goto L18
	}
L14:
	;
	m.T0[v40].(func(*base.Module, int32, int32, int32))(m, l0, v32, v39)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	goto L13
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v54 + v39
	v58 = v33 - v39
	if v58 != 0 {
		v32 = v32 + v39
		v33 = v58
		goto L5
	} else {
		goto L20
	}
L18:
	;
	if v48 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v50 | int64(2)
	v151 = v19
	goto L1
L20:
	;
	goto L6
L21:
	;
	if v71 == int32(-1) {
		v151 = int32(-1)
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v75 = int32(-1)
	v77 = F_rdbSaveLen(m, l0, base.I64_extend_i32_u(l3))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	if v77 == int32(-1) {
		v151 = v75
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if l0 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v151 = l2 + v71 + v77 + int32(1)
	goto L1
L26:
	;
	v129 = int32(-1)
	if l2 == v129 {
		v151 = v129
		goto L1
	} else {
		goto L45
	}
L27:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v83&int32(6) != 0 {
		v151 = v75
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if l2 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v89 = l1
	v95 = l2
	goto L30
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v98) < base.Ui32(v95) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L26
L32:
	;
	v100 = v98
	goto L34
L33:
	;
	v100 = v95
	goto L34
L34:
	;
	if v98 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v101 = v100
	goto L37
L36:
	;
	v101 = v95
	goto L37
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v102 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v108 = m.T0[v107].(func(*base.Module, int32, int32, int32) int32)(m, l0, v89, v101)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L15
	} else {
		goto L42
	}
L39:
	;
	m.T0[v102].(func(*base.Module, int32, int32, int32))(m, l0, v89, v101)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v114 + v101
	v118 = v95 - v101
	if v118 != 0 {
		v89 = v89 + v101
		v95 = v118
		goto L30
	} else {
		goto L44
	}
L42:
	;
	if v108 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v110 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v110 | int64(2)
	v151 = v75
	goto L1
L44:
	;
	goto L31
L45:
	;
	goto L25
}
func F_rdbSaveModulesAux(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = *(*int32)(unsafe.Add(mBase, _consts[541]))
	v14 = F_dictGetIterator(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v320
L2:
	;
	F_dictReleaseIterator(m, v14)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L3
	} else {
		goto L83
	}
L3:
	;
	return int32(0)
L4:
	;
	v25 = v14 + int32(20)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v121 == int32(0) {
		v312 = v3
		goto L2
	} else {
		goto L31
	}
L6:
	;
	v32 = v25
	v33 = v29
	goto L9
L7:
	;
	v29 = int32(1)
	goto L6
L8:
	;
	v29 = int32(0)
	goto L6
L9:
	;
	switch v33 {
	case 0:
		goto L14
	default:
		goto L13
	}
L11:
	;
	v33 = int32(0)
	goto L9
L12:
	;
	goto L5
L13:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v113
	if v113 == int32(0) {
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v37 != int32(-1) {
		v76 = v37
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v77 = int32(1)
	v78 = v76 + v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v78
	v80 = int32(0)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v84+int32(26)))))
	if v88 == int32(255) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v41 != 0 {
		v76 = int32(-1)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v43 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	if v70 != int32(-1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v50 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v42)+16)))
	v51 = int64(*(*int8)(unsafe.Add(mBase, uint32(v42)+27)))
	v52 = int64(*(*int32)(unsafe.Add(mBase, uint32(v42)+8)))
	v53 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v42)+12)))
	v54 = int64(*(*int8)(unsafe.Add(mBase, uint32(v42)+26)))
	v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v42)+4)))
	v56 = F_wangHash64(m, v55)
	mBase = m.M
	v58 = F_wangHash64(m, v54+v56)
	mBase = m.M
	v60 = F_wangHash64(m, v53+v58)
	mBase = m.M
	v62 = F_wangHash64(m, v52+v60)
	mBase = m.M
	v64 = F_wangHash64(m, v51+v62)
	mBase = m.M
	v66 = F_wangHash64(m, v50+v64)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v69 = v68
	goto L18
L20:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+24)))
	v48 = v46 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+24)) = uint16(v48)
	v69 = v42
	goto L18
L21:
	;
	v76 = v70 + int32(-1)
	goto L15
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v76 = v73
	goto L15
L23:
	;
	v103 = int32(2)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v83+v101<<(uint(v103)%32)+int32(4))))
	v32 = v108 + v102<<(uint(v103)%32)
	v33 = int32(1)
	goto L9
L24:
	;
	v92 = v80
	goto L26
L25:
	;
	v92 = v77 << (uint(v88) % 32)
	goto L26
L26:
	;
	if v78 < v92 {
		v101 = v84
		v102 = v78
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if v84 != 0 {
		v121 = v80
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	if v94 == int32(-1) {
		v121 = v80
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(4294967296)
	v101 = int32(1)
	v102 = int32(0)
	goto L23
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v117
	v121 = v113
	goto L12
L31:
	;
	v130 = v3
	v132 = v121
	goto L32
L32:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	goto L34
L33:
	;
	v312 = v199
	goto L2
L34:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	v136 = v9 + int32(8)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v137
	goto L35
L35:
	;
	v142 = v9 + int32(8)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v144 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v209 = v14 + int32(20)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v210 != 0 {
		goto L58
	} else {
		goto L59
	}
L37:
	;
	if v144 == int32(0) {
		v199 = v130
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v144+base.B2i32(v147 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v153
	goto L38
L40:
	;
	v160 = v130
	v162 = v144
	goto L41
L41:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+56))
	if v164 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v199 = v180
	goto L36
L43:
	;
	v183 = v9 + int32(8)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v185 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L44:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)+80))
	if v168&l1 == int32(0) {
		v180 = v160
		goto L43
	} else {
		goto L47
	}
L45:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v163)+76))
	if v165 == int32(0) {
		v180 = v160
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v172 = F_rdbSaveSingleModuleAux(m, l0, l1, v163)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L3
	} else {
		goto L49
	}
L48:
	;
	v180 = v172 + v160
	goto L43
L49:
	;
	if v172 != int32(-1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	F_dictReleaseIterator(m, v14)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v320 = int32(-1)
	goto L1
L52:
	;
	if v185 != 0 {
		v160 = v180
		v162 = v185
		goto L41
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v185+base.B2i32(v188 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v194
	goto L53
L55:
	;
	goto L42
L56:
	;
	if v305 != 0 {
		v130 = v199
		v132 = v305
		goto L32
	} else {
		goto L82
	}
L57:
	;
	v216 = v209
	v217 = v213
	goto L60
L58:
	;
	v213 = int32(1)
	goto L57
L59:
	;
	v213 = int32(0)
	goto L57
L60:
	;
	switch v217 {
	case 0:
		goto L65
	default:
		goto L64
	}
L62:
	;
	v217 = int32(0)
	goto L60
L63:
	;
	goto L56
L64:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v297
	if v297 == int32(0) {
		goto L62
	} else {
		goto L81
	}
L65:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v221 != int32(-1) {
		v260 = v221
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v261 = int32(1)
	v262 = v260 + v261
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v262
	v264 = int32(0)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267+v268+int32(26)))))
	if v272 == int32(255) {
		goto L75
	} else {
		goto L76
	}
L67:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v225 != 0 {
		v260 = int32(-1)
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v227 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+20))
	if v254 != int32(-1) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v234 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v226)+16)))
	v235 = int64(*(*int8)(unsafe.Add(mBase, uint32(v226)+27)))
	v236 = int64(*(*int32)(unsafe.Add(mBase, uint32(v226)+8)))
	v237 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v226)+12)))
	v238 = int64(*(*int8)(unsafe.Add(mBase, uint32(v226)+26)))
	v239 = int64(*(*int32)(unsafe.Add(mBase, uint32(v226)+4)))
	v240 = F_wangHash64(m, v239)
	mBase = m.M
	v242 = F_wangHash64(m, v238+v240)
	mBase = m.M
	v244 = F_wangHash64(m, v237+v242)
	mBase = m.M
	v246 = F_wangHash64(m, v236+v244)
	mBase = m.M
	v248 = F_wangHash64(m, v235+v246)
	mBase = m.M
	v250 = F_wangHash64(m, v234+v248)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v250
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v253 = v252
	goto L69
L71:
	;
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226)+24)))
	v232 = v230 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v226)+24)) = uint16(v232)
	v253 = v226
	goto L69
L72:
	;
	v260 = v254 + int32(-1)
	goto L66
L73:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v260 = v257
	goto L66
L74:
	;
	v287 = int32(2)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v267+v285<<(uint(v287)%32)+int32(4))))
	v216 = v292 + v286<<(uint(v287)%32)
	v217 = int32(1)
	goto L60
L75:
	;
	v276 = v264
	goto L77
L76:
	;
	v276 = v261 << (uint(v272) % 32)
	goto L77
L77:
	;
	if v262 < v276 {
		v285 = v268
		v286 = v262
		goto L74
	} else {
		goto L78
	}
L78:
	;
	if v268 != 0 {
		v305 = v264
		goto L63
	} else {
		goto L79
	}
L79:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v267)+20))
	if v278 == int32(-1) {
		v305 = v264
		goto L63
	} else {
		goto L80
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(4294967296)
	v285 = int32(1)
	v286 = int32(0)
	goto L74
L81:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v301
	v305 = v297
	goto L63
L82:
	;
	goto L33
L83:
	;
	v320 = v312
	goto L1
}
func F_rdbSaveRawString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v67 int64
	_ = v67
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v89 int64
	_ = v89
	var v94 int64
	_ = v94
	var v98 int32
	_ = v98
	var v100 int64
	_ = v100
	var v102 int32
	_ = v102
	var v110 int64
	_ = v110
	var v134 int64
	_ = v134
	var v153 int32
	_ = v153
	var v162 int64
	_ = v162
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v190 int32
	_ = v190
	var v194 int64
	_ = v194
	var v197 int64
	_ = v197
	var v202 int32
	_ = v202
	var v203 int64
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int64
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v334 int32
	_ = v334
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if base.Ui32(int32(11)) < base.Ui32(l2) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v334
L2:
	;
	v273 = int32(-1)
	v275 = F_rdbSaveLen(m, l0, base.I64_extend_i32_u(l2))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L52
	} else {
		goto L63
	}
L3:
	;
	if base.Ui32(l2) < base.Ui32(int32(21)) {
		goto L2
	} else {
		goto L58
	}
L4:
	;
	v18 = v13 + int32(8)
	v19 = int32(0)
	if base.Ui32(l2+int32(-21)) < base.Ui32(int32(-20)) {
		v153 = v19
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v153 == int32(0) {
		goto L2
	} else {
		goto L32
	}
L6:
	;
	goto L5
L7:
	;
	v31 = int32(1)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if l2 != v31 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v153 = int32(1)
	goto L6
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v134
	goto L8
L10:
	;
	if v32&int32(255) == int32(45) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v36 = v32 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v36&int32(255)) {
		v153 = v19
		goto L6
	} else {
		goto L12
	}
L12:
	;
	if v18 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v134 = base.I64_extend_i32_u(v36) & int64(255)
	goto L9
L14:
	;
	if base.Ui32(int32(8)) < base.Ui32((v55+int32(-49))&int32(255)) {
		v153 = v19
		goto L6
	} else {
		goto L17
	}
L15:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v54 = int32(2)
	v55 = v52
	v56 = l1 + int32(1)
	goto L14
L16:
	;
	v54 = v31
	v55 = v32
	v56 = l1
	goto L14
L17:
	;
	v67 = base.I64_extend_i32_u(v55+int32(-48)) & int64(255)
	if base.Ui32(l2) <= base.Ui32(v54) {
		v110 = v67
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v32&int32(255) != int32(45) {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v73 = v54
	v75 = v67
	v77 = v56
	goto L20
L20:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if base.Ui32((v79+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v153 = v19
		goto L6
	} else {
		goto L22
	}
L21:
	;
	v110 = v100
	goto L18
L22:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v75) {
		v153 = v19
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v89 = v75 * int64(10)
	v94 = base.I64_extend_i32_u(v79+int32(-48)) & int64(255)
	if base.Ui64(v94^int64(-1)) < base.Ui64(v89) {
		v153 = v19
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v98 = int32(1)
	v100 = v89 + v94
	v102 = v73 + v98
	if v102 != l2 {
		v73 = v102
		v75 = v100
		v77 = v77 + v98
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	if v110 < int64(0) {
		v153 = v19
		goto L6
	} else {
		goto L30
	}
L27:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v110) {
		v153 = v19
		goto L6
	} else {
		goto L28
	}
L28:
	;
	if v18 == int32(0) {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	v134 = int64(0) - v110
	goto L9
L30:
	;
	if v18 == int32(0) {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	v134 = v110
	goto L9
L32:
	;
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	if base.Ui64(int64(255)) < base.Ui64(v162+int64(128)) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v202))) = uint8(v203)
	if l0 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	if base.Ui64(int64(65535)) < base.Ui64(v162+int64(32768)) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v169 = int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)) = uint8(v169)
	v202 = v13 + int32(4)
	v203 = v162
	v204 = int32(2)
	goto L33
L36:
	;
	if base.Ui64(int64(4294967295)) < base.Ui64(v162+int64(2147483648)) {
		goto L2
	} else {
		goto L38
	}
L37:
	;
	v178 = int32(193)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)) = uint8(v178)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v162)
	v202 = v13 + int32(5)
	v203 = int64(base.Ui64(v162) >> (uint(int64(8)) % 64))
	v204 = int32(3)
	goto L33
L38:
	;
	v190 = int32(194)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)) = uint8(v190)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v162)
	v194 = int64(base.Ui64(v162) >> (uint(int64(16)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+6)) = uint8(v194)
	v197 = int64(base.Ui64(v162) >> (uint(int64(8)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)) = uint8(v197)
	v202 = v13 + int32(7)
	v203 = int64(base.Ui64(v162) >> (uint(int64(24)) % 64))
	v204 = int32(5)
	goto L33
L39:
	;
	v334 = v204
	goto L1
L40:
	;
	v208 = int32(-1)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v209&int32(6) != 0 {
		v334 = v208
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v216 = v13 + int32(3)
	v221 = v204
	goto L42
L42:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v224) < base.Ui32(v221) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L39
L44:
	;
	v226 = v224
	goto L46
L45:
	;
	v226 = v221
	goto L46
L46:
	;
	if v224 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v227 = v226
	goto L49
L48:
	;
	v227 = v221
	goto L49
L49:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v228 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v236 = m.T0[v235].(func(*base.Module, int32, int32, int32) int32)(m, l0, v216, v227)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L52
	} else {
		goto L55
	}
L51:
	;
	m.T0[v228].(func(*base.Module, int32, int32, int32))(m, l0, v216, v227)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	return int32(0)
L53:
	;
	goto L50
L54:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v242 + v227
	v246 = v221 - v227
	if v246 != 0 {
		v216 = v216 + v227
		v221 = v246
		goto L42
	} else {
		goto L57
	}
L55:
	;
	if v236 != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v238 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v238 | int64(2)
	v334 = v208
	goto L1
L57:
	;
	goto L43
L58:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _consts[633]))
	if v260 == int32(0) {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v264 = F_rdbSaveLzfStringObject(m, l0, l1, l2)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	if v264 == int32(-1) {
		v334 = int32(-1)
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if int32(0) < v264 {
		v334 = v264
		goto L1
	} else {
		goto L62
	}
L62:
	;
	goto L2
L63:
	;
	if v275 == int32(-1) {
		v334 = v273
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if l2 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if l0 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v334 = v275
	goto L1
L67:
	;
	if l2 == int32(-1) {
		v334 = v273
		goto L1
	} else {
		goto L85
	}
L68:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v281&int32(6) != 0 {
		v334 = v273
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v285 = l1
	v292 = l2
	goto L70
L70:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v294) < base.Ui32(v292) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L67
L72:
	;
	v296 = v294
	goto L74
L73:
	;
	v296 = v292
	goto L74
L74:
	;
	if v294 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v297 = v296
	goto L77
L76:
	;
	v297 = v292
	goto L77
L77:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v298 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v304 = m.T0[v303].(func(*base.Module, int32, int32, int32) int32)(m, l0, v285, v297)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L52
	} else {
		goto L82
	}
L79:
	;
	m.T0[v298].(func(*base.Module, int32, int32, int32))(m, l0, v285, v297)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L52
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v310 + v297
	v314 = v292 - v297
	if v314 != 0 {
		v285 = v285 + v297
		v292 = v314
		goto L70
	} else {
		goto L84
	}
L82:
	;
	if v304 != 0 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v306 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v306 | int64(2)
	v334 = v273
	goto L1
L84:
	;
	goto L71
L85:
	;
	v334 = v275 + l2
	goto L1
}
func F_rdbSaveRio(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int64
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v220 int64
	_ = v220
	var v222 int32
	_ = v222
	var v235 int32
	_ = v235
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
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int64
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v293 int32
	_ = v293
	v7 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v7
	v19 = *(*int32)(unsafe.Add(mBase, _consts[634]))
	if v19 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if base.Ui32(int32(81)) <= base.Ui32(l1) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = int32(962)
	goto L1
L3:
	;
	m.G0 = v14 + int32(48)
	return v293
L4:
	;
	v278 = int32(-1)
	if l3 == int32(0) {
		v293 = v278
		goto L3
	} else {
		goto L88
	}
L5:
	;
	F__serverAssert(m, int32(_a1128), int32(_a1122), int32(1484))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L10
	} else {
		goto L87
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
	if l1 == int32(80) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = int32(_a141)
	goto L9
L8:
	;
	v31 = int32(_a140)
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v31
	v37 = F_snprintf(m, v14+int32(38), int32(10), int32(_a1129), v14)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if l2 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v92 = F_rdbSaveInfoAuxFields(m, l2, l4, l5)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L30
	}
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
	if v43&int32(6) != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v56 = v14 + int32(38)
	v57 = int32(9)
	goto L15
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if base.Ui32(v60) < base.Ui32(v57) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L12
L17:
	;
	v62 = v60
	goto L19
L18:
	;
	v62 = v57
	goto L19
L19:
	;
	if v60 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v63 = v62
	goto L22
L21:
	;
	v63 = v57
	goto L22
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v64 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v70 = m.T0[v69].(func(*base.Module, int32, int32, int32) int32)(m, l2, v56, v63)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L27
	}
L24:
	;
	m.T0[v64].(func(*base.Module, int32, int32, int32))(m, l2, v56, v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v76 + v63
	v80 = v57 - v63
	if v80 != 0 {
		v56 = v56 + v63
		v57 = v80
		goto L15
	} else {
		goto L29
	}
L27:
	;
	if v70 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+32)) = v72 | int64(2)
	goto L4
L29:
	;
	goto L16
L30:
	;
	if v92 == int32(-1) {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v97 = l0 & int32(1)
	if v97 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if l0&int32(2) != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v99 = F_rdbSaveModulesAux(m, l2, int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	if v99 == int32(-1) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	if v97 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v105 = F_rdbSaveFunctions(m, l2)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	if v105 == int32(-1) {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	v167 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)) = uint8(v167)
	if l2 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L41:
	;
	v109 = F_clusterRDBSaveSlotImports(m, l2, l1)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	if v109 == int32(-1) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v113 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v115 <= v113 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v152 = F_rdbSaveModulesAux(m, l2, int32(2))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L10
	} else {
		goto L51
	}
L45:
	;
	v127 = v113
	goto L46
L46:
	;
	v131 = F_rdbSaveDb(m, l2, v127, l4, l1, v14+int32(20))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L10
	} else {
		goto L48
	}
L47:
	;
	goto L44
L48:
	;
	if v131 == int32(-1) {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v136 = v127 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v136 < v138 {
		v127 = v136
		goto L46
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	if v152 == int32(-1) {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	goto L40
L53:
	;
	v220 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v220
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
	if v222&int32(6) != 0 {
		goto L4
	} else {
		goto L71
	}
L54:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
	if v171&int32(6) != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v184 = v14 + int32(24)
	v185 = int32(1)
	goto L56
L56:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if base.Ui32(v188) < base.Ui32(v185) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L53
L58:
	;
	v190 = v188
	goto L60
L59:
	;
	v190 = v185
	goto L60
L60:
	;
	if v188 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v191 = v190
	goto L63
L62:
	;
	v191 = v185
	goto L63
L63:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v192 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v198 = m.T0[v197].(func(*base.Module, int32, int32, int32) int32)(m, l2, v184, v191)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L10
	} else {
		goto L68
	}
L65:
	;
	m.T0[v192].(func(*base.Module, int32, int32, int32))(m, l2, v184, v191)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v204 + v191
	v208 = v185 - v191
	if v208 != 0 {
		v184 = v184 + v191
		v185 = v208
		goto L56
	} else {
		goto L70
	}
L68:
	;
	if v198 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v200 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+32)) = v200 | int64(2)
	goto L4
L70:
	;
	goto L57
L71:
	;
	v235 = int32(8)
	v238 = v14 + int32(24)
	goto L72
L72:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if base.Ui32(v239) < base.Ui32(v235) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v241 = v239
	goto L76
L75:
	;
	v241 = v235
	goto L76
L76:
	;
	if v239 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v242 = v241
	goto L79
L78:
	;
	v242 = v235
	goto L79
L79:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v243 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v249 = m.T0[v248].(func(*base.Module, int32, int32, int32) int32)(m, l2, v238, v242)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L10
	} else {
		goto L84
	}
L81:
	;
	m.T0[v243].(func(*base.Module, int32, int32, int32))(m, l2, v238, v242)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v255 + v242
	v260 = v235 - v242
	if v260 != 0 {
		v235 = v260
		v238 = v238 + v242
		goto L72
	} else {
		goto L86
	}
L84:
	;
	if v249 != 0 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v251 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+32)) = v251 | int64(2)
	goto L4
L86:
	;
	v293 = int32(0)
	goto L3
L87:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	goto L89
L89:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v282
	v293 = v278
	goto L3
}
func F_rdbSaveStringObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch int32(base.Ui32(v6)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		v25 = F_objectGetVal(m, l1)
		mBase = m.M
		v27 = F_objectGetVal(m, l1)
		mBase = m.M
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-1)))))
		switch v30 & int32(7) {
		case 0:
			v35 = F_rdbSaveRawString(m, l0, v25, int32(base.Ui32(v30)>>(uint(int32(3))%32)))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				return v35
			}
		case 1:
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-3)))))
			v41 = F_rdbSaveRawString(m, l0, v25, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				return v41
			}
		case 2:
			v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(-5)))))
			v47 = F_rdbSaveRawString(m, l0, v25, v46)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				return v47
			}
		case 3:
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-9))))
			v53 = F_rdbSaveRawString(m, l0, v25, v52)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				return v53
			}
		case 4:
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-17))))
			v59 = v58
			v60 = F_rdbSaveRawString(m, l0, v25, v59)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				return v60
			}
		default:
			v59 = int32(0)
			v60 = F_rdbSaveRawString(m, l0, v25, v59)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				return v60
			}
		}
	case 1:
		v11 = F_objectGetVal(m, l1)
		mBase = m.M
		v13 = F_rdbSaveLongLongAsStringObject(m, l0, base.I64_extend_i32_s(v11))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v13
		}
	default:
		F__serverAssertWithInfo(m, int32(0), l1, int32(_a1123), int32(_a1122), int32(553))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
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
func F_rdbSaveToFile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	v4 = F___syscall_getpid(m)
	mBase = m.M
	v5 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[635]))
	F_moduleFireServerEvent(m, int64(1), base.B2i32(v4 == v6)<<(uint(int32(1))%32), v5)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(0)
		v18 = F_rdbSaveInternal(m, v15, l0, v15, v15)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v18 == int32(0) {
				F_moduleFireServerEvent(m, int64(1), int32(3), int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			} else {
				v22 = int32(9116376)
				v23 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				F_moduleFireServerEvent(m, int64(1), int32(4), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[5])) = v23
					return int32(-1)
				}
			}
		}
	}
}
func F_rdbSaveToReplicasSockets(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v318 int64
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v423 int64
	_ = v423
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int64
	_ = v622
	var v623 int32
	_ = v623
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v657 int32
	_ = v657
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v697 int32
	_ = v697
	var v704 int32
	_ = v704
	v14 = m.G0
	v16 = v14 - int32(128)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	goto L5
L1:
	;
	F__serverPanic_1(m, int32(_a1122), int32(3935), int32(_a1180), int32(0))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L46
	} else {
		goto L187
	}
L2:
	;
	F__serverAssert(m, int32(_a1181), int32(_a1122), int32(3776))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L46
	} else {
		goto L186
	}
L3:
	;
	m.G0 = v16 + int32(128)
	return v680
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[643]))
	if v26 != int32(-1) {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	if base.B2i32(v19 != int32(-1)) == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v680 = int32(-1)
	goto L3
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[644]))
	if v30 != int32(-1) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[645]))
	if v34 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v38 = int32(-1)
	v41 = l0 & int32(4)
	if v41 != 0 {
		v205 = v38
		v206 = v38
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v680 = int32(-1)
	goto L3
L11:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
	v213 = F_valkey_malloc(m, v210<<(uint(int32(2))%32))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L46
	} else {
		goto L47
	}
L12:
	;
	v42 = int32(-1)
	v44 = v16 + int32(112)
	v49 = m.G0
	v51 = v49 - int32(64)
	m.G0 = v51
	v54 = F_pipe(m, v44)
	mBase = m.M
	if v54 != 0 {
		v111 = v42
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v111 == int32(-1) {
		v680 = v42
		goto L3
	} else {
		goto L28
	}
L14:
	;
	m.G0 = v51 + int32(64)
	goto L13
L15:
	;
	goto L17
L16:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v105 = F_close(m, v104)
	mBase = m.M
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v107 = F_close(m, v106)
	mBase = m.M
	v111 = int32(-1)
	goto L14
L17:
	;
	goto L20
L20:
	;
	goto L24
L23:
	;
	v111 = int32(0)
	goto L14
L24:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = int32(2048)
	v88 = F_fcntl(m, v83, int32(4), v51+int32(16))
	mBase = m.M
	if v88 != 0 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	goto L23
L28:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v16)+112))
	*(*int32)(unsafe.Add(mBase, _consts[643])) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v16)+116))
	v121 = int32(-1)
	v128 = m.G0
	v130 = v128 - int32(64)
	m.G0 = v130
	v133 = F_pipe(m, v16+int32(112))
	mBase = m.M
	if v133 != 0 {
		v190 = v121
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v16)+116))
	*(*int32)(unsafe.Add(mBase, _consts[644])) = v201
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v16)+112))
	v205 = v203
	v206 = v120
	goto L11
L30:
	;
	if v190 != int32(-1) {
		goto L29
	} else {
		goto L45
	}
L31:
	;
	m.G0 = v130 + int32(64)
	goto L30
L32:
	;
	goto L34
L34:
	;
	goto L37
L37:
	;
	goto L40
L40:
	;
	v190 = int32(0)
	goto L31
L45:
	;
	v196 = F_close(m, v120)
	mBase = m.M
	v198 = *(*int32)(unsafe.Add(mBase, _consts[643]))
	v199 = F_close(m, v198)
	mBase = m.M
	v680 = v121
	goto L3
L46:
	;
	return int32(0)
L47:
	;
	*(*int32)(unsafe.Add(mBase, _consts[645])) = int32(0)
	if v41 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v228 = v16 + int32(120)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v229
	goto L50
L49:
	;
	v220 = int32(_a44)
	*(*int64)(unsafe.Add(mBase, _consts[646])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[645])) = v213
	goto L48
L50:
	;
	v234 = v16 + int32(120)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if v236 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v370 = F_serverFork(m, int32(1))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L46
	} else {
		goto L90
	}
L52:
	;
	v254 = v236
	v258 = int32(1)
	v260 = int32(0)
	goto L57
L53:
	;
	if v236 != 0 {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L53
L55:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v236+base.B2i32(v239 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v245
	goto L54
L56:
	;
	v364 = int32(0)
	v366 = int32(0)
	goto L51
L57:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v254)+8))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+104))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	if v265 != int32(6) {
		v321 = v260
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v364 = base.B2i32(v337 == int32(0))
	v366 = v338
	goto L51
L59:
	;
	v341 = v16 + int32(120)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	if v343 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L60:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+132))
	if v325 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L61:
	;
	v268 = int32(*(*int16)(unsafe.Add(mBase, uint32(v264)+162)))
	if l0 != v268 {
		v337 = v258
		v338 = v260
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v263)+104))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+160)))
	if v271&int32(1) != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v286 != l1 {
		v337 = v258
		v338 = v260
		goto L59
	} else {
		goto L70
	}
L64:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v270)+156))
	if v275 <= int32(589823) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v286 = int32(80)
	goto L63
L66:
	;
	if v275 < int32(459264) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _consts[647]))
	v286 = v279
	goto L63
L68:
	;
	v286 = int32(11)
	goto L63
L69:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _consts[648]))
	v286 = v283
	goto L63
L70:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v213+v260<<(uint(int32(2))%32)))) = v291
	if v41 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v318 = *(*int64)(unsafe.Add(mBase, _consts[40]))
	goto L78
L72:
	;
	v309 = int32(_a44)
	v311 = *(*int32)(unsafe.Add(mBase, _consts[646]))
	*(*int32)(unsafe.Add(mBase, _consts[646])) = v311 + int32(1)
	goto L71
L73:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _consts[218]))
	v300 = F_connSendTimeout(m, v291, base.I64_extend_i32_s(v296*int32(1000)))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L46
	} else {
		goto L74
	}
L74:
	;
	v302 = F_sendCurrentOffsetToReplica(m, v263)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L46
	} else {
		goto L75
	}
L75:
	;
	F_addRdbReplicaToPsyncWait(m, v263)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L46
	} else {
		goto L76
	}
L76:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	v307 = F_connBlock(m, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L46
	} else {
		goto L77
	}
L77:
	;
	goto L71
L78:
	;
	v319 = F_replicationSetupReplicaForFullResync(m, v263, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L46
	} else {
		goto L79
	}
L79:
	;
	v321 = v260 + int32(1)
	goto L60
L80:
	;
	v337 = int32(0)
	v338 = v321
	goto L59
L81:
	;
	v328 = m.T0[v325].(func(*base.Module) int32)(m)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L46
	} else {
		goto L82
	}
L82:
	;
	if v328 == int32(0) {
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v263)+104))
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+160)))
	if v333&int32(8) != 0 {
		v337 = v258
		v338 = v321
		goto L59
	} else {
		goto L84
	}
L84:
	;
	goto L80
L85:
	;
	if v343 != 0 {
		v254 = v343
		v258 = v337
		v260 = v338
		goto L57
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v343+base.B2i32(v346 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = v352
	goto L86
L88:
	;
	goto L58
L89:
	;
	v484 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v370 != int32(-1) {
		goto L139
	} else {
		goto L140
	}
L90:
	;
	if v370 != 0 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	if v41 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _consts[54]))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v390 = int32(_a127)
	v393 = int32(*(*int8)(unsafe.Add(mBase, _consts[55])))
	if v393 != 0 {
		goto L98
	} else {
		goto L99
	}
L93:
	;
	F_rioInitWithFd(m, v16+int32(32), v206)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L46
	} else {
		goto L96
	}
L94:
	;
	F_rioInitWithConnset(m, v16+int32(32), v213, v366)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L46
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _consts[643]))
	v384 = F_close(m, v383)
	mBase = m.M
	goto L92
L97:
	;
	goto L113
L98:
	;
	v395 = F_strchr(m, v389, v393)
	mBase = m.M
	if v395 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L97
L100:
	;
	goto L97
L101:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v398 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+1)))
	if v399 == int32(0) {
		goto L100
	} else {
		goto L104
	}
L103:
	;
	goto L97
L104:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, _consts[57])))
	if v402 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+2)))
	if v404 == int32(0) {
		goto L100
	} else {
		goto L107
	}
L106:
	;
	v403 = F_twobyte_strstr(m, v395, v390)
	mBase = m.M
	goto L97
L107:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, _consts[58])))
	if v407 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+3)))
	if v409 == int32(0) {
		goto L100
	} else {
		goto L110
	}
L109:
	;
	v408 = F_threebyte_strstr(m, v395, v390)
	mBase = m.M
	goto L97
L110:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, _consts[59])))
	if v412 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v414 = F_twoway_strstr(m, v395, v390)
	mBase = m.M
	goto L100
L112:
	;
	v413 = F_fourbyte_strstr(m, v395, v390)
	mBase = m.M
	goto L97
L113:
	;
	goto L115
L115:
	;
	goto L116
L116:
	;
	goto L117
L117:
	;
	if v364 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v427 = int32(0)
	v431 = F_rdbSaveRioWithEOFMark(m, l0, l1, v16+int32(32), v427, l2)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L46
	} else {
		goto L124
	}
L119:
	;
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v16)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v423 | int64(8)
	goto L118
L120:
	;
	F__exit(m, v479^int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_rioFreeFd(m, v16+int32(32))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L46
	} else {
		goto L134
	}
L122:
	;
	F_rioFreeConnset(m, v16+int32(32))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L46
	} else {
		goto L132
	}
L123:
	;
	if v41 == int32(0) {
		v464 = v427
		goto L121
	} else {
		goto L131
	}
L124:
	;
	if v431 != 0 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v436 = m.T0[v435].(func(*base.Module, int32) int32)(m, v16+int32(32))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L46
	} else {
		goto L126
	}
L126:
	;
	if v436 == int32(0) {
		goto L123
	} else {
		goto L127
	}
L127:
	;
	F_sendChildCowInfo(m, int32(2), int32(_a1182))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L46
	} else {
		goto L128
	}
L128:
	;
	if v41 == int32(0) {
		v464 = int32(1)
		goto L121
	} else {
		goto L129
	}
L129:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
	F_sendChildInfoGeneric(m, int32(5), int32(0), v449, float64(-1), int32(_a1182))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L46
	} else {
		goto L130
	}
L130:
	;
	v457 = int32(1)
	goto L122
L131:
	;
	v457 = v427
	goto L122
L132:
	;
	F_valkey_free(m, v213)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L46
	} else {
		goto L133
	}
L133:
	;
	v479 = v457
	goto L120
L134:
	;
	v469 = F_close(m, v206)
	mBase = m.M
	v471 = *(*int32)(unsafe.Add(mBase, _consts[644]))
	v472 = F_close(m, v471)
	mBase = m.M
	F_valkey_free(m, v213)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L46
	} else {
		goto L135
	}
L135:
	;
	v478 = F_read(m, v205, v16+int32(112), int32(1))
	mBase = m.M
	v479 = v464
	goto L120
L137:
	;
	v671 = int32(-1)
	if v370 == v671 {
		goto L183
	} else {
		goto L184
	}
L138:
	;
	v657 = F_close(m, v205)
	mBase = m.M
	goto L137
L139:
	;
	if int32(2) < v484 {
		goto L169
	} else {
		goto L170
	}
L140:
	;
	if int32(3) < v484 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v500 = v16 + int32(120)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	*(*int32)(unsafe.Add(mBase, uint32(v500)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v500))) = v501
	goto L146
L142:
	;
	goto L143
L143:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v491 = F___strerror_l(m, v490, v490)
	mBase = m.M
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v491
	F__serverLog(m, int32(3), int32(_a1183), v16)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L46
	} else {
		goto L145
	}
L145:
	;
	goto L141
L146:
	;
	v506 = v16 + int32(120)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v506)))
	if v508 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	if v41 != 0 {
		goto L160
	} else {
		goto L161
	}
L148:
	;
	if v508 == int32(0) {
		goto L147
	} else {
		goto L151
	}
L149:
	;
	goto L148
L150:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v508+base.B2i32(v511 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v506))) = v517
	goto L149
L151:
	;
	v525 = v508
	goto L152
L152:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v525)+8))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v534)+104))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)))
	if v536 != int32(7) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L147
L154:
	;
	v542 = v16 + int32(120)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v542)))
	if v544 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v535))) = int32(6)
	goto L154
L156:
	;
	if v544 != 0 {
		v525 = v544
		goto L152
	} else {
		goto L159
	}
L157:
	;
	goto L156
L158:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v544+base.B2i32(v547 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v542))) = v553
	goto L157
L159:
	;
	goto L153
L160:
	;
	F_valkey_free(m, v213)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L46
	} else {
		goto L163
	}
L161:
	;
	v568 = F_close(m, v206)
	mBase = m.M
	v569 = int32(_a44)
	v570 = *(*int32)(unsafe.Add(mBase, _consts[643]))
	v571 = F_close(m, v570)
	mBase = m.M
	v573 = *(*int32)(unsafe.Add(mBase, _consts[644]))
	v574 = F_close(m, v573)
	mBase = m.M
	F_valkey_free(m, v213)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L46
	} else {
		goto L162
	}
L162:
	;
	v577 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[649])) = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[645])) = int64(0)
	goto L138
L163:
	;
	v587 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	if v587 != int32(-1) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L137
L165:
	;
	goto L164
L166:
	;
	v594 = F_close(m, v587)
	mBase = m.M
	v595 = int32(_a44)
	v596 = *(*int32)(unsafe.Add(mBase, _consts[650]))
	v597 = F_close(m, v596)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[109])) = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[108])) = int64(-1)
	goto L165
L167:
	;
	v591 = *(*int32)(unsafe.Add(mBase, _consts[650]))
	if v591 == int32(-1) {
		goto L165
	} else {
		goto L168
	}
L168:
	;
	goto L166
L169:
	;
	v621 = int32(0)
	v622 = F___time(m, v621)
	mBase = m.M
	v623 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[651])) = int32(2)
	*(*int64)(unsafe.Add(mBase, _consts[652])) = v622
	if v41 == v621 {
		goto L178
	} else {
		goto L179
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v370
	if v364 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v609 = int32(_a139)
	goto L173
L172:
	;
	v609 = int32(_a1184)
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v609
	if v41 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v613 = int32(_a1185)
	goto L176
L175:
	;
	v613 = int32(_a1186)
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v613
	F__serverLog(m, int32(2), int32(_a1187), v16+int32(16))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L46
	} else {
		goto L177
	}
L177:
	;
	goto L169
L178:
	;
	v632 = F_close(m, v206)
	mBase = m.M
	v633 = int32(_a44)
	v634 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v636 = *(*int32)(unsafe.Add(mBase, _consts[643]))
	v640 = F_aeCreateFileEvent(m, v634, v636, int32(1), int32(969), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L46
	} else {
		goto L181
	}
L179:
	;
	F_valkey_free(m, v213)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L46
	} else {
		goto L180
	}
L180:
	;
	goto L137
L181:
	;
	if v640 == int32(-1) {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	goto L138
L183:
	;
	v675 = v671
	goto L185
L184:
	;
	v675 = int32(0)
	goto L185
L185:
	;
	v680 = v675
	goto L3
L186:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L187:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_rdbSavedObjectLen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	v6 = F_rdbGetObjectType(m, l0, int32(80))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(-1) {
			F__serverAssert(m, int32(_a1126), int32(_a1122), int32(1176))
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v13 = F_rdbSaveObject(m, int32(0), l0, l1, l2, v6)
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v13 == int32(-1) {
					F__serverAssertWithInfo(m, int32(0), l0, int32(_a1127), int32(_a1122), int32(1178))
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					return v13
				}
			}
		}
	}
}
func F_rdbShowGenericInfo(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	v12 = m.G0
	v14 = v12 - int32(352)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[965]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = v17
	v22 = F_iprintf(m, int32(_a1824), v14+int32(192))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[966]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = v25
	v30 = F_iprintf(m, int32(_a1825), v14+int32(176))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[967]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v33
	v38 = F_iprintf(m, int32(_a1826), v14+int32(160))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[968]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v41
	v46 = F_iprintf(m, int32(_a1827), v14+int32(144))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v48 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, _consts[969]))
	if v49 == v48 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, _consts[970]))
	if v59 == v58 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v49
	v56 = F_iprintf(m, int32(_a1828), v14+int32(128))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _consts[971]))
	goto L57
L10:
	;
	m.G0 = v14 + int32(352)
	return
L11:
	;
	v62 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, _consts[972]))
	if v64 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v83 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, _consts[973]))
	if v84 < v83 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v68 = F_dup(m, int32(1))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = int32(420)
	v72 = *(*int32)(unsafe.Add(mBase, _consts[971]))
	v76 = F_open(m, v72, int32(65), v14+int32(112))
	mBase = m.M
	if v76 == int32(-1) {
		goto L9
	} else {
		goto L15
	}
L14:
	;
	v65 = int32(-1)
	v81 = v65
	v82 = v65
	goto L12
L15:
	;
	v80 = F_dup2(m, v76, int32(1))
	mBase = m.M
	v81 = v68
	v82 = v76
	goto L12
L16:
	;
	v276 = int32(0)
	v277 = *(*int32)(unsafe.Add(mBase, _consts[972]))
	if v277 == v276 {
		goto L10
	} else {
		goto L56
	}
L17:
	;
	v89 = v62
	v92 = v84
	goto L18
L18:
	;
	v98 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, _consts[974]))
	if v99 == v98 {
		v256 = v92
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L16
L20:
	;
	if v89 < v256 {
		v89 = v89 + int32(1)
		v92 = v256
		goto L18
	} else {
		goto L55
	}
L21:
	;
	v104 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, _consts[975]))
	v109 = v99
	v115 = v104
	v116 = v106
	v117 = int32(_a1829)
	goto L22
L22:
	;
	switch v116 + int32(-1) {
	case 0:
		goto L26
	case 1:
		goto L25
	default:
		goto L24
	}
L23:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _consts[973]))
	v256 = v250
	goto L20
L24:
	;
	v146 = int32(0)
	v148 = *(*int32)(unsafe.Add(mBase, _consts[975]))
	v150 = v146
	v154 = v148
	v157 = v148
	goto L30
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v89
	v144 = F_iprintf(m, int32(_a1830), v14+int32(80))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L29
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v89
	v129 = F_snprintf(m, v14+int32(208), int32(80), int32(_a1830), v14+int32(64))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v14 + int32(208)
	v137 = F_iprintf(m, int32(_a1831), v14+int32(48))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L24
L29:
	;
	goto L24
L30:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _consts[976]))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161+v150<<(uint(int32(2))%32)+v89*int32(28))))
	switch v154 {
	case 0:
		goto L36
	case 1:
		goto L35
	case 2:
		goto L34
	default:
		goto L33
	}
L31:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v227+int32(-1)) {
		v241 = v227
		goto L51
	} else {
		goto L52
	}
L32:
	;
	v229 = v150 + int32(1)
	if v229 != int32(7) {
		v150 = v229
		v154 = v226
		v157 = v227
		goto L30
	} else {
		goto L50
	}
L33:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	F_rdbStatsPrintInfo(m, v166, v201, v14+int32(288), int32(64))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L43
	}
L34:
	;
	v198 = F_putchar(m, int32(44))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L42
	}
L35:
	;
	v195 = F_putchar(m, int32(9))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L41
	}
L36:
	;
	if v115 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v168<<(uint(int32(2))%32))+uint32(_consts[964])))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v89
	v184 = F_snprintf(m, v14+int32(208), int32(80), int32(_a1832), v14+int32(32))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v226 = int32(0)
	v227 = v157
	goto L32
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v14 + int32(208)
	v192 = F_iprintf(m, int32(_a1833), v14+int32(16))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L33
L41:
	;
	goto L33
L42:
	;
	goto L33
L43:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _consts[975]))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v14 + int32(288)
	if v208 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v216 = int32(_a1834)
	goto L46
L45:
	;
	v216 = int32(_a79)
	goto L46
L46:
	;
	v217 = F_iprintf(m, v216, v14)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _consts[975]))
	if v220 != 0 {
		v226 = v220
		v227 = v220
		goto L32
	} else {
		goto L48
	}
L48:
	;
	v222 = F_putchar(m, int32(10))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _consts[975]))
	v226 = v225
	v227 = v225
	goto L32
L50:
	;
	goto L31
L51:
	;
	v243 = v115 + int32(1)
	v245 = v243 << (uint(int32(2)) % 32)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v245)+uint32(_consts[974])))
	if v248 != 0 {
		v109 = v248
		v115 = v243
		v116 = v241
		v117 = v245 + int32(_a1829)
		goto L22
	} else {
		goto L54
	}
L52:
	;
	v237 = F_putchar(m, int32(10))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _consts[975]))
	v241 = v240
	goto L51
L54:
	;
	goto L23
L55:
	;
	goto L19
L56:
	;
	v281 = F_dup2(m, v81, int32(1))
	mBase = m.M
	v282 = F_close(m, v82)
	mBase = m.M
	v283 = F_close(m, v81)
	mBase = m.M
	goto L10
L57:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v302 = F___strerror_l(m, v301, v301)
	mBase = m.M
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+100)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v299
	v306 = *(*int32)(unsafe.Add(mBase, _consts[514]))
	v310 = F_fiprintf(m, v306, int32(_a1835), v14+int32(96))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_rdbStatsPrintInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
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
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
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
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
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
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
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
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
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
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v474 float64
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 float64
	_ = v521
	var v529 float64
	_ = v529
	var v532 float64
	_ = v532
	var v535 int64
	_ = v535
	var v539 float64
	_ = v539
	var v545 int64
	_ = v545
	var v547 int64
	_ = v547
	var v548 int32
	_ = v548
	var v551 int64
	_ = v551
	var v554 int64
	_ = v554
	var v555 int32
	_ = v555
	var v561 int64
	_ = v561
	var v565 int32
	_ = v565
	var v570 int64
	_ = v570
	var v571 int64
	_ = v571
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v609 int64
	_ = v609
	var v616 int32
	_ = v616
	var v617 int64
	_ = v617
	var v622 int32
	_ = v622
	var v623 int64
	_ = v623
	var v624 int64
	_ = v624
	var v626 int64
	_ = v626
	var v630 int32
	_ = v630
	var v639 int64
	_ = v639
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 float64
	_ = v688
	var v696 float64
	_ = v696
	var v699 float64
	_ = v699
	var v702 int64
	_ = v702
	var v706 float64
	_ = v706
	var v712 int64
	_ = v712
	var v714 int64
	_ = v714
	var v715 int32
	_ = v715
	var v718 int64
	_ = v718
	var v721 int64
	_ = v721
	var v722 int32
	_ = v722
	var v728 int64
	_ = v728
	var v732 int32
	_ = v732
	var v737 int64
	_ = v737
	var v738 int64
	_ = v738
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v761 int32
	_ = v761
	var v776 int64
	_ = v776
	var v783 int32
	_ = v783
	var v784 int64
	_ = v784
	var v789 int32
	_ = v789
	var v790 int64
	_ = v790
	var v791 int64
	_ = v791
	var v793 int64
	_ = v793
	var v797 int32
	_ = v797
	var v806 int64
	_ = v806
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 float64
	_ = v855
	var v863 float64
	_ = v863
	var v866 float64
	_ = v866
	var v869 int64
	_ = v869
	var v873 float64
	_ = v873
	var v879 int64
	_ = v879
	var v881 int64
	_ = v881
	var v882 int32
	_ = v882
	var v885 int64
	_ = v885
	var v888 int64
	_ = v888
	var v889 int32
	_ = v889
	var v895 int64
	_ = v895
	var v899 int32
	_ = v899
	var v904 int64
	_ = v904
	var v905 int64
	_ = v905
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v928 int32
	_ = v928
	var v943 int64
	_ = v943
	var v950 int32
	_ = v950
	var v951 int64
	_ = v951
	var v956 int32
	_ = v956
	var v957 int64
	_ = v957
	var v958 int64
	_ = v958
	var v960 int64
	_ = v960
	var v964 int32
	_ = v964
	var v973 int64
	_ = v973
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1074 float64
	_ = v1074
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 float64
	_ = v1121
	var v1129 float64
	_ = v1129
	var v1132 float64
	_ = v1132
	var v1135 int64
	_ = v1135
	var v1139 float64
	_ = v1139
	var v1145 int64
	_ = v1145
	var v1147 int64
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int64
	_ = v1151
	var v1154 int64
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1161 int64
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1170 int64
	_ = v1170
	var v1171 int64
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1194 int32
	_ = v1194
	var v1209 int64
	_ = v1209
	var v1216 int32
	_ = v1216
	var v1217 int64
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1223 int64
	_ = v1223
	var v1224 int64
	_ = v1224
	var v1226 int64
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1239 int64
	_ = v1239
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1288 float64
	_ = v1288
	var v1296 float64
	_ = v1296
	var v1299 float64
	_ = v1299
	var v1302 int64
	_ = v1302
	var v1306 float64
	_ = v1306
	var v1312 int64
	_ = v1312
	var v1314 int64
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1318 int64
	_ = v1318
	var v1321 int64
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1328 int64
	_ = v1328
	var v1332 int32
	_ = v1332
	var v1337 int64
	_ = v1337
	var v1338 int64
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1361 int32
	_ = v1361
	var v1376 int64
	_ = v1376
	var v1383 int32
	_ = v1383
	var v1384 int64
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1390 int64
	_ = v1390
	var v1391 int64
	_ = v1391
	var v1393 int64
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1406 int64
	_ = v1406
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1455 float64
	_ = v1455
	var v1463 float64
	_ = v1463
	var v1466 float64
	_ = v1466
	var v1469 int64
	_ = v1469
	var v1473 float64
	_ = v1473
	var v1479 int64
	_ = v1479
	var v1481 int64
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1485 int64
	_ = v1485
	var v1488 int64
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1495 int64
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1504 int64
	_ = v1504
	var v1505 int64
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1528 int32
	_ = v1528
	var v1543 int64
	_ = v1543
	var v1550 int32
	_ = v1550
	var v1551 int64
	_ = v1551
	var v1556 int32
	_ = v1556
	var v1557 int64
	_ = v1557
	var v1558 int64
	_ = v1558
	var v1560 int64
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1573 int64
	_ = v1573
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	v7 = m.G0
	v9 = v7 - int32(288)
	m.G0 = v9
	v11 = int32(_a1804)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v14 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v9 + int32(288)
	return
L2:
	;
	v60 = int32(_a1805)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v63 != 0 {
		goto L21
	} else {
		goto L22
	}
L3:
	;
	if v46-v48 != 0 {
		goto L2
	} else {
		goto L15
	}
L4:
	;
	v46 = F_tolower(m, v42)
	mBase = m.M
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v48 = F_tolower(m, v47)
	mBase = m.M
	goto L3
L5:
	;
	v16 = l1
	v17 = v11
	v18 = v14
	goto L8
L6:
	;
	v42 = int32(0)
	v43 = v11
	goto L4
L7:
	;
	v42 = v39 & int32(255)
	v43 = v38
	goto L4
L8:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v20 == int32(0) {
		v38 = v17
		v39 = v18
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v38 = v32
	v39 = int32(0)
	goto L7
L10:
	;
	v24 = v18 & int32(255)
	if v24 == v20 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = int32(1)
	v32 = v17 + v31
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v33 != 0 {
		v16 = v16 + v31
		v17 = v32
		v18 = v33
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v26 = F_tolower(m, v24)
	mBase = m.M
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v28 = F_tolower(m, v27)
	mBase = m.M
	if v26 == v28 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v38 = v17
	v39 = v30
	goto L7
L14:
	;
	goto L9
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50<<(uint(int32(2))%32))+uint32(_consts[964])))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v55
	v58 = F_snprintf(m, l2, l3, int32(_a79), v9)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
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
	v106 = int32(_a1806)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v109 != 0 {
		goto L36
	} else {
		goto L37
	}
L19:
	;
	if v95-v97 != 0 {
		goto L18
	} else {
		goto L31
	}
L20:
	;
	v95 = F_tolower(m, v91)
	mBase = m.M
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v97 = F_tolower(m, v96)
	mBase = m.M
	goto L19
L21:
	;
	v65 = l1
	v66 = v60
	v67 = v63
	goto L24
L22:
	;
	v91 = int32(0)
	v92 = v60
	goto L20
L23:
	;
	v91 = v88 & int32(255)
	v92 = v87
	goto L20
L24:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v69 == int32(0) {
		v87 = v66
		v88 = v67
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v87 = v81
	v88 = int32(0)
	goto L23
L26:
	;
	v73 = v67 & int32(255)
	if v73 == v69 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v80 = int32(1)
	v81 = v66 + v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	if v82 != 0 {
		v65 = v65 + v80
		v66 = v81
		v67 = v82
		goto L24
	} else {
		goto L30
	}
L28:
	;
	v75 = F_tolower(m, v73)
	mBase = m.M
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v77 = F_tolower(m, v76)
	mBase = m.M
	if v75 == v77 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v87 = v66
	v88 = v79
	goto L23
L30:
	;
	goto L25
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v99
	v104 = F_snprintf(m, l2, l3, int32(_a1807), v9+int32(16))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	goto L1
L33:
	;
	v152 = int32(_a1808)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v155 != 0 {
		goto L51
	} else {
		goto L52
	}
L34:
	;
	if v141-v143 != 0 {
		goto L33
	} else {
		goto L46
	}
L35:
	;
	v141 = F_tolower(m, v137)
	mBase = m.M
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v143 = F_tolower(m, v142)
	mBase = m.M
	goto L34
L36:
	;
	v111 = l1
	v112 = v106
	v113 = v109
	goto L39
L37:
	;
	v137 = int32(0)
	v138 = v106
	goto L35
L38:
	;
	v137 = v134 & int32(255)
	v138 = v133
	goto L35
L39:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v115 == int32(0) {
		v133 = v112
		v134 = v113
		goto L38
	} else {
		goto L41
	}
L40:
	;
	v133 = v127
	v134 = int32(0)
	goto L38
L41:
	;
	v119 = v113 & int32(255)
	if v119 == v115 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v126 = int32(1)
	v127 = v112 + v126
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
	if v128 != 0 {
		v111 = v111 + v126
		v112 = v127
		v113 = v128
		goto L39
	} else {
		goto L45
	}
L43:
	;
	v121 = F_tolower(m, v119)
	mBase = m.M
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v123 = F_tolower(m, v122)
	mBase = m.M
	if v121 == v123 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v133 = v112
	v134 = v125
	goto L38
L45:
	;
	goto L40
L46:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v145
	v150 = F_snprintf(m, l2, l3, int32(_a1807), v9+int32(32))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L16
	} else {
		goto L47
	}
L47:
	;
	goto L1
L48:
	;
	v198 = int32(_a1809)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v201 != 0 {
		goto L66
	} else {
		goto L67
	}
L49:
	;
	if v187-v189 != 0 {
		goto L48
	} else {
		goto L61
	}
L50:
	;
	v187 = F_tolower(m, v183)
	mBase = m.M
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	v189 = F_tolower(m, v188)
	mBase = m.M
	goto L49
L51:
	;
	v157 = l1
	v158 = v152
	v159 = v155
	goto L54
L52:
	;
	v183 = int32(0)
	v184 = v152
	goto L50
L53:
	;
	v183 = v180 & int32(255)
	v184 = v179
	goto L50
L54:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v161 == int32(0) {
		v179 = v158
		v180 = v159
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v179 = v173
	v180 = int32(0)
	goto L53
L56:
	;
	v165 = v159 & int32(255)
	if v165 == v161 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v172 = int32(1)
	v173 = v158 + v172
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
	if v174 != 0 {
		v157 = v157 + v172
		v158 = v173
		v159 = v174
		goto L54
	} else {
		goto L60
	}
L58:
	;
	v167 = F_tolower(m, v165)
	mBase = m.M
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	v169 = F_tolower(m, v168)
	mBase = m.M
	if v167 == v169 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	v179 = v158
	v180 = v171
	goto L53
L60:
	;
	goto L55
L61:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v191
	v196 = F_snprintf(m, l2, l3, int32(_a1807), v9+int32(48))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	goto L1
L63:
	;
	v244 = int32(_a1810)
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v247 != 0 {
		goto L81
	} else {
		goto L82
	}
L64:
	;
	if v233-v235 != 0 {
		goto L63
	} else {
		goto L76
	}
L65:
	;
	v233 = F_tolower(m, v229)
	mBase = m.M
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	v235 = F_tolower(m, v234)
	mBase = m.M
	goto L64
L66:
	;
	v203 = l1
	v204 = v198
	v205 = v201
	goto L69
L67:
	;
	v229 = int32(0)
	v230 = v198
	goto L65
L68:
	;
	v229 = v226 & int32(255)
	v230 = v225
	goto L65
L69:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v207 == int32(0) {
		v225 = v204
		v226 = v205
		goto L68
	} else {
		goto L71
	}
L70:
	;
	v225 = v219
	v226 = int32(0)
	goto L68
L71:
	;
	v211 = v205 & int32(255)
	if v211 == v207 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v218 = int32(1)
	v219 = v204 + v218
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+1)))
	if v220 != 0 {
		v203 = v203 + v218
		v204 = v219
		v205 = v220
		goto L69
	} else {
		goto L75
	}
L73:
	;
	v213 = F_tolower(m, v211)
	mBase = m.M
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	v215 = F_tolower(m, v214)
	mBase = m.M
	if v213 == v215 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v225 = v204
	v226 = v217
	goto L68
L75:
	;
	goto L70
L76:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v237
	v242 = F_snprintf(m, l2, l3, int32(_a1807), v9+int32(64))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L16
	} else {
		goto L77
	}
L77:
	;
	goto L1
L78:
	;
	v290 = int32(_a1811)
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v293 != 0 {
		goto L96
	} else {
		goto L97
	}
L79:
	;
	if v279-v281 != 0 {
		goto L78
	} else {
		goto L91
	}
L80:
	;
	v279 = F_tolower(m, v275)
	mBase = m.M
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	v281 = F_tolower(m, v280)
	mBase = m.M
	goto L79
L81:
	;
	v249 = l1
	v250 = v244
	v251 = v247
	goto L84
L82:
	;
	v275 = int32(0)
	v276 = v244
	goto L80
L83:
	;
	v275 = v272 & int32(255)
	v276 = v271
	goto L80
L84:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v253 == int32(0) {
		v271 = v250
		v272 = v251
		goto L83
	} else {
		goto L86
	}
L85:
	;
	v271 = v265
	v272 = int32(0)
	goto L83
L86:
	;
	v257 = v251 & int32(255)
	if v257 == v253 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v264 = int32(1)
	v265 = v250 + v264
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
	if v266 != 0 {
		v249 = v249 + v264
		v250 = v265
		v251 = v266
		goto L84
	} else {
		goto L90
	}
L88:
	;
	v259 = F_tolower(m, v257)
	mBase = m.M
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v261 = F_tolower(m, v260)
	mBase = m.M
	if v259 == v261 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	v271 = v250
	v272 = v263
	goto L83
L90:
	;
	goto L85
L91:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v283
	v288 = F_snprintf(m, l2, l3, int32(_a1807), v9+int32(80))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L16
	} else {
		goto L92
	}
L92:
	;
	goto L1
L93:
	;
	v336 = int32(_a1812)
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v339 != 0 {
		goto L111
	} else {
		goto L112
	}
L94:
	;
	if v325-v327 != 0 {
		goto L93
	} else {
		goto L106
	}
L95:
	;
	v325 = F_tolower(m, v321)
	mBase = m.M
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	v327 = F_tolower(m, v326)
	mBase = m.M
	goto L94
L96:
	;
	v295 = l1
	v296 = v290
	v297 = v293
	goto L99
L97:
	;
	v321 = int32(0)
	v322 = v290
	goto L95
L98:
	;
	v321 = v318 & int32(255)
	v322 = v317
	goto L95
L99:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	if v299 == int32(0) {
		v317 = v296
		v318 = v297
		goto L98
	} else {
		goto L101
	}
L100:
	;
	v317 = v311
	v318 = int32(0)
	goto L98
L101:
	;
	v303 = v297 & int32(255)
	if v303 == v299 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v310 = int32(1)
	v311 = v296 + v310
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+1)))
	if v312 != 0 {
		v295 = v295 + v310
		v296 = v311
		v297 = v312
		goto L99
	} else {
		goto L105
	}
L103:
	;
	v305 = F_tolower(m, v303)
	mBase = m.M
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	v307 = F_tolower(m, v306)
	mBase = m.M
	if v305 == v307 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	v317 = v296
	v318 = v309
	goto L98
L105:
	;
	goto L100
L106:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v329
	v334 = F_snprintf(m, l2, l3, int32(_a1807), v9+int32(96))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L16
	} else {
		goto L107
	}
L107:
	;
	goto L1
L108:
	;
	v382 = int32(_a1813)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v385 != 0 {
		goto L126
	} else {
		goto L127
	}
L109:
	;
	if v371-v373 != 0 {
		goto L108
	} else {
		goto L121
	}
L110:
	;
	v371 = F_tolower(m, v367)
	mBase = m.M
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	v373 = F_tolower(m, v372)
	mBase = m.M
	goto L109
L111:
	;
	v341 = l1
	v342 = v336
	v343 = v339
	goto L114
L112:
	;
	v367 = int32(0)
	v368 = v336
	goto L110
L113:
	;
	v367 = v364 & int32(255)
	v368 = v363
	goto L110
L114:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	if v345 == int32(0) {
		v363 = v342
		v364 = v343
		goto L113
	} else {
		goto L116
	}
L115:
	;
	v363 = v357
	v364 = int32(0)
	goto L113
L116:
	;
	v349 = v343 & int32(255)
	if v349 == v345 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v356 = int32(1)
	v357 = v342 + v356
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+1)))
	if v358 != 0 {
		v341 = v341 + v356
		v342 = v357
		v343 = v358
		goto L114
	} else {
		goto L120
	}
L118:
	;
	v351 = F_tolower(m, v349)
	mBase = m.M
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	v353 = F_tolower(m, v352)
	mBase = m.M
	if v351 == v353 {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	v363 = v342
	v364 = v355
	goto L113
L120:
	;
	goto L115
L121:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v375
	v380 = F_snprintf(m, l2, l3, int32(_a1807), v9+int32(112))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L16
	} else {
		goto L122
	}
L122:
	;
	goto L1
L123:
	;
	v428 = int32(_a1814)
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v431 != 0 {
		goto L141
	} else {
		goto L142
	}
L124:
	;
	if v417-v419 != 0 {
		goto L123
	} else {
		goto L136
	}
L125:
	;
	v417 = F_tolower(m, v413)
	mBase = m.M
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	v419 = F_tolower(m, v418)
	mBase = m.M
	goto L124
L126:
	;
	v387 = l1
	v388 = v382
	v389 = v385
	goto L129
L127:
	;
	v413 = int32(0)
	v414 = v382
	goto L125
L128:
	;
	v413 = v410 & int32(255)
	v414 = v409
	goto L125
L129:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	if v391 == int32(0) {
		v409 = v388
		v410 = v389
		goto L128
	} else {
		goto L131
	}
L130:
	;
	v409 = v403
	v410 = int32(0)
	goto L128
L131:
	;
	v395 = v389 & int32(255)
	if v395 == v391 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v402 = int32(1)
	v403 = v388 + v402
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+1)))
	if v404 != 0 {
		v387 = v387 + v402
		v388 = v403
		v389 = v404
		goto L129
	} else {
		goto L135
	}
L133:
	;
	v397 = F_tolower(m, v395)
	mBase = m.M
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	v399 = F_tolower(m, v398)
	mBase = m.M
	if v397 == v399 {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	v409 = v388
	v410 = v401
	goto L128
L135:
	;
	goto L130
L136:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v421
	v426 = F_snprintf(m, l2, l3, int32(_a1807), v9+int32(128))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L16
	} else {
		goto L137
	}
L137:
	;
	goto L1
L138:
	;
	v481 = int32(_a1815)
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v484 != 0 {
		goto L159
	} else {
		goto L160
	}
L139:
	;
	if v463-v465 != 0 {
		goto L138
	} else {
		goto L151
	}
L140:
	;
	v463 = F_tolower(m, v459)
	mBase = m.M
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	v465 = F_tolower(m, v464)
	mBase = m.M
	goto L139
L141:
	;
	v433 = l1
	v434 = v428
	v435 = v431
	goto L144
L142:
	;
	v459 = int32(0)
	v460 = v428
	goto L140
L143:
	;
	v459 = v456 & int32(255)
	v460 = v455
	goto L140
L144:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	if v437 == int32(0) {
		v455 = v434
		v456 = v435
		goto L143
	} else {
		goto L146
	}
L145:
	;
	v455 = v449
	v456 = int32(0)
	goto L143
L146:
	;
	v441 = v435 & int32(255)
	if v441 == v437 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v448 = int32(1)
	v449 = v434 + v448
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+1)))
	if v450 != 0 {
		v433 = v433 + v448
		v434 = v449
		v435 = v450
		goto L144
	} else {
		goto L150
	}
L148:
	;
	v443 = F_tolower(m, v441)
	mBase = m.M
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	v445 = F_tolower(m, v444)
	mBase = m.M
	if v443 == v445 {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433))))
	v455 = v434
	v456 = v447
	goto L143
L150:
	;
	goto L145
L151:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v467 != 0 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+144)) = v474
	v479 = F_snprintf(m, l2, l3, int32(_a1816), v9+int32(144))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L16
	} else {
		goto L155
	}
L153:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v474 = base.F64_promote_f32(base.F32_div(base.F32_convert_i32_u(v469), base.F32_convert_i32_u(v467)))
	goto L152
L154:
	;
	v474 = float64(0)
	goto L152
L155:
	;
	goto L1
L156:
	;
	v648 = int32(_a1817)
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v651 != 0 {
		goto L200
	} else {
		goto L201
	}
L157:
	;
	if v516-v518 != 0 {
		goto L156
	} else {
		goto L169
	}
L158:
	;
	v516 = F_tolower(m, v512)
	mBase = m.M
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513))))
	v518 = F_tolower(m, v517)
	mBase = m.M
	goto L157
L159:
	;
	v486 = l1
	v487 = v481
	v488 = v484
	goto L162
L160:
	;
	v512 = int32(0)
	v513 = v481
	goto L158
L161:
	;
	v512 = v509 & int32(255)
	v513 = v508
	goto L158
L162:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487))))
	if v490 == int32(0) {
		v508 = v487
		v509 = v488
		goto L161
	} else {
		goto L164
	}
L163:
	;
	v508 = v502
	v509 = int32(0)
	goto L161
L164:
	;
	v494 = v488 & int32(255)
	if v494 == v490 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v501 = int32(1)
	v502 = v487 + v501
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486)+1)))
	if v503 != 0 {
		v486 = v486 + v501
		v487 = v502
		v488 = v503
		goto L162
	} else {
		goto L168
	}
L166:
	;
	v496 = F_tolower(m, v494)
	mBase = m.M
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487))))
	v498 = F_tolower(m, v497)
	mBase = m.M
	if v496 == v498 {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486))))
	v508 = v487
	v509 = v500
	goto L161
L168:
	;
	goto L163
L169:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v521 = float64(99)
	v529 = float64(100)
	if base.F64_lt(v521, v529) != 0 {
		goto L173
	} else {
		goto L174
	}
L170:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+160)) = base.F64_promote_f32(base.F32_convert_i64_s(v639))
	v646 = F_snprintf(m, l2, l3, int32(_a1816), v9+int32(160))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L16
	} else {
		goto L196
	}
L171:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v520)+80))
	if v548 < int32(1) {
		goto L178
	} else {
		goto L179
	}
L172:
	;
	v547 = int64(-9223372036854775807 - 1)
	goto L171
L173:
	;
	v532 = v521
	goto L175
L174:
	;
	v532 = v529
	goto L175
L175:
	;
	v535 = *(*int64)(unsafe.Add(mBase, uint32(v520)+88))
	v539 = base.F64_add(base.F64_mul(base.F64_div(v532, float64(100)), base.F64_convert_i64_s(v535)), float64(0.5))
	if base.F64_lt(base.F64_abs(v539), float64(9.223372036854776e+18)) == int32(0) {
		goto L172
	} else {
		goto L176
	}
L176:
	;
	v545 = base.I64_trunc_f64_s(v539)
	v547 = v545
	goto L171
L177:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v520)+24))
	v617 = *(*int64)(unsafe.Add(mBase, uint32(v520)+32))
	v622 = int32(63) - (v616 + base.I32_wrap_i64(base.I64_clz(v617|v609)))
	v623 = base.I64_extend_i32_u(v622)
	v624 = v609 >> (uint(v623) % 64)
	v626 = base.I64_extend32_s(v624) << (uint(v623) % 64)
	if base.F64_eq(v521, float64(0)) != 0 {
		v639 = v626
		goto L194
	} else {
		goto L195
	}
L178:
	;
	v609 = int64(0)
	goto L177
L179:
	;
	v551 = int64(1)
	if v551 < v547 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v554 = v547
	goto L182
L181:
	;
	v554 = v551
	goto L182
L182:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v520)+96))
	v561 = int64(0)
	v565 = int32(0)
	goto L183
L183:
	;
	v570 = *(*int64)(unsafe.Add(mBase, uint32(v555+v565<<(uint(int32(3))%32))))
	v571 = v570 + v561
	if v571 < v554 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	goto L178
L185:
	;
	v594 = v565 + int32(1)
	if v594 != v548 {
		v561 = v571
		v565 = v594
		goto L183
	} else {
		goto L193
	}
L186:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v520)+28))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v520)+24))
	v579 = int32(base.Ui32(v565) >> (uint(v578) % 32))
	if v579 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v580 = v573
	goto L189
L188:
	;
	v580 = int32(0)
	goto L189
L189:
	;
	v583 = int32(1)
	if v583 < v579 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v586 = v579
	goto L192
L191:
	;
	v586 = v583
	goto L192
L192:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v520)+16))
	v609 = base.I64_extend_i32_s((v573+int32(2147483647))&v565+v580) << (uint(base.I64_extend_i32_u(v586+v587+int32(-1))) % 64)
	goto L177
L193:
	;
	goto L184
L194:
	;
	goto L170
L195:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v520)+40))
	v639 = v626 + int64(1)<<(uint(base.I64_extend_i32_u(v622+base.B2i32(v630 <= base.I32_wrap_i64(v624))))%64) + int64(-1)
	goto L194
L196:
	;
	goto L1
L197:
	;
	v815 = int32(_a1818)
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v818 != 0 {
		goto L241
	} else {
		goto L242
	}
L198:
	;
	if v683-v685 != 0 {
		goto L197
	} else {
		goto L210
	}
L199:
	;
	v683 = F_tolower(m, v679)
	mBase = m.M
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680))))
	v685 = F_tolower(m, v684)
	mBase = m.M
	goto L198
L200:
	;
	v653 = l1
	v654 = v648
	v655 = v651
	goto L203
L201:
	;
	v679 = int32(0)
	v680 = v648
	goto L199
L202:
	;
	v679 = v676 & int32(255)
	v680 = v675
	goto L199
L203:
	;
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654))))
	if v657 == int32(0) {
		v675 = v654
		v676 = v655
		goto L202
	} else {
		goto L205
	}
L204:
	;
	v675 = v669
	v676 = int32(0)
	goto L202
L205:
	;
	v661 = v655 & int32(255)
	if v661 == v657 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v668 = int32(1)
	v669 = v654 + v668
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653)+1)))
	if v670 != 0 {
		v653 = v653 + v668
		v654 = v669
		v655 = v670
		goto L203
	} else {
		goto L209
	}
L207:
	;
	v663 = F_tolower(m, v661)
	mBase = m.M
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654))))
	v665 = F_tolower(m, v664)
	mBase = m.M
	if v663 == v665 {
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653))))
	v675 = v654
	v676 = v667
	goto L202
L209:
	;
	goto L204
L210:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v688 = float64(90)
	v696 = float64(100)
	if base.F64_lt(v688, v696) != 0 {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+176)) = base.F64_promote_f32(base.F32_convert_i64_s(v806))
	v813 = F_snprintf(m, l2, l3, int32(_a1816), v9+int32(176))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L16
	} else {
		goto L237
	}
L212:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v687)+80))
	if v715 < int32(1) {
		goto L219
	} else {
		goto L220
	}
L213:
	;
	v714 = int64(-9223372036854775807 - 1)
	goto L212
L214:
	;
	v699 = v688
	goto L216
L215:
	;
	v699 = v696
	goto L216
L216:
	;
	v702 = *(*int64)(unsafe.Add(mBase, uint32(v687)+88))
	v706 = base.F64_add(base.F64_mul(base.F64_div(v699, float64(100)), base.F64_convert_i64_s(v702)), float64(0.5))
	if base.F64_lt(base.F64_abs(v706), float64(9.223372036854776e+18)) == int32(0) {
		goto L213
	} else {
		goto L217
	}
L217:
	;
	v712 = base.I64_trunc_f64_s(v706)
	v714 = v712
	goto L212
L218:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v687)+24))
	v784 = *(*int64)(unsafe.Add(mBase, uint32(v687)+32))
	v789 = int32(63) - (v783 + base.I32_wrap_i64(base.I64_clz(v784|v776)))
	v790 = base.I64_extend_i32_u(v789)
	v791 = v776 >> (uint(v790) % 64)
	v793 = base.I64_extend32_s(v791) << (uint(v790) % 64)
	if base.F64_eq(v688, float64(0)) != 0 {
		v806 = v793
		goto L235
	} else {
		goto L236
	}
L219:
	;
	v776 = int64(0)
	goto L218
L220:
	;
	v718 = int64(1)
	if v718 < v714 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v721 = v714
	goto L223
L222:
	;
	v721 = v718
	goto L223
L223:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v687)+96))
	v728 = int64(0)
	v732 = int32(0)
	goto L224
L224:
	;
	v737 = *(*int64)(unsafe.Add(mBase, uint32(v722+v732<<(uint(int32(3))%32))))
	v738 = v737 + v728
	if v738 < v721 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L219
L226:
	;
	v761 = v732 + int32(1)
	if v761 != v715 {
		v728 = v738
		v732 = v761
		goto L224
	} else {
		goto L234
	}
L227:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v687)+28))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v687)+24))
	v746 = int32(base.Ui32(v732) >> (uint(v745) % 32))
	if v746 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v747 = v740
	goto L230
L229:
	;
	v747 = int32(0)
	goto L230
L230:
	;
	v750 = int32(1)
	if v750 < v746 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v753 = v746
	goto L233
L232:
	;
	v753 = v750
	goto L233
L233:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v687)+16))
	v776 = base.I64_extend_i32_s((v740+int32(2147483647))&v732+v747) << (uint(base.I64_extend_i32_u(v753+v754+int32(-1))) % 64)
	goto L218
L234:
	;
	goto L225
L235:
	;
	goto L211
L236:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v687)+40))
	v806 = v793 + int64(1)<<(uint(base.I64_extend_i32_u(v789+base.B2i32(v797 <= base.I32_wrap_i64(v791))))%64) + int64(-1)
	goto L235
L237:
	;
	goto L1
L238:
	;
	v982 = int32(_a1819)
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v985 != 0 {
		goto L282
	} else {
		goto L283
	}
L239:
	;
	if v850-v852 != 0 {
		goto L238
	} else {
		goto L251
	}
L240:
	;
	v850 = F_tolower(m, v846)
	mBase = m.M
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847))))
	v852 = F_tolower(m, v851)
	mBase = m.M
	goto L239
L241:
	;
	v820 = l1
	v821 = v815
	v822 = v818
	goto L244
L242:
	;
	v846 = int32(0)
	v847 = v815
	goto L240
L243:
	;
	v846 = v843 & int32(255)
	v847 = v842
	goto L240
L244:
	;
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821))))
	if v824 == int32(0) {
		v842 = v821
		v843 = v822
		goto L243
	} else {
		goto L246
	}
L245:
	;
	v842 = v836
	v843 = int32(0)
	goto L243
L246:
	;
	v828 = v822 & int32(255)
	if v828 == v824 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v835 = int32(1)
	v836 = v821 + v835
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820)+1)))
	if v837 != 0 {
		v820 = v820 + v835
		v821 = v836
		v822 = v837
		goto L244
	} else {
		goto L250
	}
L248:
	;
	v830 = F_tolower(m, v828)
	mBase = m.M
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821))))
	v832 = F_tolower(m, v831)
	mBase = m.M
	if v830 == v832 {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820))))
	v842 = v821
	v843 = v834
	goto L243
L250:
	;
	goto L245
L251:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v855 = float64(50)
	v863 = float64(100)
	if base.F64_lt(v855, v863) != 0 {
		goto L255
	} else {
		goto L256
	}
L252:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+192)) = base.F64_promote_f32(base.F32_convert_i64_s(v973))
	v980 = F_snprintf(m, l2, l3, int32(_a1816), v9+int32(192))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L16
	} else {
		goto L278
	}
L253:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v854)+80))
	if v882 < int32(1) {
		goto L260
	} else {
		goto L261
	}
L254:
	;
	v881 = int64(-9223372036854775807 - 1)
	goto L253
L255:
	;
	v866 = v855
	goto L257
L256:
	;
	v866 = v863
	goto L257
L257:
	;
	v869 = *(*int64)(unsafe.Add(mBase, uint32(v854)+88))
	v873 = base.F64_add(base.F64_mul(base.F64_div(v866, float64(100)), base.F64_convert_i64_s(v869)), float64(0.5))
	if base.F64_lt(base.F64_abs(v873), float64(9.223372036854776e+18)) == int32(0) {
		goto L254
	} else {
		goto L258
	}
L258:
	;
	v879 = base.I64_trunc_f64_s(v873)
	v881 = v879
	goto L253
L259:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v854)+24))
	v951 = *(*int64)(unsafe.Add(mBase, uint32(v854)+32))
	v956 = int32(63) - (v950 + base.I32_wrap_i64(base.I64_clz(v951|v943)))
	v957 = base.I64_extend_i32_u(v956)
	v958 = v943 >> (uint(v957) % 64)
	v960 = base.I64_extend32_s(v958) << (uint(v957) % 64)
	if base.F64_eq(v855, float64(0)) != 0 {
		v973 = v960
		goto L276
	} else {
		goto L277
	}
L260:
	;
	v943 = int64(0)
	goto L259
L261:
	;
	v885 = int64(1)
	if v885 < v881 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v888 = v881
	goto L264
L263:
	;
	v888 = v885
	goto L264
L264:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v854)+96))
	v895 = int64(0)
	v899 = int32(0)
	goto L265
L265:
	;
	v904 = *(*int64)(unsafe.Add(mBase, uint32(v889+v899<<(uint(int32(3))%32))))
	v905 = v904 + v895
	if v905 < v888 {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	goto L260
L267:
	;
	v928 = v899 + int32(1)
	if v928 != v882 {
		v895 = v905
		v899 = v928
		goto L265
	} else {
		goto L275
	}
L268:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v854)+28))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v854)+24))
	v913 = int32(base.Ui32(v899) >> (uint(v912) % 32))
	if v913 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v914 = v907
	goto L271
L270:
	;
	v914 = int32(0)
	goto L271
L271:
	;
	v917 = int32(1)
	if v917 < v913 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v920 = v913
	goto L274
L273:
	;
	v920 = v917
	goto L274
L274:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v854)+16))
	v943 = base.I64_extend_i32_s((v907+int32(2147483647))&v899+v914) << (uint(base.I64_extend_i32_u(v920+v921+int32(-1))) % 64)
	goto L259
L275:
	;
	goto L266
L276:
	;
	goto L252
L277:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v854)+40))
	v973 = v960 + int64(1)<<(uint(base.I64_extend_i32_u(v956+base.B2i32(v964 <= base.I32_wrap_i64(v958))))%64) + int64(-1)
	goto L276
L278:
	;
	goto L1
L279:
	;
	v1028 = int32(_a1820)
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v1031 != 0 {
		goto L297
	} else {
		goto L298
	}
L280:
	;
	if v1017-v1019 != 0 {
		goto L279
	} else {
		goto L292
	}
L281:
	;
	v1017 = F_tolower(m, v1013)
	mBase = m.M
	v1018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014))))
	v1019 = F_tolower(m, v1018)
	mBase = m.M
	goto L280
L282:
	;
	v987 = l1
	v988 = v982
	v989 = v985
	goto L285
L283:
	;
	v1013 = int32(0)
	v1014 = v982
	goto L281
L284:
	;
	v1013 = v1010 & int32(255)
	v1014 = v1009
	goto L281
L285:
	;
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v988))))
	if v991 == int32(0) {
		v1009 = v988
		v1010 = v989
		goto L284
	} else {
		goto L287
	}
L286:
	;
	v1009 = v1003
	v1010 = int32(0)
	goto L284
L287:
	;
	v995 = v989 & int32(255)
	if v995 == v991 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1002 = int32(1)
	v1003 = v988 + v1002
	v1004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987)+1)))
	if v1004 != 0 {
		v987 = v987 + v1002
		v988 = v1003
		v989 = v1004
		goto L285
	} else {
		goto L291
	}
L289:
	;
	v997 = F_tolower(m, v995)
	mBase = m.M
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v988))))
	v999 = F_tolower(m, v998)
	mBase = m.M
	if v997 == v999 {
		goto L288
	} else {
		goto L290
	}
L290:
	;
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987))))
	v1009 = v988
	v1010 = v1001
	goto L284
L291:
	;
	goto L286
L292:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+208)) = v1021
	v1026 = F_snprintf(m, l2, l3, int32(_a1807), v9+int32(208))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L16
	} else {
		goto L293
	}
L293:
	;
	goto L1
L294:
	;
	v1081 = int32(_a1821)
	v1084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v1084 != 0 {
		goto L315
	} else {
		goto L316
	}
L295:
	;
	if v1063-v1065 != 0 {
		goto L294
	} else {
		goto L307
	}
L296:
	;
	v1063 = F_tolower(m, v1059)
	mBase = m.M
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1060))))
	v1065 = F_tolower(m, v1064)
	mBase = m.M
	goto L295
L297:
	;
	v1033 = l1
	v1034 = v1028
	v1035 = v1031
	goto L300
L298:
	;
	v1059 = int32(0)
	v1060 = v1028
	goto L296
L299:
	;
	v1059 = v1056 & int32(255)
	v1060 = v1055
	goto L296
L300:
	;
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034))))
	if v1037 == int32(0) {
		v1055 = v1034
		v1056 = v1035
		goto L299
	} else {
		goto L302
	}
L301:
	;
	v1055 = v1049
	v1056 = int32(0)
	goto L299
L302:
	;
	v1041 = v1035 & int32(255)
	if v1041 == v1037 {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1048 = int32(1)
	v1049 = v1034 + v1048
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033)+1)))
	if v1050 != 0 {
		v1033 = v1033 + v1048
		v1034 = v1049
		v1035 = v1050
		goto L300
	} else {
		goto L306
	}
L304:
	;
	v1043 = F_tolower(m, v1041)
	mBase = m.M
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034))))
	v1045 = F_tolower(m, v1044)
	mBase = m.M
	if v1043 == v1045 {
		goto L303
	} else {
		goto L305
	}
L305:
	;
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033))))
	v1055 = v1034
	v1056 = v1047
	goto L299
L306:
	;
	goto L301
L307:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1067 != 0 {
		goto L309
	} else {
		goto L310
	}
L308:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+224)) = v1074
	v1079 = F_snprintf(m, l2, l3, int32(_a1816), v9+int32(224))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L16
	} else {
		goto L311
	}
L309:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1074 = base.F64_promote_f32(base.F32_div(base.F32_convert_i32_u(v1069), base.F32_convert_i32_u(v1067)))
	goto L308
L310:
	;
	v1074 = float64(0)
	goto L308
L311:
	;
	goto L1
L312:
	;
	v1248 = int32(_a1822)
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v1251 != 0 {
		goto L356
	} else {
		goto L357
	}
L313:
	;
	if v1116-v1118 != 0 {
		goto L312
	} else {
		goto L325
	}
L314:
	;
	v1116 = F_tolower(m, v1112)
	mBase = m.M
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1113))))
	v1118 = F_tolower(m, v1117)
	mBase = m.M
	goto L313
L315:
	;
	v1086 = l1
	v1087 = v1081
	v1088 = v1084
	goto L318
L316:
	;
	v1112 = int32(0)
	v1113 = v1081
	goto L314
L317:
	;
	v1112 = v1109 & int32(255)
	v1113 = v1108
	goto L314
L318:
	;
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1087))))
	if v1090 == int32(0) {
		v1108 = v1087
		v1109 = v1088
		goto L317
	} else {
		goto L320
	}
L319:
	;
	v1108 = v1102
	v1109 = int32(0)
	goto L317
L320:
	;
	v1094 = v1088 & int32(255)
	if v1094 == v1090 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1101 = int32(1)
	v1102 = v1087 + v1101
	v1103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1086)+1)))
	if v1103 != 0 {
		v1086 = v1086 + v1101
		v1087 = v1102
		v1088 = v1103
		goto L318
	} else {
		goto L324
	}
L322:
	;
	v1096 = F_tolower(m, v1094)
	mBase = m.M
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1087))))
	v1098 = F_tolower(m, v1097)
	mBase = m.M
	if v1096 == v1098 {
		goto L321
	} else {
		goto L323
	}
L323:
	;
	v1100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1086))))
	v1108 = v1087
	v1109 = v1100
	goto L317
L324:
	;
	goto L319
L325:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1121 = float64(99)
	v1129 = float64(100)
	if base.F64_lt(v1121, v1129) != 0 {
		goto L329
	} else {
		goto L330
	}
L326:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+240)) = base.F64_promote_f32(base.F32_convert_i64_s(v1239))
	v1246 = F_snprintf(m, l2, l3, int32(_a1816), v9+int32(240))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L16
	} else {
		goto L352
	}
L327:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+80))
	if v1148 < int32(1) {
		goto L334
	} else {
		goto L335
	}
L328:
	;
	v1147 = int64(-9223372036854775807 - 1)
	goto L327
L329:
	;
	v1132 = v1121
	goto L331
L330:
	;
	v1132 = v1129
	goto L331
L331:
	;
	v1135 = *(*int64)(unsafe.Add(mBase, uint32(v1120)+88))
	v1139 = base.F64_add(base.F64_mul(base.F64_div(v1132, float64(100)), base.F64_convert_i64_s(v1135)), float64(0.5))
	if base.F64_lt(base.F64_abs(v1139), float64(9.223372036854776e+18)) == int32(0) {
		goto L328
	} else {
		goto L332
	}
L332:
	;
	v1145 = base.I64_trunc_f64_s(v1139)
	v1147 = v1145
	goto L327
L333:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+24))
	v1217 = *(*int64)(unsafe.Add(mBase, uint32(v1120)+32))
	v1222 = int32(63) - (v1216 + base.I32_wrap_i64(base.I64_clz(v1217|v1209)))
	v1223 = base.I64_extend_i32_u(v1222)
	v1224 = v1209 >> (uint(v1223) % 64)
	v1226 = base.I64_extend32_s(v1224) << (uint(v1223) % 64)
	if base.F64_eq(v1121, float64(0)) != 0 {
		v1239 = v1226
		goto L350
	} else {
		goto L351
	}
L334:
	;
	v1209 = int64(0)
	goto L333
L335:
	;
	v1151 = int64(1)
	if v1151 < v1147 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v1154 = v1147
	goto L338
L337:
	;
	v1154 = v1151
	goto L338
L338:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+96))
	v1161 = int64(0)
	v1165 = int32(0)
	goto L339
L339:
	;
	v1170 = *(*int64)(unsafe.Add(mBase, uint32(v1155+v1165<<(uint(int32(3))%32))))
	v1171 = v1170 + v1161
	if v1171 < v1154 {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	goto L334
L341:
	;
	v1194 = v1165 + int32(1)
	if v1194 != v1148 {
		v1161 = v1171
		v1165 = v1194
		goto L339
	} else {
		goto L349
	}
L342:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+28))
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+24))
	v1179 = int32(base.Ui32(v1165) >> (uint(v1178) % 32))
	if v1179 != 0 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1180 = v1173
	goto L345
L344:
	;
	v1180 = int32(0)
	goto L345
L345:
	;
	v1183 = int32(1)
	if v1183 < v1179 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1186 = v1179
	goto L348
L347:
	;
	v1186 = v1183
	goto L348
L348:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+16))
	v1209 = base.I64_extend_i32_s((v1173+int32(2147483647))&v1165+v1180) << (uint(base.I64_extend_i32_u(v1186+v1187+int32(-1))) % 64)
	goto L333
L349:
	;
	goto L340
L350:
	;
	goto L326
L351:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+40))
	v1239 = v1226 + int64(1)<<(uint(base.I64_extend_i32_u(v1222+base.B2i32(v1230 <= base.I32_wrap_i64(v1224))))%64) + int64(-1)
	goto L350
L352:
	;
	goto L1
L353:
	;
	v1415 = int32(_a1823)
	v1418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v1418 != 0 {
		goto L396
	} else {
		goto L397
	}
L354:
	;
	if v1283-v1285 != 0 {
		goto L353
	} else {
		goto L366
	}
L355:
	;
	v1283 = F_tolower(m, v1279)
	mBase = m.M
	v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280))))
	v1285 = F_tolower(m, v1284)
	mBase = m.M
	goto L354
L356:
	;
	v1253 = l1
	v1254 = v1248
	v1255 = v1251
	goto L359
L357:
	;
	v1279 = int32(0)
	v1280 = v1248
	goto L355
L358:
	;
	v1279 = v1276 & int32(255)
	v1280 = v1275
	goto L355
L359:
	;
	v1257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1254))))
	if v1257 == int32(0) {
		v1275 = v1254
		v1276 = v1255
		goto L358
	} else {
		goto L361
	}
L360:
	;
	v1275 = v1269
	v1276 = int32(0)
	goto L358
L361:
	;
	v1261 = v1255 & int32(255)
	if v1261 == v1257 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1268 = int32(1)
	v1269 = v1254 + v1268
	v1270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253)+1)))
	if v1270 != 0 {
		v1253 = v1253 + v1268
		v1254 = v1269
		v1255 = v1270
		goto L359
	} else {
		goto L365
	}
L363:
	;
	v1263 = F_tolower(m, v1261)
	mBase = m.M
	v1264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1254))))
	v1265 = F_tolower(m, v1264)
	mBase = m.M
	if v1263 == v1265 {
		goto L362
	} else {
		goto L364
	}
L364:
	;
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253))))
	v1275 = v1254
	v1276 = v1267
	goto L358
L365:
	;
	goto L360
L366:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1288 = float64(90)
	v1296 = float64(100)
	if base.F64_lt(v1288, v1296) != 0 {
		goto L370
	} else {
		goto L371
	}
L367:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+256)) = base.F64_promote_f32(base.F32_convert_i64_s(v1406))
	v1413 = F_snprintf(m, l2, l3, int32(_a1816), v9+int32(256))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L16
	} else {
		goto L393
	}
L368:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+80))
	if v1315 < int32(1) {
		goto L375
	} else {
		goto L376
	}
L369:
	;
	v1314 = int64(-9223372036854775807 - 1)
	goto L368
L370:
	;
	v1299 = v1288
	goto L372
L371:
	;
	v1299 = v1296
	goto L372
L372:
	;
	v1302 = *(*int64)(unsafe.Add(mBase, uint32(v1287)+88))
	v1306 = base.F64_add(base.F64_mul(base.F64_div(v1299, float64(100)), base.F64_convert_i64_s(v1302)), float64(0.5))
	if base.F64_lt(base.F64_abs(v1306), float64(9.223372036854776e+18)) == int32(0) {
		goto L369
	} else {
		goto L373
	}
L373:
	;
	v1312 = base.I64_trunc_f64_s(v1306)
	v1314 = v1312
	goto L368
L374:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+24))
	v1384 = *(*int64)(unsafe.Add(mBase, uint32(v1287)+32))
	v1389 = int32(63) - (v1383 + base.I32_wrap_i64(base.I64_clz(v1384|v1376)))
	v1390 = base.I64_extend_i32_u(v1389)
	v1391 = v1376 >> (uint(v1390) % 64)
	v1393 = base.I64_extend32_s(v1391) << (uint(v1390) % 64)
	if base.F64_eq(v1288, float64(0)) != 0 {
		v1406 = v1393
		goto L391
	} else {
		goto L392
	}
L375:
	;
	v1376 = int64(0)
	goto L374
L376:
	;
	v1318 = int64(1)
	if v1318 < v1314 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1321 = v1314
	goto L379
L378:
	;
	v1321 = v1318
	goto L379
L379:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+96))
	v1328 = int64(0)
	v1332 = int32(0)
	goto L380
L380:
	;
	v1337 = *(*int64)(unsafe.Add(mBase, uint32(v1322+v1332<<(uint(int32(3))%32))))
	v1338 = v1337 + v1328
	if v1338 < v1321 {
		goto L382
	} else {
		goto L383
	}
L381:
	;
	goto L375
L382:
	;
	v1361 = v1332 + int32(1)
	if v1361 != v1315 {
		v1328 = v1338
		v1332 = v1361
		goto L380
	} else {
		goto L390
	}
L383:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+28))
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+24))
	v1346 = int32(base.Ui32(v1332) >> (uint(v1345) % 32))
	if v1346 != 0 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1347 = v1340
	goto L386
L385:
	;
	v1347 = int32(0)
	goto L386
L386:
	;
	v1350 = int32(1)
	if v1350 < v1346 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1353 = v1346
	goto L389
L388:
	;
	v1353 = v1350
	goto L389
L389:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+16))
	v1376 = base.I64_extend_i32_s((v1340+int32(2147483647))&v1332+v1347) << (uint(base.I64_extend_i32_u(v1353+v1354+int32(-1))) % 64)
	goto L374
L390:
	;
	goto L381
L391:
	;
	goto L367
L392:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+40))
	v1406 = v1393 + int64(1)<<(uint(base.I64_extend_i32_u(v1389+base.B2i32(v1397 <= base.I32_wrap_i64(v1391))))%64) + int64(-1)
	goto L391
L393:
	;
	goto L1
L394:
	;
	if v1450-v1452 != 0 {
		goto L1
	} else {
		goto L406
	}
L395:
	;
	v1450 = F_tolower(m, v1446)
	mBase = m.M
	v1451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1447))))
	v1452 = F_tolower(m, v1451)
	mBase = m.M
	goto L394
L396:
	;
	v1420 = l1
	v1421 = v1415
	v1422 = v1418
	goto L399
L397:
	;
	v1446 = int32(0)
	v1447 = v1415
	goto L395
L398:
	;
	v1446 = v1443 & int32(255)
	v1447 = v1442
	goto L395
L399:
	;
	v1424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1421))))
	if v1424 == int32(0) {
		v1442 = v1421
		v1443 = v1422
		goto L398
	} else {
		goto L401
	}
L400:
	;
	v1442 = v1436
	v1443 = int32(0)
	goto L398
L401:
	;
	v1428 = v1422 & int32(255)
	if v1428 == v1424 {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1435 = int32(1)
	v1436 = v1421 + v1435
	v1437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1420)+1)))
	if v1437 != 0 {
		v1420 = v1420 + v1435
		v1421 = v1436
		v1422 = v1437
		goto L399
	} else {
		goto L405
	}
L403:
	;
	v1430 = F_tolower(m, v1428)
	mBase = m.M
	v1431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1421))))
	v1432 = F_tolower(m, v1431)
	mBase = m.M
	if v1430 == v1432 {
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v1434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1420))))
	v1442 = v1421
	v1443 = v1434
	goto L398
L405:
	;
	goto L400
L406:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1455 = float64(50)
	v1463 = float64(100)
	if base.F64_lt(v1455, v1463) != 0 {
		goto L410
	} else {
		goto L411
	}
L407:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+272)) = base.F64_promote_f32(base.F32_convert_i64_s(v1573))
	v1580 = F_snprintf(m, l2, l3, int32(_a1816), v9+int32(272))
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L16
	} else {
		goto L433
	}
L408:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+80))
	if v1482 < int32(1) {
		goto L415
	} else {
		goto L416
	}
L409:
	;
	v1481 = int64(-9223372036854775807 - 1)
	goto L408
L410:
	;
	v1466 = v1455
	goto L412
L411:
	;
	v1466 = v1463
	goto L412
L412:
	;
	v1469 = *(*int64)(unsafe.Add(mBase, uint32(v1454)+88))
	v1473 = base.F64_add(base.F64_mul(base.F64_div(v1466, float64(100)), base.F64_convert_i64_s(v1469)), float64(0.5))
	if base.F64_lt(base.F64_abs(v1473), float64(9.223372036854776e+18)) == int32(0) {
		goto L409
	} else {
		goto L413
	}
L413:
	;
	v1479 = base.I64_trunc_f64_s(v1473)
	v1481 = v1479
	goto L408
L414:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+24))
	v1551 = *(*int64)(unsafe.Add(mBase, uint32(v1454)+32))
	v1556 = int32(63) - (v1550 + base.I32_wrap_i64(base.I64_clz(v1551|v1543)))
	v1557 = base.I64_extend_i32_u(v1556)
	v1558 = v1543 >> (uint(v1557) % 64)
	v1560 = base.I64_extend32_s(v1558) << (uint(v1557) % 64)
	if base.F64_eq(v1455, float64(0)) != 0 {
		v1573 = v1560
		goto L431
	} else {
		goto L432
	}
L415:
	;
	v1543 = int64(0)
	goto L414
L416:
	;
	v1485 = int64(1)
	if v1485 < v1481 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v1488 = v1481
	goto L419
L418:
	;
	v1488 = v1485
	goto L419
L419:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+96))
	v1495 = int64(0)
	v1499 = int32(0)
	goto L420
L420:
	;
	v1504 = *(*int64)(unsafe.Add(mBase, uint32(v1489+v1499<<(uint(int32(3))%32))))
	v1505 = v1504 + v1495
	if v1505 < v1488 {
		goto L422
	} else {
		goto L423
	}
L421:
	;
	goto L415
L422:
	;
	v1528 = v1499 + int32(1)
	if v1528 != v1482 {
		v1495 = v1505
		v1499 = v1528
		goto L420
	} else {
		goto L430
	}
L423:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+28))
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+24))
	v1513 = int32(base.Ui32(v1499) >> (uint(v1512) % 32))
	if v1513 != 0 {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v1514 = v1507
	goto L426
L425:
	;
	v1514 = int32(0)
	goto L426
L426:
	;
	v1517 = int32(1)
	if v1517 < v1513 {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v1520 = v1513
	goto L429
L428:
	;
	v1520 = v1517
	goto L429
L429:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+16))
	v1543 = base.I64_extend_i32_s((v1507+int32(2147483647))&v1499+v1514) << (uint(base.I64_extend_i32_u(v1520+v1521+int32(-1))) % 64)
	goto L414
L430:
	;
	goto L421
L431:
	;
	goto L407
L432:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+40))
	v1573 = v1560 + int64(1)<<(uint(base.I64_extend_i32_u(v1556+base.B2i32(v1564 <= base.I32_wrap_i64(v1558))))%64) + int64(-1)
	goto L431
L433:
	;
	goto L1
}
func F_rdbWriteRaw(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v55
L2:
	;
	v55 = l2
	goto L1
L3:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v10&int32(6) != 0 {
		v55 = int32(-1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v16 = l1
	v19 = l2
	goto L7
L6:
	;
	return int32(0)
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v21) < base.Ui32(v19) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v23 = v21
	goto L11
L10:
	;
	v23 = v19
	goto L11
L11:
	;
	if v21 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v24 = v23
	goto L14
L13:
	;
	v24 = v19
	goto L14
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v25 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = m.T0[v32].(func(*base.Module, int32, int32, int32) int32)(m, l0, v16, v24)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L17
	} else {
		goto L20
	}
L16:
	;
	m.T0[v25].(func(*base.Module, int32, int32, int32))(m, l0, v16, v24)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	goto L15
L19:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v41 + v24
	v45 = v19 - v24
	if v45 != 0 {
		v16 = v16 + v24
		v19 = v45
		goto L7
	} else {
		goto L22
	}
L20:
	;
	if v33 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v35 | int64(2)
	return int32(-1)
L22:
	;
	goto L8
}
func F_tryExpandRdbStats(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	if base.Ui32(l2) <= base.Ui32(l1) {
		v47 = l0
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1802), int32(_a1803), int32(233))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L15
	}
L2:
	;
	return v47
L3:
	;
	v8 = F_valkey_realloc(m, l0, l2<<(uint(int32(2))%32))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v8 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v15 = l1
	goto L7
L7:
	;
	v19 = F_valkey_calloc(m, int32(48))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	v47 = v8
	goto L2
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+v15<<(uint(int32(2))%32)))) = v19
	v45 = v15 + int32(1)
	if v45 != l2 {
		v15 = v45
		goto L7
	} else {
		goto L14
	}
L10:
	;
	if v19 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v24 = base.I32_rem_u_s(v15, int32(7))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v24
	v31 = F_hdr_init(m, int64(1), int64(204800), int32(3), v19+int32(40))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v38 = F_hdr_init(m, int64(1), int64(1048576), int32(3), v19+int32(44))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	goto L8
L15:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
