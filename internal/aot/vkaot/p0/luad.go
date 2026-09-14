package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaD_growstack(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v4 < l1 {
		v9 = v4 + l1
	} else {
		v9 = v4 << (uint(int32(1)) % 32)
	}
	F_luaD_reallocstack(m, l0, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		return
	}
}
func F_luaD_pcall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = l4
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
	v17 = F_luaD_rawrunprotected(m, l0, l1, l2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v11
			return v17
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v24 = v23 + l3
			F_luaF_close(m, l0, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = v15 - v14
				switch v17 + int32(-2) {
				case 0, 1:
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v49 = *(*int64)(unsafe.Add(mBase, uint32(v46+int32(-16))))
					*(*int64)(unsafe.Add(mBase, uint32(v24))) = v49
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(-8))))
					v54 = v53
					*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v54
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v16)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24 + int32(16)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v62 = v61 + v27
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v62
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v64
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v66
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v13)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if v69 < int32(20001) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v11
						return v17
					} else {
						if int32(479975) < v27 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v11
							return v17
						} else {
							v77 = F_luaM_realloc_(m, l0, v61, v69*int32(24), int32(480000))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(20000)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v77
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v77 + int32(479976)
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v77 + (v85 - v61)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v11
								return v17
							}
						}
					}
				case 2:
					v30 = m.G3
					v34 = F_luaS_newlstr(m, l0, v30+int32(_a2256), int32(17))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v24))) = v34
						v54 = int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v54
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v16)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24 + int32(16)
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v62 = v61 + v27
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v62
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v64
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v66
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v13)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						if v69 < int32(20001) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v11
							return v17
						} else {
							if int32(479975) < v27 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v11
								return v17
							} else {
								v77 = F_luaM_realloc_(m, l0, v61, v69*int32(24), int32(480000))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(20000)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v77
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v77 + int32(479976)
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v77 + (v85 - v61)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v11
									return v17
								}
							}
						}
					}
				case 3:
					v38 = m.G3
					v42 = F_luaS_newlstr(m, l0, v38+int32(_a2257), int32(23))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v24))) = v42
						v54 = int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v54
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v16)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24 + int32(16)
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v62 = v61 + v27
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v62
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v64
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v66
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v13)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						if v69 < int32(20001) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v11
							return v17
						} else {
							if int32(479975) < v27 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v11
								return v17
							} else {
								v77 = F_luaM_realloc_(m, l0, v61, v69*int32(24), int32(480000))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(20000)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v77
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v77 + int32(479976)
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v77 + (v85 - v61)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v11
									return v17
								}
							}
						}
					}
				default:
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v16)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24 + int32(16)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v62 = v61 + v27
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v62
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v64
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v66
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v13)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if v69 < int32(20001) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v11
						return v17
					} else {
						if int32(479975) < v27 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v11
							return v17
						} else {
							v77 = F_luaM_realloc_(m, l0, v61, v69*int32(24), int32(480000))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(20000)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v77
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v77 + int32(479976)
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v77 + (v85 - v61)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v11
								return v17
							}
						}
					}
				}
			}
		}
	}
}
func F_luaD_poscall(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
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
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int64
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v292 int32
	_ = v292
	var v304 int32
	_ = v304
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v15&int32(2) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v189 = v181 + int32(-24)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v193
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v181+int32(-12))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v197
	if v191 == int32(0) {
		v304 = v192
		goto L32
	} else {
		goto L33
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v20 == int32(0) {
		v77 = v19
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v179 = l1
	v181 = v18
	goto L1
L4:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+6)))
	if v84 != 0 {
		v169 = v81
		v171 = v77
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)))
	if v23 == int32(0) {
		v77 = v19
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v36 = base.I32_div_s(v26-v33, int32(24))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(320) < v40-v28 {
		v58 = v28
		v59 = v26
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v58 + int32(320)
	v63 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v63)
	m.T0[v20].(func(*base.Module, int32, int32))(m, l0, v13+int32(12))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L14
	}
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v45 = int32(20)
	if v44 < v45 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v51 = v44 + v45
	goto L11
L10:
	;
	v51 = v44 << (uint(int32(1)) % 32)
	goto L11
L11:
	;
	F_luaD_reallocstack(m, l0, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v58 = v57
	v59 = v56
	goto L7
L14:
	;
	v69 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v69)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v72 + (v27 - v19)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72 + (v28 - v19)
	v77 = v72
	goto L4
L15:
	;
	v179 = v171 + (l1 - v19)
	v181 = v169
	goto L1
