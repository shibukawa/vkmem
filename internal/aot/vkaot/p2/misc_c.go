package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F___clock_nanosleep(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(28)
	if l0 == int32(3) {
		v62 = v14
		m.G0 = v12 + int32(16)
		return v62
	} else {
		if l2 == int32(0) {
			v62 = v14
			m.G0 = v12 + int32(16)
			return v62
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			if base.Ui32(int32(999999999)) < base.Ui32(v19) {
				v62 = v14
				m.G0 = v12 + int32(16)
				return v62
			} else {
				v22 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				if v22 < int64(0) {
					v62 = v14
					m.G0 = v12 + int32(16)
					return v62
				} else {
					if l1&int32(1) == int32(0) {
						v41 = v19
						v42 = v22
						F_emscripten_thread_sleep(m, base.F64_add(base.F64_mul(base.F64_convert_i64_s(v42), float64(1000)), base.F64_div(base.F64_convert_i32_s(v41), float64(1e+06))))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v62 = int32(0)
							m.G0 = v12 + int32(16)
							return v62
						}
					} else {
						v29 = F___clock_gettime(m, l0, v12)
						mBase = m.M
						v30 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						v31 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
						if v30 < v31 {
							v62 = int32(0)
							m.G0 = v12 + int32(16)
							return v62
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							if v30 != v31 {
								v41 = v33 - v34
								v42 = v30 - v31
								F_emscripten_thread_sleep(m, base.F64_add(base.F64_mul(base.F64_convert_i64_s(v42), float64(1000)), base.F64_div(base.F64_convert_i32_s(v41), float64(1e+06))))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v62 = int32(0)
									m.G0 = v12 + int32(16)
									return v62
								}
							} else {
								if v33 <= v34 {
									v62 = int32(0)
									m.G0 = v12 + int32(16)
									return v62
								} else {
									v41 = v33 - v34
									v42 = v30 - v31
									F_emscripten_thread_sleep(m, base.F64_add(base.F64_mul(base.F64_convert_i64_s(v42), float64(1000)), base.F64_div(base.F64_convert_i32_s(v41), float64(1e+06))))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v62 = int32(0)
										m.G0 = v12 + int32(16)
										return v62
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
func F_call_binTM(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v114 int32
	_ = v114
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	switch v9 + int32(-5) {
	case 0:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v24 = v12 + int32(16)
	default:
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v24 = v18 + v9<<(uint(int32(2))%32) + int32(152)
	case 2:
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v24 = v15 + int32(8)
	}
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = m.G398
	if v25 == int32(0) {
		v35 = v26
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+l4<<(uint(int32(2))%32))+188))
		v34 = F_luaH_getstr(m, v25, v33)
		mBase = m.M
		v35 = v34
	}
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if v36 != 0 {
		v68 = v35
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
		*(*int64)(unsafe.Add(mBase, uint32(v70))) = v71
		v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v73
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v76 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v75)+16)) = v76
		v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v75)+24)) = v78
		v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v81 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		*(*int64)(unsafe.Add(mBase, uint32(v80)+32)) = v81
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v80)+40)) = v83
		v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if int32(48) < v85-v86 {
			v96 = v86
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v96 + int32(48)
			F_luaD_call(m, l0, v96, int32(1))
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v105 = v103 + int32(-16)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v105
				v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v109 = v107 + (l3 - v69)
				v110 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
				*(*int64)(unsafe.Add(mBase, uint32(v109))) = v110
				v114 = *(*int32)(unsafe.Add(mBase, uint32(v103+int32(-8))))
				*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v114
				return int32(1)
			}
		} else {
			F_luaD_growstack(m, l0, int32(3))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v96 = v95
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v96 + int32(48)
				F_luaD_call(m, l0, v96, int32(1))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v105 = v103 + int32(-16)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v105
					v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v109 = v107 + (l3 - v69)
					v110 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
					*(*int64)(unsafe.Add(mBase, uint32(v109))) = v110
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v103+int32(-8))))
					*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v114
					return int32(1)
				}
			}
		}
	} else {
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		switch v38 + int32(-5) {
		case 0:
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v53 = v41 + int32(16)
		default:
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v53 = v47 + v38<<(uint(int32(2))%32) + int32(152)
		case 2:
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v53 = v44 + int32(8)
		}
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
		v55 = m.G398
		if v54 == int32(0) {
			v64 = v55
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+l4<<(uint(int32(2))%32))+188))
			v63 = F_luaH_getstr(m, v54, v62)
			mBase = m.M
			v64 = v63
		}
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
		if v65 != 0 {
			v68 = v64
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
			*(*int64)(unsafe.Add(mBase, uint32(v70))) = v71
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v73
			v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v76 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v75)+16)) = v76
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v75)+24)) = v78
			v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v81 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
			*(*int64)(unsafe.Add(mBase, uint32(v80)+32)) = v81
			v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v80)+40)) = v83
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if int32(48) < v85-v86 {
				v96 = v86
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v96 + int32(48)
				F_luaD_call(m, l0, v96, int32(1))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v105 = v103 + int32(-16)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v105
					v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v109 = v107 + (l3 - v69)
					v110 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
					*(*int64)(unsafe.Add(mBase, uint32(v109))) = v110
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v103+int32(-8))))
					*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v114
					return int32(1)
				}
			} else {
				F_luaD_growstack(m, l0, int32(3))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v96 = v95
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v96 + int32(48)
					F_luaD_call(m, l0, v96, int32(1))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v105 = v103 + int32(-16)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v105
						v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v109 = v107 + (l3 - v69)
						v110 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
						*(*int64)(unsafe.Add(mBase, uint32(v109))) = v110
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v103+int32(-8))))
						*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v114
						return int32(1)
					}
				}
			}
		} else {
			return int32(0)
		}
	}
}
func F_casefold(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v4 = F_casemap(m, l0, int32(1))
	if v4 != l0 {
		v8 = v4
	} else {
		v7 = F_casemap(m, l0, int32(0))
		v8 = v7
	}
	return v8
}
func F_catAppendOnlyGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v16)
	v19 = v14 | int32(1)
	v21 = base.I64_extend_i32_s(l1)
	if v21 <= int64(-1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v66 = int32(2573)
	*(*uint16)(unsafe.Add(mBase, uint32(v62+v14+int32(1)))) = uint16(v66)
	v70 = F_sdscatlen(m, l0, v14, v62+int32(3))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v62 = int32(0)
	goto L1
L4:
	;
	v43 = F_ull2string(m, v39, v40, v41)
	mBase = m.M
	if v43 == int32(0) {
		goto L2
	} else {
		goto L8
	}
L5:
	;
	goto L7
L6:
	;
	v39 = v19
	v40 = int32(31)
	v41 = v21
	v42 = int32(0)
	goto L4
L7:
	;
	v30 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v30)
	v34 = int32(1)
	v39 = v19 + v34
	v40 = int32(30)
	v41 = int64(0) - v21
	v42 = v34
	goto L4
L8:
	;
	v62 = v43 + v42
	goto L1
L10:
	;
	return int32(0)
L11:
	;
	if l1 < int32(1) {
		v210 = v70
		goto L12
	} else {
		goto L13
	}
L12:
	;
	m.G0 = v14 + int32(32)
	return v210
L13:
	;
	v82 = int32(0)
	v83 = v70
	goto L14
L14:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2+v82<<(uint(int32(2))%32))))
	v92 = F_getDecodedObject(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v210 = v197
	goto L12
L16:
	;
	v94 = int32(36)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v94)
	v96 = int32(0)
	v98 = F_objectGetVal(m, v92)
	mBase = m.M
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+int32(-1)))))
	switch v101 & int32(7) {
	case 0:
		goto L22
	case 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		v118 = v96
		goto L17
	}
L17:
	;
	v120 = base.I64_extend_i32_u(v118)
	if v120 <= int64(-1) {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v98+int32(-17))))
	v118 = v117
	goto L17
L19:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v98+int32(-9))))
	v118 = v114
	goto L17
L20:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98+int32(-5)))))
	v118 = v111
	goto L17
L21:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+int32(-3)))))
	v118 = v108
	goto L17
L22:
	;
	v118 = int32(base.Ui32(v101) >> (uint(int32(3)) % 32))
	goto L17
L23:
	;
	v165 = int32(2573)
	*(*uint16)(unsafe.Add(mBase, uint32(v161+v14+int32(1)))) = uint16(v165)
	v169 = F_sdscatlen(m, v83, v14, v161+int32(3))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L10
	} else {
		goto L32
	}
L24:
	;
	v161 = int32(0)
	goto L23
L26:
	;
	v142 = F_ull2string(m, v138, v139, v140)
	mBase = m.M
	if v142 == int32(0) {
		goto L24
	} else {
		goto L30
	}
L27:
	;
	goto L29
L28:
	;
	v138 = v19
	v139 = int32(31)
	v140 = v120
	v141 = int32(0)
	goto L26
L29:
	;
	v129 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v129)
	v133 = int32(1)
	v138 = v19 + v133
	v139 = int32(30)
	v140 = int64(0) - v120
	v141 = v133
	goto L26
L30:
	;
	v161 = v142 + v141
	goto L23
L32:
	;
	v171 = F_objectGetVal(m, v92)
	mBase = m.M
	v172 = F_objectGetVal(m, v92)
	mBase = m.M
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172+int32(-1)))))
	switch v175 & int32(7) {
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
		v192 = v96
		goto L33
	}
L33:
	;
	v193 = F_sdscatlen(m, v169, v171, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L10
	} else {
		goto L39
	}
L34:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v172+int32(-17))))
	v192 = v191
	goto L33
L35:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v172+int32(-9))))
	v192 = v188
	goto L33
L36:
	;
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172+int32(-5)))))
	v192 = v185
	goto L33
L37:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172+int32(-3)))))
	v192 = v182
	goto L33
L38:
	;
	v192 = int32(base.Ui32(v175) >> (uint(int32(3)) % 32))
	goto L33
L39:
	;
	v197 = F_sdscatlen(m, v193, int32(_a132), int32(2))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	F_decrRefCount(m, v92)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	v202 = v82 + int32(1)
	if v202 != l1 {
		v82 = v202
		v83 = v197
		goto L14
	} else {
		goto L42
	}
L42:
	;
	goto L15
}
func F_catSubCommandFullname(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = F_sdsempty(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		v16 = F_sdscatfmt(m, v9, int32(_a1579), v7)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v16
		}
	}
}
func F_changeListener(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
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
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	m.T0[v14].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L1
L5:
	;
	m.G0 = v8 + int32(16)
	return v109
L6:
	;
	v109 = int32(0)
	goto L5
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v24 = m.T0[v23].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v29 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v30 < int32(1) {
		v109 = v29
		goto L5
	} else {
		goto L11
	}
L9:
	;
	if v24 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v109 = int32(-1)
	goto L5
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v37 = v29
	goto L12
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+v37<<(uint(int32(2))%32))))
	v47 = F_aeCreateFileEvent(m, v41, v45, int32(1), v34, l0)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L15
	}
L13:
	;
	goto L6
L14:
	;
	v98 = v37 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v98 < v99 {
		v37 = v98
		goto L12
	} else {
		goto L28
	}
L15:
	;
	if v47 != int32(-1) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if v37 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v78 = m.T0[v77].(func(*base.Module) int32)(m)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L3
	} else {
		goto L23
	}
L18:
	;
	v55 = v37
	goto L19
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v61 = v55 + int32(-1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0+v61<<(uint(int32(2))%32))))
	F_aeDeleteFileEvent(m, v59, v65, int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L21
	}
L20:
	;
	goto L17
L21:
	;
	if int32(1) < v55 {
		v55 = v61
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if base.Ui32(int32(3)) < base.Ui32(v78) {
		v89 = int32(_a545)
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v89
	F__serverPanic_1(m, int32(_a1555), int32(6980), int32(_a1648), v8)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(int32(2))%32))+uint32(_consts[907])))
	v89 = v88
	goto L25
L27:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	goto L13
}
func F_checkAlreadyExpired(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	if v5 == int32(0) {
		v11 = int32(0)
		v13 = *(*int64)(unsafe.Add(mBase, _consts[98]))
		if v13 < l0 {
			v24 = v11
		} else {
			v15 = int32(_a44)
			v16 = *(*int32)(unsafe.Add(mBase, _consts[130]))
			v18 = *(*int32)(unsafe.Add(mBase, _consts[131]))
			if v18 != 0 {
				v24 = v11
			} else {
				if v16 != 0 {
					v24 = v11
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, _consts[132]))
					v24 = base.B2i32(v20 == int32(0))
				}
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+216))
		if v9 != 0 {
			v24 = int32(0)
		} else {
			v11 = int32(0)
			v13 = *(*int64)(unsafe.Add(mBase, _consts[98]))
			if v13 < l0 {
				v24 = v11
			} else {
				v15 = int32(_a44)
				v16 = *(*int32)(unsafe.Add(mBase, _consts[130]))
				v18 = *(*int32)(unsafe.Add(mBase, _consts[131]))
				if v18 != 0 {
					v24 = v11
				} else {
					if v16 != 0 {
						v24 = v11
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, _consts[132]))
						v24 = base.B2i32(v20 == int32(0))
					}
				}
			}
		}
	}
	return v24
}
func F_checkGoodReplicasStatus(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v1 = int32(0)
	v2 = int32(_a44)
	v3 = *(*int32)(unsafe.Add(mBase, _consts[667]))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[665]))
	v12 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v17 = *(*int32)(unsafe.Add(mBase, _consts[668]))
	return base.B2i32(v3 == v1) | base.B2i32(v7 == v1) | base.B2i32(v12 != v1) | base.B2i32(v7 <= v17)
}
func F_checkMultiPartAof(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int64
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	v9 = m.G0
	v11 = v9 - int32(144)
	m.G0 = v11
	v14 = F_puts(m, int32(_a1782))
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
	v17 = F_aofLoadManifestFromFile(m, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v20 = int32(0)
	v21 = base.B2i32(v19 != v20)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v22 == v20 {
		v27 = v21
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if v19 == int32(0) {
		v83 = v22
		v84 = int32(0)
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v27 = v25 + v21
	goto L4
L6:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	if v86 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v33 = F_makePath(m, l0, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v35 = F_fileIsRDB(m, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = int32(_a1182)
	goto L12
L11:
	;
	v37 = int32(_a1783)
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v37
	v42 = F_iprintf(m, int32(_a1784), v11+int32(128))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v46 = F_checkSingleAof(m, v32, v33, base.B2i32(v27 == int32(1)), l2, v35)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L19
	}
L14:
	;
	F_sdsfree(m, v33)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L24
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v32
	v77 = F_iprintf(m, int32(_a1785), v11+int32(112))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L23
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v32
	v66 = *(*int64)(unsafe.Add(mBase, _consts[960]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v66
	v71 = F_iprintf(m, int32(_a1786), v11+int32(96))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L22
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = int32(_a1787)
	v62 = F_iprintf(m, int32(_a1788), v11+int32(80))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L21
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = int32(_a1787)
	v54 = F_iprintf(m, int32(_a1789), v11+int32(64))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	switch v46 {
	default:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L15
	case 3:
		goto L16
	}
L20:
	;
	goto L14
L21:
	;
	goto L14
L22:
	;
	goto L14
L23:
	;
	goto L14
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v83 = v81
	v84 = int32(1)
	goto L6
L25:
	;
	F_aofManifestFree(m, v17)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L51
	}
L26:
	;
	v90 = F_puts(m, int32(_a1790))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v94 = v11 + int32(136)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v95
	goto L28
L28:
	;
	v100 = v11 + int32(136)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v102 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v102 == int32(0) {
		goto L25
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v102+base.B2i32(v105 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v111
	goto L30
L32:
	;
	v116 = v102
	v119 = v84
	goto L33
L33:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v125 = F_makePath(m, l0, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L40
	}
L34:
	;
	goto L25
L35:
	;
	F_sdsfree(m, v125)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L46
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v124
	v160 = F_iprintf(m, int32(_a1785), v11+int32(48))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L45
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v124
	v149 = *(*int64)(unsafe.Add(mBase, _consts[960]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v149
	v154 = F_iprintf(m, int32(_a1786), v11+int32(32))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L44
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a1791)
	v145 = F_iprintf(m, int32(_a1788), v11+int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L43
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a1791)
	v137 = F_iprintf(m, int32(_a1789), v11)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	v128 = v119 + int32(1)
	v131 = F_checkSingleAof(m, v124, v125, base.B2i32(v128 == v27), l2, int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	switch v131 {
	default:
		goto L39
	case 1:
		goto L38
	case 2:
		goto L36
	case 3:
		goto L37
	}
L42:
	;
	goto L35
L43:
	;
	goto L35
L44:
	;
	goto L35
L45:
	;
	goto L35
L46:
	;
	v165 = v11 + int32(136)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v167 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v167 != 0 {
		v116 = v167
		v119 = v128
		goto L33
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v167+base.B2i32(v170 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v176
	goto L48
L50:
	;
	goto L34
L51:
	;
	v189 = F_puts(m, int32(_a1792))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	m.G0 = v11 + int32(144)
	return
}
func F_checkOldStyleAof(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(64)
	m.G0 = v7
	v10 = F_puts(m, int32(_a1793))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v13 = F_checkSingleAof(m, l0, l0, int32(1), l1, l2)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			switch v13 {
			default:
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a1794)
				v19 = F_iprintf(m, int32(_a1789), v7)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					m.G0 = v7 + int32(64)
					return
				}
			case 1:
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(_a1794)
				v27 = F_iprintf(m, int32(_a1788), v7+int32(16))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					m.G0 = v7 + int32(64)
					return
				}
			case 2:
				*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l0
				v42 = F_iprintf(m, int32(_a1785), v7+int32(48))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					m.G0 = v7 + int32(64)
					return
				}
			case 3:
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l0
				v31 = *(*int64)(unsafe.Add(mBase, _consts[960]))
				*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v31
				v36 = F_iprintf(m, int32(_a1786), v7+int32(32))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					m.G0 = v7 + int32(64)
					return
				}
			}
		}
	}
}
func F_checkPrefixCollisionsOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int64
	_ = v46
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
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
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v408 int32
	_ = v408
	v15 = m.G0
	v17 = v15 - int32(336)
	m.G0 = v17
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(336)
	return v408
L2:
	;
	v26 = int32(0)
	goto L4
L3:
	;
	v408 = int32(1)
	goto L1
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
	if v36 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v408 = int32(1)
	goto L1
L6:
	;
	if v221 != l2 {
		v26 = v221
		goto L4
	} else {
		goto L87
	}
L7:
	;
	v408 = int32(0)
	goto L1
L8:
	;
	v221 = v26 + int32(1)
	if base.Ui32(l2) <= base.Ui32(v221) {
		goto L6
	} else {
		goto L49
	}
L9:
	;
	v40 = v17 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = int32(128)
	v46 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+12)) = v46
	*(*int64)(unsafe.Add(mBase, uint32(v40)+296)) = v46
	*(*int64)(unsafe.Add(mBase, uint32(v40)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v17 + int32(56)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+156)) = v17 + int32(200)
	goto L10
L10:
	;
	v61 = int32(0)
	v63 = F_raxSeek(m, v17+int32(32), int32(_a263), v61, v61)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v69 = l1 + v26<<(uint(int32(2))%32)
	goto L14
L13:
	;
	F_raxStop(m, v17+int32(32))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L11
	} else {
		goto L48
	}
L14:
	;
	v86 = F_raxNext(m, v17+int32(32))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	v185 = F_sdsnewlen(m, v183, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L11
	} else {
		goto L44
	}
L16:
	;
	if v86 == int32(0) {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v93 = F_objectGetVal(m, v92)
	mBase = m.M
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v96 = F_objectGetVal(m, v95)
	mBase = m.M
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+int32(-1)))))
	switch v99 & int32(7) {
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
		v116 = int32(0)
		goto L18
	}
