package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_extractLongLatOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(-1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = F_getDoubleFromObjectOrReply(m, l0, v13, l2, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 != 0 {
			v43 = v12
			m.G0 = v10 + int32(16)
			return v43
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v23 = F_getDoubleFromObjectOrReply(m, l0, v19, l2+int32(8), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 != 0 {
					v43 = v12
					m.G0 = v10 + int32(16)
					return v43
				} else {
					v25 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
					v26 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
					if base.F64_gt(base.F64_abs(v26), float64(180)) != 0 {
						*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v25
						*(*float64)(unsafe.Add(mBase, uint32(v10))) = v26
						F_addReplyErrorFormat(m, l0, int32(_a820), v10)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v43 = int32(-1)
							m.G0 = v10 + int32(16)
							return v43
						}
					} else {
						v30 = int32(0)
						if base.F64_gt(base.F64_abs(v25), float64(85.05112878)) == v30 {
							v43 = v30
							m.G0 = v10 + int32(16)
							return v43
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v25
							*(*float64)(unsafe.Add(mBase, uint32(v10))) = v26
							F_addReplyErrorFormat(m, l0, int32(_a820), v10)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v43 = int32(-1)
								m.G0 = v10 + int32(16)
								return v43
							}
						}
					}
				}
			}
		}
	}
}
func F_getLongDoubleFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v33 int32
	_ = v33
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_getLongDoubleFromObject(m, l1, v9)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v26 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
			v29 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(8))))
			*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v29
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v26
			v33 = int32(0)
			m.G0 = v9 + int32(16)
			return v33
		} else {
			if l3 == int32(0) {
				F_addReplyError(m, l0, int32(_a1060))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v33 = int32(-1)
					m.G0 = v9 + int32(16)
					return v33
				}
			} else {
				F_addReplyError(m, l0, l3)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v33 = int32(-1)
					m.G0 = v9 + int32(16)
					return v33
				}
			}
		}
	}
}
func F_getLongFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_getLongLongFromObject(m, l1, v9+int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
			if base.Ui64(int64(-4294967297)) < base.Ui64(v24+int64(-2147483648)) {
				*(*uint32)(unsafe.Add(mBase, uint32(l2))) = uint32(v24)
				v40 = int32(0)
				m.G0 = v9 + int32(16)
				return v40
			} else {
				if l3 == int32(0) {
					F_addReplyError(m, l0, int32(_a1061))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v40 = int32(-1)
						m.G0 = v9 + int32(16)
						return v40
					}
				} else {
					F_addReplyError(m, l0, l3)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v40 = int32(-1)
						m.G0 = v9 + int32(16)
						return v40
					}
				}
			}
		} else {
			if l3 != 0 {
				v20 = l3
			} else {
				v20 = int32(_a1062)
			}
			F_addReplyError(m, l0, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v40 = int32(-1)
				m.G0 = v9 + int32(16)
				return v40
			}
		}
	}
}
func F_longLatFromMemberOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 float64
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v43 int64
	_ = v43
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
	v15 = F_objectGetVal(m, l2)
	mBase = m.M
	v18 = F_zsetScore(m, l1, v15, v11+int32(24))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v18 != int32(-1) {
			v25 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
			v29 = v11 + int32(40)
			v30 = int32(26)
			*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v30)
			v32 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11+int32(44)))) = v32
			*(*int32)(unsafe.Add(mBase, uint32(v11)+41)) = v32
			v43 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
			*(*int64)(unsafe.Add(mBase, uint32(v11+int32(16)))) = v43
			if base.F64_lt(v25, float64(1.8446744073709552e+19))&base.F64_ge(v25, float64(0)) == v32 {
				v54 = int64(0)
			} else {
				v52 = base.I64_trunc_f64_u(v25)
				v54 = v52
			}
			*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v54
			*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v54
			v57 = int32(8)
			v60 = m.G0
			v61 = int32(16)
			v62 = v60 - v61
			m.G0 = v62
			v68 = *(*int64)(unsafe.Add(mBase, uint32(v11+v61)))
			*(*int64)(unsafe.Add(mBase, uint32(v62+v57))) = v68
			v70 = *(*int64)(unsafe.Add(mBase, uint32(v11+v57)))
			*(*int64)(unsafe.Add(mBase, uint32(v62))) = v70
			v72 = F_geohashDecodeToLongLatType(m, v62, l3)
			mBase = m.M
			m.G0 = v62 + v61
			if v72 != 0 {
				v86 = v32
				m.G0 = v11 + int32(48)
				return v86
			} else {
				v77 = int32(_a821)
				v81 = F_objectGetVal(m, l2)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v81
				F_addReplyErrorFormat(m, l0, v77, v11)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					v86 = int32(-1)
					m.G0 = v11 + int32(48)
					return v86
				}
			}
		} else {
			v77 = int32(_a822)
			v81 = F_objectGetVal(m, l2)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v81
			F_addReplyErrorFormat(m, l0, v77, v11)
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int32(0)
			} else {
				v86 = int32(-1)
				m.G0 = v11 + int32(48)
				return v86
			}
		}
	}
}
func F_readLong(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v101 int32
	_ = v101
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(1216)
	m.G0 = v8
	v12 = F___ftello(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, _consts[958])) = v12
		v20 = F_fgets(m, v8+int32(64), int32(128), l0)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v20 == int32(0) {
				v101 = v4
				m.G0 = v8 + int32(1216)
				return v101
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+64)))
				if v24 == l1&int32(255) {
					v61 = F_strtox_2(m, v8+int32(64)|int32(1), v8+int32(60), int32(10), int64(2147483648))
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = base.I32_wrap_i64(v61)
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
					v65 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64)+1)))
					v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64))))
					if v66 != int32(13) {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v66
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v65
						v79 = F_snprintf(m, v8+int32(192), int32(1024), int32(_a1770), v8+int32(16))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v81 = int32(0)
							v83 = *(*int64)(unsafe.Add(mBase, _consts[958]))
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = v83
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v8 + int32(192)
							v91 = F_snprintf(m, int32(_a1771), int32(1044), int32(_a1772), v8)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								v101 = v81
								m.G0 = v8 + int32(1216)
								return v101
							}
						}
					} else {
						if v65 == int32(10) {
							v93 = int32(0)
							v95 = *(*int64)(unsafe.Add(mBase, _consts[959]))
							*(*int64)(unsafe.Add(mBase, _consts[959])) = v95 + int64(1)
							v101 = int32(1)
							m.G0 = v8 + int32(1216)
							return v101
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v66
							*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v65
							v79 = F_snprintf(m, v8+int32(192), int32(1024), int32(_a1770), v8+int32(16))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v81 = int32(0)
								v83 = *(*int64)(unsafe.Add(mBase, _consts[958]))
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v83
								*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v8 + int32(192)
								v91 = F_snprintf(m, int32(_a1771), int32(1044), int32(_a1772), v8)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v101 = v81
									m.G0 = v8 + int32(1216)
									return v101
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = base.I32_extend8_s(v24)
					v37 = F_snprintf(m, v8+int32(192), int32(1024), int32(_a1773), v8+int32(48))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = int32(0)
						v41 = *(*int64)(unsafe.Add(mBase, _consts[958]))
						*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v41
						*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v8 + int32(192)
						v51 = F_snprintf(m, int32(_a1771), int32(1044), int32(_a1772), v8+int32(32))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v101 = v39
							m.G0 = v8 + int32(1216)
							return v101
						}
					}
				}
			}
		}
	}
}
func F_read_long_string(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v17 int32
	_ = v17
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
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
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_save(m, l0, v13)
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v17 + int32(-1)
	if v17 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v30
	switch v30 + int32(-10) {
	case 0, 3:
		goto L8
	default:
		goto L7
	}
L4:
	;
	v28 = F_luaZ_fill(m, v16)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v23 + int32(1)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v30 = v27
	goto L3
L6:
	;
	v30 = v28
	goto L3
L7:
	;
	v37 = m.G3
	if l1 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	F_inclinenumber(m, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v40 = int32(_a2032)
	goto L12
L11:
	;
	v40 = int32(_a2033)
	goto L12
L12:
	;
	goto L13
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v50 + int32(1) {
	case 0:
		goto L19
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13:
		goto L15
	case 11, 14:
		goto L16
	default:
		goto L20
	}
L15:
	;
	if l1 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L16:
	;
	F_save(m, l0, int32(10))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L57
	}
L17:
	;
	v186 = F_skip_sep(m, l0)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L42
	}
