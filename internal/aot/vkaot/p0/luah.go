package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaH_free(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v5 = m.G3
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v6 == v5+int32(_a2240) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
		v21 = F_luaM_realloc_(m, l0, v16, v17<<(uint(int32(4))%32), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v25 = F_luaM_realloc_(m, l0, l1, int32(40), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
		v14 = F_luaM_realloc_(m, l0, v6, int32(32)<<(uint(v11)%32), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
			v21 = F_luaM_realloc_(m, l0, v16, v17<<(uint(int32(4))%32), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v25 = F_luaM_realloc_(m, l0, l1, int32(40), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_luaH_getn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var __phi60 int32
	_ = __phi60
	var v62 int32
	_ = v62
	var __phi62 int32
	_ = __phi62
	var v71 int32
	_ = v71
	var v77 float64
	_ = v77
	var v81 float64
	_ = v81
	var v82 int64
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 float64
	_ = v99
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 float64
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v155 float64
	_ = v155
	var v156 int64
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 float64
	_ = v173
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v188 float64
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v227 int32
	_ = v227
	var v233 float64
	_ = v233
	var v237 float64
	_ = v237
	var v238 int64
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 float64
	_ = v255
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v270 float64
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v11 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v293
L2:
	;
	v49 = m.G3
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v50 != v49+int32(_a2240) {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+v11<<(uint(int32(4))%32)+int32(-8))))
	if v20 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v21 = int32(0)
	if v11 == int32(1) {
		v293 = v21
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v27 = v11
	v29 = v21
	goto L6
L6:
	;
	v38 = int32(base.Ui32(v29+v27) >> (uint(int32(1)) % 32))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-8)+v38<<(uint(int32(4))%32))))
	if v42 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v43 = v27
	goto L10
L9:
	;
	v43 = v38
	goto L10
L10:
	;
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = v38
	goto L13
L12:
	;
	v44 = v29
	goto L13
L13:
	;
	if base.Ui32(int32(1)) < base.Ui32(v43-v44) {
		v27 = v43
		v29 = v44
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v293 = v44
	goto L1
L15:
	;
	__phi60 = v11
	__phi62 = v11 + int32(1)
	v60 = __phi60
	v62 = __phi62
	goto L17
L16:
	;
	return v11
L17:
	;
	if v62 < int32(1) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v215 = v208
	goto L58
L19:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	if v128 != 0 {
		goto L33
	} else {
		goto L34
	}
L20:
	;
	v103 = v98
	goto L27
L21:
	;
	v82 = base.I64_reinterpret_f64(v81)
	v87 = int32(-1)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v94 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v82)>>(uint(int64(32))%64))+v82), v87<<(uint(v88)%32)^v87|int32(1))
	v98 = v50 + v94<<(uint(int32(5))%32)
	v99 = v81
	goto L20
L22:
	;
	v77 = base.F64_convert_i32_s(v62)
	if v62 == int32(0) {
		v98 = v50
		v99 = v77
		goto L20
	} else {
		goto L26
	}
L23:
	;
	if v62 <= v11 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v123 = v71 + v62<<(uint(int32(4))%32) + int32(-16)
	goto L19
L25:
	;
	v81 = base.F64_convert_i32_u(v62)
	goto L21
L26:
	;
	v81 = v77
	goto L21
L27:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v103)+24))
	if v111 != int32(3) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v123 = v117
	goto L19
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)+28))
	v117 = m.G398
	if v116 != 0 {
		v103 = v116
		goto L27
	} else {
		goto L32
	}
L30:
	;
	v114 = *(*float64)(unsafe.Add(mBase, uint32(v103)+16))
	if base.F64_ne(v114, v99) != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v123 = v103
	goto L19
L32:
	;
	goto L28
L33:
	;
	v208 = int32(1)
	v210 = v62 << (uint(v208) % 32)
	if base.Ui32(v210) < base.Ui32(int32(2147483646)) {
		__phi60 = v62
		__phi62 = v210
		v60 = __phi60
		v62 = __phi62
		goto L17
	} else {
		goto L57
	}