L18:
	;
	if base.Ui32(v90) < base.Ui32(v116) {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v96+int32(-17))))
	v116 = v115
	goto L18
L20:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v96+int32(-9))))
	v116 = v112
	goto L18
L21:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96+int32(-5)))))
	v116 = v109
	goto L18
L22:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+int32(-3)))))
	v116 = v106
	goto L18
L23:
	;
	v116 = int32(base.Ui32(v99) >> (uint(int32(3)) % 32))
	goto L18
L24:
	;
	v118 = v90
	goto L26
L25:
	;
	v118 = v116
	goto L26
L26:
	;
	if base.Ui32(v118) < base.Ui32(int32(4)) {
		v142 = v91
		v143 = v93
		v144 = v118
		goto L30
	} else {
		goto L31
	}
L27:
	;
	if v182 != 0 {
		goto L14
	} else {
		goto L43
	}
L28:
	;
	v182 = int32(0)
	goto L27
L29:
	;
	v154 = v149
	v155 = v150
	v156 = v151
	goto L39
L30:
	;
	if v144 == int32(0) {
		goto L28
	} else {
		goto L37
	}
L31:
	;
	if (v93|v91)&int32(3) != 0 {
		v149 = v91
		v150 = v93
		v151 = v118
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v126 = v91
	v127 = v93
	v128 = v118
	goto L33
L33:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	if v131 != v132 {
		v149 = v126
		v150 = v127
		v151 = v128
		goto L29
	} else {
		goto L35
	}
L34:
	;
	v142 = v137
	v143 = v135
	v144 = v139
	goto L30
L35:
	;
	v134 = int32(4)
	v135 = v127 + v134
	v137 = v126 + v134
	v139 = v128 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v139) {
		v126 = v137
		v127 = v135
		v128 = v139
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v149 = v142
	v150 = v143
	v151 = v144
	goto L29
L38:
	;
	v182 = v159 - v160
	goto L27
L39:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if v159 != v160 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v162 = int32(1)
	v167 = v156 + int32(-1)
	if v167 == int32(0) {
		goto L28
	} else {
		goto L42
	}
L42:
	;
	v154 = v154 + v162
	v155 = v155 + v162
	v156 = v167
	goto L39
L43:
	;
	goto L15
L44:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v188 = F_objectGetVal(m, v187)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v188
	F_addReplyErrorFormat(m, l0, int32(_a1749), v17+int32(16))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L11
	} else {
		goto L45
	}
L45:
	;
	F_sdsfree(m, v185)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	F_raxStop(m, v17+int32(32))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	goto L7
L48:
	;
	goto L8
L49:
	;
	v225 = l1 + v26<<(uint(int32(2))%32)
	v233 = v221
	goto L50
L50:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v241 = F_objectGetVal(m, v240)
	mBase = m.M
	v242 = int32(0)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v245 = F_objectGetVal(m, v244)
	mBase = m.M
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+int32(-1)))))
	switch v248 & int32(7) {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	case 4:
		goto L53
	default:
		v265 = v242
		goto L52
	}
L51:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v365 = F_objectGetVal(m, v364)
	mBase = m.M
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v367 = F_objectGetVal(m, v366)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v365
	F_addReplyErrorFormat(m, l0, int32(_a1750), v17)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L11
	} else {
		goto L86
	}
L52:
	;
	v268 = l1 + v233<<(uint(int32(2))%32)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v270 = F_objectGetVal(m, v269)
	mBase = m.M
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v272 = F_objectGetVal(m, v271)
	mBase = m.M
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272+int32(-1)))))
	switch v275 & int32(7) {
	case 0:
		goto L63
	case 1:
		goto L62
	case 2:
		goto L61
	case 3:
		goto L60
	case 4:
		goto L59
	default:
		v292 = v242
		goto L58
	}
L53:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v245+int32(-17))))
	v265 = v264
	goto L52
L54:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v245+int32(-9))))
	v265 = v261
	goto L52
L55:
	;
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v245+int32(-5)))))
	v265 = v258
	goto L52
L56:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+int32(-3)))))
	v265 = v255
	goto L52
L57:
	;
	v265 = int32(base.Ui32(v248) >> (uint(int32(3)) % 32))
	goto L52
L58:
	;
	if base.Ui32(v265) < base.Ui32(v292) {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v272+int32(-17))))
	v292 = v291
	goto L58
L60:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v272+int32(-9))))
	v292 = v288
	goto L58
L61:
	;
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v272+int32(-5)))))
	v292 = v285
	goto L58
L62:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272+int32(-3)))))
	v292 = v282
	goto L58
L63:
	;
	v292 = int32(base.Ui32(v275) >> (uint(int32(3)) % 32))
	goto L58
L64:
	;
	goto L51
L65:
	;
	v294 = v265
	goto L67
L66:
	;
	v294 = v292
	goto L67
L67:
	;
	if base.Ui32(v294) < base.Ui32(int32(4)) {
		v318 = v241
		v319 = v270
		v320 = v294
		goto L71
	} else {
		goto L72
	}
L68:
	;
	if v358 == int32(0) {
		goto L64
	} else {
		goto L84
	}
L69:
	;
	v358 = int32(0)
	goto L68
L70:
	;
	v330 = v325
	v331 = v326
	v332 = v327
	goto L80
L71:
	;
	if v320 == int32(0) {
		goto L69
	} else {
		goto L78
	}
L72:
	;
	if (v270|v241)&int32(3) != 0 {
		v325 = v241
		v326 = v270
		v327 = v294
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v302 = v241
	v303 = v270
	v304 = v294
	goto L74
L74:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	if v307 != v308 {
		v325 = v302
		v326 = v303
		v327 = v304
		goto L70
	} else {
		goto L76
	}
L75:
	;
	v318 = v313
	v319 = v311
	v320 = v315
	goto L71
L76:
	;
	v310 = int32(4)
	v311 = v303 + v310
	v313 = v302 + v310
	v315 = v304 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v315) {
		v302 = v313
		v303 = v311
		v304 = v315
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v325 = v318
	v326 = v319
	v327 = v320
	goto L70
L79:
	;
	v358 = v335 - v336
	goto L68
L80:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v335 != v336 {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	v338 = int32(1)
	v343 = v332 + int32(-1)
	if v343 == int32(0) {
		goto L69
	} else {
		goto L83
	}
L83:
	;
	v330 = v330 + v338
	v331 = v331 + v338
	v332 = v343
	goto L80
L84:
	;
	v362 = v233 + int32(1)
	if v362 == l2 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	v233 = v362
	goto L50
L86:
	;
	goto L7
L87:
	;
	goto L5
}
func F_checkint(m *base.Module, l0 int64) int32 {
	var v10 int32
	_ = v10
	var v21 int64
	_ = v21
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v10 = base.I32_wrap_i64(int64(base.Ui64(l0)>>(uint(int64(52))%64))) & int32(2047)
	if base.Ui32(v10) < base.Ui32(int32(1023)) {
		v33 = int32(0)
	} else {
		if base.Ui32(int32(1075)) < base.Ui32(v10) {
			v33 = int32(2)
		} else {
			v21 = int64(1) << (uint(base.I64_extend_i32_u(int32(1075)-v10)) % 64)
			if (v21+int64(-1))&l0 != int64(0) {
				v33 = int32(0)
			} else {
				if v21&l0 == int64(0) {
					v32 = int32(2)
				} else {
					v32 = int32(1)
				}
				v33 = v32
			}
		}
	}
	return v33
}
func F_chunk(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
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
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
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
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v239 int32
	_ = v239
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v421 int32
	_ = v421
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v825 int32
	_ = v825
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v878 int32
	_ = v878
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v997 int32
	_ = v997
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1090 int32
	_ = v1090
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1125 int32
	_ = v1125
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1143 int32
	_ = v1143
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1235 int32
	_ = v1235
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1287 int32
	_ = v1287
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1322 int32
	_ = v1322
	var v1328 int32
	_ = v1328
	var v1332 int32
	_ = v1332
	var v1340 int32
	_ = v1340
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1393 int32
	_ = v1393
	var v1412 int32
	_ = v1412
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1457 int32
	_ = v1457
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1505 int32
	_ = v1505
	var v1509 int32
	_ = v1509
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1537 int32
	_ = v1537
	var v1542 int32
	_ = v1542
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1558 int32
	_ = v1558
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1589 int32
	_ = v1589
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1624 int32
	_ = v1624
	var v1630 int32
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1642 int32
	_ = v1642
	var v1677 int32
	_ = v1677
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1710 int32
	_ = v1710
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1720 int32
	_ = v1720
	var v1725 int32
	_ = v1725
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1750 int32
	_ = v1750
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1787 int32
	_ = v1787
	var v1789 int32
	_ = v1789
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1813 int32
	_ = v1813
	var v1827 int32
	_ = v1827
	var v1831 int32
	_ = v1831
	var v1836 int32
	_ = v1836
	var v1847 int32
	_ = v1847
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1863 int32
	_ = v1863
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1893 int32
	_ = v1893
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1927 int32
	_ = v1927
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1938 int32
	_ = v1938
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1984 int32
	_ = v1984
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1993 int32
	_ = v1993
	var v1997 int32
	_ = v1997
	var v2010 int32
	_ = v2010
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2028 int32
	_ = v2028
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2038 int32
	_ = v2038
	var v2040 int32
	_ = v2040
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2067 int32
	_ = v2067
	var v2071 int32
	_ = v2071
	var v2073 int32
	_ = v2073
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2080 int32
	_ = v2080
	var v2083 int32
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2130 int32
	_ = v2130
	var v2132 int32
	_ = v2132
	var v2138 int32
	_ = v2138
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2154 int32
	_ = v2154
	var v2158 int32
	_ = v2158
	var v2165 int32
	_ = v2165
	var v2176 int32
	_ = v2176
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2194 int32
	_ = v2194
	var v2196 int32
	_ = v2196
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2212 int32
	_ = v2212
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2226 int32
	_ = v2226
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2246 int32
	_ = v2246
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2257 int32
	_ = v2257
	var v2264 int32
	_ = v2264
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2295 int32
	_ = v2295
	var v2300 int32
	_ = v2300
	var v2314 int32
	_ = v2314
	var v2317 int32
	_ = v2317
	var v2320 int32
	_ = v2320
	var v2326 int32
	_ = v2326
	var v2328 int32
	_ = v2328
	var v2332 int32
	_ = v2332
	var v2352 int32
	_ = v2352
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2379 int32
	_ = v2379
	var v2386 int32
	_ = v2386
	var v2393 int32
	_ = v2393
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2410 int32
	_ = v2410
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2419 int32
	_ = v2419
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2432 int32
	_ = v2432
	var v2439 int32
	_ = v2439
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2458 int32
	_ = v2458
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2463 int32
	_ = v2463
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2475 int32
	_ = v2475
	var v2477 int32
	_ = v2477
	var v2480 int32
	_ = v2480
	var v2491 int32
	_ = v2491
	var v2500 int32
	_ = v2500
	var v2505 int32
	_ = v2505
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2510 int32
	_ = v2510
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2518 int32
	_ = v2518
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2534 int32
	_ = v2534
	var v2551 int32
	_ = v2551
	var v2552 int32
	_ = v2552
	var v2556 int32
	_ = v2556
	var v2560 int32
	_ = v2560
	var v2561 int32
	_ = v2561
	var v2563 int32
	_ = v2563
	var v2564 int32
	_ = v2564
	var v2569 int32
	_ = v2569
	var v2589 int32
	_ = v2589
	var v2592 int32
	_ = v2592
	var v2595 int32
	_ = v2595
	var v2617 int32
	_ = v2617
	var v2618 int32
	_ = v2618
	var v2619 int32
	_ = v2619
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2631 int32
	_ = v2631
	var v2634 int32
	_ = v2634
	var v2639 int32
	_ = v2639
	var v2653 int32
	_ = v2653
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2660 int32
	_ = v2660
	var v2666 int32
	_ = v2666
	var v2670 int32
	_ = v2670
	var v2691 int32
	_ = v2691
	var v2707 int32
	_ = v2707
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2717 int32
	_ = v2717
	var v2719 int32
	_ = v2719
	var v2726 int32
	_ = v2726
	var v2732 int32
	_ = v2732
	var v2736 int32
	_ = v2736
	var v2744 int32
	_ = v2744
	var v2779 int32
	_ = v2779
	var v2783 int32
	_ = v2783
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2789 int32
	_ = v2789
	var v2791 int32
	_ = v2791
	var v2813 int32
	_ = v2813
	var v2829 int32
	_ = v2829
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2858 int32
	_ = v2858
	v19 = m.G0
	v21 = v19 - int32(800)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+52)))
	v26 = v24 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+52)) = uint16(v26)
	if base.Ui32(v26&int32(65535)) < base.Ui32(int32(201)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L6
L2:
	;
	v32 = m.G3
	F_luaX_lexerror(m, l0, v32+int32(_a2036), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	goto L1
L5:
	;
	v2855 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2856 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2855)+52)))
	v2858 = v2856 + int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2855)+52)) = uint16(v2858)
	m.G0 = v21 + int32(800)
	return
L6:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v60 = v58 + int32(-260)
	if base.Ui32(int32(27)) < base.Ui32(v60) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v58 + int32(-258) {
	case 0:
		goto L17
	case 1:
		goto L24
	default:
		goto L18
	case 6:
		goto L23
	case 7:
		goto L21
	case 8:
		goto L26
	case 10:
		goto L15
	case 14:
		goto L22
	case 15:
		goto L16
	case 19:
		goto L25
	}
L9:
	;
	if int32(1)<<(uint(v60)%32)&int32(134283271) != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v2829 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2829 != int32(59) {
		goto L434
	} else {
		goto L435
	}
L12:
	;
	v2813 = int32(1)
	goto L11
L13:
	;
	F_check_match(m, l0, int32(262), int32(264), v67)
	mBase = m.M
	v2617 = m.ExcPending
	if v2617 != 0 {
		goto L3
	} else {
		goto L417
	}
L14:
	;
	v2400 = m.G3
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v2401)+36))
	v2406 = F_luaX_newstring(m, l0, v2400+int32(_a2037), int32(15))
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
		goto L3
	} else {
		goto L379
	}
L15:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L3
	} else {
		goto L314
	}
L16:
	;
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaX_next(m, l0)
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L3
	} else {
		goto L288
	}
L17:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L3
	} else {
		goto L273
	}
L18:
	;
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_primaryexp(m, l0, v21+int32(208))
	mBase = m.M
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L3
	} else {
		goto L269
	}
L19:
	;
	v1453 = v1450
	v1457 = v90
	goto L220
L20:
	;
	v1450 = int32(1)
	goto L19
L21:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L3
	} else {
		goto L207
	}
L22:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v702)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v702)+28)) = v704
	goto L130
L23:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v533 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+786)) = uint8(v533)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+780)) = int32(-1)
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532)+50)))
	v538 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+785)) = uint8(v538)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+784)) = uint8(v537)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v532)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+776)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v532)+20)) = v21 + int32(776)
	F_luaX_next(m, l0)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L3
	} else {
		goto L85
	}
L24:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L3
	} else {
		goto L82
	}
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaX_next(m, l0)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L3
	} else {
		goto L33
	}
L26:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+776)) = int32(-1)
	F_luaX_next(m, l0)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v78 = F_subexpr(m, l0, v21+int32(200), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
	if v80 != int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_goiftrue(m, v85, v21+int32(200))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L3
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = int32(3)
	goto L29
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v21)+220))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v91 == int32(274) {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	v1450 = int32(0)
	goto L19
L33:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+28)) = v99
	goto L34
L34:
	;
	v104 = F_subexpr(m, l0, v21+int32(200), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
	if v106 != int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_goiftrue(m, v111, v21+int32(200))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L3
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = int32(3)
	goto L36
L38:
	;
	v116 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+786)) = uint8(v116)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+780)) = int32(-1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v21)+220))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+50)))
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+785)) = uint8(v122)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+784)) = uint8(v121)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+776)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v95)+20)) = v21 + int32(776)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v130 == int32(259) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L44
	}
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v135 = F_luaX_token2str(m, l0, int32(259))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v135
	v138 = m.G3
	v143 = F_luaO_pushfstring(m, v133, v138+int32(_a2038), v21+int32(16))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	F_luaX_syntaxerror(m, l0, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+210)) = uint8(v151)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = int32(-1)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+50)))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+209)) = uint8(v151)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+208)) = uint8(v155)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v150)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v150)+20)) = v21 + int32(200)
	F_chunk(m, l0)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v150)+20))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	*(*int32)(unsafe.Add(mBase, uint32(v150)+20)) = v167
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+8)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+48))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+50)))
	if base.Ui32(v172) <= base.Ui32(v169) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+9)))
	if v327 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L47:
	;
	v175 = v171 + int32(172)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v171)+24))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+24))
	v179 = int32(0)
	v182 = (v172 - v169) & int32(3)
	if v182 == v179 {
		v218 = v172
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v169-v172) {
		v292 = v218
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v187 = v172
	v201 = v179
	goto L50
L50:
	;
	v204 = v187 + int32(-1)
	v205 = int32(1)
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175+v204<<(uint(v205)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v178+v208*int32(12))+8)) = v176
	v214 = v201 + v205
	if v214 != v182 {
		v187 = v204
		v201 = v214
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v218 = v204
	goto L48
L52:
	;
	goto L51
L53:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v171)+50)) = uint8(v292)
	goto L46
