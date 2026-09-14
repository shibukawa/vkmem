package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F_luaopen_base(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	switch int32(0) {
	case 0:
		v57 = l0 + int32(72)
	case 1:
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v33
		v57 = l0 + int32(88)
	case 2:
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v57 = v27 + int32(96)
	default:
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
		v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+7)))
		v45 = m.G398
		if base.Ui32(v44) < base.Ui32(int32(0)) {
			v56 = v45
		} else {
			v56 = v43 + int32(8)
		}
		v57 = v56
	}
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	*(*int64)(unsafe.Add(mBase, uint32(v60))) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v65 + int32(16)
	v70 = m.G3
	v72 = v70 + int32(_a2194)
	F_lua_setfield(m, l0, int32(-10002), v72)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		return int32(0)
	} else {
		F_luaL_register(m, l0, v72, v70+int32(_a2291))
		mBase = m.M
		v80 = m.ExcPending
		if v80 != 0 {
			return int32(0)
		} else {
			F_lua_pushlstring(m, l0, v70+int32(_a2292), int32(7))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return int32(0)
			} else {
				F_lua_setfield(m, l0, int32(-10002), v70+int32(_a2201))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					v91 = m.G5
					F_lua_pushcclosure(m, l0, v91+int32(1238), int32(0))
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int32(0)
					} else {
						F_lua_pushcclosure(m, l0, v91+int32(1239), int32(1))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							F_lua_setfield(m, l0, int32(-2), v70+int32(_a2200))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								F_lua_pushcclosure(m, l0, v91+int32(1240), int32(0))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									F_lua_pushcclosure(m, l0, v91+int32(1241), int32(1))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										F_lua_setfield(m, l0, int32(-2), v70+int32(_a2197))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											F_lua_createtable(m, l0, int32(0), int32(1))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v146 = v143 + int32(-16)
												v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v183 = *(*int64)(unsafe.Add(mBase, uint32(v146)))
												*(*int64)(unsafe.Add(mBase, uint32(v182))) = v183
												v185 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v182)+8)) = v185
												v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v187 + int32(16)
												v192 = F_lua_setmetatable(m, l0, int32(-2))
												mBase = m.M
												v193 = m.ExcPending
												if v193 != 0 {
													return int32(0)
												} else {
													F_lua_pushlstring(m, l0, v70+int32(_a2293), int32(2))
													mBase = m.M
													v198 = m.ExcPending
													if v198 != 0 {
														return int32(0)
													} else {
														F_lua_setfield(m, l0, int32(-2), v70+int32(_a2294))
														mBase = m.M
														v203 = m.ExcPending
														if v203 != 0 {
															return int32(0)
														} else {
															F_lua_pushcclosure(m, l0, v91+int32(1242), int32(1))
															mBase = m.M
															v208 = m.ExcPending
															if v208 != 0 {
																return int32(0)
															} else {
																F_lua_setfield(m, l0, int32(-10002), v70+int32(_a2203))
																mBase = m.M
																v213 = m.ExcPending
																if v213 != 0 {
																	return int32(0)
																} else {
																	F_luaL_register(m, l0, v70+int32(_a2193), v70+int32(_a2295))
																	mBase = m.M
																	v219 = m.ExcPending
																	if v219 != 0 {
																		return int32(0)
																	} else {
																		return int32(2)
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
					}
				}
			}
		}
	}
}
func F_luaopen_math(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v3 = m.G3
	F_luaL_register(m, l0, v3+int32(_a2177), v3+int32(_a2315))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v14))) = float64(3.141592653589793)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v18 + int32(16)
		F_lua_setfield(m, l0, int32(-2), v3+int32(_a2316))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v29))) = math.Float64frombits(uint64(0x7ff0000000000000))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v33 + int32(16)
			F_lua_setfield(m, l0, int32(-2), v3+int32(_a2317))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				F_lua_getfield(m, l0, int32(-1), v3+int32(_a2318))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_lua_setfield(m, l0, int32(-2), v3+int32(_a2319))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						return int32(1)
					}
				}
			}
		}
	}
}