L18:
	;
	v164 = F_skip_sep(m, l0)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L36
	}
L19:
	;
	v56 = v11 + int32(32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v59 = v57 + int32(16)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	switch v63 + int32(-61) {
	case 0:
		goto L24
	default:
		goto L22
	case 3:
		goto L23
	}
L20:
	;
	switch v50 + int32(-91) {
	case 0:
		goto L18
	default:
		goto L15
	case 2:
		goto L17
	}
L21:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v37 + v40
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v138
	v141 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v11 + int32(32)
	v149 = F_luaO_pushfstring(m, v137, v141+int32(_a2026), v11+int32(16))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L33
	}
L22:
	;
	v90 = m.G3
	v93 = F_strcspn(m, v59, v90+int32(_a2027))
	mBase = m.M
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1013]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(40)))) = uint16(v100)
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1014])))
	*(*int64)(unsafe.Add(mBase, uint32(v56))) = v102
	v105 = int32(63)
	if base.Ui32(v93) < base.Ui32(v105) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	v75 = v57 + int32(17)
	v76 = F_strlen(m, v75)
	mBase = m.M
	v77 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v77)
	v80 = int32(72)
	if base.Ui32(v76) <= base.Ui32(v80) {
		v88 = v75
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v68 = F_strncpy(m, v56, v57+int32(17), int32(80))
	mBase = m.M
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v68+int32(79)))) = uint8(v72)
	goto L21