L54:
	;
	v239 = v218
	goto L55
L55:
	;
	v255 = int32(1)
	v257 = v239<<(uint(v255)%32) + v175
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257+int32(-2)))))
	v261 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v178+v260*v261)+8)) = v176
	v265 = int32(-4)
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257+v265))))
	*(*int32)(unsafe.Add(mBase, uint32(v178+v267*v261)+8)) = v176
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257+int32(-6)))))
	*(*int32)(unsafe.Add(mBase, uint32(v178+v274*v261)+8)) = v176
	v280 = v239 + v265
	v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175+v280<<(uint(v255)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v178+v284*v261)+8)) = v176
	if base.Ui32(v169) < base.Ui32(v280) {
		v239 = v280
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v292 = v280
	goto L53
L57:
	;
	goto L56
L58:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+50)))
	*(*int32)(unsafe.Add(mBase, uint32(v150)+36)) = v335
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	F_luaK_patchtohere(m, v150, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L3
	} else {
		goto L61
	}
L59:
	;
	v331 = int32(0)
	v333 = F_luaK_codeABC(m, v150, int32(35), v169, v331, v331)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v340 = F_luaK_jump(m, v95)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	F_luaK_patchlist(m, v95, v340, v99)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	F_check_match(m, l0, int32(262), int32(277), v67)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+20)) = v349
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+8)))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+48))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+50)))
	if base.Ui32(v354) <= base.Ui32(v351) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+9)))
	if v509 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L66:
	;
	v357 = v353 + int32(172)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v353)+24))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+24))
	v361 = int32(0)
	v364 = (v354 - v351) & int32(3)
	if v364 == v361 {
		v400 = v354
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v351-v354) {
		v474 = v400
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v369 = v354
	v383 = v361
	goto L69
L69:
	;
	v386 = v369 + int32(-1)
	v387 = int32(1)
	v390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357+v386<<(uint(v387)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v360+v390*int32(12))+8)) = v358
	v396 = v383 + v387
	if v396 != v364 {
		v369 = v386
		v383 = v396
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v400 = v386
	goto L67
L71:
	;
	goto L70
L72:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v353)+50)) = uint8(v474)
	goto L65
L73:
	;
	v421 = v400
	goto L74
L74:
	;
	v437 = int32(1)
	v439 = v421<<(uint(v437)%32) + v357
	v442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v439+int32(-2)))))
	v443 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v360+v442*v443)+8)) = v358
	v447 = int32(-4)
	v449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v439+v447))))
	*(*int32)(unsafe.Add(mBase, uint32(v360+v449*v443)+8)) = v358
	v456 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v439+int32(-6)))))
	*(*int32)(unsafe.Add(mBase, uint32(v360+v456*v443)+8)) = v358
	v462 = v421 + v447
	v466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357+v462<<(uint(v437)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v360+v466*v443)+8)) = v358
	if base.Ui32(v351) < base.Ui32(v462) {
		v421 = v462
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v474 = v462
	goto L72
L76:
	;
	goto L75
L77:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+50)))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+36)) = v517
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	F_luaK_patchtohere(m, v95, v519)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L3
	} else {
		goto L80
	}
L78:
	;
	v513 = int32(0)
	v515 = F_luaK_codeABC(m, v95, int32(35), v351, v513, v513)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	F_luaK_patchtohere(m, v95, v120)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L3
	} else {
		goto L81
	}
L81:
	;
	goto L12
L82:
	;
	F_block(m, l0)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	F_check_match(m, l0, int32(262), int32(259), v67)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L3
	} else {
		goto L84
	}
L84:
	;
	goto L12
L85:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v548 == int32(285) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_luaX_next(m, l0)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L3
	} else {
		goto L91
	}
L87:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v553 = F_luaX_token2str(m, l0, int32(285))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v553
	v556 = m.G3
	v561 = F_luaO_pushfstring(m, v551, v556+int32(_a2038), v21+int32(96))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L3
	} else {
		goto L89
	}
L89:
	;
	F_luaX_syntaxerror(m, l0, v561)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L3
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v569 + int32(-44) {
	case 0:
		goto L14
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
		goto L92
	case 17:
		goto L94
	default:
		goto L93
	}
L92:
	;
	v697 = m.G3
	F_luaX_syntaxerror(m, l0, v697+int32(_a2039))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L3
	} else {
		goto L129
	}
L93:
	;
	if v569 == int32(267) {
		goto L14
	} else {
		goto L128
	}
L94:
	;
	v572 = m.G3
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)+36))
	v578 = F_luaX_newstring(m, l0, v572+int32(_a2040), int32(11))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L3
	} else {
		goto L95
	}
L95:
	;
	F_new_localvar(m, l0, v578, int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L3
	} else {
		goto L96
	}
L96:
	;
	v586 = F_luaX_newstring(m, l0, v572+int32(_a2041), int32(11))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L3
	} else {
		goto L97
	}
L97:
	;
	F_new_localvar(m, l0, v586, int32(1))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v594 = F_luaX_newstring(m, l0, v572+int32(_a2042), int32(10))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	F_new_localvar(m, l0, v594, int32(2))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	F_new_localvar(m, l0, v566, int32(3))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L3
	} else {
		goto L101
	}
L101:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v602 == int32(61) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L3
	} else {
		goto L107
	}
L103:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v607 = F_luaX_token2str(m, l0, int32(61))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v607
	v610 = m.G3
	v615 = F_luaO_pushfstring(m, v605, v610+int32(_a2038), v21+int32(48))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L3
	} else {
		goto L105
	}
L105:
	;
	F_luaX_syntaxerror(m, l0, v615)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L3
	} else {
		goto L106
	}
L106:
	;
	goto L102
L107:
	;
	v625 = F_subexpr(m, l0, v21+int32(200), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_exp2nextreg(m, v627, v21+int32(200))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L3
	} else {
		goto L109
	}
L109:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v632 == int32(44) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L3
	} else {
		goto L115
	}
L111:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v637 = F_luaX_token2str(m, l0, int32(44))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v637
	v640 = m.G3
	v645 = F_luaO_pushfstring(m, v635, v640+int32(_a2038), v21+int32(32))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	F_luaX_syntaxerror(m, l0, v645)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L3
	} else {
		goto L114
	}
L114:
	;
	goto L110
L115:
	;
	v655 = F_subexpr(m, l0, v21+int32(200), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_exp2nextreg(m, v657, v21+int32(200))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L3
	} else {
		goto L117
	}
L117:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v662 != int32(44) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v573)+36))
	v684 = F_luaK_numberK(m, v573, float64(1))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L3
	} else {
		goto L124
	}
L119:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	v670 = F_subexpr(m, l0, v21+int32(200), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L3
	} else {
		goto L121
	}
L121:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_exp2nextreg(m, v672, v21+int32(200))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L3
	} else {
		goto L122
	}
L122:
	;
	v677 = int32(1)
	F_forbody(m, l0, v574, v67, v677, v677)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L3
	} else {
		goto L123
	}
L123:
	;
	goto L13
L124:
	;
	v686 = F_luaK_codeABx(m, v573, int32(1), v682, v684)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	F_luaK_reserveregs(m, v573, int32(1))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	v691 = int32(1)
	F_forbody(m, l0, v574, v67, v691, v691)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	goto L13
L128:
	;
	goto L92
L129:
	;
	goto L13
L130:
	;
	v706 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+780)) = v706
	v708 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+786)) = uint8(v708)
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702)+50)))
	v711 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+785)) = uint8(v711)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+784)) = uint8(v710)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v702)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v21)+776)) = v714
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+185)) = uint16(v711)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+184)) = uint8(v710)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v21 + int32(776)
	*(*int32)(unsafe.Add(mBase, uint32(v702)+20)) = v21 + int32(176)
	F_luaX_next(m, l0)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	F_chunk(m, l0)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L3
	} else {
		goto L132
	}
L132:
	;
	F_check_match(m, l0, int32(276), int32(272), v67)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L3
	} else {
		goto L133
	}
L133:
	;
	v738 = F_subexpr(m, l0, v21+int32(200), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L3
	} else {
		goto L134
	}
L134:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
	if v740 != int32(1) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_goiftrue(m, v745, v21+int32(200))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L3
	} else {
		goto L137
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = int32(3)
	goto L135
L137:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v21)+220))
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+185)))
	if v751 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v702)+20))
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1214)))
	*(*int32)(unsafe.Add(mBase, uint32(v702)+20)) = v1215
	v1217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1214)+8)))
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v702)+12))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+48))
	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+50)))
	if base.Ui32(v1220) <= base.Ui32(v1217) {
		goto L191
	} else {
		goto L192
	}
L139:
	;
	v929 = int32(0)
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v931)+20))
	if v932 == v929 {
		v963 = v929
		goto L159
	} else {
		goto L160
	}
L140:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v702)+20))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v752)))
	*(*int32)(unsafe.Add(mBase, uint32(v702)+20)) = v753
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v752)+8)))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v702)+12))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)+48))
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757)+50)))
	if base.Ui32(v758) <= base.Ui32(v755) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v752)+9)))
	if v913 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L142:
	;
	v761 = v757 + int32(172)
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v757)+24))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v757)))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v763)+24))
	v765 = int32(0)
	v768 = (v758 - v755) & int32(3)
	if v768 == v765 {
		v804 = v758
		goto L143
	} else {
		goto L144
	}
L143:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v755-v758) {
		v878 = v804
		goto L148
	} else {
		goto L149
	}
L144:
	;
	v773 = v758
	v787 = v765
	goto L145
L145:
	;
	v790 = v773 + int32(-1)
	v791 = int32(1)
	v794 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761+v790<<(uint(v791)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v764+v794*int32(12))+8)) = v762
	v800 = v787 + v791
	if v800 != v768 {
		v773 = v790
		v787 = v800
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v804 = v790
	goto L143
L147:
	;
	goto L146
L148:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v757)+50)) = uint8(v878)
	goto L141
L149:
	;
	v825 = v804
	goto L150
L150:
	;
	v841 = int32(1)
	v843 = v825<<(uint(v841)%32) + v761
	v846 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v843+int32(-2)))))
	v847 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v764+v846*v847)+8)) = v762
	v851 = int32(-4)
	v853 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v843+v851))))
	*(*int32)(unsafe.Add(mBase, uint32(v764+v853*v847)+8)) = v762
	v860 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v843+int32(-6)))))
	*(*int32)(unsafe.Add(mBase, uint32(v764+v860*v847)+8)) = v762
	v866 = v825 + v851
	v870 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761+v866<<(uint(v841)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v764+v870*v847)+8)) = v762
	if base.Ui32(v755) < base.Ui32(v866) {
		v825 = v866
		goto L150
	} else {
		goto L152
	}
L151:
	;
	v878 = v866
	goto L148
L152:
	;
	goto L151
L153:
	;
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702)+50)))
	*(*int32)(unsafe.Add(mBase, uint32(v702)+36)) = v921
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v752)+4))
	F_luaK_patchtohere(m, v702, v923)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L3
	} else {
		goto L156
	}
L154:
	;
	v917 = int32(0)
	v919 = F_luaK_codeABC(m, v702, int32(35), v755, v917, v917)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L3
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_patchlist(m, v926, v750, v704)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L3
	} else {
		goto L157
	}
L157:
	;
	goto L138
L158:
	;
	if v986 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L159:
	;
	v977 = m.G3
	F_luaX_syntaxerror(m, l0, v977+int32(_a2043))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L3
	} else {
		goto L166
	}
L160:
	;
	v937 = v932
	v939 = v929
	goto L161
L161:
	;
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937)+10)))
	if v953 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v963 = v957
	goto L159
L163:
	;
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937)+9)))
	v957 = v939 | v956
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v937)))
	if v958 != 0 {
		v937 = v958
		v939 = v957
		goto L161
	} else {
		goto L165
	}
L164:
	;
	v986 = v939
	v997 = v937
	goto L158
L165:
	;
	goto L162
L166:
	;
	v986 = v963
	v997 = v929
	goto L158
L167:
	;
	v1010 = F_luaK_jump(m, v931)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L3
	} else {
		goto L170
	}
L168:
	;
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v997)+8)))
	v1004 = int32(0)
	v1006 = F_luaK_codeABC(m, v931, int32(35), v1003, v1004, v1004)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L3
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	F_luaK_concat(m, v931, v997+int32(4), v1010)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L3
	} else {
		goto L171
	}
L171:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_patchtohere(m, v1014, v750)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L3
	} else {
		goto L172
	}
L172:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v702)+20))
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1017)))
	*(*int32)(unsafe.Add(mBase, uint32(v702)+20)) = v1018
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1017)+8)))
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v702)+12))
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+48))
	v1023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1022)+50)))
	if base.Ui32(v1023) <= base.Ui32(v1020) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v1178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1017)+9)))
	if v1178 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L174:
	;
	v1026 = v1022 + int32(172)
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+24))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1022)))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+24))
	v1030 = int32(0)
	v1033 = (v1023 - v1020) & int32(3)
	if v1033 == v1030 {
		v1069 = v1023
		goto L175
	} else {
		goto L176
	}
L175:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v1020-v1023) {
		v1143 = v1069
		goto L180
	} else {
		goto L181
	}
L176:
	;
	v1038 = v1023
	v1052 = v1030
	goto L177
L177:
	;
	v1055 = v1038 + int32(-1)
	v1056 = int32(1)
	v1059 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1026+v1055<<(uint(v1056)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1029+v1059*int32(12))+8)) = v1027
	v1065 = v1052 + v1056
	if v1065 != v1033 {
		v1038 = v1055
		v1052 = v1065
		goto L177
	} else {
		goto L179
	}
L178:
	;
	v1069 = v1055
	goto L175
L179:
	;
	goto L178
L180:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1022)+50)) = uint8(v1143)
	goto L173
L181:
	;
	v1090 = v1069
	goto L182
L182:
	;
	v1106 = int32(1)
	v1108 = v1090<<(uint(v1106)%32) + v1026
	v1111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1108+int32(-2)))))
	v1112 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v1029+v1111*v1112)+8)) = v1027
	v1116 = int32(-4)
	v1118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1108+v1116))))
	*(*int32)(unsafe.Add(mBase, uint32(v1029+v1118*v1112)+8)) = v1027
	v1125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1108+int32(-6)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1029+v1125*v1112)+8)) = v1027
	v1131 = v1090 + v1116
	v1135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1026+v1131<<(uint(v1106)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1029+v1135*v1112)+8)) = v1027
	if base.Ui32(v1020) < base.Ui32(v1131) {
		v1090 = v1131
		goto L182
	} else {
		goto L184
	}
L183:
	;
	v1143 = v1131
	goto L180
L184:
	;
	goto L183
L185:
	;
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702)+50)))
	*(*int32)(unsafe.Add(mBase, uint32(v702)+36)) = v1186
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1017)+4))
	F_luaK_patchtohere(m, v702, v1188)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L3
	} else {
		goto L188
	}
L186:
	;
	v1182 = int32(0)
	v1184 = F_luaK_codeABC(m, v702, int32(35), v1020, v1182, v1182)
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L3
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1192 = F_luaK_jump(m, v702)
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L3
	} else {
		goto L189
	}
L189:
	;
	F_luaK_patchlist(m, v1191, v1192, v704)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L3
	} else {
		goto L190
	}
L190:
	;
	goto L138
L191:
	;
	v1375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1214)+9)))
	if v1375 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L192:
	;
	v1223 = v1219 + int32(172)
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1219)+24))
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1219)))
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+24))
	v1227 = int32(0)
	v1230 = (v1220 - v1217) & int32(3)
	if v1230 == v1227 {
		v1266 = v1220
		goto L193
	} else {
		goto L194
	}
L193:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v1217-v1220) {
		v1340 = v1266
		goto L198
	} else {
		goto L199
	}
L194:
	;
	v1235 = v1220
	v1249 = v1227
	goto L195
L195:
	;
	v1252 = v1235 + int32(-1)
	v1253 = int32(1)
	v1256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1223+v1252<<(uint(v1253)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1226+v1256*int32(12))+8)) = v1224
	v1262 = v1249 + v1253
	if v1262 != v1230 {
		v1235 = v1252
		v1249 = v1262
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v1266 = v1252
	goto L193
L197:
	;
	goto L196
L198:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1219)+50)) = uint8(v1340)
	goto L191
L199:
	;
	v1287 = v1266
	goto L200
L200:
	;
	v1303 = int32(1)
	v1305 = v1287<<(uint(v1303)%32) + v1223
	v1308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1305+int32(-2)))))
	v1309 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v1226+v1308*v1309)+8)) = v1224
	v1313 = int32(-4)
	v1315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1305+v1313))))
	*(*int32)(unsafe.Add(mBase, uint32(v1226+v1315*v1309)+8)) = v1224
	v1322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1305+int32(-6)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1226+v1322*v1309)+8)) = v1224
	v1328 = v1287 + v1313
	v1332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1223+v1328<<(uint(v1303)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1226+v1332*v1309)+8)) = v1224
	if base.Ui32(v1217) < base.Ui32(v1328) {
		v1287 = v1328
		goto L200
	} else {
		goto L202
	}
L201:
	;
	v1340 = v1328
	goto L198
L202:
	;
	goto L201
L203:
	;
	v1383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702)+50)))
	*(*int32)(unsafe.Add(mBase, uint32(v702)+36)) = v1383
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+4))
	F_luaK_patchtohere(m, v702, v1385)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L3
	} else {
		goto L206
	}
L204:
	;
	v1379 = int32(0)
	v1381 = F_luaK_codeABC(m, v702, int32(35), v1217, v1379, v1379)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L3
	} else {
		goto L205
	}
L205:
	;
	goto L203
L206:
	;
	goto L12
L207:
	;
	F_singlevar(m, l0, v21+int32(200))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L3
	} else {
		goto L208
	}
L208:
	;
	goto L209
L209:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1412 == int32(46) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	F_field(m, l0, v21+int32(200))
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L3
	} else {
		goto L219
	}
L212:
	;
	if v1412 != int32(58) {
		v1423 = int32(0)
		goto L213
	} else {
		goto L214
	}
L213:
	;
	F_body(m, l0, v21+int32(776), v1423, v67)
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L3
	} else {
		goto L216
	}
L214:
	;
	F_field(m, l0, v21+int32(200))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L3
	} else {
		goto L215
	}