L34:
	;
	if base.Ui32(v62-v60) <= base.Ui32(int32(1)) {
		v293 = v60
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v135 = v60
	v136 = v62
	goto L36
L36:
	;
	v142 = v136 + v135
	v144 = int32(base.Ui32(v142) >> (uint(int32(1)) % 32))
	if base.Ui32(v142) < base.Ui32(int32(2)) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v197)+8))
	if v202 != 0 {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v177 = v172
	goto L44
L40:
	;
	v155 = base.F64_convert_i32_u(v144)
	v156 = base.I64_reinterpret_f64(v155)
	v161 = int32(-1)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v168 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v156)>>(uint(int64(32))%64))+v156), v161<<(uint(v162)%32)^v161|int32(1))
	v172 = v50 + v168<<(uint(int32(5))%32)
	v173 = v155
	goto L39
L41:
	;
	v172 = v50
	v173 = base.F64_convert_i32_u(v144)
	goto L39
L42:
	;
	if v11 < v144 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v197 = v148 + v144<<(uint(int32(4))%32) + int32(-16)
	goto L38
L44:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v177)+24))
	if v185 != int32(3) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v197 = v191
	goto L38
L46:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v177)+28))
	v191 = m.G398
	if v190 != 0 {
		v177 = v190
		goto L44
	} else {
		goto L49
	}
L47:
	;
	v188 = *(*float64)(unsafe.Add(mBase, uint32(v177)+16))
	if base.F64_ne(v188, v173) != 0 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v197 = v177
	goto L38
L49:
	;
	goto L45
L50:
	;
	v203 = v136
	goto L52
L51:
	;
	v203 = v144
	goto L52
L52:
	;
	if v202 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v204 = v144
	goto L55
L54:
	;
	v204 = v135
	goto L55
L55:
	;
	if base.Ui32(int32(1)) < base.Ui32(v203-v204) {
		v135 = v204
		v136 = v203
		goto L36
	} else {
		goto L56
	}
L56:
	;
	v293 = v204
	goto L1
L57:
	;
	goto L18
L58:
	;
	if v215 < int32(1) {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	return v215 + int32(-1)
L60:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v279)+8))
	if v286 != 0 {
		v215 = v215 + int32(1)
		goto L58
	} else {
		goto L74
	}
L61:
	;
	v259 = v254
	goto L68
L62:
	;
	v238 = base.I64_reinterpret_f64(v237)
	v243 = int32(-1)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v250 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v238)>>(uint(int64(32))%64))+v238), v243<<(uint(v244)%32)^v243|int32(1))
	v254 = v50 + v250<<(uint(int32(5))%32)
	v255 = v237
	goto L61
L63:
	;
	v233 = base.F64_convert_i32_s(v215)
	if v215 == int32(0) {
		v254 = v50
		v255 = v233
		goto L61
	} else {
		goto L67
	}
L64:
	;
	if v215 <= v11 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v279 = v227 + v215<<(uint(int32(4))%32) + int32(-16)
	goto L60
L66:
	;
	v237 = base.F64_convert_i32_u(v215)
	goto L62
L67:
	;
	v237 = v233
	goto L62
L68:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v259)+24))
	if v267 != int32(3) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v279 = v273
	goto L60
L70:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v259)+28))
	v273 = m.G398
	if v272 != 0 {
		v259 = v272
		goto L68
	} else {
		goto L73
	}
L71:
	;
	v270 = *(*float64)(unsafe.Add(mBase, uint32(v259)+16))
	if base.F64_ne(v270, v255) != 0 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v279 = v259
	goto L60
L73:
	;
	goto L69