L25:
	;
	v89 = F_strcat(m, v56, v88)
	mBase = m.M
	goto L21
L26:
	;
	v82 = F_strlen(m, v56)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v56+v82))) = int32(3026478)
	v88 = v75 + (v76 - v80)
	goto L25
L27:
	;
	v119 = F_strlen(m, v56)
	mBase = m.M
	v120 = v56 + v119
	v121 = m.G3
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[1015]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v120))) = uint16(v124)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[1016]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v120+int32(2)))) = uint8(v130)
	goto L21
L28:
	;
	v117 = F_strcat(m, v56, v59)
	mBase = m.M
	goto L27
L29:
	;
	v107 = v93
	goto L31
L30:
	;
	v107 = v105
	goto L31
L31:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v107))))
	if v109 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v112 = F_strncat(m, v56, v59, v107)
	mBase = m.M
	v113 = F_strlen(m, v112)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v112+v113))) = int32(3026478)
	goto L27
L33:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v141 + int32(_a2028)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v149
	v158 = F_luaO_pushfstring(m, v151, v141+int32(_a2029), v11)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_luaD_throw(m, v160, int32(3))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L13
L36:
	;
	if v164 != l2 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_save(m, l0, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v171 + int32(-1)
	if v171 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v183 = F_luaZ_fill(m, v170)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = v177 + int32(1)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v181
	goto L13
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v183
	goto L13
L42:
	;
	if v186 != l2 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_save(m, l0, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v193 + int32(-1)
	if v193 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v206
	if l1 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v204 = F_luaZ_fill(m, v192)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = v199 + int32(1)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	v206 = v203
	goto L45
L48:
	;
	v206 = v204
	goto L45
L49:
	;
	m.G0 = v11 + int32(112)
	return
L50:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	v219 = F_luaS_newlstr(m, v211, v213+l2, v215-l2<<(uint(int32(1))%32))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	v223 = F_luaH_setstr(m, v211, v222, v219)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v219
	goto L49
L53:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	if v225 != 0 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v226 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+8)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v226
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v211)+16))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+68))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v230)+64))
	if base.Ui32(v231) < base.Ui32(v232) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	F_luaC_step(m, v211)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L52
L57:
	;
	F_inclinenumber(m, l0)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if l1 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v249)+4)) = int32(0)
	goto L13
L60:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = v273 + int32(-1)
	if v273 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L61:
	;
	F_save(m, l0, v50)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v257 + int32(-1)
	if v257 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v269 = F_luaZ_fill(m, v256)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+4)) = v263 + int32(1)
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v267
	goto L13
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v269
	goto L13
L66:
	;
	v285 = F_luaZ_fill(m, v272)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v272)+4)) = v279 + int32(1)
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v283
	goto L13
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v285
	goto L13
}