L215:
	;
	v1423 = int32(1)
	goto L213
L216:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_storevar(m, v1428, v21+int32(200), v21+int32(776))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L3
	} else {
		goto L217
	}
L217:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1435)))
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+20))
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1437+v1438<<(uint(int32(2))%32)+int32(-4)))) = v67
	goto L218
L218:
	;
	goto L12
L219:
	;
	goto L209
L220:
	;
	switch v1453 {
	case 0:
		goto L227
	default:
		goto L226
	}
L222:
	;
	v1453 = int32(0)
	v1457 = v1715
	goto L220
L223:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v21)+776))
	F_luaK_patchtohere(m, v70, v1744)
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L3
	} else {
		goto L267
	}
L224:
	;
	F_luaK_concat(m, v70, v21+int32(776), v1457)
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L3
	} else {
		goto L266
	}
L225:
	;
	v1453 = int32(1)
	v1457 = v1720
	goto L220
L226:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L3
	} else {
		goto L231
	}
L227:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1471 = F_luaX_token2str(m, l0, int32(274))
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L3
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v1471
	v1474 = m.G3
	v1477 = F_luaO_pushfstring(m, v1469, v1474+int32(_a2038), v21)
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L3
	} else {
		goto L229
	}
L229:
	;
	F_luaX_syntaxerror(m, l0, v1477)
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L3
	} else {
		goto L230
	}
L230:
	;
	v1720 = v1457
	goto L225
L231:
	;
	F_block(m, l0)
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L3
	} else {
		goto L232
	}
L232:
	;
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1485 == int32(261) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1692 = F_luaK_jump(m, v70)
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L3
	} else {
		goto L257
	}
L234:
	;
	if v1485 != int32(260) {
		goto L224
	} else {
		goto L235
	}
L235:
	;
	v1492 = F_luaK_jump(m, v70)
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L3
	} else {
		goto L236
	}
L236:
	;
	F_luaK_concat(m, v70, v21+int32(776), v1492)
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L3
	} else {
		goto L237
	}
L237:
	;
	F_luaK_patchtohere(m, v70, v1457)
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L3
	} else {
		goto L238
	}
L238:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L3
	} else {
		goto L239
	}
L239:
	;
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1501 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+210)) = uint8(v1501)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = int32(-1)
	v1505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1500)+50)))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+209)) = uint8(v1501)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+208)) = uint8(v1505)
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1500)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v1509
	*(*int32)(unsafe.Add(mBase, uint32(v1500)+20)) = v21 + int32(200)
	F_chunk(m, l0)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L3
	} else {
		goto L240
	}
L240:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1500)+20))
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v1516)))
	*(*int32)(unsafe.Add(mBase, uint32(v1500)+20)) = v1517
	v1519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1516)+8)))
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(v1500)+12))
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1520)+48))
	v1522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1521)+50)))
	if base.Ui32(v1522) <= base.Ui32(v1519) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1516)+9)))
	if v1677 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L242:
	;
	v1525 = v1521 + int32(172)
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1521)+24))
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1521)))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+24))
	v1529 = int32(0)
	v1532 = (v1522 - v1519) & int32(3)
	if v1532 == v1529 {
		v1568 = v1522
		goto L243
	} else {
		goto L244
	}
L243:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v1519-v1522) {
		v1642 = v1568
		goto L248
	} else {
		goto L249
	}
L244:
	;
	v1537 = v1522
	v1542 = v1529
	goto L245
L245:
	;
	v1554 = v1537 + int32(-1)
	v1555 = int32(1)
	v1558 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1525+v1554<<(uint(v1555)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1528+v1558*int32(12))+8)) = v1526
	v1564 = v1542 + v1555
	if v1564 != v1532 {
		v1537 = v1554
		v1542 = v1564
		goto L245
	} else {
		goto L247
	}
L246:
	;
	v1568 = v1554
	goto L243
L247:
	;
	goto L246
L248:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1521)+50)) = uint8(v1642)
	goto L241
L249:
	;
	v1589 = v1568
	goto L250
L250:
	;
	v1605 = int32(1)
	v1607 = v1589<<(uint(v1605)%32) + v1525
	v1610 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1607+int32(-2)))))
	v1611 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v1528+v1610*v1611)+8)) = v1526
	v1615 = int32(-4)
	v1617 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1607+v1615))))
	*(*int32)(unsafe.Add(mBase, uint32(v1528+v1617*v1611)+8)) = v1526
	v1624 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1607+int32(-6)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1528+v1624*v1611)+8)) = v1526
	v1630 = v1589 + v1615
	v1634 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1525+v1630<<(uint(v1605)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1528+v1634*v1611)+8)) = v1526
	if base.Ui32(v1519) < base.Ui32(v1630) {
		v1589 = v1630
		goto L250
	} else {
		goto L252
	}
L251:
	;
	v1642 = v1630
	goto L248
L252:
	;
	goto L251
L253:
	;
	v1685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1500)+50)))
	*(*int32)(unsafe.Add(mBase, uint32(v1500)+36)) = v1685
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1516)+4))
	F_luaK_patchtohere(m, v1500, v1687)
	mBase = m.M
	v1689 = m.ExcPending
	if v1689 != 0 {
		goto L3
	} else {
		goto L256
	}
L254:
	;
	v1681 = int32(0)
	v1683 = F_luaK_codeABC(m, v1500, int32(35), v1519, v1681, v1681)
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L3
	} else {
		goto L255
	}
L255:
	;
	goto L253
L256:
	;
	goto L223
L257:
	;
	F_luaK_concat(m, v70, v21+int32(776), v1692)
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L3
	} else {
		goto L258
	}
L258:
	;
	F_luaK_patchtohere(m, v70, v1457)
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L3
	} else {
		goto L259
	}
L259:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L3
	} else {
		goto L260
	}
L260:
	;
	v1703 = F_subexpr(m, l0, v21+int32(200), int32(0))
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L3
	} else {
		goto L261
	}
L261:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
	if v1705 != int32(1) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_goiftrue(m, v1710, v21+int32(200))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L3
	} else {
		goto L264
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = int32(3)
	goto L262
L264:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v21)+220))
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1716 != int32(274) {
		goto L222
	} else {
		goto L265
	}
L265:
	;
	v1720 = v1715
	goto L225
L266:
	;
	goto L223
L267:
	;
	F_check_match(m, l0, int32(262), int32(266), v67)
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L3
	} else {
		goto L268
	}
L268:
	;
	goto L12
L269:
	;
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v21)+208))
	if v1755 != int32(13) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = int32(0)
	F_assignment(m, l0, v21+int32(200), int32(1))
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		goto L3
	} else {
		goto L272
	}
L271:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1752)))
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+12))
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v21)+216))
	v1763 = v1759 + v1760<<(uint(int32(2))%32)
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1763)))
	*(*int32)(unsafe.Add(mBase, uint32(v1763))) = v1764&int32(-8372225) | int32(16384)
	goto L12
L272:
	;
	goto L12
L273:
	;
	v1779 = int32(0)
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1781)+20))
	if v1782 == v1779 {
		v1813 = v1779
		goto L275
	} else {
		goto L276
	}
L274:
	;
	if v1836 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L275:
	;
	v1827 = m.G3
	F_luaX_syntaxerror(m, l0, v1827+int32(_a2043))
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L3
	} else {
		goto L282
	}
L276:
	;
	v1787 = v1782
	v1789 = v1779
	goto L277
L277:
	;
	v1803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1787)+10)))
	if v1803 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	v1813 = v1807
	goto L275
L279:
	;
	v1806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1787)+9)))
	v1807 = v1789 | v1806
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1787)))
	if v1808 != 0 {
		v1787 = v1808
		v1789 = v1807
		goto L277
	} else {
		goto L281
	}
L280:
	;
	v1836 = v1789
	v1847 = v1787
	goto L274
L281:
	;
	goto L278
L282:
	;
	v1836 = v1813
	v1847 = v1779
	goto L274
L283:
	;
	v1860 = F_luaK_jump(m, v1781)
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L3
	} else {
		goto L286
	}
L284:
	;
	v1853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1847)+8)))
	v1854 = int32(0)
	v1856 = F_luaK_codeABC(m, v1781, int32(35), v1853, v1854, v1854)
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L3
	} else {
		goto L285
	}
L285:
	;
	goto L283
L286:
	;
	F_luaK_concat(m, v1781, v1847+int32(4), v1860)
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L3
	} else {
		goto L287
	}
L287:
	;
	v2813 = int32(0)
	goto L11
L288:
	;
	v1868 = int32(0)
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1871 = v1869 + int32(-260)
	if base.Ui32(int32(27)) < base.Ui32(v1871) {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	F_luaK_ret(m, v1865, v1997, v1993)
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L3
	} else {
		goto L313
	}
L290:
	;
	if v1869 == int32(59) {
		v1993 = int32(0)
		v1997 = v1868
		goto L289
	} else {
		goto L293
	}
L291:
	;
	if int32(1)<<(uint(v1871)%32)&int32(134283271) == int32(0) {
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v1993 = int32(0)
	v1997 = v1868
	goto L289
L293:
	;
	v1887 = F_subexpr(m, l0, v21+int32(200), int32(0))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L3
	} else {
		goto L294
	}
L294:
	;
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1890 == int32(44) {
		goto L297
	} else {
		goto L298
	}
L295:
	;
	F_luaK_exp2nextreg(m, v1865, v21+int32(200))
	mBase = m.M
	v1989 = m.ExcPending
	if v1989 != 0 {
		goto L3
	} else {
		goto L312
	}
L296:
	;
	F_luaK_setreturns(m, v1865, v21+int32(200), int32(-1))
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L3
	} else {
		goto L308
	}
L297:
	;
	v1905 = int32(1)
	goto L301
L298:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
	if base.Ui32(v1893+int32(-13)) <= base.Ui32(int32(1)) {
		goto L296
	} else {
		goto L299
	}
L299:
	;
	v1901 = F_luaK_exp2anyreg(m, v1865, v21+int32(200))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L3
	} else {
		goto L300
	}
L300:
	;
	v1993 = int32(1)
	v1997 = v1901
	goto L289
L301:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L3
	} else {
		goto L303
	}
L302:
	;
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
	if base.Ui32(int32(2)) <= base.Ui32(v1938+int32(-13)) {
		goto L295
	} else {
		goto L307
	}
L303:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_exp2nextreg(m, v1923, v21+int32(200))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L3
	} else {
		goto L304
	}
L304:
	;
	v1931 = F_subexpr(m, l0, v21+int32(200), int32(0))
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L3
	} else {
		goto L305
	}
L305:
	;
	v1934 = v1905 + int32(1)
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1935 == int32(44) {
		v1905 = v1934
		goto L301
	} else {
		goto L306
	}
L306:
	;
	goto L302
L307:
	;
	goto L296
L308:
	;
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
	if v1966 != int32(13) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1865)+50)))
	v1993 = int32(-1)
	v1997 = v1984
	goto L289
L310:
	;
	if v1890 == int32(44) {
		goto L309
	} else {
		goto L311
	}
L311:
	;
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(v1865)))
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v1971)+12))
	v1973 = *(*int32)(unsafe.Add(mBase, uint32(v21)+208))
	v1976 = v1972 + v1973<<(uint(int32(2))%32)
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(v1976)))
	*(*int32)(unsafe.Add(mBase, uint32(v1976))) = v1977&int32(-64) | int32(29)
	goto L309
L312:
	;
	v1990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1865)+50)))
	v1993 = v1934
	v1997 = v1990
	goto L289
L313:
	;
	v2813 = int32(0)
	goto L11
L314:
	;
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2014 != int32(265) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v2165 = v2014
	v2176 = int32(0)
	goto L345
L316:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L3
	} else {
		goto L317
	}
L317:
	;
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2020 == int32(285) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_luaX_next(m, l0)
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L3
	} else {
		goto L323
	}
L319:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2025 = F_luaX_token2str(m, l0, int32(285))
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L3
	} else {
		goto L320
	}
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+144)) = v2025
	v2028 = m.G3
	v2033 = F_luaO_pushfstring(m, v2023, v2028+int32(_a2038), v21+int32(144))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L3
	} else {
		goto L321
	}
L321:
	;
	F_luaX_syntaxerror(m, l0, v2033)
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L3
	} else {
		goto L322
	}
L322:
	;
	goto L318
L323:
	;
	F_new_localvar(m, l0, v2038, int32(0))
	mBase = m.M
	v2043 = m.ExcPending
	if v2043 != 0 {
		goto L3
	} else {
		goto L324
	}
L324:
	;
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+784)) = v2044
	*(*int32)(unsafe.Add(mBase, uint32(v21)+776)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+792)) = int64(-1)
	F_luaK_reserveregs(m, v2019, int32(1))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L3
	} else {
		goto L325
	}
L325:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2053)+50)))
	v2055 = int32(1)
	v2056 = v2054 + v2055
	*(*uint8)(unsafe.Add(mBase, uint32(v2053)+50)) = uint8(v2056)
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(v2053)))
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v2058)+24))
	v2067 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2053+v2056&int32(255)<<(uint(v2055)%32)+int32(170)))))
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v2053)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2059+v2067*int32(12))+4)) = v2071
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_open_func(m, l0, v21+int32(200))
	mBase = m.M
	v2077 = m.ExcPending
	if v2077 != 0 {
		goto L3
	} else {
		goto L326
	}
L326:
	;
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v2078)+60)) = v2073
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2080 == int32(40) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L3
	} else {
		goto L332
	}
L328:
	;
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2085 = F_luaX_token2str(m, l0, int32(40))
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		goto L3
	} else {
		goto L329
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v2085
	v2088 = m.G3
	v2093 = F_luaO_pushfstring(m, v2083, v2088+int32(_a2038), v21+int32(128))
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L3
	} else {
		goto L330
	}
L330:
	;
	F_luaX_syntaxerror(m, l0, v2093)
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L3
	} else {
		goto L331
	}
L331:
	;
	goto L327
L332:
	;
	F_parlist(m, l0)
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L3
	} else {
		goto L333
	}
L333:
	;
	v2102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2102 == int32(41) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v2121 = m.ExcPending
	if v2121 != 0 {
		goto L3
	} else {
		goto L339
	}
L335:
	;
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2107 = F_luaX_token2str(m, l0, int32(41))
	mBase = m.M
	v2108 = m.ExcPending
	if v2108 != 0 {
		goto L3
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v2107
	v2110 = m.G3
	v2115 = F_luaO_pushfstring(m, v2105, v2110+int32(_a2038), v21+int32(112))
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L3
	} else {
		goto L337
	}
L337:
	;
	F_luaX_syntaxerror(m, l0, v2115)
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L3
	} else {
		goto L338
	}
L338:
	;
	goto L334
L339:
	;
	F_chunk(m, l0)
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L3
	} else {
		goto L340
	}
L340:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2124)+64)) = v2125
	F_check_match(m, l0, int32(262), int32(265), v2073)
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L3
	} else {
		goto L341
	}
L341:
	;
	F_close_func(m, l0)
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L3
	} else {
		goto L342
	}
L342:
	;
	F_pushclosure(m, l0, v21+int32(200), v21+int32(176))
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L3
	} else {
		goto L343
	}
L343:
	;
	F_luaK_storevar(m, v2019, v21+int32(776), v21+int32(176))
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L3
	} else {
		goto L344
	}
L344:
	;
	v2145 = int32(1)
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v2019)))
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v2146)+24))
	v2148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2019)+50)))
	v2154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2148<<(uint(v2145)%32)+v2019+int32(170)))))
	v2158 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2147+v2154*int32(12))+4)) = v2158
	v2813 = v2145
	goto L11
L345:
	;
	if v2165 == int32(285) {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	if v2203 != int32(61) {
		goto L358
	} else {
		goto L359
	}
L347:
	;
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_luaX_next(m, l0)
	mBase = m.M
	v2198 = m.ExcPending
	if v2198 != 0 {
		goto L3
	} else {
		goto L352
	}
L348:
	;
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2183 = F_luaX_token2str(m, l0, int32(285))
	mBase = m.M
	v2184 = m.ExcPending
	if v2184 != 0 {
		goto L3
	} else {
		goto L349
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+160)) = v2183
	v2186 = m.G3
	v2191 = F_luaO_pushfstring(m, v2181, v2186+int32(_a2038), v21+int32(160))
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L3
	} else {
		goto L350
	}
L350:
	;
	F_luaX_syntaxerror(m, l0, v2191)
	mBase = m.M
	v2194 = m.ExcPending
	if v2194 != 0 {
		goto L3
	} else {
		goto L351
	}
L351:
	;
	goto L347
L352:
	;
	F_new_localvar(m, l0, v2196, v2176)
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L3
	} else {
		goto L353
	}
L353:
	;
	v2202 = v2176 + int32(1)
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2203 != int32(44) {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	goto L346
L355:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L3
	} else {
		goto L356
	}
L356:
	;
	v2208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v2165 = v2208
	v2176 = v2202
	goto L345
L357:
	;
	F_adjust_assign(m, l0, v2202, v2264, v21+int32(200))
	mBase = m.M
	v2281 = m.ExcPending
	if v2281 != 0 {
		goto L3
	} else {
		goto L369
	}
L358:
	;
	v2257 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v2257
	v2264 = v2257
	goto L357
L359:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L3
	} else {
		goto L360
	}
L360:
	;
	v2216 = F_subexpr(m, l0, v21+int32(200), int32(0))
	mBase = m.M
	v2217 = m.ExcPending
	if v2217 != 0 {
		goto L3
	} else {
		goto L361
	}
L361:
	;
	v2218 = int32(1)
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2219 != int32(44) {
		v2264 = v2218
		goto L357
	} else {
		goto L362
	}
L362:
	;
	v2226 = v2218
	goto L363
L363:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v2241 = m.ExcPending
	if v2241 != 0 {
		goto L3
	} else {
		goto L365
	}
L365:
	;
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_exp2nextreg(m, v2242, v21+int32(200))
	mBase = m.M
	v2246 = m.ExcPending
	if v2246 != 0 {
		goto L3
	} else {
		goto L366
	}
L366:
	;
	v2250 = F_subexpr(m, l0, v21+int32(200), int32(0))
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L3
	} else {
		goto L367
	}
L367:
	;
	v2253 = v2226 + int32(1)
	v2254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2254 == int32(44) {
		v2226 = v2253
		goto L363
	} else {
		goto L368
	}
L368:
	;
	v2264 = v2253
	goto L357