L74:
	;
	goto L59
}
func F_luaH_new(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
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
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	v4 = int32(0)
	v13 = F_luaM_realloc_(m, l0, v4, v4, int32(40))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(5)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v13
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)))
		*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v17)
		v25 = v22 & int32(3)
		*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)) = uint8(v25)
		v27 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v27
		v29 = int32(255)
		*(*uint8)(unsafe.Add(mBase, uint32(v13)+6)) = uint8(v29)
		*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)) = uint8(v27)
		v35 = m.G3
		*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v35 + int32(_a2240)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v27
		if base.Ui32(int32(268435455)) < base.Ui32(l1+int32(1)) {
			v51 = F_luaM_toobig(m, l0)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				v53 = v51
				*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v53
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
				if l1 <= v55 {
				} else {
					v59 = (l1 - v55) & int32(7)
					if v59 == int32(0) {
						v88 = v55
					} else {
						v69 = v55
						v71 = int32(0)
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v53+v69<<(uint(int32(4))%32))+8)) = int32(0)
							v77 = int32(1)
							v78 = v69 + v77
							v80 = v71 + v77
							if v80 != v59 {
								v69 = v78
								v71 = v80
								continue
							} else {
								break
							}
							break
						}
						v88 = v78
					}
					if base.Ui32(int32(-8)) < base.Ui32(v55-l1) {
					} else {
						v100 = v88
						for {
							v105 = v53 + v100<<(uint(int32(4))%32)
							v106 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(24)))) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(40)))) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(56)))) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(72)))) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(88)))) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(104)))) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(120)))) = v106
							v137 = v100 + int32(8)
							if v137 != l1 {
								v100 = v137
								continue
							} else {
								break
							}
							break
						}
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = l1
				F_setnodevector(m, l0, v13, l2)
				mBase = m.M
				v150 = m.ExcPending
				if v150 != 0 {
					return int32(0)
				} else {
					return v13
				}
			}
		} else {
			v45 = int32(0)
			v49 = F_luaM_realloc_(m, l0, v45, v45, l1<<(uint(int32(4))%32))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				v53 = v49
				*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v53
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
				if l1 <= v55 {
				} else {
					v59 = (l1 - v55) & int32(7)
					if v59 == int32(0) {
						v88 = v55
					} else {
						v69 = v55
						v71 = int32(0)
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v53+v69<<(uint(int32(4))%32))+8)) = int32(0)
							v77 = int32(1)
							v78 = v69 + v77
							v80 = v71 + v77
							if v80 != v59 {
								v69 = v78
								v71 = v80
								continue
							} else {
								break
							}
							break
						}
						v88 = v78
					}
					if base.Ui32(int32(-8)) < base.Ui32(v55-l1) {
					} else {
						v100 = v88
						for {
							v105 = v53 + v100<<(uint(int32(4))%32)
							v106 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(24)))) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(40)))) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(56)))) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(72)))) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(88)))) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(104)))) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v105+int32(120)))) = v106
							v137 = v100 + int32(8)
							if v137 != l1 {
								v100 = v137
								continue
							} else {
								break
							}
							break
						}
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = l1
				F_setnodevector(m, l0, v13, l2)
				mBase = m.M
				v150 = m.ExcPending
				if v150 != 0 {
					return int32(0)
				} else {
					return v13
				}
			}
		}
	}
}
func F_luaH_next(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int64
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int64
	_ = v228
	var v231 int32
	_ = v231
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	switch v10 {
	case 0:
		v160 = int32(-1)
		goto L1
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L7
	case 4:
		goto L6
	default:
		goto L3
	}
L1:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v167 = v160 + int32(1)
	if v167 < v165 {
		goto L36
	} else {
		goto L37
	}
L2:
	;
	v103 = v97
	goto L17
L3:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v86 = int32(-1)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v93 = base.I32_rem_u_s(v85, v86<<(uint(v87)%32)^v86|int32(1))
	v97 = v84 + v93<<(uint(int32(5))%32)
	goto L2
L4:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v73 = int32(-1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v80 = base.I32_rem_u_s(v72, v73<<(uint(v74)%32)^v73|int32(1))
	v97 = v71 + v80<<(uint(int32(5))%32)
	goto L2
L5:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v62 = int32(-1)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v97 = v60 + v61&(v62<<(uint(v63)%32)^v62)<<(uint(int32(5))%32)
	goto L2
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v51 = int32(-1)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v97 = v48 + v50&(v51<<(uint(v52)%32)^v51)<<(uint(int32(5))%32)
	goto L2
L7:
	;
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	if base.F64_lt(base.F64_abs(v11), float64(2.147483648e+09)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v19 < int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v19 = int32(-2147483648)
	goto L8
L10:
	;
	v17 = base.I32_trunc_f64_s(v11)
	v19 = v17
	goto L8
L11:
	;
	if base.F64_ne(v11, float64(0)) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if base.F64_ne(v11, base.F64_convert_i32_s(v19)) != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v24 < v19 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v160 = v19 + int32(-1)
	goto L1
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v32 = base.I64_reinterpret_f64(v11)
	v37 = int32(-1)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v44 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v32)>>(uint(int64(32))%64))+v32), v37<<(uint(v38)%32)^v37|int32(1))
	v97 = v31 + v44<<(uint(int32(5))%32)
	goto L2
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v97 = v30
	goto L2