L16:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v85&int32(2) == int32(0) {
		v169 = v81
		v171 = v77
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v93 = v81
	v95 = v77
	v97 = v85
	goto L18
L18:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = v100 + int32(-1)
	if v100 == int32(0) {
		v169 = v93
		v171 = v95
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v169 = v159
	v171 = v160
	goto L15
L20:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v106 == int32(0) {
		v159 = v93
		v160 = v95
		v162 = v97
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v162&int32(2) != 0 {
		v93 = v159
		v95 = v160
		v97 = v162
		goto L18
	} else {
		goto L31
	}
L22:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)))
	if v109 == int32(0) {
		v159 = v93
		v160 = v95
		v162 = v97
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(4)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(320) < v122-v112 {
		v138 = v93
		v140 = v112
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+8)) = v140 + int32(320)
	v144 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v144)
	m.T0[v106].(func(*base.Module, int32, int32))(m, l0, v13+int32(12))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L12
	} else {
		goto L30
	}
L25:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v127 = int32(20)
	if v126 < v127 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v133 = v126 + v127
	goto L28
L27:
	;
	v133 = v126 << (uint(int32(1)) % 32)
	goto L28
L28:
	;
	F_luaD_reallocstack(m, l0, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v138 = v136
	v140 = v137
	goto L24
L30:
	;
	v150 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v150)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v152)+8)) = v153 + (v113 - v95)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v153 + (v112 - v95)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	v159 = v152
	v160 = v153
	v162 = v158
	goto L21
L31:
	;
	goto L19
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v304
	m.G0 = v13 + int32(112)
	return v191 + int32(1)
L33:
	;
	v203 = v179
	v205 = v191
	v207 = v192
	v208 = int32(0)
	goto L35
L34:
	;
	if v205 < int32(1) {
		v304 = v207
		goto L32
	} else {
		goto L39
	}