L369:
	;
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2282)+50)))
	v2284 = v2283 + v2202
	*(*uint8)(unsafe.Add(mBase, uint32(v2282)+50)) = uint8(v2284)
	v2287 = v2282 + int32(172)
	v2289 = v2284 & int32(255)
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v2282)+24))
	v2291 = *(*int32)(unsafe.Add(mBase, uint32(v2282)))
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(v2291)+24))
	v2293 = int32(0)
	v2295 = v2202 & int32(3)
	if v2295 == v2293 {
		v2332 = v2202
		goto L370
	} else {
		goto L371
	}
L370:
	;
	if base.Ui32(v2176) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L375
	}
L371:
	;
	v2300 = v2202
	v2314 = v2293
	goto L372
L372:
	;
	v2317 = int32(1)
	v2320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2287+(v2289-v2300)<<(uint(v2317)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2292+v2320*int32(12))+4)) = v2290
	v2326 = v2300 + int32(-1)
	v2328 = v2314 + v2317
	if v2328 != v2295 {
		v2300 = v2326
		v2314 = v2328
		goto L372
	} else {
		goto L374
	}
L373:
	;
	v2332 = v2326
	goto L370
L374:
	;
	goto L373
L375:
	;
	v2352 = v2332
	goto L376
L376:
	;
	v2371 = v2287 + (v2289-v2352)<<(uint(int32(1))%32)
	v2372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2371))))
	v2373 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v2292+v2372*v2373)+4)) = v2290
	v2379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2371+int32(2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2292+v2379*v2373)+4)) = v2290
	v2386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2371+int32(4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2292+v2386*v2373)+4)) = v2290
	v2393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2371+int32(6)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2292+v2393*v2373)+4)) = v2290
	v2399 = v2352 + int32(-4)
	if v2399 != 0 {
		v2352 = v2399
		goto L376
	} else {
		goto L378
	}
L378:
	;
	goto L12
L379:
	;
	F_new_localvar(m, l0, v2406, int32(0))
	mBase = m.M
	v2410 = m.ExcPending
	if v2410 != 0 {
		goto L3
	} else {
		goto L380
	}
L380:
	;
	v2415 = F_luaX_newstring(m, l0, v2400+int32(_a2044), int32(11))
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L3
	} else {
		goto L381
	}
L381:
	;
	F_new_localvar(m, l0, v2415, int32(1))
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L3
	} else {
		goto L382
	}
L382:
	;
	v2423 = F_luaX_newstring(m, l0, v2400+int32(_a2045), int32(13))
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L3
	} else {
		goto L383
	}
L383:
	;
	F_new_localvar(m, l0, v2423, int32(2))
	mBase = m.M
	v2427 = m.ExcPending
	if v2427 != 0 {
		goto L3
	} else {
		goto L384
	}
L384:
	;
	F_new_localvar(m, l0, v566, int32(3))
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L3
	} else {
		goto L385
	}
L385:
	;
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2432 != int32(44) {
		v2491 = v2432
		v2500 = int32(1)
		goto L386
	} else {
		goto L387
	}
L386:
	;
	if v2491 == int32(267) {
		goto L399
	} else {
		goto L400
	}
L387:
	;
	v2439 = int32(4)
	goto L388
L388:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		goto L3
	} else {
		goto L390
	}
L389:
	;
	v2491 = v2480
	v2500 = v2439 + int32(-2)
	goto L386
L390:
	;
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2455 == int32(285) {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_luaX_next(m, l0)
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L3
	} else {
		goto L396
	}
L392:
	;
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2460 = F_luaX_token2str(m, l0, int32(285))
	mBase = m.M
	v2461 = m.ExcPending
	if v2461 != 0 {
		goto L3
	} else {
		goto L393
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v2460
	v2463 = m.G3
	v2468 = F_luaO_pushfstring(m, v2458, v2463+int32(_a2038), v21+int32(80))
	mBase = m.M
	v2469 = m.ExcPending
	if v2469 != 0 {
		goto L3
	} else {
		goto L394
	}
L394:
	;
	F_luaX_syntaxerror(m, l0, v2468)
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L3
	} else {
		goto L395
	}
L395:
	;
	goto L391
L396:
	;
	F_new_localvar(m, l0, v2473, v2439)
	mBase = m.M
	v2477 = m.ExcPending
	if v2477 != 0 {
		goto L3
	} else {
		goto L397
	}
L397:
	;
	v2480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2480 == int32(44) {
		v2439 = v2439 + int32(1)
		goto L388
	} else {
		goto L398
	}
L398:
	;
	goto L389
L399:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v2521 = m.ExcPending
	if v2521 != 0 {
		goto L3
	} else {
		goto L404
	}
L400:
	;
	v2505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2507 = F_luaX_token2str(m, l0, int32(267))
	mBase = m.M
	v2508 = m.ExcPending
	if v2508 != 0 {
		goto L3
	} else {
		goto L401
	}
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v2507
	v2510 = m.G3
	v2515 = F_luaO_pushfstring(m, v2505, v2510+int32(_a2038), v21+int32(64))
	mBase = m.M
	v2516 = m.ExcPending
	if v2516 != 0 {
		goto L3
	} else {
		goto L402
	}
L402:
	;
	F_luaX_syntaxerror(m, l0, v2515)
	mBase = m.M
	v2518 = m.ExcPending
	if v2518 != 0 {
		goto L3
	} else {
		goto L403
	}
L403:
	;
	goto L399
L404:
	;
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2526 = F_subexpr(m, l0, v21+int32(200), int32(0))
	mBase = m.M
	v2527 = m.ExcPending
	if v2527 != 0 {
		goto L3
	} else {
		goto L405
	}
L405:
	;
	v2528 = int32(1)
	v2529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2529 != int32(44) {
		v2569 = v2528
		goto L406
	} else {
		goto L407
	}
L406:
	;
	F_adjust_assign(m, l0, int32(3), v2569, v21+int32(200))
	mBase = m.M
	v2589 = m.ExcPending
	if v2589 != 0 {
		goto L3
	} else {
		goto L414
	}
L407:
	;
	v2534 = v2528
	goto L408
L408:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v2551 = m.ExcPending
	if v2551 != 0 {
		goto L3
	} else {
		goto L410
	}
L409:
	;
	v2569 = v2563
	goto L406
L410:
	;
	v2552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_exp2nextreg(m, v2552, v21+int32(200))
	mBase = m.M
	v2556 = m.ExcPending
	if v2556 != 0 {
		goto L3
	} else {
		goto L411
	}
L411:
	;
	v2560 = F_subexpr(m, l0, v21+int32(200), int32(0))
	mBase = m.M
	v2561 = m.ExcPending
	if v2561 != 0 {
		goto L3
	} else {
		goto L412
	}
L412:
	;
	v2563 = v2534 + int32(1)
	v2564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2564 == int32(44) {
		v2534 = v2563
		goto L408
	} else {
		goto L413
	}
L413:
	;
	goto L409
L414:
	;
	F_luaK_checkstack(m, v2401, int32(3))
	mBase = m.M
	v2592 = m.ExcPending
	if v2592 != 0 {
		goto L3
	} else {
		goto L415
	}
L415:
	;
	F_forbody(m, l0, v2402, v2522, v2500, int32(0))
	mBase = m.M
	v2595 = m.ExcPending
	if v2595 != 0 {
		goto L3
	} else {
		goto L416
	}
L416:
	;
	goto L13
L417:
	;
	v2618 = *(*int32)(unsafe.Add(mBase, uint32(v532)+20))
	v2619 = *(*int32)(unsafe.Add(mBase, uint32(v2618)))
	*(*int32)(unsafe.Add(mBase, uint32(v532)+20)) = v2619
	v2621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2618)+8)))
	v2622 = *(*int32)(unsafe.Add(mBase, uint32(v532)+12))
	v2623 = *(*int32)(unsafe.Add(mBase, uint32(v2622)+48))
	v2624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2623)+50)))
	if base.Ui32(v2624) <= base.Ui32(v2621) {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v2779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2618)+9)))
	if v2779 == int32(0) {
		goto L430
	} else {
		goto L431
	}
L419:
	;
	v2627 = v2623 + int32(172)
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v2623)+24))
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v2623)))
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v2629)+24))
	v2631 = int32(0)
	v2634 = (v2624 - v2621) & int32(3)
	if v2634 == v2631 {
		v2670 = v2624
		goto L420
	} else {
		goto L421
	}
L420:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v2621-v2624) {
		v2744 = v2670
		goto L425
	} else {
		goto L426
	}
L421:
	;
	v2639 = v2624
	v2653 = v2631
	goto L422
L422:
	;
	v2656 = v2639 + int32(-1)
	v2657 = int32(1)
	v2660 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2627+v2656<<(uint(v2657)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2630+v2660*int32(12))+8)) = v2628
	v2666 = v2653 + v2657
	if v2666 != v2634 {
		v2639 = v2656
		v2653 = v2666
		goto L422
	} else {
		goto L424
	}
L423:
	;
	v2670 = v2656
	goto L420
L424:
	;
	goto L423
L425:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2623)+50)) = uint8(v2744)
	goto L418
L426:
	;
	v2691 = v2670
	goto L427
L427:
	;
	v2707 = int32(1)
	v2709 = v2691<<(uint(v2707)%32) + v2627
	v2712 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2709+int32(-2)))))
	v2713 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v2630+v2712*v2713)+8)) = v2628
	v2717 = int32(-4)
	v2719 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2709+v2717))))
	*(*int32)(unsafe.Add(mBase, uint32(v2630+v2719*v2713)+8)) = v2628
	v2726 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2709+int32(-6)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2630+v2726*v2713)+8)) = v2628
	v2732 = v2691 + v2717
	v2736 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2627+v2732<<(uint(v2707)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2630+v2736*v2713)+8)) = v2628
	if base.Ui32(v2621) < base.Ui32(v2732) {
		v2691 = v2732
		goto L427
	} else {
		goto L429
	}
L428:
	;
	v2744 = v2732
	goto L425
L429:
	;
	goto L428
L430:
	;
	v2787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532)+50)))
	*(*int32)(unsafe.Add(mBase, uint32(v532)+36)) = v2787
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(v2618)+4))
	F_luaK_patchtohere(m, v532, v2789)
	mBase = m.M
	v2791 = m.ExcPending
	if v2791 != 0 {
		goto L3
	} else {
		goto L433
	}
L431:
	;
	v2783 = int32(0)
	v2785 = F_luaK_codeABC(m, v532, int32(35), v2621, v2783, v2783)
	mBase = m.M
	v2786 = m.ExcPending
	if v2786 != 0 {
		goto L3
	} else {
		goto L432
	}
L432:
	;
	goto L430
L433:
	;
	goto L12
L434:
	;
	v2834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2834)+50)))
	*(*int32)(unsafe.Add(mBase, uint32(v2834)+36)) = v2835
	if v2813 != 0 {
		goto L6
	} else {
		goto L437
	}
L435:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v2833 = m.ExcPending
	if v2833 != 0 {
		goto L3
	} else {
		goto L436
	}
L436:
	;
	goto L434
L437:
	;
	goto L7
}
func F_clearFailoverState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v2 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[714])) = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[715])) = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[713]))
	F_valkey_free(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = int32(_a44)
		v13 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[658])) = v13
		*(*int64)(unsafe.Add(mBase, _consts[713])) = int64(0)
		if l0 == v13 {
			F_unpauseActions(m, int32(2))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				return
			}
		} else {
			F_disconnectOrRedirectAllBlockedClients(m)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_unpauseActions(m, int32(2))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_close(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v3 = m.Wasi_snapshot_preview1.Fd_close(m, l0)
	mBase = m.M
	if v3 == int32(27) {
		v6 = int32(0)
	} else {
		v6 = v3
	}
	if v6 != 0 {
		v8 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v6
		v11 = int32(-1)
	} else {
		v11 = int32(0)
	}
	return v11
}
func F_closeListeningSockets(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = *(*int32)(unsafe.Add(mBase, _consts[325]))
	if v10 == v2 {
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[326]))
		if v14 < int32(1) {
		} else {
			v19 = v2
			for {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v19<<(uint(int32(2))%32))+uint32(_consts[856])))
				v25 = F_close(m, v24)
				mBase = m.M
				v27 = v19 + int32(1)
				v29 = *(*int32)(unsafe.Add(mBase, _consts[326]))
				if v27 < v29 {
					v19 = v27
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v34 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[857]))
	if v36 == v34 {
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, _consts[858]))
		if v40 < int32(1) {
		} else {
			v45 = v34
			for {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v45<<(uint(int32(2))%32))+uint32(_consts[859])))
				v51 = F_close(m, v50)
				mBase = m.M
				v53 = v45 + int32(1)
				v55 = *(*int32)(unsafe.Add(mBase, _consts[858]))
				if v53 < v55 {
					v45 = v53
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v60 = int32(0)
	v62 = *(*int32)(unsafe.Add(mBase, _consts[860]))
	if v62 == v60 {
	} else {
		v66 = *(*int32)(unsafe.Add(mBase, _consts[861]))
		if v66 < int32(1) {
		} else {
			v71 = v60
			for {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v71<<(uint(int32(2))%32))+uint32(_consts[862])))
				v77 = F_close(m, v76)
				mBase = m.M
				v79 = v71 + int32(1)
				v81 = *(*int32)(unsafe.Add(mBase, _consts[861]))
				if v79 < v81 {
					v71 = v79
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v86 = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, _consts[863]))
	if v88 == v86 {
	} else {
		v92 = *(*int32)(unsafe.Add(mBase, _consts[864]))
		if v92 < int32(1) {
		} else {
			v97 = v86
			for {
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v97<<(uint(int32(2))%32))+uint32(_consts[865])))
				v103 = F_close(m, v102)
				mBase = m.M
				v105 = v97 + int32(1)
				v107 = *(*int32)(unsafe.Add(mBase, _consts[864]))
				if v105 < v107 {
					v97 = v105
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v112 = int32(0)
	v114 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v114 == v112 {
	} else {
		v118 = *(*int32)(unsafe.Add(mBase, _consts[866]))
		if v118 < int32(1) {
		} else {
			v123 = v112
			for {
				v128 = *(*int32)(unsafe.Add(mBase, uint32(v123<<(uint(int32(2))%32))+uint32(_consts[867])))
				v129 = F_close(m, v128)
				mBase = m.M
				v131 = v123 + int32(1)
				v133 = *(*int32)(unsafe.Add(mBase, _consts[866]))
				if v131 < v133 {
					v123 = v131
					continue
				} else {
					break
				}
				break
			}
		}
	}
	if l0 == int32(0) {
		m.G0 = v6 + int32(16)
		return
	} else {
		v140 = int32(0)
		v141 = *(*int32)(unsafe.Add(mBase, _consts[868]))
		if v141 == v140 {
			m.G0 = v6 + int32(16)
			return
		} else {
			v145 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(2) < v145 {
				v155 = v141
				v156 = F_unlink(m, v155)
				mBase = m.M
				if v156 == int32(0) {
					m.G0 = v6 + int32(16)
					return
				} else {
					v160 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					if int32(3) < v160 {
						m.G0 = v6 + int32(16)
						return
					} else {
						v164 = *(*int32)(unsafe.Add(mBase, _consts[5]))
						v165 = F___strerror_l(m, v164, v164)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v165
						F__serverLog(m, int32(3), int32(_a1576), v6)
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return
						} else {
							m.G0 = v6 + int32(16)
							return
						}
					}
				}
			} else {
				F__serverLog(m, int32(2), int32(_a1577), int32(0))
				mBase = m.M
				v152 = m.ExcPending
				if v152 != 0 {
					return
				} else {
					v154 = *(*int32)(unsafe.Add(mBase, _consts[868]))
					v155 = v154
					v156 = F_unlink(m, v155)
					mBase = m.M
					if v156 == int32(0) {
						m.G0 = v6 + int32(16)
						return
					} else {
						v160 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if int32(3) < v160 {
							m.G0 = v6 + int32(16)
							return
						} else {
							v164 = *(*int32)(unsafe.Add(mBase, _consts[5]))
							v165 = F___strerror_l(m, v164, v164)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v165
							F__serverLog(m, int32(3), int32(_a1576), v6)
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
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
func F_closeRepldbfd(m *base.Module, l0 int32) {
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
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v11 = v6 + int32(8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v12
	goto L1
L1:
	;
	v17 = v6 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = int32(-1)
	m.G0 = v6 + int32(16)
	return
L3:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v66 = F_close(m, v65)
	mBase = m.M
	goto L2
L4:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	F_bioCreateCloseJob(m, v59, int32(0), int32(1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	if v19 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19+base.B2i32(v22 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v28
	goto L6
L8:
	;
	v34 = v19
	goto L9
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	if v35 == l0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L4
L11:
	;
	v42 = v6 + int32(8)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v44 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+104))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v38 == int32(8) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if v44 != 0 {
		v34 = v44
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v44+base.B2i32(v47 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v53
	goto L15
L17:
	;
	goto L10
L18:
	;
	return
L19:
	;
	goto L2
}
func F_close_state(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_luaF_close(m, l0, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		F_luaC_freeall(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			v16 = F_luaM_realloc_(m, l0, v11, v12<<(uint(int32(2))%32), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+52))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+60))
				v21 = F_luaM_realloc_(m, l0, v18, v19, int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v4)+60)) = v23
					*(*int32)(unsafe.Add(mBase, uint32(v4)+52)) = v21
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v31 = F_luaM_realloc_(m, l0, v26, v27*int32(24), v23)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						v38 = F_luaM_realloc_(m, l0, v33, v34<<(uint(int32(4))%32), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
							v44 = m.T0[v43].(func(*base.Module, int32, int32, int32, int32) int32)(m, v40, l0, int32(376), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
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
func F_clusterscanCommand(m *base.Module, l0 int32) {
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
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
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v241 int64
	_ = v241
	var v247 int32
	_ = v247
	var v249 int64
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v263 int64
	_ = v263
	var v268 int64
	_ = v268
	var v272 int32
	_ = v272
	var v274 int64
	_ = v274
	var v276 int32
	_ = v276
	var v284 int64
	_ = v284
	var v295 int64
	_ = v295
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v321 int64
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v341 int64
	_ = v341
	var v348 int32
	_ = v348
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int64
	_ = v367
	var v368 int64
	_ = v368
	var v369 int64
	_ = v369
	var v370 int64
	_ = v370
	var v374 int64
	_ = v374
	var v379 int64
	_ = v379
	var v384 int64
	_ = v384
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v612 int64
	_ = v612
	var v613 int64
	_ = v613
	var v614 int64
	_ = v614
	var v615 int64
	_ = v615
	var v619 int64
	_ = v619
	var v624 int64
	_ = v624
	var v629 int64
	_ = v629
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v657 int32
	_ = v657
	var v665 int32
	_ = v665
	var v673 int32
	_ = v673
	var v681 int32
	_ = v681
	var v697 int64
	_ = v697
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v720 int32
	_ = v720
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(96)
	return
L2:
	;
	v23 = F_parseScanOptionsOrReply(m, l0, int32(0), int32(2), int32(1), v11+int32(48))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L6
	}
L3:
	;
	F_addReplyError(m, l0, int32(_a224))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	goto L1
L6:
	;
	if v23 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	v26 = int32(-1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	v33 = base.B2i32(v25 == v26) | base.B2i32(v28 == v26) | base.B2i32(v25 == v28)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v36 = F_objectGetVal(m, v35)
	mBase = m.M
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v37 != int32(48) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v75 = F_objectGetVal(m, v74)
	mBase = m.M
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+int32(-1)))))
	switch v78 & int32(7) {
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
		v95 = int32(0)
		goto L26
	}
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v40 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if v33 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	F_addReplyArrayLen(m, l0, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L25
	}
L13:
	;
	v49 = F_sdsempty(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	F_addReplyBulkCString(m, l0, int32(_a195))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	if v41 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v55 = int32(0)
	goto L19
L18:
	;
	v55 = v41
	goto L19
L19:
	;
	if v42 == int32(-1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v58 = v55
	goto L22
L21:
	;
	v58 = v42
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a225) + v58<<(uint(int32(2))%32)
	v64 = F_sdscatfmt(m, v49, int32(_a226), v11)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_addReplyBulkSds(m, l0, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	goto L12
L25:
	;
	goto L1
L26:
	;
	v96 = int32(45)
	v97 = F___strchrnul(m, v75, v96)
	mBase = m.M
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v99 == v96 {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-17))))
	v95 = v94
	goto L26
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-9))))
	v95 = v91
	goto L26
L29:
	;
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75+int32(-5)))))
	v95 = v88
	goto L26
