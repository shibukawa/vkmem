package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaB_assert(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_luaL_checkany(m, l0, int32(1))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v23 = v18 + int32(0)
		v24 = m.G398
		if base.Ui32(v23) < base.Ui32(v17) {
			v26 = v23
		} else {
			v26 = v24
		}
		v68 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
		switch v68 {
		case 0:
			v73 = v68
			v75 = v73
		case 1:
			v69 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v75 = base.B2i32(v69 != int32(0))
		default:
			v73 = int32(1)
			v75 = v73
		}
		if v75 != 0 {
			v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v93 = (v88 - v89) >> (uint(int32(4)) % 32)
			m.G0 = v6 + int32(16)
			return v93
		} else {
			v77 = m.G3
			v81 = F_luaL_optlstring(m, l0, int32(2), v77+int32(_a2071), int32(0))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v81
				v86 = F_luaL_error(m, l0, v77+int32(_a79), v6)
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					v93 = v86
					m.G0 = v6 + int32(16)
					return v93
				}
			}
		}
	}
}
func F_luaB_auxwrap(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v54 int32
	_ = v54
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
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int64
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	switch int32(-1) {
	case 0:
		v56 = l0 + int32(72)
	case 1:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
		v56 = l0 + int32(88)
	case 2:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v56 = v25 + int32(96)
	default:
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
		v43 = m.G398
		if base.Ui32(v42) < base.Ui32(int32(1)) {
			v54 = v43
		} else {
			v54 = v41 + int32(24)
		}
		v56 = v54
	}
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v59 != int32(8) {
		v63 = int32(0)
	} else {
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
		v63 = v62
	}
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v69 = F_auxresume(m, l0, v63, (v64-v65)>>(uint(int32(4))%32))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		return int32(0)
	} else {
		if int32(-1) < v69 {
			return v69
		} else {
			v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v92 = v89 + int32(-16)
			v126 = m.G398
			if v92 != v126 {
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
				v136 = base.B2i32(base.Ui32(v129+int32(-3)) < base.Ui32(int32(2)))
			} else {
				v136 = int32(0)
			}
			if v136 == int32(0) {
				v225 = F_lua_error(m, l0)
				mBase = m.M
				v226 = m.ExcPending
				if v226 != 0 {
					return int32(0)
				} else {
					return v69
				}
			} else {
				F_luaL_where(m, l0, int32(1))
				mBase = m.M
				v141 = m.ExcPending
				if v141 != 0 {
					return int32(0)
				} else {
					v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v162 = v159 + int32(-32)
					v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v198) <= base.Ui32(v162) {
						v215 = v198
					} else {
						v201 = v198
						for {
							v205 = v201 + int32(-16)
							v206 = *(*int64)(unsafe.Add(mBase, uint32(v205)))
							*(*int64)(unsafe.Add(mBase, uint32(v201))) = v206
							v210 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(-8))))
							*(*int32)(unsafe.Add(mBase, uint32(v201)+8)) = v210
							if base.Ui32(v162) < base.Ui32(v205) {
								v201 = v205
								continue
							} else {
								break
							}
							break
						}
						v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v215 = v213
					}
					v218 = *(*int64)(unsafe.Add(mBase, uint32(v215)))
					*(*int64)(unsafe.Add(mBase, uint32(v162))) = v218
					v220 = *(*int32)(unsafe.Add(mBase, uint32(v215)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = v220
					F_lua_concat(m, l0, int32(2))
					mBase = m.M
					v224 = m.ExcPending
					if v224 != 0 {
						return int32(0)
					} else {
						v225 = F_lua_error(m, l0)
						mBase = m.M
						v226 = m.ExcPending
						if v226 != 0 {
							return int32(0)
						} else {
							return v69
						}
					}
				}
			}
		}
	}
}
func F_luaB_cocreate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
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
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v212 int32
	_ = v212
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int64
	_ = v287
	var v289 int32
	_ = v289
	v3 = F_lua_newthread(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v15 = v10 + int32(0)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v15) < base.Ui32(v16) {
			v58 = m.G398
			if v15 != v58 {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v64 = v61
			} else {
				v64 = int32(-1)
			}
		} else {
			v64 = int32(-1)
		}
		if v64 != int32(6) {
			v134 = m.G3
			v137 = F_luaL_argerror(m, l0, int32(1), v134+int32(_a2077))
			mBase = m.M
			v138 = m.ExcPending
			if v138 != 0 {
				return int32(0)
			} else {
				v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v150 = v145 + int32(0)
				v151 = m.G398
				if base.Ui32(v150) < base.Ui32(v144) {
					v153 = v150
				} else {
					v153 = v151
				}
				v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v196 = *(*int64)(unsafe.Add(mBase, uint32(v153)))
				*(*int64)(unsafe.Add(mBase, uint32(v195))) = v196
				v198 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v195)+8)) = v198
				v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v200 + int32(16)
				if l0 == v3 {
				} else {
					v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v212 - int32(16)
					v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v280 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = v280 + int32(16)
					v286 = v279 + int32(0)
					v287 = *(*int64)(unsafe.Add(mBase, uint32(v286)))
					*(*int64)(unsafe.Add(mBase, uint32(v280))) = v287
					v289 = *(*int32)(unsafe.Add(mBase, uint32(v286)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v280)+8)) = v289
				}
				return int32(1)
			}
		} else {
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v77 = v72 + int32(0)
			v78 = m.G398
			if base.Ui32(v77) < base.Ui32(v71) {
				v80 = v77
			} else {
				v80 = v78
			}
			v123 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
			if v123 != int32(6) {
				v130 = int32(0)
			} else {
				v126 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+6)))
				v130 = base.B2i32(v127 != int32(0))
			}
			if v130 == int32(0) {
				v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v150 = v145 + int32(0)
				v151 = m.G398
				if base.Ui32(v150) < base.Ui32(v144) {
					v153 = v150
				} else {
					v153 = v151
				}
				v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v196 = *(*int64)(unsafe.Add(mBase, uint32(v153)))
				*(*int64)(unsafe.Add(mBase, uint32(v195))) = v196
				v198 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v195)+8)) = v198
				v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v200 + int32(16)
				if l0 == v3 {
				} else {
					v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v212 - int32(16)
					v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v280 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = v280 + int32(16)
					v286 = v279 + int32(0)
					v287 = *(*int64)(unsafe.Add(mBase, uint32(v286)))
					*(*int64)(unsafe.Add(mBase, uint32(v280))) = v287
					v289 = *(*int32)(unsafe.Add(mBase, uint32(v286)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v280)+8)) = v289
				}
				return int32(1)
			} else {
				v134 = m.G3
				v137 = F_luaL_argerror(m, l0, int32(1), v134+int32(_a2077))
				mBase = m.M
				v138 = m.ExcPending
				if v138 != 0 {
					return int32(0)
				} else {
					v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v150 = v145 + int32(0)
					v151 = m.G398
					if base.Ui32(v150) < base.Ui32(v144) {
						v153 = v150
					} else {
						v153 = v151
					}
					v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v196 = *(*int64)(unsafe.Add(mBase, uint32(v153)))
					*(*int64)(unsafe.Add(mBase, uint32(v195))) = v196
					v198 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v195)+8)) = v198
					v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v200 + int32(16)
					if l0 == v3 {
					} else {
						v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v212 - int32(16)
						v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v280 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = v280 + int32(16)
						v286 = v279 + int32(0)
						v287 = *(*int64)(unsafe.Add(mBase, uint32(v286)))
						*(*int64)(unsafe.Add(mBase, uint32(v280))) = v287
						v289 = *(*int32)(unsafe.Add(mBase, uint32(v286)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v280)+8)) = v289
					}
					return int32(1)
				}
			}
		}
	}
}
func F_luaB_corunning(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v3))) = l0
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v7 + int32(16)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+112))
	if base.B2i32(v12 == l0) == int32(0) {
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17 + int32(16)
	}
	return int32(1)
}
func F_luaB_cowrap(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v2 = F_luaB_cocreate(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v6 = m.G5
		F_lua_pushcclosure(m, l0, v6+int32(1244), int32(1))
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
func F_luaB_error(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	v5 = F_luaL_optinteger(m, l0, int32(2), int32(1))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v17 = v14 + int32(16)
		if base.Ui32(v17) <= base.Ui32(v13) {
		} else {
			v21 = v13
			for {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(0)
				v25 = v21 + int32(16)
				if base.Ui32(v25) < base.Ui32(v17) {
					v21 = v25
					continue
				} else {
					break
				}
				break
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v45 = v40 + int32(0)
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v45) < base.Ui32(v46) {
			v88 = m.G398
			if v45 != v88 {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
				v98 = base.B2i32(base.Ui32(v91+int32(-3)) < base.Ui32(int32(2)))
			} else {
				v98 = int32(0)
			}
		} else {
			v98 = int32(0)
		}
		if v98 == int32(0) {
			v173 = F_lua_error(m, l0)
			mBase = m.M
			v174 = m.ExcPending
			if v174 != 0 {
				return int32(0)
			} else {
				return v173
			}
		} else {
			if v5 < int32(1) {
				v173 = F_lua_error(m, l0)
				mBase = m.M
				v174 = m.ExcPending
				if v174 != 0 {
					return int32(0)
				} else {
					return v173
				}
			} else {
				F_luaL_where(m, l0, v5)
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return int32(0)
				} else {
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v116 = v111 + int32(0)
					v117 = m.G398
					if base.Ui32(v116) < base.Ui32(v110) {
						v119 = v116
					} else {
						v119 = v117
					}
					v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v162 = *(*int64)(unsafe.Add(mBase, uint32(v119)))
					*(*int64)(unsafe.Add(mBase, uint32(v161))) = v162
					v164 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v161)+8)) = v164
					v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v166 + int32(16)
					F_lua_concat(m, l0, int32(2))
					mBase = m.M
					v172 = m.ExcPending
					if v172 != 0 {
						return int32(0)
					} else {
						v173 = F_lua_error(m, l0)
						mBase = m.M
						v174 = m.ExcPending
						if v174 != 0 {
							return int32(0)
						} else {
							return v173
						}
					}
				}
			}
		}
	}
}
func F_luaB_gcinfo(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v4 = F_lua_gc(m, l0, int32(3), int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v9))) = base.F64_convert_i32_s(v4)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v14 + int32(16)
		return int32(1)
	}
}
func F_luaB_getfenv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
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
	var v210 int32
	_ = v210
	var v211 int64
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	F_getfunc(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v26 = v23 + int32(-16)
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
		if v63 != int32(6) {
			v70 = int32(0)
		} else {
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+6)))
			v70 = base.B2i32(v67 != int32(0))
		}
		if v70 == int32(0) {
			v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v160 = v157 + int32(-16)
			v196 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
			switch v196 + int32(-6) {
			case 0:
				v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v200 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
				v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v199))) = v201
				v216 = int32(5)
				v217 = v199
			case 1:
				v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v205 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
				v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v204))) = v206
				v216 = int32(5)
				v217 = v204
			case 2:
				v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v210 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
				v211 = *(*int64)(unsafe.Add(mBase, uint32(v210)+72))
				*(*int64)(unsafe.Add(mBase, uint32(v209))) = v211
				v213 = *(*int32)(unsafe.Add(mBase, uint32(v210)+80))
				v216 = v213
				v217 = v209
			default:
				v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v216 = int32(0)
				v217 = v214
			}
			*(*int32)(unsafe.Add(mBase, uint32(v217)+8)) = v216
			v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v219 + int32(16)
			return int32(1)
		} else {
			switch int32(0) {
			case 0:
				v126 = l0 + int32(72)
			case 1:
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102
				v126 = l0 + int32(88)
			case 2:
				v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v126 = v96 + int32(96)
			default:
				v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
				v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
				v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+7)))
				v114 = m.G398
				if base.Ui32(v113) < base.Ui32(int32(0)) {
					v125 = v114
				} else {
					v125 = v112 + int32(8)
				}
				v126 = v125
			}
			v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v130 = *(*int64)(unsafe.Add(mBase, uint32(v126)))
			*(*int64)(unsafe.Add(mBase, uint32(v129))) = v130
			v132 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v129)+8)) = v132
			v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v134 + int32(16)
			return int32(1)
		}
	}
}
func F_luaB_ipairs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		switch int32(-1) {
		case 0:
			v61 = l0 + int32(72)
		case 1:
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v37
			v61 = l0 + int32(88)
		case 2:
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v61 = v31 + int32(96)
		default:
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+7)))
			v49 = m.G398
			if base.Ui32(v48) < base.Ui32(int32(1)) {
				v60 = v49
			} else {
				v60 = v47 + int32(24)
			}
			v61 = v60
		}
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v65 = *(*int64)(unsafe.Add(mBase, uint32(v61)))
		*(*int64)(unsafe.Add(mBase, uint32(v64))) = v65
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v67
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v69 + int32(16)
		v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v84 = v79 + int32(0)
		v85 = m.G398
		if base.Ui32(v84) < base.Ui32(v78) {
			v87 = v84
		} else {
			v87 = v85
		}
		v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v130 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
		*(*int64)(unsafe.Add(mBase, uint32(v129))) = v130
		v132 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v129)+8)) = v132
		v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v134 + int32(16)
		v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v140))) = base.F64_convert_i32_s(int32(0))
		v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v145 + int32(16)
		return int32(3)
	}
}
func F_luaB_loadfile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	v2 = int32(0)
	v3 = int32(1)
	v7 = F_luaL_optlstring(m, l0, v3, v2, v2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_luaL_loadfile(m, l0, v7)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v11 == int32(0) {
				v103 = v3
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v16 + int32(16)
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v42 = v39 + int32(-32)
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v78) <= base.Ui32(v42) {
					v95 = v78
				} else {
					v81 = v78
					for {
						v85 = v81 + int32(-16)
						v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
						*(*int64)(unsafe.Add(mBase, uint32(v81))) = v86
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v81+int32(-8))))
						*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v90
						if base.Ui32(v42) < base.Ui32(v85) {
							v81 = v85
							continue
						} else {
							break
						}
						break
					}
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v95 = v93
				}
				v98 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
				*(*int64)(unsafe.Add(mBase, uint32(v42))) = v98
				v100 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v100
				v103 = int32(2)
			}
			return v103
		}
	}
}
func F_luaB_loadstring(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(1)
	v15 = F_luaL_checklstring(m, l0, v10, v8+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = F_luaL_optlstring(m, l0, int32(2), v15, int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			v23 = F_luaL_loadbuffer(m, l0, v15, v22, v20)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					v115 = v10
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v28 + int32(16)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v54 = v51 + int32(-32)
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v90) <= base.Ui32(v54) {
						v107 = v90
					} else {
						v93 = v90
						for {
							v97 = v93 + int32(-16)
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
							*(*int64)(unsafe.Add(mBase, uint32(v93))) = v98
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v93+int32(-8))))
							*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v102
							if base.Ui32(v54) < base.Ui32(v97) {
								v93 = v97
								continue
							} else {
								break
							}
							break
						}
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v107 = v105
					}
					v110 = *(*int64)(unsafe.Add(mBase, uint32(v107)))
					*(*int64)(unsafe.Add(mBase, uint32(v54))) = v110
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v112
					v115 = int32(2)
				}
				m.G0 = v8 + int32(16)
				return v115
			}
		}
	}
}
func F_luaB_newproxy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int64
	_ = v396
	var v400 int32
	_ = v400
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v483 int32
	_ = v483
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = v8 + int32(16)
	if base.Ui32(v11) <= base.Ui32(v7) {
	} else {
		v15 = v7
		for {
			*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(0)
			v19 = v15 + int32(16)
			if base.Ui32(v19) < base.Ui32(v11) {
				v15 = v19
				continue
			} else {
				break
			}
			break
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11
	v32 = F_lua_newuserdata(m, l0, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v46 = v41 + int32(0)
		v47 = m.G398
		if base.Ui32(v46) < base.Ui32(v40) {
			v49 = v46
		} else {
			v49 = v47
		}
		v91 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
		switch v91 {
		case 0:
			v96 = v91
			v98 = v96
		case 1:
			v92 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
			v98 = base.B2i32(v92 != int32(0))
		default:
			v96 = int32(1)
			v98 = v96
		}
		if v98 == int32(0) {
			return int32(1)
		} else {
			v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v109 = v104 + int32(0)
			v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v109) < base.Ui32(v110) {
				v152 = m.G398
				if v109 != v152 {
					v155 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
					v158 = v155
				} else {
					v158 = int32(-1)
				}
			} else {
				v158 = int32(-1)
			}
			if v158 != int32(1) {
				v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v256 = v251 + int32(0)
				v257 = m.G398
				if base.Ui32(v256) < base.Ui32(v250) {
					v259 = v256
				} else {
					v259 = v257
				}
				v301 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
				switch v301 + int32(-5) {
				case 0:
					v304 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
					v316 = v304 + int32(16)
				default:
					v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v316 = v310 + v301<<(uint(int32(2))%32) + int32(152)
				case 2:
					v307 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
					v316 = v307 + int32(8)
				}
				v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
				if v317 != 0 {
					v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v319)+8)) = int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(v319))) = v317
					v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v323 + int32(16)
					v329 = int32(1)
				} else {
					v329 = int32(0)
				}
				if v329 == int32(0) {
					v495 = m.G3
					v498 = F_luaL_argerror(m, l0, int32(1), v495+int32(_a2070))
					mBase = m.M
					v499 = m.ExcPending
					if v499 != 0 {
						return int32(0)
					} else {
						v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v512 = v507 + int32(0)
						v513 = m.G398
						if base.Ui32(v512) < base.Ui32(v506) {
							v515 = v512
						} else {
							v515 = v513
						}
						v557 = *(*int32)(unsafe.Add(mBase, uint32(v515)+8))
						switch v557 + int32(-5) {
						case 0:
							v560 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
							v572 = v560 + int32(16)
						default:
							v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v572 = v566 + v557<<(uint(int32(2))%32) + int32(152)
						case 2:
							v563 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
							v572 = v563 + int32(8)
						}
						v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
						if v573 != 0 {
							v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v575)+8)) = int32(5)
							*(*int32)(unsafe.Add(mBase, uint32(v575))) = v573
							v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v579 + int32(16)
						} else {
						}
						v588 = F_lua_setmetatable(m, l0, int32(2))
						mBase = m.M
						v589 = m.ExcPending
						if v589 != 0 {
							return int32(0)
						} else {
							return int32(1)
						}
					}
				} else {
					switch int32(-1) {
					case 0:
						v385 = l0 + int32(72)
					case 1:
						v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
						v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
						v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v361
						v385 = l0 + int32(88)
					case 2:
						v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v385 = v355 + int32(96)
					default:
						v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
						v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
						v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+7)))
						v373 = m.G398
						if base.Ui32(v372) < base.Ui32(int32(1)) {
							v384 = v373
						} else {
							v384 = v371 + int32(24)
						}
						v385 = v384
					}
					v388 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
					v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v390 = int32(-16)
					v392 = F_luaH_get(m, v388, v389+v390)
					mBase = m.M
					v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v396 = *(*int64)(unsafe.Add(mBase, uint32(v392)))
					*(*int64)(unsafe.Add(mBase, uint32(v393+v390))) = v396
					v400 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v393+int32(-8)))) = v400
					v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v421 = v418 + int32(-16)
					v457 = *(*int32)(unsafe.Add(mBase, uint32(v421)+8))
					switch v457 {
					case 0:
						v462 = v457
						v464 = v462
					case 1:
						v458 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
						v464 = base.B2i32(v458 != int32(0))
					default:
						v462 = int32(1)
						v464 = v462
					}
					v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v483 + int32(-16)
					if v464 != 0 {
						v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v512 = v507 + int32(0)
						v513 = m.G398
						if base.Ui32(v512) < base.Ui32(v506) {
							v515 = v512
						} else {
							v515 = v513
						}
						v557 = *(*int32)(unsafe.Add(mBase, uint32(v515)+8))
						switch v557 + int32(-5) {
						case 0:
							v560 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
							v572 = v560 + int32(16)
						default:
							v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v572 = v566 + v557<<(uint(int32(2))%32) + int32(152)
						case 2:
							v563 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
							v572 = v563 + int32(8)
						}
						v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
						if v573 != 0 {
							v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v575)+8)) = int32(5)
							*(*int32)(unsafe.Add(mBase, uint32(v575))) = v573
							v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v579 + int32(16)
						} else {
						}
						v588 = F_lua_setmetatable(m, l0, int32(2))
						mBase = m.M
						v589 = m.ExcPending
						if v589 != 0 {
							return int32(0)
						} else {
							return int32(1)
						}
					} else {
						v495 = m.G3
						v498 = F_luaL_argerror(m, l0, int32(1), v495+int32(_a2070))
						mBase = m.M
						v499 = m.ExcPending
						if v499 != 0 {
							return int32(0)
						} else {
							v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v512 = v507 + int32(0)
							v513 = m.G398
							if base.Ui32(v512) < base.Ui32(v506) {
								v515 = v512
							} else {
								v515 = v513
							}
							v557 = *(*int32)(unsafe.Add(mBase, uint32(v515)+8))
							switch v557 + int32(-5) {
							case 0:
								v560 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
								v572 = v560 + int32(16)
							default:
								v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v572 = v566 + v557<<(uint(int32(2))%32) + int32(152)
							case 2:
								v563 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
								v572 = v563 + int32(8)
							}
							v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
							if v573 != 0 {
								v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v575)+8)) = int32(5)
								*(*int32)(unsafe.Add(mBase, uint32(v575))) = v573
								v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v579 + int32(16)
							} else {
							}
							v588 = F_lua_setmetatable(m, l0, int32(2))
							mBase = m.M
							v589 = m.ExcPending
							if v589 != 0 {
								return int32(0)
							} else {
								return int32(1)
							}
						}
					}
				}
			} else {
				v161 = int32(0)
				F_lua_createtable(m, l0, v161, v161)
				mBase = m.M
				v164 = m.ExcPending
				if v164 != 0 {
					return int32(0)
				} else {
					v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v185 = v182 + int32(-16)
					v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v222 = *(*int64)(unsafe.Add(mBase, uint32(v185)))
					*(*int64)(unsafe.Add(mBase, uint32(v221))) = v222
					v224 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v221)+8)) = v224
					v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v226 + int32(16)
					v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v232)+8)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v232))) = int32(1)
					v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v238 + int32(16)
					F_lua_rawset(m, l0, int32(-10003))
					mBase = m.M
					v244 = m.ExcPending
					if v244 != 0 {
						return int32(0)
					} else {
						v588 = F_lua_setmetatable(m, l0, int32(2))
						mBase = m.M
						v589 = m.ExcPending
						if v589 != 0 {
							return int32(0)
						} else {
							return int32(1)
						}
					}
				}
			}
		}
	}
}
func F_luaB_next(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = v15 + int32(32)
		if base.Ui32(v18) <= base.Ui32(v14) {
		} else {
			v22 = v14
			for {
				*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(0)
				v26 = v22 + int32(16)
				if base.Ui32(v26) < base.Ui32(v18) {
					v22 = v26
					continue
				} else {
					break
				}
				break
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v18
		v39 = F_lua_next(m, l0, int32(1))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			if v39 != 0 {
				v49 = int32(2)
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v42 + int32(16)
				v49 = int32(1)
			}
			return v49
		}
	}
}
func F_luaB_pcall(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	F_luaL_checkany(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v12 = int32(-1)
		v16 = F_lua_pcall(m, l0, (v7-v8)>>(uint(int32(4))%32)+v12, v12, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = int32(0)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v21))) = base.B2i32(base.B2i32(v16 == v18) != v18)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27 + int32(16)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v42 = v37 + int32(0)
			v43 = m.G398
			if base.Ui32(v42) < base.Ui32(v36) {
				v45 = v42
			} else {
				v45 = v43
			}
			v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v87) <= base.Ui32(v45) {
				v104 = v87
			} else {
				v90 = v87
				for {
					v94 = v90 + int32(-16)
					v95 = *(*int64)(unsafe.Add(mBase, uint32(v94)))
					*(*int64)(unsafe.Add(mBase, uint32(v90))) = v95
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v90+int32(-8))))
					*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = v99
					if base.Ui32(v45) < base.Ui32(v94) {
						v90 = v94
						continue
					} else {
						break
					}
					break
				}
				v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v104 = v102
			}
			v107 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
			*(*int64)(unsafe.Add(mBase, uint32(v45))) = v107
			v109 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v109
			v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			return (v111 - v112) >> (uint(int32(4)) % 32)
		}
	}
}
func F_luaB_select(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = (v4 - v5) >> (uint(int32(4)) % 32)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = v12 + int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v17) < base.Ui32(v18) {
		v60 = m.G398
		if v17 != v60 {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
			v66 = v63
		} else {
			v66 = int32(-1)
		}
	} else {
		v66 = int32(-1)
	}
	if v66 != int32(4) {
		v93 = F_luaL_checkinteger(m, l0, int32(1))
		mBase = m.M
		v94 = m.ExcPending
		if v94 != 0 {
			return int32(0)
		} else {
			if v93 < v8 {
				v97 = v93
			} else {
				v97 = v8
			}
			if v93 < int32(0) {
				v100 = v93 + v8
			} else {
				v100 = v97
			}
			if int32(0) < v100 {
				return v8 - v100
			} else {
				v104 = m.G3
				v107 = F_luaL_argerror(m, l0, int32(1), v104+int32(_a2072))
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return int32(0)
				} else {
					return v8 - v100
				}
			}
		}
	} else {
		v71 = F_lua_tolstring(m, l0, int32(1), int32(0))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return int32(0)
		} else {
			v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
			if v75 != int32(35) {
				v93 = F_luaL_checkinteger(m, l0, int32(1))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					if v93 < v8 {
						v97 = v93
					} else {
						v97 = v8
					}
					if v93 < int32(0) {
						v100 = v93 + v8
					} else {
						v100 = v97
					}
					if int32(0) < v100 {
						return v8 - v100
					} else {
						v104 = m.G3
						v107 = F_luaL_argerror(m, l0, int32(1), v104+int32(_a2072))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							return v8 - v100
						}
					}
				}
			} else {
				v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = int32(3)
				*(*float64)(unsafe.Add(mBase, uint32(v81))) = base.F64_convert_i32_s(v8 + int32(-1))
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v86 + int32(16)
				return int32(1)
			}
		}
	}
}
func F_luaB_setmetatable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = v6 + int32(16)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v11) < base.Ui32(v12) {
		v54 = m.G398
		if v11 != v54 {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			v60 = v57
		} else {
			v60 = int32(-1)
		}
	} else {
		v60 = int32(-1)
	}
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		return int32(0)
	} else {
		switch v60 {
		case 0, 5:
			v74 = m.G3
			v77 = F_luaL_getmetafield(m, l0, int32(1), v74+int32(_a2073))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				if v77 == int32(0) {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v95 = v92 + int32(32)
					if base.Ui32(v95) <= base.Ui32(v91) {
					} else {
						v99 = v91
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = int32(0)
							v103 = v99 + int32(16)
							if base.Ui32(v103) < base.Ui32(v95) {
								v99 = v103
								continue
							} else {
								break
							}
							break
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v95
					v116 = F_lua_setmetatable(m, l0, int32(1))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return int32(0)
					} else {
						return int32(1)
					}
				} else {
					v81 = m.G3
					v85 = F_luaL_error(m, l0, v81+int32(_a2074), int32(0))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v95 = v92 + int32(32)
						if base.Ui32(v95) <= base.Ui32(v91) {
						} else {
							v99 = v91
							for {
								*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = int32(0)
								v103 = v99 + int32(16)
								if base.Ui32(v103) < base.Ui32(v95) {
									v99 = v103
									continue
								} else {
									break
								}
								break
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v95
						v116 = F_lua_setmetatable(m, l0, int32(1))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return int32(0)
						} else {
							return int32(1)
						}
					}
				}
			}
		default:
			v68 = m.G3
			v71 = F_luaL_argerror(m, l0, int32(2), v68+int32(_a2075))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				v74 = m.G3
				v77 = F_luaL_getmetafield(m, l0, int32(1), v74+int32(_a2073))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					if v77 == int32(0) {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v95 = v92 + int32(32)
						if base.Ui32(v95) <= base.Ui32(v91) {
						} else {
							v99 = v91
							for {
								*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = int32(0)
								v103 = v99 + int32(16)
								if base.Ui32(v103) < base.Ui32(v95) {
									v99 = v103
									continue
								} else {
									break
								}
								break
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v95
						v116 = F_lua_setmetatable(m, l0, int32(1))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return int32(0)
						} else {
							return int32(1)
						}
					} else {
						v81 = m.G3
						v85 = F_luaL_error(m, l0, v81+int32(_a2074), int32(0))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v95 = v92 + int32(32)
							if base.Ui32(v95) <= base.Ui32(v91) {
							} else {
								v99 = v91
								for {
									*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = int32(0)
									v103 = v99 + int32(16)
									if base.Ui32(v103) < base.Ui32(v95) {
										v99 = v103
										continue
									} else {
										break
									}
									break
								}
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v95
							v116 = F_lua_setmetatable(m, l0, int32(1))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								return int32(1)
							}
						}
					}
				}
			}
		}
	}
}
func F_luaB_unpack(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
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
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v13 = F_luaL_optinteger(m, l0, int32(2), int32(1))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v23 = v18 + int32(32)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v23) < base.Ui32(v24) {
				v66 = m.G398
				if v23 != v66 {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
					v72 = v69
				} else {
					v72 = int32(-1)
				}
			} else {
				v72 = int32(-1)
			}
			if int32(0) < v72 {
				v79 = F_luaL_checkinteger(m, l0, int32(3))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					v81 = v79
					if v81 < v13 {
						v244 = int32(0)
						return v244
					} else {
						v84 = v81 - v13
						if base.Ui32(int32(2147483646)) < base.Ui32(v84) {
							v92 = m.G3
							v96 = F_luaL_error(m, l0, v92+int32(_a2076), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								return v96
							}
						} else {
							v88 = v84 + int32(1)
							v89 = F_lua_checkstack(m, l0, v88)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								if v89 != 0 {
									v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v110 = v105 + int32(0)
									v111 = m.G398
									if base.Ui32(v110) < base.Ui32(v104) {
										v113 = v110
									} else {
										v113 = v111
									}
									v155 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
									v156 = F_luaH_getnum(m, v155, v13)
									mBase = m.M
									v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v158 = *(*int64)(unsafe.Add(mBase, uint32(v156)))
									*(*int64)(unsafe.Add(mBase, uint32(v157))) = v158
									v160 = *(*int32)(unsafe.Add(mBase, uint32(v156)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v157)+8)) = v160
									v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v162 + int32(16)
									if v81 <= v13 {
										v244 = v88
									} else {
										v168 = v13
										for {
											v173 = v168 + int32(1)
											v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v184 = v179 + int32(0)
											v185 = m.G398
											if base.Ui32(v184) < base.Ui32(v178) {
												v187 = v184
											} else {
												v187 = v185
											}
											v229 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
											v230 = F_luaH_getnum(m, v229, v173)
											mBase = m.M
											v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v232 = *(*int64)(unsafe.Add(mBase, uint32(v230)))
											*(*int64)(unsafe.Add(mBase, uint32(v231))) = v232
											v234 = *(*int32)(unsafe.Add(mBase, uint32(v230)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v231)+8)) = v234
											v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v236 + int32(16)
											if v173 != v81 {
												v168 = v173
												continue
											} else {
												break
											}
											break
										}
										v244 = v88
									}
									return v244
								} else {
									v92 = m.G3
									v96 = F_luaL_error(m, l0, v92+int32(_a2076), int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										return v96
									}
								}
							}
						}
					}
				}
			} else {
				v76 = F_lua_objlen(m, l0, int32(1))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					v81 = v76
					if v81 < v13 {
						v244 = int32(0)
						return v244
					} else {
						v84 = v81 - v13
						if base.Ui32(int32(2147483646)) < base.Ui32(v84) {
							v92 = m.G3
							v96 = F_luaL_error(m, l0, v92+int32(_a2076), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								return v96
							}
						} else {
							v88 = v84 + int32(1)
							v89 = F_lua_checkstack(m, l0, v88)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								if v89 != 0 {
									v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v110 = v105 + int32(0)
									v111 = m.G398
									if base.Ui32(v110) < base.Ui32(v104) {
										v113 = v110
									} else {
										v113 = v111
									}
									v155 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
									v156 = F_luaH_getnum(m, v155, v13)
									mBase = m.M
									v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v158 = *(*int64)(unsafe.Add(mBase, uint32(v156)))
									*(*int64)(unsafe.Add(mBase, uint32(v157))) = v158
									v160 = *(*int32)(unsafe.Add(mBase, uint32(v156)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v157)+8)) = v160
									v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v162 + int32(16)
									if v81 <= v13 {
										v244 = v88
									} else {
										v168 = v13
										for {
											v173 = v168 + int32(1)
											v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v184 = v179 + int32(0)
											v185 = m.G398
											if base.Ui32(v184) < base.Ui32(v178) {
												v187 = v184
											} else {
												v187 = v185
											}
											v229 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
											v230 = F_luaH_getnum(m, v229, v173)
											mBase = m.M
											v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v232 = *(*int64)(unsafe.Add(mBase, uint32(v230)))
											*(*int64)(unsafe.Add(mBase, uint32(v231))) = v232
											v234 = *(*int32)(unsafe.Add(mBase, uint32(v230)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v231)+8)) = v234
											v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v236 + int32(16)
											if v173 != v81 {
												v168 = v173
												continue
											} else {
												break
											}
											break
										}
										v244 = v88
									}
									return v244
								} else {
									v92 = m.G3
									v96 = F_luaL_error(m, l0, v92+int32(_a2076), int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										return v96
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