L35:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v212) <= base.Ui32(v203) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v214 = *(*int64)(unsafe.Add(mBase, uint32(v203)))
	*(*int64)(unsafe.Add(mBase, uint32(v207))) = v214
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+8)) = v216
	v220 = int32(16)
	v221 = v207 + v220
	v225 = v205 + int32(-1)
	if v225 != 0 {
		v203 = v203 + v220
		v205 = v225
		v207 = v221
		v208 = v208 + int32(1)
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v304 = v221
	goto L32
L39:
	;
	v229 = v205 & int32(7)
	if v229 == int32(0) {
		v255 = v205
		v257 = v207
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if base.Ui32(int32(-8)) < base.Ui32(v208-v191) {
		v304 = v257
		goto L32
	} else {
		goto L45
	}
L41:
	;
	v234 = int32(0)
	v236 = v205
	v238 = v207
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+8)) = int32(0)
	v246 = v238 + int32(16)
	v248 = v236 + int32(-1)
	v250 = v234 + int32(1)
	if v250 != v229 {
		v234 = v250
		v236 = v248
		v238 = v246
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v255 = v248
	v257 = v246
	goto L40
L44:
	;
	goto L43
L45:
	;
	v268 = v255
	v270 = v257
	goto L46
L46:
	;
	v275 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v270)+120)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v270)+104)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v270)+88)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v270)+72)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v270)+56)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v270)+40)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v270)+24)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v270)+8)) = v275
	v292 = v270 + int32(128)
	if base.Ui32(v268+int32(-9)) < base.Ui32(int32(-2)) {
		v268 = v268 + int32(-8)
		v270 = v292
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v304 = v292
	goto L32
L48:
	;
	goto L47
}
func F_luaD_reallocCI(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(int32(178956970)) < base.Ui32(l1+int32(1)) {
		v17 = F_luaM_toobig(m, l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = v17
			*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v19
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v19 + (v22 - v5)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v19 + l1*int32(24) + int32(-24)
			return
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v11 = int32(24)
		v15 = F_luaM_realloc_(m, l0, v5, v10*v11, l1*v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v19 = v15
			*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v19
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v19 + (v22 - v5)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v19 + l1*int32(24) + int32(-24)
			return
		}
	}
}
func F_luaD_reallocstack(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	v7 = l1 + int32(6)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(int32(268435455)) < base.Ui32(l1+int32(7)) {
		v20 = F_luaM_toobig(m, l0)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v22 = v20
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v7
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v22 + l1<<(uint(int32(4))%32)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v22 + (v29 - v8)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v33 == int32(0) {
				v52 = v22
			} else {
				v37 = v33
				for {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v41 + (v42 - v8)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
					if v46 != 0 {
						v37 = v46
						continue
					} else {
						break
					}
					break
				}
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v52 = v47
			}
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if base.Ui32(v54) < base.Ui32(v53) {
			} else {
				v57 = v53
				for {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v52 + (v61 - v8)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					*(*int32)(unsafe.Add(mBase, uint32(v57))) = v52 + (v65 - v8)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v52 + (v69 - v8)
					v74 = v57 + int32(24)
					if base.Ui32(v74) <= base.Ui32(v54) {
						v57 = v74
						continue
					} else {
						break
					}
					break
				}
			}
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v52 + (v81 - v8)
			return
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v14 = int32(4)
		v18 = F_luaM_realloc_(m, l0, v8, v13<<(uint(v14)%32), v7<<(uint(v14)%32))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v22 = v18
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v7
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v22 + l1<<(uint(int32(4))%32)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v22 + (v29 - v8)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v33 == int32(0) {
				v52 = v22
			} else {
				v37 = v33
				for {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v41 + (v42 - v8)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
					if v46 != 0 {
						v37 = v46
						continue
					} else {
						break
					}
					break
				}
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v52 = v47
			}
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if base.Ui32(v54) < base.Ui32(v53) {
			} else {
				v57 = v53
				for {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v52 + (v61 - v8)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					*(*int32)(unsafe.Add(mBase, uint32(v57))) = v52 + (v65 - v8)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v52 + (v69 - v8)
					v74 = v57 + int32(24)
					if base.Ui32(v74) <= base.Ui32(v54) {
						v57 = v74
						continue
					} else {
						break
					}
					break
				}
			}
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v52 + (v81 - v8)
			return
		}
	}
}
func F_luaD_throw(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	v2 = l1
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v4 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+160)) = v2
		F___wasm_longjmp(m, v4+int32(4), int32(1))
		mBase = m.M
		v93 = m.ExcPending
		if v93 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v2)
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+88))
		if v7 == int32(0) {
			m.Env.Exit(m, int32(1))
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v10
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
			F_luaF_close(m, l0, v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				switch v2 + int32(-2) {
				case 0, 1:
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v38 = *(*int64)(unsafe.Add(mBase, uint32(v35+int32(-16))))
					*(*int64)(unsafe.Add(mBase, uint32(v16))) = v38
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(-8))))
					v43 = v42
					*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v43
					v46 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v46)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v16 + int32(16)
					v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v51)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if v53 < int32(20001) {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
						v82 = m.T0[v81].(func(*base.Module, int32) int32)(m, l0)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							m.Env.Exit(m, int32(1))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if int32(479975) < v56-v57 {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
							v82 = m.T0[v81].(func(*base.Module, int32) int32)(m, l0)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								m.Env.Exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v64 = F_luaM_realloc_(m, l0, v57, v53*int32(24), int32(480000))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(20000)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v64 + int32(479976)
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v64 + (v72 - v57)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
								v82 = m.T0[v81].(func(*base.Module, int32) int32)(m, l0)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									m.Env.Exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				case 2:
					v19 = m.G3
					v23 = F_luaS_newlstr(m, l0, v19+int32(_a2256), int32(17))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v23
						v43 = int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v43
						v46 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v46)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v16 + int32(16)
						v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v51)
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						if v53 < int32(20001) {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
							v82 = m.T0[v81].(func(*base.Module, int32) int32)(m, l0)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								m.Env.Exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if int32(479975) < v56-v57 {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
								v82 = m.T0[v81].(func(*base.Module, int32) int32)(m, l0)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									m.Env.Exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v64 = F_luaM_realloc_(m, l0, v57, v53*int32(24), int32(480000))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(20000)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v64
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v64 + int32(479976)
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v64 + (v72 - v57)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
									v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
									v82 = m.T0[v81].(func(*base.Module, int32) int32)(m, l0)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										m.Env.Exit(m, int32(1))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				case 3:
					v27 = m.G3
					v31 = F_luaS_newlstr(m, l0, v27+int32(_a2257), int32(23))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v31
						v43 = int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v43
						v46 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v46)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v16 + int32(16)
						v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v51)
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						if v53 < int32(20001) {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
							v82 = m.T0[v81].(func(*base.Module, int32) int32)(m, l0)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								m.Env.Exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if int32(479975) < v56-v57 {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
								v82 = m.T0[v81].(func(*base.Module, int32) int32)(m, l0)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									m.Env.Exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v64 = F_luaM_realloc_(m, l0, v57, v53*int32(24), int32(480000))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(20000)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v64
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v64 + int32(479976)
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v64 + (v72 - v57)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
									v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
									v82 = m.T0[v81].(func(*base.Module, int32) int32)(m, l0)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										m.Env.Exit(m, int32(1))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				default:
					v46 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v46)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v16 + int32(16)
					v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v51)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if v53 < int32(20001) {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
						v82 = m.T0[v81].(func(*base.Module, int32) int32)(m, l0)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							m.Env.Exit(m, int32(1))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if int32(479975) < v56-v57 {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
							v82 = m.T0[v81].(func(*base.Module, int32) int32)(m, l0)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								m.Env.Exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v64 = F_luaM_realloc_(m, l0, v57, v53*int32(24), int32(480000))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(20000)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v64 + int32(479976)
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v64 + (v72 - v57)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
								v82 = m.T0[v81].(func(*base.Module, int32) int32)(m, l0)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									m.Env.Exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
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