L30:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+int32(-3)))))
	v95 = v85
	goto L26
L31:
	;
	v95 = int32(base.Ui32(v78) >> (uint(int32(3)) % 32))
	goto L26
L32:
	;
	F_addReplyError(m, l0, int32(_a227))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L4
	} else {
		goto L152
	}
L33:
	;
	if v103 == int32(0) {
		goto L32
	} else {
		goto L37
	}
L34:
	;
	v103 = v97
	goto L36
L35:
	;
	v103 = int32(0)
	goto L36
L36:
	;
	goto L33
L37:
	;
	v107 = v103 + int32(1)
	v108 = int32(45)
	v109 = F___strchrnul(m, v107, v108)
	mBase = m.M
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v111 == v108 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v115 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L39:
	;
	v115 = v109
	goto L41
L40:
	;
	v115 = int32(0)
	goto L41
L41:
	;
	goto L38
L42:
	;
	v118 = v115 - v107
	v119 = int32(0)
	if v118 < int32(1) {
		v139 = v119
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v187 = v115 + int32(1)
	v189 = v75 + v95 - v187
	v191 = v11 + int32(88)
	v199 = m.G0
	v201 = v199 - int32(16)
	m.G0 = v201
	if base.Ui32(v189+int32(-21)) < base.Ui32(int32(-20)) {
		goto L67
	} else {
		goto L68
	}
L44:
	;
	v185 = v181 & int32(16383)
	goto L43
L45:
	;
	v150 = v139 + int32(1)
	if v118 <= v150 {
		goto L55
	} else {
		goto L56
	}
L46:
	;
	v148 = F_crc16(m, v107, v118)
	mBase = m.M
	v181 = v148
	goto L44
L47:
	;
	if v139 != v118 {
		goto L45
	} else {
		goto L53
	}
L48:
	;
	v127 = v119
	goto L49
L49:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v127))))
	if v131 == int32(123) {
		v139 = v127
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v135 = v127 + int32(1)
	if v135 != v118 {
		v127 = v135
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L46
L53:
	;
	goto L46
L54:
	;
	v178 = F_crc16(m, v107+v139+int32(1), v156+(v139^int32(-1)))
	mBase = m.M
	v181 = v178
	goto L44
L55:
	;
	v171 = F_crc16(m, v107, v118)
	mBase = m.M
	v181 = v171
	goto L44
L56:
	;
	v156 = v150
	goto L58
L57:
	;
	if v156 == v118 {
		goto L55
	} else {
		goto L62
	}
L58:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v156))))
	if v158 == int32(125) {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v162 = v156 + int32(1)
	if v162 != v118 {
		v156 = v162
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L55
L62:
	;
	if v156 != v150 {
		goto L54
	} else {
		goto L63
	}
L63:
	;
	goto L55
L64:
	;
	if v348 == int32(0) {
		goto L32
	} else {
		goto L91
	}
L65:
	;
	m.G0 = v201 + int32(16)
	goto L64
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = v341
	v348 = int32(1)
	goto L65
L67:
	;
	v312 = int32(0)
	v313 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v201)+12)) = v312
	v321 = F_strtoull(m, v187, v201+int32(12), int32(10))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = v321
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	if v323 == int32(28) {
		v348 = v312
		goto L65
	} else {
		goto L88
	}
L68:
	;
	v207 = int32(1)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v189 != v207 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if v208&int32(255) != int32(45) {
		v228 = v207
		v229 = v208
		v230 = v187
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v212 = v208 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v212&int32(255)) {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v341 = base.I64_extend_i32_u(v212) & int64(255)
	goto L66
L72:
	;
	if base.Ui32(int32(8)) < base.Ui32((v229+int32(-49))&int32(255)) {
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v720 = int32(2)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	v228 = v720
	v229 = v226
	v230 = v115 + v720
	goto L72
L74:
	;
	v241 = base.I64_extend_i32_u(v229+int32(-48)) & int64(255)
	if base.Ui32(v189) <= base.Ui32(v228) {
		v284 = v241
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if v208&int32(255) != int32(45) {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	v247 = v228
	v249 = v241
	v251 = v230
	goto L77
L77:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+1)))
	if base.Ui32((v253+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L67
	} else {
		goto L79
	}
L78:
	;
	v284 = v274
	goto L75
L79:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v249) {
		goto L67
	} else {
		goto L80
	}
L80:
	;
	v263 = v249 * int64(10)
	v268 = base.I64_extend_i32_u(v253+int32(-48)) & int64(255)
	if base.Ui64(v268^int64(-1)) < base.Ui64(v263) {
		goto L67
	} else {
		goto L81
	}
L81:
	;
	v272 = int32(1)
	v274 = v263 + v268
	v276 = v247 + v272
	if v276 != v189 {
		v247 = v276
		v249 = v274
		v251 = v251 + v272
		goto L77
	} else {
		goto L82
	}
L82:
	;
	goto L78
L83:
	;
	if int64(0) <= v284 {
		v341 = v284
		goto L66
	} else {
		goto L87
	}
L84:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v284) {
		goto L67
	} else {
		goto L85
	}
L85:
	;
	v295 = int64(-1)
	if v295 < v284+v295 {
		v348 = int32(0)
		goto L65
	} else {
		goto L86
	}
L86:
	;
	v341 = int64(0)
	goto L66
L87:
	;
	goto L67
L88:
	;
	if v323 == int32(68) {
		v348 = v312
		goto L65
	} else {
		goto L89
	}
L89:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v328 == int32(0) {
		v348 = v312
		goto L65
	} else {
		goto L90
	}
L90:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	v348 = base.B2i32(v332 == int32(0))
	goto L65
L91:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, _consts[140])))
	if v364 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v103-v75 != int32(6) {
		goto L100
	} else {
		goto L101
	}
L93:
	;
	v366 = int32(_a228)
	goto L94
L94:
	;
	v367 = *(*int64)(unsafe.Add(mBase, _consts[141]))
	v368 = *(*int64)(unsafe.Add(mBase, _consts[142]))
	v369 = v367 ^ v368
	v370 = int64(21)
	v374 = v369<<(uint(v370)%64) + (v369 ^ int64(-1))
	v379 = (int64(base.Ui64(v374)>>(uint(int64(24))%64)) ^ v374) * int64(265)
	v384 = (int64(base.Ui64(v379)>>(uint(int64(14))%64)) ^ v379) * v370
	goto L95
L95:
	;
	v390 = base.I32_wrap_i64((int64(base.Ui64(v384)>>(uint(int64(28))%64)) ^ v384) * int64(2147483649))
	v391 = int32(1)
	if base.Ui32(v391) < base.Ui32(v390) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v394 = v390
	goto L98
L97:
	;
	v394 = v391
	goto L98
L98:
	;
	v395 = int32(63)
	v397 = int32(48)
	v398 = v394&v395 + v397
	*(*uint8)(unsafe.Add(mBase, _consts[143])) = uint8(v398)
	v400 = int32(0)
	v404 = int32(base.Ui32(v394)>>(uint(int32(30))%32)) | v397
	*(*uint8)(unsafe.Add(mBase, _consts[140])) = uint8(v404)
	v412 = int32(base.Ui32(v394)>>(uint(int32(6))%32))&v395 + v397
	*(*uint8)(unsafe.Add(mBase, _consts[144])) = uint8(v412)
	v420 = int32(base.Ui32(v394)>>(uint(int32(12))%32))&v395 + v397
	*(*uint8)(unsafe.Add(mBase, _consts[145])) = uint8(v420)
	v428 = int32(base.Ui32(v394)>>(uint(int32(18))%32))&v395 + v397
	*(*uint8)(unsafe.Add(mBase, _consts[146])) = uint8(v428)
	v436 = int32(base.Ui32(v394)>>(uint(int32(24))%32))&v395 + v397
	*(*uint8)(unsafe.Add(mBase, _consts[147])) = uint8(v436)
	*(*uint8)(unsafe.Add(mBase, _consts[148])) = uint8(v400)
	goto L92
L99:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	if v514 == int32(-1) {
		goto L119
	} else {
		goto L120
	}
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = int64(0)
	goto L99
L101:
	;
	v444 = int32(_a229)
	v445 = int32(6)
	goto L106
L102:
	;
	if v509 == int32(0) {
		goto L99
	} else {
		goto L118
	}
L103:
	;
	v509 = int32(0)
	goto L102
L104:
	;
	v481 = v476
	v482 = v477
	v483 = v478
	goto L114
L105:
	;
	if v466 == int32(0) {
		goto L103
	} else {
		goto L112
	}
L106:
	;
	if (v444|v75)&int32(3) != 0 {
		v476 = v75
		v477 = v444
		v478 = v445
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v453 = v75
	v454 = v444
	v455 = v445
	goto L108
L108:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	if v458 != v459 {
		v476 = v453
		v477 = v454
		v478 = v455
		goto L104
	} else {
		goto L110
	}
L109:
	;
	goto L105
L110:
	;
	v461 = int32(4)
	v462 = v454 + v461
	v464 = v453 + v461
	v466 = v455 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v466) {
		v453 = v464
		v454 = v462
		v455 = v466
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v476 = v464
	v477 = v462
	v478 = v466
	goto L104
L113:
	;
	v509 = v486 - v487
	goto L102
L114:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481))))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	if v486 != v487 {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v489 = int32(1)
	v494 = v483 + int32(-1)
	if v494 == int32(0) {
		goto L103
	} else {
		goto L117
	}
L117:
	;
	v481 = v481 + v489
	v482 = v482 + v489
	v483 = v494
	goto L114
L118:
	;
	goto L100
L119:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	if v521 == int32(-1) {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	if v185 == v514 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	F_addReplyError(m, l0, int32(_a230))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	goto L1
L123:
	;
	if v521&v514 == int32(-1) {
		goto L136
	} else {
		goto L137
	}
L124:
	;
	if v185 == v521 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	if v33&base.B2i32(v185 < v528) != int32(1) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	F_addReplyArrayLen(m, l0, int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L134
	}
L128:
	;
	F_addReplyBulkCString(m, l0, int32(_a195))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L4
	} else {
		goto L133
	}
L129:
	;
	v533 = F_sdsempty(m)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a225) + v536<<(uint(int32(2))%32)
	v544 = F_sdscatfmt(m, v533, int32(_a226), v11+int32(16))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	F_addReplyBulkSds(m, l0, v544)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	goto L127
L133:
	;
	goto L127
L134:
	;
	goto L1
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = int32(_a229)
	v697 = *(*int64)(unsafe.Add(mBase, uint32(v11)+88))
	F_scanGenericCommandWithOptions(m, l0, int32(0), v697, v11+int32(48), v11+int32(32))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L4
	} else {
		goto L151
	}
L136:
	;
	v569 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v569+v185<<(uint(int32(2))%32)+int32(52))))
	goto L138
L137:
	;
	v560 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(43)))) = uint8(v560)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+41)) = uint16(v560)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v560)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v185
	goto L135
L138:
	;
	v578 = v185
	goto L140
L139:
	;
	v598 = int32(0)
	v599 = int32(*(*uint8)(unsafe.Add(mBase, _consts[140])))
	*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(43)))) = uint8(v598)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+41)) = uint16(v598)
	v606 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v606)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v578
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v185
	if v599 != 0 {
		goto L135
	} else {
		goto L145
	}
L140:
	;
	if v578 == int32(16383) {
		goto L139
	} else {
		goto L142
	}
L141:
	;
	goto L139
L142:
	;
	v587 = v578 + int32(1)
	v589 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v589+v587<<(uint(int32(2))%32)+int32(52))))
	goto L143
L143:
	;
	if v595 == v575 {
		v578 = v587
		goto L140
	} else {
		goto L144
	}
L144:
	;
	goto L141
L145:
	;
	v611 = int32(_a228)
	goto L146
L146:
	;
	v612 = *(*int64)(unsafe.Add(mBase, _consts[141]))
	v613 = *(*int64)(unsafe.Add(mBase, _consts[142]))
	v614 = v612 ^ v613
	v615 = int64(21)
	v619 = v614<<(uint(v615)%64) + (v614 ^ int64(-1))
	v624 = (int64(base.Ui64(v619)>>(uint(int64(24))%64)) ^ v619) * int64(265)
	v629 = (int64(base.Ui64(v624)>>(uint(int64(14))%64)) ^ v624) * v615
	goto L147
L147:
	;
	v635 = base.I32_wrap_i64((int64(base.Ui64(v629)>>(uint(int64(28))%64)) ^ v629) * int64(2147483649))
	v636 = int32(1)
	if base.Ui32(v636) < base.Ui32(v635) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v639 = v635
	goto L150
L149:
	;
	v639 = v636
	goto L150
L150:
	;
	v640 = int32(63)
	v642 = int32(48)
	v643 = v639&v640 + v642
	*(*uint8)(unsafe.Add(mBase, _consts[143])) = uint8(v643)
	v645 = int32(0)
	v649 = int32(base.Ui32(v639)>>(uint(int32(30))%32)) | v642
	*(*uint8)(unsafe.Add(mBase, _consts[140])) = uint8(v649)
	v657 = int32(base.Ui32(v639)>>(uint(int32(6))%32))&v640 + v642
	*(*uint8)(unsafe.Add(mBase, _consts[144])) = uint8(v657)
	v665 = int32(base.Ui32(v639)>>(uint(int32(12))%32))&v640 + v642
	*(*uint8)(unsafe.Add(mBase, _consts[145])) = uint8(v665)
	v673 = int32(base.Ui32(v639)>>(uint(int32(18))%32))&v640 + v642
	*(*uint8)(unsafe.Add(mBase, _consts[146])) = uint8(v673)
	v681 = int32(base.Ui32(v639)>>(uint(int32(24))%32))&v640 + v642
	*(*uint8)(unsafe.Add(mBase, _consts[147])) = uint8(v681)
	*(*uint8)(unsafe.Add(mBase, _consts[148])) = uint8(v645)
	goto L135
L151:
	;
	goto L1
L152:
	;
	goto L1
}
func F_collateStringObjects(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_compareStringObjectsWithFlags(m, l0, l1, int32(2))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_commandlogInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v2 = F_listCreate(m)
	mBase = m.M
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		v4 = int32(_a44)
		*(*int64)(unsafe.Add(mBase, _consts[222])) = int64(0)
		*(*int32)(unsafe.Add(mBase, _consts[223])) = v2
		*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(111)
		v11 = F_listCreate(m)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = int32(_a44)
			*(*int64)(unsafe.Add(mBase, _consts[224])) = int64(0)
			*(*int32)(unsafe.Add(mBase, _consts[225])) = v11
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(111)
			v20 = F_listCreate(m)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = int32(_a44)
				*(*int64)(unsafe.Add(mBase, _consts[226])) = int64(0)
				*(*int32)(unsafe.Add(mBase, _consts[227])) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(111)
				return
			}
		}
	}
}
func F_condjump(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v20 = F_luaK_code(m, l0, l2<<(uint(int32(6))%32)|l1|l3<<(uint(int32(23))%32)|l4<<(uint(int32(14))%32), v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(-1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v30 = F_luaK_code(m, l0, int32(2147450902), v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v24 == int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v30
L5:
	;
	if v30 != int32(-1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v42 = v30
	goto L9
L7:
	;
	return v24
L8:
	;
	v65 = v24 + (v42 ^ int32(-1))
	v67 = v65 >> (uint(int32(31)) % 32)
	if base.Ui32(v65^v67-v67) < base.Ui32(int32(131072)) {
		v80 = v50
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v49 = v38 + v42<<(uint(int32(2))%32)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v54 = int32(base.Ui32(v50)>>(uint(int32(14))%32)) + int32(-131071)
	if v54 == int32(-1) {
		goto L8
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v59 = v42 + v54 + int32(1)
	if v59 != int32(-1) {
		v42 = v59
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v65<<(uint(int32(14))%32) | v80&int32(16383) + int32(2147467264)
	goto L4
L14:
	;
	v72 = m.G3
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v73, v72+int32(_a2034))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v80 = v78
	goto L13
}
func F_connect(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v7 = m.Env.X__syscall_connect(m, l0, l1, l2, v4, v4, v4)
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
func F_connectionByType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if base.Ui32(int32(4)) <= base.Ui32(l0) {
		F__serverAssert(m, int32(_a540), int32(_a538), int32(63))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[319])))
		if v14 != 0 {
			m.G0 = v6 + int32(16)
			return v14
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(3) < v16 {
				m.G0 = v6 + int32(16)
				return v14
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[320])))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v23
				F__serverLog(m, int32(3), int32(_a541), v6)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v14
				}
			}
		}
	}
}
func F_connectionTypeTcp(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[321]))
	if v8 != 0 {
		v34 = v8
		m.G0 = v5 + int32(16)
		return v34
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[319]))
		if v10 != 0 {
			*(*int32)(unsafe.Add(mBase, _consts[321])) = v10
			if v10 != 0 {
				v34 = v10
				m.G0 = v5 + int32(16)
				return v34
			} else {
				F__serverAssert(m, int32(_a542), int32(_a538), int32(80))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(3) < v12 {
				*(*int32)(unsafe.Add(mBase, _consts[321])) = v10
				if v10 != 0 {
					v34 = v10
					m.G0 = v5 + int32(16)
					return v34
				} else {
					F__serverAssert(m, int32(_a542), int32(_a538), int32(80))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a543)
				F__serverLog(m, int32(3), int32(_a541), v5)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[321])) = v23
					F__serverAssert(m, int32(_a542), int32(_a538), int32(80))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
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
		}
	}
}
func F_constructor(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int64
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
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
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
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
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v19 = F_luaK_codeABC(m, v14, int32(10), v3, v3, v3)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+100)) = int64(0)
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(11)
	v29 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v29
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_exp2nextreg(m, v37, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v40 == int32(123) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L9
	}
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v45 = F_luaX_token2str(m, l0, int32(123))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v45
	v48 = m.G3
	v53 = F_luaO_pushfstring(m, v43, v48+int32(_a2038), v11+int32(64))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_luaX_syntaxerror(m, l0, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L4
L9:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v60 == int32(125) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_check_match(m, l0, int32(125), int32(123), v13)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L51
	}