L17:
	;
	v109 = v103 + int32(16)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v112 == v113 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v148 = int32(0)
	v149 = m.G3
	F_luaG_runerror(m, l0, v149+int32(_a2271), v148)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L34
	} else {
		goto L35
	}
L19:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v103)+28))
	if v147 != 0 {
		v103 = v147
		goto L17
	} else {
		goto L33
	}
L20:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v160 = (v103-v141)>>(uint(int32(5))%32) + v145
	goto L1
L21:
	;
	if v131 != 0 {
		goto L20
	} else {
		goto L29
	}
L22:
	;
	switch v112 {
	case 0:
		v129 = int32(1)
		goto L24
	case 1:
		goto L27
	case 2:
		goto L26
	case 3:
		goto L28
	default:
		goto L25
	}
L23:
	;
	v131 = int32(0)
	goto L21
L24:
	;
	v131 = v129
	goto L21
L25:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v129 = base.B2i32(v126 == v127)
	goto L24
L26:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v131 = base.B2i32(v123 == v124)
	goto L21
L27:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v131 = base.B2i32(v120 == v121)
	goto L21
L28:
	;
	v117 = *(*float64)(unsafe.Add(mBase, uint32(v109)))
	v118 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v131 = base.F64_eq(v117, v118)
	goto L21
L29:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v103)+24))
	if v132 != int32(11) {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v135 < int32(4) {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v138 != v139 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	goto L20
L33:
	;
	goto L18
L34:
	;
	return int32(0)
L35:
	;
	v160 = v148
	goto L1
L36:
	;
	v169 = v165
	goto L38
L37:
	;
	v169 = v167
	goto L38
L38:
	;
	v173 = v160
	goto L41
L39:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v246)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v251
	return int32(1)
L40:
	;
	v200 = v169 - v165
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v203 = int32(1) << (uint(v202) % 32)
	if v203 <= v200 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v179 = v173 + int32(1)
	if v165 <= v179 {
		goto L40
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(l2))) = base.F64_convert_i32_s(v173 + int32(2))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v197 = v194 + v179<<(uint(int32(4))%32)
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v197)))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v198
	v246 = v197
	goto L39
L43:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181+v179<<(uint(int32(4))%32))+8))
	if v185 == int32(0) {
		v173 = v179
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	return int32(0)
L46:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v209 = v200
	goto L47
L47:
	;
	v216 = v205 + v209<<(uint(int32(5))%32)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	if v217 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L45
L49:
	;
	v231 = v209 + int32(1)
	if v231 != v203 {
		v209 = v231
		goto L47
	} else {
		goto L51
	}
L50:
	;
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v216)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v216)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v222
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v227 = v224 + v209<<(uint(int32(5))%32)
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v227)))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v228
	v246 = v227
	goto L39
L51:
	;
	goto L48
}
func F_luaH_set(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v4 = int32(0)
	v7 = F_luaH_get(m, l1, l2)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)) = uint8(v4)
	v10 = m.G398
	if v7 != v10 {
		v34 = v7
		return v34
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		v13 = m.G3
		switch v12 {
		case 0:
			v22 = v13 + int32(_a2272)
			F_luaG_runerror(m, l0, v22, int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v32 = F_newkey(m, l0, l1, l2)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = v32
					return v34
				}
			}
		default:
			v32 = F_newkey(m, l0, l1, l2)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = v32
				return v34
			}
		case 3:
			v16 = m.G3
			v17 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
			if base.F64_eq(v17, v17) != 0 {
				v32 = F_newkey(m, l0, l1, l2)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = v32
					return v34
				}
			} else {
				v22 = v16 + int32(_a2273)
				F_luaG_runerror(m, l0, v22, int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v32 = F_newkey(m, l0, l1, l2)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = v32
						return v34
					}
				}
			}
		}
	}
}
