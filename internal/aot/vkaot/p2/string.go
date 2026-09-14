package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_copy_string_from_lua_stack(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v5 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = F_lua_objlen(m, l0, int32(-1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v14 = m.G22
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v16 = m.T0[v15].(func(*base.Module, int32) int32)(m, v10+int32(1))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = F___stpncpy(m, v16, v5, v10)
				mBase = m.M
				v20 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v10+v16))) = uint8(v20)
				return v16
			}
		}
	}
}
func F_createStringObjectFromLongDouble(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(5136)
	m.G0 = v8
	v13 = F_ld2string(m, v8, int32(5120), l0, l1, base.B2i32(l2 != int32(0)))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if base.Ui32(int32(255)) < base.Ui32(v13) {
			v41 = F_sdsnewlen(m, v8, v13)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+uint32(_consts[603]))) = int32(0)
				v48 = F_zmalloc_usable(m, int32(12), v8+int32(5132))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v41
					*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(34359738368)
					v53 = v48
					m.G0 = v8 + int32(5136)
					return v53
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			if base.Ui32(int32(128)) < base.Ui32(v13+v30+int32(9)) {
				v41 = F_sdsnewlen(m, v8, v13)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+uint32(_consts[603]))) = int32(0)
					v48 = F_zmalloc_usable(m, int32(12), v8+int32(5132))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v41
						*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(34359738368)
						v53 = v48
						m.G0 = v8 + int32(5136)
						return v53
					}
				}
			} else {
				v39 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v8, v13, int32(0), int64(-1))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v53 = v39
					m.G0 = v8 + int32(5136)
					return v53
				}
			}
		}
	}
}
func F_createStringObjectFromSds(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v12 & int32(7) {
	case 0:
		v32 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
		v44 = *(*int32)(unsafe.Add(mBase, _consts[601]))
		if base.Ui32(int32(128)) < base.Ui32(v32+v44+int32(9)) {
			v57 = v32
			v58 = F_sdsnewlen(m, l0, v57)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
				v62 = int32(12)
				v65 = F_zmalloc_usable(m, v62, v7+v62)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v58
					*(*int64)(unsafe.Add(mBase, uint32(v65))) = int64(34359738368)
					v71 = v65
					m.G0 = v7 + int32(16)
					return v71
				}
			}
		} else {
			v53 = F_createEmbeddedStringObjectWithKeyAndExpire(m, l0, v32, int32(0), int64(-1))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				v71 = v53
				m.G0 = v7 + int32(16)
				return v71
			}
		}
	case 1:
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v32 = v19
		v44 = *(*int32)(unsafe.Add(mBase, _consts[601]))
		if base.Ui32(int32(128)) < base.Ui32(v32+v44+int32(9)) {
			v57 = v32
			v58 = F_sdsnewlen(m, l0, v57)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
				v62 = int32(12)
				v65 = F_zmalloc_usable(m, v62, v7+v62)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v58
					*(*int64)(unsafe.Add(mBase, uint32(v65))) = int64(34359738368)
					v71 = v65
					m.G0 = v7 + int32(16)
					return v71
				}
			}
		} else {
			v53 = F_createEmbeddedStringObjectWithKeyAndExpire(m, l0, v32, int32(0), int64(-1))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				v71 = v53
				m.G0 = v7 + int32(16)
				return v71
			}
		}
	case 2:
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v29 = v22
		if base.Ui32(int32(255)) < base.Ui32(v29) {
			v57 = v29
			v58 = F_sdsnewlen(m, l0, v57)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
				v62 = int32(12)
				v65 = F_zmalloc_usable(m, v62, v7+v62)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v58
					*(*int64)(unsafe.Add(mBase, uint32(v65))) = int64(34359738368)
					v71 = v65
					m.G0 = v7 + int32(16)
					return v71
				}
			}
		} else {
			v32 = v29
			v44 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			if base.Ui32(int32(128)) < base.Ui32(v32+v44+int32(9)) {
				v57 = v32
				v58 = F_sdsnewlen(m, l0, v57)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
					v62 = int32(12)
					v65 = F_zmalloc_usable(m, v62, v7+v62)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v58
						*(*int64)(unsafe.Add(mBase, uint32(v65))) = int64(34359738368)
						v71 = v65
						m.G0 = v7 + int32(16)
						return v71
					}
				}
			} else {
				v53 = F_createEmbeddedStringObjectWithKeyAndExpire(m, l0, v32, int32(0), int64(-1))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v71 = v53
					m.G0 = v7 + int32(16)
					return v71
				}
			}
		}
	case 3:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v29 = v25
		if base.Ui32(int32(255)) < base.Ui32(v29) {
			v57 = v29
			v58 = F_sdsnewlen(m, l0, v57)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
				v62 = int32(12)
				v65 = F_zmalloc_usable(m, v62, v7+v62)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v58
					*(*int64)(unsafe.Add(mBase, uint32(v65))) = int64(34359738368)
					v71 = v65
					m.G0 = v7 + int32(16)
					return v71
				}
			}
		} else {
			v32 = v29
			v44 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			if base.Ui32(int32(128)) < base.Ui32(v32+v44+int32(9)) {
				v57 = v32
				v58 = F_sdsnewlen(m, l0, v57)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
					v62 = int32(12)
					v65 = F_zmalloc_usable(m, v62, v7+v62)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v58
						*(*int64)(unsafe.Add(mBase, uint32(v65))) = int64(34359738368)
						v71 = v65
						m.G0 = v7 + int32(16)
						return v71
					}
				}
			} else {
				v53 = F_createEmbeddedStringObjectWithKeyAndExpire(m, l0, v32, int32(0), int64(-1))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v71 = v53
					m.G0 = v7 + int32(16)
					return v71
				}
			}
		}
	case 4:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v29 = v28
		if base.Ui32(int32(255)) < base.Ui32(v29) {
			v57 = v29
			v58 = F_sdsnewlen(m, l0, v57)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
				v62 = int32(12)
				v65 = F_zmalloc_usable(m, v62, v7+v62)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v58
					*(*int64)(unsafe.Add(mBase, uint32(v65))) = int64(34359738368)
					v71 = v65
					m.G0 = v7 + int32(16)
					return v71
				}
			}
		} else {
			v32 = v29
			v44 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			if base.Ui32(int32(128)) < base.Ui32(v32+v44+int32(9)) {
				v57 = v32
				v58 = F_sdsnewlen(m, l0, v57)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
					v62 = int32(12)
					v65 = F_zmalloc_usable(m, v62, v7+v62)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v58
						*(*int64)(unsafe.Add(mBase, uint32(v65))) = int64(34359738368)
						v71 = v65
						m.G0 = v7 + int32(16)
						return v71
					}
				}
			} else {
				v53 = F_createEmbeddedStringObjectWithKeyAndExpire(m, l0, v32, int32(0), int64(-1))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v71 = v53
					m.G0 = v7 + int32(16)
					return v71
				}
			}
		}
	default:
		v32 = int32(0)
		v44 = *(*int32)(unsafe.Add(mBase, _consts[601]))
		if base.Ui32(int32(128)) < base.Ui32(v32+v44+int32(9)) {
			v57 = v32
			v58 = F_sdsnewlen(m, l0, v57)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
				v62 = int32(12)
				v65 = F_zmalloc_usable(m, v62, v7+v62)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v58
					*(*int64)(unsafe.Add(mBase, uint32(v65))) = int64(34359738368)
					v71 = v65
					m.G0 = v7 + int32(16)
					return v71
				}
			}
		} else {
			v53 = F_createEmbeddedStringObjectWithKeyAndExpire(m, l0, v32, int32(0), int64(-1))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				v71 = v53
				m.G0 = v7 + int32(16)
				return v71
			}
		}
	}
}
func F_equalStringObjects(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
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
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = int32(240)
	v10 = v8 & v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11&v9 != int32(16) {
		if v10 == int32(16) {
			v326 = F_compareStringObjectsWithFlags(m, l0, l1, int32(1))
			mBase = m.M
			v329 = m.ExcPending
			if v329 != 0 {
				return int32(0)
			} else {
				v332 = base.B2i32(v326 == int32(0))
				return v332
			}
		} else {
			v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v148&int32(4) == int32(0) {
				v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v208 = v207
			} else {
				if v148&int32(1) != 0 {
					v157 = int32(16)
				} else {
					v157 = int32(8)
				}
				v158 = l0 + v157
				if v148&int32(2) == int32(0) {
					v190 = v158
				} else {
					v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
					v164 = v158 + v163
					v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
					switch v168 & int32(7) {
					case 0:
						v185 = int32(base.Ui32(v168) >> (uint(int32(3)) % 32))
					case 1:
						v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+int32(-2)))))
						v185 = v175
					case 2:
						v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164+int32(-4)))))
						v185 = v178
					case 3:
						v181 = *(*int32)(unsafe.Add(mBase, uint32(v164+int32(-8))))
						v185 = v181
					case 4:
						v184 = *(*int32)(unsafe.Add(mBase, uint32(v164+int32(-16))))
						v185 = v184
					default:
						v185 = int32(0)
					}
					v190 = v164 + int32(1) + v185 + int32(1)
				}
				v204 = *(*int32)(unsafe.Add(mBase, _consts[601]))
				v208 = v190 + v204
			}
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208+int32(-1)))))
			switch v215 & int32(7) {
			case 0:
				v232 = int32(base.Ui32(v215) >> (uint(int32(3)) % 32))
			case 1:
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208+int32(-3)))))
				v232 = v222
			case 2:
				v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208+int32(-5)))))
				v232 = v225
			case 3:
				v228 = *(*int32)(unsafe.Add(mBase, uint32(v208+int32(-9))))
				v232 = v228
			case 4:
				v231 = *(*int32)(unsafe.Add(mBase, uint32(v208+int32(-17))))
				v232 = v231
			default:
				v232 = int32(0)
			}
			v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v233&int32(4) == int32(0) {
				v292 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v295 = v292
			} else {
				if v233&int32(1) != 0 {
					v242 = int32(16)
				} else {
					v242 = int32(8)
				}
				v243 = l1 + v242
				if v233&int32(2) == int32(0) {
					v275 = v243
				} else {
					v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
					v249 = v243 + v248
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
					switch v253 & int32(7) {
					case 0:
						v270 = int32(base.Ui32(v253) >> (uint(int32(3)) % 32))
					case 1:
						v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249+int32(-2)))))
						v270 = v260
					case 2:
						v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v249+int32(-4)))))
						v270 = v263
					case 3:
						v266 = *(*int32)(unsafe.Add(mBase, uint32(v249+int32(-8))))
						v270 = v266
					case 4:
						v269 = *(*int32)(unsafe.Add(mBase, uint32(v249+int32(-16))))
						v270 = v269
					default:
						v270 = int32(0)
					}
					v275 = v249 + int32(1) + v270 + int32(1)
				}
				v289 = *(*int32)(unsafe.Add(mBase, _consts[601]))
				v295 = v275 + v289
			}
			v297 = int32(0)
			v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+int32(-1)))))
			switch v301 & int32(7) {
			case 0:
				v318 = int32(base.Ui32(v301) >> (uint(int32(3)) % 32))
			case 1:
				v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+int32(-3)))))
				v318 = v308
			case 2:
				v311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v295+int32(-5)))))
				v318 = v311
			case 3:
				v314 = *(*int32)(unsafe.Add(mBase, uint32(v295+int32(-9))))
				v318 = v314
			case 4:
				v317 = *(*int32)(unsafe.Add(mBase, uint32(v295+int32(-17))))
				v318 = v317
			default:
				v318 = v297
			}
			if v232 != v318 {
				v332 = v297
				return v332
			} else {
				v326 = F_compareStringObjectsWithFlags(m, l0, l1, int32(1))
				mBase = m.M
				v329 = m.ExcPending
				if v329 != 0 {
					return int32(0)
				} else {
					v332 = base.B2i32(v326 == int32(0))
					return v332
				}
			}
		}
	} else {
		if v10 != int32(16) {
			v326 = F_compareStringObjectsWithFlags(m, l0, l1, int32(1))
			mBase = m.M
			v329 = m.ExcPending
			if v329 != 0 {
				return int32(0)
			} else {
				v332 = base.B2i32(v326 == int32(0))
				return v332
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v18&int32(4) == int32(0) {
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v78 = v77
			} else {
				if v18&int32(1) != 0 {
					v27 = int32(16)
				} else {
					v27 = int32(8)
				}
				v28 = l0 + v27
				if v18&int32(2) == int32(0) {
					v59 = v28
				} else {
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
					v34 = v28 + v33
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
					switch v38 & int32(7) {
					case 0:
						v55 = int32(base.Ui32(v38) >> (uint(int32(3)) % 32))
					case 1:
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+int32(-2)))))
						v55 = v45
					case 2:
						v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(-4)))))
						v55 = v48
					case 3:
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(-8))))
						v55 = v51
					case 4:
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(-16))))
						v55 = v54
					default:
						v55 = int32(0)
					}
					v59 = v34 + int32(1) + v55 + int32(1)
				}
				v74 = *(*int32)(unsafe.Add(mBase, _consts[601]))
				v78 = v59 + v74
			}
			v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v82&int32(4) == int32(0) {
				v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				return base.B2i32(v78 == v143)
			} else {
				if v82&int32(1) != 0 {
					v91 = int32(16)
				} else {
					v91 = int32(8)
				}
				v92 = l1 + v91
				if v82&int32(2) == int32(0) {
					v123 = v92
				} else {
					v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
					v98 = v92 + v97
					v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
					switch v102 & int32(7) {
					case 0:
						v119 = int32(base.Ui32(v102) >> (uint(int32(3)) % 32))
					case 1:
						v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+int32(-2)))))
						v119 = v109
					case 2:
						v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98+int32(-4)))))
						v119 = v112
					case 3:
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v98+int32(-8))))
						v119 = v115
					case 4:
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v98+int32(-16))))
						v119 = v118
					default:
						v119 = int32(0)
					}
					v123 = v98 + int32(1) + v119 + int32(1)
				}
				v138 = *(*int32)(unsafe.Add(mBase, _consts[601]))
				return base.B2i32(v78 == v123+v138)
			}
		}
	}
}
func F_freeStringObject(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5&int32(240) != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8&int32(4) == int32(0) {
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v68 = v67
		} else {
			if v8&int32(1) != 0 {
				v17 = int32(16)
			} else {
				v17 = int32(8)
			}
			v18 = l0 + v17
			if v8&int32(2) == int32(0) {
				v49 = v18
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				v24 = v18 + v23
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				switch v28 & int32(7) {
				case 0:
					v45 = int32(base.Ui32(v28) >> (uint(int32(3)) % 32))
				case 1:
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-2)))))
					v45 = v35
				case 2:
					v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(-4)))))
					v45 = v38
				case 3:
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-8))))
					v45 = v41
				case 4:
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-16))))
					v45 = v44
				default:
					v45 = int32(0)
				}
				v49 = v24 + int32(1) + v45 + int32(1)
			}
			v64 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			v68 = v49 + v64
		}
		F_sdsfree(m, v68)
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return
		} else {
			return
		}
	}
}
func F_readString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v140 int32
	_ = v140
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(1120)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3
	v18 = F_readLong(m, l0, int32(36), v10+int32(92))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v18 == int32(0) {
			v140 = v3
			m.G0 = v10 + int32(1120)
			return v140
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
			if base.Ui32(v24) < base.Ui32(int32(2147483646)) {
				v49 = v24 + int32(2)
				v50 = F_valkey_malloc(m, v49)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v50
					v53 = int32(0)
					v55 = F___ftello(m, l0)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, _consts[958])) = v55
						v59 = F_fread(m, v50, int32(1), v49, l0)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							if v59 == v49 {
								v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v94 = int32(*(*int8)(unsafe.Add(mBase, uint32(v90+v49+int32(-1)))))
								v95 = v90 + v24
								v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v95))))
								if v96 != int32(13) {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v96
									*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v94
									v109 = F_snprintf(m, v10+int32(96), int32(1024), int32(_a1770), v10+int32(48))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v111 = int32(0)
										v113 = *(*int64)(unsafe.Add(mBase, _consts[958]))
										*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v113
										*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v10 + int32(96)
										v123 = F_snprintf(m, int32(_a1771), int32(1044), int32(_a1772), v10+int32(32))
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return int32(0)
										} else {
											v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											F_valkey_free(m, v125)
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
												v140 = v111
												m.G0 = v10 + int32(1120)
												return v140
											}
										}
									}
								} else {
									if v94 == int32(10) {
										v130 = int32(0)
										v132 = *(*int64)(unsafe.Add(mBase, _consts[959]))
										*(*int64)(unsafe.Add(mBase, _consts[959])) = v132 + int64(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v130)
										v140 = int32(1)
										m.G0 = v10 + int32(1120)
										return v140
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v96
										*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v94
										v109 = F_snprintf(m, v10+int32(96), int32(1024), int32(_a1770), v10+int32(48))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											v111 = int32(0)
											v113 = *(*int64)(unsafe.Add(mBase, _consts[958]))
											*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v113
											*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v10 + int32(96)
											v123 = F_snprintf(m, int32(_a1771), int32(1044), int32(_a1772), v10+int32(32))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return int32(0)
											} else {
												v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												F_valkey_free(m, v125)
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
													v140 = v111
													m.G0 = v10 + int32(1120)
													return v140
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v49
								*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v59
								v70 = F_snprintf(m, v10+int32(96), int32(1024), int32(_a1774), v10+int32(80))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									v73 = *(*int64)(unsafe.Add(mBase, _consts[958]))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v73
									*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v10 + int32(96)
									v83 = F_snprintf(m, int32(_a1771), int32(1044), int32(_a1772), v10+int32(64))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										F_valkey_free(m, v85)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
											v140 = v53
											m.G0 = v10 + int32(1120)
											return v140
										}
									}
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v24
				v34 = F_snprintf(m, v10+int32(96), int32(1024), int32(_a1775), v10+int32(16))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v36 = int32(0)
					v38 = *(*int64)(unsafe.Add(mBase, _consts[958]))
					*(*int64)(unsafe.Add(mBase, uint32(v10))) = v38
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v10 + int32(96)
					v46 = F_snprintf(m, int32(_a1771), int32(1044), int32(_a1772), v10)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v140 = v36
						m.G0 = v10 + int32(1120)
						return v140
					}
				}
			}
		}
	}
}
func F_stringConfigRewrite(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_rewriteConfigStringOption(m, l2, l1, v5, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_string_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	v4 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v10 = l2 + int32(256)
	v14 = base.B2i32(v10 != v4)
	if v7&int32(3) == v4 {
		v40 = v7
		v42 = v10
		v43 = v14
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if v113 != 0 {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	v113 = int32(0)
	goto L1
L3:
	;
	v91 = v84
	v93 = v86
	goto L21
L4:
	;
	if v43 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L5:
	;
	if v10 == int32(0) {
		v40 = v7
		v42 = v10
		v43 = v14
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v23 = v7
	v25 = v10
	goto L7
L7:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v28 == int32(0) {
		v84 = v23
		v86 = v25
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v40 = v35
	v42 = v31
	v43 = v33
	goto L4
L9:
	;
	v31 = v25 + int32(-1)
	v32 = int32(0)
	v33 = base.B2i32(v31 != v32)
	v35 = v23 + int32(1)
	if v35&int32(3) == v32 {
		v40 = v35
		v42 = v31
		v43 = v33
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v31 != 0 {
		v23 = v35
		v25 = v31
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v47 == int32(0) {
		v77 = v40
		v79 = v42
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v79 == int32(0) {
		goto L2
	} else {
		goto L20
	}
L14:
	;
	if base.Ui32(v42) < base.Ui32(int32(4)) {
		v77 = v40
		v79 = v42
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v57 = v40
	v59 = v42
	goto L16
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v64 = v63 ^ int32(0)
	v67 = int32(-2139062144)
	if (int32(16843008)-v64|v64)&v67 != v67 {
		v84 = v57
		v86 = v59
		goto L3
	} else {
		goto L18
	}
L17:
	;
	v77 = v72
	v79 = v74
	goto L13
L18:
	;
	v72 = v57 + int32(4)
	v74 = v59 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v74) {
		v57 = v72
		v59 = v74
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v84 = v77
	v86 = v79
	goto L3
L21:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v96 != int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L2
L23:
	;
	v101 = v93 + int32(-1)
	if v101 != 0 {
		v91 = v91 + int32(1)
		v93 = v101
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v113 = v91
	goto L1
L25:
	;
	goto L22
L26:
	;
	v115 = v113 - v7
	goto L28
L27:
	;
	v115 = v10
	goto L28
L28:
	;
	if base.Ui32(v115) < base.Ui32(l2) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v117 = v115
	goto L31
L30:
	;
	v117 = l2
	goto L31
L31:
	;
	if v117 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v122 = v7 + v115
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 + v117
	return v117
L33:
	;
	goto L32
L34:
	;
	v120 = F__emscripten_memcpy_bulkmem(m, l1, v7, v117)
	mBase = m.M
	goto L33
}