L11:
	;
	v64 = v60
	goto L12
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	if v71 == int32(0) {
		v92 = v64
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L10
L14:
	;
	if v92 == int32(91) {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	F_luaK_exp2nextreg(m, v14, v11+int32(72))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	if v80 != int32(50) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v92 = v91
	goto L14
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	F_luaK_setlist(m, v14, v84, v85, int32(50))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = int32(0)
	goto L17
L20:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v219 + int32(-44) {
	case 0, 15:
		goto L48
	default:
		goto L10
	}
L21:
	;
	v167 = F_subexpr(m, l0, v11+int32(72), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L39
	}
L22:
	;
	F_recfield(m, l0, v11+int32(72))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L38
	}
L23:
	;
	if v92 != int32(285) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	F_luaX_lookahead(m, l0)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v99 == int32(61) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_recfield(m, l0, v11+int32(72))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L37
	}
L27:
	;
	v105 = F_subexpr(m, l0, v11+int32(72), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	if v107 < int32(2147483646) {
		v146 = v107
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v149 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = v146 + v149
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v152 + v149
	goto L20
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+60))
	if v113 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	F_luaX_lexerror(m, v141, v139, int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L36
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(2147483645)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v113
	v129 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v129 + int32(_a2050)
	v137 = F_luaO_pushfstring(m, v111, v129+int32(_a2051), v11+int32(48))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(2147483645)
	v116 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v116 + int32(_a2050)
	v124 = F_luaO_pushfstring(m, v111, v116+int32(_a2052), v11+int32(32))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v139 = v124
	goto L31
L35:
	;
	v139 = v137
	goto L31
L36:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	v146 = v145
	goto L29
L37:
	;
	goto L20
L38:
	;
	goto L20
L39:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	if v169 < int32(2147483646) {
		v206 = v169
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v209 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = v206 + v209
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v212 + v209
	goto L20
L41:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+16))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+60))
	if v175 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	F_luaX_lexerror(m, v201, v199, int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L47
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(2147483645)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v175
	v189 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v189 + int32(_a2050)
	v197 = F_luaO_pushfstring(m, v173, v189+int32(_a2051), v11+int32(16))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(2147483645)
	v178 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v178 + int32(_a2050)
	v184 = F_luaO_pushfstring(m, v173, v178+int32(_a2052), v11)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v199 = v184
	goto L42
L46:
	;
	v199 = v197
	goto L42
L47:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	v206 = v205
	goto L40
L48:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v224 != int32(125) {
		v64 = v224
		goto L12
	} else {
		goto L50
	}
L50:
	;
	goto L13
L51:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	if v239 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	v273 = v19 << (uint(int32(2)) % 32)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v271+v273)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	if base.Ui32(int32(16)) <= base.Ui32(v276) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	switch v242 {
	case 0:
		v263 = v239
		goto L54
	default:
		goto L55
	case 13, 14:
		goto L56
	}
L54:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	F_luaK_setlist(m, v14, v265, v266, v263)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L60
	}
L55:
	;
	F_luaK_exp2nextreg(m, v14, v11+int32(72))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L59
	}
L56:
	;
	F_luaK_setreturns(m, v14, v11+int32(72), int32(-1))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	F_luaK_setlist(m, v14, v249, v250, int32(-1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = v254 + int32(-1)
	goto L52
L59:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	v263 = v262
	goto L54
L60:
	;
	goto L52
L61:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+12))
	v314 = v306<<(uint(int32(23))%32) | v275&int32(8388607)
	*(*int32)(unsafe.Add(mBase, uint32(v308+v273))) = v314
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v11)+100))
	if base.Ui32(int32(16)) <= base.Ui32(v316) {
		goto L73
	} else {
		goto L74
	}
L62:
	;
	if base.Ui32(v298) < base.Ui32(int32(8)) {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	v283 = v276
	v284 = int32(0)
	goto L65
L64:
	;
	v298 = v276
	v299 = int32(8)
	goto L62
L65:
	;
	v286 = int32(1)
	v287 = v284 + v286
	v289 = v283 + v286
	v291 = int32(base.Ui32(v289) >> (uint(v286) % 32))
	if base.Ui32(int32(31)) < base.Ui32(v289) {
		v283 = v291
		v284 = v287
		goto L65
	} else {
		goto L67
	}
L66:
	;
	v298 = v291
	v299 = v287<<(uint(int32(3))%32) + int32(8)
	goto L62
L67:
	;
	goto L66
L68:
	;
	v306 = v298
	goto L70
L69:
	;
	v306 = v299 | (v298 + int32(-8))
	goto L70
L70:
	;
	goto L61
L71:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v348+v273))) = v346<<(uint(int32(14))%32)&int32(8372224) | v314&int32(-8372225)
	m.G0 = v11 + int32(112)
	return
