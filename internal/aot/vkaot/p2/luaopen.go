package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaopen_bit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 float64
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v11))) = float64(1.437217655e+09)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15 + v6
	v20 = F_lua_tonumber(m, l0, int32(-1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		v27 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v20, float64(6.755399441055744e+15))))
		if v27 == int32(1437217655) {
			v52 = m.G3
			F_luaL_register(m, l0, v52+int32(_a_F_luaopen_bit_0), v52+int32(_a_F_luaopen_bit_1))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return int32(1)
			}
		} else {
			if v27 != 0 {
				v39 = m.G3
				if v27 == int32(1127743488) {
					v44 = int32(_a_F_luaopen_bit_2)
				} else {
					v44 = int32(_a_F_luaopen_bit_3)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v39 + v44
				v49 = F_luaL_error(m, l0, v39+int32(_a_F_luaopen_bit_4), v7)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v52 = m.G3
					F_luaL_register(m, l0, v52+int32(_a_F_luaopen_bit_0), v52+int32(_a_F_luaopen_bit_1))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return int32(1)
					}
				}
			} else {
				v31 = F_lua_isnumber(m, l0, int32(-1))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v31 != 0 {
						v39 = m.G3
						if v27 == int32(1127743488) {
							v44 = int32(_a_F_luaopen_bit_2)
						} else {
							v44 = int32(_a_F_luaopen_bit_3)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v39 + v44
						v49 = F_luaL_error(m, l0, v39+int32(_a_F_luaopen_bit_4), v7)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							v52 = m.G3
							F_luaL_register(m, l0, v52+int32(_a_F_luaopen_bit_0), v52+int32(_a_F_luaopen_bit_1))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(16)
								return int32(1)
							}
						}
					} else {
						v34 = m.G3
						v37 = F_luaL_typerror(m, l0, int32(-1), v34+int32(_a_F_luaopen_bit_5))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = m.G3
							if v27 == int32(1127743488) {
								v44 = int32(_a_F_luaopen_bit_2)
							} else {
								v44 = int32(_a_F_luaopen_bit_3)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v39 + v44
							v49 = F_luaL_error(m, l0, v39+int32(_a_F_luaopen_bit_4), v7)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v52 = m.G3
								F_luaL_register(m, l0, v52+int32(_a_F_luaopen_bit_0), v52+int32(_a_F_luaopen_bit_1))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(16)
									return int32(1)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_luaopen_create(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v2 = int32(0)
	F_lua_createtable(m, l0, v2, v2)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = m.G404
		F_lua_pushcclosure(m, l0, v9, int32(0))
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v14 = m.G3
			F_lua_setfield(m, l0, int32(-2), v14+int32(_a_F_luaopen_create_0))
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = m.G405
				F_lua_pushcclosure(m, l0, v19, int32(0))
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_lua_setfield(m, l0, int32(-2), v14+int32(_a_F_luaopen_create_1))
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = m.G406
						F_lua_pushcclosure(m, l0, v28, int32(0))
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_lua_setfield(m, l0, int32(-2), v14+int32(_a_F_luaopen_create_2))
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v37 = m.G407
								F_lua_pushcclosure(m, l0, v37, int32(0))
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									F_lua_setfield(m, l0, int32(-2), v14+int32(_a_F_luaopen_create_3))
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										F_lua_pushlstring(m, l0, v14+int32(_a_F_luaopen_create_4), int32(8))
										v50 = m.ExcPending
										if v50 != 0 {
											return int32(0)
										} else {
											F_lua_setfield(m, l0, int32(-2), v14+int32(_a_F_luaopen_create_5))
											v55 = m.ExcPending
											if v55 != 0 {
												return int32(0)
											} else {
												F_lua_pushlstring(m, l0, v14+int32(_a_F_luaopen_create_6), int32(18))
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													F_lua_setfield(m, l0, int32(-2), v14+int32(_a_F_luaopen_create_7))
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														F_lua_pushlstring(m, l0, v14+int32(_a_F_luaopen_create_8), int32(30))
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															F_lua_setfield(m, l0, int32(-2), v14+int32(_a_F_luaopen_create_9))
															v75 = m.ExcPending
															if v75 != 0 {
																return int32(0)
															} else {
																F_lua_pushlstring(m, l0, v14+int32(_a_F_luaopen_create_10), int32(36))
																v80 = m.ExcPending
																if v80 != 0 {
																	return int32(0)
																} else {
																	F_lua_setfield(m, l0, int32(-2), v14+int32(_a_F_luaopen_create_11))
																	v85 = m.ExcPending
																	if v85 != 0 {
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
func F_luaopen_debug(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	v3 = m.G3
	F_luaL_register(m, l0, v3+int32(_a_F_luaopen_debug_0), v3+int32(_a_F_luaopen_debug_1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_luaopen_os(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	v3 = m.G3
	F_luaL_register(m, l0, v3+int32(_a_F_luaopen_os_0), v3+int32(_a_F_luaopen_os_1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_luaopen_struct(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	v3 = m.G3
	F_luaL_register(m, l0, v3+int32(_a_F_luaopen_struct_0), v3+int32(_a_F_luaopen_struct_1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