L72:
	;
	if base.Ui32(v338) < base.Ui32(int32(8)) {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	v323 = v316
	v324 = int32(0)
	goto L75
L74:
	;
	v338 = v316
	v339 = int32(8)
	goto L72
L75:
	;
	v326 = int32(1)
	v327 = v324 + v326
	v329 = v323 + v326
	v331 = int32(base.Ui32(v329) >> (uint(v326) % 32))
	if base.Ui32(int32(31)) < base.Ui32(v329) {
		v323 = v331
		v324 = v327
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v338 = v331
	v339 = v327<<(uint(int32(3))%32) + int32(8)
	goto L72
L77:
	;
	goto L76
L78:
	;
	v346 = v338
	goto L80
L79:
	;
	v346 = v339 | (v338 + int32(-8))
	goto L80
L80:
	;
	goto L71
}
func F_cosh(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v5 int64
	_ = v5
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v50 int32
	_ = v50
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v79 float64
	_ = v79
	var v82 float64
	_ = v82
	var v86 float64
	_ = v86
	var v89 float64
	_ = v89
	var v92 int64
	_ = v92
	var v97 int32
	_ = v97
	var v100 float64
	_ = v100
	var v103 float64
	_ = v103
	var v106 int64
	_ = v106
	var v109 int64
	_ = v109
	var v110 float64
	_ = v110
	var v111 float64
	_ = v111
	var v117 float64
	_ = v117
	var v128 float64
	_ = v128
	var v136 float64
	_ = v136
	var v140 float64
	_ = v140
	var v144 float64
	_ = v144
	v4 = base.F64_abs(l0)
	v5 = base.I64_reinterpret_f64(v4)
	if base.Ui64(int64(4604418530035630079)) < base.Ui64(v5) {
		if base.Ui64(int64(4649454526309335039)) < base.Ui64(v5) {
			v136 = float64(2.247116418577895e+307)
			v140 = F_exp(m, base.F64_add(v4, float64(-1416.0996898839683)))
			mBase = m.M
			v144 = base.F64_mul(base.F64_mul(base.F64_mul(float64(1), v136), v140), v136)
			return v144
		} else {
			v28 = F_top12_1(m, v4)
			mBase = m.M
			v30 = v28 & int32(2047)
			v32 = F_top12_1(m, float64(5.551115123125783e-17))
			mBase = m.M
			v35 = F_top12_1(m, float64(512))
			mBase = m.M
			if base.Ui32(v35-v32) <= base.Ui32(v30-v32) {
				if base.Ui32(v32) <= base.Ui32(v30) {
					v43 = F_top12_1(m, float64(1024))
					mBase = m.M
					if base.Ui32(v30) < base.Ui32(v43) {
						v60 = int32(0)
						v61 = int32(0)
						v62 = *(*float64)(unsafe.Add(mBase, _consts[1025]))
						v65 = *(*float64)(unsafe.Add(mBase, _consts[1026]))
						v66 = base.F64_add(base.F64_mul(v4, v62), v65)
						v67 = base.F64_sub(v66, v65)
						v69 = *(*float64)(unsafe.Add(mBase, _consts[1027]))
						v72 = *(*float64)(unsafe.Add(mBase, _consts[1028]))
						v75 = base.F64_add(base.F64_mul(v67, v69), base.F64_add(base.F64_mul(v67, v72), v4))
						v76 = base.F64_mul(v75, v75)
						v79 = *(*float64)(unsafe.Add(mBase, _consts[1029]))
						v82 = *(*float64)(unsafe.Add(mBase, _consts[1030]))
						v86 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
						v89 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
						v92 = base.I64_reinterpret_f64(v66)
						v97 = base.I32_wrap_i64(v92) << (uint(int32(4)) % 32) & int32(2032)
						v100 = *(*float64)(unsafe.Add(mBase, uint32(v97)+uint32(_consts[1033])))
						v103 = base.F64_add(base.F64_mul(base.F64_mul(v76, v76), base.F64_add(base.F64_mul(v75, v79), v82)), base.F64_add(base.F64_mul(v76, base.F64_add(base.F64_mul(v75, v86), v89)), base.F64_add(v100, v75)))
						v106 = *(*int64)(unsafe.Add(mBase, uint32(v97)+uint32(_consts[1034])))
						v109 = v106 + v92<<(uint(int64(45))%64)
						if v60 != 0 {
							v111 = base.F64_reinterpret_i64(v109)
							v117 = base.F64_add(base.F64_mul(v111, v103), v111)
							v128 = v117
						} else {
							v110 = F_specialcase_1(m, v103, v109, v92)
							mBase = m.M
							v128 = v110
						}
					} else {
						v46 = base.I64_reinterpret_f64(v4)
						if v46 == int64(-4503599627370496) {
							v117 = float64(0)
							v128 = v117
						} else {
							v50 = F_top12_1(m, math.Float64frombits(uint64(0x7ff0000000000000)))
							mBase = m.M
							if base.Ui32(v30) < base.Ui32(v50) {
								if int64(-1) < v46 {
									v59 = F___math_oflow(m, int32(0))
									mBase = m.M
									v128 = v59
								} else {
									v57 = F___math_uflow(m, int32(0))
									mBase = m.M
									v128 = v57
								}
							} else {
								v128 = base.F64_add(v4, float64(1))
							}
						}
					}
				} else {
					v128 = base.F64_add(v4, float64(1))
				}
			} else {
				v60 = v30
				v61 = int32(0)
				v62 = *(*float64)(unsafe.Add(mBase, _consts[1025]))
				v65 = *(*float64)(unsafe.Add(mBase, _consts[1026]))
				v66 = base.F64_add(base.F64_mul(v4, v62), v65)
				v67 = base.F64_sub(v66, v65)
				v69 = *(*float64)(unsafe.Add(mBase, _consts[1027]))
				v72 = *(*float64)(unsafe.Add(mBase, _consts[1028]))
				v75 = base.F64_add(base.F64_mul(v67, v69), base.F64_add(base.F64_mul(v67, v72), v4))
				v76 = base.F64_mul(v75, v75)
				v79 = *(*float64)(unsafe.Add(mBase, _consts[1029]))
				v82 = *(*float64)(unsafe.Add(mBase, _consts[1030]))
				v86 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
				v89 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
				v92 = base.I64_reinterpret_f64(v66)
				v97 = base.I32_wrap_i64(v92) << (uint(int32(4)) % 32) & int32(2032)
				v100 = *(*float64)(unsafe.Add(mBase, uint32(v97)+uint32(_consts[1033])))
				v103 = base.F64_add(base.F64_mul(base.F64_mul(v76, v76), base.F64_add(base.F64_mul(v75, v79), v82)), base.F64_add(base.F64_mul(v76, base.F64_add(base.F64_mul(v75, v86), v89)), base.F64_add(v100, v75)))
				v106 = *(*int64)(unsafe.Add(mBase, uint32(v97)+uint32(_consts[1034])))
				v109 = v106 + v92<<(uint(int64(45))%64)
				if v60 != 0 {
					v111 = base.F64_reinterpret_i64(v109)
					v117 = base.F64_add(base.F64_mul(v111, v103), v111)
					v128 = v117
				} else {
					v110 = F_specialcase_1(m, v103, v109, v92)
					mBase = m.M
					v128 = v110
				}
			}
			return base.F64_mul(base.F64_add(v128, base.F64_div(float64(1), v128)), float64(0.5))
		}
	} else {
		if base.Ui64(v5) < base.Ui64(int64(4490088828488384512)) {
			v144 = float64(1)
			return v144
		} else {
			v11 = F_expm1(m, v4)
			mBase = m.M
			v13 = float64(1)
			v14 = base.F64_add(v11, v13)
			return base.F64_add(base.F64_div(base.F64_mul(v11, v11), base.F64_add(v14, v14)), v13)
		}
	}
}
func F_createSocketAcceptHandler(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v7 < int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v62
L2:
	;
	v62 = int32(0)
	goto L1
L3:
	;
	v12 = int32(0)
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+v12<<(uint(int32(2))%32))))
	v22 = F_aeCreateFileEvent(m, v16, v20, int32(1), l1, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L2
L6:
	;
	v50 = v12 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v50 < v51 {
		v12 = v50
		goto L4
	} else {
		goto L16
	}
L7:
	;
	return int32(0)
L8:
	;
	if v22 != int32(-1) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v12 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v32 = v12
	goto L12
L11:
	;
	return int32(-1)
L12:
	;
	v35 = int32(-1)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v39 = v32 + v35
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0+v39<<(uint(int32(2))%32))))
	F_aeDeleteFileEvent(m, v37, v43, int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	if int32(1) < v32 {
		v32 = v39
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v62 = v35
	goto L1
L16:
	;
	goto L5
}
func F_createSparklineSequence(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	v3 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+16)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = v7
		*(*int64)(unsafe.Add(mBase, uint32(v3+int32(24)))) = v7
		return v3
	}
}
func F_cronUpdateMemoryStats(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
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
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v404 int32
	_ = v404
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	v1 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v16 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	if v16 < int32(261) {
		if v16 < int32(1) {
			v93 = v1
		} else {
			v24 = v1
			v25 = v16
			v27 = v25 & int32(3)
			if base.Ui32(int32(4)) <= base.Ui32(v25) {
				v34 = int32(0)
				v36 = v24
				v37 = v34
				v41 = v34
				for {
					v44 = v37 << (uint(int32(2)) % 32)
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[282])))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[283])))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[284])))
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[285])))
					v60 = v47 + (v50 + (v53 + (v56 + v36)))
					v61 = int32(4)
					v62 = v37 + v61
					v64 = v41 + v61
					if v64 != v25&int32(2147483644) {
						v36 = v60
						v37 = v62
						v41 = v64
						continue
					} else {
						break
					}
					break
				}
				v66 = v60
				v67 = v62
			} else {
				v66 = v24
				v67 = int32(0)
			}
			if v27 == int32(0) {
				v93 = v66
			} else {
				v75 = v66
				v76 = v67
				v78 = int32(0)
				for {
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v76<<(uint(int32(2))%32))+uint32(_consts[285])))
					v87 = v86 + v75
					v88 = int32(1)
					v91 = v78 + v88
					if v91 != v27 {
						v75 = v87
						v76 = v76 + v88
						v78 = v91
						continue
					} else {
						break
					}
					break
				}
				v93 = v87
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[286]))
		v24 = v20
		v25 = int32(260)
		v27 = v25 & int32(3)
		if base.Ui32(int32(4)) <= base.Ui32(v25) {
			v34 = int32(0)
			v36 = v24
			v37 = v34
			v41 = v34
			for {
				v44 = v37 << (uint(int32(2)) % 32)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[282])))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[283])))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[284])))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[285])))
				v60 = v47 + (v50 + (v53 + (v56 + v36)))
				v61 = int32(4)
				v62 = v37 + v61
				v64 = v41 + v61
				if v64 != v25&int32(2147483644) {
					v36 = v60
					v37 = v62
					v41 = v64
					continue
				} else {
					break
				}
				break
			}
			v66 = v60
			v67 = v62
		} else {
			v66 = v24
			v67 = int32(0)
		}
		if v27 == int32(0) {
			v93 = v66
		} else {
			v75 = v66
			v76 = v67
			v78 = int32(0)
			for {
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v76<<(uint(int32(2))%32))+uint32(_consts[285])))
				v87 = v86 + v75
				v88 = int32(1)
				v91 = v78 + v88
				if v91 != v27 {
					v75 = v87
					v76 = v76 + v88
					v78 = v91
					continue
				} else {
					break
				}
				break
			}
			v93 = v87
		}
	}
	v101 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	if base.Ui32(v93) <= base.Ui32(v101) {
	} else {
		v103 = int32(0)
		v113 = *(*int32)(unsafe.Add(mBase, _consts[281]))
		if v113 < int32(261) {
			if v113 < int32(1) {
				v190 = v103
			} else {
				v121 = v103
				v122 = v113
				v124 = v122 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(v122) {
					v131 = int32(0)
					v133 = v121
					v134 = v131
					v138 = v131
					for {
						v141 = v134 << (uint(int32(2)) % 32)
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[282])))
						v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[283])))
						v150 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[284])))
						v153 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[285])))
						v157 = v144 + (v147 + (v150 + (v153 + v133)))
						v158 = int32(4)
						v159 = v134 + v158
						v161 = v138 + v158
						if v161 != v122&int32(2147483644) {
							v133 = v157
							v134 = v159
							v138 = v161
							continue
						} else {
							break
						}
						break
					}
					v163 = v157
					v164 = v159
				} else {
					v163 = v121
					v164 = int32(0)
				}
				if v124 == int32(0) {
					v190 = v163
				} else {
					v172 = v163
					v173 = v164
					v175 = int32(0)
					for {
						v183 = *(*int32)(unsafe.Add(mBase, uint32(v173<<(uint(int32(2))%32))+uint32(_consts[285])))
						v184 = v183 + v172
						v185 = int32(1)
						v188 = v175 + v185
						if v188 != v124 {
							v172 = v184
							v173 = v173 + v185
							v175 = v188
							continue
						} else {
							break
						}
						break
					}
					v190 = v184
				}
			}
		} else {
			v117 = *(*int32)(unsafe.Add(mBase, _consts[286]))
			v121 = v117
			v122 = int32(260)
			v124 = v122 & int32(3)
			if base.Ui32(int32(4)) <= base.Ui32(v122) {
				v131 = int32(0)
				v133 = v121
				v134 = v131
				v138 = v131
				for {
					v141 = v134 << (uint(int32(2)) % 32)
					v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[282])))
					v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[283])))
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[284])))
					v153 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[285])))
					v157 = v144 + (v147 + (v150 + (v153 + v133)))
					v158 = int32(4)
					v159 = v134 + v158
					v161 = v138 + v158
					if v161 != v122&int32(2147483644) {
						v133 = v157
						v134 = v159
						v138 = v161
						continue
					} else {
						break
					}
					break
				}
				v163 = v157
				v164 = v159
			} else {
				v163 = v121
				v164 = int32(0)
			}
			if v124 == int32(0) {
				v190 = v163
			} else {
				v172 = v163
				v173 = v164
				v175 = int32(0)
				for {
					v183 = *(*int32)(unsafe.Add(mBase, uint32(v173<<(uint(int32(2))%32))+uint32(_consts[285])))
					v184 = v183 + v172
					v185 = int32(1)
					v188 = v175 + v185
					if v188 != v124 {
						v172 = v184
						v173 = v173 + v185
						v175 = v188
						continue
					} else {
						break
					}
					break
				}
				v190 = v184
			}
		}
		*(*int32)(unsafe.Add(mBase, _consts[71])) = v190
	}
	v200 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v201 = base.I32_div_s(int32(1000), v200)
	if int32(99) < v201 {
		v211 = int32(0)
		v221 = *(*int32)(unsafe.Add(mBase, _consts[281]))
		if v221 < int32(261) {
			if v221 < int32(1) {
				v298 = v211
			} else {
				v229 = v211
				v230 = v221
				v232 = v230 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(v230) {
					v239 = int32(0)
					v241 = v229
					v242 = v239
					v246 = v239
					for {
						v249 = v242 << (uint(int32(2)) % 32)
						v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[282])))
						v255 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[283])))
						v258 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[284])))
						v261 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[285])))
						v265 = v252 + (v255 + (v258 + (v261 + v241)))
						v266 = int32(4)
						v267 = v242 + v266
						v269 = v246 + v266
						if v269 != v230&int32(2147483644) {
							v241 = v265
							v242 = v267
							v246 = v269
							continue
						} else {
							break
						}
						break
					}
					v271 = v265
					v272 = v267
				} else {
					v271 = v229
					v272 = int32(0)
				}
				if v232 == int32(0) {
					v298 = v271
				} else {
					v280 = v271
					v281 = v272
					v283 = int32(0)
					for {
						v291 = *(*int32)(unsafe.Add(mBase, uint32(v281<<(uint(int32(2))%32))+uint32(_consts[285])))
						v292 = v291 + v280
						v293 = int32(1)
						v296 = v283 + v293
						if v296 != v232 {
							v280 = v292
							v281 = v281 + v293
							v283 = v296
							continue
						} else {
							break
						}
						break
					}
					v298 = v292
				}
			}
		} else {
			v225 = *(*int32)(unsafe.Add(mBase, _consts[286]))
			v229 = v225
			v230 = int32(260)
			v232 = v230 & int32(3)
			if base.Ui32(int32(4)) <= base.Ui32(v230) {
				v239 = int32(0)
				v241 = v229
				v242 = v239
				v246 = v239
				for {
					v249 = v242 << (uint(int32(2)) % 32)
					v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[282])))
					v255 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[283])))
					v258 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[284])))
					v261 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[285])))
					v265 = v252 + (v255 + (v258 + (v261 + v241)))
					v266 = int32(4)
					v267 = v242 + v266
					v269 = v246 + v266
					if v269 != v230&int32(2147483644) {
						v241 = v265
						v242 = v267
						v246 = v269
						continue
					} else {
						break
					}
					break
				}
				v271 = v265
				v272 = v267
			} else {
				v271 = v229
				v272 = int32(0)
			}
			if v232 == int32(0) {
				v298 = v271
			} else {
				v280 = v271
				v281 = v272
				v283 = int32(0)
				for {
					v291 = *(*int32)(unsafe.Add(mBase, uint32(v281<<(uint(int32(2))%32))+uint32(_consts[285])))
					v292 = v291 + v280
					v293 = int32(1)
					v296 = v283 + v293
					if v296 != v232 {
						v280 = v292
						v281 = v281 + v293
						v283 = v296
						continue
					} else {
						break
					}
					break
				}
				v298 = v292
			}
		}
		*(*int32)(unsafe.Add(mBase, _consts[810])) = v298
		v306 = int32(0)
		v316 = *(*int32)(unsafe.Add(mBase, _consts[281]))
		if v316 < int32(261) {
			if v316 < int32(1) {
				v393 = v306
			} else {
				v324 = v306
				v325 = v316
				v327 = v325 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(v325) {
					v334 = int32(0)
					v336 = v324
					v337 = v334
					v341 = v334
					for {
						v344 = v337 << (uint(int32(2)) % 32)
						v347 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[282])))
						v350 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[283])))
						v353 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[284])))
						v356 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[285])))
						v360 = v347 + (v350 + (v353 + (v356 + v336)))
						v361 = int32(4)
						v362 = v337 + v361
						v364 = v341 + v361
						if v364 != v325&int32(2147483644) {
							v336 = v360
							v337 = v362
							v341 = v364
							continue
						} else {
							break
						}
						break
					}
					v366 = v360
					v367 = v362
				} else {
					v366 = v324
					v367 = int32(0)
				}
				if v327 == int32(0) {
					v393 = v366
				} else {
					v375 = v366
					v376 = v367
					v378 = int32(0)
					for {
						v386 = *(*int32)(unsafe.Add(mBase, uint32(v376<<(uint(int32(2))%32))+uint32(_consts[285])))
						v387 = v386 + v375
						v388 = int32(1)
						v391 = v378 + v388
						if v391 != v327 {
							v375 = v387
							v376 = v376 + v388
							v378 = v391
							continue
						} else {
							break
						}
						break
					}
					v393 = v387
				}
			}
		} else {
			v320 = *(*int32)(unsafe.Add(mBase, _consts[286]))
			v324 = v320
			v325 = int32(260)
			v327 = v325 & int32(3)
			if base.Ui32(int32(4)) <= base.Ui32(v325) {
				v334 = int32(0)
				v336 = v324
				v337 = v334
				v341 = v334
				for {
					v344 = v337 << (uint(int32(2)) % 32)
					v347 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[282])))
					v350 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[283])))
					v353 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[284])))
					v356 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[285])))
					v360 = v347 + (v350 + (v353 + (v356 + v336)))
					v361 = int32(4)
					v362 = v337 + v361
					v364 = v341 + v361
					if v364 != v325&int32(2147483644) {
						v336 = v360
						v337 = v362
						v341 = v364
						continue
					} else {
						break
					}
					break
				}
				v366 = v360
				v367 = v362
			} else {
				v366 = v324
				v367 = int32(0)
			}
			if v327 == int32(0) {
				v393 = v366
			} else {
				v375 = v366
				v376 = v367
				v378 = int32(0)
				for {
					v386 = *(*int32)(unsafe.Add(mBase, uint32(v376<<(uint(int32(2))%32))+uint32(_consts[285])))
					v387 = v386 + v375
					v388 = int32(1)
					v391 = v378 + v388
					if v391 != v327 {
						v375 = v387
						v376 = v376 + v388
						v378 = v391
						continue
					} else {
						break
					}
					break
				}
				v393 = v387
			}
		}
		*(*int32)(unsafe.Add(mBase, _consts[811])) = v393
		v404 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[611])) = v404
		*(*int32)(unsafe.Add(mBase, _consts[612])) = v404
		*(*int32)(unsafe.Add(mBase, _consts[610])) = v404
		*(*int32)(unsafe.Add(mBase, _consts[613])) = int32(0)
		v421 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[812])) = v421
		v425 = *(*int32)(unsafe.Add(mBase, _consts[612]))
		if v425 != 0 {
			v439 = v425
			v441 = *(*int32)(unsafe.Add(mBase, _consts[611]))
			if v441 != 0 {
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[611])) = v439
			}
			v445 = *(*int32)(unsafe.Add(mBase, _consts[610]))
			if v445 != 0 {
			} else {
				v446 = int32(0)
				v448 = *(*int32)(unsafe.Add(mBase, _consts[811]))
				*(*int32)(unsafe.Add(mBase, _consts[610])) = v448
			}
			m.G0 = v5 + int32(16)
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(0)
			F_scriptingEngineManagerForEachEngine(m, int32(1022), v5+int32(12))
			mBase = m.M
			v432 = m.ExcPending
			if v432 != 0 {
				return
			} else {
				v433 = int32(0)
				v435 = *(*int32)(unsafe.Add(mBase, _consts[810]))
				v436 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				v437 = v435 - v436
				*(*int32)(unsafe.Add(mBase, _consts[612])) = v437
				v439 = v437
				v441 = *(*int32)(unsafe.Add(mBase, _consts[611]))
				if v441 != 0 {
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[611])) = v439
				}
				v445 = *(*int32)(unsafe.Add(mBase, _consts[610]))
				if v445 != 0 {
				} else {
					v446 = int32(0)
					v448 = *(*int32)(unsafe.Add(mBase, _consts[811]))
					*(*int32)(unsafe.Add(mBase, _consts[610])) = v448
				}
				m.G0 = v5 + int32(16)
				return
			}
		}
	} else {
		v205 = *(*int32)(unsafe.Add(mBase, _consts[220]))
		v208 = base.I32_div_s(int32(100), base.I32_extend16_s(v201))
		v210 = base.I32_rem_s(v205, base.I32_extend16_s(v208))
		if v210 != 0 {
			m.G0 = v5 + int32(16)
			return
		} else {
			v211 = int32(0)
			v221 = *(*int32)(unsafe.Add(mBase, _consts[281]))
			if v221 < int32(261) {
				if v221 < int32(1) {
					v298 = v211
				} else {
					v229 = v211
					v230 = v221
					v232 = v230 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(v230) {
						v239 = int32(0)
						v241 = v229
						v242 = v239
						v246 = v239
						for {
							v249 = v242 << (uint(int32(2)) % 32)
							v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[282])))
							v255 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[283])))
							v258 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[284])))
							v261 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[285])))
							v265 = v252 + (v255 + (v258 + (v261 + v241)))
							v266 = int32(4)
							v267 = v242 + v266
							v269 = v246 + v266
							if v269 != v230&int32(2147483644) {
								v241 = v265
								v242 = v267
								v246 = v269
								continue
							} else {
								break
							}
							break
						}
						v271 = v265
						v272 = v267
					} else {
						v271 = v229
						v272 = int32(0)
					}
					if v232 == int32(0) {
						v298 = v271
					} else {
						v280 = v271
						v281 = v272
						v283 = int32(0)
						for {
							v291 = *(*int32)(unsafe.Add(mBase, uint32(v281<<(uint(int32(2))%32))+uint32(_consts[285])))
							v292 = v291 + v280
							v293 = int32(1)
							v296 = v283 + v293
							if v296 != v232 {
								v280 = v292
								v281 = v281 + v293
								v283 = v296
								continue
							} else {
								break
							}
							break
						}
						v298 = v292
					}
				}
			} else {
				v225 = *(*int32)(unsafe.Add(mBase, _consts[286]))
				v229 = v225
				v230 = int32(260)
				v232 = v230 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(v230) {
					v239 = int32(0)
					v241 = v229
					v242 = v239
					v246 = v239
					for {
						v249 = v242 << (uint(int32(2)) % 32)
						v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[282])))
						v255 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[283])))
						v258 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[284])))
						v261 = *(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_consts[285])))
						v265 = v252 + (v255 + (v258 + (v261 + v241)))
						v266 = int32(4)
						v267 = v242 + v266
						v269 = v246 + v266
						if v269 != v230&int32(2147483644) {
							v241 = v265
							v242 = v267
							v246 = v269
							continue
						} else {
							break
						}
						break
					}
					v271 = v265
					v272 = v267
				} else {
					v271 = v229
					v272 = int32(0)
				}
				if v232 == int32(0) {
					v298 = v271
				} else {
					v280 = v271
					v281 = v272
					v283 = int32(0)
					for {
						v291 = *(*int32)(unsafe.Add(mBase, uint32(v281<<(uint(int32(2))%32))+uint32(_consts[285])))
						v292 = v291 + v280
						v293 = int32(1)
						v296 = v283 + v293
						if v296 != v232 {
							v280 = v292
							v281 = v281 + v293
							v283 = v296
							continue
						} else {
							break
						}
						break
					}
					v298 = v292
				}
			}
			*(*int32)(unsafe.Add(mBase, _consts[810])) = v298
			v306 = int32(0)
			v316 = *(*int32)(unsafe.Add(mBase, _consts[281]))
			if v316 < int32(261) {
				if v316 < int32(1) {
					v393 = v306
				} else {
					v324 = v306
					v325 = v316
					v327 = v325 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(v325) {
						v334 = int32(0)
						v336 = v324
						v337 = v334
						v341 = v334
						for {
							v344 = v337 << (uint(int32(2)) % 32)
							v347 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[282])))
							v350 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[283])))
							v353 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[284])))
							v356 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[285])))
							v360 = v347 + (v350 + (v353 + (v356 + v336)))
							v361 = int32(4)
							v362 = v337 + v361
							v364 = v341 + v361
							if v364 != v325&int32(2147483644) {
								v336 = v360
								v337 = v362
								v341 = v364
								continue
							} else {
								break
							}
							break
						}
						v366 = v360
						v367 = v362
					} else {
						v366 = v324
						v367 = int32(0)
					}
					if v327 == int32(0) {
						v393 = v366
					} else {
						v375 = v366
						v376 = v367
						v378 = int32(0)
						for {
							v386 = *(*int32)(unsafe.Add(mBase, uint32(v376<<(uint(int32(2))%32))+uint32(_consts[285])))
							v387 = v386 + v375
							v388 = int32(1)
							v391 = v378 + v388
							if v391 != v327 {
								v375 = v387
								v376 = v376 + v388
								v378 = v391
								continue
							} else {
								break
							}
							break
						}
						v393 = v387
					}
				}
			} else {
				v320 = *(*int32)(unsafe.Add(mBase, _consts[286]))
				v324 = v320
				v325 = int32(260)
				v327 = v325 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(v325) {
					v334 = int32(0)
					v336 = v324
					v337 = v334
					v341 = v334
					for {
						v344 = v337 << (uint(int32(2)) % 32)
						v347 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[282])))
						v350 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[283])))
						v353 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[284])))
						v356 = *(*int32)(unsafe.Add(mBase, uint32(v344)+uint32(_consts[285])))
						v360 = v347 + (v350 + (v353 + (v356 + v336)))
						v361 = int32(4)
						v362 = v337 + v361
						v364 = v341 + v361
						if v364 != v325&int32(2147483644) {
							v336 = v360
							v337 = v362
							v341 = v364
							continue
						} else {
							break
						}
						break
					}
					v366 = v360
					v367 = v362
				} else {
					v366 = v324
					v367 = int32(0)
				}
				if v327 == int32(0) {
					v393 = v366
				} else {
					v375 = v366
					v376 = v367
					v378 = int32(0)
					for {
						v386 = *(*int32)(unsafe.Add(mBase, uint32(v376<<(uint(int32(2))%32))+uint32(_consts[285])))
						v387 = v386 + v375
						v388 = int32(1)
						v391 = v378 + v388
						if v391 != v327 {
							v375 = v387
							v376 = v376 + v388
							v378 = v391
							continue
						} else {
							break
						}
						break
					}
					v393 = v387
				}
			}
			*(*int32)(unsafe.Add(mBase, _consts[811])) = v393
			v404 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[611])) = v404
			*(*int32)(unsafe.Add(mBase, _consts[612])) = v404
			*(*int32)(unsafe.Add(mBase, _consts[610])) = v404
			*(*int32)(unsafe.Add(mBase, _consts[613])) = int32(0)
			v421 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[812])) = v421
			v425 = *(*int32)(unsafe.Add(mBase, _consts[612]))
			if v425 != 0 {
				v439 = v425
				v441 = *(*int32)(unsafe.Add(mBase, _consts[611]))
				if v441 != 0 {
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[611])) = v439
				}
				v445 = *(*int32)(unsafe.Add(mBase, _consts[610]))
				if v445 != 0 {
				} else {
					v446 = int32(0)
					v448 = *(*int32)(unsafe.Add(mBase, _consts[811]))
					*(*int32)(unsafe.Add(mBase, _consts[610])) = v448
				}
				m.G0 = v5 + int32(16)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(0)
				F_scriptingEngineManagerForEachEngine(m, int32(1022), v5+int32(12))
				mBase = m.M
				v432 = m.ExcPending
				if v432 != 0 {
					return
				} else {
					v433 = int32(0)
					v435 = *(*int32)(unsafe.Add(mBase, _consts[810]))
					v436 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					v437 = v435 - v436
					*(*int32)(unsafe.Add(mBase, _consts[612])) = v437
					v439 = v437
					v441 = *(*int32)(unsafe.Add(mBase, _consts[611]))
					if v441 != 0 {
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[611])) = v439
					}
					v445 = *(*int32)(unsafe.Add(mBase, _consts[610]))
					if v445 != 0 {
					} else {
						v446 = int32(0)
						v448 = *(*int32)(unsafe.Add(mBase, _consts[811]))
						*(*int32)(unsafe.Add(mBase, _consts[610])) = v448
					}
					m.G0 = v5 + int32(16)
					return
				}
			}
		}
	}
}
